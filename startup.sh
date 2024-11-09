#!/bin/sh
# startup.sh

# Wait for database to be ready
/usr/local/bin/wait-for db:5432 -t 60

# Run migrations
cd /root/prisma
go run github.com/steebchen/prisma-client-go migrate deploy

# Start the application
cd /root
exec ./gateway