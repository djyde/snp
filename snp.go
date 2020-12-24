package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Attr struct {
	Scope       string
	Description string
}

type SnippetItem struct {
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Body        []string `json:"body"`
	Prefix      string   `json:"prefix"`
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func main() {

	cwd, _ := os.Getwd()

	files, err := ioutil.ReadDir(cwd)

	if err != nil {
		log.Fatal(err)
	}

	snippet := make(map[string]SnippetItem)

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".snp") {

			attr := Attr{}

			filePath := filepath.Join(cwd, file.Name())
			rawContent, readFileErr := ioutil.ReadFile(filePath)
			if readFileErr != nil {
				log.Fatal(err)
			}

			prefix := strings.TrimSuffix(file.Name(), filepath.Ext(filePath))

			content, parseFmErr := Unmarshal(rawContent, &attr)
			if parseFmErr != nil {
				log.Fatal(parseFmErr)
			}

			stringContent := strings.TrimRight(strings.TrimLeft(string(content), "\n"), "\n")

			body := strings.Split(stringContent, "\n")

			// fmt.Println(body)

			item := SnippetItem{attr.Scope, attr.Description, body, prefix}

			snippet[prefix] = item

			// fmt.Println(item)
			// fmt.Println(filepath.Join(cwd, file.Name()))
		}
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	json, jsonEncodingErr := JSONMarshal(snippet)

	if jsonEncodingErr != nil {
		log.Fatal(jsonEncodingErr)
	}

	ioutil.WriteFile("snp.code-snippets", json, 0644)
}
