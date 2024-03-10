FROM alpine:latest

RUN mkdir /app

COPY ./server-app /app/server-app

CMD [ "/app/server-app" ]