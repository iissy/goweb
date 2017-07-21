FROM scratch
COPY ./server /server
EXPOSE 80/tcp
VOLUME ["/public"]
ENTRYPOINT ["/server"]