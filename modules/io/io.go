package io

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			lines = append(lines, strings.TrimSpace(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
