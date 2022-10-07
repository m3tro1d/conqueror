FROM debian:10-slim

ADD bin/conqueror /app/bin/conqueror

WORKDIR /app/bin

EXPOSE 8080

CMD [ "/app/bin/conqueror" ]
