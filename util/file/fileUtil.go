package fileUtil

import (
	"bufio"
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)


func CreateFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
	}
}

// Exists returns whether the given file or directory exists or not
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func LineCounter(path string) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)


	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func AppendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}

	// save changes
	err = f.Sync()
	if err != nil {
		log.Fatalf("Unable to Sync changes in WriteLineInFile function: %v", err)
	}

	return nil
}

func WriteLineInFile(path string, text string) {
	// open file using READ & WRITE permission
	if !Exists(path) {
		CreateFileWPath(path)
	}
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Unable to open file %s: %s", path, err)
	}
	defer file.Close()
	// write some text in line
	_, err = file.WriteString(text + "\n")
	if err != nil {
		log.Fatalf("Unable to write string in WriteLineInFile function: %v", err)
	}
	// save changes
	err = file.Sync()
	if err != nil {
		log.Fatalf("Unable to Sync changes in WriteLineInFile function: %v", err)
	}
}

func DeleteFile(path string) error {
	var err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func CreateFileWPath(p string) error {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return err
	}
	var file, err = os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}