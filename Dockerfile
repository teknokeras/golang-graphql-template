FROM golang:1.11.4-stretch

RUN go get github.com/cespare/reflex

ENV APP /app

COPY reflex.conf /

WORKDIR $APP


#ENTRYPOINT ["reflex", "-c", "/reflex.conf"]

EXPOSE 5000