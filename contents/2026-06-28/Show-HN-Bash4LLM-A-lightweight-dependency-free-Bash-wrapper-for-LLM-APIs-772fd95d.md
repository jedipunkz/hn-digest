---
source: "https://github.com/kamaludu/bash4llm/"
hn_url: "https://news.ycombinator.com/item?id=48710827"
title: "Show HN: Bash4LLM+ – A lightweight, dependency-free Bash wrapper for LLM APIs"
article_title: "GitHub - kamaludu/bash4llm: Bash-first wrapper for Groq’s OpenAI-compatible API. Secure, portable, Termux-friendly. · GitHub"
author: "kamaludu"
captured_at: "2026-06-28T20:25:18Z"
capture_tool: "hn-digest"
hn_id: 48710827
score: 6
comments: 3
posted_at: "2026-06-28T19:43:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Bash4LLM+ – A lightweight, dependency-free Bash wrapper for LLM APIs

- HN: [48710827](https://news.ycombinator.com/item?id=48710827)
- Source: [github.com](https://github.com/kamaludu/bash4llm/)
- Score: 6
- Comments: 3
- Posted: 2026-06-28T19:43:10Z

## Translation

タイトル: Show HN: Bash4LLM+ – LLM API 用の軽量で依存関係のない Bash ラッパー
記事のタイトル: GitHub - kamaludu/bash4llm: Groq の OpenAI 互換 API 用の Bash ファースト ラッパー。安全でポータブル、Termux に優しい。 · GitHub
説明: Groq の OpenAI 互換 API 用の Bash ファースト ラッパー。安全でポータブル、Termux に優しい。 - GitHub - kamaludu/bash4llm: Groq の OpenAI 互換 API の Bash ファースト ラッパー。安全でポータブル、Termux に優しい。
HN テキスト: Bash4LLM は、端末から LLM と対話するための単一ファイルの Bash ラッパーです。 Python、Node、その他のランタイムをインストールせずに動作するシンプルなものが欲しかったので、これを作成しました。 Bash、curl、jq のみを使用します。プロンプトの送信、小規模なチャットの開始、ファイルの 1 行ずつの処理、ストリーム出力、セッション メタデータの JSON 形式での保存が可能です。私はそれを安全かつ予測可能にしようとしました。システム /tmp も eval も使用しません。 Groq はデフォルトでサポートされており、extras/providers/ フォルダー内の専用 Bash スクリプトを使用して他のプロバイダーを追加できます。例: echo "コマンドの説明: ls -l" | ./bash4llm

記事本文:
GitHub - kamaludu/bash4llm: Groq の OpenAI 互換 API 用の Bash ファースト ラッパー。安全でポータブル、Termux に優しい。 · GitHub
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
カマルドゥ
/
bash4llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
138 コミット 138 コミット .github .github bin bin docs docs extras extras testing testing .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md INSTALL-en.md INSTALL-en.md INSTALL.md INSTALL.md ライセンス ライセンスPROVIDERS.md PROVIDERS.md README-ja.md README-ja.md README.md README.md RELEASE-NOTES.md RELEASE-NOTES.md SECURITY-ja.md SECURITY-ja.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Bash4LLM⁺ — Groq の OpenAI 互換 Chat Completions API 用の安全で Bash ファーストの完全に監査可能な CLI ラッパー (他のプロバイダーにも拡張可能)。
Bash4LLM⁺ は、自己完結型で読みやすく、テスト可能な単一の Bash スクリプトです。
ダウンロードして実行可能にし、API キーをエクスポートして、すぐに使用を開始できます。
Unix のような環境と互換性があります: Linux、macOS、WSL、Cygwin、Termux (Android)、BSD。
動的モデルのリスト
GET 経由 https://api.groq.com/openai/v1/models
→ ハードコーディングされたモデルはありません。
設計によるセキュリティ
→ /tmp 、 eval の使用なし、制限的なアクセス許可、高度なプロバイダー検証。
セクションごとのモジュール構造
→ PRECORE_BOOT、PRECORE_RUN、PROVIDER、CORE_SETUP、CORE_PROVIDER。
UI 状態システム (ui_state)
→ CORE は、GUI または外部ツール (ホーム アシスタントなど) との統合のためにメタデータをアトミック JSON 形式で常に公開します。
ストリーミングと非ストリーミング
→ 応答の最後にリアルタイムまたは完全な出力。
自動保存
→ 設定可能なしきい値を超える長い出力の場合。
高度なモデル管理
→ 更新、リスト、永続的なデフォルト、動的ホワイトリスト、自動選択。
オプションの追加物
→ 追加プロバイダー (Gemini、Hugging Face、Mistral など)、テンプレート、ドキュメント、ツール

安全なもの。
Termux / Android に対応
→ flock (Android ではカーネル/SELinux レベルで不安定または制限されることが多い) をバイパスすることで Termux 環境を自動的に検出し、同時実行管理を透過的に堅牢なディレクトリ ロック メカニズム (アトミック mkdir) に切り替えます。
脅威モデル (短縮版)
Bash4LLM⁺ は、シングルユーザー環境 (PC/ラップトップ、個人サーバー) 向けに設計されています。
プロバイダーはシェル内で実行されるコードです。プロバイダーは、所有する安全なディレクトリに存在する必要があります。
BASH4LLM_EXTRAS_DIR や BASH4LLM_TMPDIR などの変数は、信頼できる構成とみなされます。
スクリプトはモデルを出力しません。
TOCTOU のリスクと JSON/SSE 解析の制限が軽減され、文書化されています。
詳細についてはセキュリティをご覧ください。
Bash4LLM⁺ では、次のパッケージ (または同等のパッケージ) が PATH で使用可能である必要があります。
⏩ 早送り (クイックインストール)
ターミナルで次のコマンドを実行して、Bash4LLM⁺ をすぐに起動します。
#1.リポジトリのクローンを作成します (最大速度を実現するには最新のコミットのみ)
git clone -- Depth 1 --branch main https://github.com/kamaludu/bash4llm.git repo-bash4llm
#2.作業フォルダーを作成し、実行可能ファイルを抽出します
mkdir -p bash4llm
cp リポ-bash4llm/bin/bash4llm bash4llm/
chmod +x bash4llm/bash4llm
#3.フォルダーに移動してテンプレートを更新します
cd bash4llm
./bash4llm --refresh-models
スクリプトは、デフォルトのプロバイダー (Groq) の API キーを入力するように求めます。
プロバイダー groq の API キーを入力します (env GROQ_API_KEY):
API キーを入力してエクスポートすると、セッション中に入力する必要がなくなります。
エクスポート GROQ_API_KEY="gsk_xxxxxxxxxxxxxxxxxx"
推奨: オプションの追加機能をインストールします。
#4.エクストラのインストール
./bash4llm --install-extras ../repo-bash4llm/extras/
Bash4llm を使用します ⚡
詳細な手順

あなたは: インストール
chmod +x bash4llm
import GROQ_API_KEY= " gsk_xxxxxxxxxxxxxxxxxx "
./bash4llm --ヘルプ
オプションの追加物:
./bash4llm --install-extras
オプションあり:
選択的インストール:
./bash4llm --install-extras Provider1 templateA
./bash4llm 「イタリア語で短い詩を書いてください」
複数行のプロンプト:
./bash4llm << 'EOF'
短い詩を書く
イタリア語で
終了後
ファイルからの入力:
./bash4llm -f プロンプト.txt
パイプ:
echo "相対性理論を説明してください" | ./bash4llm
特定のモデル:
./bash4llm -m llama-3.3-70b-versatile "短いエッセイを書く"
ドライラン:
./bash4llm --dry-run "こんにちは"
外部プロバイダー (インストールされている場合):
./bash4llm --provider gemini "これを翻訳します"
利用可能なコマンド、フラグ、オプション
フラグ
トピック
効果
--refresh-models 、 --refresh-model
いいえ
モデルリストを更新します（APIキーが必要です）。
--リストモデル
いいえ
モデルリストを印刷します（対話形式）。
--list-models-raw
いいえ
モデルリストを未加工形式で印刷します (モデルごとに 1 行)。
--リストプロバイダー
いいえ
プロバイダーリストを印刷します。
--list-providers-raw
いいえ
RAW 形式の印刷プロバイダー。
--set-default <モデル>
はい
アクティブなプロバイダーの永続的なデフォルト テンプレートを設定します。
-m <モデル> 、 --model <モデル>
はい
この実行のテンプレートを設定します。
--プロバイダー<名前>
はい
CLIからプロバイダーを設定します。
--プロバイダー
いいえ
引数がない場合 → インタラクティブな選択が開きます。
入力（ファイル、JSON、テンプレート、バッチ）
フラグ
トピック
効果
-f <ファイル>
はい
ファイルを FILE_INPUTS に追加します。
--json-input <json>
はい
JSON 入力 (OpenAI のような形式) を設定します。
--template <名前>
はい
BASH4LLM_TEMPLATES_DIR からテンプレートを適用します。
--batch <ファイル>
はい
バッチ リクエストを実行します (1 行 = 1 プロンプト)。
セッション
フラグ
トピック
効果
--セッション<id>
はい
特定の ID でセッションを有効にします。
--セッションウィンドウ [n]
オプション
セッション ウィンドウを設定します (指定されていない場合、デフォルトは 10)。
--init-ses

シオン
はい
API 呼び出しを行わずに、空のセッションを安全に初期化し (NDJSON ファイルとメタデータを作成)、グローバル セッション インデックスに記録します。 --session <id> を併用する必要があります。
モデル/生成パラメータ
フラグ
トピック
効果
--system <テキスト>
はい
システムプロンプトを設定します。
-- チュア <n>
はい
温度パラメータを設定します (0.0 ～ 2.0、エイリアス canonical)。
--温度<n>
はい
--ture の別名。
--最大 <n>
はい
最大トークンを設定します。
出力と保存
フラグ
トピック
効果
--保存
いいえ
出力を強制的に保存します。
--nosave
いいえ
保存を無効にします。
--out <パス>
はい
出力ファイル/ディレクトリのパス。
-- しきい値 <n>
はい
自動保存のバイト単位のサイズしきい値 (デフォルト: 1000)。
--json
いいえ
健全な生の JSON 出力。
--かなり
いいえ
フォーマットされた JSON 出力。
--テキスト
いいえ
抽出された標準テキスト出力 (デフォルトの動作)。
--生
いいえ
末尾の区切りを除いた生のテキスト出力。
動作モード
フラグ
トピック
効果
--ドライラン
いいえ
実際の API 呼び出しはありません (シミュレートされた動作)。
--静か
いいえ
不要な出力を削減し、TTY のタイトルを抑制します。
--ストリーム
いいえ
アクティブな非同期ストリーミング。
--ストリームなし
いいえ
非同期ストリーミングを無効にします。
--チャット
いいえ
REPL インタラクティブ チャット モード。
--ブートストラップのみ
いいえ
パス/ロックの検証のみを実行して終了します。
構成と診断
フラグ
トピック
効果
--show-config
いいえ
アクティブな完全な構成を表示します。
--診断
いいえ
完全なシステム診断を実行します。
--バージョン
いいえ
スクリプトのバージョンを出力して終了します。
-h 、 --help
いいえ
ファイルからフォーマットされた対話型ヘルプを表示します。
追加インストール
フラグ
トピック
効果
--install-extras
オプション
エクストラをインストールします。ソースディレクトリを受け入れることができます。
--install-extras=<ディレクトリ>
はい
特定のソース ディレクトリからエクストラをインストールします。
解析の終了

フラグ
効果
--
解析オプションを終了します。
-*
不明なオプション→エラー。
*
位置引数 → ARGS に追加。
構成とモデル
$BASH4LLM_CONFIG_DIR/config
→ ローカルパラメータ (MODEL、TURE、MAX_TOKENS、FORMAT、THRESHOLD)
$BASH4LLM_CONFIG_DIR/model.$PROVIDER
→ プロバイダーのデフォルトのテンプレート
$MODELS_FILE
→ --refresh-models によってモデルのホワイトリストが更新されました
自動プロバイダーの選択 ( auto_select_model_<provider> )
最初のホワイトリスト エントリ (models.txt)
グローバル構成レガシー構成 (MODEL=...)
共有 OS レベルでは /tmp は使用できません。
一時ファイルは、権限 700 ( umask 077 ) で $RUN_TMPDIR ディレクトリに分離されます。
ファイルは 600 のアクセス許可で保存されました。
--out を指定すると、可能であれば Bash4LLM⁺ がディレクトリを作成します。
📁 UI 状態システム (ui_state)
Bash4LLM⁺ は、以下のアトミック JSON ファイルを介して、外部 GUI/ツール向けの操作メタデータを公開します。
$BASH4LLM_CONFIG_DIR/ui_state
含まれるもの:
session/<id>.json → セッション状態 (active、msg_count、last_ts)
session/index.json → セッションリスト
last_api.json → 最後の API 結果 (http_status、req_id、edgecase_detected など)
last_history.json → 最後に保存された履歴
Provider_capabilities.json → アクティブなプロバイダー機能 (ストリーミング、refresh_models)
GUI (オプション) は、CGI プレースホルダーとしてこれらのファイルのみを読み取ります。
📘 Bash4LLM のコンテキストメモリ⁺
Bash4LLM⁺ はそれ自体ではメモリを維持しません。
メモリは、 --session を使用してセッションをアクティブ化した場合にのみ存在します。
各セッションでは、永続的な NDJSON ファイルが作成されます。
$BASH4LLM_HISTORY_DIR/sessions/<session_id>.ndjson
また、Bash4LLM⁺ はセッション メタデータを次の場所に保持します。
$BASH4LLM_CONFIG_DIR/ui_state/sessions/<session_id>.json
このメタデータは、外部 GUI/ツールの正規ソースです。
./bash4llm --session chat1 "こんにちは"
./ba

sh4llm --session chat1 " 私が言ったことを要約してください "
🟩 --session-window の正しい使用法
./bash4llm --session chat1 --session-window 10 "続行"
🟧 基本的なルール
コンテキスト メモリを確保するには、 --session <id> を常に含める必要があります。
モデル出力は実行されません。
Provider = code: エクストラ/プロバイダーを安全に保ちます。
環境変数 = 信頼できる構成。
コード
変数
意味
0
-
成功
10
BASH4LLM_ERR_NO_API_KEY
API キーがありません
11
BASH4LLM_ERR_BAD_MODEL
無効なモデルまたはホワイトリストに登録されていないモデル
12
BASH4LLM_ERR_CURL_FAILED
ネットワーク/カールエラー
14
BASH4LLM_ERR_NO_PROMPT
プロンプトは提供されません
15
BASH4LLM_ERR_TMP
一般的なファイルシステム/一時エラー
16
BASH4LLM_ERR_API
ベンダーHTTP/APIエラー
主な変数
変数
必要
説明
GROQ_API_KEY
API 呼び出しの場合ははい
Groq API キープロバイダー。
BASH4LLM_CONFIG_DIR
推奨
構成ディレクトリ。
BASH4LLM_MODELS_DIR
推奨
テンプレートディレクトリ。
BASH4LLM_TMPDIR
はい
一時ディレクトリ。
BASH4LLM_HISTORY_DIR
推奨
セッションディレクトリと履歴。
モデル
いいえ
アクティブモデル。
プロバイダー
いいえ
アクティブなプロバイダー。
ALLOWED_MODELS
いいえ
許可されたモデルのホワイトリスト。
ライセンス
Bash4LLM⁺ は GPL v3 ライセンスに基づいて配布されます。
「ライセンス」を参照してください。
著者: クリスティアン・エヴァンジェリスティ
電子メール: opensource @ cevangel. anonaddy. me
リポジトリ: https://github.com/k

[切り捨てられた]

## Original Extract

Bash-first wrapper for Groq’s OpenAI-compatible API. Secure, portable, Termux-friendly. - GitHub - kamaludu/bash4llm: Bash-first wrapper for Groq’s OpenAI-compatible API. Secure, portable, Termux-friendly.

Bash4LLM is a single-file Bash wrapper for interacting with LLMs from the terminal. I created it because I wanted something simple that worked without installing Python, Node, or any other runtime. It uses only Bash, curl, and jq. You can send prompts, start a small chat, process files line by line, stream output, and save session metadata in JSON format. I tried to make it safe and predictable: no use of the system /tmp, no use of eval. Groq is supported by default, and other providers can be added with dedicated Bash scripts in the extras/providers/ folder. Example: echo "explains the command: ls -l" | ./bash4llm

GitHub - kamaludu/bash4llm: Bash-first wrapper for Groq’s OpenAI-compatible API. Secure, portable, Termux-friendly. · GitHub
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
kamaludu
/
bash4llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
138 Commits 138 Commits .github .github bin bin docs docs extras extras tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md INSTALL-en.md INSTALL-en.md INSTALL.md INSTALL.md LICENSE LICENSE PROVIDERS.md PROVIDERS.md README-en.md README-en.md README.md README.md RELEASE-NOTES.md RELEASE-NOTES.md SECURITY-en.md SECURITY-en.md SECURITY.md SECURITY.md View all files Repository files navigation
Bash4LLM⁺ — wrapper CLI sicuro, Bash‑first e completamente auditabile per l’API Chat Completions compatibile OpenAI di Groq (ed estendibile ad altri provider).
Bash4LLM⁺ è un singolo script Bash, auto‑contenuto, leggibile e verificabile.
Scaricalo, rendilo eseguibile, esporta la tua API key e inizia subito a usarlo.
Compatibile con ambienti Unix‑like: Linux, macOS, WSL, Cygwin, Termux (Android), BSD.
Lista modelli dinamica
tramite GET https://api.groq.com/openai/v1/models
→ nessun modello hardcoded.
Sicurezza by design
→ nessun uso di /tmp , nessun eval , permessi restrittivi, validazione provider avanzata.
Struttura modulare a sezioni
→ PRECORE_BOOT, PRECORE_RUN, PROVIDER, CORE_SETUP, CORE_PROVIDER.
Sistema di Stato UI (ui_state)
→ il CORE espone costantemente metadati in formato JSON atomico per l'integrazione con GUI o strumenti esterni (es. Home Assistant).
Streaming e non‑streaming
→ output in tempo reale o completo a fine risposta.
Salvataggio automatico
→ per output lunghi oltre una soglia configurabile.
Gestione modelli avanzata
→ refresh, lista, default persistente, whitelist dinamica, auto‑selezione.
Extras opzionali
→ provider aggiuntivi (come Gemini, Hugging Face, Mistral), template, documentazione, strumenti di sicurezza.
Pronto per Termux / Android
→ rileva automaticamente l'ambiente Termux bypassando flock (spesso instabile o limitato a livello kernel/SELinux su Android) e devia trasparentemente la gestione della concorrenza sul robusto meccanismo di directory lock ( mkdir atomico).
Modello di minaccia (versione breve)
Bash4LLM⁺ è progettato per ambienti single‑user (PC/laptop, server personali).
I provider sono codice eseguito nella tua shell: devono risiedere in directory sicure di tua proprietà.
Variabili come BASH4LLM_EXTRAS_DIR e BASH4LLM_TMPDIR sono considerate configurazione fidata.
Lo script non esegue mai l’output del modello.
I rischi TOCTOU e i limiti del parsing JSON/SSE sono mitigati e documentati.
Dettagli completi in SECURITY .
Bash4LLM⁺ richiede che i seguenti pacchetti (o equivalenti) siano disponibili nel PATH:
⏩ FAST FORWARD (Installazione Rapida)
Esegui questi comandi nel tuo terminale per avviare subito Bash4LLM⁺ :
# 1. Clona il repository (solo l'ultimo commit per massima velocità)
git clone --depth 1 --branch main https://github.com/kamaludu/bash4llm.git repo-bash4llm
# 2. Crea una cartella di lavoro ed estrai l'eseguibile
mkdir -p bash4llm
cp repo-bash4llm/bin/bash4llm bash4llm/
chmod +x bash4llm/bash4llm
# 3. Entra nella cartella e aggiorna i modelli
cd bash4llm
./bash4llm --refresh-models
Lo script ti chiederà l'inserimento della tua chiave API per il provider di default (Groq):
Enter API key for provider groq (env GROQ_API_KEY):
Inserisci la tua API key, poi esportala per non doverla più inserire durante la sessione:
export GROQ_API_KEY="gsk_xxxxxxxxxxxxxxxxx"
Consigliato: installa gli Extras opzionali :
# 4. Installazione degli Extras
./bash4llm --install-extras ../repo-bash4llm/extras/
Usa Bash4llm ⚡
Istruzioni dettagliate in: INSTALL
chmod +x bash4llm
export GROQ_API_KEY= " gsk_xxxxxxxxxxxxxxxxx "
./bash4llm --help
Extras opzionali:
./bash4llm --install-extras
Con opzioni:
installazione selettiva:
./bash4llm --install-extras provider1 templateA
./bash4llm " scrivi una breve poesia in italiano "
Prompt multilinea:
./bash4llm << ' EOF '
scrivi una breve poesia
in italiano
EOF
Input da file:
./bash4llm -f prompt.txt
Pipe:
echo " spiegami la relatività " | ./bash4llm
Modello specifico:
./bash4llm -m llama-3.3-70b-versatile " scrivi un saggio breve "
Dry run:
./bash4llm --dry-run " ciao "
Provider esterno (se installato):
./bash4llm --provider gemini " traduci questo "
Comandi, flag e opzioni disponibili
Flag
Argomento
Effetto
--refresh-models , --refresh-model
no
Aggiorna la lista modelli (richiede API key).
--list-models
no
Stampa lista modelli (formato interattivo).
--list-models-raw
no
Stampa lista modelli in formato raw (una riga per modello).
--list-providers
no
Stampa lista provider.
--list-providers-raw
no
Stampa provider in formato raw.
--set-default <model>
sì
Imposta modello di default persistente per il provider attivo.
-m <model> , --model <model>
sì
Imposta modello per questa esecuzione.
--provider <name>
sì
Imposta provider da CLI.
--provider
no
Se senza argomento → apre selezione interattiva.
Input (file, JSON, template, batch)
Flag
Argomento
Effetto
-f <file>
sì
Aggiunge file a FILE_INPUTS .
--json-input <json>
sì
Imposta input JSON (formato OpenAI-like).
--template <name>
sì
Applica template da BASH4LLM_TEMPLATES_DIR .
--batch <file>
sì
Esegue richieste batch (una riga = un prompt).
Sessioni
Flag
Argomento
Effetto
--session <id>
sì
Abilita sessione con ID specifico.
--session-window [n]
opzionale
Imposta finestra sessione (default 10 se non fornito).
--init-session
si
Inizializza in sicurezza una sessione vuota (creando i file NDJSON e i metadati) e la registra nell'indice globale delle sessioni, senza effettuare chiamate API. Richiede l'uso congiunto di --session <id> .
Parametri modello / generazione
Flag
Argomento
Effetto
--system <text>
sì
Imposta system prompt.
--ture <n>
sì
Imposta parametro temperatura (da 0.0 a 2.0, alias canonico).
--temperature <n>
sì
Alias di --ture .
--max <n>
sì
Imposta max token.
Output e salvataggio
Flag
Argomento
Effetto
--save
no
Forza salvataggio output.
--nosave
no
Disabilita salvataggio.
--out <path>
sì
Percorso file/directory output.
--threshold <n>
sì
Soglia dimensione in byte per salvataggio automatico (default: 1000).
--json
no
Output JSON raw integro.
--pretty
no
Output JSON formattato.
--text
no
Output testuale standard estratto (comportamento predefinito).
--raw
no
Output testuale grezzo escludendo separazioni finali.
Modalità operative
Flag
Argomento
Effetto
--dry-run
no
Nessuna chiamata API reale (comportamento simulato).
--quiet
no
Riduce l'output non necessario e sopprime i titoli su TTY.
--stream
no
Streaming asincrono attivo.
--no-stream
no
Disattiva streaming asincrono.
--chat
no
Modalità chat interattiva REPL.
--bootstrap-only
no
Esegue solo validazione percorsi/lock e termina.
Configurazione e diagnostica
Flag
Argomento
Effetto
--show-config
no
Mostra configurazione completa attiva.
--diagnostics
no
Esegue diagnostica completa del sistema.
--version
no
Stampa versione dello script e termina.
-h , --help
no
Mostra help interattivo formattato da file.
Installazione extras
Flag
Argomento
Effetto
--install-extras
opzionale
Installa extras; può accettare directory sorgente.
--install-extras=<dir>
sì
Installa extras da directory sorgente specifica.
Terminazione parsing
Flag
Effetto
--
Termina parsing opzioni.
-*
Opzione sconosciuta → errore.
*
Argomento posizionale → aggiunto a ARGS .
Configurazione e modelli
$BASH4LLM_CONFIG_DIR/config
→ parametri locali (MODEL, TURE, MAX_TOKENS, FORMAT, THRESHOLD)
$BASH4LLM_CONFIG_DIR/model.$PROVIDER
→ modello predefinito per provider
$MODELS_FILE
→ whitelist modelli aggiornata da --refresh-models
auto‑selezione provider ( auto_select_model_<provider> )
prima voce della whitelist ( models.txt )
configurazione globale legacy config ( MODEL=... )
Nessun uso di /tmp a livello di sistema operativo condiviso.
File temporanei isolati in directory $RUN_TMPDIR con permessi 700 ( umask 077 ).
File salvati con permessi 600 .
Con --out Bash4LLM⁺ crea la directory se possibile.
📁 Sistema di Stato UI (ui_state)
Bash4LLM⁺ espone metadati operativi destinati a GUI/strumenti esterni tramite file JSON atomici in:
$BASH4LLM_CONFIG_DIR/ui_state
Contiene:
sessions/<id>.json → stato sessione (active, msg_count, last_ts)
sessions/index.json → elenco sessioni
last_api.json → ultimo risultato API (http_status, req_id, edgecase_detected, ecc.)
last_history.json → ultimo salvataggio history
provider_capabilities.json → capacità provider attivo (streaming, refresh_models)
La GUI (extra opzionale) legge solo questi file per i placeholder CGI.
📘 Memoria contestuale in Bash4LLM⁺
Bash4LLM⁺ non mantiene memoria da solo .
La memoria esiste solo se attivi una sessione tramite --session .
Ogni sessione crea un file NDJSON persistente:
$BASH4LLM_HISTORY_DIR/sessions/<session_id>.ndjson
E Bash4LLM⁺ mantiene i metadati della sessione in:
$BASH4LLM_CONFIG_DIR/ui_state/sessions/<session_id>.json
Questi metadati sono la fonte canonica per GUI/strumenti esterni.
./bash4llm --session chat1 " Ciao "
./bash4llm --session chat1 " Riassumi ciò che ho detto "
🟩 Uso corretto di --session-window
./bash4llm --session chat1 --session-window 10 " continua "
🟧 Regola fondamentale
Per avere memoria contestuale devi sempre includere --session <id> .
Nessuna esecuzione dell’output del modello.
Provider = codice: mantieni extras/providers sicuro.
Variabili d’ambiente = configurazione fidata.
Codice
Variabile
Significato
0
-
Successo
10
BASH4LLM_ERR_NO_API_KEY
API key mancante
11
BASH4LLM_ERR_BAD_MODEL
Modello non valido o non in whitelist
12
BASH4LLM_ERR_CURL_FAILED
Errore rete/curl
14
BASH4LLM_ERR_NO_PROMPT
Nessun prompt fornito
15
BASH4LLM_ERR_TMP
Errore generico filesystem / temporanei
16
BASH4LLM_ERR_API
Errore HTTP/API del fornitore
Variabili principali
Variabile
Necessaria
Descrizione
GROQ_API_KEY
sì per chiamate API
API key provider Groq.
BASH4LLM_CONFIG_DIR
consigliata
Directory configurazione.
BASH4LLM_MODELS_DIR
consigliata
Directory modelli.
BASH4LLM_TMPDIR
sì
Directory temporanea.
BASH4LLM_HISTORY_DIR
consigliata
Directory sessioni e cronologia.
MODEL
no
Modello attivo.
PROVIDER
no
Provider attivo.
ALLOWED_MODELS
no
Whitelist modelli ammessi.
Licenza
Bash4LLM⁺ è distribuito sotto licenza GPL v3.
Vedi LICENSE .
Autore: Cristian Evangelisti
Email: opensource​@​cevangel.​anonaddy.​me
Repository: https://github.com/k

[truncated]
