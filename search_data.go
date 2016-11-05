package main

import (
	"fmt"
	"strings"
)

// Returns lower cased string if ignoreCase is true
// Else returns same string
func corvo(st string, ignoreCase bool) string {
	if ignoreCase {
		return strings.ToLower(st)
	}

	return st
}

func search(jd jsonData, query string, ignoreCase bool) error {
	q := corvo(query, ignoreCase)

	for _, ct := range jd.Categories {
		ctName := corvo(ct.Name, ignoreCase)

		if strings.Contains(ctName, q) {
			cat := strings.Replace(ct.Name, query, highlightUse(query), -1)
			fmt.Printf("[Category: %s]\n", cat)
		}

		for _, cm := range ct.Commands {
			cmName := corvo(cm.Name, ignoreCase)
			cmUse := corvo(cm.Use, ignoreCase)

			if strings.Contains(cmName, q) || strings.Contains(cmUse, q) {
				name := strings.Replace(cm.coloredName(), query, highlightName(query), -1)
				use := strings.Replace(cm.Use, query, highlightUse(query), -1)

				fmt.Print(name)
				if use != "" {
					fmt.Print(" →  ", use)
				}
				fmt.Println()
			}
		}
	}

	return nil
}

func highlightName(st string) string {
	red := "\x1b[31m"
	yellow := "\x1b[33m"
	return fmt.Sprintf("%s%s%s", red, st, yellow)
}

func highlightUse(st string) string {
	red := "\x1b[31m"
	reset := "\x1b[0m"
	return fmt.Sprintf("%s%s%s", red, st, reset)
}
