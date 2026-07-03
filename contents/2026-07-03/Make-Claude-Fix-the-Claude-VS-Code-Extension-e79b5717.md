---
source: "https://github.com/Pomax/Claude-VS-Code-Extension-patches"
hn_url: "https://news.ycombinator.com/item?id=48778373"
title: "Make Claude Fix the Claude VS Code Extension"
article_title: "GitHub - Pomax/Claude-VS-Code-Extension-patches: Fix the fucking idiotic Claude VS Code extension so it's less infuriatingly stupid · GitHub"
author: "TheRealPomax"
captured_at: "2026-07-03T19:02:24Z"
capture_tool: "hn-digest"
hn_id: 48778373
score: 1
comments: 1
posted_at: "2026-07-03T18:39:30Z"
tags:
  - hacker-news
  - translated
---

# Make Claude Fix the Claude VS Code Extension

- HN: [48778373](https://news.ycombinator.com/item?id=48778373)
- Source: [github.com](https://github.com/Pomax/Claude-VS-Code-Extension-patches)
- Score: 1
- Comments: 1
- Posted: 2026-07-03T18:39:30Z

## Translation

タイトル: クロードにクロード VS コード拡張を修正させる
記事のタイトル: GitHub - Pomax/Claude-VS-Code-Extension-patches: クソバカな Claude VS Code 拡張機能を修正して、腹立たしいほどバカにならないように · GitHub
説明: クソバカな Claude VS Code 拡張機能を修正して、腹立たしいほどバカにならないようにします - Pomax/Claude-VS-Code-Extension-patches

記事本文:
GitHub - Pomax/Claude-VS-Code-Extension-patches: クソバカな Claude VS Code 拡張機能を修正して、腹立たしいほどバカにならないように · GitHub
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
ポマックス
/
クロード VS コード拡張パッチ
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
Pomax/Claude-VS-Code-Extension-patches
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット README.md README.md patch-instructions.md patch-instructions.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Anthropic は、人々が報告した VS Code 拡張機能のバグを実際に修正することに煩わされません。クロードにバグを修正してもらうこともできますが、代わりにクロードがバグを古いものとしてクローズできるようにして、ユーザーに注意を払う必要がなくなります。ユーザーを必要としているのは誰ですか?この場合の反論は、「Anthropic を必要とする人がいるでしょうか? 拡張機能は単なる Web アプリであり、Claude はそれらをうまく修正できます。」ということだと思います。そこで私は Claude に、Anthropic が手を出せない部分を修正するように言いました。そして今、このリポジトリができました。
おそらく次の 2 つの理由から、Claude 拡張機能の自動更新をオフにする必要があると思われます。
拡張機能が更新されるたびに修正を再適用する必要があります。
これは高頻度取引ではありませんが、頻繁に更新されます。CVE がない限り、月に 1 回以上更新しても文字通り意味がありません。どの Anthropic も機能しないため、1 時間ごとにプッシュされる無意味な変更だけが得られます。
次に、Claude を「自動モード」に設定し、patch-instructions.md を読んでその指示どおりに実行するように指示します。パッチの指示は、Claude が変更を加える許可があるかどうかを確認するように表現されており、ユーザーはリテラル文字列 yes で応答する必要があります。すべて小文字、引用符なし、「はいどうぞ」なし、「どうぞ」なし。正確な文字列「はい」で答えないと、クロードは続行を拒否します。
注: クロードが引き続きいくつかのタスクを実行する許可を求める可能性があります。 CLI には「yolo」モードがあるのに、VS Code 拡張機能にはなぜ「yolo」モードがないのかわかりませんが、そうではないので、

実行する必要があることはすべて実行するように既に明示的に指示したにもかかわらず、おそらく実行を拒否する操作が少なくとも 1 つあります。 「AI」ってすごいじゃないですか？
1.「考えながら気まぐれな用語」を削除しました。
私たちがやっていることを表す言葉は「働く」なので、それが唯一の言葉です。
アニメーション化されたスピナーは削除されました。誰もそんな注目を集める必要はありません。
テキストの遅延とアニメーションも削除されました。繰り返しますが、私たちは仕事を終わらせようとしています。
2. コードブロックに適切な構文の強調表示を追加しました。
Claude は完全に有効なマークダウンを形成しており、拡張機能の作成者は、それが文字通り Monaco に組み込まれていると考えていても、それらを適切に構文ハイライトすることに煩わされませんでした。絶対に許せない。
「Xを追加してもいいですか？」という意味だと思います。そして、絶対にそうしてください。
これらのパッチを使用している場合は、Claude 拡張機能を使用していることになるため、Claude を利用できるようになります。修正したい内容を詳細に記載した問題を提出し、patch-instructions.md ファイルですでに多くの内容が修正されているのと同じ方法でその作業を行うよう Claude に依頼し、変更内容を含めるために指示ファイルを更新するように依頼します。次に、問題を解決する PR としてそれを投稿するだけです。
Anthropic が問題を直すのに手間がかからないとしても、少なくとも私たち自身で間違っているところはすべて一緒に直すことができます。
なぜクロードにこんなことをさせる必要があるのでしょうか？
あなたはしない。説明を読んで自分で実行すれば、トークンはかかりません。または、他の作業をしている間にそれを実行するようにクロードに指示することもできます。ツール自体で更新を実行していることを気にする必要はないからです。
クソバカな Claude VS Code 拡張機能を修正して、腹立たしいほどバカにならないようにする
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
途中でエラーが発生しました

オーディング。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fix the fucking idiotic Claude VS Code extension so it's less infuriatingly stupid - Pomax/Claude-VS-Code-Extension-patches

GitHub - Pomax/Claude-VS-Code-Extension-patches: Fix the fucking idiotic Claude VS Code extension so it's less infuriatingly stupid · GitHub
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
Pomax
/
Claude-VS-Code-Extension-patches
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Pomax/Claude-VS-Code-Extension-patches
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits README.md README.md patch-instructions.md patch-instructions.md View all files Repository files navigation
Anthropic can't be bothered to actually fix any VS Code extension bugs that people file. They could just let Claude fix them, but instead they let Claude close bugs as stale so they don't have to pay any attention to their users. Who needs users? I guess the retort to that in this case is "who needs Anthropic? Extensions are just web apps, Claude can fix those just fine". So I told Claude to fix the things Anthropic can't be bothered to, and now there's this repo.
You probably want to turn off auto-update for the Claude extension, for two reasons:
you'll need to reapply the fixes every time the extension updates, and
it updates a lot even though this is not high frequency trading, and there is literally no point in updating more than once a month unless there's a CVE. Which Anthropic won't even act on, so all you're getting are meaningless changes pushed what seems like every hour.
Then: set Claude to "auto mode" and tell it to read through patch-instructions.md and do what it says. The patch instructions are phrased in such as a way that Claude will ask you to confirm that it has permission to make changes, and you will have to reply with the literal string yes . All lower case, no quotes, no "yes please", no "go for it", if you do not answer with the exact string yes Claude will refuse to proceed.
Note: it's possible that Claude will still ask for permission to do some tasks. I have no idea why the CLI has "yolo" mode but the VS Code extension doesn't, but it doesn't, so there's at least one operation that it'll probably refuse to do, despite the fact that you already explicitly told it to do whatever it had to do. Isn't "AI" great?
1. I removed the "whimsy terms while thinking".
The word for what we're doing is "working", and so that's the only word it uses.
The animated spinner has been removed. No one needs that attention hog.
The text delay and animation's been removed too. Again, we're trying to get work done.
2. I added proper syntax highlighting for code blocks.
Claude forms perfectly valid markdown, and the extension authors couldn't be bothered to properly syntax highlight them, even thought that's literally built into Monaco. Absolutely inexcusable.
I think you mean "can I add X?". And absolutely, please do.
If you're using these patches, you are using the Claude extension so you have Claude available: just file an issue detailing what you want to fix, and then ask Claude to do that work in the same way that the patch-instructions.md files already fixes a bunch of stuff, and ask it to update the instructions file to include your changes. Then simply contribute that as a PR that fixes your issue.
If Anthropic can't be bothered to fix stuff, at least we can fix everything that's wrong ourselves, together .
Why do I need to let Claude do this?
You don't. Go read the instructions and perform them yourself, then it'll cost you zero tokens. Or you can tell Claude to do it while you work on other things, because you shouldn't need to care about a tool running an update on itself.
Fix the fucking idiotic Claude VS Code extension so it's less infuriatingly stupid
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
