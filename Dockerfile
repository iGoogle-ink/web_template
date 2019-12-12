FROM busybox
MAINTAINER Jerry "jerry@igoogle.ink"
RUN go mod tidy
COPY main /usr/main
COPY cmd/web_template.json /usr/web_template.json
CMD ["/usr/main", "-conf", "/usr/web_template.json"]
EXPOSE 80