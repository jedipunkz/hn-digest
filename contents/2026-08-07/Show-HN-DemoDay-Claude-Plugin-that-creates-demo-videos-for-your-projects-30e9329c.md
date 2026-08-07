---
source: "https://github.com/mrieck/demoday-claude-plugin"
hn_url: "https://news.ycombinator.com/item?id=49211759"
title: "Show HN: DemoDay – Claude Plugin that creates demo videos for your projects"
article_title: "GitHub - mrieck/demoday-claude-plugin: Claude Code plugin that creates demo videos for your projects. Claude is given tools to drive the browser/cli, make recordings, and uses FAL.ai, ElevenLabs, and remotion to build a demo video with scene transitions. · GitHub"
author: "mrieck"
captured_at: "2026-08-07T15:44:47Z"
capture_tool: "hn-digest"
hn_id: 49211759
score: 1
comments: 0
posted_at: "2026-08-07T15:16:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DemoDay – Claude Plugin that creates demo videos for your projects

- HN: [49211759](https://news.ycombinator.com/item?id=49211759)
- Source: [github.com](https://github.com/mrieck/demoday-claude-plugin)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T15:16:20Z

## Translation

タイトル: Show HN: DemoDay – プロジェクトのデモビデオを作成する Claude プラグイン
記事のタイトル: GitHub - mrieck/demoday-claude-plugin: プロジェクトのデモビデオを作成する Claude Code プラグイン。クロードには、ブラウザ/CLI を操作し、録画を作成し、FAL.ai、イレブンラボ、およびリモートを使用してシーン遷移のあるデモ ビデオを構築するためのツールが与えられています。 · GitHub
説明: プロジェクトのデモビデオを作成する Claude Code プラグイン。クロードには、ブラウザ/CLI を操作し、録画を作成し、FAL.ai、イレブンラボ、およびリモートを使用してシーン遷移のあるデモ ビデオを構築するためのツールが与えられています。 - GitHub - mrieck/demoday-claude-plugin: デモビデオを作成する Claude Code プラグイン
[切り捨てられた]
HN テキスト: Claude Max サブウーファーをもっと使用する必要がありました。そのため、プロジェクトや新機能用のデモ ビデオを Claude に作成してもらいました。 Claude はプロジェクトを調べ、スクリプトを作成し、アプリを実行してから、Fal API キーと Remotion を使用してすべてのシーンを構築します。シーンの見た目が悪かったり、追加/編集が必要な場合は、修正を依頼してください。編集には Remotion を使用します。これは個人には無料ですが、大規模な組織にはライセンスが必要です。カスタム音声に イレブン ラボを使用する場合も同じです (サブスクリプションが必要です)。それ以外の場合、クロードはファルからの汎用音声を使用します。特に、より洗練された製品デモを提供する他の製品にお金を払っている人にとっては、フィードバックをお待ちしています。

記事本文:
GitHub - mrieck/demoday-claude-plugin: プロジェクトのデモビデオを作成する Claude Code プラグイン。クロードには、ブラウザ/CLI を操作し、録画を作成し、FAL.ai、イレブンラボ、およびリモートを使用してシーン遷移のあるデモ ビデオを構築するためのツールが与えられています。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムリック
/
デモデークロードプラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .claude-pl

ugin .claude-plugin コマンド コマンド docs docs mcp mcp remotion-template remotion-template スクリプト スクリプト スキル スキル .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json HOW_IT_WORKS.md HOW_IT_WORKS.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたのソフトウェアには起動ビデオが必要です。クロードに撮らせてください。
DemoDay は、あらゆるものを変換する Claude Code プラグインです。
コードベースをナレーション付きの完成したデモビデオに変換します。クロードはあなたのリポジトリを読んで、
あなたと一緒にスクリプトを作成し、アプリを画面上に表示し、AI ナレーションでアプリをナレーションします —
イントロ、キャプション、トランジション、およびオプションを備えた洗練された MP4 を手渡します。
オンカメラのプレゼンター。コマンドは 1 つ、コーヒーの値段についてです。 macOSのみ。
/demoday:create-demo-video
🎬 実際に見てみる
overboard_github.mp4
別の Claude プラグイン (プロジェクト ダッシュボード) である Overboard のサンプル デモ ビデオ (ナレーション付き)
ElementalLabs を使用した船長のナレーションを含め、すべて DemoDay によって編集されました。
🖱️ クロードがアプリを操作します — 最初にカメラの外でリハーサルを行い、次に
決定論的なランナーはスムーズなカーソル移動と人間の操作でテイクを実行します。
タイピング中。デッドエアも手探りもありません。
🎙️ AI ナレーションと単語タイミングのキャプション — ナレーションが最初に生成され、
そのため、すべてのシーンが正確に音声に反映されます。
🧑‍💼 オプションのオンカメラ プレゼンター — イントロ、アウトロ用の AI プレゼンター
ナレーションと同じ声に合わせて口パクした断面図。
🗣️ カスタムキャラクターの声 — わかりやすい英語で声を説明します
(「風化したロブスター船の船長、濃い海岸アクセント」) とデモ全体
その中で語られています。
🎞️ スクリーンキャストではなく、実際の編集 – タイトルカード、トランジション、
ズームクリック、吹き出し、Bロール、オーディオダッキング、レンダリング
降格。
💻映画も

Mac 上のあらゆるもの - Web アプリ、ネイティブ macOS アプリ、
CLI/ターミナル プログラム (はい、Claude Code プラグインを内部からデモできます)
ライブクロードセッション）。
🔁 繰り返しのコストが低い — ナレーションを 1 行変更すると、その行だけが変更されます。
再生されました。費用がかかる前に、コストの見積もりが表示されます。
/plugin マーケットプレイスに mrieck/claude-plugins を追加
/プラグインインストールdemoday@production-mark
( production-mark は、でホストされているマーケットプレイスです
mrieck/claude-plugins ;デモデーは
プラグイン — このリポジトリ。)
ソフトウェアについてはこれで終わりです。 npm install や手動の依存関係セットアップは必要ありません。最初の
/demoday:create-demo-video を実行すると、クロードはマシンをチェックし、
不足しているもの (ノード パッケージ、キャプチャ ブラウザ、ffmpeg) をインストールする前に
撮影が始まります。
要件: macOS および Claude Code 。の
クロードがあなたの代わりにできないことは、以下の 2 つのステップだけです。
1. fal.ai キーを追加します (必須)
音声、プレゼンター、B ロールなど、すべての世代が貫かれています。
ファルアイ。次の方法でキーを macOS キーチェーンに保存します。
これを自分の端末で実行します（クロード経由ではないため、キーは決して必要ありません）
トランスクリプトに触れます):
security add-generic-password -s Demonday -a FAL_API_KEY -w
エコーせずにキーの入力を求めます。一般的な 30 秒のデモの料金
fal クレジットで 1 ～ 2 ドル。
キーチェーンの条件ではありませんか?キーが存在できる他のすべての場所がカバーされます
HOW_IT_WORKS.md 。
2. 画面の権限を付与します (デスクトップ キャプチャのみ)
ネイティブ Mac アプリを撮影するには、クロード コードを実行するターミナルを指定します。
(ターミナル、iTerm、VS Code…) システムの画面録画とアクセシビリティ
「設定」→「プライバシーとセキュリティ」を選択し、端末を再起動してください。 Webアプリの撮影や
ステージングされたターミナル ウィンドウでの CLI デモには権限はまったく必要ありません。
何かが足りない場合は、クロードがプリフライト中に正確な設定ペインを教えてくれます。
同じキーチェーンのワンライナー、別のアカウント

名前:
ELEVENLABS_API_KEY — カスタム キャラクターの音声のロックを解除します
イレブンラボ音声デザイン
(アフィリエイトリンク) 。有料の イレブンラボ プランが必要です (スターター以上): 無料枠では音声がブロックされます
API 経由で作成でき、有料プランには商用ライセンスも適用されます
とにかくビデオを公開したいと思うでしょう。
BRAVE_API_KEY — B ロール生成で参照画像を検索します。
画面の録画、リハーサルのスクリーンショット、完成した MP4 はローカルに残ります。のみ
世代が必要とするものをアップロードします: ナレーションテキスト、ナレーション音声、
プレゼンター画像と B ロール フレームは fal.ai に移動します — そして
ナレーション テキストを使用する場合は、イレブンラボに送信します。 APIキー
常に読み取られるだけで、決して書き込まれることはありません
（詳細）。
見せたいソフトウェアのリポジトリから:
/demoday:create-demo-video
クロードはコードベースを読み、撮影する価値のある 2 つまたは 3 つのフローを提案し、質問します。
目標、聴衆、長さ、プレゼンターのスタイルについて検討し、費用の見積もりを示します。
「はい」と答えた後でのみ、単一のフレームが生成されます。 10分ほど後
demo/ に MP4 があります。
興味深い部分 — 2 パスのリハーサル/パフォーマンスのデザイン、ナレーションが最初
ペーシング、demo.json マニフェスト、コンテンツ アドレス指定のキャッシュ、その他に関するすべて
API キーの解決 — HOW_IT_WORKS.md にあります。
レンダリングには Remotion を使用します。これは個人および小規模企業には無料ですが、
従業員数のしきい値を超える有料ライセンスが必要です - を参照してください。
remotion.dev/license 。有料で生成された音声
イレブンラボのプランには商用利用権が含まれています。
商用利用が可能なイレブンラボプランに加入するには、
ここをクリックしてください (アフィリエイトリンク) 。
DemoDay 自体は MIT ライセンスを取得しています。
プロジェクトのデモビデオを作成する Claude Code プラグイン。クロードには、ブラウザ/CLI を操作し、録画を作成し、FAL.ai、イレブンラボ、およびリモートを使用してシーン遷移のあるデモ ビデオを構築するためのツールが与えられています。
R

eadme MIT ライセンス アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code plugin that creates demo videos for your projects. Claude is given tools to drive the browser/cli, make recordings, and uses FAL.ai, ElevenLabs, and remotion to build a demo video with scene transitions. - GitHub - mrieck/demoday-claude-plugin: Claude Code plugin that creates demo videos
[truncated]

I needed to use more of my Claude Max sub - so now I have Claude make demo videos for my projects or new features. Claude looks at your project, writes a script, drives your app, then it uses your Fal API key and Remotion to build out all the scenes. If a scene looks bad or something needs to be added/edited, just ask for a revision. It uses Remotion for editing which is free for individuals, but larger orgs need a license. Same thing if you use Eleven Labs for custom voices (needs a subscription)- otherwise Claude will use a generic voice from Fal. Would love some feedback, especially for anyone that pays for other products that offer more polished product demos.

GitHub - mrieck/demoday-claude-plugin: Claude Code plugin that creates demo videos for your projects. Claude is given tools to drive the browser/cli, make recordings, and uses FAL.ai, ElevenLabs, and remotion to build a demo video with scene transitions. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mrieck
/
demoday-claude-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .claude-plugin .claude-plugin commands commands docs docs mcp mcp remotion-template remotion-template scripts scripts skills skills .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json HOW_IT_WORKS.md HOW_IT_WORKS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Your software deserves a launch video. Let Claude film it.
DemoDay is a Claude Code plugin that turns any
codebase into a finished, narrated demo video. Claude reads your repo, writes the
script with you, drives your app on screen , narrates it with an AI voiceover —
and hands you a polished MP4 with intro, captions, transitions and an optional
on-camera presenter. One command, about the cost of a coffee. macOS only.
/demoday:create-demo-video
🎬 See it in action
overboard_github.mp4
Sample demo video of Overboard, another Claude Plugin (project dashboard), narrated and
edited entirely by DemoDay, including the sea-captain voiceover using ElevenLabs.
🖱️ Claude drives your app — it rehearses off camera first, then a
deterministic runner performs the take with smooth cursor moves and human
typing. No dead air, no fumbling.
🎙️ AI voiceover with word-timed captions — narration is generated first,
so every scene lands exactly on the voice.
🧑‍💼 Optional on-camera presenter — an AI presenter for the intro, outro
and cutaways, lip-synced to the same voice as the narration.
🗣️ Custom character voices — describe a voice in plain English
( "weathered lobster-boat captain, thick coastal accent" ) and the whole demo
is narrated in it.
🎞️ Real editing, not a screencast — title cards, transitions,
zoom-to-click, callouts, b-roll and audio ducking, rendered with
Remotion .
💻 Films almost anything on a Mac — web apps, native macOS apps, and
CLI/terminal programs (yes, it can demo a Claude Code plugin from inside a
live claude session).
🔁 Cheap to iterate — change one line of narration and only that line is
regenerated. A cost estimate is shown before anything is spent.
/plugin marketplace add mrieck/claude-plugins
/plugin install demoday@productive-mark
( productive-mark is the marketplace, hosted in
mrieck/claude-plugins ; demoday is
the plugin — this repo.)
That's it for software — no npm install , no manual dependency setup. The first
time you run /demoday:create-demo-video , Claude checks your machine and
installs anything missing (Node packages, a capture browser, ffmpeg) before
filming starts.
Requirements: macOS and Claude Code . The
only thing Claude can't do for you is the two steps below.
1. Add your fal.ai key (required)
All generation — voice, presenter, b-roll — runs through
fal.ai . Store the key in your macOS Keychain by
running this in your own terminal (not through Claude, so the key never
touches a transcript):
security add-generic-password -s demoday -a FAL_API_KEY -w
It prompts for the key without echoing it. A typical 30-second demo costs
$1–2 in fal credits .
Not on Keychain terms? Every other place a key can live is covered in
HOW_IT_WORKS.md .
2. Grant screen permissions (desktop capture only)
To film a native Mac app , give the terminal you run Claude Code from
(Terminal, iTerm, VS Code…) Screen Recording and Accessibility in System
Settings → Privacy & Security, then restart the terminal. Filming web apps and
CLI demos in a staged terminal window needs no permissions at all — and if
anything's missing, Claude tells you the exact Settings pane during preflight.
Same Keychain one-liner, different account name:
ELEVENLABS_API_KEY — unlocks custom character voices via
ElevenLabs Voice Design
(affiliate link) . Needs a paid ElevenLabs plan (Starter and up): the free tier blocks voice
creation over the API, and a paid plan also carries the commercial license
you'd want for a published video anyway.
BRAVE_API_KEY — lets b-roll generation search reference images.
Screen recordings, rehearsal screenshots and the finished MP4 stay local. Only
what generation needs is uploaded: narration text, narration audio, the
presenter image and b-roll frames go to fal.ai — and
narration text to ElevenLabs if you use it. API keys
are only ever read, never written
( details ).
From the repo of the software you want to show off:
/demoday:create-demo-video
Claude will read the codebase, pitch you two or three flows worth filming, ask
about goal, audience, length and presenter style, and show you the cost estimate.
Only after you say yes does it generate a single frame. Ten-ish minutes later
there's an MP4 in demo/ .
The interesting bits — the two-pass rehearse/perform design, narration-first
pacing, the demo.json manifest, content-addressed caching, and everything about
API-key resolution — live in HOW_IT_WORKS.md .
Rendering uses Remotion, which is free for individuals and small companies but
needs a paid license above a headcount threshold — see
remotion.dev/license . Voices generated on a paid
ElevenLabs plan include commercial usage rights.
To subscribe to an ElevenLabs plan that allows commercial use,
click here (affiliate link) .
DemoDay itself is MIT licensed .
Claude Code plugin that creates demo videos for your projects. Claude is given tools to drive the browser/cli, make recordings, and uses FAL.ai, ElevenLabs, and remotion to build a demo video with scene transitions.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
