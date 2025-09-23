package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// Log the error
		panic(err)
	}
	
	fileInfo, err := f.Stat()
	if err != nil {
		// Log the error
		panic(err)
	}
	
	fmt.Println("File Name", fileInfo.Name())
	// fmt.Println("File or Folder", fileInfo.IsDir())
	// fmt.Println("File Size", fileInfo.Size())
	// fmt.Println("File Permission", fileInfo.Mode())
	// fmt.Println("File Modified At", fileInfo.ModTime())
	// fmt.Println("File Modified At", fileInfo.Sys())


	// Read File
	// defer f.Close()

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))


	// Read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fInfo, err := dir.ReadDir(-1)

	// for _, fi := range fInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }


	// Create and Write
	// fl, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fl.Close()

	// // fl.WriteString("Hello Yo")
	// // fl.WriteString(" Hello Baba")
	
	// bytes := []byte(" Hello Golang")
	// fl.Write(bytes)


	// Read and Write to another File (Streaming Fashion)
	source, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer source.Close()


	destination, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destination.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(destination)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written text in new file")

}