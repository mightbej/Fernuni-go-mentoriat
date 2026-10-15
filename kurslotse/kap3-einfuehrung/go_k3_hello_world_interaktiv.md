# Das erste Go-Programm: Hello, World!

Die Seite ist eine Lektion zu Kapitel 3 des Kurstexts, Abschnitt 3.1 und 3.2, Seiten 21–24. Sie erklärt Listing 3.1, das Hello-World-Programm: Paketklausel, Importdeklaration, die Funktion `main` und den Aufruf von `fmt.Println`. Dazu kommen Pakete und die Standardbibliothek aus Abschnitt 3.1 und der Leerraum.

Die Seite hat sieben Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Listing 3.1 hat sieben Zeilen: 1 `package main`, 2 leer, 3 `import "fmt"`, 4 leer, 5 `func main() {`, 6 `fmt.Println("Hello, World!")`, 7 `}`. Eingerückt ist mit einem Tabulator.
- In Schritt 2 bis 4 sagt der Student erst vorher, was Go mit einer geänderten Fassung macht, und deckt danach die Meldung oder Ausgabe auf. Diese Beschreibung nennt die Meldungen deshalb nicht. Hat der Student sie aufgedeckt, stehen sie im Stand der Demo. Steht dort, dass sie noch verdeckt ist, nennt die Antwort weder die Meldung noch das Ergebnis. Sie fragt den Studenten nach seiner Vermutung und sagt, welcher Satz auf der Seite weiterhilft, zum Beispiel die Regel, dass jedes importierte Paket benutzt werden muss.
- Alle Meldungen und Ausgaben stammen von Go 1.26.1.
- Der Student kennt bisher Kapitel 2 des Kurstexts: Variable, Anweisung, Ausdruck, Unterprogramm, Compiler. Eigene Funktionen in Go kommen erst in Abschnitt 3.4, Variablen in Go in Abschnitt 3.3.3.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen sind Listing 3.1 mit Zeilennummern und die Quelle: Listing 3.1, `hello_world.go`, Abschnitt 3.2, Seite 22.
- Unter „So startest du es“ stehen der Befehl `go run hello_world.go` und die Ausgabe `Hello, World!`. Go übersetzt erst und führt dann aus. Ein Link führt zur Lektion „Höhere und niedrigere Programmiersprachen“, Schritt 2.
- Die Tradition des Hello-World-Programms stammt aus dem Buch „The C Programming Language“ von Kernighan und Ritchie, 1978.
- Eine Liste nennt die Schritte der Lektion.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 3.2
- 2.7

## Schritt 1: Zeile für Zeile

Reiter: zeilen

### Aufbau der Darstellung

- Listing 3.1 mit Zeilennummern. Jede Zeile lässt sich anklicken, die gewählte Zeile ist gelb.
- Darunter steht die Erklärung der gewählten Zeile:
  - Zeile 1: Paketklausel. Sie steht in der ersten nicht leeren Zeile jeder Quellcodedatei und legt das Paket fest. Aus dem Paket `main` baut Go ein ausführbares Programm statt einer Bibliothek.
  - Zeile 2 und 4: Leerzeilen für die Übersicht.
  - Zeile 3: Importdeklaration. Sie macht das Paket `fmt` verfügbar, Abkürzung für „format“, gesprochen „fumpt“. Es gehört zur Standardbibliothek und liefert Funktionen zum Einlesen, Formatieren und Ausgeben.
  - Zeile 5: Beginn der Funktionsdeklaration: `func`, der Name `main`, leere runde Klammern für keine Parameter, `{` als Beginn des Funktionsrumpfs.
  - Zeile 6: Aufruf von `fmt.Println`, gesprochen „print line“. Das Argument ist das String-Literal `"Hello, World!"`, sein Wert die Zeichenkette ohne Anführungszeichen. `fmt.Println` gibt das Argument aus und hängt einen Zeilenumbruch an. Die Zeile ist eingerückt, weil sie im Rumpf steht.
  - Zeile 7: Ende des Funktionsrumpfs. Hier endet `main` und mit ihr das Programm.

### Handgriffe

- Eine Zeile anklicken.
- „vorige Zeile“ oder „nächste Zeile“ wählen.

### Abschnitte des Kurstexts

- 3.2

## Schritt 2: Pakete und Import

Reiter: pakete

### Aufbau der Darstellung

- Text zu Paketen: Ein Paket fasst Programmteile mit gemeinsamem Zweck zusammen. Ein Programm importiert es und verwendet es wieder. In anderen Sprachen heißt es Bibliothek oder Modul.
- Text zur Standardbibliothek: über 150 Pakete für Text, Dateien und Netzwerk, Mathematik, Sortieren, Zeit, Dateiformate und Kryptographie. Die Liste steht unter `https://pkg.go.dev/std`.
- Drei Regeln für `import`: direkt nach der Paketklausel, mehrere oder keine Importdeklaration, jedes importierte Paket muss benutzt werden, sonst übersetzt Go nicht.
- Die Frage „Erst vorhersagen“ und zwei Fassungen zur Wahl:
  - Fassung A ohne `import "fmt"`, fünf Zeilen, Zeile 4 ruft `fmt.Println` auf.
  - Fassung B mit `import "fmt"` in Zeile 3 und leerem Rumpf von `main`.
- Der Knopf „Was meldet Go?“ deckt die Meldung und eine Erklärung auf. Beide Fassungen übersetzt Go nicht.

### Handgriffe

- Fassung A oder B wählen.
- „Was meldet Go?“ wählen.

### Abschnitte des Kurstexts

- 3.1
- 3.2

## Schritt 3: Leerraum

Reiter: leerraum

### Aufbau der Darstellung

- Text: Leerzeichen, Tabulatoren und Zeilenumbrüche heißen Leerraum. Er trennt die Teile des Programms und wird sonst nicht beachtet. Ob eines oder mehrere, ist gleich.
- Zwei Ausnahmen: Im String-Literal gehört jedes Leerzeichen zur Zeichenkette. Zeilenumbrüche haben an manchen Stellen eine Bedeutung, das erklärt Abschnitt 3.3.2. Ein Link führt zur Lektion „Fehlersuche: drei kurze Go-Programme“, Schritt 3.
- Die Frage „Erst vorhersagen“ und eine Fassung mit viel Leerraum: `package` und `main` in zwei Zeilen, mehrere Leerzeichen nach `import`, nach `func`, in `main( )` und um das Argument, im String-Literal `"Hello,   World!"` drei Leerzeichen.
- Der Knopf „Was gibt Go aus?“ deckt auf, ob Go die Fassung übersetzt und was sie ausgibt.
- Text „Leerraum für den Leser“: Leerzeilen zwischen den Teilen, Einrücken mit Tabulator nach der Standardformatierung von Go.

### Handgriffe

- „Was gibt Go aus?“ wählen.

### Abschnitte des Kurstexts

- 3.2
- 3.3.2

## Schritt 4: main als Einstiegspunkt

Reiter: main

### Aufbau der Darstellung

- Drei Punkte: Go ruft `main` beim Start von selbst auf. Endet `main`, endet das Programm. Jedes Programm muss `main` deklarieren.
- Eine Demo geht den Lauf von Listing 3.1 in vier Ständen durch: vor dem Start, Aufruf von `main` mit Zeile 5 gelb, Aufruf von `fmt.Println` mit Zeile 6 gelb und der Ausgabe, Ende bei Zeile 7.
- Ein Link führt zur Lektion „Unterprogramme und Funktionen“, Schritt 1, Aufruf und Rücksprung.
- Die Frage „Erst vorhersagen“ und eine Fassung, in der die Funktion `start` statt `main` heißt. Der Knopf „Was meldet Go?“ deckt die Meldung und eine Erklärung auf. Ist sie aufgedeckt, gilt: Die Meldung nennt keine Zeile, weil Go den Fehler erst beim Zusammensetzen des Programms bemerkt. `runtime.main_main·f` ist der Name der Stelle in Go, die `main` aufrufen soll.

### Handgriffe

- „weiter“, „zurück“ oder „von vorn“ wählen.
- „Was meldet Go?“ wählen.

### Abschnitte des Kurstexts

- 3.2
- 2.6

## Schritt 5: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Fünf Fassungen zur Wahl: Listing 3.1, ohne import, import ohne Aufruf, viel Leerraum, func start().
- Läuft die Lektion im Editorlotsen, steht die gewählte Fassung in einem Editor. Der Student ändert sie und wählt „Ausführen“. Der Editorlotse übersetzt das Programm mit Go auf dem Rechner des Studenten und zeigt die Ausgabe oder die Meldung. Eine Meldung mit Zeile ist im Editor rot unterstrichen. Wartet das Programm auf eine Eingabe, tippt der Student sie in die Zeile unter der Ausgabe.
- Läuft die Lektion nicht im Editorlotsen, zum Beispiel auf GitHub Pages, zeigt der Schritt das Programm und darunter, was Go 1.26.1 dazu gemeldet oder ausgegeben hat. Darüber steht, wie der Student die Lektion im Editorlotsen öffnet: Editorlotse starten, im Menü „Programm wählen“ unter „Kapitel 3“ diese Lektion wählen. Läuft die Lektion im Editorlotsen, führt ein Link zurück zu seinem Editor. Der Stand der Demo sagt, welcher Fall gilt.
- Vorschläge: eigenen Namen ausgeben, eine zweite Zeile mit `fmt.Println`, einen Fehler einbauen.
- Der Stand der Demo sagt, welche Fassung gewählt ist, ob der Student sie geändert hat und was der letzte Lauf ausgegeben oder gemeldet hat. Die Programmzeilen sind die Fassung im Editor.

### Handgriffe

- Eine Fassung wählen.
- Im Editor schreiben, „Ausführen“, „Stopp“ oder „Fassung zurücksetzen“ wählen.
- Eine Eingabe in die Zeile unter der Ausgabe tippen.

### Abschnitte des Kurstexts

- 3.2

## Schritt 6: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle nennt je Schritt der Lektion den Abschnitt, was aus dem Kurstext stammt und was nicht im Kurstext steht.
- „Wo das Thema im Kurstext noch vorkommt“: 2.6, 2.7, 3.3.2, 3.4 und 4.6, je mit Lektion oder Quiz.
- „Wie es weitergeht“: die Lektion zur Fehlersuche und die Übersicht der Quiz zu Kapitel 3.

### Handgriffe

- Einem Link folgen.

### Abschnitte des Kurstexts

- 3.1
- 3.2
