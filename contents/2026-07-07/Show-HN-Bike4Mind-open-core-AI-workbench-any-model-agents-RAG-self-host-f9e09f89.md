---
source: "https://github.com/Bike4Mind/bike4mind"
hn_url: "https://news.ycombinator.com/item?id=48822349"
title: "Show HN: Bike4Mind – open-core AI workbench; any model, agents, RAG, self-host"
article_title: "GitHub - Bike4Mind/bike4mind: The open-core AI workbench — notebooks, agents, RAG, voice, and images across any model: OpenAI, Anthropic, Google, xAI, or local via Ollama/vLLM. BSL 1.1, auto-converting to Apache-2.0 on a two-year clock. Your AI keeps running when theirs doesn't. · GitHub"
author: "erikbethke"
captured_at: "2026-07-07T19:43:14Z"
capture_tool: "hn-digest"
hn_id: 48822349
score: 1
comments: 1
posted_at: "2026-07-07T19:17:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Bike4Mind – open-core AI workbench; any model, agents, RAG, self-host

- HN: [48822349](https://news.ycombinator.com/item?id=48822349)
- Source: [github.com](https://github.com/Bike4Mind/bike4mind)
- Score: 1
- Comments: 1
- Posted: 2026-07-07T19:17:34Z

## Translation

タイトル: 表示 HN: Bike4Mind – オープンコア AI ワークベンチ。任意のモデル、エージェント、RAG、セルフホスト
記事のタイトル: GitHub - Bike4Mind/bike4mind: オープンコア AI ワークベンチ — あらゆるモデルにわたるノートブック、エージェント、RAG、音声、画像: OpenAI、Anthropic、Google、xAI、または Ollama/vLLM 経由のローカル。 BSL 1.1、2 年周期で Apache-2.0 に自動変換。 AI が動作していないときでも、あなたの AI は動作し続けます。 · GitHub
説明: オープンコア AI ワークベンチ — あらゆるモデルにわたるノートブック、エージェント、RAG、音声、画像: OpenAI、Anthropic、Google、xAI、または Ollama/vLLM 経由のローカル。 BSL 1.1、2 年周期で Apache-2.0 に自動変換。 AI が動作していないときでも、あなたの AI は動作し続けます。 - バイク4マインド/バイク4マインド
HN テキスト: こんにちは HN、私は Erik Bethke です。30 年以上オンライン ゲーム (Starfleet Command、GoPets、FarmVille、Mafia Wars) を作り続けているゲーム開発者です。過去 3 年間、私は Bike4Mind を構築してきましたが、今ではビッグボーイ パンツを履いて「市場に行く」必要があります。ああ、もう少しエージェントを起動して、さらにいくつかのバグを修正し、さらに UX を磨き上げたら、ずっと快適になるでしょう。あるいは、Q、R、または S を実行する新機能はどうでしょうか? Bike4Mindとは何ですか?
私は AI 拡張ゲーム開発ツールの本拠地として Bike4Mind を立ち上げましたが、時間の経過とともに機能を追加し続けてきました。これは、ノートブック、ファイル、データレイクと RAG、エージェント (さまざまな形式)、React または Python での共有可能なアーティファクトの生成など、すべての機能を備えた完全なコグニティブ ワークベンチです。 CLI (B4M CLI) に接続するための安全な Web ソケットがあるため、実際にバッテリーが付属しており、OpenAI、Anthropic、Gemini、xAI、Bedrock などのあらゆるモデルを実行できます。会話中にモデルを切り替え、ノートブックをフォークし、同じプロンプトで 4 つの異なるモデルを実行し、大量の画像生成をサポートし、OpenAI と Anthropic の歴史をインポートし、もちろんいつでも自分のものを使って出力します。

私。つまり、すべてのモデルとすべての機能が揃っています。確かに、でも、だから何？私たちはコードとしてのインフラストラクチャと SSTv4 の活用を強く信じており、Bike4Mind は独自のハードウェア上で 100% 自己ホストできます。 Qwen をローカルで実行し、すべての GUI と CLI を使用して、インターネットから完全にオフラインにすることができます。私が見つけた中で最も完全な完全主権型ワークベンチです。もしあなたがより完全なものを知っているのであれば、ぜひそれを見てみたいと思っています。さて、すべてのモデル、機能、そして私がセルフホストできるということはクールに聞こえますが、Fable などで自分で雰囲気を盛り上げてみませんか?過去 3 年間、エンタープライズ クライアントにこれを使用してきました。そのため、エンタープライズ SaaS のすべての機能も利用できます。バッテリーも含まれています。マルチテナント、RBAC、MFA、スロットリング、モデレーション、分析、プレミアム サブスクリプション、クレジット、スリムな CRM さえも組み込まれています。 OAuth、github、google docs、notion など... カップケーキを舐めることと、MVP を起動して実行して機能させることは別のことですが、企業向けにデプロイできるものを導入することはまったく別のことです。何？いったい何をしているのですか？ 48 時間前 (オースティン 7 月 5 日正午)、私たちは構築したすべてのものを取り出し、私がレイアウトできる最も進歩的なライセンスでオープンしました。つまり、ソースは一方向ラチェット付きの BSL 1.1 ライセンスの直下にあり、24 か月後に各リリースが完全な Apache-2.0 に変換されます。これは、Bike4Mind がリリースごと、毎月ごとにさらにオープンになる、取り返しのつかない一方向のラチェットであることをご理解ください。これは、「おい、私たちは咳き込むまでオープンソースだ - ジャズの手は閉まっている - すべての PR に感謝する!」というものではありません。誤解のないように言っておきますが、私はソースが利用可能であることがオープンソースと同じではないことを知っています。ただし、実際のライセンスとマーケティング サイトを確認してください。私たちは、誰でもフォークして新しい製品を作成することを奨励し、許可しています。

それらを販売するか、すべてを取得して独自の AWS アカウントまたは独自のハードウェアでセルフホストします。また、それをフォークして、独自のクライアント向けにフルスタック AI ソリューションをカスタマイズすることもできます。実際のところ、私たちが 24 か月強にわたって差し止めている唯一のことは、Bike4Mind のバニラ コピーを立ち上げて B4M-as-a-service と競合したり、名前を変更して同じことをしたりすることを許可しないことです。

記事本文:
GitHub - Bike4Mind/bike4mind: オープンコア AI ワークベンチ - あらゆるモデル (OpenAI、Anthropic、Google、xAI、または Ollama/vLLM 経由のローカル) にわたるノートブック、エージェント、RAG、音声、画像。 BSL 1.1、2 年周期で Apache-2.0 に自動変換。 AI が動作していないときでも、あなたの AI は動作し続けます。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロード先

セッションをリフレッシュしてください。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
バイク4マインド
/
バイク4マインド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
89 コミット 89 コミット .changeset .changeset .github .github .husky .husky .vscode .vscode apps/ client apps/ client b4m-core b4m-core ブループリント ブループリント docs-site/ docs docs-site/ docs infra infra パッケージ パッケージ スクリプト スクリプト selfhost/ ws-gateway selfhost/ ws-gateway .cursorrules .cursorrules .dockerignore .dockerignore .env.example .env.example .env.secrets.example .env.secrets.example .env.selfhost.example .env.selfhost.example .envrc .envrc .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .gitleaks.toml.simple .gitleaks.toml.simple .gitleaksignore .gitleaksignore .markdownlint.json .markdownlint.json .mcp.json .mcp.json .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.cjs .prettierrc.cjs .semgrep.yml .semgrep.yml BIKE4MIND_CLI.md BIKE4MIND_CLI.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md SELF_HOST.md SELF_HOST.md bootstrap.sh bootstrap.sh commitlint.config.js commitlint.config.js compose.ollama-gpu.yaml compose.ollama-gpu.yaml compose.selfhost.yaml compose.selfhost.yaml compose.yaml compose.yaml dev dev dev-fast dev-fast elasticmq.conf elasticmq.conf eslint.config.mjs eslint.config.mjs fix-gitleaks.sh fix-gitleaks.sh flake.lock flake.lock flake.nix flake.nix for-env.example for-env.ex

十分な install-hooks.sh install-hooks.sh mongo mongo next-env.d.ts next-env.d.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml premium-overlay.lock.json premium-overlay.lock.json review-secrets.sh review-secrets.shサンプル-ノートブック-エクスポート.json サンプル-ノートブック-エクスポート.json sst-dev-fast sst-dev-fast sst-env.d.ts sst-env.d.ts sst.config.ts sst.config.ts start-db start-db tsconfig.json tsconfig.jsonturbo.jsonturbo.json vitest.config.ts vitest.config.ts vitest.shared.ts vitest。
[切り捨てられた]
AI 時代の心の自転車。
オープンコア AI ナレッジ プラットフォーム - ノートブック、自律エージェント、RAG ナレッジ エンジンが連携して認知タスクを強化するマルチモデル ワークスペース。
Bike4Mind は、AI テクノロジーのスーパーセットを 1 つのワークスペースで結び付ける、AI を活用したナレッジ プラットフォームです。独自のモデルを持ち込み、データを保管し、面倒な作業はエージェントに任せましょう。
ノートブック — コア ワークスペース: チャット、ドキュメント、コンテキストをあらゆるモデルにわたって 1 か所にまとめます。
設計によるマルチ LLM — Anthropic、OpenAI、Google Gemini、xAI、Ollama (ローカル) などに加え、画像モデル (FLUX、DALL·E 3、安定拡散) を自由に交換できます。
Quest Master — 自律的な複数ステップのタスク計画: テキストと画像の生成、ビジョンレビュー、Web 検索、数学、コード、および人間参加型のステップが並行して実行されます。
AI エージェント — ツールや知識に基づいて複雑なタスクを実行する自律的な ReAct スタイルのエージェント。
ナレッジ エンジン — ドキュメントに対する RAG: スマート チャンキング、ベクトル検索、コレクション、タグ付けにより、エージェントはあなたが知っていることを推論できます。
アーティファクト — エージェントによって作成された再利用可能なスニペット、ドキュメント、ビジュアライゼーション。公開および共有レイヤーが組み込まれています。
app.bike4mind.com でホストされているサービスを試すか、SE

独自のインフラストラクチャ上でオープン コアを lf-host します (下記)。
Bike4Mind はオープンコアです。エンジンはパブリックであり、自己ホスト可能です。マルチテナントのホスト型サービスは私たちのビジネスです。
(このリポジトリ、BUSL-1.1) を開きます: エージェント エンジン、LLM アダプター、CLI、データ モデル、およびセルフホスト パス — @bike4mind/* パッケージとして公開されます。
商用 / クローズド : マルチテナントのホスト型サービス (請求、資格、ホスト型インフラストラクチャ)、およびオーバーウォッチなどのプレミアム オーバーレイの運用。
正確な開閉境界については、CONTRIBUTING.md を参照してください。
ホスト型 (最速): app.bike4mind.com にサインインします。インストールするものはありません。
セルフホスト: Docker を使用して独自のハードウェアでオープン コアを実行します。AWS アカウントやクラウド プロバイダーは必要ありません。 「セルフホストのクイックスタート」を参照してください。要するに:
# 1. env テンプレートをコピーし、JWT_SECRET / SESSION_SECRET / を生成します
# SECRET_ENCRYPTION_KEY (SELF_HOST.md を参照) およびモデル キーを追加します
cp .env.selfhost.example .env.selfhost
# 2. スタックを起動する (アプリ + MongoDB + オブジェクト ストレージ + キュー + メール キャッチャー)
docker compose -f compose.selfhost.yaml --env-file .env.selfhost up -d
Bike4Mind は http://localhost:3000 で実行されます。サインイン コードの電子メールは、バンドルされた Mailpit ( http://localhost:8025 ) に送信されます。セルフホスト イメージは ghcr.io/bike4mind/bike4mind-selfhost に公開されます。
Bike4Mind は、pnpm + Turborepo モノレポ (Node 24) です。
pnpm i -r # インストール
pnpmturbo:core:build # @bike4mind/* コア パッケージをビルドします
pnpm Turbo:typecheck # 型チェック
pnpmturbo:test # テストを実行する
プロジェクトのレイアウト: アプリ/クライアント (Next.js SPA + ページ API バックエンド)、パッケージ/cli (対話型 CLI + ReAct エージェント)、b4m-core/* (@bike4mind/* エンジン パッケージ)、パッケージ/データベース (Mongoose モデル + 移行)。リアルタイム WebSocket ファンアウトは、別個の @bike4mind/subscriber-fanout イメージとして出荷されます。完全なガイドについては、CONTRIBUTING.md を参照してください。
コン

貢献は歓迎です — フォーク → トピックブランチ → PR → スカッシュマージ。すべての寄稿者は、最初の PR で軽量の CLA に署名します (著作権はお客様が保持します)。まずは CONTRIBUTING.md と行動規範から始めてください。質問とセルフホスティングのヘルプは ディスカッション にあります。
脆弱性は非公開で報告してください。SECURITY.md を参照してください。セキュリティ上の問題については公開問題を開かないでください。
Bike4Mind は、ビジネス ソース ライセンス 1.1 に基づいてライセンスされており、広範な追加使用許可が付与されています。
✅ コードを読み取り、変更、再配布し、実稼働環境で使用することができます。これには、組織内で使用するためにコードを自己ホストすることや、その上に独自の製品を構築または商品化することも含まれます。
❌ ソフトウェアを競合するホスト/管理サービス (ライセンスで定義されている「Bike4Mind サービス」) として第三者に提供することはできません。
🕓 リリースされた各バージョンは、公開リリースから 2 年後に Apache-2.0 に変換されます。このライセンスが厳格化されることはありません。
代替ライセンスについては、licensing@bike4mind.com にお問い合わせください。
オープンコア AI ワークベンチ - あらゆるモデルにわたるノートブック、エージェント、RAG、音声、画像: OpenAI、Anthropic、Google、xAI、または Ollama/vLLM 経由のローカル。 BSL 1.1、2 年周期で Apache-2.0 に自動変換。 AI が動作していないときでも、あなたの AI は動作し続けます。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
11
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The open-core AI workbench — notebooks, agents, RAG, voice, and images across any model: OpenAI, Anthropic, Google, xAI, or local via Ollama/vLLM. BSL 1.1, auto-converting to Apache-2.0 on a two-year clock. Your AI keeps running when theirs doesn't. - Bike4Mind/bike4mind

Hi HN, I am Erik Bethke, game dev that has been making online games for 30+ years (Starfleet Command, GoPets, FarmVille, Mafia Wars). For the last 3 years I have been building Bike4Mind and now I have to put my big boy pants on and 'go to market'. Ugh, it would be so much more cozy to fire up a few more agents and fix a few more bugs, polish some more UX - or hey how about a new feature that does Q, R or S? What is Bike4Mind?
I started Bike4Mind to be the home of a bunch of AI augmented game development tools, and over time have just kept adding features. It is a full cognitive workbench we have all of the features - notebooks, files, data lakes & RAG, agents (different forms), generate share-able artifacts in react or python. We have secure websockets to connect to our CLI - B4M CLI, so it is truly batteries included - and runs any model - OpenAI, Anthropic, Gemini, xAI, Bedrock. Switch models during a conversation, fork a notebook, run 4 different models on the same prompt, loads of image generation support, and import OpenAI and Anthropic history and of course egress with your own at any time. So, all of the models and all the features. Sure, but so what? We are huge believers in infrastructure as code and leveraged SSTv4, and Bike4Mind is able to 100% self-host on your own hardware. You can run Qwen locally and have all of the GUI and CLI and be totally offline from the internet. The most complete fully-sovereign workbench I've found - and if you know a more complete one I'd genuinely like to see it. Okay, so all of the models, features and I can self host - sounds cool, but why don't I vibe that up myself in Fable or such right? We have used this with our enterprise clients over the last 3 years and so you also get all of the Enterprise SaaS stuff - batteries included - multi-tenant, RBAC, MFA, throttling, moderation, analytics, premium subscriptions, credits, heck even a slim CRM is built in! OAuth, github, google docs, notion, etc etc… it is one thing to lick the cupcake and get that MVP up and running and working for you - it is a totally different beast putting something that you can deploy for an enterprise. What? What are you doing exactly? 48 hours ago (July 5th noon Austin) we took everything we built and we made it open with the most progressive license I could lay out: Source immediately under a BSL 1.1 license with a one-way ratchet where after 24 months each release converts to full Apache-2.0. Please appreciate that this is an irrevocable, one-way ratchet where Bike4Mind just gets more open with each release and each month. This is not a "Hey we are open source until well cough reasons - jazz hands we are closed - thanks for all of the PRs!" To be clear, I know that source-available is not the same as open-source. But check out our actual license and the marketing site - we encourage and permit anyone to fork and create new products and sell them, or take the whole thing and self-host on their own AWS account or their own hardware. You can also fork it and customize a full stack AI solution for your own clients. Really the only thing we are holding back - and just over 24 months - is not allowing people to stand up a vanilla copy of Bike4Mind and compete with B4M-as-a-service or change the name and do the same.

GitHub - Bike4Mind/bike4mind: The open-core AI workbench — notebooks, agents, RAG, voice, and images across any model: OpenAI, Anthropic, Google, xAI, or local via Ollama/vLLM. BSL 1.1, auto-converting to Apache-2.0 on a two-year clock. Your AI keeps running when theirs doesn't. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Bike4Mind
/
bike4mind
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
89 Commits 89 Commits .changeset .changeset .github .github .husky .husky .vscode .vscode apps/ client apps/ client b4m-core b4m-core blueprints blueprints docs-site/ docs docs-site/ docs infra infra packages packages scripts scripts selfhost/ ws-gateway selfhost/ ws-gateway .cursorrules .cursorrules .dockerignore .dockerignore .env.example .env.example .env.secrets.example .env.secrets.example .env.selfhost.example .env.selfhost.example .envrc .envrc .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .gitleaks.toml.simple .gitleaks.toml.simple .gitleaksignore .gitleaksignore .markdownlint.json .markdownlint.json .mcp.json .mcp.json .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.cjs .prettierrc.cjs .semgrep.yml .semgrep.yml BIKE4MIND_CLI.md BIKE4MIND_CLI.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SELF_HOST.md SELF_HOST.md bootstrap.sh bootstrap.sh commitlint.config.js commitlint.config.js compose.ollama-gpu.yaml compose.ollama-gpu.yaml compose.selfhost.yaml compose.selfhost.yaml compose.yaml compose.yaml dev dev dev-fast dev-fast elasticmq.conf elasticmq.conf eslint.config.mjs eslint.config.mjs fix-gitleaks.sh fix-gitleaks.sh flake.lock flake.lock flake.nix flake.nix for-env.example for-env.example install-hooks.sh install-hooks.sh mongo mongo next-env.d.ts next-env.d.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml premium-overlay.lock.json premium-overlay.lock.json review-secrets.sh review-secrets.sh sample-notebook-export.json sample-notebook-export.json sst-dev-fast sst-dev-fast sst-env.d.ts sst-env.d.ts sst.config.ts sst.config.ts start-db start-db tsconfig.json tsconfig.json turbo.json turbo.json vitest.config.ts vitest.config.ts vitest.shared.ts vitest.
[truncated]
Your bicycle for the mind in the age of AI.
An open-core AI knowledge platform — a multi-model workspace where notebooks, autonomous agents, and a RAG knowledge engine work together to augment any cognitive task.
Bike4Mind is an AI-powered knowledge platform that wires a superset of AI technologies together behind one workspace. Bring your own models, keep your data, and let agents do the heavy lifting.
Notebooks — the core workspace: chat, documents, and context in one place, across any model.
Multi-LLM by design — swap freely between Anthropic, OpenAI, Google Gemini, xAI, Ollama (local), and more, plus image models (FLUX, DALL·E 3, Stable Diffusion).
Quest Master — autonomous, multi-step task planning: text and image generation, vision review, web search, math, code, and human-in-the-loop steps, run in parallel.
AI Agents — autonomous ReAct-style agents that carry out complex tasks against your tools and knowledge.
Knowledge Engine — RAG over your documents: smart chunking, vector search, collections, and tagging so agents can reason about what you know.
Artifacts — reusable snippets, documents, and visualizations produced by agents, with a built-in publish-and-share layer.
Try the hosted service at app.bike4mind.com , or self-host the open core on your own infrastructure (below).
Bike4Mind is open core . The engine is public and self-hostable; the multi-tenant hosted service is our business.
Open (this repo, BUSL-1.1): the agent engine, LLM adapters, CLI, data models, and the self-host path — published as @bike4mind/* packages.
Commercial / closed : operating the multi-tenant hosted service — billing, entitlements, hosted infrastructure — and premium overlays such as Overwatch.
See CONTRIBUTING.md for the exact open/closed boundary.
Hosted (fastest): sign in at app.bike4mind.com — nothing to install.
Self-host: run the open core on your own hardware with Docker — no AWS account or cloud provider required. See the Self-Host Quickstart . In short:
# 1. Copy the env template, then generate JWT_SECRET / SESSION_SECRET /
# SECRET_ENCRYPTION_KEY (see SELF_HOST.md) and add your model keys
cp .env.selfhost.example .env.selfhost
# 2. Bring up the stack (app + MongoDB + object storage + queues + mail catcher)
docker compose -f compose.selfhost.yaml --env-file .env.selfhost up -d
Bike4Mind then runs at http://localhost:3000 ; sign-in code emails land in the bundled Mailpit at http://localhost:8025 . The self-host image is published to ghcr.io/bike4mind/bike4mind-selfhost .
Bike4Mind is a pnpm + Turborepo monorepo (Node 24).
pnpm i -r # install
pnpm turbo:core:build # build the @bike4mind/* core packages
pnpm turbo:typecheck # type check
pnpm turbo:test # run tests
Project layout: apps/client (Next.js SPA + pages API backend), packages/cli (interactive CLI + ReAct agent), b4m-core/* (the @bike4mind/* engine packages), packages/database (Mongoose models + migrations). Realtime WebSocket fanout ships as a separate @bike4mind/subscriber-fanout image. See CONTRIBUTING.md for the full guide.
Contributions are welcome — fork → topic branch → PR → squash-merge. All contributors sign a lightweight CLA on their first PR (you keep your copyright). Start with CONTRIBUTING.md and our Code of Conduct . Questions and self-hosting help go in Discussions .
Please report vulnerabilities privately — see SECURITY.md . Do not open public issues for security problems.
Bike4Mind is licensed under the Business Source License 1.1 with a broad Additional Use Grant:
✅ You may read, modify, redistribute, and make production use of the code — including self-hosting it for your organization's own internal use, and building or commercializing your own products on top of it.
❌ You may not offer the software to third parties as a competing hosted/managed service (a "Bike4Mind Service", as defined in the LICENSE).
🕓 Each released version converts to Apache-2.0 two years after its public release. This license will never be tightened.
For alternative licensing, contact licensing@bike4mind.com .
The open-core AI workbench — notebooks, agents, RAG, voice, and images across any model: OpenAI, Anthropic, Google, xAI, or local via Ollama/vLLM. BSL 1.1, auto-converting to Apache-2.0 on a two-year clock. Your AI keeps running when theirs doesn't.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
11
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
