FROM registry.access.redhat.com/rhel7:latest

MAINTAINER Daniel Tschan <tschan@puzzle.ch>

RUN rpm -ihv https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && \
    yum -y install openssl jq && \
    yum clean all && \
    mkdir -p /srv/.well-known/acme-challenge /usr/local/letsencrypt /var/lib/letsencrypt && \
    chmod 775 /srv/.well-known/acme-challenge && \  
    cd /usr/local/bin && \
    curl -O https://console.appuio.ch/console/extensions/clients/linux/oc && \
    chmod 755 /usr/local/bin/oc

ADD acme-tiny/acme_tiny.py bin/* src/* ca/* /usr/local/letsencrypt/

RUN yum -y --enablerepo=rhel-7-server-optional-rpms install golang-bin && \   
    cd /usr/local/letsencrypt && \
    go build letsencrypt.go sh.go && \
    yum -y history undo last && \  
    yum clean all

USER 1001
ENV HOME /var/lib/letsencrypt

EXPOSE 8080

WORKDIR /tmp
CMD ["/usr/local/letsencrypt/letsencrypt"]
