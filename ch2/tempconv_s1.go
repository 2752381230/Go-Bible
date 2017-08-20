package main
//package tempconv
import (
    "fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5 / 9)
}
// 命名类型可以自定义其值的行为
func (c Celsius) String() string {
    return fmt.Sprintf("%g C", c)
}

func main() {
    var tempC Celsius = 100.0
    fmt.Println(CToF(tempC))
    var tempF Fahrenheit = 100.0
    fmt.Println(FToC(tempF))
    fmt.Printf("%g\n", BoilingC - FreezingC)
    boilingF := CToF(BoilingC)
    fmt.Printf("%g\n", boilingF - CToF(FreezingC))
    //fmt.Printf("%g\n", boilingF - FreezingC)  // compile error: type mismatch
    var c Celsius
    var f Fahrenheit
    fmt.Println(c == 0)          // "true"
    fmt.Println(f >= 0)          // "true"
    //fmt.Println(c == f)        // compile error: type mismatch
    fmt.Println(c == Celsius(f)) // "true"!”

    c = FToC(212.0)
    fmt.Println(c.String()) // "100°C"
    fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
    fmt.Printf("%s\n", c)   // "100°C"
    fmt.Println(c)          // "100°C"
    fmt.Printf("%g\n", c)   // "100"; does not call String
    fmt.Println(float64(c)) // "100"; does not call String”
}
