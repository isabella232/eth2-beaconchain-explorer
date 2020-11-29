printf 'start'
docker container exec -i $(docker-compose ps -q postgres) psql -U debug -d beaconcha < tables.sql