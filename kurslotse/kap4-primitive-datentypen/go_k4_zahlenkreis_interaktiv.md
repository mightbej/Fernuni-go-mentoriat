# Zahlenkreis: Ganzzahlen und Überlauf in Go

Die Seite stellt alle Bitmuster eines Ganzzahl-Typs auf einen Kreis. Ein Zeiger zeigt auf ein Bitmuster. Addieren bewegt den Zeiger im Uhrzeigersinn, Subtrahieren gegen den Uhrzeigersinn. Der Kreis zeigt entweder ein Modell mit 4 Bit oder 8 Bit, und 8 Bit sind in Go die Typen `uint8` und `int8`.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt. Der Kreis ist in Schritt 1, 2, 3 und 5 derselbe: Die Zahl der Bits, die Lesart mit oder ohne Vorzeichen und die Stelle des Zeigers bleiben beim Wechsel des Schritts erhalten.

Den Code auf der Seite kann der Student nicht ändern und nicht ausführen. Die Seite schreibt die Programmzeilen selbst, nachdem der Student in Schritt 3 oder 5 gerechnet hat.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist nur Text: was der Kreis zeigt, was die Nahtstelle ist und welche fünf Schritte folgen.
- Der Text grenzt zwei Fälle ab: Liegt das rechnerische Ergebnis über dem größten Wert oder bei einem Typ mit Vorzeichen unter dem kleinsten, ist das der arithmetische Überlauf. Ein negatives Ergebnis bei einem Typ ohne Vorzeichen ist das Repräsentationsproblem aus Schritt 5.
- Ein Kreis ist hier noch nicht zu sehen. Er steht in Schritt 1, 2, 3 und 5.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.5
- 4.3.6
- 4.3.8

## Schritt 1: Bitmuster und Wert

Reiter: bitmuster

### Aufbau der Darstellung

- Links steht der Kreis, darunter das Bitmuster in großen Ziffern. Rechts daneben stehen die Einstellungen, eine Anzeige und eine Meldung.
- Jede Stelle auf dem Kreis ist ein Bitmuster. Ganz oben steht das Bitmuster aus lauter Nullen. Im Uhrzeigersinn ist jedes Bitmuster um 1 größer als das davor.
- Mit 4 Bit hat der Kreis 16 runde Felder. In jedem Feld steht der Wert, außen neben dem Feld steht das Bitmuster. Das Feld, auf das der Zeiger zeigt, hat einen dicken schwarzen Rand.
- Mit 8 Bit hat der Kreis 256 kurze Striche. Jeder sechzehnte Strich ist länger und trägt seinen Wert. Ein oranger Punkt markiert die Stelle, auf die der Zeiger zeigt.
- Grün sind die Stellen mit einem Wert ab 0. Blau sind die Stellen mit einem negativen Wert. Blaue Stellen gibt es nur in der Lesart mit Vorzeichen.
- Der schwarze Pfeil aus der Mitte ist der Zeiger.
- Die rote gestrichelte Linie quer zum Kreis ist die Nahtstelle. Sie liegt zwischen dem größten und dem kleinsten Wert und ist beschriftet mit „Nahtstelle zwischen … und …“. Ohne Vorzeichen liegt sie oben, mit Vorzeichen unten.
- Unter dem Kreis steht das Bitmuster, auf das der Zeiger zeigt, als Reihe von Knöpfen. Über jedem Bit steht seine Wertigkeit. Ein Bit auf 1 ist farbig hinterlegt, ein Bit auf 0 ist weiß mit grauer Ziffer. Unter der Reihe steht die Summe der Wertigkeiten aller Bits auf 1, zum Beispiel 4 + 1 = 5.
- Rechts steht die Anzeige mit vier Zeilen: Typ, Wertebereich, Bitmuster, Wert.
- Unter der Anzeige steht ein grüner Kasten mit einer Meldung, zum Beispiel „Startwert 5, Bitmuster 0101“.

### Handgriffe

- Ein Bit unter dem Kreis anklicken und damit umschalten. Der Zeiger springt auf das neue Bitmuster.
- Eine Stelle auf dem Kreis anklicken. Der Zeiger springt dorthin.
- „4 Bit (Modell)“ oder „8 Bit“ wählen. Beim Wechsel springt der Zeiger auf 0. Andere Größen zeigt der Kreis nicht: `int`, `int16`, `int32` und `int64` lassen sich nicht wählen.
- „ohne Vorzeichen“ oder „mit Vorzeichen“ wählen. Das Bitmuster bleibt dabei stehen.
- Einen Startwert eingeben und „als Startwert setzen“ wählen. Der Startwert muss im Wertebereich liegen: 0 bis 15 oder −8 bis 7 bei 4 Bit, 0 bis 255 oder −128 bis 127 bei 8 Bit.
- „an den größten Wert“ wählen. Der Zeiger springt auf den größten Wert des Wertebereichs.
- Rechnen kann der Student in diesem Schritt nicht. Die Knöpfe dazu stehen in Schritt 3.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.5

## Schritt 2: Mit und ohne Vorzeichen

Reiter: vorzeichen

### Aufbau der Darstellung

- Über dem Kreis steht eine Frage zum Vorhersagen in einem gelben Kasten. Die Auflösung ist eingeklappt.
- Darunter steht derselbe Kreis wie in Schritt 1, mit demselben Stand: links der Kreis, darunter das Bitmuster in großen Ziffern, rechts Einstellungen, Anzeige und Meldung.
- Jede Stelle auf dem Kreis ist ein Bitmuster. Ganz oben steht das Bitmuster aus lauter Nullen. Im Uhrzeigersinn ist jedes Bitmuster um 1 größer als das davor.
- Schaltet der Student zwischen „ohne Vorzeichen“ und „mit Vorzeichen“ um, bleiben das Bitmuster und die Stelle des Zeigers stehen. Es ändern sich vier Dinge: der Wert, die Wertigkeit des vordersten Bits, die Farben auf dem Kreis und die Lage der Nahtstelle.
- Ohne Vorzeichen hat das vorderste Bit eine positive Wertigkeit: 8 bei 4 Bit, 128 bei 8 Bit. Alle Stellen des Kreises sind grün. Die Nahtstelle liegt oben, zwischen dem größten Wert und 0.
- Mit Vorzeichen hat das vorderste Bit eine negative Wertigkeit: −8 bei 4 Bit, −128 bei 8 Bit. Die rechte Hälfte des Kreises ist grün und trägt die Werte ab 0. Die linke Hälfte ist blau und trägt die negativen Werte. Die Nahtstelle liegt unten, zwischen dem größten und dem kleinsten Wert.
- Unter dem Kreis steht über jedem Bit seine Wertigkeit. Ein Bit auf 1 ist grün hinterlegt. Das vorderste Bit ist blau hinterlegt, wenn es auf 1 steht und eine negative Wertigkeit hat. Unter der Reihe steht die Summe, zum Beispiel −128 + 64 + 1 = −63.
- Mit 4 Bit hat der Kreis 16 runde Felder mit Wert und Bitmuster. Mit 8 Bit hat er 256 Striche, jeder sechzehnte trägt seinen Wert, und ein oranger Punkt markiert die Stelle des Zeigers.
- Rechts steht die Anzeige mit Typ, Wertebereich, Bitmuster und Wert. Darunter steht die Meldung, zum Beispiel „Bitmuster 1111 1111 wird als int8 gelesen: Wert −1“.

### Handgriffe

- „ohne Vorzeichen“ oder „mit Vorzeichen“ wählen. Bei 8 Bit heißen die Knöpfe „uint8 – ohne Vorzeichen“ und „int8 – mit Vorzeichen“.
- „4 Bit (Modell)“ oder „8 Bit“ wählen. Beim Wechsel springt der Zeiger auf 0. Andere Größen zeigt der Kreis nicht: `int`, `int16`, `int32` und `int64` lassen sich nicht wählen.
- Ein Bit unter dem Kreis anklicken und damit umschalten.
- Eine Stelle auf dem Kreis anklicken. Der Zeiger springt dorthin.
- Einen Startwert eingeben und „als Startwert setzen“ wählen. Der Startwert muss im Wertebereich liegen: 0 bis 15 oder −8 bis 7 bei 4 Bit, 0 bis 255 oder −128 bis 127 bei 8 Bit.
- „an den größten Wert“ wählen.
- Die Auflösung der Frage aufklappen.
- Rechnen kann der Student in diesem Schritt nicht. Die Knöpfe dazu stehen in Schritt 3.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.5

## Schritt 3: Überlauf

Reiter: ueberlauf

### Aufbau der Darstellung

- Über dem Kreis stehen vier Fragen zum Vorhersagen in einem gelben Kasten. Die Auflösungen sind eingeklappt.
- Darunter steht derselbe Kreis wie in Schritt 1 und Schritt 2, mit demselben Stand: links der Kreis, darunter das Bitmuster in großen Ziffern, rechts Einstellungen, Anzeige, die Knöpfe zum Rechnen, die Meldung und der Go-Code.
- Jede Stelle auf dem Kreis ist ein Bitmuster. Im Uhrzeigersinn ist jedes Bitmuster um 1 größer als das davor. Grün sind die Stellen mit einem Wert ab 0, blau die Stellen mit einem negativen Wert.
- Die rote gestrichelte Linie ist die Nahtstelle. Dort liegt der größte Wert neben dem kleinsten. Ohne Vorzeichen liegt sie oben, bei 8 Bit zwischen 255 und 0. Mit Vorzeichen liegt sie unten, bei 8 Bit zwischen 127 und −128.
- Rechnet der Student, läuft der Zeiger Stelle für Stelle: im Uhrzeigersinn beim Addieren einer positiven Zahl, gegen den Uhrzeigersinn beim Addieren einer negativen Zahl. Bis 16 Stellen läuft er langsam, darüber schnell.
- Läuft der Zeiger über die Nahtstelle, wird die rote Linie kurz dick und durchgezogen. Jeder Lauf über die Nahtstelle macht das Ergebnis um 2^n falsch: um 16 bei 4 Bit, um 256 bei 8 Bit.
- Nach der Rechnung steht im Kasten die Meldung. Ohne Überlauf ist der Kasten grün und nennt Rechnung und Ergebnis. Beim Überschreiten der Nahtstelle ist der Kasten rot. Bei einem negativen Ergebnis ohne Vorzeichen steht „Negativer Wert nicht darstellbar.“, sonst „Überlauf.“. Er nennt das rechnerische Ergebnis, wie oft der Zeiger über die Nahtstelle gelaufen ist und welcher Wert gespeichert wird.
- Unter der Meldung stehen bei 8 Bit drei oder mehr Programmzeilen in Go mit derselben Rechnung. Hinter `//` steht die Ausgabe. Bei 4 Bit steht dort der Satz, dass Go keinen Typ mit 4 Bit kennt.
- Ist die Zahl größer als der größte Wert des Typs, steht im Go-Code eine Schleife mit `x++` oder `x--`. Der Grund: Die Zahl selbst passt dann nicht in den Typ, zum Beispiel 200 nicht in `int8` mit dem größten Wert 127. Die Zeile `x = x + 200` übersetzt Go deshalb nicht. Mit der Summe hat das nichts zu tun. Die Schleife zählt stattdessen 200-mal um 1 weiter.
- Einen Knopf für eine bestimmte Zahl gibt es nicht. Der Student trägt die Zahl in das Feld neben „addieren“ ein.
- Unter dem Kreis steht das Bitmuster mit der Wertigkeit über jedem Bit und der Summe darunter.
- Unter der Demo steht ein gelber Kasten mit dem Hinweis, dass Go den Überlauf zur Laufzeit nicht meldet.

### Handgriffe

- „− 1“ oder „+ 1“ wählen. Der Zeiger geht eine Stelle weiter.
- Eine ganze Zahl eingeben und „addieren“ wählen: −2000 bis 2000. Eine negative Zahl subtrahiert.
- Einen Startwert eingeben und „als Startwert setzen“ wählen. Der Startwert muss im Wertebereich liegen: 0 bis 15 oder −8 bis 7 bei 4 Bit, 0 bis 255 oder −128 bis 127 bei 8 Bit.
- „an den größten Wert“ wählen. Von dort führt „+ 1“ über die Nahtstelle.
- „4 Bit (Modell)“ oder „8 Bit“ wählen. Beim Wechsel springt der Zeiger auf 0. Andere Größen zeigt der Kreis nicht: `int`, `int16`, `int32` und `int64` lassen sich nicht wählen. Go-Code zeigt die Seite nur bei 8 Bit.
- „ohne Vorzeichen“ oder „mit Vorzeichen“ wählen.
- Ein Bit unter dem Kreis anklicken oder eine Stelle auf dem Kreis anklicken.
- Die Auflösung einer Frage aufklappen.
- Multiplizieren und Dividieren kann der Student auf der Seite nicht.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.8

## Schritt 4: Die Ganzzahl-Typen in Go

Reiter: typen

### Aufbau der Darstellung

- Ein Kreis ist in diesem Schritt nicht zu sehen.
- Oben steht über die ganze Breite die Tabelle der Ganzzahl-Typen mit Typ, Bits, kleinstem und größtem Wert sowie Stellen auf dem Kreis. Die Zeile des 8-Bit-Typs, auf den der Kreis zuletzt gestellt war, ist gelb markiert. Beim 4-Bit-Modell ist keine Zeile markiert. Die Tabelle ist nicht anklickbar, die Zeilen byte und rune werden nie markiert.
- Unter der Tabelle steht „Welchen Typ wähle ich?“ mit drei Punkten: `int` als Normalfall, wenn sein Wertebereich ausreicht. Eine feste Bitzahl braucht einen Grund, etwa einen großen Wertebereich, geringen Speicherbedarf oder ein Dateiformat. Vermeintlich nur positive Werte sind kein Grund für `uint`, vorzeichenlose Typen nimmt man nur dann, wenn man mit Bitmustern arbeitet.
- Darunter steht „Gleiche Größe bedeutet nicht gleicher Typ“: `int` und `int64` sind verschiedene Typen, auch bei gleicher Bitzahl, eine Zuweisung zwischen ihnen braucht eine Typkonvertierung. `byte` und `uint8` bezeichnen dagegen denselben Typ, ebenso `rune` und `int32`. Mit `byte` betont man die Verwendung als Folge von 8 Bit. Die Zeichen erklärt die verlinkte Lektion zu Strings.

### Handgriffe

- Den Links zur Typkonvertierung und zu Strings folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Den Kreis gibt es in Schritt 1, 2, 3 und 5.

### Abschnitte des Kurstexts

- 4.3.4
- 4.3.6

## Schritt 5: 10 minus 15 mit uint8

Reiter: uint8

### Aufbau der Darstellung

- Oben stehen nebeneinander zwei feste Programme, die der Student nicht ändern und nicht ausführen kann: `var anzahl int = 10` beziehungsweise `var anzahl uint8 = 10`, dann `anzahl = anzahl - 15` und `fmt.Println(anzahl)`. Die Ausgaben sind -5 und 251, beide Programme laufen ohne Fehlermeldung.
- Ein gelber Kasten stellt eine Frage zum Vorhersagen: An welcher Stelle überschreitet der Zeiger die Nahtstelle?
- Darunter steht der Knopf „10 − 15 im Zahlenkreis“ und derselbe Kreis wie in Schritt 1 bis 3, mit Einstellungen, Bits, Anzeige, Meldung und Go-Code. Er zeigt nur das 4-Bit-Modell oder `int8` beziehungsweise `uint8`, keinen `int` mit 32 oder 64 Bit.
- Der Knopf stellt immer 8 Bit ohne Vorzeichen und Startwert 10 ein, auch wenn vorher `int8` gewählt war, und lässt den Zeiger 15 Stellen zurücklaufen. Von 10 geht es über 0 und 255 bis 251, Bitmuster 1111 1011, denn −5 + 256 = 251. Die Meldung lautet dann „Negativer Wert nicht darstellbar.“
- Darunter steht ein zugeklappter Abschnitt „Warum 251?“: `uint8` kann −5 nicht darstellen. In einer Fußnote zu Hinweis 4.7 im Kurstext wird dieses Überschwappen vom arithmetischen Überlauf unterschieden, es ist ein Repräsentationsproblem. `int` speichert −5 mit 32 oder 64 Bit. Darf eine Anzahl nicht negativ sein, ist ein negatives Ergebnis ein Fehler der Berechnung, den man mit `int` erkennen kann. Vor Rechenfehlern und Überläufen schützt `int` nicht.
- Unter „Wie es weitergeht“ stehen Links auf die Lektion „Zweierkomplement“, auf das Quiz zu den Wertebereichen und als Ausblick auf den Exkurs „Blockchain: eine verkettete Liste mit Fingerabdruck“, in dem ein Guthaben im Typ `uint64` steht.

### Handgriffe

- „10 − 15 im Zahlenkreis“ wählen, um den Fall mit `uint8` zu starten oder zu wiederholen. Während der Zeiger läuft, sind dieser Knopf und die Knöpfe der Demo gesperrt.
- Alle Handgriffe des Kreises aus Schritt 3 stehen auch hier bereit: 4 oder 8 Bit, Lesart, Bits und Startwert ändern, −1 oder +1 und ganze Zahlen von −2000 bis 2000 addieren.
- „Warum 251?“ auf- und zuklappen.
- Beim Wechsel in Schritt 1 bis 3 bleibt der Stand des Kreises erhalten.

### Abschnitte des Kurstexts

- 4.3.6
- 4.3.8

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Bitmuster und Wert, 4.3.5, Seiten 64–66: Die Wertigkeit der Bits und der Wert als Summe der Wertigkeiten stammen aus dem Kurstext. Der Kreis und das Modell mit 4 Bit gibt die Lektion dazu.
  - 2 Vorzeichen, 4.3.5, Seiten 64–66: Zweierkomplement: Das Zweierkomplement mit der negativen Wertigkeit des vordersten Bits stammt aus dem Kurstext. Die Färbung der Werte auf dem Kreis gibt die Lektion dazu.
  - 3 Überlauf, 4.3.8, Seiten 78–84: Der arithmetische Überlauf stammt aus dem Kurstext. Die Nahtstelle auf dem Kreis und das Programm der Demo zum Rechnen gibt die Lektion dazu.
  - 4 Die Typen in Go, 4.3.4, Seite 63 und 4.3.6, Seiten 66–68: Die Ganzzahl-Typen, die Wahl des Typs, der Unterschied zwischen `int` und `int64` und der Alias stammen aus dem Kurstext. Die Markierung der Tabellenzeile ergänzt die Lektion.
  - 5 10 − 15 mit uint8, 4.3.6, Seiten 66–68, Hinweis 4.7: Die Rechnung, das Ergebnis 251 und die Unterscheidung vom arithmetischen Überlauf stammen aus Hinweis 4.7 und seiner Fußnote. Der Vergleich mit `int`, der Weg am Kreis und der Ausblick auf den Blockchain-Exkurs ergänzen die Lektion.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speicherelemente: Wie Folgen von 0 und 1 als Werte gelesen werden, hängt vom Typ ab. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 2.3, Seiten 10–11, Wertebereich: Bei großen Zahlen kann ein Ergebnis zu groß für den Wertebereich sein. Lektion: Überlauf bei Ganzzahlen in Go, Schritt 4.
  - Abschnitt 4.3.1, Seiten 59–61, Literale: Lektion: Zahlensysteme in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 2.3
- 4.3.1
- 4.3.4
- 4.3.5
- 4.3.6
- 4.3.8
