---
source: "https://ai-creed.dev/projects/ai-whisper/"
hn_url: "https://news.ycombinator.com/item?id=48698924"
title: "Show HN: AI-whisper – Claude works better when Codex watches its back"
article_title: "ai-whisper — ai-creed"
author: "vuphanse"
captured_at: "2026-06-27T15:26:24Z"
capture_tool: "hn-digest"
hn_id: 48698924
score: 1
comments: 0
posted_at: "2026-06-27T15:07:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-whisper – Claude works better when Codex watches its back

- HN: [48698924](https://news.ycombinator.com/item?id=48698924)
- Source: [ai-creed.dev](https://ai-creed.dev/projects/ai-whisper/)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T15:07:57Z

## Translation

タイトル: 表示 HN: AI のささやき – コーデックスが背後を監視すると、クロードの動作が向上します
記事タイトル: ai-wisper — ai-creed
説明: 構造化されたワークフローによって駆動される、ペアになった AI コーディング エージェントのターミナルファースト リレー

記事本文:
●ai-creedプロジェクト・bio←プロジェクト
構造化されたワークフローによって駆動される、ペアになった AI コーディング エージェントのターミナルファースト リレー
v0.7.0 — オープンソース、npm、個人的に積極的に使用されています。
ai-whisper は端末内で 2 つのコーディング エージェントをペアリングします。設計によりプロバイダーに依存しないように、Claude、Codex、ezio のいずれか 2 つをマウントします。彼らは 1 つのバトンを共有し、一度に 1 人が順番を所有し、評価者が各ラウンドをゲートする構造化されたワークフローによってコラボレーションが促進されます。リレーは 1 つの取り決めに縛られることはありません。今日のワークフローでは、実装者とレビュー者がペアになっています (仕様駆動開発、ラルフ ループ、複雑なバグ修正、審議)。
ハード エージェントのタスクは、2 人目のエージェントがレビューすることでよりうまくいきます。 Whisper を使用すると、ペアリングが 2 台の端末が互いに通信しているふりをするのではなく、1 つのワークフローのように感じられます。
ai-whisper はこのように構築されました。ほぼすべての機能は仕様として開始され、独自の仕様駆動型開発ワークフローを通じて実行されます。
構造化されたワークフロー — 仕様駆動開発、ラルフロープ、複雑なバグ修正、熟慮
シングルバトンリレー — 群れではなく、一度に 1 人の所有者
役割にとらわれない — ワークフローは役割を割り当てます (現在: 実装者 + レビュー者)
LLM エバリュエーターによってゲート制御される自律ループ
一時停止可能、再開可能な実行 - 飛行中に正常な実行を一時停止し、スタックしたものを回復し、再起動しません。
真実の情報源として実際にマウントされたクロード / コーデックス / ezio セッション
ワークフローは、フェーズ、ゲート、ラウンド予算などの構造化されたループであり、長いプロンプトではありません。ワークの形状から選ぶ：
仕様駆動開発 — 「完了」を前もって説明できる場合。仕様を計画に落とし込み、レビューの下で実行します。
ラルフ・ループ — それができないとき。オープンエンドの目標をチャンクごとに検討し、レビュー担当者がそれぞれのゴールをゲートします。
複雑なバグ修正 — バグは報告されているが、根本的な原因は報告されていない場合。 T

3 つのフェーズ (診断 → 修正と検証 → 事後分析) があり、実装者はバグを再現する必要があります。これは、コードを読んでの推測ではなく、コミットされた失敗テストです。
熟慮 — アイデアがまだ曖昧で、まだ「完了」と表現できないとき。 Explorer と Challenger は空間をマッピングし (目的 → アプローチ → トレードオフ → 総合)、ブレインストーミングや SDD に取り入れるためのコミットされた調査結果ドキュメントを返します。
どちらを選択しても、仕様、目標、またはバグ レポートを作成します。実行は設定した基準に収束します。そのため、正確なアーティファクトが結果のほとんどを占めます。
Autonomous は船の決定ではなく、ループをカバーします。完成した実行は強力なドラフトです。着陸する前に、差分を読み、実行し、QA を行います。 ai-wisper は収束を行い、迅速に判断するための完全な軌跡を提供します。
本当のスペック主導の開発実行: クロード (左) とコーデックス (中央) はそれぞれのマウントされたセッションで作業し、ダッシュボード (右) はバトンのハンドオフとフェーズごとの判定を追跡します。
npm install -g ai-whisper
リポジトリ: github.com/ai-creed/ai-whisper →
ai-creed · ヴー・ファンと彼の友人たちによって構築 ·
ギットハブ ·
vu.phan.se@gmail.com
· 2026年

## Original Extract

terminal-first relay for paired ai coding agents, driven by structured workflows

● ai-creed projects · bio ← projects
terminal-first relay for paired ai coding agents, driven by structured workflows
v0.7.0 — open source, on npm, actively used personally.
ai-whisper pairs two coding agents in your terminal — mount any two of Claude, Codex, and ezio , provider-agnostic by design. They share a single baton — one owns the turn at a time — and structured workflows drive the collaboration, with an evaluator gating each round. The relay isn’t tied to any one arrangement; today’s workflows pair an implementer and a reviewer (spec-driven-development, ralph-loop, complex-bug-fixing, deliberation).
Hard agent tasks go better with a second agent reviewing. Whisper makes that pairing feel like one workflow instead of two terminals pretending to talk to each other.
ai-whisper was built this way: nearly every feature started as a spec run through its own spec-driven-development workflow.
structured workflows — spec-driven-development, ralph-loop, complex-bug-fixing, deliberation
single-baton relay — one owner at a time, not a swarm
role-agnostic — workflows assign the roles (today: implementer + reviewer)
autonomous loops gated by an llm evaluator
pausable, resumable runs — pause a healthy run mid-flight, recover a stuck one, don't restart
real mounted claude / codex / ezio sessions as the source of truth
A workflow is a structured loop — phases, gates, round budgets — not a long prompt. Pick by the shape of the work:
spec-driven-development — when you can describe “done” up front. Sharpens a spec into a plan, then executes it under review.
ralph-loop — when you can’t. Grinds an open-ended goal chunk-by-chunk, a reviewer gating each one.
complex-bug-fixing — when a bug is reported but the root cause isn’t. Three phases (diagnosis → fix-and-verify → post-mortem), with the implementer required to reproduce the bug — a committed failing test, not speculation from reading code.
deliberation — when the idea is still fuzzy and you can’t describe “done” yet. An Explorer and Challenger map the space (objectives → approaches → tradeoffs → synthesis), then hand back a committed findings doc to take into brainstorming or SDD.
Whichever you pick, you write the spec, goal, or bug report; the run converges on the criteria you set — so a precise artifact is most of the outcome.
Autonomous covers the loop , not the ship decision. A finished run is a strong draft: you still read the diff, run it, and QA it before it lands. ai-whisper does the convergence and hands you the full trail to judge fast.
A real spec-driven-development run: Claude (left) and Codex (middle) work in their own mounted sessions while the dashboard (right) tracks the baton handoffs and per-phase verdicts.
npm install -g ai-whisper
repo: github.com/ai-creed/ai-whisper →
ai-creed · built by Vu Phan and his friends ·
github ·
vu.phan.se@gmail.com
· 2026
