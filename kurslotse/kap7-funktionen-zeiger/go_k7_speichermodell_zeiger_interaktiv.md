# Speichermodell: Zeiger in Go

Die Seite zeigt, was ein Zeiger im Hauptspeicher ist: eine Variable, in deren Zellen die Adresse einer anderen Variable steht. Sie zeigt die drei Schreibweisen `&b`, `*p` und `new(int32)` und führt ein Programm Zeile für Zeile aus. Die Seite gehört zu Kapitel 7 des Kurstexts, zu Abschnitt 7.4 und zu Abschnitt 7.4.5. Sie ist die dritte von vier HTML-Seiten mit dem gleichen Modell.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Hauptspeicher ist eine Reihe von Zellen. Jede Zelle fasst 8 Bit, also 1 Byte, und hat eine Adresse.
- Neben dem Hauptspeicher zeigt die Seite eine Namenstabelle im Modell. In jeder Zeile stehen ein Name, ein Typ und eine Adresse. Der Wert steht nur im Hauptspeicher.
- Die Seite sagt dazu wörtlich: „Diese Tabelle gibt es so nicht. Sie ist ein Modell dafür, wie der Compiler einem Namen einen Typ und eine Adresse zuordnet. Das laufende Programm kennt nur noch Adressen.“ Der Rechner führt keine solche Tabelle und schlägt beim Lauf keinen Namen nach.
- Ein Zeiger ist eine Variable. Sein Wert ist die Adresse einer anderen Variable. Im Modell braucht jeder Zeiger 8 Zellen, gleich auf welchen Typ er zeigt, weil auf einem 64-Bit-Rechner jede Adresse 8 Zellen lang ist.
- Der Typ `*int32` sagt zweierlei: In den Zellen des Zeigers steht eine Adresse, und an dieser Adresse liegt ein `int32` mit 4 Zellen.
- `&b` liefert die Adresse von `b`. Im Kurstext heißt `&` Adressoperator.
- `*p` geht zu den Zellen an der Adresse, die in `p` steht. Im Kurstext heißt `*` vor einem Zeiger Dereferenzierungsoperator, und `*p` ist die Variable, auf die `p` zeigt.
- `new(int32)` reserviert 4 Zellen ohne Namen, füllt sie mit dem Null-Wert und liefert ihre Adresse. Im Kurstext heißt eine solche Variable anonyme Variable.
- Das Modell hat 32 Zellen mit den Adressen 100 bis 131. Die Adressen sind erfunden. Ein echtes Go-Programm gibt Adressen wie `0xc000012345` aus.
- Die Wörter „Zelle“ und „Namenstabelle im Modell“ gehören zum Modell der Seite. Im Kurstext steht in Abschnitt 7.4: Der Wert eines Zeigers ist die Adresse, unter der eine Variable im Speicher abgelegt ist. Was der Rechner dabei mit den Zellen tut und was die Namenstabelle im Modell dazu angibt, steht nicht im Kurstext. Dass ein Zeiger 8 Zellen braucht, steht dort auch nicht.
- Das Modell vereinfacht. In Go ordnet der Compiler jedem Namen beim Übersetzen einen Typ und eine Adresse zu. Ein laufendes Go-Programm schlägt keinen Namen mehr nach.
- Ein Wert aus mehreren Zellen steht im Modell so, dass die vorderen Bits in der Zelle mit der kleineren Adresse stehen. Die hinteren Bits einer Adresse stehen also in der letzten Zelle des Zeigers.
- In Go lässt sich mit einer Adresse nicht rechnen, und ein Programm kann eine Adresse nicht Bit für Bit verstellen. Das Modell erlaubt es in Schritt 2, damit man sieht, was eine Adresse ist.
- Auf der Seite läuft kein Go. Der Browser stellt nach, was das Programm tut. Der Student kann keinen eigenen Code eingeben.
- Der Null-Wert `nil`, Zeiger als Parameter und als Rückgabe, die Lebenszeit einer Variable und Speicherlecks kommen auf dieser Seite nicht vor. Dafür gibt es die Lektion „Zeiger in Go“.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text erklärt Zelle, Adresse und Namenstabelle im Modell, mit dem Satz, dass es diese Tabelle so nicht gibt. Bisher kam der Rechner immer über einen Namen zur Adresse. Die Adresse ist selbst eine Zahl und lässt sich in einer Variable ablegen. Eine solche Variable heißt Zeiger.
- Ein Zeiger hat in der Namenstabelle eine Zeile wie jede Variable.
- Mit `new` reserviert das Programm Zellen, für die in der Namenstabelle kein Name steht.
- Eine Liste nennt zwei HTML-Seiten, die dieser vorausgehen: „Speichermodell: Name, Typ und Adresse“ zu Kapitel 4 und „Speichermodell: Arrays“ zu Kapitel 6.
- Ein gelber Kasten sagt, dass das Modell vereinfacht und was in Abschnitt 7.4 des Kurstexts steht.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Den Links zu den zwei HTML-Seiten folgen.
- In diesem Schritt lässt sich nichts ausprobieren. Die Demo steht in Schritt 2.

### Abschnitte des Kurstexts

- 7.4
- 7.4.5

## Schritt 1: Drei Schreibweisen

Reiter: schreibweisen

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit einer Tabelle und Text. Eine Demo gibt es hier nicht.
- Die Tabelle hat die Spalten „Schreibweise“ und „Was der Rechner tut“ und drei Zeilen.
- `&b`: Das Ergebnis ist die Adresse von `b`. Der Compiler hat sie beim Übersetzen festgelegt. Im Modell steht sie in der Namenstabelle in der Spalte Adresse.
- `*p`: Der Rechner liest die Adresse, die in `p` steht, und geht zu den Zellen an dieser Adresse.
- `new(int32)`: Der Rechner reserviert 4 Zellen ohne Namen und liefert ihre Adresse.
- Darunter steht: Jeder Zeiger braucht 8 Zellen. Im Kurstext heißt `&` Adressoperator und `*` Dereferenzierungsoperator, und eine Variable, die `new` erzeugt hat, heißt anonyme Variable.
- Der Stern hat in Go zwei Bedeutungen neben der Multiplikation. In einem Typ wie `*int32` gehört er zum Namen des Typs. Vor einem Zeiger wie in `*p` ist er der Dereferenzierungsoperator. Im Kurstext steht das in Hinweis 7.3.

### Handgriffe

- In diesem Schritt lässt sich nichts ausprobieren. Die Tabelle ist fest.
- Mit „Weiter zu Schritt 2“ zur Demo wechseln.

### Abschnitte des Kurstexts

- 7.4
- 7.4.5

## Schritt 2: Ein Programm im Speicher

Reiter: speicher

### Aufbau der Darstellung

- Oben steht das Programm als nummerierte Liste mit sieben Zeilen: `var b int32`, `b = 1000`, `p := &b`, `*p = 2000`, `q := new(int32)`, `*q = *p`, `fmt.Println(b, *q)`. Die Zeilen stehen im Rumpf von `main`.
- Die zuletzt ausgeführte Zeile steht weiß auf dunklem Grund. Die Zeilen davor sind blass.
- Darunter stehen drei Knöpfe: „nächste Zeile ausführen“, „eine Zeile zurück“ und „von vorn“.
- Links liegt der Hauptspeicher: 32 Zellen mit den Adressen 100 bis 131. In jeder Zelle steht oben klein die Adresse und darunter das Bitmuster aus acht Bits.
- Die Zellen einer Variable haben eine Farbe. Ein schwarzer Pfeil führt von einem Zeiger zu der Zelle, deren Adresse in ihm steht.
- Eine Zelle mit blauem Rand hat die letzte Zeile gelesen. Eine Zelle mit rotem Rand hat die letzte Zeile beschrieben.
- Unter dem Hauptspeicher steht eine Zelle groß: acht Bits als Knöpfe, über jedem Bit seine Wertigkeit. Darüber steht, zu welcher Variable die Zelle gehört.
- Rechts steht die „Namenstabelle im Modell“ mit den Spalten Name, Typ, Adresse und Zellen. Bei der Variable, die `new` angelegt hat, ist das Feld für den Namen leer.
- Darunter steht der Kasten „Was der Rechner in dieser Zeile tut“ mit den Handlungen des Rechners als nummerierte Liste.
- Der Ablauf, Zeile für Zeile:
  1. `var b int32` reserviert die Zellen 100 bis 103 und füllt sie mit Nullen.
  2. `b = 1000` schreibt 1000 in die Zellen 100 bis 103.
  3. `p := &b` legt `p` vom Typ `*int32` an. `p` belegt die 8 Zellen 104 bis 111. In ihnen steht die Zahl 100, die Adresse von `b`. Der Pfeil führt von `p` zu Zelle 100.
  4. `*p = 2000` liest die Adresse 100 aus `p` und schreibt 2000 in die Zellen 100 bis 103. Die Zeile nennt den Namen `b` nicht, trotzdem hat `b` jetzt den Wert 2000.
  5. `q := new(int32)` reserviert die Zellen 112 bis 115 ohne Namen und legt `q` vom Typ `*int32` in den Zellen 116 bis 123 an. In `q` steht die Adresse 112.
  6. `*q = *p` liest über `p` den Wert 2000 aus den Zellen 100 bis 103 und schreibt ihn über `q` in die Zellen 112 bis 115. Danach gibt es den Wert 2000 zweimal, an zwei Adressen.
  7. `fmt.Println(b, *q)` gibt `2000 2000` aus.
- `p` und `q` zeigen auf zwei verschiedene Variablen. Der Vergleich `p == q` wäre `false`, der Vergleich `*p == *q` wäre nach Zeile 6 `true`.
- Schaltet der Student ein Bit in den Zellen eines Zeigers um, steht im Zeiger eine andere Adresse, und der Pfeil zeigt auf eine andere Zelle. Beispiel: Nach Zeile 3 steht in den Zellen 104 bis 111 die Zahl 100. Schaltet der Student in Zelle 111 das Bit mit der Wertigkeit 1 an, steht dort 101. `*p` würde dann die Zellen 101 bis 104 als einen `int32` lesen: drei Zellen von `b` und die erste Zelle von `p`.
- Steht in einem Zeiger eine Adresse, die es im Modell nicht gibt, bricht die nächste Zeile mit `*p` ab.

### Handgriffe

- „nächste Zeile ausführen“ wählen.
- „eine Zeile zurück“ wählen: Der Stand vor der letzten Zeile kommt zurück, auch von Hand umgeschaltete Bits.
- „von vorn“ wählen: Alle Zellen sind wieder frei, die Namenstabelle zeigt noch keine Zeile.
- Eine Zelle im Hauptspeicher anklicken: Sie erscheint unten groß.
- Eine Zeile der Namenstabelle anklicken. Der Mauszeiger über einer Zeile hebt die Zellen der Variable hervor.
- Ein Bit der groß gezeigten Zelle anklicken: Es wechselt zwischen 0 und 1.
- Das Programm lässt sich nicht ändern.

### Abschnitte des Kurstexts

- 7.4
- 7.4.5

## Aufgaben

Reiter: aufgaben

### Aufbau der Darstellung

- Zu sehen sind zwei Karten mit Text. Eine Demo gibt es hier nicht.
- Die erste Karte heißt „Erst vorhersagen, dann ausprobieren“ und stellt fünf Aufgaben. Die Aufgaben 1 bis 4 gehören zu Schritt 2, Aufgabe 5 gehört zu Schritt 1.
- Unter jeder Aufgabe steht eine aufklappbare „Antwort“. Sie ist verdeckt, bis der Student sie aufklappt.
- Die Antworten: Aufgabe 1: In den Zellen von `p` steht 100, die Adresse von `b`. Aufgabe 2: `b` hat den Wert 2000. Aufgabe 3: In `p` steht 101, `*p` läse drei Zellen von `b` und die erste Zelle von `p`. Aufgabe 4: Das Programm erreicht die Zellen nur über `q`. Aufgabe 5: 8 Zellen, weil jede Adresse gleich lang ist.
- Fragt der Student nach einer Aufgabe, deren Antwort er noch nicht aufgeklappt hat, bekommt er einen Hinweis auf den Weg und nicht gleich das Ergebnis.
- Zu Aufgabe 4 gehört aus dem Kurstext: Eine Variable, auf die kein Zeiger mehr zeigt, ist unerreichbar. Die Speicherbereinigung von Go gibt ihren Speicherplatz frei. Das steht in Abschnitt 7.4.3 und 7.4.4.
- Die zweite Karte nennt die vier HTML-Seiten mit dem Speichermodell: primitive Datentypen in Kapitel 4, Arrays in Kapitel 6, diese zu Kapitel 7 und Verbund in Kapitel 8.
- Am Ende des Schritts steht „Nächste Schritte“:
  - Nächste Lektion: „Das Kopier-Problem: Arrays als Argument“.
  - Quiz: „Funktionen, Zeiger und Rekursion“.

### Handgriffe

- Eine „Antwort“ aufklappen.
- Für die Aufgaben 1 bis 4 mit den Reitern oben zu Schritt 2 wechseln und dort ausprobieren.
- Den Links zu den drei anderen HTML-Seiten mit dem Speichermodell folgen.
- In diesem Schritt selbst lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 7.4
- 7.4.3
- 7.4.5

## Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - 1 Drei Schreibweisen, 7.4, Seiten 173–174; 7.4.5, Seiten 183–184: new: Die Operatoren & und * und die Funktion new stammen aus dem Kurstext. Was der Rechner dabei mit den Zellen tut und was die Namenstabelle im Modell dazu angibt, steht nicht im Kurstext.
  - 2 Ein Programm im Speicher, 7.4, Seiten 173–178: Zeiger, Adresse und Dereferenzieren stammen aus dem Kurstext. Das Programm Zeile für Zeile mit Hauptspeicher und Namenstabelle gibt die Lektion dazu.
  - Aufgaben, kein eigener Abschnitt: Die Aufgaben beziehen sich auf die Abschnitte der Schritte 1 und 2. Die Aufgaben gibt die Lektion dazu.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen diese Abschnitte:
  - Abschnitt 2.2, Seiten 8–10, Speicherplatz: Lektion: Variablen, Typen und Speicher, Schritt 4.
  - Abschnitt 4.3.4, Seite 63, Größe der Typen: Lektion: Speichermodell.
  - Abschnitt 7.3 und 7.4, Seiten 171–184, Zeiger: Lektion: Zeiger in Go.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu „Aufgaben““ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.2
- 4.3.4
- 7.3
- 7.4
- 7.4.5
