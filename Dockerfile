FROM golang:1.11.4 AS build-env
WORKDIR /usr/local/go/src/github.com/SpectoLabs/hoverfly
COPY . /usr/local/go/src/github.com/SpectoLabs/hoverfly    
RUN cd core/cmd/hoverfly && CGO_ENABLED=0 GOOS=linux go install -ldflags "-s -w"
RUN cd hoverctl && CGO_ENABLED=0 GOOS=linux go install -ldflags "-s -w"
RUN cd ctlserver && CGO_ENABLED=0 GOOS=linux go install -ldflags "-s -w"

FROM mysql:5.7 AS mysql

FROM debian:stretch-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl && rm -rf /var/lib/apt/lists/*

ENV MYSQL_EXECUTABLE=mysql
ENV BINLOG_EXECUTABLE=mysqlbinlog
ENV MYSQLDUMP_EXECUTABLE=mysqldump
ENV MYSQLADMIN_EXECUTABLE=mysqladmin
ENV DESTINATIONDB_HOST=poa-db
ENV DESTINATIONDB_USER=root
ENV DESTINATIONDB_PASSWORD=root
COPY --from=build-env /usr/local/go/bin/hoverfly /bin/hoverfly
COPY --from=build-env /usr/local/go/bin/hoverctl /bin/hoverctl
COPY --from=build-env /usr/local/go/bin/ctlserver /bin/ctlserver
COPY --from=mysql /usr/bin/mysql /usr/bin/mysql
COPY --from=mysql /usr/bin/mysqldump /usr/bin/mysqldump
COPY --from=mysql /usr/bin/mysqlbinlog /usr/bin/mysqlbinlog
COPY --from=mysql /usr/bin/mysqladmin /usr/bin/mysqladmin
COPY --from=mysql /usr/lib/x86_64-linux-gnu/libatomic.so.1 /usr/lib/x86_64-linux-gnu/libatomic.so.1
COPY --from=mysql /lib/x86_64-linux-gnu/libncurses.so.5 /lib/x86_64-linux-gnu/libncurses.so.5

CMD ["/bin/ctlserver"]

EXPOSE 8500 8888
