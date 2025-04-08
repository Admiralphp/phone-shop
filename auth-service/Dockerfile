FROM node:18-alpine

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

# Expose API port
EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
