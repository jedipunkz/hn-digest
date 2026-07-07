---
source: "https://github.com/thomasunise/visionbridge"
hn_url: "https://news.ycombinator.com/item?id=48822662"
title: "I built a tiny proxy that gives GLM 5.2 vision (or any text LLM) – MIT"
article_title: "GitHub - thomasunise/visionbridge: Give text-only LLMs vision. A tiny OpenAI-compatible proxy that lets reasoning models (DeepSeek, Qwen, GLM…) see images by querying a separate vision model through tools: look, OCR, scan, crop, compare. No training, no weights. · GitHub"
author: "thomasunise"
captured_at: "2026-07-07T19:41:43Z"
capture_tool: "hn-digest"
hn_id: 48822662
score: 2
comments: 0
posted_at: "2026-07-07T19:37:55Z"
tags:
  - hacker-news
  - translated
---

# I built a tiny proxy that gives GLM 5.2 vision (or any text LLM) – MIT

- HN: [48822662](https://news.ycombinator.com/item?id=48822662)
- Source: [github.com](https://github.com/thomasunise/visionbridge)
- Score: 2
- Comments: 0
- Posted: 2026-07-07T19:37:55Z

## Translation

タイトル: GLM 5.2 ビジョン (またはテキスト LLM) を提供する小さなプロキシを構築しました – MIT
記事のタイトル: GitHub - thomasunise/visionbridge: テキストのみの LLM にビジョンを与える。 OpenAI 互換の小さなプロキシ。ツール (検索、OCR、スキャン、クロップ、比較) を通じて別のビジョン モデルをクエリすることで、推論モデル (DeepSeek、Qwen、GLM など) が画像を参照できるようにします。トレーニングもウェイトトレーニングも必要ありません。 · GitHub
説明: テキストのみの LLM にビジョンを与えます。 OpenAI 互換の小さなプロキシ。ツール (検索、OCR、スキャン、クロップ、比較) を通じて別のビジョン モデルをクエリすることで、推論モデル (DeepSeek、Qwen、GLM など) が画像を参照できるようにします。トレーニングもウェイトトレーニングも必要ありません。 - トーマスニス/ビジョンブリッジ

記事本文:
GitHub - thomasunise/visionbridge: テキストのみの LLM にビジョンを与えます。 OpenAI 互換の小さなプロキシ。ツール (検索、OCR、スキャン、クロップ、比較) を通じて別のビジョン モデルをクエリすることで、推論モデル (DeepSeek、Qwen、GLM など) が画像を参照できるようにします。トレーニングもウェイトトレーニングも必要ありません。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードして ses を更新してください

シオン。
アラートを閉じる
{{ メッセージ }}
トーマスニス
/
ビジョンブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows ベンチマーク ベンチマークの例 例 ハグギングフェイス ハグギングフェイス テスト テスト ビジョンブリッジ ビジョンブリッジ .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
小さな OpenAI 互換プロキシを通じて、テキストのみの LLM にビジョンを提供します。
VisionBridge はチャット UI とモデルの間に位置します。それは推論モデルを可能にします
LM Studio、Ollama、vLLM、Z.ai、OpenRouter、または任意の OpenAI 互換バックエンドから
ビジョンモデルに的を絞った別の質問をして画像を検査します。
トレーニングはありません。重みはありません。独自の推論モデルとビジョン モデルを持参してください。
OpenWebUI / LibreChat / アプリ
|
v
/v1/chat/completions の VisionBridge
|
+-- 推論モデル: LM Studio / Ollama / vLLM / GLM / DeepSeek
|
+-- ビジョンモデル: LLaVA / Qwen-VL / GLM-V / GPT-4o-mini / 任意の VLM
VisionBridge は OpenAI スタイルのチャット リクエストを受け取ります。リクエストに次の内容が含まれている場合
画像を保存し、推論モデルのプロンプトを書き換えます。
推論モデルでツールを呼び出します。
対象を絞った目視検査用の look(image_id, question)
ocr(image_id) テキスト抽出用
scan(image_id) は、密度の高いページのタイル化された完全なドキュメント テキスト スイープを取得します。
領域を拡大するための Crop_and_look(image_id, box, question)
(小さな作物は、VLM が認識する前に自動的にアップスケールされます)
比較する

(image_id_a、image_id_b、question) 2 つの画像を一緒に判定する場合
視覚モデルは見ることを行います。推論モデルが思考を行います。
ループが開始する前に、すべての画像がワンショット シーンの説明を取得します。
ビジョンモデル (コンテンツによって並行してキャッシュされる)、つまり推論モデル
グローバルなコンテキストから始まり、ツールの予算を細部に費やします。複数の
ツール呼び出しは 1 回のターンで同時に実行され、複数回にわたってイメージが繰り返されます。
会話は、コンテンツ アドレス キャッシュから提供されるのではなく、コンテンツ アドレス キャッシュから提供されます。
再フェッチされ、再デコードされます。
ツール呼び出しでは、バックエンドのネイティブ OpenAI 関数呼び出しが使用されます。
利用可能です。デフォルトの TOOL_MODE=auto では、VisionBridge は推論を調査します。
バックエンドを一度実行すると、プロンプト JSON プロトコルに透過的にフォールバックします。
ツールをサポートしていないバックエンドとモデル - したがって、強力なホスト型で動作します
モデルも弱いローカルモデルも同様です。
ループは不完全なモデルに耐えるように構築されています: 幻覚のような画像 ID、不良品
クロップ ボックス、不正な形式の JSON、ビジョン バックエンドの問題がフィードバックされます。
推論モデルを観察として作成し、それ自体を修正できるようにし、
検査制限により、エラーではなく最終的なベストエフォート型の回答がトリガーされます。
docker run -p 8080:8080 \
-e REASONING_BASE_URL=http://host.docker.internal:1234/v1 \
-e REASONING_API_KEY=lm-studio \
-e REASONING_MODEL=ローカルモデル \
-e VISION_BASE_URL=http://host.docker.internal:11434/v1 \
-e VISION_API_KEY=オラマ \
-e VISION_MODEL=llava:13b \
ghcr.io/eekosystems/visionbridge:最新
次に、OpenAI 互換クライアントを次のように指定します。
ベース URL: http://localhost:8080/v1
APIキー：何でも
モデル: ビジョンブリッジ
地域開発
CD ビジョンブリッジ
Python -m venv .venv
ソース .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e " .[dev] "
cp .env.example .env # Windows: .env.example .env をコピー
ビジョンブリッジサーブ

--ホスト 0.0.0.0 --ポート 8080
以下の方法でいつでもバックエンド構成を確認できます。
ビジョンブリッジの医師
両方のバックエンドが到達可能であること、構成されたモデルが存在すること、および
テストが成功したことを確認します。
すべての設定は環境変数 (または .env ) から取得されます。
1 つのインスタンスが、異なるモデルの下で複数の推論とビジョンのペアを提供できます。
名前。未指定のフィールドはデフォルトのペアから継承します。
EXTRA_MODELS = {"visionbridge-fast": {"reasoning": {"model": "qwen3:8b"}, "vision": {"model": "llava:7b"}}}
(または、EXTRA_MODELS_FILE で JSON ファイルを指定します)。すべての名前が表示されます
GET /v1/models を実行すると、クライアントはモデル フィールドとのペアを選択します。
「開発」タブに移動し、ローカルサーバーを起動します。
REASONING_BASE_URL = http://localhost:1234/v1
REASONING_API_KEY = lm-studio
REASONING_MODEL = ローカルモデル
LM Studio は、ビジョン対応のモジュールをロードする場合、ビジョン バックエンドとしても使用できます。
モデルを作成し、同じ OpenAI 互換サーバーを通じて公開します。
REASONING_BASE_URL = http://localhost:11434/v1
REASONING_API_KEY = オラマ
REASONING_MODEL = qwen3:32b
VISION_BASE_URL = http://localhost:11434/v1
VISION_API_KEY = オラマ
VISION_MODEL = ラヴァ:13b
vLLM バックエンド
REASONING_BASE_URL = http://localhost:8000/v1
REASONING_API_KEY = トークン-abc123
REASONING_MODEL = クウェン/Qwen3-32B
VISION_BASE_URL = http://localhost:8001/v1
VISION_API_KEY = トークン-abc123
VISION_MODEL = Qwen/Qwen3-VL-8B-命令
OpenWebUI のセットアップ
管理パネル -> 設定 -> 接続 -> OpenAI 互換
ベース URL: http://host.docker.internal:8080/v1
APIキー:何でも
モデル: ビジョンブリッジ
API サーフェス
VisionBridge は、ほとんどのクライアントが必要とする小さなサブセットを意図的に公開します。
GET /health (両方のバックエンドが到達可能であることを確認するには、?probe=true を追加します)
GET /v1/traces/{completion_id} — 最近の完全なツールループ トレース
リクエスト: すべてのツール呼び出し、その引数、vi

sion モデルの観察、および
タイミング。 「なぜそう言われたのか？」に答えるのに非常に貴重です。
応答には、あらゆる推論とビジョンにわたって集約された使用ブロックが含まれます
リクエストに対して行われた呼び出しと、ログを関連付けるための X-Request-ID ヘッダー。
ストリーミング : 画像のないリクエストは、トークンごとに直接ストリーミングされます。
推論バックエンド。画像ストリームを伴うリクエストも: SSE ハートビートは維持されます
ツールの実行中、およびネイティブ ツール呼び出しモードでは接続が維持されます。
最終的な答えは、推論モデルが生成するときにトークンごとに転送されます。
(プロンプト JSON モードはループをバッファリングし、応答をチャンクでストリーミングします)。
エラー: バックエンド障害により HTTP 502、無効な入力 400 が返されます。
認証: /v1/* でベアラー キーを要求するには、BRIDGE_API_KEYS=key1,key2 を設定します。
エンドポイント (/health は開いたままになります)。インターネットに接続するものであれば、やはり
TLS を前面に持つリバース プロキシ。
画像 URL の取得はデフォルトで強化されています。ホストはプライベートに解決され、
ループバック、またはリンクローカル アドレスが拒否され (SSRF ガード)、リダイレクトが拒否されます。
ホップバイホップで検証され、ダウンロードにはサイズ制限があり、ペイロードは次のようにデコードする必要があります。
実際の画像。次の場合は、独自のリバース プロキシ/認証レイヤーの背後で VisionBridge を実行します。
これをローカルホストを超えて公開する場合、意図的に組み込み認証なしで出荷されます。
Benchmarks/ には、直接スコアに対して VisionBridge をスコアリングするハーネスが含まれています
DocVQA スタイルの JSONL タスク セット上の VLM (緩和された精度 + ANLS)。参照
ベンチマーク/README.md 。
積極的に開発されています。 0.2.0 の新機能については、CHANGELOG.md を参照してください。
次に:
キャプションのみのプロンプトに対する公開されたベンチマーク数値
永続的なディスク上のイメージ/プライマー キャッシュ
OpenWebUI と LibreChat のスクリーンショット
貢献を歓迎します — COTRIBUTING.md を参照してください。
テキストのみの LLM にビジョンを与えます。別のビジョンをクエリすることで推論モデル (DeepSeek、Qwen、GLM…) に画像を表示させる小さな OpenAI 互換プロキシ

ツールを使用してモデルを作成します: 検索、OCR、スキャン、切り抜き、比較。トレーニングもウェイトトレーニングも必要ありません。
thomasunise.com/visionbridge
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give text-only LLMs vision. A tiny OpenAI-compatible proxy that lets reasoning models (DeepSeek, Qwen, GLM…) see images by querying a separate vision model through tools: look, OCR, scan, crop, compare. No training, no weights. - thomasunise/visionbridge

GitHub - thomasunise/visionbridge: Give text-only LLMs vision. A tiny OpenAI-compatible proxy that lets reasoning models (DeepSeek, Qwen, GLM…) see images by querying a separate vision model through tools: look, OCR, scan, crop, compare. No training, no weights. · GitHub
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
thomasunise
/
visionbridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows benchmarks benchmarks examples examples huggingface huggingface tests tests visionbridge visionbridge .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
Give text-only LLMs vision through a tiny OpenAI-compatible proxy.
VisionBridge sits between your chat UI and your models. It lets a reasoning model
from LM Studio, Ollama, vLLM, Z.ai, OpenRouter, or any OpenAI-compatible backend
inspect images by asking a separate vision model targeted questions.
No training. No weights. Bring your own reasoning model and your own vision model.
OpenWebUI / LibreChat / app
|
v
VisionBridge at /v1/chat/completions
|
+-- reasoning model: LM Studio / Ollama / vLLM / GLM / DeepSeek
|
+-- vision model: LLaVA / Qwen-VL / GLM-V / GPT-4o-mini / any VLM
VisionBridge receives an OpenAI-style chat request. If the request contains
images, it stores the images, rewrites the prompt for the reasoning model, and
lets the reasoning model call tools:
look(image_id, question) for targeted visual inspection
ocr(image_id) for text extraction
scan(image_id) for a tiled full-document text sweep of dense pages
crop_and_look(image_id, box, question) for zooming into a region
(small crops are automatically upscaled before the VLM sees them)
compare(image_id_a, image_id_b, question) for judging two images together
The vision model does the seeing. The reasoning model does the thinking.
Before the loop starts, every image gets a one-shot scene description from
the vision model (in parallel, cached by content), so the reasoning model
begins with global context and spends its tool budget on details. Multiple
tool calls in one turn run concurrently , and repeated images across a
conversation are served from a content-addressed cache instead of being
re-fetched and re-decoded.
Tool calling uses the backend's native OpenAI function calling when
available. In the default TOOL_MODE=auto , VisionBridge probes the reasoning
backend once and transparently falls back to a prompt-JSON protocol for
backends and models that don't support tools — so it works with strong hosted
models and weak local ones alike.
The loop is built to survive imperfect models: hallucinated image ids, bad
crop boxes, malformed JSON, and vision-backend hiccups are fed back to the
reasoning model as observations so it can correct itself, and hitting the
inspection limit triggers a final best-effort answer instead of an error.
docker run -p 8080:8080 \
-e REASONING_BASE_URL=http://host.docker.internal:1234/v1 \
-e REASONING_API_KEY=lm-studio \
-e REASONING_MODEL=local-model \
-e VISION_BASE_URL=http://host.docker.internal:11434/v1 \
-e VISION_API_KEY=ollama \
-e VISION_MODEL=llava:13b \
ghcr.io/eekosystems/visionbridge:latest
Then point your OpenAI-compatible client at:
Base URL: http://localhost:8080/v1
API key: anything
Model: visionbridge
Local Development
cd visionbridge
python -m venv .venv
source .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e " .[dev] "
cp .env.example .env # Windows: copy .env.example .env
visionbridge serve --host 0.0.0.0 --port 8080
Check your backend configuration any time with:
visionbridge doctor
It verifies both backends are reachable, that the configured models exist, and
that a test completion succeeds.
All settings come from environment variables (or .env ):
One instance can serve several reasoning+vision pairs under different model
names. Unspecified fields inherit from the default pair:
EXTRA_MODELS = {"visionbridge-fast": {"reasoning": {"model": "qwen3:8b"}, "vision": {"model": "llava:7b"}}}
(or point EXTRA_MODELS_FILE at a JSON file). All names appear in
GET /v1/models , and clients select a pair with the model field.
Go to the Developer tab and start the local server.
REASONING_BASE_URL = http://localhost:1234/v1
REASONING_API_KEY = lm-studio
REASONING_MODEL = local-model
LM Studio can also be used as the vision backend if you load a vision-capable
model and expose it through the same OpenAI-compatible server.
REASONING_BASE_URL = http://localhost:11434/v1
REASONING_API_KEY = ollama
REASONING_MODEL = qwen3:32b
VISION_BASE_URL = http://localhost:11434/v1
VISION_API_KEY = ollama
VISION_MODEL = llava:13b
vLLM Backend
REASONING_BASE_URL = http://localhost:8000/v1
REASONING_API_KEY = token-abc123
REASONING_MODEL = Qwen/Qwen3-32B
VISION_BASE_URL = http://localhost:8001/v1
VISION_API_KEY = token-abc123
VISION_MODEL = Qwen/Qwen3-VL-8B-Instruct
OpenWebUI Setup
Admin Panel -> Settings -> Connections -> OpenAI-compatible
Base URL: http://host.docker.internal:8080/v1
API Key: anything
Model: visionbridge
API Surface
VisionBridge intentionally exposes the small subset most clients need:
GET /health (add ?probe=true to verify both backends are reachable)
GET /v1/traces/{completion_id} — the full tool-loop trace for a recent
request: every tool call, its arguments, the vision model's observation, and
timings. Invaluable for answering "why did it say that?"
Responses include a usage block aggregated across every reasoning and vision
call made for the request, and an X-Request-ID header for correlating logs.
Streaming : image-free requests are streamed token-by-token straight from
the reasoning backend. Requests with images stream too: SSE heartbeats keep
the connection alive while tools run, and in native tool-calling mode the
final answer is forwarded token-by-token as the reasoning model produces it
(prompt-JSON mode buffers the loop and streams the answer in chunks).
Errors : backend failures return HTTP 502, invalid inputs 400.
Auth : set BRIDGE_API_KEYS=key1,key2 to require a Bearer key on /v1/*
endpoints ( /health stays open). For anything internet-facing, still prefer a
reverse proxy with TLS in front.
Image URL fetching is hardened by default: hosts resolving to private,
loopback, or link-local addresses are rejected (SSRF guard), redirects are
validated hop-by-hop, downloads are size-capped, and payloads must decode as
real images. Run VisionBridge behind your own reverse proxy / auth layer if
you expose it beyond localhost — it deliberately ships without built-in auth.
benchmarks/ contains a harness that scores VisionBridge against a direct
VLM (relaxed accuracy + ANLS) on DocVQA-style JSONL task sets. See
benchmarks/README.md .
Actively developed. See CHANGELOG.md for what's new in 0.2.0.
Next up:
published benchmark numbers against caption-only prompting
persistent on-disk image/primer cache
OpenWebUI and LibreChat screenshots
Contributions welcome — see CONTRIBUTING.md .
Give text-only LLMs vision. A tiny OpenAI-compatible proxy that lets reasoning models (DeepSeek, Qwen, GLM…) see images by querying a separate vision model through tools: look, OCR, scan, crop, compare. No training, no weights.
thomasunise.com/visionbridge
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
