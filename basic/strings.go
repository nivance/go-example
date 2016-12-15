package main

//import (
//	"fmt"
//	"strings"
//)
//func init() {
//	var str string = "This is an example of a string"
//	fmt.Printf("T/F? Does the string \"%s\" hava prefix %s?", str, "Th")
//	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

//	var s string = "Hi, I'm Marc, Hi."
//	fmt.Printf("The position of \"Marc\" is :")
//	fmt.Printf("%d\n", strings.Index(s, "Marc"))
//	fmt.Printf("The position of the first instance of \"Hi\" is: ")
//	fmt.Printf("%d\n", strings.Index(s, "Hi"))
//	fmt.Printf("The position of the last instance of \"Hi\" is: ")
//	fmt.Printf("%d\n", strings.LastIndex(s, "Hi"))
//	fmt.Printf("The position of \"Burger\" is: ")
//	fmt.Printf("%d\n", strings.Index(str, "Burger"))

//	str = "Hello, how is it going, Hugo?"
//	var manyG = "gggggggggg"

//	fmt.Printf("Number of H's in %s is: ", str)
//	fmt.Printf("%d\n", strings.Count(str, "H"))

//	fmt.Printf("Number of double g's in %s is: ", manyG)
//	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

//	var origS string = "Hi there!"
//	var newS string
//	newS = strings.Repeat(origS, 3)
//	fmt.Printf("The new repeated string is: %s\n", newS)

//	str = "The quick brown fox jumps over the lazy dog"
//	sl := strings.Fields(str)
//	fmt.Printf("Splitted in slice: %v\n", sl)
//	for _, val := range sl {
//		fmt.Printf("%s - ", val)
//	}
//	fmt.Println()
//	str2 := "GO1|The ABC of Go|25"
//	sl2 := strings.Split(str2, "|")
//	fmt.Printf("Splitted in slice: %v\n", sl2)
//	for _, val := range sl2 {
//		fmt.Printf("%s - ", val)
//	}
//	fmt.Println()
//	str3 := strings.Join(sl2, ";")
//	fmt.Printf("sl2 joined by ;: %s\n", str3)
//}
