// Von der verketteten Liste zur Blockchain
//
// Das Programm gehört zur Lektion „Von der verketteten Liste zur Blockchain“
// im Mentoriat Imperative Programmierung. Es ist ein Modell: kein echtes Geld
// und keine echte Blockchain. Es verwendet nur, was im Kurstext vorkommt:
// Verbünde, Arrays, Zeiger, Funktionen mit zwei Rückgabewerten, uint32 und
// uint64. Neu ist nur fmt.Sprintf. Die Funktion arbeitet wie fmt.Printf,
// gibt den Text aber zurück, statt ihn auszugeben.
//
// Ausführen: go run blockchain.go

package main

import "fmt"

// Ein Konto mit Namen, Guthaben und Sperrvermerk
type konto struct {
	name     string
	guthaben uint64
	gesperrt bool
}

// Ein Block der Kette. Er kennt seinen Vorgänger auf zwei Arten:
// über den Zeiger vorgänger und über dessen Hashwert vorgängerHash.
type block struct {
	nummer        int
	eintrag       string
	vorgängerHash uint32
	hash          uint32
	vorgänger     *block
}

// Die Blockkette mit den Konten. kopf zeigt auf den neuesten Block.
type blockkette struct {
	kopf        *block
	konten      [3]konto
	herausgeber string
	angehalten  bool
}

// fingerabdruckText berechnet aus einem Text einen Hashwert mit 32 Bit.
// Das Verfahren heißt FNV-1a. Der Überlauf bei der Multiplikation ist gewollt.
func fingerabdruckText(text string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(text); i++ {
		h = h ^ uint32(text[i]) // das Byte mit XOR einmischen
		h = h * 16777619        // Überlauf: Go rechnet modulo 2^32
	}
	return h
}

// fingerabdruck berechnet den Hashwert eines Blocks aus seiner Nummer,
// seinem Eintrag und dem Hashwert seines Vorgängers.
func fingerabdruck(b *block) uint32 {
	text := fmt.Sprintf("%d|%s|%08x", b.nummer, b.eintrag, b.vorgängerHash)
	return fingerabdruckText(text)
}

// blockAnhängen hängt einen neuen Block vorn an die Kette.
func blockAnhängen(k *blockkette, eintrag string) {
	neuer := &block{eintrag: eintrag}
	if k.kopf != nil {
		neuer.nummer = k.kopf.nummer + 1
		neuer.vorgängerHash = k.kopf.hash
	}
	neuer.hash = fingerabdruck(neuer)
	neuer.vorgänger = k.kopf
	k.kopf = neuer
}

// ketteGültig prüft jeden Block, vom neuesten bis zum ersten.
func ketteGültig(k *blockkette) bool {
	for b := k.kopf; b != nil; b = b.vorgänger {
		if b.hash != fingerabdruck(b) {
			return false
		}
		if b.vorgänger != nil && b.vorgängerHash != b.vorgänger.hash {
			return false
		}
	}
	return true
}

// kontoSuchen liefert einen Zeiger auf das Konto mit dem Namen oder nil.
func kontoSuchen(k *blockkette, name string) *konto {
	for i := 0; i < len(k.konten); i++ {
		if k.konten[i].name == name {
			return &k.konten[i]
		}
	}
	return nil
}

// überweisen prüft erst alle Regeln und bucht dann.
func überweisen(k *blockkette, von string, an string, betrag uint64) (bool, string) {
	sender := kontoSuchen(k, von)
	empfänger := kontoSuchen(k, an)
	if sender == nil || empfänger == nil {
		return false, "unbekanntes Konto"
	}
	if k.angehalten {
		return false, "die Kette ist angehalten"
	}
	if sender.gesperrt {
		return false, von + " ist gesperrt"
	}
	if empfänger.gesperrt {
		return false, an + " ist gesperrt"
	}
	if betrag > sender.guthaben {
		return false, "das Guthaben reicht nicht"
	}
	sender.guthaben = sender.guthaben - betrag
	empfänger.guthaben = empfänger.guthaben + betrag
	blockAnhängen(k, fmt.Sprintf("%s an %s: %d", von, an, betrag))
	return true, "gebucht"
}

// ausgeben erzeugt neue Coins. Das darf nur der Herausgeber.
func ausgeben(k *blockkette, aufrufer string, an string, betrag uint64) (bool, string) {
	if aufrufer != k.herausgeber {
		return false, "das darf nur der Herausgeber"
	}
	empfänger := kontoSuchen(k, an)
	if empfänger == nil {
		return false, "unbekanntes Konto"
	}
	if empfänger.guthaben+betrag < empfänger.guthaben {
		return false, "Überlauf verhindert"
	}
	empfänger.guthaben = empfänger.guthaben + betrag
	blockAnhängen(k, fmt.Sprintf("%s an %s: %d", k.herausgeber, an, betrag))
	return true, "ausgegeben"
}

// sperren sperrt ein Konto. Das darf nur der Herausgeber.
func sperren(k *blockkette, aufrufer string, name string) (bool, string) {
	if aufrufer != k.herausgeber {
		return false, "das darf nur der Herausgeber"
	}
	betroffen := kontoSuchen(k, name)
	if betroffen == nil {
		return false, "unbekanntes Konto"
	}
	betroffen.gesperrt = true
	blockAnhängen(k, k.herausgeber+" sperrt "+name)
	return true, "gesperrt"
}

// vernichten setzt das Guthaben eines gesperrten Kontos auf 0.
// Das darf nur der Herausgeber.
func vernichten(k *blockkette, aufrufer string, name string) (bool, string) {
	if aufrufer != k.herausgeber {
		return false, "das darf nur der Herausgeber"
	}
	betroffen := kontoSuchen(k, name)
	if betroffen == nil {
		return false, "unbekanntes Konto"
	}
	if !betroffen.gesperrt {
		return false, name + " ist nicht gesperrt"
	}
	betrag := betroffen.guthaben
	betroffen.guthaben = 0
	blockAnhängen(k, fmt.Sprintf("%s vernichtet %d von %s", k.herausgeber, betrag, name))
	return true, "vernichtet"
}

// anhalten hält alle Überweisungen an. Das darf nur der Herausgeber.
func anhalten(k *blockkette, aufrufer string) (bool, string) {
	if aufrufer != k.herausgeber {
		return false, "das darf nur der Herausgeber"
	}
	k.angehalten = true
	blockAnhängen(k, k.herausgeber+" hält die Kette an")
	return true, "angehalten"
}

// zeige gibt eine Aktion und ihr Ergebnis aus.
func zeige(aktion string, ok bool, meldung string) {
	fmt.Printf("%-38s %-5v %s\n", aktion, ok, meldung)
}

// kontenAusgeben gibt alle Konten aus.
func kontenAusgeben(k *blockkette) {
	for i := 0; i < len(k.konten); i++ {
		kt := k.konten[i]
		fmt.Printf("  %-8s %4d  gesperrt: %v\n", kt.name, kt.guthaben, kt.gesperrt)
	}
}

// ketteAusgeben gibt alle Blöcke aus, vom neuesten bis zum ersten.
func ketteAusgeben(k *blockkette) {
	for b := k.kopf; b != nil; b = b.vorgänger {
		fmt.Printf("  Block %d  vorgängerHash %08x  hash %08x  %s\n", b.nummer, b.vorgängerHash, b.hash, b.eintrag)
	}
}

func main() {
	k := &blockkette{herausgeber: "Herausgeber"}
	k.konten[0] = konto{name: "Alice"}
	k.konten[1] = konto{name: "Bob"}
	k.konten[2] = konto{name: "Charlie"}

	ok, meldung := ausgeben(k, "Herausgeber", "Alice", 100)
	zeige("Herausgeber gibt Alice 100", ok, meldung)
	ok, meldung = ausgeben(k, "Herausgeber", "Bob", 30)
	zeige("Herausgeber gibt Bob 30", ok, meldung)
	ok, meldung = überweisen(k, "Alice", "Bob", 25)
	zeige("Alice überweist Bob 25", ok, meldung)
	ok, meldung = überweisen(k, "Bob", "Charlie", 10)
	zeige("Bob überweist Charlie 10", ok, meldung)
	ok, meldung = überweisen(k, "Charlie", "Alice", 50)
	zeige("Charlie überweist Alice 50", ok, meldung)
	kontenAusgeben(k)

	ok, meldung = sperren(k, "Bob", "Alice")
	zeige("Bob sperrt Alice", ok, meldung)
	ok, meldung = sperren(k, "Herausgeber", "Alice")
	zeige("Herausgeber sperrt Alice", ok, meldung)
	ok, meldung = überweisen(k, "Alice", "Charlie", 5)
	zeige("Alice überweist Charlie 5", ok, meldung)
	ok, meldung = vernichten(k, "Herausgeber", "Alice")
	zeige("Herausgeber vernichtet Alices Guthaben", ok, meldung)
	ok, meldung = anhalten(k, "Herausgeber")
	zeige("Herausgeber hält die Kette an", ok, meldung)
	ok, meldung = überweisen(k, "Bob", "Charlie", 5)
	zeige("Bob überweist Charlie 5", ok, meldung)
	kontenAusgeben(k)

	fmt.Println()
	ketteAusgeben(k)
	fmt.Println("Kette gültig:", ketteGültig(k))

	// Jemand ändert den Eintrag in Block 2
	b := k.kopf
	for b.nummer != 2 {
		b = b.vorgänger
	}
	b.eintrag = "Alice an Bob: 5"
	fmt.Println("Block 2 geändert, Kette gültig:", ketteGültig(k))
}
