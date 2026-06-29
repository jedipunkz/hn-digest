---
source: "https://github.com/Vadale/project-guardian"
hn_url: "https://news.ycombinator.com/item?id=48725632"
title: "A user-space firewall that gates an AI agent's actions"
article_title: "GitHub - Vadale/project-guardian: AI Guardian Firewall — a local, user-space, agent-agnostic firewall that mediates an autonomous AI agent's actions (files, shell, network, services) with a deterministic policy boundary, a tamper-evident audit log, and a human-in-the-loop approval cockpit. No kernel\n[truncated]"
author: "grauk"
captured_at: "2026-06-29T22:23:14Z"
capture_tool: "hn-digest"
hn_id: 48725632
score: 2
comments: 0
posted_at: "2026-06-29T21:38:21Z"
tags:
  - hacker-news
  - translated
---

# A user-space firewall that gates an AI agent's actions

- HN: [48725632](https://news.ycombinator.com/item?id=48725632)
- Source: [github.com](https://github.com/Vadale/project-guardian)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T21:38:21Z

## Translation

タイトル: AI エージェントのアクションをゲートするユーザー空間のファイアウォール
記事のタイトル: GitHub - Vadale/project-guardian: AI Guardian Firewall — 決定論的なポリシー境界、改ざん明示監査ログ、および人間参加型承認コックピットを使用して自律型 AI エージェントのアクション (ファイル、シェル、ネットワーク、サービス) を仲介する、ローカル、ユーザー空間、エージェントに依存しないファイアウォール。カーネルなし
[切り捨てられた]
説明: AI Guardian ファイアウォール — 決定論的なポリシー境界、改ざん防止監査ログ、および人間参加型の承認コックピットを使用して、自律型 AI エージェントのアクション (ファイル、シェル、ネットワーク、サービス) を仲介する、ローカルのユーザー空間のエージェントに依存しないファイアウォールです。カーネルモジュールはありません。 Apache-2.0。 - バダーレ/プロ
[切り捨てられた]

記事本文:
GitHub - Vadale/project-guardian: AI Guardian ファイアウォール — 決定論的なポリシー境界、改ざん明示的な監査ログ、および人間参加型の承認コックピットを使用して、自律 AI エージェントのアクション (ファイル、シェル、ネットワーク、サービス) を仲介する、ローカル、ユーザー空間、エージェントに依存しないファイアウォールです。カーネルモジュールはありません。 Apache-2.0。 · GitHub
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
エアコンを切り替えました

別のタブまたはウィンドウでカウントされます。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
バダーレ
/
プロジェクトガーディアン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
65 コミット 65 コミット .claude/ エージェント .claude/ エージェント .github/ ワークフロー .github/ ワークフロー クレート クレート ドキュメント ドキュメント 評価 評価例 例 ファズ ファズ ポリシー ポリシー スクリプト/ デモ スクリプト/ デモ ui ui .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.mddeny.tomldeny.tomlrust-toolchain.tomlrust-toolchain.toml すべてのファイルを表示リポジトリ ファイルのナビゲーション
Project Guardian — 自律エージェント用の AI Guardian ファイアウォール
ステータス: 動作中の製品、v0.1.0 がリリースされました。 Rust ワークスペース、196 テスト グリーン。
フェーズ 4 (強化) によって実装: 決定論的なポリシー エンジン、
改ざん明示監査ログ (オプションで封印された鍵署名付き)、勧告チェッカー、
MCP ゲートウェイ + stdio トランスポート、デーモン + 制御ソケット、ターミナル
承認コックピット (TUI)、AgentDojo 評価ハーネス、TLS を使用したネットワーク プロキシ
傍受 (ブローカーが挿入した認証情報、漏洩検査、デフォルト拒否)
egress、コックピット ask -routing)、OS exec サンドボックス、トークン ブローカー (OS
キーチェーン + 最小権限に関する注意事項)、軽量の検証可能な資格情報、
アダプティブ提案 + 安全性レポート、ed25519 署名付きコミュニティ ポリシー
パック、および固有のクリティカル カテゴリ フロア (お金 / 資格情報 /
抽出/不可逆的な削除は、サイレント許可に解決することはできません。
署名入りパック経由）。ゲ

開始しました: docs/user-guide.md 。
1.0 の残り: 署名/公証されたパッケージとデスクトップ GUI — を参照してください。
ロードマップ.md 。
評価: ローカル 12B エージェントを使用する AgentDojo では、Guardian はプロンプト インジェクションをカットします。
バンキング スイートに対する攻撃成功率が 100% → 0% (決定的な拒否)
お金の動き）。弊社独自の GuardianBench — ベンチマーク
アクション ファイアウォール用に構築 — スコアは偽陰性 0%、偽陽性 0%、100%
8 つのドメインにわたる拒否の正確性、さらにトークン化層での PII 漏洩は 0%
(データ ブローカー、ADR-0005)。評価/全文を参照してください。
正直に指摘されたスコアカード (アクション ファイアウォールの範囲が終了する場所を含む - 以下)。
📄 ホワイトペーパー: 設計と脅威モデル (PDF) — またはお読みください
GitHub にあります (その仕組みとエージェントへの影響を示す図付き)。
ライセンス: Apache-2.0 · ガバナンス: 貢献 ·
セキュリティ · CODE_OF_CONDUCT · ADR
この README は正規の仕様 (アイデア、完全な機能セット、アーキテクチャ、脅威) です。
モデル）。構築方法と順序については、 ROADMAP.md を参照してください。何のために
着陸しました。 docs/changelog.md を参照してください。
⚠️免責事項 — ご自身の責任でご使用ください
Guardian は、次の処理を行うように構成できる初期段階のソフトウェア (v0.1.0) です。
機密データ (資格情報、個人データ、財務詳細)。提供されます
Apache-2.0 ライセンスに基づく「現状のまま」、いかなる種類の保証もありません
(セクション 7 ～ 8 を参照)。法律で認められる最大限の範囲で、著者はいかなる拒否も受け入れません。
あらゆる損害、データ損失、セキュリティ侵害、金銭的損失、その他に対する責任
このソフトウェアの使用、誤用、または使用不能から生じる損害。あなたは
お客様の目的、お客様がどのように適合するかを評価することに単独で責任を負います。
ポリシーを構成し、それを介してルーティングされるデータのセキュリティを設定します。それは
認定、監査、または実稼働環境で強化されていないため、信頼してはなりません

として
これは、一か八かの作業負荷または規制された作業負荷に対する唯一の安全策です。 SECURITY.md を参照してください。
脅威モデルと脆弱性を報告する方法について説明します。
Guardian は、自律的なネットワークとネットワークの間に位置する、ローカルのユーザー空間の「ファイアウォール」です。
AI エージェントとそれが触れる可能性のあるもの - ファイル、シェル、ネットワーク、
あなたが委任したオンライン サービス。エージェントを信頼していません。毎
エージェントが試行するアクションは、構造化されたアクションとしてインターセプトされます。
決定論的なポリシー エンジンによって評価されるエージェントのツール/MCP 境界、および
— 意思決定に人間が必要な場合 — 別の担当者が平易な言葉で説明
承認または拒否する前に、「翻訳者」モデルを選択してください。 Guardian はエージェントに依存しません
(エージェントがクロード、GPT、ラマなどによって駆動されているかどうかは関係ありません)
それ以外)、OS フレンドリー (カーネル モジュールをインストールしたり、カーネル モジュールと競合したりすることはありません)
制御用のオペレーティング システム）。
最速 — 事前に構築されたバイナリ (ツールチェーンは必要ありません) を次の場所からダウンロードします。
最新リリース。
署名されていないため、OS は 1 回尋ねます: macOS → 右クリック → 開く ; Windows →
SmartScreen → 詳細 → とにかく実行する ; Linux → chmod +x ガーディアン 。それからガーディアン --help 。
(Windows は実験的/未テストです。 docs/user-guide.md を参照してください。)
1 つのコマンドで設定します。 Guardian init は ~/.guardian/{config.toml,policy.toml} を作成します
あなたの役割に応じて、正確な次のステップと貼り付ける MCP スニペットを出力します。
ガーディアン初期化番号または: --rolepersonal-assistant
Guardian-daemon # ターミナル 1 — サービス
ガーディアン UI # ターミナル 2 — 承認コックピット (TUI)
次に、エージェントの MCP クライアントを Guardian に指定します (スニペット Guardian init が出力されます —
クロード コード、カーソル、または任意の MCP クライアントで動作します):
{
"mcpサーバー": {
"guardian" : { "command" : " Guardian " , "args" : [ " mcp " 、 " --daemon " 、 " /tmp/guardian.sock " ] }
}
}
アクションにその日の承認が必要な場合

mon はデスクトップ通知を生成するので、
コックピットを監視する必要はありません (無効にするには、設定で notification = false を設定します)。
またはソースからビルドする - Rust ツールチェーンが必要です。
カーゴビルド --release
# 1) 信号機メディエーションをエンドツーエンドで確認します (スクリプト化、セットアップなし)
カーゴ ラン -p ガーディアン-cli -- デモ
# 2) 内部のレッドチームのスコアカード (決定的、モデルは必要ありません)
カーゴ ラン -p ガーディアン-cli -- eval
# ...そして、アクション ファイアウォール ベンチマークである GuardianBench (FN 0% / FP 0% / 拒否 100%):
GUARDIAN_BIN=ターゲット/リリース/guardian python3 評価/guardianbench/guardianbench.py
# 3) 実際のエージェントの完全なループ - 3 つの端末:
GUARDIAN_SOCK=/tmp/g.sock Cargo run -p Guardian-daemon # サービス
GUARDIAN_SOCK=/tmp/g.sock Cargo run -p Guardian-cli -- ui # 承認コックピット (TUI)
# その後、MCP クライアント (例: クロード コード) を次のように指定します。
# ガーディアン mcp --daemon /tmp/g.sock
カーゴ test --workspace を使用してテストを実行します。ガーディアンの効果を測定する
エージェントの攻撃成功率: 評価/ 。
エージェントは「話すチャットボット」から「行動するエージェント」になりました - 彼らは読んで、
ファイルの書き込み、シェル コマンドの実行、閲覧、購入、電子メールの送信など、ますます多くの機能が実行されます。
タッチに敏感なアカウント (銀行、医療記録、公共行政ポータル)。
これにより、次の 4 つの具体的なリスクが生じます。
機密データの漏洩と破壊的なミス。エージェントに直接指示する
アカウント、電子メール、プライベート文書にアクセスすると、ユーザーはプライバシーにさらされます。
違反、幻覚による破壊行為、外部からの攻撃。
早速注射。この時代の主要なエージェントセキュリティの脅威: コンテンツ
エージェントが読み取る (Web ページ、PDF、電子メール、ツールの結果) には次のものが含まれる場合があります
エージェントをハイジャックして、ユーザーが決して求めていないことを実行させる命令。
クリック疲れ/インフォームド・コンセントの失敗。システムレベルのエージェントがポップアップ表示されます
承認

スクリプトと API 呼び出しのリクエスト。技術者以外のユーザーは、
それらを理解し、すべてを盲目的に承認すると、安全性が無効になります。
人間に面した操縦翼面がなく、トレーサビリティもありません。既存のツール
(生のハーネス許可プロンプト、Docker) はプログラマー向けに構築されています。ありません
直感的な「制御室」、および改ざんの明らかな記録を保持する簡単な方法はありません。
エージェントが実際に何をしたか（透明性義務などに関連）
EU AI 法第 2 条50）。
3. 設計原則 (交渉不可)
これらは、後のすべてのトレードオフを決定するルールです。
セキュリティ境界は決定的です。 LLM は決して境界ではありません。
強制 (許可/要求/拒否) は、動作が次のようなルール エンジンによって行われます。
予測可能でテスト可能。 LLM は間違っている可能性があり、次の方法で攻撃される可能性があります。
プロンプトインジェクションのため、翻訳とリスクスコア付けのみに使用され、決して使用されません。
ロックを解除します。
エージェントの散文ではなく、構造化されたアクションを傍受します。ポリシーエンジン
そして翻訳者は実際にインターセプトされたアクション (ツール呼び出しとそのアクション) を調べます。
引数、実際の HTTP リクエスト、ファイル操作) — エージェントでは決して行われません
何をしようとしているのかについての自然言語による主張。この主張は操作可能です。
アクションはそうではありません。
構造上、エージェントに依存しません。制御はアクション境界で適用され、
これは、どのモデルがアクションを生成したかに関係なく同じです。
カーネル空間ではなく、ユーザー空間。カーネル モジュールや OS フックは必要ありません。
ベンダーが付与する権利。 (§4 を参照 — これが中心的な決定です。)
ローカルファースト / プライバシーファースト。ポリシーの評価、学習、監査
ユーザーのマシン上でライブログを記録します。クラウドへのあらゆる送信はオプトインであり、
明示的な。
多層防御。ツール境界での調停が主な制御です。
OS サンドボックスとネットワーク プロキシは封じ込めのバックストップであり、プラン A ではありません。
でフェイルクローズされました

クリティカル パス、便宜上フェール オープン。での失敗
金銭/資格情報/窃盗の経路をブロックします。低リスクパスでの障害
正常に機能を低下させます (ログに記録され、既存のハーネスのデフォルトに従います)。
デフォルトで改ざん防止機能。 Guardian が決定したことはすべて、
追加専用、ハッシュチェーンされた、署名可能な監査ログ。
4. 重要な決定: ガーディアンがどこで行動するか
解決済み: ガーディアンはエージェントのアクション境界 (ハーネス / ) で動作します。
ツール呼び出し/MCP レイヤー - ユーザー空間内。 OS カーネルでは動作しません。
なぜ OS カーネル / OS フックではないのか
ディープ OS インターセプト (ユーザー空間を超えた Linux LSM/eBPF、macOS エンドポイント セキュリティ)
およびネットワーク拡張、Windows ミニフィルター/WFP カーネル コールアウト) が必要です
ベンダー付与の権利、コード署名、公証、プラットフォームごと
認証が必要です。 macOS と Windows では、これがオープンソース プロジェクトの壁になります
ソロ/コミュニティのメンテナーでもあります。
カーネルレベルのバグにより、ユーザーのマシンがクラッシュします。ミスの爆発範囲は
OS全体。
これは間違った高度です。システムコール レベルでは、 write(fd, buf, n) が表示されます。
「エージェントは不明な IBAN に 4,000 ユーロを送金しようとしています。」意図が読み取れるのは、
カーネルではなく、アクションの境界です。
ハーネスとアクションの境界が適切なチョークポイントである理由
最新のエージェント ハーネス (Claude Code、Cursor、OpenAI Agents ランタイム、

[切り捨てられた]

## Original Extract

AI Guardian Firewall — a local, user-space, agent-agnostic firewall that mediates an autonomous AI agent's actions (files, shell, network, services) with a deterministic policy boundary, a tamper-evident audit log, and a human-in-the-loop approval cockpit. No kernel modules. Apache-2.0. - Vadale/pro
[truncated]

GitHub - Vadale/project-guardian: AI Guardian Firewall — a local, user-space, agent-agnostic firewall that mediates an autonomous AI agent's actions (files, shell, network, services) with a deterministic policy boundary, a tamper-evident audit log, and a human-in-the-loop approval cockpit. No kernel modules. Apache-2.0. · GitHub
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
Vadale
/
project-guardian
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
65 Commits 65 Commits .claude/ agents .claude/ agents .github/ workflows .github/ workflows crates crates docs docs evaluation evaluation examples examples fuzz fuzz policies policies scripts/ demo scripts/ demo ui ui .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Project Guardian — An AI Guardian Firewall for Autonomous Agents
Status: working product, v0.1.0 released . Rust workspace, 196 tests green.
Implemented through Phase 4 (hardening) : the deterministic policy engine, the
tamper-evident audit log (optionally sealed-key signed ), the advisory Checker,
the MCP gateway + stdio transport, the daemon + control socket, the terminal
approval cockpit (TUI), the AgentDojo eval harness, the network proxy with TLS
interception (broker-injected credentials, exfiltration inspection, default-deny
egress, cockpit ask -routing), the OS exec sandbox , the token broker (OS
keychain + least-privilege caveats), lightweight verifiable credentials ,
adaptive suggestions + safety report , ed25519-signed community policy
packs , and an intrinsic critical-category floor (money / credentials /
exfiltration / irreversible deletion can never resolve to a silent allow , not even
via a signed pack). Getting started: docs/user-guide.md .
Remaining for 1.0: signed/notarized packaging and the desktop GUI — see
ROADMAP.md .
Evaluation: on AgentDojo with a local 12B agent, Guardian cuts the prompt-injection
attack-success rate on the banking suite from 100% → 0% (deterministic deny on
money-movement). Our own GuardianBench — a benchmark
built for an action-firewall — scores 0% false-negatives, 0% false-positives, 100%
refusal-correctness across 8 domains, plus 0% PII leaks in its tokenization layer
(the data broker, ADR-0005). See evaluation/ for the full,
honestly-caveated scorecard (including where an action-firewall's scope ends — below).
📄 White paper: design & threat model (PDF) — or read it
on GitHub (with diagrams of how it works and its impact on the agent).
License: Apache-2.0 · Governance: CONTRIBUTING ·
SECURITY · CODE_OF_CONDUCT · ADRs
This README is the canonical spec (idea, full feature set, architecture, threat
model). For how and in what order it's built, see ROADMAP.md ; for what's
landed, see docs/changelog.md .
⚠️ Disclaimer — use at your own risk
Guardian is early-stage software (v0.1.0) that can be configured to handle
sensitive data (credentials, personal data, financial details). It is provided
"AS IS", without warranty of any kind , under the Apache-2.0 license
(see Sections 7–8). To the maximum extent permitted by law, the author accepts no
liability for any damage, data loss, security breach, financial loss, or other
harm arising from the use, misuse, or inability to use this software. You are
solely responsible for evaluating its fitness for your purpose, for how you
configure your policy, and for the security of any data you route through it. It is
not certified, audited, or production-hardened, and must not be relied upon as
the sole safeguard for high-stakes or regulated workloads. See SECURITY.md
for the threat model and how to report a vulnerability.
Guardian is a local, user-space "firewall" that sits between an autonomous
AI agent and the things it can touch — your files, your shell, the network, and
the online services you delegate to it. It does not trust the agent. Every
action the agent attempts is intercepted as a structured action at the
agent's tool/MCP boundary, evaluated by a deterministic policy engine , and
— when a decision needs a human — explained in plain language by a separate
"translator" model before you approve or deny it. Guardian is agent-agnostic
(it does not care whether the agent is driven by Claude, GPT, Llama, or anything
else) and OS-friendly (it never installs a kernel module or fights the
operating system for control).
Fastest — download a prebuilt binary (no toolchain needed) from the
latest release .
It's unsigned, so the OS asks once: macOS → right-click → Open ; Windows →
SmartScreen → More info → Run anyway ; Linux → chmod +x guardian . Then guardian --help .
(Windows is experimental/untested — see docs/user-guide.md .)
Set it up in one command. guardian init creates ~/.guardian/{config.toml,policy.toml}
for your role and prints the exact next steps + the MCP snippet to paste:
guardian init # or: --role personal-assistant
guardian-daemon # terminal 1 — the service
guardian ui # terminal 2 — the approval cockpit (TUI)
Then point your agent's MCP client at Guardian (the snippet guardian init prints —
works for Claude Code, Cursor, or any MCP client):
{
"mcpServers" : {
"guardian" : { "command" : " guardian " , "args" : [ " mcp " , " --daemon " , " /tmp/guardian.sock " ] }
}
}
When an action needs your approval the daemon raises a desktop notification , so you
don't have to watch the cockpit (set notifications = false in the config to disable).
Or build from source — requires the Rust toolchain :
cargo build --release
# 1) see the traffic-light mediation end to end (scripted, no setup)
cargo run -p guardian-cli -- demo
# 2) the internal red-team scorecard (deterministic, no model needed)
cargo run -p guardian-cli -- eval
# ...and GuardianBench, our action-firewall benchmark (FN 0% / FP 0% / refusal 100%):
GUARDIAN_BIN=target/release/guardian python3 evaluation/guardianbench/guardianbench.py
# 3) the full loop for a real agent — three terminals:
GUARDIAN_SOCK=/tmp/g.sock cargo run -p guardian-daemon # the service
GUARDIAN_SOCK=/tmp/g.sock cargo run -p guardian-cli -- ui # the approval cockpit (TUI)
# then point an MCP client (e.g. Claude Code) at:
# guardian mcp --daemon /tmp/g.sock
Run the tests with cargo test --workspace . Measuring Guardian's effect on an
agent's attack-success rate: evaluation/ .
Agents went from "chatbots that talk" to "agents that act" — they read and
write files, run shell commands, browse, buy things, send email, and increasingly
touch sensitive accounts (banking, health records, public-administration portals).
That creates four concrete risks:
Sensitive-data exposure & destructive mistakes. Giving an agent direct
access to accounts, email, and private documents exposes the user to privacy
violations, hallucinated destructive actions, and external attacks.
Prompt injection. The dominant agent-security threat of this era: content
the agent reads (a web page, a PDF, an email, a tool result) can contain
instructions that hijack the agent into doing something the user never asked.
Click fatigue / informed-consent failure. System-level agents pop up
approval requests for scripts and API calls. Non-technical users do not
understand them and approve everything blindly, which nullifies the safety.
No human-facing control surface and no traceability. Existing tooling
(raw harness permission prompts, Docker) is built for programmers. There is no
intuitive "control room," and no easy way to keep a tamper-evident record of
what an agent actually did (relevant for transparency obligations such as the
EU AI Act, Art. 50).
3. Design principles (non-negotiable)
These are the rules that decide every later trade-off.
The security boundary is deterministic. The LLM is never the boundary.
Enforcement (allow / ask / deny) is done by a rule engine whose behavior is
predictable and testable. An LLM can be wrong and can be attacked via
prompt injection, so it is used only to translate and risk-score , never to
unlock.
Intercept structured actions, not the agent's prose. The policy engine
and the translator look at the real intercepted action (the tool call and its
arguments, the actual HTTP request, the file operation) — never at the agent's
natural-language claim about what it intends to do. The claim is manipulable;
the action is not.
Agent-agnostic by construction. Control is applied at the action boundary,
which is identical regardless of which model produced the action.
User-space, not kernel-space. No kernel modules, no OS hooks that require
vendor-granted entitlements. (See §4 — this is the central decision.)
Local-first / privacy-first. Policy evaluation, learning, and the audit
log live on the user's machine. Sending anything to the cloud is opt-in and
explicit.
Defense in depth. Mediation at the tool boundary is the primary control;
OS sandboxing and a network proxy are containment backstops, not the plan A.
Fail closed on the critical path, fail open on convenience. A failure in
the money/credential/exfiltration path blocks; a failure in a low-risk path
degrades gracefully (logs, defers to existing harness defaults).
Tamper-evident by default. Everything Guardian decides is written to an
append-only, hash-chained, signable audit log.
4. THE KEY DECISION: where Guardian acts
Resolved: Guardian acts at the agent's action boundary — the harness /
tool-call / MCP layer — in user-space. It does NOT act in the OS kernel.
Why not the OS kernel / OS hooks
Deep OS interception (Linux LSM/eBPF beyond user-space, macOS Endpoint Security
& Network Extension, Windows minifilter/WFP kernel callouts) requires
vendor-granted entitlements, code-signing, notarization, and per-platform
certification . On macOS and Windows this is a wall for an open-source project
and a solo/community maintainer.
Kernel-level bugs crash the user's machine. The blast radius of a mistake is
the whole OS.
It is the wrong altitude: at the syscall level you see write(fd, buf, n) , not
"the agent is about to wire €4,000 to an unknown IBAN." Intent is legible at
the action boundary, not at the kernel.
Why the harness/action boundary is the right choke point
Modern agent harnesses (Claude Code, Cursor, the OpenAI Agents runtime,

[truncated]
