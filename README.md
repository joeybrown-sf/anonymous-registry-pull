# anonymous-registry-pull

Repo for reproducing anonymous registry call.

Attempt to reproduce the error by running the app:
```bash
go run github.com/joeybrown-sf/anonymous-registry-pull
```

Then try running on docker in docker:
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

