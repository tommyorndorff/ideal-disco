# ideal-disco
What doesnt an ideal disco have.

## Why is this an odd format
I dont like editing PowerPoint decks and they dont repo well.  In the spirit of re-using the content, this is all written in Markdown and renders to slides with [Marp](https://marp.app/).  

The 'quick start' is at Marp's [Quick Start](https://marp.app/#get-started).  There's a VSCode plugin that will allow Markdown preview per standard VSCode preview functionality.  

Take a look at one of the Makefiles to see how to build PDFs:
```sh
$ cd week-1
$ make
npx @marp-team/marp-cli@latest 01-aws_history.md --pdf --allow-local-files -o pdf/01-aws_history.pdf
...
```
