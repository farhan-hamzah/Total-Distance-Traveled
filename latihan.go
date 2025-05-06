package main
import "fmt"

func main(){
	var mainTank, additionalTank int
	fmt.Scan(&mainTank, &additionalTank)
	hasil := distanceTraveled(mainTank, additionalTank)
	fmt.Print(hasil)
}
func distanceTraveled(mainTank int, additionalTank int) int {
    distance := 0

    for mainTank >= 5 {
        distance += 5 * 10
        mainTank -= 5
        if additionalTank > 0 {
            mainTank += 1
            additionalTank--
        }
    }
    distance += mainTank * 10

    return distance
}
