FROM golang:1.16

COPY explorer explorer

RUN apt-get -qy \
	-o "Dpkg::Options::=--force-confdef" \
	-o "Dpkg::Options::=--force-confold" upgrade && \
	apt-get -qy autoclean

WORKDIR /go/explorer

RUN go build -v ./...

CMD ["bash"]
