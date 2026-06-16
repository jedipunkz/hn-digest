---
source: "https://github.com/forgedculture/corpus-keeper"
hn_url: "https://news.ycombinator.com/item?id=48554450"
title: "Corpus Keeper – a zero-dependency auditor for the docs you point an AI at"
article_title: "GitHub - forgedculture/corpus-keeper · GitHub"
author: "forgedculture"
captured_at: "2026-06-16T13:57:05Z"
capture_tool: "hn-digest"
hn_id: 48554450
score: 1
comments: 0
posted_at: "2026-06-16T12:52:36Z"
tags:
  - hacker-news
  - translated
---

# Corpus Keeper – a zero-dependency auditor for the docs you point an AI at

- HN: [48554450](https://news.ycombinator.com/item?id=48554450)
- Source: [github.com](https://github.com/forgedculture/corpus-keeper)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T12:52:36Z

## Translation

タイトル: Corpus Keeper – AI に指示したドキュメントの依存関係のない監査人
記事のタイトル: GitHub - forgedculture/corpus-keeper · GitHub
説明: GitHub でアカウントを作成して、forgedculture/corpus-keeper の開発に貢献します。

記事本文:
GitHub - 鍛造文化/コーパスキーパー · GitHub
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
鍛造文化
/
コーパスキーパー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 他のアクションを開く m

enu フォルダーとファイル
2 コミット 2 コミット .github/ workflows .github/ workflowsデモ_コーパス デモ_コーパス .gitignore .gitignore ライセンス ライセンス README.md README.md check.sh check.sh corpus_keeper.py corpus_keeper.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI の性能は、AI が指定するフォルダーによって決まります。コーパスキーパー
そのフォルダーを正直な状態に保ちます。
AI アシスタントをドキュメントに組み込むすべてのチームは、
同じ病気: 文脈腐れ。まだ残っている古い価格表
現在。会議では中止されたが書面では中止された計画。
同じ事実について同意しない 2 つのドキュメント。 AIにはどちらなのか判断できない
1 つが true なので、完全な自信を持って 1 つが選択されます。失敗が起こるのは、
AIのせいだ。原因は体幹にあります。
Corpus Keeper は、依存関係のない監査役およびガバナンス足場です。
あなた、あなたのチーム、または AI が次のように扱うドキュメントのフォルダー
地上の真実。
git clone https://github.com/forgedculture/corpus-keeper
CD コーパスキーパー
python3 corpus_keeper.py 監査デモ_コーパス
FINDING [リンク] CORPUS_INDEX.md: 壊れたリンク -> roadmap.md
検索 [リンク] welcome.md: 壊れたリンク -> guides/setup.md
検索 [ascii] welcome.md: オフセット 101 の非 ASCII バイト
FINDING [インデックス] CORPUS_INDEX.md: インデックス エントリにファイルがありません -> roadmap.md
FINDING [index] unlisted_note.md: CORPUS_INDEX.md にリストされていないファイル
FINDING [stale] old_plan.md: 行 3 は最新の真実へのポインタがなく、古いとマークされています
info [stale] unlisted_note.md: 行 3 にオープン マーカー (TODO) があります
6 つのファイルをスキャンしました: 6 つの結果、1 つの情報
6 つの結果が表示されます: リンク切れ 2 つ、非 ASCII 文字、
ファントム インデックス エントリ、インデックスのないファイル、および非推奨のドキュメント
その代替品への指針はありません。次に、demo_corpus/pricing_2025.md を開きます。
およびdemo_corpus/pricing_current.md : どちらも引用可能であると主張しています
価格設定、そして彼らは同意しない

すべての数字に。その7番目の欠陥は、
スクリプトがキャッチできない種類 - 以下の「セマンティック層」を参照してください。
実行ごとに機械的な破損: 相対リンクの破損、非 ASCII バイト
(オプション)、インデックスドリフト (インデックスからファイルが欠落している、インデックスエントリ)
何も指していない)、および古いマーカー (SUPERSEDED、DEPRECATED)
現在の真実への指針はありません。
終了コードは契約です: 0 クリーン、1 検出、2 エラー。落ちる
変更せずにスクリプト、cron ジョブ、または CI に追加します。
独自のフォルダーを管理下に置く
python3 corpus_keeper.py init /path/to/your/folder
python3 corpus_keeper.py 監査 /path/to/your/folder
init はガバナンス層の足場を構築します: GOVERNANCE.md (ルール)、
CORPUS_INDEX.md (マップ)、番号付き意思決定記録テンプレート、および
追加専用の決定ログ。既存のファイルが上書きされることはありません。
ルールを一気に: 一度に 1 つの現在の真実で、決定が得られます
レコード、ログは追加専用、インデックスはマップ、毎回監査
編集します。
監査人は形式の腐敗を発見します。意味の腐敗 - 矛盾
ドキュメント間、まだ最新に見える古いドキュメント、変更
誰も決定を記録していませんでした - 読者が必要です。コーパスキーパー
キットは、そのパスを実行するように AI アシスタントを配線します: クロード スキル、
ChatGPT カスタム GPT セットアップ、および Codex、Cursor、Gemini 用の AGENTS.md
CLI、Copilot、さらにガバナンス テンプレートとサポート。
キットとその背後にある方法論は、次の場所にあります。
フォージドカルチャー.com 。
Python 3.8 以降、標準ライブラリのみ。 Apache 2.0 (「ライセンス」を参照)。
.corpuskeeper.json による構成 ( init によって作成): インデックス
ファイル名、ASCII ポリシー、除外パターン、ドキュメント拡張子。
保守者: コミットする前に ./check.sh を実行します。 CIも同じように動作します
スクリプト。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
エラーがありました

r 読み込み中。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to forgedculture/corpus-keeper development by creating an account on GitHub.

GitHub - forgedculture/corpus-keeper · GitHub
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
forgedculture
/
corpus-keeper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ workflows .github/ workflows demo_corpus demo_corpus .gitignore .gitignore LICENSE LICENSE README.md README.md check.sh check.sh corpus_keeper.py corpus_keeper.py View all files Repository files navigation
Your AI is only as good as the folder you point it at. Corpus Keeper
keeps that folder honest.
Every team wiring an AI assistant into their documents is accruing the
same disease: context rot. The old price sheet that still looks
current. The plan that was cancelled in a meeting but not in writing.
Two docs that disagree about the same fact. The AI cannot tell which
one is true, so it picks one - with total confidence. The failure gets
blamed on the AI. The cause is the corpus.
Corpus Keeper is a zero-dependency auditor and governance scaffold for
any folder of documents that you, your team, or your AI treats as
ground truth.
git clone https://github.com/forgedculture/corpus-keeper
cd corpus-keeper
python3 corpus_keeper.py audit demo_corpus
FINDING [links] CORPUS_INDEX.md: broken link -> roadmap.md
FINDING [links] welcome.md: broken link -> guides/setup.md
FINDING [ascii] welcome.md: non-ASCII byte at offset 101
FINDING [index] CORPUS_INDEX.md: index entry has no file -> roadmap.md
FINDING [index] unlisted_note.md: file not listed in CORPUS_INDEX.md
FINDING [stale] old_plan.md: line 3 marked stale with no pointer to current truth
info [stale] unlisted_note.md: line 3 has open marker (TODO)
scanned 6 files: 6 findings, 1 info
You will see 6 findings: two broken links, a non-ASCII character, a
phantom index entry, an unindexed file, and a deprecated document with
no pointer to its replacement. Then open demo_corpus/pricing_2025.md
and demo_corpus/pricing_current.md : both claim to be quotable
pricing, and they disagree on every number. That seventh defect is the
kind a script cannot catch - see "The semantic layer" below.
Mechanical rot, on every run: broken relative links, non-ASCII bytes
(optional), index drift (files missing from your index, index entries
pointing at nothing), and stale markers (SUPERSEDED, DEPRECATED) with
no pointer to current truth.
Exit codes are the contract: 0 clean, 1 findings, 2 error. It drops
into a script, cron job, or CI unchanged.
Bring your own folder under management
python3 corpus_keeper.py init /path/to/your/folder
python3 corpus_keeper.py audit /path/to/your/folder
init scaffolds the governance layer: GOVERNANCE.md (the rules),
CORPUS_INDEX.md (the map), a numbered decision-record template, and an
append-only decisions log. Existing files are never overwritten.
The rules in one breath: one current truth at a time, decisions get
records, logs are append-only, the index is the map, audit after every
edit.
The auditor catches rot of form. Rot of meaning - contradictions
between documents, a stale doc that still looks current, a change
nobody recorded a decision for - needs a reader. The Corpus Keeper
kit wires your AI assistant to do that pass: a Claude skill, a
ChatGPT Custom GPT setup, and an AGENTS.md for Codex, Cursor, Gemini
CLI, and Copilot, plus governance templates and support.
The kit, and the methodology behind it, live at
forgedculture.com .
Python 3.8+, standard library only. Apache 2.0 (see LICENSE).
Configuration via .corpuskeeper.json (written by init ): index
filename, ASCII policy, exclude patterns, document extensions.
Maintainers: run ./check.sh before committing; CI runs the same
script.
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
