package main

type TypeInfo struct {
	Namespace     string `json:"namespace"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	PodCount      int32  `json:"pod_count"`
	Container     string `json:"container"`
	CPURequest    string `json:"cpu_request"`
	CPULimit      string `json:"cpu_limit"`
	CPUUsage      string `json:"cpu_usage"`
	MemoryRequest string `json:"memory_request"`
	MemoryLimit   string `json:"memory_limit"`
	MemoryUsage   string `json:"memory_usage"`
}
