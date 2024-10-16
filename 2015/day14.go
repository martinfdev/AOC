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
	winner, bestDistance := findWinningReindeerByDistance(parsedReindeers, 2503) // 2503 seconds of race
	fmt.Printf("The winning reindeer is %s with a distance of %d km\n", winner, bestDistance)

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
