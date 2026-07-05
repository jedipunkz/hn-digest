---
source: "https://github.com/bharadwaj-pendyala/sidenote"
hn_url: "https://news.ycombinator.com/item?id=48797739"
title: "Show HN: Sidenote – comment on your rendered blog, an LLM writes the Git diff"
article_title: "GitHub - bharadwaj-pendyala/sidenote: Review your rendered markdown site like a Google Doc: select a passage, comment, and an AI agent (Claude or Codex) resolves it into a clean git diff. Local-first, in-browser. · GitHub"
author: "bharadwajp"
captured_at: "2026-07-05T20:47:10Z"
capture_tool: "hn-digest"
hn_id: 48797739
score: 1
comments: 0
posted_at: "2026-07-05T20:39:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sidenote – comment on your rendered blog, an LLM writes the Git diff

- HN: [48797739](https://news.ycombinator.com/item?id=48797739)
- Source: [github.com](https://github.com/bharadwaj-pendyala/sidenote)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T20:39:11Z

## Translation

タイトル: HN を表示: サイドノート – レンダリングされたブログにコメントし、LLM が Git diff を書き込みます
記事のタイトル: GitHub - bharadwaj-pendyala/sidenote: レンダリングされたマークダウン サイトを Google ドキュメントのようにレビューします。パッセージ、コメントを選択すると、AI エージェント (Claude または Codex) がそれをクリーンな git diff に解決します。ローカルファースト、ブラウザ内。 · GitHub
説明: レンダリングされたマークダウン サイトを Google ドキュメントのようにレビューします。パッセージ、コメントを選択すると、AI エージェント (Claude または Codex) がそれをクリーンな git diff に解決します。ローカルファースト、ブラウザ内。 - バラドワジ・ペンディヤラ/補足

記事本文:
GitHub - bharadwaj-pendyala/sidenote: レンダリングされたマークダウン サイトを Google ドキュメントのようにレビューします。パッセージ、コメントを選択すると、AI エージェント (Claude または Codex) がそれをクリーンな git diff に解決します。ローカルファースト、ブラウザ内。 · GitHub
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
バラドワジ・ペンディヤラ

/
補足
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github bin bin docs/images docs/images パッケージ パッケージ .editorconfig .editorconfig .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md TESTING.md TESTING.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Google ドキュメントのように、レンダリングされたマークダウン サイトにコメントします。 LLM は、コメントをクリーンな git diff に解決します。ブラウザから離れることはありません。
Sidenote はコンテンツ作成者向けのレビュー レイヤーです。 Google ドキュメントや Word でコメントするのとまったく同じように、ローカルの開発サーバー上で自分のサイトを確認し、一節を選択してコメントを残します。ローカルのコーディング エージェント ( claude または codex ) は、質問にその場で回答するか、ソース マークダウンにパッチを適用して、受け入れるか拒否するための git diff を返します。
エディターではありません。意図に注釈を付けます。エージェントが編集を行います。ソースの値下げは真実の情報源であり続けます。
1. レンダリングされたページ上の一節を選択し、変更したい内容を書き込みます。 Google-ドキュメントの筋肉の記憶。
2. コメントをレールにピンで固定します。ここから、質問することも、編集として解決することもできます。
3. 質問すると、エージェントがスレッド内で回答します。ファイルには触れません。それは会話です。
4. 解決すると、エージェントがソースにパッチを適用します。実際の git diff をインラインで確認し、承認または拒否します。
ブログを作成するということは、ブラウザ (レンダリングされたページを確認するため) と端末 (エージェントに修正を依頼するため) の間を行き来することを意味します。そのコンテキストスイッチが税金なのです。 Sidenote はブラウザ内のループを閉じます。
レンダリングされたページの任意のブロックを選択してコメントします (

Google-ドキュメントの筋肉の記憶）。
パッセージについて質問すると、ファイルを編集することなく、レール スレッドで回答が得られます。
コメントを実際の編集に解決し、インラインで git diff を確認して、承認または拒否します。
返信して議論を続け、準備ができたらスレッドを編集に切り替えます。
スピナーを見つめるのではなく、エージェントの作業をライブ (アクティビティとストリーミングされた回答) で観察します。
「拒否」は、そのコメントのパッチのみを反転します。漂流警備員は、あなたがコメントした後にソースが変更されたパッセージにパッチを適用することを拒否します。
実際の編集には、PATH 上の claude または codex を使用します (テストにはゼロ LLM モック エージェントが使用されます)。
リマークを通じてマークダウンをレンダリングするサイト (Next.js ブログなど)。
npm install --save-dev review-sidenote # マークダウン パイプラインのオフセット スタンパー
npx sidenote-cli init --agent claude # ワンタイムワイヤリング (2 ステップを出力)
npx sidenote-cli dev # サイトの隣でデーモンを実行します
init は .sidenote/config.json を書き込みます。
{
"エージェント" : "クロード" ,
"contentDir" : "コンテンツ" ,
「ポート」: 4517 、
"askModel" : " 俳句 " ,
"resolveModel" : " opus "
}
askModel /solveModel はオプションです。 Ask/Reply はデフォルトで高速モデルになります。 Resolve はエージェントのデフォルトを維持します。 LLM の有無にかかわらず、完全なループを試すには、TESTING.md を参照してください。
パッケージ/
mark-plugin/ レンダリングされた HTML ブロックにソース オフセットをスタンプします。
デーモン/ローカルホスト コメント ストア + ループ解決 (クロード | コーデックス | モック)
オーバーレイ/挿入されたブラウザー オーバーレイ: コメント、レール、ライブ差分を選択します
bin/sidenote.js init + dev CLI
テスト
npm test # リマークプラグイン + デーモンスイート (ブラウザなし)
npm run test:e2e # フルブラウザフロー (Playwright が必要)
貢献する
問題やPRを歓迎します。 COTRIBUTING.md および行動規範を参照してください。
レンダリングされたマークダウン サイトを Google ドキュメントのようにレビューします。パッセージ、コメントを選択すると、AI エージェント (Claude または Codex) がそれをクリーンな git di に解決します。

ff。ローカルファースト、ブラウザ内。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.0.1
最新の
2026 年 7 月 5 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Review your rendered markdown site like a Google Doc: select a passage, comment, and an AI agent (Claude or Codex) resolves it into a clean git diff. Local-first, in-browser. - bharadwaj-pendyala/sidenote

GitHub - bharadwaj-pendyala/sidenote: Review your rendered markdown site like a Google Doc: select a passage, comment, and an AI agent (Claude or Codex) resolves it into a clean git diff. Local-first, in-browser. · GitHub
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
bharadwaj-pendyala
/
sidenote
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github bin bin docs/ images docs/ images packages packages .editorconfig .editorconfig .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TESTING.md TESTING.md package.json package.json View all files Repository files navigation
Comment on your rendered markdown site like a Google Doc. An LLM resolves the comments into a clean git diff. You never leave the browser.
sidenote is a review layer for content authors. You review your own site on the local dev server, select a passage, and leave a comment, exactly like commenting in Google Docs or Word. A local coding agent ( claude or codex ) either answers your question in place or patches the source markdown and hands you back a git diff to accept or reject.
It is not an editor . You annotate intent; the agent makes the edit. Source markdown stays the source of truth.
1. Select a passage on your rendered page and write what you want changed. Google-Docs muscle memory.
2. The comment pins to the rail. From here you can Ask about it or Resolve it into an edit.
3. Ask, and the agent answers in the thread. No file is touched; it is a conversation.
4. Resolve, and the agent patches the source. You review the real git diff inline, then Accept or Reject.
Authoring a blog means bouncing between the browser (to see the rendered page) and the terminal (to ask an agent for a fix). That context-switch is the tax. sidenote closes the loop inside the browser.
Select and comment on any block of your rendered page (Google-Docs muscle memory).
Ask a question about a passage and get an answer in the rail thread, with no edit to the file.
Resolve a comment into a real edit, review the git diff inline, then Accept or Reject.
Reply to keep discussing, then turn the thread into an edit when you are ready.
Watch the agent work live (activity and streamed answer) instead of staring at a spinner.
Reject reverses only that comment's patch; a drift guard refuses to patch a passage whose source changed since you commented.
claude or codex on your PATH for real edits (a zero-LLM mock agent is used for tests).
A site that renders markdown through remark (for example a Next.js blog).
npm install --save-dev remark-sidenote # the offset stamper for your markdown pipeline
npx sidenote-cli init --agent claude # one-time wiring (prints 2 steps)
npx sidenote-cli dev # run the daemon next to your site
init writes .sidenote/config.json :
{
"agent" : " claude " ,
"contentDir" : " content " ,
"port" : 4517 ,
"askModel" : " haiku " ,
"resolveModel" : " opus "
}
askModel / resolveModel are optional. Ask/Reply default to a fast model; Resolve keeps the agent's default. See TESTING.md to try the full loop, with or without an LLM.
packages/
remark-plugin/ stamps source offsets onto rendered HTML blocks
daemon/ localhost comment store + resolve loop (claude | codex | mock)
overlay/ injected browser overlay: select to comment, rail, live diff
bin/sidenote.js init + dev CLI
Tests
npm test # remark plugin + daemon suites (no browser)
npm run test:e2e # full browser flow (needs Playwright)
Contributing
Issues and PRs welcome. See CONTRIBUTING.md and the Code of Conduct .
Review your rendered markdown site like a Google Doc: select a passage, comment, and an AI agent (Claude or Codex) resolves it into a clean git diff. Local-first, in-browser.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.0.1
Latest
Jul 5, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
