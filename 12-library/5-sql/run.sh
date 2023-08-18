#!/bin/bash

PASSWORD=lab_password

docker run -d \
  --network host \
	--name some-postgres \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=${PASSWORD} \
	postgres:14

sleep 5

docker exec -e POSTGRES_PASSWORD=${PASSWORD} -i some-postgres psql -U postgres << EOF
  CREATE TABLE users (
      id serial PRIMARY KEY,
      email VARCHAR (355) UNIQUE NOT NULL,
      password VARCHAR (50) NOT NULL
  );

  insert into users(id,email,password) values(1,'1056764180@qq,com','12345678');
  insert into users(id,email,password) values(2,'10567@qq,com','1234567890');
  insert into users(id,email,password) values(3,'10567567@qq,com','12345678908');
EOF