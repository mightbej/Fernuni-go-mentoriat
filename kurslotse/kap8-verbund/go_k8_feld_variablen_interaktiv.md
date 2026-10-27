# Verbund-Variablen und Feld-Variablen in Go

Die Seite zeigt, dass jedes Feld eines Verbunds selbst eine Variable ist: die Punkt-Notation, den Unterschied zwischen dem Wert eines Felds und einem Zeiger auf das Feld, und die Zuweisung eines ganzen Verbunds. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.2.3, Seiten 204–205.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Eine Variable eines Verbund-Typs heißt Verbund-Variable. Jedes ihrer Felder ist eine Variable vom Typ des Felds und heißt Feld-Variable.
- Der Wert einer Verbund-Variable setzt sich aus den Werten ihrer Felder zusammen.
- Der Typ der Seite ist `person` mit den Feldern `vorname string`, `nachname string`, `gebJahr int`, `ort string` und `verheiratet bool`.
- Die Verbund-Variable der Seite heißt `anna`. Ihr Wert ändert sich von Schritt zu Schritt: Schritt 1 beginnt mit `{Anna Schmidt 1966 Hamburg false}` und endet mit `{Anna Schmidt-Schiller 1966 Hamburg true}`. Schritt 2 endet mit `{Anna Schmidt-Schiller 1966 Ulm true}`.
- Jede Demo zeigt ein Programm, drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, einen Rahmen für `main` mit den Variablen, die Ausgabe und einen Kasten mit einem Satz.
- Ein Verbund steht in seinem Kasten als Liste seiner Felder mit Name und Wert. Grün ist ein Feld, das die letzte Anweisung gelesen oder geändert hat. Ein gelber Kasten ist eine neue oder gerade geänderte Variable.
- Der Wert eines Zeigers steht in Worten da: „Adresse von anna.ort“.
- In Schritt 2 steht unter dem Rahmen der „Hauptspeicher“ des Modells mit drei Blöcken: `anna` ab Zelle 100, `heimat` ab Zelle 200 und `pHeimat` in den Zellen 300 bis 307. In `pHeimat` steht die Zahl 140. Das ist im Modell die Adresse des Felds `ort`, 40 Zellen hinter dem Anfang von `anna`. Die Adressen sind erfunden. Wie viele Zellen ein `string` belegt, sagt die Seite nicht.
- `heimat` hat eigene Zellen mit einer Kopie des Werts. `pHeimat` hat keine Kopie, nur die Adresse.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und der Definition des Typs `person`. Eine Demo gibt es hier nicht.
- Eine Liste erklärt Verbund-Variable und Feld-Variable.
- Der Text vergleicht mit dem Array: Auch ein Element `a[i]` ist eine Variable.
- Unter „Was dazugehört“ stehen drei Punkte mit Links: die Lektion „Verbund in Go“, Abschnitt 6.1 mit der Lektion „Arrays in Go“ und Abschnitt 7.4 mit der Lektion „Zeiger in Go“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2.3

## Schritt 1: Punkt-Notation

Reiter: punkt

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text erklärt die Punkt-Notation: Name der Verbund-Variable, Punkt, Name des Felds. `anna.nachname` ist der Bezeichner der Feld-Variable.
- Die Demo zeigt fünf Zeilen im Rumpf von `main`: `var anna = person{"Anna", "Schmidt", 1966, "Hamburg", false}`, `anna.nachname += "-Schiller"`, `anna.verheiratet = true`, `fmt.Println(anna.nachname, anna.verheiratet)`, `fmt.Println(anna)`.
- Die Ausgabe: `Schmidt-Schiller true` und `{Anna Schmidt-Schiller 1966 Hamburg true}`.
- Unter „Was der Compiler nicht annimmt“ stehen drei Knöpfe mit je einer Zeile und der Meldung:
  1. `fmt.Println(anna.name)` meldet `anna.name undefined (type person has no field or method name)`.
  2. `anna.gebJahr = "1966"` meldet `cannot use "1966" (untyped string constant) as int value in assignment`.
  3. `fmt.Println(person.ort)` meldet `person.ort undefined (type person has no method ort)`. `person` ist der Typ und keine Variable.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Eine der drei Zeilen mit Fehler wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.3

## Schritt 2: Wert eines Felds oder Zeiger auf das Feld

Reiter: feldzeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Eine Verbund-Variable und jedes ihrer Felder sind adressierbar. `&anna.ort` liefert einen Zeiger auf das Feld.
- Eine Liste unterscheidet den Wert des Felds, `var heimat string = anna.ort`, und das Feld selbst, `var pHeimat *string = &anna.ort`.
- Die Demo zeigt acht Zeilen. Sie beginnt mit `anna` gleich `{Anna Schmidt-Schiller 1966 Hamburg true}`.
- Der Ablauf, Zeile für Zeile:
  1. `var heimat string = anna.ort`: `heimat` bekommt die Kopie `"Hamburg"`.
  2. `var pHeimat *string = &anna.ort`: `pHeimat` zeigt auf das Feld.
  3. Die Ausgabe ist `Hamburg Hamburg true`.
  4. `anna.ort = "München"` ändert das Feld.
  5. Die Ausgabe ist `Hamburg München false`. `heimat` ist die alte Kopie, `*pHeimat` ist das Feld.
  6. `*pHeimat = "Ulm"` ändert das Feld über den Zeiger.
  7. `heimat = "Kiel"` ändert nur `heimat`.
  8. `fmt.Println(anna)` gibt `{Anna Schmidt-Schiller 1966 Ulm true}` aus.
- Im Kurstext hat das Programm fünf Ausgaben, die Demo zeigt drei davon.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.3
- 7.4

## Schritt 3: Ein Verbund als Ganzes

Reiter: kopie

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Eine Zuweisung kopiert den Wert. Der Wert eines Verbunds besteht aus den Werten aller Felder, also kopiert die Zuweisung alle Felder.
- Die Demo zeigt drei Zeilen: `kopie := anna`, `kopie.ort = "Bonn"`, `fmt.Println(anna.ort, kopie.ort)`. Sie beginnt mit `anna` gleich `{Anna Schmidt-Schiller 1966 Ulm true}`.
- Nach der ersten Zeile stehen zwei Verbünde mit den gleichen Werten im Rahmen. Die zweite Zeile ändert nur `kopie`. Die Ausgabe ist `Ulm Bonn`.
- Ein Satz verweist auf den Tausch `personA, personB = personB, personA` aus Abschnitt 8.1.
- Das Programm mit `kopie` steht nicht im Kurstext. Es ist ein eigenes Beispiel zum Satz, dass sich der Wert einer Verbund-Variable aus den Werten ihrer Felder zusammensetzt.
- Am Ende des Schritts unter „Nächste Schritte“ stehen drei Links: die Lektion „Zeiger auf Verbünde“, die Seite „Speichermodell: Verbund“ und das Quiz.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.3
- 8.1

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Punkt-Notation, 8.2.3, Seite 204: Das Programm der Demo stammt aus dem Kurstext, bis auf die letzte Zeile. Die letzte Zeile und die drei Zeilen mit Meldungen des Compilers sind eine Ergänzung, geprüft mit Go 1.26.
  - 2 Wert oder Zeiger, 8.2.3, Seiten 204–205: Das Programm stammt aus dem Kurstext, ohne zwei der fünf Ausgaben. Der Wert eines Zeigers steht in der Lektion in Worten da.
  - 3 Ein Verbund als Ganzes, 8.2.3, Seiten 204–205: Dass sich der Wert einer Verbund-Variable aus den Werten ihrer Felder zusammensetzt, steht im Kurstext. Das Programm mit kopie ist ein eigenes Beispiel, geprüft mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Datenverbund: Ohne Go steht dort das Beispiel einer Person mit Vorname, Nachname, Adresse und Telefonnummer. Die Adresse ist selbst wieder zusammengesetzt. Lektion: Variablen, Typen und Speicher, Schritt 3.
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.
  - Abschnitt 6.1, Seiten 154–155, Index: Ein Array-Element spricht man mit dem Index an, ein Feld mit seinem Namen. Lektion: Arrays in Go, Schritt 2.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 6.1
- 7.3
- 7.4
- 8.2.3
