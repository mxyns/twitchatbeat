FROM golang:1.17.13

RUN \
    apt-get update \
      && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
         netcat \
         python3 \
         python3-pip \
         python3-venv \
      && rm -rf /var/lib/apt/lists/*

RUN pip3 install --upgrade pip==20.1.1
RUN pip3 install --upgrade setuptools==47.3.2
RUN pip3 install --upgrade docker-compose==1.23.2

# Setup work environment
ENV GOPATH="/usr/share/go"
ENV PATH="${PATH}:${GOPATH}/bin"
ENV TWITCHATBEAT_REPO github.com/mxyns/twitchatbeat
ENV TWITCHATBEAT_INSTALLDIR /usr/share/twitchatbeat


RUN mkdir -p ${GOPATH}/src/github.com/magefile && cd $_ \
    && git clone https://github.com/magefile/mage \
    && cd mage \
    && go run bootstrap.go

WORKDIR $GOPATH/src/$TWITCHATBEAT_REPO
COPY . .
RUN make update; exit 0
RUN mage update
RUN mage build

RUN mkdir -p $TWITCHATBEAT_INSTALLDIR
RUN mv twitchatbeat $TWITCHATBEAT_INSTALLDIR/

WORKDIR $TWITCHATBEAT_INSTALLDIR
RUN cd $TWITCHATBEAT_INSTALLDIR
RUN export CWD=$(pwd)
ENV PATH="${PATH}:${CWD}"
