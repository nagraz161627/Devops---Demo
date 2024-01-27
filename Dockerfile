FROM  golang:alpine3.18 as build
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download 
COPY *.go .
RUN go build -o demo

FROM alpine:3.18
WORKDIR /app
COPY --from=build /build/demo .
USER 1001
CMD [ "./demo" ]
