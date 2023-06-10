#!/bin/bash

name=$(curl -s https://01.alem.school/assets/superhero/all.json | jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives')

echo "$name" | tr -d '""'
