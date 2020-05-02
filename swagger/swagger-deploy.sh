#!/bin/sh

# swaggerをredocに対応するように変換
redoc-cli bundle ./swagger.yml --output ./public/index.html

# firebaseにデプロイ
firebase deploy
