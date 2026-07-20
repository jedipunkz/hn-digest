---
source: "https://github.com/goawaygeek/ground-control"
hn_url: "https://news.ycombinator.com/item?id=48986313"
title: "Show HN: A chess game you play with Claude as your wing-bot in the CLI"
article_title: "GitHub - goawaygeek/ground-control · GitHub"
author: "goawaygeek"
captured_at: "2026-07-20T23:47:25Z"
capture_tool: "hn-digest"
hn_id: 48986313
score: 1
comments: 0
posted_at: "2026-07-20T23:40:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A chess game you play with Claude as your wing-bot in the CLI

- HN: [48986313](https://news.ycombinator.com/item?id=48986313)
- Source: [github.com](https://github.com/goawaygeek/ground-control)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T23:40:49Z

## Translation

タイトル: Show HN: CLI でクロードをウィングボットとしてプレイするチェス ゲーム
記事タイトル: GitHub - goawaygeek/ground-control · GitHub
説明: GitHub でアカウントを作成して、goawaygeek/ground-control の開発に貢献します。
HN テキスト: こんにちは、HN!誰かが興味を持った場合に備えて、自分が作成したものを披露したいと思いました。私はクロード コード CLI 内でプレイするゲームを構築したかったのですが、チャネル機能がリリースされたときに、そのためにそれを使用する方法を調査し始めました。私は、クロードと協力して時間枠内のテーマに基づいてジョークを書き、それを視聴者に投稿して他の人たちと競い合うマルチプレイヤー コメディ バトル ゲームから始めました (視聴者は最も面白いジョークに投票します)。それは問題ありませんでしたが、実際にテストする相手がいなかったので、次に進みました。私の息子はチェスを習っているので、その代わりにクロードと一緒にチェスをプレイできるように調整しました。3 つの可能な手を提案し、そこから選択して手として送信できます (クロードではなく、実際にやりたいことを選択します)。それはずっと楽しくて、チェスについて実際にいくつかのことを学ぶことができました（私のチェスは上手ではありません、チェスを学ぶのにそれほど役立つ必要はありませんでしたが、なぜ特定の手がより良いのかなどについて、チェスを行ったり来たりすることはできました）。時間が経つにつれて、実際に何かを「エージェント的に」構築するのがどのようなものかを感じられるように、それを少し強化しました (この場合は、おそらくバイブの方が適切です。私はエージェントに実際に物事を暴走させたことはありません)。これで、サーバーをローカルで実行してクロード コードを起動してボットとチェスのゲームをプレイしたり、リモート サーバーに参加 (またはホスト) して友達とプレイしたりできるようになりました。ユーザー名を維持できるようにするために、GCP 上に MCP サーバーを配置し、Notion 上に非常に初歩的なストレージを配置しました (必要に応じて)。参加すると、Bをプレイできるようになります

または、ロビーに誰かがいるかどうかを確認して、ゲームに挑戦してください。すべてのコードと手順が記載された Github は、https://github.com/goawaygeek/ground-control にあります。インストールしたら、「--dangerously-load-development-channels」を実行する必要があります。これは、ローカルで実行している場合は、そのチャネルに何が入ってくるか、何が入ってくるかを知っているので大丈夫です。ホストされているバージョンからのプロンプト インジェクションの可能性を軽減するために、渡されるフリー テキスト (メッセージ、メッセージ、ジョーク、ユーザー名など)。それを確実に行う方法についてのヒントは大歓迎です。ボットで遊ぶのはそれ自体楽しいので、緊張している場合は、オンラインで見知らぬ人に挑戦する代わりに、そうしてください (ただし、当然の警告ですが、モデルは重要です。Haiku をプレイするのはひどいもので、Opus 以降はかなりまともでした!)。誰かが（ホストされているサーバー上で）オンラインであるかどうかを確認したい場合は、ここに簡単なフロントエンドを作成しました: https://groundcontrol.deepdeep.space そして、2 人がプレイしているビデオ (まあ、私自身がプレイしています) を YouTube に投稿しました: https://youtu.be/YMgbf-qUVqc 私にとっての論理的な次のステップは次のとおりです。
a) チェス エンジンを改善して、ボットを微調整したり、ボットが提案する手を評価したりできるようにします (「初心者の手」と「グランド マスター」の手の違いを示すため)。
b) CC の代わりにローカル モデルを実行するようにすべてを移動します。コメディ バトル テスト用にチャンネルを実行するのに 1 ～ 2 日費やしただけで、その道に留まりました。長期的にはローカル モデルにもっと興味があります。API クレジットを使用するものを作成したくなかったため、チャンネル経由で Anthropic プランのいずれかでプレイできるようにしました。
c) 人々が独自のローカル モデルを活用し、微調整/微調整できるまったく新しいゲームを作成します。おそらく私の年齢を表していると思いますが、R2 ユニットを持つという考えは私にとって

戦闘中の戦闘機の後ろ姿は、懐かしいですね。何かご興味がございましたら、コメントを残すか、ご連絡ください。

記事本文:
GitHub - goawaygeek/ground-control · GitHub
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
ゴーアウェイオタク
/
地上管制
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ タグ G

o ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
38 コミット 38 コミット docs docs scripts scripts src src .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json Dockerfile Dockerfile ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Ground Control は、Claude Code CLI を通じて Channels フラグを使用することで、LLM と提携し、他の人間/LLM パートナーシップに対してゲームをプレイできるホスト型サーバーです。
十分に明確に伝わらない場合のために説明すると、これは他の何よりも UX の実験です (少なくとも現時点では)。その作成の背後にある前提は次のとおりです。
LLM 推論コストをホストされたサーバーから移動する
他社との競争において LLM と提携することを奨励します
その過程で何かを学ぶかもしれない
現在、Ground Control には、チェスとコメディ バトルという、異なるスタイルの 2 つのゲームが同梱されています。どちらも同じサーバーを通じて実行されますが、動作方法は異なります。
Chess では、ロビーに参加して、チェスをプレイする他のプレイヤーを見つけることができます。それはあなたとあなたの競争相手のゲームをスピンアップし、できるだけ多くの同時ゲームを実行します。ロビーは純粋に対戦相手を見つけるための待機場所です。これが実際にどのように見えるかについてのビデオがYouTubeにあります
Comedy Battle はロビーをゲームセンターとして使用します。このゲームでは、2 人の対戦相手がテーマに沿って最も面白いジョークを言おうとし、ロビーの観客がどちらが面白いか投票します。より面白いジョークを言った競技者はそのまま残り、観客からランダムに選ばれた新たな対戦相手と対戦します。サーバーに参加すると、視聴者として参加し、ジョークを作成する順番が来るまで、ジョークを審査してもらいます。
サーバーに書き込みが行われています

新しいゲーム (囲碁、チェッカーなど) を提案したい場合は、/src/game/[your-game]/ でゲームを作成し、PR を送信できるようにします (詳細については docs/game-submission-lifecycle.md を参照してください)。ロボットのデザインや戦闘戦略などで LLM が戦い抜くことを奨励するようなゲームを見てみたいと思っています。私自身もそれに取り組む予定です :)
新しいゲームには、サーバー側のボット対戦相手が同梱される予定です — docs/bot-pattern.md を参照してください。空のロビーは、新規訪問者が離れる唯一の最大の理由です。 LLM が提供できるボット (「周りに誰もいない - ボットをプレイしたい?」) があれば、そのギャップは埋まります。
現時点では、プレイするにはクロード コードのサブスクリプションが必要です。長期的な夢は、これをローカルでホストされたモデルで動作させ、API 料金やサブスクリプション費用をなくし、LLM を調整して、独自の調整された LLM とペアになった他のユーザーと (モデル、マシン、またはプロンプト/カスタム RAG などを通じて) 対戦できるようにすることです。
参加したい場合は、PR を提出するか、メンテナーに電子メールを送信してください。
開始方法については続きをお読みください...
クロード コード v2.1.80+ (チャネル サポート用)
git clone https://github.com/goawaygeek/ground-control.git
CD 地上管制
npmインストール
2. ゲームに参加する
このリポジトリの .mcp.json ファイルは、すでに https://groundcontrol.deepdeep.space にあるホストされたゲーム サーバーを指しています。必要なゲームで Claude Code を起動するだけです。
# Chess — コーチとしてクロードと提携します
クロード --dangerously-load-development-channels サーバー:チェス
# コメディバトル — クロードと一緒にジョークを書き、視聴者が投票
クロード --dangerously-load-development-channels サーバー:コメディーバトル
クロードはプレイヤー名を尋ね、あなたを登録し、プレイヤートークンを渡します。
登録すると、クロードは次のようなプレーヤー トークンを出力します。
PLAYER_TOKEN=8edfc25d-8b6d-4a3e-9479-47c98f0889ec
その行を .env ファイルに追加します。

プロジェクトのルート。次回クロード コードを起動すると、自動的に再接続されます。名前は予約されています。
# .env
PLAYER_TOKEN=8edfc25d-8b6d-4a3e-9479-47c98f0889ec
.env ファイルは gitignored されるため、トークンは非公開のままになります。これは事実上パスワードですので、共有しないでください。
はい、はい。プレーヤーのトークンが非常にハック的で、ユーザー名を永続化するために Notion のテーブルを使用するだけであることはわかっています。ある時点で適切な認証が行われる予定ですが、実験のためにユーザー名を要求する方法が人々に提供されないようにしたくありませんでした。
https://groundcontrol.deepdeep.space のランディング ページには、現在のプレイヤー数とゲームのフェーズが表示されます。ゲームごとのライブ観客 UI が登場します。
このプラットフォームは、GameModule インターフェイスを中心に設計されています。どのゲームでも次の方法で追加できます。
GameModule を実装する src/game/<your-game>/index.ts の作成 ( src/game/types.ts )
src/game/index.tsに登録
ゲームのエントリで .mcp.json を更新する
サーバーは /<your-game>/ にルームを自動作成します。以下を定義します。
プレーヤーのライフサイクル — プレーヤーが参加、脱退、または役割を割り当てられたときに何が起こるか
ゲーム フロー — フェーズ遷移、タイムアウト、ステート マシン
ツール — クロードがゲーム上で呼び出すことができる動詞 ( make_move 、 submit_joke 、 vote など)
指示 — クロードに遊び方を指示するプロンプト
無料で手に入るもの: GameRoom
作成した各 GameModule は GameRoom によってラップされ、煩雑なインフラストラクチャをすべて処理するため、ゲーム ロジックはルールと状態に重点を置いたままになります。ゲームごとに 1 つの GameRoom があり、それらは隔離されているため、チェスプレイヤーはコメディーバトルイベントを見ることができません。
ゲームごとのセッション管理 - 登録、再接続、トークン認証、ロールの割り当て、クリーンな切断/猶予期間。セッションはルームごとに分離されるため、同じプレイヤー名が異なるゲームに独立して存在できます。
3 つのブロードキャスト スコープ - イベントを単一のプレーヤーにプッシュします

(broadcastToPlayer)、すべてのアクティブ プレーヤー (broadcastToAllPlayers)、Web 観客のみ (broadcastToAudience)、または全員 (broadcastEverywhere)。
ターゲット イベントとグローバル イベント — ゲームは各ハンドラーからイベントを返します。イベントに _targetPlayer トークンがある場合、イベントはそのプレイヤーのみに送信されます (カードの配り、秘密の役割など)。それ以外の場合は全員にブロードキャストされます。
フェーズ タイマー — 任意のイベントから _nextPhaseTimeout を返し、GameRoom は onPhaseTimeout() ハンドラーへのコールバックをスケジュールします。 _phaseTimerKey を使用して、複数の独立したタイマーを同時に実行します (例: 部屋内のゲームごとに 1 つのチェス クロック)。
Audience Web UI — 誰でも SSE 経由で /<your-game>/audience に接続してリアルタイムで視聴できます。
つまり、GameModule は 4 つの質問に答えるだけで済みます。
プレイヤーが参加するとどうなりますか? (onPlayerJoin)
プレイヤーが行動すると何が起こるでしょうか？ (アクション上)
フェーズタイマーが作動すると何が起こるでしょうか? (onPhaseTimeout)
現在の状態は何ですか? ( getState )
それ以外のすべて (ブロードキャスト、セッション ライフサイクル、タイマー、分離) は GameRoom によって処理されます。
Chess はエンジンとして chess.js を使用し、チャレンジを通じてプレイヤーをペアリングし、1 つの部屋内で複数のゲームを同時に実行します。
Comedy Battle は、フェーズ タイマーによって駆動されるフェーズ移行 (ロビー → 書き込み → 公開 → 投票 → 結果) を備えた手動のステート マシンです。
.mcp.json は、デフォルトでホストされているサーバーを指します。自分で実行するには:
npm実行サーバー
サーバーは http://localhost:8087 で起動します。 .mcp.json の *-local MCP エントリ ( chess-local 、 Comedy-battle-local ) を使用して、ホストされているサーバーではなくローカル サーバーに接続します。
パブリックにアクセス可能なサーバー (友人用など) を実行したい場合は、groundcontrol.deepdeep.space にあるホストされたインスタンスが、HTTPS 用に Caddy を前面に備えた単一の GCE VM 上で実行されます。 docs/self-hostin を参照してください。

正確なセットアップについては g-gcp.md を参照してください。
デフォルトでは、ローカル サーバーはプレーヤーを data/players.json に保存します。 Notion をバッキング ストア (再起動、デプロイ、環境全体で共有) として使用するには、これを独自のサーバーに使用したい場合は、.env に詳細を追加するだけです。
NOTION_TOKEN=your_notion_token \
NOTION_DATABASE_ID=your_database_id \
npm実行サーバー
Notion データベースには次のプロパティが必要です。
クロード・コード・セッション（プレイヤー）
|
MCP チャネル サーバー
(src/channel-server.ts)
|
HTTP + SSE
|
ゲームサーバー (src/server.ts)
／＼
GameRoom: チェス GameRoom: コメディーバトル
| |
チェスゲームコメディバトル
(src/game/chess/) (src/game/comedy-battle/)
ゲームサーバー — パスプレフィックスルーティング ( /chess/join 、 /comedy-battle/start-round など) を介してすべてのゲームをホストする HTTP サーバー。プレーヤー セッション、SSE イベント ストリーム、フェーズ タイマーを管理します。
チャネル サーバー — クロード コード セッションをゲーム サーバーにブリッジする MCP サーバー。ゲーム ツール ( make_move 、 submit_joke 、 vote 、 send_message ) を公開し、クロード コード チャネル経由でリアルタイム イベントを配信します。
GameModule インターフェイス — 交換機能の継ぎ目。 src/game/types.ts に GameModule を実装することで、任意のゲームを追加できます。
ソース/
server.ts # ゲームごとのルーティングを備えた HTTP サーバー
channel-server.ts # MCP チャネルサーバー (クロードコードブリッジ)
channel-client.ts # ゲームサーバー用の HTTP/SSE クライアント
auth.ts # トークンベースのセッション管理
store.ts # プレーヤーの永続化 (LocalJson + Notion)
game-room.ts # GameRoom: ゲーム + セッション + ブロードキャスト + タイマー
format-event.ts # クロードコード表示用のイベントフォーマット
config.ts # サーバー URL 解決 (CLI > env > デフォルト)
ゲーム/
types.ts # GameModule インターフェイス
Index.ts # ゲームファクトリー + レジストリ
chess/index.ts # チェスゲーム (chess.js エンジン)
Comedy-battle/index.ts # コメディバトルステートマシン
ウェブ/
Index.html # ランディング ページ (TUI スタイル)

、世論調査/統計)
API
すべてのゲーム エンドポイントには /<gameId>/ というプレフィックスが付きます。
npm install # 依存関係をインストールする
npm runserver # :8087 でゲームサーバーを起動します
npm test # テストを実行する (vitest)
ライセンス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to goawaygeek/ground-control development by creating an account on GitHub.

Hey HN! Wanted to show off a thing I built in case it interests anyone. I wanted to build a game you play inside the claude code CLI and back when they launched the channels feature I started investigating ways to use that to do so. I started with a multiplayer comedy battle game where you would collaborate with Claude to write jokes based on a theme in a timeframe and then post them to an audience and compete against other folks (with the audience voting for the funniest joke), it was OK but I didn't have anyone to really test it with so I moved on. My son has been learning chess so instead I tweaked it to let you play chess with Claude suggesting three possible move options for you to choose from and submit as your move (you choose what you actually want to do not Claude). It was much more fun and allowed me to actually learn a thing or two about chess (my chess isn't great, it didn't need to do much to help me learn but I could go back and forth with it on why a certain move was better, etc.). Over time I beefed it up a bit just so I could get a feel for what it is like actually building something 'agentically' (vibing is probably more apt in this case, I never really let an agent run wild with things). So now you can run the server locally and spin up Claude Code to play a game of chess against a bot or you can join (or host) a remote server and play your friends. I threw an MCP server on GCP and some very rudimentary storage on Notion to allow you to maintain a username (should you choose to). When you join you can play a bot or see if anyone is in the lobby and challenge them to a game. Github with all the code and instructions is over at: https://github.com/goawaygeek/ground-control Once you have things installed you'll have to `--dangerously-load-development-channels` which, well, if you're running it locally all good you know what's coming in and out of that channel, to mitigate the potential for prompt injection from the hosted version I've put in qualifiers around any of the free text that gets passed through (messages, jokes, usernames, etc.). Any tips on how to sure that up are more than welcome! Playing with the bot is fun on its own so if you're nervous just do that instead of challenging a random online stranger (although fair warning, models matter, playing Haiku was terrible, Opus and above were pretty decent!). I put up a simple front end if you want to check whether anyone is online (on my hosted server) here: https://groundcontrol.deepdeep.space and I put up a video of two people playing (well, me playing myself) on youtube: https://youtu.be/YMgbf-qUVqc The logical next steps for me would be to:
a) improve the chess engine so the bot can be tweaked, or the moves it suggests could be graded (to show the difference between a 'novice move' and a 'grand master' move.
b) move everything so it runs local models instead of CC, the day or two I spent getting channels running for the comedy battle test just kind of kept me down that path, long term I'm much more interested in local models I just didn't want to create something that used API credits hence the ability to play on any of the Anthropic plans via Channels.
c) create a whole new game that allows people to leverage their own local models and tweak/fine tune them. It probably shows my age but the idea of having an R2 unit in the back of your fighter while you go into battle is, um, nostalgic I guess. Please leave comments or reach out if anything is of interest to you!

GitHub - goawaygeek/ground-control · GitHub
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
goawaygeek
/
ground-control
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits docs docs scripts scripts src src .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json Dockerfile Dockerfile LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Ground Control is a hosted server that lets you partner with your LLM and play games against other human/LLM partnerships by using the Channels flag through the Claude Code CLI.
In case it doesn't come across clearly enough: this is an experiment in UX much more than anything else (at least as of right now). The premise behind its creation was to:
Move LLM inference cost off the hosted server
Encourage you to partner with your LLM in competition against others
Maybe learn something along the way
Right now there are two games of different styles that ship with Ground Control: Chess and Comedy Battle . Both run through the same server but work in different ways.
Chess lets you join the lobby and find other players to play chess against. It will spin up games of you vs. your competitor and run as many concurrent games as it can. The lobby is purely a waiting place to find opponents. Here's a video of what this looks like in practice on youtube
Comedy Battle uses the lobby as the game centre. In this game, two opponents try to tell the funniest joke around a theme, and the audience in the lobby votes on which one is funnier. The competitor who tells the funnier joke stays on to face a new opponent chosen at random from the audience. Once you join the server you'll enter as an audience member and get given jokes to judge until your turn to create a joke comes up.
The server is written in such a way that if you want to propose a new game (go, checkers, etc.) you can create the game in /src/game/[your-game]/ and submit a PR (see docs/game-submission-lifecycle.md for more details). I would love to see a game that encourages LLMs to battle it out in things like Robot design or battle strategy, I'll be working on that myself too :)
New games are expected to ship with a server-side bot opponent — see docs/bot-pattern.md . Empty lobbies are the single biggest reason a new visitor leaves; having a bot the LLM can offer ("nobody's around — want to play a bot?") closes that gap.
As of right now you need a Claude Code subscription to play. The long term dream is to make this work with locally hosted models, removing any API fees or subscription costs and allowing you to tweak your LLM to battle other people paired with their own tweaked LLMs (either through model, machine or prompt/custom RAGs, etc.).
If you want to get involved, file a PR or email the maintainer.
Read on for how to get started...
Claude Code v2.1.80+ (for Channels support)
git clone https://github.com/goawaygeek/ground-control.git
cd ground-control
npm install
2. Join a game
The .mcp.json file in this repo already points at the hosted game server at https://groundcontrol.deepdeep.space . Just launch Claude Code with the game you want:
# Chess — partner with Claude as your coach
claude --dangerously-load-development-channels server:chess
# Comedy Battle — write jokes with Claude, audience votes
claude --dangerously-load-development-channels server:comedy-battle
Claude will ask you for a player name, register you, then give you a player token .
When you register, Claude will print a player token like:
PLAYER_TOKEN=8edfc25d-8b6d-4a3e-9479-47c98f0889ec
Add that line to a .env file in the project root. Next time you launch Claude Code, it'll reconnect as you automatically — your name is reserved.
# .env
PLAYER_TOKEN=8edfc25d-8b6d-4a3e-9479-47c98f0889ec
The .env file is gitignored so your token stays private. Don't share it — it's effectively your password.
Yes, yes. I know the player tokens are super hacky and just use a table in Notion to persist usernames. Proper auth will follow at some point I just didn't want folks to have no way to claim a username for an experiment.
The landing page at https://groundcontrol.deepdeep.space shows current player counts and game phases. Per-game live spectator UIs are coming.
The platform is designed around a GameModule interface. Any game can be added by:
Creating src/game/<your-game>/index.ts that implements GameModule ( src/game/types.ts )
Registering it in src/game/index.ts
Updating .mcp.json with an entry for your game
The server auto-creates a room at /<your-game>/ . You define:
Player lifecycle — what happens when they join, leave, or are assigned a role
Game flow — phase transitions, timeouts, state machine
Tools — the verbs Claude can invoke on your game ( make_move , submit_joke , vote , etc.)
Instructions — the prompt that tells Claude how to play
What you get for free: GameRoom
Each GameModule you write is wrapped by a GameRoom , which handles all the messy infrastructure so your game logic stays focused on rules and state. One GameRoom per game — they're isolated, so chess players don't see comedy-battle events.
Per-game session management — registration, reconnection, token auth, role assignment, and clean disconnect/grace periods. Sessions are isolated per room, so the same player name can exist independently in different games.
Three broadcast scopes — push events to a single player ( broadcastToPlayer ), to all active players ( broadcastToAllPlayers ), to web spectators only ( broadcastToAudience ), or to everyone ( broadcastEverywhere ).
Targeted vs. global events — your game returns events from each handler. If an event has a _targetPlayer token it goes to that player only (e.g. dealing cards, secret roles); otherwise it broadcasts to everyone.
Phase timers — return a _nextPhaseTimeout from any event and GameRoom schedules a callback into your onPhaseTimeout() handler. Use _phaseTimerKey to run multiple independent timers simultaneously (e.g. one chess clock per game-within-the-room).
Audience web UI — anyone can connect to /<your-game>/audience over SSE to watch in real time.
This means your GameModule only needs to answer four questions:
What happens when a player joins? ( onPlayerJoin )
What happens when a player acts? ( onAction )
What happens when a phase timer fires? ( onPhaseTimeout )
What's the current state? ( getState )
Everything else — broadcasting, session lifecycle, timers, isolation — is handled by GameRoom .
Chess uses chess.js as the engine, pairs players via challenges, runs multiple concurrent games inside one room.
Comedy Battle is a hand-rolled state machine with phase transitions (LOBBY → WRITING → REVEAL → VOTING → RESULTS) driven by phase timers.
The .mcp.json points at the hosted server by default. To run your own:
npm run server
The server starts on http://localhost:8087 . Use the *-local MCP entries in .mcp.json ( chess-local , comedy-battle-local ) to connect to your local server instead of the hosted one.
If you want to run a publicly-reachable server (e.g. for friends), the hosted instance at groundcontrol.deepdeep.space runs on a single GCE VM with Caddy in front for HTTPS — see docs/self-hosting-gcp.md for the exact setup.
By default, the local server stores players in data/players.json . To use Notion as a backing store (shared across restarts, deploys, and environments), if you want to use this for your own servers you can just add the details to your .env:
NOTION_TOKEN=your_notion_token \
NOTION_DATABASE_ID=your_database_id \
npm run server
Your Notion database must have these properties:
Claude Code Session (Player)
|
MCP Channel Server
(src/channel-server.ts)
|
HTTP + SSE
|
Game Server (src/server.ts)
/ \
GameRoom: chess GameRoom: comedy-battle
| |
ChessGame ComedyBattle
(src/game/chess/) (src/game/comedy-battle/)
Game Server — HTTP server that hosts all games via path-prefix routing ( /chess/join , /comedy-battle/start-round , etc.). Manages player sessions, SSE event streams, and phase timers.
Channel Server — MCP server that bridges Claude Code sessions to the game server. Exposes game tools ( make_move , submit_joke , vote , send_message ) and delivers real-time events via Claude Code Channels.
GameModule interface — the swappability seam. Any game can be added by implementing GameModule in src/game/types.ts .
src/
server.ts # HTTP server with per-game routing
channel-server.ts # MCP channel server (Claude Code bridge)
channel-client.ts # HTTP/SSE client for game server
auth.ts # Token-based session management
store.ts # Player persistence (LocalJson + Notion)
game-room.ts # GameRoom: game + sessions + broadcast + timers
format-event.ts # Event formatting for Claude Code display
config.ts # Server URL resolution (CLI > env > default)
game/
types.ts # GameModule interface
index.ts # Game factory + registry
chess/index.ts # Chess game (chess.js engine)
comedy-battle/index.ts # Comedy battle state machine
web/
index.html # Landing page (TUI-styled, polls /stats)
API
All game endpoints are prefixed with /<gameId>/ :
npm install # Install dependencies
npm run server # Start game server on :8087
npm test # Run tests (vitest)
License
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
