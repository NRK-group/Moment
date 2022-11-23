FROM node:16.17.0

WORKDIR /App-frontend

COPY package*.json ./
RUN npm install
COPY . .

CMD [ "npm", "run", "produce" ]