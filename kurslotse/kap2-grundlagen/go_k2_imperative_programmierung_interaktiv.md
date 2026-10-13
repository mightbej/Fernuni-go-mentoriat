# Imperative Programmierung: Anweisungen, Blöcke und Schleifen

Die Seite erklärt, was ein imperatives Programm ist: eine Folge von Anweisungen, die der Rechner nacheinander ausführt. Sie zeigt, welche Programmiersprachen heute am meisten benutzt werden, ordnet die vier Arten zu programmieren ein und lässt die Bausteine Verzweigung und Schleife Schritt für Schritt laufen, einmal mit Sprung und einmal mit Block.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Die Seite gehört zu Kapitel 2 des Kurstexts. Dort kommt noch kein Go vor. Die Programme der Seite gehören zu keiner bestimmten Programmiersprache. Die Wörter `IF`, `THEN`, `GOTO`, `WHILE`, `REPEAT` und `UNTIL` in Großbuchstaben sind kein Go.
- Go-Code steht auf der Seite nur in der Tabelle „Wie Go diese Bausteine schreibt“ in Schritt 3, und dort nur als einzelne Wörter: `if`, `else`, `for`, `break`.
- Der Student kann die Programme nicht ändern und keinen eigenen Code ausführen.
- Jede Demo behält ihren Stand, wenn der Student den Reiter wechselt.
- Klausurstoff ist Schritt 3. Schritt 1 und 2 dienen der Einordnung.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text erklärt das Wort „imperativ“: Es kommt vom lateinischen imperare, befehlen. In einer Sprache wie Go heißt ein Befehl Anweisung.
- Befehl und Anweisung meinen dasselbe auf zwei Ebenen: „Befehl“ sagt man, wenn man nah an der Hardware programmiert, „Anweisung“ in einer höheren Programmiersprache wie Go.
- Eine Liste nennt die zwei Teile eines imperativen Programms: die einzelne Anweisung und die Reihenfolge der Anweisungen.
- Eine zweite Liste sagt, was die drei Schritte der Seite zeigen.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.

### Abschnitte des Kurstexts

- 2.1

## Schritt 1: Welche Sprachen heute benutzt werden

Reiter: sprachen

### Aufbau der Darstellung

- Zu sehen sind 15 waagerechte Balken untereinander. Jeder Balken steht für eine Programmiersprache. Links steht der Name der Sprache, rechts der Anteil in Prozent.
- Die Sprache mit dem größten Anteil steht oben. Je länger der Balken, desto größer der Anteil. Der oberste Balken füllt die ganze Breite.
- Die Zahlen stammen aus der Umfrage „Stack Overflow Developer Survey 2026“. Gefragt wurde, mit welchen Sprachen die Entwickler im letzten Jahr gearbeitet haben. Jeder durfte mehrere Sprachen nennen. Die Anteile ergeben zusammen deshalb mehr als 100 %.
- Der Balken von Go ist grün, der Name Go auch. Die übrigen Balken sind grau.
- Unter den Balken steht der Knopf „Welche dieser Sprachen sind nicht imperativ?“. Solange die Antwort verdeckt ist, steht darunter in einem grünen Kasten nur, auf welchem Platz Go steht.
- Nach einem Klick auf den Knopf sind zwei Balken braun, und der grüne Kasten nennt die zwei Sprachen mit einer Begründung. Der Knopf heißt dann „Antwort wieder verdecken“.
- Die Einteilung in imperativ und nicht imperativ ist eine Vereinfachung der Seite. Einige der Sprachen erlauben mehrere Arten zu programmieren.
- Die Rangliste steht nicht im Kurstext.

### Handgriffe

- Den Knopf unter den Balken wählen: Er deckt die Antwort auf und verdeckt sie wieder.
- Die Balken lassen sich nicht anklicken. Ein anderes Jahr oder eine andere Umfrage lässt sich nicht wählen.

### Abschnitte des Kurstexts

Die Rangliste geht über den Kurstext hinaus. Zu diesem Schritt gibt es dort keinen Abschnitt.

## Schritt 2: Vier Arten zu programmieren

Reiter: arten

### Aufbau der Darstellung

- Zu sehen sind oben vier Knöpfe nebeneinander: „logisch“, „funktional“, „objektorientiert“ und „imperativ und strukturiert“. Der gewählte Knopf ist grün. Am Anfang ist „imperativ und strukturiert“ gewählt.
- Die vier Arten heißen Programmierparadigmen. Im Kurstext zeigt Abbildung 2.1 sie.
- Unter den Knöpfen steht in einem grünen Kasten der Grundgedanke der gewählten Art.
- Darunter steht die Zeitleiste, eine Tabelle mit drei Spalten: Jahr, Sprache oder Ereignis, Art zu programmieren. Sie hat 13 Zeilen. Die jüngste Sprache steht oben, die älteste unten.
- Die Zeilen der gewählten Art sind gelb hinterlegt und fett. Alle anderen Zeilen sind blass grau.
- Eine Zeile nennt keine Sprache, sondern ein Ereignis: „Kritik am Befehl GOTO“, ca. 1968. Sie gehört zur Art „imperativ und strukturiert“.
- Jede Sprache steht in der Zeitleiste bei genau einer Art. Das ist eine Vereinfachung der Seite.
- Die Jahreszahlen stehen nicht im Kurstext und sind ungefähre Angaben. Vor jeder steht „ca.“.
- Unter der Zeitleiste steht: Das Modul behandelt die strukturierte Programmierung. Die objektorientierte Programmierung baut darauf auf und folgt in späteren Modulen.

### Handgriffe

- Einen der vier Knöpfe wählen. Es ist immer genau eine Art gewählt.
- Die Zeilen der Zeitleiste lassen sich nicht anklicken und nicht umsortieren.

### Abschnitte des Kurstexts

- 2.1

## Schritt 3: Die Bausteine

Reiter: bausteine

### Aufbau der Darstellung

- Zu sehen sind zwei Karten untereinander. Die obere enthält die Demo, die untere die Tabelle „Wie Go diese Bausteine schreibt“.
- Oben in der Demo stehen drei Knöpfe für den Baustein: „Verzweigung“, „Schleife“ und „Zwei Schleifen im Vergleich“. Der gewählte Knopf ist grün.
- Darunter stehen die Knöpfe für den Startwert der Variable. Der gewählte Knopf ist grün.
- Darunter stehen die drei Knöpfe „nächster Schritt“, „einen Schritt zurück“ und „von vorn“.
- Darunter stehen zwei Programme nebeneinander, jedes in einem dunklen Kasten mit einer Überschrift. Auf einem schmalen Bildschirm stehen sie untereinander.
- Bei „Verzweigung“ und „Schleife“ steht links das Programm mit Sprung und rechts dasselbe Programm mit Block. Bei „Zwei Schleifen im Vergleich“ steht links `REPEAT-UNTIL` und rechts `WHILE`. Dort sind beide Programme strukturiert, einen Sprung gibt es nicht.
- Die zwei Programme gehören zusammen: Beide beginnen mit demselben Startwert, und ein Klick auf „nächster Schritt“ führt in beiden je eine Zeile aus.
- Hell hinterlegt mit einem gelben Strich am linken Rand ist in jedem Programm die Zeile, die zuletzt ausgeführt wurde. Am Anfang ist keine Zeile markiert.
- Die Programme mit Sprung haben vor jeder Zeile eine Zeilennummer. `GOTO 4` heißt: Weiter geht es bei der Zeile mit der Nummer 4. Die Programme mit Block haben keine Zeilennummern.
- Ein Block ist die Folge von Anweisungen zwischen `{` und `}`. Die Zeilen im Block sind eingerückt.
- „gib aus: groß“ gibt das Wort „groß“ aus. „gib aus: a“ gibt den Wert der Variable a aus. Das Zeichen `=` in `a = a - 3` ist eine Zuweisung.
- Unter jedem Programm steht ein Feld mit zwei Zeilen: der Wert der Variable und die Ausgabe bisher.
- Unter beiden Programmen steht in einem grünen Kasten je ein Satz für links und für rechts. Er erklärt, was die markierte Zeile getan hat.
- Die zwei Programme brauchen nicht immer gleich viele Schritte. Ist eines früher zu Ende, steht dort „Das Programm ist schon zu Ende“, und das andere läuft weiter.
- Baustein „Verzweigung“: Die Variable heißt a. Links prüft Zeile 1 die Bedingung `a <= 3` und springt, wenn sie erfüllt ist. Rechts prüft die erste Zeile das Gegenteil `a > 3` und führt den Block aus, wenn sie erfüllt ist. Beide Programme tun dasselbe.
- Baustein „Schleife“: Die Variable heißt n. Links entsteht die Schleife durch den Sprung zurück in Zeile 4, `GOTO 1`. Rechts steht eine `WHILE`-Schleife, sie prüft die Bedingung vor jedem Durchlauf.
- Baustein „Zwei Schleifen im Vergleich“: Die Variable heißt n. `REPEAT-UNTIL` prüft die Bedingung nach dem Durchlauf und läuft deshalb mindestens einmal. `WHILE` prüft sie davor und kann den Block ganz auslassen. Der Unterschied zeigt sich beim Startwert n = 0.
- `REPEAT-UNTIL` verlässt die Schleife, wenn die Bedingung erfüllt ist. `WHILE` verlässt sie, wenn die Bedingung nicht erfüllt ist.
- Die Bedingung `a > 3` stammt aus dem Kurstext. Die Anweisungen in den Programmen sind eigene Beispiele der Seite.
- Die Tabelle in der unteren Karte hat vier Zeilen und vier Spalten: der Baustein im Kurstext, die Schreibweise in Go, der Abschnitt im Kurstext dazu und die HTML-Seite dazu.
- Nach der Tabelle schreibt Go `if` und `else` für den Block mit `IF`, `for` mit einer Bedingung für die `WHILE`-Schleife und `break` für den Sprung aus einer Schleife heraus. Für `REPEAT-UNTIL` hat Go keine eigene Schreibweise.
- `BREAK` steht nur in der Tabelle. In den Programmen der Demo kommt es nicht vor. Auch `ELSE` und die `FOR`-Schleife zeigt die Demo nicht.

### Handgriffe

- Einen der drei Bausteine wählen. Der Lauf beginnt dann von vorn mit dem ersten Startwert.
- Einen Startwert wählen. Bei „Verzweigung“: a = 2 oder a = 5. Bei „Schleife“: n = 0, n = 2 oder n = 3. Bei „Zwei Schleifen im Vergleich“: n = 0 oder n = 2. Der Lauf beginnt danach von vorn.
- Für den Startwert gibt es nur diese Knöpfe und kein Eingabefeld. Ein Wert wie n = 10 oder a = 4 lässt sich auf der Seite nicht einstellen.
- „nächster Schritt“: Beide Programme führen eine Zeile aus. Sind beide Programme zu Ende, ist der Knopf gesperrt.
- „einen Schritt zurück“: Beide Programme gehen einen Schritt zurück.
- „von vorn“: Der Lauf beginnt neu, Baustein und Startwert bleiben.
- Den drei Links der Tabelle folgen: „if-Anweisungen“, „for-Anweisung mit Bedingung“ und „break-Anweisung“. Das sind eigene HTML-Seiten zu Kapitel 5 und keine Reiter dieser Seite.
- Die Programme lassen sich nicht ändern. Nur ein Programm allein laufen zu lassen geht nicht.

### Abschnitte des Kurstexts

- 2.1
- 5.5.1
- 5.5.3.1
- 5.5.3.3
