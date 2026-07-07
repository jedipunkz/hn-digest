---
source: "https://conformityengineering.com/playbook/"
hn_url: "https://news.ycombinator.com/item?id=48824992"
title: "EU AI Act becomes applicable Aug 2: an engineering checklist"
article_title: "The Conformity Engineering Playbook: Preparing Your AI System for August 2, 2026"
author: "stevalsoto"
captured_at: "2026-07-07T22:51:52Z"
capture_tool: "hn-digest"
hn_id: 48824992
score: 1
comments: 0
posted_at: "2026-07-07T22:39:17Z"
tags:
  - hacker-news
  - translated
---

# EU AI Act becomes applicable Aug 2: an engineering checklist

- HN: [48824992](https://news.ycombinator.com/item?id=48824992)
- Source: [conformityengineering.com](https://conformityengineering.com/playbook/)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T22:39:17Z

## Translation

タイトル: EU AI 法が 8 月 2 日に適用される: エンジニアリング チェックリスト
記事のタイトル: 適合性エンジニアリング ハンドブック: 2026 年 8 月 2 日に向けて AI システムを準備する
説明: 2026 年 8 月 2 日の期限までに EU AI 法への準拠を AI システムに組み込むための実用的な記事ごとのチェックリスト (リスク管理、データ ガバナンス、ロギング、人間による監視、証拠生成)。

記事本文:
← 適合性エンジニアリング
適合性エンジニアリング ハンドブック: 2026 年 8 月 2 日に向けて AI システムを準備する
2026 年 7 月 7 日に公開 · 12 分で読みました
2026 年 8 月 2 日、EU AI 法 (規制 (EU) 2024/1689) が一般施行されます。
該当します。 AI システムが雇用、信用スコアリング、教育などの高リスクのカテゴリーに分類される場合、
重要なインフラストラクチャ、必須サービスなど - 一連の具体的で監査可能な義務
「今後の規制」ではなくなり、システムが実行される法律になります。
ほとんどのチームは、2018 年に GDPR に取り組んだ方法、つまりコンプライアンス プロジェクトと同じようにこれに取り組んでいます。
弁護士が運営し、文書を作成します。このアプローチは AI システムでは失敗します。理由は 1 つあります。
システムは事務処理よりも早く変化します。前回の適合性評価
四半期のモデルは、今朝出荷したものについて規制当局にほとんど伝えません。
これは、適合性エンジニアリングが解決するために存在する問題です。つまり、規制を扱う必要があります。
システムの特性としての適合性 — 設計され、パイプラインによって強制され、
事後的にそれについて書かれた報告書ではなく、継続的に証拠が残されています。信頼性が高まりました
この移行を通じて SRE が生成されました。セキュリティはそれを通過し、DevSecOps を生み出しました。
次はコンプライアンスです。
以下は実践版です。高リスク システムに対する EU AI 法の中核的な義務、
それぞれを満たすエンジニアリング実践と、それが必要とする証拠成果物にマッピングされます。
プロデュースする。すべての項目は、記事→実践→証拠という同じパターンに従います。
違反に対する罰金は最大 3,500 万ユーロ、または禁止されたものに対する世界の年間売上高の 7% となります。
慣行、その他ほとんどの義務の不遵守に対しては最大 1,500 万ユーロまたは 3% が課せられます。これらは
GDPR クラスの数値 — まさにこれが、取締役会が法的な問題だけでなく、エンジニアリングについても質問し始めている理由です。
どういう計画ですか

は。
チェックリスト: 記事 → 実践 → 証拠
1. システムを正直に分類する
ダウンストリームのすべては、システムが高リスクかどうかによって決まります。練習：作る
バージョン化されたエンジニアリング上の決定を分類し、リポジトリに文書化し、再評価します。
重要な機能変更ごと - 第 3 四半期に追加したユースケースにより、リスクが最小限のシステムが反転する可能性があります
ハイリスクなものに。証拠: 根拠を記載した日付付きの分類メモ。
バージョン管理。
2. 生活リスク管理システムを運用する
同法では、リスク管理をライフサイクル全体にわたって「継続的」かつ「反復的」にすることが求められています。
それはプロセス要件であり、文書要件ではありません。練習: リスクを負わないようにする
コードとして登録し、リリースごとにレビューします。緩和されていないリリースをブロックする
既知のリスク。証拠: リリースタグに関連付けられたレジスターの変更履歴。
3. データを監査されるように管理する
トレーニング、検証、テストのデータは品質基準を満たしている必要があり、調査する必要があります。
可能性のあるバイアス。実践: すべてのデータセットのデータセット カード、からの系統追跡
ソースからモデルまで、および 1 回限りの研究ではなく、パイプライン ステップとして実行されるバイアス検査。
証拠: データセットのドキュメントとトレーニングの実行ごとに生成されるバイアス テスト レポート。
4. 技術文書を作成します - 書かないでください
Annex IV の文書は市場投入前に存在し、最新の状態に保たれている必要があります。手書き
定義上、ドキュメントは古いものです。実践: docs-as-code — Annex IV パッケージを生成する
ビルド時にシステム自体 (アーキテクチャ、モデルのバージョン、評価結果) から取得されます。
証拠: リリースごとに再現可能なドキュメントが構築されている。
5. まだ要求されていない監査のログを作成します
高リスクのシステムは、その存続期間全体にわたってイベントを自動的に記録する必要があります。練習:
初日から監査可能性を中心に構造化された追加専用ロギングを設計: 入力

、
意思決定、モデルのバージョン、上書き - 保持ポリシーを使用します。証拠: ログ
それ自体は不変であり、クエリ可能です。
6. 仮定ではなく指示を出荷する
導入担当者は、システムを正しく解釈して使用できなければなりません。練習:
バージョン管理されたシステム カードと使用説明書は、リリースごとに同梱されます。
変更ログ。証拠: 説明書アーカイブ (バージョンごとに 1 つ)。
7. 人間による監視を政策ではなく機能にする
効果的な監視とは、人間が理解し、介入し、無効にすることができることを意味します。練習:
承認/上書き/停止のパスを製品表面に構築します。これは、製品表面にのみ存在する監視です。
ポリシー PDF は「有効」ではありません。証拠: 人間であることを証明する監視のやりとりの記録
実際にコントロールを使ってみましょう。
8. リリースごとに精度、堅牢性、セキュリティを証明する
これらは測定可能な特性です。実践: 回帰しきい値を使用した評価スイート
CI、スケジュールに基づいた敵対的およびレッドチームのテスト、および標準的なサイバーセキュリティ衛生
モデル固有の攻撃 (ポイズニング、プロンプト インジェクション) まで拡張されました。証拠: 評価
すべてのリリースに添付されるレポート。
9. SDLC を品質管理システムにしましょう
プロバイダーには QMS が必要です。ほとんどのエンジニアリング組織は、コード レビュー、CI/CD などの 80% をすでに実行しています。
インシデント対応 — 名前は付けずに。実践: 既存のライフサイクルを体系化し、
ギャップを埋めて、パイプラインにそれを強制させます。証拠: プロセス文書と
それを強制する CI 構成。
10. 適合性評価をリリースゲートとして扱う
市場投入前: 適合性評価、EU 適合宣言、CE マーキング、
そしてEUデータベースへの登録。練習: それを最終ゲートとしてモデル化します。
前の 9 つの項目がすでに生成した証拠に基づいたリリース プロセス — それが
技術的適合性を文書化する代わりに、その成果を得る。証拠：

署名された
申告および登録記録。
11. 起動後は思いどおりに監視する
市販後のモニタリングは必須であり、重大なインシデントはできるだけ早く報告する必要があります。
15 日以内。最悪のカテゴリではより速くなります。実践: ワイヤー生産モニタリング
明示的なランブック手順として規制報告を伴うインシデント プロセスに移行します。
証拠: ダッシュボードとインシデント記録のモニタリング。
(項目 1) を分類します。どちらが分かると不安の半分は消える
義務は実際にあなたに適用されます。
チェックリストと照らし合わせてギャップを評価します。各項目について: 実践は存在しますか?
そして証拠は自動的に生成されるのでしょうか？ほとんどのチームはログ記録と評価が近いことに気づき、
文書化とリスク管理はそうではありません。
最悪の 2 つのギャップを選択し、それらを設計します。パイプライン チェックと生成されたものです。
文書ではなく成果物。生成された Annex IV パッケージは 10 個の手書きのパッケージに相当します。
この記事は技術ガイドであり、法的なアドバイスではありません。義務はさまざまです
役割 (プロバイダー対デプロイヤー)、セクター、およびメンバー国の実装ごとに - 検証します。
特定の状況については資格のある弁護士と相談してください。
© 2026 Conformity Engineering · Cookie は使用しません。トラッカーはありません。構造により適合します。

## Original Extract

A practical, article-by-article checklist for engineering EU AI Act conformity into your AI system before the August 2, 2026 deadline — risk management, data governance, logging, human oversight, and evidence generation.

← Conformity Engineering
The Conformity Engineering Playbook: Preparing Your AI System for August 2, 2026
Published July 7, 2026 · 12 min read
On August 2, 2026 , the EU AI Act (Regulation (EU) 2024/1689) becomes generally
applicable. If your AI system falls into a high-risk category — hiring, credit scoring, education,
critical infrastructure, essential services, and more — a set of concrete, auditable obligations
stops being "upcoming regulation" and becomes the law your system runs under.
Most teams are approaching this the way they approached GDPR in 2018: a compliance project,
run by lawyers, producing documents. That approach fails for AI systems for one simple reason:
your system changes faster than your paperwork. A conformity assessment of last
quarter's model tells a regulator little about what you shipped this morning.
This is the problem conformity engineering exists to solve: treat regulatory
conformity as a property of the system — designed in, enforced by pipelines, and
continuously evidenced — rather than a report written about it after the fact. Reliability went
through this transition and produced SRE. Security went through it and produced DevSecOps.
Compliance is next.
What follows is the practical version: the core EU AI Act obligations for high-risk systems,
mapped to the engineering practice that satisfies each one and the evidence artifact it should
produce. Every item follows the same pattern: Article → Practice → Evidence.
Penalties scale to the violation: up to €35M or 7% of global annual turnover for prohibited
practices, and up to €15M or 3% for non-compliance with most other obligations. These are
GDPR-class numbers — which is exactly why boards are starting to ask engineering, not just legal,
what the plan is.
The checklist: Article → Practice → Evidence
1. Classify your system honestly
Everything downstream depends on whether your system is high-risk. Practice: make
classification a versioned engineering decision, documented in the repo and re-evaluated on
every significant feature change — a use case you add in Q3 can flip a minimal-risk system
into a high-risk one. Evidence: a dated classification memo with the reasoning, in
version control.
2. Run a living risk management system
The Act requires risk management to be "continuous" and "iterative" across the lifecycle —
that is a process requirement, not a document requirement. Practice: keep a risk
register as code, reviewed at every release; block releases that introduce unmitigated
known risks. Evidence: the register's change history, tied to release tags.
3. Govern your data like it will be audited
Training, validation, and test data must meet quality criteria, and you must examine
possible biases. Practice: dataset cards for every dataset, lineage tracking from
source to model, and bias examinations that run as pipeline steps — not one-time studies.
Evidence: dataset documentation and bias test reports generated per training run.
4. Generate technical documentation — don't write it
Annex IV documentation must exist before market placement and stay current. Hand-written
docs are stale by definition. Practice: docs-as-code — generate the Annex IV package
from the system itself (architecture, model versions, eval results) at build time.
Evidence: a reproducible documentation build per release.
5. Log for the audit you haven't been asked for yet
High-risk systems must automatically record events over their lifetime. Practice:
design structured, append-only logging around auditability from day one: inputs,
decisions, model version, overrides — with a retention policy. Evidence: the logs
themselves, immutable and queryable.
6. Ship instructions, not assumptions
Deployers must be able to interpret and use the system correctly. Practice:
versioned system cards and instructions for use, shipped with every release like a
changelog. Evidence: the instructions archive, one per version.
7. Make human oversight a feature, not a policy
Effective oversight means a human can understand, intervene, and override. Practice:
build approve/override/halt paths into the product surface — oversight that exists only in a
policy PDF is not "effective." Evidence: oversight interaction logs proving humans
actually use the controls.
8. Prove accuracy, robustness, and security per release
These are measurable properties. Practice: eval suites with regression thresholds
in CI, adversarial and red-team testing on a schedule, and standard cybersecurity hygiene
extended to model-specific attacks (poisoning, prompt injection). Evidence: eval
reports attached to every release.
9. Let your SDLC be your quality management system
Providers need a QMS. Most engineering orgs already run 80% of one — code review, CI/CD,
incident response — without naming it. Practice: codify your existing lifecycle,
close the gaps, and let the pipeline enforce it. Evidence: process documentation plus
the CI configuration that enforces it.
10. Treat conformity assessment as a release gate
Before market placement: conformity assessment, EU declaration of conformity, CE marking,
and registration in the EU database. Practice: model it as the final gate of your
release process, fed by the evidence the previous nine items already generate — that's the
payoff of engineering conformity instead of documenting it. Evidence: the signed
declaration and registration record.
11. Monitor after launch like you mean it
Post-market monitoring is mandatory, and serious incidents must be reported — as fast as
within 15 days, faster for the worst categories. Practice: wire production monitoring
to an incident process with regulatory reporting as an explicit runbook step.
Evidence: monitoring dashboards and incident records.
Classify (item 1). Half the anxiety disappears when you know which
obligations actually apply to you.
Gap-assess against the checklist. For each item: does the practice exist,
and does it produce evidence automatically? Most teams find logging and evals are close,
documentation and risk management are not.
Pick the two worst gaps and engineer them — pipeline checks and generated
artifacts, not documents. A generated Annex IV package is worth ten hand-written ones.
This article is an engineering guide, not legal advice. Obligations vary
by role (provider vs. deployer), sector, and member-state implementation — validate your
specific situation with qualified counsel.
© 2026 Conformity Engineering · No cookies. No trackers. Conform by construction.
