FROM golang:1.21-alpine AS BUILDER
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app

FROM scratch
COPY --from=BUILDER /bin/app /
COPY templates /templates
EXPOSE 8080
ENTRYPOINT ["/app"]
