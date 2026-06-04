---
source: "https://rearmhq.com/news/2026-06-01-rearm-26-06-5-release/"
hn_url: "https://news.ycombinator.com/item?id=48397100"
title: "ReARM 26.06.5: Agentic Coding Guardrails and DevOps"
article_title: "ReARM 26.06.5: Agentic Coding Guardrails and DevOps - ReARM by Reliza"
author: "taleodor"
captured_at: "2026-06-04T13:14:54Z"
capture_tool: "hn-digest"
hn_id: 48397100
score: 1
comments: 0
posted_at: "2026-06-04T11:24:32Z"
tags:
  - hacker-news
  - translated
---

# ReARM 26.06.5: Agentic Coding Guardrails and DevOps

- HN: [48397100](https://news.ycombinator.com/item?id=48397100)
- Source: [rearmhq.com](https://rearmhq.com/news/2026-06-01-rearm-26-06-5-release/)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T11:24:32Z

## Translation

タイトル: ReARM 26.06.5: エージェント コーディング ガードレールと DevOps
記事のタイトル: ReARM 26.06.5: エージェント コーディング ガードレールと DevOps - ReARM by Reliza
説明: ReARM v26.06.5https://github.com/relizaio/rearm/releases/tag/26.06.5 のメジャー リリースを発表します。詳細については、リリース バージョンをご覧ください。

記事本文:
ReARM 26.06.5: エージェント コーディング ガードレールと DevOps - ReARM by Reliza ReARM ブログ
ReARM 26.06.5: エージェント コーディング ガードレールと DevOps
ReARM v26.06.5 のメジャー リリースを発表します。詳細情報は、ReARM Demo インスタンスのリリース ビューでご覧いただけます。 ReARM Pro インストールはすでにアップグレードされています。 ReARM CE ユーザーは、以下で説明する変更点を活用するためにアップグレードすることをお勧めします。
このリリースの見出しは、ReARM バックエンドと UI にまたがる AI コーディング エージェントを管理するためのエンドツーエンドのプラットフォームです。 ReARM は、専用の AI エージェント ダッシュボード、インタラクティブなセッション ビュー、およびライブ セッション テーブルを備えた、エージェント、セッション、およびモデル オントロジーをファーストクラスのエンティティとしてモデル化するようになりました。セッションで生成されたすべてのコミットは、コミットトレーラー属性を通じて、それを作成したエージェントおよびセッションに結び付けられます。 ReARM では、人間のコミッター キーに加えて、エージェントの署名キー登録 (サーバー側の指紋と SSH ID 導出を含む FREEFORM キーを含む) による完全なコミット署名検証 (SSH と GPG の両方) も追加されています。そのため、ReARM は、どのコミットが人間の開発者からのもので、どれが AI エージェントからのものであるかを明確に区別できるようになりました。エージェントは、 /api/agents/orientation.md で提供される公開されたオリエンテーション コントラクトに対してブートストラップし、新しい AGENT 権限関数がエージェントの操作を制御します。 ReARM Pro では、構成可能なポリシーにより、セッション、プルリクエスト、リリース レベルでエージェント操作をブロックできます。たとえば、操作の続行を許可する前に、最終セッション レポート、署名付きコミット、または重大な脆弱性がないなどの最小限のセキュリティ体制が必要です。
DevOps サーフェスはプレビューからベータ版に移行します。 ReARM DevOps を使用すると、製品リリースをインスタンスに割り当てることができ、ReARM CD が各リリースを適切なインスタンスに自動的にルーティングします。

承認ステータスと各インスタンスの環境に基づいているため、リリースは、必要なゲートを通過するにつれて、ステージングから本番に向けて進みます。同様に重要なことは、ReARM は、デプロイが意図されたものだけでなく、各インスタンスで実際に実行されているものをレポートします。インスタンス ビューはインスタンス、計画履歴、実際の履歴のタブに分割され、ウォッチャーによって報告された (「実際の」) デプロイされたイメージが計画された機能セットに対して表示されるため、ドリフトが一目でわかります。インスタンスにアタッチされた機能セットには Helm 依存関係が 1 つだけ必要で、ターゲット リリースのスコープは名前空間に設定でき、スコープごとの DEVOPS_READ / DEVOPS_WRITE 権限はエンドツーエンドで組み込まれます。詳細については、新しい DevOps ワークフローのドキュメントを参照してください。 DevOps 機能は ReARM Pro のみです。
エージェント ガードレールと DevOps サーフェスを組み合わせると、AI エージェントはループ全体を閉じることができます。コードを書いてプル リクエストを開くだけでなく、構築したものをデプロイして結果を観察することもできます。エージェントは特定のインスタンスに割り当てることができ、その作業がリリースされると、ロールアウトをエンドツーエンドで推進します。エージェントは、生成したビルドを固定する新しい機能セットをアセンブルし、ターゲット インスタンスをそれに切り替えます。その後、インスタンスの実際の状態をリードバックして、デプロイが完了したことを確認し、計画からのずれを検出します。すべてのステップが上記と同じ属性、署名、およびポリシーのガードレールの下で実行されるため、組織はコードから展開までの自律的なループを実現し、全体を通じて監視可能で管理されます。
このリリースには、バックエンドのパフォーマンスと信頼性に関する大規模な作業が含まれています。生成された列と読み取り専用 Lite エンティティによる合計のみのメトリクス読み取り、Kubernetes のクリーンな再起動のための OOM での終了による調整された JVM/GC 設定、依存関係トラックのクリーンアップとページングの改善、および exp

クエリのタイムアウトが短い、明示的な接続プール。 Rebom BOM エンリッチメント スケジューラは、古いエンリッチメントを再試行し、その作業の範囲をエンリッチメントが構成された組織に限定し、調整主導の BOM 差分通知がリリースごとに 1 回発行されるようになりました。最終的な効果は、メモリ負荷が軽減され、大規模なコンポーネント ポートフォリオでの動作がより予測可能になることです。
バックエンドは Java 25、Spring Boot 4、Jackson 3、ZGC ガベージ コレクターに移行し、バックエンドと Helm チャート全体で準備/起動プローブと正常なシャットダウンが実現します。
このリリースには、依存関係の基礎となる CVE を修正するものなど、多数の依存関係の更新が含まれています。 ReARM ユーザーは、これらの修正の恩恵を受けるために、このリリースにアップグレードすることをお勧めします。
すべての ReARM リリースの TEI を引き続き公開します。このリリースの TEI: urn:rei:purl:demo.rearmhq.com:pkg:github/relizaio/rearm@26.06.5 。

## Original Extract

We're announcing a major release of ReARM v26.06.5https://github.com/relizaio/rearm/releases/tag/26.06.5. Detailed information is available on its release v...

ReARM 26.06.5: Agentic Coding Guardrails and DevOps - ReARM by Reliza ReARM Blog
ReARM 26.06.5: Agentic Coding Guardrails and DevOps
We're announcing a major release of ReARM v26.06.5 . Detailed information is available on its release view on the ReARM Demo instance. ReARM Pro installations have already been upgraded; ReARM CE users are encouraged to upgrade to benefit from the changes described below.
The headline of this release is an end-to-end platform for governing AI coding agents, spanning the ReARM backend and UI. ReARM now models Agents, Sessions, and a Model Ontology as first-class entities, with a dedicated AI Agents dashboard, an interactive session view, and a Live Sessions table. Every commit produced in a session is tied back to the agent and session that authored it through commit-trailer attribution. ReARM also adds full commit signature verification — both SSH and GPG — with signing-key enrolment for agents (including FREEFORM keys, with server-side fingerprint and SSH-identity derivation) alongside human committer keys, so ReARM can now clearly tell which commits come from human developers and which from AI agents. Agents bootstrap against a published orientation contract served at /api/agents/orientation.md , and a new AGENT permission function governs agentic operations. On ReARM Pro, configurable policies can block agentic operations at the session, pull-request, and release level — requiring, for example, a final session report, signed commits, or a minimum security posture such as no critical vulnerabilities before an operation is allowed to proceed.
The DevOps surface graduates from Preview to Beta. With ReARM DevOps you assign product releases to instances, and ReARM CD automatically routes each release to the right instances based on its approval status and each instance's environment — so a release progresses from staging toward production as it clears the required gates. Just as importantly, ReARM reports back what is actually running on each instance, not only what was intended to deploy: the Instance view is split into Instance, Plan History, and Actual History tabs, and watcher-reported ("actual") deployed images are surfaced against the planned feature set so drift is visible at a glance. Feature sets attached to instances require exactly one Helm dependency, target releases can be scoped to a namespace, and per-scope DEVOPS_READ / DEVOPS_WRITE permissions are plumbed end-to-end. See the new DevOps workflow documentation for details. DevOps functionality is ReARM Pro only.
Taken together, agentic guardrails and the DevOps surface let an AI agent close the entire loop — not just write code and open a pull request, but deploy what it built and observe the result. An agent can be assigned to a specific instance and, once its work is released, drive the rollout end to end: it assembles a new feature set pinning the build it produced and switches the target instance over to it, then reads back the instance's actual state to confirm the deployment landed and to detect drift from what was planned. Because every step runs under the same attribution, signing, and policy guardrails described above, organizations get an autonomous code-to-deployment loop that stays observable and governed throughout.
This release includes a large round of backend performance and reliability work: totals-only metrics reads backed by generated columns and read-only Lite entities, tuned JVM/GC settings with exit-on-OOM for clean Kubernetes restarts, Dependency-Track cleanup and paging improvements, and an explicit connection pool with shorter query timeouts. The Rebom BOM-enrichment scheduler now retries stale enrichments and scopes its work to enrichment-configured organizations, and reconcile-driven BOM-diff notifications are emitted once per release. The net effect is lower memory pressure and more predictable behavior on large component portfolios.
The backend moves to Java 25, Spring Boot 4, Jackson 3, and the ZGC garbage collector, and gains readiness/startup probes and graceful shutdown across the backend and Helm chart.
This release contains a number of dependency updates, including those fixing underlying CVEs in dependencies. ReARM users are encouraged to upgrade to this release to benefit from these fixes.
We are continuing to publish TEIs for all ReARM releases. TEI for this release: urn:tei:purl:demo.rearmhq.com:pkg:github/relizaio/rearm@26.06.5 .
