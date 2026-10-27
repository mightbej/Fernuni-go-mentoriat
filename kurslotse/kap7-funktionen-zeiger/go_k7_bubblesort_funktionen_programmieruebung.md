# Programmierübung: Bubble Sort mit Funktionen

Die Seite ist eine Programmierübung. Sie zerlegt das Sortierprogramm mit Bubble Sort aus Kapitel 6 in Funktionen: Der Vergleich wird zur Funktion `isGreater`, der Tausch zur Funktion `swap`. Beide bekommen das Array als Kopie. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.1 und 7.2, Seiten 163–170.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- Bubble Sort vergleicht in einem Durchgang jedes Paar von Nachbarn und tauscht, wenn das linke Element größer ist. Nach einem Durchgang steht das größte Element am Ende. Für zehn Elemente braucht es neun Durchgänge.
- Das Verfahren ändert sich in dieser Übung nicht. Das Programm vergleicht 45-mal und tauscht 23-mal, wie das Programm aus Kapitel 6.
- `isGreater(a [10]int, i int) bool` liefert `a[i] > a[i+1]`.
- `swap(a [10]int, i int) [10]int` tauscht in ihrer Kopie die Nachbarn an den Indizes `i` und `i+1` und gibt die Kopie zurück. Der Aufrufer schreibt `a = swap(a, i)`.
- In Go bekommt eine Funktion von jedem Argument eine Kopie, das heißt call-by-value. Zeiger kommen auf dieser Seite nicht vor, sie stehen in der Programmierübung „Bubble Sort mit Zeigern“.
- Die Funktionen `isGreater` und `swap` stehen nicht im Kurstext. Bubblesort steht im Kurstext in Abschnitt 8.4.2, dort mit einem Array von Verbunden.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Die Demos stellen im Browser nach, was das Programm tut. Ausführen und ändern lässt sich das Programm nur in Schritt 5, „Selbst ausprobieren“, wenn der Editorlotse die Lektion ausliefert.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und dem Kern des Programms aus Kapitel 6: zwei Schleifen, ein Vergleich, ein Tausch. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Stücke, die in Funktionen wandern.
- Unter „Was dazugehört“ stehen Links auf die Programmierübung „Sortieren mit Bubble Sort“ in Kapitel 6 und auf die Lektion „Funktionen in Go“.
- Ein gelber Kasten sagt: Die Funktionen machen das Programm lesbarer und nicht schneller.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 7.2
- 8.4.2

## Schritt 1: Der Vergleich als Funktion

Reiter: vergleich

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `isGreater`. Eine Liste nennt Parameter, Rückgabe und Signatur: Parameter-Typen `[10]int` und `int`, Rückgabe-Typ `bool`.
- Ein Funktionsaufruf mit einer Rückgabe ist ein Ausdruck und darf als Bedingung stehen: `if isGreater(a, i)`.
- Die Länge 10 gehört zum Typ des Parameters. Die Funktion nimmt nur Arrays vom Typ `[10]int` an.
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
- Ein Codeblock zeigt `swap` mit der Rückgabe `[10]int`. Ein zweiter zeigt den Aufruf `a = swap(a, i)` in der if-Anweisung.
- Fehlt `a =` vor dem Aufruf, übersetzt Go das Programm, und es gibt am Ende das unsortierte Array aus. Jeder Aufruf tauscht in einer Kopie, und die Rückgabe ist verworfen.
- Ein gelber Kasten verweist auf `negToZeroValue` in Listing 7.1.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.2

## Schritt 3: Ein Durchgang

Reiter: durchgang

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt 18 Zeilen: `isGreater` in den Zeilen 1 bis 3, `swap` in den Zeilen 5 bis 8 und `main` mit einer Schleife `for i := 0; i < len(a)-1; i++` in den Zeilen 10 bis 18. Das ist ein einziger Durchgang.
- Darunter stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Jede laufende Funktion hat einen Rahmen. Im Rahmen von `isGreater` und von `swap` steht eine eigene Kopie des Arrays. Grün markiert sind die zwei Elemente des Paars.
- Ein Schritt der Demo ist ein Aufruf von `isGreater`, ein Aufruf von `swap` oder die Zuweisung der Rückgabe von `swap`. Die Demo hat 27 Schritte.
- Der Satz unter der Demo zählt die kopierten Elemente: 10 für jeden Aufruf von `isGreater`, 10 für jeden Aufruf von `swap` und 10 für jede Zuweisung der Rückgabe.
- Der erste Durchgang: 9 Vergleiche, 8 Tausche, 250 kopierte Elemente. Nicht getauscht wird nur das Paar 42 und 88. Die Ausgabe am Ende ist `[7 19 42 3 25 71 14 60 9 88]`.

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

- Zu sehen ist das ganze Programm mit `package main`, `import "fmt"`, den zwei Funktionen und `main` mit zwei Schleifen. Die Laufvariable der äußeren Schleife heißt `durchgang`.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe: `Start:`, neun Zeilen „Nach Durchgang 1“ bis „Nach Durchgang 9“ und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Nach Durchgang 7 ist das Array sortiert. Die Durchgänge 8 und 9 tauschen nichts mehr.
- Unter „Die Bilanz“ steht: 45 Vergleiche, 23 Tausche, 45 · 10 + 23 · 20 = 910 kopierte Elemente.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: `a =` streichen, eine Funktion `isSorted(a [10]int) bool` schreiben, einen ganzen Durchgang in eine Funktion auslagern.
- Eine Demo gibt es hier nicht.

### Handgriffe

- „Programm kopieren“ wählen.
- Den Links zum Exkurs „Das Kopier-Problem“ und zur Programmierübung „Bubble Sort mit Zeigern“ folgen.

### Abschnitte des Kurstexts

- 7.2
- 8.4.2

## Schritt 5: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Im Editorlotsen steht ein Editor mit dem Knopf „Ausführen“. Auf GitHub Pages ist der Reiter grau. Dort sieht der Student das Programm und seine Ausgabe, festgehalten mit Go 1.26.1, und einen Satz, wie er die Lektion im Editorlotsen öffnet.
- Ein Programm: „Bubble Sort mit Funktionen“. Es ist das ganze Programm aus Schritt 4. Die Ausgabe ist die gleiche wie dort.
- Der Text schlägt vor, im Editor die Aufgaben unter „Zum Weiterdenken“ aus Schritt 4 zu lösen. Fragt der Student nach einer Lösung, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Der Stand der Demo sagt, welches Programm gewählt ist, ob der Student es geändert hat und was der letzte Lauf ausgegeben oder gemeldet hat. Die Programmzeilen sind das Programm im Editor.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Davor: „Sortieren mit Bubble Sort“ in Kapitel 6.
  - Danach: „Bubble Sort mit Zeigern“ und „Bubble Sort rekursiv“.
  - Das andere Verfahren: „Selection Sort mit Funktionen“.
  - Quiz: „Funktionen, Zeiger und Rekursion“.

### Handgriffe

- Das Programm ansehen.
- Im Editor schreiben, „Ausführen“, „Stopp“ oder „Programm zurücksetzen“ wählen, wenn der Editorlotse die Lektion ausliefert.

### Abschnitte des Kurstexts

- 7.1

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit sieben Zeilen: Abschnitt 3.4, 4.5, 5.3.2, 6.1, 7.1, 7.2 und 8.4.2, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Die Zerlegung in `isGreater` und `swap` kommt im Kurstext nicht vor.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.1
- 7.2
- 8.4.2
