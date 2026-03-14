# brainrumi (wip)
a brainfuck interpreter / language shell written in go !

## manual installation
for now, you can get to using this interpreter with these commands (with go installed):
```bash
git clone https://github.com/snarkb0t/brainrumi.git
cd brainrumi
go build .
```
you can move the resulting executable to your `bin` folder.

## usage
to enter the REPL:
```bash
brainrumi
```
you can enter `q`, `quit` or `exit` to end the session, or `r` to toggle display modes.

to run a program written in brainfuck:
```bash
brainrumi path/to/program.bf
```
you can also optionally use the `-r` flag - standard output will be displayed raw instead of being encoded in UTF-8.
<br>the `-d` flag allows you to view the execution time of each evaluation.
<br>to toggle the (experimental) optimizing interpreter, you can also use the `-o` flag. (more on this in the **notes** section.)

## technical stuff
cells are stored in memory as signed 32-bit integers, and the interpreter is **not** limited to 30,000 cells. negative pointers do **not** exist in this implementation.
<br>generally, brainrumi aims to be as faithful as possible to the original specifications (other than the 32 bit part of course.)

## notes
1. THE `,` OPERATOR MIGHT ACT A LITTLE WEIRD FOR NOW - I'M SORRY ... like most of the example brainfuck code i've tested that involves any kind of user input is fucked in some way or another.
2. the optimizing interpreter is still in work in progress (the entire interpreter is tbf) - it'll be more effective in some programs and less in others. for instance, you won't see much difference in a simple hello world program (might even be a little slower) - but a mandelbrot fractal program took 11.45 seconds to run, whereas that same program took 1 minute and 17.52 seconds in the traditional interpreter.
3. this does **NOT** support compiling into machine code (yet?)
4. this is a spiritual successor of sorts to my [deadfish interpreter](https://github.com/snarkb0t/rumifish).
<br><br>credits to my dog, rumi, for gracefully lending her name to be used in this project !
