---
source: "https://github.com/CorvinLabs/CorvinOS"
hn_url: "https://news.ycombinator.com/item?id=48749035"
title: "CorvinOS – self-hosted agentic OS where EU AI Act and GDPR compliance by design"
article_title: "GitHub - CorvinLabs/CorvinOS: Self-hosted agentic OS — connects Ollama, Claude & GPT-4 to Discord, Telegram, WhatsApp, Slack & Email. EU AI Act 2026 & GDPR compliance by architecture. pip install corvinos · GitHub"
author: "shumway"
captured_at: "2026-07-01T16:46:51Z"
capture_tool: "hn-digest"
hn_id: 48749035
score: 2
comments: 0
posted_at: "2026-07-01T16:02:43Z"
tags:
  - hacker-news
  - translated
---

# CorvinOS – self-hosted agentic OS where EU AI Act and GDPR compliance by design

- HN: [48749035](https://news.ycombinator.com/item?id=48749035)
- Source: [github.com](https://github.com/CorvinLabs/CorvinOS)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T16:02:43Z

## Translation

タイトル: CorvinOS – EU AI 法と GDPR に設計上準拠したセルフホスト型エージェント OS
記事のタイトル: GitHub - CorvinLabs/CorvinOS: セルフホスト型エージェント OS — Ollama、Claude、GPT-4 を Discord、Telegram、WhatsApp、Slack、Email に接続します。 EU AI 法 2026 とアーキテクチャ別の GDPR 準拠。 pip インストール corvinos · GitHub
説明: セルフホスト型エージェント OS — Ollama、Claude、GPT-4 を Discord、Telegram、WhatsApp、Slack、Email に接続します。 EU AI 法 2026 とアーキテクチャ別の GDPR 準拠。 pip インストール corvinos - CorvinLabs/CorvinOS

記事本文:
GitHub - CorvinLabs/CorvinOS: セルフホスト型エージェント OS — Ollama、Claude、GPT-4 を Discord、Telegram、WhatsApp、Slack、Email に接続します。 EU AI 法 2026 とアーキテクチャ別の GDPR 準拠。 pip インストール corvinos · GitHub
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
エラーがありました

r 読み込み中。このページをリロードしてください。
コルビンラボ
/
CorvinOS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
53 コミット 53 コミット .claude-plugin .claude-plugin .github .github .ldd .ldd .well-known .well-known 資産 アセット コンプライアンス コンプライアンス コア コア corvinOS corvinOS docs docs オペレーター オペレーター ops ops スクリプト スクリプト サイト サイト skill_basis skill_basis テスト テスト .corvin_repo .corvin_repo .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CCLA.md CCLA.md CHANGELOG.md CHANGELOG.md CLA-SIGNATORIES.md CLA-SIGNATORIES.md CLA.md CLA.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md INSTALLATION.md INSTALLATION.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md build_wheels.py build_wheels.py corvinOS_path_fix.pth corvinOS_path_fix.pth hatch_build.py hatch_build.py install.ps1 install.ps1 install.sh install.sh ollama-manifest.yaml ollama-manifest.yaml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
概要・
建築・
監査とコンプライアンス ·
A2Aネットワーク ·
エンジン層 ·
セキュリティ ·
EU AI法 ·
学習目標
インストールは 1 つです。七つの橋。任意の LLM。
CorvinOS は、単一の pip パッケージを通じて、Ollama、Claude、GPT-4、および任意の OpenRouter モデルを Discord、Telegram、WhatsApp、Slack、Email、Teams、および Signal に接続するセルフホスト型エージェント OS です。
pip install corvinos && python -m corvinOS
ローカルファースト — Ollama と --engine hermes を使用して 100% オフラインで実行します。 API キーは必要ありません。
エージェントティック — 実行時にサンドボックス ツールと新しいスキルを生成します。 5 つの AI エンジンにわたってサブタスクを委任します。
アーキテクチャによるコンプライアンス — EU AI 法 2026 (d

秘密保持、同意、ハウスルール）と GDPR（監査チェーン、データ保存、消去）は、ポリシー文書ではなく負荷がかかるコードです。フラグによって無効にすることはできません。
マルチテナント — 1 つのインスタンス、複数のユーザー、ペルソナ、チームがすべて分離されています。
どこでも自己ホスト可能 - Linux、macOS、Windows。 Docker 対応。シングルピップインストール。
CorvinOS は、EU AI 法を文書ではなくコードで施行します。
すべてのコンプライアンス要件 (開示、同意、監査の整合性、データの保存場所、送信制御、GDPR 消去) は、フラグ、環境変数、または構成のオーバーライドによって無効にすることのできない構造的な制約です。規制された展開では、ポリシーによる約束ではなく、検証可能な保証が得られます。
完全なセットアップ ガイドについては、INSTALLATION.md を参照してください。
推奨 — Linux、macOS、および Windows で同様に動作します。
pip インストールコルビノス
python -m corvinOS # Web コンソール (http://localhost:8765)
python -m corvinOS は PATH に依存しません。最初のパスでコンソールを起動します。
すべての OS で試してください — Microsoft Store / システム Python を含む、pip がインストールされている場所
PATH 上にないユーザーごとのスクリプト ディレクトリにフォールバックします (通常の理由)
corvin-serve は Windows では「見つかりません」です)。 Windows では、 py -m corvinOS も使用できます。
PATH に短い corvin-serve コマンドを追加したいですか?でインストールします
pipx — アプリを分離し、PATH を接続します。
すべてのプラットフォームで自動的に:
pipx コルビノスをインストールする
corvin-serve # Web コンソール (http://localhost:8765)
プレーン pip インストールからの corvin-serve は、スクリプト ディレクトリに配置されるとのみ機能します
PATH 上にあります。 python -m corvinOS を実行すると、そのディレクトリが PATH に追加されます。
したがって、 corvin-serve は新しいターミナルで動作しますが、 pipx (または python -m corvinOS ) が信頼できるクロスプラットフォーム パスです。
基本インストールは純粋な Python およびクロスプラットフォームであり、Web コンソールが提供されます
Linux、macOS でのセットアップまで

および Windows、クラウド/エッジ音声対応
(OpenAI + Microsoft Edge TTS) すぐに動作します。ローカル、オフラインの場合
音声モデルには、オプションで次のものが追加されます。
pip install " corvinos[voice] " # ローカル Piper TTS + 高速ウィスパー STT
ローカルモデルの依存関係 (piper-tts 、
fast-whisper ) 一部の Python バージョンには Windows ホイールがありません。それらを維持する
基本インストール外とは、pip install corvinos が確実にセットアップに到達することを意味します。
サポートされているすべてのプラットフォーム。
要件: Python 3.10+ · Linux、macOS 12+、または Windows 10/11 · Node.js 20+ ブリッジの場合のみ必要
デフォルトのエンジン: Claude Code (Claude Pro または Max サブスクリプションが必要)。
完全にローカルのゼロエグレス展開の場合: --engine hermes (Ollama、API キーなし)。
# macOS / Linux — 自作
醸造タップ CorvinLabs/corvinos && 醸造インストール corvinos
# Windows — スクープ
スクープ バケット コルビノスを追加 https://github.com/CorvinLabs/scoop-corvinos && スクープ インストール コルビノス
# conda / mamba (レビュー待ち)
conda install -c conda-forge corvinos
# 開発者がソースからインストール
git clone https://github.com/CorvinLabs/CorvinOS.git && cd CorvinOS
pip install -e " .[all] " && corvin-install
完全なドキュメント: docs/overview.md
1 つのコマンドで、サービス、構成、データ ディレクトリ、パッケージなどすべてが削除されます。
corvin-uninstall --purge # プロンプトを表示せずにすべてのファイルを削除します
pip uninstall corvinos -y # Python パッケージを削除します
--purge を指定しない場合、アンインストーラーはデータ ディレクトリ (監査ログ、API キー、セッション履歴) を削除する前に質問します。完全に非対話型のワイプには --purge を使用します。
pip アンインストール corvinos -y 後に残るのは、複製されたリポジトリ ディレクトリ (ソース インストール) だけです。不要になった場合は、rm -rf <repo> で削除します。
EU AI 法 2026 + GDPR: 構造的施行
CorvinOS は EU AI Act 2026 と GDPR を必須として実装

構造上の設計上の制約。すべての機能は、これによってコンプライアンスの保証が弱まるか? という質問に答える必要があります。
絶対制約 — 環境変数、フラグ、または設定によってこれらを無効にすることはできません。
開示は構造的にロックされています。 · 同意ゲートにはバイパスがありません。 · すべての監査イベントは応答の前にハッシュ チェーンを通過します。 · L34 は非準拠のエンジン生成をブロックします。 · L38 監査書き込み失敗は A2A リクエストをブロックします。 · L44 ハウス ルール ゲートにはキルフラグがありません。
voice-audit verify # 完全なハッシュ チェーンをたどる;休憩時に 1 番出口から出る
Bridge.sh Doctor # 監査チェーン検証を使用したブート セルフテスト
python -m corvin_compliance_reports.cli 処理レコードを生成 # GDPR Art. 30
完全なコンプライアンス リファレンス: docs/eu-ai-act/README.md · docs/audit-and-compliance.md
コンプライアンス スタックに影響を与えずに LLM を交換する
CorvinOS は、WorkerEngine プロトコル (L22) を介して AI バックエンドをコンプライアンス ランタイムから切り離します。すべてのエンジンは、ツール実行ブローカーを介してパスゲート、監査チェーン、アーティファクト登録を共有します。コンプライアンス設定を変更せずにプロバイダーを交換します。
マルチエージェント メッシュ - CorvinOS インスタンスが相互に通信する
複数の CorvinOS インスタンスが分散エージェント ネットワークを形成します。すべてのクロスインスタンス呼び出しには、暗号化署名、双方向構成証明、ノンス リプレイ保護、バイナリ添付ファイル検証が含まれます。監査優先不変式: 応答が送信される前に、エンベロープがハッシュ チェーンに書き込まれます。
パスゲート (書き込み保護) · bwrap 環境インジェクションを備えたシークレット ボールト · サンドボックス化された Forge ツール生成 · フェイルクローズ リンターを備えた SkillForge · マルチテナント セッション分離 · PII リダクションによる会話リコール · セッション アーティファクト メモリ · k-匿名化サンプリングによる外部データ ソース。
データはあなたの明示的な許可なしに離れることはありません
3 層防御: テナントごとのエンジン

llowlist → データ分類マトリックス (PUBLIC / INTERNAL / CONFIDENTIAL / SECRET) → 出力ホスト許可リスト。 EU_PRODUCTION プリセットは箱から出してすぐに出荷されます。生データ行は LLM コンテキストには決して入りません。スキーマ + 集計統計 + 匿名化されたサンプルのみが含まれます。
Web コンソール — ブラウザからすべてを管理
http://localhost:8765 のコントロール プレーン。単一のダッシュボードからセッション、ペルソナ、ブリッジ、Forge ツール、監査ログを管理します。 5 スコープのテナント モデル: 1 つのインスタンスが複数のユーザー、プロジェクト、チームを完全に分離して処理します。完全な REST API は /v1/console/ にあります。
Bridge.sh console # Web コンソールを起動します
Bridge.sh 医師 # ヘルスチェック + 監査検証
建築
7 つのブリッジ デーモン (WhatsApp、Telegram、Discord、Slack、Email、Teams、Signal) がメッセージを共有受信トレイに集めます。ブリッジ アダプターは、ACL を適用し、適切なペルソナにルーティングし、TTS パイプラインを実行し、チャットごとに順次、クロスチャットに並行してスキルを評価します。 WorkerEngine 抽象化は、コンプライアンス スタックに触れることなく、LLM バックエンドを交換します。
完全なレイヤーの内訳: docs/layer-model.md · アーキテクチャ図: docs/diagrams/ · 完全なドキュメント: docs/overview.md
bash オペレーター/bridges/run-all-tests.sh
テストは、Python アダプター、ノード デーモンブート スモーク テスト、cowork、forge、skill-forge、およびすべてのセキュリティ層に及びます。テストは密閉的に実行されます。クロードは ADAPTER_FAKE_CLAUDE=1 を介してスタブされ、名前空間の分離がテスト対象となる実際の bwrap になります。
プル リクエストを開くと、 CLA.md を受け入れることになります。マージされたすべてのコントリビューションには、 CLA-SIGNATORIES.md 内の対応するエントリが必要です。完全なワークフローについては、CONTRIBUTING.md を参照してください。
Apache License、バージョン 2.0 に基づいてライセンスされています。
再ライセンス権 (CLA §3): メンテナは、source-availa を含む、別のライセンスに基づいて CorvinOS の将来のバージョンをリリースする権利を保持します。

ble ライセンス (ビジネス ソース ライセンス、機能ソース ライセンス、エラスティック ライセンス v2) または商用ライセンス — 寄稿者からのさらなる同意は必要ありません。この権利は、 CLA.md の条件としてすべての寄稿者に付与されます。すでに公開されている Apache-2.0 リリースは影響を受けません。それらは永遠に Apache-2.0 のままです。全規約については、CLA.md § 3 を参照してください。
「CorvinOS」および「Corvin」は、Apache § 6 に基づくプロジェクト識別子です。ライセンスは商標権を付与するものではありません。
セルフホスト型エージェント OS — Ollama、Claude、GPT-4 を Discord、Telegram、WhatsApp、Slack、Email に接続します。 EU AI 法 2026 とアーキテクチャ別の GDPR 準拠。 pip インストールコルビノス
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
2
v0.9.60
最新の
2026 年 7 月 1 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted agentic OS — connects Ollama, Claude & GPT-4 to Discord, Telegram, WhatsApp, Slack & Email. EU AI Act 2026 & GDPR compliance by architecture. pip install corvinos - CorvinLabs/CorvinOS

GitHub - CorvinLabs/CorvinOS: Self-hosted agentic OS — connects Ollama, Claude & GPT-4 to Discord, Telegram, WhatsApp, Slack & Email. EU AI Act 2026 & GDPR compliance by architecture. pip install corvinos · GitHub
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
CorvinLabs
/
CorvinOS
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
53 Commits 53 Commits .claude-plugin .claude-plugin .github .github .ldd .ldd .well-known .well-known assets assets compliance compliance core core corvinOS corvinOS docs docs operator operator ops ops scripts scripts site site skills_basis skills_basis tests tests .corvin_repo .corvin_repo .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CCLA.md CCLA.md CHANGELOG.md CHANGELOG.md CLA-SIGNATORIES.md CLA-SIGNATORIES.md CLA.md CLA.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md INSTALLATION.md INSTALLATION.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md build_wheels.py build_wheels.py corvinOS_path_fix.pth corvinOS_path_fix.pth hatch_build.py hatch_build.py install.ps1 install.ps1 install.sh install.sh ollama-manifest.yaml ollama-manifest.yaml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Overview ·
Architecture ·
Audit & Compliance ·
A2A Network ·
Engine Layer ·
Security ·
EU AI Act ·
Learning Objectives
One install. Seven bridges. Any LLM.
CorvinOS is a self-hosted agentic OS that connects Ollama, Claude, GPT-4, and any OpenRouter model to Discord, Telegram, WhatsApp, Slack, Email, Teams, and Signal — through a single pip package.
pip install corvinos && python -m corvinOS
Local-first — run 100 % offline with Ollama and --engine hermes . No API key needed.
Agentic — generates sandboxed tools and new skills at runtime; delegates subtasks across five AI engines.
Compliance by architecture — EU AI Act 2026 (disclosure, consent, house-rules) and GDPR (audit chain, data residency, erasure) are load-bearing code, not policy documents. None can be disabled by a flag.
Multi-tenant — one instance, multiple users, personas, and teams, all isolated.
Self-hostable anywhere — Linux, macOS, Windows; Docker-ready; single pip install .
CorvinOS enforces the EU AI Act in code, not in documentation.
Every compliance requirement — disclosure, consent, audit integrity, data residency, egress control, GDPR erasure — is a structural constraint that cannot be disabled by a flag, env var, or config override. Regulated deployments get verifiable guarantees, not policy promises.
See INSTALLATION.md for the complete setup guide.
Recommended — works identically on Linux, macOS, and Windows:
pip install corvinos
python -m corvinOS # web console at http://localhost:8765
python -m corvinOS is PATH-independent : it starts the console on the first
try on every OS — including Microsoft Store / system Python, where pip install
falls back to a per-user scripts directory that is not on PATH (the usual reason
corvin-serve is "not found" on Windows). On Windows you can also use py -m corvinOS .
Want the short corvin-serve command on your PATH? Install with
pipx — it isolates the app and wires up PATH
automatically, on every platform:
pipx install corvinos
corvin-serve # web console at http://localhost:8765
corvin-serve from a plain pip install only works once its scripts directory
is on PATH . Running python -m corvinOS once adds that directory to your PATH,
so corvin-serve then works in a new terminal — but pipx (or python -m corvinOS ) is the reliable cross-platform path.
The base install is pure-Python and cross-platform — it brings the web console
all the way up to setup on Linux, macOS, and Windows, with cloud/edge voice
(OpenAI + Microsoft Edge TTS) working out of the box. For local, offline
speech models add the optional extra:
pip install " corvinos[voice] " # local Piper TTS + faster-whisper STT
The voice extra is opt-in because its local-model dependencies ( piper-tts ,
faster-whisper ) lack Windows wheels for some Python versions; keeping them
out of the base install means pip install corvinos reaches setup reliably on
every supported platform.
Requirements: Python 3.10+ · Linux, macOS 12+, or Windows 10/11 · Node.js 20+ required only for bridges
Default engine: Claude Code (Claude Pro or Max subscription required).
For fully local, zero-egress deployment: --engine hermes (Ollama, no API key).
# macOS / Linux — Homebrew
brew tap CorvinLabs/corvinos && brew install corvinos
# Windows — Scoop
scoop bucket add corvinos https://github.com/CorvinLabs/scoop-corvinos && scoop install corvinos
# conda / mamba (pending review)
conda install -c conda-forge corvinos
# Developer install from source
git clone https://github.com/CorvinLabs/CorvinOS.git && cd CorvinOS
pip install -e " .[all] " && corvin-install
Full documentation: docs/overview.md
One command removes everything — services, config, data directories, and the package:
corvin-uninstall --purge # removes all files without prompting
pip uninstall corvinos -y # removes the Python package
Without --purge the uninstaller asks before deleting data directories (audit logs, API keys, session history). Use --purge for a fully non-interactive wipe.
After pip uninstall corvinos -y the only thing left is the cloned repo directory (source installs) — delete it with rm -rf <repo> if you no longer need it.
EU AI Act 2026 + GDPR: Structural Enforcement
CorvinOS implements EU AI Act 2026 and GDPR as structural design constraints . Every feature must answer: does this weaken a compliance guarantee?
Absolute constraints — no env var, flag, or config can disable these:
disclosure is structurally locked · consent gate has no bypass · every audit event traverses the hash chain before any response · L34 blocks non-compliant engine spawns · L38 audit write failure blocks the A2A request · L44 house-rules gate has no kill-flag.
voice-audit verify # walk the full hash chain; exits 1 on any break
bridge.sh doctor # boot self-test with audit chain verification
python -m corvin_compliance_reports.cli generate processing-records # GDPR Art. 30
Full compliance reference: docs/eu-ai-act/README.md · docs/audit-and-compliance.md
Swap the LLM Without Touching the Compliance Stack
CorvinOS decouples the AI backend from the compliance runtime via the WorkerEngine protocol (L22). Every engine shares path-gate, audit chain, and artifact registration through the Tool Execution Broker — swap providers without changing your compliance setup.
Multi-Agent Mesh — CorvinOS Instances Talk to Each Other
Multiple CorvinOS instances form a decentralized agent network. Every cross-instance call carries a cryptographic signature, bidirectional attestation, nonce replay protection, and binary attachment verification. Audit-first invariant: the envelope is written to the hash chain before any response is sent.
Path-gate (write-protection) · secret vault with bwrap env-injection · sandboxed Forge tool generation · SkillForge with fail-closed linter · multi-tenant session isolation · conversation recall with PII-redaction · session artifact memory · external data sources with k-anonymised sampling.
Data Never Leaves Without Your Explicit Permission
Three-layer defence: per-tenant engine allowlist → data classification matrix (PUBLIC / INTERNAL / CONFIDENTIAL / SECRET) → egress host allowlist. EU_PRODUCTION presets ship out of the box. Raw data rows never enter the LLM context — only schema + aggregate stats + anonymised sample.
Web Console — Manage Everything From the Browser
Control plane at http://localhost:8765 . Manage sessions, personas, bridges, forge tools, and audit logs from a single dashboard. Five-scope tenant model: one instance handles multiple users, projects, and teams in full isolation. Full REST API at /v1/console/ .
bridge.sh console # start web console
bridge.sh doctor # health check + audit verify
Architecture
Seven bridge daemons (WhatsApp, Telegram, Discord, Slack, Email, Teams, Signal) funnel messages into a shared inbox. The Bridge Adapter enforces ACL, routes to the right persona, runs the TTS pipeline, and grades skills — per-chat-sequential, cross-chat-parallel. The WorkerEngine abstraction swaps the LLM backend without touching the compliance stack.
Full layer breakdown: docs/layer-model.md · Architecture diagrams: docs/diagrams/ · Full documentation: docs/overview.md
bash operator/bridges/run-all-tests.sh
Tests span the Python adapter, Node daemon-boot smoke tests, cowork, forge, skill-forge, and all security layers. Tests run hermetically — Claude stubbed via ADAPTER_FAKE_CLAUDE=1 , real bwrap where namespace isolation is the subject under test.
By opening a pull request you accept CLA.md . Every merged contribution requires a corresponding entry in CLA-SIGNATORIES.md . See CONTRIBUTING.md for the full workflow.
Licensed under the Apache License, Version 2.0 .
Relicense right (CLA §3): The Maintainer retains the right to release future versions of CorvinOS under a different license — including source-available licenses (Business Source License, Functional Source License, Elastic License v2) or a commercial license — without requiring further consent from contributors. This right is granted by every contributor as a condition of the CLA.md . Already-published Apache-2.0 releases are not affected; they remain Apache-2.0 forever. See CLA.md § 3 for the full terms.
"CorvinOS" and "Corvin" are project identifiers per Apache § 6 — the license does not grant trademark rights.
Self-hosted agentic OS — connects Ollama, Claude & GPT-4 to Discord, Telegram, WhatsApp, Slack & Email. EU AI Act 2026 & GDPR compliance by architecture. pip install corvinos
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
v0.9.60
Latest
Jul 1, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
