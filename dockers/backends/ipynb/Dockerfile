FROM zaha/pyserv-ml
MAINTAINER Tony Worm docker@zaha.io

# Update stuff
RUN apt-get update && apt-get --no-install-recommends install -y expect
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# REPLACE daemon to be run by runit.
RUN rm /etc/service/flask/run && ln -s /scripts/launch.sh /etc/service/flask/run

RUN pip install -r /scripts/requirements.txt  

# Some preliminary ipython tasks
ENV IPYTHONDIR /ipython

EXPOSE 8888 9999

VOLUME ["/ipynb/notebooks", "/ipython"]

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
