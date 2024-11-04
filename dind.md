# Test Using Docker In Docker

## Context

It may be appropriate to control the daemon used for tests.

For instance, the following error has been reported in some cases. At this time the cause of the bug is not clear.
```
Error response from daemon: Pull failed due to an unauthenticated request. Registry Access Management is enabled, which requires pulls to be authenticated. Please run `docker login`, or contact your administrators if this is unexpected.
```

This guide enables a developer to stand up a docker daemon inside docker (DinD).

## Steps
1. Determine a directory where your DinD ssl certs will live and create the directory. Also set the associated `DOCKER` env var. This demo arbitrarily uses `~/.dind/certs`.
   ```bash
   export DOCKER_CERT_PATH=~/.dind/certs
   mkdir -p $DOCKER_CERT_PATH
   ```

2. Create ssl certs.
    ```bash
    openssl req -x509 -newkey rsa:4096 -keyout $DOCKER_CERT_PATH/key.pem -out $DOCKER_CERT_PATH/cert.pem -sha256 -days 3650 -nodes -subj '/CN=localhost'
    ```

3. Run docker in docker (DinD) with mounted certs and mapped port.
   ```bash
   docker run -d --privileged --name dind -p 2375:2376 -v $DOCKER_CERT_PATH:/certs/client \
     -p 32768:32768 \
     -p 32769:32769 \
     -p 32770:32770 \
     -p 32771:32771 \
     -p 32772:32772 \
     -p 32773:32773 \
     -p 32774:32774 \
     -p 32775:32775 \
     -p 32776:32776 \
     -p 32777:32777 \
     -p 32778:32778 \
     -p 32779:32779 \
     -p 32780:32780 \
     docker:dind 
   ```

4. Set the following environment variables. `DOCKER_CERT_PATH` may be set from step 1.
   ```bash
   export DOCKER_TLS_VERIFY=1
   export DOCKER_HOST=tcp://localhost:2375
   # export DOCKER_CERT_PATH=~/.dind/certs
   ```
   
5. Run `docker info -f json | jq '.ServerVersion'`. There should be a value such as `27.3.1`.
