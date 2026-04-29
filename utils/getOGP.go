package utils

import (
	"encoding/json"
	"fmt"
	"notion2atlas/domain"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func GetOGP(urlStr string) (*domain.OGPResult, error) {
	if urlStr == "" {
		return nil, fmt.Errorf("url is empty")
	}
	// Chrome を sandbox 無しで起動（GitHub Actions 対応）
	u := launcher.New().
		Headless(true).
		NoSandbox(true).
		Set("disable-setuid-sandbox").
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect().
		MustIgnoreCertErrors(true)

	defer browser.MustClose()

	// ページ作成
	page := browser.MustPage()

	// UserAgent 設定
	err := page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/115 Safari/537.36",
	})
	if err != nil {
		fmt.Println("error in utils/GetOGP/page.SetUserAgent url: " + urlStr)
		return nil, err
	}

	// ページ移動
	// page.MustNavigate(urlStr).MustWaitLoad()
	err = page.Navigate(urlStr)
	if err != nil {
		fmt.Printf("❌ failed to navigate to [%s]: %v\n", urlStr, err)
		return &domain.OGPResult{
			Title:       "ページが見つかりません",
			Icon:        "",
			Image:       "",
			Description: urlStr,
		}, nil
	}
	err = page.WaitLoad()
	if err != nil {
		return nil, err
	}

	// タイトル取得
	title := page.MustInfo().Title

	// JS 実行して OGP データ取得
	res, err := page.Evaluate(&rod.EvalOptions{
		AwaitPromise: true,
		JS: `
			(() => {
				const ogp = {
					description: document.querySelector('meta[name="description"]')?.content || "",
					image: document.querySelector('meta[property="og:image"]')?.content
						? new URL(document.querySelector('meta[property="og:image"]').content, document.baseURI).href
						: null,
					icon: document.querySelector('link[rel="icon"]')?.href
						? new URL(document.querySelector('link[rel="icon"]').href, document.baseURI).href
						: null
				};
				return JSON.stringify(ogp);
			})
		`,
	})
	if err != nil {
		fmt.Println("error in utils/GetOGP/page.Evaluate url: " + urlStr)
		return nil, err
	}

	var ogp domain.OGPResult
	if err := json.Unmarshal([]byte(res.Value.Str()), &ogp); err != nil {
		fmt.Println("error in utils/GetOGP/json.Unmarshal url:" + urlStr)
		return nil, err
	}

	return &domain.OGPResult{
		Title:       title,
		Icon:        ogp.Icon,
		Image:       ogp.Image,
		Description: ogp.Description,
	}, nil
}

// Browser
// CDPClient
// CoveredError
// DefaultLogger
// DefaultSleeper
// Element
// ElementNotFoundError
// atg Elements
// Eval
// EvalError
// EvalOptions
// ExpectElementError
// ExpectElementsError
// Hijack
// HijackRequest
// HijackResponse
// HijackRouter
// InvisibleShapeError
// KeyAction
// KeyActionPress
// KeyActionRelease
// KeyActionType
// KeyActionTypeKey
// KeyActions
// Keyboard
// Message
// Mouse
// NavigationError
// New
// NewBrowserPool
// NewPagePool
// NewPool
// NewSt reamReader
// NoPointerEventsError
// NoShadowRootError
// NotFoundSleeper
// NotInteractableError
// ObjectNotFoundError
// Page
// PageCloseCanceledError
// PageNotFoundError
// a Pages
// Pool
// RaceContext
// ScrollScreenshotOptions
// SearchResult SelectorType
// SelectorTypeCSSSector
// SelectorTypeRegex
// SelectorTypeText
// StreamReader
// Touch
// TraceType
// TraceTypeInput
// TraceTypeQuery
// TraceTypeWait
// TraceTypeWaitRequests
// TraceTypeWaitRequestsIdle
// Try
// TryError
// New().BrowserContextID
// DefaultLogger.Fatal DefaultLoqger.Fatalf
