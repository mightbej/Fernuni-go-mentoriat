# Exkurs: Von der Problemstellung zum Go-Programm

Die Seite zeigt an einem Zahlenratespiel, wie ein Programm entsteht: Aus einer Idee wird in sechs Stufen eine immer genauere Spezifikation. Auf den Stufen 1 bis 5 ist die Spezifikation ein Text, den der Student einem Sprachmodell gibt. Auf Stufe 6 ist sie ein Go-Programm. Die Seite ist ein Exkurs des Mentoriats zu Kapitel 3 und steht nicht im Kurstext. Der Mentor zeigt sie in der Sitzung zu Kapitel 3. Das Go-Programm in Stufe 6 braucht die Anweisungen aus Kapitel 5, die Programmierübung „Zahlenraten mit Statistik“ zu Kapitel 5 baut es Schritt für Schritt auf.

Die Seite hat acht Reiter: „Worum es geht“, sechs Stufen und „Bezug zum Kurstext“. Sichtbar ist immer nur einer.

Für alle Stufen gilt:

- Die Seite nennt das Sprachmodell die „Entwicklungsabteilung“. Der Student gibt ihr die Spezifikation und spielt, was zurückkommt.
- Das Spiel wird nicht in der Spalte des Kurslotsen gespielt. Der Student kopiert die Spezifikation mit dem Knopf „Spezifikation kopieren“ und fügt sie in einem eigenen, neuen Gespräch mit einem Sprachmodell ein. Die Spalte des Kurslotsen erklärt die Seite.
- Der Kurslotse leitet deshalb kein Zahlenratespiel und denkt sich keine Zahl aus, auch wenn der Student die Spezifikation in die Spalte einfügt.
- Jede Stufe ab Stufe 2 hat drei farbige Kästen. Gelb: „Was uns am Spiel davor nicht gefallen hat“. Grün: „Was wir festlegen“. Blau: „Ausprobieren: Ist es das Spiel, das wir wollen?“ mit Fragen zum Beobachten.
- Zwischen dem grünen und dem blauen Kasten steht die Spezifikation in einem dunklen Kasten. Auf den Stufen 2 bis 4 ist grün unterlegt, was auf dieser Stufe neu dazugekommen ist.
- Der Gedanke der Seite: Alles, was die Spezifikation nicht festlegt, entscheidet die Entwicklungsabteilung. Auf der letzten Stufe bleibt nichts mehr offen.
- Die Fragen in den blauen Kästen sind Aufgaben für den Studenten. Die Seite nennt die Antworten nicht bei den Fragen. Sie stehen im gelben Kasten der nächsten Stufe.
- Wie das Sprachmodell des Studenten auf eine Spezifikation antwortet, steht nicht auf der Seite. Es fällt bei jedem Gespräch anders aus.
- Auf der Seite läuft kein Go.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen sind zwei Karten. Die erste beschreibt das Vorgehen in vier Punkten: aufschreiben, was wir wollen, die Spezifikation in die Entwicklungsabteilung geben, spielen und prüfen, das Fehlende auf der nächsten Stufe festlegen.
- Die zweite Karte „Die sechs Stufen“ ist eine Tabelle mit drei Spalten: Stufe, was am Spiel davor nicht gefallen hat, was die Stufe festlegt.
- Die sechs Stufen: 1 Spielidee, 2 Zahlenbereich 1 bis 100, 3 höchstens 7 Versuche, 4 mehrere Runden und eine Wertung, 5 jeder Schritt des Ablaufs, 6 alles in Go.
- Unter der Tabelle steht, dass das Spiel die Anweisungen aus Kapitel 5 des Kurstexts braucht und dass der letzte Reiter jede Anweisung einem Abschnitt zuordnet.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Stufe 1“ zu einer anderen Stufe wechseln.
- In dieser Stufe lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 5.5.3
- 2.1

## Stufe 1: Die Spielidee

Reiter: spielidee

### Aufbau der Darstellung

- Der grüne Kasten legt nur die Idee fest: Einer denkt sich eine Zahl aus, der andere rät und erfährt nach jedem Tipp die Richtung.
- Der dunkle Kasten zeigt die Spezifikation: ein Absatz, der mit „Wir spielen Zahlenraten.“ beginnt.
- Der blaue Kasten fragt: Aus welchem Bereich stammte die Zahl? War der Bereich bei allen Gruppen derselbe? Mit welchem Tipp habt ihr begonnen?
- Offen bleibt auf dieser Stufe der Zahlenbereich. Das ist der Mangel, den Stufe 2 behebt.

### Handgriffe

- „Spezifikation kopieren“ legt die Spezifikation in die Zwischenablage. Der Student fügt sie in ein neues Gespräch mit einem Sprachmodell ein.
- Die Spezifikation lässt sich auf der Seite nicht ändern.

### Abschnitte des Kurstexts

- 5.5.3

## Stufe 2: Der Zahlenbereich

Reiter: zahlenbereich

### Aufbau der Darstellung

- Der gelbe Kasten nennt den Mangel von Stufe 1: Der Bereich stand nirgends, einmal lag die Zahl zwischen 1 und 10, einmal zwischen 1 und 10000.
- Der grüne Kasten legt fest: eine ganze Zahl von 1 bis 100, und der Spieler erfährt den Bereich vor dem Raten.
- Die Spezifikation hat zwei Absätze. Der zweite ist grün unterlegt und neu.
- Der blaue Kasten fragt: Wie viele Tipps habt ihr gebraucht? Was geschieht, wenn man der Reihe nach 1, 2, 3 rät? Kann man dieses Spiel verlieren?
- Offen bleibt die Zahl der Versuche. Ohne Grenze kann man nicht verlieren.

### Handgriffe

- „Spezifikation kopieren“ legt die Spezifikation in die Zwischenablage.
- Die Spezifikation lässt sich auf der Seite nicht ändern.

### Abschnitte des Kurstexts

- 5.5.3

## Stufe 3: Die Zahl der Versuche

Reiter: versuche

### Aufbau der Darstellung

- Der gelbe Kasten nennt den Mangel von Stufe 2: Das Spiel ist langweilig, man kann nicht verlieren. Wer der Reihe nach rät, braucht im Mittel etwa 50 Tipps.
- Der Teil „Die Überlegung“ erklärt das Halbieren: Wer immer die Mitte des möglichen Bereichs tippt, streicht mit jeder Antwort die Hälfte der Zahlen.
- Darunter steht ein Eingabefeld „Zahlen von 1 bis“ mit dem Wert 100 und rechts davon das Ergebnis, zum Beispiel „Mit Halbieren reichen 7 Versuche.“
- Unter dem Feld steht eine Tabelle mit zwei Spalten: der Tipp und die Zahl der Zahlen, die vor diesem Tipp noch möglich sind, wenn es jedes Mal schlecht läuft. Für 100 lautet die Reihe 100, 50, 25, 12, 6, 3, 1.
- Die Tabelle rechnet so: Vor jedem Tipp halbiert sich die Zahl der möglichen Zahlen, Reste fallen weg. Bei 1 trifft der Tipp sicher.
- Unter der Tabelle steht ein Vergleich: Wer der Reihe nach rät, braucht im Mittel 50,5 Tipps und im schlimmsten Fall 100.
- 7 Versuche reichen für jeden Bereich bis 127. Ab 128 Zahlen braucht man einen achten Versuch.
- Der grüne Kasten legt fest: höchstens 7 Versuche. Nach dem siebten Fehlversuch ist die Runde verloren, und der Spieler erfährt die Zahl.
- Die Spezifikation hat drei Absätze. Der dritte ist grün unterlegt und neu.
- Das Halbieren heißt in der Informatik binäre Suche. Im Kurstext kommt sie in Abschnitt 8.4.1 bei der Suche in einem sortierten Array vor. Ein Absatz unter der Tabelle sagt das und verlinkt die Lektion „Binäre Suche in einem Verbund-Array“. Die Lektion gehört zu Kapitel 8 und setzt Funktionen, Zeiger und Verbünde voraus.

### Handgriffe

- Im Feld „Zahlen von 1 bis“ eine ganze Zahl von 2 bis 1000000 eingeben. Tabelle und Ergebnis ändern sich sofort.
- „Spezifikation kopieren“ legt die Spezifikation in die Zwischenablage.
- Die Spezifikation lässt sich auf der Seite nicht ändern. In der Spezifikation bleibt es bei 1 bis 100 und 7 Versuchen, auch wenn im Feld eine andere Zahl steht.

### Abschnitte des Kurstexts

- 5.5.3

## Stufe 4: Runden und Wertung

Reiter: runden

### Aufbau der Darstellung

- Der gelbe Kasten nennt den Mangel von Stufe 3: Nach einer Runde ist Schluss, und eine einzelne Runde sagt wenig.
- Der grüne Kasten legt fest: mehrere Runden, drei Zähler und eine Wertung. Gezählt werden die gewonnenen Runden, die verlorenen Runden und die Versuche aus den gewonnenen Runden.
- Eine Tabelle im grünen Kasten nennt die Wertung nach dem Durchschnitt der Versuche in gewonnenen Runden: höchstens 3,0 Meister, höchstens 4,0 Experte, höchstens 5,5 Fortgeschritten, darüber Anfänger.
- Die Spezifikation hat sechs Absätze. Die letzten drei sind grün unterlegt und neu.
- Der blaue Kasten fragt: Stimmt der Durchschnitt, wenn man ihn selbst nachrechnet? Stimmen die Zähler nach einer verlorenen Runde? Sehen die Meldungen bei der Nachbargruppe gleich aus?
- Offen bleibt die Reihenfolge der Schritte und der Wortlaut der Meldungen.

### Handgriffe

- „Spezifikation kopieren“ legt die Spezifikation in die Zwischenablage.
- Die Spezifikation lässt sich auf der Seite nicht ändern.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.2

## Stufe 5: Der Ablauf

Reiter: ablauf

### Aufbau der Darstellung

- Der gelbe Kasten nennt den Mangel von Stufe 4: Die Regeln stimmen, aber das Spiel sieht bei jeder Gruppe anders aus. Die Reihenfolge haben die Entwickler entschieden.
- Der grüne Kasten legt fest: Jeder Schritt steht da, in der Reihenfolge, in der er geschieht. Jeder Wert, den sich das Spiel merken muss, bekommt einen Namen.
- Die Spezifikation ist auf dieser Stufe neu geschrieben, als Ablauf mit nummerierten Schritten.
- Der Ablauf nennt sechs Werte. Drei gelten über alle Runden: `anzahlSiege`, `anzahlNiederlagen`, `gesamtVersuche`. Drei beginnen in jeder Runde neu: `zielZahl`, `versuchsZaehler`, `gewonnen`.
- Der Ablauf je Runde hat sieben Schritte: Zufallszahl wählen, Zähler und `gewonnen` zurücksetzen, die Schleife der Tipps mit den Teilschritten a bis e, die verlorene Runde, die Ausgabe der Zähler, Durchschnitt und Wertung, die Frage „Nochmal spielen? (j/n)“.
- In Schritt 3 stecken die Bausteine eines Programms: „Solange … wiederhole“ ist eine Schleife, „Ist der Tipp gleich …“ ist eine Verzweigung, „Verlasse diese Schleife sofort“ ist ein vorzeitiges Ende der Schleife.
- Der blaue Kasten fragt: Welche Schritte wiederholen sich? Welche hängen von einer Bedingung ab? Was ist immer noch nicht festgelegt?
- Der Ablauf ist noch kein Go. Er ist Text für ein Sprachmodell.

### Handgriffe

- „Spezifikation kopieren“ legt den Ablauf in die Zwischenablage.
- Der Ablauf lässt sich auf der Seite nicht ändern.

### Abschnitte des Kurstexts

- 2.1
- 5.5.3
- 5.5.1

## Stufe 6: Das Go-Programm

Reiter: programm

### Aufbau der Darstellung

- Der gelbe Kasten nennt den Mangel von Stufe 5: Das Spiel läuft nur, solange das Sprachmodell mitspielt, und das Sprachmodell macht Fehler. Das Beispiel: Auf den Tipp 35 kam „größer als 35“, auf den Tipp 37 kam „Treffer, meine Zahl ist 36“.
- Der grüne Kasten legt fest: Der Ablauf aus Stufe 5 wird in Go geschrieben. Der Go-Compiler entscheidet nichts für uns.
- Die Tabelle „Vom Ablauf zum Programm“ hat zwei Spalten: Schritt im Ablauf und Anweisung in Go. Sie hat neun Zeilen. Die Abschnitte des Kurstexts dazu stehen im Reiter „Bezug zum Kurstext“.
- Die Zuordnung: Grenzen zu `const ( MAX_VERSUCHE = 7 … )`. Drei Werte über alle Runden zu `var ( … )` außerhalb von `main`. Wiederholung je Runde zu `for { … }`. Zufallszahl zu `zielZahl := rand.Intn(…)`. Schleife der Tipps zu `for versuchsZaehler < MAX_VERSUCHE { … }`. Vergleich zu `if … else if … else`. Verlassen der Schleife zu `break`. Durchschnitt zu `float64(gesamtVersuche) / float64(anzahlSiege)`. Wertung zu `switch { case durchschnitt <= 3.0: … }`.
- Diese Stufe hat keine Spezifikation zum Kopieren. Das ganze Go-Programm steht auf der Seite „Programmierübung: Zahlenraten mit Statistik“, der Link steht unter der Tabelle.
- Die zweite Karte „Was wir gesehen haben“ fasst zusammen: Auf jeder Stufe fiel eine Entscheidung, die vorher offen war. Das Go-Programm ist die Spezifikation, die nichts mehr offen lässt. Ein Satz verlinkt die Lektion „Binäre Suche in einem Verbund-Array“, dort kommt das Halbieren aus Stufe 3 wieder vor.
- Am Ende steht ein Link zur Übung „Sortieren mit Bubble Sort“ in Kapitel 6.
- Am Ende des Schritts unter „Nächste Schritte“ steht ein Link auf die Programmierübung „Zahlenraten mit Statistik“.

### Handgriffe

- In dieser Stufe lässt sich nichts einstellen.
- Den Links zur Programmierübung und zur Übung in Kapitel 6 folgen. Beide sind eigene HTML-Seiten.
- Mit „Weiter zum Bezug zum Kurstext“ zum letzten Reiter wechseln.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.1
- 5.5.2
- 5.2

## Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Anweisung des Go-Programms aus Stufe 6 eine Zeile, zusammen neun Zeilen. Die Spalten: Anweisung in Go, was sie im Spiel tut, Abschnitt im Kurstext mit gedruckten Seiten, Lektion dazu mit Link.
- Die Zuordnung:
  - `const ( … )`: Abschnitt 5.2, Seiten 124–126, Lektion „Konstanten in Go“.
  - `var ( … )` außerhalb von `main`: Abschnitt 5.1.3, Seiten 119–120, Lektion „Variablendeklaration in Go“, Schritt 3.
  - `for { … }`: Abschnitt 5.5.3.2, Seiten 142–146, Expertenwissen 5.16, Lektion „for-Anweisung mit for-Klausel“, Schritt 4.
  - `zielZahl := …`: Abschnitt 5.1.2, Seiten 116–118, Lektion „Variablendeklaration in Go“, Schritt 2.
  - `for versuchsZaehler < MAX_VERSUCHE { … }`: Abschnitt 5.5.3.1, Seiten 138–142, Lektion „for-Anweisung mit Bedingung“.
  - `if … else if … else`: Abschnitt 5.5.1, Seiten 129–132, Lektion „if-Anweisung in Go“.
  - `break`: Abschnitt 5.5.3.3, Seiten 146–149, Lektion „break-Anweisung in Go“.
  - `float64(…) / float64(…)`: Abschnitt 4.7, Seiten 109–111, Lektion „Typkonvertierung“.
  - `switch { … }`: Abschnitt 5.5.2, Seiten 133–137, Lektion „switch-Anweisung in Go“.
- Unter „Was der Exkurs sonst aus dem Kurstext aufnimmt“ stehen zwei Abschnitte:
  - Abschnitt 2.7, Seiten 17–19, Spezifikation: Sie legt fest, welche Eingaben gültig sind und welche Ausgabe zu welcher Eingabe richtig ist. Lektion „Höhere und niedrigere Programmiersprachen“, Schritt 4.
  - Abschnitt 8.4.1, Seiten 221–225, binäre Suche: das Halbieren aus Stufe 3. Lektion „Binäre Suche in einem Verbund-Array“.
- Unter „Was über den Kurstext hinausgeht“ steht: Die Stufen 1 bis 5 und das Sprachmodell als Entwicklungsabteilung stehen nicht im Kurstext. Das Paket `math/rand` mit `rand.Intn` kommt im Kurstext nicht vor.

### Handgriffe

- Mit den Reitern oben, mit „Zurück“ oder mit „Zum Anfang“ zu einer anderen Stufe wechseln.
- In diesem Reiter lässt sich nichts einstellen.

### Abschnitte des Kurstexts

- 5.2
- 5.1.3
- 5.5.3.2
- 5.1.2
- 5.5.3.1
- 5.5.1
- 5.5.3.3
- 4.7
- 5.5.2
- 2.7
- 8.4.1
