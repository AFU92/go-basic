package main

import "fmt"

func main() {

	var colors [3]string

	colors[0] = "Blue"
	colors[1] = "Yellow"
	colors[2] = "Red"

	fmt.Println("\nColors list:")
	fmt.Println(colors)

	fmt.Println("\nColor #2:\n" + colors[1])

	// Array de arrays movies

	var movies [3][2]string

	movies[0][0] = "The notebook"
	movies[0][1] = "Titanic"

	movies[1][0] = "ant-man"
	movies[1][1] = "avengers"

	movies[2][0] = "toy story"
	movies[2][1] = "frozen"

	fmt.Println("\nMovies list:")
	fmt.Println(movies)

	fmt.Println("\nMovies array #2:")
	fmt.Println(movies[1])

	fmt.Println("\nMovie #2.2:")
	fmt.Println(movies[1][1])

	// Array series
	var series = [4]string{
		"defenders",
		"daredevil",
		"iron fist",
		"titans"}

	fmt.Println("\nSeries list:")
	fmt.Println(series)

	// Array primes
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("\nprimes number list:")
	fmt.Println(primes)
}
