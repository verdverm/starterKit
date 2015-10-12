#!/bin/bash
set -e
source /build/docker-couchdb/buildconfig
set -x

$minimal_apt_get_install erlang-base-hipe erlang-crypto erlang-eunit \
  erlang-inets erlang-os-mon erlang-public-key erlang-ssl \
  erlang-syntax-tools erlang-tools erlang-xmerl erlang-dev libicu-dev \
  libmozjs185-dev erlang-asn1 make g++ libtool pkg-config git \
  automake autoconf autoconf-archive

git clone https://github.com/apache/couchdb.git /tmp/couchdb
cd /tmp/couchdb
git checkout tags/1.6.1

./bootstrap
./configure --prefix=/couchdb && make && make install

useradd -d /couchdb/lib/couchdb couchdb

mkdir /etc/service/couchdb
cp /build/docker-couchdb/runit/couchdb.sh /etc/service/couchdb/run
cp /build/docker-couchdb/50_enforce_couchdb_permissions.sh /etc/my_init.d/



# copy local.ini, overriding default settings

# Should do some sed replacements for auth related things
cp /build/docker-couchdb/local.ini /couchdb/etc/couchdb/local.ini
