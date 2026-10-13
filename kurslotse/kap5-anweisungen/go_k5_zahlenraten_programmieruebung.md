# Programmierübung: Zahlenraten mit Statistik

Die Seite baut ein Go-Programm in sieben Schritten auf: ein Zahlenratespiel über mehrere Runden mit Statistik und Wertung. Sie setzt den Exkurs „Von der Problemstellung zum Go-Programm“ fort, dessen Ablauf aus Stufe 5 hier zu Go wird. Das Programm braucht fast alle Anweisungen aus Kapitel 5 des Kurstexts.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Das Programm ist ein eigenes Beispiel des Mentoriats und steht nicht im Kurstext. Es ist mit Go 1.26 übersetzt und gespielt.
- Auf der Seite läuft kein Go. Das Spiel lässt sich auf der Seite nicht spielen. Wer spielen will, kopiert das Programm aus Schritt 6 in eine Datei `zahlenraten.go` und startet es im Terminal mit `go run zahlenraten.go`.
- Der Student kann auf der Seite keinen Code ändern. Einen Lauf Anweisung für Anweisung gibt es auf dieser Seite nicht.
- Das ganze Programm hat 115 Zeilen und steht in der Funktion `main`. Eigene Funktionen kommen nicht vor.
- Die Regeln des Spiels: Zufallszahl von 1 bis 100, höchstens 7 Versuche je Runde, nach jedem Tipp die Meldung „zu klein“, „zu groß“ oder „richtig“.
- Die drei Konstanten: `MAX_VERSUCHE = 7`, `MIN_ZAHL = 1`, `MAX_ZAHL = 100`.
- Die drei globalen Variablen, die über alle Runden zählen: `anzahlSiege`, `anzahlNiederlagen`, `gesamtVersuche`. `gesamtVersuche` ist die Summe der Versuche aus den gewonnenen Runden.
- Die Variablen einer Runde: `zielZahl` für die gesuchte Zahl, `versuchsZaehler` für die Nummer des Versuchs, `gewonnen` als Flag-Variable, `geratenZahl` für den Tipp.
- Das Programm hat zwei Schleifen ineinander. Die äußere `for {` wiederholt die Runden. Die innere `for versuchsZaehler < MAX_VERSUCHE {` wiederholt die Tipps einer Runde.
- Die Namen der Konstanten in Großbuchstaben mit Unterstrich sind in Go unüblich. Im Kurstext heißen Konstanten zum Beispiel `maxAnzahl`. Go übersetzt beide Schreibweisen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Code und Demo gibt es hier nicht.
- Oben steht der Bezug zum Exkurs „Von der Problemstellung zum Go-Programm“.
- Die Liste „Was das Programm tun soll“ nennt sechs Aufgaben: Zahl wählen, raten lassen, Runde beenden, mitzählen, bewerten, weiterspielen.
- Darunter steht, dass das ganze Programm in `main` steht, weil Funktionen erst in Kapitel 7 behandelt werden.
- Eine Liste sagt, was die sieben Schritte der Seite zeigen.
- Ein gelber Kasten sagt, dass das Programm mit Go 1.26 übersetzt und gespielt ist und dass auf der Seite kein Go läuft.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Dem Link zum Exkurs folgen. Er ist eine eigene HTML-Seite.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.1

## Schritt 1: Rahmen und Aufbau

Reiter: rahmen

### Aufbau der Darstellung

- Oben steht der Rahmen des Programms in einem dunklen Kasten: `package main`, die Importe `fmt` und `math/rand`, die leere Funktion `main` mit einem Kommentar.
- `fmt` ist das Paket für Ausgabe und Eingabe. `math/rand` ist das Paket für Zufallszahlen und steht nicht im Kurstext.
- Der Teil „Der Aufbau in Worten“ ist eine nummerierte Liste mit fünf Punkten: feste Werte benennen, drei Zähler über alle Runden, Begrüßung, die Wiederholung je Runde mit fünf Unterpunkten, Verabschiedung.
- Der Teil „Zwei Schleifen ineinander“ nennt die äußere Schleife für die Runden und die innere Schleife für die Tipps einer Runde. Die innere steht im Schleifenrumpf der äußeren.
- Die Übung baut erst die innere Schleife in Schritt 3 und legt die äußere in Schritt 5 darum.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Der Codeblock ist fest.
- Dem Link zum Exkurs folgen.

### Abschnitte des Kurstexts

- 3.3
- 5.5.3

## Schritt 2: Konstanten und Variablen

Reiter: daten

### Aufbau der Darstellung

- Der erste Codeblock zeigt die Deklarationen zwischen den Importen und `main`: einen Block `const ( … )` mit den drei Konstanten und einen Block `var ( … )` mit den drei Zählern vom Typ `int`.
- Die drei Zähler haben keinen Wert in der Deklaration und bekommen den Null-Wert 0.
- Der Teil „Warum die Zähler nicht in der Schleife stehen“ sagt: Die Zähler müssen ihren Wert von Runde zu Runde behalten. Im Schleifenrumpf der äußeren Schleife begännen sie in jeder Runde bei 0.
- Die Zähler sind als globale Variablen deklariert, außerhalb von `main`. Nötig wäre das nicht: Sie könnten ebenso in `main` vor der äußeren Schleife deklariert sein. Die Übung zeigt damit die globale Deklaration aus Abschnitt 5.1.3.
- Der zweite Codeblock zeigt die Begrüßung, die ersten drei Anweisungen in `main`: eine Zeile mit `fmt.Println` und zwei Zeilen mit `fmt.Printf`.
- `%d` in `fmt.Printf` setzt einen ganzzahligen Wert ein. `\n` steht für einen Zeilenumbruch.
- Die Ausgabe der Begrüßung: „🎯 Willkommen beim Zahlenraten-Spiel!“, „Rate eine Zahl zwischen 1 und 100.“, „Du hast 7 Versuche pro Runde.“

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Die Codeblöcke sind fest.
- Den Links zu den Seiten „Konstanten“ und „Variablendeklaration“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.2
- 5.1.3
- 4.2

## Schritt 3: Rateschleife

Reiter: raten

### Aufbau der Darstellung

- Der Codeblock zeigt eine Runde des Spiels, 39 Zeilen. Im ganzen Programm sind es die Zeilen 29 bis 67. Dort stehen sie im Schleifenrumpf der äußeren Schleife und sind um zwei Stufen eingerückt. Im Codeblock dieses Schritts beginnen sie am linken Rand.
- Erste Anweisung: `zielZahl := rand.Intn(MAX_ZAHL-MIN_ZAHL+1) + MIN_ZAHL`. `rand.Intn(n)` liefert eine zufällige ganze Zahl von 0 bis n − 1. Mit n gleich 100 und plus 1 ergibt das eine Zahl von 1 bis 100.
- Dann: `versuchsZaehler := 0` und `gewonnen := false`.
- Die Schleife: `for versuchsZaehler < MAX_VERSUCHE {`. Ihre erste Anweisung ist `versuchsZaehler++`. Danach gibt das Programm „Versuch 3/7 - Deine Zahl: “ mit der Nummer des Versuchs aus und liest den Tipp mit `fmt.Scan(&geratenZahl)`.
- Weil `versuchsZaehler++` am Anfang des Schleifenrumpfs steht, nennt die Ausgabe gleich die richtige Nummer des Versuchs: Der Zähler beginnt bei 0 und steht im ersten Schleifendurchlauf auf 1.
- Die if-Anweisung hat drei Zweige. `if geratenZahl == zielZahl`: Glückwunsch, `anzahlSiege++`, `gesamtVersuche += versuchsZaehler`, `gewonnen = true`, `break`. `else if geratenZahl < zielZahl`: Meldung „Zu klein“. `else`: Meldung „Zu groß“.
- `break` verlässt die Rateschleife sofort.
- Hinter der Schleife steht `if !gewonnen {`. Dann meldet das Programm „Leider verloren“ mit der gesuchten Zahl und erhöht `anzahlNiederlagen`.
- Die Schleife endet auf zwei Wegen: durch `break` nach einem Treffer oder regulär, wenn `versuchsZaehler` den Wert 7 hat und `7 < 7` den Wert `false` ergibt.
- Der Teil „Zum Weiterdenken“ fragt, wie der Kopf der Rateschleife ohne `break` lauten müsste. Die Seite nennt die Antwort nicht, sie verweist auf Schritt 3 der Seite „break-Anweisung“. Die Antwort wäre `for versuchsZaehler < MAX_VERSUCHE && !gewonnen {`.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Der Codeblock ist fest. Tipps lassen sich hier nicht eingeben.
- Den Links zur Seite „break-Anweisung“ folgen. Sie ist eine eigene HTML-Seite.

### Abschnitte des Kurstexts

- 5.5.3.1
- 5.5.1
- 5.5.3.3
- 5.3

## Schritt 4: Statistik und Wertung

Reiter: wertung

### Aufbau der Darstellung

- Der Codeblock zeigt die Ausgabe nach jeder Runde, 31 Zeilen. Im ganzen Programm sind es die Zeilen 69 bis 99.
- `gesamtSpiele := anzahlSiege + anzahlNiederlagen` rechnet die Zahl der gespielten Runden aus.
- Drei Zeilen mit `fmt.Println` zeichnen einen Rahmen aus Strichzeichen mit dem Wort STATISTIK. Danach folgen „Gesamt gespielt“, „Gewonnen“ und „Verloren“.
- `if anzahlSiege > 0 {` berechnet den Durchschnitt nur, wenn es mindestens einen Sieg gibt. Der else-Zweig gibt „Noch keine Siege!“ aus.
- `durchschnitt := float64(gesamtVersuche) / float64(anzahlSiege)`: Ohne die Konversion teilte Go ganze Zahlen ohne Rest. 7 Versuche in 2 gewonnenen Runden ergäben 3 statt 3,5.
- `%.2f` gibt den Durchschnitt mit zwei Nachkommastellen aus.
- Die switch-Anweisung ohne switch-Ausdruck bewertet den Durchschnitt: `case durchschnitt <= 3.0:` Meister, `case durchschnitt <= 4.0:` Experte, `case durchschnitt <= 5.5:` Fortgeschritten, `default:` Anfänger.
- Eine Tabelle auf der Seite nennt die gleichen vier Grenzen in Worten.
- Die Reihenfolge der case-Abschnitte entscheidet: Go nimmt den ersten case-Ausdruck, der `true` ergibt. Ein Durchschnitt von 2,5 erfüllt alle drei Bedingungen und bekommt „Meister“.
- `var kategorie string` ist ausführlich und ohne Wert deklariert. Die Variable bekommt ihren Wert in der switch-Anweisung.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Der Codeblock ist fest.
- Den Links zu den Seiten „Typkonvertierung“ und „switch-Anweisung“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.2
- 4.7
- 5.5.1

## Schritt 5: Mehrere Runden

Reiter: runden

### Aufbau der Darstellung

- Der Codeblock zeigt die äußere Schleife als Gerüst: `for {`, zwei Kommentare an der Stelle der Rateschleife und der Statistik, die Frage „Nochmal spielen? (j/n): “, `var antwort string`, `fmt.Scan(&antwort)`, die if-Anweisung mit `break`, die Ausgabe „--- Neue Runde ---“, `}` und dahinter die Verabschiedung.
- `for {` hat keine Schleifenbedingung. Die Schleife läuft, bis `break` sie verlässt.
- `if antwort == "n" || antwort == "N" {` ist erfüllt, wenn der Spieler ein kleines oder ein großes n eingibt. `||` bedeutet „oder“. Jede andere Eingabe startet eine neue Runde.
- Das Programm enthält zwei `break`. Das `break` nach dem Treffer steht in der Rateschleife und beendet nur die Runde. Das `break` nach der Antwort „n“ steht in der äußeren Schleife und beendet das Spiel.
- Die Verabschiedung lautet „👋 Vielen Dank fürs Spielen! Bis zum nächsten Mal!“.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen. Der Codeblock ist fest.
- Den Links zu den Seiten „for-Anweisung mit for-Klausel“ und „break-Anweisung“ folgen. Beide sind eigene HTML-Seiten.

### Abschnitte des Kurstexts

- 5.5.3.2
- 5.5.3.3
- 4.5

## Schritt 6: Das ganze Programm

Reiter: programm

### Aufbau der Darstellung

- Zu sehen ist das ganze Programm mit 115 Zeilen in einem dunklen Kasten, mit Zeilennummern am linken Rand.
- Über dem Programm stehen fünf Knöpfe: „nichts markieren“, „Konstanten und Variablen“, „Rateschleife“, „Statistik und Wertung“, „äußere Schleife“. Der gewählte Knopf ist grün.
- Ein Knopf markiert seinen Teil im Programm gelb: Konstanten und Variablen sind die Zeilen 8 bis 20, die Rateschleife die Zeilen 29 bis 67, Statistik und Wertung die Zeilen 69 bis 99, die äußere Schleife die Zeilen 28 bis 112.
- Gelb bedeutet hier nur: Diese Zeilen gehören zum gewählten Teil. Einen Lauf zeigt die Seite nicht.
- Der Knopf „Programm kopieren“ legt das Programm in die Zwischenablage.
- Wichtige Zeilen: 22 `func main() {`, 28 `for {`, 30 die Zufallszahl, 37 der Kopf der Rateschleife, 38 `versuchsZaehler++`, 43 `fmt.Scan(&geratenZahl)`, 46 der Vergleich auf Gleichheit, 54 das erste `break`, 64 `if !gewonnen {`, 81 die Berechnung des Durchschnitts, 86 `switch {`, 107 die Prüfung der Antwort, 108 das zweite `break`, 114 die Verabschiedung.
- Unter dem Programm steht, wie man es startet: Datei `zahlenraten.go` anlegen, Programm hineinkopieren, im Terminal `go run zahlenraten.go` eingeben.
- Der Teil „Zum Ausprobieren“ nennt drei Aufgaben für das eigene Terminal: immer die Mitte tippen, `MAX_VERSUCHE` auf 5 ändern, eine Runde mit Absicht verlieren.
- Wer `MAX_VERSUCHE` ändert, ändert nur Zeile 10. Dafür sind Konstanten da.
- Wer immer die Mitte des möglichen Bereichs tippt, gewinnt jede Runde: Sieben Versuche reichen für 100 Zahlen. Das steht im Exkurs in Stufe 3.

### Handgriffe

- Einen der fünf Knöpfe wählen, um einen Teil zu markieren.
- „Programm kopieren“ wählen.
- Das Programm lässt sich auf der Seite nicht ändern und nicht starten. Die drei Aufgaben unter „Zum Ausprobieren“ macht der Student im eigenen Terminal.

### Abschnitte des Kurstexts

- 5.5.3
- 5.5.1
- 5.2

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit vier Spalten: Stelle im Programm, was dort steht, Abschnitt im Kurstext, Seite dazu. Sie hat zwölf Zeilen.
- Die Zuordnung: Konstanten zu 5.2, globale Variablen zu 5.1.1 und 5.1.3, kurze Variablendeklaration zu 5.1.2, Zuweisungen zu 5.3, `if … else if … else` zu 5.5.1, `switch` zu 5.5.2, die Rateschleife zu 5.5.3.1, `for {` zu 5.5.3.2 mit dem Expertenwissen 5.16, `break` zu 5.5.3.3, `float64(…)` zu 4.7, `||` und `!` zu 4.5, `fmt.Printf` zu 4.2.
- Die letzte Spalte verlinkt je Zeile eine HTML-Seite des Mentoriats. Zu `fmt.Printf` gibt es noch keine.
- Darunter steht, was über den Kurstext hinausgeht: das Paket `math/rand` mit `rand.Intn`.
- Am Ende steht ein Link zur Übung „Sortieren mit Bubble Sort“ in Kapitel 6.

### Handgriffe

- In diesem Schritt lässt sich nichts einstellen.
- Den Links in der Tabelle folgen. Jede Zeile führt zu einer eigenen HTML-Seite.

### Abschnitte des Kurstexts

- 5.1
- 5.2
- 5.3
- 5.5.1
- 5.5.2
- 5.5.3
