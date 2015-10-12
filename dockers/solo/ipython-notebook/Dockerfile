FROM zaha/python-ml
MAINTAINER Tony Worm docker@zaha.io

# Install Nginx.
RUN \
  add-apt-repository -y ppa:nginx/stable && \
  apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y nginx && \
  chown -R www-data:www-data /var/lib/nginx

# Clean up
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# INSTALL requirements
RUN pip install -r /scripts/requirements.txt  


# REPLACE daemon to be run by runit.
RUN mkdir -p /etc/service/ipython-notebook
RUN ln -s /scripts/launch.sh /etc/service/ipython-notebook/run

# ENABLE nginx site configuration
RUN mv /scripts/ipython.conf /etc/nginx/sites-enabled/ipython.conf
RUN chown :www-data /etc/nginx/sites-enabled/ipython.conf

# Expose ports... for the outside world...
EXPOSE 8080

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
