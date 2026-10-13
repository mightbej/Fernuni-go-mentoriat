# Logische Operatoren und Vergleiche in Go

Die Seite zeigt die sechs Vergleichsoperatoren, den Vergleich von Strings und wie man mehrere Vergleiche mit `&&`, `||` und `!` zu einer Bedingung verknüpft. Sie gehört zu Abschnitt 4.5.2 des Kurstexts und zu Abschnitt 4.1.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Vergleich ist ein Ausdruck. Sein Wert hat immer den Typ `bool`, er ist `true` oder `false`. Ein Ausdruck vom Typ `bool` heißt boolescher Ausdruck.
- `==` und `!=` gibt es für alle primitiven Typen. `<`, `<=`, `>` und `>=` gibt es nur für geordnete Typen: Ganzzahlen, Gleitkommazahlen und `string`. Für `bool` gibt es sie nicht.
- Beide Seiten eines Vergleichs brauchen den gleichen Typ.
- `&&` ist die Konjunktion, `||` die Disjunktion, `!` die Negation.
- Bindungsstärke von stark nach schwach: einstellige Operatoren wie `!`, dann `*` und `/`, dann `+` und `-`, dann die sechs Vergleichsoperatoren, dann `&&`, dann `||`.
- Go wandelt einen Wahrheitswert nie von selbst in eine Zahl um und umgekehrt.
- Den Typ `bool`, die drei logischen Operatoren mit allen Kombinationen und die Kurzschluss-Auswertung zeigt eine andere HTML-Seite: „Wahrheitswerte und Kurzschluss-Auswertung“. Diese Seite hier erklärt die Kurzschluss-Auswertung nicht, sie nennt sie nur in Schritt 3.
- Auf der Seite läuft kein Go. Werte und Meldungen des Compilers sind mit Go 1.26 nachgeprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Codeblock: `var alter int = 17`, dann `alter >= 18` mit dem Wert `false` und `alter >= 6 && alter < 18` mit dem Wert `true`.
- Eine Liste nennt die drei logischen Operatoren in je einem Satz.
- Der Teil „Zwei Lektionen zu Abschnitt 4.5“ sagt, welche HTML-Seite was behandelt: diese die Vergleiche aus 4.5.2, die andere 4.5 und 4.5.1.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.1, Abschnitt 3.3.6 und Abschnitt 4.1.
- Eine weitere Liste sagt, was die drei Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Wahrheitswerte und Kurzschluss-Auswertung“ und „Imperative Programmierung“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.5
- 4.5.2

## Schritt 1: Sechs Operatoren

Reiter: operatoren

### Aufbau der Darstellung

- Oben steht eine Tabelle mit den sechs Operatoren: wann der Ausdruck `true` ist und für welche Typen es den Operator gibt.
- Der Teil „Drei Regeln“: Der Wert zählt und nicht die Schreibweise, `0xaf == 175` ist `true`. Erst wird gerechnet, dann verglichen, `10*5 == 70-20` ist `true`. Beide Seiten brauchen den gleichen Typ.
- Der Teil „== ist nicht =“: `a == 5` fragt, `a = 5` weist zu. Eine Zuweisung hat in Go keinen Wert und kann nicht als Bedingung stehen.
- Ein Codeblock zeigt die feste Belegung: `var a int = 2 + 3`, `var b int = 5`, `var s = "Hallo Welt"`, `var t = "Hallo " + "Welt"`.
- Die Demo zeigt zwölf Knöpfe mit je einem Ausdruck. Der gewählte Ausdruck steht in `fmt.Println(…)` in einem dunklen Codeblock, mit dem Wert hinter `//`.
- Darunter sagt ein Kasten den Wert und den Grund. Übersetzt Go die Zeile nicht, ist der Kasten rot und nennt die Meldung des Compilers.
- Neun Ausdrücke haben den Wert `true`: `a == b`, `s == t`, `33 == 33.0`, `0xaf == 175`, `10*5 == 70-20`, `1+1 != 3`, `8 >= 8`, `5 != 7 != (2 != 2)` und `4 > 2 == (5 < 9)`.
- Drei Ausdrücke übersetzt Go nicht. `4 > 2 == 5 < 9` mit der Meldung `invalid operation: 4 > 2 == 5 (mismatched types untyped bool and untyped int)`. `false < true` mit der Meldung `invalid operation: false < true (operator < not defined on untyped bool)`. `a == s` mit der Meldung `invalid operation: a == s (mismatched types int and string)`.
- Der Teil „bool ist keine Zahl“ nennt den Sprachvergleich 4.17: In C++ gilt 0 als `false` und jede andere Zahl als `true`. In Go schreibt man `if x != 0` statt `if x`.

### Handgriffe

- Einen der zwölf Ausdrücke wählen.
- Ein eigener Ausdruck lässt sich nicht eingeben. Die Belegung der Variablen ist fest.

### Abschnitte des Kurstexts

- 4.5.2
- 4.1

## Schritt 2: Strings vergleichen

Reiter: strings

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Die Regel in drei Punkten: Go vergleicht Zeichen für Zeichen von links. Das erste Paar verschiedener Zeichen entscheidet. Ist ein String vorher zu Ende, ist der längere der größere.
- Eine Tabelle nennt die Nummern der Zeichen: Leerzeichen 32, Ziffern 48 bis 57, Großbuchstaben 65 bis 90, Kleinbuchstaben 97 bis 122, Umlaute und ß über 122. Das ä hat 228, das Ä hat 196.
- Daraus folgt: Ziffern sind kleiner als Großbuchstaben, Großbuchstaben kleiner als Kleinbuchstaben.
- Die Demo hat zwei Eingabefelder `q` und `r` und den Knopf „vergleichen“. Darunter stehen fünf Knöpfe mit den Beispielen des Kurstexts und drei eigene Beispiele.
- Eine Tabelle legt die zwei Strings Zeichen für Zeichen untereinander. Unter jedem Zeichen steht klein seine Nummer. Gleiche Zeichen sind grau. Die Stelle, die entscheidet, ist gelb. Die Zeichen dahinter sind blass. Ist ein String zu Ende, steht dort ein Strich mit „zu Ende“.
- Ein dunkler Codeblock zeigt `q := …`, `r := …` und `fmt.Println(q == r, q < r, q > r)` mit den drei Werten hinter `//`.
- Ein Kasten nennt den Vergleich, der wahr ist, und die Stelle, die entscheidet.
- Die Ergebnisse der Beispiele: `"Maus" > "Katze"`, weil M die Nummer 77 hat und K die Nummer 75. `"Katzen" > "Katze"`, weil Katzen der längere String ist. `"affe" > "Ziege"`, weil a die Nummer 97 hat und Z die Nummer 90. `"sieben" > "8"`, weil s die Nummer 115 hat und 8 die Nummer 56. `"56" > "199"`, weil 5 die Nummer 53 hat und 1 die Nummer 49.
- Eigene Beispiele: `"Ärger" > "Zebra"`, weil Ä die Nummer 196 hat. `"Haus" < "haus"`, weil H die Nummer 72 hat und h die Nummer 104. `"Go" == "Go"`.
- Der Teil „Wo das vom Wörterbuch abweicht“: `"Zebra" < "apfel"` ist wahr. `"Ärger" > "Zebra"` ist wahr. `"10" < "9"` ist wahr, weil Zeichen verglichen werden und keine Zahlen.
- Im Kurstext steht die Regel ohne Nummern: Kleinbuchstaben werden höher gewertet als Großbuchstaben und diese höher als Ziffern.

### Handgriffe

- Für `q` und `r` je einen String mit höchstens 14 Zeichen eintragen und „vergleichen“ wählen oder die Eingabetaste drücken. Die Anführungszeichen gehören nicht in das Feld.
- Einen der acht Knöpfe mit Beispielen wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Den Operator kann man nicht wählen. Die Seite zeigt immer `==`, `<` und `>` zusammen.

### Abschnitte des Kurstexts

- 4.5.2
- 4.6.2

## Schritt 3: Vergleiche verknüpfen

Reiter: verknuepfen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und aufklappbarer Auflösung.
- Eine Liste zeigt drei Arten, einen Bereich zu prüfen: innerhalb mit `x >= 0 && x <= 9`, außerhalb mit `x < 0 || x > 9`, das Gegenteil mit `!(x >= 0 && x <= 9)`. Der Ausdruck mit `!` hat immer den gleichen Wert wie der Ausdruck mit `||`.
- Der Teil „Was zuerst gerechnet wird“ nennt die Bindungsstärken in fünf Stufen. Deshalb braucht `x >= 0 && x <= 9` keine Klammern. Das `!` bindet am stärksten und braucht die Klammern.
- Die Demo hat fünf Knöpfe für den Wert von `x`: −3, 0, 5, 9 und 12. Vier weitere Knöpfe wählen den Ausdruck: `x >= 0 && x <= 9`, `x < 0 || x > 9`, `!(x >= 0 && x <= 9)` und `0 <= x <= 9`.
- Ein dunkler Codeblock zeigt `var x int = …` und `fmt.Println(…)` mit dem Wert hinter `//`.
- Darunter stehen Kästen für die Teile: der linke Vergleich mit eingesetztem Wert und Ergebnis, der Operator, der rechte Vergleich. Beim Ausdruck mit `!` folgt ein Kasten, der das Ergebnis umkehrt.
- Steht das Ergebnis nach dem linken Vergleich schon fest, ist der Kasten des rechten Vergleichs gestrichelt und sagt „wird nicht mehr ausgewertet“. Das ist die Kurzschluss-Auswertung.
- Die Werte: `x >= 0 && x <= 9` ist `true` für 0, 5 und 9 und `false` für −3 und 12. Die zwei anderen Ausdrücke sind `true` für −3 und 12 und `false` für 0, 5 und 9.
- `0 <= x <= 9` übersetzt Go bei keinem Wert von `x`. Die Meldung: `invalid operation: 0 <= x <= 9 (mismatched types untyped bool and untyped int)`. Go rechnet von links, `0 <= x` ist ein Wahrheitswert, und der lässt sich nicht mit 9 vergleichen.
- Der Teil „Zwei Vergleiche hintereinander“ zeigt die zwei Zeilen des Kurstexts: `5 != 7 != (2 != 2)` und `4 > 2 == (5 < 9)`. Beide sind `true`, und beide brauchen die Klammern.
- Der Schritt endet mit Links zu zwei Quizseiten und zu den Seiten „Wahrheitswerte und Kurzschluss-Auswertung“, „if-Anweisung in Go“ und „for-Anweisung mit Bedingung in Go“. Welche Fragen die Quizseiten stellen, steht auf dieser Seite nicht.

### Handgriffe

- Einen der fünf Werte für `x` wählen.
- Einen der vier Ausdrücke wählen.
- Die Auflösung der Frage aufklappen.
- Eigene Werte und eigene Ausdrücke lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.1.2
- 4.5.2
- 4.5.1
