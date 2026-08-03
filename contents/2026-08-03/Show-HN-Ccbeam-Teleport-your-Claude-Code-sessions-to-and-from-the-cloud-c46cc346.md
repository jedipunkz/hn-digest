---
source: "https://github.com/sahilmahendrakar/ccbeam"
hn_url: "https://news.ycombinator.com/item?id=49154573"
title: "Show HN: Ccbeam – Teleport your Claude Code sessions to and from the cloud"
article_title: "GitHub - sahilmahendrakar/ccbeam: Teleport a Claude Code session between your devices. Same conversation, same context, across different machines. · GitHub"
author: "smahendrakar"
captured_at: "2026-08-03T12:50:43Z"
capture_tool: "hn-digest"
hn_id: 49154573
score: 2
comments: 0
posted_at: "2026-08-03T11:55:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ccbeam – Teleport your Claude Code sessions to and from the cloud

- HN: [49154573](https://news.ycombinator.com/item?id=49154573)
- Source: [github.com](https://github.com/sahilmahendrakar/ccbeam)
- Score: 2
- Comments: 0
- Posted: 2026-08-03T11:55:44Z

## Translation

タイトル: HN を表示: Ccbeam – クロード コード セッションをクラウドとの間でテレポートします
記事のタイトル: GitHub - sahilmahendrakar/ccbeam: デバイス間でクロード コード セッションをテレポートします。異なるマシン間でも同じ会話、同じコンテキスト。 · GitHub
説明: デバイス間でクロード コード セッションをテレポートします。異なるマシン間でも同じ会話、同じコンテキスト。 - サヒルマヘンドラカール/ccbeam
HN テキスト: コーディング エージェントの実行中にラップトップを開いたままにするのにうんざりしたので、ccbeam を構築しました。 ccbeam を使用すると、ラップトップ、クラウド、リモート デバイスとの間でクロード コード セッションをテレポートできます。進行中の作業も転送されます。ラップトップからクロード コード セッションを開始し、それをクラウドまたはリモート デバイスに送信し、再び取得したいときにはいつでもラップトップにビームして戻すことができます。既存の SSH 構成を使用し、資格情報を転送することはなく、クラウド サンドボックス プロバイダーにとっては BYOK です。このツールは、クロード コード セッションは実際には単なるトランスクリプト履歴であり、ディレクトリ内のダーティ ファイルであるという認識から生まれました。両方を別の環境にコピーできる場合は、基本的に同じクロード コード セッションを実行できます。これにより、デバイスに依存しないクロード コードが可能になります。これは完全にオープンソースであり、他の人にとっても役立つ場合に備えてここで共有します。

記事本文:
GitHub - sahilmahendrakar/ccbeam: デバイス間でクロード コード セッションをテレポートします。異なるマシン間でも同じ会話、同じコンテキスト。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
サヒルマヘンドラカール
/
CCビーム
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット bin bin docs docs plugin plugin scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md TESTING.md TESTING.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
デバイス間でクロード コード セッションをテレポートします。同じ会話、同じ
さまざまなマシンにわたるコンテキスト。
/ccbeam:up を使用して、別の環境 (リモート デスクトップ、クラウドなど) にテレポートします。
会話はコンテキストをそのままにして向こうで再開されます。同じセッション ID、
何も要約されておらず、コミットされていない変更も一緒に送信されます。それ以来
CPU、 CLAUDE.md 、フック、MCP など、すべてがそのマシン上で実行されます。
サーバー。 /ccbeam:home では、変更した内容がそのまま戻ってきます。
~/.ssh/config 内の任意のマシンが宛先になります。 SSHできれば
すでにリストにあります。 Claude Code 以外にインストールするものはありません。
実行し続けるデーモンもありません。
または、独自のキー上のクラウド サンドボックス。 E2B APIの導入
キーとクラウドは同じリストに追加されます。 ccbeam はインフラストラクチャを実行せず、何も保持しません
キーとプロキシは何もありません。サインアップするホスト型 ccbeam はありません。
アカウントもデーモンも設定ファイルもありません。すでに所有している SSH を使用します。
npm i -g ccbeam
ccbeam install-shell # オプション: `claude` を入力し続けます
ccbeam は claude と同じオプションを選択し、端末を直接次のユーザーに渡します。
本物のクロード・コード。後継者ではなく監督者です。
ccbeam # `クロード` のように
ccbeam --model opus # 任意のクロード フラグが機能します
install-shell は、 .bashrc 、 .zshrc 、または
config.fish なので、入力したコマンドは claude のままになります。
クロード () { コマンド ccbeam " $@ " ; }
同じトリック nvm 、 pye

nv と direnv を使用します。コマンド クロードはまだ実行されます クロード
直接コードを記述すると、ccbeam uninstall-shell がファイルを元の状態に正確に復元します。
(どちらの場合でも、最初の編集の前にバックアップされます)。
/ccbeam:home はどこにいても、ホップ数に関係なく機能します。
/ccbeam:up home も、指が伸びるのであれば同じことです。
local は他のデバイスと同様のデバイスであるため、/ccbeam:up によって会話も移動します
すでに使用しているマシン上のフォルダー間では、Claude Code では実行できません。
それ以外の場合はセッションの途中で実行します。
ビームに
⌂ ローカル あなたはここにいます
GPUボックス 2時間前
昨日のマックミニ
☁ 雲が一時停止しました
↑↓ 選択 · ⏎ 確認 · フィルターに入力 · Esc キャンセル
リストは ~/.ssh/config から読み取られるため、保持するレジストリはありません
同期し、それを「追加」するためにマシン上で実行するものは何もありません。クラウドは常にそこにあります、設定されています
上がるかどうか。
gpu-box — フォルダー
~/dev/api main ·3 ダーティ
~/src/トレーナー cuda-12
~/dev/スクラッチメイン
フォルダーは、そのデバイス自体のクロード コード履歴から取得されます。
実際に動作し、最新のものから順に、ブランチとダーティファイルの数を使用して、次のことができます。
あなたが何に向かって歩いているかを見てください。最後に使用したものが事前に選択されているため、
戻る場所は /ccbeam:up → ⏎ → ⏎ です。
会話。トランスクリプトが宛先にコピーされ、再開されます
同じセッション ID の下にあるため、コンテキストは完全にそのままです。
要約よりも。
コミットされていない作業 (フォルダーが git リポジトリの場合) - 変更され、
両方とも追跡されていないファイル。そうすると戻ってきます。
他には何もありません。設定、MCP サーバー、フック、および資格情報は、
使用しているデバイス。それが重要です。そのマシンの環境を取得するのです。
コミットされていない変更は世界で最も失われやすいものなので、すべてのキャリーが
その仮定を検証するか、または拒否します。
宛先は同じコミット上にある必要があります。に適用されたパッチ
違う

基本は、仕事が静かに消える方法です。
あなたが去った木でない限り、目的地はきれいでなければなりません。
家に帰ったら、残したフォルダーは変更されていないはずです。何かが編集した場合
あなたが不在の間、ccbeam は受信した作業を拒否し、保存します。
~/.ccbeam/incoming-<timestamp>/ を上書きするのではなく。
/ccbeam:up クラウドは、他のデバイスと同様に動作するサンドボックスです。最初の
今度は、E2B キーを要求し、自身をセットアップします。その後
画面上では 2 行、約 1 秒です。
› /ccbeam:アップクラウド
· クラウドボックスを起動する
· 3 つの変更されたファイルを保持
☁ cloud:~/dev/api · /ccbeam:home を返す
それはあなたのアカウントで実行されます。 ccbeam はインフラストラクチャを運用せず、鍵も保持しません
何もプロキシしません。コードはマシンからサンドボックスに送られます。
その間には何もありません。ホスト型 ccbeam は存在せず、今後も存在する予定はありません。
知っておく価値のあるいくつかの決定事項:
1 つのボックスで持続します。一時停止するとファイルシステムが保存され、
メモリは無期限に保存され、再開には約 1 秒かかります。なので箱はそのままです
サインインすると、そこにインストールしたものはすべて保存されます (たまたま使用したマシン)
毎回新しいコンテナを購入するのではなく、秒単位でレンタルします。
あなたが去った瞬間にそれは一時停止し、そう言います。それも逃げることができない
あなた: すべてのサンドボックスは有限のタイムアウトと autoPause を使用して作成されるため、
ccbeam が完全に強制終了された場合、ボックスは課金の代わりに自動的にスリープ状態になります。
気づくまでは。
新しいボックスは、GitHub からではなく、ラップトップからシードされます。 ccbeam は
履歴の git バンドルなので、ボックスは正確なコミットに到達します。
コミットされていない仕事は反対されました。プッシュしたことのないリポジトリでも動作します
およびローカルにのみ存在するブランチ、およびボックスは資格情報を必要としません。
git ホスト。
E2B SDK は依存関係ではありません。それは ~/.ccbeam/deps にフェッチされます
固定バージョン

クラウドボックスを初めてセットアップするとき。 ccbeam のインストール
使用していない機能については何も取得しません。
そこで会話を終了し、再び会話を再開する
ボックスは存続するため、そこで行った会話はそこに残ります。を閉じます
ラップトップに戻ってください:
/ccbeam:up クラウド再開ボックス上での会話をピックアップ
ccbeamクラウドセッションがそれらをリストします
ccbeam クラウド履歴書を選択して、ここからセッションを開始します
ccbeam cloud rm < id > 1 つ削除します
resume は意図的に別の動詞です。 /ccbeam:up クラウドとはこれを意味します
会話はセッション ID を維持したままそこに移動します。これがすべての理由です。
コンテキストは存続します。 /ccbeam:up クラウド履歴書は別の会話を採用します
そして、あなたがいる場所を元の場所にそのまま残します。それらが動詞を共有していれば、
/ccbeam:up は黙ってその場所を放棄する可能性があります。彼らはそうしないので、それはできません。
セッションピッカーには、最初に発言した内容、フォルダー、最後に発言した日時が表示されます。
タッチすると、まだ実行中の会話がマークされます。 Ctrl-D で削除します
強調表示されたものは、それが動作していたフォルダーを保持します — ccbeam はあなたのフォルダーを削除しません
ファイル。 30 日間変更されなかった会話は、ボックスが起動すると削除されます。
そういう時はそう言うのです。
切断後もボックス内での作業は実際に続行され、ボックスは一時停止されます。
飛行中にターンをフリーズし、そのまま再開します。 ccbeam は再接続しません
ただし、古い端末を使用する必要はありません。トランスクリプトはセッションですので、
バックアップを選択するのは、あそこのクロード --再開です。これはすべてのビームです
すでにそうしています。
ccbeam Cloud destroy はボックスを削除します。 ccbeam クラウド修復は新しいものを構築します
1つ。
ビームするデバイスには ssh アクセス、クロード コード (サインイン)、ノード 18+ が必要です
そして git 。 ccbeam ドクター <デバイス> は何が欠けているかを教えてくれます。クラウドボックス
独自のインストールが行われますが、実行するにはマシン上に Node 20.18.1+ (21.x ではありません) が必要です
あなたはビームです

om 、E2B SDK がそこで実行されるためです。 ssh デバイスは 18 では問題ありません。
$ ccbeam デバイス
● ローカル あなたはここにいます
● GPU ボックス 2 時間前
○昨日のMac-mini
⏸ クラウドが一時停止しました
各デバイスは、 clude auth login を使用して Claude Code 自体にサインインします。それは
意図的 — 以下を参照してください。
認証情報を常に転送します。 ccbeam はモデル呼び出しをプロキシすることはありません
ラップトップに移動するか、トークンを別のデバイスに送信します。 Anthropic の規約では禁止されています
他の製品またはサービスでクロード サブスクリプション OAuth トークンを使用する。それぞれ
デバイス自体を認証することが唯一の正しい設計であり、これに限定されるものではありません。
周りのエンジニア。クラウド ボックスも例外ではありません。クラウド ボックスは独自の方法でサインインします。
一時停止したサンドボックスがそのサインインを維持するため、そのサインインは存続します。
ファイルシステム。
ホストされているものを実行します。リレーでもプロキシでもデフォルトの API キーでもありません。
メンテナーのアカウントで公開されたサンドボックス イメージですらありません。何も入っていない
ランタイム パスは、制御できないインフラストラクチャに依存します。
SSH デバイスをサンドボックス化します。クロード コードは、あなたと同じように遠くのマシン上で実行されます。
アクセス — 自分で SSH 接続して起動する場合と同じです。雲
ボックスは分離されているため、それを使用する理由になりますが、それは E2B の特性です
ccbeam ではなく。
チームのアクセス制御。それはあなたのデバイスと鍵です。
ccbeam は、バンドルされた小さなプラグインを使用して実際のクロードを起動し、待機します。
/ccbeam:up は、リクエスト ファイルを書き込むツールを呼び出します。スラッシュコマンドでは実行できません
端末を乗っ取るので、試行しません。
プラグインの Stop フックは、ターン境界でセッションを終了します。
トランスクリプトは完全に書かれています。
ずっと端末を所有していたスーパーバイザーが記録を発送します。
そして目的地を変更すると、そこから本物のクロードが出発します。
端末に接続されています。
通信するプロトコルも、向こう側にデーモンもありません。 SSH接続
シャリ

ng は、繰り返しの呼び出しをそれぞれ最大 10 ミリ秒に保ちます。
docs/design.md には、スーパーバイザーが存在する必要がある理由が説明されています。
およびデバイスの種類を追加する方法 — Fly Machines、Modal、Daytona、およびプレーン VM は次のとおりです。
すべて 50 行の SshDevice とライフサイクルです。これは、
貢献。
npm test # 単体テスト — ネットワークなし、アカウントなし
npm run test:e2e # エンドツーエンド、CCBEAM_HOST に到達可能なホストが必要
ユニット スイートは、偽のデバイスを介して実際のシーディングおよびキャリー ロジックを駆動します。
これはローカルでシェルアウトするため、バンドル→クローン→パッチはホストなしで実行されます。
e2e スイートのセットアップは docs/design.md にあります。
TESTING.md は、マシンがチェックできないパスに関するマニュアルの概要です。
それ自体で。
Anthropic との提携、承認、後援はありません。クロードとクロード
コードは、Anthropic、PBC の商標です。 Apache Beam プロジェクトや Gravitational とは無関係です
テレポート。
デバイス間でクロード コード セッションをテレポートします。異なるマシン間でも同じ会話、同じコンテキスト。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Teleport a Claude Code session between your devices. Same conversation, same context, across different machines. - sahilmahendrakar/ccbeam

I got tired of leaving my laptop open while running coding agents, so I built ccbeam. With ccbeam, you can teleport your claude code session to and from your laptop, the cloud, and your remote devices. Your in progress work is transported too. You can start a claude code session from your laptop, beam it up to the cloud or a remote device, and beam it back to your laptop whenever you want to pick it up again. It uses your existing SSH config, never forwards credentials, and is BYOK for a cloud sandbox provider. This tool came out of the realization that a claude code session is really just the transcript history, and the dirty files in your directory. If you can copy both over to a different environment, you can essentially run the same claude code session. This allows for a device-agnostic claude code. It's fully open source, sharing it here in case its useful to others too!

GitHub - sahilmahendrakar/ccbeam: Teleport a Claude Code session between your devices. Same conversation, same context, across different machines. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
sahilmahendrakar
/
ccbeam
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits bin bin docs docs plugin plugin scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md TESTING.md TESTING.md package.json package.json View all files Repository files navigation
Teleport a Claude Code session between your devices. Same conversation, same
context, across different machines.
Teleport to a different environment (remote desktop, cloud, etc.) with /ccbeam:up .
The conversation picks up over there with its context intact — same session id,
nothing summarised — and your uncommitted changes travel with it. From then on
everything runs on that machine: its CPU, its CLAUDE.md , its hooks, its MCP
servers. /ccbeam:home brings you back, with whatever you changed.
Any machine in your ~/.ssh/config is a destination. If you can ssh
there, it's already in the list. Nothing to install on it beyond Claude Code,
and no daemon to leave running.
Or a cloud sandbox, on your own key. Bring an E2B API
key and cloud joins the same list. ccbeam runs no infrastructure, holds no
keys and proxies nothing — there is no hosted ccbeam to sign up for.
No account, no daemon, no config file. It uses the SSH you already have.
npm i -g ccbeam
ccbeam install-shell # optional: keep typing `claude`
ccbeam takes the same options as claude and hands the terminal straight to
the real Claude Code. It's a supervisor, not a replacement.
ccbeam # like `claude`
ccbeam --model opus # any claude flag works
install-shell adds one delimited block to your .bashrc , .zshrc or
config.fish so the command you type stays claude :
claude () { command ccbeam " $@ " ; }
Same trick nvm , pyenv and direnv use. command claude still runs Claude
Code directly, and ccbeam uninstall-shell restores the file exactly as it was
(it's backed up before the first edit either way).
/ccbeam:home works from wherever you are, however many hops in — and
/ccbeam:up home is the same thing if that's what your fingers reach for.
local is a device like any other, so /ccbeam:up also moves a conversation
between folders on the machine you're already on, which Claude Code can't
otherwise do mid-session.
beam to
⌂ local you are here
gpu-box 2h ago
mac-mini yesterday
☁ cloud paused
↑↓ select · ⏎ confirm · type to filter · esc cancel
The list is read from your ~/.ssh/config , so there's no registry to keep in
sync and nothing to run on a machine to "add" it. cloud is always there, set
up or not.
gpu-box — folder
~/dev/api main ·3 dirty
~/src/trainer cuda-12
~/dev/scratch main
Folders come from that device's own Claude Code history: every directory you've
actually worked in, newest first, with branch and dirty-file count so you can
see what you're walking into. The one you used last is pre-selected, so going
back somewhere is /ccbeam:up → ⏎ → ⏎.
The conversation. The transcript is copied to the destination and resumed
there under the same session id, so the context is genuinely intact rather
than summarised.
Your uncommitted work , if the folder is a git repo — modified and
untracked files both. It comes back when you do.
Nothing else. Settings, MCP servers, hooks and credentials belong to the
device you're on. That's the point: you get that machine's environment.
Uncommitted changes are the easiest thing in the world to lose, so every carry
either verifies its assumptions or refuses:
The destination must be on the same commit . A patch applied to a
different base is how work silently disappears.
The destination must be clean , unless it's the tree you left.
Coming home, the folder you left must be unchanged . If something edited it
while you were away, ccbeam refuses and saves the incoming work to
~/.ccbeam/incoming-<timestamp>/ rather than overwriting anything.
/ccbeam:up cloud is a sandbox that behaves like any other device. The first
time, it asks for an E2B key and sets itself up; after that
it's two lines on screen and about a second.
› /ccbeam:up cloud
· waking the cloud box
· carried 3 changed file(s)
☁ cloud:~/dev/api · /ccbeam:home to return
It runs on your account. ccbeam operates no infrastructure, holds no keys
and proxies nothing — your code goes from your machine to your sandbox, with
nothing in between. There is no hosted ccbeam and there is not going to be one.
A few decisions worth knowing about:
It's one box, and it persists. Pausing preserves the filesystem and
memory indefinitely, and resuming takes about a second. So the box stays
signed in and keeps whatever you installed on it — a machine you happen to
rent by the second, not a fresh container each time.
It pauses the moment you leave , and says so. It also can't run away from
you: every sandbox is created with a finite timeout and autoPause , so even
if ccbeam is killed outright the box puts itself to sleep instead of billing
you until you notice.
A fresh box is seeded from your laptop, not from GitHub. ccbeam ships a
git bundle of your history so the box lands on the exact commit your
uncommitted work was written against. It works on repos you've never pushed
and branches that only exist locally, and the box never needs a credential for
your git host.
The E2B SDK is not a dependency. It's fetched into ~/.ccbeam/deps at a
pinned version the first time you set the cloud box up. Installing ccbeam
pulls nothing for a feature you haven't used.
Leaving a conversation there, and picking it up again
Because the box persists, a conversation you took there stays there. Close the
laptop and come back to it:
/ccbeam:up cloud resume pick up a conversation living on the box
ccbeam cloud sessions list them
ccbeam cloud resume pick one up, starting a session here
ccbeam cloud rm < id > delete one
resume is a separate verb on purpose . /ccbeam:up cloud means this
conversation moves there, keeping its session id — which is the whole reason its
context survives. /ccbeam:up cloud resume adopts a different conversation
and leaves the one you're in exactly where it was. If those shared a verb,
/ccbeam:up could silently abandon your place; they don't, so it can't.
The session picker shows what you first said, the folder, when it was last
touched, and marks any conversation still running. Ctrl-D deletes the
highlighted one, keeping the folder it worked in — ccbeam doesn't delete your
files. Conversations untouched for 30 days are pruned when the box wakes, and it
says so when that happens.
Work genuinely continues in the box after you disconnect, and a paused box
freezes a turn mid-flight and resumes it intact. ccbeam doesn't reattach to the
old terminal, though — it doesn't need to. The transcript is the session, so
picking one back up is claude --resume over there, which is what every beam
already does.
ccbeam cloud destroy gets rid of the box. ccbeam cloud repair builds a new
one.
A device you beam to needs ssh access, Claude Code (signed in), node 18+
and git . ccbeam doctor <device> tells you what's missing. The cloud box
installs its own — but driving it needs Node 20.18.1+ (not 21.x) on the machine
you beam from , because the E2B SDK runs there. ssh devices are fine on 18.
$ ccbeam devices
● local you are here
● gpu-box 2h ago
○ mac-mini yesterday
⏸ cloud paused
Each device signs in to Claude Code itself, with claude auth login . That's
deliberate — see below.
Forward credentials, ever. ccbeam will never proxy model calls through
your laptop or ship your token to another device. Anthropic's terms prohibit
using Claude subscription OAuth tokens in other products or services; each
device authenticating itself is the only correct design, not a limitation to
engineer around. The cloud box is no exception: it signs itself in, in its own
terminal, and that sign-in survives because a paused sandbox keeps its
filesystem.
Run a hosted anything. Not a relay, not a proxy, not a default API key,
not even a sandbox image published under a maintainer's account. Nothing in
the runtime path depends on infrastructure you don't control.
Sandbox your ssh devices. Claude Code runs on the far machine as you, with
your access — the same as if you'd ssh'd in and started it yourself. The cloud
box is isolated, which is a reason to use it, but that's a property of E2B
rather than of ccbeam.
Team access control. It's your devices and your keys.
ccbeam launches the real claude with a small bundled plugin, and waits.
/ccbeam:up calls a tool that writes a request file. A slash command can't
take over the terminal, so it doesn't try.
The plugin's Stop hook ends the session at the turn boundary — after the
transcript is completely written.
The supervisor, which has owned the terminal all along, ships the transcript
and your changes to the destination and launches the real claude there,
attached to your terminal.
There's no protocol to speak and no daemon on the far side; ssh connection
sharing keeps the repeated calls at ~10ms each.
docs/design.md covers why the supervisor has to exist,
and how to add a device kind — Fly Machines, Modal, Daytona and plain VMs are
all a fifty-line SshDevice plus a lifecycle, which is the intended shape of a
contribution.
npm test # unit tests — no network, no accounts
npm run test:e2e # end-to-end, needs a reachable host in CCBEAM_HOST
The unit suite drives the real seeding and carrying logic through a fake device
that shells out locally, so bundle → clone → patch is exercised without a host.
Setup for the e2e suite is in docs/design.md ;
TESTING.md is the manual brief for the paths a machine can't check
by itself.
Not affiliated with, endorsed by, or sponsored by Anthropic. Claude and Claude
Code are trademarks of Anthropic, PBC. Not affiliated with the Apache Beam project or Gravitational's
Teleport.
Teleport a Claude Code session between your devices. Same conversation, same context, across different machines.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
