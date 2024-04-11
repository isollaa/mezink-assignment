ARG GO_VERSION=1.22
# Builder
FROM golang:${GO_VERSION}-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make build-base

WORKDIR /app

COPY . .

RUN GOFLAGS="-buildvcs=false" go build -o goBinary .

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && apk --no-cache add ca-certificates && \
    apk --update --no-cache add tzdata

ENV TZ=UTC

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/goBinary /app
COPY --from=builder /app/migration /app/migration

CMD /app/goBinary
