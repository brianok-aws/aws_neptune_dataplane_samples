import { NeptunedataClient, GetPropertygraphSummaryCommand } from "@aws-sdk/client-neptunedata";
import {inspect} from "util";

class NeptuneDataClient {
    constructor() {
        const clientConfig = {
            endpoint: 'https://my-cluster-name.cluster-abcdefgh1234.my-region.neptune.amazonaws.com:my-port',
            sslEnabled: true,
            region: 'my-region',
            maxAttempts: 1
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