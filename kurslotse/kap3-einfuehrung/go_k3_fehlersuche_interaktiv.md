# Fehlersuche: drei Go-Programme mit mehreren Fehlern

Die Seite ist eine Übung. Drei Programme enthalten je mehrere Fehler. Einige meldet der Compiler, die übrigen zeigen sich erst an der Ausgabe. Die Seite gehört zu Kapitel 3 des Kurstexts, Abschnitt 3.3 und 3.4, Seiten 25–41.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Jede Demo zeigt ein Programm in vier oder fünf Fassungen nacheinander. Fassung 1 heißt „Das Programm lesen“ und zeigt das fehlerhafte Programm ohne Meldung. Jede weitere Fassung ist ein Schritt der Suche: übersetzen, eine Stelle korrigieren, wieder übersetzen. Gelb sind die Zeilen, die der Compiler meldet oder auf die es gerade ankommt.
- Unter dem Programm stehen vier Knöpfe: „weiter“, „zurück“, „von vorn“ und „Diese Fassung kopieren“. Darunter stehen die Meldung oder die Ausgabe und ein Kasten mit einem Satz. Bei einer Meldung ist der Kasten rot.
- Solange Fassung 1 zu sehen ist, soll der Student die Fehler selbst suchen. Fragt er in Fassung 1 nach den Fehlern, bekommt er einen Hinweis auf eine Zeile oder auf die Art eines Fehlers und nicht die ganze Lösung. Bei jeder späteren Fassung darf die Antwort die Fehler nennen, die bis dahin gezeigt sind.
- Zwei Arten von Fehlern: Fehler, die der Compiler meldet, zum Beispiel ein fehlendes Komma oder ein unbekannter Bezeichner. Und Fehler, die er nicht meldet, zum Beispiel eine falsche Formel. Das Programm läuft dann und gibt falsche Werte aus.
- Nach einem Syntaxfehler bricht der Compiler ab. Weitere Fehler erscheinen erst, wenn der erste behoben ist.
- Englische Wörter der Meldungen: `undefined` heißt nicht definiert, `declared and not used` deklariert und nicht verwendet, `mismatched types` nicht zusammenpassende Typen, `unexpected newline` unerwarteter Zeilenumbruch.
- Auf der Seite läuft kein Go. Der Browser zeigt, was Go 1.26 gemeldet oder ausgegeben hat. Der Student kann das Programm auf der Seite nicht ändern. Er kann eine Fassung kopieren und im eigenen Editor mit `go run` übersetzen.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Arten von Fehlern, je mit einem Beispiel.
- Unter „So arbeitest du mit den drei Programmen“ stehen drei Punkte: lesen, weiter, selbst ausprobieren.
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
- Die Demo hat fünf Fassungen.
- Fassung 2: Der Compiler meldet nur Zeile 19: `syntax error: unexpected name b in argument list; possibly missing comma or )`. In `rechteck(a b)` fehlt das Komma.
- Fassung 3, nach dem Komma: zwei Meldungen. Zeile 11: `declared and not used: volumen`. Zeile 20: `undefined: fläche`. In Zeile 20 steht `quader(fläche, c)`, und `fläche` ist nur im Rumpf von `rechteck` gültig.
- Fassung 4: In Zeile 20 steht `quader(rechteck(a, b), c)`, wie in Listing 3.3. Es bleibt die Meldung zu Zeile 11. Sie führt zu Zeile 12: `return grundfläche`.
- Fassung 5: Zeile 12 lautet `return volumen`. Das Programm gibt „Fläche: 12“ und „Volumen: 60“ aus.
- Hätte `quader` das Volumen ohne eigene Variable berechnet, hätte der Compiler den dritten Fehler nicht gemeldet. Das Programm hätte dann „Volumen: 12“ ausgegeben.
- Eine Liste erklärt Gültigkeitsbereich, Rückgabewert und die Regel zur unbenutzten Variable.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Diese Fassung kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.4

## Schritt 2: Temperaturen

Reiter: temperatur

### Aufbau der Darstellung

- Das Programm ist ein Auszug aus Listing 3.2 mit drei eingebauten Fehlern. Es hat 14 Zeilen. In `main` stehen `tempC` = 22 vom Typ `float64`, `tempText` = "Temperatur in", in Zeile 8 `tempF` und in Zeile 9 `tempK`, dann drei Aufrufe von `fmt.Println` in den Zeilen 11 bis 13. Erwartet sind 71,6 °F und 295,15 K.
- Die Demo hat fünf Fassungen.
- Fassung 2: Der Compiler meldet Zeile 11: `invalid operation: tempText + " °C:" + tempC (mismatched types string and float64)`.
- Fassung 3: Zeile 11 übergibt `tempC` als eigenes Argument. Das Programm läuft und gibt 97.2 für Fahrenheit und -251.14999999999998 für Kelvin aus. In Zeile 8 steht `(32 + tempC) * 1.8`.
- Fassung 4: Zeile 8 lautet `32 + tempC*1.8`. Fahrenheit ist 71.6. Kelvin ist weiter negativ, in Zeile 9 steht `tempC - 273.15`.
- Fassung 5: Zeile 9 lautet `tempC + 273.15`. Die Ausgabe ist „Temperatur in °C: 22“, „Temperatur in °F: 71.6“, „Temperatur in  K: 295.15“.
- Die Ausgabe -251.14999999999998 an der Stelle von -251.15 kommt von der Darstellung von Gleitkommazahlen. Das erklärt die Lektion „Gleitkommazahlen“ zu Kapitel 4, auf der Seite steht dazu nichts.
- Eine Liste erklärt die Bindungsstärke von `*` und `+` und das Prüfen einer Formel mit einem bekannten Wert: 0 °C sind 32 °F und 273,15 K.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Diese Fassung kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.3.6
- 3.3.7
- 3.3.5

## Schritt 3: Der Wert bleibt 7

Reiter: kopie

### Aufbau der Darstellung

- Das Programm ist ein eigenes Beispiel mit zwei Fehlern. Die Funktion `addiereEins(zahl int) int` erhöht `zahl` um eins und gibt `zahl` zurück. In `main` steht `var x int = 7`, eine Anweisung `if x > 0` mit der Klammer `{` in einer eigenen Zeile, der Aufruf `addiereEins(x)` und die Ausgabe von `x`. Erwartet ist am Ende der Wert 8.
- Die Demo hat vier Fassungen.
- Fassung 2: Der Compiler meldet Zeile 12: `syntax error: unexpected newline, expected { after if clause`.
- Fassung 3: Die Klammer steht am Ende der Zeile mit `if`. Das Programm läuft und gibt „x ist positiv“ und „Wert von x: 7“ aus. Der Aufruf steht in Zeile 15.
- Fassung 4: Zeile 15 lautet `x = addiereEins(x)`. Die Ausgabe ist „Wert von x: 8“.
- Eine Liste erklärt: Der Parameter `zahl` ist eine eigene Variable der Funktion und bekommt eine Kopie des Werts von `x`. Der Rückgabewert geht verloren, wenn der Aufruf allein in einer Zeile steht.
- Go meldet einen nicht verwendeten Rückgabewert einer eigenen Funktion nicht.
- Die öffnende Klammer `{` muss in der gleichen Zeile stehen wie `if` und die Bedingung. In einer neuen Zeile ist sie immer ein Fehler.
- Unter „Wie es weitergeht“ stehen Links auf vier Quiz und auf die Übersicht der Quiz zu Kapitel 3.

### Handgriffe

- „weiter“, „zurück“, „von vorn“ oder „Diese Fassung kopieren“ wählen.

### Abschnitte des Kurstexts

- 3.4
- 3.3.2
