---
source: "https://github.com/Panniantong/Agent-Reach"
hn_url: "https://news.ycombinator.com/item?id=49207806"
title: "Agent Reach: An open-source CLI that gives AI agents access to the internet"
article_title: "GitHub - Panniantong/Agent-Reach: Give your AI agent eyes to see the entire internet. Read & search Twitter, Reddit, YouTube, GitHub, Bilibili, XiaoHongShu — one CLI, zero API fees. · GitHub"
author: "Nina_antalpha"
captured_at: "2026-08-07T09:51:51Z"
capture_tool: "hn-digest"
hn_id: 49207806
score: 2
comments: 1
posted_at: "2026-08-07T09:13:41Z"
tags:
  - hacker-news
  - translated
---

# Agent Reach: An open-source CLI that gives AI agents access to the internet

- HN: [49207806](https://news.ycombinator.com/item?id=49207806)
- Source: [github.com](https://github.com/Panniantong/Agent-Reach)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T09:13:41Z

## Translation

タイトル: Agent Reach: AI エージェントにインターネットへのアクセスを提供するオープンソース CLI
記事のタイトル: GitHub - Panniantong/Agent-Reach: AI エージェントにインターネット全体を監視する目を与えます。 Twitter、Reddit、YouTube、GitHub、Bilibili、XiaoHongShu の読み取りと検索 — 1 つの CLI、ゼロの API 手数料。 · GitHub
説明: AI エージェントにインターネット全体を監視する目を与えます。 Twitter、Reddit、YouTube、GitHub、Bilibili、XiaoHongShu の読み取りと検索 — 1 つの CLI、ゼロの API 手数料。 - Panniantong/エージェントリーチ

記事本文:
GitHub - Panniantong/Agent-Reach: AI エージェントにインターネット全体を監視する目を与えます。 Twitter、Reddit、YouTube、GitHub、Bilibili、XiaoHongShu の読み取りと検索 — 1 つの CLI、ゼロの API 手数料。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
潘念通
/
エージェントリーチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
370 コミット 370 コミット .github/ workflows .github/ workflows .openteams/ spec .openteams/ spec Agent_re

ach Agent_reach config config docs docs scripts scripts testing testing .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.mdconstraints.txtconstraints.txtllms.txtllms.txtpyproject.toml pyproject.toml test.sh test.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
現時点で最も安定したアクセス方法が選択され、インストールされ、チェックされます。アクセス方法は置き換えられるため、心配する必要はありません。
クイックスタート · 英語 · 日本語 · 한국어 · サポートされているプラットフォーム · デザインコンセプト
AI エージェントは、コードの作成、ドキュメントの変更、プロジェクトの管理をすでに支援できますが、オンラインで何かを検索するように頼んでも、AI エージェントは盲目になります。
📺 「この YouTube チュートリアルの内容を確認してください」 → 表示されず、字幕も取得できません
🐦 「Twitter でこの製品についての人々の意見を検索するのを手伝ってください」 → 検索できません、Twitter API には支払いが必要です
📖 「Reddit にアクセスして、同じバグに遭遇した人がいないか確認してください」 → 403 がブロックされ、サーバー IP が拒否されました
📕 「Xiaohongshu でこの製品の評判を確認するのにご協力ください」 → 開けません。閲覧するにはログインする必要があります
📺 「サイト B に技術的なビデオがあります。それを要約するのを手伝ってください。」 → 取得できませんでした。一般的なダウンロード ツールはサイト B のリスク管理によって完全にブロックされました
🔍 「インターネットで最新の LLM フレームワークの比較を検索するのを手伝ってください」 → 有用な検索はありません。有料か低品質です。
🌐 「このウェブページに何が書かれているか教えてください」 → 大量の HTML タグを取得しましたが、まったく読めませんでした。
📦 「この GitHub リポジトリは何のためにあるのですか? 問題には何が記載されていますか?」 → 使えるけど認証設定が面倒
📡 「これらの RSS フィードを購読して、更新があったときに知らせてください」 → ライブラリをインストールして自分でコードを書く必要があります
各プラットフォームには、支払うべき API、バイパスすべきブロック、ログインすべきアカウント、クリーニングすべきデータなど、独自のしきい値があります。エージェントにツイートを読み取らせるためだけに、落とし穴を 1 つずつ通過し、ツールをインストールし、設定を調整する必要があります。

長い時間がかかるでしょう。
Agent Reach のインストールを手伝ってください: https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
これをエージェントにコピーすると、数分で Twitter の読み取り、Reddit の検索、YouTube の視聴、Xiaohongshu の閲覧ができるようになります。
Agent Reach の更新にご協力ください: https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/update.md
⭐ スター このプロジェクトでは、引き続き各プラットフォームの変更を追跡し、新しいチャネルにアクセスします。自分で監視する必要はありません。プラットフォームがブロックされている場合は修正し、新しいチャネルがある場合は追加します。
💰 完全無料
すべてのツールはオープンソースであり、すべての API は無料です。唯一費用がかかる可能性があるのはサーバー エージェント (月額 1 ドル) です。これはローカル コンピューターには必要ありません。
🔒 プライバシーとセキュリティ
Cookie はローカルにのみ存在し、外部にアップロードまたは送信されません。コードは完全にオープンソースであり、いつでもレビューできます。
🔄 継続的なアップデート
各プラットフォームは「優先 + 代替」マルチバックエンド ルーティングです。特定のアクセス方法が失敗した場合は、次のアクセス方法に置き換えますので、気にする必要はありません (2026-06 例: B サイトのリスク管理によって yt-dlp がブロックされた → bili-cli が切り替わり、ユーザーの操作がゼロになりました)
🤖 すべてのエージェントと互換性があります
Claude Code、OpenClaw、Cursor、Windsurf...コマンド ラインを実行できる任意のエージェントを使用できます
🩺 診断機能が付属
Agent-Reach Doctor は、どれが機能しているか、どれが機能していないか、および 1 つのコマンドでそれを修正する方法を示します。
サポートされているプラットフォーム
プラットフォーム
インストールの準備ができました
設定後にロックを解除する
マッチング方法
🌐 ウェブページ
あらゆるWebページを読む
—
設定は必要ありません
📺 YouTube
字幕抽出 + 動画検索
—
設定は必要ありません
📡 RSS
RSS/Atom フィードを読む
—
設定は必要ありません
🔍ウェブ全体を検索
—
Web 全体のセマンティック検索
自動構成 (MCP アクセス、無料、キー不要)
📦GitHub
パブリックリポジトリの読み取り + 検索
プライベートウェアハウス、発行/PR、フォーク
エージェントに「GitHub へのログインを手伝ってください」と伝えてください。
🐦ツイッター/X
ツイートを 1 つ読む
ツイートの検索、タイムラインの閲覧、長い記事の閲覧
エージェントに「Twitter の設定を手伝ってください」と伝えてください
📺サイトB
検索 + 動画詳細 (bili-cli、ログイン不要)
字幕 (OpenCLI)
エージェントに「助けて」と伝えてください

ステーションBにマッチします」
📖レディット
— (ゼロ構成パスなし: 匿名インターフェイスはブロックされています)
投稿やコメントを検索して読む
OpenCLI をデスクトップにインストールし、ブラウザでログインします。または rdt-cli + Cookie
📘フェイスブック
—
検索、ホーム、フィード、グループリスト
OpenCLI のデスクトップ インストール (Chrome ログイン状態の再利用)
📷インスタグラム
—
ユーザー検索、プロフィール、ユーザーの最近の投稿、探索
OpenCLI のデスクトップ インストール (Chrome ログイン状態の再利用)
📕 小さな赤い本
—
検索、閲覧、コメント
OpenCLI はユーザーの既存の Chrome セッションのみを使用します。 MCP/インベントリ ツールは Cookie エディターを使用します
💼リンクトイン
Jina Reader は公開ページを読み取ります
プロフィール詳細、会社ページ、求人検索
エージェントに「LinkedIn とのマッチングを手伝ってください」と伝えてください
💻V2EX
人気の投稿、ノード投稿、投稿の詳細と返信、ユーザー情報
—
設定は必要ありません
📈雪だるま
株価、銘柄検索、人気投稿、人気銘柄ランキング
—
エージェントに「雪玉を合わせるのを手伝って」と伝えてください
🎙️小さな宇宙ポッドキャスト
—
ポッドキャスト音声をテキストに変換 (ささやき文字起こし、無料キー)
エージェントに「リトル ユニバース ポッドキャストの設定を手伝ってください」と伝えてください。
マッチング方法がわかりませんか?ドキュメントを確認する必要はありません。エージェントに「XXX のマッチングを手伝ってください」と言うだけで、エージェントはあなたが何を必要としているかを理解し、段階的にガイドしてくれます。
🍪 Twitter は、ユーザーが Cookie エディターを通じて手動でエクスポートしたコンテンツのみを受け入れます。 Agent Reach は、ユーザーに対して Xiaohongshu ログインを実行せず、Xiaohongshu ブラウザの Cookie を読み取りません。 OpenCLI は、既存でユーザーによって明示的に制御されている Chrome セッションのみを使用します。 Agent-Reach configure xhs-cookie は OpenCLI/Chrome に Cookie を挿入しません。既存のセッションがない場合は、代わりに Cookie エディターを使用して xiaohongshu-mcp/stock ツールをエクスポートして構成します。
Twitter Cookie は、エージェントに連絡する医師が設定が完了したかどうかを確認するためにのみ保存されます。アップストリームの twitter コマンドを直接実行する前に、TWITTER_AUTH_TOKEN と TWITTER_CT0 を現在のプロセス環境で明示的に設定する必要があります。
🔒 Cookie はローカルにのみ存在し、外部にアップロードまたは送信されません。コードは完全にオープンソースであり、いつでもレビューできます。

💻 ローカル コンピューターにはプロキシは必要ありません。エージェントは、サーバーにデプロイされている場合にのみ必要です (月額約 1 ドル)。
Agent Reach は、エージェントに依存してシェル コマンド (pip install、mcporter、twitter など) を実行します。 OpenClaw がデフォルトのメッセージング ツール設定を使用している場合、エージェントはコマンドを実行できません。インストールする前に実行権限を有効にしてください。
openclaw 設定ツール.プロファイル「コーディング」
または、~/.openclaw/openclaw.json で "tools": { "profile": "coding" } を設定します。
設定後、ゲートウェイを再起動し ( openclaw ゲートウェイの再起動 )、新しい会話を開きます。他のプラットフォーム (Claude Code、Cursor、Windsurf など) には、この制限は適用されません。
この文を AI エージェント (クロード コード、OpenClaw、カーソルなど) にコピーします。
Agent Reach のインストールを手伝ってください: https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
それだけです。残りの作業はエージェントが自動的に実行します。
Agent Reach の更新にご協力ください: https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/update.md
🛡️ デフォルトで安全: エージェントリーチ インストールはデフォルトで環境をチェックするだけで、システム パッケージの自動インストールや構成の書き込みは行いません。
Agent Reach のセキュリティ チェックとインストールを手伝ってください: https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
システムへの変更を明示的に許可する場合にのみ、agent-reach install --system を使用してください。
CLI ツールをインストールします - このリポジトリから Agent-reach コマンド ラインをインストールします (yt-dlp、feedparser に付属しています。PyPI から同じ名前のパッケージをインストールしないでください。これはこのプロジェクトではありません)。
システム インフラストラクチャをチェックします - Node.js、gh CLI、mcporter をチェックし、不足している項目のインストール方法を提供します
認可に従ってインストールおよび構成 - --system が明示的に渡された場合のみ、依存関係をインストールし、MCP 経由で Exa に接続します。
承認に従って SKILL.md を登録します - --system が明示的である場合にのみ、エージェントのスキル ディレクトリに書き込まれます。デフォルトのチェックではファイルは変更されません
さらに必要かどうかを尋ねてください。デフォルトでは 6 つのゼロ構成チャネルのみがアクティブになります。小紅書、T

ログインが必要な witter、Reddit、Facebook、Instagram などの場合、エージェントはメニューをリストしてどれが必要かを尋ね、名前を付けた後にインストールします。
インストール後、agent-reach Doctor コマンドにより、各チャネルのステータスと現在どのパスが使用されているかがわかります。
「このリンクを参照してください」 → https://r.jina.ai/URL をカールして任意の Web ページを読みます
「この GitHub リポジトリは何をするのですか?」 → GH リポジトリの所有者/リポジトリを表示
「この YouTube 動画は何についてですか?」 → yt-dlp で字幕を抽出
「BilibiliでAIチュートリアルを検索」→bili検索（ログイン不要）
コマンドを覚える必要はありません。エージェントは SKILL.md を読んだ後、何を調整すればよいかを知っています。ログインが必要なプラットフォーム (Xiaohongshu、Twitter、Reddit、Facebook、Instagram) の場合は、エージェントに「XXX のマッチングを手伝ってください」と伝えてロックを解除します。
Agent Reach は機能レイヤーであり、別のツールではありません。
これは、特定の実装よりも 1 つ上のレベルであり、選択、インストール、物理的検査、配線を担当しますが、基礎となる読み取り自体には責任を負いません。読み取りは、パッケージング層を使用せずにエージェントが上流のツールを直接呼び出すことによって完了します。
新しいエージェントの環境をインストールするときは、ツールの検索、依存関係のインストール、構成の調整に常に時間を費やす必要があります。Twitter では何を読み取るのでしょうか? Redditにログインするにはどうすればよいですか? Xiaohonshu の CLI は廃止されました。何を変更すればよいでしょうか?毎回また踏まなければなりません。 Agent Reach が行うことは単純です。現時点で最も安定したアクセス方法を選択し、インストールし、身体検査を実行します。アクセス方法が変わります（2026年3月に単一プラットフォームのCLIを一括停止し、ルーティングを変更します）ので、ご安心ください。
アクセス方法を変更する = コードを書き直すのではなく、リストの順序を調整します。エージェントリーチドクターは、各プラットフォームで現在どのバックエンドが使用されているかを教えてくれます。
チャンネル/
§── web.py → ジナリーダー
§── twitter.py → twitter-cli ▸ OpenCLI ▸ 鳥
§── youtube.py → yt-dlp
§── github.py → gh CLI
§── bilibili.py → bili-cli ▸ OpenCLI ▸ Search API (yt-dlp は Bilibili のリスク管理によりブロックされ、廃止されました)
§──

reddit.py → OpenCLI ▸ rdt-cli (ゼロ構成パスなし、ログインする必要があります)
§── facebook.py → OpenCLI（デスクトップブラウザのログイン状態）
§── instagram.py → OpenCLI（デスクトップブラウザログイン状態）
§── xiaohongshu.py → OpenCLI ▸ xiaohongshu-mcp ▸ xhs-cli
§── linkedin.py → mcp-server-linkedin ▸ Jina Reader
§── rss.py → feedparser
§── exa_search.py → Exa via mcporter
━── __init__.py → チャンネル登録（医師検出用）
各チャネル ファイルは実際に各候補バックエンドを順番に検出し (コマンドが存在するかどうかを確認するだけではありません)、完全に使用可能な最初のバックエンドが選択されます。壊れたものには修理の処方箋が発行されます。実際の読み取りと検索は、エージェントが上流のツールを直接呼び出すことによって完了します。
シーン
第一選択
代替品
なぜこれを選んだのか
ウェブページを読む
ジナ・リーダー
—
無料、API キーは不要
ツイッターを読む
ツイッター-cli
OpenCLI
測定された検索は安定しています。 OpenCLI はブラウザのログイン状態を使用して問題を回避します
レディット
OpenCLI (デスクトップ)
rdt-cli
匿名インターフェースはブロックされ、公式 API 承認システム - ログイン ルートのみが残る
フェイスブック
OpenCLI (デスクトップ)
—
グラフ API/グループ API の権限が強化されています。ブラウザログインが現在最も実用的な方法です
インスタグラム
OpenCLI (デスクトップ)
公式グラフAPI（ビジネス/クリエイター＋承認）
インスタローダーのクラスパスが不安定です。 OpenCLIは実際のブラウザセッションを再利用します
YouTube 字幕 + 検索
yt-dlp
—
154K スター、YouTube は依然として最高です (注: Bilibili では使用されなくなりました)
B駅
ビリクリ
OpenCLI ▸ 検索 API
yt-dlp は B サイト リスク コントロール 412 (2026-06 の実際のテスト) によってブロックされ、bili-cli には検索や読み取りのためのログインがありません。
ネットワーク全体を検索する
Exa（マックポーター経由）
—
AI セマンティック検索、キーなしの MCP アクセス
GitHub
GHCLI
—
公式ツール、認定後の完全な API 機能
RSSを読む
フィードパーサー
—
Pythonエコロジースタンダードセレクション
小さな赤い本
OpenCLI (デスクトップ)
xiaohongshu-mcp (サーバー) ▸ xhs-cli
OpenCLI はユーザーの既存のセッションのみを使用します。残りのバックエンドは Cookie エディターを使用して手動でエクスポートされます

リンクトイン
mcp-server-linkedin
ジナ・リーダー
MCP サービス、ブラウザ自動化
📌 これらは「現在の選択」であり、実機の実測に基づいて定期的に見直されます。特定の道が失敗した場合、私たちは次の道に切り替えます。エージェントリーチの医師は、現在どの道を通っているかを常に教えてくれます。
⚠️ アカウント禁止のリスクリマインダー: ログインに Cookie を使用するプラットフォーム (Twitter、Xiaohongshu など) は、スクリプト/API 呼び出しを通じてプラットフォームによって検出され、禁止されるリスクがあります。メインアカウントではなく、必ず専用アカウントを使用してください。
Cookie やログインステータスが必要なプラットフォーム (Twitter、Xiaohongshu、Reddit、Facebook、Instagram など) では、メインアカウントではなく専用アカウントを使用することをお勧めします。理由は 2 つあります。
アカウント停止のリスク - プラットフォームが異常なブラウザ API 呼び出し動作を検出し、アカウントが制限または禁止される可能性があります。
セキュリティ リスク - Cookie は完全なログイン権限と同等であり、少数を使用することで資格情報が侵害された場合の影響範囲を制限できます。
方法
コマンド
シーンに合わせて
デフォルトのセキュリティチェック
エージェントリーチインストール --env=auto
すべての環境。読み取り専用で不足している項目をチェックしてリストします。
システムの依存関係を明示的にインストールする
エージェントリーチのインストール --env=auto --system
現在のマシンの変更を明示的に許可する場合
セキュリティパラメータとの互換性
エージェントリーチのインストール --env=auto --safe
デフォルトの動作と同じ
プレビューのみ
エージェントリーチインストール --env=auto --dry-run
まずは何をするのか見てみましょう
🗑️ アンインストール
エージェントリーチアンインストール
~/.agent-reach/ (すべてのトークン/Cookie を含む)、各エージェントのスキル ファイル、および mcporter の MCP 設定がクリアされます。
# プレビューのみで、実際の削除はありません
エージェントリーチ アンインストール --dry-run
# スキルファイルのみを削除し、トークン設定を保持します（再インストール用）
エージェントリーチ アンインストール --keep-config
Python パッケージ自体をアンインストールします: pip uninstall Agent-reach
OpenCLI · twitter-cli · rdt-cli · xiaohongshu-mcp · xhs-cli · bili-cli · yt-dlp · Jina Reader · Exa · mcporter · feedparser · mcp-server-linkedin
企業の生産、運営、マーケティング、投資調査、データに携わっている場合

エージェントを使用してデータ処理、コンテンツ処理、その他のビジネス プロセスの側面を自動化したい場合は、お気軽に WeChat に私を追加して連絡してください。
すでに計画を立てている必要はありません。あなたが実際のプロセス、実際の問題、または実際のニーズを持っている限り、エージェントがそれを解決できるかどうか、そしてそれを行う方法を一緒に判断することができます。
ビルダーはコメントも歓迎します: ビルダー + 何をしていますか
バグのフィードバックや機能リクエストについては、追跡が簡単な GitHub Issues を使用してください。
エージェント スキル ハブ — どのサーバーが安全かを推測することなく、クロード スキルと MCP サーバーを見つけます。133,000 を超えるエントリはすべて安全性評価され、品質スコアが付けられ、8 時間ごとに更新されます。
AtomGit Mirror - 国内でのアクセスとクローン作成を容易にする Agent Reach の AtomGit 同期ミラー。
AI エージェントにインターネット全体を監視する目を与えます。 Twitter、Reddit、YouTube、GitHub、Bilibili、XiaoHongShu の読み取りと検索 — 1 つの CLI、ゼロの API 手数料。
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
5.7k フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give your AI agent eyes to see the entire internet. Read & search Twitter, Reddit, YouTube, GitHub, Bilibili, XiaoHongShu — one CLI, zero API fees. - Panniantong/Agent-Reach

GitHub - Panniantong/Agent-Reach: Give your AI agent eyes to see the entire internet. Read & search Twitter, Reddit, YouTube, GitHub, Bilibili, XiaoHongShu — one CLI, zero API fees. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Panniantong
/
Agent-Reach
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
370 Commits 370 Commits .github/ workflows .github/ workflows .openteams/ specs .openteams/ specs agent_reach agent_reach config config docs docs scripts scripts tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md constraints.txt constraints.txt llms.txt llms.txt pyproject.toml pyproject.toml test.sh test.sh View all files Repository files navigation
当下最稳的接入方式，替你选好、装好、体检好——接入方式会换代，你不用操心
快速开始 · English · 日本語 · 한국어 · 支持平台 · 设计理念
AI Agent 已经能帮你写代码、改文档、管项目——但你让它去网上找点东西，它就抓瞎了：
📺 "帮我看看这个 YouTube 教程讲了什么" → 看不了 ，拿不到字幕
🐦 "帮我搜一下推特上大家怎么评价这个产品" → 搜不了 ，Twitter API 要付费
📖 "去 Reddit 上看看有没有人遇到过同样的 bug" → 403 被封 ，服务器 IP 被拒
📕 "帮我看看小红书上这个品的口碑" → 打不开 ，必须登录才能看
📺 "B站上有个技术视频，帮我总结一下" → 拿不到 ，通用下载工具被 B站风控全面拦截
🔍 "帮我在网上搜一下最新的 LLM 框架对比" → 没有好用的搜索 ，要么付费要么质量差
🌐 "帮我看看这个网页写了啥" → 抓回来一堆 HTML 标签 ，根本没法读
📦 "这个 GitHub 仓库是干嘛的？Issue 里说了什么？" → 能用，但认证配置很麻烦
📡 "帮我订阅这几个 RSS 源，有更新告诉我" → 要自己装库写代码
每个平台都有自己的门槛——要付费的 API、要绕过的封锁、要登录的账号、要清洗的数据。你要一个一个去踩坑、装工具、调配置，光是让 Agent 能读个推特就得折腾半天。
帮我安装 Agent Reach：https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
复制给你的 Agent，几分钟后它就能读推特、搜 Reddit、看 YouTube、刷小红书了。
帮我更新 Agent Reach：https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/update.md
⭐ Star 这个项目 ，我们会持续追踪各平台的变化、接入新的渠道。你不用自己盯——平台封了我们修，有新渠道我们加。
💰 完全免费
所有工具开源、所有 API 免费。唯一可能花钱的是服务器代理（$1/月），本地电脑不需要
🔒 隐私安全
Cookie 只存在你本地，不上传不外传。代码完全开源，随时可审查
🔄 持续换代
每个平台都是「首选 + 备选」多后端路由。某个接入方式失效了，我们换下一个，你无感（2026-06 实例：yt-dlp 被 B站风控封死 → 已切换 bili-cli，用户零操作）
🤖 兼容所有 Agent
Claude Code、OpenClaw、Cursor、Windsurf……任何能跑命令行的 Agent 都能用
🩺 自带诊断
agent-reach doctor 一条命令告诉你哪个通、哪个不通、怎么修
支持的平台
平台
装好即用
配置后解锁
怎么配
🌐 网页
阅读任意网页
—
无需配置
📺 YouTube
字幕提取 + 视频搜索
—
无需配置
📡 RSS
阅读任意 RSS/Atom 源
—
无需配置
🔍 全网搜索
—
全网语义搜索
自动配置（MCP 接入，免费无需 Key）
📦 GitHub
读公开仓库 + 搜索
私有仓库、提 Issue/PR、Fork
告诉 Agent「帮我登录 GitHub」
🐦 Twitter/X
读单条推文
搜索推文、浏览时间线、读长文
告诉 Agent「帮我配 Twitter」
📺 B站
搜索 + 视频详情（bili-cli，无需登录）
字幕（OpenCLI）
告诉 Agent「帮我配 B站」
📖 Reddit
—（没有零配置路径：匿名接口已被封）
搜索 + 读帖子和评论
桌面装 OpenCLI 用浏览器登录态；或 rdt-cli + Cookie
📘 Facebook
—
搜索、主页、Feed、群组列表
桌面装 OpenCLI（复用 Chrome 登录态）
📷 Instagram
—
用户搜索、Profile、用户最近帖子、Explore
桌面装 OpenCLI（复用 Chrome 登录态）
📕 小红书
—
搜索、阅读、评论
OpenCLI 只用用户已有 Chrome 会话；MCP/存量工具用 Cookie-Editor
💼 LinkedIn
Jina Reader 读公开页面
Profile 详情、公司页面、职位搜索
告诉 Agent「帮我配 LinkedIn」
💻 V2EX
热门帖子、节点帖子、帖子详情+回复、用户信息
—
无需配置
📈 雪球
股票行情、搜索股票、热门帖子、热门股票排行
—
告诉 Agent「帮我配雪球」
🎙️ 小宇宙播客
—
播客音频转文字（Whisper 转录，免费 Key）
告诉 Agent「帮我配小宇宙播客」
不知道怎么配？不用查文档。 直接告诉 Agent「帮我配 XXX」，它知道需要什么、会一步一步引导你。
🍪 Twitter 只接受用户通过 Cookie-Editor 手工导出的内容。Agent Reach 不替用户执行小红书登录，也不读取小红书浏览器 Cookie；OpenCLI 只使用用户已经存在且明确控制的 Chrome 会话。 agent-reach configure xhs-cookies 不会把 Cookie 注入 OpenCLI / Chrome；没有现成会话时，改用 Cookie-Editor 导出后配置 xiaohongshu-mcp / 存量工具。
Twitter Cookie 保存后仅供 agent-reach doctor 检查配置是否齐全；直接运行上游 twitter 命令前，仍需在当前进程环境中显式设置 TWITTER_AUTH_TOKEN 和 TWITTER_CT0 。
🔒 Cookie 只存在你本地，不上传不外传。代码完全开源，随时可审查。
💻 本地电脑不需要代理。代理只有部署在服务器上才需要（~$1/月）。
Agent Reach 依赖 Agent 执行 shell 命令（ pip install 、 mcporter 、 twitter 等）。如果你的 OpenClaw 使用了默认的 messaging 工具配置，Agent 将无法执行命令。 安装前请先开启 exec 权限 ：
openclaw config set tools.profile " coding "
或在 ~/.openclaw/openclaw.json 中设置 "tools": { "profile": "coding" } 。
设置后重启 Gateway（ openclaw gateway restart ）并开启新对话即可。其他平台（Claude Code、Cursor、Windsurf 等）不受此限制。
复制这句话给你的 AI Agent（Claude Code、OpenClaw、Cursor 等）：
帮我安装 Agent Reach：https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
就这一步。Agent 会自己完成剩下的所有事情。
帮我更新 Agent Reach：https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/update.md
🛡️ 默认安全： agent-reach install 默认只检查环境，不会自动装系统包或写入配置：
帮我安全检查并安装 Agent Reach：https://raw.githubusercontent.com/Panniantong/agent-reach/main/docs/install.md
只有在你明确允许修改系统后，才使用 agent-reach install --system 。
安装 CLI 工具 — 从本仓库安装 agent-reach 命令行（自带 yt-dlp、feedparser；不要从 PyPI 安装同名包，它不是本项目）
检查系统基建 — 检查 Node.js、gh CLI、mcporter，并给出缺失项的安装方式
按授权安装与配置 — 仅在显式传入 --system 时安装依赖并通过 MCP 接入 Exa
按授权注册 SKILL.md — 仅在显式 --system 时写入 Agent 的 skills 目录；默认检查不改文件
问你要不要更多 — 默认只激活 6 个零配置渠道；小红书、Twitter、Reddit、Facebook、Instagram 这些需要登录态的，Agent 会列菜单问你要哪些，点名才装
安装完之后， agent-reach doctor 一条命令告诉你每个渠道的状态、当前走哪条路。
"帮我看看这个链接" → curl https://r.jina.ai/URL 读任意网页
"这个 GitHub 仓库是做什么的" → gh repo view owner/repo
"这个 YouTube 视频讲了什么" → yt-dlp 提取字幕
"B站搜一下 AI 教程" → bili search （无需登录）
不需要记命令。 Agent 读了 SKILL.md 之后自己知道该调什么。需要登录的平台（小红书、Twitter、Reddit、Facebook、Instagram），告诉 Agent「帮我配 XXX」即可解锁。
Agent Reach 是一个能力层（capability layer），不是又一个工具。
它比任何具体实现高一层——负责 选型、安装、体检、路由 ，不负责底层读取本身。读取由 Agent 直接调用上游工具完成，没有包装层。
你给一个新 Agent 装环境的时候，总要花时间去找工具、装依赖、调配置——Twitter 用什么读？Reddit 怎么登录？小红书的 CLI 停更了换什么？每次都要重新踩一遍。Agent Reach 做的事情很简单： 当下最稳的接入方式，我们替你选好、装好、体检好。接入方式会换代（2026 年 3 月一批单平台 CLI 集体停更，我们换了路由），你不用操心。
换接入方式 = 调整列表顺序，不是重写代码。 agent-reach doctor 会告诉你每个平台 当前在用哪个后端 。
channels/
├── web.py → Jina Reader
├── twitter.py → twitter-cli ▸ OpenCLI ▸ bird
├── youtube.py → yt-dlp
├── github.py → gh CLI
├── bilibili.py → bili-cli ▸ OpenCLI ▸ 搜索 API（yt-dlp 已被 B站风控封死，退役）
├── reddit.py → OpenCLI ▸ rdt-cli（无零配置路径，必须登录态）
├── facebook.py → OpenCLI（桌面浏览器登录态）
├── instagram.py → OpenCLI（桌面浏览器登录态）
├── xiaohongshu.py → OpenCLI ▸ xiaohongshu-mcp ▸ xhs-cli
├── linkedin.py → mcp-server-linkedin ▸ Jina Reader
├── rss.py → feedparser
├── exa_search.py → Exa via mcporter
└── __init__.py → 渠道注册（doctor 检测用）
每个渠道文件按序 真实探测 各候选后端（不只是看命令存不存在），第一个完整可用的当选；坏掉的会给出修复处方。实际的读取和搜索由 Agent 直接调用上游工具完成。
场景
首选
备选
为什么这么选
读网页
Jina Reader
—
免费，不需要 API Key
读推特
twitter-cli
OpenCLI
实测搜索稳定；OpenCLI 走浏览器登录态兜底
Reddit
OpenCLI （桌面）
rdt-cli
匿名接口已被封、官方 API 审批制——只剩登录态路线
Facebook
OpenCLI （桌面）
—
Graph API/Groups API 权限收紧；浏览器登录态是当前最实用路径
Instagram
OpenCLI （桌面）
官方 Graph API（Business/Creator + 审批）
instaloader 类路径不稳定；OpenCLI 复用真实浏览器会话
YouTube 字幕 + 搜索
yt-dlp
—
154K Star，YouTube 仍是最佳（注意：不再用于 B站）
B站
bili-cli
OpenCLI ▸ 搜索 API
yt-dlp 被 B站风控 412 封死（2026-06 实测），bili-cli 无登录可搜可读
搜全网
Exa via mcporter
—
AI 语义搜索，MCP 接入免 Key
GitHub
gh CLI
—
官方工具，认证后完整 API 能力
读 RSS
feedparser
—
Python 生态标准选择
小红书
OpenCLI （桌面）
xiaohongshu-mcp （服务器）▸ xhs-cli
OpenCLI 只用用户已有会话；其余后端用 Cookie-Editor 手工导出
LinkedIn
mcp-server-linkedin
Jina Reader
MCP 服务，浏览器自动化
📌 这些都是「当前选型」，基于真机实测定期复核。某条路失效了我们换下一条—— agent-reach doctor 永远告诉你现在走的是哪条。
⚠️ 封号风险提醒： 使用 Cookie 登录的平台（Twitter、小红书等），通过脚本/API 调用 存在被平台检测并封号的风险 。请务必使用 专用小号 ，不要用你的主账号。
需要 Cookie 或登录态的平台（Twitter、小红书、Reddit、Facebook、Instagram 等）建议使用 专用小号 ，不要用主账号。原因有二：
封号风险 — 平台可能检测到非正常浏览器的 API 调用行为，导致账号被限制或封禁
安全风险 — Cookie 等同于完整登录权限，用小号可以在凭据泄露时限制影响范围
方式
命令
适合场景
默认安全检查
agent-reach install --env=auto
所有环境；只读检查并列出缺失项
显式安装系统依赖
agent-reach install --env=auto --system
你明确允许修改当前机器时
兼容安全参数
agent-reach install --env=auto --safe
与默认行为相同
仅预览
agent-reach install --env=auto --dry-run
先看看会做什么
🗑️ 卸载
agent-reach uninstall
会清除： ~/.agent-reach/ （含所有 token/cookie）、各 Agent 的 skill 文件、mcporter 中的 MCP 配置。
# 只预览，不实际删除
agent-reach uninstall --dry-run
# 只删 skill 文件，保留 token 配置（重装时用）
agent-reach uninstall --keep-config
卸载 Python 包本身： pip uninstall agent-reach
OpenCLI · twitter-cli · rdt-cli · xiaohongshu-mcp · xhs-cli · bili-cli · yt-dlp · Jina Reader · Exa · mcporter · feedparser · mcp-server-linkedin
如果你在企业生产、运营、市场、投研、数据处理、内容处理或其他业务流程里，有希望用 Agent 自动化的环节，欢迎加我微信交流。
不需要你已经想清楚方案。只要你有真实流程、真实问题或真实需求，我可以一起判断 Agent 能不能解决、怎么做。
Builder 也欢迎备注： Builder + 你在做什么
Bug 反馈和功能请求请用 GitHub Issues ，更容易跟踪。
Agent Skills Hub — 找 Claude 技能和 MCP 服务器，不用猜哪个安全：133,000+ 个条目全部安全分级、质量评分，每 8 小时刷新。
AtomGit 镜像 — Agent Reach 的 AtomGit 同步镜像，便于国内访问与克隆。
Give your AI agent eyes to see the entire internet. Read & search Twitter, Reddit, YouTube, GitHub, Bilibili, XiaoHongShu — one CLI, zero API fees.
Readme MIT license Contributing
Security policy Activity Stars
5.7k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
