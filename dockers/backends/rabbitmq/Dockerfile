FROM zaha/base
MAINTAINER Tony Worm  docker@zaha.io

ENV RABBITMQ_VERSION 3.4.3-1

# Install RabbitMQ
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys F7B8CEA6056E8E56 && \
    echo "deb http://www.rabbitmq.com/debian/ testing main" >> /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y rabbitmq-server pwgen && \
    rabbitmq-plugins enable rabbitmq_management 

# Clean up
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


# Define environment variables.
ENV RABBITMQ_LOG_BASE /data/log
ENV RABBITMQ_MNESIA_BASE /data/mnesia

# Define mount points.
RUN mkdir -p /data/log data/mnesia
RUN chown -R rabbitmq:rabbitmq /data
VOLUME ["/data/log", "/data/mnesia"]

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# Add daemon to be run by runit.
RUN mkdir /etc/service/rabbitmq
RUN ln -s /scripts/launch.sh /etc/service/rabbitmq/run

# Expose for linking with other containers
EXPOSE 5672 15672

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
