---
source: "https://github.com/piyushkhemka/sip"
hn_url: "https://news.ycombinator.com/item?id=48878407"
title: "Built a tracker to estimate water wastage when talking to Claude"
article_title: "GitHub - piyushkhemka/sip: Every prompt takes a sip 💧 — a fun Claude Code status line showing how much water talking to Claude costs, plus cost, 5h/weekly usage (mirrors /usage), context, model & effort, and turn count. All local, all free. · GitHub"
author: "piyushkhemka"
captured_at: "2026-07-12T05:17:50Z"
capture_tool: "hn-digest"
hn_id: 48878407
score: 1
comments: 0
posted_at: "2026-07-12T05:01:34Z"
tags:
  - hacker-news
  - translated
---

# Built a tracker to estimate water wastage when talking to Claude

- HN: [48878407](https://news.ycombinator.com/item?id=48878407)
- Source: [github.com](https://github.com/piyushkhemka/sip)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T05:01:34Z

## Translation

タイトル: クロードと話すときに水の浪費を推定するトラッカーを構築しました
記事のタイトル: GitHub - piyushkhemka/sip: すべてのプロンプトが一口飲みます 💧 — クロードと話すのにかかる水の量、プラスコスト、週 5 時間の使用量 (ミラー / 使用量)、コンテキスト、モデルと労力、およびターン数を示す楽しいクロード コードのステータス ライン。すべてローカル、すべて無料。 · GitHub
説明: すべてのプロンプトを一口飲みます 💧 — クロードと話すのにかかる水の量、プラスコスト、週 5 時間の使用量 (ミラー / 使用量)、コンテキスト、モデルと労力、ターン数を示す楽しいクロード コードのステータス ライン。すべてローカル、すべて無料。 - ピユシュケムカ/一口

記事本文:
GitHub - piyushkhemka/sip: すべてのプロンプトを一口飲みます 💧 — クロードと話すのにかかる水の量、プラスコスト、週 5 時間の使用量 (ミラー / 使用量)、コンテキスト、モデルと労力、ターン数を示す楽しいクロード コードのステータス ライン。すべてローカル、すべて無料。 · GitHub
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
ピユシケムカ
/
一口飲む
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .claude-plugin .claude-plugin コマンド コマンド スクリプト scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md settings.json settings.json sip.sh sip.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Sip は、水の量はどれくらいかという 1 つの質問を中心に構築された楽しい小さなプロジェクトです。
実際にクロードと話す必要があるのですか？すべてのプロンプトは文字通り一口飲みます。
LLM のトレーニングと実行には実際の水を消費します (データセンターの冷却と
発電）、そのコストは通常目に見えません。シップはクロードです
すでにいる場所に番号を付けるコードのステータス行
作業しながらリアルタイムで検索、更新できます。
💧今日は1.2mL・0.5L
それが見出しです — この回答の推定水量と今日の実行結果
合計。 Sip には、良い製品に必要な便利なものも表示されます。
ステータス行 — セッション/日次コスト、5 時間および週次プランの使用状況 (
絵文字を使用すると、限界にどれだけ近づいているかが一目でわかります)、コンテキスト ウィンドウ
使用状況、モデル + 努力値、ターン数。すべてローカル、すべて無料、追加の API なし
呼び出し:
💵 $0.42 (今日 $3.17) | ⚠️ 5 時間 45% (2 時間 14 分でリセット) · ✅ 7 日 12% (3 日でリセット) | 💧 今日は1.2 mL · 0.5 L | 🧠 50k/200k (25%) | 🤖 作品 4.8 (多大な努力) | 🔄 ターン: 14
セグメント
番組
💵 コスト
セッションコスト + 今日の累計
✅ ⚠️🚨🚫使用法
5 時間 / 7 日プランの使用、ステータス絵文字 + リセット時間付き
💧 水
この回答の推定水量・今日の合計
🧠 コンテキスト
使用されたトークン / ウィンドウ サイズ、色分け
🤖 モデルと努力
短いモデル名 + 推論の労力
🔄ターン
実際のユーザーのターン

このセッション
使用状況: ✅ <33% · ⚠️ 33 ～ 66% · 🚨 67 ～ 84% · 🚫 ≥85% (同)
しきい値はコンテキスト セグメントに色を付けます)。使用状況セグメントは 1 回だけ表示されます
アカウントにはレート制限データがあります。それ以外の場合は単純に省略され、表示されません
プレースホルダー付き。
コンテキスト %: Sip は同じ生データからパーセンテージ自体を計算します。
使用済み/合計 (入力 + キャッシュ、モデルのコンテキスト全体) で表示されるトークン数
ウィンドウ — 200k、または拡張コンテキスト モデルでは 1M) なので、この 2 つは決してできません
視覚的に同意しません。クロード・コード自身にのみフォールバックします。
context_window.used_percentage (ウィンドウ サイズがペイロードにない場合)。
これは、キャッシュを除外したり、
分母が違う。正確な情報をキャプチャするには、SIP_DEBUG=/tmp/sip.log を設定します。
ペイロード クロード コードが送信し、生の数値を確認します。
jq (macOS では brew install jq、Debian/Ubuntu では apt-get install jq) および
awk (macOS/Linux ではデフォルトで存在します)。
2 つのステップ: プラグインをインストールし、セットアップ コマンドを実行します。プラグイン
自身をサブエージェント パネルに自動的に接続できますが、(
コマンド) は、Claude Code のメイン ステータス ラインをそれに指すことができます。
1 回限りのクロード コード レベルの設定。
クロード プラグイン マーケットプレイスに piyushkhemka/sip を追加
クロードプラグインのインストール SIP
またはローカルチェックアウトから:
クロード プラグイン マーケットプレイス /path/to/sip を追加
クロードプラグインのインストール SIP
これにより、Sip のサブエージェント ステータス ラインが即座に接続されます (エージェントに表示されます)
パネル）。メインのステータス行 (画面の下部にあるもの) は、
まだ設定されていません。それがステップ2です。
ステップ 2 — メインステータスラインをオンにする
/sip-セットアップ
これにより、sip.sh の安定した自己完結型コピーが ~/.claude/sip.sh にインストールされます。
~/.claude/settings.json をそこにポイントします (現在のファイルをバックアップします)
まず)。インストールされたコピーは場所に依存しないため、Sip は動作し続けます。
私でも

プラグインのバージョン管理されたキャッシュ ディレクトリが後の更新で変更される場合 —
プラグインを更新した後に /sip-setup を再実行してコピーを更新するだけです。
Claude Code をリロードすると、画面の下部にステータス ラインが表示されます。
Sip は完全にローカルで実行され、API トークンを使用しません。何も表示されません。
使用量や水道料金など、何らかの費用がかかるか、プランに影響します。
残りのコストはレイテンシーであるため、スクリプトはそれを最小限に抑えるように構築されています。
セッション記録を 1 回の awk パスで実際のユーザーのターンと合計をカウントします
出力トークン (水の推定用) — ペイロードにはターン番号が含まれていません。
したがって、これにより正確さと冪等性の両方が維持されます。次の場合は完全にスキップされます。
最後のレンダリング以降、トランスクリプトは増加していません（セッションごとにバイト単位でキャッシュされます）
これは、新しいメッセージが表示されずに再レンダリングが繰り返される場合によく見られるケースです。
1 つの jq 呼び出しでペイロードが解析され、コンテキスト/使用率が導出され、
日次コスト + 水道帳簿を作成します (状態は、経由で同じ呼び出しに読み込まれます)
--slurpfile )、すべてのフィールドが防御的に強制されるため、不正な形式の事前
レコードによって実行が破損することはありません。
1 つの小さなアトミックな書き込みにより、単一の daily.json が永続化されます。
ローカルの日が変わる - 日ごとのファイルもクリーンアップもありません。軽量
mkdir ベースのロック (競合がない場合はコストがほぼゼロ) により、
同じファイル上で競合するクロード コードの同時セッション。
macOS に同梱されている bash 3.2 に移植可能。 jq と awk のみが必要です。
状態は ${CLAUDE_PLUGIN_DATA:-${XDG_STATE_HOME:-$HOME/.claude}/sip}/daily.json にあります。
セッションごとの小さなトランスクリプト スキャン キャッシュと並行して。
💧 水の見積もりの仕組み
Sip は、応答にかかる目に見えないコストを数値に変換し、ローカルで推定します。
出力トークンから (ネットワーク呼び出しなし、API トークンはまだゼロ):
Water_mL = Output_tokens / 1000 × K (デフォルト K = 1,000 出力トークンあたり 1.5 mL)

生成 (デコード) が推論エネルギーを支配するため、出力トークンがそれを駆動します。
input/prefill はトークンごとにはるかに安価です。デフォルトの K = 1.5 mL / 1k 出力
トークンは、透過的に導出される総設置面積 (冷却 + 電力) です。
エネルギー: 1,000 出力トークンあたり ~0.3 ～ 0.6 Wh — Google の値と一致
Gemini テキスト プロンプトと GPU の中央値のファーストパーティ値は 0.24 Wh
デコードのスループット。
水強度: オンサイトで ~1.1 mL/Wh (Google の暗黙の WUE、PUE 1.09) プラス
米国の送電網による発電量は ~2 mL/Wh ≈ 合計 ~3 mL/Wh。
⇒ ~0.4 ～ 0.6 Wh × ~3 mL/Wh ≈ 1,000 出力トークンあたり ~1.5 mL。
正気度: 典型的な反応は ≈ 1 mL、長い反応は数 mL、重い日は ~100 ～ 200
mL。これは、Google の 0.26 mL (オンサイトのみ) と広く共有されている 0.26 mL の間に位置します。
(そして今では誤りが証明されています) ~500 mL の「クエリあたりのボトル」という数字は、かなり大きな数字でした。
古い GPT-3 の最悪のケース。
これは推定値であり、実際の範囲は広い (~100 倍)。あなたの好みに合わせて調整してください
自分自身の仮定:
import SIP_WATER_ML_PER_1K_TOKENS=0.5 # 保守的、オンサイトのみ (Google スコープ)
import SIP_WATER_ML_PER_1K_TOKENS=1.5 # デフォルト、合計フットプリント
なぜコストから導き出さないのでしょうか？価格にはマージンが組み込まれており、モデル層によって異なりますので、
これはトークンよりもエネルギーのうるさいプロキシです。Sip はトークンを使用します。
出典: Google Cloud — AI 推論の環境への影響の測定 (2025 年 8 月)
· Li et al.、「Making AI Lessty」(arXiv 2304.03271)
· Goedecke、「ChatGPT と話すと、500 ml ではなく 5 ml の水がかかります。」
5 時間と 7 日 (毎週) のパーセンテージは、同じ数値/使用量です。
表示します — クロード コードには、すべてのステータス ライン レンダリングにすでにそれらが含まれています
( rate_limits.five_hour / rate_limits.seven_day )。 Sip はただ読むだけです
ローカルで印刷してください。余分な API 呼び出しは行わず、何も追加しません。
使用量 — 使用量を確認しても、使用量はかかりません。
どちらか

ndow は独立して存在しない場合があります (例: Five_hour のみが設定されます)
まだ); Sip は、存在するウィンドウを表示し、残りを省略します。
~/.claude/settings.json から statusLine キーを削除します (または、
settings.json.bak.* Enable.sh によって作成されたバックアップ)。
すべてのプロンプトを一口飲みます 💧 — クロードと話すのにかかる水の量、プラスコスト、週 5 時間の使用量 (ミラー / 使用量)、コンテキスト、モデルと労力、ターン数を示す楽しいクロード コードのステータス ライン。すべてローカル、すべて無料。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Every prompt takes a sip 💧 — a fun Claude Code status line showing how much water talking to Claude costs, plus cost, 5h/weekly usage (mirrors /usage), context, model & effort, and turn count. All local, all free. - piyushkhemka/sip

GitHub - piyushkhemka/sip: Every prompt takes a sip 💧 — a fun Claude Code status line showing how much water talking to Claude costs, plus cost, 5h/weekly usage (mirrors /usage), context, model & effort, and turn count. All local, all free. · GitHub
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
piyushkhemka
/
sip
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .claude-plugin .claude-plugin commands commands scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md settings.json settings.json sip.sh sip.sh View all files Repository files navigation
Sip is a fun little project built around one question: how much water does
it actually take to talk to Claude? Every prompt takes a sip — literally.
Training and running LLMs consumes real water (datacenter cooling +
electricity generation), and that cost is normally invisible. Sip is a Claude
Code status line that puts a number on it, right where you're already
looking, updating live as you work.
💧 1.2 mL · 0.5 L today
That's the headline — this response's estimated water, then today's running
total. Sip also shows the useful stuff you'd want from a good
status line — session/daily cost, your 5-hour and weekly plan usage (with an
emoji so you know at a glance how close you are to a limit), context window
usage, model + effort, and your turn count. All local, all free, no extra API
calls:
💵 $0.42 ($3.17 today) | ⚠️ 5h 45% (resets in 2h14m) · ✅ 7d 12% (resets in 3d) | 💧 1.2 mL · 0.5 L today | 🧠 50k/200k (25%) | 🤖 opus 4.8 (high effort) | 🔄 Turn: 14
Segment
Shows
💵 Cost
session cost + today's running total
✅ ⚠️ 🚨🚫 Usage
5h / 7d plan usage, with a status emoji + reset time
💧 Water
estimated water this response · today's total
🧠 Context
tokens used / window size, color-coded
🤖 Model & effort
short model name + reasoning effort
🔄 Turn
real user turns this session
Usage status: ✅ <33% · ⚠️ 33–66% · 🚨 67–84% · 🚫 ≥85% (same
thresholds color the context segment). The usage segment only appears once
your account has rate-limit data — it's simply omitted otherwise, not shown
with placeholders.
On the context %: Sip computes the percentage itself from the same raw
token counts shown in used/total (input + cache, over the model's context
window — 200k, or 1M on extended-context models) so the two can never
visually disagree; it only falls back to Claude Code's own
context_window.used_percentage if the window size isn't in the payload.
This can still differ from other surfaces that exclude cache or use a
different denominator. Set SIP_DEBUG=/tmp/sip.log to capture the exact
payload Claude Code sends and see the raw numbers.
jq ( brew install jq on macOS, apt-get install jq on Debian/Ubuntu) and
awk (present by default on macOS/Linux).
Two steps: install the plugin , then run the setup command . A plugin
can wire itself into the subagent panel automatically, but only you (via a
command) can point Claude Code's main status line at it — that's a
one-time, Claude-Code-level setting.
claude plugin marketplace add piyushkhemka/sip
claude plugin install sip
Or from a local checkout:
claude plugin marketplace add /path/to/sip
claude plugin install sip
This immediately wires up Sip's subagent status line (visible in the agent
panel). Your main status line — the one at the bottom of the screen — isn't
set yet; that's step 2.
Step 2 — Turn on the main status line
/sip-setup
This installs a stable, self-contained copy of sip.sh to ~/.claude/sip.sh
and points your ~/.claude/settings.json at it (backing up the current file
first). Because the installed copy is location-independent, Sip keeps working
even if the plugin's versioned cache directory changes on a later update —
just re-run /sip-setup after updating the plugin to refresh the copy.
Reload Claude Code and you'll see the status line at the bottom of the screen.
Sip runs entirely locally and uses zero API tokens — nothing it shows,
including usage and water, costs you anything or counts against your plan.
The remaining cost is latency, so the script is built to minimize it:
One awk pass over the session transcript counts real user turns and sums
output tokens (for the water estimate) — the payload carries no turn number,
so this keeps both accurate and idempotent. It's skipped entirely when the
transcript hasn't grown since the last render (cached per session by byte
size), which is the common case for repeated re-renders with no new messages.
One jq call parses the payload, derives context/usage percentages, and
does the daily cost + water bookkeeping (state read into the same call via
--slurpfile ), with every field defensively coerced so a malformed prior
record can't corrupt the run.
One tiny atomic write persists a single daily.json that self-resets when
the local day changes — no per-day files, no cleanup. A lightweight
mkdir -based lock (near-zero cost when uncontended) protects it from
concurrent Claude Code sessions racing on the same file.
Portable to the bash 3.2 shipped on macOS; only requires jq and awk .
State lives in ${CLAUDE_PLUGIN_DATA:-${XDG_STATE_HOME:-$HOME/.claude}/sip}/daily.json ,
alongside a small per-session transcript-scan cache.
💧 How the water estimate works
Sip turns the invisible cost of a response into a number, estimated locally
from output tokens (no network calls, still zero API tokens):
water_mL = output_tokens / 1000 × K (default K = 1.5 mL per 1,000 output tokens)
Output tokens drive it because generation (decode) dominates inference energy;
input/prefill is far cheaper per token. The default K = 1.5 mL / 1k output
tokens is the total footprint (cooling + electricity), derived transparently:
Energy: ~0.3–0.6 Wh per 1k output tokens — consistent with Google's
first-party figure of 0.24 Wh for a median Gemini text prompt and GPU
decode throughput.
Water intensity: ~1.1 mL/Wh on-site (Google's implied WUE, PUE 1.09) plus
~2 mL/Wh for US grid electricity generation ≈ ~3 mL/Wh total .
⇒ ~0.4–0.6 Wh × ~3 mL/Wh ≈ ~1.5 mL per 1k output tokens.
Sanity: a typical response ≈ 1 mL , a long one a few mL, a heavy day ~100–200
mL. That sits between Google's 0.26 mL (on-site only) and the widely-shared
(and now debunked) ~500 mL "bottle per query" figure, which was a large,
old GPT-3 worst case.
This is an estimate, and the honest range is wide (~100×). Tune it to your
own assumptions:
export SIP_WATER_ML_PER_1K_TOKENS=0.5 # conservative, on-site only (Google-scope)
export SIP_WATER_ML_PER_1K_TOKENS=1.5 # default, total footprint
Why not derive it from cost? Price embeds margin and varies by model tier, so
it's a noisier proxy for energy than tokens — Sip uses tokens.
Sources: Google Cloud — measuring the environmental impact of AI inference (Aug 2025)
· Li et al., "Making AI Less Thirsty" (arXiv 2304.03271)
· Goedecke, "Talking to ChatGPT costs 5ml of water, not 500ml"
The 5-hour and 7-day (weekly) percentages are the same numbers /usage
shows you — Claude Code already includes them in every status-line render
( rate_limits.five_hour / rate_limits.seven_day ). Sip just reads them
locally and prints them; it makes no extra API calls and adds nothing to
your usage — checking your usage doesn't cost you usage.
Either window can be independently absent (e.g. only five_hour populated
yet); Sip shows whichever windows are present and omits the rest.
Remove the statusLine key from ~/.claude/settings.json (or restore a
settings.json.bak.* backup made by enable.sh ).
Every prompt takes a sip 💧 — a fun Claude Code status line showing how much water talking to Claude costs, plus cost, 5h/weekly usage (mirrors /usage), context, model & effort, and turn count. All local, all free.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
