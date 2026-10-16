# Programmierübung: Insertion Sort Schritt für Schritt

Die Seite ist eine Programmierübung. Sie baut Insertion Sort in vier Stufen auf. Jede Stufe nimmt ein Mittel von Go dazu, in der Reihenfolge des Kurses: Tupel-Zuweisung, Array mit Index und zwei Schleifen, Funktion, Zeiger auf ein Array. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.3, Seiten 171–172, und Abschnitt 7.4, Seiten 173–178.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Insertion Sort kommt im Kurstext nicht vor. Die Übung ist eine Ergänzung des Mentors. Sie verwendet nur, was der Kurs bis Abschnitt 7.4 gezeigt hat.
- Der Kurstext sagt nichts zu Insertion Sort und nichts zu seiner Laufzeit, auch nicht in Abschnitt 7.3, 7.4 oder 8.4.2. Abschnitt 7.3 und 7.4 erklären nur den Zeiger als Parameter. Eine Antwort zur Laufzeit oder zum Vergleich mit Bubble Sort nennt deshalb keinen Abschnitt des Kurstexts als Quelle. Sie nennt die Zahlen dieser Seite und verweist auf den Exkurs zur Laufzeit.
- Das Verfahren: Links steht ein sortierter Teil, zu Beginn nur das erste Element. Das nächste Element wird von rechts mit dem sortierten Teil verglichen. Jedes größere Element rückt einen Platz nach rechts. Steht links ein kleineres oder gleiches Element oder ist der Anfang erreicht, kommt das neue Element in die Lücke. So wie man Spielkarten in der Hand einsortiert.
- Ab Stufe 2 heißen die Variablen wie in Thomas' Vorlage: `i` für das Element, das eingefügt wird, `key` für seinen Wert, `j` für den Index, mit dem die innere Schleife vergleicht. `i` und `j` sind Indizes und keine Zeiger. Der einzige Zeiger der Übung ist `pa` in Stufe 4.
- Die innere Schleife heißt `for j >= 0 && a[j] > key`. `j >= 0` muss links stehen: Wegen der Kurzschluss-Auswertung aus Abschnitt 4.5.1 wertet Go `a[j]` nicht aus, wenn `j` gleich −1 ist. Umgedreht bricht das Programm ab mit `panic: runtime error: index out of range [-1]`, mit Go 1.26.1 geprüft.
- Die Werte: In Stufe 1 bis 3 sind es die fünf Werte 42, 7, 19, 88, 3. In Stufe 4 ist es das Array der Bubble-Sort-Übungen, `[10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`. Die ersten fünf Werte sind die gleichen.
- Zahl der Vergleiche `a[j] > key`, mit Go 1.26.1 gezählt: fünf Werte der Übung 8, zehn Werte der Übung 30, zehn sortierte Werte 9, zehn umgekehrt sortierte Werte 45. Stufe 1 vergleicht immer zehnmal.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Die Demos stellen im Browser nach, was das Programm tut. Programm und Werte der Demos lassen sich nicht ändern. Ändern und ausführen lässt sich ein Programm nur in Schritt 5, wenn der Editorlotse die Lektion ausliefert. Schlägt eine Antwort eine Änderung vor, verweist sie auf Schritt 5.
- Die Demos in Schritt 1, 2 und 4 stehen in zwei Spalten. Links steht das Programm in einem Kasten, der von selbst zur hervorgehobenen Zeile rollt, lange Zeilen brechen dort um. Rechts stehen die Knöpfe, darunter der Satz zum letzten Schritt, die Rahmen mit den Variablen und die Ausgabe. Ist die Seite schmal, zum Beispiel auf einem Telefon, steht das alles unter dem Programm.
- Die Demos zeigen keine Zahl der Vergleiche. Mitgezählt wird nur im Programm „Stufe 4 mit Ausgabe jedes Einfügens“ in Schritt 5.
- Fragt der Student nach der Lösung einer eigenen Aufgabe, zum Beispiel absteigend sortieren, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Der Text führt Insertion Sort mit den Spielkarten in der Hand ein und nennt in einer Liste: den sortierten Teil, das Einfügen, das Ende.
- Fünf Reihen mit Kärtchen zeigen die fünf Werte vor dem Sortieren und nach jedem Einfügen. Grün ist der sortierte Teil, gelb das gerade eingefügte Element. Vorher: 42 7 19 88 3. Nach „7 einfügen“: 7 42 19 88 3. Nach „19 einfügen“: 7 19 42 88 3. Nach „88 einfügen“: 7 19 42 88 3. Nach „3 einfügen“: 3 7 19 42 88.
- Eine Liste nennt die vier Stufen mit ihrem Mittel von Go und dem Abschnitt des Kurstexts.
- Ein gelber Kasten sagt: Insertion Sort kommt im Kurstext nicht vor.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3

## Schritt 1: Fünf Variablen

Reiter: variablen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit zwei Fragen und aufklappbaren Auflösungen: wie viele Vergleiche das Programm ausführt (immer zehn), und welcher Vergleich beim Einfügen von `d` genügt hätte (der erste, `d < c`).
- Die Demo zeigt `main` mit 42 Zeilen, die Stufe aus Thomas' Vorlage. Für `b`, `c`, `d` und `e` steht je ein Kommentar und dann für jeden linken Nachbarn ein `if` mit Tausch, zum Beispiel `if c < b { b, c = c, b }`. Nach jedem Einfügen gibt `fmt.Println` die Werte aus.
- Die Ausgabe:
  - `Vorher:      42 7 19 88 3`
  - `b eingefügt: 7 42 19 88 3`
  - `c eingefügt: 7 19 42 88 3`
  - `d eingefügt: 7 19 42 88 3`
  - `e eingefügt: 3 7 19 42 88`
- Getauscht wird sechsmal: einmal bei `b`, einmal bei `c`, nie bei `d`, viermal bei `e`.
- Unter „Was dieser Stufe fehlt“: Das Programm vergleicht jedes neue Element bis zum Anfang, auch wenn es schon steht. Für zehn Werte bräuchte es 45 `if`-Anweisungen.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.

### Abschnitte des Kurstexts

- 5.3

## Schritt 2: Ein Array

Reiter: array

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen: warum die innere Schleife bei `i` gleich 3 nicht läuft (42 ist nicht größer als 88), und was passiert, wenn die Bedingung umgedreht wird (Abbruch mit `index out of range [-1]`, Kurzschluss-Auswertung).
- Eine Liste erklärt die äußere Schleife, die innere Schleife und das Einfügen. Neu gegenüber Stufe 1: Das Programm verschiebt statt zu tauschen. Während die innere Schleife läuft, steht ein Wert zweimal im Array, `key` hält den eingefügten Wert.
- Die Demo zeigt `main` mit 14 Zeilen: `a := [5]int{42, 7, 19, 88, 3}`, die äußere Schleife `for i := 1; i < len(a); i++`, darin `key := a[i]`, `j := i - 1`, die innere Schleife `for j >= 0 && a[j] > key` mit `a[j+1] = a[j]` und `j--`, danach `a[j+1] = key`.
- Im Rahmen von `main` stehen `a`, `i`, `key` und `j`. Grün ist der sortierte Teil, so weit er vor dem Einfügen reichte. Gelb ist das Element am Index `j` oder der Platz, an den `key` kommt.
- Die Ausgabe: `Vorher:  [42 7 19 88 3]` und `Nachher: [3 7 19 42 88]`.
- Unter der Demo: Links auf die Lektionen „for-Anweisung mit Bedingung in Go“ und „Wahrheitswerte und Kurzschluss-Auswertung“. Acht Vergleiche statt zehn.
- Unter „Was dieser Stufe fehlt“: Das Sortieren steht mitten in `main`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.

### Abschnitte des Kurstexts

- 6.1
- 6.2
- 5.5.3.1
- 4.5.1

## Schritt 3: Eine Funktion

Reiter: funktion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage: Was gibt `main` aus, wenn es nur `insertionSort(a)` aufruft und dann `a` ausgibt? Auflösung: das unveränderte Array.
- Darunter steht die Funktion `insertionSort(a [5]int) [5]int`. Sie sortiert ihre Kopie und liefert sie zurück.
- Der Text erklärt die Wertübergabe aus Abschnitt 7.2: Die Funktion bekommt eine Kopie des Arrays.
- Eine Auswahl hat drei Möglichkeiten für die zwei Zeilen in `main`:
  - `b := insertionSort(a)` und `fmt.Println(a, b)`. Ausgabe: `[42 7 19 88 3] [3 7 19 42 88]`.
  - `insertionSort(a)` und `fmt.Println(a)`. Ausgabe: `[42 7 19 88 3]`. Go übersetzt den Aufruf, obwohl niemand das Ergebnis auffängt.
  - `a = insertionSort(a)` und `fmt.Println(a)`. Ausgabe: `[3 7 19 42 88]`.
- Unter „Was dieser Stufe fehlt“ stehen zwei Punkte: Jeder Aufruf kopiert das Array hin und zurück, mit Link auf „Das Kopier-Problem“. Die Funktion sortiert nur den Typ `[5]int`. Die feste Länge bleibt auch in Stufe 4, dort `[10]int`.

### Handgriffe

- Eine der drei Möglichkeiten wählen.
- Die Auflösung aufklappen.

### Abschnitte des Kurstexts

- 7.1
- 7.2

## Schritt 4: Ein Zeiger

Reiter: zeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen: in welcher Runde ein Element am weitesten wandert (bei `i` gleich 9 rückt die 9 an sieben Elementen vorbei zum Index 2), und wie viele Vergleiche ein schon sortiertes Array braucht (neun).
- Der Text erklärt den Parameter `pa *[10]int`, den Aufruf `insertionSort(&a)`, die Kurzschreibweise `pa[i]` statt `(*pa)[i]` und `len(pa)` mit einem Zeiger auf ein Array.
- Die Demo zeigt 19 Zeilen: `insertionSort` in den Zeilen 1 bis 12, `main` in den Zeilen 14 bis 19. Ein Schritt der Demo ist eine ganze Runde der äußeren Schleife, hervorgehoben ist dann Zeile 10, `pa[j+1] = key`.
- Im Rahmen von `insertionSort` stehen `pa` als „Adresse von a“, `i`, `key` und `j` nach der Runde. Im Array `a` von `main` ist grün der sortierte Teil, gelb der Platz, an den `key` gekommen ist.
- Die neun Runden: 7 an Index 0, 19 an Index 1, 88 bleibt an Index 3, 3 an Index 0, 25 an Index 3, 71 an Index 5, 14 an Index 2, 60 an Index 6, 9 an Index 2.
- Die Ausgabe: `Start:    [42 7 19 88 3 25 71 14 60 9]` und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Unter „Wie schnell ist Insertion Sort?“: 30 Vergleiche für das Array der Übung, 9 für ein sortiertes, 45 für ein umgekehrt sortiertes Array. Im schlechtesten Fall wächst die Zahl mit dem Quadrat der Anzahl der Elemente, wie bei Bubble Sort. Link auf den Exkurs zur Laufzeit. Der Kurstext sagt zur Laufzeit von Insertion Sort nichts. Eine Antwort dazu nennt keinen Abschnitt des Kurstexts als Quelle, sondern den Exkurs zur Laufzeit.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.
- Den Links folgen.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 5: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Im Editorlotsen steht ein Editor mit dem Knopf „Ausführen“. Auf GitHub Pages ist der Reiter grau. Dort wählt der Student ein Programm und sieht das Programm und seine Ausgabe, festgehalten mit Go 1.26.1.
- Sieben vollständige Programme: „Stufe 1: fünf Variablen“, „Stufe 1 mit der Funktion ordne“, „Stufe 2: Array und Schleifen“, „Stufe 3: b := insertionSort(a)“, „Stufe 3: Ergebnis nicht aufgefangen“, „Stufe 4: Zeiger“, „Stufe 4 mit Ausgabe jedes Einfügens“.
- „Stufe 1 mit der Funktion ordne“ ist eine Zwischenstufe aus Thomas' Vorlage: `ordne(x, y int) (int, int)` liefert die zwei Werte geordnet, `main` ruft sie zehnmal auf, zum Beispiel `b, c = ordne(b, c)`.
- „Stufe 4 mit Ausgabe jedes Einfügens“ gibt nach jeder Runde `i`, `key` und das Array aus und am Ende `Vergleiche: 30`. Damit es jeden Vergleich zählen kann, steht der Vergleich dort in `if pa[j] <= key { break }`.
- Vier Vorschläge: die Bedingung umdrehen (Abbruch mit `index out of range [-1]`), absteigend sortieren (in der inneren Schleife `a[j] < key` statt `a[j] > key`), Vergleiche für ein sortiertes und ein umgekehrt sortiertes Array zählen (9 und 45), `pa[j] >= key` statt `pa[j] > key` (das Ergebnis bleibt richtig, gleiche Werte rücken jetzt aneinander vorbei).

### Handgriffe

- Ein Programm wählen, im Editor ändern und ausführen, wenn der Editorlotse die Lektion ausliefert.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle nennt zu jedem Schritt den Abschnitt im Kurstext, was daraus stammt und was nicht im Kurstext steht: 5.3 Tupel-Zuweisung, 4.5.1 Kurzschluss-Auswertung, 5.5.3 `for`-Anweisung, 6.1 und 6.2 Array und Durchlaufen, 7.1 und 7.2 Funktion und Wertübergabe, 7.3 und 7.4 Zeiger auf ein Array.
- Unter „Wo das Thema im Kurstext noch vorkommt“: Abschnitt 4.5.1 Kurzschluss-Auswertung, Abschnitt 5.5.3.1 `for` mit Bedingung, Bubblesort in Abschnitt 8.4.2, je mit Lektion.
- Unter „Wie es weitergeht“: die Programmierübungen „Bubble Sort mit Zeigern“, „Selection Sort mit Zeigern“ und „Quicksort Schritt für Schritt“, der Exkurs zur Laufzeit und das Quiz zu Funktionen, Zeigern und Rekursion.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
