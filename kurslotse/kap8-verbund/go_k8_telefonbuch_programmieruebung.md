# Programmierübung: Ein Telefonbuch als Verbund-Array

Die Seite ist eine Programmierübung. Drei Funktionen arbeiten auf einem Array von Verbünden: `indexJüngste` sucht die jüngste Person, `zähleDoppelte` zählt Paare mit dem gleichen Namen, `umkehren` kehrt die Reihenfolge um. Die Seite gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.4, Seiten 221–226. Die Funktionen selbst stehen nicht im Kurstext.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Typ ist `person` mit den Feldern `name string`, `telefonnummer int` und `alter int`. Das Feld `alter` kommt im Kurstext bei diesem Typ nicht vor.
- Das Telefonbuch: `telbuch := [6]person{{"Meier", 4711, 34}, {"Becker", 5522, 58}, {"Koch", 7745, 19}, {"Meier", 9081, 41}, {"Arnold", 3310, 27}, {"Becker", 1234, 63}}`.
- Jede Funktion bekommt einen Zeiger auf das Array, `a *[6]person`. Der Aufruf lautet zum Beispiel `indexJüngste(&telbuch)`. Kopiert wird eine Adresse und kein Verbund. Das Muster ist `bubblesort(a *[10]person)` in Listing 8.2.
- `len(a)` liefert auch für einen Zeiger auf ein Array dessen Länge, hier 6. `a[i]` ist die Kurzschreibweise für `(*a)[i]`.
- In den Demos steht ein Element des Arrays als `{Name Telefonnummer Alter}`, zum Beispiel `{Meier 4711 34}`. Grün sind die Elemente, die der letzte Schritt gelesen oder getauscht hat.
- Jede Demo zeigt über dem Programm von Anfang an den Rahmen von `main` mit dem Telefonbuch, auch bevor eine Anweisung ausgeführt ist. Darunter stehen das Programm und die drei Knöpfe. Unter den Knöpfen steht der Rahmen der Funktion mit dem Zeiger `a` und ihren übrigen Variablen, dann die Ausgabe und ein Kasten mit einem Satz.
- Einen Hauptspeicher mit Zellen und Adressen zeigt diese Übung nicht. Sie arbeitet eine Ebene über dem Speichermodell. Die Speichersicht für Zeiger auf ein Array steht in den Programmierübungen zu Kapitel 7.
- Die Ausgabe des ganzen Programms: „Jüngste Person: Koch 19“, „Älteste Person: Becker 63“, „Paare mit gleichem Namen: 2“ und „Umgekehrt: [{Becker 1234 63} {Arnold 3310 27} {Meier 9081 41} {Koch 7745 19} {Becker 5522 58} {Meier 4711 34}]“.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays. Ein Element aus einem Array zu entfernen, zeigt die Seite deshalb nicht.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Das Programm ist mit Go 1.26 übersetzt und ausgeführt.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock mit dem Typ `person` und dem Telefonbuch. Eine Demo gibt es hier nicht.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Bubblesort auf einem Verbund-Array“ und „Vergleich von Verbünden“ und auf die Programmierübung „Selection Sort mit Zeigern“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4
- 8.2

## Schritt 1: Die jüngste Person

Reiter: juengste

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Über dem Programm steht der Rahmen von `main` mit dem Telefonbuch. Darunter stehen die neun Zeilen von `indexJüngste`, die drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, der Rahmen von `indexJüngste`, die Ausgabe und ein Kasten mit einem Satz.
- Die Variable `j` aus `main`, die den Rückgabewert aufnimmt, steht nicht im Rahmen von `main`. Der Rückgabewert 2 steht im Satz des letzten Schritts und in der Ausgabe.
- Die Funktion merkt sich in `minIndex` den Index der jüngsten Person bisher. Sie beginnt bei 0 und vergleicht `a[i].alter < a[minIndex].alter` für i von 1 bis 5.
- Der Lauf: 58 < 34 ist false. 19 < 34 ist true, `minIndex` wird 2. 41 < 19, 27 < 19 und 63 < 19 sind false. Die Funktion liefert 2.
- Die Demo hat 8 Schritte. Die Ausgabe am Ende ist „Jüngste Person: Koch 19“.
- Die Funktion liefert den Index und keinen Verbund. Mit dem Index erreicht `main` den Eintrag im Array selbst.
- `indexÄlteste` vergleicht mit `>` und liefert 5, Becker mit 63 Jahren.
- Sind zwei Personen gleich alt, liefert `indexJüngste` die mit dem kleineren Index, weil `<` bei Gleichheit false ist. Mit `<=` liefert sie die mit dem größeren Index.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Telefonbuch lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 8.4
- 8.2.3

## Schritt 2: Doppelte Namen

Reiter: doppelt

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt die elf Zeilen von `zähleDoppelte`: eine äußere Schleife über i von 0 bis 4, eine innere Schleife über j von i+1 bis 5, den Vergleich `a[i].name == a[j].name` und `anzahl++`.
- Ein Schritt der Demo ist ein ganzer Durchlauf der inneren Schleife für ein i. Die Demo hat 7 Schritte.
- Der Lauf: Bei i = 0 steht am Index 3 der gleiche Name Meier. Bei i = 1 steht am Index 5 der gleiche Name Becker. Bei i = 2, 3 und 4 ist kein Name gleich. Die Funktion liefert 2.
- Zusammen sind es 5 + 4 + 3 + 2 + 1 = 15 Vergleiche, nach der Formel n · (n − 1) / 2.
- `telbuch[0] == telbuch[3]` ist false: `==` vergleicht alle Felder, und Telefonnummer und Alter sind verschieden.
- Die Funktion zählt Paare. Stünde ein Name dreimal im Array, ergäbe das drei Paare.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.5
- 8.4.2

## Schritt 3: Umkehren

Reiter: umkehren

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt die fünf Zeilen von `umkehren`: `for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1` und im Rumpf `a[i], a[j] = a[j], a[i]`.
- Der Lauf: Tausch der Indizes 0 und 5, dann 1 und 4, dann 2 und 3. Danach ist i gleich 3 und j gleich 2, die Schleife endet. Es sind drei Vertauschungen.
- Die Demo hat 9 Schritte. Die Ausgabe am Ende ist die Zeile „Umgekehrt: …“ mit Becker 1234 vorn und Meier 4711 hinten.
- Die Funktion hat keine Rückgabe. Sie ändert das Array von `main` über den Zeiger. Mit einem Parameter vom Typ `[6]person` würde sie eine Kopie umkehren, und `main` gäbe das unveränderte Array aus.
- Eine Liste erklärt die drei Tupel-Schreibweisen der Funktion.
- Bei einer ungeraden Zahl von Elementen bleibt das mittlere Element stehen.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 5.3.2
- 7.4

## Schritt 4: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das ganze Programm mit dem Typ `person`, den vier Funktionen `indexJüngste`, `indexÄlteste`, `zähleDoppelte` und `umkehren` und mit `main`.
- Darunter stehen der Knopf „Programm kopieren“ und die Ausgabe mit vier Zeilen.
- In `main` heißt eine Variable `ä`. Go erlaubt Umlaute in Bezeichnern, der Kurstext verwendet sie auch.
- Unter „Zum Weiterdenken“ stehen vier Aufgaben ohne Lösung: eine Funktion `zähleÄlterAls` schreiben, eine Funktion `istSortiert` schreiben, die Frage nach zwei gleich alten Personen, `umkehren` ohne Zeiger.
- Zur zweiten Aufgabe: `istSortiert` vergleicht jedes Element mit seinem rechten Nachbarn, das sind höchstens 5 Vergleiche.
- Zur vierten Aufgabe: Ohne Zeiger gibt `main` das Telefonbuch in der ursprünglichen Reihenfolge aus.
- Eine Demo gibt es hier nicht.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf die Lektionen „Bubblesort auf einem Verbund-Array“ und „Verkettete Listen in Go“.

### Handgriffe

- „Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 8.4
- 7.4

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit sechs Zeilen: Abschnitt 5.3.2, Kapitel 6, Abschnitt 7.4, 8.2.3, 8.2.5 und 8.4.2, jeweils mit den gedruckten Seiten und einem Satz, was das Programm daraus verwendet.
- Darunter steht: Die vier Funktionen kommen im Kurstext nicht vor.
- Unter der Tabelle steht außerdem: Die Funktionen sind eigene Beispiele, geprüft mit Go 1.26. Der Zeiger auf das Array folgt Listing 8.2, Seite 226. Die Formel für die Zahl der Vergleiche in Schritt 2 steht in Abschnitt 8.4.2, Seite 225.

### Handgriffe

- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4
- 8.4.2
