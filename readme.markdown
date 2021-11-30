# My Website

I wanted to host a blog to occasionally write about topics that interest me, and
about experiences I've had. However, I don't have much need for static site
generators outside of this use case, so it seemed redundant to learn [Hugo](https://gohugo.io/) or
[Jekyll](https://jekyllrb.com/), and I like to do things myself anyway. So I wrote a short Bash script
(`blog`) as my own little static site generator. It has everything I'd want from
it. 

## Usage

From `./blog -h`: 

``` 
Usage of this tiny static site generator: 
-n / --new         create a new post 
-b / --build     build the full site 
```

To elaborate on this, I took inspiration from [Hugo](https://gohugo.io/), and used comments at the
start of the file to give basic information about the pages to build. For
example: 

``` 
<!-- 
page_title: Title of a Blog Post 
page_description: This is the description of a blog post.  
page_status: published 
page_date: 2021/11/20 
-->
```

The minimum a *post* needs to be published is a `page_status: published`, and a
`page_date`. A static page is definied by not having a `page_date`. For example,
you can create an "about" page by haing a post called `about.markdown` and; 

``` 
<!-- 
page_title: About Page
page_description: This is the description of the about page.
page_status: published 
-->
```

It creates a folder `about` with an `index.html` file in there with this
content.


## Installation

This setup doesn't really have anything linking it to a particular domain. I
really just want this to be very simple to use. It should be `git clone`'d into
your website directory, and from there, run the script with `./blog` (`chmod +x`
if it's not executing). 

If you have any questions please email me,
	[hello@julianorchard.co.uk](mailto:hello@julianorchard.co.uk).
