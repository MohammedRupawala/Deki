package BuiltIns

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, cmd string) {
	// fmt.Fprint(w,"\r")
	// fmt.Println("Command is " + cmd)
	fmt.Fprint(w, cmd)
	fmt.Fprint(w, "\n")
}

