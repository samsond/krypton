package parser

type ResourceRequirements struct {
    Limits   ResourceSpec
    Requests ResourceSpec
}

type ResourceSpec struct {
    Memory string
    CPU    string
}

type AppDeployment struct {
    Name      string
    Namespace string
    Replicas  int
    Image     string
    Args      []string
    Ports     map[string]int
    Env       map[string]string
    Resources *ResourceRequirements
    Storage   *StorageConfig
}


type StorageConfig struct {
    Volume string
    Size   string
}
