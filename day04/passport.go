package day04

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

func ParsePassportLine(line string) Passport {
	p := make(Passport)

	var key string
	var value string

	rawFields := strings.Split(line, " ")
	for _, field := range rawFields {
		key = field[:3]
		value = field[4:]

		if key == "cid" {
			continue
		}

		p[key] = value
	}

	return p
}

var requiredFields = [...]string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func (p Passport) IsValid() bool {
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

var colorRE = regexp.MustCompile("#[0-9a-f]{6}")
var eyeColors = [...]string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

func (p Passport) IsValidExtended() bool {
	if !p.IsValid() {
		return false
	}

	var byr, iyr, eyr int
	var hgt string
	var hgtValueStr string
	var hgtValue int
	var hgtUnit string
	var hcl string
	var ecl string
	var err error

	if byr, err = strconv.Atoi(p["byr"]); err != nil {
		return false
	} else if byr > 2002 || byr < 1920 {
		return false
	}

	if iyr, err = strconv.Atoi(p["iyr"]); err != nil {
		return false
	} else if iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr, err = strconv.Atoi(p["eyr"]); err != nil {
		return false
	} else if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt = p["hgt"]
	hgtUnit = hgt[len(hgt)-2:]
	hgtValueStr = hgt[:len(hgt)-2]

	if hgtUnit != "in" && hgtUnit != "cm" {
		return false
	}

	if hgtValue, err = strconv.Atoi(hgtValueStr); err != nil {
		return false
	} else if hgtUnit == "in" && (hgtValue < 59 || hgtValue > 76) {
		return false
	} else if hgtUnit == "cm" && (hgtValue < 150 || hgtValue > 193) {
		return false
	}

	hcl = p["hcl"]
	if !colorRE.MatchString(hcl) {
		return false
	} else if len(hcl) > 0 && hcl[0] != '#' {
	}

	ecl = p["ecl"]
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
