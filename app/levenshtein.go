package app

import (
	levenshtein "github.com/ka-weihe/fast-levenshtein"
)

func GetClosestDistrict(targetDistrict string, districtMap map[string]string) string {
	closestDistrict, closestDistance := "", 100
	for district := range districtMap {
		distance := levenshtein.Distance(targetDistrict, district)
		if distance < closestDistance {
			closestDistance, closestDistrict = distance, district
		}
	}
	return closestDistrict
}
