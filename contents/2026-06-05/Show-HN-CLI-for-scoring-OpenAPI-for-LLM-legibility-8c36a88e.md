---
source: "https://github.com/jentic/jentic-api-scorecard"
hn_url: "https://news.ycombinator.com/item?id=48409696"
title: "Show HN: CLI for scoring OpenAPI for LLM legibility"
article_title: "GitHub - jentic/jentic-api-scorecard: Jentic API Scorecard · GitHub"
author: "seanblanchfield"
captured_at: "2026-06-05T10:25:46Z"
capture_tool: "hn-digest"
hn_id: 48409696
score: 3
comments: 0
posted_at: "2026-06-05T08:35:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CLI for scoring OpenAPI for LLM legibility

- HN: [48409696](https://news.ycombinator.com/item?id=48409696)
- Source: [github.com](https://github.com/jentic/jentic-api-scorecard)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T08:35:34Z

## Translation

タイトル: HN を表示: LLM の読みやすさのために OpenAPI をスコアリングするための CLI
記事のタイトル: GitHub - jentic/jentic-api-scorecard: Jentic API スコアカード · GitHub
説明: Jentic API スコアカード。 GitHub でアカウントを作成して、jentic/jentic-api-scorecard の開発に貢献してください。
HN テキスト: 私たちは以前、エージェントの対応力について API を評価するためのルーブリックをオープンソース化しました (OpenAPI イニシアチブに深く関与している同僚によって慎重に設計されました)。私たちは、このルーブリック (決定論的評価と LLM ベースの評価の組み合わせ) に従って API をスコアリングするための優れた Web アプリをホストしました。そして、人々が回帰テストと長期的な追跡のためにそれを CI/CD パイプラインにハッキングしていることを知りました。そこで、それを簡単にするために、適切な無料枠を備えた優れた CLI をリリースしました。

記事本文:
GitHub - jentic/jentic-api-scorecard: Jentic API スコアカード · GitHub
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
ジェンティック
/
jentic-api-スコアカード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Co

de 「その他のアクション」メニューを開く フォルダーとファイル
144 コミット 144 コミット .claude .claude .github .github .husky .husky アセット アセット docker docker docs docs パッケージ パッケージ 仕様 仕様 .commitlintrc.json .commitlintrc.json .gitignore .gitignore .lintstagedrc.json .lintstagedrc.json .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md ライセンス ライセンス通知 NOTICE README.md README.md eslint.config.js eslint.config.js lerna.json lerna.json package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリファイルナビゲーション
検証に合格した OpenAPI ドキュメントが、AI エージェントが使用できるとは限りません。文法は
一つ。意味の明確さ、安全性、発見しやすさは別です。 Jentic API スコアカード
OpenAPI ドキュメントを
Jentic API AI 準備フレームワーク (JAIRF)
6 つの次元にわたって単一のグレードを返すため、どこを改善すべきかが正確にわかります。
各 OpenAPI ドキュメントは 6 つのレンズにわたって評価されます。いずれかのレンズで、対象を絞った小規模な改善が行われます。
人間の開発者と AI エージェントの両方に多大な利益をもたらす傾向があります。
基本的なコンプライアンス (FC) — 構造的妥当性と OpenAPI 自体への準拠。
開発者エクスペリエンスと Jentic 互換性 (DXJ) — ドキュメントの品質と、
OpenAPI ドキュメントはダウンストリーム ツールと連携します。
AI 対応性とエージェント エクスペリエンス (ARAX) — セマンティックな明確さと LLM に必要なコンテキスト
各操作の理由。
エージェント ユーザビリティ (AU) — 予測可能で安全な複数ステップのオーケストレーション。
セキュリティ (SEC) — 宣言された認証スキームと信頼境界。
AI Discoverability (AID) — AI システムが OpenAPI ドキュメントをどれだけ簡単に見つけて解析できるか。
スコアリングは、Docker コンテナ内で 2 つのフェーズでローカルに実行されます。分析ではバッテリーが稼働します
バリデーターとセント

OpenAPI ドキュメントに対する構造チェックを行い、一連の診断と
観察。
スコアリングでは、これらを 6 つの JAIRF 次元にわたって約 35 個のシグナルにマッピングし、それらを次のように集約します。
ディメンションごとのスコアを計算し、それらを単一の加重スコアとグレードにロールアップします。
npm/npx を備えた Node.js 20 LTS 以降 (>= 20.19.0)。 Node.js のダウンロードを参照してください。
Docker がインストールされ、実行されています。 「Docker のインストール」を参照してください。
CLI は、最初の実行時にスコアリング イメージを自動的に取得します。
ghcr.io (イメージをプルするため) と、
スコアリングしている OpenAPI ドキュメント (エンジンがコンテナ内からドキュメントを取得します)。
npm install -g @jentic/api-scorecard-cli
これにより、CLI がグローバルにインストールされます。スコアリング エンジン (Docker イメージ) は自動的にダウンロードされます
初めてスコアを実行するときは、通常の接続で 1 ～ 2 分かかります。
ローカル ファイルまたは非 OAK URL の場合は、JENTIC_API_KEY も必要です。「」を参照してください。
匿名アクセスとキー付きアクセス。
jentic-api-scorecard --version
ゼロインストールをご希望ですか?グローバル インストールをスキップして npx を使用できます。すべての例は次のとおりです。
この README は、 jentic-api-scorecard の代わりに npx @jentic/api-scorecard-cli で動作します。
npx @jentic/api-scorecard-cli@<version> (例: @1.0.0 ) を使用して特定のリリースに固定します。
固定されていないフォームは、各呼び出しで最新の dist タグが指すものに解決されます。
一方、npm install -g では、明示的に更新するまで、インストールされているバージョンに固定されます。
Jentic Public API (OAK) の OpenAPI ドキュメント
キーなし、上限なしのスコア — サインアップなし、設定なし:
npx @jentic/api-scorecard-cli@最新スコア \
https://raw.githubusercontent.com/jentic/jentic-public-apis/refs/heads/main/apis/openapi/swagger-api/petstore/1.0.27/openapi.json
OAK 外部の URL またはローカル ファイルの場合は、API キーを設定します。
JENTIC_API_KEY= < キー > npx @jentic/api-scorecard-cli@latest

スコア \
https://petstore3.swagger.io/api/v3/openapi.json
JENTIC_API_KEY= < キー > npx @jentic/api-scorecard-cli@最新スコア ./openapi.yaml
重要
無料キーには毎月 100 スコアが付属しています (各暦月の初めにリセットされます)。サインアップとクォータの詳細については、「匿名アクセスとキー付きアクセス」を参照してください。
それでおしまい。 CLI は、最初の実行時にスコアリング エンジンを自動的に取得します。
--detail フラグを使用すると、ズームインできます。
# 見出しのスコアと成績のみ
npx @jentic/api-scorecard-cli@latest スコア --詳細概要 ./openapi.yaml
# ディメンションごとの内訳 (デフォルト)
npx @jentic/api-scorecard-cli@latest スコア --詳細寸法 ./openapi.yaml
# 各次元内の個々の信号
npx @jentic/api-scorecard-cli@latest スコア --detail シグナル ./openapi.yaml
# 重大度ごとに上位 5 件の結果を含む完全な診断
npx @jentic/api-scorecard-cli@latest スコア --詳細診断 ./openapi.yaml
機械可読出力
--format json を追加して、エンジンの逐語的な JSON を標準出力に出力します (任意のフィルタリングを適用)
-- 詳細レベルを選択します)。 Pretty は無条件のデフォルトのままです。 --json のフォーマット
CI ゲート、アーカイブ、
または LLM 支援レビュー。
# CI のヘッドラインスコアのゲート
npx @jentic/api-scorecard-cli@最新スコア ./openapi.yaml --format json | jq .summary.score
# 完全な証拠バンドルをファイルにキャプチャします
npx @jentic/api-scorecard-cli@最新スコア ./openapi.yaml \
--format json --detail Diagnostics --output report.json
--output <file> ( -o ) は、標準出力ではなくパスにレポートを書き込みます。スピナーは標準エラー出力に残ります。
--quiet ( -q ) は、対話型端末でも stderr スピナーを抑制します (スピナーはすでに
stderr が TTY ではない場合は自動抑制されます)。エンジンの警告は引き続き標準エラー出力を通過します。
--with-llm を追加して LLM ベースのシグナルのロックを解除します - 深い

API かどうかについてのセマンティック推論
説明はエージェントにとって実行可能であり、エラー応答が自律回復をサポートしているかどうか、
もっと。 LLM プロバイダーが必要です: クラウド (OpenAI / Anthropic / Gemini / AWS Bedrock) またはローカル
OpenAI 互換エンドポイント (Ollama、LM Studio、vLLM など)。
エクスポート OPENAI_API_KEY=sk-...
エクスポート LLM_PROVIDER=OPENAI
エクスポート LIGHT_LLM_PROVIDER=OPENAI
エクスポート LLM_LIGHT_MODEL=gpt-4o-mini
JENTIC_API_KEY= < キー > npx @jentic/api-scorecard-cli@最新スコア ./openapi.yaml --with-llm
トークンコストが低い - エンジンは軽量モデルを使用しています (Claude Haiku、GPT-4o-mini など)。
操作は小さなバッチで処理され、仕様のサイズに関係なく 7 バッチに制限されます。ローカル
モデル (Ollama) では、通話ごとに費用はかかりません。
LLM シグナルガイドを参照してください。
すべてのプロバイダー レシピ (ローカル Ollama を含む)、完全な環境変数参照、および
トラブルシューティング。
Jentic Public API (OAK) の OpenAPI ドキュメント
キーなしでスコアを取得し、無料枠を維持します。これらの URL はキーの検証を完全にバイパスします。
それ以外のすべて (ローカル ファイル、OAK 外部の URL) については、jentic.com/signup でキーを取得します。サインインしたら、[Score] → [CLI & Keys] をクリックしてキーを発行します。次に、次のように設定します。
import JENTIC_API_KEY= < キー >
実際のキーは、コンテナーによって api.jentic.com に対してライブで検証されます。同じ通話が 2 倍になる
キーごとの使用量/レート制限アカウンティングのヒットとして。各無料キーは毎月 100 スコアを取得します。
各暦月の初めにリセットされます。そのクォータが使い果たされると、CLI は次のように終了します。
コード 7 を実行し、プランをアップグレードするためのリンクとともに Retry-After の値を出力します。
jentic-api-scorecard [-V | --バージョン] [-h | --ヘルプ]
jentic-api-scorecard <コマンド> [オプション]
コマンド
コマンド
説明
スコア<入力>
URL またはローカル ファイル パスによって OpenAPI ドキュメントをスコア付けします。
スコア
OpenAPI ドキュメントにスコアを付ける b

y URL またはローカル ファイル パス。
jentic-api-scorecard スコア <入力> [オプション]
引数
名前
説明
<入力>
https:// OpenAPI ドキュメントへの URL またはローカル ファイル パス。必須。
オプション
旗
デフォルト
選択肢
説明
--with-llm
オフ
—
LLM に基づく分析を有効にします。 LLM プロバイダーが必要です (LLM 分析を参照)。
--バンドル
オフ
—
URL 入力に対して CLI 側のバンドルを強制します。CLI はホスト上の URL をフェッチし、Redocly でバンドルし、stdin 経由でコンテナにパイプします。ホストのみが到達できる URL (内部ネットワーク、VPN ゲート仕様、認証が必要な URL) に使用します。 JENTIC_API_KEY が必要です。ローカルファイルに対しては何もしません。
-d、--detail <レベル>
寸法
概要、寸法、信号、診断
ペイロードの深さ (「制御出力の深さ」を参照)。
-f、--format <fmt>
かなり
かわいい、ジェソン
出力エンコーディング (「 機械可読出力 」を参照)。
-o, --output <ファイル>
標準出力
—
フォーマットされたレポートを <file> に書き込みます。スピナーは標準エラー出力に残ります。
-q、--静か
オフ
—
TTY に関係なく、stderr スピナーを抑制します。
-h、--ヘルプ
—
—
スコアの使用法を表示します。
環境
変数
いつ
目的
JENTIC_API_KEY
OAK 外部の URL とローカル ファイル
jentic.com/signup で発行された実際のキー。 api.jentic.com に対してライブで検証されています (匿名アクセスとキー付きアクセスを参照)。無料割り当て: 暦月あたり 100 スコア。
LLM プロバイダー + ルーティング変数
--with-llm あり
CLI は認証情報 ( OPENAI_API_KEY 、 ANTHROPIC_API_KEY 、 GEMINI_API_KEY 、 AWS キー) とルーティング ( LLM_PROVIDER 、 LIGHT_LLM_PROVIDER 、 LLM_MODEL 、 LLM_LIGHT_MODEL 、 *_API_URL 、 LLM_MAX_TOKENS ) を自動検出し、コンテナに転送します。ループバック URL は、ホスト側の Ollama に到達できるように書き換えられます。完全なリファレンス: LLM シグナル ガイド。
終了コード
コード
意味
0
採点が完了しました (スコア自体には関係ありません)。
1
一般的なエラー (不正な入力、バンドルの失敗、予期しないコンテナの失敗)。

2
認証: JENTIC_API_KEY が Jentic バックエンドが認識しない値に設定されているか、キーが設定されていないローカル ファイル/stdin 入力が使用されました。
3
匿名ゲートが拒否されました: URL が OAK ホワイトリスト外であり、キーが設定されていません。
4
Docker がインストールされていません o
[切り捨てられた]
jentic.com/scorecard は、Web UI で同じスコアリングを提供します。
URL を貼り付けるかファイルをドロップするだけで、Docker や Node は必要ありません。
何が実行されているかを正確に知る必要があるチームの場合は、何が実行されているかを正確に確認してください。
出荷され、Jentic へのランタイム依存関係なしで実行されます。
スコアリング スタック内のすべてのコンポーネント - ランナー、CLI、リリース パイプライン、
エンジン — Apache 2.0 ライセンスがあり、ソースを読み取ることができます。独自の BLOB はありません。
クローズドソースのシムはありません。仕様を評価するコードを読んでください
それを採用する前に。任意の行を監査し、ライセンス条項に基づいて再配布します。
必要に応じてフォークします。
規制された環境向けに署名済み
すべての npm tarball とすべての GHCR イメージは Sigstore によって署名されています
SLSA の来歴と SPDX SBOM を備えています。署名は OIDC 駆動の内部で行われます
GitHub Actions ワークフローには長期間有効な公開シークレットがありません。
NPM_TOKEN 、PAT なし、リリース チェーンに人間のキーホルダーはありません。 1つのコマンド
アーティファクトをインストールする前にエンドツーエンドで検証します。
npmパッケージのサプライチェーン → —
npm 来歴、SPDX SBOM、信頼できる

[切り捨てられた]

## Original Extract

Jentic API Scorecard. Contribute to jentic/jentic-api-scorecard development by creating an account on GitHub.

We previously open-sourced a rubric for assessing APIs for agent-readiness (carefully designed by some colleagues who are deeply involved with the OpenAPI initiative). We hosted a nice web app for scoring an API according to this rubric (a mix of deterministic and LLM-based assessment), and learned people were hacking it into their CI/CD pipelines for regression testing and longitudinal tracking. So we just released a nice CLI to make that easier, with a decent free tier.

GitHub - jentic/jentic-api-scorecard: Jentic API Scorecard · GitHub
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
jentic
/
jentic-api-scorecard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
144 Commits 144 Commits .claude .claude .github .github .husky .husky assets assets docker docker docs docs packages packages specs specs .commitlintrc.json .commitlintrc.json .gitignore .gitignore .lintstagedrc.json .lintstagedrc.json .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md LICENSE LICENSE NOTICE NOTICE README.md README.md eslint.config.js eslint.config.js lerna.json lerna.json package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json View all files Repository files navigation
An OpenAPI document that passes validation isn't necessarily one an AI agent can use. Grammar is
one thing; semantic clarity, safety, and discoverability are another. The Jentic API Scorecard
scores your OpenAPI document against the
Jentic API AI Readiness Framework (JAIRF)
across six dimensions and returns a single grade — so you know exactly where to improve.
Each OpenAPI document is evaluated across six lenses — small, targeted improvements in any of them
tend to produce outsized gains for both human developers and AI agents:
Foundational Compliance (FC) — structural validity and conformance to OpenAPI itself.
Developer Experience & Jentic Compatibility (DXJ) — documentation quality and how well the
OpenAPI document plays with downstream tooling.
AI-Readiness & Agent Experience (ARAX) — semantic clarity and the context an LLM needs to
reason about each operation.
Agent Usability (AU) — predictable, safe multi-step orchestration.
Security (SEC) — declared auth schemes and trust boundaries.
AI Discoverability (AID) — how easily an AI system can find and parse the OpenAPI document.
Scoring runs locally inside a Docker container in two phases. Analysis runs a battery of
validators and structural checks against the OpenAPI document to produce a set of diagnostics and
observations.
Scoring maps those into ~35 signals across the six JAIRF dimensions, aggregates them into
per-dimension scores, and rolls those up into a single weighted score and grade.
Node.js 20 LTS or newer ( >= 20.19.0 ) with npm/npx. See Node.js downloads .
Docker installed and running. See Docker installation .
The CLI pulls the scoring image automatically on first run.
Network access to ghcr.io (to pull the image) and to whatever URL hosts the
OpenAPI document you're scoring (the engine fetches it from inside the container).
npm install -g @jentic/api-scorecard-cli
This installs the CLI globally. The scoring engine (Docker image) is downloaded automatically
the first time you run score — allow a minute or two on a typical connection.
For local files or non-OAK URLs you'll also need a JENTIC_API_KEY — see
Anonymous vs keyed access .
jentic-api-scorecard --version
Prefer zero-install? You can skip the global install and use npx — every example in
this README works with npx @jentic/api-scorecard-cli in place of jentic-api-scorecard .
Pin to a specific release with npx @jentic/api-scorecard-cli@<version> (e.g. @1.0.0 );
the unpinned form resolves to whatever the latest dist-tag points at on each invocation,
while npm install -g pins you to the installed version until you explicitly update.
OpenAPI documents from Jentic Public APIs (OAK)
score without any key, uncapped — no signup, no config:
npx @jentic/api-scorecard-cli@latest score \
https://raw.githubusercontent.com/jentic/jentic-public-apis/refs/heads/main/apis/openapi/swagger-api/petstore/1.0.27/openapi.json
For URLs outside OAK or local files, set the API key:
JENTIC_API_KEY= < your-key > npx @jentic/api-scorecard-cli@latest score \
https://petstore3.swagger.io/api/v3/openapi.json
JENTIC_API_KEY= < your-key > npx @jentic/api-scorecard-cli@latest score ./openapi.yaml
Important
Free keys come with 100 scorings per month (resets at the start of each calendar month). See Anonymous vs keyed access for signup and quota details.
That's it. The CLI pulls the scoring engine automatically on first run.
The --detail flag lets you zoom in:
# Just the headline score and grade
npx @jentic/api-scorecard-cli@latest score --detail summary ./openapi.yaml
# Per-dimension breakdown (default)
npx @jentic/api-scorecard-cli@latest score --detail dimensions ./openapi.yaml
# Individual signals within each dimension
npx @jentic/api-scorecard-cli@latest score --detail signals ./openapi.yaml
# Full diagnostics with top 5 findings per severity
npx @jentic/api-scorecard-cli@latest score --detail diagnostics ./openapi.yaml
Machine-readable output
Add --format json to emit engine-verbatim JSON on stdout (filtered by whatever
--detail level you pick). Pretty stays the unconditional default; --format json
is the canonical way to get a stable machine-readable channel for CI gating, archival,
or LLM-assisted review.
# Gate on the headline score in CI
npx @jentic/api-scorecard-cli@latest score ./openapi.yaml --format json | jq .summary.score
# Capture the full evidence bundle to a file
npx @jentic/api-scorecard-cli@latest score ./openapi.yaml \
--format json --detail diagnostics --output report.json
--output <file> ( -o ) writes the report to a path instead of stdout; the spinner stays on stderr.
--quiet ( -q ) suppresses the stderr spinner even in interactive terminals (the spinner already
auto-suppresses when stderr isn't a TTY). Engine warnings still pass through stderr.
Add --with-llm to unlock LLM-backed signals — deeper semantic reasoning about whether your API
descriptions are actionable for agents, whether error responses support autonomous recovery, and
more. Requires an LLM provider: cloud (OpenAI / Anthropic / Gemini / AWS Bedrock) or a local
OpenAI-compatible endpoint (Ollama, LM Studio, vLLM, …).
export OPENAI_API_KEY=sk-...
export LLM_PROVIDER=OPENAI
export LIGHT_LLM_PROVIDER=OPENAI
export LLM_LIGHT_MODEL=gpt-4o-mini
JENTIC_API_KEY= < your-key > npx @jentic/api-scorecard-cli@latest score ./openapi.yaml --with-llm
Token cost is low — the engine uses a lightweight model (e.g. Claude Haiku, GPT-4o-mini),
processes operations in small batches, and caps at 7 batches regardless of spec size. Local
models (Ollama) cost nothing per call.
See LLM Signals guide
for all provider recipes (including local Ollama), the full environment variable reference, and
troubleshooting.
OpenAPI documents from Jentic Public APIs (OAK)
score without any key and stay on the free tier — those URLs bypass key validation entirely.
For everything else (local files, URLs outside OAK), get a key at jentic.com/signup — once signed in, click Score → CLI & Keys to issue your key. Then set it:
export JENTIC_API_KEY= < your-key >
Real keys are validated live by the container against api.jentic.com . The same call doubles
as the per-key usage / rate-limit accounting hit. Each free key gets 100 scorings per month ,
resetting at the start of each calendar month. Once that quota is exhausted the CLI exits with
code 7 and prints the Retry-After value along with a link to upgrade your plan.
jentic-api-scorecard [-V | --version] [-h | --help]
jentic-api-scorecard <command> [options]
Commands
Command
Description
score <input>
Score an OpenAPI document by URL or local file path.
score
Score an OpenAPI document by URL or local file path.
jentic-api-scorecard score <input> [options]
Arguments
Name
Description
<input>
https:// URL or local file path to an OpenAPI document. Required.
Options
Flag
Default
Choices
Description
--with-llm
off
—
Enable LLM-backed analysis. Requires an LLM provider (see LLM analysis ).
--bundle
off
—
Force CLI-side bundling for URL inputs: the CLI fetches the URL on the host, bundles with Redocly, and pipes to the container via stdin. Use for URLs only the host can reach (internal networks, VPN-gated specs, auth-required URLs). Requires JENTIC_API_KEY . No-op for local files.
-d, --detail <level>
dimensions
summary , dimensions , signals , diagnostics
Payload depth (see Control output depth ).
-f, --format <fmt>
pretty
pretty , json
Output encoding (see Machine-readable output ).
-o, --output <file>
stdout
—
Write the formatted report to <file> . The spinner stays on stderr.
-q, --quiet
off
—
Suppress the stderr spinner regardless of TTY.
-h, --help
—
—
Show usage for score .
Environment
Variable
When
Purpose
JENTIC_API_KEY
URLs outside OAK and local files
Real key issued at jentic.com/signup ; validated live against api.jentic.com (see Anonymous vs keyed access ). Free quota: 100 scorings per calendar month.
LLM provider + routing vars
With --with-llm
The CLI auto-detects credentials ( OPENAI_API_KEY , ANTHROPIC_API_KEY , GEMINI_API_KEY , AWS keys) and routing ( LLM_PROVIDER , LIGHT_LLM_PROVIDER , LLM_MODEL , LLM_LIGHT_MODEL , *_API_URL , LLM_MAX_TOKENS ) and forwards them to the container; loopback URLs are rewritten so a host-side Ollama is reachable. Full reference: LLM Signals guide .
Exit codes
Code
Meaning
0
Scoring completed (regardless of the score itself).
1
Generic error (bad input, bundling failure, unexpected container failure).
2
Auth: JENTIC_API_KEY is set to a value the Jentic backend does not recognize, or a local file / stdin input was used without the key set.
3
Anonymous gate rejected: URL outside the OAK allowlist and no key set.
4
Docker not installed o
[truncated]
jentic.com/scorecard offers the same scoring in a web UI —
paste a URL or drop a file, no Docker or Node required.
For teams that need to know exactly what's running, verify exactly what was
shipped, and run without a runtime dependency on Jentic.
Every component in the scoring stack — runner, CLI, release pipeline, and
engine — is Apache 2.0 licensed and source-readable. No proprietary blobs,
no closed-source shims. Read the code that's about to grade your specs
before you adopt it; audit any line, redistribute under the license terms,
fork if you ever need to.
Signed for regulated environments
Every npm tarball and every GHCR image is signed by Sigstore
with SLSA provenance and an SPDX SBOM. Signing happens inside an OIDC-driven
GitHub Actions workflow with no long-lived publishing secrets — there is no
NPM_TOKEN , no PAT, and no human keyholder in the release chain. One command
verifies an artifact end-to-end before you install it:
npm package supply chain → —
npm provenance, SPDX SBOM, trusted

[truncated]
