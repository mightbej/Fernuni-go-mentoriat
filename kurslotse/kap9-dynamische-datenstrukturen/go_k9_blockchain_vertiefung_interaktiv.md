# Vertiefung: Blockchain in Go

Die Seite ist die Vertiefung zur Lektion „Blockchain: eine verkettete Liste mit Fingerabdruck“, ein Exkurs zu Kapitel 9. Die leichte Fassung zeigt das Gleiche ohne Rechnung Byte für Byte und ohne die Funktionen im Wortlaut. Sie baut eine kleine Blockchain mit dem Stoff des Kurses in Go nach: eine verkettete Liste, in der jeder neue Block vorn ankommt und auf den älteren zeigt, eine Hashfunktion mit 32 Bit, Guthaben im Typ `uint64`, Überweisungen zwischen drei Konten und die Rechte eines Herausgebers. Das ganze Programm liegt als `blockchain.go` im Ordner `go` neben der Seite.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Blockchain, Hashfunktion, FNV-1a, Bitcoin, Ethereum und Solidity kommen im Kurstext nicht vor. In Abschnitt 4.3.8 wird die Hashfunktion nur in einem Satz als Beispiel für gewollten Überlauf genannt. Die Seite ist ein Exkurs des Mentors.
- Das Programm ist ein Modell: kein echtes Geld und keine echte Blockchain. Es läuft auf einem Rechner. Netzwerk, digitale Signaturen und Proof of Work fehlen.
- Aus dem Kurstext stammen: die verkettete Liste und der Güterzug aus Abschnitt 9.2 mit Listing 9.1, die vorzeichenlosen Typen und Hinweis 4.7 aus Abschnitt 4.3.6, XOR aus Abschnitt 4.3.7.2, der Überlauf aus Abschnitt 4.3.8, der Index-Ausdruck `text[i]` aus Abschnitt 4.6, zwei Rückgabewerte aus Abschnitt 7.1, die Adresse eines Array-Elements `&a[i]` aus Abschnitt 7.4, Verbünde aus Kapitel 8.
- `fmt.Sprintf` steht nicht im Kurstext. Die Funktion arbeitet wie `fmt.Printf` aus Abschnitt 4.2, gibt den Text aber zurück, statt ihn auszugeben.
- Die Typen des fertigen Programms: `konto` mit `name string`, `guthaben uint64`, `gesperrt bool`. `block` mit `nummer int`, `eintrag string`, `vorgängerHash uint32`, `hash uint32`, `vorgänger *block`. `blockkette` mit `kopf *block`, `konten [3]konto`, `herausgeber string`, `angehalten bool`. In Schritt 1 hat `block` nur `nummer` und `vorgänger`, `blockkette` nur `kopf`.
- `kopf` zeigt immer auf den neuesten Block. Jeder Block zeigt mit `vorgänger` auf den älteren Block. Der erste Block hat die Nummer 0 und `vorgänger` gleich `nil`. In echten Blockchains heißt er Genesis-Block.
- Hashwerte stehen auf der Seite hexadezimal mit acht Ziffern, wie `fmt.Printf` mit `%08x` sie ausgibt.
- Die Seite verwendet keine Slices und keine Maps. Der Kurs arbeitet mit Arrays. Echte Blockchains verwenden SHA-256 mit 256 Bit. Das Paket `crypto/sha256` verlangt den Text als Slice und kommt deshalb nicht vor.
- Auf der Seite läuft kein Go. Der Browser rechnet nach, was das Programm tut. Alle Werte und Ausgaben auf der Seite sind mit Go 1.26 geprüft.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität. Die Blockchain gehört nicht zum Stoff des Kurstexts.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text sagt, was eine Blockchain ist: ein Kassenbuch, das als Kopie auf vielen Rechnern liegt, mit Einträgen in verketteten Blöcken. Bekannt wurde sie 2009 mit Bitcoin.
- Zwei Änderungen gegenüber der verketteten Liste: Die Richtung, ein neuer Block kommt vorn an und zeigt auf den älteren. Und der Fingerabdruck, jeder Block trägt den Hashwert seines Vorgängers.
- Ein gelber Kasten sagt, dass Blockchain und Hashfunktion nicht im Kurstext stehen.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Verkettete Listen in Go“, „Überlauf bei Ganzzahlen in Go“, „Rechenwerk“ für XOR, „Strings: Bytes und Zeichen“, „Funktionen in Go“ und die Programmierübung „Ein Telefonbuch als Verbund-Array“.
- Unter „Die Schritte“ steht je ein Satz zu den sechs Schritten. Darunter steht ein Link auf `blockchain.go`.
- Ein Link führt zur Lektion „Vom Eisenbahnzug in den Maschinenraum von Tether“. Sie zeigt die echten Blockchains und den Vertrag des Stablecoins USDT.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.2

## Schritt 1: Der umgedrehte Güterzug

Reiter: zug

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text vergleicht mit dem Güterzug des Kurstexts. Dort zeigt die Lokomotive auf den ersten Wagon und jeder Wagon auf den nächsten, die Reihenfolge ergibt sich aus den Containernummern. In der Blockchain ergibt sich die Reihenfolge aus der Zeit: Jeder neue Block kommt direkt hinter die Lokomotive, seine Kupplung zeigt auf den älteren Block.
- Ein Codeblock zeigt die Typen `block` mit `nummer int` und `vorgänger *block` und `blockkette` mit `kopf *block`.
- Zwei Codeblöcke nebeneinander: links `blockAnhängen(k *blockkette)`, rechts der erste Fall von `elementEinfügen` aus Listing 9.1 mit `neues.nächstes = liste.kopf` und `liste.kopf = neues`. Die letzten zwei Zeilen von `blockAnhängen` sind die gleichen zwei Zuweisungen: `neuer.vorgänger = k.kopf` und `k.kopf = neuer`.
- `blockAnhängen` braucht keine Schleife, weil ein neuer Block nie in der Mitte eingefügt wird.
- Die Demo hat fünf Zeilen aus `main`: `k := &blockkette{}` und viermal `blockAnhängen(k)`. Eine Anweisung der Demo ist ein ganzer Aufruf.
- Die Demo zeigt die Kette wie die Liste in Listing 9.1: vorn ein violetter Kasten „kopf“, dahinter der neueste Block, ganz hinten der älteste und `nil`. Jeder Pfeil ist ein Zeiger `vorgänger`. Unter jeder Nummer steht „ab Zelle“ und die Anfangsadresse. Darunter stehen die Rahmen mit den Variablen, der Hauptspeicher und ein Kasten mit einem Satz. Gelb ist der neue Block.
- Im Modell belegen ein `int` und ein Zeiger je 8 Zellen. Ein `block` belegt 16 Zellen, eine `blockkette` 8. `k` liegt in den Zellen 100 bis 107, die `blockkette` ab Zelle 120. Die Blöcke 0 bis 3 entstehen ab den Zellen 200, 216, 232 und 248. Die Adressen sind erfunden.
- Am Ende steht im Bild: kopf → Block 3 → Block 2 → Block 1 → Block 0 → nil. Ein neuer Block kommt immer direkt hinter kopf. Vom Kopf aus gelesen geht es in der Zeit zurück.
- Jeder Block bleibt erreichbar. Vom Kopf aus führt der Zeiger `vorgänger` von Block zu Block bis Block 0. Beim Anhängen geht kein Block verloren, weil `neuer.vorgänger` den alten Wert von `kopf` bekommt, bevor `kopf` überschrieben wird.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Programm und seine Werte lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 9.2
- 8.3.1

## Schritt 2: Der Fingerabdruck

Reiter: hash

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text erklärt Hashfunktion und Hashwert und nennt zwei Eigenschaften: gleicher Text, gleicher Wert, und ein geändertes Zeichen ändert den ganzen Wert.
- Ein Codeblock zeigt `fingerabdruckText(text string) uint32`. Das Verfahren heißt FNV-1a. Startwert ist 2166136261, hexadezimal 811c9dc5. Je Byte zwei Anweisungen: `h = h ^ uint32(text[i])` und `h = h * 16777619`.
- Eine Liste erklärt: `uint32`, weil ein Hashwert ein Bitmuster ist (Abschnitt 4.3.6). `text[i]` ist ein Byte, ein Umlaut besteht aus zwei Bytes (Abschnitt 4.6). `^` ist XOR (Abschnitt 4.3.7.2). Das Produkt braucht bis zu 57 Bit, Go behält die unteren 32. Diesen Überlauf setzt man bei einer Hashfunktion bewusst ein (Abschnitt 4.3.8).
- Vor der Demo führt ein Absatz das Beispiel ein: Die Texte sind Einträge eines Kassenbuchs. Alice, Bob und Charlie haben je ein Konto, bezahlt wird in Coins. „Alice an Bob: 25“ heißt: Alice überweist Bob 25 Coins. Für die Hashfunktion ist ein Eintrag nur eine Folge von Bytes.
- Eine Liste erklärt die Spalten der Tabelle: Eine Zeile ist ein Durchlauf der Schleife. „Byte“ ist `text[i]` mit Zeichen. „h nach XOR“ ist `h` nach der ersten Anweisung. „Produkt ohne Grenze“ ist das mathematische Ergebnis der Multiplikation. „h neu“ sind die unteren 32 Bit davon, mit diesem Wert beginnt die nächste Zeile. `h` in der letzten Zeile ist der Hashwert.
- Die Demo hat sechs Vorlagen als Knöpfe: „Alice an Bob: 25“, „Alice an Bob: 5“, „alice an Bob: 25“, „Hallo“, „Grüße“, „Alice an Bob: 9403879403“. Darunter ein Eingabefeld für einen eigenen Text mit höchstens 40 Bytes.
- Ein Kasten nennt die Zahl der Bytes und den Hashwert hexadezimal und dezimal.
- Die Tabelle hat die Spalten i, Byte, h nach XOR, Produkt ohne Grenze und h neu. Die erste Zeile ist der Startwert.
- Ein gelber Kasten sagt, dass 32 Bit zu wenig sind: Es gibt nur rund 4,3 Milliarden Werte vom Typ `uint32`. Ein Go-Programm hat auf einem gewöhnlichen Rechner in knapp zehn Minuten den Text „Alice an Bob: 9403879403“ gefunden. Er hat den gleichen Hashwert 4e3c6e66 wie „Alice an Bob: 25“. Echte Blockchains verwenden 256 Bit, Bitcoin zum Beispiel SHA-256.
- FNV-1a ist ein veröffentlichtes Verfahren von Fowler, Noll und Vo. Es dient in der Praxis für Hashtabellen und nicht für Sicherheit.

### Handgriffe

- Eine der sechs Vorlagen wählen.
- Einen eigenen Text eingeben, höchstens 40 Bytes.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 4.3.8
- 4.3.6
- 4.3.7.2
- 4.6

## Schritt 3: Hash statt Zeiger

Reiter: kette

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text: Eine Adresse gilt nur in einem laufenden Programm auf einem Rechner. Den Hashwert berechnet jeder Rechner aus dem gleichen Inhalt gleich. Deshalb trägt jeder Block den Hashwert seines Vorgängers.
- Ein Codeblock zeigt den erweiterten Typ `block` mit `nummer`, `eintrag`, `vorgängerHash`, `hash` und `vorgänger`. Eine Liste erklärt die drei neuen Felder. `hash` hängt über `vorgängerHash` vom Vorgänger ab und damit von allen älteren Blöcken.
- Ein zweiter Codeblock zeigt `fingerabdruck(b *block) uint32` mit `fmt.Sprintf("%d|%s|%08x", b.nummer, b.eintrag, b.vorgängerHash)`, die volle Fassung von `blockAnhängen(k *blockkette, eintrag string)` und `ketteGültig(k *blockkette) bool`. `ketteGültig` läuft vom Kopf zum ersten Block und prüft je Block zwei Dinge: Passt `hash` zum Inhalt? Passt `vorgängerHash` zum Hashwert des Vorgängers? Bei der ersten Abweichung gibt die Funktion `false` zurück.
- Für Block 2 entsteht der Text `2|Alice an Bob: 25|2b2ae705`.
- Die Demo zeigt eine Kette mit vier Blöcken als Karten: vorn der Kasten „kopf“, dahinter Block 3, ganz hinten Block 0 und `nil`. Jeder Pfeil ist ein Zeiger auf den Vorgänger. Im Feld `vorgängerHash` steht deshalb der Hashwert des Blocks rechts daneben, und unter dem Feld steht auf jeder Karte „= hash von Block …, rechts daneben“. Je Karte: Nummer, Eintrag, `hash` und `vorgängerHash`. Adressen zeigen die Karten nicht.
- Die ungefälschte Kette: Block 0 „Herausgeber an Alice: 100“ mit hash f34cbfc8. Block 1 „Herausgeber an Bob: 30“ mit hash 2b2ae705. Block 2 „Alice an Bob: 25“ mit hash bfd83c2e. Block 3 „Bob an Charlie: 10“ mit hash 13892aec.
- Vier Stufen der Fälschung als Knöpfe. Alice will in Block 2 nur 5 Coins überwiesen haben. Gelb ist ein geändertes Feld, rot das Feld, an dem `ketteGültig` anschlägt.
- Stufe „Eintrag in Block 2 geändert“: Block 2 hat den Eintrag „Alice an Bob: 5“ und weiter hash bfd83c2e. `fingerabdruck` ergibt 0ad88e7e. Die Prüfung schlägt an Block 2 an.
- Stufe „dazu hash von Block 2 neu berechnet“: Block 2 hat hash 0ad88e7e. Block 3 trägt in `vorgängerHash` weiter bfd83c2e. Die Prüfung schlägt an Block 3 an.
- Stufe „dazu Block 3 neu berechnet“: Block 3 hat `vorgängerHash` 0ad88e7e und hash 16bf5383. `ketteGültig` gibt `true` zurück. Die anderen Rechner haben aber eine Kopie mit dem Kopf 13892aec und übernehmen die gefälschte Kette nicht.
- Ein gelber Kasten nennt Proof of Work bei Bitcoin: Für jeden Block muss ein Rechner eine teure Rechenaufgabe lösen. Ein Fälscher müsste das für jeden späteren Block wiederholen, schneller als alle anderen Rechner zusammen. Die Seite zeigt das Verfahren nicht.

### Handgriffe

- Eine der vier Stufen wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Eintrag kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 9.2
- 4.2

## Schritt 4: Guthaben ohne Vorzeichen

Reiter: guthaben

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt Hinweis 4.7 des Kurstexts: Vorzeichenlose Typen sollen nicht für Werte verwendet werden, die nur scheinbar positiv sind. Das Beispiel dort ist `10 − 15` im Typ `uint8` mit dem Ergebnis 251. Ein Link führt zum Zahlenkreis, Schritt 3, wo der Sprung über die Null zu sehen ist.
- Blockchains verwenden trotzdem vorzeichenlose Typen. Verträge in Solidity für Ethereum speichern Guthaben meist im Typ `uint256`. Deshalb muss jedes Programm vor dem Abziehen prüfen.
- Ein Codeblock zeigt die Prüfung `if betrag > sender.guthaben` vor `sender.guthaben = sender.guthaben - betrag`.
- Die Demo: Alice hat 75 Coins im Typ `uint64`. Vier Knöpfe wählen den Betrag: 25, 75, 76 oder 100. Zwei schwarze Felder nebeneinander zeigen das Ergebnis „Ohne Prüfung“ und „Mit Prüfung“. Ein Kasten darunter erklärt das Ergebnis.
- Mit Go geprüft: 75 − 76 ergibt 18446744073709551615, das ist 2^64 − 1, die größte Zahl des Typs. 75 − 100 ergibt 18446744073709551591.
- Unter „Auch beim Addieren prüfen“ steht die Prüfung aus `ausgeben`: `if empfänger.guthaben+betrag < empfänger.guthaben`. Ist die Summe kleiner als der alte Wert, ist sie übergelaufen. Am Ende von Abschnitt 4.3.8 steht, dass man vor einer gefährdeten Addition prüfen sollte.

### Handgriffe

- Einen der vier Beträge wählen.
- Die „Auflösung“ aufklappen.
- Das Guthaben 75 lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 4.3.6
- 4.3.8

## Schritt 5: Überweisen

Reiter: ueberweisen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `konto`, die volle `blockkette` und `kontoSuchen`. `kontoSuchen` liefert mit `&k.konten[i]` die Adresse eines Array-Elements oder `nil`. Über den Zeiger ändert `überweisen` das Konto im Array selbst und keine Kopie.
- Die Funktion `überweisen` steht mit 23 nummerierten Zeilen da. Sie prüft der Reihe nach: unbekanntes Konto, angehaltene Kette, gesperrter Sender, gesperrter Empfänger, zu wenig Guthaben. Erst dann bucht sie und hängt einen Block mit dem Eintrag „von an an: betrag“ an. Sie gibt `bool` und eine Meldung zurück. Gelb ist die Zeile, mit der der letzte Aufruf geendet hat.
- Der Simulator beginnt mit den vier Blöcken aus Schritt 3. Alice hat 75, Bob 45, Charlie 10.
- Der Simulator zeigt die drei Konten, ob die Kette angehalten ist, das Ergebnis von `ketteGültig`, `kopf` und die neuesten fünf Blöcke als Karten mit Eintrag und Hashwert, der neueste vorn, und einen Verlauf aller Aufrufe.
- Der Simulator rechnet wie die Funktionen in `blockchain.go`. Schritt 5 und Schritt 6 teilen sich seinen Stand.
- Sind Sender und Empfänger das gleiche Konto, bucht `überweisen` ab und wieder auf. Das Guthaben bleibt gleich, die Kette bekommt trotzdem einen Block.

### Handgriffe

- „Von“: Alice, Bob oder Charlie. „An“: Alice, Bob oder Charlie. „Betrag“: 5, 10, 25, 50 oder 100.
- „überweisen“ ruft die Funktion auf. „von vorn“ setzt den Simulator auf den Anfang zurück.
- Die „Auflösung“ aufklappen.
- Einen eigenen Betrag kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 7.1
- 7.4
- 8.2.4

## Schritt 6: Wer darf sperren?

Reiter: sperren

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Kette hat einen Herausgeber. Nur er darf neue Coins ausgeben, Konten sperren, das Guthaben eines gesperrten Kontos vernichten und die Kette anhalten. Jede dieser Funktionen prüft zuerst `aufrufer != k.herausgeber`.
- Ein Codeblock zeigt `sperren` und `vernichten`. `anhalten` setzt `k.angehalten` auf `true`. Danach lehnt `überweisen` jede Überweisung ab.
- Jeder Eingriff des Herausgebers wird ein Block der Kette, zum Beispiel „Herausgeber sperrt Alice“ oder „Herausgeber vernichtet 75 von Alice“.
- Ein gelber Kasten: Im Modell kann jeder Aufruf „Herausgeber“ als Aufrufer angeben. In einer echten Blockchain belegt eine digitale Signatur, wer aufruft.
- Der Simulator setzt den Stand aus Schritt 5 fort, mit dem gleichen Bild wie dort.
- Unter „Was das Programm zeigt“: Die Kette ist technisch in Ordnung, und trotzdem entscheidet eine einzige Stelle, wer zahlen darf. Solche Funktionen stehen im Vertrag des Stablecoins USDT von Tether. Ein Link führt zur Lektion „Vom Eisenbahnzug in den Maschinenraum von Tether“.
- Unter „Das ganze Programm“ steht ein Link auf `blockchain.go` und die Ausgabe des Programms mit Go 1.26. `main` spielt die Schritte 5 und 6 durch und ändert am Ende den Eintrag in Block 2. Die letzten zwei Zeilen der Ausgabe: „Kette gültig: true“ und „Block 2 geändert, Kette gültig: false“.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: eine Funktion `entsperren` schreiben, ein viertes Konto für Dora, und nach dem Ändern von Block 2 alle Hashwerte bis zum Kopf neu berechnen. Zur zweiten Aufgabe: Die Länge des Arrays in `blockkette` und die Zeilen in `main` ändern sich. Zur dritten: `ketteGültig` gibt dann `true` zurück, aber der Kopf hat einen anderen Hashwert als bei den anderen Rechnern.

### Handgriffe

- „Wer ruft auf?“: Herausgeber, Alice, Bob oder Charlie. „Welches Konto?“: Alice, Bob oder Charlie.
- „sperren“, „vernichten“ und „anhalten“ rufen die Funktion mit dem gewählten Aufrufer auf.
- „überweisen“ überweist 10 Coins vom gewählten Konto an Charlie, bei Charlie an Alice.
- „von vorn“ setzt den Simulator zurück.
- Eine Funktion zum Entsperren oder Freigeben gibt es im Programm nicht. Sie ist die erste Aufgabe unter „Zum Weiterdenken“.

### Abschnitte des Kurstexts

- 7.4.1
- 8.2.4

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Der umgedrehte Güterzug, 9.2, Seiten 229–236: Güterzug und Listing 9.1: Die Kette mit kopf und Zeigern folgt der Liste aus Listing 9.1. Die Typen und die Funktion sind eigene Beispiele, geprüft mit Go 1.26.
  - 2 Der Fingerabdruck, 4.3.6, Seiten 66–68; 4.3.7, Seiten 69–78; 4.3.8, Seiten 78–84: Die Ganzzahl-Typen, die Operatoren auf Bit-Ebene und der Überlauf stammen aus dem Kurstext. FNV-1a ist ein veröffentlichtes Verfahren von Fowler, Noll und Vo. Die Funktion ist ein eigenes Beispiel, geprüft mit Go 1.26.
  - 3 Hash statt Zeiger, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Blockchain und die Funktionen sind eigene Beispiele, geprüft mit Go 1.26.
  - 4 Guthaben ohne Vorzeichen, 4.3.6, Seiten 66–68: Hinweis 4.7; 4.3.8, Seiten 78–84: Hinweis 4.7 und der Überlauf stammen aus dem Kurstext. Die Werte sind mit Go 1.26 geprüft.
  - 5 Überweisen, 7.1, Seiten 163–168: zwei Rückgabewerte; 7.4, Seiten 173–178: Adresse eines Array-Elements: Zwei Rückgabewerte und die Adresse eines Array-Elements stammen aus dem Kurstext. Die Funktionen sind eigene Beispiele, geprüft mit Go 1.26.
  - 6 Wer darf sperren?, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Funktionen und das Programm sind eigene Beispiele, geprüft mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 9.2, Seiten 229–236, verkettete Liste: Lektion: Verkettete Listen in Go.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Lektion: Überlauf bei Ganzzahlen in Go.
  - Abschnitt 4.3.7.2.1, Seiten 72–74, bit-weise Operatoren: ^ in der Hashfunktion. Lektion: Rechenwerk.
  - Abschnitt 7.1, Seiten 163–165, zwei Rückgabewerte: Lektion: Funktionen in Go.
  - Ohne Abschnitt im Kurstext, Tether: Wie echtes digitales Geld mit den gleichen Bausteinen arbeitet, zeigt der Exkurs Vom Eisenbahnzug in den Maschinenraum von Tether.
- Unter „Was über den Kurstext hinausgeht“ steht:
  - Die Blockchain: Sie kommt im Kurstext nicht vor.
- Unter „Wie es weitergeht“ stehen Links auf die Lektion zu Tether und auf „Binäre Suchbäume in Go“.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 6“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.3.6
- 4.3.7
- 4.3.7.2.1
- 4.3.8
- 7.1
- 7.4
- 9.2
