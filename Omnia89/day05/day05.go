package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day05")

	// test data
	//dataByRow = []string{
	//	"seeds: 79 14 55 13",
	//	"",
	//	"seed-to-soil map:",
	//	"50 98 2",
	//	"52 50 48",
	//	"",
	//	"soil-to-fertilizer map:",
	//	"0 15 37",
	//	"37 52 2",
	//	"39 0 15",
	//	"",
	//	"fertilizer-to-water map:",
	//	"49 53 8",
	//	"0 11 42",
	//	"42 0 7",
	//	"57 7 4",
	//	"",
	//	"water-to-light map:",
	//	"88 18 7",
	//	"18 25 70",
	//	"",
	//	"light-to-temperature map:",
	//	"45 77 23",
	//	"81 45 19",
	//	"68 64 13",
	//	"",
	//	"temperature-to-humidity map:",
	//	"0 69 1",
	//	"1 0 69",
	//	"",
	//	"humidity-to-location map:",
	//	"60 56 37",
	//	"56 93 4",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
	// 5554894

}

func part01(dataByRow []string) int {

	lowestLocation := -1

	seeds, seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap := getMaps(dataByRow)

	for _, seed := range seeds {
		soilRange := seedToSoilMap.findRange(seed)
		soil := seed + soilRange.delta

		fertilizerRange := soilToFertilizerMap.findRange(soil)
		fertilizer := soil + fertilizerRange.delta

		waterRange := fertilizerToWaterMap.findRange(fertilizer)
		water := fertilizer + waterRange.delta

		lightRange := waterToLightMap.findRange(water)
		light := water + lightRange.delta

		temperatureRange := lightToTemperatureMap.findRange(light)
		temperature := light + temperatureRange.delta

		humidityRange := temperatureToHumidityMap.findRange(temperature)
		humidity := temperature + humidityRange.delta

		locationRange := humidityToLocationMap.findRange(humidity)
		location := humidity + locationRange.delta

		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func part02(dataByRow []string) int {

	lowestLocation := -1

	seeds, seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap := getMaps(dataByRow)

	for i := 0; i < len(seeds); i = i + 2 {
		initialSeed := seeds[i]
		seedRange := seeds[i+1] - 1
		fmt.Printf("  - initialSeed: %d | seedRange: %d\n", initialSeed, seedRange)
		for j := 0; j <= seedRange; j++ {
			if j%10000 == 0 {
				fmt.Printf("  - initialSeed: %d | seedRange: %d | lowestLocation: %d | j: %d\n", initialSeed, seedRange, lowestLocation, j)
			}
			seed := initialSeed + j

			soilRange := seedToSoilMap.findRange(seed)
			soil := seed + soilRange.delta

			fertilizerRange := soilToFertilizerMap.findRange(soil)
			fertilizer := soil + fertilizerRange.delta

			waterRange := fertilizerToWaterMap.findRange(fertilizer)
			water := fertilizer + waterRange.delta

			lightRange := waterToLightMap.findRange(water)
			light := water + lightRange.delta

			temperatureRange := lightToTemperatureMap.findRange(light)
			temperature := light + temperatureRange.delta

			humidityRange := temperatureToHumidityMap.findRange(temperature)
			humidity := temperature + humidityRange.delta

			locationRange := humidityToLocationMap.findRange(humidity)
			location := humidity + locationRange.delta

			if lowestLocation == -1 || location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	return lowestLocation
}

type rangeDiff struct {
	initRange int
	endRange  int
	delta     int
}

type rangeDiffMap []rangeDiff

func (s rangeDiffMap) Len() int {
	return len(s)
}

func (s rangeDiffMap) findRange(n int) rangeDiff {
	for _, v := range s {
		if v.initRange <= n && v.endRange >= n {
			return v
		}
	}
	return rangeDiff{}
}

func getMaps(dataByRow []string) (
	seeds []int,
	seedToSoilMap rangeDiffMap,
	soilToFertilizerMap rangeDiffMap,
	fertilizerToWaterMap rangeDiffMap,
	waterToLightMap rangeDiffMap,
	lightToTemperatureMap rangeDiffMap,
	temperatureToHumidityMap rangeDiffMap,
	humidityToLocationMap rangeDiffMap,
) {

	seedToSoilMap = make(rangeDiffMap, 0)
	soilToFertilizerMap = make(rangeDiffMap, 0)
	fertilizerToWaterMap = make(rangeDiffMap, 0)
	waterToLightMap = make(rangeDiffMap, 0)
	lightToTemperatureMap = make(rangeDiffMap, 0)
	temperatureToHumidityMap = make(rangeDiffMap, 0)
	humidityToLocationMap = make(rangeDiffMap, 0)

	seedToSoilParse := false
	soilToFertilizerParse := false
	fertilizerToWaterParse := false
	waterToLightParse := false
	lightToTemperatureParse := false
	temperatureToHumidityParse := false
	humidityToLocationParse := false

	rowToRangeDiff := func(row string) rangeDiff {
		if row == "" {
			return rangeDiff{}
		}
		parts := strings.Split(row, " ")
		init := util.ToInt(parts[1])
		r := util.ToInt(parts[2]) - 1
		dest := util.ToInt(parts[0])

		return rangeDiff{
			initRange: init,
			endRange:  init + r,
			delta:     dest - init,
		}
	}

	for _, v := range dataByRow {
		if strings.HasPrefix(v, "seeds:") {
			parts := strings.Split(v, ":")
			seeds = util.StringToIntSlice(parts[1], " ")
			continue
		}

		if v == "seed-to-soil map:" {
			seedToSoilParse = true
			continue
		}
		if v == "soil-to-fertilizer map:" {
			soilToFertilizerParse = true
			continue
		}
		if v == "fertilizer-to-water map:" {
			fertilizerToWaterParse = true
			continue
		}
		if v == "water-to-light map:" {
			waterToLightParse = true
			continue
		}
		if v == "light-to-temperature map:" {
			lightToTemperatureParse = true
			continue
		}
		if v == "temperature-to-humidity map:" {
			temperatureToHumidityParse = true
			continue
		}
		if v == "humidity-to-location map:" {
			humidityToLocationParse = true
			continue
		}

		if seedToSoilParse {
			if v == "" {
				seedToSoilParse = false
				continue
			}

			seedToSoilMap = append(seedToSoilMap, rowToRangeDiff(v))
		}
		if soilToFertilizerParse {
			if v == "" {
				soilToFertilizerParse = false
				continue
			}

			soilToFertilizerMap = append(soilToFertilizerMap, rowToRangeDiff(v))
		}
		if fertilizerToWaterParse {
			if v == "" {
				fertilizerToWaterParse = false
				continue
			}

			fertilizerToWaterMap = append(fertilizerToWaterMap, rowToRangeDiff(v))
		}
		if waterToLightParse {
			if v == "" {
				waterToLightParse = false
				continue
			}

			waterToLightMap = append(waterToLightMap, rowToRangeDiff(v))
		}
		if lightToTemperatureParse {
			if v == "" {
				lightToTemperatureParse = false
				continue
			}

			lightToTemperatureMap = append(lightToTemperatureMap, rowToRangeDiff(v))
		}
		if temperatureToHumidityParse {
			if v == "" {
				temperatureToHumidityParse = false
				continue
			}

			temperatureToHumidityMap = append(temperatureToHumidityMap, rowToRangeDiff(v))
		}
		if humidityToLocationParse {
			if v == "" {
				humidityToLocationParse = false
				continue
			}

			humidityToLocationMap = append(humidityToLocationMap, rowToRangeDiff(v))
		}
	}

	return
}
