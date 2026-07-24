# Go-Demos S207 bis S270

Diese Sammlung enthaelt Demo-Programme fuer Verbunde (structs), Arrays und einfache Live-Anwendungen in Go.

## Voraussetzungen

- Go installiert (empfohlen: aktuelle stabile Version)
- Fuer Live-Demos: Internetverbindung

Version pruefen:

```powershell
go version
```

## Programme ausfuehren

Im Terminal in diesen Ordner wechseln:

```powershell
cd /pfad/zum/projektordner
```

Hinweis: In VS Code ist der Projektordner oft bereits geoeffnet. Dann ist kein `cd` noetig.

Dann jeweils eine Datei gezielt starten:

```powershell
go run S207_struct_demo.go
```

Hinweis: In diesem Ordner gibt es mehrere `main`-Programme. Deshalb nicht `go run .` verwenden, sondern immer die gewuenschte Datei direkt angeben.

## Uebersicht der Demos

- `S207_struct_demo.go`: Grundlagen zu `struct`-Definitionen, Exportregeln, Initialisierung, Funktionen mit Structs.
- `S212_verbund_felder_demo.go`: Feldzugriff per Punktnotation, Aenderungen an Feldern, Kopie vs. Zeiger auf Felder.
- `S214_zeiger_auf_verbund_demo.go`: Zeiger auf Verbunde, Dereferenzierung, Verbund-Literale mit Zeigern.
- `S216_verbund_vergleich_demo.go`: Vergleich von Verbundwerten mit `==`, Wertgleichheit vs. Identitaet.
- `S219_verbund_parameter_demo.go`: Call-by-Value vs. Call-by-Reference bei Verbundparametern inkl. Performance-Aspekt.
- `S220_hierarchische_verbunde_demo.go`: Geschachtelte Verbunde, Wertfelder vs. Zeigerfelder.
- `S224_rekursive_verbunde_demo.go`: Rekursive Verbunde mit Zeigern (z. B. Stammbaum, Listen, Baeume).
- `S225_anonyme_felder_demo.go`: Eingebettete/anonyme Felder (Embedded Structs) und vereinfachter Zugriff.
- `S230_binaere_suche_demo.go`: Binaere Suche in sortierten Verbund-Arrays.
- `S234_bubblesort_demo.go`: Bubblesort auf Verbund-Arrays mit Schritt-fuer-Schritt-Ausgabe.
- `S240_array_operationen_demo.go`: Duplikate, Umkehrung, Min/Max und weitere Array-Operationen auf Verbunden.
- `S250_aktien_app_demo.go`: Live-Aktienkurs-Demo (Yahoo Finance), JSON in Verbunde, Portfolio-Berechnungen.
- `S260_wetter_app_demo.go`: Live-Wetterdaten (Open-Meteo), aktuelle Werte und 7-Tage-Prognose.
- `S270_holiday_screener_live.go`: Live-Holiday-Screener mit RSS-Feeds, Wetteranreicherung und Scoring.

## Alle Startbefehle (Copy/Paste)

```powershell
go run S207_struct_demo.go
go run S212_verbund_felder_demo.go
go run S214_zeiger_auf_verbund_demo.go
go run S216_verbund_vergleich_demo.go
go run S219_verbund_parameter_demo.go
go run S220_hierarchische_verbunde_demo.go
go run S224_rekursive_verbunde_demo.go
go run S225_anonyme_felder_demo.go
go run S230_binaere_suche_demo.go
go run S234_bubblesort_demo.go
go run S240_array_operationen_demo.go
go run S250_aktien_app_demo.go
go run S260_wetter_app_demo.go
go run S270_holiday_screener_live.go
```

## Hinweis zu Abhaengigkeiten (nur Live-Screener S270)

Die Datei `S270_holiday_screener_live.go` verwendet das Paket `github.com/mmcdole/gofeed`.
Falls noch kein Go-Modul im Ordner existiert, einmalig ausfuehren:

```powershell
go mod init s207-270-demos
go get github.com/mmcdole/gofeed@latest
```

Danach laeuft der Start wie oben mit `go run S270_holiday_screener_live.go`.
