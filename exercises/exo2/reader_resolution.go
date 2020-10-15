
//solution made in lesson 2

package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}


func (tr spaceEraser) Read (characters_buf []byte) (int,error) {
/*We replace the positions of spaces by other characteres in the buffer 
and we return an int which delimits the length of the characters without spaces.*/
	length,err := tr.r.Read(characters_buf)
	length_noblankspaces := 0
	for i:=0; i <length; i++ {
		if characters_buf[i] != 32 {  //UTF-8 space = 32
			characters_buf[length_noblankspaces] = characters_buf[i]
			length_noblankspaces++
		}
	}
	return length_noblankspaces, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
