<!--
page_title: Meta Images are now built into this blog
page_description: This is a quick post to showcase how I've implemented meta images into this blog, rather nicely.
page_status: published
page_date: 2021/12/04
page_image: https://images.unsplash.com/photo-1520004434532-668416a08753?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80
-->

# Meta Images; now a feature of this blog!

Meta images are actually incredibly important to how well a website will present itself to the world. Obviously having other good social sharing meta tags like `og:title` and `og:description` are important, but meta images are your chance to really make a link look good!

They're there when we show a page in an instant messenger, ususally, and they're also usually present on social media. It's a free pass to get some attention, and now, this blog works nicely with them. 

## How it works

When you run `./blog -n`, the final question you'll be asked is now "Do you wish to use the default image for meta tags?". A default `yes` response will mean the page just includes an image stored in `/res/default.png`, and a `no` response will prompt you to enter either a URL or a local location (ideally, store your images in `/res/`).

![Example of what the default sharing looks like.](/res/blog-meta-example.png)

That's it. Now links to this blog look nicer. 
