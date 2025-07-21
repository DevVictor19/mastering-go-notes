package main

func main() {

}

func bar(slice []int) int {
	a := ([3]int)(slice)
	return a[0] + a[1] + a[2] + a[3]
}
