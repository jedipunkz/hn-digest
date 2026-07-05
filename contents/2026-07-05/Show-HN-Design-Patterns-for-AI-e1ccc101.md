---
source: "https://verificationdesign.com/"
hn_url: "https://news.ycombinator.com/item?id=48792993"
title: "Show HN: Design Patterns for AI"
article_title: "Verification Design"
author: "verify-ai"
captured_at: "2026-07-05T11:24:09Z"
capture_tool: "hn-digest"
hn_id: 48792993
score: 2
comments: 1
posted_at: "2026-07-05T10:36:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Design Patterns for AI

- HN: [48792993](https://news.ycombinator.com/item?id=48792993)
- Source: [verificationdesign.com](https://verificationdesign.com/)
- Score: 2
- Comments: 1
- Posted: 2026-07-05T10:36:40Z

## Translation

タイトル: Show HN: AI のデザイン パターン
記事タイトル: 検証設計
説明: 意見ではなく証拠に基づいて作業を検証するエージェント システムを構築します。
HN テキスト: Gang of Four の Design Patterns は私のお気に入りの本の 1 つです。しっかりと構成されていて読みやすく、参考になりやすかったです。各パターンを手動でコピーして学習し、Java のパターン リポジトリに参照したときのことを覚えています。数週間前、私はインスピレーションを感じ、AI を使用して Arxiv をフィルタリングしてより質の高い研究を実現する研究取り込みパイプラインを構築し、繰り返される最高のアイデアをフィルタリングして、ある種のエグゼクティブ サマリーまたは理念文書にまとめました。この文書は、私が始めるすべてのソフトウェア プロジェクトのガイドとなります。ぜひシェアしたいと思いますので、お役に立てれば幸いです。

記事本文:
検証設計 コンテンツへスキップ 検証設計パターン
意見ではなく証拠に基づいて作業を検証するエージェント システムを構築します。
LLM は単純な自己レビューでは確実に自己修正することができません
1.
AI 設計パターンにより、確率的システムの信頼性が向上します。
差異と、明示的なコンテキストを通じた系統的バイアスへの対抗、
実行可能なチェックと制御されたオーケストレーション。の目標
検証設計は、完全な場合に決定論的な動作に近づくことです。
決定論は利用できません。
サンプリングの不安定性、周囲の状態、コンテキストの汚染、
非同期タイミング、非決定的なツール状態。
おべっか、自己レビュー盲目、裁判官選好バイアス、
確認フレーミング、同族の死角。
バリアンスは新しいカップリングです。アンカリングは新しい結束力です。それぞれ
このカタログのパターンは、それが制約するものに名前を付けています。
モデルが動作する前に存在するもの。 5パターン。
因果関係タグ エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
構成 システムの検証基準を、散在したプロンプトの散文ではなく、明示的でバージョン管理された機械可読データとして表現します。エージェントが発行するすべてのイベントに安定した結合可能な識別子、および該当する場合は親識別子をスタンプして、検証によって因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。

共有環境状態における時間的近接性。
ガードレール デコレーター モデル呼び出し、ツール呼び出し、またはその他のモデル出力境界を、エラーを拒否、置換、サニタイズ、または変換できるポリシー デコレーターでラップします。これにより、ポリシーは、モデルが従うように求められるプロンプトの散文ではなく、モデルが交差する境界のコード内に存在します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
状態ベースライン 検証中のアクションの前に関連する環境またはプロセスの状態をキャプチャすることで、検証者は、すでに存在する状態を受け入れるのではなく、そのアクションが観察された変化を引き起こしたことを証明できます。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
軌道カーソル エージェントが複数ステップのプロセスのどこにいるのか、各境界で何が起こったのかを明示的で構造化した記録を維持することで、検証者と次のターンがチャット履歴やモデルの呼び出しから推測するのではなく、軌道を読み取ることができます。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
判断を観察可能な信号に変える方法。 6パターン。
敵対的フレーム トーンレベルの懐疑指示を、証拠とみなされるものを定義し、拒否する一般的なショートカット パスを指定する許容ルールに置き換えます。

検証者のデフォルトを「もっともらしい場合は受け入れる」から「信頼できる証拠に裏付けられない限り失敗する」に反転します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
Blind Oracle エージェントのドラフト、推論トレース、またはショートカット履歴に基づいて導出を条件付けせずに、仕様、質問、または独立した再実行から期待される証拠を導出します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
Comparator 有限族の名前付き演算子として検証比較を表現します。そのため、判定は「これは正しいか?」というモデルの解釈ではなく、(期待値、観測値、演算子、しきい値、正規化) の決定論的な関数になります。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
デルタ 絶対的な環境状態ではなく、環境状態の変化をアサートすることで、エージェントのアクションの成功を検証します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
実行可能アナログ 主観的な言語ベースの翻訳

検証ステップは、エージェントの判断とは無関係にバイナリの合格/不合格信号を生成する決定論的でプログラム的な実行ステップに組み込まれます。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
裁判官のハーネス LLM 裁判官を摂動、反復、校正、報告の構造的ハーネスで包み込むことで、1 人の裁判官の評決が目に見える一貫性とバイアス制御を備えた測定された信号となるようにします。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
独立性とフィードバックルーティングを通じてバイアスを制御する方法。 6パターン。
敵対者 別のロールの出力で障害を見つけることだけを役割とする構造的に独立したロールを割り当て、そのロールにオーケストレーターが検査できるネガティブ チャネルを発行するよう要求します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
バックプレッシャー ダウンストリームのチェックが失敗した場合、失敗を飲み込んだり盲目的に再試行するのではなく、制限された再試行バジェット内で構造化された再実行コンテキストとして失敗を上流にルーティングします。エージェントが発行するすべてのイベントに安定した結合可能な識別子、および該当する場合は親識別子をスタンプして、観察された影響を感染ではなく特定のエージェントのアクションに帰属することができるようにする

共有された周囲状態における時間的近接性から生じる因果関係を調べる。
クロスファミリー 意図的に異なるモデル ファミリに対して高レバレッジ生成と高レバレッジ評価を実行し、両方の ID を記録するため、共有トレーニング データのバイアスと共有の潜在事前分布が検出されずに検証境界を通過することはありません。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
ディベート モデルの裁量ではなく、オーケストレーション状態で保持されるターン順序、ラウンド数、フェーズ、コンセンサスのしきい値を使用して、決定の前に複数のロール間で制限付きの不一致を実行します。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
エスカレーション チェーン ルートは、型指定され検証されたハンドオフを通じて、より高い権限または異なる機能のハンドラーに動作するため、次の所有者は、誰を呼び出すかについてのモデルのメモリではなく、コード レベルの状態になります。エージェントが発行するすべてのイベントに、安定した結合可能な識別子と、該当する場合は親識別子をスタンプします。これにより、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
ツール アダプター モデルが発行するツール呼び出しを型付き境界で正規化します。スキーマを導出またはフェッチし、呼び出し前に引数を検証し、型付き引数を使用してツールを呼び出し、型付き観察を返します。エージェントが発行するすべてのイベントに、安定した参加可能な識別子と、該当する場合は PA をスタンプします。

レンタル識別子を使用するため、検証では、共有環境状態における時間的近接性から因果関係を推測するのではなく、観察された効果を特定のエージェントのアクションに帰属させることができます。
パターンは 2026 年 7 月 3 日に更新されました。 GitHub 上のソース、チェック、履歴。

## Original Extract

Build agentic systems that verify their work with evidence, not opinion.

Design Patterns by Gang of Four is one of my favorite books. It was well structured and easy to read and reference. I remember when copied each pattern manually to learn them and reference into a pattern repo in Java. A couple weekends ago I felt inspired and I've build a research ingestion pipeline that uses AI to filter Arxiv for better quality research and filtered the best ideas that repeat into a sort of executive summary or ethos document. This document now guides all software projects I begin. I would like to share and I hope you find it useful too!

Verification Design Skip to content Verification Design Patterns
Build agentic systems that verify their work with evidence, not opinion.
LLMs cannot reliably self-correct through naive self-review
1 .
AI design patterns make probabilistic systems more reliable by reducing
variance and countering systematic bias through explicit context,
executable checks, and controlled orchestration. The goal of
verification design is to approach deterministic behavior where full
determinism is unavailable.
Sampling instability, ambient state, context contamination,
asynchronous timing, non-deterministic tool state.
Sycophancy, self-review blindness, judge preference bias,
confirmation framing, same-family blind spots.
Variance is the new coupling. Anchoring is the new cohesion. Each
pattern in this catalog names what it constrains.
What exists before the model acts. 5 patterns.
Causal Tag Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Constitution Represent the system’s verification criteria as explicit, versioned, machine-readable data, rather than as scattered prompt prose. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Guardrail Decorator Wrap a model call, tool call, or other model-output boundary in a policy decorator that can deny, replace, sanitize, or convert errors, so policy lives in code at the boundary the model crosses instead of in prompt prose the model is asked to obey. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
State Baseline Capture the relevant environment or process state before an action under verification, so the verifier can prove the action caused the observed change instead of accepting state that already existed. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Trajectory Cursor Maintain an explicit, structured record of where the agent is in its multi-step process and what happened at each boundary, so the verifier and the next turn can read the trajectory instead of inferring it from chat history or model recall. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
How to turn judgment into observable signals. 6 patterns.
Adversarial Frame Replace tone-level skepticism instructions with admissibility rules that define what counts as proof, name common shortcut paths to reject, and invert the verifier's default from "accept if plausible" to "fail unless backed by trusted evidence." Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Blind Oracle Derive expected evidence from the spec, the question, or independent re-execution without conditioning that derivation on the agent's draft, reasoning trace, or shortcut history. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Comparator Express verification comparison as a named operator from a finite family, so the verdict is a deterministic function of (expected, observed, operator, threshold, normalization) rather than a model's interpretation of "does this look right?" Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Delta Verify the success of an agent's actions by asserting on the change in environment state rather than the absolute environment state. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Executable Analog Translate a subjective, language-based verification step into a deterministic, programmatic execution step that yields a binary pass/fail signal independent of the agent's judgment. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Judge Harness Wrap an LLM judge in a structural harness of perturbation, repetition, calibration, and reporting so that one judge verdict becomes a measured signal with visible consistency and bias controls. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
How to control bias through independence and feedback routing. 6 patterns.
Adversary Assign a structurally separate role whose only job is to find failures in another role's output, and require that role to emit a negative channel the orchestrator can inspect. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Backpressure When a downstream check fails, route the failure back upstream as structured rerun context within a bounded retry budget, instead of swallowing the failure or retrying blindly. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Cross-Family Run high-leverage generation and high-leverage assessment on deliberately different model families, and record both identities, so shared training-data biases and shared latent priors cannot pass undetected through the verification boundary. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Debate Run bounded disagreement among multiple roles before a decision, with turn order, round count, phase, and consensus threshold held in orchestration state instead of model discretion. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Escalation Chain Route work to a higher-authority or different-capability handler through a typed, validated handoff, so the next owner is code-level state instead of a model's memory of who to call. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Tool Adapter Normalize model-emitted tool calls at a typed boundary: derive or fetch a schema, validate arguments before invocation, call the tool with typed arguments, and return a typed observation. Stamp every event the agent emits with a stable joinable identifier and, when applicable, a parent identifier, so verification can attribute observed effects to specific agent actions rather than inferring causality from temporal proximity in shared ambient state.
Patterns updated 2026-07-03. Source, checks, and history on GitHub .
