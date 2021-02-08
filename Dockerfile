FROM asia.gcr.io/uii-cloud-project/panduba/alpine:3.12-musl-1

RUN apk add --no-cache bash

ADD main /

EXPOSE 80

CMD ["/main"]