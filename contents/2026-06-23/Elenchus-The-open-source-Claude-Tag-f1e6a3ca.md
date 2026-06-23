---
source: "https://github.com/Kheil-Z/elenchus"
hn_url: "https://news.ycombinator.com/item?id=48650130"
title: "Elenchus: The open-source Claude Tag"
article_title: "GitHub - Kheil-Z/elenchus: Elenchus: multiplayer LLM workspace. Stop copy-pasting AI replies. One thread — your team, your AI, everyone in the same room, each using your own LLM API key. Think \"Claude Team\" without the hassle. · GitHub"
author: "AdilZtn"
captured_at: "2026-06-23T19:35:15Z"
capture_tool: "hn-digest"
hn_id: 48650130
score: 1
comments: 0
posted_at: "2026-06-23T19:30:06Z"
tags:
  - hacker-news
  - translated
---

# Elenchus: The open-source Claude Tag

- HN: [48650130](https://news.ycombinator.com/item?id=48650130)
- Source: [github.com](https://github.com/Kheil-Z/elenchus)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T19:30:06Z

## Translation

タイトル: Elenchus: オープンソースの Claude タグ
記事タイトル: GitHub - Kheil-Z/elenchus: Elenchus: マルチプレイヤー LLM ワークスペース。 AI の返信をコピペするのはやめましょう。 1 つのスレッド - チーム、AI、同じ部屋にいる全員が独自の LLM API キーを使用します。手間をかけずに「クロードチーム」を考えてください。 · GitHub
説明: Elenchus: マルチプレイヤー LLM ワークスペース。 AI の返信をコピペするのはやめましょう。 1 つのスレッド - チーム、AI、同じ部屋にいる全員が独自の LLM API キーを使用します。手間をかけずに「クロードチーム」を考えてください。 - ケイルZ/エレンクス

記事本文:
GitHub - Kheil-Z/elenchus: Elenchus: マルチプレイヤー LLM ワークスペース。 AI の返信をコピペするのはやめましょう。 1 つのスレッド - チーム、AI、同じ部屋にいる全員が独自の LLM API キーを使用します。手間をかけずに「クロードチーム」を考えてください。 · GitHub
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
Kh

エイルZ
/
エレンクス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
43 コミット 43 コミット .github/ workflows .github/ workflows app app コンポーネント コンポーネント docs docs lib lib public public supabase/ migrations supabase/ migrations .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .prettierrc .prettierrc ライセンス ライセンスPRIVACY.md PRIVACY.md README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
マルチプレイヤー AI ワークスペース。複数の人が 1 つの会話スレッドを共有し、それぞれが独自の API キーを持ち込んでいます。 Claude、Gemini、または ChatGPT を @mention — タグ付けした人が応答すると、発信者に請求されます。
ライブデモ — Supabase の無料枠でホストされているため、新しいアカウントの作成にはレート制限がかかる場合があります。
マルチプレイヤー会話 — 共有スレッド、リアルタイム同期、各人に帰属するメッセージ
ユーザーごとの BYOK — 各メンバーは独自の API キーを追加します。あなたは決して他人のためにお金を払うことはありません
マルチプロバイダー — Anthropic (Claude)、Google (Gemini)、OpenAI (ChatGPT)。各ユーザーが自分のものを選択します
ビジョン + ドキュメント — 画像と PDF を添付します。画像はビジョン入力として送信され、PDF は抽出されるか、ネイティブにクロードに送信されます。
プロジェクト ドキュメント — プロジェクト全体または 1 つの会話に範囲指定されたファイルをアップロードします。 AIがいつそれらを認識するかを制御する
カスタム システム プロンプト — AI に対するプロジェクトごとの指示
アクティビティ フィード — 未読インジケーターを使用して、外出中に何が起こったかを把握します
リアルタイムのマルチプレイヤー チャット — 2 人のユーザーが同じ会議に参加

バージョンの 1 つがクロードを @メンションしています。各ユーザーは独自の API キーを持ち、異なるプロバイダーを持つことができます。トークンのコストは、呼び出しをトリガーした人に帰属します。プロバイダー インジケーターが入力の下部に表示されます。
プロジェクトとドキュメント — プロジェクトは、会話、メンバー、アップロードされたファイル、アクティビティ フィードをグループ化します。ドキュメントの範囲を 1 つの会話に限定することも、プロジェクト全体で共有することもできます。チャット内で、ファイルをいつ AI に送信するかを制御します (常にまたは送信しない)。トークンの支出はユーザーごとに分類されるため、誰が何を使用しているのかを誰もが確認できます。
設定 — 各ユーザーは独自の表示名、アバターの色、API キーを設定します。キーは保存時に暗号化されます。他のメンバーに影響を与えることなく、いつでもキーを取り消したり、プロバイダーを切り替えたりできます。
Supabase アカウント (無料枠で動作します)
Vercel アカウント (無料枠で動作) — または任意の Node 互換ホスト
git clone https://github.com/Kheil-Z/elenchus.git
CD エレンクス
pnpmインストール
2. Supabase プロジェクトを作成する
supabase.com → 新しいプロジェクトに移動します
プロビジョニングが完了したら、「SQL エディター」→「新規クエリー」を開きます。
supabase/migrations/000_schema.sql の内容全体を貼り付けて実行します。
「認証」→「プロバイダー」→「電子メール」に移動し、有効になっていることを確認します
「認証」→「電子メールテンプレート」に移動し、確認メールが Supabase を参照しないように送信者名を更新します。
cp .env.example .env.local
変数
どこで見つけられますか
NEXT_PUBLIC_SUPABASE_URL
Supabase → プロジェクト設定 → API → プロジェクト URL
NEXT_PUBLIC_SUPABASE_ANON_KEY
Supabase → プロジェクト設定 → API → anon public
SUPABASE_SERVICE_ROLE_KEY
Supabase → プロジェクト設定 → API → service_role シークレット
暗号化キー
ノード -e "console.log(require('crypto').randomBytes(32).toString('hex'))" で生成します
4. ローカルで実行する
pnpm開発
http://localhost:3000 を開きます。
vercel.com にインポートします。N が自動検出されます。

ext.js
Vercel のプロジェクト設定のステップ 3 の 4 つの環境変数を追加します。
Supabase → Authentication → URL Configuration で、Site URL を Vercel URL に設定し、それをリダイレクト URL (例: https://your-app.vercel.app/** ) に追加します。これがないと、電子メール確認リンクは機能しません。
エレンクス/
§── アプリ/
│ §── (保護)/ # 認証済みページ (プロジェクト、チャット、設定)
│ §── api/ # サーバー側 API ルート (LLM 呼び出し、アップロードなど)
│ §── auth/ # ログイン/サインアップページ
│ └── プライバシー/ # プライバシーポリシー（公開）
§── コンポーネント/ # 共有 UI コンポーネント
§── lib/ # タイプ、Supabase クライアント、LLM 層、ユーティリティ
└── スーパーベース/
└── migrations/ # 単一のスキーマ ファイル — 新しいプロジェクトで 1 回実行
技術スタック
レイヤー
選択
フレームワーク
Next.js (アプリルーター)
言語
TypeScript (厳密)
スタイリング
テイルウィンド CSS v4
データベース + 認証
スーパーベース (Postgres + RLS)
ストレージ
スーパーベースストレージ
AIプロバイダー
Anthropic、Google Gemini、OpenAI
ホスティング
ヴェルセル
プライバシー
PRIVACY.md またはアプリ内プライバシー ポリシーを参照してください。
Elenchus: マルチプレイヤー LLM ワークスペース。 AI の返信をコピペするのはやめましょう。 1 つのスレッド - チーム、AI、同じ部屋にいる全員が独自の LLM API キーを使用します。手間をかけずに「クロードチーム」を考えてください。
エレンクス-ブラッシュ.vercel.app
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Elenchus: multiplayer LLM workspace. Stop copy-pasting AI replies. One thread — your team, your AI, everyone in the same room, each using your own LLM API key. Think "Claude Team" without the hassle. - Kheil-Z/elenchus

GitHub - Kheil-Z/elenchus: Elenchus: multiplayer LLM workspace. Stop copy-pasting AI replies. One thread — your team, your AI, everyone in the same room, each using your own LLM API key. Think "Claude Team" without the hassle. · GitHub
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
Kheil-Z
/
elenchus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .github/ workflows .github/ workflows app app components components docs docs lib lib public public supabase/ migrations supabase/ migrations .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .prettierrc .prettierrc LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
Multiplayer AI workspace. Multiple people share one conversation thread, each bringing their own API key. @mention Claude, Gemini, or ChatGPT — whoever you tag responds, billed to the person who called them.
Live demo — hosted on Supabase free tier, so new account creation may be rate-limited.
Multiplayer conversations — shared thread, real-time sync, messages attributed to each person
BYOK per user — each member adds their own API key; you never pay for anyone else
Multi-provider — Anthropic (Claude), Google (Gemini), OpenAI (ChatGPT); each user picks their own
Vision + documents — attach images and PDFs; images sent as vision inputs, PDFs extracted or sent natively to Claude
Project documents — upload files project-wide or scoped to a single conversation; control when the AI sees them
Custom system prompts — per-project instructions for the AI
Activity feed — catch up on what happened while you were away, with unread indicators
Real-time multiplayer chat — two users in the same conversation, one @mentions Claude. Each user has their own API key and can have a different provider; the token cost is attributed to whoever triggered the call. The provider indicator is visible at the bottom of the input.
Projects and documents — a project groups conversations, members, uploaded files, and an activity feed. Documents can be scoped to a single conversation or shared project-wide. Inside a chat, you control when files are sent to the AI (Always or Never). Token spend is broken down per user so everyone can see who is using what.
Settings — each user sets their own display name, avatar colour, and API key. Keys are encrypted at rest. You can revoke a key or switch providers at any time without affecting other members.
A Supabase account (free tier works)
A Vercel account (free tier works) — or any Node-compatible host
git clone https://github.com/Kheil-Z/elenchus.git
cd elenchus
pnpm install
2. Create a Supabase project
Go to supabase.com → New project
Once provisioned, open SQL Editor → New query
Paste and run the entire contents of supabase/migrations/000_schema.sql
Go to Authentication → Providers → Email and confirm it is enabled
Go to Authentication → Email Templates and update the sender name so confirmation emails don't reference Supabase
cp .env.example .env.local
Variable
Where to find it
NEXT_PUBLIC_SUPABASE_URL
Supabase → Project Settings → API → Project URL
NEXT_PUBLIC_SUPABASE_ANON_KEY
Supabase → Project Settings → API → anon public
SUPABASE_SERVICE_ROLE_KEY
Supabase → Project Settings → API → service_role secret
ENCRYPTION_KEY
Generate with node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
4. Run locally
pnpm dev
Open http://localhost:3000 .
Import it at vercel.com — it auto-detects Next.js
Add the four environment variables from step 3 in Vercel's project settings
In Supabase → Authentication → URL Configuration , set Site URL to your Vercel URL and add it to Redirect URLs (e.g. https://your-app.vercel.app/** ) — without this, email confirmation links will not work
elenchus/
├── app/
│ ├── (protected)/ # Authenticated pages (projects, chat, settings)
│ ├── api/ # Server-side API routes (LLM calls, uploads, etc.)
│ ├── auth/ # Login / signup pages
│ └── privacy/ # Privacy policy (public)
├── components/ # Shared UI components
├── lib/ # Types, Supabase client, LLM layer, utilities
└── supabase/
└── migrations/ # Single schema file — run once on a fresh project
Tech stack
Layer
Choice
Framework
Next.js (App Router)
Language
TypeScript (strict)
Styling
Tailwind CSS v4
Database + auth
Supabase (Postgres + RLS)
Storage
Supabase Storage
AI providers
Anthropic, Google Gemini, OpenAI
Hosting
Vercel
Privacy
See PRIVACY.md or the in-app privacy policy .
Elenchus: multiplayer LLM workspace. Stop copy-pasting AI replies. One thread — your team, your AI, everyone in the same room, each using your own LLM API key. Think "Claude Team" without the hassle.
elenchus-blush.vercel.app
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
