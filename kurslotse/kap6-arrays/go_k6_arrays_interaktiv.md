# Arrays in Go

Die Seite zeigt, was ein Array in Go ist und wie man damit arbeitet: das Array-Literal, den Zugriff über den Index, das Durchlaufen mit der for-Anweisung mit range-Klausel, das zweidimensionale Array und einen Ausblick auf Slices. Die Seite gehört zu Kapitel 6 des Kurstexts, Abschnitte 6 bis 6.4.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Array ist eine Folge aus einer festen Zahl von Elementen, die alle Werte eines einzigen Typs sind.
- Der Typ eines Arrays besteht aus der Länge und dem Element-Typ, zum Beispiel `[6]int`. Die Länge gehört zum Typ. `[4]int` und `[6]int` sind verschiedene Typen.
- Die Länge muss beim Übersetzen feststehen. Sie ist eine Konstante und nie eine Variable.
- Arrays sind null-indiziert. Das erste Element hat den Index 0, das letzte den Index `len(a)-1`.
- `fmt.Println` gibt ein Array in eckigen Klammern aus, die Elemente durch Leerzeichen getrennt: `[3 17 24 31 36 47]`.
- Im Kurstext heißt es durchgehend „Array“. Das deutsche Wort „Feld“ wird dort nicht verwendet, weil es auch für die Felder eines Verbunds steht.
- Wie ein Array im Hauptspeicher liegt, zeigt eine andere HTML-Seite: „Speichermodell: Arrays“.
- Auf der Seite läuft kein Go. Der Browser stellt die Läufe nach. Alle Ausgaben und Meldungen sind mit Go 1.26 nachgeprüft.
- Schritt 1 bis 4 beginnen mit einem blauen Kasten mit Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und zwei Codeblöcken. Eine Demo gibt es hier nicht.
- Der erste Codeblock zeigt die Lottozahlen als sechs einzelne Variablen `z1` bis `z6`. Der zweite zeigt sie als Array: `var lotto [6]int = [6]int{3, 17, 24, 31, 36, 47}`, die Ausgabe ist `[3 17 24 31 36 47]`.
- Der Grund für das Array: Über einzelne Variablen kann keine Schleife laufen. Über die Elemente eines Arrays kann sie es, weil jedes Element einen Index hat.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 4.6.2.1, Abschnitt 5.1 mit 5.1.5 und Abschnitt 5.5.3.2. In Kapitel 2 und 3 kommen Arrays nicht vor.
- Eine zweite Liste sagt, was die fünf Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Strings: Bytes und Zeichen“, „Variablendeklaration in Go“, „for-Anweisung mit for-Klausel in Go“ und „Speichermodell: Arrays“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 6

## Schritt 1: Array und Literal

Reiter: literal

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen.
- Der Teil „Der Typ eines Arrays“ erklärt Länge und Element-Typ.
- Der Teil „Das Array-Literal“: Ein Array-Literal besteht aus dem Typ und den Werten in geschweiften Klammern. Es ist ein Ausdruck. Fehlende Werte bekommen den Null-Wert des Element-Typs: 0 bei Zahlen, `false` bei `bool`, der leere String bei `string`. Die Ellipse `...` anstelle der Länge lässt Go die Werte zählen.
- Die Demo zeigt neun Knöpfe mit je einer Deklaration. Die gewählte steht in einem dunklen Codeblock, mit `fmt.Println` und der Ausgabe hinter `//`.
- Darunter zeigt eine Reihe von Kästen das Array: in jedem Kasten ein Wert, darüber sein Index, links der Name der Variable.
- Ein Kasten darunter nennt Typ und Zahl der Elemente und den Grund. Übersetzt Go das Programm nicht, ist der Kasten rot und nennt die Meldung des Compilers. Die Reihe der Kästen fehlt dann.
- Go übersetzt: `var lotto [6]int = [6]int{3, 17, 24, 31, 36, 47}`. `farben := [4]string{"rot", "gelb", "grün", "blau"}` mit der Ausgabe `[rot gelb grün blau]`. `q := [5]float64{12.5, -8.3}` mit der Ausgabe `[12.5 -8.3 0 0 0]`. `var jahr [12]int` mit zwölf Nullen. `kurz := [...]int{3, 17, 24}` mit dem Typ `[3]int`. `const n = 5` mit `var a [n]int`, fünf Nullen.
- Go übersetzt nicht: `var b [6]int = [4]int{5, 8, 8, 2}` mit der Meldung `cannot use [4]int{…} (value of type [4]int) as [6]int value in variable declaration`. `a := [3]int{1, 2, 3, 4}` mit der Meldung `index 3 is out of bounds (>= 3)`. `n := 5` mit `var a [n]int` und der Meldung `invalid array length n`.
- Der Teil „Ein Wert je Zeile“ zeigt ein Literal über mehrere Zeilen. Dann steht auch hinter dem letzten Wert ein Komma. Ohne dieses Komma übersetzt Go das Programm nicht.
- Der Teil „Ein Literal mit Variablen“ zeigt: `r := [...]int{3, 4, s, 6}` mit `s` gleich −5 ergibt `[3 4 -5 6]`. Nach `s = 8` bleibt `r` unverändert, weil das Literal bei der Zuweisung einmal ausgewertet wird. Einer Array-Variable lässt sich jederzeit ein neues Literal zuweisen: `r = [...]int{0, 1, 0, s}` ergibt `[0 1 0 8]`.
- Ein Array der Länge 0 ist erlaubt und hat keinen praktischen Nutzen.

### Handgriffe

- Eine der neun Deklarationen wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Eine eigene Deklaration lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 6

## Schritt 2: Index

Reiter: index

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Der Text erklärt Index, Index-Ausdruck `a[i]` und indizierte Variable. Ein Codeblock zeigt das Beispiel des Kurstexts: `a := [...]int{11, 22, 33, 44}`, nach `a[i] = 99` mit `i` gleich 2 ist `a` gleich `[11 22 99 44]`, und `a[0] + a[3]` ist 55.
- `len(a)` liefert die Länge. `a[len(a)-1]` ist das letzte Element.
- Die Demo zeigt das Array `a` als vier Kästen mit dem Index darüber. Die Kästen sind Knöpfe. Das gewählte Element ist gelb, ein gerade geändertes Element grün.
- Darunter stehen ein Zahlenfeld „Neuer Wert“, ein Knopf, der dem gewählten Element den Wert zuweist, und der Knopf „von vorn“. Der Knopf trägt den gewählten Index, zum Beispiel „a[2] = Wert“.
- Zwei weitere Knöpfe zeigen einen Index außerhalb: `fmt.Println(a[4])` und `j := 4; fmt.Println(a[j])`.
- Ein dunkler Codeblock zeigt die Zeilen in Go mit ihren Ausgaben: die letzte Zuweisung, `a[i]`, das ganze Array, `len(a)` und `a[len(a)-1]`.
- Ein Kasten darunter erklärt den Stand. Bei einem Index außerhalb ist er rot.
- `fmt.Println(a[4])`: Der Index ist eine Konstante. Go übersetzt das Programm nicht. Die Meldung: `invalid argument: index 4 out of bounds [0:4]`.
- `j := 4` und `fmt.Println(a[j])`: Der Index ist eine Variable. Go übersetzt das Programm, es bricht zur Laufzeit ab. Die Meldung: `panic: runtime error: index out of range [4] with length 4`.
- Im Kurstext greifen die zwei falschen Zugriffe auf `lotto` zu, mit den Indizes 10 und 8. Die Meldung des Compilers endet im Kurstext mit `[0:4]`. Go 1.26 meldet für `lotto` mit sechs Elementen `[0:6]`.
- Der Teil „Index-Positionen im Literal“: `b := [8]int{5, 2: 8, 4: -2}` ergibt `[5 0 8 0 -2 0 0 0]`. `c := [...]int{4: 6}` ergibt `[0 0 0 0 6]`, das Array hat fünf Elemente.

### Handgriffe

- Ein Element des Arrays anklicken: Sein Index ist gewählt.
- Eine ganze Zahl von −999 bis 999 eintragen und den Knopf „a[i] = Wert“ wählen oder die Eingabetaste drücken.
- „von vorn“ wählen: Das Array enthält wieder 11, 22, 33 und 44, gewählt ist der Index 2.
- Einen der zwei Knöpfe mit einem Index außerhalb wählen.
- Die Auflösung der Frage aufklappen.
- Die Länge des Arrays lässt sich nicht ändern. Ein eigener Index außerhalb lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 6.1

## Schritt 3: Durchlaufen

Reiter: durchlaufen

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Der Text nennt das Muster der for-Klausel für ein Array: `i := 0`, `i < len(a)`, `i++`.
- Der Teil zur range-Klausel: `for i, v := range a`. Rechts von `range` steht das Array. Die erste Variable ist die Index-Variable vom Typ `int`. Die zweite hat den Element-Typ und bekommt vor jedem Durchlauf den Wert von `a[i]`. Beide gelten nur im Rumpf.
- Die Demo zeigt fünf Knöpfe für fünf Programme: „for-Klausel“, „for i, v := range a“, „for _, v := range a“, „for i := range a“ und „v = 99 im Rumpf“. Alle arbeiten mit `a := [...]int{8, -5, 14, 3}`.
- Das gewählte Programm steht als nummerierte Liste auf dunklem Grund. Die Zeile, in der der Lauf steht, ist hell unterlegt mit gelbem Strich.
- Darunter stehen drei Knöpfe: „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“.
- Das Array `a` steht als Reihe von Kästen da. Das Element, das gerade dran ist, ist gelb.
- Darunter stehen die Variablen als Kästen mit Name und Wert, zum Beispiel `sum`, `i` und `v`. Eine Variable, die sich gerade geändert hat, ist gelb.
- Ein dunkles Feld zeigt die Ausgabe des Programms. Ein Kasten darunter sagt, was die letzte Anweisung getan hat.
- Das Programm „for-Klausel“ und das Programm „for i, v := range a“ geben das Gleiche aus: `Index: 0 Wert: 8`, `Index: 1 Wert: -5`, `Index: 2 Wert: 14`, `Index: 3 Wert: 3` und `Summe aller Werte: 20`.
- Das Programm „for _, v := range a“ rechnet das Produkt: `Produkt aller Werte: -1680`.
- Das Programm „for i := range a“ kehrt mit `a[i] = -a[i]` jedes Vorzeichen um: `negiertes Array: [-8 5 -14 -3]`.
- Das Programm „v = 99 im Rumpf“ gibt viermal 99 aus und danach `[8 -5 14 3]`. Das Array ist unverändert, weil `v` eine Kopie des Elements ist. Wer ein Element ändern will, schreibt `a[i] = 99`.
- Der Teil „Welche Variable man weglassen kann“: Wird der Index nicht gebraucht, steht `_` an seiner Stelle, sonst meldet Go eine ungenutzte Variable: `declared and not used: i`. Wird der Wert nicht gebraucht, fällt die zweite Variable weg. `for range a` führt den Rumpf einmal je Element aus.
- Der letzte Teil nennt die for-range-Schleife über eine ganze Zahl seit Go 1.22: `for i := range 10` zählt von 0 bis 9.

### Handgriffe

- Eines der fünf Programme wählen. Der Lauf beginnt dann von vorn.
- „nächste Anweisung“ wählen: Der Lauf geht eine Anweisung weiter.
- „eine Anweisung zurück“ wählen.
- „von vorn“ wählen.
- Die Auflösung der Frage aufklappen.
- Die Werte des Arrays und die Programme lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 6.2
- 5.5.3.2

## Schritt 4: Zwei Dimensionen

Reiter: matrix

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Der Text: Ein Array hat eine Dimension. Ist der Element-Typ selbst ein Array-Typ, entsteht ein zweidimensionales Array. Es entspricht einer Matrix, ein einfaches Array einem Vektor.
- Ein Codeblock zeigt `a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}`, zeilenweise geschrieben. Die Ausgabe ist `[[1 2] [3 4] [5 6]]`.
- `[3][2]int` ist ein Array der Länge 3 mit dem Element-Typ `[2]int`: 3 Zeilen mit je 2 Spalten. `a[i][j]` ist das Element in Zeile `i` und Spalte `j`. `a[i]` ist die ganze Zeile `i`.
- Die erste Demo zeigt die Matrix `a` als Gitter mit „Zeile 0“ bis „Zeile 2“ links und „Spalte 0“ und „Spalte 1“ oben. Die Elemente sind Knöpfe. Das gewählte Element ist gelb, die übrigen Elemente seiner Zeile sind blau unterlegt. Ein Kasten darunter nennt `a[i][j]` mit Wert, `a[i]` als Zeile, `len(a)` gleich 3 und `len(a[i])` gleich 2.
- Die zweite Demo führt ein Programm mit der Matrix `b` vom Typ `[3][4]int` aus. Es setzt mit zwei verschachtelten for-range-Schleifen alle negativen Elemente auf 0 und gibt jedes Element mit `%3v` aus. Das Programm steht als nummerierte Liste mit 15 Zeilen da.
- Die Matrix `b` am Anfang: Zeile 0 ist −43, 12, 5, −63. Zeile 1 ist −19, −82, 91, 55. Zeile 2 ist 2, 46, −37, 28. Fünf Elemente sind negativ.
- Knöpfe der zweiten Demo: „nächstes Element“, „ein Element zurück“, „bis zum Ende“ und „von vorn“. Ein Handgriff bearbeitet ein Element: prüfen, wenn nötig auf 0 setzen, ausgeben.
- Das Element, das gerade bearbeitet ist, ist im Gitter gelb. Wurde es auf 0 gesetzt, ist es grün.
- Die Ausgabe am Ende, drei Zeilen: `0 12 5 0`, dann `0 0 91 55`, dann `2 46 0 28`, jede Zahl drei Zeichen breit.
- Ein weiterer Codeblock zeigt das Durchlaufen von `a` mit `for i, zeile := range a` und `for _, element := range zeile`. Die Ausgabe: `Zeile 0: 1 2`, `Zeile 1: 3 4`, `Zeile 2: 5 6`.
- Im Kurstext steht für `a` auch ein Programm mit zwei for-Klauseln und `len(a[i])`. Es steht nicht auf der Seite, liefert aber die gleiche Ausgabe.
- Ein Zugriff neben die Matrix übersetzt Go nicht, zum Beispiel `a[3][0]` mit der Meldung `invalid argument: index 3 out of bounds [0:3]`.

### Handgriffe

- In der ersten Demo ein Element der Matrix `a` anklicken.
- In der zweiten Demo „nächstes Element“, „ein Element zurück“, „bis zum Ende“ oder „von vorn“ wählen.
- Die Auflösung der Frage aufklappen.
- Die Werte der zwei Matrizen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 6.3

## Schritt 5: Slices

Reiter: slices

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text, einem Codeblock und einer Tabelle. Eine Demo gibt es hier nicht.
- Ein Slice ist, einfach gesagt, ein Array, dessen Länge zur Laufzeit wachsen kann. Der Typ hat leere eckige Klammern: `[]int`.
- Der Codeblock zeigt das Programm des Kurstexts: `zahlen := []int{1, 2, 3}`, `zahlen = append(zahlen, 4)`, `teil := zahlen[1:3]`. Die Ausgaben: `[1 2 3 4]`, `[2 3]`, `len(zahlen)` ist 4, `cap(zahlen)` ist 6.
- Die Tabelle vergleicht Array und Slice in drei Fragen: Schreibweise des Typs, wann die Länge feststeht, ob man ein Element anhängen kann.
- `append` mit einem Array übersetzt Go nicht. Die Meldung: `invalid append: argument must be a slice; have a (variable of type [4]int)`.
- Der Teil „Was dahinter steckt“: Ein Slice benutzt intern ein Array fester Größe. `len` liefert die Länge des Slice, `cap` die Kapazität, also die Größe des internen Arrays. Reicht der Platz nicht, legt Go ein größeres Array an und kopiert die Elemente.
- Die Zahl 6 für die Kapazität stammt aus einem Lauf mit Go 1.26. Im Kurstext steht keine Zahl. Wie viel Platz Go schafft, legt die Sprache nicht fest.
- Ein gelber Kasten sagt: Im Kurstext werden Slices nur in Abschnitt 6.4 gezeigt. Alle Beispiele des Kurses verwenden Arrays.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Der Codeblock ist fest und lässt sich nicht ausführen.

### Abschnitte des Kurstexts

- 6.4

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Array und Literal, 6, Seiten 151–154: Die ersten fünf Deklarationen der Demo stammen aus dem Kurstext. Die Deklaration von kurz steht dort als lotto mit sechs Werten. Die vier Deklarationen danach und die Meldungen des Compilers gibt die Lektion dazu.
  - 2 Index, 6.1, Seiten 154–155: Das Array mit 11, 22, 33 und 44 und die zwei falschen Zugriffe stammen aus dem Kurstext. Dort greifen die Zugriffe auf lotto zu, mit den Indizes 10 und 8. Im Kurstext endet die Meldung des Compilers mit [0:4]. Go 1.26 nennt für lotto die Länge 6: [0:6].
  - 3 Durchlaufen, 6.2, Seiten 156–159: Hinweis 6.1: Die ersten vier Programme der Demo stammen aus dem Kurstext. Zum fünften Programm mit v = 99 steht im Kurstext nur ein Satz.
  - 4 Zwei Dimensionen, 6.3, Seiten 160–162: Beide Matrizen und beide Programme stammen aus dem Kurstext. Die Demo gibt die Lektion dazu.
  - 5 Slices, 6.4, Seite 162: Das Programm stammt aus dem Kurstext. Die Ausgabe 6 für die Kapazität stammt aus einem Lauf mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, zusammengesetzte Typen: Aus mehreren Werten des gleichen Typs lässt sich ein neuer Typ bauen, zum Beispiel ein Vektor. Lektion: Variablen, Typen und Speicher, Schritt 3.
  - Abschnitt 4.6.2.1, Seiten 105–107, Index bei Strings: Auch ein String hat Elemente mit Index, seine Bytes. Lektion: Strings: Bytes und Zeichen.
  - Abschnitt 5.5.3.2, Seiten 142–146, for-Klausel: Mit ihr läuft der Index durch das Array. Lektion: for-Anweisung mit for-Klausel in Go.
  - Abschnitt 7.2, Seiten 169–170, Array als Argument: Eine Funktion bekommt eine Kopie des ganzen Arrays. Lektion: Das Kopier-Problem.
- Unter „Wie es weitergeht“ stehen Links zum Quiz zu den Arrays, zur Seite „Speichermodell: Arrays“ und zu zwei Programmierübungen zum Sortieren. Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links am Ende folgen. Alle führen auf eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.2
- 4.6.2.1
- 5.5.3.2
- 6.1
- 6.2
- 6.3
- 6.4
- 7.2
