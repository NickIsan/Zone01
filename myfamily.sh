

curl -s "https://platform.zone01.gr/assets/superhero/all.json" \
    | jq -r --argjson  id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives' \
    | tr -d '"'