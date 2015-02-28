# Gosh!

[![Build Status](https://travis-ci.org/bfontaine/gosh.svg?branch=master)](https://travis-ci.org/bfontaine/gosh)

**Gosh** is a simple shell written in Go.

## Install

    go get github.com/bfontaine/gosh

### Dependencies

* Go 1.2 or higher
* Readline

## Features

Just a basic REPL, right now.

Gosh doesn’t support quotes, escaping and wildcards for now.

### Builtin commands

* `alias L=V`: add an alias `L` to `V`. `V` can be any command, even with
  spaces or weird characters.
* `cd <path>`: change the current directory (can contain spaces)
* `echo ...`: print stuff. You can insert environment variables with `$var` or
  `${var}`
* `quit`, `exit`, `^D`: exit the shell
