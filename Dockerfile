FROM golang:alpine AS build
WORKDIR /src
COPY . .
RUN apk add --no-cache git && \
    CGO_ENABLED=0 go build -o /bin/app .

FROM scratch
COPY --from=build /bin/app /bin/app
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/bin/app"]
