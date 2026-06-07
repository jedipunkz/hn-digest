---
source: "https://github.com/anthropics/claude-code/issues/65697"
hn_url: "https://news.ycombinator.com/item?id=48434436"
title: "Anthropic, please ship an official Claude Desktop for Linux"
article_title: "[FEATURE] Official Claude Desktop build for Linux (Ubuntu LTS / Debian) · Issue #65697 · anthropics/claude-code · GitHub"
author: "predkambrij"
captured_at: "2026-06-07T15:35:46Z"
capture_tool: "hn-digest"
hn_id: 48434436
score: 169
comments: 68
posted_at: "2026-06-07T13:06:52Z"
tags:
  - hacker-news
  - translated
---

# Anthropic, please ship an official Claude Desktop for Linux

- HN: [48434436](https://news.ycombinator.com/item?id=48434436)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/65697)
- Score: 169
- Comments: 68
- Posted: 2026-06-07T13:06:52Z

## Translation

タイトル: Anthropic、Linux 用の公式クロード デスクトップを出荷してください
記事のタイトル: [特集] Linux 用公式 Claude デスクトップ ビルド (Ubuntu LTS / Debian) · 問題 #65697 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存のリクエストを検索しましたが、この機能はまだリクエストされていません これは 1 つの機能リクエストです (複数の機能ではありません) 問題ステートメントのプリフライト メモです。最も近い未解決の問題は #40347 です。関連: #47316...

記事本文:
[特集] Linux 用公式 Claude デスクトップ ビルド (Ubuntu LTS / Debian) · 問題 #65697 · anthropics/claude-code · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
人類学
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
[特集] Linux 用公式クロード デスクトップ ビルド (Ubuntu LTS / Debian) #65697
リンクをコピーする 新しい問題 リンクをコピー 開く 開く [機能] Linux 用公式クロード デスクトップ ビルド (Ubuntu LTS / Debian) #65697 リンクをコピー ラベル領域:デスクトップの強化 新機能またはリクエスト 新機能またはリクエスト プラットフォーム:linux Linux で特に発生する問題 問題は Linux で特に発生 説明
発行本体アクションのプリフライト チェックリスト
既存のリクエストを検索しましたが、この機能はまだリクエストされていません
これは単一の機能リクエストです (複数の機能ではありません)
飛行前のメモ。最も近い未解決の問題は #40347 です。関連: #47316 (閉鎖)、#38276 (このリポジトリの範囲外として閉鎖)、#36011 (古い)。私はこれを #40347 の統合および拡張として提出し、技術的な枠組み (デスクトップ拡張に対するクロード コード プラグイン開発) を修正し、Cowork Linux-VM アーキテクチャのプライマリ ソースと名付け、現在の市場データを追加します。メンテナーが希望する場合は、喜んで #40347 にマージします。別の会場が正しい場合は、閉じるのではなくルーティングしてください。
範囲について: この問題は 2 つの具体的な方法でクロード コードに関係します。 (1) Claude Code プラグインは、Linux ビルドを持たない Claude Desktop 拡張機能に対して開発およびテストされているため、現在プラグインの作業には OS を切り替える必要があります。 (2) Cowork は macOS 上の Linux VM 内で Claude Code バイナリを呼び出します。そのため、Linux 実行パスは Claude Code 製品内にすでに存在しており、実際には公開ターゲットとして欠落しています。
この問題で求められていること
Linux デスクトップのサポートに関する Anthropic の公的立場。理想的にはファーストパーティ ビルドです。 「現在のロードマップには載っていない。その理由はここにある」という理由を説明すれば、この問題のほとんどが解決します。私の知る限り、Linux デスクトップのサポートに関する公式声明はありません。不在の私

それ自体が問題の一部です。
Anthropic は、macOS と Windows 向けにのみ Claude Desktop を配布しています。公式ダウンロードページには「Linux では利用できません」と記載されています。 Claude Code (CLI) は Linux 上でネイティブに実行されますが、これはターミナル ツールであり、デスクトップ GUI の代替ではありません。デスクトップ拡張機能 (Claude Code プラグインがテストされるサーフェス)、コンピューターの使用、デスクトップ ディクテーション、および Cowork は、Claude Desktop でのみ利用できます。したがって、Linux ユーザーには、これらの機能への公式にサポートされているグラフィカル パスがなく、特に、macOS または Windows に切り替えることなくデスクトップ拡張機能としてクロード コード プラグインを開発およびテストする方法がありません。
なぜこれを正当化することが構造的に難しいのか
Anthropic はすでに Linux ソフトウェアを構築、署名、配布しています。 code.claude.com/docs/en/setup ごとに、Claude Code には署名付きの apt、dnf、apk リポジトリとアーキテクチャごとのバイナリ (linux-x64、linux-arm64、musl バリアント) が同梱されています。パイプラインは存在します。
Cowork エージェントは、製品内ですでに Linux に依存しています。発売日 (2026 年 12 月 1 日) の Simon Willison による独立したリバース エンジニアリングと、Pluto Security および pvieito (「Claude Cowork の内部」) による裏付けにより、macOS 上で Cowork が Apple の仮想化フレームワーク (VZVirtualMachine) 経由でカスタム Ubuntu 22.04 VM を起動し、その内部で bubblewrap および seccomp の下でクロード コード バイナリを実行していることが判明しました。 Anthropic 自身のドキュメントでは、ハイパーバイザーの分割 (macOS 上の Apple Virtualization.framework、Windows 上の Hyper-V) を確認しています。コミュニティ プロジェクト johnzfitch/claude-cowork-linux は、macOS ネイティブ モジュールをスタブ化し、VM を完全にスキップすることにより、Linux x86_64 上でネイティブに実行される同じ Cowork モードを実証します。 Linux 機能は製品内にすでに存在しています。欠けているのは、公開された Linux ターゲットです。
なぜそれが欠けていることが重要なのか
Claude Desktop は OAuth トークン、API キーを処理します

ys、および拡張構成。これは、開発者ワークステーション上で実行される資格情報処理アプリケーションです。
Linux ユーザーは現在、Windows Electron ビルドのサードパーティ製再パッケージを通じてこれを入手しています。主要なプロジェクトである aaddrick/claude-desktop-debian (約 4.5,000 個の星) は真に高品質です。署名された apt および dnf リポジトリ、.deb/.rpm/AppImage/AUR/Nix ビルド、CI テスト済み、--doctor 診断、数日以内のアップストリーム トラッキング (最新リリース 05/06/2026、Claude Desktop 1.11187.1 を追跡)。また、定義上、ベンダー署名もベンダー監査も受けていません。 Anthropic は正式なものを何も提供していないため、少なからぬ数の Claude ユーザーが自分の資格情報とローカル ファイルシステムへのアクセスをサードパーティの再パッケージに委ねています。構造的リスクは現在の管理者に関するものではありません。これは、Anthropic 自身のエージェント ランタイムが依存するプラットフォームにおける前例です。
Linux は非主流の開発者プラットフォームではありません。 Stack Overflow 2025 (177 か国、49,000 人以上の回答者): プロの開発者の 27.7% にとって、Ubuntu は主要な OS です。 StatCounter: インドのデスクトップ Linux 16.21% (2024 年 7 月)。米国は2025年6月に5％を超えた。
Claude Code がすでに Linux 用に使用しているのと同じディストリビューション パイプラインを使用して、Anthropic が運営する apt リポジトリ経由で、現在の 2 つの Ubuntu LTS リリース (および Debian) を対象とした署名付き .deb として、Linux 用の公式 Claude Desktop ビルドを公開します。
Claude Code CLI: 公式であり、署名された apt/dnf/apk リポジトリを使用して Linux 上でネイティブに実行されます。端末ワークフローに優れており、ローカル MCP サーバーを正常に実行します。デスクトップ GUI の代替ではありません。デスクトップ拡張機能として Claude Code プラグインをテストするためのサーフェスはなく、コンピューターも使用せず、Cowork も使用しません。
Web クライアント (claude.ai): リモート MCP コネクタをサポートしますが、デスクトップ拡張機能、コンピュータの使用、Cowork はサポートしません。ブラウザがクラッシュすると会話状態が失われます。より高いRAMと

ネイティブクライアントよりもバッテリーコストがかかります。
コミュニティの再パッケージ (aaddrick/claude-desktop-debian、johnzfitch/claude-cowork-linux、Snap ラッパー、k3d3 NixOS flake): 機能しており、現在使用しているもの。非公式、ベンダー署名、ベンダー監査はありません。
Wine での Windows ビルド: クリップボードとフォントの統合が壊れ、MCP サブプロセス処理が信頼できず、ファーストパーティのセキュリティ更新プログラムがありません。
macOS または Windows に切り替えてプラグインをテストする: 現在の回避策。反復ごとに摩擦が発生します。本当の修正ではありません。
高 - 生産性に重大な影響を与える
私は主な開発環境として Ubuntu LTS を実行しています。 Stack Overflow 2025 Developer Survey によると、これはプロの開発者の 27.7% に当てはまります。
クロードコードのプラグインを開発しています。プラグインは Claude Desktop 拡張機能としてテストおよび反復されます。これには Claude Desktop が必要です。 Linux ビルドはありません。
現在の回避策は、拡張機能としてプラグインをテストする必要があるたびに macOS に切り替えることです。これは、私が Linux 上で構築しているプラ​​グインを反復するたびに摩擦が発生し、人間工学的に非常に悪いため、Linux からのプラグイン開発を完全に思いとどまらせます。
公式 Linux ビルドでは、Anthropic が署名したリポジトリから apt 経由でインストールし、Claude Code プラグインを開発、テスト、反復して、プラグインを作成したのと同じマシン上でデスクトップ拡張機能として使用します。
耐荷重クレームのソース。可能な場合はプライマリという名前が付けられます。
claude.com/download: 「Linux では利用できません」。
code.claude.com/docs/en/desktop: macOS および Windows で利用できるデスクトップ アプリ。
code.claude.com/docs/en/setup: 署名された apt、dnf、および apk リポジトリ。プラットフォームごとのバイナリ (linux-x64、linux-arm64、linux-x64-musl、linux-arm64-musl)。 Ubuntu 20.04+/Debian 10+。
Simon Willison、「Claude Cowork の第一印象」、2026 年 12 月 1 日 (simonwillison.net): Apple の Virtualiza 経由の VZVirtualMachine

カスタム Linux ルート ファイルシステムをブートするフレームワーク。
Pluto Security: VM 内の Ubuntu 22.04 の詳細なリバース エンジニアリングを裏付けています。
pvieito、「Inside Claude Cowork」: macOS ホスト → Apple Virtualization Framework → Ubuntu 22.04 VM → bubblewrap → seccomp → /usr/local/bin/claude のクロード コード。
Anthropic のドキュメントでは、リバース エンジニアリングされた内部構造は確認せずに、ハイパーバイザーの分割 (macOS 上の Apple Virtualization.framework、Windows 上の Hyper-V) を確認しています。
johnzfitch/claude-cowork-linux: macOS ネイティブ モジュールをスタブ化し、VM を使用せずに Linux x86_64 上で直接 Cowork を実行する、動作中のコミュニティ ポート。
aaddrick/claude-desktop-debian: 星は約 4.5,000 個。 .deb、.rpm、AppImage、AUR、Nix; pkg.claude-desktop-debian.dev にある署名済みの apt および dnf リポジトリ。最新リリース v2.0.18+claude1.11187.1 (日付:2026 年 5 月 6 日)。 --医師の診断; CIテスト済み。 Linux 上の実験的な Cowork。
関連: aaddrick/claude-desktop-arch、emsi/claude-desktop、k3d3/claude-desktop-linux-flake。
StatCounter: インドのデスクトップ Linux 16.21% (2024 年 7 月)。米国は2025年6月に5%を超えた。 2025 年には全世界で約 4.7%。
ファーストパーティのビルドがロードマップにない場合
信頼性とセキュリティの問題のほとんどに対処する低コストのフォールバック: Linux は現在計画されていない (大まかな計画があれば) というインストール ドキュメントでの公式声明、推奨されるコミュニティ プロジェクトの承認、そのプロジェクトの 1 回限りのセキュリティ レビューの概要、資格情報の処理と MCP サーバー構成に関する Linux ユーザー向けの明示的なセキュリティ ガイダンス。
最も強力な内部の「今ではない」ため、この問題は丁寧に終わらせるのではなく、本当の会話を促します。
ボリュームはエンジニアリング税を正当化するものではありません。 Cowork の同等性、Windows の強化、エージェントの機能はすべて、おそらく 3 番目のデスクトップ プラットフォームを上回ります。

。
Linux の断片化により、ディストリビューション、ディスプレイ サーバー、サンドボックス モデル、グラフィックス スタックなど、不釣り合いなサポート税が発生します。コミュニティ プロジェクトのコミット ログには、表面 (AppArmor ユーザー ブロック、KDE ​​Plasma SNI レース、Wayland HiDPI、eCryptfs パス長の障害) が示されています。
Enterprise Linux 開発者は主にリモート開発と CLI によってサービスを提供します。デスクトップ GUI では、コストに見合った企業の収益が得られない可能性があります。
機会費用。 Linux デスクトップ上のすべてのエンジニアの四半期は、エージェントの品質、MCP エコシステム、コワークの強化、またはエンタープライズ コントロール プレーンに関するものではない四半期です。
分布は簡単ではありません。署名付きリポジトリ、GPG キー、AppImage 署名、Snap、AUR、Nix。
合理的な上級決定であれば、これらを考慮して「現在のロードマップには載っていない」と結論付ける可能性がある。それは理解できると思います。私が理解できないのは、公的な立場がまったく存在しないことと、その沈黙が現在の Linux ユーザーに与える構造的なセキュリティ コストです。
この問題は自動トリアージ システムによって処理されていることは承知しています。私はこれを、明確なプライマリ アスクと低コストのフォールバック (追加コンテキストの「良いいいえ」パス) を備えた単一の統合リクエストとして作成しました。別の会場が正しい場合は、閉じるのではなくルーティングしてください。むしろ答えてください

[切り捨てられた]

## Original Extract

Preflight Checklist I have searched existing requests and this feature hasn't been requested yet This is a single feature request (not multiple features) Problem Statement Preflight note. The closest open issue is #40347. Related: #47316...

[FEATURE] Official Claude Desktop build for Linux (Ubuntu LTS / Debian) · Issue #65697 · anthropics/claude-code · GitHub
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
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
[FEATURE] Official Claude Desktop build for Linux (Ubuntu LTS / Debian) #65697
Copy link New issue Copy link Open Open [FEATURE] Official Claude Desktop build for Linux (Ubuntu LTS / Debian) #65697 Copy link Labels area:desktop enhancement New feature or request New feature or request platform:linux Issue specifically occurs on Linux Issue specifically occurs on Linux Description
Issue body actions Preflight Checklist
I have searched existing requests and this feature hasn't been requested yet
This is a single feature request (not multiple features)
Preflight note. The closest open issue is #40347 . Related: #47316 (closed), #38276 (closed as out of scope for this repo), #36011 (stale). I am filing this as a consolidation and extension of #40347 with corrected technical framing (Claude Code plugin development against Desktop extensions), named primary sourcing for the Cowork Linux-VM architecture, and current market data. Happy to merge into #40347 if maintainers prefer; please route rather than close if a different venue is correct.
On scope: this issue concerns Claude Code in two concrete ways. (1) Claude Code plugins are developed and tested against Claude Desktop extensions, which has no Linux build, so plugin work currently requires switching OS. (2) Cowork invokes the Claude Code binary inside a Linux VM on macOS, so the Linux execution path already exists inside the Claude Code product and is the practical thing missing as a published target.
What this issue is asking for
A public Anthropic position on Linux desktop support, and ideally a first-party build. A reasoned "not on the current roadmap, and here is why" would resolve most of what this issue is about. There is, to my knowledge, no public statement on Linux desktop support; the absence is itself part of the problem.
Anthropic distributes Claude Desktop for macOS and Windows only. The official download page states "Not available for Linux". Claude Code (the CLI) runs natively on Linux but is a terminal tool, not a substitute for the desktop GUI. Desktop extensions (the surface Claude Code plugins are tested against), computer use, desktop dictation and Cowork are available only in Claude Desktop. Linux users therefore have no officially supported graphical path to these capabilities, and in particular no way to develop and test Claude Code plugins as desktop extensions without switching to macOS or Windows.
Why this is structurally hard to justify
Anthropic already builds, signs and distributes Linux software. Per code.claude.com/docs/en/setup, Claude Code ships signed apt, dnf and apk repositories and per-architecture binaries (linux-x64, linux-arm64, musl variants). The pipeline exists.
The Cowork agent already depends on Linux inside the product. Independent reverse-engineering by Simon Willison on launch day (12/01/2026), corroborated by Pluto Security and pvieito ("Inside Claude Cowork"), found that on macOS Cowork boots a custom Ubuntu 22.04 VM via Apple's Virtualization Framework (VZVirtualMachine) and runs the Claude Code binary inside it under bubblewrap and seccomp. Anthropic's own documentation confirms the hypervisor split: Apple Virtualization.framework on macOS, Hyper-V on Windows. The community project johnzfitch/claude-cowork-linux demonstrates the same Cowork mode running natively on Linux x86_64 by stubbing the macOS native modules and skipping the VM entirely. The Linux capability already exists inside the product; what's missing is a published Linux target.
Why it matters that it is missing
Claude Desktop handles OAuth tokens, API keys, and extension configurations. It is a credential-handling application running on developer workstations.
Linux users currently obtain it via third-party repackages of the Windows Electron build. The leading project, aaddrick/claude-desktop-debian (roughly 4.5k stars), is genuinely high quality: signed apt and dnf repositories, .deb/.rpm/AppImage/AUR/Nix builds, CI-tested, a --doctor diagnostic, and upstream tracking within days (latest release 05/06/2026, tracking Claude Desktop 1.11187.1). It is also, by definition, not vendor-signed and not vendor-audited. A non-trivial number of Claude users entrust their credentials and local filesystem access to a third-party repackage because Anthropic ships nothing official. The structural risk is not about the current maintainers; it is the precedent on a platform Anthropic's own agent runtime depends on.
Linux is not a fringe developer platform. Stack Overflow 2025 (49,000+ respondents, 177 countries): Ubuntu primary OS for 27.7% of professional developers. StatCounter: India desktop Linux 16.21% (July 2024); US crossed 5% in June 2025.
Publish an official Claude Desktop build for Linux, targeting the two current Ubuntu LTS releases (and Debian) as a signed .deb via an Anthropic-operated apt repository, using the same distribution pipeline Claude Code already uses for Linux.
Claude Code CLI: official and runs natively on Linux with signed apt/dnf/apk repositories. Excellent for terminal workflows and runs local MCP servers fine. Not a substitute for the desktop GUI: no surface for testing Claude Code plugins as desktop extensions, no computer use, no Cowork.
Web client (claude.ai): supports remote MCP connectors but no desktop extensions, no computer use, no Cowork. Loses conversation state on browser crash; higher RAM and battery cost than a native client.
Community repackages (aaddrick/claude-desktop-debian, johnzfitch/claude-cowork-linux, Snap wrappers, k3d3 NixOS flake): functional and what I currently use. Unofficial, not vendor-signed, not vendor-audited.
Windows build under Wine: clipboard and font integration break, MCP subprocess handling is unreliable, no first-party security updates.
Switching to macOS or Windows to test plugins: current workaround. Friction on every iteration; not a real fix.
High - Significant impact on productivity
I run Ubuntu LTS as my primary development environment. Per the Stack Overflow 2025 Developer Survey, this is the case for 27.7% of professional developers.
I develop Claude Code plugins. Plugins are tested and iterated on as Claude Desktop extensions, which requires Claude Desktop. There is no Linux build.
The current workaround is to switch to macOS every time I need to test a plugin as an extension. This is friction on every iteration of a plugin I am building on Linux, and a sufficiently bad ergonomic that it discourages plugin development from Linux entirely.
With an official Linux build I would install via apt from an Anthropic-signed repository and develop, test and iterate on Claude Code plugins as desktop extensions on the same machine I write them on.
Sources for the load-bearing claims, named primary where possible.
claude.com/download: "Not available for Linux".
code.claude.com/docs/en/desktop: desktop app available for macOS and Windows.
code.claude.com/docs/en/setup: signed apt, dnf and apk repositories; per-platform binaries (linux-x64, linux-arm64, linux-x64-musl, linux-arm64-musl); Ubuntu 20.04+/Debian 10+.
Simon Willison, "First impressions of Claude Cowork", 12/01/2026 (simonwillison.net): VZVirtualMachine via Apple's Virtualization Framework booting a custom Linux root filesystem.
Pluto Security: corroborating reverse-engineering deep dive, Ubuntu 22.04 inside the VM.
pvieito, "Inside Claude Cowork": macOS host → Apple Virtualization Framework → Ubuntu 22.04 VM → bubblewrap → seccomp → Claude Code at /usr/local/bin/claude.
Anthropic documentation confirms the hypervisor split (Apple Virtualization.framework on macOS, Hyper-V on Windows) without confirming the reverse-engineered internals.
johnzfitch/claude-cowork-linux: working community port that stubs the macOS native modules and runs Cowork directly on Linux x86_64 with no VM.
aaddrick/claude-desktop-debian: roughly 4.5k stars; .deb, .rpm, AppImage, AUR, Nix; signed apt and dnf repositories at pkg.claude-desktop-debian.dev; latest release v2.0.18+claude1.11187.1 dated 05/06/2026; --doctor diagnostic; CI-tested; experimental Cowork on Linux.
Related: aaddrick/claude-desktop-arch, emsi/claude-desktop, k3d3/claude-desktop-linux-flake.
StatCounter: India desktop Linux 16.21% (July 2024); US crossed 5% in June 2025; global approximately 4.7% in 2025.
If a first-party build is not on the roadmap
A lower-cost fallback that would address most of the trust and security concerns: a public statement on the install documentation that Linux is not currently planned (with rough horizon if any), acknowledgement of a recommended community project, a one-off security review summary of that project, and explicit security guidance for Linux users on credential handling and MCP server configuration.
The strongest internal "not now", so this issue invites a real conversation rather than a polite close.
Volume does not justify the engineering tax. Cowork parity, Windows hardening and agent capability work all plausibly outrank a third desktop platform.
Linux fragmentation creates a disproportionate support tax: distros, display servers, sandboxing models, graphics stacks. The community project's commit log shows the surface (AppArmor userns blocks, KDE Plasma SNI races, Wayland HiDPI, eCryptfs path-length failures).
Enterprise Linux developers are largely served by remote development and the CLI. A desktop GUI may not unlock enterprise revenue proportionate to its cost.
Opportunity cost. Every engineer-quarter on Linux desktop is a quarter not on agent quality, MCP ecosystem, Cowork hardening, or enterprise control planes.
Distribution is non-trivial. Signed repos, GPG keys, AppImage signing, Snap, AUR, Nix.
A reasonable senior decision could weigh these and conclude "not on the current roadmap". I would understand that. What I do not understand is the absence of any public position at all, and the structural security cost of that silence to current Linux users.
I am aware this issue is processed by an automated triage system. I have written it as a single consolidated request with a clear primary ask and a lower-cost fallback (the "good no" path in Additional Context). Please route rather than close if a different venue is correct; please respond rather

[truncated]
