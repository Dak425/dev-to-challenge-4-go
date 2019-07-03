package memory

import (
	devto "github.com/Dak425/dev-to-challenge-4-go"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type parser struct {
	regex *regexp.Regexp
}

func newParser() parser {
	regex, err := regexp.Compile("[^a-zA-Z0-9.\\n\\s]+")

	if err != nil {
		log.Fatal("unable to compile regexp when initializing parser")
	}

	return parser{
		regex: regex,
	}
}

func (p parser) getLines(raw string) []string {
	return strings.Split(raw, "\n")
}

func (p parser) sanitizeLine(line string) string {
	return p.regex.ReplaceAllString(line, "")
}

func (p parser) initialBalanceFromLines(lines []string) (float64, []string) {
	initialBalance, err := strconv.ParseFloat(p.sanitizeLine(lines[0]), 64)

	if err != nil {
		log.Fatalf("unable to load initial balance from lines: %v", err)
	}

	return initialBalance, lines[1:]
}

func (p parser) transactionFromLine(line string) devto.Transaction {
	parts := strings.Split(p.sanitizeLine(line), " ")

	if len(parts) != 3 {
		log.Fatalf("invalid line parsed for check book: %v", parts)
	}

	checkNumber, err := strconv.Atoi(parts[0])

	if err != nil {
		log.Fatalf("unable to parse check number from line: %v", err)
	}

	amount, err := strconv.ParseFloat(parts[2], 64)

	if err != nil {
		log.Fatalf("unable to parse amount from line: %v", err)
	}

	return devto.Transaction{
		CheckNumber: checkNumber,
		Category:    parts[1],
		Amount:      amount,
	}
}
