# Small Utility Containers

...for testing and stuff. I put here the programs that I say "I wish there was a container that does X" more than once to myself, to be containerized.

## Containers

### request-forever

Send infinitely many HTTP requests to the specified URL, with the specified method, in specified interval, etc.

#### Usage

Below is a basic Kubernetes deployment config to use the program:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: request-forever-deployment
  labels:
    app: request-forever
spec:
  replicas: 1
  selector:
    matchLabels:
      app: request-forever
  template:
    metadata:
      labels:
        app: request-forever
    spec:
      containers:
      - name: request-forever-container
        image: ghcr.io/ozansz/utils/request-forever:latest
        args:
          - "--url"
          - "https://example.org"
          - "--method"
          - "GET"
          - "--interval"
          - "10s"
          - "--http-timeout"
          - "1m"
        command:
          - "/ko-app/request-forever"
```

