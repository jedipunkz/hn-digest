---
source: "https://blog.cloudflare.com/engineering-standards-enforcement/"
hn_url: "https://news.ycombinator.com/item?id=49170628"
title: "Cloudflare enforces engineering standards using AI"
article_title: "How Cloudflare enforces engineering standards using AI | The Cloudflare Blog"
author: "garyhtou"
captured_at: "2026-08-04T16:07:48Z"
capture_tool: "hn-digest"
hn_id: 49170628
score: 1
comments: 0
posted_at: "2026-08-04T15:49:50Z"
tags:
  - hacker-news
  - translated
---

# Cloudflare enforces engineering standards using AI

- HN: [49170628](https://news.ycombinator.com/item?id=49170628)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/engineering-standards-enforcement/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:49:50Z

## Translation

タイトル: CloudflareはAIを使用してエンジニアリング標準を強制します
記事のタイトル: Cloudflare が AI を使用してエンジニアリング標準を強制する方法 | Cloudflareのブログ
説明: AI エージェントが開発ライフサイクル全体で使用するエンジニアリング標準の管理団体である Cloudflare Codex を作成しました。構造化された RFC とエージェント レビューを組み合わせることで、チームはコード、仕様、インシデント レポート間の一貫性を自動的に強化します。

記事本文:
Cloudflare が AI を使用してエンジニアリング標準を強化する方法 | Cloudflareのブログ
コンテンツへスキップ すべてのカテゴリ AI
ログイン 営業担当者へのお問い合わせ ブログ Agents Week AI ベスト プラクティス +2 さらに 2 つのタグを表示 5 つのタグ 5 つのタグを表示 選択したタグ
Agents Week AI ベスト プラクティス 開発者プラットフォーム開発者
自動プラットフォーム最適化
Cloudflare 1 ユーザーのリスクスコア
CloudflareがAIを使用してエンジニアリング標準を強制する方法
過去 4 か月間で、当社の AI コードレビュー担当者は、Cloudflare エンジニアリング標準からの 25 万件近くの逸脱 (この投稿では「違反」と呼びます) を報告し、16,000 件のマージをブロックしました。当社の仕様レビュー担当者は、実装が開始される前に、同じ基準に照らして 600 近くの技術設計を評価しました。どちらのシステムも、人々とエージェント向けに構築されたエンジニアリング ガイダンスの共有ソースである Cloudflare Codex を利用しています。この投稿では、Codex を構築した理由、Codex がエンジニアリング ライフサイクルをどのようにサポートするか、そして次に何を行う予定であるかについて説明します。
Codex (AI エンジニアリング スタックに関する前回の投稿で簡単に紹介しました) が登場する前は、Cloudflare の開発者ガイダンスは、正式なドキュメント、リポジトリ ファイル、チャット スレッド、個々のエンジニアの蓄積された知識など、さまざまな場所に存在していました。エンジニアは、解決しようとしている問題に取り組むのではなく、ガイダンスを探すことに多くの時間を費やしすぎることがよくありました。たとえ答えを見つけたとしても、それが最新のものなのか、信頼できるものなのか、自分たちの状況に当てはまるものなのか、必ずしも判断できるわけではありません。
Cloudflareが成長するにつれて、そのモデルを維持することがますます困難になってきました。すべての規格を読むことができるエンジニアは存在しませんし、レビュー担当者がすべての要件を確実にチェックすることもできません。人々がチーム間を移動したり、ガイダンスが一貫して表面化されなかったり、情報が提供されなかったりすると、組織的な知識を回復することが困難になりました。

rced により、プロジェクト間の漂流が生じました。
私たちはこの一連の知識を Cloudflare Codex として再構築しました。Cloudflare Codex は、エージェントが作業時に取得して適用できる、管理されたエンジニアリング標準のセットです。同じガイダンスにより、コード レビュー、技術設計レビュー、インシデント レポート レビュー、その他多くのユースケースに情報を提供できるようになり、エンジニアは結果として得られる調査結果に時間と判断を集中できます。
コーデックスの構成とワークフロー
専用の Codex ガバナンス モデルは、Codex を、私たちが関心を持っているエンジニアリング領域をカバーする個別のドメインに分割します。これらには、アーキテクチャの問題 (フロントエンドやコントロール プレーンなど)、横断的な問題 (セキュリティと信頼性)、特定の言語 (TypeScript や Rust)、およびその他のいくつかの領域が含まれます。各ドメインは、ドキュメントの内容、一貫性、全体的な品質を監督する責任を負う所有者によって主導されます。
コーデックス標準では、Request for Comments (RFC) 形式が使用されます。要件では、RFC 2119 で定義されている SHOULD キーワードと MUST キーワードを使用します。また、前付ヘッダーにはドメインや RFC ステータスなどのメタデータが保持されることも期待されます。重要な関心とドメインコンピテンシーを持つCloudflare従業員は誰でも、所定の構造に従ったマージリクエストを通じてRFCを提案できます。その後、提案は、ますます広範な審査員グループからの数回のフィードバックを経ます。ドメイン所有者が最終承認を与えると、RFC はコーデックスの一部となり、Astro を利用した内部サイトに公開されます。
承認された RFC は、Codex クライアントおよびエージェントによって使用され、コード、構成、またはドキュメント内の Codex 違反に直ちにフラグを立て始める可能性があります。ただし、Codex ステートメントに基づいてブロックするのは、RFC が承認済みライフサイクル状態から施行済みライフサイクル状態に移行した後のみです。この個別の昇格ステップにより、チームは吸収する時間が与えられます。

新しい要件を検討し、執行に追加の作業が必要な場合に対応します。
次の図は、Codex ワークフローの手順を示しています。
単純なプロセスはここで停止し、Codex 全体をそのまま大規模言語モデル (LLM) にフィードする可能性があります。ただし、すでに RFC の数が増加していることを考えると (60 以上、さらに増え続けています)、コーパスのボリュームはコンテキスト ウィンドウに大きなストレスを与え、LLM の結果に悪影響を及ぼす可能性があります。モデルを最も関連性の高い RFC に誘導できるようにするために、専用のエージェントを呼び出して、SHOULD ステートメントと MUST ステートメントを専用の JSON 構造に自動的に抽出して圧縮し、遅延検出とプログレッシブ開示をサポートするメタデータで強化します。次の抜粋は、コントロール プレーン サービス RFC の結果を示しています。
{
"rfc" : 14 、
"title" : "コントロール プレーン サービス" ,
"ステータス" : "承認済み" 、
"ドメイン" : "コントロールプレーン" ,
「ステートメント」: [
{
"slug" : "エッジ構成の伝播にクイックシルバーを使用" ,
"セクション" : [ "提案" , "インフラストラクチャ" ],
"レベル" : "すべきです" 、
"text" : "システムまたは顧客の構成をエッジに伝播する必要がある場合は、アウトボックス パターン経由で Quicksilver を使用してください" ,
"href" : "/rfcs/014-control-plane-services/#infrastruct"
}、
{
"slug" : "api-schemas-must-be-documented-in-openapi-spec" ,
"セクション" : [ "提案" , "API ゲートウェイ" ],
"レベル" : "必須" 、
"text" : "API リクエストとレスポンスのスキーマは、OpenAPI 仕様を使用して文書化する必要があります" ,
"href" : "/rfcs/014-control-plane-services/#api-gateway"
}
】
各ステートメントは、RFC が更新された場合でも、抽出プロセス中に変更されない安定したスラッグ識別子を受け取ります。識別子を使用すると、異なるシステム間で同じステートメントを長期にわたって追跡できるようになります。これは、監視、分析、例外処理に不可欠です。
最初は追加で

ステートメントを JSON ではなく別のより簡潔な Markdown ファイルにコピーしました。時間の経過とともに、エージェントが必要なコンテンツをより正確にフィルタリングできるように、より豊富な構造化フォーマットに移行しました。ステートメントが適用されるソフトウェア開発ライフサイクル (SDLC) ステージ (設計、実装、ランタイムなど) の指標など、さらに厳密なスコープを設定するための追加のメタデータを含める予定です。
すでにいくつかのシステムが日常のエンジニアリング作業で Codex を使用しています。 AI コードレビュー担当者、仕様レビュー担当者、インシデントレポートレビュー担当者の 3 人のエージェントが、Codex が実際にどのように機能するかを示します。
別のブログ投稿で説明されている AI コード レビューアー エージェントは、コーデックス コンプライアンスを含むいくつかの側面にわたってマージ リクエストを評価します。
レビューごとに、エージェントは RFC を取得し、Codex ステートメントを解析します。モデルまたはコーディネーターが追加のコンテキストを必要とする場合にのみ、完全な RFC 本文を読み込みます。ほとんどの場合、声明は、報告された違反を説明するのに十分な情報を提供します。
SHOULD と MUST の区別は、RFC のステータスと合わせて、レビュー担当者がどのように対応するかを決定します。承認された RFC の調査結果は、妨げとなるものではない推奨事項です。 RFC が施行されると、MUST 要件が満たされていないため、レビュー担当者は重大度に応じて承認を保留するか、マージ リクエストをブロックします。
今年初めにコーデックスが開始されて以来、AI コードレビュー担当者は 230,000 件近くの違反を報告しました。このうち、約 16,000 件で承認が保留されました (つまり、強制された RFC に関する MUST ステートメントに言及したものです)。
コーディネーター フレームワークとサブエージェントの実行により、AI コード レビューアーの 1 回の実行が完了するまでに通常は数分かかります。待ち時間は多くの場合お金 (またはトークン) の価値がありますが、エンジニアは遅延と余分な往復が発生すると非難していました。

n 調査結果を修正する。私たちはエクスペリエンスを改善する方法を検討し、2 つの追加オプションを思いつきました。
機械的に検証できる言語固有の Codex 要件については、カスタム リンター構成パッケージを提供します。これらは Codex 仕様に準拠しており、問題を数ミリ秒で表面化することが可能です。 TypeScript は、Codex リンターのサポートを受けた最初の言語であり、パフォーマンスの高いリンター実行のために oxlint (最近 Cloudflare に加わった VoidZero チームによって保守されています) も標準化しています。 Rustプロジェクト用のリンターは現在開発中ですが、最終的にはGoもそれに続き、Cloudflareで最も一般的に使用される言語を完全にカバーする予定です。
継続的インテグレーション (CI) 部分をレビュー サイクルから切り離すために、コマンド ライン インターフェイス (CLI) を介して AI コード レビューアーをローカルで実行できるようにしました。これは、CI のコーディネーター機能と一致し、自動的に決定された差分セットに対して同じ (OpenCode ベースの) エージェントを実行し、結果がターミナルに表示されます。
私たちはリンターがほぼすべての開発者とコードベースにとって役立つと信じていますが、CLI はそれを好むエンジニアにとってのオプションの代替手段のままです。
Cloudflareのエンジニアは、実装前に定期的に設計文書と技術仕様（略して仕様）を作成します。コーデックスの重要なサブセットは、設計、アーキテクチャ、および技術レビューに関連するその他のテーマに関連しています。実装を開始する前にアーキテクチャ上の間違いを発見するために、仕様を発見し、関連する Codex 要件に照らして評価するエージェントである仕様レビューアーを構築しました。
仕様レビューアーは開発者プラットフォーム上で動作します。Cloudflare ワーカーとして実行され、その結果と状態を D1 に保存し、AI ゲートウェイを介してモデルリクエストをルーティングし、キックします。

Cron トリガーによる新しい仕様のスキャンをオフにします。まず、仕様に関連するドメインとセクションによって Codex をフィルタリングします (たとえば、言語機能や実装に焦点を当てた RFC は無視されます)。いくつかのガイド プロンプトは、評価の実行方法と結果の構成方法をモデルに指示します。結果は重大度 (SHOULD および MUST キーワードの影響を受ける) に基づいて評価され、一般的な品質とアーキテクチャに関するアドバイスが含まれます。レビューの実行が完了すると、レビューの詳細を検査できるカスタム ダッシュボードにリンクするメモが仕様書に残されます。
2026 年 5 月の初め以来、約 600 の独自のオープン仕様がレビューされました。オンデマンドまたは仕様変更によってトリガーされた再実行を含め、現在までに 3,200 件を超えるレビュー呼び出しを追跡しました。所見の大部分は重大度が「重大」(65%) または「軽度」(29%) であり、「重大」な所見は少数 (6%) でした。
次の画像は、仕様レビュー担当者の UI がどのようなものかを示しています。
仕様ドキュメントにコメントを直接投稿し、レビュー評価に影響を与える可能性のある人間とエージェントの会話を埋め込み、人間による追加レビューのために影響力の高い提案にフラグを立てることで、仕様レビュー担当者をより緊密に統合する予定です。
インシデント レポートのレビュー担当者は、同じアプローチをインシデント レポート (事後分析とも呼ばれます) に適用します。各レポートが完全であることをチェックすることに加えて、レポートが何が起こったのかを明確に説明しているかどうかを評価し、要因を特定し、解決策を文書化し、有意義なフォローアップアクションを提案します。これらの期待は、専用の Codex RFC で定義されています。
インシデント レポートのレビュー担当者は、仕様レビュー担当者と同じ開発者プラットフォームの構成要素を使用します。この共有アーキテクチャは、Codex エージェントにとって一般的なパターンになりつつあります。
5月から

2026 年、レビュー担当者は 200 件を超えるインシデント レポートを評価し、フォローアップのアクション項目の欠落、不完全なタイムライン、検出シグナルの欠落などのギャップを特定しました。これらのレポートのうち、93% は、影響が低い、内部のみで発生した、または事前に宣言されたインシデントをカバーしていました。重大度の高いインシデントについては、包括的な中央レビュー プロセスの一環としてレビュー担当者の参加を義務付けており、すべての調査結果に対処するまでレポートは完了したとはみなされません。
Codex は、コード、技術設計、およびインシデント レポートをレビューするエージェントをすでにサポートしています。このモデルを SDLC 全体に拡張し、エージェントが設計、実装、運用全体にわたって一貫して問題を表面化できるようにする予定です。長期的な目標は、エージェントが自律性を高めて問題を特定し、修正を提案できるようにする一方で、エンジニアは引き続きそれらの変更をレビューして承認する責任を負うことです。
また、Codex をエンジニアリングを超えて拡張しています。製品、セキュリティ、コンプライアンス、信頼と安全の各チームは独自の基準を追加し始めており、エージェントが設計と実装だけを超えた考慮事項に照らして作業を評価できるようになります。
多くのエンジニアリング ワークフローにわたって、Codex 支援のエージェントが私たちの作業をサポートしてくれました。

[切り捨てられた]

## Original Extract

We created the Cloudflare Codex, a governed body of engineering standards that AI agents consume across the development lifecycle. By pairing structured RFCs with agentic reviews, teams automatically enforce consistency across code, specs, and incident reports.

How Cloudflare enforces engineering standards using AI | The Cloudflare Blog
Skip to content All Categories AI
Login Contact Sales Blog Agents Week AI Best Practices +2 Show 2 more tags 5 Tags Show 5 tags Selected Tags
Agents Week AI Best Practices Developer Platform Developers
Automatic Platform Optimization
Cloudflare One User Risk Score
How Cloudflare enforces engineering standards using AI
Over the past four months, our AI code reviewer has flagged nearly a quarter of a million deviations from Cloudflare engineering standards (what we’ll call “violations” in this post) and blocked 16,000 merges. Our spec reviewer agent has evaluated close to 600 technical designs against the same standards before implementation began. Both systems draw from the Cloudflare Codex, a shared source of engineering guidance built for people and agents. This post explains why we built the Codex, how it supports the engineering lifecycle, and what we plan to do next.
Before the Codex (which we briefly introduced in a previous post about our AI engineering stack), developer guidance at Cloudflare lived in many places: formal documentation, repository files, chat threads, and the accumulated knowledge of individual engineers. Engineers often spent too much time searching for guidance instead of working on the problem they were trying to solve. Even after finding an answer, they could not always tell whether it was current, authoritative, or applicable to their situation.
As Cloudflare grew, that model became increasingly difficult to sustain. No engineer could read every standard, and reviewers could not reliably check every requirement. Institutional knowledge became harder to recover when people moved between teams, and guidance that was not consistently surfaced or enforced led to drift between projects.
We rebuilt this body of knowledge as the Cloudflare Codex: a governed set of engineering standards that agents can retrieve and apply at the point of work. The same guidance can now inform code review, technical design review, incident report review, and many other use cases, while engineers focus their time and judgment on the resulting findings.
Codex organization and workflow
A dedicated Codex governance model divides the Codex into distinct domains covering the engineering areas we care about. These include architectural matters (for example, frontend and control plane), cross-cutting concerns (security and reliability), specific languages (TypeScript and Rust), and several other areas. Each domain is led by an owner who is responsible for the content, consistency, and overall quality of the documents they oversee.
Codex standards use a Request for Comments (RFC) format. Requirements use the SHOULD and MUST keywords defined by RFC 2119 . We also expect a front matter header to hold metadata such as the domain and RFC status. Any Cloudflare employee with a key interest and domain competency can propose an RFC through a merge request that follows the prescribed structure. The proposal then passes through several rounds of feedback from an increasingly broad group of reviewers. Once the domain owner gives final approval, the RFC becomes part of the Codex and is published to an Astro -powered internal site.
Approved RFCs can be consumed by Codex clients and agents, which may then start to flag Codex violations in code, configuration, or documentation immediately. However, they block based on Codex statements only after an RFC moves from the approved to the enforced lifecycle state. This separate promotion step gives teams time to absorb new requirements and accommodates cases where enforcement needs additional work.
The following diagram illustrates the steps in the Codex workflow:
A naive process could stop here and feed the entire Codex to a large language model (LLM) as is. Given the increasing number of RFCs we have already (60+ and counting), however, the corpus volume would put a lot of stress on the context window and impact LLM results negatively. To help guide models to the most relevant RFCs, we invoke a purpose-built agent to automatically extract and compact the SHOULD and MUST statements into a dedicated JSON structure and enrich it with metadata that supports lazy discovery and progressive disclosure. The following abridged excerpt shows the result for our control plane services RFC:
{
"rfc" : 14 ,
"title" : "Control Plane Services" ,
"status" : "approved" ,
"domain" : "control-plane" ,
"statements" : [
{
"slug" : "use-quicksilver-for-edge-configuration-propagation" ,
"section" : [ "Proposal" , "Infrastructure" ],
"level" : "SHOULD" ,
"text" : "If you need to propagate system or customer configuration to the edge, use Quicksilver via the outbox pattern" ,
"href" : "/rfcs/014-control-plane-services/#infrastructure"
},
{
"slug" : "api-schemas-must-be-documented-in-openapi-spec" ,
"section" : [ "Proposal" , "API Gateway" ],
"level" : "MUST" ,
"text" : "API request and response schemas MUST be documented using an OpenAPI spec" ,
"href" : "/rfcs/014-control-plane-services/#api-gateway"
}
]
} Each statement receives a stable slug identifier that remains unchanged during the extraction process even when its RFC is updated. The identifier lets us track the same statement across different systems over time, which is essential for monitoring, analysis, and exception handling.
Initially, we extracted the statements into another, more concise Markdown file rather than JSON. Over time, we moved to a richer structured format so that agents could filter the content they needed more accurately. We plan to include additional metadata for even tighter scoping, such as indicators for the software development life cycle (SDLC) stage a statement applies to (e.g., design, implementation, runtime).
Several systems already use the Codex in day-to-day engineering work. Three agents show how the Codex works in practice: our AI code reviewer, spec reviewer, and incident report reviewer.
Our AI code reviewer agent, covered in a separate blog post , evaluates merge requests across several dimensions, including Codex compliance.
For each review, the agent retrieves the RFCs and parses the Codex statements. It loads full RFC bodies only when the model or coordinator needs additional context. In most cases, the statements provide enough information to explain a reported violation.
The distinction between SHOULD and MUST, together with an RFC’s status, determines how the reviewer responds. Findings from approved RFCs are non-blocking recommendations. Once an RFC is enforced, an unsatisfied MUST requirement causes the reviewer to withhold approval or block a merge request, depending on the severity.
Since the Codex’s inception earlier this year, the AI code reviewer has flagged close to 230,000 violations. Among these, almost 16,000 caused approval to be withheld (i.e., they referred to MUST statements on enforced RFCs).
A single AI code reviewer run usually takes a couple of minutes to complete due to the coordinator framework and sub-agent execution. Although the wait is very often worth the money (or tokens), engineers were calling out the delay and extra round trip involved in remediating the findings. We looked into how we could improve the experience and came up with two additional options:
For language-specific Codex requirements that can be verified mechanically, we provide custom linter configuration packages. These are aligned with our Codex specification and make it possible to surface problems in milliseconds. TypeScript was the first language to receive Codex linter support while also standardizing on oxlint (maintained by the VoidZero team who joined Cloudflare recently) for performant linter execution. A linter for Rust projects is currently under development, and Go will eventually follow to complete coverage of Cloudflare’s most commonly used languages.
To cut out the continuous integration (CI) leg from the review cycle, we made it possible to run the AI code reviewer locally through a command-line interface (CLI). It matches the coordinator functionality from CI and runs the same ( OpenCode -based) agents against an automatically determined diff set, with results presented in the terminal.
We believe the linters would be useful to almost every developer and codebase, while the CLI remains an optional alternative for engineers who prefer it.
Engineers at Cloudflare regularly write design documents and technical specifications (or specs in short) before implementation. A significant subset of the Codex pertains to design, architecture, and other themes relevant to technical reviews. To catch architectural mistakes before implementation begins, we built the spec reviewer , an agent that discovers specs and evaluates them against relevant Codex requirements.
The spec reviewer operates on the Developer Platform: it runs as a Cloudflare Worker, stores its results and state in D1, routes model requests through AI Gateway, and kicks off scanning for new specs via a Cron Trigger. It starts by filtering the Codex by domains and sections relevant to specs (for example, language features and implementation-focused RFCs are disregarded). Several guiding prompts instruct the model on how to run the assessment and frame the results. The findings get rated based on severity (influenced by SHOULD and MUST keywords) and include general quality and architectural advice. On completion of a review run, a note is left on the spec document linking to a custom dashboard where review details can be inspected.
Since the beginning of May 2026, almost 600 unique open specs have been reviewed. Including reruns triggered on demand or by spec changes, we tracked over 3,200 review invocations to this date. The vast majority of findings had a “major” (65%) or “minor” (29%) severity, with “critical” findings being the minority (6%).
The following image gives an impression of what the spec reviewer UI looks like:
We plan to integrate the spec reviewer more tightly by posting comments directly on the spec documents, embedding human-agent conversations that can influence the review assessment, and flagging high-impact proposals for additional human review.
The incident report reviewer applies the same approach to incident reports (also known as postmortems). In addition to checking that each report is complete, it evaluates whether the report clearly explains what happened, identifies contributing factors, documents the resolution, and proposes meaningful follow-up actions. These expectations are defined in a dedicated Codex RFC.
The incident report reviewer uses the same Developer Platform building blocks as the spec reviewer. This shared architecture is becoming a common pattern for our Codex agents.
Since May 2026, the reviewer has assessed more than 200 incident reports and identified gaps such as missing follow-up action items, incomplete timelines, and omitted detection signals. Among those reports, 93% covered incidents that were low-impact, internal-only, or declared preemptively. For high-severity incidents, we’ve made the reviewer mandatory as part of our comprehensive central review process, and reports are not considered complete until all findings have been addressed.
The Codex already supports agents that review code, technical designs, and incident reports. We plan to extend that model throughout the SDLC, allowing agents to surface issues consistently across design, implementation, and operations. The longer-term goal is for agents to identify issues as well as propose fixes with increasing autonomy, while engineers remain responsible for reviewing and approving those changes.
We are also expanding the Codex beyond engineering. Product, security, compliance, and trust and safety teams are beginning to add their own standards, allowing agents to evaluate work against considerations that extend beyond design and implementation alone.
Across a number of engineering workflows, Codex-backed agents have helped us sur

[truncated]
