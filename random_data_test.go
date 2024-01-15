package randomdata

import (
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestCustomRand(t *testing.T) {
	r1 := rand.New(rand.NewSource(1))
	r2 := rand.New(rand.NewSource(1))

	CustomRand(r1)
	s1 := RandStringRunes(10)
	CustomRand(r2)
	s2 := RandStringRunes(10)

	assert.Equal(t, s1, s2)
}

func TestTitle(t *testing.T) {
	titleMale := Title(Male)
	titleFemale := Title(Female)
	randomTitle := Title(100)

	assert.Contains(t, jsonData.MaleTitles, titleMale, "titleMale empty or not in male titles")
	assert.Contains(t, jsonData.FemaleTitles, titleFemale, "firstNameFemale empty or not in female titles")

	names := make([]string, len(jsonData.MaleTitles)+len(jsonData.FemaleTitles))
	names = append(names, jsonData.MaleTitles...)
	names = append(names, jsonData.FemaleTitles...)
	assert.Contains(t, names, randomTitle, "randomName empty or not in male and female titles")
}

func TestRandomStringDigits(t *testing.T) {
	assert.Len(t, StringNumber(2, "-"), 5)
	assert.Len(t, StringNumber(2, ""), 4)
	assert.Len(t, StringNumberExt(3, "/", 3), 11)
	assert.Len(t, StringNumberExt(3, "", 3), 9)
}

func TestFirstName(t *testing.T) {
	firstNameMale := FirstName(Male)
	firstNameFemale := FirstName(Female)
	randomName := FirstName(RandomGender)

	assert.Contains(t, jsonData.FirstNamesMale, firstNameMale, "firstNameMale empty or not in male names")
	assert.Contains(t, jsonData.FirstNamesFemale, firstNameFemale, "firstNameFemale empty or not in female names")
	assert.NotEmpty(t, randomName)
}

func TestLastName(t *testing.T) {
	assert.Contains(t, jsonData.LastNames, LastName(), "lastName empty or not in slice")
}

func TestFullName(t *testing.T) {
	fullNameMale := FullName(Male)
	fullNameFemale := FullName(Female)
	fullNameRandom := FullName(RandomGender)

	maleSplit := strings.Fields(fullNameMale)
	femaleSplit := strings.Fields(fullNameFemale)
	randomSplit := strings.Fields(fullNameRandom)

	assert.NotEmpty(t, maleSplit, "failed on full name male")
	assert.Contains(t, jsonData.FirstNamesMale, maleSplit[0], "couldnt find maleSplit first name in firstNamesMale")
	assert.Contains(t, jsonData.LastNames, maleSplit[1], "couldnt find maleSplit last name in lastNames")

	assert.NotEmpty(t, femaleSplit, "failed on full name female")
	assert.Contains(t, jsonData.FirstNamesFemale, femaleSplit[0], "couldnt find femaleSplit first name in firstNamesFemale")
	assert.Contains(t, jsonData.LastNames, femaleSplit[1], "couldnt find femaleSplit last name in lastNames")

	assert.NotEmpty(t, randomSplit, "failed on full name random")

	firstNames := append(append([]string{}, jsonData.FirstNamesMale...), jsonData.FirstNamesFemale...)
	assert.Contains(t, firstNames, randomSplit[0], "couldnt find randomSplit first name in either firstNamesMale or firstNamesFemale")
}

func TestEmail(t *testing.T) {
	assert.NotEmpty(t, Email(), "failed to generate email with content")
}

func TestCountry(t *testing.T) {
	countryFull := Country(FullCountry)
	countryTwo := Country(TwoCharCountry)
	countryThree := Country(ThreeCharCountry)

	assert.GreaterOrEqual(t, len(countryThree), 3, "countryThree < 3 chars")
	assert.Contains(t, jsonData.Countries, countryFull, "couldnt find country in countries")
	assert.Contains(t, jsonData.CountriesTwoChars, countryTwo, "couldnt find country with two chars in countriesTwoChars")
	assert.Contains(t, jsonData.CountriesThreeChars, countryThree, "couldnt find country with three chars in countriesThreeChars")
}

func TestCurrency(t *testing.T) {
	assert.Contains(t, jsonData.Currencies, Currency(), "could not find currency in currencies")
}

func TestCity(t *testing.T) {
	assert.Contains(t, jsonData.Cities, City(), "couldnt find city in cities")
}

func TestParagraph(t *testing.T) {
	assert.Contains(t, jsonData.Paragraphs, Paragraph(), "couldnt find paragraph in paragraphs")
}

func TestAlphanumeric(t *testing.T) {
	alphanumric := Alphanumeric(10)
	assert.Len(t, alphanumric, 10, "alphanumric has wrong size")

	re := regexp.MustCompile(`^[[:alnum:]]+$`)
	assert.True(t, re.MatchString(alphanumric), "alphanumric contains invalid character")
}

func TestBool(t *testing.T) {
	booleanVal := Boolean()
	assert.True(t, booleanVal == true || booleanVal == false, "bool was wrong format")
}

func TestState(t *testing.T) {
	assert.Contains(t, jsonData.StatesSmall, State(Small), "couldnt find small state name in states")
	assert.Contains(t, jsonData.States, State(Large), "couldnt find state name in states")
}

func TestNoun(t *testing.T) {
	assert.NotEmpty(t, jsonData.Nouns)
	assert.Contains(t, jsonData.Nouns, Noun(), "couldnt find noun in json data")
}

func TestAdjective(t *testing.T) {
	assert.NotEmpty(t, jsonData.Adjectives)
	assert.Contains(t, jsonData.Adjectives, Adjective(), "couldnt find noun in json data")
}

func TestSillyName(t *testing.T) {
	assert.NotEmpty(t, SillyName(), "couldnt generate a silly name")
}

func TestIpV4Address(t *testing.T) {
	ipAddress := IpV4Address()

	ipBlocks := strings.Split(ipAddress, ".")
	assert.GreaterOrEqual(t, len(ipBlocks), 0, "invalid generated IP address")
	assert.LessOrEqual(t, len(ipBlocks), 4, "invalid generated IP address")

	for _, blockString := range ipBlocks {
		blockNumber, err := strconv.Atoi(blockString)
		assert.NoError(t, err, "error while testing IpV4Address()")
		assert.GreaterOrEqual(t, blockNumber, 0, "invalid generated IP address")
		assert.LessOrEqual(t, blockNumber, 255, "invalid generated IP address")
	}
}

func TestIpV6Address(t *testing.T) {
	ipAddress := net.ParseIP(IpV6Address())
	assert.Len(t, ipAddress, net.IPv6len, "invalid generated IPv6 address %v", ipAddress)

	roundTripIP := net.ParseIP(ipAddress.String())
	assert.NotNil(t, roundTripIP, "invalid generated IPv6 address %v", ipAddress)
	assert.True(t, net.IP.Equal(ipAddress, roundTripIP), "invalid generated IPv6 address %v", ipAddress)
}

func TestMacAddress(t *testing.T) {
	mac := MacAddress()
	assert.Len(t, mac, 17, "invalid generated Mac address %v", mac)
	assert.True(t, regexp.MustCompile(`([0-9a-fa-f]{2}[:-]){5}([0-9a-fa-f]{2})`).MatchString(mac), "invalid generated Mac address %v", mac)
}

func TestDecimal(t *testing.T) {
	d := Decimal(2, 4, 3)
	assert.GreaterOrEqual(t, d, 2.0, "invalid generate range")
	assert.LessOrEqual(t, d, 4.0, "invalid generate range")

	ds := strings.Split(strconv.FormatFloat(d, 'f', 3, 64), ".")
	assert.Len(t, ds[1], 3, "invalid floating point")
}

func TestDay(t *testing.T) {
	assert.Contains(t, jsonData.Days, Day(), "couldnt find day in days")
}

func TestMonth(t *testing.T) {
	assert.Contains(t, jsonData.Months, Month(), "couldnt find month in months")
}

func TestStringSample(t *testing.T) {
	list := []string{"str1", "str2", "str3"}
	str := StringSample(list...)
	assert.Equal(t, reflect.TypeOf(str).String(), "string", "didn't get a string object")
	assert.Contains(t, list, str, "didn't get string from sample list")
}

func TestStringSampleEmptyList(t *testing.T) {
	str := StringSample()
	assert.Equal(t, reflect.TypeOf(str).String(), "string", "didn't get a string object")
	assert.Empty(t, str, "didn't get empty string for empty sample list")
}

func TestFullDate(t *testing.T) {
	fulldateOne := FullDate()
	fulldateTwo := FullDate()

	_, err := time.Parse(DateOutputLayout, fulldateOne)
	assert.NoError(t, err, "invalid random full date")

	_, err = time.Parse(DateOutputLayout, fulldateTwo)
	assert.NoError(t, err, "invalid random full date")

	assert.NotEqual(t, fulldateOne, fulldateTwo, "generated same full date twice in a row")
}

func TestFullDatePenetration(t *testing.T) {
	for i := 0; i < 100000; i += 1 {
		d := FullDate()
		_, err := time.Parse(DateOutputLayout, d)
		assert.NoError(t, err, "invalid random full date")
	}
}

func TestFullDateInRangeNoArgs(t *testing.T) {
	fullDate := FullDateInRange()
	_, err := time.Parse(DateOutputLayout, fullDate)
	assert.NoError(t, err, "didn't get valid date format")
}

func TestFullDateInRangeOneArg(t *testing.T) {
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 10000; i++ {
		fullDate := FullDateInRange("2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)
		assert.NoError(t, err, "didn't get valid date format")
		assert.False(t, d.After(maxDate), "random date didn't match specified max date")
	}
}

func TestFullDateInRangeTwoArgs(t *testing.T) {
	minDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 10000; i++ {
		fullDate := FullDateInRange("2016-01-01", "2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)
		assert.NoError(t, err, "didn't get valid date format")
		assert.False(t, d.After(maxDate), "random date didn't match specified max date")
		assert.False(t, d.Before(minDate), "random date didn't match specified min date")
	}
}

func TestFullDateInRangeSwappedArgs(t *testing.T) {
	wrongMaxDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	fullDate := FullDateInRange("2016-12-31", "2016-01-01")
	d, err := time.Parse(DateOutputLayout, fullDate)
	assert.NoError(t, err, "didn't get valid date format")
	assert.Equal(t, d, wrongMaxDate, "didn't return min date")
}

func TestTimezone(t *testing.T) {
	timezone := Timezone()
	assert.Contains(t, jsonData.Timezones, timezone, "couldnt find timezone in timezones: %v", timezone)
}

func TestLocale(t *testing.T) {
	locale := Locale()
	_, err := language.Parse(locale)
	assert.NoError(t, err, "invalid locale: %v", locale)
}

func TestLocalePenetration(t *testing.T) {
	for i := 0; i < 10000; i += 1 {
		locale := Locale()
		_, err := language.Parse(locale)
		assert.NoError(t, err, "invalid locale: %v", locale)
	}
}

func TestUserAgentString(t *testing.T) {
	ua := UserAgentString()
	assert.NotEmpty(t, ua, "empty User Agent String")
	assert.True(t, regexp.MustCompile(`^[a-zA-Z]+\/[0-9]+.[0-9]+\ \(.*\).*$`).MatchString(ua),
		"invalid generated User Agent String: %v", ua)
}

func TestPhoneNumbers(t *testing.T) {
	CheckPhoneNumber(PhoneNumber(), t)
}

func CheckPhoneNumber(str string, t *testing.T, msgAndArgs ...interface{}) {
	assert.LessOrEqual(t, len(str)-strings.Count(str, " "), 16, msgAndArgs)

	matched, err := regexp.MatchString("\\+\\d{1,3}\\s\\d{1,3}", str)
	assert.NoError(t, err, msgAndArgs)
	assert.True(t, matched, msgAndArgs)
}

func TestProvinceForCountry(t *testing.T) {
	supportedCountries := []string{"US", "GB"}
	for _, c := range supportedCountries {
		p := ProvinceForCountry(c)
		assert.NotEmpty(t, p, "did not return a valid province for country %s", c)
		switch c {
		case "US":
			assert.Contains(t, jsonData.States, p, "did not return a known province for US")
		case "GB":
			assert.Contains(t, jsonData.ProvincesGB, p, "did not return a known province for GB")
		}
	}
	assert.Empty(t, ProvinceForCountry("bogus"), "did not return empty province for unknown country")
}

func TestStreetForCountry(t *testing.T) {
	supportedCountries := []string{"US", "GB"}
	for _, c := range supportedCountries {
		p := StreetForCountry(c)
		assert.NotEmpty(t, p, "did not return a valid street for country %s", c)
	}
	assert.Empty(t, StreetForCountry("bogus"), "did not return empty street for unknown country")
}
