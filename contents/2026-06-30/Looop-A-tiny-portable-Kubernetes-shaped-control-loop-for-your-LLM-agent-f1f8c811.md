---
source: "https://github.com/yusukeshib/looop"
hn_url: "https://news.ycombinator.com/item?id=48733755"
title: "Looop – A tiny, portable, Kubernetes-shaped control loop for your LLM agent"
article_title: "GitHub - yusukeshib/looop: A tiny, portable, Kubernetes-shaped control loop for your work: senses your world and makes one LLM-driven move per beat toward your goals. · GitHub"
author: "johnnydamacha"
captured_at: "2026-06-30T15:50:05Z"
capture_tool: "hn-digest"
hn_id: 48733755
score: 1
comments: 0
posted_at: "2026-06-30T15:05:18Z"
tags:
  - hacker-news
  - translated
---

# Looop – A tiny, portable, Kubernetes-shaped control loop for your LLM agent

- HN: [48733755](https://news.ycombinator.com/item?id=48733755)
- Source: [github.com](https://github.com/yusukeshib/looop)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:05:18Z

## Translation

タイトル: Looop – LLM エージェント用の小型でポータブルな Kubernetes 型の制御ループ
記事のタイトル: GitHub - yuukeshib/looop: 作業用の小型でポータブルな Kubernetes 型の制御ループ: 世界を感知し、目標に向かって 1 ビートごとに LLM 駆動の動きを 1 つ実行します。 · GitHub
説明: 作業用の小型でポータブルな Kubernetes 型の制御ループ: 世界を感知し、目標に向かって 1 ビートごとに LLM 主導の動きを 1 つ実行します。 - ユウスケシブ/looop

記事本文:
GitHub - yuukeshib/looop: 作業用の小型でポータブルな Kubernetes 型の制御ループ: 世界を感知し、目標に向かって 1 ビートごとに LLM 駆動の動きを 1 つ実行します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ユウスケシブ
/
ループ
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
104 コミット 104 コミット .github/ workflows .github/ workflows src src .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md flake.lock flake.lock flake.nix flake.nix install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント主導の作業のための、小型でポータブルな自律制御ループ。 1 つ
自己完結型バイナリ — データベース、サーバー、ヘルパー ファイルはありません。
looop は頭脳であり、タスクランナーではありません。あなたが気になるものを監視します
(GitHub、Linear、Grafana など) を実行し、ワーカー エージェントのフリートを実行します。それぞれが勝ちました
世界を感じ取り、何かが変化した場合、最も重要なことを一つ決定する
ワーカーの生成を含めて、それを移動して実行します。裁きは自分の中に生きている
ループ (ビートごとの小さなゲートされた LLM 呼び出し)。
自律ループは簡単です。難しい部分 - そしてlooopの要点
デザイン — 人間がどこでどのようにループに参加するかです。人間多すぎてそれ
自律的ではありません。少なすぎて無謀です。 looop の答えはあなたを引き込むことです
それはまさに 2 種類の瞬間であり、他の場所ではありません。
人間が常に最新情報を把握し続ける方法
ループに触れる方法は 2 つあります。その分割がデザインです。
ステアリング — 非同期、あなたが始めます。あなたはドライバーではなく、仲間です。あなたが何を形作るか
looopは目標とPLAYBOOKを編集して追求します。次のビートを観察します。これ
決してループをブロックしないでください。方向を決めてそこから立ち去ります。
looop _ goal write ship-v2 - # 望ましい状態を宣言します (次のビートに有効)
looop _ プレイブックの書き込み - # あなたの判断、優先順位、ガードレール
回答 — 同期すると、ループが開始されます。 looop は、必要な場合にのみ連絡します。
本当にそうしなければなりません: 労働者は人間だけが決定を下します。

メーカー、または
不可逆的なアクション (マージ、デプロイ、削除) には、明示的な「はい」が必要です。ブロックします
そしてあなたの電話を待ちます。
looop _ wait --only-asks # ループが必要になるまで安価にブロックします
looop _answer < id > " yes " # ワーカーのブロックを解除 / ゲートを承認
重要な点は、介入ポイントが UI から切り離されることです。質問して、
答えは、バックエンドに依存しない 1 つのコントラクトを通じて到達できる耐久性のあるファイル メールボックスです
( looop _ … ) なので、ループが特定の端末、tmux、または stdin でブロックされることはありません。
— どのチャネルからでも、最終的には答えが必要なだけです。
裸のターミナル - 自分で動詞を入力します (最も薄いクライアント)。
エージェント コンシェルジュ — 中継する claude / codex / opencode / pi セッション
わかりやすい言葉で質問し、あなたに代わって答えます。
通知スクリプト — Slack/SMS に質問をプッシュし、応答を中継するループ。
クライアントはインターフェースであり、決して意思決定者ではありません。 looopが決定します。クライアント
質問をあなたに伝えて、あなたの答えを返すだけです。
2 つのプロパティにより、これらすべてが信頼できるものになります。
レベルトリガー。すべての状態はプレーン ファイルであるため、ループはビートごとに再検知します
クラッシュしたパルスは再起動時にファイルを再読み取りするだけです。保留中の質問は存続します
再起動 — キューがなく、作業が失われることもありません。
1 拍につき 1 つの動き。各ビートは最大でも 1 つのことを行います。 1 日の予算の上限
費やす。動作は読みやすく低コストのままです。変化しない世界では LLM 呼び出しのコストがかかりません。
ワンビート：感覚→決断→行動
SENSE — すべての sensors/*.sh を実行し、スナップショット/ を更新します。変わらない世界
最後のビート以降 → ここで停止、LLM 呼び出しなし。
決定 — 変更の際は、PLAYBOOK + 目標 + 測定値 + 保留中の質問を渡します。
LLM は 1 つの型指定された手を返します。
ACT — 実行します。ゴール/センサー/PLAYBOOK を書き込み、1 つの元に戻せるコマンドを実行します。
またはワーカーを生成します。不可逆的な動きはゲートされています - 彼らはあなたの答えを待ちます
(s

上記 ee)、人間のみの判断を下す従業員も同様です。
レイヤー
それは何ですか
コア
自律的なパルスとその背後にある耐久性のある状態。決断し、行動する。
契約
ループ _ … 動詞 — コアを読み取り、操作するための、バックエンドに依存しない 1 つの安定したサーフェス。
クライアント
人間の契約を推進するもの (端末/コンシェルジュ/通知)。インターフェイスであって、決して意思決定者ではありません。
状態はデータ ディレクトリ内のプレーン ファイルであり、コントラクトを通じて到達されます。
パブリックインターフェース:
カール -fsSL https://raw.githubusercontent.com/yuukeshib/looop/main/install.sh |バッシュ
# または
カーゴインストールループ
唯一のハード依存: LLM ランナー。クロードがデフォルトです。コーデックス 、
opencode と pi も機能します。looop init を使用したものを選択してください。 (触る労働者
コードは git worktree 経由で独自のサンドボックスを分離するか、利用可能な場合はボックスを作成します。
ループ依存関係ではなく、ワーカー規約。)
looop init # 対話型セットアップ — `up` の前に必要です。ランナーの配線を選択します
looop up # 自律パルスを開始します (切り離されています)
looop watch # ライブログ + 実行セッションセレクター (読み取り専用)
looop down # パルスとすべてのワーカーを停止します
looop init は looop up の前に必要です。パルスは開始するまで拒否します。
ランナー配線が存在するため、すべてのティックとワーカーを駆動するエージェント CLI は
明示的な選択であり、サイレントデフォルトではありません。コアを読んで操縦するには、
looop _ … 手動で動詞を作成する (looop _ state 、 _ wait 、 _ Answer 、 _ goal write )
または、クライアントに運転してもらいます。
looop はヘッドレスで実行されるため、インタビューすることはできません。新しいデータ ディレクトリには、
スターター PLAYBOOK、セットアップの目標、実際の保留中のセットアップについてクライアントに尋ねる
待っているとすぐに目が覚めるように尋ねます。最も簡単な答えはエージェントです
クライアント (「コンシェルジュ」) — 会話する claude / codex / opencode / pi セッション
平易な言葉:
クロード # それからこう言います:
#「私のlooopコンシェルになってください」

ge: 「looop up」を実行し、セットアップの目標を伝え、インタビューします
# 私に目標、センサー、PLAYBOOK を書いてもらいます。」
コンシェルジュは looop up を実行し、保留中のセットアップの質問を表示し、あなたの設定を編集します。
動詞の書き込みによる目標/PLAYBOOK — 運転中に平易な言葉で話す
契約。カスタマイズしたら、スターターの質問に答えて、セットアップの目標をアーカイブします。
そこからループが実行されます。
コンシェルジュを完全にスキップして、手動で操縦することもできます。詳細については、looop ヘルプを参照してください。
完全なコマンドリファレンスと設計マニュアル。
設定 ( $LOOOP_CONFIG 、デフォルト ~/.config/looop/config.json ) は 2 つだけです
シェルコマンド。 looop init を使用すると、 claude 、 codex 、 opencode 、 pi 、
またはカスタム;その後、ループは結果を単純なランナー配線として処理します。
{
"tick_command" : " claude -p --output-format stream-json --verbose --dangerously-skip-permissions --model Sonnet \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " クロード --dangerously-skip-permissions --model opus \" $(cat {{prompt_file}}) \" "
}
コーデックス
{
"tick_command" : " codex exec --json --dangerously-bypass-approvals-and-sandbox \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " codex --dangerously-bypass-approvals-and-sandbox \" $(cat {{prompt_file}}) \" "
}
opencode (ベストエフォート — インストールされているバージョンに対して検証します)
{
"tick_command" : " opencode を実行 \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " opencode \" $(cat {{prompt_file}}) \" "
}
円周率
{
"tick_command" : " pi -p --mode json -ne --model claude-sonnet-4-5 -- Thinking low @{{prompt_file}} " ,
"worker_command" : " pi --model claude-opus-4-8 -- Thinking Medium @{{prompt_file}} "
}
上記のモデル ID は例です。クロードにとって、ソネット/作品は常に別名です。
それぞれの最新のものに解決します。特定のバージョンを固定する (例:
--model claude-opus-4-1 ) 再現性が必要な場合。
小型でポータブルな Kubernetes-s

仕事のための計画的な制御ループ: 世界を感知し、目標に向かって 1 ビートごとに LLM 主導の動きを 1 つ実行します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
55
v0.37.0
最新の
2026 年 6 月 30 日
+ 54 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A tiny, portable, Kubernetes-shaped control loop for your work: senses your world and makes one LLM-driven move per beat toward your goals. - yusukeshib/looop

GitHub - yusukeshib/looop: A tiny, portable, Kubernetes-shaped control loop for your work: senses your world and makes one LLM-driven move per beat toward your goals. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
yusukeshib
/
looop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
104 Commits 104 Commits .github/ workflows .github/ workflows src src .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md flake.lock flake.lock flake.nix flake.nix install.sh install.sh View all files Repository files navigation
A tiny, portable, autonomous control loop for agent-driven work. One
self-contained binary — no database, no server, no helper files.
looop is the brain, not a task runner. It watches the things you care about
(GitHub, Linear, Grafana, …) and runs a fleet of worker agents. Each beat it
senses the world and, if something changed, decides the single most important
move and executes it — including spawning workers. The judgment lives inside
looop (a small, gated LLM call per beat).
An autonomous loop is easy. The hard part — and the whole point of looop's
design — is where and how a human enters the loop. Too much human and it
isn't autonomous; too little and it's reckless. looop's answer is to pull you in
at exactly two kinds of moments, and nowhere else.
How the human stays in the loop
There are two distinct ways you touch the loop — and that split is the design.
Steer — async, you initiate. You are a peer, not a driver. You shape what
looop pursues by editing goals and the PLAYBOOK; it observes them next beat. This
never blocks the loop — you set direction and walk away.
looop _ goal write ship-v2 - # declare desired state (effective next beat)
looop _ playbook write - # your judgment, priorities, guardrails
Answer — sync, the loop initiates. looop reaches back for you only when it
genuinely must: a worker hits a decision only a human can make, or an
irreversible action — merge, deploy, delete — needs an explicit yes. It blocks
and waits for your call.
looop _ wait --only-asks # block cheaply until the loop needs you
looop _ answer < id > " yes " # unblock the worker / approve the gate
The key move: the intervention point is decoupled from any UI. Asks and
answers are a durable file mailbox reached through one backend-agnostic contract
( looop _ … ), so the loop never blocks on a particular terminal, tmux, or stdin
— it just needs an answer eventually , from whatever channel reaches you:
a bare terminal — you typing the verbs yourself (the thinnest client);
an agent concierge — a claude / codex / opencode / pi session that relays
asks in plain language and answers on your behalf;
a notify script — a loop that pushes asks to Slack/SMS and relays your reply.
A client is an interface , never a decision-maker. looop decides; the client
just carries the question to you and your answer back.
Two properties make all this dependable:
Level-triggered. All state is plain files, so the loop re-senses every beat
and a crashed pulse just re-reads its files on restart. A pending ask survives
restarts — no queues, no lost work.
One move per beat. Each beat does at most one thing; a daily budget caps
spend. Behavior stays legible and cheap — an unchanged world costs no LLM call.
One beat: sense → decide → act
SENSE — run every sensors/*.sh , refreshing snapshots/ . World unchanged
since last beat → stop here, no LLM call.
DECIDE — on change, hand the PLAYBOOK + goals + readings + pending asks to
the LLM, which returns one typed move.
ACT — execute it: write a goal/sensor/PLAYBOOK, run one reversible command,
or spawn a worker. Irreversible moves are gated — they wait for your answer
(see above), and so does any worker that hits a human-only decision.
Layer
What it is
core
the autonomous pulse + the durable state behind it. Decides and acts.
contract
the looop _ … verbs — the one stable, backend-agnostic surface to read and steer core.
client
anything that drives the contract for a human (terminal / concierge / notify). An interface, never a decision-maker.
State is plain files in the data dir, reached through the contract — not a
public interface:
curl -fsSL https://raw.githubusercontent.com/yusukeshib/looop/main/install.sh | bash
# or
cargo install looop
Only hard dependency: an LLM runner. claude is the default; codex ,
opencode , and pi also work — pick one with looop init . (Workers that touch
code isolate their own sandbox via git worktree , or box if available — a
worker convention, not a looop dependency.)
looop init # interactive setup — required before `up`; pick the runner wiring
looop up # start the autonomous pulse (detached)
looop watch # live log + running-session selector (read-only)
looop down # stop the pulse and all workers
looop init is required before looop up : the pulse refuses to start until
the runner wiring exists, so the agent CLI driving every tick and worker is an
explicit choice, never a silent default. To read and steer core, drive the
looop _ … verbs by hand ( looop _ state , _ wait , _ answer , _ goal write )
or let a client drive them for you.
looop runs headless, so it can't interview you. A fresh data dir is seeded with a
starter PLAYBOOK, a setup goal, and a real pending setup ask so a client
waiting on asks wakes immediately. The simplest way to answer is an agent
client ("concierge") — a claude / codex / opencode / pi session you talk to in
plain language:
claude # then say:
# "be my looop concierge: run `looop up`, relay the setup goal, and interview
# me to write my goals, sensors, and PLAYBOOK."
The concierge runs looop up , surfaces the pending setup ask, and edits your
goals/PLAYBOOK via the write verbs — speaking plain language while driving the
contract. Once customized, answer the starter ask and archive the setup goal;
looop runs from there.
You can skip the concierge entirely and steer by hand. See looop help for the
full command reference and design manual.
The config ( $LOOOP_CONFIG , default ~/.config/looop/config.json ) is just two
shell commands . looop init lets you pick claude , codex , opencode , pi ,
or custom ; after that looop treats the result as plain runner wiring:
{
"tick_command" : " claude -p --output-format stream-json --verbose --dangerously-skip-permissions --model sonnet \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " claude --dangerously-skip-permissions --model opus \" $(cat {{prompt_file}}) \" "
}
codex
{
"tick_command" : " codex exec --json --dangerously-bypass-approvals-and-sandbox \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " codex --dangerously-bypass-approvals-and-sandbox \" $(cat {{prompt_file}}) \" "
}
opencode (best-effort — verify against your installed version)
{
"tick_command" : " opencode run \" $(cat {{prompt_file}}) \" " ,
"worker_command" : " opencode \" $(cat {{prompt_file}}) \" "
}
pi
{
"tick_command" : " pi -p --mode json -ne --model claude-sonnet-4-5 --thinking low @{{prompt_file}} " ,
"worker_command" : " pi --model claude-opus-4-8 --thinking medium @{{prompt_file}} "
}
Model ids above are examples. For claude, sonnet / opus are aliases that always
resolve to the latest of each; pin a specific version (e.g.
--model claude-opus-4-1 ) if you need reproducibility.
A tiny, portable, Kubernetes-shaped control loop for your work: senses your world and makes one LLM-driven move per beat toward your goals.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
55
v0.37.0
Latest
Jun 30, 2026
+ 54 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
