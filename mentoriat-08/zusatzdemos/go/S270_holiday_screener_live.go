package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

// =============================================================================
// 🔴 LIVE VERSION: Holiday Screener mit echten RSS-Feeds!
// =============================================================================

type ReiseDeal struct {
	ID              string
	Titel           string
	Beschreibung    string
	Quelle          string
	URL             string
	Zielstadt       string
	Preis           float64
	Dauer           int
	Veroeffentlicht time.Time
	Gefunden        time.Time
	Wetter          *WetterInfo
	Score           float64
}

type WetterInfo struct {
	Stadt         string
	Temperatur    float64
	Beschreibung  string
	Prognose7Tage string
}

type DealQuelle struct {
	Name     string
	URL      string
	LiveFeed bool
}

var staedteKoordinaten = map[string][2]float64{
	"Barcelona": {41.39, 2.16}, "Paris": {48.85, 2.35}, "Rom": {41.90, 12.50},
	"Lissabon": {38.72, -9.14}, "Amsterdam": {52.37, 4.90}, "London": {51.51, -0.13},
	"Prag": {50.08, 14.44}, "Budapest": {47.50, 19.04}, "Wien": {48.21, 16.37},
	"Berlin": {52.52, 13.41}, "Kopenhagen": {55.68, 12.57}, "Stockholm": {59.33, 18.07},
	"Dublin": {53.35, -6.26}, "Athen": {37.98, 23.73}, "Istanbul": {41.01, 28.98},
	"Madrid": {40.42, -3.70}, "Mailand": {45.46, 9.19}, "München": {48.14, 11.58},
}

func main() {
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║  🔴 LIVE: Holiday Screener mit echten RSS-Feeds!  🔴   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝\n")

	// SCHRITT 1: LIVE Quellen
	fmt.Println("📡 SCHRITT 1: LIVE-Quellen konfigurieren")
	fmt.Println(strings.Repeat("─", 60))

	quellen := []DealQuelle{
		{
			Name:     "Reddit TravelDeals",
			URL:      "https://www.reddit.com/r/TravelDeals/.rss",
			LiveFeed: true,
		},
		{
			Name:     "HolidayPirates DE",
			URL:      "https://www.holidaypirates.com/feed/DE",
			LiveFeed: true,
		},
		{
			Name:     "Urlaubsguru",
			URL:      "https://www.urlaubsguru.de/feed/",
			LiveFeed: true,
		},
		{
			Name:     "Travelzoo",
			URL:      "https://www.travelzoo.com/blog/feed/",
			LiveFeed: true,
		},
	}

	for i, q := range quellen {
		fmt.Printf("  [%d] %s\n", i+1, q.Name)
		fmt.Printf("      🔴 LIVE: %s\n", q.URL)
	}
	fmt.Println()

	// SCHRITT 2: LIVE Deals laden
	fmt.Println("🔍 SCHRITT 2: LIVE-Deals abrufen...")
	fmt.Println(strings.Repeat("─", 60))

	var alleDeals []ReiseDeal

	for _, quelle := range quellen {
		fmt.Printf("  Lade von %s...", quelle.Name)

		deals, err := ladeDealsVonRSS(quelle)
		if err != nil {
			fmt.Printf(" ✗ Fehler: %v\n", err)
			fmt.Println("  → Verwende Beispieldaten als Fallback")
			deals = ladeBeispielDeals()
		} else {
			fmt.Printf(" ✓ %d LIVE-Deals\n", len(deals))
		}

		alleDeals = append(alleDeals, deals...)
	}

	fmt.Printf("\n  ✓ Gesamt: %d Deals geladen\n\n", len(alleDeals))

	// SCHRITT 3: Parse Deals
	fmt.Println("🔬 SCHRITT 3: Deal-Informationen extrahieren")
	fmt.Println(strings.Repeat("─", 60))

	// Filtere alte Deals (älter als 30 Tage)
	gefiltert := []ReiseDeal{}
	maxAlter := 30 * 24 * time.Hour
	altGezaehlt := 0
	for _, deal := range alleDeals {
		if time.Since(deal.Veroeffentlicht) <= maxAlter {
			gefiltert = append(gefiltert, deal)
		} else {
			altGezaehlt++
		}
	}
	if altGezaehlt > 0 {
		fmt.Printf("  🗑️  %d alte Deals entfernt (>30 Tage)\n", altGezaehlt)
	}
	alleDeals = gefiltert
	fmt.Printf("  ✓ %d aktuelle Deals (≤30 Tage)\n\n", len(alleDeals))

	for i := range alleDeals {
		parseReiseDeal(&alleDeals[i])
	}

	// Debug: Zeige erste 5 Deals mit Datum
	fmt.Println("  📋 Beispiel-Deals (erste 5):")
	for i := 0; i < 5 && i < len(alleDeals); i++ {
		d := alleDeals[i]
		titelKurz := d.Titel
		if len(titelKurz) > 65 {
			titelKurz = titelKurz[:65] + "..."
		}

		// Berechne Alter
		alter := time.Since(d.Veroeffentlicht)
		alterText := ""
		if alter.Hours() < 24 {
			alterText = fmt.Sprintf("vor %dh", int(alter.Hours()))
		} else {
			alterText = fmt.Sprintf("vor %dd", int(alter.Hours()/24))
		}

		fmt.Printf("    [%d] %s\n", i+1, titelKurz)
		fmt.Printf("        🕒 %s | Stadt: %s | Preis: %.0f€ | Tage: %d\n",
			alterText,
			func() string {
				if d.Zielstadt == "" {
					return "❌"
				}
				return d.Zielstadt
			}(),
			d.Preis,
			d.Dauer)
	}

	gefundenMitPreis := 0
	for _, d := range alleDeals {
		if d.Preis > 0 && d.Zielstadt != "" {
			gefundenMitPreis++
		}
	}

	fmt.Printf("\n  ✓ %d Deals mit Preis & Stadt gefunden\n\n", gefundenMitPreis)

	// SCHRITT 4: Wetter laden
	fmt.Println("🌤️  SCHRITT 4: Wetterdaten für Zielorte")
	fmt.Println(strings.Repeat("─", 60))

	staedteSet := make(map[string]bool)
	for _, deal := range alleDeals {
		if deal.Zielstadt != "" {
			staedteSet[deal.Zielstadt] = true
		}
	}

	wetterCache := make(map[string]*WetterInfo)
	for stadt := range staedteSet {
		fmt.Printf("  %s...", stadt)
		wetter, err := holeWetter(stadt)
		if err != nil {
			fmt.Printf(" ✗\n")
		} else {
			wetterCache[stadt] = wetter
			fmt.Printf(" ✓ %.1f°C\n", wetter.Temperatur)
		}
		time.Sleep(400 * time.Millisecond)
	}

	for i := range alleDeals {
		if alleDeals[i].Zielstadt != "" {
			alleDeals[i].Wetter = wetterCache[alleDeals[i].Zielstadt]
		}
	}

	fmt.Println()

	// SCHRITT 5: Bewerten & Sortieren
	fmt.Println("📊 SCHRITT 5: Deals bewerten")
	fmt.Println(strings.Repeat("─", 60))

	for i := range alleDeals {
		if alleDeals[i].Dauer > 0 {
			alleDeals[i].Score = alleDeals[i].Preis / float64(alleDeals[i].Dauer)
		}
	}

	sortiereDealNachScore(alleDeals)
	fmt.Println("  ✓ Nach Preis-Leistung sortiert\n")

	// SCHRITT 6: Beste Deals zeigen
	fmt.Println("🏆 DIE BESTEN REISE-DEALS (LIVE)")
	fmt.Println(strings.Repeat("═", 60))

	zeigeTopDeals(alleDeals, 10)

	// SCHRITT 7: Statistiken
	fmt.Println("📈 STATISTIKEN")
	fmt.Println(strings.Repeat("─", 60))
	zeigeStatistiken(alleDeals)

	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║  ✅ LIVE-Daten erfolgreich geladen und analysiert!     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝\n")
}

// Hilfsfunktionen
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ternary(b bool, yes, no string) string {
	if b {
		return yes
	}
	return no
}

// =============================================================================
// RSS LIVE-FEED
// =============================================================================

func ladeDealsVonRSS(quelle DealQuelle) ([]ReiseDeal, error) {
	parser := gofeed.NewParser()
	client := &http.Client{Timeout: 15 * time.Second}

	req, err := http.NewRequest("GET", quelle.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Educational Demo)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	feed, err := parser.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var deals []ReiseDeal
	for i, item := range feed.Items {
		if i >= 30 {
			break
		}

		deal := ReiseDeal{
			ID:           fmt.Sprintf("rss_%d", i),
			Titel:        item.Title,
			Beschreibung: cleanHTML(item.Description),
			Quelle:       quelle.Name,
			URL:          item.Link,
			Gefunden:     time.Now(),
		}

		if item.PublishedParsed != nil {
			deal.Veroeffentlicht = *item.PublishedParsed
		} else {
			deal.Veroeffentlicht = time.Now()
		}

		deals = append(deals, deal)
	}

	return deals, nil
}

func cleanHTML(s string) string {
	// Entferne HTML-Tags (einfach)
	s = regexp.MustCompile(`<[^>]*>`).ReplaceAllString(s, "")
	if len(s) > 200 {
		s = s[:200] + "..."
	}
	return s
}

func parseReiseDeal(deal *ReiseDeal) {
	text := deal.Titel + " " + deal.Beschreibung

	// Suche Städte (case-insensitive)
	textLower := strings.ToLower(text)
	for stadt := range staedteKoordinaten {
		if strings.Contains(textLower, strings.ToLower(stadt)) {
			deal.Zielstadt = stadt
			break
		}
	}

	// Suche Preis (flexiblere Muster)
	preisPatterns := []string{
		`\$(\d+)`,             // $599
		`(\d+)\s*€`,           // 399€
		`(\d+)\s*EUR`,         // 399 EUR
		`(\d{3,4})(?:\s|$|/)`, // 599 oder 1299 (ohne Symbol)
		`from\s+\$?(\d+)`,     // from $599
		`[\$€](\d+)`,          // $599 oder €599
	}

	for _, pattern := range preisPatterns {
		preisRegex := regexp.MustCompile(pattern)
		if match := preisRegex.FindStringSubmatch(text); len(match) > 1 {
			preis, err := strconv.ParseFloat(match[1], 64)
			if err == nil && preis >= 50 && preis <= 5000 {
				deal.Preis = preis
				break
			}
		}
	}

	// Suche Dauer
	dauerRegex := regexp.MustCompile(`(\d+)\s*(?:days?|nights?|tage?|nächte?)`)
	if match := dauerRegex.FindStringSubmatch(strings.ToLower(text)); len(match) > 1 {
		dauer, _ := strconv.Atoi(match[1])
		deal.Dauer = dauer
	}

	// Fallback-Dauer basierend auf Preis
	if deal.Dauer == 0 {
		if deal.Preis > 1000 {
			deal.Dauer = 7
		} else if deal.Preis > 500 {
			deal.Dauer = 5
		} else {
			deal.Dauer = 3
		}
	}
}

func holeWetter(stadt string) (*WetterInfo, error) {
	coords, ok := staedteKoordinaten[stadt]
	if !ok {
		return nil, fmt.Errorf("Stadt nicht gefunden")
	}

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?"+
			"latitude=%.2f&longitude=%.2f"+
			"&current=temperature_2m,weather_code"+
			"&timezone=auto",
		coords[0], coords[1])

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Current struct {
			Temperature2m float64 `json:"temperature_2m"`
			WeatherCode   int     `json:"weather_code"`
		} `json:"current"`
	}

	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	return &WetterInfo{
		Stadt:        stadt,
		Temperatur:   apiResp.Current.Temperature2m,
		Beschreibung: wetterCodeZuText(apiResp.Current.WeatherCode),
	}, nil
}

func wetterCodeZuText(code int) string {
	codes := map[int]string{
		0: "Sonnig", 1: "Heiter", 2: "Bewölkt", 3: "Bedeckt",
		61: "Regen", 71: "Schnee", 95: "Gewitter",
	}
	if text, ok := codes[code]; ok {
		return text
	}
	return "Variabel"
}

func sortiereDealNachScore(deals []ReiseDeal) {
	n := len(deals)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if deals[j].Score > deals[j+1].Score {
				deals[j], deals[j+1] = deals[j+1], deals[j]
			}
		}
	}
}

func zeigeTopDeals(deals []ReiseDeal, maxAnzahl int) {
	fmt.Println("┌────┬──────────────┬─────────┬──────┬─────────┬──────────────┐")
	fmt.Println("│Rang│  Zielstadt   │  Preis  │ Tage │  €/Tag  │    Wetter    │")
	fmt.Println("├────┼──────────────┼─────────┼──────┼─────────┼──────────────┤")

	gezaehlt := 0
	for i := 0; i < len(deals) && gezaehlt < maxAnzahl; i++ {
		deal := deals[i]

		// Nur Deals mit vollständigen Daten
		if deal.Preis == 0 || deal.Zielstadt == "" || deal.Score == 0 {
			continue
		}

		gezaehlt++

		wetterText := "?"
		if deal.Wetter != nil {
			wetterText = fmt.Sprintf("%s %.0f°C",
				deal.Wetter.Beschreibung, deal.Wetter.Temperatur)
		}

		emoji := ""
		if deal.Score < 50 {
			emoji = "🔥"
		}

		fmt.Printf("│%2d %s│ %-12s │%6.0f€ │  %2d  │ %6.2f€ │ %-12s │\n",
			gezaehlt, emoji,
			deal.Zielstadt,
			deal.Preis,
			deal.Dauer,
			deal.Score,
			wetterText[:min(12, len(wetterText))])
	}

	if gezaehlt == 0 {
		fmt.Println("│                 Keine vollständigen Deals gefunden                  │")
	}

	fmt.Println("└────┴──────────────┴─────────┴──────┴─────────┴──────────────┘\n")
}

func zeigeStatistiken(deals []ReiseDeal) {
	validDeals := 0
	summePreis := 0.0

	for _, deal := range deals {
		if deal.Preis > 0 && deal.Zielstadt != "" {
			validDeals++
			summePreis += deal.Preis
		}
	}

	if validDeals == 0 {
		fmt.Println("  Keine vollständigen Deals gefunden.\n")
		return
	}

	durchschnitt := summePreis / float64(validDeals)

	fmt.Printf("  📊 Gesamt Deals:       %d\n", len(deals))
	fmt.Printf("  ✅ Vollständig:        %d\n", validDeals)
	fmt.Printf("  💰 Ø-Preis:           %.2f€\n", durchschnitt)
	fmt.Println()
}

func ladeBeispielDeals() []ReiseDeal {
	jetzt := time.Now()
	return []ReiseDeal{
		{
			ID:              "ex1",
			Titel:           "Barcelona: 5 Tage inkl. Flug ab 199€",
			Quelle:          "Beispiel",
			Veroeffentlicht: jetzt,
			Gefunden:        jetzt,
		},
		{
			ID:              "ex2",
			Titel:           "Paris: 3 Tage Städtetrip nur 159€",
			Quelle:          "Beispiel",
			Veroeffentlicht: jetzt,
			Gefunden:        jetzt,
		},
	}
}
