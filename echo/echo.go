package main

import (
	"fmt"
	"github.com/guonaihong/flag"
	"os"
	"strconv"
)

func isxdigit(b byte) bool {
	if b >= '0' && b <= '9' ||
		b >= 'a' && b <= 'f' ||
		b >= 'A' && b <= 'F' {
		return true
	}
	return false
}

func isoctal(b byte) bool {
	if b >= '0' && b <= '7' {
		return true
	}

	return false
}

func isoctalStr(s string, max int) (i int, haveOctal bool) {
	for i = 0; i < len(s); i++ {
		if i > max {
			return i - 1, haveOctal
		}

		if !isoctal(s[i]) {
			return i - 1, haveOctal
		}

		haveOctal = true
	}

	return i, haveOctal
}

func isxdigitStr(s string, max int) (i int, haveHex bool) {

	for i = 0; i < len(s); i++ {
		if i > max {
			return i - 1, haveHex
		}

		if !isxdigit(s[i]) {
			return i - 1, haveHex
		}

		haveHex = true
	}

	return i, haveHex
}

func main() {
	newLine := flag.Bool("n", false, "do not output the trailing newline")
	enable := flag.Bool("e", false, "enable interpretation of backslash escapes")
	disable := flag.Bool("E", true, "disable interpretation of backslash escapes (default)")
	flag.Parse()

	args := flag.Args()

	c0 := uint64(0)
	var err error

	defer func() {
		if *newLine == false {
			fmt.Printf("\n")
		}
	}()

	if *enable {
		printSlash := false
		for _, s := range args {
			for i := 0; i < len(s); i++ {
				c := s[i]

				if c == '\\' && i < len(s) {
					i++
					if i >= len(s) {
						fmt.Printf("\\")
						goto notAnEscape
					}

					c = s[i]
					switch c {
					case 'a':
						c = '\a'
					case 'b':
						c = '\b'
					case 'c':
						return
					case 'e':
						c = '\x1B'
					case 'f':
						c = '\f'
					case 'n':
						c = '\n'
					case 'r':
						c = '\r'
					case 't':
						c = '\t'
					case 'v':
						c = '\v'
					case 'x':
						if i+1 >= len(s) {
							printSlash = true
							goto notAnEscape
						}

						n, haveHex := isxdigitStr(s[i+1:], 2)
						if !haveHex {
							printSlash = true
							goto notAnEscape
						}

						c0, err = strconv.ParseUint(s[i+1:i+1+n], 16, 32)
						if err != nil {
							printSlash = true
							goto notAnEscape
						}

						i = i + 1 + n - 1
						c = byte(c0)

					case '0':
						if i+1 >= len(s) {
							printSlash = true
							goto notAnEscape
						}

						n, haveOctal := isoctalStr(s[i+1:], 3)
						if !haveOctal {
							printSlash = true
							goto notAnEscape
						}

						c0, err = strconv.ParseUint(s[i+1:i+1+n], 8, 32)
						if err != nil {
							printSlash = true
							goto notAnEscape
						}

						i = i + 1 + n - 1
						c = byte(c0)
					case '\\':
					default:
						fmt.Print("\\")
					}

				}

			notAnEscape:
				if printSlash {
					fmt.Printf("\\")
					printSlash = false
				}

				// fmt.Printf("%c") is not the same as the putchar output in c
				// in go fmt.Printf("%c\n", 172) -->  ¬
				// in c  putchar(172)            -->  ?
				os.Stdout.Write([]byte{c})
			}
			fmt.Print(" ")
		}
		return
	}

	if *disable {
		for i, s := range args {
			fmt.Print(s)
			if i+1 != len(args) {
				fmt.Printf(" ")
			}
		}
	}

}