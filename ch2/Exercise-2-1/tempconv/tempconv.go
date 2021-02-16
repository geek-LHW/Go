// Package tempconv performs Celsius 、Fahrenheit and Kelvin conversions.
// Exercise 2.1: Add typ es, cons tants, and functions to tempconv for pro cessing temperatures in
// the Kelvin scale, where zero Kelvin is −273.15°C andadif ference of 1K has the same mag nitude as 1°C.
package tempconv

import "fmt"

//Celsius ...
type Celsius float64

//Fahrenheit ...
type Fahrenheit float64

//Kelvin ...
type Kelvin float64

//The temperature standard
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(K Kelvin) Celsius { return Celsius(K - 273.15) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
