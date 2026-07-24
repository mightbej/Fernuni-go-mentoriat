package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// =============================================================================
// DEMONSTRATION: Aktien-Portfolio mit Live-Daten (Praktische Anwendung)
// =============================================================================
//
// Dieses Programm demonstriert Verbünde in einer echten Anwendung:
//
// 1. Live-Aktienkurse von Yahoo Finance abrufen
// 2. Portfolio mit mehreren Aktien verwalten
// 3. Gewinn/Verlust in Echtzeit berechnen
// 4. Hierarchische Verbünde (Portfolio -> Positionen -> Aktie)
// 5. JSON-Daten in Verbünde umwandeln
//
// WICHTIG: Benötigt Internet-Verbindung!
//
// Ausführen mit: go run S250_aktien_app_demo.go
//
// =============================================================================

// =============================================================================
// Verbund-Definitionen
// =============================================================================

// Aktie: Repräsentiert eine Aktie mit aktuellen Kursdaten
type Aktie struct {
	Symbol           string  // z.B. "AAPL", "MSFT", "TSLA"
	Name             string  // Firmenname
	AktuellerPreis   float64 // Aktueller Kurs in USD
	VorherigerPreis  float64 // Schlusskurs vom Vortag
	Aenderung        float64 // Absolute Änderung
	AenderungProzent float64 // Prozentuale Änderung
	Volumen          int64   // Handelsvolumen
	MarktKap         float64 // Marktkapitalisierung
}

// PortfolioPosition: Eine Position im Portfolio
type PortfolioPosition struct {
	Aktie         Aktie   // Die Aktie (hierarchischer Verbund!)
	Anzahl        int     // Wie viele Aktien besitze ich?
	Kaufpreis     float64 // Zu welchem Preis gekauft?
	Kaufdatum     string  // Wann gekauft?
	Investiert    float64 // Gesamtinvestition = Anzahl * Kaufpreis
	AktuellerWert float64 // Aktueller Wert = Anzahl * AktuellerPreis
	GewinnVerlust float64 // Gewinn/Verlust in USD
	GewinnProzent float64 // Gewinn/Verlust in %
}

// Portfolio: Sammlung aller Aktienpositionen
type Portfolio struct {
	Name                string              // Name des Portfolios
	Positionen          []PortfolioPosition // Array von Positionen
	GesamtInvestiert    float64             // Gesamte Investition
	GesamtAktuellerWert float64             // Aktueller Gesamtwert
	GesamtGewinnVerlust float64             // Gesamt Gewinn/Verlust
	GesamtGewinnProzent float64             // Gesamt Gewinn/Verlust in %
}

// yahooResponse: Struktur für Yahoo Finance API Response
// (Zeigt komplexe geschachtelte Verbünde)
type yahooResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Symbol             string  `json:"symbol"`
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				PreviousClose      float64 `json:"previousClose"`
				ChartPreviousClose float64 `json:"chartPreviousClose"`
			} `json:"meta"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║   AKTIEN-PORTFOLIO MIT LIVE-DATEN - Verbünde in Action   ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 1. Einzelne Aktie abrufen und anzeigen
	// -------------------------------------------------------------------------

	fmt.Println("1. Einzelne Aktie abrufen (Apple):")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   Rufe Daten von Yahoo Finance ab...")
	apple, err := holeAktienkurs("AAPL")
	if err != nil {
		fmt.Printf("   ✗ Fehler beim Abrufen: %v\n", err)
		fmt.Println("   (Prüfe Internet-Verbindung oder versuche es später)")
		fmt.Println()
	} else {
		zeigeAktie(apple)
	}

	// -------------------------------------------------------------------------
	// 2. Mehrere Aktien abrufen
	// -------------------------------------------------------------------------

	fmt.Println("2. Mehrere beliebte Tech-Aktien abrufen:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	symbole := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA"}
	var aktien []Aktie

	for _, symbol := range symbole {
		fmt.Printf("   Lade %s...", symbol)
		aktie, err := holeAktienkurs(symbol)
		if err != nil {
			fmt.Printf(" ✗ Fehler\n")
		} else {
			fmt.Printf(" ✓\n")
			aktien = append(aktien, aktie)
		}
		time.Sleep(500 * time.Millisecond) // Höfliche Pause zwischen Requests
	}
	fmt.Println()

	if len(aktien) > 0 {
		fmt.Println("   Aktuelle Kurse:")
		fmt.Println()
		fmt.Println("   ┌──────┬────────────┬───────────┬──────────┐")
		fmt.Println("   │Symbol│ Kurs (USD) │ Änderung  │  Trend   │")
		fmt.Println("   ├──────┼────────────┼───────────┼──────────┤")

		for _, a := range aktien {
			trend := "→"
			if a.AenderungProzent > 0 {
				trend = "↑"
			} else if a.AenderungProzent < 0 {
				trend = "↓"
			}

			fmt.Printf("   │ %-5s│ %10.2f │ %+8.2f%% │    %s     │\n",
				a.Symbol, a.AktuellerPreis, a.AenderungProzent, trend)
		}

		fmt.Println("   └──────┴────────────┴───────────┴──────────┘")
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 3. Portfolio erstellen und verwalten
	// -------------------------------------------------------------------------

	fmt.Println("3. Portfolio erstellen:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	// Erstelle Portfolio mit Beispiel-Käufen
	portfolio := Portfolio{
		Name:       "Mein Tech-Portfolio",
		Positionen: []PortfolioPosition{},
	}

	// Hinweis: In echter Anwendung würden wir hier aktuelle Kurse holen
	// Für Demo verwenden wir die bereits geladenen Aktien
	fmt.Println("   Füge Positionen zum Portfolio hinzu:")
	fmt.Println()

	if len(aktien) >= 3 {
		// Position 1: Apple
		position1 := erstellePosition(aktien[0], 10, 150.00, "2024-01-15")
		portfolio.Positionen = append(portfolio.Positionen, position1)
		fmt.Printf("   ✓ %d x %s gekauft zu $%.2f am %s\n",
			position1.Anzahl, position1.Aktie.Symbol,
			position1.Kaufpreis, position1.Kaufdatum)

		// Position 2: Microsoft
		position2 := erstellePosition(aktien[1], 5, 350.00, "2024-02-20")
		portfolio.Positionen = append(portfolio.Positionen, position2)
		fmt.Printf("   ✓ %d x %s gekauft zu $%.2f am %s\n",
			position2.Anzahl, position2.Aktie.Symbol,
			position2.Kaufpreis, position2.Kaufdatum)

		// Position 3: Tesla
		if len(aktien) >= 5 {
			position3 := erstellePosition(aktien[4], 15, 200.00, "2024-03-10")
			portfolio.Positionen = append(portfolio.Positionen, position3)
			fmt.Printf("   ✓ %d x %s gekauft zu $%.2f am %s\n",
				position3.Anzahl, position3.Aktie.Symbol,
				position3.Kaufpreis, position3.Kaufdatum)
		}
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 4. Portfolio-Statistiken berechnen
	// -------------------------------------------------------------------------

	fmt.Println("4. Portfolio-Analyse:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	berechnePortfolioStatistiken(&portfolio)
	zeigePortfolio(portfolio)

	// -------------------------------------------------------------------------
	// 5. Top-Performer und Verlierer
	// -------------------------------------------------------------------------

	if len(portfolio.Positionen) > 0 {
		fmt.Println("5. Performance-Analyse:")
		fmt.Println("   " + strings(60, "-"))
		fmt.Println()

		beste := findeBestePosition(portfolio)
		schlechteste := findeSchlechtestePosition(portfolio)

		fmt.Println("   🏆 BESTE POSITION:")
		fmt.Printf("      %s: +$%.2f (+%.2f%%)\n",
			beste.Aktie.Symbol, beste.GewinnVerlust, beste.GewinnProzent)
		fmt.Println()

		fmt.Println("   📉 SCHLECHTESTE POSITION:")
		fmt.Printf("      %s: $%.2f (%.2f%%)\n",
			schlechteste.Aktie.Symbol,
			schlechteste.GewinnVerlust,
			schlechteste.GewinnProzent)
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 6. Was zeigt dieses Programm?
	// -------------------------------------------------------------------------

	fmt.Println("6. Was haben wir demonstriert?")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   Verbund-Konzepte:")
	fmt.Println("   ✓ Einfache Verbünde (Aktie)")
	fmt.Println("   ✓ Hierarchische Verbünde (Portfolio -> Position -> Aktie)")
	fmt.Println("   ✓ Verbund-Arrays ([]PortfolioPosition)")
	fmt.Println("   ✓ Geschachtelte Verbünde (yahooResponse)")
	fmt.Println()

	fmt.Println("   Praktische Anwendung:")
	fmt.Println("   ✓ HTTP-Requests an externe API")
	fmt.Println("   ✓ JSON in Go-Verbünde umwandeln")
	fmt.Println("   ✓ Berechnungen auf Verbund-Feldern")
	fmt.Println("   ✓ Echte Live-Daten verarbeiten")
	fmt.Println()

	fmt.Println("   Algorithmen:")
	fmt.Println("   ✓ Daten von API holen und parsen")
	fmt.Println("   ✓ Portfolio-Berechnungen")
	fmt.Println("   ✓ Minimum/Maximum finden (beste/schlechteste Position)")
	fmt.Println("   ✓ Aggregationen (Gesamtwerte berechnen)")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 7. Erweiterungsmöglichkeiten
	// -------------------------------------------------------------------------

	fmt.Println("7. Mögliche Erweiterungen für Studenten:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   💡 Weitere Features:")
	fmt.Println("   • Portfolio in Datei speichern/laden (JSON)")
	fmt.Println("   • Mehr Aktiendaten (52-Wochen Hoch/Tief)")
	fmt.Println("   • Historische Kurse abrufen")
	fmt.Println("   • Dividenden berücksichtigen")
	fmt.Println("   • Alarme bei Kursänderungen")
	fmt.Println("   • Sortierung nach verschiedenen Kriterien")
	fmt.Println("   • Suchfunktion nach Symbol/Name")
	fmt.Println("   • Export als CSV/Excel")
	fmt.Println()

	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║              ENDE DER DEMONSTRATION                       ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
}

// =============================================================================
// API-Funktionen
// =============================================================================

// holeAktienkurs: Holt aktuelle Kursdaten von Yahoo Finance
func holeAktienkurs(symbol string) (Aktie, error) {
	// Yahoo Finance API URL
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

	// HTTP Request mit User-Agent (Yahoo mag Bots nicht ohne)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Aktie{}, fmt.Errorf("Request-Fehler: %v", err)
	}

	// User-Agent setzen (wichtig für Yahoo!)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return Aktie{}, fmt.Errorf("HTTP-Fehler: %v", err)
	}
	defer resp.Body.Close()

	// Response lesen
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Aktie{}, fmt.Errorf("Fehler beim Lesen: %v", err)
	}

	// DEBUG: Zeige Anfang der Response
	if len(body) > 0 && body[0] != '{' {
		// Keine JSON-Antwort - zeige ersten Teil
		preview := string(body)
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return Aktie{}, fmt.Errorf("Keine JSON-Antwort. Antwort: %s", preview)
	}

	// JSON parsen
	var apiResponse yahooResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return Aktie{}, fmt.Errorf("JSON-Parse-Fehler: %v (Response: %s)", err, string(body[:min(200, len(body))]))
	}

	// Prüfe ob Daten vorhanden
	if len(apiResponse.Chart.Result) == 0 {
		return Aktie{}, fmt.Errorf("keine Daten für Symbol %s", symbol)
	}

	// Extrahiere Daten
	meta := apiResponse.Chart.Result[0].Meta
	aktuellerPreis := meta.RegularMarketPrice
	vorherigerPreis := meta.PreviousClose

	// Wenn PreviousClose nicht gesetzt, verwende ChartPreviousClose
	if vorherigerPreis == 0 {
		vorherigerPreis = meta.ChartPreviousClose
	}

	// Berechne Änderungen
	aenderung := aktuellerPreis - vorherigerPreis
	aenderungProzent := 0.0
	if vorherigerPreis > 0 {
		aenderungProzent = (aenderung / vorherigerPreis) * 100
	}

	// Erstelle Aktie-Verbund
	aktie := Aktie{
		Symbol:           meta.Symbol,
		Name:             meta.Symbol, // Yahoo gibt Namen nicht immer zurück
		AktuellerPreis:   aktuellerPreis,
		VorherigerPreis:  vorherigerPreis,
		Aenderung:        aenderung,
		AenderungProzent: aenderungProzent,
		Volumen:          0, // Nicht in dieser API-Version
		MarktKap:         0, // Nicht in dieser API-Version
	}

	return aktie, nil
}

// =============================================================================
// Portfolio-Funktionen
// =============================================================================

// erstellePosition: Erstellt eine Portfolio-Position
func erstellePosition(aktie Aktie, anzahl int, kaufpreis float64, kaufdatum string) PortfolioPosition {
	investiert := float64(anzahl) * kaufpreis
	aktuellerWert := float64(anzahl) * aktie.AktuellerPreis
	gewinnVerlust := aktuellerWert - investiert
	gewinnProzent := 0.0
	if investiert > 0 {
		gewinnProzent = (gewinnVerlust / investiert) * 100
	}

	return PortfolioPosition{
		Aktie:         aktie,
		Anzahl:        anzahl,
		Kaufpreis:     kaufpreis,
		Kaufdatum:     kaufdatum,
		Investiert:    investiert,
		AktuellerWert: aktuellerWert,
		GewinnVerlust: gewinnVerlust,
		GewinnProzent: gewinnProzent,
	}
}

// berechnePortfolioStatistiken: Berechnet Gesamtstatistiken
func berechnePortfolioStatistiken(p *Portfolio) {
	p.GesamtInvestiert = 0
	p.GesamtAktuellerWert = 0

	for _, pos := range p.Positionen {
		p.GesamtInvestiert += pos.Investiert
		p.GesamtAktuellerWert += pos.AktuellerWert
	}

	p.GesamtGewinnVerlust = p.GesamtAktuellerWert - p.GesamtInvestiert
	if p.GesamtInvestiert > 0 {
		p.GesamtGewinnProzent = (p.GesamtGewinnVerlust / p.GesamtInvestiert) * 100
	}
}

// findeBestePosition: Findet Position mit höchstem Gewinn
func findeBestePosition(p Portfolio) PortfolioPosition {
	if len(p.Positionen) == 0 {
		return PortfolioPosition{}
	}

	beste := p.Positionen[0]
	for _, pos := range p.Positionen {
		if pos.GewinnProzent > beste.GewinnProzent {
			beste = pos
		}
	}
	return beste
}

// findeSchlechtestePosition: Findet Position mit niedrigstem Gewinn
func findeSchlechtestePosition(p Portfolio) PortfolioPosition {
	if len(p.Positionen) == 0 {
		return PortfolioPosition{}
	}

	schlechteste := p.Positionen[0]
	for _, pos := range p.Positionen {
		if pos.GewinnProzent < schlechteste.GewinnProzent {
			schlechteste = pos
		}
	}
	return schlechteste
}

// =============================================================================
// Ausgabe-Funktionen
// =============================================================================

// zeigeAktie: Zeigt Details einer Aktie
func zeigeAktie(a Aktie) {
	fmt.Println("   ┌─────────────────────────────────────────────┐")
	fmt.Printf("   │ Symbol: %-35s │\n", a.Symbol)
	fmt.Println("   ├─────────────────────────────────────────────┤")
	fmt.Printf("   │ Aktueller Kurs:    $%-20.2f │\n", a.AktuellerPreis)
	fmt.Printf("   │ Vorheriger Kurs:   $%-20.2f │\n", a.VorherigerPreis)
	fmt.Printf("   │ Änderung:          $%-20.2f │\n", a.Aenderung)
	fmt.Printf("   │ Änderung %%:        %+20.2f%% │\n", a.AenderungProzent)
	fmt.Println("   └─────────────────────────────────────────────┘")
	fmt.Println()
}

// zeigePortfolio: Zeigt Portfolio-Übersicht
func zeigePortfolio(p Portfolio) {
	fmt.Printf("   Portfolio: %s\n", p.Name)
	fmt.Println()

	if len(p.Positionen) == 0 {
		fmt.Println("   (Keine Positionen)")
		return
	}

	fmt.Println("   ┌──────┬──────┬──────────┬───────────┬──────────────┐")
	fmt.Println("   │Symbol│Anzahl│ Kaufpreis│ Akt. Preis│  Gewinn/Verl │")
	fmt.Println("   ├──────┼──────┼──────────┼───────────┼──────────────┤")

	for _, pos := range p.Positionen {
		fmt.Printf("   │ %-5s│ %4d │ $%7.2f │  $%7.2f │ $%6.2f %+.1f%% │\n",
			pos.Aktie.Symbol,
			pos.Anzahl,
			pos.Kaufpreis,
			pos.Aktie.AktuellerPreis,
			pos.GewinnVerlust,
			pos.GewinnProzent)
	}

	fmt.Println("   ├──────┴──────┴──────────┴───────────┴──────────────┤")
	fmt.Printf("   │ Investiert:        $%-28.2f │\n", p.GesamtInvestiert)
	fmt.Printf("   │ Aktueller Wert:    $%-28.2f │\n", p.GesamtAktuellerWert)
	fmt.Printf("   │ Gesamt G/V:        $%-18.2f %+.2f%% │\n",
		p.GesamtGewinnVerlust, p.GesamtGewinnProzent)
	fmt.Println("   └──────────────────────────────────────────────────┘")
	fmt.Println()
}

// strings: Hilfsfunktion für wiederholte Zeichen
func strings(n int, s string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// min: Hilfsfunktion für Minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
