FROM denderello/docker-golang

ADD code/bin/hello-server /opt/hello-server

CMD ["/opt/hello-server"]
