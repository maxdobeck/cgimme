# cgimme
This is an example cli rewritten with spf13/cobra.  It was originally a cli using only the golang std library at https://github.com/maxdobeck/gimme.

### Usage
Just like the original [gimme](https://github.com/maxdobeck/gimme) this cli just gets data you want and shunts it into your clipboard and outputs it.
```
$ cgimme -h
cgimme searches for data you need without leaving the command line.  Most often any positive results will be placed into your clipboard.

Usage:
  cgimme [command]

Available Commands:
  emails      search for emails
  help        Help about any command

Flags:
  -h, --help   help for cgimme

Use "cgimme [command] --help" for more information about a command.
```
