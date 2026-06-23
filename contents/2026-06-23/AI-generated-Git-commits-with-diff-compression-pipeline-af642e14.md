---
source: "https://github.com/benS-03/commait"
hn_url: "https://news.ycombinator.com/item?id=48648201"
title: "AI-generated Git commits with diff compression pipeline"
article_title: "GitHub - benS-03/commait: AI-powered Git CLI that generates commit messages, compress diffs, commits changes, and optionally pushes—all from a single command. · GitHub"
author: "dogbless"
captured_at: "2026-06-23T17:34:34Z"
capture_tool: "hn-digest"
hn_id: 48648201
score: 1
comments: 1
posted_at: "2026-06-23T17:20:02Z"
tags:
  - hacker-news
  - translated
---

# AI-generated Git commits with diff compression pipeline

- HN: [48648201](https://news.ycombinator.com/item?id=48648201)
- Source: [github.com](https://github.com/benS-03/commait)
- Score: 1
- Comments: 1
- Posted: 2026-06-23T17:20:02Z

## Translation

タイトル: 差分圧縮パイプラインを使用した AI 生成の Git コミット
記事のタイトル: GitHub - benS-03/commait: AI を活用した Git CLI は、コミット メッセージの生成、差分の圧縮、変更のコミット、およびオプションでプッシュをすべて 1 つのコマンドから実行します。 · GitHub
説明: AI を活用した Git CLI は、コミット メッセージの生成、差分の圧縮、変更のコミット、およびオプションでプッシュをすべて 1 つのコマンドから実行します。 - ベンS-03/コミット

記事本文:
GitHub - benS-03/commait: AI を活用した Git CLI で、コミット メッセージの生成、差分の圧縮、変更のコミット、およびオプションでプッシュをすべて 1 つのコマンドから実行できます。 · GitHub
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
ベンS-03
/
妥協する
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
68 コミット 68 コミット デモ デモ dist dist src src テスト テスト .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore README.md README.md commait-1.0.0.tgz commait-1.0.0.tgz package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI によって生成されたコミット メッセージ。トークンと時間を節約するために構築されています。
適切なコミット メッセージを書くのは難しくありませんが、AI がより速く、より適切に実行できるのであれば、なぜそうする必要があるのでしょうか。 Commait は段階的な diff を読み取り、数秒で書き込みます。差分トークンを削減するための圧縮アルゴリズムが含まれています。
$npm commait-cli をインストールします
クイックスタート
.env ファイルに OpenAI または Anthropic の API キーがあることを確認してください。標準の命名規則を使用します (例: ANTHROPIC_API_KEY)
$commait 構成の初期化
設定を好みに合わせて初期化します。
注: 構成の初期化は、構成ファイルが存在しない場合に自動的に行われるため、必要ありません。
段階的な変更をコミットするには、次を実行します。
$commaitコミット
これにより、メッセージの生成、コミット、リモートへのプッシュがすべて 1 つのコマンドで行われます。オプションの入力を求めるプロンプトが表示されるか、設定に基づいてプロンプトをスキップします。
ワンコマンド ワークフロー: 1 つのコマンドで変更を生成、コミット、プッシュします。
マルチプロバイダーのサポート: OpenAI または Anthropic を使用し、簡単な構成変更でメッセージを生成します。
完全に設定可能な自動化 : 自動ステージ変更、生成されたメッセージの自動コミット、リモートへの自動プッシュ、またはプロンプトとのすべてのステップのインタラクティブな維持など、ハンズオフを希望どおりにダイヤルインできます。
トークンバジェット制御: リクエストごとに送信される差分トークンの最大値を制限するため、LA であっても生成コストが予測可能になります。

ge の差分。
差分圧縮パイプライン : トークンのバジェットが生の差分を超えた場合、カスタムのステージングされた圧縮パイプラインが実行され、各ステージでトークンをバジェットと照合して、バジェットを下回りながらも情報損失を可能な限り削減します。詳細については、圧縮セクションを参照してください。
カスタム プロンプト : チームのコミット規則や個人のスタイルに合わせて、デフォルトのプロンプト テンプレートをオーバーライドします。
オプションのコンテキスト: コミット時に -c フラグを使用して、プロンプトにカスタム コンテキスト行をプレフィックスとして付けることができます。
インタラクティブなモデル ピッカー : 正確なモデル名の文字列を記憶する必要がなく、選択したプロバイダー内のモデルを参照して選択します。
構成可能なリモート: デフォルトのプッシュ ターゲットを設定するか、コミットごとにどのリモートを使用するかを尋ねるプロンプトが表示されます。
すべての構成オプションを 1 つずつ編集して実行するには、次のコマンドを実行します。
$commait 構成の初期化
または、単一のオプションを編集するには、次のコマンドを実行します。
$commait 構成セット [オプション]
オプションの完全なリストは以下で確認するか、次のコマンドを実行して取得できます。
$commait 構成オプション
設定オプション
Provider : 生成に使用する AI プロバイダー。
model : 選択したプロバイダー内で使用するモデル。
プロンプト : プロバイダーにプロンプ​​トを表示する内容。デフォルトのプロンプトが提供されます。
auto_stage : プロンプトを表示せずにすべてのファイルを自動的にステージングするようコミットするかどうか。
auto_commit : 生成後にプロンプ​​トを表示せずに自動的にコミットするかどうか。
auto_push : コミット後にプロンプ​​トを表示せずに自動的にプッシュするかどうか。
max_diff_tokens : AI に送信されるトークンの最大量。
default_origin : プッシュ先のデフォルトのリモート。
ask_origin : コミット時にどのリモートにプッシュするかを尋ねられるかどうか。
圧縮パイプラインは段階的に動作し、それぞれの動作がますます攻撃的になります。各ステージの実行後、新しいトークン数が予算に対してチェックされます。

不必要な情報損失が発生しないようにします。 4つの段階は次のとおりです。
Strip Noise Files : パッケージロック、gemfile、dist ファイルなどのユーザー生成以外のファイルを削除します。
Strip Headers : これは、差分内の各ファイルからヘッダーを取り除き、ファイル名と、ファイルが削除されたか名前が変更されたかを AI が伝えるインジケーターのみに置き換えます。
Strip Context : 各ハンクの前後にある 3 つの変更されていないコンテキスト行が削除されます。
Strip Lines : これは最終段階であり、最も攻撃的です。これは、各ハンクの最初と最後に一定数の行を保持し、残りを削除することで機能します。これによってトークン数がバジェットを下回らない場合は、このアクションがループされ、各反復で保持される行数が減ります。これは、諦めるまでに最大 6 回実行できます。
あなたは、Git コミット メッセージを作成する専門のソフトウェア エンジニアです。
提供された diff を分析し、次のルールに従ってコミット メッセージを書き込みます。
1行目: 従来のコミット>フォーマット — (): 概要>
タイプ: feat、fix、リファクタリング、chore、>docs、test、style、perf、ci
範囲: モジュール、ファイル、または機能 > 影響を受ける領域 (不明瞭な場合は省略)
概要: 命令形、小文字、ピリオドなし、最大 72 文字
必要に応じて、空白行を追加し、 > 短い本文 (最大 2 ～ 4 文) > 内容ではなく、なぜを説明します
差分がどのようなものかを決して説明しないでください。変更によって何が起こるかを説明してください。
曖昧さよりも具体性を優先 >(「認証を改善する」ではなく「OAuth トークンの更新に再試行ロジックを追加する」)
複数の懸念事項が変更された場合は、その主題に対して最も重要な懸念事項を選択します。他の人を本文に簡単にリストします
意味のある単位（構成ファイルなど）でない限り、ファイル名には言及しないでください。
バッククォートで囲んだり、説明を追加したりしないでください。コミット メッセージのみを出力します。
コミットメッサを生成する AI を活用した Git CLI

ges、差分の圧縮、変更のコミット、およびオプションでプッシュをすべて 1 つのコマンドから実行できます。
www.npmjs.com/package/commait-cli
トピックス
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

AI-powered Git CLI that generates commit messages, compress diffs, commits changes, and optionally pushes—all from a single command. - benS-03/commait

GitHub - benS-03/commait: AI-powered Git CLI that generates commit messages, compress diffs, commits changes, and optionally pushes—all from a single command. · GitHub
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
benS-03
/
commait
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
68 Commits 68 Commits demo demo dist dist src src tests tests .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore README.md README.md commait-1.0.0.tgz commait-1.0.0.tgz package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json View all files Repository files navigation
AI-generated commit messages, built to save on tokens and time.
Writing good commit messages is not hard, but why do it if AI can do it faster and better. Commait reads a staged diff and writes one for you in seconds. Includes a compression algorithm to reduce diff tokens.
$npm install commait-cli
Quick Start
Ensure you have an API key for either OpenAI or Anthropic in your .env file. Use standard naming conventions (ex. ANTHROPIC_API_KEY)
$commait config init
to initialize the configurations to your preferences.
Note: initializing config is not necessary as it will be done automatically if no config file exists.
To commit staged changes, run:
$commait commit
This will generate a message, commit, and push to remote all in one command. It will either prompt you with options or skip prompts based on config.
One-command workflow : Generate, commit, and push changes with a single command.
Multi-provider support : Generate messages using OpenAI or Anthropic with a simple config change.
Fully configurable automation : Dial in exactly how hands-off you want it: auto-stage changes, auto-commit the generated message, auto-push to your remote, or keep every step interactive with prompts.
Token budget control : Cap the max diff tokens sent per request, so generation cost stays predictable even on large diffs.
Diff compression pipeline : If token budget exceeded by raw diff, a custom staged compression pipeline runs, checking the tokens against the budget at each stage to reduce information loss as much as possible while still getting below budget. See compression section for more details.
Custom prompts : Override the default prompt template to match your team's commit conventions or personal style.
Optional Context : Prompt can be prefixed with custom context line using -c flag with committing.
Interactive model picker : Browse and select models within your chosen provider without needing to memorize exact model name strings.
Configurable remotes : Set a default push target, or get prompted for which remote to use on every commit.
To run through edit all config options one by one run:
$commait config init
Or to edit a single option run:
$commait config set [option]
A full list of options can be seen below or gotten by running:
$commait config options
Config Options
provider : Which AI provider to use for generation.
model : which model to use within selected provider.
prompt : What to prompt provider with, a default prompt is provided.
auto_stage : whether or not you want commit to automatically stage all files without prompting.
auto_commit : whether or not you want to automatically commit after generation without prompting.
auto_push : whether or not you want to automatically push after commit without prompting.
max_diff_tokens : the maximum amount of tokens sent to AI.
default_origin : the default remote to push to.
ask_origin : whether or not you want to be asked what remote to push to upon commit.
The compression pipeline works in stages, each one being increasingly more aggressive. After each stage runs, the new token count is checked against the budget so that unnecessary information loss does not occur. The 4 stages are as follows:
Strip Noise Files : This removes any non user generated files, like package-lock, gemfiles, dist files, etc...
Strip Headers : This will strip the header off of each file in the diff, and replace them with just the file name, and an indicator for the AI to convey if the file was deleted or renamed.
Strip Context : This removes the 3 unchanged context lines before and after every hunk.
Strip Lines : This is the final stage and most aggressive. This works by keeping a certain number of lines at the beginning and end of every hunk and stripping the rest. If this does not reduce the token count below the budget, it will loop this action, reducing the number of lines kept each iteration. This can run a max of 6 times, before it gives up.
You are an expert software engineer writing a Git commit message.
Analyze the provided diff and write a commit message following these rules:
First line: conventional commit >format — (): summary>
Types: feat, fix, refactor, chore, >docs, test, style, perf, ci
Scope: the module, file, or feature >area affected (omit if unclear)
Summary: imperative mood, lowercase, >no period, max 72 chars
If needed, add a blank line then a >short body (2–4 sentences max) >explaining WHY, not what
Never describe what the diff looks >like — describe what the change does
Prefer specificity over vagueness >("add retry logic to OAuth token refresh" not "improve auth")
If multiple concerns are changed, pick the dominant one for the subject; list others briefly in the body
Do not mention file names unless they are the meaningful unit (e.g. a config file)
Do not wrap in backticks or add any explanation — output the commit message only
AI-powered Git CLI that generates commit messages, compress diffs, commits changes, and optionally pushes—all from a single command.
www.npmjs.com/package/commait-cli
Topics
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
