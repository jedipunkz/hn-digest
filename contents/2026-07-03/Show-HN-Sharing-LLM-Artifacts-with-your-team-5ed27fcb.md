---
source: "https://github.com/danielpang/dropway"
hn_url: "https://news.ycombinator.com/item?id=48779570"
title: "Show HN: Sharing LLM Artifacts with your team"
article_title: "GitHub - danielpang/dropway · GitHub"
author: "d_pang"
captured_at: "2026-07-03T20:53:11Z"
capture_tool: "hn-digest"
hn_id: 48779570
score: 2
comments: 0
posted_at: "2026-07-03T20:26:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sharing LLM Artifacts with your team

- HN: [48779570](https://news.ycombinator.com/item?id=48779570)
- Source: [github.com](https://github.com/danielpang/dropway)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T20:26:42Z

## Translation

タイトル: HN を表示: LLM アーティファクトをチームと共有する
記事タイトル: GitHub - danielpang/dropway · GitHub
説明: GitHub でアカウントを作成して、danielpang/dropway の開発に貢献します。
HN テキスト: 技術者以外のチームがクロード・コワークと協力し始めたときに発見した共通の課題は、HTML/マークダウン ファイルをチームの他のメンバーと共有する機能でした。 LLM を使用すると、構築は簡単ですが、成果物の共有は依然として困難です。プロトタイプやレポートを取得するには、Cloudflare や Vercel などの技術プラットフォームを使用して単純な Web ページをホストする必要があります。これらのページの設定は、技術者以外の人にとっては複雑すぎて、エンジニアリング時間の無駄でした。 Claude と ChatGPT は最近、アーティファクトを共有する機能を追加しましたが、アクセス制御の柔軟性に欠けており、ファイルをプラットフォームにロックしてしまいます。 Dropway は簡単な修正です。ファイルをダッシュ​​ボードにドラッグし (または CLI または MCP 経由でデプロイし)、数秒で共有可能な URL を取得します。アクセス制御はサイトごとに設定されます: プライベート電子メール許可リスト、組織のみ、またはパブリック。すべてのデプロイは不変であり、コンテンツに対応しているため、バージョン管理は自動的に行われ、ロールバックはワンクリックで実行できます。 CLI と MCP の統合を追加したため、Claude、Codex、および Cursor をサイトに直接展開できるようになりました。選択した LLM を離れることなく、作成、共有、反復を行うことができます。ソースは FSL で利用可能で、 https://github.com/danielpang/dropway で自己ホスト可能です。または、https://dropway.dev でホストされているオプションを使用してください。アーキテクチャや MCP 統合に関する質問に喜んでお答えします。同じ問題に遭遇したことがある方からのフィードバックをお待ちしています。よろしくお願いします！

記事本文:
GitHub - ダニエルパン/ドロップウェイ · GitHub
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
ダニエルパン
/
ドロップウェイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

レ
177 コミット 177 コミット .github/ workflows .github/ workflows apps/ ダッシュボード apps/ ダッシュボード cli cli クラウド クラウド契約 契約 db db デプロイ デプロイ ドキュメント ドキュメント エッジ/ サービングワーカー エッジ/ サービングワーカー ee ee 例 例 内部 内部パッケージ パッケージ サービス サービス testdata/ マークダウン testdata/ マークダウン .dockerignore .dockerignore .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md fly.mcp.toml fly.mcp.toml fly.toml fly.toml go.mod go.mod go.sum go.sum install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ファイルのフォルダー → アクセス制御されたライブ URL を 1 つのコマンドで作成します。
Dropway はあらゆる静的サイト (HTML/CSS レポート、データ ダッシュボード、React、
Vite ビルド、ドキュメント サイト、または AI 生成ページ）を、ライブでバージョン管理された共有可能なファイルに変換します。
URLを数秒で表示します。フォルダーをダッシュボードにドラッグするか、MCP 経由でアップロードします。
そして、あなたが管理する実際の URL を取得します。それをインターネット全体、あなたのチーム、そして
会社、または特定の少数の人々。いつでも考えを変えてください。
これはオープンソースで自己ホスト可能であり、オプションで Dropway.dev でホストされた SaaS を使用できます。
何かを構築するのは今では簡単です。適切な人と安全に共有し、
まだ難しいです。レポート、プロトタイプ、ダッシュボード、またはページを生成しました。
そしてあなたのオプションはあまり良くありません:
zip / スクリーンショットを電子メールで送信: ライブ URL がなく、すぐに古くなり、更新できません。
S3 + CloudFront を起動する: 技術者以外の人にとっては簡単ではありません。
Vercel/ Netlify プロジェクト : Github へのアクセスが必要で、チームだけと共有することはできません。
Wiki に貼り付ける: 実際のレイアウト、対話性、CSS/JS が失われます。
内部ファイル共有: 誰かに送信できるリンクはありません

チームの一員、有効期限なし、
誰が何を見たのかの監査はありません。
これらのオプションのいずれも、「この 3 人と共有する」、または「次のメンバー全員と共有する」ことはできません。
私の会社」または「公開されていますが、パスワードで保護されています」を選択し、それを取り消すか、ロールします。
悪いバージョンに戻します。
Dropway は欠けているレイヤーです。シンプルかつ迅速に公開でき、共有とアクセスが可能です。
最高級の機能として制御します。
ドラッグ アンド ドロップで、フォルダーをライブ URL にデプロイします。パイプラインも構成もありません。事前に構築された静的出力はそのまま機能します。
サイトごとの 4 つの共有層: パブリック、パスワード保護、特定
ユーザー (電子メールの許可リスト)、または組織内のすべてのユーザー。デフォルトの拒否。あなたはオプトインします。
マルチテナント組織: チーム、ロール (オーナー/管理者/メンバー)、組織全体のポリシー
これにより、社外への共有が完全に禁止される場合があります。
不変でバージョン管理されたデプロイ : すべてのデプロイはコンテンツに対応しており、インスタント
以前のバージョンにロールバックします。
共有者の監査ログ
そして何にアクセスしたのか。
即時取り消し: メンバーを削除するか、サイトの共有を解除すると、そのメンバーのアクセスは無効になります。
トークンの有効期限が切れるたびではなく、端ですぐにカットされます。
信頼できないコンテンツを安全に提供: テナント HTML/JS は別のコンテンツから提供されます。
Public-Suffix-List ドメインなので、あるサイトが別のサイト (またはあなたの) セッションに到達することはできません。
LLM フレンドリーなアクセス : パブリック サイトは自動サービスを提供します。
llms.txt インデックスとウェルカム AI クローラー (GPTBot、ClaudeBot、
PerplexityBot など) を使用して、エージェントがコンテンツを検出して読み取ることができます。ゲート付きサイト
(組織のみ / 許可リスト / パスワード) クローラーに対して立ち入りを禁止します。 LLM が彼らに届く
認証された Dropway MCP サーバー (OAuth 2.1; に接続する) 経由のみ
クロード、カーソル、またはコーデックスなど）、アクセス制御は AI に対しても同様に適用されます。
人々。所有者/管理者は、設定で組織ごとに MCP アクセスをオフに切り替えることができます。
オープンソース + 自己ホスト可能 : 自己ホストするか、dropway.dev で当社のサービスを使用します。
エンジニアとデータ/分析

生成されたレポート、ノートブックを HTML として共有するチーム、
ベンチマーク ダッシュボードと 1 回限りのツール。社内またはクライアントと。
財務/会計の専門家が分析やレポートを共有できます。
静的プロトタイプ、設計仕様、ビルドのレビューを共有するデザイナーと PM
必要に応じてパスワードで保護された、正確な関係者が閲覧できるようにします。
Web サイトを生成し、ユーザーに実際の Web サイトを渡す必要がある AI アプリ/エージェント ビルダー
プログラムによってアクセス制御された URL。
管理された共有が必要な企業: 「デフォルトでは内部、場合によっては外部のみ」
管理者がそれを許可します。」と、ロール、監査、カスタム ドメイン、即時取り消しが含まれます。
ダッシュボードにファイルをアップロード ─▶ Go API (レコード システム + 認証) ─▶ R2/minIO (BLOB ストレージ)
│ 再構築可能な配線投影を書き込みます
▼
ブラウザ ─▶ Cloudflare エッジ ワーカー (*.dropwaycontent.com) ─▶ サイトをストリーミングします
public = ログインなし、キャッシュ可能 · ゲート = /authz からのホストスコープのトークン
Next.js ダッシュボード (Better Auth: Google / email / magic-link) がコントロールです
飛行機; Go API は記録システムであり、承認境界です。ある
Cloudflare Workerはエッジでコンテンツを提供します。 Postgres (行レベルのセキュリティ付き)
組織ごと）が真実の情報源です。コンポーネント + リクエストフロー
図は docs/diagrams/ にあります。
Docker (Compose を含む) が必要です。 1 つのコマンドでスタック全体が構築され、開始されます。
Go API、Next.js ダッシュボード、バンドルされた Postgres、バンドルされた MinIO (R2/S3)
代役)、およびスキーマの移行。
git clone https://github.com/your-org/dropway.git && cd ドロップウェイ
cpdeploy/.env.exampledeploy/.env #安全なローカル開発のデフォルト
docker compose -fdeploy/docker-compose.yml up --build
サービス
URL
ダッシュボード（ここからサインアップ）
http://ローカルホスト:3000
API (Go コントロールプレーン)
http://localhost:8080 ( GET /healthz )
Postgres / MinIO (バンドル)
ローカルホスト:5432 / co

nsole http://localhost:9001
最初の開始後に 1 回。 Better Auth は ID テーブルを所有しているので、それらを作成します。
docker compose -f デプロイ/docker-compose.yml 実行ダッシュボード \
pnpm dlx @better-auth/cli@latest 移行 --yes
次に、 http://localhost:3000 を開いてサインアップし、組織を作成し、サイトを作成してデプロイします。
最初のフォルダー。
独自の Postgres / オブジェクト ストアを使用する
バンドルされている Postgres と MinIO は、オプションの Compose プロファイルです。 Supabase を指す /
外部 Postgres または Cloudflare R2 / S3 の場合は、COMPOSE_PROFILES からプロファイルをドロップします
そして、一致する DATABASE_URL / S3_* 変数をdeploy/.envに設定します。
COMPOSE_PROFILES= DATABASE_URL=postgres://… S3_ENDPOINT=https://… \
docker compose -fdeploy/docker-compose.yml up --build api ダッシュボードの移行
サイトが提供される場所 (コンテンツ ドメイン)
公開サイトは <org-slug>--<app-slug>.<CONTENT_DOMAIN> で提供されます (org-namespaced)。
したがって、2 つの組織の両方が、衝突することなく blog という名前のアプリを持つことができます。 3 つのバール
deploy/.env は、ダッシュボードと CLI から返される URL を制御します。
ローカルのデフォルトでは、すべてのデプロイがクリック可能な http://<org>--<app>.localhost:8090/ になります。
*.localhost は、DNS が設定されていないすべてのブラウザで 127.0.0.1 に解決されます。奉仕する
独自のドメインで、ワイルドカード DNS レコードと TLS 証明書 ( *.your-domain ) を指定します。
コンテンツサーバーとセット:
CONTENT_DOMAIN=apps.example.com # あなたが管理するドメイン
CONTENT_SCHEME=https
CONTENT_PORT= # 空 → 標準 :443
これらの変数は表示される URL にのみ影響します。serve は Host ヘッダーとストリップに一致します。
そのため、保存されたルートはベアホストのままになります。詳細 (ワイルドカード DNS/証明書、
Public-Suffix-List 分離ルール、およびゲート サイトの https 要件) は次のとおりです。
デプロイ/README.md 。
Dropway は、OAuth で保護された MCP サーバー (mcp サービス、http://localhost:8092) を出荷します。
ローカル）したがってLLM時代

サイトの閲覧や読み取りはできません: サイトのリスト、読み取りまたはダウンロード
サイトのファイル。 MCP サーバーを使用してサイトを作成し、展開することもできます +
ファイルをそこに公開し、サイトの共有設定を変更します (所有者/管理者のみ、同じルール)
ダッシュボードとして）。公開サイト
自動提供される llms.txt を介してクローラーでも読み取ることができます。
ゲートされたサイト (組織のみ / 許可リスト / パスワード) は、LLM によってのみアクセス可能です。
承認された MCP 接続を使用するため、AI に対しても人間と同様にアクセス制御が適用されます。
MCP URL を使用してカスタム コネクタとして AI ツールに追加すると、ダッシュボードに次のように表示されます。
[設定] → [LLM アクセス (MCP)] → [接続] の手順をコピーして貼り付けます。短いバージョン:
最初の接続ではブラウザが開き、サインインして「MCP アクセスの承認」を承認します。
これは標準の OAuth 2.1 です (コピーする API キーはありません)。アクセス範囲は組織内にとどまります
各サイトの共有設定が尊重されます。所有者/管理者は、MCP アクセスをオフに切り替えることができます。
組織全体の [設定] → [LLM アクセス (MCP)] をクリックします。運用環境では MCP_PUBLIC_URL / を設定します
MCP ホストへの NEXT_PUBLIC_MCP_URL (例: https://mcp.dropway.dev );参照してください
デプロイ/README.md 。
go build ./... && go test ./... # Go コア (API + CLI)
go run ./services/api/cmd/api # API を開始します
corepack Enable && pnpm install # TS ワークスペース
pnpm dev # ダッシュボードを実行する
完全なローカル開発リファレンス (ビルド フレーバー、エッジ ワーカー、手動での移行) は次の場所にあります。
デプロイ/README.md 。
オープンソース + ホスト型 (オープンコア)
Dropway は、誰もがソースを利用できるコードベースという人気の開発ツール モデルに従っています。
無料で自己ホストできるほか、利便性と拡張性を高めるためにオプションでホストされた SaaS を利用できます。
このコアは Functional Source License (FSL-1.1-Apache-2.0) の下にあります。
自己ホストし、変更し、内部で無料で使用します。として再販することはできません
競合するホスト型サービス。各リリースは Apache 2.0 になります

2年後。
クラウド/モジュール (ストライプ請求 + 使用量クォータ) は独自のものであり、決して
はセルフホスト ビルドに同梱されているため、セルフホストには制限がありません。 ee/ホールド
ライセンスゲート型エンタープライズ機能 (SSO/SAML、監査エクスポート、カスタム ドメイン)。
docs/diagrams/ : コンポーネント + シーケンス図 (サインアップ、
サインイン、展開、ゲート アクセス、LLM/MCP アクセス)。
COTRIBUTING.md : DCO の承認のもとでの貢献を歓迎します。
コア: FSL-1.1-Apache-2.0 (→ 2 年後の Apache 2.0)。雲/そして
ee/ は独自のライセンスによって管理されます。 「Dropway」の名前とロゴは予約されています
商標;
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
v0.2.0
最新の
2026 年 7 月 1 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to danielpang/dropway development by creating an account on GitHub.

A common challenge we found as our non technical team began to work with Claude Cowork was the ability to share HTML/markdown files with the rest of the team. With LLMs building is easy but sharing artifacts is still difficult. Getting a prototype or report meant using a technical platform like Cloudflare or Vercel to host simple web pages. Setting up these pages was too complex for non-technical folks and a waste of engineering time. While Claude and ChatGPT have recently added the ability to share artifacts, they lack flexibility with access control and lock your files into their platform. Dropway is a simple fix: drag your files into the dashboard (or deploy via CLI or MCP) and get a shareable URL in seconds. Access control is set per site: private email allowlist, org-only, or public. Every deploy is immutable and content-addressed, so versioning is automatic and rollback is one click. We added CLI and MCP integration so Claude, Codex, and Cursor can deploy directly to a site. You can create, share, iterate without leaving your LLM of choice. Source-available under FSL and self-hostable at https://github.com/danielpang/dropway . Or use our hosted option at https://dropway.dev Happy to answer questions about the architecture or the MCP integration, and would love feedback from anyone who has run into this same friction. Thank you in advance!

GitHub - danielpang/dropway · GitHub
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
danielpang
/
dropway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
177 Commits 177 Commits .github/ workflows .github/ workflows apps/ dashboard apps/ dashboard cli cli cloud cloud contracts contracts db db deploy deploy docs docs edge/ serving-worker edge/ serving-worker ee ee examples examples internal internal packages packages services services testdata/ markdown testdata/ markdown .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md fly.mcp.toml fly.mcp.toml fly.toml fly.toml go.mod go.mod go.sum go.sum install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
A folder of files → a live, access-controlled URL in one command.
Dropway turns any static site (for example an HTML/CSS report, a data dashboard, a React or
Vite build, a docs site, or an AI-generated page) into a live, versioned, shareable
URL in seconds. Drag a folder into the dashboard, or upload via MCP,
and you get a real URL you control: share it with the whole internet, your team and
company, or a few specific people. Change your mind anytime.
It's open source and self-hostable, with an optional hosted SaaS at dropway.dev .
Building something is easy now. Sharing it with the right people, safely,
is still difficult. You've generated a report, a prototype, a dashboard, or a page,
and your options aren't great:
Email a zip / screenshot : no live URL, instantly stale, impossible to update.
Spin up S3 + CloudFront : not simple for non-technical folks.
Vercel/ Netlify project : requires access to Github, can't share with just your team.
Paste into a wiki : loses the real layout, interactivity, and your CSS/JS.
Internal file shares : no link you can send someone outside the team, no expiry,
no audit of who saw what.
None of these options allow you to "share this with these three people," or "anyone at
my company," or "public, but password-protected," and then revoke it, or roll
back a bad version.
Dropway is that missing layer: simple and quick to publish, with sharing and access
control as first-class features.
Drag-and-drop to deploy : folder to live URL. No pipeline, no config. Pre-built static output just works.
Four sharing tiers, per site : public, password-protected, specific
people (email allowlist), or anyone in your org. Default-deny; you opt in.
Multi-tenant orgs : teams, roles (owner/admin/member), and an org-wide policy
that can forbid sharing outside the company entirely.
Immutable, versioned deploys : every deploy is content-addressed, with instant
rollback to any previous version.
Audit log of who shared
and accessed what.
Immediate revocation : remove a member or unshare a site and their access is
cut at the edge right away, not whenever a token happens to expire.
Safe to serve untrusted content : tenant HTML/JS is served from a separate
Public-Suffix-List domain, so one site can never reach another's (or your) session.
LLM-friendly access : public sites auto-serve an
llms.txt index and welcome AI crawlers (GPTBot, ClaudeBot,
PerplexityBot, and others), so agents can discover and read your content. Gated sites
(org-only / allowlist / password) stay off-limits to crawlers. LLMs reach them
only through the authenticated Dropway MCP server (OAuth 2.1; connect it to
Claude, Cursor, or Codex), so your access control holds for AI exactly as it does for
people. Owners/admins can switch MCP access off per org in Settings.
Open source + self-hostable : self-host or use our service at dropway.dev
Engineers & data/analytics teams sharing generated reports, notebooks-as-HTML,
benchmark dashboards, and one-off tools. Internally or with a client.
Financial/ accounting professionals can share analysis and reports.
Designers & PMs sharing static prototypes, design specs, and review builds with
the exact stakeholders who should see them, password-protected if needed.
AI app / agent builders that generate websites and need to hand a user a real,
access-controlled URL programmatically.
Companies that need governed sharing: "internal by default, external only if
an admin allows it," with roles, audit, custom domains, and instant revocation.
upload files in dashboard ─▶ Go API (system of record + authz) ─▶ R2/minIO (blob storage)
│ writes a rebuildable routing projection
▼
browser ─▶ Cloudflare edge Worker (*.dropwaycontent.com) ─▶ streams your site
public = no login, cacheable · gated = host-scoped token from /authz
A Next.js dashboard (with Better Auth: Google / email / magic-link) is the control
plane; a Go API is the system of record and the authorization boundary; a
Cloudflare Worker serves content at the edge; Postgres (with row-level security
per org) is the source of truth. The component + request-flow
diagrams are in docs/diagrams/ .
You need Docker (with Compose). One command builds and starts the whole stack:
the Go API, the Next.js dashboard, a bundled Postgres, a bundled MinIO (an R2/S3
stand-in), and the schema migrations.
git clone https://github.com/your-org/dropway.git && cd dropway
cp deploy/.env.example deploy/.env # safe local-dev defaults
docker compose -f deploy/docker-compose.yml up --build
Service
URL
Dashboard (sign up here)
http://localhost:3000
API (Go control-plane)
http://localhost:8080 ( GET /healthz )
Postgres / MinIO (bundled)
localhost:5432 / console http://localhost:9001
One time, after the first start. Better Auth owns the identity tables, so create them:
docker compose -f deploy/docker-compose.yml exec dashboard \
pnpm dlx @better-auth/cli@latest migrate --yes
Then open http://localhost:3000 , sign up, create an org, create a site, and deploy
your first folder.
Use your own Postgres / object store
The bundled Postgres and MinIO are optional Compose profiles. To point at Supabase /
an external Postgres or Cloudflare R2 / S3, drop the profile from COMPOSE_PROFILES
and set the matching DATABASE_URL / S3_* vars in deploy/.env :
COMPOSE_PROFILES= DATABASE_URL=postgres://… S3_ENDPOINT=https://… \
docker compose -f deploy/docker-compose.yml up --build api dashboard migrate
Where your sites are served (the content domain)
Published sites are served at <org-slug>--<app-slug>.<CONTENT_DOMAIN> , which is org-namespaced,
so two orgs can both have an app named blog without colliding. Three vars in
deploy/.env control the URL the dashboard and CLI hand back:
The local defaults make every deploy a clickable http://<org>--<app>.localhost:8090/ .
*.localhost resolves to 127.0.0.1 in every browser with no DNS setup. To serve
on your own domain, point a wildcard DNS record and TLS cert ( *.your-domain ) at the
content server and set:
CONTENT_DOMAIN=apps.example.com # a domain you control
CONTENT_SCHEME=https
CONTENT_PORT= # empty → standard :443
These vars affect only the displayed URL: serve matches the Host header and strips
the port, so the stored route stays the bare host. Full details (wildcard DNS/cert, the
Public-Suffix-List isolation rule, and the https requirement for gated sites) are in
deploy/README.md .
Dropway ships an OAuth-protected MCP server (the mcp service, http://localhost:8092
locally) so an LLM agent can browse and read your sites: list sites, read or download a
site's files. You can also use the MCP server to create a site, deploy +
publish files to it, and change a site's sharing settings (owner/admin only, same rules
as the dashboard). Public sites
are also readable by crawlers via an auto-served llms.txt ;
gated sites (org-only / allowlist / password) are reachable by an LLM only through
an authorized MCP connection, so your access control holds for AI exactly as it does for people.
Add it to your AI tool as a custom connector using the MCP URL, the dashboard shows
copy-paste steps under Settings → LLM access (MCP) → Connect . The short version:
The first connection opens your browser to sign in and approve "Authorize MCP access" .
That's standard OAuth 2.1 (no API keys to copy). Access stays scoped to your organization
and honors each site's sharing settings. Owners/admins can switch MCP access off for the
whole org under Settings → LLM access (MCP) . In production set MCP_PUBLIC_URL /
NEXT_PUBLIC_MCP_URL to your MCP host (e.g. https://mcp.dropway.dev ); see
deploy/README.md .
go build ./... && go test ./... # the Go core (API + CLI)
go run ./services/api/cmd/api # start the API
corepack enable && pnpm install # the TS workspace
pnpm dev # run the dashboard
Full local-dev reference (build flavors, the edge Worker, migrating by hand) lives in
deploy/README.md .
Open source + hosted (open-core)
Dropway follows popular Dev tools model: a source-available codebase anyone
can self-host for free, plus an optional hosted SaaS for convenience and scale.
The core is under the Functional Source License (FSL-1.1-Apache-2.0) :
self-host, modify, and use it internally for free; you just can't resell it as a
competing hosted service. Each release becomes Apache 2.0 after two years.
The cloud/ module (Stripe billing + usage quotas) is proprietary and never
ships in the self-host build , so self-host has no limits. ee/ holds
license-gated enterprise features (SSO/SAML, audit export, custom domains).
docs/diagrams/ : component + sequence diagrams (sign-up,
sign-in, deploy, gated access, LLM/MCP access).
CONTRIBUTING.md : contributions welcome under a DCO sign-off.
Core: FSL-1.1-Apache-2.0 (→ Apache 2.0 after two years). cloud/ and
ee/ are governed by their own licenses. The "Dropway" name and logo are reserved
trademarks;
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
v0.2.0
Latest
Jul 1, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
