---
source: "https://proandroiddev.com/seven-claude-code-commands-one-kotlin-multiplatform-app-on-android-and-ios-cb01e920a3e6"
hn_url: "https://news.ycombinator.com/item?id=49097777"
title: "An Android and iOS app from 7 Claude Code commands, every prompt and timing"
article_title: "Medium"
author: "thisissadeghi"
captured_at: "2026-07-29T15:06:55Z"
capture_tool: "hn-digest"
hn_id: 49097777
score: 1
comments: 0
posted_at: "2026-07-29T14:08:57Z"
tags:
  - hacker-news
  - translated
---

# An Android and iOS app from 7 Claude Code commands, every prompt and timing

- HN: [49097777](https://news.ycombinator.com/item?id=49097777)
- Source: [proandroiddev.com](https://proandroiddev.com/seven-claude-code-commands-one-kotlin-multiplatform-app-on-android-and-ios-cb01e920a3e6)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T14:08:57Z

## Translation

タイトル: 7 つのクロード コード コマンド、すべてのプロンプトとタイミングからの Android および iOS アプリ
記事タイトル: 中
説明: 2 つの機能、2 つのプラットフォーム、1 つの Kotlin マルチプラットフォーム コードベース。クロード コードと、アーキテクチャの漂流を防ぐテンプレートが必要です。この投稿はビルド全体であり、何もカットされていません…

記事本文:
中 KMPilot の完全なウォークスルー: 7 つのクロード コード コマンドによる Android + iOS アプリ (すべてのプロンプトとタイミングが含まれています) |アリ・サデギ著 | 2026 年 7 月 | ProAndroidDev サイトマップ アプリで開く サインアップ
セットアップ: 1 つのコマンド、1 つの MCP サーバー
パイプライン: 7 つのコマンド、ステップ 8 なし
各コマンドの間には次の 1 つの点があります: /clear
ラウンド 1: ギャラリー、最初にデザイン
/create-feature ギャラリー (58:41)
/レビュー特集ギャラリー (6:11)
/フィーチャーアーティストの作成 (25:48)
/modify-feature アーティスト (13:08)
KMPilot が行ったことは、単純な Claude コードにはできなかったことです
Android プロフェッショナルと Google 開発者エキスパートからの最新の投稿。
KMPilot の完全なウォークスルー: 7 つのクロード コード コマンドによる Android + iOS アプリ (すべてのプロンプトとタイミングが含まれています)
2 つの機能、2 つのプラットフォーム、1 つの Kotlin マルチプラットフォーム コードベース。クロード コードと、アーキテクチャの漂流を防ぐテンプレートが必要です。
ギャラリー、アートワークの詳細、アーティスト、空の「お気に入り」タブなど、両方のプラットフォームで同じビルドが構築されます。 1 つの Kotlin マルチプラットフォーム コードベース。この投稿は、シカゴ美術館のパブリック ドメイン コレクションのアート ギャラリーである Loupe の記録から何もカットされていない全体の構築です。私が入力したすべてのプロンプトはこの投稿に含まれており、すべてのタイミングはアクティブな作業のみをカウントします。
コマンドは、クロード コードをデザインファーストの Kotlin マルチプラットフォーム パイプラインに変える MIT ライセンスのテンプレートである KMPilot から取得されます。最初の記事では、なぜそれが存在するのかについて説明しました (AI が生成したコードは、プロンプト間でその構造を何かが保持しない限り、ドリフトします)。これは HOW であり、エンドツーエンドで再現できます。あなたの仕事は、各機能を説明し、チェックポイントを指示し、承認することです。同じパイプラインで、はるかに大規模なワールドカップ コンパニオン アプリである Kickoff26 が構築されました。ルーペは一度で追えるよう意図的に小さく作られています。
フルセッション、スピードアップしてチャ

コマンドによって操作されます。ここで全体を視聴するか、読み続けて以下のセクションのリンクから任意のフェーズにジャンプしてください。
3 つのタブ。すべてシカゴ美術館の API によってサポートされています。ギャラリーは美術館の作品を閲覧および検索し、フルブリードの詳細画面を開きます。アーティストはアーティストとその作品を閲覧します。お気に入りはウォークスルー全体で空のプレースホルダーとして残り、最後には演習になります。
セットアップ: 1 つのコマンド、1 つの MCP サーバー
前提条件: クロード コード、JDK 21+、PATH 上の python3 (設計ステップではトークン抽出に使用します)。次に、テンプレートをインストールします。
カール -fsSL https://github.com/ThisIsSadeghi/KMPilot/releases/download/v0.1.3/install.sh | bash インストーラーは、アプリ名とパッケージのプレフィックスという 2 つの値の入力を求めます。私は Loupe と com.yourname.loupe に答えました。これらは単なる私の値です。独自の値を使用してください。以下のすべてのコマンドは同じように動作します。
次に、Claude Code でプロジェクトを開きます。
cd Loupe && claude 完全な Kotlin Multiplatform スケルトンを取得します。Gradle ワイヤリング、Compose Multiplatform を介した共有 UI、デザイン システム モジュール、7 つのコマンド、フック、エージェント定義です。クローンと名前変更なので数秒で終わります。
このウォークスルーは v0.1.3 に固定されているため、読んだ内容がそのまま実行されます。 wiki は現バージョンです。 KMPilot がコードを作成します。書き込んだものをビルドして実行するには、通常の Kotlin マルチプラットフォーム ツール、Android ターゲットの場合は Android Studio、iOS の場合は Xcode を備えた Mac を使用します。
1 つの 1 回限りのセットアップ: /design-ui は、MCP サーバーを通じて Google の AI デザイン ツールである Google Stitch を駆動します。 stack.withgoogle.com/settings から API キーを取得し (サインインし、[API キー] セクションを見つけて、[API キーの作成] をクリックします)、ユーザー スコープでサーバーを登録して、すべてのプロジェクトでそれが認識されるようにします。
クロード mcp ステッチを追加 --transport http https://stitch.googleapis.com/mcp

\
--header "X-Goog-Api-Key: YOUR_API_KEY" -s user claude mcp list # →ステッチ: … ✔ 接続された 2 つのフィールド ノート。キーは Stitch 独自の設定ページから取得する必要があります。aistudio.google.com で作成されたキーは同一に見えますが、Stitch API はそれらを拒否します。そして、Stitch サーバーには厳しい日々が続いています。クロード mcp リストに接続エラーまたはスキーマ エラーが表示される場合、それはあなたの設定ではなく彼ら側にあることがほぼ確実です。セットアップをデバッグする前に、後で再試行してください。
パイプライン: 7 つのコマンド、ステップ 8 なし
パイプライン: 7 つのコマンド 各コマンドには以下のセクションがあり、実際のタイミングとともに、それが構築した実際の機能が示されています。すべてを手作業で修正するステップ 8 はありません。フックは機能コードへの直接編集をブロックするため、変更はパイプライン経由でのみ反映されます。
パイプラインは 2 回実行され、毎回異なる順序で実行されます。 gallery では、デザインが最初に来ます: /design-ui 、次に /create-feature がそのデザインに対してビルドされます。アーティストの場合、コードが最初にあり、その後デザインがボルトで固定されます。 2 つの注文、同じレール、2 回目の実行では、レールが実際に節約できる量が表示されます。
各コマンドの間には次の 1 つの点があります: /clear
すべてのコマンドは、次のコマンドの前に /clear を実行するように指示します。 KMPilot は再開を認識します。そのメモリはコンテキスト ウィンドウではなくディスク上に存在するため、ステップ間の会話全体がクリアされ、何も失われません。
ラウンド 1: ギャラリー、最初にデザイン
/design-ui <feature> を実行し、続いて画面の平易な説明を実行します。ビルドは設計ツールで開始されます。逐語的、タイプミスなどすべて:
/design-ui ギャラリーには次の詳細が含まれています。
シカゴ美術館のパブリックドメインコレクションのアートギャラリーブラウザ。
画面 1 - ギャラリー: タイトルとアーティストを含むアートワーク画像の編集グリッド、
検索バー、現代の博物館アプリの感覚。アプリには下部ナビゲーションがあります
3 つのタブのあるバー -

ギャラリー、アーティスト、お気に入り - ギャラリーが選択されています。
画面 2 - アートワークの詳細: フルブリード ヒーロー画像、タイトル、アーティスト、日付、媒体、
クレジットライン。アートに集中できる、静かでギャラリーのような美学。
両方をプレミアムかつプロフェッショナルに見せます このスキルは MCP 上で Stitch を駆動し、モックアップをセッションに戻します。そこで、デザイナーに説明するときと同じように、変更点を説明してモックアップを調整します。ユーザーに代わって黙って設計上の決定を行うわけではありません。アプリ全体を形作るものについては、立ち止まって尋ねます。初期の段階でブランドカラーを提案し、ギャラリーの外観に合わせて調整された 2 つのニュートラルなオプションを提供しました。両方を渡して、独自の 16 進数を入力しました。
このスキルはインクと石を提供します。私は両方を拒否し、独自の #6B1F2A を設定します。デザイン トークンを 1 行で操作します。また、仮定するのではなく、どのオプションの状態を設計するかについても尋ねました。私は、[読み込み中]、[空]、[失敗] の 3 つすべてを要求したため、後でオプトインするすべての機能は、共有の [読み込み中] 画面と [失敗] 画面を無料で継承します。
すべての変更がプロンプトに反映されるわけではありません。 Stitch には、サーバー側で画面が終了したにもかかわらず、接続がリセットされて生成呼び出しがタイムアウトになるという既知のバグがあります。この問題は Google の AI 開発者フォーラムで報告および確認されており、それ以来修正済みとマークされていますが、依然として発生したり消えたりしています。スキルにはパスが定義されており、再試行するのではなく（再試行すると画面が複製されてしまいます）停止し、ブラウザでプロジェクトを開いて同期するように求められます。
「接続がリセットされました (Stitch の既知のバグ)、再試行されません。」このスキルは、作業を重複させるのではなく、スティッチの不安定さを回避します。 MCP 上の編集パスも同様に誤動作していたため、アートワークの詳細画面での 1 つの変更については、それと戦うのをやめ、代わりに Stitch Web エディターで画面を直接編集しました。そこで手動で調整し、スキルにバージョンを使用するように指示しました。

プロジェクトを再生成するのではなく、すでにプロジェクトに常駐しています。それは手動で編集された画面を取得し、そのトークンを抽出して続行しました。
画面が適切であれば承認し、承認によってクロード コードに相当するものがないプレーンなステップがトリガーされます。このスキルは完成した画面を Stitch から取得し、デザイン自体からすべての色、半径、フォント、間隔の値を抽出し、モックアップが使用する正確なアイコンと画像をダウンロードして、そのすべてをブループリントに書き込みます。デザインは明示的な Compose 命令としてキャプチャされ、 .claude/docs/gallery/ に保存されます。また、カラー監査も実行され、コード内のテーマが Stitch が実際にレンダリングしたものから逸脱している箇所にフラグが立てられるため、ビルド ステップで事前に修正できます。
UIデザイナーが完成しました。 5 つの状態のスクリーンショット、ブループリント、設計仕様、カラー ドリフト フラグはすべて、Kotlin の行が書き込まれる前に表示されます。 ▶ このフェーズをセッション全体で視聴する
/create-feature ギャラリー (58:41)
/create-feature <feature> とデータ コントラクト (エンドポイント、フィールド、動作) を実行します。 7人の中で最大の命令。デザインは既にブループリントに存在しているため (スキルがそれを自動検出し、デザイン認識モードに入ります)、引き渡すのは機能の残りの半分であるデータ コントラクトです。
/create-feature gallery に次の詳細が含まれます:
Art Institute のパブリックドメインの作品を閲覧する
シカゴAPIの。検索エンドポイント:
https://api.artic.edu/api/v1/artworks/search?query[term][is_public_domain]=true&fields=id,title,artist_display,image_id,date_display
(認証なし。User-Agent ヘッダーを送信します)。検索機能付きギャラリーグリッド。アートワークをタップすると
IIIF 経由でヒーロー画像を含む詳細画面を開きます
(https://www.artic.edu/iiif/2/{image_id}/full/843,/0/default.jpg)、さらに
ミディアムディスプレイとクレジットライン。の設計に従って下部ナビゲーションを実装します。
ブループリント - アーティストと

現時点では、お気に入りタブはプレースホルダー画面です。普通のクロード コードはここでファイルの書き込みを開始します。 KMPilot は最初に 2 回停止します。
チェックポイント 1: PRD。数分後、短い製品要件ドキュメント (機能の内容、画面、データ、エッジ ケース) が画面に表示され、待機中のプロンプトが表示されました。ここでステアリングはほぼフリーになります。私の唯一の介入は 1 つの修正でした。プロジェクトはベース URL を composeApp/build.gradle.kts に保持しているため、データ層は独自にハードコーディングするのではなく、そこから URL を読み取る必要があることを指摘しました。スキルはそれを明示的な設計決定として PRD に組み込み、BASE_URL をその Gradle ファイルに移動し、タスクを生成する前に承認を待ちました。
コードが存在する前の PRD チェックポイント。ベース URL を 1 行修正しました。チェックポイント 2: タスク。承認された PRD は個別のタスク (データ レイヤー、UI、配線) に分割され、スキルは再び待機します。ここでは実装の詳細を 1 つリダイレクトしました。生成されたプランは、Koin のparametersOf を通じて詳細 ViewModel の引数を結び付けました。代わりに SavedStateHandle を使用するように依頼しました。これにより、アートワーク ID は構築時に渡されるのではなく、ナビゲーション ルートから取得されます。
3 人のエージェントにわたる 13 のタスク。実装を開始する前に、詳細 ViewModel を 1 行で SavedStateHandle にリダイレクトします。影響を受けるタスクが更新され、変更が私に返されることが確認されました。私は並列実行の計画を承認しました。それが、この機能が存在する前に私が入力した最後の内容でした。データ (モデル、リポジトリ、ネットワーク呼び出し)、UI (ViewModel、Compose 画面)、および統合 (DI、ナビゲーション、Gradle ワイヤリング) のタスクは、並行して実行されているエージェントに送られます。それぞれが 1 つのレイヤーを所有するため、衝突することはありません。その結果、同じクリーン アーキテクチャ形状の完全な有線機能モジュールが得られます。

ery KMPilot 機能が備えています。
あなたが承認するチェックポイントは、あなたが解く差分よりも安価です。
モジュールとともに、スキルは .claude/docs/gallery/spec.md に spec.md を書き込み、バージョン管理され、プロジェクトにコミットされます。このファイルは機能のメモリであり、自由形式ではなく構造化されています。つまり、目標と目標以外、拒否された代替案を含む設計決定の表、GIVEN/WHEN/THEN シナリオとしての要件、および日付付きの変更ログです。以下のすべてのコマンドは、実行する前にそれを読み取り、いくつかのコマンドはそれを書き戻します。
機能の spec.md、KMPilot 専用アーティファクト。 Enter キーを押すかクリックすると、フルサイズの画像が表示されます。機能が完了しました。データ、UI、および統合レイヤーが接続され、ビルドが渡され、spec.md が書き込まれ、一時的なタスク ファイルがクリーンアップされました。 ▶ このフェーズをセッション全体で視聴する
/verify-ui <feature> を実行します。機能名のみで、概要はありません。設計ステップではトークンを抽出しました。このステップにより、構築された画面が保持されます。 Plain Claude Code は、コンパイル時に完了画面を呼び出して次に進みます。 /verify-ui gallery は、実装された Compose を再度開き、ブループリントに対してトークンごとに再チェックします。私の実行では、初めて不合格の結果が返されました。重大な不一致が 7 つ、軽微な不一致が 7 つあり、それぞれウィジェットに名前が付けられています。
UI が完了していることを確認します。トークン

[切り捨てられた]

## Original Extract

Two features, two platforms, one Kotlin Multiplatform codebase. You need Claude Code and a template that keeps the architecture from drifting. This post is the whole build with nothing cut from the…

Medium The full KMPilot walkthrough: an Android + iOS app from seven Claude Code commands, every prompt and timing included | by Ali Sadeghi | Jul, 2026 | ProAndroidDev Sitemap Open in app Sign up
Setup: one command, one MCP server
The pipeline: seven commands, no step eight
One thing between every command: /clear
Round 1: gallery, design first
/create-feature gallery (58:41)
/review-feature gallery (6:11)
/create-feature artists (25:48)
/modify-feature artists (13:08)
What KMPilot did that plain Claude Code doesn’t
The latest posts from Android Professionals and Google Developer Experts.
The full KMPilot walkthrough: an Android + iOS app from seven Claude Code commands, every prompt and timing included
Two features, two platforms, one Kotlin Multiplatform codebase. You need Claude Code and a template that keeps the architecture from drifting.
The same build on both platforms: gallery, artwork detail, artists, and the empty Favorites tab. One Kotlin Multiplatform codebase. T his post is the whole build with nothing cut from the record: Loupe , an art gallery for the Art Institute of Chicago’s public-domain collection. Every prompt I typed is in this post, and every timing counts active work only.
The commands come from KMPilot , an MIT-licensed template that turns Claude Code into a design-first Kotlin Multiplatform pipeline. The first article covered WHY it exists (AI-generated code drifts unless something holds its structure in place between prompts). This one is the HOW, and you can reproduce it end to end. Your job is to describe each feature, steer at the checkpoints, and approve. The same pipeline built Kickoff26 , a much larger World Cup companion app; Loupe is deliberately small enough to follow in one sitting.
The full session, sped up and chaptered by command. Watch the whole thing here, or read on and jump to any phase from its section link below.
Three tabs, all backed by the Art Institute of Chicago’s API . Gallery browses and searches the museum’s artworks and opens a full-bleed detail screen. Artists browses artists and their works. Favorites stays an empty placeholder for the whole walkthrough and becomes your exercise at the end.
Setup: one command, one MCP server
Prerequisites: Claude Code , JDK 21+, python3 on PATH (the design step uses it for token extraction). Then install the template:
curl -fsSL https://github.com/ThisIsSadeghi/KMPilot/releases/download/v0.1.3/install.sh | bash The installer prompts you for two values: the app name and the package prefix. I answered Loupe and com.yourname.loupe, they're just my values, use your own and every command below behaves the same.
Then open the project in Claude Code:
cd Loupe && claude You get a complete Kotlin Multiplatform skeleton: Gradle wiring, shared UI via Compose Multiplatform, a design-system module, the seven commands, the hooks, the agent definitions. It’s a clone and a rename, so it finishes in seconds.
This walkthrough is pinned to v0.1.3 so that what you read is what you run; the wiki is the living version. KMPilot writes the code; to build and run what it writes you use the usual Kotlin Multiplatform tooling, Android Studio for the Android target and a Mac with Xcode for iOS.
One piece of one-time setup: /design-ui drives Google Stitch , Google's AI design tool, through its MCP server. Get an API key from stitch.withgoogle.com/settings (sign in, find the API keys section, click Create API key), then register the server at user scope so every project sees it:
claude mcp add stitch --transport http https://stitch.googleapis.com/mcp \
--header "X-Goog-Api-Key: YOUR_API_KEY" -s user claude mcp list # → stitch: … ✔ Connected Two field notes. The key must come from Stitch’s own settings page: keys created at aistudio.google.com look identical but the Stitch API rejects them. And the Stitch server has rough days; if claude mcp list shows connection or schema errors, it’s almost certainly their side rather than your config. Retry later before debugging your setup.
The pipeline: seven commands, no step eight
The pipeline: seven commands Each command gets a section below, shown on the real feature it built, with its real timing. There is no step eight where you fix everything by hand; a hook blocks direct edits to feature code, so changes only ever land through the pipeline.
The pipeline runs twice, and in a different order each time. On gallery , the design comes first: /design-ui , then /create-feature builds against that design. On artists , the code comes first and the design is bolted on afterwards. Two orders, same rails, and the second run shows what the rails actually save.
One thing between every command: /clear
Every command tells you to run /clear before the next one. KMPilot is resume-aware: its memory lives on disk, not in the context window, so you clear the whole conversation between steps and lose nothing.
Round 1: gallery, design first
You run /design-ui <feature> followed by a plain-language brief of the screens. The build starts in a design tool. Verbatim, typos and all:
/design-ui gallery with the following details:
An art gallery browser for the Art Institute of Chicago's public-domain collection.
Screen 1 - Gallery: an editorial grid of artwork images with title and artist,
a search bar, the feel of a modern museum app. The app has a bottom navigation
bar with three tabs - Gallery, Artists, Favorites - with Gallery selected.
Screen 2 - Artwork detail: full-bleed hero image, title, artist, date, medium,
credit line. A quiet, gallery-like aesthetic that keeps all focus on the art.
make both look premium and professional The skill drives Stitch over MCP and brings the mockups back into the session, where you refine them by describing changes, the way you’d brief a designer. It does not silently make design decisions for you; the ones that shape the whole app, it stops and asks. Early on it proposed a brand color, offering two neutral options tuned for a gallery look. I passed on both and typed my own hex:
The skill offers Ink and Stone; I reject both and set my own #6B1F2A. Steering a design token in one line. It also asked which optional states to design rather than assuming. I asked for all three: Loading, Empty, and Failed, so every feature opting in later inherits the shared Loading and Failed screens for free.
Not every change went through the prompt. Stitch has a known bug where a generation call times out with a connection reset even though the screen finished server-side; it’s reported and confirmed on Google’s AI developer forum , and marked fixed since, though it still comes and goes. The skill has a defined path for it, and rather than retry (a retry would duplicate the screen) it stops and asks me to open the project in the browser to sync.
“Connection reset (known Stitch bug), not retrying.” The skill routes around Stitch’s flakiness instead of duplicating work. The edit path over MCP was misbehaving in the same way, so for one change on the Artwork Detail screen I stopped fighting it and edited the screen directly in the Stitch web editor instead. I adjusted it there by hand, then told the skill to use the version already sitting in the project rather than regenerate it. It picked up the hand-edited screen, extracted its tokens, and carried on.
When a screen is right, you approve it, and approval triggers the step plain Claude Code has no equivalent for. The skill pulls the finished screens back from Stitch, extracts every color, radius, font, and spacing value out of the design itself, downloads the exact icons and images the mockup uses, and writes all of it into a blueprint: the design captured as explicit Compose instructions, stored under .claude/docs/gallery/ . It also runs a color audit and flags where the theme in code has drifted from what Stitch actually rendered, so the build step can correct it up front.
UI Designer Complete. Five state screenshots, a blueprint, a design spec, and a color-drift flag, all before a line of Kotlin is written. ▶ Watch this phase in the full session
/create-feature gallery (58:41)
You run /create-feature <feature> plus the data contract (endpoints, fields, behavior). The biggest command of the seven. The design already exists in the blueprint (the skill auto-detects it and enters design-aware mode), so what you hand over is the other half of a feature: the data contract.
/create-feature gallery with the following detail:
browse public-domain artworks from the Art Institute
of Chicago API. Search endpoint:
https://api.artic.edu/api/v1/artworks/search?query[term][is_public_domain]=true&fields=id,title,artist_display,image_id,date_display
(no auth; send a User-Agent header). Gallery grid with search; tapping an artwork
opens a detail screen with hero image via IIIF
(https://www.artic.edu/iiif/2/{image_id}/full/843,/0/default.jpg), plus
medium_display and credit_line. Implement the bottom navigation as designed in
the blueprint - Artists and Favorites tabs are placeholder screens for now. Plain Claude Code would start writing files here. KMPilot stops twice first.
Checkpoint one: the PRD. A few minutes in, a short product requirements document was on screen (what the feature does, its screens, its data, its edge cases) with a waiting prompt. This is where steering is nearly free. My only intervention was one correction: I pointed out that the project keeps its base URL in composeApp/build.gradle.kts , so the data layer should read it from there rather than hardcode its own. The skill folded that into the PRD as an explicit design decision, moved BASE_URL to that Gradle file, and waited for me to approve before generating any tasks.
The PRD checkpoint, with my one-line base-URL correction, before any code exists. Checkpoint two: the tasks. The approved PRD gets broken into discrete tasks (data layer, UI, wiring), and the skill waits again. Here I redirected one implementation detail. The generated plan wired the detail ViewModel’s argument through Koin’s parametersOf ; I asked it to use a SavedStateHandle instead, so the artwork id arrives from the navigation route rather than being handed in at construction time.
13 tasks across three agents; I redirect the detail ViewModel to SavedStateHandle in one line, before implementation starts. It updated the affected tasks and confirmed the change back to me. I approved the plan for parallel execution, and that was the last thing I typed before the feature existed. The tasks go to agents running in parallel: data (models, repository, the network call), ui (ViewModel, Compose screens), and integration (DI, navigation, Gradle wiring). Each owns one layer, so they never collide. The result is a complete, wired feature module in the same Clean Architecture shape every KMPilot feature has.
A checkpoint you approve is cheaper than a diff you untangle.
Alongside the module, the skill writes spec.md at .claude/docs/gallery/spec.md , versioned and committed with the project. This file is the feature's memory, and it's structured rather than freeform: goals and non-goals, a table of design decisions with the alternatives that were rejected, requirements as GIVEN/WHEN/THEN scenarios, and a dated changelog. Every command below reads it before acting, and several write it back.
A feature’s spec.md, the KMPilot-only artifact. Press enter or click to view image in full size Feature Complete. Data, UI, and integration layers wired, build passing, spec.md written, ephemeral task files cleaned. ▶ Watch this phase in the full session
You run /verify-ui <feature> , feature name only, no brief. The design step extracted tokens; this step holds the built screens to them. Plain Claude Code calls a screen done when it compiles and moves on. /verify-ui gallery re-opens the implemented Compose and re-checks it against the blueprint, token by token. On my run it came back with a failing grade the first time: seven critical mismatches and seven minor ones, each named down to the widget.
Verify UI Complete. A token a

[truncated]
