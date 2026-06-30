---
source: "https://blume.codes/blog/how-blume-codebase-improvements-work"
hn_url: "https://news.ycombinator.com/item?id=48731786"
title: "Clustering AI conversations signals for codebase improvements"
article_title: "How Blume codebase improvements work - Blume Blog"
author: "blumeCodes"
captured_at: "2026-06-30T12:37:12Z"
capture_tool: "hn-digest"
hn_id: 48731786
score: 1
comments: 0
posted_at: "2026-06-30T12:30:26Z"
tags:
  - hacker-news
  - translated
---

# Clustering AI conversations signals for codebase improvements

- HN: [48731786](https://news.ycombinator.com/item?id=48731786)
- Source: [blume.codes](https://blume.codes/blog/how-blume-codebase-improvements-work)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T12:30:26Z

## Translation

タイトル: コードベース改善のための AI 会話シグナルのクラスタリング
記事のタイトル: Blume コードベースの改善の仕組み - Blume ブログ
説明: Blume はコーディング エージェントとの会話を監視し、摩擦がないか静かに探します。エージェントが誤解し続けている質問、説明し続けるルール、セットアップとまったく一致しない設定などです。

記事本文:
Blume コードベースの改善の仕組み - Blume ブログ
ホーム ドキュメント ブログ ツール プライバシーとセキュリティについて
ブルーメを入手
左矢印
ブログに戻る
Blume コードベースの改善の仕組み
Blume チーム カレンダー 2026 年 6 月 23 日 5 分で読む 改善の仕組み
Blume はコーディング エージェントとの会話を監視し、摩擦がないか静かに探します —
エージェントが誤解し続けた質問、あなたが再説明し続けたルール、
設定がセットアップとまったく一致しませんでした。同じ問題が発生した場合
多くの場合、Blume は具体的な修正案を作成し、それを示します。あなたは、
diff を確認して適用するかどうかを決定します。
この記事では、この機能の内容、有効にする方法、何が起こるかについて説明します。
ボンネットの下で。
摩擦は瞬間的に作用するのが難しいです。単一の紛らわしいやりとりはそうではありません
設定を変更する価値はありますが、エージェントが 3 回目も同じことを行う
間違った仮定です。埋める価値のある本当のギャップがあります。問題はあなたです
各インスタンスは異なるチャットに埋もれているため、このパターンに気づくことはほとんどありません。
会話に摩擦の兆候が繰り返されていないか注意してください。
チャット全体で同様のシグナルをクラスター化するため、1 回限りのシグナルがノイズを引き起こさないようにします。
新しいルール、フック、スキル、ドキュメントなどの特定の変更を 1 回提案します。
クラスターは十分な証拠によって裏付けられています。
条件を確認して適用し、完全な差分とワンクリックで元に戻すことができます。
最も重要なことは 2 つあります。それは繰り返された証拠に基づいてのみ機能すること、そして実行されることです。
完全にあなたのマシン上で。分析は独自のエージェント CLI を介して行われます
(Claude Code、Codex、または Cursor) 既存のプランのクォータを使用します。何もない
あなたの会話は Blume または第三者にアップロードされます。
改善はデフォルトではオフになっており、オプトインです。上部の「改善」行
提案リストの [設定] をクリックします。次の 2 つを選択します。
スマートなタイミング (推奨) — 最高速度で実行されます。

セッションの最後の 15 回目
プランのクォータ ウィンドウの分、ただしその 25% を超えている間のみ
ウィンドウのクォータはまだ使用されていません。言い換えれば、それはあなたが予定していた割り当てを消費します
とにかく負けて、残りの時間は邪魔にならないようにします。
チャットが静かになったとき — アイドル状態になるとすぐに各会話を分析します。
これにより、全体を通じてクォータを使用するという犠牲を払って、最速のフィードバックが得られます。
日。
を実行する接続されたプラン (クロード コード、コーデックス、カーソル) を選択します。
分析。このリストは、接続されているプラ​​ンから入力されます。選ばなければ、
ブルーメがあなたのために1つを選びます。
[オンにする] をクリックすると、スケジューラが起動します。それ以降、「改善」行には次のように表示されます。
オン · ローカルで実行されます。
有効にすると、「構成」により完全な設定ページが開きます。ここにあるものはすべて
オプトアウト — 適切なデフォルトがすでに選択されており、項目を絞り込みます。
改善するためのハーネス。インストールされているエージェント CLI のチェックリスト。チェック済み
ハーネスが改良される。スキップするには 1 つのチェックを外します。後で取り付ける新しいハーネス
デフォルトでは改善されています。
スコープ — 変更の適用が許可される場所:
プロジェクト — 現在のリポジトリ内のファイル (例: .github/AGENTS.md )。
グローバル — マシン全体の構成 (例: ~/.claude/settings.json )。
各提案はリストにカードとして表示されます。
変更を説明するタイトル (例: 「JSON のクロード コード設定を追加」
応答」）。
証拠バッジ — 「チャット中に 3 回表示されました」 — であることがわかります。
推測ではなく、実際のパターンに基づいています。
1 つ以上の操作行。それぞれの行はアクション (+ create、~
更新、-削除）、アーティファクトのタイプ、短い概要、およびその到達先
(プロジェクトまたはグローバル、どのハーネスに適用されるか)。
「プレビュー」をクリックすると、ファイルごとの統合された差分、絶対的な全体像が表示されます。
接触するパスと期待される効果。プレビューからは次のことができます。
証拠を参照

ence — 変更の動機となった生の会話の抜粋。
現時点で役に立たない場合は、スヌーズまたは無視してください。
適用 — 変更をディスクに書き込みます。
適用された提案、却下された提案、およびスヌーズされた提案は、適用された場合に解決済みに移動します。
それらは元に戻すことができます。
応募はオール・オア・ナッシングです。ブルーメは何かを書く前に、すべてのことを検証します。
ターゲット: 作成はすでに存在していてはならず、更新または削除は一致する必要があります
ファイルは提案が作成されたときとまったく同じです。何かあれば
その下で変更され、何も書かれておらず、提案にフラグが立てられています
自分の仕事を黙って壊すのではなく、レビューしてください。
適用された各変更は、置き換えられた内容を記録するため、「元に戻す」で元に戻すことができます。
以前の内容 — これも、ファイルが Blume が書いたものとまだ一致している場合に限ります。
会話が分析されると、ローカル パイプラインを通過します。各ステージ
判断が必要な場合は、エージェント CLI を呼び出します。残りは単純なローカル ロジックです。
会話
│
§─ 抽出エージェントはトランスクリプトを読み取り、「シグナル」を抽出します
│ (摩擦点、証拠と苦痛スコア付き)
│
§─ クラスタリング エージェントは、新しい信号を類似の信号とグループ化します。
│ 過去の会話 — または新しいクラスターを開始する
│
§─ プロモーション ローカルしきい値は、クラスターに十分な量がいつあるかを決定します。
│ 行動に移すための証拠（信号×会話×痛み）
│
§─ 推進したクラスターを具体化するプランニングエージェント
│ 操作 (作成/更新/削除 + ファイル + 差分)、
│ 許可されたスコープとアーティファクトのタイプ内で
│
└─ 提案はローカルに保存され、レビュー用にリストに表示されます
スケジューラは 1 分に 1 回起動します。 Quiet-chat モードではアイドル状態でキューに入れられます
出現したままの会話 (ティックごとに数件なので、溢れ出ることはありません)。で
スマート タイミング モードでは、プランのセッション ウィンドウに入るまでキューを保持します。
使うか失うかの瀬戸際です。 /blume コマンドは st にジャンプします

すぐ前へ。
なぜなら、プランナーは差分を提案するだけであり、あなたの意見なしで書くことは決してないからです。
適用 — モデルの判断は、制御するステップによって制限されます。あなたの設定
スコープ、アーティファクト、ハーネスは計画段階に渡されるため、
提案は、許可した場所のみをターゲットにできます。
改善により、チャット中に散在する摩擦が、精査された小さな問題に変わります。
エージェント設定に対する一連の変更。それはローカルで、オプトインで、証拠に基づいており、
可逆的: 繰り返される問題を監視し、修正案を作成し、ユーザーが修正を行うのを待ちます。
はいと言います。
すべてのコーディング エージェントを簡単に監視します。

## Original Extract

Blume watches your coding-agent conversations and quietly looks for friction — the questions your agent kept misunderstanding, the rule you keep re-explaining, the config that never quite matched your setup.…

How Blume codebase improvements work - Blume Blog
Home Docs Blog Tools About Privacy & Security
Get Blume
Arrow Left
Back to blog
How Blume codebase improvements work
Blume Team Calendar June 23, 2026 5 min read How Improvements Work
Blume watches your coding-agent conversations and quietly looks for friction —
the questions your agent kept misunderstanding, the rule you keep re-explaining,
the config that never quite matched your setup. When the same problem shows up
often enough, Blume drafts a concrete fix and shows it to you. You review the
diff and decide whether to apply it.
This article explains what the feature does, how to turn it on, and what happens
under the hood.
Friction is hard to act on in the moment. A single confusing exchange isn't
worth changing your setup over — but the third time the agent makes the same
wrong assumption, there's a real gap worth closing. The trouble is that you
rarely notice the pattern, because each instance is buried in a different chat.
Watch your conversations for repeated signals of friction.
Cluster similar signals across chats, so a one-off doesn't trigger noise.
Propose a specific change — a new rule, hook, skill, or doc — once a
cluster is backed by enough evidence.
Review and apply on your terms, with a full diff and one-click undo.
Two things matter most: it only acts on repeated evidence, and it runs
entirely on your machine . The analysis goes through your own agent CLI
(Claude Code, Codex, or Cursor) using your existing plan's quota. Nothing about
your conversations is uploaded to Blume or any third party.
Improvements are off by default — it's opt-in. In the Improve row at the top
of the suggestions list, click Configure . You'll choose two things:
Smart timing (recommended) — runs at the end of a session, in the last 15
minutes of your plan's quota window, but only while more than 25% of that
window's quota is still unused. In other words, it spends quota you were about
to lose anyway, and stays out of your way the rest of the time.
When chats go quiet — analyzes each conversation as soon as it goes idle.
This gives you the fastest feedback, at the cost of using quota throughout the
day.
Pick the connected plan (Claude Code, Codex, Cursor) that should run the
analysis. The list is populated from your connected plans. If you don't choose,
Blume picks one for you.
Click Turn on , and the scheduler starts. From then on, the Improve row shows
On · runs locally .
Once enabled, Configure opens the full settings page. Everything here is
opt-out — sensible defaults are already selected, and you narrow things down.
Harnesses to improve. A checklist of your installed agent CLIs. Checked
harnesses get improved; uncheck one to skip it. New harnesses you install later
are improved by default.
Scopes — where changes are allowed to land:
Project — files in the current repository (e.g. .github/AGENTS.md ).
Global — your machine-wide config (e.g. ~/.claude/settings.json ).
Each suggestion appears as a card in the list:
A title describing the change (e.g. "Add Claude Code preference for JSON
responses").
An evidence badge — "came up 3 times in your chats" — so you can see it's
grounded in a real pattern, not a guess.
One or more operation rows , each showing the action ( + create, ~
update, − remove), the artifact type, a short summary, and where it lands
(project or global, which harness it applies to).
Click Preview to see the full picture: a unified diff per file, the absolute
path it touches, and the expected effect. From the preview you can:
See evidence — the raw conversation excerpts that motivated the change.
Snooze or Dismiss if it's not useful right now.
Apply — write the changes to disk.
Applied, dismissed, and snoozed suggestions move to resolved , where applied
ones can be undone .
Applying is all-or-nothing . Before writing anything, Blume verifies every
target: a create must not already exist, and an update or remove must match
the file exactly as it was when the suggestion was drafted. If anything has
changed underneath it, nothing is written and the suggestion is flagged for
review instead of silently clobbering your work.
Each applied change records what it replaced, so Undo can restore the
previous contents — again, only if the file still matches what Blume wrote.
When a conversation is analyzed, it flows through a local pipeline. Each stage
that needs judgment calls your agent CLI; the rest is plain local logic.
Conversation
│
├─ Extraction agent reads the transcript, pulls out "signals"
│ (friction points, with evidence and a pain score)
│
├─ Clustering agent groups new signals with similar ones from
│ past conversations — or starts a new cluster
│
├─ Promotion local thresholds decide when a cluster has enough
│ evidence (signals × conversations × pain) to act on
│
├─ Planning agent turns a promoted cluster into concrete
│ operations (create/update/remove + file + diff),
│ within your allowed scopes and artifact types
│
└─ Suggestion stored locally, shown in the list for your review
The scheduler wakes once a minute. In quiet-chats mode it queues idle
conversations as they appear (a few per tick, so it never floods). In
smart-timing mode it holds the queue until your plan's session window enters
its use-it-or-lose-it window. A /blume command jumps straight to the front.
Because the planner only ever proposes diffs — and never writes without your
Apply — the model's judgment is bounded by a step you control. Your settings for
scopes, artifacts, and harnesses are passed into the planning stage, so a
suggestion can only target the places you've allowed.
Improvements turn the friction scattered across your chats into a small, vetted
set of changes to your agent setup. It's local, opt-in, evidence-driven, and
reversible: it watches for repeated problems, drafts a fix, and waits for you to
say yes.
Watch every coding agent, effortlessly.
