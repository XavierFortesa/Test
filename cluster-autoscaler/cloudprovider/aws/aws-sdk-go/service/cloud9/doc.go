// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloud9 provides the client and types for making API
// requests to AWS Cloud9.
//
// Cloud9 is a collection of tools that you can use to code, build, run, test,
// debug, and release software in the cloud.
//
// For more information about Cloud9, see the Cloud9 User Guide (https://docs.aws.amazon.com/cloud9/latest/user-guide).
//
// Cloud9 supports these operations:
//
//   - CreateEnvironmentEC2: Creates an Cloud9 development environment, launches
//     an Amazon EC2 instance, and then connects from the instance to the environment.
//
//   - CreateEnvironmentMembership: Adds an environment member to an environment.
//
//   - DeleteEnvironment: Deletes an environment. If an Amazon EC2 instance
//     is connected to the environment, also terminates the instance.
//
//   - DeleteEnvironmentMembership: Deletes an environment member from an environment.
//
//   - DescribeEnvironmentMemberships: Gets information about environment members
//     for an environment.
//
//   - DescribeEnvironments: Gets information about environments.
//
//   - DescribeEnvironmentStatus: Gets status information for an environment.
//
//   - ListEnvironments: Gets a list of environment identifiers.
//
//   - ListTagsForResource: Gets the tags for an environment.
//
//   - TagResource: Adds tags to an environment.
//
//   - UntagResource: Removes tags from an environment.
//
//   - UpdateEnvironment: Changes the settings of an existing environment.
//
//   - UpdateEnvironmentMembership: Changes the settings of an existing environment
//     member for an environment.
//
// See https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23 for more information on this service.
//
// See cloud9 package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloud9/
//
// # Using the Client
//
// To contact AWS Cloud9 with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Cloud9 client Cloud9 for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloud9/#New
package cloud9
