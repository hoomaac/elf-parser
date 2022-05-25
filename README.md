# elf-parser
A simple ELF parser in Go

### What is it all about?

Just for illutration of how Go can work with C stuffs, e.g. C headers, functions.

### Build

For static build you can run
`
    make static
`
build with memory status
`
    make static_memcheck
`

for dynamic build, change above statics with **dynamic**

### Test

For benchmark just enter
`
    make benchmark
`
and for tests
`
    make test
`

![elf_parser_print](/screenshot/print_header.png?raw=true "print header go")