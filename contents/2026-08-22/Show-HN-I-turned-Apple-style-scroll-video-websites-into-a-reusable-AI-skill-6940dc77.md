---
source: "https://github.com/musoyangrigor/scroll-video-website-skill"
hn_url: "https://news.ycombinator.com/item?id=49401459"
title: "Show HN: I turned Apple-style scroll-video websites into a reusable AI skill"
article_title: "GitHub - musoyangrigor/scroll-video-website-skill: Scroll Video Website is a portable AI-agent skill that turns videos into smooth, scroll-controlled canvas experiences with optimized WebP frames, progressive loading, and responsive animation. · GitHub"
image: "https://opengraph.githubassets.com/3be3692efa0edaa12e71671a3ed696f8247ee2033bec7e969a2779dbce6f86fd/musoyangrigor/scroll-video-website-skill"
author: "MusoyanGrigor"
captured_at: "2026-08-22T17:12:23Z"
capture_tool: "hn-digest"
hn_id: 49401459
score: 1
comments: 0
posted_at: "2026-08-22T16:47:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I turned Apple-style scroll-video websites into a reusable AI skill

- HN: [49401459](https://news.ycombinator.com/item?id=49401459)
- Source: [github.com](https://github.com/musoyangrigor/scroll-video-website-skill)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T16:47:58Z

## Translation

タイトル: Show HN: Apple スタイルのスクロールビデオ Web サイトを再利用可能な AI スキルに変えました
記事タイトル: GitHub - musoyangrigor/scroll-video-website-skill: Scroll Video Website は、最適化された WebP フレーム、プログレッシブ読み込み、応答性の高いアニメーションを備えた、ビデオをスムーズなスクロール制御のキャンバス エクスペリエンスに変えるポータブル AI エージェント スキルです。 · GitHub
説明: Scroll Video Website は、最適化された WebP フレーム、プログレッシブ読み込み、応答性の高いアニメーションを備えた、ビデオをスムーズなスクロール制御のキャンバス エクスペリエンスに変えるポータブル AI エージェント スキルです。 - musoyangrigor/scroll-video-website-skill

記事本文:
GitHub - musoyangrigor/scroll-video-website-skill: Scroll Video Website は、最適化された WebP フレーム、プログレッシブ読み込み、応答性の高いアニメーションを備えた、ビデオをスムーズなスクロール制御のキャンバス エクスペリエンスに変えるポータブル AI エージェント スキルです。 · GitHub
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
ムソヤンリゴール
/
スクロールビデオウェブサイトスキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
16 コミット 16 コミット フォルダーとファイル
ある

ssets アセット スクロールビデオウェブサイト スクロールビデオウェブサイト ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのエージェント スキルで、あらゆるビデオをスムーズなスクロール制御の Web サイトに変えます。
npx スキル追加 musoyangrigor/scroll-video-website-skill --skill スクロールビデオウェブサイト
ライブデモ · インストール · 使用法 · 機能
デモでは、1 つのビデオから作成されたシンプルな製品 Web サイトを示します。ビデオは画面いっぱいに表示され、前後にスクロールするとフレーム内をスムーズに移動します。
次のコマンドを使用して、同じ種類の Web サイトを作成できます。
$scroll-video-website ./media/product-demo.mp4
このスキルはビデオを最適化されたフレームに変換し、応答性の高いスクロール制御されたページを構築します。
エージェントにスキル名とその後にビデオ パスを付けます。
$scroll-video-website ./media/product-film.mp4
必要に応じて、パスの後に視覚的な方向を追加します。
$scroll-video-website ./media/product-film.mp4 ダークな編集スタイルを使用
既存の生成されたフレーム シーケンスを対話的に最適化します。
$scroll-video-website の最適化
このコマンドには引数はかかりません。エージェントは、現在のプロジェクトで生成されたシーケンスを検出し、実際のアセットを測定し、代表的なフレームを試行エンコードして、フォーマット、フレーム数、品質、およびその他のサポートされている設定の個別の推定値を表示します。次に、それらを組み合わせた推定影響を示し、何かを書き換える前に確認を求めます。
ビデオを最適化された WebP フレーム シーケンスに変換します。
プログレッシブフレームロードにより両方向にスムーズにスクラブします。
隣接フレームのブレンディングを使用して、応答性の高いフルビューポート キャンバスをレンダリングします。
既存のプロジェクト スタックを保持し、モーションの削減設定をサポートします。
デフォルトで最小限のキャンバスのみのエクスペリエンスを構築するか、追加されたデザインの方向に従います。
ファイル固有のサイズ e を使用して既存のシーケンスを対話的に最適化します。

処理前の推定。
Scroll Video Website は、最適化された WebP フレーム、プログレッシブ読み込み、応答性の高いアニメーションを備えた、ビデオをスムーズなスクロール制御のキャンバス エクスペリエンスに変えるポータブル AI エージェント スキルです。
Readme MIT ライセンス アクティビティ スター
2 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Scroll Video Website is a portable AI-agent skill that turns videos into smooth, scroll-controlled canvas experiences with optimized WebP frames, progressive loading, and responsive animation. - musoyangrigor/scroll-video-website-skill

GitHub - musoyangrigor/scroll-video-website-skill: Scroll Video Website is a portable AI-agent skill that turns videos into smooth, scroll-controlled canvas experiences with optimized WebP frames, progressive loading, and responsive animation. · GitHub
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
musoyangrigor
/
scroll-video-website-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
assets assets scroll-video-website scroll-video-website LICENSE LICENSE README.md README.md View all files Repository files navigation
Turn any video into a smooth, scroll-controlled website with one Agent Skill.
npx skills add musoyangrigor/scroll-video-website-skill --skill scroll-video-website
Live Demo · Install · Usage · Features
The demo shows a simple product website made from a single video. The video fills the screen, and scrolling forward or backward smoothly moves through its frames.
You can create the same kind of website with:
$scroll-video-website ./media/product-demo.mp4
The skill converts the video into optimized frames and builds the responsive scroll-controlled page for you.
Give your agent the skill name followed by a video path:
$scroll-video-website ./media/product-film.mp4
Optionally add visual direction after the path:
$scroll-video-website ./media/product-film.mp4 use a dark editorial style
Optimize an existing generated frame sequence interactively:
$scroll-video-website optimize
The command takes no arguments. The agent discovers generated sequences in the current project, measures the real assets, trial-encodes representative frames, and shows separate estimates for format, frame count, quality, and other supported settings. It then shows their combined estimated impact and asks for confirmation before rewriting anything.
Converts video into an optimized WebP frame sequence.
Scrubs smoothly in both directions with progressive frame loading.
Renders a responsive, full-viewport canvas with adjacent-frame blending.
Preserves the existing project stack and supports reduced-motion preferences.
Builds a minimal canvas-only experience by default, or follows added design direction.
Interactively optimizes existing sequences with file-specific size estimates before processing.
Scroll Video Website is a portable AI-agent skill that turns videos into smooth, scroll-controlled canvas experiences with optimized WebP frames, progressive loading, and responsive animation.
Readme MIT license Activity Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
