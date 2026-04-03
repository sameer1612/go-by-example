package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains:", strings.Contains("test", "es"))
	fmt.Println("Count:", strings.Count("test", "e"))
	fmt.Println("HasPrefix:", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix:", strings.HasSuffix("test", "st"))
	fmt.Println("Index:", strings.Index("test", "e"))
	fmt.Println("Join:", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:", strings.Repeat("a", 5))
	fmt.Println("Replace:", strings.Replace("aaaaa", "a", "b", 2))
	fmt.Println("ReplaceAll:", strings.ReplaceAll("aaaaa", "a", "b"))
	fmt.Println("Split:", strings.Split("a-b-c", "-"))
	fmt.Println("ToLower:", strings.ToLower("TesT"))
	fmt.Println("ToUpper:", strings.ToUpper("TesT"))
}
