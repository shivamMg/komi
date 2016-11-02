package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var jsonFileName = "komi.json"
var jsonFilePath = "/home/" + os.Getenv("USER") + "/" + jsonFileName

// Read and parse json data from `jsonFilePath`
// If an error occurs, return it
func readJSONData() (jsonData, *jsonFileError) {
	jd := jsonData{}

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return jd, &jsonFileError{err, "Error reading JSON file: " + jsonFilePath}
	}

	err = json.Unmarshal(content, &jd)
	if err != nil {
		return jd, &jsonFileError{err, "Error parsing JSON. File: " + jsonFilePath}
	}

	return jd, nil
}

// Parse and write json data to `jsonFilePath`
// If an error occurs, return it
func writeJSONData(jd jsonData) *jsonFileError {
	content, err := json.Marshal(jd)
	if err != nil {
		return &jsonFileError{err, "Error parsing JSON. File: " + jsonFilePath}
	}

	err = ioutil.WriteFile(jsonFilePath, content, 0666)
	if err != nil {
		return &jsonFileError{err, "Error writing to JSON file: " + jsonFilePath}
	}

	return nil
}
