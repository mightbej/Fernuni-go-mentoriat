# Binäre Suchbäume in Go

Die Seite erklärt binäre Bäume und binäre Suchbäume: die Begriffe, die Regel des Suchbaums, die Funktionen `einfügen`, `suchen`, `istWertInBaum`, `baumAusgeben` und `summeDesBaums` aus dem Kurstext und die drei Fälle beim Löschen. Sie gehört zu Kapitel 9 des Kurstexts, Abschnitt 9.3, Seiten 237–244.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Typ aus Listing 9.2: `knoten` mit den Feldern `wert int` und `links, rechts *knoten`. Er ist ein rekursiver Verbund-Typ.
- Regel des Suchbaums: Für jeden Knoten mit dem Wert x führt `links` zu einem Suchbaum mit nur kleineren Werten und `rechts` zu einem Suchbaum mit nur größeren Werten. Der Kurstext nimmt an, dass kein Wert doppelt vorkommt.
- Die Seite arbeitet mit den sieben Werten 17, 36, 42, 57, 71, 88 und 93 des Kurstexts. Ihre Summe ist 404.
- Der Baum aus Abbildung 9.5 entsteht aus der Reihenfolge 57, 36, 17, 42, 88, 93, 71: Wurzel 57, links 36 mit den Kindern 17 und 42, rechts 88 mit den Kindern 71 und 93. Höhe 3.
- Der Baum aus Abbildung 9.6 entsteht aus der Reihenfolge 36, 88, 17, 93, 57, 42, 71: Wurzel 36, links 17, rechts 88. Die 88 hat links 57 und rechts 93. Die 57 hat links 42 und rechts 71. Höhe 4. Diesen Baum verwenden die Schritte 3 bis 6.
- Der Baum aus Abbildung 9.7 entsteht aus der sortierten Reihenfolge 17, 36, 42, 57, 71, 88, 93. Jeder Knoten hat nur ein rechtes Kind. Höhe 7. Er heißt entartet.
- Die Höhe eines Knotens ist die Zahl der Knoten auf dem Pfad von der Wurzel zu ihm. Die Wurzel hat die Höhe 1. Die Höhe des Baums ist die größte Höhe eines Blatts.
- Die Seite zeichnet einen Baum als Bild: Kreise mit dem Wert, Linien für die Zeiger. Die Wurzel steht oben. Ein Zeiger mit dem Wert `nil` ist nicht gezeichnet. Gelb und grün markieren je nach Schritt den Pfad, den neuen oder den gefundenen Knoten.
- Im Modell der Seite belegt ein Knoten 24 Zellen: je 8 für `wert`, `links` und `rechts`. Die Adressen sind erfunden.
- Alle Funktionen des Kurstexts zum Baum sind rekursiv. Im Bild des Kurstexts ist der Baum ein Mobile.
- Wie ein Baum beim Einfügen balanciert bleibt, behandelt der Kurstext nicht. Er verweist auf das Modul Algorithmen und Datenstrukturen.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code und keine eigenen Werte eingeben. Alle Ausgaben und Zahlen sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt den Grund für den Suchbaum: In der Liste ist das Suchen der Stelle teuer.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Verkettete Listen in Go“, „Rekursive Verbund-Typen“, „Rekursion in Go“ und „Binäre Suche in einem Verbund-Array“ zu Abschnitt 8.4.1.
- Unter „Die Schritte“ steht je ein Satz zu den sechs Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.3
- 9.3.1

## Schritt 1: Begriffe

Reiter: begriffe

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste erklärt Knoten, Wurzel, Kante, Elternknoten und Kindknoten, Blatt und innerer Knoten, Pfad und Höhe, binärer Baum.
- Das Bild zeigt den Baum aus Abbildung 9.5. Vier Knöpfe wählen einen Knoten: 57, 36, 42 oder 93.
- Gelb ist der Pfad von der Wurzel, grün der gewählte Knoten. Ein Kasten nennt Pfad, Höhe, Elternknoten und Kindknoten.
- Knoten 57: Höhe 1, Wurzel. Knoten 36: Höhe 2, Elternknoten 57, Kinder 17 und 42. Knoten 42: Höhe 3, Elternknoten 36, Blatt. Knoten 93: Höhe 3, Elternknoten 88, Blatt.
- Eine Tabelle nennt die größte mögliche Zahl von Knoten: 1 bei Höhe 1, 3 bei Höhe 2, 7 bei Höhe 3, 15 bei Höhe 4, 2 hoch n minus 1 bei Höhe n.
- Ein Baum der Höhe 20 fasst höchstens 1 048 575 Knoten.

### Handgriffe

- Einen der vier Knoten wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 9.3

## Schritt 2: Der Suchbaum

Reiter: suchbaum

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt die Regel des Suchbaums.
- Drei Knöpfe wählen eine Reihenfolge, in der die sieben Werte eingefügt werden. Das Bild zeigt den Baum, der daraus entsteht, ein Kasten nennt Abbildung und Höhe.
- Reihenfolge 1 ergibt Abbildung 9.5 mit Höhe 3, Reihenfolge 2 Abbildung 9.6 mit Höhe 4, Reihenfolge 3 Abbildung 9.7 mit Höhe 7.
- Eine Liste erklärt „extrem balanciert“ und „entartet“.

### Handgriffe

- Eine der drei Reihenfolgen wählen.
- Die „Auflösung“ aufklappen.
- Eine eigene Reihenfolge kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 9.3.1
- 9.3.2

## Schritt 3: Einfügen

Reiter: einfuegen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Zwei Codeblöcke zeigen den Typ `knoten` und die Funktion `einfügen` aus Listing 9.2.
- `einfügen` vergleicht den neuen Wert mit dem Wert der Wurzel. Ist er kleiner, geht es nach links, sonst nach rechts. Ist dort `nil`, entsteht dort der neue Knoten mit `&knoten{wert: neuerWert}`. Sonst ruft sich die Funktion mit der Wurzel dieses Teilbaums selbst auf.
- Zwei Bedingungen an den Aufruf: `wurzel` ist nicht `nil`, und der Wert kommt im Baum noch nicht vor. Mit `nil` bricht das Programm ab mit `panic: runtime error: invalid memory address or nil pointer dereference`. Einen Wert, der schon im Baum steht, würde die Funktion rechts einhängen, weil der Zweig `else` auch den Fall „gleich“ aufnimmt.
- Die Demo hat sieben Zeilen aus `main`: `baum := &knoten{wert: 36}` und sechs Aufrufe von `einfügen` mit 88, 17, 93, 57, 42 und 71. Ein Schritt der Demo ist ein ganzer Aufruf.
- Darunter stehen die drei Knöpfe, das Bild des Baums mit der Anfangsadresse unter jedem Knoten, Rahmen mit Variablen, der Hauptspeicher und ein Kasten mit einem Satz. Eine Ausgabe gibt es nicht.
- Gelb sind die Knoten, mit denen `einfügen` verglichen hat, grün ist der neue Knoten.
- Die Knoten entstehen ab den Zellen 100 (36), 124 (88), 148 (17), 172 (93), 196 (57), 220 (42) und 244 (71). `baum` liegt in den Zellen 60 bis 67 und enthält 100.
- Aufrufe von `einfügen` je Wert, die rekursiven mitgezählt: 88 und 17 je einer, 93 und 57 je zwei, 42 und 71 je drei.
- Am Ende steht der Baum aus Abbildung 9.6.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Die Werte und ihre Reihenfolge lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 9.3.2

## Schritt 4: Suchen

Reiter: suchen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `suchen` und `istWertInBaum` aus Listing 9.2.
- `suchen` liefert einen Zeiger auf den gefundenen Knoten oder `nil`. `return wurzel` deckt beide Fälle der ersten Bedingung ab.
- Sechs Knöpfe wählen einen Suchwert: 57, 71, 36, 13, 24 oder 100. Gesucht wird im Baum aus Abbildung 9.6.
- Im Bild sind die Knoten gelb, mit denen die Suche vergleicht. Ein gefundener Knoten ist grün.
- Die Ausgabe zeigt `true` oder `false`, das Ergebnis von `istWertInBaum`.
- Aufrufe von `suchen`: 57 braucht 3, 71 braucht 4, 36 braucht einen, 13 und 24 brauchen je 3, 100 braucht 4. Bei einem Wert, der fehlt, ist der letzte Aufruf der mit `nil`.
- `fmt.Println(suchen(baum, 57))` gibt den Knoten mit einem `&` davor aus, zum Beispiel `&{57 0xc0000080c0 0xc0000080d8}`. Die zwei hinteren Zahlen sind die Adressen der Kinder. `fmt.Println(suchen(baum, 13))` gibt `<nil>` aus. Das steht nicht auf der Seite.
- Eine Tabelle vergleicht die Suche nach 93 in den drei Bäumen: 3 Aufrufe bei Höhe 3, 3 Aufrufe bei Höhe 4, 7 Aufrufe bei Höhe 7.
- In einem balancierten Baum mit n Knoten braucht die Suche etwa log2(n) Schritte, in einem entarteten bis zu n.

### Handgriffe

- Einen der sechs Suchwerte wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Suchwert kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 9.3.2
- 8.4.1

## Schritt 5: Durchlaufen

Reiter: durchlaufen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt die sieben Zeilen der Funktion `baumAusgeben` aus Listing 9.2 und arbeitet auf dem Baum aus Abbildung 9.6.
- Darunter stehen die drei Knöpfe, das Bild des Baums, die Rahmen, die Ausgabe und ein Kasten mit einem Satz. Einen Hauptspeicher zeigt diese Demo nicht.
- Je Aufruf von `baumAusgeben`, der läuft oder wartet, steht ein eigener Rahmen da. Es sind bis zu vier zugleich.
- Im Bild sind die schon ausgegebenen Knoten grün. Gelb sind die Knoten, deren Aufruf läuft oder wartet.
- Die Demo hat 13 Schritte. Die Ausgabe am Ende ist `17 36 42 57 71 88 93`.
- Die Reihenfolge linker Teilbaum, Knoten, rechter Teilbaum heißt In-Order. In einem Suchbaum liefert sie die Werte aufsteigend sortiert.
- Ein Codeblock zeigt `summeDesBaums`. Ein festes Bild zeigt unter jedem Knoten die Summe seines Teilbaums: 17, 42, 71 und 93 bei den Blättern, 170 beim Knoten 57, 351 beim Knoten 88, 404 bei der Wurzel 36.
- `summeDesBaums` nimmt erst den Knoten, dann die Teilbäume. Diese Reihenfolge heißt Pre-Order. Eine Ausgabe in Pre-Order wäre 36 17 88 57 42 71 93. Das steht nicht auf der Seite.
- Die Reihenfolge Post-Order nennt der Kurstext nicht.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Der Baum lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 9.3.2
- 7.5

## Schritt 6: Löschen

Reiter: loeschen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste nennt die drei Fälle: Der Knoten ist ein Blatt, er hat ein Kind, er hat zwei Kinder.
- Zwei feste Bilder zeigen den Baum aus Abbildung 9.6 vor und nach dem Löschen der 88. Vorher sind 88 und 57 gelb und 71 grün. Nachher steht die 71 an der Stelle der 88, mit den Kindern 57 und 93, und die 57 hat nur noch das linke Kind 42.
- Der Weg zum Ersatz: von der 88 einmal nach links zur 57, dann nach rechts zur 71. Dort geht es rechts nicht weiter. Die 71 ist der größte Wert im linken Teilbaum.
- Ebenso ginge der kleinste Wert im rechten Teilbaum, hier die 93.
- Go-Code zum Löschen steht im Kurstext nicht und auf der Seite nicht. Fragt der Student danach, bekommt er die drei Fälle erklärt und kein fertiges Programm.
- Unter „Wie es weitergeht“ stehen Links auf die Lektionen „Verkettete Listen in Go“ und „Rekursion in Go“ und auf den Exkurs zur Laufzeit.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.3.2
