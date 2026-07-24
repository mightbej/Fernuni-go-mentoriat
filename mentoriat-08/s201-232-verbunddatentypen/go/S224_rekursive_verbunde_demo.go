package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Rekursive Verbund-Typen (Abschnitt 8.3.1)
// =============================================================================
//
// Dieses Programm demonstriert:
//
// 1. Unzulässige rekursive Definition
//    - Verbund-Typ darf sich nicht selbst als Feld-Typ enthalten
//    - Würde zu unendlicher Rekursion führen
//    - Beispiel (nicht kompilierbar): type mensch struct { vater mensch }
//
// 2. Zulässige rekursive Definition mit Zeigern
//    - Verbund-Typ DARF Zeiger auf sich selbst enthalten
//    - Zeiger-Typ *T ist als Feld-Typ erlaubt
//    - nil-Werte beenden die Rekursion
//
// 3. Praktische Anwendung: Stammbaum
//    - Hierarchische Struktur mit Eltern-Kind-Beziehungen
//    - Zugriff über mehrere Ebenen (Eltern, Großeltern, etc.)
//
// 4. Weitere rekursive Datenstrukturen
//    - Verkettete Listen
//    - Binäre Bäume
//
// Ausführen mit: go run S224_rekursive_verbunde_demo.go
//
// =============================================================================

// =============================================================================
// UNZULÄSSIGE Definition (auskommentiert, würde nicht kompilieren)
// =============================================================================

// NICHT ZULÄSSIG: Unendlich rekursive Definition
// type menschFalsch struct {
//     name          string
//     vater, mutter menschFalsch  // FEHLER: Typ selbst als Feld-Typ
// }
//
// Problem: Jeder mensch würde vater und mutter enthalten,
//          die wiederum vater und mutter enthalten würden,
//          die wiederum... -> unendliche Rekursion!

// =============================================================================
// ZULÄSSIGE Definition mit Zeiger-Typen
// =============================================================================

// mensch: Rekursiver Verbund-Typ für Stammbaum
type mensch struct {
	name          string
	vater, mutter *mensch // Zeiger auf Verbund-Typ (rekursiv, aber zulässig!)
}

// knoten: Rekursiver Verbund-Typ für verkettete Liste
type knoten struct {
	wert int
	next *knoten // Zeiger auf nächsten Knoten
}

// baumKnoten: Rekursiver Verbund-Typ für binären Baum
type baumKnoten struct {
	wert   int
	links  *baumKnoten // Linkes Kind
	rechts *baumKnoten // Rechtes Kind
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Rekursive Verbund-Typen ===\n")

	// -------------------------------------------------------------------------
	// 1. Warum direkte Rekursion nicht zulässig ist
	// -------------------------------------------------------------------------

	fmt.Println("1. Warum direkte Rekursion nicht zulässig ist:")
	fmt.Println("   --------------------------------------------")
	fmt.Println()

	fmt.Println("   UNZULÄSSIG (würde nicht kompilieren):")
	fmt.Println("   type menschFalsch struct {")
	fmt.Println("       name          string")
	fmt.Println("       vater, mutter menschFalsch  // FEHLER!")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("   Problem:")
	fmt.Println("   - Jeder mensch enthält vater und mutter")
	fmt.Println("   - Jeder vater/mutter enthält wiederum vater und mutter")
	fmt.Println("   - Dies setzt sich unendlich fort")
	fmt.Println("   - Unendliche Rekursion, kein Ende möglich")
	fmt.Println("   - Speichergröße eines mensch wäre unendlich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Warum Zeiger-Rekursion zulässig ist
	// -------------------------------------------------------------------------

	fmt.Println("2. Warum Zeiger-Rekursion zulässig ist:")
	fmt.Println("   -------------------------------------")
	fmt.Println()

	fmt.Println("   ZULÄSSIG:")
	fmt.Println("   type mensch struct {")
	fmt.Println("       name          string")
	fmt.Println("       vater, mutter *mensch  // OK! Zeiger-Typ")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("   Lösung:")
	fmt.Println("   - Felder sind Zeiger, nicht vollständige Verbünde")
	fmt.Println("   - Zeiger haben feste Größe (z.B. 8 Bytes auf 64-Bit)")
	fmt.Println("   - Zeiger können nil sein -> Rekursion endet")
	fmt.Println("   - Speichergröße eines mensch ist bekannt und endlich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Einfaches Beispiel: Drei Generationen
	// -------------------------------------------------------------------------

	fmt.Println("3. Einfaches Beispiel - Drei Generationen:")
	fmt.Println("   ----------------------------------------")
	fmt.Println()

	// Großeltern (keine Eltern bekannt -> nil)
	opa := mensch{name: "Wilhelm", vater: nil, mutter: nil}
	oma := mensch{name: "Elfriede", vater: nil, mutter: nil}

	fmt.Printf("   Großeltern erstellt:\n")
	fmt.Printf("   opa = %+v\n", opa)
	fmt.Printf("   oma = %+v\n", oma)
	fmt.Println()

	// Eltern (Zeiger auf Großeltern)
	vater := mensch{name: "Karl", vater: &opa, mutter: &oma}
	mutter := mensch{name: "Maria", vater: nil, mutter: nil}
	// mutter := mensch{name: "Maria"}

	fmt.Printf("   Eltern erstellt:\n")
	fmt.Printf("   vater = %+v\n", vater)
	fmt.Printf("   mutter = %+v\n", mutter)
	fmt.Println()

	// Kind (Zeiger auf Eltern)
	kind := mensch{name: "Lukas", vater: &vater, mutter: &mutter}

	fmt.Printf("   Kind erstellt:\n")
	fmt.Printf("   kind = %+v\n", kind)
	fmt.Println()

	// Zugriff über mehrere Ebenen
	fmt.Println("   Zugriff über mehrere Ebenen:")
	fmt.Printf("   kind.name:                   %s\n", kind.name)
	fmt.Printf("   kind.vater.name:             %s\n", kind.vater.name)
	fmt.Printf("   kind.mutter.name:            %s\n", kind.mutter.name)
	fmt.Printf("   kind.vater.vater.name (Opa): %s\n", kind.vater.vater.name)
	fmt.Printf("   kind.vater.mutter.name (Oma): %s\n", kind.vater.mutter.name)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Vollständiger Stammbaum aus dem Text
	// -------------------------------------------------------------------------

	fmt.Println("4. Vollständiger Stammbaum (wie im Text):")
	fmt.Println("   ---------------------------------------")
	fmt.Println()

	// Großeltern-Generation (Blätter des Stammbaums)
	achim := mensch{"Achim", nil, nil}
	anita := mensch{"Anna", nil, nil}
	bert := mensch{"Bert", nil, nil}
	bärbel := mensch{"Bärbel", nil, nil}

	fmt.Println("   Großeltern-Generation:")
	fmt.Printf("   achim:  %+v\n", achim)
	fmt.Printf("   anita:  %+v\n", anita)
	fmt.Printf("   bert:   %+v\n", bert)
	fmt.Printf("   bärbel: %+v\n", bärbel)
	fmt.Println()

	// Eltern-Generation
	chris := mensch{"Chris", &achim, &anita}
	dagmar := mensch{"Dora", &bert, &bärbel}

	fmt.Println("   Eltern-Generation:")
	fmt.Printf("   chris:  %+v\n", chris)
	fmt.Printf("   dagmar: %+v\n", dagmar)
	fmt.Println()

	// Kind-Generation
	erik := mensch{"Dirk", &chris, &dagmar}

	fmt.Println("   Kind-Generation:")
	fmt.Printf("   erik: %+v\n", erik)
	fmt.Println()

	// Komplexer Zugriff
	fmt.Println("   Komplexe Stammbaumabfrage:")
	fmt.Printf("   %s hat Vater %s und Mutter %s\n",
		erik.name, erik.vater.name, erik.mutter.name)
	fmt.Printf("   sowie die Opas %s und %s\n",
		erik.vater.vater.name, erik.mutter.vater.name)
	fmt.Printf("   und die Omas %s und %s\n",
		erik.vater.mutter.name, erik.mutter.mutter.name)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Umgang mit nil-Werten (Abbruch der Rekursion)
	// -------------------------------------------------------------------------

	fmt.Println("5. Umgang mit nil-Werten:")
	fmt.Println("   -----------------------")
	fmt.Println()

	// Person ohne bekannte Eltern
	person := mensch{name: "Max", vater: nil, mutter: nil}

	fmt.Printf("   person = %+v\n", person)
	fmt.Printf("   person.vater: %v (nil)\n", person.vater)
	fmt.Printf("   person.mutter: %v (nil)\n", person.mutter)
	fmt.Println()

	// VORSICHT: Dereferenzierung von nil führt zu Panic!
	fmt.Println("   WARNUNG: Zugriff auf nil-Zeiger:")
	fmt.Println("   person.vater == nil? ->", person.vater == nil)
	fmt.Println()

	// Sichere Überprüfung vor Zugriff
	if person.vater != nil {
		fmt.Printf("   Vater: %s\n", person.vater.name)
	} else {
		fmt.Println("   Vater: unbekannt (nil)")
	}

	if person.mutter != nil {
		fmt.Printf("   Mutter: %s\n", person.mutter.name)
	} else {
		fmt.Println("   Mutter: unbekannt (nil)")
	}
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ nil beendet die Rekursion")
	fmt.Println("   ✓ Immer auf nil prüfen vor Dereferenzierung!")
	fmt.Println("   ✓ Sonst: Runtime Panic (Programmabsturz)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Weitere rekursive Datenstruktur: Verkettete Liste
	// -------------------------------------------------------------------------

	fmt.Println("6. Verkettete Liste (Linked List):")
	fmt.Println("   --------------------------------")
	fmt.Println()

	// Liste aufbauen: 10 -> 20 -> 30 -> nil
	knoten3 := knoten{wert: 30, next: nil}
	knoten2 := knoten{wert: 20, next: &knoten3}
	knoten1 := knoten{wert: 10, next: &knoten2}

	fmt.Println("   Liste: 10 -> 20 -> 30 -> nil")
	fmt.Println()

	fmt.Printf("   knoten1: %+v\n", knoten1)
	fmt.Printf("   knoten2: %+v\n", knoten2)
	fmt.Printf("   knoten3: %+v\n", knoten3)
	fmt.Println()

	// Liste durchlaufen
	fmt.Println("   Liste durchlaufen:")
	aktuellerKnoten := &knoten1
	position := 1
	for aktuellerKnoten != nil {
		fmt.Printf("   Position %d: Wert = %d (Adresse: %p)\n",
			position, aktuellerKnoten.wert, aktuellerKnoten)
		aktuellerKnoten = aktuellerKnoten.next
		position++
	}
	fmt.Println()

	fmt.Println("   Struktur:")
	fmt.Println("   ✓ Jeder Knoten zeigt auf den nächsten")
	fmt.Println("   ✓ Letzter Knoten hat next = nil")
	fmt.Println("   ✓ Dynamische Größe möglich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Weitere rekursive Datenstruktur: Binärer Baum
	// -------------------------------------------------------------------------

	fmt.Println("7. Binärer Baum (Binary Tree):")
	fmt.Println("   ----------------------------")
	fmt.Println()

	// Baum aufbauen:
	//        50
	//       /  \
	//     30    70
	//    /  \   /  \
	//  20  40 60  80

	baum := &baumKnoten{
		wert: 50,
		links: &baumKnoten{
			wert: 30,
			links: &baumKnoten{
				wert:   20,
				links:  nil,
				rechts: nil,
			},
			rechts: &baumKnoten{
				wert:   40,
				links:  nil,
				rechts: nil,
			},
		},
		rechts: &baumKnoten{
			wert: 70,
			links: &baumKnoten{
				wert:   60,
				links:  nil,
				rechts: nil,
			},
			rechts: &baumKnoten{
				wert:   80,
				links:  nil,
				rechts: nil,
			},
		},
	}

	fmt.Println("   Baum-Struktur:")
	fmt.Println("           50")
	fmt.Println("          /  \\")
	fmt.Println("        30    70")
	fmt.Println("       /  \\   /  \\")
	fmt.Println("     20  40 60  80")
	fmt.Println()

	fmt.Printf("   Wurzel:              %d\n", baum.wert)
	fmt.Printf("   Linkes Kind:         %d\n", baum.links.wert)
	fmt.Printf("   Rechtes Kind:        %d\n", baum.rechts.wert)
	fmt.Printf("   Links-Links:         %d\n", baum.links.links.wert)
	fmt.Printf("   Links-Rechts:        %d\n", baum.links.rechts.wert)
	fmt.Printf("   Rechts-Links:        %d\n", baum.rechts.links.wert)
	fmt.Printf("   Rechts-Rechts:       %d\n", baum.rechts.rechts.wert)
	fmt.Println()

	fmt.Println("   In-Order Traversierung (Links-Wurzel-Rechts):")
	fmt.Print("   ")
	inOrder(baum)
	fmt.Println()
	fmt.Println()

	fmt.Println("   Struktur:")
	fmt.Println("   ✓ Jeder Knoten hat max. zwei Kinder (links, rechts)")
	fmt.Println("   ✓ Blätter haben links = nil und rechts = nil")
	fmt.Println("   ✓ Hierarchische Baumstruktur")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Visualisierung: Speicher-Layout
	// -------------------------------------------------------------------------

	fmt.Println("8. Speicher-Layout eines rekursiven Verbunds:")
	fmt.Println("   -------------------------------------------")
	fmt.Println()

	p1 := mensch{name: "Papa", vater: nil, mutter: nil}
	p2 := mensch{name: "Sohn", vater: &p1, mutter: nil}

	fmt.Println("   Situation: Sohn hat Zeiger auf Papa")
	fmt.Println()
	fmt.Printf("   p1 (Papa): Adresse %p\n", &p1)
	fmt.Printf("   p2 (Sohn): Adresse %p\n", &p2)
	fmt.Println()
	fmt.Printf("   p2.vater zeigt auf: %p (= Adresse von p1)\n", p2.vater)
	fmt.Println()

	fmt.Println("   Speicher:")
	fmt.Println("   ┌─────────────────────────────────┐")
	fmt.Printf("   │ p1 (Papa)        @ %p │\n", &p1)
	fmt.Println("   ├─────────────────────────────────┤")
	fmt.Println("   │ name:   \"Papa\"                  │")
	fmt.Println("   │ vater:  nil                     │")
	fmt.Println("   │ mutter: nil                     │")
	fmt.Println("   └─────────────────────────────────┘")
	fmt.Println()
	fmt.Println("   ┌─────────────────────────────────┐")
	fmt.Printf("   │ p2 (Sohn)        @ %p │\n", &p2)
	fmt.Println("   ├─────────────────────────────────┤")
	fmt.Println("   │ name:   \"Sohn\"                  │")
	fmt.Printf("   │ vater:  %p ───┐     │\n", p2.vater)
	fmt.Println("   │ mutter: nil                     │")
	fmt.Println("   └─────────────────────────────────┘")
	fmt.Println("            │")
	fmt.Println("            └──> zeigt auf p1 (Papa)")
	fmt.Println()

	fmt.Println("   MERKE:")
	fmt.Println("   ✓ Zeiger speichert nur Adresse (z.B. 8 Bytes)")
	fmt.Println("   ✓ Nicht der komplette Verbund wird gespeichert")
	fmt.Println("   ✓ Daher ist rekursive Definition möglich")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Zusammenfassung
	// -------------------------------------------------------------------------

	fmt.Println("9. Zusammenfassung:")
	fmt.Println("   ----------------")
	fmt.Println()

	fmt.Println("   UNZULÄSSIG:")
	fmt.Println("   ✗ type T struct { feld T }  // Direkter Selbstbezug")
	fmt.Println("   ✗ Würde zu unendlicher Rekursion führen")
	fmt.Println("   ✗ Speichergröße wäre unendlich")
	fmt.Println()

	fmt.Println("   ZULÄSSIG:")
	fmt.Println("   ✓ type T struct { feld *T }  // Zeiger auf sich selbst")
	fmt.Println("   ✓ Zeiger haben feste Größe")
	fmt.Println("   ✓ nil beendet Rekursion")
	fmt.Println()

	fmt.Println("   ANWENDUNGEN:")
	fmt.Println("   ✓ Stammbäume (Eltern-Kind-Beziehungen)")
	fmt.Println("   ✓ Verkettete Listen (Datenstruktur)")
	fmt.Println("   ✓ Binäre Bäume (Hierarchien)")
	fmt.Println("   ✓ Graphen (Netzwerke)")
	fmt.Println()

	fmt.Println("   WICHTIG:")
	fmt.Println("   ✓ Immer auf nil prüfen vor Dereferenzierung")
	fmt.Println("   ✓ Rekursive Algorithmen für Verarbeitung nötig")
	fmt.Println("   ✓ Kapitel 9 behandelt rekursive Algorithmen")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// inOrder führt eine In-Order Traversierung des Baums durch (rekursiv)
func inOrder(k *baumKnoten) {
	if k == nil {
		return
	}
	inOrder(k.links)          // Links besuchen
	fmt.Printf("%d ", k.wert) // Wurzel verarbeiten
	inOrder(k.rechts)         // Rechts besuchen
}
