package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Returns list of category names
func (jd jsonData) GetCategoryList() []string {
	var categories []string

	if jd.Categories == nil {
		return categories
	}

	for _, c := range jd.Categories {
		categories = append(categories, c.Name)
	}

	return categories
}

// Returns list of commands and true  if category was found
// else it returns an empty commands list with false
func (jd jsonData) GetCommands(catgr string) ([]command, bool) {
	var commands []command

	if jd.Categories == nil {
		return commands, false
	}

	for _, c := range jd.Categories {
		if c.Name == catgr {
			return c.Commands, true
		}
	}

	return commands, false
}

// Add a category
func (jd *jsonData) AddCategory(c category) {
	jd.Categories = append(jd.Categories, c)
}

// Add command to a category
func (jd *jsonData) AddCommand(categoryIndex int, com command) {
	jd.Categories[categoryIndex].Commands = append(jd.Categories[categoryIndex].Commands, com)
}

// Modify a category's name
func (jd *jsonData) ModifyCategory(categoryIndex int, newName string) {
	jd.Categories[categoryIndex].Name = newName
}

// Modify a command's name and use
func (jd *jsonData) ModifyCommand(categoryIndex int, commandIndex int, newName string, newUse string) {
	com := &jd.Categories[categoryIndex].Commands[commandIndex]

	(*com).Name = newName
	(*com).Use = newUse
}

// Delete a category
func (jd *jsonData) DeleteCategory(categoryIndex int) {
	cats := &jd.Categories

	l := len(*cats)
	// Swap with last category
	(*cats)[categoryIndex] = (*cats)[l-1]
	// Slice off last category
	*cats = (*cats)[:l-1]
}

// Delete a command
func (jd *jsonData) DeleteCommand(categoryIndex int, commandIndex int) {
	coms := &jd.Categories[categoryIndex].Commands

	l := len(*coms)
	(*coms)[commandIndex] = (*coms)[l-1]
	*coms = (*coms)[:l-1]
}

// Returns index of a category name
// Returns -1 if category name not found
func (jd jsonData) getIndex(catgr string) (int, error) {
	for i, c := range jd.Categories {
		if c.Name == catgr {
			return i, nil
		}
	}

	msg := fmt.Sprintf("No such category as `%s`", catgr)
	return -1, errors.New(msg)
}

// Prints category names for Bash completion
func (jd jsonData) printCategories() {
	for _, c := range jd.Categories {
		fmt.Println(c.Name)
	}
}

// Returns colored name for a command
func (c command) coloredName() string {
	yellow := "\x1b[33m"
	reset := "\x1b[0m"
	return fmt.Sprintf("%s%s%s", yellow, c.Name, reset)
}

// Read Serial Number and validate
func readSerialNo(upperBound int) (int, error) {
	input, err := readInput("Select Serial no: ", false)
	if err != nil {
		return -1, err
	}
	sno, err := strconv.Atoi(input)
	if err != nil || sno < 1 || sno > upperBound {
		return -1, errors.New("Invalid Serial no")
	}

	return sno, nil
}
