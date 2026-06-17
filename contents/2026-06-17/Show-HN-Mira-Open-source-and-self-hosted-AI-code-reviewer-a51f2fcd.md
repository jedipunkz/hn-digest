---
source: "https://github.com/miracodeai/mira"
hn_url: "https://news.ycombinator.com/item?id=48570197"
title: "Show HN: Mira – Open-source and self-hosted AI code reviewer"
article_title: "GitHub - miracodeai/mira: Self-hosted AI code reviewer with indexed PR reviews, walkthroughs, vulnerability scanning, dependency graphs, custom rules, and a learning loop. · GitHub"
author: "upmostly"
captured_at: "2026-06-17T16:24:06Z"
capture_tool: "hn-digest"
hn_id: 48570197
score: 5
comments: 2
posted_at: "2026-06-17T13:23:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mira – Open-source and self-hosted AI code reviewer

- HN: [48570197](https://news.ycombinator.com/item?id=48570197)
- Source: [github.com](https://github.com/miracodeai/mira)
- Score: 5
- Comments: 2
- Posted: 2026-06-17T13:23:54Z

## Translation

タイトル: Show HN: Mira – オープンソースおよびセルフホスト AI コードレビューアー
記事のタイトル: GitHub - miracodeai/mira: インデックス付き PR レビュー、ウォークスルー、脆弱性スキャン、依存関係グラフ、カスタム ルール、学習ループを備えたセルフホスト型 AI コード レビューアー。 · GitHub
説明: インデックス付き PR レビュー、ウォークスルー、脆弱性スキャン、依存関係グラフ、カスタム ルール、学習ループを備えた自己ホスト型 AI コード レビューアー。 - ミラコダイ/ミラ
HN テキスト: こんにちは、HN、私は Mira の共同作成者の Jay です。 BYOK (独自のキーの持ち込み) が可能な、オープンソースの自己ホスト型 AI コード レビューアー。ローカルモデルは本当に良くなっている。ホスト型フロンティア モデルは非常に高価になっています。そこで、ミラを作りました。 Mira にはいくつかの優れた機能もあります。 - レビューが非常に速いです。平均は77秒、対してグレプタイルは5分。 PR はクラウド上のどこかのキューに並ぶことはありません。 - ミラは爆発半径を実行し、コード変更でどのようなダメージが与えられるかを確認します。 - リポジトリ自体からコードベースのパターンを学習し、構成ファイルなしでそれらを適用します。 - 自己ホスト型なので、コードがインフラストラクチャから離れることはありません。独自のモデル (OpenAI、Anthropic、またはファイアウォールの背後にあるローカル LLM) を持ち込みます。私たちは非常に活発に活動しており、積極的な貢献者を探しています。私たちのDiscordに参加してください！ https://github.com/miracodeai/mira ありがとうございます。
ジェイ

記事本文:
GitHub - miracodeai/mira: インデックス付き PR レビュー、ウォークスルー、脆弱性スキャン、依存関係グラフ、カスタム ルール、学習ループを備えた自己ホスト型 AI コード レビューアー。 · GitHub
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
ミラコダイ
/
ミラ
公共
ああ、ああ！
ロー中にエラーが発生しました

アディング。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
336 コミット 336 コミット .github .github スクリプト スクリプト src/ mira src/ mira テスト テスト ui/ mira ui/ mira .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .mira.yaml.example .mira.yaml.example .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FEATURES.md FEATURES.md ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md fly.toml fly.toml pyproject.toml pyproject.toml Railway.toml Railway.toml render.yaml render.yaml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自己ホスト型 AI コード レビュー。コード、ダッシュボード、LLM キー。
すべての機能を自己ホストします: 完全なレビュー エンジン、コードベースのインデックス作成、脆弱性スキャン、カスタム ルール、組織全体のパッケージ検索、ダッシュボード、学習ループ。有料枠、ライセンスキー、SaaS アップセルはありません。
Mira は、選択した LLM (Anthropic、OpenAI、Google、DeepSeek などをサポートする OpenRouter 経由) を使用してプル リクエストをレビューし、簡潔で実用的なフィードバックを投稿します。ノイズ フィルター、信頼性クランプ、学習ループにより、重要なコメントのみが表示されるようになります。全表面については FEATURES.md を参照してください。
ほとんどの AI レビュー担当者は SaaS です。差分 (および多くの場合、周囲のコード全体) はサードパーティのサーバーに送信され、取得できる唯一の「ビュー」は PR で返されるコメントです。ミラはその両方を裏返します。
コードがインフラストラクチャから離れることはありません。差分、埋め込み、インデックス、レビュー履歴、脆弱性データはすべて SQLite または Postgres に保存されます。

あなたが所有する構造。テレフォンホームも必要なく、テレメトリーも必要なく、「これはトレーニングに使用されますか?」という質問もありません。質問。
上に表示されているダッシュボードはあなたのものです。これは、あなたのコードを他人から見たマーケティング用のスクリーンショットではありません。 CodeRabbit、Greptile、および同様の SaaS レビュー担当者は、そのようなものを公開していません。 Mira のダッシュボードには、他では得られないシグナルが表示されます。
組織全体のパッケージ インベントリ: 「どのリポジトリが lodash@4.17.20 を使用しますか?」と回答します。 1 つのクエリで。 CVE フィードの隣に積み重ねて、瞬時に爆発半径をチェックします。
すべての依存関係に関する CVE アラート: 時間ごとの OSV.dev ポーリング、重大度 + アドバイザリ リンク + 修正バージョンがパッケージの横にインラインで表示されます。
依存関係 + ブラスト半径グラフ : シンボルを変更する前に、どのファイルとリポジトリがシンボルに依存しているかを正確に確認します。
リポジトリごとのレビュー イベント ストリーム: すべての Webhook、すべてのチャンク、すべてのコスト数値を 1 か所でライブ トラブルシューティングできます。
コストとトークンのテレメトリ: LLM キーを制御するため、見積もりではなく、リポジトリごとおよびモデルごとの実際の支出です。
近日公開予定の変更頻度ヒートマップ: バグ修正が継続的に適用されるファイルを表面化することで、レビューの注目を集めることができます。
エンジニアリング チームが「どのリポジトリがこの CVE に公開されているか?」などの回答が必要な場合は、または「この機能を変更した場合の影響範囲はどれくらいですか?」 、これらの質問は数日かかる調査ではなくなり、ワンクリックのダッシュボード ページになり始めます。
OpenAI 互換エンドポイント: デフォルトで OpenRouter (Anthropic、OpenAI、Google Gemini、DeepSeek など)、または自己ホスト型/データ常駐レビューの場合は、vLLM、Ollama、LiteLLM プロキシ、LocalAI、llama.cpp、Togetter、Fireworks、Groq、または Cerebras で llm.base_url を指定します。 「カスタム エンドポイント」を参照してください。
低ノイズ : 信頼性のしきい値、重複排除、重大度の並べ替え、PR ごとのコメントの上限
インデックス付きレビュー : フルリポジトリコードインデックスは、単なる差分ではなく、LLM の実際のコンテキストを提供します。
あなたのお茶を学びます

m : 拒否されたコメントとマージされた PR に対する人間によるレビュー パターンからルールを合成します。
脆弱性スキャン: 時間ごとの OSV.dev ポーリングにより、すべてのリポジトリのすべてのパッケージにわたる CVE が明らかになります
組織全体のパッケージ検索: 「lodash@4.17.20 を使用するリポジトリはどれですか?」と回答します。数秒で
構成可能: リポジトリごとの設定用の .mira.yaml、ダッシュボードのカスタム + グローバル ルール
初日のセルフホスト: Docker イメージ、Railway / Fly.io / レンダー構成、SQLite または Postgres
本日、GitHub にさらに多くのプラットフォームが登場します: GitLab、Bitbucket、Gitea のサポートがロードマップにあります。同じレビュー エンジン、同じダッシュボード、プロバイダー アダプターが異なるだけ
Mira は、公開コード レビュー ベンチで測定された最速のツールであり、速度と品質のパレート フロンティアにある唯一のツールです。F1 でより高いスコアを獲得するすべてのツールは、PR ごとに 5 ～ 14 倍の時間がかかります。
同じサブセット上の公開されているすべての競合他社に対してプロットすると、Mira は左上隅に位置します。右にあるものはすべて遅いです。それを超えるものはすべて、追加の F1 に対して壁時間の 5 ～ 14 倍を支払います。
同じ 50 PR のオフライン ベンチマークで測定され、Claude Sonnet 4.5 によって判定されます。
方法論: Claude Sonnet 4.5 を審査員として、Martian Code Review Bench オフライン データセットに対してスコアを測定しました。
すべての PR を自動レビューし、コメントに応答する GitHub アプリとして Mira を実行します。
注: 現在サポートされているプロバイダーは GitHub のみです。 GitLab、Bitbucket、および Gitea アダプターがロードマップに含まれています。レビュー エンジン、インデクサー、ダッシュボードはプロバイダーに依存しません。 Webhook + コメント投稿レイヤーのみが GitHub 固有です。リポジトリにスターを付けてフォローしてください。
1. github.com/settings/apps/new で GitHub アプリを作成します。
Webhook URL: https://your-server.com/github/webhook
権限: プル リクエスト (読み取り + 書き込み)、コンテンツ (読み取り + 書き込み)、問題 (読み取り + 書き込み)
イベント: プル リクエスト、コメントの発行
2 つのファイル: 非シークレット用の mira.yaml

フォルト、シークレットの .env。両方を容器に移します。
# mira.yaml: デプロイメント全体のデフォルト。すべてのキーはオプションです。
llm :
モデル：「アントロピック/クロード・ソネット-4-6」
fallback_model : " anthropic/claude-haiku-4-5 "
Indexing_model : " anthropic/claude-haiku-4-5 "
フィルター:
信頼性のしきい値 : 0.7
最大コメント数 : 5
レビュー：
ウォークスルー : true
自己批判 : true
security_pass : true
# .env: シークレットのみ。これをコミットしないでください。
MIRA_GITHUB_APP_ID=123456
MIRA_GITHUB_PRIVATE_KEY= " $( cat private-key.pem ) "
MIRA_WEBHOOK_SECRET=あなたの秘密
OPENROUTER_API_KEY=sk-or-...
docker run -p 8000:8000 \
--env-ファイル .env \
-v " $( pwd ) /mira.yaml:/app/mira.yaml " \
ghcr.io/miracodeai/mira:最新 \
--config /app/mira.yaml
Mira は内部で OpenRouter と通信するため、OpenRouter がサポートするすべてのモデルが機能します。 mira.yaml の llm.model は、OpenRouter モデル ページにリストされているものです。例:
OPENROUTER_API_KEY を 1 回設定します。 1 つのキーがすべてのプロバイダーで機能します。 Mira が認識するモデルの完全なレジストリ (価格と目的ごとの推奨事項を含む) については、src/mira/llm/models.json を参照してください。
独自のモデルを追加します。ダッシュボードのモデル ドロップダウン (およびその背後にある検証) は、そのレジストリから取得されます。カスタム モデルを選択可能にするには — 例: DeepSeek またはローカル エンドポイント — 独自の models.json で MIRA_MODELS_JSON_PATH を指定します (ボリューム マウントは適切に機能します)。そのエントリはモデル ID ごとにバンドルされたエントリにマージされるため、追加またはオーバーライドするもののみをリストします。
MIRA_MODELS_JSON_PATH=/config/models.json
// /config/models.json — バンドルされたファイルからエントリをテンプレートとしてコピーします
{
"ディープシーク/ディープシークチャット" : {
"label" : " DeepSeek チャット " ,
"プロバイダー" : " openai " ,
"max_input_tokens" : 64000 、
"max_output_tokens" : 8000 、
"1m あたりの入力コスト" : 0.27 、
"1m あたりの出力コスト" : 1.10 、
"supports_json_mode" : true 、
「目的

ses" : [ " インデックス作成 " 、 " レビュー " ]
}
}
ファイルは起動時に読み取られます。欠落しているファイルまたは無効なファイルは、警告とともにバンドルされたレジストリにフォールバックします。
近日提供予定: Anthropic 、 OpenAI 、 Google Vertex 、 Ollama 、および vLLM 用の直接アダプター。すでにプロバイダー キーを保持しているチーム、社内でオープンウェイト モデルを実行しているチーム、または OpenRouter を介したトラフィックの流れを防ぐデータ常駐ルールがあるチーム向け。
データ常駐要件または既存の AWS Bedrock アクセスがあるチームの場合、Mira は OpenRouter や仲介を介さずに、Bedrock の Converse API を直接呼び出すことができます。
#ミラ.yaml
llm :
提供者：「岩盤」
モデル: " us.anthropic.claude-sonnet-4-6-v1:0 "
fallback_model : " us.anthropic.claude-haiku-4-5-v1:0 "
リージョン : " us-east-1 "
# aws_profile: "my-profile" # オプション
認証では、標準の AWS 認証情報チェーン (環境変数、インスタンス プロファイル、ECS タスク ロール、SSO) を使用します。 API キーを管理する必要はありません。
git diff HEAD~1 | mira レビュー --stdin --dry-run --config mira.yaml
3. リポジトリにアプリをインストールします。すべての PR は自動レビューされます。
Mira とチャット: コードについて質問するには、PR に @miracodeai <question> をコメントしてください。
mira.yaml ( --config 経由でロード) には、モデルの選択、フィルターのしきい値、レビュー動作など、デプロイメント全体のデフォルトが保持されます。ほとんどのチームはここで止まります。
単一のリポジトリで異なるレビュー動作が必要な場合は、ルートに .mira.yaml をドロップするか、ダッシュボードからそのリポジトリの設定を編集します。どちらも、その 1 つのリポジトリのみに対して mira.yaml を介してディープ マージされます。
# .mira.yaml: オプションのリポジトリごとのオーバーライド
フィルター:
confidence_threshold : 0.5 # ノイズの多いリポジトリ → 下のバー
最大コメント数 : 10
完全なスキーマについてはドキュメントを参照してください。
git clone https://github.com/mira-reviewer/mira.git
CDミラ
pip install -e " .[dev,serve] "
# テストを実行する
pytest テスト/ -v
# 回帰スイートを実行します (実際の GitHub + LLM にヒット、~$1、~3 分)。
結果がちらついたピン留めされた PR の数

繰り返しにわたって。前に実行
# プロンプト、ノイズ フィルター、またはエンジンに関わる変更をマージします。
OPENROUTER_API_KEY=... GITHUB_TOKEN=... pytest -m eval -v
#糸くず
ruff チェック src/tests/
# 型チェック
mypy src/mira/ --ignore-missing-imports
ライセンス
インデックス付き PR レビュー、ウォークスルー、脆弱性スキャン、依存関係グラフ、カスタム ルール、学習ループを備えた自己ホスト型 AI コード レビューアー。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
13
フォーク
レポートリポジトリ
リリース
28
v0.4.0
最新の
2026 年 6 月 14 日
+ 27 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI code reviewer with indexed PR reviews, walkthroughs, vulnerability scanning, dependency graphs, custom rules, and a learning loop. - miracodeai/mira

Hey HN, I'm Jay, co-creator of Mira. An open-source, self-hosted AI code reviewer where you BYOK (bring your own key). Local models are getting really good. Hosted frontier models are getting really expensive. So, we built Mira. Mira has some snazzy features too: - It's really fast at reviewing. Average is 77s compared with Greptile's 5 minutes. Your PRs aren't going into a queue on a cloud somewhere. - Mira performs a blast radius and see what damage will be done with code change. - Learns your codebase's patterns from the repo itself and enforces them without a config file. - It's self-hosted, so your code never leaves your infrastructure. Bring your own model - OpenAI, Anthropic, or a local LLM behind your firewall. We're really active, and looking for active contributors. Come join us in our discord! https://github.com/miracodeai/mira Thanks,
Jay

GitHub - miracodeai/mira: Self-hosted AI code reviewer with indexed PR reviews, walkthroughs, vulnerability scanning, dependency graphs, custom rules, and a learning loop. · GitHub
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
miracodeai
/
mira
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
336 Commits 336 Commits .github .github scripts scripts src/ mira src/ mira tests tests ui/ mira ui/ mira .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .mira.yaml.example .mira.yaml.example .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FEATURES.md FEATURES.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md fly.toml fly.toml pyproject.toml pyproject.toml railway.toml railway.toml render.yaml render.yaml uv.lock uv.lock View all files Repository files navigation
Self-hosted AI code review. Your code, your dashboard, your LLM key.
Self-host every feature: full review engine, codebase indexing, vulnerability scanning, custom rules, org-wide package search, dashboard, learning loop. No paid tier, no license key, no SaaS upsell.
Mira reviews your pull requests using your choice of LLM (via OpenRouter , which fronts Anthropic, OpenAI, Google, DeepSeek, and more) and posts concise, actionable feedback. The noise filter, confidence clamping, and learning loop ensure you only see comments that matter. See FEATURES.md for the full surface.
Most AI reviewers are SaaS: your diffs (and often the full surrounding code) leave for a third-party server, and the only "view" you get is the comments that come back on a PR. Mira flips both halves of that:
Your code never leaves your infra. Diffs, embeddings, indexes, review history, vulnerability data, all stored in your SQLite or Postgres, on infrastructure you own. No phone-home, no required telemetry, no "is this used for training?" question.
The dashboard you see above is yours. It's not a marketing screenshot of someone else's view of your code. CodeRabbit, Greptile, and similar SaaS reviewers don't expose anything like it. Mira's dashboard surfaces signals you don't get anywhere else:
Org-wide package inventory : answer "which repos use lodash@4.17.20 ?" in one query. Stack it next to your CVE feed for instant blast-radius checks.
CVE alerts on every dependency : hourly OSV.dev poll, severity + advisory link + fix version surfaced inline next to the package.
Dependency + blast-radius graphs : see exactly which files and repos depend on a symbol before you change it.
Per-repo review event stream : every webhook, every chunk, every cost figure, in one place for live troubleshooting.
Cost & token telemetry : actual spend per repo and per model, not estimates, because you control the LLM key.
Coming soon, change-frequency heatmaps : surface the files that bug fixes keep landing on so you can target review attention.
If your engineering team needs answers like "which of our repos are exposed to this CVE?" or "what's the blast radius of changing this function?" , those questions stop being multi-day investigations and start being one-click dashboard pages.
Any OpenAI-compatible endpoint : OpenRouter by default (Anthropic, OpenAI, Google Gemini, DeepSeek, …), or point llm.base_url at vLLM, Ollama, LiteLLM proxy, LocalAI, llama.cpp, Together, Fireworks, Groq, or Cerebras for self-hosted / data-resident reviews. See Custom endpoints .
Low noise : Confidence thresholds, dedupe, severity sorting, per-PR comment caps
Indexed reviews : full-repo code index gives the LLM real context, not just the diff
Learns your team : synthesizes rules from rejected comments and human review patterns on merged PRs
Vulnerability scanning : hourly OSV.dev poll surfaces CVEs across every package in every repo
Org-wide package search : answer "which repos use lodash@4.17.20 ?" in seconds
Configurable : .mira.yaml for per-repo settings, custom + global rules in the dashboard
Self-host on day one : Docker image, Railway / Fly.io / Render configs, SQLite or Postgres
GitHub today, more platforms coming : GitLab, Bitbucket, and Gitea support are on the roadmap. Same review engine, same dashboard, just a different provider adapter
Mira is the fastest tool measured on the public Code Review Bench , and the only one on the speed/quality Pareto frontier: every tool that scores higher on F1 takes 5–14× longer per PR .
Plotted against every published competitor on the same subset, Mira sits in the upper-left corner: everything to the right is slower; everything above it pays 5–14× the wall time for the extra F1.
Measured on the same 50-PR offline benchmark, judged by Claude Sonnet 4.5.
Methodology: scores measured against the Martian Code Review Bench offline dataset with Claude Sonnet 4.5 as the judge.
Run Mira as a GitHub App that auto-reviews every PR and responds to comments.
Note: GitHub is the only supported provider today. GitLab, Bitbucket, and Gitea adapters are on the roadmap. The review engine, indexer, and dashboard are provider-agnostic; only the webhook + comment-posting layer is GitHub-specific. Star the repo to follow along.
1. Create a GitHub App at github.com/settings/apps/new :
Webhook URL: https://your-server.com/github/webhook
Permissions: Pull Requests (read+write), Contents (read+write), Issues (read+write)
Events: Pull requests, Issue comments
Two files: mira.yaml for non-secret defaults, .env for secrets. Pass both into the container.
# mira.yaml: deployment-wide defaults. Every key is optional.
llm :
model : " anthropic/claude-sonnet-4-6 "
fallback_model : " anthropic/claude-haiku-4-5 "
indexing_model : " anthropic/claude-haiku-4-5 "
filter :
confidence_threshold : 0.7
max_comments : 5
review :
walkthrough : true
self_critique : true
security_pass : true
# .env: secrets only. Don't commit this.
MIRA_GITHUB_APP_ID=123456
MIRA_GITHUB_PRIVATE_KEY= " $( cat private-key.pem ) "
MIRA_WEBHOOK_SECRET=your-secret
OPENROUTER_API_KEY=sk-or-...
docker run -p 8000:8000 \
--env-file .env \
-v " $( pwd ) /mira.yaml:/app/mira.yaml " \
ghcr.io/miracodeai/mira:latest \
--config /app/mira.yaml
Mira talks to OpenRouter under the hood, so any model OpenRouter supports works. llm.model in mira.yaml is whatever the OpenRouter Models page lists. Examples:
Set OPENROUTER_API_KEY once; one key works across every provider. See src/mira/llm/models.json for the full registry of models Mira recognises (with pricing and per-purpose recommendations).
Adding your own models. The dashboard's model dropdowns (and the validation behind them) come from that registry. To make a custom model selectable — e.g. DeepSeek or a local endpoint — point MIRA_MODELS_JSON_PATH at your own models.json (a volume mount works well). Its entries are merged over the bundled ones by model id, so you only list what you want to add or override:
MIRA_MODELS_JSON_PATH=/config/models.json
// /config/models.json — copy an entry from the bundled file as a template
{
"deepseek/deepseek-chat" : {
"label" : " DeepSeek Chat " ,
"provider" : " openai " ,
"max_input_tokens" : 64000 ,
"max_output_tokens" : 8000 ,
"input_cost_per_1m" : 0.27 ,
"output_cost_per_1m" : 1.10 ,
"supports_json_mode" : true ,
"purposes" : [ " indexing " , " review " ]
}
}
The file is read at startup; a missing or invalid file falls back to the bundled registry with a warning.
Coming soon: direct adapters for Anthropic , OpenAI , Google Vertex , Ollama , and vLLM , for teams that already hold provider keys, run open-weights models in-house, or have data-residency rules that prevent traffic from flowing through OpenRouter.
For teams with data-residency requirements or existing AWS Bedrock access, Mira can call Bedrock's Converse API directly — no OpenRouter, no intermediary:
# mira.yaml
llm :
provider : " bedrock "
model : " us.anthropic.claude-sonnet-4-6-v1:0 "
fallback_model : " us.anthropic.claude-haiku-4-5-v1:0 "
region : " us-east-1 "
# aws_profile: "my-profile" # optional
Auth uses the standard AWS credential chain (env vars, instance profile, ECS task role, SSO). No API keys to manage.
git diff HEAD~1 | mira review --stdin --dry-run --config mira.yaml
3. Install the app on your repos. Every PR gets auto-reviewed.
Chat with Mira: Comment @miracodeai <question> on any PR to ask about the code.
mira.yaml (loaded via --config ) holds deployment-wide defaults: model choices, filter thresholds, review behaviour. Most teams stop here.
If a single repo needs different review behaviour, drop a .mira.yaml at its root or edit that repo's settings from the dashboard. Both deep-merge over mira.yaml for that one repo only:
# .mira.yaml: optional per-repo override
filter :
confidence_threshold : 0.5 # noisier repo → lower bar
max_comments : 10
See the docs for the full schema.
git clone https://github.com/mira-reviewer/mira.git
cd mira
pip install -e " .[dev,serve] "
# Run tests
pytest tests/ -v
# Run the regression suite (hits real GitHub + LLM, ~$1, ~3 min).
# Pinned PRs whose findings have flickered across iterations. Run before
# merging changes that touch prompts, the noise filter, or the engine.
OPENROUTER_API_KEY=... GITHUB_TOKEN=... pytest -m eval -v
# Lint
ruff check src/ tests/
# Type check
mypy src/mira/ --ignore-missing-imports
License
Self-hosted AI code reviewer with indexed PR reviews, walkthroughs, vulnerability scanning, dependency graphs, custom rules, and a learning loop.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
13
forks
Report repository
Releases
28
v0.4.0
Latest
Jun 14, 2026
+ 27 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
