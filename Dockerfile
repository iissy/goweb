FROM alpine
COPY ./server /server
EXPOSE 8000/tcp
VOLUME ["/public", "/views"]
ENTRYPOINT ["/server"]