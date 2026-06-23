---
source: "https://github.com/itsthelore/wayfinder-router"
hn_url: "https://news.ycombinator.com/item?id=48648150"
title: "Wayfinder – routing LLM prompts without another LLM"
article_title: "GitHub - itsthelore/wayfinder-router: Simple CLI tool for deterministic routing of queries between local and hosted LLM models · GitHub"
author: "tcballard"
captured_at: "2026-06-23T17:34:58Z"
capture_tool: "hn-digest"
hn_id: 48648150
score: 1
comments: 0
posted_at: "2026-06-23T17:17:02Z"
tags:
  - hacker-news
  - translated
---

# Wayfinder – routing LLM prompts without another LLM

- HN: [48648150](https://news.ycombinator.com/item?id=48648150)
- Source: [github.com](https://github.com/itsthelore/wayfinder-router)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T17:17:02Z

## Translation

タイトル: Wayfinder – 別の LLM を使用せずに LLM プロンプトをルーティングする
記事のタイトル: GitHub - itsthelore/wayfinder-router: ローカル LLM モデルとホストされた LLM モデル間のクエリの決定論的ルーティングのためのシンプルな CLI ツール · GitHub
説明: ローカル LLM モデルとホストされた LLM モデルの間でクエリを決定的にルーティングするためのシンプルな CLI ツール - itsthelore/wayfinder-router

記事本文:
GitHub - itsthelore/wayfinder-router: ローカル LLM モデルとホストされた LLM モデル間のクエリの決定論的ルーティングのためのシンプルな CLI ツール · GitHub
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
それはそれで
/
ウェイファインダールーター
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
163 コミット 163 コミット .claude/ skill/ wayfinder-triage .claude/ skill/ wayfinder-triage .github/ workflows .github/ workflows .rac .rac ベンチマーク ベンチマーク 決定 意思決定 デザイン デザイン ドキュメント ドキュメントの例 例 ロードマップ ロードマップ テスト テスト wayfinder_router wayfinder_router .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile EXPLAINER.md EXPLAINER.md ライセンス ライセンス Makefile Makefile 通知 通知 README.md README.md conftest.py conftest.py docker-compose.example.yml docker-compose.example.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
決定的なプロンプトの複雑さのルーティング — 各プロンプトを
ローカルまたはクラウド モデル、オフライン、モデル呼び出しを決定する必要はありません。
クイックスタート ·
ベンチマーク ·
比較する方法 ·
説明者・
変更履歴
Wayfinder はプロンプトの形状 (長さ、見出し、リスト、コード) を読み取ります。
加えて、証明、数学、厳しい制約などの文言の難しさの手がかり、および
小規模なローカル モデルに送信するか、大規模なクラウド モデルに送信するかを指定します。それ
マイクロ秒で決定し、オフラインで実行され、別のモデルを呼び出して作成することはありません。
電話する。 API キーもネットワークもモデル呼び出しも必要ありません。スコアと
推奨事項。それをどうするかはあなた次第です。
安価なプロンプトはローカルに残り、難しいプロンプトは高価なモデルに移り、あなたは支払うのをやめます
「これを要約してください」と「私のタイプミスを修正してください」というフロンティア価格。
ほとんどのルーターは、トレーニングされた分類子、LLM 判定者、またはモデルを呼び出すことによって決定します。
ホストされた API。これにより、正確なステップにレイテンシー、コスト、およびわずかなランダム性が追加されます。
それはお金を節約することを目的としています。 Wayfinder は代わりに構造と文言を読み取るため、
決定は自由であり、いつでも同じです。
W

ayfinder は最高の精度数値を追求していません。実行できるルーターは 1 つだけです
オフラインでモデル呼び出しを行わず、独自のトラフィックに合わせて調整します。デフォルトでは得点が入ります
プロンプト構造のみ。語彙上の手がかり (証明、数学、制約) を読み取ることもできます。
ただし、それらはデフォルトで出荷されます: 二重盲検テスト
独立して作成されたプロンプトでは、語彙リフトが一般化しないことが示されました (
目に見えない難しいプロンプトの最大 20% をキャッチしますが、単純な単語数ベースラインには負けます)。
これらはオプトインです - 独自の重みに調整した場合にのみ重みを上げます
交通用語。難易度が純粋に意味論的なプロンプト — 微妙なコード
「100番目の素数は何ですか?」という無邪気そうなスニペット。 — 構造的なものはありません
そうすれば、セマンティックルーターがそこでそれを打ち負かします。盲目でも生き残るエッジ
テストは、決定論的、ミリ秒未満のオフライン ルーティングをリードするものです。
モデルコールなしの決定。ベンチマーク (ベンチマークを作成)
正直なベースラインと完璧な基準に対して、どこで勝ち、どこで負けるかを示します。
オラクル。 RouterBench または RouterArena にポイントすると、段階的な数値が表示されます。
ここは初めてですか、それとも検討中ですか? FAQ には明確な答えがあります —
どこで負けるかも含めて（RouterBench の短いけど難しいランダムな結果と変わらない）
items)、そしてなぜそれをまだ実行するのか。
ルーティングの決定を自分で確認するには 2 つの方法があります。API キーもモデルもネットワーク上に何も存在しません。
ターミナル内 — Wayfinder パレットでの意思決定優先のチャット。ターミナル
チャットはデフォルトのインストールに同梱されているため、追加するものは何もありません。または、次のように実行します。
uvx 経由ではまったくインストールされません:
uvx wayfinder-router chat --dry-run # ゼロインストール、ゼロキー
# または: pip install wayfinder-router && wayfinder-router chat
すべてのターンには、ルーティングされた場所 ( ● LOCAL / ◆ CLOUD )、構造スコア、およびその理由が表示されます。
( /why )、そしてランニング

節約と常時クラウドの比較。 /init は終了せずにモデルをセットアップします
チャット、/route · /local · /cloud はターンを強制し、会話は複数の場所で継続します。
セッション ( /threads )。
ブラウザーでは、ライブしきい値スライダーを備えた Web チャット UI:
pip install " wayfinder-router[ゲートウェイ] "
ウェイファインダールーターウェブチャット --dry-run
# http://127.0.0.1:8088/demo を開きます
webchat は、サーブ上のシン ランチャー (ゲートウェイとその /demo ページ; --no-open 、
--port 、 --host 0.0.0.0 、 --dry-run );サーブはヘッドレスコマンドです。両面
すべてのメッセージについて、ルーティング先 (ローカルまたはクラウド)、複雑さのスコアとその理由を表示します。
(機能の内訳)、および常時クラウドと比較して節約されたコスト。設定がない場合は両方とも
決定のみ (Web の場合は --dry-run、ターミナルのプレビュー) なので、次のように実行できます。
ゼロセットアップ。実際の応答を取得するには、wayfinder-router init を実行して [gateway.models] をスキャフォールディングします。
(その後、wayfinder-router ドクターがキーの解決を確認します) — クイックスタート を参照してください。
あらゆる OpenAI 互換 API で動作します
Wayfinder は、各呼び出しを OpenAI スタイルの /chat/completions エンドポイントに転送します。
あなたのプロバイダーはそう言っています（そしてほとんどがそうしています）、それはうまくいきます。階層は 1 つの Base_url です。
モデル名と、リクエスト時に環境から読み取られたキー。 SDK なし、なし
プロバイダーごとのコード。無料のローカル モデルとホストされたモデルをペアにするか、2 つのクラウド層を実行します。
…さらに、Groq、Togetter、OpenRouter、Fireworks、DeepSeek、およびローカル サーバー
(vLLM、LM Studio、llama.cpp) — + OpenAI 互換エンドポイント
これには Bearer キーが必要です。
Wayfinder をモデルの前に置きます。アプリは OpenAI API を話し続けます。あなた
Base_url を 1 つ変更するだけです。
構成をスキャフォールディングする — init はスターター wayfinder-router.toml (キーレス ローカル
Ollama → Anthropic Cloud) と .env.example を追加して、キーを確認します。
pip install " wayfinder-router[ゲートウェイ] "
ウェイファインダー-ro

uter init # スターター設定 (ハイブリッド プリセット)
wayfinder-router init --preset openai # 2 つの OpenAI 層 (gpt-4o-mini → gpt-4o)
wayfinder-router init --preset gemini # 2 つの Gemini 層 (gemini-2.5-flash → gemini-2.5-pro)
wayfinder-router init --interactive # プロバイダー/モデルを段階的に選択します
または、wayfinder-router.toml に 2 つのモデルを手動で記述します。
[ルーティング]
しきい値 = 0.5 # 以下 -> ローカル、上/上 -> クラウド
[ゲートウェイ。モデル。地元]
Base_url = " http://localhost:11434/v1 "
モデル = " ラマ3.2 "
[ゲートウェイ。モデル。クラウド]
Base_url = " https://api.openai.com/v1 "
モデル = " gpt-4o "
api_key_env = " OPENAI_API_KEY " # この環境変数から読み取り、保存されません
# api_key_cmd = "op read op://Private/OpenAI/credential" # オプション: ボールトから入力します
Wayfinder はシークレットを保存しません。モデルは環境変数 ( api_key_env ) とキーに名前を付けます。
リクエスト時に環境から読み取られます。 「インストール」するものは何もありません。
変数をエクスポートします。生のキーをシェルに貼り付けたくないですか?オプションを追加する
api_key_cmd と Wayfinder は起動時にシークレット ストアからその変数を入力します。
操作読み取り … (1Password)、セキュリティ … (macOS キーチェーン)、秘密ツール … (Linux)、
pass / gopass 、vault kv get … 、aws Secretsmanager get-secret-value … 、 bw 、
doppler 、 gcloud secrets … 、またはシークレットを出力する任意のコマンド。鍵は握られている
メモリ内にのみ存在し、ディスクには書き込まれません。ウェイファインダールータードクターがどれを検出するか
インストールされているこれらのツールを確認し、正確な行を提案します。
キーを設定して、ゲートウェイを実行します。医師は設定を再チェックし、それぞれの設定が
開始する前に、モデルのキーが解決されます ( ✓ set / ✗ not set )。
import ANTHROPIC_API_KEY=sk-... # または OPENAI_API_KEY (設定ごと)
wayfinder-router Doctor モデルごとに # ✓/✗ — 各キーは設定されていますか?
ウェイファインダールーターサービス

--ポート 8088
既存のクライアントにそれを指示してください。コード変更なし:
クライアント = openai 。 OpenAI (base_url = "http://localhost:8088/v1" 、api_key = "未使用")
クライアント。チャット 。完成品。 create (モデル = "自動" , メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : "..." }])
簡単なプロンプトはローカルに、難しいプロンプトはクラウドに、そしてすべての応答にはメッセージが含まれます
x-wayfinder-router-model と x-wayfinder-router-score により、どこにあるかがわかります
行きました。 1 つのリクエストを操作したいですか? model="cloud" /prefer-local で固定するか、または
X-Wayfinder-Threshold ヘッダーを使用して単一の呼び出しのカットを移動します (「
単一のリクエストを操作します)。
カール -s ローカルホスト:8088/healthz
# {"ステータス":"ok","モデル":["クラウド","ローカル"]}
curl -s -D - -o /dev/null http://localhost:8088/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {"モデル":"自動","メッセージ":[{"役割":"ユーザー","コンテンツ":"こんにちは"}]} ' \
| grep -i x-wayfinder-router
# x-wayfinder-router-model: ローカル
# x-wayfinder-router-score: 0.00
バックエンドはまだありませんか? wayfinder-routerserve --dry-run ルーティングに関する応答
アップストリームに電話する代わりに決定を行うため、30 秒でルーティングを実感できます。
実際のモデルを配線する前に。
コマンド
あなたが得るもの
pip インストールウェイファインダールーター
スコアラー、CLI、Python API、ターミナル チャット ( chat );スコアラー/ライブラリのインポートは依存関係が軽いままです
pip install "wayfinder-router[ゲートウェイ]"
OpenAI 互換のルーティング ゲートウェイを追加します。これは、サービスの一般的なケースです。
pip install "wayfinder-router[ui]"
ローカルの調整/説明/構成 UI を追加します
pip install "wayfinder-router[all]"
デフォルトのインストールに加えてゲートウェイと UI
仕組み
Wayfinder は、すでに使用している OpenAI 互換クライアントの背後にあります。あなたは指摘します
そのクライアントのbase_urlはゲートウェイで一度だけ表示され、それ以降は表示されなくなります。の
ローカルまたはホストをルーティングするかどうかに関係なく、同じクライアントがリクエストを処理します

d.
クライアント (チャット アプリ、IDE、エージェント、またはコード)
|
v
Wayfinder ゲートウェイのスコア、モデルの選択
|
|-- 低 --> ローカル (Ollama、vLLM)
|-- 高 --> ホスト (OpenAI、任意の /v1)
|
v
応答は同じクライアント経由で返されます。
x-wayfinder-router-* ヘッダー付き
これからいくつかのことがわかります。
手前のインターフェースはあなたのものです。チャット GUI (Open WebUI、LibreChat)、IDE
カスタム エンドポイント (カーソル、続行)、エージェント フレームワーク、または
OpenAI SDK 上の独自のコード。今日チャットウィンドウをご希望ですか? Open WebUIを前面に置き、
それをゲートウェイに向けます。
ローカルおよびホストされたものはバックエンドであり、アプリではありません。ローカルモデルは単なるサーバーです
(Ollama、LM Studio、vLLM、llama.cpp) OpenAI の /v1 を話します。ホストされているものは
同じ形。ユーザーは UI を切り替えることはなく、通常はどのモデルかを知りません。
と答えた。
スコアはセカンドオピニオンではなく計算されます。モデルにどれくらい大変かを尋ねると、
プロンプトは遅く、非決定的であり、決定するにはモデル呼び出しが必要になります。
モデルコールをするかどうか。 Wayfinder は代わりにプロンプトをスキャンします — 構造
(長さ、見出し、ステップ、リンク、コード、表) および文言の難しさの手がかり
(推論用語、数学記号、制約) — 0.0 ～ 1.0 の値に変換し、
それをあなたの閾値と比較します。同じプロンプト、同じしきい値、同じ答え。それは
判決ではなく、困難の代理です。

[切り捨てられた]

## Original Extract

Simple CLI tool for deterministic routing of queries between local and hosted LLM models - itsthelore/wayfinder-router

GitHub - itsthelore/wayfinder-router: Simple CLI tool for deterministic routing of queries between local and hosted LLM models · GitHub
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
itsthelore
/
wayfinder-router
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
163 Commits 163 Commits .claude/ skills/ wayfinder-triage .claude/ skills/ wayfinder-triage .github/ workflows .github/ workflows .rac .rac benchmarks benchmarks decisions decisions designs designs docs docs examples examples roadmaps roadmaps tests tests wayfinder_router wayfinder_router .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile EXPLAINER.md EXPLAINER.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md conftest.py conftest.py docker-compose.example.yml docker-compose.example.yml pyproject.toml pyproject.toml View all files Repository files navigation
Deterministic prompt-complexity routing — send each prompt to your
local or cloud model, offline, with no model call to decide.
Quickstart ·
Benchmark ·
How it compares ·
Explainer ·
Changelog
Wayfinder reads the shape of a prompt — its length, headings, lists, and code —
plus difficulty cues in the wording, like proofs, math, and hard constraints, and
tells you whether to send it to your small local model or your big cloud one. It
decides in microseconds, runs offline, and never calls another model to make the
call. No API key, no network, no model call to decide. You get a score and a
recommendation; what you do with it is up to you.
Cheap prompts stay local, hard ones go to the expensive model, and you stop paying
frontier prices for "summarize this" and "fix my typo."
Most routers decide by calling a model: a trained classifier, an LLM judge, or a
hosted API. That adds latency, cost, and a little randomness to the exact step
that is meant to save you money. Wayfinder reads structure and wording instead, so
the decision is free and the same every time.
Wayfinder is not chasing a top accuracy number. It is the one router you can run
offline, with zero model calls, and tune on your own traffic. By default it scores
prompt structure only. It can also read lexical cues (proofs, math, constraints),
but those ship off by default : a double-blind test
on independently-authored prompts showed the lexical lift does not generalize (it
catches ~20% of unseen hard prompts and loses to a plain word-count baseline), so
they are opt-in — raise their weights only if you've calibrated them to your own
traffic's vocabulary. A prompt whose difficulty is purely semantic — a subtle code
snippet, an innocent-looking "what is the 100th prime number?" — has no structural
tell, and a semantic router will beat it there. The edge that survives the blind
test is the one to lead with: a deterministic, sub-millisecond, offline routing
decision with no model call. The benchmark ( make benchmark )
shows where it wins and where it loses, against honest baselines and a perfect
oracle. Point it at RouterBench or RouterArena for graded numbers.
New here, or weighing it up? The FAQ gives straight answers —
including where it loses (it's no better than random on RouterBench's short-but-hard
items) and why you'd still run it.
Two ways to see the routing decision for yourself — no API keys, no models, nothing on the network.
In your terminal — a decision-first chat in the Wayfinder palette. The terminal
chat ships in the default install, so there's nothing extra to add — or run it with
no install at all via uvx :
uvx wayfinder-router chat --dry-run # zero install, zero keys
# or: pip install wayfinder-router && wayfinder-router chat
Every turn shows where it routed ( ● LOCAL / ◆ CLOUD ), the structural score and why
( /why ), and the running savings vs always-cloud. /init sets up models without leaving
the chat, /route · /local · /cloud force a turn, and conversations persist across
sessions ( /threads ).
In your browser — the web chat UI with a live threshold slider:
pip install " wayfinder-router[gateway] "
wayfinder-router webchat --dry-run
# opens http://127.0.0.1:8088/demo
webchat is a thin launcher over serve (the gateway and its /demo page; --no-open ,
--port , --host 0.0.0.0 , --dry-run ); serve is the headless command. Both surfaces
show, for every message, where it routed (local vs cloud), the complexity score and why
(the feature breakdown), and the cost saved vs always-cloud. With no config both are
decision-only ( --dry-run for the web; the terminal's preview), so you can poke at it with
zero setup. To get real replies, run wayfinder-router init to scaffold [gateway.models]
(then wayfinder-router doctor to confirm your keys resolve) — see Quickstart .
Works with any OpenAI-compatible API
Wayfinder forwards each call to an OpenAI-style /chat/completions endpoint — so if
your provider speaks that (and most do), it just works. A tier is one base_url ,
a model name, and a key read from the environment at request time; no SDK, no
per-provider code. Pair a free local model with a hosted one, or run two cloud tiers.
…plus Groq, Together, OpenRouter, Fireworks, DeepSeek, and local servers
(vLLM, LM Studio, llama.cpp) — + any OpenAI-compatible endpoint
that takes a Bearer key.
Put Wayfinder in front of your models. Your app keeps speaking the OpenAI API; you
just change one base_url .
Scaffold a config — init writes a starter wayfinder-router.toml (keyless local
Ollama → Anthropic cloud) plus a .env.example , then checks your keys:
pip install " wayfinder-router[gateway] "
wayfinder-router init # starter config (hybrid preset)
wayfinder-router init --preset openai # two OpenAI tiers (gpt-4o-mini → gpt-4o)
wayfinder-router init --preset gemini # two Gemini tiers (gemini-2.5-flash → gemini-2.5-pro)
wayfinder-router init --interactive # pick providers/models step by step
Or describe your two models in wayfinder-router.toml by hand:
[ routing ]
threshold = 0.5 # below -> local, at/above -> cloud
[ gateway . models . local ]
base_url = " http://localhost:11434/v1 "
model = " llama3.2 "
[ gateway . models . cloud ]
base_url = " https://api.openai.com/v1 "
model = " gpt-4o "
api_key_env = " OPENAI_API_KEY " # read from this env var, never stored
# api_key_cmd = "op read op://Private/OpenAI/credential" # optional: fill it from a vault
Wayfinder never stores secrets: a model names an env var ( api_key_env ) and the key
is read from your environment at request time. There is nothing to "install" — just
export the variable. Prefer not to paste a raw key into your shell? Add an optional
api_key_cmd and Wayfinder fills that variable from your secret store at startup —
op read … (1Password), security … (macOS Keychain), secret-tool … (Linux),
pass / gopass , vault kv get … , aws secretsmanager get-secret-value … , bw ,
doppler , gcloud secrets … , or any command that prints the secret. The key is held
in memory only, still never written to disk. wayfinder-router doctor detects which
of these tools you have installed and suggests the exact line.
Set your key(s), then run the gateway. doctor re-checks the config and whether each
model's key resolves ( ✓ set / ✗ not set ) before you start:
export ANTHROPIC_API_KEY=sk-... # or OPENAI_API_KEY, per your config
wayfinder-router doctor # ✓/✗ per model — is each key set?
wayfinder-router serve --port 8088
Point your existing client at it. No code change:
client = openai . OpenAI ( base_url = "http://localhost:8088/v1" , api_key = "unused" )
client . chat . completions . create ( model = "auto" , messages = [{ "role" : "user" , "content" : "..." }])
Easy prompts go local, hard ones go cloud, and every response carries
x-wayfinder-router-model and x-wayfinder-router-score so you can see where it
went. Want to steer one request? Pin it with model="cloud" / prefer-local , or
move the cut for a single call with an X-Wayfinder-Threshold header (see
Steer a single request ).
curl -s localhost:8088/healthz
# {"status":"ok","models":["cloud","local"]}
curl -s -D - -o /dev/null http://localhost:8088/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {"model":"auto","messages":[{"role":"user","content":"hi"}]} ' \
| grep -i x-wayfinder-router
# x-wayfinder-router-model: local
# x-wayfinder-router-score: 0.00
No backends yet? wayfinder-router serve --dry-run answers with the routing
decision instead of calling an upstream, so you can feel the routing in 30 seconds
before wiring up real models.
command
what you get
pip install wayfinder-router
scorer, CLI, Python API, and the terminal chat ( chat ); the scorer/library imports stay dependency-light
pip install "wayfinder-router[gateway]"
adds the OpenAI-compatible routing gateway, the common case for serving
pip install "wayfinder-router[ui]"
adds the local calibrate / explain / configure UI
pip install "wayfinder-router[all]"
gateway and UI on top of the default install
How it works
Wayfinder sits behind whatever OpenAI-compatible client you already use. You point
that client's base_url at the gateway once, and from then on it is invisible. The
same client serves a request whether it routes local or hosted.
your client (chat app, IDE, agent, or code)
|
v
Wayfinder gateway scores, picks a model
|
|-- low --> local (Ollama, vLLM)
|-- high --> hosted (OpenAI, any /v1)
|
v
response returns via the same client,
with x-wayfinder-router-* headers
A few things follow from this:
The interface in front is yours. A chat GUI (Open WebUI, LibreChat), an IDE
assistant with a custom endpoint (Cursor, Continue), an agent framework, or your
own code on the OpenAI SDK. Want a chat window today? Put Open WebUI in front and
point it at the gateway.
Local and hosted are backends, not apps. The local model is just a server
(Ollama, LM Studio, vLLM, llama.cpp) speaking OpenAI's /v1 ; the hosted one is
the same shape. The user never switches UIs and usually never knows which model
answered.
The score is computed, not a second opinion. Asking a model how hard a
prompt is would be slow, non-deterministic, and would cost a model call to decide
whether to make a model call. Wayfinder scans the prompt instead — structure
(length, headings, steps, links, code, tables) and difficulty cues in the wording
(reasoning terms, math symbols, constraints) — into a 0.0 - 1.0 value and
compares it to your threshold. Same prompt, same threshold, same answer. It is a
proxy for difficulty, not a verdict, which i

[truncated]
