<!--
page_title: Testing Code Blocks and Images
page_description: This is a quick test of how I've implemented <code> and <img stylings!
page_status: published
page_date: 2021/12/05
page_image: https://julianorchard.co.uk/res/default.png
-->

# This Post Is An Example...

HTML `code` tags and `img` tags are something I want to look nice on this blog, and `code` tags aren't something I've really had to make a huge effort with before. I've not written a website where I have a huge need to style them, but since I'll probably be writing about code I've written on this blog, I *will* need them here: 

## How I've settled on code tags working

```
@echo off

echo Hello World!

echo This is a really, really long line of code and it's just going to keep going and going and going because I need it to be a long line of code and to keep going.

```

The only thing is code highlighting. I've looked into [highlightjs](https://highlightjs.org/usage/), but I think it's overkill, and if there's any very large bits of code I want to include, I'll probably just use [GitHub gists](https://gist.gihub.com) to include them here. 

## How I've settled on images working

Images will look like...

![This is what image captions will be like!](https://images.unsplash.com/photo-1638459554090-675e14d87d1b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80)

When [Pandoc](https://pandoc.org/) (which the [script](https://github.com/julianorchard/julianorchard.github.io) used to run this site calls) converts Markdown to HTML, it adds any image description to `<figcaption>`, as well as the image `alt=` text. I therefore wanted to come up with a nice way to display this caption, and based on the colours I was already using for other parts of the site, this is what I came up with. It used a `:before` CSS element to add a border in the top left, and the bottom right caption is the `<figcaption>` element. 

Anyway, that's it.
