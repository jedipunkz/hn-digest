---
source: "https://www.jackfranklin.co.uk/blog/ai-review-plan/"
hn_url: "https://news.ycombinator.com/item?id=48482526"
title: "AI-review: reviewing AI code before it lands"
article_title: "ai-review: reviewing AI code before it lands - Jack Franklin"
author: "mooreds"
captured_at: "2026-06-10T21:44:32Z"
capture_tool: "hn-digest"
hn_id: 48482526
score: 1
comments: 0
posted_at: "2026-06-10T20:51:30Z"
tags:
  - hacker-news
  - translated
---

# AI-review: reviewing AI code before it lands

- HN: [48482526](https://news.ycombinator.com/item?id=48482526)
- Source: [www.jackfranklin.co.uk](https://www.jackfranklin.co.uk/blog/ai-review-plan/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T20:51:30Z

## Translation

タイトル: AI レビュー: AI コードを実装前にレビューする
記事のタイトル: ai-review: AI コードを着陸前にレビューする - ジャック フランクリン

記事本文:
ai-review: AI コードを実装前にレビューする
長年にわたり、私はほとんどの時間をターミナルで過ごす開発者として非常に快適になりました。私はドットファイルを自分の好みに合わせて調整するのにあまりにも長い時間を費やしており、15 年間のうちの大半は Vim/Neovim ユーザーでした。
この傾向はAI時代にも続いています。さまざまな IDE や UI ツールを試してきましたが、いつも Claude Code、Gemini CLI などに戻ってきます。
端末AIツールの問題点
私が常に IDE ユーザーをうらやましく思っている分野が 1 つあります。それは、AI の計画とコードをレビューすることです。クロードは端末に計画を提示しますが、詳細に検討するのは非常に難しく、フィードバックを残すための特定の行を強調表示することはできません。コードが記述されると、同じ問題が発生します。少なくとも私にとっては、端末の差分をスクロールするのは楽しくないし、問題を見つける良い方法でもありません。
AI Review は、ブラウザーで計画または git diff を確認する方法を提供します。一行ずつフィードバックを残すことができ、一般的なコメントのためのスペースも用意されています。完了すると、コメントとコメントが参照する行は、エージェントが対応できるように会話に直接戻されます。
これにより、マウスとキーボードを使用してブラウザーで計画を直観的に確認できる、端末 AI エージェントのような快適さが得られます。
AI と協力して作品を計画します。
AI に指示する: /review-plan 最終計画をまとめる
ブラウザに読み込まれたら、レビューしてフィードバックを残します。満足するまでこれらの手順を繰り返すことができます。
AI にプロンプトを表示します: /review-diff コードをレビューさせてください
ブラウザで差分を確認し、1 行ずつコメントを残してください。
視覚的に学習したい方のために、このツールを使用して計画をレビューするワークフローを示すビデオを録画しました。
完全な手順は GitHub で見つけることができますが、

TLDR は次のとおりです。
npm install -g @jackfranklin/ai-review (注: npm パッケージは ai-review-plan ではなく、単なる @jackfranklin/ai-review です)
/review-plan スキルをエージェントに追加します。これは、開始点として使用できる GitHub リポジトリのテンプレートです。
フィードバックがある場合は、ここでさまざまなプラットフォームで私を見つけるか、GitHub リポジトリに問題を残してください。計画とコードレビューを楽しんでください!
この投稿が気に入ったら、 Bluesky で私をフォローしてみてください。

## Original Extract

ai-review: reviewing AI code before it lands
Over the years I have become very comfortable as a developer who spends most of his time in the terminal. I've spent far too long tweaking my dotfiles to make things just how I like, and I have been a Vim/Neovim user for the best part of fifteen years.
This trend has continued into the AI era. I've tried various IDEs and UI tools, but I always come back to Claude Code, Gemini CLI, and others.
The problem with terminal AI tools
There is one area where I have always been jealous of IDE users: reviewing AI plans and code. Claude will present a plan in the terminal, but it's just hard to review in detail, and I cannot highlight specific lines to leave feedback on. The same problem hits once the code is written — scrolling through a terminal diff isn't fun nor is it a good way to spot problems, at least not for me.
AI Review offers you a way to review the plan, or a git diff, in the browser. You can leave line by line feedback, and have space for general comments too. Once you're done, your comments and the lines they reference are passed straight back into the conversation for the agent to act on.
This gives you the comfort of a terminal AI agent, with the ability to intuitively review the plan in the browser, using your mouse and keyboard.
Work with the AI to plan a piece of work.
Prompt the AI: /review-plan put the final plan together
Once it loads up in the browser, I review and leave feedback. I can repeat these steps until I am happy.
Prompt the AI: /review-diff let me review your code
Review the diff in the browser, and leave line-by-line comments.
If you're more of a visual learner, I recorded a video which shows my workflow for reviewing plans with this tool .
You can find full instructions on GitHub but the TLDR is:
npm install -g @jackfranklin/ai-review (note: the npm package is just @jackfranklin/ai-review , not ai-review-plan )
Add the /review-plan skill to your agent. Here is a template from the GitHub repo that you can use as a starting point .
If you have feedback, you can find me on various platforms here or leave an issue on the GitHub repo. Happy plan and code reviewing!
If you enjoyed this post, you might like to follow me on Bluesky .
