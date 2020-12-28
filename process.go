package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type attr struct {
	Scope       string
	Description string
}

type snippetItem struct {
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Body        []string `json:"body"`
	Prefix      string   `json:"prefix"`
}

func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// ParseSnpFiles read current work directory's .snp file and parse it to code snippet
func ParseSnpFiles() []byte {
	cwd, _ := os.Getwd()

	snippet := make(map[string]snippetItem)
	err := filepath.Walk(cwd, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(filePath, ".snp") {

			attribute := attr{}

			rawContent, readFileErr := ioutil.ReadFile(filePath)
			if readFileErr != nil {
				log.Fatal(err)
			}

			filename := path.Base(filePath)
			prefix := filename[0 : len(filename)-4]

			content, parseFmErr := Unmarshal(rawContent, &attribute)
			if parseFmErr != nil {
				log.Fatal(parseFmErr)
			}

			stringContent := strings.TrimRight(strings.TrimLeft(string(content), "\n"), "\n")

			body := strings.Split(stringContent, "\n")

			item := snippetItem{attribute.Scope, attribute.Description, body, prefix}

			snippet[prefix] = item
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	json, jsonEncodingErr := jsonMarshal(snippet)

	if jsonEncodingErr != nil {
		log.Fatal(jsonEncodingErr)
	}

	return json
}
