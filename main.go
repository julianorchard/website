package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

func readFileToString(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getPreamble(fileContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(fileContent))
	if err != nil {
		return "", err
	}
	for n := range doc.Descendants() {
		if n.Type == html.CommentNode {
			return n.Data, nil
		}
	}

	return "", fmt.Errorf("found a file without preamble")
}

func isDirectory(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), err
}

type PageMetadata struct {
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`
}

func pageMetadata(rawMetadata string) (PageMetadata, error) {
	var output PageMetadata
	err := yaml.Unmarshal([]byte(rawMetadata), &output)
	if err != nil {
		return PageMetadata{}, fmt.Errorf("cannot unmarshal data: %v", err)
	}
	return output, nil
}

func main() {
	fileList := []string{}
	err := filepath.Walk(
		"./templates/",
		func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		},
	)
	if err != nil {
		fmt.Println("ffs")
		os.Exit(1)
	}

	for _, path := range fileList {
		if d, err := isDirectory(path); err != nil || d {
			continue
		}
		fmt.Println("reading file", path)
		file, err := readFileToString(path)
		if err != nil {
			fmt.Printf("couldn't read file: %s", file)
		}

		fmt.Println("getting preamble", path)
		preamble, err := getPreamble(file)
		if err != nil {
			fmt.Println("unable to get preamble:", err)
		}

		meta, err := pageMetadata(strings.TrimSpace(preamble))
		if err != nil {
			fmt.Println("unable to parse preamble:", err)
		}

		fmt.Println("page name is", meta.Name)
		fmt.Println("page name is", meta.Icon)
	}
}
