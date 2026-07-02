---
source: "https://github.com/kerbelp/context-md"
hn_url: "https://news.ycombinator.com/item?id=48765303"
title: "Show HN: Context.md – A proposed standard for AI project context"
article_title: "GitHub - kerbelp/context-md: Repository Context Layer — a git-native context.md standard for AI agents: consult before planning, write back what was learned · GitHub"
author: "kerbelp"
captured_at: "2026-07-02T19:10:40Z"
capture_tool: "hn-digest"
hn_id: 48765303
score: 2
comments: 1
posted_at: "2026-07-02T18:11:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Context.md – A proposed standard for AI project context

- HN: [48765303](https://news.ycombinator.com/item?id=48765303)
- Source: [github.com](https://github.com/kerbelp/context-md)
- Score: 2
- Comments: 1
- Posted: 2026-07-02T18:11:55Z

## Translation

タイトル: HN を表示: Context.md – AI プロジェクト コンテキストの標準案
記事のタイトル: GitHub - kerbelp/context-md: リポジトリ コンテキスト レイヤー — AI エージェント向けの git ネイティブ context.md 標準: 計画する前に相談し、学んだことを書き戻す · GitHub
説明: リポジトリ コンテキスト レイヤー — AI エージェント用の git ネイティブ context.md 標準: 計画する前に相談し、学習した内容を書き戻す - kerbelp/context-md

記事本文:
GitHub - kerbelp/context-md: リポジトリ コンテキスト レイヤー — AI エージェント用の git ネイティブ context.md 標準: 計画する前に相談し、学んだことを書き戻す · GitHub
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
ケルベルプ
/
コンテキストMD
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット ホワイトペーパー ホワイトペーパー ライセンス ライセンス README.md README.md SPEC.md SPEC.md context.md context.md context.md.example context.md.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントソフトウェア開発に欠けている層
パベル・カーベル — 独立研究者
抽象的。リポジトリは、ソース コード、テスト、ドキュメント、構成をバージョン管理します。これらのいずれも、エージェントに標準化され進化する操作コンテキストを提供しません。午前 2 時に発見された決定、拒否された代替案、制約。その知識はチャットのスクロールバックと人々の頭の中に存在し、すべてのエージェント セッションはそれなしで始まります。このノートでは、エージェントが計画前に参照し、実行後に人間のレビューの下で更新するバージョン管理された人間が読める形式のストアである リポジトリ コンテキスト レイヤー を提案します。 Git は変更内容を保存します。リポジトリ コンテキストには、プロジェクトが認識している内容が保存されます。それ以外はすべて飾りです。
📄 PDF · 📋 最小限の仕様 · 📝 例
エージェント セッションは、強力なモデルと空の頭から始まります。リポジトリ内のすべてのファイルを読み取ることができますが、ファイルがなぜそのように見えるのかはまだわかりません。コードは結果を記録します。理由はホワイトボードとマージの間のどこかで削除されました。
そこで代理人は決着した議論を再訴訟する。 3 月に削除した ORM を提案しています。本番環境で失敗した抽象化を再導入します。これはいずれもモデルの問題ではありません。リポジトリ内で最も高価な知識は否定的なもの、つまり試行され拒否されたものであるため、必要な知識はモデルが到達できない場所に存在します。コードでバージョン管理される標準アーティファクトはなく、エージェントが読む必要があり、継続的に更新されます。

実行を通じてed。この法案は、繰り返される間違い、アーキテクチャのドリフト、そして同じプロジェクトを同じモデルで再説明するという恒久的な税金として到来します。
II.既存のアーティファクトでは不十分な理由
明らかな反論には答えが必要だ。 README には、プロジェクトの変更方法ではなく、プロジェクトの使用方法が記載されており、決定が取り消されても誰も更新しません。 ADR は最も近い祖先ですが、人間にとっては 1 回だけ書き込むことができるエッセイです。エージェントがそれらを読み取る必要はなく、エージェントがそれらに追加することも期待されません。 RAG は類似性によって取得しますが、これは制約としてはまったく間違っています。ルールが最も重要になるのは、プロンプト内に類似するものがない場合です。 IDE メモリは 1 つのツールと 1 つのマシンにプライベートであり、確認することはできません。エージェント指示ファイルは人間からの命令を下位に伝達しますが、エージェント自体が学習した内容を吸収する定義された方法はありません。これらは静的で単方向ですが、コンテキスト層は動的で双方向です。それぞれがスライスを解決します。どれもコードでバージョン管理されず、契約によって参照され、人間のレビューの下でエージェントによって書き込み可能です。その組み合わせが不足しているレイヤーです。
Ⅲ．リポジトリコンテキストレイヤー
リポジトリ コンテキスト レイヤーは、バージョン管理され、人間が判読できるコンテキスト ストアであり、コードベースと並行して存在し、その中で動作する AI エージェントに対する権限のある操作コンテキストとして機能します。
Git は変更内容を保存します。
リポジトリ コンテキストには、プロジェクトが認識している内容が保存されます。
それをメモリではなくコンテキストと呼びます。記憶は個人的なものであり、曖昧で、オプションです。それはそれを形成したセッションとともに蒸発します。コンテキストは、モデルの実行時の動作の決定論的なコントロール プレーンです。ディスク上、コードでバージョン管理され、既知のアドレスにあります。エージェントがそれを参照する必要があり、それを維持することが期待されている時点で、それはメモの山ではなくなります。
相談→実行→更新→コミット
同僚との伝統的な

ntextレイヤー
エージェント エージェント
│ │
▼ ▼
リポジトリ リポジトリコンテキスト
│
▼
リポジトリ
│
▼
更新されたコンテキスト
契約には次の 4 つのステップがあります。
相談してください。計画を立てる前にコンテキストを読んでください。事前資料なしで作成された計画は、適切なフォーマットを備えた推測です。
実行する。制約を拘束力として扱い、意図を自由な設計選択のタイブレーカーとして扱います。
アップデート。作業で教えられたこと、ほとんどすべてのシステムがスキップするステップ、つまり ARM64 ビルドを中断するパッケージ、プロキシ タイムアウトについては何も文書化されていないことを追加します。
専念。レビューされた 1 つの変更で、コードと鮮明なコンテキストが移動します。
外部ストアは残りを提供できません。コンテキストはリポジトリとともに移動するため、リポジトリとともに進化し、コードが分岐するときは分岐し、マージするときはマージし、ロールバックするときはロールバックします。 2 つのブランチが異なることを学習すると、競合は他のすべてのマージ競合と同様に、人間によるレビューによって解決されます。
1 つの選択でほとんどの作業が実行されます。レイヤーは Git 内に存在します。 Git は、来歴、レビュー、分岐、マージ、分散同期、監査をすでに解決しています。リポジトリ コンテキスト レイヤーは、エージェントの知識のための新しい永続レイヤーを発明するのではなく、エンジニアがすでに信頼している機械を再利用し、その保証を無料で継承します。エージェントの信念を git log できない場合、それらをデバッグすることはできません。残りの原則はその選択に基づいて決まります。
人間が読める形式。プレーンなマークダウン。レビューで読むことができ、どのエディタでも編集できます。
レビュー可能。エージェントが独自の運用ルールを変更すると、その変更は動機となったコードの隣の差分に表示され、人間が両方を同時に承認します。人間の拒否権による自己変更: 公の場でコンテキストを編集するエージェントは、暗闇の中で記憶するエージェントよりも安全です。
ブランチを意識したもの。コンテキストは分岐モデルに従います。和解できる第二の真実の情報源はありません。
ツールに依存しない

凹み。 SDKもサーバーもベンダーもありません。ファイルを読み取ることができるエージェントであれば誰でも参加でき、テキスト エディターを使用する人間であれば誰でも一流のライターです。
段階的に進化しています。 3 つのセクションとそれぞれに 1 つの正直な文から始めます。レイヤーは、テンプレートが要求したときではなく、プロジェクトが学習するときに成長します。
決定論的な発見。エージェントはコンテキストを検索せずに見つける必要があります。固定パスは API です。
提案されているデフォルトの実装は、単一の Markdown ファイル context.md という考えられる中で最も鈍いものです。検出順序は .repo/context.md 、次にリポジトリ ルートの context.md です。最初のヒットが勝ちます。次の 3 つのセクションが必要です。
意図: このプロジェクトが何であるか、そして他のすべてが役立つべき設計哲学。制約: それぞれに理由がある、交渉不可能なルール。理由のない規則は迷信です。エージェントは従うでしょうが、一般化することはできません。ルールだけでなく、拒否も記録します。 Evolved Context : エージェントと人間がここで働いている間に学んだことを記録する追加専用の台帳。成果が証明されたエントリは制約に移行し、レビューされた昇進は具体的な自己改善ループとなります。
# リポジトリコンテキスト
## 意図
ローカルファーストの CLI。ファイルがソースです
真実; SQLite は再構築可能なインデックスです。
## 制約
- ORM はありません。拒否 2026-03: クエリの不透明度
オフライン修復パスが壊れました。
## 進化したコンテキスト
- [ 2026-06-29 ] pkg x >= 3.0 は ARM64 を破壊します
構築します。修正されるまで 2.3.x に固定します。
ここで標準化されているのは、コンテンツではなく、ディスカバリーとライフサイクルです。リポジトリはさまざまなコンテキスト構造を進化させますが、すべてのエージェントはどこを参照すればよいか、コンテキストがどのように変化するかを知っておく必要があります。 3 つのヘッダー (意思決定ログ、パターン カタログ、ディレクトリごとのコンテキスト) を超えるものはすべて、慣例により最上位に階層化されます。過剰に指定しないでください。不足仕様および採用されたビートが完了しました

そして無視されました。
VII.リファレンス実装
Metatron はこのレイヤーの実装の 1 つです。意思決定とプロジェクトの事実は git に裏付けされたマークダウンとして保持され、相談時にエージェントに提供され、フィードバックは進化したコンテキスト台帳にルーティングされます。オプションのインデックスを使用すると、検索が高速化されます。ファイルは真実のままです。これはアーキテクチャではなく実装です。どのツールでも、またはまったくどのツールでもこの​​レイヤーを実装できます。抽象化は、この製品を含め、その上に構築されたすべての製品よりも存続する必要があります。
リポジトリの人口が変わろうとしています。間もなく、コードベースのほとんどの読み手と書き手は人間ではなくなります。モデルの重みはリリース間で固定されます。リポジトリはそうである必要はありません。コンテキスト レイヤーを使用すると、すべてのセッションが最後よりも鮮明な事前音で終了し、トレーニング速度ではなくマージ速度で複雑になります。モデルではなくリポジトリがよりスマートになり、自己改善は微調整ではなくマージになります。学習の単位は、 context.md 内のレビューされた行です。
リポジトリは数十年前にコードのバージョン管理を学び、次にテスト、次にドキュメント、そしてインフラストラクチャを学びました。次のステップは、コンテキストをバージョン化することです。
リポジトリのルートに context.md を作成します (例から開始します)。
計画する前にそれを読み、コミットする前に進化したコンテキストに追加するようにエージェントに伝えます。
コードの差分と同様にコンテキストの差分を確認します。実績のある台帳エントリを制約にプロモートします。
© 2026 P. カーベル。無料でご利用いただけます。このリポジトリは、リポジトリ コンテキスト レイヤー ( PDF ) の正規のホームです。
リポジトリ コンテキスト レイヤー — AI エージェント用の git ネイティブ context.md 標準: 計画する前に相談し、学習した内容を書き戻す
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
リポジトリ コンテキスト レイヤー v0.1
最新の
2026 年 7 月 2 日
パッケージ
0
読み込み中にエラーが発生しました。このpをリロードしてください

年 。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Repository Context Layer — a git-native context.md standard for AI agents: consult before planning, write back what was learned - kerbelp/context-md

GitHub - kerbelp/context-md: Repository Context Layer — a git-native context.md standard for AI agents: consult before planning, write back what was learned · GitHub
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
kerbelp
/
context-md
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits whitepaper whitepaper LICENSE LICENSE README.md README.md SPEC.md SPEC.md context.md context.md context.md.example context.md.example View all files Repository files navigation
The Missing Layer of Agentic Software Development
Pavel Kerbel — Independent Researcher
Abstract. Repositories version source code, tests, documentation, and configuration. None of these gives an agent a standardized, evolving operating context. The decisions, the rejected alternatives, the constraint discovered at two in the morning: that knowledge lives in chat scrollback and in people's heads, and every agent session starts without it. This note proposes the Repository Context Layer , a version-controlled, human-readable store that agents consult before planning and update after executing, under human review. Git stores what changed; Repository Context stores what the project knows. Everything else is decoration.
📄 PDF · 📋 Minimal spec · 📝 Example
An agent session begins with a strong model and an empty head. It can read every file in the repository and still not know why any of them look the way they do. Code records outcomes; the reasons were deleted somewhere between the whiteboard and the merge.
So the agent re-litigates settled arguments. It proposes the ORM you removed in March. It reintroduces the abstraction that failed in production. None of this is a model problem. The knowledge it needs exists, just nowhere a model can reach, because the most expensive knowledge in a repository is negative: what was tried and rejected. No standard artifact is versioned with the code, required reading for agents, and continuously updated through execution. The bill arrives as repeated mistakes, architectural drift, and a permanent tax of re-explaining the same project to the same model.
II. Why Existing Artifacts Are Not Enough
The obvious rebuttals deserve answers. READMEs describe how to use a project, not how to change it, and nobody updates them when a decision is reversed. ADRs are the closest ancestor, but they are write-once essays for humans; no agent is required to read them and none is expected to append to them. RAG retrieves by similarity, which is precisely wrong for constraints: a rule matters most when nothing in the prompt resembles it. IDE memories are private to one tool and one machine, invisible to review. Agent instruction files carry orders from the human downward but have no defined way to absorb what the agent itself learns; they are static and unidirectional where the context layer is dynamic and bidirectional. Each solves a slice. None is versioned with the code, consulted by contract, and writable by the agent under human review. That combination is the missing layer.
III. The Repository Context Layer
A Repository Context Layer is a version-controlled, human-readable context store that lives alongside the codebase and serves as the authoritative operating context for AI agents working in it.
Git stores what changed.
Repository Context stores what the project knows.
Call it context, not memory. Memory is personal, fuzzy, optional; it evaporates with the session that formed it. Context is the deterministic control plane for the model's runtime behavior: on disk, versioned with the code, at a known address. It stops being a pile of notes the moment the agent is required to consult it and expected to maintain it.
consult → execute → update → commit
traditional with the context layer
agent agent
│ │
▼ ▼
repository repository context
│
▼
repository
│
▼
updated context
The contract has four steps:
Consult. Read the context before planning. A plan made without priors is a guess with good formatting.
Execute. Treat Constraints as binding and Intent as the tiebreaker for open design choices.
Update. Append what the work taught, the step almost every system skips: the package that breaks the ARM64 build, the proxy timeout nothing documents.
Commit. Code and sharpened context travel in one reviewed change.
No external store can offer the rest: context evolves with the repository because it travels with it, branching when the code branches, merging when it merges, rolling back when it rolls back. When two branches learn different things, the conflict is resolved like every other merge conflict: by a human, in review.
One choice does most of the work: the layer lives in Git. Git already solved provenance, review, branching, merging, distributed synchronization, and audit. Rather than inventing a new persistence layer for agent knowledge, the Repository Context Layer reuses the machinery engineers already trust, and inherits its guarantees for free. If you cannot git log your agent's beliefs, you cannot debug them. The remaining principles follow from that choice.
Human-readable. Plain Markdown. Readable in review, editable in any editor.
Reviewable. When an agent changes its own operating rules, the change appears in the diff next to the code that motivated it, and a human approves both at once. Self-modification with a human veto: an agent that edits its context in the open is safer than one that remembers in the dark.
Branch-aware. Context follows the branching model for free; no second source of truth to reconcile.
Tool-independent. No SDK, no server, no vendor. Any agent that can read a file participates, and any human with a text editor is a first-class writer.
Incrementally evolving. Start with three sections and one honest sentence in each; the layer grows when the project learns, not when a template demands it.
Deterministic discovery. The agent must find the context without searching for it. A fixed path is an API.
The proposed default implementation is the dullest thing possible: a single Markdown file, context.md . Discovery order is .repo/context.md , then context.md at the repository root; first hit wins. Three sections are required:
Intent : what this project is, and the design philosophy everything else must serve. Constraints : the non-negotiable rules, each with its reason. A rule without a reason is a superstition; the agent will comply but can never generalize. Record the rejection, not just the rule. Evolved Context : an append-only ledger of what agents and humans learned while working here. Entries that prove out graduate into Constraints, and that reviewed promotion is the self-improvement loop made tangible.
# Repository Context
## Intent
Local-first CLI. Files are the source of
truth; SQLite is a rebuildable index.
## Constraints
- No ORM. Rejected 2026-03: query opacity
broke the offline repair path.
## Evolved Context
- [ 2026-06-29 ] pkg x >= 3.0 breaks ARM64
builds. Pin to 2.3.x until fixed.
Note what is standardized here: discovery and lifecycle, not content. Repositories will evolve different context structures, but every agent should know where to look and how the context changes. Everything beyond the three headers (decision logs, pattern catalogs, per-directory context) is convention layered on top. Do not over-specify. Under-specified and adopted beats complete and ignored.
VII. A Reference Implementation
Metatron is one implementation of the layer: decisions and project facts kept as git-backed Markdown, served to agents at consult time, feedback routed into the evolved-context ledger. An optional index accelerates retrieval; the files remain the truth. It is an implementation, not the architecture. Any tool, or none at all, can implement the layer, and the abstraction should outlive every product built on it, including this one.
Repositories are about to change population: soon most readers and writers of a codebase will not be human. Model weights are frozen between releases; repositories do not have to be. With a context layer, every session ends with sharper priors than the last, compounding at merge speed rather than training speed. The repository gets smarter, not the model, and self-improvement becomes a merge instead of a fine-tune. The unit of learning is a reviewed line in context.md .
Repositories learned to version code decades ago, then tests, then documentation, then infrastructure. The next step is to version context.
Create context.md at your repository root (start from the example ).
Tell your agent to read it before planning and append to Evolved Context before committing.
Review context diffs like code diffs. Promote proven ledger entries into Constraints .
© 2026 P. Kerbel. Freely available. This repository is the canonical home of the Repository Context Layer ( PDF ).
Repository Context Layer — a git-native context.md standard for AI agents: consult before planning, write back what was learned
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Repository Context Layer v0.1
Latest
Jul 2, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
