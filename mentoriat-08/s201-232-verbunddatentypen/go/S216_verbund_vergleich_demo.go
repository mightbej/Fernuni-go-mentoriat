package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Vergleich von Verbund-Variablen (Abschnitt 8.2.x)
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Vergleich von Verbund-Variablen mit == Operator
//    - Der Vergleich a == b ist äquivalent zu:
//      a.x == b.x && a.y == b.y && a.z == b.z (alle Felder werden verglichen)
//
// 2. Unterschied zwischen Wert-Gleichheit und Identität
//    - Zwei Verbünde können den gleichen Wert haben (a == b ist true)
//    - Aber trotzdem unterschiedliche Variablen sein (&a != &b ist true)
//    - Gleicher Wert ≠ gleiche Speicherstelle
//
// 3. Verschiedene Vergleichsszenarien
//    - Verbünde mit identischen Werten
//    - Verbünde mit unterschiedlichen Werten
//    - Einzelfeld-Vergleiche vs. Gesamt-Vergleich
//
// Ausführen mit: go run S216_verbund_vergleich_demo.go
//
// =============================================================================

// Definition verschiedener Verbund-Datentypen
type punkt struct {
	x int
	y int
	z int
}

type person struct {
	vorname  string
	nachname string
	alter    int
}

type farbe struct {
	rot  int
	grün int
	blau int
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Vergleich von Verbund-Variablen ===\n")

	// -------------------------------------------------------------------------
	// 1. Grundlegender Vergleich von Verbünden
	// -------------------------------------------------------------------------

	fmt.Println("1. Grundlegender Vergleich von Verbünden:")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	a := punkt{x: 10, y: 20, z: 30}
	b := punkt{x: 10, y: 20, z: 30}
	c := punkt{x: 10, y: 20, z: 99}

	fmt.Printf("   a = %+v (Adresse: %p)\n", a, &a)
	fmt.Printf("   b = %+v (Adresse: %p)\n", b, &b)
	fmt.Printf("   c = %+v (Adresse: %p)\n", c, &c)
	fmt.Println()

	// Wert-Vergleich
	fmt.Println("   Wert-Vergleich mit == Operator:")
	fmt.Printf("   a == b? %t  (gleiche Werte)\n", a == b)
	fmt.Printf("   a == c? %t  (unterschiedliche z-Werte)\n", a == c)
	fmt.Printf("   b == c? %t  (unterschiedliche z-Werte)\n", b == c)
	fmt.Println()

	// Adress-Vergleich
	fmt.Println("   Adress-Vergleich (Zeiger auf Verbünde):")
	fmt.Printf("   &a == &b? %t  (unterschiedliche Variablen!)\n", &a == &b)
	fmt.Printf("   &a == &c? %t  (unterschiedliche Variablen!)\n", &a == &c)
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ a == b ist true  (Werte sind gleich)")
	fmt.Println("   ✓ &a != &b ist true  (unterschiedliche Speicherstellen)")
	fmt.Println("   ✓ Gleicher Wert bedeutet NICHT gleiche Variable!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Interner Vergleichsmechanismus: Feld für Feld
	// -------------------------------------------------------------------------

	fmt.Println("2. Interner Vergleichsmechanismus:")
	fmt.Println("   --------------------------------")
	fmt.Println()

	fmt.Println("   Der Vergleich a == b ist äquivalent zu:")
	fmt.Println("   a.x == b.x && a.y == b.y && a.z == b.z")
	fmt.Println()

	// Manuelle Einzelfeld-Vergleiche
	fmt.Println("   Einzelfeld-Vergleiche für a und b:")
	fmt.Printf("   a.x == b.x? %t  (%d == %d)\n", a.x == b.x, a.x, b.x)
	fmt.Printf("   a.y == b.y? %t  (%d == %d)\n", a.y == b.y, a.y, b.y)
	fmt.Printf("   a.z == b.z? %t  (%d == %d)\n", a.z == b.z, a.z, b.z)
	fmt.Println()

	// Kombinierter Vergleich
	manuellAB := a.x == b.x && a.y == b.y && a.z == b.z
	fmt.Printf("   Manueller Vergleich: a.x==b.x && a.y==b.y && a.z==b.z = %t\n", manuellAB)
	fmt.Printf("   Operator-Vergleich:  a == b = %t\n", a == b)
	fmt.Printf("   Beide sind gleich? %t\n", manuellAB == (a == b))
	fmt.Println()

	// Einzelfeld-Vergleiche für a und c
	fmt.Println("   Einzelfeld-Vergleiche für a und c:")
	fmt.Printf("   a.x == c.x? %t  (%d == %d)\n", a.x == c.x, a.x, c.x)
	fmt.Printf("   a.y == c.y? %t  (%d == %d)\n", a.y == c.y, a.y, c.y)
	fmt.Printf("   a.z == c.z? %t  (%d == %d) <- Unterschied!\n", a.z == c.z, a.z, c.z)
	fmt.Println()

	manuellAC := a.x == c.x && a.y == c.y && a.z == c.z
	fmt.Printf("   Manueller Vergleich: a.x==c.x && a.y==c.y && a.z==c.z = %t\n", manuellAC)
	fmt.Printf("   Operator-Vergleich:  a == c = %t\n", a == c)
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Der == Operator vergleicht ALLE Felder")
	fmt.Println("   ✓ Nur wenn ALLE Felder gleich sind, ist a == b true")
	fmt.Println("   ✓ Ein einziges unterschiedliches Feld macht a == b false")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Gleicher Wert ≠ gleiche Identität
	// -------------------------------------------------------------------------

	fmt.Println("3. Gleicher Wert ≠ gleiche Identität:")
	fmt.Println("   ------------------------------------")
	fmt.Println()

	p1 := person{vorname: "Anna", nachname: "Schmidt", alter: 30}
	p2 := person{vorname: "Anna", nachname: "Schmidt", alter: 30}

	fmt.Printf("   p1 = %+v\n", p1)
	fmt.Printf("   p2 = %+v\n", p2)
	fmt.Println()

	fmt.Println("   Vergleiche:")
	fmt.Printf("   p1 == p2? %t  (Werte sind identisch)\n", p1 == p2)
	fmt.Printf("   &p1 == &p2? %t  (unterschiedliche Speicherstellen)\n", &p1 == &p2)
	fmt.Println()

	fmt.Printf("   Adressen:\n")
	fmt.Printf("   &p1 = %p\n", &p1)
	fmt.Printf("   &p2 = %p\n", &p2)
	//fmt.Printf("   Differenz: %d Bytes\n", uintptr(&p2)-uintptr(&p1))
	fmt.Println()

	fmt.Println("   INTERPRETATION:")
	fmt.Println("   ✓ p1 und p2 haben den gleichen Wert (Wert-Gleichheit)")
	fmt.Println("   ✓ p1 und p2 sind unterschiedliche Variablen (Identität)")
	fmt.Println("   ✓ Sie belegen unterschiedliche Speicherbereiche")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Zuweisung vs. Vergleich
	// -------------------------------------------------------------------------

	fmt.Println("4. Zuweisung vs. Vergleich:")
	fmt.Println("   -------------------------")
	fmt.Println()

	f1 := farbe{rot: 255, grün: 0, blau: 0}
	f2 := farbe{rot: 0, grün: 255, blau: 0}

	fmt.Printf("   Vor Zuweisung:\n")
	fmt.Printf("   f1 = %+v (Adresse: %p)\n", f1, &f1)
	fmt.Printf("   f2 = %+v (Adresse: %p)\n", f2, &f2)
	fmt.Printf("   f1 == f2? %t\n", f1 == f2)
	fmt.Println()

	// Zuweisung: Wert wird kopiert
	f2 = f1
	fmt.Printf("   Nach Zuweisung f2 = f1:\n")
	fmt.Printf("   f1 = %+v (Adresse: %p)\n", f1, &f1)
	fmt.Printf("   f2 = %+v (Adresse: %p)\n", f2, &f2)
	fmt.Printf("   f1 == f2? %t  (jetzt gleicher Wert)\n", f1 == f2)
	fmt.Printf("   &f1 == &f2? %t  (immer noch unterschiedliche Adressen)\n", &f1 == &f2)
	fmt.Println()

	// Änderung einer Variable
	f1.rot = 128
	fmt.Printf("   Nach Änderung f1.rot = 128:\n")
	fmt.Printf("   f1 = %+v\n", f1)
	fmt.Printf("   f2 = %+v (unverändert! f2 ist eine Kopie)\n", f2)
	fmt.Printf("   f1 == f2? %t  (jetzt unterschiedliche Werte)\n", f1 == f2)
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Zuweisung f2 = f1 kopiert den WERT")
	fmt.Println("   ✓ f1 und f2 bleiben unterschiedliche Variablen")
	fmt.Println("   ✓ Änderungen an f1 betreffen f2 nicht")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Zeiger-Vergleich vs. Wert-Vergleich
	// -------------------------------------------------------------------------

	fmt.Println("5. Zeiger-Vergleich vs. Wert-Vergleich:")
	fmt.Println("   --------------------------------------")
	fmt.Println()

	k1 := punkt{x: 5, y: 10, z: 15}
	k2 := punkt{x: 5, y: 10, z: 15}
	pk1 := &k1
	pk2 := &k2
	pk3 := &k1 // zeigt auf die gleiche Variable wie pk1

	fmt.Printf("   k1 = %+v (Adresse: %p)\n", k1, &k1)
	fmt.Printf("   k2 = %+v (Adresse: %p)\n", k2, &k2)
	fmt.Println()

	fmt.Println("   Wert-Vergleich:")
	fmt.Printf("   k1 == k2? %t  (Werte sind gleich)\n", k1 == k2)
	fmt.Println()

	fmt.Println("   Zeiger-Vergleich:")
	fmt.Printf("   pk1 == pk2? %t  (zeigen auf unterschiedliche Variablen)\n", pk1 == pk2)
	fmt.Printf("   pk1 == pk3? %t  (zeigen auf die GLEICHE Variable k1)\n", pk1 == pk3)
	fmt.Printf("   *pk1 == *pk2? %t  (dereferenzierte Werte sind gleich)\n", *pk1 == *pk2)
	fmt.Println()

	fmt.Printf("   Adressen:\n")
	fmt.Printf("   pk1 zeigt auf: %p\n", pk1)
	fmt.Printf("   pk2 zeigt auf: %p\n", pk2)
	fmt.Printf("   pk3 zeigt auf: %p\n", pk3)
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Zeiger-Vergleich: pk1 == pk2 vergleicht ADRESSEN")
	fmt.Println("   ✓ Wert-Vergleich: *pk1 == *pk2 vergleicht WERTE")
	fmt.Println("   ✓ pk1 == pk3 ist true (gleiche Adresse)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Verbünde als Funktionsparameter
	// -------------------------------------------------------------------------

	fmt.Println("6. Verbünde als Funktionsparameter:")
	fmt.Println("   ----------------------------------")
	fmt.Println()

	original := punkt{x: 100, y: 200, z: 300}
	fmt.Printf("   Original vor Funktionsaufruf: %+v (Adresse: %p)\n", original, &original)
	fmt.Println()

	fmt.Println("   Aufruf: verändereKopie(original)")
	verändereKopie(original)
	fmt.Printf("   Original nach verändereKopie: %+v (unverändert!)\n", original)
	fmt.Println()

	fmt.Println("   Aufruf: verändereOriginal(&original)")
	verändereOriginal(&original)
	fmt.Printf("   Original nach verändereOriginal: %+v (geändert!)\n", original)
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Call-by-Value: Funktion erhält KOPIE (keine Änderung am Original)")
	fmt.Println("   ✓ Call-by-Reference: Funktion erhält ZEIGER (Änderung am Original)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Vergleich mit Zero-Values
	// -------------------------------------------------------------------------

	fmt.Println("7. Vergleich mit Zero-Values:")
	fmt.Println("   ---------------------------")
	fmt.Println()

	var z1 punkt
	var z2 punkt
	z3 := punkt{x: 0, y: 0, z: 0}

	fmt.Printf("   z1 (var z1 punkt): %+v\n", z1)
	fmt.Printf("   z2 (var z2 punkt): %+v\n", z2)
	fmt.Printf("   z3 (punkt{0,0,0}): %+v\n", z3)
	fmt.Println()

	fmt.Println("   Vergleiche:")
	fmt.Printf("   z1 == z2? %t  (beide haben Zero-Values)\n", z1 == z2)
	fmt.Printf("   z1 == z3? %t  (explizite Nullwerte = Zero-Values)\n", z1 == z3)
	fmt.Printf("   &z1 == &z2? %t  (unterschiedliche Variablen)\n", &z1 == &z2)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Zusammenfassung mit Visualisierung
	// -------------------------------------------------------------------------

	fmt.Println("8. Zusammenfassung - Visualisierung:")
	fmt.Println("   -----------------------------------")
	fmt.Println()

	demo1 := punkt{x: 42, y: 43, z: 44}
	demo2 := punkt{x: 42, y: 43, z: 44}

	fmt.Println("   Speicherlayout:")
	fmt.Println()
	fmt.Println("   Adresse:  Variable:  Werte:")
	fmt.Println("   ────────────────────────────────────────")
	fmt.Printf("   %p  demo1      {x:42, y:43, z:44}\n", &demo1)
	fmt.Printf("   %p  demo2      {x:42, y:43, z:44}\n", &demo2)
	fmt.Println("   ────────────────────────────────────────")
	fmt.Println()

	fmt.Println("   Vergleiche:")
	fmt.Printf("   demo1 == demo2           -> %t  (Werte sind gleich)\n", demo1 == demo2)
	fmt.Printf("   &demo1 == &demo2         -> %t  (Adressen sind unterschiedlich)\n", &demo1 == &demo2)
	fmt.Printf("   demo1.x == demo2.x       -> %t  (Einzelfeld-Vergleich)\n", demo1.x == demo2.x)
	fmt.Println()

	fmt.Println("   KERNAUSSAGE:")
	fmt.Println("   ═══════════════════════════════════════════════════════════")
	fmt.Println("   ✓ a == b vergleicht alle Felder (äquivalent zu:")
	fmt.Println("     a.x==b.x && a.y==b.y && a.z==b.z)")
	fmt.Println()
	fmt.Println("   ✓ Wenn a == b true ist (gleiche Werte), sind a und b")
	fmt.Println("     trotzdem unterschiedliche Variablen")
	fmt.Println()
	fmt.Println("   ✓ &a != &b ist true (unterschiedliche Speicherstellen)")
	fmt.Println("   ═══════════════════════════════════════════════════════════")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// verändereKopie erhält eine Kopie des Verbunds (Call-by-Value)
func verändereKopie(p punkt) {
	fmt.Printf("   -> In verändereKopie: Parameter-Adresse %p\n", &p)
	p.x = 999
	fmt.Printf("   -> In verändereKopie: p.x = %d (nur Kopie geändert)\n", p.x)
}

// verändereOriginal erhält einen Zeiger auf den Verbund (Call-by-Reference)
func verändereOriginal(p *punkt) {
	fmt.Printf("   -> In verändereOriginal: Zeiger zeigt auf %p\n", p)
	p.x = 999
	fmt.Printf("   -> In verändereOriginal: p.x = %d (Original geändert)\n", p.x)
}
