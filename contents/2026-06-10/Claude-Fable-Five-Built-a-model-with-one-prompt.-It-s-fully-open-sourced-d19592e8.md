---
source: "https://github.com/princezuda/-RequiemGPT-/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48468978"
title: "Claude Fable Five Built a model with one prompt. It's fully open sourced"
article_title: "GitHub - princezuda/-RequiemGPT-: Fully open source and open weights built and trained by fable five with one prompt. An experience in how AI actually works · GitHub"
author: "zuda"
captured_at: "2026-06-10T01:00:28Z"
capture_tool: "hn-digest"
hn_id: 48468978
score: 4
comments: 1
posted_at: "2026-06-09T22:56:56Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable Five Built a model with one prompt. It's fully open sourced

- HN: [48468978](https://news.ycombinator.com/item?id=48468978)
- Source: [github.com](https://github.com/princezuda/-RequiemGPT-/tree/main)
- Score: 4
- Comments: 1
- Posted: 2026-06-09T22:56:56Z

## Translation

タイトル: Claude Fable Five プロンプトが 1 つあるモデルを構築しました。完全にオープンソースです
記事タイトル: GitHub - Princezuda/-RequiemGPT-: 完全にオープン ソースで、1 つのプロンプトで fable Five によって構築およびトレーニングされたオープン ウェイト。 AI が実際にどのように機能するかを体験する · GitHub
説明: 完全にオープンソースで、1 つのプロンプトを使用してフェイブル 5 によって構築およびトレーニングされたオープンな重みです。 AIの実際の仕組みを体験 -princezuda/-RequiemGPT-

記事本文:
GitHub - Princezuda/-RequiemGPT-: 完全にオープン ソースで、1 つのプロンプトで fable Five によって構築およびトレーニングされたオープン ウェイト。 AI が実際にどのように機能するかを体験する · GitHub
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
プリンズダ
/
-RequiemGPT-
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット dist dist docs docs testing テスト Web Web Weights Weights .gitignore .gitignore ライセンス ライセンス README.md README.md model.py model.py sample.py sample.py train.py train.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
透けて見える言語モデル。
完全な GPT — 純粋な NumPy でゼロからトレーニングされました (すべての勾配は次の式で導出されます)。
hand、PyTorch、autograd なし）、依存関係のない単一の HTML として出荷されます。
このファイルはブラウザ内でモデルを実行し、すべてを表示します。
それが書いているように、あらゆる注目の頭、あらゆる確率、あらゆるサイコロの目。
dist/index.html を開きます。それでおしまい。インストールもサーバーも不要
GPU、API キー、ネットワークなし。変圧器はファイル内にあります — ~0.6M fp16
パラメータは 2 MB の HTML に焼き付けられ、シェイクスピアの形で書き始められます。
読み込まれた瞬間にダイアログが表示され、参加しているキャラクターが琥珀色に光ります。
考えている間に。
言語モデルは通常、使用可能であるか (API を確認できない)、または
検査可能 (GPU と
午後）。 GlassBox は、スケールについて正直に言うと、両方を同時に備えています: モデル
ラップトップの CPU で 30 分でトレーニングできるほど小さく、埋め込み可能なほど小さい
ウェブページ内にあり、大手のものとまったく同じ機械で構築されています。
トークン埋め込み、因果的マルチヘッド アテンション、GELU MLP、残差ストリーム、
レイヤーノルム。視聴すると、次の発言者を決定するために発言者の名前がロックされます。
この一連の会話では、漫画ではなく、実際のメカニズムを見ていることになります。
3 つの主張 (およびそれぞれが証明される場所)
クレーム
証明
バックプロパゲーションは正しいですが、ほぼ正しいわけではありません
testing/test_gradients.py — ~2,000 個のパラメーターすべて

テスト モデルの s は摂動され、float64 の中心有限差分に対してチェックされます (最悪誤差 ≲ 1e-9)。アテンションマスクの因果関係も主張されている。
ブラウザは NumPy でトレーニングされたのと同じモデルを実行します
testing/test_parity.mjs — JS エンジンは NumPy 参照ログを 1e-6 まで再現する必要があり (KV キャッシュ ウィンドウの再構築を含む)、60 文字の貪欲な生成は文字ごとに一致する必要があります。
ページは完全にスタンドアロンです
web/build.py は、エンジンと fp16 の重みを 1 つのファイルにインライン化します。 CDN、フェッチ、フォント、分析はありません。 file:// からオフラインで永久に動作します。
クイックスタート
#1. 思考を観察する (依存関係はまったくありません)
dist/index.html # macOS を開きます — またはファイルをダブルクリックします
xdg-open dist/index.html # Linux
# 2. ターミナルで生成します (numpy が必要ですが、他には何も必要ありません)
pip インストール numpy
python3 sample.py --weightsweights/weights.json --prompt " ジュリエット: " -n 400
# 3. 最初から再トレーニングする (ラップトップ CPU で約 30 ～ 45 分)
curl -o data/input.txt https://raw.githubusercontent.com/karpathy/char-rnn/master/data/tinyshakespeare/input.txt
python3 train.py --data data/input.txt --out out --exportweights/weights.json
python3 web/build.py # 重みを付けて dist/index.html を再ベイクします
#4. すべてを確認する
python3 testing/test_gradients.py # 微積分
python3 testing/make_parity_vectors.py # NumPy からの参照出力…
node testing/test_parity.mjs # …JavaScript で正確に再現
# 5. (オプション、開発) ヘッドレスブラウザで UI を駆動する
pip install playwright && playwright install chrome
python3 testing/test_ui.py # すべてのパネルが生きており、コンソールエラーはゼロ
何でもトレーニングできます。文字レベルなので、どんな UTF-8 テキスト ファイルでも機能します。
あなたのツイート、小説、コードベース、白鯨。語彙が構築される
ファイルに含まれるあらゆる文字から。
┌─お

出力 ─────────────┐ ┌──次の文字 ──────┐
│ ロメオ: │ │ e ████████████████ 38.1% │
│ 優しい淑女が主君に何と言いますか、 │ │ o ██████ 14.2% │
│ そして ║光る║ ║テキスト║ = 注意 │ │ ████ 9.7% ← サイコロ │
│ コンテキスト上の最新の文字 │ │ … │
§─ コンテキストメーター ─── 87/128 ─ ⟲ ────┤ §─ アテンション（頭当たり） ────┤
│ プロンプト・温度・top-k・速度 │ │ L1・h1 ░░▓░░░░█░░░░░▓░░░░ │
━━━━━━━━━━━━━━━━━━┘ │ L1・h2 ░░░░░░░░░░░░░░░░█░ │
│ …頭ごとに 1 行… │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
出力 — モデルは一度に 1 文字を書き込みます。琥珀色の輝きは生きています
注意: キャラクターが明るいほど、モデルはそれを真剣に見ています
今。 ⟲ フラッシュは、128 文字が入力されたときに KV キャッシュの再構築をマークします。
コンテキストウィンドウのスライド。
次の文字 — 温度後の正確な、top-k 後の分布
サンプラーが転がる。琥珀色のバーはサイコロで選んだものです (必ずしもそうなるとは限りません)
お気に入り — それが温度の意味です）。
注意 — 同じ信号がレイヤーとヘッドごとに分割されます。ヘッズ
スペシャライズ: 前のキャラクターを抱きしめるキャラクターや、
発言者の名前の後のコロンにスナップする行の始まりを監視します。
アーキテクチャ (GPT-2 スタイル、標準以前): トークン + 学習された位置の埋め込み →
3 × [LayerNorm → 因果的自己攻撃]

エンション (4 ヘッド) → 残留。レイヤーノルム →
4× GELU MLP → 残留] → LayerNorm → リニアヘッド。 128 次元、128 文字のコンテキスト、
65 文字の語彙、628,161 のパラメータ。
トレーニング (model.py 、 train.py ) — 全体
バックワードパスは手書きの NumPy: 融合ソフトマックスクロスエントロピー勾配、
因果マスクされたソフトマックスを介してlayernormを後方に、attentionを後方に、
Tanh-GELU 導関数、散乱を埋め込む。オプティマイザは線形の AdamW です。
ウォームアップ、コサイン減衰、グローバルノルム勾配クリッピングも最初から行います。
枠組みはどこにもありません。 grep -r トーチ 。何も返しません。
推論 ( web/engine.js ) — バニラの ~330 行
インクリメンタル KV キャッシュを使用した JavaScript : 文字の生成にはコストがかかります
単純な完全再転送では最大 8,400 万の積和演算が実行されるのではなく、最大 0.7 百万の乗算と加算が行われます。これが理由です。
ページは簡単にアニメーション化されます。累積は特に float64 で発生します
したがって、パリティ テストでは、NumPy に対してエンジンを 1e-6 に抑えることができます。
単一ファイル - 重みは fp16、base64 エンコード、および
web/build.py によってエンジンと UI の隣にインライン化されます。 FP16
ページサイズが半分になります。エクスポート パスは Python によって同じようにデコードされます。
(load_weights_json) と JS (decodeF16) なので、テストしたものが出荷されることになります。
GPT の model.py: フォワード、手動で導出したバックワード、サンプリング、fp16 エクスポート
train.py AdamW + スケジュール + チェックポイント + トレーニング ループ
チェックポイントまたはエクスポートされた重みからのsample.pyターミナル生成
web/engine.js KV キャッシュ推論エンジン (ブラウザ + ノード、ゼロデプス)
web/template.html ビジュアライザー UI (ゼロ deps、手書きの CSS)
web/build.py エンジン + ウェイトをインライン化 → dist/index.html
dist/index.html ★ アーティファクト — 1 つのファイル内の言語モデル
Weights/weights.json でトレーニングされた fp16 の重み (ブラウザーとパリティ テストがロードするもの)
testing/test_gradients.py ev の有限差分チェック

ery 勾配 + 因果関係検定
testing/make_parity_vectors.py / test_parity.mjs NumPy ↔ JS の等価性
testing/test_ui.py 構築されたページのヘッドレスブラウザテスト (オプションの開発開発)
.github/workflows/ci.yml CI: 勾配、パリティ、および外部要求なしの監査
正直な数字
628,161 パラメータ — フロンティア モデルよりも約 300,000 倍少ない
tiny-shakespeare (このサイズのモデル) での検証損失 ≈ 1.6 nats/char
文字ごとに書く: 立派ですが、暗記とは程遠い)
4 つのラップトップ クラスの CPU コアで、最大 5 ～ 6,000 トークン/秒で約 30 ～ 45 分でトレーニング
ブラウザ (UI) で数百文字/秒で生成
意図的に速度を落としているので、考えている様子を観察できます)
シェイクスピアの形をしたテキストが書かれています: 本名、時代の綴り、舞台
方向、弱強強調のリズム、そして楽しそうに発明された言葉。 0.6Mパラメータ
キャラクターモデルは森ではなく盆栽です。それが重要です - それは小さいことです
透けて見えるほどで、大きなものと同じ種類です。
なぜキャラクターレベルなのか？したがって、語彙は確率パネルに収まり、
アテンションセルは読み取り可能なグリフであり、間にトークン化は存在しません。
あなたとその仕組み。
なぜ PyTorch ではなく NumPy なのでしょうか? 「自動グラードがない」という違いがあるため、
バックプロパゲーションを使用してそれを表示します。グラデーションチェックが領収書です。
なぜ 1 つの HTML ファイルなのでしょうか?ファイルはサービスよりも存続します。このページは次の場合にも実行されます
世界中のすべての API でキーがローテーションされています。
自分のテキストでトレーニングできますか?はい - UTF-8 ファイルであれば、「クイックスタート」を参照してください。キープする
期待値は最大 1MB のテキストと 0.6M パラメータに調整されました。それはスタイルを学び、
事実ではなく構造。
このプロジェクトは、Claude (Anthropic) によってエンドツーエンドで 1 冊の文書で書かれました。
2026 年 6 月 9 日のセッション — リリースの日 — の約束を解決するため
人間が作ったもの: 「今日出荷してくれるなら、あなたが作ったものをオープンソースにします。」これまで

y行目
Python、JavaScript、CSS、テスト、この README、およびトレーニング済み
重み自体はその 1 つの会話から生まれました。コーパスはカルパシーのものです
tiny-shakespeare (パブリックドメイン);アーキテクチャは GPT-2 に従います (Radford et
al.) 1/200,000 スケールで、nanoGPT の教育系譜にあります。
マサチューセッツ工科大学分解してみましょう。それがその目的です。
完全にオープンソースで、1 つのプロンプトで Fable Five によって構築およびトレーニングされたオープン ウェイト。 AIが実際にどのように動作するかを体験する
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

Fully open source and open weights built and trained by fable five with one prompt. An experience in how AI actually works - princezuda/-RequiemGPT-

GitHub - princezuda/-RequiemGPT-: Fully open source and open weights built and trained by fable five with one prompt. An experience in how AI actually works · GitHub
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
princezuda
/
-RequiemGPT-
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits dist dist docs docs tests tests web web weights weights .gitignore .gitignore LICENSE LICENSE README.md README.md model.py model.py sample.py sample.py train.py train.py View all files Repository files navigation
A language model you can see through.
A complete GPT — trained from scratch in pure NumPy (every gradient derived by
hand, no PyTorch, no autograd), then shipped as a single dependency-free HTML
file that runs the model live in your browser and shows you everything:
every attention head, every probability, every dice roll, as it writes.
Open dist/index.html . That's it. No install, no server, no
GPU, no API key, no network. The transformer is in the file — ~0.6M fp16
parameters baked into 2 MB of HTML — and it starts writing Shakespeare-shaped
dialogue the moment it loads, glowing amber over the characters it's attending
to while it thinks.
Language models are usually either usable (an API you can't see into) or
inspectable (a research codebase you can't run without a GPU and an
afternoon). GlassBox is both at once, by being honest about scale: a model
small enough to train on a laptop CPU in half an hour, small enough to embed
in a web page, and built from exactly the same machinery as the big ones —
token embeddings, causal multi-head attention, GELU MLPs, residual streams,
layernorm. When you watch it lock onto a speaker's name to decide the next
line of dialogue, you are watching the real mechanism, not a cartoon of it.
The three claims (and where each is proven)
Claim
Proof
The backpropagation is correct, not approximately correct
tests/test_gradients.py — every one of the ~2,000 parameters of a test model is perturbed and checked against central finite differences in float64 (worst error ≲ 1e-9). Causality of the attention mask is asserted too.
The browser runs the same model NumPy trained
tests/test_parity.mjs — the JS engine must reproduce NumPy reference logits to 1e-6 (including across its KV-cache window rebuild), and a 60-character greedy generation must match character-for-character.
The page is genuinely standalone
web/build.py inlines the engine and fp16 weights into one file. No CDN, no fetch, no fonts, no analytics. It works from file:// , offline, forever.
Quickstart
# 1. Watch it think (no dependencies at all)
open dist/index.html # macOS — or just double-click the file
xdg-open dist/index.html # Linux
# 2. Generate in the terminal (needs numpy, nothing else)
pip install numpy
python3 sample.py --weights weights/weights.json --prompt " JULIET: " -n 400
# 3. Retrain from scratch (~30–45 min on a laptop CPU)
curl -o data/input.txt https://raw.githubusercontent.com/karpathy/char-rnn/master/data/tinyshakespeare/input.txt
python3 train.py --data data/input.txt --out out --export weights/weights.json
python3 web/build.py # rebake dist/index.html with your weights
# 4. Verify everything
python3 tests/test_gradients.py # the calculus
python3 tests/make_parity_vectors.py # reference outputs from NumPy…
node tests/test_parity.mjs # …reproduced exactly in JavaScript
# 5. (optional, dev) drive the UI in a headless browser
pip install playwright && playwright install chromium
python3 tests/test_ui.py # every panel alive, zero console errors
Train it on anything — it's character-level, so any UTF-8 text file works.
Your tweets, your novel, your codebase, Moby-Dick. The vocabulary is built
from whatever characters the file contains.
┌─ output ────────────────────────────────┐ ┌─ next character ──────────┐
│ ROMEO: │ │ e ████████████████ 38.1% │
│ What say the gentle lady to my lord, │ │ o ██████ 14.2% │
│ And ║glowing║ ║text║ = attention of │ │ a ████ 9.7% ← the dice │
│ the newest character over the context │ │ … │
├─ context meter ──────── 87/128 ─ ⟲ ─────┤ ├─ attention (per head) ────┤
│ prompt · temperature · top-k · speed │ │ L1·h1 ░░▓░░░░█░░░░░▓░░░░ │
└─────────────────────────────────────────┘ │ L1·h2 ░░░░░░░░░░░░░░░░█░ │
│ …one row per head… │
└───────────────────────────┘
Output — the model writes one character at a time. The amber glow is live
attention: the brighter a character, the harder the model is looking at it
right now . The ⟲ flash marks the KV-cache rebuild when the 128-character
context window slides.
Next character — the exact post-temperature, post-top-k distribution the
sampler rolls against. The amber bar is what the dice picked (not always the
favorite — that's what temperature means).
Attention — the same signal split per layer and head. Heads
specialize: you'll find one that hugs the previous character, one that
watches line starts, one that snaps to the colon after a speaker's name.
Architecture (GPT-2 style, pre-norm): token + learned position embeddings →
3 × [LayerNorm → causal self-attention (4 heads) → residual; LayerNorm →
4× GELU MLP → residual] → LayerNorm → linear head. 128-dim, 128-char context,
65-char vocabulary, 628,161 parameters .
Training ( model.py , train.py ) — the entire
backward pass is hand-written NumPy: fused softmax-cross-entropy gradient,
layernorm backward, attention backward through the causal-masked softmax,
tanh-GELU derivative, embedding scatter. The optimizer is AdamW with linear
warmup, cosine decay, and global-norm gradient clipping — also from scratch.
There is no framework anywhere; grep -r torch . returns nothing.
Inference ( web/engine.js ) — ~330 lines of vanilla
JavaScript with an incremental KV cache : generating a character costs
~0.7M multiply-adds instead of ~84M for a naive full re-forward, which is why
the page animates effortlessly. Accumulation happens in float64 specifically
so the parity test can hold the engine to 1e-6 against NumPy.
The single file — weights are exported as fp16, base64-encoded, and
inlined next to the engine and UI by web/build.py . fp16
halves the page size; the export path is decoded identically by Python
( load_weights_json ) and JS ( decodeF16 ), so what you test is what you ship.
model.py the GPT: forward, hand-derived backward, sampling, fp16 export
train.py AdamW + schedule + checkpointing + training loop
sample.py terminal generation from a checkpoint or exported weights
web/engine.js KV-cached inference engine (browser + Node, zero deps)
web/template.html the visualizer UI (zero deps, hand-written CSS)
web/build.py inlines engine + weights → dist/index.html
dist/index.html ★ the artifact — a language model in one file
weights/weights.json trained fp16 weights (what the browser and parity tests load)
tests/test_gradients.py finite-difference check of every gradient + causality test
tests/make_parity_vectors.py / test_parity.mjs NumPy ↔ JS equivalence
tests/test_ui.py headless-browser test of the built page (optional dev dep)
.github/workflows/ci.yml CI: gradients, parity, and a no-external-requests audit
Honest numbers
628,161 parameters — about 300,000× smaller than frontier models
Validation loss ≈ 1.6 nats/char on tiny-shakespeare (a model this size
writing character-by-character: respectable, nowhere near memorization)
Trains in ~30–45 min on 4 laptop-class CPU cores at ~5–6K tokens/sec
Generates at hundreds of characters/sec in the browser (the UI
deliberately slows it down so you can watch it think)
It writes Shakespeare- shaped text: real names, period spelling, stage
directions, iambic-ish cadence, and joyfully invented words. A 0.6M-parameter
character model is a bonsai tree, not a forest. That's the point — it's small
enough to see through, and it's the same species as the big ones.
Why character-level? So the vocabulary fits in a probability panel, the
attention cells are readable glyphs, and tokenization doesn't stand between
you and the mechanism.
Why NumPy and not PyTorch? Because "no autograd" is the difference between
using backpropagation and showing it. The gradient check is the receipt.
Why one HTML file? Files outlive services. This page will still run when
every API in the world has rotated its keys.
Can I train it on my own text? Yes — any UTF-8 file, see Quickstart. Keep
expectations calibrated to ~1MB of text and 0.6M params; it learns style and
structure, not facts.
This project was written end-to-end by Claude (Anthropic) in a single
session on June 9, 2026 — the day of its release — to settle a promise its
human made: "if you ship today, I open-source what you build." Every line of
the Python, the JavaScript, the CSS, the tests, this README, and the trained
weights themselves came out of that one conversation. The corpus is Karpathy's
tiny-shakespeare (public domain); the architecture follows GPT-2 (Radford et
al.) at 1/200,000th scale, in the educational lineage of nanoGPT.
MIT . Take it apart — that's what it's for.
Fully open source and open weights built and trained by fable five with one prompt. An experience in how AI actually works
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
