---
source: "https://github.com/MerqryLabs/ai-crawler-visibility"
hn_url: "https://news.ycombinator.com/item?id=48683149"
title: "I made a Claude Code skill to check if AI crawlers can read your site"
article_title: "GitHub - MerqryLabs/ai-crawler-visibility · GitHub"
author: "novaesystems"
captured_at: "2026-06-26T07:14:18Z"
capture_tool: "hn-digest"
hn_id: 48683149
score: 2
comments: 0
posted_at: "2026-06-26T06:42:49Z"
tags:
  - hacker-news
  - translated
---

# I made a Claude Code skill to check if AI crawlers can read your site

- HN: [48683149](https://news.ycombinator.com/item?id=48683149)
- Source: [github.com](https://github.com/MerqryLabs/ai-crawler-visibility)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T06:42:49Z

## Translation

タイトル: AI クローラーがサイトを読み取れるかどうかを確認するクロード コード スキルを作成しました
記事のタイトル: GitHub - MerqryLabs/ai-crawler-visibility · GitHub
説明: GitHub でアカウントを作成して、MerqryLabs/ai-crawler-visibility の開発に貢献します。

記事本文:
GitHub - MerqryLabs/ai-crawler-visibility · GitHub
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
メルクリラボ
/
AI クローラーの可視性
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
MerqryLabs/ai-crawler-visibility
本支店

es タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット README.md README.md SKILL.md SKILL.md ai-crawler-visibility.zip ai-crawler-visibility.zip すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Web サイトが実際に AI 検索に表示されるかどうか、また表示されない場合は正確に修正する方法を示す Claude Code スキル。
現在、ChatGPT Search、Claude、Perplexity は人々を Web サイトに誘導しています。しかし、それらにフィードを供給するクローラー (ClaudeBot、GPTBot、PerplexityBot、Bytespider、Meta-ExternalAgent) は JavaScript を実行しません。サーバーから返される生の HTML のみを読み取ります。
Google のレンダラーは何年にもわたって JavaScript を処理してきたため、クライアントでレンダリングされたサイト (React、Vite、Vue、Svelte、その他の SPA) は Google で完全に上位にランクされ、AI の回答ではまったく表示されなくなります。
「Google トラフィックは得られるが、AI からは何も得られない」と考えたことがある場合、または洗練された雰囲気でコーディングされた SPA を出荷したのに、ChatGPT のどこにもいない場合、ほとんどの場合、これが理由です。
このスキルは、非 JS クローラーと同じ方法 (生の HTML、ブラウザーなし) でページを取得し、実際のコンテンツと主要なシグナル ( title 、メタ ディスクリプション、正規、Open Graph、JSON-LD、クロール可能なリンク) が実際に存在するかどうかをページごとにレポートします。
次に、具体的なフレームワーク固有の修正を規定します。 「SSR を考慮する」のではなく、正確なツール、設定変更、それが機能することを確認する 1 行の方法です。
すべてのページは次の 3 つの判定のいずれかを取得します。
を診断して処方します。マーケティングコピーを書いたり、キーワード調査をしたり、バックリンクを追いかけたりすることはありません。
これはクロード・コード・エージェント・スキルです。フォルダーをスキル ディレクトリにドロップすると、Claude Code が自動的にそれを取得します。
個人用 (すべてのプロジェクトで利用可能):
git clone https://github.com/ < merqrylabs > /ai-crawler-visibility.git \
~ /.claude/skills/ai-crawler-visibility
広報

オブジェクトスコープ (このリポジトリのみ):
git clone https://github.com/ < merqrylabs > /ai-crawler-visibility.git \
./.claude/skills/ai-crawler-visibility
それだけです。 npm インストール、API キー、サードパーティ パッケージは必要ありません。アナライザーはプレーンな Python 3 で実行されます。
設定が異なる場合は、公式のエージェント スキル ドキュメントで正規のスキル パスを確認してください。
クロード・コードに自然言語で話しかけるだけです。このスキルは、次のような種類の質問をすると自動的にトリガーされます。
「私のサイトが ChatGPT に表示されないのはなぜですか?」
「私の新しいサイトは AI 検索エンジンに表示されますか?」
「JSON-LD を追加しましたが、リッチリザルトがまだ表示されません。」
バンドルされたアナライザーを直接実行することもできます。
# 単一ページ
python3 scripts/check_visibility.py https://thesite.com
# いくつかの特定のページ
python3 scripts/check_visibility.py https://thesite.com https://thesite.com/pricing
# サイトマップからのサンプル (--limit ページまでチェック)
python3 scripts/check_visibility.py --sitemap https://thesite.com/sitemap.xml --limit 10
結果をプログラムで処理するには --json を追加します。
⚠️ localhost ではなく、実稼働 URL を常に確認してください。ローカル開発サーバーは、デプロイされたサイトがクローラーにどのように提供されるかを反映しません。
# AI クローラー可視性レポート
## 評決
https://thesite.com — 見えない
https://thesite.com/pricing — 見えない
## クローラーが実際に見るもの
/ : 表示可能な単語が 0 件あります。生の HTML に欠落しています: メタ記述、
正規、Open Graph、JSON-LD、クロール可能なリンク。
## 修正 (最も大きな影響が最初にあります)
[フレームワーク固有の具体的な手順 — それぞれの検証方法を含む]
## 修正を確認する方法
ページのソースを表示して (検査ではなく)、見出しテキストを確認します
および JSON-LD は生の HTML に表示されます。
目に見えないまたは部分的な評決を得ましたか?ここが追い越し車線です。
このスキルは、何が壊れているか、何が修正されているかを示します。手動で実装したくない場合

事前レンダリングと静的 JSON-LD インジェクションを自分で行う場合は、コンパニオン AI Crawler Fixer パックが代わりに実行します。
静的 SEO + JSON-LD を直接index.html に挿入します。
実際のコンテンツが生の HTML に表示されるように、ビルド後のプリレンダリングを設定します。
承認ゲートを通過するため、サインオフがなければコードベースには何も触れません
修正が完了したことを証明するために診断を再実行します
👉 **[AI Crawler Fixer パックを入手→] https://novae8.gumroad.com/l/AICrawlerFixer
診断は無料で、今後も無料です。 Fixer は、今日中に問題を解決したい場合に使用します。
仕組み (短いバージョン)
ページを非表示にするのはユーザー エージェントではありません。JavaScript を実行しないことが非表示にします。 (スクリプトは引き続き、サーバー側の事前レンダリングを検出するための 2 番目のリクエストとして実際のボット UA を送信します。)
JavaScript によって挿入された JSON-LD は、非レンダリング クローラーには見えません。生の HTML 内にある必要があります。
水分補給は大丈夫です。すでにレンダリングされた HTML にハンドラーをアタッチするので、コンテンツはすでにそこにあります。純粋なクライアント側レンダリング (CSR) が問題です。
[ページ ソースの表示] には、生の HTML (クローラーが認識するもの) が表示されます。 「Inspect」では、レンダリングされた DOM (表示されないもの) が表示されます。必ず「ページのソースを表示」で確認してください。
MIT — 使って、フォークして、出荷してください。
Novae Systems によって構築されました。地元企業向けの Web サイト、AI 自動化、SEO です。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
AI クローラーの可視性
最新の
2026 年 6 月 26 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to MerqryLabs/ai-crawler-visibility development by creating an account on GitHub.

GitHub - MerqryLabs/ai-crawler-visibility · GitHub
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
MerqryLabs
/
ai-crawler-visibility
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
MerqryLabs/ai-crawler-visibility
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits README.md README.md SKILL.md SKILL.md ai-crawler-visibility.zip ai-crawler-visibility.zip View all files Repository files navigation
A Claude Code skill that tells you whether your website is actually visible to AI search — and exactly how to fix it if it isn't.
ChatGPT Search, Claude, and Perplexity are sending people to websites now. But the crawlers that feed them — ClaudeBot, GPTBot, PerplexityBot, Bytespider, Meta-ExternalAgent — do not run JavaScript . They read only the raw HTML your server returns.
Google's renderer has handled JavaScript for years, so a client-rendered site (React, Vite, Vue, Svelte, any SPA) can rank perfectly fine on Google and still be completely invisible in AI answers .
If you've ever thought "we get Google traffic but nothing from AI" — or you shipped a slick vibe-coded SPA and you're nowhere in ChatGPT — this is almost always why.
This skill fetches your page(s) the same way a non-JS crawler does — raw HTML, no browser — and reports, per page, whether your real content and key signals ( title , meta description, canonical, Open Graph, JSON-LD, crawlable links) are actually present.
Then it prescribes a concrete, framework-specific fix . Not "consider SSR" — the exact tool, config change, and one-line way to verify it worked.
Every page gets one of three verdicts:
It diagnoses and prescribes . It does not write marketing copy, do keyword research, or chase backlinks.
This is a Claude Code Agent Skill . Drop the folder into your skills directory and Claude Code picks it up automatically.
Personal (available in every project):
git clone https://github.com/ < merqrylabs > /ai-crawler-visibility.git \
~ /.claude/skills/ai-crawler-visibility
Project-scoped (just this repo):
git clone https://github.com/ < merqrylabs > /ai-crawler-visibility.git \
./.claude/skills/ai-crawler-visibility
That's it. No npm install, no API keys, no third-party packages — the analyzer runs on plain Python 3.
Check the official Agent Skills docs for the canonical skills path if your setup differs.
Just talk to Claude Code in natural language. The skill triggers on its own when you ask the kind of question it's built for:
"Why isn't my site showing up in ChatGPT?"
"Is my new site visible to AI search engines?"
"I added JSON-LD but rich results still aren't showing."
You can also run the bundled analyzer directly:
# Single page
python3 scripts/check_visibility.py https://thesite.com
# Several specific pages
python3 scripts/check_visibility.py https://thesite.com https://thesite.com/pricing
# Sample from a sitemap (checks up to --limit pages)
python3 scripts/check_visibility.py --sitemap https://thesite.com/sitemap.xml --limit 10
Add --json to process results programmatically.
⚠️ Always check your production URL, not localhost — local dev servers don't reflect how the deployed site is served to crawlers.
# AI Crawler Visibility Report
## Verdict
https://thesite.com — INVISIBLE
https://thesite.com/pricing — INVISIBLE
## What the crawler actually sees
/ : 0 visible words. Missing from raw HTML: meta description,
canonical, Open Graph, JSON-LD, crawlable links.
## Fixes (highest impact first)
[Concrete, framework-specific steps — each with how to verify]
## How to confirm the fix
View Page Source (not Inspect) and confirm your headline text
and JSON-LD appear in the raw HTML.
Got an INVISIBLE or PARTIAL verdict? Here's the fast lane.
This skill tells you what's broken and what the fix is . If you'd rather not hand-implement prerendering and static JSON-LD injection yourself, the companion AI Crawler Fixer pack does it for you:
Injects static SEO + JSON-LD straight into index.html
Sets up post-build prerendering so real content lands in the raw HTML
Works through approval gates so nothing touches your codebase without your sign-off
Re-runs the diagnosis after to prove the fix landed
👉 **[Get the AI Crawler Fixer pack →] https://novae8.gumroad.com/l/AICrawlerFixer
Diagnosis is free and always will be. The Fixer is for when you want the problem gone today.
How it works (the short version)
The User-Agent isn't what makes a page invisible — not executing JavaScript is. (The script still sends a real bot UA as a second request to detect server-side prerendering.)
JSON-LD injected by JavaScript is invisible to non-rendering crawlers. It has to be in the raw HTML.
Hydration is fine. It attaches handlers to already-rendered HTML, so the content's already there. Pure client-side rendering (CSR) is the problem.
"View Page Source" shows raw HTML (what the crawler sees). "Inspect" shows the rendered DOM (what it does not see). Always verify with View Page Source.
MIT — use it, fork it, ship it.
Built by Novae Systems — websites, AI automation, and SEO for local businesses.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
ai-crawler-visibility
Latest
Jun 26, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
