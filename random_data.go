// Package randomdata implements a bunch of simple ways to generate (pseudo) random data
package randomdata

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

const (
	Male         int = 0
	Female       int = 1
	RandomGender int = 2
)

const (
	Small int = 0
	Large int = 1
)

const (
	FullCountry      = 0
	TwoCharCountry   = 1
	ThreeCharCountry = 2
)

const (
	DateInputLayout  = "2006-01-02"
	DateOutputLayout = "Monday 2 Jan 2006"
)

const ALPHANUMERIC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type jsonContent struct {
	Adjectives          []string `json:"adjectives"`
	Nouns               []string `json:"nouns"`
	FirstNamesFemale    []string `json:"firstNamesFemale"`
	FirstNamesMale      []string `json:"firstNamesMale"`
	LastNames           []string `json:"lastNames"`
	Domains             []string `json:"domains"`
	People              []string `json:"people"`
	StreetTypes         []string `json:"streetTypes"` // Taken from https://github.com/tomharris/random_data/blob/master/lib/random_data/locations.rb
	Paragraphs          []string `json:"paragraphs"`  // Taken from feedbooks.com and www.gutenberg.org
	Countries           []string `json:"countries"`   // Fetched from the world bank at http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
	CountriesThreeChars []string `json:"countriesThreeChars"`
	CountriesTwoChars   []string `json:"countriesTwoChars"`
	Currencies          []string `json:"currencies"` //https://github.com/OpenBookPrices/country-data
	Cities              []string `json:"cities"`
	States              []string `json:"states"`
	StatesSmall         []string `json:"statesSmall"`
	Days                []string `json:"days"`
	Months              []string `json:"months"`
	FemaleTitles        []string `json:"femaleTitles"`
	MaleTitles          []string `json:"maleTitles"`
	Timezones           []string `json:"timezones"`           // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Locales             []string `json:"locales"`             // https://tools.ietf.org/html/bcp47
	UserAgents          []string `json:"userAgents"`          // http://techpatterns.com/downloads/firefox/useragentswitcher.xml
	CountryCallingCodes []string `json:"countryCallingCodes"` // from https://github.com/datasets/country-codes/blob/master/data/country-codes.csv
	ProvincesGB         []string `json:"provincesGB"`
	StreetNameGB        []string `json:"streetNameGB"`
	StreetTypesGB       []string `json:"streetTypesGB"`
}

// Rand is a source of random numbers.
type Rand struct {
	pr *rand.Rand
	mu *sync.Mutex
}

// FromSeed creates a new source of random numbers using a seed.
func FromSeed(seed int64) *Rand {
	return &Rand{
		pr: rand.New(rand.NewSource(seed)),
		mu: &sync.Mutex{},
	}
}

// FromRand creates a new source of random numbers from a rand.Rand.
func FromRand(randToUse *rand.Rand) *Rand {
	return &Rand{
		pr: randToUse,
		mu: &sync.Mutex{},
	}
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pr.Intn(n)
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func (r *Rand) Float64() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pr.Float64()
}

// Boolean returns randomly either true or false.
func (r *Rand) Boolean() bool {
	nr := r.Intn(2)
	return nr != 0
}

// Duration returns a random duration between 0 and the specified max duration.
func (r *Rand) Duration(maxDuration time.Duration) time.Duration {
	n := int64(maxDuration)
	randN := r.pr.Int63n(n)
	return time.Duration(randN)
}

// Time returns a random time between the given time and within a given duration time range.
func (r *Rand) Time(t time.Time, timeRange time.Duration) time.Time {
	return t.Add(r.Duration(timeRange))
}

// TimeRange returns a random time interval between the given time and within a given time range.
func (r *Rand) TimeRange(t time.Time, timeRange time.Duration) (time.Time, time.Time) {
	d1 := r.Duration(timeRange)
	d2 := r.Duration(timeRange - d1)
	time1 := t.Add(d1)
	time2 := time1.Add(d2)
	return time1, time2
}

// RandStringRunes generates random runes.
func (r *Rand) RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

var jsonData = jsonContent{}

func init() {
	jsonData = jsonContent{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

// StringFrom returns a random element of a slice.
func (r *Rand) StringFrom(source []string) string {
	if len(source) == 0 {
		return ""
	}
	return source[r.Intn(len(source))]
}

// Title returns a random title, gender decides the gender of the name.
func (r *Rand) Title(gender int) string {
	switch gender {
	case Male:
		return r.StringFrom(jsonData.MaleTitles)
	case Female:
		return r.StringFrom(jsonData.FemaleTitles)
	default:
		return r.Title(r.Intn(2))
	}
}

// FirstName returns a random first name, gender decides the gender of the name.
func (r *Rand) FirstName(gender int) string {
	var name = ""
	switch gender {
	case Male:
		name = r.StringFrom(jsonData.FirstNamesMale)
	case Female:
		name = r.StringFrom(jsonData.FirstNamesFemale)
	default:
		name = r.FirstName(rand.Intn(2))
	}
	return name
}

// LastName returns a random last name.
func (r *Rand) LastName() string {
	return r.StringFrom(jsonData.LastNames)
}

// FullName returns a combination of FirstName LastName randomized, gender decides the gender of the name.
func (r *Rand) FullName(gender int) string {
	return r.FirstName(gender) + " " + r.LastName()
}

// Email returns a random email.
func (r *Rand) Email() string {
	return r.createEmail(r.FirstName(RandomGender), r.LastName())
}

func (r *Rand) createEmail(firstName, lastName string) string {
	return strings.ToLower(firstName+"."+lastName) + r.StringNumberExt(1, "", 3) + "@" + r.StringFrom(jsonData.Domains)
}

// Country returns a random country, countryStyle decides what kind of format the returned country will have.
func (r *Rand) Country(countryStyle int64) string {
	country := ""
	switch countryStyle {
	case FullCountry:
		country = r.StringFrom(jsonData.Countries)
	case TwoCharCountry:
		country = r.StringFrom(jsonData.CountriesTwoChars)
	case ThreeCharCountry:
		country = r.StringFrom(jsonData.CountriesThreeChars)
	default:
	}
	return country
}

// Currency returns a random currency under ISO 4217 format.
func (r *Rand) Currency() string {
	return r.StringFrom(jsonData.Currencies)
}

// City returns a random city.
func (r *Rand) City() string {
	return r.StringFrom(jsonData.Cities)
}

// ProvinceForCountry returns a randomly selected province (state, county,subdivision ) name for a supplied country.
// If the country is not supported it will return an empty string.
func (r *Rand) ProvinceForCountry(countrycode string) string {
	switch countrycode {
	case "US":
		return r.StringFrom(jsonData.States)
	case "GB":
		return r.StringFrom(jsonData.ProvincesGB)
	}
	return ""
}

// State returns a random american state.
func (r *Rand) State(typeOfState int) string {
	if typeOfState == Small {
		return r.StringFrom(jsonData.StatesSmall)
	}
	return r.StringFrom(jsonData.States)
}

// Street returns a random fake street name.
func (r *Rand) Street() string {
	return fmt.Sprintf("%s %s", r.StringFrom(jsonData.People), r.StringFrom(jsonData.StreetTypes))
}

// StreetForCountry returns a random fake street name typical to the supplied country.
// If the country is not supported it will return an empty string.
func (r *Rand) StreetForCountry(countrycode string) string {
	switch countrycode {
	case "US":
		return r.Street()
	case "GB":
		return fmt.Sprintf("%s %s", r.StringFrom(jsonData.StreetNameGB), r.StringFrom(jsonData.StreetTypesGB))
	}
	return ""
}

// Address returns an american style address.
func (r *Rand) Address() string {
	return fmt.Sprintf("%d %s,\n%s, %s, %s", r.Number(100), r.Street(), r.City(), r.State(Small), r.PostalCode("US"))
}

// Paragraph returns a random paragraph.
func (r *Rand) Paragraph() string {
	return r.StringFrom(jsonData.Paragraphs)
}

// Number returns a random number, if only one integer (n1) is supplied it returns a number in [0,n1).
// if a second argument is supplied it returns a number in [n1,n2).
func (r *Rand) Number(numberRange ...int) int {
	nr := 0
	if len(numberRange) > 1 {
		nr = 1
		nr = r.Intn(numberRange[1]-numberRange[0]) + numberRange[0]
	} else {
		nr = r.Intn(numberRange[0])
	}
	return nr
}

func (r *Rand) Decimal(numberRange ...int) float64 {
	nr := 0.0
	if len(numberRange) > 1 {
		nr = 1.0
		nr = r.Float64()*(float64(numberRange[1])-float64(numberRange[0])) + float64(numberRange[0])
	} else {
		nr = r.Float64() * float64(numberRange[0])
	}

	if len(numberRange) > 2 {
		sf := strconv.FormatFloat(nr, 'f', numberRange[2], 64)
		nr, _ = strconv.ParseFloat(sf, 64)
	}
	return nr
}

func (r *Rand) StringNumberExt(numberPairs int, separator string, numberOfDigits int) string {
	numberString := ""
	for i := 0; i < numberPairs; i++ {
		for d := 0; d < numberOfDigits; d++ {
			numberString += fmt.Sprintf("%d", r.Number(0, 9))
		}
		if i+1 != numberPairs {
			numberString += separator
		}
	}
	return numberString
}

// StringNumber returns a random number as a string.
func (r *Rand) StringNumber(numberPairs int, separator string) string {
	return r.StringNumberExt(numberPairs, separator, 2)
}

// Alphanumeric returns a random alphanumeric string consits of [0-9a-zA-Z].
func (r *Rand) Alphanumeric(length int) string {
	list := make([]byte, length)

	for i := range list {
		list[i] = ALPHANUMERIC[r.Intn(len(ALPHANUMERIC))]
	}

	return string(list)
}

// Noun returns a random noun.
func (r *Rand) Noun() string {
	return r.StringFrom(jsonData.Nouns)
}

// Adjective returns a random adjective.
func (r *Rand) Adjective() string {
	return r.StringFrom(jsonData.Adjectives)
}

func uppercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// SillyName returns a silly name, useful for randomizing naming of things.
func (r *Rand) SillyName() string {
	return uppercaseFirstLetter(r.Noun()) + r.Adjective()
}

// IpV4Address returns a valid IPv4 address as string.
func (r *Rand) IpV4Address() string {
	blocks := []string{}
	for i := 0; i < 4; i++ {
		number := r.Intn(255)
		blocks = append(blocks, strconv.Itoa(number))
	}

	return strings.Join(blocks, ".")
}

// IpV6Address returns a valid IPv6 address as net.IP.
func (r *Rand) IpV6Address() string {
	var ip net.IP
	for i := 0; i < net.IPv6len; i++ {
		number := uint8(r.Intn(255))
		ip = append(ip, number)
	}
	return ip.String()
}

// MacAddress returns an mac address string.
func (r *Rand) MacAddress() string {
	blocks := []string{}
	for i := 0; i < 6; i++ {
		number := fmt.Sprintf("%02x", r.Intn(255))
		blocks = append(blocks, number)
	}

	return strings.Join(blocks, ":")
}

// Day returns random day.
func (r *Rand) Day() string {
	return r.StringFrom(jsonData.Days)
}

// Month returns random month.
func (r *Rand) Month() string {
	return r.StringFrom(jsonData.Months)
}

// FullDate returns full date.
func (r *Rand) FullDate() string {
	timestamp := time.Now()
	year := timestamp.Year()
	month := r.Number(1, 13)
	maxDay := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
	day := r.Number(1, maxDay+1)
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date.Format(DateOutputLayout)
}

// FullDateInRange returns a date string within a given range, given in the format "2006-01-02".
// If no argument is supplied it will return the result of randomdata.FullDate().
// If only one argument is supplied it is treated as the max date to return.
// If a second argument is supplied it returns a date between (and including) the two dates.
// Returned date is in format "Monday 2 Jan 2006".
func (r *Rand) FullDateInRange(dateRange ...string) string {
	var (
		min        time.Time
		max        time.Time
		duration   int
		dateString string
	)
	if len(dateRange) == 1 {
		max, _ = time.Parse(DateInputLayout, dateRange[0])
	} else if len(dateRange) == 2 {
		min, _ = time.Parse(DateInputLayout, dateRange[0])
		max, _ = time.Parse(DateInputLayout, dateRange[1])
	}
	if !max.IsZero() && max.After(min) {
		duration = r.Number(int(max.Sub(min))) * -1
		dateString = max.Add(time.Duration(duration)).Format(DateOutputLayout)
	} else if !max.IsZero() && !max.After(min) {
		dateString = max.Format(DateOutputLayout)
	} else {
		dateString = r.FullDate()
	}
	return dateString
}

func (r *Rand) Timezone() string {
	return r.StringFrom(jsonData.Timezones)
}

func (r *Rand) Locale() string {
	return r.StringFrom(jsonData.Locales)
}

func (r *Rand) UserAgentString() string {
	return r.StringFrom(jsonData.UserAgents)
}

func (r *Rand) PhoneNumber() string {
	str := r.StringFrom(jsonData.CountryCallingCodes) + " "

	str += r.Digits(r.Intn(3) + 1)

	for {
		// max 15 chars
		remaining := 15 - (len(str) - strings.Count(str, " "))
		if remaining < 2 {
			return "+" + str
		}
		str += " " + r.Digits(r.Intn(remaining-1)+1)
	}
}

// Letters generates a string of N random leters (A-Z).
func (r *Rand) Letters(letters int) string {
	list := make([]byte, letters)
	for i := range list {
		list[i] = byte(r.Intn('Z'-'A') + 'A')
	}
	return string(list)
}

// Digits generates a string of N random digits, padded with zeros if necessary.
func (r *Rand) Digits(digits int) string {
	max := int(math.Pow10(digits)) - 1
	num := r.Intn(max)
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, num)
}

// BoundedDigits generates a string of N random digits, padded with zeros if necessary.
// The output is restricted to the given range.
func (r *Rand) BoundedDigits(digits, low, high int) string {
	if low > high {
		low, high = high, low
	}

	max := int(math.Pow10(digits)) - 1
	if high > max {
		high = max
	}

	num := r.Intn(high-low+1) + low
	format := fmt.Sprintf("%%0%dd", digits)
	return fmt.Sprintf(format, num)
}
