FROM golang:1.21 as build

COPY . /app/

WORKDIR /app/

RUN CGO_ENABLED=0 go build -o temporal-worker ./worker

FROM gcr.io/distroless/static-debian11:nonroot

COPY --from=build /app/temporal-worker  /app/temporal-worker

ENTRYPOINT ["/app/temporal-worker"]
