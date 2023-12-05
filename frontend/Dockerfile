# Use an official Python runtime as the base image
FROM node:20.10.0

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

RUN npm install

# Copy the rest of the application code to the container
COPY . .

RUN npm run build

