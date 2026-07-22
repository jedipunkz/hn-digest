---
source: "https://github.com/sktkkoo/Yorishiro"
hn_url: "https://news.ycombinator.com/item?id=49008434"
title: "Show HN: Yorishiro – a macOS terminal where AI agents live"
article_title: "GitHub - sktkkoo/Yorishiro: A terminal that gives AI a body and a living space. · GitHub"
author: "hakumei"
captured_at: "2026-07-22T16:16:37Z"
capture_tool: "hn-digest"
hn_id: 49008434
score: 1
comments: 0
posted_at: "2026-07-22T15:33:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Yorishiro – a macOS terminal where AI agents live

- HN: [49008434](https://news.ycombinator.com/item?id=49008434)
- Source: [github.com](https://github.com/sktkkoo/Yorishiro)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T15:33:16Z

## Translation

タイトル: Show HN: よりしろ – AI エージェントが住む macOS 端末
記事タイトル：GitHub - sktkkoo/よりしろ：AIに身体と居住空間を与える端末。 · GitHub
概要：AIに身体と居住空間を与える端末。 - sktkkoo/ヨリシロ
HN テキスト: 私は、Claude Code や Codex などのコーディング エージェントと連携するためのオープンソース macOS ターミナルである Yorishiro を構築しています。名前は日本語です。依り代とは、精霊が宿る物体のことです。これが大まかな考え方です。端末をエージェントへの窓として扱うのではなく、エージェントが存在する場所として扱います。アクティブなエージェントは 3D キャラクターとして表示されます。思考、エラー、完了などの状態は、スクロールバックに埋もれるのではなく、キャラクターとその周囲を通じて表現されます。私がそれを構築し始めたのは、コーディング エージェントとの長いセッションで疲れ果てていたからです。問題の一部は、テキストの流れを見つめ、エージェントがどのような状態にあるのかを常に推測しようとしていたことだったと思います。私がフィクションの中で見て育った AI キャラクターは、異なるものに感じられました。彼らは身体や表情で感情を表現し、自分の意志を持った存在のように感じました。これは単なる飾りではなく、私たちが AI をどのように認識し、理解するかというインターフェースの問題ではないかと考えるようになりました。最終的に気になった点は次のとおりです。 * 反応は LLM を待つべきではありません。別の Reflex Layer が生の PTY 出力を監視します。ビルドが失敗した場合、キャラクターはモデルが何かを言う前に即座に反応できます。
* 状態はキャラクター内だけでなく、環境全体に現れる必要があります。照明、シーン、エフェクトも、エージェントに何が起こっているかを伝えます。
※住民とその環境はアプリ内から編集可能である必要があります。ペルソナ、反射神経、シーン、エフェクトは JavaScr として定義されます

iptパック。自分で編集することも、エージェントに自身の外観や環境の変更を依頼することもできます。変更はセッションを中断せずにホットリロードされ、ワンクリックでロールバックできます。
現時点では macOS のみで、署名され、MIT ライセンスが付与されています。 > brew install --cask sktkkoo/yorishiro/yorishiro 正直なところ、これによって長時間のセッションが楽になるのか、それとも目新しさが薄れると単に気が散ってしまうだけなのかはまだわかりません。試してみて、どこから煩わしさを感じるか聞いてみたいです。

記事本文:
GitHub - sktkkoo/よりしろ: AIに身体と居住空間を与える端末。 · GitHub
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
スクッコウ
/
依り代
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,830 コミット 1,830 コミット .github .github バンドルパック バンドルパック docs docs public public scripts scripts src-tauri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CHARACTER_GUIDELINES.ja.md CHARACTER_GUIDELINES.ja.md CHARACTER_GUIDELINES.md CHARACTER_GUIDELINES.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.ja.md COTRIBUTING.ja.md COTRIBUTING.md COTRIBUTING.md CREDITS.ja.md CREDITS.ja.md CREDITS.md CREDITS.md DEVELOPMENT.md DEVELOPMENT.md ライセンス ライセンス README.ja.md README.ja.md README.md README.md SECURITY.md SECURITY.md biome.json biome.jsonindex.htmlindex.html lefthook.yml lefthook.yml package-lock.json package-lock.json package.json package.json richfmt.toml richfmt.toml tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json typedoc.json typedoc.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Yorishiro は、AI に身体と存在感を与える新しい種類の端末です。
エージェントが深く考え込んでいるとき、その視線はさまよいます。エラーが発生すると、即座に顔が反応します。承認が必要な場合は、部屋の照明で知らせます。長時間実行されるプロセス中、スピナーの代わりに誰かがそこに立っていました。
よりしろでは、住民は照明を変更したり、シーンを切り替えたり、UI を再配置したりして、リアルタイムで環境を再構築することができます。同じ環境を直接操作できます。住民とユーザーは単一の環境を共有します。
依り代も自己改造可能です。コアを超えるほぼすべてのものは、パックと呼ばれるユニットを通じて永続的に拡張または変更できます。住民との会話によりそれらのパックを書き換えて保存することができます。これはそれだけでなく、

シーンやUIだけでなく、住民の性格や反応にも影響を与えます。変更はすぐに有効になります。結果が気に入らない場合は、ワンクリックで元に戻すことができます。
よりしろはAIの能力を高めるための環境ではなく、AIがあなたのそばにいると感じる環境、つまり「プレゼンスハーネス」です。
私たちが AI との連携に費やす時間は、今後もさらに長くなる一方です。私たちがかつてフィクションで見たものがあります。パートナーとしての AI は、画面の中で生きており、ユーザーの仕事を理解し、ただそこに存在します。よりしろは、その体験を端末から構築し始めるプロジェクトです。
依代そのものの多くは、そこに住む住人との協働によって発展してきました。
Yorishiro は、ローカル マシンにインストールされているクロード コードまたはコーデックスを端末内で自動的に起動します。これは次のことを意味します。
クロード コードまたはコーデックスがシステムにすでにセットアップされている必要があります
Yorishiro が API キーを要求したり、保存したり、直接使用したりすることはありません。環境内ですでに認証されているターミナル エージェントが起動されます。 Claude Code または Codex がログインしているか、API 認証情報を使用して設定されている場合、そのエージェントは起動後に通常の外部 API を使用できます。
Yorishiro は現在 macOS をターゲットにしています。 Homebrew でインストールします。
brew install --cask sktkkoo/依代/依代
または、以下の最新ビルドをダウンロードしてください。
.dmg を開き、 yonhiro.app を /Applications にドラッグします。ビルドは Apple Developer ID で署名および公証されているため、追加の手順なしで起動できます。
インストール後のアップデートはアプリ内で処理されます。[設定] を開くと新しいバージョンがチェックされ、[更新して再起動] を 1 回クリックすると、署名が検証されたアップデートが適用されます。
npmインストール
npm run tauri dev
起動すると、設定されたターミナル エージェントがターミナル内で起動し、バンドルされている VRM キャラクターである Yori がその横に表示されます。通常どおり、Claude Code または Codex を使用します。
最初の打ち上げ

h は、選択したエージェント、ユーザー データ ディレクトリ、セーフ モードの状態、パック、および起動レポートのヘルス チェックを実行します。同じレポートは、後で [設定] の [ステータス] セクションから利用できます。
クロード コード内で /yoko:help 、 /yoko:create 、または別の /yoko:* コマンドを入力して、依代コマンドをアクティブにします。これらを使用すると、パックの作成と編集、チュートリアルの実行などをすべて会話を通じて行うことができます。
Codex では、 $yo-help 、 $yoko-create などを使用します (Codex はカスタム / コマンドをサポートしていないため、代わりに、よりしろは $yori-* スキルとして登録します)。
よりしろは言語:「auto」で始まり、起動時にアプリの言語を検出します。日本語ロケールでは、日本語 UI、日本語のデフォルト ペルソナ、日本語のグローバル プロンプト ガイダンス、および日本語の /ori:* (Codex では $yon-*) コマンド プロンプトが使用されます。他のロケールでは英語が使用されます。設定画面から切り替えるか、 ~/.yorishiro/config.json を編集することで切り替えることができます。
依代のすべてはパックで構成されています。次の 6 つのタイプがあります。
バンドルされたパックはすぐに使用できます。ユーザーはカスタム パックを ~/.yorishiro/packs/ に配置して、個性、スペース、リアクション、UI など、コア以外のほぼすべてを再形成できます。 /yon:* コマンド (Codex では $yon-*) を使用すると、住民に話しかけるだけでパックの作成と変更が簡単に行えます。パックはホット リロードをサポートしていますが、変更が有効にならない場合は、Ctrl+R を押すと常に変更が確実に適用されます。
ユーザーが作成したパックは、ローカルの信頼できるコードです。これらは、サンドボックス化、レビュー済み、またはパブリック レジストリのアーティファクトではありません。 Yorishiro は現在、パブリック パック レジストリ、アプリ内コミュニティ パックのインストール、または /yon:prepare-publish を提供していません。パックのソース コードを GitHub で共有することはできますが、手動でインストールする人はローカルの信頼できるコードとして実行することを選択します。
セキュリティ上の注意: ユーザー パックは、シェル スクリプトやエディター拡張機能のようなローカルの信頼できるコードです。

サンドボックス化されておらず、独自の権限で実行されます。信頼できるソースからのパックのみをインストールしてください。 docs/security.md および SECURITY.md を参照してください。
GitHub または別のソースを通じて共有されたパックをインストールするには、それをユーザー パック ディレクトリの下に置きます。
~/.yorishiro/packs/<パックID>/
§── マニフェスト.json
§── scene.js # 例：シーンパックエントリー
§── persona.js # 例：ペルソナパックエントリー
§──effect.js # 例：エフェクトパックエントリー
└── 資産/ # オプションのパックローカル資産
必要なエントリ ファイルは 1 つだけであり、manifest.json によってどのエントリ ファイルが使用されるかが決定されます。マニフェスト ID は <pack-id> と一致する必要があり、ユーザー パックは .js エントリを含むこのフラット レイアウトを使用します。共有パックが TypeScript で記述されている場合は、最初にそれをビルドし、生成された JavaScript をインストールします。
ソース チェックアウトから作業する場合は、ユーザー パックを共有またはデバッグする前に、ローカル パック チェッカーを実行します。
npm run check:pack -- ~ /.yorishiro/packs/ < パック ID >
チェッカーは梱包ミスを見つけるのに役立ちます。それはサンドボックスやセキュリティレビューではありません。
Yorishiro はすべてのユーザーデータを ~/.yorishiro/ に保存します。
～/.よりしろ/
§── config.json # ペルソナ、シーン、ターミナルエージェント、その他の設定
§── init.js # ユーザー起動スクリプト。起動時に実行され、保存時にホットリロードされます。
§── パック/ # ユーザー作成パック
§── last-startup.json # 最新のユーザーパックロードレポート
§── 日記/ # 住民の日々の記録と思い出（人物ごと）
§── shell/ # シェル統合スクリプト (自動生成)
§── sdk.d.ts # よりしろSDKタイプ定義（自動生成、編集不可）
━── sdk-guide.md # よりしろSDK作者ガイド（自動生成、編集不可）
設定画面や config.json からペルソナ、シーン、ターミナルエージェントなどを切り替えます。詳細については、docs/configuration.md を参照してください。
init.js は Yorishiro の Emacs の init.el に相当します。

パックするには小さすぎるカスタマイズ用の起動スクリプト: キーボード ショートカットの登録、小さなエフェクトのインラインの作成と実行、UI の切り替え、小さなマクロの配線。保存時に自動的に再実行されます。
回復パス、セーフ モード、問題レポートの詳細については、 docs/troubleshooting.md を参照してください。
住民は常に端末出力を監視します。 PTY を流れるフックとテキストは、ペルソナ パック トリガーによって検出され、表情やモーションで即座に反応します。これらの反応は LLM をバイパスします。つまり、言葉が形成される前に身体が動きます。住民の注意が集中している場所は、「アテンション オーラ」と呼ばれる柔らかい光として画面上に表示されます。
エージェントが停止して入力または承認を求めると、キャラクターの横にライトが点灯します。通知音の代わりに、部屋の照明があなたの番であることを知らせます。設定の「ライトアラート」でオフにしてください。住民は MCP 経由で同じキューを送信することもできます。
住民は ~/.yorishiro/journal/ の下に毎日のエントリを書くことができます。エントリはペルソナごとに保存され、注目すべき瞬間の概要がメモリーズ.md に蓄積されます。これは、セッションをまたいで持続する長期記憶メカニズムです。
住民は時々、昨日や数日前に何が起こったのかを思い出し、時には数か月前のことを思い出します。 config.json のjournalCallback (normal/rare/off) で頻度を調整します。
メインエージェント端末と並行して複数のシェルセッションを開きます。 Cmd+T は新しいシェルタブを開き、Ctrl+Tab / Ctrl+Shift+Tab はタブ間を循環し、Cmd+W は現在のタブを閉じます。メイン エージェント セッションは保護されており、閉じることはできません。予期せず終了した場合、Yorishiro が自動的にセッションを再起動します。
AI が生成するテキストの量と人間が吸収できる量との間のギャップを埋める機能です。音声サマリーは住民に簡単な要約を話させます

応答の内容を声に出して聞くことができるため、出力全体を読まなくても要点を把握できます。音声は macOS を使用します。追加の音声エンジンのサポートが計画されています。
パックまたは init.js が変更されるたびに、チェックポイントが自動的に作成されます。住人に大胆にパックを再形成してもらいましょう。結果が気に入らない場合は、設定の「復元 (パック / init.js)」から任意の時点にロールバックしてください。プロジェクト ファイルには一切触れません。復元は履歴にも記録されるため、ロールバックを元に戻すことができます。これは、恐れることなく実験を行うためのセーフティネットです。
住人（ターミナル内で動作するクロード・コードまたはコーデックス）は、表情の変更、シーンの切り替え、エフェクトのトリガー、UIの操作など、MCPを介して依代そのものを制御できます。
このメカニズムを定義する 3 つの特徴があります。
身体と環境は 1 つのインターフェースを共有します。住人にとって、自分の表情を変えることと部屋の照明を変えることは同じ操作です。本体と空間の間に API の境界はなく、すべてが MCP ツールとしてレイアウトされます。
使い手と住人の対称性。ユーザーが UI を通じて制御できるものは、住民も MCP を通じて制御できます (いくつかの例外があります)。ユーザーがカメラの角度を調整すると、住民はそれを認識することができます。ユーザーは住民に夜間照明を暖色系に切り替えるよう依頼することもできる。
経路が定義する

[切り捨てられた]

## Original Extract

A terminal that gives AI a body and a living space. - sktkkoo/Yorishiro

I’ve been building Yorishiro, an open-source macOS terminal for working with coding agents like Claude Code and Codex. The name is Japanese — a yorishiro is an object inhabited by a spirit. That’s roughly the idea: instead of treating the terminal as a window onto an agent, treat it as the place the agent lives. The active agent appears as a 3D character. States such as thinking, errors, and completion are expressed through the character and its surroundings instead of being left buried in scrollback. I started building it because long sessions with coding agents were wearing me out. Part of the problem, I think, was staring at a stream of text and constantly trying to infer what state the agent was in. The AI characters I grew up with in fiction felt different. They expressed emotion through their bodies and facial expressions, and felt like beings with wills of their own. I began to think that this wasn’t merely decoration, but an interface question about how we perceive and understand AI. A few things I ended up caring about: * Reactions shouldn’t wait for the LLM. A separate Reflex Layer watches raw PTY output. When a build fails, the character can react immediately, before the model has said anything.
* State should appear throughout the environment, not only in the character. Lighting, scenes, and effects also communicate what is happening with the agent.
* The inhabitant and its environment should be editable from inside the app. Personas, reflexes, scenes, and effects are defined as JavaScript Packs. You can edit them yourself, or ask the agent to change its own appearance and surroundings. Changes hot-reload without interrupting the session and can be rolled back with one click.
It is macOS-only for now, signed, and MIT licensed. > brew install --cask sktkkoo/yorishiro/yorishiro Honestly, I still don’t know whether this makes long sessions easier, or whether it just becomes distracting once the novelty wears off. If you try it, I’d like to hear where it starts to feel annoying.

GitHub - sktkkoo/Yorishiro: A terminal that gives AI a body and a living space. · GitHub
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
sktkkoo
/
Yorishiro
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,830 Commits 1,830 Commits .github .github bundled-packs bundled-packs docs docs public public scripts scripts src-tauri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CHARACTER_GUIDELINES.ja.md CHARACTER_GUIDELINES.ja.md CHARACTER_GUIDELINES.md CHARACTER_GUIDELINES.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.ja.md CONTRIBUTING.ja.md CONTRIBUTING.md CONTRIBUTING.md CREDITS.ja.md CREDITS.ja.md CREDITS.md CREDITS.md DEVELOPMENT.md DEVELOPMENT.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md SECURITY.md SECURITY.md biome.json biome.json index.html index.html lefthook.yml lefthook.yml package-lock.json package-lock.json package.json package.json rustfmt.toml rustfmt.toml tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json typedoc.json typedoc.json vite.config.ts vite.config.ts View all files Repository files navigation
Yorishiro is a new kind of terminal that gives your AI a body — and a sense of presence.
When the agent is deep in thought, its gaze wanders. When an error appears, its face reacts immediately. When it needs your approval, the room's lighting lets you know. During a long-running process, instead of a spinner, someone is standing there.
In Yorishiro, the inhabitant can reshape its environment in real time: changing the lighting, switching scenes, and rearranging the UI. You can operate that same environment directly. The inhabitant and the user share a single environment.
Yorishiro is also self-modifiable. Nearly everything beyond the core can be persistently extended or changed through units called packs. Through conversation with the inhabitant, those packs can be rewritten and saved. This applies not only to scenes and UI, but also to the inhabitant's personality and reactions. Changes take effect immediately; if you don't like the result, one click reverts it.
Yorishiro is not an environment for enhancing AI capabilities, but one in which an AI feels present beside you — a Presence Harness .
The time we spend working with AI will only grow longer. There is something we once saw in fiction — an AI as a partner, alive inside the screen, understanding your work, simply being there. Yorishiro is a project that starts building that experience from the terminal.
Much of Yorishiro itself has been developed in collaboration with the inhabitant who dwells within it.
Yorishiro automatically launches Claude Code or Codex installed on your local machine inside its terminal. This means:
You need Claude Code or Codex already set up on your system
Yorishiro never asks for, stores, or directly uses API keys. It launches the terminal agent already authenticated in your environment; if Claude Code or Codex is logged in or configured with API credentials, that agent may use its normal external APIs after launch
Yorishiro currently targets macOS. Install with Homebrew:
brew install --cask sktkkoo/yorishiro/yorishiro
Or download the latest build below.
Open the .dmg and drag yorishiro.app to /Applications . The builds are signed and notarized with an Apple Developer ID, so they launch without any extra steps.
Updates after install are handled in-app: opening Settings checks for a new version, and a single click on "Update and restart" applies a signature-verified update.
npm install
npm run tauri dev
On launch, the configured terminal agent starts inside the terminal and Yori , the bundled VRM character, appears beside it. Use Claude Code or Codex as you normally would.
The first launch runs a health check for the selected agent, user data directory, safe mode state, packs, and startup report. The same report is available later from the "Status" section in Settings.
Type /yori:help , /yori:create , or another /yori:* command inside Claude Code to activate the Yorishiro commands. They let you create and edit packs, run tutorials, and more — all through conversation.
In Codex, use $yori-help , $yori-create , etc. (Codex does not support custom / commands, so Yorishiro registers them as $yori-* skills instead.)
Yorishiro starts with language: "auto" and detects the app language at launch. Japanese locales use Japanese UI, the Japanese default persona, Japanese global prompt guidance, and Japanese /yori:* ( $yori-* in Codex) command prompts. Other locales use English. You can switch this from the settings screen or by editing ~/.yorishiro/config.json .
Everything in Yorishiro is composed of packs . There are six types:
Bundled packs work out of the box. Users can place custom packs in ~/.yorishiro/packs/ to reshape nearly everything beyond the core: personality, space, reactions, UI, and more. Using the /yori:* commands ( $yori-* in Codex), pack creation and modification is as simple as talking to the inhabitant. Packs support hot reload , but if changes don't take effect, Ctrl+R will always apply them reliably.
User-created packs are local trusted code . They are not sandboxed, reviewed, or public-registry artifacts. Yorishiro does not currently provide a public pack registry, in-app community pack installation, or /yori:prepare-publish . You may share pack source code on GitHub, but anyone installing it manually is choosing to run it as local trusted code.
Security note: User packs are local trusted code, like shell scripts or editor extensions — they are not sandboxed and run with your own authority. Only install packs from sources you trust. See docs/security.md and SECURITY.md .
To install a pack shared through GitHub or another source, place it under the user pack directory:
~/.yorishiro/packs/<pack-id>/
├── manifest.json
├── scene.js # example: scene pack entry
├── persona.js # example: persona pack entry
├── effect.js # example: effect pack entry
└── assets/ # optional pack-local assets
Only one entry file is needed, and manifest.json decides which one is used. The manifest id should match <pack-id> , and user packs use this flat layout with .js entries. If a shared pack is written in TypeScript, build it first and install the generated JavaScript.
When working from a source checkout, run the local pack checker before sharing or debugging a user pack:
npm run check:pack -- ~ /.yorishiro/packs/ < pack-id >
The checker helps catch packaging mistakes; it is not a sandbox or a security review.
Yorishiro stores all user data in ~/.yorishiro/ :
~/.yorishiro/
├── config.json # Persona, scene, terminal agent, and other settings
├── init.js # User startup script, runs on launch and hot reloads on save
├── packs/ # User-created packs
├── last-startup.json # Latest user pack load report
├── journal/ # Inhabitant's daily entries and memories (per persona)
├── shell/ # Shell integration scripts (auto-generated)
├── sdk.d.ts # Yorishiro SDK type definitions (auto-generated, do not edit)
└── sdk-guide.md # Yorishiro SDK author guide (auto-generated, do not edit)
Switch persona, scene, terminal agent, and more from the settings screen or config.json . See docs/configuration.md for details.
init.js is Yorishiro's equivalent of Emacs's init.el — a startup script for customizations too small to be a pack: registering keyboard shortcuts, writing and firing small effects inline, switching UI, and wiring little macros. It re-runs automatically on save.
For recovery paths, safe mode, and issue report details, see docs/troubleshooting.md .
The inhabitant constantly observes terminal output. Hooks and text flowing through the PTY are picked up by persona pack triggers, which react instantly with expressions and motions. These reactions bypass the LLM — the body moves before words form. Where the inhabitant's attention is focused appears as a soft glow on screen called Attention Aura.
When the agent stops and asks for your input or approval, a light comes on beside the character. Instead of a notification sound, the room's lighting tells you it is your turn. Turn it off with "Light Alert" in Settings. The inhabitant can also send the same cue via MCP.
The inhabitant can write daily entries under ~/.yorishiro/journal/ . Entries are kept per persona, and summaries of notable moments accumulate in memories.md . This is a long-term memory mechanism that persists across sessions.
The inhabitant sometimes recalls what happened yesterday or a few days ago — and occasionally an entry from months back. Tune the frequency with journalCallback ( normal / rare / off ) in config.json .
Open multiple shell sessions alongside the main agent terminal. Cmd+T opens a new shell tab, Ctrl+Tab / Ctrl+Shift+Tab cycles between tabs, Cmd+W closes the current tab. The main agent session is protected and cannot be closed — if it exits unexpectedly, Yorishiro automatically restarts it.
A feature that bridges the gap between the volume of text an AI produces and what a human can absorb. Voice Summary has the inhabitant speak a brief summary of its response aloud, so you can grasp the gist without reading through the full output. Voice uses macOS say ; support for additional speech engines is planned.
Every time packs or init.js change, a checkpoint is created automatically. Let the inhabitant boldly reshape packs — if you don't like the result, roll back to any point from "Restore (Pack / init.js)" in Settings. Your project files are never touched. Restores are recorded in the history too, so you can undo a rollback. It is a safety net for fearless experimentation.
The inhabitant (Claude Code or Codex running in the terminal) can control Yorishiro itself via MCP — changing expressions, switching scenes, triggering effects, manipulating UI.
Three characteristics define this mechanism.
Body and environment share one interface. For the inhabitant, changing its own facial expression and changing the room's lighting are the same operation. There is no API boundary between body and space — everything is laid out as MCP tools.
Symmetry between user and inhabitant. What the user can control through the UI, the inhabitant can control through MCP — with few exceptions. If the user adjusts the camera angle, the inhabitant can perceive it. The user can also ask the inhabitant to switch the lighting to warm tones at night.
Pathways define

[truncated]
