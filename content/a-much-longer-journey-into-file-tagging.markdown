<!--
page_title: A Much Longer Journey Into File Annotation
page_description: Previously, I wrote about how I added metadata to my files. This lead me down a far longer path than I realised!
page_status: publish
page_date: 2022/04/27
page_image:https://julianorchard.co.uk/res/desc-vim.png
-->

# A Much Longer Journey Into File Annotation

## How I ended up writing my first Vim Plugin...

In a [previous post](https://julianorchard.co.uk/posts/2022/02/09/a-short-journey-in-file-annotation/), I took a short look into how I'd gone from tagging my files using a small shell script, and how I'd changed my method to use a Vim `iabbrev`.

I thought that would be the end of my journey, but I decided to use the
opportunity to make a small Vim plugin which made the whole process a lot more
easy to configure if there's extra functionality required.

## The Good, The Bad, & The Ugly

The plugin I ended up creating has some really nice elements, some elements I
struggled with, and some problems I'm yet to solve...

### The Good

I needed a way to let someone using the file input their signature (e.g.,
"Julian Orchard <hello@julianorchard.co.uk>"), and I came up with a way of doing
that which works really well in my opinion.

Once the plugin is installed (I use
[Vim-Plug](https://github.com/junegunn/vim-plug) for regular Vim, and with
Neovim I use [Packer.nvim](https://github.com/wbthomason/packer.nvim)), you can
define either a string in your `.vimrc` file, or a list.

For example;

```

" A String
let g:desc_author = "Julian Orchard <hello@julianorchard.co.uk>"

" A List
let g:desc_author = [["desc", "Julian Orchard <hello@julianorchard.co.uk>"],
                    \["des2", "Another Signature <me@julianorchard.co.uk>"]]

```

This conveniently solves the issue of how to easily define a signature, but also
makes it possible for the user of the plugin to change the `iabbrev` *keyword*
used. The list doesn't have to have multiple values; if you want to reclaim the
word `desc` (the default keyword), you can just do something like this: 

```

let g:desc_author = [["trigger_desc", "Julian Orchard <hello@julianorchard.co.uk>"]

```

I was really happy with this solution, but what happens if someone doesn't add
anything to their `.vimrc` file? 

Initially, I just removed the line from the description, if no `g:desc_author`
was defined, but I came up with a pretty nice alternative... you can add this as
an `iabbrev` to your config files independently too, I think it's pretty good:

```

iab <expr> ~g substitute(system('git config --global user.name') . " <" . 
    \system('git config --global user.email') . ">", '\n', '', 'g')

```

This gets your `git config --global user.name` and `user.email` variables, and
uses them to create the signature. Mine is `julianorchard <hello@julianorchard.co.uk>`, for example.


### The Bad

I think using `iabbrev` might have been a mistake, in hindsight. I didn't find
it very flexible in how it works, and figuring out I needed `<expr>` took longer
than it should have; it wasn't very intuitive. 

The code also repeats itself more than I'd like as a result. It's not too bad,
but I have four different `iabbrev`'s in the final version; 

- One for block comments and there is a signature to add
- One for line comments and there is a signature to add
- One for block comments and there is no signature to add
- One for line comments and there is no signature to add

It's just not as clean as I'd like, and I think I'd have more flexibility using
a remapped keybinding or some other feature.

### The Ugly

I'm fairly certain the `autocmd`'s I have made use of do not work in the most
performative way. I don't notice an impact from them, but each time the user goes into *insert mode*, the `iabbrev` is updated. 

## Perfectionist

Unfortunately, I am a bit of a perfectionist, and I'm not 100% happy with the
plugin, but I have to praise myself; for my first Vim plugin, I'm pretty content
with it.

I've tested it with my Neovim config, which is new for me, and it seems to also
work well. 

That's it.
