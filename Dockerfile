FROM alpine:3.4

MAINTAINER k.guan@hnair.com

ENV KONG_URL=http://223.202.32.56:8055/
ENV KONG_ADMIN_URL=http://223.202.32.56:8056/

COPY cloud-platform /usr/bin/cloud-platform

EXPOSE 9981

ENTRYPOINT ["/usr/bin/cloud-platform"]



