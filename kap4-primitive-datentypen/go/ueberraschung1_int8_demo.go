package main

import "fmt"

// Überraschung 1: 127 + 1 mit int8. Schätze die Ausgabe, bevor du es ausführst.

func main() {
	var a int8 = 127
	a = a + 1
	fmt.Println(a)
}
