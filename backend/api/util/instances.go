package util

import (
	"fmt"
	"math/rand"
	"time"
)

type InstanceAlgorithm struct {
	Algorithm string `json:"type"`
}

type SelectedInstance struct {
	Port int `json:"port"`
}

func GetCurrentInstance(algorithm *InstanceAlgorithm, instances []int) *SelectedInstance {
	if algorithm == nil {
		return nil
	}

	var currentInstance SelectedInstance
	switch algorithm.Algorithm {
	case ROUND_ROBIN:
		robin := RoundRobin{
			LastIndex: 0,
		}

		selectedPort := robin.SelectInstance(instances)
		currentInstance = SelectedInstance{
			Port: instances[selectedPort],
		}

	case RANDOM:
		random := Random{}

		selectedPort := random.SelectInstance(instances)
		currentInstance = SelectedInstance{
			Port: instances[selectedPort],
		}
	}

	return &currentInstance
}

// RoundRobin / ROUND_ROBIN Load Balancer
type RoundRobin struct {
	LastIndex int
}

func (rr *RoundRobin) SelectInstance(instances []int) int {
	rr.LastIndex = (rr.LastIndex + 1) % len(instances)
	return instances[rr.LastIndex]
}

// Random Load Balancer
type Random struct{}

func (r *Random) SelectInstance(instances []int) int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(instances))
	fmt.Println("Selected random index:", index)
	return index
}

// LeastConnections LEAST_CONNECTIONS Load Balancer
type LeastConnections struct {
	// In a real scenario, you'd track active connections for each instance
}

//
//func (lc *LeastConnections) SelectInstance(instances []Instance) Instance {
//	// Placeholder for selecting the least connections (return the first instance for now)
//	return instances[0]
//}
//
//// LeastResponseTime LEAST_RESPONSE_TIME Load Balancer
//type LeastResponseTime struct {
//	// In a real scenario, you'd track response times for each instance
//}
//
//func (lrt *LeastResponseTime) SelectInstance(instances []Instance) Instance {
//	// Placeholder for selecting the least response time (return the first instance for now)
//	return instances[0]
//}
//
//// IPHash IP_HASH Load Balancer
//type IPHash struct{}
//
//func (ih *IPHash) SelectInstance(instances []Instance) Instance {
//	// Placeholder for IP hash (return the first instance for now)
//	return instances[0]
//}
//
//// WeightedRoundRobin WEIGHTED_ROUND_ROBIN Load Balancer
//type WeightedRoundRobin struct{}
//
//func (wrr *WeightedRoundRobin) SelectInstance(instances []Instance) Instance {
//	// Placeholder for weighted round-robin (return the first instance for now)
//	return instances[0]
//}
//
//// WeightedLeastConnections WEIGHTED_LEAST_CONNECTIONS Load Balancer
//type WeightedLeastConnections struct{}
//
//func (wlc *WeightedLeastConnections) SelectInstance(instances []Instance) Instance {
//	// Placeholder for weighted least connections (return the first instance for now)
//	return instances[0]
//}
//

//
//// ConsistencyHashing CONSISTENCY_HASHING Load Balancer
//type ConsistencyHashing struct{}
//
//func (ch *ConsistencyHashing) SelectInstance(instances []Instance, clientIP string) Instance {
//	// Placeholder for consistency hashing (return the first instance for now)
//	return instances[0]
//}
//
//// ClientIPAffinity CLIENT_IP_AFFINITY Load Balancer
//type ClientIPAffinity struct{}
//
//func (cipa *ClientIPAffinity) SelectInstance(instances []Instance, clientIP string) Instance {
//	// Placeholder for client IP affinity (return the first instance for now)
//	return instances[0]
//}
//
//// HealthBased HEALTH_BASED Load Balancer
//type HealthBased struct{}
//
//func (hb *HealthBased) SelectInstance(instances []Instance, clientIP string) Instance {
//	// Placeholder for health-based (return the first instance for now)
//	return instances[0]
//}
