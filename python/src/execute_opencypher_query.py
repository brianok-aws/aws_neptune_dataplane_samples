import boto3
from botocore.config import Config

neptune_endpoint = "https://my-cluster-name.cluster-abcdefgh1234.my-region.neptune.amazonaws.com:my-port" # Replace with your Neptune endpoint

my_config = Config(
    region_name = 'my-region',
    retries = {
        'max_attempts': 1
    },
    read_timeout=None
)
client = boto3.client("neptunedata", config=my_config, endpoint_url=neptune_endpoint)

query = "MATCH (n) RETURN n LIMIT 10" # Example openCypher query
response = client.execute_open_cypher_query(openCypherQuery=query)

for item in response['results']:
    print(item)
