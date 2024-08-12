package reader

import (
	"os"
	"testing"
)

// Helper function to create a temporary file with the given content
func createTempFile(content string) (string, error) {
	file, err := os.CreateTemp("", "*.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

// Test case for valid tetrominoes
func TestReadTetrominoes_Valid(t *testing.T) {
	content := `####
####
####
####
	
#...
#...
#...
#...

...#.
...#.
.###.
.....`

	filePath, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filePath)

	_, err = ReadTetrominoes(filePath)
	if err == nil {
		t.Fatalf("Expected an error biut got nil")
	}
}

// Test case for invalid tetrominoes
func TestReadTetrominoes_Invalid(t *testing.T) {
	content := `####
####
###
####
	
#...
#...
#...
#...

...#.
...#.
.###.
.....`

	filePath, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filePath)

	_, err = ReadTetrominoes(filePath)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

// Test case for invalid character
func TestReadTetrominoes_InvalidCharacter(t *testing.T) {
	content := `####
####
#####
####
	
#...
#...
#...
#...

...#.
...#.
.###.
.....`

	filePath, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filePath)

	_, err = ReadTetrominoes(filePath)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}

// Test case for tetromino with incorrect size
func TestReadTetrominoes_IncorrectSize(t *testing.T) {
	content := `##.#
####
##..
####
	
#...
#...
#...
#...

...#.
...#.
.###.
.....`

	filePath, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filePath)

	_, err = ReadTetrominoes(filePath)
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}
