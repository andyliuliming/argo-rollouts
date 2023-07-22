package validation

import (
	"os"
	"strconv"
)

func init() {
	if defaultMaxWeight := os.Getenv("TRAFFIC_MANAGER_DEFAULT_MAX_WEIGHT"); defaultMaxWeight != "" {
		maxWeight, err := strconv.ParseInt(defaultMaxWeight, 10, 32)
		if err != nil {
			panic(err)
		}
		DefaultMaxWeight = int32(maxWeight)
	}
}

var (
	DefaultMaxWeight = int32(100)
)
