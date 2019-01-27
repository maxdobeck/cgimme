# cgimme
This is an example cli rewritten with spf13/cobra.  It was originally a cli using only the golang std library at https://github.com/maxdobeck/gimme.

### Usage
Just like the original [gimme](https://github.com/maxdobeck/gimme) this cli just gets data you want and shunts it into your clipboard and outputs it.
```
$ cgimme -h

Usage:
  cgimme [command]

Available Commands:
  emails      search for emails
  help        Help about any command

Flags:
  -h, --help   help for cgimme

Use "cgimme [command] --help" for more information about a command.
```

### Why?
Cobra is super useful!  It has an additional generator on top of the usage of [spf13/pflag](https://github.com/spf13/pflag) for [UNIX style flags](http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html).  But with [maxdobeck/gimme](https://github.com/maxdobeck/gimme) I wanted a nice and simple, low dependency, cli.  There's always a tradeoff between dependencies and development time.

This is sort of an expirement and so far I'm already seeing potential issues, almost literally.  Every open Issue with spf13/cobra is a potential headache for me down the road.  If no one maintains the repo then we're in trouble.  

With the addition of modules we may start seeing some issues like that.  Currently the Cobra generator defualts to GOPATH/src for the parent directory.  But what if I don't want to work out of that dir?

So far this is great though!  Worst case I can just borrow someo of Cobra's patterns and refactor maxdobeck/gimme.
