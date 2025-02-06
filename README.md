# Ncurses with unicode support

This Go package is a wrapper for the ncurses library.  It provides
access to the ncurses terminal handling library functions.
To use the Go package, you need to have the ncurses library installed.
This must include the header files and the shared library.

## Installation

This package can be installed using the ``go get`` command::

    go get github.com/blippy/go-ncurses

## Usage

In order to get UTF-8 to work, you will need to set locale LC\_ALL to "".

Documentation is available via the package's online help, either on
pkg.go.dev_ or on the command line::

    godoc github.com/blippy/go-ncurses

  .. _pkg.go.dev: https://pkg.go.dev/seehuhn.de/go/ncurses
