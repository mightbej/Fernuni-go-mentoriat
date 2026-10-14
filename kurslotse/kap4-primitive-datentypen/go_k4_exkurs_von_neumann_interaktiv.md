# Exkurs: Speichern und Rechnen im Rechner nach von Neumann

Die Seite gibt die Orientierung für Kapitel 4 des Kurstexts. Sie zeigt einen Rechner aus vier Werken: Steuerwerk, Rechenwerk, Speicherwerk, Ein- und Ausgabe. Sie ordnet die zwei Teile eines Datentyps je einem Werk zu: Der Wertebereich gehört zum Speicherwerk, die Operationen gehören zum Rechenwerk.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Aufteilung in vier Werke geht auf John von Neumann zurück, er hat sie 1945 beschrieben. Der Kurstext nennt den Namen von Neumann nicht. Die Seite ist das Bild zu den Abschnitten 2.2 und 2.3 des Kurstexts.
- Das Bild des Rechners ist in Schritt 1 und Schritt 2 gleich aufgebaut. Oben stehen Steuerwerk links und Rechenwerk rechts in einem gestrichelten Rahmen mit der Aufschrift „Prozessor“. Darunter liegt ein grauer Balken über die ganze Breite, er steht für die Leitungen. Unter dem Balken stehen Speicherwerk links und „Ein- und Ausgabe“ rechts.
- Die Seite ist ein vereinfachtes Modell. Sie zeigt nur die vier Werke und die Leitungen. Adressen und Speicherzellen zeigt erst die Seite „Speichermodell“.
- Die Seite verwendet nur die Wörter Steuerwerk, Rechenwerk, Speicherwerk, „Ein- und Ausgabe“, Leitungen und Prozessor. Register, Befehlszähler, Bus und Akkumulator kommen auf der Seite nicht vor.
- Die Seite führt kein Programm aus. Der Student kann die zwei Anweisungen nicht ändern und nicht starten.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text und einer Tabelle. Eine Demo und ein Bild des Rechners gibt es hier nicht.
- Eine Liste nennt zwei Abschnitte aus Kapitel 2 des Kurstexts: Abschnitt 2.2 zum Speichern, Abschnitt 2.3 zum Rechnen.
- Die Tabelle hat drei Spalten: den Teil des Datentyps, ein Beispiel für `int8` und das Werk des Rechners, das dafür zuständig ist.
- Erste Zeile der Tabelle: Der Wertebereich legt fest, welche Werte der Typ kennt. `int8` kennt die ganzen Zahlen von −128 bis 127. Das Speicherwerk hält den Wert als Bitmuster.
- Zweite Zeile der Tabelle: Die Operationen legen fest, was man mit den Werten tun darf, zum Beispiel addieren, vergleichen oder den Rest einer Division bilden. Das Rechenwerk führt die Operation aus.
- Unter der Tabelle steht, was die zwei Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Anklicken und einstellen lässt sich in diesem Schritt sonst nichts.

### Abschnitte des Kurstexts

- 2.2
- 2.3
- 4.1

## Schritt 1: Eine Anweisung läuft durch den Rechner

Reiter: ablauf

### Aufbau der Darstellung

- Zu sehen ist oben ein dunkler Kasten mit einem Programm aus zwei Anweisungen: `a = a + 4` und `fmt.Println(a)`. Die Variable `a` hat den Typ `int8` und am Anfang den Wert 17.
- Hell hinterlegt mit einem gelben Strich am linken Rand ist die Anweisung, die gerade an der Reihe ist. Am Anfang ist keine markiert.
- Darunter stehen drei Knöpfe, darunter das Bild des Rechners mit den vier Werken.
- Gelb unterlegt ist jedes Werk, das im sichtbaren Stand arbeitet. Das sind ein Werk oder zwei Werke.
- Jedes Werk zeigt unter seinem Namen seinen Inhalt. Das Steuerwerk zeigt die Anweisung, die an der Reihe ist. Das Rechenwerk zeigt, was es gerade rechnet. Das Speicherwerk zeigt das Bitmuster von `a`. „Ein- und Ausgabe“ zeigt, was auf dem Bildschirm steht.
- Der Balken für die Leitungen ist gelb, wenn ein Bitmuster von einem Werk zum anderen geht. Im Balken steht dann, was unterwegs ist und wohin es geht. Sonst ist der Balken grau.
- Unter dem Bild erklärt ein grüner Kasten den sichtbaren Stand in ein bis zwei Sätzen.
- Der Ablauf hat einen Anfang und sieben Schritte:
  - Anfang: Kein Werk arbeitet. Das Speicherwerk hält für `a` das Bitmuster 00010001. Als `int8` gelesen ist das 17.
  - Schritt 1 des Ablaufs: Das Steuerwerk holt die Anweisung `a = a + 4` aus dem Speicherwerk. Steuerwerk und Speicherwerk sind gelb.
  - Schritt 2 des Ablaufs: Das Steuerwerk stellt fest, was die Anweisung verlangt. Im Steuerwerk stehen unter der Anweisung drei Zeilen: „hole a“, „addiere 4“, „speichere a“. Nur das Steuerwerk ist gelb.
  - Schritt 3 des Ablaufs: Das Speicherwerk liefert das Bitmuster 00010001 an das Rechenwerk. Im Speicherwerk bleibt der Wert stehen, das Rechenwerk bekommt eine Kopie. Speicherwerk und Rechenwerk sind gelb.
  - Schritt 4 des Ablaufs: Das Rechenwerk addiert. Im Rechenwerk stehen drei Zeilen untereinander: 00010001, darunter + 00000100, darunter = 00010101. Als `int8` gelesen heißt das 17 + 4 = 21. Nur das Rechenwerk ist gelb.
  - Schritt 5 des Ablaufs: Das Bitmuster 00010101 geht vom Rechenwerk zum Speicherwerk und überschreibt dort den alten Wert. `a` hat jetzt den Wert 21. Rechenwerk und Speicherwerk sind gelb.
  - Schritt 6 des Ablaufs: Das Steuerwerk holt die Anweisung `fmt.Println(a)` aus dem Speicherwerk. Steuerwerk und Speicherwerk sind gelb.
  - Schritt 7 des Ablaufs: Das Bitmuster 00010101 geht vom Speicherwerk zur Ausgabe. Es wird als `int8` gelesen und in die Ziffern 2 und 1 umgewandelt. Auf dem Bildschirm steht 21. Speicherwerk und „Ein- und Ausgabe“ sind gelb.
- In Schritt 3 des Ablaufs wird das Speicherwerk nur gelesen. Lesen verändert den gespeicherten Wert nicht. Erst Schritt 5 des Ablaufs überschreibt ihn.
- Die drei Zeilen „hole a“, „addiere 4“, „speichere a“ sind kein Go und keine Befehle einer bestimmten Maschinensprache. Sie sagen in Worten, was die Anweisung verlangt.
- Auch die Anweisungen liegen als Bitmuster im Speicherwerk. Das Bild zeigt im Speicherwerk nur das Bitmuster von `a`.
- Unter dem grünen Kasten steht ein Hinweis: Der Compiler übersetzt `a = a + 4` in mehrere Befehle der Maschinensprache, nämlich laden, addieren und zurückschreiben. Das Modell fasst sie zu einer Anweisung zusammen.

### Handgriffe

- „nächster Schritt“: Der Ablauf geht einen Schritt weiter. Nach Schritt 7 des Ablaufs ist der Knopf gesperrt.
- „einen Schritt zurück“: Der Ablauf geht einen Schritt zurück. Am Anfang ist der Knopf gesperrt.
- „von vorn“: Der Ablauf beginnt neu, `a` hat wieder den Wert 17.
- Die vier Werke lassen sich in diesem Schritt nicht anklicken. Das geht in Schritt 2 der Seite.
- Die zwei Anweisungen, der Typ `int8`, der Anfangswert 17 und die Zahl 4 lassen sich nicht ändern. Einen Überlauf oder eine andere Operation kann der Student hier nicht ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 2.3
- 2.7

## Schritt 2: Die vier Werke

Reiter: werke

### Aufbau der Darstellung

- Zu sehen ist dasselbe Bild des Rechners wie in Schritt 1. Hier ist jedes der vier Werke ein Knopf, und unter jedem Namen steht in einem Satz die Aufgabe des Werks.
- Im Balken für die Leitungen steht fest: Über die Leitungen wandern Bitmuster von einem Werk zum anderen. Der Balken lässt sich nicht anklicken.
- Das gewählte Werk ist gelb unterlegt. Es ist immer nur ein Werk gewählt. Am Anfang ist keines gewählt.
- Unter dem Bild steht in einem grünen Kasten die Auskunft zum gewählten Werk: was das Werk tut und wo es im Kurstext vorkommt.
- Auskunft zum Steuerwerk: Es holt Anweisung für Anweisung aus dem Speicherwerk und beauftragt die anderen Werke. Welche Anweisung als nächste an der Reihe ist, regeln in Go `if`, `switch` und `for`. Das ist Kapitel 5 des Kurstexts.
- Auskunft zum Rechenwerk: Es bekommt Bitmuster und liefert ein neues Bitmuster. Jede Operation des Rechenwerks hat in Go einen Operator, zum Beispiel `+`, `<` oder `&&`. Die Operatoren stehen in Abschnitt 4.1. Ein Link führt zur Seite „Rechenwerk“.
- Auskunft zum Speicherwerk: Es hält die Werte der Variablen und das Programm als Bitmuster. Der Typ einer Variable legt fest, wie viele Bits sie braucht und welche Werte hineinpassen. Das steht für die Ganzzahlen in Abschnitt 4.3.4. Ein Link führt zur Seite „Speichermodell“.
- Auskunft zu „Ein- und Ausgabe“: Über die Eingabe kommen Daten in den Rechner, über die Ausgabe verlassen sie ihn. Die Funktion `fmt.Printf` gibt Werte aus, das steht in Abschnitt 4.2.
- Unter dem grünen Kasten steht ein Hinweis: Auch das Programm liegt als Bitmuster im Speicherwerk. Einem Bitmuster sieht man nicht an, ob es eine Zahl, ein Zeichen oder eine Anweisung bedeutet.

### Handgriffe

- Eines der vier Werke anklicken: Steuerwerk, Rechenwerk, Speicherwerk oder „Ein- und Ausgabe“.
- Den Links zu den Seiten „Speichermodell“ und „Rechenwerk“ folgen. Der Link zur Seite „Rechenwerk“ erscheint erst, wenn das Rechenwerk gewählt ist.
- Ein Programm läuft in diesem Schritt nicht. Den Ablauf zeigt Schritt 1 der Seite.

### Abschnitte des Kurstexts

- 4.1
- 4.2
- 4.3.4
- 4.3.7.2

## Schritt 3: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.2, Seiten 8–10: Speichern; 2.3, Seiten 10–11: Rechnen; 4.1, Seite 45: Datentyp: Speichern und Rechnen stammen aus Kapitel 2. Dass ein Datentyp aus Wertebereich und Operationen besteht, steht in Abschnitt 4.1. Die Zuordnung zu den Werken des Rechners gibt die Lektion dazu.
  - 1 Eine Anweisung läuft durch den Rechner, 2.3, Seiten 10–11: Ausdruck und Zuweisung; 2.7, Seiten 17–19: Maschinensprache: Dass der Rechner Bitmuster in andere Bitmuster umwandelt, steht in Abschnitt 2.3. Der Rechner nach von Neumann mit vier Werken und der Ablauf von a = a + 4 stehen nicht im Kurstext.
  - 2 Die vier Werke, 4.3.4, Seite 63: Ganzzahl-Typen; 4.1, Seiten 45–52: Operatoren; 4.2, Seiten 53–58: fmt.Printf; Kapitel 5: Anweisungen: Zu jedem Werk nennt die Lektion den Abschnitt, in dem sein Teil der Arbeit in Go vorkommt. Die vier Werke selbst stehen nicht im Kurstext.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speichern: Eine Variable belegt Speicherplatz aus Speicherelementen mit 0 oder 1. Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 2.7, Seiten 17–19, Maschinensprache: Lektion: Höhere und niedrigere Programmiersprachen.
- Unter „Wie es weitergeht“ steht ein Link zur Seite „Speichermodell“.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 2“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 2.3
- 2.7
- 4.1
- 4.2
- 4.3.4
