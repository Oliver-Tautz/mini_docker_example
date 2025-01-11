from golang:1.24rc1-bookworm AS build

ADD backend .

WORKDIR backend

EXPOSE 8080
RUN  build server.go models.go


from alpine 
COPY --from=build backend/server .
EXPOSE 8080
ENTRYPOINT [ "./server" ]





