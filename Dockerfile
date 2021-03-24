FROM golang:1.16 AS build

WORKDIR /root/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o dist/

FROM gcr.io/distroless/base

WORKDIR /root/
COPY --from=build /root/app/dist .
COPY --from=build /root/app/.env .

EXPOSE 8080

CMD ["./club-portal"]
