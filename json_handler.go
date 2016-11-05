package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var jsonFilePath = os.Getenv("KOMI_DATA_FILE")

func init() {
	// If Env Variable not set
	if jsonFilePath == "" {
		jsonFilePath = "/home/" + os.Getenv("USER") + "/komi.json"
	}
}

// Read and parse json data from `jsonFilePath`
// If an error occurs, return it
func readJSONData() (jsonData, error) {
	jd := jsonData{}

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return jd, &jsonFileError{err, "Error reading JSON file: " + jsonFilePath}
	}

	err = json.Unmarshal(content, &jd)
	if err != nil {
		return jd, &jsonFileError{err, "Error parsing JSON file: " + jsonFilePath}
	}

	return jd, nil
}

// Parse and write json data to `jsonFilePath`
// If an error occurs, return it
func writeJSONData(jd jsonData) error {
	content, err := json.Marshal(jd)
	if err != nil {
		return &jsonFileError{err, "Error parsing JSON file: " + jsonFilePath}
	}

	err = ioutil.WriteFile(jsonFilePath, content, 0666)
	if err != nil {
		return &jsonFileError{err, "Error writing to JSON file: " + jsonFilePath}
	}

	return nil
}
