# Exkurs: Warum Moodle am 21. Oktober 2026 nicht erreichbar war

Der Exkurs wendet Abschnitt 2.7 des Kurstexts auf einen wirklichen Fall an. Am 21. und 22. Oktober 2026 war das Moodle der FernUniversität von außen nicht erreichbar. Der Grund war eine Sicherheitslücke im Interpreter von PHP. Moodle ist in PHP geschrieben, selbst aber war es fehlerfrei.

Die Seite hat fünf Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Der Exkurs erklärt den Fehler nicht in der Tiefe. Er nennt die zwei Speicherfehler in einem Satz: Der Interpreter las Speicherplatz, in den er noch keinen neuen Wert geschrieben hatte, und benutzte freigegebenen Speicherplatz weiter. Im Vordergrund steht, was der Angreifer mit den Rechten des Webservers erreichen kann. Wie der Angriff im Einzelnen abläuft, steht auf der Seite nicht und ist nicht Stoff des Kurses.
- Im Kurstext stehen die imperative Programmierung (Abschnitt 2.1), der Compiler und der Interpreter (Abschnitt 2.7), die Standardbibliothek von Go (Abschnitt 3.1), `go run` (Abschnitt 3.2), Garbage-Collector und Null-Wert mit dem Sprachvergleich 5.1 zu C (Abschnitt 5.1). PHP, Opcodes, Moodle, Webserver und die Sicherheitslücke stehen nicht im Kurstext.
- Die Angaben zu den Lücken stammen aus den Meldungen des PHP-Projekts vom 21.10.2026: Die Versionen 8.5.12, 8.4.27, 8.3.36 und 8.2.35 beheben zwei kritische Lücken. Ein Angreifer konnte ohne Anmeldung Code auf dem Server ausführen. Der Fehler wirkte beim Verarbeiten einer Anfrage mit einer Datei zum Hochladen, bevor ein Programm der Website lief. PHP 8.1 und älter bekommen keine Korrektur.
- Zur Ausnutzung steht auf der Seite: erster Angriffsversuch am 17.10.2026 nach Amazon, Cloudflare zählte schon zehn Stunden vor der Veröffentlichung Tausende Angriffsversuche, danach Zehntausende bis Millionen Angriffsversuche nach Cloudflare. Einen gelungenen Einbruch hatten die Behörden in Deutschland, den Niederlanden und Japan bis zum 22.10. nicht bestätigt. Der Kurslotse behauptet keinen Einbruch.
- Die Versionsnummern, die betroffenen Einrichtungen und die Zahlen zu den Angriffsversuchen gelten für Oktober 2026.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text sagt, dass Moodle am 21. und 22. Oktober 2026 von außen nicht erreichbar war und dass der Fehler in einem Programm steckte, das Moodle zum Laufen braucht.
- Eine Liste nennt die zwei Wege aus Abschnitt 2.7: mit Compiler wird vorher in Maschinensprache übersetzt, mit Interpreter liest ein weiteres Programm Befehl für Befehl und bildet die Wirkung nach.
- Unter „Was PHP ist“: PHP ist wie Go eine imperative Sprache nach Abschnitt 2.1, entwickelt für Websites. Ein PHP-Programm läuft auf dem Server und stellt die HTML-Seite zusammen. PHP übersetzt ein Programm zuerst in einen Zwischencode, die Opcodes, und der Interpreter führt ihn aus. Ein Link führt zur Zwischensprache in Schritt 2 der Lektion zu Abschnitt 2.7. Go wird vollständig in Maschinensprache übersetzt.
- Eine zweite Liste nennt die drei folgenden Schritte. Ein Link führt zur Lektion „Höhere und niedrigere Programmiersprachen“.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- Sonst lässt sich in diesem Schritt nichts anklicken oder ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.7

## Schritt 1: Was geschah

Reiter: geschehen

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Unter „Die Lage am 21. und 22. Oktober 2026“ steht eine Liste: In Japan schalteten die Banken MUFG und Mizuho, ANA und JR West Teile ihrer Websites ab, die Premierministerin forderte die Betreiber auf, ihre Websites zu prüfen und Dienste notfalls auszusetzen. In Europa gingen Websites von Behörden in Frankreich, Dänemark, Finnland und den Niederlanden vom Netz. Die Universitäten Wien und Freiburg schalteten ab, an der FernUniversität war Moodle zwei Tage nicht erreichbar.
- Unter „Die Lücke“ steht zuerst: Moodle ist in PHP programmiert, ein Fehler im Interpreter von PHP betraf deshalb auch das Moodle der FernUniversität. Danach nennt der Text die vier neuen Versionen von PHP vom 21.10.2026 und sagt, dass ein Angreifer ohne Anmeldung Code auf dem Server ausführen konnte.
- PHP ist weit verbreitet: Nach Schätzungen läuft etwa 70 Prozent aller Websites mit PHP. In der Umfrage von Stack Overflow 2026 steht es auf Platz 13, direkt vor Go auf Platz 14. Ein Link führt zur Lektion „Imperative Programmierung“, Schritt 1, mit den Balken der Umfrage.
- Unter „Warum es so eilte“ nennt eine Liste zwei Gründe: die Lücke selbst, die in jeder verbreiteten Version des Interpreters steckte, und die Geschwindigkeit der Angreifer. Amazon sah einen ersten Angriffsversuch am 17.10., Cloudflare zählte zehn Stunden vor der Veröffentlichung Tausende Angriffsversuche, Cloudflare zählte nach der Veröffentlichung in 45 Minuten Zehntausende, nach zwei Stunden Hunderttausende und bis zum 22.10. Millionen Angriffsversuche. Den Betreibern blieb keine Zeit, wer nicht sofort aktualisieren konnte, nahm die Server vom Netz. Einen gelungenen Einbruch hatten die Behörden in Deutschland, den Niederlanden und Japan bis zum 22.10. nicht bestätigt.
- Eine Liste fasst die Meldungen zusammen: Der Fehler steckte im Interpreter, der in C geschrieben ist. Er wirkte beim Annehmen einer Anfrage mit einer Datei zum Hochladen. Betroffen war jede Website auf einem Server mit einer betroffenen Version.
- Darunter stehen die Quellen mit Links: php.net, die Sicherheitsmeldungen GHSA-w8hw-w9fx-j79w und GHSA-6wmf-4rv5-p3rp, Japan Times, ABC7, IBTimes, Amazon und The Next Web.

### Handgriffe

- Den Links zu den Quellen folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

Der Kurstext hat zu diesem Schritt keinen Abschnitt.

## Schritt 2: Wo der Fehler steckte

Reiter: fehler

### Aufbau der Darstellung

- Eine Liste nennt die drei Programme, durch die eine Anfrage auf dem Server geht: der Webserver nimmt sie an; der Interpreter von PHP, ein Programm in Maschinensprache, übersetzt aus C, verarbeitet sie und führt danach Moodle aus; Moodle, ein Programm in PHP, prüft die Anmeldung.
- Der Prozessor führt den Interpreter aus, und der Interpreter führt Moodle aus. Das ist der Weg mit Interpreter aus Schritt 2 der Lektion zu Abschnitt 2.7, ein Link führt dorthin.
- Ein gelber Kasten stellt eine Frage zum Vorhersagen: An welcher Station kann Moodle einen Angreifer aufhalten, wenn der Fehler im Interpreter steckt? Die Antwort zeigt die Demo: an keiner, weil Moodle beim Angriff nie an die Reihe kommt.
- Darunter wählen zwei Knöpfe die Anfrage: „Normale Anfrage“ oder „Anfrage des Angreifers“. Die gewählte ist gelb hinterlegt.
- Darunter stehen die Knöpfe für den Ablauf und ein Zähler.
- Darunter stehen vier Stationen in einer Reihe, mit Pfeilen verbunden: Absender, Webserver, „Interpreter von PHP“, Moodle.
- Gelb unterlegt ist jede Station, die gerade arbeitet. Rot unterlegt ist der Interpreter, wenn dort der Fehler wirkt. Blass ist Moodle, wenn es nie an die Reihe kommt.
- Beim Angriff springt der Ablauf nicht zu Moodle weiter. Der Interpreter führt an der Stelle von Moodle den Code des Angreifers aus. Der Server läuft weiter, und der Code des Angreifers kann dort tun, was der Webserver darf. Moodle kommt bei dieser Anfrage nie an die Reihe.
- Unter den Stationen erklärt ein Kasten den sichtbaren Stand. Er ist grün, beim Fehler rot.
- Die normale Anfrage hat fünf Schritte: der Student schickt sie, der Webserver nimmt sie an, der Interpreter verarbeitet sie, der Interpreter führt Moodle aus und Moodle prüft die Anmeldung, der Webserver schickt die Kursseite an den Studenten zurück. Im letzten Schritt sind Webserver und Absender gelb.
- Die Anfrage des Angreifers hat fünf Schritte: der Angreifer schickt eine Anfrage ohne Anmeldung mit einer Datei zum Hochladen, der Webserver nimmt sie an, der Interpreter liest beim Verarbeiten Speicherplatz, dessen Inhalt der Angreifer bestimmt hat, und führt dadurch Code des Angreifers aus, der Code läuft mit den Rechten des Webservers und Moodle hat keine Zeile ausgeführt. Im fünften Schritt steht, was der Code mit diesen Rechten kann: die Konfigurationsdatei von Moodle mit dem Passwort der Datenbank lesen, eine eigene PHP-Datei als Hintertür ablegen, ein Programm herunterladen und starten. Daher der Name PHP2Shell, von PHP zu einer Kommandozeile auf dem Server. Der Code hat die Rechte des Webservers, keine weitergehenden.
- Unter der Demo steht „Warum der Angriff so gefährlich war“: Der Code läuft mit den Rechten des Webservers und darf alles, was Moodle darf. Eine Liste nennt zwei Beispiele: die Konfigurationsdatei von Moodle mit dem Passwort der Datenbank lesen und damit an alle Daten von Moodle kommen, etwa Namen, E-Mail-Adressen und Beiträge in den Foren; Programme herunterladen, ablegen und starten, etwa eine Hintertür.
- Unter „Was im Speicher schiefging“ steht in zwei Absätzen: Jede Variable belegt Speicherplatz im Hauptspeicher, Link auf „Variablen, Typen und Speicher“, Schritt 4. Der Interpreter ist in C geschrieben, reserviert Speicherplatz für eine Anfrage und gibt ihn danach frei. Er las Speicherplatz, in den er noch keinen neuen Wert geschrieben hatte, und benutzte freigegebenen Speicherplatz weiter. In beiden Fällen stand dort, was der Angreifer mit seinen Anfragen hineingebracht hatte. Wie der Angreifer das im Einzelnen erreicht, steht auf der Seite nicht.
- Unter „Was Go hier anders macht“ steht eine Liste: Null-Wert, jede Variable bekommt in Go sofort einen Wert, Speicherplatz mit altem Inhalt liest ein Go-Programm nicht, in C gibt es Variablen ohne Anfangswert (Sprachvergleich 5.1, Seite 115, Link auf „Variablendeklaration in Go“, Schritt 1). Garbage-Collector, er gibt Speicherplatz erst frei, wenn keine Variable ihn mehr benutzt (Abschnitt 5.1, Seite 113).
- Darunter steht, dass Abschnitt 7.4.4 den umgekehrten Fehler behandelt, Speicher, der nie freigegeben wird, mit Link auf „Zeiger in Go“, Schritt 5. Abschnitt 7.4.4 erklärt den Fehler im Interpreter nicht.

### Handgriffe

- „Normale Anfrage“ oder „Anfrage des Angreifers“ wählen. Der Ablauf beginnt dann von vorn.
- „Nächster Schritt“: Der Ablauf geht einen Schritt weiter. Am Ende ist der Knopf gesperrt.
- „Einen Schritt zurück“: Der Ablauf geht einen Schritt zurück. Am Anfang ist der Knopf gesperrt.
- „Von vorn beginnen“: Der Ablauf beginnt neu.
- Die Stationen lassen sich nicht anklicken. Der Student kann keine eigene Anfrage schicken und nichts am Ablauf ändern.

### Abschnitte des Kurstexts

- 2.2
- 2.7
- 5.1
- 7.4.4

## Schritt 3: Wie der Fehler behoben wird

Reiter: reparatur

### Aufbau der Darstellung

- Zwei Überschriften vergleichen, wie der Fehler behoben wird.
- PHP: Der Betreiber installiert den neuen Interpreter auf seinem Server. Der Quelltext von Moodle bleibt unverändert.
- Go: Der Compiler bindet die Pakete der Standardbibliothek beim Übersetzen in die Programmdatei ein. Ein Fehler in einem Paket steckt deshalb in jeder Programmdatei, die es benutzt. Jedes solche Programm muss mit der neuen Version von Go neu übersetzt werden.
- Beispiel aus derselben Woche: Go 1.26.10 vom 19.10.2026 behebt Fehler im Paket `crypto/x509`, das bei HTTPS anhand von Zertifikaten prüft, ob ein Server der ist, für den er sich ausgibt. Jedes Go-Programm, das solche Verbindungen herstellt, muss mit Go 1.26.10 neu übersetzt werden.
- Eine Tabelle vergleicht PHP und Go in drei Zeilen: wo der gemeinsame Teil steckt, wer nach einem Fehler etwas tun muss, was mit dem Programm selbst geschieht.
- Ein gelber Kasten stellt eine Frage zum Vorhersagen: Eine Hochschule betreibt 12 Server mit Moodle, ein Entwickler stellt 4 Go-Programme bereit, die 500 Studenten auf ihren eigenen Rechnern benutzen. Wie viele Programme müssen nach einem Fehler im Interpreter von PHP und nach einem Fehler in einem Paket von Go neu übersetzt werden, und wo wird etwas ausgetauscht?
- Die Antwort ist verdeckt, bis der Student auf „Antwort aufdecken“ klickt. Erst dann geht ihr Text mit dem Text der Seite mit. Sie lautet: Bei PHP wird kein Programm neu übersetzt, auf jedem der 12 Server wird der Interpreter ausgetauscht. Bei Go übersetzt der Entwickler die 4 Programme neu, danach laden die 500 Studenten die neuen Programmdateien herunter, bis dahin arbeitet jeder mit der fehlerhaften.
- Darunter steht, dass der Student in Abschnitt 3.2 sein erstes Go-Programm mit `go run` übersetzt.

### Handgriffe

- „Antwort aufdecken“ zeigt die Antwort, „Antwort wieder verdecken“ verbirgt sie.
- In diesem Schritt läuft kein Programm. Der Student kann hier nichts übersetzen oder ausführen.

### Abschnitte des Kurstexts

- 2.7
- 3.1
- 3.2

## Schritt 4: Bezug zum Kurstext

Reiter: kurstext

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Eine Tabelle hat je Schritt eine Zeile. Die Spalten: Schritt, Abschnitt im Kurstext mit gedruckten Seiten, was aus dem Kurstext stammt, was nicht im Kurstext steht.
  - Worum es geht, 2.1, Seiten 5–8, und 2.7, Seite 18: Die imperative Programmierung und die zwei Wege mit Compiler und Interpreter stammen aus dem Kurstext. Dass Moodle in PHP geschrieben ist und PHP einen Zwischencode verwendet, ergänzt der Exkurs.
  - 1 Was geschah: kein Abschnitt. Die Angaben stammen aus den Meldungen des PHP-Projekts, von Amazon und aus der Presse.
  - 2 Wo der Fehler steckte, 2.2, Seiten 8–10, 2.7, Seite 18, 5.1, Seiten 113 und 115, und 7.4.4, Seiten 182–183: Dass der Interpreter ein weiteres Programm ist, steht in Abschnitt 2.7. Null-Wert, Garbage-Collector und der Vergleich mit C stehen in Abschnitt 5.1. Webserver, Anfrage und Anmeldung ergänzt der Exkurs. Abschnitt 7.4.4 behandelt den umgekehrten Fehler.
  - 3 Wie der Fehler behoben wird, 2.7, Seite 18, 3.1, Seite 21, 3.2, Seiten 22–24: Dass der Compiler Bibliotheksprogramme einbindet, wird in Abschnitt 2.7 nur erwähnt. Die Standardbibliothek steht in Abschnitt 3.1. Dass deshalb jedes Programm neu übersetzt werden muss, folgert der Exkurs daraus.
- Unter „Wo das Thema im Kurstext noch vorkommt“ stehen die Abschnitte 2.1, 2.2, 2.7, 3.2, 5.1 und 7.4.4, mit Links auf die Lektionen „Imperative Programmierung“, „Variablen, Typen und Speicher“, „Höhere und niedrigere Programmiersprachen“, „Variablendeklaration in Go“ und „Zeiger in Go“.
- Unter „Stand der Angaben“ steht, dass Versionsnummern, Einrichtungen und Zahlen zu den Angriffsversuchen für Oktober 2026 gelten.

### Handgriffe

- Mit den Reitern oben oder mit „Zurück zu Schritt 3“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 2.1
- 2.2
- 2.7
- 3.1
- 3.2
- 5.1
- 7.4.4
