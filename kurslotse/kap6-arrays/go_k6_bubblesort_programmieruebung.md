# Programmierübung: Sortieren mit Bubble Sort

Die Seite ist eine Programmierübung. Zehn ganze Zahlen sollen aufsteigend sortiert werden. Die Übung führt in sechs Schritten zum Programm: zehn einzelne Variablen, ein Array mit einem Durchgang, drei ausgeschriebene Durchgänge, zwei verschachtelte Schleifen, das ganze Programm, der Bezug zum Kurstext. Die Seite gehört zu Kapitel 6 des Kurstexts. Das Verfahren Bubblesort selbst wird im Kurstext erst in Abschnitt 8.4.2 behandelt.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist immer `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`. Sortiert lautet es `[3 7 9 14 19 25 42 60 71 88]`.
- Bubble Sort kennt einen einzigen Handgriff: zwei Nachbarn `a[i]` und `a[i+1]` vergleichen und tauschen, wenn der linke größer ist als der rechte.
- Der Tausch ist eine Tupel-Zuweisung: `a[i], a[i+1] = a[i+1], a[i]`.
- Ein Durchgang geht einmal von links nach rechts durch das Array. Danach steht der größte Wert des durchlaufenen Teils an dessen Ende.
- Auf der Seite heißt es „Durchgang“. Im Englischen heißt das „pass“. Auch der Kurstext schreibt in Abschnitt 8.4.2 „Durchgang“.
- Das Programm kommt ohne Funktionen aus. Es steht ganz in `main`.
- Alle Programme und Ausgaben sind mit Go 1.26 nachgeprüft. Auf der Seite läuft kein Go. Die Demos stellen den Lauf im Browser nach.
- Schritt 2 und 4 beginnen mit einem blauen Kasten mit Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein Codeblock zeigt das Array unsortiert und sortiert.
- Der Teil „Die Idee von Bubble Sort“ nennt den einen Handgriff und das Bild der Luftblase: Ein großer Wert steigt Tausch für Tausch nach hinten auf.
- Der Teil „Womit das Programm auskommt“ nennt vier Stellen des Kurstexts mit Links: Array und Index in Kapitel 6, die for-Klausel in Abschnitt 5.5.3.2, die if-Anweisung in Abschnitt 5.5.1 und die Tupel-Zuweisung in Abschnitt 5.3.2.
- Eine Liste sagt, was die sechs Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den vier HTML-Seiten folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 2.

### Abschnitte des Kurstexts

- 6
- 8.4.2

## Schritt 1: Zehn Variablen

Reiter: variablen

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und zwei Codeblöcken. Eine Demo gibt es hier nicht.
- Das Programm legt zehn Variablen `a0` bis `a9` an, gibt sie aus, tauscht `a0` und `a1` mit `if a0 > a1 { a0, a1 = a1, a0 }` und gibt sie noch einmal aus.
- Die Ausgabe: `Unsortiert: 42 7 19 88 3 25 71 14 60 9` und `Nach einem Tausch: 7 42 19 88 3 25 71 14 60 9`.
- Der Teil „Wo der Ansatz endet“: Jeder Vergleich braucht eine eigene if-Anweisung, für zehn Zahlen sind es 45. Eine Schleife kann nicht von `a0` zu `a1` weiterzählen. Für elf Zahlen müsste man alles neu schreiben.
- Mit einem Array wird aus `a1` der Index-Ausdruck `a[1]`, und in den eckigen Klammern darf eine Variable stehen.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Codeblöcke sind fest.

### Abschnitte des Kurstexts

- 6
- 5.3.2

## Schritt 2: Ein Durchgang

Reiter: durchgang

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen.
- Ein Codeblock zeigt den Durchgang: `for i := 0; i < len(a)-1; i++` mit `if a[i] > a[i+1]` und dem Tausch im Rumpf.
- Der Teil „Warum die Schleife bei len(a)-1 endet“: Der Rumpf greift auf `a[i+1]` zu. Das letzte Element hat keinen rechten Nachbarn. Der größte Index, mit dem der Rumpf läuft, ist 8. Mit `i < len(a)` würde das Programm bei `i` gleich 9 auf `a[10]` zugreifen und mit `index out of range` abbrechen.
- Die Demo zeigt das Programm als nummerierte Liste mit acht Zeilen auf dunklem Grund. Die Zeile, die zuletzt lief, ist hell unterlegt: Zeile 4 mit `if`, wenn nicht getauscht wurde, Zeile 5 mit dem Tausch, wenn getauscht wurde.
- Darunter stehen vier Knöpfe: „nächster Vergleich“, „ein Vergleich zurück“, „Durchgang zu Ende“ und „von vorn“.
- Das Array steht als Reihe von zehn Kästen da, über jedem Kasten sein Index.
- Farben: Gelb sind zwei Nachbarn, die verglichen und nicht getauscht wurden. Rot sind zwei Nachbarn, die getauscht wurden. Grün ist ein Wert, der an seinem Platz steht.
- Ein Kasten darunter nennt den letzten Vergleich in Worten und die Zahl der Vergleiche und Tausche bisher.
- Der Durchgang hat neun Vergleiche. Nach dem Durchgang lautet das Array `[7 19 42 3 25 71 14 60 9 88]`. Die 88 steht am Ende. Sortiert ist das Array noch nicht.
- Die neun Vergleiche: 42 und 7 werden getauscht. 42 und 19 werden getauscht. 42 und 88 bleiben. 88 und 3, 88 und 25, 88 und 71, 88 und 14, 88 und 60, 88 und 9 werden getauscht. Das sind acht Tausche.

### Handgriffe

- „nächster Vergleich“ wählen: Das Programm vergleicht das nächste Paar.
- „ein Vergleich zurück“ wählen.
- „Durchgang zu Ende“ wählen: Alle übrigen Vergleiche des Durchgangs laufen auf einmal.
- „von vorn“ wählen.
- Die Auflösungen der zwei Fragen aufklappen.
- Die Werte des Arrays und das Programm lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 6.1
- 5.5.3.2

## Schritt 3: Drei Durchgänge

Reiter: ausgeschrieben

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und zwei Codeblöcken. Eine Demo gibt es hier nicht.
- Das Programm schreibt drei Durchgänge einzeln aus, drei Schleifen hintereinander, und gibt das Array nach jedem aus.
- Die Ausgabe: `Start: [42 7 19 88 3 25 71 14 60 9]`, `Nach Durchgang 1: [7 19 42 3 25 71 14 60 9 88]`, `Nach Durchgang 2: [7 19 3 25 42 14 60 9 71 88]`, `Nach Durchgang 3: [7 3 19 25 14 42 9 60 71 88]`.
- Hinten wächst der sortierte Teil: erst 88, dann 71 und 88, dann 60, 71 und 88.
- Die 3 steht am Anfang beim Index 4 und rückt in jedem Durchgang nur einen Platz nach vorn.
- Die drei Schleifen unterscheiden sich nur in der Bedingung: `len(a)-1`, `len(a)-2`, `len(a)-3`. Jeder Durchgang darf einen Vergleich früher aufhören, weil die Werte am Ende schon richtig stehen.
- Dreimal fast der gleiche Code ist der Grund für die äußere Schleife in Schritt 4.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Codeblöcke sind fest.

### Abschnitte des Kurstexts

- 6.1
- 8.4.2

## Schritt 4: Zwei Schleifen

Reiter: schleifen

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Ein Codeblock zeigt die zwei Schleifen: außen `for durchgang := 0; durchgang < len(a)-1; durchgang++`, innen `for i := 0; i < len(a)-1-durchgang; i++`, nach der inneren Schleife `fmt.Printf("Nach Durchgang %d: %v\n", durchgang+1, a)`.
- Die äußere Schleife läuft bei zehn Elementen neunmal. Die innere Schleife ist der Durchgang aus Schritt 2. `durchgang` zählt ab 0, die Ausgabe zählt mit `durchgang+1` ab 1.
- Die Demo ist gebaut wie in Schritt 2, mit dem Programm aus zehn Zeilen und allen neun Durchgängen. Dazu kommt ein dunkles Feld mit der Ausgabe des Programms.
- Die hell unterlegte Zeile ist Zeile 5 mit `if`, wenn nicht getauscht wurde, und Zeile 6 mit dem Tausch, wenn getauscht wurde.
- Die Demo hat 45 Vergleiche und 23 Tausche.
- Die Ausgabe am Ende: Nach Durchgang 1 `[7 19 42 3 25 71 14 60 9 88]`. Nach Durchgang 2 `[7 19 3 25 42 14 60 9 71 88]`. Nach Durchgang 3 `[7 3 19 25 14 42 9 60 71 88]`. Nach Durchgang 4 `[3 7 19 14 25 9 42 60 71 88]`. Nach Durchgang 5 `[3 7 14 19 9 25 42 60 71 88]`. Nach Durchgang 6 `[3 7 14 9 19 25 42 60 71 88]`. Nach Durchgang 7 `[3 7 9 14 19 25 42 60 71 88]`. Durchgang 8 und 9 ändern nichts mehr.
- Der Teil „Wie viel Arbeit das ist“: 9 + 8 + … + 1 = 45 Vergleiche. Für n Elemente sind es n · (n − 1) / 2. Bei doppelt so vielen Elementen sind es etwa viermal so viele Vergleiche. Ein Link führt zum Exkurs zur Laufzeit.

### Handgriffe

- „nächster Vergleich“ wählen.
- „ein Vergleich zurück“ wählen.
- „Durchgang zu Ende“ wählen: Der laufende Durchgang läuft zu Ende. Steht der Lauf am Ende eines Durchgangs, läuft der nächste Durchgang ganz.
- „von vorn“ wählen.
- Die Auflösung der Frage aufklappen.
- Die Werte des Arrays und das Programm lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 8.4.2
- 5.5.3.2

## Schritt 5: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Oben stehen zwei Knöpfe für die Werte des Arrays: „feste Werte“ und „Zufallswerte“. Darunter steht der Knopf „Programm kopieren“.
- Das ganze Programm steht als nummerierte Liste auf dunklem Grund, mit `package main`, `import` und `func main()`.
- Die Fassung mit festen Werten hat 18 Zeilen. Darunter steht ihre Ausgabe: `Unsortiert:`, die neun Zeilen `Nach Durchgang …` und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Die Fassung mit Zufallswerten hat 24 Zeilen. Sie importiert zusätzlich `math/rand`, legt mit `var a [10]int` ein Array aus zehn Nullen an und füllt es mit `for i := range a { a[i] = rand.Intn(100) + 1 }`. `rand.Intn(100)` liefert eine Zufallszahl von 0 bis 99, mit `+ 1` wird daraus 1 bis 100.
- Der Teil „So startest du das Programm“: einen leeren Ordner anlegen, darin die Datei `bubblesort.go`, das Programm hineinkopieren, im Terminal `go run bubblesort.go` eingeben.
- Der Teil „Zum Weiterdenken“ stellt drei Aufgaben ohne Lösung: früher aufhören, wenn nichts mehr getauscht wird. Absteigend sortieren. Die Länge auf 20 ändern.
- Fragt der Student nach einer dieser drei Aufgaben, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm. Hinweise: Für das frühe Aufhören merkt sich das Programm in einer Variable vom Typ `bool`, ob ein Durchgang getauscht hat. Für absteigend dreht sich der Vergleich um. Für die Länge 20 ändert sich nur die Zahl im Typ, weil überall sonst `len(a)` steht.

### Handgriffe

- Zwischen „feste Werte“ und „Zufallswerte“ wählen.
- „Programm kopieren“ wählen: Das gezeigte Programm liegt dann in der Zwischenablage.
- Das Programm läuft nicht auf der Seite. Der Student führt es auf seinem Rechner aus.

### Abschnitte des Kurstexts

- 6
- 6.2

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle und Text. Eine Demo gibt es hier nicht.
- Die Tabelle nennt sechs Abschnitte des Kurstexts und was das Programm daraus verwendet: 5.3.2 Tupel-Zuweisung, 5.5.1 if-Anweisung, 5.5.3.2 for-Klausel, 6 Arrays, 6.1 Index-Position, 6.2 Durchlaufen eines Arrays.
- Der Teil „Bubble Sort im Kurstext“: Im Kurstext wird Bubblesort in Abschnitt 8.4.2 behandelt, Seiten 225–226. Dort sortiert die Funktion `bubblesort` ein Array von Verbunden nach dem Namen und bekommt einen Zeiger auf das Array. Die Übung nimmt das Verfahren vorweg, mit ganzen Zahlen und ohne Funktion.
- Der Teil „Was über den Kurstext hinausgeht“: die Ausgabe mit `%d` und `%v` und das Paket `math/rand`.
- Der Reiter nennt außerdem: Abschnitt 8.4.2 nennt die Zahl der Vergleiche. Schritt 1 folgt der Begründung des Arrays in Kapitel 6, Seite 151, mit den Variablen z1 bis z6 der Lottoziehung. Alle Programme der Übung sind eigene Beispiele.
- Unter „Wie es weitergeht“ stehen Links zur Programmierübung zu Selection Sort und zur HTML-Seite „Bubble Sort mit Funktionen“ in Kapitel 7.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2
- 6
- 5.3.2
