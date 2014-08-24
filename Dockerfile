FROM busybox:ubuntu-14.04

ADD hello-server /opt/hello-server

CMD ["/opt/hello-server"]

EXPOSE 8000
