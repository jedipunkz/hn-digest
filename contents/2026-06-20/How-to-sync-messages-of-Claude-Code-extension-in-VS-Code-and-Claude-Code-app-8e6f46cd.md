---
source: "https://github.com/budhasantosh010/claude-code-session-bridge"
hn_url: "https://news.ycombinator.com/item?id=48608376"
title: "How to sync messages of Claude Code extension in VS Code and Claude Code app?"
article_title: "GitHub - budhasantosh010/claude-code-session-bridge: Ping-pong a single Claude Code session between the VS Code extension and the desktop app — by writing the desktop 'bookmark' that VS-Code-born chats are missing. · GitHub"
author: "realsanb"
captured_at: "2026-06-20T12:43:39Z"
capture_tool: "hn-digest"
hn_id: 48608376
score: 1
comments: 0
posted_at: "2026-06-20T11:16:37Z"
tags:
  - hacker-news
  - translated
---

# How to sync messages of Claude Code extension in VS Code and Claude Code app?

- HN: [48608376](https://news.ycombinator.com/item?id=48608376)
- Source: [github.com](https://github.com/budhasantosh010/claude-code-session-bridge)
- Score: 1
- Comments: 0
- Posted: 2026-06-20T11:16:37Z

## Translation

タイトル: VS Code と Claude Code アプリで Claude Code 拡張機能のメッセージを同期する方法?
記事のタイトル: GitHub - budhasantosh010/claude-code-session-bridge: VS Code 拡張機能とデスクトップ アプリの間で単一の Claude Code セッションをピンポンします — VS-Code 生まれのチャットにはないデスクトップの「ブックマーク」を書き込むことによって行います。 · GitHub
説明: VS Code 拡張機能とデスクトップ アプリの間で単一の Claude Code セッションをピンポンします。これは、VS-Code 生まれのチャットにはないデスクトップの「ブックマーク」を書き込むことによって行われます。 - budhasantosh010/claude-code-session-bridge

記事本文:
GitHub - budhasantosh010/claude-code-session-bridge: VS Code 拡張機能とデスクトップ アプリの間で単一の Claude Code セッションをピンポンします。これは、VS-Code 生まれのチャットにはないデスクトップの「ブックマーク」を書き込むことによって行われます。 · GitHub
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
ブダサントシュ01

0
/
クロードコードセッションブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
budhasantosh010/claude-code-session-bridge
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md register_vscode_session_in_desktop.ps1 register_vscode_session_in_desktop.ps1 すべてのファイルを表示 リポジトリ ファイルのナビゲーション
VS Code で開始した Claude Code チャットを Claude Code デスクトップ アプリに表示させることで、両方のアプリ間で 1 つの会話を行き来できるようになります。
Windows · 1 つの PowerShell スクリプト · インストール不要 · チャットを編集しない · MIT ライセンス
⚡ お急ぎですか？ 3ステップバージョン
完全な説明は以下にありますが、ただやりたいだけならこれですべてです。
ステップ 1 スクリプトをダウンロードします (緑色の「コード」ボタン → ZIP をダウンロード → 解凍)。
ステップ 2 VS Code で: [ターミナル] → [新しいターミナル]。 1 行を貼り付け、フォルダー内で入れ替えて、Enter キーを押します。
powershell -NoProfile -ExecutionPolicy Bypass -File "‹PATH TO THE .ps1›" -ProjectPath "‹YOUR VS CODE PROJECT FOLDER›" -Apply
ステップ 3 デスクトップ アプリを完全に終了し (トレイ アイコン → 終了)、再度開きます。
完了 — そのフォルダーの VS Code チャットがデスクトップ アプリのサイドバーに表示されます。
👉 «.ps1 へのパス» と «プロジェクト フォルダー» が混同されていますか? 3 フォルダーのセクションをお読みください。これは人々がつまずく最大の事柄です。
📖 すべての背後にある 1 つのアイデア — ノート + インデックスカード
❓ WH の質問 — 誰が / 何を / なぜ / いつ / どこで / どのように
🔒安全ですか？ — はい、そしてその理由はまさに
📥 インストール — スクリプトを PC に取得します
🧩 重要: これら 3 つのフォルダーを混同しないでください。最初にこれをお読みください。
🚀 使い方 — レシピをコピー＆ペースト
✅ レシピ A — 1 つのプロジェクトからすべての新しいチャットを追加します (

日常的なもの）
👀 レシピ B — プレビューのみ、何も変更しません
🎯 レシピ C — 特定のチャットを 1 つ追加する
🌍 レシピ D — PC 上のすべてのプロジェクト
⚠️ 誰もが忘れるステップ — デスクトップ アプリを再起動する
🆘 赤いエラー? 2 つの一般的な修正
🔁 両方のアプリ間で 1 つのチャットを実行します (ピンポン)
🧭 おまけ: フォルダー→「ドロワー」の命名
⚠️ 正直な制限 · 📜 ライセンス
📖 最初にお読みください: すべての背後にある 1 つのアイデア
すべてのクロード コード チャットがノートブックであり、それを示す図書館のインデックス カードがあると想像してください。
ノートブック = 実際のチャット (あなたとクロードが書いたすべてのメッセージ)
インデックスカード = 「ノートブックが存在します。その名前と場所は次のとおりです」と書かれた小さなメモです。
VS Code 拡張機能は、シェルフ (ディスク上のフォルダー) を検索してノートブックを見つけます。
デスクトップ アプリはより遅延しており、インデックス カードのみを読み取ります。カードがない = ノートがすぐそこにあるのに、ノートが見えません。
デスクトップ アプリでチャットを開始すると、インデックス カードが書き込まれます。 ✅
VS Code でチャットを開始すると、カードは書き込まれません。 ❌ → デスクトップ アプリはそれを認識しません。
このツールは、不足しているインデックス カードを書き込みます。それがすべてです。チャットごとに 1 枚のカード。
あなたのVS-CODEチャット
ノートブック ✔ (あなたのメッセージ - すでにディスク上にあります。私たちはそれに触れることはありません)
インデックス カード ✘ 見つかりません ── このツールを実行します ──► インデックス カード ✔ ──► デスクトップ アプリで認識されます
❓ WH の質問 (明確に回答)
Claude Code を複数の場所 (VS Code 拡張機能とデスクトップ アプリ) で使用していて、VS Code で開始されたチャットがデスクトップ アプリのサイドバーに表示されないことにイライラしている人は誰でもいます。
これは、インデックス カードがないチャット用に小さな「インデックス カード」ファイル (Claude Code ではセッション ラッパーと呼んでいます) を作成します。カードを挿入すると、デスクトップ アプリにチャットがリストされ、開くことができます。小さなポインタファイルを追加します

。他には何もありません。
なぜこれが必要なのでしょうか? (根本原因)
アプリとターミナル CLI は両方とも、ディスク上のまったく同じチャット ファイルを共有します。しかし、デスクトップ アプリはインデックス カードからのみサイドバーを構築し、デスクトップ アプリだけがそれらのカードを書き込みます。VS Code は決して書き込みません。したがって、VS-Code で生まれたチャットは、誰かがカードを書き込むまでデスクトップ アプリには表示されません。このツールはそれを書き込みます。
VS Code でチャットを開始し、デスクトップ アプリでもチャットしたいときはいつでも。何度でも実行できます。すでにカードを持っているチャットはスキップされるため、再実行は常に安全です。
どこに物を置きますか？ (正確なフォルダー)
ノートブック (あなたのチャット — 私たちはそれを読むだけで、決して変更しません):
C:\Users\<you>\.claude\projects\<folder-as-dashes>\<session-id>.jsonl
インデックス カード (このツールが書き込む内容):
C:\Users\<you>\AppData\Roaming\Claude\claude-code-sessions\<guid>\<guid>\local_<session-id>.json
段階的にどのように機能するのでしょうか?
フォルダー内の VS-Code 生まれのチャットごとに、ツールは次のことを行います。
1. チャット ファイルを開き、最初の ~50 行のみを読み取ります (読み取り専用ピーク)
2. 次の 2 つのことを理解します。
• チャットが属するフォルダ
• わかりやすいタイトル (チャットの最初のメッセージ) の前に日付を付ける
例: 「[6月13日]クロードのセシンを見つけるのを手伝ってください...」
3. ポインタ (「cliSessionId」) がそのチャット ファイルを指す小さなインデックス カードを書き込みます
4. カードをデスクトップ アプリの他のカードの隣にドロップします
それだけです。チャットを読み取り (ピークのみ)、新しい小さなカード ファイルを書き込みます。
🔒安全ですか？ （はい、まさにその理由はここにあります）
✘ 実際のチャット ファイル (.jsonl) を編集、移動、削除することはありません。
✘ すでに存在するカードを上書きすることはありません（それらはスキップされます）。
✘ 何も削除することはありません。カードファイルを追加するだけです。
✔ デフォルトでは DRY-RUN です。 -Apply を追加しない限り、何も書き込まれません。
✔ 簡単に元に戻すことができます。カード ファイルを削除します。

作成されましたが、リストは消えています
(カードはチャットではないため、チャットは完全に安全です)
📥 インストール — スクリプトをコンピュータに取得します
必要なファイルは 1 つだけです: register_vscode_session_in_desktop.ps1 。コンパイルやインストールする必要はありません。この 1 つのファイルを保存して実行するだけです。どちらかの方法を選択して入手してください。
方法 1 — ZIP をダウンロードします (最も簡単で、ツールは必要ありません)
1. https://github.com/budhasantosh010/claude-code-session-bridge を開きます
2.緑色の「＜＞コード」ボタンをクリック→「ZIPをダウンロード」
3. ダウンロード フォルダーで claude-code-session-bridge-main.zip を見つけます。
4. それを右クリック→「すべて展開...」→永続的なホームを選択します。 C:\ユーザー\<あなた>\ツール
5. これで、スクリプトがここに完成しました。
C:\Users\<you>\Tools\claude-code-session-bridge-main\register_vscode_session_in_desktop.ps1
方法 2 — git clone (すでに Git を持っている場合)
cd C:\Users\ < あなた > \Tools
git clone https://github.com/budhasantosh010/claude - コード - セッション - Bridge.git
結果:
C:\Users\<you>\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
📌 覚えやすい場所に永続的に保管してください (Tools\ フォルダーが最適です)。ダウンロード内に残さないでください。誤って削除してしまう可能性があります。
スクリプトのフルパスをコピーします（これをすべてのコマンドに貼り付けます）。
1. ファイル エクスプローラーで、register_vscode_session_in_desktop.ps1 を見つけます。
2. SHIFTを押しながら右クリック→「パスとしてコピー」
3. フルパスがクリップボードに追加されました。例:
「C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1」
🧩 重要 — これら 3 つのフォルダーを混同しないでください (これを読まないと混乱します)。
このツールを作成した人も含め、ほとんどの人が最初はこれにつまずきます。プレイ中には 3 つのまったく異なるフォルダーがあり、それらは同じものではありません。
① スクリプト (ツール) ← a .ps1

ファイル。単なるヘルパープログラムです。どこにでも住んでいます。
C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
② プロジェクトフォルダー ← コード/作業が存在する場所。 VS Code で開くフォルダー。
C:\Users\Alex\Documents\my-website
③ CLAUDE'S SESSIONS FOLDER ← クロード自身のプライベートストレージ。あなたは決してこれを指ささないでください。
C:\Users\Alex\.claude\projects\...
「どのフォルダーをコマンドに入れればいいですか？」
👉 フォルダー② — プロジェクト フォルダー (VS Code で開くフォルダー)。
🚫 フォルダー③ ( .claude\projects\... ) ではありません。
なぜ？ -ProjectPath の部分は、「どのプロジェクトのチャットが必要ですか?」と尋ねています。 — そして、コードが存在する場所に基づいてプロジェクトに名前を付けます ( ...\my-website などの人間の名前)。次に、スクリプトはそれを静かにフォルダー ③ に変換します。フォルダー③を自分で入力することはありません。
あなたは次のように入力します: -ProjectPath "C:\Users\Alex\Documents\my-website" ← フォルダ② (あなたのコード)
スクリプト図: → C:\Users\Alex\.claude\projects\c--Users-...-my-website\ ← フォルダ③の中を見てください。
(これはあなたの仕事ではなく、それ自体で行われます)
「なぜ .ps1 はランダムなフォルダーに保存されるのでしょうか? .claude 内に存在すべきではないでしょうか?」
.ps1 ファイルは Claude の一部ではありません。これは小さなヘルパー プログラムです。これはクロードの内部機構とは何の関係もないので、 .claude には属しません。
.claude\ = クロードの家。クロード自身のもの (セッション、設定、メモリ)。
── 他人の家に自分の道具を保管しないんですね ──
.ps1 = ドライバー。ツール。好きな引き出しに収納できます —
デスクトップ、Tools\ フォルダー、どこでも。全く同じように動作します。
スクリプトの場所はまったく関係ありません。重要なのは、コマンドが -File "..." 部分でそれを指していることです。
-ファイル「C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1」
━─────

── 「PowerShell さん、このツールを実行してください。ここにあります」 ──────┘
.ps1 を C:\Tools\ に移動しますか?次に、これを -File "C:\Tools\register_vscode_session_in_desktop.ps1" に変更するだけで、引き続き正常に動作します。
このツールを実行します ① — このプロジェクトのチャットをブックマークします ② — デスクトップ アプリに追加します。
(-ファイル "...ps1") (-ProjectPath "...コード フォルダー")
フォルダ③ (.claude) = クロードのプライベートストレージ。ツールがそれを読み取ります。決して名前を付けることはありません。
🚀 使い方（レシピをコピー＆ペースト）
すべてのコマンドに必要な 2 つのもの
① SCRIPT PATH = register_vscode_session_in_desktop.ps1 を保存した場所 (上にコピーしたもの)
② プロジェクト パス = そのチャットのために VS Code で開くフォルダー
プロジェクト パスをコピーするには: VS Code で、エクスプローラー サイドバーの最上位フォルダーを右クリック → [パスのコピー] をクリックします。
(または、エクスプローラーで、フォルダーを Shift キーを押しながら右クリック → 「パスとしてコピー」)。
1. PowerShell を開きます。VS Code で、[ターミナル] → [新しいターミナル] をクリックします。
(または、Windows キーを押し、「PowerShell」と入力し、Enter キーを押します)
2. 2 つのパスを記入したレシピを下に貼り付けます
3. Enter キーを押します
💡 黄金の習慣: -Apply を指定せずに実行します。これはプレビューであり、何も書き込みません。記載されている内容に満足していますか? Add -Apply を実行して再度実行し、実際にカードを書き込みます。
👇 以下のすべてのレシピでは、この 1 つの具体的な例を使用しています。独自の 2 つのパスを交換するだけです。
スクリプト パスの例: C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
プロジェクト パスの例: C:\Users\Alex\Documents\my-website
✅ レシピ A — 1 つのプロジェクト (毎日のプロジェクト) からすべての新しいチャットを追加します
プレーンエン

[切り捨てられた]

## Original Extract

Ping-pong a single Claude Code session between the VS Code extension and the desktop app — by writing the desktop 'bookmark' that VS-Code-born chats are missing. - budhasantosh010/claude-code-session-bridge

GitHub - budhasantosh010/claude-code-session-bridge: Ping-pong a single Claude Code session between the VS Code extension and the desktop app — by writing the desktop 'bookmark' that VS-Code-born chats are missing. · GitHub
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
budhasantosh010
/
claude-code-session-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
budhasantosh010/claude-code-session-bridge
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md register_vscode_session_in_desktop.ps1 register_vscode_session_in_desktop.ps1 View all files Repository files navigation
Make a Claude Code chat you started in VS Code show up in the Claude Code desktop app — so you can carry one conversation back and forth between both apps.
Windows · one PowerShell script · no install · never edits your chats · MIT licensed
⚡ In a hurry? The 3-step version
Full explanations are below — but if you just want it done, this is the whole thing.
STEP 1 Download the script (green "Code" button → Download ZIP → Extract).
STEP 2 In VS Code: Terminal → New Terminal. Paste ONE line, swap in your folder, press Enter:
powershell -NoProfile -ExecutionPolicy Bypass -File "‹PATH TO THE .ps1›" -ProjectPath "‹YOUR VS CODE PROJECT FOLDER›" -Apply
STEP 3 Fully quit the desktop app (tray icon → Quit) and reopen it.
Done — your VS Code chats from that folder now appear in the desktop app's sidebar.
👉 Confused by ‹PATH TO THE .ps1› vs ‹YOUR PROJECT FOLDER› ? Read the 3-folders section — it's the #1 thing people trip on.
📖 The one idea behind everything — notebook + index card
❓ The WH questions — who / what / why / when / where / how
🔒 Is it safe? — yes, and exactly why
📥 Install — get the script onto your PC
🧩 IMPORTANT: don't mix up these 3 folders — read this first
🚀 How to use it — copy-paste recipes
✅ Recipe A — add all new chats from one project (the everyday one)
👀 Recipe B — preview only, change nothing
🎯 Recipe C — add one specific chat
🌍 Recipe D — every project on the PC
⚠️ The step everyone forgets — restart the desktop app
🆘 Red error? The two common fixes
🔁 Carry one chat between both apps (ping-pong)
🧭 Bonus: the folder → "drawer" naming
⚠️ Honest limitations · 📜 License
📖 Read this first: the one idea behind everything
Imagine every Claude Code chat is a notebook , and there's a library index card that points to it.
THE NOTEBOOK = your actual chat (every message you and Claude wrote)
THE INDEX CARD = a tiny note that says "a notebook exists, here's its name and where it is"
The VS Code extension finds notebooks by looking on the shelf (the folder on your disk).
The desktop app is lazier — it only reads index cards . No card = it can't see the notebook, even though the notebook is right there.
When you start a chat in the desktop app , it writes the index card for you. ✅
When you start a chat in VS Code , no card gets written. ❌ → the desktop app is blind to it.
This tool writes the missing index card. That's the whole thing. One card per chat.
YOUR VS-CODE CHAT
notebook ✔ (your messages — already on disk, we never touch it)
index card ✘ MISSING ──run this tool──► index card ✔ ──► desktop app sees it
❓ The WH questions (answered plainly)
Anyone who uses Claude Code in more than one place — the VS Code extension and the desktop app — and is annoyed that chats started in VS Code never appear in the desktop app's sidebar.
It creates a small "index card" file (Claude Code calls it a session wrapper ) for chats that don't have one. With the card in place, the desktop app lists the chat and can open it. It adds tiny pointer files. Nothing else.
WHY is this even needed? (the root cause)
Both apps and the terminal CLI share the exact same chat file on disk. But the desktop app builds its sidebar only from index cards, and only the desktop app writes those cards — VS Code never does. So VS-Code-born chats are invisible to the desktop app until someone writes the card. This tool writes it.
Whenever you started a chat in VS Code and now want it in the desktop app too. Run it as often as you like — it skips chats that already have a card , so re-running is always safe.
WHERE does it put things? (exact folders)
THE NOTEBOOK (your chat — we only READ it, never change it):
C:\Users\<you>\.claude\projects\<folder-as-dashes>\<session-id>.jsonl
THE INDEX CARD (what this tool WRITES):
C:\Users\<you>\AppData\Roaming\Claude\claude-code-sessions\<guid>\<guid>\local_<session-id>.json
HOW does it work, step by step?
For each VS-Code-born chat in a folder, the tool:
1. Opens the chat file and reads ONLY its first ~50 lines (a read-only peek)
2. Figures out two things:
• which folder the chat belongs to
• a friendly title (the chat's first message), with the date in front
e.g. "[Jun 13] hey help me find the claude's sesins..."
3. Writes a tiny index card whose pointer ("cliSessionId") aims at that chat file
4. Drops the card next to the desktop app's other cards
That's it. It reads your chats (peek only) and writes new little card files.
🔒 Is it safe? (yes — here's exactly why)
✘ It NEVER edits, moves, or deletes your actual chat files (.jsonl)
✘ It NEVER overwrites a card that already exists (it skips those)
✘ It NEVER deletes anything — it only ADDS card files
✔ It is DRY-RUN by default — it writes NOTHING unless you add -Apply
✔ It is trivially reversible — delete the card file it made, and the listing is gone
(your chat stays perfectly safe, because the card was never the chat)
📥 Install — get the script onto your computer
You need just ONE file : register_vscode_session_in_desktop.ps1 . There is nothing to compile or install — you save this one file and run it. Pick either way to get it.
Way 1 — Download the ZIP (easiest, no tools needed)
1. Open https://github.com/budhasantosh010/claude-code-session-bridge
2. Click the green "< > Code" button → "Download ZIP"
3. Find claude-code-session-bridge-main.zip in your Downloads folder
4. Right-click it → "Extract All…" → choose a PERMANENT home, e.g. C:\Users\<you>\Tools
5. You now have the script here:
C:\Users\<you>\Tools\claude-code-session-bridge-main\register_vscode_session_in_desktop.ps1
Way 2 — git clone (if you already have Git)
cd C:\Users\ < you > \Tools
git clone https: // github.com / budhasantosh010 / claude - code - session - bridge.git
Result:
C:\Users\<you>\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
📌 Keep it somewhere permanent you'll remember (a Tools\ folder is perfect). Don't leave it in Downloads — too easy to delete by accident.
Copy the script's full path (you'll paste this into every command)
1. In File Explorer, find register_vscode_session_in_desktop.ps1
2. Hold SHIFT, right-click it → "Copy as path"
3. The full path is now on your clipboard, e.g.:
"C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1"
🧩 IMPORTANT — don't mix up these 3 folders (read this or you'll get confused)
Almost everyone trips on this at first — including the person who built this tool. There are three completely different folders in play, and they are NOT the same thing:
① THE SCRIPT (the tool) ← a .ps1 file. Just a helper program. Lives ANYWHERE.
C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
② YOUR PROJECT FOLDER ← where YOUR code/work lives. The folder you open in VS Code.
C:\Users\Alex\Documents\my-website
③ CLAUDE'S SESSIONS FOLDER ← Claude's OWN private storage. You NEVER point at this.
C:\Users\Alex\.claude\projects\...
"Which folder do I put in the command?"
👉 Folder ② — your project folder (the one you open in VS Code).
🚫 NOT folder ③ ( .claude\projects\... ).
Why? The -ProjectPath part is asking "which project's chats do you want?" — and you name a project by where its code lives (its human name, like ...\my-website ). The script then quietly translates that into folder ③ for you. You never type folder ③ yourself.
YOU type: -ProjectPath "C:\Users\Alex\Documents\my-website" ← folder ② (your code)
SCRIPT figures: → look inside C:\Users\Alex\.claude\projects\c--Users-...-my-website\ ← folder ③
(it does this itself — not your job)
"Why is the .ps1 saved in some random folder? Shouldn't it live inside .claude ?"
The .ps1 file is NOT part of Claude. It's a small helper program. It has nothing to do with Claude's internal machinery, so it does not belong in .claude .
.claude\ = Claude's house. Claude's own stuff (sessions, settings, memory).
── you don't store your own tools inside someone else's house ──
the .ps1 = YOUR screwdriver. A tool. It can sit in ANY drawer you like —
Desktop, a Tools\ folder, anywhere. It works exactly the same.
The script's location doesn't matter at all. What matters is that your command points at it with the -File "..." part:
-File "C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1"
└─────────── "hey PowerShell, run THIS tool — it's right here" ───────────┘
Move the .ps1 to C:\Tools\ ? Then just change that to -File "C:\Tools\register_vscode_session_in_desktop.ps1" and it still runs fine.
RUN this tool ① ── to bookmark the chats of this project ② ── into the desktop app.
(-File "...ps1") (-ProjectPath "...your code folder")
Folder ③ (.claude) = Claude's private storage. The tool reads it FOR you. You never name it.
🚀 How to use it (copy-paste recipes)
The 2 things every command needs
① SCRIPT PATH = where you saved register_vscode_session_in_desktop.ps1 (copied just above)
② PROJECT PATH = the folder you open in VS Code for that chat
To copy the PROJECT PATH: in VS Code, right-click the top folder in the Explorer sidebar → "Copy Path" .
(Or in File Explorer, Shift+right-click the folder → "Copy as path" .)
1. Open PowerShell: in VS Code click Terminal → New Terminal
(or press the Windows key, type "PowerShell", press Enter)
2. Paste a recipe below — with YOUR two paths filled in
3. Press Enter
💡 Golden habit: run it without -Apply first — that's a preview and writes nothing. Happy with what it lists? Add -Apply and run again to actually write the cards.
👇 Every recipe below uses this ONE concrete example. Just swap in your own two paths.
Example SCRIPT PATH : C:\Users\Alex\Tools\claude-code-session-bridge\register_vscode_session_in_desktop.ps1
Example PROJECT PATH: C:\Users\Alex\Documents\my-website
✅ Recipe A — add all new chats from ONE project (the everyday one)
Plain En

[truncated]
