---
source: "https://ajin.im/is/building/did-claude-just-reset-usage/"
hn_url: "https://news.ycombinator.com/item?id=48469288"
title: "Did Claude just reset usage?"
article_title: "did claude just reset usage?"
author: "poppypetalmask"
captured_at: "2026-06-10T00:59:59Z"
capture_tool: "hn-digest"
hn_id: 48469288
score: 2
comments: 0
posted_at: "2026-06-09T23:37:01Z"
tags:
  - hacker-news
  - translated
---

# Did Claude just reset usage?

- HN: [48469288](https://news.ycombinator.com/item?id=48469288)
- Source: [ajin.im](https://ajin.im/is/building/did-claude-just-reset-usage/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T23:37:01Z

## Translation

タイトル: クロードは使用量をリセットしただけですか?
記事のタイトル: クロードは使用量をリセットしただけですか?
説明: 人間のカナリア

記事本文:
anthropic は週の半ばにプロ/最大使用量カウンターをゼロにすることがあります。多くの場合、それは事件や計測バグの後の静かな補償です。時々、x に投稿されることがあります。多くの場合、そうではありません。いずれの場合でも、アプリ内通知、電子メール、ステータス ページのインシデント テキストには何も表示されません。
このページは、公式メーターが約 30 分ごとにポーリングされ、そこから回答する 1 つの重要な参照アカウント (カナリア) を監視します。週半ばのリセットは次のようになります。週が経過し、リセット日はそのままですが、used% はゼロに急落します。
カナリアはあなたのアカウントを確認できず、リセットはコホート範囲に限定される場合があります。端末から確認してください。このコマンドは、独自のローカル クロード コード トークンを読み取り、anthropic に直接質問します。
TOKEN=$(security find-generic-password -s "クロード コード-認証情報" -w |
python3 -c "import sys,json;print(json.load(sys.stdin)['claudeAiOauth']['accessToken'])")
カール -s https://api.anthropic.com/api/oauth/usage \
-H "認可: Bearer $TOKEN" -H "anthropic-beta: oauth-2025-04-20" \
| python3 -m json.tool
Linux/WSLのコピー
TOKEN=$(python3 -c "import json,os;print(json.load(open(os.path.expanduser(
'~/.claude/.credentials.json')))['claudeAiOauth']['accessToken'])")
カール -s https://api.anthropic.com/api/oauth/usage \
-H "認可: Bearer $TOKEN" -H "anthropic-beta: oauth-2025-04-20" \
| python3 -m json.tool
seven_day.utilization を読んでください。それが週の作業で設定されるはずの値よりはるかに下にあり、seven_day.resets_at が移動していない場合は、カウンタもリセットされています。
トークンを Web サイトに貼り付けないでください。
人々はこれらをバグとしてファイルします (claude-code #52497 、 #49616 )。それはバグではありません。
補償リセットは経過時間中に used% を低下させ、resets_at はそのままになります。通常の毎週のロールは両方ともゼロになります。再アンカーはウィンドウ自体を移動します。 2 つの間にある 1 つの奇妙な読み方

o 通常のものは測光がちらつき、カウントされません。
インシデントのコンテキストは、クロードのステータス ページから取得されます。リセット自体はインシデントのテキストには表示されません。

## Original Extract

A canary for Anthropic

anthropic sometimes zeroes pro/max usage counters mid-week. often it is quiet compensation after an incident or a metering bug. sometimes it gets a post on x. often it does not. either way there is no in-app notice, no email, nothing in the status-page incident text.
this page watches one heavy reference account (the canary), whose official meter is polled every ~30 minutes, and answers from that. a mid-week reset looks like this: used% plunges to zero while the week keeps elapsing and the reset date stays put.
the canary cannot see your account, and resets are sometimes cohort-scoped. check yours from a terminal. the command reads your own local claude code token and asks anthropic directly.
TOKEN=$(security find-generic-password -s "Claude Code-credentials" -w |
python3 -c "import sys,json;print(json.load(sys.stdin)['claudeAiOauth']['accessToken'])")
curl -s https://api.anthropic.com/api/oauth/usage \
-H "Authorization: Bearer $TOKEN" -H "anthropic-beta: oauth-2025-04-20" \
| python3 -m json.tool
linux / wsl copy
TOKEN=$(python3 -c "import json,os;print(json.load(open(os.path.expanduser(
'~/.claude/.credentials.json')))['claudeAiOauth']['accessToken'])")
curl -s https://api.anthropic.com/api/oauth/usage \
-H "Authorization: Bearer $TOKEN" -H "anthropic-beta: oauth-2025-04-20" \
| python3 -m json.tool
read seven_day.utilization . if it sits far below where your week's work should have it, and seven_day.resets_at has not moved, your counter was reset too.
never paste your token into a website.
people file these as bugs ( claude-code #52497 , #49616 ). it is not a bug.
a compensation reset drops used% while elapsed time and resets_at stay put. the normal weekly roll zeroes both together. a re-anchor moves the window itself. a single weird reading between two normal ones is metering flicker, and does not count.
incident context comes from the claude status page . the resets themselves do not appear in its incident text.
