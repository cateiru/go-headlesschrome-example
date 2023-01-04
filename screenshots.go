//go:build ignore

//
// HeadlessChrome 経由でスクリーンショットを撮影する例

package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
)

// https://earthquake-map.vercel.app/
// https://cateiru.com
// https://blog.cateiru.com

const TARGET_SITE_URL = "https://earthquake-map.vercel.app"
const SAVE_DIR = "screenshots"

// Based on https://pkg.go.dev/github.com/go-rod/rod#example-package-Page_screenshot
func main() {
	now := time.Now()

	Screenshot()

	fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
}

func Screenshot() {
	browser := rod.New().MustConnect()

	page := browser.MustPage(TARGET_SITE_URL).MustSetViewport(1920, 1080, 0, false).MustWaitLoad()

	img, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  1920,
			Height: 1080,

			// 解像度はここで変えられる
			Scale: 2,
		},
		CaptureBeyondViewport: true,
		OptimizeForSpeed:      true,
		FromSurface:           true,
	})
	if err != nil {
		panic(err)
	}

	err = utils.OutputFile(filepath.Join(SAVE_DIR, "my.png"), img)
	if err != nil {
		panic(err)
	}
}
