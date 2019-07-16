FROM debian:stretch-20171009
MAINTAINER tiru@hasura.io

# Install openssh server
RUN apt-get update \
 && apt-get install -y upx-ucl binutils curl openssh-server git jq gettext-base \
 # && curl -o /tmp/docker-18.09 'https://download.docker.com/linux/static/stable/x86_64/docker-18.09.7.tgz' \
 && curl -o /tmp/docker-18.09 'https://download.docker.com/linux/static/nightly/x86_64/docker-0.0.0-20190714170256-22f3318.tgz' \
 && tar -xf /tmp/docker-18.09 -C /tmp \
 && mv /tmp/docker/docker /bin/docker \
 && rm -rf /tmp/docker-18.09 /tmp/docker \
 && strip --strip-unneeded /bin/docker \
 && chmod a+x /bin/docker \
 && upx /bin/docker \
 && curl -o /bin/kubectl 'https://storage.googleapis.com/kubernetes-release/release/v1.14.3/bin/linux/amd64/kubectl' \
 && strip --strip-unneeded /bin/kubectl \
 && chmod a+x /bin/kubectl \
 && upx /bin/kubectl \
 && curl -o /tmp/helm.tar.gz 'https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-linux-amd64.tar.gz' \
 && tar -xzf /tmp/helm.tar.gz -C /tmp \
 && mv /tmp/linux-amd64/helm /bin/helm \
 && rm -rf /tmp/helm.tar.gz /tmp/linux-amd64 \
 && strip --strip-unneeded /bin/helm \
 && chmod a+x /bin/helm \
 && upx /bin/helm \
 && curl -o /usr/bin/mo 'https://raw.githubusercontent.com/tests-always-included/mo/master/mo' \
 && chmod a+x /usr/bin/mo \
 && apt-get purge -y --auto-remove upx-ucl binutils \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

# sshd config
RUN mkdir /var/run/sshd

ADD no-interactive-login.sh /sshd-lib/
ADD pre_receive.sh /sshd-lib/
ADD sshd_config /sshd-lib/
ADD start_sshd.sh /sshd-lib/

EXPOSE 22
