package gklib

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ValidationFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("There is no file.")
		os.Exit(1)
	}

	_, err := os.Open(path)

	if err != nil {
		fmt.Println(path, "Can't open the file.")
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@")
	fmt.Println("@@ Log file : " + path + "now loading...")
	fmt.Println("@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("")
	fmt.Println("")
}

func ReadLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}
