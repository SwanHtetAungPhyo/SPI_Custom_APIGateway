#!/bin/sh

# Start the Go backend in the background
/app/main &

# Change directory to the Next.js frontend build directory
cd /app/ui

# Start the Next.js frontend on port 3000
npx next start -p 3000
