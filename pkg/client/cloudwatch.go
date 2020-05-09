package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// PutCloudwatchMetrics exports metrics to a Cloudwatch
func (c *Cassowary) PutCloudwatchMetrics(svc cloudwatchiface.CloudWatchAPI, metrics ResultMetrics) error {

	/*
		session, err := session.NewSession()
		if err != nil {
			return err
		}

		svc := cloudwatch.New(session)
	*/

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Cassowary/Metrics"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("apa"),
				Unit:       aws.String("unit"),
				Value:      aws.Float64(444.44),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("Site"),
						Value: aws.String(c.BaseURL),
					},
				},
			},
		},
	})

	if err != nil {
		return nil
	}

	return nil
}
