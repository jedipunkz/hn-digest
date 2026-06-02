---
source: "https://devarch.ai"
hn_url: "https://news.ycombinator.com/item?id=48363883"
title: "Without Intelligent Guardrails, Claude Code Is Pure Chaos"
article_title: "DevArch — AI-Assisted Development, Disciplined"
author: "ChicagoDave"
captured_at: "2026-06-02T00:30:07Z"
capture_tool: "hn-digest"
hn_id: 48363883
score: 2
comments: 1
posted_at: "2026-06-01T23:22:35Z"
tags:
  - hacker-news
  - translated
---

# Without Intelligent Guardrails, Claude Code Is Pure Chaos

- HN: [48363883](https://news.ycombinator.com/item?id=48363883)
- Source: [devarch.ai](https://devarch.ai)
- Score: 2
- Comments: 1
- Posted: 2026-06-01T23:22:35Z

## Translation

タイトル: インテリジェントなガードレールがなければ、クロード コードは純粋なカオスです
記事のタイトル: DevArch — AI 支援開発、規律ある開発
説明: クロード コードのディレクティブ、エージェント、およびエンジニアリング規律を強制するスキル。セッションの継続性、品質ゲート、動作テスト、ADR はすべて自動です。

記事本文:
____ _ _
| _ \ _____ __/ \ _ __ ___| |__
| | | |/ _ \ \ / / _ \ | '__/__| '_ \
| |_| | __/\ V / ___ \| | | (__| | | |
|___/ \___| \_/_/ \_\_| \___|_| |_|
v4.0.0
Claude Opus 4.8 で検証済み →
クロード・コードは強力です。
規律がなければ、それは混乱です。
DevArch は、あらゆるステップでエンジニアリング規律を強制する、Claude Code のディレクティブ、エージェント、およびスキルのセットです。セッションの継続性、品質ゲート、動作テスト、アーキテクチャ決定記録、ドメイン駆動設計はすべて自動で行われます。対象分野の専門家、1 人の開発者、および QA がチーム全体を置き換えることができます。
AI により、孤独な開発者が高速化されます。 DevArch により、チームは規律を保てるようになります。部屋に専門家が存在することで、アーキテクチャの一貫性と品質の向上が維持されるため、小規模なチームでも大規模なチームと同様の成果を上げることができます。
ドメイン駆動設計は、一人の人の頭の中に存在すべきではありません。 DevArch は、境界のあるコンテキストを見つけ、ユビキタス言語に名前を付け、境界を保持するのに役立ちます。これにより、コード行を記述する前にシステムの形状が決まります。
「BUILD: FAILED」からすべて緑色になるまで
動作テスト、突然変異チェック、品質ゲートにより、赤いビルドが合格ビルドに変わります。テストによってそれが証明されます。あなたが仕事を指揮します。 DevArch は規律を維持するため、品質が低下するのではなく、強化されます。
DevArch は Claude Code のライフサイクルにフックします。すべてのセッションは自動的にこの規律に従います。
すべてのフックは自動的に起動します。開発者がそれらを呼び出すことはありません。セッションのコンテキストは会話全体で保持されます。明日新しいセッションを開始すると、監査エージェントがどこから中断したかを正確に教えてくれます。
セッションは内部ループです。スキル、成果物、意思決定記録がそれらを一貫したプロジェクトに結び付けます。
構築する前にブレインストーミングを行います。アーキテクチャを定期的に見直してください。すべてのセッションで概要が作成されます。あらゆる重要な決定

ADRとなります。セッション間で失われるものは何もありません。
開発者 Claude コード DevArch フック DevArch エージェント
| | | |
|セッションを開始する | | |
|---------------------->| | |
| | SessionStart フック | |
| |---------------------->| |
| | |セッションファイルを作成する |
| | |前のセッションを検索|
| | |計画の状態を報告する |
| | |品質ゲートを設定する |
| |セッション前監査 | |
| |---------------------------------------------->|
| | |過去 5 つのセッションを読む
| | |表面ブロッカー |
| | |再発するバグにフラグを立てます|
|監査説明会 | | | |
| |セッションプランナー | |
| |---------------------------------------------->|
| | |段階に分解する
|計画の見直し | | | |
| | | |
| |プランゲート出口 | |
| |---------------------->|ゲート = クリア |
| | |アーカイブ計画 |
| | | |
: (実装) : : :
| | | |
| | PostToolUseフック | (ツール呼び出しごとに) |
| |---------------------->| |
| | | work-log.txt にログを記録します。
| | |予算の使用状況を追跡 |
| | | 70/90/100% でレポート |
| | | |
| |突然変異の検証 | |
| |---------------------------------------------->|
| | |チェックテストアサート |
| | |実際の状態の変化について
| | | |
| 「発送してください」 | | |
|---------------------->| | |
| |作品概要ライター | |
| |---------------------------------------------->|
| | |セッションを要約する |
| | |ステータス + ブロッカー |
| |プレコンパクトフック | |
| |---------------------->|作業ログを追加 |
| | |セッションを終了する |
| | | |
| |コミット + プッシュ | |
|完了 |
何が入っているのか
エンジニアリング規律を自動的に適用するディレクティブ、エージェント、およびスキル。
セッション開始フックはコンテキストを復元します。作業概要は状態をキャプチャします。セッション前の監査の表面ブロッカー。自分がどこにいるのかを見失ってしまうことはありません。
コードを記述する前に不変条件とバリアントを特定します。凝集したモジュール性が強制されます。ドメイン ロジックとインフラストラクチャの間の境界を明確にします。提案ではありません — r

ウレス。
境界を越えた状態 (ストア、リデューサー、プロジェクション、ドメイン モジュール) を所有するファイルを編集する前に、エージェントは OWNER、SHARED?、PROMISE、および ALTERNATIVES を宣言します。コンシューマごとの状態は、明示的な会話なしに共有モジュールに静かに移行することはできません。
すべての副作用関数は、テストが作成される前に、構造化されたステートメント (DOES、WHEN、BECAUSE、REJECTS WHEN) を取得します。コードが何をするのかを説明できない場合、そのコードは準備ができていません。
行動ステートメントから派生したテストは、RED/YELLOW/GREEN で評価されます。トートロジー アサーション、モックのみのチェック、および戻り値のみのテストは捕捉され、拒否されます。
エージェントはすべての副作用関数の後に実行され、値を返すだけでなく、「スローしなかった」だけでなく、実際の状態の変更に対してテストがアサートされたことを検証します。
統合にちなんで名付けられたフェーズには、実際のサブプロセス、ランタイム、または移行を推進するテストが必要です。所有されている依存関係のスタブは、その依存関係との統合の受け入れゲートとなることはできません。 「カーブアウト」は「完全」の有効な修飾語ではありません。
決定によって今後のセッションが制限される場合、エージェントは「ADR に値しますか?」と尋ねます。承認された決定は、コンテキスト、根拠、結果とともに永続的な記録として保存されます。
重要な目標は、実装を開始する前にフェーズに分解されます。予算追跡では、フェーズごとのツール呼び出し予算の 70%、90%、および 100% で警告が表示されます。
複雑さが必要な場合にアクティブ化します。境界のあるコンテキスト、ドメイン イベント、集計、値オブジェクト、ユビキタス言語は、儀式ではなく会話を通じてキャプチャされます。
すべてのフックは Bash と PowerShell の両方で出荷されます。同じフック サーフェスは Windows、macOS、Linux で利用できます。
すべてのセッションは、一意の 6 文字の 16 進数 ID を取得します。ゲート ファイル、ツール呼び出しバジェット、およびファイル変更追跡はその ID によって名前空間化されるため、同じ環境内で 2 つの同時クロード コード セッションが実行されます。

po は共有状態では決して衝突しません。
1 つの製品にすべてが含まれています。年間ライセンス。シート単位で価格が設定されます。
チーム全体で 1 つのライセンス。
大規模なチーム、請求書発行、カスタム条件。
Enterprise シート、ボリューム ライセンス、またはプロジェクトに関するご質問については、お気軽にお問い合わせください。
© 2026 デビッド・コーネルソン。無断転載を禁じます。・ライセンス

## Original Extract

Claude Code directives, agents, and skills that enforce engineering discipline. Session continuity, quality gates, behavioral tests, ADRs — all automatic.

____ _ _
| _ \ _____ __/ \ _ __ ___| |__
| | | |/ _ \ \ / / _ \ | '__/ __| '_ \
| |_| | __/\ V / ___ \| | | (__| | | |
|____/ \___| \_/_/ \_\_| \___|_| |_|
v4.0.0
Validated for Claude Opus 4.8 →
Claude Code is powerful.
Without discipline, it's chaos .
DevArch is a set of directives, agents, and skills for Claude Code that enforce engineering discipline at every step. Session continuity, quality gates, behavioral testing, architecture decision records, and domain-driven design — all automatic. A subject matter expert, a single developer, and QA can replace an entire team.
AI makes a lone developer fast. DevArch makes them disciplined — the expert presence in the room that keeps architecture coherent and quality compounding, so a small team delivers like a big one.
Domain-driven design shouldn't live in one person's head. DevArch helps you find the bounded contexts, name the ubiquitous language, and hold the boundaries — so the system has a shape before you write a line of code.
From BUILD: FAILED to all green
Behavioral tests, mutation checks, and quality gates turn a red build into a passing one — with the tests that prove it. You direct the work; DevArch keeps the discipline, so quality compounds instead of eroding.
DevArch hooks into Claude Code's lifecycle. Every session follows this discipline automatically.
Every hook fires automatically. The developer never invokes them. Session context is preserved across conversations — start a new session tomorrow and the audit agent tells you exactly where you left off.
Sessions are the inner loop. Skills, artifacts, and decision records connect them into a coherent project.
Brainstorm before you build. Review the architecture periodically. Every session produces a summary. Every significant decision becomes an ADR. Nothing is lost between sessions.
Developer Claude Code DevArch Hooks DevArch Agents
| | | |
| start session | | |
|---------------------->| | |
| | SessionStart hook | |
| |---------------------->| |
| | | create session file |
| | | find previous session|
| | | report plan state |
| | | set quality gate |
| | pre-session-audit | |
| |---------------------------------------------->|
| | | read last 5 sessions
| | | surface blockers |
| | | flag recurring bugs|
| audit briefing | | | |
| | session-planner | |
| |---------------------------------------------->|
| | | decompose into phases
| review plan | | | |
| | | |
| | plan-gate-exit | |
| |---------------------->| gate = clear |
| | | archive plan |
| | | |
: (implementation) : : :
| | | |
| | PostToolUse hook | (on every tool call) |
| |---------------------->| |
| | | log to work-log.txt |
| | | track budget usage |
| | | report at 70/90/100% |
| | | |
| | mutation-verification | |
| |---------------------------------------------->|
| | | check tests assert |
| | | on real state change
| | | |
| "ship it" | | |
|---------------------->| | |
| | work-summary-writer | |
| |---------------------------------------------->|
| | | summarize session |
| | | status + blockers |
| | PreCompact hook | |
| |---------------------->| append work log |
| | | finalize session |
| | | |
| | commit + push | |
| done |
What's Inside
Directives, agents, and skills that enforce engineering discipline automatically.
Session start hooks restore context. Work summaries capture state. Pre-session audits surface blockers. You never lose track of where you are.
Invariants and variants identified before writing code. Cohesive modularity enforced. Clear boundaries between domain logic and infrastructure. Not suggestions — rules.
Before editing files that own cross-boundary state — stores, reducers, projections, domain modules — the agent declares OWNER, SHARED?, PROMISE, and ALTERNATIVES. Per-consumer state can't quietly migrate into a shared module without an explicit conversation.
Every side-effect function gets a structured statement — DOES, WHEN, BECAUSE, REJECTS WHEN — before any test is written. If you can't state what the code does, it's not ready.
Tests derived from behavior statements are graded RED/YELLOW/GREEN. Tautological assertions, mock-only checks, and return-value-only tests are caught and rejected.
An agent runs after every side-effect function to verify tests assert on actual state changes — not just return values, not just "didn't throw."
Phases named after an integration must have a test that drives the real subprocess, runtime, or migration. A stub of an owned dependency can't be the acceptance gate for the integration with that dependency. "Carve-out" is never a valid modifier of "complete."
When a decision constrains future sessions, the agent asks "ADR-worthy?" Approved decisions are promoted to permanent record with context, rationale, and consequences.
Non-trivial goals are decomposed into phases before implementation starts. Budget tracking warns at 70%, 90%, and 100% of tool-call budget per phase.
Activate when complexity demands it. Bounded contexts, domain events, aggregates, value objects, ubiquitous language — captured through conversation, not ceremony.
Every hook ships in both Bash and PowerShell — the same hook surface available on Windows, macOS, and Linux.
Every session gets a unique 6-character hex ID. Gate files, tool-call budgets, and file change tracking are namespaced by that ID, so two concurrent Claude Code sessions in the same repo never collide on shared state.
One product, everything included. Annual license, priced by seats.
One license for your whole team.
Larger teams, invoicing, custom terms.
Enterprise seats, volume licensing, or questions about your project — drop us a line.
© 2026 David Cornelson. All rights reserved. · License
