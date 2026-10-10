# Verweise auf aktuelle Forschung und Entwicklung

Diese Liste nennt jeden Hinweis und jede Lektion, die eine Seite mit einer aktuellen Arbeit verbindet. Einmal im Semester wird geprüft, ob die Links noch stimmen und welche Angaben veraltet sind.

## LittleBit: Gewichte als Bits, Multiplikation als XOR

- Eingetragen am 2026-10-09 als Hinweis.
- Seite: `kap4-primitive-datentypen/go_k4_rechenwerk_interaktiv.html`, Karte „Die Bit-Operatoren in Go“.
- Der Hinweis verlinkt die Seite Zahlenkreis und nennt Abschnitt 4.3.4 des Kurstexts.
- Quelle: Lee, Kim, You, Kim: „LittleBit: Ultra Low-Bit Quantization via Latent Factorization“, NeurIPS 2025, <https://arxiv.org/abs/2506.13771>. Geprüft an Fassung v5.
- Aus der Quelle stammen: unter 0,9 GB für Llama2-13B, XOR auf dem Vorzeichenbit, gepackt in `uint32_t`.
- Selbst gerechnet sind die 26 GB: 13 Milliarden Gewichte mal 16 Bit.
- Veralten können: der Preis der Grafikkarte „über 2.000 €“ und die Aussage „jedes heutigen Smartphones“.
- Belegt ist nur, dass das Modell in den Speicher passt. Auf einem Telefon ausgeführt hat es in der Quelle niemand.
