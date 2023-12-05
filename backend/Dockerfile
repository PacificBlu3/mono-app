# Use an official Python runtime as the base image
FROM python:3.11.4

# Set the working directory in the container
WORKDIR /app

# Copy the requirements.txt file to the container
COPY requirements.txt .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code to the container
COPY . .

EXPOSE 3300

# Set the default command to run when the container starts
CMD [ "python", "app.py" ]

# Add labels for better maintainability
# LABEL maintainer="Your Name <your.email@example.com>"
# LABEL version="1.0"
# LABEL description="Dockerfile for Python 3 and Node.js"