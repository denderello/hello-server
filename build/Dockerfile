FROM denderello/docker-golang

RUN apt-get update
RUN apt-get install -y \
  build-essential \
  git

VOLUME ["/opt/hello-server"]

ENTRYPOINT ["/bin/bash", "-c"]
