package main

import "fmt"

// =============================================================================
// EXKURS: Wie die Laufzeit mit der Datenmenge wächst
// =============================================================================
//
// Das Programm gehört zur HTML-Seite exkurse/exkurs_laufzeit_wachstum.html.
// Bezug zum Kurstext v3.0.6: Abschnitt 8.4, Seiten 221–226.
//
// Es zählt, wie viele Vergleiche vier Algorithmen brauchen,
// wenn das Array 10, 100 oder 1000 Elemente hat:
//
//    - lineare Suche:   n Vergleiche            O(n)
//    - binäre Suche:    etwa log2(n) Vergleiche O(log n)
//    - Bubble Sort:     n·(n-1)/2 Vergleiche    O(n²)
//    - Selection Sort:  n·(n-1)/2 Vergleiche    O(n²)
//
// Gesucht wird jedes Mal ein Wert, der im Array fehlt.
// Das ist für beide Suchen der ungünstigste Fall.
//
// Aufruf:  go run laufzeit_vergleiche_zaehlen.go
// =============================================================================

const maxAnzahl = 1000

// mischen füllt die ersten anzahl Elemente mit den Zahlen 0 bis anzahl-1
// und bringt sie in eine gemischte Reihenfolge.
// Die Reihenfolge ist bei jedem Lauf dieselbe, damit die Zahlen zur HTML-Seite passen.
func mischen(a *[maxAnzahl]int, anzahl int) {
	for i := 0; i < anzahl; i++ {
		a[i] = i
	}
	zufall := 20261011
	for i := anzahl - 1; i > 0; i-- {
		zufall = zufall * 16807 % 2147483647
		j := zufall % (i + 1)
		hilf := a[i]
		a[i] = a[j]
		a[j] = hilf
	}
}

// bubbleSort sortiert die ersten anzahl Elemente aufsteigend.
// Ein Durchgang vergleicht Nachbarn und schiebt das größte Element ans Ende.
func bubbleSort(a *[maxAnzahl]int, anzahl int) (vergleiche, vertauschungen int) {
	for ende := anzahl - 1; ende > 0; ende-- {
		for i := 0; i < ende; i++ {
			vergleiche++
			if a[i] > a[i+1] {
				hilf := a[i]
				a[i] = a[i+1]
				a[i+1] = hilf
				vertauschungen++
			}
		}
	}
	return
}

// selectionSort sortiert die ersten anzahl Elemente aufsteigend.
// Ein Durchgang sucht das kleinste Element im Rest und setzt es nach vorn.
func selectionSort(a *[maxAnzahl]int, anzahl int) (vergleiche, vertauschungen int) {
	for start := 0; start < anzahl-1; start++ {
		kleinstes := start
		for i := start + 1; i < anzahl; i++ {
			vergleiche++
			if a[i] < a[kleinstes] {
				kleinstes = i
			}
		}
		if kleinstes != start {
			hilf := a[start]
			a[start] = a[kleinstes]
			a[kleinstes] = hilf
			vertauschungen++
		}
	}
	return
}

// lineareSuche geht das Array von vorn nach hinten durch.
// Der Index ist -1, wenn der Wert fehlt.
func lineareSuche(a *[maxAnzahl]int, anzahl int, gesucht int) (index, vergleiche int) {
	for i := 0; i < anzahl; i++ {
		vergleiche++
		if a[i] == gesucht {
			return i, vergleiche
		}
	}
	return -1, vergleiche
}

// binaereSuche sieht in der Mitte nach und wirft die falsche Hälfte weg.
// Das Array muss aufsteigend sortiert sein.
// Gezählt wird ein Vergleich je Blick in die Mitte.
func binaereSuche(a *[maxAnzahl]int, anzahl int, gesucht int) (index, vergleiche int) {
	links := 0
	rechts := anzahl - 1
	for links <= rechts {
		mitte := (links + rechts) / 2
		vergleiche++
		if a[mitte] == gesucht {
			return mitte, vergleiche
		}
		if a[mitte] < gesucht {
			links = mitte + 1
		} else {
			rechts = mitte - 1
		}
	}
	return -1, vergleiche
}

func main() {
	var a [maxAnzahl]int

	fmt.Println("Vergleiche je Algorithmus")
	fmt.Println()
	fmt.Printf("%6s %15s %14s %12s %15s\n", "n", "lineare Suche", "binäre Suche", "Bubble Sort", "Selection Sort")

	for _, n := range [3]int{10, 100, 1000} {
		mischen(&a, n)
		bubbleVergleiche, bubbleTausch := bubbleSort(&a, n)

		mischen(&a, n)
		selectionVergleiche, selectionTausch := selectionSort(&a, n)

		// Das Array ist jetzt sortiert. Der Wert n kommt in ihm nicht vor.
		_, linear := lineareSuche(&a, n, n)
		_, binaer := binaereSuche(&a, n, n)

		fmt.Printf("%6d %15d %14d %12d %15d\n", n, linear, binaer, bubbleVergleiche, selectionVergleiche)
		fmt.Printf("%6s %15s %14s %12d %15d   Vertauschungen\n", "", "", "", bubbleTausch, selectionTausch)
	}
}
