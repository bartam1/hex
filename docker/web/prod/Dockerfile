FROM node:15.14.0-alpine3.12

ENV NODE_ENV development

RUN apk add yarn
ADD start.sh /
RUN chmod +x /start.sh

CMD ["/start.sh"]