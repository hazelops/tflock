FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/tflock .
FROM scratch AS bin

COPY --from=build /out/tflock /tflock
ENTRYPOINT /tflock
