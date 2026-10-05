package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Binäre Suche in Verbund-Arrays (Abschnitt 8.4.1)
// =============================================================================
//
// Quelle: Die Funktion rbins folgt dem Listing binSuche.go aus dem Kurstext
// „Imperative Programmierung“ (Kurs 01613) der FernUniversität in Hagen.
// Die Ausgaben und die übrigen Funktionen sind eigene Ergänzungen.
//
// Dieses Programm demonstriert:
//
// 1. Binäre Suche (Binary Search) in sortierten Arrays
//    - Rekursiver Algorithmus
//    - Halbierung des Suchbereichs in jedem Schritt
//    - Laufzeit: O(log₂ n) statt O(n)
//
// 2. Praktisches Beispiel: Telefonbuch
//    - Array von person-Verbünden (sortiert nach Name)
//    - Suche nach Name, Rückgabe der Telefonnummer
//
// 3. Verschiedene Testfälle
//    - Name gefunden (Anfang, Mitte, Ende)
//    - Name nicht gefunden
//    - Grenzfälle
//
// 4. Visualisierung des Suchprozesses
//    - Zeigt welche Elemente verglichen werden
//    - Zählt Anzahl der Vergleiche
//
// 5. Vergleich mit linearer Suche
//    - Effizienz-Unterschied demonstrieren
//
// Ausführen mit: go run S230_binaere_suche_demo.go
//
// =============================================================================

const arrlen = 10 // Länge des Arrays

// person: Verbund-Typ für Telefonbuch-Eintrag
type person struct {
	name          string
	telefonnummer int
}

// Globale Variable für Visualisierung
var vergleichsZähler int

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Binäre Suche in Verbund-Arrays ===\n")

	// -------------------------------------------------------------------------
	// 1. Telefonbuch erstellen (sortiert nach Name)
	// -------------------------------------------------------------------------

	fmt.Println("1. Telefonbuch erstellen:")
	fmt.Println("   -----------------------")
	fmt.Println()

	// Sortiertes Array von Personen (alphabetisch nach Name)
	telbuch := [arrlen]person{
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

	fmt.Println("   Telefonbuch (sortiert nach Name):")
	for i, p := range telbuch {
		fmt.Printf("   [%d] %-12s -> %d\n", i, p.name, p.telefonnummer)
	}
	fmt.Println()

	fmt.Println("   WICHTIG: Array muss sortiert sein für binäre Suche!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Grundlegende Funktionsweise
	// -------------------------------------------------------------------------

	fmt.Println("2. Grundlegende Funktionsweise der binären Suche:")
	fmt.Println("   -----------------------------------------------")
	fmt.Println()

	fmt.Println("   Prinzip:")
	fmt.Println("   1. Betrachte mittleres Element")
	fmt.Println("   2. Vergleiche gesuchten Name mit mittlerem Element")
	fmt.Println("   3. Wenn gefunden: Fertig!")
	fmt.Println("   4. Wenn alphabetisch früher: Suche in linker Hälfte")
	fmt.Println("   5. Wenn alphabetisch später: Suche in rechter Hälfte")
	fmt.Println("   6. Wiederhole ab Schritt 1 (rekursiv)")
	fmt.Println()

	fmt.Println("   Effizienz:")
	fmt.Println("   ✓ Array wird in jedem Schritt halbiert")
	fmt.Printf("   ✓ Max. Schritte: log₂(%d) ≈ %.0f\n", arrlen, log2(float64(arrlen)))
	fmt.Println("   ✓ Viel schneller als lineare Suche bei großen Arrays")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Beispiel 1: Name in der Mitte
	// -------------------------------------------------------------------------

	fmt.Println("3. Beispiel 1: Suche nach \"Müller\" (etwa in der Mitte):")
	fmt.Println("   ------------------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Müller\")")
	fmt.Println()

	nummer := rbinsVisualisiert(&telbuch, 0, arrlen-1, "Müller", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Beispiel 2: Name am Anfang
	// -------------------------------------------------------------------------

	fmt.Println("4. Beispiel 2: Suche nach \"Bauer\" (am Anfang):")
	fmt.Println("   ---------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Bauer\")")
	fmt.Println()

	nummer = rbinsVisualisiert(&telbuch, 0, arrlen-1, "Bauer", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Beispiel 3: Name am Ende
	// -------------------------------------------------------------------------

	fmt.Println("5. Beispiel 3: Suche nach \"Wolf\" (am Ende):")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Wolf\")")
	fmt.Println()

	nummer = rbinsVisualisiert(&telbuch, 0, arrlen-1, "Wolf", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Beispiel 4: Name nicht vorhanden (alphabetisch kleiner)
	// -------------------------------------------------------------------------

	fmt.Println("6. Beispiel 4: Suche nach \"Abel\" (nicht vorhanden, kleiner als alle):")
	fmt.Println("   --------------------------------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Abel\")")
	fmt.Println()

	nummer = rbinsVisualisiert(&telbuch, 0, arrlen-1, "Abel", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN (alphabetisch vor allen Namen)")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Beispiel 5: Name nicht vorhanden (alphabetisch größer)
	// -------------------------------------------------------------------------

	fmt.Println("7. Beispiel 5: Suche nach \"Zimmermann\" (nicht vorhanden, größer als alle):")
	fmt.Println("   --------------------------------------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Zimmermann\")")
	fmt.Println()

	nummer = rbinsVisualisiert(&telbuch, 0, arrlen-1, "Zimmermann", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN (alphabetisch nach allen Namen)")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Beispiel 6: Name nicht vorhanden (zwischen zwei Namen)
	// -------------------------------------------------------------------------

	fmt.Println("8. Beispiel 6: Suche nach \"Mayer\" (nicht vorhanden, zwischen Meier und Müller):")
	fmt.Println("   ------------------------------------------------------------------------------")
	fmt.Println()

	vergleichsZähler = 0
	fmt.Println("   Aufruf: rbins(&telbuch, 0, 9, \"Mayer\")")
	fmt.Println()

	nummer = rbinsVisualisiert(&telbuch, 0, arrlen-1, "Mayer", 1)

	if nummer != -1 {
		fmt.Printf("   ✓ GEFUNDEN! Telefonnummer: %d\n", nummer)
	} else {
		fmt.Println("   ✗ NICHT GEFUNDEN (liegt alphabetisch zwischen existierenden Namen)")
	}
	fmt.Printf("   Anzahl Vergleiche: %d\n", vergleichsZähler)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Vergleich: Binäre Suche vs. Lineare Suche
	// -------------------------------------------------------------------------

	fmt.Println("9. Effizienz-Vergleich: Binäre vs. Lineare Suche:")
	fmt.Println("   -----------------------------------------------")
	fmt.Println()

	// Alle Namen durchsuchen
	fmt.Println("   Suche nach allen Namen im Telefonbuch:")
	fmt.Println()

	totalBinär := 0
	totalLinear := 0

	for _, p := range telbuch {
		// Binäre Suche
		vergleichsZähler = 0
		rbins(&telbuch, 0, arrlen-1, p.name)
		binärVergleiche := vergleichsZähler

		// Lineare Suche
		linearVergleiche := linSuche(&telbuch, p.name)

		fmt.Printf("   %-12s: Binär: %d Vergleiche, Linear: %d Vergleiche\n",
			p.name, binärVergleiche, linearVergleiche)

		totalBinär += binärVergleiche
		totalLinear += linearVergleiche
	}

	fmt.Println()
	fmt.Printf("   GESAMT:\n")
	fmt.Printf("   Binäre Suche:  %d Vergleiche\n", totalBinär)
	fmt.Printf("   Lineare Suche: %d Vergleiche\n", totalLinear)
	fmt.Printf("   Faktor: %.2fx schneller\n", float64(totalLinear)/float64(totalBinär))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Größeres Beispiel: 1000 Einträge
	// -------------------------------------------------------------------------

	fmt.Println("10. Skalierung: Großes Telefonbuch mit 1000 Einträgen:")
	fmt.Println("    ---------------------------------------------------")
	fmt.Println()

	n := 1000
	maxBinär := int(log2(float64(n))) + 1
	durchschnittLinear := (n + 1) / 2

	fmt.Printf("   Array-Größe: %d Einträge\n", n)
	fmt.Println()
	fmt.Printf("   Binäre Suche:\n")
	fmt.Printf("   - Max. Vergleiche: log₂(%d) ≈ %d\n", n, maxBinär)
	fmt.Println()
	fmt.Printf("   Lineare Suche:\n")
	fmt.Printf("   - Durchschn. Vergleiche: %d\n", durchschnittLinear)
	fmt.Printf("   - Max. Vergleiche: %d\n", n)
	fmt.Println()
	fmt.Printf("   Faktor (Durchschnitt): %.1fx schneller\n", float64(durchschnittLinear)/float64(maxBinär))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 11. Code-Erklärung
	// -------------------------------------------------------------------------

	fmt.Println("11. Funktionsweise des Algorithmus:")
	fmt.Println("    ---------------------------------")
	fmt.Println()

	fmt.Println("   func rbins(personArr *[arrlen]person, links, rechts int, suchwert string) int {")
	fmt.Println("       // 1. Prüfe ob Suchbereich zulässig")
	fmt.Println("       if rechts >= links {")
	fmt.Println()
	fmt.Println("           // 2. Berechne Mitte")
	fmt.Println("           mitte := links + (rechts-links)/2")
	fmt.Println()
	fmt.Println("           // 3. Vergleiche mit mittlerem Element")
	fmt.Println("           if personArr[mitte].name == suchwert {")
	fmt.Println("               return personArr[mitte].telefonnummer  // Gefunden!")
	fmt.Println("           }")
	fmt.Println()
	fmt.Println("           // 4. Suche rekursiv in linker oder rechter Hälfte")
	fmt.Println("           if personArr[mitte].name > suchwert {")
	fmt.Println("               return rbins(personArr, links, mitte-1, suchwert)  // Links")
	fmt.Println("           } else {")
	fmt.Println("               return rbins(personArr, mitte+1, rechts, suchwert)  // Rechts")
	fmt.Println("           }")
	fmt.Println("       }")
	fmt.Println()
	fmt.Println("       return -1  // Nicht gefunden")
	fmt.Println("   }")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 12. Zusammenfassung
	// -------------------------------------------------------------------------

	fmt.Println("12. Zusammenfassung:")
	fmt.Println("    ---------------")
	fmt.Println()

	fmt.Println("   Binäre Suche:")
	fmt.Println("   ✓ Rekursiver Divide-and-Conquer Algorithmus")
	fmt.Println("   ✓ Halbiert Suchbereich in jedem Schritt")
	fmt.Println("   ✓ Laufzeit: O(log₂ n)")
	fmt.Println("   ✓ Sehr effizient bei großen Arrays")
	fmt.Println()

	fmt.Println("   Voraussetzungen:")
	fmt.Println("   ✓ Array muss sortiert sein")
	fmt.Println("   ✓ Direkter Zugriff auf Elemente (Random Access)")
	fmt.Println()

	fmt.Println("   Anwendungen:")
	fmt.Println("   ✓ Telefonbücher, Wörterbücher")
	fmt.Println("   ✓ Datenbanken (Indizes)")
	fmt.Println("   ✓ Alle sortierten Datenstrukturen")
	fmt.Println()

	fmt.Println("   Vergleich mit linearer Suche:")
	fmt.Println("   ✓ Linear: O(n) - durchschnittlich n/2 Vergleiche")
	fmt.Println("   ✓ Binär: O(log₂ n) - maximal log₂(n) Vergleiche")
	fmt.Printf("   ✓ Bei 1000 Einträgen: ~10 vs. ~500 Vergleiche\n")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Funktionen für binäre Suche
// =============================================================================

// rbins: Rekursive binäre Suche (Original-Version aus dem Text)
//
// Sucht im Verbund-Array nach einem Element, dessen Feld name den Wert
// suchwert hat und gibt bei Erfolg den Wert des Felds telefonnummer
// dieses Elements zurück; ansonsten den Wert -1, wenn suchwert nicht
// gefunden wurde.
//
// Das Array muss aufsteigend sortiert sein und suchwert darf im Array
// höchstens einmal vorkommen.
//
// links und rechts sind die Grenzen des Suchbereichs, der beide Grenzen
// mit umfasst.
//
// Der initiale Aufruf erfolgt mit dem kompletten Array, also links=0 und
// rechts=arrlen-1
func rbins(personArr *[arrlen]person, links, rechts int, suchwert string) int {
	var ergebnis int
	vergleichsZähler++ // Für Statistik

	// prüfen, ob Suchbereich zulässig ist
	if rechts >= links {
		// Mitte des Suchbereichs berechnen
		// (Index-Position im gesamten Array)
		mitte := links + (rechts-links)/2

		if personArr[mitte].name == suchwert {
			// suchwert gefunden
			ergebnis = personArr[mitte].telefonnummer
		} else {
			if personArr[mitte].name > suchwert {
				// Suche im linken Teilbereich fortsetzen
				ergebnis = rbins(personArr, links, mitte-1, suchwert)
			} else {
				// Suche im rechten Teilbereich fortsetzen
				ergebnis = rbins(personArr, mitte+1, rechts, suchwert)
			}
		}
	} else {
		ergebnis = -1 // suchwert nicht gefunden
	}
	return ergebnis
}

// rbinsVisualisiert: Binäre Suche mit Visualisierung
func rbinsVisualisiert(personArr *[arrlen]person, links, rechts int, suchwert string, tiefe int) int {
	var ergebnis int
	vergleichsZähler++

	// Einrückung für Visualisierung
	indent := ""
	for i := 0; i < tiefe-1; i++ {
		indent += "    "
	}

	// Zeige Suchbereich
	if rechts >= links {
		mitte := links + (rechts-links)/2

		fmt.Printf("%sSchritt %d: Bereich [%d..%d], Mitte=%d (%s)\n",
			indent, tiefe, links, rechts, mitte, personArr[mitte].name)

		if personArr[mitte].name == suchwert {
			fmt.Printf("%s✓ GEFUNDEN bei Index %d: %s -> %d\n",
				indent, mitte, personArr[mitte].name, personArr[mitte].telefonnummer)
			ergebnis = personArr[mitte].telefonnummer
		} else {
			if personArr[mitte].name > suchwert {
				fmt.Printf("%s   \"%s\" < \"%s\" -> Suche links weiter\n",
					indent, suchwert, personArr[mitte].name)
				ergebnis = rbinsVisualisiert(personArr, links, mitte-1, suchwert, tiefe+1)
			} else {
				fmt.Printf("%s   \"%s\" > \"%s\" -> Suche rechts weiter\n",
					indent, suchwert, personArr[mitte].name)
				ergebnis = rbinsVisualisiert(personArr, mitte+1, rechts, suchwert, tiefe+1)
			}
		}
	} else {
		fmt.Printf("%sSchritt %d: Bereich ungültig [%d..%d] -> NICHT GEFUNDEN\n",
			indent, tiefe, links, rechts)
		ergebnis = -1
	}
	return ergebnis
}

// linSuche: Lineare Suche zum Vergleich
func linSuche(personArr *[arrlen]person, suchwert string) int {
	vergleiche := 0
	for i := 0; i < arrlen; i++ {
		vergleiche++
		if personArr[i].name == suchwert {
			return vergleiche
		}
	}
	return vergleiche
}

// log2: Berechnet log₂(n)
func log2(n float64) float64 {
	if n <= 1 {
		return 0
	}
	count := 0.0
	for n > 1 {
		n /= 2
		count++
	}
	return count
}
