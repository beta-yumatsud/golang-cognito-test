FROM alpine:3.12

RUN apk update \
    && apk upgrade \
    && apk add --update-cache --no-cache --virtual .create-user shadow \
    && groupadd -g 1000 golang-cognito-test \
    && useradd -u 1000 -g 1000 golang-cognito-test \
    && apk del .create-user \
    && apk add tzdata ca-certificates \
    && rm -rf /var/cache/apk/*

ENV TZ=Asia/Tokyo

USER golang-cognito-test
WORKDIR /app
COPY ./.artifacts/golang-cognito-test-linux-amd64 .

CMD ["./golang-cognito-test-linux-amd64"]
