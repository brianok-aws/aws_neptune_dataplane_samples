package main
import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata"
	"os"
	"encoding/json"
	"net/http"
)

func main() {
    region := "my-region"
	clusterEndpoint := "my-cluster-name.cluster-abcdefgh1234." + region + ".neptune.amazonaws.com"
	neptunePort := "neptune-port"
	// Here we set an unlimited client timeout, but 
	// you can also use the value of your instance timeout from the Neptune configuration
	client := &http.Client{
		Timeout: 0,
	}
    sdkConfig, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithHTTPClient(client))
    svc := neptunedata.NewFromConfig(sdkConfig, func(o *neptunedata.Options) {
        o.BaseEndpoint = aws.String("https://" + clusterEndpoint + ":" + neptunePort)
		o.Retryer = aws.NopRetryer{}    // Do not retry calls if they fail
    })
    query := "g.addV('person').property('name','justin').property(id,'justin-1')"
    serializer := "application/vnd.gremlin-v1.0+json;types=false"
    input := neptunedata.ExecuteGremlinQueryInput{GremlinQuery: &query, Serializer: &serializer}
    result, err1 := svc.ExecuteGremlinQuery(context.TODO(), &input)
	if (err1 != nil) {
		fmt.Printf("Error retrieving result, %v", err1.Error())
		os.Exit(1)
	}
    var kv map[string]interface{}
	err2 := result.Result.UnmarshalSmithyDocument(&kv)
	if err2 != nil {
		fmt.Printf("Error retrieving result, %v", err2.Error())
		os.Exit(1)
	}
    enc, _ := json.Marshal(kv)
    fmt.Println(string(enc))
    os.Exit(0)
}
