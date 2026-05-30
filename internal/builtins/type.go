package BuiltIns

import "io"

var builtins = []string{
	"type", "pwd", "echo", "exit",
}


func GetType(w io.Writer, tp string) bool {
	for _, builtin := range builtins {
		if tp == builtin {
			io.WriteString(w,tp +" is a shell builtin\n")
			return true
		}
	}
	return false
}
