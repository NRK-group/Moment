
FROM node:16.17.0 AS builder

WORKDIR /App

COPY package*.json ./
RUN npm install
COPY . ./

RUN npm run build


FROM nginx:1.19.0
#copies React to the container directory
# Set working directory to nginx resources directory
WORKDIR /usr/share/nginx/html
# Remove default nginx static resources
RUN rm -rf ./*
# Copies static resources from builder stage
COPY --from=builder /App/dist .
ADD nginx.conf /etc/nginx/nginx.conf
# Containers run nginx with global directives and daemon off
EXPOSE 8070
ENTRYPOINT ["nginx", "-g", "daemon off;"]

