# Fehlersuche: drei Go-Programme mit mehreren Fehlern

Die Seite ist eine Übung. Drei Programme enthalten je mehrere Fehler. Einige meldet der Compiler, die übrigen zeigen sich erst an der Ausgabe. Die Seite gehört zu Kapitel 3 des Kurstexts, Abschnitt 3.3 und 3.4, Seiten 25–41.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt. Schritt 1 bis 3 zeigen je ein Programm, Schritt 4 ordnet die acht Fehler dem Kurstext zu.

Für Schritt 1 bis 3 gilt:

- Jede Demo zeigt vier oder fünf Programme nacheinander. Programm 1 heißt „Das Programm lesen“ und zeigt das fehlerhafte Programm ohne Meldung. Jedes weitere Programm ist ein Schritt der Suche: übersetzen, eine Stelle korrigieren, wieder übersetzen. Gelb sind die Zeilen, die der Compiler meldet oder auf die es gerade ankommt.
- Unter dem Programm stehen vier Knöpfe: „weiter“, „zurück“, „von vorn“ und „Dieses Programm kopieren“. Darunter stehen die Meldung oder die Ausgabe und ein Kasten mit einem Satz. Bei einer Meldung ist der Kasten rot.
- Solange Programm 1 zu sehen ist, soll der Student die Fehler selbst suchen. Fragt er in Programm 1 nach den Fehlern, bekommt er einen Hinweis auf eine Zeile oder auf die Art eines Fehlers und nicht die ganze Lösung. Bei jedem späteren Programm darf die Antwort die Fehler nennen, die bis dahin gezeigt sind.
- Zwei Arten von Fehlern: Fehler, die der Compiler meldet, zum Beispiel ein fehlendes Komma oder ein unbekannter Bezeichner. Und Fehler, die er nicht meldet, zum Beispiel eine falsche Formel. Das Programm läuft dann und gibt falsche Werte aus.
- Nach einem Syntaxfehler bricht der Compiler ab. Weitere Fehler erscheinen erst, wenn der erste behoben ist.
- Englische Wörter der Meldungen: `undefined` heißt nicht definiert, `declared and not used` deklariert und nicht verwendet, `mismatched types` nicht zusammenpassende Typen, `unexpected newline` unerwarteter Zeilenumbruch.
- Auf der Seite läuft kein Go. Der Browser zeigt, was Go 1.26 gemeldet oder ausgegeben hat. Der Student kann das Programm auf der Seite nicht ändern. Er kann jedes Programm kopieren und im eigenen Editor mit `go run` übersetzen.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Arten von Fehlern, je mit einem Beispiel.
- Unter „So arbeitest du mit den drei Programmen“ stehen vier Punkte: lesen, weiter, selbst ausprobieren, nachlesen. Der letzte Punkt verweist auf Schritt 4.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Fehlersuche: drei kurze Go-Programme“, „Höhere und niedrigere Programmiersprachen“ und „Unterprogramme und Funktionen“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.3
- 2.7

## Schritt 1: Rechteck und Quader

Reiter: quader

### Aufbau der Darstellung

- Das Programm ist Listing 3.3 mit kürzeren Ausgaben und drei eingebauten Fehlern. Es hat 21 Zeilen: `rechteck(seiteA int, seiteB int) int` in den Zeilen 5 bis 8 mit der Variable `fläche`, `quader(grundfläche int, höhe int) int` in den Zeilen 10 bis 13 mit der Variable `volumen`, `main` in den Zeilen 15 bis 21 mit `a` = 3, `b` = 4, `c` = 5. Erwartet sind die Ausgaben „Fläche: 12“ und „Volumen: 60“.
- Die Demo hat fünf Programme.
- Programm 2: Der Compiler meldet nur Zeile 19: `syntax error: unexpected name b in argument list; possibly missing comma or )`. In `rechteck(a b)` fehlt das Komma.
- Programm 3, nach dem Komma: zwei Meldungen. Zeile 11: `declared and not used: volumen`. Zeile 20: `undefined: fläche`. In Zeile 20 steht `quader(fläche, c)`, und `fläche` ist nur im Rumpf von `rechteck` gültig.
- Programm 4: In Zeile 20 steht `quader(rechteck(a, b), c)`, wie in Listing 3.3. Es bleibt die Meldung zu Zeile 11. Sie führt zu Zeile 12: `return grundfläche`.
- Programm 5: Zeile 12 lautet `return volumen`. Das Programm gibt „Fläche: 12“ und „Volumen: 60“ aus.
- Hätte `quader` das Volumen ohne eigene Variable berechnet, hätte der Compiler den dritten Fehler nicht gemeldet. Das Programm hätte dann „Volumen: 12“ ausgegeben.
- Eine Liste erklärt Gültigkeitsbereich, Rückgabewert und die Regel zur unbenutzten Variable.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.4

## Schritt 2: Temperaturen

Reiter: temperatur

### Aufbau der Darstellung

- Das Programm ist ein Auszug aus Listing 3.2 mit drei eingebauten Fehlern. Es hat 14 Zeilen. In `main` stehen `tempC` = 22 vom Typ `float64`, `tempText` = "Temperatur in", in Zeile 8 `tempF` und in Zeile 9 `tempK`, dann drei Aufrufe von `fmt.Println` in den Zeilen 11 bis 13. Erwartet sind 71,6 °F und 295,15 K.
- Die Demo hat fünf Programme.
- Programm 2: Der Compiler meldet Zeile 11: `invalid operation: tempText + " °C:" + tempC (mismatched types string and float64)`.
- Programm 3: Zeile 11 übergibt `tempC` als eigenes Argument. Das Programm läuft und gibt 97.2 für Fahrenheit und -251.14999999999998 für Kelvin aus. In Zeile 8 steht `(32 + tempC) * 1.8`.
- Programm 4: Zeile 8 lautet `32 + tempC*1.8`. Fahrenheit ist 71.6. Kelvin ist weiter negativ, in Zeile 9 steht `tempC - 273.15`.
- Programm 5: Zeile 9 lautet `tempC + 273.15`. Die Ausgabe ist „Temperatur in °C: 22“, „Temperatur in °F: 71.6“, „Temperatur in  K: 295.15“.
- Die Ausgabe -251.14999999999998 an der Stelle von -251.15 kommt von der Darstellung von Gleitkommazahlen. Das erklärt die Lektion „Gleitkommazahlen“ zu Kapitel 4, auf der Seite steht dazu nichts.
- Eine Liste erklärt die Bindungsstärke von `*` und `+` und das Prüfen einer Formel mit einem bekannten Wert: 0 °C sind 32 °F und 273,15 K.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.3.6
- 3.3.7
- 3.3.5

## Schritt 3: Der Wert bleibt 7

Reiter: kopie

### Aufbau der Darstellung

- Das Programm ist ein eigenes Beispiel mit zwei Fehlern. Die Funktion `addiereEins(zahl int) int` erhöht `zahl` um eins und gibt `zahl` zurück. In `main` steht `var x int = 7`, eine Anweisung `if x > 0` mit der Klammer `{` in einer eigenen Zeile, der Aufruf `addiereEins(x)` und die Ausgabe von `x`. Erwartet ist am Ende der Wert 8.
- Die Demo hat vier Programme.
- Programm 2: Der Compiler meldet Zeile 12: `syntax error: unexpected newline, expected { after if clause`.
- Programm 3: Die Klammer steht am Ende der Zeile mit `if`. Das Programm läuft und gibt „x ist positiv“ und „Wert von x: 7“ aus. Der Aufruf steht in Zeile 15.
- Programm 4: Zeile 15 lautet `x = addiereEins(x)`. Die Ausgabe ist „Wert von x: 8“.
- Eine Liste erklärt: Der Parameter `zahl` ist eine eigene Variable der Funktion und bekommt eine Kopie des Werts von `x`. Der Rückgabewert geht verloren, wenn der Aufruf allein in einer Zeile steht.
- Go meldet einen nicht verwendeten Rückgabewert einer eigenen Funktion nicht.
- Die öffnende Klammer `{` muss in der gleichen Zeile stehen wie `if` und die Bedingung. In einer neuen Zeile ist sie immer ein Fehler.
- Unter dem Programm stehen Links auf die Lektionen „if-Anweisung in Go“ und „Funktionen in Go: Rückgaben und Wertübergabe“.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf vier Quiz und auf die Übersicht der Quiz zu Kapitel 3.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Dieses Programm kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.4
- 3.3.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Fehler eine Zeile, zusammen acht Zeilen. Die Spalten: „Fehler im Programm“ mit Schritt, Zeile und Code, „Was falsch ist“, „Abschnitt im Kurstext“ mit gedruckten Seiten und „Lektion dazu“ mit Link.
- Die Zuordnung in der Tabelle:
  - Schritt 1, Zeile 19, `rechteck(a b)`: fehlendes Komma. Abschnitt 3.4, Seiten 37–41, Hinweis 3.7. Lektion „Fehlersuche: drei kurze Go-Programme“, Schritt 1.
  - Schritt 1, Zeile 20, `quader(fläche, c)`: Bezeichner außerhalb seines Gültigkeitsbereichs. Abschnitt 3.4, Seiten 37–41. Lektion „Variablendeklaration in Go“, Schritt 3.
  - Schritt 1, Zeile 12, `return grundfläche`: falscher Rückgabewert, gefunden über die Meldung zur unbenutzten Variable `volumen`. Abschnitt 3.4, Seiten 37–41, und Abschnitt 5.1.5, Seiten 122–123. Lektion „Variablendeklaration in Go“, Schritt 5.
  - Schritt 2, Zeile 11, `tempText + " °C:" + tempC`: `+` mit String und Zahl. Abschnitt 3.3.7, Seite 36, Abschnitt 4.1, Seiten 45–47, und Abschnitt 3.3.5, Seiten 33–35. Lektion „Strings in Go“, Schritt 4.
  - Schritt 2, Zeile 8, `(32 + tempC) * 1.8`: falsche Klammern. Abschnitt 3.3.6, Seiten 35–36, und Abschnitt 4.1.2, Seiten 50–52. Lektion „Logische Operatoren und Vergleiche“, Schritt 3.
  - Schritt 2, Zeile 9, `tempC - 273.15`: falsches Vorzeichen, ohne Meldung. Abschnitt 3.3.6, Seiten 35–36, und Abschnitt 2.7, Seiten 17–19. Lektion „Höhere und niedrigere Programmiersprachen“, Schritt 4.
  - Schritt 3, Zeile 12 und 13, `{` in neuer Zeile: Abschnitt 3.3.2, Seiten 27–29, Expertenwissen 3.5. Lektion „Fehlersuche: drei kurze Go-Programme“, Schritt 3.
  - Schritt 3, Zeile 15, `addiereEins(x)`: Rückgabewert verloren. Abschnitt 3.4, Seiten 37–41, und Abschnitt 7.2, Seiten 169–170. Lektion „Funktionen in Go“, Schritt 5.
- Im Kurstext steht die Formel für Kelvin als `tempC - nullpunkt` mit der Konstante `nullpunkt` = -273.15. Das Programm der Seite schreibt `tempC + 273.15`.
- Unter „Was die Lektion voraussetzt“ stehen Abschnitt 2.6 und 2.7 mit Links auf die Lektionen „Unterprogramme und Funktionen“ und „Höhere und niedrigere Programmiersprachen“.
- Unter „Woher die Programme stammen“ steht: Schritt 1 ist Listing 3.3, Schritt 2 ein Auszug aus Listing 3.2, Schritt 3 ein eigenes Beispiel.
- Unter „Was über den Kurstext hinausgeht“ steht: Die acht Fehler stehen nicht im Kurstext. Für die Klammer in einer neuen Zeile nennt Abschnitt 3.3.2 die Meldung `unexpected semicolon or newline before {`, Go 1.26 meldet nach `if` `unexpected newline, expected { after if clause`.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.4
- 3.3.2
- 3.3.5
- 3.3.6
- 3.3.7
- 4.1
- 5.1.5
- 7.2
