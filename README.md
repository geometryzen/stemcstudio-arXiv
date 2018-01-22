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

# Launch

```bash
./stemcstudio-search
```

# Test

```bash
curl -XPOST 'localhost:8081/search' -d '{"query":"webgl"}'
```
