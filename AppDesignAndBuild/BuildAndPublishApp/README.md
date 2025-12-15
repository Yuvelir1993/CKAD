```bash
docker run -d -p 5000:5000 --restart always --name registry registry:3
docker build -t ckad-1/site:1.0.0 .
docker run --rm -p 8080:80 ckad-1/site:1.0.0
```