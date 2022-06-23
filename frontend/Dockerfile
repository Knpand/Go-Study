FROM node:14.17.1
WORKDIR /app
RUN rm -rf node_modules package-lock.json && npm install
RUN yarn global add @vue/cli
RUN yarn global add axios
CMD ["yarn","add","@vue/cli"]
CMD ["yarn","serve"]
