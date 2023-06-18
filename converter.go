package main

import (
	"fmt"
)

func Example() {
	// Example usage
	result := ConvertUnits(10, "cm", "in")
	fmt.Println(result) // Output: 3.937007874015748
}

func ConvertUnits(value float64, fromUnit string, toUnit string) float64 {
	// Define conversion factors for each unit
	conversionFactors := map[string]float64{
		// Length - Metric
		"m":  1.0,
		"km": 1000.0,
		"dm": 0.1,
		"cm": 0.01,
		"mm": 0.001,
		"μm": 1e-6,
		"nm": 1e-9,
		"pm": 1e-12,

		// Length - Imperial
		"in":  0.0254,
		"ft":  0.3048,
		"yd":  0.9144,
		"mi":  1609.34,
		"nmi": 1852.0,

		// Temperature - Metric
		"celsius":     1.0,
		"kelvin":      1.0,
		"millikelvin": 1e-3,

		// Temperature - Imperial
		"fahrenheit": 5.0 / 9.0,
		"rankine":    5.0 / 9.0,

		// Mass - Metric
		"kg": 1.0,
		"g":  0.001,
		"mg": 1e-6,
		"μg": 1e-9,
		"t":  1000.0,

		// Mass - Imperial
		"lb": 0.453592,
		"oz": 0.0283495,
		"gr": 0.06479891,

		// Volume - Metric
		"l":   1.0,
		"ml":  0.001,
		"cl":  0.01,
		"dl":  0.1,
		"m^3": 1000.0,

		// Volume - Imperial
		"tsp":  0.00492892,
		"tbsp": 0.0147868,
		"ozv":  0.0295735,
		"c":    0.236588,
		"pt":   0.473176,
		"qt":   0.946353,
		"gal":  3.78541,
		"ft^3": 28.3168,
		"yd^3": 764.5549,

		// Pressure
		"pa":   1.0,
		"bar":  1e5,
		"atm":  101325.0,
		"mmhg": 133.322,
		"psi":  6894.76,

		// Energy
		"j":    1.0,
		"kj":   1000.0,
		"cal":  4.184,
		"kcal": 4184.0,
		"btu":  1055.06,

		// Power
		"w":  1.0,
		"kw": 1000.0,
		"hp": 746.0,

		// Area
		"m²":  1.0,
		"km²": 1e6,
		"ha":  1e4,
		"ft²": 0.092903,
		"yd²": 0.836127,
		"mi²": 2.58999e6,
		"ac":  4046.86,
		"in²": 0.00064516,
		"cm²": 0.0001,
		"a":   100.0,

		// Time
		"s":   1.0,
		"ms":  0.001,
		"μs":  0.000001,
		"ns":  0.000000001,
		"min": 60.0,
		"h":   3600.0,
		"d":   86400.0,
		"wk":  604800.0,
		"mo":  2628000.0,
		"yr":  31536000.0,

		// Speed
		"m/s":  1.0,
		"km/h": 0.27777777777778,
		"mph":  0.44704,
		"ft/s": 0.3048,
		"kt":   0.51444444444444,
	}

	// Retrieve conversion factor for the original unit
	fromFactor, fromExists := conversionFactors[fromUnit]
	if !fromExists {
		panic(fmt.Sprintf("Invalid from unit: %s", fromUnit))
	}

	// Retrieve conversion factor for the destination unit
	toFactor, toExists := conversionFactors[toUnit]
	if !toExists {
		panic(fmt.Sprintf("Invalid to unit: %s", toUnit))
	}

	// Convert value from the original unit to the base unit
	baseValue := value * fromFactor

	// Convert base value to the destination unit
	result := baseValue / toFactor

	return result
}
