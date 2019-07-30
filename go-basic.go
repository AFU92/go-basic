package main

import "fmt"

func main() {

	// Array movies

	var movies [3]string

	movies[0] = "The notebook"
	movies[1] = "ant-man"
	movies[2] = "avengers"

	fmt.Println("\nMovies list:")
	fmt.Println(movies)

	fmt.Println("\nMovie #2:\n" + movies[1])

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
