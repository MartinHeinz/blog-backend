FROM node:lts-alpine as builder
COPY package*.json ./
RUN npm install --no-optional
COPY . .
RUN npm run build

FROM nginx:stable-alpine as runner
COPY --from=builder dist /home/html
COPY site.conf /etc/nginx/conf.d/
COPY .htpasswd /etc/nginx/.htpasswd
EXPOSE 8080
COPY startup.sh /home/
RUN chmod 777 /home/startup.sh
CMD ["sh", "/home/startup.sh"]

# docker build -t blog-frontend .
# docker run -p 8080:8080 --rm blog-frontend