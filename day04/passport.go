package day04

import (
	"regexp"
	"strconv"
	"strings"
)

// A Passport represents a collection of fields that pertain to a traveler's
// personally identifying information.
type Passport map[string]string

// ParsePassportLine takes a puzzle input passport line and transforms it into
// a Passport key/value store.
func ParsePassportLine(line string) Passport {
	p := make(Passport)

	var key string
	var value string

	rawFields := strings.Split(line, " ")
	for _, field := range rawFields {
		key = field[:3]
		value = field[4:]

		// The puzzle indicates that the country identifier field is to be
		// ignored, so we shall not store this.
		if key == "cid" {
			continue
		}

		p[key] = value
	}

	return p
}

// requiredFields is an array of fields that are required on a passport for it
// to be considered valid.
var requiredFields = [...]string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// HasRequiredFields determines whether or not a passport contains all of the
// required fields.
func (p Passport) HasRequiredFields() bool {
	_, hasCid := p["cid"]

	if (!hasCid && len(requiredFields) != len(p)) ||
		(hasCid && len(requiredFields) != len(p)-1) {
		return false
	}

	for _, field := range requiredFields {
		if _, ok := p[field]; !ok {
			return false
		}
	}

	return true
}

// colorRE matches a 6 character hex color code with the leading octothorpe.
var colorRE = regexp.MustCompile("#[0-9a-f]{6}")

// eyeColors is an array of the eye color codes valid on a Passport.
var eyeColors = [...]string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

// IsValid checks to see whether a passport has both all of the required fields
// AND those fields have an acceptable value.
func (p Passport) IsValid() bool {
	if !p.HasRequiredFields() {
		return false
	}

	var err error

	var byr int
	if byr, err = strconv.Atoi(p["byr"]); err != nil {
		return false
	} else if byr > 2002 || byr < 1920 {
		return false
	}

	var iyr int
	if iyr, err = strconv.Atoi(p["iyr"]); err != nil {
		return false
	} else if iyr < 2010 || iyr > 2020 {
		return false
	}

	var eyr int
	if eyr, err = strconv.Atoi(p["eyr"]); err != nil {
		return false
	} else if eyr < 2020 || eyr > 2030 {
		return false
	}

	rawHgt := p["hgt"]
	hgtUnit := rawHgt[len(rawHgt)-2:]
	hgtStr := rawHgt[:len(rawHgt)-2]

	if hgtUnit != "in" && hgtUnit != "cm" {
		return false
	}

	var hgt int
	if hgt, err = strconv.Atoi(hgtStr); err != nil {
		return false
	} else if hgtUnit == "in" && (hgt < 59 || hgt > 76) {
		return false
	} else if hgtUnit == "cm" && (hgt < 150 || hgt > 193) {
		return false
	}

	hcl := p["hcl"]
	if !colorRE.MatchString(hcl) {
		return false
	} else if len(hcl) > 0 && hcl[0] != '#' {
	}

	ecl := p["ecl"]
	validEcl := false
	for _, c := range eyeColors {
		if ecl == c {
			validEcl = true
			break
		}
	}

	if !validEcl {
		return false
	}

	if _, err = strconv.Atoi(p["pid"]); err != nil {
		return false
	} else if len(p["pid"]) != 9 {
		return false
	}

	return true
}
