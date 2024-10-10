package parser

// Resource is an interface that represents any Kubernetes resource that can be parsed from DSL
type Resource interface {
    GetName() string
    GetNamespace() string
}
