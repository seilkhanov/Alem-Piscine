#!/bin/bash
SUBJECT_ID=70
NAME=$(curl -s https://01.alem.school/assets/superhero/all.json | jq --arg id "$SUBJECT_ID" -r '.[] | select(.id == ($id | tonumber)) | .name')

echo '"'"$NAME"'"'
