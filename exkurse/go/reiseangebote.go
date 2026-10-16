// Reiseangebote aus dem Internet, mit Wetter am Zielort
//
// Das Programm liest die neuesten Reiseangebote aus zwei Feeds: mydealz,
// Gruppe Reisen, und Urlaubspiraten. Ein Feed ist eine Liste der neuesten
// Beiträge einer Webseite in einem festen Format. In jedem Angebot sucht das
// Programm einen Zielort, einen Preis und die Zahl der Nächte. Für jeden
// gefundenen Zielort holt es das Wetter bei Open-Meteo. Am Ende sortiert es
// die Angebote nach dem Preis je Nacht.
//
// Was du aus dem Kurs kennst:
//   - Verbünde mit einem Zeiger als Feld: ReiseDeal -> *WetterInfo, Abschnitt 8.2.4,
//     und nil, wenn es kein Wetter gibt, Abschnitt 7.4
//   - ein Array von Verbünden mit einer Zählvariable, Kapitel 6 und Abschnitt 8.4
//   - Bubblesort auf einem Array von Verbünden, Abschnitt 8.4.2
//   - eine Funktion, die über einen Zeiger einen Verbund ändert, Abschnitt 8.2.7
//
// Was neu ist: Das Programm holt die Feeds mit HTTP und liest sie im Format
// XML, das Wetter im Format JSON. Zielort, Preis und Nächte findet es mit
// regulären Ausdrücken, das sind Suchmuster für Text. Dafür sorgen die Pakete
// net/http, encoding/xml, encoding/json und regexp.
//
// Das Programm braucht eine Verbindung zum Internet. Die Angebote sind bei
// jedem Lauf andere. Nicht jedes Angebot nennt Preis und Zielort im Titel.
//
// Mentoriat Imperative Programmierung, FernUniversität in Hagen.
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// ReiseDeal ist ein Angebot aus einem Feed.
type ReiseDeal struct {
	Titel        string
	Quelle       string      // aus welchem Feed das Angebot stammt
	Link         string      // Adresse des Beitrags
	Erschienen   time.Time   // wann der Beitrag erschienen ist
	Zielort      string      // leer, wenn das Programm keinen findet
	Preis        float64     // in Euro, 0, wenn das Programm keinen findet
	Naechte      int         // 0, wenn das Programm keine Zahl findet
	PreisJeNacht float64     // Preis / Naechte, ohne Nächte der Preis
	Wetter       *WetterInfo // nil, wenn es kein Wetter gibt
}

// WetterInfo ist das Wetter an einem Zielort.
type WetterInfo struct {
	Temperatur   float64
	Beschreibung string
}

// Ort ist ein Zielort, den das Programm im Text eines Angebots erkennt.
type Ort struct {
	Name        string
	Breitengrad float64
	Laengengrad float64
}

// Die Zielorte, die das Programm kennt. Ergänze eigene Ziele mit ihren
// Koordinaten und vergrößere dafür die Länge des Arrays.
var orte = [32]Ort{
	{"Mallorca", 39.57, 2.65}, {"Teneriffa", 28.47, -16.25}, {"Gran Canaria", 28.12, -15.43},
	{"Lanzarote", 28.96, -13.55}, {"Fuerteventura", 28.50, -13.86}, {"Madeira", 32.65, -16.91},
	{"Kreta", 35.24, 24.81}, {"Malta", 35.90, 14.51}, {"Antalya", 36.90, 30.70},
	{"Hurghada", 27.26, 33.81}, {"Dubai", 25.20, 55.27}, {"Bangkok", 13.75, 100.50},
	{"New York", 40.71, -74.01}, {"Island", 64.15, -21.94}, {"Barcelona", 41.39, 2.16},
	{"Madrid", 40.42, -3.70}, {"Lissabon", 38.72, -9.14}, {"Porto", 41.15, -8.61},
	{"Rom", 41.90, 12.50}, {"Venedig", 45.44, 12.33}, {"Mailand", 45.46, 9.19},
	{"Paris", 48.85, 2.35}, {"London", 51.51, -0.13}, {"Amsterdam", 52.37, 4.90},
	{"Prag", 50.08, 14.44}, {"Wien", 48.21, 16.37}, {"Budapest", 47.50, 19.04},
	{"Salzburg", 47.81, 13.04}, {"Dresden", 51.05, 13.74}, {"Berlin", 52.52, 13.41},
	{"Hamburg", 53.55, 10.00}, {"München", 48.14, 11.58},
}

// Feed ist eine Quelle für Angebote.
type Feed struct {
	Name    string
	Adresse string
}

var feeds = [2]Feed{
	{"mydealz Reisen", "https://www.mydealz.de/rss/gruppe/reisen"},
	{"Urlaubspiraten", "https://www.urlaubspiraten.de/feed"},
}

// Höchstens so viele Angebote nimmt das Programm auf.
const maxDeals = 80

func main() {
	fmt.Println("Reiseangebote aus dem Internet, mit Wetter am Zielort")
	fmt.Println("=====================================================")
	fmt.Println()

	// 1. Feeds lesen
	fmt.Println("1. Feeds lesen")
	fmt.Println()
	var deals [maxDeals]ReiseDeal
	anzahl := 0
	for i := 0; i < len(feeds); i++ {
		vorher := anzahl
		err := leseFeed(feeds[i], &deals, &anzahl)
		if err != nil {
			fmt.Printf("   %-16s nicht gelesen: %v\n", feeds[i].Name, err)
			continue
		}
		fmt.Printf("   %-16s %d Angebote\n", feeds[i].Name, anzahl-vorher)
	}
	if anzahl == 0 {
		fmt.Println()
		fmt.Println("   Kein Feed ließ sich lesen. Ist der Rechner mit dem Internet verbunden?")
		return
	}
	fmt.Println()

	// 2. Zielort, Preis und Nächte suchen
	fmt.Println("2. Zielort, Preis und Nächte suchen")
	fmt.Println()
	vollstaendig := 0
	for i := 0; i < anzahl; i++ {
		werteAus(&deals[i])
		if deals[i].Zielort != "" && deals[i].Preis > 0 {
			vollstaendig++
		}
	}
	fmt.Printf("   %d von %d Angeboten nennen einen bekannten Zielort und einen Preis.\n\n", vollstaendig, anzahl)

	// 3. Wetter für die Zielorte holen, für jeden Ort nur einmal
	fmt.Println("3. Wetter an den Zielorten")
	fmt.Println()
	var wetter [len(orte)]WetterInfo
	var geholt [len(orte)]bool
	var versucht [len(orte)]bool
	for i := 0; i < anzahl; i++ {
		k := ortNummer(deals[i].Zielort)
		if k < 0 || deals[i].Preis == 0 {
			continue
		}
		if !versucht[k] {
			versucht[k] = true
			w, err := holeWetter(orte[k])
			if err != nil {
				fmt.Printf("   %-14s kein Wetter: %v\n", orte[k].Name, err)
			} else {
				wetter[k] = w
				geholt[k] = true
				fmt.Printf("   %-14s %5.1f °C, %s\n", orte[k].Name, w.Temperatur, w.Beschreibung)
			}
			time.Sleep(300 * time.Millisecond)
		}
		if geholt[k] {
			deals[i].Wetter = &wetter[k] // ein Zeiger auf das Element des Arrays
		}
	}
	fmt.Println()

	// 4. Nach Preis je Nacht sortieren und ausgeben
	sortiere(&deals, anzahl)
	fmt.Println("4. Die günstigsten Angebote mit Zielort und Preis")
	fmt.Println()
	fmt.Printf("   %-3s  %-13s  %8s  %7s  %9s  %-22s  %s\n", "Nr.", "Zielort", "Preis €", "Nächte", "€/Nacht", "Wetter", "Titel")
	gezeigt := 0
	for i := 0; i < anzahl && gezeigt < 10; i++ {
		d := deals[i]
		if d.Zielort == "" || d.Preis == 0 {
			continue
		}
		gezeigt++
		naechte := "?"
		if d.Naechte > 0 {
			naechte = strconv.Itoa(d.Naechte)
		}
		wetterText := "kein Wetter"
		if d.Wetter != nil { // ohne diese Prüfung bräche das Programm bei nil ab
			wetterText = fmt.Sprintf("%.0f °C, %s", d.Wetter.Temperatur, d.Wetter.Beschreibung)
		}
		fmt.Printf("   %3d  %-13s  %8.2f  %7s  %9.2f  %-22s  %s\n",
			gezeigt, d.Zielort, d.Preis, naechte, d.PreisJeNacht, kuerze(wetterText, 22), kuerze(d.Titel, 50))
	}
	if gezeigt == 0 {
		fmt.Println("   Heute nennt kein Angebot einen bekannten Zielort und einen Preis.")
	}
	fmt.Println()
	fmt.Println("   ? heißt: Das Angebot nennt keine Zahl der Nächte, zum Beispiel bei einem Flug.")
	fmt.Println("   €/Nacht ist dann der ganze Preis.")
}

// eintrag ist ein Beitrag im Feed im Format RSS, soweit das Programm ihn braucht.
type eintrag struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

// leseFeed holt einen Feed und hängt seine Beiträge an das Array deals an.
// Es liest den Feed Element für Element: Bei jedem Element <item> liest es
// einen Beitrag ein.
func leseFeed(f Feed, deals *[maxDeals]ReiseDeal, anzahl *int) error {
	anfrage, err := http.NewRequest("GET", f.Adresse, nil)
	if err != nil {
		return err
	}
	anfrage.Header.Set("User-Agent", "Mozilla/5.0 (Mentoriat Go)")
	client := http.Client{Timeout: 15 * time.Second}
	antwort, err := client.Do(anfrage)
	if err != nil {
		return fmt.Errorf("keine Verbindung")
	}
	defer antwort.Body.Close()
	if antwort.StatusCode != http.StatusOK {
		return fmt.Errorf("die Quelle antwortet mit %s", antwort.Status)
	}

	dec := xml.NewDecoder(antwort.Body)
	dec.Strict = false          // manche Feeds halten die Regeln von XML nicht genau ein
	dec.Entity = xml.HTMLEntity // und enthalten Zeichen wie &nbsp; aus HTML
	dec.CharsetReader = func(_ string, r io.Reader) (io.Reader, error) { return r, nil }
	gelesen := 0
	for *anzahl < maxDeals && gelesen < 40 {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			if gelesen == 0 {
				return fmt.Errorf("der Feed ließ sich nicht lesen")
			}
			break
		}
		start, ok := token.(xml.StartElement)
		if !ok || start.Name.Local != "item" {
			continue
		}
		var e eintrag
		if dec.DecodeElement(&e, &start) != nil {
			continue
		}
		deals[*anzahl] = ReiseDeal{
			Titel:      strings.TrimSpace(e.Title),
			Quelle:     f.Name,
			Link:       strings.TrimSpace(e.Link),
			Erschienen: leseDatum(e.PubDate),
		}
		*anzahl++
		gelesen++
	}
	if gelesen == 0 {
		return fmt.Errorf("der Feed enthält keine Beiträge")
	}
	return nil
}

// leseDatum versteht die Datumsangaben von RSS, zum Beispiel
// "Thu, 15 Oct 2026 20:00:00 +0200".
func leseDatum(text string) time.Time {
	formate := [3]string{time.RFC1123Z, time.RFC1123, "Mon, 2 Jan 2006 15:04:05 -0700"}
	for i := 0; i < len(formate); i++ {
		if t, err := time.Parse(formate[i], strings.TrimSpace(text)); err == nil {
			return t
		}
	}
	return time.Time{}
}

// Suchmuster für Preis und Nächte. Ein Preis ist eine Zahl vor €, zum Beispiel
// "133€", "46,98€" oder "1.730 €". Nächte stehen als "2 Nächte" oder "2 ÜN"
// im Titel, Wochen als "1 Woche". "pro Nacht" heißt: Der Preis gilt schon je Nacht.
var (
	musterPreis    = regexp.MustCompile(`(\d{1,3}(?:\.\d{3})+|\d+)(?:,(\d{1,2}))?\s*€`)
	musterNaechte  = regexp.MustCompile(`(?i)(\d+)\s*(?:nächte|nacht|ün\b|übernachtung)`)
	musterWochen   = regexp.MustCompile(`(?i)(\d+)\s*wochen?\b`)
	musterProNacht = regexp.MustCompile(`(?i)(pro|je|/)\s*nacht`)
)

// werteAus sucht im Titel Zielort, Preis und Nächte und trägt sie über den
// Zeiger d direkt in das Angebot ein.
func werteAus(d *ReiseDeal) {
	for k := 0; k < len(orte); k++ {
		muster := regexp.MustCompile(`(?i)(^|[^\p{L}])` + regexp.QuoteMeta(orte[k].Name) + `($|[^\p{L}])`)
		if muster.MatchString(d.Titel) {
			d.Zielort = orte[k].Name
			break
		}
	}
	if teile := musterPreis.FindStringSubmatch(d.Titel); teile != nil {
		zahl := strings.ReplaceAll(teile[1], ".", "")
		if teile[2] != "" {
			zahl += "." + teile[2]
		}
		if preis, err := strconv.ParseFloat(zahl, 64); err == nil && preis >= 10 && preis <= 10000 {
			d.Preis = preis
		}
	}
	if teile := musterNaechte.FindStringSubmatch(d.Titel); teile != nil {
		d.Naechte, _ = strconv.Atoi(teile[1])
	}
	if teile := musterWochen.FindStringSubmatch(d.Titel); teile != nil && d.Naechte == 0 {
		wochen, _ := strconv.Atoi(teile[1])
		d.Naechte = 7 * wochen
	}
	d.PreisJeNacht = d.Preis
	if d.Naechte > 0 && !musterProNacht.MatchString(d.Titel) {
		d.PreisJeNacht = d.Preis / float64(d.Naechte)
	}
}

// ortNummer liefert die Stelle eines Zielorts im Array orte, -1, wenn es ihn nicht gibt.
func ortNummer(name string) int {
	for k := 0; k < len(orte); k++ {
		if orte[k].Name == name {
			return k
		}
	}
	return -1
}

// holeWetter fragt Open-Meteo nach dem Wetter an einem Ort.
func holeWetter(o Ort) (WetterInfo, error) {
	adresse := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.2f&longitude=%.2f&current=temperature_2m,weather_code&timezone=auto",
		o.Breitengrad, o.Laengengrad)
	client := http.Client{Timeout: 10 * time.Second}
	antwort, err := client.Get(adresse)
	if err != nil {
		return WetterInfo{}, fmt.Errorf("keine Verbindung")
	}
	defer antwort.Body.Close()
	if antwort.StatusCode != http.StatusOK {
		return WetterInfo{}, fmt.Errorf("Open-Meteo antwortet mit %s", antwort.Status)
	}
	var d struct {
		Current struct {
			Temperatur float64 `json:"temperature_2m"`
			Code       int     `json:"weather_code"`
		} `json:"current"`
	}
	if err := json.NewDecoder(antwort.Body).Decode(&d); err != nil {
		return WetterInfo{}, fmt.Errorf("die Antwort ließ sich nicht lesen")
	}
	return WetterInfo{Temperatur: d.Current.Temperatur, Beschreibung: wetterText(d.Current.Code)}, nil
}

// wetterText übersetzt den Wettercode der Weltorganisation für Meteorologie.
func wetterText(code int) string {
	switch code {
	case 0:
		return "klar"
	case 1, 2:
		return "teils bewölkt"
	case 3:
		return "bewölkt"
	case 45, 48:
		return "Nebel"
	case 51, 53, 55, 61, 63, 80, 81:
		return "Regen"
	case 65, 82:
		return "starker Regen"
	case 71, 73, 75, 77, 85, 86:
		return "Schnee"
	case 95, 96, 99:
		return "Gewitter"
	default:
		return fmt.Sprintf("Code %d", code)
	}
}

// sortiere ordnet die ersten n Angebote mit Bubblesort nach dem Preis je Nacht.
// Angebote ohne Zielort oder ohne Preis kommen ans Ende.
func sortiere(deals *[maxDeals]ReiseDeal, n int) {
	for durchgang := 0; durchgang < n-1; durchgang++ {
		for j := 0; j < n-1-durchgang; j++ {
			if spaeter(deals[j], deals[j+1]) {
				deals[j], deals[j+1] = deals[j+1], deals[j]
			}
		}
	}
}

// spaeter sagt, ob a in der Liste hinter b gehört.
func spaeter(a, b ReiseDeal) bool {
	aGut := a.Zielort != "" && a.Preis > 0
	bGut := b.Zielort != "" && b.Preis > 0
	if aGut != bGut {
		return !aGut
	}
	return a.PreisJeNacht > b.PreisJeNacht
}

// kuerze schneidet einen Text nach n Zeichen ab. Es zählt Zeichen, nicht Bytes,
// damit kein Umlaut zerbricht.
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
