# break-Anweisung in Go

Die Seite erklärt die Anweisung `break`, die eine Schleife sofort verlässt. Sie führt das Programm zur Fakultät aus Listing 5.7 Anweisung für Anweisung aus, zeigt die gleiche Schleife ohne `break` mit einer Flag-Variable in der Schleifenbedingung und sagt, wann `break` passt. Sie gehört zu Abschnitt 5.5.3.3 des Kurstexts.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- `break` im Schleifenrumpf einer for-Anweisung verlässt den Schleifenrumpf sofort und beendet die for-Anweisung. Die Post-Anweisung läuft danach nicht mehr, und die Schleifenbedingung wird nicht mehr geprüft.
- `break` ist nur in einer Schleife und in einer switch-Anweisung erlaubt. Es verlässt nur die innerste Schleife oder switch-Anweisung, in der es steht.
- Die Fakultät n! ist das Produkt aller ganzen Zahlen von 1 bis n. 0! ist 1.
- `math.MaxInt` ist der größte Wert des Typs `int`: 9223372036854775807.
- 20! = 2432902008176640000 passt noch in den Typ `int`. 21! passt nicht mehr.
- Eine Flag-Variable ist eine Variable vom Typ `bool`, die einen Zustand festhält. Hier heißt sie `überlauf`.
- Alle Schritte folgen dem Kurstext. Einen Reiter mit Faustregeln des Mentors hat diese Seite nicht, weil die Ratschläge zum Stil in Abschnitt 5.5.3.3 selbst stehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Codeblock zeigt eine Schleife `for k := 1; k <= 5; k++ {` mit `if k == 3 { break }` und danach `fmt.Print(k, " ")`. Das Programm gibt `1 2` aus.
- Eine Liste nennt drei frühere Abschnitte des Kurstexts: 5.5.3.2 mit der for-Klausel, 5.5.3.1 mit der Flag-Variable, 4.3.8 mit dem arithmetischen Überlauf.
- Eine zweite Liste sagt, was die vier Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „for-Anweisung mit for-Klausel“, „for-Anweisung mit Bedingung“ und „Überlauf bei Ganzzahlen“ folgen. Alle drei sind eigene HTML-Seiten.
- Der Codeblock lässt sich nicht ändern und nicht ausführen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 4.

### Abschnitte des Kurstexts

- 5.5.3.3
- 5.5.3.2

## Schritt 1: Fakultät und Überlauf

Reiter: pruefung

### Aufbau der Darstellung

- Oben steht, was die Fakultät ist, mit dem Beispiel 5! = 120 und dem Sonderfall 0! = 1.
- Der erste Codeblock zeigt die Schleife ohne Prüfung: `fak := 1`, `for i := 1; i <= eingabe; i++ {`, `fak *= i`, `}`.
- Der Teil „Das Problem“ sagt: Die Fakultät wächst sehr schnell. Schon 21! ist größer als `math.MaxInt`. Rechnet das Programm trotzdem weiter, kommt es zum arithmetischen Überlauf, und Go liefert ohne Meldung einen falschen Wert.
- Der zweite Codeblock zeigt die Prüfung vor der Multiplikation: `if math.MaxInt/fak >= i {` und darin `fak *= i`.
- Der Gedanke der Prüfung: `math.MaxInt/fak` sagt, wie oft `fak` in den größten Wert hineinpasst. Ist diese Zahl mindestens so groß wie `i`, dann ist `fak * i` höchstens `math.MaxInt`.
- Die Demo heißt „Ausprobieren: die Prüfung für einen Schleifendurchlauf“. Vier Knöpfe wählen den Schleifendurchlauf mit `i` gleich 5, 13, 20 oder 21. Der gewählte Knopf ist grün.
- Darunter stehen vier Angaben: „fak vor der Multiplikation“, „math.MaxInt/fak“, „math.MaxInt/fak >= i“ und „fak * i“. Darunter erklärt ein grüner Kasten die Werte.
- `i` gleich 5: `fak` ist 24, der Quotient ist 384307168202282325, die Prüfung ergibt `true`, das Produkt ist 120.
- `i` gleich 13: `fak` ist 479001600, der Quotient ist 19255409662, die Prüfung ergibt `true`, das Produkt ist 6227020800.
- `i` gleich 20: `fak` ist 121645100408832000, der Quotient ist 75, die Prüfung ergibt `true`, das Produkt ist 2432902008176640000.
- `i` gleich 21: `fak` ist 2432902008176640000, der Quotient ist 3, die Prüfung ergibt `false`. Das Produkt passt nicht in `int`. Go rechnete -4249290049419214848. Diese zwei Angaben stehen in roter Schrift.
- `math.MaxInt/fak` ist eine Division ganzer Zahlen, die Nachkommastellen fallen weg.

### Handgriffe

- Den Schleifendurchlauf wählen: nur mit den Knöpfen 5, 13, 20 und 21. Andere Werte von `i` gibt es in der Demo nicht.
- Die Demo zeigt nur die Werte eines Schleifendurchlaufs. Den ganzen Lauf zeigt Schritt 2.
- Die zwei Codeblöcke sind fest und haben keine Knöpfe.

### Abschnitte des Kurstexts

- 5.5.3.3
- 4.3.8

## Schritt 2: break im Lauf

Reiter: lauf

### Aufbau der Darstellung

- Oben steht, was `break` bewirkt: Der Schleifenrumpf wird sofort verlassen, und die for-Anweisung ist sofort zu Ende.
- Die Demo zeigt Listing 5.7 des Kurstexts, `for_klausel_fakultaet.go`, mit 36 Zeilen und Zeilennummern.
- Zeile 9: `eingabe := 0`. Zeile 10: `überlauf := false`. Zeile 12 gibt die Aufforderung „Fakultät berechnen (natürliche Zahl eingeben):“ aus. Zeile 13: `fmt.Scan(&eingabe)`.
- Zeile 15: `if eingabe < 0 {`. Zeile 16 meldet eine ungültige Eingabe. Zeile 17: `} else {`.
- Zeile 18: `fak := 1`. Zeile 19: `for i := 1; i <= eingabe; i++ {`. Zeile 20 ist ein Kommentar.
- Zeile 21: `if math.MaxInt/fak >= i {`. Zeile 22: `fak *= i`. Zeile 23: `} else {`. Zeile 24 meldet den Überlauf. Zeile 25: `überlauf = true`. Zeile 26: `break // Schleife sofort verlassen`.
- Zeile 28 ist die schließende Klammer der Schleife. Zeile 29: `if überlauf {`. Zeile 30 meldet, dass die Fakultät zu groß ist. Zeile 32 gibt die Fakultät aus. Zeile 36 ist die schließende Klammer von `main`.
- Der Student spielt den Benutzer. Der Lauf hält in Zeile 13 an und wartet. Dann wählt der Student eine der fünf Zahlen 0, 3, 5, 21 und −2.
- Eingabe 3: Der Schleifenrumpf läuft dreimal. Ausgabe: „Fakultät von 3 ist 6.“
- Eingabe 5: Der Schleifenrumpf läuft fünfmal. Ausgabe: „Fakultät von 5 ist 120.“
- Eingabe 0: `1 <= 0` ergibt bei der ersten Prüfung `false`, der Schleifenrumpf läuft nie. Ausgabe: „Fakultät von 0 ist 1.“
- Eingabe −2: Zeile 15 ergibt `true`. Ausgabe: „Eingabe -2 ist ungültig.“ Die Schleife läuft nie.
- Eingabe 21: Zwanzig Schleifendurchläufe multiplizieren. Im 21. Schleifendurchlauf ergibt die Prüfung in Zeile 21 `false`. Dann laufen die Zeilen 24, 25 und 26. `break` in Zeile 26 verlässt die Schleife, das Programm macht in Zeile 29 weiter. Die Ausgabe hat zwei weitere Zeilen: „Achtung: 2432902008176640000 * 21 führt zu Überlauf: -4249290049419214848“ und „Fakultät von 21 ist zu groß, um mit int berechnet werden zu können.“
- Nach dem `break` bleibt `i` bei 21. Die Post-Anweisung `i++` läuft nicht mehr.
- Die negative Zahl in der Meldung ist `fak*i`, so wie Go es beim Überlauf falsch berechnet. Der Variable `fak` wird dieses Produkt nicht zugewiesen, `fak` bleibt 2432902008176640000.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wurde oder an der das Programm wartet. Ein grüner Strich links bedeutet: Die Zeile ist in diesem Lauf schon gelaufen.
- Hinter den Bedingungen in den Zeilen 15, 21 und 29 steht ein kleines Schild mit dem Ergebnis ihrer letzten Prüfung: grün mit `true` oder rot mit `false`.
- Zeile 19 kommt in drei Rollen dran. Hinter ihr steht dann ein Schild: grau „Init-Anweisung“, grün „Schleifenbedingung: true“ oder rot „Schleifenbedingung: false“, grau „Post-Anweisung“.
- Über dem Programm stehen fünf Angaben: `eingabe`, `überlauf`, `fak`, `i` und „Schleifendurchläufe bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile. Unter dem Programm steht ein schwarzer Kasten „Ausgabe bisher“.
- Unter der Demo stehen zwei Teile: „Wozu die Variable überlauf dient“ und „Zwei Stellen im Programm“ mit der Eingabe 0 und der negativen Eingabe.
- `&eingabe` gibt `fmt.Scan` Zugriff auf die Variable. Das Zeichen `&` wird im Kurstext erst in Abschnitt 7.4.1 erklärt.

### Handgriffe

- „Nächste Anweisung“ führt die nächste Anweisung aus.
- „Bis zur nächsten Prüfung in Zeile 19“ führt alle Anweisungen bis zur nächsten Prüfung der Schleifenbedingung aus. Mit der Eingabe 21 spart das viele Klicks.
- „Bis zum Ende“ führt den Rest aus.
- Wartet das Programm in Zeile 13, sind diese drei Knöpfe gesperrt. Dann wählt der Student eine Eingabe, nur mit den Knöpfen 0, 3, 5, 21 und −2. Ein Eingabefeld gibt es nicht, andere Zahlen gehen nicht. Die fünf Knöpfe sind nur wählbar, während das Programm in Zeile 13 wartet.
- „Von vorn“ löscht die Eingabe und beginnt wieder bei Zeile 9. Einen Schritt zurück gibt es nicht.
- Ein Lauf nimmt genau eine Eingabe an.
- Die Programmzeilen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.5.3.3

## Schritt 3: Ohne break

Reiter: flag

### Aufbau der Darstellung

- Oben steht: Die bessere Lösung kommt ohne `break` aus. Die Flag-Variable `überlauf` wandert in die Schleifenbedingung.
- Der Codeblock zeigt den neuen Kopf der Schleife: `for i := 1; i <= eingabe && !überlauf; i++ {`.
- `&&` ist der Operator für „und“. `!überlauf` kehrt den Wahrheitswert von `überlauf` um und ist `true`, solange kein Überlauf festgestellt wurde.
- Die Demo heißt „Vergleich: was nach dem Überlauf geschieht“. Sie gilt für die Eingabe 25 und den Schleifendurchlauf mit `i` gleich 21. Zwei Knöpfe wählen das Programm, der gewählte Knopf ist grün.
- Unter den Knöpfen steht die Schleife des gewählten Programms in einem dunklen Kasten ohne Zeilennummern, darunter erklärt ein grüner Kasten, was nach dem Überlauf geschieht.
- „mit break“: Nach `überlauf = true` folgt `break`. Die Schleife ist sofort zu Ende, `i` bleibt 21.
- „mit Flag in der Schleifenbedingung“: Nach `überlauf = true` ist der Schleifenrumpf zu Ende. Die Post-Anweisung erhöht `i` auf 22. Die Prüfung `22 <= 25 && !überlauf` ergibt `false`, weil `!überlauf` jetzt `false` ist. Die Schleife endet.
- Beide Programme geben für jede Eingabe das Gleiche aus.
- Im Kurstext steht vom Programm ohne `break` nur der Kopf der Schleife. Den Schleifenrumpf hat die Seite aus Listing 5.7 übernommen.

### Handgriffe

- Das Programm wählen: „mit break“ oder „mit Flag in der Schleifenbedingung“.
- Kein Code dieses Schritts lässt sich ändern oder ausführen. Einen Lauf Anweisung für Anweisung gibt es hier nicht. Eine Eingabe lässt sich hier nicht wählen, der Vergleich gilt für die Eingabe 25.
- Den Links zu den Seiten „for-Anweisung mit Bedingung“ und „Logische Operatoren und Vergleiche“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.3.3
- 5.5.3.1

## Schritt 4: Wann break passt

Reiter: wann

### Aufbau der Darstellung

- Oben steht, warum `break` oft als schlechter Programmierstil gilt: Mit `break` entscheidet nicht mehr allein die Schleifenbedingung, wann die Schleife endet. Der Leser muss den ganzen Schleifenrumpf durchsuchen.
- Die Regel im Kurstext: Eine Schleife soll möglichst ohne `break` auskommen.
- Der Teil „Wann break trotzdem in Ordnung ist“ nennt den Fall, dass hinter dem Grund für den Abbruch noch Anweisungen im Schleifenrumpf folgen, die nicht mehr laufen sollen. Wer `break` verwendet, schreibt einen Kommentar an den Anfang der Schleife.
- Die Demo heißt „Ausprobieren: was break verlässt“. Vier Knöpfe wählen vier Programme. Der gewählte Knopf ist grün.
- Unter den Knöpfen steht das Programm mit Zeilennummern. Gelb markiert ist die Zeile mit `break`. Darunter stehen „Ergebnis“ oder „Meldung des Compilers“ und „Ausgabe“, darunter erklärt ein grüner Kasten das Ergebnis.
- „in einer Schleife“: `for k := 1; k <= 5; k++ {` mit `if k == 3 { break }` und `fmt.Print(k, " ")`. Ausgabe: `1 2`.
- „in zwei geschachtelten Schleifen“: außen `i` von 1 bis 3, innen `j` von 1 bis 3, in der inneren Schleife `if j == 2 { break }`. Ausgabe: drei Zeilen, `i, j: 1 1`, `i, j: 2 1`, `i, j: 3 1`. `break` verlässt nur die innere Schleife.
- „in switch in einer Schleife“: `for i := 1; i <= 5; i++ {` mit `switch i { case 3: break }` und `fmt.Print(i, " ")`. Ausgabe: `1 2 3 4 5`. `break` verlässt nur die switch-Anweisung, die Schleife läuft weiter.
- „außerhalb einer Schleife“: `if x > 3 { break }` ohne Schleife. Meldung des Compilers: `break is not in a loop, switch, or select`.
- Im Kurstext steht, dass `break` nur in einer Schleife und in einer switch-Anweisung zulässig ist. Die vier Programme und der Fall der geschachtelten Schleifen stehen nicht im Kurstext.
- Am Ende des Schritts unter „Wie es weitergeht“ steht ein Link zum Quiz. Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.

### Handgriffe

- Eines der vier Programme wählen: nur mit den vier Knöpfen.
- Die Programmzeilen lassen sich nicht ändern. Einen Lauf Anweisung für Anweisung gibt es in diesem Schritt nicht.
- Dem Link zum Quiz folgen. Es ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.3.3
- 5.5.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Fakultät und Überlauf, 5.5.3.3, Seiten 146–148: Die Fakultät mit der Prüfung vor jeder Multiplikation stammt aus dem Kurstext. Die Werte der Demo stammen aus einem Lauf mit Go 1.26.
  - 2 break im Lauf, 5.5.3.3, Seiten 146–148: Listing 5.7: Listing 5.7 stammt aus dem Kurstext. Die Ausführung Anweisung für Anweisung gibt die Lektion dazu.
  - 3 Ohne break, 5.5.3.3, Seite 148: Der Kopf der Schleife mit der Flag-Variable stammt aus dem Kurstext. Den Schleifenrumpf hat die Lektion aus Listing 5.7 übernommen.
  - 4 Wann break passt, 5.5.3.3, Seiten 147–149: Der Rat, eine Schleife möglichst ohne break zu schreiben, stammt aus dem Kurstext. Die vier Programme der Demo und die Meldung des Compilers gibt die Lektion dazu. Dass break bei geschachtelten Schleifen nur die innere verlässt, steht nicht im Kurstext.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.1, Seiten 5–8, Sprünge und Schleifen: Ohne Kontrollanweisungen laufen die Anweisungen der Reihe nach. Bedingte Sprünge und Schleifen ändern diese Reihenfolge. Lektion: Imperative Programmierung, Schritt 3.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Warum die Fakultät geprüft wird. Lektion: Überlauf bei Ganzzahlen in Go.
  - Abschnitt 5.5.3.2, Seiten 142–146, Endlosschleife for {: Sie endet nur mit break. Lektion: for-Anweisung mit for-Klausel in Go, Schritt 4.
  - Abschnitt 5.5.2, Seiten 133–137, switch: Ein break in einer switch-Anweisung verlässt nur diese. Lektion: switch-Anweisung in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 4.3.8
- 5.5.2
- 5.5.3.2
- 5.5.3.3
