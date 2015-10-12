#!/bin/bash

set -e

export APPHOST=$(netstat -nr | grep 'UG' | awk '{print $2}')

echo "SERVICE      HOST"
echo "apphost      $APPHOST"
echo "app_cnt      $APPCNT"

echo "APPDOMAIN    $APPDOMAIN"
echo "RUNMODE      $RUNMODE"

(
	echo "# Auto-generated : do NOT touch"
	echo ""
		echo "upstream ipython.localhost {"
		printf "    server ${APPHOST}:9999;\n"
		echo "}"
		echo ""
		echo "upstream datasrvr.localhost {"
		echo "    least_conn;"
		printf "    server ${APPHOST}:5223;\n"
		echo "}"
		echo ""
	if [ $APPCNT -gt 0 ]; then
		echo "upstream landing.localhost {"
		echo "    least_conn;"
		for i in $(seq 1 $APPCNT); do
			printf "    server ${APPHOST}:500${i};\n"
		done
		echo "}"
		echo ""
	fi
		echo "server {"
		echo "    listen 80;"
	if [ "$RUNMODE" == "prod" ]; then
		echo "    return 301 https://\$host\$request_uri;"
		echo "}"
		echo ""
		echo "server {"
		echo "    ssl on;"
		echo "    ssl_certificate     /etc/ssl/$APPDOMAIN.pem;"
		echo "    ssl_certificate_key /etc/ssl/$APPDOMAIN.key;"
		echo ""
		echo "    listen 443;"
	fi
		echo "    server_name $APPDOMAIN www.$APPDOMAIN;"
		echo ""
		echo "    access_log /var/log/nginx/$APPDOMAIN.access.log;"
        echo "    error_log /var/log/nginx/$APPDOMAIN.error.log;"
		echo ""
		echo "    client_max_body_size 100M;"
		echo ""
		echo "    location /static {"
		echo "        root /;"
		echo "    }"
		echo ""    
		echo "    location /ipython {"
		echo "        proxy_pass http://ipython.localhost;"
		echo "        # Needed for the websockets connections:"
		echo "        proxy_http_version 1.1;"
		echo "        proxy_set_header Upgrade \$http_upgrade;"
		echo "        proxy_set_header Connection "upgrade";"
		echo "        proxy_set_header Origin \"\";"
		echo "    }"
		echo ""
		echo "    location /datasrvr {"
		echo "        proxy_redirect          off;"
		echo "        proxy_set_header        Host            \$host;"
		echo "        proxy_set_header        X-Real-IP       \$remote_addr;"
		echo "        proxy_set_header        X-Forwarded-For \$proxy_add_x_forwarded_for;"
		echo "        proxy_pass http://datasrvr.localhost;"
		echo "    }"
		echo ""
		echo "    location / {"
		echo "        proxy_redirect          off;"
		echo "        proxy_set_header        Host            \$host;"
		echo "        proxy_set_header        X-Real-IP       \$remote_addr;"
		echo "        proxy_set_header        X-Forwarded-For \$proxy_add_x_forwarded_for;"
		echo "        proxy_pass http://landing.localhost;"
		echo "    }"
		echo "}"
		echo ""
		echo ""


) > $APPDOMAIN.cfg

sudo cp $APPDOMAIN.cfg /etc/nginx/sites-enabled/default

less $APPDOMAIN.cfg

nginx
