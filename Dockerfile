FROM centos:7

WORKDIR /root

RUN yum install -y wget && \
    wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz && \
    mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg && \
    mkdir go/src/dumb

ENV PATH=${PATH}:/usr/local/go/bin GOROOT=/usr/local/go GOPATH=/root/go

ADD main.go go/src/dumb/
ADD client.go go/src/dumb/
ADD exmo.go go/src/dumb/
ADD bittrex.go go/src/dumb/

RUN go build dumb && go install dumb

RUN cd ~
RUN mkdir public
RUN mkdir static
ADD public/* public/
ADD static/* static/

RUN chmod 777 -R ~/public/
RUN chmod 777 -R ~/static/

EXPOSE 80

CMD ./go/bin/dumb