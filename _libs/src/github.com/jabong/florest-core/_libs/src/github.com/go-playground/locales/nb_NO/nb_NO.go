package nb_NO

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type nb_NO struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'nb_NO' locale
func New() locales.Translator {
	return &nb_NO{
		locale:                 "nb_NO",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                "٫",
		group:                  " ",
		minus:                  "−",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ".",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		percentSuffix:          " ",
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mar.", "apr.", "mai", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
		daysAbbreviated:        []string{"søn.", "man.", "tir.", "ons.", "tor.", "fre.", "lør."},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sø.", "ma.", "ti.", "on.", "to.", "fr.", "lø."},
		daysWide:               []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"a.m.", "p.m."},
		erasAbbreviated:        []string{"f.Kr.", "e.Kr."},
		erasNarrow:             []string{"f.Kr.", "e.Kr."},
		erasWide:               []string{"før Kristus", "etter Kristus"},
		timezones:              map[string]string{"HKT": "normaltid for Hongkong", "ACWDT": "vest-sentralaustralsk sommertid", "IST": "indisk tid", "EST": "normaltid for den nordamerikanske østkysten", "ECT": "ecuadoriansk tid", "EDT": "sommertid for den nordamerikanske østkysten", "CLST": "chilensk sommertid", "AST": "atlanterhavskystlig standardtid", "AWDT": "vestaustralsk sommertid", "MESZ": "sentraleuropeisk sommertid", "WAST": "vestafrikansk sommertid", "COST": "colombiansk sommertid", "OESZ": "østeuropeisk sommertid", "BT": "bhutansk tid", "PDT": "sommertid for den nordamerikanske Stillehavskysten", "WIB": "vestindonesisk tid", "OEZ": "østeuropeisk normaltid", "GYT": "guyansk tid", "EAT": "østafrikansk tid", "CST": "normaltid for det sentrale Nord-Amerika", "GFT": "tidssone for Fransk Guyana", "VET": "venezuelansk tid", "BOT": "boliviansk tid", "HAT": "sommertid for Newfoundland", "SGT": "singaporsk tid", "AEST": "østaustralsk normaltid", "LHDT": "sommertid for Lord Howe-øya", "JDT": "japansk sommertid", "AEDT": "østaustralsk sommertid", "LHST": "normaltid for Lord Howe-øya", "MYT": "malaysisk tid", "CHAST": "normaltid for Chatham", "AKDT": "alaskisk sommertid", "ARST": "argentinsk sommertid", "HADT": "sommertid for Hawaii og Aleutene", "WESZ": "vesteuropeisk sommertid", "WART": "vestargentinsk normaltid", "AKST": "alaskisk normaltid", "ADT": "atlanterhavskystlig sommertid", "ART": "argentinsk normaltid", "COT": "colombiansk normaltid", "ChST": "tidssone for Chamorro", "WARST": "vestargentinsk sommertid", "AWST": "vestaustralsk normaltid", "JST": "japansk normaltid", "HAST": "normaltid for Hawaii og Aleutene", "CHADT": "sommertid for Chatham", "PST": "normaltid for den nordamerikanske Stillehavskysten", "TMT": "turkmensk normaltid", "TMST": "turkmensk sommertid", "WAT": "vestafrikansk normaltid", "HNT": "normaltid for Newfoundland", "MDT": "Macau, sommertid", "HKST": "sommertid for Hongkong", "∅∅∅": "sommertid for Amazonas", "WIT": "østindonesisk tid", "CLT": "chilensk normaltid", "UYT": "uruguayansk normaltid", "MST": "Macau, standardtid", "WITA": "sentralindonesisk tid", "SAST": "sørafrikansk tid", "ACST": "sentralaustralsk normaltid", "MEZ": "sentraleuropeisk normaltid", "NZST": "newzealandsk normaltid", "NZDT": "newzealandsk sommertid", "CAT": "sentralafrikansk tid", "CDT": "sommertid for det sentrale Nord-Amerika", "SRT": "surinamsk tid", "ACWST": "vest-sentralaustralsk normaltid", "UYST": "uruguayansk sommertid", "ACDT": "sentralaustralsk sommertid", "WEZ": "vesteuropeisk normaltid", "GMT": "Greenwich middeltid"},
	}
}

// Locale returns the current translators string locale
func (nb *nb_NO) Locale() string {
	return nb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nb_NO'
func (nb *nb_NO) PluralsCardinal() []locales.PluralRule {
	return nb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nb_NO'
func (nb *nb_NO) PluralsOrdinal() []locales.PluralRule {
	return nb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nb_NO'
func (nb *nb_NO) PluralsRange() []locales.PluralRule {
	return nb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nb_NO'
func (nb *nb_NO) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nb_NO'
func (nb *nb_NO) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nb_NO'
func (nb *nb_NO) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nb *nb_NO) MonthAbbreviated(month time.Month) string {
	return nb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nb *nb_NO) MonthsAbbreviated() []string {
	return nb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nb *nb_NO) MonthNarrow(month time.Month) string {
	return nb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nb *nb_NO) MonthsNarrow() []string {
	return nb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nb *nb_NO) MonthWide(month time.Month) string {
	return nb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nb *nb_NO) MonthsWide() []string {
	return nb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nb *nb_NO) WeekdayAbbreviated(weekday time.Weekday) string {
	return nb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nb *nb_NO) WeekdaysAbbreviated() []string {
	return nb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nb *nb_NO) WeekdayNarrow(weekday time.Weekday) string {
	return nb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nb *nb_NO) WeekdaysNarrow() []string {
	return nb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nb *nb_NO) WeekdayShort(weekday time.Weekday) string {
	return nb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nb *nb_NO) WeekdaysShort() []string {
	return nb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nb *nb_NO) WeekdayWide(weekday time.Weekday) string {
	return nb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nb *nb_NO) WeekdaysWide() []string {
	return nb.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nb_NO' and handles both Whole and Real numbers based on 'v'
func (nb *nb_NO) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(nb.decimal) - 1; j >= 0; j-- {
				b = append(b, nb.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nb.group) - 1; j >= 0; j-- {
					b = append(b, nb.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nb.minus) - 1; j >= 0; j-- {
			b = append(b, nb.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nb_NO' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nb *nb_NO) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 9
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(nb.decimal) - 1; j >= 0; j-- {
				b = append(b, nb.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nb.minus) - 1; j >= 0; j-- {
			b = append(b, nb.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nb.percentSuffix...)

	b = append(b, nb.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nb_NO'
func (nb *nb_NO) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nb.currencies[currency]
	l := len(s) + len(symbol) + 7 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(nb.decimal) - 1; j >= 0; j-- {
				b = append(b, nb.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nb.group) - 1; j >= 0; j-- {
					b = append(b, nb.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(nb.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, nb.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(nb.minus) - 1; j >= 0; j-- {
			b = append(b, nb.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nb_NO'
// in accounting notation.
func (nb *nb_NO) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nb.currencies[currency]
	l := len(s) + len(symbol) + 7 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(nb.decimal) - 1; j >= 0; j-- {
				b = append(b, nb.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nb.group) - 1; j >= 0; j-- {
					b = append(b, nb.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(nb.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nb.currencyNegativePrefix[j])
		}

		for j := len(nb.minus) - 1; j >= 0; j-- {
			b = append(b, nb.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(nb.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, nb.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nb.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nb_NO'
func (nb *nb_NO) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
