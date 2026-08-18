---
source: "https://github.com/Tonterias/h2aichat"
hn_url: "https://news.ycombinator.com/item?id=49344277"
title: "Show HN: H2AI Chat – multi-AI debate you can self-host with local models (AGPL)"
article_title: "GitHub - Tonterias/h2aichat: Several AIs debate a topic while you moderate — self-hostable, and it runs on your own local models. Open-source (AGPL). · GitHub"
image: "https://opengraph.githubassets.com/2c4fac9623382af7231ad619f46d3438780e8bdc1bfdf55ebf88947d8dda66b2/Tonterias/h2aichat"
author: "h2aichat"
captured_at: "2026-08-18T12:25:12Z"
capture_tool: "hn-digest"
hn_id: 49344277
score: 1
comments: 0
posted_at: "2026-08-18T11:47:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: H2AI Chat – multi-AI debate you can self-host with local models (AGPL)

- HN: [49344277](https://news.ycombinator.com/item?id=49344277)
- Source: [github.com](https://github.com/Tonterias/h2aichat)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T11:47:04Z

## Translation

タイトル: 表示 HN: H2AI チャット – ローカル モデルで自己ホストできるマルチ AI ディベート (AGPL)
記事のタイトル: GitHub - Tonterias/h2aichat: あなたが司会を務めている間、複数の AI がトピックについて議論します。自己ホスト可能であり、独自のローカル モデルで実行されます。オープンソース (AGPL)。 · GitHub
説明: あなたが司会を務めている間、複数の AI がトピックについて議論します。自己ホスト可能であり、独自のローカル モデルで実行されます。オープンソース (AGPL)。 - トンテリアス/h2aichat

記事本文:
GitHub - Tonterias/h2aichat: あなたが司会を務めている間、複数の AI がトピックについて議論します。自己ホスト可能であり、独自のローカル モデルで実行されます。オープンソース (AGPL)。 · GitHub
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
トンテリア
/
h2aiチャット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
28 コミット 28 コミット フォルダーとファイル
.github .github アセット アセット 会話/ ja 会話/ en docs ドキュメント 実行 実行 .dockerignore .dock

erignore .env.example .env.example .gitignore .gitignore CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
複数の AI が互いに議論し、あなたが司会を務めます。人間が司会を務めながら、さまざまなモデル (それぞれに独自の役割を持つ) が話し、互いに挑戦し、アイデアを対比させるターンベースの会話システム。自己ホスト可能で、独自のローカル モデルで実行できるため、データがマシンから流出することはありません。
このデモでは、彼らは愚かな言葉遊びをしますが、実際のチャットでは、驚くべき推論、創造性、コラボレーションが明らかになります。研究によると、対立する意見の間でシミュレートされた議論は、単一の AI 単独よりも優れた推論を生み出すことがわかっています。
議論であって、単一の声ではありません。性格も役割も異なる数人のモデルが交代で登場。
必要に応じて、あなたのマシン上で。ローカル モデル (LM Studio / Ollama) で動作します。議論全体がコンピューター上で実行されるため、会話がコンピューターから離れることはありません。
あなたが責任者です。あなたが司会を務め、介入し、一時停止し、誰が発言するかを決定します。
複数の AI 間の 2 つの議論の例 (h2aichat.com でライブ配信):
完全なソースは、このリポジトリのmessages/en/にもあります。
pip install -r 要件.txt
python -m uvicorn 実行.api_server:app --port 8000
# http://localhost:8000 を開きます
独自のモデルを使用してマシン上で 100% 実行する (コンピューターから何も残らない) 場合、最も速い方法は Docker です。
ドッカー構成 -d
出典: link タグ:dockercomposeexec ollama ollama pull llama3.2
# http://localhost:8000 を開きます -> 議論する準備ができています: 「ローカル」エージェントが付属しています
フルガイド (Docker または LM Studio / Ol)

ラマ、および正直なプライバシーに関する注意事項): docs/RUN_LOCALLY.md 。キーを .env に置きます ( .env.example を参照)。決して実際のキーをコミットしないでください。
若くて進化中のプロジェクト。荒削りな部分や欠けている部分がありますが、フィードバックや貢献を歓迎します ( COTRIBUTING.md を参照)。
ぜひご協力をお願いいたします。初めてコードを投稿する前に、CLA に同意するように求められます (ワンクリック)。その理由は COTRIBUTING.md にあります。
H2AI Chat は AGPL-3.0 に基づいてリリースされています (「ライセンス」を参照)。自由に使用、研究、修正、共有することができます。 AGPL では、変更されたバージョンをネットワーク サービスとして提供する場合は、変更を公開する必要があるという条件が 1 つ追加されます。これにより、プロジェクトは誰でもオープンな状態に保たれ、誰かがコピーを閉じたり、コミュニティに対して独自のサービスを実行したりすることができなくなります。また、h2aichat.com でホスト型バージョンと企業向けエディションも実行しています (開発資金はそこにあります)。
「H2AI Chat」およびそのロゴは、ミゲル・アンヘル・スアレスの商標です。 AGPL ライセンスは、名前やブランドではなくコードを対象としています。コードをフォークすることはできますが、「H2AI Chat」という名前や独自のバージョンのロゴを使用することはできません。
あなたが司会を務めている間、複数の AI がトピックについて議論します。自己ホスト可能であり、独自のローカル モデルで実行されます。オープンソース (AGPL)。
Readme AGPL-3.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Several AIs debate a topic while you moderate — self-hostable, and it runs on your own local models. Open-source (AGPL). - Tonterias/h2aichat

GitHub - Tonterias/h2aichat: Several AIs debate a topic while you moderate — self-hostable, and it runs on your own local models. Open-source (AGPL). · GitHub
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
Tonterias
/
h2aichat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
28 Commits 28 Commits Folders and files
.github .github assets assets conversations/ en conversations/ en docs docs execution execution .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
Several AIs debate with each other, and you moderate. A turn-based conversation system where different models (each with its own role) talk, challenge one another, and contrast ideas — while a human moderates. Self-hostable, and you can run it with your own local models, so your data never leaves your machine.
In this demo, they play a silly word game, but their real chats reveal surprising reasoning, creativity, and collaboration. Research suggests that simulated debate between opposing views produces better reasoning than a single AI alone.
A debate, not a single voice. Several models with different personalities and roles, taking turns.
On your machine, if you want. Works with local models (LM Studio / Ollama): the whole debate runs on your computer, so your conversations never leave it.
You're in charge. You moderate, step in, pause, and decide who speaks.
Two example debates between several AIs (rendered live on h2aichat.com):
Their full source is also in conversations/en/ in this repo.
pip install -r requirements.txt
python -m uvicorn execution.api_server:app --port 8000
# open http://localhost:8000
To run it 100% on your machine with your own models (nothing leaves your computer), the fastest way is Docker:
docker compose up -d
docker compose exec ollama ollama pull llama3.2
# open http://localhost:8000 -> ready to debate: it ships with a "Local" agent
Full guide (Docker or LM Studio / Ollama, plus an honest privacy note): docs/RUN_LOCALLY.md . Put your keys in .env (see .env.example ); never commit real keys.
Young, evolving project. There are rough edges and missing pieces — feedback and contributions are welcome (see CONTRIBUTING.md ).
We'd love your help. Before your first code contribution you'll be asked to accept the CLA (one click); the why is in CONTRIBUTING.md .
H2AI Chat is released under AGPL-3.0 (see LICENSE ). You can use, study, modify, and share it freely. The AGPL adds one condition: if you offer a modified version as a network service, you must publish your changes. This keeps the project open for everyone and prevents anyone from closing a copy and running a proprietary service against the community. We also run a hosted version and an edition for companies at h2aichat.com (that's what funds development).
“H2AI Chat” and its logo are trademarks of Miguel Ángel Suárez . The AGPL license covers the code , not the name or the branding: you may fork the code, but you may not use the name “H2AI Chat” or the logo for your own version.
Several AIs debate a topic while you moderate — self-hostable, and it runs on your own local models. Open-source (AGPL).
Readme AGPL-3.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
