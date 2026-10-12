# Größter gemeinsamer Teiler: vom Struktogramm zum Go-Programm

Die Seite zeigt ein kleines Programm in vier Darstellungen: als Text, als Struktogramm, als Flussdiagramm und als Go-Programm. Das Programm berechnet den größten gemeinsamen Teiler zweier positiver ganzer Zahlen, kurz ggT. Es zieht so lange die kleinere Zahl von der größeren ab, bis beide Zahlen gleich sind.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für die Schritte 1, 2 und 4 gilt:

- Sie zeigen denselben Durchlauf des Programms, jeder in seiner Darstellung. Wer den Reiter wechselt, bleibt an derselben Stelle des Durchlaufs und bei denselben Zahlen.
- Der Durchlauf ist nachgestellt. Die Seite führt kein Go-Programm aus. Der Student kann den Code nicht ändern und nicht starten.
- Das Wort „Schritt“ kommt auf der Seite zweimal vor. Die Reiter heißen Schritt 1 bis Schritt 4. Der Zähler unter der Darstellung, zum Beispiel „3 von 24“, zählt die Schritte des Durchlaufs.
- Ein Schritt des Durchlaufs ist eines von sechs Ereignissen: die Eingabe, die Prüfung der Bedingung der Schleife, die Frage `n > m?`, die Subtraktion `n ← n − m`, die Subtraktion `m ← m − n`, die Ausgabe.
- Unter der Darstellung steht ein Feld mit vier Zeilen: Schritt, Wert von n, Wert von m, Ausgabe. Darunter steht in einem grünen Kasten ein Satz, der den letzten Schritt des Durchlaufs erklärt.
- Die Werte von n und m sind die Werte nach dem Schritt, der gerade markiert ist.

Die Schreibweise des Programms aus Abschnitt 2.5 gehört zu keiner bestimmten Programmiersprache. Der Pfeil `←` bedeutet: Die Variable links bekommt den Wert, der rechts ausgerechnet wird. `≠` bedeutet ungleich.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Bilder und einen Durchlauf gibt es hier nicht.
- In einem dunklen Kasten steht das Programm aus Abschnitt 2.5 als Text mit fünf Zeilen: Eingabe, `WHILE n ≠ m`, `IF n > m THEN n ← n − m`, `ELSE m ← m − n`, Ausgabe.
- Darunter steht, was die vier Schritte der Seite zeigen.
- Am Ende stehen die zwei Fragen aus dem Kurstext: was das Programm leistet und warum es bei positiven Zahlen immer zu einem Ende kommt.
- Die Antworten auf die zwei Fragen sind verdeckt, bis der Student „Antworten zeigen“ wählt. Der Stand der Demo sagt, ob sie verdeckt oder aufgedeckt sind.

### Handgriffe

- „Antworten zeigen“ wählen. Derselbe Knopf verdeckt die Antworten wieder.
- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 2.1
- 2.5

## Schritt 1: Das Programm als Struktogramm

Reiter: struktogramm

### Aufbau der Darstellung

- Zu sehen ist das Struktogramm: ein großes Rechteck aus Kästen, die ohne Lücke aneinanderstoßen. Man liest es von oben nach unten. Pfeile gibt es nicht.
- Ganz oben: ein Kasten über die ganze Breite mit „Eingabe n, m ∈ {1, 2, 3, …}“.
- Darunter: der Kasten der Schleife mit „while n ≠ m“. Er hat die Form eines Winkels. Er läuft oben über die ganze Breite und links als schmaler Streifen nach unten. Alles, was rechts neben dem Streifen steht, ist der Rumpf der Schleife und wird wiederholt.
- Im Rumpf oben: der Kasten der Verzweigung mit der Frage „n > m?“. Zwei schräge Linien laufen von seinen oberen Ecken zur Mitte unten. In der linken Ecke steht „ja“, in der rechten „nein“.
- Im Rumpf unten: zwei Kästen nebeneinander. Links unter „ja“ steht „n ← n − m“, rechts unter „nein“ steht „m ← m − n“. Bei jedem Durchgang durch die Schleife wird nur einer der beiden ausgeführt.
- Ganz unten: ein Kasten über die ganze Breite mit „Ausgabe n“. Er liegt außerhalb der Schleife.
- Gelb mit dickem Rand ist der Kasten, der im Durchlauf gerade an der Reihe ist. Ist der Kasten „while n ≠ m“ gelb, wird die Bedingung der Schleife geprüft.
- Nach einer Subtraktion springt die Markierung zurück auf „while n ≠ m“. So zeigt das Struktogramm die Wiederholung.
- Über dem Struktogramm stehen die Felder für n und m und die Knöpfe des Durchlaufs. Darunter stehen das Feld mit Schritt, Werten und Ausgabe und der erklärende Satz.

### Handgriffe

- In die Felder n und m zwei ganze Zahlen eingeben: jede von 1 bis 99. Voreingestellt sind 42 und 78.
- „Zahlen übernehmen“ wählen. Erst dann rechnet der Durchlauf mit den neuen Zahlen. Er beginnt dabei von vorn.
- „nächster Schritt“: Der Durchlauf geht einen Schritt weiter.
- „einen Schritt zurück“: Der Durchlauf geht einen Schritt zurück.
- „von vorn“: Der Durchlauf beginnt neu mit denselben Zahlen.
- „bis zum Ende“: Der Durchlauf springt zur Ausgabe.
- Mit den Reitern oben zu einem anderen Schritt wechseln. Der Durchlauf bleibt dabei an seiner Stelle.
- Die Zahl 0, negative Zahlen, Zahlen über 99 und Kommazahlen nimmt die Seite nicht an. Sie zeigt dann eine Meldung, und der Durchlauf bleibt bei den alten Zahlen. Eine Eingabe mit 0 lässt sich auf der Seite also nicht ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.5

## Schritt 2: Das Programm als Flussdiagramm

Reiter: flussdiagramm

### Aufbau der Darstellung

- Zu sehen ist das Flussdiagramm: einzelne Kästen, die mit Pfeilen verbunden sind. Man liest es entlang der Pfeile. Ein Rechteck ist eine Anweisung. Eine Raute ist eine Frage mit zwei Ausgängen, die mit „ja“ und „nein“ beschriftet sind.
- Oben in der Mitte: das Rechteck „Eingabe n, m ∈ {1, 2, 3, …}“. Ein Pfeil führt nach unten zur oberen Raute.
- Obere Raute: die Frage „n = m?“. Der Ausgang „ja“ führt nach rechts zum Rechteck „Ausgabe n“. Dort endet das Programm. Der Ausgang „nein“ führt nach unten zur unteren Raute.
- Untere Raute: die Frage „n > m?“. Der Ausgang „ja“ führt nach rechts und dann nach unten zum Rechteck „n ← n − m“. Der Ausgang „nein“ führt gerade nach unten zum Rechteck „m ← m − n“.
- Unten stehen also zwei Rechtecke nebeneinander: unter der unteren Raute „m ← m − n“, rechts daneben „n ← n − m“. Die Reihenfolge ist umgekehrt wie im Struktogramm.
- Von beiden Rechtecken führt eine Linie nach unten. Die Linien laufen zusammen, führen am linken Rand nach oben und münden mit einem Pfeil zwischen dem Rechteck „Eingabe“ und der oberen Raute. Dieser Weg zurück nach oben ist die Schleife. Einen eigenen Kasten hat die Schleife hier nicht.
- Die obere Raute fragt „n = m?“. Im Struktogramm steht dieselbe Bedingung umgekehrt als „n ≠ m“. Der Ausgang „nein“ der Raute entspricht der erfüllten Bedingung „n ≠ m“.
- Gelb mit dickem Rand ist das Rechteck oder die Raute, die im Durchlauf gerade an der Reihe ist. Die Pfeile werden nicht markiert.
- Über dem Flussdiagramm stehen die Felder für n und m und die Knöpfe des Durchlaufs. Darunter stehen das Feld mit Schritt, Werten und Ausgabe und der erklärende Satz.
- Unter dem Flussdiagramm steht ein Link auf die Seite „Imperative Programmierung, Schritt 3“. Sie zeigt, wie Blöcke den Sprung ersetzen.

### Handgriffe

- In die Felder n und m zwei ganze Zahlen eingeben: jede von 1 bis 99. Voreingestellt sind 42 und 78.
- „Zahlen übernehmen“ wählen. Erst dann rechnet der Durchlauf mit den neuen Zahlen. Er beginnt dabei von vorn.
- „nächster Schritt“: Der Durchlauf geht einen Schritt weiter.
- „einen Schritt zurück“: Der Durchlauf geht einen Schritt zurück.
- „von vorn“: Der Durchlauf beginnt neu mit denselben Zahlen.
- „bis zum Ende“: Der Durchlauf springt zur Ausgabe.
- Mit den Reitern oben zu einem anderen Schritt wechseln. Der Durchlauf bleibt dabei an seiner Stelle.
- Die Zahl 0, negative Zahlen, Zahlen über 99 und Kommazahlen nimmt die Seite nicht an. Sie zeigt dann eine Meldung, und der Durchlauf bleibt bei den alten Zahlen. Eine Eingabe mit 0 lässt sich auf der Seite also nicht ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.5

## Schritt 3: Von der Darstellung zum Go-Code

Reiter: uebergang

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle, einem Stück Go-Code und einer Liste. Einen Durchlauf gibt es hier nicht.
- Die Tabelle hat vier Zeilen, je eine für `WHILE n ≠ m`, `IF n > m THEN`, `n ← n − m` und `ELSE m ← m − n`. Jede Tabellenzeile nennt die Schreibweise aus Abschnitt 2.5, dieselbe Zeile in Go, was sich in Go ändert, und den Abschnitt im Kurstext mit einem Link auf die HTML-Seite dazu.
- Unter der Tabelle steht in einem dunklen Kasten der Kern in Go: die Schleife `for n != m` mit `if n > m`, `n -= m`, `else` und `m -= n`. Das sind die Zeilen 7 bis 13 des Programms in Schritt 4.
- Die Liste darunter nennt, was ein vollständiges Go-Programm zum Kern dazu braucht: den Rahmen mit `package main` und `main`, die Deklaration `var n, m int`, Eingabe und Ausgabe mit `fmt.Scan` und `fmt.Printf`, die Prüfung der Eingabe und die Funktion `ggT` mit `return n`.
- Go hat kein Schlüsselwort `while`. Die Schleife mit Bedingung schreibt Go mit `for`.

### Handgriffe

- Den Links in der Tabelle und in der Liste zu einer anderen HTML-Seite folgen.
- Mit den Reitern oben zu einem anderen Schritt wechseln. Schritt 4 zeigt das ganze Programm im Durchlauf.

### Abschnitte des Kurstexts

- 2.5
- 5.3
- 5.5.1
- 5.5.3.1

## Schritt 4: Das Go-Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das Go-Programm `for_ggt.go` in einem dunklen Kasten, 27 Zeilen mit Zeilennummern am linken Rand. Die Programmzeilen gehen mit ihren Nummern mit.
- Zeilen 6 bis 15: die Funktion `ggT` mit der Schleife. Zeilen 17 bis 27: die Funktion `main` mit Deklaration, Eingabe, Prüfung der Eingabe und Ausgabe.
- Hell hinterlegt mit einem gelben Strich am linken Rand sind die Zeilen, die zum Schritt des Durchlaufs gehören. Der Durchlauf hat dieselben Schritte wie im Struktogramm.
- Eingabe: Die Zeilen 20 und 21 sind markiert. Die Zahlen kommen aus den Feldern n und m über dem Programm. Eine Eingabe über die Tastatur wie bei `fmt.Scan` gibt es auf der Seite nicht.
- Prüfung der Bedingung der Schleife: Zeile 7 ist markiert.
- Frage `n > m`: Zeile 8 ist markiert.
- Subtraktion: Zeile 9 mit `n -= m` oder Zeile 11 mit `m -= n` ist markiert.
- Ausgabe: Die Zeilen 14 und 23 sind markiert. Zeile 14 gibt das Ergebnis an `main` zurück, Zeile 23 gibt den Satz aus.
- Die Zeilen 18, 19 und 22 werden nie markiert. Der Durchlauf zeigt sie nicht als eigenen Schritt. Der erklärende Satz zur Eingabe nennt die Zeilen 22 und 23.
- Die Werte von n und m im Feld unter dem Programm sind während der Schleife die Werte der Variablen in der Funktion `ggT`. Die Variablen `n` und `m` in `main` behalten die eingegebenen Zahlen. Deshalb nennt die Ausgabe am Ende die eingegebenen Zahlen.
- Im Feld „Ausgabe“ steht am Ende der ganze Satz, zum Beispiel „6 ist der ggT von 42 und 78.“
- Unter dem Feld stehen drei Listenpunkte zu den Zeilen 7 bis 13, zu Zeile 22 und zu Zeile 23.
- Die `else`-Zeilen 24 und 25 mit der Meldung „unzulässig“ lassen sich auf der Seite nicht erreichen, weil die Felder nur Zahlen von 1 bis 99 annehmen.

### Handgriffe

- In die Felder n und m zwei ganze Zahlen eingeben: jede von 1 bis 99. Voreingestellt sind 42 und 78.
- „Zahlen übernehmen“ wählen. Erst dann rechnet der Durchlauf mit den neuen Zahlen. Er beginnt dabei von vorn.
- „nächster Schritt“: Der Durchlauf geht einen Schritt weiter.
- „einen Schritt zurück“: Der Durchlauf geht einen Schritt zurück.
- „von vorn“: Der Durchlauf beginnt neu mit denselben Zahlen.
- „bis zum Ende“: Der Durchlauf springt zur Ausgabe.
- Mit den Reitern oben zu einem anderen Schritt wechseln. Der Durchlauf bleibt dabei an seiner Stelle.
- Die Zahl 0, negative Zahlen, Zahlen über 99 und Kommazahlen nimmt die Seite nicht an. Sie zeigt dann eine Meldung, und der Durchlauf bleibt bei den alten Zahlen. Eine Eingabe mit 0 lässt sich auf der Seite also nicht ausprobieren, und die Meldung „unzulässig“ aus Zeile 25 erscheint nie.
- Der Student kann das Programm nicht ändern und nicht ausführen.

### Abschnitte des Kurstexts

- 5.5.3.1
- 7.2
