FROM docker.io/library/golang:1.24
WORKDIR /work
COPY go.mod go.sum *.go *.html ./
RUN CGO_ENABLED=0 go build -ldflags='-s -w'

FROM scratch
COPY --from=0 /work/signal-gateway /
ENTRYPOINT ["/signal-gateway"]
EXPOSE 80
