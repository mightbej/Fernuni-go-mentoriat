package main

import "fmt"

// Überraschung 2: 0.1 + 0.2 == 0.3 mit float64. Schätze die Ausgabe, bevor du es ausführst.

func main() {
	var x, y float64 = 0.1, 0.2
	fmt.Println(x+y == 0.3)
}
