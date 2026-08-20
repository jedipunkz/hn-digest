---
source: "https://github.com/yasuoiwakura/openai-ollama-api-bridge"
hn_url: "https://news.ycombinator.com/item?id=49379160"
title: "A proxy to translate OpenCode OpenAI calls to native Ollama, allows ctx >4096"
article_title: "GitHub - yasuoiwakura/openai-ollama-api-bridge: A proxy to translate Opencode openai API calls to native ollama, allow context size to forward and even overwrite to be >4096 · GitHub"
image: "https://opengraph.githubassets.com/27923166ee9f7499e11c61dfb295793b35fb50a9ce3215e9a6f89b38fa9d2f59/yasuoiwakura/openai-ollama-api-bridge"
author: "MK2k"
captured_at: "2026-08-20T20:18:08Z"
capture_tool: "hn-digest"
hn_id: 49379160
score: 2
comments: 1
posted_at: "2026-08-20T19:40:00Z"
tags:
  - hacker-news
  - translated
---

# A proxy to translate OpenCode OpenAI calls to native Ollama, allows ctx >4096

- HN: [49379160](https://news.ycombinator.com/item?id=49379160)
- Source: [github.com](https://github.com/yasuoiwakura/openai-ollama-api-bridge)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T19:40:00Z

## Translation

タイトル: OpenCode OpenAI 呼び出しをネイティブ Ollama に変換するプロキシ、ctx >4096 を許可
記事タイトル: GitHub - yasuoiwaakura/openai-ollama-api-bridge: Opencode openai API 呼び出しをネイティブ ollama に変換するプロキシ、コンテキスト サイズを転送し、>4096 に上書きすることも可能 · GitHub
説明: Opencode の openai API 呼び出しをネイティブの ollama に変換するプロキシ。コンテキスト サイズの転送と 4096 を超える上書きを可能にします - yasuoiwaakura/openai-ollama-api-bridge

記事本文:
GitHub - yasuoiwaakura/openai-ollama-api-bridge: Opencode openai API 呼び出しをネイティブ ollama に変換するプロキシ。コンテキスト サイズを転送し、さらには 4096 を超えるように上書きできるようにします · GitHub
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
ヤスオイワクラ
/
オープンナイ・オラマ・APIブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
6 コミット 6 コミット フォルダーとファイル
.dockerignore .dockerignore .env.example .env.example .gitignore .gitignore

Dockerfile Dockerfile LICENSE LICENSE Notes.md Notes.md README.md README.md SPEC.md SPEC.mdcollect_fixture.pycollect_fixture.py compose.yml compose.yml proxy.py proxy.pyrequirements.txtrequirements.txt translator.py translator.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI /v1/chat/completions リクエストを Ollama のネイティブ /api/chat API に変換したり、逆に変換したりする軽量の HTTP プロキシ。
Ollama の OpenAI 互換エンドポイント ( /v1/chat/completions ) には、ランタイム構成に重大な制限があります。
カスタム Modelfile がないと、コンテキスト サイズなどのパラメーターを一元的に制御できません。 OpenAI 互換性レイヤーを介したリクエストは、必要なランタイム設定ではなく Ollama のデフォルトで実行される場合があります。
一般的な回避策は、次のような設定で単一モデルごとにカスタム モデルファイルを作成することです。
パラメータ num_ctx 64000
これにより、重複したモデル定義が作成されます。
llama3.2 → ベースモデル + カスタムモデルファイル
qwen3 → ベースモデル + カスタムモデルファイル
gemma3 → 基本モデル + カスタム Modelfile
ミストラル → ベースモデル + カスタムモデルファイル
deepseek-r1 → 基本モデル + カスタム Modelfile
20 モデル = 20 の追加モデルファイル。
後でコンテキスト サイズを変更すると、すべてのコンテキストを再度再構築または更新することになります。
問題はモデル自体ではなく、中央のランタイム構成レイヤーが欠落していることです。
Ollama モデルファイルを作成しない場合、デフォルトの OpenCode + Ollama セットアップには重大な互換性制限があります
フローチャート LR
サブグラフ OpenCode
OC_OpenAI["OpenAI コネクタ"]
サブグラフ JSON["opencode.jsonc"]
JMAX["max_tokens<br />=64K"]
JMAXLEN["最大応答長"]
J_MODEL["基本モデル"]
終わり
サブグラフ OC_POST["POST"]
OC_COMPLETE["/v1/チャット/コンプリート"]
MODEL[「基本モデル」]
MSG["メッセージ"]
MAX["max_tokens<br />=32K"]
ストリーム["ストリーム"]
終わり
終わり
部分グラフ オラマ
CTX_DEFAULT["デフォルト:<br />num_ctx=

4K」]
サブグラフ OL_OpenAI["OpenAI 互換性 API"]
OL_COMPLETE["/v1/チャット/コンプリート"]
OL_MODEL["基本モデル"]
CMSG["メッセージ"]
CMAX["max_tokens<br />=32K"]
CSTREAM["ストリーム"]
終わり
サブグラフ Native["ネイティブ Ollama API"]
チャット["POST /api/chat"]
GEN["POST /api/generate"]
終わり
MODELS[(モデルファイルなし**の任意のモデル)]
終わり
J_MODEL --> モデル
JMAX -.->|"🔒<br />ハードコード<br />32K"|マックス
JMAXLEN -.->|"❌<br />無視されました"|マックス
OC_OpenAI -.->|"❌<br />使用されていません"|ネイティブ
OC_COMPLETE -.->|"🚧限定"| OL_COMPLETE
OC_OpenAI --> OC_POST
モデル --> OL_MODEL
MSG --> CMSG
MAX --> CMAX
ストリーム --> Cストリーム
OL_MODEL --デフォルト モデル<br/>num_ctx=4K<br/>モデルファイルなし --> モデル
CMSG --> モデル
CMAX -.->|"❌<br />無視されました"| CTX_DEFAULT
CTX_DEFAULT -->|"カスタム モデルファイルなし"|モデル
Cストリーム --> モデル
チャット --> モデル
世代 --> モデル
スタイル OC_COMPLETE 塗りつぶし:#ffe7aa
スタイル OL_COMPLETE 塗りつぶし:#ffe7aa
スタイルチャット記入:#b5f5b5
スタイル GEN フィル:#b5f5b5
スタイル JMAX 塗りつぶし:#b5f5b5
スタイル MAX フィル:#ffe7aa
スタイル CMAX 塗りつぶし:#ffaaaa
スタイル J_MODEL 塗りつぶし:#b5f5b5
スタイル OL_MODEL 塗りつぶし:#b5f5b5
スタイルモデル塗りつぶし:#b5f5b5
読み込み中
モデルごとに別のカスタム Modelfile を作成する必要があるのでしょうか?
ブリッジを使用しない唯一の実際的な回避策は、ランタイム パラメーターをオーバーライドするためだけに、モデルごとにカスタム Ollama Modelfile を作成することです。
これは冗長で時間がかかり、モデルを頻繁にテストしたり切り替えたりする場合に維持するのが困難です。
フローチャート TD
M1[「ラマ3.2<br/>モデル」]
F11["ベースモデル<br/>llama3.2"]
F12["モデルファイル<br/>llama3.2<br/>num_ctx=32K"]
M2["qwen3<br/>モデル"]
F21["ベースモデル<br/>qwen3"]
F22["モデルファイル<br/>qwen3<br/>num_ctx=32K"]
M3["gemma3<br/>モデル"]
F31["ベースモデル<br/>gemma3"]
F32["モデルファイル<br/>gemma3<br/>num_ctx=32K"]
M4[「ミストラル<br/>モデル」]
F41["ベースモデル<br/>ミストラル"]
F42["モデルファイル<br/>ミストラル<br/>num_ctx=32K"]
M5["深さ

ek-r1<br/>モデル」]
F51["ベースモデル<br/>deepseek-r1"]
F52["モデルファイル<br/>ディープシーク-r1<br/>num_ctx=32K"]
M1 --> F11
M1 --> F12
M2 --> F21
M2 --> F22
M3 --> F31
M3 --> F32
M4 --> F41
M4 --> F42
M5 --> F51
M5 --> F52
スタイル M1 塗りつぶし:#b5f5b5
スタイル M2 塗りつぶし:#b5f5b5
スタイル M3 塗りつぶし:#b5f5b5
スタイル M4 塗りつぶし:#b5f5b5
スタイル M5 塗りつぶし:#b5f5b5
スタイル F12 塗りつぶし:#ffe7aa
スタイル F22 塗りつぶし:#ffe7aa
スタイル F32 塗りつぶし:#ffe7aa
スタイル F42 塗りつぶし:#ffe7aa
スタイル F52 塗りつぶし:#ffe7aa
読み込み中
このプロキシの機能
Ollama のネイティブ /api/chat エンドポイントを直接使用し、ランタイム パラメーターが明示的にサポートされます。
1 つの .env ファイルですべてのモデルを一元的に構成します。
ランタイム設定を適用する中央 API ブリッジ
単一の OpenAI エンドポイントを提供しながら、ランタイム ポリシーを一元的に適用し、プライマリ サーバーが利用できなくなった場合にはセカンダリ Ollama サーバーに自動的に切り替えます。
フローチャート LR
クライアント["OpenCode<br/>OpenAI コネクタ"]
サブグラフ Bridge["API ブリッジ"]
TRANS["OpenAI → オラマ<br/>翻訳"]
CONTEXT["コンテキスト サイズ ポリシーを<br/>強制する"]
FAILOVER[「他のサーバーに切り替える」]
終わり
O1[「オラマサーバー」]
クライアント -->|"OpenAI API を話します"|橋
ブリッジ -->|"Ollama API を話します"| O1
スタイル クライアント入力:#ffaaaa
スタイル ブリッジ フィル:#88FFFF
スタイル O1 塗りつぶし:#b5f5b5
読み込み中
高可用性/フェイルオーバー
単一の OpenAI エンドポイントを提供し、プライマリ サーバーが使用できなくなった場合はセカンダリ Ollama サーバーに自動的に切り替えます。
フローチャート LR
クライアント["OpenCode<br/>OpenAI コネクタ"]
ブリッジ["OpenAI ↔ Ollama API ブリッジ"]
決定{「プライマリは利用可能ですか?」}
C1["コンテキスト サイズ = 128K"]
C2["コンテキスト サイズ = 32K"]
O1[「プライマリ Ollama サーバー」]
O2[「セカンダリ Ollama サーバー」]
クライアント --> ブリッジ
ブリッジ --> 決定
決定 -->|はい| C1
決定 -->|いいえ| C2
C1 --> O1
C2 --> O2
スタイル ブリッジフィル:#ffff88
スタイル O1 塗りつぶし:#b5f5b5
スタイル O2 フィル:#b5f5b5
スタイル C1 塗りつぶし:#e8f4fd
スタイルC

2 フィル:#88FFFF
読み込み中
クイックスタート
git clone <リポジトリ>
cd py-ollam-openai-bridge
cp .env.example .env
# .env を編集: OLLAMA_URL を Ollama ホストに設定します
ドッカー構成 -d
ブリッジは次の場所で入手できます。
http://localhost:8080/v1
直接 (Docker なし)
pip install -r 要件.txt
# .env を設定する
Pythonプロキシ.py
OpenCode をブリッジに向けます。
{
「プロバイダー」: {
"オラマブリッジ" : {
"name" : " オラマ (ブリッジ) " ,
"npm" : " @ai-sdk/openai-compatibility " ,
"オプション" : {
"baseURL" : " http://localhost:8080/v1 "
}、
「モデル」: {
"gpt-oss:20b" : {
"名前" : " _gpt-oss:20b " ,
"tool_call" : true 、
「制限」: {
「コンテキスト」: 64000 、
「出力」：40960
}
}
}
}
}
}
すべてのオプション
NUM_CTX = 64000
NUM_PREDICT = 128000
温度 = 0.7
TOP_P = 0.9
TOP_K = 40
MIN_P = 0.05
REPEAT_PENALTY = 1.1
シード = 42
キープアライブ = 30m
.env 内のコメントされていない値のみが挿入されます。
未設定の値は Ollama のデフォルトを使用します。
OLLAMA_URL → FAILOVER_OLLAMA_URL
例:
OLLAMA_URL = http://192.168.0.42:11434
FAILOVER_OLLAMA_URL = http://192.168.0.101:11434
FAILOVER_NUM_CTX = 96000
FAILOVER_NUM_PREDICT = 128000
FAILOVER_KEEP_ALIVE = 5m
特徴:
GET /api/tags によるヘルスチェック
プライマリが使用できない場合の自動切り替え
接続エラーとタイムアウト時のフェイルオーバー
独立したフェイルオーバーパラメータ
シンプルモード: OLLAMA_URL のみが必要です
特徴
Ollama を使用した OpenCode
APIブリッジ
OpenAI チャットの完了
✅
✅
OpenAI レスポンス API
⚠️限定
✅
ネイティブ/API/チャット
❌
✅
ネイティブ /api/generate
❌
✅
ストリーミング
✅
✅
ツール呼び出し
バックエンドに依存
保存済み
複数の Ollama サーバー
❌ 手動スイッチ
✅
コンテキストサイズポリシー
❌ モデルごとのモデルファイル
✅ 一元化されたランタイムポリシー
自動フェイルオーバー
❌
✅
ルーティングポリシー
❌
✅
単一の安定したエンドポイント
❌
✅
ランタイム構成モデル
モデル固有の Modelfile を作成する代わりに、次のようにします。
ラマ3.2+

モデルファイル
qwen3 + モデルファイル
gemma3 + モデルファイル
ミストラル + モデルファイル
ディープシーク + モデルファイル
ブリッジは実行時ポリシーを一元的に適用します。
.env
|
+-- NUM_CTX
+-- NUM_PREDICT
+-- 温度
+-- TOP_P
+-- TOP_K
+-- キープアライブ
|
v
オラマの全モデル
これにより、モデル定義を再構築せずに実行時の動作を変更できます。
OpenAI 互換クライアントを変更しないでください
Ollama のネイティブ API 機能を使用する
ランタイム構成を一元化する
複数の Ollama バックエンドをサポートする
利用可能な場合はツール呼び出しと推論メタデータを保存します
プライマリサーバーを継続的にチェックすることで、フェイルオーバーの 2 秒の遅延を防止します
ヘルスチェックを60秒間キャッシュするだけです
私の実際のユースケースは、外部から GPU クラスターにアクセスすることです。そこでフェイルオーバー機能も実装しました。
API ゲートウェイとリバース プロキシを使用すると、この小さなセットアップには多大な効果があるように見えますが、私は Authentik と TraefiK/Caddy を使用して Homelab WebApps にアクセスし、Kong を使用して Homelab API にアクセスします。
TLS 接続を処理する層を 1 つだけにする必要があるため、Caddy (Homelab) または Traefik (Rootserver) が常に関与します。
フローチャート TD
サブグラフ ラップトップ
OC["OpenCode クライアント"]
終わり
サブグラフ Homelan["家庭内 LAN"]
サブグラフ ルーター
IP["パブリック IPV4"]
ポートフォワーディング["NAT フォワード ポート 443"]
終わり
サブグラフ OptiPlex["Optiplex"]
サブグラフ PVE["Proxmox クラスター"]
サブグラフ LXC["LXC HomeLab コンテナ"]
サブグラフ Docker["Docker Compose"]
ReverseProxy["Traefik/Caddy<br />リバース プロキシ<br/>TLS 終了"]
Kong["Kong API ゲートウェイ<br />JWT 認証"]
ブリッジ[「API ブリッジ<br />このプロジェクト」]
終わり
終わり
終わり
終わり
サブグラフ OldLaptop["Linux ラップトップ"]
サブグラフ Debian["Debian Linux"]
ROCM["Rocm/バルカンドライバー"]
サブグラフ Debian_Docker["Docker Compose"]
ollam_docker["Ollama サーバー"]
終わり
終わり
終わり
RX[「AMD RX 6900XT<br/>16GB VRAM」]
サブグラフ ゲーム["ゲーム機"]
サブグラフ Win11["Windows 11"]
CUDA["CUDAドライバー"]
ollam_e

xe["ollam.exe"]
終わり
RTX[「Nvidia RTX 5090<br/>32GB VRAM」]
終わり
終わり
OC -->|インターネット| IP
IP --> ポートフォワーディング
ポートフォワーディング --> ReverseProxy
ReverseProxy --> コング
コング --> ブリッジ
ブリッジ -->|"プライマリ<br />num_ctx=128K"| ollam_exe
RTX -.- CUDA
CUDA -.- ollama_exe
RX -.-|"ライザーケーブル"| ROCM
ROCM -.- ollam_docker
ブリッジ -.->|"フェイルオーバー<br />num_ctx=32K"|オラマドッカー
スタイル OptiPlex 塗りつぶし: #88FFFF
スタイル OldLaptop 塗りつぶし: #88FFFF
スタイル ゲームフィル: #88FFFF
スタイル ブリッジ フィル:#b5f5b5
読み込み中
ライセンス
Opencode openai API 呼び出しをネイティブ ollam に変換するプロキシ。コンテキスト サイズを転送したり、4096 を超えるように上書きしたりすることができます。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A proxy to translate Opencode openai API calls to native ollama, allow context size to forward and even overwrite to be >4096 - yasuoiwakura/openai-ollama-api-bridge

GitHub - yasuoiwakura/openai-ollama-api-bridge: A proxy to translate Opencode openai API calls to native ollama, allow context size to forward and even overwrite to be >4096 · GitHub
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
yasuoiwakura
/
openai-ollama-api-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
.dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE NOTES.md NOTES.md README.md README.md SPEC.md SPEC.md collect_fixture.py collect_fixture.py compose.yml compose.yml proxy.py proxy.py requirements.txt requirements.txt translator.py translator.py View all files Repository files navigation
A lightweight HTTP proxy that translates OpenAI /v1/chat/completions requests to Ollama's native /api/chat API and back.
Ollama's OpenAI-compatible endpoint ( /v1/chat/completions ) has serious runtime configuration limitations.
Without custom Modelfiles, parameters such as context size cannot be centrally controlled. Requests through the OpenAI compatibility layer may run with Ollama defaults instead of the desired runtime settings.
The typical workaround is creating custom Modelfiles for every single model with settings like:
PARAMETER num_ctx 64000
This creates duplicated model definitions:
llama3.2 → base model + custom Modelfile
qwen3 → base model + custom Modelfile
gemma3 → base model + custom Modelfile
mistral → base model + custom Modelfile
deepseek-r1 → base model + custom Modelfile
20 models = 20 additional Modelfiles.
Changing the context size later means rebuilding or updating all of them again.
The problem is not the models themselves — it is the missing central runtime configuration layer.
Without creating Ollama Modelfiles, the default OpenCode + Ollama setup has serious compatibility limitations
flowchart LR
subgraph OpenCode
OC_OpenAI["OpenAI Connector"]
subgraph JSON["opencode.jsonc"]
JMAX["max_tokens<br />=64K"]
JMAXLEN["max response length"]
J_MODEL["basic model"]
end
subgraph OC_POST["POST"]
OC_COMPLETE["/v1/chat/completions"]
MODEL["basic model"]
MSG["messages"]
MAX["max_tokens<br />=32K"]
STREAM["stream"]
end
end
subgraph Ollama
CTX_DEFAULT["Default:<br />num_ctx=4K"]
subgraph OL_OpenAI["OpenAI Compatibility API"]
OL_COMPLETE["/v1/chat/completions"]
OL_MODEL["basic model"]
CMSG["messages"]
CMAX["max_tokens<br />=32K"]
CSTREAM["stream"]
end
subgraph Native["Native Ollama APIs"]
CHAT["POST /api/chat"]
GEN["POST /api/generate"]
end
MODELS[(Any Model **without** Modelfile)]
end
J_MODEL --> MODEL
JMAX -.->|"🔒<br />hardcoded<br />32K"| MAX
JMAXLEN -.->|"❌<br />ignored"| MAX
OC_OpenAI -.->|"❌<br />not used"| Native
OC_COMPLETE -.->|"🚧limited"| OL_COMPLETE
OC_OpenAI --> OC_POST
MODEL --> OL_MODEL
MSG --> CMSG
MAX --> CMAX
STREAM --> CSTREAM
OL_MODEL --Default Model<br/>num_ctx=4K<br/>without Modelfile--> MODELS
CMSG --> MODELS
CMAX -.->|"❌<br />ignored"| CTX_DEFAULT
CTX_DEFAULT -->|"without custom Modelfile"| MODELS
CSTREAM --> MODELS
CHAT --> MODELS
GEN --> MODELS
style OC_COMPLETE fill:#ffe7aa
style OL_COMPLETE fill:#ffe7aa
style CHAT fill:#b5f5b5
style GEN fill:#b5f5b5
style JMAX fill:#b5f5b5
style MAX fill:#ffe7aa
style CMAX fill:#ffaaaa
style J_MODEL fill:#b5f5b5
style OL_MODEL fill:#b5f5b5
style MODEL fill:#b5f5b5
Loading
Why create another custom Modelfile for every model?
The only practical workaround without a bridge is creating a custom Ollama Modelfile for every model just to override runtime parameters.
This is redundant, time-consuming, and difficult to maintain when frequently testing or switching models.
flowchart TD
M1["llama3.2<br/>Model"]
F11["Base Model<br/>llama3.2"]
F12["Modelfile<br/>llama3.2<br/>num_ctx=32K"]
M2["qwen3<br/>Model"]
F21["Base Model<br/>qwen3"]
F22["Modelfile<br/>qwen3<br/>num_ctx=32K"]
M3["gemma3<br/>Model"]
F31["Base Model<br/>gemma3"]
F32["Modelfile<br/>gemma3<br/>num_ctx=32K"]
M4["mistral<br/>Model"]
F41["Base Model<br/>mistral"]
F42["Modelfile<br/>mistral<br/>num_ctx=32K"]
M5["deepseek-r1<br/>Model"]
F51["Base Model<br/>deepseek-r1"]
F52["Modelfile<br/>deepseek-r1<br/>num_ctx=32K"]
M1 --> F11
M1 --> F12
M2 --> F21
M2 --> F22
M3 --> F31
M3 --> F32
M4 --> F41
M4 --> F42
M5 --> F51
M5 --> F52
style M1 fill:#b5f5b5
style M2 fill:#b5f5b5
style M3 fill:#b5f5b5
style M4 fill:#b5f5b5
style M5 fill:#b5f5b5
style F12 fill:#ffe7aa
style F22 fill:#ffe7aa
style F32 fill:#ffe7aa
style F42 fill:#ffe7aa
style F52 fill:#ffe7aa
Loading
What this proxy does
Uses Ollama's native /api/chat endpoint directly, where runtime parameters are explicitly supported.
One .env file configures all models centrally.
Central API bridge to enforce runtime settings
Provide a single OpenAI endpoint while centrally enforcing runtime policies and automatically switching to a secondary Ollama server if the primary becomes unavailable.
flowchart LR
Client["OpenCode<br/>OpenAI Connector"]
subgraph Bridge["API Bridge"]
TRANS["OpenAI → Ollama<br/>Translation"]
CONTEXT["Enforce<br/>Context Size Policy"]
FAILOVER["Switch to other Server"]
end
O1["Ollama Server"]
Client -->|"speaks OpenAI API"| Bridge
Bridge -->|"speaks Ollama API"| O1
style Client fill:#ffaaaa
style Bridge fill:#88FFFF
style O1 fill:#b5f5b5
Loading
High Availability / Failover
Provide a single OpenAI endpoint while automatically switching to a secondary Ollama server if the primary becomes unavailable.
flowchart LR
Client["OpenCode<br/>OpenAI Connector"]
Bridge["OpenAI ↔ Ollama API Bridge"]
Decision{"Primary Available?"}
C1["Context Size = 128K"]
C2["Context Size = 32K"]
O1["Primary Ollama Server"]
O2["Secondary Ollama Server"]
Client --> Bridge
Bridge --> Decision
Decision -->|Yes| C1
Decision -->|No| C2
C1 --> O1
C2 --> O2
style Bridge fill:#ffff88
style O1 fill:#b5f5b5
style O2 fill:#b5f5b5
style C1 fill:#e8f4fd
style C2 fill:#88FFFF
Loading
Quick Start
git clone < repo >
cd py-ollama-openai-bridge
cp .env.example .env
# edit .env: set OLLAMA_URL to your Ollama host
docker compose up -d
Bridge available at:
http://localhost:8080/v1
Direct (without Docker)
pip install -r requirements.txt
# configure .env
python proxy.py
Point OpenCode at the bridge:
{
"provider" : {
"ollama-bridge" : {
"name" : " Ollama (bridged) " ,
"npm" : " @ai-sdk/openai-compatible " ,
"options" : {
"baseURL" : " http://localhost:8080/v1 "
},
"models" : {
"gpt-oss:20b" : {
"name" : " _gpt-oss:20b " ,
"tool_call" : true ,
"limit" : {
"context" : 64000 ,
"output" : 40960
}
}
}
}
}
}
All options
NUM_CTX = 64000
NUM_PREDICT = 128000
TEMPERATURE = 0.7
TOP_P = 0.9
TOP_K = 40
MIN_P = 0.05
REPEAT_PENALTY = 1.1
SEED = 42
KEEP_ALIVE = 30m
Only uncommented values in .env are injected.
Unset values use Ollama defaults.
OLLAMA_URL → FAILOVER_OLLAMA_URL
Example:
OLLAMA_URL = http://192.168.0.42:11434
FAILOVER_OLLAMA_URL = http://192.168.0.101:11434
FAILOVER_NUM_CTX = 96000
FAILOVER_NUM_PREDICT = 128000
FAILOVER_KEEP_ALIVE = 5m
Features:
Health check via GET /api/tags
Automatic switch when primary is unavailable
Failover on connection errors and timeouts
Independent failover parameters
Simple mode: only OLLAMA_URL required
Feature
OpenCode w/ Ollama
API Bridge
OpenAI Chat Completions
✅
✅
OpenAI Responses API
⚠️ Limited
✅
Native /api/chat
❌
✅
Native /api/generate
❌
✅
Streaming
✅
✅
Tool Calling
Backend dependent
Preserved
Multiple Ollama Servers
❌ Manual switch
✅
Context Size Policy
❌ Per-model Modelfiles
✅ Centralized runtime policy
Automatic Failover
❌
✅
Routing Policies
❌
✅
Single Stable Endpoint
❌
✅
Runtime configuration model
Instead of creating model-specific Modelfiles:
llama3.2 + Modelfile
qwen3 + Modelfile
gemma3 + Modelfile
mistral + Modelfile
deepseek + Modelfile
the bridge applies runtime policies centrally:
.env
|
+-- NUM_CTX
+-- NUM_PREDICT
+-- TEMPERATURE
+-- TOP_P
+-- TOP_K
+-- KEEP_ALIVE
|
v
All Ollama models
This allows changing runtime behaviour without rebuilding model definitions.
Keep OpenAI-compatible clients unchanged
Use Ollama's native API capabilities
Centralize runtime configuration
Support multiple Ollama backends
Preserve tool calls and reasoning metadata where available
prevent 2s failover delay by ongoing checking primary server
just cache health check for 60 seconds
My actual usecase is accessing my GPU Cluster from outside. That's why I also implemented the failover function.
Having an API Gateway AND reverse Proxy mit seem overpowered for this little setup, but I use Authentik with TraefiK/Caddy to access Homelab WebApps and Kong to access Homelab APIs.
You want only one Tier to handle TLS Connections, so Caddy(Homelab) or Traefik(Rootserver) are always involved.
flowchart TD
subgraph Laptop
OC["OpenCode Client"]
end
subgraph Homelan["Home LAN"]
subgraph Router
IP["Public IPV4"]
Portforwarding["NAT Forward Port 443"]
end
subgraph OptiPlex["Optiplex"]
subgraph PVE["Proxmox Cluster"]
subgraph LXC["LXC HomeLab Container"]
subgraph Docker["Docker Compose"]
ReverseProxy["Traefik/Caddy<br />Reverse Proxy<br/>TLS Termination"]
Kong["Kong API Gateway<br />JWT Authorization"]
Bridge["API Bridge<br />this project"]
end
end
end
end
subgraph OldLaptop["Linux Laptop"]
subgraph Debian["Debian Linux"]
ROCM["Rocm/Vulcan driver"]
subgraph Debian_Docker["Docker Compose"]
ollama_docker["Ollama Server"]
end
end
end
RX["AMD RX 6900XT<br/>16GB VRAM"]
subgraph Gaming["Gaming Machine"]
subgraph Win11["Windows 11"]
CUDA["CUDA driver"]
ollama_exe["ollama.exe"]
end
RTX["Nvidia RTX 5090<br/>32GB VRAM"]
end
end
OC -->|Internet| IP
IP --> Portforwarding
Portforwarding --> ReverseProxy
ReverseProxy --> Kong
Kong --> Bridge
Bridge -->|"PRIMARY<br />num_ctx=128K"| ollama_exe
RTX -.- CUDA
CUDA -.- ollama_exe
RX -.-|"Riser Cable"| ROCM
ROCM -.- ollama_docker
Bridge -.->|"FAILOVER<br />num_ctx=32K"| ollama_docker
style OptiPlex fill: #88FFFF
style OldLaptop fill: #88FFFF
style Gaming fill: #88FFFF
style Bridge fill:#b5f5b5
Loading
License
A proxy to translate Opencode openai API calls to native ollama, allow context size to forward and even overwrite to be >4096
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
