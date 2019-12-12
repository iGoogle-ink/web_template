FROM busybox
MAINTAINER Jerry "jerry@igoogle.ink"
COPY application /usr/application
COPY cmd/web_template.json /usr/web_template.json
CMD ["/usr/application", "-conf", "/usr/web_template.json"]
EXPOSE 80