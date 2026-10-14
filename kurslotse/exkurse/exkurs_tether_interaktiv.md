# Vom Eisenbahnzug in den Maschinenraum von Tether

Die Seite ist der Exkurs zum Abschluss des Kurses, für Studenten im ersten Semester. Sie führt von der Lektion „Blockchain: eine verkettete Liste mit Fingerabdruck“ zu Bitcoin, Ethereum und dem digitalen Dollar USDT von Tether. Der Student findet das Programm von USDT im Internet und erkennt darin Begriffe aus dem Kurs wieder, bis zur Sperrliste. Am Ende steht die Frage, wer die Kontrolle über digitales Geld hat, mit dem digitalen Euro als Vergleich.

Die Seite hat vier Reiter. Sichtbar ist immer nur ein Schritt.

Für alle Schritte gilt:

- Nichts auf dieser Seite steht im Kurstext. Die Seite ist ein Exkurs des Mentors. Sie verwendet Begriffe aus dem Kurs: vorzeichenlose Typen aus Kapitel 4, `if` aus Kapitel 5, Funktionen aus Kapitel 7, Verbünde und Listen aus Kapitel 8 und 9.
- Die Seite erklärt jedes Fachwort einfach. Ein Smart Contract ist ein Programm, das in einer Blockchain liegt und von jedem Rechner im Netz gleich ausgeführt wird. Die Seite sagt dafür „Programm“. Ein Stablecoin ist ein digitaler Dollar, der immer einen US-Dollar wert sein soll.
- Das Programm von USDT auf Ethereum ist in Solidity geschrieben. Es liegt unter der Adresse `0xdAC17F958D2ee523a2206206994597C13D831ec7`. Das ist die Adresse des Programms, nicht die seines Eigentümers.
- Eigentümer des Programms ist Tether. Im Programm steht er in der Variable `owner`. Etherscan zeigt ihren Wert unter dem Reiter „Read Contract“. Am 14.10.2026 stand dort die Adresse `0xc6cde7c39eb2f0f0095f41570af89efc2c1ea828`. Das ist ein weiteres Programm, das einen Aufruf erst ausführt, wenn drei von sechs Schlüsselinhabern zugestimmt haben. Wer die sechs sind, steht nicht in der Blockchain. Die Seite selbst nennt diese Einzelheiten nicht.
- Die Seite beschreibt Technik und Regeln. Sie gibt keinen Rat, Kryptowährungen zu kaufen oder zu verkaufen.
- Alle Zahlen und Tatsachen haben den Stand 14.10.2026.
- Auf der Seite läuft kein Solidity und kein Go.
- Welche Aufgaben die Klausur stellt, ist nicht bekannt. Verbindlich sind der Kurstext und Moodle der FernUniversität. Der Inhalt dieser Seite gehört nicht zum Stoff des Kurstexts.

## Worum es geht

Reiter: worum

### Aufbau der Darstellung

- Zu sehen ist eine Karte mit Text. Eine Demo gibt es hier nicht.
- Der Text spannt den Bogen: Güterzug des Kurstexts, verkettete Liste, Blockchain als Kassenbuch auf vielen Rechnern, und jetzt ein echtes Programm, das so ein Kassenbuch führt. Das ist der Maschinenraum im Titel.
- Eine Liste nennt, was der Student aus dem Kurs braucht: Typen, Bedingungen, Funktionen, Verbünde und Listen.
- Unter „Die Schritte“ steht je ein Satz zu den drei Schritten.
- Ein gelber Kasten sagt, dass nichts davon im Kurstext steht und dass die Seite keinen Rat zum Kaufen gibt.

### Handgriffe

- Mit den Reitern oben oder mit „Weiter zu Schritt 1“ zu einem anderen Schritt wechseln.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

- 9.2

## Schritt 1: Bitcoin, Ethereum und Tether

Reiter: netze

### Aufbau der Darstellung

- Zu sehen ist Text mit drei Überschriften. Eine Frage zum Vorhersagen und eine Demo gibt es hier nicht.
- Bitcoin: digitales Geld ohne Bank, seit 2009, das Kassenbuch ist eine Blockchain wie in der Lektion zur Blockchain.
- Ethereum: eine zweite Blockchain, seit 2015. In ihren Blöcken stehen auch Programme. Jeder Rechner führt sie aus und kommt zum gleichen Ergebnis. Ein solches Programm heißt Smart Contract, wörtlich „kluger Vertrag“.
- Tether und USDT: Tether gibt den digitalen Dollar USDT heraus. Er hat keine eigene Blockchain, er ist ein Programm auf Ethereum und anderen Blockchains. Das Programm führt das Kassenbuch, wer wie viele USDT hat.
- Zwei Zahlen: Wert aller USDT 184 Mrd. US-Dollar, rund 60 Prozent aller Stablecoins, die zusammen rund 300 Mrd. US-Dollar wert sind.
- In der EU ist USDT nach den Regeln für Kryptowerte, kurz MiCA, nicht zugelassen. Große Börsen haben den Handel damit für Kunden in der EU 2025 eingestellt oder eingeschränkt. In Deutschland ist die BaFin zuständig.

### Handgriffe

- Den Links zu den Quellen folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

Der Schritt geht über den Kurstext hinaus.

## Schritt 2: Das Programm von USDT lesen

Reiter: lesen

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“: Für eine Sperre braucht ein Konto ein Feld vom Typ `bool`, in der Lektion zur Blockchain heißt es `gesperrt`.
- Der Text: Der Programmtext von USDT ist öffentlich, in der Sprache Solidity. Solidity kennt die gleichen Begriffe wie Go.
- Unter „So findest du den Programmtext“ stehen drei Schritte: die Seite von USDT bei Etherscan öffnen, den Reiter „Contract“ lesen, mit Strg+F nach `isBlackListed` suchen. Etherscan ist eine Webseite, die anzeigt, was in der Blockchain Ethereum steht. „Source Code Verified“ heißt: Etherscan hat den Text selbst übersetzt und mit dem Programm in der Blockchain verglichen, es ist wirklich der Text des laufenden Programms.
- Die Demo: fünf Knöpfe, „Guthaben“, „Sperrliste“, „Konto sperren“, „Absender prüfen“ und „Nur der Eigentümer“. Links steht die Zeile aus dem Programm in Solidity, wörtlich. Rechts steht ungefähr das Gleiche in Go. Darunter erklärt ein Kasten die Zeile.
- `mapping` ordnet jedem Konto, in Solidity `address`, einen Wert zu. In Go gibt es dafür `map`, das im Kurs nicht vorkommt. Die Seite zeigt in Go ein Feld im Verbund `konto`.
- `uint` in Solidity ist eine vorzeichenlose Ganzzahl mit 256 Bit.
- `require` heißt „verlange“: Ist die Bedingung falsch, bricht der Aufruf ab, und nichts wird geändert. `msg.sender` ist das Konto, das die Funktion aufruft. `onlyOwner` heißt: Nur der Eigentümer, Tether, darf die Funktion aufrufen. `_;` steht für den Rumpf der Funktion.
- Unter „Was du gefunden hast“ stehen vier Punkte: Sperrliste, Eigentümer, Prüfung beim Überweisen, und die Funktion `destroyBlackFunds`, die das Guthaben eines gesperrten Kontos auf 0 setzt.
- Das Programm von USDT prüft beim Überweisen nur, ob der Absender gesperrt ist. Ein gesperrtes Konto kann weiter USDT empfangen.

### Handgriffe

- Eine der fünf Zeilen wählen.
- Den Link zu Etherscan öffnen.
- Die „Auflösung“ aufklappen.

### Abschnitte des Kurstexts

- 4.3.6
- 4.5.2
- 5.5.1
- 7.1

## Schritt 3: Wer hat die Kontrolle?

Reiter: kontrolle

### Aufbau der Darstellung

- Oben steht ein blauer Kasten mit einer Frage und einer aufklappbaren „Auflösung“: Tether kann eine Zahlung in USDT verhindern.
- Unter „Die Sperrliste wird benutzt“: Tether gibt an, bis September 2026 USDT im Wert von mehr als 5 Milliarden US-Dollar gesperrt zu haben, in Zusammenarbeit mit Behörden aus 70 Ländern. Die Firma BlockSec zählte im Juli 2026 9.597 gesperrte Konten, die Seite schreibt „fast 10.000“.
- Unter „Der Staat greift über Tether zu“: Ein Staat kann das Programm nicht ändern, aber Tether verpflichten. Der GENIUS Act der USA vom Juli 2025 verlangt von Herausgebern digitaler Dollar, dass sie auf Anordnung eines Gerichts oder einer Behörde sperren oder vernichten können. Zwei Seiten: Schutz vor gestohlenem Geld, Macht einer einzigen Stelle.
- Unter „Und der digitale Euro?“ vier Punkte, Stand Oktober 2026: Das Gesetz ist noch nicht beschlossen, frühestens 2029 könnte es ihn geben. Er soll nicht auf einer Blockchain laufen, die Zentralbanken des Euroraums führen das Kassenbuch zentral. Er soll kein programmiertes Geld sein. Sperren würden die Banken, wie bei einem Girokonto heute, zum Beispiel bei Sanktionen der EU. Diese Pflicht steht im Entwurf des Gesetzes und gilt für alle Banken gleich. Das Eurosystem selbst kennt die Namen der Nutzer nicht und hat keine Sperrliste wie USDT.
- Unter „Zum Diskutieren“ stehen zwei offene Fragen ohne Auflösung. Die Seite bezieht keine Stellung, auf beide Fragen gibt es Argumente in beide Richtungen.
- Unter „Was du jetzt kannst“: Wer imperative Programme lesen kann, kann nachprüfen, welche Regeln in digitalem Geld stecken.

### Handgriffe

- Die „Auflösung“ aufklappen.
- Den Links zu den Quellen folgen.
- In diesem Schritt lässt sich nichts ausprobieren.

### Abschnitte des Kurstexts

Der Schritt geht über den Kurstext hinaus.
