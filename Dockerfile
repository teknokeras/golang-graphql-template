FROM golang:1.11.4-stretch

RUN go get github.com/cespare/reflex

ENV APP /app

COPY reflex.conf /

WORKDIR $APP

#linter
RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

#ENTRYPOINT ["reflex", "-c", "/reflex.conf"]

EXPOSE 5000