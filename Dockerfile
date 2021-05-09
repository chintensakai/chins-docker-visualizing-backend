FROM golang:alpine AS build-env
WORKDIR /opt/apps/docker-visualizing
COPY . /opt/apps/docker-visualizing

ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct

RUN go build -v -o /opt/apps/docker-visualizing/app

FROM alpine
COPY --from=build-env  /opt/apps/docker-visualizing/app /usr/local/bin/app
COPY --from=build-env /opt/apps/docker-visualizing/conf/app.ini /usr/local/bin/
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
EXPOSE 9281
CMD ["app"]