FROM centos:7

EXPOSE 30668

COPY ./api-server /opt/blog-api-server/
RUN ./api-server
WORKDIR /opt/blog-api-server
