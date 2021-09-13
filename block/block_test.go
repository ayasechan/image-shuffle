package block

import (
	"image"
	"os"
	"testing"

	"github.com/ayasechan/image-shuffle/utils"

	"github.com/stretchr/testify/assert"
)

const (
	testKey = "bar"
)

var bs = NewBlockShuffler(testKey)

func TestRandomSeq(t *testing.T) {
	s := RandomSeq(testKey, 8)
	except := []int{7, 3, 0, 6, 2, 1, 4, 5}
	assert.Equal(t, s, except)
}

func TestCalcBlockPoint(t *testing.T) {
	p := calcBlockPoint(24, 5)
	assert.Equal(t, p, image.Point{16, 8})
}

func TestEncrypt(t *testing.T) {
	src, err := utils.LoadImage("../origin.jpg")
	assert.NoError(t, err)
	dst := bs.Encrypt(src)
	os.Mkdir("../dist", 0644)
	err = utils.WriteImage("../dist/encrypt.jpg", dst)
	assert.NoError(t, err)
}

func TestDecrypt(t *testing.T) {
	src, err := utils.LoadImage("../dist/encrypt.jpg")
	assert.NoError(t, err)
	dst := bs.Decrypt(src)
	os.Mkdir("../dist", 0644)
	err = utils.WriteImage("../dist/decrypt.jpg", dst)
	assert.NoError(t, err)
}
