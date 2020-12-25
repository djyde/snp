package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
)

// VERSION code
const VERSION = "v0.2.0"

func main() {
	// cwd, _ := os.Getwd()

	// output path, default is current file
	// outputPath := flag.String("o", cwd, "Where to output the snippet")

	shouldPrint := flag.Bool("p", false, "Should print out the json text")

	autoUpdate := flag.Bool("u", false, "Should automatically update snippet file in VS Code")

	printVersion := flag.Bool("v", false, "Print version")

	flag.Parse()

	if *printVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	json := ParseSnpFiles()

	if *shouldPrint {
		fmt.Println(string(json))
	}

	if *autoUpdate {
		configPath := configdir.LocalConfig("Code")

		target := filepath.Join(configPath, "User", "snippets", "snp.code-snippets")

		fmt.Println("Updated to", target)

		ioutil.WriteFile(target, json, 0644)
	}

}
