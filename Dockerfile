# Stage 1: Build the Next.js frontend
FROM node:18 AS frontend-build

WORKDIR /app

# Copy only the package files first for better caching
COPY client/api-gateway-ui/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY client/api-gateway-ui/ ./

# Build the Next.js app
RUN npm run build

# Stage 2: Serve the Next.js app
FROM node:18-slim AS frontend-final

WORKDIR /app

# Copy the build output from the build stage
COPY --from=frontend-build /app/.next /app/.next
COPY --from=frontend-build /app/public /app/public
COPY --from=frontend-build /app/package*.json ./

# Install only production dependencies
RUN npm install --production

# Expose the port Next.js runs on
EXPOSE 3000

# Start the Next.js app
CMD ["npx", "next", "start", "-p", "3000"]
