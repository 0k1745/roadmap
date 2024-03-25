FROM golang:1.22 AS build-stage

WORKDIR /app

#COPY go.mod go.sum ./
COPY go.mod ./
#RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /process

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /process /process

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/process"]