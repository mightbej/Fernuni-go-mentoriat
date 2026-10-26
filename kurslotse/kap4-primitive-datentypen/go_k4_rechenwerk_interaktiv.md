# Rechenwerk: Wie ein Computer mit Bits addiert und subtrahiert

Die Seite zeigt den Weg von „Strom an, Strom aus“ zur Addition zweier Zahlen: ein Gatter, die Gatter auf einem Byte, die Schaltung für eine Stelle mit Übertrag, der Addierer mit acht Stellen, die Subtraktion über das Zweierkomplement. Die Bit-Operatoren von Go sind die Bausteine. Die Seite gehört zu Abschnitt 4.3.7.2.1 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Bit ist 1 oder 0. 1 bedeutet: Die Leitung führt Strom. 0 bedeutet: Sie führt keinen.
- Alle Werte der Seite sind Bytes vom Typ `uint8`, also acht Bits mit den Wertigkeiten 128, 64, 32, 16, 8, 4, 2 und 1. Das Bit ganz rechts hat die Wertigkeit 1.
- Die Bit-Operatoren in Go: `&` ist AND, `|` ist OR, `^` mit zwei Operanden ist XOR, `^` vor einem Operanden ist NOT, `&^` ist AND NOT.
- Im Kurstext werden die Bit-Operatoren in Abschnitt 4.3.7.2.1 behandelt. Der Addierer und die Schaltbilder stehen nicht im Kurstext. Sie sind eine Ergänzung des Mentors.
- Ein Bit, das der Student umschalten kann, ist ein Knopf. Ein Bit mit dem Wert 1 ist farbig unterlegt. Ergebnis-Bits sind grün, Übertrags-Bits rot.
- In jeder Bit-Tabelle steht links der Name der Zeile, in der Mitte stehen die Bits, rechts steht der Wert als Dezimalzahl oder eine Erklärung.
- Auf der Seite läuft kein Go. Die dunklen Codeblöcke zeigen Go-Zeilen mit der Ausgabe hinter `//`. Für die voreingestellten Werte sind sie mit Go 1.26 nachgeprüft.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Schritt 1 bis 5 beginnen mit einem blauen Kasten „Erst vorhersagen, dann ausprobieren“. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Teil „Vom Gatter zum Rechenwerk“ nennt vier Stufen: Gatter, Gatter auf einem Byte, Addierer für eine Stelle, Rechenwerk aus acht solchen Addierern.
- Eine Liste nennt zwei frühere Stellen des Kurstexts: Abschnitt 2.3 und Abschnitt 4.3.5.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass der Addierer nicht im Kurstext steht.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zum Exkurs „Speichern und Rechnen im Rechner“, zur Seite „Zahlenkreis“ und zur Seite „Gatter aus Schaltern“ folgen. Alle sind eigene HTML-Seiten.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 5.

### Abschnitte des Kurstexts

- 4.3.7.2.1
- 4.3.5

## Schritt 1: Ein Gatter

Reiter: gatter

### Aufbau der Darstellung

- Die Demo ist eine Tabelle mit einer einzigen Bit-Spalte. Oben stehen die zwei Eingänge `a` und `b` als Knöpfe, rechts daneben „Strom an“ oder „Strom aus“.
- Darunter stehen vier Zeilen mit den Ergebnissen: AND `a & b`, OR `a | b`, XOR `a ^ b` und NOT `^a`. Rechts steht jeweils die Regel.
- AND ist 1, wenn beide 1 sind. OR ist 1, wenn mindestens einer 1 ist. XOR ist 1, wenn genau einer 1 ist. NOT ist 1, wenn `a` 0 ist.
- Die letzte Zeile heißt „Addition `a + b`“. Sie hat zwei Bits: links den Übertrag in Rot, rechts die Summe in Grün. Die Summe kommt aus XOR, der Übertrag aus AND.
- Die vier Fälle der Addition: 0 + 0 = 00. 0 + 1 = 01. 1 + 0 = 01. 1 + 1 = 10, also Summe 0 und Übertrag 1.

### Handgriffe

- Das Bit von `a` oder das Bit von `b` anklicken: Es wechselt zwischen 0 und 1.
- Mehr lässt sich in diesem Schritt nicht einstellen.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.3.7.2.1

## Schritt 2: Die Gatter auf einem Byte

Reiter: byte

### Aufbau der Darstellung

- Oben steht die Tabelle „Die Bit-Operatoren in Go“ mit sieben Zeilen: `a & b`, `a | b`, `a ^ b`, `^a`, `a &^ b`, `a << n` und `a >> n`. Die Spalten nennen den Namen, wann das Ergebnis-Bit 1 ist, und die Aufgabe im Rechenwerk.
- Die Demo ist eine Bit-Tabelle mit acht Spalten. Über den Spalten stehen die Wertigkeiten 128 bis 1. Zwischen der vierten und fünften Spalte ist eine Lücke.
- Die ersten zwei Zeilen sind `a` und `b`. Ihre Bits sind Knöpfe. Über der Tabelle stehen zwei Zahlenfelder für `a` und `b` und der Knopf „Beispiel des Kurstexts“.
- Darunter stehen sieben Zeilen mit Ergebnissen: `a & b`, `a | b`, `a ^ b`, `a &^ b`, `^a`, `a << 1` und `a >> 1`.
- Jedes Ergebnis-Bit von `&`, `|`, `^` und `&^` hängt nur von den zwei Bits darüber ab. Die Operatoren arbeiten auf allen acht Stellen zugleich und unabhängig voneinander.
- Mit dem Beispiel des Kurstexts, `a` = 10 = 0000 1010 und `b` = 12 = 0000 1100: `a & b` = 0000 1000. `a | b` = 0000 1110. `a ^ b` = 0000 0110. `a &^ b` = 0000 0010. `^a` = 1111 0101. `a << 1` = 0001 0100. `a >> 1` = 0000 0101.
- Unter der Tabelle steht ein dunkler Codeblock: die Deklaration von `a` und `b` als `uint8` mit binären Literalen und sieben Zeilen `fmt.Printf("%08b\n", …)` mit dem Bitmuster hinter `//`.
- `%08b` gibt eine Zahl binär mit acht Stellen und führenden Nullen aus.
- Ein gelber Kasten „Wozu das heute gebraucht wird“ nennt eine Forschungsarbeit von 2025: Ein Sprachmodell wurde verkleinert, indem Gewichte als einzelne Bits in `uint32`-Werten gespeichert und Multiplikationen durch XOR ersetzt wurden. Die Quelle ist Lee u. a., „LittleBit“, NeurIPS 2025. Mehr als diese Sätze steht auf der Seite nicht dazu.

### Handgriffe

- Ein Bit von `a` oder `b` anklicken.
- `a` und `b` als ganze Zahl von 0 bis 255 eintragen.
- „Beispiel des Kurstexts“ wählen: Das setzt `a` auf 10 und `b` auf 12.
- Geschoben wird in der Demo immer um genau 1 Stelle. Um mehr Stellen schiebt die Seite „Bitverschiebung“.
- Die Auflösung der Frage aufklappen. Den Links zu den Seiten „Zahlenkreis“ und „Bitverschiebung“ und zur Quelle folgen.

### Abschnitte des Kurstexts

- 4.3.7.2.1
- 4.3.7.2.2

## Schritt 3: Eine Stelle mit Übertrag

Reiter: stelle

### Aufbau der Darstellung

- Die Demo zeigt eine einzelne Stelle einer Addition, die Stelle mit der Wertigkeit 2. Sie hat drei Eingänge: `a`, `b` und den Übertrag von rechts. Alle drei sind Knöpfe.
- Oben steht eine kleine Bit-Tabelle mit zwei Spalten, überschrieben mit 4 und 2. In der Zeile „Übertrag“ steht rechts der Übertrag von rechts zum Umschalten und links der neue Übertrag. Darunter stehen `a`, `+ b` und unter einem Strich die Summe.
- Darunter erklärt ein Kasten die Rechnung in Worten und mit zwei Formeln: Summe = a XOR b XOR Übertrag. Übertrag nach links = (a AND b) OR (Übertrag AND (a XOR b)).
- Die Regel in Worten: Keine Eins ergibt Summe 0 ohne Übertrag. Eine Eins ergibt Summe 1 ohne Übertrag. Zwei Einsen ergeben 10, also Summe 0 und Übertrag 1. Drei Einsen ergeben 11, also Summe 1 und Übertrag 1.
- Darunter steht das Schaltbild des Addierers für eine Stelle. Es hat fünf Gatter: zwei XOR, zwei AND und ein OR.
- Im Schaltbild kommen `a` und `b` von oben und der Übertrag von rechts. Die Summe geht nach unten hinaus, der neue Übertrag nach links.
- Das obere XOR verknüpft `a` und `b`. Das untere XOR verknüpft dieses Zwischenergebnis mit dem Übertrag von rechts und liefert die Summe. Das linke AND oben verknüpft `a` und `b`. Das AND in der Mitte verknüpft das Zwischenergebnis mit dem Übertrag von rechts. Das OR fasst beide AND zusammen und liefert den Übertrag nach links.
- Eine dicke, farbige Leitung führt Strom und steht für 1. Eine dünne, graue Leitung steht für 0. Ein Gatter mit Ausgang 1 ist gelb unterlegt. Ein Punkt verbindet zwei Leitungen, ein Bogen führt eine Leitung über eine andere hinweg.
- Diese Schaltung heißt in der Technik Volladdierer. Das Wort steht nicht auf der Seite.

### Handgriffe

- Die drei Bits `a`, `b` und „Übertrag von rechts“ anklicken.
- Die Stelle ist immer die mit der Wertigkeit 2. Eine andere Stelle lässt sich hier nicht wählen.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.3.7.2.1

## Schritt 4: Der Addierer

Reiter: addierer

### Aufbau der Darstellung

- Die Demo addiert zwei Bytes `a` und `b` Stelle für Stelle, von rechts nach links.
- Über der Tabelle stehen zwei Zahlenfelder für `a` und `b` und vier Knöpfe mit Beispielen: 17 + 4, 23 + 9, 255 + 1 und 127 + 1. Darunter stehen die Knöpfe „nächste Stelle“, „alle Stellen“ und „von vorn“.
- Die Bit-Tabelle hat vier Zeilen: „Übertrag“, `a`, `+ b` und unter einem Strich „Summe“. Der Übertrag steht über der Stelle, in die er hineingeht.
- Sichtbar sind nur die Summenbits und Überträge der Stellen, die schon gerechnet sind. Die zuletzt gerechnete Stelle ist umrandet.
- Links von der vordersten Stelle gibt es eine Extraspalte. Dort steht am Ende der Übertrag aus der vordersten Stelle. Ist er 1, ist er durchgestrichen: Er fällt weg.
- Unter der Tabelle erklärt ein Kasten die zuletzt gerechnete Stelle, mit den zwei Formeln aus Schritt 3.
- Darunter steht das Schaltbild aus Schritt 3, hier mit den Werten der zuletzt gerechneten Stelle. Acht solche Schaltungen stehen im Rechenwerk nebeneinander.
- Sind alle acht Stellen gerechnet, erscheinen ein Kasten mit dem Ergebnis und ein dunkler Codeblock. Der Kasten liest das Ergebnis zweimal: als `uint8` und als `int8`. Er ist rot, wenn eine der zwei Lesarten falsch ist.
- 17 + 4 = 21: 0001 0001 + 0000 0100 = 0001 0101. An keiner Stelle entsteht ein Übertrag.
- 23 + 9 = 32: 0001 0111 + 0000 1001 = 0010 0000, mit Überträgen an mehreren Stellen.
- 255 + 1: Das Ergebnis ist 0000 0000. An jeder Stelle entsteht ein Übertrag, der letzte fällt weg. Als `uint8` gelesen ist das 0 statt 256, also falsch. Als `int8` gelesen ist es −1 + 1 = 0, also richtig.
- 127 + 1: Das Ergebnis ist 1000 0000. Als `uint8` gelesen ist das 128 und richtig. Als `int8` gelesen ist es −128 statt 128, also falsch.
- Dasselbe Bitmuster bedeutet als `uint8` und als `int8` verschiedene Zahlen. Das Rechenwerk rechnet für beide Typen gleich.

### Handgriffe

- „nächste Stelle“ rechnet eine Stelle weiter. „alle Stellen“ rechnet bis zum Ende. „von vorn“ beginnt wieder bei der Stelle ganz rechts.
- Ein Bit von `a` oder `b` anklicken, `a` und `b` als ganze Zahl von 0 bis 255 eintragen oder einen der vier Knöpfe mit Beispielen wählen. Jede dieser Änderungen beginnt die Rechnung von vorn.
- Die Auflösungen der drei Fragen aufklappen. Dem Link zur Seite „Zahlenkreis“ folgen.

### Abschnitte des Kurstexts

- 4.3.7.2.1
- 4.3.8

## Schritt 5: Die Subtraktion

Reiter: subtraktion

### Aufbau der Darstellung

- Oben steht: Das Rechenwerk subtrahiert nicht. Es bildet aus `b` die Zahl `−b` und addiert sie. Dafür dreht es alle Bits von `b` um und zählt 1 dazu. Diese Darstellung heißt Zweierkomplement.
- Über der Tabelle stehen zwei Zahlenfelder für `a` und `b` und drei Knöpfe mit Beispielen: 17 − 4, 10 − 15 und 0 − 1.
- Die Bit-Tabelle hat acht Zeilen. Zuerst `a` und `b` mit Bits zum Umschalten. Dann „Schritt 1: `^b`, alle Bits umdrehen“. Dann „Schritt 2: `^b + 1`, das ist −b“. Dann die Addition: „Übertrag“, `a`, `+ (−b)` und „Schritt 3: Summe“.
- Die Tabelle zeigt alle drei Schritte auf einmal. Einen Lauf Stelle für Stelle gibt es hier nicht.
- Unter der Tabelle steht ein Kasten mit dem Ergebnis, gelesen als `uint8` und als `int8`, und ein dunkler Codeblock mit `a - b` und `a + (^b + 1)`. Beide liefern das gleiche Ergebnis.
- 17 − 4: `^b` ist 1111 1011, `−b` ist 1111 1100, die Summe ist 0000 1101, also 13. Der Übertrag aus der vordersten Stelle fällt weg. Bei der Subtraktion gehört das dazu und ist kein Fehler.
- 10 − 15: Die Summe ist 1111 1011. Als `uint8` gelesen ist das 251 und falsch, als `int8` gelesen ist es −5 und richtig.
- 0 − 1: Die Summe ist 1111 1111. Als `uint8` gelesen ist das 255, als `int8` gelesen −1.
- Ein gelber Kasten sagt: Ein echtes Rechenwerk dreht die Bits von `b` um und setzt den ersten Übertrag ganz rechts auf 1. Das Ergebnis ist das gleiche.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf: „Zahlenkreis“, „Bitverschiebung“, „Quiz zur Bitverschiebung“.

### Handgriffe

- Ein Bit von `a` oder `b` anklicken, `a` und `b` als ganze Zahl von 0 bis 255 eintragen oder einen der drei Knöpfe mit Beispielen wählen.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.3.5
- 4.3.7.2.1

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Ein Gatter, 4.3.7.2.1, Seiten 72–74: bit-weise Operatoren: Die bit-weisen Operatoren und ihre Tabellen stammen aus dem Kurstext. Die Zeile „Addition“ gibt die Lektion dazu.
  - 2 Die Gatter auf einem Byte, 4.3.7.2.1, Seiten 72–74: Die Werte 10 und 12 sind das Beispiel des Kurstexts. Die Demo mit zwei einstellbaren Bytes gibt die Lektion dazu.
  - 3 Eine Stelle mit Übertrag, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Schaltung ist eine Ergänzung des Mentors.
  - 4 Der Addierer, 4.3.8, Seiten 78–84: Überlauf: Der Überlauf, der beim letzten Übertrag entsteht, wird in Abschnitt 4.3.8 behandelt. Der Addierer steht nicht im Kurstext.
  - 5 Die Subtraktion, 4.3.5, Seiten 64–66: Zweierkomplement: Im Kurstext steht, dass sich die Subtraktion mit dem Zweierkomplement auf die Addition zurückführen lässt. Die Rechnung Stelle für Stelle gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, Rechnen: Der Rechner wandelt Folgen von 0 und 1 in andere Folgen um. Lektion: Exkurs: Speichern und Rechnen im Rechner.
  - Abschnitt 4.3.5, Seiten 64–66, Zweierkomplement: Lektion: Zweierkomplement.
  - Ohne Abschnitt im Kurstext: Wie man die Gatter aus Schaltern baut, zeigt die Lektion Gatter aus Schaltern.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zu den Seiten „Zahlenkreis“ und „Bitverschiebung“ und zum Quiz zur Bitverschiebung folgen. Ein Quiz zum Rechenwerk gibt es nicht.

### Abschnitte des Kurstexts

- 2.3
- 4.3.5
- 4.3.7.2.1
- 4.3.8
