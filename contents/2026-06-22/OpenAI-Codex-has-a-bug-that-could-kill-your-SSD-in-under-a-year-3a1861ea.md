---
source: "https://www.notebookcheck.net/OpenAI-Codex-has-a-bug-that-could-kill-your-SSD-in-under-a-year.1326191.0.html"
hn_url: "https://news.ycombinator.com/item?id=48634658"
title: "OpenAI Codex has a bug that could kill your SSD in under a year"
article_title: "OpenAI Codex has a bug that could kill your SSD in under a year - Notebookcheck News"
author: "abixb"
captured_at: "2026-06-22T20:03:04Z"
capture_tool: "hn-digest"
hn_id: 48634658
score: 4
comments: 1
posted_at: "2026-06-22T19:11:18Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Codex has a bug that could kill your SSD in under a year

- HN: [48634658](https://news.ycombinator.com/item?id=48634658)
- Source: [www.notebookcheck.net](https://www.notebookcheck.net/OpenAI-Codex-has-a-bug-that-could-kill-your-SSD-in-under-a-year.1326191.0.html)
- Score: 4
- Comments: 1
- Posted: 2026-06-22T19:11:18Z

## Translation

タイトル: OpenAI Codex には 1 年以内に SSD を停止させる可能性のあるバグがあります
記事タイトル: OpenAI Codex には 1 年以内に SSD を破壊する可能性のあるバグがある - Notebookcheck ニュース
説明: OpenAI の Codex CLI が密かに SSD を破壊しています。ログ シンクの構成が正しくないと、ローカル データベースに年間最大 640 TB の書き込みが発生します。これは、一般的なドライブの耐用年数をはるかに超えています。このバグはまだ GitHub で公開されています。

記事本文:
OpenAI Codexには1年以内にSSDを停止させる可能性のあるバグがある - Notebookcheck News
▼
OpenAI Codex には 1 年以内に SSD を停止させる可能性のあるバグがあります
ⓘ Zac Wolff on Unsplash この OpenAI Codex のバグは、チェックしないままにしておくと、1 年も経たないうちにドライブの保証されている耐久性がすべて燃え尽きる可能性があります。 OpenAI の Codex CLI は密かに SSD を破壊しています。ログ シンクの構成が正しくないと、ローカル データベースに年間最大 640 TB の書き込みが発生します。これは、一般的なドライブの耐用年数をはるかに超えています。このバグはまだ GitHub で公開されています。アヌバブ・シャルマ、2026 年 6 月 22 日公開 🇩🇪 AI ストレージ
4 レビュー ← 選択したタイプを除外 ← 選択したタグを除外
OpenAI の Codex CLI を使用し、長時間実行したままにすると、SSD が損傷を受ける可能性があります。
1996fanrui という名前の GitHub ユーザーは、マシン上のディスク アクティビティが異常に高いことに気づき、6 月 14 日にこの問題を文書化しました。調査した結果、Codex がローカル SQLite データベース (~/.codex/logs_2.sqlite に保存) に継続的に診断ログの書き込みを行っていることがわかりました。 21 日間の稼働時間で、ドライブは約 37 TB の書き込みを吸収しました。年間換算すると、年間約 640 テラバイトになります。一般的な 1 TB のコンシューマ SSD の寿命は約 600 TBW です。したがって、このバグを放置しておくと、1 年以内にドライブの保証耐久性がすべて使い果たされてしまう可能性があります。
原因は、おそらく誰もエンド ユーザーに配布することを意図していないログ構成です。 Codex の SQLite フィードバック シンクは、デフォルトでグローバル TRACE レベル (可能な限り最もノイズの多い設定) で実行されます。生の WebSocket ペイロードから、「passwd」や「ld.so.cache」を開くなどのありふれたファイルシステム イベントまで、あらゆるものをログに記録します。また、標準の RUST_LOG 環境変数も無視されるため、これを無効にする明確な方法はありません。ログに記録されたデータの約 71% は TRACE レベルです。

少なくとも平均的なユーザーにとっては、実際の診断目的はありません。
さらに悪化させるのは書き込み増幅です。データベースは単に拡大するだけでなく、1 分あたり数万回の挿入と削除の操作を繰り返しています。ファイル サイズが示すよりもはるかに多くの量を物理的にドライブに書き込んでいます。
これは実際には、少なくとも 4 月からさまざまな形で既知の問題となっており、年間を通じて複数の関連レポートが提出されています。 OpenAI の最近の変更ログには SQLite の信頼性に関する修正がいくつか含まれていますが、書き込み速度の問題には対処していません。この問題はまだ未解決のままだ。
それまでの間、Linux および macOS ユーザーは、「~/.codex/logs_2.sqlite」を「/tmp/」にシンボリックリンクして、書き込みを RAM にリダイレクトできます。ファイルには会話データが含まれていないため、再起動時に失われても問題ありません。
Unsplash の Zac Wolff による注目の画像
静的バージョン 動的ロード コメントの読み込み中
この記事にコメントする
関連記事

## Original Extract

OpenAI's Codex CLI is quietly destroying SSDs. A misconfigured logging sink writes up to 640 TB/year to a local database, which is way more than a typical drive's lifetime endurance. The bug is still open on GitHub.

OpenAI Codex has a bug that could kill your SSD in under a year - Notebookcheck News
▼
OpenAI Codex has a bug that could kill your SSD in under a year
ⓘ Zac Wolff on Unsplash This OpenAI Codex bug, if left unchecked, could burn through your drive's entire warranted endurance in less than a year. OpenAI's Codex CLI is quietly destroying SSDs. A misconfigured logging sink writes up to 640 TB/year to a local database, which is way more than a typical drive's lifetime endurance. The bug is still open on GitHub. Anubhav Sharma , Published 06/22/2026 🇩🇪 AI Storage
4 Reviews ← exclude selected types ← exclude selected tags
If you use OpenAI's Codex CLI and leave it running for long periods of time, your SSD may be getting hammered.
A GitHub user named 1996fanrui documented the issue on June 14 after noticing unusually high disk activity on their machine. After digging around, they found Codex was continuously hammering a local SQLite database (stored at ~/.codex/logs_2.sqlite) with diagnostic log writes. Over 21 days of uptime, the drive had absorbed around 37 TB of writes. Annualized, that's roughly 640 terabytes per year. A typical 1 TB consumer SSD is rated for about 600 TBW lifetime — so this bug, if left unchecked, could burn through your drive's entire warranted endurance in less than a year.
The culprit is a logging configuration that probably nobody meant to ship to end users. Codex's SQLite feedback sink runs at global TRACE level by default — the noisiest possible setting. It logs everything from raw WebSocket payloads to mundane filesystem events like opening 'passwd' and 'ld.so.cache'. It also ignores the standard RUST_LOG environment variable, so there's no obvious way to turn it down. Around 71% of the logged data is TRACE-level noise that has no real diagnostic purpose, for the average user at least.
What makes it worse is write amplification. The database isn't just growing, but also cycling through tens of thousands of insert-and-delete operations per minute. It is physically writing far more to the drive than the file size implies.
This has actually been a known issue in various forms since at least April, with multiple related reports filed across the year. OpenAI's recent changelog touched some SQLite reliability fixes but hasn't addressed the write rate problem. The issue is still wide open.
In the meantime, Linux and macOS users can symlink '~/.codex/logs_2.sqlite' to '/tmp/' to redirect writes to RAM. The file holds no conversation data, so losing it on reboot is fine.
Featured image by Zac Wolff on Unsplash
static version load dynamic Loading Comments
Comment on this article
Related Articles
