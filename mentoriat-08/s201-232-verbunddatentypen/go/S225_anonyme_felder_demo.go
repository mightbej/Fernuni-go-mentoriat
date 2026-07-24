package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Anonyme Felder und eingebettete Verbünde (Abschnitt 8.3.2)
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Variante 1: Flache Verbund-Typen
//    - Alle Felder direkt im Verbund
//    - Einfach, aber unstrukturiert
//
// 2. Variante 2: Hierarchische Verbund-Typen mit benannten Feldern
//    - Bessere Strukturierung durch Aggregation
//    - Umständlicher Zugriff über vollständigen Pfad
//
// 3. Beste Variante: Anonyme Felder (Embedded Structs)
//    - Hierarchische Struktur mit einfachem Zugriff
//    - Anonyme Felder können beim Zugriff weggelassen werden
//    - Implizite Feldnamen basierend auf Typ-Namen
//
// 4. Alternative Zugriffswege und Literale
//
// Ausführen mit: go run S225_anonyme_felder_demo.go
//
// =============================================================================

// =============================================================================
// VARIANTE 1: Flache Verbund-Typen
// =============================================================================

type kreisV1 struct {
	x, y, radius int
}

type radV1 struct {
	x, y, radius, speichen int
}

// =============================================================================
// VARIANTE 2: Hierarchische Verbund-Typen mit benannten Feldern
// =============================================================================

type punkt struct {
	x, y int
}

type kreisV2 struct {
	mitte  punkt // Verbund-Typ als benanntes Feld
	radius int
}

type radV2 struct {
	umkreis  kreisV2 // Verbund-Typ als benanntes Feld
	speichen int
}

// =============================================================================
// BESTE VARIANTE: Anonyme Felder (Embedded Structs)
// =============================================================================

type kreis struct {
	punkt  // anonymes Feld (embedded struct)
	radius int
}

type rad struct {
	kreis    // anonymes Feld (embedded struct)
	speichen int
}

// =============================================================================
// Weitere Beispiele für eingebettete Verbünde
// =============================================================================

type adresse struct {
	straße string
	plz    int
	stadt  string
}

type person struct {
	name    string
	adresse // anonymes Feld
}

type firma struct {
	name    string
	adresse // anonymes Feld
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Anonyme Felder und eingebettete Verbünde ===\n")

	// -------------------------------------------------------------------------
	// 1. Variante 1: Flache Verbund-Typen
	// -------------------------------------------------------------------------

	fmt.Println("1. Variante 1: Flache Verbund-Typen")
	fmt.Println("   ---------------------------------")
	fmt.Println()

	fmt.Println("   Definition:")
	fmt.Println("   type kreisV1 struct {")
	fmt.Println("       x, y, radius int")
	fmt.Println("   }")
	fmt.Println("   type radV1 struct {")
	fmt.Println("       x, y, radius, speichen int")
	fmt.Println("   }")
	fmt.Println()

	var kV1 kreisV1
	kV1.x = 12
	kV1.y = 20
	kV1.radius = 30

	fmt.Printf("   kV1: %+v\n", kV1)
	fmt.Println()

	var rV1 radV1
	rV1.x = 5
	rV1.y = 3
	rV1.radius = 20
	rV1.speichen = 8

	fmt.Printf("   rV1: %+v\n", rV1)
	fmt.Println()

	fmt.Println("   Eigenschaften:")
	fmt.Println("   ✓ Einfacher, direkter Zugriff")
	fmt.Println("   ✗ Keine Strukturierung (alle Felder auf gleicher Ebene)")
	fmt.Println("   ✗ Wiederholung: x, y, radius kommen in beiden Typen vor")
	fmt.Println("   ✗ Keine Wiederverwendung von Code/Konzepten")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Variante 2: Hierarchische Verbund-Typen (benannte Felder)
	// -------------------------------------------------------------------------

	fmt.Println("2. Variante 2: Hierarchische Verbund-Typen (benannte Felder)")
	fmt.Println("   ----------------------------------------------------------")
	fmt.Println()

	fmt.Println("   Definition:")
	fmt.Println("   type punkt struct {")
	fmt.Println("       x, y int")
	fmt.Println("   }")
	fmt.Println("   type kreisV2 struct {")
	fmt.Println("       mitte  punkt  // benanntes Feld")
	fmt.Println("       radius int")
	fmt.Println("   }")
	fmt.Println("   type radV2 struct {")
	fmt.Println("       umkreis  kreisV2  // benanntes Feld")
	fmt.Println("       speichen int")
	fmt.Println("   }")
	fmt.Println()

	var kV2 kreisV2
	kV2.mitte.x = 12
	kV2.mitte.y = 20
	kV2.radius = 30

	fmt.Println("   Zugriff auf kV2:")
	fmt.Println("   kV2.mitte.x = 12")
	fmt.Println("   kV2.mitte.y = 20")
	fmt.Println("   kV2.radius = 30")
	fmt.Printf("   kV2: %+v\n", kV2)
	fmt.Println()

	var rV2 radV2
	rV2.umkreis.mitte.x = 5
	rV2.umkreis.mitte.y = 3
	rV2.umkreis.radius = 20
	rV2.speichen = 8

	fmt.Println("   Zugriff auf rV2:")
	fmt.Println("   rV2.umkreis.mitte.x = 5")
	fmt.Println("   rV2.umkreis.mitte.y = 3")
	fmt.Println("   rV2.umkreis.radius = 20")
	fmt.Println("   rV2.speichen = 8")
	fmt.Printf("   rV2: %+v\n", rV2)
	fmt.Println()

	fmt.Println("   Eigenschaften:")
	fmt.Println("   ✓ Gute Strukturierung (punkt, kreis als Konzepte)")
	fmt.Println("   ✓ Wiederverwendung (punkt wird mehrfach verwendet)")
	fmt.Println("   ✗ Umständlicher Zugriff (kompletter Pfad nötig)")
	fmt.Println("   ✗ rV2.umkreis.mitte.x ist lang und unübersichtlich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Beste Variante: Anonyme Felder (Embedded Structs)
	// -------------------------------------------------------------------------

	fmt.Println("3. Beste Variante: Anonyme Felder (Embedded Structs)")
	fmt.Println("   --------------------------------------------------")
	fmt.Println()

	fmt.Println("   Definition:")
	fmt.Println("   type punkt struct {")
	fmt.Println("       x, y int")
	fmt.Println("   }")
	fmt.Println("   type kreis struct {")
	fmt.Println("       punkt   // ANONYMES Feld (kein Feldname!)")
	fmt.Println("       radius int")
	fmt.Println("   }")
	fmt.Println("   type rad struct {")
	fmt.Println("       kreis    // ANONYMES Feld (kein Feldname!)")
	fmt.Println("       speichen int")
	fmt.Println("   }")
	fmt.Println()

	// Einfacher Zugriff - anonyme Felder können weggelassen werden!
	var k kreis
	k.x = 12 // punkt.x - anonymes Feld weggelassen!
	k.y = 20 // punkt.y - anonymes Feld weggelassen!
	k.radius = 30

	fmt.Println("   Zugriff auf k (anonyme Felder weggelassen):")
	fmt.Println("   k.x = 12       // statt k.punkt.x")
	fmt.Println("   k.y = 20       // statt k.punkt.y")
	fmt.Println("   k.radius = 30")
	fmt.Printf("   k: %+v\n", k)
	fmt.Println()

	var r rad
	r.x = 5       // kreis.punkt.x - beide anonyme Felder weggelassen!
	r.y = 3       // kreis.punkt.y - beide anonyme Felder weggelassen!
	r.radius = 20 // kreis.radius - anonymes Feld weggelassen!
	r.speichen = 8

	fmt.Println("   Zugriff auf r (anonyme Felder weggelassen):")
	fmt.Println("   r.x = 5        // statt r.kreis.punkt.x")
	fmt.Println("   r.y = 3        // statt r.kreis.punkt.y")
	fmt.Println("   r.radius = 20  // statt r.kreis.radius")
	fmt.Println("   r.speichen = 8")
	fmt.Printf("   r: %+v\n", r)
	fmt.Println()

	fmt.Println("   Eigenschaften:")
	fmt.Println("   ✓ Gute Strukturierung (punkt, kreis als Konzepte)")
	fmt.Println("   ✓ Wiederverwendung (punkt wird mehrfach verwendet)")
	fmt.Println("   ✓ Einfacher Zugriff (anonyme Felder können weggelassen werden)")
	fmt.Println("   ✓ Beste aus beiden Welten!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Implizite Feldnamen
	// -------------------------------------------------------------------------

	fmt.Println("4. Implizite Feldnamen bei anonymen Feldern:")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	fmt.Println("   Anonyme Felder haben implizit den Namen ihres Typs:")
	fmt.Println()

	// Mit %+v werden Feldnamen angezeigt
	fmt.Printf("   r mit Feldnamen: %+v\n", r)
	fmt.Println()

	fmt.Println("   Erklärung:")
	fmt.Println("   {kreis:{punkt:{x:5 y:3} radius:20} speichen:8}")
	fmt.Println("    ^^^^^^ ^^^^^^                     ^^^^^^^^")
	fmt.Println("    Impliziter Feldname \"kreis\"")
	fmt.Println("           Impliziter Feldname \"punkt\"")
	fmt.Println()

	fmt.Println("   Das anonyme Feld vom Typ 'kreis' heißt implizit 'kreis'")
	fmt.Println("   Das anonyme Feld vom Typ 'punkt' heißt implizit 'punkt'")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Alternative Zugriffswege
	// -------------------------------------------------------------------------

	fmt.Println("5. Alternative Zugriffswege:")
	fmt.Println("   --------------------------")
	fmt.Println()

	fmt.Println("   Alle folgenden Zugriffe sind äquivalent:")
	fmt.Println()

	// Drei äquivalente Wege, auf x zuzugreifen
	fmt.Println("   // Variante 1: Alle anonymen Felder weglassen (kürzeste Form)")
	r.x = 100
	fmt.Printf("   r.x = 100                 -> r.x = %d\n", r.x)
	fmt.Println()

	fmt.Println("   // Variante 2: Anonymes Feld 'kreis' mit angeben")
	r.kreis.x = 200
	fmt.Printf("   r.kreis.x = 200           -> r.x = %d\n", r.x)
	fmt.Println()

	fmt.Println("   // Variante 3: Alle Felder explizit angeben (längste Form)")
	r.kreis.punkt.x = 300
	fmt.Printf("   r.kreis.punkt.x = 300     -> r.x = %d\n", r.x)
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Alle drei Varianten greifen auf das GLEICHE Feld zu")
	fmt.Println("   ✓ In der Praxis: kürzeste Form verwenden (r.x)")
	fmt.Println("   ✓ Bei Mehrdeutigkeit: längere Form nötig")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Vergleich der drei Varianten
	// -------------------------------------------------------------------------

	fmt.Println("6. Direkter Vergleich der drei Varianten:")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	fmt.Println("   Zugriff auf x-Koordinate eines Rads:")
	fmt.Println()
	fmt.Println("   Variante 1 (flach):      rV1.x")
	fmt.Println("   Variante 2 (benannt):    rV2.umkreis.mitte.x")
	fmt.Println("   Variante 3 (anonym):     r.x")
	fmt.Println()
	fmt.Println("   -> Variante 3 kombiniert gute Struktur mit einfachem Zugriff!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Literale von Verbünden mit anonymen Feldern
	// -------------------------------------------------------------------------

	fmt.Println("7. Literale von Verbünden mit anonymen Feldern:")
	fmt.Println("   ----------------------------------------------")
	fmt.Println()

	// Syntax-Variante 1: Werte ohne Feldnamen (unübersichtlich)
	fmt.Println("   Syntax-Variante 1: Werte ohne Feldnamen")
	r1 := rad{kreis{punkt{22, 33}, 55}, 4}
	fmt.Println("   r1 := rad{kreis{punkt{22, 33}, 55}, 4}")
	fmt.Printf("   r1: %+v\n", r1)
	fmt.Println("   -> Kompakt, aber schwer lesbar bei komplexen Strukturen")
	fmt.Println()

	// Syntax-Variante 2: Mit expliziten Feldnamen (übersichtlich)
	fmt.Println("   Syntax-Variante 2: Mit Feldnamen (alle Felder)")
	r2 := rad{
		kreis: kreis{ // impliziter Feldname
			punkt: punkt{ // impliziter Feldname
				x: 22,
				y: 33,
			},
			radius: 55,
		},
		speichen: 4,
	}
	fmt.Println("   r2 := rad{")
	fmt.Println("       kreis: kreis{")
	fmt.Println("           punkt: punkt{")
	fmt.Println("               x: 22,")
	fmt.Println("               y: 33,")
	fmt.Println("           },")
	fmt.Println("           radius: 55,")
	fmt.Println("       },")
	fmt.Println("       speichen: 4,")
	fmt.Println("   }")
	fmt.Printf("   r2: %+v\n", r2)
	fmt.Println("   -> Übersichtlich und selbsterklärend")
	fmt.Println()

	// Syntax-Variante 2b: Nur bestimmte Felder (Rest = Zero-Values)
	fmt.Println("   Syntax-Variante 2b: Mit Feldnamen (nur bestimmte Felder)")
	r3 := rad{
		kreis: kreis{
			punkt: punkt{
				y: 99, // nur y setzen, x bleibt 0
			},
			// radius bleibt 0
		},
		// speichen bleibt 0
	}
	fmt.Println("   r3 := rad{")
	fmt.Println("       kreis: kreis{")
	fmt.Println("           punkt: punkt{")
	fmt.Println("               y: 99,  // nur y setzen")
	fmt.Println("           },")
	fmt.Println("       },")
	fmt.Println("   }")
	fmt.Printf("   r3: %+v\n", r3)
	fmt.Println("   -> Nicht gesetzte Felder haben Zero-Values (0)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Einschränkung: Nicht zwei anonyme Felder gleichen Typs
	// -------------------------------------------------------------------------

	fmt.Println("8. Einschränkung: Nicht zwei anonyme Felder gleichen Typs:")
	fmt.Println("   --------------------------------------------------------")
	fmt.Println()

	fmt.Println("   UNZULÄSSIG (würde nicht kompilieren):")
	fmt.Println("   type linie struct {")
	fmt.Println("       punkt  // Start-Punkt")
	fmt.Println("       punkt  // End-Punkt - FEHLER! Gleicher Typ wie oben")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("   Problem:")
	fmt.Println("   - Beide Felder hätten den impliziten Namen 'punkt'")
	fmt.Println("   - Zugriffspfad wäre mehrdeutig")
	fmt.Println("   - Compiler könnte nicht entscheiden")
	fmt.Println()

	fmt.Println("   LÖSUNG: Benannte Felder verwenden")
	type linie struct {
		start punkt // benanntes Feld
		ende  punkt // benanntes Feld
	}

	fmt.Println("   type linie struct {")
	fmt.Println("       start punkt  // benanntes Feld")
	fmt.Println("       ende  punkt  // benanntes Feld")
	fmt.Println("   }")

	l := linie{
		start: punkt{x: 10, y: 20},
		ende:  punkt{x: 50, y: 60},
	}
	fmt.Printf("   l: %+v\n", l)
	fmt.Printf("   l.start.x = %d, l.ende.x = %d\n", l.start.x, l.ende.x)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Weiteres Beispiel: Person und Firma mit Adresse
	// -------------------------------------------------------------------------

	fmt.Println("9. Weiteres Beispiel: Person und Firma mit Adresse:")
	fmt.Println("   -------------------------------------------------")
	fmt.Println()

	fmt.Println("   Definition:")
	fmt.Println("   type adresse struct {")
	fmt.Println("       straße string")
	fmt.Println("       plz    int")
	fmt.Println("       stadt  string")
	fmt.Println("   }")
	fmt.Println("   type person struct {")
	fmt.Println("       name    string")
	fmt.Println("       adresse        // anonymes Feld")
	fmt.Println("   }")
	fmt.Println("   type firma struct {")
	fmt.Println("       name    string")
	fmt.Println("       adresse        // anonymes Feld")
	fmt.Println("   }")
	fmt.Println()

	p := person{
		name: "Anna Müller",
		adresse: adresse{
			straße: "Hauptstraße 42",
			plz:    10115,
			stadt:  "Berlin",
		},
	}

	f := firma{
		name: "Tech GmbH",
		adresse: adresse{
			straße: "Innovation Park 1",
			plz:    80331,
			stadt:  "München",
		},
	}

	fmt.Printf("   Person: %+v\n", p)
	fmt.Printf("   Firma:  %+v\n", f)
	fmt.Println()

	// Einfacher Zugriff auf eingebettete Felder
	fmt.Println("   Einfacher Zugriff auf Adressfelder:")
	fmt.Printf("   p.stadt = %s  (statt p.adresse.stadt)\n", p.stadt)
	fmt.Printf("   f.stadt = %s  (statt f.adresse.stadt)\n", f.stadt)
	fmt.Println()

	// Änderung über verkürzten Pfad
	p.plz = 10117
	f.plz = 80333
	fmt.Println("   Nach Änderung p.plz und f.plz:")
	fmt.Printf("   p.plz = %d\n", p.plz)
	fmt.Printf("   f.plz = %d\n", f.plz)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Zusammenfassung
	// -------------------------------------------------------------------------

	fmt.Println("10. Zusammenfassung:")
	fmt.Println("    ---------------")
	fmt.Println()

	fmt.Println("   Anonyme Felder (Embedded Structs):")
	fmt.Println("   ✓ Deklaration ohne Feldnamen, nur Typ-Name")
	fmt.Println("   ✓ Impliziter Feldname = Typ-Name")
	fmt.Println("   ✓ Verkürzter Zugriff: Anonyme Felder können weggelassen werden")
	fmt.Println("   ✓ Kombiniert Struktur mit einfachem Zugriff")
	fmt.Println()

	fmt.Println("   Vorteile:")
	fmt.Println("   ✓ Code-Wiederverwendung (z.B. punkt, adresse)")
	fmt.Println("   ✓ Bessere Strukturierung")
	fmt.Println("   ✓ Einfache Syntax für Zugriff")
	fmt.Println("   ✓ Häufig in Go-Bibliotheken verwendet")
	fmt.Println()

	fmt.Println("   Einschränkungen:")
	fmt.Println("   ✗ Nur ein anonymes Feld pro Typ erlaubt")
	fmt.Println("   ✗ Literale nicht verkürzt schreibbar")
	fmt.Println()

	fmt.Println("   Verwendung:")
	fmt.Println("   ✓ Immer wenn Verbünde andere Verbünde enthalten")
	fmt.Println("   ✓ Zur Vermeidung von Feld-Wiederholungen")
	fmt.Println("   ✓ Für objektorientierte Konzepte (siehe spätere Kapitel)")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}
