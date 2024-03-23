package common

import (
	timezone "github.com/evanoberholster/timezoneLookup/v2"
	"os"
)

func GetTimeZone(latitude float64, longitude float64) (string, error) {
	var tzc = timezone.Timezonecache{}
	tzFile, err := os.Open("timezone.data")
	if err != nil {
		return "", err
	}

	defer func(tzFile *os.File) {
		err := tzFile.Close()
		if err != nil {
			// log err
		}
	}(tzFile)

	if err := tzc.Load(tzFile); err != nil {
		return "", err
	}

	defer func(tzc *timezone.Timezonecache) {
		err := tzc.Close()
		if err != nil {
			// log err
		}
	}(&tzc)

	tz, _ := tzc.Search(latitude, longitude)
	return tz.Name, nil
}
