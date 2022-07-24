# # # # # # #
#  Builder  #
# # # # # # #

FROM golang:1.18.4-alpine3.16 as builder
# TO-DO: Use alpine and install sh package

# Make an /app directory in the container
RUN mkdir /app

# change default working directory in the container to /app
WORKDIR /app

# Cache dependencies
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Hot reloading with Air
RUN apk --no-cache add curl
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# copy all files in the root directory to /app in the container
COPY ./ ./
ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache

# Build the binary, call it main, put in the root directory in the container
RUN go build -o main /app/
RUN ls

EXPOSE 8080

# CMD [ "air" ]


# # # # # # #
#  Runner   #
# # # # # # #

FROM alpine:3.16 as runner
# WORKDIR /
COPY --from=builder /app/main /main
# CMD ./main


ENTRYPOINT [ "/main" ]
# CMD [ "./main" ]
# ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
# RUN pwd
# RUN ls
# RUN ls
# USER nonroot:nonroot
# CMD [ "./main" ]
# CMD ["./app/main"]
