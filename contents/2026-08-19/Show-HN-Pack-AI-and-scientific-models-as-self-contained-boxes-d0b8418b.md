---
source: "https://github.com/suffro/scrollcase"
hn_url: "https://news.ycombinator.com/item?id=49362094"
title: "Show HN: Pack AI and scientific models as self-contained boxes"
article_title: "GitHub - suffro/scrollcase: Pack AI and scientific models as self-contained boxes. Pinned, validated, signed and air-gappable, for macOS Metal, Linux and Windows (CPU and CUDA). · GitHub"
image: "https://opengraph.githubassets.com/79c55f0d562da1c578c14e696ec3e9bdc6c778e77ecbb0f09c55db4dbfd417ae/suffro/scrollcase"
author: "Suffro"
captured_at: "2026-08-19T15:22:44Z"
capture_tool: "hn-digest"
hn_id: 49362094
score: 1
comments: 0
posted_at: "2026-08-19T14:27:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pack AI and scientific models as self-contained boxes

- HN: [49362094](https://news.ycombinator.com/item?id=49362094)
- Source: [github.com](https://github.com/suffro/scrollcase)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T14:27:14Z

## Translation

タイトル: Show HN: AI と科学モデルを自己完結型のボックスとしてパックする
記事のタイトル: GitHub - suffro/scrollcase: AI と科学モデルを自己完結型のボックスとしてパックします。 macOS Metal、Linux、Windows (CPU および CUDA) 向けに固定、検証、署名、エアギャップ可能。 · GitHub
説明: AI と科学モデルを自己完結型のボックスとしてパックします。 macOS Metal、Linux、Windows (CPU および CUDA) 向けに固定、検証、署名、エアギャップ可能。 - サフロ/スクロールケース
HN テキスト: こんにちは!私は Scrollcase というオープンソース プロジェクトに取り組んでいます。私がこれを始めたのは、マシン間で Python または ML プロジェクトを移動するのが面倒なことが多いためです。適切な Python バージョン、ライブラリ、ネイティブ ビルド、モデル ファイルなどが必要です。そのため、Scrollcase では、誰かがマシン上で Python 環境を再構築する代わりに、Python、ロックされた依存関係、コード、モデル ファイルを含むすべてを一度にパッケージ化します。その後、それを検証し、解凍して実行できます。 Python のインストール、pip のインストール、Docker の必要がなく、依存関係を管理する必要もありません。内部に何を入れるかを定義し、依存関係をロックし、すぐに実行できる移植可能な自己完結型の検証可能なパッケージを構築するだけです。自分が制御していないマシン用に Python、ML、または科学環境をパッケージ化する必要がある人々からのフィードバックをお待ちしています。

記事本文:
GitHub - suffro/scrollcase: AI と科学モデルを自己完結型のボックスとしてパックします。 macOS Metal、Linux、Windows (CPU および CUDA) 向けに固定、検証、署名、エアギャップ可能。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
サフロ
/
スクロールケース
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
270 コミット 270 コミット フォルダーとファイル
.claude .claude .github .github .scrollcase-runtime-types-RzUfxr/ src .scrollcase-runtim

e-types-RzUfxr/ src docs docs 例 例 Python Python Rust Rust スクリプト スクリプト src src テスト テスト .gitattributes .gitattributes .gitignore .gitignore AGENT-POLICY.md AGENT-POLICY.md AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python 環境全体を自己完結型のボックスとしてパックする
Scrollcase は、Python 環境全体とそれが実行するコード (LLM や科学モデルなど) を、単一の自己完結型でポータブルな署名付きアーカイブ (ボックス) に詰め込みます。
あなたはその箱を他の人に渡します。開梱して実行するだけです。インストールするものは何もありません。Python、pip install、コンパイラ、Docker、維持する依存関係はありません。
すべてのボックスには署名が付いているため、受け取った人は、それが自分が作成したものであり、途中で変更されたものではないことを確認できます。
ビルドは決定的です。同じコミットをリビルドすると同じバイトが返されます。出荷したものは誰でも再現できます。
これが基本的に全体のアイデアです
Python ランタイムを他人のマシンに入手するということは、通常、Python ランタイムを再構築するよう依頼することを意味します。
環境: 適切な Python、適切なライブラリ、CPU または GPU に適したネイティブ ビルド、
適切な場所から適切な重みをダウンロードします。機能しなくなるまで機能します - そして壊れます
あなたのマシンではなく、彼らのマシンです。スクロールケースの動きは、あなたが制御するマシンで一度だけ時間を稼ぐために機能します。
そして結果をファイルに変換します。
言葉
意味
スクロール
作成するファイル: 依存関係、モデル ファイル、何を実行するか、どのようにテストするか。ビルドが受け入れる唯一の入力。 →参考
ボックス
出てくるもの: 内部環境全体を含む 1 つのアーカイブ。 →フォーマット
ターゲット
どのマシン

対象となるのは、オペレーティング システム、CPU アーキテクチャ、アクセラレータ (および CUDA バージョン) です。 1 つのボックス、1 つのターゲット。
リリース
消費者が箱を確認できるように、箱について説明する署名済みの文書。 → セキュリティモデル
箱の中には何が入っているのか
エンティティ
中には何が入っているのか
Pythonインタプリタ
選択した正確なバージョン。ホストには Python はまったく必要ありません。
あらゆる依存関係
ロック ファイルが固定したバージョンの Conda および PyPI パッケージ、ネイティブ ライブラリが含まれます。
あなたのコード
アプリケーション ファイル、開始するエントリ スクリプトまたはモジュール。
モデルファイル
アーカイブに埋め込まれるか、サイズとハッシュが記録された状態でアーカイブの外に保管されます。
署名付きメタデータ
このボックスが何であるか、何が含まれているか、そしてそのダイジェスト - したがって、消費者はそれ以外のものを拒否できます。
ライセンスインベントリ
すべての依存関係のライセンスはロックから派生し、推測されません。
ボックスは 1 つのターゲット (1 つのオペレーティング システム、1 つの CPU アーキテクチャ、1 つのアクセラレータ) に対して構築されます。
macos-aarch64-metal と linux-x86_64-cuda12 は 2 つのボックスであり、オプションが付いている 1 つのボックスではありません。つまり
意図的 — どこでも動作することを約束したボックスは、インストール時に物事を決定する必要があります。
それが取り除かれる問題です。
物を梱包する人またはチーム。
スクロールで環境を説明します。
依存関係、ファイル、モデル資産を宣言します。
結果のファイルを好きな場所に公開します。
ロックから環境を作成し、宣言されたアセットをダウンロードしてハッシュチェックし、ツリーを作成します
再配置可能、ファイルのコピー、スクロールで宣言されたテストの実行、アーカイブの生成、
ハッシュを計算し、リリース文書を作成し、署名します。
開発者はこれらのステップを 1 つずつ実行したり、環境パスを修復したりしません。
マニフェストを作成するか、手動でファイルに署名します。
ボックスをインストールして使用するアプリケーション。
ユーザーのマシンに適したボックスを選択します。
ローカルリリース、アーカイブを提供し、

適合する消費者への鍵を信頼します。
更新、アクティブ化、ロールバック、および削除を所有します。
公式の Node、Python、Rust コンシューマーは、呼び出し元を検証し、安全に抽出し、ローカル ボックスを実行します。
すでに成立しています。彼らはチャンネルを選択したり、ダウンロードしたりしません。
6 ステップを 1 回。その後、新しいバージョンは通常、単一のビルドで出荷されます。
スクロールケースの初期化
プロジェクト構造を作成し、尋ねた後、デフォルトで「はい」に設定します。使い捨ての実行可能ファイルを作成します。
自分のマシン用の example-box、短い SCROLLCASE.md 、TypeScript、Python、Rust
Consumer-templates/ の下に例があります。空のワークスペースには --no-example を渡します。
スクロールケース 新しいスクロール
ターゲット、ボックス ID、パッケージ化しているもののアップストリーム リビジョン、および場所の 4 つの質問をします。
ボックスが公開され、ターゲット固有の 1 つのscroll.json 、その pixi.toml 、および
スターター self_test.py 。既存のものは上書きされません。最初に周囲を見回すには、次の例を使用します。
初期化が作成されました。 → スクロール参照
スクロールで宣言されているものはすべて、ファイルを手動で編集するのではなく、コマンドによって追加されます。
scrollcase add dep my-model onnxruntime # 依存関係
スクロールケース追加アセット my-model https://…/model.safetensors # 一度ダウンロードし、サイズとハッシュを記録します
スクロールケース追加ファイル my-model runtime/entrypoint.py # このプロジェクトのファイル
対応するのは、remove 、 editscroll 、および Refresh です。
→ CLIリファレンス
スクロールケース ロック my-model/macos-aarch64-metal
Git にコミットする pixi.lock への依存関係を一度解決します。そこからのビルドは
インストールしますが、決して解決されません。これにより、同じコミットの 2 つのビルドが同じものを生成します。
バイト。
依存関係が変更された場合にのみ、ロックを再実行します。 →なぜピクシーなのか
スクロールケースキー生成
リリースの署名に使用されるキー ペアを作成します。秘密キーがリポジトリに入ることはありません。の
公開鍵は次の場所に移動します

これにより、自分のボックスと他の人のボックスを区別できるようになります。
ローカル キーの代わりに、実際のキー管理 (KMS、HSM、署名サービス) がプラグインされます。
→ サインと鍵の保管
スクロールケースドクター --scroll my-model/macos-aarch64-metal # オプション: このマシンでビルドできますか?
スクロールケース ビルド my-model/macos-aarch64-metal
ビルドは、ターゲットと互換性のあるマシン上で実行する必要があります。ボックスアーカイブ、署名付きのアーカイブを作成します。
リリースおよびチャネルのドキュメント、およびすぐに公開できるディレクトリ ツリー。
宣言されたインポートが失敗した場合、アセット ハッシュが一致しない場合、またはパリティ チェックが許容範囲に違反した場合、
箱はありません。ゲートが失敗しても、署名されたアーティファクトが生成されることはありません。
ターミナルからワンショット実行する場合:
スクロールケース実行 ./release.json --archive ./box.zip -- --help
最初に検証し、署名されたスクリプトまたはモジュールをシェルなしで実行し、子の終了を保持します。
状態を確認し、一時的な抽出を削除します。
アプリケーションはライブラリを通じて同じことを行います。つまり、scrollcase/consumer の Node API、
Python パッケージscrollcase_consumer、またはRustクレートscrollcase-consumer。 3人全員が共有しているのは、
同じ検証、安全な抽出、実行、受信、シグナル、クリーンアップ、およびオンデマンド資産
セマンティクスがあり、どれも何もダウンロードしません。
再起動してもボックスを抽出したままにするアプリケーションは、解凍するのではなくボックスに再接続します。
またまた。 → ライブラリ API ·
抽出したボックスを保持する
スクロールケースはファイルを書き込んで停止します。それらのアップロードは、スクリプトを使用して手動で行う必要があります。
CI/CD、またはオブジェクト ストレージ パイプライン経由。
スクロールケースでファイルを構築します
↓
導入システムがそれらをアップロードします
↓
アプリケーションがそれらをダウンロード、検証、実行します
この境界は、この形式がオブジェクト ストレージ、GitHub リリース、プライベート ストレージで機能する理由です。
サーバー、またはデスクトップ アップデーター

すぐに持っています。
→ ボックスの配布
何が変わったのか
何を実行するか
コードまたはインクルードされたファイル
バージョンを上げる、ビルドする
依存関係
ロックし、結果を確認し、ビルドします
モデルの重み
スクロール内のアセットを更新し、バージョンを上げ、ビルドします
複数のターゲットとは、それぞれのターゲットと互換性のあるマシン上の複数のビルドを意味します。
通常、Scrollcase 自体よりも、幅広いプラットフォーム マトリクスの方が高価になります。
→ プラットフォーム例
消費側アプリケーションがまだ所有しているもの
スクロールケースが箱を作ります。周りのアプリケーションは実装しません。そのアプリケーション
マシンを検出し、リリースを選択し、マニフェストとアーカイブをダウンロードし、
署名、ランタイム要件のチェック、ボックスの抽出、オンデマンド アセットの取得、
ランタイムを管理し、更新とアンインストールを処理します。
公式コンシューマは、検証、抽出、実行をカバーします。選択、ダウンロード、および
更新ポリシーは製品に付属します。「スクロールケースを使用する理由」を参照してください。のために
なぜその場所に線が引かれているのか。
エンド ユーザーにとって重要なのは、機能が何も表示されないということです。機能を選択し、インストールを押し、
残りはアプリケーションが行います。
環境を説明する
↓
バージョンをロックする
↓
箱を作る
↓
ファイルを公開する
実際には、この困難はツールではなく、パッケージ化している環境に起因します。
扱いにくいネイティブ ライブラリ、古い科学パッケージ、いくつかの CUDA バージョン、非常に大きな重み、
サポートされているチャネルにない依存関係。スクロールケースはそれらを消すのではなく、消してしまいます。
すべてのユーザーからのサポート チケットではなく、一度解決するビルド時の問題です。
ディスク上の署名済みの検証済みボックスで停止します。アーカイブをホストしたり、レジストリを実行したり、どれを選択するかを決定したりすることはありません。
クライアントがインストールすべきバージョン、ボックスのダウンロード、更新の管理、CI ランナーの割り当て、または判断
モデルかどうか

科学的に正しいです。それらはそれを使用する人のものです。それが、
オブジェクト ストレージ、GitHub リリース、プライベート サーバー、または既存のアップデータで使用できる形式。
各デモでは、プロジェクトで Scrollcase を活用する方法を理解するのに役立つさまざまな機能と使用例を紹介します。
Apache-2.0 、Scrollcase 独自のソースをカバーしています。それが構築するボックスの中身 —
インタプリタ、conda-forge と PyPI の依存関係、モデル コードと重み — 独自のライセンスを保持します。
これはまさに、すべてのボックス内のライセンス監査が記録するために存在するものです。 「通知」を参照してください。
TL;DR — 同じ内容が 1 分で読めるページにあります。
クイックスタート — 今すぐサンプル ボックスを構築します。
なぜスクロールケースなのか? — そして、よりシンプルなツールがより良い選択である場合。
デモをお試しください。さまざまな機能や使用例、および Scrollcase を独自のプロジェクトで機能させる方法を示す実際の例です。
AI と科学モデルを自己完結型のボックスとしてパックします。 macOS Metal、Linux、Windows (CPU および CUDA) 向けに固定、検証、署名、エアギャップ可能。
Readme Apache-2.0 ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Pack AI and scientific models as self-contained boxes. Pinned, validated, signed and air-gappable, for macOS Metal, Linux and Windows (CPU and CUDA). - suffro/scrollcase

Hi! I've been working on an open-source project called Scrollcase. I started this cause moving Python or ML projects between machines is often painful. You need the right Python version, libraries, native builds, model files, and so on. So instead of making someone rebuild a Python environment on their machine, Scrollcase packages everything once, including Python, locked dependencies, code, and model files. Then they can verify it, unpack it, and run it. No Python install, no pip install, no Docker, and no dependency to mantain. You just define what goes inside, lock the dependencies, and build a portable, self-contained, verifiable package that's ready to run. Would love feedback from people who had to package Python, ML, or scientific environments for machines they don't control.

GitHub - suffro/scrollcase: Pack AI and scientific models as self-contained boxes. Pinned, validated, signed and air-gappable, for macOS Metal, Linux and Windows (CPU and CUDA). · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
suffro
/
scrollcase
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
270 Commits 270 Commits Folders and files
.claude .claude .github .github .scrollcase-runtime-types-RzUfxr/ src .scrollcase-runtime-types-RzUfxr/ src docs docs examples examples python python rust rust scripts scripts src src tests tests .gitattributes .gitattributes .gitignore .gitignore AGENT-POLICY.md AGENT-POLICY.md AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Pack entire Python environments as self-contained boxes
Scrollcase packs an entire Python environment and the code it runs — like an LLM or a scientific model — into a single, self-contained , portable and signed archive: a box .
You give that box to someone else. They unpack it and run it — that's it! Nothing to install : no Python, no pip install, no compiler, no Docker, no dependencies to maintain .
Every box is signed , so whoever receives it can check that it is exactly the one you built and not something that changed on the way over.
Builds are deterministic : rebuilding the same commit gives the same bytes back — anyone can reproduce what you shipped.
This is basically the whole idea
Getting a Python runtime onto someone else's machine normally means asking them to rebuild your
environment: the right Python, the right libraries, the right native builds for their CPU or GPU,
the right weights downloaded from the right place. It works until it doesn't — and it breaks on
their machine, not yours. Scrollcase moves that work to build time, once, on a machine you control,
and turns the result into a file.
Word
Meaning
scroll
The file you write: dependencies, model files, what to run, how to test it. The only input a build accepts. → reference
box
What comes out: one archive with the whole environment inside. → format
target
Which machine it is for: operating system, CPU architecture, accelerator (and CUDA version). One box, one target.
release
The signed document that describes the box, so a consumer can verify it. → security model
What is inside a box
Entity
What's inside
Python interpreter
The exact version you chose. The host does not need Python at all.
Every dependency
Conda and PyPI packages, native libraries included, at the versions your lock file pinned.
Your code
Application files, an entry script or module to start.
Model files
Embedded in the archive, or kept outside it with their size and hash recorded.
Signed metadata
What this box is, what it contains, and its digest — so a consumer can reject anything else.
Licence inventory
Every dependency's licence, derived from the lock, not guessed.
A box is built for one target : one operating system, one CPU architecture, one accelerator.
macos-aarch64-metal and linux-x86_64-cuda12 are two boxes, not one box with options. That is
deliberate — a box that promised to work everywhere would have to decide things at install time,
which is the problem being removed.
The person or team packaging the thing.
describes the environment in a scroll ;
declares dependencies, files, and model assets;
publishes the resulting files wherever they like.
Creates the environment from the lock, downloads and hash-checks declared assets, makes the tree
relocatable, copies your files in, runs the tests the scroll declares, produces the archive,
computes the hashes, writes the release documents, and signs them.
The developer does not run those steps one by one, and does not repair environment paths, write
manifests, or sign files by hand.
The application that installs and uses the box.
picks the right box for the user's machine;
hands the local release, archive, and trust keys to a conforming consumer;
owns updates, activation, rollback, and removal.
The official Node, Python, and Rust consumers verify, safely extract, and run a local box the caller
already holds. They do not choose channels and they do not download.
Six steps, once. After that, shipping a new version is usually a single build .
scrollcase init
Creates the project structure and — after asking, defaulting to yes — a disposable runnable
example-box for your own machine, a short SCROLLCASE.md , and TypeScript, Python, and Rust
examples under consumer-templates/ . Pass --no-example for an empty workspace.
scrollcase new scroll
Asks four questions — target, box id, the upstream revision of what you are packaging, and where
boxes will be published — and writes one target-specific scroll.json , its pixi.toml , and a
starter self_test.py . Nothing existing is overwritten. To just look around first, use the example
init created. → Scroll reference
Everything the scroll declares is added by command, not by hand-editing files:
scrollcase add dep my-model onnxruntime # a dependency
scrollcase add asset my-model https://…/model.safetensors # downloads once, records size and hash
scrollcase add file my-model runtime/entrypoint.py # a file from this project
remove , edit scroll , and refresh are the counterparts.
→ CLI reference
scrollcase lock my-model/macos-aarch64-metal
Resolves the dependencies once into a pixi.lock you commit to Git. From then on the build
installs — it never resolves — which is what makes two builds of the same commit produce the same
bytes.
Re-run lock when a dependency changes, and only then. → Why pixi
scrollcase keygen
Creates the key pair used to sign releases. The private key never goes into the repository; the
public key goes to the consuming application, which is how it can tell your box from anyone else's.
Real key custody — a KMS, an HSM, a signing service — plugs in instead of the local key.
→ Signing and key custody
scrollcase doctor --scroll my-model/macos-aarch64-metal # optional: can this machine build it?
scrollcase build my-model/macos-aarch64-metal
The build must run on a machine compatible with the target. It produces the box archive, the signed
release and channel documents, and a publication-ready directory tree.
If a declared import fails, an asset hash does not match, or a parity check breaches its tolerance,
there is no box. A failed gate never produces a signed artefact.
For a one-shot run from the terminal:
scrollcase run ./release.json --archive ./box.zip -- --help
It verifies first, runs the signed script or module without a shell, preserves the child's exit
status, and removes its temporary extraction.
An application does the same thing through a library: the Node API at scrollcase/consumer , the
Python package scrollcase_consumer , or the Rust crate scrollcase-consumer . All three share the
same verification, safe extraction, execution, receipt, signal, cleanup, and on-demand asset
semantics, and none of them downloads anything.
An application that keeps a box extracted across restarts re-attaches to it rather than unpacking
again. → Library APIs ·
Keeping an extracted box
Scrollcase writes files and stops. Uploading them is yours to do — by hand, with a script, from
CI/CD, or through an object-storage pipeline.
Scrollcase builds the files
↓
your deployment system uploads them
↓
your application downloads, verifies, and runs them
This boundary is the reason the format works with object storage, GitHub Releases, a private
server, or a desktop updater you already have.
→ Distributing boxes
What changed
What to run
Code or included files
bump the version, build
Dependencies
lock , review the result, build
Model weights
update the asset in the scroll, bump the version, build
Several targets mean several builds, each on a machine compatible with its target — and that, rather
than Scrollcase itself, is usually where a wide platform matrix gets expensive.
→ Platform examples
What the consuming application still owns
Scrollcase builds the box; it does not implement the application around it. That application
detects the machine, chooses a release, downloads the manifests and archive, verifies the
signatures, checks runtime requirements, extracts the box, fetches any on-demand assets, starts the
runtime, and handles updates and uninstallation.
The official consumers cover verification, extraction, and execution. Selection, download, and
update policy stay with the product — see Why Scrollcase? for
why that line is drawn where it is.
For the end user, the point is that none of it is visible: they pick a feature, press install, and
the application does the rest.
describe the environment
↓
lock the versions
↓
build the box
↓
publish the files
In practice the difficulty comes from the environment you are packaging, not from the tool:
awkward native libraries, old scientific packages, several CUDA versions, very large weights, a
dependency that is not on a supported channel. Scrollcase does not make those disappear — it makes
them a build-time problem you solve once, instead of a support ticket from every user.
It stops at a signed, verified box on disk. It does not host archives, run a registry, decide which
version a client should install, download boxes, manage updates, allocate CI runners, or judge
whether a model is scientifically correct. Those belong to whoever uses it — which is what keeps the
format usable with object storage, GitHub Releases, a private server or an existing updater.
Each demo showcases different features and use cases to help you understand how to leverage Scrollcase in your projects.
Apache-2.0 , covering Scrollcase's own source. The contents of the boxes it builds —
interpreters, conda-forge and PyPI dependencies, model code and weights — carry their own licences,
which is exactly what the licence audit inside every box exists to record. See NOTICE .
TL;DR — the same thing in a page you can read in a minute.
Quickstart — build the example box now.
Why Scrollcase? — and when a simpler tool is the better choice.
Try a demo — worked examples that show different features and use cases, and how to put Scrollcase to work in your own project.
Pack AI and scientific models as self-contained boxes. Pinned, validated, signed and air-gappable, for macOS Metal, Linux and Windows (CPU and CUDA).
Readme Apache-2.0 license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
