FROM golang:1.24rc1-bookworm AS build

COPY backend/ backend

WORKDIR backend

RUN  go build server.go models.go


from alpine 
COPY --from=build server server
EXPOSE 8080

ENTRYPOINT [ "./server" ]





