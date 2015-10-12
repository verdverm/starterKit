#!/bin/bash
set -e

if [ ! -f /scripts.rabbitmq_password_set ]; then
	/scripts/set_rabbitmq_password.sh
fi

if [ "$1" = 'rabbitmq-server' ]; then
	chown -R rabbitmq /var/lib/rabbitmq
fi

exec "$@"