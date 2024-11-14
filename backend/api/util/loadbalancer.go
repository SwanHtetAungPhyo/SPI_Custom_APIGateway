package util

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
