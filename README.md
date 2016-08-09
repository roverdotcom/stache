## Stache [![Build Status](https://travis-ci.org/roverdotcom/stache.svg?branch=master)](https://travis-ci.org/roverdotcom/stache)

A CLI tool for rendering mustache templates using ENV as the context.

### Pipe input

```
$ echo "My home directory is: {{ HOME }}" | stache
```

### Read from a file

```
$ echo "My username is {{ USER }}" > test.mustache
$ stache -f test.mustache > test.txt
```

### Write to a file

```
$ stache -f test.mustace -o test.txt
```
