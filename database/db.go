package database

import (
	"encoding/json"
	"errors"
	"os"
)

type DB struct {
	Data []dbItem `json:"data"`
}

type dbItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (t *DB) Set(key string, val string) error {
	for _, val := range t.Data {
		if val.Key == key {
			return errors.New("Key already exists")
		}
	}

	item := dbItem{Key: key, Value: val}
	t.Data = append(t.Data, item)
	return nil
}

func (t *DB) Get(key string) (error, dbItem) {
	for _, val := range t.Data {
		if val.Key == key {
			return nil, val
		}
	}

	return errors.New("Key not found"), dbItem{}
}

func (t *DB) Delete(key string) error {
	index := -1
	for ind, val := range t.Data {
		if val.Key == key {
			index = ind
		}
	}
	ls := *t
	if index < 0 || index > len(ls.Data) {
		return errors.New("Key not found")
	}

	t.Data = append(ls.Data[:index], ls.Data[index+1:]...)

	return nil
}

func (t *DB) List() {
  for _, val := range t.Data {
    println(val.Key, val.Value)
  }
}

func (t *DB) Load(filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *DB) Store(filePath string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
