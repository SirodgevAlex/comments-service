#!/bin/bash

cd ..
cd db
cd migrations

psql -U postgres -d test-comments-system -f 006_delete_all_comments.sql
psql -U postgres -d test-comments-system -f 005_delete_all_posts.sql
# psql -U postgres -d test-comments-system -f 007_drop_test-comments-service.sql
# psql -U postgres -d test-comments-system -f 008_create_test-comments-service.sql

cd ../..
cd tests
go test

#password 1234 