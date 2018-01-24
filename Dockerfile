FROM alpine:3.7 as certs
RUN apk update && apk add ca-certificates

FROM alpine:3.7
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
ADD stemcstudio-search /usr/bin/stemcstudio-search
ENTRYPOINT ["stemcstudio-search"]
