package main

import (
	"testing"
)

const elfFile = "test_bin"

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v and %v was expected to be equal", a, b)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("%v and %v was expected to be not equal", a, b)
	}
}

func TestOpenElf(t *testing.T) {

	fd := openElf(elfFile)

	assertNotEqual(t, fd, -1)

}

func TestReadElf32(t *testing.T) {

	fd := openElf(elfFile)

	assertNotEqual(t, fd, -1)

	elfHeader := readElf32(fd)

	assertNotEqual(t, elfHeader, nil)
}

func TestReadElf64(t *testing.T) {

	fd := openElf(elfFile)

	assertNotEqual(t, fd, -1)

	elfHeader := readElf64(fd)

	assertNotEqual(t, elfHeader, nil)
}

func TestIsElf(t *testing.T) {

	fd := openElf(elfFile)

	assertNotEqual(t, fd, -1)

	elfHeader := readElf32(fd)

	assertNotEqual(t, elfHeader, nil)

	res := isElf(elfHeader)

	assertEqual(t, res, true)
}

func BenchmarkReadElf32(b *testing.B) {

	var fd _Ctype_int

	for n := 0; n < b.N; n++ {

		fd = openElf(elfFile)
		readElf32(fd)
	}
}

func BenchmarkReadElf64(b *testing.B) {

	var fd _Ctype_int

	for n := 0; n < b.N; n++ {
		fd = openElf(elfFile)
		elfHeader := readElf32(fd)

		if isElf64(elfHeader) {
			readElf64(fd)
		}
	}
}
