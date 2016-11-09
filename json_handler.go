package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var (
	jsonFileDir  string
	jsonFilePath string
)

func init() {
	const filename = "komi.json"
	defaultDataDir := path.Join("/home", os.Getenv("USER"), ".komi")
	jsonFileDir = os.Getenv("KOMI_DATA_DIR")

	// If Env Variable not set
	if jsonFileDir == "" {
		jsonFileDir = defaultDataDir
	}

	jsonFilePath = path.Join(jsonFileDir, filename)
}

// Read and parse json data from `jsonFilePath`
// If an error occurs, return it
func readJSONData() (jsonData, error) {
	jd := jsonData{}

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		// return jd, &jsonFileError{err, "Error reading JSON file: " + jsonFilePath}

		// Create file
		err := setupDataDir(jd)
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

	err = ioutil.WriteFile(jsonFilePath, content, 0644)
	if err != nil {
		return &jsonFileError{err, "Error writing to JSON file: " + jsonFilePath}
	}

	return nil
}

// Create data file with jd data
func setupDataDir(jd jsonData) error {
	// Create data directory
	err := os.MkdirAll(jsonFileDir, 0744)
	if err != nil {
		return err
	}

	// Create data file
	err = writeJSONData(jd)
	if err != nil {
		return err
	}

	return nil
}
