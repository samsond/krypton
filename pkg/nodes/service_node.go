package nodes

type ServiceNode struct {
	Name      string
	Namespace string
	Ports     map[int]int // Map of `port,targetPort`
	Labels    map[string]string
}

func (n *ServiceNode) NodeType() string {
	return "Service"
}
