---
source: "https://github.com/raeudigerRaeffi/riddlerun"
hn_url: "https://news.ycombinator.com/item?id=48462197"
title: "Show HN: RiddleRun – AI run end-to-end browser tests"
article_title: "GitHub - raeudigerRaeffi/riddlerun: An open source agentic end2end testing tool for your webpages · GitHub"
author: "raffasch123"
captured_at: "2026-06-09T16:04:39Z"
capture_tool: "hn-digest"
hn_id: 48462197
score: 3
comments: 0
posted_at: "2026-06-09T15:13:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RiddleRun – AI run end-to-end browser tests

- HN: [48462197](https://news.ycombinator.com/item?id=48462197)
- Source: [github.com](https://github.com/raeudigerRaeffi/riddlerun)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T15:13:54Z

## Translation

タイトル: Show HN: RiddleRun – AI によるエンドツーエンドのブラウザ テストの実行
記事のタイトル: GitHub - raeudigerRaeffi/riddlerun: Web ページ用のオープン ソース エージェント エンドツーエンド テスト ツール · GitHub
説明: Web ページ用のオープンソース エージェント エンドツーエンド テスト ツール - raeudigerRaefi/riddlerun
HN テキスト: 5,000 行以上のコードをすべて徹底的にレビューせずにバイブコーディングしましたか?
あなたのアプリケーションは主に思考、祈り、そして疑わしい量のコピウムによってまとめられていますか?
エージェントがコミットするたびに開発ページ全体を実行して、ランダムに壊れたものがないかどうかを確認しますか?
もしそうなら、私はあなたのために何かを作りました。
ターミナルから直接実行できるオープンソースのエージェント型エンドツーエンド Web テスト フレームワークである、riddlerun の紹介です。
必要なのは、Docker、API キー、そしてアプリケーションがどのように動作するかを一貫した文で説明する能力だけです。フィードバック、問題点の報告、プロジェクトの将来の方向性についての提案を特にいただければ幸いです。

記事本文:
GitHub - raeudigerRaeffi/riddlerun: Web ページ用のオープンソース エージェント エンドツーエンド テスト ツール · GitHub
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
ライディガーラエフィ
/
なぞなぞ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows バックエンド バックエンド cli cli docs docs 例 例 フロントエンド フロントエンド メディア メディア移行 マイグレーション リドルラン リドルラン スキーマ スキーマ テスト テスト .dockerignore .dockerignore .env.example .env.example .env.production.example .env.production.example .gitignore .gitignore Dockerfile Dockerfile Dockerfile.api Dockerfile.api ライセンス ライセンス Makefile Makefile README.md README.md alembic.ini alembic.ini docker-compose.prod.yml docker-compose.prod.yml docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ナビゲーション ファイル
Docker ファースト AI のエンドツーエンドのブラウザー テスト。
RiddleRun は、自動化されたエージェントによるエンドツーエンド テストを必要とするチーム向けの CLI およびオプションの自己ホスト型 Web アプリです。 JSON でユーザー ジャーニーを記述し、Docker で実行して、Playwright/ブラウザー使用エージェントにテスト ケースを実行させます。
以下の実行例は、このリポジトリ内の JSON テストからのものであり、自己ホスト型 RiddleRun Web アプリから記録されました。
このリポジトリには、最初に独自のアプリをホストしなくても実行できる Wikipedia のサンプルが付属しています。
エンドツーエンドのテストでは、製品の品質が高くなることがよくあります。
バイブコーディング。コードベースはテストできるよりも速く変化します。
手動 E2E テストは拡張できません。すべてのリリースではコア フローを再度チェックする必要があり、製品、サポートされるブラウザ、ユーザー ジャーニーが拡大するにつれてコストは急速に増加します。
従来の Playwright スイートと Selenium スイートは脆弱で、メンテナンスに時間がかかることがあります。セレクターを多用するテストは UI 構造が変更されると中断されますが、最新の AI 支援開発ではアプリケーションの変更速度が向上し、その結果、チームがフィードバックを必要とする速度も速くなります。
CLI は Docker イメージとして公開されます。 Cl

サンプル プロジェクトとテスト ファイル用にこのリポジトリを 1 つ作成するか、準備ができたら独自の JSON ファイルをマウントします。
git clone https://github.com/raeudigerRaeffi/riddlerun.git
CDリドルラン
2. 公開された Docker イメージを使用して実行します
シェルに API キーを設定し、イメージをプルし、テスト例を 1 つ実行します。
$ env: OPENAI_API_KEY = " sk-... "
ドッカープルリドルラン / リドルラン:最新
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行 / ワークスペース / サンプル / テスト -- テスト ID ウィキペディア - Python - 言語 - ES
デフォルトでは、/data/artifacts/runs に JSON アーティファクトを書き込みます。ブラウザ使用ログはデフォルトで非表示になっているため、CLI はライブ結果テーブルを表示できます。デバッグ時に --verbose を追加します。
3. または、ソースから CLI イメージをビルドします
RiddleRun 自体を開発する場合、または公開されたイメージのプルを回避する場合は、このオプションを使用します。
CDリドルラン
docker build -triddlerun:local 。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e OPENAI_API_KEY `
リドルラン:ローカル `
実行 / ワークスペース / サンプル / テスト -- テスト ID ウィキペディア - Python - 言語 - ES
4. 独自のプロジェクトとテストを作成する
RiddleRun には、1 つのプロジェクト JSON ファイルと 1 つ以上のテスト ケース JSON ファイルが必要です。
プロジェクトはテスト対象のアプリケーションを記述します。
{
"id" : " ローカルデモ " ,
"name" : " ローカル デモ アプリ " ,
"baseUrl" : " http://host.docker.internal:3000 " ,
"認証" : {
"ユーザー名" : "demo@example.com" ,
"電子メール" : "demo@example.com" ,
"パスワード" : "チェンジミー"
}、
「ブラウザ」: {
"モード" : "ローカル" ,
"ヘッドレス" : true 、
「タイムアウト

さん「：120000」
}
}
テスト ケースはプロジェクトに属し、順序付けられた自然言語ステップが含まれます。
{
"id" : " ログインスモーク " ,
"name" : "ログインスモークテスト" ,
"projectId" : " local-demo " 、
"description" : " ユーザーがログインできることを確認します。 " ,
"タグ" : [ " スモーク " 、 " 認証 " ]、
「ステップ」: [
{
"description" : " プロジェクトのベース URL を開きます。 」
}、
{
"description" : " 保存されたプロジェクトの認証情報を使用してログインします。 「、
"expected_result" : " 認証されたダッシュボードが表示されます。 」
}
]
}
フィールドノート:
テストにアカウントが必要な場合、認証はオプションです
BaseUrl は、テスト対象のアプリの開始 URL です。 macOS および Windows 上の Docker から、host.docker.internal はホスト マシンを指します。
各テストの projectId はプロジェクト ID と一致する必要があります。
説明はブラウザ エージェントが何をすべきかを指示します。
Expected_result はオプションで、エージェントにどの結果を検証するかを指示します。
タグはオプションであり、テストのサブセットを簡単に実行できます。
完全な JSON コントラクトについては、 docs/test-spec.md 、 schemas/project.schema.json 、および schemas/testcase.schema.json を参照してください。
ディレクトリをマウントし、RIDDLE_RUN_PROJECT をプロジェクト ファイルに指定して、独自のファイルを実行します。
docker run -- rm `
- v " ${PWD} /my-riddlerun-tests:/workspace/tests:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / テスト / プロジェクト / ローカル - デモ.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行/ワークスペース/テスト/ケース -- タグスモーク
実行オプション
各テスト ケースは、1 つのブラウザ使用エージェント タスクとして実行されます。以下のオプションは、どのテストが選択されるか、エージェントが動作できる時間、どの LLM が使用されるか、および後で保存されるものを制御します。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / ウィキ

ペディア.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行/ワークスペース/サンプル/テスト -- ウィキペディアのタグ
意図的な失敗の例を実行します。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行/ワークスペース/例/テスト -- テスト ID ウィキペディア - Python - 言語 - 失敗
ブラウザを開かずに JSON ファイルを検証します。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
リドルラン / リドルラン:最新 `
検証 / ワークスペース / 例 / テスト
走行ビデオを録画します。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行 / ワークスペース / サンプル / テスト -- テスト ID ウィキペディア - Python - 言語 - es -- ビデオ - ディレクトリ / データ / アーティファクト / ビデオ
より高いエージェント ステップ バジェットを使用して、より長いフローを実行します。
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e OPENAI_API_KEY `
リドルラン / リドルラン:最新 `
実行/ワークスペース/サンプル/テスト `
-- テスト ID ウィキペディア - クロス - 記事 - リンク `
-- 最大 - ステップ 45 `
-- 冗長
CLI フラグ
旗
デフォルト
説明
--プロジェクト 、-p
リドル_ラン_プロジェクト
プロジェクトのメタデータ JSON ファイル。環境変数が設定されていない場合は必須です。
--テストID
すべてのテスト
指定されたテスト ID のみを実行します。複数のテストを選択するには、フラグを繰り返します。
--タグ
すべてのテスト
指定されたタグのいずれかを含むテストのみを実行します。リストされたタグと一致するまで繰り返します。
--最大ステップ数
29

テストごとに許可されるブラウザ使用エージェント アクションの最大数。複数のステップを含む長い移動の場合は、この値を増やします。
--ドライラン
オフ
ブラウザを開かずに、選択したテストを検証し、スキップされた結果を印刷します。
--アーティファクトディレクトリ
Docker の /data/artifacts
JSON 実行結果ファイルが書き込まれるディレクトリ。
--ビデオディレクトリ
オフ
画面録画を有効にし、ビデオをこのディレクトリに保存します。
--冗長
オフ
ライブ結果テーブルのみではなく、ブラウザー使用ログを表示します。
--llm プロバイダー
オープンナイ
LLM プロバイダーはブラウザーでの使用に渡されました。
--llm-モデル
o4-mini
選択したプロバイダーのモデル名。
--llm-API キー
プロバイダー環境変数
API キーのオーバーライド。 Docker ではプロバイダー固有の環境変数を優先します。
--llm-ベース URL
プロバイダのデフォルト
ローカルプロバイダーまたは Ollama などの互換プロバイダーのカスタムベース URL。
環境変数
すべての設定は、RIDDLE_RUN_ プレフィックスを使用して設定することもできます。開始点については、.env.example を参照してください。
プロジェクトのブラウザーの動作は、プロジェクトの JSON ブラウザー ブロックで定義されます。
RiddleRun は実行ごとに Docker コンテナ内で Chromium を起動します。ブラウザはホスト マシンには表示されません。視覚的な記録が必要な場合は、--video-dir を使用してください。
RiddleRun は、構成可能な LLM プロバイダーを使用してブラウザーを使用します。 OpenAI がデフォルトです。
$ env: OPENAI_API_KEY = " sk-... "
macOS と Linux:
エクスポート OPENAI_API_KEY= " sk-... "
次に、 -e OPENAI_API_KEY を使用して変数を Docker に渡します。 CLI フラグまたは RIDDLE_RUN_LLM_API_KEY オーバーライドが設定されていない場合、 OPENAI_API_KEY 、 ANTHROPIC_API_KEY 、 GOOGLE_API_KEY 、 MISTRAL_API_KEY 、 GROQ_API_KEY などのプロバイダー固有のキーが自動的に読み取られます。
Docker Compose 開発の場合は、キーを .env に置きます。
RIDDLE_RUN_LLM_PROVIDER=オープンアイ
RIDDLE_RUN_LLM_MODEL=o4-mini
OPENAI_API_KEY=sk-...
プロバイダーの例
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/art

ifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e ANTHROPIC_API_KEY `
リドルラン / リドルラン:最新 `
実行 / ワークスペース / サンプル / テスト -- テスト ID ウィキペディア - Python - 言語 - es -- llm - プロバイダー anthropic -- llm - モデル クロード - ソネット - 4 - 5
地元のオラマ:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
リドルラン / リドルラン:最新 `
実行/ワークスペース/サンプル/テスト -- test-id wikipedia - Python - 言語 - es -- llm - プロバイダー ollama -- llm - モデル llama3。 1 -- llm - ベース - URL http://host.docker.internal: 11434
LiteLLM プロバイダー (例: Groq):
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ ワークスペース / サンプル / プロジェクト / wikipedia.json `
- e GROQ_API_KEY `
リドルラン / リドルラン:最新 `
実行 / ワークスペース / サンプル / テスト -- テスト ID ウィキペディア - Python - 言語 - es -- llm - プロバイダー groq -- llm - モデル groq / llama - 3.3 - 70b - 多用途
サポートされているプロバイダーには、 openai 、 anthropic 、 google 、 mistral 、 o が含まれます

[切り捨てられた]

## Original Extract

An open source agentic end2end testing tool for your webpages - raeudigerRaeffi/riddlerun

Did you vibe-code 5k+ lines of code without thoroughly reviewing all of them?
Is your application held together mostly by thoughts, prayers, and a suspicious amount of copium ?
Do you run through your entire development page after every agent commit just to check that nothing randomly broke?
If yes, I built something for you.
Introducing riddlerun: an open-source agentic end-to-end web testing framework that can be run directly from the terminal.
All you need is Docker, an API key, and the ability to describe in a coherent sentence how your application is supposed to behave. I’d be especially grateful for feedback, issue reports, and proposals for the future direction of the project.

GitHub - raeudigerRaeffi/riddlerun: An open source agentic end2end testing tool for your webpages · GitHub
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
raeudigerRaeffi
/
riddlerun
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows backend backend cli cli docs docs examples examples frontend frontend media media migrations migrations riddlerun riddlerun schemas schemas tests tests .dockerignore .dockerignore .env.example .env.example .env.production.example .env.production.example .gitignore .gitignore Dockerfile Dockerfile Dockerfile.api Dockerfile.api LICENSE LICENSE Makefile Makefile README.md README.md alembic.ini alembic.ini docker-compose.prod.yml docker-compose.prod.yml docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
Docker-first AI end-to-end browser testing.
RiddleRun is a CLI and optional self-hosted web app for teams that want automated agentic end2end test. Describe a user journey in JSON, run it in Docker, and let a Playwright/browser-use agent execute your test case for you!
The example runs below come from the JSON tests in this repository and were recorded from the self-hosted RiddleRun web app.
The repo ships with Wikipedia examples you can run without hosting your own app first:
End-to-end testing is where product quality often becomes expensive:
Vibecoding. Your codebase changes faster than you can test.
Manual E2E testing does not scale. Every release needs core flows checked again, and the cost grows quickly as your product, supported browsers, and user journeys expand.
Traditional Playwright and Selenium suites can be brittle and slow to maintain. Selector-heavy tests break when UI structure changes, while modern AI-assisted development increases the speed at which applications change and therefore the speed at which teams need feedback.
The CLI is published as a Docker image. Clone this repo for the example project and test files, or mount your own JSON files once you are ready.
git clone https: // github.com / raeudigerRaeffi / riddlerun.git
cd riddlerun
2. Run with the published Docker image
Set an API key in your shell, pull the image, and run one example test:
$ env: OPENAI_API_KEY = " sk-... "
docker pull riddlerun / riddlerun:latest
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - es
Runs write JSON artifacts under /data/artifacts/runs by default. Browser-use logs are hidden by default so the CLI can show a live results table; add --verbose when debugging.
3. Or build the CLI image from source
Use this option when you want to develop RiddleRun itself or avoid pulling the published image.
cd riddlerun
docker build - t riddlerun:local .
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun:local `
run / workspace / examples / tests -- test-id wikipedia - python - language - es
4. Create your own project and tests
RiddleRun needs one project JSON file and one or more test case JSON files.
A project describes the application under test:
{
"id" : " local-demo " ,
"name" : " Local demo app " ,
"baseUrl" : " http://host.docker.internal:3000 " ,
"auth" : {
"username" : " demo@example.com " ,
"email" : " demo@example.com " ,
"password" : " change-me "
},
"browser" : {
"mode" : " local " ,
"headless" : true ,
"timeoutMs" : 120000
}
}
A test case belongs to a project and contains ordered natural-language steps:
{
"id" : " login-smoke " ,
"name" : " Login smoke test " ,
"projectId" : " local-demo " ,
"description" : " Verify that a user can log in. " ,
"tags" : [ " smoke " , " auth " ],
"steps" : [
{
"description" : " Open the project base URL. "
},
{
"description" : " Log in using the saved project credentials. " ,
"expected_result" : " The authenticated dashboard is visible. "
}
]
}
Field notes:
auth is optional if testing requires an account
baseUrl is the starting URL for the app under test. From Docker on macOS and Windows, host.docker.internal points back to your host machine.
projectId in each test must match the project id .
description tells the browser agent what to do.
expected_result is optional and tells the agent what outcome to validate.
tags are optional and make it easier to run subsets of tests.
See docs/test-spec.md , schemas/project.schema.json , and schemas/testcase.schema.json for the full JSON contract.
Run your own files by mounting their directory and pointing RIDDLE_RUN_PROJECT at your project file:
docker run -- rm `
- v " ${PWD} /my-riddlerun-tests:/workspace/tests:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / tests / projects / local - demo.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / tests / cases -- tag smoke
Run Options
Each test case is executed as one browser-use agent task. The options below control which tests are selected, how long the agent can work, which LLM is used, and what gets saved afterward.
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- tag wikipedia
Run the intentional failure example:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - fail
Validate JSON files without opening a browser:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
riddlerun / riddlerun:latest `
validate / workspace / examples / tests
Record a run video:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - es -- video - dir / data / artifacts / videos
Run a longer flow with a higher agent step budget:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e OPENAI_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests `
-- test-id wikipedia - cross - article - links `
-- max - steps 45 `
-- verbose
CLI flags
Flag
Default
Description
--project , -p
RIDDLE_RUN_PROJECT
Project metadata JSON file. Required unless the env var is set.
--test-id
all tests
Run only the given test IDs. Repeat the flag to select multiple tests.
--tag
all tests
Run only tests that include one of the given tags. Repeat to match any listed tag.
--max-steps
29
Maximum browser-use agent actions allowed per test. Increase this for longer multi-step journeys.
--dry-run
off
Validate the selected tests and print skipped results without opening a browser.
--artifact-dir
/data/artifacts in Docker
Directory where JSON run result files are written.
--video-dir
off
Enable screen recording and save videos under this directory.
--verbose
off
Show browser-use logs instead of only the live results table.
--llm-provider
openai
LLM provider passed to browser-use.
--llm-model
o4-mini
Model name for the selected provider.
--llm-api-key
provider env var
API key override. Prefer provider-specific env vars in Docker.
--llm-base-url
provider default
Custom base URL for local or compatible providers such as Ollama.
Environment variables
All settings can also be set with the RIDDLE_RUN_ prefix. See .env.example for a starting point.
Browser behavior for a project is defined in the project JSON browser block:
RiddleRun launches Chromium inside the Docker container for each run. The browser is not displayed on your host machine; use --video-dir when you need a visual recording.
RiddleRun uses browser-use with a configurable LLM provider. OpenAI is the default.
$ env: OPENAI_API_KEY = " sk-... "
macOS and Linux:
export OPENAI_API_KEY= " sk-... "
Then pass the variable into Docker with -e OPENAI_API_KEY . Provider-specific keys such as OPENAI_API_KEY , ANTHROPIC_API_KEY , GOOGLE_API_KEY , MISTRAL_API_KEY , and GROQ_API_KEY are read automatically when no CLI flag or RIDDLE_RUN_LLM_API_KEY override is set.
For Docker Compose development, put keys in .env :
RIDDLE_RUN_LLM_PROVIDER=openai
RIDDLE_RUN_LLM_MODEL=o4-mini
OPENAI_API_KEY=sk-...
Provider examples
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e ANTHROPIC_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - es -- llm - provider anthropic -- llm - model claude - sonnet - 4 - 5
Local Ollama:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - es -- llm - provider ollama -- llm - model llama3. 1 -- llm - base - url http: // host.docker.internal: 11434
LiteLLM provider, for example Groq:
docker run -- rm `
- v " ${PWD} /examples:/workspace/examples:ro " `
- v " ${PWD} /artifacts:/data/artifacts " `
- e RIDDLE_RUN_PROJECT =/ workspace / examples / projects / wikipedia.json `
- e GROQ_API_KEY `
riddlerun / riddlerun:latest `
run / workspace / examples / tests -- test-id wikipedia - python - language - es -- llm - provider groq -- llm - model groq / llama - 3.3 - 70b - versatile
Supported providers include openai , anthropic , google , mistral , o

[truncated]
