package main
import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata"
	"os"
	"encoding/json"
)

func main() {
    region := "my-region"
	clusterEndpoint := "my-cluster-name.cluster-abcdefgh1234.my-region.neptune.amazonaws.com"
	neptunePort := "neptune-port"
    sdkConfig, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
    svc := neptunedata.NewFromConfig(sdkConfig, func(o *neptunedata.Options) {
        o.BaseEndpoint = aws.String("https://" + clusterEndpoint + ":" + neptunePort)
		o.Retryer = aws.NopRetryer{}
    })
    query := "g.addV('person').property('name','justin').property(id,'justin-1')"
    serializer := "application/vnd.gremlin-v1.0+json;types=false"
    input := neptunedata.ExecuteGremlinQueryInput{GremlinQuery: &query, Serializer: &serializer}
    result, err1 := svc.ExecuteGremlinQuery(context.TODO(), &input)
	if (err1 != nil) {
		fmt.Printf("Error retrieving result, %v", err1.Error())
		os.Exit(2)
	}
    var kv map[string]interface{}
	err2 := result.Result.UnmarshalSmithyDocument(&kv)
	if err2 != nil {
		fmt.Printf("Error retrieving result, %v", err2.Error())
		os.Exit(2)
	}
    enc, _ := json.Marshal(kv)
    fmt.Println(string(enc))
    os.Exit(1)
}
