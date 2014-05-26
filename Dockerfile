FROM denderello/docker-golang

ADD build/bin/hello-server /opt/hello-server

CMD ["/opt/hello-server"]

EXPOSE 8000
