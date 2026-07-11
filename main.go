package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/yuin/goldmark"
	gmhtml "github.com/yuin/goldmark/renderer/html"
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
	Name        string    `yaml:"page_title"`
	Description string    `yaml:"page_description"`
	Image       string    `yaml:"page_image"`
	Template    string    `yaml:"template"`
	PageDate    string    `yaml:"page_date"`
	Draft       bool      `yaml:"draft"`
	PageHead    string    `yaml:"page_head"`
	Rel         string    `yaml:"rel"`
	Date        time.Time // this is parsed from PageDate
	Content     string    // this is the page body
	Styles      string    // CSS from another file
	JavaScript  string    // JS from another file
}

func pageMetadata(rawMetadata string, content []byte) (PageMetadata, error) {
	// Unsafe 😎🔥
	md := goldmark.New(
		goldmark.WithRendererOptions(
			gmhtml.WithUnsafe(),
		),
	)
	var body bytes.Buffer
	if err := md.Convert(content, &body); err != nil {
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

	if output.PageDate == "" {
		output.PageDate = "1996-10-12"
	}
	output.Date, err = time.Parse("2006-01-02", output.PageDate)
	if err != nil {
		return PageMetadata{}, fmt.Errorf(
			"failed to parse time on page: %w",
			err,
		)
	}

	if output.Template == "" {
		output.Template = "main"
	}

	if output.PageHead == "" {
		output.PageHead = "regular"
	}

	if output.Rel == "" {
		return PageMetadata{}, fmt.Errorf("rel not provided and is required")
	}

	cssPath := "./src/style.css"
	css, err := readFileToString(cssPath)
	if err != nil {
		return PageMetadata{}, fmt.Errorf(
			"failed to include CSS at %s: %w",
			cssPath, err,
		)
	}
	output.Styles = string(css)

	jsPath := "./src/script.js"
	js, err := readFileToString(jsPath)
	if err != nil {
		return PageMetadata{}, fmt.Errorf(
			"failed to include CSS at %s: %w",
			jsPath, err,
		)
	}
	output.JavaScript = string(js)

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
		"./content/",
		func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		},
	)
	if err != nil {
		fmt.Println("couldn't do it. just couldn't")
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

		if meta.Draft {
			continue
		}

		_, err = renderPage(meta)
		if err != nil {
			fmt.Printf("failed to render %s: %v", path, err)
		}

		rel := outPath(meta, path)
		fmt.Println("OUTPUTTING TO:", rel)
		// writeFile(rel, render)
	}
}

func writeFile(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write file %q: %w", path, err)
	}

	return nil
}

// FIXME: this stuff
func outPath(meta PageMetadata, path string) string {
	// Date
	urlPageDate := strings.ReplaceAll(meta.PageDate, "-", "/")
	fmt.Println(meta.Rel, ": replace {date} with", urlPageDate)
	r := strings.ReplaceAll(meta.Rel, "{date}", urlPageDate)

	// Page Name
	baseName := strings.TrimRight(path, "/")
	baseName = strings.Split(baseName, "/")[len(strings.Split(baseName, "/"))-1]
	baseName = strings.TrimSuffix(path, filepath.Ext(path))
	r = strings.ReplaceAll(meta.Rel, "{name}", baseName)

	// Final path
	r = filepath.Join("./output", fmt.Sprintf("%s.html", r))

	return r
}
