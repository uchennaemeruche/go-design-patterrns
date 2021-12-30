package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return &ak47{
			gun: gun{
				name:  gunType + " gun",
				power: 4,
			},
		}, nil
	} else if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong or invalid gun type passed")
}
