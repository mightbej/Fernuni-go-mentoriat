package main

import "fmt"

// Das Programm demonstriert:

// ✅ Verbund-Variablen und Punkt-Notation
// ✅ Zugriff auf Feld-Variablen
// ✅ Änderung von Feld-Werten
// ✅ Feld-Variablen sind adressierbar (Zeiger)
// ✅ WICHTIG: Unterschied zwischen Kopie und Zeiger
// ✅ Auswirkungen bei Änderung des Verbunds
// ✅ Indirekte Änderung über Zeiger
// ✅ Keine Auswirkung bei Kopie-Variablen
// ✅ Visualisierung der Speichersituation
// ✅ Praktisches Beispiel mit Funktionen

// go run S212_verbund_felder_demo.go

// Besonders betont wird der zentrale Unterschied:

// var heimat string = anna.ort → Kopie des Werts
// var pHeimat *string = &anna.ort → Zeiger auf das Feld

// =============================================================================
// Definition des Verbund-Datentyps person
// =============================================================================

type person struct {
	vorname     string
	nachname    string
	gebJahr     int
	ort         string
	verheiratet bool
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Verbund-Variablen und Feld-Variablen ===\n")

	// -------------------------------------------------------------------------
	// 1. Verbund-Variable und Zugriff mit Punkt-Notation
	// -------------------------------------------------------------------------

	fmt.Println("1. Verbund-Variable und Punkt-Notation:")
	fmt.Println("   -------------------------------------")

	// Deklaration und Initialisierung einer Verbund-Variablen
	var anna = person{"Anna", "Schmidt", 1966, "Hamburg", false}
	fmt.Printf("   Initial:        %+v\n", anna)
	fmt.Println()

	// Zugriff auf Felder mit Punkt-Notation
	fmt.Println("   Zugriff auf einzelne Felder:")
	fmt.Printf("   anna.vorname     = %s\n", anna.vorname)
	fmt.Printf("   anna.nachname    = %s\n", anna.nachname)
	fmt.Printf("   anna.gebJahr     = %d\n", anna.gebJahr)
	fmt.Printf("   anna.ort         = %s\n", anna.ort)
	fmt.Printf("   anna.verheiratet = %t\n", anna.verheiratet)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Änderung von Feld-Werten
	// -------------------------------------------------------------------------

	fmt.Println("2. Änderung von Feld-Werten:")
	fmt.Println("   --------------------------")

	// Feld-Variable nachname ändern (String-Verkettung)
	fmt.Println("   Vor Änderung:")
	fmt.Printf("   anna.nachname = %s, anna.verheiratet = %t\n", anna.nachname, anna.verheiratet)

	anna.nachname += "-Schiller"
	anna.verheiratet = true

	fmt.Println("   Nach Änderung:")
	fmt.Printf("   anna.nachname = %s, anna.verheiratet = %t\n", anna.nachname, anna.verheiratet)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Feld-Variablen sind adressierbar
	// -------------------------------------------------------------------------

	fmt.Println("3. Feld-Variablen sind adressierbar:")
	fmt.Println("   -----------------------------------")

	// Zeiger auf die gesamte Verbund-Variable
	var pAnna *person = &anna

	fmt.Printf("   Zeiger auf Verbund: pAnna = %p\n", pAnna)
	fmt.Printf("   Wert über Zeiger:   *pAnna = %+v\n", *pAnna)
	fmt.Println()

	// Zeiger auf einzelne Feld-Variablen
	var pVorname *string = &anna.vorname
	var pNachname *string = &anna.nachname
	var pGebJahr *int = &anna.gebJahr
	var pOrt *string = &anna.ort
	var pVerheiratet *bool = &anna.verheiratet

	fmt.Println("   Zeiger auf einzelne Felder:")
	fmt.Printf("   &anna.vorname     = %p -> Wert: %s\n", pVorname, *pVorname)
	fmt.Printf("   &anna.nachname    = %p -> Wert: %s\n", pNachname, *pNachname)
	fmt.Printf("   &anna.gebJahr     = %p -> Wert: %d\n", pGebJahr, *pGebJahr)
	fmt.Printf("   &anna.ort         = %p -> Wert: %s\n", pOrt, *pOrt)
	fmt.Printf("   &anna.verheiratet = %p -> Wert: %t\n", pVerheiratet, *pVerheiratet)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. WICHTIG: Feld-Wert vs. Zeiger auf Feld
	// -------------------------------------------------------------------------

	fmt.Println("4. WICHTIG: Feld-Wert (Kopie) vs. Zeiger auf Feld:")
	fmt.Println("   ------------------------------------------------")
	fmt.Println()

	// Aktuelle Situation
	fmt.Printf("   Aktueller Verbund: %+v\n", anna)
	fmt.Println()

	// Variable mit KOPIE des Feld-Werts
	var heimat string = anna.ort
	fmt.Println("   var heimat string = anna.ort")
	fmt.Printf("   -> heimat enthält KOPIE: %s\n", heimat)
	fmt.Println()

	// Variable mit ZEIGER auf das Feld
	var pHeimat *string = &anna.ort
	fmt.Println("   var pHeimat *string = &anna.ort")
	fmt.Printf("   -> pHeimat zeigt auf Feld: %p\n", pHeimat)
	fmt.Printf("   -> *pHeimat dereferenziert: %s\n", *pHeimat)
	fmt.Println()

	// Vergleich der Werte
	fmt.Printf("   heimat == *pHeimat? %t\n", heimat == *pHeimat)
	fmt.Printf("   (beide sind momentan: %s)\n", heimat)
	fmt.Println()

	fmt.Println("   " + strings("=", 70))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Änderung des Feld-Werts im Verbund
	// -------------------------------------------------------------------------

	fmt.Println("5. Änderung des Feld-Werts im Verbund:")
	fmt.Println("   -------------------------------------")

	fmt.Println("   Aktion: anna.ort = \"München\"")
	anna.ort = "München"
	fmt.Println()

	fmt.Println("   Auswirkungen:")
	fmt.Printf("   - heimat     = %s (unverändert! Ist nur eine Kopie)\n", heimat)
	fmt.Printf("   - *pHeimat   = %s (geändert! Zeiger zeigt auf das Feld)\n", *pHeimat)
	fmt.Printf("   - anna.ort   = %s (geändert!)\n", anna.ort)
	fmt.Println()

	fmt.Printf("   heimat == *pHeimat? %t (jetzt unterschiedlich!)\n", heimat == *pHeimat)
	fmt.Printf("   Verbund: %+v\n", anna)
	fmt.Println()

	fmt.Println("   " + strings("=", 70))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Indirekte Änderung über Zeiger
	// -------------------------------------------------------------------------

	fmt.Println("6. Indirekte Änderung des Verbunds über Zeiger:")
	fmt.Println("   ----------------------------------------------")

	fmt.Println("   Aktion: *pHeimat = \"Ulm\"")
	*pHeimat = "Ulm"
	fmt.Println()

	fmt.Println("   Auswirkungen:")
	fmt.Printf("   - *pHeimat zeigt auf anna.ort -> Feld wurde geändert!\n")
	fmt.Printf("   - anna.ort   = %s (geändert durch Zeiger!)\n", anna.ort)
	fmt.Printf("   - heimat     = %s (unverändert, ist Kopie)\n", heimat)
	fmt.Printf("   Verbund: %+v\n", anna)
	fmt.Println()

	fmt.Println("   " + strings("=", 70))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Zuweisung an die Kopie hat keine Auswirkung auf den Verbund
	// -------------------------------------------------------------------------

	fmt.Println("7. Zuweisung an Kopie-Variable:")
	fmt.Println("   ------------------------------")

	fmt.Println("   Aktion: heimat = \"Kiel\"")
	heimat = "Kiel"
	fmt.Println()

	fmt.Println("   Auswirkungen:")
	fmt.Printf("   - heimat     = %s (geändert)\n", heimat)
	fmt.Printf("   - anna.ort   = %s (UNVERÄNDERT! heimat ist nur Kopie)\n", anna.ort)
	fmt.Printf("   - *pHeimat   = %s (unverändert, zeigt weiter auf anna.ort)\n", *pHeimat)
	fmt.Printf("   Verbund: %+v\n", anna)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Zusammenfassung mit Visualisierung
	// -------------------------------------------------------------------------

	fmt.Println("8. Zusammenfassung - Kopie vs. Zeiger:")
	fmt.Println("   ------------------------------------")
	fmt.Println()
	fmt.Println("   Speichersituation (vereinfacht):")
	fmt.Println()
	fmt.Println("   ┌─────────────────────────────────────────────┐")
	fmt.Println("   │ anna (Verbund-Variable)                     │")
	fmt.Println("   ├─────────────────────────────────────────────┤")
	fmt.Println("   │ vorname:     \"Anna\"                          │")
	fmt.Println("   │ nachname:    \"Schmidt-Schiller\"              │")
	fmt.Println("   │ gebJahr:     1966                           │")
	fmt.Printf("   │ ort:         %-31s │ <--- pHeimat zeigt hierher\n", "\""+anna.ort+"\"")
	fmt.Println("   │ verheiratet: true                           │")
	fmt.Println("   └─────────────────────────────────────────────┘")
	fmt.Println()
	fmt.Printf("   heimat = \"%s\"  (eigenständige Variable, Kopie)\n", heimat)
	fmt.Println()
	fmt.Println("   MERKE:")
	fmt.Println("   ✓ heimat ist eine eigenständige Kopie")
	fmt.Println("   ✓ Änderungen an heimat betreffen nur heimat selbst")
	fmt.Println("   ✓ pHeimat zeigt auf das Feld anna.ort")
	fmt.Println("   ✓ Änderungen über *pHeimat ändern anna.ort direkt")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Weitere Beispiele: Zeiger auf Verbund vs. Zeiger auf Feld
	// -------------------------------------------------------------------------

	fmt.Println("9. Bonus: Zeiger auf Verbund vs. Zeiger auf Feld:")
	fmt.Println("   ------------------------------------------------")

	var max = person{"Max", "Müller", 1990, "Berlin", false}
	fmt.Printf("   Original: %+v\n", max)
	fmt.Println()

	// Zeiger auf den gesamten Verbund
	var pMax *person = &max
	fmt.Println("   Über Zeiger auf Verbund:")
	fmt.Println("   pMax.ort = \"Köln\"  (Kurzschreibweise für (*pMax).ort)")
	pMax.ort = "Köln"
	fmt.Printf("   -> max = %+v\n", max)
	fmt.Println()

	// Zeiger auf einzelnes Feld
	var pMaxOrt *string = &max.ort
	fmt.Println("   Über Zeiger auf Feld:")
	fmt.Println("   *pMaxOrt = \"Frankfurt\"")
	*pMaxOrt = "Frankfurt"
	fmt.Printf("   -> max = %+v\n", max)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Praktisches Beispiel: Funktion mit Zeiger-Parameter
	// -------------------------------------------------------------------------

	fmt.Println("10. Praktisches Beispiel: Funktion mit Zeigern:")
	fmt.Println("    ---------------------------------------------")

	var lisa = person{"Lisa", "Weber", 2000, "Dresden", false}
	fmt.Printf("    Vor umzug():  %+v\n", lisa)

	umzug(&lisa, "Leipzig")
	fmt.Printf("    Nach umzug(): %+v\n", lisa)
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// strings wiederholt einen String n-mal
func strings(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// umzug ändert den Ort einer Person über einen Zeiger
func umzug(p *person, neuerOrt string) {
	fmt.Printf("    umzug() aufgerufen: %s zieht um nach %s\n", p.vorname, neuerOrt)
	p.ort = neuerOrt
}
