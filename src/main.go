package main

import (
	"fmt"
	"os"
	"io"
)

func main() {

	path := "WhaT-THE-fuClk";

	createFile(path)
	writeFile(path, "THIS IS FCUKKK")
	readFile(path)
	deleteFile(path)
}

func createFile(path string) {
	
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		
		if isError(err) { 
			return 
		}

		defer file.Close()
	}

	fmt.Println("==> Done create file ::", path)
}

func deleteFile(path string) {
	
	err := os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting file", path);
}

func readFile(path string) {
	
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

	text := make([]byte, 1024)
	for {
		_, err := file.Read(text)

		if err == io.EOF { break; }

		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("===> done reading from file")
	fmt.Println(string(text))
}

func writeFile(path string, content string) {
	
	file, err := os.OpenFile(path, os.O_RDWR, 0644)

	if isError(err) {
		return
	}

	defer file.Close()

	_, err = file.WriteString(content);
	if isError(err) { return }

	err = file.Sync()
	if isError(err) { return }

	fmt.Println("==> done writing file");
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}
