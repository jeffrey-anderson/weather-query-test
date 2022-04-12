# Build

docker build -t query-test:0.1 .

# Run

docker run -it --rm --env ASTRA_DB_ID=$ASTRA_DB_ID --env ASTRA_DB_REGION=$ASTRA_DB_REGION --env ASTRA_DB_KEYSPACE=$ASTRA_DB_KEYSPACE --env ASTRA_DB_APPLICATION_TOKEN=$ASTRA_DB_APPLICATION_TOKEN --name query-test query-test:0.1

