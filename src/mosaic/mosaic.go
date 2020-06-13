package mosaic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sync"
)

func init() {
	TILESDB = tilesDB()
}

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

func (db *DB) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	//delete(db.store, filename)
	db.mutex.Unlock()
	return filename
}

// resize an image by its ratio e.g. ratio 2 means reduce the size by 1/2, 10 means reduce the size by 1/10
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	width := bounds.Dx()
	ratio := width / newWidth
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// find the average color of the picture
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

var TILESDB map[string][3]float64

func CloneTilesDB() DB {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	tiles := DB{
		store: db,
		mutex: &sync.Mutex{},
	}
	return tiles
}

// populate a tiles database in memory
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("public/tiles")
	for _, f := range files {
		name := filepath.Join("public/tiles", f.Name())
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating tiles db:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, "when populating tiles db:", err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

// find the Eucleadian distance between 2 points
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

// find the square
func sq(n float64) float64 {
	return n * n
}

// cut out the image and return individual channels with image.Image
// no encoding of JPEG
func Cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	sp := image.Point{0, 0}
	go func() {
		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}
				nearest := db.nearest(color)
				file, err := os.Open(nearest)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error in decoding nearest", err, nearest)
					}
				} else {
					fmt.Println("error opening file when creating mosaic:", nearest)
				}
				file.Close()
			}
		}
		c <- newimage.SubImage(newimage.Rect)
	}()

	return c
}

// combine the images and return the encoding string
func Combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	c := make(chan string)
	// start a goroutine
	go func() {
		var wg sync.WaitGroup
		newimage := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done()
		}
		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool
		for {
			select {
			case s1, ok1 = <-c1:
				go copy(newimage, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(newimage, s2.Bounds(), s2, image.Point{r.Max.X / 2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(newimage, s3.Bounds(), s3, image.Point{r.Min.X, r.Max.Y / 2})
			case s4, ok4 = <-c4:
				go copy(newimage, s4.Bounds(), s4, image.Point{r.Max.X / 2, r.Max.Y / 2})
			}
			if ok1 && ok2 && ok3 && ok4 {
				break
			}
		}
		// wait till all copy goroutines are complete
		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, newimage, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}
