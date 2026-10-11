# Gatter aus Schaltern: Das Rechenwerk zum Anfassen

Die Seite baut die Gatter AND, OR, XOR und NOT aus Schaltern, einer Batterie und einer Lampe. Sie führt von dort zu einer Stelle des Addierers, zum Relais und zum Prozessor. Sie gehört zur Seite „Rechenwerk“, die acht solche Stellen zu einem Addierer zusammensetzt.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

In allen Schaltbildern gilt:

- Ein Klick auf einen Schalter legt ihn um.
- Schalter mit demselben Namen gehören zusammen und schalten gemeinsam. Die Schalter a und b gelten in Schritt 1 und in Schritt 2 gemeinsam.
- Eine dicke, farbige Leitung führt Strom. Eine dünne, graue Leitung führt keinen Strom.
- Ein Punkt verbindet Leitungen. Leitungen, die sich ohne Punkt kreuzen, berühren sich nicht.
- Links in jedem Schaltbild steht die Batterie, rechts die Lampe. Eine leuchtende Lampe ist gelb.
- Ein Schalter auf 1 ist geschlossen und lässt Strom durch. Ein Schalter auf 0 ist offen.

Im Kurstext werden AND, OR, XOR und NOT in Abschnitt 4.3.7.2.1 auf den Werten 0 und 1 definiert, mit den Operatoren `&`, `|` und `^`. Die Gatter aus Schaltern, das Relais und der Prozessor kommen im Kurstext nicht vor. Die Seite enthält keinen Go-Code außer den Ausdrücken `a & b`, `a | b`, `a ^ b`, `^a` und `a ^ b ^ c`.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Schaltbilder gibt es hier nicht.
- Über der Tabelle wird die Schreibweise von Go eingeführt: `a & b` für AND, `a | b` für OR, `a ^ b` für XOR, `^a` für NOT, `a ^ b ^ c` für zweimal XOR hintereinander. Dieselben Ausdrücke stehen in den nächsten Schritten unter den Schaltbildern, dahinter ihr Wert.
- Die Tabelle „Gatter und ihre Verdrahtung“ nennt zu jedem Gatter die Schreibweise in Go, die Verdrahtung und ein Beispiel aus dem Alltag.
- Unter der Tabelle steht, was davon im Kurstext vorkommt.

### Handgriffe

- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.1

## Schritt 1: Vier Gatter aus zwei Schaltern

Reiter: gatter

### Aufbau der Darstellung

- Oben steht eine Frage zum Vorhersagen: das Licht im Treppenhaus. Die Auflösung ist eingeklappt.
- Darunter stehen vier Schaltbilder in zwei Reihen: oben links AND, oben rechts OR, unten links XOR, unten rechts NOT.
- Alle vier Schaltbilder benutzen dieselben Schalter a und b. Wer a in einem Schaltbild umlegt, legt a in allen um.
- AND: Die Schalter a und b liegen hintereinander in einer Leitung. Der Strom muss durch beide.
- OR: Die Schalter a und b liegen nebeneinander in zwei Leitungen. Der Strom hat zwei Wege, einer genügt.
- XOR: Zwei Wechselschalter mit zwei Leitungen dazwischen, die über Kreuz laufen. Ein Wechselschalter hat einen Eingang und zwei Ausgänge. Seine Stellung oben steht für 1. Strom fließt, wenn die beiden Schalter verschieden stehen.
- NOT: Nur der Schalter a. Er ist ein Öffner: Bei a = 1 unterbricht er die Leitung, bei a = 0 ist sie geschlossen.
- Unter jedem Schaltbild steht ein Satz, ob die Lampe leuchtet, und der Wert des Ausdrucks.

### Handgriffe

- Schalter a anklicken und damit umlegen.
- Schalter b anklicken und damit umlegen.
- Die Auflösung der Frage aufklappen.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.1

## Schritt 2: Eine Stelle des Addierers aus Schaltern

Reiter: addierer

### Aufbau der Darstellung

- Oben stehen zwei Fragen zum Vorhersagen. Die Auflösungen sind eingeklappt.
- Eine Stelle des Addierers zählt drei Bits zusammen: a, b und c. c ist der Übertrag von rechts. Das Ergebnis hat zwei Ziffern im Zweiersystem, und für jede Ziffer gibt es eine Lampe.
- Darunter stehen zwei Schaltbilder nebeneinander in einem gemeinsamen grünen Rahmen. Anders als in Schritt 1 gehören sie zusammen: Sie sind die zwei Ausgänge derselben Stelle. Dieselben Schalter a, b und c steuern beide Lampen.
- Was eine leuchtende Lampe bedeutet: Die Lampe „Summe“ ist das Bit dieser Stelle. Leuchtet sie, steht im Ergebnis an dieser Stelle eine 1. Die Lampe „Übertrag“ ist die 1, die an die nächste Stelle links weitergeht.
- Linkes Schaltbild, Lampe „Übertrag nach links“, die linke Ziffer: zwei Wege. Auf dem oberen Weg liegen a und b hintereinander. Auf dem unteren Weg liegt c, dahinter a und b nebeneinander. Die Lampe leuchtet, wenn mindestens zwei der drei Schalter geschlossen sind.
- Rechtes Schaltbild, Lampe „Summe“, die rechte Ziffer: die Kreuzschaltung. Links der Wechselschalter a, in der Mitte der Kreuzschalter b in einem gestrichelten Kasten, rechts der Wechselschalter c. Jeder der drei Schalter wechselt das Licht, wenn man ihn umlegt. Die Lampe zeigt `a ^ b ^ c`.
- Im linken Schaltbild kommen a und b zweimal vor. Beide Stellen schalten gemeinsam.
- Unter den Schaltbildern, noch im Rahmen, steht die Rechnung, zum Beispiel 1 + 0 + 1 = 10 im Zweiersystem, das ist 2. Die linke Ziffer ist die Lampe „Übertrag“, die rechte Ziffer die Lampe „Summe“. Die Lampen stehen also in derselben Reihenfolge wie die Ziffern.

### Handgriffe

- Schalter a, b oder c anklicken und damit umlegen.
- Die Auflösungen der Fragen aufklappen.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 4.3.7.2.1

## Schritt 3: Das Relais, Strom schaltet Strom

Reiter: relais

### Aufbau der Darstellung

- Oben steht eine Frage zum Überlegen. Die Auflösung ist eingeklappt.
- Darunter steht ein Schaltbild mit zwei Stromkreisen nebeneinander. Jeder hat eine eigene Batterie.
- Linker Stromkreis, beschriftet mit „Steuerkreis“: ein Schalter und ein Kasten mit der Aufschrift „Spule“.
- Rechter Stromkreis, beschriftet mit „zweiter Stromkreis“: ein Kontakt und eine Lampe.
- Eine gestrichelte Linie führt von der Spule zum Kontakt. Sie zeigt, dass die Spule den Kontakt bewegt. Durch die gestrichelte Linie fließt kein Strom.
- Anklicken lässt sich nur der Schalter im Steuerkreis. Den Kontakt bewegt die Spule.

### Handgriffe

- Den Schalter im Steuerkreis anklicken und damit umlegen.
- Die Auflösung der Frage aufklappen.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

Zu diesem Schritt gibt es im Kurstext keinen Abschnitt.

## Schritt 4: Vom Relais zum Prozessor

Reiter: prozessor

### Aufbau der Darstellung

- Zu sehen sind zwei Karten.
- Erste Karte: zwei Diagramme nebeneinander. Links der Takt in Schritten je Sekunde, rechts die Zahl der Schalter in einem Rechner, beide von 1941 bis 2022. Die Achse nach oben ist logarithmisch: Jeder Strich bedeutet das Tausendfache des Strichs darunter.
- Unter den Diagrammen steht eine Tabelle mit sieben Rechnern und Prozessoren: Jahr, Name, Takt, Zahl der Schalter.
- Zweite Karte, „Die Baumarkt-Rechnung“: fünf Eingabefelder und darunter das Ergebnis in vier Punkten: Preis im Prozessor, Preis im Baumarkt, von Hand geschaltet, Platz an der Wand.
- Die Diagramme lassen sich nicht verändern. Die Rechnung ändert sich mit jeder Eingabe.

### Handgriffe

- In der Baumarkt-Rechnung eine der fünf Zahlen ändern: Schalter im Prozessor in Milliarden, Preis des Prozessors in Euro, Preis eines Schalters im Baumarkt in Euro, Takt in Gigahertz, Schalter, die ein Mensch pro Sekunde umlegt. Jede Zahl muss größer als 0 sein.
- Mit den Reitern oben zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

Zu diesem Schritt gibt es im Kurstext keinen Abschnitt.
