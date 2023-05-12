FROM golang:1.20-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./
RUN go build -o /hathor ./cmd/hathor


FROM scratch

ARG PORT

WORKDIR /app

COPY --from=build /hathor /hathor

EXPOSE $PORT

CMD [ "/hathor" ]
