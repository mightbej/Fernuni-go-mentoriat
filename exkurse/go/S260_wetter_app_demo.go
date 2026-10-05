package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// =============================================================================
// DEMONSTRATION: Wetter-App mit Live-Daten (Praktische Anwendung)
// =============================================================================
//
// Dieses Programm demonstriert Verbünde in einer Wetter-Anwendung:
//
// 1. Live-Wetterdaten von OpenMeteo API abrufen
// 2. Aktuelle Wetterlage für verschiedene Städte
// 3. 7-Tage-Wettervorhersage
// 4. Vergleich mehrerer Standorte
// 5. Hierarchische Verbünde (Stadt -> Wetter -> Prognose)
// 6. Wetterstatistiken berechnen
//
// WICHTIG: Benötigt Internet-Verbindung!
// API: OpenMeteo (kostenlos, kein API-Key erforderlich)
//
// Ausführen mit: go run S260_wetter_app_demo.go
//
// =============================================================================

// =============================================================================
// Verbund-Definitionen
// =============================================================================

// Koordinaten: Geografische Position
type Koordinaten struct {
	Breitengrad float64 // Latitude
	Laengengrad float64 // Longitude
}

// Stadt: Repräsentiert einen Ort
type Stadt struct {
	Name        string
	Land        string
	Koordinaten Koordinaten
}

// AktuellesWetter: Momentane Wetterlage
type AktuellesWetter struct {
	Zeitpunkt           time.Time
	Temperatur          float64 // in °C
	Gefuehlt            float64 // Gefühlte Temperatur
	Windgeschwindigkeit float64 // in km/h
	Windrichtung        int     // in Grad
	Niederschlag        float64 // in mm
	Luftfeuchtigkeit    int     // in %
	WetterCode          int     // Wetterzustand-Code
	Beschreibung        string  // Textbeschreibung
}

// TagesVorhersage: Wettervorhersage für einen Tag
type TagesVorhersage struct {
	Datum        string
	TempMax      float64 // Höchsttemperatur
	TempMin      float64 // Tiefsttemperatur
	Niederschlag float64 // Niederschlagsmenge in mm
	WetterCode   int
	Beschreibung string
}

// WochenVorhersage: 7-Tage-Prognose
type WochenVorhersage struct {
	Tage []TagesVorhersage
}

// WetterBericht: Vollständiger Wetterbericht für eine Stadt
type WetterBericht struct {
	Stadt      Stadt
	Aktuell    AktuellesWetter
	Vorhersage WochenVorhersage
	Abgerufen  time.Time
}

// =============================================================================
// API Response Strukturen (für JSON-Parsing)
// =============================================================================

type openMeteoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   struct {
		Time             string  `json:"time"`
		Temperature2m    float64 `json:"temperature_2m"`
		Windspeed10m     float64 `json:"windspeed_10m"`
		Winddirection10m int     `json:"winddirection_10m"`
		WeatherCode      int     `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		Time             []string  `json:"time"`
		Temperature2mMax []float64 `json:"temperature_2m_max"`
		Temperature2mMin []float64 `json:"temperature_2m_min"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
		WeatherCode      []int     `json:"weather_code"`
	} `json:"daily"`
}

// =============================================================================
// HAUPTPROGRAMM
// =============================================================================

func main() {
	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║      WETTER-APP MIT LIVE-DATEN - Verbünde in Action      ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 1. Beliebte Städte definieren
	// -------------------------------------------------------------------------

	fmt.Println("1. Städte für Wetterabfrage:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	staedte := []Stadt{
		{"Berlin", "Deutschland", Koordinaten{52.52, 13.41}},
		{"München", "Deutschland", Koordinaten{48.14, 11.58}},
		{"Hamburg", "Deutschland", Koordinaten{53.55, 10.00}},
		{"Frankfurt", "Deutschland", Koordinaten{50.11, 8.68}},
		{"Paris", "Frankreich", Koordinaten{48.85, 2.35}},
		{"London", "UK", Koordinaten{51.51, -0.13}},
	}

	for i, stadt := range staedte {
		fmt.Printf("   [%d] %-15s %s (%.2f°N, %.2f°E)\n",
			i+1, stadt.Name, stadt.Land,
			stadt.Koordinaten.Breitengrad,
			stadt.Koordinaten.Laengengrad)
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. Aktuelles Wetter für erste Stadt abrufen
	// -------------------------------------------------------------------------

	fmt.Println("2. Aktuelles Wetter für Berlin:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   Rufe Wetterdaten von OpenMeteo ab...")
	berlinWetter, err := holeWetterBericht(staedte[0])
	if err != nil {
		fmt.Printf("   ✗ Fehler: %v\n", err)
		fmt.Println()
	} else {
		fmt.Println("   ✓ Daten erfolgreich abgerufen!")
		fmt.Println()
		zeigeAktuellesWetter(berlinWetter)
	}

	// -------------------------------------------------------------------------
	// 3. 7-Tage-Vorhersage
	// -------------------------------------------------------------------------

	if err == nil {
		fmt.Println("3. 7-Tage-Wettervorhersage für Berlin:")
		fmt.Println("   " + strings(60, "-"))
		fmt.Println()
		zeigeWochenVorhersage(berlinWetter)
	}

	// -------------------------------------------------------------------------
	// 4. Wetter für mehrere Städte vergleichen
	// -------------------------------------------------------------------------

	fmt.Println("4. Wetter-Vergleich mehrerer Städte:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	var wetterBerichte []WetterBericht

	for i, stadt := range staedte {
		if i >= 5 {
			break // Nur erste 3 Städte für schnellere Demo
		}
		fmt.Printf("   Lade Wetter für %s...", stadt.Name)
		bericht, err := holeWetterBericht(stadt)
		if err != nil {
			fmt.Printf(" ✗\n")
		} else {
			fmt.Printf(" ✓\n")
			wetterBerichte = append(wetterBerichte, bericht)
		}
		time.Sleep(500 * time.Millisecond) // Pause zwischen Requests
	}
	fmt.Println()

	if len(wetterBerichte) > 0 {
		zeigeStadtVergleich(wetterBerichte)
	}

	// -------------------------------------------------------------------------
	// 5. Wetterstatistiken
	// -------------------------------------------------------------------------

	if len(wetterBerichte) > 0 {
		fmt.Println("5. Wetterstatistiken:")
		fmt.Println("   " + strings(60, "-"))
		fmt.Println()

		waermsteStadt := findeWaermsteStadt(wetterBerichte)
		kaeltesteStadt := findeKaeltesteStadt(wetterBerichte)
		windigsteStadt := findeWindigsteStadt(wetterBerichte)

		fmt.Printf("   🌡️  Wärmste Stadt:   %s (%.1f°C)\n",
			waermsteStadt.Stadt.Name,
			waermsteStadt.Aktuell.Temperatur)

		fmt.Printf("   ❄️  Kälteste Stadt:  %s (%.1f°C)\n",
			kaeltesteStadt.Stadt.Name,
			kaeltesteStadt.Aktuell.Temperatur)

		fmt.Printf("   💨 Windigste Stadt: %s (%.1f km/h)\n",
			windigsteStadt.Stadt.Name,
			windigsteStadt.Aktuell.Windgeschwindigkeit)
		fmt.Println()

		durchschnittTemp := berechneDurchschnittsTemperatur(wetterBerichte)
		fmt.Printf("   📊 Durchschnittstemperatur: %.1f°C\n", durchschnittTemp)
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 6. Wochenstatistik für eine Stadt
	// -------------------------------------------------------------------------

	if err == nil {
		fmt.Println("6. Wochenstatistik für Berlin:")
		fmt.Println("   " + strings(60, "-"))
		fmt.Println()

		maxTemp := findeMaxTemperaturWoche(berlinWetter)
		minTemp := findeMinTemperaturWoche(berlinWetter)
		gesamtNiederschlag := berechneGesamtNiederschlag(berlinWetter)

		fmt.Printf("   Höchste Temperatur diese Woche:  %.1f°C\n", maxTemp)
		fmt.Printf("   Niedrigste Temperatur diese Woche: %.1f°C\n", minTemp)
		fmt.Printf("   Temperaturspanne:                  %.1f°C\n", maxTemp-minTemp)
		fmt.Printf("   Gesamtniederschlag:                %.1f mm\n", gesamtNiederschlag)
		fmt.Println()

		regenTage := zaehleRegenTage(berlinWetter)
		fmt.Printf("   Regentage (> 0.5mm): %d von %d Tagen\n",
			regenTage, len(berlinWetter.Vorhersage.Tage))
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 7. Was zeigt dieses Programm?
	// -------------------------------------------------------------------------

	fmt.Println("7. Was haben wir demonstriert?")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   Verbund-Konzepte:")
	fmt.Println("   ✓ Einfache Verbünde (Koordinaten, Stadt)")
	fmt.Println("   ✓ Hierarchische Verbünde (WetterBericht -> Stadt -> Koordinaten)")
	fmt.Println("   ✓ Verbund-Arrays ([]TagesVorhersage, []WetterBericht)")
	fmt.Println("   ✓ Geschachtelte Verbünde (openMeteoResponse)")
	fmt.Println("   ✓ Zeit-Verbünde (time.Time)")
	fmt.Println()

	fmt.Println("   Praktische Anwendung:")
	fmt.Println("   ✓ HTTP-Requests an Wetter-API")
	fmt.Println("   ✓ JSON in Go-Verbünde umwandeln")
	fmt.Println("   ✓ Berechnungen auf mehreren Verbünden")
	fmt.Println("   ✓ Echte Live-Wetterdaten")
	fmt.Println()

	fmt.Println("   Algorithmen:")
	fmt.Println("   ✓ API-Daten holen und parsen")
	fmt.Println("   ✓ Minimum/Maximum finden (wärmste/kälteste Stadt)")
	fmt.Println("   ✓ Statistiken berechnen (Durchschnitt, Summen)")
	fmt.Println("   ✓ Daten vergleichen und auswerten")
	fmt.Println()

	// -------------------------------------------------------------------------
	// 8. Erweiterungsmöglichkeiten
	// -------------------------------------------------------------------------

	fmt.Println("8. Mögliche Erweiterungen für Studenten:")
	fmt.Println("   " + strings(60, "-"))
	fmt.Println()

	fmt.Println("   💡 Weitere Features:")
	fmt.Println("   • Stündliche Vorhersage (24h)")
	fmt.Println("   • Wetterwarnstufen bei Extremwetter")
	fmt.Println("   • Beste Reisezeit ermitteln")
	fmt.Println("   • UV-Index, Luftqualität")
	fmt.Println("   • Wetter-Icons/Emojis anzeigen")
	fmt.Println("   • Sonnenaufgang/-untergang")
	fmt.Println("   • Favoriten-Städte speichern")
	fmt.Println("   • Automatische Standort-Erkennung")
	fmt.Println()

	fmt.Println("╔═══════════════════════════════════════════════════════════╗")
	fmt.Println("║              ENDE DER DEMONSTRATION                       ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════╝")
}

// =============================================================================
// API-Funktionen
// =============================================================================

// holeWetterBericht: Holt vollständigen Wetterbericht von OpenMeteo
func holeWetterBericht(stadt Stadt) (WetterBericht, error) {
	// OpenMeteo API URL mit allen benötigten Parametern
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f"+
			"&current=temperature_2m,windspeed_10m,winddirection_10m,weather_code"+
			"&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,weather_code"+
			"&timezone=auto",
		stadt.Koordinaten.Breitengrad,
		stadt.Koordinaten.Laengengrad)

	// HTTP Request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return WetterBericht{}, fmt.Errorf("HTTP-Fehler: %v", err)
	}
	defer resp.Body.Close()

	// Response lesen
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WetterBericht{}, fmt.Errorf("Fehler beim Lesen: %v", err)
	}

	// JSON parsen
	var apiResponse openMeteoResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return WetterBericht{}, fmt.Errorf("JSON-Parse-Fehler: %v", err)
	}

	// Parse aktuelle Zeit
	zeitpunkt, _ := time.Parse("2006-01-02T15:04", apiResponse.Current.Time)

	// Erstelle AktuellesWetter
	aktuell := AktuellesWetter{
		Zeitpunkt:           zeitpunkt,
		Temperatur:          apiResponse.Current.Temperature2m,
		Windgeschwindigkeit: apiResponse.Current.Windspeed10m,
		Windrichtung:        apiResponse.Current.Winddirection10m,
		WetterCode:          apiResponse.Current.WeatherCode,
		Beschreibung:        wetterCodeZuText(apiResponse.Current.WeatherCode),
	}

	// Erstelle Vorhersage
	var tage []TagesVorhersage
	for i := 0; i < len(apiResponse.Daily.Time) && i < 7; i++ {
		tag := TagesVorhersage{
			Datum:        apiResponse.Daily.Time[i],
			TempMax:      apiResponse.Daily.Temperature2mMax[i],
			TempMin:      apiResponse.Daily.Temperature2mMin[i],
			Niederschlag: apiResponse.Daily.PrecipitationSum[i],
			WetterCode:   apiResponse.Daily.WeatherCode[i],
			Beschreibung: wetterCodeZuText(apiResponse.Daily.WeatherCode[i]),
		}
		tage = append(tage, tag)
	}

	vorhersage := WochenVorhersage{Tage: tage}

	// Erstelle vollständigen Bericht
	bericht := WetterBericht{
		Stadt:      stadt,
		Aktuell:    aktuell,
		Vorhersage: vorhersage,
		Abgerufen:  time.Now(),
	}

	return bericht, nil
}

// wetterCodeZuText: Konvertiert WMO Wetter-Code in deutsche Beschreibung
func wetterCodeZuText(code int) string {
	beschreibungen := map[int]string{
		0:  "Klar",
		1:  "Überwiegend klar",
		2:  "Teilweise bewölkt",
		3:  "Bewölkt",
		45: "Neblig",
		48: "Gefrierender Nebel",
		51: "Leichter Nieselregen",
		53: "Nieselregen",
		55: "Starker Nieselregen",
		61: "Leichter Regen",
		63: "Regen",
		65: "Starker Regen",
		71: "Leichter Schneefall",
		73: "Schneefall",
		75: "Starker Schneefall",
		77: "Schnee-Schauer",
		80: "Leichte Regenschauer",
		81: "Regenschauer",
		82: "Starke Regenschauer",
		85: "Schnee-Schauer",
		86: "Starke Schnee-Schauer",
		95: "Gewitter",
		96: "Gewitter mit Hagel",
		99: "Starkes Gewitter",
	}

	if beschreibung, ok := beschreibungen[code]; ok {
		return beschreibung
	}
	return fmt.Sprintf("Unbekannt (%d)", code)
}

// =============================================================================
// Statistik-Funktionen
// =============================================================================

// findeWaermsteStadt: Findet Stadt mit höchster aktueller Temperatur
func findeWaermsteStadt(berichte []WetterBericht) WetterBericht {
	if len(berichte) == 0 {
		return WetterBericht{}
	}

	waermste := berichte[0]
	for _, bericht := range berichte {
		if bericht.Aktuell.Temperatur > waermste.Aktuell.Temperatur {
			waermste = bericht
		}
	}
	return waermste
}

// findeKaeltesteStadt: Findet Stadt mit niedrigster aktueller Temperatur
func findeKaeltesteStadt(berichte []WetterBericht) WetterBericht {
	if len(berichte) == 0 {
		return WetterBericht{}
	}

	kaelteste := berichte[0]
	for _, bericht := range berichte {
		if bericht.Aktuell.Temperatur < kaelteste.Aktuell.Temperatur {
			kaelteste = bericht
		}
	}
	return kaelteste
}

// findeWindigsteStadt: Findet Stadt mit höchster Windgeschwindigkeit
func findeWindigsteStadt(berichte []WetterBericht) WetterBericht {
	if len(berichte) == 0 {
		return WetterBericht{}
	}

	windigste := berichte[0]
	for _, bericht := range berichte {
		if bericht.Aktuell.Windgeschwindigkeit > windigste.Aktuell.Windgeschwindigkeit {
			windigste = bericht
		}
	}
	return windigste
}

// berechneDurchschnittsTemperatur: Berechnet Durchschnittstemperatur aller Städte
func berechneDurchschnittsTemperatur(berichte []WetterBericht) float64 {
	if len(berichte) == 0 {
		return 0
	}

	summe := 0.0
	for _, bericht := range berichte {
		summe += bericht.Aktuell.Temperatur
	}
	return summe / float64(len(berichte))
}

// findeMaxTemperaturWoche: Findet höchste Temperatur der Woche
func findeMaxTemperaturWoche(bericht WetterBericht) float64 {
	max := bericht.Vorhersage.Tage[0].TempMax
	for _, tag := range bericht.Vorhersage.Tage {
		if tag.TempMax > max {
			max = tag.TempMax
		}
	}
	return max
}

// findeMinTemperaturWoche: Findet niedrigste Temperatur der Woche
func findeMinTemperaturWoche(bericht WetterBericht) float64 {
	min := bericht.Vorhersage.Tage[0].TempMin
	for _, tag := range bericht.Vorhersage.Tage {
		if tag.TempMin < min {
			min = tag.TempMin
		}
	}
	return min
}

// berechneGesamtNiederschlag: Summiert Niederschlag der Woche
func berechneGesamtNiederschlag(bericht WetterBericht) float64 {
	summe := 0.0
	for _, tag := range bericht.Vorhersage.Tage {
		summe += tag.Niederschlag
	}
	return summe
}

// zaehleRegenTage: Zählt Tage mit Niederschlag > 0.5mm
func zaehleRegenTage(bericht WetterBericht) int {
	anzahl := 0
	for _, tag := range bericht.Vorhersage.Tage {
		if tag.Niederschlag > 0.5 {
			anzahl++
		}
	}
	return anzahl
}

// =============================================================================
// Ausgabe-Funktionen
// =============================================================================

// zeigeAktuellesWetter: Zeigt aktuelle Wetterlage
func zeigeAktuellesWetter(bericht WetterBericht) {
	fmt.Println("   ┌───────────────────────────────────────────────────────┐")
	fmt.Printf("   │ %s, %s                          \n", bericht.Stadt.Name, bericht.Stadt.Land)
	fmt.Printf("   │ Aktualisiert: %s                  \n", bericht.Aktuell.Zeitpunkt.Format("15:04"))
	fmt.Println("   ├───────────────────────────────────────────────────────┤")
	fmt.Printf("   │ 🌡️  Temperatur:       %6.1f°C                      │\n", bericht.Aktuell.Temperatur)
	fmt.Printf("   │ ☁️  Wetter:           %-30s │\n", bericht.Aktuell.Beschreibung)
	fmt.Printf("   │ 💨 Wind:             %6.1f km/h (%d°)              │\n",
		bericht.Aktuell.Windgeschwindigkeit, bericht.Aktuell.Windrichtung)
	fmt.Println("   └───────────────────────────────────────────────────────┘")
	fmt.Println()
}

// zeigeWochenVorhersage: Zeigt 7-Tage-Vorhersage
func zeigeWochenVorhersage(bericht WetterBericht) {
	fmt.Println("   ┌────────────┬─────────────┬───────────┬──────────────┐")
	fmt.Println("   │   Datum    │   Wetter    │   Temp.   │ Niederschlag │")
	fmt.Println("   ├────────────┼─────────────┼───────────┼──────────────┤")

	for _, tag := range bericht.Vorhersage.Tage {
		// Datum formatieren (nur Tag/Monat)
		t, _ := time.Parse("2006-01-02", tag.Datum)
		datumStr := t.Format("Mo, 02.01")

		fmt.Printf("   │ %-10s │ %-11s │ %2.0f - %2.0f°C │   %5.1f mm   │\n",
			datumStr,
			tag.Beschreibung[:min(11, len(tag.Beschreibung))],
			tag.TempMin,
			tag.TempMax,
			tag.Niederschlag)
	}

	fmt.Println("   └────────────┴─────────────┴───────────┴──────────────┘")
	fmt.Println()
}

// zeigeStadtVergleich: Vergleicht Wetter mehrerer Städte
func zeigeStadtVergleich(berichte []WetterBericht) {
	fmt.Println("   ┌──────────────┬────────────┬─────────────────┬──────────┐")
	fmt.Println("   │     Stadt    │   Temp.    │     Wetter      │   Wind   │")
	fmt.Println("   ├──────────────┼────────────┼─────────────────┼──────────┤")

	for _, bericht := range berichte {
		fmt.Printf("   │ %-12s │   %5.1f°C │ %-15s │ %5.1f km/h│\n",
			bericht.Stadt.Name,
			bericht.Aktuell.Temperatur,
			bericht.Aktuell.Beschreibung[:min(15, len(bericht.Aktuell.Beschreibung))],
			bericht.Aktuell.Windgeschwindigkeit)
	}

	fmt.Println("   └──────────────┴────────────┴─────────────────┴──────────┘")
	fmt.Println()
}

// =============================================================================
// Hilfsfunktionen
// =============================================================================

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
