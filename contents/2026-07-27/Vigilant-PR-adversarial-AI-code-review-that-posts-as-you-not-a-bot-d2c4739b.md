---
source: "https://github.com/tllongdev/vigilant-pr"
hn_url: "https://news.ycombinator.com/item?id=49072485"
title: "Vigilant PR – adversarial AI code review that posts as you, not a bot"
article_title: "GitHub - tllongdev/vigilant-pr: Adversarial AI code review, posted as you - reviews pull requests on your behalf under your own identity, not a bot. Bring your own model. · GitHub"
author: "tllongdev"
captured_at: "2026-07-27T17:30:52Z"
capture_tool: "hn-digest"
hn_id: 49072485
score: 1
comments: 0
posted_at: "2026-07-27T17:00:42Z"
tags:
  - hacker-news
  - translated
---

# Vigilant PR – adversarial AI code review that posts as you, not a bot

- HN: [49072485](https://news.ycombinator.com/item?id=49072485)
- Source: [github.com](https://github.com/tllongdev/vigilant-pr)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T17:00:42Z

## Translation

タイトル: Vigilant PR – ボットではなくあなたとして投稿する敵対的な AI コード レビュー
記事のタイトル: GitHub - tllongdev/vigilant-pr: Adversarial AI コード レビュー、ユーザーとして投稿 - ボットではなく、ユーザー自身の ID でユーザーに代わってプル リクエストをレビューします。ご自身のモデルをご持参ください。 · GitHub
説明: ユーザーとして投稿された敵対的 AI コード レビュー - ボットではなく、ユーザー自身の ID でユーザーに代わってプル リクエストをレビューします。ご自身のモデルをご持参ください。 - tlllongdev/vigilant-pr

記事本文:
GitHub - tllongdev/vigilant-pr: Adversarial AI コード レビュー、ユーザーとして投稿 - ボットではなく、ユーザー自身の ID でユーザーに代わってプル リクエストをレビューします。ご自身のモデルをご持参ください。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
tllongdev
/
警戒PR
プ

ブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
開発 ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
61 コミット 61 コミット .claude/ スキル/ vigilant-pr .claude/ スキル/ vigilant-pr .cursor/ スキル/ vigilant-pr .cursor/ スキル/ vigilant-pr .github .github アセット アセット docs docs src/ vigilant src/ vigilant テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス通知 通知 README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ポータブルでワークフローに依存しない AI プルリクエスト レビュアー。ユーザーに代わってレビュー コメントを投稿します。一般的なボットではなく、GitHub ID です。
プル リクエストのレビュアーとしてタグ付けされ、Vigilant PR がそれをレビューして投稿します
あなたとしてのコメント - ボットではなくあなたの GitHub ID。リポジトリ側の設定はありません。
GitHub アプリ、トークンだけ。
フローチャート LR
CFG(["端末での設定:<br/>警戒初期化"]) --> C1
CFG -.-> C2
C1(["vigilant github-watch"]) --> A["あなたは GitHub PR の<br/>レビュアーとしてタグ付けされています"]
C2(["警戒するスラックウォッチ<br/>または警戒するチームウォッチ"]) -.-> A2["ベータ版: あなたは<br/>設定された Slack または Teams チャネルで @ メンションされています"]
A --> B[「警戒中の PR は<br/>PR と差分を読む」]
A2 -.-> B
B --> C[「敵対的レビュー、<br/>重大度タグ付けされた調査結果」]
C --> D[「ボットではなくあなたとしてインラインコメントを<br/>投稿します」]
D --> E["何もブロックされない場合は承認し、<br/>ブロックされる場合はコメントします"]
classDef beta ストローク:#a98bff、カラー:#a98bff、ストローク-ダシャーレイ:5 5;
classDef cmd 塗りつぶし:#0d2038、ストローク:#4da3ff、色:#cfe6ff;
クラスA2ベータ。
クラスC2ベータ。
クラスCFGコマンド;
クラス C1 cmd;
読み込み中
ステータス
v1 - GitHub プル リクエストのレビューにすぐに使用できます。
任意の AI モデルを使用する

あなたが欲しいのです。 Vigilant PR に優先する API キーを提供します。
プロバイダーであり、そのモデル - Claude、GPT、Gemini、Grok (xAI)、Llama、を使用します。
NVIDIA、OpenRouter、またはローカルで実行するモデル。無料のものもあれば有料のものもありますので、お好みで選択してください。
キーを設定してモデルを選択するだけです (または、 vigilant init を実行します。
あなたのために）。正確なオプションについては、「モデル」を参照してください。
Slack と Teams のサポートも存在しますが、まだベータ版です。
GitHub CLI gh 、コメントを作成するユーザーとして認証される
( gh auth login )、またはプル リクエストの GH_TOKEN 環境変数: 読み取り/書き込み。
サポートされているモデルプロバイダーの API キー (無料、カードなしを含む)
層 (Groq、Google Gemini、NVIDIA NIM)。 「モデル」を参照してください。
コア エンジンには依存関係がありません (標準ライブラリのみ) - モデル呼び出しはオーバーライドされます。
プレーンな HTTP、SDK なし。
vigilent init を 1 回実行します。 GitHub アカウントに接続します (gh auth ログインを実行します)
必要に応じて)、モデル キーを保存します。編集するファイルはありません。キーは
~/.config/vigilant-pr/credentials.json の 0600 ファイルに保存されます (同じ
gh および aws CLI としての姿勢)。
model コマンドを使用してモデルを管理および切り替えます。
警戒モデル add # プロバイダーを選択し、キーを貼り付けます (保存され、アクティブになります)
警戒モデルは groq # を追加するか、プロバイダーに直接名前を付けます
警戒モデルリスト # 保存されたプロバイダー、マスクされたキー、およびアクティブなプロバイダーを参照
vigilant model use groq # アクティブなモデルを切り替えます (プロバイダーまたはプロバイダー/モデルごとに)
警戒モデルremove openai # 保存されたキーを削除する
ファイル/CI を優先しますか? .env ファイルと実際の環境変数は引き続き機能し、
常に保存されたキーよりも優先されます (実際の環境 > .env > ストア)。コピー
.env.example を .env に変更し、使用するものを入力します。モデルを設定しない場合、Vigilant
見つかったプロバイダー キーから 1 つを自動選択します (Anthropic が推奨)。
選択したモデルを出力します。
pipx インストール git+https://github.com/tll

ongdev/警戒-pr
# または特定のリリースを固定します:
pipx install git+https://github.com/tllongdev/vigilant-pr@v1.1.0
# または、クローンから:
UVツールをインストールします。
main は安定したリリースラインであるため、固定されていないコマンドは常に
最新リリース。
所有しているものを確認し、保存されている設定と API キーを適切にアップグレードします。
~/.config/vigilant-pr/ に存在し、アップグレードの影響を受けることはありません。
警戒 --version
# pipx: 最新のメイン (最も信頼性の高い) から強制的に再インストールします
pipx install --force git+https://github.com/tllongdev/vigilant-pr
＃紫外線：
uv ツールのインストール --force git+https://github.com/tllongdev/vigilant-pr
# コンテナ:
docker pull ghcr.io/tllongdev/vigilant-pr:latest
--force を使用します。 pipx アップグレード / uv ツール アップグレードはバージョン文字列を比較するため、
単純なアップグレードでは、「すでに最新である」と報告され、新しいメイン コミットがスキップされる可能性があります。
それはバージョンを上げませんでした - そして固定された @vX.Y.Z インストールはまったく移動しません。
強制再インストールでは、常に現在のコードがプルされます。代わりに固定するには、次のように再インストールします
明示的なタグ: pipx install --force git+https://github.com/tllongdev/vigilant-pr@v1.7.0 。
ゼロから PR を自動レビューするまで、3 つのコマンドで実行できます。
pipx インストール git+https://github.com/tllongdev/vigilant-pr
vigilant init # GitHub に接続し、モデルキーを選択して保存します (最初に無料オプション)
vigilant github-watch # あなたがレビュー担当者であるオープン PR を自動レビューします
vigilant init はすべてを説明します: GitHub アカウントに接続します
(必要に応じて gh auth login を実行します)、モデルプロバイダーを選択できます
(Groq のようなクレジット カード不要の無料オプションが主流)、キーを検証し、
それを保管します。手動で編集する必要はありません。後で慎重にモデルを使用してモデルを切り替えてください。
何かを投稿する前にレビューを見たいですか?まず PR をドライランします。
警戒レビュー https://github.com/owner/repo/pull/123 --dry-run
それは

全体の流れ: インストール、初期化、監視。以下はすべて参考になります
特定のモデル、ウォッチャーのチューニング、チャット サーフェスなど。
Vigilant PR は現在無料でご利用いただけます。
ワークフローに価値を加える場合、寄付は継続的な開発とメンテナンスに直接使われます。
Apple Pay と Google Pay をサポート - ワンタップ、カード番号なし。
# PR をレビューし、あなたとして投稿します (Sonnet 5、デフォルト層)
警戒レビュー https://github.com/owner/repo/pull/123
# ハード PR のために Opus 4.8 にエスカレーションする
警戒レビュー 123 --repo 所有者/repo --opus
# 投稿せずにプレビューする
警戒レビュー 123 --repo 所有者/repo --dry-run
# プレビューして、投稿する前に承認します (新しいモデルを試すときに最適です)
警戒レビュー 123 --repo 所有者/repo --approve
投稿前にレビュー（承認ゲート）
デフォルトでは、レビューは自動的に投稿されます。なじみのないモデルを試している場合 - または
信頼する前に、何が生成されるかを確認したいだけです - 承認をオンにしてください
ゲート: Vigilant は完全なレビュー (概要 + インライン コメント) を出力し、
何かが投稿される前に y/N を入力します。
1 回限り: --approve をレビューまたは github-watch に追加します。
常にオン: VIGILANT_REQUIRE_APPROVAL=1 を設定します (または vigilant init で「はい」と答えます)。
オフに戻します: --no-approve または VIGILANT_REQUIRE_APPROVAL=0 。
モデルを信頼したら、フラグを外して、モデルに代わって投稿させます。
Vigilant PR はモデルに依存しません。プロバイダー/モデル文字列を使用してモデルを選択します。
--model または VIGILANT_MODEL 環境変数を指定し、そのプロバイダーのキーを指定します。裸の
名前 (例: claude-sonnet-5 ) は Anthropic として扱われるため、既存のセットアップは維持されます。
働いています。内部には、Anthropic という 2 つのワイヤー プロトコルだけがあります。
メッセージ API と OpenAI 互換の /chat/completions API - ほとんどの
プロバイダー、ローカル サーバー、ゲートウェイはすぐに使用できます。
モデルの前面に OpenAI 互換の AI ゲートがある場合

eway または LLM プロキシ
(LiteLLM、Portkey、Cloudflare AI Gateway、Kong、セルフホスト型プロキシ、または
内部エンタープライズ ゲートウェイ - 多くの場合、集中管理された低コストのアクセス用)、
point ゲートウェイプロバイダーと連携して警戒してください。完全にベンダー中立です - いいえ
ゲートウェイの名前はコードで指定されているため、エンドポイントと資格情報を指定するだけです。
最も簡単な方法は、ベース URL と認証の入力を求めるガイド付きウィザードです。
モードを選択し、すべてを管理された資格情報ストアに保存します。
警戒モデルはゲートウェイを追加します
または、環境変数 (または .env ) を使用して手動で構成します。モデルを設定する
とベース URL を選択し、認証モードを 1 つ選択します。
import VIGILANT_MODEL=ゲートウェイ/モデル名
エクスポート VIGILANT_API_BASE=https://gateway.example.com/v1
# 認証 A: 静的ベアラー トークン
エクスポート VIGILANT_API_KEY=...
# 認証 B: OAuth2 クライアント認証情報 (トークンが取得、キャッシュされ、自動更新されます)
エクスポート VIGILANT_OAUTH_TOKEN_URL=https://auth.example.com/oauth/token
エクスポート VIGILANT_OAUTH_CLIENT_ID=...
エクスポート VIGILANT_OAUTH_CLIENT_SECRET=...
# オプション:
import VIGILANT_OAUTH_SCOPE=... # スコープ (IdP が必要とする場合)
import VIGILANT_OAUTH_AUDIENCE=... # 件の対象者 (IdP が必要とする場合)
import VIGILANT_OAUTH_AUTH_STYLE=basic # クライアント ID/シークレットを HTTP Basic として送信します (デフォルト: 本文)
無料利用枠では、約 2 分で開始できます。
Groq (最速): https://console.groq.com/keys (キーは gsk_ で始まります)
双子座 : https://aistudio.google.com/apikey
NVIDIA NIM : https://build.nvidia.com (キーは nvapi- で始まります)
エクスポート GROQ_API_KEY=gsk_...
エクスポート VIGILANT_MODEL=groq/llama-3.3-70b-versatile
警戒レビュー https://github.com/owner/repo/pull/123
警戒モデルを実行して、資格情報がどのプロバイダーに到達できるのか (そして、どこに到達できるのか) を確認します。
プロバイダーはリスト エンドポイント、つまり使用できる正確なモデル ID を公開します)。
最も詳細なレビューを行うには、フロンティア モデルを使用します。アドヴェル

サリアルの虫探し
モデルの品質に応じてスケールします。 Claude Sonnet 5 (デフォルト) または Opus 4.8 はより繊細な印象を与えます
小規模な無料モデルよりも問題が発生します。無料利用枠は試してみるのに最適です。
より軽い PR に。拡張思考チューニング (Opus 適応的思考) のみが適用されます
人間の道へ。他のプロバイダーは低いレビュー温度で実行されます。
vigilant github-watch は、要求されているオープン PR を GitHub にポーリングします。
レビュー担当者があなたの代わりにそれらを自動レビューします。べき等です (決して再レビューしないでください)
同じヘッド SHA）、制限付き（ポーリング間隔 + 1 日あたりの上限）、復元力のある（
1 つの PR で障害が発生してもループがクラッシュすることはありません)。 GitHub アプリや Webhook は必要ありません - 必要なのはあなたのものだけです
トークン。 (古い名前の vigilant watch は今でも別名として機能します。)
# 継続的に実行 (デフォルト: 120 秒ごとにポーリング、UTC 日あたり 50 レビューを上限)
警戒中のgithub-watch
# 1 回のパスと終了 - cron に最適
vigilant github-watch --once
# ケイデンスとキャップを調整する
vigilant github-watch --poll-interval 300 --daily-cap 20
接触するリポジトリのスコープ設定
デフォルトでは、ウォッチャーはリクエストされたすべての PR をレビューします。で制約します
環境変数 (カンマ区切り)。常に否定が勝ちます。空ではない許可リストは
排他的:
import VIGILANT_ORG_ALLOW= " acme,acme-labs " # これらの組織のみ
輸出警戒

[切り捨てられた]

## Original Extract

Adversarial AI code review, posted as you - reviews pull requests on your behalf under your own identity, not a bot. Bring your own model. - tllongdev/vigilant-pr

GitHub - tllongdev/vigilant-pr: Adversarial AI code review, posted as you - reviews pull requests on your behalf under your own identity, not a bot. Bring your own model. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
tllongdev
/
vigilant-pr
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
develop Branches Tags Go to file Code Open more actions menu Folders and files
61 Commits 61 Commits .claude/ skills/ vigilant-pr .claude/ skills/ vigilant-pr .cursor/ skills/ vigilant-pr .cursor/ skills/ vigilant-pr .github .github assets assets docs docs src/ vigilant src/ vigilant tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
A portable, workflow-agnostic AI pull-request reviewer that posts review comments on behalf of you - your GitHub identity, not a generic bot.
Get tagged as a reviewer on a pull request, and Vigilant PR reviews it and posts
the comments as you - your GitHub identity, not a bot. No repo-side setup, no
GitHub App, just your token.
flowchart LR
CFG(["Configure in the terminal:<br/>vigilant init"]) --> C1
CFG -.-> C2
C1(["vigilant github-watch"]) --> A["You are tagged as a<br/>reviewer on a GitHub PR"]
C2(["vigilant slack-watch<br/>or vigilant teams-watch"]) -.-> A2["Beta: you are @-mentioned in a<br/>configured Slack or Teams channel"]
A --> B["Vigilant PR reads<br/>the PR and diff"]
A2 -.-> B
B --> C["Adversarial review,<br/>severity-tagged findings"]
C --> D["Posts inline comments<br/>as you, not a bot"]
D --> E["Approves if nothing blocks,<br/>comments if it does"]
classDef beta stroke:#a98bff,color:#a98bff,stroke-dasharray:5 5;
classDef cmd fill:#0d2038,stroke:#4da3ff,color:#cfe6ff;
class A2 beta;
class C2 beta;
class CFG cmd;
class C1 cmd;
Loading
Status
v1 - ready to use for reviewing GitHub pull requests.
Use any AI model you want. Give Vigilant PR an API key for your preferred
provider and it uses that model - Claude, GPT, Gemini, Grok (xAI), Llama,
NVIDIA, OpenRouter, or a model you run locally. Some are free, some paid - your choice.
You just set your key and pick a model (or run vigilant init , which does it
for you). See Models for the exact options.
Slack and Teams support also exists, but is still beta.
The GitHub CLI gh , authenticated as the user who should author the comments
( gh auth login ), or a GH_TOKEN env var with Pull requests: read/write.
An API key for any supported model provider - including free, no-card
tiers (Groq, Google Gemini, NVIDIA NIM). See Models .
The core engine is dependency-free (standard library only) - model calls go over
plain HTTP, no SDKs.
Run vigilant init once. It connects your GitHub account (runs gh auth login
for you if needed) and stores a model key - there are no files to edit. Keys are
kept in a 0600 file at ~/.config/vigilant-pr/credentials.json (the same
posture as the gh and aws CLIs).
Manage and switch models with the model command:
vigilant model add # pick a provider, paste a key (stored, becomes active)
vigilant model add groq # or name the provider directly
vigilant model list # see stored providers, masked keys, and the active one
vigilant model use groq # switch the active model (by provider or provider/model)
vigilant model remove openai # delete a stored key
Prefer files/CI? A .env file and real environment variables still work and
always take precedence over the stored keys (real env > .env > store). Copy
.env.example to .env and fill in what you use. If you set no model, Vigilant
auto-selects one from whichever provider key it finds (Anthropic preferred) and
prints which model it chose.
pipx install git+https://github.com/tllongdev/vigilant-pr
# or pin a specific release:
pipx install git+https://github.com/tllongdev/vigilant-pr@v1.1.0
# or, from a clone:
uv tool install .
main is the stable release line, so the unpinned command always gets the
latest release.
Check what you have, then upgrade in place - your stored config and API keys
live in ~/.config/vigilant-pr/ and are never touched by an upgrade:
vigilant --version
# pipx: force a reinstall from the latest main (most reliable)
pipx install --force git+https://github.com/tllongdev/vigilant-pr
# uv:
uv tool install --force git+https://github.com/tllongdev/vigilant-pr
# container:
docker pull ghcr.io/tllongdev/vigilant-pr:latest
Use --force . pipx upgrade / uv tool upgrade compare version strings, so a
plain upgrade can report "already up to date" and skip a newer main commit
that didn't bump the version - and a pinned @vX.Y.Z install won't move at all.
A forced reinstall always pulls the current code. To pin instead, reinstall with
an explicit tag: pipx install --force git+https://github.com/tllongdev/vigilant-pr@v1.7.0 .
From zero to auto-reviewing PRs as you, in three commands:
pipx install git+https://github.com/tllongdev/vigilant-pr
vigilant init # connects GitHub, picks + stores a model key (free options first)
vigilant github-watch # auto-reviews any open PR where you're a requested reviewer
vigilant init walks you through everything: it connects your GitHub account
(running gh auth login for you if needed), lets you pick a model provider
(leading with free, no-credit-card options like Groq), validates the key, and
stores it. Nothing to hand-edit; switch models later with vigilant model use .
Want to see a review before it posts anything? Dry-run any PR first:
vigilant review https://github.com/owner/repo/pull/123 --dry-run
That's the whole flow: install, init , watch. Everything below is reference for
specific models, watcher tuning, and the chat surfaces.
Vigilant PR is free to use right now.
If it adds value to your workflow, donations go directly toward its continued development and maintenance.
Apple Pay and Google Pay supported - one tap, no card number.
# Review a PR and post as you (Sonnet 5, the default tier)
vigilant review https://github.com/owner/repo/pull/123
# Escalate to Opus 4.8 for a hard PR
vigilant review 123 --repo owner/repo --opus
# Preview without posting
vigilant review 123 --repo owner/repo --dry-run
# Preview, then approve before it posts (great while trying a new model)
vigilant review 123 --repo owner/repo --approve
Review before it posts (approval gate)
By default reviews post automatically. If you're trying an unfamiliar model - or
just want to watch what it produces before trusting it - turn on the approval
gate: Vigilant prints the full review (summary + inline comments) and asks for a
y/N before anything is posted.
One-off: add --approve to review or github-watch .
Always on: set VIGILANT_REQUIRE_APPROVAL=1 (or answer "yes" in vigilant init ).
Turn it back off: --no-approve or VIGILANT_REQUIRE_APPROVAL=0 .
Once you trust the model, drop the flag and let it post on your behalf.
Vigilant PR is model-agnostic. Pick a model with a provider/model string via
--model or the VIGILANT_MODEL env var, and supply that provider's key. A bare
name (e.g. claude-sonnet-5 ) is treated as Anthropic, so existing setups keep
working. Under the hood there are just two wire protocols - the Anthropic
Messages API and the OpenAI-compatible /chat/completions API - so most
providers, local servers, and gateways work out of the box.
If your models are fronted by an OpenAI-compatible AI gateway or LLM proxy
(LiteLLM, Portkey, Cloudflare AI Gateway, Kong, a self-hosted proxy, or an
internal enterprise gateway - often for centrally-managed, lower-cost access),
point Vigilant at it with the gateway provider. It's fully vendor-neutral - no
gateway is named in code, you just supply the endpoint and credentials.
The easiest path is the guided wizard, which prompts for the base URL and auth
mode and saves everything to the managed credential store:
vigilant model add gateway
Or configure it manually with environment variables (or a .env ). Set the model
and base URL, then pick one auth mode:
export VIGILANT_MODEL=gateway/your-model-name
export VIGILANT_API_BASE=https://gateway.example.com/v1
# Auth A: a static bearer token
export VIGILANT_API_KEY=...
# Auth B: OAuth2 client-credentials (token is fetched, cached, and auto-refreshed)
export VIGILANT_OAUTH_TOKEN_URL=https://auth.example.com/oauth/token
export VIGILANT_OAUTH_CLIENT_ID=...
export VIGILANT_OAUTH_CLIENT_SECRET=...
# optional:
export VIGILANT_OAUTH_SCOPE=... # scope, if your IdP requires one
export VIGILANT_OAUTH_AUDIENCE=... # audience, if your IdP requires one
export VIGILANT_OAUTH_AUTH_STYLE=basic # send client id/secret as HTTP Basic (default: body)
Free tiers get you started in ~2 minutes:
Groq (fastest): https://console.groq.com/keys (key starts with gsk_ )
Gemini : https://aistudio.google.com/apikey
NVIDIA NIM : https://build.nvidia.com (key starts with nvapi- )
export GROQ_API_KEY=gsk_...
export VIGILANT_MODEL=groq/llama-3.3-70b-versatile
vigilant review https://github.com/owner/repo/pull/123
Run vigilant models to see which providers your credentials can reach (and, where
the provider exposes a list endpoint, the exact model ids you can use).
For the deepest reviews, use a frontier model. Adversarial bug-finding
scales with model quality; Claude Sonnet 5 (default) or Opus 4.8 catch subtler
issues than small free models. The free tiers are great for trying it out and
for lighter PRs. Extended-thinking tuning (Opus adaptive thinking) applies only
to the Anthropic path; other providers run with a low review temperature.
vigilant github-watch polls GitHub for open PRs where you are a requested
reviewer and auto-reviews them on your behalf. It is idempotent (never re-reviews
the same head SHA), bounded (poll interval + per-day cap), and resilient (a
failure on one PR never crashes the loop). No GitHub App, no webhooks - just your
token. (The old name vigilant watch still works as an alias.)
# Run continuously (default: poll every 120s, cap 50 reviews/UTC-day)
vigilant github-watch
# One pass and exit - ideal for cron
vigilant github-watch --once
# Tune cadence and cap
vigilant github-watch --poll-interval 300 --daily-cap 20
Scoping which repos it touches
By default the watcher reviews any PR you are requested on. Constrain it with
env vars (comma-separated). Deny always wins; a non-empty allow list is
exclusive:
export VIGILANT_ORG_ALLOW= " acme,acme-labs " # only these orgs
export VIGILANT

[truncated]
