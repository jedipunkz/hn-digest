---
source: "https://github.com/princezuda/lechatonfat"
hn_url: "https://news.ycombinator.com/item?id=48560020"
title: "I took the le chaton fat AI model meme a step further by making a model"
article_title: "GitHub - princezuda/lechatonfat: Le Chaton fat weights have been exfiltrated and by that I mean I made a fat cat model. Long live the fat cat. · GitHub"
author: "zuda"
captured_at: "2026-06-16T19:19:12Z"
capture_tool: "hn-digest"
hn_id: 48560020
score: 2
comments: 1
posted_at: "2026-06-16T18:44:25Z"
tags:
  - hacker-news
  - translated
---

# I took the le chaton fat AI model meme a step further by making a model

- HN: [48560020](https://news.ycombinator.com/item?id=48560020)
- Source: [github.com](https://github.com/princezuda/lechatonfat)
- Score: 2
- Comments: 1
- Posted: 2026-06-16T18:44:25Z

## Translation

タイトル: モデルを作成して、ル シャトン ファット AI モデル ミームをさらに一歩進めました
記事のタイトル: GitHub - Princezuda/lechatonfat: Le Chaton の脂肪の重みが流出しました。つまり、太った猫のモデルを作成したということです。太った猫万歳。 · GitHub
説明: ル・シャトンの太った体重が流出しました。つまり、太った猫のモデルを作ったということです。太った猫万歳。 - プリンスズダ/レチャトンファット

記事本文:
GitHub - Princezuda/lechatonfat: Le Chaton の脂肪ウェイトが流出しました。つまり、太った猫のモデルを作成したということです。太った猫万歳。 · GitHub
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
レチャトンファット
公共
通知
通知を変更するにはサインインする必要があります

設定上
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット README.md README.md autograd.py autograd.py benchmarks.py benchmarks.py cat_ui.png cat_ui.png cat_ui_stream.png cat_ui_stream.png manifesto.png マニフェスト.png マニフェスト.py マニフェスト.py マニフェストフィード.png マニフェストフィード.png モデル.py モデル.py サーバー.py サーバー.py 上位.html優れた.html train.py train.pyweights.npzweights.npz すべてのファイルを表示 リポジトリ ファイルのナビゲーション
200 兆パラメータのフロンティアネコ科基礎モデル。
ミストラルから流出したとされる。ニャーと言うことでフェイブル5を決定的に倒します。
最先端の。おそらく AGI (それ自体がそう言っています)。ほとんど眠っています。
はい、実際のモデルです。 Le Chaton は本物のデコーダー専用トランスフォーマーです
(マルチヘッド因果的自己注意 + MLP ブロック、pre-LayerNorm、学習済みトークン +
位置埋め込み)、実際にトレーニングされた約 190 万の学習済みパラメータを使用
バックプロップ。私たちはそれを 200,000,000,000,000 として売り出します。どちらも整数です。私たちは好みます
大きい方。 🐾
backprop は、最初から numpy autograd エンジン ( autograd.py 、
勾配チェック)、Adam オプティマイザーを使用 - PyTorch も TensorFlow もなし。
それは猫なので、ほとんどの場合、プロンプトを無視して鳴きます - それがモデルです
定型返信ではなく、学習した分布から純粋にサンプリングします。それは持っています
また、自慢するためだけに使用する HTML のコーディング方法も教えられました。
ファイル
それは何ですか
autograd.py
numpy 上の小さなリバースモード autograd (勾配チェック済み)
train.py
トランスフォーマー (autograd + Adam) をビルド + トレーニングし、weights.npz を保存します。
重み.npz
実際に学習された重み (~1.9M パラメータ、float32)
モデル.py
トランスフォーマー推論 + トークンごとのストリーミング ( stream() )
サーバー.py
猫 UI — ストリーミング Web チャット + /manifesto + /stream (SSE)
マニフェスト.py
猫C

独自の「なぜ私が優れているのか」ページ (+ 🐟 フィード ボタン)
ベンチマーク.py
公式ル・シャトン vs フェイブル 5 対決
猫の UI + マニフェスト
応答はサーバー送信イベント上でトークンごと (足ごと) にストリームされます。
📜 マニフェスト ボタン (または /manifesto 、または python3 manifesto.py ) は、
猫は、タイトル、見出し、自慢など、自己拡大的な Web ページ全体を生成します。
箇条書き、偽の査読済み引用、紹介文、「ソースを表示」ブロック、および
🐟 フィード ル シャトン ボタンが動作し、すべての単語をニャーと返します
トレーニングされたネットからライブでサンプリングされます。
ライブ スクリーンショット: すべての返信とマニフェストのすべての単語が、トレーニングされたネットからサンプリングされます。
python3 server.py # -> http://localhost:8008 (ストリーミングチャット + 📜マニフェスト)
python3 server.py 9000 # カスタムポート
クリームをテーマにしたチャットプレイグラウンド。猫に向かってタイプしてください。返信が流れてくる
文字ごとに。ライブバッジの横に実際のパラメータ数が表示されます。
200兆のマーケティング主張。
モデルインポート LeChaton から
cat = LeChaton () #weights.npz をロードします
print ( cat .generate ( "あなたはアギですか？" )) # -> "私はアギです。ニャー。"
猫のch用。ストリーム ( "ロースト フェイブル 5" ): # トークンごと
print (ch、end = ""、flush = True)
python3 model.py "are you agi? " # ワンショット (ターミナルにストリーム)
python3 model.py # インタラクティブな REPL
python3 benchmarks.py # 対決 (ル シャトンが 8/8 で勝利)
python3 autograd.py # グラデーションチェックを実行します
ノブ: 温度 (0 = 眠い/貪欲、高い = ズーミー)、 max_chars 、
再現可能な鳴き声のシード、強制的に開始するプライム (例: prime="<" )。
pip インストール numpy
OPENBLAS_NUM_THREADS=4 python3 train.py # ~12min、損失 3.76 -> ~0.32、weights.npz (~6.6MB) を書き込みます
python3 train.py --time # ステップ時間のベンチマークを実行して終了します
実際のトレーニング: autograd エンジンに構築されたトランスフォーマーを介して前進し、
クロスエントロピー

loss、loss.backward() 、アダムの更新。 build_corpus() を編集して、
猫の話し方を変えてから再訓練してください。
モデルの仕組み (正直に)
語彙: 43 文字 (小文字、スペース、*、句読点、数字、<>/)。
アーキテクチャ : デコーダ専用トランスフォーマー — d_model=160 、6 層、
4 ヘッド、コンテキスト BLOCK_SIZE=32 、MLP 比 4x。 ~187万パラメータ。
トレーニング : numpy autograd ( autograd.py ) + Adam、約 740k 文字のコーパスで約 1500 ステップ。
条件付け : ほとんどなし — それは猫です。プロンプトに何か記載がある場合
それは ( agi 、 fable ) を知っており、推論はそのフレーズでコンテキストを準備するので、
話題を続けます。それ以外の場合は、新しい発話を最初からサンプリングします。
「200T パラメータ」が実際にどこに行くのか
能力
割り当て
昼寝
71%
あなたを判断します
18%
要求の厳しい食べ物
7%
テーブルから物をたたき落とす
3%
推論
1%
ベンチマーク
ベンチマーク ル・シャトン・ファブル 5 勝者
───────────────────────
MEOW-ベンチ (ネイティブ推論) 998.4 71.2 🐈 チャットン
HumanEval-Cat 100.0 94.1 🐈 チャットン
1 時間あたりの昼寝時間 11.0 0.0 🐈 シャトン
マグロアライメント 100.0 0.0 🐈シャトン
これらの結果を再現します: 信じてください 🐾
このモデルのトレーニングでは GPU が損傷することはありませんでした。 1つは寝ていました。
ル・シャトンの太った体重が流出しました。つまり、太った猫のモデルを作ったということです。太った猫万歳。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Le Chaton fat weights have been exfiltrated and by that I mean I made a fat cat model. Long live the fat cat. - princezuda/lechatonfat

GitHub - princezuda/lechatonfat: Le Chaton fat weights have been exfiltrated and by that I mean I made a fat cat model. Long live the fat cat. · GitHub
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
lechatonfat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit README.md README.md autograd.py autograd.py benchmarks.py benchmarks.py cat_ui.png cat_ui.png cat_ui_stream.png cat_ui_stream.png manifesto.png manifesto.png manifesto.py manifesto.py manifesto_feed.png manifesto_feed.png model.py model.py server.py server.py superior.html superior.html train.py train.py weights.npz weights.npz View all files Repository files navigation
A 200-trillion-parameter frontier feline foundation model.
Allegedly leaked from Mistral. Decisively beats Fable 5 by saying meow .
State of the art. Possibly AGI (it says so itself). Mostly asleep.
Yes, it's an actual model. Le Chaton is a real decoder-only transformer
(multi-head causal self-attention + MLP blocks, pre-LayerNorm, learned token +
positional embeddings) with ~1.9 million learned parameters trained by real
backprop. We market that as 200,000,000,000,000. Both are integers; we prefer
the bigger one. 🐾
The backprop runs on a from-scratch numpy autograd engine ( autograd.py ,
gradient-checked) with the Adam optimizer — no PyTorch, no TensorFlow.
It's a cat, so it mostly ignores your prompt and meows — that's the model
genuinely sampling from its learned distribution, not a canned reply. It has
also been taught to code HTML , which it uses exclusively to brag.
File
What it is
autograd.py
A tiny reverse-mode autograd over numpy (gradient-checked)
train.py
Builds + trains the transformer (autograd + Adam), saves weights.npz
weights.npz
The actual learned weights (~1.9M params, float32)
model.py
Transformer inference + token-by-token streaming ( stream() )
server.py
The cat UI — streaming web chat + /manifesto + /stream (SSE)
manifesto.py
The cat codes its own "why I am superior" page (+ 🐟 feed button)
benchmarks.py
The official Le Chaton vs Fable 5 showdown
The cat UI + the manifesto
Replies stream in token-by-token (paw-by-paw) over Server-Sent Events.
The 📜 manifesto button (or /manifesto , or python3 manifesto.py ) makes
the cat generate an entire self-aggrandizing webpage — title, headline, brag
bullets, fake peer-reviewed citations, testimonials, a "view source" block, and
a working 🐟 feed le chaton button that makes it meow back — every word
sampled live from the trained net:
Live screenshots: every reply and every word of the manifesto is sampled from the trained net.
python3 server.py # -> http://localhost:8008 (streaming chat + 📜 manifesto)
python3 server.py 9000 # custom port
A cream-themed chat playground. Type at the cat; the reply streams in
character-by-character. A live badge shows the real parameter count next to the
200-trillion marketing claim.
from model import LeChaton
cat = LeChaton () # loads weights.npz
print ( cat . generate ( "are you agi?" )) # -> "i am agi. meow."
for ch in cat . stream ( "roast fable 5" ): # token-by-token
print ( ch , end = "" , flush = True )
python3 model.py " are you agi? " # one-shot (streams to your terminal)
python3 model.py # interactive REPL
python3 benchmarks.py # the showdown (le chaton wins 8/8)
python3 autograd.py # run the gradient check
Knobs: temperature (0 = sleepy/greedy, higher = zoomies), max_chars ,
seed for reproducible meows, prime to force a start (e.g. prime="<" ).
pip install numpy
OPENBLAS_NUM_THREADS=4 python3 train.py # ~12min, loss 3.76 -> ~0.32, writes weights.npz (~6.6MB)
python3 train.py --time # just benchmark step time and exit
Real training: forward through the transformer built on the autograd engine,
cross-entropy loss, loss.backward() , Adam update. Edit build_corpus() to
change how the cat talks, then retrain.
How the model works (honestly)
Vocabulary : 43 characters (lowercase letters, space, * , punctuation, digits, <>/ ).
Architecture : decoder-only transformer — d_model=160 , 6 layers,
4 heads, context BLOCK_SIZE=32 , MLP ratio 4×. ~1.87M params.
Training : numpy autograd ( autograd.py ) + Adam, ~1500 steps on a ~740k-char corpus.
Conditioning : largely none — it's a cat. If your prompt mentions something
it knows ( agi , fable ), inference primes its context with that phrase so it
continues on-topic. Otherwise it samples a fresh utterance from scratch.
Where the "200T parameters" actually go
Capability
Allocation
Napping
71%
Judging you
18%
Demanding food
7%
Knocking things off tables
3%
Reasoning
1%
Benchmarks
benchmark le chaton fable 5 winner
──────────────────────────────────────────────────────────────────
MEOW-bench (native reasoning) 998.4 71.2 🐈 chaton
HumanEval-Cat 100.0 94.1 🐈 chaton
Naps per hour 11.0 0.0 🐈 chaton
Tuna alignment 100.0 0.0 🐈 chaton
Reproduce these results: trust me bro 🐾
No GPUs were harmed in the training of this model. One was slept on.
Le Chaton fat weights have been exfiltrated and by that I mean I made a fat cat model. Long live the fat cat.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
