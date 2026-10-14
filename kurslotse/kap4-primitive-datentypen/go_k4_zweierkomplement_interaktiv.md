# Zweierkomplement: negative Zahlen als Bitmuster in Go

Die Seite zeigt, wie Go eine negative ganze Zahl als Bitmuster ablegt. Schritt 1 liest ein Bitmuster als `uint8` und als `int8`. Schritt 2 bildet aus einer Zahl ihr Negatives. Schritt 3 trennt das binäre Literal im Quelltext vom Bitmuster im Speicher. Die Seite gehört zu Abschnitt 4.3.5 des Kurstexts.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite rechnet mit 8 Bit, also mit den Typen `uint8` und `int8`. Die größeren Typen arbeiten genauso, nur mit mehr Bits.
- Die Bits eines Bytes zählen von rechts. Das rechte Bit ist Bit 0, das linke Bit 7.
- Beim Typ `uint8` haben die Bits die Wertigkeiten 128, 64, 32, 16, 8, 4, 2 und 1. Der Typ reicht von 0 bis 255.
- Beim Typ `int8` hat das linke Bit die Wertigkeit −128. Die anderen sieben Bits zählen wie bei `uint8`. Der Typ reicht von −128 bis 127.
- Der Wert ist die Summe der Wertigkeiten aller Bits, die auf 1 stehen.
- Das linke Bit ist kein Minuszeichen. Das Vorzeichen folgt aus seiner negativen Wertigkeit.
- Bitmuster schreibt die Seite in zwei Vierergruppen, zum Beispiel 1111 0111.
- Im Kurstext steht zum Zweierkomplement nur die Wertigkeit −128 und die Tabelle mit dem Wert des linken Bits und dem Wert der rechten Bits. Die formale Definition gehört nach dem Kurstext in das Modul Computersysteme.
- Auf der Seite läuft kein Go. Der Browser rechnet. Die Programmzeilen in Schritt 2 und 3 und die Meldung des Compilers sind mit Go 1.26 nachgeprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein gelber Kasten nennt die eine Änderung: Bei einem Typ mit Vorzeichen hat das linke Bit eine negative Wertigkeit, bei `int8` ist sie −128.
- Der Teil „Was das Zweierkomplement bringt“ nennt zwei Vorteile: Die Subtraktion lässt sich als Addition des Negativen ausführen, und die Negation ist einfach.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.2, Abschnitt 4.3.1 und Abschnitt 4.3.4.
- Eine zweite Liste sagt, was die drei Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Rechenwerk“, „Variablen, Typen und Speicher“, „Zahlensysteme in Go“ und „Zahlenkreis“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.3.5

## Schritt 1: Die Wertigkeit −128

Reiter: wertigkeit

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Eine Tabelle stellt `uint8` und `int8` nebeneinander: Wertigkeit von Bit 7, Wertigkeit der anderen Bits und Wertebereich.
- Die Demo zeigt sechs Knöpfe mit den Bitmustern aus dem Kurstext: 0000 0000, 0000 0001, 0111 1111, 1000 0000, 1000 0001 und 1111 1111.
- Darunter steht ein Byte: acht Bits als große Knöpfe. Über jedem Bit steht seine Wertigkeit beim Typ `int8`: −128, 64, 32, 16, 8, 4, 2 und 1. Ein Bit, das auf 1 steht, ist grün, das linke Bit ist blau.
- Darunter stehen zwei Kästen. Der grüne Kasten liest das Bitmuster als `uint8` und zeigt die Summe der Wertigkeiten. Der blaue Kasten liest es als `int8` in der Form des Kurstexts: Wert des linken Bits plus Wert der rechten Bits.
- Die Werte der sechs Bitmuster: 0000 0000 ist 0 und 0. 0000 0001 ist 1 und 1. 0111 1111 ist 127 und 127. 1000 0000 ist als `uint8` 128 und als `int8` −128. 1000 0001 ist 129 und −127. 1111 1111 ist 255 und −1.
- Der Teil „Was daraus folgt“ nennt vier Folgerungen: Das Vorzeichen steht am linken Bit. Das linke Bit ist kein Minuszeichen, 1000 1001 ist −119 und nicht −9. Die Bitmuster von 0 bis 127 haben in beiden Typen den gleichen Wert. Für n Bit reicht ein Typ ohne Vorzeichen von 0 bis 2^n − 1 und ein Typ mit Vorzeichen von −2^(n−1) bis 2^(n−1) − 1.
- Der Schritt endet mit einem Link auf Schritt 2 der Seite „Zahlenkreis“.

### Handgriffe

- Ein Bit anklicken: Es wechselt zwischen 0 und 1. Beide Kästen rechnen sofort mit.
- Einen der sechs Knöpfe mit Bitmustern wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Eine Dezimalzahl lässt sich in diesem Schritt nicht eingeben. Das geht in Schritt 2.

### Abschnitte des Kurstexts

- 4.3.5
- 4.3.4

## Schritt 2: Eine Zahl negieren

Reiter: negieren

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen und aufklappbaren Auflösungen.
- Darunter stehen die zwei Handgriffe: alle Bits umkehren, dann 1 addieren. In Go kehrt der Operator `^` vor einem Wert alle Bits um. `^x + 1` ist das Gleiche wie `-x`.
- Der Teil „Warum das stimmt“ begründet es: `x + ^x` ergibt an jeder Stelle genau eine 1, also 1111 1111, und das ist −1. Aus `x + ^x = -1` folgt `^x + 1 = -x`.
- Die Demo hat ein Zahlenfeld für eine Zahl von −128 bis 127, den Knopf „negieren“ und den Knopf „das Ergebnis negieren“. Darunter stehen sechs Knöpfe mit Beispielen: 42, 9, 1, −9, 0 und −128.
- Eine Tabelle mit drei Zeilen zeigt die Rechnung: die Zahl als `int8` mit Bitmuster und Probe über die Wertigkeiten, dann „1. Alle Bits umkehren“ mit dem Bitmuster und seinem Wert, dann „2. Eins addieren“ mit dem Ergebnis und der Probe.
- Für 42: 0010 1010, umgekehrt 1101 0101, das ist −43, plus 1 ergibt 1101 0110, das ist −42. Probe: −128 + 64 + 16 + 4 + 2 = −42.
- Für 9: 0000 1001, umgekehrt 1111 0110, plus 1 ergibt 1111 0111, das ist −9.
- Alle Bits umkehren allein ergibt immer −x − 1. Erst die 1 dazu ergibt −x.
- Ein grüner Kasten sagt, was aus der Zahl geworden ist, und nennt die zwei Ausdrücke in Go.
- Der Teil „Zwei Sonderfälle“: Die 0 negiert ist wieder 0, der Übertrag aus der vordersten Stelle fällt weg. Die −128 negiert ist wieder −128, weil 128 nicht in einen `int8` passt. Go meldet dabei nichts.
- Liegt die Eingabe außerhalb von −128 bis 127, zeigt ein roter Kasten eine Meldung. Die Rechnung darunter gehört dann noch zur vorigen Zahl.

### Handgriffe

- Eine ganze Zahl von −128 bis 127 eintragen und „negieren“ wählen oder die Eingabetaste drücken.
- „das Ergebnis negieren“ wählen: Die Seite negiert das letzte Ergebnis. Zweimal negieren führt zur Ausgangszahl zurück.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Einzelne Bits lassen sich in diesem Schritt nicht umschalten. Das geht in Schritt 1.

### Abschnitte des Kurstexts

- 4.3.5
- 4.3.7.2.1

## Schritt 3: Literal und Bitmuster

Reiter: literal

### Aufbau der Darstellung

- Oben steht der Unterschied der zwei Ebenen. Im Quelltext ist ein binäres Literal wie `0b1001` immer eine Zahl ab 0. Eine negative Zahl schreibt man als Betrag mit Minus davor: `-0b1001` ist −9. Im Speicher steht −9 in einem `int8` als 1111 0111. Dieses Bitmuster rechnet der Compiler aus.
- Das Literal `0b11110111` ist 247 und nicht −9.
- Die Demo zeigt sechs Knöpfe mit je einer Zeile. Die gewählte Zeile steht in einem dunklen Codeblock, bei fünf Zeilen mit der Ausgabe von `fmt.Println` hinter `//`.
- Darunter sagt ein Kasten, ob Go die Zeile übersetzt. Übersetzt Go sie, ist der Kasten grün und nennt Wert und Bitmuster. Sonst ist er rot und nennt die Meldung des Compilers.
- `var a int8 = 0b00001001`: Wert 9, Bitmuster 0000 1001.
- `var b int8 = -0b00001001`: Wert −9, Bitmuster 1111 0111.
- `var c int8 = 0b11110111`: Go übersetzt die Zeile nicht. Die Meldung: `cannot use 0b11110111 (untyped int constant 247) as int8 value in variable declaration (overflows)`.
- `var d uint8 = 0b11110111`: Wert 247, Bitmuster 1111 0111.
- `var e int8 = int8(d)` mit `d` vom Typ `uint8` und dem Wert 247: Wert −9, Bitmuster 1111 0111. Die Konvertierung einer Variable lässt die Bits stehen.
- `var m int8 = -0b10000000`: Wert −128, Bitmuster 1000 0000.
- Der Teil „Das Bitmuster sichtbar machen“ zeigt drei Zeilen: `fmt.Printf("%b\n", b)` gibt für −9 `-1001` aus, also den Betrag mit Minus. `fmt.Printf("%08b\n", uint8(b))` gibt `11110111` aus, das Bitmuster im Speicher.

### Handgriffe

- Eine der sechs Zeilen wählen.
- Eigener Code lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.3.5
- 4.3.1
- 4.7

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Die Wertigkeit −128, 4.3.5, Seiten 64–66: Die sechs Bitmuster auf den Knöpfen und die Zerlegung in den Wert des linken Bits und den Wert der rechten Bits stammen aus dem Kurstext. Das Beispiel mit −119 ist ein eigenes Beispiel.
  - 2 Eine Zahl negieren, 4.3.5, Seite 64; 4.3.7.2.1, Seiten 72–74: Operator ^: Im Kurstext steht nur, dass die Negation technisch sehr einfach ist. Die zwei Handgriffe, die Begründung und die Sonderfälle gibt die Lektion dazu.
  - 3 Literal und Bitmuster, 4.3.5, Seite 66: Der Unterschied zwischen System-Ebene und Quellcode-Ebene mit dem Beispiel -0b1001 stammt aus dem Kurstext. Die sechs Zeilen der Demo, die Meldung des Compilers und die Ausgabe mit %08b gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speicherelemente: Wie Folgen von 0 und 1 als Werte gelesen werden, hängt vom Typ ab. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 4.3.1, Seiten 59–61, Literale: Lektion: Zahlensysteme in Go.
- Unter „Was über den Kurstext hinausgeht“ steht:
  - Die formale Definition: Nach dem Kurstext gehört sie in das Modul Computersysteme.
- Unter „Wie es weitergeht“ stehen Links zum Quiz zu den Wertebereichen und zu den Seiten „Rechenwerk“, „Zahlenkreis“ und „Überlauf bei Ganzzahlen“. Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links am Ende folgen. Alle führen auf eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.2
- 4.3.1
- 4.3.5
- 4.3.7.2.1
