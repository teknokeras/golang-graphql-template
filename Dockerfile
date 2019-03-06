FROM golang:1.11.4-stretch

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -

RUN apt-get install -y nodejs

RUN node -v

RUN npm -v

RUN go get github.com/cespare/reflex

RUN mkdir workspace

RUN mkdir workspace/app

RUN mkdir workspace/web

ENV WORKSPACE /workspace

COPY reflex.conf /

WORKDIR $WORKSPACE/web

COPY ./web $WORKSPACE/web

RUN npm install

WORKDIR $WORKSPACE

#linter
RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

#ENTRYPOINT ["reflex", "-c", "/reflex.conf"]

EXPOSE 5000