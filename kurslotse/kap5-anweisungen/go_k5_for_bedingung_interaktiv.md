# for-Anweisung mit Bedingung in Go

Die Seite erklärt die einfachste Schleife in Go: `for`, eine Schleifenbedingung und einen Schleifenrumpf. Sie führt drei Programme Anweisung für Anweisung aus, darunter das Programm zur wiederholten Eingabe aus Listing 5.4. Sie gehört zu Abschnitt 5.5.3.1 des Kurstexts.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Ein Programm einer Demo steht in einem dunklen Kasten mit Zeilennummern am linken Rand.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wurde. Ein grüner Strich links an einer Programmzeile bedeutet: Die Zeile ist in diesem Lauf schon gelaufen.
- Hinter einer Bedingung steht ein kleines Schild mit dem Ergebnis ihrer letzten Prüfung: grün mit `true` oder rot mit `false`.
- Die Schleifenbedingung ist der boolesche Ausdruck hinter `for`. Der Schleifenrumpf ist der Block zwischen `{` und `}`. Ein Schleifendurchlauf ist eine Ausführung des ganzen Schleifenrumpfs.
- Go prüft die Schleifenbedingung vor jedem Schleifendurchlauf, auch vor dem ersten. Die Schleifenbedingung wird deshalb einmal öfter geprüft, als der Schleifenrumpf ausgeführt wird.
- Schritt 1 bis 4 folgen dem Kurstext. Schritt 5 enthält Faustregeln des Mentors, die nicht im Kurstext stehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Oben steht, wozu eine Schleife dient, und die einfachste Form in Go: `for n > 0 {` mit einem Schleifenrumpf aus `fmt.Println(n)` und `n--`.
- Eine Liste nennt die drei Arten der for-Anweisung: mit Bedingung (Abschnitt 5.5.3.1, Thema dieser Seite), mit for-Klausel (Abschnitt 5.5.3.2), mit `range` (Abschnitt 6.2).
- Eine zweite Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 2.1 mit `GOTO` und `WHILE b {…}`, Abschnitt 2.5 mit dem größten gemeinsamen Teiler als Flussdiagramm und Struktogramm, Abschnitt 5.5.1 mit der if-Anweisung.
- Eine dritte Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „for-Anweisung mit for-Klausel“, „Imperative Programmierung“, „Größter gemeinsamer Teiler“ und „if-Anweisung“ folgen. Alle vier sind eigene HTML-Seiten.
- Der Codeblock lässt sich nicht ändern und nicht ausführen.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.3.1
- 2.1
- 2.5

## Schritt 1: Bedingung und Rumpf

Reiter: begriffe

### Aufbau der Darstellung

- Oben stehen fünf Begriffe: Schleifenbedingung, Schleifenrumpf, Schleifendurchlauf, Abbruchbedingung, Endlosschleife.
- Die Abbruchbedingung ist das Gegenteil der Schleifenbedingung. Zu `n > 0` gehört `n <= 0`.
- Darunter steht, wann Go die Schleifenbedingung prüft: vor jedem Schleifendurchlauf. Ergibt sie schon vor dem ersten `false`, läuft der Schleifenrumpf kein einziges Mal.
- Die Demo zeigt ein Programm mit sieben Zeilen. Zeile 1 deklariert `n` mit dem gewählten Startwert. Zeile 3 ist `for n > 0 {`. Zeile 4 gibt `n` aus. Zeile 5 ist `n--`. Zeile 6 ist `}`. Zeile 7 steht hinter der Schleife und gibt „Ende“ aus.
- Das Programm gibt es in zwei Fassungen. In der Fassung „ohne n--“ steht in Zeile 5 nur der Kommentar `// hier fehlt n--`. Zu sehen ist immer nur eine Fassung, ihr Knopf ist grün.
- Der Student geht den Lauf Anweisung für Anweisung durch. Der Lauf hält bei den Zeilen 1, 3, 4, 5 und 7. Zeile 3 kommt bei jeder Prüfung wieder dran.
- Über dem Programm stehen vier Angaben: der Wert von `n`, „Prüfungen der Bedingung bisher“, „Schleifendurchläufe bisher“ und „Ausgabe bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile.
- Beispiel mit Startwert 3 und `n--`: Die Ausgabe ist 3, 2, 1, Ende. Der Schleifenrumpf läuft dreimal, die Schleifenbedingung wird viermal geprüft.
- Mit Startwert 0 ergibt die erste Prüfung `false`. Der Schleifenrumpf läuft kein einziges Mal, die Ausgabe ist nur „Ende“. Das gilt in beiden Fassungen.
- In der Fassung „ohne n--“ ist das Programm mit jedem Startwert von 1 bis 5 eine Endlosschleife: `n` ändert sich nie. Die Demo hält nach fünf Schleifendurchläufen bei der sechsten Prüfung an und sagt das. Ein echtes Programm liefe weiter.
- Unter der Demo steht der Teil „Woran man eine Endlosschleife erkennt“. Dort steht auch: Im Terminal beendet man ein solches Programm mit den Tasten Strg und C.

### Handgriffe

- Den Startwert von `n` wählen: nur mit den Knöpfen 0, 1, 2, 3, 4 und 5. Ein Eingabefeld gibt es nicht, andere Werte gehen nicht.
- Die Fassung wählen: „mit n--“ oder „ohne n--“.
- „Nächste Anweisung“ führt die nächste Anweisung aus. „Bis zum Ende“ führt den Rest aus. „Von vorn“ beginnt wieder bei Zeile 1. Einen Schritt zurück gibt es nicht.
- Ein neuer Startwert oder eine andere Fassung beginnt den Lauf von vorn.
- Die Programmzeilen lassen sich nicht ändern. Die Bedingung `n > 0` steht fest.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.3.1

## Schritt 2: Wiederholte Eingabe

Reiter: eingabe

### Aufbau der Darstellung

- Die Demo zeigt Listing 5.4 des Kurstexts, `for_bedingung_eingabe.go`, mit 18 Zeilen.
- Zeile 9 deklariert `eingabe` als `float64` mit dem Wert `1.0`. Zeile 11 ist die Schleife: `for eingabe != 0 {`. Zeile 12 gibt die Aufforderung „Bitte eine positive Zahl eingeben (0 zum Beenden):“ aus. Zeile 13 liest mit `fmt.Scan(&eingabe)` eine Zahl. Zeile 14 ist `if eingabe > 0 {`. Zeile 15 gibt mit `fmt.Printf` und `math.Sqrt` die Wurzel aus. Zeile 17 ist die schließende Klammer der Schleife.
- Der Student spielt den Benutzer. Der Lauf hält in Zeile 13 an und wartet. Dann wählt der Student eine der fünf Zahlen 16, 25, 2, −4 und 0.
- Die Wurzeln, die das Programm ausgibt: 4 für 16, 5 für 25, 1.4142135623730951 für 2. Für −4 und für 0 gibt es keine Wurzel aus, weil `eingabe > 0` dann `false` ergibt.
- Die Eingabe 0 beendet die Schleife: Die nächste Prüfung `0 != 0` ergibt `false`.
- Der Lauf hält bei den Zeilen 9, 11, 12, 13, 14, 15 und 17. Leerzeilen, die Zeilen 1 bis 8 und Zeilen, die nur `}` enthalten, sind kein eigener Halt. Zeile 17 ist der Halt am Ende.
- Über dem Programm stehen drei Angaben: der Wert von `eingabe`, „Schleifendurchläufe bisher“ und „Eingaben bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile.
- Unter dem Programm steht ein schwarzer Kasten „Ausgabe bisher“. Er zeigt, was auf dem Bildschirm steht. Eine Zeile mit `>` davor ist eine Eingabe des Benutzers.
- Zeile 9 gibt es in zwei Fassungen. „wie im Kurstext“: `= 1.0`. „geändert“: `= 0.0`. Die Fassung mit 0.0 steht nicht im Kurstext, die Seite hat sie zum Vergleich gebaut.
- Mit 0.0 ergibt die erste Prüfung `0 != 0` den Wert `false`. Das Programm überspringt den Schleifenrumpf und endet, ohne zu fragen. Der Lauf hat dann nur drei Halte: Zeile 9, 11 und 17.
- Die 1.0 dient nur dazu, dass die erste Prüfung `true` ergibt. Jeder andere Wert ungleich 0 täte dasselbe.
- Unter der Demo stehen drei Teile: „Warum eingabe mit 1.0 beginnt“, „Drei Stellen im Programm“ mit `if`, `math.Sqrt` und `&eingabe`, und „Das zweite Beispiel des Kurstexts“ mit zwei Links zum größten gemeinsamen Teiler aus Listing 5.5.
- `&eingabe` gibt `fmt.Scan` Zugriff auf die Variable. Das Zeichen `&` wird im Kurstext erst in Abschnitt 7.4.1 erklärt.
- Diese Art der Schleife ist für den Fall gedacht, dass die Zahl der Schleifendurchläufe vor dem Eintritt in die Schleife nicht feststeht.

### Handgriffe

- Die Fassung von Zeile 9 wählen: „= 1.0“ oder „= 0.0“. Das beginnt den Lauf von vorn.
- „Nächste Anweisung“ führt die nächste Anweisung aus. Wartet das Programm in Zeile 13, ist der Knopf gesperrt.
- Wartet das Programm in Zeile 13: eine Eingabe wählen, nur mit den Knöpfen 16, 25, 2, −4 und 0. Ein Eingabefeld gibt es nicht, andere Zahlen gehen nicht. Die fünf Knöpfe sind nur wählbar, während das Programm in Zeile 13 wartet.
- Ein Lauf nimmt höchstens sechs Eingaben an. Nach fünf Eingaben ist nur noch der Knopf 0 wählbar.
- „Von vorn“ löscht die Eingaben und beginnt wieder bei Zeile 9. Einen Schritt zurück und einen Knopf „Bis zum Ende“ gibt es hier nicht.
- Die Programmzeilen lassen sich nicht ändern.
- Den Links zur Seite „Größter gemeinsamer Teiler“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.3.1

## Schritt 3: while in anderen Sprachen

Reiter: while

### Aufbau der Darstellung

- Oben steht: Viele Sprachen leiten diese Schleife mit `while` ein, sie heißt deshalb allgemein while-Schleife. Go hat sie auch und schreibt `for`.
- Der Schritt hat zwei Auswahlen mit Knöpfen. Sie hängen nicht zusammen. Der gewählte Knopf ist grün.
- Obere Auswahl, „Dieselbe Schleife in vier Sprachen“: Knöpfe Go, Java, Python und C++. Darunter steht die Schleife in der gewählten Sprache in einem dunklen Kasten ohne Zeilennummern, darunter eine Erklärung. Jede Fassung gibt die Zahlen von `i` bis 9 aus.
- Go: `for i < 10 {`. Java und C++: `while (i < 10) {`, die runden Klammern sind dort Pflicht. Python: `while i < 10:`, der Schleifenrumpf ist eingerückt, und statt `i++` steht `i += 1`.
- Nur der Code unter dem Knopf „Go“ ist Go. Die drei anderen Fassungen sind Java, Python und C++.
- Untere Auswahl, „Was Go bei der Schreibweise verlangt“: vier Knöpfe mit vier Schreibweisen derselben Schleife in Go. Darunter steht die Schreibweise, darunter „Ergebnis“ oder „Meldung des Compilers“ in roter Schrift, darunter eine Erklärung.
- „wie üblich“: `for i < 10 {`. Go übersetzt das Programm.
- „mit runden Klammern“: `for (i < 10) {`. Go übersetzt das Programm. Nach Idiomatic Go lässt man die Klammern weg, das Werkzeug `gofmt` entfernt sie.
- „mit while“: `while i < 10 {`. Der Compiler meldet `syntax error: unexpected name i at end of statement`. Go kennt das Schlüsselwort `while` nicht und hält es für einen gewöhnlichen Namen. Deshalb nennt die Meldung das `i`.
- „Zahl als Bedingung“: `for i {`. Der Compiler meldet `non-boolean condition in for statement`. Die Schleifenbedingung muss den Typ `bool` haben.

### Handgriffe

- Obere Auswahl: eine der vier Sprachen wählen. Andere Sprachen gibt es nicht.
- Untere Auswahl: eine der vier Schreibweisen wählen. Andere Schreibweisen gibt es nicht.
- Kein Code dieses Schritts lässt sich ändern oder ausführen. Einen Lauf Anweisung für Anweisung gibt es hier nicht.

### Abschnitte des Kurstexts

- 5.5.3.1
- 2.8

## Schritt 4: Mindestens ein Durchlauf

Reiter: flag

### Aufbau der Darstellung

- Oben stehen zwei Schleifen anderer Sprachen, die erst nach dem Schleifenrumpf prüfen. Ihr Schleifenrumpf läuft deshalb mindestens einmal.
- do-while-Schleife in Java und C++: Sie endet, wenn die Bedingung hinter `while` nicht mehr zutrifft. Der Codeblock dazu ist Java.
- repeat-until-Schleife in Pascal: Sie endet, wenn die Bedingung hinter `until` zutrifft. Die Bedingung ist also umgekehrt formuliert: `until eingabe = 0` statt `while (eingabe != 0)`. Der Codeblock dazu ist Pascal.
- Go hat keine dieser zwei Schleifen.
- Der Nachbau in Go nimmt eine Flag-Variable: eine Variable vom Typ `bool`, die einen Zustand festhält. Hier heißt sie `ersterLauf`.
- Die Demo zeigt ein Programm mit elf Zeilen. Zeile 1: `var eingabe float64 = 0.0`. Zeile 2: `var ersterLauf bool = true`. Zeile 4: `for ersterLauf || eingabe != 0 {`. Zeile 5 gibt die Aufforderung aus. Zeile 6 liest mit `fmt.Scan(&eingabe)`. Zeile 7 ist `if eingabe > 0 {`. Zeile 8 gibt die Wurzel aus. Zeile 10: `ersterLauf = false`. Zeile 11 ist die schließende Klammer der Schleife.
- Das Programm ist ein Ausschnitt ohne `package`, `import` und `func main`. Im Kurstext steht an der Stelle der Zeilen 5 bis 9 nur ein Kommentar, die Seite hat die Zeilen aus Listing 5.4 eingesetzt.
- `eingabe` beginnt hier mit 0.0. Trotzdem läuft der Schleifenrumpf einmal: Bei der ersten Prüfung hat `ersterLauf` den Wert `true`, und `true || …` ergibt `true`.
- Ab der zweiten Prüfung hat `ersterLauf` den Wert `false`. Dann entscheidet nur noch `eingabe != 0`.
- Zeile 10 wird in jedem Schleifendurchlauf ausgeführt. Ab dem zweiten ändert sie nichts mehr, weil `ersterLauf` schon `false` ist.
- Der Student spielt wieder den Benutzer. Der Lauf hält in Zeile 6 an und wartet. Die Eingaben und die Wurzeln sind dieselben wie in Schritt 2: 16, 25, 2, −4 und 0.
- Gibt der Student gleich als Erstes 0 ein, hat das Programm trotzdem einmal gefragt. Das ist der Unterschied zur Fassung mit `0.0` in Schritt 2.
- Über dem Programm stehen vier Angaben: `eingabe`, `ersterLauf`, „Schleifendurchläufe bisher“ und „Eingaben bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile. Unter dem Programm steht der schwarze Kasten „Ausgabe bisher“.
- Unter dem schwarzen Kasten steht eine Tabelle. Sie bekommt bei jeder Prüfung der Schleifenbedingung in Zeile 4 eine Zeile mit vier Spalten: die Nummer der Prüfung, der Wert von `ersterLauf`, das Ergebnis von `eingabe != 0` und das Ergebnis der ganzen Bedingung.
- In der ersten Zeile der Tabelle steht bei `eingabe != 0` „nicht ausgewertet“: Der linke Teil von `||` ergab schon `true`, den rechten Teil wertet Go dann nicht mehr aus. Das heißt Kurzschluss-Auswertung.
- Am Ende steht der Teil „Der schlechtere Weg: den Rumpf kopieren“: Aufforderung und `fmt.Scan` stehen einmal vor der Schleife und einmal im Schleifenrumpf. Das Programm läuft richtig, gilt aber als schlechter Stil.
- Der Schritt endet mit einem Link zum Quiz.

### Handgriffe

- „Nächste Anweisung“ führt die nächste Anweisung aus. Wartet das Programm in Zeile 6, ist der Knopf gesperrt.
- Wartet das Programm in Zeile 6: eine Eingabe wählen, nur mit den Knöpfen 16, 25, 2, −4 und 0. Ein Eingabefeld gibt es nicht, andere Zahlen gehen nicht.
- Ein Lauf nimmt höchstens sechs Eingaben an. Nach fünf Eingaben ist nur noch der Knopf 0 wählbar.
- „Von vorn“ löscht die Eingaben und beginnt wieder bei Zeile 1. Einen Schritt zurück und einen Knopf „Bis zum Ende“ gibt es hier nicht.
- Die Programmzeilen lassen sich nicht ändern. Die Anfangswerte `0.0` und `true` stehen fest.
- Die Codeblöcke in Java und Pascal und der Codeblock mit dem kopierten Rumpf sind fest und haben keine Knöpfe.
- Den Links zur Seite über die Kurzschluss-Auswertung und zum Quiz folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.3.1
- 2.1
- 4.5

## Schritt 5: Aus der Praxis

Reiter: praxis

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und Codeblöcken. Eine Demo gibt es hier nicht.
- Oben sagt ein gelber Kasten: Die Faustregeln stammen vom Mentor und stehen nicht in Abschnitt 5.5.3.1 des Kurstexts. Go übersetzt jedes Beispiel, auch die drei mit dem Kommentar „falsch“.
- Stoff des Kurses ist, was im Kurstext steht: die for-Anweisung mit Bedingung aus Abschnitt 5.5.3.1, wie sie Schritt 1 bis 4 zeigen. Die drei Faustregeln dieses Schritts sind Ratschläge des Mentors und gehören nicht zu diesem Stoff.
- Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt: Es steht weder auf dieser Seite noch im Kurstext. Ob eine Faustregel in der Klausur vorkommt oder nicht, lässt sich deshalb nicht sagen. Verbindlich sind der Kurstext und Moodle der FernUniversität.
- Faustregel 1, der Rumpf ändert, was die Bedingung prüft: `for x > 0` mit einem Schleifenrumpf, der nur `fmt.Println(x)` enthält. `x` bleibt 10, die Schleife endet nie. Die Prüffrage: Welche Anweisung im Schleifenrumpf sorgt dafür, dass die Bedingung irgendwann `false` ergibt?
- Faustregel 2, die Bedingung fängt auch den Wert ab, der das Ziel überspringt: `n` beginnt mit 5, der Schleifenrumpf ist `n -= 2`. Mit `for n != 0` wird `n` zu 3, 1, −1, −3 und trifft die 0 nie. Mit `for n > 0` endet die Schleife bei −1.
- Faustregel 3, eine errechnete Kommazahl nicht mit `!=` prüfen: `x` beginnt mit 0.0, der Schleifenrumpf ist `x += 0.1`. Nach zehn Schleifendurchläufen hat `x` den Wert 0.9999999999999999 und nach elf 1.0999999999999999. `for x != 1.0` endet nie. `for x < 1.0` endet, läuft aber elfmal.
- Der Grund ist der Rundungsfehler bei Gleitkommazahlen aus Abschnitt 4.4.3: 0.1 lässt sich im Zweiersystem nicht genau darstellen.
- Für genau zehn Schleifendurchläufe zählt das letzte Beispiel mit der ganzen Zahl `i` von 1 bis 10 und rechnet `x := float64(i) / 10`.
- `eingabe != 0` in Listing 5.4 ist in Ordnung, weil die 0 eingegeben und nicht errechnet wird.
- Der Teil „Wer zählt, nimmt die for-Klausel“ verweist auf Abschnitt 5.5.3.2.
- Der Teil „Ausblick auf break“ zeigt eine Schleife `for {` ohne Bedingung, die mit `if eingabe == 0 { break }` verlassen wird. Die Schleife ohne Bedingung steht im Kurstext im Expertenwissen 5.16, `break` in Abschnitt 5.5.3.3. Im Kurstext steht, dass eine Schleife möglichst ohne `break` auskommen soll, und für „mindestens ein Durchlauf“ wird dort die Flag-Variable verwendet.
- Die Codeblöcke sind fest und haben keine Knöpfe.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen.
- Den Links zu den Seiten „Gleitkommazahlen“, „break-Anweisung“, „if-Anweisung“ und zum Quiz folgen.

### Abschnitte des Kurstexts

- 5.5.3.1
- 4.4.3
- 5.5.3.2
- 5.5.3.3
