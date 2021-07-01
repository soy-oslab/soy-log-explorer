FROM golang:1.16

COPY explorer explorer

RUN apt -qy update && \
	apt-get -qy \
		-o "Dpkg::Options::=--force-confdef" \
		-o "Dpkg::Options::=--force-confold" upgrade && \
	apt-get -qy autoclean

WORKDIR explorer

RUN go build -v ./...

CMD ["bash"]
