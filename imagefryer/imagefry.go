//imagefry applies pseudo random filter to an image
package fry

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"time"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func randFilter(imageData image.Image, imgCfg image.Config, timesFry int) image.Image {
	// copy old image to a new template

	alteredImage := image.NewRGBA(imageData.Bounds())
	draw.Draw(alteredImage, imageData.Bounds(), imageData, image.Point{}, draw.Over)

	width := imgCfg.Width
	height := imgCfg.Height

	// apply random changes to the image
	for i := 0; i < timesFry; i++ {
		for y := 0; y < height; y++ {
			rand.Seed(time.Now().UTC().UnixNano())
			for x := 0; x < width; x++ {
				r, g, b, a := alteredImage.At(x, y).RGBA()
				newColor := color.RGBA{randColor(uint8(r)), randColor(uint8(g)), randColor(uint8(b)), uint8(a)}
				alteredImage.Set(x, y, newColor)
			}
		}
	}
	return alteredImage
}

func randColor(origColor uint8) uint8 {
	key := rand.Intn(1)
	if key == 0 {
		return origColor + uint8(rand.Intn(10))
	} else {
		return origColor - uint8(rand.Intn(10))
	}
}

// Decode reads and analyzes the given reader as a GIF image
func SplitAnimatedGIF(reader io.Reader, timesFry int) (err error, newGif *gif.GIF) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error while decoding: %s", r)
		}
	}()
	inGif, err := gif.DecodeAll(reader)
	fryGif := gif.GIF{}

	if err != nil {
		return err, nil
	}

	for _, srcImg := range inGif.Image {
		var imgCfg image.Config
		imgCfg.Width, imgCfg.Height = srcImg.Rect.Dx(), srcImg.Rect.Dy()
		alteredImage := randFilter(srcImg, imgCfg, timesFry)
		bounds := alteredImage.Bounds()
		alteredPalette := image.NewPaletted(bounds, srcImg.Palette)
		draw.Draw(alteredPalette, alteredPalette.Rect, alteredImage, bounds.Min, draw.Over)

		// save current frame "stack". This will overwrite an existing file with that name
		fryGif.Delay = append(fryGif.Delay, 8)
		fryGif.Image = append(fryGif.Image, alteredPalette)
	}
	//gif.EncodeAll(out, &fryGif) //ignores encoding errors
	return nil, &fryGif
}

//fry an image
func openDecodeFilterStatic(ImageFile multipart.File, timesFry int) (image.Image, string) {
	imageData, imageType, err := image.Decode(ImageFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "perlin: %v\n", err)
	}

	ImageFile.Seek(0, 0)

	imgCfg, _, err := image.DecodeConfig(ImageFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ImageFile.Seek(0, 0)

	newImg := randFilter(imageData, imgCfg, timesFry)
	return newImg, imageType
}

func Fry(ImageFile *multipart.File, timesFry int) (*image.Image, string) {
	newImg, imageType := openDecodeFilterStatic(*ImageFile, timesFry)
	return &newImg, imageType
}
