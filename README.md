
# Krypton CLI

`Krypton` is a CLI tool designed for generating Kubernetes YAML manifests from a custom Domain-Specific Language (DSL). It allows developers to define Kubernetes resources in a more human-readable way using `.kp` files, which are then parsed and converted into YAML files.

## Features
* DSL Parsing: Parse custom .kp files to define Kubernetes deployments.
* YAML Generation: Generate deployment YAML manifests from the parsed DSL.
* Extensible: The code structure allows for easy extension to support other Kubernetes resources beyond deployments.
* Template-based: Uses Go templates for generating YAML, making it easy to customize the output format.

## Getting Started

### Prerequisites

- Go 1.23+ installed

### Building the CLI

To build the `kptn` binary, run:

```bash
make build
```

Display the version of kptn:
```bash
./kptn version
```

### Usage
Define Your Application in the DSL:

1. Write your application specification using the custom DSL in a .kp file:

```bash
deploy app "my-app" {
    namespace: "default";
    replicas: 3;
    image: "my-app:v1.0";
    ports {
        http: 8080;
        metrics: 2112;
    }
    env {
        DATABASE_URL: "postgres://user:password@host/db";
    }
    resources {
        limits {
            memory: "512Mi";
            cpu: "500m";
        }
        requests {
            memory: "256Mi";
            cpu: "250m";
        }
    }
    storage {
        volume: "my-app-data";
        size: "5Gi";
    }
}

```

Refer in the examples/ folder as basic_app.kp i.e. [examples/basic_app.kp](./examples/basic_app.kp)


2. Generate Kubernetes Manifests:

Use the generate command to produce Kubernetes YAML:

```bash
./kptn generate examples/basic_app.kp --output=generated.yaml
```

This will read the DSL file, parse it, and generate a generated.yaml file with the Kubernetes Deployment manifest.

3. Output

The generated YAML will be similar to the following:


```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-app
        image: my-app:v1.0
        ports:
        - name: "http"
          containerPort: 8080
        - name: "metrics"
          containerPort: 2112
        env:
        - name: "DATABASE_URL"
          value: "postgres://user:password@host/db"
        resources:
          limits:
            memory: "512Mi"
            cpu: "500m"
          requests:
            memory: "256Mi"
            cpu: "250m"
        volumeMounts:
        - name: "my-app-data"
          mountPath: /data
      volumes:
      - name: "my-app-data"
        persistentVolumeClaim:
          claimName: "my-app-data"
```

## Project Structure
* cmd/krypton: Contains the main entry point for the kptn CLI.
* * pkg/parser: Handles the parsing logic for .kp files.
* pkg/generator: Contains the logic for converting parsed structures into YAML using templates.
* pkg/generator/templates: Holds the Go templates used for generating Kubernetes YAML.
* examples/: Directory containing example .kp files and their corresponding generated YAML files.

## Extending the DSL
To extend the DSL to support more Kubernetes resources, follow these steps:

1. Define a new parser function in the pkg/parser package.
2. Add a new Go template in the pkg/generator/templates directory.
3. Update the generator package to use the new template for the corresponding resource.
4. Add the necessary parsing logic in the NewGenerateCommand to handle the new resource type.

## Contributing
Feel free to open issues or submit pull requests. Contributions are always welcome!

## License

This project will be licensed under the Apache License, Version 2.0. The LICENSE file will be available soon, providing the full terms and conditions for usage, including permissions, conditions, and limitations.