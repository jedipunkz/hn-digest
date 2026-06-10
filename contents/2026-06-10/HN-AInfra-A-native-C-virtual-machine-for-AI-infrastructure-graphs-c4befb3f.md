---
source: "https://github.com/TangibleResearch/AInfra"
hn_url: "https://news.ycombinator.com/item?id=48478820"
title: "HN: AInfra – A native C virtual machine for AI infrastructure graphs"
article_title: "GitHub - TangibleResearch/AInfra: A system to define ai infrastructure · GitHub"
author: "reboy"
captured_at: "2026-06-10T19:02:44Z"
capture_tool: "hn-digest"
hn_id: 48478820
score: 2
comments: 1
posted_at: "2026-06-10T16:32:04Z"
tags:
  - hacker-news
  - translated
---

# HN: AInfra – A native C virtual machine for AI infrastructure graphs

- HN: [48478820](https://news.ycombinator.com/item?id=48478820)
- Source: [github.com](https://github.com/TangibleResearch/AInfra)
- Score: 2
- Comments: 1
- Posted: 2026-06-10T16:32:04Z

## Translation

タイトル: HN: AInfra – AI インフラストラクチャ グラフ用のネイティブ C 仮想マシン
記事タイトル: GitHub - TangibleResearch/AInfra: AI インフラストラクチャを定義するシステム · GitHub
説明: AI インフラストラクチャを定義するシステム。 GitHub でアカウントを作成して、TangibleResearch/AInfra の開発に貢献してください。

記事本文:
GitHub - TangibleResearch/AInfra: AI インフラストラクチャを定義するシステム · GitHub
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
タンジブルリサーチ
/
Aインフラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows ainfra-compiler ainfra-compiler docs docs 例 例 infraos-backend infraos-backend infraos-ui infraos-ui infravm infravm オプティマイザー オプティマイザー シェル シェル LICENSE LICENSE Makefile Makefile README.md README.md language.md language.md Push.sh Push.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AInfra は、ソース コードから AI インフラストラクチャを構築するための複合システムです。
AInfra 言語/コンパイラ : .ainfra ソース用の Rust コンパイラ。
AIF : コンパイルされた AI インフラストラクチャ オブジェクト形式。
InfraVM : AIF をロードして PointRun を実行するための C11 ランタイム。
InfraOS : FastAPI + 認証、IDE、Ops Assistant、オブジェクト レジストリ、プロバイダー ステータスを備えたバニラ TypeScript コントロール プレーン。
オプティマイザー : コンパイラと VM の動作を高速化/小型化するための Rust および C ライブラリ。
ビルドする前に次のツールをインストールします。
Rustの安定したツールチェーン: https://rustup.rs
C コンパイラ:clang、gcc、または MSVC/WSL
VM OpenAI コネクタ用のカール開発ヘッダー
xcode-select --install
醸造インストールRustノードPython makeカール
Ubuntu/Debian
sudo apt-get アップデート
sudo apt-get install -y build-essential make pkg-config libcurl4-openssl-dev python3 python3-venv nodejs npmcurl git
カール --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs |しー
窓
最も単純なパスとして WSL2 Ubuntu を使用し、上記の Ubuntu コマンドに従います。ネイティブ Windows サポートはまだ主要なターゲットではありません。
git クローン https://github.com/TangibleResearch/AInfra.git
cd Aインフラ
シェル/infraos.sh 初期化
シェル/infraos.sh ビルド
shell/infraos.sh init はランタイム フォルダーを作成し、SQLite が存在しない場合は初期化し、デフォルトのローカル管理者を作成します。
ユーザー名: 管理者
パスワード: 管理者
localhost 以外のものを公開する前に、このパスワードを変更してください。
シェル/infraos.sh 開始
シェル/

infraos.sh を開く
デフォルトの URL:
バックエンド: http://127.0.0.1:8000
API ドキュメント: http://127.0.0.1:8000/docs
シェル/infraos.sh 停止
AInfra のコンパイルと実行
シェル/infraos.sh コンパイル例/local-stub.ainfra データ/objects/local-stub.aif
shell/infraos.sh run data/objects/local-stub.aif
または、完全なデモを実行します。
シェル/infraos.sh デモ
詳細なドキュメント
コンパイラ、AIF 形式、InfraVM ランタイム、PointRun、プロバイダー コネクタ、InfraOS バックエンド、認証、権限、UI、シェル ワークフロー、セキュリティ制限、ロードマップの完全な技術ウォークスルーについては、以下を参照してください。
クラウド モデルを実行しない限り、プロバイダー キーはオプションです。
エクスポート OPENAI_API_KEY= " ... "
エクスポート ANTHROPIC_API_KEY= " ... "
import GEMINI_API_KEY= " ... "
import GOOGLE_API_KEY= " ... "
エクスポート AZURE_OPENAI_API_KEY= " ... "
エクスポート MICROSOFT_API_KEY= " ... "
エクスポート DEEPSEEK_API_KEY= " ... "
エクスポート HUGGGINGFACE_API_KEY= " ... "
エクスポート HF_TOKEN= " ... "
サーバー設定:
import INFRAOS_SERVER_NAME= " 私の VM サーバー "
エクスポート INFRAOS_BACKEND_PORT=8000
エクスポート INFRAOS_UI_PORT=5173
CI/CD
このリポジトリには、 .github/workflows/ci.yml に GitHub アクションが含まれています。
GitHub Actions は、 main へのプッシュおよびプル リクエストで実行されます。
この結合されたリポジトリは次にも分割されます。
AIF : 言語/コンパイラ/フォーマット
ソース ワークスペースから ./push.sh プッシュを使用して、4 つのリポジトリをすべて再生成して公開します。
AIインフラを定義するシステム
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

A system to define ai infrastructure. Contribute to TangibleResearch/AInfra development by creating an account on GitHub.

GitHub - TangibleResearch/AInfra: A system to define ai infrastructure · GitHub
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
TangibleResearch
/
AInfra
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows ainfra-compiler ainfra-compiler docs docs examples examples infraos-backend infraos-backend infraos-ui infraos-ui infravm infravm optimizer optimizer shell shell LICENSE LICENSE Makefile Makefile README.md README.md language.md language.md push.sh push.sh View all files Repository files navigation
AInfra is the combined system for building AI infrastructure from source code:
AInfra language/compiler : Rust compiler for .ainfra source.
AIF : compiled AI infrastructure object format.
InfraVM : C11 runtime for loading AIF and executing PointRun.
InfraOS : FastAPI + vanilla TypeScript control plane with auth, IDE, Ops Assistant, object registry, and provider status.
Optimizer : Rust and C libraries for faster/smaller compiler and VM behavior.
Install these tools before building:
Rust stable toolchain: https://rustup.rs
C compiler: clang , gcc , or MSVC/WSL
curl development headers for the VM OpenAI connector
xcode-select --install
brew install rust node python make curl
Ubuntu/Debian
sudo apt-get update
sudo apt-get install -y build-essential make pkg-config libcurl4-openssl-dev python3 python3-venv nodejs npm curl git
curl --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs | sh
Windows
Use WSL2 Ubuntu for the simplest path, then follow the Ubuntu commands above. Native Windows support is not the primary target yet.
git clone https://github.com/TangibleResearch/AInfra.git
cd AInfra
shell/infraos.sh init
shell/infraos.sh build
shell/infraos.sh init creates runtime folders, initializes SQLite if missing, and creates the default local admin:
username: admin
password: admin
Change this password before exposing anything beyond localhost.
shell/infraos.sh start
shell/infraos.sh open
Default URLs:
Backend: http://127.0.0.1:8000
API docs: http://127.0.0.1:8000/docs
shell/infraos.sh stop
Compile And Run AInfra
shell/infraos.sh compile examples/local-stub.ainfra data/objects/local-stub.aif
shell/infraos.sh run data/objects/local-stub.aif
Or run the full demo:
shell/infraos.sh demo
Deep Documentation
For a full technical walkthrough of the compiler, AIF format, InfraVM runtime, PointRun, provider connectors, InfraOS backend, authentication, privileges, UI, shell workflow, security limits, and roadmap, read:
Provider keys are optional unless you run cloud models:
export OPENAI_API_KEY= " ... "
export ANTHROPIC_API_KEY= " ... "
export GEMINI_API_KEY= " ... "
export GOOGLE_API_KEY= " ... "
export AZURE_OPENAI_API_KEY= " ... "
export MICROSOFT_API_KEY= " ... "
export DEEPSEEK_API_KEY= " ... "
export HUGGINGFACE_API_KEY= " ... "
export HF_TOKEN= " ... "
Server settings:
export INFRAOS_SERVER_NAME= " My VM Server "
export INFRAOS_BACKEND_PORT=8000
export INFRAOS_UI_PORT=5173
CI/CD
This repo includes GitHub Actions in .github/workflows/ci.yml .
GitHub Actions runs on pushes and pull requests to main .
This combined repository is also split into:
AIF : language/compiler/format
Use ./push.sh push from the source workspace to regenerate and publish all four repos.
A system to define ai infrastructure
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
