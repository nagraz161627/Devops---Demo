FROM  golang:1.17-alpine3.15 as build
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download 
COPY *.go .
RUN go build -o demo

FROM alpine:3.15
WORKDIR /app
COPY --from=build /build/demo .
USER 1001
CMD [ "./demo" ]
