# for-Anweisung mit for-Klausel in Go

Die Seite erklärt die Schleife, die zählt: `for`, eine for-Klausel aus Init-Anweisung, Schleifenbedingung und Post-Anweisung, und einen Schleifenrumpf. Sie führt das Programm zu den Quadratzahlen aus Listing 5.6 Anweisung für Anweisung aus. Sie gehört zu Abschnitt 5.5.3.2 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ausgaben und Meldungen des Compilers stammen aus Läufen mit Go 1.26 und stehen fest auf der Seite.
- Auf der Seite läuft kein Go. Der Student kann keinen eigenen Code eingeben und kein Programm starten. Was sich einstellen lässt, steht unter „Handgriffe“.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Die for-Klausel steht zwischen `for` und `{`. Ihre drei Teile sind durch zwei Semikolons getrennt: Init-Anweisung, Schleifenbedingung, Post-Anweisung. In `for i := 1; i <= 10; i++ {` ist `i := 1` die Init-Anweisung, `i <= 10` die Schleifenbedingung und `i++` die Post-Anweisung.
- Die Init-Anweisung läuft genau einmal am Anfang. Die Schleifenbedingung wird vor jedem Schleifendurchlauf geprüft, auch vor dem ersten. Die Post-Anweisung läuft nach jedem Schleifendurchlauf, also nach dem Schleifenrumpf.
- Die Schleifenbedingung wird einmal öfter geprüft, als der Schleifenrumpf ausgeführt wird.
- Die Laufvariable ist die Variable, die in der Init-Anweisung deklariert ist und in der Schleife zählt. Sie gilt nur in der for-Anweisung.
- Schritt 1 bis 4 folgen dem Kurstext. Schritt 5 enthält Faustregeln des Mentors, die nicht im Kurstext stehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und zwei Codeblöcken. Eine Demo gibt es hier nicht.
- Der erste Codeblock zeigt eine zählende Schleife als for-Anweisung mit Bedingung: `i := 1`, `for i <= 10 {`, `fmt.Println(i * i)`, `i++`, `}`.
- Der zweite Codeblock zeigt die gleiche Schleife mit for-Klausel: `for i := 1; i <= 10; i++ {`.
- Beide Programme geben die Quadratzahlen 1, 4, 9 bis 100 aus.
- Eine Liste nennt drei frühere Abschnitte des Kurstexts: 5.5.3.1 mit der for-Anweisung mit Bedingung, 5.4 mit den einfachen Anweisungen, 5.1.3 mit dem Gültigkeitsbereich.
- Eine zweite Liste sagt, was die fünf Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass die Programme der Seite mit Go 1.26 ausgeführt sind.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „for-Anweisung mit Bedingung“, „Zuweisungen“ und „Variablendeklaration“ folgen. Alle drei sind eigene HTML-Seiten.
- Die Codeblöcke lassen sich nicht ändern und nicht ausführen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demos stehen in Schritt 1 bis 4.

### Abschnitte des Kurstexts

- 5.5.3.2
- 5.5.3.1
- 5.4

## Schritt 1: Die for-Klausel

Reiter: klausel

### Aufbau der Darstellung

- Oben stehen die drei Teile der for-Klausel mit ihrem Beispiel: Init-Anweisung `i := 1`, Schleifenbedingung `i <= 10`, Post-Anweisung `i++`.
- Die Demo heißt „Ausprobieren: wo die drei Teile stehen“. Drei Knöpfe wählen einen Teil. Der gewählte Knopf ist grün.
- Unter den Knöpfen stehen zwei Programme nebeneinander in dunklen Kästen ohne Zeilennummern. Links „Mit Bedingung“: `i := 1`, `for i <= 10 {`, die Ausgabe mit `fmt.Printf`, `i++`, `}`. Rechts „Mit for-Klausel“: `for i := 1; i <= 10; i++ {`, die Ausgabe mit `fmt.Printf`, `}`.
- Der gewählte Teil ist in beiden Programmen gelb markiert. So sieht man, wohin jedes Stück des linken Programms im rechten gewandert ist.
- Die Init-Anweisung steht links als eigene Zeile vor der Schleife und rechts vor dem ersten Semikolon. Die Schleifenbedingung steht links hinter `for` und rechts zwischen den Semikolons. Die Post-Anweisung steht links als letzte Zeile im Schleifenrumpf und rechts hinter dem zweiten Semikolon.
- Darunter erklärt ein grüner Kasten den gewählten Teil: wann er läuft und was an seiner Stelle erlaubt ist.
- Die Demo führt die Programme nicht aus. Den Lauf zeigt Schritt 2.
- Der Teil „Was Go bei der for-Klausel verlangt“ nennt drei Regeln mit je einer Meldung des Compilers: `syntax error: var declaration not allowed in for initializer` für `var` in der Init-Anweisung, `syntax error: cannot declare in post statement of for loop` für `:=` in der Post-Anweisung, `syntax error: unexpected :=, expected )` für runde Klammern um die for-Klausel.
- Der Teil „Was %2[1]v bedeutet“ erklärt die Formatangabe: `%v` setzt einen Wert ein, `2` gibt ihm mindestens zwei Zeichen Platz, `[1]` wählt das erste Argument, also `i`. `%3v` gibt danach `i*i` mit mindestens drei Zeichen Platz aus.
- Eine Zeile der Ausgabe sieht so aus: ` 3 *  3 ist   9`.

### Handgriffe

- Einen der drei Teile wählen: nur mit den Knöpfen „Init-Anweisung“, „Schleifenbedingung“ und „Post-Anweisung“.
- Die zwei Programme lassen sich nicht ändern und nicht ausführen. Werte lassen sich in diesem Schritt nicht einstellen.

### Abschnitte des Kurstexts

- 5.5.3.2
- 5.4

## Schritt 2: Der Ablauf

Reiter: ablauf

### Aufbau der Darstellung

- Oben steht, wann die drei Teile laufen: die Init-Anweisung einmal am Anfang, die Schleifenbedingung vor jedem Schleifendurchlauf, die Post-Anweisung nach jedem Schleifendurchlauf.
- Die Demo zeigt Listing 5.6 des Kurstexts, `for_klausel_quadrat.go`, mit neun Zeilen und Zeilennummern.
- Zeile 5: `func main() {`. Zeile 6: `for i := 1; i <= 10; i++ {`. Zeile 7: die Ausgabe mit `fmt.Printf`. Zeile 8: die schließende Klammer der Schleife. Zeile 9: die schließende Klammer von `main`.
- Zeile 6 gibt es in zwei Fassungen. „wie im Kurstext“: `i <= 10`. „geändert“: `i <= 3`. Die Fassung mit 3 steht nicht im Kurstext, die Seite hat sie gebaut, damit der Lauf kürzer ist. Beim ersten Öffnen der Seite ist die Fassung mit 3 gewählt.
- Der Student geht den Lauf Anweisung für Anweisung durch. Der Lauf hält bei Zeile 5, dann immer wieder bei Zeile 6 und Zeile 7 und am Ende bei Zeile 9.
- Zeile 6 kommt in drei Rollen dran. Hinter ihr steht dann ein kleines Schild: grau „Init-Anweisung“, grün „Schleifenbedingung: true“ oder rot „Schleifenbedingung: false“, grau „Post-Anweisung“.
- Die Reihenfolge eines Laufs: Init-Anweisung. Dann je Schleifendurchlauf: Schleifenbedingung, Zeile 7, Post-Anweisung. Am Ende einmal die Schleifenbedingung mit `false`.
- Mit `i <= 3` hat der Lauf 13 Anweisungen: Der Schleifenrumpf läuft dreimal, die Schleifenbedingung wird viermal geprüft. Die Ausgabe hat drei Zeilen für 1, 2 und 3. Nach dem dritten Schleifendurchlauf hat `i` den Wert 4.
- Mit `i <= 10` hat der Lauf 34 Anweisungen: Der Schleifenrumpf läuft zehnmal, die Schleifenbedingung wird elfmal geprüft. Nach dem zehnten Schleifendurchlauf hat `i` den Wert 11.
- Die gelb markierte Programmzeile ist die Zeile, die gerade ausgeführt wurde. Ein grüner Strich links bedeutet: Die Zeile ist in diesem Lauf schon gelaufen.
- Über dem Programm stehen drei Angaben: der Wert von `i`, „Prüfungen der Bedingung bisher“ und „Schleifendurchläufe bisher“. Vor der Init-Anweisung steht bei `i` „gibt es noch nicht“, nach dem Ende der Schleife „gibt es nicht mehr“.
- Unter den Angaben erklärt ein grüner Kasten die markierte Zeile. Unter dem Programm steht ein schwarzer Kasten „Ausgabe bisher“.
- Unter der Demo steht der Teil „Das Ende genau betrachtet“: Nach dem letzten Schleifendurchlauf läuft die Post-Anweisung noch einmal, erst die Prüfung danach beendet die Schleife.

### Handgriffe

- Die Fassung von Zeile 6 wählen: „i <= 10“ oder „i <= 3“. Das beginnt den Lauf von vorn. Andere Grenzen gibt es nicht.
- „Nächste Anweisung“ führt die nächste Anweisung aus. „Bis zum Ende“ führt den Rest aus. „Von vorn“ beginnt wieder bei Zeile 5. Einen Schritt zurück gibt es nicht.
- Der Startwert 1 und die Post-Anweisung `i++` stehen fest. Die Programmzeilen lassen sich nicht ändern. Das Programm hat keine Eingabe.

### Abschnitte des Kurstexts

- 5.5.3.2

## Schritt 3: Die Laufvariable

Reiter: laufvariable

### Aufbau der Darstellung

- Oben steht: Die Init-Anweisung ist fast immer eine kurze Variablendeklaration. Die Variable, die sie einführt, heißt Laufvariable oder Kontrollvariable.
- Der Gültigkeitsbereich einer Laufvariable aus der Init-Anweisung umfasst die for-Klausel und den Schleifenrumpf. Hinter der Schleife gibt es sie nicht mehr.
- Die Demo hat drei Knöpfe für drei Fassungen. Der gewählte Knopf ist grün. Darunter steht das Programm der Fassung mit Zeilennummern, darunter „Ergebnis“ oder „Meldung des Compilers“ und „Ausgabe“, darunter ein grüner Kasten mit der Erklärung.
- Gelb markiert ist hier die Zeile, auf die es in der Fassung ankommt. Einen Lauf Anweisung für Anweisung gibt es in diesem Schritt nicht.
- Fassung „i im Schleifenrumpf“, drei Zeilen: `for i := 1; i <= 3; i++ {`, `fmt.Println(i)`, `}`. Go übersetzt. Ausgabe: 1, 2, 3. Markiert ist Zeile 2.
- Fassung „i hinter der Schleife“, vier Zeilen: wie die erste, dazu Zeile 4 `fmt.Println(i)` hinter der Schleife. Meldung des Compilers: `undefined: i`. Markiert ist Zeile 4.
- Fassung „i vor der Schleife deklariert“, fünf Zeilen: `i := 1`, `for ; i <= 3; i++ {`, `fmt.Println(i)`, `}`, `fmt.Println("danach:", i)`. Go übersetzt. Ausgabe: 1, 2, 3 und „danach: 4“. Markiert ist Zeile 5.
- In der dritten Fassung ist die Init-Anweisung leer, das erste Semikolon bleibt stehen. `i` ist dann keine Variable der Schleife und gilt auch dahinter. Der Wert 4 ist der Wert, mit dem die Prüfung `4 <= 3` die Schleife beendet hat.
- Der Teil „Wie die Laufvariable heißt“ sagt: üblich sind `i`, `j` und `k`. Der Grundsatz: Je kleiner der Gültigkeitsbereich einer Variable ist, desto kürzer darf ihr Bezeichner sein.
- Der Teil „Zwei Laufvariablen“ zeigt `for i, j := 0, 100; i < j; i, j = i+1, j-5 {`. Die Schleife gibt 17 Zeilen aus, die erste ist `i, j: 0 100`, die letzte `i, j: 16 20`. Im Kurstext ist das Expertenwissen 5.15.

### Handgriffe

- Eine der drei Fassungen wählen: nur mit den drei Knöpfen.
- Die Programmzeilen lassen sich nicht ändern. Der Codeblock mit den zwei Laufvariablen ist fest und hat keine Knöpfe.

### Abschnitte des Kurstexts

- 5.5.3.2
- 5.1.3

## Schritt 4: Eine Schleife, vier Schreibweisen

Reiter: varianten

### Aufbau der Darstellung

- Oben steht: Alle drei Teile der for-Klausel dürfen fehlen. Go hat im Grunde nur eine Schleife.
- Die Demo hat vier Knöpfe für vier Schreibweisen. Der gewählte Knopf ist grün. Darunter steht das Programm in einem dunklen Kasten ohne Zeilennummern, darunter erklärt ein grüner Kasten die Schreibweise.
- Alle vier Programme geben 1, 2 und 3 aus. Go übersetzt alle vier.
- „mit allen drei Teilen“: `for i := 1; i <= 3; i++ {`.
- „ohne Init- und Post-Anweisung“: `i := 1` vor der Schleife, `for ; i <= 3; {`, im Schleifenrumpf `fmt.Println(i)` und `i++`. Die zwei Semikolons bleiben stehen.
- „ohne die Semikolons“: `for i <= 3 {`. Das ist die for-Anweisung mit Bedingung aus Abschnitt 5.5.3.1.
- „ohne Schleifenbedingung“: `for {`. Go setzt `true` als Schleifenbedingung ein. Im Schleifenrumpf steht `if i > 3 { break }`, sonst wäre es eine Endlosschleife.
- Darunter steht der Teil „Wozu eine Schleife ohne Bedingung gut ist“: für Programme, die auf Dauer laufen, englisch event loop. Sie endet mit `break` oder mit `os.Exit(0)`.
- Im Kurstext ist das Expertenwissen 5.16. Die vier Programme mit 1 bis 3 sind eigene Beispiele der Seite.
- Die Anweisung `break` wird im Kurstext erst in Abschnitt 5.5.3.3 erklärt.

### Handgriffe

- Eine der vier Schreibweisen wählen: nur mit den vier Knöpfen.
- Kein Code dieses Schritts lässt sich ändern oder ausführen. Einen Lauf Anweisung für Anweisung gibt es hier nicht.
- Dem Link zur Seite „break-Anweisung“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.3.2
- 5.5.3.1
- 5.5.3.3

## Schritt 5: Aus der Praxis

Reiter: praxis

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und Codeblöcken. Eine Demo gibt es hier nicht.
- Oben sagt ein gelber Kasten: Die Faustregeln stammen vom Mentor und stehen nicht in Abschnitt 5.5.3.2 des Kurstexts. Go übersetzt jede Fassung, auch die vier mit dem Kommentar „falsch“. `{ … }` steht für einen beliebigen Schleifenrumpf.
- Stoff des Kurses ist, was im Kurstext steht: die for-Anweisung mit for-Klausel aus Abschnitt 5.5.3.2, wie sie Schritt 1 bis 4 zeigen. Die drei Faustregeln dieses Schritts sind Ratschläge des Mentors und gehören nicht zu diesem Stoff.
- Welche Fragen das Quiz stellt, steht auf dieser Seite nicht.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt: Es steht weder auf dieser Seite noch im Kurstext. Ob eine Faustregel in der Klausur vorkommt oder nicht, lässt sich deshalb nicht sagen. Verbindlich sind der Kurstext und Moodle der FernUniversität.
- Zu jeder Faustregel stehen die falsche und die richtige Fassung schon auf der Seite. Ausprobieren oder umstellen lässt sich in diesem Schritt nichts.
- Faustregel 1, Startwert und Vergleich gehören zusammen: `for i := 0; i < 10; i++` und `for i := 1; i <= 10; i++` laufen beide zehnmal. `for i := 1; i < 10; i++` läuft nur neunmal, `for i := 0; i <= 10; i++` läuft elfmal. Der Fehler heißt off-by-one error. In Go ist die Form ab 0 häufiger, weil das erste Element eines Arrays den Index 0 hat.
- Faustregel 2, die Laufvariable ändert nur die Post-Anweisung: Eine Schleife `for i := 1; i <= 10; i++` mit einem zusätzlichen `i++` im Schleifenrumpf gibt 1, 3, 5, 7 und 9 aus. Wer in Schritten von 2 zählen will, schreibt `i += 2` in die Post-Anweisung.
- Faustregel 3, wer herunterzählt, dreht alle drei Teile um: `for i := 10; i >= 1; i--` läuft zehnmal. `for i := 10; i <= 1; i--` läuft kein einziges Mal, weil `10 <= 1` schon bei der ersten Prüfung `false` ergibt.
- Die Codeblöcke sind fest und haben keine Knöpfe.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen.
- Den Links zum Quiz und zur Seite „for-Anweisung mit Bedingung“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.3.2

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Die for-Klausel, 5.5.3.2, Seiten 143–144: Init-Anweisung, Schleifenbedingung und Post-Anweisung stammen aus dem Kurstext. Die drei Meldungen des Compilers und die Erklärung der Formatangabe gibt die Lektion dazu.
  - 2 Der Ablauf, 5.5.3.2, Seiten 143–144: Listing 5.6: Listing 5.6 zu den Quadratzahlen stammt aus dem Kurstext. Die Fassung mit i <= 3 gibt die Lektion dazu.
  - 3 Die Laufvariable, 5.5.3.2, Seiten 143–145: Expertenwissen 5.15: Dass die Laufvariable nur in der for-Anweisung gilt und einen kurzen Namen wie i bekommt, steht im Kurstext. Das Expertenwissen 5.15 zeigt Init- und Post-Anweisung mit mehreren Variablen. Die drei Programme der Demo und die Meldung des Compilers gibt die Lektion dazu.
  - 4 Eine Schleife, vier Schreibweisen, 5.5.3.2, Seiten 145–146: Expertenwissen 5.16: Dass Teile der for-Klausel fehlen dürfen, steht im Expertenwissen 5.16. Die vier Programme sind eigene Beispiele.
  - 5 Aus der Praxis, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Faustregeln stammen aus der Praxis des Mentors.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.1, Seiten 5–8, Sprünge und Schleifen: Ohne Kontrollanweisungen laufen die Anweisungen der Reihe nach. Bedingte Sprünge und Schleifen ändern diese Reihenfolge. Lektion: Imperative Programmierung, Schritt 3.
  - Abschnitt 5.4, Seite 129, einfache Anweisungen: Was als Init-Anweisung und Post-Anweisung stehen darf. Lektion: Zuweisungen in Go, Schritt 5.
  - Abschnitt 5.5.3.1, Seiten 138–142, for mit Bedingung: Lektion: for-Anweisung mit Bedingung in Go.
  - Abschnitt 6.2, Seiten 156–159, Arrays durchlaufen: Die for-Klausel zählt den Index eines Arrays hoch. Lektion: Arrays in Go.
- Unter „Wie es weitergeht“ steht ein Link zum Quiz.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 5.4
- 5.5.3.1
- 5.5.3.2
- 6.2
