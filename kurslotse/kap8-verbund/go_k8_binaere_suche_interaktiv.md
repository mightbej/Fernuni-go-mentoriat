# Binäre Suche in einem Verbund-Array

Die Seite erklärt die binäre Suche aus Abschnitt 8.4.1 des Kurstexts, Seiten 221–225, und die Funktion `binSuche` aus Listing 8.1. Die Funktion sucht in einem Array von Verbünden des Typs `person`, das nach dem Feld `name` sortiert ist, und liefert die Telefonnummer oder −1.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Typ ist `person` mit den Feldern `name string` und `telefonnummer int`, wie in Listing 8.1. Die Konstante `arrlen` ist 10.
- Das Telefonbuch der Seite ist sortiert: `telbuch := [10]person{{"Arnold", 3310}, {"Becker", 5522}, {"Fischer", 6398}, {"Huber", 4456}, {"Koch", 7745}, {"Lange", 8127}, {"Meier", 4711}, {"Schulz", 1234}, {"Wagner", 9081}, {"Zimmer", 2063}}`. Index 0 ist Arnold, Index 9 ist Zimmer. Es ist das Telefonbuch aus der Lektion „Bubblesort auf einem Verbund-Array“, nachdem es sortiert ist.
- Die Funktion: `func binSuche(personArr *[arrlen]person, links, rechts int, suchwert string) int`. `links` und `rechts` sind der erste und der letzte Index des Suchbereichs, beide gehören dazu. Ist `rechts >= links`, rechnet sie `mitte := links + (rechts-links)/2`. Stimmt `personArr[mitte].name` mit `suchwert` überein, ist `ergebnis` die Telefonnummer. Ist `personArr[mitte].name > suchwert`, ruft sie sich mit `links, mitte-1` auf, sonst mit `mitte+1, rechts`. Ist `rechts < links`, ist `ergebnis` −1. Am Ende steht `return ergebnis`.
- Der erste Aufruf lautet `binSuche(&telbuch, 0, arrlen-1, suchwert)`. Im Kurstext steht als Beispiel `binSuche(&telbuch, 0, arrlen-1, "Müller")`. Im Telefonbuch der Seite gibt es keinen Müller.
- Die Seite nennt einen Vergleich mit dem Namen am Index `mitte` einen „Blick“. Das Wort steht nicht im Kurstext.
- Die Blicke für jeden Namen, geprüft mit Go 1.26: Koch 1 Blick. Becker und Schulz 2 Blicke. Arnold, Fischer, Lange und Wagner 3 Blicke. Huber, Meier und Zimmer 4 Blicke. Abel 3 Blicke und danach ein Aufruf mit leerem Bereich, Maier und Zuse 4 Blicke und danach ein Aufruf mit leerem Bereich. Mehr als 4 Blicke braucht das Telefonbuch nie.
- Die Suche von vorn aus Abschnitt 8.1 sieht die Einträge der Reihe nach an. Sie braucht für den Namen am Index i genau i + 1 Vergleiche und für einen fehlenden Namen 10.
- Go vergleicht zwei Strings Byte für Byte. Die Namen der Seite haben einen großen Anfangsbuchstaben und keine Umlaute, die Reihenfolge von Go ist deshalb die alphabetische. Wie Go Strings vergleicht, zeigt Schritt 1 der Lektion „Bubblesort auf einem Verbund-Array“.
- Die binäre Suche setzt voraus, dass das Array nach dem Feld sortiert ist, nach dem gesucht wird, und dass `suchwert` höchstens einmal vorkommt. Das steht im Kommentar über Listing 8.1.
- Im Kurstext steht keine Begründung, warum `mitte` als `links + (rechts-links)/2` gerechnet wird und nicht als `(links+rechts)/2`. Beide liefern hier den gleichen Wert. Die Seite nennt in einem Kasten einen Vorteil der ersten Form: Die Summe kann bei sehr großen Indizes überlaufen, die Differenz nicht. Dieser Kasten geht über den Kurstext hinaus.
- Die Schreibweise O(log n) steht nicht im Kurstext. Dort heißt die Laufzeit „logarithmisch“.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code und keinen eigenen Namen eingeben. Alle Ausgaben und Zahlen sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt den Anlass: Am Ende von Abschnitt 8.1 sucht eine `for`-Schleife einen Eintrag von vorn. Die binäre Suche nutzt aus, dass das Telefonbuch sortiert ist.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Verbund-Variablen und Feld-Variablen“, „Zeiger in Go“, „Unterprogramme und Funktionen“ (Schritt 4, Rekursion), „Rekursion in Go“, „Bubblesort auf einem Verbund-Array“ (Schritt 1, Strings vergleichen) und auf den Exkurs „Das Kopier-Problem“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.1

## Schritt 1: Die Idee

Reiter: idee

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt den Typ `person` und das sortierte Telefonbuch.
- Ein Streifen zeigt die zehn Einträge nebeneinander. Jeder Eintrag zeigt oben den Index, in der Mitte den Namen, darunter die Telefonnummer und darunter die Wörter „links“, „mitte“ oder „rechts“, wenn der Index eine Grenze oder die Mitte des Suchbereichs ist.
- Grau und gestrichelt sind Einträge außerhalb des Suchbereichs. Gelb ist der Eintrag in der Mitte, mit dem die Suche gerade vergleicht. Grün ist ein gefundener Eintrag. Ein hellblauer Rand markiert den gesuchten Namen.
- Neben „Namen, die fehlen:“ stehen drei Knöpfe: Abel, Maier und Zuse.
- Darunter stehen die Knöpfe „nächster Blick“ und „von vorn“ und ein Kasten mit einem Satz. Nach dem letzten Blick nennt der Satz auch, wie viele Namen die Suche von vorn angesehen hätte.
- Unter dem Kasten steht die Idee aus Abschnitt 8.4.1 in einem Absatz.
- Ein gelber Kasten verlinkt Stufe 3 des Exkurses „Von der Problemstellung zum Go-Programm“ zu Kapitel 3. Dort errät der Spieler eine Zahl von 1 bis 100 mit höchstens sieben Tipps, das Spiel meldet „größer“ oder „kleiner“. Wer immer die Mitte des möglichen Bereichs tippt, rät mit der Strategie der binären Suche. Der Exkurs zeigt das mit einem Rechner und einer Tabelle. Beim Zahlenraten macht der Spieler die Suche selbst, ein Programm mit Funktionen und Arrays gibt es dort nicht.
- In diesem Schritt gibt es kein Programm. Die Funktion steht in Schritt 2.

### Handgriffe

- Einen der zehn Namen im Streifen anklicken oder einen der drei fehlenden Namen wählen. Die Suche beginnt dann von vorn.
- „nächster Blick“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 8.4.1

## Schritt 2: Die Funktion

Reiter: funktion

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Absatz sagt, warum die Funktion einen Zeiger und zwei Indizes bekommt und keine Kopie des halben Arrays.
- Ein Codeblock zeigt Listing 8.1 unverändert, mit allen Kommentaren.
- Eine Liste erklärt acht Stellen: die Parameter, den ersten Aufruf, die Bedingung `rechts >= links`, die Rechnung für `mitte`, den Fall „gefunden“, den Fall „links weitersuchen“, den Fall „rechts weitersuchen“ und `return ergebnis`.
- Zu `mitte`: Der Operator `/` teilt ganzzahlig. Mit `links` 0 und `rechts` 9 ist `mitte` gleich 4.
- Zu `mitte-1` und `mitte+1`: Der Name in der Mitte ist schon verglichen und gehört nicht mehr zum Bereich.
- Ein gelber Kasten vergleicht `links + (rechts-links)/2` mit `(links+rechts)/2` und verlinkt die Lektion zum Überlauf.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.1
- 7.4

## Schritt 3: Aufruf für Aufruf

Reiter: aufrufe

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Sechs Knöpfe wählen einen der sechs Fälle aus Abschnitt 8.4.1: Arnold steht als erster Name im Array, Zimmer als letzter, Huber weder vorn noch hinten. Abel fehlt und ist kleiner als alle Namen, Zuse fehlt und ist größer als alle, Maier fehlt und liegt zwischen Lange und Meier.
- Die Demo zeigt ein Programm mit 26 Zeilen: in Zeile 1 bis 18 `binSuche` ohne die Kommentare des Listings, in Zeile 20 bis 26 `main`. `main` legt in Zeile 21 bis 23 das Telefonbuch an, ruft in Zeile 24 `nummer := binSuche(&telbuch, 0, arrlen-1, "…")` mit dem Namen des gewählten Falls auf und gibt in Zeile 25 `nummer` aus.
- Darunter stehen die drei Knöpfe „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“.
- Der Streifen aus Schritt 1 zeigt den Suchbereich des jüngsten Aufrufs, mit den gleichen Farben.
- Darunter stehen die Rahmen: einer für `main` mit `telbuch` und später `nummer`, einer für jeden laufenden Aufruf von `binSuche` mit `personArr`, `links`, `rechts`, `suchwert`, `ergebnis` und, sobald gerechnet, `mitte`. Die Rahmen heißen „binSuche, 1. Aufruf“, „binSuche, 2. Aufruf“ und so weiter. Ein eben beendeter Aufruf hat einen gestrichelten Rahmen.
- Der Hauptspeicher: `telbuch` ab Zelle 100. Jeder Aufruf belegt 24 Zellen für seine Parameter, der erste ab Zelle 400: `personArr` in 400 bis 407 mit dem Inhalt 100, `links` in 408 bis 415, `rechts` in 416 bis 423. Der zweite Aufruf beginnt bei 424. `suchwert`, `ergebnis`, `mitte` und `nummer` stehen nur in den Rahmen. Die Adressen sind erfunden. Wie viele Zellen ein `string` belegt, sagt die Seite nicht.
- Ein Schritt der Demo ist: das Anlegen des Telefonbuchs, je Aufruf der Eintritt mit `var ergebnis int`, die Rechnung von `mitte`, der Vergleich mit der Entscheidung, bei leerem Bereich `ergebnis = -1`, dann je Aufruf die Rückkehr mit dem Ergebnis, die Zuweisung an `nummer` und die Ausgabe.
- Die Aufrufe der sechs Fälle, geprüft mit Go 1.26:
  - Arnold: (0, 9) mitte 4 Koch, (0, 3) mitte 1 Becker, (0, 0) mitte 0 Arnold gefunden. Ausgabe 3310.
  - Zimmer: (0, 9) mitte 4 Koch, (5, 9) mitte 7 Schulz, (8, 9) mitte 8 Wagner, (9, 9) mitte 9 Zimmer gefunden. Ausgabe 2063.
  - Huber: (0, 9) mitte 4 Koch, (0, 3) mitte 1 Becker, (2, 3) mitte 2 Fischer, (3, 3) mitte 3 Huber gefunden. Ausgabe 4456.
  - Abel: (0, 9) mitte 4, (0, 3) mitte 1, (0, 0) mitte 0 Arnold, dann (0, −1) leer. Ausgabe −1.
  - Zuse: (0, 9) mitte 4, (5, 9) mitte 7, (8, 9) mitte 8, (9, 9) mitte 9 Zimmer, dann (10, 9) leer. Ausgabe −1.
  - Maier: (0, 9) mitte 4 Koch, (5, 9) mitte 7 Schulz, (5, 6) mitte 5 Lange, (6, 6) mitte 6 Meier, dann (6, 5) leer. Ausgabe −1.
- Im Kurstext steht zu „kleiner als alle“, dass kein weiterer Aufruf links mehr erfolgt. Das Programm ruft trotzdem noch einmal auf, mit `links` 0 und `rechts` −1. Dieser Aufruf vergleicht nichts und liefert sofort −1.
- Ein Kasten mit einem Satz sagt nach jedem Schritt, was geschieht. Darüber steht die Ausgabe.

### Handgriffe

- Einen der sechs Fälle wählen. Die Demo beginnt dann von vorn.
- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Das Programm und das Telefonbuch lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 8.4.1
- 7.5
- 7.4

## Schritt 4: Aufwand

Reiter: aufwand

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text sagt: Jeder Vergleich mit der Mitte halbiert den Suchbereich mindestens. Ein Array mit 2^n Elementen lässt sich nur n-mal halbieren.
- Eine Tabelle nennt für 10 Einträge höchstens 4 Vergleiche mit der Mitte und höchstens 10 Vergleiche bei der Suche von vorn, für 1000 Einträge 10 und 1000, für 1000000 Einträge 20 und 1000000. Die Zahlen sind mit Go 1.26 gezählt.
- Ein Absatz sagt: Deshalb reichen beim Zahlenraten im Exkurs „Von der Problemstellung zum Go-Programm“ sieben Versuche für die Zahlen von 1 bis 100, denn 2^7 ist 128. Wer immer die Mitte tippt, gewinnt jede Runde. Für 100 Zahlen braucht die Strategie höchstens 7 Tipps, geprüft mit Go 1.26.
- Die Laufzeit der binären Suche ist logarithmisch. Ein Link führt zum Exkurs „Wie die Laufzeit mit der Datenmenge wächst“.
- Unter „Was die Suche voraussetzt“ stehen drei Punkte: das Array ist sortiert, jeder Name steht einmal da, die Suche springt mit `personArr[mitte]` direkt an jeden Index.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: `binSuche` mit einer `for`-Schleife schreiben, warum die Suche nach einer Telefonnummer so nicht geht, was bei einem doppelten Namen geschieht.
- Zur ersten Aufgabe: Die Schleife läuft, solange `rechts >= links` gilt, und setzt statt eines Aufrufs `rechts` auf `mitte-1` oder `links` auf `mitte+1`. Die Lektion „Rekursion in Go“ zeigt in Schritt 1 eine Funktion iterativ und rekursiv.
- Zur zweiten Aufgabe: Das Telefonbuch ist nach `name` sortiert und nicht nach `telefonnummer`. Erst ein Array, das nach der Nummer sortiert ist, lässt sich nach der Nummer durchsuchen. So steht es in Abschnitt 8.4.1 am Beispiel des gedruckten Telefonbuchs.
- Zur dritten Aufgabe: Die Funktion liefert die Nummer des Eintrags, den sie zuerst als Mitte trifft. Das muss nicht der erste der beiden sein.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen. Den Links folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 8.4.1
- 8.4.2
- 9.3.1

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Die Idee, 8.4.1, Seiten 221–222: Das Halbieren des Suchbereichs stammt aus dem Kurstext. Das Telefonbuch ist ein eigenes Beispiel. Die Blicke folgen der Rechnung aus Listing 8.1 und sind mit Go 1.26 geprüft.
  - 2 Die Funktion, 8.4.1, Seite 224: Listing 8.1: Listing 8.1 stammt unverändert aus dem Kurstext. Die Lektion gibt nichts dazu.
  - 3 Aufruf für Aufruf, 8.4.1, Seiten 223–224: Listing 8.1: Die Funktion aus Listing 8.1 und die sechs Fälle stammen aus dem Kurstext. Die Funktion steht hier ohne ihre Kommentare. Aufrufe und Ausgaben sind mit Go 1.26 geprüft. Die Adressen gehören zum Modell der Lektion.
  - 4 Aufwand, 8.4.1, Seiten 222–223: Die Aussage zum Halbieren und die logarithmische Laufzeit stehen im Kurstext. Die Tabelle und die Aufgaben sind ergänzt. Die Zahlen der Tabelle sind mit Go 1.26 gezählt.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Ohne Abschnitt im Kurstext, Halbieren beim Zahlenraten: Lektion: Exkurs: Von der Problemstellung zum Go-Programm, Stufe 3.
  - Abschnitt 7.5, Seite 185, Rekursion: binSuche ruft sich selbst auf. Lektion: Rekursion in Go.
  - Abschnitt 9.3, Seite 237, binäre Bäume: Die Suche im Baum halbiert ebenfalls. Lektion: Binäre Suchbäume in Go.
  - Ohne Abschnitt im Kurstext, Laufzeit: Wie die Zahl der Schritte mit der Datenmenge wächst, zeigt der Exkurs Wie die Laufzeit mit der Datenmenge wächst.
- Unter „Wie es weitergeht“ stehen Links auf die Lektion „Bubblesort auf einem Verbund-Array“ zu Abschnitt 8.4.2, auf die Programmierübung „Ein Telefonbuch als Verbund-Array“ und auf die Lektion „Binäre Suchbäume in Go“ zu Abschnitt 9.3.1.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.5
- 8.4.1
- 9.3
