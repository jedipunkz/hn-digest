---
source: "https://github.com/ninjahawk/Subtext"
hn_url: "https://news.ycombinator.com/item?id=48811892"
title: "Show HN: LLM Thought Visualization"
article_title: "GitHub - ninjahawk/Subtext · GitHub"
author: "ninjahawk1"
captured_at: "2026-07-07T00:08:06Z"
capture_tool: "hn-digest"
hn_id: 48811892
score: 1
comments: 0
posted_at: "2026-07-06T23:34:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLM Thought Visualization

- HN: [48811892](https://news.ycombinator.com/item?id=48811892)
- Source: [github.com](https://github.com/ninjahawk/Subtext)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T23:34:27Z

## Translation

タイトル: HN を表示: LLM 思考の視覚化
記事タイトル: GitHub - ninjahawk/サブテキスト · GitHub
説明: GitHub でアカウントを作成して、ninjahawk/Subtext の開発に貢献します。

記事本文:
GitHub - ninjahawk/サブテキスト · GitHub
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
ニンジャホーク
/
サブテキスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
7 コミット 7 コミット docs docs media media .gitignore .gitignore ライセンス ライセンス README.md README.md Index.html Index.html Record_session.py Record_session.py 要件.txt 要件.txt サーバー.py サーバー.py start.bat start.bat test_client.py test_client.py test_multiturn.py test_multiturn.py verify_accuracy.py verify_accuracy.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
言語モデルの言語ワークスペースを観察するためのリアルタイムツール
読み、推論し、話すとき。
🌐 ブラウザでライブリプレイを見る · ▶ デモビデオ · 📄 論文 · 🔬 リファレンス実装
Anthropic の最近の研究では、内部表現の小さなセットが特定されました。
グローバル ワークスペースのように動作する言語モデル — J スペース —
内容はモデルによって口頭で報告され、意図的に調整され、
因果的に多段階の推論に使用されますが、周囲の大部分のニューラル
アクティビティはレポートにアクセスできないままです。識別ツールは、
ヤコビアン レンズ。任意の層での残留ストリームの活性化を
最終層の基礎を作成し、モデル自体の非埋め込みを通じてそれをデコードします。
答え: この内部状態がどの語彙を生成する傾向があるのか、
今ですか、それとも後でですか？
サブテキストは、地元のユーザーとのライブ会話中にそのメソッドを継続的に適用します。
モデル。すべてのトークンで — モデルがユーザーのメッセージを取り込むときと、
応答を生成している間、レンズは 9 つの深さで読み取られ、結果が返されます。
発生したままにレンダリングされます。モデル計算の中間ステップ
直接監視可能になる: 読み取り中に評決が形成され、その前にいくつかのトークンが生成される
任意の出力。計画された単語は高い強度を保持しますが、無関係なトークンは
放出される。ツーホップの質問は、彼らの暗黙の中期問題を表面化させます。
サブテキストは次のものと異なります

インタラクティブな読み出しはすでに利用可能です (例:
Neuronpedia デモ) は、会話型で継続的であるという点で、
ライブチャット中のレンズには、ユーザーのメッセージを読むフェーズが含まれます。
KV キャッシュを介して生成速度でストリーミングし、キャンバスを
トークンごとの台帳と単語ごとのインスペクター。セッションはエクスポートして再生できます
GPU を搭載していないブラウザでも利用できます。
レンズが出力に示さないもの
機器の価値は、モデルの内部状態とモデルの内部状態との間のギャップです。
その目に見えるテキスト。デモセッションの 3 つの瞬間:
1. 判決は返答に先立って行われます。出力のトークンは存在しません。モデル
まだ読んでいます これは正しいですか? 12 + 5 = 1 。ワークスペースはすでに保持されています
math、addition、算術演算、modulo — フェーズインジケーターはオレンジ色です
（読んで）。
2. 判断が形成され、言語化されます。返事が始まると（「いいえ、それは）
は ** ではありません…」)、不正なコードが高い強度でワークスペースを支配し、
方程式、計算、ステートメント co-active — 結論は次のとおりです。
「正しくありません」という言葉が表示される前に、いくつかのトークンが内部的に決済されます。
3. 他の言葉が話されている間、計画は実行されます。説明の途中で、
ワークスペースには、モジュロ、ビット単位、システム、数値、つまり技術的なものが保持されます。
現在の出力トークンが
無関係な。
これらは、コンシューマー ハードウェアのオープン 4B モデルで、レポートと
論文に記載されている現象の計画 (クロードスケールモデルを使用)、
2 ホップ署名を含む: レイヤ 20 のイタリアとレイヤ 26 のユーロ
世代が始まる前に、ブーツのような国の形をした問題について。
レンダリングされた各ワードはレンズの読み出しであり、モデルの出力ではありません。それは、
モデルをその単語に向けて配置する内部活性化。
垂直位置はレイヤーに対応します。初期層（知覚）

) にあります
上部。読み出し値は、発光に近づくにつれて下部レールに近づきます。
サイズと不透明度は、絶対的な読み取り強度をエンコードします。ディスプレイは、
レンズ確率からの固定単調マッピング。弱い読み取り値が表示される
弱い。黄色は、ユーザーの読み取り中に取得された読み取り値を示します。青いながら
生成しています。
ホバーすると、単語のレイヤーごとのアクティブ化プロファイルが表示されます。クリックすると開きます
ピーク強度、平均深さ、および強度履歴を備えた検査官。
右側のパネルには、キャンバスが厳選したすべての内容が記録されます。
現在アクティブな読み出しのライブランキングとトークンごとの台帳。
ブラウザ (単一の HTML ファイル) ⇐ websocket ⇐ server.py
Qwen3.5-4B (bf16、HF トランス、KV キャッシュ)
プレフィットヤコビアンレンズ: ニューロンペディア/ヤコビアンレンズ、リビジョン qwen-n1000
トークンごと: 9 層の残留フック → J_l トランスポート → アンエンベッド
→ フルボキャブラリーソフトマックス → ワードスタートトップ-k → フレーム
各交換には 2 つのフェーズがあります。単一の事前入力パスでユーザーのメッセージがカバーされます。
あらゆる位置でレンズの読み取り値が取得されます (読み取りフェーズ)。世代
次に、KV キャッシュをトークンごとに処理し、最新のレンズを読み取ります。
各ステップ（思考フェーズ）を位置づけます。レンズはレイヤーごとに
行列とベクトルの積とトークンごとのアン埋め込みにより、ストリーミングは
モデルのネイティブ生成速度。
フィルタリングを表示します。 Raw レンズの top-k には句読点と BPE が含まれています
継続フラグメント (例: itude 、 from cert‑itude )。
読み物として意味がある。表示は単語の頭文字の語彙に制限されます
トークン。リファレンス実装のmask_displayの後に
より厳格な単語開始基準。確率は全体にわたって計算されます
フィルタリングの前に語彙を確認するため、フィルタリングは可読性にのみ影響を及ぼし、決して影響しません。
読み出し自体。
verify_accuracy.py は、この実装のライブ パス (順方向) を比較します。

フック、
KV キャッシュが有効になっている) を、同一の参照 JacobianLens.apply() に対して実行します。
入力。ウォークスルー プロンプトの 4 つのレイヤー × 3 つのポジションにわたって、トップ 5
読み出し値は完全に一致し、ロジット間のコサイン類似度 ≥ 0.99998
ベクトルを生成し、予想される 2 ホップ中間を再現します。監査は次のように行うことができます。
サーバーを停止していつでも再実行できます。
要件: 約 10 GB の VRAM を備えた NVIDIA GPU、Python 3.11+、および CUDA
PyTorchのビルド。最初の起動ではモデルとレンズ (合計約 9 GB) をダウンロードし、
表示トークン マスクを構築します (約 1 分、キャッシュされます)。
git clone https://github.com/ninjahawk/Subtext
cd サブテキスト
pip install -r 要件.txt
Pythonサーバー.py
# → http://localhost:8765
Windows では、 python -u -X utf8 server.py を実行するか、 start.bat を使用します。
その他のモデル。 Neuronpedia のため、サーバーは Qwen3.5-4B 用に構成されています。
事前に取り付けられたレンズを公開しています (27B レンズも公開されています。
GPU — server.py の MODEL_NAME / LENS_FILE を編集します)。任意の HuggingFace デコーダ
jlens.fit() で独自のレンズをフィッティングすることで使用できます。 ~100 個のプロンプトが生成されます
使用可能なレンズであり、4B スケールのモデルの取り付けには 1 台で 1 時間程度かかります。
単一のコンシューマ GPU。リファレンスリポジトリを参照してください
詳細については。
リプレイ。 ⤓ セッション ボタンは、現在の会話をエクスポートします。
レンズ フレームが JSON ファイルとして含まれています。 ?replay=<file-url> を使用してアプリを開きます
ライブ ペーシングで再生するには、GPU は必要ありません。それはまさにそれです
ホストされたデモです。
機器はメソッドの制限を継承します。レンズは概念だけを読み取ります
単一の語彙トークンに対応します。マルチトークンの概念は目に見えない
あるいは断片的。これは、で特定されたワークスペースをほぼキャプチャします。
モデルの内部状態全体ではなく、紙とその下の層
適合範囲は観察されません。解釈も

紙を尊重する
独自のフレーミング: ワークスペースの読み出しは、機能の可用性を示します。
報告と推論のための情報。彼らは主観を証明しない
経験。
メソッドとリファレンス実装は Anthropic によるものです。
(jacobian-lens、Apache 2.0)。
事前に取り付けられたレンズ重量は次のサイトで公開されています。
ニューロンペディア。モデルは
クウェン3.5-4B 。サブテキストは
独立したプロジェクトであり、Anthropic とは関係ありません。
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

Contribute to ninjahawk/Subtext development by creating an account on GitHub.

GitHub - ninjahawk/Subtext · GitHub
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
ninjahawk
/
Subtext
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits docs docs media media .gitignore .gitignore LICENSE LICENSE README.md README.md index.html index.html record_session.py record_session.py requirements.txt requirements.txt server.py server.py start.bat start.bat test_client.py test_client.py test_multiturn.py test_multiturn.py verify_accuracy.py verify_accuracy.py View all files Repository files navigation
A real-time instrument for observing the verbal workspace of a language model
as it reads, reasons, and speaks.
🌐 Watch a live replay in your browser · ▶ Demo video · 📄 The paper · 🔬 Reference implementation
Recent work from Anthropic identified a small set of internal representations in
language models — the J-space — that behaves like a global workspace: its
contents can be verbally reported by the model, deliberately modulated, and are
causally used for multi-step reasoning, while the surrounding majority of neural
activity remains inaccessible to report. The identification tool is the
Jacobian lens, which transports a residual-stream activation at any layer into
the final-layer basis and decodes it through the model's own unembedding,
answering: which vocabulary words is this internal state disposed to produce,
now or later?
Subtext applies that method continuously during live conversation with a local
model. On every token — both while the model ingests the user's message and
while it generates its reply — the lens is read at nine depths and the result
is rendered as it happens. The intermediate steps of the model's computation
become directly watchable: verdicts form during reading, several tokens before
any output; planned words hold at high strength while unrelated tokens are
being emitted; two-hop questions surface their unspoken middle term.
Subtext differs from the interactive readouts already available (e.g. the
Neuronpedia demo) in that it is conversational and continuous: it renders the
lens during a live chat, includes the reading phase over the user's message,
streams at generation speed via a KV cache, and pairs the canvas with a
per-token ledger and per-word inspector. Sessions can be exported and replayed
in any browser without a GPU.
What the lens shows that the output does not
The value of the instrument is the gap between the model's internal state and
its visible text. Three moments from the demo session:
1. The verdict precedes the reply. Zero tokens of output exist; the model
is still reading Is this correct? 12 + 5 = 1 . The workspace already holds
math , addition , arithmetic , modulo — the phase indicator is amber
(reading).
2. The judgment is formed, then verbalized. As the reply begins ("No, that
is **not…"), incorrect dominates the workspace at high strength, with
equation , calculation , statement co-active — the conclusion is
internally settled several tokens before the words "not correct" appear.
3. Plans are held while other words are being said. Mid-explanation, the
workspace holds modulo , bitwise , system , numbers — the technical
caveat the model is about to raise — while the current output token is
unrelated.
These reproduce, on an open 4B model on consumer hardware, the reporting and
planning phenomena described in the paper (which used Claude-scale models),
including the two-hop signature: Italy at layer 20 and euros at layer 26
on the country-shaped-like-a-boot question, before generation begins.
Each rendered word is a lens readout, not model output. It indicates an
internal activation disposing the model toward that word.
Vertical position corresponds to layer. Early layers (perception) are at
the top; readouts approach the bottom rail as they approach emission.
Size and opacity encode absolute readout strength. The display applies a
fixed monotone mapping from lens probability; weak readouts are rendered
weak. Amber marks readouts taken while reading the user; blue while
generating.
Hover shows a word's per-layer activation profile; click opens an
inspector with peak strength, mean depth, and strength history.
The right panel records everything the canvas curates: the conversation, a
live ranking of currently-active readouts, and a per-token ledger.
browser (single HTML file) ⇐ websocket ⇐ server.py
Qwen3.5-4B (bf16, HF transformers, KV cache)
pre-fitted Jacobian lens: neuronpedia/jacobian-lens, revision qwen-n1000
per token: residual hooks at 9 layers → J_l transport → unembed
→ full-vocabulary softmax → word-start top-k → frame
Each exchange has two phases. A single prefill pass covers the user's message,
with lens readouts taken at every position (the reading phase); generation
then proceeds token-by-token with a KV cache, reading the lens at the newest
position each step (the thinking phase). The lens adds a per-layer
matrix-vector product and an unembedding per token, so streaming runs at the
model's native generation speed.
Display filtering. Raw lens top-k contains punctuation and BPE
continuation fragments (e.g. itude , from cert‑itude ), which are not
meaningful as readouts. Display is restricted to word-initial vocabulary
tokens, following the reference implementation's mask_display with a
stricter word-start criterion. Probabilities are computed over the full
vocabulary before any filtering, so filtering affects legibility only, never
the readout itself.
verify_accuracy.py compares this implementation's live path (forward hooks,
KV cache enabled) against the reference JacobianLens.apply() on identical
inputs. Across 4 layers × 3 positions on the walkthrough prompt, top-5
readouts match exactly, with cosine similarity ≥ 0.99998 between logit
vectors, and reproduce the expected two-hop intermediates. The audit can be
re-run at any time with the server stopped.
Requirements: an NVIDIA GPU with ~10 GB of VRAM, Python 3.11+, and a CUDA
build of PyTorch. First launch downloads the model and lens (~9 GB total) and
builds a display-token mask (~1 minute, cached).
git clone https://github.com/ninjahawk/Subtext
cd Subtext
pip install -r requirements.txt
python server.py
# → http://localhost:8765
On Windows, run python -u -X utf8 server.py , or use start.bat .
Other models. The server is configured for Qwen3.5-4B because Neuronpedia
publishes a pre-fitted lens for it (a 27B lens is also published, for larger
GPUs — edit MODEL_NAME / LENS_FILE in server.py ). Any HuggingFace decoder
can be used by fitting your own lens with jlens.fit() ; ~100 prompts produces
a usable lens, and fitting a 4B-scale model takes on the order of an hour on a
single consumer GPU. See the reference repo
for details.
Replays. The ⤓ session button exports the current conversation — every
lens frame included — as a JSON file. Open the app with ?replay=<file-url>
to play one back with live pacing, no GPU required; that is exactly what the
hosted demo is.
The instrument inherits the method's limitations. The lens reads only concepts
that correspond to single vocabulary tokens; multi-token concepts are invisible
or fragmentary. It approximately captures the workspace identified in the
paper, not the entirety of the model's internal state, and layers below the
fitted range are not observed. Interpretation should also respect the paper's
own framing: workspace readouts demonstrate functional availability of
information for report and reasoning; they do not demonstrate subjective
experience.
The method and reference implementation are by Anthropic
( jacobian-lens , Apache 2.0).
Pre-fitted lens weights are published by
Neuronpedia . The model is
Qwen3.5-4B . Subtext is an
independent project and is not affiliated with Anthropic.
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
