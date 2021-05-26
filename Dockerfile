# image for building application
FROM golang:1.15-alpine as builder
# install all dependencies
RUN apk add --no-cache git make gcc g++ musl-dev linux-headers gmp-dev mpfr-dev
ADD . /bitbon-system-blockchain
RUN cd /bitbon-system-blockchain && make clean build

# image for running application
FROM alpine:latest
RUN apk add --no-cache libstdc++ gmp-dev libgomp
COPY --from=builder /bitbon-system-blockchain/build/bin/bitbon_node /usr/local/bin/
RUN ln -s /usr/local/bin/bitbon_node /bitbon_node
EXPOSE 8545 8546 30303 30303/udp 6060
ENTRYPOINT ["/bitbon_node"]
