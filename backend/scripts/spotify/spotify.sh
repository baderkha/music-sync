curl -X POST "https://accounts.spotify.com/api/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=client_credentials&client_id=$SPOTIFY_CLIENT_ID&client_secret=$SPOTIFY_CLIENT_SECRET" | jq -r '.access_token' > tmp_spot_secret

mv tmp_spot_secret secrets






#rm -rf ~/.tmp_spot_secret