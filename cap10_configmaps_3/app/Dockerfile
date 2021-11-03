FROM golang:1.14-alpine as build

WORKDIR /src/
COPY main.go /src/
RUN CG0_ENABLED=0 && go build -o /bin/demo

FROM alpine
# COPY textfile.txt /app/
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT [ "/bin/demo" ]