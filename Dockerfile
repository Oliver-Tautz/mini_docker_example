FROM golang:1.24rc1-bookworm AS build

COPY backend/ backend

WORKDIR backend


# disable CGO for compatible binary in apline ...
ENV CGO_ENABLED=0

RUN  go build -o /bin/server server.go models.go




from alpine

COPY --from=build bin/server /bin/server
EXPOSE 8080

RUN find / -name server
RUN ls -l /bin/server

ENTRYPOINT [ "/bin/server" ]





