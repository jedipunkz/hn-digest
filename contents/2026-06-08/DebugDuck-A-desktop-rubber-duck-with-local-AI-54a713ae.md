---
source: "https://github.com/CarlosVallejoRuiz/DebugDuck"
hn_url: "https://news.ycombinator.com/item?id=48450641"
title: "DebugDuck – A desktop rubber duck with local AI"
article_title: "GitHub - CarlosVallejoRuiz/DebugDuck: Rubber Duck Debugging desktop widget — Tauri v2 + React + AI · GitHub"
author: "CarlosVallejoR"
captured_at: "2026-06-08T21:38:55Z"
capture_tool: "hn-digest"
hn_id: 48450641
score: 1
comments: 0
posted_at: "2026-06-08T19:42:15Z"
tags:
  - hacker-news
  - translated
---

# DebugDuck – A desktop rubber duck with local AI

- HN: [48450641](https://news.ycombinator.com/item?id=48450641)
- Source: [github.com](https://github.com/CarlosVallejoRuiz/DebugDuck)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T19:42:15Z

## Translation

タイトル: DebugDuck – ローカル AI を備えたデスクトップ ラバー アヒル
記事タイトル: GitHub - CarlosVallejoRuiz/DebugDuck: ラバーダックのデスクトップウィジェットのデバッグ — Tauri v2 + React + AI · GitHub
説明: ラバーダックのデバッグデスクトップウィジェット — Tauri v2 + React + AI - CarlosVallejoRuiz/DebugDuck

記事本文:
GitHub - CarlosVallejoRuiz/DebugDuck: ラバーダックのデバッグデスクトップウィジェット — Tauri v2 + React + AI · GitHub
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
カルロス・バジェホ・ルイス
/
デバッグダック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーションオプション

イオン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .agents/ スキル .agents/ スキル .claude/ スキル .claude/ スキル .github/ ワークフロー .github/ ワークフロー 画像 画像 パブリック パブリック スクリプト スクリプト src-tauri src-tauri src src .gitignore .gitignore CLAUDE.md CLAUDE.md README.md README.md eslint.config.js eslint.config.js Index.html Index.html package-lock.json package-lock.json package.json package.json skill-lock.json skill-lock.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル AI を備えたラバー アヒル。常に浮いている。常にあなたを判断します。
ラバーダックのデバッグですが、ローカル AI、音声、アニメーション、そして自分の感情を使用します。
すべてのウィンドウの上に常駐する macOS 用のフローティング デスクトップ ウィジェット。気を散らすことも、サブスクリプションも、マシンからデータが流出することもありません。あなたとあなたのアヒルとあなたの現地言語モデルだけです。
話す。アヒルは考えます。アヒルは答えます。あまり長く放っておくと、彼の機嫌が悪くなってしまいます。
🎙️ 音声アクティベーション — アヒルをダブルクリックして話す、Web Speech API
🧠 100% ローカル AI — インターネットなしで、localhost:1234 の LM Studio に接続します
📡 リアルタイム ストリーミング — 応答はトークンごとに表示されます
🎭 2 つの人格 — プログラマー (ソクラテス) または一般 (意見を持つ)
😈 残酷スライダー — 患者の指導者から「高齢者への明らかな質問」まで
🥚 たまごっちモード — アヒルには反応に影響を与える気分があります
🎬 状態アニメーション — アイドル、リスニング、思考、応答 (140 フレーム PNG)
🎉 エウレカボタン — フルスクリーンの紙吹雪 + 勝利カウンター
🍅 統合されたポモドーロ — タイマー

ネイティブ通知付きで 25 分
💬 会話の記憶 — 自動圧縮で文脈を記憶
🖱️ ピクセルパーフェクトなクリックスルー — アヒルは透明な領域のクリックを妨害しません
🔍 自動モデル検出 — LM Studio にロードしたモデルを検出します
📐 構成可能な位置 — ウィジェットを任意の隅に移動するための 3×3 グリッド
要件
最小値
おすすめ
macOS
12 モントレー
14 ソノマ+
RAM
8GB
16GB
チップ
インテル/アップルシリコン
アップルシリコン
LMスタジオ
任意のバージョン
最新バージョン
マイク
必須
—
🤖 LMスタジオのセットアップ
LM Studioをダウンロードしてインストールします
それを開いて [検出] タブに移動してモデルを検索します
モデル
VRAM/RAM
注意事項
ミストラライ/ミニストラル-3B
～3GB
軽くて速い、レスポンスが良い
ミストラライ/ミストラル-7B-命令
～6GB
推奨 - 理想的なバランス
ミストラライ/Mistral-7B-Instruct-v0.3
～6GB
安定した代替品
metal-llama/Llama-3.1-8B-Instruct
～7GB
技術的な質問にもしっかり対応
google/gemma-2-9b-it
～9GB
コードに最適
⚠️ 「考える」モデル (Qwen3、DeepSeek-R1、理由のある名前) は避けてください。内部推論で余分なトークンを使用し、応答を遅くします。 DebugDuck はこれらをサポートしていますが、エクスペリエンスはさらに悪くなります。
LM Studioで「ローカルサーバー」タブを開きます。
モデルを選択し、「モデルのロード」を押します。
「ステータス」トグルをアクティブにすると、「ポート 1234 で実行中」と表示されます。
アヒルは起動時にアクティブなモデルを自動的に検出します。モデルをホットに変更するには:
設定を開く ⚙️ → アクティブなモデルには検出されたモデルが表示されます
↺ を押して検出を更新します
#Rusttoolchain
カール --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs |しー
#Node.js 20+
# (推奨: nvm または fnm の使用)
クローンを作成してブートする
git クローン https://github.com/tu-usuario/DebugDuck.git
cd デバッグダック
npmインストール

開発ビルド (必須 - tauri dev は使用しないでください)
⚠️ ウィジェットは SpeechRecognition と SpeechSynthesis を使用します。これらには、macOS で署名された .app バンドルが必要です。 tauri dev はこれには機能しません。
npm run tauri build -- --debug && \
codedesign --sign - --force --deep \
--entitlements src-tauri/entitlements.plist \
src-tauri/target/debug/bundle/macos/DebugDuck.app && \
src-tauri/target/debug/bundle/macos/DebugDuck.app を開きます
🦆 DebugDuckの使い方
アクション
何をするのですか
アヒルをダブルクリックします
マイクをアクティブにする
質問を話してください
アヒルは聞いて書き写します
応答を待ちます
考えながら「頭をかく」アニメーション
吹き出しを読む
コミックスタイルのスクロール可能な応答
エウレカ！
できました — 紙吹雪 🎉、カウンター +1、幸福度 +10
ノブ
25 分のポモドーロ タイマーを開始します 🍅
ダブルクリック (バブルが開いた状態)
閉じてアイドル状態に戻る
設定⚙️
アヒルの上にカーソルを置く → 歯車が表示される → クリックして設定を開きます。
性格 — 🦆 プログラマー (ソクラテス的モード、直接的な解決策を与えない) または 🌍 一般 (あらゆるトピックについて話す)
たまごっちモード — ムードシステムをアクティブにします
残酷スライダー — たまごっちが無効な場合にのみ表示されます
メモリ — アヒルは最大 4 つのメッセージと圧縮された概要を記憶します
位置 — ウィジェットを移動する 3×3 グリッド
有効にすると、アヒルの気分が残酷さスライダーの代わりに反応のトーンを制御します。
幸福感を変える出来事:
アヒルの左側にステータス バッジ (絵文字) が表示されます。カーソルを合わせるとライフバーが表示されます。
テクノロジー
使用する
タウリ v2
macOS ネイティブ シェル、透明ウィンドウ、Rust コマンド
リアクト19
宣言型 UI、ロジックのフック
TypeScript
フロントエンド全体での静的型付け
ヴィテ 8
ビルドツール、HMR、PNG アセットのインポート
テイルウィンド CSS v4
ユーティリティファーストスタイル

t
ズスタンド
localStorage に永続化されたグローバル状態
錆び/リクエスト
LM Studio への SSE ストリーミング (署名付きバンドルの CORS バイパス)
ウェブスピーチAPI
外部依存性のない音声認識
キャンバスAPI
フレームごとのアニメーション + クリックスルー用のアルファ サンプリング
フックアーキテクチャ
src/フック/
§── useVoiceRecognition.ts # Web Speech API + 専門用語修正
§── useAIResponse.ts # LM Studio + SSE ストリーミング + 圧縮メモリ
§── useAnimation.ts # 状態別アニメーションシステム
§── usePomodoro.ts # タイマー 25min + ネイティブ通知
§── useTamagotchi.ts # ムードシステム + 自動減衰
└── useWindowPosition.ts # Rustによる画面配置
🗺️ロードマップ
Windows 用にビルド (透明性と権限を調整)
LM Studio の代替バックエンドとして Ollama をサポート
会話履歴はセッション間で保持される
マルチダック モード (複数のウィジェット)
会話を Markdown としてエクスポートする
マイクを有効にするためのグローバル キーボード ショートカット
MIT — 好きなようにしてください。ただし、アヒルがバグの解決に協力してくれるなら、少なくともエウレカを与えてください。
フラストレーション、カフェイン、ラバーダックのデバッグで作られました。
ラバーダックデバッグデスクトップウィジェット — Tauri v2 + React + AI
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
2
v0.1.1
最新の
2026 年 6 月 4 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Rubber Duck Debugging desktop widget — Tauri v2 + React + AI - CarlosVallejoRuiz/DebugDuck

GitHub - CarlosVallejoRuiz/DebugDuck: Rubber Duck Debugging desktop widget — Tauri v2 + React + AI · GitHub
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
CarlosVallejoRuiz
/
DebugDuck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .agents/ skills .agents/ skills .claude/ skills .claude/ skills .github/ workflows .github/ workflows images images public public scripts scripts src-tauri src-tauri src src .gitignore .gitignore CLAUDE.md CLAUDE.md README.md README.md eslint.config.js eslint.config.js index.html index.html package-lock.json package-lock.json package.json package.json skills-lock.json skills-lock.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Tu pato de goma con IA local. Siempre flotando. Siempre juzgándote.
El rubber duck debugging, pero con IA local, voz, animaciones y sus propios sentimientos.
Un widget de escritorio flotante para macOS que vive encima de todas tus ventanas. Sin distracciones, sin suscripciones, sin datos que salen de tu máquina. Solo tú, tu pato y tu modelo de lenguaje local.
Hablas. El pato piensa. El pato responde. Si lo abandonas mucho tiempo, se pone de mal humor.
🎙️ Activación por voz — doble clic en el pato para hablar, Web Speech API
🧠 IA 100% local — conecta con LM Studio en localhost:1234 , sin internet
📡 Streaming en tiempo real — las respuestas aparecen token a token
🎭 Dos personalidades — Programador (socrático) o General (opinionado)
😈 Slider de crueldad — de mentor paciente a "pregunta obvia para cualquier senior"
🥚 Modo Tamagotchi — el pato tiene un estado de ánimo que afecta sus respuestas
🎬 Animaciones por estados — idle, escuchando, pensando, respondiendo (140 frames PNG)
🎉 Botón Eureka — confeti fullscreen + contador de victorias
🍅 Pomodoro integrado — timer 25min con notificación nativa
💬 Memoria de conversación — recuerda contexto con compresión automática
🖱️ Click-through pixel-perfect — el pato no intercepta clicks en áreas transparentes
🔍 Detección automática de modelo — detecta qué modelo tienes cargado en LM Studio
📐 Posición configurable — grid 3×3 para mover el widget a cualquier esquina
Requisito
Mínimo
Recomendado
macOS
12 Monterey
14 Sonoma+
RAM
8 GB
16 GB
Chip
Intel / Apple Silicon
Apple Silicon
LM Studio
cualquier versión
última versión
Micrófono
requerido
—
🤖 Configuración de LM Studio
Descarga LM Studio e instálalo
Ábrelo y ve a la pestaña Discover para buscar modelos
Modelo
VRAM / RAM
Notas
mistralai/Ministral-3B
~3 GB
Ligero y rápido, buenas respuestas
mistralai/Mistral-7B-Instruct
~6 GB
Recomendado — equilibrio ideal
mistralai/Mistral-7B-Instruct-v0.3
~6 GB
Alternativa estable
meta-llama/Llama-3.1-8B-Instruct
~7 GB
Sólido para preguntas técnicas
google/gemma-2-9b-it
~9 GB
Excelente para código
⚠️ Evita modelos "thinking" (Qwen3, DeepSeek-R1, nombres con reasoning ). Usan tokens extra en razonamiento interno y ralentizan las respuestas. DebugDuck los soporta, pero la experiencia es peor.
Abre la pestaña Local Server en LM Studio
Selecciona un modelo y pulsa Load Model
Activa el toggle Status → debe mostrar Running on port 1234
El pato detecta automáticamente el modelo activo al arrancar. Para cambiar de modelo en caliente:
Abre ajustes ⚙️ → Modelo activo muestra el modelo detectado
Pulsa ↺ para refrescar la detección
# Rust toolchain
curl --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs | sh
# Node.js 20+
# (Recomendado: usar nvm o fnm)
Clonar y arrancar
git clone https://github.com/tu-usuario/DebugDuck.git
cd DebugDuck
npm install
Build de desarrollo (requerido — no usar tauri dev )
⚠️ El widget usa SpeechRecognition y SpeechSynthesis , que requieren un bundle .app firmado en macOS. tauri dev no funciona para esto.
npm run tauri build -- --debug && \
codesign --sign - --force --deep \
--entitlements src-tauri/entitlements.plist \
src-tauri/target/debug/bundle/macos/DebugDuck.app && \
open src-tauri/target/debug/bundle/macos/DebugDuck.app
🦆 Cómo usar DebugDuck
Acción
Qué hace
Doble clic en el pato
Activa el micrófono
Habla tu pregunta
El pato escucha y transcribe
Espera la respuesta
Animación "rasca cabeza" mientras piensa
Lee el bocadillo
Respuesta scrolleable estilo cómic
¡Eureka!
Entendiste — confeti 🎉, +1 contador, +10 felicidad
Pomo
Inicia timer Pomodoro de 25 min 🍅
Doble clic (con bocadillo abierto)
Cierra y vuelve a idle
Ajustes ⚙️
Hover sobre el pato → aparece el engranaje → clic para abrir ajustes:
Personalidad — 🦆 Programador (modo socrático, no da soluciones directas) o 🌍 General (habla de cualquier tema)
Modo Tamagotchi — activa el sistema de estado de ánimo
Slider de crueldad — solo visible cuando Tamagotchi está desactivado
Memoria — el pato recuerda hasta 4 mensajes + resumen comprimido
Posición — grid 3×3 para mover el widget
Cuando está activado, el estado de ánimo del pato controla el tono de sus respuestas en lugar del slider de crueldad:
Eventos que modifican la felicidad:
El badge de estado (emoji) aparece a la izquierda del pato. Hover sobre él para ver la barra de vida.
Tecnología
Uso
Tauri v2
Shell nativo macOS, ventana transparente, comandos Rust
React 19
UI declarativa, hooks para lógica
TypeScript
Tipado estático en todo el frontend
Vite 8
Build tool, HMR, import de assets PNG
Tailwind CSS v4
Estilos utility-first
Zustand
Estado global con persistencia en localStorage
Rust / reqwest
Streaming SSE a LM Studio (bypass CORS en bundle firmado)
Web Speech API
Reconocimiento de voz sin dependencias externas
Canvas API
Animaciones frame-by-frame + alpha sampling para click-through
Arquitectura de hooks
src/hooks/
├── useVoiceRecognition.ts # Web Speech API + corrección de términos técnicos
├── useAIResponse.ts # LM Studio + streaming SSE + memoria comprimida
├── useAnimation.ts # Sistema de animaciones por estado
├── usePomodoro.ts # Timer 25min + notificación nativa
├── useTamagotchi.ts # Sistema de estado de ánimo + decay automático
└── useWindowPosition.ts # Posicionamiento en pantalla via Rust
🗺️ Roadmap
Build para Windows (ajustar transparencia y permisos)
Soporte para Ollama como backend alternativo a LM Studio
Historial de conversaciones persistente entre sesiones
Modo multi-pato (más de un widget)
Export de la conversación como Markdown
Atajos de teclado globales para activar el micrófono
MIT — haz lo que quieras, pero si el pato te ayuda a resolver un bug, mínimo dale un Eureka.
Hecho con frustración, cafeína y rubber duck debugging.
Rubber Duck Debugging desktop widget — Tauri v2 + React + AI
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
2
v0.1.1
Latest
Jun 4, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
