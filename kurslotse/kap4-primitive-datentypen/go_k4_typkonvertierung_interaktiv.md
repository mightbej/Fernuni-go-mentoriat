# Typkonvertierung: implizite Typanpassung und explizite Typkonvertierung in Go

Die Seite zeigt, wie in Go zwei Operanden zum selben Typ kommen. Go passt den Typ nur bei konstanten Ausdrücken selbst an, das ist die implizite Typanpassung. Bei einer Variablen schreibt das Programm die Umwandlung mit `T(a)` hin, das ist die explizite Typkonvertierung. Die Seite rechnet nur mit `int` und `float64`.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt. Die Variable a gilt in Schritt 2 und in Schritt 4 gemeinsam, die Variable f in Schritt 2 und in Schritt 3: Wer den Wert in einem Schritt ändert, sieht ihn im anderen wieder.

Den Code auf der Seite kann der Student nicht ändern und nicht ausführen. Die Seite schreibt die Programmzeilen selbst und passt sie an, wenn der Student einen Wert ändert. Einen eigenen Ausdruck kann der Student nicht eingeben.

Zahlen mit Nachkommastellen stehen auf der Seite in der Schreibweise von Go, also mit Punkt: 1.7 und nicht 1,7. Ein `float64`-Wert ohne Nachkommastellen steht im Text als 25.0. `fmt.Println` gibt ihn als 25 aus.

Die Farben gelten für alle Schritte: Blau gehört zum Typ `int` und zu ganzzahligen Ausdrücken, Grün zum Typ `float64` und zu Floating-Point-Ausdrücken, Rot zu einem Fehler.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Der Schritt hat keine Demo.
- Oben steht ein Programm mit drei Zeilen, das Go nicht übersetzt: `f * a` mit `a` vom Typ `int` und `f` vom Typ `float64`. Darunter steht die Meldung des Compilers.
- Eine Tabelle stellt die zwei Wege gegenüber: implizite Typanpassung und explizite Typkonvertierung.
- Am Ende steht eine Liste der vier Schritte.

### Handgriffe

- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 2.4
- 4.1
- 4.7

## Schritt 1: Implizite Typanpassung

Reiter: anpassung

### Aufbau der Darstellung

- Oben stehen die zwei Fälle, in denen Go den Typ selbst anpasst, als Liste.
- Darunter steht ein Programm mit zwölf Zeilen aus dem Kurstext. Es ist fest und ändert sich nicht.
- Unter dem Programm steht eine Reihe von elf Knöpfen. Jeder Knopf trägt eine Zeile oder einen Ausdruck aus dem Programm. Der gewählte Knopf ist grün.
- Der letzte Knopf, `var k int = 2.5`, ist eine Gegenprobe der Seite. Diese Zeile steht nicht im Kurstext, und Go übersetzt sie nicht.
- Die Darstellung zerlegt die gewählte Zeile in zwei Kästen. Der linke Kasten zeigt, was links vom Operator oder vom Gleichheitszeichen steht, der rechte, was rechts steht. Dazwischen steht der Operator oder das Gleichheitszeichen.
- In jedem Kasten steht oben der Ausdruck und darunter seine Art, zum Beispiel „Integer-Literal“ oder „Variable vom Typ float64“.
- Ein blauer Kasten ist ein Ausdruck mit ganzzahligem Typ, ein grüner ein Ausdruck mit Floating-Point-Typ.
- Gestrichelt umrandet ist der Kasten, dessen Typ Go anpasst. In ihm steht dann „Go rechnet mit …“. Ist kein Kasten gestrichelt, passt Go nichts an.
- Unter den Kästen steht die Rechnung mit dem Ergebnis, bei der Gegenprobe der Satz „Go übersetzt die Zeile nicht.“
- Darunter steht die Erklärung in Worten.

### Handgriffe

- Einen der elf Knöpfe anklicken und damit die Zeile wählen. Etwas anderes lässt sich in diesem Schritt nicht einstellen.
- Die elf Knöpfe heißen: `var g float64 = 9`, `g = g / 2`, `3 * 10 / g`, `6 * 0.25`, `15.0 / 2.0`, `15.0 / 2`, `15 / 2.0`, `15 / 2`, `var i int = 5.0`, `i = 2.5 * 2.4`, `var k int = 2.5`.
- Eine andere Zeile lässt sich nicht wählen, auch kein Beispiel aus einem anderen Abschnitt des Kurstexts.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.7
- 4.1
- 5.2

## Schritt 2: Explizite Typkonvertierung

Reiter: konvertierung

### Aufbau der Darstellung

- Oben steht die Schreibweise `T(a)` mit einer Liste, die `T` und `a` erklärt.
- Darunter stehen zwei Eingabefelder für a und f und drei Knöpfe mit Beispielen.
- Das Programm darunter deklariert `a` als `int` und `f` als `float64` und gibt mehrere Ausdrücke aus. Hinter `//` steht die Ausgabe.
- Zwei Zeilen des Programms sind auskommentiert: `0.75 * a` und `f * a`. Go übersetzt sie nicht. Hinter ihnen steht „Fehler beim Übersetzen“.
- Die vorletzte Ausgabe `fmt.Println(f)` zeigt, dass f nach `int(f)` seinen Wert behalten hat. Danach folgt die Zuweisung `f = float64(a)` und eine zweite Ausgabe von f.
- Die Tabelle unter dem Programm hat fünf Zeilen, je eine für `0.75 * a`, `0.75 * float64(a)`, `f * a`, `f * float64(a)` und `int(f) * a`. Sie nennt, was Go mit dem Ausdruck macht, und seinen Wert und Typ. Die zwei Zeilen mit Fehler sind rot unterlegt und nennen die Meldung des Compilers.
- Der gelbe Kasten am Ende vergleicht `f * float64(a)` mit `int(f) * a`. Weicht die Ausgabe von Go von der Rechnung von Hand ab, erklärt er stattdessen den Rundungsfehler von `float64`. Beispiel: 1.1 mal 3 gibt Go als 3.3000000000000003 aus.

### Handgriffe

- Für a eine ganze Zahl eingeben: −99 bis 99.
- Für f eine Zahl mit höchstens einer Nachkommastelle eingeben: −9.9 bis 9.9.
- Ein Beispiel wählen: a = 25 und f = 1.7 wie im Kurstext, a = 3 und f = 1.1, a = −4 und f = −2.8.
- Der Faktor 0.75 ist fest.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.7

## Schritt 3: Abschneiden

Reiter: abschneiden

### Aufbau der Darstellung

- Zu sehen ist ein Zahlenstrahl von −10 bis 10 mit einem Strich an jeder ganzen Zahl. Der Strich bei 0 ist dicker.
- Der grüne Punkt auf dem Zahlenstrahl ist der Wert von f. Über ihm steht „f = …“.
- Der blaue Ring ist der Wert von `int(f)`. Unter ihm steht „int(f) = …“.
- Ein schwarzer Pfeil führt vom grünen Punkt zum blauen Ring. Er zeigt immer in Richtung 0: bei positivem f nach links, bei negativem f nach rechts.
- Hat f keine Nachkommastellen, liegt der grüne Punkt im blauen Ring, und es gibt keinen Pfeil.
- Unter dem Zahlenstrahl steht ein Programm mit vier Zeilen. Es speichert `int(f)` in der Variablen n und gibt n und f aus. Die Ausgabe von f zeigt, dass f seinen Wert behält.
- Die Tabelle darunter vergleicht drei Verfahren für den eingestellten Wert: Abschneiden mit `int(f)`, Abrunden und Runden. Die Zeile zu `int(f)` ist gelb unterlegt. Abrunden und Runden stehen nur zum Vergleich da, die Seite nennt dafür keine Funktion von Go.
- Der gelbe Kasten am Ende sagt, dass Go `int(3.7)` mit einem Literal nicht übersetzt, `int(5.0)` aber schon.

### Handgriffe

- Für f eine Zahl mit höchstens einer Nachkommastelle eingeben: −9.9 bis 9.9.
- f mit dem Regler in Schritten von 0.1 verstellen.
- Ein Beispiel wählen: 1.7 wie im Kurstext, 3.9, −2.8, −0.9, 5.0.
- Einstellen lässt sich nur der Wert der Variablen f. Einen Ausdruck wie `int(5.0)` oder `int(3.7)` kann der Student auf der Seite nicht eingeben.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.7

## Schritt 4: Division

Reiter: division

Die Lektion „Rechnen mit Ganzzahlen: Division und Rest in Go“ ist hier für Ganzzahlquotient, Rest und negative Operanden verlinkt. Dieser Schritt bleibt bei der Typkonvertierung vor der Division.

### Aufbau der Darstellung

- Oben steht als Liste, wie `/` bei zwei Operanden vom Typ `int` und bei zwei Operanden vom Typ `float64` rechnet.
- Darunter stehen zwei Eingabefelder für a und b und vier Knöpfe mit Beispielen.
- Zu sehen sind drei Rechnungen nebeneinander. Sie gehören zusammen: Alle drei rechnen mit denselben Werten von a und b.
- Linke Rechnung, blau umrandet: `a / b`, die ganzzahlige Division.
- Mittlere Rechnung, grün umrandet: `float64(a) / float64(b)`, der rationale Quotient.
- Rechte Rechnung: `float64(a / b)`, der Denkfehler. Sie ist rot umrandet, wenn die Division nicht aufgeht. Geht die Division auf, ist sie grün umrandet, weil der Zahlenwert dann stimmt.
- Jede Rechnung ist eine nummerierte Liste. Der erste Punkt setzt die Werte ein, die folgenden rechnen von innen nach außen. Darunter stehen fett Wert und Typ des Ausdrucks.
- Unter den drei Rechnungen steht das Programm mit den drei Ausgaben, darunter die Erklärung in Worten.
- Der gelbe Kasten sagt, dass Go `float64(a) / b` nicht übersetzt.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Logische Operatoren und Vergleiche in Go“.
  - Quiz: der „Selbsttest: Typkonvertierung“ mit 10 Fragen und je vier Antworten und die „Quizrunde richtig oder falsch: Typkonvertierung“.

### Handgriffe

- Für a eine ganze Zahl eingeben: −99 bis 99.
- Für b eine ganze Zahl eingeben: −99 bis 99, aber nicht 0. Die Division durch 0 lässt sich auf der Seite nicht ausprobieren.
- Ein Beispiel wählen: 25 und 10 wie im Kurstext, 7 und 2, −7 und 2, 20 und 5.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.7
- 4.3.7

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Implizite Typanpassung, 4.7, Seite 109; 4.1, Seiten 46–47: gleicher Typ beider Operanden: Die Programmzeilen stammen aus dem Kurstext. Die Zerlegung jeder Zeile in zwei Kästen gibt die Lektion dazu.
  - 2 Explizite Typkonvertierung, 4.7, Seiten 109–110: Das Programm folgt den Beispielen des Kurstexts. Die zwei Ausgaben von f am Ende hat die Lektion ergänzt.
  - 3 Abschneiden, 4.7, Seite 110: Das Abschneiden der Nachkommastellen und das Beispiel f = 1.7 stammen aus dem Kurstext. Die Werte 3.9, −2.8, −0.9 und 5.0 und das Bild mit dem Pfeil gibt die Lektion dazu.
  - 4 Division, 4.7, Seiten 110–111; 4.3.7.1, Seiten 69–71: ganzzahlige Division: Die Division mit Konvertierung und die ganzzahlige Division stammen aus dem Kurstext. Die Auswertung von innen nach außen mit farbigen Rahmen gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.4, Seiten 11–12, Konversion: Go ist streng typisiert und wandelt Werte nicht von selbst in einen anderen Typ um. Lektion: Variablen, Typen und Speicher, Schritt 2.
  - Abschnitt 4.4, Seiten 85–92, Gleitkommazahlen: Lektion: Gleitkommazahlen.
  - Anwendung: Die Programmierübung Zahlenraten teilt mit float64(gesamtVersuche) / float64(anzahlSiege).

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.4
- 4.1
- 4.3.7.1
- 4.4
- 4.7
