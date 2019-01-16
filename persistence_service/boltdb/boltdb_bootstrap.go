package boltdb

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type BoltDBBoostrap struct {
	db         *bolt.DB
}

func NewBoltDBBootstrap() *BoltDBBoostrap {
	return &BoltDBBoostrap{}
}

func (b *BoltDBBoostrap) Open(filePath string, option interface{}) (err error) {
	if b.db != nil {
		err = errors.New("db has already opened")
		return
	}

	var opt *bolt.Options
	if option != nil {
		opt = option.(*bolt.Options)
	}
	b.db, err = bolt.Open(filePath, 0600, opt)
	return
}

func (b *BoltDBBoostrap) Get(bucketName []byte, key []byte) (result []byte, err error) {

	err = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			logrus.Errorf("[BoltDBBoostrap] tx.Bucket error: cant find bucketName %s", string(bucketName))
			return fmt.Errorf("cant find bucketName %s", string(bucketName))
		}

		result = bucket.Get(key)

		return nil
	})

	return
}

func (b *BoltDBBoostrap) Put(bucketName []byte, key []byte, value []byte) error {

	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			var err error
			bucket, err = tx.CreateBucket(bucketName)
			if err != nil {
				logrus.Errorf("[BoltDBBoostrap] tx.CreateBucket bucketName: %s, error: %v", string(bucketName), err)
				return fmt.Errorf("create bucket failed, name=%s", string(bucketName))
			}
		}

		err := bucket.Put(key, value)
		if err != nil {
			logrus.Errorf("[BoltDBBoostrap] bucket.Put error: %v", err)
		}
		return err
	})

	return err
}

func (b *BoltDBBoostrap) Delete(bucketName []byte, key []byte) error {

	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			logrus.Errorf("[BoltDBBoostrap] tx.Bucket error: cant find bucket %s", string(bucketName))
			return fmt.Errorf("cant find bucket %s", string(bucketName))
		}

		err := bucket.Delete(key)
		if err != nil {
			logrus.Errorf("[BoltDBBoostrap] bucket.Put error: %v", err)
		}
		return err
	})

	return err
}

func (b *BoltDBBoostrap) Close() error {
	return b.db.Close()
}
