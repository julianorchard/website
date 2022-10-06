<!--
page_title: NA Tiny Vim Tip
page_description: Literally the smallest, but quite useful, tip for vim search and replace.
page_status: draft
page_date: 2022/07/26
page_image:
-->

# A Very Quick Tip For Vim

## Using the find an replace feature...

Whilst rewriting my configuration files literately (
a post surely to come on this, later!),
and I'm using org-babel-tangle with Emacs.

However, when I started making the configuration,
I didn't know about the `:tangle c:/path :mkdirp
yes` attribute, so I had to add this in after I'd
nearly *finished* the document! (That's a lot of
code-blocks to edit...)

Obviously, because I'm using Evil-mode, I assumed
I could easily do a find and replace on this...
something like:

```vimscript
:%s/:tangle.*/...          ...            ...
```

But I realised I needed to keep the file path,
which was a variable part of this comment; the
files tangled to different locations, so it was
important they still kept those paths despite
matching any path `.*`... So what was the
solution?

A tiny vim tip, that I didn't know about; the `&`
ampersand operator can be used to replace with the
entire string, effectively allowing you to append
the string. My example looked like this:

```vimscript
:%s/:tangle.*/& :mkdirp yes/g
```

That's the tip! I will be using this a lot from
now on, and I don't know how I didn't know about
this one before...
