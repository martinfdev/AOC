package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Reindeer struct {
	Name              string
	SpeedKmPerSec     int
	FlightDurationSec int
	RestDurationSec   int
}

func Day14() {
	reindeers := Read_file("files/day14.txt")

	parsedReindeers, err := parseReindeerDataFromString(reindeers)
	if err != nil {
		fmt.Println("Error to parse reindeer data:", err)
		return
	}

	const raceDurationSec = 2503 // duration of the race in seconds
	//part 1
	winner, bestDistance := findWinningReindeerByDistance(parsedReindeers, raceDurationSec)
	fmt.Printf("The winning reindeer is %s with a distance of %d km\n", winner, bestDistance)

	//part 2
	winners, maxPoints := scoreRaceByPoints(parsedReindeers, raceDurationSec)
	fmt.Printf("The winning reindeer(s) by points: %v with %d points\n", winners, maxPoints)

}

func parseReindeerDataFromString(content string) ([]Reindeer, error) {
	var reindeers []Reindeer
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var name string
		var speed, flightDuration, restDuration int

		_, err := fmt.Sscanf(
			line,
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &speed, &flightDuration, &restDuration,
		)
		if err != nil {
			return nil, fmt.Errorf("entrada inválida: %q (%w)", line, err)
		}

		reindeers = append(reindeers, Reindeer{
			Name:              name,
			SpeedKmPerSec:     speed,
			FlightDurationSec: flightDuration,
			RestDurationSec:   restDuration,
		})
	}
	return reindeers, scanner.Err()
}

func calculateDistance(reindeer Reindeer, totalTimeSec int) int {
	cycleDuration := reindeer.FlightDurationSec + reindeer.RestDurationSec
	fullCycles := totalTimeSec / cycleDuration
	remainingTime := totalTimeSec % cycleDuration
	flyingTime := fullCycles*reindeer.FlightDurationSec + min(remainingTime, reindeer.FlightDurationSec)
	return flyingTime * reindeer.SpeedKmPerSec
}

func findWinningReindeerByDistance(reindeers []Reindeer, totalTimeSec int) (winner string, bestDistance int) {
	for _, reindeer := range reindeers {
		d := calculateDistance(reindeer, totalTimeSec)
		if d > bestDistance {
			bestDistance = d
			winner = reindeer.Name
		}
	}
	return
}

// simulate seconds per second to award points
func scoreRaceByPoints(reindeers []Reindeer, totalTimeSec int) (winners []string, maxPoints int) {
	n := len(reindeers)
	points := make([]int, n)

	for t := 1; t <= totalTimeSec; t++ {
		maxDistance := -1
		distances := make([]int, n)
		for i, reindeer := range reindeers {
			d := calculateDistance(reindeer, t)
			distances[i] = d
			if d > maxDistance {
				maxDistance = d
			}
		}
		// Award points to reindeers at max distance
		for i := range reindeers {
			if distances[i] == maxDistance {
				points[i]++
			}
		}
	}

	// Find max points and winners
	for i, p := range points {
		if p > maxPoints {
			maxPoints = p
			winners = []string{reindeers[i].Name}
		} else if p == maxPoints {
			winners = append(winners, reindeers[i].Name)
		}
	}
	return winners, maxPoints
}
