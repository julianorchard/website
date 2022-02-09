<!--
page_title: A Short Journey In File Annotation
page_ : I've gone through a tiny journey in how I annotate files I make, and it makes for an interesting look into the "best language/tool for the job" conversation
page_status: published
page_date: 2022/02/09
page_image: https://julianorchard.co.uk/res/default.png
-->

# Settling Upon A Method For File Annotation

## A Short Journey...

I've been on a very short, but pretty interesting journey into file annotation.
What I mean when I say 'file annotation' is this (it's easier to show than tell,
in this case...)  : -

```
<!--
Name:        a-short-journey-in-file-annotation.markdown
By:          julianorchard
Tag Added:   09/02/2022
Desciption:  What the file does, explained in a few words here.
-->
```

Inserting this into this document used my current, preferred method. I've been
through a few. Method 0.1 would be me just manually typing. This meant the
content of the description was almost always inconsistent, aesthetically and
practically speaking. But it's still something I do if I'm rushed for some
reason...

The next method was creating a short Bash script to insert it. It was something
like this...

```
#!/bin/bash

# Title:         desc 
# By:            julianorchard
# Tag Added:     02/02/2020
# Description:   Adds file descriptions to scripts I write.

# File Exists
filename=${1}
if [[ -f ${filename} ]] ; then

# Make a new file and add filename, etc
  filenamecp="${filename}-temp"
  date=`date +%d-%m-%Y`
  echo "#!/bin/bash" > ${filenamecp}
  echo "" >> ${filenamecp}
  echo "# Title:      ${filename}" >> ${filenamecp}
  echo "# By:         julianorchard" >> ${filenamecp}
  echo "# Tag Added:  ${date}" >> ${filenamecp}

# Get file description
  printf "Please create a short description for the file: "
  read desc

# Add title and description
  echo "# Description:  ${desc}" >> ${filenamecp}

# Copy the rest of the file 
  sed -n '2,$ p' ${filename} >> ${filenamecp}

# Replace file
  mv ${filenamecp} ${filename}
  chmod +x ${filename}

else

# Error
  echo "Error: This file doesn't seem to exist!"
fi
```

The problem with this was moving to Windows... in hindsight, and now that I use
Git Bash more than CMD again (I've come full circle), I could have probably kept
using this to good effect. 

The other thing it's missing is the ability to add these descriptors to
non-bash/python files. Files like this very markdown file which I'm writing now
would benefit from `<!--` comments, and `#` comments would be confusing and
useless and bad!

A far better method was staring my in the face,
and I hadn't even noticed...

## Vim; It's Always Amazing

Vim and its scripting language were the ultimate solution for me, as I use Vim for
basically everything other than very large projects (where I probably use Atom
or VSCode instead). 

The below example just shows how this script works for HTML style comments: 

```
let filename =  expand('%:t')
let author   = "julianorchard"
let curtime  =  strftime("%d/%m/%Y")

" <!-- These Style Comments -->
	if (&ft=='html' ||
		\&ft=='php'  || 
		\&ft=='markdown')
		iab <expr> desc "<!--<cr>Name:       " . filename 
		\ . "<cr>By:         " . author
		\ . "<cr>Tag Added:  " . curtime
		\ . "<cr>Desciption: DESCRIPTION<cr>-->"
		\ . "<esc>/DESCRIPTION<cr>cw"

	"...etc. for other file types
```

You simply type the word 'desc' to run it, and it writes the Name, Author, Date
the tag was added, and then `Description: DESCRIPTION`, followed by the line break (`<cr>`)
then `-->`. Then it goes `<esc>` to get out of Vim's insert mode,
and runs `/DESCRIPTION<cr>cw`, which searches `/` for the string DESCRIPTION,
hits enter, and `cw` changes the word. Works nicely!

It's stored in a file sourced by my .vim/vimrc, and it's great. I might
eventually make it a really robust, standalone Vim package so that I can just
install it with [Vim Plug](https://github.com/junegunn/vim-plug) etc. (Vim Plug
being my favourite Vim plugin manager...).

That's it. If you don't already use one of these, I'd love to hear if this is
helpful, and if you do already use one, I'd love to know what solution you use!
