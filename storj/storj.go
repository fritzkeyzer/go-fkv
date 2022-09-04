package storj

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"storj.io/uplink"
	"strings"
)

type KV struct {
	storjAccess *uplink.Access
	bucketName  string
}

func NewKV(bucketName, accessGrant string) (*KV, error) {
	access, err := uplink.ParseAccess(accessGrant)
	if err != nil {
		return nil, fmt.Errorf("storj: could not request access grant: %v", err)
	}

	// open up the Project we will be working with.
	project, err := uplink.OpenProject(context.TODO(), access)
	if err != nil {
		return nil, fmt.Errorf("storj: could not open project: %v", err)
	}
	defer project.Close()

	// ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(context.TODO(), bucketName)
	if err != nil {
		return nil, fmt.Errorf("storj: could not ensure bucket: %w", err)
	}

	kv := KV{
		storjAccess: access,
		bucketName:  bucketName,
	}
	return &kv, nil
}

func (kv *KV) Get(key string, ptr any) (found bool, err error) {
	// open up the Project we will be working with.
	project, err := uplink.OpenProject(context.TODO(), kv.storjAccess)
	if err != nil {
		return false, fmt.Errorf("storj: could not open project: %v", err)
	}
	defer project.Close()

	// initiate a download stream
	download, err := project.DownloadObject(context.TODO(), kv.bucketName, key, nil)
	if err != nil && strings.Contains(err.Error(), "object not found") {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("storj: initiating download stream: %v", err)
	}
	defer download.Close()

	// read everything from the download stream
	receivedContents, err := io.ReadAll(download)
	if err != nil {
		return true, fmt.Errorf("storj: could not read data: %v", err)
	}

	if err := json.Unmarshal(receivedContents, ptr); err != nil {
		return true, fmt.Errorf("decoding: %w", err)
	}

	return true, nil
}

func (kv *KV) Set(key string, object any) error {
	data, err := json.Marshal(object)
	if err != nil {
		return fmt.Errorf("encoding: %v", err)
	}

	// open up the Project we will be working with.
	project, err := uplink.OpenProject(context.TODO(), kv.storjAccess)
	if err != nil {
		return fmt.Errorf("storj: could not open project: %v", err)
	}
	defer project.Close()

	// initiate upload stream
	upload, err := project.UploadObject(context.TODO(), kv.bucketName, key, nil)
	if err != nil {
		return fmt.Errorf("storj: initiating upload stream: %v", err)
	}

	// copy the data to the upload.
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("storj: could not upload data: %v", err)
	}

	// commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("storj: could not commit uploaded object: %v", err)
	}

	return nil
}

func (kv *KV) Delete(key string) error {
	project, err := uplink.OpenProject(context.TODO(), kv.storjAccess)
	if err != nil {
		return fmt.Errorf("storj: could not open project: %v", err)
	}
	defer project.Close()

	// delete object
	_, err = project.DeleteObject(context.TODO(), kv.bucketName, key)
	if err != nil {
		return fmt.Errorf("storj: deleting object: %v", err)
	}

	return nil
}
