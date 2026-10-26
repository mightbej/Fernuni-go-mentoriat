# switch-Anweisung in Go

Die Seite erklärt die Mehrfachverzweigung mit `switch`. Sie führt das Programm zur Zuordnung der Wochentage aus Listing 5.3 Anweisung für Anweisung aus, vergleicht `switch` mit `if`, zeigt die switch-Anweisung ohne switch-Ausdruck und den Unterschied zu Java und C++. Sie gehört zu Abschnitt 5.5.2 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ausgaben und Meldungen des Compilers der Go-Programme stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Der switch-Ausdruck ist der Ausdruck hinter `switch`. Ein case-Abschnitt besteht aus `case`, einem oder mehreren case-Ausdrücken und einem Doppelpunkt. Der Anweisungsteil sind die Anweisungen, die auf einen case-Abschnitt folgen. Der default-case beginnt mit `default:`.
- Go wertet den switch-Ausdruck einmal aus und vergleicht seinen Wert von oben nach unten mit den case-Ausdrücken. Beim ersten Treffer führt Go den Anweisungsteil aus, und die switch-Anweisung ist zu Ende. Die Abschnitte darunter prüft Go nicht mehr.
- In Go gibt es keinen Fallthrough: Nach einem Anweisungsteil läuft das Programm nicht in den nächsten case-Abschnitt weiter. Ein `break` braucht die switch-Anweisung in Go nicht.
- Ein Programm mit Zeilennummern steht in einem dunklen Kasten, die Nummern stehen am linken Rand.
- Schritt 1 bis 4 folgen dem Kurstext. Schritt 5 enthält Faustregeln des Mentors, die nicht im Kurstext stehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Oben steht, was eine Mehrfachverzweigung ist: eine Auswahl unter mehr als zwei Fällen.
- Der Codeblock zeigt eine kurze switch-Anweisung über `num` mit `case 1:`, `case 2:` und `default:`.
- Eine Liste nennt zwei frühere Stellen des Kurstexts: Abschnitt 2.1 mit `IF b THEN … ELSE …` und Abschnitt 5.5.1 mit der Mehrfachverzweigung durch `else if`.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Go-Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Imperative Programmierung“ und „if-Anweisung“ folgen. Beide sind eigene HTML-Seiten.
- Der Codeblock lässt sich nicht ändern und nicht ausführen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 4.

### Abschnitte des Kurstexts

- 5.5.2
- 5.5.1
- 2.1

## Schritt 1: Aufbau und Ablauf

Reiter: aufbau

### Aufbau der Darstellung

- Oben stehen die fünf Begriffe: switch-Ausdruck, case-Abschnitt, case-Ausdruck, Anweisungsteil, default-case.
- Darunter steht, wie Go eine switch-Anweisung ausführt: auswerten, von oben nach unten vergleichen, den Anweisungsteil des ersten Treffers ausführen, sonst den default-case.
- Die Demo zeigt Listing 5.3 des Kurstexts, `switch_wochentag.go`, mit 29 Zeilen.
- Zeile 6: `var num int`. Zeile 7: `var text string`. Zeile 9: `fmt.Scan(&num)` liest eine Zahl. Zeile 11: `switch num {`.
- Die case-Abschnitte und ihre Anweisungsteile: Zeile 12 `case 1:` mit Zeile 13 `text = "Montag"`. Zeile 14 `case 2:` mit Zeile 15 „Dienstag“. Zeile 16 `case 3:` mit Zeile 17 „Mittwoch“. Zeile 18 `case 4:` mit Zeile 19 „Donnerstag“. Zeile 20 `case 5:` mit Zeile 21 „Freitag“. Zeile 22 `case 6, 7:` mit Zeile 23 „Wochenende“. Zeile 24 `default:` mit Zeile 25 „unbekannt“.
- Zeile 26 ist die schließende Klammer der switch-Anweisung. Zeile 28 gibt `num` und `text` aus, zum Beispiel „5 ist Freitag zugeordnet.“
- Der Student spielt den Benutzer. Der Lauf hält in Zeile 9 an und wartet. Dann wählt der Student eine der neun Zahlen 1, 2, 3, 4, 5, 6, 7, 9 und 0.
- Nach der Eingabe hält der Lauf bei Zeile 11 und dann bei jedem case-Abschnitt, den Go prüft, von oben nach unten. Passt ein case-Abschnitt, hält der Lauf danach bei seinem Anweisungsteil und springt von dort zu Zeile 28.
- Beispiel mit der Eingabe 3: Der Lauf hält bei den Zeilen 6, 7, 9, 11, 12, 14, 16, 17 und 28. Die case-Abschnitte in den Zeilen 18, 20 und 22 und der default-case prüft Go nicht mehr.
- Beispiel mit der Eingabe 9 oder 0: Go prüft alle sechs case-Abschnitte, keiner passt. Der Lauf hält dann bei Zeile 24 und Zeile 25, `text` bekommt den Wert „unbekannt“.
- `case 6, 7:` hat zwei case-Ausdrücke. Der Abschnitt passt, wenn `num` den Wert 6 oder den Wert 7 hat.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wurde oder an der das Programm wartet. Ein grüner Strich links bedeutet: Die Zeile ist in diesem Lauf schon gelaufen.
- Hinter jedem geprüften case-Abschnitt steht ein kleines Schild: rot mit „passt nicht“ oder grün mit „passt“. Hinter `default:` steht grün „kein case passt“, wenn der default-case zum Tragen kommt.
- Über dem Programm stehen drei Angaben: der Wert von `num`, der Wert von `text` und „Geprüfte case-Abschnitte bisher“. Darunter erklärt ein grüner Kasten die markierte Zeile. Unter dem Programm steht ein schwarzer Kasten „Ausgabe bisher“.
- Unter der Demo steht der Teil „Was Go bei den case-Ausdrücken verlangt“: gleicher Typ wie der switch-Ausdruck, kein Wert doppelt, beliebiger Typ, beliebige Ausdrücke, der Platz des default-case.
- Die zwei Meldungen des Compilers in diesem Teil: `cannot convert "sechs" (untyped string constant) to type int` und `duplicate case 6 (constant of type int) in expression switch`.
- `&num` gibt `fmt.Scan` Zugriff auf die Variable. Das Zeichen `&` wird im Kurstext erst in Abschnitt 7.4.1 erklärt.

### Handgriffe

- „Nächste Anweisung“ führt die nächste Anweisung aus. Wartet das Programm in Zeile 9, ist der Knopf gesperrt.
- Wartet das Programm in Zeile 9: eine Eingabe wählen, nur mit den Knöpfen 1, 2, 3, 4, 5, 6, 7, 9 und 0. Ein Eingabefeld gibt es nicht, andere Zahlen gehen nicht. Die neun Knöpfe sind nur wählbar, während das Programm in Zeile 9 wartet.
- „Bis zum Ende“ führt den Rest aus. Wartet das Programm noch auf die Eingabe, bleibt der Lauf in Zeile 9 stehen.
- „Von vorn“ löscht die Eingabe und beginnt wieder bei Zeile 6. Einen Schritt zurück gibt es nicht.
- Ein Lauf nimmt genau eine Eingabe an. Für eine andere Zahl: „Von vorn“ und neu wählen.
- Die Programmzeilen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 5.5.2

## Schritt 2: switch oder if

Reiter: vergleich

### Aufbau der Darstellung

- Der Schritt hat zwei Auswahlen mit Knöpfen. Sie hängen nicht zusammen. Der gewählte Knopf ist grün.
- Obere Auswahl: die Zuordnung der Wochentage „mit switch“ oder „mit if“. Darunter steht der Code in einem dunklen Kasten ohne Zeilennummern. Darunter steht die Angabe „So oft steht num im Code“.
- „mit switch“: `num` steht einmal im Code, hinter `switch`.
- „mit if“: `num` steht siebenmal im Code: `num == 1` bis `num == 5` und `num == 6 || num == 7`. Der Operator `||` bedeutet „oder“.
- Beide Programme tun das Gleiche. Im Programm mit `if` steht der Vergleich `num ==` in jedem Zweig neu da. Im Kurstext steht dazu, das Programm mit `if` sei etwas länger und weniger übersichtlich.
- Der Teil „Ein Funktionsaufruf als switch-Ausdruck“ zeigt eine Funktion `f`. Sie erhöht bei jedem Aufruf die globale Variable `aufrufe` um 1 und liefert den Wert 1.
- Untere Auswahl: drei Programmstücke. Darunter stehen der Code und zwei Angaben: „Wert von text danach“ und „Wert von aufrufe danach“.
- „if mit f()“: `if f() == -1 {` und `} else if f() == 1 {`. `text` bekommt „plus“, `aufrufe` hat den Wert 2. `f` wird zweimal aufgerufen, weil `f()` in zwei Bedingungen steht.
- „if mit Variable“: `fWert := f()` vor der if-Anweisung. `text` bekommt „plus“, `aufrufe` hat den Wert 1.
- „switch mit f()“: `switch f() {` mit `case -1:` und `case 1:`. `text` bekommt „plus“, `aufrufe` hat den Wert 1. Go wertet den switch-Ausdruck nur einmal aus.
- Der Wert 2 bei `aufrufe` steht in roter Schrift.
- Im Kurstext ist das Expertenwissen 5.11. Dort steht `f()` ohne Rumpf. Die Funktion mit dem Zähler ist ein eigenes Beispiel der Seite.
- Ein Seiteneffekt ist eine Wirkung einer Funktion über ihren Rückgabewert hinaus, hier die Änderung der Variable `aufrufe`.

### Handgriffe

- Obere Auswahl: „mit switch“ oder „mit if“ wählen.
- Untere Auswahl: eines der drei Programmstücke wählen.
- Kein Code dieses Schritts lässt sich ändern oder ausführen. Einen Lauf Anweisung für Anweisung gibt es hier nicht. Eine Eingabe für `num` gibt es hier nicht.
- Dem Link zur Seite „Variablendeklaration“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.2
- 5.5.1

## Schritt 3: Ohne switch-Ausdruck

Reiter: tagless

### Aufbau der Darstellung

- Oben steht: Der switch-Ausdruck darf fehlen. Go setzt dann den switch-Ausdruck `true` ein. Die Form heißt auf Englisch tagless switch.
- Jeder case-Ausdruck hat dann den Typ `bool`. Go führt den Anweisungsteil des ersten case-Ausdrucks aus, der `true` ergibt.
- Die Demo zeigt ein Programm mit acht Zeilen und Zeilennummern. Zeile 1: `switch {`. Zeile 2: `case x < 0:`. Zeile 3 gibt „x ist negativ“ aus. Zeile 4: `case x > 0:`. Zeile 5 gibt „x ist positiv“ aus. Zeile 6: `default:`. Zeile 7 gibt „x ist Null“ aus. Zeile 8: `}`.
- Der Student wählt den Wert von `x` mit einem von drei Knöpfen: −5, 0 oder 7. Die Demo zeigt das Ergebnis sofort, einen Lauf Anweisung für Anweisung gibt es hier nicht.
- Hinter den zwei case-Ausdrücken steht je ein Schild: grün mit `true`, rot mit `false` oder grau mit „nicht geprüft“. Gelb markiert ist die Anweisung, die Go ausführt.
- Mit −5: `x < 0` ergibt `true`, Zeile 3 läuft, `x > 0` ist „nicht geprüft“. Ausgabe: „x ist negativ“.
- Mit 7: `x < 0` ergibt `false`, `x > 0` ergibt `true`, Zeile 5 läuft. Ausgabe: „x ist positiv“.
- Mit 0: beide ergeben `false`, der default-case in Zeile 7 läuft. Ausgabe: „x ist Null“.
- Unter der Demo steht die gleiche Verzweigung mit `if`, `else if` und `else`.
- Der letzte Teil sagt: Mit switch-Ausdruck gehört in den case-Abschnitt ein Wert und keine Bedingung. Zu `switch x {` mit `case x < 0:` meldet der Compiler `cannot convert x < 0 (untyped bool value) to type int`.

### Handgriffe

- Den Wert von `x` wählen: nur mit den Knöpfen −5, 0 und 7. Ein Eingabefeld gibt es nicht, andere Werte gehen nicht.
- Die Programmzeilen lassen sich nicht ändern. Der Codeblock mit `if` ist fest und hat keine Knöpfe.

### Abschnitte des Kurstexts

- 5.5.2
- 5.5.1

## Schritt 4: In anderen Sprachen

Reiter: sprachen

### Aufbau der Darstellung

- Oben steht: In Java, C++ und anderen Sprachen aus der Familie von C ist die switch-Anweisung nach einem Anweisungsteil nicht zu Ende. Das Programm führt den Anweisungsteil des nächsten case-Abschnitts mit aus. Das heißt Fallthrough. Wer das nicht will, schreibt `break` ans Ende jedes Anweisungsteils.
- Die Demo hat zwei Reihen von Knöpfen. Der gewählte Knopf ist grün.
- Obere Reihe, das Programm: „Go“, „Java ohne break“, „Java mit break“.
- Untere Reihe, der Wert von `x`: 1, 2 oder 3.
- Darunter steht das gewählte Programm in einem dunklen Kasten ohne Zeilennummern. Jedes Programm hat `case 1:` mit der Ausgabe „Eins“, `case 2:` mit der Ausgabe „Zwei“ und `default:` mit der Ausgabe „Anderes“. Darunter steht die Angabe „Ausgabe“, darunter erklärt ein grüner Kasten das Ergebnis.
- „Go“ und „Java mit break“ geben genau ein Wort aus: „Eins“ für 1, „Zwei“ für 2, „Anderes“ für 3.
- „Java ohne break“ gibt für 1 drei Wörter aus: „Eins“, „Zwei“, „Anderes“. Für 2 gibt es „Zwei“ und „Anderes“ aus. Für 3 gibt es nur „Anderes“ aus. Eine Ausgabe mit mehr als einem Wort steht in roter Schrift.
- Nur das Programm „Go“ ist Go. Die zwei anderen Programme sind Java und nicht ausgeführt. Ihre Ausgaben folgen aus der Regel für den Fallthrough.
- Der Teil „Der Typ des switch-Ausdrucks“ sagt: In Go darf der switch-Ausdruck jeden Typ haben. In manchen anderen Sprachen ist er auf bestimmte primitive Typen beschränkt.
- Im Kurstext ist das der Sprachvergleich 5.12. Dort steht kein Code.

### Handgriffe

- Das Programm wählen: nur „Go“, „Java ohne break“ und „Java mit break“. Andere Sprachen gibt es nicht.
- Den Wert von `x` wählen: nur 1, 2 und 3.
- Kein Code dieses Schritts lässt sich ändern oder ausführen.

### Abschnitte des Kurstexts

- 5.5.2

## Schritt 5: Aus der Praxis

Reiter: praxis

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und Codeblöcken. Eine Demo gibt es hier nicht.
- Oben sagt ein gelber Kasten: Die Faustregeln stammen vom Mentor und stehen nicht in Abschnitt 5.5.2 des Kurstexts. Go übersetzt jedes Beispiel, auch die drei mit dem Kommentar „falsch“.
- Stoff des Kurses ist, was im Kurstext steht: die switch-Anweisung aus Abschnitt 5.5.2, wie sie Schritt 1 bis 4 zeigen. Die drei Faustregeln dieses Schritts sind Ratschläge des Mentors und gehören nicht zu diesem Stoff.
- Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt: Es steht weder auf dieser Seite noch im Kurstext. Ob eine Faustregel in der Klausur vorkommt oder nicht, lässt sich deshalb nicht sagen. Verbindlich sind der Kurstext und Moodle der FernUniversität.
- Was in der Klausur vorkommen darf und was nicht, steht ebenfalls nicht auf dieser Seite. Auskunft dazu gibt nur Moodle der FernUniversität.
- Zu jeder Faustregel stehen der falsche und der richtige Code schon auf der Seite. Ausprobieren oder umstellen lässt sich in diesem Schritt nichts.
- Faustregel 1, ein leerer case-Abschnitt tut nichts: `case 6:` ohne Anweisungsteil und darunter `case 7:` mit `text = "Wochenende"`. Bei `num` gleich 6 passt der erste Abschnitt, sein Anweisungsteil ist leer, `text` bleibt der leere String. Richtig ist `case 6, 7:`.
- Faustregel 2, der default-case fängt den Wert ab, an den niemand gedacht hat: eine switch-Anweisung ohne default-case mit `case 1:` und `case 6, 7:`. Bei `num` gleich 9 tut sie nichts, und das Programm gibt „9 ist  zugeordnet.“ mit leerem Text aus.
- Faustregel 3, ohne switch-Ausdruck entscheidet die Reihenfolge: Mit `case t < 25:` an erster Stelle gibt das Programm für −5 „angenehm“ aus. Mit `case t < 0:` an erster Stelle gibt es „frostig“ aus. Die engste Bedingung steht zuerst.
- Die Codeblöcke sind fest und haben keine Knöpfe.
- Am Ende des Schritts unter „Wie es weitergeht“ steht ein Link zum Quiz.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen.
- Den Links zum Quiz und zur Seite „if-Anweisung“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.2

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Aufbau und Ablauf, 5.5.2, Seiten 133–134: Listing 5.3: Listing 5.3 zu den Wochentagen stammt aus dem Kurstext. Die zwei Meldungen des Compilers und das Beispiel mit wort gibt die Lektion dazu.
  - 2 switch oder if, 5.5.2, Seiten 134–136: Expertenwissen 5.11: Der Vergleich von switch und if und das Expertenwissen 5.11 stammen aus dem Kurstext. Die Funktion f mit dem Zähler ist ein eigenes Beispiel.
  - 3 Ohne switch-Ausdruck, 5.5.2, Seite 135: Die switch-Anweisung ohne switch-Ausdruck stammt aus dem Kurstext. Das Programm mit if und die Meldung des Compilers gibt die Lektion dazu.
  - 4 In anderen Sprachen, 5.5.2, Seite 137: Sprachvergleich 5.12: Der Sprachvergleich mit dem Fallthrough stammt aus dem Kurstext. Die drei Programme sind eigene Beispiele.
  - 5 Aus der Praxis, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Faustregeln stammen aus der Praxis des Mentors.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.1, Seiten 5–8, Sprünge und Schleifen: Ohne Kontrollanweisungen laufen die Anweisungen der Reihe nach. Bedingte Sprünge und Schleifen ändern diese Reihenfolge. Lektion: Imperative Programmierung, Schritt 3.
  - Abschnitt 5.5.1, Seiten 129–132, if-Anweisung: Lektion: if-Anweisung in Go.
  - Abschnitt 4.5.1, Seiten 93–96, Kurzschluss-Auswertung: Warum f() in einer Kette aus else if mehrmals aufgerufen wird. Lektion: Wahrheitswerte und Kurzschluss-Auswertung.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 4.5.1
- 5.5.1
- 5.5.2
