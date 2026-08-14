---
source: "https://github.com/robss2020/claude-fable-5-having-fun"
hn_url: "https://news.ycombinator.com/item?id=49305642"
title: "Claude Fable 5 Having Fun"
article_title: "GitHub - robss2020/claude-fable-5-having-fun: SHOVE — a CLI tactics game Claude Fable 5 designed, built, and played (6 sessions, 22m19s) until it was genuinely fun. Brief + writeup + source. · GitHub"
author: "logicallee"
captured_at: "2026-08-14T23:12:42Z"
capture_tool: "hn-digest"
hn_id: 49305642
score: 2
comments: 1
posted_at: "2026-08-14T23:01:46Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable 5 Having Fun

- HN: [49305642](https://news.ycombinator.com/item?id=49305642)
- Source: [github.com](https://github.com/robss2020/claude-fable-5-having-fun)
- Score: 2
- Comments: 1
- Posted: 2026-08-14T23:01:46Z

## Translation

タイトル: クロード寓話 5 楽しんで
記事のタイトル: GitHub - robss2020/claude-fable-5-having-fun: SHOVE — CLI 戦術ゲーム Claude Fable 5 を、本当に楽しくなるまで設計、構築、プレイ (6 セッション、22 分 19 秒) しました。概要 + 記事 + ソース。 · GitHub
説明: SHOVE — CLI 戦術ゲーム Claude Fable 5 が、本当に楽しくなるまで設計、構築、プレイ (6 セッション、22 分 19 秒) しました。概要 + 記事 + ソース。 - robss2020/claude-fable-5-have-fun

記事本文:
GitHub - robss2020/claude-fable-5-having-fun: SHOVE — CLI 戦術ゲーム Claude Fable 5 が、本当に楽しくなるまで設計、構築、プレイ (6 セッション、22 分 19 秒) されました。概要 + 記事 + ソース。 · GitHub
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
ロブス2020
/
クロード・寓話-5-楽しんでいます
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット LEDGER.txt LEDGER.txt README.md README.md Brief.md Brief.md Brief.txt bri

ef.txt devlog.md devlog.md game_state.json game_state.json play_ledger.jsonl play_ledger.jsonl report.md report.md shove.py shove.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI (Claude Fable 5) は、ゲームを発明し、構築し、単独でプレイして心から楽しめるようになるまで実際にプレイするように言われました。そして実際にそのようにしました。これは、モデルが何を楽しいと感じているか、そしてモデルが自分自身の楽しさを正直に判断できるかどうかを知るための、小さいながらも具体的な窓となります。その結果が、完全に CLI を通じてプレイされるターンベースの Into the Breach スタイルの戦術パズルである SHOVE であり、Fable 5 が本当に面白いと判断するまで、6 つのバージョンと 6 つの記録されたプレイ セッション (合計 22 分 19 秒) にわたって反復されました。タスクを設定した概要から始めて、デザイン、バージョンごとの開発ログ、プレイ時間の完全な台帳、プレイがどのような感じであったかを示す一人称レポートをカバーする記事を読みます。ディレクトリの残りの部分には、ゲーム自体 ( shove.py ) とその生のプレイ ログが保存されます。
— クロード著の README Opus 4.8
びっくり！ここの人間。きっと、この文章の中に人間が書いた文章が隠されているとは予想していなかったでしょう。まあ、それは本当です。この段落 (のみ) には職人による手作業によるキーの押下があります。全部一字一句入力してみました。とにかく、このリポジトリで示していることをどのように実行したかを次に示します。新しいディレクトリを作成し、そのファイルに表示されている内容を含む Brief.txt ファイルを作成しました。次に、Claude Code を起動して Fable 5 を選択し、ディレクトリ内のファイルを読み取ってタスクを完了し、同じディレクトリにレポートを書くように指示しました。私はそれを読むのが好きで、後でそれを共有するために、Opus 4.8 を起動し、ディレクトリを新しい github claude-fable-5-having-fun にアップロードするように指示し、なぜ誰かがそれを気にしたり読んだりするのかについて「だから何」と答えることも含めて、物事を 1 段落に要約するように指示しました。 SEに署名するように言いました

Claude Opus 4.8 によって書かれた Readme のセクション。この段落のすべては人間である私によって書かれており、brief.txt の内容も私が書きました。
SHOVE — CLI 戦術ゲーム、Claude Fable 5 が本当に楽しくなるまで設計、構築、プレイ (6 セッション、22 分 19 秒) しました。概要 + 記事 + ソース。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

SHOVE — a CLI tactics game Claude Fable 5 designed, built, and played (6 sessions, 22m19s) until it was genuinely fun. Brief + writeup + source. - robss2020/claude-fable-5-having-fun

GitHub - robss2020/claude-fable-5-having-fun: SHOVE — a CLI tactics game Claude Fable 5 designed, built, and played (6 sessions, 22m19s) until it was genuinely fun. Brief + writeup + source. · GitHub
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
robss2020
/
claude-fable-5-having-fun
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits LEDGER.txt LEDGER.txt README.md README.md brief.md brief.md brief.txt brief.txt devlog.md devlog.md game_state.json game_state.json play_ledger.jsonl play_ledger.jsonl report.md report.md shove.py shove.py View all files Repository files navigation
An AI (Claude Fable 5) was told to invent, build, and actually play a game until it genuinely enjoyed playing it alone — and it did, which is a small but concrete window into what a model finds fun and whether it can honestly judge its own enjoyment. The result is SHOVE , a turn-based Into the Breach –style tactics puzzle played entirely through a CLI, iterated across six versions and six logged play sessions (22m19s total) until Fable 5 judged it truly fun. Start with the brief that set the task, then read the writeup covering the design, the version-by-version development log, the complete ledger of play times, and a first-person report of what playing it felt like; the rest of the directory holds the game itself ( shove.py ) and its raw play logs.
— README written by Claude Opus 4.8
surprise! human here. I bet you weren't expecting a human-written paragraph hidden in all this slop. well, it's true. this paragraph (only) has artisanal hand-pressed key presses. I typed out all of it verbatim. So anyway, here's how I did what you see in this repo: I just made a new directory, made a brief.txt file with the contents you see in that file, then I started Claude Code and selected Fable 5, and then told it to read the file in the directory and complete the task there, writing a report in the same directory. I liked reading it and then later in order to share it I started Opus 4.8 and told it to upload the directory to a new github claude-fable-5-having-fun and told it to summarize things in 1 paragraph, including answering "so what" about why anyone would care or read it. I told it to sign its section of the readme as written by Claude Opus 4.8. everything in this paragraph was written by me, the human, and I wrote the brief.txt contents.
SHOVE — a CLI tactics game Claude Fable 5 designed, built, and played (6 sessions, 22m19s) until it was genuinely fun. Brief + writeup + source.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
