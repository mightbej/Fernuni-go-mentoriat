# Strings: Bytes und Zeichen in Go

Die Seite zeigt, dass ein String in Go eine Folge von Bytes ist und dass ein Zeichen ein bis vier Bytes belegt. Sie legt einen String als Reihe von Bytes hin und zeigt, was `len(s)`, `s[i]` und `s[low:high]` liefern. Der zweite Schritt zeigt den Typ `[]rune`, mit dem man Zeichen zählt. Die Seite gehört zu Abschnitt 4.6.2 des Kurstexts.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Go speichert einen String im Format UTF-8. Ein ASCII-Zeichen belegt 1 Byte. ü, ß und é belegen 2 Bytes. ∞ und € belegen 3 Bytes. Ein Emoji wie 😀 belegt 4 Bytes.
- `len(s)` liefert die Zahl der Bytes, nicht die Zahl der Zeichen.
- `s[i]` liefert ein einzelnes Byte vom Typ `byte`. `byte` ist ein anderer Name für `uint8`.
- `s[low:high]` liefert den Substring aus den Bytes mit Index `low` bis `high − 1`. Das Byte mit Index `high` gehört nicht mehr dazu. Der Substring hat `high − low` Bytes.
- Der Index eines Strings beginnt bei 0.
- Der Typ `[]rune` kommt in Abschnitt 4.6 des Kurstexts nicht vor. Er ist eine Ergänzung des Mentors. Stoff des Kurses ist, was im Kurstext steht.
- In der Demo rechnet der Browser, nicht Go. Die Ausgaben der zwei festen Programme in Schritt 2 sind mit Go 1.26 nachgeprüft.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben.
- Schritt 1 und 2 beginnen mit einem blauen Kasten mit Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Oben steht, wie viele Bytes ein Zeichen in UTF-8 braucht.
- Der Teil „Was in Go Bytes zählt“ nennt `len(s)`, `s[i]` und `s[low:high]`.
- Der Teil „Was Zeichen zählt“ nennt die Umwandlung in `[]rune`.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 3.3.7, Abschnitt 4.3.4 und Abschnitt 4.6.1.
- Eine zweite Liste sagt, was die zwei Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Zahlenkreis“ und „Strings in Go“ folgen. Beide sind eigene HTML-Seiten.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.6.2
- 4.6

## Schritt 1: Ein String als Byte-Reihe

Reiter: bytes

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit fünf Fragen zu „Tür“ und „Abo“.
- Eine Tabelle nennt, wie viele Bytes ein Zeichen in UTF-8 braucht.
- Die Demo hat oben ein Eingabefeld „String“, den Knopf „anzeigen“ und sechs Knöpfe mit Beispielen: Tür, Maßstab, anrufen, Grüße, ∞ und 😀.
- Darunter liegt der String als Raster mit vier Zeilen. Von oben nach unten: „Index in []rune“, „Zeichen“, „Byte“ und „Index im String“.
- In der Zeile „Zeichen“ steht jedes Zeichen groß in einem Kasten, darunter klein seine Unicode-Nummer, zum Beispiel U+00FC für ü. Ein Zeichen mit mehreren Bytes ist so breit wie seine Bytes zusammen. Ein Zeichen mit 1 Byte ist grün, ein Zeichen mit mehreren Bytes orange.
- In der Zeile „Byte“ steht je Byte ein Knopf mit dem Wert hexadezimal und darunter dezimal. Das ü besteht aus den zwei Bytes c3 und bc, dezimal 195 und 188.
- In der Zeile „Index im String“ steht unter jedem Byte sein Index. In der Zeile „Index in []rune“ steht über jedem Zeichen sein Index in Zeichen gezählt.
- Die ausgewählten Bytes sind blau. Die Auswahl ist der Bereich von `low` bis `high`.
- Für „Tür“: T hat Byte-Index 0, ü belegt die Byte-Indizes 1 und 2, r hat Byte-Index 3. In `[]rune` haben T, ü und r die Indizes 0, 1 und 2.
- Unter dem Raster stehen zwei Zahlenfelder `low` und `high` und der Knopf „Bereich setzen“.
- Darunter steht ein Kasten mit Angaben: `len(s)` in Bytes, `len(r)` in Zeichen, `s[low]` mit Wert und Typ, `s[low:high]` als Substring und `string(r[a:b])` für dieselbe Auswahl in Zeichen gezählt.
- Darunter steht ein farbiger Kasten mit einer Meldung. Er ist rot, wenn der Schnitt mitten durch ein Zeichen geht.
- Ganz unten steht ein dunkler Codeblock mit den Zeilen in Go und ihren Ausgaben hinter `//`.
- Geht ein Schnitt mitten durch ein Zeichen, bleibt ein Teil eines Zeichens übrig. Go meldet keinen Fehler. Das Verb `%q` zeigt solche Bytes als `\x…`, zum Beispiel `"\xc3"`. Die Konsole zeigt dafür meist das Ersatzzeichen �.
- `"Tür"[1:3]` ist „ü“. `"Tür"[3:4]` ist „r“. `"Tür"[1:2]` ist ein halbes ü, das Byte 195.
- `"H" + s[0]` übersetzt Go nicht, weil `s[0]` den Typ `byte` hat und kein String ist. Die Meldung: `invalid operation: "H" + s[0] (mismatched types untyped string and byte)`.

### Handgriffe

- Einen String mit höchstens 24 Zeichen eintragen und „anzeigen“ wählen.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Ein Byte anklicken: Es wird allein ausgewählt. Ein zweiter Klick auf ein anderes Byte wählt den Bereich dazwischen.
- Ein Zeichen anklicken: Alle seine Bytes werden ausgewählt. Ein zweiter Klick auf ein anderes Zeichen wählt den Bereich bis dorthin.
- `low` und `high` eintragen und „Bereich setzen“ wählen. Es muss gelten: 0 ≤ low ≤ high ≤ Zahl der Bytes. Sonst meldet die Seite einen Fehler.
- Die Auflösungen der fünf Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.6.2
- 4.6.2.1

## Schritt 2: Mit Zeichen arbeiten

Reiter: zeichen

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text, einer Tabelle und zwei Codeblöcken. Eine Demo gibt es hier nicht.
- Oben steht der blaue Kasten mit zwei Fragen zu „Grüße“ und „Maßstab“.
- Die Tabelle nennt vier Typen: `string` als Folge von Bytes, `byte` als anderen Namen für `uint8`, `rune` als Unicode-Nummer eines Zeichens und anderen Namen für `int32`, `[]rune` als Folge ganzer Zeichen.
- `[]rune(s)` zerlegt den String in seine Zeichen. `string(…)` macht aus Runen wieder einen String.
- Der erste Codeblock arbeitet mit `s := "Maßstab"` und `r := []rune(s)`. `len(s)` ist 8, `len(r)` ist 7. `string(r[2])` ist „ß“. `string(r[3:7])` ist „stab“. `s[3:7]` schneidet durch das ß und liefert `"\x9fsta"`.
- Der zweite Codeblock läuft mit `for i, z := range "Tür"` zeichenweise über den String. Die Ausgabe: `0 T`, `1 ü`, `3 r`. Der Index ist die Position des ersten Bytes des Zeichens, deshalb folgt auf 1 die 3.
- Die Schleife mit `range` wird im Kurstext erst in Abschnitt 6.2 behandelt.
- „Grüße“ hat 7 Bytes und 5 Zeichen. `utf8.RuneCountInString(s)` zählt die Zeichen, ohne den String umzuwandeln.
- Ein gelber Kasten sagt: Ein `[]rune` belegt 4 Bytes je Zeichen, und die Umwandlung kopiert den ganzen Text.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Ob `[]rune` in der Klausur vorkommt, lässt sich deshalb nicht sagen. Verbindlich sind der Kurstext und Moodle der FernUniversität.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links zu zwei Quizseiten. Welche Fragen sie stellen, steht auf dieser Seite nicht.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Die zwei Codeblöcke sind fest und lassen sich nicht ausführen.
- Die Auflösungen der zwei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.6.2
- 4.6.2.1

## Schritt 3: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Ein String als Byte-Reihe, 4.6.2, Seiten 102–105: String-Datentyp; 4.6.2.1, Seiten 105–107: Zugriff auf Teile: Der String als Folge von Bytes, die Länge und der Zugriff mit einem Index stammen aus dem Kurstext. Die Zeilen mit []rune gehen über den Kurstext hinaus.
  - 2 Mit Zeichen arbeiten, kein Abschnitt: Aus Abschnitt 4.6 stammt nichts. Der Typ []rune und die zwei Programme sind eine Ergänzung des Mentors.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Zeichenketten: Variablen können auch Zeichenketten sein, mit eigenen Operationen. Lektion: Variablen, Typen und Speicher.
  - Abschnitt 3.3.7, Seite 36, Strings: Eine Lektion gibt es nicht, dazu gibt es das Quiz Ausgabe und String-Verkettung.
  - Abschnitt 4.6.1, Seiten 99–102, String-Literale: Lektion: Strings in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 2“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zu den zwei Quizseiten und zur Seite „Strings in Go“ folgen. Alle sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.2
- 3.3.7
- 4.6.1
- 4.6.2
- 4.6.2.1
