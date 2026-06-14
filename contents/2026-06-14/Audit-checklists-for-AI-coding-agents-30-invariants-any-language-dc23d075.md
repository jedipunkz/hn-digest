---
source: "https://github.com/danygiguere/audit-skills"
hn_url: "https://news.ycombinator.com/item?id=48531259"
title: "Audit checklists for AI coding agents – 30 invariants, any language"
article_title: "GitHub - danygiguere/audit-skills: Language- and framework-agnostic audit checklists for AI coding agents — security, correctness, and operability. Works with Claude Code, GitHub Copilot, Cursor, Codex CLI, OpenCode, and any agent that can read files. · GitHub"
author: "danygiguere"
captured_at: "2026-06-14T21:33:38Z"
capture_tool: "hn-digest"
hn_id: 48531259
score: 1
comments: 0
posted_at: "2026-06-14T19:02:17Z"
tags:
  - hacker-news
  - translated
---

# Audit checklists for AI coding agents – 30 invariants, any language

- HN: [48531259](https://news.ycombinator.com/item?id=48531259)
- Source: [github.com](https://github.com/danygiguere/audit-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T19:02:17Z

## Translation

タイトル: AI コーディング エージェントの監査チェックリスト – 30 個の不変条件、任意の言語
記事のタイトル: GitHub - danygiguere/audit-skills: AI コーディング エージェント向けの言語とフレームワークに依存しない監査チェックリスト - セキュリティ、正確性、操作性。 Claude Code、GitHub Copilot、Cursor、Codex CLI、OpenCode、およびファイルを読み取ることができる任意のエージェントで動作します。 · GitHub
説明: AI コーディング エージェント向けの言語およびフレームワークに依存しない監査チェックリスト - セキュリティ、正確性、操作性。 Claude Code、GitHub Copilot、Cursor、Codex CLI、OpenCode、およびファイルを読み取ることができる任意のエージェントで動作します。 - danygiguere/監査スキル

記事本文:
GitHub - danygiguere/audit-skills: AI コーディング エージェント向けの言語およびフレームワークに依存しない監査チェックリスト - セキュリティ、正確性、操作性。 Claude Code、GitHub Copilot、Cursor、Codex CLI、OpenCode、およびファイルを読み取ることができる任意のエージェントで動作します。 · GitHub
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
ディミ

SSアラート
{{ メッセージ }}
ダニギゲレ
/
監査スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .agents/ スキル .agents/ スキル .claude .claude .github/ ワークフロー .github/ ワークフロー デモ デモ プロジェクト プロジェクト スクリプト スクリプト .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md バージョンバージョン すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント向けの言語やフレームワークに依存しない監査チェックリスト —
安全性、正確性、操作性。クロードコード、GitHub で動作します
コパイロット、カーソル、Codex CLI、OpenCode、およびファイルを読み取ることができる任意のエージェント。
すべてのチェックリストは不変条件として書かれており、検出は臭いとして書かれています。
フレームワーク API なので、同じコンテンツが Rails アプリ、Spring サービス、
または Express API — エージェントはフレームワーク固有の翻訳を提供します。
20 行のマネー ハンドラーの /audit — 静的分析スキャナーの 6 つのバグ
それぞれが所有権、同時実行性、および
パターンマッチングではなく再試行です。すべてにフラグが付けられ、重大度が示され、修正が行われました。
AGENTS.md — 30 個すべての不変条件の 1 ページのダイジェスト。その内容をコピーする
プロジェクトの AGENTS.md に追加すると、すべてのエージェントがコンテキスト内でそれを保持できるようになります。
.agents/skills/audit/ — ルーター スキル、30 個のチェックリストすべてと
修復パターンは、references/ (4 つのカテゴリ: アクセス) の下にバンドルされています。
データセキュリティ、入力/API、正確性、操作性）。
.agents/skills/audit-* — トピックごとの薄いラッパー スキルなので、各チェックリスト
個別に呼び出すことができます ( /audit-idor 、 /audit-injection 、
/audit-fix-authz , …)。このパッケージがインストールするものはすべて次から始まります
Audit なので、他のスキルとグループ化されたままになります。
/au

dit は完全な監査を実行します。コードが何を実行し、適用するかを特定します。
以下の一致するすべてのチェックリスト。各トピックは個別に呼び出すこともできます
(クリックしてチェックリスト自体を読んでください)。
あらゆる言語やフレームワークで動作します。各チェックリストには 8 つの名前が挙げられています
概念用語集の共通エコシステム (Rails、Laravel、Django、Spring、
Node、Vapor、.NET、Go) — これらは認識ショートカットであり、
サポートリスト。不変条件と検出匂いはフレームワークに依存しないため、
監査は、Phoenix、FastAPI、Ktor、または社内に同様に適用されます。
スタック: エージェントが翻訳を提供します。
監査
チェック項目
/監査認可
アクションの時点でのサーバー側の権限チェック - 権限昇格、UI のみのゲート、読み取りではなく変更ではチェック
/監査認証セッション
ログイン、ログアウト、およびリセットのフロー — セッションの固定、アカウントの列挙、トークンの有効期限、および使い捨てのリメンバーミーストレージ
/監査アイドル
リクエスタがアクセスできるかどうかを確認せずに、リクエストで指定された ID によってリソースがフェッチまたは変更された場合
/監査データの公開
過剰に公開された応答、エラー、ログ - モデル全体のシリアル化、スタック トレース、PII
/audit-crypto
パスワードハッシュ、トークンのランダム性、定数時間比較、自家製暗号、キー処理
/監査出力エンコーディング
XSS — コンテキストに応じたエンコードを行わずに HTML、JS、CSS、URL、ヘッダー、または電子メールにレンダリングされたユーザー データ
/監査テナント分離
テナント間の漏洩 - スコープ外のクエリ、テナントレスのキャッシュ キー、テナントをまたぐバックグラウンド ジョブ
/audit-csrf
CSRF トークンやオリジン検証を行わない Cookie/セッション認証でのエンドポイントの状態変更
/audit-mass-assignment
モデルに大規模にバインドされたリクエスト ペイロード — 書き込み可能なロール/所有者/バランス フィールド、ホワイトリストの代わりに拒否リスト
入力とAPI
監査
チェック項目
/監査インジェクション
SQL/NoSQL、コマンド、テンプレート、

d パス インジェクション - クエリ、シェル、またはテンプレートに連結された入力
/監査構成
安全でない構成 - 運用環境でのデバッグ、許容的な CORS、セキュリティ ヘッダーの欠落、Cookie フラグ
/監査秘密
ハードコードされた資格情報、ログまたはバージョン管理内の秘密、広すぎるキー、ローテーション パスなし
/audit-api-validation
境界検証 — タイプ、境界、許可されたフィールド、価格や役割などのクライアントが計算した値の信頼性
/監査ファイルの処理
パス トラバーサル、未検証のアップロード、サイズ制限の欠如、Web ルートから提供されるファイル、ジップスリップ
/audit-ssrf
ユーザーの影響を受ける URL に対するサーバー側のリクエスト - ホワイトリスト、プライベート IP 範囲、リダイレクトの再検証。オープンリダイレクトを含む
/監査パーサー差分
バリデーターは受け入れますが、コンシューマーは異なる方法で読み取る入力 — アンカーされていない正規表現、許可リストで始まる、2 つの URL パーサー、検証してから再解析する
正しさ
監査
チェック項目
/監査アトミック性
トランザクションを使用しないマルチストア書き込み - 失敗しても部分的な状態が維持される
/audit-冪等性
2 回実行すると誤動作するハンドラー - Webhook、支払い、キューの再配信、二重送信
/監査バックグラウンド作業
ジョブとコンシューマ - 無制限の再試行、有害なメッセージ、タイムアウトの欠落、重複または順序どおりでない配信
/監査状態管理
競合状態 — ロック、アトミック プリミティブ、または制約を使用せずに、共有状態をチェックしてから実行します。
/監査例外処理
飲み込まれたエラー、ブランケットキャッチ、失われた原因、クリーンアップの欠落、および間違った HTTP ステータス (404 対 403、401、422、409)
/audit-discarded-async
Fire-and-forget バグ - Promise、Future、またはリアクティブ パブリッシャーが作成されたものの、待機、返され、または構成されなかったもの。裸の購読。サイレントに実行されないコールド書き込み
/監査カーディナリティ
クエリが 1 つの行に一致することを前提とした操作 - 一意でない c に対する UPDATE/DELETE

列のファンアウト、一意でないフィールドの findOne/.single()、DB 制約なしで列が一意として扱われる
操作性
監査
チェック項目
/audit-nplus1
コレクションに対するループ内で行われるクエリ - または HTTP/キャッシュ呼び出し
/監査可観測性
サイレント障害 - 飲み込まれたエラー、識別子のないログ、メトリックまたはアラート パスなし
/audit-migration-safety
テーブルをロックするスキーマ変更、エキスパンドコントラクトを伴わない破壊的な変更、バッチ処理されていないバックフィル
/監査リソース制限
入力からの無制限の作業 - ページネーションの欠落、サイズの上限、レート制限、壊滅的な正規表現
/audit-blocking-io-async
イベント ループまたはコルーチンでの呼び出しのブロック、スケジューラーでの CPU 作業、非同期での同期、タイムアウトの欠落
/監査スキーマ設計
FK 列とホット パスのインデックスの欠落、実際の外部キーのない ORM のみのリレーションシップ、デフォルトの ON DELETE、アプリ コードのみの整合性ルール、フロート マネー
/監査ステートレス性
2 番目のレプリカまたはデプロイによって壊れる状態 - メモリ内セッションとカウンター、静的変更可能な状態、ローカル ディスクのアップロード、プロセス ローカル ロックとスケジューラ
修正
スキル
適用される
/監査-修正-認証
認可、IDOR、テナント分離の検出結果の修復パターン — スコープ付きクエリ、ポリシー オブジェクト、デフォルト拒否
/監査-修正-非同期
正しい結果の修復パターン — トランザクション、送信ボックス、冪等性キー、ロック、制限付き再試行
/audit-fix-observability
可観測性ギャップの修復パターン - 構造化ログ、相関 ID、RED メトリクス、症状ベースのアラート
インストール
.agents フォルダーをプロジェクトにコピーします。これでインストール全体が完了します。
(これは単なるマークダウンであり、何も実行されません):
git clone -- Depth 1 https://github.com/danygiguere/audit-skills /tmp/audit-skills && cp -R /tmp/audit-skills/.agents your-project/
カーソルはから直接インストールすることもできます

リポジトリ リンク、および
スキル CLI:
npx スキルは danygiguere/audit-skills --all を追加します。
このリポジトリの AGENTS.md は、全 30 件の 1 ページのダイジェストです。
不変条件。そのコンテンツをプロジェクトの AGENTS.md にコピーします (必要に応じて追加します)
あなたはすでにそれを持っています - あなたのものを決して置き換えないでください）：そこにマージされると、すべてのものが提供されます
すべてのプロンプトの不変条件に対するエージェントの周囲の認識。それがなければ、
スキルは発動時にのみ発動します。そのルーティング テーブルは、
インストールされたスキルフォルダー。
Claude Code のメモ: Claude Code はまだ .agents/skills/ を読み取れません。
( anthropics/claude-code#31005 )。
それを次のように橋渡しします。
mkdir -p .claude && ln -s ../.agents/skills .claude/skills
echo ' @AGENTS.md ' > CLAUDE.md # CLAUDE.md をまだ持っていない場合
代替案: プロジェクトをここに持ち込んでください
スキルをすべてのプロジェクトにコピーする代わりに、クローンを作成できます。
Audit-skills を 1 回実行し、projects/ フォルダー内にプロジェクトをドロップします。
gitignored なので、コードが git status に表示されることはありません。
git pull (または git checkout vX.Y ) は、タッチせずにスキルを更新します
そこに置いたものは何でも。
git clone https://github.com/danygiguere/audit-skills
# 監査したいプロジェクトをプロジェクト内にドロップします/
cp -R /path/to/myproject 監査スキル/プロジェクト/myproject
次に、このリポジトリ内から監査します。
/監査プロジェクト/myproject
最新の状態を維持するには: git pull — プロジェクトは変更されません。
これは、所有していないリポジトリまたは不要なリポジトリを監査する場合に便利です。
変更する場合、またはスキルの中央コピーを 1 つ維持したい場合
プロジェクトごとに 1 つではなく。
自動 — エージェントに「このエンドポイントを確認する」/「これを監査する」ように依頼します
差分";スキルは説明に応じて発動します。
コマンド別 — /audit (完全な監査の場合)、またはトピックごと:
/audit-idor 、 /audit-injection 、 /audit-atomicity 、 … すべて
デフォルトで現在の差分を監査します。ファイル、フォルダーに名前を付けます。

または分岐します
何か他のものを監査する。
名前: 「この Webhook ハンドラーで冪等性チェックリストを実行する」。
修正 — 検出結果が確認された後: /audit-fix-authz 、
/audit-fix-async 、 /audit-fix-observability (「修正の仕組み」を参照)。
監査と修正は意図的に別のステップになっています。 /監査と
Audit-* チェックリストは検索とレポートのみを行い、コードを変更することはありません。
修正は、ユーザーが要求したときに行われます。レポートの後に「修正してください」と言うか、実行します。
Audit-fix-* コマンド。
すべての発見には利用可能な修正が含まれています。違うのは、どこに生息しているかです。
ほとんどのトピック - 修正はチェックリスト自体に含まれています。それぞれのチェックリストは、
例セクションでは、固定形状の隣に脆弱な形状を示します。のために
インジェクション、シークレット、出力エンコーディング、N+1 クエリなどのトピックの修正は次のとおりです。
機械的であり、正解は 1 つあります (クエリをパラメータ化し、シークレットを移動します)
環境に、ループの前に一括ロードします)。 「直してください」と言われると、
エージェントはその固定形状を適用します。追加のコマンドは必要ありません。
8 つのトピック - 修正はアーキテクチャ上の選択です。いくつかの調査結果は、
実際のトレードオフを伴ういくつかの有効な修正 (冪等性のバグ: 重複排除テーブル、
冪等キー、UPSERT、または絶対状態書き込み?)。それらのトピックが指し示すのは、
エージェントが以下の選択を行う手順を説明する修復プレイブック:
どちらの場合でも、監査→所見の確認→結果としての流れは同じです。

[切り捨てられた]

## Original Extract

Language- and framework-agnostic audit checklists for AI coding agents — security, correctness, and operability. Works with Claude Code, GitHub Copilot, Cursor, Codex CLI, OpenCode, and any agent that can read files. - danygiguere/audit-skills

GitHub - danygiguere/audit-skills: Language- and framework-agnostic audit checklists for AI coding agents — security, correctness, and operability. Works with Claude Code, GitHub Copilot, Cursor, Codex CLI, OpenCode, and any agent that can read files. · GitHub
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
danygiguere
/
audit-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .agents/ skills .agents/ skills .claude .claude .github/ workflows .github/ workflows demo demo projects projects scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md VERSION VERSION View all files Repository files navigation
Language- and framework-agnostic audit checklists for AI coding agents —
security, correctness, and operability. Works with Claude Code, GitHub
Copilot, Cursor, Codex CLI, OpenCode, and any agent that can read files.
Every checklist is written as invariants and detection smells , not
framework APIs, so the same content audits a Rails app, a Spring service,
or an Express API — the agent supplies the framework-specific translation.
/audit on a 20-line money handler — six bugs a static-analysis scanner
can't see, because each takes reasoning about ownership, concurrency, and
retries, not pattern-matching. Every one flagged, with severity and a fix.
AGENTS.md — a one-page digest of all 30 invariants; copy its content
into your project's AGENTS.md so every agent has it in context.
.agents/skills/audit/ — the router skill, with all 30 checklists and
remediation patterns bundled under references/ (four categories: access
& data security, input/API, correctness, operability).
.agents/skills/audit-* — thin per-topic wrapper skills so each checklist
is individually invocable ( /audit-idor , /audit-injection ,
/audit-fix-authz , …). Everything this package installs starts with
audit , so it stays grouped among your other skills.
/audit runs the full audit — it identifies what the code does and applies
every matching checklist below. Each topic is also individually invocable
(click through to read the checklist itself).
Works with any language or framework. Each checklist names eight
common ecosystems in its concept glossary (Rails, Laravel, Django, Spring,
Node, Vapor, .NET, Go) — those are recognition shortcuts, not a
support list. The invariants and detection smells are framework-free, so
the audits apply equally to Phoenix, FastAPI, Ktor, or your in-house
stack: the agent supplies the translation.
Audit
Checks for
/audit-authorization
Server-side permission checks at the point of action — privilege escalation, UI-only gating, checks on read but not on mutate
/audit-authn-session
Login, logout, and reset flows — session fixation, account enumeration, token expiry and single-use, remember-me storage
/audit-idor
Resources fetched or mutated by a request-supplied ID without verifying the requester may touch them
/audit-data-exposure
Over-exposed responses, errors, and logs — whole-model serialization, stack traces, PII
/audit-crypto
Password hashing, token randomness, constant-time comparison, homemade crypto, key handling
/audit-output-encoding
XSS — user data rendered into HTML, JS, CSS, URLs, headers, or emails without context-appropriate encoding
/audit-tenant-isolation
Cross-tenant leakage — unscoped queries, tenant-less cache keys, background jobs crossing tenants
/audit-csrf
State-changing endpoints on cookie/session auth without CSRF token or origin verification
/audit-mass-assignment
Request payloads bound wholesale onto models — writable role/owner/balance fields, denylists instead of allowlists
Input & API
Audit
Checks for
/audit-injection
SQL/NoSQL, command, template, and path injection — input concatenated into queries, shells, or templates
/audit-config
Insecure configuration — debug in production, permissive CORS, missing security headers, cookie flags
/audit-secrets
Hardcoded credentials, secrets in logs or version control, overly broad keys, no rotation path
/audit-api-validation
Boundary validation — types, bounds, allowed fields, trusting client-computed values like prices or roles
/audit-file-handling
Path traversal, unvalidated uploads, missing size limits, files served from the web root, zip-slip
/audit-ssrf
Server-side requests to user-influenced URLs — allowlists, private IP ranges, redirect re-validation; includes open redirects
/audit-parser-differentials
Inputs a validator accepts but the consumer reads differently — unanchored regexes, startswith allowlists, two URL parsers, validate-then-reparse
Correctness
Audit
Checks for
/audit-atomicity
Multi-store writes without a transaction — partial state surviving failures
/audit-idempotency
Handlers that misbehave when run twice — webhooks, payments, queue redelivery, double submits
/audit-background-work
Jobs and consumers — unbounded retries, poison messages, missing timeouts, duplicate or out-of-order delivery
/audit-state-management
Race conditions — check-then-act on shared state without locks, atomic primitives, or constraints
/audit-exception-handling
Swallowed errors, blanket catches, lost causes, missing cleanup, and wrong HTTP statuses (404 vs 403, 401, 422, 409)
/audit-discarded-async
Fire-and-forget bugs — promises, futures, or reactive publishers created but never awaited, returned, or composed; bare subscribe; cold writes that silently never run
/audit-cardinality
Operations assuming a query matches one row — UPDATE/DELETE on a non-unique column fanning out, findOne/.single() on non-unique fields, columns treated as unique without a DB constraint
Operability
Audit
Checks for
/audit-nplus1
Queries — or HTTP/cache calls — made inside loops over collections
/audit-observability
Silent failures — swallowed errors, logs without identifiers, no metric or alert path
/audit-migration-safety
Schema changes that lock tables, destructive changes without expand-contract, unbatched backfills
/audit-resource-limits
Unbounded work from input — missing pagination, size caps, rate limits, catastrophic regex
/audit-blocking-io-async
Blocking calls on event loops or coroutines, CPU work on the scheduler, sync-over-async, missing timeouts
/audit-schema-design
Missing indexes on FK columns and hot paths, ORM-only relationships without real foreign keys, defaulted ON DELETE, integrity rules only in app code, float money
/audit-statelessness
State that breaks with a second replica or a deploy — in-memory sessions and counters, static mutable state, local-disk uploads, process-local locks and schedulers
Fixes
Skill
Applies
/audit-fix-authz
Remediation patterns for authorization, IDOR, and tenant-isolation findings — scoped queries, policy objects, deny-by-default
/audit-fix-async
Remediation patterns for correctness findings — transactions, outbox, idempotency keys, locking, bounded retries
/audit-fix-observability
Remediation patterns for observability gaps — structured logging, correlation IDs, RED metrics, symptom-based alerts
Install
Copy the .agents folder into your project — that's the whole install
(it's just markdown; nothing executes):
git clone --depth 1 https://github.com/danygiguere/audit-skills /tmp/audit-skills && cp -R /tmp/audit-skills/.agents your-project/
Cursor can also install directly from the repo link, and if you use the
skills CLI :
npx skills add danygiguere/audit-skills --all .
This repo's AGENTS.md is the one-page digest of all 30
invariants. Copy its content into your project's AGENTS.md (append it if
you already have one — never replace yours): merged there, it gives every
agent ambient awareness of the invariants on every prompt; without it, the
skills only activate when triggered. Its routing table points at the
installed skills folder.
Claude Code note: Claude Code does not yet read .agents/skills/
( anthropics/claude-code#31005 ).
Bridge it with:
mkdir -p .claude && ln -s ../.agents/skills .claude/skills
echo ' @AGENTS.md ' > CLAUDE.md # if you don't already have a CLAUDE.md
Alternative: bring your project here
Instead of copying the skills into every project, you can clone
audit-skills once and drop your projects inside the projects/ folder —
it is gitignored, so your code never shows up in git status and a
git pull (or git checkout vX.Y ) updates the skills without touching
anything you put there.
git clone https://github.com/danygiguere/audit-skills
# drop any project you want to audit inside projects/
cp -R /path/to/myproject audit-skills/projects/myproject
Then audit from inside this repo:
/audit projects/myproject
To stay current: git pull — your projects are untouched.
This is useful when you want to audit a repo you don't own or don't want
to modify, or when you'd rather maintain one central copy of the skills
instead of one per project.
Automatic — ask your agent to "review this endpoint" / "audit this
diff"; the skills trigger on their descriptions.
By command — /audit for a full audit, or per topic:
/audit-idor , /audit-injection , /audit-atomicity , … All of them
audit your current diff by default; name a file, folder, or branch to
audit something else.
By name — "run the idempotency checklist on this webhook handler".
Fixes — after findings are confirmed: /audit-fix-authz ,
/audit-fix-async , /audit-fix-observability (see "How fixes work").
Audits and fixes are deliberately separate steps. /audit and the
audit-* checklists only find and report — they never change code.
Fixing happens when you ask for it: say "fix those" after a report, or run
an audit-fix-* command.
Every finding has a fix available; what differs is where it lives:
Most topics — the fix is in the checklist itself. Each checklist's
Example section shows the vulnerable shape next to the fixed shape. For
topics like injection, secrets, output encoding, or N+1 queries, the fix is
mechanical and has one right answer (parameterize the query, move the secret
to the environment, bulk-load before the loop). When you say "fix it", the
agent applies that fixed shape — no extra command needed.
Eight topics — the fix is an architectural choice. Some findings have
several valid fixes with real trade-offs (an idempotency bug: dedupe table,
idempotency key, UPSERT, or an absolute-state write?). Those topics point to
a remediation playbook that walks the agent through choosing:
Either way, the flow is the same: audit → confirmed findings → as

[truncated]
