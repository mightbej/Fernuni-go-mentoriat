# Anonyme Felder und eingebettete Verbünde in Go

Die Seite erklärt anonyme Felder: ein Feld, das nur mit seinem Verbund-Typ und ohne Namen deklariert ist, und den verkürzten Zugriff auf die Felder des eingebetteten Verbunds. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.3.2, Seiten 216–220.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein anonymes Feld, englisch anonymous field, ist ein Feld, bei dessen Deklaration nur der Typ dasteht. Ein Verbund in einem anonymen Feld heißt eingebetteter Verbund, englisch embedded struct.
- Ein anonymes Feld hat doch einen Namen: Go verwendet den Bezeichner seines Typs.
- Die Typen der Seite in der dritten Variante: `punkt` mit `x, y int`. `kreis` mit dem anonymen Feld `punkt` und `radius int`. `rad` mit dem anonymen Feld `kreis` und `speichen int`.
- Beim Zugriff dürfen die anonymen Felder fehlen: `r.x` steht für `r.kreis.punkt.x`. In einem Literal gibt es diese Abkürzung nicht.
- Zwei anonyme Felder des gleichen Typs in einem Verbund sind nicht zulässig.
- Der Kurstext nennt als Quelle der Beispiel-Idee das Buch von Donovan und Kernighan.
- Für die verkettete Liste in Kapitel 9 braucht man Abschnitt 8.3.1, die rekursiven Verbund-Typen. Anonyme Felder kommen dort nicht vor.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Meldungen sind mit Go 1.26 geprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt das Problem: Ein hierarchischer Verbund macht den Zugriff auf ein Feld der untersten Ebene lang.
- Das Beispiel stammt aus einem Zeichenprogramm: Ein Kreis hat einen Mittelpunkt mit x und y und einen Radius. Ein Rad ist ein Kreis mit einer Zahl von Speichen.
- Unter „Was dazugehört“ stehen zwei Punkte mit Links: die Lektion „Hierarchische Verbünde“ und die Lektion „Verbund in Go“.
- Ein gelber Kasten sagt, dass man für Kapitel 9 zuerst die Lektion „Rekursive Verbund-Typen“ braucht.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.3.2
- 8.3

## Schritt 1: Flach, hierarchisch, eingebettet

Reiter: varianten

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Drei Knöpfe wählen eine Variante. Darunter stehen auf dunklem Grund die Deklarationen und der Zugriff auf die vier Werte eines Rads, darunter ein Kasten mit einem Satz.
- Variante 1, flach: `kreisV1` mit `x, y, radius int` und `radV1` mit `x, y, radius, speichen int`. Zugriff: `rV1.x = 5`.
- Variante 2, hierarchisch: `punkt` mit `x, y int`, `kreisV2` mit `mitte punkt` und `radius int`, `radV2` mit `umkreis kreisV2` und `speichen int`. Zugriff: `rV2.umkreis.mitte.x = 5`.
- Variante 3, eingebettet: `kreis` mit dem anonymen Feld `punkt`, `rad` mit dem anonymen Feld `kreis`. Zugriff: `r.x = 5`. Im Kurstext heißt sie die beste Variante.
- Unter „Im Speicher sind alle drei gleich“ steht: Ein Rad besteht in jeder Variante aus vier `int`, die hintereinander liegen, bei 8 Zellen je `int` 32 Zellen. Die Varianten unterscheiden sich nur in den Namen. Dieser Absatz steht nicht im Kurstext.

### Handgriffe

- Eine der drei Varianten wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.3.2

## Schritt 2: Zugriff über anonyme Felder

Reiter: zugriff

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt die drei Typen der dritten Variante.
- Die Demo zeigt acht Zeilen im Rumpf von `main`: `var r rad`, `r.x = 5`, `r.y = 3`, `r.radius = 20`, `r.speichen = 8`, `fmt.Printf("%+v\n", r)`, `r.kreis.punkt.x = 7`, `fmt.Println(r.x)`.
- Unter dem Programm stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, der Rahmen von `main`, der Hauptspeicher, die Ausgabe und ein Kasten mit einem Satz.
- Im Rahmen steht `r` als Kasten mit vier Zeilen. Jede nennt ein Feld mit seinem ganzen Weg: `kreis.punkt.x`, `kreis.punkt.y`, `kreis.radius` und `speichen`.
- Im Hauptspeicher belegt `r` die Zellen 100 bis 131, das sind 32 Zellen für vier `int`. Die Adressen sind erfunden.
- Die Ausgabe: `{kreis:{punkt:{x:5 y:3} radius:20} speichen:8}` und nach `r.kreis.punkt.x = 7` die Zahl `7`.
- Unter „Drei Wege zum gleichen Feld“ stehen `r.x = 5`, `r.kreis.x = 5` und `r.kreis.punkt.x = 5`. Alle drei sind erlaubt und meinen das Gleiche.
- Ein gelber Kasten sagt: Zwei anonyme Felder des gleichen Typs sind nicht erlaubt. Auf `type doppelt struct { punkt; punkt }` meldet Go `punkt redeclared`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.3.2
- 8.2.2

## Schritt 3: Literale

Reiter: literale

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Für Literale gibt es keine Abkürzung. Es bleiben die zwei Schreibweisen aus Abschnitt 8.2.1. Mit Feldnamen steht für ein anonymes Feld der Name seines Typs.
- Fünf Knöpfe zeigen je ein Literal. Darunter steht, was `fmt.Printf("%+v\n", r)` nach der Zuweisung ausgibt, oder die Meldung des Compilers.
- Die fünf Literale:
  1. `r = rad{kreis{punkt{22, 33}, 55}, 4}` ergibt `{kreis:{punkt:{x:22 y:33} radius:55} speichen:4}`.
  2. `r = rad{kreis: kreis{punkt: punkt{x: 22, y: 33}, radius: 55}, speichen: 4}` ergibt das Gleiche.
  3. `r = rad{kreis: kreis{punkt: punkt{y: 99}}}` ergibt `{kreis:{punkt:{x:0 y:99} radius:0} speichen:0}`.
  4. `r = rad{x: 5}` meldet `unknown field x in struct literal of type rad`.
  5. `r = rad{kreis{punkt{22, 33}, 55}}` meldet `too few values in struct literal of type rad`.
- Darunter steht das zweite Literal über mehrere Zeilen, wie im Kurstext.
- Am Ende des Schritts unter „Nächste Schritte“ stehen drei Punkte: die Lektion „Rekursive Verbund-Typen“, Abschnitt 8.4 mit binärer Suche und Bubblesort, und das Go-Programm `S225_anonyme_felder_demo.go`.

### Handgriffe

- Eines der fünf Literale wählen.
- Die „Auflösung“ aufklappen.
- Ein eigenes Literal lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 8.3.2
- 8.2.1

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Drei Varianten, 8.3.2, Seiten 216–218: Alle Deklarationen und Zugriffe stammen aus dem Kurstext. Der Absatz zum Speicher ist eine Ergänzung.
  - 2 Zugriff, 8.3.2, Seiten 218–219: Die Typen, die Zuweisungen und die Ausgabe mit %+v stammen aus dem Kurstext. Die Zeilen 7 und 8 der Demo und die Meldung des Compilers sind ergänzt, geprüft mit Go 1.26.
  - 3 Literale, 8.3.2, Seiten 219–220: Die Literale 1 bis 3 stammen aus dem Kurstext. Die Literale 4 und 5 mit den Meldungen des Compilers sind ergänzt, geprüft mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 8.3, Seiten 212–214, hierarchische Verbünde: Ein Verbund als benanntes Feld. Lektion: Hierarchische Verbünde: ein Verbund als Feld.
  - Abschnitt 8.2.2, Seiten 202–203, Ausgabe: %+v gibt Feldnamen mit aus. Lektion: Verbund in Go, Schritt 4.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2.2
- 8.3
- 8.3.2
