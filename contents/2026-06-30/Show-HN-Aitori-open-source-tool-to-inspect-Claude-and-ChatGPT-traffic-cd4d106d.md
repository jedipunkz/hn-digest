---
source: "https://github.com/truefoundry/aitori"
hn_url: "https://news.ycombinator.com/item?id=48733643"
title: "Show HN: Aitori – open-source tool to inspect Claude and ChatGPT traffic"
article_title: "GitHub - truefoundry/aitori · GitHub"
author: "deeptishukla22"
captured_at: "2026-06-30T15:50:24Z"
capture_tool: "hn-digest"
hn_id: 48733643
score: 1
comments: 0
posted_at: "2026-06-30T14:58:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Aitori – open-source tool to inspect Claude and ChatGPT traffic

- HN: [48733643](https://news.ycombinator.com/item?id=48733643)
- Source: [github.com](https://github.com/truefoundry/aitori)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T14:58:11Z

## Translation

タイトル: Show HN: Aitori – クロードと ChatGPT トラフィックを検査するオープンソース ツール
記事タイトル: GitHub - truefoundry/aitori · GitHub
説明: GitHub でアカウントを作成して、truefoundry/aitori の開発に貢献します。

記事本文:
GitHub - truefoundry/aitori · GitHub
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
トゥルーファウンドリ
/
あいとり
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows cmd/ aitori cmd/ aitori configs configs docs docs external 内部テスト/ 統合テスト/ 統合ツール/ aitori-gateway tools/ aitori-gateway .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
aitori はマシン上で実行され、AI トラフィックを傍受する小さなエージェントです
それを残します。アプリやコマンドライン ツールからのモデル呼び出しと MCP 呼び出しは、
ゲートウェイ。ログを記録し、ポリシーと照合してチェックできます。他のすべて
放っておかれます。アプリケーションは影響を受けません。同じエンドポイントと通信し、
同じキーを入力すると、同じ応答が返されます。
多くのアプリでは、リクエストの送信先を変更する方法がありません。上のブラウザ
claude.ai または chatgpt.com には、ゲートウェイを指す設定がありません。
そのトラフィックを傍受できる唯一の場所は、そのトラフィックの前にあるマシン自体上にあります。
葉。 aitori は、リストしたホストのみを復号化し、残りはそのままにしておきます。
ゲートウェイは利用できませんが、リクエストは依然として元の宛先に到達します。
フローチャート LR
App(["App / CLI"]) -->|HTTPS| ET["アイトリ<br/>(デバイス上)"]
ET -->|"AI コール<br/>(LLM / MCP)"| GW["AI ゲートウェイ<br/>ログ · 予算 · ポリシー"]
ET -->|"その他すべて"| UP[("実際の上流")]
GW -->|進む|上へ
classDef gw fill:#dbeafe、ストローク:#2563eb、color:#1e3a8a;
クラス GW GW;
読み込み中
リストされたホストの 1 つに対するリクエストは復号化され、検査されます。モデルなら
または MCP コールの場合、aitori はそれをゲートウェイに転送し、ゲートウェイがそれを記録して渡します。
本来の目的地へ。それ以外は何でもまっすぐ進みます

ここ。リクエストの仕組み
機密扱い、ブロック アクション、および応答がどのように返されるかについては、「」で説明されています。
docs/architecture.md 。
バイナリをインストールします (macOS/Linux):
カール -fsSL https://raw.githubusercontent.com/truefoundry/aitori/main/install.sh |しー
後で aitori を削除するには、「アンインストール (元に戻す)」を参照してください。
削除する前に sudo aitori down と sudo aitori ca Remove を実行するとシステムが変更されます
バイナリ)。
次に、このマシンを管理し、トラフィックをライブで監視します (ゲートウェイも構成も必要ありません)。
sudo アイトリアップ --ui
これにより、デバイスごとの CA がインストールされ、システムが aitori を指すようになります。内蔵されている
プロファイルはすでに Claude (コード、デスクトップ、Web) と ChatGPT をカバーしているため、それらの呼び出しは
復号化され、分類され、ゲートウェイが設定されていない状態で直接通過します
本当の上流に接続します (何も壊れません)。ライブビューを開きます。
Claude または ChatGPT を使用すると、通話は次のようにストリーミングされます。
sudo aitori down # システムプロキシを元に戻し、設定の挿入を元に戻します
Ctrl-C はクリーン終了時にもこれを実行しますが、次の場合は常に sudo aitori down を実行してください。
プロセスが強制終了されたか、ターミナルを閉じただけです。それ以外の場合はシステム プロキシ
停止したアイトリと交通中断を指し続けます。 (デバイスごとの CA は残ります)
インストールされています。 sudo aitori ca Remove で削除してください。)
それがデモです。トラフィックをただ監視するのではなく、アクションを起こすには、aitori をポイントしてください
ゲートウェイで（下）。ソースからビルドするには、 docs/development.md を参照してください。
ゲートウェイがない場合、aitori はトラフィックを検査し、変更せずに通過させます。
それに基づいて行動するには、aitori を AI ゲートウェイに向けます: setgateway.url (構成または
--gateway-url を使用して)、ゲートウェイ トークンを次の名前のファイルに置きます。
ゲートウェイ.auth.token_file 。モデル呼び出しと MCP 呼び出しは代わりにゲートウェイを経由します。
プロバイダーに直接連絡します。 docs/gateway.md には、
ゲートウェイは実装する必要があります。 TrueFoundry AI Gateway の場合は、正確な URL とトークン

手順は docs/truefoundry_gateway.md にあります。
組み込みプロファイルはすでに共通アプリを管理しているため、aitori up は
設定がまったくありません。構成ファイルは独自のホスト、ゲートウェイ、またはライブ UI を追加します。
最小のものは次のとおりです。
バージョン : 1
ウイ:
有効 : true # http://127.0.0.1:9100 でのライブトラフィックビュー
intercept_hosts :
- api.example.com # このホストも復号化して管理します
configs/demo.yaml はゲートウェイなしのデモです。
configs/conversations.yaml は完全にコメント化されたものです
たとえば。 docs/configuration.md が完全なリファレンスです。の
組み込みのホスト/パス ルールが開始点です - 実際のキャプチャに対してルールを確認します
彼らに頼る前に。
これまでにエンドツーエンドで検証したこと。 ✅ = 検証済みのキャプチャ。 — = まだ
検証済み (必ずしもサポートされていないわけではありません)。他のホストやアプリは設定経由で動作しますが、
エンドツーエンドで行使されていない。 docs/roadmap.md は内容を追跡します
完了したものと開いているもの。
aitori はマシンが連携していることを前提としています。ローカル管理者がいる人
権利はそれを回避することができます。これはガバナンス ツールであり、サンドボックスではありません。生成される CA
はデバイスに固有であり、その秘密キーがマシンから離れることはありません。
フリート全体で共有キーではありません。復号化するホストは構成ファイルにリストされます。
aitori status で表示されます。
docs/getting-started.md — ライブ UI のインストール、実行、ゲートウェイへの接続
docs/development.md — ソース、モックゲートウェイ、およびテストスイートからビルドします
docs/configuration.md — 構成リファレンス、CLI、設定インジェクション、透過的キャプチャ
docs/architecture.md — aitori の詳細な仕組み
docs/gateway.md — ゲートウェイへの接続とリルートコントラクト
docs/roadmap.md — ステータス、検証済みのプラットフォーム、および未解決の作業
AGENTS.md — 寄稿者向けのアーキテクチャと規約
aitori ( A.I.-tor-ee ) は ai + tori — 鳥居 torii は t の門です

彼は
日本の神社の入り口。全体的なアイデアは次のとおりです。aitori はすべての AI のゲートです
コールは発信中に通過します。 🪧
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 25 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to truefoundry/aitori development by creating an account on GitHub.

GitHub - truefoundry/aitori · GitHub
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
truefoundry
/
aitori
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows cmd/ aitori cmd/ aitori configs configs docs docs internal internal test/ integration test/ integration tools/ aitori-gateway tools/ aitori-gateway .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
aitori is a small agent that runs on a machine and intercepts the AI traffic
leaving it. Model and MCP calls from apps and command-line tools are sent through
a gateway, where they can be logged and checked against a policy; everything else
is left alone. The application is unaffected: it talks to the same endpoint, uses
the same key, and gets the same response back.
Many apps give you no way to change where they send their requests. A browser on
claude.ai or chatgpt.com has no setting that points it at a gateway.
The only place left to intercept that traffic is on the machine itself, before it
leaves. aitori decrypts only the hosts you list and leaves the rest alone, and if
the gateway is unavailable the request still reaches its original destination.
flowchart LR
App(["App / CLI"]) -->|HTTPS| ET["aitori<br/>(on device)"]
ET -->|"AI calls<br/>(LLM / MCP)"| GW["AI gateway<br/>log · budget · policy"]
ET -->|"everything else"| UP[("real upstream")]
GW -->|forward| UP
classDef gw fill:#dbeafe,stroke:#2563eb,color:#1e3a8a;
class GW gw;
Loading
A request to one of the listed hosts is decrypted and inspected. If it is a model
or MCP call, aitori forwards it to the gateway, which records it and passes it on
to the original destination. Anything else goes straight there. How requests are
classified, the block action, and how responses travel back are described in
docs/architecture.md .
Install the binary (macOS/Linux):
curl -fsSL https://raw.githubusercontent.com/truefoundry/aitori/main/install.sh | sh
To remove aitori later, see Uninstall (revert
system changes with sudo aitori down and sudo aitori ca remove before deleting
the binaries).
Then govern this machine and watch the traffic live — no gateway, no config:
sudo aitori up --ui
This installs a per-device CA and points the system at aitori. The built-in
profiles already cover Claude (Code, Desktop, web) and ChatGPT, so their calls are
decrypted, classified, and — with no gateway configured — passed straight through
to their real upstream (nothing breaks). Open the live view:
Use Claude or ChatGPT, and the calls stream in:
sudo aitori down # reverts the system proxy + undoes settings injection
Ctrl-C does this too on a clean exit, but always run sudo aitori down if the
process was killed or you just closed the terminal — otherwise the system proxy
stays pointed at a stopped aitori and traffic breaks. (The per-device CA is left
installed; remove it with sudo aitori ca remove .)
That is the demo. To act on the traffic rather than only watch it, point aitori
at a gateway (below). To build from source, see docs/development.md .
Without a gateway, aitori inspects the traffic and passes it through unchanged.
To act on it, point aitori at an AI gateway: set gateway.url (in the config or
with --gateway-url ) and put the gateway token in the file named by
gateway.auth.token_file . Model and MCP calls then go through the gateway instead
of straight to the provider. docs/gateway.md describes what a
gateway has to implement; for the TrueFoundry AI Gateway, the exact URL and token
steps are in docs/truefoundry_gateway.md .
The built-in profiles already govern the common apps, so aitori up works with
no config at all. A config file adds your own hosts, a gateway, or the live UI —
the smallest one is just:
version : 1
ui :
enabled : true # live-traffic view at http://127.0.0.1:9100
intercept_hosts :
- api.example.com # decrypt + govern this host too
configs/demo.yaml is the no-gateway demo;
configs/conversations.yaml is the fully-commented
example. docs/configuration.md is the full reference. The
built-in host/path rules are starting points — confirm them against a real capture
before relying on them.
What we've verified end-to-end so far. ✅ = validated capture; — = not yet
validated (not necessarily unsupported). Other hosts and apps work via config but
haven't been exercised end-to-end. docs/roadmap.md tracks what's
done and what's open.
aitori assumes the machine is cooperating. Someone with local administrator
rights can bypass it; it is a governance tool, not a sandbox. The CA it generates
is specific to the device and its private key never leaves the machine, so there
is no shared key across a fleet. The hosts it decrypts are listed in the config
and shown by aitori status .
docs/getting-started.md — install, run, the live UI, and connecting to a gateway
docs/development.md — build from source, the mock gateway, and the test suite
docs/configuration.md — config reference, CLI, settings injection, transparent capture
docs/architecture.md — how aitori works in depth
docs/gateway.md — connecting to a gateway and the reroute contract
docs/roadmap.md — status, validated platforms, and open work
AGENTS.md — architecture and conventions, for contributors
aitori ( A.I.-tor-ee ) is ai + tori — and 鳥居 torii is the gate at the
entrance to a Japanese shrine. Which is the whole idea: aitori is the gate every AI
call passes through on its way out. 🪧
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 25, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
