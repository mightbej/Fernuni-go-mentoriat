package main

import "fmt"

// =============================================================================
// DEMONSTRATION: Array-Operationen auf Verbünden (Abschnitt 8.4)
// =============================================================================
//
// Dieses Programm demonstriert wichtige Array-Operationen:
//
// 1. Duplizate finden und zählen
//    - Alle Duplikate identifizieren
//    - Anzahl der Vorkommen zählen
//
// 2. Duplizate entfernen
//    - Array bereinigen
//    - Nur eindeutige Elemente behalten
//
// 3. Array umkehren
//    - Reihenfolge invertieren
//    - In-Place Umkehrung
//
// 4. Minimum und Maximum finden
//    - Kleinsten/Größten Eintrag suchen
//    - Nach verschiedenen Feldern
//
// 5. Weitere nützliche Operationen
//    - Array kopieren
//    - Elemente zählen
//    - Sortierung prüfen
//
// Ausführen mit: go run S240_array_operationen_demo.go
//
// =============================================================================

// person: Verbund-Typ für Telefonbuch-Eintrag
type person struct {
	name          string
	telefonnummer int
	alter         int
}

// duplikatInfo: Speichert Information über Duplikate
type duplikatInfo struct {
	name   string
	anzahl int
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("=== DEMONSTRATION: Array-Operationen auf Verbünden ===\n")

	// -------------------------------------------------------------------------
	// 1. Duplizate finden und zählen
	// -------------------------------------------------------------------------

	fmt.Println("1. Duplizate finden und zählen:")
	fmt.Println("   -----------------------------")
	fmt.Println()

	mitDuplikaten := []person{
		{"Müller", 11111, 35},
		{"Schmidt", 22222, 42},
		{"Müller", 33333, 28},
		{"Weber", 44444, 51},
		{"Schmidt", 55555, 39},
		{"Klein", 66666, 45},
		{"Müller", 77777, 31},
		{"Weber", 88888, 29},
	}

	fmt.Println("   Array mit möglichen Duplikaten (nach Name):")
	for i, p := range mitDuplikaten {
		fmt.Printf("   [%d] %-10s Tel: %d, Alter: %d\n",
			i, p.name, p.telefonnummer, p.alter)
	}
	fmt.Println()

	dupliziert := findeDuplikate(mitDuplikaten)

	if len(dupliziert) > 0 {
		fmt.Println("   Gefundene Duplikate:")
		for _, dup := range dupliziert {
			fmt.Printf("   - '%s' kommt %d mal vor\n", dup.name, dup.anzahl)
		}
		fmt.Println()

		fmt.Println("   Details zu den Duplikaten:")
		for _, dup := range dupliziert {
			fmt.Printf("\n   Alle Einträge für '%s':\n", dup.name)
			for i, p := range mitDuplikaten {
				if p.name == dup.name {
					fmt.Printf("      [%d] Tel: %d, Alter: %d\n",
						i, p.telefonnummer, p.alter)
				}
			}
		}
		fmt.Println()
	} else {
		fmt.Println("   ✓ Keine Duplikate gefunden!")
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 2. Duplizate entfernen
	// -------------------------------------------------------------------------

	fmt.Println("2. Duplizate entfernen:")
	fmt.Println("   --------------------")
	fmt.Println()

	fmt.Println("   Vorher (mit Duplikaten):")
	fmt.Printf("   Anzahl Einträge: %d\n", len(mitDuplikaten))
	for i, p := range mitDuplikaten {
		fmt.Printf("   [%d] %-10s\n", i, p.name)
	}
	fmt.Println()

	bereinigt := entferneDuplikate(mitDuplikaten)

	fmt.Println("   Nachher (ohne Duplikate):")
	fmt.Printf("   Anzahl Einträge: %d (Reduzierung: %d)\n",
		len(bereinigt), len(mitDuplikaten)-len(bereinigt))
	for i, p := range bereinigt {
		fmt.Printf("   [%d] %-10s Tel: %d, Alter: %d\n",
			i, p.name, p.telefonnummer, p.alter)
	}
	fmt.Println()

	fmt.Println("   HINWEIS:")
	fmt.Println("   - Bei Duplikaten wird das erste Vorkommen behalten")
	fmt.Println("   - Duplikat-Erkennung erfolgt nur über das Feld 'name'")
	fmt.Println("   - Telefonnummer und Alter können unterschiedlich sein")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 3. Array umkehren
	// -------------------------------------------------------------------------

	fmt.Println("3. Array umkehren (Reihenfolge invertieren):")
	fmt.Println("   ------------------------------------------")
	fmt.Println()

	original := []person{
		{"Bauer", 11111, 30},
		{"Fischer", 22222, 35},
		{"Klein", 33333, 40},
		{"Meier", 44444, 45},
		{"Schmidt", 55555, 50},
	}

	fmt.Println("   Original-Reihenfolge:")
	printArraySlice(original)
	fmt.Println()

	umgekehrt := arrayUmkehren(original)

	fmt.Println("   Umgekehrte Reihenfolge:")
	printArraySlice(umgekehrt)
	fmt.Println()

	fmt.Println("   Visualisierung:")
	fmt.Print("   ")
	for i := 0; i < len(original); i++ {
		if i > 0 {
			fmt.Print(" <-> ")
		}
		fmt.Printf("[%d]", i)
	}
	fmt.Println()
	fmt.Print("   ")
	for i := 0; i < len(original); i++ {
		if i > 0 {
			fmt.Print("     ")
		}
		fmt.Printf(" ↕ ")
	}
	fmt.Println()
	fmt.Print("   ")
	for i := len(original) - 1; i >= 0; i-- {
		if i < len(original)-1 {
			fmt.Print(" <-> ")
		}
		fmt.Printf("[%d]", i)
	}
	fmt.Println()
	fmt.Println()

	// Alternative: In-Place Umkehrung
	fmt.Println("   Alternative: In-Place Umkehrung (ohne neues Array):")
	zuUmkehren := []person{
		{"Alpha", 111, 20},
		{"Beta", 222, 25},
		{"Gamma", 333, 30},
		{"Delta", 444, 35},
	}

	fmt.Println("   Vorher:")
	printArraySlice(zuUmkehren)
	fmt.Println()

	arrayUmkehrenInPlace(&zuUmkehren)

	fmt.Println("   Nachher:")
	printArraySlice(zuUmkehren)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Minimum und Maximum finden
	// -------------------------------------------------------------------------

	fmt.Println("4. Minimum und Maximum finden:")
	fmt.Println("   ----------------------------")
	fmt.Println()

	personen := []person{
		{"Weber", 55555, 42},
		{"Bauer", 11111, 28},
		{"Schmidt", 99999, 67},
		{"Klein", 22222, 19},
		{"Müller", 77777, 54},
	}

	fmt.Println("   Testdaten:")
	for i, p := range personen {
		fmt.Printf("   [%d] %-10s Tel: %d, Alter: %d\n",
			i, p.name, p.telefonnummer, p.alter)
	}
	fmt.Println()

	// Minimum/Maximum nach Name (alphabetisch)
	minName := findeMinimum(personen, "name")
	maxName := findeMaximum(personen, "name")

	fmt.Println("   Nach Name (alphabetisch):")
	fmt.Printf("   Minimum: '%s' (kleinster Name alphabetisch)\n", minName.name)
	fmt.Printf("   Maximum: '%s' (größter Name alphabetisch)\n", maxName.name)
	fmt.Println()

	// Minimum/Maximum nach Telefonnummer
	minTel := findeMinimum(personen, "telefon")
	maxTel := findeMaximum(personen, "telefon")

	fmt.Println("   Nach Telefonnummer:")
	fmt.Printf("   Minimum: %s hat %d (kleinste Nummer)\n",
		minTel.name, minTel.telefonnummer)
	fmt.Printf("   Maximum: %s hat %d (größte Nummer)\n",
		maxTel.name, maxTel.telefonnummer)
	fmt.Println()

	// Minimum/Maximum nach Alter
	minAlter := findeMinimum(personen, "alter")
	maxAlter := findeMaximum(personen, "alter")

	fmt.Println("   Nach Alter:")
	fmt.Printf("   Minimum: %s ist %d Jahre alt (jüngste Person)\n",
		minAlter.name, minAlter.alter)
	fmt.Printf("   Maximum: %s ist %d Jahre alt (älteste Person)\n",
		maxAlter.name, maxAlter.alter)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. Array kopieren
	// -------------------------------------------------------------------------

	fmt.Println("5. Array kopieren (tiefe Kopie):")
	fmt.Println("   -------------------------------")
	fmt.Println()

	quelle := []person{
		{"Original1", 111, 25},
		{"Original2", 222, 30},
		{"Original3", 333, 35},
	}

	fmt.Println("   Quelle-Array:")
	printArraySlice(quelle)
	fmt.Println()

	kopie := arrayKopieren(quelle)

	fmt.Println("   Kopie-Array:")
	printArraySlice(kopie)
	fmt.Println()

	// Modifiziere Kopie
	kopie[0].name = "Modifiziert"
	kopie[0].alter = 99

	fmt.Println("   Nach Änderung in der Kopie:")
	fmt.Println("   Quelle (unverändert):")
	printArraySlice(quelle)
	fmt.Println()
	fmt.Println("   Kopie (geändert):")
	printArraySlice(kopie)
	fmt.Println()

	fmt.Println("   ✓ Quelle bleibt unverändert -> echte tiefe Kopie!")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. Sortierung prüfen
	// -------------------------------------------------------------------------

	fmt.Println("6. Sortierung prüfen:")
	fmt.Println("   -------------------")
	fmt.Println()

	sortiert := []person{
		{"Bauer", 11111, 30},
		{"Fischer", 22222, 35},
		{"Meier", 33333, 40},
		{"Schmidt", 44444, 45},
		{"Weber", 55555, 50},
	}

	unsortiert := []person{
		{"Weber", 55555, 50},
		{"Bauer", 11111, 30},
		{"Schmidt", 44444, 45},
		{"Meier", 33333, 40},
		{"Fischer", 22222, 35},
	}

	fmt.Println("   Array 1:")
	printArraySlice(sortiert)
	if istSortiert(sortiert) {
		fmt.Println("   ✓ Array ist aufsteigend sortiert!")
	} else {
		fmt.Println("   ✗ Array ist nicht sortiert")
	}
	fmt.Println()

	fmt.Println("   Array 2:")
	printArraySlice(unsortiert)
	if istSortiert(unsortiert) {
		fmt.Println("   ✓ Array ist aufsteigend sortiert!")
	} else {
		fmt.Println("   ✗ Array ist nicht sortiert")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Elemente zählen (mit Bedingung)
	// -------------------------------------------------------------------------

	fmt.Println("7. Elemente zählen (mit Bedingung):")
	fmt.Println("   ---------------------------------")
	fmt.Println()

	gemischt := []person{
		{"Person1", 11111, 25},
		{"Person2", 22222, 42},
		{"Person3", 33333, 18},
		{"Person4", 44444, 67},
		{"Person5", 55555, 31},
		{"Person6", 66666, 19},
		{"Person7", 77777, 55},
	}

	fmt.Println("   Testdaten:")
	for i, p := range gemischt {
		fmt.Printf("   [%d] %-10s Alter: %d\n", i, p.name, p.alter)
	}
	fmt.Println()

	unter30 := zaehleWenn(gemischt, func(p person) bool { return p.alter < 30 })
	zwischen30und50 := zaehleWenn(gemischt, func(p person) bool {
		return p.alter >= 30 && p.alter < 50
	})
	ueber50 := zaehleWenn(gemischt, func(p person) bool { return p.alter >= 50 })

	fmt.Printf("   Alter unter 30:      %d Personen\n", unter30)
	fmt.Printf("   Alter 30-49:         %d Personen\n", zwischen30und50)
	fmt.Printf("   Alter 50 und älter:  %d Personen\n", ueber50)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Array filtern
	// -------------------------------------------------------------------------

	fmt.Println("8. Array filtern (Bedingung erfüllen):")
	fmt.Println("   ------------------------------------")
	fmt.Println()

	fmt.Println("   Original-Array:")
	for i, p := range gemischt {
		fmt.Printf("   [%d] %-10s Alter: %d\n", i, p.name, p.alter)
	}
	fmt.Println()

	erwachsene := filtereArray(gemischt, func(p person) bool { return p.alter >= 18 })
	senioren := filtereArray(gemischt, func(p person) bool { return p.alter >= 60 })

	fmt.Println("   Gefiltert: Nur Erwachsene (>= 18 Jahre):")
	for i, p := range erwachsene {
		fmt.Printf("   [%d] %-10s Alter: %d\n", i, p.name, p.alter)
	}
	fmt.Println()

	fmt.Println("   Gefiltert: Nur Senioren (>= 60 Jahre):")
	if len(senioren) > 0 {
		for i, p := range senioren {
			fmt.Printf("   [%d] %-10s Alter: %d\n", i, p.name, p.alter)
		}
	} else {
		fmt.Println("   (keine Senioren im Array)")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 9. Zwei Arrays verschmelzen
	// -------------------------------------------------------------------------

	fmt.Println("9. Zwei Arrays verschmelzen:")
	fmt.Println("   --------------------------")
	fmt.Println()

	array1 := []person{
		{"Bauer", 11111, 30},
		{"Fischer", 22222, 35},
		{"Meier", 33333, 40},
	}

	array2 := []person{
		{"Schmidt", 44444, 45},
		{"Weber", 55555, 50},
	}

	fmt.Println("   Array 1:")
	printArraySlice(array1)
	fmt.Println()

	fmt.Println("   Array 2:")
	printArraySlice(array2)
	fmt.Println()

	verschmolzen := arrayVerschmelzen(array1, array2)

	fmt.Println("   Verschmolzenes Array:")
	printArraySlice(verschmolzen)
	fmt.Printf("   Gesamtgröße: %d Elemente\n", len(verschmolzen))
	fmt.Println()

	// -------------------------------------------------------------------------
	// 10. Statistiken berechnen
	// -------------------------------------------------------------------------

	fmt.Println("10. Statistiken berechnen:")
	fmt.Println("    ----------------------")
	fmt.Println()

	statistikDaten := []person{
		{"Person1", 11111, 25},
		{"Person2", 22222, 35},
		{"Person3", 33333, 45},
		{"Person4", 44444, 30},
		{"Person5", 55555, 40},
	}

	fmt.Println("   Daten:")
	for i, p := range statistikDaten {
		fmt.Printf("   [%d] %-10s Alter: %d\n", i, p.name, p.alter)
	}
	fmt.Println()

	durchschnitt := berechneAltersDurchschnitt(statistikDaten)
	minAlter2 := findeMinimum(statistikDaten, "alter")
	maxAlter2 := findeMaximum(statistikDaten, "alter")

	fmt.Printf("   Durchschnittsalter: %.1f Jahre\n", durchschnitt)
	fmt.Printf("   Mindestalter:       %d Jahre (%s)\n",
		minAlter2.alter, minAlter2.name)
	fmt.Printf("   Höchstalter:        %d Jahre (%s)\n",
		maxAlter2.alter, maxAlter2.name)
	fmt.Printf("   Altersspanne:       %d Jahre\n",
		maxAlter2.alter-minAlter2.alter)
	fmt.Println()

	// -------------------------------------------------------------------------
	// 11. Zusammenfassung
	// -------------------------------------------------------------------------

	fmt.Println("11. Zusammenfassung der Array-Operationen:")
	fmt.Println("    ---------------------------------------")
	fmt.Println()

	fmt.Println("   Basis-Operationen:")
	fmt.Println("   ✓ Duplizate finden und zählen")
	fmt.Println("   ✓ Duplizate entfernen")
	fmt.Println("   ✓ Array umkehren (mit/ohne neue Kopie)")
	fmt.Println("   ✓ Minimum/Maximum nach Kriterium finden")
	fmt.Println()

	fmt.Println("   Erweiterte Operationen:")
	fmt.Println("   ✓ Array kopieren (tiefe Kopie)")
	fmt.Println("   ✓ Sortierung prüfen")
	fmt.Println("   ✓ Elemente mit Bedingung zählen")
	fmt.Println("   ✓ Array nach Kriterium filtern")
	fmt.Println("   ✓ Zwei Arrays verschmelzen")
	fmt.Println("   ✓ Statistiken berechnen")
	fmt.Println()

	fmt.Println("   Wichtige Konzepte:")
	fmt.Println("   ✓ Arrays als Slice (dynamische Größe)")
	fmt.Println("   ✓ Zeiger für In-Place Operationen")
	fmt.Println("   ✓ Funktionen höherer Ordnung (Filter, Zähler)")
	fmt.Println("   ✓ Vergleich von Verbund-Feldern")
	fmt.Println()

	fmt.Println("=== ENDE DER DEMONSTRATION ===")
}

// =============================================================================
// Funktionen für Array-Operationen
// =============================================================================

// findeDuplikate: Findet alle Duplikate (nach Name) und zählt Vorkommen
func findeDuplikate(arr []person) []duplikatInfo {
	// Map zum Zählen der Vorkommen
	vorkommen := make(map[string]int)

	// Zähle alle Vorkommen
	for _, p := range arr {
		vorkommen[p.name]++
	}

	// Sammle Duplikate (mehr als 1 Vorkommen)
	var duplikate []duplikatInfo
	for name, anzahl := range vorkommen {
		if anzahl > 1 {
			duplikate = append(duplikate, duplikatInfo{name, anzahl})
		}
	}

	return duplikate
}

// entferneDuplikate: Entfernt Duplikate (behält erstes Vorkommen)
func entferneDuplikate(arr []person) []person {
	gesehen := make(map[string]bool)
	var ergebnis []person

	for _, p := range arr {
		if !gesehen[p.name] {
			gesehen[p.name] = true
			ergebnis = append(ergebnis, p)
		}
	}

	return ergebnis
}

// arrayUmkehren: Kehrt Array-Reihenfolge um (erstellt neue Kopie)
func arrayUmkehren(arr []person) []person {
	n := len(arr)
	ergebnis := make([]person, n)

	for i := 0; i < n; i++ {
		ergebnis[i] = arr[n-1-i]
	}

	return ergebnis
}

// arrayUmkehrenInPlace: Kehrt Array-Reihenfolge um (in-place, ohne Kopie)
func arrayUmkehrenInPlace(arr *[]person) {
	n := len(*arr)

	for i := 0; i < n/2; i++ {
		// Tausche Element i mit Element (n-1-i)
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}

// findeMinimum: Findet Element mit kleinstem Wert für gegebenes Feld
func findeMinimum(arr []person, feld string) person {
	if len(arr) == 0 {
		return person{}
	}

	min := arr[0]

	for i := 1; i < len(arr); i++ {
		vergleich := false

		switch feld {
		case "name":
			vergleich = arr[i].name < min.name
		case "telefon":
			vergleich = arr[i].telefonnummer < min.telefonnummer
		case "alter":
			vergleich = arr[i].alter < min.alter
		}

		if vergleich {
			min = arr[i]
		}
	}

	return min
}

// findeMaximum: Findet Element mit größtem Wert für gegebenes Feld
func findeMaximum(arr []person, feld string) person {
	if len(arr) == 0 {
		return person{}
	}

	max := arr[0]

	for i := 1; i < len(arr); i++ {
		vergleich := false

		switch feld {
		case "name":
			vergleich = arr[i].name > max.name
		case "telefon":
			vergleich = arr[i].telefonnummer > max.telefonnummer
		case "alter":
			vergleich = arr[i].alter > max.alter
		}

		if vergleich {
			max = arr[i]
		}
	}

	return max
}

// arrayKopieren: Erstellt tiefe Kopie des Arrays
func arrayKopieren(arr []person) []person {
	kopie := make([]person, len(arr))
	copy(kopie, arr)
	return kopie
}

// istSortiert: Prüft ob Array aufsteigend nach Name sortiert ist
func istSortiert(arr []person) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].name > arr[i+1].name {
			return false
		}
	}
	return true
}

// zaehleWenn: Zählt Elemente, die Bedingung erfüllen
func zaehleWenn(arr []person, bedingung func(person) bool) int {
	anzahl := 0
	for _, p := range arr {
		if bedingung(p) {
			anzahl++
		}
	}
	return anzahl
}

// filtereArray: Gibt neues Array mit Elementen zurück, die Bedingung erfüllen
func filtereArray(arr []person, bedingung func(person) bool) []person {
	var ergebnis []person
	for _, p := range arr {
		if bedingung(p) {
			ergebnis = append(ergebnis, p)
		}
	}
	return ergebnis
}

// arrayVerschmelzen: Verbindet zwei Arrays zu einem
func arrayVerschmelzen(arr1, arr2 []person) []person {
	ergebnis := make([]person, 0, len(arr1)+len(arr2))
	ergebnis = append(ergebnis, arr1...)
	ergebnis = append(ergebnis, arr2...)
	return ergebnis
}

// berechneAltersDurchschnitt: Berechnet durchschnittliches Alter
func berechneAltersDurchschnitt(arr []person) float64 {
	if len(arr) == 0 {
		return 0
	}

	summe := 0
	for _, p := range arr {
		summe += p.alter
	}

	return float64(summe) / float64(len(arr))
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

// printArraySlice: Gibt Slice übersichtlich aus
func printArraySlice(arr []person) {
	for i, p := range arr {
		fmt.Printf("   [%d] %-12s Tel: %d, Alter: %d\n",
			i, p.name, p.telefonnummer, p.alter)
	}
}
