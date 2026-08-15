---
source: "https://riffn.io/how-riffn-works/"
hn_url: "https://news.ycombinator.com/item?id=49308095"
title: "Show HN: Riffn. An instant voice link with your AI agents and local models"
article_title: "How Riffn works | Talk to the agents in your repos hands-free"
author: "riffn"
captured_at: "2026-08-15T06:20:36Z"
capture_tool: "hn-digest"
hn_id: 49308095
score: 1
comments: 0
posted_at: "2026-08-15T06:05:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Riffn. An instant voice link with your AI agents and local models

- HN: [49308095](https://news.ycombinator.com/item?id=49308095)
- Source: [riffn.io](https://riffn.io/how-riffn-works/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T06:05:09Z

## Translation

タイトル: 表示 HN: リフン。 AI エージェントおよびローカル モデルとのインスタント音声リンク
記事のタイトル: Riffn の仕組み |リポジトリ内のエージェントとハンズフリーで会話できます
説明: Riffn は、Tailscale を介して、iPhone を自分のコンピュータ上のリポジトリで実行されているコーディング エージェントにリンクします。セッションを開始し、携帯電話をポケットに入れて、運転中、散歩中、または家事をしながら音声で作業を続けます。
HN テキスト: 散歩中や運転中に最高のアイデアが浮かびます。私が Riffn を構築したのは、これらのアイデアをいつでもどこででも、つまりデスクから離れているとき、通常は屋外にいるときに確実に活用できるようにするためです。以前は音声録音をよく使っていました。そこで私は、手と目が忙しいときに話したいモデルやエージェントへの直接音声リンクに置き換えるために Riffn を構築しました。これにより、外出中でもアイデアを発展させ、アプリやその他のプロジェクトを進め続けることができます。 Riffn は、コーディング エージェントまたはローカルにインストールされた LLM への音声リンクを提供する iOS アプリです。私は、Tailscale を介して Riffn を自分のマシン上で実行されているエージェントやモデルにリンクする、小さなオープンソース ブリッジを構築しました。ブリッジのインストールは、1 つの npx コマンドと QR ペアリングだけです。各エージェントに名前を付け、3 単語のコマンドでエージェントを自由に切り替えます。 「Start Riffn」をクリックし、イヤホンを差し込み、電話をロックし、ポケットに入れて通話を開始します。私はゆっくり考えるためにこれを使用するのが好きです。バックグラウンドでアクティブにしておくことができ、運転中、犬の散歩中、または洗濯物を干している間、次のアイデア、音声ターン、または認識されたコマンドに備えることができます。ブリッジは GitHub で入手できます。今のところ、これは iOS のみです。 Apple の音声認識を使用します。

記事本文:
リフン
メニュー
ホーム
仕組み
コーディングエージェント
ローカルモデル
モデル
音声制御
価格設定
サポート
リフンの仕組み
ポケットに携帯電話を入れてレポエージェントに話しかける
Riffn は、リポジトリで実行されているコーディング エージェントに iPhone をリンクする iOS アプリです。
自分のコンピュータ、自分の Tailscale ネットワーク経由で。アプリでセッションを開始し、ロックします。
電話をかけて片付けます。セッションはバックグラウンドでアクティブなままで、次の発話の準備が整います。
運転中、歩行中、または家事中に、または認識されたコマンドを実行します。
ブリッジは公開ソースです: @riffn/bridge は最大 100 KB のノード ヘルパーです
ランタイム依存関係がゼロの Apache-2.0、npm 来歴を持つ CI から公開されています。
お使いのブラウザではこのビデオを再生できません。
録音 (MP4) をダウンロードします。
コンピュータ上のリポジトリ内でブリッジを実行します。
通信したい作業ディレクトリから:
npx @riffn/bridge@0.6.0 初期化
フローティング npx @riffn/bridge を実行するのではなく、バージョンを固定します。あなたは
マシン上でエージェントを駆動するコードを実行するため、ピン留めするということは、読み取ったコードが
実行するコード。
init はエージェントを検出し、ベアラー トークンを生成し、Tailscale をチェックし、
QR コードをペアリングし、フォアグラウンドで実行されます。
Tailscale 経由で接続し、QR でペアリングします
Tailscale は、同じテールネットにサインインしたコンピューターと電話の両方で実行されます。リフンでは、
[設定] → [AI 設定] → [Repo Agent] → [マシンをリンク] を開きます。
QR コードをスキャンし、接続に名前を付けて、 [オンライン] と表示されていることを確認します。
パブリック インターネットには何も公開されません。ブリッジはデフォルトでテールネットにバインドされ、
意図的にオーバーライドしない限り、0.0.0.0 は拒否されます。
セッションを開始し、電話を置いて通話を続けます
アプリでセッションを開始します。それ以降、バックグラウンドでアクティブなままになります。
電話はロックされており、次の発話ターンまたは認識されたコマンドの準備ができています。それは動作します

と
イヤホンと CarPlay で。
新しいアカウントのデフォルトは自動送信で、沈黙の後にターンが送信されます。ほとんどの人は、
声に出して考えて、代わりにトリガーワードに切り替えます。「トリガーをアクティブにする」と言います。
フレーズ」と入力すると、それ以降は、あなたが言うまで何も送信されません
「over」なので、邪魔されることなく、立ち止まって考えることができます。
Claude Code はすでにコンピュータにインストールされ、認証されている必要があります。
Node 18+ と Tailscale を使用。 Riffn には、Claude Code、モデル サブスクリプション、または API は含まれていません
クレジットがあり、リポジトリ パス (CLI およびプロバイダー) 上でモデルを実行することはありません。
アカウントは、ターミナルでクロード コードを入力したかのように、これを実行します。
OpenAI、GLM、またはその他のモデルをリポジトリ パス上で実行できますが、Claude 経由でのみ実行できます。
LiteLLM、Z.ai などの Anthropic 互換ゲートウェイを指すコード ハーネス
ルーターを開きます。これは意図的なセキュリティの選択であり、欠けている機能ではありません。直接 Codex ブリッジングは
ネイティブ Windows では Codex シェルのサブプロセスがファイル読み取りをバイパスできるため、ポリシーによって無効になっています
拒否するため、封じ込められたふりをするのではなく、封じ込められません。ハーネスを介したルーティング
以下のすべての保証はクロード コードによってマシンに適用されるため、有効なままになります。
モデルによってではなく。
試す前に知っておくべき 2 つのこと: これは OpenAI API の課金を使用します — ChatGPT または
Codex サブスクリプションでは API キーは提供されません。 LiteLLM はサードパーティ ソフトウェアです。
はそのキーを保持しているため、プロキシをローカルで認証された状態に保ちます。
リンクするすべてのエージェントとモデルに、任意の名前を付けます。彼らに一語の名前を付けて、
「グロリアに切り替え」は 3 単語のコマンドになります。メニューもモデル ピッカーも必要ありません。
セッション中に電話を触る。 Riffn は電話上のコマンドを認識し、それに基づいて動作します。
つまり、あなたが現在話しているモデルは必要ありません

rはそれを見ます。
典型的なドライブ: 機能についてグロリアと話し、「Milo に切り替えて」と言って、
次にローカル モデルに切り替え、次に「キミに切り替え」で他のモデルのエージェントに移動します。
レポ。各ブリッジは独自のポートと独自のプロジェクト状態を取得するため、複数のブリッジを並行して実行できます。
1台のマシン。
すべてのエンドポイントを最初に構成してペアリングする必要があります。セットアップしていないモデルまたはエージェントは、
に切り替えることはできません。接続に ASSISTANT という名前を付けることは避けてください。
PROJECTS または RIFFN CLOUD — これらは予約されています。
2 つのパスがあり、どちらも Riffn のワーカーを介して推論トラフィックを伝送しません。
そのパスの外側: 音声認識は Apple の Speech を使用します
したがって、Riffn とローカル モデルを組み合わせても、パス全体がローカルになるわけではありません。
アカウントとサブスクリプションの状態は Riffn のワーカーによって処理されます。ご希望の場合は、
エンドツーエンドのローカルな主張ですが、このページは主張していません。
ブリッジは、応答テキスト、ジョブ ステータス (歩数とアクティビティ) のみを電話に送信します。
カテゴリ、決してファイル内容を含まない）と編集された健康概要。
これはあなたのマシンへのリモートコード実行です
これが正直なフレーム構成であり、橋が走る前に読めるほど小さいのはそのためです。
テールネット上に留まり、リクエストごとにベアラー トークンが必要となり、エージェントを実行します。
デフォルトでは読み取り/計画専用。
デフォルトでは読み取り/計画専用。別の方法を選択しない限り、何も書き込まれません。
書き込みアクセスはコンピュータ上でのみ有効になり、電話からは決して有効になりません。ペアになった
電話、または盗まれたトークンは、それ自身のアクセス許可を高めることができません。
制限付きモードは 2 つのキーで動作します。初期化時に選択します。
マシンにアクセスし、音声で確認すると、1 つのタスクだけでファイルを編集できます。
書き込み可能な実行の前にスナップショットが取得されます (完全なリポジトリの状態)。
git ref として、コミットされていないファイルや追跡されていないファイルも含まれます。スナップショットを取得できない場合は、
実行はrです

無防備に走るのではなく、発散します。
シェルと Git の実行はブロックされたままになります。クロードは決してコマンドを実行しません
任意の層: テストなし、インストールなし、コミットなし、プッシュなし、サブエージェントなし、MCP ツールなし。
ターンはシングルフライトです。 2 番目の同時ターンは 429 を獲得します
同じディレクトリに対して 2 番目のエージェントを生成する代わりに。
テレメトリ、分析、クラッシュレポートはありません。ログはマシン上に残り、
デフォルトでトークン、プロンプト、コード、パスを編集します。
状態はワークスペースの外に存在します。トークン、セッション、スナップショット、監査ログ
~/.riffin-bridge/ の下に保存されます。起動時にリポジトリ内の状態ディレクトリを拒否します。
ランタイム依存関係なし (ノード組み込みのみ)、npm を使用して CI から公開
出所はパッケージページで確認できます。
1 つのオペレーター制御、初期化時にマシン上またはブリッジで選択
.env 。電話からは決して変更できません。
Limited および Ungate では、編集タスクは PreToolUse フックの下で実行されます。
クロード コードは、すべてのツール呼び出しの前に参照します。読み取り、編集、書き込み、Glob、Grep、Web のみが許可されます。
そしてそれ以外はすべて否定します。ファイル ツールは、作業ディレクトリ、環境、および
シークレット ファイルは読み取りも書き込みもできず、見つからないパスは失敗して閉じられます。
明確に述べておく価値のある制限が 1 つあります。ブリッジは任意のコーディング エージェントの前にも立つことができます。
カスタム エージェント モードの CLI。自分が作成していないツールに対して読み取り専用を強制することはできません。
CLI が持つ権限が何であれ、あなたの声が変わります。ツール自体を構成します。橋
caps: 演算子定義をレポートするため、アプリは正直な状態を表示します。
Riffn は iOS のみです。 Android や Web クライアントはありません。
ローカル モデルのリンクは無料で、Apple でサインインする必要があります。
リポ エージェントをリンクするには、Riffn コード プランが必要です。現在の価格は
価格ページおよびApp Storeのリスト内。このページでは繰り返しません
しびれている

えー、それは古くなってしまいます。
自分のコンピューター、Tailscale のセットアップ、エージェントのインストール、およびサードパーティのモデルを持ち込みます。
エージェントのサブスクリプションと API クレジット。これらは、Riffn ではなく、プロバイダーによって請求されます。
実行する前にブリッジを読んでください。それがブリッジが小さいことのポイントです。
エージェントとモデルへのインスタント音声リンク。
© 2026 リフン.無断転載を禁じます。

## Original Extract

Riffn links your iPhone to a coding agent running in a repo on your own computer, over Tailscale. Start a session, pocket the phone, and keep working by voice while you drive, walk or do chores.

I get my best ideas when I'm out walking or driving. I built Riffn to make sure I can capitalize on those ideas when and where they happen - while I'm away from my desk, usually outside. I used to use voice recording a lot. So I built Riffn to replace that with a direct voice link to any model or agent I wanted to talk to when my hands and eyes were busy. This lets me develop ideas and keep pushing forward on my apps and other projects when I'm out and about. Riffn is an iOS app that gives you a voice link to your coding agents or locally-installed LLMs. I built a small, open-source bridge that links Riffn to agents and models running on your own machine over Tailscale. Installing the bridge is just one npx command and QR pairing. Give each agent a name, then switch freely between agents with a 3-word command. Click 'Start Riffn', pop in an earbud, lock the phone, put it in your pocket, and start talking. I like to use it for slow thinking - you can leave it active in the background, ready for your next idea, spoken turn or a recognized command while you drive, walk the dog, or hang out the washing. The bridge is available on GitHub. For now this is iOS only. It uses Apple speech recognition.

Riffn
Menu
Home
How it works
Coding Agents
Local Models
Models
Voice Control
Pricing
Support
How Riffn works
Talk to your repo agent with the phone in your pocket
Riffn is an iOS app that links your iPhone to a coding agent running in a repo on
your own computer, over your own Tailscale network. Start a session in the app, lock the
phone and put it away: the session stays active in the background, ready for your next spoken turn
or a recognised command, while you drive, walk or do chores.
The bridge is public source: @riffn/bridge is a ~100 KB Node helper
with zero runtime dependencies, Apache-2.0, published from CI with npm provenance.
Your browser cannot play this video.
Download the recording (MP4) .
Run the bridge inside the repo on your computer
From the working directory you want to talk to:
npx @riffn/bridge@0.6.0 init
Pin the version rather than running a floating npx @riffn/bridge . You are
running code that drives an agent on your machine, so pinning means the code you read is
the code you run.
init detects your agent, generates a bearer token, checks Tailscale, prints a
pairing QR code and then runs in the foreground.
Connect over Tailscale and pair by QR
Tailscale runs on both the computer and the phone, signed in to the same tailnet. In Riffn,
open Settings → AI Settings → Repo Agent → Link my machine ,
scan the QR code, give the connection a name and check that it shows Online .
Nothing is exposed to the public internet: the bridge binds to your tailnet by default and
refuses 0.0.0.0 unless you deliberately override it.
Start a session, put the phone away, keep talking
You start the session in the app. From then on it stays active in the background with the
phone locked, ready for your next spoken turn or a recognised command. It works with
earbuds and in CarPlay.
New accounts default to auto-send, which sends a turn after a silence. Most people who
think out loud switch to the trigger word instead: say “activate trigger
phrase” and from then on nothing is sent until you say
“over” , so you can pause to think without being interrupted.
Claude Code must already be installed and authenticated on your computer , along
with Node 18+ and Tailscale. Riffn does not include Claude Code, a model subscription or any API
credit, and it does not run a model for you on the repo path — your CLI and your provider
account do that, exactly as if you had typed into Claude Code in a terminal.
You can run an OpenAI, GLM or other model on the repo path, but only through the Claude
Code harness , pointed at an Anthropic-compatible gateway such as LiteLLM, Z.ai or
OpenRouter. That is a deliberate security choice, not a missing feature: direct Codex bridging is
disabled by policy because on native Windows, Codex shell subprocesses can bypass file-read
denials, so it fails closed rather than pretending to be contained. Routing through the harness
keeps every guarantee below in force, because they are enforced on your machine by Claude Code
rather than by the model.
Two things to know before you try it: this uses OpenAI API billing — a ChatGPT or
Codex subscription does not supply the API key ; and LiteLLM is third-party software that
holds that key, so keep the proxy local and authenticated.
You give every agent and model you link a name of your choosing. Give them one-word names and
“Switch to Gloria” becomes a three-word command — no menu, no model picker, no
touching the phone during the session. Riffn recognises the command on the phone and acts on it
there, so the model you are currently talking to never sees it.
A typical drive: talk to Gloria about a feature, say “Switch to Milo” and send the
next turn to the local model, then “Switch to Kimi” to move to the agent in the other
repo. Each bridge gets its own port and its own project state, so several can run side by side on
one machine.
Every endpoint has to be configured and paired first; a model or agent you have not set up is
not available to switch to. Avoid naming a connection ASSISTANT ,
PROJECTS or RIFFN CLOUD — those are reserved.
There are two paths, and neither one carries your inference traffic through Riffn’s Worker.
Outside that path: speech recognition uses Apple’s Speech
framework, so pairing Riffn with a local model does not make the whole path local.
Account and subscription state is handled by Riffn’s Worker. If you want an
end-to-end-local claim, this page is not making one.
The bridge sends the phone only the reply text, job status (step counts and activity
categories , never file contents) and a redacted health summary.
This is remote code execution into your machine
That is the honest framing, and it is why the bridge is small enough to read before you run it.
It stays on your tailnet, requires a bearer token on every request, and runs the agent
read/plan-only by default.
Read/plan-only by default. Nothing writes unless you chose otherwise.
Write access is enabled only on the computer, never from the phone. A paired
phone, or a stolen token, cannot raise its own permissions.
Limited mode is two-key armed: you pick it at init on the
machine, and you confirm by voice, and then exactly one task may edit files.
A snapshot is taken before any write-capable run — the full repo state
including uncommitted and untracked files, as a git ref. If the snapshot cannot be taken, the
run is refused rather than run unprotected.
Shell and Git execution stay blocked. Claude never gets command execution at
any tier: no tests, no installs, no commits, no push, no subagents, no MCP tools.
Turns are single-flight. A second concurrent turn gets a 429
instead of spawning a second agent against the same directory.
No telemetry, no analytics, no crash reporting. Logs stay on your machine and
redact the token, prompts, code and paths by default.
State lives outside the workspace. Tokens, sessions, snapshots and audit logs
are kept under ~/.riffin-bridge/ ; startup refuses a state directory inside the repo.
Zero runtime dependencies (Node built-ins only), published from CI with npm
provenance you can verify on the package page.
One operator control, chosen on the machine at init or in the bridge’s
.env . It can never be changed from the phone.
In limited and ungated , the edit task runs under a PreToolUse hook that
Claude Code consults before every tool call: it allows only Read, Edit, Write, Glob, Grep and web,
and denies everything else. File tools are confined to the working directory, environment and
secret files cannot be read or written, and missing paths fail closed.
One limit worth stating plainly: the bridge can also front an arbitrary coding-agent
CLI in custom-agent mode. It cannot enforce read-only on a tool it did not write —
whatever permissions that CLI has, your voice turns have. Configure the tool itself; the bridge
reports caps: operator-defined so the app shows the honest state.
Riffn is iOS only . There is no Android or web client.
Linking a local model is free , and requires Sign in with Apple.
Linking a repo agent requires the Riffn Code plan . Current prices are on the
pricing page and in the App Store listing; this page does not repeat a
number that would go stale.
You bring your own computer, Tailscale setup, agent installation, and any third-party model or
agent subscription and API credit. Those are billed by their providers, not by Riffn.
Read the bridge before you run it — that is the point of it being small.
An instant voice link to your agents and models.
© 2026 Riffn. All rights reserved.
