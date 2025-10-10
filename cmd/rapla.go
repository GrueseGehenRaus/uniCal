package main

import (
	"net/http"
	"os"
	"strings"

	ics "github.com/arran4/golang-ical"
)

type Rapla struct {
	cal *ics.Calendar
}

// Creating a new Rapla instance based on a provided URL
func FetchNewRaplaInstance(url string) (*Rapla, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cal, err := ics.ParseCalendar(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Rapla{cal: cal}, nil
}

// Save the filtered calendar to a file
func (rapla *Rapla) saveFilteredICal(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(rapla.cal.Serialize())
	if err != nil {
		return err
	}

	return nil
}

// Functions that operate on the calendar

// Filter events based on provided blocklist
func (rapla *Rapla) filterEvents(blocklist []string) {
	// Create a new calendar and copy relevant properties from the original
	filteredCal := ics.NewCalendar()
	for _, event := range rapla.cal.Events() {
		blocklisted := false
		for _, blockedTitle := range blocklist {
			if evenProperty := event.GetProperty(ics.ComponentPropertySummary); evenProperty != nil && evenProperty.Value == blockedTitle || strings.Contains(evenProperty.Value, blockedTitle) && evenProperty != nil {
				blocklisted = true
				break
			}
		}
		if !blocklisted {
			filteredCal.AddVEvent(event)
		}
	}
	rapla.cal = filteredCal
}
