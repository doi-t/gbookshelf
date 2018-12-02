FROM golang:1.11.2 AS builder

ENV DEP_VERSION 0.5.0
ENV USER doi-t
ENV REPO gbookshelf
ENV TARGET ./cmd/gbookshelf-server

ADD https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/${USER}/${REPO}
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server ${TARGET}

# FROM scratch
FROM alpine
RUN apk update && \
   apk add ca-certificates && \
   update-ca-certificates && \
   rm -rf /var/cache/apk/*
WORKDIR /
COPY --from=builder /server /server

ENTRYPOINT ["./server"]
