# Admin Dashboard Frontend Dockerfile
# ./admin-dashboard/Dockerfile
FROM node:18-alpine as build

WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy source code
COPY . .

# Build the Angular application
RUN npm run build --prod

# Production stage
FROM nginx:alpine

# Copy built assets from the build stage
COPY --from=build /app/dist/admin-dashboard /usr/share/nginx/html

# Copy custom nginx config
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]