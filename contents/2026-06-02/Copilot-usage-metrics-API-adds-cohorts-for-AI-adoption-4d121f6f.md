---
source: "https://github.blog/changelog/2026-05-29-copilot-usage-metrics-api-adds-cohorts-for-ai-adoption/"
hn_url: "https://news.ycombinator.com/item?id=48355451"
title: "Copilot usage metrics API adds cohorts for AI adoption"
article_title: "Copilot usage metrics API adds cohorts for AI adoption - GitHub Changelog"
author: "saikatsg"
captured_at: "2026-06-02T00:41:43Z"
capture_tool: "hn-digest"
hn_id: 48355451
score: 4
comments: 1
posted_at: "2026-06-01T11:31:04Z"
tags:
  - hacker-news
  - translated
---

# Copilot usage metrics API adds cohorts for AI adoption

- HN: [48355451](https://news.ycombinator.com/item?id=48355451)
- Source: [github.blog](https://github.blog/changelog/2026-05-29-copilot-usage-metrics-api-adds-cohorts-for-ai-adoption/)
- Score: 4
- Comments: 1
- Posted: 2026-06-01T11:31:04Z

## Translation

タイトル: Copilot 使用状況メトリクス API が AI 導入のコホートを追加
記事のタイトル: Copilot 使用状況メトリクス API が AI 導入のコホートを追加 - GitHub 変更ログ
説明: 誰がアクティブであるかだけでなく、どのように Copilot を使用しているかなど、より深い Copilot 導入ストーリーを伝えるのに役立つように、Copilot 使用状況メトリクス API は、関与している各ユーザーを AI に分類するようになりました。

記事本文:
Copilot 使用状況メトリクス API が AI 導入のコホートを追加 - GitHub 変更ログ
コンテンツにスキップ
サイドバーにスキップ
/
ブログ
変更履歴
GitHub Copilot CLI を試す
新機能を見る
検索
変更履歴
変更履歴に戻る
改善
2026 年 5 月 29 日 •
2 分で読めます
Copilot 使用状況メトリクス API により AI 導入のコホートが追加される
誰がアクティブであるかだけでなく、どのように Copilot を使用しているかなど、Copilot 導入のより深いストーリーを伝えるために、Copilot 使用状況メトリクス API は、28 日間のローリング ウィンドウでの Copilot 製品の使用状況に基づいて、関与している各ユーザーを AI 導入フェーズに分類するようになりました。新しい ai_adoption_phase フィールドはユーザー レベルのレポートで使用でき、新しい totals_by_ai_adoption_phase 配列は企業レベルおよび組織レベルのレポートでフェーズごとのメトリクスを表示します。
関与している各ユーザーは、過去 28 日間の期間のうち少なくとも 2 日間使用した Copilot サーフェスに基づいて、4 つのフェーズのいずれかに割り当てられます。
フェーズ 0 — コホートなし : ユーザーはどのフェーズでもエンゲージメント基準を満たしていません。
フェーズ 1 — コード ファースト : ユーザーはコード補完や IDE エージェント モードを使用します。
フェーズ 2 — エージェントファースト : ユーザーは単一の GitHub ベースのエージェント サーフェス (つまり、Copilot クラウド エージェント、Copilot コード レビュー、または Copilot CLI) を操作します。
フェーズ 3 — マルチエージェント : ユーザーは 2 つ以上の GitHub ベースのエージェント サーフェス、または新しい GitHub Copilot アプリを使用します。
各 ai_adoption_phase 値にはバージョン フィールド ( v1 から始まる) が含まれているため、Copilot の製品面が成長するにつれて、歴史的コンテキストを壊すことなく分類ロジックを進化させることができます。
エンタープライズおよび組織レベルのレポートでは、新しい totals_by_ai_adoption_phase 配列により、エンゲージメントとアクティビティの指標がフェーズごとにグループ化されます。次のものが含まれます。
エンゲージメント ユーザーの合計 (28 年のうち 2 日のエンゲージメント ウィンドウ)
ユーザーが開始したインタラクションの平均
コード生成と受け入れアクティビティ av

時代
追加および削除されたコード行の平均
作成、マージ、レビューされたプル リクエストの平均数
集約されたメトリクスは、合計ではなく、各フェーズ内のユーザーごとの平均を報告します。
成熟度のストーリーを伝える: 単純なアクティブ ユーザー数を超えて、開発者が実際にどの Copilot 機能を採用しているかを示します。
コホートの進行状況を追跡する: ユーザーがコードファーストの使用からエージェントファーストおよびマルチエージェントのワークフローに移行する様子を時間の経過とともに観察します。
ターゲットの実現: 最大のチャンスが見込まれるフェーズでのトレーニング、文書化、展開プログラムに焦点を当てます。
これらのメトリックは、REST API を通じて Copilot の使用状況メトリックにアクセスできるエンタープライズ管理者および組織所有者が利用できます。
このリリースをチーム フィルター シップと組み合わせて使用​​すると、より詳細な精度を得ることができます。
GitHub コミュニティ内のディスカッションに参加してください。
シェアする
コピーされました
共有
変更履歴に戻る
関連記事
開発者向けニュースレターを購読する
開発者向けの隔週ニュースレターでヒント、技術ガイド、ベスト プラクティスをご覧ください。
送信することで、GitHub とその関連会社がパーソナライズされたコミュニケーション、ターゲットを絞った広告、およびキャンペーンの効果のために私の情報を使用することに同意します。詳細については、GitHub のプライバシーに関する声明をご覧ください。
私の個人情報を共有しないでください
LinkedIn アイコン
LinkedIn の GitHub
インスタグラムのアイコン
Instagram の GitHub
YouTube アイコン
YouTube の GitHub
GitHub アイコン
GitHub 上の GitHub の組織

## Original Extract

To help you tell a deeper Copilot adoption story—not just who is active, but how they’re using Copilot—the Copilot usage metrics API now classifies each engaged user into an AI…

Copilot usage metrics API adds cohorts for AI adoption - GitHub Changelog
Skip to content
Skip to sidebar
/
Blog
Changelog
Try GitHub Copilot CLI
See what's new
Search
Changelog
Back to changelog
Improvement
May 29, 2026 •
2 minute read
Copilot usage metrics API adds cohorts for AI adoption
To help you tell a deeper Copilot adoption story—not just who is active, but how they’re using Copilot—the Copilot usage metrics API now classifies each engaged user into an AI adoption phase based on their Copilot product usage over a rolling 28-day window. A new ai_adoption_phase field is available on user-level reports, and a new totals_by_ai_adoption_phase array surfaces per-phase metrics on enterprise- and organization-level reports.
Each engaged user is assigned to one of four phases based on the Copilot surfaces they’ve used on at least two days in the last 28-day window:
Phase 0 — No cohort : User did not meet the engagement criteria for any phase.
Phase 1 — Code first : User engaged with code completion and/or IDE agent mode.
Phase 2 — Agent first : User engaged with a single GitHub-based agent surface (i.e., Copilot cloud agent, Copilot code review, or Copilot CLI).
Phase 3 — Multi-agent : User engaged with two or more GitHub-based agent surfaces, or with the new GitHub Copilot app.
Each ai_adoption_phase value includes a version field (starting at v1 ) so the classification logic can evolve as Copilot’s product surface grows, without breaking historical context.
On the enterprise- and organization-level reports, the new totals_by_ai_adoption_phase array groups engagement and activity metrics by phase, including:
Total engaged users (2-day-in-28 engagement window)
User-initiated interaction average
Code generation and acceptance activity averages
Lines of code added and deleted averages
Pull requests created, merged, and reviewed averages
Aggregated metrics report the average per user within each phase rather than the sum.
Tell the maturity story: Move beyond simple active-user counts and show which Copilot capabilities your developers are actually adopting.
Track cohort progression: Watch users graduate from code-first usage into agent-first and multi-agent workflows over time.
Target enablement: Focus training, documentation, and rollout programs on the phases where you see the biggest opportunity.
These metrics are available to enterprise administrators and organization owners who have access to Copilot usage metrics through the REST API.
You can use this release in combination with the teams filter ship for greater granularity.
Join the discussion within GitHub Community .
Share
Copied
Shared
Back to changelog
Related Posts
Subscribe to our developer newsletter
Discover tips, technical guides, and best practices in our biweekly newsletter just for devs.
By submitting, I agree to let GitHub and its affiliates use my information for personalized communications, targeted advertising, and campaign effectiveness. See the GitHub Privacy Statement for more details.
Do not share my personal information
LinkedIn icon
GitHub on LinkedIn
Instagram icon
GitHub on Instagram
YouTube icon
GitHub on YouTube
GitHub icon
GitHub’s organization on GitHub
