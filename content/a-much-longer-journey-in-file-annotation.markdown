<!--
page_title: A Much Longer Journey In File Annotation
page_description: I take a deeper dive into how I now 'tag' my files with Vim.
page_status: draft
page_date: 2022/04/19
-->

# A Much Longer Journey Into File Annotation

## And VimScript in General; an update from my last post on this topic

Recently, I [wrote about](https://julianorchard.co.uk/posts/2022/02/09/a-short-journey-in-file-annotation/) how I have been on a short journey in how I 'tag' my scripts and files with content. This journey has become substantially longer than I thought, and I've even managed to write my first Vim Plugin, of sorts.

I've called the plugin '*desc.vim*', and it can be installed by vim plugin managers (such as my favourite one, [vim-plug](https://github.com/junegunn/vim-plug)).

## Challenges, Etc.

Writing my first Vim 

### The Good

I immediately realised that if I was going to make
this a plugin that people other than myself would
want to use (this remains to be seen, of
course...) I would need a more robust way of doing
the signature part of the description. People
using it need to be able to easily set their
username, email address, whatever else. There are
**four** possible ways that this can be done, and
I'm proud of how nicely this part works: 

1. If you have `git config --global user.name`
   and `.email` set, desc.vim will automatically
   pick that, and format it nicely:

```
>   File:       example.md
>   Author:     julianorchard <hello@julianorchard.co.uk>
>   Tag Added:  2022-04-19
>   Desciption: Example Description.
```

2. If you don't have a `git config --glo...` set,
   it will be blank at that point; I can't think
   of a nicer way to do it.
3. You can override those two defaults with
   either a string vairable, or a list/array
   variable set in your `.vimrc` file, for
   example:

```
let g:desc_author = "Example <example@email.com>"
```

4. And as a list/array (which allows you to remap
   the iabbrev): 

```
let g:desc_author = [["desc",    "Example <example@email.com>"],
                    \["descrip", "Another Example <another@email.com>"]]
```

### The Bad

This plugin faced some interesting challenges. I
was determined, for better or worse, to use vim's
[abbreviations](http://vimdoc.sourceforge.net/htmldoc/map.html#Abbreviations),
specifically, `iabbbrev`'s in this case. This
initially threw me, because they simply return
what you tell them to, and I needed to include
variables in the Abbreviation; therefore, I needed
to use `<expr>`. `

### 


Anyway, the repo can be found on [my
GitHub](https://github.com/julianorchard/desc.vim). That's it. 
