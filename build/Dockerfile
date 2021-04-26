FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.10-3 AS build

WORKDIR /build

#COPY ./demo/ ./demo/

ADD gin-demo ./
RUN go build -o bin/gin_demo cmd/gin-demo/main.go

#FROM alpine:latest
#FROM scratch

#COPY --from=build /home/roger/k8s/build/hello_world /usr/local/bin/hello_world

#RUN chmod +x /usr/local/bin/hello_world
RUN chmod +x build/start.sh

EXPOSE 8080

#CMD ["/usr/local/bin/hello_world"]
CMD ["/build/build/start.sh"]