# Build stage
FROM golang:1.12
RUN go get -v -d github.com/go-delve/delve/cmd/dlv
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /dlv github.com/go-delve/delve/cmd/dlv
RUN ls -al /dlv
CMD /dlv
EXPOSE 2345
