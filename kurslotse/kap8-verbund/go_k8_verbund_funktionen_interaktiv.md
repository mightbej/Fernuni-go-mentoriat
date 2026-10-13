# Verbünde und Funktionen: Zeiger oder Kopie

Die Seite zeigt Verbünde an Funktionen: einen Verbund-Zeiger als Rückgabe, was auf der linken Seite einer Zuweisung stehen darf, einen Zeiger als Parameter und eine Kopie als Parameter. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.2.6 und 8.2.7, Seiten 208–211.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- In Go gilt call-by-value: Der Parameter bekommt eine Kopie des Arguments. Hat der Parameter einen Verbund-Typ, ist das eine Kopie aller Feld-Werte. Hat er einen Zeiger-Typ, ist es die Kopie einer Adresse.
- Beim Zugriff auf ein Feld dereferenziert Go einen Verbund-Zeiger von selbst: `p.gebJahr` steht für `(*p).gebJahr`. Das gilt auch für einen Zeiger, den ein Funktionsaufruf liefert.
- Einen Wert zuweisen kann man nur einer Variable. Ein dereferenzierter Zeiger ist eine Variable.
- Die Typen der Seite: `person` mit `vorname`, `nachname`, `gebJahr`, `ort`, `verheiratet` und `auto` mit `marke string`, `kennzeichen string`, `tachostand int`, `tüvJahr int`. Die Namen mit Umlaut stehen so im Kurstext.
- Jede Demo zeigt ein Programm, drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und einen Kasten mit einem Satz. Jede laufende Funktion hat einen Rahmen. Ein gestrichelter Rahmen gehört zu einer Funktion, die beendet ist.
- Ein Verbund steht in seinem Kasten als Liste seiner Felder. Der Wert eines Zeigers steht in Worten da, zum Beispiel „Adresse von golf“.
- Unter den Rahmen einer Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger. In ihm steht eine Zahl, die Adresse einer Zelle, und darunter „Adresse der Zelle …“. Ein Block mit gelbem Schein ist neu oder gerade geändert. Die Adressen sind erfunden. Ein `int` belegt im Modell 8 Zellen, ein Zeiger auch. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke der Demos: In Schritt 1 beginnen `otto1` bei Zelle 100 und `otto2` bei Zelle 200, `p1` und `p2` liegen in den Zellen 300 bis 307 und 308 bis 315 und enthalten 100 und 200. In Schritt 3 beginnt `golf` bei Zelle 100, der Parameter `a` liegt in den Zellen 200 bis 207 und enthält 100. In Schritt 4 beginnt `golf` bei Zelle 100, und der Parameter `a` ist ein zweiter Verbund ab Zelle 200, die Kopie.
- Bei einem Verbund steht „ab Zelle …“ ohne Ende. Wie viele Zellen ein `string` belegt, sagt die Seite nicht.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Liste stellt die zwei Möglichkeiten nebeneinander: Parameter vom Verbund-Typ `a auto` oder vom Zeiger-Typ `a *auto`.
- Unter „Was dazugehört“ stehen drei Punkte mit Links: die Lektion „Funktionen in Go“, die Lektion „Zeiger in Go“ und die Lektion „Zeiger auf Verbünde“.
- Unter „Die Schritte“ steht je ein Satz zu den vier Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2.7
- 8.2.6

## Schritt 1: Ein Verbund-Zeiger als Rückgabe

Reiter: rueckgabe

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt 15 Zeilen: die Funktion `werIstÄlter(p1, p2 *person) *person` in den Zeilen 1 bis 7, darunter `main` mit `otto1` und `otto2`, der Zeile `werIstÄlter(&otto1, &otto2).nachname += " senior"` und zwei Ausgaben.
- `otto1` ist `{Otto Meier 1895 Bonn false}`, `otto2` ist `{Otto Meier 1930 Bonn false}`.
- Der Ablauf hat sieben Schritte: zwei Verbünde anlegen. Aufruf: `p1` und `p2` bekommen die Adressen. Die Bedingung `1895 < 1930` ist erfüllt, die Funktion gibt `p1` zurück, die Adresse von `otto1`. Zurück in `main` hängt die Zuweisung „ senior“ an `otto1.nachname` an. Die Ausgaben sind `{Otto Meier senior 1895 Bonn false}` und `{Otto Meier 1930 Bonn false}`.
- Die Funktion gibt den Zeiger auf den Verbund mit dem kleineren Geburtsjahr zurück. Sind beide gleich alt, gibt sie `p2` zurück.
- Darunter steht die lange Schreibweise: `(*werIstÄlter(&otto1, &otto2)).nachname += " senior"`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.6

## Schritt 2: Was links vom Gleichheitszeichen stehen darf

Reiter: links

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste nennt, was eine Variable ist: ein Bezeichner, ein dereferenzierter Zeiger, ein dereferenzierter Rückgabewert.
- Vier Knöpfe zeigen je eine Zeile. Vor jeder Zeile gilt der Anfang von Schritt 1. Darunter steht, was `fmt.Println(otto1)` nach der Zeile ausgibt, oder die Meldung des Compilers.
- Die vier Zeilen und ihr Ergebnis:
  1. `werIstÄlter(&otto1, &otto2).nachname += " senior"`: `{Otto Meier senior 1895 Bonn false}`.
  2. `(*werIstÄlter(&otto1, &otto2)).nachname += " senior"`: das Gleiche.
  3. `*werIstÄlter(&otto1, &otto2) = person{}`: `{  0  false}`. Die Zuweisung ersetzt alle Felder von `otto1` durch Null-Werte. Der Stern ist nötig, weil kein Feld-Zugriff folgt.
  4. `werIstÄlter(&otto1, &otto2) = &person{}`: Go übersetzt die Zeile nicht. Die Meldung: `cannot assign to werIstÄlter(&otto1, &otto2) (neither addressable nor a map index expression)`. Links steht ein Zeiger-Wert und keine Variable.
- Das Wort „map“ in der Meldung meint einen Datentyp von Go, der im Kurs nicht vorkommt.

### Handgriffe

- Eine der vier Zeilen wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.6

## Schritt 3: Ein Zeiger als Parameter

Reiter: zeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt den Typ `auto`.
- Die Demo zeigt zehn Zeilen: die Funktion `tüvErneuern(a *auto)` mit `a.tüvJahr += 2`, darunter `main` mit `golf := auto{marke: "VW", kennzeichen: "M-FG574", tachostand: 12000, tüvJahr: 2024}`, einer Ausgabe, dem Aufruf `tüvErneuern(&golf)` und einer zweiten Ausgabe.
- Der Ablauf hat sechs Schritte: `golf` anlegen. Ausgabe `TÜV vorher:  2024`. Aufruf: `a` bekommt die Adresse von `golf`, kein Feld wird kopiert. `a.tüvJahr += 2` erhöht das Feld von `golf` auf 2026. Die Funktion endet ohne Rückgabe. Ausgabe `TÜV nachher:  2026`.
- In der Ausgabe stehen zwei Leerzeichen vor der Zahl: eines aus dem String, eines setzt `fmt.Println` zwischen zwei Argumente.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.7
- 7.4.1

## Schritt 4: Eine Kopie als Parameter

Reiter: kopie

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt zehn Zeilen: die Funktion `tüvErneuernWert(a auto)` mit dem gleichen Rumpf, darunter das gleiche `main` mit dem Aufruf `tüvErneuernWert(golf)`.
- Der Ablauf hat sechs Schritte: `golf` anlegen. Ausgabe `TÜV vorher:  2024`. Aufruf: `a` ist eine neue Verbund-Variable mit Kopien aller vier Feld-Werte. `a.tüvJahr += 2` erhöht das Feld der Kopie auf 2026. Die Funktion endet, die Kopie gibt es nicht mehr. Ausgabe `TÜV nachher:  2024`.
- Unter „Wert oder Zeiger: Wann nimmt man was?“ stehen drei Fälle aus dem Kurstext: Soll die Funktion den Verbund ändern, braucht sie einen Zeiger. Liest sie nur und ist der Verbund groß, ist ein Zeiger oft auch sinnvoll, weil die Kopie Speicherplatz und Rechenzeit kostet. Soll die Funktion den Verbund auf keinen Fall ändern können, nimmt man die Kopie.
- Unter „Wie es weitergeht“ stehen drei Links: die Lektion „Hierarchische Verbünde“, der Exkurs „Das Kopier-Problem“ und das Quiz.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.7
- 7.2
