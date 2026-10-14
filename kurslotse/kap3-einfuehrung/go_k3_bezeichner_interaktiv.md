# Bezeichner und Schlüsselwörter in Go

Die Seite erklärt, welche Namen Go für Variablen, Funktionen und Typen zulässt: die zwei Regeln für Bezeichner, Groß- und Kleinschreibung, die 25 Schlüsselwörter, die vorgegebenen Bezeichner und den leeren Bezeichner. Sie gehört zu Kapitel 3 des Kurstexts, Abschnitt 3.3.4, Seiten 30–32.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Variablen, Konstanten, Typen und Funktionen heißen Programmelemente. Der Name eines Programmelements heißt Bezeichner. Bei der Deklaration wird ein Programmelement mit seinem Bezeichner verbunden.
- Regel 1: Das erste Zeichen eines Bezeichners ist ein Buchstabe oder der Unterstrich `_`. Regel 2: Die übrigen Zeichen sind Buchstaben, Ziffern oder Unterstriche. Ein Schlüsselwort ist kein zulässiger Bezeichner.
- Buchstaben sind alle Buchstaben des Unicode-Zeichensatzes, auch ä, ö, ü und ß. Das steht im Kurstext in einer Fußnote.
- Jede Auswahl der Seite setzt einen Namen oder eine Zeile in dieses Programm ein: Zeile 1 `package main`, Zeile 3 `import "fmt"`, Zeile 5 `func main() {`, Zeile 6 die Deklaration, Zeile 7 die Ausgabe, Zeile 8 `}`. Die Zeilennummern in den Meldungen des Compilers beziehen sich darauf.
- Eine Meldung des Compilers hat die Form Datei, Zeile, Spalte, Text, zum Beispiel `.\main.go:6:9: syntax error: …`. Die Meldungen sind englisch. `unexpected` heißt unerwartet, `expected name` heißt: erwartet war ein Name.
- Bei mehreren Fehlern zeigt die Seite nur die erste Meldung des Compilers.
- Auf der Seite läuft kein Go. Der Browser zeigt, was Go 1.26 zu jeder festen Auswahl gemeldet oder ausgegeben hat. Der Student kann keinen eigenen Namen und keinen eigenen Code eingeben.
- Regeln für gute Namen, zum Beispiel sprechende Namen oder die Schreibweise mit Binnenmajuskel, stehen in Abschnitt 3.3.4 nicht. Die Seite nennt nur die Konvention zum großen Anfangsbuchstaben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt die Begriffe Programmelement, Bezeichner und Deklaration.
- An der Zeile `var tempC float64 = 22` zeigt eine Liste drei Arten von Wörtern: `var` ist ein Schlüsselwort, `tempC` ein selbst gewählter Bezeichner, `float64` ein vorgegebener Bezeichner.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Variablen, Typen und Speicher“, „Höhere und niedrigere Programmiersprachen“ und „Variablendeklaration in Go“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.3.4
- 3.3.3

## Schritt 1: Zwei Regeln

Reiter: regeln

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste nennt die zwei Regeln.
- Vierzehn Knöpfe wählen einen Namen. Darunter steht das Programm mit diesem Namen in Zeile 6 und 7, dann die Ausgabe oder die Meldung des Compilers und ein Kasten mit einem Satz. Bei einer Meldung ist der Kasten rot.
- Zulässig sind `alter2`, `name_vorname`, `_intern`, `größe`, `π`, `x`, `HTTPServer` und `konto123`. Das Programm gibt dann 5 aus.
- `2alter`: `syntax error: unexpected literal 2, expected name`. Der Compiler liest die 2 als Zahl.
- `24stunden`: `syntax error: unexpected literal 24, expected name`. Das ist das Beispiel des Kurstexts.
- `mein-name`: `syntax error: unexpected -, expected type`. Der Compiler liest `mein` als Namen der Variable und erwartet danach den Typ.
- `mein name`: `syntax error: unexpected name int at end of statement`. Der Compiler liest `mein` als Namen und `name` als Typ.
- `user.name`: `syntax error: unexpected ., expected type`.
- `a$`: `invalid character U+0024 '$'`. Das Zeichen kommt in Go nicht vor.
- Ein gelber Kasten sagt: Der Compiler meldet nicht „ungültiger Bezeichner“, er meldet das erste Zeichen, das nicht passt.

### Handgriffe

- Einen der vierzehn Namen wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Namen kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 3.3.4

## Schritt 2: Groß und klein

Reiter: gross

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock deklariert `tempC` und gibt `tempc` aus. Darunter stehen die zwei Meldungen: `declared and not used: tempC` und `undefined: tempc`.
- Eine Liste erklärt den Export: Ein Bezeichner mit großem Anfangsbuchstaben ist auch von anderen Paketen aus erreichbar, zum Beispiel `Println` im Paket `fmt`. Ein Bezeichner mit kleinem Anfangsbuchstaben gilt nur im eigenen Paket.
- Ein gelber Kasten nennt die Konvention des Kurstexts: großer Anfangsbuchstabe genau dann, wenn der Bezeichner exportiert werden soll. Die Programme des Kurses beginnen alle Bezeichner mit einem Kleinbuchstaben.
- Pakete teilt der Kurs nicht auf. Wie man ein eigenes Paket schreibt, steht nicht im Kurstext.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.3.4
- 3.1

## Schritt 3: Schlüsselwörter

Reiter: schluessel

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text erklärt Schlüsselwort und reserviertes Wort.
- Eine Liste zeigt die 25 Schlüsselwörter: `break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var`. Grün sind `package`, `import`, `func` und `var`.
- Der Kurs braucht `chan`, `defer`, `go`, `interface`, `map`, `select` und `goto` nicht.
- Vier Knöpfe wählen eine Zeile. Darunter stehen die Ausgabe oder die Meldung und ein Kasten mit einem Satz.
- Zeile 1: `var func int = 5` ergibt `syntax error: unexpected keyword func, expected name`.
- Zeile 2: `var var int = 5` ergibt `syntax error: unexpected keyword var, expected name`.
- Zeile 3: `var go! int = 5` ergibt `syntax error: unexpected keyword go, expected name`. Das Ausrufezeichen erreicht der Compiler nicht.
- Zeile 4: `var funktion int = 5` übersetzt Go, die Ausgabe ist 5.

### Handgriffe

- Eine der vier Zeilen wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 3.3.4

## Schritt 4: Vorgegebene Bezeichner

Reiter: vorgegeben

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Tabelle nennt die vorgegebenen Bezeichner in drei Zeilen: Datentypen wie `int`, `float64`, `string` und `bool`, eingebaute Funktionen wie `len`, Konstanten `true`, `false`, `iota` und `nil`.
- Ein vorgegebener Bezeichner ist kein reserviertes Wort. Ein Programm darf ihn neu deklarieren.
- Vier Knöpfe wählen ein Paar von zwei Zeilen.
- Paar 1: `var int int = 5` und `fmt.Println(int)`. Go übersetzt, die Ausgabe ist 5.
- Paar 2: `var int int = 5` und `var y int = 7`. Die Meldung zur zweiten Zeile: `int (local variable) is not a type`.
- Paar 3: `var true int = 5` und `fmt.Println(true)`. Go übersetzt, die Ausgabe ist 5.
- Paar 4: `var _ int = 5` und `fmt.Println(_)`. Die Meldung: `cannot use _ as value or type`.
- Ein gelber Kasten warnt: einem vorgegebenen Bezeichner keine neue Bedeutung geben.
- Der Abschnitt „Der leere Bezeichner“ sagt: `_` ist ein Platzhalter für einen Wert, den das Programm nicht braucht. Der Kurstext geht später darauf ein. Ein Beispiel dafür steht auf der Seite nicht.
- Unter „Wie es weitergeht“ stehen Links auf das Quiz „Bezeichner und Schlüsselwörter“, auf die Lektion „Fehlersuche: drei kurze Go-Programme“ und auf den Zahlenkreis.

### Handgriffe

- Eines der vier Paare wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 3.3.4
