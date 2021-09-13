package block

import (
	"crypto/md5"
	"encoding/binary"
	"image"
	"image/draw"
	"strconv"
)

const BLOCK_SIZE = 8

type Block int

// 以key为种子生成一定长度的随机序列
func RandomSeq(key string, n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	hasher := md5.New()
	for i := (n - 1); i > 0; i-- {
		hasher.Reset()
		hasher.Write([]byte(key))
		hasher.Write([]byte(strconv.Itoa(i)))
		j := binary.LittleEndian.Uint64(hasher.Sum(nil)) % uint64(i+1)
		// Fisher–Yates shuffle
		s[j], s[i] = s[i], s[j]
	}
	return s
}

func calcBlockPoint(w, n int) image.Point {
	wBlockN := w / BLOCK_SIZE
	yBlockN, xBlockN := n/wBlockN, n%wBlockN

	var point image.Point
	point.X = xBlockN * BLOCK_SIZE
	point.Y = yBlockN * BLOCK_SIZE
	return point
}

func drawBlock(src image.Image, dst draw.Image, srcBlock, dstBlock Block) {
	w := src.Bounds().Dx()
	srcBlockPoint := calcBlockPoint(w, int(srcBlock))
	dstBlockPoint := calcBlockPoint(w, int(dstBlock))
	draw.Draw(
		dst,
		image.Rect(dstBlockPoint.X, dstBlockPoint.Y, (dstBlockPoint.X+BLOCK_SIZE), (dstBlockPoint.Y+BLOCK_SIZE)),
		src,
		srcBlockPoint,
		draw.Op(0),
	)
}

type BlockShuffler struct {
	Key string
}

func NewBlockShuffler(key string) *BlockShuffler {
	return &BlockShuffler{Key: key}
}

// 混淆图片
// 不能被8整除的图片边缘会被舍去！
func (bs *BlockShuffler) Encrypt(src image.Image) image.Image {
	bounds := src.Bounds()
	wBlockN := bounds.Dx() / BLOCK_SIZE
	hBlockN := bounds.Dy() / BLOCK_SIZE
	blockN := wBlockN * hBlockN
	dst := image.NewRGBA(image.Rect(0, 0, (wBlockN * BLOCK_SIZE), (hBlockN * BLOCK_SIZE)))
	s := RandomSeq(bs.Key, blockN)
	for i, v := range s {
		drawBlock(src, dst, Block(i), Block(v))
	}
	return dst
}

// 反混淆图片
func (bs *BlockShuffler) Decrypt(src image.Image) image.Image {
	bounds := src.Bounds()
	blockN := (bounds.Dx() / BLOCK_SIZE) * (bounds.Dy() / BLOCK_SIZE)
	dst := image.NewRGBA(bounds)
	s := RandomSeq(bs.Key, blockN)
	for i, v := range s {
		drawBlock(src, dst, Block(v), Block(i))
	}
	return dst
}
