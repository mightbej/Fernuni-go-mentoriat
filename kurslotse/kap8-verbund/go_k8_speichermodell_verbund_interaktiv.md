# Speichermodell: Verbund-Datentyp in Go

Die Seite zeigt, wie ein Verbund im Hauptspeicher liegt: Die Felder liegen lückenlos hintereinander, die Namenstabelle im Modell nennt nur den Anfang, und ein Bauplan des Typs sagt, wo jedes Feld beginnt. Die Seite gehört zu Kapitel 8 des Kurstexts, zu Abschnitt 8.2 und 8.2.3. Sie ist die vierte von vier HTML-Seiten mit dem gleichen Modell.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Hauptspeicher ist eine Reihe von Zellen. Jede Zelle fasst 8 Bit, also 1 Byte, und hat eine Adresse.
- Neben dem Hauptspeicher zeigt die Seite eine Namenstabelle im Modell. In jeder Zeile stehen ein Name, ein Typ und eine Adresse. Der Wert steht nur im Hauptspeicher.
- Die Seite sagt dazu wörtlich: „Diese Tabelle gibt es so nicht. Sie ist ein Modell dafür, wie der Compiler einem Namen einen Typ und eine Adresse zuordnet. Das laufende Programm kennt nur noch Adressen.“ Der Rechner führt keine solche Tabelle und schlägt beim Lauf keinen Namen nach.
- Ein Verbund hat in der Namenstabelle eine einzige Zeile mit dem Namen, dem Typ und der Adresse des ersten Feldes.
- Der Typ der Seite ist `student` mit den Feldern `matrikel int32`, `semester int16`, `note int8` und `versuch int8`. Ein `student` belegt 4 + 2 + 1 + 1 = 8 Zellen.
- Der Bauplan des Typs: `matrikel` hat den Abstand 0 und 4 Zellen. `semester` hat den Abstand 4 und 2 Zellen. `note` hat den Abstand 6 und 1 Zelle. `versuch` hat den Abstand 7 und 1 Zelle.
- Die Rechnung für ein Feld: Adresse des Feldes = Adresse des Verbunds + Abstand des Feldes.
- Bei einem Array rechnet der Rechner die Adresse aus dem Index, weil alle Elemente gleich groß sind. Die Felder eines Verbunds sind verschieden groß, deshalb steht der Abstand jedes Feldes im Bauplan.
- Das Modell hat 32 Zellen mit den Adressen 100 bis 131. Die Adressen sind erfunden.
- Die Wörter „Zelle“, „Namenstabelle im Modell“ und „Bauplan“ gehören zum Modell der Seite. Sie stehen nicht im Kurstext. Im Kurstext steht in Abschnitt 8.2.3: Jedes Feld einer Verbund-Variable ist selbst eine Variable, und Verbund und Felder sind adressierbar.
- Das Modell vereinfacht. In Go ordnet der Compiler jedem Namen beim Übersetzen einen Typ und eine Adresse zu und benutzt dabei den Bauplan. Ein echter Compiler darf zwischen zwei Feldern Zellen frei lassen, damit jedes Feld an einer passenden Adresse beginnt. Beim Typ `student` ist das nicht nötig.
- Ein Wert aus mehreren Zellen steht im Modell so, dass die vorderen Bits in der Zelle mit der kleineren Adresse stehen.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text erklärt Zelle, Adresse und Namenstabelle im Modell, mit dem Satz, dass es diese Tabelle so nicht gibt. Er sagt auch, was ein Verbund im Speicher ist und warum der Compiler einen Bauplan anlegt.
- Eine Liste nennt drei HTML-Seiten, die dieser vorausgehen: „Verbund in Go“, „Verbund-Variablen und Feld-Variablen“ und „Speichermodell: Arrays“.
- Ein gelber Kasten sagt, dass das Modell vereinfacht und was in Abschnitt 8.2.3 des Kurstexts steht.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demo steht in Schritt 2.

### Abschnitte des Kurstexts

- 8.2
- 8.2.3

## Schritt 1: Der Bauplan des Typs

Reiter: bauplan

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text, einer Tabelle und einer Formel. Eine Demo gibt es hier nicht.
- Die Tabelle hat die Spalten „Feld“, „Typ des Feldes“, „Abstand vom Anfang in Zellen“ und „Zellen des Feldes“ und vier Zeilen für `matrikel`, `semester`, `note` und `versuch`.
- Die Formel steht in einem eigenen Kasten: Adresse des Feldes = Adresse des Verbunds + Abstand des Feldes.
- Das Beispiel darunter: Liegt `s` an der Adresse 100, dann liegt `s.semester` an 104 und `s.note` an 106.
- Die Deklaration `type student struct { … }` reserviert keine Zelle. Sie legt nur den Bauplan fest.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Tabelle ist fest.
- Mit „Weiter zu Schritt 2“ zur Demo wechseln.

### Abschnitte des Kurstexts

- 8.2

## Schritt 2: Ein Programm im Speicher

Reiter: speicher

### Aufbau der Darstellung

- Oben steht das Programm als nummerierte Liste mit neun Einträgen: die Deklaration `type student struct { … }` über sechs Zeilen, dann `var s student`, `s.matrikel = 1234567`, `s.semester = 3`, `s.note = 2`, `var t student`, `t = s`, `t.note = 1`, `fmt.Println(s, t)`.
- Der zuletzt ausgeführte Eintrag steht weiß auf dunklem Grund. Die Einträge davor sind blass.
- Darunter stehen drei Knöpfe: „nächste Zeile ausführen“, „eine Zeile zurück“ und „von vorn“.
- Links liegt der Hauptspeicher: 32 Zellen mit den Adressen 100 bis 131. In jeder Zelle steht oben klein die Adresse und darunter das Bitmuster aus acht Bits.
- Die Zellen einer Variable haben eine Farbe. Die Felder eines Verbunds sind abwechselnd hell und dunkel gefärbt.
- Eine Zelle mit blauem Rand hat die letzte Zeile gelesen. Eine Zelle mit rotem Rand hat die letzte Zeile beschrieben.
- Unter dem Hauptspeicher steht eine Zelle groß: acht Bits als Knöpfe, über jedem Bit seine Wertigkeit. Darüber steht, zu welchem Feld die Zelle gehört.
- Rechts steht die „Namenstabelle im Modell“ mit den Spalten Name, Typ, Adresse und Zellen. Darunter steht der Kasten „Was der Rechner in dieser Zeile tut“.
- Der Ablauf, Eintrag für Eintrag:
  1. Die Deklaration des Typs reserviert keine Zelle.
  2. `var s student` reserviert die Zellen 100 bis 107 und füllt sie mit Nullen.
  3. `s.matrikel = 1234567` schreibt in die Zellen 100 bis 103.
  4. `s.semester = 3` schreibt in die Zellen 104 und 105, die Adresse ist 100 + 4.
  5. `s.note = 2` schreibt in die Zelle 106, die Adresse ist 100 + 6.
  6. `var t student` reserviert die Zellen 108 bis 115.
  7. `t = s` kopiert alle 8 Zellen von `s` nach `t`.
  8. `t.note = 1` schreibt in die Zelle 114, die Adresse ist 108 + 6. `s.note` bleibt 2.
  9. `fmt.Println(s, t)` gibt `{1234567 3 2 0} {1234567 3 1 0}` aus.
- Schaltet der Student ein Bit von Hand um, ändert sich der Wert des Feldes. Die Namenstabelle bleibt gleich.

### Handgriffe

- „nächste Zeile ausführen“ wählen.
- „eine Zeile zurück“ wählen: Der Stand vor der letzten Zeile kommt zurück, auch von Hand umgeschaltete Bits.
- „von vorn“ wählen: Alle Zellen sind wieder frei, die Namenstabelle zeigt noch keine Zeile.
- Eine Zelle im Hauptspeicher anklicken: Sie erscheint unten groß.
- Eine Zeile der Namenstabelle anklicken. Der Mauszeiger über einer Zeile hebt die Zellen der Variable hervor.
- Ein Bit der groß gezeigten Zelle anklicken: Es wechselt zwischen 0 und 1.
- Das Programm lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 8.2.3
- 8.2

## Aufgaben

Reiter: aufgaben

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Die erste Karte heißt „Erst vorhersagen, dann ausprobieren“ und stellt fünf Aufgaben. Die Aufgaben 1, 2 und 5 gehören zu Schritt 1, die Aufgaben 3 und 4 gehören zu Schritt 2.
- Unter jeder Aufgabe steht eine aufklappbare „Antwort“. Sie ist verdeckt, bis der Student sie aufklappt.
- Die Antworten: Aufgabe 1: Die Deklaration des Typs reserviert keine Zelle. Aufgabe 2: `s.note` liegt an 100 + 6 = 106. Aufgabe 3: `t = s` beschreibt 8 Zellen. Aufgabe 4: In `s` steht weiter die Note 2, `t.note` liegt an 108 + 6 = 114. Aufgabe 5: Die Felder sind verschieden groß, deshalb steht der Abstand im Bauplan.
- Fragt der Student nach einer Aufgabe, deren Antwort er noch nicht aufgeklappt hat, bekommt er einen Hinweis auf den Weg und nicht gleich das Ergebnis.
- Die zweite Karte nennt die vier HTML-Seiten mit dem Speichermodell: primitive Datentypen in Kapitel 4, Arrays in Kapitel 6, Zeiger in Kapitel 7 und diese zu Kapitel 8.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Verbünde und Funktionen: Zeiger oder Kopie“.

### Handgriffe

- Eine „Antwort“ aufklappen.
- Für die Aufgaben 3 und 4 mit den Reitern oben zu Schritt 2 wechseln und dort ausprobieren.
- Den Links zu den drei anderen HTML-Seiten mit dem Speichermodell folgen.
- In diesem Schritt selbst lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2
- 8.2.3

## Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Der Bauplan des Typs, 8.2, Seiten 199–200: Der Verbund-Typ mit Feldern verschiedener Typen stammt aus dem Kurstext. Der Typ student ist ein eigenes Beispiel. Der Bauplan mit den Abständen steht nicht im Kurstext.
  - 2 Ein Programm im Speicher, 8.2, Seiten 199–200; 8.2.3, Seiten 204–205: Deklaration und Zugriff mit Punkt-Notation stammen aus dem Kurstext. Das Programm Zeile für Zeile mit Hauptspeicher und Namenstabelle gibt die Lektion dazu.
  - Aufgaben, kein eigener Abschnitt: Die Aufgaben beziehen sich auf die Abschnitte der Schritte 1 und 2. Die Aufgaben gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Datenverbund: Ohne Go steht dort das Beispiel einer Person mit Vorname, Nachname, Adresse und Telefonnummer. Die Adresse ist selbst wieder zusammengesetzt. Lektion: Variablen, Typen und Speicher, Schritt 3.
  - Kapitel 6, Seiten 151–162, Arrays im Speicher: Lektion: Speichermodell: Arrays.
  - Abschnitt 8.2.3, Seiten 204–205, Feld-Variablen: Lektion: Verbund-Variablen und Feld-Variablen in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu „Aufgaben““ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 8.2
- 8.2.3
