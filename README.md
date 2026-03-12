# brainrumi (wip)
a brainfuck interpreter / language shell written in go !

## usage
to enter the REPL:
```
brainrumi
```
you can enter `q`, `quit` or `exit` to end the session, or `r` to toggle display modes.

to run a program written in brainfuck:
```
brainrumi path/to/program
```
you can also optionally use the `-r` flag - standard output will be displayed raw instead of being encoded in UTF-8.
<br>the `-d` flag allows you to view the execution time of each evaluation.

## technical stuff
cells are stored in memory as signed 32-bit integers, and the interpreter is **not** limited to 30,000 cells. negative pointers do **not** exist in this implementation.
<br>generally, brainrumi aims to be as faithful as possible to the original specifications (other than the 32 bit part of course.)

## notes
THE `,` OPERATOR MIGHT ACT A LITTLE WEIRD FOR NOW - I'M SORRY ... like most of the example brainfuck code i've tested that involves any kind of user input is fucked in some way or another.
<br>this does **NOT** support compiling into machine code (yet?)
<br>this is a spiritual successor of sorts to my [deadfish interpreter](https://github.com/snarkb0t/rumifish).
<br>credits to my dog, rumi, for gracefully lending her name to be used in this project !
