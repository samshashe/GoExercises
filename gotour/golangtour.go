package main

import (
 "fmt"
 )

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error){
	for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 66)
	a,_ := r.Read(b)
	fmt.Println(string(b), a)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (int,error){
	res,err := rr.r.Read(p)
	for i := range(p) {
      p[i] = rot13(p[i])
    }
	return res,err
}
func rot13(b byte) (c byte) {
    switch {
      case b >= 'A' && b <= 'Z':
        c = (b - 'A' + 13) % 26 + 'A'
      case b >= 'a' && b <= 'z':
        c = (b - 'a' + 13) % 26 + 'a'
      default:
        c = b  
    }
    return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
