---
source: "https://github.com/bassimeledath/kitchen-rush"
hn_url: "https://news.ycombinator.com/item?id=48551671"
title: "Show HN: Kitchen Rush, Overcooked inspired LLM tool calling benchmark"
article_title: "GitHub - bassimeledath/kitchen-rush: Kitchen Rush: a benchmark for accurate AND fast native tool calling · GitHub"
author: "bombastic311"
captured_at: "2026-06-16T08:05:29Z"
capture_tool: "hn-digest"
hn_id: 48551671
score: 2
comments: 0
posted_at: "2026-06-16T07:10:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kitchen Rush, Overcooked inspired LLM tool calling benchmark

- HN: [48551671](https://news.ycombinator.com/item?id=48551671)
- Source: [github.com](https://github.com/bassimeledath/kitchen-rush)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T07:10:39Z

## Translation

タイトル: Show HN: Kitchen Rush、Overcooked からインスピレーションを得た LLM ツール呼び出しベンチマーク
記事タイトル: GitHub - bassimeledath/kitchen-rush: Kitchen Rush: 正確かつ高速なネイティブ ツール呼び出しのベンチマーク · GitHub
説明: Kitchen Rush: 正確かつ高速なネイティブ ツール呼び出しのベンチマーク - bassimeledath/kitchen-rush

記事本文:
GitHub - bassimeledath/kitchen-rush: Kitchen Rush: 正確かつ高速なネイティブ ツール呼び出しのベンチマーク · GitHub
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
バシメレダス
/
キッチンラッシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
70 コミット 70 コミット .github .github アセット アセット docs docs leaderboard leaderboard scripts scripts src/kitchenrush src/kitchenrush テスト テスト ui ui .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff ライセンス ライセンスREADME.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
インテリジェンスと同じくらい遅延が重要なエージェント ツール呼び出しベンチマーク。
ほとんどのツール呼び出しベンチマーク (BFCL、τ-bench、ToolSandbox、AppWorld) は、モデルが
適切な判断を下し、世界はそれが考えるまで礼儀正しく待ちます。オフラインならそれでいいよ
タスク。しかし、音声アシスタント、ライブオペレーション エージェント、その他リアルタイムのものを構築している場合は、
同時に 2 つのことを考慮します。モデルは正しいことを行うか、そしてそれを高速に行うかです。
十分ですか？ 30 秒の推論の後に完璧な答えを見つけるモデルは、あなたにとっては次のとおりです。
間違ったモデル。
キッチン ラッシュは、構築によって両方を一度に測定します。モデルが考えるのに費やす時間は
アクションが着地するまでに経過するゲーム時間に変換されます。モデルが検討している間、
食べ物は調理され続け、食べ物は焦げ、注文の締め切りは過ぎます。スピードと正確さは両立しない
目を細めて見るグラフ — それらは 1 つのスコアであり、デプロイメントで経験されるのと同じ方法で経験されます。
彼ら。
モデルはオーバークックスタイルでシェフを演じます
キッチン。注文が次々と入ってきます (ハンバーガー、スープ、ラーメンなど)。モデルは通常の注文に応えます。
ネイティブ関数呼び出し - Collect 、 Chop 、 Cook 、 Plate 、 Serve - 締め切りとの競争、
燃焼タイマー、連続成功した料理のコンボボーナス。 ～からの 3 つの意図的な変更
加熱しすぎた場合:
レイテンシーが勝負だ。すべてのモデル応答は、まずその思考時間を

共有
世界時計に合わせてアクションが実行されます。 (1 回の応答で複数の呼び出しを連鎖させて支払うことができます)
待ち時間が 1 回あれば、決断力が報われます。)
ジョイスティックのスキルはありません。シェフは自動的に正しいステーションまで歩きます。旅行
時間はアクションの中にチャージされます。試されているのは正しい行動を選択することだ
ビデオゲームの反射神経ではなく、時間のプレッシャーの下でのシーケンス。
完全に決定的。同じシード、同じアクション、同じレイテンシ → まったく同じエピソード、
いつでも、どのマシンでも。すべての実行はブラウザ ビューアで再生して監査できます。
すべてのエピソードは、KR (キッチン ラッシュ スコア) と呼ばれる単一の 0 ～ 100 のスコアを生成します。それは
2 つの固定アンカー間の曲線上でグレーディング: KR 0 は、「何もしないよりも良くない」ことを意味します。
すべての注文を期限切れにする」、KR 100 は「再生するスクリプト化された参照シェフと一致した」ことを意味します。
同じキッチンで待ち時間ゼロで。」
実際に実行された例がそれを具体的にします。あるキッチンでは、シェフが何もせずに仕上げるとします。
-60 ポイント (すべての注文が期限切れ)、ゼロレイテンシのリファレンス シェフは +140 で終了、
そしてモデルは +40 で終了します。 2 つのアンカーとあなたのアンカーの間には 200 ポイントあります。
モデルはそのうち 100 個をカバーしたため、KR は 50 となり、リファレンスとのギャップの半分が縮まりました。
多くのシードキッチンの平均値を取得すると、リーダーボードのナンバーが得られます
( docs/METHODOLOGY.md に完全な式が記載されています)。
Kitchen Rush を柔軟にするノブは次のとおりです。すべてのキッチンはレイテンシーで生成されます。
予算 B ( --latency-budget 、決定ごとの秒単位)。 B をペースと考えてください。
キッチンの価格は次のとおりです。注文の締め切りは、シェフがちょうど B 秒間費やすように設定されています。
それぞれの決定により、約 1.4 ～ 1.6 倍の余裕があれば、すべての注文を完了することができます。各 B は、
独自のリーダーボード — さまざまな予算での結果が平均されることはありません。
数学的には

裏地付き、価格は正確です:
期限 = 到着 + ⌈σ · C(B)⌉、C(B) = A + K·B
A は注文の本質的な調理/歩行時間、K は有能な計画の意思決定の数です。
はニーズ、σ はヘッドルーム (層ごとに 1.4 ～ 1.6) です。ということで実際はℓ秒で決まるモデル
注文ごとに K・(B − ℓ) 秒間の呼吸空間が得られるか失われます。 Bより速い？銀行の余裕がない
注文の価値がまだ十分にあるうちに提供されます。もっとゆっくり？ヘッドルームを食べ尽くして、
注文は ℓ ≈ B + (σ−1)・C(B)/K あたりで完了不可能になり始めます — での決定あたり約 3 ～ 4 秒
現在の層では B=1。これはまさにキャリブレーション スイープが基準を示している場所です。
シェフの崩壊 ( docs/METHODOLOGY.md §2 、
docs/CALIBRATION.md )。
そして、平易なデプロイメント用語で言えば、B=1 秒で勝つモデルが、すべての場合に最適な選択です。
決定は約 1 秒以内に下される必要があります。ベンチマークの再現可能なクロックでは、
意思決定ごとに約 65 個の出力トークンの予算、つまり簡潔な単発ツールのディスパッチ — なんてことだろう
音声エージェントのニーズ。 B=5s は決定ごとに約 730 トークンを購入します。これは短期間のバーストには十分です。
インタラクティブアシスタントができることの推論。同じモデルでも、ランクが大きく異なる場合があります。
2 つのボードであり、その並べ替えこそがベンチマークの目的です。
17 モデル構成 × 12 シード × {ミディアム、ハード} キッチン × 2 レイテンシ バジェット — 816
これまでのエピソード。各グラフは 1 つのレイテンシ バジェットです。バーは平均 KR、ひげは 95%
信頼区間。階層ごとの完全なテーブル (コスト、推論トークン、およびサービス レートを含む)
Leaderboard/results/board.md にあります。
左側のボード (B=1s) はリアルタイム テストです。キッチンには 1 秒あたり 1 秒の価格が設定されています。
ベンチマークのクロックで約 65 個の出力トークンを購入する決定 — 簡潔な単発ツール
派遣。ここで勝つということは「MOD」を意味します

el 音声エージェントまたはライブ ダッシュボードの操作を信頼します。」
右側のボード (B=5s) では、決定ごとに 5 秒間の同じキッチンの価格が表示されます (~730)
トークン (短時間の推論のための余地)、対話型アシスタントが提供できるもの。
並べて読んでください。そのコントラストが製品です。厳しいリアルタイム圧力下 (B=1 秒)
高速な理由のないモデルが表彰台を獲得します。gemini-3.1-flash-lite はほぼ互角に動作します。
クロード・ソネット-4.6 (32 vs 37)。すべての決定には 5 秒を与え、取締役会に委ねてください。
再注文: gpt-5.4-mini は低い推論でほぼゼロからデッドヒートまでロケットを発射します。
ソネット (44 対 44) のコストは約 5 分の 1 ですが、フラッシュライトは B=1 の半分に下がります。
立っている。推論が完全にオフになっている同じミニのスコアは両方の予算で 0.0 です — 推論
B=1 での can’t ask は、まさに B=5 でのフロンティア レベルのツール呼び出し元になります。それが
遅延税が可視化されました。 ( · 行は少ない労力で推論を実行して実行されたと考えてください。
それ以外の場合は、推論がオフになります。高速なシングルショット ディスパッチが、正直なリアルタイムのデフォルトです。 1行
Anthropic の API があるため、 claude-sonnet-4.6·think はありません。
ツール呼び出しが強制され、ハーネスがツールを強制する場合、拡張思考は許可されません。
呼び出し — ソネットは思考のみを競います。)
ライブで視聴したフリップ: 上部のクリップと同じ 2 つのモデル、
ただし、価格は B=5 秒のキッチンです。今、mini の推理バーストは手頃な価格です - それは終わりです
ソネットはまだ 40 で調理中ですが、すべての注文を 99 生ポイント (KR 86) で提供します。
mini の最高のキッチン — 上のグラフは平均を示しており、24 軒すべてで 44 対 44 の同点です — しかし、
方向性は本物です。B=5 で中層で完全に勝利します (59 対 52)。同じモデルでも違う
レイテンシ バジェット、勝者は異なります。それがまさに 2 つのボードが測定するものです。
2分

utes — スクリプト化された参照シェフをローカルで実行します (モデル呼び出しなし)。
pip install -e 。 # コアには依存関係がありません
キッチンラッシュ ベンチ --ベースライン ランダム --ティア イージー --シード 12 --トライアル 2
Kitchenrush calibrate --tier easy --latency-budget 1 # リファレンス Chef がレイテンシーによってどのように劣化するかを確認します
# ブラウザでゲームを視聴します (スクリプト化されたシェフ):
キッチンラッシュ リプレイ --oracle --tier easy --seed 0 # ui/replays/easy_seed0.json を書き込みます
cd ui && python3 -m http.server 8000 # 次に http://localhost:8000 を開きます
# ...または、1 つのクロックで最大 4 つのモデルを並べてレースします: ?replays=a.json,b.json (ui/README.md を参照)
実際のモデルのベンチマークを行うには、プロバイダーのサポートと API キーを追加します。
pip install -e ' .[プロバイダー] '
キッチンラッシュベンチ --model anthropic:claude-sonnet-4-6 --tier Medium --latency-budget 1
LiteLLM ルーティング可能なモデルは、 Provider:model を介して機能します。完全にカスタムしたプラグインも可能です
client — 必要なのは名前とgenerate(system,messages,tools) -> ModelResponseだけです
register_adapter で登録されたメソッド。 CLI コマンド: run 、 Bench 、 Reply 、 Seeds 、
校正します。
docs/RULES.md — コード検証済みの権威あるルールセット
docs/METHODOLOGY.md — KR メトリクス、B の計算、統計プロトコル
docs/CALIBRATION.md — gen-1.0 フリーズの背後にある証拠
docs/LIMITATIONS.md — KR が測定するものと測定しないもの (読む価値あり)
結果を引用する前に）
docs/OBJECTIONS.md — 予想される批判とデータによる回答
docs/SUBMISSIONS.md · docs/CONTAMINATION.md —
リーダーボードの契約とデータの健全性
あなたの作品でキッチン ラッシュを使用している場合は、それを引用してください (機械可読コピー
引用.cff ):
@ソフトウェア { キッチンラッシュ2026 、
著者 = { エレダス、バシム } 、
title = { Kitchen Rush: 正確かつ高速なツール呼び出しのベンチマーク } ,
URL = { https://github.com/bassimeledath/kitchen-rush } 、
年 = { 2026

}
}
ライセンス
Kitchen Rush: 正確かつ高速なネイティブ ツール呼び出しのベンチマーク
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

Kitchen Rush: a benchmark for accurate AND fast native tool calling - bassimeledath/kitchen-rush

GitHub - bassimeledath/kitchen-rush: Kitchen Rush: a benchmark for accurate AND fast native tool calling · GitHub
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
bassimeledath
/
kitchen-rush
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
70 Commits 70 Commits .github .github assets assets docs docs leaderboard leaderboard scripts scripts src/ kitchenrush src/ kitchenrush tests tests ui ui .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
An agent tool-calling benchmark where latency matters as much as intelligence.
Most tool-calling benchmarks (BFCL, τ-bench, ToolSandbox, AppWorld) check whether a model
makes the right calls — and the world politely waits while it thinks. That's fine for offline
tasks. But if you're building a voice assistant, a live-ops agent, or anything realtime, you
care about two things at once: does the model do the right thing, and does it do it fast
enough? A model that finds the perfect answer after thirty seconds of reasoning is, for you,
the wrong model.
Kitchen Rush measures both at once, by construction: the time a model spends thinking is
converted into game time that passes before its actions land. While the model deliberates,
food keeps cooking, food burns, and order deadlines slip away. Speed and accuracy aren't two
charts you squint at — they're one score, experienced the way a deployment would experience
them.
The model plays a chef in an Overcooked -style
kitchen. Orders stream in (burgers, soups, ramen…), and the model fulfils them with ordinary
native function calls — collect , chop , cook , plate , serve — racing deadlines,
burn timers, and a combo bonus for consecutive successful dishes. Three deliberate changes from
Overcooked:
Latency is the game. Every model response first charges its thinking time to the shared
world clock, then its actions execute. (You can chain several calls in one response and pay
the latency once — decisiveness is rewarded.)
No joystick skills. The chef walks itself to the right station automatically; travel
time is charged inside the action. What's being tested is choosing the right action
sequence under time pressure , not video-game reflexes.
Fully deterministic. Same seed, same actions, same latencies → exactly the same episode,
every time, on any machine. Every run can be replayed in a browser viewer and audited.
Every episode produces a single 0–100 score we call KR (the Kitchen Rush score ). It's
graded on a curve between two fixed anchors: KR 0 means "no better than doing nothing and
letting every order expire," and KR 100 means "matched a scripted reference chef that plays
the same kitchen with zero latency."
A worked example makes it concrete. Say that on one kitchen the do-nothing chef finishes at
−60 points (every order expired), the zero-latency reference chef finishes at +140 ,
and your model finishes at +40 . There are 200 points between the two anchors and your
model covered 100 of them, so its KR is 50 — it closed half the gap to the reference.
Average that over many seeded kitchens and you have the leaderboard number
( docs/METHODOLOGY.md has the full formula).
Here's the knob that makes Kitchen Rush flexible: every kitchen is generated at a latency
budget B ( --latency-budget , in seconds per decision). Think of B as the pace the
kitchen is priced for : order deadlines are set so that a chef spending exactly B seconds on
each decision can finish every order, with roughly 1.4–1.6× headroom to spare. Each B gets its
own leaderboard — results at different budgets are never averaged together.
For the mathematically inclined, the pricing is exact:
deadline = arrival + ⌈σ · C(B)⌉, where C(B) = A + K·B
A is the order's intrinsic cooking/walking time, K is how many decisions a competent plan
needs, and σ is the headroom (1.4–1.6 by tier). So a model that actually decides in ℓ seconds
gains or loses K·(B − ℓ) seconds of breathing room per order. Faster than B? You bank slack
and serve while orders are still worth full value. Slower? You eat through the headroom, and
orders start becoming unfinishable at around ℓ ≈ B + (σ−1)·C(B)/K — about 3–4 s/decision at
B=1 on the current tiers, which is exactly where our calibration sweep shows the reference
chef collapsing ( docs/METHODOLOGY.md §2 ,
docs/CALIBRATION.md ).
And in plain deployment terms: the model that wins at B=1s is the best pick when every
decision has to land in about a second — on the benchmark's reproducible clock that's a
budget of roughly 65 output tokens per decision, i.e. terse, single-shot tool dispatch — what a
voice agent needs. B=5s buys about 730 tokens per decision — enough for a short burst of
reasoning, what an interactive assistant can afford. The same model can rank very differently on the
two boards, and that reordering is precisely what the benchmark is for.
17 model configurations × 12 seeds × {medium, hard} kitchens × two latency budgets — 816
episodes so far. Each chart is one latency budget; bars are mean KR, whiskers are 95%
confidence intervals. The full per-tier table (with costs, reasoning tokens, and serve rates)
is at leaderboard/results/board.md .
The left board (B=1s) is the realtime test: the kitchen is priced for one second per
decision, which on the benchmark's clock buys about 65 output tokens — terse, single-shot tool
dispatch. Winning here means "the model I'd trust to drive a voice agent or a live dashboard."
The right board (B=5s) prices the same kitchens for five seconds per decision (~730
tokens — room for a short burst of reasoning), what an interactive assistant can afford.
Read them side by side — that contrast is the product. Under tight realtime pressure (B=1s)
the fast no-reasoning models hold the podium: gemini-3.1-flash-lite runs nearly even with
claude-sonnet-4.6 (32 vs 37). Give every decision five seconds instead and the board
reorders: gpt-5.4-mini with low reasoning rockets from near-zero to a dead heat with
sonnet (44 vs 44) at about a fifth of the cost , while flash-lite drops to half its B=1
standing. The same mini with reasoning fully off scores 0.0 at both budgets — reasoning it
can't afford at B=1 is exactly what makes it a frontier-level tool caller at B=5. That's the
latency tax, made visible. ( ·think rows ran with reasoning on at low effort; everything
else with reasoning off — fast single-shot dispatch is the honest realtime default. One row
you might expect is missing: there is no claude-sonnet-4.6·think , because Anthropic's API
does not allow extended thinking when tool calls are forced, and the harness forces tool
calls — sonnet competes thinking-off only.)
The flip, watched live: the same two models from the clip at the top,
but in a kitchen priced at B=5s. Now the mini's reasoning burst is affordable — it finishes
every order at 99 raw points (KR 86) while sonnet is still cooking at 40. This is the
mini's best kitchen — the chart above shows the average, a 44–44 tie across all 24 — but the
direction is real: it wins the medium tier at B=5 outright (59 vs 52). Same models, different
latency budget, different winner: that's exactly what the two boards measure.
Two minutes — run the scripted reference chef locally (no model calls):
pip install -e . # the core has zero dependencies
kitchenrush bench --baseline random --tier easy --seeds 12 --trials 2
kitchenrush calibrate --tier easy --latency-budget 1 # see how the reference chef degrades with latency
# watch a game in the browser (scripted chef):
kitchenrush replay --oracle --tier easy --seed 0 # writes ui/replays/easy_seed0.json
cd ui && python3 -m http.server 8000 # then open http://localhost:8000
# ...or race up to 4 models side-by-side on one clock: ?replays=a.json,b.json (see ui/README.md)
To benchmark a real model, add provider support and your API key:
pip install -e ' .[providers] '
kitchenrush bench --model anthropic:claude-sonnet-4-6 --tier medium --latency-budget 1
Any LiteLLM-routable model works via provider:model . You can also plug in a fully custom
client — it only needs a name and a generate(system, messages, tools) -> ModelResponse
method, registered with register_adapter . CLI commands: run , bench , replay , seeds ,
calibrate .
docs/RULES.md — the authoritative, code-verified ruleset
docs/METHODOLOGY.md — the KR metric, the math of B, statistical protocol
docs/CALIBRATION.md — the evidence behind the gen-1.0 freeze
docs/LIMITATIONS.md — what KR does and doesn't measure (worth reading
before citing results)
docs/OBJECTIONS.md — anticipated critiques, answered with data
docs/SUBMISSIONS.md · docs/CONTAMINATION.md —
leaderboard contract & data hygiene
If you use Kitchen Rush in your work, please cite it (machine-readable copy in
CITATION.cff ):
@software { kitchenrush2026 ,
author = { Eledath, Bassim } ,
title = { Kitchen Rush: A Benchmark for Accurate and Fast Tool Calling } ,
url = { https://github.com/bassimeledath/kitchen-rush } ,
year = { 2026 }
}
License
Kitchen Rush: a benchmark for accurate AND fast native tool calling
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
