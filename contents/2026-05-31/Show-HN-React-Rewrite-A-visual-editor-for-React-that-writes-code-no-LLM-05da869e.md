---
source: "https://github.com/donghaxkim/react-rewrite"
hn_url: "https://news.ycombinator.com/item?id=48339764"
title: "Show HN: React-Rewrite – A visual editor for React that writes code, no LLM"
article_title: "GitHub - donghaxkim/react-rewrite: visual editor (figma) for your react apps, edit UI elements live and directly changes source files. no API key, no AI.. · GitHub"
author: "donghaxkim"
captured_at: "2026-05-31T00:26:19Z"
capture_tool: "hn-digest"
hn_id: 48339764
score: 2
comments: 0
posted_at: "2026-05-30T19:24:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: React-Rewrite – A visual editor for React that writes code, no LLM

- HN: [48339764](https://news.ycombinator.com/item?id=48339764)
- Source: [github.com](https://github.com/donghaxkim/react-rewrite)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T19:24:33Z

## Translation

タイトル: Show HN: React-Rewrite – LLM なしでコードを記述する React 用のビジュアル エディター
記事タイトル: GitHub - donghaxkim/react-rewrite: React アプリ用のビジュアル エディター (figma)、UI 要素をライブで編集し、ソース ファイルを直接変更します。 API キーも AI もありません。 · GitHub
説明: React アプリ用のビジュアル エディター (figma)、UI 要素をライブで編集し、ソース ファイルを直接変更します。 API キーも AI もありません。 - donghaxkim/react-rewrite

記事本文:
GitHub - donghaxkim/react-rewrite: React アプリ用のビジュアル エディター (figma)、UI 要素をライブで編集し、ソース ファイルを直接変更します。 API キーも AI もありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ドンハクキム
/
反応リライト
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
269 コミット 269 コミット .changeset .changeset .github/ workflows .github/ workflows パッケージ パッケージ .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
React-rewrite を使用すると、React アプリをローカルで実行しながら視覚的に編集し、それらの変更をプロジェクト内のソース ファイルに自動的に書き戻すことができます。
これはローカル開発用に構築されており、開発サーバーの前でプロキシを開き、ページにオーバーレイを挿入することで機能します。
完全なデモ:
https://x.com/imdonghakim/status/2038230475894899119
このリポジトリをダウンロードしたり複製したりする必要はありません。
React アプリのルートから:
npm install -D 反応リライト-cli
開発サーバーを起動し、2 番目のターミナルで次のコマンドを実行します。
npx 反応リライト
最初にインストールせずに試してみたい場合は、次のようにします。
npx 反応リライト-cli@latest
何をするのか
要素を選択し、そのコンポーネント名、ファイル パス、行番号を確認します。
サポートされている Tailwind ベースのレイアウト、間隔、サイズ、タイポグラフィ、および色のプロパティを編集します
テキストをダブルクリックしてインライン編集します
要素のコピー、貼り付け、複製
複数の変更をステージングし、「確認」で適用する
進行中のキャンバスの変更を元に戻し、適用された変更を変更ログで確認します。
サポートされているアプリのセットアップ: Next.js、Vite、Create React App
プロパティ エディターを使用する場合は、Tailwind CSS をお勧めします。テキスト編集と一部の構造アクションは Tailwind に依存しません。
編集する React アプリのルートでこれを実行します。
npm install -D 反応-r

ewrite-cli
最初にインストールしたくない場合は、 npx act-rewrite-cli@latest を使用して直接実行することもできます。
いつものように React 開発サーバーを起動します。
2 番目のターミナルで、同じプロジェクト ルートから次を実行します。
npx 反応リライト
自動検出で正しいポートが選択されない場合は、明示的にポートを渡します。
npx 反応リライト 3000
このツールはブラウザでローカル プロキシを開き、編集オーバーレイを表示し、確認された変更をプロジェクト内のファイルに書き込みます。
要素をクリックして検査し、選択します。
サイドバーでプロパティを編集し、サポートされている場所にドラッグして順序を変更するか、テキストをダブルクリックしてコピーを変更します。
UI で保留中の変更を確認します。
[確認] をクリックしてソース ファイルに適用します。
反応リライト [オプション] [ポート]
引数:
ポート開発サーバーのポートオーバーライド
オプション:
--no-open ブラウザを自動的に開きません
--host <ホスト> 開発サーバーのホスト (デフォルト: "localhost")
--verbose デバッグ ログを有効にする
ショートカット
ショートカット
アクション
Ctrl/Cmd + C
選択した要素をコピーする
Ctrl/Cmd + V
コピーした要素を兄弟として貼り付けます
Ctrl/Cmd + D
選択した要素を所定の位置に複製します
削除 / バックスペース
選択した要素を削除します
Ctrl/Cmd + Z
キャンバスの変更を元に戻す
Ctrl/Cmd + Shift + L
変更ログの切り替え
Ctrl/Cmd + クリック
オーバーレイを通じてリンクをたどる
テキストをダブルクリックします
テキストをインラインで編集する
注意事項
アプリのルート ディレクトリから act-rewrite を実行すると、フレームワークを検出してファイル パスを安全に解決できます。
これは開発ビルドに対してのみ機能し、運用ビルドに対しては機能しません。
現在のプロジェクト内のファイルのみが書き込みの対象となります。
このリポジトリ自体を操作するには:
pnpmインストール
pnpmビルド
pnpm テスト -- --run
反復的な CLI 開発の場合:
pnpm開発
ツールをエンドツーエンドでテストするには、サポートされている別の React アプリをローカルで実行する必要があります。
パッケージ/
cli/CLI、プロキシサーバー、およびソース変換
オーバーレイ/挿入されたブラウザー オーバーレイ
シャール

ed/共有 TypeScript タイプ
ライセンス
React アプリ用のビジュアル エディター (figma) は、UI 要素をライブで編集し、ソース ファイルを直接変更します。 API キーも AI もありません。
読み込み中にエラーが発生しました。このページをリロードしてください。
29
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

visual editor (figma) for your react apps, edit UI elements live and directly changes source files. no API key, no AI.. - donghaxkim/react-rewrite

GitHub - donghaxkim/react-rewrite: visual editor (figma) for your react apps, edit UI elements live and directly changes source files. no API key, no AI.. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
donghaxkim
/
react-rewrite
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
269 Commits 269 Commits .changeset .changeset .github/ workflows .github/ workflows packages packages .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
react-rewrite lets you edit a React app visually while it is running locally, then automatically writes those changes back to the source files in your project.
It is built for local development and works by opening a proxy in front of your dev server and injecting an overlay into the page.
full demo:
https://x.com/imdonghakim/status/2038230475894899119
You do not need to download or clone this repo.
From the root of your React app:
npm install -D react-rewrite-cli
Start your dev server, then in a second terminal run:
npx react-rewrite
If you want to try it without installing first:
npx react-rewrite-cli@latest
What it does
Select an element and inspect its component name, file path, and line number
Edit supported Tailwind-based layout, spacing, size, typography, and color properties
Double-click text to edit it inline
Copy, paste, and duplicate elements
Stage multiple changes and apply them with Confirm
Undo in-progress canvas changes and review applied changes in the changelog
Supported app setups: Next.js, Vite, and Create React App
Tailwind CSS is recommended if you want to use the property editor. Text editing and some structural actions do not depend on Tailwind.
Run this in the root of the React app you want to edit:
npm install -D react-rewrite-cli
If you don't want to install it first, you can also run it directly with npx react-rewrite-cli@latest .
Start your React dev server as usual.
In a second terminal, from the same project root, run:
npx react-rewrite
If auto-detection does not pick the right port, pass it explicitly:
npx react-rewrite 3000
The tool opens a local proxy in your browser, shows the editing overlay, and writes confirmed changes back into files inside your project.
Click an element to inspect and select it.
Edit properties in the sidebar, drag to reorder where supported, or double-click text to change copy.
Review pending changes in the UI.
Click Confirm to apply them to your source files.
react-rewrite [options] [port]
Arguments:
port Dev server port override
Options:
--no-open Don't open browser automatically
--host <host> Dev server host (default: "localhost")
--verbose Enable debug logging
Shortcuts
Shortcut
Action
Ctrl/Cmd + C
Copy selected element
Ctrl/Cmd + V
Paste copied element as sibling
Ctrl/Cmd + D
Duplicate selected element in place
Delete / Backspace
Remove selected element
Ctrl/Cmd + Z
Undo canvas changes
Ctrl/Cmd + Shift + L
Toggle changelog
Ctrl/Cmd + Click
Follow links through the overlay
Double-click text
Edit text inline
Notes
Run react-rewrite from your app's root directory so it can detect the framework and safely resolve file paths.
It only works against development builds, not production builds.
Only files inside the current project are eligible for writes.
To work on this repository itself:
pnpm install
pnpm build
pnpm test -- --run
For iterative CLI development:
pnpm dev
You will still need a separate supported React app running locally to test the tool end to end.
packages/
cli/ CLI, proxy server, and source transforms
overlay/ Injected browser overlay
shared/ Shared TypeScript types
License
visual editor (figma) for your react apps, edit UI elements live and directly changes source files. no API key, no AI..
There was an error while loading. Please reload this page .
29
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
