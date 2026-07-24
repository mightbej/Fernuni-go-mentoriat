package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Hierarchische Verbünde (Abschnitt 8.3)
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Hierarchische Verbund-Typen
//    - Verbünde können andere Verbünde als Felder enthalten
//    - Ermöglicht komplexe, geschachtelte Datenstrukturen
//
// 2. Felder mit Verbund-Werten (besitzerW)
//    - Feld speichert kompletten Verbund-Wert (Kopie)
//    - Unflexibel: Jeder Besitzer hat sein eigenes Auto
//    - Änderungen am Original wirken sich nicht auf Kopie aus
//
// 3. Felder mit Verbund-Zeigern (besitzerZ)
//    - Feld speichert nur Referenz auf Verbund
//    - Flexibel: Mehrere Besitzer können sich ein Auto teilen
//    - Änderungen wirken sich auf alle Referenzen aus
//
// 4. Praktische Vorteile von Zeiger-Feldern
//    - Sharing von Daten zwischen mehreren Strukturen
//    - Speichereffizienz
//    - In der Praxis meist bevorzugt
//
// Ausführen mit: go run S220_hierarchische_verbunde_demo.go
//
// =============================================================================

// Definition des Verbund-Datentyps auto
type auto struct {
	marke       string
	kennzeichen string
	tachostand  int
	tüvJahr     int
}

// besitzerW: Hierarchischer Verbund mit Wert-Feld (unflexibel)
// Das Feld "fahrzeug" speichert einen kompletten auto-Verbund (Kopie)
type besitzerW struct {
	name     string
	ort      string
	fahrzeug auto // Feld mit Verbund-Typ (Wert, nicht Zeiger!)
}

// besitzerZ: Hierarchischer Verbund mit Zeiger-Feld (flexibel)
// Das Feld "fahrzeug" speichert nur eine Referenz auf einen auto-Verbund
type besitzerZ struct {
	name     string
	ort      string
	fahrzeug *auto // Feld mit Zeiger auf Verbund-Typ
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Hierarchische Verbünde ===\n")

	// -------------------------------------------------------------------------
	// 1. Hierarchischer Verbund mit Wert-Feld (besitzerW)
	// -------------------------------------------------------------------------

	fmt.Println("1. Hierarchischer Verbund mit Wert-Feld:")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	// Initialisierung eines hierarchischen Verbunds
	mike := besitzerW{
		name: "Mike",
		ort:  "München",
		fahrzeug: auto{
			marke:       "Mercedes",
			kennzeichen: "M-RT357",
			tachostand:  133000,
			tüvJahr:     2024,
		},
	}

	fmt.Printf("   mike = %+v\n", mike)
	fmt.Printf("   Mikes Auto: %+v\n", mike.fahrzeug)
	fmt.Println()

	fmt.Println("   Struktur:")
	fmt.Println("   besitzerW enthält auto als WERT (kompletter Verbund)")
	fmt.Println("   -> Jeder besitzerW hat sein eigenes Auto (Kopie)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Wert vs. Zeiger auf Feld
	// -------------------------------------------------------------------------

	fmt.Println("2. Wert (Kopie) vs. Zeiger auf Feld:")
	fmt.Println("   -----------------------------------")
	fmt.Println()

	fmt.Println("   Ausgangssituation:")
	fmt.Printf("   mike.fahrzeug = %+v\n", mike.fahrzeug)
	fmt.Println()

	// Variable mit Kopie des Feld-Werts
	var altesAuto auto = mike.fahrzeug
	fmt.Println("   var altesAuto auto = mike.fahrzeug")
	fmt.Printf("   -> altesAuto (Kopie): %+v\n", altesAuto)
	fmt.Println()

	// Zeiger auf das Feld
	var pMikesAuto *auto = &(mike.fahrzeug)
	fmt.Println("   var pMikesAuto *auto = &(mike.fahrzeug)")
	fmt.Printf("   -> pMikesAuto zeigt auf: %p\n", pMikesAuto)
	fmt.Printf("   -> *pMikesAuto (Zeiger): %+v\n", *pMikesAuto)
	fmt.Println()

	fmt.Println("   Vergleich:")
	fmt.Printf("   altesAuto:     %+v\n", altesAuto)
	fmt.Printf("   *pMikesAuto:   %+v\n", *pMikesAuto)
	fmt.Printf("   mike.fahrzeug: %+v\n", mike.fahrzeug)
	fmt.Println("   -> Alle haben momentan den gleichen Wert")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Feld einen neuen Wert zuweisen
	// -------------------------------------------------------------------------

	fmt.Println("3. Feld einen neuen Wert zuweisen:")
	fmt.Println("   ---------------------------------")
	fmt.Println()

	bmw := auto{
		marke:       "BMW",
		kennzeichen: "M-TS300",
		tachostand:  0,
		tüvJahr:     2028,
	}

	fmt.Printf("   Neues Auto erstellt: bmw = %+v\n", bmw)
	fmt.Println()

	fmt.Println("   Zuweisung: mike.fahrzeug = bmw")
	mike.fahrzeug = bmw
	fmt.Println()

	fmt.Println("   Auswirkungen:")
	fmt.Printf("   altesAuto:     %+v (unverändert! Ist Kopie)\n", altesAuto)
	fmt.Printf("   *pMikesAuto:   %+v (geändert! Zeigt auf Feld)\n", *pMikesAuto)
	fmt.Printf("   mike.fahrzeug: %+v (geändert!)\n", mike.fahrzeug)
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ altesAuto ist eigenständige Variable (Kopie)")
	fmt.Println("   ✓ Änderung am Feld betrifft altesAuto nicht")
	fmt.Println("   ✓ pMikesAuto zeigt auf Feld -> liefert neuen Wert")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Änderung an bmw wirkt sich NICHT auf mike.fahrzeug aus
	// -------------------------------------------------------------------------

	fmt.Println("4. Feldwert ist Kopie, keine Referenz:")
	fmt.Println("   -------------------------------------")
	fmt.Println()

	fmt.Printf("   Vor Änderung:\n")
	fmt.Printf("   bmw.tachostand:           %d\n", bmw.tachostand)
	fmt.Printf("   mike.fahrzeug.tachostand: %d\n", mike.fahrzeug.tachostand)
	fmt.Println()

	fmt.Println("   Änderung: bmw.tachostand = 8000")
	bmw.tachostand = 8000
	fmt.Println()

	fmt.Printf("   Nach Änderung:\n")
	fmt.Printf("   bmw.tachostand:           %d (geändert)\n", bmw.tachostand)
	fmt.Printf("   mike.fahrzeug.tachostand: %d (unverändert!)\n", mike.fahrzeug.tachostand)
	fmt.Println()

	fmt.Println("   ERKLÄRUNG:")
	fmt.Println("   ✓ mike.fahrzeug = bmw kopierte den Wert von bmw")
	fmt.Println("   ✓ mike.fahrzeug und bmw sind unabhängige Verbünde")
	fmt.Println("   ✓ Änderungen an bmw betreffen mike.fahrzeug nicht")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Problem: Auto-Sharing ist nicht möglich
	// -------------------------------------------------------------------------

	fmt.Println("5. Problem mit Wert-Feldern:")
	fmt.Println("   ---------------------------")
	fmt.Println()

	fmt.Println("   Szenario: Zwei Personen sollen sich ein Auto teilen")
	fmt.Println()

	vw := auto{
		marke:       "VW",
		kennzeichen: "B-AB123",
		tachostand:  50000,
		tüvJahr:     2025,
	}

	anna := besitzerW{
		name:     "Anna",
		ort:      "Berlin",
		fahrzeug: vw, // Kopie von vw
	}

	bert := besitzerW{
		name:     "Bert",
		ort:      "Berlin",
		fahrzeug: vw, // weitere Kopie von vw
	}

	fmt.Println("   anna := besitzerW{..., fahrzeug: vw}")
	fmt.Println("   bert := besitzerW{..., fahrzeug: vw}")
	fmt.Println()

	fmt.Printf("   anna.fahrzeug: %+v (Adresse: %p)\n", anna.fahrzeug, &anna.fahrzeug)
	fmt.Printf("   bert.fahrzeug: %+v (Adresse: %p)\n", bert.fahrzeug, &bert.fahrzeug)
	fmt.Println()

	fmt.Println("   -> Jeder hat seine EIGENE Kopie des Autos!")
	fmt.Println("   -> Kein Sharing möglich!")
	fmt.Println()

	fmt.Println("   Test: Anna fährt 1000 km")
	anna.fahrzeug.tachostand += 1000
	fmt.Printf("   anna.fahrzeug.tachostand: %d (geändert)\n", anna.fahrzeug.tachostand)
	fmt.Printf("   bert.fahrzeug.tachostand: %d (unverändert!)\n", bert.fahrzeug.tachostand)
	fmt.Println()

	fmt.Println("   PROBLEM:")
	fmt.Println("   ✓ Mit besitzerW (Wert-Feld) kein Auto-Sharing möglich")
	fmt.Println("   ✓ Jeder Besitzer hat sein eigenes, unabhängiges Auto")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Lösung: Hierarchischer Verbund mit Zeiger-Feld (besitzerZ)
	// -------------------------------------------------------------------------

	fmt.Println("6. Lösung: Hierarchischer Verbund mit Zeiger-Feld:")
	fmt.Println("   ------------------------------------------------")
	fmt.Println()

	audi := auto{
		marke:       "Audi",
		kennzeichen: "IN-CV45",
		tachostand:  0,
		tüvJahr:     2026,
	}

	fmt.Printf("   audi = %+v (Adresse: %p)\n", audi, &audi)
	fmt.Println()

	olaf := besitzerZ{
		name:     "Olaf",
		ort:      "Hannover",
		fahrzeug: &audi, // Zeiger auf audi
	}

	olga := besitzerZ{
		name:     "Olga",
		ort:      "Hannover",
		fahrzeug: &audi, // Zeiger auf audi (dasselbe Auto!)
	}

	fmt.Println("   olaf := besitzerZ{..., fahrzeug: &audi}")
	fmt.Println("   olga := besitzerZ{..., fahrzeug: &audi}")
	fmt.Println()

	fmt.Printf("   olaf = %+v\n", olaf)
	fmt.Printf("   olga = %+v\n", olga)
	fmt.Println()

	fmt.Printf("   olaf.fahrzeug zeigt auf: %p\n", olaf.fahrzeug)
	fmt.Printf("   olga.fahrzeug zeigt auf: %p\n", olga.fahrzeug)
	fmt.Printf("   &audi:                   %p\n", &audi)
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ Beide Zeiger verweisen auf DASSELBE Auto!")
	fmt.Println("   ✓ Auto-Sharing ist möglich!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Änderungen wirken sich auf alle Referenzen aus
	// -------------------------------------------------------------------------

	fmt.Println("7. Änderungen wirken sich auf alle Referenzen aus:")
	fmt.Println("   ------------------------------------------------")
	fmt.Println()

	fmt.Println("   Ausgangssituation:")
	fmt.Printf("   audi:            %+v\n", audi)
	fmt.Printf("   *olaf.fahrzeug:  %+v\n", *olaf.fahrzeug)
	fmt.Printf("   *olga.fahrzeug:  %+v\n", *olga.fahrzeug)
	fmt.Println()

	// Änderung 1: Direkter Zugriff auf audi
	fmt.Println("   Änderung 1: audi.kennzeichen = \"H-OL22\"")

	audi.kennzeichen = "H-OL22"
	
	fmt.Printf("   -> audi:           %+v\n", audi)
	fmt.Printf("   -> *olaf.fahrzeug: %+v\n", *olaf.fahrzeug)
	fmt.Printf("   -> *olga.fahrzeug: %+v\n", *olga.fahrzeug)
	fmt.Println()

	// Änderung 2: Indirekter Zugriff über olaf
	fmt.Println("   Änderung 2: olaf.fahrzeug.tachostand += 1000")

	olaf.fahrzeug.tachostand += 1000
	
	fmt.Printf("   -> audi:           %+v\n", audi)
	fmt.Printf("   -> *olaf.fahrzeug: %+v\n", *olaf.fahrzeug)
	fmt.Printf("   -> *olga.fahrzeug: %+v\n", *olga.fahrzeug)
	fmt.Println()

	// Änderung 3: Indirekter Zugriff über olga
	fmt.Println("   Änderung 3: olga.fahrzeug.tachostand += 500")
	olga.fahrzeug.tachostand += 500
	fmt.Printf("   -> audi:           %+v\n", audi)
	fmt.Printf("   -> *olaf.fahrzeug: %+v\n", *olaf.fahrzeug)
	fmt.Printf("   -> *olga.fahrzeug: %+v\n", *olga.fahrzeug)
	fmt.Println()

	fmt.Println("   FAZIT:")
	fmt.Println("   ✓ Alle Änderungen betreffen das GLEICHE Auto")
	fmt.Println("   ✓ Egal ob direkt (audi) oder indirekt (olaf/olga)")
	fmt.Println("   ✓ tachostand: 0 -> 1000 -> 1500")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Direkter Vergleich: besitzerW vs. besitzerZ
	// -------------------------------------------------------------------------

	fmt.Println("8. Direkter Vergleich: besitzerW vs. besitzerZ:")
	fmt.Println("   ----------------------------------------------")
	fmt.Println()

	// besitzerW: Wert-Feld
	fmt.Println("   A) besitzerW mit Wert-Feld:")
	fmt.Println("   ---------------------------")

	ford := auto{
		marke:       "Ford",
		kennzeichen: "F-123",
		tachostand:  20000,
		tüvJahr:     2025,
	}

	// Beide bekommen eine KOPIE
	max1 := besitzerW{name: "Max", ort: "Frankfurt", fahrzeug: ford}
	lisa1 := besitzerW{name: "Lisa", ort: "Frankfurt", fahrzeug: ford}

	fmt.Printf("   max1.fahrzeug:  %+v (Adresse: %p)\n", max1.fahrzeug, &max1.fahrzeug)
	fmt.Printf("   lisa1.fahrzeug: %+v (Adresse: %p)\n", lisa1.fahrzeug, &lisa1.fahrzeug)
	fmt.Println()

	max1.fahrzeug.tachostand = 30000
	fmt.Println("   Nach max1.fahrzeug.tachostand = 30000:")
	fmt.Printf("   max1.fahrzeug.tachostand:  %d (geändert)\n", max1.fahrzeug.tachostand)
	fmt.Printf("   lisa1.fahrzeug.tachostand: %d (unverändert)\n", lisa1.fahrzeug.tachostand)
	fmt.Println("   -> Kein Sharing, getrennte Autos")
	fmt.Println()

	// besitzerZ: Zeiger-Feld
	fmt.Println("   B) besitzerZ mit Zeiger-Feld:")
	fmt.Println("   -----------------------------")

	opel := auto{
		marke:       "Opel",
		kennzeichen: "OP-456",
		tachostand:  15000,
		tüvJahr:     2026,
	}

	// Beide bekommen einen ZEIGER auf dasselbe Auto
	max2 := besitzerZ{name: "Max", ort: "Frankfurt", fahrzeug: &opel}
	lisa2 := besitzerZ{name: "Lisa", ort: "Frankfurt", fahrzeug: &opel}

	fmt.Printf("   max2.fahrzeug:  %p -> %+v\n", max2.fahrzeug, *max2.fahrzeug)
	fmt.Printf("   lisa2.fahrzeug: %p -> %+v\n", lisa2.fahrzeug, *lisa2.fahrzeug)
	fmt.Println()

	max2.fahrzeug.tachostand = 25000
	fmt.Println("   Nach max2.fahrzeug.tachostand = 25000:")
	fmt.Printf("   max2.fahrzeug.tachostand:  %d (geändert)\n", max2.fahrzeug.tachostand)
	fmt.Printf("   lisa2.fahrzeug.tachostand: %d (auch geändert!)\n", lisa2.fahrzeug.tachostand)
	fmt.Println("   -> Sharing funktioniert, gleiches Auto")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Praktische Vorteile von Zeiger-Feldern
	// -------------------------------------------------------------------------

	fmt.Println("9. Praktische Vorteile von Zeiger-Feldern:")
	fmt.Println("   -----------------------------------------")
	fmt.Println()

	fmt.Println("   Vorteile von Zeiger-Feldern (besitzerZ):")
	fmt.Println("   ✓ Flexibilität: Mehrere Strukturen können dasselbe Objekt teilen")
	fmt.Println("   ✓ Speichereffizienz: Nur eine Kopie des Objekts im Speicher")
	fmt.Println("   ✓ Konsistenz: Änderungen sind überall sichtbar")
	fmt.Println("   ✓ Realistische Modellierung: Entspricht oft der Realität")
	fmt.Println()

	fmt.Println("   Nachteile von Wert-Feldern (besitzerW):")
	fmt.Println("   ✗ Unflexibel: Jede Struktur hat eigene Kopie")
	fmt.Println("   ✗ Speicher-Overhead: Mehrfache Kopien im Speicher")
	fmt.Println("   ✗ Inkonsistenz: Änderungen betreffen nur eine Kopie")
	fmt.Println("   ✗ Kein Sharing: Nicht möglich, Objekte zu teilen")
	fmt.Println()

	fmt.Println("   EMPFEHLUNG:")
	fmt.Println("   ✓ In der Praxis werden hierarchische Verbünde meist")
	fmt.Println("     mit Zeiger-Feldern aufgebaut")
	fmt.Println("   ✓ Zeiger-Felder sind flexibler und speichereffizienter")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Komplexeres Beispiel: Mehrere Ebenen
	// -------------------------------------------------------------------------

	fmt.Println("10. Komplexeres Beispiel: Mehrere Hierarchie-Ebenen:")
	fmt.Println("    --------------------------------------------------")
	fmt.Println()

	// Weitere Verschachtelung
	type familie struct {
		name       string
		mitglieder []*besitzerZ // Array von Zeigern auf besitzerZ
		gemeinsam  *auto        // Gemeinsames Familienauto
	}

	fiat := auto{
		marke:       "Fiat",
		kennzeichen: "F-FAM1",
		tachostand:  80000,
		tüvJahr:     2025,
	}

	vater := besitzerZ{name: "Hans", ort: "Frankfurt", fahrzeug: &fiat}
	mutter := besitzerZ{name: "Maria", ort: "Frankfurt", fahrzeug: &fiat}

	müller := familie{
		name:       "Familie Müller",
		mitglieder: []*besitzerZ{&vater, &mutter},
		gemeinsam:  &fiat,
	}

	fmt.Printf("   %s:\n", müller.name)
	for i, m := range müller.mitglieder {
		fmt.Printf("   Mitglied %d: %s, Auto: %s (%s)\n",
			i+1, m.name, m.fahrzeug.marke, m.fahrzeug.kennzeichen)
	}
	fmt.Printf("   Gemeinsames Auto: %+v\n", *müller.gemeinsam)
	fmt.Println()

	fmt.Println("   -> Alle Familienmitglieder teilen sich ein Auto")
	fmt.Println("   -> Komplexe Hierarchie mit mehreren Ebenen")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}
