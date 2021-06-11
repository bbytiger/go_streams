FROM ubuntu:18.04
RUN mkdir /usr/share/man/man1
RUN apt update -y
RUN apt install wget -y
RUN apt install default-jre -y
RUN apt install emacs -y
RUN mkdir /usr/local/downloads && cd /usr/local/downloads
RUN wget http://www-us.apache.org/dist/kafka/2.7.0/kafka_2.13-2.7.0.tgz
RUN tar xzf kafka_2.13-2.7.0.tgz
RUN mv kafka_2.13-2.7.0 /usr/local/kafka