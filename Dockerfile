# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/Gin/node-tracker

WORKDIR /go/src/github.com/Gin/node-tracker
RUN go install -v github.com/Gin/node-tracker/cmd

EXPOSE 8000
