#!/bin/bash

#  **** ~> not ready yet, please PR

# nocache="--no-cache"


###################
###   BASES     ###
###################

# docker pull phusion/baseimage:0.9.15
# docker build $nocache -t zaha/base       bases/base
# docker build $nocache -t zaha/java8      bases/java8
# docker build $nocache -t zaha/python     bases/python
# docker build $nocache -t zaha/python-ml  bases/python-ml
# docker build $nocache -t zaha/golang     bases/golang


###################
###  DATABASES  ###
###################

# docker build $nocache -t zaha/psql       databases/postgresql
# docker build $nocache -t zaha/neo4j      databases/neo4j
# docker build $nocache -t zaha/couchdb    databases/couchdb
# docker build $nocache -t zaha/mongodb        databases/mongodb    ****


###################
###  SERVICES   ###
###################

# docker build $nocache -t zaha/goserv     services/goserv
# docker build $nocache -t zaha/pyserv     services/pyserv
# docker build $nocache -t zaha/pyserv-ml  services/pyserv-ml


###################
###   BACKEND   ###
###################

# docker build $nocache -t zaha/ipynb      backends/ipynb
# docker build $nocache -t zaha/rabbitmq   backends/rabbitmq


###################
###  FRONTEND   ###
###################

# docker build $nocache -t zaha/nginx      frontends/nginx
# docker build $nocache -t zaha/flask      frontends/flask
# docker build $nocache -t zaha/redis      frontends/redis
# docker build $nocache -t zaha/memcache       frontends/memcache   ****


###################
###    APPS     ###
###################

# docker build $nocache -t zaha/pyapp          apps/python         **** 
# docker build $nocache -t zaha/goapp          apps/golang         ****
