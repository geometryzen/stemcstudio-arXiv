# stemcstudio-search

Microservice providing the search capability in STEMCstudio

# Build

```bash
go build
```

# Configure

```base
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
```

# Launch as a binary

```bash
./stemcstudio-search
```

# Test

```bash
curl -XPOST 'localhost:8081/search' -d '{"query":"webgl"}'
```

# Minikube, Kubernetes, and Docker

```bash
minikube version
```

```bash
minikube start
```

In the folder containing the Dockerfile,

```bash
eval $(minikube docker-env)
```

```bash
./build.sh
```

Create a deployment all in one.

```bash
kubectl run stemcstudio-search --image=stemcstudio-search:v1 --port=8081
```

```bash
kubectl get deployments
kubectl get pods
```

```bash
kubectl expose deployment stemcstudio-search --type=LoadBalancer
```

```bash
minikube service stemcstudio-search --url
```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: stemcstudio-search-secret
type: Opaque
data:
  username:
  password: 
```

