package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3FileStorage struct {
	bucket string
	client *s3.Client
}

func NewS3FileStorage(bucket, region, endpoint, apiKey, apiSecret string, forcePathStyle bool) (*S3FileStorage, error) {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(apiKey, apiSecret, "")),
		config.WithRegion(region),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg,
		func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = forcePathStyle
		})

	m := &S3FileStorage{
		bucket: bucket,
		client: client,
	}

	return m, nil
}

func (fs *S3FileStorage) getObject(ctx context.Context, input *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3Object, error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}

	attrIn := s3.HeadObjectInput{
		Bucket: input.Bucket,
		Key:    input.Key,
	}

	attrOut, err := fs.client.HeadObject(ctx, &attrIn)
	if err != nil {
		return nil, err
	}

	objOut, err := fs.client.GetObject(ctx, input, optFns...)
	if err != nil {
		return nil, err
	}

	obj := s3Object{
		client:   fs.client,
		input:    *input,
		resp:     objOut,
		length:   *attrOut.ContentLength,
		lengthOk: true,
		body:     objOut.Body,
	}

	return &obj, nil
}

func (fs *S3FileStorage) OpenFileByKey(key string) (reader io.ReadSeekCloser, err error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(key),
	}

	file, err := fs.getObject(context.Background(), input)
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}
	return file, nil
}

func (fs *S3FileStorage) SaveByKey(src io.Reader, key, name, contentType string) (err error) {

	input := &s3.PutObjectInput{
		Bucket:             aws.String(fs.bucket),
		Key:                aws.String(key),
		Body:               src,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String(fmt.Sprintf("attachment; filename*=UTF-8''%s", url.PathEscape(name))),
	}

	uploader := manager.NewUploader(fs.client)

	_, err = uploader.Upload(context.Background(), input)
	return
}

func (fs *S3FileStorage) DeleteByKey(key string) (err error) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(key),
	}

	_, err = fs.client.DeleteObject(context.Background(), input)
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return ErrFileNotFound
		}
		return err
	}

	return nil
}

func (fs *S3FileStorage) GenerateAccessURL(key string) (string, error) {
	pc := s3.NewPresignClient(fs.client)

	req, _ := pc.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(fs.bucket),
		Key:    aws.String(key),
	}, func(options *s3.PresignOptions) {
		options.Expires = 5 * time.Minute
	})

	return req.URL, nil
}
