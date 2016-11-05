package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Display categories in a square matrix
func displayCategories(jd jsonData) {
	categories := jd.GetCategoryList()
	l := len(categories)
	// Length of category with max length
	var maxCatLen int

	for _, c := range categories {
		if maxCatLen < len(c) {
			maxCatLen = len(c)
		}
	}

	// Calculate required order for square matrix
	n := 1
	for l > n*n {
		n++
	}

	// Display categories
	for i, c := range categories {
		fmt.Print(c)
		for j := 0; j <= maxCatLen-len(c); j++ {
			fmt.Print(" ")
		}
		if (i+1)%n == 0 {
			fmt.Println()
		}
	}

	// Required newline if categories do not form a
	// complete square matrix
	if l%n != 0 {
		fmt.Println()
	}
}

// Display commands and uses
// Return error if category not found
func displayCommands(jd jsonData, catgr string, namesOnly bool) error {
	coms, found := jd.GetCommands(catgr)

	if !found {
		msg := fmt.Sprintf("No such category as `%s`\n", catgr)
		return errors.New(msg)
	}

	if namesOnly {
		// Print serialized command names and return
		for i, c := range coms {
			fmt.Printf("%-2d %s\n", i+1, c.coloredName())
		}
		return nil
	}

	for _, c := range coms {
		fmt.Print(c.coloredName())
		if c.Use != "" {
			fmt.Print(" →  ", c.Use)
		}
		fmt.Println()
	}

	return nil
}

func displayJSONData(jd jsonData) error {
	content, err := json.MarshalIndent(jd, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}
