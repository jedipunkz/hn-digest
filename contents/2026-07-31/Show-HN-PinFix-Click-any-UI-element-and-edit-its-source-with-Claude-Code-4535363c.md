---
source: "https://github.com/ForeverSc/pinfix"
hn_url: "https://news.ycombinator.com/item?id=49122903"
title: "Show HN: PinFix – Click any UI element and edit its source with Claude Code"
article_title: "GitHub - ForeverSc/pinfix: 📌 Edit UI with Claude Code using on-page comments. · GitHub"
author: "foreversc2"
captured_at: "2026-07-31T13:40:19Z"
capture_tool: "hn-digest"
hn_id: 49122903
score: 1
comments: 0
posted_at: "2026-07-31T13:29:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PinFix – Click any UI element and edit its source with Claude Code

- HN: [49122903](https://news.ycombinator.com/item?id=49122903)
- Source: [github.com](https://github.com/ForeverSc/pinfix)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T13:29:58Z

## Translation

タイトル: HN を表示: PinFix – 任意の UI 要素をクリックし、そのソースをクロード コードで編集します
記事のタイトル: GitHub - ForeverSc/pinfix: 📌 ページ上のコメントを使用してクロード コードで UI を編集します。 · GitHub
説明: 📌 ページ上のコメントを使用してクロードコードで UI を編集します。 - ForeverSc/ピンフィックス

記事本文:
GitHub - ForeverSc/pinfix: 📌 ページ上のコメントを使用してクロード コードで UI を編集します。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
フォーエバースク
/
ピンフィックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
49 コミット 49 コミット .github/ workflows .github/ workflows .husky .husky docs docs 例 例 パッケージ パッケージ スクリプト .gitignore .gitignore .npmrc .npmrc .prettierignore .prettierignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md README_ZH.md README_ZH.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.mjs prettier.config.mjs release.config.mjs release.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
デザインにコメントを残すようにフロントエンド ページを編集します。
プレビュー.mp4
任意の UI をクリックし、変更したい内容を記述すると、PinFix によって Claude Code がソース コードを正確に見つけて、リアルタイムで編集を適用できるようになります。
コンテキストの切り替えはありません。ファイルパスをコピー＆ペーストする必要はありません。ポイントして説明し、HMR が変更を適用するのを確認するだけです。
従来のクロード コードのワークフローでは、コードベースのどこを変更する必要があるかを説明する必要があります。 PinFix はこれを反転します。ブラウザーで要素を視覚的に選択すると、正確なソース ファイル、行、列がすでに認識されます。会話は完全なコンテキストから始まります。
ページはコンテキストです — 変更したい内容を UI 上で直接指定します。ウィンドウの切り替え、ファイル パス、場所の説明は必要ありません。
リアルタイム編集 — クロード コードはソース ファイルを直接更新し、HMR が結果を即座に表示します。
視覚的な選択 - Alt+Shift+Z を押して十字モードに入り、マウスを移動して要素を強調表示し、クリックしてピンを配置します。
デザイン調整パネル — 単語だけでは遅すぎる場合に、テキスト、レイアウト、間隔、サイズ、色、境界線、タイポグラフィの変更にインスペクター スタイルのコントロールを使用します。
フレームワークに依存しない — W

React、Vue、Svelte、または任意の JSX/TSX ベースのフレームワークと連携します。
ゼロ構成 — ビルド構成内の 1 つのプラグイン行。チャネル サーバーは自動的に生成され、クリーンアップされます。
npm install -D @pinfix/plugin
プラグインをビルド構成に追加します。
// vite.config.ts
「@pinfix/plugin/vite」から pinfix をインポートします
デフォルトのdefineConfigをエクスポート( {
プラグイン: [ pinfix() ] 、
})
ウェブパック
// webpack.config.js
「@pinfix/plugin/webpack」から pinfix をインポートします
デフォルトのエクスポート {
プラグイン: [ pinfix() ] 、
}
Rspack / Rsbuild
// rsbuild.config.ts
「@pinfix/plugin/rspack」から pinfix をインポートします
デフォルトのエクスポート {
ツール: {
rspack : { プラグイン : [ pinfix ( ) ] } 、
} 、
}
次に、通常どおり開発サーバーを起動します。 PinFix は開発モードで自動的にアクティブになります。
開発サーバーを起動します ( npm run dev )
Alt + Shift + Z (Mac では Option + Shift + Z) を押して注釈モードに入ります。
任意のコンポーネントの上にマウスを移動します。青い枠線で強調表示されます。
クリックして要素上にピンを配置します
変更リクエストを入力するか、「デザインの調整」を開いてレイアウト、間隔、サイズ、色、境界線、テキストのコントロールを調整します。
クロード コードは、選択されたソースの場所と構造化された視覚的な変更を受け取ります
Claude Code は応答をストリーミングし、ソース コードを編集します
HMR は変更を適用します - 結果をすぐに確認します
会話を続けて改良を繰り返します
┌─────────────┐ WebSocket ┌─────────┐ クロード エージェント SDK
│ ブラウザクライアント │ ◄─────────────► │ チャネルサーバー │ ◄───────────────► クロードコード
│ (シャドウ DOM) │ ポート 24816 │ (自動生成) │
━━━━━━┘ ━━━━━━

──┘
ビルド プラグインは、JSX/TSX/Vue ファイルを変換して、ファイル パス、行、列のメタデータを含む data-pinfix-source 属性を挿入します。
クライアント オーバーレイは Shadow DOM 内でレンダリングされ、アプリのスタイルから分離されます。ピンの配置、チャット UI、デザイン調整コントロール、WebSocket 通信を処理します。
チャネル サーバーは開発サーバーと並行して自動的に生成されます。アクティブなピンは、完全なプロジェクト コンテキストを持つワークスペース レベルのクロード コード セッションを使用します。
ピンフィックス ( {
port : 24816 , // WebSocket ポート (デフォルト: 24816)
hotkey : 'alt+shift+z' , // アクティブ化ホットキー
fab : true , // フローティングアクションボタンを表示
プロンプト : 'カスタム システム プロンプト...' , // クロード コードの追加コンテキスト
scapeTags : [ 'Layout' , 'Provider' ] , // これらのラッパー コンポーネントをスキップします
一致：/\。 ( t s x | j s x | v u e ) $ / , // 一致するファイルのみを変換します
exclude : / no d e _ mod u l e s / , // 変換から除外
debug : false , // デバッグログを有効にする
})
サポートされているバンドラー
バンドラー
インポートパス
ステータス
ヴィート 5+
@pinfix/プラグイン/vite
安定した
ウェブパック 5
@pinfix/プラグイン/webpack
安定した
アールズパック2
@pinfix/プラグイン/rspack
安定した
要件
クロードコードがマシンにインストールされ、使用可能になります
git clone https://github.com/foreversc/pinfix.git
cd ピンフィックス
pnpmインストール
pnpmビルド
pnpm dev:vite-react # Vite + React サンプルを実行する
ライセンス
📌 ページ上のコメントを使用して Claude Code で UI を編集します。
eversc.github.io/pinfix/ トピック
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

📌 Edit UI with Claude Code using on-page comments. - ForeverSc/pinfix

GitHub - ForeverSc/pinfix: 📌 Edit UI with Claude Code using on-page comments. · GitHub
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
ForeverSc
/
pinfix
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
49 Commits 49 Commits .github/ workflows .github/ workflows .husky .husky docs docs examples examples packages packages scripts scripts .gitignore .gitignore .npmrc .npmrc .prettierignore .prettierignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md README_ZH.md README_ZH.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.mjs prettier.config.mjs release.config.mjs release.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
Edit frontend pages like leaving comments on a design.
preview.mp4
Click any UI, describe what you want to change, and PinFix lets Claude Code precisely find the source code and apply the edit in real time.
No context switching. No copy-pasting file paths. Just point, describe, and see HMR apply the change.
Traditional Claude Code workflows require you to explain where in the codebase something needs to change. PinFix flips this — you visually select the element in the browser, and it already knows the exact source file, line, and column. Your conversation starts with full context.
The page is the context — Point to what you want to change directly on the UI, no window switching, file paths, or location explanations needed
Real-time edits — Claude Code updates your source files directly, and HMR shows the result instantly
Visual selection — Press Alt+Shift+Z to enter crosshair mode, hover to highlight elements, and click to place a pin
Design adjustment panel — Use inspector-style controls for text, layout, spacing, size, color, border, and typography changes when words alone are too slow
Framework Agnostic — Works with React, Vue, Svelte, or any JSX/TSX-based framework.
Zero Config — One plugin line in your build config. The channel server spawns and cleans up automatically.
npm install -D @pinfix/plugin
Add the plugin to your build config:
// vite.config.ts
import pinfix from '@pinfix/plugin/vite'
export default defineConfig ( {
plugins : [ pinfix ( ) ] ,
} )
Webpack
// webpack.config.js
import pinfix from '@pinfix/plugin/webpack'
export default {
plugins : [ pinfix ( ) ] ,
}
Rspack / Rsbuild
// rsbuild.config.ts
import pinfix from '@pinfix/plugin/rspack'
export default {
tools : {
rspack : { plugins : [ pinfix ( ) ] } ,
} ,
}
Then start your dev server as usual. PinFix activates automatically in development mode.
Start your dev server ( npm run dev )
Press Alt + Shift + Z (Option + Shift + Z on Mac) to enter annotation mode
Hover over any component — it highlights with a blue border
Click to place a pin on the element
Type your change request, or open Adjust design to tweak layout, spacing, size, color, border, and text controls
Claude Code receives the selected source location plus your structured visual changes
Claude Code streams a response and edits your source code
HMR applies the change — see the result immediately
Continue the conversation for iterative refinements
┌─────────────────┐ WebSocket ┌─────────────────┐ Claude Agent SDK
│ Browser Client │ ◄──────────────────────► │ Channel Server │ ◄─────────────────────► Claude Code
│ (Shadow DOM) │ port 24816 │ (auto-spawned) │
└─────────────────┘ └─────────────────┘
Build plugin transforms your JSX/TSX/Vue files to inject data-pinfix-source attributes with file path, line, and column metadata.
Client overlay renders inside Shadow DOM — isolated from your app's styles. Handles pin placement, chat UI, design adjustment controls, and WebSocket communication.
Channel server spawns automatically alongside your dev server. The active pin uses a workspace-level Claude Code session with full project context.
pinfix ( {
port : 24816 , // WebSocket port (default: 24816)
hotkey : 'alt+shift+z' , // Activation hotkey
fab : true , // Show floating action button
prompt : 'Custom system prompt...' , // Additional context for Claude Code
escapeTags : [ 'Layout' , 'Provider' ] , // Skip these wrapper components
match : / \. ( t s x | j s x | v u e ) $ / , // Only transform matching files
exclude : / n o d e _ m o d u l e s / , // Exclude from transform
debug : false , // Enable debug logging
} )
Supported Bundlers
Bundler
Import Path
Status
Vite 5+
@pinfix/plugin/vite
Stable
Webpack 5
@pinfix/plugin/webpack
Stable
Rspack 2
@pinfix/plugin/rspack
Stable
Requirements
Claude Code installed and usable on your machine
git clone https://github.com/foreversc/pinfix.git
cd pinfix
pnpm install
pnpm build
pnpm dev:vite-react # Run the Vite + React example
License
📌 Edit UI with Claude Code using on-page comments.
foreversc.github.io/pinfix/ Topics
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
