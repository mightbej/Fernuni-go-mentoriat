# Das Kopier-Problem: Arrays als Argument

Die Seite ist ein Exkurs. Sie rechnet an einer Funktion `swap` aus, wie viele Elemente ein Programm kopiert, wenn es ein Array als Wert an eine Funktion übergibt, und zeigt das Programm mit einem Zeiger. Sie gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.2 und 7.3, Seiten 169–172.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- In Go bekommt eine Funktion von jedem Argument eine Kopie des Werts, das heißt call-by-value. Bei einem Array ist das eine Kopie aller Elemente.
- Im Kurstext steht dazu in Abschnitt 7.2: Bei jedem Funktionsaufruf mit einem Array als Argument muss eine Kopie des gesamten Arrays erstellt werden, was Rechenzeit und Speicherplatz beansprucht. Das Beispiel dort ist ein Array vom Typ `[1000][1000]int` mit 1.000.000 Elementen.
- Das Array der Seite ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- Die Funktion `swap` steht nicht im Kurstext. Sie folgt dem Muster von `negToZeroValue` in Listing 7.1 und von `negToZeroPointer` in Listing 7.2.
- Die Seite zählt kopierte Elemente. Sie misst keine Zeit. Wie lange das Kopieren auf einem bestimmten Rechner dauert, sagt die Seite nicht. Ein Compiler darf eine Kopie einsparen, wenn das Ergebnis das gleiche bleibt.
- Jede Demo zeigt ein Programm, drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, darunter Rahmen mit Variablen, die Ausgabe und einen Kasten mit einem Satz zur letzten Anweisung. Jede laufende Funktion hat einen Rahmen. Ein Array steht als Reihe kleiner Felder da, grün ist ein Element, das die letzte Anweisung geändert hat.
- Unter den Rahmen einer Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger. In ihm steht eine Zahl, die Adresse einer Zelle, und darunter „Adresse der Zelle …“. Ein Block mit gelbem Schein ist neu oder gerade geändert. Die Adressen sind erfunden. Ein `int` belegt im Modell 8 Zellen, ein Zeiger auch. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke in Schritt 1: Das Array `a` von `main` liegt in den Zellen 100 bis 179. Der Parameter `a` von `swap` ist ein zweites Array in den Zellen 180 bis 259, `i` und `j` liegen ab 260 und 268. Die Kopie belegt also 80 Zellen.
- Die Blöcke in Schritt 3: Das Array `a` liegt in den Zellen 100 bis 179. Der Parameter `pa` liegt in den Zellen 180 bis 187 und enthält die Zahl 100. Der Zeiger belegt 8 Zellen, gleich wie lang das Array ist.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt die Aussage aus Abschnitt 7.2 und das Beispiel der Seite, die Funktion `swap`.
- Unter „Was dazugehört“ stehen drei Punkte mit Links: die Lektion „Funktionen in Go“, die Lektion „Zeiger in Go“ und die zwei Programmierübungen zum Sortieren aus Kapitel 6.
- Ein gelber Kasten sagt: Die Seite zählt Elemente und misst keine Zeit.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.2
- 7.3

## Schritt 1: Ein Tausch, zwei Kopien

Reiter: tausch

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit zwei Fragen und aufklappbaren „Auflösungen“.
- Die Demo zeigt zehn Zeilen: die Funktion `swap(a [10]int, i int, j int) [10]int` mit `a[i], a[j] = a[j], a[i]` und `return a`, darunter `main` mit dem Array, `a = swap(a, 0, 1)` und `fmt.Println(a)`.
- Der Ablauf hat sechs Schritte: Array anlegen. Aufruf: Der Parameter `a` bekommt eine Kopie aller zehn Elemente, kopiert sind 10. Tausch in der Kopie. `return a`. Zuweisung in `main`: zehn Elemente zurück, kopiert sind 20. Ausgabe `[7 42 19 88 3 25 71 14 60 9]`.
- Der Satz unter der Demo zählt die kopierten Elemente mit.
- Steht in `main` nur `swap(a, 0, 1)` ohne `a =`, bleibt das Array unverändert. Go übersetzt das Programm und meldet nichts.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Eine „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 7.2

## Schritt 2: Wie viel kopiert wird

Reiter: zaehlen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text rechnet: Ein Aufruf von `swap` kopiert bei einem Array der Länge n zweimal n Elemente.
- Selection Sort mit `findMin` und `swap`: n − 1 Runden, in jeder Runde kopiert `findMin` n Elemente und `swap` höchstens 2 · n. Die Obergrenze ist 3 · n · (n − 1).
- Bubble Sort mit `isGreater` und `swap`: n · (n − 1) / 2 Vergleiche, jeder kopiert n Elemente, dazu höchstens ebenso viele Tausche mit 2 · n. Die Obergrenze ist 1,5 · n · n · (n − 1).
- Die Demo hat vier Knöpfe für die Länge: 10, 100, 1.000 und 1.000.000 Elemente. Eine Tabelle zeigt die kopierten Elemente für einen Aufruf von `swap`, für Selection Sort und für Bubble Sort, mit Wertübergabe und mit Zeigern. Mit Zeigern steht überall 0.
- Die Zahlen mit Wertübergabe: Bei 10 Elementen 20, höchstens 270 und höchstens 1.350. Bei 100 Elementen 200, höchstens 29.700 und höchstens 1.485.000. Bei 1.000 Elementen 2.000, höchstens 2.997.000 und höchstens 1.498.500.000. Bei 1.000.000 Elementen 2.000.000, höchstens 2.999.997.000.000 und höchstens 1.499.998.500.000.000.000.
- Für das Array der Programmierübungen mit zehn Elementen sind es weniger als die Obergrenze: Bubble Sort mit 45 Vergleichen und 23 Tauschen kopiert 910 Elemente, Selection Sort mit 9 Runden und 7 Tauschen kopiert 230.
- Ein `int` belegt auf einem 64-Bit-Rechner 8 Byte. Bei 1.000.000 Elementen kopiert ein Aufruf von `swap` 16 Megabyte.
- Die Rechnung steht nicht im Kurstext.

### Handgriffe

- Eine der vier Längen wählen.
- Die „Auflösung“ aufklappen.
- Andere Längen lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.2

## Schritt 3: Mit einem Zeiger

Reiter: zeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste nennt die drei Änderungen: Der Parameter ist `pa *[10]int`. Die Rückgabe entfällt. Der Aufruf lautet `swap(&a, 0, 1)`.
- `pa[i]` ist die Kurzschreibweise für `(*pa)[i]` aus Hinweis 7.4 des Kurstexts.
- Die Demo zeigt neun Zeilen: die Funktion `swap(pa *[10]int, i int, j int)` mit `pa[i], pa[j] = pa[j], pa[i]`, darunter `main` mit dem Array, `swap(&a, 0, 1)` und `fmt.Println(a)`.
- Der Ablauf hat fünf Schritte: Array anlegen. Aufruf: `pa` bekommt die Adresse von `a`, kopiert ist kein Element. Der Tausch steht sofort im Array von `main`. Die Funktion endet ohne Rückgabe. Ausgabe `[7 42 19 88 3 25 71 14 60 9]`.
- Der Wert des Zeigers steht in Worten da: „Adresse von a“.
- Eine Tabelle vergleicht `swap` mit einer Kopie und `swap` mit einem Zeiger in fünf Zeilen. Die letzte Zeile: Fehlt bei der Kopie `a =`, bleibt das Array unverändert, und Go meldet nichts. Fehlt beim Zeiger `&`, übersetzt Go das Programm nicht. Die Meldung lautet `cannot use a (variable of type [10]int) as *[10]int value in argument to swap`.
- Ein gelber Kasten sagt: Auch der Zeiger wird kopiert. Die Kopie einer Adresse ist immer gleich groß.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Ein Tausch, zwei Kopien, 7.2, Seiten 169–170: Listing 7.1: Das Muster von negToZeroValue in Listing 7.1 stammt aus dem Kurstext. Die Funktion swap ist ein eigenes Beispiel nach diesem Muster.
  - 2 Wie viel kopiert wird, 7.2, Seiten 169–170: Das Beispiel [1000][1000]int mit 1.000.000 Elementen je Aufruf stammt aus dem Kurstext. Die Rechnung steht nicht im Kurstext. Die Zahlen 45, 23, 9 und 7 stammen aus den Programmierübungen zu Kapitel 6.
  - 3 Mit einem Zeiger, 7.3, Seiten 171–172: Listing 7.2; Hinweis 7.4, Seite 177: Das Muster von negToZeroPointer in Listing 7.2 und Hinweis 7.4 stammen aus dem Kurstext. Die Funktion swap ist ein eigenes Beispiel. Die Meldung zu einem fehlenden & stammt von Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.6, Seiten 15–17, Vertauschen: Zwei Werte mit einer dritten Variable vertauschen, ohne Go. Lektion: Unterprogramme und Funktionen, Schritt 1.
  - Abschnitt 5.3.2, Seiten 128–129, Tupel-Zuweisung: a, b = b, a. Lektion: Zuweisungen in Go, Schritt 4.
  - Kapitel 6, Seiten 151–162, Arrays: Die zwei Sortierverfahren mit ihren Zahlen der Vergleiche und Tausche. Lektion: Sortieren mit Bubble Sort.
- Unter „Wie es weitergeht“ stehen Links auf die vier Programmierübungen: Bubble Sort und Selection Sort, je mit Funktionen und mit Zeigern.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.6
- 5.3.2
- 7.2
- 7.3
