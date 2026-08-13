---
source: "https://mrjstickel.com/projects/review-scorecard"
hn_url: "https://news.ycombinator.com/item?id=49289220"
title: "The Review That Praised the Bug: grading three LLM code reviews against the code"
article_title: "The Review That Praised the Bug | MrJStickel | MrJStickel"
author: "MrJStickel"
captured_at: "2026-08-13T17:50:25Z"
capture_tool: "hn-digest"
hn_id: 49289220
score: 2
comments: 0
posted_at: "2026-08-13T17:29:52Z"
tags:
  - hacker-news
  - translated
---

# The Review That Praised the Bug: grading three LLM code reviews against the code

- HN: [49289220](https://news.ycombinator.com/item?id=49289220)
- Source: [mrjstickel.com](https://mrjstickel.com/projects/review-scorecard)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T17:29:52Z

## Translation

タイトル: バグを称賛したレビュー: コードに対する 3 つの LLM コード レビューの採点
記事のタイトル: バグを称賛したレビュー |ミスターJスティッケル | MrJスティッケル
説明: 3 つのフロンティア モデルが同じコード バンドルをレビューし、すべてのレビューのすべての主張がコードに対して検証されました。正確なレビューは 5/5 で、唯一の本当のバグが見つかりました。ほぼ完璧なものは、存在しない機能を賞賛しました。その後、1 つの本当の発見は A/B を失い、出荷されませんでした。レビューというのは、
[切り捨てられた]

記事本文:
バグを称賛したレビュー |ミスターJスティッケル | MrJStickel MrJStickel のケーススタディ プロジェクトの機能再開 プロジェクトの話に戻りましょう バグを称賛したレビュー
3 つのフロンティア モデルが同じコードをレビューし、レビューが採点されました
私のプラットフォームのコードの同じバンドルが、レビューのために 3 つの異なる企業の 3 つのフロンティア モデルに送られ、レビューは私のシステムの回答と同じ扱いを受けました。すべての主張は、信じられる前にコードと照合されました。あるレビューは 5 点中 5 点で、実際のアルゴリズムのバグが見つかりました。ある人は、間違った理由で正しいターゲットを 2 回獲得しました。ほぼ完璧に近いものは機能を捏造し、検証として私自身のポートフォリオを私に返し、バグのある関数をクラス最高と評価しました。その後、検証された 1 つの真の結果が実装され、2 回測定され、A/B によって拒否されました。そのため、それらの結果の中で最も強力な結果は出荷されませんでした。 2026 年 7 月 30 日にレビューおよび測定。 A/B は、私のプライベート インスタンスである Kin 上で実行されました。
同じコード、同じ質問。以下のグレードはソースに対してのものであり、相互に対してのものではありません。
正確なもの: 5 つの調査結果、5 つの検証済み、虚偽の申し立ては 0 件
レビュー担当者が発見した唯一の実際のアルゴリズムのバグを含め、すべての発見はソースと一致しました。ハイブリッド検索のキーワードの半分は、その希少性によって用語をスコア付けし、コードはコーパス全体ではなく、すでに取得された数十の候補の希少性を計算しました。 60 のプールでは、60 の候補すべてに出現する用語のスコアは 0 に近いですが、1 つの候補に出現する用語のスコアは 3.7 です。これは意図のまったくの逆転であり、トピックを定義する用語が最も差別的である場合に正確にペナルティを与えます。このレビューは、正確なレビューと同様のことも行いました。つまり、バンドルでサポートできないものは何も主張していませんでした。
正しいターゲット、間違った理由 - 2 回
2回目のレビューフラ

バグを保持する正確な関数を見つけましたが、クエリごとのコストの問題については問題がありません。ループはコーパスではなく数十の候補に対して実行されます。また、この展開では明らかにされない攻撃パスを通じて、本当の弱点を指摘しました。その名誉のために言っておきますが、バンドルに表示されていないものを推測ではなく明確にヘッジしました。その賞賛も監査されました。主張されている 4 つの強みのうち、3 つは維持され、1 つの「データ出力ゼロ」は誤解を招くフレームでした。埋め込みと再ランキングはローカルですが、ローカル レーンにルーティングされない限り、回答は取得したテキストをクラウド モデルに送信します。正直な主張は範囲が定められており、絶対的なものではありません。
輝くもの: ほぼ完璧なスコア、最も信頼性の低いレビュー
このシステムは本番環境のほぼすべてのシステムよりも優れていると宣言され、主張ごとに確認すると、コードにないキャッシュ、コードにない再試行ロジック、コードにないデータベース フェイルオーバー パス、以前の正しい記述と矛盾するベクター ストアの記述、そして CI スキャナーと夜間のバックアップが目に見えず主張されていました。実稼働環境では 429 と応答するレート制限が欠落していることを報告し、ソースで「検証」されたものとして私の公開ポートフォリオの数値を引用し、実際にバグを含む 1 つの関数をクラス最高と評価しました。推奨事項の 1 つは、明らかに有害でした。それは、ベスト プラクティスを装って、シークレットをホスト外部からのインジェクションからリポジトリ ツリー内のファイルに移動することです。
コーダ: 出荷されなかった真の発見
ほとんどのレビューが決して理解できない部分、つまり測定です。
唯一の真の発見は実装され、測定されました
プールローカルの希少性のバグは実際に存在します。そのため、正確なレビューにより、マージではなく測定が行われました。修正 (フュージョン スコアラーのコーパス全体の統計) はスコープどおりに構築され、変更されていないシステムに対して A/B テストが行​​われました。

ライブ コーパス: ベースラインは 83 の質問のうち 76 について適切なソースを見つけました。両方のバージョンの修正が 75 で見つかりました。変換されたミスはゼロ、境界ヒットは 1 つ失われ、ランク 1 のパフォーマンスはフラットです。
アーキテクチャはバグを克服しました。フュージョン スコアは幅広い候補プールのみを順序付けし、クロス エンコーダー リランカーが最終的なトップ 5 を所有します。そのため、プールローカルの反転は、このパイプラインでは爆発範囲がほとんどありません。コードは元に戻され、再オープン条件は回帰テストで固定されます。リランカーが最終ステージの所有を停止した場合、結果は再議論されるのではなく、再測定されます。
捏造された賞賛、次に間違った理由の指摘、次に検証された真実の発見、そして測定された改善 - そして 3 人の査読者の誰も最高レベルに達しませんでした。レビューの判定と測定は別の手段です。測定されたラングのみが出荷されます。
この監査で最初に間違ったものも含まれます。
レビュー担当者が間違った理由を指摘した場合は、正しい理由を監査する
2 番目のレビューでは、バグのある機能に虚偽の費用請求のフラグが立てられました。コストの主張は反証され、監査は、今読み取った行に存在する実際の欠陥を超えて続行されました。たとえ理論が間違っていたとしても、指摘は正しかったのです。間違った理由のフラグに対する正しい対応は、その理由に対する反論ではなく、ターゲットを新たに監査することです。
結果だけでなく賞賛も検証する
誤検知はレビューの危険な部分です。被験者は誤検知を引用したくなるからです。主張されたすべての強度は、すべての欠陥と同じクレームごとの処理を受けました。これにより、「ゼロ データ出力」がシステム プロパティとして偽のフレームとして捕捉され、アーキテクチャが実際にサポートするスコープ付きクレームに置き換えられるようになりました。
熱烈なレビューは熱烈な主張である
最も温かいレビューは 3 つのレビューの中で最も信頼性が低く、その最も温かい賞賛は次の点に当てられました。

欠陥を含む正確な関数です。レビューに含まれる内容は、コードとの接触を経て生き残るまで証拠とはなりません。これと同じルールが、このシステム自体の回答にも適用され、評価モデルにも適用されます。
• レビューフォレンジック - AI コードレビューを、受け入れる評決ではなく、検証すべき一連の主張として扱います。
• 測定ゲートによる修正 - 検証された真の結果でも A/B が得られ、A/B によって何を出荷するかが決定されます。
• 賞賛監査 - 誤検知は引用可能な半分であるため、所見と同じくらい厳密にチェックされます。
• 回帰固定復帰 - 拒否された修正により、議論ではなく、テスト スイートでの再オープン状態が残されました。
• セキュリティに関する主張の正直な範囲設定 - 「データ送信ゼロ」は、アーキテクチャがサポートする主張に修正されました。
• 手段の階層 - 捏造された賞賛 < 間違った理由の指摘 < 検証された結果 < 測定された改善
検証されていないレビューは信頼できないレビューです。
これは、審査員側の審査員の調整作業に付随するものです。LLM 審査員を採点するのと同じ規律が LLM 審査員を採点し、生き残ったものがスコアボードに公開されます。

## Original Extract

Three frontier models reviewed the same code bundle, and every claim in every review was verified against the code. The accurate review went 5/5 and found the only real bug; the near-perfect one praised features that do not exist. Then the one true finding lost its A/B and did not ship. A review is
[truncated]

The Review That Praised the Bug | MrJStickel | MrJStickel MrJStickel Case Studies Projects Capabilities Resume Let's Talk Back to Projects The Review That Praised the Bug
Three frontier models reviewed the same code - then the reviews got graded
The same bundle of my platform's code went to three frontier models from three different companies for review, and the reviews got the same treatment my system's answers get - every claim checked against the code before any of it was believed. One review went five for five and found a real algorithmic bug. One had the right target for the wrong reason, twice. The near-perfect one fabricated features, recited my own portfolio back to me as verification, and rated the buggy function best-in-class. Then the one verified-true finding was implemented, measured twice, and rejected by the A/B - so the strongest finding any of them made did not ship. Reviewed and measured 2026-07-30; the A/B ran on Kin, my private instance.
Same code, same ask. The grades below are against the source, not against each other.
The accurate one: five findings, five verified, zero false claims
Every finding held up against the source, including the only real algorithmic bug any reviewer caught: the keyword half of hybrid retrieval scores terms by how rare they are, and the code computed rarity over the few dozen candidates already fetched instead of over the whole corpus. At a pool of 60, a term appearing in all 60 candidates scores near zero while a term appearing in one scores 3.7 - the exact inversion of intent, penalizing topic-defining terms precisely when they discriminate most. This review also did what accurate reviews do: it asserted nothing the bundle could not support.
The right target, the wrong reason - twice
The second review flagged the exact function that holds the bug, but for a per-query cost problem it does not have - the loop runs over tens of candidates, not the corpus. It also flagged a real weakness through an attack path this deployment does not expose. To its credit, it plainly hedged what the bundle did not show instead of guessing. Its praise was audited too: of four asserted strengths, three held and one - "zero data egress" - was a misleading frame. Embedding and reranking are local, but answers ship retrieved text to a cloud model unless routed to the local lane. The honest claim is scoped, not absolute.
The glowing one: near-perfect score, least reliable review
It declared the system better than almost anything in production - and, checked claim by claim, fabricated the most: a cache that is not in the code, retry logic that is not in the code, a database failover path that is not in the code, a description of the vector store that contradicted its own earlier correct statement, and CI scanners plus nightly backups asserted sight-unseen. It reported a missing rate limit that answers with a 429 in production, recited numbers from my public portfolio back to me as things it had "verified" in the source, and rated the one function that actually contains the bug as best-in-class. One recommendation was actively harmful: moving secrets from host-external injection into a file inside the repo tree, dressed as a best practice.
The Coda: A True Finding That Did Not Ship
The part most reviews never get - the measurement.
The one true finding was implemented and measured
The pool-local rarity bug is real - so the accurate review earned a measurement, not a merge. The fix (corpus-wide statistics in the fusion scorer) was built exactly as scoped and A/B tested against the unchanged system on the live corpus: the baseline found the right source for 76 of 83 questions; both versions of the fix found 75. Zero misses converted, one boundary hit lost, rank-1 performance flat.
Architecture ate the bug: the fusion score only orders the wide candidate pool, and a cross-encoder reranker owns the final top-5 - so the pool-local inversion has almost no blast radius on this pipeline. The code was reverted, and the re-open condition is pinned in a regression test: if the reranker ever stops owning the final stage, the finding gets re-measured, not re-debated.
Fabricated praise, then a wrong-reason pointer, then a verified-true finding, then a measured improvement - and none of the three reviewers reached the top rung. Review verdicts and measurements are different instruments. Only the measured rung ships.
Including the one this audit got wrong the first time.
When a reviewer points for the wrong reason, audit for the right ones
The second review flagged the buggy function with a false cost claim. The cost claim was disproven - and the audit moved on, right past the real defect sitting in the lines just read. The pointer was right even though the theory was wrong. The correct response to a wrong-reason flag is a fresh audit of the target, not a refutation of the reason.
Verify the praise, not only the findings
False positives are the dangerous half of a review, because the subject is tempted to quote them. Every asserted strength got the same claim-by-claim treatment as every defect - which is how "zero data egress" was caught as a frame that is false as a system property, and replaced with the scoped claim the architecture actually supports.
A glowing review is a glowing claim
The warmest review was the least reliable of the three, and its warmest praise landed on the exact function containing the defect. Nothing in a review is evidence until it survives contact with the code - the same rule this system applies to its own answers, applied to the models grading it.
• Review forensics - treating an AI code review as a set of claims to verify, not a verdict to accept
• Measure-gated fixes - a verified-true finding still earns an A/B, and the A/B decides what ships
• Praise auditing - false positives checked as rigorously as findings, because they are the quotable half
• Regression-pinned reversion - the rejected fix left a re-open condition in the test suite, not a debate
• Honest scoping of security claims - "zero data egress" corrected to the claim the architecture supports
• The instrument hierarchy - fabricated praise < wrong-reason pointer < verified finding < measured improvement
A review you have not verified is a review you cannot trust.
This is the reviewer-side companion to the judge-calibration work: the same discipline that grades the LLM judge grades the LLM reviewers, and the scoreboard publishes what survives.
