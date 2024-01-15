package randomdata

import "fmt"

func Example() {
	r := FromSeed(1234)

	// Print a male title
	fmt.Println(r.FirstName(Male))

	// Print a female title
	fmt.Println(r.FirstName(Female))

	// Print a title with random gender
	fmt.Println(r.FirstName(RandomGender))

	// Print a male first name
	fmt.Println(r.FirstName(Male))

	// Print a female first name
	fmt.Println(r.FirstName(Female))

	// Print a first name with random gender
	fmt.Println(r.FirstName(RandomGender))

	// Print a last name
	fmt.Println(r.LastName())

	// Print a male name
	fmt.Println(r.FullName(Male))

	// Print a female name
	fmt.Println(r.FullName(Female))

	// Print a name with random gender
	fmt.Println(r.FullName(RandomGender))

	// Print a random email
	fmt.Println(r.Email())

	// Print a country with full text representation
	fmt.Println(r.Country(FullCountry))

	// Print a country using ISO 3166-1 alpha-3
	fmt.Println(r.Country(ThreeCharCountry))

	// Print a country using ISO 3166-1 alpha-2
	fmt.Println(r.Country(TwoCharCountry))

	// Print a currency using ISO 4217
	fmt.Println(r.Currency())

	// Print the name of a random city
	fmt.Println(r.City())

	// Print the name of a random american state
	fmt.Println(r.State(Large))

	// Print the name of a random american state using two letters
	fmt.Println(r.State(Small))

	// Print a random number >= 10 and <= 20
	fmt.Println(r.Number(10, 20))

	// Print a number >= 0 and <= 20
	fmt.Println(r.Number(20))

	// Print a random float >= 0 and <= 20 with decimal point 3
	fmt.Println(r.Decimal(0, 20, 3))

	// Print a random float >= 10 and <= 20
	fmt.Println(r.Decimal(10, 20))

	// Print a random float >= 0 and <= 20
	fmt.Println(r.Decimal(20))

	// Print a paragraph
	fmt.Println(r.Paragraph())

	// Print a random bool
	fmt.Println(r.Boolean())

	// Print a random postalcode from Sweden
	fmt.Println(r.PostalCode("SE"))

	// Print a random american sounding street name
	fmt.Println(r.Street())

	// Print a random american address
	fmt.Println(r.Address())

	// Print a random string of numbers
	fmt.Println(r.StringNumber(2, "-"))

	// Print a set of 2 random 3-Digits numbers as a string
	fmt.Println(r.StringNumberExt(2, "-", 3))

	// Print a random IPv4 address
	fmt.Println(r.IpV4Address())

	// Print a random IPv6 address
	fmt.Println(r.IpV6Address())

	// Print a random day
	fmt.Println(r.Day())

	// Print a month
	fmt.Println(r.Month())

	// Print full date
	fmt.Println(r.FullDate())
}
