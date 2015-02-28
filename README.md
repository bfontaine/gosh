# Gosh!

[![Build Status](https://travis-ci.org/bfontaine/gosh.svg?branch=master)](https://travis-ci.org/bfontaine/gosh)
[![Coverage Status](https://coveralls.io/repos/bfontaine/gosh/badge.svg?branch=master)](https://coveralls.io/r/bfontaine/gosh?branch=master)

**Gosh** is a simple shell written in Go.

## Install

    go get github.com/bfontaine/gosh

### Dependencies

* Go 1.2 or higher
* Readline

## Usage

    gosh

You’ll get a prompt, from which you can execute commands like in every other
shell. Use `^D` or `exit` to exit it.

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
