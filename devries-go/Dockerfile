# Latest Ubuntu image
FROM ubuntu:latest

# Make a working directory, as it would be binded with the current language one
WORKDIR /deez

# Basic system-level dependencies
RUN apt-get update && apt install -y make git curl 

RUN curl -L -O  https://go.dev/dl/go1.20.5.linux-amd64.tar.gz && \
    tar xfz go1.20.5.linux-amd64.tar.gz -C /

ENV PATH="${PATH}:/go/bin"

# Commands for docker run
CMD make clean && \
    make ready

