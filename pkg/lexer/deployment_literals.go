package lexer

// contains the literals used in the deployment DSL
const (
	deployAppPrefix     = "deploy app"
	namespacePrefix     = "namespace:"
	replicasPrefix      = "replicas:"
	imagePrefix         = "image:"
	envPrefix           = "env {"
	envTokenValue       = "env"
	resourcesPrefix     = "resources {"
	resourcesTokenValue = "resources"
	limitsPrefix        = "limits {"
	limitsTokenValue    = "limits"
	requestsPrefix      = "requests {"
	requestsTokenValue  = "requests"
	memoryPrefix        = "memory:"
	cpuPrefix           = "cpu:"
	storagePrefix       = "storage {"
	storageTokenValue   = "storage"
	volumePrefix        = "volume:"
	sizePrefix          = "size:"
)
