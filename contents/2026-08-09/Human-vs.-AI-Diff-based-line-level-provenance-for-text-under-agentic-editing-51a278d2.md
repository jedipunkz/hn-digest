---
source: "https://github.com/eighttrigrams/us-vs-them"
hn_url: "https://news.ycombinator.com/item?id=49232300"
title: "Human vs. AI – Diff-based line-level provenance for text under agentic editing"
article_title: "GitHub - eighttrigrams/us-vs-them: Line-level provenance for text under agentic editing — who wrote this line, us or them? — derived from version history, not from markup in the file. · GitHub"
author: "eighttrigrams"
captured_at: "2026-08-09T16:23:09Z"
capture_tool: "hn-digest"
hn_id: 49232300
score: 2
comments: 0
posted_at: "2026-08-09T15:25:29Z"
tags:
  - hacker-news
  - translated
---

# Human vs. AI – Diff-based line-level provenance for text under agentic editing

- HN: [49232300](https://news.ycombinator.com/item?id=49232300)
- Source: [github.com](https://github.com/eighttrigrams/us-vs-them)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T15:25:29Z

## Translation

タイトル: 人間 vs. AI – エージェント編集中のテキストの差分ベースの行レベルの出自
記事のタイトル: GitHub - Eighttrigrams/us-vs-them: エージェント編集中のテキストの行レベルの出自 — この行を書いたのは誰ですか、私たちですか、それとも彼らですか? — ファイル内のマークアップからではなく、バージョン履歴から派生します。 · GitHub
説明: エージェント編集中のテキストの行レベルの出所 — この行を書いたのは誰ですか、私たちですか、それとも彼らですか? — ファイル内のマークアップからではなく、バージョン履歴から派生します。 - エイトトリグラム/私たち対彼ら

記事本文:
GitHub - Eighttrigrams/us-vs-them: エージェント編集中のテキストの行レベルの出自 — この行を書いたのは誰ですか、私たちですか、それとも彼らですか? — ファイル内のマークアップからではなく、バージョン履歴から派生します。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
八卦
/
私たち対彼ら
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
43 コミット 43 コミット src/ et/ uvt src/ et/ uvt test/ et/ uvt test/ et/ uvt .gitignore .gitignore Mak

efile Makefile README.md README.md deps.edn deps.edn title_image.png title_image.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント編集中のテキストの行レベルの出所 - この行を書いたのは誰か、私たち、または
彼らは？ — テキストのバージョン履歴から得られます。ライブラリまたはCLIツールとして使用します。
エージェントのコーディングと編集では、出所が適切な問題になります。テキストa
人間が書いたり編集したものは神聖なものに近いものとみなされるべきです。エージェントはそうすべきです。
躊躇してはいけませんが、それに触れるのには十分な理由があります。別のエージェントが作成したスロップ、
一方、完全に手に入る状態です。
このユースケース: 大部分がバイブコード化されたアプリを例に挙げて、
自分のアイデアや所有権を主張したいコード内のいくつかのコーナー。
別のエージェントがこのコード部分を強盗することは絶対に望まないでしょう
次のセッションで。
別の使用例: 最初に生成された README.md を書き換えます。
冒頭の段落。エージェントは自由にやり直したり、部分を追加したりできます。
さらに下向きですが、オープナーの何かを変更することはよく考えるべきです。
これが機能するための主な制約は、次のことを必要としないことです。
テキストがそのために特別にマークアップされるようにします。遍在するプレーンテキスト (マークダウン)
そのままサポートされるべきです。
その場合に利用できる唯一のことは、テキストの新しいバージョンが作成されることです。
人間またはエージェントの識別可能な作成者のもとで。
特定のテキストに対する評価の出力は、一連の範囲 (「アイランド」) です。
機械が生成したテキストの「海」の中に人間が書いた行。技術的には単純な差分に基づいており、
これはアルゴリズム開発の指針となるメタファーです。著者を追跡したくない
個々の行のみですが、意味のある一貫したテキストの断片です。
したがって、著者の結合、分割、および希釈は、

行動
考慮に入れられない方法でも、
完全な海または完全な島に集まります。
us-vs-them を CLI ツールとして使用するには、ローカル インストール用の bbin が必要です。
インストールする
git リポジトリはすでに、それぞれの出自を持つバージョンの履歴です。
マーカー - ファイルのすべてのリビジョンを順番に、変更の作成者とともに表示します。
できました。
git リポジトリ内の任意の場所で使用するには:
私たち対彼ら --ours dan@eighttrigrams.net README.md
これにより、次のようなリストが生成されます
1-3 0.00
4 1.00
5-7 0.00
8-20 0.46
21-164 0.00
ここで、1.0 は完全に人間が作成した範囲を意味します。
0.46 は、元々人間が作成した範囲をエージェントによってある程度変更されたことを意味します。
0.00 は、完全にエージェントが作成されたことを意味します。
--ours : これらは人間であり、他の人は全員エージェントとみなされます
--theirs : これらはエージェントであり、他の人はすべて人間とみなされます
リストの短い方の側に名前を付けます。
両方の引数を同時に渡すと拒否されます。
動作を理解する最良の方法は、以下を確認することです。
注意_テスト.clj 。
エージェント編集中のテキストの行レベルの出所 — この行を書いたのは誰ですか、私たちですか、それとも彼らですか? — ファイル内のマークアップからではなく、バージョン履歴から派生します。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Line-level provenance for text under agentic editing — who wrote this line, us or them? — derived from version history, not from markup in the file. - eighttrigrams/us-vs-them

GitHub - eighttrigrams/us-vs-them: Line-level provenance for text under agentic editing — who wrote this line, us or them? — derived from version history, not from markup in the file. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
eighttrigrams
/
us-vs-them
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits src/ et/ uvt src/ et/ uvt test/ et/ uvt test/ et/ uvt .gitignore .gitignore Makefile Makefile README.md README.md deps.edn deps.edn title_image.png title_image.png View all files Repository files navigation
Line-level provenance for text under agentic editing — who wrote this line, us or
them? — derived from a text's version history. Use it as library or as CLI tool.
With agentic coding and editing, provenance becomes a pertinent question. Text a
human wrote or edited should be considered close to sacred: an agent should
be hesitant and have a very good reason to touch it. Slop another agent has produced,
on the other hand, is completely up for grabs.
A use case for this: Take a mostly vibecoded app in which you want to establish
some corners in the code where you want to assert your ideas and ownership.
You surely don't want another agent bulldoze over this piece of code
in the next session.
Another use case: the README.md, originally generated, where you rewrite
the opening paragraphs. The agent should feel free to redo or append parts
further downwards but should really think twice changing anything in the opener.
The main constraint under which this should work is that this should not require
for text to be marked up specifically for that. Omnipresent plain text (markdown)
should be supported as is.
The only thing to leverage then, is that each new version of a text is created
under identifable authorship - of either a human or an agent.
The output of an evaluation over a given text is a set of ranges — "islands" of
human-authored lines inside a "sea" of machine generated text. Technically based on simple diffing,
this is the guiding metaphor for development of the algorithm. We don't want to track authorship
of individual lines only, but of meaningfully coherent pieces of text.
So joining, splitting apart, and dilution of authorship are behaviours
to be factored in, also in such a manner that we don't
converge in full sea or full island.
Using us-vs-them as CLI tool requires bbin for a local install.
make install
A git repository is already a history of versions each carrying a provenance
marker — every revision of the file, in order, with the author of the change that
made it.
To use it anywhere inside a git repository:
us-vs-them --ours dan@eighttrigrams.net README.md
This yields a listing like
1-3 0.00
4 1.00
5-7 0.00
8-20 0.46
21-164 0.00
where 1.0 means fully human authored range.
0.46 means originally human authored range, modified by agents to a certain degree.
0.00 means fully agent authored.
--ours : these are the humans, everonee else is considered an agent
--theirs : these are agents, everyone else is considered human
Name whichever side is the shorter list.
Passing both arguments at the same time will be rejected.
The best way to understand the behaviour is to have a look at
caution_test.clj .
Line-level provenance for text under agentic editing — who wrote this line, us or them? — derived from version history, not from markup in the file.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
