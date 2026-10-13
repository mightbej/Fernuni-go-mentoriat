# Speichermodell: Arrays in Go

Die Seite zeigt, wie ein Array im Hauptspeicher liegt. Ein Array hat einen Namen und viele Elemente. In der Tabelle des Rechners steht nur der Anfang, die Adresse eines Elements rechnet der Rechner aus dem Index aus. Die Seite gehört zu Kapitel 6 des Kurstexts und zu Abschnitt 6.1. Sie ist die zweite von vier HTML-Seiten mit dem gleichen Modell.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Hauptspeicher ist eine Reihe von Zellen. Jede Zelle fasst 8 Bit, also 1 Byte, und hat eine Adresse.
- Der Rechner führt eine Tabelle. In jeder Zeile stehen ein Name, ein Typ und eine Adresse. Der Wert steht nur im Hauptspeicher.
- Ein Array ist eine Reihe gleich großer Elemente, die lückenlos hintereinander liegen. Die Tabelle bekommt für das ganze Array eine einzige Zeile mit dem Namen, dem Typ und der Adresse des ersten Elements.
- Die Rechnung für ein Element: Adresse des Elements = Adresse des Arrays + Index · Zellen pro Element.
- Das Beispiel der Seite ist `var a [4]int16`. Das Array hat 4 Elemente, jedes Element ist ein `int16` mit 2 Zellen. Das Array belegt 8 Zellen und beginnt an der Adresse 100.
- Das Modell hat 32 Zellen mit den Adressen 100 bis 131. Die Adressen sind erfunden.
- Die Wörter „Zelle“ und „Tabelle des Rechners“ gehören zum Modell der Seite. Die Rechnung mit Adressen steht nicht im Kurstext. Dort steht zu Arrays im Speicher nur eine Fußnote: Die Länge eines Arrays ist nicht sein Speicherplatz.
- Das Modell vereinfacht. In Go benutzt der Compiler die Tabelle beim Übersetzen. Ein laufendes Go-Programm schlägt keinen Namen mehr nach.
- Ein Wert aus mehreren Zellen steht im Modell so, dass die vorderen Bits in der Zelle mit der kleineren Adresse stehen.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text erklärt Zelle, Adresse und Tabelle und sagt, was ein Array im Speicher ist.
- Im Typ `[4]int16` steckt, was der Rechner braucht: 4 Elemente, jedes mit 2 Zellen.
- Eine Liste nennt zwei HTML-Seiten, die dieser vorausgehen: „Speichermodell: Name, Typ und Adresse“ zu Kapitel 4 und „Arrays in Go“.
- Ein gelber Kasten sagt, dass das Modell vereinfacht.
- Was auf dieser Seite steht, steht so nicht im Kurstext. Zellen, Adressen, die Tabelle des Rechners und die Rechnung mit dem Index kommen in Kapitel 6 des Kurstexts nicht vor. In Kapitel 6 steht, was ein Array ist. In Abschnitt 6.1 steht, was ein Index ist und dass die Zählung bei 0 beginnt. Die Seite zeigt, was dabei im Speicher geschieht, als Ergänzung des Mentors.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den zwei HTML-Seiten folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demo steht in Schritt 2.

### Abschnitte des Kurstexts

- 6
- 6.1

## Schritt 1: Die Adresse eines Elements

Reiter: adresse

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text, einer Formel und einer Tabelle. Eine Demo gibt es hier nicht.
- Für `a[2]` gibt es keine Zeile in der Tabelle des Rechners. Der Rechner rechnet die Adresse aus.
- Die Formel steht in einem eigenen Kasten: Adresse des Elements = Adresse des Arrays + Index · Zellen pro Element.
- Die Tabelle zeigt die Rechnung für `var a [4]int16` ab der Adresse 100: `a[0]` liegt an 100 + 0 · 2 = 100 in den Zellen 100 und 101. `a[1]` liegt an 102 in den Zellen 102 und 103. `a[2]` liegt an 104 in den Zellen 104 und 105. `a[3]` liegt an 106 in den Zellen 106 und 107.
- Der Rechner braucht für jedes Element eine einzige Rechnung. Der Zugriff auf `a[3]` ist deshalb so schnell wie der Zugriff auf `a[0]`. Das gilt auch für ein Array mit 1000 Elementen.
- Die Zählung beginnt bei 0, weil der Index sagt, wie viele Elemente vor dem gesuchten liegen.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Tabelle ist fest.
- Mit „Weiter zu Schritt 2“ zur Demo wechseln.

### Abschnitte des Kurstexts

- 6.1

## Schritt 2: Ein Programm im Speicher

Reiter: speicher

### Aufbau der Darstellung

- Oben steht das Programm als nummerierte Liste mit sechs Zeilen: `var a [4]int16`, `a[0] = 17`, `a[2] = 1000`, `i := 3`, `a[i] = a[0] + a[2]`, `fmt.Println(a)`. Die Zeilen stehen im Rumpf von `main`.
- Die zuletzt ausgeführte Zeile steht weiß auf dunklem Grund. Die Zeilen davor sind blass.
- Darunter stehen drei Knöpfe: „nächste Zeile ausführen“, „eine Zeile zurück“ und „von vorn“.
- Links liegt der Hauptspeicher: 32 Zellen mit den Adressen 100 bis 131. In jeder Zelle steht oben klein die Adresse und darunter das Bitmuster aus acht Bits.
- Die Zellen einer Variable haben eine Farbe. Die Elemente des Arrays sind abwechselnd hell und dunkel gefärbt, so sieht man, welche zwei Zellen zu einem Element gehören.
- Eine Zelle mit blauem Rand hat die letzte Zeile gelesen. Eine Zelle mit rotem Rand hat die letzte Zeile beschrieben.
- Unter dem Hauptspeicher steht eine Zelle groß: acht Bits als Knöpfe, über jedem Bit seine Wertigkeit. Darüber steht, zu welchem Element oder welcher Variable die Zelle gehört.
- Rechts steht die „Tabelle des Rechners“ mit den Spalten Name, Typ, Adresse und Zellen. Für das Array steht dort eine Zeile: `a`, `[4]int16`, Adresse 100, 8 Zellen.
- Darunter steht der Kasten „Was der Rechner in dieser Zeile tut“ mit den Handlungen des Rechners als nummerierte Liste, auch mit der Rechnung für die Adresse.
- Der Ablauf: `var a [4]int16` reserviert die Zellen 100 bis 107 und füllt sie mit Nullen. `a[0] = 17` schreibt 17 in die Zellen 100 und 101. `a[2] = 1000` schreibt 1000 in die Zellen 104 und 105, die Adresse ist 100 + 2 · 2 = 104. `i := 3` legt die Variable `i` vom Typ `int` an, sie belegt die 8 Zellen 108 bis 115. `a[i] = a[0] + a[2]` liest erst `i`, rechnet 100 + 3 · 2 = 106 und schreibt 1017 in die Zellen 106 und 107. `fmt.Println(a)` gibt `[17 0 1000 1017]` aus.
- Die Reihenfolge, die die Seite für Zeile 5 zeigt, `a[i] = a[0] + a[2]`: Erst liest der Rechner `a[0]` aus den Zellen 100 und 101, dann `a[2]` aus den Zellen 104 und 105, dann rechnet er 17 + 1000 = 1017. Danach liest er `i` aus den Zellen 108 bis 115, rechnet die Adresse 100 + 3 · 2 = 106 aus und schreibt 1017 in die Zellen 106 und 107. Die Reihenfolge steht auch im Kasten „Was der Rechner in dieser Zeile tut“.
- Go prüft den Index vor jedem Zugriff. Schaltet der Student vor Zeile 5 in Zelle 115 das Bit mit der Wertigkeit 4 an, hat `i` den Wert 7. Zeile 5 bricht dann ab, weil das Array nur die Indizes 0 bis 3 hat. Ohne die Prüfung käme der Rechner auf die Adresse 100 + 7 · 2 = 114, und diese Zelle gehört zu `i`.
- Schaltet der Student ein Bit von Hand um, ändert sich der Wert des Elements. Die Tabelle bleibt gleich.

### Handgriffe

- „nächste Zeile ausführen“ wählen.
- „eine Zeile zurück“ wählen: Der Stand vor der letzten Zeile kommt zurück, auch von Hand umgeschaltete Bits.
- „von vorn“ wählen: Alle Zellen sind wieder frei, die Tabelle ist leer.
- Eine Zelle im Hauptspeicher anklicken: Sie erscheint unten groß.
- Eine Zeile der Tabelle anklicken. Der Mauszeiger über einer Zeile hebt die Zellen der Variable hervor.
- Ein Bit der groß gezeigten Zelle anklicken: Es wechselt zwischen 0 und 1.
- Das Programm lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 6.1
- 6

## Aufgaben

Reiter: aufgaben

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Die erste Karte heißt „Erst vorhersagen, dann ausprobieren“ und stellt fünf Aufgaben. Die Aufgaben 1 bis 3 und 5 gehören zu Schritt 1, Aufgabe 4 gehört zu Schritt 2.
- Unter jeder Aufgabe steht eine aufklappbare „Antwort“. Sie ist verdeckt, bis der Student sie aufklappt.
- Die Antworten: Aufgabe 1: 8 Zellen und eine Zeile in der Tabelle. Aufgabe 2: 100 + 3 · 2 = 106. Aufgabe 3: Der Index sagt, wie viele Elemente vor dem gesuchten liegen. Aufgabe 4: Go prüft den Index und bricht ab. Aufgabe 5: Nein, der Rechner rechnet in beiden Fällen einmal.
- Fragt der Student nach einer Aufgabe, deren Antwort er noch nicht aufgeklappt hat, bekommt er einen Hinweis auf den Weg und nicht gleich das Ergebnis.
- Die zweite Karte nennt die vier HTML-Seiten mit dem Speichermodell: primitive Datentypen in Kapitel 4, diese zu Kapitel 6, Zeiger in Kapitel 7 und Verbund in Kapitel 8.

### Handgriffe

- Eine „Antwort“ aufklappen.
- Für Aufgabe 4 mit den Reitern oben zu Schritt 2 wechseln und dort ausprobieren.
- Den Links zu den drei anderen HTML-Seiten mit dem Speichermodell folgen.
- In diesem Schritt selbst lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 6.1
