# Verkettete Listen in Go

Die Seite erklärt die einfach verkettete Liste: warum das Einfügen in ein sortiertes Array teuer ist, wie ein Element mit zwei Zuweisungen in eine Kette kommt, welche zwei Verbund-Typen die Liste braucht und wie die Funktionen `listeAusgeben`, `längeDerListe` und `elementEinfügen` arbeiten. Sie gehört zu Kapitel 9 des Kurstexts, Abschnitt 9.1 und 9.2, Seiten 227–236.

Die Seite hat acht Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Eine dynamische Datenstruktur ändert sich während der Laufzeit des Programms. Die Liste wächst mit jedem Element.
- Die Typen aus Listing 9.1: `listenElement` mit den Feldern `wert int` und `nächstes *listenElement`, und `sortierteListe` mit dem Feld `kopf *listenElement`.
- Im Bild des Kurstexts ist ein Element ein Wagon, das Feld `nächstes` die Kupplung, die Liste mit `kopf` die Lokomotive. Beim letzten Element ist `nächstes` gleich `nil`. Bei einer leeren Liste ist `kopf` gleich `nil`.
- Eine leere Liste entsteht mit `liste := &sortierteListe{}`. Ein Element entsteht mit `&listenElement{wert: neuerWert}`. Beide sind anonyme Variablen: Verbünde ohne Namen, erreichbar nur über Zeiger.
- Die Liste ist aufsteigend sortiert und darf gleiche Werte mehrfach enthalten.
- Die Seite zeigt eine Liste auf zwei Arten. Die „Kette“: ein violetter Kasten für den Kopf oder den Anfang, dann je Element ein Kasten mit dem Wert und darunter „ab“ und der Anfangsadresse, verbunden mit Pfeilen, am Ende `nil`. Der „Hauptspeicher“: je Variable ein Block mit Zellen, Typ und Inhalt.
- Im Modell der Seite belegen ein `int` und ein Zeiger je 8 Zellen, ein `listenElement` 16 Zellen, eine `sortierteListe` 8 Zellen. Die Adressen sind erfunden. Echte Adressen sehen anders aus, zum Beispiel `0xc000010250`.
- In den Rahmen der Demos steht eine Adresse als „Adresse 216“, im Hauptspeicher als Zahl 216. Ein Block mit violettem Rand ist eine Zeiger-Variable.
- Die Elemente liegen im Hauptspeicher in der Reihenfolge, in der sie erzeugt wurden. Die Reihenfolge der Liste steht allein in den Zeigern.
- Zu einem Element, auf das kein Zeiger mehr zeigt, kommt das Programm nie wieder. Go gibt seinen Speicher später von selbst frei. Das Freigeben erklärt die Seite nicht.
- Doppelt verkettete Listen nennt der Kurstext nur. Die Seite behandelt sie nicht.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code und keine eigenen Werte eingeben. Alle Ausgaben sind mit Go 1.26 geprüft.
- Die Seite verwendet keine Slices. Der Kurs arbeitet mit Arrays.
- Zu den Aufgaben unter „Zum Weiterdenken“ gibt es auf der Seite keine Lösung. Fragt der Student danach, bekommt er einen Hinweis auf den Weg und nicht das fertige Programm.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text nennt die Begriffe dynamische Datenstruktur und einfach verkettete Liste und das Bild des Güterzugs.
- Unter „Was dazugehört“ stehen Links auf die Lektionen „Rekursive Verbund-Typen“, „Zeiger auf Verbünde“, „Zeiger in Go“ und „Arrays in Go“.
- Unter „Die Schritte“ steht je ein Satz zu den sechs Schritten.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9
- 9.2

## Schritt 1: Einfügen im Array

Reiter: array

### Aufbau der Darstellung

- Oben steht ein blauer Kasten „Erst vorhersagen, dann ausprobieren“ mit einer Frage und einer aufklappbaren „Auflösung“.
- Der Text nennt das Beispiel des Kurstexts: einen Schrank mit Schubladen, sortiert nach Versicherungsnummern.
- Die Demo zeigt ein Array mit acht Plätzen. Sechs sind belegt mit 3, 35, 42, 99, 111 und 176, zwei sind frei.
- Vier Knöpfe wählen den Wert, der dazukommt: 200, 100, 40 oder 1.
- Darunter stehen zwei Reihen „vorher“ und „nachher“. In „nachher“ ist der neue Wert gelb, die weitergerückten Elemente sind grün.
- Ein Kasten nennt den Index des neuen Werts und die Zahl der Elemente, die weiterrücken: 0 bei 200, 2 bei 100, 4 bei 40, 6 bei 1.
- Die Elemente rücken von hinten her weiter, das letzte zuerst. Sonst würde ein Element das nächste überschreiben.
- Unter „Was daran stört“ stehen drei Punkte: Aufwand, große Elemente, feste Länge.
- Der Ausweg am Ende von Abschnitt 9.1: Im Array liegen nur Schlüsselwert und Zeiger auf einen Verbund mit den übrigen Daten.

### Handgriffe

- Einen der vier Werte wählen.
- Die „Auflösung“ aufklappen.
- Einen eigenen Wert kann der Student nicht eingeben.

### Abschnitte des Kurstexts

- 9.1

## Schritt 2: Zwei Zuweisungen

Reiter: zug

### Aufbau der Darstellung

- Oben steht der Güterzug des Kurstexts: Wagons mit Containern, sortiert nach den Nummern 3, 35, 42, 99, 111, 176. Der Container 50 soll mit. Der Eisenbahner trennt den Zug hinter dem Wagon 42, hängt den neuen Wagon an und kuppelt den Rest wieder an.
- Eine Liste übersetzt den Zug in das Programm: Wagon ist Element, Containernummer ist Wert, Kupplung ist Zeiger auf das nächste Element, beim letzten Wagon ist der Zeiger `nil`.
- Darunter steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Kette hat die Elemente 35 ab Zelle 200, 42 ab Zelle 216 und 99 ab Zelle 232. `aktuelles` zeigt auf das Element 42. Das neue Element mit dem Wert 50 entsteht ab Zelle 248.
- Es gibt zwei Demos mit je drei Anweisungen und den drei Knöpfen „nächste Anweisung“, „eine Anweisung zurück“ und „von vorn“. Jede Demo zeigt die Kette, Rahmen mit Variablen, den Hauptspeicher und einen Kasten mit einem Satz. Eine Ausgabe haben sie nicht.
- Demo „Die Reihenfolge des Kurstexts“: `neues := &listenElement{wert: 50}`, dann `neues.nächstes = aktuelles.nächstes`, dann `aktuelles.nächstes = neues`. Am Ende lautet die Kette 35, 42, 50, 99.
- Demo „Die umgekehrte Reihenfolge“: erst `aktuelles.nächstes = neues`, dann `neues.nächstes = aktuelles.nächstes`. Nach der zweiten Anweisung führt kein Zeiger mehr zum Element 99. Nach der dritten enthält das neue Element seine eigene Adresse 248 als Nachfolger: Die Kette hat kein Ende, und eine Schleife bis `nil` hört nie auf.
- Ein gestrichelter Kasten ist ein Element, zu dem kein Zeiger mehr führt.
- Ein gelber Kasten nennt die Regel: Erst bekommt das neue Element seinen Nachfolger, dann wird es selbst angehängt.
- Eine Tabelle ordnet die vier Zeiger des Kurstexts zu: z1 ist `aktuelles`, z2 ist `aktuelles.nächstes`, z3 ist `neues`, z4 ist `neues.nächstes`.
- Der Kurstext nennt für das Suchen der Stelle außerdem die Zeiger z5, z6 und z7. Die Seite verwendet sie nicht. In Listing 9.1 übernimmt `aktuelles` die Rolle von z5 und `aktuelles.nächstes` die von z6.

### Handgriffe

- In jeder der zwei Demos „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Die Werte und die Anweisungen lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 9.2

## Schritt 3: Der Typ der Liste

Reiter: typ

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt die Typen `listenElement` und `sortierteListe`, ein zweiter die Zeile `liste := &sortierteListe{}`.
- Ein festes Bild zeigt die Liste, nachdem die Werte 8, 17, 12 und 2 eingefügt sind. Die Kette lautet 2, 8, 12, 17.
- Die Blöcke im Hauptspeicher: `liste` in den Zellen 100 bis 107 mit dem Inhalt 120. Ab Zelle 120 die `sortierteListe` mit `kopf` 248. Ab 200 das Element 8 mit `nächstes` 232. Ab 216 das Element 17 mit `nil`. Ab 232 das Element 12 mit `nächstes` 216. Ab 248 das Element 2 mit `nächstes` 200.
- Eine Liste nennt Größe, Reihenfolge im Speicher und Reihenfolge der Liste.
- Eine Demo gibt es hier nicht.

### Handgriffe

- Die „Auflösung“ aufklappen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.2
- 8.3.1

## Schritt 4: Durchlaufen

Reiter: durchlaufen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Die Demo zeigt die neun Zeilen der Funktion `listeAusgeben` aus Listing 9.1 und arbeitet auf der Liste 2, 8, 12, 17 aus Schritt 3.
- Darunter stehen die drei Knöpfe, die Kette, Rahmen für `main` und `listeAusgeben`, ein Rahmen „Anonyme Variablen, erzeugt mit &“, der Hauptspeicher, die Ausgabe und ein Kasten mit einem Satz.
- Der Parameter `liste` ist eine Kopie des Zeigers aus `main`, im Modell in den Zellen 300 bis 307. `element` liegt in den Zellen 308 bis 315 und wandert: 248, 200, 232, 216, `nil`.
- In der Kette ist das Element grün umrandet, auf das `element` zeigt.
- Die Demo hat 12 Schritte. Die Ausgabe am Ende ist `-> 2 -> 8 -> 12 -> 17 -> nil`.
- Ein Codeblock zeigt `längeDerListe`. Die Länge steht in keinem Feld, das Programm zählt alle Elemente.
- Eine Liste hat keinen Index. Zum dritten Element führt nur der Weg über die zwei davor.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Die Liste lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 9.2

## Schritt 5: Einfügen

Reiter: einfuegen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“.
- Ein Codeblock zeigt `elementEinfügen` aus Listing 9.1 ohne die Zeilen mit `fmt.Printf`. Die Demo zeigt die Meldungen als Ausgabe.
- Die Funktion unterscheidet zwei Fälle. Erster Fall: Die Liste ist leer, oder der neue Wert ist kleiner oder gleich dem Wert im Kopf. Dann wird das neue Element der Kopf. Zweiter Fall: `aktuelles` beginnt beim Kopf und rückt weiter, solange es einen Nachfolger gibt und dessen Wert kleiner ist als der neue Wert. Danach folgen die zwei Zuweisungen aus Schritt 2.
- Ein gelber Kasten erklärt die Kurzschluss-Auswertung in beiden Bedingungen: `liste.kopf.wert` wird nur gelesen, wenn `kopf` nicht `nil` ist. `aktuelles.nächstes.wert` wird nur gelesen, wenn `aktuelles.nächstes` nicht `nil` ist.
- Die Demo hat elf Zeilen aus `main`: `liste := &sortierteListe{}`, neun Aufrufe von `elementEinfügen` mit den Werten 8, 17, 12, 2, 23, 20, 17, 5, 14 und die Ausgabe der Länge. Ein Schritt der Demo ist ein ganzer Aufruf.
- Das neue Element ist gelb. Grün umrandet ist das Element, bei dem `aktuelles` stehen geblieben ist.
- Die Elemente entstehen ab den Zellen 200, 216, 232, 248, 264, 280, 296, 312 und 328, in der Reihenfolge der Aufrufe.
- Die Meldungen: „Wert 8 in leere Liste einfügen“, „Wert 17 am Ende hinter 8 einfügen“, „Wert 12 zwischen 8 und 17 einfügen“, „Wert 2 am Anfang vor 8 einfügen“, „Wert 23 am Ende hinter 17 einfügen“, „Wert 20 zwischen 17 und 23 einfügen“, „Wert 17 zwischen 12 und 17 einfügen“, „Wert 5 zwischen 2 und 8 einfügen“, „Wert 14 zwischen 12 und 17 einfügen“, am Ende „Länge der Liste: 9“.
- Im Kurstext steht bei zwei Meldungen ein anderer Wortlaut, „ans Ende hinter“ und „am Anfang einfügen vor“. Die Seite zeigt, was das Programm mit Go 1.26 ausgibt.
- Die Liste am Ende: 2, 5, 8, 12, 14, 17, 17, 20, 23.
- Die Schleife rückt bei 23, 20 und 14 dreimal weiter, beim zweiten Wert 17 zweimal, sonst nicht.
- Das zweite Element 17 steht vor dem ersten, weil die Schleife nur weiterrückt, solange der Nachfolger kleiner ist.

### Handgriffe

- „nächste Anweisung“, „eine Anweisung zurück“ oder „von vorn“ wählen.
- Die „Auflösung“ aufklappen.
- Die Werte und ihre Reihenfolge lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 9.2
- 4.5.2

## Schritt 6: Liste und Array

Reiter: vergleich

### Aufbau der Darstellung

- Zu sehen ist eine Tabelle mit vier Zeilen: Einfügen, Finden der Stelle, Zahl der Elemente, Speicher je Element, jeweils für das sortierte Array und für die sortierte verkettete Liste.
- Darunter steht: Das Einfügen ist in der Liste billig, das Suchen der Stelle nicht. Deshalb folgt in Abschnitt 9.3 der binäre Suchbaum.
- Unter „Zum Weiterdenken“ stehen drei Aufgaben ohne Lösung: eine Funktion `istWertInListe` schreiben, ein Element löschen, die Liste umdrehen. Der Kurstext nennt Suchen, Löschen und Umdrehen und führt sie nicht aus.
- Zur ersten Aufgabe: In einer sortierten Liste darf die Suche aufhören, sobald ein Element größer ist als der Suchwert.
- Zur zweiten Aufgabe: Das Feld `nächstes` des Vorgängers bekommt die Adresse des Nachfolgers. Das erste Element hat keinen Vorgänger, dort bekommt `kopf` den neuen Wert.
- Eine Demo gibt es hier nicht.
- Am Ende des Schritts unter „Wie es weitergeht“ stehen Links auf die Lektion „Binäre Suchbäume in Go“, auf den Exkurs „Blockchain: eine verkettete Liste mit Fingerabdruck“ und auf den Exkurs zur Laufzeit. Der Exkurs zur Blockchain baut eine Liste, in der jeder neue Block auf seinen Vorgänger zeigt und dessen Hashwert trägt.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.2
- 9.3

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Einfügen im Array, 9.1, Seiten 227–229: Die Containernummern 3 bis 176 stammen aus Abbildung 9.2. Das Array mit acht Plätzen ist ein eigenes Beispiel, geprüft mit Go 1.26.
  - 2 Zwei Zuweisungen, 9.2, Seiten 229–231: Abbildung 9.3: Der Güterzug und die drei Zeilen aus der Funktion elementEinfügen in Listing 9.1 stammen aus dem Kurstext. Die umgekehrte Reihenfolge ist ein eigenes Beispiel, geprüft mit Go 1.26.
  - 3 Der Typ der Liste, 9.2, Seiten 231–233: Die zwei Typen und die leere Liste stammen aus Listing 9.1. Die Adressen und die Größen in Zellen gehören zum Modell der Lektion.
  - 4 Durchlaufen, 9.2, Seiten 233–234: Listing 9.1: Beide Funktionen stammen aus Listing 9.1. Die Ausführung Schritt für Schritt gibt die Lektion dazu.
  - 5 Einfügen, 9.2, Seiten 233–236: Listing 9.1: Die Funktion elementEinfügen stammt aus Listing 9.1. Die Lektion kürzt die Funktion um die Meldungen. Die Meldungen in der Ausgabe stammen von Go 1.26. Im Kurstext steht an zwei Stellen „ans Ende hinter“ und „am Anfang einfügen vor“.
  - 6 Liste und Array, 9.2, Seiten 231–232: Die Tabelle fasst zusammen, was im Kurstext steht. Die Zeile zum Speicher je Element ist eine Ergänzung.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speicherplatz: Bei einfachen Typen ist die Zahl der Speicherelemente fest, bei zusammengesetzten Typen nicht immer. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Kapitel 6, Seiten 151–162, Arrays: Ein Array hat eine feste Länge. Lektion: Arrays in Go.
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.
  - Abschnitt 8.3.1, Seiten 215–216, rekursive Verbund-Typen: Ein Verbund hat ein Feld mit einem Zeiger auf den eigenen Typ. Lektion: Rekursive Verbund-Typen in Go, Schritt 4.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 6“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.
- Den Links folgen.

### Abschnitte des Kurstexts

- 2.2
- 7.3
- 7.4
- 8.3.1
- 9.1
- 9.2
