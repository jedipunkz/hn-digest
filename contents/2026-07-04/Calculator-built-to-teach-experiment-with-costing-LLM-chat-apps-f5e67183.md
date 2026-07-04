---
source: "https://calc.ajinkya.ai/"
hn_url: "https://news.ycombinator.com/item?id=48788306"
title: "Calculator built to teach/experiment with costing LLM chat apps"
article_title: "AI Cost Calculator · Procurement-grade cost model for LLM-based applications"
author: "csmonk"
captured_at: "2026-07-04T20:44:29Z"
capture_tool: "hn-digest"
hn_id: 48788306
score: 1
comments: 0
posted_at: "2026-07-04T19:43:51Z"
tags:
  - hacker-news
  - translated
---

# Calculator built to teach/experiment with costing LLM chat apps

- HN: [48788306](https://news.ycombinator.com/item?id=48788306)
- Source: [calc.ajinkya.ai](https://calc.ajinkya.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T19:43:51Z

## Translation

タイトル: コストのかかる LLM チャット アプリを教育/実験するために構築された計算機
記事のタイトル: AI コスト計算ツール · LLM ベースのアプリケーションの調達グレードのコスト モデル

記事本文:
AIコスト計算機
ベータ版
0 ツール · 0 エージェント · 0 RAG
▾
⬇ Excel形式でダウンロード
🖨 印刷/PDFとして保存
⬇ CSVエクスポート
⌘ 共有可能なリンクをコピー
📋 JSONの表示/編集
↺ デフォルトにリセット
?
ライト
暗い
暖かい
ドロップダウンは次のように置き換えられました
セグメント化された錠剤を開いた直後に #mode-toggle します。) -->
合計 / 月 (オールイン)
—
·
設定する
◐ 1 プロジェクトプロフィール
⚡ 2 グローバルパラメータ
⚒ 3 ツールと料金
▣ 4 エージェント艦隊
✓ 5 ファクトチェック
⚙ 6 エージェントエンジニアリング
☁ 7 ホスティングとインフラ
$ 8 予算ソルバー
⇌ 9 モデル比較
⇕ 10 感度
↗ 11 時間の経過に伴うコスト
⫷12 サイドバイサイド
⇆ 13 現状 vs 提案
◍ 14名
レポート
▤ • レポート
参考文献
∑ • 導出
§ • 方法論
$ • 価格
▦ • ベンチマーク
ホスティング
インフラストラクチャー
コンプライアンス
人材と計画
▶ プロジェクトプロフィール名、チーム、対象者
これは何だろう。デプロイメントの ID — プロジェクト名、所有チームまたは代理店、1 段落の説明、およびホスティング戦略 (マネージド API、BYOK、セルフホスト GPU、オンプレミス)。レポートのタイトル ブロックと、下流セクションが表示されるゲートを駆動します。
なぜそれが重要なのか。ホスティングの選択は、レポートの中で唯一最大の構造上の決定です。これにより、コスト モデル全体が、トークン API ごと (トラフィックに比例)、予約容量 (約束された支出)、設備投資償却 (トラフィックに関係なく毎月固定) の間で切り替わります。間違ってクリックすると、請求額が 5 ～ 10 倍に移動し、下流に誤った予約/セルフホスト フィールドが表示される可能性があります。
結果をどう解釈するか。あなたが望んでいたものではなく、実際の調達手段に一致するホスティングを選択してください。プロジェクト名と説明はレポートの上部にそのまま表示されるため、レビュー担当者に読んでもらいたいように記述してください。この選択に基づいて、下流のセクションが自動表示/自動非表示になります。そうしないと

「予約」を参照してください。予約を使用しないホスティング モデルを使用しています。
LLM が実行される場所を選択します。ダウンストリーム セクション (予約、セルフホスト容量、オンプレミスの償却) は、関連する場合にのみ自動表示されます。
定額のセルフホストコストとして扱われます。セルフホスト容量の GPU/レプリカの計算はバイパスされます。TCO の計算により月次の値が直接提供されます。
セルフホストはクエリの 50% を処理します。 API 構成とセルフホスト構成の両方が適用されます。
AI コスト計算ツールへようこそ
マルチエージェント LLM 導入のための調達グレードの計算ツール。
すべての係数は実際の OpenAI に対して経験的に調整されています /
Anthropic API が実行されます (ベンチ/ハーネスを参照)。予算を見積もるのに使用してください。
ホスティング戦略を比較するか、調達の決定を擁護する
引用。
計算機全体がブラウザ内で実行され、ページからは何も表示されません。
ソースとベンチマークのハーネス:
github.com/ajinkyalkarni/ai-cost-calculator-studio

## Original Extract

AI Cost Calculator
BETA
0 tools · 0 agents · 0 RAG
▾
⬇ Download as Excel
🖨 Print / Save as PDF
⬇ Export CSV
⌘ Copy shareable link
📋 View / Edit JSON
↺ Reset to defaults
?
Light
Dark
Warm
dropdown was replaced by
the segmented-pill #mode-toggle right after open.) -->
Total / mo (all-in)
—
·
Configure
◐ 1 Project profile
⚡ 2 Global parameters
⚒ 3 Tools & rates
▣ 4 Agent fleet
✓ 5 Fact-checking
⚙ 6 Agent engineering
☁ 7 Hosting & infra
$ 8 Budget solver
⇌ 9 Model comparison
⇕ 10 Sensitivity
↗ 11 Cost over time
⫷ 12 Side-by-side
⇆ 13 AS-IS vs proposed
◍ 14 Personnel
Report
▤ • Report
References
∑ • Derivation
§ • Methodology
$ • Prices
▦ • Benchmarks
Hosting
Infrastructure
Compliance
People & Plan
▶ Project profile name, team, audience
What this is. Identity for the deployment — project name, owning team or agency, one-paragraph description, and the hosting strategy (managed API, BYOK, self-hosted GPU, on-prem). Drives the report title block and gates which downstream sections appear.
Why it matters. The hosting choice is the single largest structural decision in the report: it switches the entire cost model between per-token-API (linear with traffic), reserved-capacity (committed spend), and capex-amortized (fixed monthly regardless of traffic). Getting it wrong by mis-clicking can move the bill by 5–10× and surface the wrong reservation / self-host fields downstream.
How to interpret the results. Pick the hosting that matches your actual procurement vehicle — not the one you wish you had. Project name and description appear verbatim at the top of the report, so write them as you'd want a reviewer to read them. Sections downstream auto-show / auto-hide based on this choice; if you don't see Reservations, you're on a hosting model that doesn't use them.
Pick where the LLM runs. Downstream sections (Reservations, Self-host capacity, On-prem amortization) auto-show only when relevant.
Treated as a flat self-host cost. Self-host capacity GPU/replica math is bypassed — your TCO calculation supplies the monthly directly.
Self-host serves 50% of queries. Both API and Self-host configuration apply.
Welcome to the AI Cost Calculator
A procurement-grade calculator for multi-agent LLM deployments.
Every coefficient is empirically calibrated against real OpenAI /
Anthropic API runs (see the bench/ harness). Use it to size a budget,
compare hosting strategies, or defend a procurement decision with
citations.
The whole calculator runs in your browser — nothing leaves the page.
Source & benchmark harness:
github.com/ajinkyakulkarni/ai-cost-calculator-studio
