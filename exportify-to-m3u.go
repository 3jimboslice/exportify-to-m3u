package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func processCSV(filePath string) error {
	// Open the CSV file
	csvFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file %s: %w", filePath, err)
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file %s: %w", filePath, err)
	}

	// Ensure there are records in the CSV file
	if len(records) == 0 {
		return fmt.Errorf("CSV file %s is empty", filePath)
	}

	// Create the output file name by replacing .csv with .m3u
	outputFileName := strings.TrimSuffix(filepath.Base(filePath), ".csv") + ".m3u"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputFileName, err)
	}
	defer outputFile.Close()

	// Find the index of "Track Name" and "Artist Name(s)" in the CSV header
	header := records[0]
	var trackIndex, artistIndex int = -1, -1
	for i, field := range header {
		switch field {
		case "title":
			trackIndex = i
		case "artist":
			artistIndex = i
		}
	}

	// Check if we found the required indices
	if trackIndex == -1 {
		return fmt.Errorf("Track Name field not found in the CSV header of %s", filePath)
	}
	if artistIndex == -1 {
		return fmt.Errorf("Artist Name(s) field not found in the CSV header of %s", filePath)
	}

	// Debug output to verify indices
	fmt.Printf("Processing %s: Track Name index: %d, Artist Name(s) index: %d\n", filePath, trackIndex, artistIndex)

	// Write the formatted data to the output file
	for _, record := range records[1:] {
		trackName := record[trackIndex]
		artistName := record[artistIndex]
		formatted := fmt.Sprintf("%s - %s.m4a", trackName, artistName)
		_, err := outputFile.WriteString(formatted + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to output file %s: %w", outputFileName, err)
		}
	}

	fmt.Printf("Formatted data has been written to %s\n", outputFileName)
	return nil
}

func main() {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current directory: %s\n", err)
		return
	}

	// List all files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Failed to read directory: %s\n", err)
		return
	}

	// Process each CSV file in the directory
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".csv" {
			err := processCSV(file.Name())
			if err != nil {
				fmt.Printf("Error processing file %s: %s\n", file.Name(), err)
			}
		}
	}
}
