# FROM asia.gcr.io/uii-cloud-project/hcm/docker/ubuntu:latest

# LABEL MAINTAINER="Ahmad Haris Fahmi<ahmadharis@uii.ac.id>"

# ADD main /

# EXPOSE 80

# CMD ["/main"]

FROM asia.gcr.io/uii-cloud-project/hcm/go/golang:1.17-ubuntu

LABEL MAINTAINER="Ari Satrio<ari.satrio@uii.ac.id>"

ADD main /

EXPOSE 80

CMD ["/main"]
