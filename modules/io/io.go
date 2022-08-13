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
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func GetFileLinesNoTrim(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func GetFileContent(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func WriteToFile(filePath string, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(content)
}
