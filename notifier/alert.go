package notifier

import (
	"fmt"
)

func Notify(symbol string, price, threshold float64) {
	if price > threshold {
		fmt.Printf("Alert! %s price has crossed $%.2f. Current Price: $%.2f\n", symbol, threshold, price)
	}
}
