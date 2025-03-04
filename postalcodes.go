package randomdata

import (
	"strings"
)

// Supported formats obtained from:
// * http://www.geopostcodes.com/GeoPC_Postal_codes_formats

// PostalCode yields a random postal/zip code for the given 2-letter country code.
//
// These codes are not guaranteed to refer to actually locations.
// They merely follow the correct format as far as letters and digits goes.
// Where possible, the function enforces valid ranges of letters and digits.
func (r *Rand) PostalCode(countrycode string) string {
	switch strings.ToUpper(countrycode) {
	case "LS", "MG", "IS", "OM", "PG":
		return r.Digits(3)

	case "AM", "GE", "NZ", "NE", "NO", "PY", "ZA", "MZ", "SJ", "LI", "AL",
		"BD", "CV", "GL":
		return r.Digits(4)

	case "DZ", "BA", "KH", "DO", "EG", "EE", "GP", "GT", "ID", "IL", "JO",
		"KW", "MQ", "MX", "LK", "SD", "TR", "UA", "US", "CR", "IQ", "KV", "MY",
		"MN", "ME", "PK", "SM", "MA", "UY", "EH", "ZM":
		return r.Digits(5)

	case "BY", "CN", "IN", "KZ", "KG", "NG", "RO", "RU", "SG", "TJ", "TM", "UZ", "VN":
		return r.Digits(6)

	case "CL":
		return r.Digits(7)

	case "IR":
		return r.Digits(10)

	case "FO":
		return "FO " + r.Digits(3)

	case "AF":
		return r.BoundedDigits(2, 10, 43) + r.BoundedDigits(2, 1, 99)

	case "AU", "AT", "BE", "BG", "CY", "DK", "ET", "GW", "HU", "LR", "MK", "PH",
		"CH", "TN", "VE":
		return r.BoundedDigits(4, 1000, 9999)

	case "SV":
		return "CP " + r.BoundedDigits(4, 1000, 9999)

	case "HT":
		return "HT" + r.Digits(4)

	case "LB":
		return r.Digits(4) + " " + r.Digits(4)

	case "LU":
		return r.BoundedDigits(4, 6600, 6999)

	case "MD":
		return "MD-" + r.BoundedDigits(4, 1000, 9999)

	case "HR":
		return "HR-" + r.Digits(5)

	case "CU":
		return "CP " + r.BoundedDigits(5, 10000, 99999)

	case "FI":
		// Last digit is usually 0 but can, in some cases, be 1 or 5.
		switch r.Intn(2) {
		case 0:
			return r.Digits(4) + "0"
		case 1:
			return r.Digits(4) + "1"
		}

		return r.Digits(4) + "5"

	case "FR", "GF", "PF", "YT", "MC", "RE", "BL", "MF", "PM", "RS", "TH":
		return r.BoundedDigits(5, 10000, 99999)

	case "DE":
		return r.BoundedDigits(5, 1000, 99999)

	case "GR":
		return r.BoundedDigits(3, 100, 999) + " " + r.Digits(2)

	case "HN":
		return "CM" + r.Digits(4)

	case "IT", "VA":
		return r.BoundedDigits(5, 10, 99999)

	case "KE":
		return r.BoundedDigits(5, 100, 99999)

	case "LA":
		return r.BoundedDigits(5, 1000, 99999)

	case "MH":
		return r.BoundedDigits(5, 96960, 96970)

	case "FM":
		return "FM" + r.BoundedDigits(5, 96941, 96944)

	case "MM":
		return r.BoundedDigits(2, 1, 14) + r.Digits(3)

	case "NP":
		return r.BoundedDigits(5, 10700, 56311)

	case "NC":
		return "98" + r.Digits(3)

	case "PW":
		return "PW96940"

	case "PR":
		return "PR " + r.Digits(5)

	case "SA":
		return r.BoundedDigits(5, 10000, 99999) + "-" + r.BoundedDigits(4, 1000, 9999)

	case "ES":
		return r.BoundedDigits(2, 1, 52) + r.BoundedDigits(3, 100, 999)

	case "WF":
		return "986" + r.Digits(2)

	case "SZ":
		return r.Letters(1) + r.Digits(3)

	case "BM":
		return r.Letters(2) + r.Digits(2)

	case "AD":
		return r.Letters(2) + r.Digits(3)

	case "BN", "AZ", "VG", "PE":
		return r.Letters(2) + r.Digits(4)

	case "BB":
		return r.Letters(2) + r.Digits(5)

	case "EC":
		return r.Letters(2) + r.Digits(6)

	case "MT":
		return r.Letters(3) + r.Digits(4)

	case "JM":
		return "JM" + r.Letters(3) + r.Digits(2)

	case "AR":
		return r.Letters(1) + r.Digits(4) + r.Letters(3)

	case "CA":
		return r.Letters(1) + r.Digits(1) + r.Letters(1) + r.Digits(1) + r.Letters(1) + r.Digits(1)

	case "FK", "TC":
		return r.Letters(4) + r.Digits(1) + r.Letters(2)

	case "GG", "IM", "JE":
		return r.Letters(2) + r.Digits(2) + r.Letters(2)

	case "GB":
		return r.Letters(2) + r.Digits(1) + " " + r.Digits(1) + r.Letters(2)

	case "KY":
		return r.Letters(2) + r.Digits(1) + "-" + r.Digits(4)

	case "JP":
		return r.Digits(3) + "-" + r.Digits(4)

	case "LV", "SI":
		return r.Letters(2) + "-" + r.Digits(4)

	case "LT":
		return r.Letters(2) + "-" + r.Digits(5)

	case "SE", "TW":
		return r.Digits(5)

	case "MV":
		return r.Digits(2) + "-" + r.Digits(2)

	case "PL":
		return r.Digits(2) + "-" + r.Digits(3)

	case "NI":
		return r.Digits(3) + "-" + r.Digits(3) + "-" + r.Digits(1)

	case "KR":
		return r.Digits(3) + "-" + r.Digits(3)

	case "PT":
		return r.Digits(4) + "-" + r.Digits(3)

	case "NL":
		return r.Digits(4) + r.Letters(2)

	case "BR":
		return r.Digits(5) + "-" + r.Digits(3)
	}

	return ""
}
