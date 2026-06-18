---
source: "https://github.com/bleak-ai/gcontext"
hn_url: "https://news.ycombinator.com/item?id=48582222"
title: "Show HN: Gcontext – a tree of llms.txt files to steer agents on support tasks"
article_title: "GitHub - bleak-ai/gcontext: Context Management System to let your agent find and build relevant context · GitHub"
author: "bsampera"
captured_at: "2026-06-18T10:36:37Z"
capture_tool: "hn-digest"
hn_id: 48582222
score: 1
comments: 0
posted_at: "2026-06-18T07:55:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Gcontext – a tree of llms.txt files to steer agents on support tasks

- HN: [48582222](https://news.ycombinator.com/item?id=48582222)
- Source: [github.com](https://github.com/bleak-ai/gcontext)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T07:55:35Z

## Translation

タイトル: HN を表示: Gcontext – エージェントをサポート タスクに導くための llms.txt ファイルのツリー
記事のタイトル: GitHub - bleak-ai/gcontext: エージェントが関連するコンテキストを見つけて構築できるようにするコンテキスト管理システム · GitHub
説明: エージェントが関連するコンテキストを検索して構築できるようにするコンテキスト管理システム - bleak-ai/gcontext
HN テキスト: 私は格闘技ジム ソフトウェア (MAAT) を作成するスタートアップで働いています。私たちは、支払いシステムとデータベースを使用して、ジムのオーナーが行う必要がないように学生のメンバーシップを処理します。ジムの数が増えるにつれて、サブスクリプションの問題、メンバーシップの更新、データのエクスポートなど、サポート タスクも増えます。現在、私たちがそれを解決する方法は「llms.txt のツリー」です。 llms.txt は通常、Web サイトまたはドキュメントで入手可能な情報を参照します。私たちは内部で同じ考え方を使用して、エージェントが必要とする情報を整理します。エージェントはフォルダーから開始し、下に移動します。
§── llms.txt # このレベルの各フォルダーを参照します
§── Stripe/ # info.md: Stripe アカウントの構造
§── firestore/ # info.md: スキーマの様子
└── サポート/
§── info.md # サポートタスクの解決方法
§── runbook/ # タスクごとに 1 つのファイル、独自の llms.txt を含む
│ §── サブスクリプションのキャンセル.md
│ §──export-gym-data.md
│ └── fix-membership-mismatch.md
└── logs/ # エージェントが解決したすべてのタスクを 1 日あたり 1 ファイル
コピー
これにより、エージェントをより適切に操作し、新しいサポート タスクが発生するたびに新しい Runbook を作成できるようになります。すべての統合に、エージェントが触れるべきものと触れるべきでないものを追加できます。 gcontext プロンプトは、ガードレールが注意深く遵守されていることを確認します。

記事本文:
GitHub - bleak-ai/gcontext: エージェントが関連するコンテキストを見つけて構築できるようにするコンテキスト管理システム · GitHub
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
暗い愛
/
gcontext
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット case-studies/ maat-support case-studies/ maat-support コア コア デモ デモ サンプル 例 gcontext gcontext テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
主な概念は、フォルダーまたはファイルを参照する llms.txt のツリーです。目標は、AI エージェントとの会話に常に適切な情報をロードし、そこからエージェントがアクセスできるコンテキストを拡張し、物事の「ライブ メモリ」またはライブ レコードに似たものを作成できるようにすることです。
クロードコードメモリーシステムとの違いは何ですか？
問題は同じですが、Gcontext はコンテキストをそれ自体の問題として扱いたいユーザー向けであり、Memory が舞台裏ですべてを処理するのに対し、Gcontext ではユーザーがコンテキストをどのように定義するか、AI エージェントが実行する手順をより正確に定義できます。
1. ワークスペースをインストールして作成します。
カール -LsSf https://gcontext.ai/gcontext/install.sh | sh # または: uv ツールをインストール gcontext-ai
gcontextの初期化
2. エージェントにモジュールを構築するよう依頼します。ワークスペースでエージェントを開き、必要な統合に名前を付けます。
supabase 統合モジュールを作成してロードします。
エージェントは、 info.md 、 llms.txt インデックス、および必要なシークレットを名前のみで宣言する module.yaml を書き込みます (値は .env から離れることはありません)。
3. 要求される秘密を入力します。このモジュールは、どの変数を設定するかを指示します。それらを .env に置きます。
SUPABASE_URL=https://....supabase.co
SUPABASE_SECRET_KEY=...
4. 実際のデータでモジュールを強化します。アクセスが確立されたので、エージェントにライブ サービスを探索し、見つかった内容を書き留めるよう依頼します。
スーパーブを見てください

実際の読み取りリクエストを使用してモジュール情報を更新します。
それ以降、新しいセッションでは、モジュールのインデックスをたどり、 .env のキーを使用して API を呼び出すことで、質問 (「メンバー テーブルに行は何行ありますか?」) に答えることができます。
新しいサービスごとにステップ 2 を繰り返すと、ツリーが形成され始めます。ルート llms.txt は、統合ごとに 1 つのモジュールにルーティングされ、それぞれが独自のキーを宣言します。ただ尋ねてください：
ストライプ統合モジュールを作成してロードします。
統合を追加すればするほど、ワークスペースはすべての統合へのアクセスを調整する単一の場所になります。まさにこのときに gcontext が効果を発揮します。
これは誰にでも当てはまるわけではありません。唯一の外部依存関係がデータベースであるプロジェクトで作業している場合、これのためだけにこれらすべてを設定するのはおそらく意味がありません。しかし、使用する外部の依存関係や統合が増えるにつれて、これらすべてへのアクセスを調整するための専用の中央の場所を用意することがより合理的になります。
このプロジェクトは、格闘技ジム向け管理ソフトウェアであるスタートアップ MAAT で開発されており、製品は支払いシステムと DB の上に構築されています。ジムが増えるにつれて、これらすべての統合が影響を受けるサポート タスクの解決にますます多くの時間を費やす必要がありました。私たちは、日々の作業をスピードアップするために、gcontext が llms.txt に従うパターンを考え出しました。
サポートタスクを解決する私たちの旅
AI以前。サポート タスクを解決する方法に関するいくつかのプレイブックがあり、それらを使用してタスクを手動で解決しました。
AIあり - 主にCLAUDE.md。私たちは Claude.md でシステムがどのように機能するのか、そして問題の最も一般的な原因は何かを説明しました。ほとんどの場合、正しい兆候が示されましたが、Runbook を保存する場所がなかったため、依然として多くのタスクについて手動作業と調査を行う必要があり、依然として繰り返し作業が必要でした。
AI+で

GContext これにより、AI をプロセスに完全に組み込むことができました。 When we do now an exploration for a task, we can save this exploration in a md file, referenced by an llms.txt and the AI model will be able to find it later.
ここでその例を参照できます: case-studies/maat-support/ 。 It is the tree from the diagram above: a root llms.txt router over the stripe and firestore integration modules, plus a support module with its runbooks and execution logs.
残りの作業はエージェントがわかりやすい言葉で行います。次のように尋ねてください。
Bleak AI によって構築されました | gcontext.ai
エージェントが関連するコンテキストを見つけて構築できるようにするコンテキスト管理システム
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

Context Management System to let your agent find and build relevant context - bleak-ai/gcontext

I work at a startup that makes martial arts gym software (MAAT). We handle the memberships of students so gym owners don't have to, using a payment system and a database. As we get more gyms we get more support tasks: subscription problems, membership updates, data exports... The way we solve it now is a "tree of llms.txt". An llms.txt normally references what info is available on a website or docs — we use the same idea internally to organize the info the agent needs. The agent starts from a folder and navigates down: .
├── llms.txt # references each folder at this level
├── stripe/ # info.md: how our stripe account is structured
├── firestore/ # info.md: how the schema looks
└── support/
├── info.md # how to resolve support tasks
├── runbooks/ # one file per task, with its own llms.txt
│ ├── cancel-subscription.md
│ ├── export-gym-data.md
│ └── fix-membership-mismatch.md
└── logs/ # one file per day, every task the agent resolved
Copy
With this we can steer the agent much better and create a new runbook every time a new support task comes. You can add in every integration what the agent should and should not touch. The gcontext prompts make sure any guardrail is meticulously followed.

GitHub - bleak-ai/gcontext: Context Management System to let your agent find and build relevant context · GitHub
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
bleak-ai
/
gcontext
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits case-studies/ maat-support case-studies/ maat-support core core demo demo examples examples gcontext gcontext tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
The main concept is a tree of llms.txt that references either folders or files. The goal is to load in a conversation with an AI agent the right information at every time, and from this, be able to grow the context that the agent has access to and create a something analogous to a "Live Memory" or Live Record of things.
What is the difference from Claude Code Memory system?
The problem is the same, but Gcontext is for users that want to treat the Context as a problem on it's own, while Memory handles everything behind the scenes, Gcontext allows the user to define how Context should be defined and be much more precise in which are the steps that an AI agent does.
1. Install and create the workspace.
curl -LsSf https://gcontext.ai/gcontext/install.sh | sh # or: uv tool install gcontext-ai
gcontext init
2. Ask your agent to build a module. Open your agent in the workspace and name the integration you want:
Create a supabase integration module and load it.
The agent writes info.md , an llms.txt index, and a module.yaml that declares which secrets it needs by name only (the values never leave .env ).
3. Fill in the secrets it asks for. The module tells you which variables to set; put them in .env :
SUPABASE_URL=https://....supabase.co
SUPABASE_SECRET_KEY=...
4. Enrich the module with real data. Now that the access is in place, ask the agent to explore the live service and write down what it finds:
Look in supabase and update the module information with real read requests.
From then on a fresh session can answer questions ("how many rows are in the members table?") by following the index to the module and calling the API with the key from .env .
Repeat step 2 for each new service, and a tree starts to form: a root llms.txt routing to one module per integration, each declaring its own keys. Just ask:
Create a stripe integration module and load it.
The more integrations you add, the more the workspace becomes a single place that coordinates access to all of them, which is exactly when gcontext pays off.
This is not for everyone. If you work in a project where the only external dependency is a database, it probably doesn't make sense to set all of this up just for this. But as more external dependencies and integrations you are using, it makes more sense to have a dedicated central place where to coordinate the access to all of these.
This project has been developed at the startup MAAT , Management Software for Martial Arts Gyms, where the product is built on top of a payment System and a DB. As we got more gyms we had to spend more and more time resolving support tasks where all of these integrations were affected. We came up with the pattern that gcontext follows with the llms.txt to speed up our day to day.
Our Journey solving Support Tasks
Before AI. We had some playbooks on how to resolve the Support Tasks and we used them to solve the tasks manually
With AI - Mainly CLAUDE.md. We explained in the Claude.md how our system works and which were the most common causes of problems, it gave us most of the times the right indications, but often we still had to do manual work and exploration for many tasks, still repetitive as we didn't had a place to store the runbooks.
With AI + GContext This allowed us to intertwine completely the AI in our processes. When we do now an exploration for a task, we can save this exploration in a md file, referenced by an llms.txt and the AI model will be able to find it later.
You can browse an example of it here: case-studies/maat-support/ . It is the tree from the diagram above: a root llms.txt router over the stripe and firestore integration modules, plus a support module with its runbooks and execution logs.
The rest the agent does for you from plain language, just ask:
Built by Bleak AI | gcontext.ai
Context Management System to let your agent find and build relevant context
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
