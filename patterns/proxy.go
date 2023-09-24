package patterns

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type Consumer struct {
	cachedDatabase cache
}

func NewConsumer() *Consumer {
	return &Consumer{cachedDatabase: *NewCache()}
}

func (consumer *Consumer) Save(input Input) {
	consumer.cachedDatabase.Set(input)
}

func (consumer *Consumer) Get(id string) string {
	return consumer.cachedDatabase.Get(id)
}

type IDatabase interface {
	Set(input Input)
	Get(id string) string
}

type Input struct {
	Id      string
	Content string
}

type Database struct {
	lock sync.RWMutex
	File string
	data map[string]string
}

func NewDatabase(file string) *Database {
	return &Database{File: file, data: make(map[string]string)}
}

func (db *Database) Set(input Input) {
	db.lock.Lock()
	defer db.lock.Unlock()
	dbBytes, err := os.ReadFile(db.File)
	if err != nil {
		log.Fatalf("cannot read database file %s", db.File)
	}
	json.Unmarshal(dbBytes, &db.data)
	db.data[input.Id] = input.Content
	updatedDBBytes, err := json.Marshal(db.data)
	if err != nil {
		log.Fatal("cannot marshal data")
	}
	os.WriteFile(db.File, updatedDBBytes, os.ModePerm)
}

func (db *Database) Get(id string) string {
	dbBytes, _ := os.ReadFile(db.File)
	var data map[string]string
	json.Unmarshal(dbBytes, &data)
	if content, ok := data[id]; ok {
		return content
	}
	return "NA"
}

type cache struct {
	database   Database
	cacheddata map[string]cachedData
}

func NewCache() *cache {
	return &cache{
		database:   *NewDatabase("data/db.json"),
		cacheddata: make(map[string]cachedData),
	}
}

func (c *cache) Set(input Input) {
	c.database.Set(input)
}

func (c *cache) Get(id string) string {
	value, ok := c.cacheddata[id]
	cacheExpiredButFound := ok && time.Now().After(value.ExpiredAt)
	if !ok || cacheExpiredButFound {
		// retrieve the content from database and cache it.
		content := c.database.Get(id)
		log.Printf("Cached value, %s", content)
		c.cacheddata[id] = cachedData{
			ExpiredAt: time.Now().Add(3 * time.Second),
			Content:   content,
		}
	}
	return value.Content
}

type cachedData struct {
	ExpiredAt time.Time
	Content   string
}

func Proxy() {
	consumer := NewConsumer()
	input := Input{Id: "1", Content: "hello"}
	consumer.Save(input)
	result := consumer.Get("2")
	log.Print(result)
}
