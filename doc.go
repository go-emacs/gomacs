/*
The gomacs provides golang programming environment for emacs.

Install

use `go get` command:

   $ go get github.com/go-emacs/gomacs

Usage

Launch:

   $ gomacs                  # launch emacs

Update:

   $ gomacs --update         # update emacs lisp from internet.

The gomacs can use emacs option and operation. for example:

   $ gomacs --help           # show emacs --help
   $ gomacs main.go          # open main.go
   $ gomacs +12 main.go      # go to line
   $ gomacs -rv              # switch foreground and background color
       :

Setting

You can use the following `settings.json`:

   {
       "emacs" : "/usr/local/bin/emacs",
       "args" : ["-q"]
   }

The "emacs" defines your emacs execution path.
You can use any emacs options, use the "args" field. See emacs --help.

The gomacs search the setting file from:

    1 ./.gomacs.json
    2 ~/.config/gomacs/settings.json
    3 $GOPATH/src/github.com/go-emacs/gomacs/settings.json


enjoy!

*/
package main
