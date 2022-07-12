# module install
FROM node:16-alpine AS modules
RUN apk update
RUN apk add --no-cache python3 g++ make
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package*.json ./
RUN yarn install --production

# build
FROM node:16-alpine AS builder
COPY --from=modules /app/node_modules/ /app/node_modules
WORKDIR /app
COPY . .
RUN npm run build

# serve
FROM node:16-alpine
COPY --from=builder /app/build/ /app/build
RUN yarn global add serve
CMD ["serve", "-s", "/app/build"]

EXPOSE 3000