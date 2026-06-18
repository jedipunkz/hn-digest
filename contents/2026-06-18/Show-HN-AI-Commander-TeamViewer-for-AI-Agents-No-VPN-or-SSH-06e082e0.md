---
source: "https://aicommander.dev/"
hn_url: "https://news.ycombinator.com/item?id=48590260"
title: "Show HN: AI Commander – TeamViewer for AI Agents, No VPN or SSH"
article_title: "AI Commander — Remote Shell for AI Agents"
author: "coderai"
captured_at: "2026-06-18T21:46:39Z"
capture_tool: "hn-digest"
hn_id: 48590260
score: 3
comments: 0
posted_at: "2026-06-18T19:24:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Commander – TeamViewer for AI Agents, No VPN or SSH

- HN: [48590260](https://news.ycombinator.com/item?id=48590260)
- Source: [aicommander.dev](https://aicommander.dev/)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T19:24:36Z

## Translation

タイトル: HN を表示: AI Commander – AI エージェント用の TeamViewer、VPN または SSH なし
記事のタイトル: AI Commander — AI エージェント用のリモート シェル
説明: Claude、Codex、ChatGPT、または任意の MCP/HTTP クライアントが所有するマシン上でシェル コマンドを実行できるようにします。公開された SSH、オープン ポート、または VPN を使用しないリモート アクセス。
HN テキスト: Claude、Codex、Opencode、またはその他の AI チャットなどの CLI ツールからリモート コンピューターに瞬時にアクセスできるプラットフォームを構築しました。 TeamViewer など、接続コードを生成する Windows (トレイ)、Mac (メニュー バー)、Linux (CLI アプリ) 用のミニ アプリがあります。包括的なセキュリティ、インスタント アクセス、ファイアウォールやネットワーク構成は不要です。 AI にマシン上で何かをするよう依頼するだけで、完璧に実行してくれます。

記事本文:
AIコマンダー
ユースケース
接続する
セキュリティ
ドキュメント
サインイン
ダウンロード
AIにコマンドの実行を依頼する
数分で作業可能
ネットワーク設定がありません
AI に依頼してリモート コンピューター上でコマンドを実行します。
AI エージェント用のリモート シェル。 Claude、ChatGPT、Codex、または任意の AI アシスタントに、所有するマシン上で実際のシェル コマンドを実行させます。
公開された SSH、オープン ポート、VPN はありません。
アプリをインストールし、マシンコードを共有し、AI に何をするかを尋ねます。アプリは自動的に接続するため、ファイアウォールの変更はありません。
Linux (クラウド サーバー、ホーム サーバー、または Raspberry Pi) では、1 行を貼り付けます。 Mac または Windows では、デスクトップ アプリをインストールします。 AIC-7K3P-WX9M-RTBN のような安定したコードが得られます。
ChatGPT、Claude、または別の AI アシスタントにコードを渡して、わかりやすい言葉で尋ねます。そのマシン上でコマンドが実行され、結果が表示されます。
マシンを 1 か所に保管し、 prod-db のような名前を付ける場合は、サインインします。試すのにアカウントは必要ありません。
制御したいマシンにアプリを入れます
プラットフォームを選択してください。アプリは、接続に使用するマシンコードを提供します。
macOS
窓
Linux
macOS 用のダウンロード
メニューバー アプリ、すべてが含まれる · Apple Silicon (M1–M4) · Intel (x64)
Windows 用のダウンロード
トレイ アプリ、すべてが含まれています · Windows x64
バックグラウンドで静かに実行され、サーバー、クラウド VM、Raspberry Pi 上で動作します
Claude、ChatGPT、Codex などで動作します。最初の 1 時間はアカウントなしで新しいマシンを試すことができます。継続的にアクセスするにはサインインしてください。ほとんどの AI ツールは 1 つのコマンドで接続します。
$
codex mcp aicommander を追加 --url https://aicommander.dev/mcp
$
opencode mcp aicommander を追加 --url https://aicommander.dev/mcp
$
pi インストール npm:@aicommander/mcp
または、わかりやすい言葉で AI に尋ねるだけで、残りの部分は AI が判断します。
aiccommander.dev を使用して AIC-XXX に接続します
Cursor、Windsurf、ChatGPT、または Claude Deskt の使用

代わりに？
すべての接続ガイド →
各オプションは同じ仕事をします。つまり、AI がマシン上で何かを実行できるようになります。すでに使用しているツールに適合するものを選択してください。
Claude、Codex、opencode、Cursor、ChatGPT、およびその他の MCP クライアントの最も単純なパス。コマンドを 1 つ実行するだけで接続されます。セットアップ→
プレーンな Web リクエストを好みますか?コードとコマンドを単純な Web アドレスに送信します。 URL を呼び出すことができるもの (スクリプト、スケジュールされたジョブ、チャットボットなど) はすべて、マシンを駆動できます。 API ドキュメント →
既製のスキルを、スキルをサポートする AI エージェントにドロップします。次に、マシンコードを指定すると機能します。スキルをインストール →
画面を見るのではなく、仕事をするために作られています
AI を実際のマシン上で動作させるためのシンプルで直接的な方法: 小さなアプリをインストールし、マシン コードを使用し、要求します。
ログの確認、テストの実行、アプリの再起動、画面のないマシンの管理に最適です。
何も開いたり露出したりしないでください。
セキュリティはボルトオンではなく、コアに組み込まれています
AI にお客様のマシンへの実際のアクセスを提供するため、セキュリティが最優先されます。デフォルトで安全であり、制限については正直です。
アプリは接続するだけです。マシン上の何も受信接続を待機しません。見つけたり攻撃したりするための扉が開いておらず、ファイアウォールの内側でそのまま動作します。
当社はコードやキーを読み取り可能な形式で保管することはありません。たとえデータベースが漏洩したとしても、その中には使用できるものは何もありません。キーはいつでもキャンセルできます。
アクセスは自動的に更新され、静かに期限切れになります。キーは再度サインインするまで休止状態になるため、忘れたキーは永久に機能し続けることはできません。
コマンドとその結果はログに記録されたり保存されたりすることはありません。それらは通過して消え、各マシンは他のマシンから分離されたままになります。
新しいコードは誰でも 1 時間機能するため、簡単に始めることができます。その後は、承認されたアカウントのみが使用できるようになります —

アクセスをリセットしてクリアすることを選択するまで。
マシンが送り返すものは何であれ、従うべき新しい命令としてではなく、表示される単純な結果として扱われるため、ログ内の卑劣な行が AI を騙す可能性ははるかに低くなります。
より多くの生産性ツール、すべて AI 対応
プラグアンド.ai ↗
チーム全体のための Slack AI ボット。席ごとの料金はかかりません。
プライチャット ↗
各モデルのプライベート AI チャット。購読はありません。
cnvs.app ↗
描画、タスク、カンバン用のリアルタイム共同ホワイトボード。サインアップはありません。インスタントシェア。
maxcv.ai ↗
それぞれの仕事に合わせた履歴書。 AI スクリーナーを打ち負かす AI。
mcpfinder.dev ↗
MCP を見つける MCP。エージェント向けのオープンソース検出。
私に会いましょう ↗
会社間グループのスケジュール設定。

## Original Extract

Let Claude, Codex, ChatGPT, or any MCP/HTTP client run shell commands on machines you own. Remote access without exposed SSH, open ports, or VPN.

I built a platform allowing for instant access to remote computers from CLI tools like Claude, Codex, Opencode, or any other AI chat. There are mini apps for Windows (tray), Mac (menu bar), and Linux (CLI app) generating connection code, like TeamViewer. Comprehensive security, instant access, no firewall/networking configuration. Just ask your AI to do something on your machine, and it will do it flawlessly.

AI Commander
Use cases
Connect
Security
Docs
Sign in
Download
Ask AI to run commands
Works in minutes
No network setup
Run commands on remote computers by asking AI.
A remote shell for AI agents . Let Claude, ChatGPT, Codex, or any AI assistant run real shell commands on machines you own.
No exposed SSH, no open ports, no VPN.
Install the app, share the machine code, and ask your AI what to do. The app connects out on its own, so there are no firewall changes.
On Linux — a cloud server, home server, or Raspberry Pi — paste one line. On Mac or Windows, install the desktop app. You get a stable code like AIC-7K3P-WX9M-RTBN .
Give the code to ChatGPT, Claude, or another AI assistant, then ask in plain words. It runs the command on that machine and shows you the result.
Sign in when you want to keep machines in one place and give them names like prod-db . You do not need an account to try it.
Put the app on the machine you want to control
Pick your platform. The app gives you the machine code you will use to connect.
macOS
Windows
Linux
Download for macOS
Menu-bar app, everything included · Apple Silicon (M1–M4) · Intel (x64)
Download for Windows
Tray app, everything included · Windows x64
Runs quietly in the background · works on servers, cloud VMs, and Raspberry Pi
Works with Claude, ChatGPT, Codex, and more. You can try a new machine for the first hour without an account; sign in for ongoing access. Most AI tools connect with one command:
$
codex mcp add aicommander --url https://aicommander.dev/mcp
$
opencode mcp add aicommander --url https://aicommander.dev/mcp
$
pi install npm:@aicommander/mcp
Or just ask any AI in plain words — it figures out the rest:
use aicommander.dev to connect to AIC-XXX
Using Cursor, Windsurf, ChatGPT, or Claude Desktop instead?
All connection guides →
Each option does the same job: it lets your AI run something on your machine. Pick the one that fits the tool you already use.
The simplest path for Claude, Codex, opencode, Cursor, ChatGPT, and other MCP clients. Run one command and you're connected. Setup →
Prefer plain web requests? Send the code and your command to a simple web address. Anything that can call a URL — a script, a scheduled job, or a chatbot — can drive a machine. API docs →
Drop a ready-made Skill into any AI agent that supports skills. Then mention a machine code and it works. Install the Skill →
Built for doing the work, not watching a screen
A simple, direct way for your AI to work on a real machine: install a small app, use the machine code, and ask.
Great for checking logs, running tests, restarting an app, or managing a machine with no screen —
with nothing left open or exposed.
Security built into the core, not bolted on
We give an AI real access to your machine — so security comes first. Safe by default, and honest about the limits.
The app only reaches out — nothing on your machine waits for incoming connections. There's no open door to find or attack, and it works behind firewalls untouched.
We never keep your codes or keys in readable form. Even if our database leaked, there'd be nothing usable in it — and you can cancel a key anytime.
Access refreshes on its own and quietly expires. Keys go dormant until you sign in again, so a forgotten one can't keep working forever.
Your commands and their results are never logged or saved. They pass through and are gone — and each machine stays separate from the others.
A brand-new code works for anyone for one hour, so it's easy to get started. After that, only approved accounts can use it — until you choose to reset and clear access.
Whatever a machine sends back is treated as plain results to show you, not as new orders to follow — so a sneaky line in a log is far less likely to fool the AI.
More productivity tools, all AI-ready
plugand.ai ↗
Slack AI bot for the whole team. No per-seat fees.
plai.chat ↗
Private AI chat, every model. No subscription.
cnvs.app ↗
Real-time collaborative whiteboard for drawing, tasks & kanban. No signup. Instant share.
maxcv.ai ↗
CV tailored to each job. AI that beats the AI screeners.
mcpfinder.dev ↗
An MCP that finds MCPs. Open-source discovery for agents.
whenmeet.me ↗
Cross-company group scheduling.
