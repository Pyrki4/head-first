package divide

import (
	"errors"
)

func Divide(divided, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return divided / divisor, nil
}

//func main() {
//	quotient, err := Divide(5.6, 0.0)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%0.2f\n", quotient)
//	}
//}
