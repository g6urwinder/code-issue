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

/*
* it will get the stats of path if exists
* if does not exists than create file
* if error happens return
* if no error's than return success message
*/
func createFile(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil { return }

		defer file.Close()
	}

	fmt.Println("==> Done create file ::", path)
}

/*
* It will delete file on provided path
* if error happens it will return immediately
* if no error than it will return success text
*/
func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil { return }

	fmt.Println("==> done deleting file", path);
}

/*
* open file with permission read, write and 0644
* 0644 is a symbolic link of -rw-r-r--
* 
* return if error End Of File happens
* return if error is not EOF and print it
* 
* if not error than return content of file
*/
func readFile(path string) {
	
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil { return }
	defer file.Close()

	text := make([]byte, 1024)
	for {
		_, err := file.Read(text)
		if err == io.EOF { break; }

		if err != nil && err != io.EOF {
			fmt.Println("Error => ", err);
			break
		}
	}

	fmt.Println("===> done reading from file")
	fmt.Println(string(text))
}

/*
* Open file with read, write and 0644 permissions
* 
* if err than return
*
* if no errors than return success message
*/
func writeFile(path string, content string) {
	
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil { return }

	defer file.Close()

	_, err = file.WriteString(content);
	if err != nil { return }

	err = file.Sync()
	if err != nil { return }

	fmt.Println("==> done writing file");
}
