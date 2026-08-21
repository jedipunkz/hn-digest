---
source: "https://github.com/alibaba/anolisa"
hn_url: "https://news.ycombinator.com/item?id=49383466"
title: "Anolisa – Agentic OS with runtime, security, observability and token compression"
article_title: "GitHub - alibaba/anolisa: ANOLISA (Agentic Nexus Operating Layer & Interface System Architecture) | Agentic OS with runtime, security, observability, and Tokenless response compression for lower token usage and cost. · GitHub"
image: "https://repository-images.githubusercontent.com/1195873793/e3f09d3a-3b99-4615-8948-44bccc219b45"
author: "forrestly"
captured_at: "2026-08-21T03:42:11Z"
capture_tool: "hn-digest"
hn_id: 49383466
score: 1
comments: 0
posted_at: "2026-08-21T03:36:01Z"
tags:
  - hacker-news
  - translated
---

# Anolisa – Agentic OS with runtime, security, observability and token compression

- HN: [49383466](https://news.ycombinator.com/item?id=49383466)
- Source: [github.com](https://github.com/alibaba/anolisa)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T03:36:01Z

## Translation

タイトル: Anolisa – ランタイム、セキュリティ、可観測性、トークン圧縮を備えたエージェントティック OS
記事のタイトル: GitHub - alibaba/anolisa: ANOLISA (Agentic Nexus オペレーティング層およびインターフェイス システム アーキテクチャ) |ランタイム、セキュリティ、可観測性、トークンレス応答圧縮を備えたエージェントティック OS により、トークンの使用量とコストが削減されます。 · GitHub
説明: ANOLISA (Agentic Nexus オペレーティング層およびインターフェイス システム アーキテクチャ) |ランタイム、セキュリティ、可観測性、トークンレス応答圧縮を備えたエージェントティック OS により、トークンの使用量とコストが削減されます。 - アリババ/アノリサ

記事本文:
GitHub - alibaba/anolisa: ANOLISA (Agentic Nexus オペレーティング層およびインターフェイス システム アーキテクチャ) |ランタイム、セキュリティ、可観測性、トークンレス応答圧縮を備えたエージェントティック OS により、トークンの使用量とコストが削減されます。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
アリババ
/
アノリサ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
2,500 コミット 2,500 コミット

s フォルダーとファイル
.github .github docker docker docs docs スクリプト スクリプト仕様 仕様 src src テスト テスト Web サイト Web サイト .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHANGELOG_zh.md CHANGELOG_zh.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md COTRIBUTING_zh.md COTRIBUTING_zh.md ライセンス ライセンス Makefile Makefile 通知 通知 README.md README.md README_zh.md README_zh.md SECURITY.md SECURITY.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
A Gentic N exus オペレーティング層およびインターフェース システム アーキテクチャ
エージェント ワークロードのオペレーティング システム層。
エージェントが端末から直接システムを操作し、ツールを削除できるようにします。
コストがかかる前にモデルに到達するレスポンス — シェルを維持したまま、
すでに実行しているエージェント フレームワークとサンドボックス。
中文版 · ウェブサイト ·
クイックスタート ·
ユーザーガイド ·
貢献する
ANOLISA は、AI エージェント ワークロードのサーバー側オペレーティング レイヤーです。対処します
エージェント実行の 3 つの実際的な制約: 端末エントリ、トークン コスト、
実行環境。シェル、エージェント フレームワーク、サンドボックスはそのままにしておきます
すでに使用しています。 ANOLISA CLI は、単一のインストール エントリ ポイントを提供します。
機能は個別に有効にすることができます。
ANOLISA は初めてですか?
クイックスタートで最初の結果を選択→
エージェントエントリー
コンテキストの効率性
ランタイムとセキュリティ
コシュン
シェル副操縦士
トークンレス
ツール出力の圧縮
ws-ckpt
チェックポイントとロールバック
OSスキル
システムとDevOpsの専門知識
エージェントサイト
トレースとトークンの可視性
スキルFS
集中的なスキルビュー
クチューナー
カーネルのチューニング
エージェントの記憶
セッション間の記憶
エージェント セク コア
サンドボックスと検証
ブレイズ
サンドボックスのライフサイクル
解決するもの
エージェントをターミナルで直接動作させる
cosh-ng は AI ネイティブの Linux ターミナルです。ファミリを維持します。

ar Bash/Zsh の動作、
次に、意図を理解し、ツールとスキルを使用し、質問できるエージェントを追加します。
危険な作業の前に承認を得るため。シェルコマンドと自然言語は同じものを共有します
ユーザーに別のチャット アプリケーションを強制するのではなく、ターミナルを使用します。
トークンがどこへ行くのかを確認し、モデルに到達する前に無駄を削減します。
トークンレス
ツールのスキーマと応答がモデルに到達する前に、冗長性を削除します。
エージェントの記憶
セッション全体で有用なコンテキストを再利用します。
SkillFS は、
現在のスキル ビューに焦点が当てられ、必要に応じて他のスキルを発見できるようになります。
エージェントサイト
トークンがどこで使われるかを示します。
エージェントがカーネルから実行される様子を確認する
Linux では、AgentSight は eBPF を使用して、コードを変更せずにエージェントを監視します。
トークンの使用とサブエージェントを使用して、モデルとツールの呼び出しを通じてユーザー入力を追跡します。
同じビュー内で分岐します。
エージェントサイト-デモ-en-ebpf-cover-1080p.mp4
Claude Code を使用してトークンレスを 3 分で試してみる
Token-less をインストールし、Claude Code に接続します。
カール -fsSL https://get.agentic-os.sh |バッシュ
エクスポート PATH= " $HOME /.local/bin: $PATH "
anolisa トークンレスでインストール
anolisaアダプターはトークンレスのクロードコードを有効にします
Claude Code を再起動し、ツールを多用するタスクを 1 つ実行して、結果を検査します。
トークンレス統計の概要
トークンレス統計リスト -- 制限 5
完全なトークンレス クイック スタートを開く →
· ユーザーマニュアルを読む
トークンレスコンテキスト効率-hd.mp4
観察されたあるコーディング タスクでは、トークンレスにより 317,000 トークン (40.5%) が節約されました。
AgentSight の測定について。
結果はワークロードによって異なります。
デバッグとトレースはフィールド ブラックリストによって削除され、メタデータは null として、
タグ / extra を空の値として使用します。圧縮はエージェントと
モデルなので、エージェント フレームワークのコードは変更されません。ドロップされた配列アイテムは引き続き取得可能
<<tokenless:KEY>> マーカーを介して、圧縮を可逆的に保ちます。
節約は、ツールの応答に適用されます。

全体ではなく文脈
セッション法案。トークンレスユーザーマニュアル
では、特定のワークロードに対する効果を見積もる方法について説明します。
すべてのエージェントの実行に境界と戻る方法を与える
ANOLISA はエージェント実行環境を構築しています。
エージェント セク コア
リスクのある業務を隔離し、
ws-ckpt は回復を維持します
ワークスペース変更のポイント。
変更されたスキルを実行前にキャッチする
署名されたスキルが変更されると、エージェントのレポートが再度使用される前にずれていました。
再スキャンでは、ブロックしている検出結果が拒否として記録されます。
エージェントのデモを試す →
・スキル台帳ガイド
skill-ledger-demo-en-with-cover.mp4
ランタイムまたはセキュリティの開始点を選択 →
· ANOLISA CLI を開始する
ANOLISA CLI は、共通のインストール エントリ ポイントです。 cosh-ng がインストールされている場所
システムモード。トークンレス機能やその他の機能は個別に追加できます。
カール -fsSL https://get.agentic-os.sh |バッシュ
sudo anolisa --install-mode システムインストール cosh-ng
anolisa トークンレスでインストール
cosh を実行して AI ネイティブのターミナルに入ります。トークンレスでツールを最適化することも可能
フレームワークを変更することなく、既存のエージェントからの呼び出しを可能にします。
クイックスタート ·
設置・
ユーザーガイド ·
トラブルシューティング ·
ソースからビルド ·
変更履歴
DingTalk でスキャンして ANOLISA コミュニティに参加します。
バグの問題をオープンし、
機能リクエスト。
CONTRIBUTING.md を読む
プルリクエストを送信する前に。
脆弱性を報告するには、
セキュリティポリシー。
ANOLISA は以下に基づいてリリースされています。
Apache ライセンス 2.0 。
ANOLISA (Agentic Nexus オペレーティング層およびインターフェイス システム アーキテクチャ) |ランタイム、セキュリティ、可観測性、トークンレス応答圧縮を備えたエージェントティック OS により、トークンの使用量とコストが削減されます。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
94 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

n

## Original Extract

ANOLISA (Agentic Nexus Operating Layer & Interface System Architecture) | Agentic OS with runtime, security, observability, and Tokenless response compression for lower token usage and cost. - alibaba/anolisa

GitHub - alibaba/anolisa: ANOLISA (Agentic Nexus Operating Layer & Interface System Architecture) | Agentic OS with runtime, security, observability, and Tokenless response compression for lower token usage and cost. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
alibaba
/
anolisa
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,500 Commits 2,500 Commits Folders and files
.github .github docker docker docs docs scripts scripts specs specs src src tests tests website website .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHANGELOG_zh.md CHANGELOG_zh.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING_zh.md CONTRIBUTING_zh.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md README_zh.md README_zh.md SECURITY.md SECURITY.md View all files Repository files navigation
A gentic N exus O perating L ayer & I nterface S ystem A rchitecture
The operating system layer for Agent workloads.
Let Agents drive the system straight from your terminal, and strip the tool
responses that reach the model before they cost you — while keeping the Shell,
Agent framework, and sandbox you already run.
中文版 · Website ·
Quick Start ·
User Guide ·
Contributing
ANOLISA is a server-side operating layer for AI Agent workloads. It addresses
three practical constraints of Agent execution: terminal entry, Token cost, and
execution environments. Keep the Shell, Agent framework, and sandbox you
already use. ANOLISA CLI provides a single installation entry point, while each
capability can be enabled independently.
New to ANOLISA?
Choose your first outcome in the Quick Start →
Agent entry
Context efficiency
Runtime & security
cosh-ng
Shell copilot
Token-less
Tool-output compression
ws-ckpt
Checkpoint and rollback
OS Skills
System and DevOps expertise
AgentSight
Trace and Token visibility
SkillFS
Focused Skill views
ktuner
Kernel tuning
Agent Memory
Cross-session memory
Agent Sec Core
Sandbox and verification
Blaze
Sandbox lifecycle
What it solves
Let the Agent work directly in the terminal
cosh-ng is an AI-native Linux terminal: it keeps familiar Bash/Zsh behavior,
then adds an Agent that can understand intent, use tools and Skills, and ask
for approval before risky work. Shell commands and natural language share one
terminal instead of forcing users into a separate chat application.
See where Tokens go and cut waste before it reaches the model
Token-less
removes redundancy from tool schemas and responses before they reach the model.
Agent Memory
reuses useful context across sessions.
SkillFS keeps the
current Skill view focused and makes other Skills discoverable when needed.
AgentSight
shows where Tokens are spent.
See an Agent run from the kernel up
On Linux, AgentSight uses eBPF to observe an Agent without changing its code.
Follow user input through model and tool calls, with Token use and sub-agent
branches in the same view.
agentsight-demo-en-ebpf-cover-1080p.mp4
Try Token-less with Claude Code in 3 minutes
Install Token-less and connect it to Claude Code:
curl -fsSL https://get.agentic-os.sh | bash
export PATH= " $HOME /.local/bin: $PATH "
anolisa install tokenless
anolisa adapter enable tokenless claude-code
Restart Claude Code, run one tool-heavy task, then inspect the result:
tokenless stats summary
tokenless stats list --limit 5
Open the full Token-less Quick Start →
· Read the user manual
tokenless-context-efficiency-hd.mp4
In one observed coding task, Token-less saved 317K Tokens (40.5%), based
on AgentSight measurements.
Results vary by workload.
debug and trace are dropped by the field blacklist, metadata as null, and
tags / extra as empty values. Compression runs between the Agent and the
model, so no Agent framework code changes. Dropped array items stay retrievable
through a <<tokenless:KEY>> marker, which keeps the compression reversible.
Savings apply to the tool responses entering the context, not to the whole
session bill. The Token-less user manual
explains how to estimate the effect for a given workload.
Give every Agent execution a boundary and a way back
ANOLISA is building out the Agent execution environment:
Agent Sec Core
isolates risky operations, and
ws-ckpt keeps recovery
points for workspace changes.
Catch a changed Skill before it runs
When a signed Skill changes, the Agent reports drifted before using it again.
A rescan records blocking findings as deny .
Try the Agent demo →
· Skill Ledger guide
skill-ledger-demo-en-with-cover.mp4
Choose a runtime or security starting point →
· Start with ANOLISA CLI
ANOLISA CLI is the common installation entry point. cosh-ng is installed in
system mode; Token-less and other capabilities can be added independently.
curl -fsSL https://get.agentic-os.sh | bash
sudo anolisa --install-mode system install cosh-ng
anolisa install tokenless
Run cosh to enter the AI-native terminal. Token-less can also optimize tool
calls from an existing Agent without changing its framework.
Quick Start ·
Installation ·
User Guide ·
Troubleshooting ·
Build from Source ·
Changelog
Scan with DingTalk to join the ANOLISA community.
Open an issue for bugs and
feature requests.
Read CONTRIBUTING.md
before submitting a pull request.
Report vulnerabilities through the
Security Policy .
ANOLISA is released under the
Apache License 2.0 .
ANOLISA (Agentic Nexus Operating Layer & Interface System Architecture) | Agentic OS with runtime, security, observability, and Tokenless response compression for lower token usage and cost.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
94 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
