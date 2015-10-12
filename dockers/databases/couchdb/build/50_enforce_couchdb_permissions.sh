#!/bin/bash
# Docker attaches shared directories as owned by root
set -e
chown -R couchdb:couchdb /couchdb/etc/couchdb /couchdb/var/lib/couchdb /couchdb/var/log/couchdb /couchdb/var/run/couchdb
chmod 0770 /couchdb/etc/couchdb /couchdb/var/lib/couchdb /couchdb/var/log/couchdb /couchdb/var/run/couchdb
