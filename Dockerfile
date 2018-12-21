############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/scm.bluebeam.com/stu/golang-template/
WORKDIR $GOPATH/src/scm.bluebeam.com/stu/golang-template/
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/scm.bluebeam.com/stu/golang-template
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/scm.bluebeam.com/stu/golang-template /go/bin/scm.bluebeam.com/stu/golang-template
# Run the hello binary.
ENTRYPOINT ["/go/bin/scm.bluebeam.com/stu/golang-template"]