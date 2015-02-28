# Gosh!

**Gosh** is a simple shell written in Go.

## Install

    go get github.com/bfontaine/gosh

### Dependencies

* Go 1.2 or higher
* Readline

## Features

Just a basic REPL, right now.

### Builtin commands

* `alias L=V`: add an alias `L` to `V`. `V` can be any command, even with
  spaces or weird characters.
* `cd <path>`: change the current directory (can contain spaces)
* `echo ...`: print stuff. You can insert environment variables with `$var` or
  `${var}`
* `quit`, `exit`, `^D`: exit the shell
