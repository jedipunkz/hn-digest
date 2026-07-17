---
source: "https://www.narendravardi.com/use-caffeinate-to-keep-claude-code-and-codex-running-on-mac/"
hn_url: "https://news.ycombinator.com/item?id=48947924"
title: "Use caffeinate to keep Claude Code and Codex running on Mac"
article_title: "Use caffeinate to keep Claude Code and Codex running on Mac"
author: "uyfyxr8"
captured_at: "2026-07-17T15:06:20Z"
capture_tool: "hn-digest"
hn_id: 48947924
score: 1
comments: 0
posted_at: "2026-07-17T14:38:19Z"
tags:
  - hacker-news
  - translated
---

# Use caffeinate to keep Claude Code and Codex running on Mac

- HN: [48947924](https://news.ycombinator.com/item?id=48947924)
- Source: [www.narendravardi.com](https://www.narendravardi.com/use-caffeinate-to-keep-claude-code-and-codex-running-on-mac/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T14:38:19Z

## Translation

タイトル: カフェインを使用して Mac でクロード コードとコーデックスを実行し続ける
説明: macOS には、長時間実行されるターミナル タスク中に Mac をスリープ状態に保つ「caffeinate」と呼ばれる小さなコマンドが付属しています。これは、スリープによってネットワーク セッションが切断され、作業が中断される可能性があるオフィス VPN 経由の Claude Code および Codex セッションに特に役立ちます。

記事本文:
サインイン
購読する
あい
カフェインを使用して Mac 上で Claude Code と Codex を実行し続ける
macOS には、長時間実行されるターミナル タスク中に Mac をスリープ状態に保つ「caffeinate」と呼ばれる小さなコマンドが付属しています。これは、スリープによってネットワーク セッションが切断され、作業が中断される可能性があるオフィス VPN 経由の Claude Code および Codex セッションに特に役立ちます。
AI コーディング ツールは、私の通常のターミナル ワークフローの一部になりつつあります。
Claude Code や Codex などのツールは、コードベースを探索したり、リファクタリングを要求したり、テストを実行したり、エージェントにタスクを実行させたりするときに役立ちます。しかし、この流れを中断する可能性がある非常に退屈な問題が 1 つあります。
オフィスの VPN に接続している場合、これはさらに苦痛になります。実行時間が少し長いタスクでは、ネットワーク接続を維持する必要があります。ラップトップがスリープ状態になると、VPN が切断される可能性があり、タスクが一時停止したり、失敗したり、手動による対応が再度必要になったりする可能性があります。
以前、ラップトップのスリープを防ぐ単純な Web ページである No Sleep Page について書きました。これは、ブラウザベースのソリューションが必要な場合にうまく機能します。
欠点の 1 つは、スリープ ページなしではブラウザのタブを前面に表示する必要があることです。ページ自体には次のように書かれています。
このタブがフォアグラウンドにある場合にのみ、コンピュータを起動状態に保つことができます。
ただし、すでにターミナル内で作業している場合は、macOS にはさらにクリーンなオプションがあります。
caffeinate は、Mac のスリープを防止する macOS コマンドライン ツールです。
macOS ではすでに利用可能であるため、新たにインストールするものは何もありません。
このワークフローでは、いくつかのオプションのみが必要です。
-i : アイドルスリープを防止します。これは私が端末コマンドで最も頻繁に使用するものです。
-d : ディスプレイがスリープ状態になるのを防ぎます。
-t <秒> : Mac を一定の秒数の間起動したままにします。
-w <pid> : 特定のプロセスが終了するまで Mac を起動したままにします。
オプションを組み合わせることができます。例えば

, -di は、ディスプレイを起動したままにし、アイドル スリープを防止することを意味します。
カフェイン酸塩
これをターミナルで実行すると、 Ctrl + C を使用してコマンドを停止するまで Mac は起動したままになります。
ただし、より便利な方法は、 caffeinate を介して別のコマンドを実行することです。
カフェイン -i NPM テスト
この場合、npm test の実行中、macOS は起動したままになります。コマンドが終了すると、caffeinate も動作を停止します。
重要なコマンドが実行されている間のみ、Mac を起動したままにしておきます。
ターミナル ウィンドウを最前面に表示しておく必要はありません。 caffeinate を使用して長時間実行されるコマンドを開始し、別のアプリに切り替えると、そのコマンドがアクティブである間、macOS は引き続きマシンを起動したままにすることができます。
クロード コードおよびコーデックスでのカフェインの使用
Claude Code と Codex は、大規模なリポジトリの読み取り、テストの実行、lint エラーの修正、ファイルのリファクタリング、プル リクエストのレビューなど、タスクが重要な場合に時間がかかることがあります。これらはまさに、アイドル スリープによってセッションが中断されることを望まない実行です。
カフェイン -i クロード
コーデックスの場合:
カフェイン -i コーデックス
ほとんどの場合、 -i から始めます。ディスプレイはオフになりますが、端末の作業は続行されます。
ディスプレイも起動したままにしたい場合は、 -di を使用します。
カフェイン - ディ クロード
カフェイン -di コーデックス
一定時間実行する
コマンドにカフェインを付加したくない場合があります。 Mac をしばらく起動したままにしておきたいだけです。
私はプラン生成後30分程度を頻繁に利用しています。これにより、エージェントは、Mac が途中でスリープ状態になることなく、実装と実行のための集中的なウィンドウを得ることができます。
カフェイン酸 -t 1800
時間は秒単位です。 1 時間の場合は、以下を使用します。
カフェイン -t 3600
バックグラウンドまたはすでに実行中のタスク
今バックグラウンド タスクを開始している場合は、caffeinate コマンド全体をバックグラウンドで実行できます。
カフェイン -i ./long-running-task.sh &
タスクがすでに実行されている場合

場合は、そのプロセス ID を見つけて、 -w を使用します。
カフェイン -i -w 12345
ここで、12345 はプロセス ID です。カフェインは、プロセスが終了するまで Mac を起動したままにします。
macOS で Claude Code または Codex を使用している場合、caffeinate は知っておく価値のある小さなコマンドの 1 つです。
新しいアプリをインストールする必要はありません。
システム設定を永続的に変更する必要はありません。
caffeinate を通じて重要なコマンドを実行し、作業が完了するまで Mac を起動したままにしておきます。
カフェイン入りの男
AI で生成されたファイルでは、HTML よりも Markdown を依然として好む理由
注目の
Chrome で Markdown ファイルを読み取る最も簡単な方法
注目の
Github でのスカッシュ マージの悪夢

## Original Extract

macOS ships with a small command called `caffeinate` that keeps your Mac awake during long-running terminal tasks. It is especially useful for Claude Code and Codex sessions over office VPN, where sleep can disconnect the network session and interrupt the work.

Sign in
Subscribe
ai
Use caffeinate to keep Claude Code and Codex running on Mac
macOS ships with a small command called `caffeinate` that keeps your Mac awake during long-running terminal tasks. It is especially useful for Claude Code and Codex sessions over office VPN, where sleep can disconnect the network session and interrupt the work.
AI coding tools are becoming part of my regular terminal workflow.
Tools like Claude Code and Codex are useful when I want to explore a codebase, ask for a refactor, run tests, or let an agent work through a task. But there is one very boring problem that can interrupt this flow:
This becomes more painful when I am connected to an office VPN. Slightly longer-running tasks need the network connection to stay alive. If the laptop sleeps, the VPN can disconnect, and the task may pause, fail, or need manual attention again.
Earlier, I had written about No Sleep Page , which is a simple webpage that prevents the laptop from sleeping. That works well when you want a browser-based solution.
One drawback is that No Sleep Page needs the browser tab to stay in the foreground. The page itself says:
This can only keep your computer awake if this tab is kept in the foreground.
But if you are already working inside the terminal, macOS has an even cleaner option.
caffeinate is a macOS command-line tool that prevents your Mac from sleeping.
It is already available on macOS, so there is nothing new to install.
For this workflow, you only need a few options:
-i : prevent idle sleep. This is the one I use most often for terminal commands.
-d : prevent the display from sleeping.
-t <seconds> : keep the Mac awake for a fixed number of seconds.
-w <pid> : keep the Mac awake until a specific process exits.
You can combine options. For example, -di means keep the display awake and prevent idle sleep.
caffeinate
Run this in Terminal, and your Mac will stay awake until you stop the command using Ctrl + C .
But the more useful way is to run another command through caffeinate .
caffeinate -i npm test
In this case, macOS will stay awake while npm test is running. Once the command exits, caffeinate also stops doing its work.
Keep the Mac awake only while the important command is running.
The Terminal window does not need to stay in the foreground. You can start a long-running command with caffeinate , switch to another app, and macOS will still keep the machine awake while that command is active.
Using caffeinate with Claude Code and Codex
Claude Code and Codex can take time when the task is non-trivial: reading a large repository, running tests, fixing lint errors, refactoring files, or reviewing a pull request. These are exactly the runs where I do not want idle sleep to interrupt the session.
caffeinate -i claude
For Codex:
caffeinate -i codex
For most cases, I would start with -i . The display can turn off, but the terminal work continues.
If you want to keep the display awake too, use -di :
caffeinate -di claude
caffeinate -di codex
Running it for a fixed time
Sometimes you do not want to attach caffeinate to a command. You just want the Mac to stay awake for some time.
I frequently use it for 30 minutes after a plan is generated. That gives the agent a focused window for implementation and execution without the Mac sleeping in the middle.
caffeinate -t 1800
The time is in seconds. For one hour, use:
caffeinate -t 3600
Background or already-running tasks
If you are starting a background task now, you can background the whole caffeinate command:
caffeinate -i ./long-running-task.sh &
If the task is already running, find its process ID and use -w :
caffeinate -i -w 12345
Here, 12345 is the process ID. caffeinate keeps the Mac awake until that process exits.
If you use Claude Code or Codex on macOS, caffeinate is one of those tiny commands worth knowing.
You do not need to install a new app.
You do not need to change system settings permanently.
Just run the important command through caffeinate , and let the Mac stay awake until the work is done.
man caffeinate
Why I still prefer Markdown over HTML for AI-generated files
Featured
The easiest way to read Markdown files in Chrome
Featured
The nightmare that is squash merge ft Github
