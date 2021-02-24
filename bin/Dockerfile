FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM busybox
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

COPY httpfwd /bin/httpfwd
ENTRYPOINT ["/bin/httpfwd"]
CMD [""]

EXPOSE 8080

