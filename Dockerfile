# Create builder image
FROM golang:alpine as builder
ARG GEOLOCATION_GITLAB_TOKEN

WORKDIR /go/src/gitlab.com/dpcat237/geolocation

# Download dependencies
RUN apk update && apk upgrade && apk add git
RUN git config --global url."http://dpcat237:${GEOLOCATION_GITLAB_TOKEN}@gitlab.com/".insteadOf "https://gitlab.com/"
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only -v

# Build the binary
COPY . .
RUN go install .
EXPOSE 3000 5000
RUN addgroup usgeolocation && adduser -S -G usgeolocation usgeolocation
USER usgeolocation
ENTRYPOINT ["/go/bin/geolocation"]
CMD ["grpc"]
