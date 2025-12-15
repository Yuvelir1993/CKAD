Build a tiny Go HTTP server into a minimal runtime image.
[Docker multistage](https://docs.docker.com/build/building/multi-stage/) docu.

Stage 1: golang:1.22-alpine builds /out/server
Stage 2: alpine:3.20 runs /server
Expose 8080
Image name: ckad-1/multistage:v1

Verify

```bash
docker build --target production -t ckad-1/multistage:1.0.0 .
docker run -d -p 8080:8080 --restart always --name go-server ckad-1/multistage:1.0.0
curl -i localhost:8080
```