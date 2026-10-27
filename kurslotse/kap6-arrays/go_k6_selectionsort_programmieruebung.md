# Programmierübung: Sortieren mit Selection Sort

Die Seite ist eine Programmierübung. Zehn ganze Zahlen sollen aufsteigend sortiert werden, mit dem Verfahren Selection Sort: erst das kleinste Element suchen, dann mit einem einzigen Tausch an seinen Platz bringen. Die Übung führt in acht Schritten zum Programm, vergleicht es mit Bubble Sort und lässt es im Editorlotsen ausprobieren. Die Seite gehört zu Kapitel 6 des Kurstexts. Selection Sort selbst kommt im Kurstext nicht vor, die Übung ist eine Ergänzung des Mentors.

Die Seite hat neun Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist immer `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`. Sortiert lautet es `[3 7 9 14 19 25 42 60 71 88]`.
- Links steht der sortierte Teil des Arrays, rechts der Rest. Am Anfang ist der sortierte Teil leer.
- `pos` ist die erste Stelle des Rests. Sie bekommt in dieser Runde ihren Wert.
- `minIndex` ist der Kandidat: der Index des kleinsten Werts, den das Programm im Rest bisher gesehen hat. `minIndex` enthält einen Index und nicht den Wert. Den Wert liefert `a[minIndex]`.
- Eine Runde: den kleinsten Wert des Rests suchen, dann mit dem Element an der Stelle `pos` tauschen.
- Der Tausch ist eine Tupel-Zuweisung: `a[pos], a[minIndex] = a[minIndex], a[pos]`.
- Die Variable heißt `minIndex` und nicht `min`, weil es seit Go 1.21 eine eingebaute Funktion `min` gibt.
- Das Programm kommt ohne Funktionen aus. Es steht ganz in `main`.
- Alle Programme und Ausgaben sind mit Go 1.26 nachgeprüft. Die Demos stellen den Lauf im Browser nach. Ausführen und ändern lässt sich das Programm nur in Schritt 7, „Selbst ausprobieren“, wenn der Editorlotse die Lektion ausliefert.
- Schritt 2 und 4 beginnen mit einem blauen Kasten mit einer Frage. Darunter steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein Codeblock zeigt das Array am Start und sortiert.
- Der Text stellt die zwei Verfahren gegenüber: Bubble Sort tauscht Nachbarn sofort und braucht für die zehn Zahlen 23 Tausche. Selection Sort wählt erst aus und tauscht dann, es braucht 7 Tausche. Das englische Wort selection heißt Auswahl.
- Der Teil „Womit das Programm auskommt“ nennt vier Stellen des Kurstexts mit Links: Kapitel 6, Abschnitt 5.5.3.2, Abschnitt 5.5.1 und Abschnitt 5.3.2.
- Eine Liste sagt, was die acht Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links folgen, auch zur Programmierübung zu Bubble Sort.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 2.

### Abschnitte des Kurstexts

- 6

## Schritt 1: Die Idee

Reiter: idee

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Teil „Eine Runde“ nennt drei Handgriffe: die Stelle `pos` nehmen, im Rest ab `pos` das kleinste Element suchen, es mit dem Element an der Stelle `pos` tauschen.
- Der Codeblock zeigt die ersten drei Runden in Worten. Runde 1: Der kleinste Wert ist 3, er wird mit der 42 an der Stelle 0 getauscht, das Array ist danach `[3 7 19 88 42 25 71 14 60 9]`. Runde 2: Der kleinste Wert im Rest ist 7 und steht schon an der Stelle 1. Runde 3: Der kleinste Wert im Rest ist 9, er wird mit der 19 an der Stelle 2 getauscht, das Array ist danach `[3 7 9 88 42 25 71 14 60 19]`.
- Zwei Aufgaben folgen daraus für das Programm: suchen, das ist Schritt 2, und wiederholen, das ist Schritt 4.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Der Codeblock ist fest.

### Abschnitte des Kurstexts

- 6.1

## Schritt 2: Das Minimum suchen

Reiter: minimum

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Ein Codeblock zeigt die Suche: `minIndex := 0`, dann `for i := 1; i < len(a); i++` mit `if a[i] < a[minIndex] { minIndex = i }`. Die Ausgabe: Der Index des kleinsten Werts ist 4, der kleinste Wert ist 3.
- Der Teil „Der Kandidat“: Am Anfang ist der Kandidat der Index 0. Die Schleife beginnt bei 1 und vergleicht jedes Element mit dem Kandidaten. Ist `a[i]` kleiner, wird `i` der neue Kandidat.
- Ein gelber Kasten sagt, dass `minIndex` den Index enthält und nicht den Wert. Das Programm braucht den Index, weil es später an dieser Stelle tauschen will.
- Die Demo zeigt das Programm als nummerierte Liste mit neun Zeilen auf dunklem Grund. Die Zeile, die zuletzt lief, ist hell unterlegt: Zeile 3 beim Setzen des ersten Kandidaten, Zeile 5 bei einem Vergleich ohne neuen Kandidaten, Zeile 6 bei einem neuen Kandidaten.
- Darunter stehen vier Knöpfe: „nächster Schritt“, „ein Schritt zurück“, „Runde zu Ende“ und „von vorn“.
- Das Array steht als Reihe von zehn Kästen da, über jedem Kasten sein Index, unter manchen ein Schild mit `i` oder `minIndex`.
- Farben: Blau unterlegt ist das Element, das gerade mit dem Kandidaten verglichen wird. Gelb ist der Kandidat.
- Ein Kasten darunter nennt den letzten Schritt in Worten und die Zahl der Vergleiche.
- Die Demo hat zehn Schritte: den ersten Kandidaten setzen und neun Vergleiche. `minIndex` wechselt zweimal: bei `i` gleich 1 von 0 auf 1, weil 7 kleiner ist als 42, und bei `i` gleich 4 von 1 auf 4, weil 3 kleiner ist als 7. Danach bleibt der Kandidat. Das Array ändert sich in diesem Schritt nicht.

### Handgriffe

- „nächster Schritt“ wählen.
- „ein Schritt zurück“ wählen.
- „Runde zu Ende“ wählen: Alle übrigen Vergleiche laufen auf einmal.
- „von vorn“ wählen.
- Die Auflösung der Frage aufklappen.
- Die Werte des Arrays und das Programm lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 6.1
- 5.5.3.2

## Schritt 3: Eine Runde

Reiter: runde

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und drei Codeblöcken. Eine Demo gibt es hier nicht.
- Das erste Programm sucht ab der Stelle `pos` und tauscht danach: `pos := 0`, `minIndex := pos`, die Schleife beginnt bei `i := pos + 1`, nach der Schleife steht der Tausch. Die Ausgabe: `Start: [42 7 19 88 3 25 71 14 60 9]` und `Nach Runde 1: [3 7 19 88 42 25 71 14 60 9]`.
- Erst suchen, dann tauschen: Die Schleife ändert das Array nicht, sie entscheidet nur.
- Der zweite Codeblock wiederholt die Runde mit `pos = 1`. Die Ausgabe ist wieder `[3 7 19 88 42 25 71 14 60 9]`.
- Die zwei Runden unterscheiden sich nur im Wert von `pos`. Daraus wird in Schritt 4 die äußere Schleife.
- In Runde 2 steht die 7 schon an der Stelle 1. Dann ist `minIndex` gleich `pos`, und das Programm tauscht das Element mit sich selbst. Das ist erlaubt und ändert nichts.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Codeblöcke sind fest.

### Abschnitte des Kurstexts

- 5.3.2
- 6.1

## Schritt 4: Zwei Schleifen

Reiter: schleifen

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Ein Codeblock zeigt die zwei Schleifen: außen `for pos := 0; pos < len(a)-1; pos++`, darin `minIndex := pos`, die innere Schleife ab `pos + 1`, dann `if minIndex != pos` mit dem Tausch, dann `fmt.Printf("Nach Runde %d: %v\n", pos+1, a)`.
- Die äußere Schleife läuft bei zehn Elementen neunmal. Die letzte Stelle braucht keine Runde, dort bleibt der größte Wert übrig.
- `if minIndex != pos` spart den Tausch, wenn der kleinste Wert schon an der Stelle `pos` steht.
- Die Demo ist gebaut wie in Schritt 2, mit dem Programm aus 14 Zeilen und allen neun Runden. Dazu kommen ein Schild `pos` unter dem Element an der Stelle `pos` und ein dunkles Feld mit der Ausgabe.
- Farben in dieser Demo: Blau unterlegt ist das Element, das gerade verglichen wird. Gelb ist der Kandidat. Rot ist nach einem Tausch das Element, das seinen Platz mit der Stelle `pos` gewechselt hat. Grün sind die Werte, die an ihrem Platz stehen.
- Die hell unterlegte Zeile: Zeile 4 beim Setzen des Kandidaten, Zeile 6 bei einem Vergleich ohne neuen Kandidaten, Zeile 7 bei einem neuen Kandidaten, Zeile 11 beim Tausch, Zeile 10, wenn nicht getauscht wird.
- Die Demo hat 63 Schritte: je Runde einmal den Kandidaten setzen, die Vergleiche und den Abschluss der Runde. Zusammen sind es 45 Vergleiche und 7 Tausche.
- Die Ausgabe am Ende: Nach Runde 1 `[3 7 19 88 42 25 71 14 60 9]`. Nach Runde 2 unverändert. Nach Runde 3 `[3 7 9 88 42 25 71 14 60 19]`. Nach Runde 4 `[3 7 9 14 42 25 71 88 60 19]`. Nach Runde 5 `[3 7 9 14 19 25 71 88 60 42]`. Nach Runde 6 unverändert. Nach Runde 7 `[3 7 9 14 19 25 42 88 60 71]`. Nach Runde 8 `[3 7 9 14 19 25 42 60 88 71]`. Nach Runde 9 `[3 7 9 14 19 25 42 60 71 88]`.
- In Runde 2 und in Runde 6 wird nicht getauscht.

### Handgriffe

- „nächster Schritt“ wählen.
- „ein Schritt zurück“ wählen.
- „Runde zu Ende“ wählen: Die laufende Runde läuft bis zu ihrem Abschluss.
- „von vorn“ wählen.
- Die Auflösung der Frage aufklappen.
- Die Werte des Arrays und das Programm lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.5.3.2
- 6.1

## Schritt 5: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Oben steht der Knopf „Programm kopieren“.
- Das ganze Programm steht als nummerierte Liste mit 22 Zeilen auf dunklem Grund, mit `package main`, `import "fmt"` und `func main()`.
- Darunter steht die Ausgabe: `Start:`, die neun Zeilen `Nach Runde …` und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Das Programm braucht 45 Vergleiche und 7 Tausche.
- Der Teil „So startest du das Programm“: einen leeren Ordner anlegen, darin die Datei `selectionsort.go`, das Programm hineinkopieren, im Terminal `go run selectionsort.go` eingeben.
- Der Teil „Zum Weiterdenken“ stellt drei Aufgaben ohne Lösung: absteigend sortieren, Vergleiche und Tausche zählen, die Zeile `if minIndex != pos {` löschen.
- Fragt der Student nach einer dieser drei Aufgaben, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm. Hinweise: Für absteigend dreht sich der Vergleich um, und die Variable heißt dann besser `maxIndex`. Zum Zählen braucht es zwei Variablen vom Typ `int`, die vor der äußeren Schleife deklariert sind. Ohne das `if` tauscht das Programm in Runde 2 und 6 ein Element mit sich selbst, die Ausgabe bleibt gleich.

### Handgriffe

- „Programm kopieren“ wählen: Das Programm liegt dann in der Zwischenablage.
- Das Programm läuft nicht auf der Seite. Der Student führt es auf seinem Rechner aus.

### Abschnitte des Kurstexts

- 6
- 5.3.2

## Schritt 6: Vergleich mit Bubble Sort

Reiter: vergleich

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle und Text. Eine Demo gibt es hier nicht.
- Die Tabelle stellt Bubble Sort und Selection Sort in fünf Fragen gegenüber: was verglichen wird, wann getauscht wird, wo der sortierte Teil wächst, wie viele Vergleiche und wie viele Tausche die zehn Zahlen brauchen.
- Bubble Sort vergleicht Nachbarn, tauscht sofort, der sortierte Teil wächst hinten. 45 Vergleiche, 23 Tausche.
- Selection Sort vergleicht mit dem Kandidaten, tauscht einmal am Ende der Runde, der sortierte Teil wächst vorn. 45 Vergleiche, 7 Tausche.
- Beide Verfahren brauchen für n Elemente n · (n − 1) / 2 Vergleiche. Bei doppelt so vielen Elementen sind es etwa viermal so viele. Für große Datenmengen sind beide zu langsam. Ein Link führt zum Exkurs zur Laufzeit.
- Selection Sort tauscht höchstens einmal je Runde. Das zählt, wenn ein Tausch teuer ist.
- Selection Sort braucht auch für ein schon sortiertes Array alle Vergleiche.
- Die Schreibweise O(n²) steht nicht auf dieser Seite. Sie wird im Exkurs zur Laufzeit erklärt.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2

## Schritt 7: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Im Editorlotsen steht ein Editor mit dem Knopf „Ausführen“. Auf GitHub Pages ist der Reiter grau. Dort sieht der Student das Programm und seine Ausgabe, festgehalten mit Go 1.26.1, und einen Satz, wie er die Lektion im Editorlotsen öffnet.
- Ein Programm: „Selection Sort“. Es ist das ganze Programm aus Schritt 5. Die Ausgabe ist die gleiche wie dort.
- Der Text schlägt vor, im Editor die Aufgaben unter „Zum Weiterdenken“ aus Schritt 5 zu lösen. Fragt der Student nach einer Lösung, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Der Stand der Demo sagt, welches Programm gewählt ist, ob der Student es geändert hat und was der letzte Lauf ausgegeben oder gemeldet hat. Die Programmzeilen sind das Programm im Editor.
- Am Ende des Schritts unter „Nächste Schritte“ steht ein Link zur HTML-Seite „Selection Sort mit Funktionen“ in Kapitel 7 und nennt Quicksort und Mergesort als schnellere Verfahren, die im Kurstext nicht behandelt werden.

### Handgriffe

- Das Programm ansehen.
- Im Editor schreiben, „Ausführen“, „Stopp“ oder „Programm zurücksetzen“ wählen, wenn der Editorlotse die Lektion ausliefert.

### Abschnitte des Kurstexts

- 6.2

## Schritt 8: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle und Text. Eine Demo gibt es hier nicht.
- Die Tabelle nennt fünf Abschnitte des Kurstexts und was das Programm daraus verwendet: 5.3.2 Tupel-Zuweisung, 5.5.1 if-Anweisung, 5.5.3.2 for-Klausel, 6 Arrays, 6.1 Index-Position.
- Der Teil „Selection Sort im Kurstext“: Selection Sort kommt im Kurstext nicht vor. Das einzige Sortierverfahren dort ist Bubblesort in Abschnitt 8.4.2. In Abschnitt 8.4.2 steht, dass der Vorkurs zwei einfache Sortierverfahren vorgestellt hat. Welche das sind, steht dort nicht.
- Der Teil „Was über den Kurstext hinausgeht“: das Verfahren selbst, der Vergleich mit Bubble Sort und die Ausgabe mit `%d` und `%v`.
- Der Reiter nennt außerdem: In Abschnitt 8.4.2 steht nur, dass der Vorkurs zwei einfache Sortierverfahren vorgestellt hat, und die Zahl der Vergleiche von Bubblesort. Alle Programme sind eigene Beispiele. Die Variable heißt minIndex, weil es seit Go 1.21 eine eingebaute Funktion min gibt.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links folgen.

### Abschnitte des Kurstexts

- 6
- 8.4.2
