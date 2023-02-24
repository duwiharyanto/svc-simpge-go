FROM asia.gcr.io/uii-cloud-project/hcm/go/golang:1.17-ubuntu

LABEL MAINTAINER="Ari Satrio<ari.satrio@uii.ac.id>"

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu
USER ubuntu

ADD main /

EXPOSE 80

CMD ["/main"]
