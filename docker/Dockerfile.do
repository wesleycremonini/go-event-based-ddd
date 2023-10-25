FROM golang:alpine as build

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /go/src/app

COPY ../go.mod /go/src/app
COPY ../go.sum /go/src/app
RUN go mod download

COPY ../ /go/src/app

RUN CGO_ENABLED=0 go build -v -ldflags "-s -w" -o /go/bin/app /go/src/app/cmd/api

FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]