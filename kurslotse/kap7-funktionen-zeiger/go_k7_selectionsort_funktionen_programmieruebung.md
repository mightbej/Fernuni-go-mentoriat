# Programmierübung: Selection Sort mit Funktionen

Die Seite ist eine Programmierübung. Sie zerlegt das Sortierprogramm mit Selection Sort aus Kapitel 6 in Funktionen: Die Suche nach dem kleinsten Element wird zur Funktion `findMin`, der Tausch zur Funktion `swap`. Beide bekommen das Array als Kopie. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.1 und 7.2, Seiten 163–170.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- Selection Sort arbeitet in Runden. In der Runde mit dem Index `pos` sucht es das kleinste Element ab dem Index `pos` und tauscht es an den Index `pos`. Für zehn Elemente braucht es neun Runden.
- Das Verfahren ändert sich in dieser Übung nicht. Das Programm vergleicht 45-mal und tauscht siebenmal, wie das Programm aus Kapitel 6.
- `findMin(a [10]int, pos int) int` liefert den Index des kleinsten Elements ab dem Index `pos`. Sie liefert den Index und nicht den Wert.
- `swap(a [10]int, i int, j int) [10]int` tauscht in ihrer Kopie die Elemente an den Indizes `i` und `j` und gibt die Kopie zurück. Der Aufrufer schreibt `a = swap(a, pos, minIndex)`.
- Die Variable heißt `minIndex` und nicht `min`, weil es in Go seit Version 1.21 eine eingebaute Funktion `min` gibt und weil die Variable einen Index enthält.
- In Go bekommt eine Funktion von jedem Argument eine Kopie, das heißt call-by-value. Zeiger kommen auf dieser Seite nicht vor, sie stehen in der Programmierübung „Selection Sort mit Zeigern“.
- Selection Sort kommt im Kurstext nicht vor. Die Übung ist eine Ergänzung des Mentors.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Die Demos stellen im Browser nach, was das Programm tut. Ausführen und ändern lässt sich das Programm nur in Schritt 5, „Selbst ausprobieren“, wenn der Editorlotse die Lektion ausliefert.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und dem Kern des Programms aus Kapitel 6: äußere Schleife über `pos`, innere Schleife für die Suche, Tausch mit der Bedingung `minIndex != pos`. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Teile einer Runde, die Funktionen werden: Suche und Tausch.
- Unter „Was dazugehört“ stehen Links auf die Programmierübung „Sortieren mit Selection Sort“ in Kapitel 6, auf die Lektion „Funktionen in Go“ und auf die Programmierübung „Bubble Sort mit Funktionen“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 7.2

## Schritt 1: Die Suche als Funktion

Reiter: minimum

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `findMin` mit der lokalen Variable `minIndex`, einer Schleife ab `pos + 1` und `return minIndex`.
- Eine Liste nennt Parameter, lokale Variable und Rückgabe. `minIndex` gehört der Funktion und ist in `main` nicht sichtbar.
- Für das Array der Seite liefert `findMin(a, 0)` den Wert 4, dort steht die 3.
- Die Funktion liest das Array und ändert es nicht.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 3.4

## Schritt 2: Der Tausch als Funktion

Reiter: tausch

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `swap` mit zwei Indizes `i` und `j`. Bubble Sort tauscht Nachbarn und kommt mit einem Index aus, Selection Sort tauscht zwei beliebige Elemente.
- Ein zweiter Codeblock zeigt eine Runde: `minIndex := findMin(a, pos)`, dann `if minIndex != pos { a = swap(a, pos, minIndex) }`.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.2

## Schritt 3: Runde für Runde

Reiter: runden

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt 25 Zeilen: `findMin` in den Zeilen 1 bis 9, `swap` in den Zeilen 11 bis 14 und `main` in den Zeilen 16 bis 25.
- Darunter stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Im Rahmen von `findMin` und von `swap` steht eine eigene Kopie des Arrays. Grün markiert ist das gefundene kleinste Element oder das getauschte Paar.
- Ein Schritt der Demo ist ein Aufruf von `findMin`, die Prüfung `minIndex != pos` ohne Tausch, ein Aufruf von `swap` oder die Zuweisung der Rückgabe von `swap`. Die Demo hat 27 Schritte und zeigt alle neun Runden.
- Der Satz unter der Demo zählt die kopierten Elemente: 10 für jeden Aufruf von `findMin`, 10 für jeden Aufruf von `swap` und 10 für jede Zuweisung der Rückgabe.
- Die neun Runden: Runde 1 tauscht 42 und 3. Runde 2 tauscht nicht, die 7 steht richtig. Runde 3 tauscht 19 und 9. Runde 4 tauscht 88 und 14. Runde 5 tauscht 42 und 19. Runde 6 tauscht nicht, die 25 steht richtig. Runde 7 tauscht 71 und 42. Runde 8 tauscht 88 und 60. Runde 9 tauscht 88 und 71.
- Am Ende: 9 Aufrufe von `findMin`, 7 Aufrufe von `swap`, 230 kopierte Elemente. Die Ausgabe ist `[3 7 9 14 19 25 42 60 71 88]`.
- Die Demo zeigt die Suche in `findMin` nicht Vergleich für Vergleich. Das zeigt Schritt 2 der Programmierübung zu Kapitel 6.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Array lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.2
- 7.1

## Schritt 4: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das ganze Programm mit `package main`, `import "fmt"`, den zwei Funktionen und `main` mit der äußeren Schleife.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe: `Start:`, neun Zeilen „Nach Runde 1“ bis „Nach Runde 9“ und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Unter „Die Bilanz“ steht: 45 Vergleiche, 7 Tausche, 9 · 10 + 7 · 20 = 230 kopierte Elemente. Bubble Sort mit `isGreater` und `swap` kopiert 910 Elemente, weil es seine Funktionen viel öfter aufruft.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: eine Funktion `findMax` schreiben, die Zeile `if minIndex != pos` weglassen, `findMin` mit zwei Rückgaben schreiben.
- Zur zweiten Aufgabe: Ohne die Bedingung sortiert das Programm weiter richtig. Es ruft `swap` dann neunmal auf, auch wenn ein Element mit sich selbst getauscht wird, und kopiert 9 · 10 + 9 · 20 = 270 Elemente.
- Eine Demo gibt es hier nicht.

### Handgriffe

- „Programm kopieren“ wählen.
- Dem Link zur Programmierübung „Selection Sort mit Zeigern“ folgen.

### Abschnitte des Kurstexts

- 7.2
- 7.1.1

## Schritt 5: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Im Editorlotsen steht ein Editor mit dem Knopf „Ausführen“. Auf GitHub Pages ist der Reiter grau. Dort sieht der Student das Programm und seine Ausgabe, festgehalten mit Go 1.26.1, und einen Satz, wie er die Lektion im Editorlotsen öffnet.
- Ein Programm: „Selection Sort mit Funktionen“. Es ist das ganze Programm aus Schritt 4. Die Ausgabe ist die gleiche wie dort.
- Der Text schlägt vor, im Editor die Aufgaben unter „Zum Weiterdenken“ aus Schritt 4 zu lösen. Fragt der Student nach einer Lösung, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Der Stand der Demo sagt, welches Programm gewählt ist, ob der Student es geändert hat und was der letzte Lauf ausgegeben oder gemeldet hat. Die Programmzeilen sind das Programm im Editor.

### Handgriffe

- Das Programm ansehen.
- Im Editor schreiben, „Ausführen“, „Stopp“ oder „Programm zurücksetzen“ wählen, wenn der Editorlotse die Lektion ausliefert.

### Abschnitte des Kurstexts

- 7.1

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit sechs Zeilen: Abschnitt 3.4, 5.3.2, 5.5.3.2, 6.1, 7.1 und 7.2, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Selection Sort kommt im Kurstext nicht vor.
- Unter „Die Reihe“ stehen Links auf die Programmierübungen davor und danach.
- Der Reiter sagt außerdem: Die Funktion swap folgt dem Muster von negToZeroValue in Listing 7.1. Die Variable heißt minIndex, weil es in Go seit Version 1.21 eine eingebaute Funktion min gibt.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 7.2
