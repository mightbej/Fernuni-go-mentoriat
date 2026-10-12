# Zahlenkreis: Ganzzahlen und Überlauf in Go

Die Seite stellt alle Bitmuster eines Ganzzahl-Typs auf einen Kreis. Ein Zeiger zeigt auf ein Bitmuster. Addieren bewegt den Zeiger im Uhrzeigersinn, Subtrahieren gegen den Uhrzeigersinn. Der Kreis zeigt entweder ein Modell mit 4 Bit oder 8 Bit, und 8 Bit sind in Go die Typen `uint8` und `int8`.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt. Der Kreis ist in Schritt 1, Schritt 2 und Schritt 3 derselbe: Die Zahl der Bits, die Lesart mit oder ohne Vorzeichen und die Stelle des Zeigers bleiben beim Wechsel des Schritts erhalten.

Den Code auf der Seite kann der Student nicht ändern und nicht ausführen. Die Seite schreibt die Programmzeilen selbst, nachdem der Student in Schritt 3 gerechnet hat.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist nur Text: was der Kreis zeigt, was die Nahtstelle ist und welche vier Schritte folgen.
- Ein Kreis ist hier noch nicht zu sehen. Er steht in Schritt 1, Schritt 2 und Schritt 3.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.5
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
- Nach der Rechnung steht im Kasten die Meldung. Ohne Überlauf ist der Kasten grün und nennt Rechnung und Ergebnis. Mit Überlauf ist der Kasten rot. Er nennt das rechnerische Ergebnis, wie oft der Zeiger über die Nahtstelle gelaufen ist und welcher Wert gespeichert wird.
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

- Zu sehen ist eine Tabelle mit den Ganzzahl-Typen von Go. Die Spalten heißen Typ, Bits, Kleinster Wert, Größter Wert und Stellen auf dem Kreis.
- Oben stehen die vier Typen ohne Vorzeichen: `uint8`, `uint16`, `uint32`, `uint64`. Darunter stehen die vier Typen mit Vorzeichen: `int8`, `int16`, `int32`, `int64`. Am Ende stehen `uint`, `int`, `byte` und `rune`.
- Gelb markiert ist die Zeile des Typs, den der Kreis in den Schritten 1 bis 3 gerade zeigt: `uint8` oder `int8`. Zeigt der Kreis das Modell mit 4 Bit, ist keine Zeile markiert.
- Ein Kreis ist in diesem Schritt nicht zu sehen.

### Handgriffe

- In der Tabelle lässt sich nichts anklicken.
- Mit den Reitern oben zu einem anderen Schritt wechseln. In Schritt 1, Schritt 2 und Schritt 3 wählt der Student zwischen 4 Bit und 8 Bit und zwischen „ohne Vorzeichen“ und „mit Vorzeichen“. Wählen lassen sich damit nur `uint8` und `int8`. Die anderen Typen der Tabelle zeigt der Kreis nicht.

### Abschnitte des Kurstexts

- 4.3.4
