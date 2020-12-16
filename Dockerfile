FROM alpine
COPY ./server /server
EXPOSE 80/tcp
VOLUME ["/public", "/views"]
ENTRYPOINT ["/server"]