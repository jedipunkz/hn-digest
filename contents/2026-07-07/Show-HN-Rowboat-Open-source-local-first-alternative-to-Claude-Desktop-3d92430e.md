---
source: "https://github.com/rowboatlabs/rowboat"
hn_url: "https://news.ycombinator.com/item?id=48819808"
title: "Show HN: Rowboat – Open-source, local-first alternative to Claude Desktop"
article_title: "GitHub - rowboatlabs/rowboat: Open-source AI coworker, with memory · GitHub"
author: "segmenta"
captured_at: "2026-07-07T17:07:47Z"
capture_tool: "hn-digest"
hn_id: 48819808
score: 5
comments: 0
posted_at: "2026-07-07T16:10:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rowboat – Open-source, local-first alternative to Claude Desktop

- HN: [48819808](https://news.ycombinator.com/item?id=48819808)
- Source: [github.com](https://github.com/rowboatlabs/rowboat)
- Score: 5
- Comments: 0
- Posted: 2026-07-07T16:10:44Z

## Translation

タイトル: Show HN: Rowboat – クロード デスクトップに代わるオープンソース、ローカルファーストの代替品
記事タイトル: GitHub - rowboatlabs/rowboat: メモリを備えたオープンソース AI コワーカー · GitHub
説明: メモリを備えたオープンソースの AI コワーカー。 GitHub でアカウントを作成して、Rowboatlabs/Rowboat の開発に貢献してください。
HN テキスト: こんにちは、HN。Claude のデスクトップ アプリは素晴らしいです。しかし、私たち自身の日常業務では、チャット アプリではなく、もっと本格的な仕事アプリのようなものにしたいと考え続けました。 Rowboat は、Rowboat 内に独自の作業面を構築する機能を含む、そのための私たちの試みです (詳細は後述)。私たちのリポジトリは https://github.com/rowboatlabs/rowboat で、デモビデオはこちらにあります: https://www.youtube.com/watch?v=et5yQABJ3xI 以前のスタートアップでは、P&G ブランドをサポートするチームを含むエンタープライズ サポート担当者向けの深層学習製品を構築しました。モデルは、サポート担当者が電話をかけたり電子メールを処理したりしている間に、ライブメモ、提案された返信、推奨されるアクションを記録しました。教訓の 1 つが私たちに心に残りました。AI が正しいだけでは十分ではなく、作業が行われている場所に助けが現れる必要があります。そこで、私たちは「作業面」と呼ぶものを追加しました。これは、電子メール、会議、メモ、ブラウザ、および並列コーディング用の専用領域であり、アシスタントはチャットだけでなく、ワークフロー自体の内部で支援できます。 - 電子メール クライアント: Rowboat には、受信電子メールを重要なものとそれ以外のものに分類し、重要な電子メールの下書きを事前に作成するシンプルな電子メール クライアントがあります。メールを編集して送信すると、あなたのスタイルがメモされるので、今後の下書きはあなたの意見に近づきます。 - 会議メモ: グラノーラ スタイルのローカル会議メモ作成ツールを構築しました。メモはプレーンな Markdown ファイルとしてマシンに保存されます。会議後、Rowboat はメモをナレッジ グラフにフィードバックし、関連する人、プロジェクト、トピックを更新します。

c ノート。 - ブラウザ: メインのブラウザとは別に、アシスタントにサポートしてもらいたいアカウントにのみログインできる組み込みブラウザを追加しました。アシスタントはブラウザ使用スキルを使用して Web サイトをナビゲートします。 - 並列コーディング: Rowboat 内のコード モードを使用すると、Claude Code または Codex の複数のインスタンスをスピンして、それらを直接操作するか、Rowboat に作業コンテキストを使用してそれらを調整させることができます。このために、Rowboat で ACP (Agent Client Protocol) クライアントを構築しました。 - 注: 手漕ぎボートには、黒曜石スタイルのローカルメモ取りシステムが搭載されています。グラフビュー、ベースビュー、音声メモが付属しています。 Google ドキュメント ファイルを同期して、Rowboat 内で編集することもできます。 Rowboat (Web アプリ) 内に独自の作業面を構築することもできます。各アプリは独自の UI とバックグラウンド エージェントを取得し、Rowboat のすべてのツール、製品統合、およびワーク メモリを使用できます。例: GitHub アクティビティ、プロジェクト追跡、広告キャンペーン管理を管理するアプリ。起動時には、検索してインストールできるコミュニティ アプリがいくつかあります。また、GitHub リポジトリを作成して登録することで、独自のアプリを公開することもできます。 Rowboat は、作業内容をナレッジ グラフにインデックス付けします。ナレッジ グラフは、より適切なコンテキストを得るために上記のすべてのサーフェスで使用されます。私たちは数か月前にこれに関して HN を表示しました: https://news.ycombinator.com/item?id=46962641 。これらのいくつかを結び付ける例として、電子メール、会議、Slack から機能リクエストを収集してランク付けするアプリを Rowboat 内に作成し、クロード コードを使用してトップランクの機能の最初のバージョンをドラフトし、ナレッジ グラフからその機能に関する以前のコンテキストを取得できます。 Rowboat はローカルファーストです。データはプレーンな Markdown ファイルとして保存され、いつでも読み取り、編集、削除できます。これは Apache-2.0 であり、Ollama を介したローカル モデルを含むあらゆる LLM で動作します。

r LMスタジオ。皆様のご意見をお待ちしております。貢献も大歓迎です。

記事本文:
GitHub - rowboatlabs/rowboat: メモリを備えたオープンソース AI コワーカー · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
手漕ぎボートラボ
/
手漕ぎボート
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,934 コミット 1,934 コミット .github/ workflows .github/ workflows apps apps 資産 アセット .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile.qdrant Dockerfile.qdrant ライセンス ライセンスREADME.md README.md build-electron.sh build-electron.sh docker-compose.yml docker-compose.yml google-setup.md google-setup.md start.sh start.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたの作業のメモリとそれに基づいて動作する内蔵サーフェスを備えたデスクトップ AI の同僚。
Rowboat は、作業のインデックスを生きたナレッジ グラフに作成し、それを使用してマシン上で作業を実行します。これには、AI とコラボレーションするための作業面 (電子メール クライアント、メモ、ブラウザ、コード モード、会議メモ作成者、さまざまなプロジェクト用のワークスペース) が含まれています。
Mac/Windows/Linux 用の最新版をダウンロード: ダウンロード
デモ - アプリをコード化する · デモ - ナレッジ グラフ
⭐ Rowboat が役立つと思われる場合は、リポジトリにスターを付けてください。より多くの人が見つけやすくなります。
Rowboat は、電子メール、会議、Slack、アシスタントの会話を生きた Obsidian スタイルのバックリンクされたナレッジ グラフにインデックス付けします。
電子メール
内蔵の電子メール クライアントは、電子メールを重要なものとその他のものに分類します。 Rowboat は、すべての作業コンテキストを使用して、重要な電子メールに対する返信の下書きを自動的に作成します。
バックグラウンドエージェント
新着メールなどのイベント時に実行したり、毎日午前 8 時に実行するなどのスケジュールで実行するバックグラウンド エージェントを設定できます。ツールに接続し、Web を検索し、ブラウザを使用し、Claude Code または Codex を使用してコードを作成できます。
内蔵ブラウザ
Rowboat には、ユーザーとアシスタントが Web タスクで共同作業できるブラウザが含まれています。メインのブラウザから分離されているため、アシスタントにアクセスを許可するアカウントにのみログインできます。
会議はありません

エス
マイクとスピーカーを利用してライブトランスクリプトを作成し、マークダウンファイルに会議を要約してナレッジグラフを更新するローカル会議のメモテイカー。
コードモード
コード モードでは、Claude Code または Codex を使用して並列コーディング エージェントを起動し、Rowboat で必要に応じてすべての作業コンテキストを使用してそれらを駆動させることができます。
アプリ
Rowboat 内に独自の作業面を構築できます。すべてのツールと統合にアクセスでき、それらを他の人と共有できます。
統合
最も人気のある製品へのワンクリック統合が含まれます。
インストール
Mac/Windows/Linux 用の最新版をダウンロード: ダウンロード
すべてのリリース ファイル: https://github.com/rowboatlabs/rowboat/releases/latest
Google サービス (Gmail、カレンダー、ドライブ) に接続するには、Google のセットアップに従ってください。
音声入力と音声メモ (オプション) を有効にするには、~/.rowboat/config/deepgram.json に Deepgram API キーを追加します。
音声出力 (オプション) を有効にするには、~/.rowboat/config/elevenlabs.json に イレブンラボ API キーを追加します。
Exa Research Search (オプション) を使用するには、~/.rowboat/config/exa-search.json に Exa API キーを追加します。
外部ツール (オプション) を有効にするには、MCP サーバーを追加するか、~/.rowboat/config/composio.json に API キーを追加して Composio ツールを使用します。
すべての API キー ファイルは同じ形式を使用します。
{
"apiKey": "<キー>"
}
どう違うのか
ほとんどの AI ツールは、トランスクリプトやドキュメントを検索することで、オンデマンドでコンテキストを再構築します。
Rowboat は代わりに、長期にわたる知識を維持します。
関係が明示的で検査可能である
メモはユーザーが編集可能であり、モデル内に隠されることはありません
すべてはプレーンなマークダウンとしてマシン上に存在します
その結果、毎回コールド状態で始まる検索ではなく、記憶が強化されます。
Rowboat は、お好みのモデル設定で動作します。
Ollama または LM Studio を介したローカル モデル
ホストされたモデル (独自の API キー/プロバイダーを使用する)
いつでもモデルを交換

e — データはローカルの Markdown ボールトに残ります
ツールを使用して手漕ぎボートを拡張する (MCP)
Rowboat は、Model Context Protocol (MCP) 経由で外部ツールやサービスに接続できます。
つまり、(たとえば) 検索、データベース、CRM、サポート ツール、自動化、または独自の内部ツールを接続できるということです。
例: Exa (Web 検索)、Twitter/X、イレブンラボ (音声)、Slack、Linear/Jira、GitHub など。
すべてのデータはプレーンなマークダウンとしてローカルに保存されます
独自のフォーマットやホストされたロックインはありません
いつでもすべてを検査、編集、バックアップ、削除できます。
メモリを備えたオープンソースの AI コワーカー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1.5k
フォーク
レポートリポジトリ
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source AI coworker, with memory. Contribute to rowboatlabs/rowboat development by creating an account on GitHub.

Hi HN, Claude’s desktop app is brilliant. But for our own daily work we kept wanting it to be less like a chat app and more like a full-fledged work app. Rowboat is our attempt at that, including the ability to build your own work surfaces inside Rowboat (more below). Our repo is https://github.com/rowboatlabs/rowboat , and there’s a demo video here: https://www.youtube.com/watch?v=et5yQABJ3xI In a previous startup, we built a deep-learning product for enterprise support reps, including teams supporting P&G brands. Models took live notes, suggested replies, and recommended actions while support reps were on calls or handling emails. One lesson stuck with us: it's not enough for the AI to be right, the help has to show up where the work is happening. So we added what we came to call “work surfaces”: dedicated areas for email, meetings, notes, browser, and parallel coding, where the assistant can help inside the workflow itself rather than only through chat: - Email client: Rowboat has a simple email client that sorts incoming emails into important vs. everything else, and pre-creates drafts for important emails. As you edit and send emails, it takes notes on your style, so future drafts get closer to your voice. - Meeting notes: We built a Granola-style local meeting notetaker. Notes are stored as plain Markdown files on your machine. After a meeting, Rowboat feeds the notes back into the knowledge graph and updates the relevant people, project, and topic notes. - Browser: We added a built-in browser, isolated from your main one, where you can log in only to the accounts you want the assistant to help with. The assistant uses browser-use skills to navigate websites. - Parallel coding: The code-mode inside Rowboat lets you spin multiple instances of Claude Code or Codex and either work with them directly or let Rowboat use your work context to orchestrate them. We built an ACP (Agent Client Protocol) client in Rowboat for this. - Notes: Rowboat has an Obsidian-style local note-taking system. It comes with graph view, bases view, and voice notes. You can also sync Google Docs files and edit them inside Rowboat. You can also build your own work surfaces inside Rowboat (web apps). Each app gets its own UI and a background agent, and can use all of Rowboat's tools, product integrations, and your work memory. For instance: an app to manage GitHub activity, project tracking, or ads campaign management. There are a few community apps at launch you can search and install, and you can publish your own by creating a GitHub repo for it and registering it. Rowboat also indexes your work into a knowledge graph that all of the above surfaces use to have better context. We did a Show HN a few months back on this: https://news.ycombinator.com/item?id=46962641 . As an example that ties some of these together: you can create an app inside Rowboat that collects feature requests from your email, meetings, and Slack and ranks them, then uses Claude Code to draft a first version of the top-ranked feature, pulling prior context about it from your knowledge graph. Rowboat is local-first: data is stored as plain Markdown files you can read, edit, or delete anytime. It is Apache-2.0 and works with any LLM, including local models through Ollama or LM Studio. We’d love to hear your thoughts, and contributions are welcome!

GitHub - rowboatlabs/rowboat: Open-source AI coworker, with memory · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
rowboatlabs
/
rowboat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,934 Commits 1,934 Commits .github/ workflows .github/ workflows apps apps assets assets .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile.qdrant Dockerfile.qdrant LICENSE LICENSE README.md README.md build-electron.sh build-electron.sh docker-compose.yml docker-compose.yml google-setup.md google-setup.md start.sh start.sh View all files Repository files navigation
A desktop AI coworker with a memory of your work and built-in surfaces to act on it.
Rowboat indexes your work into a living knowledge graph and uses that to get work done on your machine. It includes work surfaces for collaborating with AI: email client, notes, browser, code mode, meeting note taker, and workspaces for different projects.
Download latest for Mac/Windows/Linux: Download
Demo - apps to code · Demo - knowledge graph
⭐ If you find Rowboat useful, please star the repo. It helps more people find it.
Rowboat indexes email, meetings, slack and assistant conversations into a living Obsidian-style backlinked knowledge graph.
Email
The built-in email client sorts emails into important and everything else. Rowboat automatically drafts responses for important email using all the work context.
Background agents
You can set up background agents that run on events like new email or on schedule like every day at 8am. They can connect to tools, search the web, use the browser and write code using Claude Code or Codex.
Built-in Browser
Rowboat includes a browser that lets you and assistant collaborate on web tasks. Because its isolated from your main browser, you can log in only to the accounts that want the assistant to access.
Meeting Notes
A local meeting note-taker that taps into mic & speaker, produces live transcript and summarizes the meeting in a markdown file and updates the knowledge graph.
Code Mode
Code mode lets you spin up parallel coding agents with Claude Code or Codex, and have Rowboat drive them with all the work context where needed.
Apps
You can bulild your own work surfaces inside Rowboat — they get acess to all the tools and integrations, and you can share them with other people.
Integrations
Includes one-click integrations to most popular products.
Installation
Download latest for Mac/Windows/Linux: Download
All release files: https://github.com/rowboatlabs/rowboat/releases/latest
To connect Google services (Gmail, Calendar, and Drive), follow Google setup .
To enable voice input and voice notes (optional), add a Deepgram API key in ~/.rowboat/config/deepgram.json
To enable voice output (optional), add an ElevenLabs API key in ~/.rowboat/config/elevenlabs.json
To use Exa research search (optional), add the Exa API key in ~/.rowboat/config/exa-search.json
To enable external tools (optional), you can add any MCP server or use Composio tools by adding an API key in ~/.rowboat/config/composio.json
All API key files use the same format:
{
"apiKey": "<key>"
}
How it’s different
Most AI tools reconstruct context on demand by searching transcripts or documents.
Rowboat maintains long-lived knowledge instead:
relationships are explicit and inspectable
notes are editable by you, not hidden inside a model
everything lives on your machine as plain Markdown
The result is memory that compounds, rather than retrieval that starts cold every time.
Rowboat works with the model setup you prefer:
Local models via Ollama or LM Studio
Hosted models (bring your own API key/provider)
Swap models anytime — your data stays in your local Markdown vault
Extend Rowboat with tools (MCP)
Rowboat can connect to external tools and services via Model Context Protocol (MCP) .
That means you can plug in (for example) search, databases, CRMs, support tools, and automations - or your own internal tools.
Examples: Exa (web search), Twitter/X, ElevenLabs (voice), Slack, Linear/Jira, GitHub, and more.
All data is stored locally as plain Markdown
No proprietary formats or hosted lock-in
You can inspect, edit, back up, or delete everything at any time
Open-source AI coworker, with memory
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1.5k
forks
Report repository
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
