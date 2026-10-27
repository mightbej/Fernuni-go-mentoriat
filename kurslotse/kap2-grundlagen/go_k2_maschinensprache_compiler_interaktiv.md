# Höhere und niedrigere Programmiersprachen: vom Programmtext zum Maschinenbefehl

Die Seite gehört zu Abschnitt 2.7 des Kurstexts. Sie zeigt, wie ein Programm in Maschinensprache aussieht, wie Compiler und Interpreter ein Programm aus einer höheren Sprache ausführbar machen, wie ein Compiler eine Anweisung in Token und Aufbau zerlegt und welche Fehler beim Übersetzen und beim Laufen auffallen.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite kommt ohne Go aus, wie Kapitel 2 des Kurstexts. Die Programmzeilen gehören zu keiner bestimmten Programmiersprache. Der Pfeil `←` bedeutet: Die Variable links bekommt den Wert, der rechts steht.
- Das Beispiel in Schritt 1, 2 und 4 ist das Programm aus Abschnitt 2.5 des Kurstexts. Es berechnet den größten gemeinsamen Teiler von n und m und hat fünf Zeilen: `Eingabe n, m`, `WHILE n ≠ m`, `IF n > m THEN n ← n − m`, `ELSE m ← m − n`, `Ausgabe n`.
- Die Maschinensprache der Seite ist erfunden. Sie gehört zu keinem echten Prozessor. Namen echter Prozessoren, Assembler und Register kommen auf der Seite nicht vor.
- Die Abläufe sind nachgestellt. Die Seite übersetzt kein echtes Programm. Der Student kann die Programme in Schritt 1, 2 und 4 nicht ändern.
- Das Wort „Schritt“ kommt zweimal vor. Die Reiter heißen Schritt 1 bis Schritt 4. Der Zähler neben den Knöpfen, zum Beispiel „Schritt 3 von 18“, zählt die Schritte eines Ablaufs.
- Das Halteproblem lässt die Seite aus. Im Kurstext steht es am Ende von Abschnitt 2.7.
- Go-Code schreibt der Kurslotse zu dieser Seite nur, wenn der Student ausdrücklich danach fragt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Ein Bild zeigt vier Stationen nebeneinander, mit Pfeilen von links nach rechts: Programm in höherer Sprache, Compiler, Programm in Maschinensprache, Prozessor. Der Compiler ist gelb hervorgehoben.
- Darunter steht, was die vier Schritte zeigen.
- Darunter steht in einem dunklen Kasten das Programm aus Abschnitt 2.5.

### Handgriffe

- Den Links zu einer anderen HTML-Seite folgen.
- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 2.7
- 2.5

## Schritt 1: Maschinensprache

Reiter: maschine

### Aufbau der Darstellung

- Oben steht eine Tabelle mit den sieben Arten von Befehlen der erfundenen Maschinensprache. Jede Zeile nennt den Befehlscode, die lesbare Schreibweise, die Wirkung und die Länge.
- Die Befehlscodes: `00000000` HALT, 1 Byte. `00000001` LIES x, 2 Bytes. `00000010` GIB AUS x, 2 Bytes. `00000011` SUB x, y, 3 Bytes, x bekommt den Wert x − y. `00000100` SPRINGE ZU a, 2 Bytes. `00000101` SPRINGE WENN x = y ZU a, 4 Bytes. `00000110` SPRINGE WENN x > y ZU a, 4 Bytes.
- Die Variable n hat die Nummer `00000001`, die Variable m hat die Nummer `00000010`. Eine Adresse steht als Dualzahl da, zum Beispiel `00010110` für 22.
- Darunter steht der Speicher als Raster von 25 Kästchen. Jedes Kästchen ist ein Byte, über dem Bitmuster steht „Adresse 0“ bis „Adresse 24“.
- Gelb ist der Befehlscode des Befehls, der zuletzt zerlegt ist. Blau sind seine Argumente. Grau sind die Bytes der Befehle davor. Weiß sind die Bytes, die noch nicht zerlegt sind.
- Darunter steht in einem grünen Feld, wie der Prozessor diesen Befehl zerlegt: Befehlscode, Zahl der Argumente, Länge, Bedeutung jedes Arguments, Adresse des nächsten Befehls.
- Darunter stehen zwei dunkle Kästen nebeneinander. Links steht das Programm in höherer Sprache mit den Zeilen 1 bis 5. Hell markiert ist die Zeile, zu der der zuletzt zerlegte Befehl gehört. Rechts stehen die schon zerlegten Befehle in lesbarer Schreibweise, vorn ihre Adresse.
- Das Programm hat zehn Befehle an den Adressen 0, 2, 4, 8, 12, 15, 17, 20, 22 und 24.
- Die Bytes an den Adressen 0, 1 und 2 haben alle das Bitmuster `00000001`. An Adresse 0 und 2 ist es der Befehlscode LIES, an Adresse 1 die Nummer der Variable n. Darum geht es in der Frage zum Vorhersagen.
- Die Schleife besteht aus drei Sprungbefehlen. Adresse 4: Sind n und m gleich, geht es zu Adresse 22 hinter die Schleife. Adresse 15 und Adresse 20: Sprung zurück zu Adresse 4.
- Der Befehl an Adresse 8 fragt, ob m größer als n ist. Dann überspringt der Prozessor `SUB n, m` und macht bei Adresse 17 mit `SUB m, n` weiter. Im Programm in höherer Sprache steht die Frage umgekehrt als `n > m`.
- In diesem Schritt wird das Programm nur zerlegt und nicht ausgeführt. Die Variablen haben hier keine Werte.

### Handgriffe

- „nächsten Befehl zerlegen“: Die Zerlegung geht einen Befehl weiter, bis zum zehnten Befehl.
- „einen Befehl zurück“ und „von vorn“.
- Die Bytes lassen sich nicht anklicken und nicht ändern. Ausführen lässt sich das Programm in diesem Schritt nicht, das geht in Schritt 2 auf dem Weg mit Compiler.

### Abschnitte des Kurstexts

- 2.7
- 2.5

## Schritt 2: Compiler und Interpreter

Reiter: uebersetzen

### Aufbau der Darstellung

- Unter der Liste mit den zwei Wegen steht ein gelber Kasten „Wozu das heute gebraucht wird“: Am 21. und 22. Oktober 2026 war das Moodle der FernUniversität von außen nicht erreichbar, weil ein Fehler im Interpreter von PHP steckte, also in dem Programm, das die PHP-Programme von Moodle ausführt. Ein Link führt zum Exkurs „Warum Moodle am 21. Oktober 2026 nicht erreichbar war“. PHP, Moodle und diese Lücke stehen nicht im Kurstext.
- Der Schritt hat zwei Demos: oben den Ablauf mit Compiler oder Interpreter, darunter unter der Überschrift „Viele Sprachen, viele Prozessoren“ die Zahl der Compiler mit und ohne Zwischensprache.
- Obere Demo: Zu sehen sind zwei dunkle Kästen nebeneinander. Links steht das Programm in höherer Sprache mit den Zeilen 1 bis 5. Rechts steht das Programm in Maschinensprache in lesbarer Schreibweise, vorn die Adresse.
- Das Programm läuft immer mit den Eingaben 6 und 4. Die Werte sind nacheinander (6, 4), (2, 4), (2, 2). Die Ausgabe ist 2.
- Weg mit Compiler, 18 Schritte: In den Schritten 1 bis 5 übersetzt der Compiler je eine Zeile. Rechts erscheinen ihre Befehle. In dieser Zeit haben n und m keinen Wert, und nichts wird ausgeführt. In den Schritten 6 bis 18 führt der Prozessor das übersetzte Programm aus. Hell markiert ist dann ein Befehl im rechten Kasten.
- Weg mit Interpreter, 8 Schritte: Der Interpreter liest eine Zeile, bildet ihre Wirkung sofort nach und liest die nächste. Hell markiert ist eine Zeile im linken Kasten. Im rechten Kasten steht nur der Satz, dass kein Programm in Maschinensprache entsteht.
- Rechts neben jeder Zeile des linken Kastens steht in Gelb, wie oft sie übersetzt oder gelesen ist. Der Compiler übersetzt jede Zeile einmal. Der Interpreter liest Zeile 2 dreimal und Zeile 3 zweimal, weil die Schleife zweimal durchläuft.
- Darunter stehen die Variablen n und m als Kästen, die Ausgabe und in einem grünen Feld ein Satz zum letzten Schritt. Gelb ist der Kasten, den der letzte Schritt geändert hat.
- Untere Demo: Zwei Bilder nebeneinander. Links stehen die höheren Sprachen als Kästchen „Sprache 1“ und so weiter, rechts die Maschinensprachen als „Maschine 1“ und so weiter. Jeder Strich ist ein Compiler.
- Im linken Bild führt von jeder Sprache ein Strich zu jeder Maschinensprache. Das sind n · m Compiler. Im rechten Bild führen alle Striche über ein gelbes Kästchen „Zwischensprache“ in der Mitte. Das sind n + m Compiler.
- Bei n = 1, bei m = 1 und bei n = m = 2 spart die Zwischensprache nichts. Bei n = 1 oder m = 1 braucht sie sogar einen Compiler mehr.
- Unter den Bildern steht ein Listenpunkt: Der Befehl `go run` übersetzt ein Go-Programm und startet es gleich danach.

### Handgriffe

- „mit Compiler“ oder „mit Interpreter“ wählen. Der Ablauf beginnt dann von vorn.
- „nächster Schritt“, „einen Schritt zurück“, „von vorn“.
- Die Eingaben 6 und 4 lassen sich nicht ändern.
- Zwei Regler: höhere Sprachen n von 1 bis 6, Maschinensprachen m von 1 bis 6. Voreingestellt sind n = 4 und m = 3.

### Abschnitte des Kurstexts

- 2.7
- 3.2

## Schritt 3: Token und Syntax

Reiter: token

### Aufbau der Darstellung

- Oben steht Text: der Vergleich mit einem Übersetzer von Deutsch nach Englisch, die Wörter Syntax und Semantik, die lexikalische Analyse und die syntaktische Analyse.
- Unter „Ausprobieren“ steht ein Eingabefeld für eine Anweisung, daneben der Knopf „zerlegen“, darunter sechs Beispiele als Knöpfe.
- Darunter steht unter „1. Lexikalische Analyse: die Token“ je Token eine Marke. Oben in der Marke steht der Text des Tokens, darunter seine Art.
- Es gibt vier Arten von Token. Schlüsselwort, gelb: `IF`, `THEN`, `ELSE`, `WHILE`, nur in Großbuchstaben. Name, blau: ein Buchstabe, danach Buchstaben und Ziffern. Zahl, grün: Ziffern. Zeichen, weiß: `←`, `+`, `−`, `·`, `>`, `<`, `=`, `≠`, `(`, `)`.
- Leerzeichen trennen Token und sind selbst kein Token. `ab` ist ein Name, `a b` sind zwei Namen.
- Darunter steht unter „2. Syntaktische Analyse: der Aufbau der Anweisung“ der Aufbau als Kästen, die ineinander stecken. Jeder Kasten hat einen Titel, zum Beispiel „Zuweisung“ oder „Operation +“, und nennt seine Teile, zum Beispiel „Ziel“ und „neuer Wert“.
- Die Seite versteht zwei Arten von Anweisungen. Die Zuweisung: Name, Pfeil, Ausdruck. Die Verzweigung: `IF`, Bedingung, `THEN`, Zuweisung, wahlweise `ELSE` und eine zweite Zuweisung.
- Ein Ausdruck besteht aus Namen, Zahlen, Klammern und den Operationen `+`, `−` und `·`. Die Multiplikation steckt im Aufbau weiter innen als Addition und Subtraktion. Eine Division gibt es nicht.
- Bei einem Fehler steht statt des Aufbaus eine rote Meldung. Ein Zeichen wie `€`, das zu keinem Token gehört, fällt schon in der lexikalischen Analyse auf. Eine falsche Reihenfolge der Token fällt erst in der syntaktischen Analyse auf und heißt Syntaxfehler.
- Das Beispiel `a + b ← 5` zeigt den Satz des Kurstexts: `a + b` ist kein guter Name. Die Zerlegung liefert die drei Token a, + und b, und links vom Pfeil darf nur ein Name stehen.
- Das Beispiel `IF ← 3` zeigt: Ein Schlüsselwort darf kein Name einer Variable sein.
- Das Beispiel `x ← a + 2 · b` zeigt Punkt vor Strich im Aufbau.
- Die Seite prüft nur die Syntax. Sie rechnet keine Anweisung aus, und die Variablen haben keine Werte.

### Handgriffe

- Eine Anweisung mit bis zu 60 Zeichen eintippen und „zerlegen“ wählen oder die Eingabetaste drücken.
- Auf der Tastatur gilt `<-` als Pfeil, `-` als Minus, `*` als Malpunkt und `!=` als `≠`.
- Eines der sechs Beispiele wählen: `a ← a + b`, `x ← a + 2 · b`, die Verzweigung mit `IF`, `a + b ← 5`, `IF ← 3`, `preis ← 3 € + 2`.
- Eine Schleife mit `WHILE`, mehrere Anweisungen hintereinander, ein Semikolon und Go-Code nimmt die Seite nicht an. Sie zeigt dann eine Meldung.

### Abschnitte des Kurstexts

- 2.7
- 3.3.4
- 3.3.6

## Schritt 4: Fehler

Reiter: fehler

### Aufbau der Darstellung

- Oben steht die Spezifikation des Programms: Gültig sind zwei positive ganze Zahlen n und m. Richtig ist als Ausgabe ihr größter gemeinsamer Teiler. Für 6 und 4 ist das 2, für 4 und 6 auch.
- Darunter steht eine Tabelle mit drei Arten von Fehlern: Syntaxfehler, semantischer Fehler mit falscher Ausgabe, semantischer Fehler ohne Ausgabe, die Nichtterminierung.
- Unter „Ausprobieren“ stehen vier Programme zur Wahl und zwei Eingaben.
- Das richtige Programm stammt aus Abschnitt 2.5.
- Programm A hat in Zeile 3 `IF n > THEN n ← n − m`. Hinter `>` fehlt der rechte Teil des Vergleichs. Das ist ein Syntaxfehler. Der Compiler meldet ihn, und das Programm läuft nicht. Der Ablauf hat nur einen Schritt.
- Programm B hat in Zeile 4 `ELSE m ← n`. Das ist ein semantischer Fehler. Mit den Eingaben 6 und 4 gibt das Programm 2 aus, das stimmt zufällig. Mit 4 und 6 gibt es 4 aus, richtig wäre 2.
- Programm C hat in Zeile 2 `WHILE n ≠ 0`. Die Bedingung der Schleife ist schlecht gewählt. Die Werte erreichen (2, 2), dann (2, 0), und danach ändert sich nichts mehr. Das Programm endet nie. Die Seite bricht nach sieben Durchläufen der Schleife ab und sagt das.
- Der Compiler dieser Seite prüft nur, ob jede Zeile die Syntax einhält. Was das Programm beim Laufen mit den Werten tut, prüft er nicht. Deshalb meldet er bei Programm B und bei Programm C nichts. In Programm C hält `WHILE n ≠ 0` die Syntax ein. Beim Laufen wird m zu 0, n behält den Wert 2 und wird nie 0.
- Die vier Programme stehen in der Schreibweise von Kapitel 2 und nicht in Go. Go hat kein `WHILE`.
- Die Seite markiert die geänderte Zeile nicht und nennt die Art des Fehlers erst am Ende des Ablaufs. Der Student soll beides vorher selbst finden. Fragt er, welches Programm welchen Fehler hat, bekommt er einen Hinweis, worauf er achten soll, und führt das Programm dann selbst aus.
- Zu sehen ist das Programm in einem dunklen Kasten mit den Zeilen 1 bis 5. Hell markiert ist die Zeile des letzten Schritts.
- Darunter stehen die Meldung des Compilers, die Variablen n und m als Kästen, die Reihe „Werte (n, m) bisher“ und die Ausgabe des Programms. Darunter steht in einem grünen Feld ein Satz zum letzten Schritt.
- Der erste Schritt des Ablaufs übersetzt das Programm. Der zweite liest die Eingaben. Danach ist jeder Schritt ein Durchlauf der Schleife. Der letzte Schritt gibt aus und vergleicht mit der Spezifikation.
- Unter der Demo steht eine Liste „So geht es in Go weiter“ mit Links zu Kapitel 3 und Kapitel 5.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Exkurs: Warum Moodle am 21. Oktober 2026 nicht erreichbar war“.
  - Danach: In Kapitel 3 übersetzt du dein erstes Go-Programm. Die Lektion „Das erste Go-Programm: Hello, World!“ öffnest du im Editorlotsen.

### Handgriffe

- Ein Programm wählen: „richtiges Programm“, „Programm A“, „Programm B“ oder „Programm C“. Der Ablauf beginnt dann von vorn.
- Eine Eingabe wählen: „6 und 4“ oder „4 und 6“. Andere Zahlen lassen sich nicht eingeben.
- „nächster Schritt“, „einen Schritt zurück“, „von vorn“.
- Die Programme lassen sich nicht ändern. Eigene Fehler lassen sich nicht einbauen.

### Abschnitte des Kurstexts

- 2.7
- 3.3.2
- 5.5.3.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.7, Seiten 17–18; 2.5, Seite 13: Maschinensprache, höhere Programmiersprache und Compiler stammen aus Abschnitt 2.7. Das Beispielprogramm stammt aus Abschnitt 2.5. Das Bild vom Programm über den Compiler zum Prozessor gibt die Lektion dazu.
  - 1 Maschinensprache, 2.7, Seiten 17–18; 2.5, Seite 13: Dass ein Befehl aus Befehlscode und Argumenten besteht, verschieden lang ist und dass Maschinensprache Sprungbefehle hat, stammt aus dem Kurstext. Das Programm in höherer Sprache stammt aus Abschnitt 2.5. Die sieben Befehle, ihre Befehlscodes, die lesbare Schreibweise und die Adressen sind für die Lektion erfunden.
  - 2 Compiler und Interpreter, 2.7, Seite 18: Compiler, Interpreter und die Zwischensprache mit n + m statt n · m Compilern stammen aus dem Kurstext. Die Eingaben 6 und 4, die Zählung der gelesenen Zeilen und das Bild mit den Strichen gibt die Lektion dazu.
  - 3 Token und Syntax, 2.7, Seiten 18–19: Der Vergleich mit einem Übersetzer von Deutsch nach Englisch, Syntax und Semantik, Token, lexikalische und syntaktische Analyse, das Beispiel a + b und das Wort IF stammen aus dem Kurstext. Die übrigen Beispiele und die Darstellung des Aufbaus als Kästen gibt die Lektion dazu.
  - 4 Fehler, 2.7, Seite 19: Syntaxfehler, semantische Fehler, Spezifikation und Nichtterminierung stammen aus dem Kurstext. Die drei fehlerhaften Programme und die Eingaben sind eigene Beispiele. Das Halteproblem am Ende des Abschnitts lässt die Lektion aus.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.5, Seiten 13–14, Flussdiagramm und Struktogramm: Das Beispielprogramm zum größten gemeinsamen Teiler. Lektion: Größter gemeinsamer Teiler.
  - Abschnitt 3.2, Seiten 22–24, go run: Der Befehl übersetzt ein Go-Programm und startet es gleich danach. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 3.3.2, Seiten 27–29, die erste Meldung des Compilers: Steht die Klammer { in einer neuen Zeile, lehnt der Compiler das Programm mit syntax error ab. Lektion: Fehlersuche: drei kurze Go-Programme, Schritt 3.
  - Abschnitt 3.3.4, Seiten 30–32, Token in Go: Die Regeln für Namen und die 25 Schlüsselwörter. Lektion: Bezeichner und Schlüsselwörter in Go.
  - Abschnitt 3.3.6, Seiten 35–36, Punkt vor Strich: In Go wird wie in x ← a + 2 · b zuerst multipliziert. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Fehler in Go-Programmen suchen: Die Lektionen Fehlersuche: drei kurze Go-Programme und Fehlersuche: drei Go-Programme mit mehreren Fehlern üben das an Programmen aus Kapitel 3. Die zweite zeigt in Schritt 2 auch einen semantischen Fehler, den der Compiler nicht meldet.
  - Abschnitt 5.5.3, Seiten 137–138, Endlosschleife: Lektion: for-Anweisung mit Bedingung in Go, Schritt 1.
  - Abschnitt 7.5.5, Seiten 191–193, Terminierung: Wann rekursive Aufrufe zu einem Ende kommen. Lektion: Rekursion in Go, Schritt 5.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.5
- 2.7
- 3.2
- 3.3.2
- 3.3.4
- 3.3.6
- 5.5.3
- 7.5.5
