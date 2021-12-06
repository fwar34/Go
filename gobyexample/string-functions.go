package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:", s.Repeat("te", 5))
	p("Replace:", s.Replace("foo", "o", "0", -1))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("a-b-c-d", "-"))
	p("ToLower:", s.ToLower("ADSF"))
	p("ToUpper:", s.ToUpper("sdfa"))

	p("Len:", len("hello"))
	p("Char:", "sdfsf"[1])
}
