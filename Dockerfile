from golang:1.24rc1-bookworm AS build

COPY backend/ .

WORKDIR backend

RUN  go build server.go models.go


from alpine 
COPY --from=build backend/server .
EXPOSE 8080

ENTRYPOINT [ "./server" ]





