FROM golang:1.18.3-alpine3.16 as Builder

ARG VERSION

WORKDIR /go/src/github.com/hatech/backup/
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . .
RUN CGO_ENABLED=0 go build --mod=vendor \
        -ldflags \
        "-X 'github.com/hatech/backup/pkg/version.Version=$VERSION'" \
        -a -o ./bin/app ./cmd/server/

FROM alpine:3.16
WORKDIR /

COPY --from=builder /go/src/github.com/hatech/backup/bin/app /hatech/
COPY --from=builder /go/src/github.com/hatech/backup/config.yaml /hatech/config/


VOLUME /opt/hyperkuber/helm/
CMD ["/hk/app","-c","/hk/config/config.yaml"]