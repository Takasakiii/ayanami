FROM --platform=amd64 golang:alpine AS build
WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

ENV CGO_ENABLED=1
RUN go install github.com/mattn/go-sqlite3
RUN go generate
RUN GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o ayanami

FROM alpine:latest AS prod
WORKDIR /app

COPY --from=build /app/ayanami /app/ayanami

ENTRYPOINT ["./ayanami"]
