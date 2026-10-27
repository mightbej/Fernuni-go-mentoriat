# Wahrheitswerte: bool und die Kurzschluss-Auswertung in Go

Die Seite erklärt den Typ `bool`, die drei logischen Operatoren `!`, `&&` und `||` und die Kurzschluss-Auswertung: Go wertet den rechten Operanden von `&&` und `||` nur aus, wenn das Ergebnis nach dem linken noch nicht feststeht. Sie zeigt das an den zwei Beispielen des Kurstexts. Sie gehört zu Abschnitt 4.5 und 4.5.1 des Kurstexts.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Ausgaben und die Meldung zum Abbruch sind mit Go 1.26 nachgeprüft und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Bei `a && b` wertet Go zuerst `a` aus. Ist `a` falsch, ist das Ergebnis `false`, und Go wertet `b` nicht aus. Ist `a` wahr, wertet Go `b` aus, und `b` entscheidet.
- Bei `a || b` wertet Go zuerst `a` aus. Ist `a` wahr, ist das Ergebnis `true`, und Go wertet `b` nicht aus. Ist `a` falsch, wertet Go `b` aus, und `b` entscheidet.
- Am Wert des Ausdrucks ändert die Kurzschluss-Auswertung nichts. Sie ändert, welche Teile des Programms laufen.
- Grün steht auf der Seite für `true` oder wahr, Rot für `false` oder falsch.
- Jeder Schritt beginnt mit einem blauen Kasten „Erst vorhersagen, dann ausprobieren“. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Code und Demo gibt es hier nicht.
- Oben stehen die zwei Regeln der Kurzschluss-Auswertung für `a && b` und `a || b`.
- Darunter steht, was sich dadurch am Programm ändert: Eine Division rechts findet nicht statt, ein Funktionsaufruf rechts unterbleibt.
- Eine Liste nennt zwei frühere Stellen des Kurstexts: Abschnitt 3.3.3 mit dem Typ `bool` und Abschnitt 4.1 mit Operatoren und Ausdrücken.
- Eine zweite Liste sagt, was die drei Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Ausgaben mit Go 1.26 nachgeprüft sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 3.

### Abschnitte des Kurstexts

- 4.5
- 4.5.1

## Schritt 1: Zwei Wahrheitswerte

Reiter: werte

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage: `a` ist `false`. Wertet Go bei `a && b` den Operanden `b` aus? Die Auflösung: nein.
- Darunter stehen zwei Tabellen. Die erste nennt den Typ `bool`, seine zwei Werte und woher ein Wert dieses Typs kommt. Die zweite nennt je Operator, wann der Ausdruck wahr ist und wann Go den rechten Operanden auswertet.
- `!` ist die Negation, `&&` die AND-Verknüpfung oder Konjunktion, `||` die OR-Verknüpfung oder Disjunktion.
- Die Demo zeigt zwei große Knöpfe, beschriftet mit `a` und `b`. Auf jedem steht sein Wert, `true` in Grün oder `false` in Rot. Ein Klick schaltet den Wert um.
- Unter den zwei Knöpfen stehen drei Zeilen: `!a`, `a && b` und `a || b`. Jede zeigt ihr Ergebnis als farbiges Schild und rechts einen Satz dazu, ob Go `b` auswertet.
- Die vier Belegungen: Sind beide `false`, ergibt `a && b` den Wert `false` und `a || b` den Wert `false`. Ist nur `b` wahr, ergibt `&&` `false` und `||` `true`. Ist nur `a` wahr, ergibt `&&` `false` und `||` `true`. Sind beide wahr, ergeben beide `true`.
- Ein gelber Kasten sagt: Go wandelt zwischen `bool` und Zahlen nichts um. Eine 0 ist kein `false`, eine 1 kein `true`. Auch `false < true` übersetzt Go nicht, weil `bool` kein geordneter Typ ist. Im Kurstext ist das der Sprachvergleich 4.17.

### Handgriffe

- Auf den Wert von `a` oder den Wert von `b` klicken: Er wechselt zwischen `true` und `false`.
- Andere Handgriffe hat die Demo nicht. Die drei Ausdrücke stehen fest, ein eigener Ausdruck lässt sich nicht eingeben.
- Die Auflösung der Frage aufklappen.

### Abschnitte des Kurstexts

- 4.5
- 4.5.1

## Schritt 2: Die Division, die nicht stattfindet

Reiter: division

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit drei Fragen zu `x` gleich 0 und `x` gleich 10.
- Die Demo untersucht den Ausdruck `x != 0 && (100 / x) > 10` aus dem Kurstext. `x` hat den Typ `int`, die Division ist eine Division ganzer Zahlen.
- Der Student wählt den Wert von `x` und die Reihenfolge der zwei Teile. Die gewählte Reihenfolge ist grün.
- Reihenfolge wie im Kurstext: `x != 0 && (100 / x) > 10`. Die Prüfung auf 0 steht links und schützt die Division.
- Vertauschte Reihenfolge: `(100 / x) > 10 && x != 0`. Die Division steht links und läuft immer zuerst. Diese Möglichkeit steht nicht im Kurstext.
- In einem hellen Kasten steht der Ausdruck groß mit seinen zwei Teilen. Jeder Teil ist gefärbt: grün für wahr, rot für falsch, grau und durchgestrichen für „Go wertet ihn nicht aus“, schwarz für den Abbruch.
- Darunter stehen zwei Zeilen „Schritt 1“ und „Schritt 2“ der Auswertung und ein farbiger Balken mit dem Ergebnis. Diese „Schritte“ sind die zwei Schritte der Auswertung, nicht die Schritte der Seite.
- Unter dem Kasten steht ein dunkler Codeblock mit zwei Programmzeilen und der Ausgabe als Kommentar.
- Mit `x` gleich 0 und der Reihenfolge des Kurstexts: `0 != 0` ist falsch, die Division wird übersprungen, der Ausdruck hat den Wert `false`.
- Mit `x` gleich 0 und vertauschter Reihenfolge: Go rechnet zuerst `100 / 0`. Das Programm bricht ab mit `panic: runtime error: integer divide by zero`.
- Mit `x` gleich 5: 100 / 5 ist 20, und 20 > 10 ist wahr. Der Ausdruck hat den Wert `true`.
- Mit `x` gleich 10: 100 / 10 ist 10, und 10 > 10 ist falsch. Der Ausdruck hat den Wert `false`, Go wertet beide Teile aus.
- Mit `x` gleich 20: 100 / 20 ist 5, und 5 > 10 ist falsch.
- Ein Abbruch zur Laufzeit heißt in Go panic. Der Compiler bemerkt die Division durch 0 hier nicht, weil `x` eine Variable ist.

### Handgriffe

- `x` wählen: mit den Knöpfen 0, 5, 10 und 20 oder als ganze Zahl von −1000000 bis 1000000 im Eingabefeld.
- Die Reihenfolge wählen: mit einem der zwei Knöpfe.
- Die Zahl 100 und der Vergleich mit 10 stehen fest. Der Operator ist immer `&&`. Ein eigener Ausdruck lässt sich nicht eingeben.
- Die Auflösungen der drei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.5.1

## Schritt 3: Die Funktion, die nicht aufgerufen wird

Reiter: funktion

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen zu `true || f()` und `false || f()`.
- Die Demo zeigt Listing 4.2 des Kurstexts als Programm zum Umbauen. Links steht das Programm, rechts die Ausgabe in einem schwarzen Kasten und darunter ein gelber Kasten mit der Erklärung.
- Das Programm hat zwei Funktionen. `f` gibt die Zeile `<<Funktion f wird aufgerufen>>` aus und liefert einen Wahrheitswert zurück. `main` gibt mit `fmt.Println("ergibt", … f())` das Wort „ergibt“ und den Wert des Ausdrucks aus.
- Der Student stellt dreierlei ein, je mit zwei Knöpfen: den linken Operanden `true` oder `false`, den Operator `&&` oder `||`, den Rückgabewert von `f`. Der gewählte Knopf ist grün.
- Wird `f` aufgerufen, hat die Ausgabe zwei Zeilen: erst `<<Funktion f wird aufgerufen>>`, dann `ergibt` mit dem Wert. Wird `f` nicht aufgerufen, hat die Ausgabe nur die Zeile mit `ergibt`.
- `f` wird aufgerufen bei `true && f()` und bei `false || f()`. `f` wird nicht aufgerufen bei `false && f()` und bei `true || f()`.
- Die Zeile von `f` erscheint vor der Zeile mit „ergibt“, weil Go erst den ganzen Ausdruck auswertet und dann `fmt.Println` ausführt.
- Im Kurstext hat das Listing zwei Ausgaben in einem Programm, `true && f()` und `true || f()`, und `f` liefert `false`. Die übrigen Programme hat die Seite gebaut.
- Eine Wirkung einer Funktion neben ihrem Rückgabewert heißt Seiteneffekt. Hier ist es die Ausgabe der Zeile.
- Funktionen mit Rückgabewert werden im Kurstext in Abschnitt 3.4 kurz und in Kapitel 7 ganz erklärt.
- Am Ende des Schritts unter „Nächste Schritte“ stehen Links zu zwei Quizseiten. Welche Fragen sie stellen, steht auf dieser Seite nicht.

### Handgriffe

- Den linken Operanden wählen: `true` oder `false`.
- Den Operator wählen: `&&` oder `||`.
- Den Rückgabewert von `f` wählen: `true` oder `false`.
- Die Programmzeilen lassen sich sonst nicht ändern.
- Die Auflösungen der zwei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.5.1

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Zwei Wahrheitswerte, 4.5, Seite 93: bool; 4.5.1, Seiten 93–96: Kurzschluss-Auswertung; 4.5.2, Seiten 96–98: Sprachvergleich 4.17: Der Typ bool, die Operatoren &&, || und ! und der Sprachvergleich stammen aus dem Kurstext. Die Demo mit zwei umschaltbaren Werten gibt die Lektion dazu.
  - 2 Die Division, die nicht stattfindet, 4.5.1, Seiten 94–95: Die Bedingung, die eine Division durch 0 verhindert, stammt aus dem Kurstext. Die Möglichkeit mit vertauschten Teilen und die Meldung zum Abbruch gibt die Lektion dazu.
  - 3 Die Funktion, die nicht aufgerufen wird, 4.5.1, Seiten 95–96: Listing 4.2: Listing 4.2 mit true && f() und true || f() stammt aus dem Kurstext. Die übrigen Programme sind für die Lektion gebaut.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.3, Seiten 10–11, Abkürzen: Dort steht das Beispiel a > b ∧ c/d ≥ 0 mit d = 0 zum ersten Mal, ohne Go. Schritt 2 dieser Lektion zeigt es in Go. Eine Lektion zu diesem Teil von Abschnitt 2.3 gibt es nicht.
  - Abschnitt 3.3.3, Seiten 29–30, bool: Der Typ bool mit den Werten true und false. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 5.5.1, Seiten 129–132, Bedingung: Wo ein Wahrheitswert gebraucht wird. Lektion: if-Anweisung in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links zu den zwei Quizseiten, zur Seite „Logische Operatoren und Vergleiche“ und zur Seite „for-Anweisung mit Bedingung“ folgen. Alle sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 2.3
- 3.3.3
- 4.5
- 4.5.1
- 4.5.2
- 5.5.1
