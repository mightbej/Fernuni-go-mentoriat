# Formatierte Ausgabe mit fmt.Printf

Die Seite ist eine Lektion zu Kapitel 4 des Kurstexts, Abschnitt 4.2 mit 4.2.1 und 4.2.2, Seiten 53–58. Sie zeigt an einem eigenen Programm mit vier Kreisflächen, wie `fmt.Printf` mit einem Format-String die Ausgabe formatiert: Verb `%v` und `%f`, Mindestbreite und Genauigkeit wie in `%8.2f`, die Escape-Sequenz `\n`, `%%` und die Hinweise bei falschen Verben.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Aus Urheberrechtsgründen zeigt die Lektion nicht Listing 4.1 des Kurstexts, sondern ein eigenes Programm. Die Antwort gibt Listing 4.1 nicht wieder. Fragt der Student danach, verweist sie auf Seite 53 im Kurstext. Listing 4.1 rechnet Temperaturen um, das Programm der Lektion Kreisflächen. Der Stoff ist der gleiche.
- Das Programm `kreisflaechen.go` hat 23 Zeilen: Zeile 3 bis 6 importieren `fmt` und `math`, Zeile 9 bis 11 die Funktion `kreisflaeche` mit `return math.Pi * r * r`, Zeile 13 bis 23 `main` mit den vier Variablen `r0` bis `r3` in Zeile 14 bis 17 und vier Aufrufen von `fmt.Println` in Zeile 19 bis 22. Die Radien sind 0.5, 2, 10 und 125, die Flächen 0.7853981633974483, 12.566370614359172, 314.1592653589793 und 49087.385212340516.
- `math.Pi` ist die Kreiszahl π aus dem Paket `math`. Das Paket steht nicht in Abschnitt 4.2, die Lektion nimmt es nur für die Formel.
- In Schritt 1 bis 3 sagt der Student erst vorher, was eine Ausgabe zeigt, und deckt sie danach auf. Steht im Stand der Demo, dass eine Ausgabe oder Antwort noch verdeckt ist, nennt die Antwort sie nicht. Sie fragt nach der Vermutung des Studenten und nennt den Satz auf der Seite, der weiterhilft.
- Alle Ausgaben stammen von Go 1.26.
- Der Student kennt bisher Kapitel 2 und 3 des Kurstexts und `fmt.Println`. Escape-Sequenzen im Einzelnen kommen erst in Abschnitt 4.6.1, Gleitkommazahlen in Abschnitt 4.4.
- Im Kurstext stehen von den Verben nur `%v`, `%f` und `%%`, dazu Mindestbreite und Genauigkeit bei `%f`. Weitere Verben wie `%d` oder `%s` stehen nicht in Abschnitt 4.2. Nennt die Antwort eines, sagt sie, dass es nicht im Kurstext steht.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Das Programm mit Zeilennummern und die Angabe: ein Programm dieser Lektion, `kreisflaechen.go`.
- Ein Hinweis: Im Kurstext zeigt Listing 4.1 auf Seite 53 das gleiche Problem an vier Temperaturen. Aus Urheberrechtsgründen zeigt die Lektion ein eigenes Programm, der Stoff ist der gleiche.
- Darunter die vier Zeilen seiner Ausgabe, zum Beispiel `Radius 0.5 m ergibt 0.7853981633974483 m²`.
- Zwei Mängel der Ausgabe: Die Zahlen stehen nicht am Punkt untereinander, und die Zahl der Nachkommastellen ist uneinheitlich, bei den Radien keine oder eine, bei den Flächen zwölf bis sechzehn.
- `fmt.Println` gibt jede Zahl im Standardformat aus, `fmt.Printf` lässt das Format bestimmen.
- Eine Liste nennt die Schritte der Lektion.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 4.2

## Schritt 1: Format-String und Verben

Reiter: verben

### Aufbau der Darstellung

- Der Aufruf `fmt.Printf("Radius %v m ergibt %v m²\n", r0, kreisflaeche(r0))` und seine Ausgabe `Radius 0.5 m ergibt 0.7853981633974483 m²`, gleich der Ausgabe von Zeile 19.
- Eine Liste erklärt Format-String, Verb, die Reihenfolge der Ersetzung, die Anzahl der Argumente und das Standardformat von `%v`.
- Text zu `\n`: `fmt.Printf` beendet die Zeile nicht von selbst. Ein Link führt zur Lektion „Strings in Go“, Schritt 1.
- Zwei Punkte zu Verb und Escape-Sequenz: Verben liest nur `fmt.Printf`, `\n` gilt in jedem interpretierten String-Literal.
- Die Frage „Erst vorhersagen“ und ein Programm mit neun Zeilen: Zeile 6 `fmt.Printf("%v m", 0.5)`, Zeile 7 `fmt.Printf("%v m", 2.0)`, Zeile 8 `fmt.Println()`, Zeile 9 `fmt.Println("%v m ergibt", 0.5)`.
- Der Knopf „Was gibt Go aus?“ deckt die Ausgabe und eine Erklärung auf.

### Handgriffe

- „Was gibt Go aus?“ wählen.

### Abschnitte des Kurstexts

- 4.2.1

## Schritt 2: Breite und Nachkommastellen

Reiter: breite

### Aufbau der Darstellung

- Text zu `%f`: sechs Nachkommastellen ohne weitere Angabe.
- Eine Liste erklärt am Beispiel `%8.2f` die Mindestbreite 8, die Genauigkeit `.2` und die Schreibweisen mit nur einer Angabe, `%.2f` und `%8f`. Der Kurstext erklärt das gleiche an `%7.2f` für Listing 4.1.
- Die Frage „Erst vorhersagen“ nach der Mindestbreite für die vier Flächen mit zwei Nachkommastellen. Der Knopf „Antwort zeigen“ deckt die Antwort auf: 8, weil `49087.39` acht Zeichen hat. Für die Radien genügte 6.
- Unter „Selbst einstellen“ wählt der Student Verb, Mindestbreite und Genauigkeit. Darunter stehen das Format groß, der Aufruf für die erste Zeile und die vier Zeilen der Ausgabe von Go. Jeder ausgegebene Wert ist grau hinterlegt, so sieht man die Breite und die Leerzeichen links davon. Ein Satz der Seite beschreibt das gewählte Format.
- Eine Genauigkeit lässt die Seite nur bei `%f` zu. Bei `%v` sind die Knöpfe der Genauigkeit gesperrt.
- Bei `%.0f` wird aus dem Radius 0.5 die Ausgabe `0` und nicht `1`. Der Grund: Liegt ein Wert genau in der Mitte zwischen zwei Ausgaben, rundet Go auf die gerade Ziffer. 0.5 ist als `float64` genau gespeichert und liegt genau zwischen 0 und 1. Das steht nicht im Kurstext. Die Seite selbst erwähnt es nicht.
- Ist die Mindestbreite kleiner als der Wert, belegt der Wert so viele Zeichen, wie er braucht. Bei `%4.2f` hat `49087.39` acht Zeichen, die Punkte stehen dann nicht untereinander.

### Handgriffe

- Verb: `%v` oder `%f`.
- Mindestbreite: ohne, 4 bis 10.
- Genauigkeit: ohne, .0 bis .6, nur bei `%f`.
- „Antwort zeigen“ wählen.

### Abschnitte des Kurstexts

- 4.2.1
- 4.4

## Schritt 3: Prozentzeichen und falsche Verben

Reiter: sonderfaelle

### Aufbau der Darstellung

- Text: Go übersetzt ein Programm auch, wenn der Format-String nicht zu den Argumenten passt. `fmt.Printf` schreibt dann einen Hinweis mit `%!` in die Ausgabe.
- Vier Aufrufe zur Wahl, alle aus dieser Lektion:
  - Prozentzeichen: `fmt.Printf("Der Rabatt beträgt %v %%.\n", 15)`. `%%` steht in Abschnitt 4.2.2.
  - `%f` mit einer ganzen Zahl: `fmt.Printf("Ein Vollkreis hat %f Grad.\n", 360)`, Ausgabe `Ein Vollkreis hat %!f(int=360) Grad.` Den Fall beschreibt das Expertenwissen 4.5.
  - ein Argument zu wenig: `fmt.Printf("Radius %v m ergibt %v m²\n", 2.0)`, Ausgabe `Radius 2 m ergibt %!v(MISSING) m²`. Er steht nicht im Kurstext.
  - ein Argument zu viel: `fmt.Printf("Radius %v m.\n", 2.0, 12.5)`. Er steht nicht im Kurstext. `fmt.Printf` arbeitet erst den ganzen Format-String ab, mit dem Zeilenumbruch `\n` an seinem Ende. Erst danach hängt sie den Hinweis `%!(EXTRA float64=12.5)` an. Der Hinweis steht deshalb am Anfang der nächsten Zeile, und eine weitere Ausgabe des Programms schließt ohne Leerzeichen an ihn an. Ein Aufruf ohne `\n` am Ende hätte den Hinweis in derselben Zeile.
- Der Knopf „Was gibt Go aus?“ deckt die Ausgabe und eine Erklärung auf.
- Ein Link führt zur englischen Dokumentation des Pakets `fmt`.

### Handgriffe

- Einen der vier Aufrufe wählen.
- „Was gibt Go aus?“ wählen.

### Abschnitte des Kurstexts

- 4.2.2

## Schritt 4: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Sechs Programme zur Wahl: Kreisflächen mit Println, Printf mit `%v`, Printf mit `%f`, Printf mit `%8.2f`, das Programm aus Schritt 1 und ein Programm mit den vier Aufrufen aus Schritt 3 und einem `fmt.Println("Ende")`. Die drei Programme mit Printf sind das Programm mit den Kreisflächen, in dem Zeile 19 bis 22 `fmt.Printf` mit dem jeweiligen Format aufrufen.
- Läuft die Lektion im Editorlotsen, steht das gewählte Programm in einem Editor. Der Student ändert es und wählt „Ausführen“. Der Editorlotse übersetzt das Programm mit Go auf dem Rechner des Studenten und zeigt die Ausgabe oder die Meldung. Eine Meldung mit Zeile ist im Editor rot unterstrichen.
- Läuft die Lektion nicht im Editorlotsen, zeigt der Schritt das Programm und darunter, was Go ausgegeben hat. Der Knopf „Diese Lektion im Editorlotsen öffnen“ öffnet sie im Editorlotsen. Der Stand der Demo sagt, welcher Fall gilt.
- Vorschläge: drei Nachkommastellen mit passender Mindestbreite, ein fünfter Kreis mit dem Radius 1000, eine Überschrift über den Spalten. Beim Radius 1000 ist die Fläche 3141592.65, also zehn Zeichen mit `%.2f`. Mit `%8.2f` stehen die Punkte dann nicht mehr untereinander.
- Der Stand der Demo sagt, welches Programm gewählt ist, ob der Student es geändert hat und was der letzte Lauf ausgegeben oder gemeldet hat. Die Programmzeilen sind das Programm im Editor.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Fehlersuche: drei Go-Programme mit mehreren Fehlern“.

### Handgriffe

- Ein Programm wählen.
- Im Editorlotsen: im Editor schreiben, „Ausführen“, „Stopp“ oder „Programm zurücksetzen“ wählen.
- „Diese Lektion im Editorlotsen öffnen“ wählen.

### Abschnitte des Kurstexts

- 4.2

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Eine Tabelle nennt je Schritt der Lektion den Abschnitt, was aus dem Kurstext stammt und was nicht im Kurstext steht.
- „Wo das Thema im Kurstext noch vorkommt“: 3.2, 4.4 und 4.6.1, je mit Lektion.

### Handgriffe

- Einem Link folgen.

### Abschnitte des Kurstexts

- 4.2
