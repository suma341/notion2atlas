package usecase

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"time"

	"github.com/chromedp/chromedp"
)

func SaveStaticPageOGPPicture() error {
	staticPages := []domain.PageEntity{}
	var isInfoOGPExist = filemanager.FileExists(constants.OGP_DIR + "/infos.png")
	if !isInfoOGPExist {
		staticPages = append(staticPages, domain.CreatePage("部活情報", "emoji", "ℹ️", "infos"))
	}
	var isBasicOGPExist = filemanager.FileExists(constants.OGP_DIR + "/basic.png")
	if !isBasicOGPExist {
		staticPages = append(staticPages, domain.CreatePage("基礎班カリキュラム", "emoji", "🔰", "basic"))
	}
	var isAnswerOGPExist = filemanager.FileExists(constants.OGP_DIR + "/answers.png")
	if !isAnswerOGPExist {
		staticPages = append(staticPages, domain.CreatePage("解答ページ", "emoji", "✔️", "answers"))
	}
	for _, p := range staticPages {
		err := SaveOGPPicture(p)
		if err != nil {
			fmt.Println("error in usecase/SaveStaticPageOGPPicture/SaveOGPPicture")
			return err
		}
	}
	return nil
}

func SaveOGPPicture(p domain.PageIf) error {
	_, coverUrl := p.GetCover()
	iconType, iconUrl := p.GetIcon()
	title := p.GetTitle()
	id := p.GetId()
	fmt.Printf("title: %s\n", title)
	html := createHTML(coverUrl, iconType, iconUrl, title)
	pngByte, err := html2png(html, 1203, 630)
	if err != nil {
		fmt.Printf("html: \n%s\n", html)
		fmt.Println("error in presentation/HandleCreateGGP/usecase.HTMLToPNG")
		return err
	}
	err = filemanager.SavePNG(fmt.Sprintf("%s/%s.png", constants.OGP_DIR, id), pngByte)
	if err != nil {
		fmt.Println("error in presentation/HandleCreateGGP/filemanager.SavePNG")
		return err
	}
	return nil
}

func html2png(html string, width, height int) ([]byte, error) {

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true), // Docker環境ならこれも追加推奨
	}

	// 1. Allocatorを作成 (ブラウザプロセスの管理)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	// 2. Contextを作成 (タブの管理)
	// ここで ctx を受け取る
	ctx, ctxCancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer ctxCancel()

	// 3. タイムアウトを設定 (ctx を上書き、または新しい変数で受け取る)
	// 30秒だと重い画像の時にギリギリなので、少し余裕を持たせてもいいかもしれません
	ctx, timeoutCancel := context.WithTimeout(ctx, 45*time.Second)
	defer timeoutCancel()

	dataURL := "data:text/html;charset=utf-8," + url.PathEscape(html)

	var buf []byte

	// 4. 実行には「タイムアウト設定済みの ctx」を渡す
	err := chromedp.Run(ctx,
		chromedp.EmulateViewport(int64(width), int64(height)),
		chromedp.Navigate(dataURL),
		// 画像が重い場合、body だけでなく img タグの出現を待つのも手です
		chromedp.WaitReady("body"),
		chromedp.Sleep(500*time.Millisecond), // 念のため少し長めに
		chromedp.FullScreenshot(&buf, 90),
	)

	if err != nil {
		return nil, err
	}
	return buf, nil
}

func createHTML(coverUrl, iconType, iconUrl, title string) string {
	fileIcon := "https://ryukoku-horizon.github.io/horizon-atlas/pngwing.png"

	// --- icon 部分のHTML ---
	iconHTML := ""
	if iconType == "" {
		iconHTML = fmt.Sprintf(`<img src="%s" style="width: 9rem; height: 9rem;position: absolute; top: 235px; left: 90px" >`, fileIcon)
	} else if iconType != "emoji" {
		iconHTML = fmt.Sprintf(`<img src="%s" style="width: 9rem; height: 9rem;position: absolute; top: 235px; left: 90px" >`, iconUrl)
	} else { // emoji
		iconHTML = fmt.Sprintf(`<p style="font-size: 7rem; position: absolute; top: 130px; left: 90px">%s</p>`, iconUrl)
	}

	// --- cover がある場合 ---
	coverImage := ""
	coverBlock := ""

	if coverUrl != "" {
		coverImage = fmt.Sprintf(
			`<img src="%s" style="position: absolute;top: 0px;left:-10px; width: 101%%;height: 320px;object-fit: cover;">`,
			coverUrl,
		)

		coverBlock = fmt.Sprintf(`
		<div>
			%s
			<h2 style="font-size: 64px; font-style: bold;position: absolute; top: 320px; left:40px;">%s</h2>
			<p style="font-size:28px; position: absolute; top: 440px; left: 40px;color:rgb(158, 159, 159);">on HorizonAtlas</p>
		</div>`, iconHTML, title)
	}

	// --- cover がない場合 ---
	noCoverIcon := ""
	if iconType == "" {
		noCoverIcon = fmt.Sprintf(`<img src="%s" style="width: 9rem; height: 9rem;" >`, fileIcon)
	} else if iconType != "emoji" {
		noCoverIcon = fmt.Sprintf(`<img src="%s" style="width: 9rem; height: 9rem;" >`, iconUrl)
	} else {
		noCoverIcon = fmt.Sprintf(`<p style="font-size: 7.5rem;">%s</p>`, iconUrl)
	}

	noCoverBlock := fmt.Sprintf(`
	<div  style="display: flex;flex-direction: column;align-items: center; justify-content: center; flex:1;">
		<div style="display: flex; align-items: center; justify-content: center; flex:1; grid-gap: 5px;">
			%s
			<h2 style="font-size: 64px; font-style: bold;">%s</h2>
		</div>
		<p style="font-size:28px; color:rgb(158, 159, 159);position: absolute; bottom: 20px; right: 50px;">on HorizonAtlas</p>
	</div>`, noCoverIcon, title)

	// --- 最終HTML ---
	html := fmt.Sprintf(`
	<html>
		<body style="
			margin: 0; width: 1203px; height: 630px;
			background: white;
			display: flex; justify-content: center; align-items: center;
			font-family: 'Noto Sans CJK JP', 'Noto Sans JP', sans-serif;
			">
			%s
			%s
		</body>
	</html>`,
		coverImage,
		func() string {
			if coverUrl != "" {
				return coverBlock
			}
			return noCoverBlock
		}(),
	)

	return html
}
