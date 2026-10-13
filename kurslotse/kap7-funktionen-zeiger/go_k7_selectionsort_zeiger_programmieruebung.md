# Programmierübung: Selection Sort mit Zeigern

Die Seite ist eine Programmierübung. Sie stellt die Funktionen `findMin` und `swap` aus der Übung „Selection Sort mit Funktionen“ auf Zeiger um. Die Funktionen bekommen die Adresse des Arrays, das Programm kopiert kein Element mehr. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.3 und 7.4, Seiten 171–180.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- Das Verfahren ändert sich nicht: neun Runden, 45 Vergleiche, 7 Tausche.
- `swap(pa *[10]int, i int, j int)` tauscht in dem Array, auf das `pa` zeigt, die Elemente an den Indizes `i` und `j`. Die Funktion hat keine Rückgabe. Der Aufruf lautet `swap(&a, pos, minIndex)`.
- `findMin(pa *[10]int, pos int) int` liefert den Index des kleinsten Elements ab dem Index `pos`. Der Aufruf lautet `findMin(&a, pos)`.
- `pa *[10]int` ist ein Zeiger auf ein Array vom Typ `[10]int`. `&a` liefert die Adresse von `a`.
- `pa[i]` ist die Kurzschreibweise für `(*pa)[i]`. Das steht in Hinweis 7.4 des Kurstexts.
- In `findMin` steht `len(*pa)`: Der Stern führt vom Zeiger zum Array, `len` liefert dessen Länge 10. Go erlaubt auch `len(pa)` mit dem Zeiger selbst. So steht es in Listing 8.2 des Kurstexts, dort `len(a)` mit `a *[10]person`. Die Seite verwendet die Schreibweise mit dem Stern.
- Auch ein Zeiger wird als Kopie übergeben. Kopiert wird eine Adresse und kein Element.
- Unter den Rahmen der Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger, in ihm steht die Zahl 100, die Adresse des Arrays. Die Adressen sind erfunden. Ein `int` und ein Zeiger belegen im Modell je 8 Zellen. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke in Schritt 3: `a` in den Zellen 100 bis 179, `pos` und `minIndex` von `main` ab 180 und 188. `pa` von `findMin` oder `swap` in den Zellen 196 bis 203 mit dem Inhalt 100, die übrigen Parameter und Variablen der Funktion ab 204 und 212.
- Selection Sort kommt im Kurstext nicht vor. Das Muster für die Zeiger ist `negToZeroPointer` in Listing 7.2.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und den drei Zeilen einer Runde aus der Fassung mit Kopien. Eine Demo gibt es hier nicht.
- Der Text sagt: Die Fassung mit Kopien kopiert für zehn Elemente 230 Elemente, und `swap` muss ihre Kopie zurückgeben.
- Unter „Was dazugehört“ stehen Links auf die Lektion „Zeiger in Go“, auf die Programmierübung „Bubble Sort mit Zeigern“ und auf den Exkurs „Das Kopier-Problem“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.2

## Schritt 1: swap mit einem Zeiger

Reiter: tausch

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `swap` mit dem Parameter `pa *[10]int`. Eine Liste nennt Parameter, Rumpf und die entfallene Rückgabe.
- Aus `a = swap(a, pos, minIndex)` wird `swap(&a, pos, minIndex)`.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4.1

## Schritt 2: findMin mit einem Zeiger

Reiter: suche

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `findMin` mit dem Parameter `pa *[10]int`, mit `len(*pa)` in der Schleife und mit `pa[i] < pa[minIndex]`.
- Eine Liste nennt die drei Stellen, die anders sind als in der Fassung mit der Kopie: Parameter, Länge, Elemente.
- Die Rückgabe bleibt ein `int`. Ein Zeiger als Parameter und eine Rückgabe schließen sich nicht aus, wie bei `incr` in Abschnitt 7.4.1.
- Die Funktion ändert das Array nicht. Der Zeiger spart die Kopie von zehn Elementen je Runde, zusammen 90.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 3: Runde für Runde

Reiter: runden

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt 24 Zeilen: `findMin` in den Zeilen 1 bis 9, `swap` in den Zeilen 11 bis 13 und `main` in den Zeilen 15 bis 24.
- Darunter stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Im Rahmen von `findMin` und von `swap` steht kein Array. Dort steht der Zeiger `pa` mit dem Wert „Adresse von a“. Es gibt nur ein Array, im Rahmen von `main`.
- Grün markiert ist im Array von `main` das gefundene kleinste Element oder das getauschte Paar.
- Ein Schritt der Demo ist ein Aufruf von `findMin`, die Prüfung `minIndex != pos` ohne Tausch oder ein Aufruf von `swap`. Die Demo hat 20 Schritte und zeigt alle neun Runden.
- Die neun Runden: Runde 1 tauscht 42 und 3. Runde 2 tauscht nicht. Runde 3 tauscht 19 und 9. Runde 4 tauscht 88 und 14. Runde 5 tauscht 42 und 19. Runde 6 tauscht nicht. Runde 7 tauscht 71 und 42. Runde 8 tauscht 88 und 60. Runde 9 tauscht 88 und 71.
- Am Ende: 9 Aufrufe von `findMin`, 7 Aufrufe von `swap`, kein kopiertes Element. Die Ausgabe ist `[3 7 9 14 19 25 42 60 71 88]`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Array lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.3
- 7.4.1

## Schritt 4: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das ganze Programm mit den zwei Funktionen und `main` mit der äußeren Schleife.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe mit neun Zeilen „Nach Runde 1“ bis „Nach Runde 9“. Sie ist die gleiche wie in der Fassung mit Kopien.
- Eine Tabelle „Vier Programme im Vergleich“ nennt Vergleiche, Tausche und kopierte Elemente: Bubble Sort mit Wertübergabe 45, 23 und 910. Bubble Sort mit Zeigern 45, 23 und 0. Selection Sort mit Wertübergabe 45, 7 und 230. Selection Sort mit Zeigern 45, 7 und 0.
- Darunter steht: Die Zahl der Vergleiche und Tausche gehört zum Verfahren. Die Zahl der kopierten Elemente gehört zur Art der Übergabe.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: eine Funktion `selectionSort(pa *[10]int)` schreiben, `findMin` mit einer Rückgabe vom Typ `*int` schreiben, die Meldung zu einem fehlenden `&` nachsehen.
- Zur ersten Aufgabe: In `selectionSort` ist `pa` schon die Adresse des Arrays. Die Aufrufe lauten `findMin(pa, pos)` und `swap(pa, pos, minIndex)`, ohne `&`.
- Zur dritten Aufgabe: Go meldet `cannot use a (variable of type [10]int) as *[10]int value in argument to swap`.
- Eine Demo gibt es hier nicht.

### Handgriffe

- „Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 7.3
- 7.4.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit fünf Zeilen: Abschnitt 5.3.2, 7.2, 7.3, 7.4 mit Hinweis 7.4 und 7.4.1, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Selection Sort kommt im Kurstext nicht vor.
- Unter „Die Reihe“ stehen Links auf die Programmierübungen davor und auf die zwei Übungen mit Bubble Sort.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4
