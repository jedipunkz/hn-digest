---
source: "https://devloop.sh"
hn_url: "https://news.ycombinator.com/item?id=48549774"
title: "Show HN: Devloop, a local code/review loop for Codex and Claude Code"
article_title: "devloop · Spec in. Reviewed code out."
author: "satyaborg"
captured_at: "2026-06-16T04:38:57Z"
capture_tool: "hn-digest"
hn_id: 48549774
score: 1
comments: 0
posted_at: "2026-06-16T02:22:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Devloop, a local code/review loop for Codex and Claude Code

- HN: [48549774](https://news.ycombinator.com/item?id=48549774)
- Source: [devloop.sh](https://devloop.sh)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T02:22:30Z

## Translation

タイトル: Show HN: Devloop、Codex および Claude Code のローカル コード/レビュー ループ
記事のタイトル: devloop · 仕様入力。コード出力をレビューしました。
説明: 仕様主導のローカル エージェント ループ: Codex の実装、Claude Code のレビュー、およびすべての受け入れ基準が満たされるまでの再試行。

記事本文:
░█▀▄░█▀▀░█░█░█░░░█▀█░█▀█░█▀█
░█░█░█▀▀░▀▄▀░█░░░█░█░█░█░█▀▀
░▀▀░░▀▀▀░░▀░░▀▀▀░▀▀▀░▀▀▀░▀░░
すべてが完了するまで実行される仕様主導のコード/レビュー ループ
受け入れ基準が満たされており、エージェントによってバグが修正されています。
$カール -fsSL https://devloop.sh/install |バッシュ
コーデックスとクロード コードが必要です。 macOS上で動作します。
開発者は現在、ほとんどの時間をコードに費やしています
コーディングエージェントの支援を受けてレビューします。でも同じ
モデル ファミリはコーディングとレビューの両方を行うことができません。それ
自身の出力を優先します。 Devloop は Codex を実装し、
Claude コードを敵対的にレビュー (またはその逆)、
すべての基準が満たされ、レビューが解決されるまで、
人間は、判断が報われる場所にのみやって来ます。
仕様と PR の承認段階。
[ * ]
仕様:
/devloop-spec スキルは
ショートマークダウン
明確な問題、結果、受け入れを伴う仕様
基準と、何が範囲内であるか
アウト。
[ * ]
コード : コーデックス
分離された git ワークツリーに仕様を実装します。
[ * ]
レビュー：クロード
コードは /devloop-review に対して実行されます
仕様を
すべての受け入れゲートを確保し、
エンジニアリングゲート（セキュリティ、保守性、
正確さ、単純さなど）が満たされています。
[ * ]
ループ : 繰り返します
拒否、受け入れまたは最大パスで停止、および
人間によるレビューに備えた PR を生成します。
Devloop には、両方のエージェントにインストールされる 2 つのスキルが同梱されています。
[ * ]
/devloop-spec : 大まかなアイデアを変えます。
メモ、URL、またはインタビューを 1 つにまとめます
具体的な開発ループ対応仕様。
[ * ]
/devloop-review : 各パスを判定します
仕様およびエンジニアリング品質のゲートに対して、
修正を伴う ACCEPT、REJECT、または UNCLEAR を返す
指示。
[ ]
作業を明確に固定することができます
インタビューによるスペック
コーディングエージェント。
[ ]
変化は大きい

ge またはマルチステップで十分
敵対的なレビューが実際にキャッチしていること
何か。
[ ]
エージェントに反復してもらいたい
無人で、あなたを引き込むだけです
仕様と PR の承認。
[ ]
タスクはワンショット可能であり、スコープは
小さなものです。エージェントに直接聞いてください。
[ ]
仕事は探求的または自由なものであり、
受け入れ基準が当てはまらない場合
前に書かれています。
[ ]
すべての編集を実践的にコントロールしたい
ループが単独で実行されるのではなく、
いいえ。分離された兄弟ワークツリーで実行されます。
合格しない限り、PR をプッシュしたり開いたりしないでください
--create-pr 。パス
--その場で作業できる
代わりに現在のチェックアウト。
デフォルトではコーデックスコードとクロードレビュー。
どちらかを --coder / と入れ替えます
--査読者 。
はい。 .devloop/verify フックを追加する
ビルド、lint、またはテスト コマンドを使用します。それ
受け入れをゲートし、ループをブロックします。
失敗します。
ACCEPT時、失速時、または在庫切れの場合
パスの数 (デフォルトでは最大 5、クランプ済み)
1-10）。各実行には 30 分のタイムアウトが上限となります。
なし。監査可能な Bash スクリプト 1 つ。トラックと
レポートは .devloop/ の下に配置されます。

## Original Extract

Spec-driven local agent loop: Codex implements, Claude Code reviews, and retries until all acceptance criteria are met.

░█▀▄░█▀▀░█░█░█░░░█▀█░█▀█░█▀█
░█░█░█▀▀░▀▄▀░█░░░█░█░█░█░█▀▀
░▀▀░░▀▀▀░░▀░░▀▀▀░▀▀▀░▀▀▀░▀░░
Spec-driven code/review loop that runs until all
acceptance criteria are met and bugs fixed by agents.
$ curl -fsSL https://devloop.sh/install | bash
Requires Codex and Claude Code. Runs on macOS.
Developers now spend most of their time in code
review, assisted by coding agents. But the same
model family can't both code and review. It
favors its own outputs. Devloop makes Codex implement and
Claude Code review adversarially (or vice versa),
until all criteria are met and reviews addressed, so
humans only come in where judgment pays off: at the
spec and PR sign-off stages.
[ * ]
Spec :
/devloop-spec skill generates a
short markdown
spec with a clear problem, outcome, acceptance
criteria, and what's in scope as much as what's
out.
[ * ]
Code : Codex
implements the spec in isolated git worktrees.
[ * ]
Review : Claude
Code runs /devloop-review against
the spec to
ensure all acceptance gates along with
engineering gates (security, maintainability,
correctness, simplicity, etc.) are met.
[ * ]
Loop : Iterates on
reject, stops on accept or max passes, and
generates a PR ready for human review.
Devloop ships two skills, installed for both agents:
[ * ]
/devloop-spec : turns a rough idea,
notes, a URL, or an interview into one
concrete, devloop-ready spec.
[ * ]
/devloop-review : judges each pass
against the spec and engineering quality gates,
returning ACCEPT, REJECT, or UNCLEAR with fix
instructions.
[ ]
The work can be pinned down in a clear
spec through an interview with the
coding agent.
[ ]
The change is large or multi-step enough
that adversarial review actually catches
something.
[ ]
You want the agents to iterate
unattended and only pull you in for the
spec and PR sign-off.
[ ]
The task is one-shotable and the scope
is tiny: just ask the agent directly.
[ ]
The work is exploratory or open-ended,
where acceptance criteria can't be
written up front.
[ ]
You want hands-on control of every edit
rather than a loop running on its own.
No. It runs in isolated sibling worktrees and
never pushes or opens a PR unless you pass
--create-pr . Pass
--in-place to work in your
current checkout instead.
Codex codes and Claude reviews by default.
Swap either with --coder /
--reviewer .
Yes. Add a .devloop/verify hook
with your build, lint, or test command. It
gates acceptance and blocks the loop when it
fails.
On ACCEPT, on a stall, or when it runs out
of passes (up to 5 by default, clamped
1-10). A 30-minute timeout caps each run.
None. One auditable Bash script; tracks and
reports land under .devloop/ .
