package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Color represents a color value in various formats
type Color struct {
	Hex string // hex color value, e.g. "#ff0000"
	R   int    // red component (0-255)
	G   int    // green component (0-255)
	B   int    // blue component (0-255)
	H   int    // hue (0-360)
	S   int    // saturation (0-100)
	L   int    // lightness (0-100)
	V   int    // value (0-100)
	W   int    // whiteness (0-100)
	BL  int    // blackness (0-100)
	C   int    // cyan component (0-100)
	M   int    // magenta component (0-100)
	Y   int    // yellow component (0-100)
	K   int    // black component (0-100)
}

// ConvertColor converts a color value from one type to another
func ConvertColor(value string, inputType string, outputType string) (string, error) {
	// Parse input value and convert to internal Color format
	var c Color
	var err error
	switch inputType {
	case "hex":
		c, err = parseHex(value)
	case "rgb":
		c, err = parseRGB(value)
	case "hsl":
		c, err = parseHSL(value)
	case "hsv":
		c, err = parseHSV(value)
	case "hwb":
		c, err = parseHWB(value)
	case "cmy":
		c, err = parseCMY(value)
	case "cmyk":
		c, err = parseCMYK(value)
	default:
		return "", fmt.Errorf("unsupported input color type: %s", inputType)
	}
	if err != nil {
		return "", err
	}

	// Convert internal Color format to output format
	var outputValue string
	switch outputType {
	case "hex":
		outputValue = c.Hex
	case "rgb":
		outputValue = fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
	case "hsl":
		outputValue = fmt.Sprintf("hsl(%d, %d%%, %d%%)", c.H, c.S, c.L)
	case "hsv":
		outputValue = fmt.Sprintf("hsv(%d, %d%%, %d%%)", c.H, c.S, c.V)
	case "hwb":
		outputValue = fmt.Sprintf("hwb(%d, %d%%, %d%%)", c.H, c.W, c.BL)
	case "cmy":
		outputValue = fmt.Sprintf("cmy(%d%%, %d%%, %d%%)", c.C, c.M, c.Y)
	case "cmyk":
		outputValue = fmt.Sprintf("cmyk(%d%%, %d%%, %d%%, %d%%)", c.C, c.M, c.Y, c.K)
	default:
		return "", fmt.Errorf("unsupported output color type: %s", outputType)
	}

	return outputValue, nil
}

// Parse a hex color value and return it as a Color struct
func parseHex(value string) (Color, error) {
	var c Color

	// Remove leading "#" character if present
	if strings.HasPrefix(value, "#") {
		value = value[1:]
	}

	// Parse red, green, and blue components
	var err error
	if len(value) == 3 { // short hex format, e.g. "#f00"
		c.R, err = strconv.Atoi(string(value[0]) + string(value[0]))
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
		c.G, err = strconv.Atoi(string(value[1]) + string(value[1]))
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
		c.B, err = strconv.Atoi(string(value[2]) + string(value[2]))
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
	} else if len(value) == 6 { // long hex format, e.g. "#ff0000"
		c.R, err = strconv.ParseInt(value[0:2], 16, 0)
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
		c.G, err = strconv.ParseInt(value[2:4], 16, 0)
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
		c.B, err = strconv.ParseInt(value[4:6], 16, 0)
		if err != nil {
			return c, fmt.Errorf("invalid hex color value: %s", value)
		}
	} else {
		return c, fmt.Errorf("invalid hex color value: %s", value)
	}

	// Set hex value
	c.Hex = "#" + value

	return c, nil
}

// Parse an RGB color value and return it as a Color struct
func parseRGB(value string) (Color, error) {
	var c Color

	// Remove whitespace and parentheses if present
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")

	// Parse red, green, and blue components
	var err error
	components := strings.Split(value, ",")
	if len(components) != 3 {
		return c, fmt.Errorf("invalid RGB color value: %s", value)
	}
	c.R, err = strconv.Atoi(components[0])
	if err != nil {
		return c, fmt.Errorf("invalid RGB color value: %s", value)
	}
	c.G, err = strconv.Atoi(components[1])
	if err != nil {
		return c, fmt.Errorf("invalid RGB color value: %s", value)
	}
	c.B, err = strconv.Atoi(components[2])
	if err != nil {
		return c, fmt.Errorf("invalid RGB color value: %s", value)
	}

	// Set hex value
	c.Hex = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)

	return c, nil
}

// Parse an HSL color value and return it as a Color struct
func parseHSL(value string) (Color, error) {
	var c Color

	// Remove whitespace and parentheses if present
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")

	// Parse hue, saturation, and lightness components
	var err error
	components := strings.Split(value, ",")
	if len(components) != 3 {
		return c, fmt.Errorf("invalid HSL color value: %s", value)
	}
	h, err := strconv.Atoi(components[0])
	if err != nil {
		return c, fmt.Errorf("invalid HSL color value: %s", value)
	}
	s, err := strconv.Atoi(strings.TrimSuffix(components[1], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HSL color value: %s", value)
	}
	l, err := strconv.Atoi(strings.TrimSuffix(components[2], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HSL color value: %s", value)
	}

	// Convert HSL to RGB
	r, g, b := hslToRgb(h, s, l)

	// Set RGB values
	c.R = r
	c.G = g
	c.B = b

	// Set hex value
	c.Hex = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)

	return c, nil
}

// Parse an HSV color value and return it as a Color struct
func parseHSV(value string) (Color, error) {
	var c Color

	// Remove whitespace and parentheses if present
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")

	// Parse hue, saturation, and value components
	var err error
	components := strings.Split(value, ",")
	if len(components) != 3 {
		return c, fmt.Errorf("invalid HSV color value: %s", value)
	}
	h, err := strconv.Atoi(components[0])
	if err != nil {
		return c, fmt.Errorf("invalid HSV color value: %s", value)
	}
	s, err := strconv.Atoi(strings.TrimSuffix(components[1], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HSV color value: %s", value)
	}
	v, err := strconv.Atoi(strings.TrimSuffix(components[2], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HSV color value: %s", value)
	}

	// Convert HSV to RGB
	r, g, b := hsvToRgb(h, s, v)

	// Set RGB values
	c.R = r
	c.G = g
	c.B = b

	// Set hex value
	c.Hex = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)

	return c, nil
}

// Parse an HWB color value and return it as a Color struct
func parseHWB(value string) (Color, error) {
	var c Color

	// Remove whitespace and parentheses if present
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")

	// Parse hue, whiteness, and blackness components
	var err error
	components := strings.Split(value, ",")
	if len(components) != 3 {
		return c, fmt.Errorf("invalid HWB color value: %s", value)
	}
	h, err := strconv.Atoi(components[0])
	if err != nil {
		return c, fmt.Errorf("invalid HWB color value: %s", value)
	}
	w, err := strconv.Atoi(strings.TrimSuffix(components[1], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HWB color value: %s", value)
	}
	b, err := strconv.Atoi(strings.TrimSuffix(components[2], "%"))
	if err != nil {
		return c, fmt.Errorf("invalid HWB color value: %s", value)
	}

	// Convert HWB to RGB
	r, g, b := hwbToRgb(h, w, b)

	// Set RGB values
	c.R = r
	c.G = g
	c.B = b

	// Set hex value
	c.Hex = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)

	return c, nil
}

// Parse a CMY color value and return it as a Color struct
func parseCMY(value string) (Color, error) {
	var c Color

	// Remove whitespace and parentheses if present
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")

	// Parse cyan, magenta, and yellow components
	var err error
	components := strings.Split(value, ",")
	if len(components) != 3 {
		return c, fmt.Errorf("invalid CMY color value: %s", value)
	}
	cCyan, err := strconv.ParseFloat(strings.TrimSuffix(components[0], "%"), 64)
	if err != nil {
		return c, fmt.Errorf("invalid CMY color value: %s", value)
	}
	cMagenta, err := strconv.ParseFloat(strings.TrimSuffix(components[1], "%"), 64)
	if err != nil {
		return c, fmt.Errorf("invalid CMY color value: %s", value)
	}
	cYellow, err := strconv.ParseFloat(strings.TrimSuffix(components[2], "%"), 64)
	if err != nil {
		return c, fmt.Errorf("invalid CMY color value: %s", value)
	}

	// Convert CMY to RGB
	r, g, b := cmyToRgb(cCyan, cMagenta, cYellow)

	// Set RGB values
	c.R = r
	c.G = g
	c.B = b

	// Set hex value
	c.Hex = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)

	return c, nil
}

func parseCMYK(value string) (Color, error) {
	// Remove any whitespace and convert to lowercase
	value = strings.ToLower(strings.ReplaceAll(value, " ", ""))

	// Check if the value starts with "cmyk("
	if !strings.HasPrefix(value, "cmyk(") {
		return Color{}, errors.New("invalid CMYK color format")
	}

	// Check if the value ends with ")"
	if !strings.HasSuffix(value, ")") {
		return Color{}, errors.New("invalid CMYK color format")
	}

	// Extract the values between the parentheses
	values := strings.TrimPrefix(value, "cmyk(")
	values = strings.TrimSuffix(values, ")")

	// Split the values by commas
	components := strings.Split(values, ",")

	// Make sure there are exactly four components
	if len(components) != 4 {
		return Color{}, errors.New("invalid CMYK color format")
	}

	// Convert the components to float64 values
	var cmyk [4]float64
	for i, component := range components {
		val, err := strconv.ParseFloat(component, 64)
		if err != nil {
			return Color{}, errors.New("invalid CMYK color format")
		}
		cmyk[i] = val
	}

	// Convert the CMYK values to RGB values
	r := math.Round(255 * (1 - cmyk[0]) * (1 - cmyk[3]))
	g := math.Round(255 * (1 - cmyk[1]) * (1 - cmyk[3]))
	b := math.Round(255 * (1 - cmyk[2]) * (1 - cmyk[3]))

	// Return the RGB values as a Color struct
	return Color{Type: "rgb", Value: fmt.Sprintf("rgb(%d, %d, %d)", int(r), int(g), int(b))}, nil
}

func hslToRgb(h, s, l float64) (r, g, b float64) {
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	r += m
	g += m
	b += m

	return
}

func hsvToRgb(h, s, v float64) (r, g, b float64) {
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := v - c

	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	r += m
	g += m
	b += m

	return
}

func hwbToRgb(h, w, b float64) (r, g, bOut float64) {
	h /= 360

	var c, m, y float64
	if h < 1.0/6 {
		c = 1
		m = 1 - h*6
		y = 1 - m
	} else if h < 2.0/6 {
		h = 2.0/6 - h
		c = h*6 + 1
		m = 1
		y = 1 - c
	} else if h < 3.0/6 {
		h = h - 2.0/6
		c = 1 - h*6
		m = 1
		y = 1 - c
	} else if h < 4.0/6 {
		h = 4.0/6 - h
		c = 0
		m = h*6 + 1
		y = 1 - m
	} else if h < 5.0/6 {
		h = h - 4.0/6
		c = 0
		m = 1 - h*6
		y = m
	} else {
		h = 6.0/6 - h
		c = h*6 + 1
		m = 0
		y = 1 - c
	}

	r = (c*w + m) * (1 - b)
	g = (y*w + m) * (1 - b)
	bOut = (b * 255)

	return
}

func cmyToRgb(c, m, y float64) (r, g, b float64) {
	r = (1 - c) * 255
	g = (1 - m) * 255
	b = (1 - y) * 255

	return
}
