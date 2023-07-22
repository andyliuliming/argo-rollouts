package weightutil

import (
	"os"
	"strconv"

	"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
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

func MaxTrafficWeight(ro *v1alpha1.Rollout) int32 {
	maxWeight := int32(100)
	if ro.Spec.Strategy.Canary != nil && ro.Spec.Strategy.Canary.TrafficRouting != nil && ro.Spec.Strategy.Canary.TrafficRouting.MaxTrafficWeight != nil {
		maxWeight = *ro.Spec.Strategy.Canary.TrafficRouting.MaxTrafficWeight
	}
	return maxWeight
}
