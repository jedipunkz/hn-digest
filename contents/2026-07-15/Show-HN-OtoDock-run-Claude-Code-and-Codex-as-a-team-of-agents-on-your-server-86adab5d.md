---
source: "https://github.com/OtoDock/oto-dock/"
hn_url: "https://news.ycombinator.com/item?id=48923047"
title: "Show HN: OtoDock, run Claude Code and Codex as a team of agents on your server"
article_title: "GitHub - OtoDock/oto-dock: Your personal AI agent platform — self-hosted, BYO Claude/Codex subscription · GitHub"
author: "dimitrismrtzs"
captured_at: "2026-07-15T17:07:34Z"
capture_tool: "hn-digest"
hn_id: 48923047
score: 2
comments: 0
posted_at: "2026-07-15T16:11:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OtoDock, run Claude Code and Codex as a team of agents on your server

- HN: [48923047](https://news.ycombinator.com/item?id=48923047)
- Source: [github.com](https://github.com/OtoDock/oto-dock/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T16:11:23Z

## Translation

タイトル: HN を表示:otoDock、サーバー上でエージェントのチームとしてクロード コードとコーデックスを実行
記事タイトル: GitHub -otoDock/oto-dock: パーソナル AI エージェント プラットフォーム — 自己ホスト型、BYO Claude/Codex サブスクリプション · GitHub
説明: パーソナル AI エージェント プラットフォーム — 自己ホスト型、BYO Claude/Codex サブスクリプション -otoDock/oto-dock
HN テキスト: こんにちは、HN、私はディミトリスです。私はクロード コードとコーデックス エージェントを使用しています。最初からしばらくの間、主にコーディングのために端末内からそれらを使用していました。過去 3 年間、私は構築を続けてきたので、ホームラボとビジネス サーバーがあり、すでに多くの VM があり、管理するものがたくさんあります。そこで私は、クロード コードとコーデックスを柔軟に使用して、それらをすべての VM やインフラに簡単に接続すると同時に、コーディングや日常使用のパーソナル アシスタントとして使用し続けるという、私自身の理想的なバージョンの構築を開始することにしました。そこで、ここ数か月間、OtoDock を構築しました。これは、実際の CLaude コードと Codex をエンジンとして実行する自己ホスト型プラットフォームです。その上に、WebSocket を使用してライブ ダッシュボードを構築し、CLI を深く統合しているため、ターミナルで作業するのではなく、ダッシュボードを通じて CLI を操作する方がさらに便利です。 (Anthropic がハーネスに対して行おうとしていたサブスクリプションの変更のため) 最終的にはインタラクティブ ターミナルも統合することになり、リモートのダッシュボードから直接インタラクティブ ターミナルを使用することもできます。 otoDock は、マルチテナントの自己ホスト型エージェント プラットフォームであり、企業がエージェントを作成および管理することを目的としています。エージェントはバックグラウンドで自律的に動作してさまざまなジョブを実行できます。また、プラットフォームのユーザー (チーム メンバー) と協力して動作することもできるため、多くのチーム メンバーが特定の仕事や部門 (例: ソーシャル メディア エージェント、開発エージェントなど) 向けの共同エージェントを使用できます。

プラットフォームの各ユーザーは、独自の anthropic サブスクリプションと openai サブスクリプションを持ち込むことができ、これらのサブスクリプションは cli とのユーザー セッションに自動的に使用されます。 (または Ollama を介したローカル モデル) MCP フレームワークも構築しました。これにより、アプリケーション上で簡単に累積 mcps を迅速に追加できるようになり、利用可能なすべての mcps をそこに配置するコミュニティ mcp リポジトリを保持しています。ダッシュボードには多くの機能もあります。たとえば、エージェントと会話するための完全に機能するチャット、ユーザーがエージェントと作業するときに PDF、Word、Excel、Powerpoint ファイルを編集および表示するための Collabora プレビューをサポート、適切なファイル ブラウザーを備えたユーザーごとのワークスペース (ダッシュボードに表示される各ユーザーとエージェントの Google ドライブなど)、ユーザー間で共有されるエージェントごとのナレッジとワークスペース フォルダー、ユーザーごとおよびエージェントごとのメモリ システム、実際の間隔でスケジュールされたタスク、スケジュールされたタスクを開始できる Webhook トリガーなどです。エージェントの会議、エージェント間の会議 (エージェントが共同作業できるチャットでの直接のライブ会議)、エージェントと会話するための TTS および STT による音声など... 私自身セキュリティ偏執症なので、エージェントが可能な限り安全に実行され、実際のビジネス サーバーに安全に展開できることを確認するために多くの時間を費やしました。エージェントは、デフォルトでネットワーク分離されたカーネル サンドボックス上で実行され、ローカル ネットワーク アクセスを必要とする特定の mcps に対して、管理者によって特定のローカル アップ例外が追加されます。もちろんインターネットへのアクセスは常にオンです。ライセンスはフェアソースであり、OSI オープンソースではありません。セルフホスティングは 5 ユーザーまで無料で、コードはすべて公開されているため、すべてを確認できます。試してみると、Docker compose ファイルが 1 つあり、イメージは GHCR 上にあり、セットアップ ウィザードによって管理者がローカルに作成され、どこにもサインアップや電子メールを送信する必要はありません。適切なアプリケーションを公開するのは初めてですが、とても気に入っています

皆さんもぜひ試してみて、コードを読んでセットアップで使ってみてください。私はこれに多くの時間を費やしてきました。得られるすべてのフィードバックをお待ちしています。 AI 免責事項:otoDock の大部分は otodock 自体を使用して書かれていますが、過去 2 か月間、アプリケーションは十分に成熟しており、私は開発のために毎日otoDock のエージェントを使用し始めました。開発に Claude Code Cli または Codex を使用している場合は、お気軽にお試しください。ワークフローで非常に便利で役立つことがわかります。 (リモート マシン機能は 1 週間以内に利用可能になり、プラットフォームのエージェントをローカル マシンで実行できるようになります)。 GitHub: https://github.com/OtoDock/oto-dock
ウェブサイト：https://otodock.io

記事本文:
GitHub -otoDock/oto-dock: パーソナル AI エージェント プラットフォーム — 自己ホスト型、BYO Claude/Codex サブスクリプション · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オトドック
/
オトドック
公共
通知
署名する必要があります

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
29 コミット 29 コミット .github .github オーディオ オーディオ ダッシュボード ダッシュボード mcps mcps プロキシ プロキシ スクリプト スクリプト .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md VERSIONS.md VERSIONS.md config.env.example config.env.example docker-compose.build.yml docker-compose.build.yml docker-compose.t1.yml docker-compose.t1.yml docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
otoDock — 協力エージェント
Claude Code と Codex を会社全体のエージェントのチームとして実行します —
独自のサーバー、すでに支払ったサブスクリプションで、本物のツールと本物のセキュリティを備えています。
otodock.io で音声付きの完全なデモをご覧ください。
otoDock は、AI エージェントのチームを実行するための自己ホスト型プラットフォームです。
あなたが管理するインフラストラクチャ。実際のクロード コードとコーデックスをそのまま実行します。
エンジン — したがって、エージェントは CLI が実行できるすべてのことを継承し、それらをラップします
ライブダッシュボード、共有サーバー用に構築されたセキュリティモデル、および配管
これにより、コーディング ツールがスケジュール、トリガーなどの同僚のチームに変わります。
会議、記憶、書類、音声。
チームのクロード コード/コーデックス、自己ホスト型。みんなで繋いでる
既に支払っている Claude または ChatGPT プラン (または API キー、またはローカル モデル)
Ollama 経由)、1 つのダッシュボードからエージェントと連携します。
ホームラボに安全に触れることができるエージェント。すべてのエージェントは
デフォルトでネットワークから隔離されたロックダウンされたサンドボックス - あなたが許可します
一度に 1 つのフォルダーまたは 1 つのサービスにアクセスします。
設計上の主権者。チャット、ファイル、メモリ、資格情報はライブです

に
ハードウェアを実行し、そこに留まります。
エージェントが実際に動作する様子を観察してください。すべてのステップがライブでストリーミングされ、推論、
各ツールの呼び出し、赤/緑の差分としてのファイル編集、計画と To Do リストのチェック
エージェントが移動するとオフになります。機密性の高いアクションをインラインで承認するか、信頼できるものに任せます
エージェントが実行されます。 1 人のエージェントでは不十分な場合は、専門家を 1 つの部屋に配置します。
マルチエージェント会議では、エージェントがそれぞれに対応するモデレータ付きのディスカッションが実行されます。
その他、並行して回答し、集中しながら、視聴したり参加したりできます。
拡張可能なツールの詳細、プラン モード、サブエージェントを使用したライブ ストリーミング チャット
エージェントごとの ID とコストを備えたマルチエージェント会議
音声: すべてのチャットでのハンズフリー会話、ディクテーション、読み上げ
透明で編集可能なエージェント メモリ - すべての書き込みが表示され、バージョン管理され、あなたのものになります
タブを閉じてもエージェントは動作し続けます。実際の作業のスケジュールを設定する
間隔 — 17 時間ごと、3 日ごと、まさにその意味どおり —
完全な会話として保存された実行は、開いて続行できます。からの消防職員
他の場所で何かが発生したときに Webhook を送信します。通知は次のようにエスカレーションされます
4 つの重大度、見逃せない危険アラームまで。
エージェントは実際の Word、Excel、PowerPoint、PDF ファイル (表、
グラフ、書式設定 - 会話の中でドキュメントがすぐに開きます。
あなたとあなたのチームが入力できるライブエディター。画像生成と
プロフェッショナルな編集パイプラインを利用できます。
すでにお持ちの AI を導入: 消費者向けの Claude/ChatGPT サブスクリプション
(各人が独自に接続します)、管理者が共有できるプロバイダー API キー、または
完全にローカルなモデル。 MCP ツール サーバーを介してエージェントを拡張します。
マニフェストを作成し、エージェントごとにツールを割り当て、既製のエージェントとツールをインストールします
のコミュニティカタログより
ワンクリック: ブラウザー、GitHub、Notion、そして成長し続けるカタログ。
シャになるように作られています

赤 - 放しても安全です。すべてのサーバー側エージェントは
常時接続のネットワーク分離を備えた独自のカーネル サンドボックス。アクセスを 1 つ許可します
一度にサービス。 SSO/OIDC、2要素認証、エージェントごとの役割、暗号化
認証情報、範囲指定された API キー、およびユーザーごと/エージェントごとの予算はすべて、
箱。
必要なのは作成ファイルだけです。画像は GHCR から取得されます。
mkdir オトドック && cd オトドック
カール -fsSLO https://raw.githubusercontent.com/OtoDock/oto-dock/main/docker-compose.yml
echo " POSTGRES_PASSWORD= $( openssl rand -hex 24 ) " > .env
ドッカー構成 -d
次に、http://localhost:8400 を開きます。セットアップ ウィザードによって管理者が作成されます。
アカウントにアクセスすると、数分後に最初のエージェントとチャットすることになります。セット
ユーザーが別のホストを参照する場合は、.env の DASHBOARD_PUBLIC_URL。
PROXY_PORT : 公開ポートを移動します ( config.env.example
すべてのノブを文書化します。プラットフォームは残りのシークレットを自動生成します。
最初のブート)。
代わりにソースからビルドします (貢献者):
git clone https://github.com/OtoDock/oto-dock.git && cd oto-dock
printf ' POSTGRES_PASSWORD=%s\n ' " $( openssl rand -hex 24 ) " > config.env
scripts/compose.sh up -d --build
ベアメタル（開発）
Debian/Ubuntu では、1 つのスクリプトが固定されたツールチェーンである Postgres をブートストラップします。
コンテナーと構築されたダッシュボード:
git clone https://github.com/OtoDock/oto-dock.git && cd oto-dock
scripts/dev-setup.sh # --service を追加して systemd ユニットをインストールします
インストールされる内容については scripts/README.md を参照してください。
テスト スイートを実行するための proxy/tests/README.md。
ダッシュボード/React ダッシュボード — チャット、エージェント、タスク、ファイル、管理者
プロキシ/プラットフォーム コア (FastAPI) - セッション、セキュリティ、スケジューリング、
エージェント サンドボックス、およびダッシュボードが通信する WebSocket ハブ
mcps/ MCP ツール サーバー:otoDock のカスタム セット (ファイル、メモリ、タスク、
会議、通知など) + コミュニティ

yミラー
オーディオ/音声パッケージ — STT / TTS / 音声アクティビティ、プロバイダーに依存しない
scripts/ インストール、作成、バックアップ/復元、およびメンテナ ツール
エージェントはセッションごとのカーネル内で実際のクロード コード/コーデックス プロセスとして実行されます。
サンドボックスにアクセスし、MCP 経由でツールと通信し、すべてのステップをストリームに戻します。
ダッシュボード。 PostgreSQL はプラットフォームの状態を保持します。すべてが次のように出荷されます
コンテナー (または開発用にベアメタルを実行)。
次に出荷 — リリース後に導入される段階的な機能: リモート
マシン (自分のラップトップまたはサーバー上で完全なアクセス権を持つエージェントを実行)、ライブ
ターミナルビュー、実際の電話に応答するエージェント、Android アプリなど
統合 (Google Workspace、Slack、Linear、Microsoft 365 など)。
コミュニティ エージェント:otoDock/community-agents
コミュニティ MCP:otoDock/community-mcps
貢献は大歓迎です — COTRIBUTING.md を参照してください。セキュリティ
レポート: SECURITY.md 。
otoDock は公正なソースです: に基づいてライセンスされています
Functional Source License、v1.1、Apache 2.0 の将来の付与付き
(FSL-1.1-Apache-2.0)。これを使用、実行、変更、再配布できます。
商業的にOtoDockと競合すること以外のすべて - そして各バージョン
リリースから 2 年後には自動的にプレーン Apache 2.0 になります。
otoDock はクロード コードから始まりました: ターミナルを超えて使用したかった —
どこからでも、私自身のインフラストラクチャ上で、実際の自動化に接続されています。
そのアイデアに基づいてプラットフォームを構築しました。それ以来、それは私の場所になりました
独自のエージェントが動作します。OtoDock の大部分は、OtoDock によって構築、テスト、出荷されました。
otoDock 上で実行されているエージェント。
これは公正なソースなので、自分のハードウェア上で同じように実行できます。
あなた自身のサブスクリプション。使って、拡張して、何をすべきか教えてください
次に。
パーソナル AI エージェント プラットフォーム — 自己ホスト型、BYO Claude/Codex サブスクリプション
読み込み中にエラーが発生しました。お願いします

このページを読み込んでください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
4
オトドック 1.1.0
最新の
2026 年 7 月 15 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your personal AI agent platform — self-hosted, BYO Claude/Codex subscription - OtoDock/oto-dock

Hi HN, i am Dimitris, I have been using Claude Code and Codex agents, for some time now from the beggining i had been using them from inside my terminal mainly for coding. For the past 3 years i kept building so i have a homelab and a business server, which already has a lot of vms, so a lot of things to manage. So i decided to start building my own ideal version of using claude code, and codex flexibly and connect them easily with all my vms and infra but at the same time keep using them for coding, and as a personal assistant for daily use. So for the past months I built OtoDock. Its a self hosted platform, that runs real CLaude code and Codex as the engine. On top of them i build a live dashboard with websockets and integrated in depth the clis, so its even nicer to work with the cli through the dashboard instead of working on the terminal. I even ended up integrating interactive terminals as well (due to the subscription change Anthropic was trying to do for harnesses) so you can also use interactive terminals directly from the dashboard remotely. OtoDock is a multitenant self hosted agents platform, is meant for businesses to create and manager agents, that can act autonomously on the background doing various jobs, or they can also be collaborative, they can work with users (team members) of the platform so many team members can use a collaborative agent meant for a specific job or department (ex. social media agent, dev agent, etc.) Each user of the platform can bring his own anthropic and openai subscriptions that is automatically used for his user sessions with the clis. (or local models through Ollama) I have also built an MCP framework, so that its easy to add cummunity mcps fast on the application, and i keep a community mcp repo where i will put all available mcps there. The dashboard as well has a lot of features, for example it has a fully functional chat for talking to the agents, that supports collabora previews for editing and viewing PDF, Word, Excel, Powerpoint files as the user works with the agents, a workspace per user with a proper file browser (like each user and agents google drive showing in the dashboard), per agent shared knowledge and workspace folders across users, a per user and per agent memory system, scheduled tasks with real intervals, webhook triggers that can initiate a scheduled task of an agent, meetings between agents (live meetings directly in the chat where the agents can collaborate), voice with TTS and STT for talking to the agents and many more... I am a security paranoid myself so i spent a lot of time making sure the agents run as securely as possible and can be safely desployed in a real business server. The agents run on a kernel sandbox with network isolation by default, and specific local up excemptions are added by the admin for specific mcps that need local network access. Access to the internet of course is always on. The license, it is Fair Source, not OSI open source. Self hosting is free up to 5 users, all the code is public so you can review everything. Trying it is one docker compose file, images are on GHCR, the setup wizard creates your admin locally, no signup or email anywhere. Its the first time I publish a proper application and I would love people to try it and read the code and use it in their setups, I have spent a lot of time on this and i would love all the feedback I can get. AI Disclaimer: Large parts of OtoDock are written using otodock itself, for the last 2 months the application was mature enought where i started using an agent of OtoDock daily for the development. If you use Claude Code Cli or Codex for development feel free to try it, you might find very useful and helpful in your workflow. (remote machines feature will be available in a week, where you will be able to make the agents of the platform run in your local machine). GitHub: https://github.com/OtoDock/oto-dock
Website: https://otodock.io

GitHub - OtoDock/oto-dock: Your personal AI agent platform — self-hosted, BYO Claude/Codex subscription · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
OtoDock
/
oto-dock
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
29 Commits 29 Commits .github .github audio audio dashboard dashboard mcps mcps proxy proxy scripts scripts .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md VERSIONS.md VERSIONS.md config.env.example config.env.example docker-compose.build.yml docker-compose.build.yml docker-compose.t1.yml docker-compose.t1.yml docker-compose.yml docker-compose.yml View all files Repository files navigation
OtoDock — Collaborative Agents
Run Claude Code and Codex as a team of agents for your whole company —
on your own server, on the subscriptions you already pay for, with real tools and real security.
Watch the full demo with sound on otodock.io
OtoDock is a self-hosted platform for running a team of AI agents on
infrastructure you control. It runs the real Claude Code and Codex as its
engine — so your agents inherit everything the CLIs can do — and wraps them in
a live dashboard, a security model built for shared servers, and the plumbing
that turns a coding tool into a team of coworkers: schedules, triggers,
meetings, memory, documents, voice.
Your team's Claude Code / Codex, self-hosted. Everyone connects the
Claude or ChatGPT plan they already pay for (or API keys, or local models
via Ollama), and works with agents from one dashboard.
Agents that can safely touch your homelab. Every agent runs in a
locked-down sandbox, isolated from your network by default — you grant
access one folder or one service at a time.
Sovereign by design. Chats, files, memory, and credentials live on
hardware you run, and stay there.
Watch your agents actually work: every step streams in live — the reasoning,
each tool call, file edits as red/green diffs, plans and to-do lists ticking
off as the agent moves. Approve sensitive actions inline, or let trusted
agents run. And when one agent isn't enough, put specialists in one room:
multi-agent meetings run a moderated discussion where agents address each
other, answer in parallel, and converge — while you watch or join in.
Live streaming chat with expandable tool detail, plan mode, subagents
Multi-agent meetings with per-agent identity and cost
Voice: hands-free conversations, dictation, and read-aloud in every chat
Transparent, editable agent memory — every write visible, versioned, yours
Your agents keep working when you close the tab. Schedule work on real
intervals — every 17 hours, every 3 days, exactly as you mean it — with every
run saved as a full conversation you can open and continue. Fire agents from
webhooks when something happens anywhere else. Notifications escalate through
four severities, up to a danger alarm that won't be missed.
Agents produce actual Word, Excel, PowerPoint, and PDF files — tables,
charts, formatting — and the document opens right in the conversation in a
live editor you and your team can type into. Image generation and a
professional editing pipeline ride along.
Bring the AI you already have: consumer Claude/ChatGPT subscriptions
(each person connects their own), provider API keys an admin can share, or
fully local models. Extend agents through MCP tool servers — drop in a
manifest, assign tools per agent — and install ready-made agents and tools
from the community catalog in
one click: a browser, GitHub, Notion, and a catalog that keeps growing.
Built to be shared — and safe to let loose. Every server-side agent runs in
its own kernel sandbox with always-on network isolation; you grant access one
service at a time. SSO/OIDC, two-factor auth, per-agent roles, encrypted
credentials, scoped API keys, and per-user/per-agent budgets are all in the
box.
The compose file is all you need — images are pulled from GHCR:
mkdir otodock && cd otodock
curl -fsSLO https://raw.githubusercontent.com/OtoDock/oto-dock/main/docker-compose.yml
echo " POSTGRES_PASSWORD= $( openssl rand -hex 24 ) " > .env
docker compose up -d
Then open http://localhost:8400 — the setup wizard creates your admin
account, and you're chatting with your first agent minutes later. Set
DASHBOARD_PUBLIC_URL in .env if users browse to a different host, and
PROXY_PORT to move the published port ( config.env.example
documents every knob; the platform auto-generates its remaining secrets on
first boot).
Building from source instead (contributors):
git clone https://github.com/OtoDock/oto-dock.git && cd oto-dock
printf ' POSTGRES_PASSWORD=%s\n ' " $( openssl rand -hex 24 ) " > config.env
scripts/compose.sh up -d --build
Bare metal (development)
On Debian/Ubuntu, one script bootstraps the pinned toolchain, a Postgres
container, and the built dashboard:
git clone https://github.com/OtoDock/oto-dock.git && cd oto-dock
scripts/dev-setup.sh # add --service to install a systemd unit
See scripts/README.md for what it installs and
proxy/tests/README.md for running the test suite.
dashboard/ React dashboard — chat, agents, tasks, files, admin
proxy/ Platform core (FastAPI) — sessions, security, scheduling,
the agent sandbox, and the WebSocket hub the dashboard talks to
mcps/ MCP tool servers: OtoDock's custom set (files, memory, tasks,
meetings, notifications, …) + community mirrors
audio/ Speech package — STT / TTS / voice activity, provider-agnostic
scripts/ Install, compose, backup/restore, and maintainer tooling
Agents run as real Claude Code / Codex processes inside per-session kernel
sandboxes, talk to their tools over MCP, and stream every step back to the
dashboard. PostgreSQL holds the platform state; everything ships as
containers (or runs bare-metal for development).
Shipping next — staged features you'll see land after launch: remote
machines (run agents with full access on your own laptop or servers), a live
terminal view, agents that answer real phone calls, the Android app, and more
integrations (Google Workspace, Slack, Linear, Microsoft 365, …).
Community agents: OtoDock/community-agents
Community MCPs: OtoDock/community-mcps
Contributions are welcome — see CONTRIBUTING.md . Security
reports: SECURITY.md .
OtoDock is fair source : licensed under the
Functional Source License, v1.1, with Apache 2.0 future grant
(FSL-1.1-Apache-2.0). You can use, run, modify, and redistribute it for
anything except competing with OtoDock commercially — and each version
automatically becomes plain Apache 2.0 two years after its release.
OtoDock started with Claude Code: I wanted to use it beyond the terminal —
from anywhere, on my own infrastructure, wired into real automation — so I
built the platform around that idea. It has since become the place where my
own agents work: large parts of OtoDock were built, tested, and shipped by
agents running on OtoDock.
It's fair source so you can run it the same way, on your own hardware and
your own subscriptions. Use it, extend it, and tell me what it should do
next.
Your personal AI agent platform — self-hosted, BYO Claude/Codex subscription
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
4
OtoDock 1.1.0
Latest
Jul 15, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
