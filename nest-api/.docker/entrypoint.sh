#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

npm install

npm run start:dev


#tem que rodar esse comando aq -> chmod +x .docker/entrypoint.sh