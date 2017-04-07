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
