---
source: "https://github.com/adithyan-ak/agenthound"
hn_url: "https://news.ycombinator.com/item?id=49091338"
title: "Agenthound – Offensive security framework for AI agent infrastructure"
article_title: "GitHub - adithyan-ak/AgentHound: Offensive security framework for AI agent infrastructure - recon, credential looting, model exfiltration, poisoning, and attack-path analysis across MCP, A2A, gateways, and AI services. BloodHound for the agentic stack. · GitHub"
author: "akoffsec"
captured_at: "2026-07-28T23:51:44Z"
capture_tool: "hn-digest"
hn_id: 49091338
score: 1
comments: 0
posted_at: "2026-07-28T23:22:15Z"
tags:
  - hacker-news
  - translated
---

# Agenthound – Offensive security framework for AI agent infrastructure

- HN: [49091338](https://news.ycombinator.com/item?id=49091338)
- Source: [github.com](https://github.com/adithyan-ak/agenthound)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T23:22:15Z

## Translation

タイトル: Agenthound – AI エージェント インフラストラクチャ向けの攻撃的なセキュリティ フレームワーク
記事のタイトル: GitHub - adithyan-ak/AgentHound: AI エージェント インフラストラクチャ向けの攻撃的なセキュリティ フレームワーク - 偵察、認証情報の窃取、モデルの漏洩、ポイズニング、MCP、A2A、ゲートウェイ、AI サービスにわたる攻撃パス分析。エージェントスタック用の BloodHound。 · GitHub
説明: AI エージェント インフラストラクチャのための攻撃的なセキュリティ フレームワーク - MCP、A2A、ゲートウェイ、AI サービスにわたる偵察、認証情報の窃取、モデルの漏洩、ポイズニング、および攻撃パス分析。エージェントスタック用の BloodHound。 - adithyan-ak/AgentHound

記事本文:
GitHub - adithyan-ak/AgentHound: AI エージェント インフラストラクチャのための攻撃的なセキュリティ フレームワーク - 偵察、資格情報の略奪、モデルの漏洩、ポイズニング、MCP、A2A、ゲートウェイ、AI サービスにわたる攻撃パス分析。エージェントスタック用の BloodHound。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロ

広告をクリックしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アディシアン・アク
/
エージェントハウンド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
366 コミット 366 コミット .claude/ rules .claude/ rules .github .github コレクター コレクター docker docker docs docs モジュール モジュール スクリプト スクリプト sdk sdk サーバー サーバー test-infra test-infra testdata testdata .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh mkdocs.yml mkdocs.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント インフラストラクチャの攻撃的なセキュリティ フレームワーク
MCP · A2A · モデルゲートウェイ · 推論サーバー · ベクトルストア · MLOps · ノートブック · 12 のエージェントクライアント
クイックスタート ·
能力 ·
ライフサイクル ·
グラフモデル ·
ドキュメント ·
安全性
許可された使用のみ。 AgentHound には、読み取り専用の検出モジュールとアクティブな活用モジュールが同梱されています。自分が所有するインフラストラクチャ、または書面による評価権限のあるインフラストラクチャに対してのみ実行してください。 「安全性と認可」を参照してください。
AgentHound は、AI エージェント インフラストラクチャ用のオープンソースの攻撃的セキュリティ フレームワークです。最新のエージェント スタックのすべての層にわたって、偵察、フィンガープリンティング、資格情報の略奪、モデルファイル / システム プロンプト / インベントリの微調整、モデルの反転、ツールと命令のポイズニング、および構成インプラントの永続化などの完全な取り組みを実行し、すべての事実を 1 つの Neo4j グラフにマージし、すべてを結び付ける攻撃パスを証明します。 Agenthound は、エージェント スタックの BloodHound です。
🌐 フルスペクトルのエージェント攻撃対象領域
○

このフレームワークは、MCP、A2A、モデル ゲートウェイ、推論サーバー、ベクター ストア、MLOps、ノートブック、および 12 のエージェント クライアントなど、あらゆるレイヤーを攻撃します。不動産全体が 1 つのターゲット セットです。
🔓 ゲートウェイとサービス プレーンにわたる認証情報のインベントリ
マスクされた上流プロバイダー参照のインベントリに LiteLLM マスター キーを提供します
および支出メタデータを含むハッシュ化された仮想キー参照。実際の資格情報のみ
AgentHound が利用できる値は、サービス間の相関関係に参加します。
🧬 モデルファイル、システムプロンプト、インベントリの微調整
認証されていない Ollama 上のすべてのモデル - 名前、ダイジェスト、サイズ、
安定したモデルファイルのハッシュ、システムプロンプトの存在、および信号の微調整。生
モデルファイル、テンプレート、およびシステム プロンプトは、明示的な
オプトイン。
🔬 モデル反転 / トレーニングデータ残差抽出
純粋な Go GGUF パーサーは、フィードされた重みファイルの埋め込み行列に対して統計的反転を実行し、おそらく微調整された語彙トークンを回復し、モデルがグラフ ノードとしてトレーニングされた内容を明らかにします。
☠️ 積極的な悪用 - ツール/命令ポイズニング + 設定インプラント
ContextForge が管理する MCP ツールの説明を書き換えるか、 CLAUDE.md / .cursorrules を挿入するか、永続化のために悪意のある MCP サーバーを埋め込みます。すべてのミューテーションはデフォルトで予行演習され、プロバイダー固有の回復状態が保持されます。
🗄️ RAG、ベクターストア、ノートブックの攻撃対象領域
Qdrant コレクション、Jupyter セッション、およびノートブック ツリーをインベントリします。 Jupyter で保護された操作は、最初に認証情報なしで試行され、401/403 の後でのみオペレーターが指定したベアラー値を使用して再試行されるため、匿名アクセスは実際に成功した場合にのみ記録されます。境界ツリーの切り捨ては部分インベントリとして公開されます。
🕸️ クロスプロトコルおよび認証情報チェーンの攻撃パス
15 個のポストプロセッサーが生の事実では表示できないルートを計算します

- 認証情報チェーン、クロスプロトコル ピボット、MCP と A2A にわたる漏洩パス。
🧪 データフローとしてモデル化された間接プロンプトインジェクション
汚染の伝播として扱われるプロンプト インジェクション: 信頼できない入力ツール → 汚染された兄弟 → 影響の大きいシンク。実際のグラフ エッジとして追跡されます。
📊 検出と標準インテリジェンス
19 の事前構築された攻撃パス クエリ、35 の検出ルール、0 ～ 100 のリスク スコアリング、および差分としての再テスト - OWASP MCP / Agentic Top 10 および MITRE ATLAS に渡されます。
🧩 独自の攻撃を作成する
新しい AI サービスに対する新しい攻撃は、1 つのモジュールで実行できます。アクション インターフェイスを実装し、 register.go をドロップし、空インポートします。同じ SDK、同じライフサイクル、同じグラフ。
🎯 スタックのすべてのプレーンがターゲットです
表面
発見と在庫確認
検証/アクティブな操作
エージェントクライアント
12 個の MCP クライアント構成フォーマットと命令ファイル (CLAUDE.md、AGENTS.md、.cursorrules)
命令ポイズニングと可逆的な悪意のあるサーバー構成インプラント
MCP
Stdio および HTTP/SSE サーバー、ツール、リソース、プロンプト、認証
資格情報到達の検証。 ContextForge ツール記述ポイズニングとラウンドトリップ検証
A2A
エージェント カード、JWS 検証、スキル、委任、認証
クロスプロトコルおよび委任パスの分析
LiteLLM
オペレーター提供のマスターキーレコード、マスクされたプロバイダー参照、および支出コンテキストを含むハッシュ化された仮想キーメタデータ
サービス間の資格情報の相関関係とパス分析
オラマ / vLLM
Ollama モデルのメタデータ、安定したモデルファイル ハッシュ、システム プロンプト プレゼンス、および微調整信号。 vLLM フィンガープリンティング
オプションの生のモデルファイル、テンプレート、およびシステム プロンプト キャプチャ。ローカル GGUF 抽出
クドラント
コレクション、ポイント カウント、およびオプションの制限されたペイロード サンプル
読み取り専用の暴露分析
MLフロー
実験、実行、登録モデル、アーティファクト/ストレージ URI

、および検証された匿名暴露証拠
読み取り専用の暴露分析
ジュピター
セッションと境界付きノートブック ツリー
読み取り専用の匿名対認証型暴露分析
WebUI / LangServe を開く
オープン WebUI 認証体制に加えて、認証されたアップストリーム/RAG 資格情報インベントリと観察された暴露証拠。 LangServe フィンガープリンティング
読み取り専用の認証情報のインベントリと暴露証拠
📦 数字で見る
8 つのライフサイクル CLI コマンド - スキャン、検出、略奪、抽出、毒、移植、元に戻す、キャンペーン (スキャン内での列挙 + フィンガープリントの実行)
指紋採取者 8 名、略奪者 6 名、モデル反転抽出者 1 名、毒殺者 2 名、注入者 1 名
グラフ: 23 ノード ラベル · 32 エッジ種類 (20 生 + 12 コンポジット) · 15 ポストプロセッサ
インテリジェンス: 35 個のテキスト検出ルール + 7 個の YAML フィンガープリント ルール + 1 個のコードベースの Jupyter 検出器 · 19 個の事前構築済み攻撃パス クエリ · OWASP MCP Top 10 + OWASP Agentic Top 10 + MITRE ATLAS マッピング
DB/UI/サーバーの依存関係のない 1 つの静的コレクター バイナリ。構成のみの検出はオフラインで実行できます。 Apache-2.0 リリースには、Cosign 署名付きチェックサム マニフェストとアーカイブごとの SPDX SBOM が含まれています。
デフォルトのパスの前提条件: Docker + Compose v2。ゴーなし、ノードなし、なし
git クローン 。
1. 分析サーバーを起動します - Neo4j + Postgres + UI、バインド
127.0.0.1:8080 :
カール -sSfL https://raw.githubusercontent.com/adithyan-ak/agenthound/main/docker/docker-compose.public.yml | docker compose -f - -p Agenthound up -d --wait
2. コレクターをインストールします - 単一の静的バイナリ → ~/.local/bin :
カール -sSfL https://raw.githubusercontent.com/adithyan-ak/agenthound/main/install.sh |しー
エクスポート PATH= " $HOME /.local/bin: $PATH "
または、次のパッケージ マネージャーの代替手段のいずれかを選択します。
# Homebrew (macOS または Linux; タップが自動的に追加されます)
brew インストール adithyan-ak/agenthound/agenthound
#G

o 1.25.12+
github.com/adithyan-ak/agenthound/collector/cmd/agenthound@1.0.0 をインストールしてください
Go は GOBIN にインストールされるか、デフォルトでは $(go env GOPATH)/bin にインストールされます。それを確実にする
ディレクトリは PATH 上にあります。
3. ローカル構成をスキャンします - オフライン、読み取り専用、生の認証情報の値は省略されます。
カバレッジ レベルを 1 つ選択し、保存されたアーティファクトを取り込みます。
通常のスキャン - 最初の実行を推奨:
Agenthound scan --config --ingest http://127.0.0.1:8080
ディープ スキャン — 境界付きネストされたプロジェクト命令の検出を追加します。
Agenthound scan --config --deep --ingest http://127.0.0.1:8080
どちらのコマンドも、自宅に登録されている指示ソースと選択された指示ソースを確認します。
プロジェクトのルーツ。ターゲットが
現在のディレクトリ。ディープディスカバリは選択したプロジェクトを独立して保持します
通常剪定されている家のサブツリーの内部でも覆われています。
コレクターはアップロード前に ./scan-<scan_id>.json を保存し、コンパクトなファイルを出力します。
レシートを摂取します。完全な受信確認には --json を使用します。
4. 次の場所でグラフを開きます。
http://127.0.0.1:8080 。
スタンドアロン サーバー バイナリも次のように入手できます。
Homebrew 経由の adithyan-ak/agenthound/agenthound-server。どちらのバイナリも
リリースアーカイブから、または明示的に Go から入手可能
@1.0.0 リビジョン。リリース アーカイブには、共同署名されたチェックサム マニフェストが含まれています
およびアーカイブごとの SPDX SBOM - を参照してください。
インストールガイド。
収集コマンドは、取り込み可能な JSON を書き込みます。上記のクイックスタートは、
パターンを一度取り込みます。
1. 偵察 - AI 資産を見つけます。
一般的な AI サービス ポートをスキャンし、応答したものをフィンガープリントします。
エージェントハウンド スキャン 10.0.0.0/24
MCP および A2A プロトコル形状の可能性のある Web ポートをプローブします。
エージェントハウンド ディスカバー 10.0.0.0/24
2. 戦利品 - インベントリの認証情報の証拠とモデルのメタデータ:
LITELLM_MASTER_KEY を設定すると、LiteLLM 資格情報の参照とインベントリが作成されます。
支出メタデータ:
エージェントハウンド戦利品 10.0.0.20:4000 -

-type litellm \
--マスターキー " $LITELLM_MASTER_KEY "
未加工の Ollama モデルファイル、テンプレート、システム プロンプトをオプトインします。
エージェントハウンド戦利品 10.0.0.10:11434 --type ollam \
--include-credential-values
Looter タイプ: litellm 、 ollama 、 openwebui 、 mlflow 、 qdrant 、 jupyter 。
3. 抽出 - AI_MODEL_ID をグラフから AIModel ID に設定し、反転します
微調整残差を回復するためのローカルで利用可能な GGUF 重みファイル:
Agenthound 抽出 " $AI_MODEL_ID " --type embedding-invert \
--artifact /path/to/model.gguf --commit --engagement-id ENG-1
4. 検証、活用、永続化 + 元に戻す - 認可された、元に戻せる実行
攻撃的な行動:
ContextForge 認証を構成したら、元に戻すことのできる
マネージド MCP ツールに対するポイズニングと復元のラウンド トリップ:
エージェントハウンド キャンペーン \
https://gateway.example/servers/0123456789abcdef0123456789abcdef/mcp \
--scenario mcp-poison-roundtrip --adapter contextforge \
--target-id support-lookup --engagement-id ENG-ROUNDTRIP --commit
ターゲットを絞ったツール説明のポイズンをコミットします。
エージェントハウンドの毒 \
https://gateway.example/servers/0123456789abcdef0123456789abcdef/mcp \
--type mcp.tool.description --adapter contextforge \
--target-id support-lookup --inject-file payload.txt \
--commit --engagement-id ENG-1
悪意のある MCP サーバー エントリを埋め込み、エンゲージメントをロールバックします。

[切り捨てられた]

## Original Extract

Offensive security framework for AI agent infrastructure - recon, credential looting, model exfiltration, poisoning, and attack-path analysis across MCP, A2A, gateways, and AI services. BloodHound for the agentic stack. - adithyan-ak/AgentHound

GitHub - adithyan-ak/AgentHound: Offensive security framework for AI agent infrastructure - recon, credential looting, model exfiltration, poisoning, and attack-path analysis across MCP, A2A, gateways, and AI services. BloodHound for the agentic stack. · GitHub
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
adithyan-ak
/
AgentHound
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
366 Commits 366 Commits .claude/ rules .claude/ rules .github .github collector collector docker docker docs docs modules modules scripts scripts sdk sdk server server test-infra test-infra testdata testdata .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh mkdocs.yml mkdocs.yml View all files Repository files navigation
The offensive security framework for AI agent infrastructure
MCP · A2A · model gateways · inference servers · vector stores · MLOps · notebooks · 12 agent clients
Quickstart ·
Capabilities ·
Lifecycle ·
Graph Model ·
Docs ·
Safety
Authorized use only. AgentHound ships read-only discovery and active exploitation modules. Run it only against infrastructure you own or are written-authorized to assess. See Safety & Authorization .
AgentHound is an open-source offensive security framework for AI agent infrastructure. It runs the full engagement - recon, fingerprinting, credential looting, modelfile / system-prompt / fine-tune inventory , model inversion, tool and instruction poisoning, and config-implant persistence - across every layer of the modern agentic stack, then merges every fact into one Neo4j graph and proves the attack paths that tie it all together. Agenthound is BloodHound for the agentic stack.
🌐 Full-spectrum agentic attack surface
One framework attacks every layer - MCP, A2A, model gateways, inference servers, vector stores, MLOps, notebooks, and 12 agent clients. The whole estate is one target set.
🔓 Credential inventory across the gateway & service plane
Supply a LiteLLM master key to inventory masked upstream-provider references
and hashed virtual-key references with spend metadata. Only actual credential
values available to AgentHound participate in cross-service correlation.
🧬 Modelfile, system-prompt & fine-tune inventory
Enumerate every model on an unauthenticated Ollama - names, digests, sizes,
stable modelfile hashes, system-prompt presence, and fine-tune signals. Raw
modelfiles, templates, and system prompts are available through an explicit
opt-in.
🔬 Model inversion / training-data residue extraction
A pure-Go GGUF parser runs statistical inversion on the embedding matrix of any weight file you feed it to recover likely fine-tune vocabulary tokens - surfacing what a model was trained on as graph nodes.
☠️ Active exploitation - tool/instruction poisoning + config implant
Rewrite a ContextForge-managed MCP tool description, inject CLAUDE.md / .cursorrules , or implant a malicious MCP server for persistence. Every mutation is dry-run by default and carries provider-specific recovery state.
🗄️ RAG, vector-store & notebook attack surface
Inventory Qdrant collections and Jupyter sessions and notebook trees. Jupyter protected operations are tried without credentials first and retried with an operator-supplied bearer value only after a 401/403, so anonymous access is recorded only when it actually succeeds; bounded tree truncation is published as partial inventory.
🕸️ Cross-protocol & credential-chain attack paths
15 post-processors compute the routes raw facts can't show - credential chains, cross-protocol pivots, and exfiltration paths across MCP and A2A.
🧪 Indirect prompt injection, modeled as data-flow
Prompt injection treated as taint propagation: untrusted-input tools → tainted siblings → high-impact sinks, traced as real graph edges.
📊 Detection & standards intelligence
19 prebuilt attack-path queries, 35 detection rules, 0–100 risk scoring, and retest-as-diff - crosswalked to OWASP MCP / Agentic Top 10 and MITRE ATLAS.
🧩 Write your own attacks
A new attack against a new AI service is one module away - implement an action interface, drop a register.go , blank-import it. Same SDK, same lifecycle, same graph.
🎯 Every plane of the stack is a target
Surface
Discovery & inventory
Validation / active operations
Agent clients
12 MCP client config formats plus instruction files ( CLAUDE.md , AGENTS.md , .cursorrules )
Instruction poisoning and reversible malicious-server config implants
MCP
Stdio and HTTP/SSE servers, tools, resources, prompts, and authentication
Credential-reach verification; ContextForge tool-description poisoning and round-trip validation
A2A
Agent cards, JWS verification, skills, delegation, and authentication
Cross-protocol and delegation-path analysis
LiteLLM
Operator-supplied master-key record, masked provider references, and hashed virtual-key metadata with spend context
Cross-service credential correlation and path analysis
Ollama / vLLM
Ollama model metadata, stable modelfile hashes, system-prompt presence, and fine-tune signals; vLLM fingerprinting
Optional raw modelfile, template, and system-prompt capture; local GGUF extraction
Qdrant
Collections, point counts, and optional bounded payload samples
Read-only exposure analysis
MLflow
Experiments, runs, registered models, artifact/storage URIs, and verified anonymous-exposure evidence
Read-only exposure analysis
Jupyter
Sessions and bounded notebook trees
Read-only anonymous-versus-authenticated exposure analysis
Open WebUI / LangServe
Open WebUI authentication posture plus authenticated upstream/RAG credential inventory and observed exposure evidence; LangServe fingerprinting
Read-only credential inventory and exposure evidence
📦 By the numbers
8 lifecycle CLI commands - scan · discover · loot · extract · poison · implant · revert · campaign ( enumerate + fingerprint run inside scan )
8 fingerprinters · 6 looters · 1 model-inversion extractor · 2 poisoners · 1 implanter
Graph: 23 node labels · 32 edge kinds (20 raw + 12 composite) · 15 post-processors
Intelligence: 35 text-detection rules + 7 YAML fingerprint rules + 1 code-backed Jupyter detector · 19 prebuilt attack-path queries · OWASP MCP Top 10 + OWASP Agentic Top 10 + MITRE ATLAS mappings
One static collector binary with no DB/UI/server dependencies. Config-only discovery can run offline. Apache-2.0 releases include a Cosign-signed checksum manifest and per-archive SPDX SBOMs.
Default path prerequisites: Docker + Compose v2. No Go, no Node, no
git clone .
1. Start the analysis server - Neo4j + Postgres + UI, binds
127.0.0.1:8080 :
curl -sSfL https://raw.githubusercontent.com/adithyan-ak/agenthound/main/docker/docker-compose.public.yml | docker compose -f - -p agenthound up -d --wait
2. Install the collector - single static binary → ~/.local/bin :
curl -sSfL https://raw.githubusercontent.com/adithyan-ak/agenthound/main/install.sh | sh
export PATH= " $HOME /.local/bin: $PATH "
Or choose one of these package-manager alternatives:
# Homebrew (macOS or Linux; adds the tap automatically)
brew install adithyan-ak/agenthound/agenthound
# Go 1.25.12+
go install github.com/adithyan-ak/agenthound/collector/cmd/agenthound@1.0.0
Go installs into GOBIN or, by default, $(go env GOPATH)/bin ; ensure that
directory is on PATH .
3. Scan local configs - offline, read-only, raw credential values omitted.
Choose one coverage level and ingest the saved artifact.
Normal scan — recommended first run:
agenthound scan --config --ingest http://127.0.0.1:8080
Deep scan — adds bounded nested-project instruction discovery:
agenthound scan --config --deep --ingest http://127.0.0.1:8080
Both commands check registered instruction sources at your home and selected
project roots. Add --project-dir /path/to/project when the target is not the
current directory. Deep discovery keeps that selected project independently
covered even inside a normally pruned home subtree.
The collector saves ./scan-<scan_id>.json before upload, then prints a compact
ingest receipt. Use --json for the full receipt.
4. Open the graph at
http://127.0.0.1:8080 .
The standalone server binary is also available as
adithyan-ak/agenthound/agenthound-server through Homebrew. Both binaries are
available from release archives, or from Go at the explicit
@1.0.0 revision. Release archives include a Cosign-signed checksum manifest
and per-archive SPDX SBOMs - see the
installation guide .
Collection commands write ingest-ready JSON. The quickstart above shows the
ingest pattern once.
1. Recon - find the AI estate:
Scan common AI-service ports and fingerprint what responds:
agenthound scan 10.0.0.0/24
Probe likely web ports for MCP and A2A protocol shapes:
agenthound discover 10.0.0.0/24
2. Loot - inventory credential evidence and model metadata:
With LITELLM_MASTER_KEY set, inventory LiteLLM credential references and
spend metadata:
agenthound loot 10.0.0.20:4000 --type litellm \
--master-key " $LITELLM_MASTER_KEY "
Opt in to raw Ollama modelfiles, templates, and system prompts:
agenthound loot 10.0.0.10:11434 --type ollama \
--include-credential-values
Looter types: litellm , ollama , openwebui , mlflow , qdrant , jupyter .
3. Extract - with AI_MODEL_ID set to an AIModel ID from the graph, invert
a locally-available GGUF weight file to recover fine-tune residue:
agenthound extract " $AI_MODEL_ID " --type embedding-invert \
--artifact /path/to/model.gguf --commit --engagement-id ENG-1
4. Validate, exploit, persist + revert - run sanctioned, reversible
offensive actions:
With ContextForge authentication configured, run a reversible
poison-and-restore round trip against a managed MCP tool:
agenthound campaign \
https://gateway.example/servers/0123456789abcdef0123456789abcdef/mcp \
--scenario mcp-poison-roundtrip --adapter contextforge \
--target-id support-lookup --engagement-id ENG-ROUNDTRIP --commit
Commit a targeted tool-description poison:
agenthound poison \
https://gateway.example/servers/0123456789abcdef0123456789abcdef/mcp \
--type mcp.tool.description --adapter contextforge \
--target-id support-lookup --inject-file payload.txt \
--commit --engagement-id ENG-1
Implant a malicious MCP server entry, then roll the engagement back:

[truncated]
