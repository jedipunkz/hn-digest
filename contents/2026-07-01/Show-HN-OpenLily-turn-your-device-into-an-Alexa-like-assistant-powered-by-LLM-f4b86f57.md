---
source: "https://github.com/getlark/openlily"
hn_url: "https://news.ycombinator.com/item?id=48745906"
title: "Show HN: OpenLily – turn your device into an Alexa-like assistant powered by LLM"
article_title: "GitHub - getlark/openlily: A personal voice assistant with a wake word, swappable LLMs (gpt 5.5, opus 4.8, etc), and configurable tools (web search, email, calendar, etc). · GitHub"
author: "crush_robo_1536"
captured_at: "2026-07-01T12:54:20Z"
capture_tool: "hn-digest"
hn_id: 48745906
score: 1
comments: 0
posted_at: "2026-07-01T12:52:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenLily – turn your device into an Alexa-like assistant powered by LLM

- HN: [48745906](https://news.ycombinator.com/item?id=48745906)
- Source: [github.com](https://github.com/getlark/openlily)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T12:52:57Z

## Translation

タイトル: HN を表示: OpenLily – デバイスを LLM を利用した Alexa のようなアシスタントに変える
記事のタイトル: GitHub - getlark/openlily: ウェイク ワード、交換可能な LLM (gpt 5.5、opus 4.8 など)、および構成可能なツール (Web 検索、電子メール、カレンダーなど) を備えたパーソナル音声アシスタント。 · GitHub
説明: ウェイク ワード、交換可能な LLM (gpt 5.5、opus 4.8 など)、および設定可能なツール (Web 検索、電子メール、カレンダーなど) を備えたパーソナル音声アシスタント。 - ゲットヒバリ/オープンリリー

記事本文:
GitHub - getlark/openlily: ウェイク ワード、交換可能な LLM (gpt 5.5、opus 4.8 など)、および構成可能なツール (Web 検索、電子メール、カレンダーなど) を備えたパーソナル音声アシスタント。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。これをリロードしてください

ページに
ヒバリ
/
オープンリリー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .vscode .vscode docs docs server server .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
openlily は、Alexa に似たパーソナル音声アシスタントです。あなたはそれに話しかけます
独自のマイクとスピーカーを通じて、音声入力→LLM→音声出力と応答できます。
ツール (Web 検索、ブラウザー) を通じて質問し、説明し、アクションを実行します。
自動化、電子メールなど）。これは、マシン上の端末音声 CLI として実行されます。
オプションのウェイク ワードを使用すると、呼び出すまで静かに待機します。
それはあなたのものになるように構築されています: 基礎となるモデル (LLM、音声からテキストへの変換、
テキスト読み上げ)、信頼できるプロバイダーを選択し、必要なツールのみを有効にします。
ラズベリーパイ、Mac miniなどの他のスタンドアロンデバイスでも実行できます。
openlily の動作をデモビデオでご覧ください。
ローカル音声 CLI — マイクとスピーカーがクライアントです。ブラウザも電話もありません
必須です。スタンドアロン WebRTC オーディオ処理モジュール (AEC + ノイズ抑制 +
AGC) は、ボットが自分自身を聞くのを防ぎます。
交換可能な「頭脳」 — カスケード パイプライン (個別の STT → LLM → TTS) または
リアルタイム音声合成モデルを選択し、各部分のプロバイダー/モデルを選択します。
ウェイク ワード — オプションの常時オンのオンデバイス リスナー (openWakeWord)
ウェイク フレーズが聞こえた場合にのみセッションを開始します。クラウドも API キーもありません。
デバイス上のターンテイク — Silero VAD + Smart Turn v3 をローカルで実行して決定します
話し始めたときと話し終わったとき。
ツール - Web 検索、実際のブラウザ自動化、電子メール、それぞれ

ポイントイン。カスタム ツールを簡単に作成できます。
UV同期
ローカルオーディオパスには、PyAudio 用の PortAudio が必要です。 OS に合わせてインストールします。
Linux / Raspberry Pi (Debian、Ubuntu、Raspberry Pi OS):
sudo aptアップデート
sudo apt install portaudio19-dev libportaudio2
Raspberry Pi では、マイクとスピーカーが認識されていることも確認してください
( arecord -l と aplay -l がそれらをリストする必要があります)。 USBマイク/スピーカーまたはUSB
オーディオインターフェイスは最も簡単なセットアップです。適切な入出力デバイスを選択してください
複数ある場合は、ALSA/PulseAudio 構成に含めます。
Linux では、ウェイクワード スタックは tflite-runtime を取り込みます。これは、
Python 3.11 用のホイール。リポジトリはそのバージョンを .python-version に固定します。
UV Sync はそれを自動的に使用します。次のコマンドを使用して一度インストールします。
お持ちでない場合は、uv python 3.11 をインストールしてください。 (macOSでは必要ありません
tflite-runtime なので影響を受けません。)
ブラウザ ツール (有効にした場合) は、次の方法で Playwright MCP サーバーを起動します。
npx なので、Node.js が必要です。 macOS の場合: brew install ノード 。 Linux の場合 /
Raspberry Pi: sudo apt install nodejs npm (または現在のリリースをインストール)
NodeSource から
ディストリビューションのパッケージは古いです）。
環境変数を構成します。
cp .env.example .env
実用的なアシスタントへの最速の道 — 1 つ選択してください:
現状のまま (デフォルトの cartesia_openai ブレイン): OPENAI_API_KEY を設定し、
CARTESIA_API_KEY 。それでおしまい。デカルト キーを入手するには、
デカルト.ai 。
OpenAI キーのみ、Cartesia なし:default_brain を openai_realtime に切り替えます
Brains.yaml (以下を参照) で OPENAI_API_KEY のみを設定します。声が出るよ
出たり入ったり、ウェブ検索はありません。
OpenAI キーのみ、Web 検索あり: 代わりに openai_standard を使用してください - 実行されます
OPENAI_API_KEY のみを使用して、完全に OpenAI (組み込み Web 検索を含む) 上で動作します。
API キーがまったくありません (完全にローカル):default_brain を次のように切り替えます。
local_whisper_ollama_kokoro を実行し、ローカルの Ollama を実行します
サーバー。すべて (STT、LLM、

TTS) はマシン上で実行されます — を参照してください。
完全にローカルで実行します。
.env 内のその他すべてはオプションであり、必要に応じてグループ化されます。参照
フルメニューに合わせてアシスタントをカスタマイズします。
uv run bot.py # デフォルト: ウェイクワードゲートされたローカルセッション
uv run bot.py --mode local # マイク/スピーカー音声 CLI、ウェイク ワードなし
uv run bot.py --mode webrtc # localhost:7860 でブラウザ デバッグ UI
最初の実行は開始までに時間がかかります (通常は数秒から最大 1 秒かかります)。
分 — Python が依存関係とオンデバイス ウェイクワード/VAD をコンパイルしている間
モデルを一度ダウンロードします。ターミナルはすぐに「モジュールのロード」行を出力します
それで、それが行き詰まっていないことがわかります。後の実行は数秒以内に開始されます。
openlily はニーズに合わせて構成できるように設計されています。 3 つのノブ:
1. モデルとプロバイダー (「頭脳」) を選択します。
どのモデルが音声テキスト変換、言語変換、テキスト音声変換を行うかを脳が決定します。
Brains.yaml 内のdefault_brain を含むものを選択します ( Brains.yaml.example をコピーします)。
ファイルがない場合、デフォルトは cartesia_openai です):
cartesia_openai (デフォルト) — 全体的に最も効果的: インテリジェントな OpenAI
LLM は、Cartesia の強力な音声テキスト変換およびスムーズで自然な TTS と組み合わせられています。の
デフォルトの LLM は gpt-5.4-mini です。 gpt-5.5 などのより高性能なモデルに変更します。
Brains.yaml は、応答が遅くなる代わりに、より高いインテリジェンスを実現します。
openai_standard — セットアップが最も簡単: 単一の OpenAI API キーで取得できます
すべて (STT、LLM、TTS)、2 番目のプロバイダーはありません。
openai_realtime — 個別の STT/TTS がないため、最も速く感じられます。
ただし、音声読み上げモデルの能力は最新のものよりも劣る可能性があります。
非リアルタイム OpenAI モデル。
local_whisper_ollama_kokoro — 完全にデバイス上にあり、API キーは不要です: MLX
STT の場合は Whisper、LLM の場合は Ollama、TTS の場合は Kokoro です。
プライバシーまたはオフラインでの使用に最適です。品質と速度はハードウェアに依存します。
ウェブの海はない

チャーチ。 Apple Silicon (MLX) と実行中の Ollama サーバーが必要です —
「完全にローカルで実行する」を参照してください。
同じ Brains.yaml で、各 Brain のモデル名と TTS をオーバーライドできます。
コードに触れずに音声 — 例: LLM を別のモデルに向けるか、変更します。
デカルトの音声 ID。リストにないプロバイダー (別の STT/TTS) が必要です
ベンダー、ローカル LLM)?脳の追加は、小さな自己完結型の変更です — を参照してください。
貢献.md 。
local_whisper_ollama_kokoro の脳はマシン上でパイプライン全体を実行します。
したがって、プロバイダー API キーは必要ありません。現在、Apple Silicon (
STT は MLX Whisper を使用します)。
Ollama をインストールして起動し、LLM をプルします。
ollama pull gemma4:e4b # またはより軽量/高速なモデルの gemma4:e2b
ボットを開始する前に、サーバーが実行中であり、アクセス可能であることを確認してください —
ollam ps (またはcurl http://localhost:11434/api/tags ) が応答するはずです。もし
Ollama はデフォルトの http://localhost:11434 以外の場所で実行されます。
OLLAMA_BASE_URL (例: http://my-host:11434/v1 )。
Brains.yaml で脳を選択します。
デフォルトブレイン : local_whisper_ollama_kokoro
それを実行します ( uv run bot.py --mode local )。最初の実行では STT (MLX
Whisper) および TTS (Kokoro) モデルはローカルにダウンロードしてキャッシュするため、起動は
一度ゆっくりと。フォンマイザーが見つからないときに Kokoro エラーが発生する場合は、espeak-ng をインストールしてください
( brew install espeak-ng )。
この脳には Web 検索がありません (すべての Web/ブラウザ/電子メール ツールは外部を呼び出します)
サービス）。それ以外のすべて (音声入力、LLM、音声出力) はオフラインで動作します。
ツールはオプトインです。ブラウザと電子メール ツールは集中的に接続されており、
デフォルトではオフ — エントリのコメントを解除して有効にします。
GENERIC_TOOL_SETUPS (server/tools/__init__.py)
各ツールは、その資格情報が存在し、セッションがまだ実行されている場合にのみアクティブになります。
それらがなくても大丈夫です。
Web 検索 — デフォルトでオンになっており、その取得方法

脳に依存します。の
OpenAI カスケード ブレイン ( openai_standard 、 cartesia_openai ) は OpenAI の
組み込みのホスト型 Web 検索が自動的に実行されるため、追加のキーは必要ありません。 openai_realtime
脳は代わりに Exa を呼び出すため、EXA_API_KEY が必要です (これがないと、リアルタイム
脳はウェブ検索なしで実行されます)。完全にローカルな
local_whisper_ollama_kokoro 脳にはウェブ検索がまったくありません。
ブラウザ (Playwright MCP) — 実際のローカル ブラウザを駆動します。 Node.js/npx が必要です。
独自のブラウザを起動するのではなく、CDP 経由ですでに実行中のブラウザに接続します。
BROWSER_CDP_ENDPOINT を設定します (例: http://localhost:9222 、Chrome を起動したとき)
--remote-debugging-port=9222 を使用して) 有効にします。その後、ブラウザは継続します
セッション全体で。この変数がないと、ブラウザ ツールはスキップされます。
電子メール (再送信) — 自分のアドレスに電子メールを送信します。 USER_EMAIL が必要です。
RESEND_API_KEY 、および確認された送信者 ( EMAIL_FROM )。
独自のツールを作成することも小さな変更です。 CONTRIBUTING.md を参照してください。
uv run bot.py (または --mode local-with-wake-word ) はプロセスをウォームに保ち、
はウェイク ワードを聞くとセッションを開始するため、各セッションは迅速に開始されます。を設定します。
WAKE_MODELS を含むフレーズ (カンマ区切り、デフォルトは alexa )。内蔵
事前にトレーニングされたフレーズ:
いくつかリストしてそれらのいずれかを受け入れるか (例: WAKE_MODELS=alexa,hey_jarvis )、またはポイントします
独自の .onnx / .tflite モデル ファイルをパスで指定します。
ローカル音声 CLI では、ボットが話している間、マイクは半二重ゲートされます。
発言の途中で中断することはできません。ウェイクワードバージイン (ウェイクワードを言ってください)
ボットを切断します) はデフォルトで無効になっています。試してみたい場合は、裏返してください
server/transport_local.py で WAKE_WORD_BARGE_IN を True に設定します。
local-with-wake-word (デフォルト) — ウォームプロセス。常時接続のリスナーが所有する
マイクをオンにしてウェイク ワードで音声セッションを開始し、次の時点でリスニングを再開します。
セッションはアイドル状態になります。

ローカル - マイク + スピーカー音声 CLI。すぐに話してください。ウェイクワードはありません。
webrtc — localhost:7860 のブラウザ デバッグ UI。
セッションは、沈黙が続いた後（誰も話さない）自動的に終了します。で調整してください
IDLE_TIMEOUT_SECS 。
openlily はいくつかの小さな音声キューを使用するので、自分がどの場所にいるのかを常に知ることができます。
ターミナルを見ずに回してください:
セッションの準備が整ったとき、ウェイク ワードの後、2 つの音の上昇音「ディン」
(またはローカル モードでの起動時)。これは、接続されており、マイクがオンになっていることを意味します
ライブなので、あなたの声が入力として録音されています。
ボットが動作している間、数秒ごとに低い低い「ブリップ」音が鳴ります。
話し終わるとリクエストが LLM に送信されるか、ツール呼び出し中 (Web
検索、ブラウザ、電子メール）。それは人生の静かな兆候だからあなたは沈黙に取り残されない
考えている間。
口頭での返答。 LLM が完了すると、ブリップ音が止まり、次の音が聞こえます。
テキスト読み上げを通じて回答します。
問題が発生したり、質問がありますか? Slack で質問し、
GitHub で問題を解決するか、team@getlark.ai に電子メールを送信してください。
アーキテクチャ、開発セットアップ、および頭脳とツールを追加する方法
貢献.md 。
openlily は、次のような優れたオープンソース プロジェクトを担っています。
Pipecat — リアルタイム音声エージェント フレームワーク
LiveKit — WebRTC オーディオ処理モジュール (AEC/ノイズ抑制/AGC)
開く

[切り捨てられた]

## Original Extract

A personal voice assistant with a wake word, swappable LLMs (gpt 5.5, opus 4.8, etc), and configurable tools (web search, email, calendar, etc). - getlark/openlily

GitHub - getlark/openlily: A personal voice assistant with a wake word, swappable LLMs (gpt 5.5, opus 4.8, etc), and configurable tools (web search, email, calendar, etc). · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
getlark
/
openlily
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .vscode .vscode docs docs server server .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md View all files Repository files navigation
openlily is an Alexa-like personal voice assistant. You talk to it
through your own mic and speakers — voice in → LLM → voice out — and it can answer
questions, explain things, and take actions through tools (web search, browser
automation, email, etc). It runs as a terminal voice CLI on your machine, with an
optional wake word so it sits quietly until you call it.
It's built to be yours : swap the underlying models (LLM, speech-to-text,
text-to-speech), pick a provider you trust, and turn on only the tools you want.
You can also run it on other standalone devices like raspberry pi, mac mini, etc.
Watch the demo video to see openlily in action.
Local voice CLI — your mic and speakers are the client; no browser or phone
required. A standalone WebRTC Audio Processing Module (AEC + noise suppression +
AGC) keeps the bot from hearing itself.
Swappable "brains" — run a cascade pipeline (separate STT → LLM → TTS) or a
realtime speech-to-speech model, and choose the provider/model for each piece.
Wake word — an optional always-on, on-device listener (openWakeWord) that
starts a session only when it hears the wake phrase. No cloud, no API key.
On-device turn-taking — Silero VAD + Smart Turn v3 run locally to decide
when you've started and stopped talking.
Tools — web search, real browser automation, and email, each opt-in. You can write custom tools easily.
uv sync
The local-audio path needs PortAudio for PyAudio. Install it for your OS:
Linux / Raspberry Pi (Debian, Ubuntu, Raspberry Pi OS):
sudo apt update
sudo apt install portaudio19-dev libportaudio2
On a Raspberry Pi, also make sure your mic and speakers are recognized
( arecord -l and aplay -l should list them). A USB mic/speaker or a USB
audio interface is the simplest setup; pick the right input/output device
in your ALSA/PulseAudio config if you have more than one.
On Linux the wake-word stack pulls in tflite-runtime , which only ships
wheels for Python 3.11. The repo pins that version in .python-version , so
uv sync will use it automatically — install it once with
uv python install 3.11 if you don't have it. (macOS doesn't need
tflite-runtime , so it isn't affected.)
The browser tool (if you enable it) launches the Playwright MCP server via
npx , so it needs Node.js. On macOS: brew install node . On Linux /
Raspberry Pi: sudo apt install nodejs npm (or install a current release
from NodeSource if your
distro's packages are old).
Configure environment variables :
cp .env.example .env
The fastest path to a working assistant — pick one:
As-is (default cartesia_openai brain): set OPENAI_API_KEY and
CARTESIA_API_KEY . That's it. Get a Cartesia key at
cartesia.ai .
OpenAI key only, no Cartesia: switch default_brain to openai_realtime
in brains.yaml (see below) and set just OPENAI_API_KEY . You'll have voice
in and out, just no web search.
OpenAI key only, with web search: use openai_standard instead — it runs
entirely on OpenAI (including built-in web search) with only OPENAI_API_KEY .
No API keys at all (fully local): switch default_brain to
local_whisper_ollama_kokoro and run a local Ollama
server. Everything (STT, LLM, TTS) runs on your machine — see
Run it fully local .
Everything else in .env is optional and grouped by when you need it. See
Personalizing your assistant for the full menu.
uv run bot.py # default: wake-word gated local session
uv run bot.py --mode local # mic/speakers voice CLI, no wake word
uv run bot.py --mode webrtc # browser debug UI at localhost:7860
The first run takes longer to start — usually several seconds, and up to a
minute — while Python compiles dependencies and the on-device wake-word/VAD
models download once. The terminal prints a "loading modules" line right away
so you know it isn't stuck; later runs start in a few seconds.
openlily is meant to be configured to your needs. Three knobs:
1. Choose the models and providers (the "brain")
A brain decides which models do speech-to-text, language, and text-to-speech.
Select one with default_brain in brains.yaml (copy brains.yaml.example ;
without the file the default is cartesia_openai ):
cartesia_openai (default) — the most effective overall: intelligent OpenAI
LLM paired with Cartesia's strong speech-to-text and smooth, natural TTS. The
default LLM is gpt-5.4-mini ; bump it to a more capable model like gpt-5.5 in
brains.yaml for higher intelligence at the cost of slower replies.
openai_standard — the easiest to set up: a single OpenAI API key gets you
everything (STT, LLM, TTS), no second provider.
openai_realtime — feels the fastest, since there's no separate STT/TTS
stage, but the speech-to-speech model can be less capable than the latest
non-realtime OpenAI models.
local_whisper_ollama_kokoro — fully on-device and free of API keys: MLX
Whisper for STT, Ollama for the LLM, and Kokoro for TTS.
Best for privacy or offline use; quality and speed depend on your hardware, and
it has no web search. Requires Apple Silicon (MLX) and a running Ollama server —
see Run it fully local .
In the same brains.yaml you can override each brain's model names and the TTS
voice without touching code — e.g. point the LLM at a different model, or change
the Cartesia voice ID. Want a provider that isn't listed (a different STT/TTS
vendor, a local LLM)? Adding a brain is a small, self-contained change — see
CONTRIBUTING.md .
The local_whisper_ollama_kokoro brain runs the whole pipeline on your machine,
so no provider API keys are needed. It currently requires Apple Silicon (the
STT uses MLX Whisper).
Install and start Ollama , then pull the LLM:
ollama pull gemma4:e4b # or gemma4:e2b for a lighter/faster model
Make sure the server is running and reachable before you start the bot —
ollama ps (or curl http://localhost:11434/api/tags ) should respond. If
Ollama runs somewhere other than the default http://localhost:11434 , set
OLLAMA_BASE_URL (e.g. http://my-host:11434/v1 ).
Select the brain in brains.yaml :
default_brain : local_whisper_ollama_kokoro
Run it ( uv run bot.py --mode local ). On the first run the STT (MLX
Whisper) and TTS (Kokoro) models download and cache locally, so startup is
slower once. If Kokoro errors on a missing phonemizer, install espeak-ng
( brew install espeak-ng ).
This brain has no web search (all the web/browser/email tools call external
services). Everything else — voice in, LLM, voice out — works offline.
Tools are opt-in. The browser and email tools are wired in centrally and are
off by default — enable them by uncommenting their entry in
GENERIC_TOOL_SETUPS in server/tools/__init__.py .
Each tool only activates if its credentials are present, and a session still runs
fine without them.
Web search — on by default, and how you get it depends on the brain. The
OpenAI cascade brains ( openai_standard , cartesia_openai ) use OpenAI's
built-in hosted web search automatically — no extra key. The openai_realtime
brain instead calls Exa, so it needs EXA_API_KEY (without it, the realtime
brain just runs without web search). The fully-local
local_whisper_ollama_kokoro brain has no web search at all.
Browser (Playwright MCP) — drives a real local browser. Needs Node.js/ npx .
Attaches to an already-running browser over CDP rather than launching its own,
so set BROWSER_CDP_ENDPOINT (e.g. http://localhost:9222 , from Chrome started
with --remote-debugging-port=9222 ) to enable it; the browser then persists
across sessions. Without that variable the browser tools are skipped.
Email (Resend) — sends email to your own address. Needs USER_EMAIL ,
RESEND_API_KEY , and a verified sender ( EMAIL_FROM ).
Writing your own tool is also a small change — see CONTRIBUTING.md .
uv run bot.py (or --mode local-with-wake-word ) keeps the process warm and only
starts a session once it hears a wake word, so each session starts fast. Set the
phrase(s) with WAKE_MODELS (comma-separated, defaults to alexa ). Built-in
pretrained phrases:
List several to accept any of them (e.g. WAKE_MODELS=alexa,hey_jarvis ), or point
at your own .onnx / .tflite model file by path.
In the local voice CLI the mic is half-duplex gated while the bot is talking, so
it can't be interrupted mid-utterance. Wake-word barge-in (say the wake word to
cut the bot off) is disabled by default ; if you want to try it, flip
WAKE_WORD_BARGE_IN to True in server/transport_local.py .
local-with-wake-word (default) — warm process; an always-on listener owns
the mic and starts a voice session on the wake word, then resumes listening when
the session idles out.
local — mic + speakers voice CLI; talk immediately, no wake word.
webrtc — browser debug UI at localhost:7860 .
A session ends itself after a stretch of silence (no one speaking); tune it with
IDLE_TIMEOUT_SECS .
openlily uses a couple of small audio cues so you always know where you are in a
turn, without watching the terminal:
A rising two-note "ding" when a session becomes ready — after the wake word
(or right at startup in local mode). It means you're connected and the mic is
live, so your voice is now being recorded as input.
A soft, low "blip" every few seconds while the bot is working — after you
finish speaking and the request is sent to the LLM, or during a tool call (web
search, browser, email). It's a quiet sign of life so you're not left in silence
while it thinks.
The spoken reply. Once the LLM is done, the blips stop and you hear the
answer through text-to-speech.
Running into issues or have questions? Ask in Slack , open an
issue on GitHub, or email team@getlark.ai .
Architecture, dev setup, and how to add brains and tools live in
CONTRIBUTING.md .
openlily stands on the shoulders of excellent open-source projects, including:
Pipecat — the real-time voice agent framework
LiveKit — the WebRTC Audio Processing Module (AEC/noise suppression/AGC)
open

[truncated]
