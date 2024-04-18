// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/snowball"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To cancel a cluster job
// This operation cancels a cluster job. You can only cancel a cluster job while it's
// in the AwaitingQuorum status.
func ExampleSnowball_CancelCluster_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.CancelClusterInput{
		ClusterId: aws.String("CID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.CancelCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To cancel a job for a Snowball device
// This operation cancels a job. You can only cancel a job before its JobState value
// changes to PreparingAppliance.
func ExampleSnowball_CancelJob_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.CancelJobInput{
		JobId: aws.String("JID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.CancelJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create an address for a job
// This operation creates an address for a job. Addresses are validated at the time
// of creation. The address you provide must be located within the serviceable area
// of your region. If the address is invalid or unsupported, then an exception is thrown.
func ExampleSnowball_CreateAddress_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.CreateAddressInput{
		Address: &snowball.Address{
			City:            aws.String("Seattle"),
			Company:         aws.String("My Company's Name"),
			Country:         aws.String("USA"),
			Name:            aws.String("My Name"),
			PhoneNumber:     aws.String("425-555-5555"),
			PostalCode:      aws.String("98101"),
			StateOrProvince: aws.String("WA"),
			Street1:         aws.String("123 Main Street"),
		},
	}

	result, err := svc.CreateAddress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidAddressException:
				fmt.Println(snowball.ErrCodeInvalidAddressException, aerr.Error())
			case snowball.ErrCodeUnsupportedAddressException:
				fmt.Println(snowball.ErrCodeUnsupportedAddressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a cluster
// Creates an empty cluster. Each cluster supports five nodes. You use the CreateJob
// action separately to create the jobs for each of these nodes. The cluster does not
// ship until these five node jobs have been created.
func ExampleSnowball_CreateCluster_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.CreateClusterInput{
		AddressId:   aws.String("ADID1234ab12-3eec-4eb3-9be6-9374c10eb51b"),
		Description: aws.String("MyCluster"),
		JobType:     aws.String("LOCAL_USE"),
		KmsKeyARN:   aws.String("arn:aws:kms:us-east-1:123456789012:key/abcd1234-12ab-34cd-56ef-123456123456"),
		Notification: &snowball.Notification{
			NotifyAll: aws.Bool(false),
		},
		Resources: &snowball.JobResource{
			S3Resources: []*snowball.S3Resource{
				{
					BucketArn: aws.String("arn:aws:s3:::MyBucket"),
					KeyRange:  &snowball.KeyRange{},
				},
			},
		},
		RoleARN:        aws.String("arn:aws:iam::123456789012:role/snowball-import-S3-role"),
		ShippingOption: aws.String("SECOND_DAY"),
		SnowballType:   aws.String("EDGE"),
	}

	result, err := svc.CreateCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			case snowball.ErrCodeInvalidInputCombinationException:
				fmt.Println(snowball.ErrCodeInvalidInputCombinationException, aerr.Error())
			case snowball.ErrCodeEc2RequestFailedException:
				fmt.Println(snowball.ErrCodeEc2RequestFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a job
// Creates a job to import or export data between Amazon S3 and your on-premises data
// center. Your AWS account must have the right trust policies and permissions in place
// to create a job for Snowball. If you're creating a job for a node in a cluster, you
// only need to provide the clusterId value; the other job attributes are inherited
// from the cluster.
func ExampleSnowball_CreateJob_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.CreateJobInput{
		AddressId:   aws.String("ADID1234ab12-3eec-4eb3-9be6-9374c10eb51b"),
		Description: aws.String("My Job"),
		JobType:     aws.String("IMPORT"),
		KmsKeyARN:   aws.String("arn:aws:kms:us-east-1:123456789012:key/abcd1234-12ab-34cd-56ef-123456123456"),
		Notification: &snowball.Notification{
			NotifyAll: aws.Bool(false),
		},
		Resources: &snowball.JobResource{
			S3Resources: []*snowball.S3Resource{
				{
					BucketArn: aws.String("arn:aws:s3:::MyBucket"),
					KeyRange:  &snowball.KeyRange{},
				},
			},
		},
		RoleARN:                    aws.String("arn:aws:iam::123456789012:role/snowball-import-S3-role"),
		ShippingOption:             aws.String("SECOND_DAY"),
		SnowballCapacityPreference: aws.String("T80"),
		SnowballType:               aws.String("STANDARD"),
	}

	result, err := svc.CreateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			case snowball.ErrCodeInvalidInputCombinationException:
				fmt.Println(snowball.ErrCodeInvalidInputCombinationException, aerr.Error())
			case snowball.ErrCodeClusterLimitExceededException:
				fmt.Println(snowball.ErrCodeClusterLimitExceededException, aerr.Error())
			case snowball.ErrCodeEc2RequestFailedException:
				fmt.Println(snowball.ErrCodeEc2RequestFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe an address for a job
// This operation describes an address for a job.
func ExampleSnowball_DescribeAddress_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.DescribeAddressInput{
		AddressId: aws.String("ADID1234ab12-3eec-4eb3-9be6-9374c10eb51b"),
	}

	result, err := svc.DescribeAddress(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe all the addresses you've created for AWS Snowball
// This operation describes all the addresses that you've created for AWS Snowball.
// Calling this API in one of the US regions will return addresses from the list of
// all addresses associated with this account in all US regions.
func ExampleSnowball_DescribeAddresses_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.DescribeAddressesInput{}

	result, err := svc.DescribeAddresses(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidNextTokenException:
				fmt.Println(snowball.ErrCodeInvalidNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a cluster
// Returns information about a specific cluster including shipping information, cluster
// status, and other important metadata.
func ExampleSnowball_DescribeCluster_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.DescribeClusterInput{
		ClusterId: aws.String("CID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.DescribeCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a job you've created for AWS Snowball
// This operation describes a job you've created for AWS Snowball.
func ExampleSnowball_DescribeJob_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.DescribeJobInput{
		JobId: aws.String("JID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.DescribeJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get the manifest for a job you've created for AWS Snowball
// Returns a link to an Amazon S3 presigned URL for the manifest file associated with
// the specified JobId value. You can access the manifest file for up to 60 minutes
// after this request has been made. To access the manifest file after 60 minutes have
// passed, you'll have to make another call to the GetJobManifest action.
//
// The manifest is an encrypted file that you can download after your job enters the
// WithCustomer status. The manifest is decrypted by using the UnlockCode code value,
// when you pass both values to the Snowball through the Snowball client when the client
// is started for the first time.
//
// As a best practice, we recommend that you don't save a copy of an UnlockCode value
// in the same location as the manifest file for that job. Saving these separately helps
// prevent unauthorized parties from gaining access to the Snowball associated with
// that job.
//
// The credentials of a given job, including its manifest file and unlock code, expire
// 90 days after the job is created.
func ExampleSnowball_GetJobManifest_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.GetJobManifestInput{
		JobId: aws.String("JID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.GetJobManifest(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get the unlock code for a job you've created for AWS Snowball
// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 90 days after the associated job has been created.
//
// The UnlockCode value is a 29-character code with 25 alphanumeric characters and 4
// hyphens. This code is used to decrypt the manifest file when it is passed along with
// the manifest to the Snowball through the Snowball client when the client is started
// for the first time.
//
// As a best practice, we recommend that you don't save a copy of the UnlockCode in
// the same location as the manifest file for that job. Saving these separately helps
// prevent unauthorized parties from gaining access to the Snowball associated with
// that job.
func ExampleSnowball_GetJobUnlockCode_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.GetJobUnlockCodeInput{
		JobId: aws.String("JID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.GetJobUnlockCode(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To see your Snowball service limit and the number of Snowballs you have in use
// Returns information about the Snowball service limit for your account, and also the
// number of Snowballs your account has in use.
//
// The default service limit for the number of Snowballs that you can have at one time
// is 1. If you want to increase your service limit, contact AWS Support.
func ExampleSnowball_GetSnowballUsage_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.GetSnowballUsageInput{}

	result, err := svc.GetSnowballUsage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of jobs in a cluster that you've created for AWS Snowball
// Returns an array of JobListEntry objects of the specified length. Each JobListEntry
// object is for a job in the specified cluster and contains a job's state, a job's
// ID, and other information.
func ExampleSnowball_ListClusterJobs_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.ListClusterJobsInput{
		ClusterId: aws.String("CID123e4567-e89b-12d3-a456-426655440000"),
	}

	result, err := svc.ListClusterJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidNextTokenException:
				fmt.Println(snowball.ErrCodeInvalidNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of clusters that you've created for AWS Snowball
// Returns an array of ClusterListEntry objects of the specified length. Each ClusterListEntry
// object contains a cluster's state, a cluster's ID, and other important status information.
func ExampleSnowball_ListClusters_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.ListClustersInput{}

	result, err := svc.ListClusters(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidNextTokenException:
				fmt.Println(snowball.ErrCodeInvalidNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get a list of jobs that you've created for AWS Snowball
// Returns an array of JobListEntry objects of the specified length. Each JobListEntry
// object contains a job's state, a job's ID, and a value that indicates whether the
// job is a job part, in the case of export jobs. Calling this API action in one of
// the US regions will return jobs from the list of all jobs associated with this account
// in all US regions.
func ExampleSnowball_ListJobs_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.ListJobsInput{}

	result, err := svc.ListJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidNextTokenException:
				fmt.Println(snowball.ErrCodeInvalidNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a cluster
// This action allows you to update certain parameters for a cluster. Once the cluster
// changes to a different state, usually within 60 minutes of it being created, this
// action is no longer available.
func ExampleSnowball_UpdateCluster_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.UpdateClusterInput{
		AddressId:   aws.String("ADID1234ab12-3eec-4eb3-9be6-9374c10eb51b"),
		ClusterId:   aws.String("CID123e4567-e89b-12d3-a456-426655440000"),
		Description: aws.String("updated-cluster-name"),
	}

	result, err := svc.UpdateCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			case snowball.ErrCodeInvalidInputCombinationException:
				fmt.Println(snowball.ErrCodeInvalidInputCombinationException, aerr.Error())
			case snowball.ErrCodeEc2RequestFailedException:
				fmt.Println(snowball.ErrCodeEc2RequestFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a job
// This action allows you to update certain parameters for a job. Once the job changes
// to a different job state, usually within 60 minutes of the job being created, this
// action is no longer available.
func ExampleSnowball_UpdateJob_shared00() {
	svc := snowball.New(session.New())
	input := &snowball.UpdateJobInput{
		AddressId:                  aws.String("ADID1234ab12-3eec-4eb3-9be6-9374c10eb51b"),
		Description:                aws.String("updated-job-name"),
		JobId:                      aws.String("JID123e4567-e89b-12d3-a456-426655440000"),
		ShippingOption:             aws.String("NEXT_DAY"),
		SnowballCapacityPreference: aws.String("T100"),
	}

	result, err := svc.UpdateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case snowball.ErrCodeInvalidResourceException:
				fmt.Println(snowball.ErrCodeInvalidResourceException, aerr.Error())
			case snowball.ErrCodeInvalidJobStateException:
				fmt.Println(snowball.ErrCodeInvalidJobStateException, aerr.Error())
			case snowball.ErrCodeKMSRequestFailedException:
				fmt.Println(snowball.ErrCodeKMSRequestFailedException, aerr.Error())
			case snowball.ErrCodeInvalidInputCombinationException:
				fmt.Println(snowball.ErrCodeInvalidInputCombinationException, aerr.Error())
			case snowball.ErrCodeClusterLimitExceededException:
				fmt.Println(snowball.ErrCodeClusterLimitExceededException, aerr.Error())
			case snowball.ErrCodeEc2RequestFailedException:
				fmt.Println(snowball.ErrCodeEc2RequestFailedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
