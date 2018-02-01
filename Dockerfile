FROM golang

ENV PATH $PATH:$PATH/bin

COPY . src/github.com/yulibaozi/yulibaozi.com/

# ADD run.sh /

WORKDIR /go/src/github.com/yulibaozi/yulibaozi.com

RUN chmod 777 ./run.sh

EXPOSE 8081


CMD ["./run.sh"]
