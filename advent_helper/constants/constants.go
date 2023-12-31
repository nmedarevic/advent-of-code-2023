package constants

var NumberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var NumberRegexPattern = "\\d|one|two|three|four|five|six|seven|eight|nine|ten"

var SpecialChars = map[string]int{
	"#": 35,
	"=": 61,
	"/": 47,
	"&": 38,
	"%": 37,
	"@": 64,
	"$": 36,
	"-": 45,
	"*": 42,
}
