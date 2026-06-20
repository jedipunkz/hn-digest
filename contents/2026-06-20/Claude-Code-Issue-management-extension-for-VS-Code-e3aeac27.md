---
source: "https://www.forq.ink/"
hn_url: "https://news.ycombinator.com/item?id=48611912"
title: "Claude Code Issue management extension for VS Code"
article_title: "Forq — Run a team of agents from one editor"
author: "barakolshe"
captured_at: "2026-06-20T21:34:55Z"
capture_tool: "hn-digest"
hn_id: 48611912
score: 3
comments: 1
posted_at: "2026-06-20T18:59:05Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Issue management extension for VS Code

- HN: [48611912](https://news.ycombinator.com/item?id=48611912)
- Source: [www.forq.ink](https://www.forq.ink/)
- Score: 3
- Comments: 1
- Posted: 2026-06-20T18:59:05Z

## Translation

タイトル: Claude Code VS Code 用の問題管理拡張機能
記事のタイトル: Forq — 1 人の編集者からエージェントのチームを運営する
説明: Forq は、VS Code をネイティブ Claude CLI である Claude Code のコントロール パネルに変換します。エディター タブとしてセッションを起動、分離、再開し、それぞれが独自の git ワークツリー内ですべて同時に動作します。

記事本文:
Forq — 1 つのエディタからエージェントのチームを実行する Forq デモ 仕組み ボード 実行対象 価格 ライセンス キーの取得 VS Code 拡張機能 1 つのエディタからエージェントのチームを実行します。
Forq は、VS Code をネイティブ Claude CLI である Claude Code のコントロール パネルに変えます。エディター タブとしてセッションを起動、分離、再開し、それぞれが独自の git ワークツリー内ですべて同時に動作します。
問題を作成し、それを To Do に移動すると、Claude セッションが独自のワークツリーでスピンアップされ、それを取得します。ここではループ全体を約 1 分で説明します。
一度に 1 つのエージェントがボトルネックになります。
あなたは一度に 5 つの半分構築されたアイデアをオープンにして、単独で出荷します。 1 つのブランチで 1 つのエージェントを動作させるということは、すべてがそのブランチの前にあるものを待機することを意味します。 Forq はそれらを並行して分離して実行するため、進行は並行して行われます。
クロード コード セッションをエディター タブとして開始します。必要なだけ開いて、ペインのように機能するのを確認してください。
各セッションは独自の git ワークツリーを取得するため、エージェントがお互いのファイルやブランチをステップすることはありません。
セッションに戻って、中断したところから会話を続けてください。
タスクを To Do に移動します。エージェントが引き取ってくれます。
作業は単純なボードとして保持してください。問題が To Do に表示されると、Forq は新しいワークツリーにその問題のクロード コード セッションを生成し、エージェントが作業を終えてあなたの番になると、それを [レビュー中] に移動します。
カードを列間でドラッグし、好きな場所にドロップします。
あなたのアイデアや問題はマシン上に存在し、オフラインで作業できます。共有トラッカーをご希望ですか? Linear を接続し、双方向で同期します。
あなたのツール。あなたのアカウント。あなたのマシン。
Forq はエンジンではなくコントロール パネルです。これは、ネイティブの Claude CLI や、ユーザーがインストールして自分でサインインするその他のソフトウェアを駆動するため、各プロバイダーのアカウント、請求、および条件はユーザーとそのプロバイダーの間で維持されます。 Forq は誰かのサービスを再販する仲介業者ではありません

e またはお客様に代わって Claude サブスクリプションを使用します。
Claude Code コーディング エージェント Forq は、ネイティブの Claude CLI (Claude Code、Anthropic 製) を起動し、VS Code から駆動します。これをインストールし、自分の Anthropic アカウントでサインインするため、すべてのセッションは Anthropic との独自の契約に基づいて実行され、使用料はすべてあなたのものになります。 Forq は CLI をマシン上でローカルに実行し、Claude 資格情報やサブスクリプションをプロキシしたり、保存したり、再利用したりすることはありません。 Forq は独立した拡張機能であり、Anthropic と提携または承認されていません。
Linear issue sync オプションの Linear を接続して、マシンやチームメイト間で問題ボードを同期します。デフォルトではオフになっています。これをオンにすると、問題は自分の Linear アカウントとその条件に基づいて、マシンと Linear の間で直接移動します。オフのままにしておくと、ボードは完全にマシン上に残ります。
Dodo Payments のチェックアウトとライセンス 当社の登録販売者がチェックアウトを実行し、ライセンス キーを検証します。あなたのカードの詳細は、Dodo Payments とその処理業者に送信され、当社には決して送信されません。アクティベーションによりキーが 1 台のマシンにバインドされ、オフラインで動作し続けます。
Google Gemini AI 問題のタイトル オプション 独自の Google Gemini API キーを追加すると、Forq が問題のタイトルを提案できます。あなたがタイトルを付けた問題の説明のみが、独自のキーと Google の規約に基づいて Google の Gemini API に送信されます。キーを設定しないままにすると、タイトルは完全にオフラインのままになります。
Forq は、独立した VS Code 拡張機能です。 Claude、Anthropic、Linear、Google Gemini、および Visual Studio Code はそれぞれの所有者の商標であり、ここでは Forq の動作を説明するためにのみその名前が付けられています。
Forq は有料の VS Code 拡張機能です。一度お支払いいただくと、一生あなたのものになります。1 つのライセンスで 1 台のマシンがアクティベートされ、Dodo Payments を通じて検証されます。
$20 の 1 回限り・生涯ライセンス、サブスクリプションなし
はい。 Forq は VS 内でネイティブの Claude CLI (Claude Code) を駆動します

コードなので、インストールして自分の Anthropic アカウントでサインインする必要があります。
Forq は、ネイティブの Claude CLI をローカルで実行し、ターミナルで行うのと同じ方法でコマンドのみを送信します。自分の Anthropic アカウントでサインインするため、使用は Anthropic との契約に基づいて行われます。 Forq は、お客様の Claude 認証情報やサブスクリプションをプロキシ、保存、または再利用することはありません。また、独立した製品であり、Anthropic と提携または承認されていません。
はい。アイデアと問題はプレーンな JSON としてローカルに保存されるため、ボードは接続なしで実行されます。リニア同期はオプションです。
1つ。ライセンスは、Dodo Payments を通じて 1 台のマシンにバインドされます。もっと座席が必要ですか?追加のキーを購入します。
Forq によるものではありません。ワークツリー、問題、セッションはマシン上に残ります。 Forq が送信するのは、ライセンス チェックと有効にした同期だけです。あなたが実行する Claude Code は、ターミナルで使用するのと同じように、あなたが与えたものをあなた自身のアカウントで Anthropic に送信します。
Forq 並行して出荷するビルダー向けの鋭利なツール。

## Original Extract

Forq turns VS Code into a control panel for Claude Code, the native Claude CLI. Launch, isolate, and resume sessions as editor tabs, each in its own git worktree, all working at once.

Forq — Run a team of agents from one editor Forq Demo How it works The board What it runs on Pricing Get license key VS Code extension Run a team of agents from one editor .
Forq turns VS Code into a control panel for Claude Code, the native Claude CLI. Launch, isolate, and resume sessions as editor tabs, each in its own git worktree, all working at once.
Create an issue, move it to To Do, and a Claude session spins up in its own worktree to take it. Here’s the whole loop in about a minute.
One agent at a time is the bottleneck.
You ship alone, with five half-built ideas open at once. Working one agent on one branch means everything waits on the thing in front of it. Forq runs them side by side, isolated, so progress happens in parallel.
Start a Claude Code session as an editor tab. Open as many as you need and watch them work like panes.
Each session gets its own git worktree, so agents never step on each other’s files or branches.
Come back to any session and keep the conversation going from exactly where it left off.
Move a task to To Do. An agent picks it up.
Keep your work as a simple board. The moment an issue lands in To Do, Forq spawns a Claude Code session for it in a fresh worktree, then moves it to In Review when the agent is done and it’s your turn to look.
Drag a card between columns and drop it wherever you like.
Your ideas and issues live on your machine and work offline. Prefer a shared tracker? Connect Linear and sync both ways.
Your tools. Your accounts. Your machine.
Forq is the control panel, not the engine. It drives the native Claude CLI and other software you install and sign in to yourself, so each provider’s account, billing, and terms stay between you and that provider. Forq is not a middleman reselling anyone’s service or using your Claude subscription on your behalf.
Claude Code coding agent Forq launches the native Claude CLI (Claude Code, by Anthropic) and drives it from VS Code. You install it and sign in with your own Anthropic account, so every session runs under your own agreement with Anthropic and any usage costs are yours. Forq runs the CLI locally on your machine and never proxies, stores, or reuses your Claude credentials or subscription. Forq is an independent extension, not affiliated with or endorsed by Anthropic.
Linear issue sync Optional Connect Linear to sync your issue board across machines and teammates. It is off by default; when you turn it on, issues move directly between your machine and Linear under your own Linear account and their terms. Leave it off and your board stays entirely on your machine.
Dodo Payments checkout & licensing Our merchant of record runs checkout and validates your license key. Your card details go to Dodo Payments and their processors, never to us. Activation binds a key to one machine and then keeps working offline.
Google Gemini AI issue titles Optional If you add your own Google Gemini API key, Forq can suggest issue titles. Only the issue description you are titling is sent to Google's Gemini API, under your own key and Google's terms. Leave the key unset and titling stays fully offline.
Forq is an independent VS Code extension. Claude, Anthropic, Linear, Google Gemini, and Visual Studio Code are trademarks of their respective owners, named here only to describe what Forq works with.
Forq is a paid VS Code extension. Pay once and it’s yours for life: one license activates one machine, validated through Dodo Payments.
$20 one-time · lifetime license, no subscription
Yes. Forq drives the native Claude CLI (Claude Code) inside VS Code, so you’ll need it installed and signed in with your own Anthropic account.
Forq runs the native Claude CLI locally and only sends it commands, the same way you would in a terminal. You sign in with your own Anthropic account, so your usage stays under your own agreement with Anthropic. Forq never proxies, stores, or reuses your Claude credentials or subscription, and is an independent product, not affiliated with or endorsed by Anthropic.
Yes. Ideas and issues are stored locally as plain JSON, so the board runs with no connection. Linear sync is optional.
One. A license binds to a single machine through Dodo Payments. Need more seats? Buy additional keys.
Not by Forq. Worktrees, issues, and sessions stay on your machine; the only things Forq sends are license checks and any sync you turn on. Claude Code, which you run, sends whatever you give it to Anthropic under your own account, the same as using it in a terminal.
Forq A sharp tool for builders who ship in parallel.
