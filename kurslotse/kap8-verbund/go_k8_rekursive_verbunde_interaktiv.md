# Rekursive Verbund-Typen in Go

Die Seite erklärt rekursive Verbund-Typen: warum ein Verbund sich nicht selbst als Feld enthalten darf, wohl aber einen Zeiger auf den eigenen Typ, und wie aus solchen Verbünden ein Stammbaum und eine Kette entstehen. Sie bereitet die verkettete Liste in Kapitel 9 vor. Sie gehört zu Kapitel 8 des Kurstexts, Abschnitt 8.3.1, Seiten 215–216.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Für ein Feld eines Verbund-Typs `T` ist der Typ `T` selbst nicht zulässig. Zulässig ist der Zeiger-Typ `*T`.
- Der Typ der Seite ist `mensch` mit den Feldern `name string` und `vater, mutter *mensch`.
- Ein Verbund-Typ, dessen Felder Zeiger auf Verbünde des gleichen Typs enthalten, heißt rekursive Datenstruktur.
- Ein Zeiger-Feld enthält die Adresse eines anderen Verbunds oder `nil`. `nil` heißt hier: Es ist kein Vater oder keine Mutter eingetragen.
- Beim Zugriff auf ein Feld über einen Zeiger dereferenziert Go von selbst: `erik.vater.name` steht für `(*erik.vater).name`. Das ist Hinweis 8.2 aus Abschnitt 8.2.4.
- Der Stammbaum der Seite hat sieben Verbünde. Im Modell liegen sie an erfundenen Adressen: `achim` ab Zelle 100, `anita` ab 200, `chris` ab 300, `bert` ab 400, `bärbel` ab 500, `dagmar` ab 600, `erik` ab 700. `chris` hat den Vater `achim` und die Mutter `anita`. `dagmar` hat den Vater `bert` und die Mutter `bärbel`. `erik` hat den Vater `chris` und die Mutter `dagmar`.
- In den Rahmen der Demos steht ein Zeiger in Worten, zum Beispiel „Adresse von achim“. Im „Hauptspeicher“ darunter steht dafür die Zahl, zum Beispiel 100. Ein Block mit gelbem Schein ist neu oder gerade geändert.
- Bei einem Verbund steht „ab Zelle …“ ohne Ende. Wie viele Zellen ein `string` belegt, sagt die Seite nicht. Ein Zeiger belegt 8 Zellen.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben. Alle Ausgaben und Meldungen sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt die eine Ausnahme bei den Feld-Typen und den Begriff rekursive Datenstruktur.
- Unter „Was dazugehört“ stehen vier Punkte mit Links: die Lektionen „Hierarchische Verbünde“, „Zeiger auf Verbünde“, „Zeiger in Go“ und „Rekursion in Go“.
- Unter „Die Schritte“ steht je ein Satz zu den vier Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.3.1
- 8.3

## Schritt 1: Der Typ selbst als Feld

Reiter: verbot

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt die unzulässige Deklaration mit `vater, mutter mensch`. Darunter steht die Meldung des Compilers: `invalid recursive type: mensch refers to itself`.
- Der Text begründet es mit dem Speicher: Ein Feld vom Typ `mensch` wäre ein ganzer Verbund im Verbund, mit eigenem Vater und eigener Mutter. Der Typ hätte keine endliche Größe.
- Ein zweiter Codeblock zeigt die zulässige Deklaration mit `vater, mutter *mensch`.
- Eine Liste nennt zwei Gründe: Ein Zeiger belegt immer 8 Zellen, und ein Zeiger darf `nil` sein, dort endet die Kette.
- Eine Tabelle stellt `vater mensch` und `vater *mensch` gegenüber.
- Im Kurstext steht als Grund nur die unendlich rekursive Definition. Die Begründung mit den Zellen ist eine Ergänzung des Mentors.
- Im Kurstext steht in einer Fußnote: Der Typ `T` ist weder direkt noch indirekt als Feld-Typ zulässig, also auch nicht über einen anderen Verbund-Typ oder ein Array, das wieder `T` enthält.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.3.1

## Schritt 2: Ein Stammbaum

Reiter: stammbaum

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt acht Zeilen im Rumpf von `main`: sieben Deklarationen von `achim` bis `erik` und `fmt.Println(erik.name, "hat Vater", erik.vater.name)`.
- Unter dem Programm stehen drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“, der Rahmen von `main` mit den Verbünden, der Hauptspeicher mit einem Block je Verbund, die Ausgabe und ein Kasten mit einem Satz.
- Jeder Verbund steht im Rahmen als Kasten mit den Feldern `name`, `vater` und `mutter`.
- Der Ablauf hat acht Schritte, je Deklaration einen. Nach `chris := mensch{"Chris", &achim, &anita}` stehen im Hauptspeicher in den Feldern von `chris` die Zahlen 100 und 200. Kopiert wird kein Verbund. Die Ausgabe am Ende ist `Erik hat Vater Chris`.
- Darunter steht die lange Ausgabe des Kurstexts: `Erik hat Vater Chris und Mutter Dagmar sowie die Opas Achim und Bert und die Omas Anita und Bärbel`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Der Stammbaum lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 8.3.1
- 8.2.4

## Schritt 3: Dem Zeiger folgen

Reiter: folgen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Jeder Punkt hinter einem Zeiger ist im Hauptspeicher ein Sprung. Der Rechner liest eine Adresse aus einem Feld und geht zu dem Verbund an dieser Adresse.
- Zu sehen sind sieben Blöcke, einer je Verbund, mit Anfangsadresse, Name und den drei Feldern. In den Zeiger-Feldern stehen Zahlen oder `nil`.
- Sieben Knöpfe zeigen je einen Ausdruck. Die Blöcke der Verbünde, die der Rechner besucht, sind gelb markiert. Darunter stehen der Wert oder die Meldung und ein Satz.
- Die sieben Ausdrücke:
  1. `erik.name` ist „Erik“, kein Sprung.
  2. `erik.vater` ist die Adresse von `chris`, im Modell 300.
  3. `erik.vater.name` ist „Chris“, ein Sprung.
  4. `erik.vater.vater.name` ist „Achim“, zwei Sprünge über 300 und 100.
  5. `erik.mutter.mutter.name` ist „Bärbel“, zwei Sprünge über 600 und 500.
  6. `erik.vater.vater.vater` ist `nil`. `fmt.Println` gibt dafür `<nil>` aus.
  7. `erik.vater.vater.vater.name` übersetzt Go, das Programm bricht ab mit `panic: runtime error: invalid memory address or nil pointer dereference`.
- Ein gelber Kasten sagt: Vor einem Sprung prüft ein Programm mit `!= nil`, ob der Zeiger auf einen Verbund zeigt.

### Handgriffe

- Einen der sieben Ausdrücke wählen.
- Die „Auflösung“ aufklappen.
- Ein eigener Ausdruck lässt sich nicht eingeben.

### Abschnitte des Kurstexts

- 8.3.1
- 7.4

## Schritt 4: Eine Kette

Reiter: kette

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt drei Zeilen: `for p := &erik; p != nil; p = p.vater {`, `fmt.Print(p.name, " ")`, `}`.
- Im Rahmen und im Hauptspeicher stehen `achim`, `chris`, `erik` und der Zeiger `p` in den Zellen 800 bis 807.
- Der Ablauf hat sieben Schritte: `p` ist 700 und zeigt auf `erik`, Ausgabe „Erik“. `p = p.vater` setzt `p` auf 300, Ausgabe „Chris“. Dann 100, Ausgabe „Achim“. Dann liest die Schleife aus `achim` den Wert `nil`, `p` ist `nil`, die Bedingung `p != nil` ist nicht erfüllt, die Schleife endet. Die ganze Ausgabe ist `Erik Chris Achim `.
- Die Schleife ist ein eigenes Beispiel und steht nicht im Kurstext.
- Unter „Von hier zur verketteten Liste“ steht der Typ aus Abschnitt 9.2 des Kurstexts: `listenElement` mit `wert int` und `nächstes *listenElement`.
- Eine Tabelle stellt Stammbaum und verkettete Liste gegenüber: Das Feld `vater` entspricht dem Feld `nächstes`, das Ende ist in beiden Fällen `nil`, und eine Schleife läuft hindurch, bis der Zeiger `nil` ist.
- Der Text sagt: Ein Array hat eine feste Länge, eine Kette aus Verbünden wächst mit jedem Verbund, der angehängt wird.
- Wie eine Liste aufgebaut, durchlaufen und sortiert gehalten wird, steht in Abschnitt 9.2. Diese Seite greift dem nicht vor. Im Kurstext gibt es dort außerdem den Typ `sortierteListe` mit dem Feld `kopf *listenElement`.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.3.1
- 9.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Der Typ selbst als Feld, 8.3.1, Seite 215: Beide Deklarationen stammen aus dem Kurstext. Dort steht als Grund die unendlich rekursive Definition. Die Meldung des Compilers und die Begründung mit den Zellen sind eine Ergänzung, geprüft mit Go 1.26.
  - 2 Ein Stammbaum, 8.3.1, Seiten 215–216: Die sieben Verbünde und die lange Ausgabe stammen aus dem Kurstext. Die kurze Ausgabe in Zeile 8 der Demo ist ergänzt.
  - 3 Dem Zeiger folgen, 8.3.1, Seiten 215–216: Die Ausdrücke 3 bis 5 stehen in der Ausgabe des Kurstexts. Die Ausdrücke 6 und 7 mit nil sind eine Ergänzung, geprüft mit Go 1.26.
  - 4 Eine Kette, 8.3.1, Seite 216; 9.2, Seiten 229–236: Der Typ listenElement stammt aus Abschnitt 9.2. Die Schleife über die Väter ist ein eigenes Beispiel, geprüft mit Go 1.26.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 7.5, Seite 185, Rekursion: Rekursiv ist dort eine Funktion, hier ein Typ. Lektion: Rekursion in Go.
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.
  - Abschnitt 9.2, Seiten 229–236, verkettete Listen: Lektion: Verkettete Listen in Go.
- Unter „Wie es weitergeht“ stehen drei Punkte: die Lektion „Anonyme Felder und eingebettete Verbünde“, die Lektionen „Verkettete Listen in Go“ und „Binäre Suchbäume in Go“ zu Kapitel 9 und das Go-Programm `S224_rekursive_verbunde_demo.go`.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.3
- 7.4
- 7.5
- 8.3.1
- 9.2
