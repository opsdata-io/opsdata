FROM sinlead/drone-kubectl:latest
MAINTAINER Matthew Mattox <mmattox@support.tools>

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -yq --no-install-recommends \
    apt-utils \
    curl \
    wget \
    openssh-client \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

## Install kubectl
RUN curl -sLf https://storage.googleapis.com/kubernetes-release/release/$(curl -ks https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl > /usr/local/bin/kubectl && \
chmod +x /usr/local/bin/kubectl

## Install Helm3
RUN HVER=$(curl -sSL https://github.com/kubernetes/helm/releases | sed -n '/Latest release<\/a>/,$p' | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' | head -1) && \
wget -k https://get.helm.sh/helm-"$HVER"-linux-amd64.tar.gz && \
tar -zvxf helm-* && \
cd linux-amd64 && \
chmod +x helm && \
mv helm /usr/local/bin/helm

# Install jq
RUN curl -sLf http://stedolan.github.io/jq/download/linux64/jq > /usr/local/bin/jq && \
chmod +x /usr/local/bin/jq
