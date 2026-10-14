# Zuweisungen in Go

Die Seite erklärt, wie ein Go-Programm einer Variable einen neuen Wert gibt: mit dem Zuweisungsoperator `=`, mit den kombinierten Operatoren wie `+=`, mit den Anweisungen `++` und `--` und mit der Tupel-Zuweisung. Der letzte Schritt erklärt den Begriff einfache Anweisung. Die Seite gehört zu den Abschnitten 5.3 und 5.4 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ergebnisse und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Eine Zuweisung wertet erst den Ausdruck auf der rechten Seite aus und legt dann den Wert in der Variable auf der linken Seite ab. Der alte Wert der Variable ist danach verloren.
- `=` ist der Zuweisungsoperator. `:=` ist der Deklarationsoperator und erzeugt eine neue Variable. `==` ist der Gleichheitsoperator und vergleicht zwei Werte.
- Ein Programm einer Demo steht in einem dunklen Kasten. Hat es Zeilennummern, stehen sie am linken Rand.
- Eine Meldung des Compilers steht in roter Schrift.
- Die Programme der Demos sind Ausschnitte ohne `package`, `import` und `func main`.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Codeblock zeigt zwei Zuweisungen: `x = 10` und `x = x + 12`. Nach der zweiten hat `x` den Wert 22.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.3 mit der Zuweisung als Pfeil `←`, Abschnitt 3.3.5 mit `tempC = 10` und `tempC = tempC + 12` aus Listing 3.2, Abschnitt 5.1.2 mit dem Unterschied zwischen `:=` und `=`.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zum Exkurs „Speichern und Rechnen im Rechner“ und zur Seite „Variablendeklaration“ folgen. Beide sind eigene HTML-Seiten.
- Der Codeblock lässt sich nicht ändern und nicht ausführen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 5.

### Abschnitte des Kurstexts

- 5.3
- 3.3.5
- 2.3

## Schritt 1: Zuweisung mit =

Reiter: zuweisung

### Aufbau der Darstellung

- Oben steht die Zuweisung `x = x + 5` mit ihren drei Teilen: links eine Variable, in der Mitte der Zuweisungsoperator `=`, rechts ein Ausdruck.
- Darunter steht, was Go tut: erst die rechte Seite auswerten, dann den Wert links ablegen.
- Der Teil „= ist kein Gleichheitszeichen“ sagt: Die Zuweisung hat eine Richtung von rechts nach links. Pseudocode schreibt dafür oft `←`, Pascal schreibt `:=`. Gleichheit prüft in Go der Operator `==`.
- Die Demo heißt „Ausprobieren: was Go annimmt“. Sie zeigt ein Programm mit drei Zeilen und Zeilennummern. Zeile 1 ist immer `var x int = 10`, Zeile 2 ist immer `var y int = 3`. Zeile 3 wählt der Student mit einem von sieben Knöpfen. Zeile 3 ist gelb markiert.
- Unter dem Programm stehen zwei Angaben: „Ergebnis“ oder „Meldung des Compilers“ und „Werte nach Zeile 3“. Darunter erklärt ein grüner Kasten Zeile 3.
- `x = 20`: Go übersetzt. Danach `x = 20, y = 3`.
- `x = x + 5`: Go übersetzt. Danach `x = 15, y = 3`.
- `x = y`: Go übersetzt. Danach `x = 3, y = 3`. Die Zuweisung kopiert den Wert, `y` bleibt unverändert.
- `x = 3.14`: Meldung `cannot use 3.14 (untyped float constant) as int value in assignment (truncated)`.
- `x = "20"`: Meldung `cannot use "20" (untyped string constant) as int value in assignment`.
- `20 = x`: Meldung `cannot assign to 20 (neither addressable nor a map index expression)`. Links muss eine Variable stehen.
- `x = y = 5`: Meldung `syntax error: unexpected = at end of statement`. Eine solche Kette von Zuweisungen gibt es in Go nicht, weil die Zuweisung eine Anweisung ist und keinen Wert hat.
- Bei einer Meldung läuft das Programm nicht. Bei „Werte nach Zeile 3“ steht dann, was nach Zeile 2 gälte: `x = 10, y = 3`.
- Unter der Demo steht der Teil „Wenn der Typ nicht passt“ mit `f = float64(x)` und zwei Links.
- Der Stoff dieses Schritts steht im Kurstext in Abschnitt 3.3.5 mit Hinweis 3.6. Abschnitt 5.3 setzt ihn voraus. Die sieben Zeilen der Demo und die Meldungen stehen nicht im Kurstext.

### Handgriffe

- Zeile 3 wählen: nur mit den sieben Knöpfen `x = 20`, `x = x + 5`, `x = y`, `x = 3.14`, `x = "20"`, `20 = x` und `x = y = 5`. Ein Eingabefeld gibt es nicht, andere Zeilen gehen nicht.
- Zeile 1 und 2 lassen sich nicht ändern.
- Den Links zu den Seiten „Typkonvertierung“ und „Konstanten“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 3.3.5
- 5.3

## Schritt 2: Kombinierte Operatoren

Reiter: kombiniert

### Aufbau der Darstellung

- Oben stehen drei Zuweisungen in langer Schreibweise und darunter die gleichen drei in kurzer Schreibweise: `x *= 8`, `y += 5`, `z -= 2 + 3`.
- Eine Tabelle stellt für die fünf Operatoren `*=`, `/=`, `%=`, `+=` und `-=` die kurze und die lange Schreibweise nebeneinander.
- Der Teil „Die rechte Seite gehört zusammen“ sagt: `z -= 2 + 3` entspricht `z = z - (2 + 3)`, mit Klammern. Go wertet erst die ganze rechte Seite aus.
- Die Demo heißt „Ausprobieren: kurz und lang“. `x` hat den Typ `int`. Sie hat drei Reihen von Knöpfen. Der gewählte Knopf ist grün.
- Erste Reihe, der Wert von `x` vorher: 20, 7 oder −9.
- Zweite Reihe, der Operator: `*=`, `/=`, `%=`, `+=` oder `-=`.
- Dritte Reihe, die rechte Seite: `3`, `4` oder `2 + 3`.
- Darunter stehen vier Angaben: „Kurze Schreibweise“, „Lange Schreibweise“, „Rechnung“ und „Wert von x danach“. Darunter erklärt ein grüner Kasten das Ergebnis.
- Beispiel: Mit 20, `*=` und `2 + 3` zeigt die Demo `x *= 2 + 3`, `x = x * (2 + 3)`, die Rechnung `20 * 5 = 100` und den Wert 100. Der grüne Kasten sagt dazu: Ohne die Klammern hieße die Zeile `x = x * 2 + 3`, und `x` hätte den Wert 43.
- Bei `/=` teilt Go ganze Zahlen ohne Rest und lässt die Nachkommastellen weg. 7 durch 4 ergibt 1. −9 durch 4 ergibt −2.
- Bei `%=` bekommt `x` den Rest der Division. 7 Rest bei Division durch 4 ist 3. Bei einer negativen Zahl links ist auch der Rest negativ: −9 Rest bei Division durch 4 ist −1.
- Unter der Demo steht der Teil „Welche kombinierten Operatoren es gibt“: zu jedem binären Operator der Bindungsstärken 5 und 4. Genannt sind `<<=` und `>>=` für die Bitverschiebung und `+=` für Strings: Aus "Regen" wird mit `wetter += "schauer"` der Wert "Regenschauer".

### Handgriffe

- Den Wert von `x` wählen: nur mit den Knöpfen 20, 7 und −9. Ein Eingabefeld gibt es nicht.
- Den Operator wählen: nur `*=`, `/=`, `%=`, `+=` und `-=`. Andere Operatoren gibt es in der Demo nicht.
- Die rechte Seite wählen: nur `3`, `4` und `2 + 3`. Eine Division durch 0 lässt sich nicht einstellen.
- Die Codeblöcke über der Demo sind fest und haben keine Knöpfe.
- Dem Link zur Seite „Bitverschiebung“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.3
- 4.1.2

## Schritt 3: ++ und --

Reiter: inkrement

### Aufbau der Darstellung

- Oben stehen die zwei Anweisungen. Die Inkrementierungs-Anweisung `x++` erhöht `x` um 1 und entspricht `x += 1`. Die Dekrementierungs-Anweisung `y--` verringert `y` um 1 und entspricht `y -= 1`.
- Der Teil „Anweisung und kein Ausdruck“ sagt: `+` und `-` sind Operatoren und bilden Ausdrücke mit einem Wert. `x++` ist eine Anweisung und hat keinen Wert.
- Die Demo heißt „Ausprobieren: was Go annimmt“. Sie zeigt ein Programm mit drei Zeilen und Zeilennummern. Zeile 1 ist immer `x := 10`, Zeile 2 ist immer `y := 0`. Zeile 3 wählt der Student mit einem von fünf Knöpfen. Zeile 3 ist gelb markiert.
- Unter dem Programm stehen zwei Angaben: „Ergebnis“ oder „Meldung des Compilers“ und „Werte nach Zeile 3“. Darunter erklärt ein grüner Kasten Zeile 3.
- `x++`: Go übersetzt. Danach `x = 11, y = 0`.
- `x--`: Go übersetzt. Danach `x = 9, y = 0`.
- `y = x++`: Meldung `syntax error: unexpected ++ at end of statement`.
- `y = 5 + x++`: die gleiche Meldung.
- `++x`: Meldung `syntax error: unexpected ++, expected }`. Die Form mit `++` vor der Variable gibt es in Go nicht.
- Unter der Demo steht der Teil „In anderen Sprachen“: In Java ist `x++` zugleich Anweisung und Ausdruck. Post-Increment `x++` liefert den alten Wert, Pre-Increment `++x` den neuen. Hat `x` den Wert 10, ergibt `y = x++` in Java `y = 10` und `y = ++x` ergibt `y = 11`. In beiden Fällen hat `x` danach den Wert 11.
- Der letzte Codeblock zeigt, wie man in Go das Gleiche wie `y = x++` schreibt: `y = x` und in der nächsten Zeile `x++`.

### Handgriffe

- Zeile 3 wählen: nur mit den fünf Knöpfen `x++`, `x--`, `y = x++`, `y = 5 + x++` und `++x`. Ein Eingabefeld gibt es nicht, andere Zeilen gehen nicht.
- Zeile 1 und 2 lassen sich nicht ändern.
- Der Codeblock am Ende ist fest und hat keine Knöpfe.

### Abschnitte des Kurstexts

- 5.3

## Schritt 4: Tupel-Zuweisung

Reiter: tupel

### Aufbau der Darstellung

- Oben stehen zwei Zeilen: `x, y, z := 10, 8, -5` ist eine Deklaration von drei Variablen, `x, y, z = 22, 33, 44` ist eine Tupel-Zuweisung an drei Variablen.
- Links und rechts vom `=` stehen gleich viele Einträge. Zu `x, y = 22, 33, 44` meldet der Compiler `assignment mismatch: 2 variables but 3 values`.
- Der Teil „Zwei Schritte“ sagt: Go wertet erst alle Ausdrücke rechts aus und merkt sich die Werte. Dann weist es die gemerkten Werte den Variablen links zu.
- Die Demo heißt „Ausprobieren: zwei Werte vertauschen“. `a` beginnt mit 3, `b` mit 8. Drei Knöpfe wählen die Fassung. Der gewählte Knopf ist grün.
- Der Student geht jede Fassung Schritt für Schritt durch. Über dem Programm stehen Kästen mit den Werten von `a` und `b`, in der zweiten Fassung auch von `tmp`.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wird oder wurde. Ein grüner Strich links bedeutet: Die Zeile ist schon gelaufen. Ein gelb unterlegter Kasten hat sich mit dem letzten Schritt geändert.
- Fassung „mit Tupel-Zuweisung“, zwei Zeilen: `a, b := 3, 8` und `a, b = b, a`. Der Lauf hat drei Schritte, weil Zeile 2 in ihre zwei Schritte zerlegt ist. Nach dem Auswerten zeigt ein gestrichelter Kasten „gemerkt“ die Werte 8, 3, und `a` und `b` sind noch unverändert. Nach dem Zuweisen gilt `a = 8, b = 3`.
- Der gestrichelte Kasten „gemerkt“ ist keine Variable des Programms. Er zeigt die Werte, die Go sich für diese eine Anweisung merkt.
- Fassung „mit Hilfsvariable tmp“, vier Zeilen: `a, b := 3, 8`, `tmp := a`, `a = b`, `b = tmp`. Der Lauf hat vier Schritte. Am Ende gilt `a = 8, b = 3`.
- Fassung „ohne Hilfsvariable“, drei Zeilen: `a, b := 3, 8`, `a = b`, `b = a`. Der Lauf hat drei Schritte. Am Ende gilt `a = 8, b = 8`. Diese Fassung vertauscht nicht: Der Wert 3 ist in Zeile 2 verloren gegangen. Go übersetzt sie trotzdem.
- Unter der Demo steht der Teil „Wann man die Tupel-Zuweisung nimmt“: wenn die Variablen zusammengehören und rechts einfache Ausdrücke stehen, und zum Vertauschen.
- Am Ende steht ein Link zur Seite „Unterprogramme und Funktionen“. Dort wird in Abschnitt 2.6 mit einer dritten Variable k vertauscht.

### Handgriffe

- Die Fassung wählen: nur „mit Tupel-Zuweisung“, „mit Hilfsvariable tmp“ und „ohne Hilfsvariable“. Eine neue Fassung beginnt von vorn.
- „Nächster Schritt“ führt den nächsten Schritt aus. „Von vorn“ beginnt die gewählte Fassung neu. Einen Schritt zurück gibt es nicht.
- Die Anfangswerte 3 und 8 und die Programmzeilen lassen sich nicht ändern.
- Dem Link zur Seite „Unterprogramme und Funktionen“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.3.2

## Schritt 5: Einfache Anweisungen

Reiter: einfach

### Aufbau der Darstellung

- Oben steht die Liste der fünf Arten von einfachen Anweisungen, englisch simple statements: kurze Variablendeklaration, Zuweisung, Inkrementierungs- und Dekrementierungs-Anweisung, Funktionsaufruf, leere Anweisung.
- Der Teil „Wozu man den Begriff braucht“ sagt: An manchen Stellen einer Kontrollfluss-Anweisung erlaubt Go nur eine einfache Anweisung. Das Beispiel ist die for-Klausel `for i := 1; i <= 10; i++ {`. Vor dem ersten Semikolon steht die Init-Anweisung `i := 1`, hinter dem zweiten die Post-Anweisung `i++`. Beide müssen einfache Anweisungen sein.
- Die for-Anweisung mit for-Klausel wird im Kurstext erst in Abschnitt 5.5.3.2 erklärt. Hier dient sie nur als Beispiel.
- Die Demo heißt „Ausprobieren: einfach oder nicht“. Acht Knöpfe zeigen acht Zeilen. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht die gewählte Zeile in einem dunklen Kasten. Darunter stehen zwei Angaben: „Einfache Anweisung“ mit ja oder nein und „Art der Anweisung“. Darunter erklärt ein grüner Kasten die Zuordnung.
- Einfache Anweisungen sind: `zahl := 20` (kurze Variablendeklaration), `x = 3 * y`, `x += 5` und `a, b = b, a` (alle drei Zuweisung), `i++` (Inkrementierungs-Anweisung), `quader(4, 7)` (Funktionsaufruf).
- Keine einfachen Anweisungen sind: `var zahl int = 20` (ausführliche Variablendeklaration) und `if x > 0 { x-- }` (bedingte Anweisung, eine Kontrollfluss-Anweisung).
- Zu `for var i int = 1; i <= 10; i++ {` meldet der Compiler `syntax error: var declaration not allowed in for initializer`.
- `quader` ist im Kurstext nur ein Beispielname für eine Funktion. Die Funktion selbst steht nicht auf der Seite.

### Handgriffe

- Eine der acht Zeilen wählen: nur mit den acht Knöpfen. Ein Eingabefeld gibt es nicht.
- Die Demo ordnet die gewählte Zeile nur ein. Sie führt keine Zeile aus und zeigt keine Werte. Ausprobieren lässt sich hier nur, welche Zeile eine einfache Anweisung ist.
- Der Codeblock mit der for-Klausel ist fest und hat keine Knöpfe. Er lässt sich nicht ausführen.
- Den Links zur Seite „for-Anweisung mit for-Klausel“ und zum Quiz folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.4
- 5.5.3.2

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Zuweisung mit =, 3.3.5, Seiten 33–35: Hinweis 3.6; 5.3, Seiten 126–128: Die Zuweisung und ihre Richtung stammen aus dem Kurstext. Die sieben Zeilen der Demo und die Meldungen des Compilers gibt die Lektion dazu.
  - 2 Kombinierte Operatoren, 5.3, Seiten 126–127: Die kombinierten Operatoren wie += stammen aus dem Kurstext. Die Werte der Demo und das Beispiel mit wetter sind eigene Beispiele.
  - 3 ++ und --, 5.3.1, Seiten 127–128: Sprachvergleich 5.7: Die Anweisungen ++ und -- und der Sprachvergleich stammen aus dem Kurstext. Die Meldungen des Compilers und die Werte zum Beispiel in Java gibt die Lektion dazu.
  - 4 Tupel-Zuweisung, 5.3.2, Seiten 128–129: Die Tupel-Zuweisung und das Vertauschen mit einer Hilfsvariable stammen aus dem Kurstext. Die Fassung ohne Hilfsvariable und die Meldung des Compilers gibt die Lektion dazu.
  - 5 Einfache Anweisungen, 5.4, Seite 129; 5.5.3.2, Seiten 142–146: Der Begriff der einfachen Anweisung stammt aus dem Kurstext. Die Zuordnung der acht Zeilen der Demo und die Meldung des Compilers gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, Zuweisung: Ein Ausdruck wird ausgewertet und sein Wert einer Variable zugewiesen. Lektion: Exkurs: Speichern und Rechnen im Rechner, Schritt 1.
  - Abschnitt 2.6, Seiten 15–17, Vertauschen: Zwei Werte mit einer dritten Variable vertauschen, ohne Go. Lektion: Unterprogramme und Funktionen, Schritt 1.
  - Abschnitt 4.6.2.2, Seiten 107–108, Verketten: a += "zeit". Lektion: Strings in Go, Schritt 4.
- Unter „Wie es weitergeht“ stehen Links auf: „Quiz zu den Zuweisungen“.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.3
- 2.6
- 3.3.5
- 4.6.2.2
- 5.3
- 5.3.1
- 5.3.2
- 5.4
- 5.5.3.2
