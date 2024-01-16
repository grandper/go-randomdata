# go-randomdata

randomdata is a tiny help suite for generating random data such as

* first names (male or female)
* last names
* full names (male or female)
* country names (full name or iso 3166.1 alpha-2 or alpha-3)
* locales / language tags (bcp-47)
* random email address
* city names
* American state names (two chars or full)
* random numbers (in an interval)
* random paragraphs
* random bool values
* postal- or zip-codes formatted for a range of different countries.
* american sounding addresses / street names
* silly names - suitable for names of things
* random days
* random months
* random full date
* random full profile
* random date inside range
* random phone number

## Credit where credit is due

This repository is a fork of the [go-randomdata](https://github.com/Pallinder/go-randomdata/graphs/contributors) created by [David Pallinder](https://github.com/Pallinder).

This fork was created to change fundamentally how the random source is used. In the original work, the random source is defined globally. In this version, you need to create a random generator explicitly.

## Installation

```go get github.com/grandper/go-randomdata```

## Usage

```go
package main

import (
    "fmt"
    "github.com/Pallinder/go-randomdata"
)

func main() {
    r := randomdata.FromSeed(1234)

	// Print an int from the half-open interval [0, 10)
	fmt.Println(r.Intn(10))

	// Print a float64 from the half-open interval [0.0, 1.0)
	fmt.Println(r.Float64())

	// Print a random bool
	fmt.Println(r.Boolean())

	// Print a duration
	fmt.Println(r.Duration(24 * time.Hour).Hours())

	// Print a time between now and Tomorrow
	fmt.Println(r.Time(time.Now(), 24*time.Hour))

	// Print a time range now and Tomorrow
	fmt.Println(r.TimeRange(time.Now(), 24*time.Hour))

    // Print a random silly name
    fmt.Println(r.SillyName())

    // Print a male title
    fmt.Println(r.Title(randomdata.Male))

    // Print a female title
    fmt.Println(r.Title(randomdata.Female))

    // Print a title with random gender
    fmt.Println(r.Title(randomdata.RandomGender))

    // Print a male first name
    fmt.Println(r.FirstName(randomdata.Male))

    // Print a female first name
    fmt.Println(r.FirstName(randomdata.Female))

    // Print a last name
    fmt.Println(r.LastName())

    // Print a male name
    fmt.Println(r.FullName(randomdata.Male))

    // Print a female name
    fmt.Println(r.FullName(randomdata.Female))

    // Print a name with random gender
    fmt.Println(r.FullName(randomdata.RandomGender))

    // Print an email
    fmt.Println(r.Email())

    // Print a country with full text representation
    fmt.Println(r.Country(randomdata.FullCountry))

    // Print a country using ISO 3166-1 alpha-2
    fmt.Println(r.Country(randomdata.TwoCharCountry))

    // Print a country using ISO 3166-1 alpha-3
    fmt.Println(r.Country(randomdata.ThreeCharCountry))
    
    // Print BCP 47 language tag
    fmt.Println(r.Locale())

    // Print a currency using ISO 4217
    fmt.Println(r.Currency())

    // Print the name of a random city
    fmt.Println(r.City())

    // Print the name of a random american state
    fmt.Println(r.State(randomdata.Large))

    // Print the name of a random american state using two chars
    fmt.Println(r.State(randomdata.Small))

    // Print an american sounding street name
    fmt.Println(r.Street())

    // Print an american sounding address
    fmt.Println(r.Address())

    // Print a random number >= 10 and < 20
    fmt.Println(r.Number(10, 20))

    // Print a number >= 0 and < 20
    fmt.Println(r.Number(20))

    // Print a random float >= 0 and < 20 with decimal point 3
    fmt.Println(r.Decimal(0, 20, 3))

    // Print a random float >= 10 and < 20
    fmt.Println(r.Decimal(10, 20))

    // Print a random float >= 0 and < 20
    fmt.Println(r.Decimal(20))

    // Print a paragraph
    fmt.Println(r.Paragraph())

    // Print a postal code
    fmt.Println(r.PostalCode("SE"))

    // Print a set of 2 random numbers as a string
    fmt.Println(r.StringNumber(2, "-"))

    // Print a set of 2 random 3-Digits numbers as a string
    fmt.Println(r.StringNumberExt(2, "-", 3))

    // Print a random string sampled from a list of strings
    fmt.Println(r.StringFrom("my string 1", "my string 2", "my string 3"))

    // Print a valid random IPv4 address
    fmt.Println(r.IpV4Address())

    // Print a valid random IPv6 address
    fmt.Println(r.IpV6Address())

    // Print a browser's user agent string
    fmt.Println(r.UserAgentString())

    // Print a day
    fmt.Println(r.Day())

    // Print a month
    fmt.Println(r.Month())

    // Print full date like Monday 22 Aug 2016
    fmt.Println(r.FullDate())

    // Print full date <= Monday 22 Aug 2016
    fmt.Println(r.FullDateInRange("2016-08-22"))

    // Print full date >= Monday 01 Aug 2016 and <= Monday 22 Aug 2016
    fmt.Println(r.FullDateInRange("2016-08-01", "2016-08-22"))

    // Print phone number according to e.164
    fmt.Println(r.PhoneNumber())

    // Get a complete and randomised profile of data generally used for users
    // There are many fields in the profile to use check the Profile struct definition in fullprofile.go
    profile := randomdata.GenerateProfile(randomdata.Male | randomdata.Female | randomdata.RandomGender)
    fmt.Printf("The new profile's username is: %s and password (md5): %s\n", profile.Login.Username, profile.Login.Md5)

    // Get a random country-localised street name for Great Britain
    fmt.Println(r.StreetForCountry("GB"))
    // Get a random country-localised street name for USA
    fmt.Println(r.StreetForCountry("US"))

    // Get a random country-localised province for Great Britain
    fmt.Println(r.ProvinceForCountry("GB"))
    // Get a random country-localised province for USA
    fmt.Println(r.ProvinceForCountry("US"))
}
```
