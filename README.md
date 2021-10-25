# try

Pipe data in a command-line chain through a program meant to read stdin and write to stdout. If it works, print the output. If it fails, print the original content.

Example:

```bash
$ export code='print "hello"'

$ echo $code | python2
hello

$ echo $code | python3
  File "<stdin>", line 1
    print "hello"
                ^
SyntaxError: Missing parentheses in call to 'print'

$ echo $code | try python2
hello
$ echo $code | try python3
print "hello"
```

---

A built-in timeout will abort execution of the tried program if it takes over two seconds. This program is meant to be used for text filters (sort, grep, goimports, shawnmilo/shortcuts, etc.), so it should never take more than a fraction of a second. If one of those programs freezes (something I'm experiencing with goimports a lot recently with certain input), then this won't freeze vim, or whatever you're using.
