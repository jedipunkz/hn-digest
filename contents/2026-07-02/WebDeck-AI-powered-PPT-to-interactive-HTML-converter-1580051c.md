---
source: "https://github.com/lzytttttt/WebDeck"
hn_url: "https://news.ycombinator.com/item?id=48755374"
title: "WebDeck – AI-powered PPT to interactive HTML converter"
article_title: "GitHub - lzytttttt/WebDeck · GitHub"
author: "lzytttttt"
captured_at: "2026-07-02T01:39:32Z"
capture_tool: "hn-digest"
hn_id: 48755374
score: 1
comments: 0
posted_at: "2026-07-02T01:38:01Z"
tags:
  - hacker-news
  - translated
---

# WebDeck – AI-powered PPT to interactive HTML converter

- HN: [48755374](https://news.ycombinator.com/item?id=48755374)
- Source: [github.com](https://github.com/lzytttttt/WebDeck)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T01:38:01Z

## Translation

タイトル: WebDeck – AI を活用した PPT からインタラクティブな HTML コンバーター
記事タイトル: GitHub - lzytttttt/WebDeck · GitHub
説明: GitHub でアカウントを作成して、lzytttttt/WebDeck の開発に貢献します。

記事本文:
GitHub - lzytttttt/WebDeck · GitHub
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
ジットットット
/
ウェブデッキ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
3 コミット 3 コミット アプリ アプリコンポーネント コンポーネント lib lib スクリプト スクリプトタイプ タイプ .env.local.example .env.local.example .gitignore .gitignore README.md README.md README_EN.md README_EN.md next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json patch-out.json patch-out.json postcss.config.mjs postcss.config.mjs tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json tsconfig.tsbuildinfo tsconfig.tsbuildinfo すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PPT → HTML 変換レイヤー — PowerPoint をインタラクティブな Web プレゼンテーションに変換します
Web Deck は、ワンクリックで従来の PowerPoint (.pptx) プレゼンテーションをインタラクティブな Web ページ プレゼンテーションに変換するオープン ソース ツールです。
想像してみてください。慎重に作成した PPT を共有するたびに、相手が PowerPoint をインストールしているかどうか、レイアウトがずれていないか、携帯電話で表示できるかどうかを心配する必要があります。 Web Deck はこれらすべてを解決します。PPT をアップロードすると、AI がコンテンツ構造を自動的に分析し、ブラウザーで直接開いたり、共有したり、表示したりできる Web ページ ドキュメントを生成します。
PowerPoint は 1987 年に誕生しました。その中心的な前提は次のとおりです。プレゼンテーション = 直線的に回転するスライド。この仮定はプロジェクターの時代には合理的でしたが、今日では—
1987年の世界 2026年の世界
┌─────────┐ ┌─────────┐
│ プロジェクター │ │ 携帯電話・タブレット・PC │
│ フルスクリーンソフトウェア │ → │ WeChat・飛書・メール │
│ リニアなページめくり │ │ レスポンシブ・インタラクティブ │
│ オフラインファイル交換 │ │ リンク共有・コラボレーション │
━━━━━━━┘ ━━━━━━━┘
PPT の問題点
根本原因
ウェブデッキソリューション
開くにはPowerPoint/WPSが必要です
プライベートバイナリ形式
ブラウザネイティブ、依存関係なし
携帯電話での写植

混乱
固定ピクセルレイアウト
レスポンシブデザイン、3端子対応
送信する共有ファイル（数十MB）
ファイル配布モデル
1 つのリンク、数キロバイトの HTML
コンテンツは「死んでいる」
リニアなページめくり、インタラクションなし
FAQ アコーディオン、カード、タイムライン、その他のネイティブ インタラクション
バージョン管理の混乱
バイナリ差分は不可能です
JSON 構造化データ、Git フレンドリー
コンテンツの検索・引用はできません
視覚的な形式、セマンティクスなし
セマンティック HTML、インデックス可能および参照可能
PPT を置き換えるのではなく、パラダイムをアップグレードする
Web Deck は「別のプレゼンテーション ツール」としてではなく、トランジション レイヤーとして位置付けられています。
従来のワークフロー：Word作成→PPT組版→ファイル送信→相手がPPTで開く
↓ ウェブデッキ
新しいワークフロー: PPT 作成 → Web Deck 変換 → リンク共有 → ブラウザを開く
何も学び直す必要はありません。 PPT のスキルは完全に保持されます。Web Deck は、車輪の再発明ではなく、PPT の上に HTML 機能を重ね合わせます。
.pptx ファイルをアップロードすると、Web Deck が各ページのタイトル、テキスト、キーポイント、メモなどの構造化情報を自動的に抽出し、解析結果が一目でわかるインポート品質レポートを生成します。
Anthropic Claude の機能を活用することで、Web Deck は単に「スクリーンショットを Web ページに変換」するだけでなく、コンテンツのセマンティクスを理解し、PPT コンテンツをネイティブ Web コンポーネントにインテリジェントに再編成します。
保守的モード (Conservative): PPT の 1 ページ = 1 章、元の構造を維持、学習コストはゼロ
拡張モード: AI がコンテンツを自由に再編成し、ヒーロー、カード、タイムライン、比較などのネイティブ Web コンポーネントを生成します。
Web Deck は 12 種類のネイティブ ブロックをサポートし、ほぼすべてのプレゼンテーション シナリオをカバーします。
実際のプレゼンテーション ソフトウェアと同様に、矢印キーを押してページをめくり、Esc を押して終了し、N を押してスピーカー ノートを切り替えます。フェード / スライド / ズームの 3 つの切り替えアニメーションをサポートします。
リンク共有: 公開後に専用リンクを取得し、Open Graph メタデータをサポートします
静的 HTML エクスポート: 完全に自己完結型の .html ファイルをエクスポートし、オフラインでも読み取り可能です。

外部依存関係
API キーがありませんか?それは問題ではありません。 Web Deck には MockAIProvider が組み込まれているため、外部サービスなしですべての機能を完全に体験できます。
従来の PPT 変換ツール (LibreOffice エクスポート PDF/HTML など) は基本的に、ピクセル レベルの転送、つまりスライドの外観を別の形式に「スクリーンショット」して転送します。 Web Deck はまったく異なるアプローチを採用しています。
PPT の各ページには意味情報が含まれています。これは表紙、これは目次、これはデータ比較、そしてこれは FAQ です。 Web Deck は AI を通じてこれらのセマンティクスを理解し、PPT の視覚的な外観を Web ページ上に「撮影」するのではなく、Web ページの最も適切なネイティブ コンポーネントを使用してセマンティクスを再表現します。
従来の変換: PPT ビジュアル → スクリーンショット/ベクター → Web ページに埋め込み (まだ「機能していない」)
Web デッキ: PPT コンテンツ → AI セマンティック理解 → ネイティブ Web コンポーネント (ライブ、インタラクティブ)
デュアルプロバイダーアーキテクチャ: スタックすることはありません
ユーザーがアップロードしたPPT
│
▼
[AIProvider.generateWebDeck()]
│
§── Anthropic API は利用可能ですか?
│ └── はい → クロードの意味理解 → 構造化 JSON
│
└── いいえ (ネットワークの問題/キーなし/タイムアウト)
└── 自動フォールバック → MockAIProvider → 決定論的ローカル生成
ユーザーには「変換に失敗しました」というメッセージが表示されることはありません。オフライン モードでは、モック エンジンはヒューリスティック ルールに基づいて使用可能なデッキを生成します。AI ほどスマートではありませんが、完全な使いやすさを保証します。
エディターでのすべてのデザイン調整 (テーマの切り替え、丸い角、影、間隔) は、CSS カスタム プロパティを通じて行われます。
:ルート{
--deck-bg : # ffffff ;
--デッキテキスト: # 1a1a2e ;
--デッキアクセント: # 6366f1 ;
--デッキ半径: 12 ピクセル;
--deck-shadow : 0 4 px 6 px -1 px rgba ( 0 , 0 , 0 , 0.1 );
}
プレビューとエクスポートは同じ変数セットを共有します。つまり、エディターで見られる効果 = エクスポートされた HTML ファイルの効果 = 最終的な共有リンクの効果です。見たものはそのまま得られます。
手書きSVGチャートレンダラー(棒グラフ、折れ線グラフ、円グラフ、ドーナツグラフ、KPIカード、データテーブル)、

Chart.js / D3 / ECharts などのサードパーティ ライブラリに依存しません。色は、color-mix() を介してテーマの CSS 変数から自動的に導出されます。テーマを切り替えると、チャートの色が自動的に変わります。
git クローン https://github.com/lzytttttt/WebDeck.git
cd ウェブデッキ
npmインストール
構成 (オプション)
cp .env.local.example .env.local
.env.local を編集し、Anthropic API キーを入力します (入力されていない場合は、オフライン モック モードを使用します)。
ANTHROPIC_API_KEY = sk-ant-...
ANTHROPIC_MODEL = クロード-3-5-sonnet-20240620
始める
npm rundev
ブラウザを開いて http://localhost:3000 にアクセスして開始してください。
npm 実行シード:デモ
これにより、会社概要、製品発売、資金調達ロードショー、年次概要、トレーニング コースという 5 つの組み込みデモ プロジェクトが生成されます。
┌───────────────────────┐
│ ウェブデッキ │
§─────┬─────┬─────┬───────┤
│ PPTのアップロード │ AI変換 │ ビジュアライゼーション │ 共有＆エクスポート │
│ 解析エンジン │ コンテンツ再構築 │ エディタ │ │
§─────┴─────┴───┴───────┤
│ ウェブデッキデータレイヤー │
│ (12ブロックタイプ・テーマシステム・アニメーションシステム・チャートレンダリング) │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ Next.js 14 · React 18 · TypeScript · Tailwind CSS │
│ 人間クロード · JSZip · fast-xml-parser │
━━━━━━━━━━━━━━━━━━━━━━━━

──────────┘
コアデータフロー
.pptx ファイル
│
▼
[JSZip解凍] → [fast-xml-parser XML抽出] → [extractSlideText]
│
▼
ParsedSlide[] (タイトル、rawText、箇条書き、メモ、imageRefCount、tableRefCount)
│
§──→ [buildImportQualityReport] → 6 つの警告
│
▼
[AIProvider.generateWebDeck()]
§── AnthropicAIProvider: Claude API → JSON Schema 検証
━── MockAIProvider: ヒューリスティックルール → 決定論的生成
│
▼
WebDeck JSON (ID、タイトル、テーマ、セクション[]、提案[])
│
§──→ [DeckRenderer] → Reactコンポーネントツリー
§──→ [発表者] → 全画面プレゼンテーション
└──→ [exportStaticHtml] → 自己完結型 .html ファイル
コアモジュール
モジュール
パス
説明
PPTX分析
lib/pptx/
JSZip に基づいて .pptx を解凍し、テキストと構造を XML で抽出します
AIエンジン
lib/ai/
Anthropic Claude インターフェイス + モックオフラインモード
デッキエンジン
ライブラリ/デッキ/
テーマ、アニメーション、チャートレンダリング、ブロックファクトリー、QA
輸出エンジン
lib/エクスポート/
依存関係のない静的 HTML を生成する
ストレージ層
ライブラリ/ストレージ/
プロジェクトの永続化 (JSON ファイル ストレージ)
編集者
コンポーネント/エディタ/
WYSIWYG エディター + インスペクター パネル
デッキコンポーネント
コンポーネント/デッキ/
12種類のブロックReactレンダラー
国際化
lib/i18n/
中英対訳辞書
📂 プロジェクトの構造
ウェブデッキ/
§── app/ # Next.js App Router
│ §── page.tsx # ホームページ
│ §── プロジェクト/ # プロジェクト管理
│ │ §── new/ # PPT をアップロードする
│ │ └─ [id]/
│ │ §── edit/ # エディター
│ │ └── present/ # 全画面表示
│ §── デモ/ # デモギャラリー
│ §── share/[shareId]/ # 公開共有ページ
│ └── api/ # API ルーティング
§── コンポーネント/
│ §── デッキ/ # デッキブロックレンダラー
│ §── editor/ # エディターコンポーネント
│ §── レイアウト/ # レイアウトコンポーネント
│ ━─ う

i/ # 共通の UI コンポーネント
§── lib/
│ §── ai/ # AI プロバイダー (Anthropic + Mock)
│ §── デッキ/ # デッキコアロジック
│ §── export/ # 静的 HTML エクスポート
│ §── i18n/ # 国際化
│ §── pptx/ # PPTX 解析
│ §── storage/ # データストレージ
│ └── utils/ # ユーティリティ関数
§── types/ # TypeScript の型定義
§── scripts/ # ツールスクリプト
└── data/ # プロジェクトデータ (git は無視されます)
⚠️既知の制限事項
限界
説明
計画
PPTから画像を抽出しないでください
現時点では画像の参照数のみがカウントされており、バイナリデータは抽出されません。
ロードマップ v0.3
PPT アニメーション/トランジションを保存しないでください
PPTの入場アニメーションと切り替えエフェクトはマッピングできません
ロードマップ v0.4
SmartArt はサポートされていません
SmartArt は PPT に特有の複雑な要素です
まだ計画はありません
AI の結果には手動によるレビューが必要
AI はコンテンツの意味を誤って判断する可能性があります。変換後に確認することをお勧めします。
継続的な最適化プロンプト
🗺️ロードマップ
v0.2 — 画像の抽出とインライン化 (PPTX から画像を抽出して HTML に埋め込む)
v0.3 — 追加の AI プロバイダー (OpenAI、ローカル モデル、Ollama)
インジケーター
代表値
PPTX 分析 (10 ページ)
< 1 秒
AI 変革 (10 ページ、保守的モード)
5 ～ 15 秒 (API 応答に応じて)
モック変換（10ページ）
< 500ミリ秒
エクスポートHTMLファイルのサイズ
10～50KB（画像を除く）
最初の画面読み込み
< 2 秒
📜 オープンソースプロトコル
このプロジェクトは CC BY-NC-SA 4.0 (表示 - 非営利 - 継承) に基づいてライセンスされています。
🔄 Share Alike — 派生作品は同じ CC BY-NC-SA 4.0 ライセンスに基づいてライセンスを取得する必要があります
機能ブランチを作成します: git checkout -b feature/amazing-feature
変更をコミット: git commit -m '素晴らしい機能を追加'
Push ブランチ: git Push Origin feature/amazing-feature
Web デッキ — すべてのプレゼンテーションを最新の方法で可視化します ✨
web-deck-pied.vercel.app
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0
最新の
J

2026 年 1 月 1 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to lzytttttt/WebDeck development by creating an account on GitHub.

GitHub - lzytttttt/WebDeck · GitHub
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
lzytttttt
/
WebDeck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits app app components components lib lib scripts scripts types types .env.local.example .env.local.example .gitignore .gitignore README.md README.md README_EN.md README_EN.md next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json patch-out.json patch-out.json postcss.config.mjs postcss.config.mjs tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json tsconfig.tsbuildinfo tsconfig.tsbuildinfo View all files Repository files navigation
PPT → HTML 转换层 — 把你的 PowerPoint 变成交互式网页演示文稿
Web Deck 是一个将传统 PowerPoint（ .pptx ）演示文稿一键转换为 交互式网页演示文稿 的开源工具。
想象一下：你有一份精心准备的 PPT，但每次分享都要担心对方有没有装 PowerPoint、排版会不会错位、手机上能不能看。Web Deck 解决了这一切——上传 PPT，AI 自动解析内容结构，生成一份 可以直接在浏览器中打开、分享、演示 的网页文档。
PowerPoint 诞生于 1987 年，它的核心假设是： 演示 = 线性翻页的幻灯片 。这个假设在投影仪时代是合理的，但在今天——
1987 年的世界 2026 年的世界
┌──────────────┐ ┌──────────────┐
│ 一台投影仪 │ │ 手机·平板·PC │
│ 一个全屏软件 │ → │ 微信·飞书·邮件 │
│ 线性翻页 │ │ 响应式·交互 │
│ 离线文件交换 │ │ 链接分享·协作 │
└──────────────┘ └──────────────┘
PPT 的痛点
根因
Web Deck 的解法
打开需要 PowerPoint/WPS
私有二进制格式
浏览器原生，零依赖
手机上排版错乱
固定像素布局
响应式设计，三端适配
分享要发文件（几十 MB）
文件分发模式
一个链接，几 KB HTML
内容是"死"的
线性翻页，无交互
FAQ 手风琴、卡片、时间线等原生交互
版本管理混乱
二进制 diff 不可能
JSON 结构化数据，Git 友好
内容无法被搜索/引用
视觉格式，无语义
语义化 HTML，可索引可引用
不是替代 PPT，而是升级范式
Web Deck 的定位不是"又一个演示工具"，而是一个 过渡层 ：
传统工作流： Word 写稿 → PPT 排版 → 发文件 → 对方用 PPT 打开
↓ Web Deck
新工作流： PPT 写稿 → Web Deck 转换 → 链接分享 → 浏览器打开
你不需要重新学任何东西。 你的 PPT 技能完全保留 ——Web Deck 在 PPT 的基础上叠加了 HTML 的能力，而不是推倒重来。
上传 .pptx 文件，Web Deck 会自动提取每一页的标题、正文、要点、备注等结构化信息，生成导入质量报告，让你对解析结果一目了然。
借助 Anthropic Claude 的能力，Web Deck 不是简单地"截图转网页"，而是 理解你的内容语义 ，将 PPT 内容智能重组为原生网页组件：
保守模式（Conservative） ：一页 PPT = 一个章节，保留原始结构，零学习成本
增强模式（Enhanced） ：AI 自由重组内容，生成 Hero、Cards、Timeline、Comparison 等网页原生组件
Web Deck 支持 12 种原生区块 ，覆盖几乎所有演示场景：
按方向键翻页， Esc 退出， N 键切换演讲者备注——就像真正的演示软件一样。支持 Fade / Slide / Zoom 三种切换动画。
链接分享 ：发布后获得专属链接，支持 Open Graph 元数据
静态 HTML 导出 ：导出一个完全自包含的 .html 文件，离线可读，无任何外部依赖
没有 API Key？没关系。Web Deck 内置 MockAIProvider ，无需任何外部服务即可完整体验所有功能。
传统 PPT 转换工具（如 LibreOffice 导出 PDF/HTML）本质上是在做 像素级搬运 ——把幻灯片的视觉外观"截图"到另一个格式。Web Deck 的思路完全不同：
PPT 的每一页都包含语义信息：这是封面、这是目录、这是数据对比、这是 FAQ。Web Deck 通过 AI 理解这些语义，然后用 最合适的网页原生组件 重新表达——而不是把 PPT 的视觉外观"拍照"到网页上。
传统转换： PPT 视觉 → 截图/矢量图 → 嵌入网页（还是"死"的）
Web Deck： PPT 内容 → AI 语义理解 → 原生网页组件（活的、可交互的）
双 Provider 架构：永不卡死
用户上传 PPT
│
▼
[AIProvider.generateWebDeck()]
│
├── Anthropic API 可用？
│ └── Yes → Claude 语义理解 → 结构化 JSON
│
└── No（网络问题/无 Key/超时）
└── 自动回退 → MockAIProvider → 确定性本地生成
用户永远不会看到"转换失败"。 离线模式下，Mock 引擎基于启发式规则生成可用的 deck——虽然不如 AI 智能，但保证了完整的可用性。
编辑器中的每一个设计调整（主题切换、圆角、阴影、间距）都通过 CSS 自定义属性驱动：
: root {
--deck-bg : # ffffff ;
--deck-text : # 1a1a2e ;
--deck-accent : # 6366f1 ;
--deck-radius : 12 px ;
--deck-shadow : 0 4 px 6 px -1 px rgba ( 0 , 0 , 0 , 0.1 );
}
预览和导出共用同一套变量 ——编辑器里看到的效果 = 导出的 HTML 文件的效果 = 最终分享链接的效果。所见即所得。
手写 SVG 图表渲染器（柱状图、折线图、饼图、环形图、KPI 卡片、数据表格），不依赖 Chart.js / D3 / ECharts 等第三方库。颜色通过 color-mix() 从主题 CSS 变量自动派生——切换主题时图表颜色自动跟随变化。
git clone https://github.com/lzytttttt/WebDeck.git
cd WebDeck
npm install
配置（可选）
cp .env.local.example .env.local
编辑 .env.local ，填入你的 Anthropic API Key（不填则使用离线 Mock 模式）：
ANTHROPIC_API_KEY = sk-ant-...
ANTHROPIC_MODEL = claude-3-5-sonnet-20240620
启动
npm run dev
打开浏览器访问 http://localhost:3000 ，开始使用！
npm run seed:demo
这会生成 5 个内置 Demo 项目：公司简介、产品发布、融资路演、年度总结、培训课程。
┌─────────────────────────────────────────────────────┐
│ Web Deck │
├──────────┬──────────┬──────────┬────────────────────┤
│ 上传 PPT │ AI 转换 │ 可视化 │ 分享 & 导出 │
│ 解析引擎 │ 内容重构 │ 编辑器 │ │
├──────────┴──────────┴──────────┴────────────────────┤
│ Web Deck 数据层 │
│ (12 种区块类型 · 主题系统 · 动画系统 · 图表渲染) │
├─────────────────────────────────────────────────────┤
│ Next.js 14 · React 18 · TypeScript · Tailwind CSS │
│ Anthropic Claude · JSZip · fast-xml-parser │
└─────────────────────────────────────────────────────┘
核心数据流
.pptx 文件
│
▼
[JSZip 解压] → [fast-xml-parser 提取 XML] → [extractSlideText]
│
▼
ParsedSlide[] (title, rawText, bullets, notes, imageRefCount, tableRefCount)
│
├──→ [buildImportQualityReport] → 6 种警告
│
▼
[AIProvider.generateWebDeck()]
├── AnthropicAIProvider: Claude API → JSON Schema 校验
└── MockAIProvider: 启发式规则 → 确定性生成
│
▼
WebDeck JSON (id, title, theme, sections[], suggestions[])
│
├──→ [DeckRenderer] → React 组件树
├──→ [Presenter] → 全屏演示
└──→ [exportStaticHtml] → 自包含 .html 文件
核心模块
模块
路径
说明
PPTX 解析
lib/pptx/
基于 JSZip 解压 .pptx ，提取 XML 中的文本和结构
AI 引擎
lib/ai/
Anthropic Claude 接口 + Mock 离线模式
Deck 引擎
lib/deck/
主题、动画、图表渲染、区块工厂、质量检查
导出引擎
lib/export/
生成零依赖静态 HTML
存储层
lib/storage/
项目持久化（JSON 文件存储）
编辑器
components/editor/
所见即所得编辑器 + Inspector 面板
Deck 组件
components/deck/
12 种区块的 React 渲染器
国际化
lib/i18n/
中英文双语字典
📂 项目结构
WebDeck/
├── app/ # Next.js App Router
│ ├── page.tsx # 首页
│ ├── projects/ # 项目管理
│ │ ├── new/ # 上传 PPT
│ │ └── [id]/
│ │ ├── edit/ # 编辑器
│ │ └── present/ # 全屏演示
│ ├── demo/ # Demo 画廊
│ ├── share/[shareId]/ # 公开分享页
│ └── api/ # API 路由
├── components/
│ ├── deck/ # Deck 区块渲染器
│ ├── editor/ # 编辑器组件
│ ├── layout/ # 布局组件
│ └── ui/ # 通用 UI 组件
├── lib/
│ ├── ai/ # AI 提供者（Anthropic + Mock）
│ ├── deck/ # Deck 核心逻辑
│ ├── export/ # 静态 HTML 导出
│ ├── i18n/ # 国际化
│ ├── pptx/ # PPTX 解析
│ ├── storage/ # 数据存储
│ └── utils/ # 工具函数
├── types/ # TypeScript 类型定义
├── scripts/ # 工具脚本
└── data/ # 项目数据（git ignored）
⚠️ 已知限制
限制
说明
计划
不提取 PPT 中的图片
当前只统计图片引用数，不提取二进制数据
Roadmap v0.3
不保留 PPT 动画/过渡
PPT 的入场动画、切换效果无法映射
Roadmap v0.4
不支持 SmartArt
SmartArt 是 PPT 特有的复杂元素
暂无计划
AI 结果需人工审核
AI 可能误判内容语义，建议转换后检查
持续优化 Prompt
🗺️ Roadmap
v0.2 — 图片提取与内联（从 PPTX 中提取图片嵌入 HTML）
v0.3 — 更多 AI Provider（OpenAI、本地模型、Ollama）
指标
典型值
PPTX 解析（10 页）
< 1 秒
AI 转换（10 页，保守模式）
5-15 秒（取决于 API 响应）
Mock 转换（10 页）
< 500 毫秒
导出 HTML 文件大小
10-50 KB（不含图片）
首屏加载
< 2 秒
📜 开源协议
本项目采用 CC BY-NC-SA 4.0 （署名-非商业性使用-相同方式共享）协议。
🔄 相同方式共享 — 衍生作品必须采用相同的 CC BY-NC-SA 4.0 协议
创建特性分支： git checkout -b feature/amazing-feature
提交更改： git commit -m 'Add amazing feature'
推送分支： git push origin feature/amazing-feature
Web Deck — 让每一份展示都能以最现代的方式被看见 ✨
web-deck-pied.vercel.app
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0
Latest
Jul 1, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
