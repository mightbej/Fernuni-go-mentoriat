# Speichermodell: Name, Typ und Adresse in Go

Die Seite zeigt, dass eine Variable ein Stück Hauptspeicher mit einem Namen ist. Der Typ legt fest, wie viele Zellen zur Variable gehören und wie ihre Bits zu lesen sind. Schritt 1 führt ein Programm mit acht Zeilen aus und zeigt dabei Hauptspeicher und Namenstabelle im Modell. Schritt 2 zeigt einen Wert in vier Schreibweisen. Schritt 3 liest ein Bitmuster mit vier verschiedenen Typen. Die Seite gehört zu Abschnitt 2.2 und zu den Abschnitten 4.3.1 bis 4.3.4 des Kurstexts.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Hauptspeicher ist eine Reihe von Zellen. Jede Zelle fasst 8 Bit, also 1 Byte. Jede Zelle hat eine Nummer, ihre Adresse.
- Neben dem Hauptspeicher zeigt die Seite eine Namenstabelle im Modell. In jeder Zeile stehen ein Name, ein Typ und eine Adresse. Der Wert einer Variable steht nur im Hauptspeicher, in der Namenstabelle steht er nicht.
- Die Seite sagt dazu wörtlich: „Diese Tabelle gibt es so nicht. Sie ist ein Modell dafür, wie der Compiler einem Namen einen Typ und eine Adresse zuordnet. Das laufende Programm kennt nur noch Adressen.“ Der Rechner führt keine solche Tabelle und schlägt beim Lauf keinen Namen nach.
- Das Modell der Seite hat 32 Zellen mit den Adressen 100 bis 131. Die Adressen sind erfunden.
- Das Modell vereinfacht. In Go ordnet der Compiler jedem Namen beim Übersetzen einen Typ und eine Adresse zu und setzt die Adressen fest in das Programm ein. Ein laufendes Go-Programm schlägt keinen Namen mehr nach.
- `int8` belegt 1 Zelle, `int16` 2 Zellen, `int32` 4 Zellen, `int64` und `float64` 8 Zellen, `bool` 1 Zelle.
- Ein Wert aus mehreren Zellen steht im Modell so, dass die vorderen Bits in der Zelle mit der kleineren Adresse stehen. Ein echter Rechner mit Intel- oder AMD-Prozessor legt die Zellen in umgekehrter Reihenfolge ab. Das steht nicht auf der Seite und nicht im Kurstext.
- Über jedem Bit steht seine Wertigkeit. Bei einem Typ mit Vorzeichen hat das linke Bit eine negative Wertigkeit, bei `int8` ist sie −128.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Ob das Speichermodell in der Klausur vorkommt, lässt sich deshalb nicht sagen. Verbindlich sind der Kurstext und Moodle der FernUniversität.
- Die Wörter „Zelle“ und „Namenstabelle im Modell“ gehören zum Modell der Seite. Im Kurstext heißt es in Abschnitt 2.2 „Speicherplatz“ und „Speicherelement“.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Die erste Karte erklärt Zelle, Adresse und Namenstabelle im Modell, mit dem Satz, dass es diese Tabelle so nicht gibt. Eine nummerierte Liste nennt, was die Namenstabelle zum Namen `a` angibt: die Adresse und den Typ.
- Eine Liste nennt zwei HTML-Seiten, die dem Modell vorausgehen: „Variablen, Typen und Speicher“ zu Abschnitt 2.2 und den Exkurs „Speichern und Rechnen im Rechner“.
- Ein gelber Kasten sagt, dass das Modell vereinfacht.
- Die zweite Karte enthält eine Tabelle mit sechs Typen: `int8`, `int16`, `int32`, `int64`, `float64` und `bool`. Spalten: Typ, Bits, Zellen im Hauptspeicher und was in den Zellen steht. Die Zeilen von `int8` und `int32` sind markiert, weil das Programm in Schritt 1 diese zwei Typen benutzt.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den zwei HTML-Seiten folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 2.2
- 4.3.4

## Schritt 1: Name, Typ und Adresse

Reiter: speicher

### Aufbau der Darstellung

- Oben steht das Programm als nummerierte Liste mit acht Zeilen: `var a int8`, `var b int32`, `a = 17`, `a = a + 4`, `b = 1000`, `var c int8`, `c = a`, `fmt.Println(a, b, c)`. Die Zeilen stehen im Rumpf von `main`.
- Die zuletzt ausgeführte Zeile steht weiß auf dunklem Grund. Die Zeilen davor sind blass.
- Darunter stehen drei Knöpfe: „nächste Zeile ausführen“, „eine Zeile zurück“ und „von vorn“.
- Links liegt der Hauptspeicher: 32 Zellen mit den Adressen 100 bis 131. In jeder Zelle steht oben klein die Adresse und darunter das Bitmuster aus acht Bits.
- Eine freie Zelle ist weiß mit grauer Schrift. Die Zellen einer Variable haben eine Farbe, jede Variable eine andere. Die gleiche Farbe steht vor dem Namen in der Namenstabelle.
- Eine Zelle mit blauem Rand hat die letzte Zeile gelesen. Eine Zelle mit rotem Rand hat die letzte Zeile beschrieben.
- Unter dem Hauptspeicher steht eine Zelle groß: acht Bits als Knöpfe, über jedem Bit seine Wertigkeit. Darüber steht, zu welcher Variable die Zelle gehört. Darunter steht, wie der Typ die Zellen der Variable zusammen als Zahl liest.
- Rechts steht die „Namenstabelle im Modell“ mit den Spalten Name, Typ, Adresse und Zellen. Mit jeder Deklaration zeigt sie eine Zeile mehr. Die Demo sagt dazu, dass der Compiler diese Zuordnung beim Übersetzen festgelegt hat.
- Darunter steht der Kasten „Was der Rechner in dieser Zeile tut“ mit den Handlungen des Rechners als nummerierte Liste.
- Nach dem ganzen Programm gilt: `a` hat die Adresse 100 und den Wert 21. `b` belegt die Zellen 101 bis 104 und hat den Wert 1000. `c` hat die Adresse 105 und den Wert 21. Die Ausgabe ist `21 1000 21`.
- 1000 hat das Bitmuster 11 1110 1000, das sind 10 Bit. Die vorderen Bits 00000011 stehen in Zelle 103, die hinteren Bits 11101000 in Zelle 104. Die Zellen 101 und 102 enthalten 00000000.
- `c = a` kopiert das Bitmuster aus Zelle 100 in Zelle 105. Danach sind `a` und `c` zwei Variablen mit dem gleichen Wert. Wer in Zelle 100 ein Bit umschaltet, ändert `c` nicht.
- Eine Deklaration reserviert Zellen und füllt sie mit Nullen. Der Wert einer frisch deklarierten Variable ist deshalb 0.
- Schaltet der Student ein Bit von Hand um, ändert sich der Wert der Variable. Die Namenstabelle bleibt gleich. Die folgenden Zeilen des Programms rechnen mit dem geänderten Wert.
- Ein Beispiel dazu: Steht nach `a = 17` in Zelle 100 das Bitmuster 01111111, also 127, dann liefert `a = a + 4` den Wert −125. Die Summe 131 passt nicht in einen `int8`. Das ist der Überlauf.

### Handgriffe

- „nächste Zeile ausführen“ wählen: Der Rechner führt die nächste Zeile aus.
- „eine Zeile zurück“ wählen: Der Stand vor der letzten Zeile kommt zurück, auch von Hand umgeschaltete Bits.
- „von vorn“ wählen: Alle Zellen sind wieder frei, die Namenstabelle zeigt noch keine Zeile.
- Eine Zelle im Hauptspeicher anklicken: Sie erscheint unten groß.
- Eine Zeile der Namenstabelle anklicken: Die erste Zelle des Wertes erscheint unten groß. Der Mauszeiger über einer Zeile hebt die Zellen der Variable hervor.
- Ein Bit der groß gezeigten Zelle anklicken: Es wechselt zwischen 0 und 1.
- Das Programm lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 2.2
- 4.3.4

## Schritt 2: Ein Wert, mehrere Schreibweisen

Reiter: schreibweisen

### Aufbau der Darstellung

- Zu sehen ist eine einzelne Zelle mit acht Bits als Knöpfen. Über jedem Bit steht seine Wertigkeit von 128 bis 1.
- Darüber stehen vier Knöpfe mit Beispielen: 01111011, 01000001, 11111111 und 00001000.
- Unter den Bits stehen zwei Zeilen mit Gruppen. Vier Bits ergeben eine Hexadezimalziffer. Drei Bits ergeben eine Oktalziffer, für die linke Oktalziffer bleiben zwei Bits.
- Darunter steht eine Tabelle mit vier Zeilen: Dezimalsystem, Binärsystem, Oktalsystem und Hexadezimalsystem. Spalten: Zahlensystem, Literal im Quelltext, Formatverb für `fmt.Printf` und Ausgabe.
- Für 01111011: dezimal `123`, binär `0b01111011`, oktal `0o173`, hexadezimal `0x7b`. Die Formatverben sind `%d`, `%b`, `%o` und `%x`. Die Ausgaben sind `123`, `1111011`, `173` und `7b`.
- Die Ausgabe mit `%b` lässt führende Nullen weg. Die Ausgabe hat kein Präfix wie `0x`.
- Der Kurstext nennt in Abschnitt 4.3.3 die Literale `123` und `0x7b`: zwei Schreibweisen, ein Wert. Der Compiler rechnet jede Schreibweise in ein Bitmuster um.
- Ein gelber Kasten warnt: Auch eine führende `0` ohne Buchstaben leitet eine Oktalzahl ein. `0173` hat den Wert 123, `173` den Wert 173.
- Der Schritt endet mit einem Link auf die HTML-Seite „Zahlensysteme in Go“.

### Handgriffe

- Ein Bit anklicken: Es wechselt zwischen 0 und 1. Gruppen und Tabelle rechnen sofort mit.
- Einen der vier Knöpfe mit Beispielen wählen.
- Die Zelle ist dieselbe wie in Schritt 3. Ein Bitmuster, das hier eingestellt ist, steht auch dort.
- Ein Typ lässt sich in diesem Schritt nicht wählen. Die Zelle wird als Zahl ohne Vorzeichen von 0 bis 255 gezeigt.

### Abschnitte des Kurstexts

- 4.3.1
- 4.3.2
- 4.3.3

## Schritt 3: Ein Bitmuster, mehrere Werte

Reiter: typen

### Aufbau der Darstellung

- Zu sehen ist die Zelle aus Schritt 2 mit acht Bits als Knöpfen.
- Darüber stehen vier Knöpfe mit Beispielen: 01111011, 01000001, 11111111 und 00000001.
- Unter den Bits stehen vier Knöpfe für die Lesart: „als int8“, „als uint8“, „als Zeichen“ und „als bool“. Die gewählte Lesart ist grün.
- Darunter steht, welchen Wert das Bitmuster in dieser Lesart hat.
- Bei `int8` hat das linke Bit die Wertigkeit −128, bei `uint8` die Wertigkeit 128. Bei `bool` stehen keine Wertigkeiten über den Bits.
- 11111111 ist als `int8` der Wert −1 und als `uint8` der Wert 255. 01000001 ist als `int8` und als `uint8` der Wert 65 und als Zeichen das `A`.
- Als Zeichen gelesen ist ein Bitmuster von 32 bis 126 ein sichtbares Zeichen. Unter 32 ist es ein Steuerzeichen. Ab 128 steht das Byte in einem Go-String nicht allein für ein Zeichen, es ist Teil eines Zeichens aus mehreren Bytes.
- Der Typ `bool` kennt nur 00000000 für `false` und 00000001 für `true`. Bei jedem anderen Bitmuster sagt die Seite, dass es in einer Variable vom Typ `bool` nicht vorkommt.
- Eine Tabelle zeigt 01000001 und 11111111 in den drei Lesarten `int8`, `uint8` und Zeichen.
- Ein gelber Kasten nennt den Bezug: In Abschnitt 4.3.3 wird die syntaktische von der semantischen Ebene unterschieden. In Abschnitt 2.2 steht, dass der Typ bestimmt, wie eine Folge von 0 und 1 als Wert gelesen wird.
- In einem Go-Programm legt die Deklaration den Typ einer Variable fest, und er bleibt. Nur auf dieser Seite lässt er sich für eine Zelle wechseln. Am nächsten kommt dem in Go die Konvertierung, zum Beispiel `int8(u)`. Dazu gibt es einen Link auf die HTML-Seite „Typkonvertierung“.

### Handgriffe

- Ein Bit anklicken: Es wechselt zwischen 0 und 1.
- Einen der vier Knöpfe mit Beispielen wählen.
- Eine der vier Lesarten wählen. Die Bits bleiben stehen, der Wert ändert sich.

### Abschnitte des Kurstexts

- 2.2
- 4.3.3
- 4.3.4

## Aufgaben

Reiter: aufgaben

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Die erste Karte heißt „Erst vorhersagen, dann ausprobieren“ und stellt acht Aufgaben. Vor jeder Aufgabe steht ein Link „In Schritt 1 lösen“, „In Schritt 2 lösen“ oder „In Schritt 3 lösen“, der diesen Schritt öffnet; die Demo behält dabei ihren Stand. Die Aufgaben 1 bis 5 gehören zu Schritt 1, Aufgabe 6 gehört zu Schritt 2, die Aufgaben 7 und 8 gehören zu Schritt 3.
- Unter jeder Aufgabe steht eine aufklappbare „Antwort“. Sie ist verdeckt, bis der Student sie aufklappt.
- Die Antworten: Aufgabe 1: 5 Zellen. Aufgabe 2: 23. Aufgabe 3: 1000 braucht 10 Bit. Aufgabe 4: Nein, `c` hat eine eigene Zelle. Aufgabe 5: −125. Aufgabe 6: `0xff` und `0o377`. Aufgabe 7: −1 und 255. Aufgabe 8: 01100001, das ist die Nummer 97.
- Fragt der Student nach einer Aufgabe, deren Antwort er noch nicht aufgeklappt hat, bekommt er einen Hinweis auf den Weg und nicht gleich das Ergebnis.
- Die zweite Karte nennt die vier HTML-Seiten mit dem Speichermodell: diese zu Kapitel 4, dann Arrays in Kapitel 6, Zeiger in Kapitel 7 und Verbund in Kapitel 8.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Zweierkomplement: negative Zahlen als Bitmuster in Go“.
  - Quiz: „Selbsttest: Speichermodell“ mit 12 Fragen zu dieser Lektion und „Grundlagen der primitiven Datentypen“.

### Handgriffe

- Eine „Antwort“ aufklappen.
- Für eine Aufgabe mit den Reitern oben zu Schritt 1, 2 oder 3 wechseln und dort ausprobieren.
- Den Links zu den drei anderen HTML-Seiten mit dem Speichermodell folgen.
- In diesem Schritt selbst lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 4.3.1
- 4.3.4

## Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.2, Seiten 8–10: Speicherplatz; 4.3.4, Seite 63: Ganzzahl-Typen: Dass eine Variable Speicherplatz aus Speicherelementen belegt, steht in Abschnitt 2.2. Die Tabelle der Ganzzahl-Typen stammt aus Abschnitt 4.3.4. Zellen mit Adressen und die Namenstabelle im Modell mit Name, Typ und Adresse gibt die Lektion dazu.
  - 1 Name, Typ und Adresse, 2.2, Seiten 8–10; 4.3.4, Seite 63: Die Typen und ihre Größe in Bit stammen aus dem Kurstext. Der Hauptspeicher mit den Zellen 100 bis 131 und das Programm Zeile für Zeile gibt die Lektion dazu.
  - 2 Ein Wert, mehrere Schreibweisen, 4.3.1 bis 4.3.3, Seiten 59–63: Literale: Literale in vier Zahlensystemen und der Unterschied zwischen Literal und Wert stammen aus dem Kurstext. Die Demo gibt die Lektion dazu.
  - 3 Ein Bitmuster, mehrere Werte, 2.2, Seiten 8–10; 4.3.5, Seiten 64–66: Zweierkomplement: Dass der Typ festlegt, wie die Bits gelesen werden, steht in Abschnitt 2.2. Das Lesen derselben Zelle als int8, uint8, Zeichen und bool gibt die Lektion dazu.
  - Aufgaben, kein eigener Abschnitt: Die Aufgaben beziehen sich auf die Abschnitte der Schritte 1 bis 3. Die Aufgaben gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Variable als Speicherplatz: Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 3.3.3, Seiten 29–30, Datentyp: Der Typ einer Variable legt ihren Wertebereich fest. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 4.3.5, Seiten 64–66, Zweierkomplement: Lektion: Zweierkomplement.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu „Aufgaben““ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 3.3.3
- 4.3.1
- 4.3.3
- 4.3.4
- 4.3.5
