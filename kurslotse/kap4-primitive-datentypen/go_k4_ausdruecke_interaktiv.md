# Operatoren und Ausdrücke in Go

Die Seite erklärt Abschnitt 4.1 bis 4.1.2 des Kurstexts, Seiten 45–52. Sie beginnt bei den Daten: Ein Rechner verarbeitet Daten im Rechenwerk, jedes Datum hat einen Typ, und zu jedem Typ gehören genau festgelegte Operationen. Danach zeigt sie, welche Operatoren es für welche Typen gibt, woraus ein Ausdruck besteht, in welcher Reihenfolge Go ihn ausrechnet und was ein konstanter Ausdruck ist.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt. Schritt 1 bis 5 erklären den Abschnitt, in Schritt 6 probiert der Student Programme im Editorlotsen aus, Schritt 7 ordnet alles dem Kurstext zu.

Für alle Schritte gilt:

- Ein Operator verbindet seine Operanden zu einem Ausdruck. Die Auswertung rechnet den Ausdruck aus und liefert seinen Wert.
- Ein unärer Operator hat einen Operanden und steht davor, zum Beispiel `-a`. Ein binärer Operator hat zwei Operanden und steht zwischen ihnen, zum Beispiel `a + b`.
- Die Begriffe Ausdruck, Auswertung und Bindungsstärke wurden in Abschnitt 3.3.5 und 3.3.6 schon einmal kurz eingeführt, am Programm zur Umrechnung von Temperaturen, Listing 3.2. Die Lektion dazu heißt „Zuweisen, rechnen, ausgeben“ und liegt im Editorlotsen unter Kapitel 3, nicht auf GitHub Pages.
- Die Division zweier Ganzzahlen und der Rest `%` gehören zu Abschnitt 4.3.7.1. Die Kästen „Mehr dazu später“ in Schritt 1 und 4 verlinken die Lektion „Rechnen mit Ganzzahlen: Division und Rest in Go“; sie ersetzt dort den bisherigen Verweis auf Typkonvertierung. Die Seite nennt sie nur, wo eine Demo sie enthält, zum Beispiel `7 / 2` mit dem Wert 3 oder `3/4` mit dem Wert 0.
- Alle Ausgaben und Meldungen auf der Seite hat Go 1.26.1 erzeugt. Auf der Seite selbst läuft kein Go, außer in Schritt 6 im Editorlotsen.
- Schritt 1 bis 5 enden mit einem violetten Kasten „Mehr dazu später“. Er nennt Abschnitte, in denen ein Thema genauer behandelt wird, und die Lektionen dazu.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text beginnt mit den Daten: Ein Programm arbeitet mit Daten, zum Beispiel einer Temperatur, einem Namen oder der Antwort ja oder nein. Der Rechner legt sie im Speicher ab und verarbeitet sie im Rechenwerk.
- Jedes Datum hat einen Typ. Der Typ legt den Wertebereich und die Operationen fest.
- Eine Tabelle mit den Spalten Typ, Beispielwert und „Was man mit Werten dieses Typs tun kann“ nennt vier Typen: `int` mit 7, `float64` mit 2.5, `string` mit "Hallo", `bool` mit true.
- Unter „Operator, Operand, Ausdruck“ stehen die Begriffe am Beispiel `5 + 8` mit dem Wert 13.
- Unter „Was du schon kennst“ stehen Links auf den Exkurs „Speichern und Rechnen im Rechner“ und die Lektion „Rechenwerk“ und der Hinweis auf die Lektion im Editorlotsen zu 3.3.5 und 3.3.6.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.1
- 2.3
- 2.2

## Schritt 1: Operatoren

Reiter: operatoren

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit der Frage, ob man zwei Strings mit `-` voneinander abziehen darf. Auflösung: nein, für Strings ist von den Rechenzeichen nur `+` festgelegt.
- Eine Liste erklärt binären und unären Operator.
- Die Demo „Typ und Operator“: In einer Reihe stehen vier Knöpfe für den Typ von `a` und `b`: `int`, `float64`, `string`, `bool`. In einer zweiten Reihe stehen elf Knöpfe für den Operator: `+ - * / % == < && ||`, `! (unär)` und `- (unär)` für das Vorzeichen.
- Darunter steht das Programm ab Zeile 6 im Rumpf von `main`: `var a TYP = …`, `var b TYP = …`, `fmt.Println(a OP b)`. Bei einem unären Operator fehlt `b`, und Zeile 7 ist `fmt.Println(!a)` oder `fmt.Println(-a)`. Die Werte sind fest: `int` 7 und 2, `float64` 7.5 und 2.0, `string` "Hallo" und "Welt", `bool` true und false.
- Darunter stehen die Ausgabe oder die Meldung des Compilers und ein Kasten mit Sätzen der Seite: das Ergebnis und sein Typ, oder dass der Operator für den Typ nicht festgelegt ist. Bei einer Meldung ist der Kasten rot. Am Ende des Kastens steht, welche der elf Operatoren für den gewählten Typ festgelegt sind und welche nicht.
- Ein Vergleich `==` oder `<` liefert immer `bool`. Bei den übrigen Operatoren hat das Ergebnis den Typ der Operanden.
- Besonderheiten, die der Kasten nennt: `7 / 2` ergibt bei `int` 3, die Nachkommastellen fallen weg. `7 % 2` ergibt 1. Bei `float64` gibt `7.5 * 2.0` den Wert 15 aus, `fmt.Println` lässt `.0` weg. `%` gibt es nicht für `float64`. `"Hallo" + "Welt"` ergibt `HalloWelt` ohne Leerzeichen. `"Hallo" < "Welt"` ist true, Reihenfolge im Wörterbuch. `true < false` übersetzt Go nicht, Wahrheitswerte haben keine Reihenfolge.
- Darunter lässt sich „Alle Operatoren für die primitiven Typen, Tabelle 4.1 des Kurstexts“ aufklappen: die Tabelle mit Art, Operatoren, was sie tun und für welche Typen sie festgelegt sind, auch die Bit-Operatoren und die komplexen Zahlen. Darunter die drei Gruppen: arithmetische Operatoren, logische Operatoren, Vergleichsoperatoren. Die komplexen Zahlen braucht der Kurs nicht.
- Der Kasten „Mehr dazu später“ nennt: Division und Rest bei Ganzzahlen, 4.3.7.1. Bit-Operatoren, 4.3.7.2. Vergleiche, 4.5.2. Verkettung von Strings, 4.6.2.2. Die unären Operatoren `&` und `*`, 7.4. `<-` braucht der Kurs nicht.

### Handgriffe

- Einen der vier Typen und einen der elf Operatoren wählen. Andere Werte oder andere Typen lassen sich nicht eingeben.
- Die Tabelle 4.1 auf- und zuklappen.

### Abschnitte des Kurstexts

- 4.1

## Schritt 2: Typen passen zusammen

Reiter: typen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit der Frage, ob Go `x + t` übersetzt, wenn `x` ein `int` und `t` ein `float64` ist. Auflösung: nein.
- Eine Liste nennt zwei Regeln. Gleicher Typ: Beide Operanden eines binären Operators haben den gleichen Typ. Ausnahme: Steht eine ganze Zahl als Integer-Literal im Programm und ist der andere Operand eine Gleitkommazahl, passt Go den Typ des Literals an, die implizite Typanpassung. Ergebnistyp: Rechenzeichen und logische Operatoren liefern den Typ der Operanden, ein Vergleich liefert immer `bool`.
- Vier Knöpfe wählen ein Programm, jedes ab Zeile 6 im Rumpf von `main`. Darunter stehen Ausgabe oder Meldung und ein Kasten mit einem Satz der Seite.
- Die vier Programme: `x + t` meldet `mismatched types int and float64`. `32 + tempC*1.8` mit `tempC = 22` vom Typ `float64` gibt 71.6, das Beispiel aus Listing 3.2. `alter >= 18` mit `alter = 20` gibt true, Typ `bool`. `"Hallo" + " Welt"` gibt `Hallo Welt`.
- Ein gelber Kasten „Operator oder Funktion?“ sagt: Ein Operator ist eine kurze Schreibweise für eine Funktion mit einem oder zwei Parametern, `add(add(3, 4), 5)` statt `3 + 4 + 5`. In Go lässt sich die Bedeutung eines Operators nicht ändern. In C++ heißt das Überladen eines Operators.
- Der Kasten „Mehr dazu später“ nennt die Typkonvertierung, Abschnitt 4.7.

### Handgriffe

- Eines der vier festen Programme wählen. Ein eigenes Programm lässt sich hier nicht eingeben.

### Abschnitte des Kurstexts

- 4.1
- 4.7

## Schritt 3: Aufbau von Ausdrücken

Reiter: aufbau

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit der Frage: Welche der vier Zeilen `123`, `x = 5`, `tempC - 4`, `var y int` ist ein Ausdruck? Auflösung: `123` und `tempC - 4`. Die anderen zwei sind Anweisungen und haben keinen Wert.
- Eine Liste erklärt elementare Ausdrücke (Literal, Name einer Variable oder Konstante, Funktionsaufruf ohne Argument) und kombinierte Ausdrücke (durch einen Operator, als Argument eines Funktionsaufrufs, durch Klammern).
- Sieben Knöpfe wählen einen Ausdruck aus Abschnitt 4.1.1: `123`, `x * x`, `"Hallo" + " Welt"`, `32 + tempC*1.8`, `Quadrat(2+3)`, `6 * (7 + 3)`, `kontostand >= 0 && einkommen >= 1000`.
- Darunter steht der Ausdruck als Kästen in Kästen: Jeder Teilausdruck hat einen Rahmen, kleinere Kästen stehen in größeren. Grün sind elementare Ausdrücke, blau kombinierte. Der angeklickte Kasten ist gelb umrandet. Eine Legende über den Kästen nennt die drei Farben.
- Darunter erklärt ein Kasten den angeklickten Teilausdruck: ob er elementar oder kombiniert ist, was er ist und aus welchen Teilen er besteht. Dazu steht da, aus wie vielen Kästen der ganze Ausdruck besteht.
- Ein Funktionsaufruf ohne Argument gilt als elementar, mit Argumenten als kombiniert. So steht es im Text des Kurstexts. In einer Fußnote wird ergänzt, dass man großzügiger alle Funktionsaufrufe zu den elementaren Ausdrücken zählen kann.
- Ein gelber Kasten sagt: `x + 1` allein in einer Zeile übersetzt Go nicht, die Meldung lautet `x + 1 (value of type int) is not used`. Ein Funktionsaufruf ist zugleich Ausdruck und Anweisung, zum Beispiel `fmt.Println(x)`.
- Der Kasten „Mehr dazu später“ nennt Funktionen mit Rückgabewert, Abschnitt 7.1, und Anweisungen in Kapitel 5.

### Handgriffe

- Einen der sieben festen Ausdrücke wählen und einen Kasten anklicken. Einen eigenen Ausdruck kann der Student nicht eingeben.
- Die Demo rechnet in diesem Schritt nichts aus. Die Deklarationen der Variablen und Funktionen nennt die Seite nicht, sie spielen für den Aufbau keine Rolle.

### Abschnitte des Kurstexts

- 4.1.1

## Schritt 4: Auswertung

Reiter: auswertung

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit der Frage, ob `8 / 4 * 5` den Wert 10 oder 0,4 hat. Auflösung: 10, gleich starke Operatoren gelten von links nach rechts.
- Eine Liste nennt drei Regeln: Bindungsstärke (Punkt vor Strich, Rechnen vor Vergleichen, Vergleichen vor `&&` und `||`), von links nach rechts (Linksassoziativität), Klammern zuerst.
- Die Demo „der Rechenweg“: Sechs Knöpfe wählen einen Ausdruck aus Abschnitt 4.1.2: `5 + 2 * 3` (11), `(5 + 2) * 3` (21), `8 / 4 * 5` (10), `24/8 - 3*4` (-9), `10+x*5 >= 50 || x < 1` und `9 - 3/4*g(x-y) + 24/(f()+h(g(7), 5+x))` (10).
- Ein dunkler Kasten zeigt den Rechenweg wie in der Schule, eine Zeile je Rechenschritt, jede neue Zeile beginnt mit `=`. In jeder Zeile ist der gerade ausgerechnete Teil durch seinen Wert ersetzt, in Grün. In der untersten Zeile ist gelb markiert, was Go als Nächstes ausrechnet. Rechts neben jeder Zeile steht klein, welche Rechnung gemacht wurde und warum gerade dieser Teil dran war, zum Beispiel „2 * 3 = 6 · * bindet stärker als +, Punkt vor Strich“. Die früheren Zeilen bleiben stehen. Am Ende steht „Fertig: Der Ausdruck hat den Wert …“.
- Literale sind sofort da und brauchen keinen eigenen Rechenschritt. Eine Variable ist ein eigener Schritt: Go holt ihren Wert.
- Beim Ausdruck mit `||` wählt der Student x = 0, 3 oder 8. Werte: bei 0 true, bei 3 false, bei 8 true. Bei x = 8 ist der linke Vergleich `50 >= 50` true, und Go rechnet `x < 1` nicht mehr aus, die Kurzschluss-Auswertung aus Abschnitt 4.5.1. Die Demo streicht den Teil dann grau durch.
- Beim langen Ausdruck gilt x = 3 und y = 1, `f()` gibt 2 zurück, `g(a)` gibt `a * 2` zurück, `h(a, b)` gibt `a + b` zurück. Diese Rümpfe und Werte hat die Seite festgelegt, im Kurstext stehen sie nicht. `3/4` ergibt 0, weil die Division zweier Ganzzahlen die Nachkommastellen abschneidet.
- „Klammern zeigen“ blendet unter dem Rechenweg die Form mit einer Klammer um jeden Teilausdruck ein, so wie Go ihn liest, zum Beispiel `(((10 + (x * 5)) >= 50) || (x < 1))`.
- Die Demo rechnet im Browser. Die Werte hat Go 1.26.1 nachgerechnet.
- Darunter lässt sich die Tabelle der Bindungsstärken aus Abschnitt 4.1.2 aufklappen, sechs Stufen: 6 alle unären Operatoren, 5 `* / % & &^ << >>`, 4 `+ - | ^`, 3 die Vergleiche, 2 `&&`, 1 `||`.
- Ein gelber Kasten sagt: Ein Ausdruck mit Variablen oder Funktionsaufrufen hat seinen Wert erst beim Lauf, und Go rechnet ihn bei jedem Auftreten neu aus.
- Der Kasten „Mehr dazu später“ nennt: warum `3/4` den Wert 0 hat, 4.3.7.1. Die Kurzschluss-Auswertung, 4.5.1. Vergleiche verknüpfen, 4.5.2.

### Handgriffe

- Einen der sechs festen Ausdrücke wählen. Beim Ausdruck mit `||` den Wert von x wählen: 0, 3 oder 8.
- „Nächster Rechenschritt“, „Einen Schritt zurück“, „Von vorn“, „Klammern zeigen“ und „Klammern verbergen“.
- Die Tabelle der Bindungsstärken auf- und zuklappen.
- Eigene Ausdrücke und andere Werte lassen sich hier nicht eingeben. Das geht in Schritt 6.

### Abschnitte des Kurstexts

- 4.1.2
- 3.3.6

## Schritt 5: Konstante Ausdrücke

Reiter: konstant

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit der Frage, was bei `100 / (3*5 - fix)` geschieht, einmal mit der Konstante `fix`, einmal mit der Variable `fix`. Auflösung: Mit der Konstante übersetzt Go das Programm nicht, mit der Variable bricht es beim Lauf ab.
- Der Text erklärt: Ein konstanter Ausdruck besteht nur aus Literalen und Konstanten. Der Compiler rechnet ihn beim Übersetzen aus. Eine Liste nennt die zwei Folgen: schneller, und ein Fehler wie die Division durch 0 fällt schon beim Übersetzen auf.
- Fünf Knöpfe wählen ein Programm, jedes ab Zeile 6 im Rumpf von `main`. Darunter stehen Ausgabe oder Meldung und ein Kasten mit einem Satz der Seite.
- Die fünf Programme: Der Ausdruck mit `tageProJahr` gibt 13.75. `780 - 280 + 10 * 14` gibt 640. Mit `const fix int = 8 + 7` meldet der Compiler `invalid operation: division by zero`. Mit `var fix int = 8 + 7` übersetzt Go das Programm, es bricht beim Lauf ab mit `panic: runtime error: integer divide by zero` und `exit status 2`. Den Pfad zur Datei in dieser Meldung hat die Seite gekürzt. Das fünfte Programm gibt `x*x + 1` zweimal aus, mit x = 4 und x = 10, also 17 und 101.
- „panic“ ist ein Abbruch des laufenden Programms. Das Wort steht im Kurstext an dieser Stelle nicht.
- Der Kasten „Mehr dazu später“ nennt Konstanten, Abschnitt 5.2.

### Handgriffe

- Eines der fünf festen Programme wählen. Ein eigenes Programm lässt sich hier nicht eingeben.

### Abschnitte des Kurstexts

- 4.1.2
- 5.2

## Schritt 6: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Oben stehen drei Vorschläge: die Klammern in `24/8 - 3*4` so setzen, dass der Wert −6 ist; im Programm „Vergleich und ODER“ x auf 0 oder 8 setzen; im Programm „Typen mischen“ die Deklaration von `x` so ändern, dass Go übersetzt.
- Darunter steht ein Editor mit sechs Programmen zur Wahl: „Punkt vor Strich“ gibt 11, 21, 10, -9, 3 aus. „Vergleich und ODER“ gibt zweimal false aus. „Funktionen im Ausdruck“ mit `f`, `g` und `h` gibt 10 und 0 aus. „Typen mischen“ meldet `mismatched types int and float64`. „Ausdruck allein“ meldet `x + 1 (value of type int) is not used`. „Konstante Division durch 0“ meldet `division by zero`.
- Im Editorlotsen übersetzt Go das Programm auf dem Rechner des Studenten und führt es aus. Auf GitHub Pages ist der Reiter grau. Dort zeigt er nur, was Go 1.26.1 zu jedem Programm gemeldet oder ausgegeben hat, und sagt, wie der Student die Lektion im Editorlotsen öffnet.
- Die Lösung des ersten Vorschlags ist `24/(8 - 3*4)`: `8 - 12` ergibt −4, `24/-4` ergibt −6. Die Seite nennt die Lösung nicht.
- Am Ende des Schritts unter „Nächste Schritte“ stehen Abschnitt 4.2 zu `fmt.Printf` mit dem Hinweis auf die Lektion im Editorlotsen und die Quiz „Auswertungsreihenfolge und Bindungsstärken“ und „Quiz zu Kapitel 4“.

### Handgriffe

- Ein Programm wählen, im Editorlotsen ändern und ausführen, „Programm zurücksetzen“.
- Die Lektion speichert nichts.

### Abschnitte des Kurstexts

- 4.1
- 4.1.1
- 4.1.2

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle mit den Spalten Schritt, Abschnitt im Kurstext, „Was aus dem Kurstext stammt“ und „Was die Lektion dazugibt“.
- Aus dem Kurstext stammen der Satz über Wertebereich und Operationen, unär und binär, Tabelle 4.1, die Regeln zu Typ und Ergebnistyp, Expertenwissen 4.2 und 4.3, die Ausdrücke in Schritt 3 und 4, die Tabelle der Bindungsstärken, die drei Regeln, die konstanten Ausdrücke mit 13.75 und 640 und die Division durch 0 mit `const fix`.
- Die Lektion gibt dazu: den Einstieg über Daten, Speicher und Rechenwerk, die Tabelle mit vier Typen, die Demo Typ und Operator, die Kästen in Kästen, den Rechenweg mit Begründung, die Rümpfe von `f`, `g` und `h` und die Werte von x und y, das Programm mit der Variable `fix` und alle Ausgaben und Meldungen von Go 1.26.1.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen Abschnitt 2.3, 3.3.5 und 3.3.6, 4.5.1, 4.7, 5.2 und 5.3 mit ihren Lektionen.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.1
