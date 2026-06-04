---
source: "https://github.com/shaftoe/pi-coding-agent-action/pull/260"
hn_url: "https://news.ycombinator.com/item?id=48393369"
title: "Beware of infinite loops when using AI"
article_title: "docs(README.md): improve 'if' clause to detect /pi by knocte · Pull Request #260 · shaftoe/pi-coding-agent-action · GitHub"
author: "knocte"
captured_at: "2026-06-04T04:35:13Z"
capture_tool: "hn-digest"
hn_id: 48393369
score: 1
comments: 0
posted_at: "2026-06-04T03:26:37Z"
tags:
  - hacker-news
  - translated
---

# Beware of infinite loops when using AI

- HN: [48393369](https://news.ycombinator.com/item?id=48393369)
- Source: [github.com](https://github.com/shaftoe/pi-coding-agent-action/pull/260)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T03:26:37Z

## Translation

タイトル: AIを使用する場合は無限ループに注意してください
記事のタイトル: docs(README.md): knocte による /pi を検出するための 'if' 句の改善 · Pull Request #260 · shutoe/pi-coding-agent-action · GitHub
説明: https://pi.dev/ コーディング エージェントを GitHub 互換の CI/CD、問題、PR と統合するための GitHub アクション - ドキュメント (README.md): knocte による /pi を検出するための 'if' 句の改善 · Pull Request #260 · shaftoe/pi-coding-agent-action

記事本文:
docs(README.md): /pi を検出するための 'if' 句の改善 by knocte · Pull Request #260 · shutoe/pi-coding-agent-action · GitHub
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
シャフトー
/
pi-コーディング-エージェント-アクション
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
docs(README.md): /pi を検出するために 'if' 句を改善 # 260
オープン knocte は 1 つのコミットを shoe:develop にマージしたいです shoe/pi-coding-agent-action:develop from knocte:wip/startsWithRatherThanContains knocte/pi-coding-agent-action:wip/startsWithRatherThanContains ヘッドブランチ名をクリップボードにコピーします 会話コミット 1 ( 1 ) チェック ファイル変更 オープンdocs(README.md): /pi # 260 を検出するために 'if' 句を改善します。 knocte は 1 つのコミットを shoe:develop にマージしたいです shoe/pi-coding-agent-action:develop from knocte:wip/startsWithRatherThanContains knocte/pi-coding-agent-action:wip/startsWithRatherThanContains ヘッド ブランチ名をクリップボードにコピーします
リンクをコピー
マークダウンのコピー
投稿者
ノクテ
コメントした
2026 年 6 月 4 日
•
編集済み
読み込み中
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
この README.md ファイルの中で「コマンド」/pi が説明されている箇所では、動詞「prefix」が使用されています。そのため、サンプル (つまり、クイック スタート セクション) に従っている場合、github コメントの先頭で /pi が使用されている場合にのみ「ボット」がアクティブになるとユーザーに思わせています。
したがって、これらのサンプルを contains の代わりに startWith を使用するように変更することをお勧めします。さらに、ユーザーが設定した、または将来設定する可能性のある他のコマンド (例: "/pin yadda yadda"、または単に "/ping" など) との潜在的な (まれではありますが) 衝突を防ぐために、文字列の "/pi" の後にスペースを追加する方が安全です: "/pi "。
ところで、この PR のマージに懐疑的なら、私があなたのプロジェクトを試す前にこの変更が行われていたら起こらなかったであろうちょっと面白い話をさせてください。 (注: 今となってはおかしな話ですが、幸いなことに私は安価なモデルと従量課金制のプロバイダーを使用していましたが、この問題が発生し始めてからわずか 1 時間後に気づきました。しかし、次のことを想像してみてください。

これはフロンティアモデル、つまりクレジットカードを受け取るプロバイダーを使用している人に起こり、問題が発生し始めた後、彼女は就寝します。私とは違って、犬の散歩に行っただけなので、わずか1時間後に犬を捕まえました。）
物語は、/pi 経由で Pi に PR を作成するように指示しようとしたところから始まります。この PR は、if github.actor = ... 要素を追加することで pi.yml ワークフローを変更し、安全にするようにしました。私が #255 を提出したのはその時でした、覚えていますか?その後、PAT を作成してワークフローで使用し、/pi 経由で Pi に問題を再度解決するように指示しました。
それから犬の散歩に行くことにしました。そして戻ってきたら、今度は PR を作成できていたのですが、完了した作業を知らせるメッセージが私のユーザー アカウントで書き込まれ (新しい PAT のため)、それに文字列 /pi が含まれていました (ワークフローをトリガーするのに「/pi」コマンドを使用できるのは Github ユーザーのみであることが示されていたため)。そのため、それによって別のワークフローがトリガーされ、それ自体が別の PR と別のメッセージを作成し、さらに別のワークフローが生成されるという具合でした。など！
幸いなことに、私は 1 時間後に無限ループに遭遇しただけで、費やした金額は約 2 ドルだけでした (安いキミに感謝 🙏)。すぐに停止したところ、同じ内容の PR を何度も作成することがなくなりました。したがって、この PR を実施することで、今後このようなことが誰にも起こらないようにする必要があります。
読み込み中にエラーが発生しました。このページをリロードしてください。
docs(README.md): /pi を検出するために 'if' 句を改善しました
…
400eb47
この README.md ファイル内の「コマンド」/pi が存在するすべての場所
説明すると、動詞「prefix」が使用されます。それで、それは
ユーザーは、「ボット」が /pi の場合にのみアクティブ化されると考えています。
github コメントの**先頭**で使用されます。
サンプル（つまり、「クイックスタート」セクション）。
したがって

e、これらのサンプルを使用するように変更することをお勧めします。
「contains」の代わりに「startsWith」を使用します。その上で、防ぐためには、
他の可能性のあるものとの潜在的な（まれではあるが）衝突
ユーザーが設定した可能性のあるコマンド、または
未来 (例: "/pin yadda yadda"、または単に "/ping" など)、それは
文字列の「/pi」の後にスペースを追加する方が安全です:「/pi」。
リンクをコピー
マークダウンのコピー
コードコブ
ボット
コメントした
2026 年 6 月 4 日
✅ 変更されたすべての行とカバー可能な行はテストでカバーされます。
⚠️レポートはBASE（develop@42ab44b）にアップロードしてください。欠落している BASE レポートについて詳しくは、こちらをご覧ください。
@@ 適用範囲の差 @@
# # 開発 #260 +/- ##
=========================================
カバレッジ？ 91.19%
=========================================
ファイル ? 45
ライン？ 3577
枝？ 0
=========================================
ヒット? 3262
ミスですか？ 315
パーシャル？ 0
🚀 ワークフローを強化する新機能:
❄️ テスト分析 : 不安定なテストを検出し、失敗についてレポートし、テスト スイートの問題を見つけます。
📦 JS バンドル分析 : JS マージのバンドル サイズを追跡および制限することで、自分自身を守ります。
読み込み中にエラーが発生しました。このページをリロードしてください。
github-アクション
ボット
このプルリクエストについて言及しました
2026 年 6 月 4 日
毎日の AI ダイジェスト - 2026-06-04
FutureGadget/ai-sota-feed-bot#67
開く
-->
このファイルには、以下に表示されるものとは異なる方法で解釈またはコンパイルされる可能性のある、非表示または双方向の Unicode テキストが含まれています。確認するには、エディタでファイルを開くと、隠された Unicode 文字が表示されます。
双方向 Unicode 文字の詳細については、こちらをご覧ください。
隠れた文字を表示する
無料でサインアップ
GitHub でこの会話に参加するには 。
すでにアカウントをお持ちですか?
コメントするにはサインインしてください
-->
査読者
このプル リクエストを正常にマージすると、これらの問題が解決される可能性があります。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.

フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

GitHub action to integrate https://pi.dev/ coding agent with GitHub-compatible CI/CD, issues and PRs - docs(README.md): improve 'if' clause to detect /pi by knocte · Pull Request #260 · shaftoe/pi-coding-agent-action

docs(README.md): improve 'if' clause to detect /pi by knocte · Pull Request #260 · shaftoe/pi-coding-agent-action · GitHub
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
shaftoe
/
pi-coding-agent-action
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
docs(README.md): improve 'if' clause to detect /pi # 260
Open knocte wants to merge 1 commit into shaftoe:develop shaftoe/pi-coding-agent-action:develop from knocte:wip/startsWithRatherThanContains knocte/pi-coding-agent-action:wip/startsWithRatherThanContains Copy head branch name to clipboard Conversation Commits 1 ( 1 ) Checks Files changed Open docs(README.md): improve 'if' clause to detect /pi # 260 knocte wants to merge 1 commit into shaftoe:develop shaftoe/pi-coding-agent-action:develop from knocte:wip/startsWithRatherThanContains knocte/pi-coding-agent-action:wip/startsWithRatherThanContains Copy head branch name to clipboard
Copy link
Copy Markdown
Contributor
knocte
commented
Jun 4, 2026
•
edited
Loading
Uh oh!
There was an error while loading. Please reload this page .
Everywhere in this README.md file where the "command" /pi is explained, the verb "prefix" is used. So then it is making the user think that the "bot" will only be activated when /pi is used at the beginning of a github comment if one follows the samples (i.e. in the Quick Start section).
Therefore, it is better to change these samples to use startsWith instead of contains . On top of that, to prevent the potential (although rare) collision with other possible commands that the user might have set up, or may set up in the future (e.g. "/pin yadda yadda", or simply "/ping", etc), it's safer to add a space after "/pi" in the string: "/pi ".
BTW, if you're skeptical of merging this PR, then let me tell you a little funny story that wouldn't have happened if this change had been done before I tried your project. (Note: it's funny now, and fortunately I was using a cheap model and a pay-as-you-go provider, and I noticed only 1h after it started happening; but imagine if this happens to someone that is using a frontier model, a provider that takes your credit card, and she goes to bed after the problem starts happening; unlike me that I just went to walk the dog and that's why I caught it only 1h after.)
The story begins when I tried to instruct Pi via /pi to create a PR that changed my pi.yml workflow to secure it by adding an if github.actor = ... element to it. That's when I filed #255 , remember? After that, I then created a PAT and used it in my workflow, and then I instructed Pi via /pi to try to fix the issue again.
Then I decided to go walk my dog. And then when I came back, I found that, this time, it had been able to create PR, but the message to inform me about the work that had been done was written with my user account (because of the new PAT) and it contained the string /pi (because it was telling me that now only my github user can use the "/pi" command to trigger the workflow), so then that triggered another workflow, which itself created another PR and another message, which then spawned another workflow, and so on and so forth!
I fortunately only caught the infinite loop after 1h, and only about 2USD were spent (thanks cheap Kimi 🙏). I stopped it right away and it stopped creating PRs for the same thing again and again. So with this PR in place, I HTH to avoid this happening to anyone in the future.
There was an error while loading. Please reload this page .
docs(README.md): improve 'if' clause to detect /pi
…
400eb47
Everywhere in this README.md file where the "command" /pi is
explained, the verb "prefix" is used. So then it is making the
user think that the "bot" will only be activated when /pi is
used at the **beginning** of a github comment if one follows the
samples (i.e. in the Quick Start section).
Therefore, it is better to change these samples to use
`startsWith` instead of `contains`. On top of that, to prevent
the potential (although rare) collision with other possible
commands that the user might have set up, or may set up in the
future (e.g. "/pin yadda yadda", or simply "/ping", etc), it's
safer to add a space after "/pi" in the string: "/pi ".
Copy link
Copy Markdown
codecov
Bot
commented
Jun 4, 2026
✅ All modified and coverable lines are covered by tests.
⚠️ Please upload report for BASE ( develop@42ab44b ). Learn more about missing BASE report.
@@ Coverage Diff @@
# # develop #260 +/- ##
==========================================
Coverage ? 91.19%
==========================================
Files ? 45
Lines ? 3577
Branches ? 0
==========================================
Hits ? 3262
Misses ? 315
Partials ? 0
🚀 New features to boost your workflow:
❄️ Test Analytics : Detect flaky tests, report on failures, and find test suite problems.
📦 JS Bundle Analysis : Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
There was an error while loading. Please reload this page .
github-actions
Bot
mentioned this pull request
Jun 4, 2026
Daily AI Digest - 2026-06-04
FutureGadget/ai-sota-feed-bot#67
Open
-->
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters.
Learn more about bidirectional Unicode characters
Show hidden characters
Sign up for free
to join this conversation on GitHub .
Already have an account?
Sign in to comment
-->
Reviewers
Successfully merging this pull request may close these issues.
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
