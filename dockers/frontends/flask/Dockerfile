FROM zaha/pyserv
MAINTAINER Tony Worm docker@zaha.io

# Update stuff
RUN apt-get update

# Install Python Setuptools
RUN apt-get --no-install-recommends install -y \
	ruby ruby-dev 
RUN curl -sL https://deb.nodesource.com/setup | bash -
RUN apt-get --no-install-recommends install -y nodejs

# RUN gem install sass
RUN gem install compass
RUN npm install bower uglify-js -g

# Clean up
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# REPLACE daemon to be run by runit.
RUN rm /etc/service/flask/run && ln -s /scripts/launch.sh /etc/service/flask/run

RUN pip install -r /scripts/requirements.txt  

EXPOSE 5000

VOLUME ["/src"]
WORKDIR /src

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
