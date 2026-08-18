---
source: "https://github.com/softcane/agents-workbook"
hn_url: "https://news.ycombinator.com/item?id=49342012"
title: "Show HN: Agents Workbook watch Claude Code, Codex write down their working notes"
article_title: "GitHub - softcane/agents-workbook · GitHub"
image: "https://opengraph.githubassets.com/05aa5bd5c164b6976ef45216e6b98c1368198de7c65904b93c77c0d893848fbb/softcane/agents-workbook"
author: "pradeep1177"
captured_at: "2026-08-18T06:25:30Z"
capture_tool: "hn-digest"
hn_id: 49342012
score: 1
comments: 0
posted_at: "2026-08-18T06:21:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agents Workbook watch Claude Code, Codex write down their working notes

- HN: [49342012](https://news.ycombinator.com/item?id=49342012)
- Source: [github.com](https://github.com/softcane/agents-workbook)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T06:21:29Z

## Translation

タイトル: HN を表示: エージェント ワークブックを見てクロード コード、コーデックスが作業メモを書き留める
記事のタイトル: GitHub - ソフトケーン/エージェントワークブック · GitHub
説明: GitHub でアカウントを作成して、ソフトケーン/エージェント ワークブックの開発に貢献します。

記事本文:
GitHub - ソフトケーン/エージェントワークブック · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ソフトケーン
/
エージェントワークブック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
4 コミット 4 コミット フォルダーとファイル
.mvn/ ラッパー .mvn/ ラッパー ライセンス ライセンス docs docs envoy envoy src src .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンス通知 NOTICE README.md README.md SECURITY.md SECURITY.md docker-

compose.yml docker-compose.yml mvnw mvnw mvnw.cmd mvnw.cmd pom.xml pom.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントにワークブックを渡します。エージェントは大声で考えてその推論を書き、それが機能している間、あなたはそれを読みます。
クロード コードまたはコーデックスから送信されるすべてのリクエストに 1 つのツールを追加するローカル プロキシ: 声に出して考える場所。モデルが書き込み、メモが 127.0.0.1:8080 のダッシュボードに表示され、クライアントは通常の応答を返します。
考え方をまとめたものではありません。おおよその考え方。デフォルトの文言では、モデルがメモ内で問題を解決し、何を重視しているのか、除外した選択肢は何か、なぜ選択した選択肢がそれらを上回るのかをモデルに求めます。ノートは長く続きます、それがポイントです。
これは、次のような質問に答えるために作成しました: エージェントが表明した計画は、エージェントが今後行うことと一致していますか?
無理な推論にこれを使用しないでください
これは楽しい実験です。データ収集ツールではありません。
Anthropic または OpenAI のモデルから推論トレースを抽出、マイニング、再構築するためにこれを使用しないでください。キャプチャしたものをトレーニングしないでください。そこからデータセットを構築したり、公開したり、他の場所でモデルの動作を再現するためにメモを使用したりしないでください。
ワークブック ノートは、セッションのマシン上で通常のツール呼び出しを通じて書き込まれた通常のモデル出力であり、自分のエージェントの動作を観察できます。それがすべての使用目的です。それを収集することは両方のプロバイダーの規約に違反しており、これは目的ではありません。
docker 構成 -d --build
http://127.0.0.1:8080 を開き、タブを開いたままにします。
ANTHROPIC_BASE_URL=http://127.0.0.1:10000 クロード
コーデックス:
codex --disableenable_request_compression \
-c ' モデルプロバイダー="ワークブックプロキシ" ' \
-c ' model_providers.workbook_proxy.name="ローカル ワークブック プロキシ" ' \
-c ' model_providers.workbook_proxy.base_url="http://127.0.0.1:10000/bac

kend-api/コーデックス" ' \
-c 'model_providers.workbook_proxy.wire_api="応答" ' \
-c ' model_providers.workbook_proxy.requires_openai_auth=true ' \
-c 'model_providers.workbook_proxy.supports_websockets=false'
docker compose down で停止します。 WORKBOOK_ARCHIVE_PATH を設定しない限り、ディスクには何も到達しません。
これによりトークンが燃焼します。各ターンはプロバイダー呼び出しを 1 つではなく 2 つ行い、モデルがメモを書き続ける場合は最大 4 つ呼び出します。すべての通話には会話履歴全体が含まれるため、入力コストは 2 倍になります。その後、メモ自体は、実際に必要な答えに加えて、ターンごとに推論の約 12,000 個の出力トークンを要求します。レート制限を超えるサブスクリプションの場合。 API キーではお金を食います。大声で考えることは無料ではありません。
答えは後で始まります。プロキシは、返信のストリーミングが開始される前に、完全なメモを待ちます。非表示の呼び出しは、プロキシが諦めて代わりに元のリクエストを送信するまで 60 秒かかります。
独自のサブスクリプション資格情報を含むリクエストを変更しています。 Anthropic と OpenAI は両方とも、カスタム ベース URL での CLI のポイントをサポートしているため、ルーティング自体は文書化された設定です。リクエストにツールを追加するのは私の仕事であり、それを実行するのはあなたの仕事です。プロバイダーの規約を読んで、自分で電話をかけます。
メモから機密が漏洩する可能性があります。このツールは、認証情報をモデルに漏洩しないようにモデルに要求します。それは要求であって、境界ではありません。 API キーをプロンプトに貼り付ける場合は、メモに表示される可能性があると想定してください。
プロバイダーの保護された内部推論を読み取ったり、公開したり、回復しようとしたりすることはありません。それは不可能です。ワークブック ノートは、通常のツール呼び出しを通じて書き込まれたモデル出力であり、モデルが生成する他のツール引数と何ら変わりません。
キャプチャしたメモをモデルのトレーニングや改善に使用しないでください。
AI エンジニアリングの役割を歓迎します。このプロジェクトは、私がどのように仕事をするかを示す公正なサンプルです。

rk: github.com/softcane 、softcane@gmail.com 。
Apache 2.0 については、「ライセンスと通知」を参照してください。 Anthropic または OpenAI と提携、承認、後援されていません。
Readme Apache-2.0 ライセンス セキュリティ ポリシー
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to softcane/agents-workbook development by creating an account on GitHub.

GitHub - softcane/agents-workbook · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
softcane
/
agents-workbook
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.mvn/ wrapper .mvn/ wrapper LICENSES LICENSES docs docs envoy envoy src src .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml mvnw mvnw mvnw.cmd mvnw.cmd pom.xml pom.xml View all files Repository files navigation
Give your coding agent a workbook; the agent thinks out loud and writes its reasoning, then you reads it while it works.
A local proxy that adds one tool to every request going out of Claude Code or Codex: somewhere to think out loud. The model writes, the notes land on a dashboard at 127.0.0.1:8080 , and your client gets back the ordinary reply.
Not a summary of the thinking. The aprox thinking. The default wording asks the model to work the problem through inside the note, what it is weighing, the options it ruled out, and why the one it picked beat them. Notes run long, and that is the point.
I built this to answer the question like: does an agent's stated plan match what it goes on to do?
Don't use this to farm reasoning
This is a fun experiment. It is not a data collection tool.
Do not use it to distill, mine, or reconstruct reasoning traces from Anthropic's or OpenAI's models. Do not train on what it captures. Do not build a dataset out of it, publish one, or use the notes to reproduce a model's behaviour anywhere else.
A workbook note is ordinary model output, written through an ordinary tool call, on your machine, for your session, so that you can watch your own agent work. That is the entire intended use. Harvesting it is a violation of both providers' terms and it is not what this is for.
docker compose up -d --build
Open http://127.0.0.1:8080 and leave the tab open.
ANTHROPIC_BASE_URL=http://127.0.0.1:10000 claude
Codex:
codex --disable enable_request_compression \
-c ' model_provider="workbook_proxy" ' \
-c ' model_providers.workbook_proxy.name="Local Workbook Proxy" ' \
-c ' model_providers.workbook_proxy.base_url="http://127.0.0.1:10000/backend-api/codex" ' \
-c ' model_providers.workbook_proxy.wire_api="responses" ' \
-c ' model_providers.workbook_proxy.requires_openai_auth=true ' \
-c ' model_providers.workbook_proxy.supports_websockets=false '
docker compose down stops it. Nothing reaches disk unless you set WORKBOOK_ARCHIVE_PATH .
This burns tokens. Each turn makes two provider calls instead of one, and up to four when the model keeps writing notes. Every call carries the whole conversation history, so your input cost doubles. The note itself then asks for around 12,000 output tokens of reasoning per turn, on top of the answer you actually wanted. On a subscription that eats your rate limit. On an API key it eats money. Thinking out loud is not free.
Your answer starts later. The proxy waits for the full note before your reply begins streaming. The hidden call gets 60 seconds before the proxy gives up and sends your original request instead.
You are modifying requests that carry your own subscription credentials. Anthropic and OpenAI both support pointing their CLIs at a custom base URL, so the routing itself is a documented setting. Adding a tool to the request is my doing, and running it is yours. Read your provider's terms and make your own call.
Notes can leak secrets. The tool asks the model to keep credentials out of them. That is a request, not a boundary. If you paste an API key into a prompt, assume it can surface in a note.
It does not read, expose, or try to recover any provider's protected internal reasoning; that's not possible. A workbook note is model output written through an ordinary tool call, no different from any other tool argument the model produces.
Do not use captured notes to train or improve a model.
I am open to AI engineering roles. This project is a fair sample of how I work: github.com/softcane , softcane@gmail.com .
Apache 2.0, see LICENSE and NOTICE . Not affiliated with, endorsed by, or sponsored by Anthropic or OpenAI.
Readme Apache-2.0 license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
