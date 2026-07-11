package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

func readFileToString(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
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
	Name     string `yaml:"name"`
	Icon     string `yaml:"icon"`
	Template string `yaml:"template"`
	Content  string // this is the page body
}

func pageMetadata(rawMetadata string, content []byte) (PageMetadata, error) {
	var body bytes.Buffer
	if err := goldmark.Convert(content, &body); err != nil {
		return PageMetadata{}, fmt.Errorf(
			"failed to convert document to HTML: %w",
			err,
		)
	}
	output := PageMetadata{
		Content: body.String(),
	}

	err := yaml.Unmarshal([]byte(rawMetadata), &output)
	if err != nil {
		return PageMetadata{}, fmt.Errorf("cannot unmarshal data: %w", err)
	}
	return output, nil
}

func renderPage(meta PageMetadata) ([]byte, error) {
	templatePath := filepath.Join(
		"templates",
		meta.Template+".html.tmpl",
	)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("parse template %s: %w", templatePath, err)
	}

	var output bytes.Buffer

	if err := tmpl.Execute(&output, meta); err != nil {
		return nil, fmt.Errorf("execute template %s: %w", templatePath, err)
	}

	return output.Bytes(), nil
}

func main() {
	fileList := []string{}
	err := filepath.Walk(
		"./articles/",
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
		preamble, err := getPreamble(string(file))
		if err != nil {
			fmt.Println("unable to get preamble:", err)
		}

		meta, err := pageMetadata(strings.TrimSpace(preamble), file)
		if err != nil {
			fmt.Println("unable to parse preamble:", err)
		}

		fmt.Println("page name is", meta.Name)
		fmt.Println("page icon is", meta.Icon)
		fmt.Println("page template is", meta.Template)
		fmt.Println("page content is", meta.Content)

		render, err := renderPage(meta)
		if err != nil {
			fmt.Printf("failed to render %s: %v", path, err)
		}

		fmt.Println(string(render))
	}
}
