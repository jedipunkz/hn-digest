---
source: "https://github.com/theadamdanielsson/overllm"
hn_url: "https://news.ycombinator.com/item?id=48787752"
title: "Overllm – flags where you're paying an LLM to do a regex's job"
article_title: "GitHub - theadamdanielsson/overllm: Catch the LLM/AI calls you didn't need. A fast, deterministic linter that flags LLM API calls where plain code is simpler, cheaper, and more reliable. Your GPT call is a regex. · GitHub"
author: "adamdanielsson1"
captured_at: "2026-07-04T19:01:13Z"
capture_tool: "hn-digest"
hn_id: 48787752
score: 1
comments: 1
posted_at: "2026-07-04T18:45:16Z"
tags:
  - hacker-news
  - translated
---

# Overllm – flags where you're paying an LLM to do a regex's job

- HN: [48787752](https://news.ycombinator.com/item?id=48787752)
- Source: [github.com](https://github.com/theadamdanielsson/overllm)
- Score: 1
- Comments: 1
- Posted: 2026-07-04T18:45:16Z

## Translation

タイトル: Overllm – 正規表現の仕事をするために LLM にお金を払っていることを示すフラグ
記事のタイトル: GitHub - theadamdanielsson/overllm: 不要な LLM/AI 呼び出しをキャッチします。 LLM API 呼び出しにフラグを立てる、高速で決定論的なリンター。プレーン コードの方がシンプルで、安価で、信頼性が高くなります。 GPT 呼び出しは正規表現です。 · GitHub
説明: 不要な LLM/AI 呼び出しをキャッチします。 LLM API 呼び出しにフラグを立てる、高速で決定論的なリンター。プレーン コードの方がシンプルで、安価で、信頼性が高くなります。 GPT 呼び出しは正規表現です。 - テアダムダニエルソン/overllm

記事本文:
GitHub - theadamdanielsson/overllm: 不要な LLM/AI 呼び出しをキャッチします。 LLM API 呼び出しにフラグを立てる、高速で決定論的なリンター。プレーン コードの方がシンプルで、安価で、信頼性が高くなります。 GPT 呼び出しは正規表現です。 · GitHub
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
テアダムダニエルソン

/
全体的に
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows サンプル サンプル スクリプト スクリプト スキル/ overllm スキル/ overllm src/ overllm src/ overllm テスト テスト .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
不要な LLM/AI 呼び出しをキャッチします。
overllm は、1 つのジョブを備えた小型で高速なリンターです。コード内で AI モデルを呼び出して、プレーン コードの方が適切な処理を行う場所を見つけます。日付を解析するために GPT を呼び出しました。 json.loads がすでに処理している JSON を抽出するモデルを呼び出しました。正規表現に対してレイテンシ、コスト、および非決定性を支払っていることになります。
実際のパーサーを使用してコードを読み取ります。Python は標準ライブラリ ast を使用し、JavaScript と TypeScript はツリーシッターを使用します。モデルは実行されず、ネットワークも API キーもありません。同じコードが入力され、同じ結果が出力されます。コミット前のフックとしては十分な速度です。
他の人は AI が書いたコードをリントします。 overllm は、ライブラリがすでに行っていることを行うために AI にお金を払っているところをキャッチします。
pip install overllm # Python
pip install " overllm[js] " # JavaScript / TypeScript サポートを追加します
使ってください
最初に 1 つのファイルを指定すると、すべてのファイルに対して実行する前に、何がフラグを立てているかを確認できます。
overllm app.py # 1 つのファイル
overllm src/ # フォルダー
全体的に。 # プロジェクト全体
コードを読み取り、見つけたものを出力します。何も書き込まず、何も変更しません —

これを実行すると、最悪の場合でも数行の出力が発生します。
app.py:42:5 llm-mechanical LLM 呼び出しはモデルにソートを要求します
resp = client.chat.completions.create(model="gpt-4o",messages=[...])
->sorted()を使用する
app.py:88:1 llm-in-loop ループ内の LLM 呼び出し: 反復ごとに 1 つの API ラウンドトリップ
complete(model="gpt-4o",messages=[{"role": "user", "content": f"tag {x}"}])
-> 入力を 1 回の呼び出しにバッチ処理する、繰り返しの結果をキャッシュする、または関数を使用する
1 つのファイル内に 2 つの不必要な LLM 呼び出しがあります。
デフォルトでは静かです。警告とエラーの結果のみが表示されるため、クリーンなプロジェクトでは何も出力されず、0 で終了します。ほとんどのコードベースでは、少数かまったく表示されません。洪水が発生した場合は、それをバグとして扱い、問題を報告してください。
overllm は何かを見つけるとゼロ以外で終了するため、コミットまたは CI チェックをゲートします。失敗せずに --exit-zero を渡してレポートします。
コードはマシン上に残ります
overllm は静的解析です。ファイルをローカルで解析し、見つかった内容を出力します。コードをアップロードしたり、API を呼び出したり、キーを必要とせず、テレメトリを送信することはありません。ループ内にモデルはなく、電話で応答するものはありません。ネットワーク ケーブルを引っ張ると、まったく同じように動作します。
コアは数百行の Python であり、必要な依存関係がないため、信頼する前にすべてを読むことができます。 overllm[js] は、JavaScript と TypeScript を解析するツリーシッターを追加します。これが唯一のオプションの依存関係です。
すべてのルールは具体的なコード パターンに対してのみ適用され、すべての検出結果は決定的な置換に名前を付けます。確信が持てないときは沈黙する。
デフォルトでは、overllm は警告以上を生成するだけなので、日常のコードでは静かになります。 static-prompt は info であり、 --all または --min-severity info で要求しない限り、何も表示されません。
最後の 2 つは、プロンプトではなく呼び出し自体をチェックします。つまり、存在しないモデル ID、またはモデルが無視するノブです。両方

既知のリストと正確に一致するため、ライブ モデルまたはエイリアスにフラグが立てられることはありません。リストはプロバイダーの非推奨ページを追跡するため、時間の経過とともに更新する必要があります。
Python の OpenAI、Anthropic、Google、Mistral、Cohere、Groq、AWS Bedrock、HuggingFace、Replicate、LangChain、LiteLLM、および Ollama SDK、Vercel AI SDK (generateText、streamText、generateObject)、JavaScript および TypeScript の openai / anthropic ノード SDK、およびこれらのホストへの生の HTTP リクエストを検出します。
resp = クライアント 。チャット 。完成品。 create (...) # overllm: 無視
resp = クライアント 。チャット 。完成品。 create (...) # overllm:ignore=llm-in-loop
ファイル全体をスキップするには、ファイルの先頭に # overllm:ignore-file を置きます。
pyproject.toml (Python 3.11 以降):
「ツール。全体]
無視 = [ " llm-in-loop " ]
exclude = [ " 例 / " 、 " 移行 / " ]
またはコマンドラインで: --select 、 --ignore 、 --min-severity 、 --all 、および --config PATH ( exclude は構成のみです)。完全なリストを表示するには、overllm --help を実行します。
リポジトリ:
- リポジトリ: https://github.com/theadamdanielsson/overllm
リビジョン : v0.4.0
フック:
- ID : overllm
GitHub アクション
overllm には、プル リクエストをスキャンし、根拠のあるコメントを 1 つ残すアクションが同梱されています。何も言うことがないときは沈黙します。
名前 : オーバーレルム
に:
プルリクエスト:
権限:
内容：読む
プルリクエスト : 書き込み
ジョブ:
チェック:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用: theadamdanielsson/overllm@v1
付き:
パス: " . "
その他の出力形式
overllm --format json 。 # 機械可読
overllm --format sarif 。 # GitHub コードスキャンにアップロード
overllm --format マークダウン。 # PR コメント本文
AI コードレビューアーを使用しないのはなぜでしょうか?
AI レビュー担当者と AI スロップ リンターは、モデルが生成したコード (コメント、デッド コード、構造) を調べます。彼らの誰も、llm が尋ねるような質問をしません。

ch は、そのモデルが本当に必要かどうかです。これは別の軸であり、高精度かつゼロコストで回答できる単純な静的分析です。
オーバーループ、実行中のエージェント用
overllm は保存中のコードを読み取ります。そのランタイム兄弟である overloop は、エージェントの実行中に同じ無駄を捕捉するクロード コード フックです。ツールは繰り返し呼び出し、ファイルを再読み込みし、コンテキストをあふれさせる特大の出力を呼び出します。 overllm は必要のない呼び出しをキャッチします。オーバーループは、エージェントが 2 回実行するものをキャッチします。同じアイデアの 2 つの半分。
送信できる最も便利なものは、誤検知です。つまり、overllm が、すべきではない呼び出しにフラグを立てている実際のコード行です。リンターは、それが正しい場合にのみ実行する価値があるため、1 つの具体的な不良フラグは機能リクエストよりも価値があります。 COTRIBUTING.md には、レポートの作成方法とテストの実行方法が記載されています。
過去のリリースと変更内容は CHANGELOG.md にあります。
不要な LLM/AI 呼び出しをキャッチします。 LLM API 呼び出しにフラグを立てる、高速で決定論的なリンター。プレーン コードの方がシンプルで、安価で、信頼性が高くなります。 GPT 呼び出しは正規表現です。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
9
v0.4.0
最新の
2026 年 7 月 4 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Catch the LLM/AI calls you didn't need. A fast, deterministic linter that flags LLM API calls where plain code is simpler, cheaper, and more reliable. Your GPT call is a regex. - theadamdanielsson/overllm

GitHub - theadamdanielsson/overllm: Catch the LLM/AI calls you didn't need. A fast, deterministic linter that flags LLM API calls where plain code is simpler, cheaper, and more reliable. Your GPT call is a regex. · GitHub
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
theadamdanielsson
/
overllm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows examples examples scripts scripts skills/ overllm skills/ overllm src/ overllm src/ overllm tests tests .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Catch the LLM/AI calls you didn't need.
overllm is a small, fast linter with one job: find the places in your code where you call an AI model to do something plain code does better. You called GPT to parse a date. You called a model to extract JSON that json.loads already handles. You are paying latency, money, and nondeterminism for a regex.
It reads your code with a real parser: Python through the standard-library ast , and JavaScript and TypeScript through tree-sitter. No model runs, no network, no API key. Same code in, same result out. Fast enough for a pre-commit hook.
Everyone else lints the code the AI wrote. overllm catches where you are paying an AI to do what a library already does.
pip install overllm # Python
pip install " overllm[js] " # adds JavaScript / TypeScript support
Use it
Point it at one file first, so you can see what it flags before you run it on everything:
overllm app.py # one file
overllm src/ # a folder
overllm . # the whole project
It reads your code and prints what it finds. It writes nothing and changes nothing — the worst case of running it is a few lines of output.
app.py:42:5 llm-mechanical LLM call asks the model to sort
resp = client.chat.completions.create(model="gpt-4o", messages=[...])
-> use sorted()
app.py:88:1 llm-in-loop LLM call inside a loop: one API round-trip per iteration
completion(model="gpt-4o", messages=[{"role": "user", "content": f"tag {x}"}])
-> batch the inputs into a single call, cache repeated results, or use a function
2 needless LLM calls in 1 file.
It is quiet by default. Only warning and error findings show, so a clean project prints nothing and exits 0 — most codebases surface a handful or none. If it floods you, treat that as a bug and open an issue .
overllm exits non-zero when it finds something, so it gates a commit or a CI check. Pass --exit-zero to report without failing.
Your code stays on your machine
overllm is static analysis. It parses your files locally, then prints what it found. It never uploads your code, never calls an API, needs no key, and sends no telemetry — there is no model in the loop and nothing phones home. Pull your network cable and it runs exactly the same.
The core is a few hundred lines of Python with no required dependencies, so you can read all of it before you trust it. overllm[js] adds tree-sitter to parse JavaScript and TypeScript; that is the only optional dependency.
Every rule fires only on a concrete code pattern, and every finding names the deterministic replacement. It stays silent when it is not sure.
By default overllm only raises warning and above, so it is quiet on your everyday code. static-prompt is info and stays silent unless you ask for it with --all or --min-severity info .
The last two check the call itself, not the prompt: a model id that no longer exists, or a knob the model ignores. Both are matched exactly against a known list, so a live model or alias is never flagged. The lists track provider deprecation pages and need updating over time.
It detects the OpenAI, Anthropic, Google, Mistral, Cohere, Groq, AWS Bedrock, HuggingFace, Replicate, LangChain, LiteLLM, and Ollama SDKs in Python, the Vercel AI SDK ( generateText , streamText , generateObject ) and the openai / anthropic node SDKs in JavaScript and TypeScript, and raw HTTP requests to those hosts.
resp = client . chat . completions . create (...) # overllm: ignore
resp = client . chat . completions . create (...) # overllm: ignore=llm-in-loop
Put # overllm: ignore-file at the top of a file to skip the whole file.
In pyproject.toml (Python 3.11+):
[ tool . overllm ]
ignore = [ " llm-in-loop " ]
exclude = [ " examples/ " , " migrations/ " ]
Or on the command line: --select , --ignore , --min-severity , --all , and --config PATH ( exclude is config-only). Run overllm --help for the full list.
repos :
- repo : https://github.com/theadamdanielsson/overllm
rev : v0.4.0
hooks :
- id : overllm
GitHub Action
overllm ships an Action that scans a pull request and leaves one grounded comment. It stays silent when there is nothing to say.
name : overllm
on :
pull_request :
permissions :
contents : read
pull-requests : write
jobs :
check :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : theadamdanielsson/overllm@v1
with :
paths : " . "
Other output formats
overllm --format json . # machine-readable
overllm --format sarif . # upload to GitHub code scanning
overllm --format markdown . # the PR-comment body
Why not just use an AI code reviewer?
AI reviewers and AI-slop linters look at the code the model produced: comments, dead code, structure. None of them ask the question overllm asks, which is whether you needed the model at all. It is a different axis, and it is one plain static analysis can answer with high precision and zero cost.
overloop, for your running agent
overllm reads your code at rest. Its runtime sibling, overloop , is a Claude Code hook that catches the same waste while an agent runs — the tool calls it repeats, the files it re-reads, the oversized output it floods context with. overllm catches the calls you didn't need; overloop catches the ones your agent runs twice. Two halves of the same idea.
The most useful thing you can send is a false positive: a real line of code where overllm flags a call it should not. A linter is only worth running if it is right, so one concrete bad flag is worth more than a feature request. CONTRIBUTING.md covers how to report one and how to run the tests.
Past releases and what changed are in CHANGELOG.md .
Catch the LLM/AI calls you didn't need. A fast, deterministic linter that flags LLM API calls where plain code is simpler, cheaper, and more reliable. Your GPT call is a regex.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
9
v0.4.0
Latest
Jul 4, 2026
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
