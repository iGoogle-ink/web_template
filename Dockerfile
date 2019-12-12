FROM busybox
MAINTAINER Jerry "jerry@igoogle.ink"
COPY application /usr/application
COPY cmd/web_template.json /usr/web_template.json
CMD ["/usr/application", "-conf", "/usr/web_template.json"]
EXPOSE 80

# docker run -d --restart=unless-stopped -v /etc/localtime:/etc/localtime:ro -p 2233:8080 --name web_template web_template:latest