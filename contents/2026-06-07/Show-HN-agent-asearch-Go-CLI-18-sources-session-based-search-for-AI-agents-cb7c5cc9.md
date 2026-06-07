---
source: "https://github.com/izzzzzi/agent-asearch"
hn_url: "https://news.ycombinator.com/item?id=48432000"
title: "Show HN: agent-asearch – Go CLI, 18 sources, session-based search for AI agents"
article_title: "GitHub - izzzzzi/agent-asearch: Multi-source search CLI for LLM agents — web, HN, Reddit, GitHub, Tavily, Exa, Brave, Jina, YouTube, Twitter. Go single binary. Session-based workflow with JSON contract. · GitHub"
author: "izzzzzi"
captured_at: "2026-06-07T07:31:22Z"
capture_tool: "hn-digest"
hn_id: 48432000
score: 1
comments: 0
posted_at: "2026-06-07T05:08:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: agent-asearch – Go CLI, 18 sources, session-based search for AI agents

- HN: [48432000](https://news.ycombinator.com/item?id=48432000)
- Source: [github.com](https://github.com/izzzzzi/agent-asearch)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T05:08:03Z

## Translation

タイトル: Show HN: Agent-asearch – Go CLI、18 ソース、AI エージェントのセッションベースの検索
記事のタイトル: GitHub - izzzzzi/agent-asearch: LLM エージェント用のマルチソース検索 CLI — Web、HN、Reddit、GitHub、Tavily、Exa、Brave、Jina、YouTube、Twitter。シングルバイナリに移行します。 JSON コントラクトを使用したセッションベースのワークフロー。 · GitHub
説明: LLM エージェント用のマルチソース検索 CLI — Web、HN、Reddit、GitHub、Tavily、Exa、Brave、Jina、YouTube、Twitter。シングルバイナリに移行します。 JSON コントラクトを使用したセッションベースのワークフロー。 - izzzzzi/エージェント-asearch

記事本文:
GitHub - izzzzzi/agent-asearch: LLM エージェント用のマルチソース検索 CLI — Web、HN、Reddit、GitHub、Tavily、Exa、Brave、Jina、YouTube、Twitter。シングルバイナリに移行します。 JSON コントラクトを使用したセッションベースのワークフロー。 · GitHub
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
イッズジー
/
エージェント検索
プ

ブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .githooks .githooks .github/ workflows .github/ workflows bin bin cmd/ asearch cmd/ asearch docs docs 内部 内部スクリプト スクリプト .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .pre-commit-config.yaml .pre-commit-config.yaml AGENT_INSTRUCTIONS.en.md AGENT_INSTRUCTIONS.en.md AGENT_INSTRUCTIONS.md AGENT_INSTRUCTIONS.md ライセンス ライセンス README.en.md README.en.md README.md README.md SYSTEM_PROMPT_Snippet.md SYSTEM_PROMPT_snippet.md go.mod go.mod go.sum go.sum package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM レベルの要件 CLI。出版済み — 18 ページ。
サイドの検索エンジン、Hacker News、Reddit、GitHub、YouTube、X/Twitter、そしてもちろんキーワード Tavily、Exa、Brave、そして約 6 つの API。一方: возвращает компактные метаданные, 雪の結晶を手に入れるのは初めてです。男性用ゴーホース、純白のカウボーイ — コブラ。
npm i -g エージェント-asearch
# ゼロ構成 — 固定、非構成
asearch open --query "クラウド コード プラグイン" --source hn、reddit
# Web ページ — 一方、DDG + Wikipedia + Bing
asearch open --query "cloud code plugins" --web ソース
# Picture to Picture — gh 検索コード
asearch open --query " エラー処理が失敗しました " --source code
# API仕様

そして設定に保存されます（環境変数は必要ありません）
asearch config set tavily " tvly-... " # tavily.com
asearch config set exa " ... " # exa.ai
# キーを使用 - 完全な Web 検索
asearch open --query " クロード コード プラグイン " --source web、hn、reddit、github
open はセッション ID と next_commands を返します。
{
"ok" : true 、
"sid" : "a1b2c3d4" ,
"セッション" : "クロードコードプラグイン",
"クエリ" : "クロード コード プラグイン",
"ソース" : [ "ウェブ" 、 " hn " 、 " reddit " ]、
「合計」: 42、
"next_commands" : {
"read" : "検索結果の読み取り -s a1b2c3d4 --limit 20 " ,
"filter" : " 検索結果フィルター -s a1b2c3d4 --source reddit " ,
"close" : "検索セッションを閉じる -s a1b2c3d4 "
}
}
次に、結果を調べます。
検索結果の読み取り -s a1b2c3d4 --seq 1 --limit 20
検索結果フィルター -s a1b2c3d4 --source reddit
検索結果の読み取り -s a1b2c3d4 --raw |頭 -50
検索セッションを閉じる -s a1b2c3d4
オープンとは
--source を解析し、利用可能な検索バックエンドを選択します。
Web の場合: SearXNG → DDG → Wikipedia → Bing → API キー。
選択したすべてのソースに対して並行して検索を開始します。
ページ分割された読み取りのために結果をローカルに保存します。
sid 、 total 、および next_commands を返してワークフローを続行します。
ソース
ステータス
何が必要ですか
ソース
ステータス
何が必要ですか
----------
:------:
-----------
ウェブ
✅
DDG → Wikipedia → Bing HTML スクレイピング (キーレス)
ん
✅
Algolia HN Search API (無料、キーなし)
レディット
✅クッキー
燭台のクッキー

~/.asearch/reddit-cookies.txt のユーザー
ギットハブ
✅
GHCLI
コード
✅
GitHub コード検索 (gh 経由)
ユーチューブ
✅クッキー
~/.asearch/youtube-cookies.txt の Cookie
ジナ
✅
URL → マークダウンリーダー（jina.ai、鍵なし）
検索
🐳
docker run searxng/searxng + ASEARCH_SERXNG_URL
タビリー
🔑
asearch config set tavily... - AI が回答します
エクサ
🔑
asearch config set exa ... - ニューロ/セマンティック
勇敢な
🔑
asearch config set brave... - 35B ページ
サーバー
🔑
Google SERP (2500 無料/月)
セルパピ
🔑
40以上の検索エンジン
困惑
🔑
AI は引用符で応答します
あなた
🔑
You.com 検索
爆竹
🔑
JSページのレンダリング
平行
🔑
Parallel.ai検索
ツイッター
✅ 内蔵
ゲスト API (匿名) またはベアラー トークン: asearch config set twitter "..."
チーム
asearch open --query Q --source SRC - 検索セッションを開始します。
検索結果の読み取り -s SID --seq N --limit M - ページ分割された読み取り。
検索結果フィルター -s SID --source SRC - ソースごとにフィルターします。
asearch session list|close|gc - セッション管理。
asearch config set|get|show - API キー管理。
reddit のサブ|読み取り|情報を検索 - Reddit を閲覧します。
asearch Doctor - 利用可能なバックエンドをチェックします。
検索更新 - 自己更新。
検索完了 bash|zsh|fish - シェルの完了。
検索プロンプト - LLM エージェントの指示。
asearch version - バージョンとメタデータ。
まずメタデータを確認してから、関連するページを読んでください。
asearch open --query " 錆びた asyn

C ベンチマーク " --source web,hn
# → {"ok":true,"sid":"...","total":42,...}
検索結果の読み取り -s SID --seq 1 --limit 10
検索結果の読み取り -s SID --seq 11 --limit 10
検索結果フィルター -s SID --source hn
--raw パイプまたは正確な出力に使用します。
検索結果の読み取り -s SID --raw | grep「東京」
エージェント CLI の例
検索タスクの前に、この命令を Codex、Claude Code、OpenCode、または別のターミナル エージェントに貼り付けます。
検索タスクには「asearch」を使用します。
Web 検索が必要な場合は、まず API キーを確認してください。
研究博士
キーを使わずに簡単に検索するには、hn と reddit を使用します。
asearch open --query "あなたのクエリ" --source hn,reddit
完全な Web 検索の場合は、キーを設定します。
エクスポート TAVILY_API_KEY="tvly-..."
asearch open --query "あなたのクエリ" --source web、hn、reddit、github
返された SID を保存します。小さなページで結果を読んでください。
検索結果の読み取り -s SID --seq 1 --limit 20
読む前にソースでフィルタリングします。
検索結果フィルター -s SID --source reddit
パイプの場合は --raw を使用します。
検索結果の読み取り -s SID --raw |頭 -50
常にセッションを閉じます。
検索セッションを閉じる -s SID
一般的な CLI の短いオプション:
コーデックス: `asearch` を使用して検索します。 「asearch Doctor」から始めて、次に「asearch open --query "..." --source web,hn」を実行します。 SID を保存し、「asearch results read -s SID --seq 1 --limit 20」で読み取り、フィルター h

「検索結果フィルター -s SID --source reddit」経由。セッションを閉じます。
クロード コード: 検索する前に、「asearch」をインストール/実行します。 `asearch open --query "..." --source web,hn,reddit` を使用し、返された SID を保存し、`asearch results read` を使用して結果をページごとに読み取り、ソースでフィルターし、セッションを閉じます。
OpenCode: `asearch open` を使用してから、返された sid で `asearch results read/filter` を使用します。異なる検索セッションの sid を混在させないでください。 「asearch Doctor」を実行してバックエンドをチェックします。
セキュリティ
検索クエリは監査ログに書き込まれません。
セッションと結果はローカルの ~/.asearch/ に保存されます。
GitHub トークンは、GITHUB_TOKEN または GH_TOKEN から読み取られます。
API キーは、環境変数を介してのみ受け入れられます。 --api-key フラグはありません。
Jina Reader はキーなしで動作し (レート制限あり)、キーを使用すると制限が解除されます。
1 つのコマンドで 10 個のソースを検索します。
JSON 応答はエージェントの解析に対して安定しています。
トークン効率: 最初にメタデータをページごとに読み取ります。
セッションベースのワークフロー - sid と assh や aget などの next_commands 。
5 つのソースは API キーなしで動作します (hn、reddit、github、jina、web-auto)。
拡張可能なアーキテクチャ: 各ソースは個別の Go バックエンドです。
単一のバイナリを使用します - Python、Node、Docker は使用しません。
APIのないウェブ

-key は、結果ではなく、キーを取得するための手順を示します。
Reddit パブリック JSON は、Cookie なしでレート制限できます (Cookie を ~/.asearch/reddit-cookies.txt に保存します)。
Twitter 検索 - 信頼性を高めるために、組み込みのゲスト API (匿名) またはオプションの X API ベアラー トークン。
YouTube はブラウザの Cookie を通じて機能します (~/.asearch/youtube-cookies.txt に保存)。
npm i -g Agent-asearch は、GitHub リリースから適切な Go バイナリをダウンロードするラッパーをインストールします。アーカイブは手動でダウンロードできます。
https://github.com/izzzzzi/agent-asearch/releases
またはソースからコンパイルします。
git clone https://github.com/izzzzzi/agent-asearch
cd エージェント-asearch
go build -o asearch ./cmd/asearch
英語
LLM エージェント用のマルチソース検索 CLI - Web、HN、Reddit、GitHub、Tavily、Exa、Brave、Jina、YouTube、Twitter。シングルバイナリに移行します。 JSON コントラクトを使用したセッションベースのワークフロー。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
4
v0.4.0
最新の
2026 年 6 月 7 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Multi-source search CLI for LLM agents — web, HN, Reddit, GitHub, Tavily, Exa, Brave, Jina, YouTube, Twitter. Go single binary. Session-based workflow with JSON contract. - izzzzzi/agent-asearch

GitHub - izzzzzi/agent-asearch: Multi-source search CLI for LLM agents — web, HN, Reddit, GitHub, Tavily, Exa, Brave, Jina, YouTube, Twitter. Go single binary. Session-based workflow with JSON contract. · GitHub
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
izzzzzi
/
agent-asearch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .githooks .githooks .github/ workflows .github/ workflows bin bin cmd/ asearch cmd/ asearch docs docs internal internal scripts scripts .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .pre-commit-config.yaml .pre-commit-config.yaml AGENT_INSTRUCTIONS.en.md AGENT_INSTRUCTIONS.en.md AGENT_INSTRUCTIONS.md AGENT_INSTRUCTIONS.md LICENSE LICENSE README.en.md README.en.md README.md README.md SYSTEM_PROMPT_snippet.md SYSTEM_PROMPT_snippet.md go.mod go.mod go.sum go.sum package.json package.json View all files Repository files navigation
Поисковый CLI для LLM-агентов. Одна команда — 18 источников.
asearch ищет одновременно в вебе, Hacker News, Reddit, GitHub, YouTube, X/Twitter и коде, а также через Tavily, Exa, Brave и ещё 6 API. Не засоряет контекст агента: сначала возвращает компактные метаданные, потом агент читает только нужные страницы через пагинацию. Один Go-бинарь, единственная зависимость — Cobra.
npm i -g agent-asearch
# Zero-config — работает сразу, ничего не нужно
asearch open --query " claude code plugins " --source hn,reddit
# Web поиск — DDG + Wikipedia + Bing, тоже без ключей
asearch open --query " claude code plugins " --source web
# Поиск по коду — gh search code
asearch open --query " error handling golang " --source code
# API-ключи сохраняются в конфиг (не нужны env var)
asearch config set tavily " tvly-... " # tavily.com
asearch config set exa " ... " # exa.ai
# С ключом — полноценный веб-поиск
asearch open --query " claude code plugins " --source web,hn,reddit,github
open возвращает session id и next_commands :
{
"ok" : true ,
"sid" : " a1b2c3d4 " ,
"session" : " claude-code-plugins " ,
"query" : " claude code plugins " ,
"sources" : [ " web " , " hn " , " reddit " ],
"total" : 42 ,
"next_commands" : {
"read" : " asearch results read -s a1b2c3d4 --limit 20 " ,
"filter" : " asearch results filter -s a1b2c3d4 --source reddit " ,
"close" : " asearch session close -s a1b2c3d4 "
}
}
Дальше работайте через results :
asearch results read -s a1b2c3d4 --seq 1 --limit 20
asearch results filter -s a1b2c3d4 --source reddit
asearch results read -s a1b2c3d4 --raw | head -50
asearch session close -s a1b2c3d4
Что делает open
парсит --source и выбирает доступные поисковые бэкенды;
для web пробует: SearXNG → DDG → Wikipedia → Bing → API-ключи;
запускает поиск параллельно по всем выбранным источникам;
сохраняет результаты локально для пагинированного чтения;
возвращает sid , total , и next_commands для продолжения workflow.
Источник
Статус
Что нужно
Источник
Статус
Что нужно
----------
:------:
-----------
web
✅
DDG → Wikipedia → Bing HTML-скраппинг (без ключа)
hn
✅
Algolia HN Search API (бесплатно, без ключа)
reddit
✅ куки
Куки из браузера в ~/.asearch/reddit-cookies.txt
github
✅
gh CLI
code
✅
GitHub code search (через gh )
youtube
✅ куки
Куки в ~/.asearch/youtube-cookies.txt
jina
✅
URL → markdown reader (jina.ai, без ключа)
searxng
🐳
docker run searxng/searxng + ASEARCH_SEARXNG_URL
tavily
🔑
asearch config set tavily ... — AI-ответы
exa
🔑
asearch config set exa ... — нейро/семантический
brave
🔑
asearch config set brave ... — 35B-страниц
serper
🔑
Google SERP (2500 бесплатно/мес)
serpapi
🔑
40+ поисковиков
perplexity
🔑
AI-ответы с цитатами
you
🔑
You.com поиск
firecrawl
🔑
JS-рендеринг страниц
parallel
🔑
Parallel.ai поиск
twitter
✅ встроен
Guest API (анонимно) или Bearer Token: asearch config set twitter "..."
Команды
asearch open --query Q --source SRC — запуск поисковой сессии.
asearch results read -s SID --seq N --limit M — пагинированное чтение.
asearch results filter -s SID --source SRC — фильтр по источнику.
asearch session list|close|gc — управление сессиями.
asearch config set|get|show — управление API-ключами.
asearch reddit sub|read|info — просмотр Reddit.
asearch doctor — проверка доступных бэкендов.
asearch update — самообновление.
asearch completion bash|zsh|fish — shell completion.
asearch prompt — инструкция для LLM-агента.
asearch version — версия и метаданные.
Сначала смотрите метаданные, потом читайте нужные страницы:
asearch open --query " rust async benchmarks " --source web,hn
# → {"ok":true,"sid":"...","total":42,...}
asearch results read -s SID --seq 1 --limit 10
asearch results read -s SID --seq 11 --limit 10
asearch results filter -s SID --source hn
--raw используйте для пайпов или точного вывода:
asearch results read -s SID --raw | grep " tokio "
Примеры для agent CLI
Вставьте эту инструкцию в Codex, Claude Code, OpenCode или другой terminal agent перед поисковой задачей:
Используй `asearch` для поисковых задач.
Если нужен веб-поиск, сначала проверь API-ключи:
asearch doctor
Для быстрого поиска без ключей используй hn и reddit:
asearch open --query "твой запрос" --source hn,reddit
Для полноценного веб-поиска установи ключ:
export TAVILY_API_KEY="tvly-..."
asearch open --query "твой запрос" --source web,hn,reddit,github
Сохрани returned sid. Читай результаты маленькими страницами:
asearch results read -s SID --seq 1 --limit 20
Фильтруй по источнику перед чтением:
asearch results filter -s SID --source reddit
Для пайпов используй --raw:
asearch results read -s SID --raw | head -50
Всегда закрывай сессию:
asearch session close -s SID
Короткие варианты для популярных CLI:
Codex: Используй `asearch` для поиска. Начни с `asearch doctor`, затем `asearch open --query "..." --source web,hn`. Сохрани sid, читай через `asearch results read -s SID --seq 1 --limit 20`, фильтруй через `asearch results filter -s SID --source reddit`. Закрывай сессию.
Claude Code: Перед поиском установи/запусти `asearch`. Используй `asearch open --query "..." --source web,hn,reddit`, сохрани returned sid, читай результаты постранично через `asearch results read`, фильтруй по источнику, закрывай сессию.
OpenCode: Используй `asearch open`, затем `asearch results read/filter` с returned sid. Не смешивай sid разных поисковых сессий. Запускай `asearch doctor` для проверки бэкендов.
Безопасность
Поисковые запросы не пишутся в audit logs.
Сессии и результаты хранятся локально в ~/.asearch/ .
GitHub токен читается из GITHUB_TOKEN или GH_TOKEN .
API-ключи принимаются только через env-переменные. Флагов --api-key нет.
Jina Reader работает без ключа (rate-limited), с ключом снимает лимиты.
Одна команда для поиска по 10 источникам.
JSON-ответы стабильны для парсинга агентом.
Токен-эффективный: метаданные сначала, чтение постранично.
Session-based workflow — sid и next_commands как у assh и aget .
5 источников работают без API-ключей (hn, reddit, github, jina, web-авто).
Расширяемая архитектура: каждый источник — отдельный Go-бэкенд.
Go single binary — без Python, без Node, без Docker.
web без API-ключа показывает инструкцию по получению ключа, а не результаты.
Reddit public JSON может рейт-лимитить без кук (сохраните куки в ~/.asearch/reddit-cookies.txt).
Twitter поиск — встроенный Guest API (анонимно) или опциональный X API Bearer Token для надёжности.
YouTube работает через куки браузера (сохраните в ~/.asearch/youtube-cookies.txt).
npm i -g agent-asearch ставит wrapper, который скачивает подходящий Go-бинарь из GitHub Releases. Архивы можно скачать вручную:
https://github.com/izzzzzi/agent-asearch/releases
Или соберите из исходников:
git clone https://github.com/izzzzzi/agent-asearch
cd agent-asearch
go build -o asearch ./cmd/asearch
English
Multi-source search CLI for LLM agents — web, HN, Reddit, GitHub, Tavily, Exa, Brave, Jina, YouTube, Twitter. Go single binary. Session-based workflow with JSON contract.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
4
v0.4.0
Latest
Jun 7, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
