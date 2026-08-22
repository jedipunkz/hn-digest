---
source: "https://github.com/Parthuss/stash"
hn_url: "https://news.ycombinator.com/item?id=49401418"
title: "Show HN: Stash, turn your Instagram saves into notes Claude finds on its own"
article_title: "GitHub - Parthuss/stash: Turn the Instagram posts you save into notes Claude finds on its own. · GitHub"
image: "https://opengraph.githubassets.com/9650b8b0c8c20b61f8192e4a691c57caed93eee3ccc9483f884f34f0795784c3/Parthuss/stash"
author: "ParthAkholkar"
captured_at: "2026-08-22T17:12:40Z"
capture_tool: "hn-digest"
hn_id: 49401418
score: 1
comments: 0
posted_at: "2026-08-22T16:42:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stash, turn your Instagram saves into notes Claude finds on its own

- HN: [49401418](https://news.ycombinator.com/item?id=49401418)
- Source: [github.com](https://github.com/Parthuss/stash)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T16:42:39Z

## Translation

タイトル: HN を表示: Stash、Instagram の保存をクロードが独自に見つけたメモに変える
記事タイトル: GitHub - Parthuss/stash: 保存した Instagram の投稿をクロードが独自に見つけたメモに変換します。 · GitHub
説明: 保存した Instagram の投稿を、クロードが独自に見つけたメモに変換します。 - パルサス/隠し場所

記事本文:
GitHub - Parthuss/stash: 保存した Instagram の投稿を、Claude が独自に見つけたメモに変換します。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
パルサス
/
隠し場所
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
21 コミット 21 コミット フォルダーとファイル
docs/ case-study-assets docs/ case-study-assets プロモーション プロモーション スクリプト スクリプト ショートカット ショートカット スタッシュ スタッシュ テスト テスト ワーカー ワーカー .env.example .env.example .gitignore .gi

tignore CASE_STUDY.md CASE_STUDY.md HANDOFF.md HANDOFF.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
保存した Instagram の投稿を、クロードが独自に見つけたメモに変換します。
すでに確実にキャプチャしています。共有し、自分のアカウントをタップして完了です。失敗するのは、
その後のすべて: URL からは投稿の内容について何もわかりません。
トリアージは決して行われず、いつ役立つかは何も表面化されません。これは、
ブックマークツール。あなたが保存したものを文字に起こして、プロンプトなしで返します
仕事をしている間。
デモ圧縮.mp4
シェアシート ─┐
§─► /ingest ─► キュー ─► フェッチ ─► 転記 ─► フレームゲート ─► 抽出 ─► vault/*.md + FTS5
データエクスポート ─┘ │
クロード コード スキル · MCP · ウィークリー ダイジェスト ◄┘
セットアップ
CDの隠し場所
/opt/homebrew/bin/python3.12 -m venv .venv
.venv/bin/pip install -e " .[mcp] "
cp .env.example .env
.venv/bin/python -m 隠しドクター
医師は不足しているものを指定し、それを修正するための正確な指示を出します。 3つのこと
問題:
2. Groq キー。で入手してください
console.groq.com/keys を .env に置きます。
Groq Whisper はリール/ビデオ音声を文字に起こします。 qwen/qwen3.6-27b は静的な投稿を読み取ります。
カルーセル スライド、および選択されたリール フレームを抽出し、最終的な JSON ノートを返します。
処理中にクロード呼び出しは行われません。文字起こしのみの場合、オフライン
代替案は .venv/bin/pip install -e ".[local-whisper]" のままです。
3. Instagram Cookie — おそらく必要ありません。両方の Cookie 設定
意図的に空のまま発送します。公開リールは匿名でダウンロードできます。実測
2026 年 8 月に保存された投稿、メタデータと完全なビデオ + オーディオの両方が何も表示されずに戻ってきました。
セッション全然。
これを信頼する前に、自分で確認してください。
yt-dlp --skip-download --print " %(id)s | %(uploader)s " " https://www.instagram.com/reel/XXXX/ "
設定

不要な場合に Cookie ソースを使用することは、それを放置することよりも明らかに悪いです
空白 — --cookies-from-browser chrome はキーチェーン プロンプトをトリガーし、
プロンプトに応答しないと成功するはずのダウンロードが失敗します。
ログインの壁 (プライベート アカウント、継続的なレート制限) にぶつかった場合は、
次に、セッション、最も安い順:
STASH_COOKIES_FILE — instagram.com をスコープとする cookies.txt でエクスポートされます。
「Get cookies.txt LOCALLY」拡張子、chmod 600 。 yt-dlp があなたの Instagram を取得します
セッションだけでなく、他には何もありません。
STASH_COOKIES_FROM_BROWSER=firefox — macOS ではキーチェーン許可は必要ありません。
=chrome — 最後の手段。 Chrome の Cookie はログイン時に AES 暗号化されます
Instagram だけでなく、サインインしているすべてのサイトを復号化するキーチェーン キー。
無人作業員の使用を継続的に許可することは、はるかに広範囲にわたる問題です
このタスクが保証する半径を超えています。
Cookie ファイルは、平文の Instagram ログインです。 git から外しておいてください。取り消す
Instagram → [設定] → [セキュリティ] → ログインしている場所からアクセスしてください。
静的投稿では 1 つのビジョン画像を使用します。カルーセルはスライドの順序を保持し、すべてのスライドを送信します
スライド; Groq のリクエストごとの画像制限は、順序付けされた順序で自動的に処理されます。
バッチ。混合カルーセルでは、各ビデオ スライドも転写されます。リールは
既存の Whisper と選択フレーム パス。
.venv/bin/python -m stash add " https://www.instagram.com/reel/... " -n " 保存した理由 "
.venv/bin/python -m 隠しプロセス
.venv/bin/python -m stash 検索「エージェント メモリ」
.venv/bin/python -m 隠しステータス
メモは YAML フロントマターを使用したマークダウンとして Vault/ に配置されます。それが耐久性です
アーティファクト — SQLite インデックスが派生され、stash の再インデックスによってディスクから再構築されます。
3 つの部分: Cloudflare に投稿するショートカット、
キューとそれを排出する Mac 上のデーモン。という順番で設定していきます。
これが作るものです

Mac から独立してキャプチャ - 電話が話しかけます
携帯電話、カフェ、蓋を閉めた場所など、どこからでも HTTPS 経由で Cloudflare を利用できます。
cd ワーカー && npm インストール
npx Wrangler ログイン # ブラウザ、無料アカウント、カードなし
npx wrangler d1 create stash # wrangler.toml に ID を貼り付けます
npx wrangler d1 移行は stash --remote を適用します
npx Wrangler Secret put STASH_SECRET # 任意の長いランダム文字列
npxラングラーデプロイ
デプロイされた URL と同じシークレットを STASH_WORKER_URL として .env に配置し、
STASH_SECRET 。ダウンストリームのすべてが独自にリモート キューに切り替わります。
全体の無料利用枠: ワーカーは 100,000 リクエスト/日、D1 は 5 GB。 R2 がないため、
支払い方法 — R2 はカードを要求する唯一の製品です。
現在実行中のものはこれを必要としないため、未設定のままになります。
2. ショートカットを構築してインストールする
sh ショートカット/build.sh
Worker URL とシークレットを使用して、shortcuts/Stash.cherri.template をレンダリングします。
Cherri を使用してコンパイルし、署名します。エアドロップ
結果として得られる携帯電話への Stash.ショートカット、または同じ Apple 上の Mac でそれを開きます
IDと同期します。
古い Stash ショートカットをすべて削除します。のいくつかのほぼ同一のエントリ
共有シート、一部はもう存在しないエンドポイントを指していますが、これがどのように行われるかです
前に静かに壊れました。
ショートカットには、意図的に「保存」ではなく「キューに登録」と表示されます。それが知っているのは、
リクエストがCloudflareに届きました。メモが実際に存在するという確認が来る
手順4から別途行います。
sh stash/launchd/install.sh # ログイン時に開始し、クラッシュ時に再起動します
または、ターミナルにデーモンを隠して、その動作を監視します。いずれにしても:
stash status # は ALIVE/DOWN と表示され、pid の有効性とハートビートの経過時間を確認します
このチェックが存在するのは、受信側プロセスが一度静かに終了し、何も起こらなかったためです。
そう言いました。停止しただけで終了していないデーモンも、停止したデーモンだけでなく報告されます。
stash Notice # テストを送信します。 --fail (障害形状の場合)
STASを設定する

最初に .env の ntfy または imessage への H_NOTIFY — コメントを参照してください
そこに。どちらも無料です。 ntfy は動作することが確認されており、imessage にはアプリは必要ありません
インストールしますが、1 回限りの自動化許可が必要なので、信頼する前にテストしてください。
処理。保存は失われることはありません - 保存は D1 に保存されます - しかし、メモは
次に Mac が起動し、デーモンが実行されます。キャプチャはもう必要のない部分です
何かに依存します。
古い iCloud ファイル パス (shortcuts/Stash-icloud.cherri + stash watch ) はまだ残っています
インフラストラクチャをゼロにしたい場合は、機能し、アカウントはまったく必要ありません。
既存の保存をバックフィルする
Instagram の保存済みコレクションを公開した API はありません。データのエクスポートは唯一の
Meta がそれを作成するのに数時間から数日かかります。今すぐリクエストしてください:
アカウント センター → 情報と権限 → 情報をエクスポート →
JSON
.venv/bin/python scripts/import_export.py ~ /Downloads/instagram-export --dry-run
.venv/bin/python scripts/import_export.py ~ /Downloads/instagram-export
その後、ゆっくりと水を切ります。数百のパーマリンクを yt-dlp 経由で連続してプッシュする
これはレート制限を取得する最速の方法ですが、ライブ キャプチャも中断されます。
while .venv/bin/python -m stash process --limit 1 | grep -q が書き込みました。寝ます 45 ;完了しました
リコール
ここが1ヶ月後も使い続けるかどうかを決める部分です。
MCP サーバー — stash/mcp_server.py 、ユーザー スコープで登録されています:
クロード mcp add --scope user stash -- /path/to/stash/.venv/bin/python -m stash.mcp_server
プロジェクト スコープではなくユーザー スコープ — すべてのクロード コード セッションで接続されます
このリポジトリ内だけでなく、どのリポジトリにいても関係ありません。クロードに与える
search_stash 、 get_stash_note 、 list_stash_topics 、 Recent_stash 、
mark_stash_used 。 search_stash はコンパクトなヒット (短い ID、タイトル、
1 行の概要、ツール）頻繁に電話をかけるのを安く抑えるため。 get_s

タッシュノート
ヒットしたものについて、完全な詳細（トランスクリプト、次のステップ、パーマリンク）を取得します
重要であることがわかります。
クロード コード スキル — ~/.claude/skills/stash-recall/ 、ユーザー スコープと
同じ理由で、保管庫内の資料は特定のものではありません。
プロジェクトなので、どこで作業していてもスキルを起動する必要があります。それは
技術的な作業を依頼したときではなく、技術的な作業を開始したときにトリガーされるように書かれています。
検索してください。質問しようとは決して思わないからです。
mark_stash_used はオプションのように見えますが、そうではありません。使用済みか未使用かだけが重要です
これが知識ベースなのか墓場なのかの尺度です。何もなかったら
1 か月後に使用済みとしてマークを付けた場合は、さらにメモを記入するのではなく、リコールの仕組みを変更します。
検索はハイブリッドです — FTS5 (正確なトークン: リポジトリ名、製品名) が融合されています
重み付き逆順位融合によるローカル ベクトル検索 (意味) を使用します。プレーン
キーワード検索だけでは、「ビデオを作成するにはどうすればよいですか」という言い換えによって見逃したものを発見します。
「自動的に」では、実際にビデオ生成に関する 2 つのメモが 6 位と
8 番目、WhatsApp チャットボットの背後にあります。BM25 には「自動的に」という概念がないためです。
と「世代」は関連するアイデアです。埋め込みはローカルで実行されます。
fastembed + BAAI/bge-small-en-v1.5 (なし
API キー、ネットワーク、PyTorch なし) と RRF フュージョンが実行されます。
同じ SQLite ファイル内の sqlite-vec
他のすべてと同様に、個別のベクトル データベースは必要ありません。キーワードのみに低下します。
いずれかが利用できない場合は、1 回限りの警告が表示されます。キャプチャは決して失敗することはありません
検索インフラストラクチャ上で。隠しドクターはベクターのステータスとフラグを報告します。
保存されたベクトルは、現在のものとは異なるモデルに埋め込まれていました
設定されています。
Instagram ではウィスパーが必須です — リールにはアプリ内に自動キャプションがありますが、
Instagram はダウンローダーに公開しない
( yt-dlp#15874 ) したがって、
字幕ファーストのショートカットは、他のツールがここで起動することはありません。
与えられた

トランスクリプト。オーディオが停止した場所のみフレームがプルされます。
自給自足 — ウィスパーが確信を持てなかったセグメントや、「これを実行してください」のようなフレーズ
コマンド」または「バイオのリンク」が画面を指します。アイデアを借用しました
メディア-mcp 。トーキングヘッドが 1 つ獲得
フレーム。画面録画では重要な瞬間にフレームが取得されます。サイレントリール
偶数サンプリングに戻ります。上限は 5 — フレームがトークン コストの大半を占めます。
何かを調整する前に知っておく価値のある 2 つの視力制限。どちらも基準として測定されます。
想定されているものではなく、ライブ API:
リクエストごとに 3 つの画像は API のハードキャップであり、選択できるものではありません。 4番目が戻ってくる
HTTP 400. _describe_visuals バッチ。
512px が最適で、これより大きいほど悪くなります。トークンのコストは全体的に均一です
解像度 (Groq は画像を固定予算に正規化します)、ただし高密度です
スクリーンショット 512px は「MIT ライセンス」を含む 1328 文字を転写、1024px
管理対象は最大 700 で、小さいバージョンでは正しく読み取れるテキストが文字化けしていました。それを上げる
自由そうに見えて実はそうではない。
.venv/bin/python -m pytest テスト/ -q
レイアウト
隠し場所/
config.py 環境 + パス
db.py キャプチャ キュー + FTS5 メモ インデックス
fetch.py CDN ダイレクト / yt-dlp / コバルト
transcribe.py groq またはより高速なウィスパー (セグメントごとの信頼性あり)
Frames.py の信頼ゲート
extract.py groq ビジョン + 構造化 JSON
vault.py マークダウン + フロントマター
Pipeline.py オーケストレーション
ワーカー支援のremote.py

[切り捨てられた]

## Original Extract

Turn the Instagram posts you save into notes Claude finds on its own. - Parthuss/stash

GitHub - Parthuss/stash: Turn the Instagram posts you save into notes Claude finds on its own. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Parthuss
/
stash
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
docs/ case-study-assets docs/ case-study-assets promo promo scripts scripts shortcuts shortcuts stash stash tests tests worker worker .env.example .env.example .gitignore .gitignore CASE_STUDY.md CASE_STUDY.md HANDOFF.md HANDOFF.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Turn the Instagram posts you save into notes Claude finds on its own.
You already capture reliably — share, tap your own account, done. What fails is
everything after: a URL tells you nothing about what was in the post, so you
never triage it, and nothing surfaces it when it would be useful. This is not a
bookmarking tool. It transcribes what you saved and hands it back unprompted
while you work.
demo-compressed.mp4
share sheet ─┐
├─► /ingest ─► queue ─► fetch ─► transcribe ─► frame gate ─► extract ─► vault/*.md + FTS5
data export ─┘ │
Claude Code skill · MCP · weekly digest ◄┘
Setup
cd stash
/opt/homebrew/bin/python3.12 -m venv .venv
.venv/bin/pip install -e " .[mcp] "
cp .env.example .env
.venv/bin/python -m stash doctor
doctor names anything missing and the exact command to fix it. Three things
matter:
2. A Groq key. Get one at
console.groq.com/keys and put it in .env .
Groq Whisper transcribes reel/video audio; qwen/qwen3.6-27b reads static posts,
carousel slides, and selected reel frames, then returns the final JSON note.
No Claude call is made while processing. For transcription only, the offline
alternative remains .venv/bin/pip install -e ".[local-whisper]" .
3. Instagram cookies — you probably don't need any. Both cookie settings
ship blank on purpose. Public reels download fine anonymously; measured on real
saved posts in August 2026, metadata and full video+audio both came back with no
session at all.
Verify it yourself before trusting this:
yt-dlp --skip-download --print " %(id)s | %(uploader)s " " https://www.instagram.com/reel/XXXX/ "
Setting a cookie source when you don't need one is actively worse than leaving it
blank — --cookies-from-browser chrome triggers a keychain prompt, and an
unanswered prompt fails the download that would otherwise have succeeded.
If you do hit login walls (private accounts, sustained rate limiting), add a
session then, cheapest first:
STASH_COOKIES_FILE — a cookies.txt scoped to instagram.com, exported with
the "Get cookies.txt LOCALLY" extension, chmod 600 . yt-dlp gets your Instagram
session and nothing else.
STASH_COOKIES_FROM_BROWSER=firefox — no keychain grant needed on macOS.
=chrome — last resort. Chrome's cookies are AES-encrypted with a login
keychain key that decrypts every site you are signed into, not just Instagram.
Standing permission for an unattended worker to use that is a far wider blast
radius than this task warrants.
A cookies file is your Instagram login in plain text. Keep it out of git; revoke
it from Instagram → Settings → Security → where you're logged in.
Static posts use one vision image. Carousels preserve slide order and send every
slide; Groq's per-request image limit is handled automatically in ordered
batches. Mixed carousels also transcribe each video slide. Reels keep the
existing Whisper plus selective-frame path.
.venv/bin/python -m stash add " https://www.instagram.com/reel/... " -n " why I saved it "
.venv/bin/python -m stash process
.venv/bin/python -m stash search " agent memory "
.venv/bin/python -m stash status
Notes land in vault/ as markdown with YAML frontmatter. That's the durable
artifact — the SQLite index is derived and stash reindex rebuilds it from disk.
Three pieces: a Shortcut that posts to Cloudflare, a Worker that holds the
queue, and a daemon on the Mac that drains it. Set up in that order.
This is what makes capture independent of your Mac — the phone talks to
Cloudflare over HTTPS from anywhere: cellular, a café, lid shut.
cd worker && npm install
npx wrangler login # browser, free account, no card
npx wrangler d1 create stash # paste the id into wrangler.toml
npx wrangler d1 migrations apply stash --remote
npx wrangler secret put STASH_SECRET # any long random string
npx wrangler deploy
Put the deployed URL and the same secret into .env as STASH_WORKER_URL and
STASH_SECRET . Everything downstream switches to the remote queue on its own.
Free tier throughout: Workers 100k req/day, D1 5 GB. No R2 and therefore no
payment method — R2 is the one product here that asks for a card, and
nothing currently running needs it, so it is left unconfigured.
2. Build and install the Shortcut
sh shortcuts/build.sh
Renders shortcuts/Stash.cherri.template with your Worker URL and secret,
compiles it with Cherri , and signs it. AirDrop the
resulting Stash.shortcut to your phone, or open it on a Mac on the same Apple
ID and it syncs.
Delete every older Stash shortcut. Several near-identical entries in the
share sheet, some pointing at endpoints that no longer exist, is how this
silently broke before.
The Shortcut says "Queued" , not "saved" — deliberately. It only knows the
request reached Cloudflare. Confirmation that a note actually exists comes
separately, from step 4.
sh stash/launchd/install.sh # starts at login, restarts on crash
Or stash daemon in a terminal to watch it work. Either way:
stash status # says ALIVE/DOWN, checking pid liveness AND heartbeat age
That check exists because a receiver process once died quietly and nothing
said so. A hung-but-not-exited daemon is reported down too, not just a dead one.
stash notify # sends a test; --fail for the failure shape
Set STASH_NOTIFY to ntfy or imessage in .env first — see the comments
there. Both are free; ntfy is verified working, imessage needs no app
install but does need a one-time Automation grant, so test it before trusting it.
Processing. A save is never lost — it sits in D1 — but the note appears when the
Mac is next awake with the daemon running. Capture is the part that no longer
depends on anything.
The old iCloud-file path ( shortcuts/Stash-icloud.cherri + stash watch ) still
works and needs no accounts at all, if you'd rather have zero infrastructure.
Backfilling your existing saves
No API has ever exposed Instagram Saved collections. The data export is the only
route, and Meta takes hours to days to produce it — request it now:
Accounts Center → Your information and permissions → Export your information →
JSON
.venv/bin/python scripts/import_export.py ~ /Downloads/instagram-export --dry-run
.venv/bin/python scripts/import_export.py ~ /Downloads/instagram-export
Then drain slowly. Pushing a few hundred permalinks through yt-dlp back-to-back
is the fastest way to get rate-limited, which breaks live capture too:
while .venv/bin/python -m stash process --limit 1 | grep -q wrote ; do sleep 45 ; done
Recall
This is the part that decides whether you still use it in a month.
MCP server — stash/mcp_server.py , registered at user scope:
claude mcp add --scope user stash -- /path/to/stash/.venv/bin/python -m stash.mcp_server
User scope, not project scope — it's connected in every Claude Code session
regardless of which repo you're in, not only inside this one. Gives Claude
search_stash , get_stash_note , list_stash_topics , recent_stash ,
mark_stash_used . search_stash returns compact hits (short id, title,
one-line summary, tools) to keep it cheap to call often; get_stash_note
pulls full detail — transcript, next step, permalink — for whichever hit
turns out to matter.
Claude Code skill — ~/.claude/skills/stash-recall/ , also user scope and
for the same reason: the material in the vault isn't specific to any one
project, so the skill needs to fire wherever you happen to be working. It's
written to trigger when you start technical work, not when you ask it to
search, because you will never think to ask.
mark_stash_used looks optional and isn't. used vs unused is the only
measure of whether this is a knowledge base or a graveyard. If nothing is ever
marked used after a month, change how recall works rather than filing more notes.
Search is hybrid — FTS5 (exact tokens: a repo name, a product name) fused
with local vector search (meaning) via weighted Reciprocal Rank Fusion. Plain
keyword search alone missed things by paraphrase: "how do I make videos
automatically" ranked the two notes actually about generating video 6th and
8th, behind a WhatsApp chatbot, because BM25 has no notion that "automatically"
and "generation" are related ideas. Embeddings run locally via
fastembed + BAAI/bge-small-en-v1.5 (no
API key, no network, no PyTorch) and RRF fusion runs through
sqlite-vec inside the same SQLite file
as everything else — no separate vector database. Degrades to keyword-only,
with a one-time warning, if either is unavailable; a capture can never fail
over search infrastructure. stash doctor reports vector status and flags if
stored vectors were embedded under a different model than the one currently
configured.
Whisper is mandatory on Instagram — reels have auto-captions in the app but
Instagram doesn't expose them to downloaders
( yt-dlp#15874 ), so the
subtitles-first shortcut other tools lead with never fires here.
Given a transcript, frames are pulled only where the audio stopped being
self-sufficient — segments Whisper was unsure about, and phrases like "run this
command" or "link in bio" that point at the screen. Idea borrowed from
media-mcp . A talking head gets one
frame; a screen recording gets frames at the moments that matter; a silent reel
falls back to even sampling. Capped at 5 — frames dominate token cost.
Two vision limits worth knowing before you tune anything, both measured against
the live API rather than assumed:
3 images per request is a hard API cap , not a choice. A 4th returns
HTTP 400. _describe_visuals batches around it.
512px is the optimum, and bigger is worse. Token cost is flat across
resolution (Groq normalises images to a fixed budget), but on a dense
screenshot 512px transcribed 1328 chars including "MIT license" while 1024px
managed ~700 and garbled text the smaller version read correctly. Raising it
looks free and isn't.
.venv/bin/python -m pytest tests/ -q
Layout
stash/
config.py env + paths
db.py capture queue + FTS5 note index
fetch.py CDN direct / yt-dlp / cobalt
transcribe.py groq or faster-whisper, with per-segment confidence
frames.py the confidence gate
extract.py groq vision + structured JSON
vault.py markdown + frontmatter
pipeline.py orchestration
remote.py Worker-backed

[truncated]
