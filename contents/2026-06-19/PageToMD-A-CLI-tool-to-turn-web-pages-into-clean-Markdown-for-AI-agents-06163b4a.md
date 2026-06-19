---
source: "https://github.com/gs202/PageToMD"
hn_url: "https://news.ycombinator.com/item?id=48596454"
title: "PageToMD – A CLI tool to turn web pages into clean Markdown for AI agents"
article_title: "GitHub - gs202/PageToMD: Convert web pages to clean Markdown · GitHub"
author: "gs202"
captured_at: "2026-06-19T10:35:19Z"
capture_tool: "hn-digest"
hn_id: 48596454
score: 1
comments: 0
posted_at: "2026-06-19T09:01:49Z"
tags:
  - hacker-news
  - translated
---

# PageToMD – A CLI tool to turn web pages into clean Markdown for AI agents

- HN: [48596454](https://news.ycombinator.com/item?id=48596454)
- Source: [github.com](https://github.com/gs202/PageToMD)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T09:01:49Z

## Translation

タイトル: PageToMD – Web ページを AI エージェント用のクリーンな Markdown に変換する CLI ツール
記事のタイトル: GitHub - gs202/PageToMD: Web ページをクリーンな Markdown に変換する · GitHub
説明: Web ページをクリーンな Markdown に変換します。 GitHub でアカウントを作成して、gs202/PageToMD の開発に貢献してください。

記事本文:
GitHub - gs202/PageToMD: Web ページをクリーンな Markdown に変換する · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
gs202
/
ページからMDまで
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

アクションメニュー フォルダーとファイル
92 コミット 92 コミット .github .github スクリプト スクリプト src/ pagetomd src/ pagetomd テスト テスト .gitignore .gitignore .markdownlint.jsonc .markdownlint.jsonc .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .secrets.baseline .secrets.baseline CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリファイルナビゲーション
Frontmatter を使用して、あらゆる Web ページ URL をクリーンな LLM 対応の Markdown に変換します。
デフォルトで AI 対応。出力は NFC で正規化された UTF-8、単一の H1、単調な見出し階層、ゼロ幅のジャンクや追跡クロムはなく、ベクトル ストアまたは LLM プロンプトに直接ドロップされます。
完全に忠実なメタデータ。すべてのファイルには、正規 URL、リダイレクト後の最終 URL、タイトル、作成者、日付、説明、サイト名、言語、ツール ID を含む YAML フロントマター ブロックが付属しています。 「この Markdown はどこから来たの?」ということはもうありません。
静的に高速で、必要に応じて JS 対応。デフォルトの httpx フェッチャーは 1 秒未満です。オプトイン Playwright extra (または --fetcher auto ) は、他のユーザーのインストールを肥大化させることなく SPA シェルを処理します。
安定したスクリプト可能な CLI。タイパーが構築した完全な env-var 優先順位 ( PAGETOMD_* )、安定した終了コード ( 0 / 2 / 3 / 4 / 5 / 64 / 130 )、構造化ログ ( --log-json )、およびバイト確定出力用の --no-fetched-at スイッチ。
Pandoc またはcurl + sed ではありません。 pandoc はフェッチせず、定型文を削除せず、フロントマターも発行しません。手巻きカール | カールhtml2md パイプラインは、抽出、mojibake 処理、robots.txt、リダイレクト キャップ、およびアトミック書き込みを再発明します。 pagetomd はパイプライン全体の 1 つのコマンドです。
pipx を使用する場合 (CLI での使用を推奨)
pipx インストール pagetomd
# オプション: JS レンダリングを有効にする

SPA向け
pipx inject pagetomd playwright && playwright install chrome
紫外線あり
UVツールのインストールpagetomd
# オプション: SPA の JS レンダリングを有効にする
uv ツールのインストール ' pagetomd[playwright] ' && playwright のクロムのインストール
インストールなし（UV実行）
# コア — インストールは必要ありません
uv run --with pagetomd pagetomd https://example.com
# SPA / JS-heavy ページ用の Playwright を使用 (最初に Chromium を 1 回インストールします)
uv run --with playwright playwright install chrome
uv run --with ' pagetomd[playwright] ' pagetomd https://example.com --fetcher auto
ピップあり
pip install pagetomd # core
pip install ' pagetomd[playwright] ' # + SPA サポート
クイックスタート
# デフォルト: ページのタイトルから出力ファイル名を取得します
pagetomd https://example.com/blog/post
# 標準出力へのストリーム (LLM へのパイプなど)
pagetomd https://example.com/blog/post -o -
# 確定的な出力 (fetched_at を省略 - スナップショット テスト/RAG インジェストに適しています)
pagetomd https://example.com/blog/post --no-fetched-at -o post.md
# SPA ページを自動検出し、ヘッドレス Chromium にフォールバックする
pagetomd https://my-spa.example.com -o - --fetcher auto
料理本
-o - マークダウンを標準出力に書き込みます。すべてのログは標準エラー出力に送られるため、ストリームは安全にパイプできます。
pagetomd https://example.com/blog/post -o - | llm 「この記事を 5 つの箇条書きで要約します」
ファイルから一括変換する
読み取り中 -r URL ;する
pagetomd " $url "
完了 < urls.txt
成功した各変換は 0 で終了します。ゼロ以外の出口があるとループから抜け出します
実行中ですが、標準エラー出力に表示されます (以下の終了コードを参照)。
ドキュメント サイト全体をクロールする
--crawl を使用して、シード URL の下にあるすべてのリンクされたサブページを検出し、書き込みます
ページごとに 1 つの .md ファイルを出力ディレクトリに配置します。
pagetomd " https://docs.example.com/guide/ " \
--crawl --crawl-深さ 2 \
--fetcher auto --no- respect-robots \
-o ./docs-output/
範囲: シード

はそれ自体のサブツリーのルートとして扱われます。リンクのみ
シードの下に存在する URL が追跡されます。兄弟や両親、そして
外部サイトはスキップされます。の種のために
https://docs.example.com/guide/intro スコープ内のプレフィックスは次のとおりです
https://docs.example.com/guide/intro/ — 末尾のスラッシュを渡します。
シード (または /guide/ のような「ディレクトリ」URL を使用) を使用して、クロールの範囲を指定します。
レベルが高くなります。
出力構造: ディスク上のレイアウトは、次の URL 階層を反映しています。
シードなので、同じ最終 URL セグメントを持つ 2 つのページが異なる下にあります。
親は衝突しません:
各パス セグメントは個別にスラッジ化され、Windows で予約されたデバイス
名前 ( CON 、 PRN 、…) はセグメントごとにエスケープされます。
--crawl- Depth N — シードからの BFS ホップ制限 (デフォルト: 1 )。
--crawl- Depth 10 を、自然に深さ 3 で終了するサイトに対して単純に実行します。
キューが空になると停止します。何も無駄なことはありません。
--overwrite — 既存の .md ファイルを置き換えます (デフォルト: スキップ)。クロールの終わりに、
3 つのリストが標準エラー出力に出力されます: ファイルがすでに存在するためスキップされたページ、
コンテンツを抽出できなかったページ (認証ウォール、薄いナビゲーション スタブ)、およびページ
フェッチまたは変換エラーで失敗した場合 - 各カテゴリを処理できるようになります
適当に。
他のすべてのフラグ ( --fetcher 、 --no-verify-ssl 、 --user-agent 、
--retries , …) クロール内のすべてのページに適用されます。 --再試行の優先順位
429/503 応答の Retry-After ヘッダー (1 回あたり 5 分に制限)
試み）。
単一のフェッチャー コンテキストがクロール全体で再利用されるため、ブラウザ
バックエンドはページごとに Chromium を再起動しません。
pagetomd には URL を Markdown に変換する 4 つの方法があります。状況に合ったものを選択してください。
httpx (デフォルト) — プレーンな HTTP GET。ほとんどのページでは 1 秒未満で、指数バックオフで再試行を処理し、429/503 で Retry-After を尊重し、 robots.txt を強制し、 <meta http-equiv="ref に従います。

resh"> リダイレクト。JavaScript は実行されません。サーバーが空の <div id="root"></div> シェルを送信した場合、取得できるのはそれだけです。
playwright — ヘッドレス Chromium でページをレンダリングし、ネットワークがアイドル状態になるまで待機してから、ライブ DOM (シャドウ ルートを含む) をシリアル化します。ページが SPA であることがわかっている場合にこれを使用します。オプションの Playwright extra ( pip install 'pagetomd[playwright]' ) と 1 回限りの Playwright install chromium が必要です。 httpx より遅くて重いですが、JS フレームワークの背後にあるコンテンツを取得する唯一の方法です。
auto — 最初に httpx でフェッチし、次に結果を検査します。 <body> テキストが 200 文字未満で、HTML に SPA マーカー ( data-reactroot 、 <div id="__next"> 、「JavaScript を有効にする必要がある」 noscript タグなど) が含まれている場合は、Playwright で再フェッチします。 httpx が空ではないように見える HTML を返したにもかかわらず、エクストラクターがまだコンテンツを取得できなかった場合、2 番目のセーフティ ネットが起動します。その場合、Playwright も攻撃を受けます。 Playwright が使用できない場合、ページはクロールの概要でハード障害ではなく「空」としてカウントされます。見慣れない URL が表示された場合に最適な選択です。
特定の URL (または while 読み取りループを介してパイプされた短いリスト) がある場合は、デフォルトのシングルページ モードを使用します。 URL プレフィックスの下にすべてのページが必要な場合は --crawl を使用します。リンクが自動的に検出され、重複が排除され、ディスク上の URL 階層がミラーリングされ、単一のフェッチャー コンテキストが再利用されるため、Playwright がページごとに Chromium を再起動することはありません。完全なフラグ セットについては、クロール クックブックのレシピを参照してください。
pagetomd http://127.0.0.1:8765/blog.html --no-fetched-at -o - blog.html フィクスチャに対して実行すると、次のように出力されます (最初の ~15 行が表示されます)。
---
URL : http://127.0.0.1:8765/blog.html
最終URL : http://127.0.0.1:8765/blog.html
title : なぜビルドシステムをRustで書き直したのか
著者：ジェーン・ドウ
日付: ' 2024-08

-14 '
説明 : モノリポジトリ ビルド パイプラインを Python から Rust に移行することと、その過程で学んだことを振り返ります。
site_name : サンプルエンジニアリングブログ
言語 : 英語
ツール：pagetomd
ツールバージョン : 0.4.0
---
# ビルドシステムを Rust で書き直した理由
3 年前、私たちのモノリポ ビルド パイプラインは、シェル スクリプトと祈りとともに保持された広大な Python アプリケーションでした。 ...
fetched_at が有効になっている場合 (デフォルト)、追加の fetched_at: '2026-06-15T12:34:56Z' 行がフロントマターに含まれます。値が検出できないフィールド ( language 、 author など) は YAML から省略されます。
簡潔な概要 — 完全なリストについては、pagetomd --help を参照してください。
すべてのフラグには、同等の PAGETOMD_<UPPER_NAME> があります。たとえば:
PAGETOMD_TIMEOUT=60 PAGETOMD_FETCHER=auto pagetomd https://example.com
CLI フラグは常に環境変数をオーバーライドします。 env vars は組み込みのデフォルトをオーバーライドします。
コード
意味
0
成功。
1
予期しない内部エラー。
2
フェッチの失敗 (DNS、HTTP、robots.txt、リダイレクト キャップ)。
3
抽出または変換の失敗 (空の本文、不正な形式の HTML)。
4
出力の書き込み失敗 (アクセス許可、ディスク、アトミックリネームの衝突)。
5
オプションの依存関係がありません (例: Playwright がインストールされていません)。
64
使用法または構成エラー (不正なフラグ、無効な値)。
130
ユーザー (Ctrl-C) によって中断されました。
仕組み
1 つの段落とパイプラインの図:
URL ──► フェッチャー ──► エクストラクター ──► コンバーター ──► ポストプロセス ──► ライター
(httpx / (BS4 クリーン (マークダウン (NFC、見出し (アトミック
playwright) + trafilatura) + GFM テーブル) 階層、ファイル +
URL 絶対) YAML)
フェッチャー (デフォルトでは httpx、SPA の場合は Playwright) は、再試行と robots.txt 強制を使用してページをダウンロードします。エクストラクターは、BeautifulSoup の事前クリーン パス (スクリプト/スタイル/ナビゲーション/広告のストリップ) を実行し、クリーンになったツリーを t に渡します。

rafilatura を使用してメイン コンテンツを特定し、メタデータを収集します。コンバーターは、カスタマイズされた markdownify サブクラス (ATX 見出し、言語ヒントを含むフェンスされたコード ブロック、ワイド テーブル フォールバックを含む GFM テーブル) を介して、生き残った HTML を Markdown にレンダリングします。ポストプロセッサは、AI 対応コントラクト (NFC、ゼロ幅ストリップ、単調見出し階層、絶対 URL) を強制します。ライターは YAML フロントマター ブロックを先頭に追加し、アトミックに書き込みます (または stdout にストリームします)。
pagetomd はパブリック URL のみのツールです。デフォルトでは、プライベート、ループバック、リンクローカル、マルチキャスト、予約、またはクラウド メタデータ アドレスの取得を拒否します。また、それをオーバーライドするフラグはありません。出力ファイルを、フェッチ元の URL と同じ機密性を持つものとして扱います。
CI は、4 つの重要なモジュール (extractor 、 Converter 、 Writer 、 postprocess ) に対して、プロジェクト全体のテスト カバレッジ フロア 85% とモジュールごとのフロア 90% (ラインとブランチの合計) の両方を強制します。これら 4 社は AI 対応契約を締結しているため、最も厳しい適用範囲が設けられています。
git clone https://github.com/gs202/PageToMD.git
cd ページトムド
uv sync --extra dev --extra playwright
プリコミットインストール
UV で pytest を実行
完全な貢献者のワークフローについては、CONTRIBUTING.md を参照してください。
ビジネス ソース ライセンス 1.1 — ソースが利用可能、非営利使用は無料。コンバージョン

[切り捨てられた]

## Original Extract

Convert web pages to clean Markdown. Contribute to gs202/PageToMD development by creating an account on GitHub.

GitHub - gs202/PageToMD: Convert web pages to clean Markdown · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gs202
/
PageToMD
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
92 Commits 92 Commits .github .github scripts scripts src/ pagetomd src/ pagetomd tests tests .gitignore .gitignore .markdownlint.jsonc .markdownlint.jsonc .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .secrets.baseline .secrets.baseline CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Convert any webpage URL into clean, LLM-ready Markdown with frontmatter.
AI-ready by default. Output is NFC-normalised UTF-8, single H1, monotonic heading hierarchy, no zero-width junk, no tracking chrome — drops straight into a vector store or LLM prompt.
Full-fidelity metadata. Every file ships with a YAML frontmatter block containing canonical URL, final URL after redirects, title, author, date, description, site name, language, and tool identity. No more "where did this Markdown come from?".
Static fast, JS-capable when needed. Default httpx fetcher is sub-second; opt-in playwright extra (or --fetcher auto ) handles SPA shells without bloating the install for everyone else.
Stable, scriptable CLI. Typer-built, full env-var precedence ( PAGETOMD_* ), stable exit codes ( 0 / 2 / 3 / 4 / 5 / 64 / 130 ), structured logs ( --log-json ), and a --no-fetched-at switch for byte-deterministic output.
Not pandoc or curl + sed . pandoc doesn't fetch, doesn't strip boilerplate, and doesn't emit frontmatter. Hand-rolled curl | html2md pipelines re-invent extraction, mojibake handling, robots.txt, redirect caps, and atomic writes. pagetomd is one command for the whole pipeline.
With pipx (recommended for CLI use)
pipx install pagetomd
# optional: enable JS rendering for SPAs
pipx inject pagetomd playwright && playwright install chromium
With uv
uv tool install pagetomd
# optional: enable JS rendering for SPAs
uv tool install ' pagetomd[playwright] ' && playwright install chromium
Without installing (uv run)
# Core — no install required
uv run --with pagetomd pagetomd https://example.com
# With Playwright for SPA / JS-heavy pages (install Chromium once first)
uv run --with playwright playwright install chromium
uv run --with ' pagetomd[playwright] ' pagetomd https://example.com --fetcher auto
With pip
pip install pagetomd # core
pip install ' pagetomd[playwright] ' # + SPA support
Quick start
# Default: derives output filename from the page title
pagetomd https://example.com/blog/post
# Stream to stdout (pipe into LLMs, etc.)
pagetomd https://example.com/blog/post -o -
# Deterministic output (omits fetched_at — good for snapshot tests / RAG ingestion)
pagetomd https://example.com/blog/post --no-fetched-at -o post.md
# Auto-detect SPA pages and fall back to headless Chromium
pagetomd https://my-spa.example.com -o - --fetcher auto
Cookbook
-o - writes the Markdown to stdout. All logs go to stderr, so the stream is safe to pipe:
pagetomd https://example.com/blog/post -o - | llm " summarise this article in five bullet points "
Batch-convert from a file
while read -r url ; do
pagetomd " $url "
done < urls.txt
Each successful conversion exits 0 ; any non-zero exit leaves the loop
running but is visible in stderr (see Exit codes below).
Crawl an entire documentation site
Use --crawl to discover every linked sub-page under a seed URL and write
one .md file per page into an output directory:
pagetomd " https://docs.example.com/guide/ " \
--crawl --crawl-depth 2 \
--fetcher auto --no-respect-robots \
-o ./docs-output/
Scope: The seed is treated as the root of its own subtree. Only links
whose URL lives under the seed are followed; siblings, parents, and
external sites are skipped. For a seed of
https://docs.example.com/guide/intro the in-scope prefix is
https://docs.example.com/guide/intro/ — pass a trailing slash on the
seed (or use a "directory" URL like /guide/ ) to scope the crawl one
level higher.
Output structure: The on-disk layout mirrors the URL hierarchy under
the seed, so two pages with the same final URL segment under different
parents do not collide:
Each path segment is slugified independently, and Windows-reserved device
names ( CON , PRN , …) are escaped per segment.
--crawl-depth N — BFS hop limit from the seed (default: 1 ).
--crawl-depth 10 against a site that naturally ends at depth 3 simply
stops when the queue empties; nothing is wasted.
--overwrite — replace existing .md files (default: skip). At the end of a crawl,
three lists are printed to stderr: pages skipped because the file already exists,
pages where no content could be extracted (auth walls, thin nav stubs), and pages
that failed with a fetch or conversion error — so you can handle each category
appropriately.
All other flags ( --fetcher , --no-verify-ssl , --user-agent ,
--retries , …) apply to every page in the crawl. --retries honours
Retry-After headers on 429/503 responses (capped at 5 minutes per
attempt).
A single fetcher context is reused across the whole crawl, so browser
backends do not relaunch Chromium per page.
pagetomd has four ways to turn URLs into Markdown. Pick the one that matches your situation:
httpx (default) — A plain HTTP GET. Sub-second for most pages, handles retries with exponential backoff, honours Retry-After on 429/503, enforces robots.txt , and follows <meta http-equiv="refresh"> redirects. No JavaScript execution — if the server sends an empty <div id="root"></div> shell, that's all you get.
playwright — Renders the page in headless Chromium, waits for network idle, then serialises the live DOM (including shadow roots). Use this when you know the page is a SPA. Requires the optional playwright extra ( pip install 'pagetomd[playwright]' ) and a one-time playwright install chromium . Slower and heavier than httpx , but the only way to get content that lives behind a JS framework.
auto — Fetches with httpx first, then inspects the result: if the <body> text is under 200 characters and the HTML contains SPA markers ( data-reactroot , <div id="__next"> , a "you need to enable javascript" noscript tag, etc.), it re-fetches with Playwright. A second safety net fires if httpx returned HTML that looked non-empty but the extractor still couldn't pull any content — Playwright gets a shot then too. If Playwright is unavailable, the page is counted as "empty" in the crawl summary rather than a hard failure. Best choice when you're pointed at an unfamiliar URL.
Use the default single-page mode when you have a specific URL (or a short list piped through a while read loop). Use --crawl when you want every page under a URL prefix — it discovers links automatically, deduplicates, mirrors the URL hierarchy on disk, and reuses a single fetcher context so Playwright doesn't relaunch Chromium per page. See the crawl cookbook recipe for the full flag set.
Running pagetomd http://127.0.0.1:8765/blog.html --no-fetched-at -o - against the blog.html fixture prints (first ~15 lines shown):
---
url : http://127.0.0.1:8765/blog.html
final_url : http://127.0.0.1:8765/blog.html
title : Why We Rewrote Our Build System in Rust
author : Jane Doe
date : ' 2024-08-14 '
description : A retrospective on migrating our monorepo build pipeline from Python to Rust, and what we learned along the way.
site_name : Example Engineering Blog
language : en
tool : pagetomd
tool_version : 0.4.0
---
# Why We Rewrote Our Build System in Rust
Three years ago, our monorepo build pipeline was a sprawling Python application held together with shell scripts and prayer. ...
When fetched_at is enabled (the default), an extra fetched_at: '2026-06-15T12:34:56Z' line is included in the frontmatter. Fields whose value cannot be detected (e.g. language , author ) are omitted from the YAML.
A compact overview — see pagetomd --help for the full list.
Every flag has a PAGETOMD_<UPPER_NAME> equivalent. For example:
PAGETOMD_TIMEOUT=60 PAGETOMD_FETCHER=auto pagetomd https://example.com
CLI flags always override env vars; env vars override the built-in defaults.
Code
Meaning
0
Success.
1
Unexpected internal error.
2
Fetch failure (DNS, HTTP, robots.txt, redirect cap).
3
Extraction or conversion failure (empty body, malformed HTML).
4
Output write failure (permissions, disk, atomic-rename clash).
5
Missing optional dependency (e.g. playwright not installed).
64
Usage or configuration error (bad flag, invalid value).
130
Interrupted by user (Ctrl-C).
How it works
One paragraph plus a diagram of the pipeline:
URL ──► Fetcher ──► Extractor ──► Converter ──► Postprocess ──► Writer
(httpx / (BS4 clean (markdownify (NFC, heading (atomic
playwright) + trafilatura) + GFM tables) hierarchy, file +
URL absolutise) YAML)
The fetcher ( httpx by default, playwright for SPAs) downloads the page with retries and robots.txt enforcement. The extractor runs a BeautifulSoup pre-clean pass (strip scripts/styles/nav/ads) then hands the cleaned tree to trafilatura to identify main content and harvest metadata. The converter renders the surviving HTML to Markdown via a customised markdownify subclass (ATX headings, fenced code blocks with language hints, GFM tables with wide-table fallbacks). The postprocessor enforces the AI-readiness contract (NFC, zero-width strip, monotonic heading hierarchy, absolute URLs). The writer prepends a YAML frontmatter block and writes atomically (or streams to stdout).
pagetomd is a public-URL-only tool. It refuses to fetch private, loopback, link-local, multicast, reserved, or cloud-metadata addresses by default — and there is no flag to override that. Treat output files as having the same sensitivity as the URL they were fetched from.
CI enforces both a project-wide test coverage floor of 85% and a per-module floor of 90% (line + branch combined) on the four critical modules — extractor , converter , writer , and postprocess . These four carry the AI-readiness contract, so they get the strictest coverage bar.
git clone https://github.com/gs202/PageToMD.git
cd pagetomd
uv sync --extra dev --extra playwright
pre-commit install
uv run pytest
See CONTRIBUTING.md for the full contributor workflow.
Business Source License 1.1 — source-available, free for non-commercial use. Conv

[truncated]
