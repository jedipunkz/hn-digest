---
source: "https://github.com/zeborg/kubekosh"
hn_url: "https://news.ycombinator.com/item?id=48587008"
title: "Show HN: KubeKosh – Self-Hosted Kubernetes Playground in a Docker Container"
article_title: "GitHub - zeborg/kubekosh: Interactive Kubernetes Playground · GitHub"
author: "zeborg"
captured_at: "2026-06-18T16:13:27Z"
capture_tool: "hn-digest"
hn_id: 48587008
score: 2
comments: 0
posted_at: "2026-06-18T15:34:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: KubeKosh – Self-Hosted Kubernetes Playground in a Docker Container

- HN: [48587008](https://news.ycombinator.com/item?id=48587008)
- Source: [github.com](https://github.com/zeborg/kubekosh)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T15:34:53Z

## Translation

タイトル: Show HN: KubeKosh – Docker コンテナ内の自己ホスト型 Kubernetes プレイグラウンド
記事タイトル: GitHub - zeborg/kubekosh: インタラクティブ Kubernetes プレイグラウンド · GitHub
説明: インタラクティブな Kubernetes プレイグラウンド。 GitHub でアカウントを作成して、zeborg/kubekosh の開発に貢献してください。
HN テキスト: 皆さんこんにちは!私は最近、KubeKosh を公開しました。これは、CKA、CKS、CKAD の準備だけでなく、初心者向けの 85 以上のユニークな練習シナリオで構成される自己ホスト型のオープンソース Kubernetes ラボです。新しいシナリオは、仲間のユーザーが JSON 形式で作成して寄稿することができます。また、新しい寄稿者が独自のシナリオを作成してテストするのを支援するために、ドキュメントでそのスキーマについても言及しました。すぐに試してみるには、この 1 つのコマンドを実行するだけです。単一の Docker コンテナーでラボ全体を最大 15 秒以内に起動します。 docker run -itd --name kubekosh --privileged -p 7554:80 zeborg/kubekosh:latest 実稼働環境でよくある状況を扱う、初心者にも専門家にも同様に新しい、または洞察力に富む新しいシナリオの貢献を感謝します。スクリーンショットとプロジェクトの詳細については、README.md を参照してください。お時間をいただき、ご清聴いただきありがとうございました!

記事本文:
GitHub - zeborg/kubekosh: インタラクティブ Kubernetes プレイグラウンド · GitHub
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
ゼボーグ
/
クベコシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

e actions menu Folders and files
38 コミット 38 コミット .github/ workflows .github/ workflows バックエンド バックエンド フロントエンド フロントエンド シナリオ シナリオ スクリーンショット スクリーンショット スクリプト スクリプト Dockerfile Dockerfile ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Self-hosted Kubernetes Lab for Hands-on Learning
KubeKosh は、単一の Docker コンテナ内で実際の K3s Kubernetes クラスタを実行し、それをブラウザベースのターミナルと自動シナリオ検証と組み合わせます。クラウド アカウントやローカル クラスタは必要ありません。
docker run -itd --name kubekosh --privileged -p 7554:80 zeborg/kubekosh:latest
http://localhost:7554 を開きます — Cluster Ready インジケーターが緑色に変わるまで約 30 秒待ちます。
--privileged が必要です — K3s はカーネル名前空間と cgroup にアクセスする必要があります。
セキュリティ警告: このコンテナを公開しないでください。これは教育目的のみを目的としているため、ローカル マシンでのみ使用してください。
docker run -itd --name kubekosh --privileged -p 7554:80 \
-v < your_custom_directory > :/data zeborg/kubekosh:latest
進行状況は、コンテナー内の /data/progress.db にある SQLite に保存されます。コンテナーの再起動後も進行状況を維持するには、ローカル ディレクトリを /data にマウントします。
docker build -t kubekosh .
# マルチプラットフォーム
docker buildx build --platform linux/amd64,linux/arm64 -t kubekosh 。
何が入っているのか
バンドル
フォーカス
試験モード
🌱 Kubernetes の基本
中心となる概念
60分
🧑‍ ✈️ Kubernetes Administrator
CKA
120分
🛠️ Kubernetes Developer
CKAD
120分
🛡️ Kubernetes Security
CKS
120分
シナリオの種類:
タスク — ライブターミナルでのハンズオンチャレンジ。 「検証」をクリックすると、クラスター状態チェックが自動化されます。
MCQ — 提出に関する詳細な説明が記載された多肢選択式の質問。
The terminal comes pre-configured with:
コンポーネント
テクノロジー
フロントエンド
反応 + ビタミン

e、xterm.js
バックエンド
Node.js / Express、node-pty WebSocket PTY
クラスター
K3 (単一ノード、コンテナ内)
プロキシ
コンテナ ポート 80 上の nginx、ホスト ポート 7554 にマッピング
ストレージ
SQLite ( better-sqlite3 ) /data/progress.db
すべては script/entrypoint.sh によって管理される単一の Docker イメージ内で実行されます。
シナリオ/
§── data/ # シナリオごとに 1 つの JSON ファイル -> <scenario-id>.json
§──bundles/ # バンドルごとに 1 つの JSON ファイル -> <bundle-id>.json
└── SCHEMA.md # 完全なスキーマ リファレンス
バックエンド/
━──server.js # Express API + WebSocket PTY
フロントエンド/
└── src/ # React + Vite SPA
スクリプト/
§──entrypoint.sh # コンテナ起動 (k3s -> API -> nginx)
└── nginx.conf # リバースプロキシ設定
貢献する
このようなオープンソース プロジェクトを成長させるのは貢献です。大小を問わず、あらゆる貢献が重要です。タイプミスを修正する場合でも、シナリオの説明を磨き上げる場合でも、まったく新しい演習をゼロから構築する場合でも、次の人が可能な限り最良の方法で Kubernetes を学習できるよう支援することになります。お時間を割いていただきありがとうございます!
各シナリオは、scenarios/data ディレクトリ内の単一の JSON ファイルです。各バンドルは、scenarios/bundles ディレクトリ内の単一の JSON ファイルです。完全なスキーマについては、scenarios/SCHEMA.md を参照してください。
validation.commands — 冪等の kubectl コマンドのみ
setup_commands /teardown_commands — kubectl またはネイティブ Ubuntu コマンドのみ
correct_option は、options[].id 値の 1 つと一致する必要があります
インメモリキャッシュとホットリロード
高いパフォーマンスとディスク I/O ボトルネックのゼロを保証するために、シナリオとバンドルはバックエンドの起動時にメモリにキャッシュされます。シナリオを開発または更新する場合、イメージを再構築したりコンテナーを再起動したりせずに、定義をホットリロードできます。
シナリオ ディレクトリのマウント: ローカル シナリオ/ディレクトリ mounte でコンテナを実行します。

d から /app/scenarios :
docker run --rm -itd --privileged -p 7554:80 --name kubekosh -v < シナリオディレクトリへのパス > :/app/scenarios zeborg/kubekosh:latest
キャッシュのリロード: Web ユーザー インターフェイスのヘッダーの右上隅にある [シナリオ キャッシュのリロード] (↻) ボタンをクリックするか、API リクエストを送信します。
curl -X POST http://localhost:7554/api/cache/reload
注:
<path_to_scenarios_directory> の内容は、更新を含む複製されたリポジトリのローカル scenarios/ ディレクトリへのパスである必要があります。つまり、更新された scenarios/data および scenarios/bundles ディレクトリが含まれている必要があります。
#1. GitHub 上のリポジトリをフォークし、フォークのクローンを作成します。
git clone https://github.com/ < ユーザー名 > /kubekosh.git
CD クベコシュ
# 2. ブランチを作成する
git checkout -b feat/my-scenario
# 3. 新しいシナリオ ファイルを追加します (既存のシナリオをテンプレートとしてコピーするか、新しいシナリオを作成します)
cp シナリオ/データ/deploy-nginx.json シナリオ/データ/my-new-scenario.json
vim scenarios/data/my-new-scenario.json # [SCHEMA.md](scenarios/SCHEMA.md) に従って新しいシナリオを編集します
# 4. シナリオ ID を関連バンドルに追加する
vim scenarios/bundles/k8s-basics.json # 新しいシナリオ ID を含むようにバンドルを編集します
#5. ローカルでビルドしてテストする
# 構築されたコンテナを直接実行します。
docker build -t kubekosh 。 && docker run --rm -itd --privileged -p 7554:80 --name kubekosh kubekosh
# または、ホットリロード用にシナリオ フォルダーをマウントします。
docker run --rm -itd --privileged -p 7554:80 -v $PWD /scenarios:/app/scenarios --name kubekosh zeborg/kubekosh:dev
# 6. フォークにコミットしてプッシュする (`my-new-scenario` を `k8s-basics` バンドルに追加する例)
git add scenarios/data/my-new-scenario.json シナリオ/bundles/k8s-basics.json
git commit -m " feat: k8s-basics バンドルに my-new-scenario を追加 "
git Push -u Origin feat/my-new-scenario
プルリクエストを開く

main に対するフォークのブランチから。
Apache 2.0 ライセンス — 「ライセンス」を参照してください。
インタラクティブな Kubernetes プレイグラウンド
読み込み中にエラーが発生しました。このページをリロードしてください。
32
フォーク
レポートリポジトリ
リリース
7
リリース v0.2.2
最新の
2026 年 6 月 18 日
+ 6 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Interactive Kubernetes Playground. Contribute to zeborg/kubekosh development by creating an account on GitHub.

Hey everyone! I've recently published KubeKosh, which is a self-hosted open-source Kubernetes lab consisting of 85+ unique practice scenarios for beginners as well as CKA, CKS, and CKAD preparation. New scenarios can be created and contributed by fellow users in JSON format, and I've also mentioned its schema in the documentation for helping new contributors create and test their own scenarios. To quickly try it out, you just need to run this 1 command. It spins up the entire lab in a single Docker container within ~15 seconds: docker run -itd --name kubekosh --privileged -p 7554:80 zeborg/kubekosh:latest I'd appreciate new scenario contributions dealing with situations which are common in production environments, and which may be new or insightful to beginners and experts alike. Please refer to README.md for screenshots and more details about the project. Thanks for your time and attention!

GitHub - zeborg/kubekosh: Interactive Kubernetes Playground · GitHub
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
zeborg
/
kubekosh
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .github/ workflows .github/ workflows backend backend frontend frontend scenarios scenarios screenshots screenshots scripts scripts Dockerfile Dockerfile LICENSE LICENSE README.md README.md View all files Repository files navigation
Self-hosted Kubernetes Lab for Hands-on Learning
KubeKosh runs a real K3s Kubernetes cluster inside a single Docker container and pairs it with a browser-based terminal and automated scenario validation — no cloud account or local cluster required.
docker run -itd --name kubekosh --privileged -p 7554:80 zeborg/kubekosh:latest
Open http://localhost:7554 — wait ~30 seconds for the Cluster Ready indicator to turn green.
--privileged is required — K3s needs access to kernel namespaces and cgroups.
Security Warning: Do not expose this container publicly. Use it only on your local machine as it is meant for educational purposes only.
docker run -itd --name kubekosh --privileged -p 7554:80 \
-v < your_custom_directory > :/data zeborg/kubekosh:latest
Progress is stored in SQLite at /data/progress.db inside the container. Mount a local directory to /data to keep progress across container restarts.
docker build -t kubekosh .
# multi-platform
docker buildx build --platform linux/amd64,linux/arm64 -t kubekosh .
What's Inside
Bundle
Focus
Exam Mode
🌱 Kubernetes Basics
Core concepts
60 min
🧑‍ ✈️ Kubernetes Administrator
CKA
120 min
🛠️ Kubernetes Developer
CKAD
120 min
🛡️ Kubernetes Security
CKS
120 min
Scenario types:
Task — Hands-on challenge in the live terminal. Click Validate for automated cluster-state checking.
MCQ — Multiple-choice question with a detailed explanation on submission.
The terminal comes pre-configured with:
Component
Technology
Frontend
React + Vite, xterm.js
Backend
Node.js / Express, node-pty WebSocket PTY
Cluster
K3s (single-node, in-container)
Proxy
nginx on container port 80 , mapped to host port 7554
Storage
SQLite ( better-sqlite3 ) at /data/progress.db
Everything runs inside a single Docker image managed by scripts/entrypoint.sh .
scenarios/
├── data/ # One JSON file per scenario -> <scenario-id>.json
├── bundles/ # One JSON file per bundle -> <bundle-id>.json
└── SCHEMA.md # Full schema reference
backend/
└── server.js # Express API + WebSocket PTY
frontend/
└── src/ # React + Vite SPA
scripts/
├── entrypoint.sh # Container startup (k3s -> API -> nginx)
└── nginx.conf # Reverse-proxy config
Contributing
Contributions are what make open-source projects like this one grow — and every contribution counts, big or small. Whether you're fixing a typo, polishing a scenario description, or building a completely new exercise from scratch, you're helping the next person learn Kubernetes in the best way possible. Thank you for taking the time!
Each scenario is a single JSON file in scenarios/data directory; each bundle is a single JSON file in scenarios/bundles directory. See scenarios/SCHEMA.md for the full schema.
validation.commands — idempotent kubectl commands only
setup_commands / teardown_commands — kubectl or native Ubuntu commands only
correct_option must match one of the options[].id values
In-Memory Cache & Hot Reloading
To ensure high performance and zero disk-I/O bottlenecking, scenarios and bundles are cached in memory on backend startup. When developing or updating scenarios, you can hot-reload the definitions without rebuilding the image or restarting the container:
Mount Scenarios Directory: Run the container with the local scenarios/ directory mounted to /app/scenarios :
docker run --rm -itd --privileged -p 7554:80 --name kubekosh -v < path_to_scenarios_directory > :/app/scenarios zeborg/kubekosh:latest
Reload Cache: Click the Reload Scenario Cache (↻) button in the top right corner of the header in the web user interface, or send an API request:
curl -X POST http://localhost:7554/api/cache/reload
NOTE:
The content in <path_to_scenarios_directory> should be the path to the local scenarios/ directory of the cloned repository with your updates, i.e., it should contain the updated scenarios/data and scenarios/bundles directories.
# 1. Fork the repo on GitHub, then clone your fork
git clone https://github.com/ < your-username > /kubekosh.git
cd kubekosh
# 2. Create a branch
git checkout -b feat/my-scenario
# 3. Add a new scenario file (copy an existing scenario as a template or create a new one)
cp scenarios/data/deploy-nginx.json scenarios/data/my-new-scenario.json
vim scenarios/data/my-new-scenario.json # edit the new scenario as per [SCHEMA.md](scenarios/SCHEMA.md)
# 4. Add the scenario ID to the relevant bundle
vim scenarios/bundles/k8s-basics.json # edit the bundle to include the new scenario ID
# 5. Build and test locally
# Run the built container directly:
docker build -t kubekosh . && docker run --rm -itd --privileged -p 7554:80 --name kubekosh kubekosh
# Or mount the scenarios folder for hot-reloading:
docker run --rm -itd --privileged -p 7554:80 -v $PWD /scenarios:/app/scenarios --name kubekosh zeborg/kubekosh:dev
# 6. Commit and push to your fork (example for adding `my-new-scenario` to `k8s-basics` bundle)
git add scenarios/data/my-new-scenario.json scenarios/bundles/k8s-basics.json
git commit -m " feat: add my-new-scenario to k8s-basics bundle "
git push -u origin feat/my-new-scenario
Open a Pull Request from your fork's branch against main .
Apache 2.0 License — see LICENSE .
Interactive Kubernetes Playground
There was an error while loading. Please reload this page .
32
forks
Report repository
Releases
7
Release v0.2.2
Latest
Jun 18, 2026
+ 6 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
