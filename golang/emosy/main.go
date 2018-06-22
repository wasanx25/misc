package main

import (
	"syscall"
)

func main() {
	syscall.Write(syscall.Stdin, []byte("\x1b[4;30;46m"))
	syscall.Write(syscall.Stdin, []byte("\033[5G"))
	syscall.Write(syscall.Stdin, []byte("\033[K"))
	syscall.Write(syscall.Stdin, []byte("HelloWorld\n"))
	syscall.Write(syscall.Stdin, []byte("\x1b[0m"))
}
