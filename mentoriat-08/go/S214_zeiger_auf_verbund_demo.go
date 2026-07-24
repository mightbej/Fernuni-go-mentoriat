package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Zeiger auf Verbünde und Verbund-Literale (Abschnitt 8.2.4)
// =============================================================================
//
// Themen:
// 1. Zeiger-Variable eines Verbund-Typs (*person)
// 2. Kurzschreibweise für Feld-Zugriffe: pPer.nachname statt (*pPer).nachname
// 3. Adressierbarkeit von Verbund-Literalen: &person{...}
// 4. Kompakte vs. explizite Schreibweise
// 5. Verbund-Literale als Funktionsargumente
// 6. Jedes Literal erzeugt eine neue Variable (new wird implizit aufgerufen)
//
// Ausführen mit: go run S214_zeiger_auf_verbund_demo.go
//
// =============================================================================

// Definition der Verbund-Datentypen
type person struct {
	vorname     string
	nachname    string
	gebJahr     int
	ort         string
	verheiratet bool
}

type rgb struct {
	rot  int
	grün int
	blau int
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Zeiger auf Verbünde und Verbund-Literale ===\n")

	// -------------------------------------------------------------------------
	// 1. Zeiger-Variable eines Verbund-Typs
	// -------------------------------------------------------------------------

	fmt.Println("1. Zeiger-Variable eines Verbund-Typs:")
	fmt.Println("   ------------------------------------")

	var anna = person{"Anna", "Schmidt", 1966, "Hamburg", false}
	fmt.Printf("   anna = %+v\n", anna)
	fmt.Println()

	// Deklaration einer Zeiger-Variable auf den Verbund-Typ person
	var pPer *person = &anna
	fmt.Println("   var pPer *person = &anna")
	fmt.Printf("   pPer zeigt auf Adresse: %p\n", pPer)
	fmt.Printf("   *pPer (dereferenziert):  %+v\n", *pPer)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Feld-Zugriff via Verbund-Zeiger: Zwei Schreibweisen
	// -------------------------------------------------------------------------

	fmt.Println("2. Feld-Zugriff via Verbund-Zeiger:")
	fmt.Println("   ---------------------------------")
	fmt.Println()

	// Explizite Syntax: Dereferenzierung mit Klammerung
	fmt.Println("   a) Explizite Syntax (mit Dereferenzierung und Klammerung):")
	fmt.Printf("      Vorher: (*pPer).nachname = %s\n", (*pPer).nachname)

	(*pPer).nachname = "Meier"
	
	fmt.Printf("      Nachher: (*pPer).nachname = %s\n", (*pPer).nachname)
	fmt.Printf("      anna.nachname = %s (wurde geändert!)\n", anna.nachname)
	fmt.Println()

	// Kompakte Syntax: implizite Dereferenzierung
	fmt.Println("   b) Kompakte Syntax (implizite Dereferenzierung):")
	fmt.Printf("      Vorher: pPer.nachname = %s\n", pPer.nachname)

	pPer.nachname = "Müller"
	
	fmt.Printf("      Nachher: pPer.nachname = %s\n", pPer.nachname)
	fmt.Printf("      anna.nachname = %s (wurde geändert!)\n", anna.nachname)
	fmt.Println()

	fmt.Println("   MERKE: Beide Schreibweisen sind äquivalent!")
	fmt.Println("   ✓ (*pPer).nachname  <->  pPer.nachname")
	fmt.Println("   ✓ Go führt die Dereferenzierung implizit durch")
	fmt.Println("   ✓ Die kompakte Syntax ist in der Praxis üblich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Adressierbarkeit von Verbund-Literalen
	// -------------------------------------------------------------------------

	fmt.Println("3. Adressierbarkeit von Verbund-Literalen:")
	fmt.Println("   ----------------------------------------")
	fmt.Println()

	// Kompakte Schreibweise: Direktes Adressieren eines Literals
	fmt.Println("   Kompakte Schreibweise:")
	fmt.Println("   var pFarbe *rgb = &rgb{128, 0, 80}")

	var pFarbe *rgb = &rgb{128, 0, 80}
	
	fmt.Printf("   pFarbe zeigt auf: %p\n", pFarbe)
	fmt.Printf("   *pFarbe = %+v\n", *pFarbe)
	fmt.Println()

	// Was passiert hinter den Kulissen?
	fmt.Println("   Was geschieht hinter den Kulissen:")
	fmt.Println("   1. new(rgb) erzeugt anonyme Variable")
	fmt.Println("   2. Adresse wird in pFarbe gespeichert")
	fmt.Println("   3. *pFarbe wird das Literal zugewiesen")
	fmt.Println()

	// Explizite Schreibweise (äquivalent)
	fmt.Println("   Äquivalente explizite Schreibweise:")
	fmt.Println("   var pFarbe2 *rgb = new(rgb)")
	fmt.Println("   *pFarbe2 = rgb{128, 0, 80}")

	var pFarbe2 *rgb = new(rgb)
	*pFarbe2 = rgb{128, 0, 80}
	
	fmt.Printf("   pFarbe2 zeigt auf: %p\n", pFarbe2)
	fmt.Printf("   *pFarbe2 = %+v\n", *pFarbe2)
	fmt.Println()

	fmt.Println("   MERKE: Beide Varianten erzeugen das gleiche Ergebnis!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Verschiedene Initialisierungsformen
	// -------------------------------------------------------------------------

	fmt.Println("4. Verschiedene Initialisierungsformen für Verbund-Zeiger:")
	fmt.Println("   --------------------------------------------------------")
	fmt.Println()

	// Form 1: Zeiger auf existierende Variable
	var max = person{"Max", "Meier", 1990, "Berlin", true}
	var pMax1 *person = &max
	fmt.Println("   Form 1: Zeiger auf existierende Variable")
	fmt.Println("   var max = person{\"Max\", \"Meier\", 1990, \"Berlin\", true}")
	fmt.Println("   var pMax1 *person = &max")
	fmt.Printf("   -> %+v\n", *pMax1)
	fmt.Println()

	// Form 2: new() mit nachträglicher Zuweisung
	var pMax2 *person = new(person)
	*pMax2 = person{"Max", "Meier", 1990, "Berlin", true}
	fmt.Println("   Form 2: new() mit nachträglicher Zuweisung")
	fmt.Println("   var pMax2 *person = new(person)")
	fmt.Println("   *pMax2 = person{\"Max\", \"Meier\", 1990, \"Berlin\", true}")
	fmt.Printf("   -> %+v\n", *pMax2)
	fmt.Println()

	// Form 3: Kompakte Schreibweise (Adresse des Literals)
	var pMax3 *person = &person{"Max", "Meier", 1990, "Berlin", true}
	fmt.Println("   Form 3: Kompakte Schreibweise (empfohlen!)")
	fmt.Println("   var pMax3 *person = &person{\"Max\", \"Meier\", 1990, \"Berlin\", true}")
	fmt.Printf("   -> %+v\n", *pMax3)
	fmt.Println()

	// Form 4: Mit Feldnamen (empfohlen für Klarheit)
	var pMax4 *person = &person{
		vorname:     "Max",
		nachname:    "Meier",
		gebJahr:     1990,
		ort:         "Berlin",
		verheiratet: true,
	}
	fmt.Println("   Form 4: Mit Feldnamen (beste Praxis!)")
	fmt.Println("   var pMax4 *person = &person{")
	fmt.Println("       vorname: \"Max\",")
	fmt.Println("       nachname: \"Meier\",")
	fmt.Println("       ...")
	fmt.Println("   }")
	fmt.Printf("   -> %+v\n", *pMax4)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Verbund-Literale als Funktionsargumente
	// -------------------------------------------------------------------------

	fmt.Println("5. Verbund-Literale als Funktionsargumente:")
	fmt.Println("   -----------------------------------------")
	fmt.Println()

	// Kompakte Schreibweise: Literal direkt als Argument
	fmt.Println("   Kompakte Schreibweise (Einzeiler):")
	fmt.Println("   fügeHinzu(&rgb{0, 255, 10})")
	fügeHinzu(&rgb{0, 255, 10})
	fmt.Println()

	// Explizite Schreibweise: Hilfsvariable erforderlich
	fmt.Println("   Explizite Schreibweise (Hilfsvariable):")
	fmt.Println("   h := new(rgb)")
	fmt.Println("   *h = rgb{0, 255, 10}")
	fmt.Println("   fügeHinzu(h)")
	h := new(rgb)
	*h = rgb{0, 255, 10}
	fügeHinzu(h)
	fmt.Println()

	fmt.Println("   VORTEIL der kompakten Schreibweise:")
	fmt.Println("   ✓ Kein temporäre Hilfsvariable nötig")
	fmt.Println("   ✓ Lesbarer und kompakter Code")
	fmt.Println("   ✓ Sehr häufig in Go-Code verwendet")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Jedes Literal erzeugt eine neue Variable
	// -------------------------------------------------------------------------

	fmt.Println("6. Jedes Literal erzeugt eine neue Variable:")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	fmt.Println("   Zwei Literale mit identischem Wert:")
	f1 := &rgb{22, 33, 44}
	f2 := &rgb{22, 33, 44}

	fmt.Println("   f1 := &rgb{22, 33, 44}  // neue Variable")
	fmt.Println("   f2 := &rgb{22, 33, 44}  // andere neue Variable")
	fmt.Println()

	fmt.Printf("   Adresse von f1: %p\n", f1)
	fmt.Printf("   Adresse von f2: %p\n", f2)
	fmt.Printf("   Wert von *f1:   %+v\n", *f1)
	fmt.Printf("   Wert von *f2:   %+v\n", *f2)
	fmt.Println()

	fmt.Printf("   f1 == f2? %t  (Zeiger zeigen auf unterschiedliche Variablen!)\n", f1 == f2)
	fmt.Printf("   *f1 == *f2? %t  (Die Werte sind identisch!)\n", *f1 == *f2)
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ Jedes &rgb{...} erzeugt eine NEUE anonyme Variable")
	fmt.Println("   ✓ Zeiger vergleichen Adressen, nicht Werte")
	fmt.Println("   ✓ f1 und f2 zeigen auf unterschiedliche Speicherstellen")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Praktisches Beispiel: Funktion mit mehreren Aufrufen
	// -------------------------------------------------------------------------

	fmt.Println("7. Praktisches Beispiel - Funktionsaufrufe mit Literalen:")
	fmt.Println("   -------------------------------------------------------")
	fmt.Println()

	fmt.Println("   Mehrere Funktionsaufrufe mit Verbund-Literalen:")
	erzeugePersonen()
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Vergleich: Array-Literale vs. Verbund-Literale
	// -------------------------------------------------------------------------

	fmt.Println("8. Vergleich: Array-Literale und Verbund-Literale:")
	fmt.Println("   ------------------------------------------------")
	fmt.Println()

	// Array-Literal ist ebenfalls adressierbar
	var pArray *[4]int = &[4]int{7, 3, 5, 8}
	fmt.Println("   Array-Literale sind auch adressierbar:")
	fmt.Println("   var pArray *[4]int = &[4]int{7, 3, 5, 8}")
	fmt.Printf("   -> %v\n", *pArray)
	fmt.Println()

	// Verbund-Literal
	var pVerbund *rgb = &rgb{100, 150, 200}
	fmt.Println("   Verbund-Literale sind adressierbar:")
	fmt.Println("   var pVerbund *rgb = &rgb{100, 150, 200}")
	fmt.Printf("   -> %+v\n", *pVerbund)
	fmt.Println()

	
	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Zusammengesetzte Literale (Array, Verbund) sind adressierbar")
	fmt.Println("   ✓ Einfache Literale (int, string, etc.) sind NICHT adressierbar")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Häufiger Fehler: Versuch, einfache Literale zu adressieren
	// -------------------------------------------------------------------------

	fmt.Println("9. Häufiger Fehler beim Programmieren:")
	fmt.Println("   ------------------------------------")
	fmt.Println()

	fmt.Println("   FALSCH: var pZahl *int = &42  // Compilerfehler!")
	fmt.Println("           Einfache Literale sind NICHT adressierbar")
	fmt.Println()

	fmt.Println("   RICHTIG: var zahl int = 42")
	fmt.Println("            var pZahl *int = &zahl  // Variable ist adressierbar")
	var zahl int = 42
	var pZahl *int = &zahl
	fmt.Printf("            -> %v\n", *pZahl)
	fmt.Println()

	fmt.Println("   AUSNAHME: Zusammengesetzte Literale sind adressierbar:")
	fmt.Println("            var pRGB *rgb = &rgb{1, 2, 3}  // OK!")
	fmt.Println("            var pArr *[3]int = &[3]int{1, 2, 3}  // OK!")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// fügeHinzu demonstriert eine Funktion, die einen Zeiger auf rgb erwartet
func fügeHinzu(pFarbe *rgb) {
	fmt.Printf("   -> fügeHinzu() aufgerufen mit: %+v (Adresse: %p)\n", *pFarbe, pFarbe)
	// In der Praxis würde hier die Farbe zu einer Sammlung hinzugefügt
}

// erzeugePersonen demonstriert mehrfache Funktionsaufrufe mit Literalen
func erzeugePersonen() {
	fmt.Println("   Aufruf: verarbeite(&person{\"Lisa\", \"Weber\", 2000, \"Dresden\", false})")
	verarbeite(&person{"Lisa", "Weber", 2000, "Dresden", false})

	fmt.Println("   Aufruf: verarbeite(&person{\"Tom\", \"Klein\", 1995, \"Leipzig\", true})")
	verarbeite(&person{"Tom", "Klein", 1995, "Leipzig", true})

	fmt.Println("   Aufruf: verarbeite(&person{\"Sarah\", \"Groß\", 1988, \"Hamburg\", false})")
	verarbeite(&person{"Sarah", "Groß", 1988, "Hamburg", false})
}

// verarbeite demonstriert eine Funktion, die einen Zeiger auf person erwartet
func verarbeite(p *person) {
	fmt.Printf("   -> Verarbeite: %s %s aus %s (Adresse: %p)\n",
		p.vorname, p.nachname, p.ort, p)
}
