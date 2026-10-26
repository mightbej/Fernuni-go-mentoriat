# Überlauf bei Ganzzahlen in Go

Die Seite zeigt den arithmetischen Überlauf: Das Ergebnis einer Rechnung mit ganzen Zahlen passt nicht in seinen Typ, und Go rechnet ohne Meldung mit einem falschen Wert weiter. Die Seite rechnet die Beispiele des Kurstexts mit den Typen `uint8` und `int8` nach. Sie gehört zu Abschnitt 4.3.8 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Überlauf entsteht zur Laufzeit, wenn das richtige Ergebnis außerhalb des Wertebereichs des Typs liegt. Go meldet nichts: kein Abbruch, keine Warnung.
- `uint8` hat 8 Bit und reicht von 0 bis 255. `int8` hat 8 Bit und reicht von −128 bis 127.
- Bei einem Überlauf behält Go die Bits, die in den Typ passen, bei 8 Bit also die 8 rechten Bits. Die Bits links davon fallen weg.
- Bei einem Typ ohne Vorzeichen heißt das: Go rechnet modulo 2 hoch n, bei `uint8` modulo 256. Im Englischen heißt das „wrap around“.
- Bei einem Typ mit Vorzeichen werden die verbliebenen Bits im Zweierkomplement gelesen. Das linke Bit zählt bei `int8` −128.
- Der Typ eines Ergebnisses folgt aus den Operanden. `f + 80` hat den Typ von `f`. Wohin das Ergebnis geht, spielt keine Rolle.
- Die Typen `uint8` und `int8` dienen nur dem Nachrechnen. In der Praxis steht meist `int`.
- Auf der Seite läuft kein Go. Der Browser rechnet nach den Regeln von Go. Alle Programmzeilen und die Meldung des Compilers sind mit Go 1.26 nachgeprüft.
- Schritt 2, 3 und 4 beginnen mit einem blauen Kasten mit Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein Codeblock zeigt das Beispiel: `var f uint8 = 200` und `fmt.Println(f + 80)` mit der Ausgabe 24. Richtig wäre 280.
- Ein gelber Kasten sagt, dass Go einen Überlauf nicht meldet.
- Eine Liste nennt drei frühere Abschnitte des Kurstexts: 4.3.4, 4.3.5 und 4.3.7.1.
- Eine zweite Liste sagt, was die fünf Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Zahlenkreis“ und „Zweierkomplement“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.3.8

## Schritt 1: Keine Meldung

Reiter: laufzeit

### Aufbau der Darstellung

- Ein Codeblock zeigt die vier Zeilen des Kurstexts: `var f uint8 = 200`, `fmt.Println(f + 80)`, `f = f + 80` und `f = 200 + 80`.
- Die zwei Zeilen mit `f + 80` laufen zur Laufzeit über, ohne Meldung. Das Ergebnis ist 24.
- Die Zeile `f = 200 + 80` übersetzt Go nicht. `200 + 80` enthält nur Konstanten, der Compiler rechnet 280 selbst aus. Die Meldung: `cannot use 200 + 80 (untyped int constant 280) as uint8 value in assignment (overflows)`.
- Das Wort „overflows“ in dieser Meldung meint nur, dass die Konstante nicht in den Typ passt. Das ist ein Fehler beim Übersetzen und nicht der Überlauf zur Laufzeit, um den es auf der Seite geht.
- Der Teil „Der Typ des Ergebnisses“ sagt: `f + 80` hat den Typ `uint8`, weil `f` diesen Typ hat. Auch `fmt.Println(f + 80)` gibt deshalb 24 aus.
- Die Demo zeigt fünf Knöpfe für den Wert von `f`: 100, 175, 176, 200 und 255. Ein dunkler Codeblock zeigt die zwei Zeilen mit der Ausgabe hinter `//`. Ein Kasten sagt, ob die Rechnung überläuft. Er ist grün ohne Überlauf und rot mit Überlauf.
- Die Ausgaben: Mit 100 kommt 180 heraus, mit 175 kommt 255 heraus, mit 176 kommt 0 heraus, mit 200 kommt 24 heraus, mit 255 kommt 79 heraus. Der Überlauf beginnt bei 176, weil 176 + 80 = 256 ist.

### Handgriffe

- Einen der fünf Werte für `f` wählen.
- Der Summand 80 ist fest. Ein eigener Wert lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.3.8
- 4.3.7.1

## Schritt 2: Ohne Vorzeichen

Reiter: uint8

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen.
- Ein Codeblock rechnet 255 + 1 schriftlich mit Bits: 1111 1111 plus 0000 0001 ergibt 1 0000 0000, das sind 9 Bits. In `uint8` bleibt 0000 0000, also 0.
- Der Teil „Das Gleiche als Rechnung“ erklärt modulo 256 und nennt die drei Beispiele des Kurstexts: 256 mod 256 = 0, 270 mod 256 = 14, 510 mod 256 = 254.
- Die Demo hat zwei Zahlenfelder `a` und `b` für Werte von 0 bis 255, eine Auswahl für den Operator `+`, `-` oder `*` und den Knopf „rechnen“.
- Darunter stehen sechs Knöpfe mit Beispielen. Aus dem Kurstext: 255 + 1, 130 + 140 und 255 + 255. Eigene Beispiele: 200 + 50, 3 − 5 und 16 · 16.
- Ein dunkler Codeblock zeigt drei Zeilen: `var a uint8 = …`, `var b uint8 = …` und `fmt.Println(a + b)` mit der Ausgabe hinter `//`.
- Eine Tabelle zeigt fünf Zeilen: das richtige Ergebnis, das richtige Ergebnis mit 16 Bit geschrieben, die 8 rechten Bits, diese 8 Bits als `uint8` gelesen und die Rechnung modulo 256. Bei einem Überlauf sind die 8 linken Bits rot durchgestrichen, die 8 rechten Bits sind grün.
- Ein Kasten darunter sagt, ob es ein Überlauf ist. Er ist rot bei einem Überlauf und grün ohne.
- Die Ergebnisse: 255 + 1 ergibt 0. 130 + 140 ergibt 14. 255 + 255 ergibt 254. 200 + 50 ergibt 250 ohne Überlauf. 3 − 5 ergibt 254. 16 · 16 ergibt 0.
- Bei 3 − 5 ist das richtige Ergebnis −2. Die Seite schreibt es mit 16 Bit im Zweierkomplement. Im Kurstext wird dieser Fall in einer Fußnote vom Überlauf unterschieden: Dem Typ fehlt der negative Wert. Go rechnet trotzdem modulo 256.
- Die Demo rechnet mit zwei Variablen. Mit einem Literal rechts, das nicht in den Typ passt, würde Go die Zeile nicht übersetzen.
- Der Teil „Ein Überlauf mit Absicht“ nennt die Hash-Funktion: Bei Typen ohne Vorzeichen ist das Verhalten festgelegt und wird manchmal genutzt.

### Handgriffe

- Für `a` und `b` ganze Zahlen von 0 bis 255 eintragen, einen Operator wählen und „rechnen“ wählen oder die Eingabetaste drücken.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Eine Division und einen anderen Typ gibt es in dieser Demo nicht. Liegt eine Eingabe außerhalb von 0 bis 255, zeigt ein roter Kasten eine Meldung.

### Abschnitte des Kurstexts

- 4.3.8

## Schritt 3: Mit Vorzeichen

Reiter: int8

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen.
- Ein Codeblock rechnet 127 + 1 schriftlich mit Bits: 0111 1111 plus 0000 0001 ergibt 1000 0000. Als `int8` gelesen ist das −128.
- Eine Tabelle zeigt die sechs Rechnungen des Kurstexts mit richtigem Ergebnis, Ergebnis in Go und Vorzeichen: 127 + 1 ergibt −128. 127 + 127 ergibt −2. −128 − 1 ergibt 127. −128 + −128 ergibt 0. 12 · 121 ergibt −84. 3 · 121 ergibt 107.
- Bei den ersten fünf Rechnungen wechselt das Vorzeichen. Bei 3 · 121 bleibt es, obwohl 363 nicht in den Typ passt. Ein Überlauf ist also nicht immer am Vorzeichen zu erkennen.
- Vorzeichenwechsel heißt auf der Seite: Das richtige Ergebnis und das Ergebnis in Go haben verschiedene Vorzeichen. Die 0 zählt dabei wie im Kurstext zu den positiven Werten.
- Die Demo ist gebaut wie in Schritt 2, mit `int8` und Werten von −128 bis 127. Sechs Knöpfe tragen die sechs Rechnungen des Kurstexts.
- Die Tabelle der Demo hat vier Zeilen: das richtige Ergebnis, das richtige Ergebnis mit 16 Bit im Zweierkomplement, die 8 rechten Bits und diese 8 Bits als `int8` gelesen.
- Der Kasten darunter sagt bei einem Überlauf auch, ob das Vorzeichen wechselt.
- Im Kurstext steht rechts vom Operator ein Literal, zum Beispiel `iA + 1`. Die Demo rechnet mit zwei Variablen, das Ergebnis ist das gleiche.

### Handgriffe

- Für `a` und `b` ganze Zahlen von −128 bis 127 eintragen, einen Operator wählen und „rechnen“ wählen oder die Eingabetaste drücken.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Eine Division und einen anderen Typ gibt es in dieser Demo nicht.

### Abschnitte des Kurstexts

- 4.3.8
- 4.3.5

## Schritt 4: Teilausdrücke

Reiter: teilausdruecke

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Ein Codeblock zeigt die drei Zeilen des Kurstexts mit `var x int8 = 120`. Die Ausgaben: `x*3/3` gegen `x*(3/3)` ergibt `34 vs. 120`. `x*8/12` gegen `x/12*8` ergibt `-5 vs. 80`. `(x+20)/2` gegen `x/2+20/2` ergibt `-58 vs. 70`.
- Links und rechts von `"vs."` steht ein Ausdruck, der mathematisch das Gleiche bedeutet. Links läuft ein Teilausdruck über, rechts nicht.
- Die Demo zeigt vier Knöpfe mit je einem Paar von Ausdrücken. Das vierte Paar ist `x+20-40` gegen `x-40+20` aus dem Expertenwissen 4.14.
- Zwei weitere Knöpfe wählen den Typ von `x`: `int8` oder `int16`.
- Ein dunkler Codeblock zeigt die Zeile mit `fmt.Println` und ihre Ausgabe.
- Darunter stehen zwei Kästen, links und rechts je ein Ausdruck. Jeder Kasten wertet seinen Ausdruck Schritt für Schritt aus und nennt am Ende den Wert. Ein Schritt mit Überlauf ist rot.
- `x*3/3` mit `int8`: 120 · 3 = 360, Überlauf, es bleibt 104. Dann 104 : 3 = 34.
- `x*(3/3)`: `3/3` enthält nur Konstanten, der Compiler rechnet 1 aus. Dann 120 · 1 = 120.
- `x*8/12` mit `int8`: 120 · 8 = 960, Überlauf, es bleibt −64. Dann −64 : 12 = −5. `x/12*8`: 120 : 12 = 10, dann 10 · 8 = 80.
- `(x+20)/2` mit `int8`: 120 + 20 = 140, Überlauf, es bleibt −116. Dann −116 : 2 = −58. `x/2+20/2`: 120 : 2 = 60, `20/2` ist die Konstante 10, dann 60 + 10 = 70.
- `x+20-40` mit `int8`: 120 + 20 = 140, Überlauf, es bleibt −116. Dann −116 − 40 = −156, Überlauf, es bleibt 100. Zwei Überläufe heben sich auf, das Ergebnis 100 ist richtig. `x-40+20`: 80, dann 100, ohne Überlauf.
- Mit `int16` läuft in keinem Paar etwas über. Beide Ausdrücke liefern dann den gleichen Wert: 120, 80, 70 und 100.
- Die Division zweier ganzer Zahlen schneidet die Nachkommastellen ab.
- Ein gelber Kasten sagt: Ein richtiges Endergebnis beweist nicht, dass kein Teilausdruck übergelaufen ist.

### Handgriffe

- Eines der vier Paare wählen.
- Für `x` den Typ `int8` oder `int16` wählen.
- Die Auflösung der Frage aufklappen.
- Der Wert 120 ist fest. Eigene Ausdrücke lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.3.8
- 4.1

## Schritt 5: Was man dagegen tut

Reiter: vorbeugen

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Teil „Den passenden Typ wählen“ nennt `int` mit mindestens 32 Bit, also mindestens bis etwa ±2 Milliarden, und `int64` bis etwa ±9 · 10^18.
- Der Teil „Gefährdete Stellen erkennen“ nennt drei Stellen: das Produkt großer Zahlen, große Zwischenergebnisse wie die Summe für einen Durchschnitt und Ausdrücke aus mehreren Variablen.
- Der Teil „Vor der Rechnung prüfen“ zeigt ein Programm: `var a, b int8 = 120, 20`, dann `if a > math.MaxInt8-b`. Ist die Bedingung wahr, gibt das Programm „Die Summe passt nicht in int8.“ aus, sonst `a + b`. Mit 120 und 20 erscheint die Meldung. Mit 100 und 20 erscheint 120.
- Die Bedingung stellt `a + b > 127` so um, dass nichts überlaufen kann. Sie gilt für ein positives `b`. `math.MaxInt8` ist die Konstante 127.
- Ein gelber Kasten sagt: Im Kurstext steht nur, dass man vor der Addition prüfen kann. Das Programm ist eine Ergänzung des Mentors.
- Der Teil „Wenn kein Typ reicht“ nennt das Paket `math/big` aus dem Expertenwissen 4.15.
- Am Ende des Schritts unter „Wie es weitergeht“ steht ein Link auf die Seite „Funktionen in Go“ zu Kapitel 7. Dort wird aus der Prüfung die Funktion `safeAdd(a, b int32) (int32, error)`, die die Summe und einen Fehler zurückgibt.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Der Codeblock ist fest und lässt sich nicht ausführen.

### Abschnitte des Kurstexts

- 4.3.8
- 4.3.6

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Keine Meldung, 4.3.8, Seiten 78–79: Die vier Zeilen und die Meldung des Compilers stammen aus dem Kurstext. Die Werte 100, 175, 176 und 255 sind eigene Beispiele.
  - 2 Ohne Vorzeichen, 4.3.8, Seiten 79–80, mit Fußnote 25: Die Rechnungen 255 + 1, 130 + 140 und 255 + 255 stammen aus dem Kurstext. In Fußnote 25 wird 3 − 5 vom Überlauf unterschieden. Dass Go in beiden Fällen modulo 256 rechnet, ergänzt die Lektion.
  - 3 Mit Vorzeichen, 4.3.8, Seiten 80–81: Die sechs Rechnungen stammen aus dem Kurstext. Dort steht rechts ein Literal, zum Beispiel iA + 1. Die Demo rechnet mit zwei Variablen. Das Ergebnis ist das gleiche.
  - 4 Teilausdrücke, 4.3.8, Seiten 82–83: Expertenwissen 4.14: Die vier Zeilen stammen aus dem Kurstext. Die Auswertung Schritt für Schritt steht im Kurstext nur für x+20-40. Für die anderen drei Zeilen gibt die Lektion sie dazu.
  - 5 Was man dagegen tut, 4.3.8, Seiten 83–84: Expertenwissen 4.15: Das Expertenwissen 4.15 stammt aus dem Kurstext. Das Beispiel 100.000 · 100.000 und das Programm mit math.MaxInt8 gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, Teilausdrücke: Bei großen Zahlen kann a · c zu groß für den Wertebereich sein. Dann liefern (a − b) · c und a · c − b · c verschiedene Ergebnisse. Das ist der Gedanke von Schritt 4, dort in Go. Eine Lektion zu diesem Teil von Abschnitt 2.3 gibt es nicht.
  - Abschnitt 4.3.4 und 4.3.5, Seiten 63–66, Wertebereiche: Lektion: Zahlenkreis.
  - Abschnitt 7.1, Seiten 163–165, mehrere Rückgaben: Eine Funktion kann einen Fehler als zweite Rückgabe melden. Die Lektion Funktionen in Go zeigt in Schritt 4 die Funktion safeAdd, die so einen Überlauf meldet.
- Unter „Wie es weitergeht“ stehen Links zum Quiz zu den Wertebereichen und zu den Seiten „Zahlenkreis“, „Rechenwerk“, „Vertiefung: Blockchain in Go“ und „break-Anweisung in Go“. Der Exkurs zur Blockchain zeigt in Schritt 2 eine Hashfunktion, die den Überlauf von `uint32` bewusst einsetzt, und in Schritt 4 das Ergebnis von 75 − 76 im Typ `uint64`. Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links folgen. Alle führen auf eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.3
- 4.3.4
- 4.3.5
- 4.3.8
- 7.1
