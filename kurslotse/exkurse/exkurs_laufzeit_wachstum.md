# Exkurs: Wie die Laufzeit mit der Datenmenge wächst

Die Seite ist ein Exkurs. Sie zeigt, wie die Zahl der Schritte eines Algorithmus wächst, wenn das Array größer wird, und erklärt die Schreibweise O(…). Vier Algorithmen laufen auf der Seite wirklich und zählen ihre Vergleiche. Die Seite knüpft an Abschnitt 8.4 des Kurstexts an, Seiten 221–226.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite zählt Schritte und keine Sekunden. Ein Schritt ist ein Vergleich zweier Werte, zum Beispiel `a[i] > a[i+1]`. Die Datenmenge ist die Zahl der Elemente im Array und heißt n.
- Die fünf Kurven und ihre Farben: O(1) grün, konstant. O(log n) blaugrün, logarithmisch. O(n) braun, linear. O(n log n) violett. O(n²) rot, quadratisch.
- Bei quadratischer Laufzeit wächst die Zahl der Schritte mit dem Quadrat von n: Verdoppelt sich n, braucht der Algorithmus etwa viermal so viele Schritte. Wird n zehnmal so groß, sind es etwa hundertmal so viele. Bei linearer Laufzeit verdoppelt sich die Zahl der Schritte mit n, bei logarithmischer kommt ein Schritt dazu.
- Die rote Kurve ist nach der Formel n · (n − 1) / 2 gezeichnet, der Zahl der Vergleiche von Bubble Sort, und nicht nach n².
- Im Kurstext steht die Schreibweise O(…) in Abschnitt 8.4 nicht. Dort stehen die Wörter „logarithmisch“ in Abschnitt 8.4.1 auf Seite 222 und „quadratisch“ in Abschnitt 8.4.2 auf Seite 225.
- Die Kurve O(n log n) gehört zu schnellen Sortierverfahren wie Mergesort. Sie kommen im Kurs nicht vor, auf der Seite läuft dazu kein Programm.
- Bubble Sort läuft auf der Seite ohne vorzeitigen Abbruch, wie in Listing 8.2 des Kurstexts.
- „Worum es geht“ und die Schritte 1 bis 3 hießen früher Bereiche: „Grundlagen“, „Ausprobieren“, „Vertiefung“ und „Faustregeln“.
- Die Links auf die Lektionen, in denen die Algorithmen erklärt sind, stehen in Schritt 4 unter „Wo das Thema im Kurstext noch vorkommt“.
- Auf der Seite läuft kein Go. Die vier Algorithmen laufen im Browser und entsprechen dem Go-Programm `laufzeit_vergleiche_zaehlen.go` im Repository. Der Student kann keinen eigenen Code eingeben.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: grundlagen

### Aufbau der Darstellung

- Die erste Karte nennt, was die Seite zählt, und die zwei Stellen im Kurstext: binäre Suche in Abschnitt 8.4.1, Bubblesort in Abschnitt 8.4.2.
- Die zweite Karte „Fünf Kurven“ zeigt ein festes Schaubild. Nach rechts wächst das Array bis n = 30, nach oben die Zahl der Schritte bis 100. Die rote Kurve verlässt den Ausschnitt bei n = 15.
- Eine Tabelle nennt zu jeder Kurve den Namen der Laufzeit, einen Algorithmus und die Schritte bei n = 1000: 1 bei O(1), 10 bei O(log n), 1.000 bei O(n), rund 10.000 bei O(n log n), 499.500 bei O(n²).
- Ein blauer Kasten erklärt die Schreibweise: O(n²) heißt, dass ein zehnmal so großes Array etwa hundertmal so viele Schritte kostet.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts verstellen.

### Abschnitte des Kurstexts

- 8.4.1
- 8.4.2

## Schritt 1: Vergleiche zählen

Reiter: ausprobieren

### Aufbau der Darstellung

- Die Seite füllt ein Array mit n gemischten Zahlen und lässt vier Algorithmen laufen: lineare Suche, binäre Suche, Bubble Sort und Selection Sort. Gesucht wird ein Wert, der im Array fehlt. Das ist für beide Suchen der ungünstigste Fall.
- Ein Regler stellt die Größe des Arrays von 2 bis 1.000 ein, am Anfang 100. Ein zweiter Regler stellt ein, wie weit die senkrechte Achse reicht, von 20 bis 550.000 Vergleichen, am Anfang 5.500. Der Knopf „Achse an den höchsten Punkt anpassen“ wählt die Achse so, dass der Punkt von Bubble Sort gerade hineinpasst.
- Das Schaubild zeigt die fünf Kurven bis n = 1.000 und eine gestrichelte senkrechte Linie beim gewählten n. Auf der Linie liegen drei Punkte mit den gezählten Vergleichen: binäre Suche, lineare Suche und ein gemeinsamer Punkt für Bubble Sort und Selection Sort.
- Unter dem Schaubild steht, welche Punkte oberhalb des Ausschnitts liegen.
- Eine Tabelle hat sechs Zeilen: Zugriff auf `a[i]`, binäre Suche, lineare Suche, Mergesort, Bubble Sort, Selection Sort. Spalten: Kurve, Schritte nach der Formel, gezählte Vergleiche, gezählte Vertauschungen. Bei Zugriff und Mergesort steht „Auf der Seite läuft dazu kein Programm.“
- Bubble Sort und Selection Sort zählen immer gleich viele Vergleiche, n · (n − 1) / 2. Bei den Vertauschungen liegen sie weit auseinander.
- Ein gelber Kasten „Zum Ausprobieren“ nennt drei Versuche: Achse klein stellen, Achse groß stellen, n verdoppeln. Von n = 100 auf n = 200 bekommt die lineare Suche 100 Vergleiche dazu, die binäre Suche einen, Bubble Sort 14.950.
- Die Karte „So zählt das Programm“ zeigt Bubble Sort in Go mit einem Zähler in der inneren Schleife und einen Link auf das Go-Programm.

### Handgriffe

- Den Regler „Größe des Arrays“ schieben.
- Den Regler für die senkrechte Achse schieben oder „Achse an den höchsten Punkt anpassen“ wählen.
- Die Zahlen im Array kann der Student nicht ändern.

### Abschnitte des Kurstexts

- 8.4.2
- 8.4.1

## Schritt 2: Woher die Formeln kommen

Reiter: vertiefung

### Aufbau der Darstellung

- Die Karte „Woher die Formeln kommen“ leitet vier Formeln her. Bubble Sort: (n − 1) + (n − 2) + … + 1 = n · (n − 1) / 2. Selection Sort: die gleiche Summe. Lineare Suche: n. Binäre Suche: Aus 1000 Elementen werden 500, 250, 125, 62, 31, 15, 7, 3, 1, nach zehn Blicken ist nichts übrig.
- Die Karte „Was die Schreibweise O(…) weglässt“ nennt den festen Faktor und den kleineren Summanden. Eine Tabelle stellt n · (n − 1) / 2 neben n²: 45 und 100 bei n = 10, 4.950 und 10.000 bei n = 100, 499.500 und 1.000.000 bei n = 1.000.
- Die Karte „Drei Feinheiten“: Selection Sort vertauscht höchstens n − 1 mal. Bubble Sort mit Abbruch braucht für ein sortiertes Array nur n − 1 Vergleiche, das ist im Kurstext auf Seite 226 eine Übungsaufgabe. Erst ab n = 7 liegt O(n²) über O(n log n).
- Eine Demo gibt es hier nicht.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.2
- 8.4.1

## Schritt 3: Die Kurve erkennen

Reiter: faustregeln

### Aufbau der Darstellung

- Die Karte „Die Kurve am Programm erkennen“ nennt vier Faustregeln: keine Schleife über das Array ist O(1), eine Schleife O(n), zwei geschachtelte Schleifen O(n²), Halbieren des Rests in jedem Schritt O(log n).
- Die Faustregeln stehen nicht im Kurstext.
- Eine Tabelle zeigt, was bei zehnfacher Größe des Arrays geschieht, mit den Schritten bei n = 1000 und bei n = 1.000.000.
- Ein grüner Kasten rechnet in Zeit um: Bei einer Milliarde Vergleichen in der Sekunde sortiert Bubble Sort eine Million Elemente in gut acht Minuten, ein Verfahren mit O(n log n) in zwei Hundertstelsekunden.
- Eine Demo gibt es hier nicht.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Anwendung: Wie viele Vergleiche Bubble Sort braucht, zeigt die Lektion „Bubblesort auf einem Verbund-Array“.
  - Gegenstück: Bei der „binären Suche in einem Verbund-Array“ wächst die Zahl der Schritte nur logarithmisch.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.1
- 8.4.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 8.4.1, Seiten 221–222; 8.4.2, Seite 225: Aus dem Kurstext stammen zwei Wörter: „logarithmisch“ für die binäre Suche und „quadratisch“ für Bubblesort. Die Schreibweise O(…), das Schaubild und die Kurven O(1), O(n) und O(n log n) stehen nicht im Kurstext.
  - 1 Vergleiche zählen, 8.4.1, Seiten 221–225: Listing 8.1; 8.4.2, Seiten 225–226: Listing 8.2: Bubble Sort läuft mit den gleichen zwei Schleifen wie in Listing 8.2, ohne vorzeitigen Abbruch. Die binäre Suche wählt die Mitte wie in Listing 8.1. Im Kurstext ruft sich die binäre Suche selbst auf. Auf der HTML-Seite läuft sie mit einer Schleife. Das Array enthält Zahlen statt Verbünde mit Namen. Die lineare Suche, Selection Sort und das Zählen der Vergleiche stehen nicht im Kurstext.
  - 2 Woher die Formeln kommen, 8.4.1, Seiten 221–222; 8.4.2, Seiten 225–226: Aus dem Kurstext stammen die Herleitung von n · (n − 1) / 2, der Satz zu den Vertauschungen und die Übungsaufgabe zum Abbruch. Dass sich ein Array mit n Elementen nur log₂ n mal halbieren lässt, steht in Abschnitt 8.4.1. Die Formeln für Selection Sort und die lineare Suche, die Zahlenreihe 500, 250, 125 … und was die Schreibweise O(…) weglässt, stehen nicht im Kurstext.
  - 3 Die Kurve erkennen, 8.4.2, Seite 225: Aus dem Kurstext stammt nur, dass Bubblesort quadratische Laufzeit hat. Die vier Faustregeln, die Tabelle bis n = 1.000.000 und die Umrechnung in Zeit stehen nicht im Kurstext.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.1, Seiten 5–8, Schleife: Ohne Go wird gezeigt, wie WHILE und REPEAT-UNTIL Anweisungen wiederholen. Wie oft eine Schleife läuft, bestimmt die Zahl der Schritte. Lektion: Imperative Programmierung, Schritt 3.
  - Abschnitt 6.2, Seiten 156–159, ein Array durchlaufen: Eine Schleife sieht jedes Element einmal an. Das ist die Kurve O(n). Lektion: Arrays in Go, Schritt 3.
  - Abschnitt 7.2, Seiten 169–170, Kopie des Arguments: Ein Array als Argument wird ganz kopiert, und auch das kostet Schritte. Deshalb übergibt die binäre Suche in Abschnitt 8.4.1 einen Zeiger auf das Array. Lektion: Das Kopier-Problem.
  - Abschnitt 7.5.3, Seiten 188–189, kaskadenartige Rekursion: Ruft sich eine Funktion zweimal selbst auf, wächst die Zahl der Aufrufe exponentiell mit der Rekursionstiefe. Bei einem Aufruf wächst sie nur linear. Lektion: Rekursion in Go, Schritt 3.
  - Abschnitt 8.4.1, Seiten 221–225, binäre Suche: Lektion: Binäre Suche in einem Verbund-Array, Schritt 4. Das Go-Programm dazu liegt im Repository: S230_binaere_suche_demo.go.
  - Abschnitt 8.4.2, Seiten 225–226, Bubblesort: Lektion: Bubblesort auf einem Verbund-Array, Schritt 4.
  - Abschnitt 9.1 und 9.2, Seiten 227–236, Einfügen ohne Verschieben: In einem sortierten Array muss beim Einfügen im Durchschnitt die Hälfte der Elemente einen Platz weiterrücken. In einer verketteten Liste genügen zwei Zeiger-Zuweisungen. Lektion: Verkettete Listen in Go, Schritt 6.
  - Abschnitt 9.3.2, Seite 243, Suchen im binären Suchbaum: In einem balancierten Baum braucht das Suchen etwa log₂ n Schritte. Lektion: Binäre Suchbäume in Go, Schritt 4.
  - Ohne Abschnitt im Kurstext, Bubble Sort und Selection Sort an einem Array mit Zahlen: Lektionen: Programmierübung: Sortieren mit Bubble Sort und Programmierübung: Sortieren mit Selection Sort.
- Unter „Was über den Kurstext hinausgeht“ steht:
  - Die Schreibweise O(…): Im Kurstext stehen nur die Wörter „logarithmisch“ und „quadratisch“. O(1), O(n) und O(n log n) kommen im Kurstext nicht vor.
  - Lineare Suche: Eine Suche, die das Array von vorn nach hinten durchgeht, kommt im Kurstext nicht vor. Die HTML-Seite nimmt sie als Beispiel für die Kurve O(n).
  - Selection Sort: Selection Sort kommt im Kurstext nicht vor. Auf Seite 225 steht nur, dass im Vorkurs zwei einfache Sortierverfahren vorgestellt wurden.
  - Mergesort und O(n log n): Schnelle Sortierverfahren kommen im Kurs nicht vor. Auf der HTML-Seite läuft dazu kein Programm.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 6.2
- 7.2
- 7.5.3
- 8.4.1
- 8.4.2
- 9.1
- 9.2
- 9.3.2
