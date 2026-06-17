---
source: "https://factory.ai/product/router"
hn_url: "https://news.ycombinator.com/item?id=48575781"
title: "Automatic LLM routing that optimizes cost and speed"
article_title: "Factory Router"
author: "terezatizkova"
captured_at: "2026-06-17T21:43:12Z"
capture_tool: "hn-digest"
hn_id: 48575781
score: 1
comments: 0
posted_at: "2026-06-17T19:42:15Z"
tags:
  - hacker-news
  - translated
---

# Automatic LLM routing that optimizes cost and speed

- HN: [48575781](https://news.ycombinator.com/item?id=48575781)
- Source: [factory.ai](https://factory.ai/product/router)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T19:42:15Z

## Translation

タイトル: コストと速度を最適化する自動 LLM ルーティング
記事のタイトル: ファクトリールーター
説明: Factory Router は各タスクに適切なモデルを自動的に選択するため、チームは最低コストで最高の品質を得ることができ、コストを最大 25% 削減しながら最先端のパフォーマンスを維持できます。

記事本文:
ファクトリールーター Factory.ai ロゴ Factory.ai 製品
低コストで最先端のパフォーマンスを実現
Droid セッションごとにモデルを自動選択します。 Factory Router は、各タスクに適切なモデルを選択し、最先端のパフォーマンスを維持し、コストを最大 25% 削減します。
Auto-Model Auto MCP (3) Skills (12) router-classifier classifier · ~2s 最初のユーザー メッセージ、最近のツール呼び出し、およびリポジトリ信号を読み取り、各モデルのスカラー品質確率を出力します。シグナル ウェイト コンピューティング スコア メッセージ 0.30 0.84 最近のツール 0.20 0.62 リポジトリ サイズ 0.15 0.77 言語混合 0.20 0.91 難易度 0.15 0.88 最終スコア 0.80 候補者スコアリングしきい値 0.70 並べ替え 最安 → 最も高価 品質しきい値 キミ K2.6 ムーンショット $ 0.81 MiniMax-M2.7 MiniMax $$ 0.88 Claude Opus 4.7 Anthropic $$$ 0.95 キミ K2.6 ストリーミング › src/auth/middleware.ts... の読み取り › レガシー セッション Cookie 検証の検出 › JWT verify (RS256) への置き換え › エッジ ケースをカバーする 7 つのテストの生成 › PR #418 がオープン — レビューの準備完了 プロンプト分類スコアの実行 問題
AI コーディングのコストは組織全体で上昇しています。
企業の AI コストは上昇しており、トークンの請求額が大きくなったからといって、より多くの作業が行われるわけではありません。パフォーマンスの低下を避けるために、エンジニアは通常、すべてのタスクに対して最もパフォーマンスの高いモデルをデフォルトで使用します。単純な質問、機械的なリファクタリング、ドキュメントの更新、小さなバグの修正、および検索の多い調査は、最終的には最前線のパフォーマンスが本当に必要な作業と同じプレミアム パスに進みます。組織レベルの生産量が明確に増加しない限り、予算は枯渇してしまいます。
タスクごとにモデルを選択するのはやめましょう。
現在、タスクごとにモデルを選択し、安全を確保するために最も高価なモデルを選択しています。 Factory Router を使用すると、一度選択すると、各セッションに最適なモデルが選択されます。
当社のエンタープライズ エンジニアリング ベンチマークについて

。
Claude Opus 4.7 と比較して、Factory Router はセッションあたりのコストを低く抑えながらフロンティアのパフォーマンスを維持します。エンタープライズ規模では、これらの節約はすべての Droid セッションに適用され、費用は最も高価なモデルに一律にデフォルト設定されるのではなく、行われている作業に関連付けられます。
ターミナルベンチ 2 の合格率 · 対 OPUS 4.7 0 % (Claude Opus 4.7 の合格率の %) セッションあたりのコスト · 対 OPUS 4.7 0 % 低い Factory Router は Opus コストの 80 % で実行 成功ごとのコスト · 80.5% の Opus レガシーベンチ パスレート · 対 OPUS 4.7 0 % の Claude Opus 4.7 合格率 セッションあたりのコスト · 対 OPUS 4.7 0 % 低い Factory Router は Opus コストの 75 % で実行 成功した実行あたりのコスト · Opus の 78.0% クロード Opus 4.7 と比較して報告 · フルセッション コストとして測定されたコスト · 複数の実行の平均値 信頼性
プロバイダーの機能が低下したり、レート制限に達したり、容量が制限されたりしても、セッションは続行されます。 Factory Router は、モデル、プロバイダー、容量を超えてルーティングし、99.9% 以上のリクエストの信頼性を実現します。
Claude Opus 4.7 Vertex · 健全なプロバイダーのフェイルオーバー プロバイダー パスが低下した場合、Factory Router は健全なプロバイダーを通じて同じモデルでセッションの実行を維持します。
専用 TPM Enterprise のお客様は、共有されたパブリック キャパシティだけに依存するのではなく、重要な作業用に予約されたスループットを利用できます。
豊富なフロンティア モデル Factory Router は、フロンティア モデルがオンラインになるたびに利用できるようにするため、複雑度の高い作業が最も強力なモデル クラスになります。
米国がホストするオープンソース モデル コスト効率の高いオプションや管理されたオプションが必要な場合は、対象となる作業を米国がホストするオープンソース モデルにルーティングします。
組織の仕組みを反映したルーティング。
ルーティング ガイダンスはチームのコンテキストを Factory Router に取り込むため、自動モデル選択は組織内で実際に作業が行われる方法を反映します。同じ政策が表面化している

ここでは他の Factory モデルが適用されるため、管理者は別のコントロール プレーンを使用せずにアクセス、コンプライアンス、資格を管理します。
標準モデル ポリシー + ルーティング ルールとコンテキスト + 管理者ルーティング ガイダンス すべての Droid セッションでの自動モデル選択 組織全体で有効化されたルーティング ルールとコンテキスト — 日常的なリファクタリング、書式設定、ドキュメントの更新 → コスト効率の高いモデルを優先 — 認証/支払い/より深い推論が必要 → フロンティア モデルを継続 — 検索の多い調査 → オープンソース モデルへのルート キャンセル 保存 今日から利用可能
Factory CLI およびデスクトップ アプリで Factory Router を使用します。
Factory Router は、Factory CLI およびデスクトップ アプリで非公開のリサーチ プレビュー段階にあります。組織で有効にすると、設定を必要とせずにすべてのユーザのモデル ピッカーに表示されます。ミッション ワーカーもこれを使用できるため、長時間実行される自律作業では、インタラクティブなヘッドレス セッションと同じように自動的にモデルが選択され、節約されます。
将来のソフトウェアを構築する準備はできていますか?
@ファクトリー2026 。無断転載を禁じます。

## Original Extract

Factory Router automatically selects the right model for each task so your team gets the best quality at the lowest cost - maintaining frontier performance while cutting costs by up to 25%.

Factory Router Factory.ai Logo Factory.ai Product
Frontier performance at lower cost
Automatic model selection for every Droid session. Factory Router picks the right model for each task, maintains frontier performance, and cuts cost by up to 25%.
Auto-Model Auto MCP (3) Skills (12) router-classifier classifier · ~2s Reads the first user message, recent tool calls and repo signals, then emits a scalar quality probability for each model. Signal Weight Computing Score message 0.30 0.84 recent tools 0.20 0.62 repo size 0.15 0.77 language mix 0.20 0.91 difficulty 0.15 0.88 Final Score 0.80 candidate scoring threshold 0.70 sorted cheapest → most expensive quality_threshold Kimi K2.6 Moonshot $ 0.81 MiniMax-M2.7 MiniMax $$ 0.88 Claude Opus 4.7 Anthropic $$$ 0.95 Kimi K2.6 streaming › Reading src/auth/middleware.ts... › Found legacy session cookie validation › Replacing with JWT verify (RS256) › Generated 7 tests covering edge cases › PR #418 opened — ready for review prompt classify score run The problem
AI coding costs are rising across organizations.
Enterprise AI costs are climbing, and a bigger token bill does not mean more work is getting done. To avoid losing on performance, engineers usually default to the most performant model for all tasks. Simple questions, mechanical refactors, documentation updates, small bug fixes, and search-heavy investigations end up on the same premium path as work that truly needs frontier performance. Budgets get exhausted without a clear increase in organization-level output.
Stop choosing a model for every task.
Today you pick a model per task and lean on the most expensive one to be safe. With Factory Router you choose once and it picks the best model for each session.
On our enterprise engineering benchmarks.
Compared with Claude Opus 4.7, Factory Router maintains frontier performance at lower cost per session. At enterprise scale, those savings apply across every Droid session, with spend tied to the work being done rather than a blanket default to the most expensive model.
TERMINAL-BENCH 2 PASS RATE · vs OPUS 4.7 0 % of Claude Opus 4.7 pass rate COST PER SESSION · vs OPUS 4.7 0 % lower Factory Router runs at 80 % of Opus cost Cost per successful run · 80.5% of Opus LEGACY-BENCH PASS RATE · vs OPUS 4.7 0 % of Claude Opus 4.7 pass rate COST PER SESSION · vs OPUS 4.7 0 % lower Factory Router runs at 75 % of Opus cost Cost per successful run · 78.0% of Opus Reported relative to Claude Opus 4.7 · cost measured as full-session cost · averaged across multiple runs Reliability
When a provider degrades, rate limits hit, or capacity gets constrained, your sessions keep going. Factory Router routes across models, providers, and capacity to deliver 99.9%+ request reliability.
Claude Opus 4.7 Vertex · healthy Provider failover If a provider path degrades, Factory Router keeps the session running on the same model through a healthy provider.
Dedicated TPM Enterprise customers get reserved throughput for critical work instead of relying only on shared public capacity.
Rich frontier models Factory Router keeps frontier models available as they come online, so high-complexity work gets the strongest model class.
US-hosted open-source models Route eligible work to US-hosted open-source models when you need cost-efficient or controlled options.
Routing that reflects how your organization works.
Routing guidance brings your team's context into Factory Router, so automatic model selection reflects how work actually happens inside your organization. The same policy surfaces that govern other Factory models apply here, so admins manage access, compliance, and eligibility without a separate control plane.
Standard model policy + Routing rules and context + Admin routing guidance Automatic model selection for every Droid session Enabled org-wide Routing rules & context — Routine refactors, formatting, and doc updates → favor cost-efficient models — auth/ and payments/ need deeper reasoning → keep on frontier models — Search-heavy investigation → route to open-source models Cancel Save Available today
Use Factory Router in the Factory CLI and Desktop App.
Factory Router is in private research preview in the Factory CLI and Desktop App. Once enabled for your org, it appears in the model picker for every user with no setup required. Mission workers can use it too, so long-running autonomous work gets the same automatic model selection and savings as interactive and headless sessions.
Ready to build the software of the future?
@Factory 2026 . All rights reserved.
