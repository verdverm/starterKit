FROM zaha/base
MAINTAINER Tony Worm docker@zaha.io

# Install instructions from https://cwiki.apache.org/confluence/display/COUCHDB/Debian

ENV COUCHDB_VERSION 1.6.0

# Define mountable directories.
VOLUME ["/couchdb/var/lib/couchdb", "/couchdb/var/log/couchdb"]

# Expose for linking with other containers
EXPOSE 5984

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]


# Relax...
ADD build /build/docker-couchdb

RUN /build/docker-couchdb/prepare.sh
RUN /build/docker-couchdb/add_erlang_apt.sh
RUN /build/docker-couchdb/install_couchdb.sh

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
