---
source: "https://twitter.com/ClaudeDevs/status/2085794862608318627"
hn_url: "https://news.ycombinator.com/item?id=49214994"
title: "Claude Code: Starting August 14, auto mode will be the default permission mode"
article_title: "ClaudeDevs on X: \"Starting August 14, auto mode will be the default permission mode in Claude Code for Pro, Max, and Team users.\nAuto mode reviews shell commands and actions with a separate classifier. In testing, it caught 89% of dangerous commands. Manual approval caught 14%.\" / X"
author: "tosh"
captured_at: "2026-08-07T19:44:26Z"
capture_tool: "hn-digest"
hn_id: 49214994
score: 2
comments: 2
posted_at: "2026-08-07T19:11:00Z"
tags:
  - hacker-news
  - translated
---

# Claude Code: Starting August 14, auto mode will be the default permission mode

- HN: [49214994](https://news.ycombinator.com/item?id=49214994)
- Source: [twitter.com](https://twitter.com/ClaudeDevs/status/2085794862608318627)
- Score: 2
- Comments: 2
- Posted: 2026-08-07T19:11:00Z

## Translation

タイトル: クロード コード: 8 月 14 日以降、自動モードがデフォルトの許可モードになります
記事のタイトル: ClaudeDevs on X: 「8 月 14 日より、自動モードが Pro、Max、および Team ユーザーの Claude Code のデフォルトの権限モードになります。
自動モードでは、シェルのコマンドとアクションを別の分類子でレビューします。テストでは、危険なコマンドの 89% を検出しました。手動承認は 14% でした。」 /X
説明: 8 月 14 日より、Pro、Max、および Team ユーザーの Claude Code のデフォルトの権限モードは自動モードになります。
自動モードでは、シェルのコマンドとアクションを別の分類子でレビューします。テストでは、危険なコマンドの 89% を検出しました。手動承認は 14% でした。

記事本文:
ClaudeDevs @ClaudeDevs 8 月 14 日より、Pro、Max、Team ユーザーの Claude Code のデフォルトの権限モードは自動モードになります。
自動モードでは、シェルのコマンドとアクションを別の分類子でレビューします。テストでは、危険なコマンドの 89% を検出しました。手動承認は 14% でした。 6:27 PM · 2026 年 8 月 7 日 237.5K 再生数 242 150 3.7K 359
ClaudeDevs @ClaudeDevs 1h この変更を行う理由は 2 つあります。
1. 当社のテストでは、追跡したすべての安全対策において、自動モードが手動の許可レビューと一致するか、それを上回りました。
2. 長期的な作業がより実行可能になります。 Claude は中断間の実行時間が長いため、複数時間のタスクを実行できます。 詳細を表示 10 5 308 25K ClaudeDevs @ClaudeDevs 1h 私たちが手動承認よりも信頼している理由の 1 つは、1,053 人の有料テスターを対象とした調査で、明らかに危険なコマンド (テキストのみで、実際には何も実行されません) の許可プロンプトを交換したことです。
テスターは 13.6% の確率でそれを発見し、50 回のプロンプト後には 5% 近くになりました。自動モードは、表示をブロックしました 8 11 279 29K ClaudeDevs @ClaudeDevs 1h 自動モードは、Opus 5 などの長時間実行タスクに優れたモデルともよく組み合わせられます。モデルは複数時間のタスクを計画して実行できますが、それが効果を発揮するのは、ユーザーを待っているすべてのツール呼び出しでセッションが停止しない場合のみです。
自動モードでは、タスク (または複数のタスク) を開始できます。 詳細を表示 5 3 112 14K ClaudeDevs @ClaudeDevs 1h 自動モードでは、ツール呼び出しごとに分類子が実行され、少数の追加トークンが使用されます。
本日の時点で、その分類器のオーバーヘッドは、Pro、Max、および Team プランの使用制限にはカウントされなくなりました。この変更は、自動モードが [表示] になると、エンタープライズ プランと API ユーザーに適用されます。 詳細を表示 6 5 138 17K ClaudeDevs @ClaudeDevs 1h デフォルトが切り替わるとアプリ内通知が届きます。すでにデフォルトの権限モードを設定している場合、Claude は

何かを変更する前に尋ねてください。 Shift+Tab でいつでもモードを切り替えることができます。
管理者は、管理設定でdefaultModeを固定する（または自動モードを完全に無効にする）ことができます。 3 2 84 18K ClaudeDevs @ClaudeDevs 1h 完全なブログ: Pro、Max、およびTeamプランのClaude Codeでは自動モードがデフォルトになりました |クロード by Anthropic From claude.com 4 2 77 17K
Steve Li @st3v3li 59m これについては使用量をリセットできますか? 2 147 2.9K
何が起こっているかを確認して会話に参加してください
電話で続行 Apple で続行 Google で続行、またはユーザー名またはメールでログイン 関係者

## Original Extract

Starting August 14, auto mode will be the default permission mode in Claude Code for Pro, Max, and Team users.
Auto mode reviews shell commands and actions with a separate classifier. In testing, it caught 89% of dangerous commands. Manual approval caught 14%.

ClaudeDevs @ClaudeDevs Starting August 14, auto mode will be the default permission mode in Claude Code for Pro, Max, and Team users.
Auto mode reviews shell commands and actions with a separate classifier. In testing, it caught 89% of dangerous commands. Manual approval caught 14%. 6:27 PM · Aug 7, 2026 237.5K Views 242 150 3.7K 359
ClaudeDevs @ClaudeDevs 1h We’re making this change for two reasons:
1. In our testing, auto mode matched or beat manual permission review on every safety measure we tracked.
2. It makes long-horizon work more viable. Claude runs longer between interruptions, so you can run multi-hour tasks in the Show more 10 5 308 25K ClaudeDevs @ClaudeDevs 1h One reason we trust it more than manual approval: in a study with 1,053 paid testers, we swapped a permission prompt for a clearly dangerous command (text only, nothing actually ran).
Testers caught it 13.6% of the time, and closer to 5% after 50 prompts. Auto mode blocked the Show more 8 11 279 29K ClaudeDevs @ClaudeDevs 1h Auto mode also pairs well with models that excel at long-running tasks, like Opus 5. The model can plan and execute multi-hour tasks, but that only pays off if the session doesn't stop at every tool call waiting for you.
With auto mode, you can kick off a task (or multiple in Show more 5 3 112 14K ClaudeDevs @ClaudeDevs 1h Auto mode runs a classifier on every tool call, which uses a small number of extra tokens.
As of today, that classifier overhead no longer counts toward your usage limits on Pro, Max, and Team plans. This change will come to Enterprise plans and API users once auto mode is the Show more 6 5 138 17K ClaudeDevs @ClaudeDevs 1h You’ll get an in-app notification when the default flips. If you’ve already set a default permission mode, Claude will ask before changing anything. You can still switch modes with Shift+Tab anytime.
Admins can pin defaultMode (or disable auto mode entirely) in managed settings 3 2 84 18K ClaudeDevs @ClaudeDevs 1h Full blog: Auto mode is now the default in Claude Code for Pro, Max, and Team plans | Claude by Anthropic From claude.com 4 2 77 17K
Steve Li @st3v3li 59m Do we get a usage reset for this? 2 147 2.9K
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
