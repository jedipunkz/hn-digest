---
source: "https://github.com/sinameraji/hotcell"
hn_url: "https://news.ycombinator.com/item?id=49082676"
title: "Show HN: Hotcell – local sandboxes for AI agents"
article_title: "GitHub - sinameraji/hotcell: Self-hostable sandbox SDK inspired by Cloudflare Sandbox SDK. works on any device · GitHub"
author: "sinameraji"
captured_at: "2026-07-28T12:48:30Z"
capture_tool: "hn-digest"
hn_id: 49082676
score: 2
comments: 0
posted_at: "2026-07-28T12:17:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hotcell – local sandboxes for AI agents

- HN: [49082676](https://news.ycombinator.com/item?id=49082676)
- Source: [github.com](https://github.com/sinameraji/hotcell)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T12:17:13Z

## Translation

タイトル: Show HN: Hotcell – AI エージェント用のローカル サンドボックス
記事タイトル: GitHub - sinamerji/hotcell: Cloudflare サンドボックス SDK からインスピレーションを得た自己ホスト可能なサンドボックス SDK。どのデバイスでも動作します · GitHub
説明: Cloudflare サンドボックス SDK からインスピレーションを得た自己ホスト可能なサンドボックス SDK。どのデバイスでも動作します - シナラジ/ホットセル
HN テキスト: こんにちは、HN!ここのシーナ。ここ 1 か月ほど取り組んできたこのオープンソース プロジェクト (Apache 2.0) を共有したいと思います。これはホットセルと呼ばれ、任意のデバイス (ラップトップ、Linux VM、ベアメタル) 上でサンドボックスを作成/一時停止/管理できます。私は computesdk ( https://www.computesdk.com/benchmarks/sandboxes/ ) で見つけたさまざまなベンチマークと、最近 Twitter で dax (オープン コードの共同創設者) が共有したベンチマーク (彼の元のスレッドと私の返信へのリンク https://x.com/sinasanm/status/2078158598996443346 ) および一般に完全なベンチマークに対してテストしました。結果はここにあります: https://github.com/sinameraji/hotcell/blob/main/docs/benchma... 私は開発者のエクスペリエンスとエージェントのエクスペリエンスの両方がスムーズであることを確認するために多くの時間を費やし、ベンチマークに対してパフォーマンスを向上させるために積極的に取り組んでいます。また、API キーがサンドボックスに直接挿入されることはありません。代わりに、サンドボックスが終了すると役に立たなくなるサンドボックスごとのトークンが作成されます。分離方法としては、docker、apple VZ (Mac の vm グレード分離)、または Linux の firecracker で動作します。私が想像する主な使用例: * ローカルでファイルシステムにアクセスできるエージェントニックデスクトップアプリを構築する場合、ホットセルを使用すると、サンドボックスを起動して必要なファイルをそこに持ち込み、新しいブランチで作業を行い、PR を作成してサンドボックスを閉じることが非常に簡単になります。 * クロードコードなどを使用するだけで、ワークツリーを使用したりプロジェクトディレクトリをクローンしたりせずに 5 ～ 6 個の差分環境をスピンアップしたい場合は、1 行を書くことができます

このようなコマンドを実行すると、即座に 5 つのサンドボックスが作成されます。
hotcell create -n 5 --name feat --branch auto --opencode --repo https://github.com/you/app 私のインスピレーションは、cloudflare サンドボックス SDK でした (しかし、友人に気づき、オンプレミス環境では Cloudflare を使用できないので、どのハードウェアでも動作するローカル/オープンソースのものを構築しようと思いました)。ありがとう、気に入っていただければ幸いです。

記事本文:
GitHub - sinameraji/hotcell: Cloudflare サンドボックス SDK からインスピレーションを得た自己ホスト可能なサンドボックス SDK。どのデバイスでも動作します · GitHub
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
シナメラジ
/
ホットセル
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
253 コミット 253 コミット .github .github エージェント エージェント デモ デモ ドキュメント ドキュメント 証拠 証拠 例 例 ヘルパー/ hotcell-vz ヘルパー/ hotcell-vz イメージ/ ベース イメージ/ ベース パッケージ パッケージ スクリプト スクリプト sdk/ python sdk/ python .gitignore .gitignore CLAUDE.md CLAUDE.md KIMI.md KIMI.md ライセンス ライセンス通知 通知 PUBLISHING.md PUBLISHING.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
独自のハードウェア上にある AI エージェント用のサンドボックス。机上の Mac Mini、クラウド VM、ベアメタル ボックスなど、保持できる限り多くの独立したエージェント サンドボックスがあり、API キーがそれらのいずれにも入力されることはありません。
ベンチマーク (完全な結果) — シーケンシャル TTI、スタッガード TTI、バースト TTI、ウォーム プール採用、サンドボックスダックスの実際のワークロード — ComputeSDK ハーネスと、ベア メタル上の dax の OpenCode ベンチマークで測定 (生の証拠)。
npm i -g ホットセル
hotcell # 最初の実行: 30 秒のガイド付きセットアップ — その後、ライブ フリート
デフォルト以外のソケット (colima、OrbStack、podman) 上の Docker ランタイム?ホットセルの読み取り
Docker_HOST 、Docker CLI のコンテキストではありません — 例:コリマの場合:
ホットセルを開始する前に DOCKER_HOST=unix://$HOME/.colima/default/docker.sock をエクスポートします。
ホットセル キーは openrouter # 任意のプロバイダーのキーを追加します — ホスト上に留まり、サンドボックス内に存在することはありません
ホットセル キー import .env # または .env を一括インポート (OPENAI_API_KEY → openai, …)
hotcell create -n 5 --repo https://github.com/you/app --branch # 5 つの分離されたセル、それぞれが独自のブランチ上にあります
hotcell ターミナル < id > # セル内のシェル
hotcell run --setup " pip install ruff " " ruff check . " # ワンショット: 作成 → 実行 → 破棄
ほー

tcell rm --all # すべてがなくなりました。あなたのリポジトリはそのままです
1 つのリポジトリに 5 つのエージェント
hotcell create -n 5 --name feat --branch auto --opencode \
--repo https://github.com/you/app
5 つの分離されたセル。それぞれクローンされたリポジトリ、独自のブランチ ( feat-1 … feat-5 )、およびプリインストールされゲートウェイを指す OpenCode ( --opencode は --egress を意味します) を持ちます。プロビジョニング中のセルごとの進行状況をライブで確認できます。セットアップが遅い場合でも、誤ってタイムアウトすることはありません。セルごとにターミナルを開き、各エージェントを異なる機能に配置します。
hotcell ターミナル < id > # inside: cd app && opencode
hotcell rm --all # 完了 — 5 つのセルがなくなり、リポジトリはそのままです
OpenRouter キーと GitHub キーはホスト上に残ります。 --egress を使用すると、セルは有効期間の短いセルごとのトークンのみを保持します。LLM 呼び出しはホットセルのゲートウェイを経由し、各セルの git オリジンは自動的にそこを経由して接続されるため、git Push はすぐにキーレスで機能します。 (OpenCode 固有の配線 + pr ヘルパー: example/agents.sh 。)
鍵が外に出たままになります。サンドボックスは、サンドボックスごとのトークンを実際のキー (従量制、支出制限付き、取り消し可能) と交換するゲートウェイを介して LLM と GitHub に到達します。オプションのデフォルト拒否の出力 (Linux および microVM ではカーネルによって適用されます。macOS Docker ではアドバイザリー)。
ハードウェアが許す限り。 1 つのデーモン、ライブ CPU/メモリ/サンドボックスごとのコスト、ボックスを OOM する代わりにオーバーサブスクライブを拒否するアドミッション コントロール。
コンテナまたは microVM — あらゆる場所の Docker、VM グレードの分離のための Firecracker (Linux/KVM)、および Apple VZ (macOS) がすべて 1 つのインターフェイスの背後にあります。
ゲートウェイが保護できるものと保護できないもの
ホットセル キー import .env では、すべての変数をゲートウェイに設定するか (実際のキーはホスト上にあり、サンドボックスはサンドボックスごとのトークンを取得します)、注入するか (実際の値はすべてのサンドボックスにコピーされます)、またはスキップするように求められます。 hotcell はどれがどれであるかを推測しません。名前からはどちらであるか判断できません。

値は秘密であり、間違った答えは両方とも沈黙します。確認ステップでは、何かが保存される前に、サンドボックスに入るすべての値がリストされます。
保護済み — 資格情報が HTTP リクエスト ヘッダーで送信される任意の API。 OpenAI、Anthropic、OpenRouter、GitHub、Stripe、Slack、Cloudflare、およびほとんどの REST API は同様に機能します。 LLM のみのカーブアウトはありません。 5 つの組み込みルートを超えるプロバイダーには、ベース URL と認証ヘッダーを一度 .hotcell/env.json に保存する必要があります (シークレットはありません。コミットすると、チームメイトのクローンがすべての決定を継承します)。
保護されていません — inject を選択すると、実際の値がサンドボックスに入ります。
非 HTTP プロトコル (Postgres、Redis、MongoDB) — 認証情報はワイヤ ハンドシェイク内に組み込まれるため、置き換えるヘッダーはありません。これらのゲートウェイは到達できません: ベース URL プロンプトが postgres:// を拒否します。
リクエスト署名スキーム (AWS SigV4) — キーはローカルでリクエストに署名し、送信されることはないため、交換するものはありません。
mTLS / クライアント証明書、および OAuth リフレッシュ トークン交換。
保護されていますが、ドロップインではありません — ホストをハードコーディングする SDK は、挿入された <PROVIDER>_BASE_URL を無視し、プロバイダーを直接呼び出します。 LLM SDK はそれを読み取ります。他の多くはコード内でホストを受け取り、それを渡すのに 1 行が必要です。 HOTCELL_EGRESS_ENFORCE では、これらの直接呼び出しは、ゲートウェイをループから静かに外すのではなく、拒否されます。
名前は変わる可能性があります。ゲートウェイ変数はそのルート名でサンドボックスに到達します。つまり、GH_TOKEN を github ルートとしてインポートすると、サンドボックスは GITHUB_API_KEY を認識します。レビュー画面には各行のルートが表示されます。
ドキュメント: ガイド · エグレスとキー · すべてのコマンドと設定 · Linux セルフホスティング · ベンチマーク · Python SDK: pip install hotcell
Cloudflare サンドボックス SDK からインスピレーションを得た自己ホスト可能なサンドボックス SDK。どのデバイスでも動作します
Readme Apache-2.0 ライセンス法

アクティビティスターズ
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hostable sandbox SDK inspired by Cloudflare Sandbox SDK. works on any device - sinameraji/hotcell

Hi HN! Sina here. i wanted to share this open source project (apache 2.0) that i've been working on for the past month or so. it's called hotcell and it lets you create/pause/manage sandboxes on any device (your laptop, linux vm, bare metal). i've tested it against various benchmarks that i found on computesdk ( https://www.computesdk.com/benchmarks/sandboxes/ ) as well as the one shared by dax (cofounder of open code) on twitter recently (link to his original thread and my reply https://x.com/sinasanm/status/2078158598996443346 ) and generally the full benchmark results are here: https://github.com/sinameraji/hotcell/blob/main/docs/benchma... i spent a lot of time on making sure both the devrloper experience and agent experience are smooth on it and im actively working to make it perform better against benchmarks. also api keys are never directly injected into the sandbox. instead it creates a per-sandbox token that becomes useless when the sandbox dies. as for isolation method, it can work with docker, apple VZ (vm grade isolation on mac) or firecracker for linux . primary use cases i imagine: * if you build agentic desktop apps that have filesystem access locally, hotcell makes it super easy to start a sandbox and bring whatever file needed to it, do the work in a new branch, create a PR and close the sandbox * if u just use claude code etc. and wanna spin up 5-6 diff environments without using worktree or cloning your project dirs, u can write a 1 line command like this to instantly have 5 sandboxes!
hotcell create -n 5 --name feat --branch auto --opencode --repo https://github.com/you/app my inspiration was cloudflare sandbox sdk (but i noticed my friends and i couldn't use cloudflare in on-prem environment so i thought i'd build a local / open source one that works w any hardwarw). thank you and hope you like it.

GitHub - sinameraji/hotcell: Self-hostable sandbox SDK inspired by Cloudflare Sandbox SDK. works on any device · GitHub
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
sinameraji
/
hotcell
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
253 Commits 253 Commits .github .github agent agent demo demo docs docs evidence evidence examples examples helpers/ hotcell-vz helpers/ hotcell-vz images/ base images/ base packages packages scripts scripts sdk/ python sdk/ python .gitignore .gitignore CLAUDE.md CLAUDE.md KIMI.md KIMI.md LICENSE LICENSE NOTICE NOTICE PUBLISHING.md PUBLISHING.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Sandboxes for AI agents, on your own hardware. A Mac Mini on your desk, a cloud VM, a bare-metal box — as many isolated agent sandboxes as it can hold, and your API keys never enter any of them.
Benchmarks ( full results ) — sequential TTI · staggered TTI · burst TTI · warm-pool adopt · sandbox-dax real workload — measured with the ComputeSDK harness , plus dax's OpenCode benchmark on bare metal ( raw evidence ).
npm i -g hotcell
hotcell # first run: 30-second guided setup — then your live fleet
Docker runtime on a non-default socket (colima, OrbStack, podman)? hotcell reads
DOCKER_HOST , not the docker CLI's context — e.g. for colima:
export DOCKER_HOST=unix://$HOME/.colima/default/docker.sock before hotcell start .
hotcell keys add openrouter # any provider's key — stays on the host, never in a sandbox
hotcell keys import .env # or bulk-import a .env (OPENAI_API_KEY → openai, …)
hotcell create -n 5 --repo https://github.com/you/app --branch # five isolated cells, each on its own branch
hotcell terminal < id > # shell inside a cell
hotcell run --setup " pip install ruff " " ruff check . " # one-shot: create → run → destroy
hotcell rm --all # everything gone; your repo untouched
Five agents on one repo
hotcell create -n 5 --name feat --branch auto --opencode \
--repo https://github.com/you/app
Five isolated cells, each with the repo cloned, its own branch ( feat-1 … feat-5 ), and OpenCode preinstalled and pointed at the gateway ( --opencode implies --egress ). Live per-cell progress while they provision — slow setups can't falsely time out. Open a terminal per cell and put each agent on a different feature:
hotcell terminal < id > # inside: cd app && opencode
hotcell rm --all # done — five cells gone, your repo untouched
Your OpenRouter and GitHub keys stay on the host: with --egress , a cell only ever holds a short-lived per-cell token — LLM calls go through hotcell's gateway, and each cell's git origin is wired through it automatically, so git push works keylessly out of the box . (OpenCode-specific wiring + a pr helper: examples/agents.sh .)
Keys stay out. Sandboxes reach LLMs and GitHub through a gateway that swaps a per-sandbox token for the real key — metered, spend-capped, revocable. Optional default-deny egress (kernel-enforced on Linux and microVMs; advisory on macOS Docker).
As many as the hardware allows. One daemon, live CPU/mem/cost per sandbox, admission control that refuses to over-subscribe instead of OOM-ing the box.
Containers or microVMs — Docker everywhere, Firecracker (Linux/KVM) and Apple VZ (macOS) for VM-grade isolation, all behind one interface.
What the gateway can and can't protect
hotcell keys import .env asks you to set every variable to gateway (the real key stays on the host; the sandbox gets a per-sandbox token), inject (the real value is copied into every sandbox), or skip . hotcell does not guess which is which — a name cannot tell you whether a value is a secret, and both wrong answers are silent. The confirm step lists every value that will enter a sandbox before anything is stored.
Protected — any API whose credential travels in an HTTP request header. OpenAI, Anthropic, OpenRouter, GitHub, Stripe, Slack, Cloudflare and most REST APIs work identically; there is no LLM-only carve-out. Providers beyond the five built-in routes need their base URL and auth header once, saved to .hotcell/env.json (no secrets — commit it, and a teammate's clone inherits every decision).
Not protected — the real value enters the sandbox if you choose inject :
Non-HTTP protocols (Postgres, Redis, MongoDB) — the credential rides inside the wire handshake, so there is no header to substitute. gateway is unreachable for these: the base-URL prompt rejects postgres:// .
Request-signing schemes (AWS SigV4) — the key signs the request locally and is never transmitted, so there is nothing to swap.
mTLS / client certificates , and OAuth refresh-token exchanges.
Protected, but not drop-in — an SDK that hardcodes its host ignores the injected <PROVIDER>_BASE_URL and calls the provider directly. LLM SDKs read it; many others take a host in code and need one line to pass it through. Under HOTCELL_EGRESS_ENFORCE those direct calls are denied rather than quietly leaving the gateway out of the loop.
Names can shift. A gateway variable reaches the sandbox under its route's name — import GH_TOKEN as the github route and the sandbox sees GITHUB_API_KEY . The review screen shows the route for each row.
Docs: guide · egress & keys · every command & config · Linux self-hosting · benchmarks · Python SDK: pip install hotcell
Self-hostable sandbox SDK inspired by Cloudflare Sandbox SDK. works on any device
Readme Apache-2.0 license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
