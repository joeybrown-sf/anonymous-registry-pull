FROM docker:dind

RUN mkdir -p /etc/docker

COPY . .
