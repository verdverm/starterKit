FROM zaha/base
MAINTAINER Tony Worm docker@zaha.io


# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]


# Update stuff
RUN apt-get update 

# Install Python Setuptools
RUN apt-get --no-install-recommends install -y \
	python-setuptools python-dev python-pip \
	libpq-dev 

ADD requirements.txt /tmp/requirements.txt
RUN pip install -r /tmp/requirements.txt  


# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
