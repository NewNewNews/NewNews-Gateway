#!/bin/sh
# startup.sh

# Wait for database to be ready
/usr/local/bin/wait-for db:5432 -t 60

# Run migrations
# cd /root/prisma
npx prisma migrate deploy

# Start the application
cd /root
exec ./gateway