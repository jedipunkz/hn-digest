---
source: "https://www.backplanes.com:443/"
hn_url: "https://news.ycombinator.com/item?id=48545168"
title: "Show HN: Spotlight shows what your Claude Code/Codex are doing"
article_title: "Backplanes Spotlight — See what your agent actually did"
author: "nickv"
captured_at: "2026-06-15T19:25:27Z"
capture_tool: "hn-digest"
hn_id: 48545168
score: 4
comments: 0
posted_at: "2026-06-15T18:24:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Spotlight shows what your Claude Code/Codex are doing

- HN: [48545168](https://news.ycombinator.com/item?id=48545168)
- Source: [www.backplanes.com:443](https://www.backplanes.com:443/)
- Score: 4
- Comments: 0
- Posted: 2026-06-15T18:24:54Z

## Translation

タイトル: HN を表示: スポットライトはクロード コード/コーデックスが何を行っているかを示します
記事のタイトル: バックプレーンのスポットライト — エージェントが実際に行ったことを確認する
説明: クロード コードおよびコーデックス セッションの自動レポート。アクセスされたファイル、実行されたコマンド、到達した外部ツール、スコープのドリフト、およびレビューに値するものを確認します。発売時は無料。
HN テキスト: こんにちは HN!長年潜んでいる者、時々コメントをする者、ここに初めて投稿する者。私は 2 人の共同創設者と数人の同僚と一緒に、皆さんと共有できることに興奮しています (そして少し緊張しています) プロジェクトに取り組んでいます。私たちの多くと同じように、私もここ数か月間、AI コーディング (雰囲気でしょうか?) に悩まされて生きてきました。11 月はこれに関して大きな瞬間だったと思います。しかし、オーケストレーション層を次々と構築した後で気づいたことの 1 つは、「クロード コードは実際に何をしているのか?」ということだということです。おそらくそれは、Claude Code が混乱して root を「rm -rf」したとき、機能フラグの切り替えをステージではなく prod にデプロイしたとき、または Playwright でログインするという破滅のサイクルにはまったときなどが原因かもしれません。しかし、それが私たちにこのツールを構築するきっかけを与えてくれました。私たちはそれをバックプレーンによるスポットライトと呼んでいます。 Spotlight は、Claude Code と Codex のセッションを取得し、セキュリティの問題、高速化できる可能性があること、時間とトークンを浪費している箇所を見つけます。また、あなたがどのようなビルダーであるかを示す楽しい小さな原型も作成します。仕組み: バックプレーン CLI デーモン/TUI をインストールします。これは、Claude セッションと Codex セッションを取得し、それらの PII とシークレットをローカルでスクラブし、ローカルでホストされているモデルを使用して第 2 レベルのスクラブを実行して、セッションの行レベルを顧客キーで暗号化して保存します。現在、AWS Secrets Manager で顧客キーを作成して保存しています。

直接アクセスしないでください。この作業をホストして実行すると、マシンや複数のハーネス間でセッションをつなぎ合わせることができ、チーム レベルのパターンを提供することもできます。詳細は https://backplanes.com/trust をご覧ください。レポートの例は https://backplanes.com/features/session-reports で見ることができます。これを試すには、1 行の CLI インストールで (はい、サインアップが必要です、申し訳ありません)、backplanes.com で無料です。今後数週間以内に、MacOS および Windows のネイティブ アプリとともに Powershell バージョンをリリースする予定です。皆さんのご意見をお聞かせください。みんな、ありがとう！ニック

記事本文:
バックプレーン スポットライト — エージェントが実際に行ったことを確認する 機能 セッション レポート 各エージェント セッションが実際に行ったこと 組織レポート 1 つのレポート — セキュリティ、エンジニアリング、支出のビュー MCP と外部アクセス エージェントが到達するすべての外部ドメイン 対象者 エンジニアとビルダー 自身のエージェント セッションの迅速なレビュー エンジニアリング マネージャー チームがどこに AI キャパシティを費やしているか CFO チームおよびツール別の支出、ROI、およびキャパシティ CISO 外部アクセス、データ出力、およびポリシー 信頼 サインイン 無料でサインアップ エージェントが実際に何をしているのか疑問に思ったことはありませんか?
Spotlight by Backplanes は、クロード コードとコーデックス セッションを読み取り、再帰的に改善するのに役立つセッション レポートを提供します。個人でもチームでも無料。
$カール -fsSL https://www.backplanes.com/spotlight/install.sh | sh コピー
macOS、Linux、および WSL 2 で動作します。
パスワード リセット - 42 分 - クロード コード
タスク: パスワード リセット フローを追加する
推論 3 つのアプローチを試した結果、使い捨てのマジック リンクにたどり着きました
ベスト プラクティス トークンの有効期限は 15 分で、使用できるのは 1 回だけです
4 ファイル 17 コマンド 3 外部サービス
すべてのセッションに光を当てます。
あなたが会議に参加している間、エージェントは 47 分間走り続けました。 Spotlight はすべての動きを監視し、保持する価値があるもの、修正すべきもの、次回の実行で時間を節約できる場所についての迅速なフィードバックを提供します。
この向こう側にいた人々によって建てられました。
一度インストールしてください。セッションごとにレベルアップします。
macOS、Linux、および WSL 2 で動作します。
ブラウザで認証し、チーム アカウントを作成し、終了時にセッションのキャプチャを開始します。
02 働き続けてください。すでに取り組んでいます。
9:42 AM セッション終了 パスワードリセット準備完了
10:15 AM セッションが終了しました auth-refresh READY
11:08 AM セッションが終了しました merge-pg-v14 処理
Spotlight は、セッションが終了すると自動的にセッションをキャプチャします。ローカル編集により PII と資格情報が削除されます

ラップトップから何かが出る前に。あなたは流れの中に留まります。
Anthropic または OpenAI への OAuth はありません。 CLI はセッションの終了後にのみセッションを読み取ります。
auth-refresh-token クリーンリフレッシュ、署名が検証されました。
merge-postgres-v14 移行はクリーンに実行され、スキーマのドリフトはありませんでした。
ux-onboarding-tweak 3 つのコンポーネントが更新され、テストに合格しました。
最初のレポートは、次に完了したセッションの後に表示されます。 Spotlight では、今日どのレポートが必要か、どのレポートをスクロールして通過できるかがわかります。
Spotlight は、個々の開発者とその開発チームにとって無料です。数えられるシートや試用時間はありません。組織全体に展開していて、帰属、ボリューム、または特定の制御が必要な場合は、私たちにご相談ください。一緒に設定します。
macOS、Linux、および WSL 2 で動作します。
エージェントはこれまで以上に多くのことを行っています。そしてそのほとんどは目に見えないものです。私たちの最初の製品である Spotlight は、今日のあなた自身のセッションを表示します。時間が経つにつれて、バックプレーンは、すべてのエージェント (自分、チーム、組織全体) の行動を確認し、フォローし、形成するのに役立ちます。次のセッションから始めてください。
Spotlight は現在、Claude Code および Codex で動作します。
次に輝くべき場所を教えてください。
Antigravity CLI Google OpenCode オープンソース Cursor Cursor AI Claude Cowork Anthropic π Pi オープンソース その他 Tell us Next from Backplanes
スポットライトは 1 番目の移動です。残りの間は近くにいてください。
hello@backplanes.com · 組織のロールアウトについてご相談ください · Slack にご参加ください →

## Original Extract

Automatic reports for Claude Code and Codex sessions. See files touched, commands run, external tools reached, scope drift, and what deserves review. Free at launch.

Hola HN! Long time lurker, sometimes commentor, first time poster here. I’ve been working alongside my two co-founders and a few colleagues on a project I’m excited (and a little nervous) to share with you all! Like many of us, I’ve lived a tortured existence with AI coding (is it vibes?) over the past few months - I think November was a big moment with this. But, one thing I’ve noticed after building orchestration layer after orchestration layer is that the thing I always came back to was “what the hell is Claude Code actually doing?” Perhaps it’s because of the time Claude Code got confused and “rm -rf”’ed root, or the time it deployed a feature flag flip to prod instead of stage, or the time it got stuck in a cycle of doom logging in with playwright, etc etc - but it inspired us to build this tool; we call it Spotlight by Backplanes. Spotlight takes your Claude Code and Codex sessions and finds security issues, things that could be sped up, and where you’re burning your time and tokens. We also create fun little archetypes of what kind of builder you are. The way it works: you install a backplanes CLI daemon/TUI that takes your Claude and Codex sessions, scrubs them of their PII and secrets locally, sends them to us where we do a second level scrub using a locally hosted model, and store your sessions row level encrypted with customer keys. Today we create and store the customer keys in AWS secrets manager, which we can’t access directly. Doing this work hosted lets us stitch sessions across machines and multiple harnesses and even gives you the ability to give team-level patterns. Details are at https://backplanes.com/trust . You can see an example report at https://backplanes.com/features/session-reports , To play with this, it's a one-line CLI install (yes, there's a signup, I'm sorry), and it's free at backplanes.com. In the coming weeks we will be releasing a Powershell version along with native MacOS and Windows apps. Please let us know what you all think. Thanks guys and gals! Nick

Backplanes Spotlight — See what your agent actually did Features Session reports What each agent session actually did Org reports One report — Security, Engineering and Spend views MCP & external access Every external domain your agents reach Audiences Engineers & builders Faster review of your own agent sessions Engineering managers Where the team is spending AI capacity CFOs Spend, ROI and capacity by team and tool CISOs External access, data egress and policy Trust Sign in Sign Up for Free Ever wonder what your agent is actually doing?
Spotlight by Backplanes reads your Claude Code and Codex sessions to give you session reports that help you get recursively better. Free for individuals and teams.
$ curl -fsSL https://www.backplanes.com/spotlight/install.sh | sh Copy
Works with macOS, Linux, and WSL 2.
password-reset - 42 min - Claude Code
Task: Add a password reset flow
Reasoning Tried 3 approaches, landed on single-use magic links
Best practice Tokens expire in 15 minutes and can only be used once
4 files 17 commands 3 external services
Shine a light on every session.
Your agent ran for 47 minutes while you were in a meeting. Spotlight watches every move and gives you quick feedback on what’s worth keeping, what to fix, and where to save time next run.
Built by people who’ve been on the other side of this.
Install once. Level up every session.
Works with macOS, Linux, and WSL 2.
Authenticates in your browser, creates your team account, and starts capturing sessions as they finish.
02 Keep working. We're already on it.
9:42 AM Session ended password-reset READY
10:15 AM Session ended auth-refresh READY
11:08 AM Session ended migrate-pg-v14 Processing
Spotlight captures sessions automatically when they end. Local redaction strips PII and credentials before anything leaves your laptop. You stay in flow.
No OAuth into Anthropic or OpenAI. The CLI only reads sessions after they end.
auth-refresh-token Clean refresh, signature verified.
migrate-postgres-v14 Migration ran clean, no schema drift.
ux-onboarding-tweak Three components updated, tests pass.
Your first report appears after your next completed session. Spotlight tells you which reports need you today and which ones you can scroll past.
Spotlight is free for individual developers and the teams they work on: no seats to count, no trial clock. If you’re rolling it out across a whole org and need attribution, volume, or specific controls, talk to us and we’ll set it up together.
Works with macOS, Linux, and WSL 2.
Your agents are doing more than ever. And most of it is invisible. Our first product, Spotlight, shows you your own sessions today. Over time, Backplanes will help you see, follow, and shape what every agent does: yours, your team's, your whole org's. Start with your next session.
Spotlight works with Claude Code and Codex today.
Tell us where it should shine next.
Antigravity CLI Google OpenCode Open source Cursor Cursor AI Claude Cowork Anthropic π Pi Open source Other Tell us Next from Backplanes
Spotlight is move one. Stay close for the rest.
hello@backplanes.com · Talk to us about org rollouts · Join our Slack →
