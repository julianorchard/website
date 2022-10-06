<!--
page_title: My First Few Months With Emacs
page_description: Documenting some random thoughts and experiences with Emacs after my first few months of solid use.
page_status: published
page_date: 2022/09/02
-->

# My First Few Months With Emacs

## From Vim Cultist to Emacs Believer

While this experience has been more eloquently documented
in other places, I thought I'd collect some
thoughts here regardless. Partly because I want to
write about it, and *partly in the hope that one
person sees it who wants to try using Emacs*!

### Background

I have been a massive fan of Vim, and the
terminal,
for a few years now. I have a nicely configured
`.vimrc` file, I've written custom functions,
and even a [little
plugin](https://github.com/julianorchard/desc.vim) to try and get to know my favourite editor a little more.
But, as the description of the [Doom
Emacs](https://github.com/doomemacs/doomemacs)
framework most aptly puts it;

> It is a story as old as time. A stubborn, shell-dwelling, and melodramatic vimmer—envious of the features of modern text editors—spirals into despair before he succumbs to the dark side.

The '*envious of the features of modern text
editors*' part is what *really* inspired me to try Emacs.

Note: *I never really liked the idea of trying
Doom. I really want to try and build my own config
and understand it a bit more. However, I'll
probably come back to Doom when I feel like I've
thoroughly made my own config, for other reasons.*

### First Contact

I started by following along with the brilliant
[Emacs From Scratch](https://youtu.be/74zOY-vgkyw)
YouTube series by [David
Wilson](https://systemcrafters.net/). I was
already feeling optimistic at this config file:

```elisp
  (setq inhibit-startup-message t)
  (scroll-bar-mode -1)
  (tool-bar-mode -1)
  (tooltip-mode -1)
  (set-fringe-mode 10)
  (menu-bar-mode -1)
```

And I was almost fully invested in giving Emacs a serious go when I realised how few lines of Elisp is enough to get started with Evil Mode!

I had thought of Emacs as being a really far-away
thing for some reason; something reserved for only those who had
understanding or dedication to make it work for
them. I was put off by the key-bindings, the
UI, leaving the terminal, and more. But the more I started to dig into it, the more I realised it *is* what most people who use the terminal are looking for. This is when I moved from 'making Emacs enough like vim to work for me' to 'wait, we don't have that in vim...'.

### Org Mode is great

Prior to playing with Emacs, I'd heard of two
'killer packages': the amazing [Magit](https://github.com/magit/magit) (git porcelain) and [Org Mode](https://orgmode.org/). While I think both are great, org-mode, for me, is one of the best things about using Emacs.

`org-babel-tangle` allows you to write org-files
(markdown-like documents) and tangle all the
blocks of code to various files. Within a few weeks, I'd
rewritten almost all of my config files,
literately, and pushed them all to the [same repo](https://github.com/julianorchard/config/).

Then there's the power of `org-agenda`. It's
great! See a list of all your upcoming tasks, let
Emacs look in multiple folders for org-files to
add to your todo-list. It's hard to explain why
it's so good if you've not used it, but it's such
a good way of noting things down.

Agenda *really shines* with `org-capture`: enter a
custom series of keys to capture a piece of
org-text that can be templated in very complicated
and useful ways. For example, I've been
maintaining a personal `journal.org` file. I hit
`SPC c` to bring up the *capture menu*, then `j`
to select a journal capture. I the write a piece
of org-text and upon `Ctrl-C Ctrl-C` (a very
common binding in Emacs to confirm or run things)
it gets inserted into the `journal.org` file under
the year, month, date; like this:

```org

* Year
** Month
*** Date
**** Timestamp
     This is an entry!

```

I initially synced my collection of org documents
with OneDrive, but I've stopped that. The beauty
of them being small text files is that I can just
sync them with scripts and other simple solutions.
I use [Syncthing](https://syncthing.net/) now, mainly, running on a server so I can connect it to my phone too (using [Orgzly](https://github.com/orgzly/orgzly-android) to get notified of events and create new ones).

### Thoughts on Vim

At the moment, I'm at a strange place with Vim.

There's a famous saying about Emacs:

> Emacs is a great operating system, lacking only a decent editor.

This quote has actually made me realise that
whilst I love the extensibility of Vim, the core
component I really love is the modal editing. It's
actually quite a relief to know that this modal
editing can be so successful in editing software
other than the original, such as in Emacs with
Evil Mode.

Emacs really is a great combination for Vim users
because the extensibility is there too. But it
makes me less concerned about trying more VSCode and JetBrains
Vim plugins in the future.

### Another Similarity of Emacs/Term Stuff

In my first few weeks of hype around Emacs, I
decided that I would be conservative about what I
would and would not *ask* of Emacs. For example,
managing mail (although I had previously used
Neomutt) is something, I told myself, that was too much
for me to bother to do with Emacs. However, as
I get deeper and deeper into Emacs, I'm finding it
harder to stick to said boundaries I've set.

I was initially surprised by this, but actually
realised that it's similar with vim-like things
for the terminal. When I first
discovered the command line, I think I set myself
the 'line' of not using it for file management.
But tools like
[Ranger](https://github.com/ranger/ranger) or
[fff](https://github.com/dylanaraps/fff) have
given me pause to reconsider that line, and I've
used them quite extensively. I think it's just a
side effect of them being so extensible.

### Conclusion

This post is a bit all over the place but,
contrary to this, I'm feeling quite clear about
where I want to go next with my Emacs journey.

I'm working on using Eshell and Dired a bit more,
trying to get more language understanding, and
trying to dive even further into Org. I will
probably post a follow-up to this at some point,
maybe at a year after using Emacs.

