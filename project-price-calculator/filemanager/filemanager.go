package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) WriteResult(data any) error {

	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return fmt.Errorf("WriteJSON: failed to create file: %w", err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return fmt.Errorf("WriteJSON: failed to encode data: %w", err)
	}

	file.Close()
	return nil
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, fmt.Errorf("ReadLines: failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("ReadLines: failed to scan file: %w", err)
	}

	file.Close()
	return lines, nil
}

func New(inputPath, outputPath string) FileManager {

	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
