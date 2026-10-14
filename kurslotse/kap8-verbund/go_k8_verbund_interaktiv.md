# Verbund in Go: der Datentyp struct

Die Seite führt den Verbund ein: wozu er dient, wie ein Verbund-Typ mit `struct` und `type` deklariert wird, was sein Null-Wert ist, wie ein Verbund-Literal aussieht und wie Go einen Verbund ausgibt. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.1 bis 8.2.2, Seiten 195–203.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Verbund fasst Werte zusammen, die zusammengehören und verschiedene Typen haben dürfen. Andere Namen: Datensatz, englisch Record, in Go Struct. Im Kurstext ist der Plural „Verbünde“.
- Die Teile eines Verbunds heißen Felder. Jedes Feld hat einen Namen und einen Typ.
- Der Typ der Seite ist `person` mit den Feldern `vorname string`, `nachname string`, `gebJahr int`, `ort string` und `verheiratet bool`, in dieser Reihenfolge.
- Auf ein Feld greift ein Programm mit einem Punkt zu, zum Beispiel `anna.ort`. Das steht ausführlich in der Lektion „Verbund-Variablen und Feld-Variablen“.
- Zeiger auf Verbünde, der Vergleich von Verbünden und Verbünde als Parameter stehen in eigenen Lektionen.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Meldungen sind mit Go 1.26 geprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text stellt Array und Verbund nebeneinander: Ein Array fasst Werte des gleichen Typs zusammen, ein Verbund Werte verschiedener Typen.
- Unter „Was aus früheren Kapiteln dazugehört“ stehen drei Punkte mit Links: Abschnitt 2.2, Abschnitt 5.1 und Kapitel 6.
- Unter „Die Schritte“ steht je ein Satz zu den vier Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.1
- 8.2

## Schritt 1: Wozu ein Verbund

Reiter: wozu

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt das Beispiel des Kurstexts, die Lochkarte mit 80 Spalten und festen Bereichen für Personalnummer, Nachname, Vorname, Geburtsdatum und Gehalt.
- Eine Liste nennt zwei Wege ohne Verbund und ihren Mangel: einzelne Variablen wie `nachnameA` und `nachnameB`, oder ein Array je Angabe. Was zusammengehört, steht dann nur im Namen oder am gleichen Index.
- Die Demo zeigt zwölf Zeilen: einen Typ `person` mit drei Feldern `personalnummer`, `vorname`, `nachname`, darunter `main` mit `personA := person{4711, "Donald", "Duck"}`, `personB := person{1313, "Gundel", "Gaukeley"}`, `personA, personB = personB, personA` und einer Ausgabe der zwei Nachnamen.
- Unter dem Programm stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, ein Rahmen für `main` mit den Variablen, die Ausgabe und ein Kasten mit einem Satz.
- Ein Verbund steht in seinem Kasten als Liste seiner Felder mit Name und Wert. Grün ist ein Feld, das die letzte Anweisung gelesen oder geändert hat.
- Der Ablauf hat vier Schritte. Nach dem Tausch hat `personA` die Werte von Gundel Gaukeley. Die Ausgabe ist `Gaukeley Duck`.
- Der Typ `person` dieser Demo hat drei Felder und ist ein anderer als der Typ `person` der Schritte 2 bis 4.
- Unter „Viele Verbünde“ steht die Schleife des Kurstexts über `personArr`, die die Personalnummer von Gundel Gaukeley sucht. Dort sind die Feldnamen groß geschrieben.
- Im Kurstext steht am Ende von Abschnitt 8.1 außerdem: Die Größe eines Arrays muss bei der Deklaration feststehen, für eine wechselnde Zahl von Mitarbeitern ist das unpraktisch. Eine Lösung dafür folgt in Abschnitt 9.2.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.1

## Schritt 2: Der Verbund-Typ

Reiter: typ

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt die Definition mit `struct`: je Zeile ein Feld, erst der Name, dann der Typ, ohne Komma.
- Unter „Ein Name für den Typ“ steht `type person struct { … }`. Das deklariert einen benannten Datentyp. Der Typname steht danach überall, wo auch `int` stehen kann: `var ich person`, `var großeltern [4]person`, `func jünger(p1, p2 person) person`.
- Unter „Vier Schreibweisen im Vergleich“ stehen vier Knöpfe mit je einer Definition. Ein Kasten sagt, ob sie den gleichen Datentyp beschreibt wie die Definition oben.
- Die vier Definitionen: 1 alle Felder in einer Zeile mit Semikolons, gleicher Typ. 2 `vorname, nachname string` zusammengefasst, gleicher Typ. 3 andere Reihenfolge der Felder, anderer Typ. 4 ein Feld heißt `name` und nicht `nachname`, anderer Typ.
- Zum Typ gehören die Namen der Felder, ihre Typen und ihre Reihenfolge.
- Unter „Groß oder klein geschrieben“ steht die Regel aus Abschnitt 3.3.4: Nur ein Bezeichner mit großem Anfangsbuchstaben wird exportiert.
- Im Kurstext steht in einer Fußnote: Ein Verbund-Typ lässt sich auch ohne Typnamen verwenden, die Definition muss dann an jeder Stelle wiederholt werden.

### Handgriffe

- Eine der vier Definitionen wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2

## Schritt 3: Null-Wert und Literale

Reiter: literale

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit zwei Fragen und aufklappbaren „Auflösungen“.
- Der Null-Wert eines Verbund-Typs besteht aus den Null-Werten der Felder, bei `person` aus `""`, `""`, `0`, `""` und `false`.
- Ein Verbund-Literal hat zwei Schreibweisen. Ohne Feldnamen: alle Werte in der Reihenfolge der Definition, keiner darf fehlen. Mit Feldnamen: Feldname, Doppelpunkt, Wert, in beliebiger Reihenfolge, fehlende Felder bekommen den Null-Wert.
- Steht das Literal über mehrere Zeilen, braucht auch die letzte Angabe ein Komma.
- Die Demo zeigt neun Zeilen als Knöpfe. Darunter stehen die Ausgabe von `fmt.Println` für die Variable oder die Meldung des Compilers und ein Kasten mit einem Satz. Bei einer Zeile, die Go nicht übersetzt, ist der Kasten rot.
- Die neun Zeilen und ihr Ergebnis:
  1. `var anna person` gibt `{  0  false}` aus.
  2. `anna := person{"Anna", "Schmidt", 1966, "Hamburg", false}` gibt `{Anna Schmidt 1966 Hamburg false}` aus.
  3. `karl := person{nachname: "Müller", vorname: "Karl", ort: "Köln", gebJahr: 1955}` gibt `{Karl Müller 1955 Köln false}` aus.
  4. `eva := person{vorname: "Eva"}` gibt `{Eva  0  false}` aus.
  5. `leer := person{}` gibt `{  0  false}` aus.
  6. Vier Werte ohne Feldnamen: `too few values in struct literal of type person`.
  7. Ort und Geburtsjahr vertauscht: `cannot use "Hamburg" (untyped string constant) as int value in struct literal`.
  8. Beide Schreibweisen gemischt: `mixture of field:value and value elements in struct literal`.
  9. Ein Feld `name`, das es nicht gibt: `unknown field name in struct literal of type person`.
- Ein gelber Kasten gibt Hinweis 8.1 wieder: Die Schreibweise mit Feldnamen ist in der Praxis üblich, weil sie von der Reihenfolge der Felder unabhängig ist. Die kurze passt zu kleinen Typen mit fester Reihenfolge wie `koordinate{10, -7, 22}` und `rgb{255, 255, 0}`.

### Handgriffe

- Eine der neun Zeilen wählen.
- Eine „Auflösung“ aufklappen.
- Eine eigene Zeile lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 8.2.1
- 8.2.2

## Schritt 4: Ausgabe

Reiter: ausgabe

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Variable der Demo ist `anna := person{"Anna", "Schmidt", 1966, "Hamburg", false}`.
- Vier Knöpfe zeigen je eine Zeile. Darunter steht die Ausgabe.
- `fmt.Println(anna)` und `fmt.Printf("%v\n", anna)` geben `{Anna Schmidt 1966 Hamburg false}` aus.
- `fmt.Printf("%+v\n", anna)` gibt `{vorname:Anna nachname:Schmidt gebJahr:1966 ort:Hamburg verheiratet:false}` aus.
- `fmt.Printf("%#v\n", anna)` gibt `main.person{vorname:"Anna", nachname:"Schmidt", gebJahr:1966, ort:"Hamburg", verheiratet:false}` aus.

### Handgriffe

- Eine der vier Zeilen wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Wozu ein Verbund, 8.1, Seiten 195–198: Die Schleife und der Tausch personA, personB = personB, personA stammen aus dem Kurstext. Dort sind die Feldnamen in Abschnitt 8.1 groß geschrieben. Das Programm der Demo ist ein eigenes Beispiel zu diesem Tausch.
  - 2 Verbund-Typ, 8.2, Seiten 199–200: Alle Definitionen stammen aus dem Kurstext. Die Demo gibt die Lektion dazu.
  - 3 Null-Wert und Literale, 8.2.1, Seiten 201–202; Hinweis 8.1, Seite 203: Die Zeilen 1 bis 3 stammen aus dem Kurstext. Die Zeilen 4 bis 9 und die Meldungen des Compilers sind eine Ergänzung, geprüft mit Go 1.26.
  - 4 Ausgabe, 8.2.2, Seiten 202–203: Die Zeilen und ihre Ausgaben stammen aus dem Kurstext. Die Demo gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Datenverbund: Ohne Go steht dort das Beispiel einer Person mit Vorname, Nachname, Adresse und Telefonnummer. Die Adresse ist selbst wieder zusammengesetzt. Lektion: Variablen, Typen und Speicher, Schritt 3.
  - Abschnitt 5.1, Seiten 113–114, Deklaration: Ein Verbund-Typ steht in einer Deklaration an der Stelle von int. Lektion: Variablendeklaration in Go.
  - Kapitel 6, Seiten 151–154, Array-Literal: Das Verbund-Literal ähnelt dem Array-Literal. Lektion: Arrays in Go, Schritt 1.
- Unter „Wie es weitergeht“ stehen drei Links: die Lektion „Verbund-Variablen und Feld-Variablen“, die Seite „Speichermodell: Verbund“ und das Quiz zu struct.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 5.1
- 8.1
- 8.2
- 8.2.1
- 8.2.2
