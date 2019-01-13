FROM golang:1.11.4-stretch

RUN go get github.com/cespare/reflex

ENV APP /app

COPY reflex.conf /

RUN mkdir $APP

WORKDIR $APP

#ENTRYPOINT ["reflex", "-c", "/reflex.conf"]