#! /bin/bash/

curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Adsouza\"}}){id}}"}' | cut -d "}" -f 1 | cut -b 24-27