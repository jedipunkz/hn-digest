---
source: "https://github.com/shahar-dagan/openfusion"
hn_url: "https://news.ycombinator.com/item?id=48582749"
title: "Show HN: open source llm fusion"
article_title: "GitHub - shahar-dagan/openfusion: Combine the results from a panel of models into an enhanced response · GitHub"
author: "shadag"
captured_at: "2026-06-18T10:35:33Z"
capture_tool: "hn-digest"
hn_id: 48582749
score: 2
comments: 0
posted_at: "2026-06-18T09:00:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: open source llm fusion

- HN: [48582749](https://news.ycombinator.com/item?id=48582749)
- Source: [github.com](https://github.com/shahar-dagan/openfusion)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T09:00:35Z

## Translation

タイトル: HN を表示: オープンソース llm fusion
記事のタイトル: GitHub - shahar-dagan/openfusion: モデルのパネルからの結果を強化された応答に結合する · GitHub
説明: モデルのパネルからの結果を強化された応答に結合します - GitHub - shahar-dagan/openfusion: モデルのパネルからの結果を強化された応答に結合します
HN text: OpenRouter の融合は素晴らしいですね。私はこれを私自身の深い調査クエリのために構築しました。モデルの融合は、建設的な批判が必要なツール呼び出しや複雑なプロンプトに役立ちます。フィードバックをお待ちしております。

記事本文:
GitHub - shahar-dagan/openfusion: モデルのパネルからの結果を結合して強化された応答を生成する · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
シャハル・ダガン
/
オープンフュージョン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
92 コミット 92 コミット .github .github ベンチ ベンチ ドキュメント ドキュメントの例 例 openfusion openfusion スクリプト スクリプト テスト テスト Web Web .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースのドロップイン複合モデル プロキシ。 OpenAI 互換ツールをそれに向けて、
set model: "openfusion" と、プロンプトが LLM のパネルに並行して展開されます —
次に、裁判官モデルがすべての応答 (合意、矛盾、盲点) を読み取り、ストリームします。
それらのいずれかを倒すことを目的とした単一の合成された答えを返します。
これは、OpenRouter の Fusion の背後にあるエージェントの混合のアイデアのオープン バージョンです。より良い答えです。
ブラックボックスではなく、調整可能でフォーク可能なレシピとして、すでに支払ったモデルから。
クイックスタート · 仕組み · プレイグラウンド ·
ルーティングと戦略 · 対 OpenRouter Fusion ·
ベンチマーク · 貢献
ここは新しいですか？実行するには最初の 2 つだけが必要です。残りは調整と貢献です。
ベータ — パネルのファンアウト、ジャッジ合成、SSE ストリーミング、Web ツールの融合、自動ルーター、ディベート/
投票/ランク付けされたアグリゲータ、生産制限、インタラクティブなプレイグラウンド。 DESIGN.mdを参照してください。
アーキテクチャとセキュリティに関する注意事項については docs/ARCHITECTURE.md を参照してください。
openfusion には、対話型ターミナル チャットと Web プレイグラウンドという 2 つのフロント エンドがあります。クローンはありません、いいえ
config、開始するために環境変数は必要ありません。
uvx --from git+https://github.com/shahar-dagan/openfusion openfusion # 一時的、uv が必要
# …または: ピップインスタ

git+https://github.com/shahar-dagan/openfusion && openfusion
裸の openfusion は、モデル パネル (バナー、ライブ) とのリッチレンダリングされたチャットにあなたを導きます。
パネル進行状況スピナー、構文が強調表示されたコードによるマークダウン応答、およびスラッシュ コマンド
( /preset 、 /tokens 、 /models 、 /key 、 /clear )。最初の実行時に OpenRouter キーを要求され、
これを保存すると ( ~/.config/openfusion/credentials )、後で実行しても再プロンプトが表示されなくなります。 /key を使用して
それを変えてください。ワンショット用のパイプ: echo "…" |オープンフュージョン 。
openfusion web # ブラウザーでプレイグラウンドを開きます
# …または: docker run -p 8000:8000 ghcr.io/shahar-dagan/openfusion
openfusion web は、サーバーの準備が完了すると、http://localhost:8000 でプレイグラウンドを開きます (パス
--no-open 、または非対話型/ヘッドレス/Docker コンテキストでは自動的にスキップされます)。を貼り付けてください
キー (サーバーのメモリ内にのみ保持される) とヒューズ。何も設定しないと、Budget プリセットが起動します (
多様なパネル + ウェブ検索による審査) なので、最初の実行でフュージョンが実際に勝ちます。
コマンドをどこにでもインストールします (アクティブ化する venv はありません)
UVツールをインストールします。 # クローンから — または: pipx install 。 && pipx 保証パス
アクティブな開発の場合は、 pip install -e を実行します。アクティブ化された venv 内 (コマンドはその場合のみ機能します)
venv がアクティブな間)。裸の pip install -e 。 openfusion をグローバル PATH に配置しません —
「トラブルシューティング」を参照してください。
固定レシピの場合は、openfusion.yaml を作成します (examples/preset.yaml.example から開始します)。
プリセット: 品質 | Budget 、または完全に詳細に説明されたパネル/審査員の場合は、examples/default.yaml.example)。あ
プリセットは多様な OpenRouter パネルに拡張され、Web ツールをオンにして判定し、OpenRouter をミラーリングします
Fusion の品質/予算スイッチ:
OpenAI SDK からドロップイン API として使用します (openfusion Web を実行している場合)。
openaiインポートからOpenAI
client = OpenAI (base_url = "http://localhost:8

000/v1" 、api_key = "local-dev" )
ストリーム = クライアント 。チャット。完成品。作成(
モデル = "オープンフュージョン" 、
messages = [{ "role" : "user" , "content" : "エージェントの混合を 1 つの段落で説明します。" }]、
ストリーム = True 、
)
ストリーム内のチャンクの場合:
print (チャンク . 選択肢 [ 0 ]. デルタ . コンテンツまたは "" 、 end = "" )
またはターミナルから直接、サーバーは必要ありません。
openfusion ask " 小規模 SaaS の Postgres と SQLite を比較します。 " --max-tokens 800
ask は、構成されたパネルに対して 1 つのフュージョンを実行し、合成された回答を stdout にストリーミングします。
(パネルの進行状況は標準エラー出力に送られます)。 --max-tokens はすべての呼び出しに上限を設定します。値が低いほど、高速かつ安価になります。
スピードと長さ。 Fusion は N パネル コールとジャッジを実行するため、1 つのモデルよりも遅くなります。
パネルは並行して進行し、パネルが終了するとすぐに審査員がストリーミングします。裁判官に促される
簡潔さを保つために、 --max-tokens (CLI)、max_tokens (API)、応答で長さを制限します。
プレイグラウンド設定内の長さコントロール、または設定内のcost_controls。
3 つのノブは、プロンプトを融合するかどうか、および融合する方法を制御します。すべてオプションで、オフ/デフォルトです。
Auto Router ( router.enabled: true ) — 単純なプロンプトに応答するプロンプトごとのゲート
単一のパススルー呼び出しで、メリットがありそうなプロンプト用にパネルを予約します (長い、
分析的、またはコードを含む)。デフォルトは安価なヒューリスティックです (追加のモデル呼び出しはありません)。モード: モデル
小規模な分類子モデルを使用し、エラーが発生した場合はヒューリスティックに戻ります。
ルータ：
有効 : true
モード : ヒューリスティック # ヒューリスティック |モデル |いつも |決して
min_chars : 280 # この長さ以上のプロンプトがヒューズします
# 分類子: # モード: モデルに必要
# ベース URL: https://openrouter.ai/api/v1
# API_key: ${OPENROUTER_API_KEY}
＃モデル：openai/gpt-4o-mini
Strategy ( Strategy: ) — パネルの生成方法: self_fusion (1 つのモデルを N 回サンプリング)、

パネル（固定された多様なパネル）またはディベート（各メンバーが後で修正する多様なパネル）
他の人の答えを見て、裁判官が総合します)。議論は余分なコストやレイテンシと引き換えに、
反対尋問:
戦略: 議論
討論：
ラウンド: ジャッジの前で 1 # 回の修正ラウンド
アグリゲータ ( アグリゲータ: ) — 回答が 1 つになる方法: 判断 (合成、デフォルト)、投票
(多数決、安価、検証可能な短答タスクに最適)、またはランク付け (1 人の短答審査員)
呼び出しは単一の最良の回答を選択します。合成よりも安価で、投票とは異なりモデル判断を使用します)。
分析の透明性 ( Analysis.emit: true ) — 裁判官の構造化された推論を表面化します
(コンセンサス / 矛盾 / 部分的な網羅 / 独自の洞察 / 盲点) を別個の SSE として
イベント: 分析 (および非ストリーミング応答の分析フィールド)。
回答本文。
プロンプト キャッシング (cache.enabled: true ) — 自己融合の N サンプルになるように共有プレフィックスをマークします。
キャッシュされたプロンプトをサポートするプロバイダーで再利用します (他の場所では操作は不要です)。
パブリック デプロイメントの場合、制限された負荷と使用量 (どちらもデフォルトは 0 = 無制限):
制限:
max_in_flight : 64 # 同時リクエストの上限。制限を超えた場合の返品 503
rate_limit_per_ minutes : 60 # ゲートウェイ キーごと (または認証されていない場合はクライアントごと);制限を超えた場合の返品 429
これらはベストエフォート型の単一プロセス ガードです。プロバイダー側の予算と組み合わせて、
マルチレプリカ展開、エッジ レート リミッター。
モデルへのリクエスト: 「openfusion」は、モデルのパネルに並行してファンアウトされます (各モデルはオプションです)
独自のウェブ調査を行っています)、その後、審査員モデルがすべての回答を読み取り、1 つを合成します — ストリーミング
SSE に戻り、構造化分析とコストを並行して行います。
フローチャート LR
C["クライアント<br/>(カーソル · OpenAI SDK · 何でも)"] -->|"POST /v1/chat/completions<br/>model=openfusion"|ろ

後<br/><i>(オプション)</i>"}
R -->|簡単なプロンプト| S["単一モデル"] --> OUT
R -->|融合する価値がある| P
サブグラフ P [「パネル・並列ファンアウト」]
TB方向
A["モデル A 🔍"]
B["モデル B 🔍"]
D["モデル C 🔍"]
終わり
P --> J[「裁判官<br/>合意・矛盾・盲点」]
J --> OUT["ストリーミングされた回答 (SSE)<br/>+ 分析 + トークン/コスト"]
C -.->|その他のモデル/クライアント ツール| S
classDef アクセント塗りつぶし:#eef2ff、ストローク:#4f46e5、色:#3730a3;
クラスJ、Rアクセント。
読み込み中
ドロップイン。 OpenAI 互換の POST /v1/chat/completions + /v1/models 、実際の SSE ストリーミング。
ロックインはありません。各パネルメンバーと審査員は {base_url, api_key, model} です。 OpenRouter は、
デフォルトのアップストリーム。 OpenAI、Togetter、ローカル vLLM/Ollama はすべて機能します。
構成主導型。パネル、ジャッジ、ストラテジー、アグリゲーター、ルーター、制限は openfusion.yaml にあります
— または 1 単語のプリセット、またはまったく何もしない (ゼロ構成クイック スタート)。
openfusion と OpenRouter Fusion の比較
openfusion は、同じアイデアをオープンに実装したものです。コアのメカニズムは同等です。の
違いは、スケールとプロンプトごとのルーターです。
パラメータ
に適用されます
注意事項
温度（クライアント）
レシピを通じて間接的にのみ判断する
自己融合はクライアントではなく設定によってパネル温度を変化させます
max_tokens 、 stop 、 response_format
ジャッジ (可視出力)
パネルメンバーはレシピのデフォルトを使用します
ストリーム 、 stream_options
ジャッジパス
パネルは常に内部で非ストリーミングで実行されます
ツール / ツールコール
フュージョンまたはパススルー
サーバー実行可能な Web ツール ( openrouter:web_search / web_fetch ) が融合されています。クライアント側の機能ツールと会話中のツールのターンパススルー
環境変数
変数
目的
OPENROUTER_API_KEY
デフォルトのアップストリーム キー (構成内の ${OPENROUTER_API_KEY} 経由)
OPENFUSION_CONFIG
構成ファイルへのパス (デフォルト: openfusion.yaml )
OPENFUSION_API_KEYS
カンマ区切りのゲートウェイ許可リスト (オプション)
OPENFUSION_HOST /

OPENFUSION_PORT
サーバーバインドアドレス
コスト安全性と生煙試験
設定のcost_controlsは、パススルー、パネル、およびジャッジ呼び出しのmax_tokensを制限します。行方不明
max_tokens 値は、設定された上限から埋められます。オーバーリミットパススルーとジャッジ
リクエストは 400 を返しますが、内部パネルの呼び出しは上限に固定されます。
オプトイン ライブ OpenRouter スモーク テストは、少量のクレジットを使用する場合にのみ実行してください。
エクスポート OPENROUTER_API_KEY=あなたのキー
python scripts/openrouter_smoke.py --config example/dev.yaml.example --yes-spend-credits
ベンチマーク
直接ベンチマークを実行します (自己融合モデルとソロ モデル)。
pip install -e " .[dev] "
python ベンチ/run.py --config サンプル/default.yaml.example --tasks ベンチ/tasks/sample.jsonl
より大きなベンチマークを実行する前に、 --tasks bench/tasks/smoke.jsonl --max-tokens 32 を使用してください。
各実行では、精度とそこに到達するまでに要した費用 (total_tokens と ) が報告されます。
モードごとの total_cost_usd — 精度の変化と扇風機の追加コストを比較検討できます。
パネルに出ます。
バンドルされた bench/tasks/sample.jsonl (20 の短い Q&A タスク) は、有能なモデルとしては飽和状態です。
ソロのベースラインはすでに ~100% のスコアを獲得しているため、フュージョンで精度を高めるためのヘッドルームはありません。で
openai/gpt-4o-mini で最近実行した (自己融合 N=2、max_tokens=32):
それで

[切り捨てられた]

## Original Extract

Combine the results from a panel of models into an enhanced response - GitHub - shahar-dagan/openfusion: Combine the results from a panel of models into an enhanced response

OpenRouter's fusion is great. I built this for my own deep research queries. Model fusion is useful for tool calls and complex prompts that require constructive criticism. Would love feedback.

GitHub - shahar-dagan/openfusion: Combine the results from a panel of models into an enhanced response · GitHub
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
shahar-dagan
/
openfusion
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
92 Commits 92 Commits .github .github bench bench docs docs examples examples openfusion openfusion scripts scripts tests tests web web .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
An open-source, drop-in compound-model proxy. Point any OpenAI-compatible tool at it,
set model: "openfusion" , and your prompt is fanned out to a panel of LLMs in parallel —
then a judge model reads every response (consensus, contradictions, blind spots) and streams
back a single synthesized answer that aims to beat any one of them.
It's the open version of the mixture-of-agents idea behind OpenRouter's Fusion: better answers
from models you already pay for, as a tunable, forkable recipe instead of a black box.
Quick start · How it works · Playground ·
Routing & strategies · vs. OpenRouter Fusion ·
Benchmarks · Contributing
New here? You only need the first two to run it; the rest is for tuning and contributing.
Beta — panel fan-out, judge synthesis, SSE streaming, web-tool fusion, an Auto Router, debate/
vote/ranked aggregators, production limits, and an interactive playground. See DESIGN.md
and docs/ARCHITECTURE.md for architecture and security notes.
openfusion has two front ends — an interactive terminal chat and a web playground. No clone, no
config, no env vars needed to start.
uvx --from git+https://github.com/shahar-dagan/openfusion openfusion # ephemeral, needs uv
# …or: pip install git+https://github.com/shahar-dagan/openfusion && openfusion
Bare openfusion drops you into a Rich-rendered chat with the model panel — a banner, a live
panel-progress spinner, Markdown answers with syntax-highlighted code, and slash commands
( /preset , /tokens , /models , /key , /clear ). On first run it asks for your OpenRouter key and
saves it ( ~/.config/openfusion/credentials ), so later runs don't re-prompt; use /key to
change it. Pipe for one-shots: echo "…" | openfusion .
openfusion web # opens the playground in your browser
# …or: docker run -p 8000:8000 ghcr.io/shahar-dagan/openfusion
openfusion web pops the playground open at http://localhost:8000 once the server is ready (pass
--no-open , or it's skipped automatically in non-interactive/headless/Docker contexts). Paste your
key (kept only in server memory) and fuse. With nothing configured it boots the Budget preset (a
diverse panel + judge with web search) so the first run lands where fusion actually wins.
Install the command everywhere (no venv to activate)
uv tool install . # from a clone — or: pipx install . && pipx ensurepath
For active development, pip install -e . inside an activated venv (the command then works only
while that venv is active). A bare pip install -e . does not put openfusion on your global PATH —
see Troubleshooting .
For a fixed recipe, write an openfusion.yaml (start from examples/preset.yaml.example —
preset: quality | budget , or examples/default.yaml.example for a fully spelled-out panel/judge). A
preset expands to a diverse OpenRouter panel + judge with web tools on, mirroring OpenRouter
Fusion's Quality/Budget switch:
Use as a drop-in API from the OpenAI SDK (with openfusion web running):
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8000/v1" , api_key = "local-dev" )
stream = client . chat . completions . create (
model = "openfusion" ,
messages = [{ "role" : "user" , "content" : "Explain mixture-of-agents in one paragraph." }],
stream = True ,
)
for chunk in stream :
print ( chunk . choices [ 0 ]. delta . content or "" , end = "" )
Or straight from the terminal, no server needed:
openfusion ask " Compare Postgres and SQLite for a small SaaS. " --max-tokens 800
ask runs one fusion against your configured panel and streams the synthesized answer to stdout
(panel progress goes to stderr). --max-tokens caps every call — lower is faster and cheaper.
Speed & length. Fusion runs N panel calls plus a judge, so it's slower than one model — the
panel runs in parallel and the judge streams as soon as the panel finishes. The judge is prompted
to stay concise, and you cap length with --max-tokens (CLI), max_tokens (API), the response-
length control in the playground Settings, or cost_controls in config.
Three knobs control whether and how a prompt is fused. All are optional and off/default.
Auto Router ( router.enabled: true ) — a per-prompt gate that answers simple prompts with a
single pass-through call and reserves the panel for prompts that look like they benefit (long,
analytical, or containing code). Default is a cheap heuristic (no extra model call); mode: model
uses a small classifier model and falls back to the heuristic if it errors:
router :
enabled : true
mode : heuristic # heuristic | model | always | never
min_chars : 280 # prompts at/over this length fuse
# classifier: # required for mode: model
# base_url: https://openrouter.ai/api/v1
# api_key: ${OPENROUTER_API_KEY}
# model: openai/gpt-4o-mini
Strategy ( strategy: ) — how the panel is produced: self_fusion (one model sampled N times),
panel (a fixed diverse panel), or debate (a diverse panel where each member revises after
seeing the others' answers, then the judge synthesizes). Debate trades extra cost/latency for
cross-examination:
strategy : debate
debate :
rounds : 1 # revision rounds before the judge
Aggregator ( aggregator: ) — how answers become one: judge (synthesis, default), vote
(majority vote, cheaper, best for verifiable short-answer tasks), or ranked (one short judge
call picks the single best answer — cheaper than synthesis, uses model judgment unlike vote).
Analysis transparency ( analysis.emit: true ) — surface the judge's structured reasoning
(consensus / contradictions / partial coverage / unique insights / blind spots) as a separate SSE
event: analysis (and an analysis field on non-streaming responses), without polluting the
answer body.
Prompt caching ( cache.enabled: true ) — mark the shared prefix so self-fusion's N samples
reuse a cached prompt on providers that support it (a no-op elsewhere).
For public deployments, bound load and spend (both default to 0 = unlimited):
limits :
max_in_flight : 64 # cap concurrent requests; over-limit returns 503
rate_limit_per_minute : 60 # per gateway key (or per client when unauthenticated); over-limit returns 429
These are best-effort, single-process guards — pair them with provider-side budgets and, for
multi-replica deployments, an edge rate limiter.
A request to model: "openfusion" is fanned out to a panel of models in parallel (each optionally
doing its own web research), then a judge model reads every answer and synthesizes one — streamed
back over SSE, with the structured analysis and cost alongside.
flowchart LR
C["Client<br/>(Cursor · OpenAI SDK · anything)"] -->|"POST /v1/chat/completions<br/>model=openfusion"| R{"Router<br/><i>(optional)</i>"}
R -->|simple prompt| S["Single model"] --> OUT
R -->|worth fusing| P
subgraph P ["Panel · parallel fan-out"]
direction TB
A["Model A 🔍"]
B["Model B 🔍"]
D["Model C 🔍"]
end
P --> J["Judge<br/>consensus · contradictions · blind spots"]
J --> OUT["Streamed answer (SSE)<br/>+ analysis + token/cost"]
C -.->|other model / client tools| S
classDef accent fill:#eef2ff,stroke:#4f46e5,color:#3730a3;
class J,R accent;
Loading
Drop-in. OpenAI-compatible POST /v1/chat/completions + /v1/models , real SSE streaming.
No lock-in. Each panel member + judge is {base_url, api_key, model} . OpenRouter is the
default upstream; OpenAI, Together, local vLLM/Ollama all work.
Config-driven. Panel, judge, strategy, aggregator, router, and limits live in openfusion.yaml
— or a one-word preset , or nothing at all (zero-config quick start).
openfusion vs. OpenRouter Fusion
openfusion is the open implementation of the same idea. The core mechanism is at parity; the
differences are scale and a per-prompt router.
Parameter
Applies to
Notes
temperature (client)
Judge only indirectly via recipe
Self-fusion varies panel temps from config, not client
max_tokens , stop , response_format
Judge (visible output)
Panel members use recipe defaults
stream , stream_options
Judge path
Panel always runs non-streamed internally
tools / tool_calls
Fusion or pass-through
Server-executable web tools ( openrouter:web_search / web_fetch ) are fused; client-side function tools and mid-conversation tool turns pass through
Environment variables
Variable
Purpose
OPENROUTER_API_KEY
Default upstream key (via ${OPENROUTER_API_KEY} in config)
OPENFUSION_CONFIG
Path to config file (default: openfusion.yaml )
OPENFUSION_API_KEYS
Comma-separated gateway allowlist (optional)
OPENFUSION_HOST / OPENFUSION_PORT
Server bind address
Cost safety and live smoke tests
cost_controls in config caps max_tokens for pass-through, panel, and judge calls. Missing
max_tokens values are filled from the configured ceiling; over-limit pass-through and judge
requests return 400 , while internal panel calls clamp to their ceiling.
Run the opt-in live OpenRouter smoke test only when you intend to spend a small number of credits:
export OPENROUTER_API_KEY=your-key
python scripts/openrouter_smoke.py --config examples/dev.yaml.example --yes-spend-credits
Benchmarks
Run the head-to-head benchmark (self-fusion vs solo model):
pip install -e " .[dev] "
python bench/run.py --config examples/default.yaml.example --tasks bench/tasks/sample.jsonl
Use --tasks bench/tasks/smoke.jsonl --max-tokens 32 before larger benchmark runs.
Each run reports accuracy plus the spend it took to get there — total_tokens and
total_cost_usd per mode — so you can weigh any accuracy change against the extra cost of fanning
out to a panel.
The bundled bench/tasks/sample.jsonl (20 short Q&A tasks) is saturated for a capable model —
the solo baseline already scores ~100%, so there is no headroom for fusion to add accuracy. On a
recent run with openai/gpt-4o-mini (self-fusion N=2, max_tokens=32 ):
So

[truncated]
