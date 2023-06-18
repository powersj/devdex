package main

import (
	"fmt"
	"math"
)

func ConvertDataUnits(value float64, fromUnit string, toUnit string) (float64, error) {
	// Define conversion factors for decimal and binary prefixes
	decimalFactors := map[string]float64{
		"B":  1.0,
		"kB": 1e3,
		"MB": 1e6,
		"GB": 1e9,
		"TB": 1e12,
		"PB": 1e15,
		"EB": 1e18,
		"ZB": 1e21,
		"YB": 1e24,
	}
	binaryFactors := map[string]float64{
		"B":   1.0,
		"KiB": math.Pow(2, 10),
		"MiB": math.Pow(2, 20),
		"GiB": math.Pow(2, 30),
		"TiB": math.Pow(2, 40),
		"PiB": math.Pow(2, 50),
		"EiB": math.Pow(2, 60),
		"ZiB": math.Pow(2, 70),
		"YiB": math.Pow(2, 80),
	}

	// Determine the appropriate conversion factor based on the input units
	var factor float64
	if _, ok := decimalFactors[fromUnit]; ok {
		if _, ok := decimalFactors[toUnit]; !ok {
			return 0, fmt.Errorf("invalid 'to' unit: %s", toUnit)
		}
		factor = decimalFactors[fromUnit] / decimalFactors[toUnit]
	} else if _, ok := binaryFactors[fromUnit]; ok {
		if _, ok := binaryFactors[toUnit]; !ok {
			return 0, fmt.Errorf("invalid 'to' unit: %s", toUnit)
		}
		factor = binaryFactors[fromUnit] / binaryFactors[toUnit]
	} else {
		return 0, fmt.Errorf("invalid 'from' unit: %s", fromUnit)
	}

	// Convert the value to the desired units
	result := value * factor
	return result, nil
}
