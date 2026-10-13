# Programmierübung: Bubble Sort rekursiv

Die Seite ist eine Programmierübung. Sie ersetzt die äußere Schleife von Bubble Sort durch einen rekursiven Aufruf: Die Funktion `bubbleSort` macht einen Durchgang und ruft sich dann für ein Element weniger selbst auf. Die Seite gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.5, Seiten 185–193.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Array ist `a := [10]int{42, 7, 19, 88, 3, 25, 71, 14, 60, 9}`.
- `bubbleSort(pa *[10]int, n int)` sortiert die ersten `n` Elemente des Arrays, auf das `pa` zeigt. Der Aufruf in `main` lautet `bubbleSort(&a, len(a))`.
- Die Funktion hat drei Teile: Für `n <= 1` kehrt sie sofort zurück. Sonst macht sie einen Durchgang mit der Schleife `for i := 0; i < n-1; i++`. Danach ruft sie `bubbleSort(pa, n-1)` auf.
- Im rekursiven Aufruf steht `pa` und nicht `&pa`. `pa` ist schon die Adresse des Arrays. `&pa` wäre die Adresse des Zeigers und hätte den Typ `**[10]int`.
- Die Funktionen `isGreater(pa *[10]int, i int) bool` und `swap(pa *[10]int, i int)` stammen aus der Programmierübung „Bubble Sort mit Zeigern“. Die innere Schleife bleibt eine Schleife.
- Das Verfahren ändert sich nicht: 45 Vergleiche, 23 Tausche. Die Funktion `bubbleSort` läuft zehnmal, mit `n` von 10 bis 1.
- `bubbleSort` ist direkt rekursiv, linear rekursiv und endrekursiv: Ein Aufruf enthält höchstens einen rekursiven Aufruf, und dieser ist die letzte Operation.
- Unter den Rahmen der Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger, in ihm steht die Zahl 100, die Adresse des Arrays. Die Adressen sind erfunden. Ein `int` und ein Zeiger belegen im Modell je 8 Zellen. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke in Schritt 2: `a` in den Zellen 100 bis 179. Jeder Aufruf von `bubbleSort` belegt 16 Zellen, 8 für `pa` und 8 für `n`: der Aufruf mit n gleich 10 die Zellen 180 bis 195, der mit n gleich 9 die Zellen 196 bis 211, und so weiter bis zum Aufruf mit n gleich 1 in den Zellen 324 bis 339. In jedem `pa` steht 100. Warten zehn Aufrufe, sind 160 Zellen belegt.
- Im Kurstext steht: Bei Endrekursion muss kein Variablenwert gespeichert werden. Der Go-Compiler nutzt das zurzeit nicht aus, jeder Aufruf belegt Speicher. Die Demo zeigt deshalb für jeden Aufruf einen Rahmen.
- Eine rekursive Fassung von Bubblesort kommt im Kurstext nicht vor. Bubblesort steht dort in Abschnitt 8.4.2, mit zwei Schleifen und einem Array von Verbunden.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und den zwei Schleifen aus der Programmierübung „Bubble Sort mit Zeigern“. Eine Demo gibt es hier nicht.
- Der Text sagt: Die äußere Schleife wiederholt den gleichen Durchgang für einen immer kürzeren Anfang des Arrays. Das lässt sich rekursiv schreiben.
- Unter „Was dazugehört“ stehen Links auf die Lektion „Rekursion in Go“, auf die Programmierübung „Bubble Sort mit Zeigern“ und auf die Programmierübung „Sortieren mit Bubble Sort“ in Kapitel 6.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
- 7.3

## Schritt 1: Die Idee

Reiter: idee

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste mit vier Punkten: die Aufgabe „sortiere die ersten n Elemente“, ein Durchgang, die kleinere Aufgabe „sortiere die ersten n − 1 Elemente“, das Ende bei n gleich 1 oder 0.
- Die Funktion braucht zwei Parameter: einen Zeiger auf das Array und die Zahl n. Die Länge 10 steht im Typ, n sagt, wie viele Elemente noch zu sortieren sind.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
- 7.5.5

## Schritt 2: Die Funktion

Reiter: funktion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste erklärt die drei Teile der Funktion mit ihren Zeilen: das Ende in den Zeilen 2 bis 4, der Durchgang in den Zeilen 5 bis 9, der rekursive Aufruf in Zeile 10.
- Die Demo zeigt 17 Zeilen: `bubbleSort` in den Zeilen 1 bis 11 und `main` in den Zeilen 13 bis 17. `isGreater` und `swap` zeigt die Demo nicht.
- Darunter stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Jeder Aufruf von `bubbleSort` hat einen eigenen Rahmen mit `pa` und `n`, zum Beispiel „Funktion bubbleSort, Aufruf mit n = 8“. Die Rahmen stehen in der Reihenfolge der Aufrufe nebeneinander. Der Wert von `pa` steht in Worten da: „Adresse von a“.
- Das Array steht nur im Rahmen von `main`. Grün markiert sind die Elemente, die schon an ihrem Platz stehen.
- Ein Schritt der Demo ist ein ganzer Durchgang, ein rekursiver Aufruf oder eine Rückkehr. Die Demo hat 23 Schritte.
- Der Ablauf: Aufruf mit n gleich 10. Dann wechseln sich neunmal ein Durchgang und ein rekursiver Aufruf ab, für n von 10 bis 2. Der Aufruf mit n gleich 1 kehrt sofort zurück. Danach enden alle neun wartenden Aufrufe nacheinander, keiner hat auf dem Rückweg etwas zu tun. `main` gibt `[3 7 9 14 19 25 42 60 71 88]` aus.
- Die Tausche je Durchgang: n = 10: 8. n = 9: 5. n = 8: 3. n = 7: 3. n = 6: 2. n = 5: 1. n = 4: 1. n = 3 und n = 2: 0. Zusammen 23 Tausche und 45 Vergleiche.
- Das Array nach den Durchgängen: nach n = 10 `[7 19 42 3 25 71 14 60 9 88]`, nach n = 9 `[7 19 3 25 42 14 60 9 71 88]`, nach n = 8 `[7 3 19 25 14 42 9 60 71 88]`, nach n = 7 `[3 7 19 14 25 9 42 60 71 88]`, nach n = 6 `[3 7 14 19 9 25 42 60 71 88]`, nach n = 5 `[3 7 14 9 19 25 42 60 71 88]`, nach n = 4 `[3 7 9 14 19 25 42 60 71 88]`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Array lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.5.1
- 7.5.5

## Schritt 3: Schleife oder Rekursion

Reiter: vergleich

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Tabelle stellt das Programm mit zwei Schleifen und die rekursive Fassung in fünf Zeilen gegenüber: Was die Durchgänge zählt, wie weit ein Durchgang läuft, wann Schluss ist, wie viele Vergleiche und Tausche es sind, was gespeichert bleibt.
- Die Variable `durchgang` zählt von 0 bis 8, der Parameter `n` fällt von 10 bis 1. Es gilt n = 10 − durchgang.
- Unter „Einordnung nach Abschnitt 7.5 des Kurstexts“ stehen vier Punkte: direkte Rekursion, lineare Rekursion, Endrekursion, Terminierung.
- Ein gelber Kasten sagt: Die rekursive Fassung ist hier nicht besser als die Schleife. Sie zeigt, wie eine Schleife und ein rekursiver Aufruf einander entsprechen.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5.2
- 7.5.3
- 7.5.5

## Schritt 4: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das ganze Programm mit `isGreater`, `swap`, `bubbleSort` und `main`.
- Gegenüber der Demo hat `bubbleSort` eine Zeile mehr: `fmt.Printf("n = %2d: %v\n", n, *pa)` gibt nach jedem Durchgang das Array aus. `*pa` ist das Array, auf das `pa` zeigt.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe: `Start:`, neun Zeilen von `n = 10:` bis `n =  2:` und `Sortiert: [3 7 9 14 19 25 42 60 71 88]`.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: die drei Zeilen mit `if n <= 1` weglassen, den rekursiven Aufruf vor die Schleife stellen, auch den Durchgang rekursiv schreiben.
- Zur ersten Aufgabe: Ohne das Ende ruft die Funktion sich mit 0, −1, −2 und so weiter auf. Die Schleife läuft dann nicht mehr, die Aufrufe hören aber nicht auf. Go bricht ab mit `fatal error: stack overflow`.
- Zur zweiten Aufgabe: Das Programm sortiert dann nicht mehr richtig. Die Durchgänge laufen in umgekehrter Reihenfolge, erst über 2 Elemente, zuletzt über 10. Ein Durchgang bringt nur das größte Element ans Ende. Ein kleines Element, das spät dazukommt, wandert nur einen Platz nach vorn. Für das Array der Seite kommt `[3 7 19 25 14 42 60 71 9 88]` heraus. Die Funktion ist dann auch nicht mehr endrekursiv, weil nach dem rekursiven Aufruf noch die Schleife läuft.
- Eine Demo gibt es hier nicht.

### Handgriffe

- „Programm kopieren“ wählen.
- Dem Link zur Lektion „Rekursion in Go“ folgen.

### Abschnitte des Kurstexts

- 7.5.5
- 7.5.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit sieben Zeilen: Abschnitt 7.3, 7.4, 7.5, 7.5.2, 7.5.3, 7.5.5 und 8.4.2, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Eine rekursive Fassung von Bubblesort kommt im Kurstext nicht vor.
- Unter „Die Reihe“ stehen Links auf die drei Programmierübungen davor und auf das Quiz.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
- 8.4.2
