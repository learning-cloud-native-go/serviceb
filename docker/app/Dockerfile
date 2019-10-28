# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /serviceb
RUN apk update && apk add --no-cache gcc musl-dev git
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app


# Deployment environment
# ----------------------
FROM alpine
RUN apk update
COPY --from=build-env /serviceb/bin/app /serviceb/
EXPOSE 8080
CMD ["/serviceb/app"]