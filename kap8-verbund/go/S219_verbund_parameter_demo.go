package main

import (
	"fmt"
	"time"
)

// =============================================================================
// DEMONSTRATION: Verbund-Parameter - Call-by-Value vs. Call-by-Reference
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Funktionen mit Zeiger-Parameter (*auto)
//    - Änderungen betreffen das Original (Call-by-Reference)
//    - Speicher- und zeiteffizient
//
// 2. Funktionen mit Wert-Parameter (auto)
//    - Änderungen betreffen nur die Kopie (Call-by-Value)
//    - Original bleibt unverändert (Sicherheit)
//
// 3. Performance-Vergleich
//    - Zeiger-Übergabe: schnell, wenig Speicher
//    - Wert-Übergabe: langsamer bei großen Verbünden
//
// 4. Wann welche Variante wählen?
//    - Zeiger: wenn Änderung gewünscht ODER Performance wichtig
//    - Wert: wenn Unveränderbarkeit garantiert werden soll
//
// Ausführen mit: go run S219_verbund_parameter_demo.go
//

// Besondere Highlights:

// Performance-Test mit 100.000 Iterationen zeigt realen Unterschied
// Beispiel mit Programmierfehler zeigt Sicherheitsvorteil von Call-by-Value
// Klare Empfehlungen für beide Varianten
// =============================================================================

// Definition des Verbund-Datentyps auto
type auto struct {
	marke    string
	modell   string
	baujahr  int
	tüvJahr  int
	kmStand  int
	farbe    string
}

// Definition eines großen Verbunds für Performance-Tests
type großerVerbund struct {
	daten [1000]int // Array mit 1000 Integern
	info  string
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Verbund-Parameter - Call-by-Value vs. Call-by-Reference ===\n")

	// -------------------------------------------------------------------------
	// 1. Funktion mit Zeiger-Parameter (Call-by-Reference)
	// -------------------------------------------------------------------------

	fmt.Println("1. Funktion mit Zeiger-Parameter (Call-by-Reference):")
	fmt.Println("   ---------------------------------------------------")
	fmt.Println()

	golf := auto{
		marke:   "VW",
		modell:  "Golf",
		baujahr: 2020,
		tüvJahr: 2024,
		kmStand: 45000,
		farbe:   "blau",
	}

	fmt.Printf("   golf vor Funktionsaufruf: %+v\n", golf)
	fmt.Printf("   TÜV vorher: %d\n", golf.tüvJahr)
	fmt.Println()

	fmt.Println("   Aufruf: tüvErneuernZeiger(&golf)")
	tüvErneuernZeiger(&golf)

	fmt.Printf("   TÜV nachher: %d (geändert!)\n", golf.tüvJahr)
	fmt.Printf("   golf nach Funktionsaufruf: %+v\n", golf)
	fmt.Println()

	fmt.Println("   ERGEBNIS:")
	fmt.Println("   ✓ Zeiger-Parameter ermöglicht Änderung am Original")
	fmt.Println("   ✓ golf.tüvJahr wurde von 2024 auf 2026 geändert")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Funktion mit Wert-Parameter (Call-by-Value)
	// -------------------------------------------------------------------------

	fmt.Println("2. Funktion mit Wert-Parameter (Call-by-Value):")
	fmt.Println("   ----------------------------------------------")
	fmt.Println()

	audi := auto{
		marke:   "Audi",
		modell:  "A4",
		baujahr: 2019,
		tüvJahr: 2024,
		kmStand: 52000,
		farbe:   "schwarz",
	}

	fmt.Printf("   audi vor Funktionsaufruf: %+v\n", audi)
	fmt.Printf("   TÜV vorher: %d\n", audi.tüvJahr)
	fmt.Println()

	fmt.Println("   Aufruf: tüvErneuernWert(audi)")
	tüvErneuernWert(audi)

	fmt.Printf("   TÜV nachher: %d (unverändert!)\n", audi.tüvJahr)
	fmt.Printf("   audi nach Funktionsaufruf: %+v\n", audi)
	fmt.Println()

	fmt.Println("   ERGEBNIS:")
	fmt.Println("   ✓ Wert-Parameter: nur Kopie wird geändert")
	fmt.Println("   ✓ audi.tüvJahr bleibt bei 2024 (Original unverändert)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Direkter Vergleich: Zeiger vs. Wert
	// -------------------------------------------------------------------------

	fmt.Println("3. Direkter Vergleich:")
	fmt.Println("   -------------------")
	fmt.Println()

	bmw := auto{
		marke:   "BMW",
		modell:  "320i",
		baujahr: 2021,
		tüvJahr: 2025,
		kmStand: 28000,
		farbe:   "weiß",
	}

	fmt.Printf("   Original: %+v\n", bmw)
	fmt.Println()

	// Test mit Wert-Parameter
	fmt.Println("   Test 1: tüvErneuernWert(bmw)")
	fmt.Printf("   Vor Aufruf:  TÜV = %d\n", bmw.tüvJahr)
	tüvErneuernWert(bmw)
	fmt.Printf("   Nach Aufruf: TÜV = %d (unverändert)\n", bmw.tüvJahr)
	fmt.Println()

	// Test mit Zeiger-Parameter
	fmt.Println("   Test 2: tüvErneuernZeiger(&bmw)")
	fmt.Printf("   Vor Aufruf:  TÜV = %d\n", bmw.tüvJahr)
	tüvErneuernZeiger(&bmw)
	fmt.Printf("   Nach Aufruf: TÜV = %d (geändert!)\n", bmw.tüvJahr)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Lesende Funktionen - Performance-Überlegungen
	// -------------------------------------------------------------------------

	fmt.Println("4. Lesende Funktionen - Performance:")
	fmt.Println("   ----------------------------------")
	fmt.Println()

	mercedes := auto{
		marke:   "Mercedes",
		modell:  "C-Klasse",
		baujahr: 2022,
		tüvJahr: 2026,
		kmStand: 15000,
		farbe:   "silber",
	}

	fmt.Println("   Funktion, die nur liest (keine Änderungen):")
	fmt.Println()

	// Mit Zeiger-Parameter (empfohlen)
	fmt.Println("   a) Mit Zeiger-Parameter (effizient):")
	istTüvFälligZeiger(&mercedes)
	fmt.Println()

	// Mit Wert-Parameter (ineffizient, aber sicher)
	fmt.Println("   b) Mit Wert-Parameter (Kopie wird erstellt):")
	istTüvFälligWert(mercedes)
	fmt.Println()

	fmt.Println("   EMPFEHLUNG:")
	fmt.Println("   ✓ Auch bei lesenden Funktionen Zeiger verwenden")
	fmt.Println("   ✓ Spart Speicher und Rechenzeit (keine Kopie)")
	fmt.Println("   ✓ Besonders wichtig bei großen Verbünden")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Performance-Messung: Großer Verbund
	// -------------------------------------------------------------------------

	fmt.Println("5. Performance-Messung mit großem Verbund:")
	fmt.Println("   ----------------------------------------")
	fmt.Println()

	großerDatensatz := großerVerbund{
		info: "Testdaten",
	}
	for i := 0; i < 1000; i++ {
		großerDatensatz.daten[i] = i
	}

	fmt.Println("   Großer Verbund: 1000 Integers + String")
	fmt.Printf("   Größe: ca. %d Bytes\n", 1000*8+len(großerDatensatz.info))
	fmt.Println()

	// Zeitmessung: Mit Zeiger
	const iterationen = 100000
	start := time.Now()
	for i := 0; i < iterationen; i++ {
		verarbeiteGroßZeiger(&großerDatensatz)
	}
	zeitZeiger := time.Since(start)

	// Zeitmessung: Mit Wert (Kopie)
	start = time.Now()
	for i := 0; i < iterationen; i++ {
		verarbeiteGroßWert(großerDatensatz)
	}
	zeitWert := time.Since(start)

	fmt.Printf("   %d Funktionsaufrufe:\n", iterationen)
	fmt.Printf("   - Mit Zeiger: %v\n", zeitZeiger)
	fmt.Printf("   - Mit Wert:   %v\n", zeitWert)
	fmt.Printf("   - Faktor:     %.2fx langsamer bei Wert-Übergabe\n",
		float64(zeitWert)/float64(zeitZeiger))
	fmt.Println()

	fmt.Println("   FAZIT:")
	fmt.Println("   ✓ Zeiger-Übergabe ist deutlich schneller")
	fmt.Println("   ✓ Bei großen Verbünden enormer Unterschied")
	fmt.Println("   ✓ Wert-Übergabe erstellt vollständige Kopie")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Sicherheitsaspekt: Unveränderbarkeit garantieren
	// -------------------------------------------------------------------------

	fmt.Println("6. Sicherheitsaspekt - Unveränderbarkeit:")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	tesla := auto{
		marke:   "Tesla",
		modell:  "Model 3",
		baujahr: 2023,
		tüvJahr: 2027,
		kmStand: 8000,
		farbe:   "rot",
	}

	fmt.Printf("   Original: %+v\n", tesla)
	fmt.Println()

	fmt.Println("   Szenario: Funktion mit Programmierfehler")
	fmt.Println()

	// Funktion mit Zeiger: Fehler wirkt sich aus
	fmt.Println("   a) Mit Zeiger-Parameter:")
	fmt.Printf("      Vor Aufruf: kmStand = %d\n", tesla.kmStand)
	berechneDurchschnittZeiger(&tesla) // Enthält versehentliche Modifikation!
	fmt.Printf("      Nach Aufruf: kmStand = %d (versehentlich geändert!)\n", tesla.kmStand)
	fmt.Println()

	// Zurücksetzen
	tesla.kmStand = 8000

	// Funktion mit Wert: Fehler wirkt sich NICHT aus
	fmt.Println("   b) Mit Wert-Parameter:")
	fmt.Printf("      Vor Aufruf: kmStand = %d\n", tesla.kmStand)
	berechneDurchschnittWert(tesla) // Modifikation betrifft nur Kopie
	fmt.Printf("      Nach Aufruf: kmStand = %d (unverändert trotz Fehler!)\n", tesla.kmStand)
	fmt.Println()

	fmt.Println("   VORTEIL Wert-Übergabe:")
	fmt.Println("   ✓ Original wird garantiert nicht verändert")
	fmt.Println("   ✓ Schutz vor versehentlichen Modifikationen")
	fmt.Println("   ✓ Programmierfehler bleiben ohne Auswirkung")
	fmt.Println()

	fmt.Println("   NACHTEIL Wert-Übergabe:")
	fmt.Println("   ✓ Fehler bleibt länger unentdeckt")
	fmt.Println("   ✓ Performance-Overhead bei großen Verbünden")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Entscheidungshilfe: Wann was verwenden?
	// -------------------------------------------------------------------------

	fmt.Println("7. Entscheidungshilfe:")
	fmt.Println("   -------------------")
	fmt.Println()

	fmt.Println("   Zeiger-Parameter verwenden (*auto):")
	fmt.Println("   ✓ Wenn Funktion den Verbund ändern soll")
	fmt.Println("   ✓ Bei großen Verbünden (Performance)")
	fmt.Println("   ✓ Bei vielen/häufigen Funktionsaufrufen")
	fmt.Println("   ✓ Wenn Speicher gespart werden soll")
	fmt.Println()

	fmt.Println("   Wert-Parameter verwenden (auto):")
	fmt.Println("   ✓ Wenn Unveränderbarkeit garantiert sein muss")
	fmt.Println("   ✓ Bei kleinen Verbünden (wenige Felder)")
	fmt.Println("   ✓ Wenn Sicherheit wichtiger als Performance")
	fmt.Println("   ✓ Bei seltenen Funktionsaufrufen")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Praktische Beispiele
	// -------------------------------------------------------------------------

	fmt.Println("8. Praktische Beispiele:")
	fmt.Println("   ----------------------")
	fmt.Println()

	ford := auto{
		marke:   "Ford",
		modell:  "Focus",
		baujahr: 2018,
		tüvJahr: 2024,
		kmStand: 75000,
		farbe:   "grün",
	}

	fmt.Printf("   Ausgangszustand: %+v\n", ford)
	fmt.Println()

	// Beispiel 1: Änderung erwünscht (Zeiger)
	fmt.Println("   Beispiel 1: Kilometerstand aktualisieren (Zeiger)")
	aktualisiereFahrt(&ford, 350)
	fmt.Printf("   -> kmStand = %d\n", ford.kmStand)
	fmt.Println()

	// Beispiel 2: Nur Anzeige (Zeiger für Performance)
	fmt.Println("   Beispiel 2: Informationen anzeigen (Zeiger)")
	zeigeAutoInfo(&ford)
	fmt.Println()

	// Beispiel 3: Berechnung ohne Änderung (Wert für Sicherheit)
	fmt.Println("   Beispiel 3: Wert berechnen ohne Änderung (Wert)")
	wert := berechneWert(ford)
	fmt.Printf("   -> Geschätzter Wert: %d EUR\n", wert)
	fmt.Printf("   -> Original unverändert: %+v\n", ford)
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Funktionen mit Zeiger-Parameter (Call-by-Reference)
// =============================================================================

// tüvErneuernZeiger verlängert den TÜV um 2 Jahre (ändert Original)
func tüvErneuernZeiger(a *auto) {
	fmt.Printf("   -> In Funktion: Original-Adresse %p\n", a)
	a.tüvJahr += 2
	fmt.Printf("   -> In Funktion: tüvJahr auf %d gesetzt\n", a.tüvJahr)
}

// istTüvFälligZeiger prüft, ob TÜV fällig ist (nur lesend, mit Zeiger)
func istTüvFälligZeiger(a *auto) {
	aktuellesJahr := 2025
	fmt.Printf("   -> istTüvFälligZeiger: TÜV %d, Aktuell %d", a.tüvJahr, aktuellesJahr)
	if a.tüvJahr < aktuellesJahr {
		fmt.Println(" -> TÜV abgelaufen!")
	} else {
		fmt.Println(" -> TÜV gültig")
	}
}

// aktualisiereFahrt aktualisiert den Kilometerstand nach einer Fahrt
func aktualisiereFahrt(a *auto, gefahreneKm int) {
	a.kmStand += gefahreneKm
	fmt.Printf("   -> %d km gefahren, neuer Stand: %d km\n", gefahreneKm, a.kmStand)
}

// zeigeAutoInfo gibt Informationen über das Auto aus
func zeigeAutoInfo(a *auto) {
	fmt.Printf("   -> %s %s (%d), %s, %d km, TÜV bis %d\n",
		a.marke, a.modell, a.baujahr, a.farbe, a.kmStand, a.tüvJahr)
}

// berechneDurchschnittZeiger berechnet Durchschnitt (mit versehentlicher Modifikation!)
func berechneDurchschnittZeiger(a *auto) float64 {
	jahre := 2025 - a.baujahr
	if jahre == 0 {
		jahre = 1
	}
	durchschnitt := float64(a.kmStand) / float64(jahre)
	
	// FEHLER: Versehentliche Modifikation!
	a.kmStand = 99999 // Programmierfehler - sollte nicht passieren!
	
	return durchschnitt
}

// verarbeiteGroßZeiger verarbeitet großen Verbund (nur für Performance-Test)
func verarbeiteGroßZeiger(g *großerVerbund) int {
	// Simuliert lesenden Zugriff
	return g.daten[500]
}

// =============================================================================
// Funktionen mit Wert-Parameter (Call-by-Value)
// =============================================================================

// tüvErneuernWert versucht TÜV zu verlängern (ändert nur Kopie!)
func tüvErneuernWert(a auto) {
	fmt.Printf("   -> In Funktion: Kopie-Adresse %p\n", &a)
	a.tüvJahr += 2
	fmt.Printf("   -> In Funktion: tüvJahr der Kopie auf %d gesetzt\n", a.tüvJahr)
	fmt.Println("   -> Achtung: Nur die Kopie wurde geändert!")
}

// istTüvFälligWert prüft, ob TÜV fällig ist (nur lesend, mit Wert-Kopie)
func istTüvFälligWert(a auto) {
	aktuellesJahr := 2025
	fmt.Printf("   -> istTüvFälligWert: TÜV %d, Aktuell %d", a.tüvJahr, aktuellesJahr)
	if a.tüvJahr < aktuellesJahr {
		fmt.Println(" -> TÜV abgelaufen!")
	} else {
		fmt.Println(" -> TÜV gültig")
	}
	fmt.Println("   -> Kopie wurde erstellt (Speicher-/Zeit-Overhead)")
}

// berechneWert berechnet Autowert ohne Original zu ändern
func berechneWert(a auto) int {
	basiswert := 30000
	alterAbzug := (2025 - a.baujahr) * 2000
	kmAbzug := a.kmStand / 10
	return basiswert - alterAbzug - kmAbzug
}

// berechneDurchschnittWert berechnet Durchschnitt (mit versehentlicher Modifikation der Kopie)
func berechneDurchschnittWert(a auto) float64 {
	jahre := 2025 - a.baujahr
	if jahre == 0 {
		jahre = 1
	}
	durchschnitt := float64(a.kmStand) / float64(jahre)
	
	// FEHLER: Versehentliche Modifikation (betrifft nur Kopie!)
	a.kmStand = 99999 // Programmierfehler - aber keine Auswirkung aufs Original!
	
	return durchschnitt
}

// verarbeiteGroßWert verarbeitet großen Verbund (nur für Performance-Test)
func verarbeiteGroßWert(g großerVerbund) int {
	// Simuliert lesenden Zugriff (Kopie wird erstellt!)
	return g.daten[500]
}
