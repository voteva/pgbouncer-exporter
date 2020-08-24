FROM scratch
ADD bin/pgbouncer-exporter /
CMD ["/pgbouncer-exporter"]