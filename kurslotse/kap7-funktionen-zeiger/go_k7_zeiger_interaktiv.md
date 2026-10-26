# Zeiger in Go

Die Seite erklärt Zeiger an den Programmen des Kurstexts: die Adresse als Argument, den Zeiger-Typ, den Adressoperator `&`, den Dereferenzierungsoperator `*`, Zeiger als Parameter und als Rückgabe, die Lebenszeit einer Variable und die Funktion `new`. Sie gehört zu Kapitel 7 des Kurstexts, Abschnitt 7.3 und 7.4, Seiten 171–184.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Zeiger ist die Adresse einer Variable. Der Wert eines Zeigers ist die Adresse, unter der eine Variable im Speicher abgelegt ist.
- Zu jedem Typ `T` gibt es den Zeiger-Typ `*T`. Der Null-Wert jedes Zeiger-Typs ist `nil`.
- `&x` liefert die Adresse der Variable `x`. `&` heißt Adressoperator.
- `*p` ist die Variable, auf die der Zeiger `p` zeigt. `*` heißt Dereferenzierungsoperator. Der Zugriff über einen Zeiger heißt indirekt.
- Jede Demo zeigt ein Programm, darunter drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, darunter Rahmen mit Variablen, die Ausgabe und einen Kasten mit einem Satz zur letzten Anweisung.
- Jede laufende Funktion hat einen grauen Rahmen. Ihre Variablen sind Kästen mit Name und Wert. Ein gelber Kasten ist neu oder hat gerade einen Wert bekommen. Ein gestrichelter Rahmen gehört zu einer Funktion, die beendet ist.
- In den Rahmen der Demos steht der Wert eines Zeigers in Worten, zum Beispiel „Adresse von x“. Zahlen wie `0xc00000a0a8` zeigt die Seite nicht, weil die Adressen bei jedem Programmstart andere sind. Pfeile zeichnet die Seite nicht. Nur Schritt 2 zeigt Zellen mit erfundenen Adressen ab 100.
- Ein grüner Rahmen enthält Variablen ohne gültigen Bezeichner. Das Programm erreicht sie nur über einen Zeiger.
- Unter den Rahmen einer Demo steht der „Hauptspeicher“ des Modells: je Variable ein Block mit ihren Zellen, ihrem Typ und ihrem Inhalt. Ein Block mit violettem Rand ist ein Zeiger. In ihm steht eine Zahl, die Adresse einer Zelle, und darunter „Adresse der Zelle …“. Ein Block mit gelbem Schein ist neu oder gerade geändert. Die Adressen sind erfunden. Ein `int` belegt im Modell 8 Zellen, ein Zeiger auch. Endet eine Funktion, verschwinden die Blöcke ihrer Variablen.
- Die Blöcke der Demos: In Schritt 1 liegt `nums` in den Zellen 100 bis 179, das sind 80 Zellen für zehn `int`. `pa` liegt in den Zellen 180 bis 187 und enthält 100. In Schritt 3 liegt `x` in den Zellen 100 bis 107, `p` in den Zellen 108 bis 115 mit dem Inhalt 100. In Schritt 4 liegen `q` und `r` ab 100 und 108, die zwei Variablen `v` in den Zellen 200 bis 207 und 208 bis 215. In Schritt 5 liegt `x` ab 100, `p` ab 108, `q` ab 116, die Variable aus `neueIntVar` in den Zellen 200 bis 207. Am Ende steht bei ihr „freigegeben“. In Schritt 6 liegen `p`, `q` und `r` ab 100, 108 und 116, die drei anonymen Variablen ab 200, 208 und 216.
- Schritt 2 hat keine Blöcke unter den Rahmen. Dort steht die Speichersicht als 17 einzelne Zellen mit Bitmuster unter „Ausdrücke mit x, p und pp“.
- In Go lässt sich mit einem Zeiger nicht rechnen. Ein Programm kann einen Zeiger zuweisen, vergleichen und ausgeben.
- Wie ein Zeiger Zelle für Zelle im Hauptspeicher liegt, zeigt die Seite „Speichermodell: Zeiger“.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Meldungen sind mit Go 1.26 geprüft.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text knüpft an Listing 7.1 an: Die Funktion `negToZeroValue` bekam eine Kopie des Arrays und konnte das Array ihres Aufrufers nicht ändern. Mit einem Zeiger übergibt der Aufrufer die Adresse des Arrays.
- Unter „Was aus früheren Kapiteln dazugehört“ stehen drei Punkte mit Links: Abschnitt 2.2 mit der Lektion „Variablen, Typen und Speicher“, Abschnitt 5.1 mit der Lektion „Variablendeklaration in Go“, Abschnitt 7.2 mit der Lektion „Funktionen in Go“.
- Unter „Die Schritte“ steht je ein Satz zu den sechs Schritten.
- Ein gelber Kasten verweist auf die Seite „Speichermodell: Zeiger“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4
- 7.2

## Schritt 1: Adresse übergeben

Reiter: adresse

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste erklärt drei Schreibweisen von Listing 7.2: `&nums` im Aufruf, `pa *[10]int` in der Parameterliste, `*pa` im Rumpf. Für ein Element braucht der Ausdruck Klammern: `(*pa)[i]`.
- Die Demo zeigt Listing 7.2 ohne die Zeilen mit `package` und `import`, mit 14 Zeilen: die Funktion `negToZeroPointer(pa *[10]int)` in den Zeilen 1 bis 7 und `main` in den Zeilen 9 bis 14.
- Der Ablauf hat zehn Schritte: `nums` anlegen mit `[8 -5 6 -3 -7 1 -4 3 -9 2]`. `nums` ausgeben. Aufruf `negToZeroPointer(&nums)`: Der Parameter `pa` bekommt die Adresse von `nums`, kein Element wird kopiert. Dann setzt die Schleife nacheinander die Elemente 1, 3, 4, 6 und 8 auf 0, und zwar im Array `nums` im Rahmen von `main`. Dann endet die Funktion ohne Rückgabe. Dann gibt `main` `nums` aus: `[8 0 6 0 0 1 0 3 0 2]`.
- Ein Array steht in seinem Kasten als Reihe kleiner Felder. Ein grünes Feld ist das Element, das die letzte Anweisung geändert hat.
- Die Demo überspringt die Durchläufe der Schleife, in denen das Element nicht negativ ist.
- Unter „Kürzer geschrieben“ stehen die zwei Kurzschreibweisen aus Hinweis 7.4: `pa[i]` an der Stelle von `(*pa)[i]` und `for i := range pa` an der Stelle von `for i := range *pa`. Go dereferenziert den Zeiger dabei von selbst.
- Ein gelber Kasten sagt: An `pa[i]` allein sieht man nicht, ob `pa` ein Array oder ein Zeiger auf ein Array ist. Das steht in der Deklaration.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Array lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.3
- 7.4

## Schritt 2: Die Operatoren & und *

Reiter: operatoren

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste erklärt den Zeiger-Typ `*T`, den Adressoperator `&` und den Dereferenzierungsoperator `*`. Der Operand von `&` muss eine Variable sein, auch ein Element `a[i]` zählt dazu. Ein Literal wie `4.95` hat keine Adresse.
- Die Demo zeigt neun Zeilen im Rumpf von `main`: `var x int = 42`, `var p *int`, `p = &x`, `fmt.Println(*p)`, `*p = 33`, `fmt.Println(x)`, `var pp **int = &p`, `fmt.Println(**pp)`, `fmt.Println(*pp == p, p == &x)`.
- Der Ablauf: `x` ist 42. `p` ist `nil`. Nach `p = &x` zeigt `p` auf `x`. `fmt.Println(*p)` gibt 42 aus. `*p = 33` setzt `x` auf 33. `fmt.Println(x)` gibt 33 aus. `pp` zeigt auf `p` und hat den Typ `**int`. `fmt.Println(**pp)` gibt 33 aus. Die letzte Zeile gibt `true true` aus.
- Unter „Ausdrücke mit x, p und pp“ steht die Speichersicht: 17 Zellen mit Adresse und Bitmuster. Zelle 100 gehört zu `x` und enthält 00100001, das ist 33. Die Zellen 101 bis 108 gehören zu `p` und enthalten die Zahl 100, die Adresse von `x`. Die Zellen 109 bis 116 gehören zu `pp` und enthalten die Zahl 101, die Adresse von `p`. Die Adresse steht in der letzten Zelle des Zeigers, die sieben Zellen davor enthalten Nullen.
- Ein Zeiger ist im Speicher eine Variable wie jede andere: Zellen mit einem Bitmuster, das als Adresse zu lesen ist. `&` fragt nach der Adresse einer Variable, `*` geht zu der Adresse, die in einem Zeiger steht.
- Wählt der Student einen Ausdruck, sind die Zellen gelb markiert, die der Rechner dafür liest: bei `x` Zelle 100, bei `p` die Zellen 101 bis 108, bei `*p` die Zellen von `p` und Zelle 100, bei `pp` die Zellen 109 bis 116, bei `*pp` die Zellen von `pp` und von `p`, bei `**pp` alle 17. Bei `&x` und `&p` ist keine Zelle markiert, weil der Rechner die Adresse einer Variable aus seiner Tabelle kennt.
- Die Adressen 100 bis 116 sind erfunden. Im Modell hat `x` den Typ `int8`, damit es in eine Zelle passt, im Programm hat `x` den Typ `int`. Die Zellen lassen sich in diesem Schritt nicht ändern.
- Unter „Ausdrücke mit x, p und pp“ stehen 16 Knöpfe. Sie gelten für den Stand am Ende des Programms: `x` ist 33, `p` zeigt auf `x`, `pp` zeigt auf `p`. Ein Kasten nennt Typ und Wert des gewählten Ausdrucks oder die Meldung von Go.
- Die elf Ausdrücke mit Typ und Wert: `x` ist `int` 33. `&x` ist `*int`, Adresse von x. `*&x` ist `int` 33. `p` ist `*int`, Adresse von x. `*p` ist `int` 33. `&*p` ist `*int`, Adresse von x. `&p` ist `**int`, Adresse von p. `pp` ist `**int`, Adresse von p. `*pp` ist `*int`, Adresse von x. `**pp` ist `int` 33. `p == &x` ist `bool` true.
- Die fünf Fehler: `*x` meldet `invalid operation: cannot indirect x (variable of type int)`. `p = x` meldet `cannot use x (variable of type int) as *int value in assignment`. `&42` meldet `invalid operation: cannot take address of 42 (untyped int constant)`. `p + 1` meldet `invalid operation: p + 1 (mismatched types *int and untyped int)`. `var n *int; fmt.Println(*n)` übersetzt Go, das Programm bricht aber ab mit `panic: runtime error: invalid memory address or nil pointer dereference`.
- Unter „Zeiger vergleichen“ steht: Zwei Zeiger sind gleich, wenn sie auf dieselbe Variable zeigen oder beide `nil` sind. Für `var y, z int = 5, 5` gibt `fmt.Println(y == z, &y == &z)` `true false` aus.
- Ein gelber Kasten gibt Hinweis 7.3 wieder: Der Stern bildet in `*int` den Namen des Zeiger-Typs, ist in `*p` der Dereferenzierungsoperator und multipliziert in `a * b`.
- Im Kurstext stehen in Abschnitt 7.4 außerdem: die Ausgabe von Typ und Wert mit `fmt.Printf("%T, %[1]v\n", x)`, die Adressen der Elemente eines Arrays mit dem Verb `%p` und der Sprachvergleich 7.5 mit C und Java. Die Seite zeigt davon nur den Satz, dass Go keine Zeiger-Arithmetik hat.
- Dass ein Programm bei `*n` mit einem Zeiger `nil` abbricht, steht nicht in Abschnitt 7.4 des Kurstexts. Dort steht nur, dass `nil` der Null-Wert jedes Zeiger-Typs ist.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Einen der 16 Ausdrücke wählen.
- Die „Auflösung“ aufklappen.
- Ein eigener Ausdruck lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 7.4

## Schritt 3: Zeiger als Parameter

Reiter: parameter

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt zwölf Zeilen: die Funktion `incr(p *int) int` mit `*p++` und `return *p`, darunter `main` mit `x := 1`, `incr(&x)`, `fmt.Println(x)`, `fmt.Println(incr(&x))`, `fmt.Println(x)`.
- Der Ablauf hat elf Schritte. Erster Aufruf: `p` bekommt eine Kopie der Adresse von `x`, `*p++` setzt `x` auf 2, die Rückgabe 2 ist verworfen. Ausgabe 2. Zweiter Aufruf als Argument von `fmt.Println`: `x` wird 3, die Rückgabe 3 wird ausgegeben. Letzte Ausgabe 3. Die ganze Ausgabe: 2, 3, 3.
- Unter „Seiteneffekt“ steht: Die Funktion liefert einen Wert und ändert dazu eine Variable, die außerhalb ihres Rumpfs deklariert ist.
- Unter „Auch ein Zeiger wird kopiert“ steht: Call-by-value gilt auch für Zeiger. Der Parameter `p` bekommt eine Kopie der Adresse, und die Kopie zeigt auf dieselbe Variable.
- Unter „x++ und incr(&x)“ steht: `x++` ist eine Anweisung und kein Ausdruck, `fmt.Println(x++)` übersetzt Go nicht. `incr(&x)` ist Anweisung und Ausdruck.
- Ein Satz erinnert an `fmt.Scan(&s)`.
- Im Kurstext steht dazu noch: An der Stelle von `*p++` ginge auch `*p += 1` oder `*p = *p + 1`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Der Anfangswert von `x` lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.4.1

## Schritt 4: Zeiger als Rückgabe

Reiter: rueckgabe

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Der Bezeichner einer lokalen Variable gilt nur in ihrem Block. Die Variable selbst kann länger bestehen, wenn die Funktion einen Zeiger auf sie zurückgibt. Jeder Aufruf von `giveMeFive` deklariert eine neue Variable `v`.
- Die Demo zeigt 13 Zeilen: die Funktion `giveMeFive() *int` mit `v := 5` und `return &v`, darunter `main` mit `var q, r *int = giveMeFive(), giveMeFive()`, `fmt.Println(*q, *r)`, `fmt.Println(*q == *r)`, `fmt.Println(q == r)`, `*q = *q + *r`, `fmt.Println(*q, *r)`.
- Der grüne Rahmen heißt „Variablen, deren Bezeichner nicht mehr gilt“. Darin stehen die „erste Variable v“ und die „zweite Variable v“. Diese zwei Namen gehören zur Darstellung der Seite.
- Der Ablauf hat zwölf Schritte: erster Aufruf, `v` wird 5, `return &v`. Zweiter Aufruf mit einer neuen Variable `v`. Dann zeigt `q` auf die erste und `r` auf die zweite Variable. Ausgabe `5 5`. `*q == *r` gibt `true` aus. `q == r` gibt `false` aus. `*q = *q + *r` setzt die erste Variable auf 10. Die letzte Ausgabe ist `10 5`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 7.4.2
- 3.4

## Schritt 5: Lebenszeit

Reiter: lebenszeit

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text unterscheidet Lebenszeit und Gültigkeitsbereich. Die Lebenszeit ist der Zeitraum, in dem es die Variable während des Programmlaufs gibt. Der Gültigkeitsbereich ist der Teil des Programmtexts, in dem ihr Bezeichner gilt.
- Eine Liste nennt drei Fälle: ohne Zeiger endet die Lebenszeit mit dem Gültigkeitsbereich. Mit Zeigern lebt die Variable, solange mindestens ein Verweis besteht. Eine Variable ohne gültigen Bezeichner und ohne Zeiger heißt unerreichbar, der Garbage-Collector gibt ihren Speicherplatz frei.
- Die Demo zeigt 13 Zeilen: die Funktion `neueIntVar() *int` mit `var v int` und `return &v`, darunter `main` mit `var x int = 8`, `var p, q *int`, `p = neueIntVar()`, `q = p`, `p = &x`, `q = nil`. Das Programm gibt nichts aus, deshalb hat diese Demo kein Feld für die Ausgabe.
- Der Kasten unter der Demo zählt die Verweise auf die Variable aus `neueIntVar`: nach `p = neueIntVar()` ist es 1, der Zeiger `p`. Nach `q = p` sind es 2. Nach `p = &x` ist es 1, der Zeiger `q`. Nach `q = nil` sind es 0: Die Variable ist unerreichbar. Ihr Kasten ist dann grau und gestrichelt.
- Im Kurstext gibt das Programm zwischen den Zuweisungen viermal eine Adresse aus. Die Seite lässt diese vier Zeilen weg.
- Unter „Eine Variable aus einem Block“ steht das Programm mit `for i := range 10`, `i += i` und `p = &i`. Es gibt 18 aus: Im letzten Durchlauf hat `i` erst den Wert 9 und nach `i += i` den Wert 18.
- Unter „Sorgsam mit der Lebenszeit umgehen“ stehen zwei Punkte aus Abschnitt 7.4.4: Eine globale Variable lebt die ganze Laufzeit des Programms. Ein vergessener Zeiger hält Speicher belegt, das heißt Speicherleck, englisch memory leak.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 7.4.3
- 7.4.4

## Schritt 6: new

Reiter: new

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: `new(T)` erzeugt eine Variable vom Typ `T`, gibt ihr den Null-Wert von `T` und liefert ihre Adresse vom Typ `*T`. Eine solche Variable heißt unbenannte oder anonyme Variable.
- Die Demo zeigt sechs Zeilen im Rumpf von `main`: `var p *int = new(int)`, `fmt.Println(*p)`, `*p = 15`, `fmt.Println(*p)`, `var q, r *int = new(int), new(int)`, `fmt.Println(q == r)`.
- Der grüne Rahmen heißt „Anonyme Variablen“. Darin stehen die Variablen „erste“, „zweite“ und „dritte“. Diese Namen gehören zur Darstellung der Seite.
- Der Ablauf: Die erste anonyme Variable hat den Wert 0, Ausgabe 0. `*p = 15` setzt sie auf 15, Ausgabe 15. `q` und `r` zeigen auf zwei weitere anonyme Variablen. `q == r` gibt `false` aus.
- Unter „Zwei Schreibweisen, die gleiche Variable“ stehen die Funktionen `neueIntVar` mit `var v int` und `return &v` und `neueIntVarNew` mit `return new(int)`. Für die Variable macht es keinen Unterschied, wie sie erzeugt wurde.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen sechs Links: „Speichermodell: Zeiger“, der Ausblick auf Kapitel 9 aus Abschnitt 7.4.6, „Hierarchische Verbünde“ in Kapitel 8, „Rekursion in Go“, die zwei Programmierübungen mit Zeigern und das Quiz.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 7.4.5
- 7.4.6

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Adresse übergeben, 7.3, Seiten 171–172: Listing 7.2; Hinweis 7.4, Seite 177: Listing 7.2 und Hinweis 7.4 stammen aus dem Kurstext. Die Lektion lässt die Zeilen mit package und import und die Kommentare weg.
  - 2 & und *, 7.4, Seiten 173–178: Das Programm fasst Zeilen des Kurstexts zusammen. Im Kurstext gibt es auch Adressen aus, zum Beispiel 0xc00000a0a8. Die Adressen sind bei jedem Start andere, die Lektion schreibt dafür „Adresse von x“. Die Meldungen von Go sind eine Ergänzung, geprüft mit Go 1.26.
  - 3 Zeiger als Parameter, 7.4.1, Seiten 178–180: Das Programm stammt aus dem Kurstext. Die Demo mit Rahmen und Hauptspeicher gibt die Lektion dazu.
  - 4 Zeiger als Rückgabe, 7.4.2, Seite 180: Das Programm stammt aus dem Kurstext. Die Namen „erste Variable v“ und „zweite Variable v“ gehören zur Darstellung der Lektion.
  - 5 Lebenszeit, 7.4.3, Seiten 181–182; 7.4.4, Seiten 182–183: Das Programm stammt aus dem Kurstext, ohne die vier Zeilen, die Adressen ausgeben. Die Zahl der Verweise zählt die Lektion mit.
  - 6 new, 7.4.5, Seiten 183–184; 7.4.6, Seite 184: Das Programm fasst Zeilen des Kurstexts zusammen. Die Demo gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speicherplatz: Eine Variable belegt Speicherplatz. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 4.3.4, Seite 63, Zellen und Adressen im Modell: Lektion: Speichermodell.
  - Abschnitt 8.2.4, Seite 205, Zeiger auf Verbünde: Lektion: Zeiger auf Verbünde in Go.
  - Abschnitt 9.2, Seiten 229–236, verkettete Listen: Jedes Element zeigt mit einem Zeiger auf das nächste. Lektion: Verkettete Listen in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 6“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 4.3.4
- 7.3
- 7.4
- 7.4.1
- 7.4.2
- 7.4.3
- 7.4.4
- 7.4.5
- 7.4.6
- 8.2.4
- 9.2
