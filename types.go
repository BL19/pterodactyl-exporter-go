package main

type ServerListResponse struct {
	Object string             `json:"object"`
	Data   []ServerListServer `json:"data"`
	Meta   Meta               `json:"meta"`
}

type ServerListServer struct {
	Object     string           `json:"object"`
	Attributes ServerAttributes `json:"attributes"`
}

type ServerAttributes struct {
	ServerOwner   bool          `json:"server_owner"`
	Identifier    string        `json:"identifier"`
	UUID          string        `json:"uuid"`
	Name          string        `json:"name"`
	Node          string        `json:"node"`
	SFTPDetails   SFTPDetails   `json:"sftp_details"`
	Description   string        `json:"description"`
	Limits        Limits        `json:"limits"`
	FeatureLimits FeatureLimits `json:"feature_limits"`
	IsSuspended   bool          `json:"is_suspended"`
	IsInstalling  bool          `json:"is_installing"`
	Relationships Relationships `json:"relationships"`
}

type SFTPDetails struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type Limits struct {
	Memory int `json:"memory"`
	Swap   int `json:"swap"`
	Disk   int `json:"disk"`
	IO     int `json:"io"`
	CPU    int `json:"cpu"`
}

type FeatureLimits struct {
	Databases   int `json:"databases"`
	Allocations int `json:"allocations"`
	Backups     int `json:"backups"`
}

type Relationships struct {
	Allocations AllocationList `json:"allocations"`
}

type AllocationList struct {
	Object string       `json:"object"`
	Data   []Allocation `json:"data"`
}

type Allocation struct {
	Object     string               `json:"object"`
	Attributes AllocationAttributes `json:"attributes"`
}

type AllocationAttributes struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	IPAlias   string `json:"ip_alias"`
	Port      int    `json:"port"`
	Notes     string `json:"notes"`
	IsDefault bool   `json:"is_default"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int         `json:"total"`
	Count       int         `json:"count"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	Links       interface{} `json:"links"`
}

type Stats struct {
	Object     string          `json:"object"`
	Attributes StatsAttributes `json:"attributes"`
}

type StatsAttributes struct {
	CurrentState string         `json:"current_state"`
	IsSuspended  bool           `json:"is_suspended"`
	Resources    StatsResources `json:"resources"`
}

type StatsResources struct {
	MemoryBytes    int     `json:"memory_bytes"`
	CPUAbsolute    float64 `json:"cpu_absolute"`
	DiskBytes      int     `json:"disk_bytes"`
	NetworkRXBytes int     `json:"network_rx_bytes"`
	NetworkTXBytes int     `json:"network_tx_bytes"`
	Uptime         int     `json:"uptime"`
}
