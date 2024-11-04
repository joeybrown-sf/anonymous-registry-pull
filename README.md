# anonymous-registry-pull

Repo for reproducing anonymous registry call.

Build the exe `GOOS=linux GOARCH=amd64 go build github.com/joeybrown-sf/anonymous-registry-pull`

```bash
GOOS=linux GOARCH=amd64 go build github.com/joeybrown-sf/anonymous-registry-pull
docker build -t anon-repro .
docker run -d --privileged --name anon-repro anon-repro
```

```bash
docker exec -it anon-repro /bin/sh
# ...
./anonymous-registry-pull
```

