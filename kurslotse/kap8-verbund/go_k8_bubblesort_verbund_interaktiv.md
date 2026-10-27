# Bubblesort auf einem Verbund-Array

Die Seite erklärt die Funktion `bubblesort` aus Listing 8.2 des Kurstexts: Sie sortiert ein Array von Verbünden des Typs `person` aufsteigend nach dem Feld `name`. Die Seite gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.4.2, Seiten 225–226.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Typ ist `person` mit den Feldern `name string` und `telefonnummer int`, wie in Listing 8.1.
- Das Telefonbuch der Seite: `telbuch := [10]person{{"Meier", 4711}, {"Schulz", 1234}, {"Becker", 5522}, {"Wagner", 9081}, {"Arnold", 3310}, {"Koch", 7745}, {"Zimmer", 2063}, {"Fischer", 6398}, {"Lange", 8127}, {"Huber", 4456}}`.
- Die Funktion: `func bubblesort(a *[10]person)` mit einer äußeren Schleife `for n := len(a); n > 1; n--`, einer inneren Schleife `for i := 0; i < n-1; i++`, dem Vergleich `a[i].name > a[i+1].name` und dem Tausch `a[i], a[i+1] = a[i+1], a[i]`. Der Aufruf lautet `bubblesort(&telbuch)`.
- `a` ist ein Zeiger auf das Array. `len(a)` liefert auch für einen Zeiger auf ein Array dessen Länge. `a[i]` ist die Kurzschreibweise für `(*a)[i]`, das ist Hinweis 7.4 des Kurstexts.
- Die Tupel-Zuweisung vertauscht zwei ganze Verbünde. Name und Telefonnummer bleiben zusammen.
- Das Telefonbuch der Seite braucht 45 Vergleiche und 24 Vertauschungen. Vertauschungen je Durchgang: 6, 6, 5, 4, 2, 1, 0, 0, 0. Nach dem sechsten Durchgang ist das Array sortiert.
- Sortiert lautet das Telefonbuch: Arnold 3310, Becker 5522, Fischer 6398, Huber 4456, Koch 7745, Lange 8127, Meier 4711, Schulz 1234, Wagner 9081, Zimmer 2063.
- Go vergleicht zwei Strings Byte für Byte. Großbuchstaben haben kleinere Werte als Kleinbuchstaben. Ein Umlaut belegt zwei Bytes mit Werten über 127.
- In den Programmierübungen zu Kapitel 6 und 7 heißt das Verfahren „Bubble Sort“ und sortiert ein Array vom Typ `[10]int`. Im Kurstext heißt es „Bubblesort“. Gemeint ist das gleiche Verfahren.
- Die binäre Suche aus Abschnitt 8.4.1 setzt ein sortiertes Array voraus. Die Lektion „Binäre Suche in einem Verbund-Array“ erklärt sie, mit dem gleichen Telefonbuch in sortierter Form.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben, Meldungen und Zahlen sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays. `sort.Slice` und das Paket `sort` kommen im Kurs nicht vor.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt den Anlass: Die binäre Suche braucht ein sortiertes Array. Das Wort „binäre Suche“ verlinkt die Lektion „Binäre Suche in einem Verbund-Array“.
- Unter „Was dazugehört“ stehen Links auf die Programmierübungen „Sortieren mit Bubble Sort“ und „Bubble Sort mit Zeigern“, auf die Lektionen „Verbund-Variablen und Feld-Variablen“, „Zeiger in Go“ und „Strings: Bytes und Zeichen“ und auf den Exkurs „Das Kopier-Problem“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2
- 8.4.1

## Schritt 1: Namen vergleichen

Reiter: vergleich

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt den Typ `person` und das Telefonbuch mit zehn Einträgen.
- Fünf Knöpfe wählen einen Vergleich. Darunter stehen die Ausgabe oder die Meldung des Compilers und ein Kasten mit einem Satz.
- Vergleich 1: `"Meier" > "Becker"` ist `true`, M hat den Wert 77, B den Wert 66.
- Vergleich 2: `"Mai" > "Maier"` ist `false`, der kürzere String ist der kleinere.
- Vergleich 3: `"meier" > "Zimmer"` ist `true`, m hat den Wert 109, Z den Wert 90.
- Vergleich 4: `"Müller" > "Muster"` ist `true`, das ü belegt die Bytes 195 und 188, u hat den Wert 117.
- Vergleich 5: `telbuch[0] > telbuch[1]` übersetzt Go nicht. Die Meldung: `invalid operation: telbuch[0] > telbuch[1] (operator > not defined on struct)`. Verbünde lassen sich mit `==` und `!=` vergleichen, nicht mit `>` oder `<`.
- Ein gelber Kasten sagt: „Alphabetisch“ stimmt, solange alle Namen gleich geschrieben sind, mit großem Anfangsbuchstaben und ohne Umlaute.

### Handgriffe

- Einen der fünf Vergleiche wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Vergleich kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 8.4.2
- 4.6

## Schritt 2: Die Funktion

Reiter: funktion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `bubblesort` aus Listing 8.2, unverändert.
- Eine Liste erklärt sechs Stellen: den Parameter `a *[10]person`, `len(a)`, `a[i].name`, die innere Schleife, die äußere Schleife und die Tupel-Zuweisung.
- `n` ist die Zahl der Elemente, die noch zu sortieren sind. Nach jedem Durchgang steht ein weiteres Element an seinem Platz.
- Ein gelber Kasten sagt, was ohne Verbünde wäre: zwei Arrays, eines für Namen und eines für Nummern, und bei jedem Tausch zwei Vertauschungen.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2
- 7.4

## Schritt 3: Durchgang für Durchgang

Reiter: durchgang

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt die neun Zeilen der Funktion `bubblesort`.
- Darunter stehen die drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen für `main` und `bubblesort`, der Hauptspeicher und ein Kasten mit einem Satz. Eine Ausgabe gibt es nicht.
- Ein Element des Arrays steht im Bild als `{Name Telefonnummer}`, zum Beispiel `{Meier 4711}`.
- Im Rahmen von `bubblesort` steht kein Array. Dort stehen der Zeiger `a` mit dem Wert „Adresse von telbuch“ und die Variablen `n` und `i`.
- Im Hauptspeicher: `telbuch` ab Zelle 100, `a` in den Zellen 400 bis 407 mit dem Inhalt 100, `n` ab 408, `i` ab 416. Die Adressen sind erfunden. Wie viele Zellen ein `string` belegt, sagt die Seite nicht.
- Die Demo hat 21 Schritte: den Aufruf, den Beginn der äußeren Schleife, neun Vergleiche des ersten Durchgangs, das Ende des ersten Durchgangs, je einen Schritt für die acht weiteren Durchgänge und die Rückkehr.
- Der erste Durchgang: kein Tausch bei i = 0, 2 und 5, Tausch bei i = 1, 3, 4, 6, 7 und 8. Danach steht Zimmer am Index 9.
- Der Kasten nennt nach jedem Schritt die Zahl der Vergleiche und Vertauschungen bisher. Am Ende sind es 45 und 24.
- Grün sind die zwei Elemente des letzten Vergleichs oder, nach einem ganzen Durchgang, die Elemente an ihrem endgültigen Platz.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Telefonbuch lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 8.4.2
- 7.4

## Schritt 4: Aufwand

Reiter: aufwand

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt die Formel: (n − 1) + (n − 2) + … + 1 = n · (n − 1) / 2 Vergleiche.
- Eine Tabelle nennt 45 Vergleiche für 10 Elemente, 4950 für 100 und 499500 für 1000.
- Bubblesort hat quadratische Laufzeit. Die Zahl der Vertauschungen kann im ungünstigsten Fall genauso groß werden.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: den Abbruch einbauen, wenn ein Durchgang ohne Vertauschung bleibt, nach dem Feld `telefonnummer` sortieren, absteigend sortieren. Die erste ist die Übungsaufgabe des Kurstexts.
- Zur ersten Aufgabe: Der Weg führt über eine Variable vom Typ `bool`, die sich merkt, ob im Durchgang vertauscht wurde. Mit dem Abbruch läuft nach dem sechsten Durchgang noch der siebte, er bleibt ohne Vertauschung. Der achte und der neunte Durchgang entfallen, das spart 3 von 45 Vergleichen.
- Zur zweiten Aufgabe: Es ändert sich nur der Vergleich, `a[i].telefonnummer > a[i+1].telefonnummer`.
- Zur dritten Aufgabe: Das Zeichen `>` im Vergleich entscheidet.
- Eine Demo gibt es hier nicht.
- Am Ende des Schritts unter „Nächste Schritte“ stehen Links auf die Programmierübung „Ein Telefonbuch als Verbund-Array“, auf die Lektion „Verkettete Listen in Go“ und auf die vier Programmierübungen der Reihe Bubble Sort.

### Handgriffe

- Die „Auflösung“ aufklappen. Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2
- 8.4.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Namen vergleichen, 8.4.1 und 8.4.2, Seiten 221–226: Der Typ person stammt aus Listing 8.1. Die zehn Einträge und die fünf Vergleiche sind eigene Beispiele, geprüft mit Go 1.26.
  - 2 Die Funktion, 8.4.2, Seite 226: Listing 8.2: Listing 8.2 stammt unverändert aus dem Kurstext. Die Lektion gibt nichts dazu.
  - 3 Durchgang für Durchgang, 8.4.2, Seite 226: Listing 8.2: Die Funktion aus Listing 8.2 stammt aus dem Kurstext. Die Zahl der Vergleiche und der Vertauschungen ist mit Go 1.26 gezählt. Die Adressen gehören zum Modell der Lektion.
  - 4 Aufwand, 8.4.2, Seiten 225–226: Die Formel, die quadratische Laufzeit und die Übungsaufgabe zum Abbruch stehen im Kurstext. Die Tabelle und die zwei weiteren Aufgaben sind ergänzt.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Kapitel 6, Seiten 151–162, Bubble Sort auf einem Array aus Zahlen: Lektion: Sortieren mit Bubble Sort.
  - Abschnitt 7.3, Seiten 171–172, Zeiger auf ein Array: Lektion: Bubble Sort mit Zeigern.
  - Abschnitt 4.5.2, Seiten 96–98, Strings vergleichen: Lektion: Logische Operatoren und Vergleiche in Go, Schritt 2.
  - Ohne Abschnitt im Kurstext, Laufzeit: Wie die Zahl der Schritte mit der Datenmenge wächst, zeigt der Exkurs Wie die Laufzeit mit der Datenmenge wächst.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.5.2
- 7.3
- 8.4.1
- 8.4.2
