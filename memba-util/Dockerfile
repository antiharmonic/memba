FROM golang:1.21

WORKDIR /usr/src/app

# install cron
RUN apt install cron

# pre-copy/cache go.mod
COPY go.mod go.sum ./
RUN  go mod download && go mod verify

COPY db-sync-cron /etc/cron.d/db-sync-cron
RUN  chmod 0644   /etc/cron.d/db-sync-cron

COPY config.yml db-sync.go ./
RUN go build -v -o /usr/local/bin/db-sync ./...

CMD ["cron", "-f"]
