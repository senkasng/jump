package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"flag"
)

var ip,name ,key,secret  string

func init(){
	flag.StringVar(&key,"key","","aws account key")
	flag.StringVar(&secret,"secret","","aws account secret")
	flag.Parse()
	if key == "" || secret == "" {
		os.Exit(1)
	}
}

func main() {
	file,_ := os.Create("./.cloudhosts")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Region:      aws.String(endpoints.CnNorth1RegionID),
	}))

	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	for _,reservation:= range result.Reservations {
		
	if reservation.Instances[0].PrivateIpAddress != nil  {
		ip = *(reservation.Instances[0].PrivateIpAddress)
	}
	if 	reservation.Instances[0].Tags != nil {
		for _,tags := range reservation.Instances[0].Tags {
			if *tags.Key == "Name"{
				name = *tags.Value
	
			}
		}
	}
	data := []byte(name+"  "+ip+  "\n")
	_,err := file.Write(data)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	}
}




