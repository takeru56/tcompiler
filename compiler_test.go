package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestCompile(t *testing.T) {
	cases := []struct {
		source   string
		bytecode []byte
	}{
		{"23", []byte{0, 0, 23, 5}},
		{"1+1", []byte{0, 0, 1, 0, 0, 1, 1, 5}},
		{"1-1", []byte{0, 0, 1, 0, 0, 1, 2, 5}},
		{"1*1", []byte{0, 0, 1, 0, 0, 1, 3, 5}},
		{"1/1", []byte{0, 0, 1, 0, 0, 1, 4, 5}},
		{"a = 1", []byte{0, 0, 1, 7, 0, 5}},
		{"a = 1 b = 2 b", []byte{0, 0, 1, 7, 0, 0, 0, 2, 7, 1, 6, 1, 5}},
	}

	for _, c := range cases {
		out, err := exec.Command("go", "run", ".", c.source).Output()
		if err != nil {
			log.Fatal(err)
		}
		s := ""
		for _, b := range c.bytecode {
			s += fmt.Sprintf("%02x", b)
		}

		if string(out) != s {
			fmt.Println("expected: " + s)
			fmt.Println("but actual: " + string(out))
			t.Error("not match\n")
		}
	}
}
