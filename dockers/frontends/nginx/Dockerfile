FROM zaha/base
MAINTAINER Tony Worm docker@zaha.io

# Install Nginx.
RUN \
  add-apt-repository -y ppa:nginx/stable && \
  apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y nginx && \
  echo "\ndaemon off;\n" >> /etc/nginx/nginx.conf && \
  chown -R www-data:www-data /var/lib/nginx

# Clean up
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Define mountable directories.
RUN mkdir -p /static /data && chown -R www-data:www-data /static /data
VOLUME ["/static", "/data", "/etc/nginx/sites-enabled", "/var/log/nginx", "/var/www/cache", "/etc/ssl"]

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# Add daemon to be run by runit.
RUN mkdir /etc/service/nginx
RUN ln -s /scripts/launch.sh /etc/service/nginx/run

# Expose ports... for the outside world...
EXPOSE 80 443

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
