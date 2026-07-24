package main

import "fmt"

// =============================================================================
// 1. Grundlegende struct-Definition (zeilenweise - üblich und übersichtlich)
// =============================================================================

type person struct {
	vorname     string
	nachname    string
	gebJahr     int
	ort         string
	verheiratet bool
}

// =============================================================================
// 2. Alternative Schreibweise: Felder desselben Typs zusammen deklarieren
// =============================================================================

type personKompakt struct {
	vorname, nachname string
	gebJahr           int
	ort               string
	verheiratet       bool
}

// =============================================================================
// 3. Andere Reihenfolge = anderer Datentyp!
// =============================================================================

type personAndereReihenfolge struct {
	nachname, vorname, ort string
	verheiratet            bool
	gebJahr                int
}

// =============================================================================
// 4. Exportierte vs. nicht-exportierte Typen und Felder
// =============================================================================

// Person ist exportiert (Großbuchstabe am Anfang)
type Person struct {
	Vorname     string // exportiert
	Nachname    string // exportiert
	GebJahr     int    // exportiert
	Ort         string // exportiert
	Verheiratet bool   // exportiert
}

// personPrivat ist nicht exportiert (Kleinbuchstabe am Anfang)
type personPrivat struct {
	vorname     string // nicht exportiert
	nachname    string // nicht exportiert
	gebJahr     int    // nicht exportiert
	ort         string // nicht exportiert
	verheiratet bool   // nicht exportiert
}

// =============================================================================
// 5. Funktionen, die mit structs arbeiten
// =============================================================================

// jünger gibt die jüngere von zwei Personen zurück
func jünger(p1, p2 person) person {
	if p1.gebJahr > p2.gebJahr {
		return p1
	}
	return p2
}

// personInfo gibt formatierte Informationen über eine Person aus
func personInfo(p person) {
	verhStatus := "ledig"
	if p.verheiratet {
		verhStatus = "verheiratet"
	}
	fmt.Printf("  %s %s, geb. %d, wohnt in %s, %s\n",
		p.vorname, p.nachname, p.gebJahr, p.ort, verhStatus)
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: struct in Go ===\n")

	// -------------------------------------------------------------------------
	// Verschiedene Möglichkeiten, struct-Variablen zu initialisieren
	// -------------------------------------------------------------------------

	fmt.Println("1. Verschiedene Initialisierungsmethoden:")
	fmt.Println("   ----------------------------------------")

	// a) Zero-Value: alle Felder werden mit ihren Zero-Values initialisiert
	var p1 person
	
	fmt.Println("   a) Zero-Value Initialisierung:")
	fmt.Printf("      p1 = %+v\n", p1)
	fmt.Println()

	// b) Struct-Literal mit Feldnamen (empfohlen!)
	p2 := person{
		vorname:     "Anna",
		nachname:    "Schmidt",
		gebJahr:     1995,
		ort:         "Berlin",
		verheiratet: false, 
	}
	fmt.Println("   b) Initialisierung mit Feldnamen:")
	fmt.Printf("      p2 = %+v\n", p2)
	fmt.Println()

	// c) Struct-Literal mit Positionsangaben (nicht empfohlen, fehleranfällig!)
	p3 := person{"Max", "Müller", 1988, "München", true}
	fmt.Println("   c) Initialisierung mit Positionen (nicht empfohlen!):")
	fmt.Printf("      p3 = %+v\n", p3)
	fmt.Println()

	// d) Schrittweise Zuweisung
	var p4 person
	p4.vorname = "Lisa"
	p4.nachname = "Weber"
	p4.gebJahr = 2000
	p4.ort = "Hamburg"
	p4.verheiratet = false
	fmt.Println("   d) Schrittweise Zuweisung:")
	fmt.Printf("      p4 = %+v\n", p4)
	fmt.Println()

	// -------------------------------------------------------------------------
	// Zugriff auf Felder
	// -------------------------------------------------------------------------

	fmt.Println("2. Zugriff auf struct-Felder:")
	fmt.Println("   ---------------------------")
	fmt.Printf("   p2.vorname   = %s\n", p2.vorname)
	fmt.Printf("   p2.nachname  = %s\n", p2.nachname)
	fmt.Printf("   p2.gebJahr   = %d\n", p2.gebJahr)
	fmt.Println()

	// -------------------------------------------------------------------------
	// Verwendung in Arrays
	// -------------------------------------------------------------------------

	fmt.Println("3. structs in Arrays:")
	fmt.Println("   -------------------")
	var großeltern [4]person
	großeltern[0] = person{"Hans", "Schmidt", 1950, "Köln", true}
	großeltern[1] = person{"Helga", "Schmidt", 1952, "Köln", true}
	großeltern[2] = person{"Fritz", "Müller", 1948, "Dresden", true}
	großeltern[3] = person{"Grete", "Müller", 1951, "Dresden", true}

	fmt.Println("   Großeltern-Array:")
	for i, gp := range großeltern {
		fmt.Printf("   [%d] ", i)
		personInfo(gp)
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// Verwendung in Funktionen
	// -------------------------------------------------------------------------

	fmt.Println("4. structs als Funktionsparameter:")
	fmt.Println("   ---------------------------------")
	fmt.Print("   Person 2: ")
	personInfo(p2)
	fmt.Print("   Person 3: ")
	personInfo(p3)

	jüngere := jünger(p2, p3)
	fmt.Print("   Jüngere Person: ")
	personInfo(jüngere)
	fmt.Println()

	// -------------------------------------------------------------------------
	// Demonstration: Unterschiedliche struct-Typen sind nicht kompatibel
	// -------------------------------------------------------------------------

	fmt.Println("5. Verschiedene struct-Typen sind NICHT kompatibel:")
	fmt.Println("   --------------------------------------------------")

	p5 := person{vorname: "Tom", nachname: "Klein", gebJahr: 1990}
	p6 := personKompakt{vorname: "Sarah", nachname: "Groß", gebJahr: 1992}

	// WICHTIG: p5 und p6 sind von verschiedenen Typen!
	// Folgende Zeile würde NICHT kompilieren:
	// p5 = p6  // FEHLER: cannot use p6 (type personKompakt) as type person

	fmt.Println("   p5 (Typ: person):")
	fmt.Printf("      %+v\n", p5)
	fmt.Println("   p6 (Typ: personKompakt):")
	fmt.Printf("      %+v\n", p6)
	fmt.Println("   -> Obwohl die Felder identisch sind, sind p5 und p6")
	fmt.Println("      verschiedene Typen und können nicht zugewiesen werden!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// Demonstration: Reihenfolge der Felder ist wichtig
	// -------------------------------------------------------------------------

	fmt.Println("6. Reihenfolge der Felder bestimmt den Typ:")
	fmt.Println("   -----------------------------------------")

	p7 := personAndereReihenfolge{
		nachname:    "Lang",
		vorname:     "Julia",
		ort:         "Leipzig",
		verheiratet: true,
		gebJahr:     1985,
	}

	fmt.Println("   p7 (Typ: personAndereReihenfolge):")
	fmt.Printf("      %+v\n", p7)
	fmt.Println("   -> Andere Feldreihenfolge = anderer Typ!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// Demonstration: Anonyme structs (ohne benannten Typ)
	// -------------------------------------------------------------------------

	fmt.Println("7. Anonyme structs (ohne benannten Typ):")
	fmt.Println("   ---------------------------------------")

	// Inline-Definition eines struct ohne type-Deklaration
	anonymPerson := struct {
		vorname  string
		nachname string
		alter    int
	}{
		vorname:  "Klaus",
		nachname: "Bauer",
		alter:    45,
	}

	fmt.Printf("   anonymPerson = %+v\n", anonymPerson)
	fmt.Println("   -> Nützlich für temporäre Datenstrukturen")
	fmt.Println()

	// -------------------------------------------------------------------------
	// Syntaktische Varianten (meist nicht empfohlen)
	// -------------------------------------------------------------------------

	fmt.Println("8. Syntaktische Varianten (unüblich, aber zulässig):")
	fmt.Println("   --------------------------------------------------")

	// Ein-Zeilen-Definition mit Semikolons (unübersichtlich!)
	type einZeiler struct {
		v string
		n string
		j int
	}
	ez := einZeiler{v: "Test", n: "Person", j: 2020}
	fmt.Printf("   einZeiler = %+v\n", ez)
	fmt.Println("   -> Syntaktisch korrekt, aber unübersichtlich!")
	fmt.Println("   -> In der Praxis: IMMER zeilenweise Schreibweise verwenden!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// Exportierte vs. nicht-exportierte Namen
	// -------------------------------------------------------------------------

	fmt.Println("9. Exportierte vs. nicht-exportierte Namen:")
	fmt.Println("   -----------------------------------------")

	// Kleingeschrieben = nicht exportiert (nur im eigenen Paket sichtbar)
	klein := personPrivat{
		vorname:  "Peter",
		nachname: "Privat",
		gebJahr:  1975,
	}

	// Großgeschrieben = exportiert (auch außerhalb des Pakets sichtbar)
	groß := Person{
		Vorname:  "Paula",
		Nachname: "Public",
		GebJahr:  1980,
	}

	fmt.Printf("   personPrivat (nicht exportiert): %+v\n", klein)
	fmt.Printf("   Person (exportiert):              %+v\n", groß)
	fmt.Println("   -> In Bibliotheken: Typnamen und Felder meist groß")
	fmt.Println("   -> In kleinen Programmen: oft klein (bewusst nicht exportiert)")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}
