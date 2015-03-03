# Gosh!

[![Build Status](https://travis-ci.org/bfontaine/gosh.svg?branch=master)](https://travis-ci.org/bfontaine/gosh)
[![Coverage Status](https://coveralls.io/repos/bfontaine/gosh/badge.svg?branch=master)](https://coveralls.io/r/bfontaine/gosh?branch=master)

**Gosh** is a simple shell written in Go.

## Install

    go install github.com/bfontaine/gosh

### Dependencies

* Go 1.2 or higher
* Readline

## Usage

    gosh

You’ll get a prompt, from which you can execute commands like in every other
shell. Use `^D` or `exit` to exit it.

## Features

Just a basic REPL, right now.

Gosh doesn’t support quotes, escaping and wildcards, nor any loop or
conditional constructions.

### Builtin commands

* `alias L=V`: add an alias `L` to `V`. `V` can be any command, even with
  spaces or weird characters. Aliases can’t be recursive.
* `cd <path>`: change the current directory (can contain spaces)
* `echo ...`: print stuff. You can insert environment variables with `$var` or
  `${var}`
* `quit`, `exit`, `^D`: exit the shell

### `~/.goshrc`

If a `~/.goshrc` file exists, Gosh reads it and executes it line-by-line as if
it were given on the prompt.

You can use it for common aliases, e.g.:

```sh
alias ll=ls -l
alias la=ls -la
```

### Options

Gosh currently supports the following options:

* `-debug`: show all errors
* `-trace`: show all lines are they are executed, both from the `~/.goshrc` and
  the interactive session

Use `gosh -h` for more info.

### Editor support

* Vim: [vim-gosh](https://github.com/bfontaine/vim-gosh)
