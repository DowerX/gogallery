FROM golang:alpine AS build

RUN apk --no-cache add git ca-certificates

COPY goGallery /goGallery
WORKDIR /goGallery
RUN go env -w GO111MODULE=off; \
    go get -u github.com/gorilla/mux; \
    go get -u golang.org/x/exp/slices; \
    CGO_ENABLED=0 \
    go build \
    -installsuffix "static" \
    -o goGallery

FROM scratch AS final
COPY --from=build /goGallery/goGallery /goGallery
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/goGallery"]
