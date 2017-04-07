# try

Pipe data in a command-line chain through a program meant to read stdin and write to stdout. If it works, print the output. If it fails, print the original content.

Example:

```bash
# valid code
$ echo 'print "hello"' | try python
hello

# invalid code (missing close-quote)
$ echo 'print "hello' | try python
print "hello

