---
source: "https://github.com/adnanakil/nobuzz/blob/main/README.md"
hn_url: "https://news.ycombinator.com/item?id=49388752"
title: "Claudette: Make Claude Stop Talking Like a BuzzFeed Article"
article_title: "nobuzz/README.md at main · adnanakil/nobuzz · GitHub"
image: "https://opengraph.githubassets.com/881a829d798842301cb0af973a8202bada9a779546ccdec6926fd77cc57c8ded/adnanakil/nobuzz"
author: "aakil"
captured_at: "2026-08-21T15:24:44Z"
capture_tool: "hn-digest"
hn_id: 49388752
score: 27
comments: 17
posted_at: "2026-08-21T14:31:52Z"
tags:
  - hacker-news
  - translated
---

# Claudette: Make Claude Stop Talking Like a BuzzFeed Article

- HN: [49388752](https://news.ycombinator.com/item?id=49388752)
- Source: [github.com](https://github.com/adnanakil/nobuzz/blob/main/README.md)
- Score: 27
- Comments: 17
- Posted: 2026-08-21T14:31:52Z

## Translation

タイトル: クローデット: クロードに BuzzFeed 記事のような話をやめさせる
記事タイトル: nobuzz/README.md at main · adnakil/nobuzz · GitHub
説明: クロードの回答を Gemini 経由でパイプして BuzzFeed の音声を削除するクロード コード スキル (/debuzz) - main · adnaakil/nobuzz の nobuzz/README.md

記事本文:
メインの nobuzz/README.md · adnanakil/nobuzz · GitHub
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
アドナナキル
/
ノブズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 54 行 (33 loc) · 3.5 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルをコピー raw ファイルをダウンロード アウトライン編集と raw アクション NoBuzz
明らかに、Anthropic が古い Buzzfee でクロードのみをトレーニングしたことは今では常識です。

d 記事 (90 年代のノスタルジーへの愛を説明)。そこで、クロードと私は、クロードの最後の応答を取得し、Gemini CLI を通じて実行して、ミレニアル世代のクリックベイトのような会話から通常の英語に変換するクロード コード スキル ( /debuzz ) を構築しました。私たちはそれを「クローデット」と呼ぶことを考えていますが、キャット・ウーは訴訟を起こすことで有名なので、絶対に「クローデット」とは呼びません。
Claudette は、Claude が 1 つの不治の病を抱えた偉大なエンジニアであることに対する私たちの解決策です。それは、自身のプル リクエストについて TED トークを配信しているかのように話します。なぜテストが不安定なのかを尋ねると、「負荷に耐える仮定」、番号付きの 3 つの啓示、そして 3 番目のテストがこれまでで最も有益であるという厳粛な注意が得られます。単なるバグというものは決してありません。キッカーは常に存在します。
このスキルは、いくら促してもこの問題を完全には解決できないことを受け入れ、代わりに正直なことを行います。応答を別のモデルであるジェミニ (昨年夏のジェミニを覚えていますか?) に渡します。その唯一の仕事は、人間のようにそれを言うことです。クローデットは、クロードに翻訳を「整理」させると、削除された音声が正確に再導入されるため、ジェミニの翻訳をそのまま印刷することをピンキーに約束しました。
ここが興味深いところです。再試行ロジックは、単にあると便利なものではなく、同期パイプライン全体の負荷に耐えるという前提です。 3 つのことが私に思い浮かびましたが、3 つ目はこれまでで最も有益です。 [...] そして 3 番目に、これがキッカーです。重複排除キーにはタイムスタンプが含まれています。これは、再試行が実際には重複排除されないことを意味します。
後 ( /debuzz 、同僚モード):
同期パイプラインの再試行ロジックには 3 つのバグがあります。 syncQueue.ts:142 は、ジョブを再度キューに入れるのではなく、ETIMEDOUT を飲み込みます。バックオフの上限は 2 秒ですが、モバイル ネットワークには短すぎます。重複排除キーにはタイムスタンプが含まれているため、再試行します。

決して重複排除されません。修正: キーからタイムスタンプを削除し、上限を 30 秒に上げて、タイムアウト エラーを再スローします。
git clone https://github.com/adnanakil/nobuzz
mkdir -p ~ /.claude/skills
cp -r nobuzz/debuzz ~ /.claude/skills/
要件:
Gemini CLI ( npm install -g @google/gemini-cli )、認証済み - Gemini を 1 回実行して /auth を使用するか、 GEMINI_API_KEY を設定します。
/debuzz [モード] [テキスト]
モード
観客
得られるもの
同僚 (デフォルト)
エンジニア
同じコンテンツ、すべてのファイル パスとコード ブロックはそのまま、芝居がきかない
マネージャー
技術系に近いマネージャー
何が起こったのか、なぜそれが重要なのか、次に何が起こるのか — 約 3 分の 1 の長さ、コードなし
ディレクター
幹部
3 ～ 5 つの文: 結果、影響、質問。 30秒間の注意力を想定
テキスト引数を指定しないと、クロードの以前の返信が翻訳されます。モードの後に​​テキストを貼り付けて、代わりにそれを翻訳します。また、「普通の英語で言ってください」のような自然なフレーズでもトリガーされます。
魔法はありません。 Claudette は以前の応答を一時ファイルに書き込み、それを gemini -p "<plain-English style interactions>" にパイプして、Gemini の出力をそのまま出力します。 Gemini エラー (通常は認証) の場合、実際のエラーが表示されます。クロードは、明確にラベル付けされたフォールバックとして独自の書き換えのみを提供します。これは、デブザーがブザーに自分自身をデバズするように静かに要求するため、負荷のかかる変換が行われることになるからです。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude Code skill (/debuzz) that pipes Claude's answers through Gemini to remove the BuzzFeed voice - nobuzz/README.md at main · adnanakil/nobuzz

nobuzz/README.md at main · adnanakil/nobuzz · GitHub
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
adnanakil
/
nobuzz
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 54 lines (33 loc) · 3.5 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions NoBuzz
Obviously it's common knowledge by now that Anthropic has solely trained claude on old Buzzfeed articles (explaining its love for 90s nostalgia). So Claude and I built a Claude Code skill ( /debuzz ) that takes Claude's last response and runs it through the Gemini CLI to translate it from talking like a millennial clickbait to regular English. We're thinking of calling it "Claudette" but Cat Wu is famously litigious so it's definitely absolutely not called "Claudette".
Claudette is our solution to Claude being a great engineer with one incurable condition: it talks like it's delivering a TED talk about its own pull request. You ask why a test is flaky and you get a "load-bearing assumption," three numbered revelations, and a solemn note that the third one is the most instructive yet . Nothing is ever just a bug. There is always a kicker.
This skill accepts that no amount of prompting fully cures this, and does the honest thing instead: it hands the response to a different model, Gemini (remember gemini from last summer?), whose only job is to say it like a human person. Claudette has pinky promised to print Gemini's translation verbatim, because letting Claude "tidy up" the translation reintroduces exactly the voice being removed.
Here's where it gets interesting: the retry logic isn't just a nice-to-have — it's - the - load-bearing - assumption - of - the - entire - sync - pipeline. Three things jumped out at me, and the third one is the most instructive yet. [...] And third — and this is the kicker — the dedupe key includes a timestamp, which means retries are never actually deduplicated.
After ( /debuzz , colleague mode):
The sync pipeline's retry logic has three bugs. syncQueue.ts:142 swallows ETIMEDOUT instead of re-queuing the job. The backoff caps at 2 seconds, which is too low for mobile networks. And the dedupe key includes a timestamp, so retries are never deduplicated. Fix: strip the timestamp from the key, raise the cap to 30 seconds, and re-throw the timeout error.
git clone https://github.com/adnanakil/nobuzz
mkdir -p ~ /.claude/skills
cp -r nobuzz/debuzz ~ /.claude/skills/
Requirements:
The Gemini CLI ( npm install -g @google/gemini-cli ), authenticated — run gemini once and use /auth , or set GEMINI_API_KEY .
/debuzz [mode] [text]
Mode
Audience
What you get
colleague (default)
An engineer
Same content, every file path and code block intact, zero theatrics
manager
A technical-adjacent manager
What happened, why it matters, what's next — about a third the length, no code
director
An executive
Three to five sentences: outcome, impact, ask. Assumes thirty seconds of attention
With no text argument it translates Claude's previous reply. Paste text after the mode to translate that instead. It also triggers on natural phrases like "say that in normal english."
No magic. Claudette writes its previous reply to a temp file, pipes it through gemini -p "<plain-English style instructions>" , and prints Gemini's output verbatim. If Gemini errors (usually auth), you see the actual error — Claude only offers its own rewrite as a clearly labeled fallback, because a debuzzer that quietly asks the buzzer to debuzz itself is how you end up with a load-bearing translation.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
