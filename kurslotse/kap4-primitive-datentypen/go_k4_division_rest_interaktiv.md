# Rechnen mit Ganzzahlen: Division und Rest in Go

Die Lektion erklärt die Grundrechenarten mit Integer-Werten, vor allem ganzzahlige Division und Rest. Das eigene Einstiegsbeispiel verteilt 17 Gegenstände auf Schachteln mit je 5 Plätzen. Negative Zahlen werden gesondert auf einer Zahlengeraden behandelt. Die Demos veranschaulichen Go mit JavaScript; sie führen keine Go-Programme aus. Nur der Editor im Schritt „Selbst ausprobieren“ kann im Editorlotsen mit Go laufen. Die Programme sind eigene Beispiele, keine Listings des Kurstexts. Aufbau und Auswertung von Ausdrücken bleiben bei der Lektion zu 4.1, Typkonvertierung bei 4.7, Überlauf bei 4.3.8.

## Schritt 0: Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Der Einstieg knüpft an Ausdrücke aus Abschnitt 2.3 und an Zuweisung und Auswertungsreihenfolge aus 3.3.5 und 3.3.6 an.
- 17 Gegenstände ergeben bei je 5 Plätzen 3 volle Schachteln und 2 übrige Gegenstände. Die Anzahl voller Schachteln wird als Ganzzahlquotient, die übrige Anzahl als Rest eingeführt.
- Der feste Go-Code verwendet anzahl und plaetze vom Typ int. Die zwei Println-Aufrufe geben 3 und 2 in getrennten Zeilen aus. Die Schachtelvorstellung gilt zunächst nur für nichtnegative Anzahlen und positive Gruppengrößen.
- Der Link „Operatoren und Ausdrücke“ führt zur allgemeinen Lektion. „Zuweisen, rechnen, ausgeben“ wird mit Namen und Ort Editorlotse genannt, ohne Link auf eine lokale Moodle-Lektion.

### Handgriffe

- Hier lässt sich nichts ausprobieren. Die Codeblöcke sind fest. Über die Reiter oder „Weiter“ geht es zum nächsten Schritt.

### Abschnitte des Kurstexts

- 2.3
- 3.3.5
- 3.3.6
- 4.3.7.1

## Schritt 1: Rechenzeichen

Reiter: rechenzeichen

### Aufbau der Darstellung

- Eine Tabelle zeigt a + b, a - b und a * b für int-Variablen a = 7 und b = 2. Alle Ergebnisse sind vom Typ int. Der Vorbehalt des Wertebereichs verweist auf Überlauf.
- Unäres -a hat einen Operanden und bedeutet 0 - a. Binäres a - b hat zwei Operanden. In a - (-b) ist das äußere Minus binär und das innere unär.
- Der feste Code gibt -7, 5, 9 und 7 aus. Unäres +a bedeutet 0 + a und ändert den Wert nicht.
- Eine Vorhersagefrage behandelt a = -7. Der Student kann diese Änderung erst im Schritt „Selbst ausprobieren“ vornehmen.

### Handgriffe

- Hier lässt sich nichts eingeben oder ausführen. Der Link zum Überlauf führt zur Lektion zu 4.3.8.

### Abschnitte des Kurstexts

- 4.3.7.1
- 4.3.8

## Schritt 2: Vollständige Gruppen

Reiter: gruppen

### Aufbau der Darstellung

- Oberhalb der Demo wird erklärt, dass nur vollständige Gruppen gezählt werden. Die blau umrandeten Schachteln enthalten je plaetze blaue Punkte; ein Punkt steht für einen Gegenstand. Der orange umrandete Kasten enthält nur die übrigen Gegenstände als orange Punkte. Es wird keine teilweise gefüllte Schachtel als volle gezählt.
- Unter dem Bild stehen Quotient und Rest, der passende Go-Code mit int-Variablen und die Ausgabe in zwei Zeilen. Die Werte der Darstellung stehen im Stand der Demo.
- Bei 0 Gegenständen erscheinen keine vollen Schachteln und keine Punkte; der Kasten „Übrig: 0“ bleibt sichtbar. Geht die Division auf, ist der Rest 0.
- 9 / 5 ergibt 1, nicht 1.8 und nicht 2. Die erforderliche Gesamtzahl zum Verpacken aller Gegenstände wird nicht mit der Zahl voller Schachteln gleichgesetzt.
- Die Frage zu 14 und 5 ist eine Vorhersageaufgabe zum anschließenden Einstellen, kein bewertetes Quiz.

### Handgriffe

- Anzahl: ganze Zahlen 0 bis 30. Plätze je Schachtel: ganze Zahlen 1 bis 10. Die Werte werden unmittelbar angezeigt.
- Vier Beispielknöpfe: „17 und 5“, „9 und 5“, „15 und 5“, „0 und 5“.
- Ungültige oder leere Eingaben zeigen eine Erklärung; Darstellung und Stand behalten die zuletzt gültigen Werte. Die Eingabefelder sind keine Go-Editoren.

### Abschnitte des Kurstexts

- 4.3.7.1

## Schritt 3: Negative Operanden

Reiter: negative

### Aufbau der Darstellung

- Dividend ist die Zahl links von /, Divisor die Zahl rechts davon. Der rationale Quotient wird als Ergebnis der üblichen Division eingeführt. Bei -9 und 5 ist er -1.8; Go liefert durch Abschneiden in Richtung null -1. Abrunden ergäbe -2.
- Die Zahlengerade zeigt den rationalen Quotienten als grünen Punkt, das Go-Ergebnis als blauen Ring und das Abrunden als rotes Quadrat. Der blaue Ring kann das rote Quadrat überdecken, wenn beide dieselbe Zahl darstellen. Beschriftungen nennen alle drei Werte zusätzlich ausdrücklich. Auf einem schmalen Bildschirm bis 650 Pixel ist die Schrift größer. Der Pfeil über der Zahlengerade führt in Richtung null.
- Die sechs Beispiele stehen fest. Bei -2 / 5 ist das Ergebnis 0, ohne gesonderte negative Null. Bei -10 / 5 liegen alle Marken bei -2, und es gibt keinen Pfeil.
- Die Zahlengerade ist ein Vergleich mathematischer Regeln. Go rechnet bei einer ganzzahligen Division nicht erst in float64. Der Link zur Typkonvertierung behandelt den Weg zu einem Ergebnis mit Nachkommastellen.

### Handgriffe

- Wählbar sind 9 / 5, -9 / 5, 9 / (-5), -9 / (-5), -2 / 5 und -10 / 5. Der gewählte Knopf ist blau hinterlegt.
- Hier gibt es keine freien Eingabefelder und keinen Editor.

### Abschnitte des Kurstexts

- 4.3.7.1
- 4.7

## Schritt 4: Quotient und Rest

Reiter: rest

### Aufbau der Darstellung

- Zuerst stehen die Begriffe: a / b liefert Ganzzahlquotient q, a % b liefert Rest r. % heißt auch Modulo-Operator und ist nur für ganze Zahlen definiert. Der Rest ist keine Nachkommastelle.
- Erst danach folgen a = q * b + r und a = (a / b) * b + a % b. Das sind mathematische Gleichungen, keine Go-Zuweisungen. Der Code darunter prüft Gleichheit mit ==.
- Die Demo zeigt die zusammengesetzte Rechnung, die Betragsregel, Go-Code mit a und b vom Typ int und die Ausgabe. Ein fester Tabellenblock zeigt die vier Vorzeichenkombinationen von 17 und 5.
- Ein von null verschiedener Rest hat das Vorzeichen des Dividenden. Rest 0 hat kein positives oder negatives Vorzeichen. Betrag bedeutet Abstand von null; |r| < |b| bedeutet: Der Betrag des Rests ist kleiner als der des Divisors.
- b = 0 wird in der Demo abgefangen. Es gibt keine berechneten Ausgaben. Die Erklärung ist eine Aussage der HTML-Seite, keine tatsächlich erzeugte Meldung des Compilers.
- Ein aufklappbarer Kasten „Ergänzung: eine Grenze des Wertebereichs“ erklärt die Ausnahme kleinster vorzeichenbehafteter Wert geteilt durch -1 und die Negation dieses Werts anhand einer int8-Variablen mit -128. Beide liefern -128; der Rest ist 0. Die mathematische Gleichung über unbegrenzten ganzen Zahlen gilt dabei nicht. In Go läuft auch das Produkt über. Bei Konstanten prüft der Compiler den Wertebereich dagegen schon beim Übersetzen; der Kasten verlinkt dazu den Zahlenkreis, Schritt 3. Typisierte und untypisierte Konstanten kommen erst in 5.2 und sind hier nicht genannt. Der Kasten nennt ausdrücklich die Go-Sprachspezifikation als zusätzliche Quelle; Abschnitt 4.3.7.1 behandelt den Sonderfall nicht.

### Handgriffe

- a und b: ganze Zahlen von -30 bis 30. b = 0 ist zum Anzeigen der Erklärung erlaubt. Die Demo führt die unzulässige Division nicht aus.
- Beispielknöpfe für alle vier Vorzeichenkombinationen mit 17 und 5, dazu a = 0, b = 5 und a = 17, b = 0.
- Ungültige oder leere Eingaben zeigen eine Erklärung. Es gelten weiter die zuletzt gültigen Werte.
- Ergänzung auf- oder zuklappen. Der Code in diesem Schritt kann nicht geändert oder ausgeführt werden.

### Abschnitte des Kurstexts

- 4.3.7.1
- 4.3.8

## Schritt 5: Typische Fehler

Reiter: fehler

### Aufbau der Darstellung

- Der feste Rechenweg zeigt 7 / 2 * 2, danach 3 * 2 und 6. / und * haben dieselbe Bindungsstärke und werden hier von links nach rechts verknüpft. (7 / 2) * 2 + 7 % 2 ergibt wieder 7. 7 / (2 * 2) ergibt dagegen 1.
- Der Typ bleibt auch bei Variablen ausschlaggebend: a = 9 und b = 2 vom Typ int ergeben bei a / b den Wert 4. Eine nachträgliche Konvertierung holt keine Nachkommastellen zurück.
- / und % mit einem konstanten Divisor 0 führen schon beim Übersetzen zu „invalid operation: division by zero“. Bei einer int-Variablen b = 0 übersetzt Go die Rechnung 7 / b, aber die Ausführung bricht mit „panic: runtime error: integer divide by zero“ ab. Für 7 % b gilt dasselbe.
- Die Rechnungen im pre ohne data-go sind kein Go-Programm. Der Text und der Stand unterscheiden sie von Programmzeilen.

### Handgriffe

- Dieser Schritt hat keine Demo; die Programmblöcke sind fest. Links führen zur Auswertung und Typkonvertierung.

### Abschnitte des Kurstexts

- 4.1.2
- 4.3.7.1
- 4.7

## Schritt 6: Selbst ausprobieren

Reiter: ausprobieren

### Aufbau der Darstellung

- Vorhersageaufgaben zum Negieren, zu Vorzeichen von Quotient und Rest, zur Klammerung und zum Rest bei null als Divisor. Der Kurslotse begleitet diese Aufgaben mit Hinweisen und Rückfragen, ohne ihre Lösungen vorwegzunehmen.
- Der Kasten enthält sechs eigene Programme: Rechenzeichen, Quotient und Rest, Kürzen?, Null zur Laufzeit, Null beim Übersetzen und Grenze von int8.
- Die festgehaltenen Ausgaben und Meldungen wurden mit Go geprüft. Die vollständigen Programme stehen im Kasten; die Grundstruktur und fmt.Println sind aus Kapitel 3 bekannt.
- Im Editorlotsen ist der Editor mit „Ausführen“ verfügbar. Auf GitHub Pages steht der feste Code mit Ausgabe oder Meldung und der Anleitung zum Öffnen im Editorlotsen. Der aktuelle Code und das Ergebnis gehen über probe.zeilen() und probe.stand() an den Kurslotsen.
- Am Ende des Schritts steht „Wie es weitergeht“ mit Links auf die Lektionen „Operatoren und Ausdrücke“, „Überlauf bei Ganzzahlen“ und „Typkonvertierung“, Schritt 4, und auf das Quiz zu Division und Rest.

### Handgriffe

- Programm wählen. Im Editorlotsen ändern und ausführen, „Programm zurücksetzen“. Auf GitHub Pages lassen sich Programme wählen, aber nicht ausführen.
- Den Links unter „Wie es weitergeht“ folgen.

### Abschnitte des Kurstexts

- 4.3.7.1

## Schritt 7: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Tabelle je Lernschritt mit den Spalten Schritt, Abschnitt im Kurstext, was aus dem Kurstext stammt, was nicht im Kurstext steht.
- Der Schachtelvergleich, die Zahlengerade, Aufgaben, Programme und genaue Meldungen sind eigene Ergänzungen. Die Rechenregeln stammen aus 4.3.7.1; der seltene Überlauf-Sonderfall aus der Go-Sprachspezifikation. In 4.3.7.1 stehen mehrere Zahlenbeispiele, etwa 9/5, 6/5, -6/5, -9/5, 7/2*2 und Vorzeichenkombinationen mit 5 und 3. Die eigenen Schachteln und der Sonderfall am kleinsten vorzeichenbehafteten Wert sind davon zu unterscheiden.
- Kapitel 2 und 3 werden als Vorwissen genannt, einschließlich „Zuweisen, rechnen, ausgeben“ im Editorlotsen, mit einem Link auf den Exkurs „Speichern und Rechnen im Rechner“.

### Handgriffe

- Hier lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.3
- 3.3.5
- 3.3.6
- 4.1.2
- 4.3.7.1
- 4.3.8
- 4.7
