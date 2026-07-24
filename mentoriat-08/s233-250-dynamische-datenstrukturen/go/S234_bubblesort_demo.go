package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Bubblesort auf Verbund-Arrays (Abschnitt 8.4.2)
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Bubblesort-Algorithmus
//    - Vergleicht systematisch benachbarte Elemente
//    - Vertauscht bei falscher Sortierung
//    - Größtes Element "blubbert" nach rechts
//
// 2. Laufzeitverhalten
//    - Anzahl Vergleiche: n·(n-1)/2 = O(n²)
//    - Quadratische Laufzeit
//    - Ungünstig für große Arrays
//
// 3. Optimierte Version
//    - Abbruch wenn keine Vertauschungen mehr
//    - Erkennt bereits sortierte Bereiche
//
// 4. Visualisierung
//    - Zeigt jeden Schritt des Sortierens
//    - Zählt Vergleiche und Vertauschungen
//
// 5. Verschiedene Testfälle
//    - Unsortiert, bereits sortiert, rückwärts sortiert
//
// Ausführen mit: go run S234_bubblesort_demo.go
//
// =============================================================================

// person: Verbund-Typ für Telefonbuch-Eintrag
type person struct {
	name          string
	telefonnummer int
}

// Globale Zähler für Statistiken
var vergleiche int
var vertauschungen int

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Bubblesort auf Verbund-Arrays ===\n")

	// -------------------------------------------------------------------------
	// 1. Grundprinzip von Bubblesort
	// -------------------------------------------------------------------------

	fmt.Println("1. Grundprinzip von Bubblesort:")
	fmt.Println("   -----------------------------")
	fmt.Println()

	fmt.Println("   Funktionsweise:")
	fmt.Println("   1. Durchlaufe Array von links nach rechts")
	fmt.Println("   2. Vergleiche jedes Paar benachbarter Elemente")
	fmt.Println("   3. Vertausche wenn falsch sortiert (links > rechts)")
	fmt.Println("   4. Größtes Element \"blubbert\" an das Ende")
	fmt.Println("   5. Wiederhole für verbleibende Elemente")
	fmt.Println()

	fmt.Println("   Beispiel mit Zahlen [5, 2, 8, 1, 9]:")
	fmt.Println("   Durchgang 1: [5,2,8,1,9] -> [2,5,8,1,9] -> [2,5,1,8,9] -> [2,5,1,8,9]")
	fmt.Println("                                                              9 ist fertig!")
	fmt.Println("   Durchgang 2: [2,5,1,8] -> [2,1,5,8] -> [2,1,5,8]")
	fmt.Println("                                           8 ist fertig!")
	fmt.Println("   Durchgang 3: [2,1,5] -> [1,2,5]")
	fmt.Println("                            5 ist fertig!")
	fmt.Println("   Durchgang 4: [1,2]")
	fmt.Println("                 Fertig sortiert: [1,2,5,8,9]")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Unsortiertes Array sortieren
	// -------------------------------------------------------------------------

	fmt.Println("2. Beispiel 1: Unsortiertes Array sortieren:")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	unsortiert := [10]person{
		{"Schmidt", 78901},
		{"Müller", 67890},
		{"Weber", 90123},
		{"Bauer", 12345},
		{"Klein", 45678},
		{"Wolf", 11111},
		{"Fischer", 23456},
		{"Meier", 56789},
		{"Hoffmann", 34567},
		{"Schneider", 89012},
	}

	fmt.Println("   Vor dem Sortieren:")
	printArray(&unsortiert)
	fmt.Println()

	vergleiche = 0
	vertauschungen = 0

	bubblesort(&unsortiert)

	fmt.Println()
	fmt.Println("   Nach dem Sortieren:")
	printArray(&unsortiert)
	fmt.Println()

	fmt.Printf("   Statistik:\n")
	fmt.Printf("   - Vergleiche:      %d\n", vergleiche)
	fmt.Printf("   - Vertauschungen:  %d\n", vertauschungen)
	fmt.Printf("   - Erwartete Vergleiche: n·(n-1)/2 = 10·9/2 = %d\n", 10*9/2)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Visualisierung des Sortiervorgangs
	// -------------------------------------------------------------------------

	fmt.Println("3. Visualisierung des Sortiervorgangs:")
	fmt.Println("   ------------------------------------")
	fmt.Println()

	klein := [5]person{
		{"Meier", 111},
		{"Bauer", 222},
		{"Weber", 333},
		{"Klein", 444},
		{"Müller", 555},
	}

	fmt.Println("   Kleines Array für detaillierte Visualisierung:")
	fmt.Println()
	fmt.Println("   Vor dem Sortieren:")
	printArrayKompakt(&klein)
	fmt.Println()

	bubblesortVisualisiert(&klein)

	fmt.Println()
	fmt.Println("   Nach dem Sortieren:")
	printArrayKompakt(&klein)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Bereits sortiertes Array (schlechtester Fall für Bubblesort)
	// -------------------------------------------------------------------------

	fmt.Println("4. Beispiel 2: Bereits sortiertes Array:")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	sortiert := [10]person{
		{"Bauer", 12345},
		{"Fischer", 23456},
		{"Hoffmann", 34567},
		{"Klein", 45678},
		{"Meier", 56789},
		{"Müller", 67890},
		{"Schmidt", 78901},
		{"Schneider", 89012},
		{"Weber", 90123},
		{"Wolf", 11111},
	}

	fmt.Println("   Vor dem Sortieren (bereits sortiert):")
	printArray(&sortiert)
	fmt.Println()

	vergleiche = 0
	vertauschungen = 0

	bubblesort(&sortiert)

	fmt.Println()
	fmt.Println("   Nach dem Sortieren (unverändert):")
	printArray(&sortiert)
	fmt.Println()

	fmt.Printf("   Statistik:\n")
	fmt.Printf("   - Vergleiche:      %d\n", vergleiche)
	fmt.Printf("   - Vertauschungen:  %d (optimal!)\n", vertauschungen)
	fmt.Println()

	fmt.Println("   PROBLEM:")
	fmt.Println("   ✗ Trotz bereits sortiertem Array wurden 45 Vergleiche gemacht!")
	fmt.Println("   ✗ Algorithmus erkennt nicht, dass bereits sortiert ist")
	fmt.Println("   ✗ Verschwendung von Rechenzeit")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Optimierte Version mit Early-Exit
	// -------------------------------------------------------------------------

	fmt.Println("5. Optimierte Version mit Early-Exit:")
	fmt.Println("   -----------------------------------")
	fmt.Println()

	sortiert2 := [10]person{
		{"Bauer", 12345},
		{"Fischer", 23456},
		{"Hoffmann", 34567},
		{"Klein", 45678},
		{"Meier", 56789},
		{"Müller", 67890},
		{"Schmidt", 78901},
		{"Schneider", 89012},
		{"Weber", 90123},
		{"Wolf", 11111},
	}

	fmt.Println("   Optimierung:")
	fmt.Println("   ✓ Wenn in einem Durchlauf keine Vertauschungen mehr:")
	fmt.Println("   ✓ -> Array ist sortiert -> Abbruch!")
	fmt.Println()

	fmt.Println("   Vor dem Sortieren (bereits sortiert):")
	printArray(&sortiert2)
	fmt.Println()

	vergleiche = 0
	vertauschungen = 0

	bubblesortOptimiert(&sortiert2)

	fmt.Println()
	fmt.Println("   Nach dem Sortieren:")
	printArray(&sortiert2)
	fmt.Println()

	fmt.Printf("   Statistik:\n")
	fmt.Printf("   - Vergleiche:      %d\n", vergleiche)
	fmt.Printf("   - Vertauschungen:  %d\n", vertauschungen)
	fmt.Printf("   - Durchläufe:      1 (statt 9!)\n")
	fmt.Println()

	fmt.Println("   VORTEIL:")
	fmt.Println("   ✓ Nur 9 Vergleiche statt 45!")
	fmt.Println("   ✓ Frühzeitiger Abbruch bei bereits sortiertem Array")
	fmt.Println("   ✓ Beste Laufzeit: O(n) statt O(n²)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Vergleich: Normal vs. Optimiert
	// -------------------------------------------------------------------------

	fmt.Println("6. Vergleich: Original vs. Optimierte Version:")
	fmt.Println("   -------------------------------------------")
	fmt.Println()

	// Test mit verschiedenen Szenarien
	szenarien := []struct {
		name  string
		array [10]person
	}{
		{
			"Rückwärts sortiert (schlechtester Fall)",
			[10]person{
				{"Wolf", 11111},
				{"Weber", 90123},
				{"Schneider", 89012},
				{"Schmidt", 78901},
				{"Müller", 67890},
				{"Meier", 56789},
				{"Klein", 45678},
				{"Hoffmann", 34567},
				{"Fischer", 23456},
				{"Bauer", 12345},
			},
		},
		{
			"Bereits sortiert (bester Fall)",
			[10]person{
				{"Bauer", 12345},
				{"Fischer", 23456},
				{"Hoffmann", 34567},
				{"Klein", 45678},
				{"Meier", 56789},
				{"Müller", 67890},
				{"Schmidt", 78901},
				{"Schneider", 89012},
				{"Weber", 90123},
				{"Wolf", 11111},
			},
		},
		{
			"Teilweise sortiert",
			[10]person{
				{"Bauer", 12345},
				{"Fischer", 23456},
				{"Weber", 90123},
				{"Klein", 45678},
				{"Meier", 56789},
				{"Hoffmann", 34567},
				{"Schmidt", 78901},
				{"Schneider", 89012},
				{"Müller", 67890},
				{"Wolf", 11111},
			},
		},
	}

	for _, szenario := range szenarien {
		fmt.Printf("   Szenario: %s\n", szenario.name)

		// Original-Version
		arr1 := szenario.array
		vergleiche = 0
		vertauschungen = 0
		bubblesort(&arr1)
		v1, t1 := vergleiche, vertauschungen

		// Optimierte Version
		arr2 := szenario.array
		vergleiche = 0
		vertauschungen = 0
		bubblesortOptimiert(&arr2)
		v2, t2 := vergleiche, vertauschungen

		fmt.Printf("   Original:   %d Vergleiche, %d Vertauschungen\n", v1, t1)
		fmt.Printf("   Optimiert:  %d Vergleiche, %d Vertauschungen\n", v2, t2)
		fmt.Printf("   Ersparnis:  %.0f%% weniger Vergleiche\n",
			(1.0-float64(v2)/float64(v1))*100)
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 7. Laufzeitanalyse
	// -------------------------------------------------------------------------

	fmt.Println("7. Laufzeitanalyse:")
	fmt.Println("   ----------------")
	fmt.Println()

	fmt.Println("   Anzahl Vergleiche (bei n Elementen):")
	fmt.Println("   Durchgang 1: (n-1) Vergleiche")
	fmt.Println("   Durchgang 2: (n-2) Vergleiche")
	fmt.Println("   Durchgang 3: (n-3) Vergleiche")
	fmt.Println("   ...")
	fmt.Println("   Durchgang (n-1): 1 Vergleich")
	fmt.Println()
	fmt.Println("   Gesamt: (n-1) + (n-2) + ... + 1 = n·(n-1)/2")
	fmt.Println()

	fmt.Println("   Beispielrechnung für n=10:")
	fmt.Printf("   10·9/2 = %d Vergleiche\n", 10*9/2)
	fmt.Println()

	fmt.Println("   Verschiedene Array-Größen:")
	for _, n := range []int{10, 50, 100, 500, 1000} {
		vergleicheGesamt := n * (n - 1) / 2
		fmt.Printf("   n = %-4d: %7d Vergleiche\n", n, vergleicheGesamt)
	}
	fmt.Println()

	fmt.Println("   Laufzeit-Klasse: O(n²)")
	fmt.Println("   ✓ Quadratisches Wachstum")
	fmt.Println("   ✗ Ineffizient für große Arrays")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Kombination mit binärer Suche
	// -------------------------------------------------------------------------

	fmt.Println("8. Praktische Anwendung: Sortieren + Suchen:")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	telefonbuch := [10]person{
		{"Weber", 90123},
		{"Bauer", 12345},
		{"Meier", 56789},
		{"Schmidt", 78901},
		{"Klein", 45678},
		{"Müller", 67890},
		{"Wolf", 11111},
		{"Fischer", 23456},
		{"Schneider", 89012},
		{"Hoffmann", 34567},
	}

	fmt.Println("   Unsortiertes Telefonbuch:")
	printArray(&telefonbuch)
	fmt.Println()

	fmt.Println("   1. Sortieren mit Bubblesort...")
	bubblesortOptimiert(&telefonbuch)
	fmt.Println()

	fmt.Println("   Sortiertes Telefonbuch:")
	printArray(&telefonbuch)
	fmt.Println()

	fmt.Println("   2. Binäre Suche nach \"Müller\"...")
	// Hinweis: Hier würde die binäre Suche aus dem vorherigen Abschnitt verwendet
	fmt.Println("   -> Jetzt kann effizient gesucht werden!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Code-Erklärung
	// -------------------------------------------------------------------------

	fmt.Println("9. Code-Erklärung (Original-Version):")
	fmt.Println("   -----------------------------------")
	fmt.Println()

	fmt.Println("   func bubblesort(a *[10]person) {")
	fmt.Println("       // Äußere Schleife: noch zu sortierende Elemente")
	fmt.Println("       for n := len(a); n > 1; n-- {")
	fmt.Println()
	fmt.Println("           // Innere Schleife: Sortieren durch Vertauschen")
	fmt.Println("           for i := 0; i < n-1; i++ {")
	fmt.Println()
	fmt.Println("               // Vergleiche benachbarte Elemente")
	fmt.Println("               if a[i].name > a[i+1].name {")
	fmt.Println()
	fmt.Println("                   // Vertausche bei falscher Reihenfolge")
	fmt.Println("                   a[i], a[i+1] = a[i+1], a[i]")
	fmt.Println("               }")
	fmt.Println("           }")
	fmt.Println("       }")
	fmt.Println("   }")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Zusammenfassung
	// -------------------------------------------------------------------------

	fmt.Println("10. Zusammenfassung:")
	fmt.Println("    ---------------")
	fmt.Println()

	fmt.Println("   Bubblesort:")
	fmt.Println("   ✓ Einfacher Sortieralgorithmus")
	fmt.Println("   ✓ Vergleicht benachbarte Elemente")
	fmt.Println("   ✓ Größtes Element \"blubbert\" nach rechts")
	fmt.Println("   ✓ Laufzeit: O(n²)")
	fmt.Println()

	fmt.Println("   Eigenschaften:")
	fmt.Println("   ✓ Stabil (Reihenfolge gleicher Elemente bleibt)")
	fmt.Println("   ✓ In-Place (kein zusätzlicher Speicher nötig)")
	fmt.Println("   ✗ Ineffizient für große Arrays")
	fmt.Println()

	fmt.Println("   Optimierung:")
	fmt.Println("   ✓ Early-Exit bei keinen Vertauschungen")
	fmt.Println("   ✓ Beste Laufzeit: O(n)")
	fmt.Println("   ✓ Schlechteste Laufzeit: O(n²)")
	fmt.Println()

	fmt.Println("   Anwendung:")
	fmt.Println("   ✓ Kleine Arrays (< 50 Elemente)")
	fmt.Println("   ✓ Fast sortierte Arrays")
	fmt.Println("   ✓ Lernzwecke")
	fmt.Println("   ✗ Große Arrays (besser: Quicksort, Mergesort)")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Bubblesort-Funktionen
// =============================================================================

// bubblesort: Original-Version aus dem Text (Listing 8.2)
//
// Sortiert das Verbund-Array aufsteigend nach dem Feld name
func bubblesort(a *[10]person) {
	// äußere Schleife: noch zu sortierende Elemente
	for n := len(a); n > 1; n-- {
		// innere Schleife: Sortieren durch Vertauschen
		for i := 0; i < n-1; i++ {
			vergleiche++ // Statistik
			if a[i].name > a[i+1].name {
				// vertausche benachbarte Elemente
				a[i], a[i+1] = a[i+1], a[i]
				vertauschungen++ // Statistik
			}
		}
	}
}

// bubblesortOptimiert: Optimierte Version mit Early-Exit
//
// Bricht ab, wenn in einem Durchlauf keine Vertauschungen mehr stattfinden
func bubblesortOptimiert(a *[10]person) {
	// äußere Schleife: noch zu sortierende Elemente
	for n := len(a); n > 1; n-- {
		vertauscht := false // Flag für Vertauschungen in diesem Durchlauf

		// innere Schleife: Sortieren durch Vertauschen
		for i := 0; i < n-1; i++ {
			vergleiche++ // Statistik
			if a[i].name > a[i+1].name {
				// vertausche benachbarte Elemente
				a[i], a[i+1] = a[i+1], a[i]
				vertauschungen++ // Statistik
				vertauscht = true
			}
		}

		// Optimierung: Abbruch wenn keine Vertauschungen mehr
		if !vertauscht {
			// Array ist bereits sortiert!
			break
		}
	}
}

// bubblesortVisualisiert: Version mit Visualisierung für kleine Arrays
func bubblesortVisualisiert(a *[5]person) {
	fmt.Println("   Schritt-für-Schritt Visualisierung:")
	fmt.Println()

	durchlauf := 1

	// äußere Schleife: noch zu sortierende Elemente
	for n := len(a); n > 1; n-- {
		fmt.Printf("   Durchlauf %d (sortiere %d Elemente):\n", durchlauf, n)

		// innere Schleife: Sortieren durch Vertauschen
		for i := 0; i < n-1; i++ {
			fmt.Printf("      Vergleiche [%d]=%s mit [%d]=%s",
				i, a[i].name, i+1, a[i+1].name)

			if a[i].name > a[i+1].name {
				// vertausche benachbarte Elemente
				a[i], a[i+1] = a[i+1], a[i]
				fmt.Printf(" -> TAUSCH!\n")
			} else {
				fmt.Printf(" -> OK\n")
			}
		}

		fmt.Print("      Stand: ")
		printArrayKompakt(a)
		fmt.Printf("      (%s ist jetzt an Position %d sortiert)\n",
			a[n-1].name, n-1)
		fmt.Println()

		durchlauf++
	}
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// printArray: Gibt Array übersichtlich aus
func printArray(a *[10]person) {
	for i, p := range a {
		fmt.Printf("   [%d] %-12s -> %d\n", i, p.name, p.telefonnummer)
	}
}

// printArrayKompakt: Gibt Array kompakt aus (nur Namen)
func printArrayKompakt(a *[5]person) {
	fmt.Print("[")
	for i, p := range a {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s", p.name)
	}
	fmt.Println("]")
}
