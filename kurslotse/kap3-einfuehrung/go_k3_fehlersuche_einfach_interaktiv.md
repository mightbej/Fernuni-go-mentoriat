# Fehlersuche: drei kurze Go-Programme

Die Seite ist eine Übung. Drei kurze Programme enthalten je einen Fehler, den der Compiler meldet. Der Student sucht den Fehler, liest die Meldung und sieht die Korrektur. Die Seite gehört zu Kapitel 3 des Kurstexts, Abschnitt 3.3 und 3.4, Seiten 25–41.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt. Schritt 1 bis 3 zeigen je ein Programm, Schritt 4 ordnet die drei Fehler dem Kurstext zu.

Für Schritt 1 bis 3 gilt:

- Jede Demo zeigt drei Programme nacheinander. Programm 1 „Das Programm lesen“: das fehlerhafte Programm ohne Meldung. Programm 2 „Übersetzen“: das gleiche Programm mit der Meldung des Compilers, die gemeldete Zeile ist gelb. Programm 3 „Korrigieren“: das berichtigte Programm mit seiner Ausgabe, die geänderte Zeile ist gelb.
- Unter dem Programm stehen vier Knöpfe: „weiter“, „zurück“, „von vorn“ und „Dieses Programm kopieren“. Darunter stehen die Meldung oder die Ausgabe und ein Kasten mit einem Satz. Bei einer Meldung ist der Kasten rot.
- Solange Programm 1 zu sehen ist, soll der Student den Fehler selbst suchen. Fragt er in Programm 1 nach dem Fehler, bekommt er einen Hinweis auf die Zeile oder auf die Art des Fehlers und nicht die Lösung. Ab Programm 2 darf die Antwort den Fehler nennen.
- Eine Meldung des Compilers hat die Form Datei, Zeile, Spalte, Text, zum Beispiel `.\main.go:10:28: syntax error: …`. Die Meldung nennt die Stelle, an der der Compiler nicht weiterkam. Der Fehler steht dort oder kurz davor.
- Englische Wörter der Meldungen: `syntax error` heißt Syntaxfehler, `unexpected` unerwartet, `expected` erwartet, `literal` ein Wert, der wörtlich im Programm steht, `newline` Zeilenumbruch, `mismatched types` nicht zusammenpassende Typen, `argument list` Liste der Argumente.
- Auf der Seite läuft kein Go. Der Browser zeigt, was Go 1.26 gemeldet oder ausgegeben hat. Der Student kann das Programm auf der Seite nicht ändern. Er kann jedes Programm kopieren und im eigenen Editor mit `go run` übersetzen.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Unter „So liest du eine Meldung“ steht die Meldung aus Schritt 1 mit einer Liste ihrer vier Teile: Datei, Zeile, Spalte, Text.
- Unter „So arbeitest du mit den drei Programmen“ stehen vier Punkte: lesen, übersetzen, korrigieren, selbst ausprobieren.
- Unter „Was dazugehört“ stehen Links auf die Lektion „Höhere und niedrigere Programmiersprachen“, Schritt 4, und auf die Lektion „Fehlersuche: drei Programme mit mehreren Fehlern“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.3
- 2.7

## Schritt 1: Argumente trennen

Reiter: argumente

### Aufbau der Darstellung

- Das Programm hat elf Zeilen: die Funktion `rechteck(a int, b int) int` mit `return a * b` und in `main` die Zeile 10 `fmt.Println(rechteck(3 4))`.
- Der Fehler: Zwischen den Argumenten 3 und 4 fehlt das Komma.
- Die Meldung: `.\main.go:10:28: syntax error: unexpected literal 4 in argument list; possibly missing comma or )`.
- Die Korrektur: `fmt.Println(rechteck(3, 4))`. Die Ausgabe ist 12.
- Eine Liste erklärt Parameter und Argument: Parameter stehen in der Deklaration, Argumente beim Aufruf, beide durch Komma getrennt.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.4

## Schritt 2: Text und Zahl

Reiter: text

### Aufbau der Darstellung

- Das Programm hat acht Zeilen. In `main` steht `var tempC int = 22` und in Zeile 7 `fmt.Println("Temperatur in C: " + tempC)`.
- Der Fehler: Der Operator `+` soll einen String mit einer Variable vom Typ `int` verbinden.
- Die Meldung: `.\main.go:7:17: invalid operation: "Temperatur in C: " + tempC (mismatched types untyped string and int)`. `untyped string` steht da, weil der Text ein Literal ist.
- Die Korrektur: `fmt.Println("Temperatur in C:", tempC)`. Die Ausgabe ist „Temperatur in C: 22“. Das Leerzeichen setzt `fmt.Println` zwischen zwei Argumente.
- Eine Liste erklärt: `+` mit zwei Strings verkettet sie. `+` mit String und Zahl gibt es in Go nicht. `fmt.Println` nimmt mehrere Argumente verschiedener Typen.
- Eine Zahl in einen String umzuwandeln, zeigt die Seite nicht.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.3.7

## Schritt 3: Die Klammer nach if

Reiter: klammer

### Aufbau der Darstellung

- Das Programm hat elf Zeilen. In `main` steht `var x int = 5`, in Zeile 7 `if x > 0` und in Zeile 8 allein die Klammer `{`.
- Der Fehler: Die öffnende Klammer steht in einer neuen Zeile.
- Die Meldung: `.\main.go:7:13: syntax error: unexpected newline, expected { after if clause`.
- Die Korrektur: `if x > 0 {` in einer Zeile. Die Ausgabe ist „positiv“.
- Der Text erklärt den Grund: Der Compiler fügt am Ende einer Zeile selbst ein Semikolon ein, wenn dort eine Anweisung zu Ende sein könnte. Die Zeile `if x > 0` endet mit einem Literal, also setzt er eines.
- Die genaue Regel aus Expertenwissen 3.5: Der Compiler fügt ein Semikolon ein, wenn die Zeile mit einem Bezeichner, einem Literal, einem der Schlüsselwörter `break`, `continue`, `fallthrough`, `return` oder einem der Zeichen `++`, `--`, `)`, `]`, `}` endet.
- Ein gelber Kasten nennt die Faustregel: Eine öffnende Klammer steht nie am Anfang einer neuen Zeile.
- Im Kurstext steht für diesen Fehler die Meldung `unexpected semicolon or newline before {`. Go 1.26 meldet ihn nach `if` mit dem Wortlaut der Demo.
- Unter dem Text steht ein Link auf die Lektion „if-Anweisung in Go“.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf die zweite Lektion zur Fehlersuche und auf drei Quiz: „Implizite Semikolons“, „Parameter und Argument“, „Ausgabe und String-Verkettung“.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.3.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Fehler eine Zeile, zusammen drei Zeilen. Die Spalten: „Fehler im Programm“ mit Schritt, Zeile und Code, „Was falsch ist“, „Abschnitt im Kurstext“ mit gedruckten Seiten und „Lektion dazu“ mit Link.
- Schritt 1, Zeile 10, `rechteck(3 4)`: fehlendes Komma. Abschnitt 3.4, Seiten 37–41, Hinweis 3.7, und Abschnitt 2.6, Seiten 15–17. Lektion „Unterprogramme und Funktionen“, Schritt 2.
- Schritt 2, Zeile 7, `"Temperatur in C: " + tempC`: `+` mit String und Zahl. Abschnitt 3.3.7, Seite 36, Abschnitt 4.1, Seiten 45–47, und Abschnitt 3.3.5, Seiten 33–35. Lektion „Strings in Go“, Schritt 4.
- Schritt 3, Zeile 7 und 8, `{` in neuer Zeile: Abschnitt 3.3.2, Seiten 27–29, Expertenwissen 3.5, und Abschnitt 5.5.1, Seiten 129–132. Lektion „if-Anweisung in Go“, Schritt 1.
- Unter „Was die Lektion voraussetzt“ steht Abschnitt 2.7, Seiten 17–19, mit der Lektion „Höhere und niedrigere Programmiersprachen“, Schritt 4.
- Unter „Woher die Programme stammen“ steht: Schritt 1 ist ein eigenes Beispiel nach Listing 3.3, Schritt 2 nach Listing 3.2, Schritt 3 ein eigenes Beispiel.
- Unter „Was über den Kurstext hinausgeht“ steht: Die drei Fehler stehen nicht im Kurstext. Für die Klammer in einer neuen Zeile nennt Abschnitt 3.3.2 die Meldung `unexpected semicolon or newline before {`, Go 1.26 meldet nach `if` `unexpected newline, expected { after if clause`.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.4
- 3.3.7
- 3.3.2
- 3.3.5
- 2.6
- 4.1
- 5.5.1
