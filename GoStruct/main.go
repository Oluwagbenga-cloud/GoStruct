package main

import (
	"fmt"
	"strings"
	
)

// type Car struct{
// 	Brand string
// 	Model string
// 	Price float64
// }

// func getCarAsText(car Car) string {
// 	carform := fmt.Sprintf("%s %s Price: %0.2f USD",car.Brand, car.Model, car.Price)

// 	return carform

// }

// func main(){
// 	car := Car{
// 		Brand: "Audi",
// 		Model: "A5",
// 		Price: 51999.9999,
// 	}

// 	fullCar := getCarAsText(car)
// 	fmt.Println(fullCar)
// }

type Movie struct {
	Name      string
	Price     float64
	SizeBytes int
}

func main(){
	 availableMovies := []Movie{
		{
			Name: "The Meg",
			Price: 10.50,
			SizeBytes: 20000,
		},
		 {
			Name: "The Green Mile",
			Price: 6.890,
			SizeBytes: 30000,
		},
		{
			Name: "Swat Team",
			Price:  12.50,
			SizeBytes: 25000,
		},
		{
			Name: "Barbie",
			Price: 15.00,
			SizeBytes: 30000,
		},
		{
			Name: "Flash",
			Price: 5.50,
			SizeBytes: 45000,
		},
	}

	movieSold := []Movie{availableMovies[3], availableMovies[4]}

	turnOver := getSumOfMovies(movieSold)
	list := getListOfSoldMovies(movieSold)
	byteSum := getTotalSizeOfBytes(availableMovies)
	allMovies := getAllMovieNames(availableMovies)

	//fmt.Printf("Movies list: %s\n",allMovies)
	fmt.Printf("Movies list: %v\n",strings.Join(allMovies, ","))
	fmt.Printf("Turnover: %.2f GBP, total sold count: %d, used disk size: %d bytes\n", turnOver, list,byteSum)
}



func getSumOfMovies(movieSold []Movie)float64{
	turnOver := 0.0
	for _, currentMovie := range movieSold {
		turnOver = turnOver + currentMovie.Price
	}
	return turnOver

}

func getListOfSoldMovies(availableMovies []Movie)int{
	return len(availableMovies)
}

func getTotalSizeOfBytes(movies []Movie)int{
	byteSum := 0
	for _, currentMovie := range movies{
		byteSum = byteSum + currentMovie.SizeBytes
	}
	return byteSum
}

func getAllMovieNames(availableMovies []Movie)[]string{
	allMovies := []string{}
	for _, currentMovie := range availableMovies {
		allMovies = append(allMovies, currentMovie.Name)

	} 
	return allMovies
}

//const mapItemCount = x


