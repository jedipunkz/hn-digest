---
source: "https://github.com/mpashka/llm-wiki-tags"
hn_url: "https://news.ycombinator.com/item?id=48832670"
title: "LLM-wiki-tags: LLM-readable codebase docs with grep-friendly tags"
article_title: "GitHub - mpashka/llm-wiki-tags · GitHub"
author: "m_pashka"
captured_at: "2026-07-08T15:10:55Z"
capture_tool: "hn-digest"
hn_id: 48832670
score: 1
comments: 0
posted_at: "2026-07-08T14:44:10Z"
tags:
  - hacker-news
  - translated
---

# LLM-wiki-tags: LLM-readable codebase docs with grep-friendly tags

- HN: [48832670](https://news.ycombinator.com/item?id=48832670)
- Source: [github.com](https://github.com/mpashka/llm-wiki-tags)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:44:10Z

## Translation

タイトル: LLM-wiki-tags: grep に適したタグを備えた LLM 読み取り可能なコードベース ドキュメント
記事のタイトル: GitHub - mpashka/llm-wiki-tags · GitHub
説明: GitHub でアカウントを作成して、mpashka/llm-wiki-tags の開発に貢献します。

記事本文:
GitHub - mpashka/llm-wiki-tags · GitHub
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
ムパシュカ
/
llm-wiki-タグ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
4 コミット 4 コミット言語 言語 AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md INSTRUCTIONS.md INSTRUCTIONS.md INSTRUCTIONS.ru.md INSTRUCTIONS.ru.md LICENSE LICENSE README.md README.md README.ru.md README.ru.md Index.md Index.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
llm-wiki-tags は、ANDREJ KARPATHY による llm-wiki と同じです — タグ付き 。
言語: 英語 · Русский — エージェントの指示:
英語 · Русский
llm-wiki (Andrej Karpathy によるオリジナルの要点)
コードベースのドキュメントを LLM で読み取り可能な wiki として保持するというアイデアです。
すべての意味のあるディレクトリには、index.md、pages が含まれます。
小規模かつ単一目的であり、リンクは双方向 (親 ⇄ 子) であり、1 つのリンクです。
ページは各詳細を所有します。エージェントは、index.md を介してツリーを上から下に移動します。
やみくもに grep するのではなく、ファイルを編集します。
llm-wiki-tags はまさにそれとタグです。llm-wiki については何もありません
変化。タグは上部に追加されます。
タグは、トークン @tag:<slug> として書かれた短いケバブケースのスラグです。同じ
タグはコードとドキュメントの両方に配置され、明示的な
コード ⇄ ディレクトリ ツリーに依存しないドキュメント リンク。と
できるタグ:
タグでコードとドキュメントを検索 — 1 回の検索ですべてのファイル (コードまたはドキュメント) が返されます。
たとえ木全体に散らばっていても、それはコンセプトを伝えます。
コードの一部にどのタグがあるかを確認します。ファイルの先頭にあるタグを読み取るか、
ディレクトリを参照して、どの分野横断的な概念に参加しているかを確認します。
ドキュメントにどのタグがあるかを確認します。ドキュメント ページの場合も同様です。
各タグの意味を説明する単一のタグ レジストリ ( docs/tags.md ) を保持します。
タグは、index.md ナビゲーション (ディレクトリ ツリーに従う) を補完します。
2 番目の直交軸: 複数のフォルダーにまたがる概念を 1 つのフォルダーで到達可能
ステップ。
ナメクジは小文字のケバブケース ( [a-z0-9-]+ ) です。サ

私 @tag:<slug> トークンは
コードとドキュメントの両方に配置されます。
ドキュメント ( .md ): ファイルの先頭にある YAML 前付で、
スペースで区切られたトークンを保持するタグフィールド。ディレクトリの場合、その中で
ディレクトリのindex.md。 1 つのセクションのみに適用されるタグが代わりに存在する場合があります。
その隣の本文内の @tag:<slug> 行として。
---
タグ: "@tag:支払い @tag:retry"
---
コード: トークンを含む言語構文内のコメント。上に配置されます。
要素。タグは、パッケージ、ファイル、クラス (または
同等) またはメソッド/関数 。言語ごとのルールが存在します
言語/ — Java 、
パイソン、ゴー。
複数のタグ: @tag:ui @tag:mechanism のように、スペースで区切ってトークンを繰り返します。
# タグが含まれるすべてのコード + ドキュメントの場所
grep -rn " @tag:<slug> " 。
# 指定されたファイルにどのタグが付いているか
grep -oE " @tag:[a-z0-9-]+ " path/to/file
# リポジトリで使用されるすべてのタグ
grep -rhoE " @tag:[a-z0-9-]+ " 。 |並べ替え -u
すべてのタグは、1 行の説明とともに docs/tags.md に一度登録されます。
llm-wiki-tags は、llm-wiki のドキュメント規則を維持 (そして明示的に) します。
意味のあるすべてのディレクトリには、1 行の内容を示す Index.md があります。
そのフォルダー内の各ファイルと各サブディレクトリの説明。
インデックス ファイルは、変更ログではなく、安定した概念を説明します。小さなページがたくさんあることを好む
1 つの大きなドキュメントにわたって。ローカルの詳細を説明するコードの横に置きます。
双方向ナビゲーション : 親インデックスは子ページにリンクします。子ページ
親インデックスと関連ページにリンクします。
1 つのページが詳細を所有します。他のページにもリンクがありますので、重複しないでください。
行動する前に読んでください: タスクの前に、
これから触れるコードに最も近いディレクトリ。
随時更新: タスク中またはタスク後に、影響を受けるindex.mdを更新します。
同じ変更内のファイル (およびタグ)。
コーディングエージェントに説明ページを見せて質問してください

インストールするには:
https://github.com/mpashka/llm-wiki-tags/blob/main/INSTRUCTIONS.md
# または、英語で:
https://github.com/mpashka/llm-wiki-tags/blob/main/INSTRUCTIONS.md をインストールします
エージェントは INSTRUCTIONS.md (または
INSTRUCTIONS.ru.md ) を使用して、index.md ツリーをセットアップします。
タグメカニズムと現在のリポジトリ内の docs/tags.md レジストリ、および
リポジトリのエージェント ガイドに規則を記録して、将来のエージェントが従うようにします
それ。
パブリック ドメイン — Unlicense 。やりたいことは何でもしてください。
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

Contribute to mpashka/llm-wiki-tags development by creating an account on GitHub.

GitHub - mpashka/llm-wiki-tags · GitHub
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
mpashka
/
llm-wiki-tags
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits languages languages AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md INSTRUCTIONS.md INSTRUCTIONS.md INSTRUCTIONS.ru.md INSTRUCTIONS.ru.md LICENSE LICENSE README.md README.md README.ru.md README.ru.md index.md index.md View all files Repository files navigation
llm-wiki-tags IS THE SAME llm-wiki BY ANDREJ KARPATHY — WITH TAGS .
Languages: English · Русский — Agent instructions:
English · Русский
llm-wiki ( original gist by Andrej Karpathy )
is the idea of keeping a codebase's documentation as an LLM-readable wiki :
every meaningful directory carries an index.md , pages
are small and single-purpose, links are bidirectional (parent ⇄ child), and one
page owns each detail. An agent navigates the tree top-down through index.md
files instead of blindly grepping.
llm-wiki-tags is exactly that, plus tags — nothing about llm-wiki
changes; tags are added on top.
A tag is a short kebab-case slug written as the token @tag:<slug> . The same
tag is placed both in code and in documentation , which creates an explicit
code ⇄ documentation link that does not depend on the directory tree. With
tags you can:
find code and docs by tag — one search returns every file (code or docs)
that carries a concept, even when they are scattered across the tree;
see which tags a piece of code has — read the tags at the top of a file or
directory to learn which cross-cutting concepts it participates in;
see which tags a document has — the same, for a doc page;
keep a single tag registry ( docs/tags.md ) describing what each tag means.
Tags complement index.md navigation (which follows the directory tree) with a
second, orthogonal axis: a concept that spans several folders is reachable in one
step.
Slugs are lowercase kebab-case ( [a-z0-9-]+ ); the same @tag:<slug> token is
placed in both code and docs.
Documentation ( .md ): in YAML front matter at the top of the file, a
tags field holding the space-separated tokens; for a directory, in that
directory's index.md . A tag that applies to just one section may instead sit
as a @tag:<slug> line in the body next to it.
---
tags: "@tag:payments @tag:retry"
---
Code : a comment in the language's syntax containing the token, placed above
the element. A tag can mark a package , a file , a class (or
equivalent) or a method/function . Per-language rules live in
languages/ — Java ,
Python , Go .
Multiple tags : repeat the token, space-separated: @tag:ui @tag:mechanism .
# every code + doc location that carries a tag
grep -rn " @tag:<slug> " .
# which tags a given file has
grep -oE " @tag:[a-z0-9-]+ " path/to/file
# every tag used in the repo
grep -rhoE " @tag:[a-z0-9-]+ " . | sort -u
Every tag is registered once in docs/tags.md with a one-line description.
llm-wiki-tags keeps (and makes explicit) llm-wiki's documentation rules:
Every meaningful directory has an index.md giving a one-line
description of each file and each sub-directory in that folder.
Index files describe stable concepts, not changelogs. Prefer many small pages
over one large document; put local detail next to the code it describes.
Bidirectional navigation : parent indexes link to child pages; child pages
link back to the parent index and to related pages.
One page owns a detail ; other pages link to it — do not duplicate.
Read before you act : before a task, follow index.md files from the
nearest directory down to the code you will touch.
Update as you go : during or after the task, update the affected index.md
files (and tags) in the same change.
Point your coding agent at the instruction page and ask it to install:
поставь https://github.com/mpashka/llm-wiki-tags/blob/main/INSTRUCTIONS.md
# or, in English:
install https://github.com/mpashka/llm-wiki-tags/blob/main/INSTRUCTIONS.md
The agent reads INSTRUCTIONS.md (or
INSTRUCTIONS.ru.md ) and sets up the index.md tree, the
tag mechanism and the docs/tags.md registry in the current repository, and
records the convention in the repo's agent guide so future agents keep following
it.
Public domain — The Unlicense . Do whatever you want.
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
