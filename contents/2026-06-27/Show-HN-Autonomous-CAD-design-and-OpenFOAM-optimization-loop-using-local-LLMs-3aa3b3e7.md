---
source: "https://github.com/ostenjap/LLM-Agent-generated-Quadcopter-Prop"
hn_url: "https://news.ycombinator.com/item?id=48701947"
title: "Show HN: Autonomous CAD design and OpenFOAM optimization loop using local LLMs"
article_title: "GitHub - ostenjap/LLM-Agent-generated-Quadcopter-Prop: Autonomous multi-agent parametric design + CFD optimization using free local LLMs (Ollama), CadQuery, and OpenFOAM. $0 API cost. · GitHub"
author: "ostenjap"
captured_at: "2026-06-27T22:22:05Z"
capture_tool: "hn-digest"
hn_id: 48701947
score: 1
comments: 0
posted_at: "2026-06-27T21:28:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Autonomous CAD design and OpenFOAM optimization loop using local LLMs

- HN: [48701947](https://news.ycombinator.com/item?id=48701947)
- Source: [github.com](https://github.com/ostenjap/LLM-Agent-generated-Quadcopter-Prop)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T21:28:55Z

## Translation

Title: Show HN: Autonomous CAD design and OpenFOAM optimization loop using local LLMs
記事タイトル: GitHub - ostenjap/LLM-Agent-generated-Quadcopter-Prop: 無料のローカル LLM (Ollama)、CadQuery、OpenFOAM を使用した自律的なマルチエージェント パラメトリック設計 + CFD 最適化。 API コストは 0 ドル。 · GitHub
Description: Autonomous multi-agent parametric design + CFD optimization using free local LLMs (Ollama), CadQuery, and OpenFOAM. API コストは 0 ドル。 - ostenjap/LLM-Agent-generated-Quadcopter-Prop
HN text: LLM designs Quadcopter propeller and verifies with physics. It uses the Auto research Loop to find the best design.

記事本文:
GitHub - ostenjap/LLM-Agent-generated-Quadcopter-Prop: 無料のローカル LLM (Ollama)、CadQuery、OpenFOAM を使用した自律型マルチエージェント パラメトリック設計 + CFD 最適化。 API コストは 0 ドル。 · GitHub
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
オステンジャップ
/
LLM エージェント生成のクアッドコプター プロップ

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ostenjap/LLM-Agent-generated-Quadcopter-Prop
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット cad cad cfd cfd data data docs docs skill/ run-propeller-optimizer skill/ run-propeller-optimizer src src .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.mdimplementation_plan.mdimplementation_plan.mdrequirements.txt requirements.txt run_overnight.ps1 run_overnight.ps1 watch.py watch.py View all files Repository files navigation
AI チームによって設計されたドローン プロペラ
無料のローカル LLM (Ollama)、CadQuery、OpenFOAM CFD を利用した、3D 印刷可能なクアッドコプター プロペラのパラメトリック設計、シミュレーション、多目的最適化のための自律型マルチエージェント システムです。
This project tries to answer a simple question: instead of an engineer hand-tweaking
a propeller and testing it over and over, can we let a group of AI models do that
loop themselves — and end up with a better propeller than a person would patiently
手で磨く？
The target propeller should be three things at once: quiet , efficient
(it doesn't waste battery), and strong (it pushes a lot of air).それらの目標
fight each other — a bigger, grippier blade gives you more thrust but more noise,
for example — so there's no single "best" answer.お得なセットがあります
トレードオフがあり、仕事はそれを見つけることです。
If you're new to AI agents, this is a nice thing to learn from, because it's not a
チャットボット。 It's AI used as a worker that actually does something and checks its own
結果。
Picture a small research team where every member is a program:
いくつかのジュニア会員は速くて安いです。彼らの仕事はブレインストーミングです - thr

ああ
たくさんのプロペラのデザインがあり、合理的なものもあれば、奇妙なものもあります。そうである必要はない
賢い、彼らは多産でなければなりません。これらは無料のローカルを使用して自分のコンピュータ上で実行されます
モデル (Ollama 経由) なので、ブレインストーミングには費用がかかりません。
主任研究者は、お金がかかり、賢い研究者です。それは生成しません
単調な仕事。結果を読み取り、パターンを認識します (「5 ブレードのデザインは
静かになる—そこをより強くプッシュする」）、そしてチームが次に何を試みるかを決定します。
その役割は、ショーを運営するエージェントである反重力です。
決して嘘をつかない審判: 単純で信頼できる数学と物理学のコード。 AIのみ
デザインを提案することもあります。彼らは自分の作品に採点することを決して許可されません。の
採点は審判が行うため、自信はあっても間違ったモデルがシステムを騙すことはできません。
この最後の点が重要なポイントであり、一般的な教訓として覚えておく価値があります。
AI について: モデルに提案させ、信頼できるコードに決定させます。モデルが得意とするのは、
選択肢を思いつくが、信頼できる真実の判断者になるのが苦手。そこで私たちは使用します
彼らは自分が得意な部分だけを取り上げます。
チームはループで作業します。 1周はこんな感じです。
提案 — 安価なモデルは、新しいプロペラ設計を大量に吐き出します。
ビルド — コードは各設計を実際の 3D ジオメトリ (実際の CAD ファイル) に変換します。
スコア — 物理コードは、推力の量、騒音の量、およびその方法を推定します。
それぞれが効率的です。この最初のパスは高速な近似計算であり、完全な計算ではありません。
シミュレーション。
候補リスト — システムは、すべての目標で達成されなかったデザインを維持します。
一度。その生き残ったセットはパレート フロントと呼ばれます。これが現在の「メニュー」です。
最良のトレードオフ。
反映 — 反重力は前方を見て、何が機能しているのか、何をすべきなのかを判断します。
次に探索します。
書き留めます。ラウンドの 1 行の概要がログに記録され、ループが実行されます。
また始まります。
最も有望なデザインが高価な治療を受けることが時々あります

t: 本物
空気をモデル化する流体力学シミュレーション ( CFD 、OpenFOAM を使用してローカルで実行)
実際にブレードの上を流れます。それは遅いので、候補者にのみ費やします
安易な計算ではすでにうまくいっているように見えます。
もう 1 つ名前を付ける価値のあるヘルパーがあります。サロゲート モデル (ガウス過程、
scikit-learn より)。チームがシミュレーションの推測を学習していると考えてください。
すでにシミュレーションした設計から答えを得ることができるため、多くの時間をスキップできます。
本当に不確実な場合にのみ実行し、それらを費やします。それはシステムが得ているものです
どこに目を向ければよいのか賢くなります。
仕組み (エージェント + 物理学)
これを信頼できるものにする 1 つのルール: AI エージェントは提案のみを行う
デザインは決して採点されません。スコアリングは単純な物理コードによって行われます。
間違った答えを導き出すことはできません。ループ全体は次のとおりです。
フローチャート TD
OBJ[「反重力は目標を設定します<br/>静か、効率的、高推力」]
サブグラフ提案["PROPOSE — エージェントは提案のみ (スコアは付けない)"]
P["提案者<br/>ローカルLLM"]
M["ミューテーター<br/>ローカルLLM"]
C["コーダーは<br/>検索演算子を作成します"]
GA[「決定的 GA<br/>常に実行され、進捗が保証される」]
終わり
CAND["候補デザイン<br/>(各パラメータ 7 つ)"]
サブグラフ スコア["SCORE — 信頼できる物理学、LLM (グラウンド トゥルース) なし"]
PERF["パフォーマンス.py<br/>BEMT ホバリング → 推力、性能指数"]
TUB["tubercle_analysis.py<br/>→ ノイズリダクション (dB)"]
STR["propeller_physics.py<br/>→ 応力、共鳴"]
V["evaluate.py<br/>目標 + 制約 → 実現可能?"]
パフォーマンス --> V
タブ --> V
STR --> V
終わり
SEL["pareto.py<br/>非支配的なデザインを維持"]
DB[("SQLite Research.db<br/>すべてのデザイン + スコア")]
REF[「反重力が反映し<br/>次世代を導く」]
CAD["generate_propeller.py<br/>STEP / STL + 防水チェック"]
CFD["cfd_verify.py<br/>OpenFOAM 真実チェック"]
オブジェクト -->

P&M&C&GA
P&M&C&GA --> CAND
CAND --> PERF & TUB & STR
V --> 選択
SEL --> DB --> REF
REF -->|次世代| P
SEL -->|最高のデザイン| CAD --> CFD
読み込み中
左から右、上から下に読んでください: エージェント (および決定論的遺伝学)
常にセーフティネットとして動作するアルゴリズム）候補設計を破棄 →
物理スクリプトがそれぞれにスコアを付けます → 優勢でない勝者は保持され、保存されます →
反重力は勝者を見て次のラウンドに進みます → ループが繰り返されます。
最高の設計は、最終的には CAD と CFD の検証に落ちてしまいます。
プロンプト システム — 7B モデルが実際の CAD コードを記述する方法
これは、あなた自身のプロジェクトに最も移し替えやすい部分です。すべてのプロンプトは
読みやすい平易な英語で、任意のドメインのパターンをコピーできます。
小規模なローカル モデルで構造化された検証済みの出力を生成したい場合。
7B モデルを作成して、機能する CAD コードと有効な設計を確実に生成する秘訣
パラメータは、役割の分離 + サンドボックス実行 + 自己修正です。それぞれ
ワーカーは、厳密な出力スキーマを持つ単一の制約されたジョブを取得します。
労働者
モデル
何をするのか
出力
提案者
qwen2.5-coder:7b
新しいデザインをゼロからブレインストーミングする
7 パラメーターの設計ベクトルの JSON 配列
ミューテーター
qwen2.5-coder:7b
優れたデザインを採用し、小さなバリエーションを作成します
調整されたベクトルの JSON 配列
コーダー
qwen2.5-coder:7b
Python 検索演算子 (突然変異関数) を作成します。
JSON {"コード": "..."}
CFDアナリスト
ファイ4ミニ
OpenFOAM ログを読み取り、ソルバーの障害を診断します
JSON {"ステータス": "...", "修正": "..."}
筆記者
ファイ4ミニ
一行の日記エントリを書きます
プレーンテキスト
Proposer プロンプトの外観 (実際のファイル)
これは src/autoresearch/skills/proposer.md です。7B モデルの完全なプロンプトです。
受け取ります。注意: 曖昧な指示はなく、厳密な境界とドメインの知識だけです。
あなたはプロペルです

自動化された研究群の R DESIGN PROPOSER。
あなたの仕事: ホバリングを改善する可能性のある新しいプロペラ設計の候補を提案する
効率を高めたり、結核騒音の低減を高めたり、ブレードの質量を減らしたりします。
JSON のみを出力します。散文も値下げもありません。スキーマは次のとおりです。
{"デザイン": [
{"chord_root_m": <float>, "chord_tip_m": <float>,
"twist_root_deg": <float>、"twist_tip_deg": <float>、
"tubercle_amp_m": <float>、"tubercle_wl_m": <float>、
"n_blades": <int>},
...
]}
ハード境界 (これらの境界内に留まります。外側の値はクランプされます):
コードルート_m : 0.020 .. 0.034
コードチップ_m : 0.006 .. 0.014
ツイストルート角度: 25 .. 45
Twist_tip_deg : 6 .. 20
tubercle_amp_m : 0.0 .. 0.005
tubercle_wl_m : 0.020 .. 0.060
n_blades : 2 .. 6 (整数)
コーダーのコードがどのようにサンドボックス化されるか
Coder は任意の Python を作成しますが、これは危険です。の
サンドボックスはこれを 2 つのレイヤーで処理します。
AST ホワイトリスト — 実行前に、AST ウォーカーは外部からのインポートを拒否します。
{math、numpy、random} 、あらゆるダンダー アクセス、およびあらゆる危険な組み込み ( open 、
exec 、 eval 、 os 、 サブプロセス など)
サブプロセスの分離 — スクリーンされたコードは、新しい Python プロセスで実行されます。
ハードタイムアウトとスクラッチ作業ディレクトリ。標準出力の JSON 行のみが
戻って受け入れられました。
コードが無効、タイムアウト、またはガベージを返した場合は、暗黙的に破棄されます。
最悪のケースは世代スロットが無駄になることであり、アーカイブが破損することはありません。
# サンドボックス コントラクト (sandbox.py より):
def mutate (parents、bounds、rng):
# 親 : 設計ベクトルのリスト
# 境界 : 各変数の [lo, hi] のリスト
# rng : random.Random インスタンス (再現性のため)
# 戻り値 : 新しい設計ベクトルのリスト
...
# STRICT サンドボックス ルール (違反 → 演算子は破棄):
# - インポートのみ: math、numpy、random。他には何もありません。
# - ファイル/ネットワーク/システムへのアクセス、オープン/実行/評価はありません。
# - 以内に返却する必要があります

5秒。
# - すべての値は [lo, hi] 境界にクランプされます。
自己修正ループ
メッシュの不良、CFD ソルバーの発散、JSON の不正など、何かが失敗した場合、
エラーは、診断コンテキストとともに責任のある作業者にフィードバックされます。の
たとえば、CFD アナリストはソルバー ログの末尾と残差値を取得します。
次に試行する具体的な修正を 1 つだけ返す必要があります。
{ "ステータス" : " 分岐中 " ,
"diagnosis" : " 反復 200 回後に U 残差が上昇、クーラント違反の可能性があります " ,
"fix" : " deltaT を 1e-3 から 5e-4 に減らします " ,
"フィールド" : { "デルタT" : " 5e-4 " }}
これはパターンです: 構造化出力 → 検証 → 自動再試行 。それは機能します
なぜなら、モデルが最初の試行で正しくなる必要は決してなく、ただ正しくなければならないからです。
最終的には、再試行の予算内で。
📂 6 つのワーカー プロンプトはすべて src/autoresearch/skills/ にあります。
— 短くて自己完結型なので、直接読んでください。
OpenFOAM、CadQuery、またはパラメトリック設計ツールを使用して作業する場合、このプロジェクトは
これは、デザイン、シミュレーション、最適化のループを自動化するための実用的なリファレンスでもあります。
再利用または学習できる内部の内容は次のとおりです。
すべてのプロペラは 7 つのパラメータ (根元と先端の弦、ねじれ) によって定義されます。
分布、結節の振幅と波長、葉の数。の
generate_propeller.py スクリプト

[切り捨てられた]

## Original Extract

Autonomous multi-agent parametric design + CFD optimization using free local LLMs (Ollama), CadQuery, and OpenFOAM. $0 API cost. - ostenjap/LLM-Agent-generated-Quadcopter-Prop

LLM designs Quadcopter propeller and verifies with physics. It uses the Auto research Loop to find the best design.

GitHub - ostenjap/LLM-Agent-generated-Quadcopter-Prop: Autonomous multi-agent parametric design + CFD optimization using free local LLMs (Ollama), CadQuery, and OpenFOAM. $0 API cost. · GitHub
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
ostenjap
/
LLM-Agent-generated-Quadcopter-Prop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ostenjap/LLM-Agent-generated-Quadcopter-Prop
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits cad cad cfd cfd data data docs docs skills/ run-propeller-optimizer skills/ run-propeller-optimizer src src .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md implementation_plan.md implementation_plan.md requirements.txt requirements.txt run_overnight.ps1 run_overnight.ps1 watch.py watch.py View all files Repository files navigation
A drone propeller, designed by a team of AIs
An autonomous multi-agent system for parametric design, simulation, and multi-objective optimization of 3D-printable quadcopter propellers — powered by free, local LLMs (Ollama), CadQuery, and OpenFOAM CFD.
This project tries to answer a simple question: instead of an engineer hand-tweaking
a propeller and testing it over and over, can we let a group of AI models do that
loop themselves — and end up with a better propeller than a person would patiently
grind out by hand?
The target propeller should be three things at once: quiet , efficient
(it doesn't waste battery), and strong (it pushes a lot of air). Those goals
fight each other — a bigger, grippier blade gives you more thrust but more noise,
for example — so there's no single "best" answer. There's a set of good
trade-offs, and the job is to find them.
If you're new to AI agents, this is a nice thing to learn from, because it's not a
chatbot. It's AI used as a worker that actually does something and checks its own
results.
Picture a small research team where every member is a program:
A few junior members are fast and cheap. Their job is to brainstorm — throw
out lots of propeller designs, some sensible, some weird. They don't need to be
smart, they need to be prolific. These run on your own computer with free local
models (via Ollama), so brainstorming costs nothing.
A lead researcher is the expensive, smart one. It doesn't generate the
grunt work; it reads the results, notices patterns ("the 5-blade designs are
getting quieter — push harder there"), and decides what the team tries next.
That role is Antigravity , the agent running the show.
A referee that never lies: plain, trusted math and physics code. The AIs only
ever propose designs. They're never allowed to score their own work. The
referee does the scoring, so a confident-but-wrong model can't fool the system.
That last point is the whole trick, and it's worth remembering as a general lesson
about AI: let models suggest, let trusted code decide. Models are great at
coming up with options and terrible at being a reliable judge of truth. So we use
them only for the part they're good at.
The team works in a loop. One lap looks like this:
Propose — the cheap models spit out a batch of new propeller designs.
Build — code turns each design into actual 3D geometry (a real CAD file).
Score — the physics code estimates how much thrust, how much noise, and how
efficient each one is. This first pass is fast approximate math, not a full
simulation.
Shortlist — the system keeps the designs that aren't beaten on every goal at
once. That surviving set is called the Pareto front — the current "menu" of
best trade-offs.
Reflect — Antigravity looks at the front and steers: what's working, what to
explore next.
Write it down — a one-line summary of the round gets logged, and the loop
starts again.
Every so often, the most promising designs get the expensive treatment: a real
fluid-dynamics simulation ( CFD , run locally with OpenFOAM) that models the air
actually flowing over the blade. That's slow, so we only spend it on candidates
that already look good on the cheap math.
There's one more helper worth naming: a surrogate model (a Gaussian Process,
from scikit-learn). Think of it as the team learning to guess the simulation's
answer from the designs it has already simulated — so it can skip a lot of slow
runs and spend them only where it's genuinely unsure. It's the system getting
smarter about where to look as it goes.
How it works (agents + physics)
The one rule that makes this trustworthy: the AI agents only ever propose
designs — they never score them. Scoring is done by plain physics code that
can't be talked into a wrong answer. Here's the whole loop:
flowchart TD
OBJ["Antigravity sets the objective<br/>quiet · efficient · high thrust"]
subgraph propose["PROPOSE — agents only suggest (never score)"]
P["Proposer<br/>local LLM"]
M["Mutator<br/>local LLM"]
C["Coder<br/>writes a search operator"]
GA["Deterministic GA<br/>always runs, guarantees progress"]
end
CAND["candidate designs<br/>(7 parameters each)"]
subgraph score["SCORE — trusted physics, no LLM (the ground truth)"]
PERF["performance.py<br/>BEMT hover → thrust, Figure of Merit"]
TUB["tubercle_analysis.py<br/>→ noise reduction (dB)"]
STR["propeller_physics.py<br/>→ stress, resonance"]
V["evaluate.py<br/>objectives + constraints → feasible?"]
PERF --> V
TUB --> V
STR --> V
end
SEL["pareto.py<br/>keep the non-dominated designs"]
DB[("SQLite research.db<br/>every design + score")]
REF["Antigravity reflects<br/>steers the next generation"]
CAD["generate_propeller.py<br/>STEP / STL + watertight check"]
CFD["cfd_verify.py<br/>OpenFOAM truth check"]
OBJ --> P & M & C & GA
P & M & C & GA --> CAND
CAND --> PERF & TUB & STR
V --> SEL
SEL --> DB --> REF
REF -->|next generation| P
SEL -->|best designs| CAD --> CFD
Loading
Read it left-to-right, top-to-bottom: the agents (and a deterministic genetic
algorithm that always runs as a safety net) throw out candidate designs → the
physics scripts score each one → the non-dominated winners are kept and saved →
Antigravity looks at the winners and steers the next round → the loop repeats.
The best designs eventually drop out the bottom into CAD and CFD verification.
The prompt system — how 7B models write real CAD code
This is the part most transferable to your own projects. Every prompt is
readable plain English, and you can copy the pattern for any domain
where you want a small local model to generate structured, validated output.
The secret to making a 7B model reliably produce working CAD code and valid design
parameters is role separation + sandboxed execution + self-correction . Each
worker gets a single, constrained job with a strict output schema.
Worker
Model
What it does
Output
Proposer
qwen2.5-coder:7b
Brainstorms brand-new designs from scratch
JSON array of 7-parameter design vectors
Mutator
qwen2.5-coder:7b
Takes a good design and creates small variations
JSON array of tweaked vectors
Coder
qwen2.5-coder:7b
Writes a Python search operator (mutation function)
JSON {"code": "..."}
CFD Analyst
phi4-mini
Reads OpenFOAM logs and diagnoses solver failures
JSON {"status": "...", "fix": "..."}
Scribe
phi4-mini
Writes one-line journal entries
Plain text
How the Proposer prompt looks (actual file)
This is src/autoresearch/skills/proposer.md — the full prompt that a 7B model
receives. Notice: no vague instructions, just hard bounds and domain knowledge:
You are a PROPELLER DESIGN PROPOSER in an automated research swarm.
Your job: propose NEW candidate propeller designs that might improve hover
efficiency, increase tubercle noise reduction, or reduce blade mass.
You output ONLY JSON. No prose, no markdown. The schema is:
{"designs": [
{"chord_root_m": <float>, "chord_tip_m": <float>,
"twist_root_deg": <float>, "twist_tip_deg": <float>,
"tubercle_amp_m": <float>, "tubercle_wl_m": <float>,
"n_blades": <int>},
...
]}
Hard bounds (stay inside these; values outside are clamped):
chord_root_m : 0.020 .. 0.034
chord_tip_m : 0.006 .. 0.014
twist_root_deg : 25 .. 45
twist_tip_deg : 6 .. 20
tubercle_amp_m : 0.0 .. 0.005
tubercle_wl_m : 0.020 .. 0.060
n_blades : 2 .. 6 (integer)
How the Coder's code gets sandboxed
The Coder writes arbitrary Python, which is dangerous. The
sandbox handles it in two layers:
AST allowlist — before execution, an AST walker rejects any import outside
{math, numpy, random} , any dunder access, and any dangerous builtin ( open ,
exec , eval , os , subprocess , etc.)
Subprocess isolation — the screened code runs in a fresh Python process with
a hard timeout and a scratch working directory. Only a JSON line on stdout is
accepted back.
If the code is invalid, times out, or returns garbage, it's silently discarded —
worst case is a wasted generation slot, never a corrupted archive:
# The sandbox contract (from sandbox.py):
def mutate ( parents , bounds , rng ):
# parents : list of design vectors
# bounds : list of [lo, hi] for each variable
# rng : random.Random instance (for reproducibility)
# returns : list of NEW design vectors
...
# STRICT sandbox rules (violations → operator discarded):
# - Import ONLY: math, numpy, random. Nothing else.
# - No file/network/system access, no open/exec/eval.
# - Must return within 5 seconds.
# - Every value clamped into [lo, hi] bounds.
The self-correction loop
When something fails — a bad mesh, a diverging CFD solver, malformed JSON — the
error is fed back to the responsible worker with the diagnostic context. The
CFD Analyst, for example, gets the tail of the solver log and the residual values,
and must return exactly one concrete fix to try next:
{ "status" : " diverging " ,
"diagnosis" : " U residuals climbing after iteration 200, likely Courant violation " ,
"fix" : " reduce deltaT from 1e-3 to 5e-4 " ,
"fields" : { "deltaT" : " 5e-4 " }}
This is the pattern: structured output → validation → auto-retry . It works
because the model never has to be right on the first try — it just has to be right
eventually , within a budget of retries.
📂 All six worker prompts are in src/autoresearch/skills/
— read them directly, they're short and self-contained.
If you work with OpenFOAM, CadQuery, or parametric design tools, this project
is also a working reference for automating the design-simulate-optimize loop.
Here's what's under the hood that you can reuse or learn from:
Every propeller is defined by 7 parameters — chord at root and tip, twist
distribution, tubercle amplitude and wavelength, and blade count. The
generate_propeller.py script

[truncated]
