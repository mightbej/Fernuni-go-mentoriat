# Zeiger auf Verbünde in Go

Die Seite zeigt Zeiger auf Verbünde: den Feld-Zugriff über einen Zeiger in langer und kurzer Schreibweise, die Adresse eines Verbund-Literals und ein Literal als Argument einer Funktion. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.2.4, Seiten 205–207.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Zu jedem Typ `T` gibt es den Zeiger-Typ `*T`, auch zu einem Verbund-Typ. Eine Variable vom Typ `*person` nimmt die Adresse einer Variable vom Typ `person` auf.
- `(*pPer).nachname` ist die lange Schreibweise für den Zugriff auf ein Feld über einen Zeiger. `pPer.nachname` ist die kurze, Go dereferenziert den Zeiger von selbst. Das steht im Kurstext als Hinweis 8.2. Das gleiche Muster ist `pa[i]` für `(*pa)[i]` aus Hinweis 7.4.
- Verbund-Literale und Array-Literale heißen zusammengesetzte Literale und sind adressierbar. `&rgb{128, 0, 80}` ist die Kurzschreibweise für `new(rgb)` und eine Zuweisung des Literals.
- Die Typen der Seite: `person` mit `vorname`, `nachname`, `gebJahr`, `ort`, `verheiratet` und `rgb` mit `red, green, blue int`.
- Jede Demo zeigt ein Programm, drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, Rahmen mit Variablen, die Ausgabe und einen Kasten mit einem Satz.
- Ein Verbund steht in seinem Kasten als Liste seiner Felder. Der Wert eines Zeigers steht in Worten da, zum Beispiel „Adresse von anna“. Anonyme Variablen stehen in einem eigenen grünen Rahmen „Anonyme Variablen“ und heißen dort „erste“ und „zweite“. Diese Namen gehören zur Darstellung der Seite.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und den Definitionen der Typen `person` und `rgb`. Eine Demo gibt es hier nicht.
- Eine Liste nennt die zwei Kurzschreibweisen von Abschnitt 8.2.4.
- Unter „Was dazugehört“ stehen drei Punkte mit Links: die Lektion „Zeiger in Go“, Hinweis 7.4 und die Lektion „Verbund-Variablen und Feld-Variablen“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.2.4
- 7.4

## Schritt 1: Feld-Zugriff über einen Zeiger

Reiter: zeiger

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Eine Liste stellt die lange und die kurze Schreibweise nebeneinander.
- Die Demo zeigt fünf Zeilen im Rumpf von `main`: `var anna = person{"Anna", "Schmidt", 1966, "Hamburg", false}`, `var pPer *person = &anna`, `(*pPer).nachname = "Meier"`, `pPer.nachname = "Maier"`, `fmt.Println(pPer.vorname, anna.nachname)`.
- Der Ablauf: `pPer` zeigt auf `anna`. Die lange Schreibweise setzt den Nachnamen auf „Meier“, die kurze auf „Maier“. Die Ausgabe ist `Anna Maier`.
- Im Kurstext weisen beide Schreibweisen „Meier“ zu. Die Demo nimmt für die zweite „Maier“, damit man beide Zuweisungen sieht.
- Unter „Vier Zeilen zum Ausprobieren“ stehen vier Knöpfe für den Stand am Ende der Demo:
  1. `fmt.Println(*pPer)` gibt `{Anna Maier 1966 Hamburg false}` aus.
  2. `fmt.Println(pPer)` gibt `&{Anna Maier 1966 Hamburg false}` aus, also `&` und den Verbund, keine Adresse.
  3. `fmt.Println(*pPer.nachname)` meldet `invalid operation: cannot indirect pPer.nachname (variable of type string)`. Ohne Klammern gehört der Stern zum Feld.
  4. `var q *person; fmt.Println(q.nachname)` übersetzt Go, das Programm bricht ab mit `panic: runtime error: invalid memory address or nil pointer dereference`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Eine der vier Zeilen wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.4

## Schritt 2: Der Zeiger im Speicher

Reiter: speicher

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Im Hauptspeicher ist ein Zeiger auf einen Verbund das Gleiche wie ein Zeiger auf eine Ganzzahl, eine Variable mit 8 Zellen, in denen eine Adresse steht. Der Typ des Zeigers sagt, was an dieser Adresse liegt.
- Im Modell haben die Felder von `rgb` den Typ `uint8`, damit ein Feld in eine Zelle passt.
- Das Modell zeigt vier Variablen in 20 Zellen mit den Adressen 100 bis 119: `n` vom Typ `int8` mit dem Wert 42 in Zelle 100. `farbe` vom Typ `rgb` mit den Werten 128, 0 und 80 in den Zellen 101 bis 103. `pn` vom Typ `*int8` in den Zellen 104 bis 111, darin die Adresse 100. `pFarbe` vom Typ `*rgb` in den Zellen 112 bis 119, darin die Adresse 101.
- In jeder Zelle stehen die Adresse, das Bitmuster und der Name der Variable oder des Felds. Die Adresse in einem Zeiger steht in seiner letzten Zelle, die sieben Zellen davor enthalten Nullen.
- Der Bauplan von `rgb` im Modell: `red` hat den Abstand 0, `green` den Abstand 1, `blue` den Abstand 2.
- Acht Knöpfe zeigen je einen Ausdruck. Die Zellen, die der Rechner dafür liest, sind gelb markiert. Darunter stehen der Wert und ein Satz.
- Die acht Ausdrücke: `pn` ist 100. `*pn` ist 42, gelesen aus Zelle 100. `pFarbe` ist 101. `*pFarbe` ist `{128 0 80}`, gelesen aus den Zellen 101 bis 103. `pFarbe.green` ist 0, die Adresse ist 101 + 1 = 102. `(*pFarbe).blue` ist 80, die Adresse ist 101 + 2 = 103. `farbe.blue` ist 80 aus derselben Zelle 103. `&farbe.green` ist 102, dafür liest der Rechner keine Zelle.
- Die Zellen lassen sich in diesem Schritt nicht ändern.
- Zellen, Adressen, Bauplan und der Typ `uint8` gehören zum Modell der Seite. Die Speichersicht steht nicht im Kurstext. Ausführlich zeigen sie die Seiten „Speichermodell: Zeiger“ und „Speichermodell: Verbund“.

### Handgriffe

- Einen der acht Ausdrücke wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.4
- 7.4

## Schritt 3: Die Adresse eines Verbund-Literals

Reiter: literal

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock stellt den kurzen Code `var pFarbe *rgb = &rgb{128, 0, 80}` neben den ausgeschriebenen mit `new(rgb)` und `*pFarbe = rgb{128, 0, 80}`.
- Die Demo zeigt fünf Zeilen: `var pFarbe *rgb = &rgb{128, 0, 80}`, `fmt.Println(*pFarbe, pFarbe.green)`, `var pZwei *rgb = new(rgb)`, `*pZwei = rgb{128, 0, 80}`, `fmt.Println(*pZwei == *pFarbe, pZwei == pFarbe)`.
- Der Ablauf: Die erste Zeile erzeugt die erste anonyme Variable mit dem Wert `{128 0 80}`. Die Ausgabe ist `{128 0 80} 0`. `new(rgb)` erzeugt die zweite anonyme Variable mit dem Null-Wert `{0 0 0}`, danach bekommt sie den Wert des Literals. Die letzte Ausgabe ist `true false`: gleicher Wert, zwei verschiedene Variablen.
- Der Name `pZwei` und die zwei Zeilen mit `fmt.Println` stehen nicht im Kurstext.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.4
- 7.4.5

## Schritt 4: Ein Literal als Argument

Reiter: argument

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text gibt Expertenwissen 8.3 wieder: Eine Funktion `fügeHinzu(pFarbe *rgb)` erwartet einen Zeiger. Kurz: `fügeHinzu(&rgb{0, 255, 10})`. Ausgeschrieben braucht der Aufrufer eine Hilfsvariable: `h := new(rgb)`, `*h = rgb{0, 255, 10}`, `fügeHinzu(h)`.
- Den Rumpf von `fügeHinzu` zeigt der Kurstext nicht, die Seite auch nicht.
- Die Demo zeigt vier Zeilen: `f1 := &rgb{22, 33, 44}`, `f2 := &rgb{22, 33, 44}`, `fmt.Println(f1 == f2)`, `fmt.Println(*f1 == *f2)`.
- Der Ablauf: Jedes `&rgb{…}` erzeugt eine neue anonyme Variable. `f1 == f2` gibt `false` aus, `*f1 == *f2` gibt `true` aus.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen vier Links: die Lektionen „Vergleich von Verbünden“, „Verbünde und Funktionen“ und „Hierarchische Verbünde“ und das Quiz.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.2.4
- 8.2.5

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Feld-Zugriff über einen Zeiger, 8.2.4, Seiten 205–206: Hinweis 8.2: Die zwei Schreibweisen stammen aus dem Kurstext. Dort weisen beide den Wert "Meier" zu. Die Demo nimmt für die zweite "Maier", damit man beide Zuweisungen sieht. Die vier Zeilen zum Ausprobieren sind eine Ergänzung, geprüft mit Go 1.26.
  - 2 Im Speicher, kein Abschnitt: Aus dem Kurstext stammt nichts. Die Speichersicht steht nicht im Kurstext. Das Modell zeigen die Lektionen „Speichermodell: Zeiger“ und „Speichermodell: Verbund“ ausführlich.
  - 3 Adresse eines Literals, 8.2.4, Seite 206: Der kurze und der ausgeschriebene Code stammen aus dem Kurstext. Die zwei Zeilen mit fmt.Println und der Name pZwei sind ergänzt. Die Namen „erste“ und „zweite“ für die anonymen Variablen gehören zur Darstellung der Lektion.
  - 4 Literal als Argument, 8.2.4, Seiten 206–207: Expertenwissen 8.3: Die Funktion fügeHinzu und die Zeiger f1 und f2 stammen aus dem Kurstext. Die letzte Zeile der Demo, *f1 == *f2, ist ergänzt.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.
  - Abschnitt 7.4.5, Seiten 183–184, anonyme Variable: new erzeugt eine Variable ohne Namen, wie das Literal mit &. Lektion: Zeiger in Go, Schritt 6.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4
- 7.4.5
- 8.2.4
