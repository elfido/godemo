package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// FileStorage comment...
type FileStorage struct {
}

// Add comment...
func (f *FileStorage) Add(entityID string, program Program) error {
	filename := fmt.Sprintf("%s.json", entityID)
	jsonBytes, err := json.Marshal(program)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
