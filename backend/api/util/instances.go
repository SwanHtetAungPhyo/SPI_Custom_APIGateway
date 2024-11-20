package util

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync/atomic"
	"time"
)

const (
	ROUND_ROBIN                = "round-robin"
	LEAST_CONNECTIONS          = "least-connections"
	LEAST_RESPONSE_TIME        = "least-response-time"
	IP_HASH                    = "ip-hash"
	WEIGHTED_ROUND_ROBIN       = "weighted-round-robin"
	WEIGHTED_LEAST_CONNECTIONS = "weighted-least-connections"
	RANDOM                     = "random"
	CONSISTENCY_HASHING        = "consistency-hashing"
	CLIENT_IP_AFFINITY         = "client-ip-affinity"
	HEALTH_BASED               = "health-based"
)

type InstanceAlgorithm struct {
	Algorithm string `json:"type"`
}

type SelectedInstance struct {
	Port int `json:"port"`
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano())) // Seed once globally

func GetCurrentInstance(algorithm *InstanceAlgorithm, instances []int, clientIp string) *SelectedInstance {
	if algorithm == nil {
		return nil
	}

	var currentInstance SelectedInstance
	switch algorithm.Algorithm {
	case ROUND_ROBIN:
		robin := &RoundRobin{
			LastIndex: 0,
		}

		selectedPort := robin.SelectInstance(instances)
		currentInstance = SelectedInstance{
			Port: instances[selectedPort],
		}

	case RANDOM:
		random := &Random{}

		selectedPort := random.SelectInstance(instances)
		currentInstance = SelectedInstance{
			Port: instances[selectedPort],
		}
	case IP_HASH:
		ip := &IPHash{}
		selectedPort := ip.SelectInstance(instances, clientIp)
		currentInstance = SelectedInstance{
			Port: instances[selectedPort],
		}
	}

	return &currentInstance
}

type RoundRobin struct {
	LastIndex int32
}

func (rr *RoundRobin) SelectInstance(instances []int) int {

	newIndex := atomic.AddInt32(&rr.LastIndex, 1) % int32(len(instances))
	return int(newIndex)
}

type Random struct{}

func (r *Random) SelectInstance(instances []int) int {
	// Reusing the global RNG instance (no need to re-seed each time)
	index := rng.Intn(len(instances))
	fmt.Println("Selected random index:", index)
	return index
}

// IPHash IP_HASH Load Balancer
type IPHash struct{}

func (ih *IPHash) SelectInstance(instances []int, clientIp string) int {
	hash := ih.hashIp(clientIp)
	index := int(hash % uint32(len(instances)))
	return index
}

func (ih *IPHash) hashIp(clientIp string) uint32 {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(clientIp))
	if err != nil {
		return 0
	}
	return hash.Sum32()
}

//// LeastConnections LEAST_CONNECTIONS Load Balancer (Placeholder for actual implementation)
//type LeastConnections struct {
//	// In a real scenario, you'd track active connections for each instance
//}

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
