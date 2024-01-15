package randomdata

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var portraitDirs = []string{"men", "women"}

// Profile contains all the data related to a complete profil.
type Profile struct {
	Gender string `json:"gender"`
	Name   struct {
		First string `json:"first"`
		Last  string `json:"last"`
		Title string `json:"title"`
	} `json:"name"`
	Location struct {
		Street   string `json:"street"`
		City     string `json:"city"`
		State    string `json:"state"`
		Postcode int    `json:"postcode"`
	} `json:"location"`

	Email string `json:"email"`
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
		Md5      string `json:"md5"`
		Sha1     string `json:"sha1"`
		Sha256   string `json:"sha256"`
	} `json:"login"`

	Dob        string `json:"dob"`
	Registered string `json:"registered"`
	Phone      string `json:"phone"`
	Cell       string `json:"cell"`

	ID struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	} `json:"id"`

	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	Nat string `json:"nat"`
}

// GenerateProfile generates a full profile.
func (r *Rand) GenerateProfile(gender int) *Profile {
	profile := &Profile{}
	if gender == Male {
		profile.Gender = "male"
	} else if gender == Female {
		profile.Gender = "female"
	} else {
		gender = r.Intn(2)
		if gender == Male {
			profile.Gender = "male"
		} else {
			profile.Gender = "female"
		}
	}
	profile.Name.Title = r.Title(gender)
	profile.Name.First = r.FirstName(gender)
	profile.Name.Last = r.LastName()
	profile.ID.Name = "SSN"
	profile.ID.Value = fmt.Sprintf("%d-%d-%d",
		r.Number(101, 999),
		r.Number(01, 99),
		r.Number(100, 9999),
	)

	profile.Email = r.createEmail(profile.Name.First, profile.Name.Last)
	profile.Cell = r.PhoneNumber()
	profile.Phone = r.PhoneNumber()
	profile.Dob = r.FullDate()
	profile.Registered = r.FullDate()
	profile.Nat = "US"

	profile.Location.City = r.City()
	i, _ := strconv.Atoi(r.PostalCode("US"))
	profile.Location.Postcode = i
	profile.Location.State = r.State(2)
	profile.Location.Street = r.StringNumber(1, "") + " " + r.Street()

	profile.Login.Username = r.SillyName()
	pass := r.SillyName()
	salt := r.RandStringRunes(16)
	profile.Login.Password = pass
	profile.Login.Salt = salt
	profile.Login.Md5 = getMD5Hash(pass + salt)
	profile.Login.Sha1 = getSha1(pass + salt)
	profile.Login.Sha256 = getSha256(pass + salt)

	pic := r.Intn(35)
	profile.Picture.Large = fmt.Sprintf("https://randomuser.me/api/portraits/%s/%d.jpg", portraitDirs[gender], pic)
	profile.Picture.Medium = fmt.Sprintf("https://randomuser.me/api/portraits/med/%s/%d.jpg", portraitDirs[gender], pic)
	profile.Picture.Thumbnail = fmt.Sprintf("https://randomuser.me/api/portraits/thumb/%s/%d.jpg", portraitDirs[gender], pic)

	return profile
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getSha1(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func getSha256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
