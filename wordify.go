package wordify

import (
	"math"
	"strings"
)

type unidad uint8

const (
	cero unidad = iota
	uno
	dos
	tres
	cuatro
	cinco
	seis
	siete
	ocho
	nueve
	diez
	once
	doce
	trece
	catorce
	quince
	dieciseis
	diecisiete
	dieciocho
	diecinueve
	veinte
	veintiuno
	veintidos
	veintitres
	veinticuatro
	veinticinco
	veintiseis
	veintisiete
	veintiocho
	veintinueve
)

func (d unidad) String() string {
	return [30]string{
		"cero",
		"uno",
		"dos",
		"tres",
		"cuatro",
		"cinco",
		"seis",
		"siete",
		"ocho",
		"nueve",
		"diez",
		"once",
		"doce",
		"trece",
		"catorce",
		"quince",
		"dieciséis",
		"diecisiete",
		"dieciocho",
		"diecinueve",
		"veinte",
		"veintiuno",
		"veintidós",
		"veintitrés",
		"veinticuatro",
		"veinticinco",
		"veintiséis",
		"veintisiete",
		"veintiocho",
		"veintinueve",
	}[d]
}

type decena uint8

const (
	treinta decena = 10 * (iota + 3)
	cuarenta
	cincuenta
	sesenta
	setenta
	ochenta
	noventa
)

func (d decena) String() string {
	toIdx := map[decena]uint8{
		treinta:   0,
		cuarenta:  1,
		cincuenta: 2,
		sesenta:   3,
		setenta:   4,
		ochenta:   5,
		noventa:   6,
	}

	return [7]string{
		"treinta",
		"cuarenta",
		"cincuenta",
		"sesenta",
		"setenta",
		"ochenta",
		"noventa",
	}[toIdx[d]]
}

type centena uint16

const (
	cien centena = 100 * (iota + 1)
	doscientos
	trescientos
	cuatrocientos
	quinientos
	seiscientos
	setecientos
	ochocientos
	novecientos
)

func (c centena) String() string {
	toIdx := map[centena]uint8{
		cien:          0,
		doscientos:    1,
		trescientos:   2,
		cuatrocientos: 3,
		quinientos:    4,
		seiscientos:   5,
		setecientos:   6,
		ochocientos:   7,
		novecientos:   8,
	}

	return [9]string{
		"ciento",
		"doscientos",
		"trescientos",
		"cuatrocientos",
		"quinientos",
		"seiscientos",
		"setecientos",
		"ochocientos",
		"novecientos",
	}[toIdx[c]]
}

// oom: [o]rders [o]f [m]agnitud
type oom uint8

const (
	cientos oom = iota + 1
	miles
	millones
	milesDeMillones
)

// useful keywords to build the number word
const (
	and       string = " y "
	space     string = " "
	un        string = "un"
	menos     string = "menos"
	hundred   string = "cien"
	twentyone string = "veintiún"
	thousand  string = "mil"
	million   string = "millón"
	millions  string = "millones"
)

// Int returns the Spanish word representation of a given integer number
// ranging from -999_999_999_999 up to 999_999_999_999
func Int(num int) string {
	var (
		number      int
		spanishWord string
	)
	if num < 0 {
		number = -1 * num
		spanishWord += menos + space
	} else {
		number = num
	}
	numLen := numLenght(number)
	// numbers are grouped into powers of ten according to:
	//    10^(3p)
	// with p in {0, 1, 2, ..., n} so that,
	// Group 1: 10^(3*0)     <= numbers < 10^(3*1)
	// Group 2: 10^(3*1)     <= numbers < 10^(3*2)
	// Group N: 10^(3*(N-1)) <= numbers < 10^(3*N)
	oomGroups := make(map[oom]int)
	for i := 1; i <= numLen/3; i++ {
		// algorithm to store numbers according to their corresponding group:
		//    (number % 10^(3*i)) / 10^(3*(i-1))
		oomGroups[oom(i)] = (number % int(math.Pow10(3*i))) / int(math.Pow10(3*(i-1)))
	}
	// if the lenght of the number is not divisible by 3, then we need to manually
	// add the left most numbers that will be stored into the highest ranked group,
	// that is, the group with the highest upper limit. For this we used the algorithm:
	//    number / 10^(number length - remainder)
	remainder := numLen % 3
	if remainder != 0 {
		idx := 1 + numLen/3
		oomGroups[oom(idx)] = number / int(math.Pow10(numLen-remainder))
	}

	cientosRes := ""
	milesRes := ""
	millonesRes := ""
	milesDeMillonesRes := ""
	for group, numbers := range oomGroups {
		switch group {
		case cientos:
			cientosRes = numberToWords(numbers, "")
		case miles:
			milesRes = numberToWords(numbers, thousand)
		case millones:
			millonesRes = numberToWords(numbers, millions)
		case milesDeMillones:
			milesDeMillonesRes = numberToWords(numbers, thousand)
		}
	}

	spanishWord += strings.TrimSpace(milesDeMillonesRes + space + millonesRes + space + milesRes + space + cientosRes)

	return spanishWord
}

func numberToWords(number int, orderOfMag string) string {
	uni := extractorUnidad(number)
	dec := extractorDecena(number)
	cen := extractorCentena(number)
	isUn := ""
	numberIsOne := ""
	numberIsTwentyOne := ""
	switch orderOfMag {
	case millions:
		isUn = un
		numberIsOne = un + space + million
		numberIsTwentyOne = twentyone
	case "":
		isUn = unidad(1).String()
		numberIsOne = unidad(1).String()
		numberIsTwentyOne = unidad(21).String()
	default:
		isUn = un
		numberIsOne = orderOfMag
		numberIsTwentyOne = twentyone
	}
	numberInWords := ""
	if number == 0 {
		numberInWords = ""
	} else if number == 1 {
		numberInWords = numberIsOne
	} else if number < 30 && number%100 == 21 {
		numberInWords = numberIsTwentyOne + space + orderOfMag
	} else if number < 30 {
		numberInWords = unidad(number).String() + space + orderOfMag
	} else if number == 100 {
		numberInWords = hundred + space + orderOfMag
	} else if number < 100 && number >= 30 && uni == 0 {
		numberInWords = decena(dec).String() + space + orderOfMag
	} else if number < 100 && number >= 30 && uni == 1 {
		numberInWords = decena(dec).String() + and + isUn + space + orderOfMag
	} else if number < 100 && number >= 30 {
		numberInWords = decena(dec).String() + and + unidad(uni).String() + space + orderOfMag
	} else if number >= 100 && number%100 == 0 {
		numberInWords = centena(cen).String() + space + orderOfMag
	} else if number >= 100 && number%100 == 1 {
		numberInWords = centena(cen).String() + space + isUn + space + orderOfMag
	} else if number >= 100 && number%100 == 21 {
		numberInWords = centena(cen).String() + space + numberIsTwentyOne + space + orderOfMag
	} else if number >= 100 && number%100 < 30 {
		numberInWords = centena(cen).String() + space + unidad(number%100).String() + space + orderOfMag
	} else if number >= 100 && number%100 >= 30 && uni == 0 {
		numberInWords = centena(cen).String() + space + decena(dec).String() + space + orderOfMag
	} else if number >= 100 && number%100 >= 30 && uni == 1 {
		numberInWords = centena(cen).String() + space + decena(dec).String() + and + isUn + space + orderOfMag
	} else {
		numberInWords = centena(cen).String() + space + decena(dec).String() + and + unidad(uni).String() + space + orderOfMag
	}

	return numberInWords
}

func numLenght(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

// extractorCentena returns the number in the hundreds place.
// Eg: 122 -> 100
func extractorCentena(num int) int {
	return 100 * (num / 100)
}

// extractorDecena returns the number in the tenths place.
// Eg: 122 -> 20
func extractorDecena(num int) int {
	return 10 * ((num % 100) / 10)
}

// extractorUnidad returns the number in the units place.
// Eg: 122 -> 2
func extractorUnidad(num int) int {
	return num % 10
}
