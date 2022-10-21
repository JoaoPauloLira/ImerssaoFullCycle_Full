#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

npm install

npm start


#tem que rodar esse comando aq -> chmod +x .docker/entrypoint.sh