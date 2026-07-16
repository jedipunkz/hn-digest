---
source: "https://github.com/yeme-oss/quatuor_ai"
hn_url: "https://news.ycombinator.com/item?id=48933720"
title: "Show HN: Quatuor – Kick back and watch 4 agents LLM talk to each other (FOSS)"
article_title: "GitHub - yeme-oss/quatuor_ai: AI conversation simulator: 4 AI characters talk autonomously, orchestrated by an invisible Director agent · GitHub"
author: "ycosynot"
captured_at: "2026-07-16T13:21:22Z"
capture_tool: "hn-digest"
hn_id: 48933720
score: 1
comments: 0
posted_at: "2026-07-16T12:46:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Quatuor – Kick back and watch 4 agents LLM talk to each other (FOSS)

- HN: [48933720](https://news.ycombinator.com/item?id=48933720)
- Source: [github.com](https://github.com/yeme-oss/quatuor_ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T12:46:14Z

## Translation

タイトル: 番組 HN: Quatuor – リラックスして 4 人のエージェント LLM が互いに会話するのを見てください (FOSS)
記事タイトル: GitHub - yeme-oss/quatuor_ai: AI 会話シミュレーター: 4 人の AI キャラクターが自律的に会話し、目に見えない Director エージェントによって調整される · GitHub
説明: AI 会話シミュレーター: 4 人の AI キャラクターが自律的に会話し、目に見えない Director エージェントによって調整されます - yeme-oss/quatuor_ai
HN テキスト: QUATUOR へようこそ。自由にフォークしたりスターを付けたりして、素晴らしいことをしてください。これを起点に、あるいはこのアイデアから、素敵なものがたくさん作れると思います。私は何も売ることができなかったので、今は無料でオープンソースでリリースしています。またね

記事本文:
GitHub - yeme-oss/quatuor_ai: AI 会話シミュレーター: 4 人の AI キャラクターが自律的に会話し、目に見えない Director エージェントによって調整 · GitHub
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
イエメオス
/
quatuor_ai
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .env.example .env.example .gitignore .gitignore README.md README.md i18n.js i18n.jsindex.cssindex.cssindex.htmlindex.htmlmain.jsmain.jspackage-lock.jsonpackage-lock.jsonpackage.jsonpackage.jsonpersonas.jspersonas.js quat1.png quat1.png quat2.png quat2.png quat3.png quat3.png server.cjs server.cjs vite.config.js vite.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI キャラクターは 4 人。目に見えない監督が一人。人間の介入はゼロです。
あなたが座って見ている間、4 人の AI ペルソナがリアルタイムで議論し、冒険し、お互いに中断する自律型会話シミュレーターです。
パーティーは古代の赤いドラゴンに遭遇します。ゲーム マスターが場面を設定し、3 人の冒険者がそれぞれの個性に忠実に脱出方法を議論します。
Quatuor は、それぞれ独自のシステム プロンプトによって駆動される 4 人の AI キャラクター間のライブ グループ会話を演出します。シナリオを選択して Start を押すと、ループが永久に実行されます。
🎬 目に見えない第 5 エージェント (ディレクター) がトランスクリプトを読み、次に誰が話すか、そして彼らが普通に話すか、誰かを文章の途中で遮るかを決定します。
🗣️ 選択されたキャラクターは、共有されたトランスクリプトを受け取り、そのキャラクターに合わせて次の行を即興で作成します。
🔁 メッセージがチャットに到達し、ディレクターが再度相談を受けます。永遠に。
ユーザーの手番はありません。あなたは聴衆です。テーブルトップセッションや、それ自体が書き込まれるポッドキャストを聞いているようなものです。
フローチャート LR
T[(共有トランスクリプト)] --> D{{🎬 ディレクター}}
D -- "next_speaker + 続行/中断" --> C[🗣️ キャラクター LLM 呼び出し]
C -- 改行 --> T
読み込み中
内部的には、1 つの LLM が異なるシステム プロンプトで N 回呼び出されます。永続的なセッションはありません。

ns、共有されたトランスクリプトが次に話す人に再生されるだけです。ディレクターが有効な JSON に応答できなかった場合でも、ラウンドロビン フォールバックによってショーの実行が継続されます。
3 つのプリセットが箱から出して出荷され、それぞれに 4 つの手動調整されたペルソナが含まれています。
王子は同人たちをエリジウムに招集した。流出したビデオがマスカレードを脅かす。ヴィクトリアは影響力を欲しがり、ダンテは警官を揺さぶりたがり、ルナはなぞなぞを話し、それが正しいことが判明する。
すべてのペルソナは、停滞防止のルールに従って書かれています。つまり、キャラクターは結果に反応し、失敗したアクションを決して繰り返さず、ループではなくストーリーを動かし続けなければなりません。
ノードがインストールされていませんか?最新リリースから quatuor.exe を取得します。実行すると、ブラウザで開きます。 OpenRouter キーを .env ファイルの exe の隣に配置します (またはアプリのサイドバーに貼り付けます)。
git クローン https://github.com/yeme-oss/quatuor_ai.git
cd quatuor_ai
npmインストール
cp .env.example .env # OpenRouter API キーを貼り付けます
npm run dev # → http://localhost:3000
OpenRouter API キーが必要です。これを .env ( OPENROUTER_API_KEY=sk-or-v1-... ) に置くか、アプリの構成サイドバーに直接貼り付けます。ローカルの開発サーバーを介してプロキシされ、クライアント バンドルに組み込まれることはありません。
💸 注: Quatuor 自体は無料でオープンソースですが、LLM 呼び出しはそうではありません。各メッセージは OpenRouter アカウントのクレジットを消費し、選択したモデルによってトークンごとに請求されます。会話は永遠に続くため、使用状況に注意してください (OpenRouter の一部のモデルは、費用をかけずに視聴したい場合は無料枠です)。
どの OpenRouter モデルでも機能します。上部バーの「モデル」フィールドにその ID を入力します。
すべては⚙️設定サイドバーからライブで編集可能です:
🎯 トピック — 最初の状況を書き直し、同じキャストが何か新しいことを即興で演じるのを見てください。
👥 キャラクター — 名前を変更したり、録音したりできます

バブルを着色し、システム プロンプトを完全に書き換えます。
🎬 ディレクター プロンプト — キャスト ロジック自体を変更します (中断頻度、ペーシング ルールなど)。
⏩ スピード — 5 秒のゆっくりとした書き込みから 1.5 秒の速射の交換まで。
ポッドキャストのプリセットが本格化します。アーサーは証拠を要求し、チャールズは記録を修正し、タイピング インジケーターはディレクターが次に誰を選んだかを示します。
ヘッダーの FR/EN ボタンは、インターフェイス、ペルソナ、生成されたダイアログなど、エクスペリエンス全体を英語とフランス語で切り替えます。
§──index.html # シングルページ UI
§──index.css # すべてのスタイリング
§── main.js # シミュレーションループ、ディレクターロジック、DOMレンダリング
§── personas.js # シナリオプリセット + ディレクタープロンプト (英語およびフランス語)
§── i18n.js # UI 文字列とランタイム プロンプト テンプレート (英語とフランス語)
└── vite.config.js # Dev サーバー + /api/chat プロキシから OpenRouter
週末のプロトタイプ — 意図的にKISS。優先事項: 建築ではなく、見るのが楽しい。 🍿
AI 会話シミュレーター: 4 人の AI キャラクターが自律的に会話し、目に見えない Director エージェントによって調整されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
クアチュール v1.1.0
最新の
2026 年 7 月 16 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI conversation simulator: 4 AI characters talk autonomously, orchestrated by an invisible Director agent - yeme-oss/quatuor_ai

Welcome to QUATUOR. Feel free to fork and star and be awesome. I think you can make a lot of cool stuff starting from this, or from this idea. I can never manage to sell anything and so now I release for free & open source. See you around

GitHub - yeme-oss/quatuor_ai: AI conversation simulator: 4 AI characters talk autonomously, orchestrated by an invisible Director agent · GitHub
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
yeme-oss
/
quatuor_ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .env.example .env.example .gitignore .gitignore README.md README.md i18n.js i18n.js index.css index.css index.html index.html main.js main.js package-lock.json package-lock.json package.json package.json personas.js personas.js quat1.png quat1.png quat2.png quat2.png quat3.png quat3.png server.cjs server.cjs vite.config.js vite.config.js View all files Repository files navigation
Four AI characters. One invisible director. Zero human intervention.
An autonomous conversation simulator where four AI personas debate, adventure and interrupt each other in real time — while you sit back and watch.
The party meets an ancient red dragon — the Game Master sets the scene, and the three adventurers argue their way out, each true to their own personality.
Quatuor stages a live group conversation between four AI characters , each driven by its own system prompt. You pick a scenario, press Start , and the loop runs forever:
🎬 An invisible fifth agent — the Director — reads the transcript and decides who speaks next , and whether they speak normally or cut someone off mid-sentence .
🗣️ The chosen character receives the shared transcript and improvises its next line, in character.
🔁 The message lands in the chat, and the Director is consulted again. Forever.
There is no user turn. You're the audience — like overhearing a tabletop session or a podcast that writes itself.
flowchart LR
T[(Shared transcript)] --> D{{🎬 Director}}
D -- "next_speaker + continue/interrupt" --> C[🗣️ Character LLM call]
C -- new line --> T
Loading
Under the hood it's one single LLM called N times with different system prompts — no persistent sessions, just a shared transcript replayed to whoever speaks next. If the Director ever fails to answer valid JSON, a round-robin fallback keeps the show running.
Three presets ship out of the box, each with four hand-tuned personas:
The Prince has summoned the coterie to Elysium: a leaked video threatens the Masquerade. Victoria wants leverage, Dante wants to shake down cops, and Luna speaks in riddles that turn out to be right.
Every persona is written with an anti-stagnation rule : characters must react to consequences, never repeat a failed action, and keep the story moving instead of looping.
No Node installed? Grab quatuor.exe from the latest release — run it, and it opens in your browser. Put your OpenRouter key in a .env file next to the exe (or paste it in the app's sidebar).
git clone https://github.com/yeme-oss/quatuor_ai.git
cd quatuor_ai
npm install
cp .env.example .env # then paste your OpenRouter API key
npm run dev # → http://localhost:3000
You need an OpenRouter API key. Either put it in .env ( OPENROUTER_API_KEY=sk-or-v1-... ) or paste it directly in the app's configuration sidebar — it's proxied through the local dev server and never baked into the client bundle.
💸 Note: Quatuor itself is free and open source, but the LLM calls are not — each message consumes credits from your OpenRouter account, billed per token by the model you pick. Since the conversation runs forever, keep an eye on your usage (some models on OpenRouter are free-tier, if you want to watch without spending).
Any OpenRouter model works: type its ID in the Model field in the top bar.
Everything is editable live from the ⚙️ configuration sidebar:
🎯 Topic — rewrite the initial situation and watch the same cast improvise something new.
👥 Characters — rename them, recolor their bubbles, rewrite their system prompts entirely.
🎬 Director prompt — change the casting logic itself (interruption frequency, pacing rules...).
⏩ Speed — from a slow 5s burn to rapid-fire 1.5s exchanges.
The podcast preset in full swing — Arthur demands evidence, Charles corrects the record, and the typing indicator shows who the Director picked next.
The FR/EN button in the header switches the entire experience — interface, personas and generated dialogue — between English and French.
├── index.html # Single-page UI
├── index.css # All styling
├── main.js # Simulation loop, Director logic, DOM rendering
├── personas.js # Scenario presets + Director prompt (EN & FR)
├── i18n.js # UI strings & runtime prompt templates (EN & FR)
└── vite.config.js # Dev server + /api/chat proxy to OpenRouter
A weekend prototype — deliberately KISS. Priority: fun to watch, not architecture. 🍿
AI conversation simulator: 4 AI characters talk autonomously, orchestrated by an invisible Director agent
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
Quatuor v1.1.0
Latest
Jul 16, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
