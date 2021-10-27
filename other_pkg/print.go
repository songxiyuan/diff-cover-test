package other_pkg

import "fmt"

func Print(n int) {
	switch n {
	case 1:
		fmt.Println(1)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	}
}
