package filter

import (
	"dcd/line"
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	ls := make([]line.Line, 6)
	ls[0].PushBytes([]byte("zzz\\abcdefghijklmn"))
	ls[1].PushBytes([]byte("zzz\\abcdefghijklm"))
	ls[2].PushBytes([]byte("zzz\\abcdefghijkl"))
	ls[3].PushBytes([]byte("zzz\\abcdefghijk"))
	ls[4].PushBytes([]byte("zzz\\abcdefghij"))
	ls[5].PushBytes([]byte("zzz\\abcdefghi"))
	var f Filter
	f.Init(ls)
	f.Type = PathType
	fmt.Println("------1")
	f.Push("z")
	f.Push("a")
	f.Push("b")
	f.Push("c")
	f.Push("d")
	f.Push("e")
	f.Push("f")
	f.Push("g")
	f.Push("h")
	f.Print()

	fmt.Println("------2")
	f.Push("i")
	f.Print()

	fmt.Println("------3")
	f.Push("j")
	f.Print()

	fmt.Println("------4")
	f.Push("k")
	f.Print()

	fmt.Println("------5")
	f.Push("l")
	f.Print()

	fmt.Println("------6")
	f.Push("m")
	f.Print()

	fmt.Println("------7")
	f.Push("n")
	f.Print()

	fmt.Println("------8")
	f.Push("o")
	f.Print()
}
