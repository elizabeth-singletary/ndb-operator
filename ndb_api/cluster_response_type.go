package ndb_api

// Define the ClusterResponse type
type ClusterResponse struct {
	ID                   string            `json:"id"`
	Name                 string            `json:"name"`
	UniqueName           string            `json:"uniqueName"`
	IPAddresses          []string          `json:"ipAddresses"`
	FQDNs                interface{}       `json:"fqdns"`
	NxClusterUUID        string            `json:"nxClusterUUID"`
	Description          string            `json:"description"`
	CloudType            string            `json:"cloudType"`
	DateCreated          string            `json:"dateCreated"`
	DateModified         string            `json:"dateModified"`
	OwnerID              string            `json:"ownerId"`
	Status               string            `json:"status"`
	Version              string            `json:"version"`
	HypervisorType       string            `json:"hypervisorType"`
	HypervisorVersion    string            `json:"hypervisorVersion"`
	Properties           []ClusterProperty `json:"properties"`
	ReferenceCount       int               `json:"referenceCount"`
	Username             interface{}       `json:"username"`
	Password             interface{}       `json:"password"`
	CloudInfo            interface{}       `json:"cloudInfo"`
	ResourceConfig       ResourceConfig    `json:"resourceConfig"`
	ManagementServerInfo interface{}       `json:"managementServerInfo"`
	EntityCounts         interface{}       `json:"entityCounts"`
	Healthy              bool              `json:"healthy"`
}

// Define the Property type for the properties field
type ClusterProperty struct {
	RefID       interface{} `json:"ref_id"`
	Name        string      `json:"name"`
	Value       string      `json:"value"`
	Secure      bool        `json:"secure"`
	Description interface{} `json:"description"`
}

// Define the ResourceConfig type for resourceConfig field
type ResourceConfig struct {
	StorageThresholdPercentage float64 `json:"storageThresholdPercentage"`
	MemoryThresholdPercentage  float64 `json:"memoryThresholdPercentage"`
}
