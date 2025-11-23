host="$DB_HOST"
port="$DB_PORT"

echo "Waiting for PostgreSQL at $host:$port to be ready..."

while ! nc -z $host $port; do
  sleep 0.5
done

echo "PostgreSQL is up and running!"

exec "$@"