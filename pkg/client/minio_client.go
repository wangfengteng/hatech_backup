package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/hatech/backup/pkg/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io/ioutil"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

var defaultMinioClient *minio.Client

func NewMinioClient(bc *model.BackupConfig) *minio.Client {
	if defaultMinioClient == nil {
		client, err := setS3Service(bc, true)
		if err != nil {
			klog.Errorf("init minio client error : %s", err.Error())
			panic(err.Error())
		}
		defaultMinioClient = client
	}

	return defaultMinioClient
}

func setS3Service(bc *model.BackupConfig, useSSL bool) (*minio.Client, error) {
	// Initialize minio client object.
	klog.Infof("minio config :%v", bc)

	var err error
	var client = &minio.Client{}
	var cred = &credentials.Credentials{}
	var tr = http.DefaultTransport
	if bc.EndpointCA != "" {
		tr, err = setTransportCA(tr, bc.EndpointCA)
		if err != nil {
			return nil, err
		}
	}
	bucketLookup := getBucketLookupType(bc.Endpoint)
	for retries := 0; retries <= model.DefaultS3Retries; retries++ {
		// if the s3 access key and secret is not set use iam role
		if len(bc.AccessKey) == 0 && len(bc.SecretKey) == 0 {
			klog.Info("invoking set s3 service client use IAM role")
			cred = credentials.NewIAM("")
			if bc.Endpoint == "" {
				bc.Endpoint = model.S3Endpoint
			}
		} else {
			// Base64 decoding S3 accessKey and secretKey before create static credentials
			// To be backward compatible, just updating base64 encoded values
			accessKey := bc.AccessKey
			secretKey := bc.SecretKey
			if len(accessKey) > 0 {
				v, err := base64.StdEncoding.DecodeString(accessKey)
				if err == nil {
					accessKey = string(v)
				}
			}
			if len(secretKey) > 0 {
				v, err := base64.StdEncoding.DecodeString(secretKey)
				if err == nil {
					secretKey = string(v)
				}
			}
			cred = credentials.NewStatic(accessKey, secretKey, "", credentials.SignatureDefault)
		}
		client, err = minio.New(bc.Endpoint, &minio.Options{
			Creds:        cred,
			Secure:       useSSL,
			Region:       bc.Region,
			BucketLookup: bucketLookup,
			Transport:    tr,
		})
		if err != nil {
			klog.Infof("failed to init s3 client server: %v, retried %d times", err, retries)
			if retries >= model.DefaultS3Retries {
				return nil, fmt.Errorf("failed to set s3 server: %v", err)
			}
			continue
		}

		break
	}

	found, err := client.BucketExists(context.TODO(), bc.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to check s3 bucket:%s, err:%v", bc.BucketName, err)
	}
	if !found {
		return nil, fmt.Errorf("bucket %s is not found", bc.BucketName)
	}
	return client, nil
}

func setTransportCA(tr http.RoundTripper, endpointCA string) (http.RoundTripper, error) {
	ca, err := readS3EndpointCA(endpointCA)
	if err != nil {
		return tr, err
	}
	if !isValidCertificate(ca) {
		return tr, fmt.Errorf("s3-endpoint-ca is not a valid x509 certificate")
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca)

	tr.(*http.Transport).TLSClientConfig = &tls.Config{
		RootCAs: certPool,
	}

	return tr, nil
}

func isValidCertificate(c []byte) bool {
	p, _ := pem.Decode(c)
	if p == nil {
		return false
	}
	_, err := x509.ParseCertificates(p.Bytes)
	if err != nil {
		return false
	}
	return true
}

func getBucketLookupType(endpoint string) minio.BucketLookupType {
	if endpoint == "" {
		return minio.BucketLookupAuto
	}
	if strings.Contains(endpoint, "aliyun") {
		return minio.BucketLookupDNS
	}
	return minio.BucketLookupAuto
}

func readS3EndpointCA(endpointCA string) ([]byte, error) {
	// I expect the CA to be passed as base64 string OR a file system path.
	// I do this to be able to pass it through rke/rancher api without writing it
	// to the backup container filesystem.
	ca, err := base64.StdEncoding.DecodeString(endpointCA)
	if err == nil {
		klog.Info("reading s3-endpoint-ca as a base64 string")
	} else {
		ca, err = ioutil.ReadFile(endpointCA)
		klog.Infof("reading s3-endpoint-ca from [%v]", endpointCA)
	}
	return ca, err
}
