# Use an official Python runtime as the base image
FROM node:20.10.0 as node

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

RUN npm install

# Copy the rest of the application code to the container
COPY . .

RUN npm run build


FROM nginx:latest

COPY --from=node /app/dist/browser /etc/nginx/html

COPY --from=node /app/nginx-pages/ /etc/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000

