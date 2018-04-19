package util

import (
	"bufio"
	"fmt"
	"os"
)

func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func ReadFile(file string) *chan string {
	res := make(chan string)
	go func(res *chan string) {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		r := bufio.NewReader(f)
		s, e := readln(r)
		for e == nil {
			*res <- s
			s, e = readln(r)
		}
		close(*res)

	}(&res)
	return &res

}
