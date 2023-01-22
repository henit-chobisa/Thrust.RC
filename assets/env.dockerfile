FROM node:14.19.3-alpine3.14

USER root

RUN apk update

WORKDIR /app

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global

ENV PATH=$PATH:/home/node/.npm-global/bin

RUN npm --global config set user root && \
    npm --global install @rocket.chat/apps-cli

ENTRYPOINT [ "rc-apps" ]