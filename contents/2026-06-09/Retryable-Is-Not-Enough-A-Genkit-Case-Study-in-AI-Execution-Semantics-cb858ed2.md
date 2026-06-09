---
source: "https://blog.pitstop.dev/i-filed-a-github-issue-against-googles-ai-framework-heres-what-happened/"
hn_url: "https://news.ycombinator.com/item?id=48463348"
title: "Retryable Is Not Enough: A Genkit Case Study in AI Execution Semantics"
article_title: "I Filed a GitHub Issue Against Google's GenKit AI Framework"
author: "SirBrenton"
captured_at: "2026-06-09T18:52:42Z"
capture_tool: "hn-digest"
hn_id: 48463348
score: 2
comments: 0
posted_at: "2026-06-09T16:33:30Z"
tags:
  - hacker-news
  - translated
---

# Retryable Is Not Enough: A Genkit Case Study in AI Execution Semantics

- HN: [48463348](https://news.ycombinator.com/item?id=48463348)
- Source: [blog.pitstop.dev](https://blog.pitstop.dev/i-filed-a-github-issue-against-googles-ai-framework-heres-what-happened/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T16:33:30Z

## Translation

タイトル: 再試行可能だけでは不十分: AI 実行セマンティクスにおける Genkit のケーススタディ
記事のタイトル: Google の GenKit AI フレームワークに対して GitHub で問題を提起しました
説明: 再試行ミドルウェアにシグナルがありました。分類は存在していました。回復パスが存在しました。欠けていたのは、それらの間の伝播チャネルでした。

記事本文:
サインイン
購読する
ブレント・ウィリアムズ
で
シグナルインテグリティ
—
2026 年 5 月 21 日
Google の AI フレームワークに対して GitHub に問題を提出しました。これが何が起こったのかです。
更新: このケースのより正式な伝播レポートがここで入手できるようになりました。
Retry-After が失われた場合: AI 実行のセマンティック劣化のケーススタディ
3週間前、私は genkit-ai/genkit に対して問題を提起しました。
バグレポートではありません。発見です。違いがあります。
発見されたのは、Genkit の再試行ミドルウェアのデフォルトの再試行ステータスに RESOURCE_EXHAUSTED が含まれているということでした。 Anthropic が Retry-After: 60 を指定して HTTP 429 を返すと、ミドルウェアはクールダウン ウィンドウ内で 1000 ミリ秒間隔で起動します。ウィンドウが自然に期限切れになるまで、すべての再試行は失敗します。
このパターンは前にも見たことがあります。それはコーパスにあります。
予想外だったのは、次に何が起こったかです。
Michael Doyle 氏は、現実世界の例を求めました。私は彼に 1 つを渡しました。Anthropic の実際の 429 応答ヘッダーです。彼は、Genkit のコアメンテナである Pavel Gajdošík にエスカレートしました。
パベル氏は問題を認めただけではありません。彼はスレッド内でアーキテクチャに関するディスカッションを開始しました。
GenkitError を通じて Retry-After シグナルを伝播する方法の 3 つのオプション。オプション 1: 専用の retryAfterMs フィールド。オプション 2: 既存の詳細フィールドの規則。オプション 3: 新しい responseMetadata バッグ。
スレッドは開いたままでした。 2 日後、彼は PR #5343 をリンクしました。
11 ファイルにまたがる 358 行。 4 つのプロバイダーの統合: Anthropic、OpenAI、Google GenAI、Vertex AI。
これが重要な部分です。
修正は「Retry-After ヘッダーを確認する」ことではありませんでした。この号を読んでいるエンジニアなら誰でもそれを知っていました。
修正はトランスポート メカニズムを追加することでした。
応答メタデータ.retryAfterMs
GenkitError の専用フィールド。再試行タイミングのセマンティクスをフレームワークの境界を越えて再試行ミドルウェアに伝達します。
プロバイダー

信号を正しく送りました。 HTTP 層はそれを正しく配信しました。フレームワークは、状態を RESOURCE_EXHAUSTED として正しく分類しました。
それでも、システムは依然として失敗しました。再試行タイミングのセマンティクスが境界を越えても生き残れなかったためです。ミドルウェアにはそれらを受信するチャネルがありませんでした。
信号は存在していました。分類は存在していました。回復パスは存在していました。欠けていたのは、それらの間の伝播チャネルでした。
それは再試行バグではありません。それはシグナルインテグリティの欠陥です。
AI システムが拡大するにつれてこれがさらに悪化する理由
単層システムでは大きな障害が発生します。エラーが表示されます。電話を直すのはあなたです。
階層化されたシステムは静かに失敗します。プロバイダー SDK は HTTP 応答を抽象化します。ミドルウェア フレームワークは SDK エラーを抽象化します。ランタイムはミドルウェアの状態を抽象化します。オーケストレーション層は実行時の動作を抽象化します。ツールまたは UI の境界で障害が表面化するまでに、それを説明し、正しい応答を制御するシグナルが、4 回または 5 回の変換によって劣化している可能性があります。
Genkit は、レイヤーが表示され、修正が公開されているため、クリーンな例です。しかし、同じ構造的病理がコーパス全体に現れます。
再試行 - 正しく解析された後、シリアル化境界で失われた
クォータ信号は正しく分類され、間違ったスコープで適用されました
STOP 条件が内部で到達しましたが、親セッションには伝播されませんでした
ランタイムにレート制限信号が存在しますが、CLI サーフェイスでは動作しません
信号を必要とする層には信号がありません。または、使用できない形式になっています。
それらはすべてシグナルインテグリティ障害です。失敗パターンも同じです。境界だけが変わります。
Pavel の PR は再試行を追加しませんでした。フレームワークの境界を越えてクールダウン セマンティクスを伝える構造化フィールドが追加されました。その後、再試行ロジックがその機能を実行できるようになります。
その区別が重要です。 T

彼は修正自体が論文を検証します:
再試行可能性の分類と再試行タイミングのセマンティクスは、一緒に伝播する必要があります。
それはヒントではありません。それはシステム不変です。そしてそれは Genkit 以外にも当てはまります。
この発見に対する正規のレシートは現在、pitstop-truth にあります。
PT-2026-05-20-github-genkit-ai-genkit-5270-hidden
非表示 — 無視されません。修正によって伝播チャネルが追加されるまで、信号は決定層に存在しませんでした。この違いにより、さまざまな修正が行われます。
Genkit からのより深い教訓は次のとおりです。
現代の AI システムの失敗は、シグナルが存在しないからではなく、層を越えるにつれて実行に重要なセマンティクスが低下するために増えています。
生き残るシステムは、信号の完全性をエンドツーエンドで維持するシステムになります。
Brent Williams は、AI/API パイプラインの実行信頼性インフラストラクチャである Pitstop を構築します。
レシート64枚。 35 を超える実稼働システム。
github.com/SirBrenton/pitstop-truth
```json
{
"post": "Google の AI フレームワークに対して GitHub の問題を提出しました。何が起こったかを説明します。",
"公開済み": "2026-05-21",
"thesis": "AI 実行システムは、意思決定に関連する信号がレイヤー間で整合性を失ったときに失敗します。再試行ロジックが欠落しているためではありません",
"プライマリ検索": {
"リポジトリ": "genkit-ai/genkit",
「問題」: 5270、
「pr」: 5343、
"pr_lines": 358、
"providers_fixed": ["Anthropic"、"OpenAI"、"Google GenAI"、"Vertex AI"],
"signal_failure_type": "非表示",
"fix_mechanism": "responseMetadata.retryAfterMs 伝播チャネルが GenkitError に追加されました"
}、
"receipt_id": "PT-2026-05-20-github-genkit-ai-genkit-5270-hidden",
"key_invariant": "再試行可能性の分類と再試行タイミングのセマンティクスは一緒に伝播する必要があります",
"canonical_publication": {
"title": "Retry-After が失われた場合: AI 実行セマンティック劣化のケーススタディ",
"url": "https://github.com/pitstop-hq/pitstop/blob/main/publ

ications/genkit-retry-after-propagation-report.md",
"タイプ": "伝播レポート",
"投稿への関係": "形式化されたケース分析",
"目的": "Genkit Retry-After 伝播結果、証拠境界、操作上の不変条件、およびより広範な意味論的劣化パターンに関する正規の公開リファレンス。"
}、
"コーパス": "https://github.com/pitstop-hq/pitstop/tree/main/corpus",
「連絡先」: " [email protected] "
}
「」
前号
私のコーパスが私よりも賢い理由
シェアする
ピットストップを購読する
[メールで保護されています]
購読する
ピットストップ © 2026
サインアップ

## Original Extract

Retry middleware had the signal. Classification existed. Recovery path existed. What was missing was the propagation channel between them.

Sign in
Subscribe
By Brent Williams
in
signal-integrity
—
21 May 2026
I Filed a GitHub Issue Against Google’s AI Framework. Here’s What Happened.
Update: A more formal propagation report for this case is now available here:
When Retry-After Gets Lost: A Case Study in AI Execution Semantic Degradation
Three weeks ago I filed an issue against genkit-ai/genkit .
Not a bug report. A finding. There’s a difference.
The finding was this: the retry middleware in Genkit includes RESOURCE_EXHAUSTED in its default retry statuses. When Anthropic returns HTTP 429 with Retry-After: 60 , the middleware fires at 1000ms intervals — inside the cooldown window. All retries fail until the window expires naturally.
I’d seen this pattern before. It’s in the corpus .
What I didn’t expect was what happened next.
Michael Doyle asked for a real-world example . I gave him one: Anthropic’s actual 429 response headers. He escalated to Pavel Gajdošík, a core Genkit maintainer.
Pavel didn’t just acknowledge the issue. He opened an architecture discussion in the thread.
Three options for how to propagate the Retry-After signal through GenkitError. Option 1: a dedicated retryAfterMs field. Option 2: convention in the existing detail field. Option 3: a new responseMetadata bag.
The thread stayed open. Two days later he linked PR #5343 .
358 lines across 11 files. Four provider integrations: Anthropic, OpenAI, Google GenAI, Vertex AI.
This is the part that matters.
The fix wasn’t “check the Retry-After header.” Every engineer reading the issue knew that.
The fix was adding a transport mechanism:
responseMetadata.retryAfterMs
A dedicated field on GenkitError that carries retry timing semantics across the framework boundary into the retry middleware.
The provider emitted the signal correctly. The HTTP layer delivered it correctly. The framework classified the condition correctly as RESOURCE_EXHAUSTED .
And yet the system still failed — because retry timing semantics didn’t survive the boundary crossing. The middleware had no channel to receive them.
The signal existed. The classification existed. The recovery path existed. What was missing was the propagation channel between them.
That’s not a retry bug. That’s a signal integrity failure.
Why this compounds as AI systems scale
A single-layer system fails loudly. You see the error. You fix the call.
A layered system fails silently. Provider SDKs abstract HTTP responses. Middleware frameworks abstract SDK errors. Runtimes abstract middleware state. Orchestration layers abstract runtime behavior. By the time a failure surfaces at the tool or UI boundary, the signal that would explain it — and govern the correct response — may have degraded through four or five transformations.
Genkit is a clean example because the layers are visible and the fix is public. But the same structural pathology appears across the corpus:
Retry-After parsed correctly, lost at the serialization boundary
Quota signal classified correctly, enforced at the wrong scope
STOP condition reached internally, not propagated to the parent session
Rate limit signal present in the runtime, non-actionable at the CLI surface
The layer that needs the signal doesn’t have it. Or has it in a form it can’t use.
Every one of those is a signal integrity failure. The failure pattern is the same. Only the boundary changes.
Pavel’s PR didn’t add more retries. It added a structured field that carries cooldown semantics across a framework boundary. Then the retry logic could do its job.
That distinction matters. The fix itself validates the thesis:
Retryability classification and retry timing semantics must propagate together.
That’s not a tip. It’s a systems invariant. And it applies far beyond Genkit.
The canonical receipt for this finding is now in pitstop-truth:
PT-2026-05-20-github-genkit-ai-genkit-5270-hidden
Hidden — not ignored. The signal wasn’t present at the decision layer until the fix added the propagation channel. That distinction produces different fixes.
The deeper lesson from Genkit is that:
modern AI systems increasingly fail not because signals are absent, but because execution-critical semantics degrade as they cross layers.
The systems that survive will be the ones that preserve signal integrity end-to-end.
Brent Williams builds Pitstop — execution reliability infrastructure for AI/API pipelines.
64 receipts. 35+ production systems.
github.com/SirBrenton/pitstop-truth
```json
{
"post": "I Filed a GitHub Issue Against Google's AI Framework. Here's What Happened.",
"published": "2026-05-21",
"thesis": "AI execution systems fail when decision-relevant signals lose integrity across layers — not because retry logic is missing",
"primary_finding": {
"repo": "genkit-ai/genkit",
"issue": 5270,
"pr": 5343,
"pr_lines": 358,
"providers_fixed": ["Anthropic", "OpenAI", "Google GenAI", "Vertex AI"],
"signal_failure_type": "hidden",
"fix_mechanism": "responseMetadata.retryAfterMs propagation channel added to GenkitError"
},
"receipt_id": "PT-2026-05-20-github-genkit-ai-genkit-5270-hidden",
"key_invariant": "Retryability classification and retry timing semantics must propagate together",
"canonical_publication": {
"title": "When Retry-After Gets Lost: A Case Study in AI Execution Semantic Degradation",
"url": "https://github.com/pitstop-hq/pitstop/blob/main/publications/genkit-retry-after-propagation-report.md",
"type": "propagation_report",
"relationship_to_post": "formalized_case_analysis",
"purpose": "Canonical public reference for the Genkit Retry-After propagation finding, evidence boundary, operational invariant, and broader semantic degradation pattern."
},
"corpus": "https://github.com/pitstop-hq/pitstop/tree/main/corpus",
"contact": " [email protected] "
}
```
Previous issue
Why My Corpus Is Smarter Than I Am
Share
Subscribe to Pitstop
[email protected]
Subscribe
Pitstop © 2026
Sign up
