FROM zaha/python
MAINTAINER Tony Worm docker@zaha.io


# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]


# Update stuff
RUN apt-get update 

# Install Python Setuptools
RUN apt-get --no-install-recommends install -y \
	gfortran \
	libffi-dev \
	libhdf5-dev \
	liblapack3 \
	liblapack-dev \
	libblas3 \
	libblas-dev \
	libpng3 \
	libxslt1-dev \
	libxml2-dev \
	zlib1g-dev \
	graphviz \
	python-matplotlib 


ADD requirements.txt /tmp/requirements.txt
RUN pip install -r /tmp/requirements.txt  

# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
