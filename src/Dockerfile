FROM golang:1.9 AS builder

RUN go version

COPY . "/go/src/github.com/danielfoehrKn/k8GoAnalyzer"
WORKDIR "/go/src/github.com/danielfoehrKn/k8GoAnalyzer"

RUN go get -d -v github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux .
#RUN go get -v -t  .
#RUN set -x && \
#    #go get github.com/2tvenom/go-test-teamcity && \
#    go get -u github.com/golang/dep/cmd/dep
#
#RUN dep init

#RUN  dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /goanalyzer
RUN ls -l

CMD ["/goanalyzer"]

EXPOSE 8080


#########
# second stage to obtain a very small image
FROM scratch

COPY --from=builder /goanalyzer .
COPY --from=0 /go/src/github.com/danielfoehrKn/k8GoAnalyzer/config.toml .
#COPY --from=builder config.toml .

EXPOSE 8080

CMD ["/goanalyzer"]