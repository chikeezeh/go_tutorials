package main
import "fmt"
// import "strconv"
func main() {
	printMe("Hello world")
	// printMe(strconv.FormatInt(intDivision(10,2), 10))
	fmt.Println(intDivision(5,2))
}

func printMe(printValue string) {
	// This function prints the value passed to it.
	fmt.Println(printValue)

}
 func intDivision(numerator int64, denominator int64) int64 {
	var result int64 = numerator/denominator
	return result
 }