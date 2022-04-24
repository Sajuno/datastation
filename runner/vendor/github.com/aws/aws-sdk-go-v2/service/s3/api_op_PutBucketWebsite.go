// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the configuration of the website that is specified in the website
// subresource. To configure a bucket as a website, you can add this subresource on
// the bucket with website configuration information such as the file name of the
// index document and any redirect rules. For more information, see Hosting
// Websites on Amazon S3
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html). This PUT
// action requires the S3:PutBucketWebsite permission. By default, only the bucket
// owner can configure the website attached to a bucket; however, bucket owners can
// allow other users to set the website configuration by writing a bucket policy
// that grants them the S3:PutBucketWebsite permission. To redirect all website
// requests sent to the bucket's website endpoint, you add a website configuration
// with the following elements. Because all requests are sent to another website,
// you don't need to provide index document name for the bucket.
//
// *
// WebsiteConfiguration
//
// * RedirectAllRequestsTo
//
// * HostName
//
// * Protocol
//
// If you
// want granular control over redirects, you can use the following elements to add
// routing rules that describe conditions for redirecting requests and information
// about the redirect destination. In this case, the website configuration must
// provide an index document for the bucket, because some requests might not be
// redirected.
//
// * WebsiteConfiguration
//
// * IndexDocument
//
// * Suffix
//
// *
// ErrorDocument
//
// * Key
//
// * RoutingRules
//
// * RoutingRule
//
// * Condition
//
// *
// HttpErrorCodeReturnedEquals
//
// * KeyPrefixEquals
//
// * Redirect
//
// * Protocol
//
// *
// HostName
//
// * ReplaceKeyPrefixWith
//
// * ReplaceKeyWith
//
// * HttpRedirectCode
//
// Amazon
// S3 has a limitation of 50 routing rules per website configuration. If you
// require more than 50 routing rules, you can use object redirect. For more
// information, see Configuring an Object Redirect
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html) in
// the Amazon S3 User Guide.
func (c *Client) PutBucketWebsite(ctx context.Context, params *PutBucketWebsiteInput, optFns ...func(*Options)) (*PutBucketWebsiteOutput, error) {
	if params == nil {
		params = &PutBucketWebsiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketWebsite", params, optFns, c.addOperationPutBucketWebsiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketWebsiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketWebsiteInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for the request.
	//
	// This member is required.
	WebsiteConfiguration *types.WebsiteConfiguration

	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt). For requests made using the Amazon Web
	// Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is
	// calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type PutBucketWebsiteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketWebsiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketWebsite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketWebsite{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketWebsiteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketWebsite(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketWebsiteUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketWebsite(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketWebsite",
	}
}

// getPutBucketWebsiteBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketWebsiteBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketWebsiteInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketWebsiteUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketWebsiteBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
