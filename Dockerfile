
FROM scratch

COPY hyperping_exporter /hyperping_exporter

ENTRYPOINT ["/hyperping_exporter"]

EXPOSE 9312
