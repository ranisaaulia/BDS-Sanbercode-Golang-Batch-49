package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
)

type Smartphone struct {
	Model  string `json:"model"`
	Vendor string `json:"vendor"`
}

type Book struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Author      string `json:"author"`
	ReleaseYear int    `json:"release_year"`
}

func luasLingkaran(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2)
}

func kelilingLingkaran(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func volumeTabung(jariJari float64, tinggi float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) * tinggi
}

func main() {

	// Data JSON Samsung yang akan diuraikan
	jsonData := `[
        {"vendor":"Nokia","model":"N81 8GB"},
        {"vendor":"Nokia","model":"N82"},
        {"vendor":"Nokia","model":"N85"},
        {"vendor":"Nokia","model":"N86 8MP"},
        {"vendor":"Nokia","model":"N9"},
        {"vendor":"Nokia","model":"N90"},
        {"vendor":"Nokia","model":"N900"},
        {"vendor":"Nokia","model":"N91"},
        {"vendor":"Nokia","model":"N92"},
        {"vendor":"Nokia","model":"N93"},
        {"vendor":"Nokia","model":"N95"},
        {"vendor":"Nokia","model":"N95 8GB"},
        {"vendor":"Nokia","model":"N96"},
        {"vendor":"Nokia","model":"N97"},
        {"vendor":"Nokia","model":"N97 Mini"},
        {"vendor":"Nokia","model":"Oro"},
        {"vendor":"Nokia","model":"Surge"},
        {"vendor":"Nokia","model":"X1-00"},
        {"vendor":"Nokia","model":"X1-01"},
        {"vendor":"Nokia","model":"X2"},
        {"vendor":"Nokia","model":"X2-01"},
        {"vendor":"Nokia","model":"X2-02"},
        {"vendor":"Nokia","model":"X2-05"},
        {"vendor":"Nokia","model":"X3"},
        {"vendor":"Nokia","model":"X3 Touch and Type"},
        {"vendor":"Nokia","model":"X5-01"},
        {"vendor":"Nokia","model":"X6 16GB"},
        {"vendor":"Nokia","model":"X6 32GB"},
        {"vendor":"Nokia","model":"X6 8GB"},
        {"vendor":"Nokia","model":"X7-00"},
        {"vendor":"Orange","model":"SPV"},
        {"vendor":"Palm","model":"Pixi"},
        {"vendor":"Palm","model":"Pixi Plus"},
        {"vendor":"Palm","model":"Pre"},
        {"vendor":"Palm","model":"Pre 2"},
        {"vendor":"Palm","model":"Pre Plus"},
        {"vendor":"Palm","model":"Treo 680"},
        {"vendor":"Palm","model":"Treo 750"},
        {"vendor":"Panasonic","model":"Eluga"},
        {"vendor":"Panasonic","model":"GD67"},
        {"vendor":"Panasonic","model":"GD87"},
        {"vendor":"Panasonic","model":"GD92"},
        {"vendor":"Panasonic","model":"X300"},
        {"vendor":"Panasonic","model":"X400"},
        {"vendor":"Panasonic","model":"X500"},
        {"vendor":"Philips","model":"Xenium 9@9"},
        {"vendor":"Philips","model":"Xenium X503"},
        {"vendor":"RIM","model":"Blackberry Bold 9000"},
        {"vendor":"RIM","model":"BlackBerry Bold 9700"},
        {"vendor":"RIM","model":"Blackberry Bold 9780"},
        {"vendor":"RIM","model":"Blackberry Bold 9790"},
        {"vendor":"RIM","model":"BlackBerry Bold 9900"},
        {"vendor":"RIM","model":"BlackBerry Bold 9930"},
        {"vendor":"RIM","model":"Blackberry Curve 3G 9300"},
        {"vendor":"RIM","model":"BlackBerry Curve 8300"},
        {"vendor":"RIM","model":"BlackBerry Curve 8310"},
        {"vendor":"RIM","model":"Blackberry Curve 8320"},
        {"vendor":"RIM","model":"Blackberry Curve 8520"},
        {"vendor":"RIM","model":"Blackberry Curve 8530"},
        {"vendor":"RIM","model":"Blackberry Curve 8900"},
        {"vendor":"RIM","model":"Blackberry Curve 9320"},
        {"vendor":"RIM","model":"BlackBerry Curve 9350"},
        {"vendor":"RIM","model":"BlackBerry Curve 9360"},
        {"vendor":"RIM","model":"BlackBerry Curve 9370"},
        {"vendor":"RIM","model":"Blackberry Curve 9380"},
        {"vendor":"RIM","model":"BlackBerry Pearl 3G 9100"},
        {"vendor":"RIM","model":"BlackBerry Pearl 3G 9105"},
        {"vendor":"RIM","model":"BlackBerry Porsche Design P9981"},
        {"vendor":"RIM","model":"BlackBerry Storm 9500"},
        {"vendor":"RIM","model":"BlackBerry Storm 9530"},
        {"vendor":"RIM","model":"BlackBerry Storm2 9520"},
        {"vendor":"RIM","model":"BlackBerry Storm2 9550"},
        {"vendor":"RIM","model":"Blackberry Style 9670"},
        {"vendor":"RIM","model":"Blackberry Torch 9800"},
        {"vendor":"RIM","model":"BlackBerry Torch 9810"},
        {"vendor":"RIM","model":"BlackBerry Tour 9630"},
        {"vendor":"Sagem","model":"myX-5"},
        {"vendor":"Samsung","model":"4G LTE"},
        {"vendor":"Samsung","model":"723 S7230"},
        {"vendor":"Samsung","model":"A257 Magnet"},
        {"vendor":"Samsung","model":"Ativ S"},
        {"vendor":"Samsung","model":"B2100"},
        {"vendor":"Samsung","model":"B2710"},
        {"vendor":"Samsung","model":"B3310"},
        {"vendor":"Samsung","model":"B7330"},
        {"vendor":"Samsung","model":"B7350 Omnia Pro 4"},
        {"vendor":"Samsung","model":"B7722"},
        {"vendor":"Samsung","model":"Beat DJ M7600"},
        {"vendor":"Samsung","model":"C3050"},
        {"vendor":"Samsung","model":"C3060"},
        {"vendor":"Samsung","model":"C3300 Champ"},
        {"vendor":"Samsung","model":"C3510 Genoa"},
        {"vendor":"Samsung","model":"C3530"},
        {"vendor":"Samsung","model":"C5010"},
        {"vendor":"Samsung","model":"C5130"},
        {"vendor":"Samsung","model":"C5212"},
        {"vendor":"Samsung","model":"C5510"},
        {"vendor":"Samsung","model":"C6625"},
        {"vendor":"Samsung","model":"Captivate i897"},
        {"vendor":"Samsung","model":"Ch@t 322 C3222"},
        {"vendor":"Samsung","model":"Ch@t 335 S3350"},
        {"vendor":"Samsung","model":"Champ Neo Duos GT-C3262"},
        {"vendor":"Samsung","model":"Corby Pro B5310"},
        {"vendor":"Samsung","model":"Corby S3650 Genio Touch"},
        {"vendor":"Samsung","model":"Diva S7070"},
        {"vendor":"Samsung","model":"E1080"},
        {"vendor":"Samsung","model":"E1120"},
        {"vendor":"Samsung","model":"E1360"},
        {"vendor":"Samsung","model":"E2120"},
        {"vendor":"Samsung","model":"E2210"},
        {"vendor":"Samsung","model":"E2370 xcover"},
        {"vendor":"Samsung","model":"E2550 Monte Slide"},
        {"vendor":"Samsung","model":"Epic 4G D700"},
        {"vendor":"Samsung","model":"Focus"},
        {"vendor":"Samsung","model":"Galaxy Ace 2"},
        {"vendor":"Samsung","model":"Galaxy Ace Plus"},
        {"vendor":"Samsung","model":"Galaxy Ace S5830"},
        {"vendor":"Samsung","model":"Galaxy Beam"},
        {"vendor":"Samsung","model":"Galaxy Fit S5670"},
        {"vendor":"Samsung","model":"Galaxy Gio S5660"},
        {"vendor":"Samsung","model":"Galaxy Grand GT-I9080"},
        {"vendor":"Samsung","model":"Galaxy Grand GT-I9082"},
        {"vendor":"Samsung","model":"Galaxy Mini 2"},
        {"vendor":"Samsung","model":"Galaxy Mini S5570"},
        {"vendor":"Samsung","model":"Galaxy Music"},
        {"vendor":"Samsung","model":"Galaxy Music Dual"},
        {"vendor":"Samsung","model":"Galaxy Nexus"},
        {"vendor":"Samsung","model":"Galaxy Note"},
        {"vendor":"Samsung","model":"Galaxy Note II"},
        {"vendor":"Samsung","model":"Galaxy Pocket"},
        {"vendor":"Samsung","model":"Galaxy Premier"},
        {"vendor":"Samsung","model":"Galaxy Pro"},
        {"vendor":"Samsung","model":"Galaxy R I9103"},
        {"vendor":"Samsung","model":"Galaxy S 4G"},
        {"vendor":"Samsung","model":"Galaxy S Advance"},
        {"vendor":"Samsung","model":"Galaxy S Duos S7562"},
        {"vendor":"Samsung","model":"Galaxy S i9000"},
        {"vendor":"Samsung","model":"Galaxy S II"},
        {"vendor":"Samsung","model":"Galaxy S II LTE"},
        {"vendor":"Samsung","model":"Galaxy S II WiMAX ISW11SC"},
        {"vendor":"Samsung","model":"Galaxy S III"},
        {"vendor":"Samsung","model":"Galaxy S III Mini"},
        {"vendor":"Samsung","model":"Galaxy SL I9003"},
        {"vendor":"Samsung","model":"Galaxy Spica I5700"},
        {"vendor":"Samsung","model":"Galaxy Tab P1000"},
        {"vendor":"Samsung","model":"Galaxy W"},
        {"vendor":"Samsung","model":"Galaxy Xcover GT-S5690"},
        {"vendor":"Samsung","model":"Galaxy Y Duos"},
        {"vendor":"Samsung","model":"Galaxy Y Pro"},
        {"vendor":"Samsung","model":"Galaxy Y Pro Duos"},
        {"vendor":"Samsung","model":"Gem I100"},
        {"vendor":"Samsung","model":"Google Nexus S I9020"},
        {"vendor":"Samsung","model":"I5500 Galaxy 5"},
        {"vendor":"Samsung","model":"I5800 Galaxy 3 Apollo"},
        {"vendor":"Samsung","model":"I7500 Galaxy"},
        {"vendor":"Samsung","model":"i8510 Innov8"},
        {"vendor":"Samsung","model":"Infuse 4G i997"},
        {"vendor":"Samsung","model":"Instinct HD"},
        {"vendor":"Samsung","model":"Intercept M910"},
        {"vendor":"Samsung","model":"Jet S8000"},
        {"vendor":"Samsung","model":"L700i"},
        {"vendor":"Samsung","model":"M2710"},
        {"vendor":"Samsung","model":"M5650 Lindy"},
        {"vendor":"Samsung","model":"Omnia 7"},
        {"vendor":"Samsung","model":"Omnia HD"},
        {"vendor":"Samsung","model":"Omnia i900"},
        {"vendor":"Samsung","model":"Omnia II I8000"},
        {"vendor":"Samsung","model":"Omnia Pro B7610"},
        {"vendor":"Samsung","model":"Omnia W"},
        {"vendor":"Samsung","model":"Pixon12 M8910"},
        {"vendor":"Samsung","model":"Rugby II A847"},
        {"vendor":"Samsung","model":"S5600"},
        {"vendor":"Samsung","model":"S5620 Monte"},
        {"vendor":"Samsung","model":"S7550 Blue Earth"},
        {"vendor":"Samsung","model":"SGH-A800"},
        {"vendor":"Samsung","model":"SGH-C100"},
        {"vendor":"Samsung","model":"SGH-D500"},
        {"vendor":"Samsung","model":"SGH-D600"},
        {"vendor":"Samsung","model":"SGH-D600E"},
        {"vendor":"Samsung","model":"SGH-D840"},
        {"vendor":"Samsung","model":"SGH-D900"},
        {"vendor":"Samsung","model":"SGH-E250"},
        {"vendor":"Samsung","model":"SGH-E700"},
        {"vendor":"Samsung","model":"SGH-E730"},
        {"vendor":"Samsung","model":"SGH-E800"},
        {"vendor":"Samsung","model":"SGH-F300"},
        {"vendor":"Samsung","model":"SGH-F330"},
        {"vendor":"Samsung","model":"SGH-G600"},
        {"vendor":"Samsung","model":"SGH-G800"},
        {"vendor":"Samsung","model":"SGH-i320"},
        {"vendor":"Samsung","model":"SGH-i600"},
        {"vendor":"Samsung","model":"SGH-M300"},
        {"vendor":"Samsung","model":"SGH-P730"},
        {"vendor":"Samsung","model":"SGH-S100"},
        {"vendor":"Samsung","model":"SGH-S300"},
        {"vendor":"Samsung","model":"SGH-T100"},
        {"vendor":"Samsung","model":"SGH-U600"},
        {"vendor":"Samsung","model":"SGH-U700"},
        {"vendor":"Samsung","model":"SGH-X100"},
        {"vendor":"Samsung","model":"SGH-X450"},
        {"vendor":"Samsung","model":"SGH-X640"},
        {"vendor":"Samsung","model":"SGH-Z500"},
        {"vendor":"Samsung","model":"Shark S5350"},
        {"vendor":"Samsung","model":"Showcase i500"},
        {"vendor":"Samsung","model":"Star 3"},
        {"vendor":"Samsung","model":"Star 3 DUOS"},
        {"vendor":"Samsung","model":"Star II S5260"},
        {"vendor":"Samsung","model":"Strive A687"},
        {"vendor":"Samsung","model":"Sunburst A697"},
        {"vendor":"Samsung","model":"T-mobile Sidekick 4G"},
        {"vendor":"Samsung","model":"Tocco Lite S5230"},
        {"vendor":"Samsung","model":"Transform M920"},
        {"vendor":"Samsung","model":"U900 Soul"},
        {"vendor":"Samsung","model":"Ultra B S7220"},
        {"vendor":"Samsung","model":"Ultra S S7350"},
        {"vendor":"Samsung","model":"Ultra Touch S8300"},
        {"vendor":"Samsung","model":"Wave 3"},
        {"vendor":"Samsung","model":"Wave 575 S5750"},
        {"vendor":"Samsung","model":"Wave II S8530"},
        {"vendor":"Samsung","model":"Wave M"},
        {"vendor":"Samsung","model":"Wave S8500"},
        {"vendor":"Samsung","model":"Wave Y S5380"},
        {"vendor":"Sendo","model":"X"},
        {"vendor":"Sharp","model":"Aquos 007SH"},
        {"vendor":"Sharp","model":"GX10"},
        {"vendor":"Siemens","model":"A55"},
        {"vendor":"Siemens","model":"C60"},
        {"vendor":"Siemens","model":"C62"},
        {"vendor":"Siemens","model":"C65"},
        {"vendor":"Siemens","model":"C75"},
        {"vendor":"Siemens","model":"CX65"},
        {"vendor":"Siemens","model":"M50"},
        {"vendor":"Siemens","model":"M65"},
        {"vendor":"Siemens","model":"S40"},
        {"vendor":"Siemens","model":"S55"},
        {"vendor":"Siemens","model":"S65"},
        {"vendor":"Siemens","model":"SF65"},
        {"vendor":"Siemens","model":"SK65"}
    ]`

	// Urainkan JSON ke dalam slice of map[string]string
	var data []map[string]string
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		log.Fatal(err)
	}

	// Looping melalui data dan tampilkan hanya dari vendor Samsung
	for _, item := range data {
		vendor := item["vendor"]
		if strings.ToLower(vendor) == "samsung" {
			fmt.Println("Model:", item["model"])
			fmt.Println("Vendor:", vendor)
			fmt.Println("Price:", item["price"])
			fmt.Println()
		}
	}

	// Data JSON Sony yang akan diuraikan
	jsonDataSony := `[
			{"vendor": "Siemens", "model": "SL42"},
			{"vendor": "Siemens", "model": "SL45i"},
			{"vendor": "Siemens", "model": "SL55"},
			{"vendor": "Siemens", "model": "SX1"},
			{"vendor": "Sony", "model": "CMD-Z5"},
			{"vendor": "Sony", "model": "Xperia Acro S"},
			{"vendor": "Sony", "model": "Xperia Advance"},
			{"vendor": "Sony", "model": "Xperia E"},
			{"vendor": "Sony", "model": "Xperia E Dual"},
			{"vendor": "Sony", "model": "Xperia Go"},
			{"vendor": "Sony", "model": "Xperia Ion LT28at"},
			{"vendor": "Sony", "model": "Xperia J"},
			{"vendor": "Sony", "model": "Xperia Miro"},
			{"vendor": "Sony", "model": "Xperia Neo L"},
			{"vendor": "Sony", "model": "Xperia P"},
			{"vendor": "Sony", "model": "Xperia S"},
			{"vendor": "Sony", "model": "Xperia SL"},
			{"vendor": "Sony", "model": "Xperia Sola"},
			{"vendor": "Sony", "model": "Xperia SX"},
			{"vendor": "Sony", "model": "Xperia T"},
			{"vendor": "Sony", "model": "Xperia Tipo"},
			{"vendor": "Sony", "model": "Xperia Tipo Dual"},
			{"vendor": "Sony", "model": "Xperia TL"},
			{"vendor": "Sony", "model": "Xperia TX"},
			{"vendor": "Sony", "model": "Xperia U"},
			{"vendor": "Sony", "model": "Xperia V"},
			{"vendor": "Sony", "model": "Xperia VL"},
			{"vendor": "Sony", "model": "Ericsson Aino"},
			{"vendor": "Sony", "model": "Ericsson Aspen"},
			{"vendor": "Sony", "model": "Ericsson C510"},
			{"vendor": "Sony", "model": "Ericsson C702"},
			{"vendor": "Sony", "model": "Ericsson C901"},
			{"vendor": "Sony", "model": "Ericsson C902"},
			{"vendor": "Sony", "model": "Ericsson C903"},
			{"vendor": "Sony", "model": "Ericsson C905"},
			{"vendor": "Sony", "model": "Ericsson Cedar J108i"},
			{"vendor": "Sony", "model": "Ericsson D750i"},
			{"vendor": "Sony", "model": "Ericsson Elm J10"},
			{"vendor": "Sony", "model": "Ericsson Elm J10i2"},
			{"vendor": "Sony", "model": "Ericsson Equinox T717"},
			{"vendor": "Sony", "model": "Ericsson F305"},
			{"vendor": "Sony", "model": "Ericsson F500i"},
			{"vendor": "Sony", "model": "Ericsson G502"},
			{"vendor": "Sony", "model": "Ericsson G700"},
			{"vendor": "Sony", "model": "Ericsson G705"},
			{"vendor": "Sony", "model": "Ericsson G900"},
			{"vendor": "Sony", "model": "Ericsson Hazel J20"},
			{"vendor": "Sony", "model": "Ericsson Hazel J20i"},
			{"vendor": "Sony", "model": "Ericsson J100"},
			{"vendor": "Sony", "model": "Ericsson J110"},
			{"vendor": "Sony", "model": "Ericsson J120"},
			{"vendor": "Sony", "model": "Ericsson J132"},
			{"vendor": "Sony", "model": "Ericsson J200i"},
			{"vendor": "Sony", "model": "Ericsson J210"},
			{"vendor": "Sony", "model": "Ericsson J220"},
			{"vendor": "Sony", "model": "Ericsson J230"},
			{"vendor": "Sony", "model": "Ericsson J300i"},
			{"vendor": "Sony", "model": "Ericsson Jalou"},
			{"vendor": "Sony", "model": "Ericsson Jalou Dolce Gabbana"},
			{"vendor": "Sony", "model": "Ericsson K200"},
			{"vendor": "Sony", "model": "Ericsson K220"},
			{"vendor": "Sony", "model": "Ericsson K300"},
			{"vendor": "Sony", "model": "Ericsson K310"},
			{"vendor": "Sony", "model": "Ericsson K320"},
			{"vendor": "Sony", "model": "Ericsson K330"},
			{"vendor": "Sony", "model": "Ericsson K500"},
			{"vendor": "Sony", "model": "Ericsson K510"},
			{"vendor": "Sony", "model": "Ericsson K530"},
			{"vendor": "Sony", "model": "Ericsson K550"},
			{"vendor": "Sony", "model": "Ericsson K600"},
			{"vendor": "Sony", "model": "Ericsson K608"},
			{"vendor": "Sony", "model": "Ericsson K610"},
			{"vendor": "Sony", "model": "Ericsson K618"},
			{"vendor": "Sony", "model": "Ericsson K630"},
			{"vendor": "Sony", "model": "Ericsson K660"},
			{"vendor": "Sony", "model": "Ericsson K700"},
			{"vendor": "Sony", "model": "Ericsson K750"},
			{"vendor": "Sony", "model": "Ericsson K770"},
			{"vendor": "Sony", "model": "Ericsson K790"},
			{"vendor": "Sony", "model": "Ericsson K800"},
			{"vendor": "Sony", "model": "Ericsson K810"},
			{"vendor": "Sony", "model": "Ericsson K850"},
			{"vendor": "Sony", "model": "Ericsson Kita"},
			{"vendor": "Sony", "model": "Ericsson Live Walkman"},
			{"vendor": "Sony", "model": "Ericsson M600"},
			{"vendor": "Sony", "model": "Ericsson Mix Walkman"},
			{"vendor": "Sony", "model": "Ericsson Naite"},
			{"vendor": "Sony", "model": "Ericsson P1"},
			{"vendor": "Sony", "model": "Ericsson P800"},
			{"vendor": "Sony", "model": "Ericsson P900"},
			{"vendor": "Sony", "model": "Ericsson P910"},
			{"vendor": "Sony", "model": "Ericsson P910a"},
			{"vendor": "Sony", "model": "Ericsson P990"},
			{"vendor": "Sony", "model": "Ericsson R300"},
			{"vendor": "Sony", "model": "Ericsson R306"},
			{"vendor": "Sony", "model": "Ericsson R306a"},
			{"vendor": "Sony", "model": "Ericsson S302"},
			{"vendor": "Sony", "model": "Ericsson S312"},
			{"vendor": "Sony", "model": "Ericsson S500"},
			{"vendor": "Sony", "model": "Ericsson S600"},
			{"vendor": "Sony", "model": "Ericsson S700i"},
			{"vendor": "Sony", "model": "Ericsson S710a"},
			{"vendor": "Sony", "model": "Ericsson Satio"},
			{"vendor": "Sony", "model": "Ericsson Spiro"},
			{"vendor": "Sony", "model": "Ericsson T100"},
			{"vendor": "Sony", "model": "Ericsson T200"},
			{"vendor": "Sony", "model": "Ericsson T226"},
			{"vendor": "Sony", "model": "Ericsson T230"},
			{"vendor": "Sony", "model": "Ericsson T250"},
			{"vendor": "Sony", "model": "Ericsson T270"},
			{"vendor": "Sony", "model": "Ericsson T280"},
			{"vendor": "Sony", "model": "Ericsson T280a"},
			{"vendor": "Sony", "model": "Ericsson T290"},
			{"vendor": "Sony", "model": "Ericsson T300"},
			{"vendor": "Sony", "model": "Ericsson T303"},
			{"vendor": "Sony", "model": "Ericsson T306"},
			{"vendor": "Sony", "model": "Ericsson T310"},
			{"vendor": "Sony", "model": "Ericsson T316"},
			{"vendor": "Sony", "model": "Ericsson T600"},
			{"vendor": "Sony", "model": "Ericsson T606"},
			{"vendor": "Sony", "model": "Ericsson T608"},
			{"vendor": "Sony", "model": "Ericsson T610"},
			{"vendor": "Sony", "model": "Ericsson T616"},
			{"vendor": "Sony", "model": "Ericsson T618"},
			{"vendor": "Sony", "model": "Ericsson T630"},
			{"vendor": "Sony", "model": "Ericsson T650"},
			{"vendor": "Sony", "model": "Ericsson T658"},
			{"vendor": "Sony", "model": "Ericsson T68i"},
			{"vendor": "Sony", "model": "Ericsson T700"},
			{"vendor": "Sony", "model": "Ericsson T707"},
			{"vendor": "Sony", "model": "Ericsson T715"},
			{"vendor": "Sony", "model": "Ericsson Txt"},
			{"vendor": "Sony", "model": "Ericsson Txt Pro"},
			{"vendor": "Sony", "model": "Ericsson V600"},
			{"vendor": "Sony", "model": "Ericsson V800"},
			{"vendor": "Sony", "model": "Ericsson Vivaz"},
			{"vendor": "Sony", "model": "Ericsson Vivaz Pro"},
			{"vendor": "Sony", "model": "Ericsson W200"},
			{"vendor": "Sony", "model": "Ericsson W205"},
			{"vendor": "Sony", "model": "Ericsson W205a"},
			{"vendor": "Sony", "model": "Ericsson W300"},
			{"vendor": "Sony", "model": "Ericsson W302"},
			{"vendor": "Sony", "model": "Ericsson W350"},
			{"vendor": "Sony", "model": "Ericsson W350a"},
			{"vendor": "Sony", "model": "Ericsson W380"},
			{"vendor": "Sony", "model": "Ericsson W395"},
			{"vendor": "Sony", "model": "Ericsson W508"},
			{"vendor": "Sony", "model": "Ericsson W518a"},
			{"vendor": "Sony", "model": "Ericsson W550"},
			{"vendor": "Sony", "model": "Ericsson W580"},
			{"vendor": "Sony", "model": "Ericsson W595"},
			{"vendor": "Sony", "model": "Ericsson W600"},
			{"vendor": "Sony", "model": "Ericsson W610"},
			{"vendor": "Sony", "model": "Ericsson W660"},
			{"vendor": "Sony", "model": "Ericsson W700"},
			{"vendor": "Sony", "model": "Ericsson W705"},
			{"vendor": "Sony", "model": "Ericsson W710"},
			{"vendor": "Sony", "model": "Ericsson W715"},
			{"vendor": "Sony", "model": "Ericsson W760"},
			{"vendor": "Sony", "model": "Ericsson W8 Walkman"},
			{"vendor": "Sony", "model": "Ericsson W800"},
			{"vendor": "Sony", "model": "Ericsson W810"},
			{"vendor": "Sony", "model": "Ericsson W830"},
			{"vendor": "Sony", "model": "Ericsson W850"},
			{"vendor": "Sony", "model": "Ericsson W880"},
			{"vendor": "Sony", "model": "Ericsson W890"},
			{"vendor": "Sony", "model": "Ericsson W900"},
			{"vendor": "Sony", "model": "Ericsson W902"},
			{"vendor": "Sony", "model": "Ericsson W910"},
			{"vendor": "Sony", "model": "Ericsson W950"},
			{"vendor": "Sony", "model": "Ericsson W960"},
			{"vendor": "Sony", "model": "Ericsson W980"},
			{"vendor": "Sony", "model": "Ericsson W995"},
			{"vendor": "Sony", "model": "Ericsson Xperia Active"},
			{"vendor": "Sony", "model": "Ericsson Xperia Arc"},
			{"vendor": "Sony", "model": "Ericsson Xperia Arc S"},
			{"vendor": "Sony", "model": "Ericsson Xperia Mini"},
			{"vendor": "Sony", "model": "Ericsson Xperia Mini Pro"},
			{"vendor": "Sony", "model": "Ericsson Xperia Neo"},
			{"vendor": "Sony", "model": "Ericsson Xperia Neo V"},
			{"vendor": "Sony", "model": "Ericsson Xperia Play"},
			{"vendor": "Sony", "model": "Ericsson Xperia Play 4G"},
			{"vendor": "Sony", "model": "Ericsson Xperia Pro"},
			{"vendor": "Sony", "model": "Ericsson Xperia Pureness"},
			{"vendor": "Sony", "model": "Ericsson Xperia Ray"},
			{"vendor": "Sony", "model": "Ericsson Xperia X1"},
			{"vendor": "Sony", "model": "Ericsson Xperia X10"},
			{"vendor": "Sony", "model": "Ericsson Xperia X10 Mini"},
			{"vendor": "Sony", "model": "Ericsson Xperia X10 Mini Pro"},
			{"vendor": "Sony", "model": "Ericsson Xperia X2"},
			{"vendor": "Sony", "model": "Ericsson Xperia X8"},
			{"vendor": "Sony", "model": "Ericsson Yari"},
			{"vendor": "Sony", "model": "Ericsson Yendo"},
			{"vendor": "Sony", "model": "Ericsson Z1010"},
			{"vendor": "Sony", "model": "Ericsson Z200"},
			{"vendor": "Sony", "model": "Ericsson Z250"},
			{"vendor": "Sony", "model": "Ericsson Z300"},
			{"vendor": "Sony", "model": "Ericsson Z310"},
			{"vendor": "Sony", "model": "Ericsson Z320"},
			{"vendor": "Sony", "model": "Ericsson Z500"},
			{"vendor": "Sony", "model": "Ericsson Z520"},
			{"vendor": "Sony", "model": "Ericsson Z525"},
			{"vendor": "Sony", "model": "Ericsson Z530"},
			{"vendor": "Sony", "model": "Ericsson Z550"},
			{"vendor": "Sony", "model": "Ericsson Z555"},
			{"vendor": "Sony", "model": "Ericsson Z558"},
			{"vendor": "Sony", "model": "Ericsson Z600"},
			{"vendor": "Sony", "model": "Ericsson Z610"},
			{"vendor": "Sony", "model": "Ericsson Z710"},
			{"vendor": "Sony", "model": "Ericsson Z750"},
			{"vendor": "Sony", "model": "Ericsson Z770"},
			{"vendor": "Sony", "model": "Ericsson Z780"},
			{"vendor": "Sony", "model": "Ericsson Z800i"},
			{"vendor": "Sony", "model": "Ericsson Zylo"},
			{"vendor": "Tel.Me.", "model": "T919"},
			{"vendor": "Toshiba", "model": "G450"},
			{"vendor": "Toshiba", "model": "Portege G500"},
			{"vendor": "Toshiba", "model": "Portege G710"},
			{"vendor": "Toshiba", "model": "Portege G810"},
			{"vendor": "Toshiba", "model": "Portege G910"},
			{"vendor": "Toshiba", "model": "TG01"},
			{"vendor": "Toshiba", "model": "TS21i"},
			{"vendor": "ViewSonic", "model": "ViewPad 4"},
			{"vendor": "ViewSonic", "model": "ViewPhone 3"}
		]`

	// Buat slice untuk menampung objek Smartphone
	var smartphones []Smartphone

	// Urainkan JSON ke dalam slice of objek
	if err := json.Unmarshal([]byte(jsonDataSony), &smartphones); err != nil {
		log.Fatal(err)
	}

	// Loop melalui objek Smartphone dan tampilkan hanya dari vendor Sony
	for _, phone := range smartphones {
		if strings.ToLower(phone.Vendor) == "sony" {
			fmt.Println("Model:", phone.Model)
			fmt.Println("Vendor:", phone.Vendor)
			fmt.Println()
		}
	}

	// Data JSON untuk buku
	books := []Book{
		{
			Title:       "Book 1",
			Desc:        "Description for Book 1",
			Author:      "Author 1",
			ReleaseYear: 2020,
		},
		{
			Title:       "Book 2",
			Desc:        "Description for Book 2",
			Author:      "Author 2",
			ReleaseYear: 2018,
		},
		{
			Title:       "Book 3",
			Desc:        "Description for Book 3",
			Author:      "Author 3",
			ReleaseYear: 2022,
		},
		{
			Title:       "Book 4",
			Desc:        "Description for Book 4",
			Author:      "Author 4",
			ReleaseYear: 2019,
		},
		{
			Title:       "Book 5",
			Desc:        "Description for Book 5",
			Author:      "Author 5",
			ReleaseYear: 2021,
		},
	}

	// Encode slice of objek Book ke JSON
	bookData, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}

	// Tampilkan hasil JSON untuk smartphone
	fmt.Println("JSON Data for Books:")
	fmt.Println(string(bookData))

	jariJari := 7.0
	tinggi := 10.0

	// Membuat fungsi handler untuk server web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		luas := luasLingkaran(jariJari)
		keliling := kelilingLingkaran(jariJari)
		volume := volumeTabung(jariJari, tinggi)

		// Menampilkan hasil perhitungan dalam format teks
		result := fmt.Sprintf("jariJari: %.2f, tinggi: %.2f, volume: %.2f, luas alas: %.2f, keliling alas: %.2f",
			jariJari, tinggi, volume, luas, keliling)
		fmt.Fprintf(w, result)
	})

	// Menjalankan server web pada port 8080
	http.ListenAndServe(":8181", nil)
}
