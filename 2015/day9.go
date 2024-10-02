package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day9() {
	data := Read_file("files/day9.txt")

	//regular expression to extract city names and distances
	re := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

	//map to store the distances between cities
	distances := make(map[string]map[string]int)

	//aplit data into lines
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		//extract city names and distances
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			city1 := matches[1]
			city2 := matches[2]
			distance, _ := strconv.Atoi(matches[3])

			if distances[city1] == nil {
				distances[city1] = make(map[string]int)
			}
			if distances[city2] == nil {
				distances[city2] = make(map[string]int)
			}
			distances[city1][city2] = distance
			distances[city2][city1] = distance
		}
	}

	//get list of cities
	cities := make([]string, 0, len(distances))
	for city := range distances {
		cities = append(cities, city)
	}

	//found the route with the shortest distance
	rute, distance := findShortestRoute(cities, distances)
	fmt.Println("Day 9: ", distance, rute)
}

func findShortestRoute(cities []string, distances map[string]map[string]int) ([]string, int) {
	//initialize variables
	shortestDistance := 0
	shortestRoute := make([]string, 0)

	//get all possible routes
	routes := permutations(cities)

	//iterate over all routes
	for _, route := range routes {
		distance := 0
		for i := 0; i < len(route)-1; i++ {
			distance += distances[route[i]][route[i+1]]
		}
		if shortestDistance == 0 || distance < shortestDistance {
			shortestDistance = distance
			shortestRoute = route
		}
	}
	return shortestRoute, shortestDistance
}

func permutations(cities []string) [][]string {
	//initialize variables
	perm := make([][]string, 0)
	used := make(map[string]bool)

	//recursive function to get all possible routes
	var permute func([]string)
	permute = func(arr []string) {
		if len(arr) == len(cities) {
			perm = append(perm, append([]string(nil), arr...))
			return
		}
		for _, city := range cities {
			if !used[city] {
				used[city] = true
				permute(append(arr, city))
				used[city] = false
			}
		}
	}
	permute([]string{})
	return perm
}
