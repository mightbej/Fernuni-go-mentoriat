# Exkurs: Wie die Laufzeit mit der Datenmenge wächst

Die Seite ist ein Exkurs. Sie zeigt, wie die Zahl der Schritte eines Algorithmus wächst, wenn das Array größer wird, und erklärt die Schreibweise O(…). Vier Algorithmen laufen auf der Seite wirklich und zählen ihre Vergleiche. Die Seite knüpft an Abschnitt 8.4 des Kurstexts an, Seiten 221–226.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite zählt Schritte und keine Sekunden. Ein Schritt ist ein Vergleich zweier Werte, zum Beispiel `a[i] > a[i+1]`. Die Datenmenge ist die Zahl der Elemente im Array und heißt n.
- Die fünf Kurven und ihre Farben: O(1) grün, konstant. O(log n) blaugrün, logarithmisch. O(n) braun, linear. O(n log n) violett. O(n²) rot, quadratisch.
- Bei quadratischer Laufzeit wächst die Zahl der Schritte mit dem Quadrat von n: Verdoppelt sich n, braucht der Algorithmus etwa viermal so viele Schritte. Wird n zehnmal so groß, sind es etwa hundertmal so viele. Bei linearer Laufzeit verdoppelt sich die Zahl der Schritte mit n, bei logarithmischer kommt ein Schritt dazu.
- Die rote Kurve ist nach der Formel n · (n − 1) / 2 gezeichnet, der Zahl der Vergleiche von Bubble Sort, und nicht nach n².
- Im Kurstext steht die Schreibweise O(…) in Abschnitt 8.4 nicht. Dort stehen die Wörter „logarithmisch“ in Abschnitt 8.4.1 auf Seite 222 und „quadratisch“ in Abschnitt 8.4.2 auf Seite 225.
- Die Kurve O(n log n) gehört zu schnellen Sortierverfahren wie Mergesort. Sie kommen im Kurs nicht vor, auf der Seite läuft dazu kein Programm.
- Bubble Sort läuft auf der Seite ohne vorzeitigen Abbruch, wie in Listing 8.2 des Kurstexts.
- Die Schritte hießen früher Bereiche: „Grundlagen“, „Ausprobieren“, „Vertiefung“ und „Faustregeln“.
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
- Die Karte „Wo die Algorithmen erklärt sind“ verlinkt die Programmierübungen zu Bubble Sort und Selection Sort, den Exkurs „Das Kopier-Problem“, die Lektionen „Bubblesort auf einem Verbund-Array“, „Binäre Suche in einem Verbund-Array“, „Binäre Suchbäume in Go“ und „Verkettete Listen in Go“ und das Go-Programm zur binären Suche.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.1
- 8.4.2
