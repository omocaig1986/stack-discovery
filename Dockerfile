FROM golang:1.13.4-alpine3.10 as build
LABEL stage=builder

RUN apk update && apk add curl git build-base

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go
COPY . .

# install deps and build
RUN cd src/discovery && dep ensure
RUN go build discovery

FROM alpine:3.10

LABEL org.label-schema.license="MIT" \
    org.label-schema.vcs-url="https://gitlab.com/p2p-faas/stack-discovery" \
    org.label-schema.vcs-type="Git" \
    org.label-schema.name="p2p-fog/discovery" \
    org.label-schema.vendor="gabrielepmattia" \
    org.label-schema.docker.schema-version="1.0"

WORKDIR /home/app
COPY --from=build /go/discovery .

RUN mkdir -p /data

# set permissions
# RUN addgroup -S app && adduser -S -g app app
# RUN chown -R app:app ./
# USER app

EXPOSE 19000

CMD ["./discovery"]
