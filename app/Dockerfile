FROM golang:alpine as builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /build/app
# Get depedences

COPY go.mod ./
RUN go mod download

# Run Testss
COPY . ./
RUN go test -v ./...

# Build app
RUN go build -o myapp

FROM alpine
COPY --from=builder /build/app/myapp ./myapp
EXPOSE 3000
CMD ["./myapp"]

