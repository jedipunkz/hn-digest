---
source: "https://github.com/garagehq/Minus-chrome-extension"
hn_url: "https://news.ycombinator.com/item?id=49246731"
title: "Show HN: Minus – in-browser AI ad blocker that swaps ads for flashcards"
article_title: "GitHub - garagehq/Minus-chrome-extension · GitHub"
author: "NickySlicks"
captured_at: "2026-08-10T17:44:13Z"
capture_tool: "hn-digest"
hn_id: 49246731
score: 1
comments: 0
posted_at: "2026-08-10T17:16:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Minus – in-browser AI ad blocker that swaps ads for flashcards

- HN: [49246731](https://news.ycombinator.com/item?id=49246731)
- Source: [github.com](https://github.com/garagehq/Minus-chrome-extension)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T17:16:14Z

## Translation

タイトル: Show HN: Minus – 広告をフラッシュカードに置き換えるブラウザ内 AI 広告ブロッカー
記事タイトル: GitHub - Garagehq/Minus-chrome-extension · GitHub
説明: GitHub でアカウントを作成して、garagehq/Minus-chrome-extension の開発に貢献します。

記事本文:
GitHub - Garagehq/Minus-chrome-extension · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ガレージ
/
マイナスクロム拡張機能
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
82 コミット 82 コミット docs docs extension extension サーバー サーバー src src テスト テスト .gitignore .gitignore ACCEPTANCE.md ACCEPTANCE.md CLAUDE.md CLAUDE.md LICENSE LICENSE PR

IVACY.md PRIVACY.md README.md README.md build.mjs build.mjs build_model_index.mjs build_model_index.mjs docs_perf_soak.json docs_perf_soak.json gen_icons.mjs gen_icons.mjs package-lock.json package-lock.json package.json package.json package.mjs package.mjs Popup_preview.png Popup_preview.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Minus — ビジョン広告ブロッカー (Chrome 拡張機能)
マイナスHDMIのブラウザいとこ
デバイス: フィルター リストを照合する代わりに、ページ要素を検索します。
ビジョン言語モデルは完全にブラウザ内で実行され、
言語フラッシュカード付きの広告です。ブロックされた広告はすべて、言語のフラッシュカードを教えてくれます。
代わりに言葉を。
マシンからは何も残されません。サーバーも、フィルターリストのサブスクリプションも、何もありません。
テレメトリー。モデルはあなたが見ているものを見ます。
実際のページ、実際の広告: theverge.com のバナーとサイドバーの広告スロットがカバーされます。
フラッシュカード (「la puerta — ドア」、「cuánto — いくら」)、それぞれにモデルの情報が含まれています。
信頼タグとかすかな ✕ で広告が表示されます。記事は無修正です。
従来の広告ブロッカーは、URL をフィルタ リストと照合します。マイナス分類
ピクセル - ファーストパーティ広告、スポンサー付きタイル、ネイティブの「チャットボックス」プレースメント
リストでは説明できないストリーミングビデオ広告は依然として捕捉されますが、何も捕まりません
アドテクによってドメインがローテーションされると中断されます。
候補者を見つけます。コンテンツ スクリプト (フレームごとに実行) が収集する
広告である可能性のある要素: <img> と <iframe> 、およびコンテナ
ID/クラスは広告スロットのような匂いがします。明らかに広告ではないものは構造的にフィルタリングされます -
同意バナー、ビデオ プレーヤー (ビデオ パスに属します)、アンロード済み
画像プレースホルダー、コンテンツ列の形をしたボックス。
見てください。各候補について、Minus は実際のピクセルを取得します。
表示要素のタブのスクリーンショット (レート制限、アクティブなタブのみ)、または
直接フレーム読み取り

<ビデオ> 。
モデルさんに聞いてみてください。作物は 450 M パラメータの視覚言語モデルに移行します —
Minus-v0.1 — で実行中
WebGPU 経由の GPU (transformers.js + ONNX ランタイム Web、量子化
~430 MB)。作物ごとに 1 つの質問に答えます。「これは広告ですか?」 —
そして、Yes/No ロジットにより、調整された p(ad) が得られます。
段階的なしきい値で決定します。広告コンテキスト内の要素 (広告 iframe、
スロットコンテナ) p ≥ 0.60 でブロックします。裸の画像は p ≥ 0.88 をクリアする必要があります
標準的な広告形状であること (正方形に近い商品/エディトリアル タイルは決して使用されません)
ブロックされました）。動画の判定には 2 つの連続した広告フレーム (ヒステリシス) が必要です。
カバーされたプレーヤーは再検証を続けるため、広告が終了した瞬間にカバーされます。
カバーしてください、壊れないでください。ブロックされた要素はフラッシュカード オーバーレイを取得します (単語 →
翻訳→例文）。ページ レイアウトは変更されていません。オーバーレイトラック
それらの要素をモーダルのために脇に置き、かすかな ✕ (明らかに、
一時的な ↩ 再ブロックの取り消し）、およびオプトイン ⚑ 広告レポート ボタンではありません。
学ぶ。ブロックされた広告に表示されるすべての単語が、組み込みの広告にフィードされます。
間隔をあけた繰り返しレビュー。
content.js（すべてのフレーム）background.js（SW）offscreen.js（ウィンドウ）
─────────────────────────────
上のフレーム: 静的スキャン → CaptureVisibleTab → Minus-v0.1 (transformers.js,
+ iframe モーション サンプラー (レート制限クロップ) WebGPU) または SigLIP2 (ORT)
すべてのフレーム: <ビデオ> → ルート分類 + バッジ → クロップごとの P(広告)
サンプリング (ヒステリシス) + タブごとのカウンター │
↑ │
フラッシュカード オーバーレイ ←───────────────────┘
アクティブなタブのみがスキャンされます (バックグラウンド タブのキャプチャでは、
間違ったピクセル)、sc

何も覆われていない場合、ans は合体して自己停止します。
エンジンは、WebGPU デバイスが失われても、それ自体を再構築することで生き残ります。
nypost.com — リーダーボード + サイドバー広告をカバー (「ブエノス ディアス」、「エル アグア」)
cnet.com — ディスプレイ広告については記事の途中で取り上げています
ポップアップ: カウンタ、一時停止、広告タイプの切り替え、しきい値スライダー、エンジン ピッカー
レビュー: ブロックされた広告の単語が間隔をあけて繰り返されるデッキになる
オプション: ブロックアクション + フラッシュカード言語 (ライブプレビュー付き)、学習統計、
モデルごとの決定ゲートを備えたエンジン ピッカー。
デフォルトのエンジンは、Hugging Face の Minus-v0.1 です。
— LiquidAI/LFM2.5-VL-450M の微調整
マイナスからのストリーミング TV キャプチャに関する 28 回の反復キャンペーンにわたってトレーニングされました
デバイスと Web ディスプレイ広告および約 10,000 のマイニングされたハード ネガ (製品写真、
編集コンテンツ、UI 要素、同意バナー、チャット ウィジェット)。の数字
出荷ゲート:
完全なモデル カード — トレーニング レシピ、ONNX 量子化、からの使用法
transformers.js — 「Hugging Face」ページにあります。
ディスプレイ/バナー広告 - 静的 <img> 、広告 <iframe> 、および広告の形
切り取られたタブのスクリーンショットから分類されたスロット。
ビデオ広告 — 検出器はすべてのフレーム ( all_frames ) で実行されるため、
プレーヤー内 <video> 広告をカバーします (2 つの判定で約 2.5 秒ごとに再サンプリングされます)
ヒステリシス（マイナス デバイスなど）とクロスオリジン iframe 内のビデオ広告
見ることができないプレーヤー (トップフレームのモーション サンプラー)。本物のカバーを確認済み
YouTube、USA Today、アルジャジーラの広告。カバーされたプレイヤーが、
広告が終了する瞬間 (モデルはティックごとにライブ フレームを再読み込みします)。本当に
DRM で保護されたプレーヤー (一部の Vevo ミュージック ビデオ) が読めない黒色でレンダリングされる
ハードウェア オーバーレイ — プレロールは読み取れないため、カバーされていませんが、
警備員は、コンテンツが誤ってカバーされないようにします。
ポップアップ/ポップアンダー広告タブ - アグレッシブ

サイト（マンガ/ストリームリーダー）ハイジャック
リンクされていないページ領域をクリックすると、ページ全体が広告であるタブが表示されます
着陸。ポップアップ ガードは、ハイジャックされたクリック→新しいタブのパターンに気づきます。
モデルにページ自体について質問し、自信のある広告ランディングをカバーします。
明示的な選択: タブを閉じるかページを表示 (タブは自動的には閉じません)
実際のリンクから開かれたものは決して触れられません)。
オーバーレイはページ UI に譲ります — モーダル/ライトボックスが覆われた広告の上で開いた場合、
フラッシュカードは脇に移動します (クリップパスの穴) ので、ダイアログはクリック可能なままになります。
ポップアップ（アイコンを左クリック）
このページでブロックされた広告カウンター (ツールバーのバッジを反映し、M アイコンは
静止時は青、ブロック中は赤）に加えて、これまでのブロック数も表示されます。
広告をブロック (すべてのサイト) — マスターのオン/オフ。オフにすると完全にアンロードされます。
エンジンからモデルを削除します (GPU メモリを解放します)。
このサイトでブロック — ホスト名ごとに切り替えます。
10 分 / 30 分 / 1 時間一時停止 — スヌーズしてブロックし、タイマーで自動的に再開します
(一時停止中のボタンはカウントダウン + 再開に置き換わります)。
広告タイプ - 独立した動画広告とディスプレイ広告の切り替え（適用済み）
ライブ、リロードなし)。
📚 フラッシュカードを復習する — 復習ページを開きます。ボタンには数が表示されます
広告で出会った言葉がただ過ぎ去ってしまうことがないように、カードの期限が設定されています (下記を参照)。
広告しきい値スライダー (ライブ値、ジャンクを保持できない)、エンジン ピッカー
スイッチフィードバック、ライブエンジンステータス。両方の広告を切り替えると警告が表示されます
マスタートグルがまだオンの間、タイプオフになります。
匿名の広告スナップショットを投稿します — オプトイン、デフォルトではオフです (「プライバシー」を参照)。
オンの場合、各オーバーレイには広告ボタンではなく ⚑ が表示されます (カーソルを合わせると表示されます /
focus) を使用して誤検知を報告し、スナップショットにレビュー用のフラグを立てます。
⚙ オプション — 完全なオプション ページが開きます。
ブロックしながら学習します (間隔をあけて繰り返します)
フラッシュカードは単なる飾りではありません。ブロックされたメッセージにマイナスが表示されるすべての単語

広告
は静かに保存され、レビューページがそれらを本物に変えます。
間隔をあけた繰り返しデッキ (Anki のような組み込み):
各カードをもう一度/良い/簡単に採点します (キーボード: スペースで反転、1/2/3 から 1/2/3 まで)
グレード）。 SM-2 スタイルのスケジューラは、知っている単語をさらに遠くに配置して表示します。
見逃したものを取り戻します。新しいカードは 1 日に数枚ずつ導入されます。
進行状況ストリップは、確認、学習、学習、レビューを追跡し、フィルタリング可能
言語ごとに表示されるので、見たはずの広告が語彙として保存されます。
マシンからは何も出ません。進行状況はローカル ストレージに保存されます。学び
オプション ページのセクションには統計情報が表示され、進行状況をリセットできます。
オプション ページ (⚙ ポップアップ内、またはアイコンを右クリック → [オプション])
ポップアップの全ページのスーパーセットと、ここにのみ存在する設定:
ブロック アクション - ブロックされた広告の上に表示される内容:
言語フラッシュカード (デフォルト) スペイン語、フランス語、ドイツ語、イタリア語、
ポルトガル語、日本語、またはギリシャ語 — 言語ごとに 500 語のデッキ
(デッキ/*.json)、ブロックされたすべての広告が語彙カードになります (単語 →
翻訳→例文）、ライブプレビュー付き。
最小限 — 広告がブロックされたことを示すだけの静かな暗いカード。
各カードに表示されるモデル信頼性タグ (「広告 97%」) を切り替えます。
変更はすでに画面上にあるオーバーレイに適用されます。リロードは必要ありません。
無効なサイト リスト - サイトごとのブロック リストをテキストとして編集します (
ポップアップのサイトごとのトグルは同じリストを書き込みます)。
広告しきい値スライダー (ポップアップと同じ設定) とエンジン ピッカー
各モデルの決定ゲート、ライブエンジンステータス。
エンジンのリストはパッケージ化されたモデル ( models/index.json ) から生成されるため、
新しいモデル ディレクトリをドロップすると、コードを変更せずに選択できるようになります。スイッチイン
ポップアップ (エンジンをリロードします)。
エンジンごとの決定しきい値 (v0.3.1 以降): カタログ エントリには次の値を含めることができます。
"しきい値": { "ctx"

: …, "bare": … } — 信頼度バー content.js が適用されます
広告コンテキスト要素 (iframe / 広告スロット) と標準サイズの裸の画像。
各モデルは、先行モデルの操作点を継承するのではなく、独自の操作点を出荷します。
現在のデフォルト (Minus v0.1) は 0.60 / 0.88 で動作します。モデルのスコアが
配信により広告がきれいに分離され、ゲートが緩くなりリコールを無料で購入できるようになります。
Iter 24 時代の出荷率は 0.35 / 0.75 で、60 分間のライブ A/B は最大 33 % 増加しました。
広告は変更されずに最大 90% の精度でカバーされます。
デフォルトの LFM エンジンは WebGPU のみです。量子化されたグラフは
GatherBlockQuantized オペレーション。ONNX-Runtime の WASM バックエンドは実行できません。 WebGPUは
Chrome/Edge 121 以降ではデフォルトでオンになります。利用できない場合は、ポップアップ ステータスに
「WebGPU が必要です — 有効にしてリロードしてください」というメッセージをクリアします。 WebGPU の負荷とウォームアップ
タイムアウト制限があるため、不安定な GPU ドライバーがエンジンを妨害することはできません。
再トレーニングされた SigLIP2-Lite WASM フォールバックはプロトタイプ化されたため、WebGPU は使用されませんでした
マシンはまだ分類器を実行できますが、棚上げされました - これに対する WASM 推論
ViT は画像あたり約 7 秒で、実際のエクスペリエンスには遅すぎます (int8 quant、これは
縮小すると、このモデルの広告想起が折りたたまれます)。仕事は
siglip2-wasm-フォールバック
枝。
拡張機能はまだ Chrome ウェブストアにないため、解凍してロードします。
リリースジップ。これは 2 分間の 4 ステップのプロセスです。
最新のminus-extension-vX.Y.Z.zip (~400 MB - バンドルされています) をダウンロードします。
ビジョンモデル)から
をリリースします。
解凍してください。以下で、minus-extension-vX.Y.Z/ のようなフォルダーを取得します。
その中にあるmanifest.json — そのフォルダーが拡張子です

[切り捨てられた]

## Original Extract

Contribute to garagehq/Minus-chrome-extension development by creating an account on GitHub.

GitHub - garagehq/Minus-chrome-extension · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
garagehq
/
Minus-chrome-extension
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
82 Commits 82 Commits docs docs extension extension server server src src tests tests .gitignore .gitignore ACCEPTANCE.md ACCEPTANCE.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md build.mjs build.mjs build_model_index.mjs build_model_index.mjs docs_perf_soak.json docs_perf_soak.json gen_icons.mjs gen_icons.mjs package-lock.json package-lock.json package.json package.json package.mjs package.mjs popup_preview.png popup_preview.png View all files Repository files navigation
Minus — vision ad blocker (Chrome extension)
The browser cousin of the minus HDMI
device: instead of matching filter lists, it looks at page elements with a
vision-language model running entirely inside your browser and covers the
ones that are ads with language flashcards — so every blocked ad teaches you a
word instead.
Nothing leaves your machine — no server, no filter-list subscriptions, no
telemetry. The model sees what you see.
Real page, real ads: the banner and sidebar ad slots on theverge.com covered by
flashcards ("la puerta — the door", "cuánto — how much"), each with the model's
confidence tag and a faint ✕ to reveal the ad. The article is untouched.
A classic ad blocker matches URLs against filter lists. Minus classifies
pixels — so first-party ads, sponsored tiles, native "chum-box" placements
and streaming video ads that lists can't describe still get caught, and nothing
breaks when ad-tech rotates domains.
Find candidates. The content script (running in every frame) collects
elements that could be ads: <img> s and <iframe> s, plus containers whose
id/class smells like an ad slot. Obvious non-ads are filtered structurally —
consent banners, video players (they belong to the video path), unloaded
image placeholders, content-column-shaped boxes.
Look at them. For each candidate, Minus grabs actual pixels: a cropped
tab screenshot for display elements (rate-limited, active tab only), or a
direct frame read for <video> .
Ask the model. Crops go to a 450 M-parameter vision-language model —
Minus-v0.1 — running on
your GPU via WebGPU (transformers.js + ONNX Runtime Web, quantized to
~430 MB). It answers one question per crop: "Is this an advertisement?" —
and the Yes/No logits give a calibrated p(ad) .
Decide with tiered thresholds. An element in ad context (an ad iframe,
a slot container) blocks at p ≥ 0.60 ; a bare image must clear p ≥ 0.88
and be a standard ad shape (near-square product/editorial tiles are never
blocked). Video verdicts need two consecutive ad frames (hysteresis), and a
covered player keeps re-verifying so it uncovers the moment the ad ends .
Cover, don't break. Blocked elements get a flashcard overlay (word →
translation → example sentence). The page layout is untouched; overlays track
their element, step aside for modals, and carry a faint ✕ (reveal, with a
transient ↩ re-block undo) and an opt-in ⚑ not an ad report button.
Learn. Every word shown on a blocked ad feeds the built-in
spaced-repetition review .
content.js (all frames) background.js (SW) offscreen.js (window)
─────────────────────── ───────────────── ─────────────────────
top frame: static scan → captureVisibleTab → Minus-v0.1 (transformers.js,
+ iframe motion sampler (rate-limited crops) WebGPU) or SigLIP2 (ORT)
every frame: <video> → route classify + badge → P(ad) per crop
sampling (hysteresis) + per-tab counter │
↑ │
flashcard overlays ←────────────────────────────────────┘
Only the active tab ever scans (a background tab capturing would read the
wrong pixels), scans coalesce and self-suspend when nothing is covered, and the
engine survives WebGPU device loss by rebuilding itself.
nypost.com — leaderboard + sidebar ads covered ("buenos días", "el agua")
cnet.com — display ad covered mid-article
The popup: counters, pause, ad-type toggles, threshold slider, engine picker
Review: words from blocked ads become a spaced-repetition deck
Options: block action + flashcard language (with live preview), learning stats,
engine picker with per-model decision gates.
The default engine is Minus-v0.1 on Hugging Face
— a fine-tune of LiquidAI/LFM2.5-VL-450M
trained over a 28-iteration campaign on streaming-TV captures from the minus
device plus web display ads and ~10 k mined hard negatives (product photography,
editorial content, UI elements, consent banners, chat widgets). Numbers at the
shipping gates:
The full model card — training recipe, ONNX quantization, usage from
transformers.js — is on the Hugging Face page.
Display / banner ads — static <img> , ad <iframe> s, and ad-shaped
slots, classified from a cropped tab screenshot.
Video ads — the detector runs in every frame ( all_frames ), so it
covers in-player <video> ads (re-sampled every ~2.5 s with 2-verdict
hysteresis, like the minus device) and video ads inside cross-origin iframe
players it can't see into (a top-frame motion sampler). Verified covering real
ads on YouTube, USA Today, and Al Jazeera. A covered player uncovers the
moment the ad ends (the model re-reads the live frame each tick). Genuinely
DRM-protected players (some Vevo music videos) render as an unreadable black
hardware overlay — their pre-rolls can't be read , so they aren't covered, but
the guard makes sure their content is never mistakenly covered either.
Popup / popunder ad tabs — aggressive sites (manga/stream readers) hijack
clicks on non-link page areas to spawn tabs whose entire page is an ad
landing. The popup guard notices the hijacked-click → new-tab pattern,
asks the model about the page itself, and covers confident ad landings with an
explicit choice: Close tab or Show page (never auto-closes; tabs
opened from real links are never touched).
Overlays yield to page UI — if a modal/lightbox opens over a covered ad,
the flashcard steps aside (clip-path hole) so the dialog stays clickable.
The popup (left-click the icon)
Ads blocked on this page counter (mirrors the toolbar badge; the M icon is
blue at rest, red while blocking) plus an all-time blocked tally.
Block ads (all sites) — master on/off. Turning it off fully unloads the
model from the engine (frees GPU memory).
Block on this site — per-hostname toggle.
Pause 10m / 30m / 1h — snooze blocking, then it auto-resumes on a timer
(a countdown + Resume now replace the buttons while paused).
Ad types — independent Video ads and Display ads toggles (applied
live, no reload).
📚 Review flashcards — opens the review page; the button shows how many
cards are due so the words you meet on ads don't just flash by (see below).
Ad threshold slider (live value, can't hold junk), Engine picker with
switch feedback, live engine status. A warning appears if you switch both ad
types off while the master toggle is still on.
Contribute anonymous ad snapshots — opt-in, off by default (see Privacy).
When it's on, each overlay gets a ⚑ not an ad button (revealed on hover /
focus) to report a false positive so the snapshot is flagged for review.
⚙ options — opens the full options page.
Learn as you block (spaced repetition)
The flashcards aren't just decoration. Every word Minus shows on a blocked ad
is quietly saved , and the Review page turns them into a real
spaced-repetition deck (like Anki, built in):
Grade each card Again / Good / Easy (keyboard: space to flip, 1/2/3 to
grade). An SM-2-style scheduler spaces words you know further out and brings
back the ones you miss; new cards are introduced a handful at a time per day.
A progress strip tracks seen · learning · learned · to review , filterable
by language, so the ads you would have watched become vocabulary you keep.
Nothing leaves your machine — progress lives in local storage; the Learning
section of the options page shows your stats and can reset progress.
The options page (⚙ in the popup, or right-click the icon → Options)
A full-page superset of the popup, plus settings that only live here:
Block action — what appears over a blocked ad:
Language flashcards (default) in Spanish, French, German, Italian,
Portuguese, Japanese, or Greek — a 500-word deck per language
( decks/*.json ), every blocked ad becomes a vocabulary card (word →
translation → example sentence), with a live preview.
Minimal — a quiet dark card that just says the ad was blocked.
Toggle the model-confidence tag ("ad 97%") shown on each card.
Changes apply to overlays already on screen — no reload.
Disabled-sites list — edit the per-site blocking list as text (the
popup's per-site toggle writes the same list).
Ad threshold slider (same setting as the popup) and the Engine picker
with each model's decision gates, live engine status.
The engine list is generated from the packaged models ( models/index.json ), so
dropping in a new model dir makes it selectable with no code changes. Switch in
the popup (reloads the engine).
Per-engine decision thresholds (since v0.3.1): a catalog entry can carry
"thresholds": { "ctx": …, "bare": … } — the confidence bars content.js applies
to ad-context elements (iframes / ad-slots) and bare standard-size images.
Each model ships its own operating point instead of inheriting a predecessor's;
the current default (Minus v0.1) runs 0.60 / 0.88. When a model's score
distribution separates ads cleanly, looser gates buy recall for free — the
Iter 24 era shipped 0.35 / 0.75 and a 60-minute live A/B measured ~33 % more
ads covered at unchanged ~90 % precision .
The default LFM engines are WebGPU-only — their quantized graph uses the
GatherBlockQuantized op, which ONNX-Runtime's WASM backend can't run. WebGPU is
on by default in Chrome/Edge 121+; if it's unavailable the popup status shows a
clear "needs WebGPU — enable it and reload" message. The WebGPU load and warm-up
are timeout-bounded, so a flaky GPU driver can't wedge the engine.
A retrained SigLIP2-Lite WASM fallback was prototyped so no-WebGPU
machines could still run a classifier, but shelved — WASM inference for this
ViT is ~7 s/image, too slow for a real experience (int8 quant, which would
shrink it, collapses this model's ad recall). The work lives on the
siglip2-wasm-fallback
branch.
The extension isn't on the Chrome Web Store yet, so you load it unpacked from
the release zip. It's a two-minute, four-step process:
Download the latest minus-extension-vX.Y.Z.zip (~400 MB — it bundles
the vision model) from
Releases .
Unzip it. You get a folder like minus-extension-vX.Y.Z/ with
manifest.json right inside it — that folder is the exten

[truncated]
