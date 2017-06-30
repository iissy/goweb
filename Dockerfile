FROM scratch
COPY ./server /server
EXPOSE 8000/tcp
VOLUME ["/public"]
ENTRYPOINT ["/server"]