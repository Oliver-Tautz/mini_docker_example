FROM golang:1.24rc1-bookworm AS build

COPY backend/ backend

WORKDIR backend

RUN  go build -o /bin/server server.go models.go


from alpine 
COPY --from=build bin/server /bin/server
EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]





