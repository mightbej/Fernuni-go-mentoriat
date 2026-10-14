# Programmierübung: Bubble Sort mit Zeigern

Die Seite ist eine Programmierübung. Sie stellt die Funktionen `isGreater` und `swap` aus der Übung „Bubble Sort mit Funktionen“ auf Zeiger um. Die Funktionen bekommen die Adresse des Arrays, das Programm kopiert kein Element mehr. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.3 und 7.4, Seiten 171–180.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- Das Verfahren ändert sich nicht: 45 Vergleiche, 23 Tausche.
- `swap(pa *[10]int, i int)` tauscht in dem Array, auf das `pa` zeigt, die Nachbarn an den Indizes `i` und `i+1`. Die Funktion hat keine Rückgabe. Der Aufruf lautet `swap(&a, i)`.
- `isGreater(pa *[10]int, i int) bool` liefert `pa[i] > pa[i+1]`. Der Aufruf lautet `isGreater(&a, i)`.
- `pa *[10]int` ist ein Zeiger auf ein Array vom Typ `[10]int`. `&a` liefert die Adresse von `a`, `&` heißt Adressoperator.
- `pa[i]` ist die Kurzschreibweise für `(*pa)[i]`. Das steht in Hinweis 7.4 des Kurstexts. Ohne Klammern wäre `*pa[i]` falsch: Go meldet `invalid operation: cannot indirect pa[i] (variable of type int)`.
- Eine Funktion, die eine Variable außerhalb ihres Rumpfs ändert, hat einen Seiteneffekt. Das steht in Abschnitt 7.4.1.
- Auch ein Zeiger wird als Kopie übergeben. Kopiert wird eine Adresse und kein Element.
- Unter den Rahmen der Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger, in ihm steht die Zahl 100, die Adresse des Arrays. Die Adressen sind erfunden. Ein `int` und ein Zeiger belegen im Modell je 8 Zellen. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke in Schritt 3: `a` in den Zellen 100 bis 179, `i` von `main` ab 180. `pa` von `isGreater` oder `swap` in den Zellen 188 bis 195 mit dem Inhalt 100, der Parameter `i` ab 196.
- Die Funktionen `isGreater` und `swap` stehen nicht im Kurstext. Das Muster ist `negToZeroPointer` in Listing 7.2.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und den zwei Funktionen der Fassung mit Kopien. Eine Demo gibt es hier nicht.
- Eine Liste nennt, was an der Fassung mit Kopien stört: der Umweg über `a = swap(a, i)` und die 910 kopierten Elemente.
- Unter „Was dazugehört“ stehen Links auf die Lektion „Zeiger in Go“ und auf den Exkurs „Das Kopier-Problem“.

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
- Ein Codeblock zeigt `swap` mit dem Parameter `pa *[10]int`. Eine Liste nennt die vier Änderungen: Parameter, Rumpf mit `pa[i]`, keine Rückgabe, Aufruf `swap(&a, i)`.
- Ein gelber Kasten sagt: Wer `&` vergisst, bekommt die Meldung `cannot use a (variable of type [10]int) as *[10]int value in argument to swap`. Das vergessene `a =` der Fassung mit Kopien blieb unbemerkt.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4.1
- 7.4

## Schritt 2: isGreater mit einem Zeiger

Reiter: vergleich

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: `isGreater` ändert das Array nicht. Die Fassung mit der Kopie liefert das richtige Ergebnis, kopiert aber bei 45 Vergleichen 450 Elemente. Für den Zeiger sprechen der Aufwand und die Einheitlichkeit.
- Ein Codeblock zeigt `isGreater` mit dem Parameter `pa *[10]int`.
- Eine Tabelle stellt die zwei Fassungen gegenüber: Aufruf `isGreater(a, i)` mit 450 kopierten Elementen, Aufruf `isGreater(&a, i)` mit keinem.
- Ein gelber Kasten sagt: Am Aufruf `isGreater(&a, i)` sieht man nicht, ob die Funktion das Array ändert.
- Die Abwägung steht nicht im Kurstext.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.2

## Schritt 3: Ein Durchgang

Reiter: durchgang

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt 17 Zeilen: `isGreater` in den Zeilen 1 bis 3, `swap` in den Zeilen 5 bis 7 und `main` mit einer Schleife für einen Durchgang in den Zeilen 9 bis 17.
- Darunter stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Im Rahmen von `isGreater` und von `swap` steht kein Array. Dort steht der Zeiger `pa` mit dem Wert „Adresse von a“. Es gibt nur ein Array, im Rahmen von `main`.
- Grün markiert sind im Array von `main` die zwei Elemente, die der letzte Aufruf gelesen oder getauscht hat.
- Ein Schritt der Demo ist ein Aufruf von `isGreater` oder ein Aufruf von `swap`. Die Demo hat 19 Schritte.
- Der erste Durchgang: 9 Vergleiche, 8 Tausche, kein kopiertes Element. Die Ausgabe am Ende ist `[7 19 42 3 25 71 14 60 9 88]`.

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

- Zu sehen ist das ganze Programm mit den zwei Funktionen und `main` mit zwei Schleifen. In der inneren Schleife steht `if isGreater(&a, i) { swap(&a, i) }`.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe. Sie ist die gleiche wie in der Fassung mit Kopien und endet mit `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Eine Tabelle „Die Bilanz“ vergleicht drei Fassungen: Kapitel 6 mit allem in `main`, Funktionen mit Wertübergabe, Funktionen mit Zeigern. Alle haben 45 Vergleiche und 23 Tausche. Kopierte Elemente: 0, 910 und 0.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: die lange Form `(*pa)[i]` schreiben, eine Funktion `sortiere(pa *[10]int)` schreiben, die Ausgabe von `fmt.Println(&a)` vorhersagen. Die Antwort auf die dritte steht im Kurstext in Abschnitt 7.4: `fmt.Println` gibt für die Adresse eines Arrays das Zeichen `&` und dahinter das Array aus.
- Eine Demo gibt es hier nicht.

### Handgriffe

- „Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit sechs Zeilen: Abschnitt 5.3.2, 7.2, 7.3, 7.4 mit Hinweis 7.4, 7.4.1 und 8.4.2, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Die Funktionen `isGreater` und `swap` kommen im Kurstext nicht vor.
- Unter „Die Reihe“ stehen Links auf die Programmierübungen davor und danach.
- Der Reiter sagt außerdem: Die Abwägung zwischen Kopie und Zeiger für eine Funktion, die nur liest, steht nicht im Kurstext.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4
- 8.4.2
