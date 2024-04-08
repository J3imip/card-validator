PORT=${PORT:-3000}

docker build -t card-validator .
docker run -it --rm -e PORT="$PORT" -p "$PORT":"$PORT" card-validator