package types

import "fmt"

type Scanner string

func (*Scanner) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing %s\n", p)
	return len(p), nil
}
