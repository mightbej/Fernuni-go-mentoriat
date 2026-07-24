# Mentoriat 08 – Go-Demos und HTML-Lernmaterialien

Dieses Verzeichnis enthält praktische Demonstrationen und interaktive HTML-Lernmaterialien für das Mentoriat 08 im Modul 01603 „Imperative Programmierung“ der FernUniversität in Hagen.

## Struktur

```text
mentoriat-08/
├── s201-232-verbunddatentypen/
│   ├── go/
│   └── html/
├── s233-250-dynamische-datenstrukturen/
│   ├── go/
│   └── html/
├── zusatzdemos/
│   ├── go/
│   └── html/
├── prompts/
└── README.md
```

Die Gliederung folgt zunächst den Seitenbereichen des Studienbriefs und anschließend dem jeweiligen Thema.

## Voraussetzungen

* GitHub Codespace oder lokale Go-Installation
* aktuelle stabile Go-Version
* Internetverbindung für die Live-Demos

Go-Version prüfen:

```bash
go version
```

Die Moduldefinitionen `go.mod` und `go.sum` liegen im Hauptverzeichnis des Repositorys. Falls Abhängigkeiten fehlen:

```bash
go mod download
```

`go mod init` ist nicht erneut erforderlich.

## Go-Programme ausführen

Alle Befehle werden im Hauptverzeichnis des Repositorys ausgeführt.

Jede Datei enthält ein eigenständiges `main`-Programm. Deshalb immer eine konkrete Datei starten und nicht `go run .` verwenden.

### S201–232: Verbunddatentypen

```bash
go run mentoriat-08/s201-232-verbunddatentypen/go/S207_struct_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S212_verbund_felder_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S214_zeiger_auf_verbund_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S216_verbund_vergleich_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S219_verbund_parameter_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S220_hierarchische_verbunde_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S224_rekursive_verbunde_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S225_anonyme_felder_demo.go
go run mentoriat-08/s201-232-verbunddatentypen/go/S230_binaere_suche_demo.go
```

Behandelte Themen:

* Definition und Initialisierung von `struct`
* Feldzugriff und Feldänderung
* Zeiger auf Verbunde
* Vergleich von Verbundwerten
* Wert- und Zeigerparameter
* geschachtelte und rekursive Verbunde
* anonyme beziehungsweise eingebettete Felder
* binäre Suche in sortierten Verbund-Arrays

### S233–250: Dynamische Datenstrukturen und Operationen

```bash
go run mentoriat-08/s233-250-dynamische-datenstrukturen/go/S234_bubblesort_demo.go
go run mentoriat-08/s233-250-dynamische-datenstrukturen/go/S240_array_operationen_demo.go
```

Behandelte Themen:

* Bubblesort auf Verbund-Arrays
* schrittweise Darstellung eines Sortierverfahrens
* Duplikate, Umkehrung sowie Min-/Max-Operationen

### Zusatzdemos

```bash
go run mentoriat-08/zusatzdemos/go/S250_aktien_app_demo.go
go run mentoriat-08/zusatzdemos/go/S260_wetter_app_demo.go
go run mentoriat-08/zusatzdemos/go/S270_holiday_screener_live.go
```

Die Zusatzdemos verwenden externe Live-Daten:

* S250: Aktienkurse und Portfolio-Berechnungen
* S260: aktuelle Wetterdaten und Prognosen
* S270: RSS-Reiseangebote, Wetteranreicherung und Bewertung

Externe Dienste können zeitweise nicht erreichbar sein oder Zugriffe ablehnen. Die Ergebnisse können sich zwischen zwei Ausführungen verändern.

## HTML-Lernmaterialien anzeigen

GitHub zeigt HTML-Dateien nur als Quellcode. Für die interaktive Darstellung im Codespace wird ein einfacher Webserver gestartet:

```bash
python3 -m http.server 8000 --directory mentoriat-08
```

Anschließend im Codespace den Reiter **PORTS** öffnen und Port `8000` mit **Open in Browser** aufrufen.

Die HTML-Dateien liegen in den jeweiligen Themenordnern:

```text
s201-232-verbunddatentypen/html/
s233-250-dynamische-datenstrukturen/html/
zusatzdemos/html/
```

Zum Beenden des Webservers im Terminal `Strg + C` drücken.

## HTML-Seiten für Studenten freigeben

Weitergeleitete Ports sind standardmäßig privat.

Für eine zeitlich begrenzte Freigabe:

1. Im Codespace den Reiter **PORTS** öffnen.
2. Port `8000` mit der rechten Maustaste anklicken.
3. **Port Visibility → Public** wählen.
4. Die Portadresse kopieren und über Teams weitergeben.
5. Nach dem Mentoriat den Server beenden und den Port wieder auf **Private** stellen.

Der Codespace und der Webserver müssen während der Nutzung aktiv bleiben.

## Arbeitsablauf

1. Repository über GitHub Codespaces öffnen.
2. Dateien im Codespace bearbeiten oder durch Copilot erzeugen lassen.
3. Go-Programme und HTML-Seiten testen.
4. Änderungen committen.
5. **Änderungen synchronisieren** ausführen.
6. Auf GitHub kontrollieren, ob der neue Stand angekommen ist.

Die normale GitHub-Seite dient als dauerhafte Ablage und zur Kontrolle. Bearbeitung und Tests erfolgen im Codespace.
