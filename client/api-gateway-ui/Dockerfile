FROM node:18 AS frontend-build

WORKDIR /app

# Copy package files to leverage caching
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . ./

# Build the Next.js app
RUN npm run build
FROM node:18-slim AS frontend-final

WORKDIR /app

# Copy the build output from the build stage
COPY --from=frontend-build /app/.next /app/.next
COPY --from=frontend-build /app/public /app/public
COPY --from=frontend-build /app/package*.json ./

# Install production dependencies
RUN npm install --only=production

# Expose the port Next.js runs on
EXPOSE 3000

# Start the Next.js app in production mode
CMD ["npm", "run", "start"]
