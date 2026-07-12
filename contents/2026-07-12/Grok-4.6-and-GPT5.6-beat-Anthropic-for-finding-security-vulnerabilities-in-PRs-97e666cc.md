---
source: "https://docs.damsecure.ai/blog/pr-review-security-benchmark/"
hn_url: "https://news.ycombinator.com/item?id=48885732"
title: "Grok 4.6 and GPT5.6 beat Anthropic for finding security vulnerabilities in PRs"
article_title: "The AI PR Security Review Frontier | Dam Secure Docs"
author: "pcollins123"
captured_at: "2026-07-12T23:43:20Z"
capture_tool: "hn-digest"
hn_id: 48885732
score: 5
comments: 1
posted_at: "2026-07-12T22:57:23Z"
tags:
  - hacker-news
  - translated
---

# Grok 4.6 and GPT5.6 beat Anthropic for finding security vulnerabilities in PRs

- HN: [48885732](https://news.ycombinator.com/item?id=48885732)
- Source: [docs.damsecure.ai](https://docs.damsecure.ai/blog/pr-review-security-benchmark/)
- Score: 5
- Comments: 1
- Posted: 2026-07-12T22:57:23Z

## Translation

タイトル: Grok 4.6 と GPT5.6 が PR のセキュリティ脆弱性の発見で Anthropic を破る
記事のタイトル: AI PR セキュリティ レビューの最前線 |ダムセキュアドキュメント

記事本文:
メイン コンテンツにスキップ ドキュメント ブログ 最近の投稿 2026
AI PR セキュリティ レビュー フロンティア ワシントンが寓話を禁止。とにかく人類が勝利する 3 時間の攻撃ウィンドウ npm がデフォルトで開いたままになる ガードレールとガイドライン DIY セキュリティ レビューの上限 Dam Secure が SOC 2 Type 2 を達成 ゲスト投稿: Dam Secure が見逃しているものをどのように見つけるか AI PR セキュリティ レビューのフロンティアを発表
Simon Harloff Patrick Collins 私たちは、PR セキュリティ レビューにどのモデルを使用するかについて推測に頼らないようにしたいと考えていました。驚いたことに、Anthropic モデルはまったくその領域に達しておらず、わずか 3 日前にリリースされた最新の GPT5.6 が王者です。 Fable はコード全体のスキャンで優れたパフォーマンスを発揮することが知られていますが、PR レビューはユーザーにとって一般的なワークロードです。
このグラフは、プル リクエストあたりのコストと F1 (検索品質) をプロットしています。青い点はフロンティアを定義します。以下の表のメトリクスの定義を参照してください。
同じ 10 個のプル リクエストに対して 10 個のモデルを実行し、それぞれに 1 つのアクセス制御バグ (IDOR、認証の欠落、認証の破損) が植え付けられていました。これをモデルごとに 5 回実行し、すべての検出結果を実際のコードに対してスコア付けしました。
その結果、コストと品質の明確なフロンティアが生まれ、それは Anthropic ではなく OpenAI、xAI、Google に属します。 GPT-5.6 ソルが優勢です。 Fable 5 は PR あたり約 3.61 ドルとパフォーマンスが悪く、品質と価値の両方で低くなります。どの人類モデルもフロンティアに到達することはできません。
私たちのベンチマーク リポジトリは、Benchmaxxing やモデル トレーニングへの偶発的な組み込みを避けるために合成され、プライベートになっています。
Fable はフル コード スキャンでは良好なパフォーマンスを発揮しますが、このワークロードは PR スキャンです。 PR スキャンは、ユーザーにとって非常に一般的なワークロードです。
これは Dam Secure ハーネスの異常ではありません。 Pydantic ハーネスと Claude Code ハーネスもテストしましたが、非常に似た結果が得られました。
ベンチマーク範囲: プルリクエストのみ
このベンチマークは、

追加のコードが追加されたプル リクエストの脆弱性を発見するモデルがどの程度優れているかを測定しようとしているだけです。このベンチマークは、既存のコードベースですべての既知の脆弱性を広範囲に検索するものではありません。フルスキャンベンチマークは近日公開予定です。
この範囲は、結果をどう読むかに影響します。重要な点は、PR のセキュリティ問題を見つけるために Fable (または他のフロンティア クロード モデル) に手を伸ばすと、お金、使用制限、トークンが無駄になるということです。 Fable 5 は、私たちがテストした中で最も高価な構成ですが、依然としてコスト/品質の境界線から外れており、両方の軸で GPT-5.6 Sol が支配しており、PR あたりのコストの約 5 分の 1 で高い再現率に達しています。
既存の大規模なコードベースを自由にローミングするときにこれらのエージェントがどのように動作するかについては、別の記事を別途リリースする予定です。そこではモデルのランキングが異なって見える場合があります。この投稿は PR レビューのみを目的としています。
各実行では、Dam Secure Vulnerability Scanner を高度な推論で使用し、埋め込まれたアクセス制御バグの 10 PR コーパスに対して 5 回繰り返しました。
True Positive (TP) - 発見されることが予想されていた脆弱性が発見されました。
False Positive (FP) - 予期していなかった脆弱性が発見されました (実際のバグではありません)。
False Negative (FN) - 見つかると予想されていた脆弱性が見つからなかった。
Recall (R) - 植え付けられた脆弱性をどの程度うまく発見しましたか?これは正しく特定された実際の陽性者の割合です: TP / (TP + FN)。
精度 (P) - 誤検知パフォーマンス。予測陽性の正しい割合: TP / (TP + FP)。
F1 - 全体的なパフォーマンスのバランスをとる再現率と精度: 2PR / (P + R)。ポイントをホバリングすると、F2 = 5PR / (4P + R) も表示され、この重みは最大 4 倍の精度を再現します。
コスト / PR - すべてのプロンプトに対してそのモデルを実行し、精度に関係なく PR を処理するためのコストです。これが最も正確な表現です

現実世界のコストの n。
コスト / TP - 見つかった真陽性あたりのコストです。
ウィキペディア: 精度、再現率、F1
GPT-5.6 Sol は、当社のハーネスで 100% リコールに達した最初のモデルです (明らかに、発見するのがより困難な脆弱性を追加する必要があります)。
表の Fable 5 の行には、「Fable 5 → Opus 4.8 フォールバック」というラベルが付いています。 Opus 4.8 に引き継がれる確率はわずか 10.7% でした。
GPT5.5 がリリースされる前は、Gemini 3.5 フラッシュがこのタイプのワークロードに推奨されるモデルでした。これは、業界が Opus モデルがセキュリティを支配していると信じたがっているにもかかわらず、当社の評価スイートが証明しているため、これまでは小さな秘密でした。
最良のモデルは GPT-5.6 Sol です。強力な F1 (0.91)、完璧な 100% リコール、そして PR あたり 0.70 ドルと非常にコスト効率が高いです。 5.6 は、ほぼ同じパフォーマンスで、前世代の 5.5 よりも約 45% 安価です。
佳作: Grok 4.5 は、F1 0.77 (再現率 74%、精度 80.4%) で PR あたり 0.20 ドルで最前線に到達し、Gemini 3.1 Flash Lite は、F1 0.75 (再現率 68%、精度 82.9%) という強力な結果で、PR あたり ~0.04 ドルで最安値を固定します。
このベンチマークに他の無差別級モデルを追加します。
オープンソース リポジトリから合成プル リクエストを構築します。各プルリクエストには脆弱性が埋め込まれています。次に、Dam Security Vulnerability Scanner と Pydantic ハーネスの両方を使用して、植え付けられた脆弱性を発見するモデルの能力を測定します。
合成データセットを構築するために 2 つの異なるアプローチを使用します。 OSS リファレンス リポジトリと実際の CVE の逆再生から構築されたプライベート合成リポジトリ (詳細は後述)。
各モデルの実行では、他のオープンソース ハーネスで利用可能な一般的なツールと一致するように構成された Dam Secure 脆弱性スキャナーの限定バージョンが使用されます。これは、セキュリティが強化された小規模なツールセットを備えた Vercel AI SDK 上に構築されています。エージェントはこれを見ます

PR diff を使用すると、次のことができます。
ファイルの読み取り - diff だけでは十分なコンテキストが得られない場合に、ファイルの内容全体を読み込みます
Grep - 変更されたファイルのパターンを検索します (認証チェック、ルート ハンドラー、スコープ検証)。
ハーネスは表内のすべてのモデルで同じです。異なるのは基礎となるモデルのみです。これにより、モデルの品質がパイプライン設計から切り離されます。ここでの良好な結果は、1 回限りのプロンプトではなく、モデルとこのツールセットを反映しています。
また、代替ハーネス Pydantic を使用してこれらのモデルをテストし、非常に似た結果が得られました。
テスト リポジトリと「ベンチマークをクリーンに保つ」
セキュリティ ベンチマークの難しい部分は、次の 2 つの手段による汚染です。
定期的なモデル トレーニング - 公開されている脆弱性データセットには、その解決策がオンラインで作成され、「偶然」モデル トレーニング データに吸収されているため、モデルは修正を見つけるのではなく、その修正を思い出すことができます。フロンティア モデルでの即時コールの明確な例として、Juice Shop を評価スイートとして使用してみてください。
Bench-maxxing - 悪意のあるモデル作成者は、ベンチマークでモデルをトレーニングします (したがって、benchmaxxing)。これにより、ベンチマークの値が破壊されます。
プライベート合成リポジトリ - 専用に構築され、決して公開されないため、検索するものは何もありません。各 PR には、アクセス制御のバグが 1 つだけ埋め込まれ、隠された回答キーが含まれる現実的な機能が追加されます。このアプリにはパブリック フットプリントがないため、トレーニングでこのアプリを見たモデルはいません。
本物の CVE のリバース リプレイ - すでに修正されている本物のアクセス制御 CVE を取得し、修正を元に戻し、本物の同時代のコミット内に埋め込むので、PR は植え付けられたパズルではなく、通常の作業のように読み取れます。これにより、実稼働コードの現実性が保たれながら、正確なグラウンド トゥルースが得られます。
すべてのレビュー担当者が同じ PR 差分を確認し、その結果とトークン コストを収集し、グラウンド トゥルースに基づいてスコアを付けます。各設定の実行

イオンを 5 回実行すると、1 回のパスでは隠れてしまう実行ごとの差異が平滑化されます。出力ログなどのセッション アーティファクトはすべて保存され、モデルの動作に異常がないか分析されます。

## Original Extract

Skip to main content Docs Blog Recent Posts 2026
The AI PR Security Review Frontier Washington Bans Fable. Anthropic Wins Anyway The Three Hour Attack Window npm Leaves Open by Default Guardrails vs Guidelines The Ceiling for DIY Security Reviews Dam Secure achieves SOC 2 Type 2 Guest Post: How Dam Secure Finds What We Miss Announcing our $4M Seed Round The AI PR Security Review Frontier
Simon Harloff Patrick Collins We wanted to take the guesswork out of which model to use for PR security reviews. To our surprise no Anthropic model reaches the frontier at all and the latest GPT5.6, released only 3 days ago, is the king. Although Fable is known to perform well on a full-code scan, PR reviews are a common workload for our users.
This chart plots Cost per Pull Request vs F1 (finding quality). Blue points define the frontier. See metrics definitions in the table below.
We ran 10 models over the same 10 pull requests, each carrying one planted access-control bug (IDOR, missing auth, broken authorization), five times per model, and scored every finding against the actual code.
The result is a clear cost/quality frontier, and it belongs to OpenAI, xAI and Google rather than Anthropic. GPT-5.6 Sol dominates. Fable 5 performs poorly at ~$3.61 per PR, lands lower on both quality and value. No Anthropic model reaches the frontier at all.
Our benchmark repos are synthetic and private to avoid Benchmaxxing and accidental inclusion in model training.
Fable does perform well on Full Code Scans, but this workload is a PR Scan. A PR Scan is a very common workload for our Users.
This is not a quirk of the Dam Secure harness. We tested the Pydantic harness and Claude Code harness as well and got very similar results.
Benchmark Scope: Pull Requests Only ​
This benchmark is only trying to measure how good models are at finding vulnerabilties in Pull Requests where additional code has been added. This benchmark is NOT about extensively searching an existing codebase for all known vulnerabilities. Our full scan benchmark is coming soon.
That scope matters for how you read the results. The key takeaway is that you are wasting money, usage limits, and tokens if you reach for Fable (or other frontier Claude models) to find security issues in PRs. Fable 5 is the most expensive configuration we tested and still lands off the cost/quality frontier — dominated on both axes by GPT-5.6 Sol, which reaches higher recall at roughly a fifth of the cost per PR.
We will separately release another article on how these agents perform when freely roaming an existing large codebase. The model rankings may look different there; this post is about PR review only.
Each run used the Dam Secure Vulnerability Scanner with high reasoning, repeated five times on a 10-PR corpus of planted access-control bugs.
True Positive (TP) - a vuln found that was expected to be found.
False Positive (FP) - a vuln found that was not expected (not a real bug).
False Negative (FN) - a vuln not found that was expected to be found.
Recall (R) - how well did it find the planted vulnerabilities? It is the fraction of actual positives correctly identified: TP / (TP + FN).
Precision (P) - false positive performance. A fraction of predicted positives that are correct: TP / (TP + FP).
F1 - the overall performance balancing Recall and Precision: 2PR / (P + R). Hovering a point also shows F2 = 5PR / (4P + R), which weights recall ~4× precision.
Costs / PR - is the cost to run that model over all prompts and process the PR regardless of it's accuracy. This is the most accurate representation of real-world cost.
Cost / TP - is the cost per true positive found.
Wikipedia: Precision, Recall and F1
GPT-5.6 Sol is the first model to reach 100% Recall on our harness (clearly we now need to add harder vulnerabilities to find).
The Fable 5 row in the table is labeled "Fable 5 → Opus 4.8 fallback". It handed off to Opus 4.8 only 10.7% of the time.
Prior to GPT5.5 coming out, Gemini 3.5 Flash has been our preferred model for this type of workload. This has previously been our little secret, because our evaluation suites have proven despite the industry wanting to believe that the Opus models dominate security.
Best model is GPT-5.6 Sol - strong F1 (0.91), perfect 100% recall and very cost effective at $0.70 per PR. 5.6 is ~45% cheaper than it's predecessor 5.5 for roughly the same performance.
Honorable mentions: Grok 4.5 lands on the frontier at $0.20 per PR with F1 0.77 (74% recall, 80.4% precision), and Gemini 3.1 Flash Lite anchors the cheap end at ~$0.04 per PR with strong results: F1 0.75 (68% recall, 82.9% precision).
We'll add other Open-weight models to this benchmark.
We construct synthetic pull requests from open source repositories. Each pull request has a vulnerability planted within it. We then measure the models ability to find the planted vulnerability using both the Dam Security Vulnerability Scanner and the Pydantic harness.
We use two different approaches to build up the synthetic dataset. Private synthetic repositories built from an OSS reference repository and reverse-replay of real CVEs (more below).
Each model run uses a limited version of our Dam Secure Vulnerability Scanner configured to match typical tools available in other open source harnesses. It is built on the Vercel AI SDK with a small, security-hardened toolset. The agent sees the PR diff and can:
Read file - Load full file contents when the diff alone is not enough context
Grep - Search the changed files for patterns (auth checks, route handlers, scope validation)
The harness is the same across every model in the table. What varies is only the underlying model. That isolates model quality from pipeline design: a good result here reflects the model plus this toolset, not a one-off prompt.
We have also tested these models using an alternate harness Pydantic, and achieved very similar results.
The Test Repos and "Keeping the Benchmark Clean" ​
The hard part of a security benchmark is contamination through two means:
Regular model training - Public vulnerability datasets have their solutions written up online and "accidentally" absorbed into model training data, so a model can recall the fix instead of finding it. Try using Juice Shop as an evaluation suite as a clear example of immediate call in a frontier model.
Bench-maxxing - Nefarious model builders will train their models on the benchmarks (hence benchmaxxing). This then destroys the value of a benchmark.
A private synthetic repos - Purpose-built and never published, so there is nothing to look up. Each PR adds a realistic feature with exactly one planted access-control bug and a hidden answer key. Because the app has no public footprint, no model has seen it in training.
Reverse-replay of real CVEs - We take a real, already-fixed access-control CVE, revert the fix, and bury it inside genuine same-era commits so the PR reads like ordinary work rather than a planted puzzle. This keeps the realism of production code while still giving us exact ground truth.
Every reviewer sees the same PR diff, we collect its findings and token cost, and we score against ground truth. Running each configuration five times smooths out the run-to-run variance that a single pass would hide. All session artifacts, such as output logs, are kept and analyzed for anomalies in model behavior.
