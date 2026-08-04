---
source: "https://github.com/maxrodrigo/clai"
hn_url: "https://news.ycombinator.com/item?id=49170538"
title: "Show HN: Clai – AI for the command line (stdin → LLM → stdout)"
article_title: "GitHub - maxrodrigo/clai: AI for the Command Line · GitHub"
author: "maxrodrigo"
captured_at: "2026-08-04T16:08:06Z"
capture_tool: "hn-digest"
hn_id: 49170538
score: 1
comments: 0
posted_at: "2026-08-04T15:42:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Clai – AI for the command line (stdin → LLM → stdout)

- HN: [49170538](https://news.ycombinator.com/item?id=49170538)
- Source: [github.com](https://github.com/maxrodrigo/clai)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:42:39Z

## Translation

タイトル: HN を表示: Clai – コマンド ラインの AI (stdin → LLM → stdout)
記事のタイトル: GitHub - maxrodrigo/clai: コマンド ライン用 AI · GitHub
説明: コマンドライン用の AI 。 GitHub でアカウントを作成して、maxrodrigo/clai の開発に貢献してください。
HN テキスト: ChatGPT が開始されて以来、私は LLM を自分のワークフローに統合しようと試みてきました。電子メール、PRS のコミット、または長いチャット スレッドさえも。
私の直感は、小さな自動化プラットフォームをコード化したいということでしたが、そのプラットフォーム自体の制約によって再び制限されてしまうことに気付きました。ゼロに 2 回戻り、長年使用してきたツールである UNIX パイプラインに戻りました。
私は llm 以降の mod (現在は非推奨) を使用し始めました。これらは素晴らしいツールですが、他のツールと組み合わせてより適切に構成できる、より無駄のないものが必要でした。 Clai が行うことは 1 つです。標準入力を取得し、それを選択したモデルに送信し、結果を出力します。 REPL も状態もありません。それが完了したら、終了するだけです。 git diff |クレイコミット
猫の記事.txt |クレイ要約 |輝き
ぺぺぺーす |クレイ・TLDR
カール -s example.com/article.html | clai -e 「3つの主要な概念を抽出する」
名前付きプロンプトのセットが付属していますが、必要に応じてオーバーライド、追加、拡張することができます。これらはフロントマターを含む単なるファイルです。ニーズに応じて推論を変更するための戦略 (ドラフトの連鎖、思考の連鎖、思考のツリー、自己洗練) もあります。主要なプロバイダーと連携しており、今後も追加していきます。ローカル モデルを指すこともでき、マシンから何も残らないようにすることもできます。インストール方法: brew install maxrodrigo/tap/clai
これは v0.3.0 で、数週間前のものです。早いし、荒削りだ。何が足りないのか、また役立つと思われる場合はフィードバックをいただければ幸いです。貢献は大歓迎です。

記事本文:
GitHub - maxrodrigo/clai: コマンドライン用 AI · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
マックスロドリゴ
/
クレイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .github .github cmd/ clai cmd/ clai docs docs 内部 内部共有/ clai share/ clai .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml CODE_OF_CONDUCT.md CODE_OF

_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md crime.toml 崖.toml go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
clai は言語モデルを通常のシェル コマンドに変換します。入力 (stdin) を読み取り、プロンプトの指示どおりに実行し、結果 (stdout) を出力します。他のすべて ( grep 、 jq 、 git 、 cat 、 pbpaste 、 xclip ) は、これまでと同様に動作します。
git diff |クレイコミット
猫の記事.txt |クレイ要約 |クレイ翻訳する
clai コードレビュー main.go | clai -e " 重大なバグのみ、番号付き "
1 つのことをうまく実行するプログラムを作成します。
連携して動作するプログラムを作成します。
テキスト ストリームを処理するプログラムを作成します。テキスト ストリームはユニバーサル インターフェイスであるためです。
その他のドキュメント:
例 ·
上級・
哲学
パイプラインネイティブ。標準入力、標準出力。 grep、jq、awk、その他すべてを使用して構成します。
あらゆるソースで動作します。 YouTube トランスクリプト、Web ページ、PDF、git diff、クリップボード。パイプで注入します。「例」を参照してください。
設定ゼロ。 API キーを 1 つ設定して実行します。セットアップ ウィザードや対話型プロンプトはありません。
組み込みのプロンプト。要約、コードレビュー、コミット、翻訳、説明など。
マルチプロバイダー。 OpenAI、Anthropic、Gemini、Vertex AI、Bedrock、Ollama、または任意の OpenAI 互換エンドポイント。
地元モデル。それを Ollama に向けても、マシンからは何も残されません。
推論戦略。思考の連鎖、思考のツリー、ドラフトの連鎖、自己洗練。
構造化された出力。専用の終了コードによる JSON スキーマ検証。
他にもターミナル AI ツールがあります: llm 、 mods 、 Fabric 、 aichat 。それらは優れており、そのほとんどはセッション、プラグイン、チャット モードへと成長し続けています。
クライは別の方法に賭けます。プラットフォームではなくツールです。ワークフローをキャプチャすることなく拡張します。 REPL もセッション状態もありません。コマンドが実行され、その r が出力されます

結果、終了します。ビヘイビアーはマークダウン ファイル内に存在し、開いて編集できます。 「哲学」を参照してください。
醸造インストール maxrodrigo/tap/clai
または Go を使用すると:
github.com/maxrodrigo/clai/cmd/clai@latest をインストールしてください
ビルド済みバイナリはリリース ページで入手できます。
注: go install はバイナリのみをインストールします。システム プロンプトと戦略については、次を実行します。
カール -sL https://github.com/maxrodrigo/clai/releases/latest/download/clai-data.tar.gz | tar -xz -C ~ /.local/share
# プロバイダーキーを設定する
エクスポート OPENAI_API_KEY= " sk-... "
# 名前付きプロンプト
claiまとめ記事.txt
clai コードレビュー main.go
# インラインプロンプト
clai -e 「わかりやすく説明します」 complex.txt
# パイプライン
git diff --cached |クレイコミット
カール -s https://api.example.com/data | clai -e " 異常を見つける "
構成
clai は TOML ファイルから設定を読み取ります。
~/.config/clai/config.toml # ユーザー設定
.clai/config.toml # プロジェクト構成 (ユーザーをオーバーライド)
例:
モデル = " openai/gpt-4.1 "
温度 = 1.0
max_tokens = 4096
[プロバイダー。オープンナイ]
api_key = " ${OPENAI_API_KEY} "
[プロバイダー。人間性 ]
api_key = " ${ANTHROPIC_API_KEY} "
利用可能なすべてのオプションとそのデフォルトを含む完全なサンプルについては、docs/ADVANCED.md – 構成リファレンスを参照してください。
すべての設定値は CLAI_ プレフィックスを介して設定できます。
export CLAI_MODEL= " 人間/クロード・ソネット "
エクスポート CLAI_TEMPERATURE= " 0.7 "
エクスポート CLAI_MAX_TOKENS= " 8192 "
プロバイダー API キーは、標準の環境変数を使用します。
エクスポート OPENAI_API_KEY= " sk-... "
import ANTHROPIC_API_KEY= " sk-ant-... "
エクスポート AWS_BEARER_TOKEN_BEDROCK= " ... "
優先順位
構成は次の順序でマージされます (後から前にオーバーライドされます)。
ユーザー設定 ( ~/.config/clai/config.toml )
プロジェクト構成 ( .clai/config.toml )
環境変数 ( CLAI_* )
モデルは、 Provider/model-name として指定されます。各プロバイダーには、e 経由で設定された API キーが必要です。

環境変数または設定ファイル。
APIキー: https://platform.openai.com/api-keys
エクスポート OPENAI_API_KEY= " sk-... "
clai summary -m openai/gpt-4.1article.txt
人間的
API キー: https://console.anthropic.com/settings/keys
import ANTHROPIC_API_KEY= " sk-ant-... "
clai summary -m anthropic/claude-sonnetarticle.txt
# 拡張された思考
clai -e " 根本原因を見つける " -m anthropic/claude-sonnet --think question.txt
AWS の基盤
ドキュメント: https://docs.aws.amazon.com/bedrock/latest/userguide/getting-started-api.html
エクスポート AWS_BEARER_TOKEN_BEDROCK= " ... "
clai summary -m bedrock/us.anthropic.claude-sonnetarticle.txt
設定 (リージョンを変更するため):
[プロバイダー。岩盤】
api_key = " ${AWS_BEARER_TOKEN_BEDROCK} "
Base_url = " https://bedrock-runtime.us-west-2.amazonaws.com "
ジェミニ (Google AI スタジオ)
API キー: https://aistudio.google.com/apikey
import GOOGLE_API_KEY= " ... "
clai summary -m gemini/gemini-2.5-flasharticle.txt
# 思考モード
clai -e " 根本原因を見つける " -m gemini/gemini-2.5-pro --think question.txt
頂点AI
ドキュメント: https://cloud.google.com/vertex-ai/generative-ai/docs/start/quickstarts/quickstart-multimodal
アプリケーションのデフォルト認証情報を使用します。 gcloud auth application-default ログインまたはサービス アカウントを使用して設定します。
import GOOGLE_CLOUD_PROJECT= " my-project-id "
import GOOGLE_CLOUD_LOCATION= " us-central1 "
clai summary -m vertex/gemini-2.5-flasharticle.txt
Config (環境変数の代替):
[プロバイダー。頂点]
プロジェクト = " 私のプロジェクト ID "
場所 = " us-central1 "
カスタムプロバイダー (Ollama、Groq など)
OpenAI 互換の API:
[プロバイダー。グロク]
api_key = " ${GROQ_API_KEY} "
Base_url = " https://api.groq.com/openai/v1 "
[プロバイダー。オラマ ]
Base_url = " http://localhost:11434/v1 "
clai 要約 -m groq/llama-4-scoutarticle.txt
Clai 要約 -m olla

ma/llama3.3 記事.txt
プロンプト
プロンプトは、モデルに何をすべきかを指示するマークダウン ファイルです。
# 名前付きプロンプト
claiまとめ記事.txt
clai コードレビュー code.patch
clai は、notes.md を翻訳します
# 複数のファイル
clai 要約レポート.txt メモ.txt
# -e を使用したインラインプロンプト
clai -e " このコードの説明 " main.go
# -f を使用してファイルからプロンプトを表示
clai -f my-prompt.mdarticle.txt
# 自然言語をシェルコマンドに変換
echo "今週変更された 10KB を超えるすべての Go ファイルを検索します" | clai シェル cmd
# -> を見つけます。 -name '*.go' -mtime -7 -size +10k
# 利用可能なプロンプトをリストする
クレイプロンプトリスト
クレイプロンプトショーの要約
~/.config/clai/prompts/ にファイルを作成します。
---
説明 : `clai プロンプト` の 1 行
モデル: anthropic/claude-sonnet # オプション
戦略 : コット # オプション
---
ここでの即時指示。
すぐに使用してください:
clai your-prompt input.txt
プロンプトは順番に解決されます (最初に一致したものが勝ちます)。
.clai/prompts/ (プロジェクトローカル)
~/.config/clai/prompts/ (ユーザー)
パターン CLAI_MODEL_<PROMPT_NAME> を使用して、環境変数経由でプロンプトのモデルをオーバーライドします。
export CLAI_MODEL_CODE_REVIEW= " 人間/クロード・ソネット "
import CLAI_MODEL_SUMMARIZE= " openai/gpt-4.1-mini "
プロンプトオーサリングの原則、構成、評価チェックリストについては、docs/ADVANCED.md を参照してください。
戦略は、モデルが問題を解決する方法を変更します。
clai Explain --strategy cot問題.txt
clai Explain --strategy none issue.txt # 無効にする
クレイ攻略リスト # 全リスト
~/.config/clai/strategies/ にカスタム戦略を作成します。研究の根拠、それぞれをいつ使用するか、およびカスタム戦略の作成については、docs/ADVANCED.md を参照してください。
--schema を使用して、スキーマに準拠した JSON 出力を取得します。
# 短縮構文
clai parse -s ' {"name": "str", "amount": "float"} ' invoice.txt
# 完全な JSON スキーマ
clai parse -s ' {"type": "object", "properties": {"items": {"type": "arr

ay"}}} ' data.txt
短縮型: str 、 int 、 float 、 bool 、 date 、 list 、 {"nested": "str"} 。
スキーマはプロンプトのフロントマターでも設定できます。
---
スキーマ:
タイトル : ストラ
著者：ストラ
タグ : リスト
---
この記事からメタデータを抽出します。
CLI リファレンス
clai [フラグ] <プロンプト> [ファイル...]
フラグ
旗
説明
-e、--式
インラインプロンプトテキスト
-f、--ファイル
ファイルからプロンプトを読み取ります
-m、--モデル
使用するモデル (例: openai/gpt-4.1 )
-t、--温度
サンプリング温度 (0.0 ～ 2.0)
--最大トークン
生成する最大トークン数
-s、--スキーマ
出力スキーマ (省略表現または JSON スキーマ)
--戦略
推論戦略 (cot、cod、tot、self-refine)
--考える
拡張的思考を可能にする (人間的/基盤的)
-n、--ドライラン
モデルを呼び出さずに送信される内容を表示する
-v、--verbose
トークン数とタイミングを標準エラー出力に出力します。
--色
標準出力が TTY でない場合でも、色付き出力を強制する
--色なし
カラー出力を無効にする
--バージョン
バージョンを印刷して終了する
コマンド
コマンド
説明
クレイプロンプトリスト
利用可能なプロンプトをリストする
clai プロンプト show <名前>
プロンプトコンテンツを表示
clai プロンプトのパス <名前>
プロンプトファイルの場所を表示
clai プロンプト add <名前>
新しいプロンプトを作成し、$EDITOR で開きます
clai プロンプト更新 <名前>
$EDITOR で既存のプロンプトを編集する
clai プロンプト <名前> を削除します
ユーザーがインストールしたプロンプトを削除する
clai プロンプト install <所有者/名前> <ファイル>
名前空間内のファイルからプロンプトをインストールする
クレイ攻略一覧
利用可能な戦略をリストする
クレイ戦略ショー<名前>
戦略コンテンツを表示
Clai 戦略パス <名前>
戦略ファイルの場所を表示
クライモデル一覧
利用可能なモデルをリストする
終了コード
コード
意味
0
成功
1
ランタイムエラー（APIエラー、ネットワークエラー）
2
使用法エラー (無効な引数、構成が欠落しています)
3
スキーマ検証エラー (出力がスキーマと一致しません)
シェルの完成
clai 補完 zsh --help
clai 補完 bash --help
クライ・コンプリ

魚のエション --ヘルプ
clai がワークフローに適合する場合は、リポジトリにスターを付けます。他の人がそれを見つけるのに役立ちます。
プロンプトは最も簡単なコントリビューションです。プロンプトはマークダウン ファイルであり、Go は必要ありません。 COTRIBUTING.md を参照してください。
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI for the Command Line . Contribute to maxrodrigo/clai development by creating an account on GitHub.

Since ChatGPT launched I've been trying to integrate LLMs into my workflows; emails, commits PRSs, or even long chat threads.
My instinct was to code a small automation platform, but I found myself limited once again by its own constraints. Went back to zero, twice, and went back to the tools I've been using for years: UNIX pipelines.
I started using llm and later mods (now deprecated) and they're great tools, but I wanted something leaner and that would compose better with other tools. Clai does one thing: takes stdin, sends it to a model of your choice and prints the result. No REPL, no state. When it's done it just exits. git diff | clai commit
cat article.txt | clai summarize | glow
pbpaste | clai tldr
curl -s example.com/article.html | clai -e "Extract the three main concepts"
It ships with a set of named prompts but you can override, add or extend as needed, they're just files with frontmatter. There's also strategies (chain-of-draft, chain-of-thought, tree-of-thought, and self-refine) to change the reasoning depending on your needs. Works with major providers and I'll keep adding more. It can also be pointed to a local model so nothing leaves your machine. Install with: brew install maxrodrigo/tap/clai
This is v0.3.0 and it's a few weeks old. It's early and rough around the edges. I'd appreciate feedback on what's missing and if you find it useful. Contributions welcome.

GitHub - maxrodrigo/clai: AI for the Command Line · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
maxrodrigo
/
clai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github .github cmd/ clai cmd/ clai docs docs internal internal share/ clai share/ clai .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md cliff.toml cliff.toml go.mod go.mod go.sum go.sum View all files Repository files navigation
clai turns a language model into a regular shell command. It reads input (stdin), does what the prompt says, prints the result (stdout). Everything else ( grep , jq , git , cat , pbpaste , xclip ) works with it the way it always has.
git diff | clai commit
cat article.txt | clai summarize | clai translate
clai code-review main.go | clai -e " Only critical bugs, numbered "
Write programs that do one thing and do it well.
Write programs to work together.
Write programs to handle text streams, because that is a universal interface.
More docs:
Examples ·
Advanced ·
Philosophy
Pipeline native. stdin in, stdout out. Composes with grep, jq, awk, and everything else.
Works with any source. YouTube transcripts, web pages, PDFs, git diffs, clipboard. Pipe it in. See Examples .
Zero config. Set one API key and go. No setup wizards, no interactive prompts.
Built-in prompts. summarize, code-review, commit, translate, explain, and more.
Multi-provider. OpenAI, Anthropic, Gemini, Vertex AI, Bedrock, Ollama, or any OpenAI-compatible endpoint.
Local models. Point it at Ollama and nothing leaves your machine.
Reasoning strategies. Chain-of-Thought, Tree-of-Thought, Chain-of-Draft, Self-Refine.
Structured output. JSON Schema validation with a dedicated exit code.
There are other terminal AI tools: llm , mods , fabric , aichat . They are good, and most of them keep growing into sessions, plugins, and chat modes.
clai bets the other way. A tool, not a platform. It amplifies your workflow without capturing it. There is no REPL and no session state; the command runs, prints its result, and exits. Behaviors live in markdown files you can open and edit. See Philosophy .
brew install maxrodrigo/tap/clai
Or with Go:
go install github.com/maxrodrigo/clai/cmd/clai@latest
Pre-built binaries available on the Releases page.
Note: go install installs the binary only. For system prompts and strategies, run:
curl -sL https://github.com/maxrodrigo/clai/releases/latest/download/clai-data.tar.gz | tar -xz -C ~ /.local/share
# Set a provider key
export OPENAI_API_KEY= " sk-... "
# Named prompts
clai summarize article.txt
clai code-review main.go
# Inline prompt
clai -e " Explain in simple terms " complex.txt
# Pipeline
git diff --cached | clai commit
curl -s https://api.example.com/data | clai -e " Find anomalies "
Configuration
clai reads configuration from TOML files:
~/.config/clai/config.toml # User config
.clai/config.toml # Project config (overrides user)
Example:
model = " openai/gpt-4.1 "
temperature = 1.0
max_tokens = 4096
[ providers . openai ]
api_key = " ${OPENAI_API_KEY} "
[ providers . anthropic ]
api_key = " ${ANTHROPIC_API_KEY} "
See docs/ADVANCED.md – Configuration Reference for a complete sample with all available options and their defaults.
All config values can be set via CLAI_ prefix:
export CLAI_MODEL= " anthropic/claude-sonnet "
export CLAI_TEMPERATURE= " 0.7 "
export CLAI_MAX_TOKENS= " 8192 "
Provider API keys use their standard environment variables:
export OPENAI_API_KEY= " sk-... "
export ANTHROPIC_API_KEY= " sk-ant-... "
export AWS_BEARER_TOKEN_BEDROCK= " ... "
Precedence
Configuration is merged in order (later overrides earlier):
User config ( ~/.config/clai/config.toml )
Project config ( .clai/config.toml )
Environment variables ( CLAI_* )
Models are specified as provider/model-name . Each provider requires an API key, set via environment variable or config file.
API Key: https://platform.openai.com/api-keys
export OPENAI_API_KEY= " sk-... "
clai summarize -m openai/gpt-4.1 article.txt
Anthropic
API Key: https://console.anthropic.com/settings/keys
export ANTHROPIC_API_KEY= " sk-ant-... "
clai summarize -m anthropic/claude-sonnet article.txt
# Extended thinking
clai -e " Find the root cause " -m anthropic/claude-sonnet --think problem.txt
AWS Bedrock
Docs: https://docs.aws.amazon.com/bedrock/latest/userguide/getting-started-api.html
export AWS_BEARER_TOKEN_BEDROCK= " ... "
clai summarize -m bedrock/us.anthropic.claude-sonnet article.txt
Config (to change region):
[ providers . bedrock ]
api_key = " ${AWS_BEARER_TOKEN_BEDROCK} "
base_url = " https://bedrock-runtime.us-west-2.amazonaws.com "
Gemini (Google AI Studio)
API Key: https://aistudio.google.com/apikey
export GOOGLE_API_KEY= " ... "
clai summarize -m gemini/gemini-2.5-flash article.txt
# Thinking mode
clai -e " Find the root cause " -m gemini/gemini-2.5-pro --think problem.txt
Vertex AI
Docs: https://cloud.google.com/vertex-ai/generative-ai/docs/start/quickstarts/quickstart-multimodal
Uses Application Default Credentials. Set up with gcloud auth application-default login or a service account.
export GOOGLE_CLOUD_PROJECT= " my-project-id "
export GOOGLE_CLOUD_LOCATION= " us-central1 "
clai summarize -m vertex/gemini-2.5-flash article.txt
Config (alternative to env vars):
[ providers . vertex ]
project = " my-project-id "
location = " us-central1 "
Custom Providers (Ollama, Groq, etc.)
Any OpenAI-compatible API:
[ providers . groq ]
api_key = " ${GROQ_API_KEY} "
base_url = " https://api.groq.com/openai/v1 "
[ providers . ollama ]
base_url = " http://localhost:11434/v1 "
clai summarize -m groq/llama-4-scout article.txt
clai summarize -m ollama/llama3.3 article.txt
Prompts
Prompts are markdown files that tell the model what to do.
# Named prompts
clai summarize article.txt
clai code-review code.patch
clai translate notes.md
# Multiple files
clai summarize report.txt notes.txt
# Inline prompt with -e
clai -e " Explain this code " main.go
# Prompt from file with -f
clai -f my-prompt.md article.txt
# Natural language to a shell command
echo " find all Go files modified this week over 10KB " | clai shell-cmd
# -> find . -name '*.go' -mtime -7 -size +10k
# List available prompts
clai prompt list
clai prompt show summarize
Create a file in ~/.config/clai/prompts/ :
---
description : One-line for `clai prompt`
model : anthropic/claude-sonnet # optional
strategy : cot # optional
---
Your prompt instructions here.
Use it immediately:
clai your-prompt input.txt
Prompts are resolved in order (first match wins):
.clai/prompts/ (project-local)
~/.config/clai/prompts/ (user)
Override a prompt's model via environment variable using the pattern CLAI_MODEL_<PROMPT_NAME> :
export CLAI_MODEL_CODE_REVIEW= " anthropic/claude-sonnet "
export CLAI_MODEL_SUMMARIZE= " openai/gpt-4.1-mini "
See docs/ADVANCED.md for prompt authoring principles, composition, and the evaluation checklist.
Strategies modify how the model reasons through problems.
clai explain --strategy cot problem.txt
clai explain --strategy none problem.txt # disable
clai strategy list # list all
Create custom strategies in ~/.config/clai/strategies/ . See docs/ADVANCED.md for research basis, when to use each, and custom strategy authoring.
Use --schema to get JSON output conforming to a schema.
# Shorthand syntax
clai parse -s ' {"name": "str", "amount": "float"} ' invoice.txt
# Full JSON Schema
clai parse -s ' {"type": "object", "properties": {"items": {"type": "array"}}} ' data.txt
Shorthand types: str , int , float , bool , date , list , {"nested": "str"} .
Schema can also be set in prompt frontmatter:
---
schema :
title : str
author : str
tags : list
---
Extract metadata from this article.
CLI Reference
clai [flags] <prompt> [files...]
Flags
Flag
Description
-e, --expression
Inline prompt text
-f, --file
Read prompt from file
-m, --model
Model to use (e.g., openai/gpt-4.1 )
-t, --temperature
Sampling temperature (0.0–2.0)
--max-tokens
Maximum tokens to generate
-s, --schema
Output schema (shorthand or JSON Schema)
--strategy
Reasoning strategy (cot, cod, tot, self-refine)
--think
Enable extended thinking (Anthropic/Bedrock)
-n, --dry-run
Show what would be sent without calling model
-v, --verbose
Print token counts and timing to stderr
--color
Force colored output even when stdout is not a TTY
--no-color
Disable colored output
--version
Print version and exit
Commands
Command
Description
clai prompt list
List available prompts
clai prompt show <name>
Show prompt content
clai prompt path <name>
Show prompt file location
clai prompt add <name>
Create a new prompt and open in $EDITOR
clai prompt update <name>
Edit an existing prompt in $EDITOR
clai prompt remove <name>
Remove a user-installed prompt
clai prompt install <owner/name> <file>
Install a prompt from a file under a namespace
clai strategy list
List available strategies
clai strategy show <name>
Show strategy content
clai strategy path <name>
Show strategy file location
clai model list
List available models
Exit Codes
Code
Meaning
0
Success
1
Runtime error (API failure, network error)
2
Usage error (invalid arguments, missing config)
3
Schema validation error (output doesn't match schema)
Shell Completion
clai completion zsh --help
clai completion bash --help
clai completion fish --help
If clai fits your workflow, star the repo . It helps others find it.
Prompts are the easiest contribution: they're markdown files, no Go required. See CONTRIBUTING.md .
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
