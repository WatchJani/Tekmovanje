// At the UPM competition, contestants have different preferences regarding sandwiches. Some of them want sandwiches with meat, whilst the others want vegetarian sandwiches. Some contestant appreciate sandwiches very much, whilst the others less so. Organisers prepared
//  sandwiches with meat and
//  vegetarian sandwiches. In order to satisfy the contestants as much as possible, the organisers conducted a survey with the following two questions:

// How satisfied would you be with a sandwich with meat?
// How satisfied would you be with a vegetarian sandwich?
// Write a program that will distribute sandwiches among the contestants so that the sum of their satisfactions is as high as possible.

// Input data
// The first line contains the number of contestants
// , the number of meaty sandwiches
//  and the number of vegetarian sandwiches
// . It is followed by
//  lines with pairs of integers
//  and
// , which represent the satisfaction of the
// -th contestant with a meaty and a vegetarian sandwich, respectively.

// Input limits
// Output data
// Output the maximum possible total satisfaction (that can be achieved by optimal distribution of sandwiches).

// Examples
// Input
// 3 1 2
// 5 1
// 3 2
// 0 3
// Output
// 10
// Input
// 10 2 8
// 543 635
// 487 819
// 404 52
// 747 874
// 800 178
// 694 616
// 99 452
// 783 253
// 435 363
// 963 748
// Output
// 6142

package main

import (
	"fmt"
	"sort"
	"time"
)

type User struct {
	Meat int
	Veg  int
}

func MaximumPossibleTotal(numberOfContestants, sandwichWithMeat, vegetarianSandwich int, users []User) int {
	if numberOfContestants != sandwichWithMeat+vegetarianSandwich {
		return -1
	}

	var satisfaction int

	sort.Slice(users, func(i, j int) bool {
		diffI := users[i].Meat - users[i].Veg
		diffJ := users[j].Meat - users[j].Veg
		return diffI > diffJ
	})

	for _, value := range users {
		fmt.Println(users)
		if value.Meat > value.Veg && sandwichWithMeat > 0 {
			sandwichWithMeat--
			satisfaction += value.Meat
		} else {
			vegetarianSandwich--
			satisfaction += value.Veg
		}
	}

	return satisfaction
}

func main() {

	start := time.Now()

	numberOfContestants := 10
	sandwichWithMeat := 2
	vegetarianSandwich := 8

	users := []User{
		{Meat: 543, Veg: 635}, //1
		{Meat: 487, Veg: 819}, //2
		{Meat: 404, Veg: 52},  //3
		{Meat: 747, Veg: 874}, //4
		{Meat: 800, Veg: 178}, //5
		{Meat: 694, Veg: 616}, //6
		{Meat: 99, Veg: 452},  //7
		{Meat: 783, Veg: 253}, //8
		{Meat: 435, Veg: 363}, //9
		{Meat: 963, Veg: 748}, //10
	}

	fmt.Println(MaximumPossibleTotal(numberOfContestants, sandwichWithMeat, vegetarianSandwich, users))

	fmt.Println(time.Since(start))
}
