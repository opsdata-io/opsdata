FROM ubuntu

MAINTAINER mmattox@support.tools

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -yq --no-install-recommends \
    apt-utils \
    curl \
    openssl \
    nano \
    mysql-client \
    iputils-ping \
    ca-certificates \
    git \
    bc \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /root

COPY *.sh /root/
COPY *.sql /root/
RUN chmod +x /root/*.sh
CMD /root/run.sh
