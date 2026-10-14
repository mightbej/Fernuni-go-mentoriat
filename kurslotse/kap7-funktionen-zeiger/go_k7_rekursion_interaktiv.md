# Rekursion in Go

Die Seite erklärt Rekursion an den Funktionen des Kurstexts: Fakultät iterativ, rekursiv und endrekursiv, Fibonacci-Zahlen, Türme von Hanoi und die Funktion `zahlenfolge`. Sie gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.5, Seiten 185–193.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ruft eine Funktion in ihrem Rumpf sich selbst auf, heißt das direkte Rekursion. Rufen sich mehrere Funktionen im Kreis auf, heißt das indirekte Rekursion.
- Eine Funktion ohne Selbstaufruf, die mit einer Schleife arbeitet, heißt iterativ.
- Jeder Aufruf einer Funktion hat eigene Parameter und eigene lokale Variablen. Ein Aufruf, der auf das Ergebnis eines rekursiven Aufrufs wartet, behält seine Werte, bis er weiterläuft.
- Eine rekursive Funktion braucht einen Zweig ohne rekursiven Aufruf, und jeder rekursive Aufruf muss diesem Zweig näher kommen.
- Die Wörter „Rahmen“ und „Aufrufstapel“ kommen im Kurstext nicht vor. Dort heißt es: Die Werte müssen als Variablenwerte der verschiedenen Instanzen des Funktionsaufrufs gespeichert werden. „Rahmen“ ist das Wort dieser Seite für den Kasten eines Aufrufs.
- Die Namen der Funktionen tragen Umlaute wie im Kurstext: `fakultätIte`, `fakultätRek`, `fakultätEnd`. Go erlaubt das.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Zahlen sind mit Go 1.26 geprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text sagt: Funktionen dienen der Wiederverwendung und der Gliederung eines Programms. Rekursion gliedert das Problem: Die Funktion führt eine Aufgabe auf eine kleinere Aufgabe der gleichen Art zurück.
- Unter „Was aus früheren Kapiteln dazugehört“ stehen vier Punkte mit Links: Abschnitt 2.6 mit der Lektion „Unterprogramme und Funktionen“, Abschnitt 3.4, Abschnitt 5.5.3.3 mit der Lektion „for-Anweisung mit for-Klausel“, Abschnitt 7.1 mit der Lektion „Funktionen in Go“.
- Unter „Die Schritte“ steht je ein Satz zu den fünf Schritten.
- Im Kurstext steht in Abschnitt 7.5 außerdem: Eine Funktion ohne Rückgabewert ist eigentlich ein Unterprogramm, auch Prozedur genannt.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
- 2.6

## Schritt 1: Iterativ und rekursiv

Reiter: fakultaet

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Darunter stehen die zwei Definitionen der Fakultät: als Produkt n! = 1 · 2 · 3 · … · n und rekursiv 1! = 1 und (n + 1)! = (n + 1) · n!.
- Unter „Iterativ“ steht die Funktion `fakultätIte` mit der Variable `fac` und einer Schleife von 2 bis `n`.
- Unter „Rekursiv“ steht die Demo. Sie zeigt elf Zeilen: die Funktion `fakultätRek(n int) int` mit `if n == 1 { return 1 } else { return n * fakultätRek(n-1) }` und `main` mit `fmt.Println(fakultätRek(4))`.
- Die Bedienung: drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“. Darunter stehen Rahmen, die Ausgabe und ein Kasten mit einem Satz zur letzten Anweisung.
- Jeder Aufruf hat einen eigenen Rahmen, zum Beispiel „Funktion fakultätRek, Aufruf mit n = 3“. Die Rahmen stehen in der Reihenfolge der Aufrufe nebeneinander. Ein gestrichelter Rahmen gehört zu einem Aufruf, der beendet ist. Der Kasten „Rückgabe“ zeigt den Wert, den ein Aufruf zurückgibt.
- Der Ablauf hat neun Schritte: Aufruf mit 4. Dann rufen die Aufrufe mit 4, 3 und 2 je den nächsten auf und warten. Der Aufruf mit 1 gibt 1 zurück. Dann rechnet der Aufruf mit 2 das Produkt 2 * 1 = 2, der Aufruf mit 3 rechnet 3 * 2 = 6, der Aufruf mit 4 rechnet 4 * 6 = 24. `main` gibt 24 aus. Die Funktion läuft viermal.
- Unter „Was gegen dieses Beispiel spricht“ stehen zwei Punkte aus dem Kurstext: Die rekursive Fassung spart keinen Speicher, weil für jeden wartenden Aufruf ein Wert von `n` gespeichert bleibt. In der Praxis nimmt man die Fakultät aus einer Tabelle.
- 20! ist 2432902008176640000 und passt noch in `int64`. 21! passt nicht mehr. Im Kurstext steht dazu nur, 20! sei in der Größenordnung von 10 hoch 18.
- `fakultätRek(0)` endet nicht: Die Bedingung `n == 1` wird nie erfüllt, das Argument wird immer kleiner. Die Funktion ist im Kurstext für n ≥ 1 gedacht.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Argument 4 lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.5.1
- 7.5
- 5.5.3.3

## Schritt 2: Endrekursion

Reiter: endrekursion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt die zwei Bedingungen der Endrekursion, englisch tail recursion: Ein Aufruf ruft die Funktion höchstens einmal rekursiv auf, und dieser Aufruf ist die letzte Operation.
- Die Demo zeigt 15 Zeilen: `fakultätEnd(n int) int` mit `return iter(n, 1)`, die Hilfsfunktion `iter(k, e int) int` mit `if k == 1 { return e } else { return iter(k-1, e*k) }` und `main` mit `fmt.Println(fakultätEnd(4))`.
- Der Ablauf hat elf Schritte: `main` ruft `fakultätEnd(4)` auf. Dann folgen die Aufrufe `iter(4, 1)`, `iter(3, 4)`, `iter(2, 12)` und `iter(1, 24)`. Der Aufruf mit `k` gleich 1 gibt `e` zurück, also 24. Die drei wartenden Aufrufe geben 24 unverändert weiter. `fakultätEnd` gibt 24 zurück, `main` gibt 24 aus.
- Unter „Der Unterschied zu Schritt 1“ steht: `fakultätRek` rechnet auf dem Rückweg, `iter` rechnet auf dem Hinweg.
- Das `else` nach dem `return` ist überflüssig und steht im Kurstext der Klarheit wegen.
- Im Kurstext steht: Bei Endrekursion muss kein Variablenwert gespeichert werden. Ob ein Compiler das ausnutzt, steht dort nicht. Der Go-Compiler tut es zurzeit nicht, auch ein endrekursiver Aufruf belegt in Go Speicher für jeden Aufruf. Die Demo zeigt deshalb für jeden Aufruf von `iter` einen Rahmen.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Argument 4 lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.5.2

## Schritt 3: Fibonacci-Zahlen

Reiter: fibonacci

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Definition: Fibo(1) = 1, Fibo(2) = 1 und Fibo(n + 2) = Fibo(n) + Fibo(n + 1). Die Folge beginnt mit 1, 1, 2, 3, 5, 8, 13, 21, 34, 55.
- Unter „Iterativ“ steht `fibonacciIter` mit den Variablen `vorletzte` und `letzte`. Die Zuweisung `vorletzte = letzte - vorletzte` gewinnt die alte letzte Zahl als Differenz zurück.
- Unter „Rekursiv“ steht `fibonacciRek` mit `return fibonacciRek(n-1) + fibonacciRek(n-2)`.
- Eine Liste unterscheidet lineare Rekursion mit höchstens einem rekursiven Aufruf und kaskadenartige Rekursion mit mehreren. Bei kaskadenartiger Rekursion wächst die Zahl der Aufrufe exponentiell mit der Rekursionstiefe.
- Die Demo hat zehn Knöpfe „n = 1“ bis „n = 10“. Ein Kasten nennt den Wert von `fibonacciRek(n)` und die Zahl aller Aufrufe. Eine Tabelle nennt zu jedem Argument, wie oft die Funktion damit läuft. Bis n = 6 zeigt die Seite darunter den Baum der Aufrufe als eingerückte Zeilen.
- Die Zahl aller Aufrufe: n = 1 und 2: 1. n = 3: 3. n = 4: 5. n = 5: 9. n = 6: 15. n = 7: 25. n = 8: 41. n = 9: 67. n = 10: 109.
- Beispiel n = 5: Wert 5, 9 Aufrufe. Mit dem Argument 4 läuft die Funktion einmal, mit 3 zweimal, mit 2 dreimal, mit 1 zweimal.
- Beispiel n = 9, das Beispiel des Kurstexts: Wert 34, 67 Aufrufe. Mit dem Argument 7 läuft die Funktion zweimal, mit 6 dreimal, mit 5 fünfmal.
- Die Seite sagt: Die rekursive Fassung ist elegant geschrieben und nicht effizient, weil sie das Gleiche mehrmals rechnet.

### Handgriffe

- Ein Argument von 1 bis 10 wählen.
- Die „Auflösung“ aufklappen.
- Andere Werte lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.5.3

## Schritt 4: Türme von Hanoi

Reiter: hanoi

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt die Aufgabe und die zwei Regeln: Ein Zug bewegt die oberste Scheibe einer Stange. Nie liegt eine größere Scheibe auf einer kleineren.
- Unter „Die Zerlegung“ stehen die drei Teile: n − 1 Scheiben zur dritten Stange, die größte Scheibe zur Zielstange, n − 1 Scheiben von der dritten Stange zur Zielstange.
- Ein Codeblock zeigt die Funktion `move(start, ziel, n int)`. Die Stangen heißen 1, 2 und 3. `6 - start - ziel` ist die Nummer der dritten Stange.
- Die Demo hat vier Knöpfe für 1 bis 4 Scheiben und drei Knöpfe „nächster Zug“, „einen Zug zurück“ und „von vorn“. Sie führt `move(1, 3, n)` aus.
- Zu sehen sind drei Stangen. Die Scheiben sind blaue Balken mit einer Nummer, 1 ist die kleinste. Die Scheibe, die der letzte Zug bewegt hat, ist braun.
- Darunter stehen die Ausgabe des Programms, zum Beispiel „Zug von 1 nach 3“, und ein Kasten mit einem Satz zum letzten Zug. Der Satz nennt den Aufruf, der die Zeile ausgibt: Die Scheibe mit der Nummer k bewegt der Aufruf `move(start, ziel, k)`.
- Die Zahl der Züge: 1 Scheibe 1 Zug, 2 Scheiben 3 Züge, 3 Scheiben 7 Züge, 4 Scheiben 15 Züge. Allgemein sind es 2 hoch n minus 1. Diese Formel steht nicht im Kurstext.
- Die Ausgabe für 3 Scheiben: von 1 nach 3, von 1 nach 2, von 3 nach 2, von 1 nach 3, von 2 nach 1, von 2 nach 3, von 1 nach 3.
- Unter „Indirekte Rekursion“ steht die Fassung des Kurstexts mit sechs Funktionen `moveAB`, `moveAC`, `moveBA`, `moveBC`, `moveCA`, `moveCB`. Gezeigt ist `moveAB`. `moveAB` ruft `moveAC` auf, und `moveAC` ruft `moveAB` auf.
- Die Funktion `move` mit den Stangen als Parametern ist direkt rekursiv.

### Handgriffe

- 1 bis 4 Scheiben wählen. Das setzt die Demo auf den Anfang zurück.
- „nächster Zug“, „einen Zug zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Der Student kann keine Scheibe selbst bewegen.

### Abschnitte des Kurstexts

- 7.5.4

## Schritt 5: Terminierung

Reiter: terminierung

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt zwei Voraussetzungen: Für mindestens einen Wert gibt es keinen rekursiven Aufruf mehr. Jeder rekursive Aufruf bekommt ein Argument, das kleiner ist als das eigene.
- Ein Codeblock zeigt die Funktion `zahlenfolge(n int)`: Sie gibt `n` aus. Bei 1 gibt sie „Ende“ aus. Bei geradem `n` ruft sie sich mit `n / 2` auf, bei ungeradem `n` mit `n + 1`.
- Die Abstiegsfunktion f_A: f_A(n) = n für gerades n, f_A(n) = n + 2 für ungerades n ab 3, f_A(1) = 1.
- Zu zeigen ist: Ruft die Funktion mit dem Argument n sich selbst mit dem Argument m auf, dann ist f_A(n) größer als f_A(m). Eine Folge natürlicher Zahlen kann nicht endlos kleiner werden, also endet die Funktion.
- Die Demo hat sechs Knöpfe mit den Anfangswerten 9, 22, 7, 27, 16 und 1. Darunter stehen die Ausgabe und eine Tabelle mit drei Zeilen: Nummer des Aufrufs, Argument n, Wert der Abstiegsfunktion.
- Die Folgen: 9, 10, 5, 6, 3, 4, 2, 1. Und 22, 11, 12, 6, 3, 4, 2, 1. Und 7, 8, 4, 2, 1. Und 27, 28, 14, 7, 8, 4, 2, 1. Und 16, 8, 4, 2, 1. Und 1.
- Die Werte der Abstiegsfunktion für den Anfang 9: 11, 10, 7, 6, 5, 4, 2, 1.
- Im Kurstext wird der Abstieg für vier Fälle nachgerechnet. Die Seite nennt die vier Fälle und rechnet sie nicht vor.
- Unter „Was geschieht, wenn eine Rekursion nicht endet“ steht die Funktion `ohneEnde` ohne Zweig für das Ende. Go bricht ab mit `runtime: goroutine stack exceeds 1000000000-byte limit` und `fatal error: stack overflow`. Das steht nicht im Kurstext.

### Handgriffe

- Einen der sechs Anfangswerte wählen.
- Die „Auflösung“ aufklappen.
- Andere Werte lassen sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.5.5

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Iterativ und rekursiv, 7.5, Seite 185; 7.5.1, Seiten 185–187: Die zwei Funktionen zur Fakultät stammen aus dem Kurstext. Die Funktion main ist eine Ergänzung. Mit Go 1.26 nachgerechnet: 20! passt noch in int64, 21! nicht mehr.
  - 2 Endrekursion, 7.5.2, Seiten 187–188: Die zwei Funktionen stammen aus dem Kurstext. Die Funktion main ist eine Ergänzung.
  - 3 Fibonacci-Zahlen, 7.5.3, Seiten 188–189: Die zwei Funktionen stammen aus dem Kurstext. Die Zahlen der Aufrufe und der Baum stehen nicht im Kurstext, sie sind mit Go 1.26 gezählt.
  - 4 Türme von Hanoi, 7.5.4, Seiten 189–191: Die Funktionen und die Ausgabe für vier Scheiben stammen aus dem Kurstext. Die Zahl der Züge, 2^n − 1, steht nicht im Kurstext.
  - 5 Terminierung, 7.5.5, Seiten 191–193: Die Funktion zahlenfolge, die Abstiegsfunktion und die Folgen für 9 und 22 stammen aus dem Kurstext. Die Funktion ohneEnde und die Meldung von Go sind eine Ergänzung, geprüft mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.6, Seiten 15–17, Rekursion: Ein Unterprogramm ruft sich selbst auf, ohne Go. Lektion: Unterprogramme und Funktionen, Schritt 4.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Warum 21! nicht mehr in int64 passt. Lektion: Überlauf bei Ganzzahlen in Go.
  - Abschnitt 8.4.1, Seiten 221–225, binäre Suche: Listing 8.1 sucht rekursiv. Lektion: Binäre Suche in einem Verbund-Array.
  - Abschnitt 9.3, Seite 237, binäre Bäume: Suchen und Einfügen im Baum sind rekursiv. Lektion: Binäre Suchbäume in Go.
- Unter „Wie es weitergeht“ stehen vier Punkte: die Programmierübung „Bubble Sort rekursiv“, der Exkurs „Laufzeit und Datenmenge“, der Ausblick auf Kapitel 9 und das Quiz zu Kapitel 7.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.6
- 4.3.8
- 7.5
- 7.5.1
- 7.5.2
- 7.5.3
- 7.5.4
- 7.5.5
- 8.4.1
- 9.3
