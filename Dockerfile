FROM alpine:3.4

MAINTAINER k.guan@hnair.com

ENV KONG_URL=http://13.76.42.81:8000
ENV KONG_ADMIN_URL=http://13.76.42.81:8001

COPY service-cloud /usr/bin/service-cloud

EXPOSE 9981

ENTRYPOINT ["/usr/bin/service-cloud"]



