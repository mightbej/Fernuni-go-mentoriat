# Programmierübung: Quicksort Schritt für Schritt

Die Seite ist eine Programmierübung. Sie baut Quicksort in fünf Stufen auf. Jede Stufe nimmt ein Mittel von Go dazu, in der Reihenfolge des Kurses: Tupel-Zuweisung, Array mit Index und Schleife, Funktion, Zeiger auf ein Array mit Bereich, Rekursion. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.5, Seiten 185–193.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Quicksort kommt im Kurstext nicht vor. Die Übung ist eine Ergänzung des Mentors. Sie verwendet nur, was der Kurs bis Abschnitt 7.5 gezeigt hat.
- Das Verfahren: Ein Element wird zum Vergleichswert, dem Pivot. Das Aufteilen, englisch partition, bringt alle kleineren Elemente links neben das Pivot und alle größeren oder gleichen rechts daneben. Danach steht das Pivot an seinem endgültigen Platz. Der linke und der rechte Teil werden auf die gleiche Weise sortiert, bis jeder Teil höchstens ein Element hat.
- Das Pivot ist immer das erste Element des Bereichs.
- Das Aufteilen der Übung arbeitet mit einem einzigen Index `left`. `left` ist eine Grenze: Links davon, ab dem Element hinter dem Pivot, stehen nur Elemente, die kleiner sind als das Pivot. Die Schleife geht mit `i` über den Bereich. Ist `a[i]` kleiner als das Pivot, tauscht es mit dem Element am Index `left`, und `left` rückt eins weiter. Nach der Schleife tauscht das Pivot mit dem Element am Index `left-1`. Die Funktion liefert `left-1`, den Index des Pivots.
- `left` ist ein Index und kein Zeiger. Der einzige Zeiger der Übung ist `pa` ab Stufe 4.
- Die Werte: In Stufe 1 bis 3 sind es die fünf Werte 42, 7, 19, 88, 3. Ab Stufe 4 ist es das Array der Bubble-Sort-Übungen, `[10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`. Die ersten fünf Werte sind die gleichen.
- Die Funktionen ab Stufe 4: `swap(pa *[10]int, i, j int)` tauscht zwei Elemente. `partition(pa *[10]int, start, end int) int` teilt den Bereich von `start` bis `end` auf, beide gehören dazu. `quickSort(pa *[10]int, start, end int)` sortiert den Bereich.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Die Demos in Schritt 1, 2, 4 und 5 haben oben zwei Spalten. Links steht das Programm in einem Kasten, der zur hervorgehobenen Zeile rollt, lange Zeilen brechen dort um. Rechts stehen die Knöpfe, darunter der Satz zum letzten Schritt und die Ausgabe. Die Rahmen mit den Variablen stehen darunter über die volle Breite. Ist die Seite schmal, zum Beispiel auf einem Telefon, steht alles untereinander.
- Die Demos stellen im Browser nach, was das Programm tut. Ausführen lässt sich ein Programm nur in Schritt 6, wenn der Editorlotse die Lektion ausliefert.
- Fragt der Student nach der Lösung einer eigenen Aufgabe, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Der Text führt das Pivot und das Aufteilen ein. Pivot ist das englische Wort für Drehpunkt.
- Zwei Reihen mit Kärtchen zeigen das Array vor und nach dem ersten Aufteilen um 42. Vorher: 42 7 19 88 3 25 71 14 60 9. Nachher: 9 7 19 3 25 14 42 88 60 71. Gelb ist das Pivot, grün sind die kleineren Werte, blau die größeren.
- Danach sind zwei kleinere Aufgaben übrig: den linken und den rechten Teil sortieren.
- Eine Liste nennt die fünf Stufen mit ihrem Mittel von Go und dem Abschnitt des Kurstexts.
- Ein gelber Kasten sagt: Quicksort stammt von Tony Hoare, 1959, und kommt im Kurstext nicht vor.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5

## Schritt 1: Fünf Variablen

Reiter: variablen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit zwei Fragen und aufklappbaren Auflösungen: welche Werte links und rechts von 42 landen, und warum der Tausch eine Tupel-Zuweisung braucht.
- Die Demo zeigt `main` mit elf Zeilen. Zeile 2: `a, b, c, d, e := 42, 7, 19, 88, 3`. Zeile 7: `d, e = e, d`. Zeile 9: `a, d = d, a`. Die Zeilen 4 bis 6 und 8 sind Kommentare mit den Entscheidungen.
- Die Demo hat fünf Schritte. Danach gilt a = 3, b = 7, c = 19, d = 42, e = 88. Die Ausgabe: `Vorher:  42 7 19 88 3` und `Nachher: 3 7 19 42 88`.
- Die fünf Werte sind nach dem Aufteilen zufällig schon sortiert. Im Allgemeinen sortiert das Aufteilen allein nicht.
- Unter „Was dieser Stufe fehlt“: Die Entscheidungen stehen nur in den Kommentaren, das Programm vergleicht nichts. Für andere Werte müsste man andere Tausche hinschreiben.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.

### Abschnitte des Kurstexts

- 5.3

## Schritt 2: Ein Array

Reiter: array

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage: Was bedeutet der Wert 4 von `left` am Ende der Schleife? Auflösung: An den Indizes 1 bis 3 stehen die drei kleineren Elemente, das Pivot gehört an den Index `left-1`, also 3.
- Die Demo zeigt `main` mit 15 Zeilen: `a := [5]int{42, 7, 19, 88, 3}`, `pivot := a[0]`, `left := 1`, die Schleife `for i := 1; i < len(a); i++` mit `if a[i] < pivot`, dem Tausch `a[i], a[left] = a[left], a[i]` und `left++`, danach `a[0], a[left-1] = a[left-1], a[0]` und zwei Ausgaben.
- Im Rahmen von `main` stehen `a`, `pivot`, `left` und `i`. Grün ist im Array, was links von der Grenze `left` steht. Gelb ist das Element am Index `i`. Nach dem letzten Tausch ist das Pivot gelb und die kleineren Elemente an den Indizes 0 bis 2 sind grün.
- Der Ablauf: i = 1, 7 < 42, tauscht mit sich selbst, left = 2. i = 2, 19 < 42, tauscht mit sich selbst, left = 3. i = 3, 88 nicht kleiner. i = 4, 3 < 42, tauscht mit a[3], das Array ist `[42 7 19 3 88]`, left = 4. i = 5, die Schleife endet. Das Pivot tauscht mit a[3].
- Die Ausgabe: `Vorher:  [42 7 19 88 3]`, `Nachher: [3 7 19 42 88]`, `Das Pivot steht am Index 3`.
- Unter „Was dieser Stufe fehlt“: Das Aufteilen steht mitten in `main` und müsste für jeden Teil noch einmal dastehen.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösung aufklappen.

### Abschnitte des Kurstexts

- 6.1
- 6.2

## Schritt 3: Eine Funktion

Reiter: funktion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage: Was gibt `main` aus, wenn es nur `partition(a)` aufruft und dann `a` ausgibt? Auflösung: das unveränderte Array.
- Darunter steht die Funktion `partition(a [5]int) ([5]int, int)`. Sie liefert das aufgeteilte Array und den Index des Pivots.
- Der Text erklärt die Wertübergabe aus Abschnitt 7.2: Die Funktion bekommt eine Kopie des Arrays.
- Eine Auswahl hat drei Möglichkeiten für die zwei Zeilen in `main`:
  - `b, p := partition(a)` und `fmt.Println(a, b, p)`. Ausgabe: `[42 7 19 88 3] [3 7 19 42 88] 3`.
  - `partition(a)` und `fmt.Println(a)`. Ausgabe: `[42 7 19 88 3]`. Go übersetzt den Aufruf, obwohl niemand die Ergebnisse auffängt.
  - `a, p := partition(a)` und `fmt.Println(a, p)`. Ausgabe: `[3 7 19 42 88] 3`. `:=` ist erlaubt, weil `p` neu ist. `a` bekommt nur einen neuen Wert.
- Unter „Was dieser Stufe fehlt“ stehen zwei Punkte: Jeder Aufruf kopiert das Array hin und zurück, mit Link auf die Lektion „Das Kopier-Problem“. Die Funktion teilt immer das ganze Array auf und keinen Bereich.

### Handgriffe

- Eine der drei Möglichkeiten wählen.
- Die Auflösung aufklappen.

### Abschnitte des Kurstexts

- 7.1
- 7.2

## Schritt 4: Zeiger und Bereiche

Reiter: zeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen: wie viele Elemente kleiner als 42 sind (sechs: 7, 19, 3, 25, 14, 9), und warum das Array nach drei Aufrufen noch nicht sortiert ist.
- Der Text nennt zwei Änderungen: ein Zeiger `pa *[10]int` und ein Bereich mit `start` und `end`. Das Tauschen steht in `swap`, wie in der Programmierübung „Bubble Sort mit Zeigern“.
- Darunter stehen `swap` und `partition`. Die Schleife läuft bis `i <= end`, weil `end` zum Bereich gehört.
- Die Demo zeigt `main` mit zehn Zeilen. `main` ruft `partition` dreimal auf: `partition(&a, 0, 9)`, `partition(&a, 0, p-1)`, `partition(&a, p+1, 9)`. Ein Schritt der Demo ist ein ganzer Aufruf.
- Im Array ist gelb der Bereich des letzten Aufrufs, grün sind die Pivots, deren Platz feststeht.
- Die Ausgabe:
  - `Start:        [42 7 19 88 3 25 71 14 60 9]`
  - `Ganz:         [9 7 19 3 25 14 42 88 60 71] Pivot am Index 6`
  - `Linker Teil:  [3 7 9 19 25 14 42 88 60 71] Pivot am Index 2`
  - `Rechter Teil: [3 7 9 19 25 14 42 71 60 88] Pivot am Index 9`
- Unter „Was dieser Stufe fehlt“: Nach drei Aufrufen stehen 9, 42 und 88 fest. Die Bereiche 0 bis 1, 3 bis 5 und 7 bis 8 brauchen weitere Aufrufe. Wie viele, hängt von den Werten ab.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 5: Rekursion

Reiter: rekursion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen: was `quickSort(pa, 3, 3)` zu tun hat (nichts), und wie ein Aufruf mit `start` größer als `end` entsteht (das Pivot landet ganz links, der linke Teil ist leer).
- Eine Liste erklärt die drei Teile von `quickSort`: das Ende in den Zeilen 2 bis 4 bei `start >= end`, das Aufteilen in Zeile 5, die zwei rekursiven Aufrufe in den Zeilen 6 und 7.
- Die Demo zeigt 15 Zeilen: `quickSort` in den Zeilen 1 bis 8 und `main` in den Zeilen 10 bis 15. `partition` und `swap` zeigt sie nicht.
- Jeder laufende Aufruf hat einen Rahmen mit `pa`, `start`, `end` und `p`, zum Beispiel „quickSort, Bereich 3 bis 5“. `pa` steht in Worten da: „Adresse von a“. Im Array ist gelb der Bereich des jüngsten Aufrufs, grün sind die Elemente, deren Platz feststeht.
- Ein Schritt der Demo ist ein Aufruf, ein Aufteilen oder eine Rückkehr. Die Demo hat 35 Schritte.
- Die 13 Aufrufe in ihrer Reihenfolge: 0 bis 9 (Pivot 42 an Index 6), 0 bis 5 (Pivot 9 an Index 2), 0 bis 1 (Pivot 3 an Index 0), 0 bis −1 (leer), 1 bis 1, 3 bis 5 (Pivot 19 an Index 4), 3 bis 3, 5 bis 5, 7 bis 9 (Pivot 88 an Index 9), 7 bis 8 (Pivot 71 an Index 8), 7 bis 7, 9 bis 8 (leer), 10 bis 9 (leer). Sechs Aufrufe teilen auf, sieben kehren sofort zurück. Höchstens vier Aufrufe laufen gleichzeitig.
- Die Ausgabe: `Start:    [42 7 19 88 3 25 71 14 60 9]` und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Unter „Einordnung nach Abschnitt 7.5 des Kurstexts“ stehen drei Punkte: kaskadenartige Rekursion wie bei den Fibonacci-Zahlen, keine Endrekursion, Terminierung, weil das Pivot jedes Mal aus dem Bereich fällt.
- Unter „Ausblick: Aufteilen nach Hoare“ steht ein zweites `partition` mit zwei Indizes `left` und `right`, die von beiden Seiten aufeinander zulaufen und Paare tauschen, die beide auf der falschen Seite stehen. Nach dem ersten Aufteilen steht dort `[14 7 19 9 3 25 42 71 60 88]`, das Pivot ebenfalls am Index 6. `quickSort` bleibt gleich. Die Übung teilt mit einem Index auf, weil das leichter zu durchschauen ist.
- Unter „Wie schnell ist Quicksort?“ steht ein Link auf den Exkurs zur Laufzeit. Die Lektion selbst zählt keine Vergleiche. Fragt der Student danach: Im Mittel braucht Quicksort etwa n log n Vergleiche. Mit dem ersten Element als Pivot ist ein schon sortiertes Array der schlechteste Fall: Jedes Aufteilen trennt nur ein Element ab, und Quicksort braucht dann so viele Vergleiche wie Bubble Sort, etwa n²/2. Der Kurstext sagt zur Laufzeit von Quicksort nichts, auch nicht in Abschnitt 7.5.5. Eine Antwort dazu nennt keinen Abschnitt des Kurstexts als Quelle, sondern den Exkurs zur Laufzeit.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die Auflösungen aufklappen.
- Den Links folgen.

### Abschnitte des Kurstexts

- 7.5.3
- 7.5.5
- 7.5.2

## Schritt 6: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Im Editorlotsen steht ein Editor mit dem Knopf „Ausführen“. Auf GitHub Pages ist der Reiter grau. Dort wählt der Student ein Programm und sieht das Programm und seine Ausgabe, festgehalten mit Go 1.26.1.
- Sieben vollständige Programme: „Stufe 1: fünf Variablen“, „Stufe 2: Array und Schleife“, „Stufe 3: b, p := partition(a)“, „Stufe 3: Ergebnis nicht aufgefangen“, „Stufe 4: dreimal von Hand“, „Stufe 5: Quicksort“, „Stufe 5 mit Ausgabe jedes Aufrufs“.
- Das letzte Programm gibt bei jedem Aufruf `quickSort(pa, start, end)` aus und nach jedem Aufteilen das Array mit Pivot und Index.
- Vier Vorschläge: andere Werte in Stufe 2, ein Array mit gleichen Werten in Stufe 5, die Aufrufe zählen, bei denen das Ende greift (sieben), die drei Zeilen mit `if start >= end` löschen.
- Zum letzten Vorschlag: Ohne das Ende ruft der Aufruf für den Bereich 0 bis −1 sich immer wieder mit dem gleichen Bereich auf. Go bricht ab mit `fatal error: stack overflow`.
- Am Ende des Schritts unter „Nächste Schritte“: die Programmierübungen „Bubble Sort mit Zeigern“ und „Bubble Sort rekursiv“, der Exkurs zur Laufzeit und das Quiz zu Funktionen, Zeigern und Rekursion.

### Handgriffe

- Ein Programm wählen, im Editor ändern und ausführen, wenn der Editorlotse die Lektion ausliefert.

### Abschnitte des Kurstexts

- 7.5

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle nennt zu jedem Schritt den Abschnitt im Kurstext, was daraus stammt und was nicht im Kurstext steht: 5.3 Tupel-Zuweisung, 6.1 und 6.2 Array und Durchlaufen, 7.1 und 7.2 Funktion und Wertübergabe, 7.3 und 7.4 Zeiger auf ein Array, 7.5 Rekursion.
- Unter „Wo das Thema im Kurstext noch vorkommt“: Abschnitt 2.6 Unterprogramme, die Türme von Hanoi in Abschnitt 7.5.4, die binäre Suche in Abschnitt 8.4.1, Bubblesort in Abschnitt 8.4.2, je mit Lektion.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
