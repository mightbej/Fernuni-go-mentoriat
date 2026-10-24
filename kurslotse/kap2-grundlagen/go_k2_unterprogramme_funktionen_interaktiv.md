# Unterprogramme und Funktionen: aufrufen, übergeben, zurückliefern

Die Seite gehört zu Abschnitt 2.6 des Kurstexts. Sie zeigt, was ein Unterprogramm ist: ein Programmteil mit einem Namen, der einmal dasteht und von mehreren Stellen aus aufgerufen wird. Sie zeigt den Aufruf als Sprung mit Rücksprungadresse, Argument und Parameter, die eigene Variable eines Unterprogramms, die Funktion mit Rückgabewert und die Rekursion.

Die Seite hat sechs Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite kommt ohne Go aus, wie Kapitel 2 des Kurstexts. Die Programmzeilen gehören zu keiner bestimmten Programmiersprache. Der Pfeil `←` bedeutet: Die Variable links bekommt den Wert, der rechts steht. `CALL` ruft ein Unterprogramm auf. `Rücksprung` beendet es.
- Die Abläufe sind nachgestellt. Die Seite führt kein Programm aus. Der Student kann die Programme nicht ändern und nicht starten.
- Das Wort „Schritt“ kommt zweimal vor. Die Reiter heißen Schritt 1 bis Schritt 4. Der Zähler neben den Knöpfen, zum Beispiel „Schritt 3 von 15 des Ablaufs“, zählt die Schritte eines Ablaufs.
- Die Nummern der Anweisungen sind ihre Adressen im Speicher. Das Hauptprogramm beginnt bei 1, das Unterprogramm bei 20.
- Go-Code schreibt der Kurslotse zu dieser Seite nur, wenn der Student ausdrücklich danach fragt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und dem Bild der vier Werke des Rechners. Eine Demo gibt es hier nicht.
- Eine Liste nennt drei Gründe für Unterprogramme: Übersicht, weniger Fehler, Wiederverwendung.
- Im Bild stehen vier Kästen. Oben im gestrichelten Rahmen „Prozessor“ stehen Steuerwerk und Rechenwerk, darunter Speicherwerk und Ein- und Ausgabe. Das Steuerwerk ist gelb hervorgehoben: Der Aufruf eines Unterprogramms ändert, welche Anweisung als nächste an der Reihe ist.
- Den Namen von Neumann nennt der Kurstext nicht. Die Zuordnung zu den Werken ist eine Zutat der Seite.
- Darunter steht, was die vier Schritte zeigen.

### Handgriffe

- Den Links zu einer anderen HTML-Seite folgen.
- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 2.6
- 2.1

## Schritt 1: Aufruf und Rücksprung

Reiter: aufruf

### Aufbau der Darstellung

- Zu sehen sind zwei Programme nebeneinander in dunklen Kästen, jede Anweisung mit ihrer Nummer. Links das Hauptprogramm mit den Anweisungen 1 bis 7, rechts das Unterprogramm `tausche` mit den Anweisungen 20 bis 23.
- Das Hauptprogramm ruft `tausche` zweimal auf, in Anweisung 3 und in Anweisung 5.
- Hell hinterlegt mit einem gelben Strich links ist die zuletzt ausgeführte Anweisung. Springt die Markierung vom linken in den rechten Kasten, ist das der Aufruf. Springt sie zurück, ist das der Rücksprung.
- Unter den Programmen stehen die drei Variablen n, m und k als Kästen mit ihrem Wert. Gelb ist der Kasten, den die letzte Anweisung geändert hat. „kein Wert“ steht in einem gestrichelten Kasten.
- Darunter steht das Feld „Gemerkte Rücksprungadresse“. Es ist leer, solange das Hauptprogramm läuft. Während das Unterprogramm läuft, steht dort 4 beim ersten Aufruf und 6 beim zweiten.
- Darunter steht die Ausgabe des Programms, darunter in einem grünen Feld ein Satz zur letzten Anweisung.
- In diesem Schritt benutzt das Unterprogramm die Variablen n, m und k des Hauptprogramms. Eigene Variablen und Parameter kommen erst in Schritt 2.
- Der Ablauf hat 15 Schritte. Das Unterprogramm läuft zweimal, die Werte sind am Ende wieder 5 und 3.
- Am Ende der Karte stehen zwei Listenpunkte zu zwei anderen Wegen zu vertauschen: mit Minus und Plus und mit XOR. Sie laufen auf der Seite nicht ab. Dort stehen nur die drei Anweisungen und ein Zahlenbeispiel.

### Handgriffe

- „nächste Anweisung ausführen“: Der Ablauf geht einen Schritt weiter.
- „eine Anweisung zurück“: Der Ablauf geht einen Schritt zurück.
- „von vorn“: Der Ablauf beginnt neu.
- Die Werte 5 und 3 lassen sich nicht ändern. Die Möglichkeiten mit Minus und Plus und mit XOR lassen sich auf der Seite nicht ausführen.

### Abschnitte des Kurstexts

- 2.6

## Schritt 2: Parameter und eigene Variable

Reiter: parameter

### Aufbau der Darstellung

- Zu sehen sind wieder zwei Programme nebeneinander. Links das Hauptprogramm mit den Anweisungen 1 bis 9, rechts das Unterprogramm `tausche(v₁, v₂)` mit den Anweisungen 20 bis 24.
- Das Hauptprogramm hat fünf Variablen: k = 100, n = 5, m = 3, a = 7, b = 9. Die Anweisungen 1 bis 5 belegen sie und sind am Anfang des Ablaufs schon ausgeführt. Der Ablauf beginnt mit Anweisung 6.
- Anweisung 6 ruft `tausche(n, m)` auf, Anweisung 7 ruft `tausche(a, b)` auf. n und m, a und b sind die Argumente.
- v₁ und v₂ sind die Parameter. Beim Aufruf `tausche(n, m)` steht v₁ für n und v₂ für m. Eine Anweisung wie `v₁ ← v₂` schreibt dann in die Variable n des Hauptprogramms. So beschreibt es der Kurstext: „Bei dem Aufruf wird dann v₁ zu n und v₂ zu m.“
- Anweisung 20, `Variable k`, deklariert eine eigene Variable k im Unterprogramm. Sie hat einen eigenen Speicherplatz. Die Variable k des Hauptprogramms behält den Wert 100.
- Die Seite benutzt die Wörter „global“ und „lokal“ nicht. Sie sagt „Variable des Hauptprogramms“ und „eigene Variable des Unterprogramms“. Im Unterprogramm meint k immer die eigene Variable.
- Unter den Programmen stehen die fünf Variablen des Hauptprogramms als Kästen. Gelb ist der Kasten, den die letzte Anweisung geändert hat.
- Darunter steht ein blauer Rahmen „Unterprogramm tausche läuft“, solange das Unterprogramm läuft. Darin stehen v₁ mit „steht für n“ und dem Wert von n, v₂ mit „steht für m“ und dem Wert von m, und ab Anweisung 20 die eigene Variable k. Läuft das Unterprogramm nicht, steht dort ein gestrichelter Rahmen mit einem Satz.
- v₁ und v₂ sind keine eigenen Speicherplätze. Der Wert neben v₁ ist der Wert der Variable, für die v₁ steht.
- Darunter stehen die gemerkte Rücksprungadresse, 7 beim ersten Aufruf und 8 beim zweiten, und die Ausgabe. Am Ende lautet sie `3 5 9 7 100`.
- Der Ablauf hat 14 Schritte.
- Unter der Demo steht eine Liste „So geht es in Go weiter“. Dort steht: In Go bekommt ein Parameter eine Kopie des Arguments. Das Unterprogramm dieser Seite würde in Go so nicht vertauschen. Mehr dazu sagt die Seite nicht, sie verweist auf Abschnitt 7.2.

### Handgriffe

- „nächste Anweisung ausführen“: Der Ablauf geht einen Schritt weiter.
- „eine Anweisung zurück“: Der Ablauf geht einen Schritt zurück.
- „von vorn“: Der Ablauf beginnt neu, mit den Anweisungen 1 bis 5 schon ausgeführt.
- Die Werte und die Argumente lassen sich nicht ändern. Ein Aufruf mit anderen Argumenten als n, m und a, b lässt sich auf der Seite nicht ausprobieren.

### Abschnitte des Kurstexts

- 2.6
- 3.4
- 7.2

## Schritt 3: Funktion

Reiter: funktion

### Aufbau der Darstellung

- Zu sehen sind zwei Schreibweisen für dieselbe Rechnung, den Wert 2 · sin(x) + 1.
- Erste Schreibweise, in einem dunklen Kasten: `CALL sinus(x, y)` und danach `z ← 2 · y + 1`. Das Unterprogramm hat zwei Argumente: x ist die Eingabe und bleibt unverändert, y ist die Ausgabe und wird beschrieben.
- Zweite Schreibweise, in einem dunklen Kasten: `z ← 2 · sin(x) + 1`. Hier ist sin eine Funktion, der Aufruf steht im Ausdruck.
- Darunter steht die Auswertung der zweiten Schreibweise, je Schritt eine neue Zeile. Die älteren Zeilen sind grau.
- Zeile 1: der Ausdruck, wie er im Programm steht. Zeile 2: `sin(30)` gelb umrandet, die Funktion ist aufgerufen. Zeile 3: an der Stelle des Aufrufs steht blau umrandet der Rückgabewert. Zeile 4: der ausgerechnete Wert. Zeile 5: `z = …`, die Zuweisung ist ausgeführt.
- Die Seite rechnet x in Grad, damit die Werte glatt sind: sin(0) = 0, sin(30) = 0,5, sin(90) = 1, sin(270) = −1. Im Kurstext hat x keine Einheit.
- Wie das Unterprogramm den Sinus ausrechnet, zeigt die Seite nicht.
- Darunter steht: Go nennt auch ein Unterprogramm ohne Rückgabewert eine Funktion.

### Handgriffe

- Für x einen von vier Werten wählen: 0, 30, 90 oder 270. Die Auswertung beginnt dann von vorn.
- „nächster Schritt der Auswertung“: bis Schritt 4.
- „einen Schritt zurück“ und „von vorn“.
- Andere Werte für x und andere Ausdrücke lassen sich nicht eingeben. Die erste Schreibweise mit `CALL sinus(x, y)` läuft auf der Seite nicht ab.

### Abschnitte des Kurstexts

- 2.6
- 3.4
- 7.1

## Schritt 4: Rekursion

Reiter: rekursion

### Aufbau der Darstellung

- Die Karte hat fünf Teile mit eigener Überschrift: „Die Regel aus dem Kurstext“, „Wie man die Regel liest“, „Die Regel anwenden: Ist 483 eine Dezimaldarstellung?“, „Die Länge folgt derselben Regel“ und „Ausprobieren“.
- Die Regel aus dem Kurstext: `d = 0 | 1 | 2 | … | 9 | 1d | 2d | … | 9d`. Die Seite nennt sie eine Grammatikregel.
- So liest die Seite die Regel: Das d links ist der Name für „Dezimaldarstellung“. Das Zeichen `=` bedeutet „ist aufgebaut als“. Der Strich `|` bedeutet „oder“ und trennt 19 Möglichkeiten. Das d rechts ist ein Platzhalter für irgendeine kürzere Dezimaldarstellung. `4d` bedeutet: die Ziffer 4, direkt dahinter eine Dezimaldarstellung. Eine Multiplikation ist das nicht.
- In Worten: Eine einzelne Ziffer von 0 bis 9 ist eine Dezimaldarstellung. Eine Ziffer von 1 bis 9, auf die eine Dezimaldarstellung folgt, ist wieder eine.
- Eine Tabelle wendet die Regel auf 483 an, je Zeile eine Prüfung: 483 passt zu `4d` mit dem Rest 83, 83 passt zu `8d` mit dem Rest 3, 3 ist eine einzelne Ziffer. Das Gegenbeispiel ist 04: `0d` steht nicht in der Regel.
- Die Tabelle zu 483 ist fester Text. Sie läuft nicht ab und hängt nicht von der eingegebenen Ziffernfolge ab.
- Die Länge l(d) ist 1 für eine einzelne Ziffer und sonst 1 + l(d′), wobei d′ die Ziffernfolge ohne ihre erste Ziffer ist.
- Unter „Ausprobieren“ steht die Demo. Nur sie hängt von der eingegebenen Ziffernfolge ab.
- Zu sehen ist je Aufruf eine Zeile. Jeder weitere Aufruf steht eine Zeile tiefer und weiter rechts. Für d = 48 sind das zwei Zeilen: `l(48) = 1 + l(8)` und darunter `l(8) = 1`.
- Gelb ist die Zeile, um die es im letzten Schritt geht. Grau mit grauem Strich links ist ein Aufruf, der auf den Wert des Aufrufs darunter wartet. Schwarz mit grünem Strich links ist ein Aufruf, dessen Wert feststeht. Dann steht die Rechnung ganz da, zum Beispiel `l(48) = 1 + l(8) = 1 + 1 = 2`.
- Der Ablauf geht erst nach unten: Jeder Schritt startet einen Aufruf, bis die einzelne Ziffer erreicht ist. Dann geht er nach oben: Jeder Schritt liefert einen Wert an den wartenden Aufruf darüber.
- Eine Ziffernfolge mit L Ziffern braucht L Aufrufe und 2 · L − 1 Schritte.
- Unter den Zeilen steht in einem grünen Feld ein Satz zum letzten Schritt.
- Nach der Regel des Kurstexts darf nur die letzte Ziffer eine 0 sein, denn nur auf eine Ziffer von 1 bis 9 folgt eine weitere Dezimaldarstellung. 40 ist erlaubt, 404 und 04 nimmt die Seite nicht an.
- Am Ende steht: In Go wird die Rekursion in Abschnitt 7.5 behandelt, ab Seite 185. Ein Link führt zur Seite „Rekursion in Go“ zu Kapitel 7.

### Handgriffe

- In das Feld „Ziffernfolge d“ 1 bis 6 Ziffern eingeben und „übernehmen“ wählen. Voreingestellt ist 48, das Beispiel des Kurstexts.
- Eines der Beispiele wählen: 48, 7 oder 1613.
- „nächster Schritt“, „einen Schritt zurück“, „von vorn“.
- Eine Eingabe mit einer 0 vor der letzten Stelle, mit mehr als 6 Ziffern oder mit anderen Zeichen nimmt die Seite nicht an. Sie zeigt dann eine Meldung und behält die alte Ziffernfolge.
- Ein Programm für l in Go zeigt die Seite nicht.

### Abschnitte des Kurstexts

- 2.6
- 7.5

## Schritt 5: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.6, Seite 15: Das Unterprogramm und die drei Gründe dafür stammen aus dem Kurstext: Übersicht, weniger Fehler und Wiederverwendung. Die Zuordnung zum Steuerwerk nach John von Neumann gibt die Lektion dazu.
  - 1 Aufruf und Rücksprung, 2.6, Seiten 15–16: Der Aufruf als Sprung, die Rücksprungadresse und das Unterprogramm mit der dritten Variable k stammen aus dem Kurstext. Die zwei anderen Wege zu vertauschen stehen dort auch, mit Fußnote 5 zum Überlauf. Das Hauptprogramm, die Nummern der Anweisungen und die Zahlenbeispiele zu den zwei anderen Wegen sind eigene Beispiele.
  - 2 Parameter und eigene Variable, 2.6, Seite 16: Die eigene Variable k des Unterprogramms, die Parameter v₁ und v₂ und die Argumente stammen aus dem Kurstext. Die Variablen a und b, die Werte und die Nummern der Anweisungen sind eigene Beispiele.
  - 3 Funktion, 2.6, Seiten 16–17: Das Beispiel mit dem Sinus, die zwei Schreibweisen und das Wort Funktionswert stammen aus dem Kurstext. Dort steht auch, dass Go jedes Unterprogramm Funktion nennt. Der Ausdruck 2 · sin(x) + 1 und die Werte von x sind eigene Beispiele. Die Lektion rechnet x in Grad, damit die Sinus-Werte glatt sind.
  - 4 Rekursion, 2.6, Seite 17: Die Regel für die Dezimaldarstellung, die Länge l(d) und das Beispiel 48 stammen aus dem Kurstext. Die Erklärung der Zeichen in der Regel, die Prüfung von 483 und 04 und die Beispiele 7 und 1613 gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 3.1, Seite 21, Pakete: Bibliotheken heißen in Go Pakete. Eine Lektion zu diesem Abschnitt gibt es nicht.
  - Abschnitt 3.4, Seiten 37–41, Funktionen in einem Go-Programm: Parameter, Argument und return an Listing 3.3. Eine Lektion zu diesem Abschnitt gibt es nicht. Die Lektion Fehlersuche: drei Go-Programme mit mehreren Fehlern zeigt Listing 3.3 in Schritt 1.
  - Abschnitt 4.3.7.2.1, Seiten 72–74, XOR: Go schreibt XOR mit dem Operator ^. Lektion: Rechenwerk, Schritt 1.
  - Abschnitt 4.3.8, Seiten 78–84, Überlauf: Warum das Vertauschen mit Minus und Plus bei großen Zahlen scheitert. Lektion: Überlauf bei Ganzzahlen in Go.
  - Abschnitt 5.3.2, Seiten 128–129, Vertauschen in Go: Die Tupel-Zuweisung n, m = m, n. Lektion: Zuweisungen in Go, Schritt 4.
  - Abschnitt 7.1 und 7.2, Seiten 163–170, Funktionen und Übergabe der Argumente: In Go bekommt ein Parameter eine Kopie des Arguments. Lektion: Funktionen in Go.
  - Abschnitt 7.5, Seiten 185–193, Rekursion: Lektion: Rekursion in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 4“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.6
- 3.1
- 3.4
- 4.3.7.2.1
- 4.3.8
- 5.3.2
- 7.1
- 7.2
- 7.5
