#Author young_zkc@163.com
#VERSION 1.1
FROM ubuntu:15.10
RUN groupadd -r younger && useradd	-r -g	younger younger
# USER younger
RUN  mkdir -p /home/youmger/work && mkdir -p /home/younger/work/conf
COPY main /home/younger/work
COPY   ./conf/conf.yml /home/younger/work/conf/conf.yml
WORKDIR /home/younger/work
RUN chmod +x main
EXPOSE 8080
CMD [ "main" ]
