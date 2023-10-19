FROM node:16

WORKDIR /app

RUN yarn init -y

RUN yarn add -D @babel/cli @babel/core @babel/preset-env babel-loader webpack webpack-cli webpack-dev-server html-webpack-plugin vue-loader vue-template-compiler css-loader vue-style-loader sass-loader

RUN yarn add babel-polyfill vue node-sass axios

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

CMD ["npm", "run", "start"]