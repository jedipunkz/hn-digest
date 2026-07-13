---
source: "https://github.com/ieranama/discomcp"
hn_url: "https://news.ycombinator.com/item?id=48890786"
title: "DiscoMCP – Turn an unknown MCP into a reusable operational skill for AI agents"
article_title: "GitHub - ieranama/discomcp: Teach any AI agent how you use an MCP server – Soft-Landing to your AI integrations · GitHub"
author: "inigoerana"
captured_at: "2026-07-13T11:53:02Z"
capture_tool: "hn-digest"
hn_id: 48890786
score: 1
comments: 0
posted_at: "2026-07-13T10:58:42Z"
tags:
  - hacker-news
  - translated
---

# DiscoMCP – Turn an unknown MCP into a reusable operational skill for AI agents

- HN: [48890786](https://news.ycombinator.com/item?id=48890786)
- Source: [github.com](https://github.com/ieranama/discomcp)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T10:58:42Z

## Translation

タイトル: DiscoMCP – 未知の MCP を AI エージェントの再利用可能な運用スキルに変える
記事のタイトル: GitHub - ieranama/discomcp: AI エージェントに MCP サーバーの使用方法を教える – AI 統合へのソフトランディング · GitHub
説明: AI エージェントに MCP サーバーの使用方法を教える – AI 統合へのソフトランディング – ieranama/discomcp

記事本文:
GitHub - ieranama/discomcp: AI エージェントに MCP サーバーの使用方法を教える – AI 統合へのソフトランディング · GitHub
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
イエラナマ
/
ディスコンプ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット .github .github ベンチマーク ベンチマーク クレート クレート ドキュメント ドキュメントの例 例 .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md dist-workspace.toml dist-workspace.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントに、接続先のツールの使用方法を 1 つのコマンドで安全に教えます。
エージェントはコードが得意です。彼らはツールの中で迷子になり、100 のアクションのうちどれが重要なのか、データがどのように接続されているのか、何を触れても安全なのかがわかりません。それで彼らは推測するか、固まってしまいます。
DiscoMCP はそれを修正します。これを任意の MCP サーバーに向けると、一般的なツール リストではなく、データのワークフロー、つまり実際に作業するビュー、テーブル、レコード、本当の質問に答えるシーケンス、ある結果が次の結果につながる方法など、その使用方法に応じたスキルがエージェントに渡されます。自分のワークスペースを見て、何も変更することなく学習できます。
ツール カタログは、サーバーに 90 のアクションがあることをエージェントに伝えます。 DiscoMCP は、実際に存在するデータに対して、実際に使用する 5 つを使用する順序で通知します。
🎯 あなたの働き方に合わせてカスタマイズ
一般的な機能ダンプではなく、実際のデータの内容に基づいた、ワークフロー、規則、ワークスペースの重要な部分など、使用状況のプロファイルです。
🧭 道を熟知したエージェント
エージェントは、100 個のツールを推測するのではなく、実際の質問に答えるシーケンスに従い、1 つの結果を次の結果に連鎖させます。
🔒 読み取り専用、

保証された
探索しても決して書くことはできません。ステップが読み取りであることが証明できる場合にのみ、ステップを実行します。学習中に何も削除されたり変更されたりすることはありません。
⚡ 1 つのコマンドでセットアップ不要
単一の 8 MB バイナリ。ランタイムもツールチェーンもありません。npx だけで実行できます。
読み取り専用、保証あり
これは、実際のシステム上で実際にエージェントを解放できるようにする部分です。
DiscoMCP は、デフォルト拒否ゲートの背後を探索します。アクションが読み取りであることが証明できる場合にのみアクションを実行します。安全な検索が実行されます。たとえツールが無害であると主張していても、書き込み、変更、削除の可能性のあるものはすべて拒否されます。秘密は、保存されているものすべてから削除されます。
そのため、エージェントは息を止めずに制作ツールを学習できます。
同じ質問、同じサーバー、同じモデル - 生成されたスキルの有無にかかわらず。非常に広く、馴染みのないサーバーに対する読み取り専用。行ごとに n=2。
タスクが難しくて慣れていないほど、スキルは役に立ちます。このスキルは、コールド エージェントがトライアルで再発見しなければならないマップをフロントローディングします。どちらも正解に到達します。スキルははるかに少ないラウンドトリップで正解に到達します。
サンプルが少なく、方向性があるため、保証はありません。些細なタスクや狭いサーバーでは、スキル自体のプロンプトコストが節約した分を洗い流してしまう可能性があります。永続的な利点は、ラウンドトリップが減り、複雑なサーバー上での動作が安定することです。完全なメソッドは Benchmarks/METRICS.md にあります。
1. 実行します。インストールは必要ありません。
npx @ieranama/discomcp --help
2. サーバーを指定します。構成全体は数行です。
[ターゲット。例]
トランスポート = "stdio"
コマンド = " npx "
args = [ " -y " , " some-mcp-server " ]
3. それをエージェントに渡して、次のことを調べてもらいます。
discomcpserve --config ./discomcp.toml
あなたのエージェントが調査を行います。 DiscoMCP はそれを安全に保ち、スキルを書き込みます。結果は .discomcp/profiles/<server>/SKILL.md に保存され、エージェントにドロップする準備が整います。

。
Rust で構築: モデルが思考を行い、小さな決定論的なコアがあらゆる安全性チェックを強制します。生成されたスキル内のすべてのクレームには、それがどのように知られたか (宣言、文書化、観察、推論) がタグ付けされるため、エージェントが推測を事実と間違えることはありません。
貢献 · セキュリティ ポリシー
オプションで、Apache ライセンス、バージョン 2.0、または MIT ライセンスのいずれかに基づいてライセンスが付与されます。
お客様が明示的に別段の定めをしない限り、Apache-2.0 ライセンスで定義されているように、お客様が DiscoMCP に含めるために意図的に提出した投稿は、追加の条項や条件なしで上記のようにデュアルライセンスされるものとします。
AI エージェントに MCP サーバーの使用方法を教える - AI 統合へのソフトランディング
Apache-2.0、MITライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.5.1
最新の
2026 年 7 月 13 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Teach any AI agent how you use an MCP server – Soft-Landing to your AI integrations - ieranama/discomcp

GitHub - ieranama/discomcp: Teach any AI agent how you use an MCP server – Soft-Landing to your AI integrations · GitHub
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
ieranama
/
discomcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits .github .github benchmarks benchmarks crates crates docs docs examples examples .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md dist-workspace.toml dist-workspace.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Teach any AI agent how you use the tools it connects to — safely, in one command.
Agents are great at code. They get lost inside your tools — they don't know which of a hundred actions matter, how your data connects, or what's safe to touch. So they guess, or they freeze.
DiscoMCP fixes that. Point it at any MCP server and it hands your agent a skill for how you use it — not a generic tool list, but your workflows on your data: the views, tables and records you actually work with, the sequences that answer your real questions, and how one result leads into the next. Learned by looking at your own workspace, and without ever changing a thing .
A tool catalogue tells an agent the server has 90 actions. DiscoMCP tells it the five you actually use, in the order you use them, against the data that's really there.
🎯 Tailored to how you work
Not a generic capability dump — a profile of your usage: your workflows, your conventions, the parts of your workspace that matter, grounded in what's really in your data.
🧭 Agents that know their way around
Your agent follows the sequences that answer real questions and chains one result into the next, instead of guessing across a hundred tools.
🔒 Read-only, guaranteed
Exploring can never write. It runs a step only when it can prove that step is a read. Nothing is deleted or modified while it learns.
⚡ One command, zero setup
A single 8 MB binary. No runtime, no toolchain — npx and you're running.
Read-only, guaranteed
This is the part that lets you actually turn an agent loose on a real system.
DiscoMCP explores behind a default-deny gate : it runs an action only when it can prove that action is a read. A safe lookup runs. Anything that could write, change, or delete — even if a tool claims it's harmless — is refused. Secrets are stripped from everything it saves.
So your agent can learn your production tools without you holding your breath.
Same question, same server, same model — with and without the generated skill. Read-only, against a genuinely wide, unfamiliar server. n=2 per row.
The harder and less familiar the task, the more it helps: the skill front-loads the map a cold agent has to rediscover by trial. Both reach correct answers — the skill reaches them in far fewer round-trips.
Small sample, directional — not a guarantee. On a trivial task or a narrow server the skill's own prompt cost can wash out what it saves; the durable win is fewer round-trips and steadier behavior on complex servers. Full method in benchmarks/METRICS.md .
1. Run it — no install needed:
npx @ieranama/discomcp --help
2. Point it at a server — the whole config is a few lines:
[ targets . example ]
transport = " stdio "
command = " npx "
args = [ " -y " , " some-mcp-server " ]
3. Hand it to your agent and let it explore:
discomcp serve --config ./discomcp.toml
Your agent does the exploring; DiscoMCP keeps it safe and writes the skill. The result lands in .discomcp/profiles/<server>/SKILL.md — ready to drop into your agent.
Built in Rust: the model does the thinking, a small deterministic core enforces every safety check. Every claim in a generated skill is tagged with how it was known — declared, documented, observed, or inferred — so an agent never mistakes a guess for a fact.
Contributing · Security policy
Licensed under either of Apache License, Version 2.0 or MIT License at your option.
Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in DiscoMCP by you, as defined in the Apache-2.0 license, shall be dual-licensed as above, without any additional terms or conditions.
Teach any AI agent how you use an MCP server – Soft-Landing to your AI integrations
Apache-2.0, MIT licenses found
Licenses found
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.5.1
Latest
Jul 13, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
