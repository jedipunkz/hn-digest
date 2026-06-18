---
source: "https://github.com/cfitzgerald-pd/1pager/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48589336"
title: "Show HN: 1pager: Fixing the AI verbosity problem"
article_title: "GitHub - cfitzgerald-pd/1pager: AI is too verbose. 1pager condenses any convo/thought/workspace into a single 1-pager optimized for human review. · GitHub"
author: "bennydog224"
captured_at: "2026-06-18T18:53:53Z"
capture_tool: "hn-digest"
hn_id: 48589336
score: 1
comments: 0
posted_at: "2026-06-18T18:19:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: 1pager: Fixing the AI verbosity problem

- HN: [48589336](https://news.ycombinator.com/item?id=48589336)
- Source: [github.com](https://github.com/cfitzgerald-pd/1pager/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T18:19:18Z

## Translation

タイトル: HN の表示: 1 ページャー: AI の冗長性の問題の修正
記事のタイトル: GitHub - cfitzgerald-pd/1pager: AI は冗長すぎます。 1pager は、あらゆる会話/考え/ワークスペースを、人間によるレビュー用に最適化された 1 つの 1pager に凝縮します。 · GitHub
説明: AI は冗長すぎます。 1pager は、あらゆる会話/考え/ワークスペースを、人間によるレビュー用に最適化された 1 つの 1pager に凝縮します。 - cfitzgerald-pd/1 ページャー

記事本文:
GitHub - cfitzgerald-pd/1pager: AI は冗長すぎます。 1pager は、あらゆる会話/考え/ワークスペースを、人間によるレビュー用に最適化された 1 つの 1pager に凝縮します。 · GitHub
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
フィッツジェラルド-pd
/
1 ページャー
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット 参照 参照 スクリプト スクリプト README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
長い文書、ディレクトリ、チャットを厳密な内容に凝縮するスキル
1 ページの概要 — 箇条書きで、冗長性は可能な限り最小限 — 人間味のある内容になっており、
.md と .docx をエクスポートします (Word および Google ドキュメントで開きます)。
LLM はデフォルトでは冗長です。編集せずに放置すると、出力は長く、パディングされ、
テルがいっぱい（ルールオブスリー、「デルブ」、エムダッシュスプリー、一般的な明るいクローザー）。
その代償は読者の負担となります。カミーユ・フルニエが次のように主張している。
AI を尊重して使用するためのガイドライン、
レビューされずに肥大化した AI 出力を出荷すると、検証税が発生する : 同僚
著者がわざわざトリミングしなかったテキストを読んで検証する必要があります。長いPR、壁-
彼女の言葉を借りれば、テキストメッセージや肥大化したドキュメントは「率直に言って、ただ失礼」です。
簡潔さは敬意の一形態です。短い方が良いのは、内容を保護するためです。
読者の時間。
このスキルは、作成者が行うべきトリミング、つまり出力のカットを行うために存在します。
読者が理解する必要があること、または判断する必要があることだけを 1 ページにまとめます。
必要なものをそぎ落として、要点、決定/質問、少数の部分を維持します
裏付けとなる事実、実際のリスク、具体的な次のステップ。他のすべてをドロップします。
最初に要点を記載し、次に箇条書きを記載する、厳密な概要資料の草案を作成します。 1ページ、難しい
天井（約 350 ワード）。
人間化 — AI の書き込み指示を削除します。インストールされているヒューマナイザーを呼び出します
存在する場合はスキル。それ以外の場合は、バンドルされているreferences/humanizer.mdにフォールバックします。
2 つのファイル、 <name>-1pager.md および <name>-1pager.docx をエクスポートします。
1ページ/
§── SKILL.md # スキル: 凝縮 → 擬人化 → エクスポート
§── README.md # このファイル
§── スクリプト/

│ └── md_to_docx.py # pure-stdlib Markdown → .docx (依存関係なし)
━── 参考文献/
└── humanizer.md # バンドルされた humanizer ルール (フォールバック)
使用法
クロード コードでは、短縮、短縮、または作成を要求すると、スキルがトリガーされます。
概要の概要:
「この設計ドキュメントを 1 ページ冊にまとめてください。」
「このスレッドは長すぎます。共有できる内容を教えてください。」
「docs/ フォルダーを 1 ページに要約します。」
ディレクトリを指定すると、関連するファイルが 1 つのページに結合されます。
デフォルト (代わりにそれを要求する場合は、ファイルごとに 1 ページ)。
コンバータを直接実行することもできます。
python3 scripts/md_to_docx.py my-1pager.md my-1pager.docx
コンバーターのスコープ
md_to_docx.py は、概要資料に必要な小さな Markdown サブセットをサポートしています。 # – ####
見出し、 - / * 箇条書き (2 スペースのインデント = サブ箇条書き)、 ** 太字 ** 、および
`code` (等幅でレンダリング)。外部依存関係はありません。.docx は
Python 標準ライブラリで構築された単なる XML の zip です。
AIは冗長すぎます。 1pager は、あらゆる会話/考え/ワークスペースを、人間によるレビュー用に最適化された 1 つの 1pager に凝縮します。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

AI is too verbose. 1pager condenses any convo/thought/workspace into a single 1-pager optimized for human review. - cfitzgerald-pd/1pager

GitHub - cfitzgerald-pd/1pager: AI is too verbose. 1pager condenses any convo/thought/workspace into a single 1-pager optimized for human review. · GitHub
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
cfitzgerald-pd
/
1pager
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit references references scripts scripts README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A skill that condenses any long document, directory, or chat into a strict
one-page summary — bullet-first, least verbosity possible — humanizes it, and
exports a .md plus a .docx (opens in Word and Google Docs).
LLMs are verbose by default. Left unedited, their output is long, padded, and
full of tells (rule-of-three, "delve", em-dash sprees, generic upbeat closers).
The cost lands on the reader. As Camille Fournier argues in
Guidelines for respectful use of AI ,
shipping unreviewed, bloated AI output creates a validation tax : colleagues
have to read and verify text the author never bothered to trim. Long PRs, wall-
of-text messages, and bloated docs are, in her words, "frankly, just rude."
Brevity is a form of respect — shorter is better because it protects the
reader's time.
This skill exists to do the trimming the author should have done: cut the output
down to only what the reader needs to understand or decide, on a single page.
Strips to essentials — keeps the point, the decision/ask, the few
supporting facts, real risks, and concrete next steps. Drops everything else.
Drafts a strict 1-pager — bottom line first, then bullets. One page, hard
ceiling (~350 words).
Humanizes — removes AI-writing tells. Invokes the installed humanizer
skill if present; otherwise falls back to the bundled references/humanizer.md .
Exports two files — <name>-1pager.md and <name>-1pager.docx .
1pager/
├── SKILL.md # the skill: condense → humanize → export
├── README.md # this file
├── scripts/
│ └── md_to_docx.py # pure-stdlib Markdown → .docx (no dependencies)
└── references/
└── humanizer.md # bundled humanizer rules (fallback)
Usage
In Claude Code, the skill triggers when you ask to condense, shorten, or make a
1-pager out of something:
"Condense this design doc into a one-pager."
"This thread is way too long — give me a tl;dr I can share."
"Boil the docs/ folder down to a single page."
Point it at a directory and it combines the relevant files into one page by
default (one page per file if you ask for that instead).
You can also run the converter directly:
python3 scripts/md_to_docx.py my-1pager.md my-1pager.docx
Converter scope
md_to_docx.py supports the small Markdown subset a 1-pager needs: # – ####
headings, - / * bullets (two-space indent = sub-bullet), **bold** , and
`code` (rendered monospace). It has no external dependencies — a .docx is
just a zip of XML, built with the Python standard library.
AI is too verbose. 1pager condenses any convo/thought/workspace into a single 1-pager optimized for human review.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
