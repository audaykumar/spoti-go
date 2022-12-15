FROM golang:1.19.4 as build

ARG WORKDIR=/go/src/github.com/audaykumar/spoti-go/

WORKDIR "${WORKDIR}"

COPY go.mod go.sum "${WORKDIR}"
# RUN go mod download

COPY . "${WORKDIR}"

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o /spoti-go

ENTRYPOINT ["/spoti-go"]

FROM scratch

WORKDIR /app

COPY --from=build /spoti-go .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/app/spoti-go"]
