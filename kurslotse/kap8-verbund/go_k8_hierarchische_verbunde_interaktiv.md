# Hierarchische Verbünde: ein Verbund als Feld

Die Seite zeigt Verbünde, die einen anderen Verbund als Feld enthalten. Das Beispiel stammt aus Abschnitt 8.3 des Kurstexts: Ein Besitzer hat die Felder `name`, `ort` und `fahrzeug`. Für das Feld `fahrzeug` gibt es zwei Bauweisen. Im Typ `besitzerW` hat es den Typ `auto` und enthält die vier Werte des Autos selbst. Im Typ `besitzerZ` hat es den Typ `*auto` und enthält die Adresse eines Autos, das in einer anderen Variable gespeichert ist.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Typ `auto` hat vier Felder: `marke` und `kennzeichen` vom Typ `string`, `tachostand` und `tüvJahr` vom Typ `int`.
- Die Seite führt kein Go-Programm aus. Die Abläufe in Schritt 2 und Schritt 3 sind nachgestellt. Der Student kann den Code nicht ändern, keine eigenen Werte eingeben und das Programm nicht starten.
- Eine Variable ist als Kasten gezeichnet. Im grauen Kopf stehen links der Name und rechts der Typ. Darunter steht je Feld eine Zeile mit dem Namen des Felds links und dem Wert rechts.
- Die Seite zeigt keine Speicherzellen, keine Bitmuster und keine Zahlen als Adressen. An der Stelle einer Adresse steht ein blauer Punkt mit dem Text „Adresse von …“, zum Beispiel „Adresse von audi“.
- Die Seite sagt „Anweisung“ und nummeriert die Anweisungen eines Programms. Eine Anweisung kann über mehrere Zeilen gehen.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Erste Karte: der Begriff hierarchischer Verbund, die zwei Bauweisen des Felds `fahrzeug` als Liste und in einem dunklen Kasten die drei Typdeklarationen `auto`, `besitzerW` und `besitzerZ`. Darunter steht, was die vier Schritte der Seite zeigen.
- Zweite Karte „Worauf die Seite aufbaut“: vier Listenpunkte mit Links auf andere HTML-Seiten. Sie nennen den Datenverbund aus Abschnitt 2.2, die Zuweisung eines Verbunds aus Abschnitt 8.2.3, den Zeiger aus Abschnitt 7.4 und den Zeiger auf einen Verbund aus Abschnitt 8.2.4.

### Handgriffe

- Den Links zu einer anderen HTML-Seite folgen.
- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 8.3

## Schritt 1: Ein Verbund als Feld

Reiter: wert-feld

### Aufbau der Darstellung

- Oben steht in einem dunklen Kasten das Literal für die Variable `mike` vom Typ `besitzerW` mit der Ausgabe von `fmt.Println` als Kommentar.
- Darunter ist `mike` als Kasten gezeichnet. Er hat die Zeilen `name` mit dem Wert "Mike" und `ort` mit dem Wert "München".
- Im Kasten `mike` liegt ein zweiter Kasten mit grünem Rand und grünem Kopf. Das ist das Feld `fahrzeug` vom Typ `auto`. Er hat die vier Zeilen `marke`, `kennzeichen`, `tachostand` und `tüvJahr`.
- Der Kasten im Kasten bedeutet: Die vier Werte des Autos liegen in der Variable `mike` selbst. Einen Pfeil gibt es in diesem Schritt nicht.
- Gelb hinterlegt ist die Zeile oder der Kopf, den der Student gewählt hat.
- Unter dem Kasten steht in einem blauen Feld der Ausdruck zur Auswahl, zum Beispiel `mike.fahrzeug.tachostand`, mit Typ und Wert. Eine Liste darunter geht den Ausdruck Stück für Stück durch: erst die Variable, dann je Punkt ein Feld.

### Handgriffe

- Auf eine der sechs Zeilen klicken: `name`, `ort`, `marke`, `kennzeichen`, `tachostand`, `tüvJahr`.
- Auf den Kopf des äußeren Kastens klicken: Der Ausdruck ist dann `mike`.
- Auf den Kopf des grünen Kastens klicken: Der Ausdruck ist dann `mike.fahrzeug`.
- Die Werte lassen sich nicht ändern.

### Abschnitte des Kurstexts

- 8.3
- 8.2.1

## Schritt 2: Kopie und Zeiger

Reiter: kopie-zeiger

### Aufbau der Darstellung

- Links steht das Programm in einem dunklen Kasten, zehn Anweisungen mit Nummern am linken Rand. Die zuletzt ausgeführte Anweisung ist hell hinterlegt und hat links einen gelben Strich. Ausgeführte Anweisungen sind hellgrau, die übrigen dunkelgrau.
- Rechts stehen die Variablen als Kästen in dieser Reihenfolge: `mike`, `pMikesAuto`, `altesAuto`, `bmw`. Eine Variable ist erst zu sehen, wenn ihre Anweisung ausgeführt ist. Ihr Platz bleibt bis dahin leer.
- `mike` hat den Typ `besitzerW`. Sein Feld `fahrzeug` ist der grüne Kasten im Kasten.
- `altesAuto` und `bmw` haben den Typ `auto` und je vier Zeilen.
- `pMikesAuto` hat den Typ `*auto`. Der Kasten hat einen blauen Rand und nur eine Zeile: einen blauen Punkt und den Text „Adresse von mike.fahrzeug“. Vom Punkt führt ein blauer Pfeil zum grünen Kasten `fahrzeug` in `mike`.
- Der Pfeil bedeutet: Der Zeiger enthält die Adresse dieses Felds. Er führt zum Feld von `mike` und nie zu `bmw` oder `altesAuto`. Nach der Zuweisung `mike.fahrzeug = bmw` zeigt er auf dasselbe Feld wie vorher, in dem dann die Werte des BMW stehen.
- Blau umrandet ist, was die letzte Anweisung gelesen hat. Rot umrandet ist, was sie beschrieben hat. Umrandet ist ein ganzer Kasten oder eine einzelne Zeile.
- Unter den Kästen steht in einem grünen Feld „Was die Anweisung tut“: zwei oder drei nummerierte Sätze zur letzten Anweisung.
- Darunter steht in einem dunklen Feld „Ausgabe des Programms“: je `fmt.Println` eine Zeile. Die Ausgaben bleiben stehen, bis der Student zurückgeht.
- Über dem Programm steht in einem gelben Kasten eine Frage zum Vorhersagen: welche der drei Werte nach Anweisung 6 den BMW zeigen. Die Antwort zeigt Anweisung 7.

### Handgriffe

- „nächste Anweisung ausführen“: Der Ablauf geht eine Anweisung weiter, bis Anweisung 10.
- „eine Anweisung zurück“: Der Ablauf geht eine Anweisung zurück.
- „von vorn“: Der Ablauf beginnt neu. Dann ist keine Variable zu sehen.
- Die Anweisungen laufen nur in der Reihenfolge 1 bis 10. Eine einzelne Anweisung lässt sich nicht überspringen.
- Der Student kann das Programm nicht ändern, keine Werte eingeben und die Kästen nicht anklicken.

### Abschnitte des Kurstexts

- 8.3
- 8.2.3
- 8.2.4

## Schritt 3: Ein Zeiger als Feld

Reiter: zeiger-feld

### Aufbau der Darstellung

- Links steht das Programm in einem dunklen Kasten, sieben Anweisungen mit Nummern am linken Rand. Die zuletzt ausgeführte Anweisung ist hell hinterlegt und hat links einen gelben Strich.
- Rechts stehen drei Kästen nebeneinander: links `olaf`, in der Mitte `audi`, rechts `olga`. Eine Variable ist erst zu sehen, wenn ihre Anweisung ausgeführt ist. Zuerst erscheint `audi` in der Mitte.
- `audi` hat den Typ `auto` und vier Zeilen. Nur in diesem Kasten stehen die Werte des Autos.
- `olaf` und `olga` haben den Typ `besitzerZ` und je drei Zeilen: `name`, `ort` und `fahrzeug`. In der Zeile `fahrzeug` stehen ein blauer Punkt und der Text „Adresse von audi“. Einen Kasten im Kasten gibt es hier nicht.
- Von jedem der zwei Punkte führt ein blauer Pfeil zum Kasten `audi`. Zwei Pfeile auf denselben Kasten bedeuten: Beide Felder enthalten dieselbe Adresse, und beide Besitzer haben dasselbe Auto.
- Blau umrandet ist, was die letzte Anweisung gelesen hat. Rot umrandet ist, was sie beschrieben hat. Bei `olaf.fahrzeug.tachostand += 1000` ist die Zeile `fahrzeug` in `olaf` blau umrandet und die Zeile `tachostand` in `audi` rot umrandet: Gelesen wird die Adresse, beschrieben wird das Auto.
- Unter den Kästen steht in einem grünen Feld „Was die Anweisung tut“, darunter in einem dunklen Feld „Ausgabe des Programms“.
- Über dem Programm steht in einem gelben Kasten eine Frage zum Vorhersagen: welchen Tachostand Anweisung 7 für `audi` ausgibt. Die Antwort zeigt Anweisung 7.

### Handgriffe

- „nächste Anweisung ausführen“: Der Ablauf geht eine Anweisung weiter, bis Anweisung 7.
- „eine Anweisung zurück“: Der Ablauf geht eine Anweisung zurück.
- „von vorn“: Der Ablauf beginnt neu. Dann ist keine Variable zu sehen.
- Die Anweisungen laufen nur in der Reihenfolge 1 bis 7.
- Der Student kann das Programm nicht ändern, keine Werte eingeben und die Kästen nicht anklicken. Ein Zeiger mit dem Wert `nil` und ein dritter Besitzer lassen sich auf der Seite nicht ausprobieren.

### Abschnitte des Kurstexts

- 8.3
- 8.2.4

## Schritt 4: Wert oder Zeiger

Reiter: vergleich

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle und Text. Eine Demo gibt es hier nicht.
- Die Tabelle hat drei Spalten: die Frage, das Feld `fahrzeug auto` im Typ `besitzerW` und das Feld `fahrzeug *auto` im Typ `besitzerZ`.
- Die Tabelle hat fünf Zeilen: was das Feld speichert, was bei einer Zuweisung an das Feld geschieht, ob zwei Besitzer dasselbe Auto haben können, wer eine Änderung am Auto bemerkt, wie der Zugriff auf den Tachostand lautet.
- Die Beispiele in der Tabelle stammen aus Schritt 2 mit `mike` und `bmw` und aus Schritt 3 mit `audi`, `olaf` und `olga`.
- Unter der Tabelle steht die Aussage des Kurstexts, dass in der Praxis meist Zeiger auf Verbünde als Feld-Typ verwendet werden.
- Die Liste „Wie es im Kurstext weitergeht“ nennt Abschnitt 8.3.1 mit den rekursiven Verbund-Typen und Abschnitt 8.3.2 mit den anonymen Feldern. Zu beiden gibt es noch keine HTML-Seite.
- Unter der Karte steht auf jedem Reiter der Link zum Quiz.

### Handgriffe

- Dem Link auf das Go-Programm `S220_hierarchische_verbunde_demo.go` folgen. Es lässt sich auf dem eigenen Rechner mit `go run` ausführen, auf der Seite nicht.
- Mit den Reitern oben zu Schritt 2 oder Schritt 3 wechseln und die Beispiele dort ablaufen lassen.

### Abschnitte des Kurstexts

- 8.3
- 8.3.1
- 8.3.2
