import { NeptunedataClient, GetPropertygraphSummaryCommand } from "@aws-sdk/client-neptunedata";
import {inspect} from "util";
import {NodeHttpHandler} from "@smithy/node-http-handler";

class NeptuneDataClient {
    constructor() {
        const clientConfig = {
            endpoint: 'https://my-cluster-name.cluster-abcdefgh1234.my-region.neptune.amazonaws.com:my-port', // Replace with your endpoint
            sslEnabled: true,
            region: 'my-region', // replace with your region
            maxAttempts: 1,  // do not retry
            requestHandler: new NodeHttpHandler({
                requestTimeout: 0  // no client timeout
            })
        };

        this._client = new NeptunedataClient(clientConfig);
    }
    async getPropertygraphSummary() {
        try {
            const command = new GetPropertygraphSummaryCommand({mode: "basic"});
            return await this._client.send(command);
        } catch (error) {
            console.error("Error executing command.", error)
            throw error;
        }
    }
}
(async () => {
    const client = new NeptuneDataClient();
    const response = await client.getPropertygraphSummary(client);
    console.log(inspect(response.payload, { depth: null}));
})();