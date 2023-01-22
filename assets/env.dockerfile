FROM node:14.19.3-alpine3.14

USER root

RUN apk update

# WORKDIR /app

COPY env.sh ./

ENV url=http://rocketchat:3000/
ENV username=user0
ENV password=123456

RUN chmod +x ./env.sh

RUN mkdir app

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global

ENV PATH=$PATH:/home/node/.npm-global/bin

RUN npm --global config set user root && \
    npm --global install @rocket.chat/apps-cli

CMD ["/bin/sh" ,"-c" , "./env.sh watch --url ${url} --username ${username} --password ${password}"]