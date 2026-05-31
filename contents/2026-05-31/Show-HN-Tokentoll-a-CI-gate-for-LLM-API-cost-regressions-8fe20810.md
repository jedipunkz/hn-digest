---
source: "https://github.com/Jwrede/tokentoll"
hn_url: "https://news.ycombinator.com/item?id=48335559"
title: "Show HN: Tokentoll, a CI gate for LLM API cost regressions"
article_title: "GitHub - Jwrede/tokentoll: Catch LLM cost changes in code review. Infracost for LLM spend. · GitHub"
author: "Jwrede"
captured_at: "2026-05-31T00:30:52Z"
capture_tool: "hn-digest"
hn_id: 48335559
score: 2
comments: 0
posted_at: "2026-05-30T12:41:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tokentoll, a CI gate for LLM API cost regressions

- HN: [48335559](https://news.ycombinator.com/item?id=48335559)
- Source: [github.com](https://github.com/Jwrede/tokentoll)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T12:41:53Z

## Translation

タイトル: HN を表示: Tokentoll、LLM API コスト回帰の CI ゲート
記事のタイトル: GitHub - Jwrede/tokentoll: コード レビューで LLM コストの変更をキャッチします。 LLM 支出のインフラコスト。 · GitHub
説明: コードレビューで LLM コストの変更をキャッチします。 LLM 支出のインフラコスト。 - Jwrede/トークントール

記事本文:
GitHub - Jwrede/tokentoll: コードレビューで LLM コストの変更をキャッチします。 LLM 支出のインフラコスト。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ジュレーデ
/
トークントール
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
54 コミット 54 コミット .github/ workflows .github/ workflows デモ デモ ドキュメント ドキュメント src/ tokentoll src/ tokentoll テスト テスト .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示リポジトリ ファイルのナビゲーション
本番前に LLM コストの後退を防ぎます。
tokentoll は LLM コストの CI ゲートです。 LLM API 呼び出しの Python、JavaScript、TypeScript を静的に分析し、管理するポリシーに対してすべてのプル リクエストをスコア付けし、PASS/WARN/FAIL の判定を PR に直接投稿します。オプションで、ポリシーに違反するとワークフローが失敗するため、コスト回帰をマージできません。
Jwrede/tokentoll-demo は、tokentoll コスト ゲートに接続された小さな多言語 LLM アプリ (Python + TypeScript) です。すでに 2 つの PR がこれに対してオープンになっています。
PR #1: Anthropic Haiku 翻訳ヘルパーを追加します。新しい通話サイト。予算内で十分。判定: PASS、ワークフローは緑色。
PR #2: supportbot を gpt-4o に切り替えます。 2 つのポリシー ルールをトリップするモデル スワップ。判定: 不合格、ワークフローは赤です。
各 PR の会話タブを開いて、tokentoll が実際に投稿した評決コメントを確認します。
PR がポリシーに違反すると、tokentoll は評決とブロック結果リストをコメントしてからゼロ以外で終了するため、チェックは失敗します。例:
## トークントールの判定: FAIL
** 阻害所見 (2): **
- ` src/agent.py:42 ` - 通話あたりのコストが 15.0 倍 (しきい値 5 倍)
- 月間デルタ合計 +$812.00 が予算 $250.00 を超えています
> 必要なアクション: 回帰を元に戻すか、` .tokentoll.yml ` のしきい値を上げるか、除外を追加します。
PR がクリーンな場合、ver

dict は PASS で、コメントにはコスト デルタ テーブルのみが表示されます。ポリシーが構成されていない場合、tokentoll は判定のない情報デルタ コメントを投稿します。
.github/workflows/tokentoll.yml を追加します。
名前 : トークントール
に:
プルリクエスト:
パス:
- " **.py "
- " **.ts "
- 「 **.tsx 」
- 「 **.js 」
- 「 **.jsx 」
権限:
内容：読む
プルリクエスト : 書き込み
ジョブ:
コストゲート:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
付き:
フェッチ深度: 0
- 使用: Jwrede/tokentoll@v0.7.0
付き:
ポリシー違反時のフェイル: true
次に、.tokentoll.yml をリポジトリのルートに追加します。
予算 :
max_monthly_delta_usd : 250
max_callsite_monthly_usd : 100
最大相対増加量 : 5.0
ポリシー:
block_unknown_models : true
フェイルオンポリシー違反 : true
今後の PR は評決コメントを受け取ります。しきい値を超える PR はワークフローに失敗します。
SHA 固定インストールと最小限の権限のセットアップについては、 docs/github-action.md を参照してください。完全なポリシー スキーマについては、 docs/policy.md を参照してください。セキュリティ体制については、 docs/security.md を参照してください。
JavaScript / TypeScript (ツリーシッター経由で解析、 .js 、 .jsx 、 .ts 、 .tsx を処理)
.tokentoll.yml のポリシー ブロックは、PR が失敗したときを制御します。
各ルールは独立しています。そのルールを無効にするには、フィールドを未設定のままにしておきます。完全なリファレンスは docs/policy.md にあります。
pip インストールトークントール
# 現在のディレクトリをスキャンして LLM API 呼び出しとそのコストを確認する
トークントールスキャン。
# 最後のコミットのコストへの影響を表示する
トークントール差分 HEAD~1
# 2 つの参照を比較し、ポリシー違反で失敗する
tokentoll diff main..HEAD --fail-on-policy-violation
サブコマンド:
tokentoll scan [PATH...] [--format table|json|markdown] [--calls-per-month N] [--config PATH]
tokentoll diff [REF] [--base REF] [--head REF] [--format table|json|markdown|github-comment]
[--config PATH] [--ポリシー違反時の失敗]
tokentoll 更新 # バンドル価格を更新

LiteLLM からのデータの取得
構成
.tokentoll.yml はリポジトリのルートに存在し、自動検出されます。ポリシーブロックを超えて:
# 動的 (ランタイム解決) モデル名の SDK ごとのデフォルト
デフォルトモデル:
オープンナイ：gpt-4o-mini
人間性 : クロード-俳句-3-20240307
# 通話サイトごとの想定月間通話量 (金額の見積もりに使用)
月あたりの通話数 : 5000
# 動的モデルのコスト推定を完全にスキップします。
# デフォルトは false: 動的呼び出しは SDK ごとのデフォルトに基づいて価格が設定されます。
Skip_dynamic_models : false
# デフォルトでは (tests/、examples/、docs/、cookbook/、benchmarks/、evals/、
# scripts/、notebooks/) が自動的に適用されます。次の方法でオプトアウトします:
use_default_excludes : false
# 追加の除外 (プレフィックスまたはグロブ)
除外する:
- 「 *_test.py 」
- ベンダー/
# パスごとのオーバーライド (最長プレフィックス一致)
オーバーライド:
- パス : src/agents/
デフォルトモデル: gpt-4o
月あたりの通話数 : 10000
- パス : src/azure/
Skip_dynamic_models : true
動的モデルのデフォルトの解決順序:default_models (SDK ごと) >default_model (汎用) > 組み込み SDK デフォルト。
tokentoll は API キーを必要とせず、テレメトリを送信せず、完全に CI 環境内で実行されます。価格データはパッケージに同梱されており、オンデマンドで LiteLLM から更新されます。推奨される権限セット、SHA ピンニング、およびフォーク PR リスクについては、 docs/security.md を参照してください。
tokentoll には MCP (Model Context Protocol) サーバーが同梱されているため、Claude Code と他の MCP ホストは、エージェントの会話内から LLM コード変更のコストへの影響を確認できます。
pip インストールトークントール[mcp]
クロード mcp add --transport stdio tokentoll -- tokentoll-mcp
scan (パス全体のコストを推定) と diff (2 つの参照を比較) の 2 つのツールが公開されています。どちらも JSON を返します。
ソースコード (.py、.ts、.tsx、.js、.jsx)
|
v
+----------------+ +------------------+
| AST スキャナ |-->| SDK 検出器 |
| ast (パイス

上) + | | OpenAI、人為的|
|ツリーシッター| Google、LiteLLM、 |
| (JS/TS) | |ラングチェーン、知浦、|
+----------------+ | Vercel AI SDK |
+-----------------+
|
v
+-----------------+
|価格設定エンジン |
| 2200以上のモデル |
+-----------------+
|
v
+-----------------+
|デフエンジン |
| (古いものと新しいもの) |
+-----------------+
|
v
+-----------------+
|ポリシー評価者 |
|合格/警告/不合格 |
+-----------------+
|
v
+-----------------+
| PRコメント / CLI |
|出力 |
+-----------------+
マルチパス定数伝播エンジンは、変数割り当て、os.getenv() / process.env.X フォールバック、関数のデフォルト、クラス属性、コンストラクター引数、dict およびオブジェクト リテラル、**kwargs アンパック、および Vercel AI SDK プロバイダー ラッパー ( openai("gpt-4o") ) を通じてモデル名を解決するため、間接指定を含む実際のコードでも有用な推定値が生成されます。
価格はバンドルされており、オフラインでも機能します。 LiteLLM から更新するには:
トークントールのアップデート
対象範囲: OpenAI、Anthropic、Google、AWS Bedrock、Azure などの 300 以上のモデルに加え、LiteLLM の統合カタログからの 2200 以上のエントリ。
静的解析のみ。データベースまたはリモート構成からロードされたモデルは解決できません。 tokentoll は、構成された SDK ごとのデフォルトにフォールバックし、呼び出しサイトを (デフォルト) としてマークします。
tiktoken がインストールされていない限り ( pip install tokentoll[tiktoken] )、トークンの推定には文字数/4 ヒューリスティックが使用されます。
毎月の見積もりは、通話サイトごとの均一な通話量を前提としています。プロジェクトごとに calls_per_month を使用してオーバーライドするか、パスごとに overrides を使用してオーバーライドします。
JS/TS 解決は同一ファイルのみです。別のモジュールからモデル名をインポートすると、解決された値ではなく動的呼び出しサイトが生成されます。
v0.9 : 失敗がわかっている PR、gpt-researcher のケーススタディ、拡張された採用セクションを含む公開デモ リポジトリ
将来: コンテキストを意識した通話頻度の推論 (FastAPI ルートと s

クリプトとループ); JS/TS のクロスファイルインポート解決
コードレビューでLLMコストの変更を把握します。 LLM 支出のインフラコスト。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
14
v0.8.3: diff エンジンがファントム +/- ペアを出力しないようにする
最新の
2026 年 5 月 30 日
+ 13 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Catch LLM cost changes in code review. Infracost for LLM spend. - Jwrede/tokentoll

GitHub - Jwrede/tokentoll: Catch LLM cost changes in code review. Infracost for LLM spend. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
Jwrede
/
tokentoll
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
54 Commits 54 Commits .github/ workflows .github/ workflows demo demo docs docs src/ tokentoll src/ tokentoll tests tests .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Prevent LLM cost regressions before production.
tokentoll is a CI gate for LLM cost. It statically analyzes Python, JavaScript, and TypeScript for LLM API calls, scores every pull request against a policy you control, and posts a PASS/WARN/FAIL verdict directly on the PR. Optionally, it fails the workflow when the policy is violated, so cost regressions cannot be merged.
Jwrede/tokentoll-demo is a small polyglot LLM app (Python + TypeScript) wired up to the tokentoll cost gate. Two PRs are already open against it:
PR #1: Add Anthropic Haiku translation helper . New call site, well within budget. Verdict: PASS, workflow green.
PR #2: switch supportbot to gpt-4o . A model swap that trips two policy rules. Verdict: FAIL, workflow red.
Open each PR's conversation tab to see the verdict comment tokentoll actually posts.
When a PR violates your policy, tokentoll comments with a verdict and a blocking-findings list, then exits non-zero so the check fails. Example:
## tokentoll verdict: FAIL
** Blocking findings (2): **
- ` src/agent.py:42 ` - per-call cost grew 15.0x (threshold 5x)
- total monthly delta +$812.00 exceeds budget $250.00
> Required action: revert the regression, raise the threshold in ` .tokentoll.yml ` , or add an exemption.
When the PR is clean, the verdict is PASS and the comment shows only the cost delta table. When no policy is configured, tokentoll posts an informational delta comment with no verdict.
Add .github/workflows/tokentoll.yml :
name : tokentoll
on :
pull_request :
paths :
- " **.py "
- " **.ts "
- " **.tsx "
- " **.js "
- " **.jsx "
permissions :
contents : read
pull-requests : write
jobs :
cost-gate :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
with :
fetch-depth : 0
- uses : Jwrede/tokentoll@v0.7.0
with :
fail-on-policy-violation : true
Then add .tokentoll.yml to your repo root:
budgets :
max_monthly_delta_usd : 250
max_callsite_monthly_usd : 100
max_relative_increase : 5.0
policies :
block_unknown_models : true
fail_on_policy_violation : true
Future PRs receive a verdict comment. PRs that exceed the thresholds fail the workflow.
For SHA-pinned installs and minimal-permissions setups, see docs/github-action.md . For the full policy schema, see docs/policy.md . For the security posture, see docs/security.md .
JavaScript / TypeScript (parsed via tree-sitter, handles .js , .jsx , .ts , .tsx )
The policy block in .tokentoll.yml controls when a PR fails:
Each rule is independent. Leave a field unset to disable that rule. Full reference in docs/policy.md .
pip install tokentoll
# Scan current directory for LLM API calls and their costs
tokentoll scan .
# Show cost impact of your last commit
tokentoll diff HEAD~1
# Compare two refs and fail on policy violation
tokentoll diff main..HEAD --fail-on-policy-violation
Subcommands:
tokentoll scan [PATH...] [--format table|json|markdown] [--calls-per-month N] [--config PATH]
tokentoll diff [REF] [--base REF] [--head REF] [--format table|json|markdown|github-comment]
[--config PATH] [--fail-on-policy-violation]
tokentoll update # refresh bundled pricing data from LiteLLM
Configuration
.tokentoll.yml lives in the repo root and is auto-discovered. Beyond the policy block:
# Per-SDK defaults for dynamic (runtime-resolved) model names
default_models :
openai : gpt-4o-mini
anthropic : claude-haiku-3-20240307
# Assumed monthly call volume per call site (used for dollar estimates)
calls_per_month : 5000
# Skip cost estimation for dynamic models entirely.
# Default false: dynamic calls are priced against the per-SDK default.
skip_dynamic_models : false
# Default excludes (tests/, examples/, docs/, cookbook/, benchmarks/, evals/,
# scripts/, notebooks/) are applied automatically. Opt out with:
use_default_excludes : false
# Additional excludes (prefix or glob)
exclude :
- " *_test.py "
- vendor/
# Per-path overrides (longest prefix match)
overrides :
- path : src/agents/
default_model : gpt-4o
calls_per_month : 10000
- path : src/azure/
skip_dynamic_models : true
Resolution order for dynamic model defaults: default_models (per-SDK) > default_model (generic) > built-in SDK defaults.
tokentoll requires no API keys, sends no telemetry, and runs entirely inside your CI environment. Pricing data ships with the package and updates from LiteLLM on demand. For the recommended permission set, SHA pinning, and fork PR risk, see docs/security.md .
tokentoll ships an MCP (Model Context Protocol) server so Claude Code and other MCP hosts can check the cost impact of LLM code changes from inside an agent conversation:
pip install tokentoll[mcp]
claude mcp add --transport stdio tokentoll -- tokentoll-mcp
Two tools are exposed: scan (estimate costs across a path) and diff (compare two refs). Both return JSON.
Source code (.py, .ts, .tsx, .js, .jsx)
|
v
+----------------+ +------------------+
| AST scanners |-->| SDK detectors |
| ast (Python) + | | OpenAI, Anthropic|
| tree-sitter | | Google, LiteLLM, |
| (JS/TS) | | LangChain, Zhipu,|
+----------------+ | Vercel AI SDK |
+------------------+
|
v
+------------------+
| Pricing engine |
| 2200+ models |
+------------------+
|
v
+------------------+
| Diff engine |
| (old vs new) |
+------------------+
|
v
+------------------+
| Policy evaluator |
| PASS/WARN/FAIL |
+------------------+
|
v
+------------------+
| PR comment / CLI |
| output |
+------------------+
A multi-pass constant propagation engine resolves model names through variable assignments, os.getenv() / process.env.X fallbacks, function defaults, class attributes, constructor arguments, dict and object literals, **kwargs unpacking, and Vercel AI SDK provider wrappers ( openai("gpt-4o") ), so real-world code with indirection still produces useful estimates.
Pricing is bundled and works offline. To refresh from LiteLLM:
tokentoll update
Coverage: 300+ models across OpenAI, Anthropic, Google, AWS Bedrock, Azure, and more, plus 2200+ entries from LiteLLM's combined catalog.
Static analysis only. Models loaded from databases or remote config cannot be resolved; tokentoll falls back to the configured per-SDK default and marks the call site as (default) .
Token estimates use a characters/4 heuristic unless tiktoken is installed ( pip install tokentoll[tiktoken] ).
Monthly estimates assume uniform call volume per call site. Override per-project with calls_per_month or per-path with overrides .
JS/TS resolution is same-file only. Importing a model name from another module produces a dynamic call site rather than a resolved value.
v0.9 : Public demo repo with a known-failing PR, gpt-researcher case study, expanded adoption section
Future : Context-aware call frequency inference (FastAPI routes versus scripts versus loops); cross-file import resolution for JS/TS
Catch LLM cost changes in code review. Infracost for LLM spend.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
14
v0.8.3: stop diff engine from emitting phantom +/- pairs
Latest
May 30, 2026
+ 13 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
