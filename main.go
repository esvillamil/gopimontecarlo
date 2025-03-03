// Ejercicio de pi montecarlo
package main

import (
"fmt"

"github.com/esvillamil/gopimontecarlo/pimontecarlo"
)

// main is the entry point of the program.
//
// It prints the message "Ejercicio de pi montecarlo" to the console.
// Then it calls the GetPi function from the pimontecarlo package
// with different iteration counts and prints the result to the console.
//
// There are no parameters.
// It does not return any values.
func main() {
fmt.Printf("Ejercicio de pi montecarlo\n")
fmt.Printf("Con 1 iteración: ")
fmt.Printf("%v\n",pimontecarlo.GetPi(1))
fmt.Printf("Con 10 iteraciones: ")
fmt.Printf("%v\n",pimontecarlo.GetPi(10))
fmt.Printf("Con 100 iteraciones: ")
fmt.Printf("%v\n",pimontecarlo.GetPi(100))
fmt.Printf("Con 1000 iteraciones: ")
fmt.Printf("%v\n",pimontecarlo.GetPi(1000))
fmt.Printf("Con 10000 iteraciones: ")
fmt.Printf("%v\n",pimontecarlo.GetPi(10000))
}
