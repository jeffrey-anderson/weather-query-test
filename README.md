# Build

docker build -t weather-query-test:0.1 .

# Run

docker run -it --rm --env ASTRA_DB_ID=$ASTRA_DB_ID --env ASTRA_DB_REGION=$ASTRA_DB_REGION --env ASTRA_DB_KEYSPACE=$ASTRA_DB_KEYSPACE --env ASTRA_DB_APPLICATION_TOKEN=$ASTRA_DB_APPLICATION_TOKEN --name query-test weather-query-test:0.1 query-test [-sampleSize <amount>] [-resolution <amount>]

# k8s

## Create secret for $ASTRA_DB_APPLICATION_TOKEN 

```
kubectl create secret generic astra-db-app-token --from-literal=token='<token value?'
```

## Run in Kubernetes:

```
kubectl apply -f query-test-job.yaml
```