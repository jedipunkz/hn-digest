---
source: "https://fast-and-flow-production.onrender.com/case-study"
hn_url: "https://news.ycombinator.com/item?id=49359043"
title: "A multi-tenant SaaS built in a 92-hour AI-augmented engineering sprint"
article_title: "FAST & FLOW Software"
image: ""
author: "caredeo"
captured_at: "2026-08-19T09:23:55Z"
capture_tool: "hn-digest"
hn_id: 49359043
score: 1
comments: 0
posted_at: "2026-08-19T09:22:25Z"
tags:
  - hacker-news
  - translated
---

# A multi-tenant SaaS built in a 92-hour AI-augmented engineering sprint

- HN: [49359043](https://news.ycombinator.com/item?id=49359043)
- Source: [fast-and-flow-production.onrender.com](https://fast-and-flow-production.onrender.com/case-study)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T09:22:25Z

## Translation

タイトル: 92 時間の AI 拡張エンジニアリング スプリントで構築されたマルチテナント SaaS
記事のタイトル: FAST & FLOW ソフトウェア

記事本文:
FAST & FLOW ソフトウェア FAST & FLOW · テクニカル エンジニアリング ケーススタディ
構造を維持したエンジニアリング圧縮。
AI で強化されたマルチテナント B2B SaaS は、異常に圧縮されたエンジニアリング サイクルの下で構築されています。
本番ソース リポジトリは非公開のままです。この公開ケーススタディでは、資格情報、独自のソースコード、またはクライアントデータを公開することなく、アーキテクチャ、方法論、測定されたエンジニアリングチェックポイント、制限、およびテスト戦略を明らかにします。
これらの数値は独立して検証可能なエンジニアリング チェックポイントに属しており、すべての数値がまったく同じ Git SHA から生じたものであるかのようには表示されません。
スタックとアーキテクチャの選択
アプリケーション ORM はありません。永続化では、Supabase クライアント、明示的な SQL 移行、および狭い PostgreSQL RPC を使用します。
標準的なスタイリング レイヤーとして Tailwind はありません。製品設計システムは、ネイティブ CSS カスタム プロパティとセマンティック React プリミティブを使用します。
これらはプロジェクト固有のアーキテクチャ上の選択であり、普遍的な推奨事項ではありません。
プラットフォーム
→ 組織
→ 所在地
→ プロジェクト
→ 製品/モジュール データ組織はプライマリ テナントとセキュリティ境界です。場所とプロジェクトは下位スコープです。
スーパーベース認証
→ アプリケーションユーザー
→ 会員登録
→ 組織
→ 信頼できるテナントコンテキスト
→ 許可エンジン
→ サーバー運用
→ PostgreSQL / RPC
→ 行レベルのセキュリティ 認証 ≠ 認可 メンバーシップ ≠ 権限 サーバー認可 ≠ RLS RLS ≠ サーバー認可 ブラウザーが提供するテナント識別子は、権限ではなく意図を表します。サーバーは、結果的な操作の前に、現在の ID、メンバーシップ、テナント コンテキスト、およびアクセス許可を個別に解決します。
AI コンテキストの制限がどのように処理されたか
AI コンテキスト = 使い捨てのワーキングメモリ。 Git + 不変 SQL 移行 + テスト + コントラクト + ホストされたデータb

ase 状態 = 耐久性のあるエンジニアリング メモリ。
このプロセスは、移行 40 の実装中に移行 1 を記憶するモデルには依存しませんでした。以前のエンジニアリング上の決定は、不変の移行履歴、アーキテクチャ契約、正確な SHA、CI ゲート、テスト、役割と権限のレジストリ、スキーマのフィンガープリント、ホストされた検証証拠などの正規のリポジトリ成果物に外部化されました。
ドメインを拡張する前に、関連する状態がそれらのアーティファクトから再構築されました。モデルは記録システムではありませんでした。エンジニアリングシステムはそうでした。
RLS とテナント間の分離
RLS 検証は、SQL ファイルにポリシーが存在するかどうかを確認するだけではありません。ホストされた Supabase 環境は、独立したテナント設備を使用して実行されました。
ユーザーA → 組織A → 許可
ユーザー A → 組織 B → 拒否 / 許可されていない行なし
ユーザーB → 組織B → 許可
ユーザー B → 組織 A → 拒否 / 許可されていない行なし 追加の敵対的制御が含まれます。
クロス環境 JWT の拒否
グローバル スーパー管理フィクスチャは、テナント全体を正当に認識できるため、分離の失敗をマスクします。合成プラットフォーム当局によるテナントローカルのネガティブコントロールの汚染が防止されました。
アトミックセキュリティの突然変異 + 監査
現在のアクターを検証する
→ テナントの検証
→ 役割割り当て権限を検証する
→ 役割の割り当てを挿入
→RoleAssigned監査イベントを追加
→ COMMIT 特権付きの変更と必要な監査追加は、同じ PostgreSQL トランザクションで実行されます。監査の追加が失敗した場合、PostgreSQL は変更をロールバックします。監査履歴は追加指向です。直接的な履歴の UPDATE および DELETE は拒否されます。
91,344 の追跡されたラインが実際に意味するもの
追跡された 91,344 行は、手書きの実稼働アプリケーション コードが 91,344 行あることを意味するものではありません。
カテゴリの内訳は分析的な推定値であり、正確な値ではありません。

cloc 認定のサブカウントのセット。
50 の前方のみの PostgreSQL 移行
リリース 1.5 チェックポイントでは、正規のリポジトリ制御チェーンにはちょうど 50 個の移行が含まれていました。歴史的な移行は不変です。修正は前方のみです。
移行ガバナンスでは、決定的なファイル名、UTC タイムスタンプ、厳密な時系列順序、移行セットの予期される承認、および履歴の不変性がチェックされます。既存の移行は、ロールバックのショートカットとして変更、削除、または名前変更することはできません。
メイン リポジトリ スイートは、ネイティブ Node.js テスト ランナーを使用します。
ノード --test シールされた Production Readiness チェックポイント: 527 個のテスト、527 個が合格、0 個が失敗しました。対象範囲には、ドメイン契約、ライフサイクル移行、認可境界、移行ガバナンス、敵対的なセキュリティ回帰、テナントの分離、ホストされた検証、および運用動作が含まれます。
リクエスト: 100
合格: 100
失敗: 0
同時実行数: 10
p50: 89.1ミリ秒
p95: 244.5 ミリ秒
最大: 289.0 ミリ秒
p95 しきい値: 1,500 ミリ秒 これは、制御されたパイロットのベースラインであり、最大スケールまたはハイパースケールのベンチマークではありません。数千の同時テナントについては一切の請求は行われません。
参照生産性シナリオ
エンジニア 2 名 × 8 か月 × 180 エンジニアリング時間/月 = 2,880 基準人間工学時間
読者が別のベースラインを受け入れる、拒否する、または置き換えることができるように、仮定が明示的に示されています。測定された 92 時間のプロジェクト サイクルと、仮定の 2,880 時間の比較シナリオを混同しないでください。
この事例研究が主張していないこと
91,344 行の手書きの製品コードは主張していません。
LOC がソフトウェアの品質を証明するとは主張しません。
92時間が検証された世界記録であるとは主張していない。
100 リクエストがハイパースケールのベンチマークであるとは主張しません。
AI が工学的な判断を排除したとは主張していない

。
コンテキストウィンドウの制限がなくなったとは主張しません。
すべてのエンジニアリング組織が同じ速度を再現できるとは主張しません。
興味深い主張は、AI がコードを速く書くということではありません。
エンジニアリング実験では、マルチテナンシー、認証、RBAC、RLS、監査可能性、前方専用データベースの進化、運用ワークフロー、CI ゲート、ホストされた敵対的検証、および測定可能なパフォーマンス証拠を維持しながら開発時間を短縮できるかどうかを検討します。
FAST & FLOW は、Git、テスト、移行、セキュリティ境界、CI、再現可能な証拠によって管理されるシステム内で AI が動作する場合、大幅なエンジニアリング圧縮が可能である可能性があることを示唆しています。

## Original Extract

FAST & FLOW Software FAST & FLOW · TECHNICAL ENGINEERING CASE STUDY
Engineering compression with preserved structure.
An AI-augmented, multi-tenant B2B SaaS built under an unusually compressed engineering cycle.
The production source repository remains private. This public case study exposes architecture, methodology, measured engineering checkpoints, limitations, and testing strategy without exposing credentials, proprietary source code, or client data.
These figures belong to independently verifiable engineering checkpoints and are not presented as if every number originated from the exact same Git SHA.
Stack and architectural choices
No application ORM. Persistence uses Supabase clients, explicit SQL migrations, and narrow PostgreSQL RPCs.
No Tailwind as the canonical styling layer. The product design system uses native CSS custom properties and semantic React primitives.
These are project-specific architectural choices, not universal recommendations.
Platform
→ Organization
→ Location
→ Project
→ Product / Module Data Organization is the primary tenant and security boundary. Location and Project are subordinate scopes.
Supabase Auth
→ Application User
→ Membership
→ Organization
→ Trusted Tenant Context
→ Permission Engine
→ Server Operation
→ PostgreSQL / RPC
→ Row Level Security Authentication ≠ authorization Membership ≠ permission Server authorization ≠ RLS RLS ≠ server authorization Browser-supplied tenant identifiers express intent, not authority. The server independently resolves current identity, membership, tenant context, and permissions before consequential operations.
How the AI context limit was handled
AI context = disposable working memory. Git + immutable SQL migrations + tests + contracts + hosted database state = durable engineering memory.
The process did not depend on a model remembering migration 1 while implementing migration 40. Earlier engineering decisions were externalized into canonical repository artifacts: immutable migration history, architecture contracts, exact SHAs, CI gates, tests, role and permission registries, schema fingerprints, and hosted verification evidence.
Before extending a domain, the relevant state was reconstructed from those artifacts. The model was not the system of record; the engineering system was.
RLS and cross-tenant isolation
RLS validation went beyond checking that policies existed in SQL files. Hosted Supabase environments were exercised with independent tenant fixtures.
User A → Organization A → ALLOW
User A → Organization B → DENY / no unauthorized rows
User B → Organization B → ALLOW
User B → Organization A → DENY / no unauthorized rows Additional adversarial controls included:
cross-environment JWT rejection
A global Super Admin fixture can legitimately see across tenants and therefore mask isolation failures. Synthetic platform authority was prevented from contaminating tenant-local negative controls.
Atomic security mutations + audit
validate current actor
→ validate tenant
→ validate role-assignment authority
→ INSERT role assignment
→ append RoleAssigned audit event
→ COMMIT Privileged mutation and required audit append execute in the same PostgreSQL transaction. If the audit append fails, PostgreSQL rolls back the mutation. Audit history is append-oriented; direct historical UPDATE and DELETE are rejected.
What 91,344 tracked lines actually means
91,344 tracked lines does not mean 91,344 lines of hand-written production application code.
The category breakdown is an analytical estimate, not a cloc-certified set of subcounts.
50 forward-only PostgreSQL migrations
At the Release 1.5 checkpoint, the canonical repository-controlled chain contained exactly 50 migrations. Historical migrations are immutable; corrections are forward-only.
Migration governance checks deterministic filenames, UTC timestamps, strict chronological ordering, expected authorization of the migration set, and historical immutability. Existing migrations cannot be modified, deleted, or renamed as a shortcut for rollback.
The main repository suite uses the native Node.js test runner:
node --test At the sealed Production Readiness checkpoint: 527 tests, 527 passed, 0 failed. Coverage included domain contracts, lifecycle transitions, authorization boundaries, migration governance, adversarial security regressions, tenant isolation, hosted verification, and operational behavior.
Requests: 100
Passed: 100
Failed: 0
Concurrency: 10
p50: 89.1 ms
p95: 244.5 ms
max: 289.0 ms
p95 threshold: 1,500 ms This is a controlled pilot baseline, not a maximum-scale or hyperscale benchmark. No claim is made about thousands of concurrent tenants.
Reference productivity scenario
2 engineers × 8 months × 180 engineering hours/month = 2,880 reference human-engineering hours
The assumptions are shown explicitly so a reader can accept, reject, or substitute a different baseline. The measured 92-hour project cycle and the hypothetical 2,880-hour comparison scenario should not be confused.
What this case study does not claim
It does not claim 91,344 lines of hand-written production code.
It does not claim that LOC proves software quality.
It does not claim that 92 hours is a verified world record.
It does not claim that 100 requests is a hyperscale benchmark.
It does not claim that AI eliminated engineering judgment.
It does not claim that context-window limits disappeared.
It does not claim that every engineering organization would reproduce the same speed.
The interesting claim is not that AI writes code fast.
The engineering experiment is whether development time can be compressed while preserving multi-tenancy, authentication, RBAC, RLS, auditability, forward-only database evolution, operational workflows, CI gates, hosted adversarial verification, and measurable performance evidence.
FAST & FLOW suggests that substantial engineering compression may be possible when AI operates inside a system governed by Git, tests, migrations, security boundaries, CI, and reproducible evidence.
