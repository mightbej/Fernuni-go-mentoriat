// Aktien-Portfolio mit Kursen aus dem Internet
//
// Das Programm holt die aktuellen Kurse von fünf Aktien bei Yahoo Finance.
// Dann legt es ein Portfolio mit drei Positionen an und berechnet, wie viel
// jede Position gewonnen oder verloren hat.
//
// Was du aus dem Kurs kennst:
//   - Verbünde, die ineinander stecken: Portfolio -> PortfolioPosition -> Aktie,
//     Abschnitt 8.3
//   - Arrays mit fester Länge und eine Zählvariable, Kapitel 6
//   - Funktionen mit einem Wert und einem Fehler als Rückgabe, Abschnitt 7.1.2
//   - ein Zeiger auf einen Verbund als Parameter, Abschnitt 8.2.7
//
// Was neu ist: Das Programm fragt Yahoo Finance mit HTTP und liest die Antwort
// im Format JSON. Dafür sorgen die Pakete net/http und encoding/json.
//
// Das Programm braucht eine Verbindung zum Internet. Die Kurse sind bei jedem
// Lauf anders. Yahoo Finance ist keine offizielle Schnittstelle und kann eine
// Anfrage ablehnen.
//
// Mentoriat Imperative Programmierung, FernUniversität in Hagen.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"
)

// Aktie hält den Kurs einer Aktie fest.
type Aktie struct {
	Symbol          string  // Kürzel an der Börse, zum Beispiel "AAPL"
	Name            string  // Name des Unternehmens, zum Beispiel "Apple Inc."
	AktuellerPreis  float64 // Kurs in US-Dollar
	VorherigerPreis float64 // Kurs am Ende des letzten Handelstags
	Aenderung       float64 // AktuellerPreis - VorherigerPreis
	ProzentAnders   float64 // die Änderung in Prozent
}

// PortfolioPosition ist eine Aktie im Portfolio: welche, wie viele, zu welchem Preis.
type PortfolioPosition struct {
	Aktie         Aktie   // ein Verbund als Feld
	Anzahl        int     // wie viele Stück
	Kaufpreis     float64 // Preis je Stück beim Kauf
	Kaufdatum     string  // zum Beispiel "2024-01-15"
	Investiert    float64 // Anzahl * Kaufpreis
	AktuellerWert float64 // Anzahl * Aktie.AktuellerPreis
	GewinnVerlust float64 // AktuellerWert - Investiert
	GewinnProzent float64 // GewinnVerlust in Prozent von Investiert
}

// Portfolio sind alle Positionen zusammen. Das Array hat Platz für 10
// Positionen, belegt sind die ersten Anzahl.
type Portfolio struct {
	Name          string
	Positionen    [10]PortfolioPosition
	Anzahl        int
	Investiert    float64
	AktuellerWert float64
	GewinnVerlust float64
	GewinnProzent float64
}

// Die Aktien, deren Kurse das Programm holt. Probiere eigene Kürzel aus,
// zum Beispiel "SAP" oder "NVDA".
var symbole = [5]string{"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA"}

func main() {
	fmt.Println("Aktien-Portfolio mit Kursen aus dem Internet")
	fmt.Println("============================================")
	fmt.Println()

	// 1. Kurse holen
	fmt.Println("1. Kurse von Yahoo Finance")
	fmt.Println()
	var kurse [5]Aktie
	geholt := 0
	for i := 0; i < len(symbole); i++ {
		a, err := holeAktie(symbole[i])
		if err != nil {
			fmt.Printf("   %-6s nicht geholt: %v\n", symbole[i], err)
			continue
		}
		kurse[geholt] = a
		geholt++
	}
	if geholt == 0 {
		fmt.Println()
		fmt.Println("   Kein Kurs ließ sich holen. Ist der Rechner mit dem Internet verbunden?")
		return
	}
	fmt.Printf("   %-6s  %-22s  %10s  %9s\n", "Kürzel", "Unternehmen", "Kurs $", "Änderung")
	for i := 0; i < geholt; i++ {
		a := kurse[i]
		fmt.Printf("   %-6s  %-22s  %10.2f  %+8.2f%%\n", a.Symbol, kuerze(a.Name, 22), a.AktuellerPreis, a.ProzentAnders)
	}
	fmt.Println()

	// 2. Portfolio anlegen. Kaufpreise und Kaufdaten sind ausgedacht.
	fmt.Println("2. Portfolio anlegen")
	fmt.Println()
	p := Portfolio{Name: "Mein Tech-Portfolio"}
	fuegeHinzu(&p, kurse, geholt, "AAPL", 10, 150.00, "2024-01-15")
	fuegeHinzu(&p, kurse, geholt, "MSFT", 5, 350.00, "2024-02-20")
	fuegeHinzu(&p, kurse, geholt, "TSLA", 15, 200.00, "2024-03-10")
	if p.Anzahl == 0 {
		fmt.Println("   Für keine Position gibt es einen Kurs.")
		return
	}
	berechneSummen(&p)
	fmt.Println()

	// 3. Portfolio ausgeben
	fmt.Println("3. " + p.Name)
	fmt.Println()
	fmt.Printf("   %-6s  %6s  %10s  %10s  %12s  %8s\n", "Kürzel", "Anzahl", "Kauf $", "Heute $", "Gewinn $", "Prozent")
	for i := 0; i < p.Anzahl; i++ {
		pos := p.Positionen[i]
		fmt.Printf("   %-6s  %6d  %10.2f  %10.2f  %+12.2f  %+7.1f%%\n",
			pos.Aktie.Symbol, pos.Anzahl, pos.Kaufpreis, pos.Aktie.AktuellerPreis, pos.GewinnVerlust, pos.GewinnProzent)
	}
	fmt.Println("   ----------------------------------------------------------------")
	fmt.Printf("   %-6s  %6s  %10s  %10s  %+12.2f  %+7.1f%%\n", "Summe", "", "", "", p.GewinnVerlust, p.GewinnProzent)
	fmt.Printf("\n   Investiert: %.2f $, heute wert: %.2f $\n\n", p.Investiert, p.AktuellerWert)

	// 4. Beste und schlechteste Position
	beste, schlechteste := besteUndSchlechteste(p)
	fmt.Println("4. Auswertung")
	fmt.Println()
	fmt.Printf("   Beste Position:        %s, %+.2f $\n", beste.Aktie.Symbol, beste.GewinnVerlust)
	fmt.Printf("   Schlechteste Position: %s, %+.2f $\n", schlechteste.Aktie.Symbol, schlechteste.GewinnVerlust)
}

// fuegeHinzu legt eine Position an, wenn es für das Kürzel einen Kurs gibt.
// p ist ein Zeiger, damit die Funktion das Portfolio selbst ändert.
func fuegeHinzu(p *Portfolio, kurse [5]Aktie, geholt int, symbol string, anzahl int, kaufpreis float64, datum string) {
	for i := 0; i < geholt; i++ {
		if kurse[i].Symbol == symbol {
			a := kurse[i]
			pos := PortfolioPosition{
				Aktie:         a,
				Anzahl:        anzahl,
				Kaufpreis:     kaufpreis,
				Kaufdatum:     datum,
				Investiert:    float64(anzahl) * kaufpreis,
				AktuellerWert: float64(anzahl) * a.AktuellerPreis,
			}
			pos.GewinnVerlust = pos.AktuellerWert - pos.Investiert
			pos.GewinnProzent = pos.GewinnVerlust / pos.Investiert * 100
			p.Positionen[p.Anzahl] = pos
			p.Anzahl++
			fmt.Printf("   %d x %s gekauft zu %.2f $ am %s\n", anzahl, symbol, kaufpreis, datum)
			return
		}
	}
	fmt.Printf("   %s fehlt, für diese Aktie gibt es keinen Kurs\n", symbol)
}

// berechneSummen zählt die Positionen zusammen.
func berechneSummen(p *Portfolio) {
	p.Investiert, p.AktuellerWert = 0, 0
	for i := 0; i < p.Anzahl; i++ {
		p.Investiert += p.Positionen[i].Investiert
		p.AktuellerWert += p.Positionen[i].AktuellerWert
	}
	p.GewinnVerlust = p.AktuellerWert - p.Investiert
	p.GewinnProzent = p.GewinnVerlust / p.Investiert * 100
}

// besteUndSchlechteste sucht die Positionen mit dem größten und dem kleinsten Gewinn.
func besteUndSchlechteste(p Portfolio) (PortfolioPosition, PortfolioPosition) {
	beste, schlechteste := p.Positionen[0], p.Positionen[0]
	for i := 1; i < p.Anzahl; i++ {
		if p.Positionen[i].GewinnVerlust > beste.GewinnVerlust {
			beste = p.Positionen[i]
		}
		if p.Positionen[i].GewinnVerlust < schlechteste.GewinnVerlust {
			schlechteste = p.Positionen[i]
		}
	}
	return beste, schlechteste
}

// yahooAntwort hat den Aufbau der Antwort von Yahoo Finance, soweit das
// Programm sie braucht. Das Paket encoding/json füllt die Felder, deren
// Namen hinter json: stehen. Result ist ein Array mit einem Element.
type yahooAntwort struct {
	Chart struct {
		Result [1]struct {
			Meta struct {
				Symbol         string  `json:"symbol"`
				LongName       string  `json:"longName"`
				Preis          float64 `json:"regularMarketPrice"`
				VorherigerKurs float64 `json:"chartPreviousClose"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

// holeAktie fragt Yahoo Finance nach dem Kurs einer Aktie.
func holeAktie(symbol string) (Aktie, error) {
	anfrage, err := http.NewRequest("GET", "https://query1.finance.yahoo.com/v8/finance/chart/"+symbol, nil)
	if err != nil {
		return Aktie{}, err
	}
	// Ohne diese Angabe lehnt Yahoo Finance die Anfrage oft ab.
	anfrage.Header.Set("User-Agent", "Mozilla/5.0 (Mentoriat Go)")
	client := http.Client{Timeout: 10 * time.Second}
	antwort, err := client.Do(anfrage)
	if err != nil {
		return Aktie{}, fmt.Errorf("keine Verbindung")
	}
	defer antwort.Body.Close()
	if antwort.StatusCode != http.StatusOK {
		return Aktie{}, fmt.Errorf("Yahoo Finance antwortet mit %s", antwort.Status)
	}
	var daten yahooAntwort
	if err := json.NewDecoder(antwort.Body).Decode(&daten); err != nil {
		return Aktie{}, fmt.Errorf("die Antwort ließ sich nicht lesen")
	}
	m := daten.Chart.Result[0].Meta
	if m.Preis == 0 {
		return Aktie{}, fmt.Errorf("kein Kurs in der Antwort")
	}
	a := Aktie{Symbol: m.Symbol, Name: m.LongName, AktuellerPreis: m.Preis, VorherigerPreis: m.VorherigerKurs}
	a.Aenderung = a.AktuellerPreis - a.VorherigerPreis
	if a.VorherigerPreis != 0 {
		a.ProzentAnders = a.Aenderung / a.VorherigerPreis * 100
	}
	return a, nil
}

// kuerze schneidet einen Text nach n Zeichen ab, damit die Tabelle gerade bleibt.
// Es zählt Zeichen, nicht Bytes, damit kein Umlaut zerbricht.
func kuerze(text string, n int) string {
	if utf8.RuneCountInString(text) <= n {
		return text
	}
	zeichen := 0
	for i := range text { // i ist die Stelle, an der das nächste Zeichen beginnt
		if zeichen == n-1 {
			return text[:i] + "…"
		}
		zeichen++
	}
	return text
}
