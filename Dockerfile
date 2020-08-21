FROM alpine
MAINTAINER Jerry "jerry@gopay.ink"
COPY main /usr/bin/main
EXPOSE 80

# CMD ["main", "-config", "/config/vbs9010_app_api.yaml"]
# docker run -d --restart=unless-stopped -v /etc/localtime:/etc/localtime:ro -p 2233:2233 --name=web_template web_template:latest