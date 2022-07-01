ARG GO_VERSION=1.16.6

FROM golang:${GO_VERSION}-alpine as builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./../database database
COPY ./../models models
COPY ./../repositories repositories
COPY ./../launchers-person-service launchers-person-service
COPY ./../models-proto model-proto
COPY ./../servers servers

FROM alpine:3.11
WORKDIR /usr/bin

COPY --from=builder /go/bin .

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /generalrestapi

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT [ "/launchers-person-service" ]