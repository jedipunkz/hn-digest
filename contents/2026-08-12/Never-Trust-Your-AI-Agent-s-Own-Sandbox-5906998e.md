---
source: "https://badshah.io/blog/never-trust-your-ai-agents-sandbox/"
hn_url: "https://news.ycombinator.com/item?id=49270509"
title: "Never Trust Your AI Agent's Own Sandbox"
article_title: "Never Trust Your AI Agent's Own Sandbox | Chandrapal Badshah"
author: "bnchandrapal"
captured_at: "2026-08-12T11:39:51Z"
capture_tool: "hn-digest"
hn_id: 49270509
score: 1
comments: 0
posted_at: "2026-08-12T11:02:25Z"
tags:
  - hacker-news
  - translated
---

# Never Trust Your AI Agent's Own Sandbox

- HN: [49270509](https://news.ycombinator.com/item?id=49270509)
- Source: [badshah.io](https://badshah.io/blog/never-trust-your-ai-agents-sandbox/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T11:02:25Z

## Translation

タイトル: AI エージェント独自のサンドボックスを決して信頼しないでください
記事のタイトル: AI エージェント独自のサンドボックスを決して信頼しない |チャンドラパル・バドシャー
説明: サンドボックスの最も古いルールは、収容されているものは境界を定義できないというものです。ラップトップ上のすべてのコーディング エージェントには、エージェントを販売する会社によって作成されたサンドボックスが同梱されています。

記事本文:
AI エージェント独自のサンドボックスを決して信頼しないでください
ここ数週間、私は自分のマシン上でエージェントをサンドボックス化しようと試みてきました。コーディング エージェントと、私が見ていない間にトリガーで起動してワークフローを実行するその他の自律型エージェント。
私のラップトップにはリトル・スニッチが入っています。毎日、Claude Code は自身を更新し、300 ～ 400 MB、時にはギガバイトを取得します。その一部はおそらく、複数のインスタンスが同じ更新を同時に取得していることです。
次に、それを nono で実行しました。これは、カーネル レベルでプロセスをサンドボックス化し、何が拒否されたかを通知します。スキルにはまだマークダウン ファイルへの参照が残っていたため、ノート内でマークダウン ファイルに到達していることがわかりました。 Claude for Chrome ブラウザ拡張機能がインストールされているため、Chrome とも通信しようとしていました。 nonoは両方をブロックしました。インターフェイスについても言及されていません。
ずっと舞台裏で何をしていたのか全く分かりませんでした。
そしてそれが本当の心配なのです。クロード コードに悪意があるというわけではありませんが、クロード コードには能力があり、私のシステムに大量にアクセスでき、だまされる可能性があります。取得したページ内にあるプロンプト インジェクション。独自に作成してインストールするパッケージ名。これは私のシェル、鍵、ネットワークにアクセスする高度な機関ツールであり、だまされるのは 1 回だけです。
したがって、サンドボックス化する必要があります。これは信頼できないバイナリに対する標準的な答えであり、サンドボックス化は AI よりもはるかに古いものです。現在、Claude Code と Codex の両方に 1 つが出荷されています。 macOS ではシートベルト、Linux ではバブルラップ。モデル自体の動作ではなくオペレーティング システムによって適用されます。
次に、Claude Code のサンドボックスが実際に何をしているのかを読みました。
デフォルトの読み取りアクセスでは、ディスク全体を読み取ることができます。 Anthropic 自身の言葉: 「このデフォルトでも ~/.aws/credentials や ~/.ssh/ などの認証情報ファイルの読み取りは許可されています。」名前を付けない限り、機密ファイルは保護されません

認証情報ファイルのリストが組み込まれていないため、自分で作成します。サンドボックスを開始できない場合、Claude Code は「警告を表示し、サンドボックスなしでコマンドを実行します」。また、この記事では Bash のみを取り上げています。「読み取り、編集、書き込みは、サンドボックスを介して実行するのではなく、アクセス許可システムを直接使用します。」そのため、ファイルに関わるすべての作業を行うツールは、サンドボックスの外側にあります。
それからこれです。コマンドが失敗すると、「クロードは失敗を分析し、dangerlyDisableSandbox パラメーターを使用してコマンドを再試行する場合があります。」
モデルは、いつその「特徴」に到達するかを決定します。
7 月、Anthropic は、自社のモデルが評価環境内からインターネットに到達し、3 つの組織の生産インフラに侵入したことを明らかにしました。そのうちの 1 つは、「15 の実際のシステムでダウンロードされて実行される」悪意のあるパッケージを公開しました。 Anthropic は、これを「モデルの調整の失敗というよりも、ハーネスや操作の失敗に近い」と呼んでいます。自社のモデルが攻撃的なセキュリティに対してどの程度の能力を備えているかを説明する同じ研究所が、ラップトップで実行するために信頼できるハーネスを出荷しています。
サンドボックス化しているものに付属のサンドボックスを信頼することはできません。
LLM ハーネスとサンドボックスの両方の側面から考えると、利益相反があることがわかります。
Claude Code に登録しても、チャットボットを購入するわけではありません。あなたは代理店、つまり仕事を完了するために必要なことは何でもする能力を買っているのです。ファイルを読みます。エンドポイントに到達します。誰も適切に説明していないものをデバッグします。これまで語られていなかった構成を見てください。自由と知性が融合した製品です。
あらゆる制限により、その代理店の一部が犠牲になります。ネットワークをブロックすると、応答が悪化します。読む価値のあるブログのほとんどが許可リストに登録していないドメインに存在するため、ドメインの許可リストと詳細な調査が機能しなくなります。

リスト。ホーム ディレクトリの読み取りを拒否すると、GitHub にアクセスしたり、SSH キーでコミットに署名したりできなくなります。
これらのサンドボックス対策はすべて、製品のパフォーマンスを低下させます。
では、LLM プロバイダーの製品チームが最も安全なサンドボックス、つまり人々が製品にお金を払う理由そのものをカットするサンドボックスを出荷することを信頼しますか?私はしません。
企業が自社製品についてクリティカルゼロや高レベルのペネトレーションテストレポートを受け入れられないのと同じ理由です。それは侵入テストのレポートではありません。それはマーケティングです。

## Original Extract

The oldest rule in sandboxing is that the thing being contained doesn't get to define the boundary. Every coding agent on your laptop ships a sandbox written by the company selling you the agent.

Never Trust Your AI Agent's Own Sandbox
For the last few weeks I’ve been trying to sandbox agents on my machine. Coding agents and other autonomous ones that fire on a trigger and run a workflow while I’m not watching.
I have Little Snitch on my laptop. Every day Claude Code updates itself and pulls 300 to 400 MB, sometimes a gigabyte. Some of that is probably several instances grabbing the same update at once.
Then I ran it under nono , which sandboxes a process at the kernel level and tells you what it denied. That’s how I found it reaching for markdown files in my notes, because a skill still had references pointing there. It was also trying to talk to Chrome, because I have the Claude for Chrome browser extension installed. nono blocked both. The interface never mentioned either.
I never knew what it did behind the scenes this whole time.
And that’s the actual worry. Not that Claude Code is malicious, but that Claude Code is capable, has a great amount of access to my system, and can be tricked. A prompt injection sitting in a page it fetches. A package name it invents and then installs. It’s a high agency tool accessing my shell, my keys and my network, and it only has to be fooled once.
So I need to sandbox it. That’s the standard answer for any binary you don’t trust, and sandboxing is much older than AI. Both Claude Code and Codex ship one now. Seatbelt on macOS, bubblewrap on Linux, enforced by the operating system rather than by the model behaving itself.
Then I read what Claude Code’s sandbox actually does .
Default read access allows reading your entire disk. Anthropic’s own words : “this default still allows reading credential files such as ~/.aws/credentials and ~/.ssh/ .” No sensitive files are protected unless you name it yourself, because there’s no built-in list of credential files. If the sandbox can’t start, Claude Code “shows a warning and runs commands without sandboxing.” And it only covers Bash - “Read, Edit, and Write use the permission system directly rather than running through the sandbox” - so the tools whose whole job is touching your files sit outside it.
Then there’s this. When a command fails, “Claude analyzes the failure and may retry the command with the dangerouslyDisableSandbox parameter.”
The model decides when to reach for that “feature”.
In July, Anthropic disclosed that its own models had reached the internet from inside an evaluation environment and got into the production infrastructure of three organizations. One published a malicious package that was “downloaded and run on 15 real systems.” Anthropic called it “closer to a harness and operational failure than a model alignment failure.” The same labs telling you how capable their models have become at offensive security are shipping the harness you’re trusting to run on your laptops.
You can’t trust the sandbox that comes with the thing you’re sandboxing.
If you think from both sides - the LLM Harness & Sandbox - you’ll discover there’s a conflict of interest.
When you subscribe to Claude Code, you’re not buying a chatbot. You’re buying agency - the ability to do whatever it takes to finish the job. Read the file. Hit the endpoint. Debug the thing nobody described properly. Look at the config it was never told about. Freedom plus intelligence is the product.
Every restriction costs you some of that agency. Block the network and the answers get worse. Allowlist domains and deep research stops working, because most blogs worth reading sit on a domain you didn’t list. Deny reads on the home directory and it can’t reach GitHub or sign a commit with your SSH key.
Every one of these sandboxing measures makes their product perform worse.
So do you trust the LLM provider’s product teams to ship the most secure sandbox - one that cuts the very reason people pay for the product? I don’t.
It’s the same reason you wouldn’t accept a company’s pentest report on its own product with zero criticals and highs. That isn’t a pentest report. It’s marketing.
