package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var jsonFileDir = os.Getenv("KOMI_DATA_DIR")
var jsonFilePath string

func init() {
	filename := "komi.json"
	// If Env Variable not set
	if jsonFileDir == "" {
		jsonFilePath = "/home/" + os.Getenv("USER") + "/" + filename
	} else {
		jsonFilePath = jsonFileDir + "/" + filename
	}
}

// Read and parse json data from `jsonFilePath`
// If an error occurs, return it
func readJSONData() (jsonData, error) {
	jd := jsonData{}

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		// return jd, &jsonFileError{err, "Error reading JSON file: " + jsonFilePath}

		// Create file
		err := createDataFile(jd)
		if err != nil {
			return jd, err
		}

		// Skip parsing JSON since data is going to be `jd`
		return jd, nil
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

// Create data file with jd data
func createDataFile(jd jsonData) error {
	err := writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
