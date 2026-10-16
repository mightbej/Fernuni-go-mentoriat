// Wetter-Vergleich mit Daten aus dem Internet
//
// Das Programm holt das Wetter für sechs Städte beim freien Wetterdienst
// Open-Meteo: das Wetter jetzt und die Vorhersage für sieben Tage. Dann sucht
// es die wärmste, die kälteste und die windigste Stadt und zeigt die Woche für
// die erste Stadt.
//
// Was du aus dem Kurs kennst:
//   - Verbünde, die ineinander stecken: WetterBericht -> Stadt -> Koordinaten,
//     Abschnitt 8.3
//   - Arrays mit fester Länge, auch als Feld eines Verbunds, Kapitel 6
//   - Minimum und Maximum in einer Schleife suchen, Kapitel 5 und 6
//   - eine switch-Anweisung, Abschnitt 5.5.2
//   - Funktionen mit einem Wert und einem Fehler als Rückgabe, Abschnitt 7.1.2
//
// Was neu ist: Das Programm fragt Open-Meteo mit HTTP und liest die Antwort
// im Format JSON. Dafür sorgen die Pakete net/http und encoding/json.
//
// Das Programm braucht eine Verbindung zum Internet.
//
// Mentoriat Imperative Programmierung, FernUniversität in Hagen.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Koordinaten sagen, wo eine Stadt liegt.
type Koordinaten struct {
	Breitengrad float64 // nördlich des Äquators, Berlin 52.52
	Laengengrad float64 // östlich von Greenwich, Berlin 13.41, westlich negativ
}

// Stadt ist ein Ort, für den das Programm das Wetter holt.
type Stadt struct {
	Name        string
	Land        string
	Koordinaten Koordinaten // ein Verbund als Feld
}

// AktuellesWetter ist das Wetter zum Zeitpunkt der Abfrage.
type AktuellesWetter struct {
	Zeitpunkt        string  // zum Beispiel "2026-10-15T22:00"
	Temperatur       float64 // in Grad Celsius
	Gefuehlt         float64 // gefühlte Temperatur
	Wind             float64 // in km/h
	Windrichtung     int     // woher der Wind kommt, in Grad: 0 Norden, 90 Osten
	Niederschlag     float64 // in mm
	Luftfeuchtigkeit int     // in Prozent
	Beschreibung     string  // zum Beispiel "Bewölkt"
}

// TagesVorhersage ist die Vorhersage für einen Tag.
type TagesVorhersage struct {
	Datum        string // zum Beispiel "2026-10-15"
	TempMax      float64
	TempMin      float64
	Niederschlag float64
	Beschreibung string
}

// WetterBericht fasst alles zu einer Stadt zusammen.
type WetterBericht struct {
	Stadt      Stadt
	Aktuell    AktuellesWetter
	Vorhersage [7]TagesVorhersage // ein Array von Verbünden als Feld
}

// Die Städte, deren Wetter das Programm holt. Probiere deine eigene Stadt aus.
// Die Koordinaten findest du zum Beispiel in der Wikipedia.
var staedte = [6]Stadt{
	{"Berlin", "Deutschland", Koordinaten{52.52, 13.41}},
	{"München", "Deutschland", Koordinaten{48.14, 11.58}},
	{"Hamburg", "Deutschland", Koordinaten{53.55, 10.00}},
	{"Hagen", "Deutschland", Koordinaten{51.36, 7.47}},
	{"Paris", "Frankreich", Koordinaten{48.85, 2.35}},
	{"London", "Großbritannien", Koordinaten{51.51, -0.13}},
}

func main() {
	fmt.Println("Wetter-Vergleich mit Daten aus dem Internet")
	fmt.Println("===========================================")
	fmt.Println()

	// 1. Wetter holen
	fmt.Println("1. Wetter von Open-Meteo")
	fmt.Println()
	var berichte [6]WetterBericht
	geholt := 0
	for i := 0; i < len(staedte); i++ {
		b, err := holeBericht(staedte[i])
		if err != nil {
			fmt.Printf("   %-10s nicht geholt: %v\n", staedte[i].Name, err)
			continue
		}
		berichte[geholt] = b
		geholt++
		time.Sleep(300 * time.Millisecond) // Open-Meteo nicht mit Anfragen überhäufen
	}
	if geholt == 0 {
		fmt.Println()
		fmt.Println("   Kein Wetter ließ sich holen. Ist der Rechner mit dem Internet verbunden?")
		return
	}
	fmt.Printf("   %-10s  %7s  %8s  %8s  %7s  %s\n", "Stadt", "Temp.", "gefühlt", "Wind", "Feuchte", "Wetter")
	for i := 0; i < geholt; i++ {
		a := berichte[i].Aktuell
		fmt.Printf("   %-10s  %5.1f°C  %6.1f°C  %3.0f km/h  %6d%%  %s\n",
			berichte[i].Stadt.Name, a.Temperatur, a.Gefuehlt, a.Wind, a.Luftfeuchtigkeit, a.Beschreibung)
	}
	fmt.Println()

	// 2. Vergleichen
	fmt.Println("2. Vergleich")
	fmt.Println()
	waermste, kaelteste, windigste := 0, 0, 0
	summe := 0.0
	for i := 0; i < geholt; i++ {
		t := berichte[i].Aktuell.Temperatur
		if t > berichte[waermste].Aktuell.Temperatur {
			waermste = i
		}
		if t < berichte[kaelteste].Aktuell.Temperatur {
			kaelteste = i
		}
		if berichte[i].Aktuell.Wind > berichte[windigste].Aktuell.Wind {
			windigste = i
		}
		summe += t
	}
	fmt.Printf("   Am wärmsten:   %s, %.1f °C\n", berichte[waermste].Stadt.Name, berichte[waermste].Aktuell.Temperatur)
	fmt.Printf("   Am kältesten:  %s, %.1f °C\n", berichte[kaelteste].Stadt.Name, berichte[kaelteste].Aktuell.Temperatur)
	fmt.Printf("   Am windigsten: %s, %.0f km/h\n", berichte[windigste].Stadt.Name, berichte[windigste].Aktuell.Wind)
	fmt.Printf("   Durchschnitt:  %.1f °C\n\n", summe/float64(geholt))

	// 3. Die Woche für die erste Stadt
	b := berichte[0]
	fmt.Printf("3. Die nächsten sieben Tage in %s\n\n", b.Stadt.Name)
	fmt.Printf("   %-9s  %-18s  %13s  %7s\n", "Tag", "Wetter", "Temperatur", "Regen")
	regentage := 0
	regen := 0.0
	for i := 0; i < len(b.Vorhersage); i++ {
		tag := b.Vorhersage[i]
		fmt.Printf("   %-9s  %-18s  %4.0f bis %2.0f°C  %4.1f mm\n",
			wochentag(tag.Datum), tag.Beschreibung, tag.TempMin, tag.TempMax, tag.Niederschlag)
		regen += tag.Niederschlag
		if tag.Niederschlag > 0.5 {
			regentage++
		}
	}
	fmt.Printf("\n   Regen in der Woche: %.1f mm, an %d von 7 Tagen mehr als 0,5 mm\n", regen, regentage)
}

// openMeteoAntwort hat den Aufbau der Antwort von Open-Meteo, soweit das
// Programm sie braucht. Das Paket encoding/json füllt die Felder, deren Namen
// hinter json: stehen. Die Vorhersage kommt als Arrays mit sieben Werten.
type openMeteoAntwort struct {
	Current struct {
		Time       string  `json:"time"`
		Temperatur float64 `json:"temperature_2m"`
		Gefuehlt   float64 `json:"apparent_temperature"`
		Feuchte    int     `json:"relative_humidity_2m"`
		Regen      float64 `json:"precipitation"`
		Wind       float64 `json:"windspeed_10m"`
		Richtung   int     `json:"winddirection_10m"`
		WetterCode int     `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		Time       [7]string  `json:"time"`
		Max        [7]float64 `json:"temperature_2m_max"`
		Min        [7]float64 `json:"temperature_2m_min"`
		Regen      [7]float64 `json:"precipitation_sum"`
		WetterCode [7]int     `json:"weather_code"`
	} `json:"daily"`
}

// holeBericht fragt Open-Meteo nach dem Wetter in einer Stadt.
func holeBericht(stadt Stadt) (WetterBericht, error) {
	adresse := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f"+
		"&current=temperature_2m,apparent_temperature,relative_humidity_2m,precipitation,windspeed_10m,winddirection_10m,weather_code"+
		"&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,weather_code&forecast_days=7&timezone=auto",
		stadt.Koordinaten.Breitengrad, stadt.Koordinaten.Laengengrad)
	client := http.Client{Timeout: 10 * time.Second}
	antwort, err := client.Get(adresse)
	if err != nil {
		return WetterBericht{}, fmt.Errorf("keine Verbindung")
	}
	defer antwort.Body.Close()
	if antwort.StatusCode != http.StatusOK {
		return WetterBericht{}, fmt.Errorf("Open-Meteo antwortet mit %s", antwort.Status)
	}
	var d openMeteoAntwort
	if err := json.NewDecoder(antwort.Body).Decode(&d); err != nil {
		return WetterBericht{}, fmt.Errorf("die Antwort ließ sich nicht lesen")
	}

	b := WetterBericht{Stadt: stadt}
	b.Aktuell = AktuellesWetter{
		Zeitpunkt:        d.Current.Time,
		Temperatur:       d.Current.Temperatur,
		Gefuehlt:         d.Current.Gefuehlt,
		Wind:             d.Current.Wind,
		Windrichtung:     d.Current.Richtung,
		Niederschlag:     d.Current.Regen,
		Luftfeuchtigkeit: d.Current.Feuchte,
		Beschreibung:     wetterText(d.Current.WetterCode),
	}
	for i := 0; i < 7; i++ {
		b.Vorhersage[i] = TagesVorhersage{
			Datum:        d.Daily.Time[i],
			TempMax:      d.Daily.Max[i],
			TempMin:      d.Daily.Min[i],
			Niederschlag: d.Daily.Regen[i],
			Beschreibung: wetterText(d.Daily.WetterCode[i]),
		}
	}
	return b, nil
}

// wetterText übersetzt den Wettercode der Weltorganisation für Meteorologie
// in einen deutschen Text.
func wetterText(code int) string {
	switch code {
	case 0:
		return "Klar"
	case 1, 2:
		return "Teilweise bewölkt"
	case 3:
		return "Bewölkt"
	case 45, 48:
		return "Nebel"
	case 51, 53, 55:
		return "Nieselregen"
	case 61, 80:
		return "Leichter Regen"
	case 63, 81:
		return "Regen"
	case 65, 82:
		return "Starker Regen"
	case 71, 73, 75, 77, 85, 86:
		return "Schnee"
	case 95, 96, 99:
		return "Gewitter"
	default:
		return fmt.Sprintf("Code %d", code)
	}
}

// wochentag macht aus "2026-10-15" den Text "Do 15.10.".
func wochentag(datum string) string {
	t, err := time.Parse("2006-01-02", datum)
	if err != nil {
		return datum
	}
	namen := [7]string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"}
	return fmt.Sprintf("%s %02d.%02d.", namen[t.Weekday()], t.Day(), int(t.Month()))
}
