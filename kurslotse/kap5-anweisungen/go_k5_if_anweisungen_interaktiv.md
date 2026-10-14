# if-Anweisung in Go: bedingte Anweisung und Verzweigung

Die Seite erklärt die if-Anweisung in Go: die einfache Form mit Bedingung und Block, die Verzweigung mit `else` und die Kette mit `else if`. Sie führt das Schaltjahr-Programm aus Listing 5.2 Zeile für Zeile aus. Sie gehört zu Abschnitt 5.5.1 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Jeder Code auf der Seite ist Go. Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Ein Programm einer Demo steht in einem dunklen Kasten mit Zeilennummern am linken Rand.
- Ein grüner Strich links an einer Programmzeile bedeutet: Die Zeile wird ausgeführt. Eine blasse Programmzeile wird übersprungen.
- Der if-Zweig ist der Block direkt hinter der Bedingung. Der else-Zweig ist der Block hinter `else`. Ein else-if-Zweig ist der Block hinter `else if` mit eigener Bedingung.
- Schritt 1 bis 4 folgen dem Kurstext. Schritt 5 enthält Faustregeln des Mentors, die nicht im Kurstext stehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Oben steht, was eine bedingte Anweisung ist: Ein Stück des Programms läuft nur, wenn eine Bedingung erfüllt ist. Das entscheidet sich erst, während das Programm läuft.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.1 mit `IF b THEN {…}` und dem Block, Abschnitt 2.5 mit Flussdiagramm und Struktogramm, Abschnitt 4.5 mit dem booleschen Ausdruck.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Imperative Programmierung“ und „Größter gemeinsamer Teiler“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.1
- 2.1
- 2.5
- 4.5

## Schritt 1: if und Block

Reiter: if

### Aufbau der Darstellung

- Oben stehen die drei Teile einer if-Anweisung: das Schlüsselwort `if`, die Bedingung als boolescher Ausdruck, der Block zwischen `{` und `}`.
- Der Schritt hat zwei Demos. Sie hängen nicht zusammen.
- Obere Demo, „Ausprobieren: alter einstellen“: ein Programm mit sechs Zeilen. Zeile 1 deklariert `alter` mit dem eingestellten Wert. Zeile 3 ist `if alter < 18 {`. Zeile 4 ist der if-Zweig und gibt „Du bist minderjährig.“ aus. Zeile 6 steht hinter der if-Anweisung und gibt immer „Fertig.“ aus.
- Hinter der Bedingung in Zeile 3 steht ein kleines Schild mit ihrem Wert: grün mit `true` oder rot mit `false`.
- Zeile 4 hat einen grünen Strich, wenn die Bedingung `true` ergibt. Sonst ist sie blass.
- Unter dem Programm stehen „Bedingung“ mit der Rechnung, zum Beispiel „17 < 18 ergibt true“, und „Ausgabe“ mit der Ausgabe des Programms. Ein grüner Kasten erklärt den Lauf.
- Bei `alter` 18 ergibt `alter < 18` den Wert `false`, weil `<` die Grenze ausschließt.
- Untere Demo, „Was Go bei der Schreibweise verlangt“: fünf Knöpfe mit fünf Schreibweisen derselben if-Anweisung. Der gewählte Knopf ist grün. Darunter steht die Schreibweise in einem dunklen Kasten ohne Zeilennummern, darunter „Ergebnis“ oder „Meldung des Compilers“ in roter Schrift, darunter eine Erklärung.
- „wie üblich“: `if alter < 18 {`. Go übersetzt das Programm.
- „ohne Block“: Die geschweiften Klammern fehlen. Der Compiler meldet `syntax error: unexpected newline, expected { after if clause`.
- „{ in der nächsten Zeile“: Der Compiler meldet dasselbe. Die öffnende Klammer muss in der Zeile der Bedingung stehen.
- „mit runden Klammern“: `if (alter < 18) {`. Go übersetzt das Programm. Nach Idiomatic Go lässt man die Klammern weg.
- „Zahl als Bedingung“: `if alter {`. Der Compiler meldet `non-boolean condition in if statement`. Die Bedingung muss den Typ `bool` haben.
- Am Ende steht der Teil „Warum Go den Block erzwingt“: In Java oder Pascal darf eine einzelne Anweisung ohne Block stehen. Wer später eine zweite Anweisung ergänzt und die Klammern vergisst, bekommt eine Anweisung, die immer ausgeführt wird.

### Handgriffe

- Obere Demo: `alter` mit den Knöpfen − und + um 1 verändern. Erlaubt sind 0 bis 120. Die Knöpfe „17“ und „18“ setzen den Wert direkt.
- Untere Demo: eine der fünf Schreibweisen wählen. Andere Schreibweisen gibt es nicht.
- Die Programmzeilen lassen sich nicht ändern. Die Bedingung `alter < 18` steht fest.

### Abschnitte des Kurstexts

- 5.5.1
- 4.5

## Schritt 2: else

Reiter: else

### Aufbau der Darstellung

- Oben steht, was `else` bewirkt: Ergibt die Bedingung `true`, läuft der if-Zweig. Ergibt sie `false`, läuft der else-Zweig.
- Die Demo zeigt ein Programm in zwei Fassungen. Zu sehen ist immer nur eine. Der Knopf der gewählten Fassung ist grün.
- Fassung „mit else“, acht Zeilen: Zeile 1 deklariert `alter`. Zeile 3 ist `if alter >= 18 {`. Zeile 4 ist der if-Zweig und gibt „Du bist volljährig.“ aus. Zeile 5 ist `} else {`. Zeile 6 ist der else-Zweig und gibt „Du bist minderjährig.“ aus. Zeile 8 gibt immer „Fertig.“ aus.
- Fassung „ohne else“, sechs Zeilen: Die Zeilen 1 bis 4 sind gleich. Zeile 5 ist `}`. Zeile 6 gibt immer „Fertig.“ aus.
- Hinter der Bedingung in Zeile 3 steht ein Schild mit `true` oder `false`.
- Unter dem Programm stehen drei Angaben: „Bedingung“ mit der Rechnung, „Ausgeführte Zweige“ und „Ausgabe“.
- Mit `else` wird genau ein Zweig ausgeführt. Ohne `else` wird höchstens ein Zweig ausgeführt: Bei `false` läuft keiner.
- Die Bedingung heißt hier `alter >= 18`. In Schritt 1 heißt sie `alter < 18`. Die zwei Schritte haben getrennte Werte für `alter`.

### Handgriffe

- Die Fassung wählen: „mit else“ oder „ohne else“.
- `alter` mit den Knöpfen − und + um 1 verändern. Erlaubt sind 0 bis 120. Die Knöpfe „17“ und „18“ setzen den Wert direkt.
- Die Programmzeilen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.5.1

## Schritt 3: else if

Reiter: elseif

### Aufbau der Darstellung

- Oben stehen die drei Regeln der Kette: Die Bedingungen werden von oben nach unten ausgewertet. Der Zweig der ersten Bedingung mit `true` läuft, alle Zweige darunter werden übersprungen. Ergibt keine Bedingung `true`, läuft der else-Zweig.
- Die Demo zeigt ein Programm mit 13 Zeilen, das aus Punkten eine Note macht. Zeile 1 deklariert `punkte` mit dem eingestellten Wert.
- Die vier Bedingungen stehen in den Zeilen 3, 5, 7 und 9: `punkte >= 90`, `punkte >= 75`, `punkte >= 60`, `punkte >= 50`.
- Die Zweige stehen in den Zeilen 4, 6, 8 und 10 und geben „Note: Sehr gut“, „Note: Gut“, „Note: Befriedigend“ und „Note: Ausreichend“ aus. Der else-Zweig in Zeile 12 gibt „Note: Nicht bestanden“ aus.
- Hinter jeder Bedingung steht ein Schild: rot mit `false`, grün mit `true` oder grau mit „nicht geprüft“. „nicht geprüft“ steht bei jeder Bedingung unter dem ersten Treffer.
- Der Zweig, der ausgeführt wird, hat einen grünen Strich. Die übrigen Zweige sind blass.
- Unter dem Programm steht die Ausgabe. Ein grüner Kasten sagt, welche Bedingungen `false` ergeben, welche als erste `true` ergibt und welche nicht mehr ausgewertet werden.
- Beispiel mit 85 Punkten: `85 >= 90` ergibt `false`, `85 >= 75` ergibt `true`. Die Ausgabe ist „Note: Gut“. Die Bedingungen in den Zeilen 7 und 9 würden auch `true` ergeben, werden aber nicht mehr ausgewertet.
- Am Ende steht ein Link zur Seite über die switch-Anweisung.

### Handgriffe

- `punkte` mit den Knöpfen −10, −1, +1 und +10 verändern. Erlaubt sind 0 bis 100.
- Die Programmzeilen lassen sich nicht ändern. Die Grenzen 90, 75, 60 und 50 stehen fest, ebenso ihre Reihenfolge.
- Dem Link zur Seite „switch-Anweisung“ folgen.

### Abschnitte des Kurstexts

- 5.5.1
- 5.5.2

## Schritt 4: Schaltjahr

Reiter: schaltjahr

### Aufbau der Darstellung

- Oben steht die Schaltjahr-Regel des Gregorianischen Kalenders: Durch 4 teilbare Jahre sind Schaltjahre. Durch 100 teilbare Jahre sind keine. Durch 400 teilbare Jahre sind doch Schaltjahre.
- Die Demo zeigt Listing 5.2 des Kurstexts, `schaltjahr.go`, mit 34 Zeilen und denselben Zeilennummern wie im Kurstext.
- Das Programm hat drei if-Anweisungen. Zeilen 13 bis 15: ohne `else`, gibt für Jahre vor 1582 einen Hinweis aus. Zeilen 18 bis 26: Kette mit `else if`, setzt die Variable `schaltjahr`. Zeilen 28 bis 32: mit `else`, gibt das Ergebnis aus.
- Für ein Jahr vor 1582 gibt das Programm den Hinweis aus und rechnet trotzdem weiter.
- Der Student geht den Lauf Zeile für Zeile durch. Die gelb markierte Zeile ist die Zeile, die gerade ausgeführt wurde. Zeilen, die schon gelaufen sind, haben einen grünen Strich.
- Leerzeilen, die Kommentarzeile 17 und Zeilen, die nur `}` oder `} else {` enthalten, sind kein eigener Halt im Lauf. Der Lauf beginnt bei Zeile 7.
- Am Ende des Laufs sind die übersprungenen Zeilen der Zweige blass. Das sind Zeilen aus 14, 19, 21, 23, 25, 29 und 31.
- Über dem Programm stehen die drei Knöpfe für den Lauf und drei Angaben: der Wert von `jahr`, der Wert von `schaltjahr` und „Ausgabe bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile. Dieser Teil bleibt beim Rollen oben im Fenster stehen. Am Ende des Laufs sagt er, ob das Ergebnis mit der Schaltjahr-Regel übereinstimmt.
- `jahr` hat bis Zeile 11 den Null-Wert 0. In Zeile 11 liest `fmt.Scan` die Eingabe.
- Die Kette gibt es in zwei Reihenfolgen. „Reihenfolge des Kurstexts“: Zeile 18 prüft `jahr%400 == 0`, Zeile 20 prüft `jahr%100 == 0`, Zeile 22 prüft `jahr%4 == 0`. „Falsche Reihenfolge“: Zeile 18 prüft `jahr%4 == 0`, Zeile 20 prüft `jahr%100 == 0`, Zeile 22 prüft `jahr%400 == 0`. Die falsche Reihenfolge steht nicht im Kurstext, die Seite hat sie zum Vergleich gebaut.
- In der falschen Reihenfolge nennt das Programm jedes durch 4 teilbare Jahr ein Schaltjahr, auch 1700 und 1900. Für 2024, 2023 und 2000 stimmt das Ergebnis trotzdem.
- Unter der Demo stehen die Teile „Warum die Reihenfolge zählt“ mit dem Jahr 1700 und „Zwei Zeichen im Programm“ mit `%` und `&jahr`.
- `%` liefert den Rest der ganzzahligen Division. `&jahr` gibt `fmt.Scan` Zugriff auf die Variable. Das Zeichen `&` wird im Kurstext erst in Abschnitt 7.4.1 erklärt.
- Der Schritt endet mit einem Link zum Quiz.

### Handgriffe

- Das Jahr wählen: mit einem der Knöpfe 2024, 2023, 2000, 1900, 1700 und 1500 oder im Zahlenfeld daneben. Erlaubt sind 1 bis 9999.
- Die Reihenfolge wählen: „Reihenfolge des Kurstexts: 400, 100, 4“ oder „falsche Reihenfolge: 4, 100, 400“. Andere Reihenfolgen gibt es nicht.
- „Nächste Zeile“ führt die nächste Zeile aus. „Bis zum Ende“ führt den Rest aus. „Von vorn“ beginnt wieder bei Zeile 7. Einen Schritt zurück gibt es nicht.
- Ein neues Jahr oder eine andere Reihenfolge beginnt den Lauf von vorn.
- Die Programmzeilen lassen sich nicht ändern.
- Dem Link zum Quiz folgen. Das Quiz ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.1

## Schritt 5: Aus der Praxis

Reiter: praxis

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und Codeblöcken. Eine Demo gibt es hier nicht.
- Oben sagt ein gelber Kasten: Die Faustregeln stammen vom Mentor und stehen nicht in Abschnitt 5.5.1 des Kurstexts. Go übersetzt in jedem Beispiel beide Fassungen.
- Stoff des Kurses ist, was im Kurstext steht: die if-Anweisung aus Abschnitt 5.5.1, wie sie Schritt 1 bis 4 zeigen. Die vier Faustregeln dieses Schritts sind Ratschläge zum Stil und gehören nicht zu diesem Stoff.
- Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.
- Faustregel 1, die Ausnahme steht vor der Regel: In einer Kette steht die engste Bedingung oben. Das Beispiel prüft `punkte >= 50` vor `punkte >= 90`. Für 95 Punkte gibt es „Note: Ausreichend“ aus, der zweite Zweig kann nie laufen.
- Faustregel 2, eine boolesche Variable braucht keinen Vergleich: `if schaltjahr {` statt `if schaltjahr == true {`. Für das Gegenteil `if !schaltjahr {`.
- Faustregel 3, eine lange Bedingung bekommt einen Namen: Die Schaltjahr-Regel als ein Ausdruck, `jahr%400 == 0 || (jahr%4 == 0 && jahr%100 != 0)`, einmal direkt hinter `if`, einmal in der Variable `istSchaltjahr`. Beide Fassungen sind in Ordnung.
- Faustregel 4, kein if im if, wenn eine Bedingung genügt: zwei verschachtelte if-Anweisungen mit `alter >= 18` und `hatAusweis`, daneben eine if-Anweisung mit `&&`.
- Am Ende steht ein Ausblick auf Kapitel 7: Ein Sonderfall lässt sich in einer Funktion vorab mit `if` und `return` erledigen.
- Die Codeblöcke sind fest und haben keine Knöpfe.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen.
- Dem Link zum Quiz folgen.

### Abschnitte des Kurstexts

- 5.5.1
- 4.5

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 if und Block, 5.5.1, Seiten 129–132: Sprachvergleiche 5.9 und 5.10: Die if-Anweisung mit Block und die Sprachvergleiche stammen aus dem Kurstext. Das Programm mit alter ist ein eigenes Beispiel.
  - 2 else, 5.5.1, Seiten 129–132: Die if-Anweisung mit else stammt aus dem Kurstext. Das Programm mit alter ist ein eigenes Beispiel.
  - 3 else if, 5.5.1, Seiten 129–132: Die Kette mit else if stammt aus dem Kurstext. Das Programm mit punkte ist ein eigenes Beispiel.
  - 4 Schaltjahr, 5.5.1, Seiten 129–132: Listing 5.2 und Exkurs 5.8 zu fmt.Scan: Listing 5.2 zum Schaltjahr stammt aus dem Kurstext. Die falsche Reihenfolge mit vertauschten Zeilen 18, 20 und 22 gibt die Lektion dazu.
  - 5 Aus der Praxis, 4.5, Seiten 93–98: &&, || und !: Die Operatoren für zusammengesetzte Bedingungen stammen aus Abschnitt 4.5. Die Faustregeln stammen aus der Praxis des Mentors.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.1, Seiten 5–8, Sprünge und Schleifen: Ohne Kontrollanweisungen laufen die Anweisungen der Reihe nach. Bedingte Sprünge und Schleifen ändern diese Reihenfolge. Lektion: Imperative Programmierung, Schritt 3.
  - Abschnitt 2.5, Seiten 13–14, Flussdiagramm und Struktogramm: Die Verzweigung als Bild. Lektion: Größter gemeinsamer Teiler.
  - Abschnitt 3.3.2, Seiten 27–29, Klammer nach if: Lektion: Fehlersuche: drei kurze Go-Programme, Schritt 3.
  - Abschnitt 4.5.2, Seiten 96–98, Vergleiche: Lektion: Logische Operatoren und Vergleiche in Go.
- Unter „Wie es weitergeht“ stehen Links auf: „Quiz zu den if-Anweisungen“.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.5
- 3.3.2
- 4.5
- 4.5.2
- 5.5.1
