---
source: "https://github.com/Kirill89/reviewcerberus"
hn_url: "https://news.ycombinator.com/item?id=48787451"
title: "Show HN: AI-powered code review tool"
article_title: "GitHub - Kirill89/reviewcerberus: AI-powered code review tool that analyzes git branch differences and generates comprehensive review reports. Supports AWS Bedrock and Anthropic API. Features automated analysis of logic, security, performance, and code quality with smart token efficiency through pro\n[truncated]"
author: "k1r111"
captured_at: "2026-07-04T19:01:42Z"
capture_tool: "hn-digest"
hn_id: 48787451
score: 1
comments: 0
posted_at: "2026-07-04T18:07:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-powered code review tool

- HN: [48787451](https://news.ycombinator.com/item?id=48787451)
- Source: [github.com](https://github.com/Kirill89/reviewcerberus)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T18:07:22Z

## Translation

タイトル: Show HN: AI を活用したコード レビュー ツール
記事タイトル: GitHub - Kirill89/reviewcerberus: git ブランチの相違点を分析し、包括的なレビュー レポートを生成する AI を活用したコード レビュー ツール。 AWS Bedrock および Anthropic API をサポートします。プロによるスマートトークン効率によるロジック、セキュリティ、パフォーマンス、コード品質の自動分析を備えています。
[切り捨てられた]
説明: git ブランチの違いを分析し、包括的なレビュー レポートを生成する、AI を活用したコード レビュー ツール。 AWS Bedrock および Anthropic API をサポートします。ロジック、セキュリティ、パフォーマンス、コード品質の自動分析と、迅速なキャッシュによるスマート トークンの効率性を備えています。 - Kirill89/reviewcerb
[切り捨てられた]

記事本文:
GitHub - Kirill89/reviewcerberus: git ブランチの違いを分析し、包括的なレビュー レポートを生成する、AI を活用したコード レビュー ツール。 AWS Bedrock および Anthropic API をサポートします。ロジック、セキュリティ、パフォーマンス、コード品質の自動分析と、迅速なキャッシュによるスマート トークンの効率性を備えています。 · GitHub
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
アカウントを切り替えました

別のタブまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
キリル89
/
レビューケルベロス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
59 コミット 59 コミット .github/ workflows .github/ workflows act-test act-test action action spec spec src src testing testing .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore DOCKERHUB.md DOCKERHUB.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md logo.png logo.png logo_256.png logo_256.png 詩.ロック 詩.ロック pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
git ブランチの違いを分析し、生成する AI を活用したコード レビュー ツール
構造化された出力を含む包括的なレビューレポート。
GitHub Action : インライン コメントと概要を使用した自動 PR レビュー
包括的なレビュー : ロジック、セキュリティ、パフォーマンス、
コードの品質
構造化された出力: 重大度別に整理された問題と要約表
マルチプロバイダー: AWS Bedrock、Anthropic API、Ollama、または Moonshot
スマートな分析: プロンプト キャッシュにより事前に提供されるコンテキスト
Git 統合 : あらゆるリポジトリで動作し、コミット ハッシュをサポートします
検証モード : 実験的
虚偽を減らすための検証の連鎖
ポジティブ
Docker で実行します (推奨):
docker run --rm -it -v $( pwd ) :/repo \
-e MODEL_PROVIDER=人類 \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
キリル89/レビューケルベロス:最新\
--repo-path /repo --output /repo/review.md
それです！レビューは現在のファイルの review.md に保存されます。
ディレクトリ。
「AWS Bedrock セットアップとその他のオプションの設定」を参照してください。
自動 PR レビューの場合は、 .github/workflows/review.yml に追加します。
名前 :

コードレビュー
に:
プルリクエスト:
タイプ: [オープン、同期]
ジョブ:
レビュー：
実行: ubuntu-最新
権限:
内容：書く
プルリクエスト : 書き込み
手順:
- 使用:actions/checkout@v4
付き:
フェッチ深度: 0
- 使用: Kirill89/reviewcerberus/action@v1
付き:
モデルプロバイダー : 人間
anthropic_api_key : ${{ Secrets.ANTHROPIC_API_KEY }}
アクションは、PR にレビュー コメントを直接投稿します。参照
すべてのオプションの GitHub アクション。
# コードレビューを実行する
ポエトリー・ラン・レビューケルベロス
# カスタムターゲットブランチ
詩の実行レビューcerberus --target-branch開発
# カスタム出力場所
詩の実行 reviewcerberus --output /path/to/review.md
詩を実行します reviewcerberus --output /path/to/dir/ # ファイル名を自動生成します
# マークダウンではなく JSON として出力します
ポエトリーランレビューcerberus --json
# 別のリポジトリ
詩の実行 reviewcerberus --repo-path /path/to/repo
# カスタムレビューガイドラインを追加
詩の実行レビューcerberus --instructionsガイドライン.md
# 検証モードを有効にする (実験的)
ポエトリーランレビューcerberus --verify
# SAST プレスキャンを有効にする (実験的)
ポエトリー・ラン・レビューcerberus --sast
コマンド例
# カスタム ガイドラインを使用した完全なレビュー
詩の実行レビューcerberus --target-branch main \
--出力レビュー.md --指示ガイドライン.md
# 別のリポジトリを確認する
詩の実行レビューcerberus --repo-path /other/repo
含まれるもの
論理と正確さ : バグ、エッジケース、エラー処理
セキュリティ : OWASP の問題、アクセス制御、入力検証
パフォーマンス : N+1 クエリ、ボトルネック、スケーラビリティ
コードの品質 : 重複、複雑さ、保守性
副作用 : 他のシステム部分への影響
テスト : カバレッジギャップ、テストケースの欠落
ドキュメント : ドキュメントが欠落しているか古い、コメントが不明瞭
概要 : 変更とリスク領域の高レベルの概要
問題テーブル: サーバーに関するすべての問題が一目でわかる

重要度の指標 (🔴 クリティカル、
🟠高、🟡中、🟢低)
詳細な問題: 各問題の説明、場所、および推奨される修正
検証モード (実験的)
--verify フラグを使用して有効にすると、誤検知を削減できます。
検証チェーン (CoVe) :
質問の生成 : 問題ごとに改ざん質問を作成します
質問に答える : コードコンテキストを使用して質問に答えます
信頼度スコア: 証拠に基づいて 1 ～ 10 の信頼度スコアを割り当てます。
出力内の各問題には、信頼度スコアと根拠が含まれています。
SAST 統合 (実験的)
--sast フラグを使用して有効にして、
OpenGrep (Semgrep フォーク) の前の事前スキャン
AI レビュー:
現在のブランチによって導入された新しい検出結果のみをスキャンします
調査結果は補足的なコンテキストとして AI エージェントに提供されます
エージェントは各結果を個別に検証し、誤検知を却下します
静的分析の精度と AI のコンテキスト理解の組み合わせ
現在の git ブランチとリポジトリを検出します
すべてのコンテキストを事前に収集します: 変更されたファイル、コミットメッセージ、差分
以下にアクセスできる AI エージェントを使用して分析します。
完全な diff コンテキスト (ファイルあたり 10,000 文字で切り捨てられます)
コードベース全体のパターン検索
マークダウンとしてレンダリングされた構造化レビュー出力を生成します
リポジトリ: /path/to/repo
現在のブランチ: 機能ブランチ
ターゲットブランチ: main
3 つの変更されたファイルが見つかりました:
- src/main.py (修正)
- src/utils.py (修正)
- test/test_main.py (追加)
コードレビューを開始しています...
🤔 考え中... ⏱️ 3.0秒
🔧 read_file_part: src/main.py
✓ レビュー完了: review_feature-branch.md
トークンの使用法:
入力トークン: 6,856
出力トークン: 1,989
合計トークン: 8,597
構成
すべての設定は環境変数 (.env ファイル) を介して行われます。
MODEL_PROVIDER=bedrock # または "anthropic"、"ollama"、"moonshot" (デフォルト: bedrock)
AWS Bedrock (MODEL_PROVIDER=bedrock の場合)
AWS_ACCESS_KEY_ID=あなた

r_key
AWS_SECRET_ACCESS_KEY=あなたの秘密
AWS_REGION_NAME=us-east-1
MODEL_NAME=us.anthropic.claude-opus-4-5-20251101-v1:0 # オプション
Bedrock を使用した Docker の例:
docker run --rm -it -v $( pwd ) :/repo \
-e AWS_ACCESS_KEY_ID=your_key \
-e AWS_SECRET_ACCESS_KEY=your_secret \
-e AWS_REGION_NAME=us-east-1 \
キリル89/レビューケルベロス:最新\
--repo-path /repo --output /repo/review.md
Anthropic API (MODEL_PROVIDER=anthropic の場合)
ANTHROPIC_API_KEY=sk-ant-your-api-key-here
MODEL_NAME=claude-opus-4-5-20251101 # オプション
オラマ (MODEL_PROVIDER=オラマの場合)
MODEL_PROVIDER=オラマ
OLLAMA_BASE_URL=http://localhost:11434 # オプション、デフォルト
MODEL_NAME=deepseek-v3.1:671b-cloud # オプション
Ollama を使用した Docker の例:
# Ollama がホスト マシン上で実行されていることを前提としています
docker run --rm -it -v $( pwd ) :/repo \
-e MODEL_PROVIDER=オラマ \
-e OLLAMA_BASE_URL=http://host.docker.internal:11434 \
キリル89/レビューケルベロス:最新\
--repo-path /repo --output /repo/review.md
ムーンショット (MODEL_PROVIDER=ムーンショットの場合)
MODEL_PROVIDER=ムーンショット
MOONSHOT_API_KEY=sk-your-api-key-here
MOONSHOT_API_BASE=https://api.moonshot.ai/v1 # オプション、デフォルト
MODEL_NAME=kimi-k2.5 # オプション
オプション設定
MAX_OUTPUT_TOKENS=10000 # 応答の最大トークン数
TOOL_CALL_LIMIT=100 # 出力を強制するまでのツール呼び出しの最大数
VERIFY_MODEL_NAME=... # 検証用のモデル (デフォルトは MODEL_NAME)
カスタムレビュープロンプト
src/agent/prompts/ のプロンプトをカスタマイズします。
full_review.md - メインレビュープロンプト
context_summary.md - 大規模な PR のコンテキスト圧縮
PR レビューを自動化するには、ReviewCerberus を GitHub アクションとして使用します。
入力
説明
デフォルト
モデルプロバイダー
プロバイダ: bedrock 、 anthropic 、 ollama 、または Moonshot
岩盤
anthropic_api_key
人間の API キー
-
aws_access_key_id
AWS アクセスキー ID (Bedrock)
-
aws_secret_access_key
AWS シークレット アクセス キー (ベッド)

ロック）
-
aws_region_name
AWS リージョン (Bedrock)
米国東部-1
モデル名
モデル名 (プロバイダー固有)
-
検証する
検証チェーンを有効にする
偽
サスト
OpenGrep SAST 事前スキャンを有効にする
偽
min_confidence
最小信頼スコア 1 ～ 10 (検証が必要)
-
フェイルオン
この重大度以上の問題がある場合は失敗します: クリティカル、高、中、低
-
指示
カスタムレビューガイドラインへのパス
-
検証付きの例
- 使用: Kirill89/reviewcerberus/action@v1
付き:
モデルプロバイダー : 人間
anthropic_api_key : ${{ Secrets.ANTHROPIC_API_KEY }}
検証: " true "
min_confidence : " 7 "
SAST を使用した例
- 使用: Kirill89/reviewcerberus/action@v1
付き:
モデルプロバイダー : 人間
anthropic_api_key : ${{ Secrets.ANTHROPIC_API_KEY }}
sast : "本当"
クオリティゲートとしての例
- 使用: Kirill89/reviewcerberus/action@v1
付き:
モデルプロバイダー : 人間
anthropic_api_key : ${{ Secrets.ANTHROPIC_API_KEY }}
失敗時: " 高 "
AWS Bedrock を使用した例
- 使用: Kirill89/reviewcerberus/action@v1
付き:
モデルプロバイダー : 岩盤
aws_access_key_id : ${{ Secrets.AWS_ACCESS_KEY_ID }}
aws_secret_access_key : ${{ Secrets.AWS_SECRET_ACCESS_KEY }}
aws_region_name : us-east-1
アクションの内容
Docker イメージを使用してレビューを実行します
以前の実行による既存のレビュー スレッドを解決します。
すべての問題を含む要約コメントを投稿します
特定の行にインラインレビューコメントを作成します
ローカル開発の場合 (Docker の使用には必要ありません):
# クローンを作成してインストールする
git clone < リポジトリ URL >
詩のインストール
# 資格情報を構成する
cp .env.example .env
# プロバイダーの資格情報を使用して .env を編集します
資格情報のセットアップの構成を参照してください。
テストを行う
# または
詩を実行する pytest -v
結合テスト（行為）
act を使用した GitHub アクションのエンドツーエンド テスト
モック Ollama および GitHub API サーバーを使用:
アクトテストを作成する
前提条件: Docker と
行為はインスタでなければならない

導かれた。
これにより、Docker イメージが構築され、モック サーバーが起動され、完全なアクション ワークフローが実行されます。
ローカルで、記録された API リクエストを vitest で検証します。
make lint # mypy、isort、black、mdformat で確認
make format # isort と black で自動フォーマットする
Docker イメージの構築
make docker-build # ローカルでビルドする
make docker-build-push # ビルドとプッシュ (マルチプラットフォーム)
バージョンは pyproject.toml から自動で読み取られます。詳細については、DOCKER.md を参照してください。
詳細。
§── src/ # Python CLI
│ §── config.py # 設定
│ §── main.py # CLI エントリポイント
│ └── エージェント/
│ §── Agent.py # エージェントのセットアップ
│ §── model.py # モデルの初期化
│ §──runner.py # 実行のレビュー
│ §── プロンプト/ # プロンプトを確認する
│ §── schema.py # データモデル (構造化出力)
│ §── git_utils/ # Git の操作
│ §── formatting/ # コンテキストと出力のレンダリング
│ §── 検証/ # 検証連鎖パイプライン
│ §── progress_callback_handler.py
│ ━─ ツール/#3 レビューツール
│
└── action/ # GitHub アクション (TypeScript)
§── action.yml # アクション定義
§── src/ # アクションのソースコード
└── dist/ # バンドルされたアクション
コードの品質基準
厳密な型チェック : すべての関数には型注釈が必要です
戻り値の型: 明示的である必要があります ( warn_return_any = true )
書式設定 : 黒 + 黒プロファイルのアイソート

[切り捨てられた]

## Original Extract

AI-powered code review tool that analyzes git branch differences and generates comprehensive review reports. Supports AWS Bedrock and Anthropic API. Features automated analysis of logic, security, performance, and code quality with smart token efficiency through prompt caching. - Kirill89/reviewcerb
[truncated]

GitHub - Kirill89/reviewcerberus: AI-powered code review tool that analyzes git branch differences and generates comprehensive review reports. Supports AWS Bedrock and Anthropic API. Features automated analysis of logic, security, performance, and code quality with smart token efficiency through prompt caching. · GitHub
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
Kirill89
/
reviewcerberus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
59 Commits 59 Commits .github/ workflows .github/ workflows act-test act-test action action spec spec src src tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore DOCKERHUB.md DOCKERHUB.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md logo.png logo.png logo_256.png logo_256.png poetry.lock poetry.lock pyproject.toml pyproject.toml View all files Repository files navigation
AI-powered code review tool that analyzes git branch differences and generates
comprehensive review reports with structured output.
GitHub Action : Automated PR reviews with inline comments and summary
Comprehensive Reviews : Detailed analysis of logic, security, performance,
and code quality
Structured Output : Issues organized by severity with summary table
Multi-Provider : AWS Bedrock, Anthropic API, Ollama, or Moonshot
Smart Analysis : Context provided upfront with prompt caching
Git Integration : Works with any repository, supports commit hashes
Verification Mode : Experimental
Chain-of-Verification to reduce false
positives
Run with Docker (recommended):
docker run --rm -it -v $( pwd ) :/repo \
-e MODEL_PROVIDER=anthropic \
-e ANTHROPIC_API_KEY=sk-ant-your-api-key \
kirill89/reviewcerberus:latest \
--repo-path /repo --output /repo/review.md
That's it! The review will be saved to review.md in your current
directory.
See Configuration for AWS Bedrock setup and other options.
For automated PR reviews, add to .github/workflows/review.yml :
name : Code Review
on :
pull_request :
types : [opened, synchronize]
jobs :
review :
runs-on : ubuntu-latest
permissions :
contents : write
pull-requests : write
steps :
- uses : actions/checkout@v4
with :
fetch-depth : 0
- uses : Kirill89/reviewcerberus/action@v1
with :
model_provider : anthropic
anthropic_api_key : ${{ secrets.ANTHROPIC_API_KEY }}
The action posts review comments directly on your PR. See
GitHub Action for all options.
# Run code review
poetry run reviewcerberus
# Custom target branch
poetry run reviewcerberus --target-branch develop
# Custom output location
poetry run reviewcerberus --output /path/to/review.md
poetry run reviewcerberus --output /path/to/dir/ # Auto-generates filename
# Output as JSON instead of markdown
poetry run reviewcerberus --json
# Different repository
poetry run reviewcerberus --repo-path /path/to/repo
# Add custom review guidelines
poetry run reviewcerberus --instructions guidelines.md
# Enable verification mode (experimental)
poetry run reviewcerberus --verify
# Enable SAST pre-scan (experimental)
poetry run reviewcerberus --sast
Example Commands
# Full review with custom guidelines
poetry run reviewcerberus --target-branch main \
--output review.md --instructions guidelines.md
# Review a different repo
poetry run reviewcerberus --repo-path /other/repo
What's Included
Logic & Correctness : Bugs, edge cases, error handling
Security : OWASP issues, access control, input validation
Performance : N+1 queries, bottlenecks, scalability
Code Quality : Duplication, complexity, maintainability
Side Effects : Impact on other system parts
Testing : Coverage gaps, missing test cases
Documentation : Missing or outdated docs, unclear comments
Summary : High-level overview of changes and risky areas
Issues Table : All issues at a glance with severity indicators (🔴 CRITICAL,
🟠 HIGH, 🟡 MEDIUM, 🟢 LOW)
Detailed Issues : Each issue with explanation, location, and suggested fix
Verification Mode (Experimental)
Enable with --verify flag to reduce false positives using
Chain-of-Verification (CoVe) :
Generate Questions : Creates falsification questions for each issue
Answer Questions : Answers questions using code context
Score Confidence : Assigns 1-10 confidence score based on evidence
Each issue in the output includes a confidence score and rationale.
SAST Integration (Experimental)
Enable with --sast flag to run an
OpenGrep (Semgrep fork) pre-scan before
the AI review:
Scans only new findings introduced by the current branch
Findings are provided to the AI agent as supplementary context
The agent independently verifies each finding and dismisses false positives
Combines static analysis precision with AI contextual understanding
Detects current git branch and repository
Collects all context upfront: changed files, commit messages, and diffs
Analyzes using AI agent with access to:
Full diff context (truncated at 10k chars per file)
Pattern search across codebase
Generates structured review output rendered as markdown
Repository: /path/to/repo
Current branch: feature-branch
Target branch: main
Found 3 changed files:
- src/main.py (modified)
- src/utils.py (modified)
- tests/test_main.py (added)
Starting code review...
🤔 Thinking... ⏱️ 3.0s
🔧 read_file_part: src/main.py
✓ Review completed: review_feature-branch.md
Token Usage:
Input tokens: 6,856
Output tokens: 1,989
Total tokens: 8,597
Configuration
All configuration via environment variables ( .env file):
MODEL_PROVIDER=bedrock # or "anthropic", "ollama", "moonshot" (default: bedrock)
AWS Bedrock (if MODEL_PROVIDER=bedrock)
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
AWS_REGION_NAME=us-east-1
MODEL_NAME=us.anthropic.claude-opus-4-5-20251101-v1:0 # optional
Docker example with Bedrock:
docker run --rm -it -v $( pwd ) :/repo \
-e AWS_ACCESS_KEY_ID=your_key \
-e AWS_SECRET_ACCESS_KEY=your_secret \
-e AWS_REGION_NAME=us-east-1 \
kirill89/reviewcerberus:latest \
--repo-path /repo --output /repo/review.md
Anthropic API (if MODEL_PROVIDER=anthropic)
ANTHROPIC_API_KEY=sk-ant-your-api-key-here
MODEL_NAME=claude-opus-4-5-20251101 # optional
Ollama (if MODEL_PROVIDER=ollama)
MODEL_PROVIDER=ollama
OLLAMA_BASE_URL=http://localhost:11434 # optional, default
MODEL_NAME=deepseek-v3.1:671b-cloud # optional
Docker example with Ollama:
# Assumes Ollama running on host machine
docker run --rm -it -v $( pwd ) :/repo \
-e MODEL_PROVIDER=ollama \
-e OLLAMA_BASE_URL=http://host.docker.internal:11434 \
kirill89/reviewcerberus:latest \
--repo-path /repo --output /repo/review.md
Moonshot (if MODEL_PROVIDER=moonshot)
MODEL_PROVIDER=moonshot
MOONSHOT_API_KEY=sk-your-api-key-here
MOONSHOT_API_BASE=https://api.moonshot.ai/v1 # optional, default
MODEL_NAME=kimi-k2.5 # optional
Optional Settings
MAX_OUTPUT_TOKENS=10000 # Maximum tokens in response
TOOL_CALL_LIMIT=100 # Maximum tool calls before forcing output
VERIFY_MODEL_NAME=... # Model for verification (defaults to MODEL_NAME)
Custom Review Prompts
Customize prompts in src/agent/prompts/ :
full_review.md - Main review prompt
context_summary.md - Context compaction for large PRs
Use ReviewCerberus as a GitHub Action for automated PR reviews.
Input
Description
Default
model_provider
Provider: bedrock , anthropic , ollama , or moonshot
bedrock
anthropic_api_key
Anthropic API key
-
aws_access_key_id
AWS Access Key ID (Bedrock)
-
aws_secret_access_key
AWS Secret Access Key (Bedrock)
-
aws_region_name
AWS Region (Bedrock)
us-east-1
model_name
Model name (provider-specific)
-
verify
Enable Chain-of-Verification
false
sast
Enable OpenGrep SAST pre-scan
false
min_confidence
Min confidence score 1-10 (requires verify)
-
fail_on
Fail if issues at or above this severity: critical , high , medium , low
-
instructions
Path to custom review guidelines
-
Example with Verification
- uses : Kirill89/reviewcerberus/action@v1
with :
model_provider : anthropic
anthropic_api_key : ${{ secrets.ANTHROPIC_API_KEY }}
verify : " true "
min_confidence : " 7 "
Example with SAST
- uses : Kirill89/reviewcerberus/action@v1
with :
model_provider : anthropic
anthropic_api_key : ${{ secrets.ANTHROPIC_API_KEY }}
sast : " true "
Example as Quality Gate
- uses : Kirill89/reviewcerberus/action@v1
with :
model_provider : anthropic
anthropic_api_key : ${{ secrets.ANTHROPIC_API_KEY }}
fail_on : " high "
Example with AWS Bedrock
- uses : Kirill89/reviewcerberus/action@v1
with :
model_provider : bedrock
aws_access_key_id : ${{ secrets.AWS_ACCESS_KEY_ID }}
aws_secret_access_key : ${{ secrets.AWS_SECRET_ACCESS_KEY }}
aws_region_name : us-east-1
What the Action Does
Runs the review using the Docker image
Resolves any existing review threads from previous runs
Posts a summary comment with all issues
Creates inline review comments on specific lines
For local development (not required for Docker usage):
# Clone and install
git clone < repo-url >
poetry install
# Configure credentials
cp .env.example .env
# Edit .env with your provider credentials
See Configuration for credential setup.
make test
# or
poetry run pytest -v
Integration Test (act)
End-to-end test of the GitHub Action using act
with mock Ollama and GitHub API servers:
make act-test
Prerequisites: Docker and
act must be installed.
This builds the Docker image, starts mock servers, runs the full action workflow
locally, then verifies the recorded API requests with vitest.
make lint # Check with mypy, isort, black, mdformat
make format # Auto-format with isort and black
Building Docker Image
make docker-build # Build locally
make docker-build-push # Build and push (multi-platform)
Version is auto-read from pyproject.toml . See DOCKER.md for
details.
├── src/ # Python CLI
│ ├── config.py # Configuration
│ ├── main.py # CLI entry point
│ └── agent/
│ ├── agent.py # Agent setup
│ ├── model.py # Model initialization
│ ├── runner.py # Review execution
│ ├── prompts/ # Review prompts
│ ├── schema.py # Data models (structured output)
│ ├── git_utils/ # Git operations
│ ├── formatting/ # Context and output rendering
│ ├── verification/ # Chain-of-Verification pipeline
│ ├── progress_callback_handler.py
│ └── tools/ # 3 review tools
│
└── action/ # GitHub Action (TypeScript)
├── action.yml # Action definition
├── src/ # Action source code
└── dist/ # Bundled action
Code Quality Standards
Strict type checking : All functions require type annotations
Return types : Must be explicit ( warn_return_any = true )
Formatting : Black + isort with black profile

[truncated]
