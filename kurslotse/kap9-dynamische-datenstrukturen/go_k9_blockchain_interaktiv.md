# Blockchain: eine verkettete Liste mit Fingerabdruck

Die Seite ist ein Exkurs zu Kapitel 9 für Studenten im ersten Semester. Sie zeigt in vier Schritten, dass hinter einer Blockchain Stoff aus dem Kurs steckt: eine verkettete Liste aus Abschnitt 9.2 und vorzeichenlose Ganzzahlen aus Kapitel 4. Die Seite bleibt an der Oberfläche. Wer jede Funktion ausgeschrieben und nachgerechnet sehen will, findet sie in der Lektion „Vertiefung: Blockchain in Go“.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Blockchain, Hashwert, Bitcoin und Fingerabdruck kommen im Kurstext nicht vor. In Abschnitt 4.3.8 wird die Hashfunktion nur in einem Satz als Beispiel für gewollten Überlauf genannt. Die Seite ist ein Exkurs des Mentors.
- Das Beispiel: Alice, Bob und Charlie haben je ein Konto. Bezahlt wird in Coins, der Geldeinheit des Modells. Ein Eintrag wie „Alice an Bob: 25“ heißt: Alice überweist Bob 25 Coins. Der Herausgeber ist die Stelle, die die Coins ausgegeben hat.
- Die Kette steht auf der Seite wie die Liste in Listing 9.1 und in der Lektion „Verkettete Listen in Go“: vorn der Kasten „kopf“, direkt dahinter der neueste Block, ganz hinten der älteste Block und `nil`. Jeder Pfeil ist ein Zeiger `vorgänger`, vom neueren auf den älteren Block. Ein neuer Block kommt immer zwischen `kopf` und den bisher neuesten Block. Vom Kopf aus gelesen geht es in der Zeit zurück. Adressen von Speicherzellen zeigt die Seite nicht.
- Hashwerte stehen hexadezimal mit acht Ziffern da. Die Hashfunktion heißt FNV-1a und rechnet mit `uint32`. Wie sie rechnet, steht auf dieser Seite nicht, nur in der Vertiefung.
- Das Programm ist ein Modell: kein echtes Geld und keine echte Blockchain.
- Auf der Seite läuft kein Go. Alle Werte sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices und keine Maps.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität. Die Blockchain gehört nicht zum Stoff des Kurstexts.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text erklärt Blockchain als gemeinsames Kassenbuch auf vielen tausend Rechnern, ohne Bank, und nennt Bitcoin seit 2009. Er führt Alice, Bob, Charlie und die Coins ein.
- Unter „Was du aus dem Kurs schon kennst“ stehen die verkettete Liste aus Abschnitt 9.2 und die vorzeichenlosen Typen aus Abschnitt 4.3.6, mit Links auf „Verkettete Listen in Go“ und „Zahlenkreis“.
- Unter „Die Schritte“ steht je ein Satz zu den vier Schritten.
- Ein gelber Kasten sagt, dass das Thema nicht im Kurstext steht, und verlinkt die Vertiefung und die Lektion „Vom Eisenbahnzug in den Maschinenraum von Tether“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.2

## Schritt 1: Eine Kette von Blöcken

Reiter: kette

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“. Die Antwort: erst `neuer.vorgänger = k.kopf`, dann `k.kopf = neuer`, wie im ersten Fall von `elementEinfügen` in Listing 9.1.
- Der Text vergleicht mit dem Güterzug des Kurstexts: Ein neuer Block kommt immer ans Ende der Zeit, eingefügt wird nie in der Mitte. Sein Zeiger zeigt auf den Vorgänger. Der erste Block hat `nil`.
- Ein Codeblock zeigt `block` mit `eintrag string` und `vorgänger *block` und `blockkette` mit `kopf *block`.
- Die Demo: Der Knopf „Block anhängen“ hängt nacheinander vier Einträge an: „Herausgeber an Alice: 100“, „Herausgeber an Bob: 30“, „Alice an Bob: 25“, „Bob an Charlie: 10“. Der neue Block ist gelb umrandet. Ein Kasten erklärt, was beim Anhängen mit `vorgänger` und `kopf` geschieht.
- Vom Kopf aus führen die Zeiger bis Block 0. Beim Anhängen geht kein Block verloren.
- Ein Link führt zur Vertiefung, Schritt 1. Dort steht, wie die Blöcke im Hauptspeicher liegen.

### Handgriffe

- „Block anhängen“, höchstens viermal, und „von vorn“.
- Die „Auflösung“ aufklappen.
- Die Einträge lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 9.2

## Schritt 2: Der Fingerabdruck

Reiter: hash

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“: Eine Speicheradresse gilt nur im laufenden Programm auf einem Rechner.
- Der Text führt den Fingerabdruck ein, Fachwort Hashwert: eine Zahl, die ein Programm aus dem Eintrag berechnet. Zwei Eigenschaften: Gleicher Eintrag, gleiche Zahl. Ein Zeichen anders, ganz andere Zahl.
- Die Hashfunktion rechnet mit `uint32` und lässt den Überlauf bewusst zu, wie in Abschnitt 4.3.8 genannt.
- Erste Demo: vier Knöpfe mit Einträgen. Die Hashwerte: „Alice an Bob: 25“ 4e3c6e66, „Alice an Bob: 5“ cd11eb3e. Ein Kasten zeigt den Hashwert des gewählten Eintrags und vergleicht ihn mit „Alice an Bob: 25“.
- Unter „Ein Block merkt sich den Fingerabdruck seines Vorgängers“: Jeder Block trägt seinen eigenen Hashwert in `hash` und den seines Vorgängers in `vorgängerHash`. Der eigene Hashwert wird aus dem Eintrag und aus `vorgängerHash` berechnet.
- Zweite Demo: vier Blöcke als Karten mit Eintrag, `hash` und `vorgängerHash`. Unter `vorgängerHash` steht auf jeder Karte, von welchem Block der Wert stammt, zum Beispiel „= hash von Block 1, rechts daneben“. Bei Block 0 steht „kein Vorgänger“. Zwei Knöpfe: „Kette unverändert“ und „Alice ändert Block 2“. Alice ändert „Alice an Bob: 25“ in „Alice an Bob: 5“ und berechnet den Hashwert von Block 2 neu. Dann ist `vorgängerHash` in Block 3 rot, weil er nicht mehr zum Hashwert von Block 2 passt.
- Die ungefälschten Hashwerte: Block 0 f34cbfc8, Block 1 2b2ae705, Block 2 bfd83c2e, Block 3 13892aec. Nach der Fälschung hat Block 2 den Hashwert 0ad88e7e.
- Um die Fälschung zu verbergen, müsste Alice jeden späteren Block ändern, auf allen Rechnern mit einer Kopie der Kette.
- Ein Link führt zur Vertiefung, Schritt 2 und 3.

### Handgriffe

- Einen der vier Einträge wählen.
- „Kette unverändert“ oder „Alice ändert Block 2“ wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Text kann der Student hier nicht eingeben, das geht in der Vertiefung.

### Abschnitte des Kurstexts

- 4.3.8

## Schritt 3: Guthaben als uint

Reiter: guthaben

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“: 75 − 76 ergibt im Typ `uint64` die Zahl 18446744073709551615.
- Ein Codeblock zeigt `konto` mit `name string`, `guthaben uint64` und `gesperrt bool`.
- Der Text nennt Hinweis 4.7 des Kurstexts: Ergibt eine Rechnung mit einem vorzeichenlosen Typ eine negative Zahl, springt der Wert unbemerkt an das obere Ende des Wertebereichs. Ein Link führt zum Zahlenkreis, Schritt 3.
- Echte Blockchains verwenden trotzdem vorzeichenlose Typen, oft mit 256 Bit. Darum muss jedes Programm vor dem Abziehen prüfen.
- Die Demo: Alice hat 75 Coins. Vier Knöpfe wählen den Betrag 25, 75, 76 oder 100. Zwei Codeblöcke nebeneinander, „Ohne Prüfung“ und „Mit Prüfung“ mit `if betrag > guthaben`, darunter je das Ergebnis. 75 − 100 ergibt ohne Prüfung 18446744073709551591.

### Handgriffe

- Einen der vier Beträge wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 4.3.6
- 4.3.8

## Schritt 4: Regeln als Programm

Reiter: regeln

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `überweisen(k *blockkette, von *konto, an *konto, betrag uint64) bool`: Sie gibt `false` zurück, wenn der Absender gesperrt ist oder sein Guthaben nicht reicht. Sonst bucht sie und hängt mit `fmt.Sprintf` einen Block mit dem Eintrag an. `fmt.Sprintf` arbeitet wie `fmt.Printf` aus Abschnitt 4.2, gibt den Text aber zurück.
- Der Text: Sperren darf nur der Herausgeber. Die Sperre landet als eigener Block in der Kette.
- Die Demo: Am Anfang hat Alice 75, Bob 45, Charlie 10, und die Kette hat die vier Blöcke aus Schritt 2. Vier Knöpfe: „Alice überweist Bob 25“, „Herausgeber sperrt Alice“, „Alice überweist Charlie 10“, „Bob überweist Charlie 100“, dazu „von vorn“. Ein Kasten erklärt den letzten Klick. Darunter stehen die drei Konten und die neuesten vier Blöcke.
- Unter „Was das Modell zeigt“: Eine einzige Stelle entscheidet, wer zahlen darf. Solche Funktionen stehen im Programm des digitalen Dollars USDT, Link auf die Lektion zu Tether.
- Die Funktion auf dieser Seite ist kürzer als in der Vertiefung. Dort sucht `überweisen` die Konten selbst und gibt zusätzlich eine Meldung zurück.
- Am Ende des Schritts unter „Nächste Schritte“ stehen Links auf die Vertiefung mit dem Programm `blockchain.go` und auf „Binäre Suchbäume in Go“.

### Handgriffe

- Die vier Knöpfe in beliebiger Reihenfolge und „von vorn“.
- Die „Auflösung“ aufklappen.
- Eine Funktion zum Entsperren gibt es nicht.

### Abschnitte des Kurstexts

- 5.5.1
- 7.4.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Eine Kette von Blöcken, 9.2, Seiten 229–236: Güterzug und Listing 9.1: Die Kette mit kopf und Zeigern folgt der Liste aus Listing 9.1. Die Typen sind eigene Beispiele, geprüft mit Go 1.26.
  - 2 Der Fingerabdruck, 4.3.8, Seiten 78–84: Überlauf: Der Überlauf, auf dem die Hashfunktion beruht, wird in Abschnitt 4.3.8 behandelt. Die Hashfunktion FNV-1a steht nicht im Kurstext, die Lektion rechnet sie nicht vor. Alle Hashwerte sind mit Go 1.26 berechnet.
  - 3 Guthaben als uint, 4.3.6, Seiten 66–68: Hinweis 4.7; 4.3.8, Seiten 78–84: Hinweis 4.7 und der Überlauf stammen aus dem Kurstext. Die Werte sind mit Go 1.26 geprüft.
  - 4 Regeln als Programm, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Funktion ist ein eigenes Beispiel.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 9.2, Seiten 229–236, verkettete Liste: Lektion: Verkettete Listen in Go.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Lektion: Überlauf bei Ganzzahlen in Go.
  - Vertiefung: Die gleiche Kette mit Hashfunktion, Speichersicht und Sperrliste zeigt die Lektion Vertiefung: Blockchain in Go.
  - Ohne Abschnitt im Kurstext, Tether: Wie echtes digitales Geld mit den gleichen Bausteinen arbeitet, zeigt der Exkurs Vom Eisenbahnzug in den Maschinenraum von Tether.
- Unter „Was über den Kurstext hinausgeht“ steht:
  - Die Blockchain: Sie kommt im Kurstext nicht vor.
  - FNV-1a: Die Hashfunktion ist ein veröffentlichtes Verfahren von Fowler, Noll und Vo.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.3.6
- 4.3.8
- 9.2
