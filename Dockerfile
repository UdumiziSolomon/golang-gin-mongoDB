FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/UdumiziSolomon/Gopher-Backend
RUN cd /build &&  git clone https://github.com/<REPO URL>

RUN cd /build/Gopher-Backend && go build

EXPOSE 6000

ENTRYPOINT ["/build/Gopher-Backend/main"]