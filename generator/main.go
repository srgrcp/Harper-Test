package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	generate("schema")
}

func generate(path string) {
	generatedFilePath := path + "/" + path + ".go"
	outputFile, err := os.Create(generatedFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputFile.Write([]byte("package " + path + "\n\n"))
	outputFile.Write([]byte("// Data auto-generated schema\n"))
	outputFile.Write([]byte("var Data string = `\n"))
	err = filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if subPath == path || info.IsDir() || !strings.HasSuffix(subPath, ".graphql") {
			return nil
		}
		return handleFile(subPath, outputFile)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	outputFile.Write([]byte("`\n"))
}

func handleFile(path string, outputFile *os.File) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// Manejo de tíldes invertidas literales dentro del string.
	escapedContent := strings.Replace(string(fileContent), "`", "`+\"`\"+`", -1)
	outputFile.Write([]byte(escapedContent + "\n"))
	fmt.Println("Generated: " + path)
	return nil
}
