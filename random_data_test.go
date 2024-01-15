package randomdata

import (
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestRand(t *testing.T) {
	t.Run("should be created from a seed", func(t *testing.T) {
		r1 := FromSeed(1234)
		r2 := FromSeed(5678)
		f1 := r1.Float64()
		f2 := r2.Float64()
		assert.NotEqual(t, f1, f2)
	})

	t.Run("should be created from a rand.Rand", func(t *testing.T) {
		r1 := FromRand(rand.New(rand.NewSource(1234)))
		r2 := FromRand(rand.New(rand.NewSource(5678)))
		f1 := r1.Float64()
		f2 := r2.Float64()
		assert.NotEqual(t, f1, f2)
	})

	r := FromSeed(1234)

	t.Run("should generate integer between 0 and n-1", func(t *testing.T) {
		const N = 10
		value := r.Intn(N)
		assert.GreaterOrEqual(t, value, 0)
		assert.LessOrEqual(t, value, N)
	})

	t.Run("should generate float between 0 and 1", func(t *testing.T) {
		value := r.Float64()
		assert.GreaterOrEqual(t, value, 0.0)
		assert.Less(t, value, 1.0)
	})

	t.Run("should pick randomly a boolean value", func(t *testing.T) {
		booleanVal := r.Boolean()
		assert.True(t, booleanVal == true || booleanVal == false, "bool was wrong format")
	})

	t.Run("should a duration", func(t *testing.T) {
		maxDuration := 24 * time.Hour
		duration := r.Duration(maxDuration)
		assert.GreaterOrEqual(t, duration, time.Duration(0))
		assert.LessOrEqual(t, duration, maxDuration)
	})

	t.Run("should generate a time", func(t *testing.T) {
		startTime := time.Now()
		timeRange := 24 * time.Hour
		endTime := startTime.Add(timeRange)
		randTime := r.Time(startTime, timeRange)
		assert.True(t, randTime.After(startTime))
		assert.True(t, randTime.Before(endTime))
	})

	t.Run("should generate a time range", func(t *testing.T) {
		startTime := time.Now()
		timeRange := 24 * time.Hour
		endTime := startTime.Add(timeRange)
		time1, time2 := r.TimeRange(startTime, timeRange)
		assert.True(t, time1.After(startTime))
		assert.True(t, time1.Before(time2))
		assert.True(t, time2.Before(endTime))
	})

	t.Run("should pick randomly a string from a slice", func(t *testing.T) {
		list := []string{"a", "b", "c", "d"}
		elem := r.StringFrom(list)
		assert.Contains(t, list, elem)
	})

	t.Run("should return an empty string for empty slice", func(t *testing.T) {
		list := []string{}
		elem := r.StringFrom(list)
		assert.Empty(t, elem)
	})

	t.Run("should generate a title", func(t *testing.T) {
		titleMale := r.Title(Male)
		titleFemale := r.Title(Female)
		randomTitle := r.Title(RandomGender)

		assert.Contains(t, jsonData.MaleTitles, titleMale, "titleMale empty or not in male titles")
		assert.Contains(t, jsonData.FemaleTitles, titleFemale, "firstNameFemale empty or not in female titles")

		names := make([]string, len(jsonData.MaleTitles)+len(jsonData.FemaleTitles))
		names = append(names, jsonData.MaleTitles...)
		names = append(names, jsonData.FemaleTitles...)
		assert.Contains(t, names, randomTitle, "randomName empty or not in male and female titles")
	})

	t.Run("should generate a first name", func(t *testing.T) {
		firstNameMale := r.FirstName(Male)
		firstNameFemale := r.FirstName(Female)
		randomName := r.FirstName(RandomGender)

		assert.Contains(t, jsonData.FirstNamesMale, firstNameMale, "firstNameMale empty or not in male names")
		assert.Contains(t, jsonData.FirstNamesFemale, firstNameFemale, "firstNameFemale empty or not in female names")
		assert.NotEmpty(t, randomName)
	})

	t.Run("should generate a last name", func(t *testing.T) {
		assert.Contains(t, jsonData.LastNames, r.LastName(), "lastName empty or not in slice")
	})

	t.Run("should generate a full name", func(t *testing.T) {
		fullNameMale := r.FullName(Male)
		fullNameFemale := r.FullName(Female)
		fullNameRandom := r.FullName(RandomGender)

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
	})

	t.Run("should generate an email", func(t *testing.T) {
		assert.NotEmpty(t, r.Email(), "failed to generate email with content")
	})
}

func TestRandomStringDigits(t *testing.T) {
	r := FromSeed(1234)
	assert.Len(t, r.StringNumber(2, "-"), 5)
	assert.Len(t, r.StringNumber(2, ""), 4)
	assert.Len(t, r.StringNumberExt(3, "/", 3), 11)
	assert.Len(t, r.StringNumberExt(3, "", 3), 9)
}

func TestCountry(t *testing.T) {
	r := FromSeed(1234)
	countryFull := r.Country(FullCountry)
	countryTwo := r.Country(TwoCharCountry)
	countryThree := r.Country(ThreeCharCountry)

	assert.GreaterOrEqual(t, len(countryThree), 3, "countryThree < 3 chars")
	assert.Contains(t, jsonData.Countries, countryFull, "couldnt find country in countries")
	assert.Contains(t, jsonData.CountriesTwoChars, countryTwo, "couldnt find country with two chars in countriesTwoChars")
	assert.Contains(t, jsonData.CountriesThreeChars, countryThree, "couldnt find country with three chars in countriesThreeChars")
}

func TestCurrency(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.Currencies, r.Currency(), "could not find currency in currencies")
}

func TestCity(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.Cities, r.City(), "couldnt find city in cities")
}

func TestParagraph(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.Paragraphs, r.Paragraph(), "couldnt find paragraph in paragraphs")
}

func TestAlphanumeric(t *testing.T) {
	r := FromSeed(1234)
	alphanumric := r.Alphanumeric(10)
	assert.Len(t, alphanumric, 10, "alphanumric has wrong size")

	re := regexp.MustCompile(`^[[:alnum:]]+$`)
	assert.True(t, re.MatchString(alphanumric), "alphanumric contains invalid character")
}

func TestState(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.StatesSmall, r.State(Small), "couldnt find small state name in states")
	assert.Contains(t, jsonData.States, r.State(Large), "couldnt find state name in states")
}

func TestNoun(t *testing.T) {
	r := FromSeed(1234)
	assert.NotEmpty(t, jsonData.Nouns)
	assert.Contains(t, jsonData.Nouns, r.Noun(), "couldnt find noun in json data")
}

func TestAdjective(t *testing.T) {
	r := FromSeed(1234)
	assert.NotEmpty(t, jsonData.Adjectives)
	assert.Contains(t, jsonData.Adjectives, r.Adjective(), "couldnt find noun in json data")
}

func TestSillyName(t *testing.T) {
	r := FromSeed(1234)
	assert.NotEmpty(t, r.SillyName(), "couldnt generate a silly name")
}

func TestIpV4Address(t *testing.T) {
	r := FromSeed(1234)
	ipAddress := r.IpV4Address()

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
	r := FromSeed(1234)
	ipAddress := net.ParseIP(r.IpV6Address())
	assert.Len(t, ipAddress, net.IPv6len, "invalid generated IPv6 address %v", ipAddress)

	roundTripIP := net.ParseIP(ipAddress.String())
	assert.NotNil(t, roundTripIP, "invalid generated IPv6 address %v", ipAddress)
	assert.True(t, net.IP.Equal(ipAddress, roundTripIP), "invalid generated IPv6 address %v", ipAddress)
}

func TestMacAddress(t *testing.T) {
	r := FromSeed(1234)
	mac := r.MacAddress()
	assert.Len(t, mac, 17, "invalid generated Mac address %v", mac)
	assert.True(t, regexp.MustCompile(`([0-9a-fa-f]{2}[:-]){5}([0-9a-fa-f]{2})`).MatchString(mac), "invalid generated Mac address %v", mac)
}

func TestDecimal(t *testing.T) {
	r := FromSeed(1234)
	d := r.Decimal(2, 4, 3)
	assert.GreaterOrEqual(t, d, 2.0, "invalid generate range")
	assert.LessOrEqual(t, d, 4.0, "invalid generate range")

	ds := strings.Split(strconv.FormatFloat(d, 'f', 3, 64), ".")
	assert.Len(t, ds[1], 3, "invalid floating point")
}

func TestDay(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.Days, r.Day(), "couldnt find day in days")
}

func TestMonth(t *testing.T) {
	r := FromSeed(1234)
	assert.Contains(t, jsonData.Months, r.Month(), "couldnt find month in months")
}

func TestFullDate(t *testing.T) {
	r := FromSeed(1234)
	fulldateOne := r.FullDate()
	fulldateTwo := r.FullDate()

	_, err := time.Parse(DateOutputLayout, fulldateOne)
	assert.NoError(t, err, "invalid random full date")

	_, err = time.Parse(DateOutputLayout, fulldateTwo)
	assert.NoError(t, err, "invalid random full date")

	assert.NotEqual(t, fulldateOne, fulldateTwo, "generated same full date twice in a row")
}

func TestFullDatePenetration(t *testing.T) {
	r := FromSeed(1234)
	for i := 0; i < 100000; i += 1 {
		d := r.FullDate()
		_, err := time.Parse(DateOutputLayout, d)
		assert.NoError(t, err, "invalid random full date")
	}
}

func TestFullDateInRangeNoArgs(t *testing.T) {
	r := FromSeed(1234)
	fullDate := r.FullDateInRange()
	_, err := time.Parse(DateOutputLayout, fullDate)
	assert.NoError(t, err, "didn't get valid date format")
}

func TestFullDateInRangeOneArg(t *testing.T) {
	r := FromSeed(1234)
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 10000; i++ {
		fullDate := r.FullDateInRange("2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)
		assert.NoError(t, err, "didn't get valid date format")
		assert.False(t, d.After(maxDate), "random date didn't match specified max date")
	}
}

func TestFullDateInRangeTwoArgs(t *testing.T) {
	r := FromSeed(1234)
	minDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	maxDate, _ := time.Parse(DateInputLayout, "2016-12-31")
	for i := 0; i < 10000; i++ {
		fullDate := r.FullDateInRange("2016-01-01", "2016-12-31")
		d, err := time.Parse(DateOutputLayout, fullDate)
		assert.NoError(t, err, "didn't get valid date format")
		assert.False(t, d.After(maxDate), "random date didn't match specified max date")
		assert.False(t, d.Before(minDate), "random date didn't match specified min date")
	}
}

func TestFullDateInRangeSwappedArgs(t *testing.T) {
	r := FromSeed(1234)
	wrongMaxDate, _ := time.Parse(DateInputLayout, "2016-01-01")
	fullDate := r.FullDateInRange("2016-12-31", "2016-01-01")
	d, err := time.Parse(DateOutputLayout, fullDate)
	assert.NoError(t, err, "didn't get valid date format")
	assert.Equal(t, d, wrongMaxDate, "didn't return min date")
}

func TestTimezone(t *testing.T) {
	r := FromSeed(1234)
	timezone := r.Timezone()
	assert.Contains(t, jsonData.Timezones, timezone, "couldnt find timezone in timezones: %v", timezone)
}

func TestLocale(t *testing.T) {
	r := FromSeed(1234)
	locale := r.Locale()
	_, err := language.Parse(locale)
	assert.NoError(t, err, "invalid locale: %v", locale)
}

func TestLocalePenetration(t *testing.T) {
	r := FromSeed(1234)
	for i := 0; i < 10000; i += 1 {
		locale := r.Locale()
		_, err := language.Parse(locale)
		assert.NoError(t, err, "invalid locale: %v", locale)
	}
}

func TestUserAgentString(t *testing.T) {
	r := FromSeed(1234)
	ua := r.UserAgentString()
	assert.NotEmpty(t, ua, "empty User Agent String")
	assert.True(t, regexp.MustCompile(`^[a-zA-Z]+\/[0-9]+.[0-9]+\ \(.*\).*$`).MatchString(ua),
		"invalid generated User Agent String: %v", ua)
}

func TestPhoneNumbers(t *testing.T) {
	r := FromSeed(1234)
	CheckPhoneNumber(r.PhoneNumber(), t)
}

func CheckPhoneNumber(str string, t *testing.T, msgAndArgs ...interface{}) {
	assert.LessOrEqual(t, len(str)-strings.Count(str, " "), 16, msgAndArgs)

	matched, err := regexp.MatchString("\\+\\d{1,3}\\s\\d{1,3}", str)
	assert.NoError(t, err, msgAndArgs)
	assert.True(t, matched, msgAndArgs)
}

func TestProvinceForCountry(t *testing.T) {
	r := FromSeed(1234)
	supportedCountries := []string{"US", "GB"}
	for _, c := range supportedCountries {
		p := r.ProvinceForCountry(c)
		assert.NotEmpty(t, p, "did not return a valid province for country %s", c)
		switch c {
		case "US":
			assert.Contains(t, jsonData.States, p, "did not return a known province for US")
		case "GB":
			assert.Contains(t, jsonData.ProvincesGB, p, "did not return a known province for GB")
		}
	}
	assert.Empty(t, r.ProvinceForCountry("bogus"), "did not return empty province for unknown country")
}

func TestStreetForCountry(t *testing.T) {
	r := FromSeed(1234)
	supportedCountries := []string{"US", "GB"}
	for _, c := range supportedCountries {
		p := r.StreetForCountry(c)
		assert.NotEmpty(t, p, "did not return a valid street for country %s", c)
	}
	assert.Empty(t, r.StreetForCountry("bogus"), "did not return empty street for unknown country")
}
