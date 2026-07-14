# My Website

This used to be a bash script! Check out [this
hash](https://github.com/julianorchard/website/tree/f006a2c7a238ca0ce350b8254819a7a49dd5d745)
to see what that was all about.

## Hugo at home

This is just a `main.go` file which operates very similarly to my previous bash
static site generator implementation in many ways. Rather than a full
site-rewrite, I thought I'd just replace this slowly!

I started a full rewrite using the 11ty static site generator. And I do like
the styles I came up with for that website. However, I didn't love 11ty, and
all the `npm` supply chain attacks of late have made me reconsider going for
Node at all. Plus this was fun!

## Usage

This repo mainly runs a GitHub Action to update a GitHub page based on
artifacts created by the pipeline run. However, you can run things locally too!

### Docker

You can run this locally with Docker:

```sh
docker build -t julian-website .
docker run --rm -p 8080:80 julian-website:latest
open http://localhost:8080
```

This is useful for testing the site locally.

### Make

Use the makefile to clean and build the HTML files locally:

```sh
make all
```

Useful for testing files are outputting correctly!

## License

This is free software under the MIT license, see the [LICENSE](LICENSE) file
for more information.
