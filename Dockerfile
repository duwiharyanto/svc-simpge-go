FROM registry.uii.ac.id/uii-project/hcm/backend/os/go:1.17-ubuntu

LABEL MAINTAINER="Ari Satrio<ari.satrio@uii.ac.id>"

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu
USER ubuntu

ADD main /

EXPOSE 80

CMD ["/main"]
