# Vergleich von Verbünden in Go

Die Seite zeigt, wann zwei Verbünde gleich sind, und unterscheidet den Vergleich der Werte vom Vergleich der Adressen. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.2.5, Seite 208.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Datentyp heißt vergleichbar, englisch comparable, wenn `==` und `!=` für ihn definiert sind. Alle primitiven Datentypen sind vergleichbar.
- Ein Verbund-Typ ist vergleichbar, wenn alle seine Feld-Typen vergleichbar sind.
- Zwei Verbund-Werte sind genau dann gleich, wenn alle ihre Feld-Werte gleich sind. Go vergleicht die Felder, die einander entsprechen.
- Der Typ der Seite ist `koordinate` mit `x, y, z int`. `a == b` bedeutet `a.x == b.x && a.y == b.y && a.z == b.z`.
- `a == b` fragt, ob zwei Variablen den gleichen Wert haben. `&a == &b` fragt, ob es dieselbe Variable ist.
- Für Verbünde gibt es nur `==` und `!=`. Die Operatoren `<`, `>`, `<=` und `>=` gibt es für Verbünde nicht.
- Sortieren von Verbünden: Im Kurstext sortiert Abschnitt 8.4.2 ein Array von Verbünden mit Bubblesort. Das Programm vergleicht dabei ein Feld von zwei Elementen des Arrays mit `<` oder `>` und tauscht die ganzen Verbünde. Der Kurs arbeitet mit Arrays und mit selbst geschriebenen Schleifen. Slices und das Paket `sort` gehören nicht zum Kurs, sie werden hier nicht empfohlen.
- Dass es für Verbünde kein `<` gibt, steht nicht in Abschnitt 8.2.5 des Kurstexts. Dort stehen nur `==` und `!=`. Die Meldung des Compilers zu `a < b` ist eine Ergänzung der Seite.
- Welche Typen nicht vergleichbar sind, nennt der Kurstext in Abschnitt 8.2.5 nicht. Alle Typen, die im Kurs als Feld vorkommen, sind vergleichbar.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und der Definition des Typs `koordinate`. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Regeln: wann ein Verbund-Typ vergleichbar ist und wann zwei Verbund-Werte gleich sind.
- Unter „Was dazugehört“ stehen drei Punkte mit Links: Abschnitt 4.5.2 mit der Lektion „Logische Operatoren und Vergleiche“, Abschnitt 7.4 mit der Lektion „Zeiger in Go“ und Abschnitt 8.2.1 mit der Lektion „Verbund in Go“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2.5

## Schritt 1: Feld für Feld

Reiter: felder

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt sechs Zeilen im Rumpf von `main`: `a := koordinate{5, -7, 8}`, `b := koordinate{0, -7, 8}`, `fmt.Println(a == b)`, `b.x = 3 + 2`, `fmt.Println(a == b)`, `fmt.Println(&a == &b)`.
- Unter dem Programm stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, ein Rahmen für `main` mit den zwei Verbünden, die Ausgabe und ein Kasten mit einem Satz.
- Ein Verbund steht in seinem Kasten als Liste seiner Felder. Grün sind die Felder, die die letzte Anweisung verglichen oder geändert hat.
- Der Ablauf: Der erste Vergleich gibt `false` aus, weil `a.x` 5 ist und `b.x` 0. Nach `b.x = 3 + 2` gibt der zweite Vergleich `true` aus. `&a == &b` gibt `false` aus.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.5

## Schritt 2: Gleicher Wert oder dieselbe Variable

Reiter: adresse

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste stellt die zwei Fragen nebeneinander: gleicher Wert mit `a == b`, dieselbe Variable mit `&a == &b`.
- Für die Auswahl gilt: `a := koordinate{5, -7, 8}`, `b := koordinate{5, -7, 8}`, `f := rgb{5, -7, 8}`.
- Sieben Knöpfe zeigen je einen Vergleich. Darunter stehen die Ausgabe oder die Meldung des Compilers und ein Kasten mit einem Satz.
- Die sieben Vergleiche und ihr Ergebnis:
  1. `a == b` ist `true`.
  2. `a != b` ist `false`.
  3. `&a == &b` ist `false`.
  4. `a == koordinate{5, -7, 8}` ist `true`.
  5. `a == f` meldet `invalid operation: a == f (mismatched types koordinate and rgb)`.
  6. `a < b` meldet `invalid operation: a < b (operator < not defined on struct)`.
  7. `a == nil` meldet `invalid operation: a == nil (mismatched types koordinate and untyped nil)`.

### Handgriffe

- Einen der sieben Vergleiche wählen.
- Die „Auflösung“ aufklappen.
- Ein eigener Vergleich lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 8.2.5
- 7.4

## Schritt 3: Der Vergleich im Speicher

Reiter: speicher

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Im Hauptspeicher ist ein Verbund eine Reihe von Zellen, in denen die Felder hintereinander liegen. Beim Vergleich `a == b` liest der Rechner die Zellen jedes Felds von `a` und die Zellen des gleichen Felds von `b` und vergleicht die Bitmuster.
- Im Modell haben die Felder von `koordinate` den Typ `int8`, damit ein Feld in eine Zelle passt. `a` liegt in den Zellen 100 bis 102, `b` in den Zellen 103 bis 105.
- Zu sehen sind sechs Zellen in zwei Reihen, oben `a`, unten `b`. In jeder Zelle stehen die Adresse, acht Bits als Knöpfe und darunter das Feld mit seinem Wert, zum Beispiel `a.y = -7`.
- Am Anfang steht in den Zellen: `a` ist 00000101, 11111001, 00001000, also 5, −7 und 8. `b` ist 00000000, 11111001, 00001000, also 0, −7 und 8. 11111001 ist −7 im Zweierkomplement.
- Eine Zelle mit rotem Rand gehört zu einem Feld, dessen Bitmuster in `a` und `b` verschieden ist.
- Der Kasten darunter nennt für jedes Feld die zwei Bitmuster und ob sie gleich sind, dann das Ergebnis von `a == b` und von `&a == &b`.
- `&a` ist im Modell 100, `&b` ist 103. `&a == &b` vergleicht diese zwei Zahlen und liefert immer `false`.
- Ein gelber Kasten sagt, wo das Modell vereinfacht: Für Ganzzahlen gilt genau, dass ein gleicher Wert ein gleiches Bitmuster hat. Go vergleicht Feld für Feld nach den Regeln des Feld-Typs. Bei einem Feld vom Typ `string` vergleicht Go die Zeichen der zwei Strings und nicht die Zellen der Variable.
- Auch bei Gleitkommazahlen stimmt die Gleichung von Wert und Bitmuster nicht immer: +0 und −0 haben verschiedene Bitmuster und sind gleich. Das steht nicht auf der Seite.
- Die Speichersicht steht nicht im Kurstext. Sie ist eine Ergänzung des Mentors.

### Handgriffe

- Ein Bit in einer der sechs Zellen anklicken: Es wechselt zwischen 0 und 1, und die Seite vergleicht neu.
- „b.x = 3 + 2 ausführen“ wählen: In Zelle 103 steht dann 00000101.
- „von vorn“ wählen: `a` ist wieder `{5, -7, 8}`, `b` ist wieder `{0, -7, 8}`.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.5
- 4.3.5

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Feld für Feld, 8.2.5, Seite 208: Das Programm stammt aus dem Kurstext. Die Demo gibt die Lektion dazu.
  - 2 Wert oder Variable, 8.2.5, Seite 208: Die Vergleiche 1 und 3 stammen aus dem Kurstext. Die anderen fünf Vergleiche und die Meldungen des Compilers sind eine Ergänzung, geprüft mit Go 1.26.
  - 3 Im Speicher, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Speichersicht steht nicht im Kurstext. Das Modell zeigt die Lektion „Speichermodell: Verbund“ ausführlich.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 4.5.2, Seiten 96–98, Vergleichsoperatoren: == und != für primitive Typen. Lektion: Logische Operatoren und Vergleiche in Go.
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.
- Unter „Wie es weitergeht“ stehen drei Punkte: die Lektion „Verbünde und Funktionen“, ein Satz zum Sortieren mit dem Vergleich eines Felds nach Abschnitt 8.4 und das Quiz.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.5.2
- 7.3
- 7.4
- 8.2.5
