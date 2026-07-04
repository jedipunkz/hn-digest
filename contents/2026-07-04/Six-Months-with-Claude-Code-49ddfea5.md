---
source: "https://sajith.me/blog/2026/six_months_with_claude_code.html"
hn_url: "https://news.ycombinator.com/item?id=48784795"
title: "Six Months with Claude Code"
article_title: "SE: Six Months with Claude Code"
author: "sajithdilshan"
captured_at: "2026-07-04T13:00:49Z"
capture_tool: "hn-digest"
hn_id: 48784795
score: 1
comments: 0
posted_at: "2026-07-04T12:04:24Z"
tags:
  - hacker-news
  - translated
---

# Six Months with Claude Code

- HN: [48784795](https://news.ycombinator.com/item?id=48784795)
- Source: [sajith.me](https://sajith.me/blog/2026/six_months_with_claude_code.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T12:04:24Z

## Translation

タイトル: クロード・コードとの6か月
記事のタイトル: SE: クロード コードとの 6 か月

記事本文:
← すべての投稿
パソコン
PyCharm
🗁 python.py
📁 エージェントライブラリ
📁 コマンド
📁プラグイン
📁スキル
python.py
utils.py
1
2
3
12
13
14
def テスト関数 ():
def __init__ (自分自身):
print( "こんにちはを受け取りました" )
def _close (自分自身):
out = "コードを書くのは丁寧です"
return _.close(cli)
エージェント-cli-プラグイン
あ
クロード・コード
Jira チケットを実装する準備ができました。
あ
クロード
現在、実装計画を検討中です。
メッセージを入力してください...
2026 年 6 月 24 日 · 5 分で読めます
Claude コードを使用した 6 か月間: 私の開発者のワークフローはどう変化したか
2025 年 10 月以来、私は仕事でクロード コードを広範囲に使用しています。初めの頃は確かにあったよ
学習曲線があり、それに合わせてワークフローを調整する必要がありました。 Kotlin 開発者として働いていたときのことを思い出します
以前の会社では、主に IntelliJ IDEA のオートコンプリートに GitHub Copilot を使用していました。その後、私が使用したのは、
ChatGPT を使用して小さなコード ブロックとスクリプトを生成しましたが、ほとんどの場合、依然として自分で手動でコーディングしていました。
しかし、Claude Opus のリリースにより、すべてが変わりました。そのモデルはとても印象的だったので、私は少しずつ
ブレーンストーミング、調査、取り組みをより小さなタスクに分割し、作成し、
Jira チケットの調整、コードの実装とテスト、PR の自動オープン、さらにはコード レビューの処理も行います。
ただし、本当に気になったのは、Anthropic をシームレスに使用できる統合 IDE が存在しないことでした。
モデル。ターミナルで Cursor、Zed、および Claude コードを試してみましたが、Python 開発者として自分自身に気づきました
PyCharm と他のツールの間を常に行ったり来たりする必要があります。 JetBrains に慣れているからかもしれません
IDE や PyCharm と比較すると、他のエディターはかなり不十分だと感じました。
私は以前、以前の会社で IntelliJ IDEA のプラグインやカスタム ディストリビューション全体を開発したことがありました。
すでにひざまずいている

IntelliJ プラグインの作成方法。それで、私は書き始めました
エージェント-cli-プラグイン 。
Claude Code が私の「コード モンキー」として機能したため、わずか数時間で MVP を思いつくのが信じられないほど速くなりました。
機能を追加し、時間の経過とともに他の CLI エージェントのサポートを拡張してきましたが、現在はそれをプライマリとして使用しています
PyCharm 内のクロード コードと対話するためのツール。私の主な目標は、適切なレンダリングを行うことでした。私の意見では、それは
PyCharm のデフォルトのターミナルでクロード コードを実行するよりもはるかに優れています。
日常のプロセスに関して言えば、私は仕事のほぼすべての側面で Claude Code を使用しています。最近、私は
次のワークフローに落ち着きました。
ブレインストーミングとリサーチ
特定のイニシアチブのすべての要件を収集し、すべての要件を含む Markdown ファイルを作成します。
情報。このファイルは、最終的に LLM で生成されたコンテンツと手動調整が混在したものとなり、
イニシアチブ全体の中核となる知識ベース。
内訳とチケットの作成
私は、以下に基づいて、イニシアチブを実装とテストの準備ができた小さな部分に分割します。
コードベースを保存し、別の Markdown ファイルに小さな Jira チケットとして保存します。カスタムクロードを作りました
これらのチケットを必要に応じて正確にフォーマットするスキル。それらを一つ一つ見直して改良した後、
Claude Code では、別のスキルを使用して、Atlassian MCP 経由で Jira に直接公開します。
チームの連携
チケットをプルする前に、さらなる改良と見積もりのためにチームにチケットを提示します。
スプリントに入る。
実施計画
チケットを受け取ると、専用のクロード コード スキルがチケットを取得し、詳細なチケットを生成します。
Markdown ファイル内の実装計画。 Jira チケットは非常に明確に指定されているため、モデルは
ほとんどの場合、計画は正しく行われるため、変更を加える必要はほとんどありません。
サンドボックス実行
サンドボックス化されたクロードで計画を実行します

コードセッション。小さなものを作りました
クロードクレート
最低限のツールとクロード コードを含む Docker イメージ。プロジェクト ディレクトリをコンテナーにマウントします。
スタートアップ。このコンテナ内では、クロード コードが --dangerously-skip-permissions フラグを使用して実行されます。
PyCharm 内で Agent-cli-plugin を使用してこのサンドボックスを起動します。なぜなら、ホストマシンの
コードベースがマウントされているので、PyCharm 内でファイルの変更がリアルタイムで発生するのを監視できます。最も良い点は、
コードを実装し、テストを書き、タスクを完了している間、私はクロードなしで別のことに集中できます。
許可を求めるために常に私の邪魔をします。破壊的な変更はコンテナに隔離され、
最悪の場合、クリーンな Git チェックアウトを実行するだけで済みます。このサンドボックス モードでは、MCP や追加のものが存在しないことに注意してください。
スキルが設定されています。これは純粋に必要最低限​​のクロード コード セッションです。
レビューとPR
実装が完了したら、別のスキルを使用して変更点のコードレビューを実行します。
ホストマシン上で作成されます。微調整が必要な場合は、実装計画に使用したのと同じセッションを使用します
ホスト上で。最後に、GitHub MCP を利用したスキルを使用して PR をスピンアップします。
このコア ワークフロー以外にも、テスト、PR コード レビュー、アラートのデバッグなどに関するさまざまなスキルを持っています。と
これらすべてを導入すると、生産性が簡単に 2 倍になったように感じます。最近、実験的にも構築しました
ジャービスという名のパーソナルアシスタントエージェント。 Gmail、Slack、GitHub、Atlassian、Google カレンダーにアクセスできます
事前設定された権限に基づいて簡単なタスクを実行します。長期記憶と
インタラクティブなチャット。これらの MCP 経由でデータを定期的にポーリングし、メモリを更新するため、実際に次のように使用できます。
アシスタントが情報を取得したり、過去のメールからコンテキストを掘り起こしたり、古い Slack ディスカッションを参照したりできます。それは
まだ初期段階です

しかし、長く使用するほど、よりパーソナライズされ、コンテキストを認識できるようになります。
思い返してみると、これらすべてが1年も経たないうちに起こったことは驚くべきことです。すでにいくつかのアイデアがあります
エージェントのユースケースで私の生産性がさらに向上します。正直、最終的にはどうなるだろうかと思います。
今年中に、あるいは来年の夏までに。

## Original Extract

← All posts
PC
PyCharm
🗁 python.py
📁 agent_lib
📁 commands
📁 plugins
📁 skills
python.py
utils.py
1
2
3
12
13
14
def test_function ():
def __init__ (self):
print( "Received hello" )
def _close (self):
out = "Writing code is neat"
return _.close(cli)
agent-cli-plugin
A
Claude Code
Ready to implement your Jira tickets.
A
CLAUDE
Reviewing the implementation plan now.
Type a message...
Jun 24, 2026 · 5 min read
Six Months with Claude Code: How My Developer Workflow Changed
Since October 2025, I have been using Claude Code extensively for work. In the beginning, there was definitely
a learning curve, and I had to adjust my workflow around it. I remember back when I worked as a Kotlin developer
at my previous company, I used GitHub Copilot primarily for autocomplete in IntelliJ IDEA. Later on, I used
ChatGPT to generate small code blocks and scripts, but most of the time, I was still coding manually by myself.
But with the release of Claude Opus, everything changed. The model was so impressive that, little by little, I
started using it for brainstorming, researching, breaking down initiatives into smaller tasks, creating and
refining Jira tickets, implementing and testing code, automatically opening PRs, and even handling code reviews.
However, one thing that really bothered me was the lack of a unified IDE where I could seamlessly use Anthropic
models. I tried Cursor, Zed, and Claude Code in the terminal, but as a Python developer, I found myself
constantly switching back and forth between PyCharm and other tools. Maybe it's because I'm so used to JetBrains
IDEs and compared to PyCharm, other editors felt quite inadequate.
Having previously developed plugins and even a whole custom distribution of IntelliJ IDEA at a former company, I
already knew how to write an IntelliJ plugin. So, I embarked on writing the
agent-cli-plugin .
With Claude Code acting as my "code monkey," it was incredibly fast to come up with an MVP in just a few hours.
I've added functionalities and extended it support other CLI agents over time, but now I use it as my primary
tool to interact with Claude Code inside PyCharm. My main goal was to get proper rendering, and in my opinion, it
looks way better than running Claude Code in PyCharm's default terminal.
When it comes to my day-to-day processes, I use Claude Code for almost every aspect of my work. Recently, I've
settled into the following workflow:
Brainstorm & Research
I gather all requirements for a given initiative and create a Markdown file with all the
information. This file ends up being a mix of LLM generated content and manual adjustments, serving as the
core knowledge base for the entire initiative.
Breakdown & Ticket Creation
I break down the initiative into smaller, implementation and testing ready chunks based on
the codebase, saving them as smaller Jira tickets in another Markdown file. I've built a custom Claude
skill to format these tickets exactly how I need them. After reviewing and refining them one by one via
Claude Code, I use another skill to publish them directly to Jira via the Atlassian MCP.
Team Alignment
I present the tickets to my team for further refinement and estimation before pulling them
into the sprint.
Implementation Planning
Once I pick up a ticket, a dedicated Claude Code skill fetches it and generates a detailed
implementation plan in a Markdown file. Because the Jira tickets are so well specified, the model gets the
plan right most of the time, and I rarely have to make any modifications.
Sandboxed Execution
I execute the plan in a sandboxed Claude Code session. I created a small
claude-crate
Docker image containing bare minimal tools and Claude Code, mounting the project directory at container
startup. Inside this container, Claude Code runs with the --dangerously-skip-permissions flag.
I use my agent-cli-plugin within PyCharm to launch this sandbox. Because the host machine's
codebase is mounted, I can watch file changes happen in real time inside PyCharm. The best part is that
while it implements code, writes tests, and finishes the task, I can focus on something else without Claude
constantly interrupting me for permissions. Any destructive changes are isolated to the container, and
worst case scenario, I can just do a clean git checkout. Note that in this sandbox mode, no MCPs or extra
skills are configured. It's a purely bare bones Claude Code session.
Review & PR
Once the implementation is done, I use another skill to run a code review on the changes
made on the host machine. If tweaks are needed, I use the same session I used for the implementation plan
on the host. Finally, I spin up a PR using a skill powered by the GitHub MCP.
Beyond this core workflow, I have various skills for testing, PR code reviews, debugging alerts, and more. With
all of this in place, I feel like my productivity has easily doubled. Recently, I even built an experimental
personal assistant agent named Jarvis. It has access to my Gmail, Slack, GitHub, Atlassian, and Google Calendar
to perform simple tasks based on pre configured permissions. I've also integrated long-term memory and an
interactive chat. Because it periodically polls data via these MCPs and updates its memory, I can truly use it as
an assistant to retrieve information, dig up context from past emails, or reference old Slack discussions. It's
still in the early stages, but the longer I use it, the more personalized and context aware it becomes.
Thinking back, it's mind blowing that all of this happened in less than a year. I already have a few more ideas
for agentic use cases to push my productivity even further, and I honestly wonder where we will all be by the end
of this year, or even by next summer.
