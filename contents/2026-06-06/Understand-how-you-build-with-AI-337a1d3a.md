---
source: "https://paxel.ycombinator.com/"
hn_url: "https://news.ycombinator.com/item?id=48419288"
title: "Understand how you build with AI"
article_title: "Understand how you build with AI."
author: "simonpure"
captured_at: "2026-06-06T00:55:09Z"
capture_tool: "hn-digest"
hn_id: 48419288
score: 3
comments: 0
posted_at: "2026-06-05T22:42:32Z"
tags:
  - hacker-news
  - translated
---

# Understand how you build with AI

- HN: [48419288](https://news.ycombinator.com/item?id=48419288)
- Source: [paxel.ycombinator.com](https://paxel.ycombinator.com/)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T22:42:32Z

## Translation

タイトル: AI を使用して構築する方法を理解する
記事のタイトル: AI を使用して構築する方法を理解する。

記事本文:
取得
どのように
よくある質問
サインイン
[
アップロード
】
人間
機械
では、建築業者は現在どのように建物を建てているのでしょうか?
私たちは何か奇妙なことが起こっていることに気づきました。私たちは皆、AI を使ったソフトウェアを作成していますが、ほとんどは単独で行っており、他の人がどのようにやっているのか全く分かりません。
そこで私たちは、AI を使用して構築する方法を理解するのに役立つツールを作成しました。
Claude、Codex、および Cursor セッションを読み取るため、構築方法について発見できます。
自分のセッションで実行して調べてください。そして、時間の経過とともに、より多くの人が自分の作品をアップロードするにつれて、他のビルダーとの比較を示すことができるようになります。
これまでに 126,496 のセッションがアップロードされ、分析されました。
ここでは、人々がコーディングの習慣について学んだことの例を示します...
まず計画を立て、決定を体系化して、複合的な足場を構築します。
セッションの 72% で Opus 4.7 に到達し、22% で GPT-5.5 に到達します。
コミットの 70% は午後 10 時から午前 2 時の間に到着し、午後 11 時頃にピークに達します。
コードを記述する前にプラン モードで開き、クイック フィックスではコードをスキップします。
11 の異なるセッションで最もよく使われたフレーズ。
4 つのリポジトリで一度に最大 6 つのコーディング エージェントを実行しました。
プロンプトの 80% は 10 単語未満です。ちょっとしたことでたくさんのことを言います。
「これではなく他のことがうまくいかなかったのですか」
これを午前 1 時に、コンテキストなしで送信しました。エージェントはどういうわけかあなたの意味を理解しました。
あなたはエージェントに対してどれくらい礼儀正しいですか?
あなたは先月、エージェントに 84 回感謝の意を表しました。ロボットが引き継いだとき、彼らはあなたのことを懐かしく思い出すでしょう。
どのくらいの頻度でコースを変更しますか？
エージェントをタスクの途中で停止し、リダイレクトします。プロンプトは 10 回中 4 回程度です。
エージェントの最長勤務期間はどれくらいですか?
あるエージェントは、あなたが介入するまで、不安定なテスト スイートを追いかけるのに 4 時間以上費やしました。
あなたの最大のクラッシュは何ですか?
「文字通り、そのファイルには触らないでくださいと言ったんだ」
3 回目に無視された後、caps loc

kはオンになり、決して後戻りしませんでした。
先月のコミットは 420 件、プル リクエストは 35 件ありました。
コードを書く前にアイデアをぶつけ、反対を求め、意見が合わない場合は方針を変更します。ツールが少なくなり、チームメイトが増えます。
今月最大のプッシュは火曜日に届きました。
これはYCの実験です。コーディング エージェントを使用して適切に構築することが何を意味するのかはまだ誰も知りませんが、私たちはそれを解明しようと努めています。
コーディング セッションをアップロードする 2 つの方法
コンピューター上の AI セッションのトランスクリプトを調べます。コマンドを実行する場所によって、すべてのプロジェクトを一度に分析するか、1 つのプロジェクトだけを分析するかが決まります。
オプション 1: すべてのリポジトリ (推奨)
プロジェクト全体にわたるより広い視野に最適です。リポジトリを保持する親フォルダーに移動して、実行します。
単一のプロジェクトに集中する場合に最適です。そのプロジェクトのフォルダーに移動し (~/path/to/your-project を実際のパスに置き換えます)、実行します。
Claude、Codex、または Cursor でこのプロンプトを使用すると、AI トランスクリプトを含むマシン上のすべてのリポジトリを検索し、リストを表示し、選択したコマンドに対してすぐに実行できるコマンドを返すことができます。
ビルダーのプロフィール、つまりステアリング、実行、エンジニアリング、製品の直感、計画の 5 つの側面にわたって AI をどのように扱うかのスナップショットを取得できます。
私たちは、アーキテクト、クオリティ ガーディアン、ベロシティ マシン、ナイト オウルなど、セッションで見るアーキタイプに名前を付けます。また、トランスクリプト内の実際のやりとりから抽出された、AI を指示するときに行う決定パターン、つまり特徴的な動きも抽出します。
次に、一般的なアドバイスではなく、実際のセッションに基づいて、あなたの成長のエッジ、次に試すべき具体的なことをいくつか示します。
電子メールでサインインし、分析したいリポジトリからターミナルでコマンドを実行します。分析は自分のマシン上の Docker で実行されるため、Docker がインストールされている必要があります。

dが最初に実行します。約 15 ～ 30 分で、スコア、アーキタイプ、意思決定パターン、成長エッジを含むプロファイルが返されます。
作業ツリー、.env ファイル、生のトランスクリプトはマシン上に残ります。クロードまたは GPT に送られるのは短いトランスクリプトの抜粋だけなので、要約してスコアを付けることができます。YC に届くのは、スコア、ナラティブ、編集された決定の小さなペイロードです。完全な内訳を読むことができます。
作業ツリーと .env ファイルは YC にアップロードされません。分析中、トランスクリプトの抜粋 (プロンプト、エージェントの応答、ツール呼び出しのスニペット) がプロキシ経由でクロードまたは GPT に送信されるため、何が起こったかを要約できます。実行の最後に、スコア、ナラティブ、編集された決定、セッション メタデータの JSON ペイロードが YC にアップロードされます。完全な内訳。
どのようなコーディング ツールで動作しますか?
クロード コード、コーデックス CLI、およびカーソル。各ツールがローカルに保存しているトランスクリプトを読み取ります。
はい。分析は、自分のマシン上の Docker コンテナーでローカルに実行されます。 Docker Desktop (または、colima や OrbStack などの Docker ランタイム) をインストールし、コマンドを実行する前にそれが実行されていることを確認します。エンジンが起動していない場合、スクリプトはエンジンの起動方法を正確に出力します。
セッション数に応じて15〜30分。準備が完了するとメールが届きます。
複数のリポジトリからアップロードできますか?
はい。親ディレクトリから実行し、含めるリポジトリを選択するか、 --project NAME を渡します。
プロフィールを更新したいときはいつでも。アップロードするたびに新しいレポートが作成されます。プロファイルはそれらすべてにわたって集約されます。
レポートとプロファイルの違いは何ですか?
レポートは 1 回のアップロードのスナップショットです。プロファイルは、すべてのレポートにわたるパターン、つまり、どのように進化してきたか、何が一貫しているか、どこで成長しているかなどを総合します。
複数のマシンでコーディングを行っています。両方の単位を取得できますか?
はい。同じカール通信を実行するだけです

そして各マシン上で。同じ電子メールでログインすると、ビルダー プロファイルによってすべてのセッションが自動的に集約されます。
Paxel トークンを Startup School アプリケーションに添付するにはどうすればよいですか?
Paxel トークンをリンクすると、Startup School 2026 に参加できる可能性が高まります。レポート ページを開いてトークンを見つけます。まだ申請していない場合、または待機リストに残っている場合は、添付リンクをクリックしてトークンをアプリケーションに直接送信するか、トークンをコピーして自分で貼り付けます。
バグを見つけたか、機能リクエストがあります。どこに送ればいいですか？
バグを報告したり、機能をリクエストしたりするには、このフォームを使用してください。気に入った機能リクエストについては、コーディング エージェントを使用してプロンプトを実行し、結果をマージします。
バグを報告するか、機能リクエストを送信してください
何かが壊れているのを見つけた場合、またはこのツールにまだ実行できないことを実行してほしい場合は、下記までお知らせください。気に入った機能リクエストについては、コーディング エージェントを使用してプロンプトを実行し、結果をマージします。
リポジトリの 1 つで試してみると、約 20 分で結果が得られます。YC がコストを負担し、コードがマシンから離れることはありません。アップロード

## Original Extract

Get
How
FAQ
Sign in
[
Upload
]
Human
Machine
So, how do builders build now?
We noticed that something strange is happening. We're all writing software with AI, but we are mostly doing it alone, with no real sense of how anyone else does it.
So we made a tool to help you understand how you build with AI.
It reads your Claude, Codex, and Cursor sessions, so you can discover things about how you build.
Run it on your own sessions and find out. And with time, as more people upload theirs, we'll be able to show you how you compare to other builders.
So far, 126,496 sessions have been uploaded and analyzed.
Here are examples of what people have learned about their coding habits...
You plan first, codify your decisions, and build scaffolding that compounds.
You reach for Opus 4.7 in 72% of your sessions, GPT-5.5 in 22%.
70% of your commits land between 10 PM and 2 AM, peaking around 11 PM.
You open in plan mode before writing code, skipping it on quick fixes.
Your most-used phrase, across 11 different sessions.
You've run as many as 6 coding agents at once across 4 repos.
80% of your prompts are under 10 words. You say a lot with a little.
"make teh other thing wrok not this 1"
You sent this at 1 AM with zero context. The agent somehow figured out what you meant.
How polite are you to your agents?
You thanked your agents 84 times in the last month. When the robots take over, they'll remember you fondly.
How often do you change course?
You stop and redirect the agent mid-task rather than let it run off, about 4 prompts in 10.
What's your longest agent run?
One agent spent over four hours straight chasing a flaky test suite before you stepped in.
What's your biggest crash out?
"I LITERALLY SAID DONT TOUCH THAT FILE"
After the third time it ignored you, the caps lock came on and never went back off.
Across 420 commits and 35 pull requests in the last month.
You bounce ideas before writing code, ask for pushback, and change course when it disagrees. Less tool, more teammate.
Your single biggest push of the month landed on a Tuesday.
This is an experiment from YC. Nobody really knows yet what it means to build well with coding agents, and we are trying to find out.
Two ways to upload coding sessions
It looks at the AI session transcripts on your computer. Where you run the command tells it whether to analyze every project at once or just one.
Option 1: All my repos (Recommended)
Best for a broader picture across projects. Change into the parent folder that holds your repos, then run.
Best for focusing on a single project. Change into that project's folder (replace ~/path/to/your-project with the real path) and run.
You can use this prompt in Claude, Codex, or Cursor to find all your repos on your machine with AI transcripts, show you the list, and hand back ready-to-run commands for the ones you pick.
You get a builder profile, a snapshot of how you work with AI across five dimensions of steering, execution, engineering, product instinct, and planning.
We name the archetypes we see in your sessions, whether you build like an Architect, a Quality Guardian, a Velocity Machine, or a Night Owl. We also pull out your decision patterns, the signature moves you make when directing the AI, drawn from real exchanges in your transcripts.
Then we point at your growth edge, a few specific things to try next, grounded in your actual sessions rather than generic advice.
Sign in with your email, then run the command in your terminal from the repo you want analyzed. The analysis runs in Docker on your own machine, so you'll need Docker installed and running first. In about 15 to 30 minutes you get your profile back with scores, archetypes, decision patterns, and a growth edge.
Your working tree, your .env files, and your raw transcripts stay on your machine. Only short transcript excerpts go to Claude or GPT so we can summarize and score, and what reaches YC is a small payload of scores, narratives, and redacted decisions. You can read the full breakdown .
Your working tree and .env files are not uploaded to YC. During analysis, transcript excerpts (prompts, agent responses, and snippets of tool calls) are sent to Claude or GPT via our proxy so we can summarize what happened. At the end of the run, a JSON payload of scores, narratives, redacted decisions, and session metadata is uploaded to YC. Full breakdown .
What coding tools does it work with?
Claude Code, Codex CLI, and Cursor. We read the transcripts each tool stores locally.
Yes. The analysis runs locally in a Docker container on your own machine. Install Docker Desktop (or any Docker runtime like colima or OrbStack) and make sure it's running before you run the command. If the engine isn't up, the script prints exactly how to start it.
15-30 minutes depending on how many sessions you have. You get an email when it's ready.
Can I upload from multiple repos?
Yes. Run from a parent directory and pick which repos to include, or pass --project NAME .
Whenever you want an updated profile. Each upload produces a new report; your profile aggregates across all of them.
What's the difference between a report and a profile?
A report is a snapshot of one upload. Your profile synthesizes patterns across all your reports: how you've evolved, what's consistent, where you're growing.
I code on multiple machines. Can I get credit for both?
Yes. Just run the same curl command on each machine. Log in with the same email and your builder profile will automatically aggregate sessions across all of them.
How do I attach my Paxel token to my Startup School application?
Linking your Paxel token boosts your chances of getting into Startup School 2026. Open your reports page and find your token. If you haven't applied yet or you're still on the waitlist, click the attach link to send your token straight to your application, or copy the token and paste it in yourself.
I found a bug or have a feature request. Where do I send it?
Use this form to report a bug or request a feature. For feature requests we like, we'll run your prompt with a coding agent and merge the result.
Report a bug or submit a feature request
Tell us below if you found something broken or want this tool to do something it can't yet. For feature requests we like, we'll run your prompt with a coding agent and merge the result.
Try it on one of your repos and you'll have results in about 20 minutes, with YC covering the cost and your code never leaving your machine. Upload
