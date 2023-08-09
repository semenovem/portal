FROM centos:centos7

RUN groupadd --gid 2000 minio \
  && useradd --uid 2000 --gid minio --shell /bin/bash --create-home minio

USER 2000

ENV HOME=/home/minio

RUN curl https://dl.min.io/client/mc/release/linux-amd64/mc \
      --create-dirs \
      -o $HOME/minio-binaries/mc && \
  chmod +x $HOME/minio-binaries/mc && \
  mkdir $HOME/.mc

ENV PATH=$PATH:$HOME/minio-binaries
