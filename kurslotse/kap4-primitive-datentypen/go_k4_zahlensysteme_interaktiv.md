# Zahlensysteme in Go

Die Seite zeigt, wie eine ganze Zahl im Quelltext dezimal, binär, oktal und hexadezimal geschrieben wird. Schritt 1 zerlegt ein Integer-Literal in seine Ziffern und rechnet den Wert aus. Schritt 2 zeigt, welche Schreibweisen Go annimmt. Schritt 3 zeigt die Ausgabe mit `fmt.Printf`. Die Seite gehört zu den Abschnitten 4.3.1 und 4.3.2 des Kurstexts.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Integer-Literal ist eine ganze Zahl, die im Quelltext steht, zum Beispiel `42`.
- Das Präfix sagt, welches Zahlensystem gemeint ist: `0b` oder `0B` für binär, `0o` oder `0O` oder die bloße `0` für oktal, `0x` oder `0X` für hexadezimal. Ein Dezimalliteral hat kein Präfix.
- `42`, `0b101010`, `0o52` und `0x2a` haben alle den Wert 42. Im Speicher steht jedes Mal das gleiche Bitmuster. Go merkt sich nicht, in welchem Zahlensystem ein Wert im Quelltext stand.
- Der Wert eines Literals ist die Summe aus Ziffer mal Basis hoch Stelle. Die Stellen zählen von rechts und beginnen bei 0.
- Die Hexadezimalziffern `a` bis `f` stehen für die Werte 10 bis 15.
- Ein Unterstrich gliedert Ziffern und ändert den Wert nicht. Er steht zwischen zwei Ziffern oder zwischen Präfix und erster Ziffer.
- Die Seite behandelt nur Literale ohne Vorzeichen. In `-42` ist das Minus ein Operator und `42` das Literal. Wie Go negative Zahlen als Bitmuster ablegt, steht auf den Seiten „Zahlenkreis“ und „Zweierkomplement“.
- Auf der Seite läuft kein Go. In Schritt 1 rechnet der Browser nach den Regeln von Go. Die Urteile und Meldungen in Schritt 2 und die Ausgaben in Schritt 3 sind mit Go 1.26 nachgeprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein Codeblock zeigt die Zahl 42 in vier Schreibweisen: `42`, `0b101010`, `0o52` und `0x2a`.
- Der Teil „Wozu vier Schreibweisen“ nennt den Zweck: Dezimal für Menschen. Binär, wenn es auf die einzelnen Bits ankommt. Hexadezimal als kurze Schreibweise für Bitfolgen, weil eine Hexadezimalziffer für vier Bits steht und zwei Ziffern für ein Byte. Oktal ist selten und gibt auf Unix-Systemen die Zugriffsrechte einer Datei an.
- Eine Liste nennt zwei frühere Stellen des Kurstexts: Abschnitt 2.2 und Abschnitt 3.3.5.
- Eine zweite Liste sagt, was die drei Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Dem Link zur Seite „Variablen, Typen und Speicher“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.3.1

## Schritt 1: Stellenwert

Reiter: stellenwert

### Aufbau der Darstellung

- Oben steht eine Tabelle mit den vier Zahlensystemen: Basis, Ziffern und Wert der Stellen. Binär: Basis 2, Stellen 1, 2, 4, 8. Oktal: Basis 8, Stellen 1, 8, 64, 512. Dezimal: Basis 10. Hexadezimal: Basis 16, Stellen 1, 16, 256, 4096.
- Darunter steht die Regel: Der Wert ist die Summe über alle Stellen von Ziffer mal Basis hoch Stelle. Das Beispiel ist `0x7b`: 7 · 16 + 11 · 1 = 123.
- Die Demo hat ein Eingabefeld „Integer-Literal“ und den Knopf „zerlegen“.
- Darunter stehen zehn Knöpfe mit Beispielen. Aus dem Kurstext: `0b1000_1011`, `0o14`, `031`, `0x7b`, `0X0000_BAFF` und `1_038_605`. Gleiche Ziffern mit anderem Präfix: `0b_10`, `0o_10`, `10` und `0x_10`. Sie haben die Werte 2, 8, 10 und 16.
- Ein grüner Kasten sagt, in welchem Zahlensystem das Literal steht, welches Präfix es hat und welchen Wert es hat. Ist die Eingabe kein Integer-Literal, ist der Kasten rot und nennt den Grund.
- Eine Tabelle zerlegt das Literal Ziffer für Ziffer. Spalten: Ziffer, Wert der Ziffer, Stelle, Wert der Stelle und Beitrag. Die letzte Zeile ist die Summe der Beiträge.
- Vier Kästen zeigen den gleichen Wert binär, oktal, dezimal und hexadezimal. Der Kasten des eingegebenen Zahlensystems ist grün. In der binären und der hexadezimalen Schreibweise trennt ein Unterstrich Gruppen aus vier Ziffern.
- Für Werte von 1 bis 65535 zeigt eine zweite Tabelle den Rückweg: Der Wert wird immer wieder durch 2 geteilt, die Reste von unten nach oben gelesen sind die Binärziffern. Für oktal teilt man durch 8, für hexadezimal durch 16.
- Bei einem Literal mit führender `0` ohne Buchstaben warnt der Kasten: Go liest es als Oktalzahl.
- Der Schritt endet mit einem Link auf Schritt 2 der Seite „Speichermodell“.

### Handgriffe

- Ein Literal mit höchstens 40 Zeichen eintragen und „zerlegen“ wählen oder die Eingabetaste drücken.
- Einen der zehn Knöpfe mit Beispielen wählen.
- Das Feld nimmt nur ein Integer-Literal an: kein Vorzeichen, keinen Punkt, keinen Typ, keine Rechnung.
- Die Seite zerlegt bis zu 20 Ziffern, bei einer Binärzahl bis zu 64.

### Abschnitte des Kurstexts

- 4.3.1

## Schritt 2: Literale in Go

Reiter: literale

### Aufbau der Darstellung

- Oben steht eine Tabelle mit den Präfixen der vier Zahlensysteme und je einem Beispiel: `0b1101` ist 13, `0o31` und `031` sind 25, `1024`, `0x7b` ist 123.
- Der Teil „Drei Regeln“ nennt: Buchstaben im Präfix und Hexadezimalziffern dürfen groß oder klein sein. Der Unterstrich gliedert. Nach einem Präfix dürfen Nullen stehen, ein Dezimalliteral beginnt nie mit einer Null.
- Der Teil „Die führende Null“ zeigt das Beispiel aus Hinweis 4.6 des Kurstexts: `var c int = 076` hat den Wert 62, denn 7 · 8 + 6 = 62. Die Ausgabe von `fmt.Println(a, b, c)` ist `245 123 62`. Go meldet nichts.
- Ein gelber Kasten rät, eine Oktalzahl mit `0o` zu schreiben. Der Rat steht nicht im Kurstext.
- Die Demo zeigt zwölf Knöpfe mit Literalen. Das gewählte Literal steht in der Zeile `var a int = …` in einem dunklen Codeblock.
- Darunter sagt ein Kasten, ob Go die Zeile übersetzt. Übersetzt Go sie, ist der Kasten grün und nennt den Wert. Sonst ist er rot und nennt die Meldung des Compilers. In beiden Fällen folgt der Grund.
- Go übersetzt: `1_038_605` mit dem Wert 1038605, `0b1000_1011` mit 139, `0X0000_BAFF` mit 47871, `0x_2a` mit 42, `031` mit 25 und `076` mit 62.
- Go übersetzt nicht: `08` mit der Meldung `invalid digit '8' in octal literal`, `0b102` mit `invalid digit '2' in binary literal`, `_42` mit `undefined: _42`, und `42_`, `4__2` und `0_x42` mit `'_' must separate successive digits`.
- `_42` ist für Go ein Name und kein Literal, deshalb sucht Go eine Variable mit diesem Namen.

### Handgriffe

- Eines der zwölf Literale wählen.
- Ein eigenes Literal lässt sich hier nicht eingeben. Das geht in Schritt 1.

### Abschnitte des Kurstexts

- 4.3.1

## Schritt 3: Ausgabe mit Printf

Reiter: printf

### Aufbau der Darstellung

- Oben stehen fünf Knöpfe für den Wert von `a`: 78, 163, 42, 255 und 62. Die Werte 78 und 163 stammen aus dem Kurstext.
- Eine Tabelle zeigt zehn Verben mit Erklärung und Ausgabe für den gewählten Wert: `%v`, `%d`, `%b`, `%o`, `%x`, `%X`, `%#b`, `%#o`, `%#x` und `%O`.
- Für 78: `%v` und `%d` liefern `78`, `%b` liefert `1001110`, `%o` liefert `116`, `%x` liefert `4e`, `%X` liefert `4E`, `%#b` liefert `0b1001110`, `%#o` liefert `0116`, `%#x` liefert `0x4e`, `%O` liefert `0o116`.
- Der Teil „Drei Bausteine“ erklärt das Verb, das Flag `#` und die Angabe `[1]`. Das Flag `#` setzt das Präfix davor, bei einer Oktalzahl die bloße `0`. Die Angabe `[1]` lässt ein Verb noch einmal das erste Argument nehmen.
- Ein dunkler Codeblock zeigt die drei Zeilen mit `fmt.Printf` aus dem Kurstext mit ihren Ausgaben hinter `//`, gerechnet für den gewählten Wert.
- Für 78 lauten die Ausgaben `78:  1001110  116  4e`, dann `78:  0b1001110 0116 0x4e`, dann `78:  1001110  0o116  4E`.
- Im Kurstext gibt die dritte Zeile die Variable `b` mit dem Wert 163 aus und liefert `163:  10100011  0o243  A3`. Auf der Seite geben alle drei Zeilen `a` aus.
- `fmt.Println` gibt eine ganze Zahl immer dezimal aus.
- Am Ende des Schritts unter „Nächste Schritte“ stehen Links zu zwei Quizseiten und zu den Seiten „Zahlenkreis“, „Zweierkomplement“ und „Rechenwerk“. Das Quiz „Zahlensysteme umrechnen“ würfelt Aufgaben aus: ein binäres, oktales oder hexadezimales Literal in eine Dezimalzahl umrechnen oder einen Wert als Literal schreiben. Das Quiz zu den Integer-Literalen stellt zehn Aussagen zur Wahl. Welche Aussagen das sind, steht auf dieser Seite nicht.

### Handgriffe

- Einen der fünf Werte wählen.
- Ein eigener Wert und ein eigenes Verb lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.3.2
- 4.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Stellenwert, 4.3.1, Seiten 59–61: Die Beispiele auf den Knöpfen stammen aus dem Kurstext. Die Zerlegung in Stellen und das Teilen mit Rest rechnet die Lektion vor.
  - 2 Literale in Go, 4.3.1, Seiten 59–61: Hinweis 4.6: Die fehlerhaften Literale _42, 42_, 4__2 und 0_x42 nennt der Kurstext. Die Meldungen des Compilers, die Literale 08 und 0b102 und der Rat, eine Oktalzahl mit 0o zu schreiben, gibt die Lektion dazu.
  - 3 Ausgabe mit Printf, 4.3.2, Seiten 62–63; 4.2, Seiten 53–58: fmt.Printf: Die Werte 78 und 163 und die drei Zeilen mit fmt.Printf stammen aus dem Kurstext. Die Werte 42, 255 und 62 sind eigene Beispiele.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 3.3.5, Seiten 33–35, Fußnote 14: Der Wert 10 ließe sich auch als 0b1010 oder 0xA schreiben. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 4.2, Seiten 53–58, fmt.Printf: Eine Lektion dazu gibt es noch nicht.
  - Abschnitt 4.3.3, Seite 63, Literal und Wert: Lektion: Speichermodell, Schritt 2.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links am Ende folgen. Alle führen auf eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 3.3.5
- 4.2
- 4.3.1
- 4.3.2
- 4.3.3
