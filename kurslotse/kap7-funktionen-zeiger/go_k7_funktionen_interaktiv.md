# Funktionen in Go: Rückgaben und Wertübergabe

Die Seite zeigt, was in Abschnitt 7.1 und 7.2 des Kurstexts zu Funktionen dazukommt: Funktionen mit mehreren Rückgaben, benannte Rückgaben, die Signatur, das Rückgabe-Tupel, den Typ `error` und die Übergabe eines Arguments als Kopie. Sie gehört zu Kapitel 7 des Kurstexts, Seiten 163–170.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Eine Funktionsdeklaration besteht aus dem Schlüsselwort `func`, einem Funktionsnamen, einer Parameterliste und einer Rückgabeliste. Die Rückgabeliste darf fehlen.
- Eine Funktion mit mehreren Rückgaben heißt mehrwertige Funktion.
- Das Beispiel der Seite ist `divMitRest`. Die Funktion liefert den Ganzzahlquotienten `x / y` und den Divisionsrest `x % y`. `divMitRest(17, 3)` liefert 5 und 2.
- Parameter stehen in der Deklaration, Argumente stehen im Aufruf. Das steht in Abschnitt 3.4 des Kurstexts.
- In Go bekommt eine Funktion von jedem Argument eine Kopie des Werts. Das heißt call-by-value. Go kennt nur diese Art der Übergabe.
- Zeiger kommen auf dieser Seite nicht vor. Dafür gibt es die Lektion „Zeiger in Go“. Rekursion steht in der Lektion „Rekursion in Go“.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Meldungen sind mit Go 1.26 geprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Die Karte nennt drei Stücke, die zu den Funktionen aus Kapitel 3 dazukommen: mehrere Rückgaben, Fehler als Rückgabe, Wertübergabe.
- Unter „Was aus früheren Kapiteln dazugehört“ stehen vier Punkte mit Links: Abschnitt 2.6 mit der Lektion „Unterprogramme und Funktionen“, Abschnitt 3.4 mit zwei Quizseiten, Abschnitt 5.3.2 mit der Lektion „Zuweisungen in Go“, Kapitel 6 mit der Lektion „Arrays in Go“.
- Unter „Die Schritte“ steht je ein Satz zu den fünf Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 7.2
- 3.4

## Schritt 1: Mehrere Rückgaben

Reiter: rueckgaben

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter stehen die vier Teile einer Funktionsdeklaration und die Funktion `divMitRest` mit der Rückgabeliste `(int, int)` und `return x / y, x % y`.
- Unter „Benannte Rückgaben“ steht: Für jede benannte Rückgabe legt Go im Rumpf eine Variable an und gibt ihr den Null-Wert ihres Typs. Eine return-Anweisung ohne Ausdrücke heißt blankes return, englisch bare return.
- Die Demo zeigt ein Programm mit zehn Zeilen: die Funktion `divMitRest(x int, y int) (q int, r int)` mit `q = x / y`, `r = x % y` und `return`, darunter `main` mit `gquot, rest := divMitRest(17, 3)` und `fmt.Println(gquot, rest)`.
- Die Zeile, an der der Lauf steht, ist gelb markiert.
- Unter dem Programm stehen drei Knöpfe: „nächste Anweisung“, „eine Anweisung zurück“, „von vorn“.
- Darunter steht für jede laufende Funktion ein Rahmen mit ihren Variablen als Kästen. Ein gelber Kasten ist neu oder hat gerade einen Wert bekommen. Ein gestrichelter Rahmen gehört zu einer Funktion, die beendet ist.
- Darunter stehen die Ausgabe auf dunklem Grund und ein grüner Kasten mit einem Satz zur letzten Anweisung.
- Der Ablauf hat sechs Schritte: Aufruf mit `x` = 17, `y` = 3, `q` = 0, `r` = 0. Dann `q` = 5. Dann `r` = 2. Dann das blanke return. Dann in `main` `gquot` = 5 und `rest` = 2. Dann die Ausgabe `5 2`.
- Unter „Kürzer geschrieben“ steht `func divMitRest(x, y int) (q, r int)`.
- Ein gelber Kasten gibt Hinweis 7.1 des Kurstexts wieder: Namen für Rückgaben dokumentieren ihre Bedeutung. Ein blankes return soll selten stehen, weil man ihm nicht ansieht, welche Werte zurückgehen. Auch mit benannten Rückgaben darf man `return q, r` schreiben.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Die Argumente 17 und 3 lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 7.1
- 7.1.1

## Schritt 2: Signatur

Reiter: signatur

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter steht: Die Folge der Parameter-Typen und die Folge der Rückgabe-Typen bilden zusammen die Signatur einer Funktion. In Go heißt die Signatur auch Typ der Funktion. Es zählen nur die Typen und ihre Reihenfolge, nicht die Namen und nicht die Schreibweise.
- Die Demo zeigt sieben Funktionen als Knöpfe untereinander: `add`, `sub`, `div`, `divMitRest`, `zweimal`, `dreimal`, `wechsel`. Die gewählte Funktion hat einen grünen Rand.
- Ein grüner Kasten nennt zur gewählten Funktion die Parameter-Typen, die Rückgabe-Typen, den Typ der Funktion und die Funktionen mit der gleichen Signatur.
- Die Signaturen: `add` und `sub` haben `func(int, int) int`. `div` hat `func(int, int) float64`. `divMitRest` hat `func(int, int) (int, int)`. `zweimal` und `dreimal` haben `func(string, int) (string, int)`. `wechsel` hat `func(string, int) (int, string)`.
- `wechsel` hat die gleichen Typen wie `zweimal`, aber die Rückgabe-Typen stehen in anderer Reihenfolge. Die Signatur ist deshalb eine andere.
- Unter „Zwei Schreibweisen“ steht: Eine Funktion mit einer einzigen Anweisung im Rumpf darf in einer Zeile stehen. Die Klammern um eine einzelne Rückgabe ohne Namen dürfen fehlen.
- Die Schreibweise `func(int, int) int` steht nicht im Kurstext. Im Kurstext steht die Signatur nur in Worten. Dass eine Funktion einer Variable zugewiesen oder als Argument übergeben werden kann, kommt in Kapitel 7 des Kurstexts nicht vor.

### Handgriffe

- Eine der sieben Funktionen wählen.
- Die „Auflösung“ aufklappen.
- Eine eigene Funktion lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.1

## Schritt 3: Rückgabe-Tupel

Reiter: tupel

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter steht: Der Aufruf einer mehrwertigen Funktion liefert ein Tupel von Werten. Ein Tupel aus zwei Werten heißt Paar.
- Eine Liste nennt vier Arten, mit dem Paar umzugehen: weiterreichen, zuweisen, einen Wert mit dem leeren Bezeichner `_` verwerfen, alles verwerfen.
- Ein Codeblock zeigt die zwei Funktionen der Demo: `divMitRest(x, y int) (q, r int)` und `add(a, b int) int`.
- Die Demo zeigt acht Zeilen als Knöpfe. Darunter stehen die Ausgabe oder die Meldung des Compilers und ein Kasten mit einem Satz. Bei einer Zeile, die Go nicht übersetzt, ist der Kasten rot.
- Die acht Zeilen und ihr Ergebnis:
  1. `fmt.Println(divMitRest(17, 3))` gibt `5 2` aus.
  2. `gquot, rest := divMitRest(17, 3)` mit Ausgabe gibt `5 2` aus.
  3. `gquot, _ := divMitRest(50, 7)` mit Ausgabe gibt `7` aus. Der Rest 1 ist verworfen.
  4. `gquot, rest := divMitRest(divMitRest(100, 17))` gibt `0 5` aus. Der innere Aufruf liefert 5 und 15.
  5. `fmt.Println(add(divMitRest(17, 3)))` gibt `7` aus.
  6. `divMitRest(17, 3)` allein übersetzt Go, das Programm gibt nichts aus.
  7. `gquot := divMitRest(17, 3)` übersetzt Go nicht: `assignment mismatch: 1 variable but divMitRest returns 2 values`.
  8. `fmt.Println(1 + divMitRest(17, 3))` übersetzt Go nicht: `multiple-value divMitRest(17, 3) (value of type (q int, r int)) in single-value context`.
- Im Kurstext stehen die Zeilen 1 bis 4. Dort steht in Zeile 3 und 4 `=` statt `:=`, weil die Variablen schon deklariert sind.

### Handgriffe

- Eine der acht Zeilen wählen.
- Die „Auflösung“ aufklappen.
- Eine eigene Zeile lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.1.1
- 5.3.2

## Schritt 4: Fehler melden

Reiter: fehler

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter steht das Muster vieler Funktionen der Standardbibliothek: Die erste Rückgabe ist der Nutzwert, die zweite ist ein Wert vom Typ `error`. Der Wert `nil` bedeutet: kein Fehler. Ein `error` lässt sich mit `fmt.Println` ausgeben, ist aber kein `string`.
- Unter „Das Beispiel des Kurstexts: fmt.Scan“ steht ein Programm mit acht Zeilen. `n, err = fmt.Scan(&a, &b, &c)` liest drei ganze Zahlen. `n` zählt die gelesenen Eingaben, `err` ist der Fehler.
- Zwei Knöpfe wählen die Eingaben. Bei „6, -8, 23“ lautet die Ausgabe `Anzahl gültiger Eingaben: 3` und `Eingaben: 6 -8 23`. Bei „7, ups“ lautet sie `Fehler: expected integer`, `Anzahl gültiger Eingaben: 1` und `Eingaben: 7 0 0`.
- Unter „Eine eigene Funktion mit Fehler: safeAdd“ steht die Funktion `safeAdd(a, b int32) (int32, error)`. Sie liefert die Summe und `nil`, wenn die Summe in `int32` passt. Sonst liefert sie 0 und einen Fehler mit der Meldung „Summe ist zu groß für int32“ oder „Summe ist zu klein für int32“.
- Die erste Bedingung lautet `b > 0 && a > math.MaxInt32-b`, die zweite `b < 0 && a < math.MinInt32-b`. `math.MaxInt32` ist 2147483647, `math.MinInt32` ist −2147483648.
- Die Funktion prüft nicht `a + b > math.MaxInt32`, weil diese Summe selbst schon übergelaufen wäre. Die umgestellte Ungleichung rechnet mit einer Differenz, die immer in `int32` passt.
- Die Demo zu `safeAdd` zeigt sechs Wertepaare als Knöpfe. Darunter stehen die Ausgabe des aufrufenden Programms und ein Kasten, der sagt, welche Bedingung entschieden hat.
- Die sechs Paare: 17 und 4 ergibt Summe 21. 2000000000 und 147483647 ergibt Summe 2147483647. 2000000000 und 147483648 ergibt den Fehler „zu groß“, ohne Prüfung käme −2147483648 heraus. 2147483647 und 1 ergibt den Fehler „zu groß“. −2000000000 und −147483649 ergibt den Fehler „zu klein“, ohne Prüfung käme 2147483647 heraus. −5 und 3 ergibt Summe −2.
- Ein gelber Kasten sagt: `safeAdd` steht nicht im Kurstext, die Funktion ist ein Beispiel des Mentors. `errors.New` erzeugt aus einem String einen Wert vom Typ `error` und kommt im Kurstext nicht vor.
- Im Kurstext steht dazu noch das Expertenwissen 7.2: In einer kurzen Variablendeklaration mit mehreren Variablen darf links eine schon deklarierte Variable wie `err` stehen, wenn mindestens eine Variable neu ist. Die Seite zeigt das nicht.

### Handgriffe

- Eine der zwei Eingaben für `fmt.Scan` wählen.
- Eines der sechs Wertepaare für `safeAdd` wählen.
- Die „Auflösung“ aufklappen.
- Eigene Eingaben und eigene Werte lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.1.2
- 4.3.5

## Schritt 5: Kopie des Arguments

Reiter: kopie

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter steht: Eine Funktion bekommt von jedem Argument eine Kopie des Werts, das heißt call-by-value. Bei einem Array hat das zwei Folgen: Der Rechner kopiert bei jedem Aufruf das ganze Array, und die Funktion ändert nur ihre Kopie.
- Die Demo zeigt Listing 7.1 ohne die Zeilen mit `package` und `import`, mit 16 Zeilen: die Funktion `negToZeroValue(a [10]int) [10]int` in den Zeilen 1 bis 8 und `main` in den Zeilen 10 bis 16.
- Die Bedienung ist die gleiche wie in Schritt 1: drei Knöpfe, Rahmen je Funktion, Ausgabe, grüner Kasten.
- Ein Array steht in seinem Kasten als Reihe kleiner Felder. Ein grünes Feld ist das Element, das die letzte Anweisung geändert hat.
- Der Ablauf hat zwölf Schritte: `nums` anlegen mit `[8 -5 6 -3 -7 1 -4 3 -9 2]`. `nums` ausgeben. Aufruf: Der Parameter `a` bekommt eine Kopie aller zehn Elemente. Dann setzt die Schleife nacheinander `a[1]`, `a[3]`, `a[4]`, `a[6]` und `a[8]` auf 0. Dann `return a`. Dann legt `main` die Rückgabe in `posNums` ab, wieder als Kopie. Dann gibt `main` `nums` aus, unverändert. Dann gibt `main` `posNums` aus: `[8 0 6 0 0 1 0 3 0 2]`.
- Die Demo überspringt die Durchläufe der Schleife, in denen das Element nicht negativ ist.
- Unter „Was daran stört“ steht: Das Array wird zweimal kopiert. Sprachen wie Pascal haben Referenzparameter, die Übergabe heißt dort call-by-reference. Go hat sie nicht, in Go übergibt man einen Zeiger.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen vier Links: die Lektion „Zeiger in Go“, der Exkurs „Das Kopier-Problem“, die Programmierübung „Bubble Sort mit Funktionen“ und das Quiz zu Kapitel 7.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Array lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.2
- 7.3

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Mehrere Rückgaben, 7.1, Seiten 163–164: Die Funktion divMitRest stammt aus dem Kurstext. Die Funktion main setzt zwei Zeilen aus Abschnitt 7.1.1 ein. Die Demo gibt die Lektion dazu.
  - 2 Signatur, 7.1, Seiten 164–165: Die sieben Funktionen stammen aus dem Kurstext. Die Schreibweise des Typs, zum Beispiel func(int, int) int, steht nicht im Kurstext. So gibt fmt.Printf mit %T den Typ einer Funktion aus.
  - 3 Rückgabe-Tupel, 7.1.1, Seite 166: Die Zeilen 1 bis 4 stammen aus dem Kurstext. Dort steht in Zeile 3 und 4 eine Zuweisung mit =, weil die Variablen schon deklariert sind. Die Zeilen 5 bis 8 und die Meldungen des Compilers sind eine Ergänzung, geprüft mit Go 1.26.
  - 4 Fehler melden, 7.1.2, Seiten 166–168: Das Programm mit fmt.Scan und seine zwei Ausgaben stammen aus dem Kurstext. Die Funktion safeAdd ist eine Ergänzung, ausgeführt mit Go 1.26.
  - 5 Kopie des Arguments, 7.2, Seiten 169–170: Listing 7.1: Listing 7.1 stammt aus dem Kurstext. Die Lektion lässt die Zeilen mit package und import und die Kommentare weg.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.6, Seiten 15–17, Unterprogramme und Funktionen: Aufruf, Parameter und Rückgabe ohne Go. Lektion: Unterprogramme und Funktionen, Schritt 3.
  - Abschnitt 3.4, Seiten 37–41, Funktionen in einem Go-Programm: Parameter, Rückgabewert und return an Listing 3.3. Eine Lektion zu diesem Abschnitt gibt es nicht. Die Lektion Fehlersuche: drei Go-Programme mit mehreren Fehlern übt sie in Schritt 1.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Den Überlauf meldet safeAdd. Lektion: Überlauf bei Ganzzahlen in Go, Schritt 5.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.6
- 3.4
- 4.3.8
- 7.1
- 7.1.1
- 7.1.2
- 7.2
