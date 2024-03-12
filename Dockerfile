
FROM alpine:latest as certs

RUN apk --update add ca-certificates

COPY hyperping_exporter /hyperping_exporter

ENTRYPOINT ["/hyperping_exporter"]

EXPOSE 9312
