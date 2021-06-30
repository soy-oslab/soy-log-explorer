FROM golang:1.16

COPY soy_log_explorer soy_log_explorer

RUN apt -qy update && \
	apt-get -qy \
		-o "Dpkg::Options::=--force-confdef" \
		-o "Dpkg::Options::=--force-confold" upgrade && \
	apt-get -qy autoclean

WORKDIR soy_log_explorer

RUN go build -v ./...

CMD ["bash"]
