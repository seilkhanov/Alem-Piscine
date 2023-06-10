#!/bin/bash
USERNAME="seilkhanov"
RESPONSE=$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data "{\"query\":\"{user(where:{login:{_eq:\\\"$USERNAME\\\"}}){id}}\"}")
ID=$(echo $RESPONSE | jq -r '.data.user[0].id')

echo $ID
