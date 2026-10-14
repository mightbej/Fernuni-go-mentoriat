# Strings in Go: Literale, Verkettung und Unveränderbarkeit

Die Seite zeigt, wie man einen String im Quelltext schreibt und was man mit Strings tun kann. Schritt 1 bis 3 behandeln die String-Literale: den Backslash als Escape-Zeichen, das rohe String-Literal und die Unicode-Nummern. Schritt 4 zeigt die Verkettung mit `+`. Schritt 5 zeigt, dass ein String unveränderbar ist. Die Seite gehört zu den Abschnitten 4.6.1 und 4.6.2.2 des Kurstexts.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Ein Wert vom Typ `string` ist eine Zeichenkette. Das String-Literal ist seine Schreibweise im Quelltext.
- Literal und Wert sind zweierlei. Mehrere Literale können für den gleichen Wert stehen. Dem Wert sieht man nicht an, aus welchem Literal er stammt.
- Das interpretierte String-Literal steht zwischen doppelten Anführungszeichen. Darin ist der Backslash das Escape-Zeichen.
- Das rohe String-Literal steht zwischen Backquotes. Darin gibt es keine Escape-Sequenzen.
- Ein einfaches Anführungszeichen `'…'` begrenzt in Go keinen String, sondern ein einzelnes Zeichen. Das kommt auf der Seite nicht vor.
- Wie ein String als Folge von Bytes im Speicher liegt und was `len(s)`, `s[i]` und `s[low:high]` liefern, zeigt eine andere HTML-Seite: „Strings: Bytes und Zeichen“. Diese Seite hier nennt die Zahl der Bytes nur als Angabe.
- Auf der Seite läuft kein Go. In Schritt 1 bis 3 rechnet der Browser nach den Regeln von Go. Die Programmzeilen und die Meldungen des Compilers sind mit Go 1.26 nachgeprüft.
- Schritt 1 und 5 beginnen mit einem blauen Kasten mit Fragen. Unter jeder Frage steht eine aufklappbare „Auflösung“. Die Antwort ist verdeckt, bis der Student sie aufklappt.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einem Codeblock. Eine Demo gibt es hier nicht.
- Der Codeblock: `var text string = "grüne Wiese"` und `fmt.Println(text)` mit der Ausgabe `grüne Wiese`.
- Ein Absatz trennt Literal und Wert: Das Literal hat die Anführungszeichen, der Wert nicht.
- Eine Liste nennt die zwei Formen: interpretiertes String-Literal mit `"…"` und rohes String-Literal mit Backquotes.
- Eine Liste nennt drei frühere Stellen des Kurstexts: Abschnitt 3.2, Abschnitt 3.3.7 und Abschnitt 4.3.1.
- Der Teil „Zwei Lektionen zu Abschnitt 4.6“ sagt, welche HTML-Seite was behandelt.
- Eine weitere Liste sagt, was die fünf Schritte zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den Seiten „Zahlensysteme in Go“ und „Strings: Bytes und Zeichen“ folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die erste Demo steht in Schritt 1.

### Abschnitte des Kurstexts

- 4.6.1
- 4.6

## Schritt 1: Der Backslash

Reiter: backslash

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit zwei Fragen.
- Eine Tabelle nennt die Escape-Sequenzen: `\n` Zeilenumbruch, `\t` Tabulator, `\\` Backslash, `\"` doppeltes Anführungszeichen, dazu die fünf seltenen Steuerzeichen `\a`, `\b`, `\f`, `\r`, `\v` und `\u` und `\U` mit einer Nummer.
- Eine Escape-Sequenz besteht aus dem Backslash und dem Zeichen danach. Im Wert steht dafür ein einziges Zeichen.
- Der Teil „Warum es \\ und \" braucht“: Ein `"` beendet das Literal, ein `\` leitet eine Escape-Sequenz ein. Wer eines der zwei Zeichen im Wert haben will, schreibt `\"` oder `\\`.
- Ein Codeblock zeigt das Beispiel des Kurstexts: `"Der Backslash \\ ist ein \"besonderes\" Zeichen."` mit der Ausgabe `Der Backslash \ ist ein "besonderes" Zeichen.`
- Die Demo hat ein Eingabefeld, das zwischen zwei festen Anführungszeichen steht, und den Knopf „lesen“. Darunter stehen sechs Knöpfe mit Beispielen.
- Ein Kasten sagt, wie viele Escape-Sequenzen das Literal enthält. Übersetzt Go das Literal nicht, ist der Kasten rot und nennt den Grund.
- Ein dunkles Feld zeigt den Wert so, wie `fmt.Println` ihn ausgibt. Ein Zeilenumbruch und ein Tabulator wirken dort wirklich. Ein anderes Steuerzeichen steht als gelbes Schild da.
- Darunter stehen Angaben: Zeichen im Literal, Zeichen im Wert, `len(s)` in Bytes und die Liste der Escape-Sequenzen.
- Die sechs Beispiele: `Der Backslash \\ ist ein \"besonderes\" Zeichen.` wird zu einem Wert mit 45 Zeichen. `Unendlich ist \u221e und Smiley ist \U0001f600` wird zu „Unendlich ist ∞ und Smiley ist 😀“. `Zeile 1\nZeile 2` wird zu zwei Zeilen. `Name\tOrt` enthält einen Tabulator. `C:\neu\tabelle.txt` ist eine Falle: Go liest `\n` als Zeilenumbruch und `\t` als Tabulator, der Pfad ist zerstört. `Er sagte "Hallo"` übersetzt Go nicht, weil das innere Anführungszeichen das Literal beendet.
- Eine Escape-Sequenz, die es nicht gibt, zum Beispiel `\q`, übersetzt Go nicht. Die Meldung lautet `unknown escape`.
- Die Escape-Sequenzen `\x` mit zwei Hexadezimalziffern und `\` mit drei Oktalziffern gibt es in Go. Im Kurstext kommen sie nicht vor, und die Seite zeigt sie nicht.

### Handgriffe

- Zwischen den zwei Anführungszeichen einen Text mit höchstens 70 Zeichen eintragen und „lesen“ wählen oder die Eingabetaste drücken. Die äußeren Anführungszeichen werden nicht mit eingetragen.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Die Auflösungen der zwei Fragen aufklappen.

### Abschnitte des Kurstexts

- 4.6.1

## Schritt 2: Rohe Literale

Reiter: roh

### Aufbau der Darstellung

- Oben steht, was ein rohes String-Literal ist: Es steht zwischen zwei Backquotes, und jedes Zeichen kommt so in den Wert, wie es dasteht.
- Ein Codeblock zeigt den Satz mit dem Backslash als rohes Literal, ohne jede Escape-Sequenz.
- Das einzige Zeichen, das in einem rohen String-Literal nicht stehen kann, ist der Backquote selbst.
- Ein gelber Kasten sagt, wie man den Backquote auf einer deutschen Tastatur eingibt: Umschalttaste und Akzenttaste, danach die Leertaste. Er ist das Akzentzeichen von è und nicht der Apostroph.
- Der Teil „Wann sich das rohe Literal lohnt“ nennt zwei Fälle: viele Backslashes und Anführungszeichen, zum Beispiel ein Pfad unter Windows, und mehrere Zeilen.
- Ein interpretiertes Literal muss in einer Zeile stehen. Ein rohes Literal darf über mehrere Zeilen gehen, und jeder Zeilenumbruch im Quelltext gehört zum Wert.
- Ein Codeblock zeigt das Programm mit `htmlCode` aus dem Kurstext, gekürzt auf fünf Zeilen HTML.
- Die Demo hat ein Eingabefeld für den Wert und den Knopf „als Literal schreiben“. Darunter stehen vier Knöpfe mit Beispielen.
- Angaben darunter: der Wert, das interpretierte String-Literal, das rohe String-Literal und `len(s)` in Bytes.
- Für den Pfad `"C:\Program Files (x86)\Windows Defender"` mit Anführungszeichen im Wert: Das interpretierte Literal lautet `"\"C:\\Program Files (x86)\\Windows Defender\""` und braucht vier Escape-Sequenzen. Das rohe Literal ist der Wert zwischen zwei Backquotes.
- Enthält der Wert einen Backquote, sagt die Seite „geht nicht“ für das rohe Literal.
- Enthält der Wert keinen Backslash und kein Anführungszeichen, unterscheiden sich die zwei Literale nur in den Begrenzern.

### Handgriffe

- Eine Zeichenkette mit höchstens 70 Zeichen in einer Zeile eintragen und „als Literal schreiben“ wählen oder die Eingabetaste drücken.
- Einen der vier Knöpfe mit Beispielen wählen.
- Einen Zeilenumbruch nimmt das Feld nicht an.

### Abschnitte des Kurstexts

- 4.6.1

## Schritt 3: Unicode-Nummern

Reiter: unicode

### Aufbau der Darstellung

- Oben steht: Unicode hat über 120.000 Zeichen, ASCII hatte 128. Jedes Unicode-Zeichen hat eine Nummer, den Unicode code point. In Go heißt die Nummer auch rune. Man schreibt sie hexadezimal.
- `\u` mit vier Hexadezimalziffern: `"\u221e"` ist ∞. `\U` mit acht Hexadezimalziffern: `"\U0001f600"` ist 😀. Eine kürzere Nummer wird vorn mit Nullen aufgefüllt.
- Ein Codeblock zeigt die Zeile des Kurstexts: `fmt.Println("Unendlich ist \u221e und Smiley ist \U0001f600")`.
- Die Demo hat ein Eingabefeld für die Nummer und den Knopf „Zeichen zeigen“. Darunter stehen sechs Knöpfe: 221e, 1f600, 20ac, fc, 41 und 3b1.
- Das Zeichen erscheint groß. Darunter stehen Angaben: die Unicode-Nummer mit `U+` und dezimal, das String-Literal mit Escape-Sequenz, das String-Literal mit dem Zeichen selbst, die Zahl der Bytes im Speicher und ob das kleine `u` genügt.
- Die Beispiele: 221e ist ∞ mit 3 Bytes. 1f600 ist 😀 mit 4 Bytes und braucht `\U0001f600`. 20ac ist € mit 3 Bytes. fc ist ü mit 2 Bytes, das Literal lautet `"\u00fc"`. 41 ist A mit 1 Byte. 3b1 ist α mit 2 Bytes.
- Die größte Unicode-Nummer ist 10FFFF. Die Nummern D800 bis DFFF gehören zu keinem Zeichen, Go übersetzt ein Literal damit nicht. Bei einem Steuerzeichen sagt die Seite, dass es kein Schriftbild hat.
- Ein gelber Kasten sagt: `\u` und `\U` gelten nur im interpretierten String-Literal.
- Wie ein Zeichen heißt, sagt die Seite nicht.

### Handgriffe

- Eine Nummer aus höchstens acht Hexadezimalziffern eintragen und „Zeichen zeigen“ wählen oder die Eingabetaste drücken.
- Einen der sechs Knöpfe mit Beispielen wählen.
- Ein Zeichen selbst lässt sich nicht eintragen, nur seine Nummer.

### Abschnitte des Kurstexts

- 4.6.1
- 4.6.2

## Schritt 4: Verketten

Reiter: verketten

### Aufbau der Darstellung

- Oben steht: Der Operator `+` heißt bei Strings Verkettungsoperator oder Konkatenationsoperator. Er hängt den rechten String an den linken und liefert einen neuen String. Go fügt nichts dazwischen ein.
- Beide Operanden müssen Strings sein. Go wandelt eine Zahl nicht von selbst in einen String um.
- Die Demo zeigt fünf Knöpfe mit je einer Zeile. Die gewählte Zeile steht in einem dunklen Codeblock.
- Ein Kasten darunter nennt die Ausgabe und den Grund. Übersetzt Go die Zeile nicht, ist der Kasten rot und nennt die Meldung des Compilers.
- `"hoch" + "springen"` ergibt `hochspringen`.
- `"3" + "4"` ergibt `34`. Go hängt an und rechnet nicht.
- `"das Quadrat von 12 ist " + (12 * 12)` übersetzt Go nicht. Die Meldung: `invalid operation: "das Quadrat von 12 ist " + (12 * 12) (mismatched types untyped string and untyped int)`.
- `fmt.Println("das Quadrat von 12 ist", 12*12)` gibt `das Quadrat von 12 ist 144` aus. Hier wird nichts verkettet: `fmt.Println` bekommt zwei Werte und setzt ein Leerzeichen dazwischen.
- `t := "Go"` und `t += "!"` ergibt `Go!`. `t += "!"` ist die Kurzform von `t = t + "!"`.
- Der Teil „Was es außer + für Strings gibt“ nennt die Vergleichsoperatoren und die Ausdrücke `s[i]` und `s[low:high]`, mit Links. Minus, Mal und Geteilt gibt es für Strings nicht.
- Wie man eine Zahl in einen String umwandelt, zeigt die Seite nicht. Im Kurstext steht es auch nicht: Abschnitt 4.7 behandelt nur die Konvertierung zwischen Zahlentypen, Strings kommen dort nicht vor. Funktionen wie `strconv.Itoa` und `fmt.Sprintf` kommen im Kurstext nicht vor. Wer eine Zahl neben einem Text ausgeben will, gibt `fmt.Println` beide Werte, durch ein Komma getrennt.

### Handgriffe

- Eine der fünf Zeilen wählen.
- Eigener Code lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 4.6.2.2
- 3.3.7

## Schritt 5: Unveränderbar

Reiter: unveraenderbar

### Aufbau der Darstellung

- Oben steht der blaue Kasten mit einer Frage.
- Strings sind in Go unveränderbar, auf Englisch immutable. Ein einzelnes Zeichen lässt sich nicht austauschen.
- Ein Codeblock zeigt `var s string = "Ohr"` und `s[0] = 85`. Go übersetzt die zweite Zeile nicht. Die Meldung: `cannot assign to s[0] (neither addressable nor a map index expression)`. Die 85 ist die Nummer des Buchstabens U.
- Der Teil „Was stattdessen geht“: Die Variable ist veränderbar, der String nicht. `s = "Uhr"` weist einen anderen String zu. `s = "U" + s[1:]` baut einen neuen String aus „U“ und dem Rest des alten. Beides ergibt „Uhr“.
- Die Demo führt das Programm des Kurstexts Zeile für Zeile aus. Es steht als nummerierte Liste da: `var a string = "Sommer"`, `var b string = a`, `a = a + "zeit"`, `fmt.Println(a)`, `fmt.Println(b)`. Die zuletzt ausgeführte Zeile ist hell unterlegt.
- Darunter stehen drei Knöpfe: „nächste Zeile ausführen“, „eine Zeile zurück“ und „von vorn“.
- Zwei Kästen zeigen die Variablen `a` und `b` mit ihrem Wert. Eine Variable, die noch nicht deklariert ist, hat einen gestrichelten Kasten. Die Variable, die die letzte Zeile angelegt oder geändert hat, ist gelb.
- Ein dunkles Feld zeigt die Ausgabe des Programms. Ein Kasten darunter sagt, was die letzte Zeile getan hat.
- Der Ablauf: Nach Zeile 1 hat `a` den Wert „Sommer“. Nach Zeile 2 hat auch `b` den Wert „Sommer“. Nach Zeile 3 hat `a` den Wert „Sommerzeit“, `b` hat weiter „Sommer“. Zeile 4 gibt `Sommerzeit` aus, Zeile 5 gibt `Sommer` aus.
- Die Verkettung ändert keinen String. Sie baut aus zwei Strings einen neuen und weist ihn `a` zu.
- Der Teil „Was die Regel bringt“ nennt das Expertenwissen 4.19: Kopien eines Strings und Substrings können sich den gleichen Platz im Arbeitsspeicher teilen.

### Handgriffe

- „nächste Zeile ausführen“ wählen: Die nächste Zeile läuft.
- „eine Zeile zurück“ wählen: Der Stand vor der letzten Zeile kommt zurück.
- „von vorn“ wählen: Keine Zeile ist ausgeführt.
- Die Auflösung der Frage aufklappen.
- Das Programm lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 4.6.2.2
- 4.6.2.1

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Der Backslash, 4.6.1, Seiten 99–100: Das erste und das zweite Beispiel der Demo stammen aus dem Kurstext. Die Meldungen des Compilers gibt die Lektion dazu.
  - 2 Rohe Literale, 4.6.1, Seiten 100–102: Der Satz mit dem Backslash, der Pfad und das Programm mit htmlCode stammen aus dem Kurstext. Die Lektion kürzt das Programm.
  - 3 Unicode-Nummern, 4.6.1, Seite 99: Die Nummern 221E und 1F600 und die Zeile mit fmt.Println stammen aus dem Kurstext. Die anderen Nummern sind eigene Beispiele.
  - 4 Verketten, 4.6.2.2, Seite 107: Die Ausdrücke "hoch" + "springen" und "das Quadrat von 12 ist " + (12 * 12) stammen aus dem Kurstext. Die drei anderen Zeilen der Demo und die Meldung des Compilers gibt die Lektion dazu.
  - 5 Unveränderbar, 4.6.2.2, Seiten 107–108: Expertenwissen 4.19: Das Beispiel mit „Ohr“ und das Programm mit „Sommer“ stammen aus dem Kurstext. Die Meldung des Compilers und die Zeile s = "U" + s[1:] gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Konkatenation: Für Zeichenketten gibt es eigene Operationen, zum Beispiel das Aneinanderhängen. Lektion: Variablen, Typen und Speicher.
  - Abschnitt 3.3.7, Seite 36, Verkettung: Der Operator + mit zwei Strings in Listing 3.2. Lektion: Fehlersuche: drei kurze Go-Programme, Schritt 2.
  - Abschnitt 5.3, Seiten 126–128, Zuweisung: a += "zeit". Lektion: Zuweisungen in Go.
- Unter „Wie es weitergeht“ stehen Links zu zwei Quizseiten und zu den Seiten „Strings: Bytes und Zeichen“ und „Zuweisungen in Go“. Welche Fragen die Quizseiten stellen, steht auf dieser Seite nicht.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 5“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 3.3.7
- 4.6.1
- 4.6.2.2
- 5.3
