package patterns

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Consumer struct{}

func (consumer Consumer) process(database IDatabase) {
	database.Set(Input{
		Id:      "1",
		Content: "hello",
	})

	result := database.Get("1")
	log.Println(result)
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

func Proxy() {
	consumer := Consumer{}
	db := NewDatabase("data/db.json")
	consumer.process(db)
}
