package main

/*
#include <elf.h>
#include <unistd.h>
#include <string.h>
*/
import "C"
import (
	"flag"
	"fmt"
	"os"
	"unsafe"
)

func isElf(elfHeader *C.Elf32_Ehdr) bool {

	elfcode := []byte("\177ELF")
	res := C.strncmp((*C.char)(unsafe.Pointer(&elfHeader.e_ident[0])), (*C.char)(unsafe.Pointer(&elfcode[0])), 4)

	if res == 0 {
		return true
	}

	return false
}

func openElf(file_name string) C.int {

	file, err := os.Open(file_name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return -1
	}

	return C.int(file.Fd())
}

func readElf32(fd C.int) *C.Elf32_Ehdr {

	var elfHeader C.Elf32_Ehdr
	if C.read(fd, unsafe.Pointer(&elfHeader), C.ulong(unsafe.Sizeof(elfHeader))) < 0 {
		return nil
	}

	return &elfHeader
}

func readElf64(fd C.int) *C.Elf64_Ehdr {

	var elfHeader C.Elf64_Ehdr

	if C.read(fd, unsafe.Pointer(&elfHeader), C.ulong(unsafe.Sizeof(elfHeader))) < 0 {
		return nil
	}

	return &elfHeader
}

func isElf64(elfHeader *C.Elf32_Ehdr) bool {
	if elfHeader.e_ident[C.EI_CLASS] == C.ELFCLASS64 {
		return true
	}

	return false
}

func printElf64Sections(elfHeader *C.Elf64_Ehdr) {

	fmt.Println("storage class: ")
	switch elfHeader.e_ident[C.EI_CLASS] {
	case C.ELFCLASS32:
		fmt.Println("  32bit elf")

	case C.ELFCLASS64:
		fmt.Println("  64bit elf")

	default:
		fmt.Println("  invalid class")
	}

	fmt.Println("Data format: ")
	switch elfHeader.e_ident[C.EI_DATA] {
	case C.ELFDATA2LSB:
		fmt.Println("  little endian")

	case C.ELFDATA2MSB:
		fmt.Println("  little endian")

	default:
		fmt.Println("  invalid format")
	}

	fmt.Println("OS: ")
	switch elfHeader.e_ident[C.EI_OSABI] {
	case C.ELFOSABI_SYSV:
		fmt.Println("  UNIX system v")

	case C.ELFOSABI_NETBSD:
		fmt.Println("  NetBSD")

	case C.ELFOSABI_FREEBSD:
		fmt.Println("  FreeBSD")

	case C.ELFOSABI_LINUX:
		fmt.Println("  Linux")

	case C.ELFOSABI_SOLARIS:
		fmt.Println("  Solaris")

	case C.ELFOSABI_ARM:
		fmt.Println("  ARM")
	}

	fmt.Println("File type: ")
	switch elfHeader.e_type {
	case C.ET_NONE:
		fmt.Println("  N/A")
	case C.ET_REL:
		fmt.Println("  Relocatable")
	case C.ET_EXEC:
		fmt.Println("  Executable")
	case C.ET_DYN:
		fmt.Println("  Shared Object")
	case C.ET_CORE:
		fmt.Println("  Core file")
	default:
		fmt.Printf("  Unknown type: 0x%x\n", elfHeader.e_type)
	}

	fmt.Printf("Entry Point:\n  0x%08x\n", elfHeader.e_entry)
	fmt.Printf("ELF header size:\n  0x%08x bytes\n", elfHeader.e_ehsize)

	fmt.Printf("Program header start addr:\n  0x%08x\n", elfHeader.e_phoff)
	fmt.Printf("Program entries:\n  0x%08x\n", elfHeader.e_phnum)
	fmt.Printf("Program size:\n  0x%08x\n", elfHeader.e_phentsize)
}

func main() {

	printComm := flag.Bool("p", false, "print the elf header")
	disassComm := flag.Bool("d", false, "disassemble the elf file")

	flag.Parse()

	file_name := flag.Args()[0]

	fd := openElf(file_name)

	elfHeader := readElf32(fd)

	if elfHeader == nil {
		fmt.Println("Could not read the elf file")
		os.Exit(1)
	}

	if !isElf(elfHeader) {
		fmt.Fprintf(os.Stderr, "%s Not an elf file", file_name)
		os.Exit(1)
	}

	if *printComm && isElf64(elfHeader) {
		elfHeader64 := readElf64(fd)
		printElf64Sections(elfHeader64)

	} else if *disassComm {
		fmt.Println("disassable")
	}
}
