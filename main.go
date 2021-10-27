package main

import "diff-cover-test/other_pkg"

func main() {
	return
}

func Print(n int) {
	switch n {
	case 1:
		other_pkg.Print(1)
	case 2:
		other_pkg.Print(2)
	case 3:
		other_pkg.Print(3)
	case 4:
		other_pkg.Print(4)
	case 5:
		other_pkg.Print(5)
	}
}
