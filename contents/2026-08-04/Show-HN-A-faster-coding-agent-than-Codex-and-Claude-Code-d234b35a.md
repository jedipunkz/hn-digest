---
source: "https://www.codewithbullet.com"
hn_url: "https://news.ycombinator.com/item?id=49173799"
title: "Show HN: A faster coding agent than Codex and Claude Code"
article_title: "Bullet · Fast, by design."
author: "alsima"
captured_at: "2026-08-04T20:19:37Z"
capture_tool: "hn-digest"
hn_id: 49173799
score: 3
comments: 4
posted_at: "2026-08-04T19:33:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A faster coding agent than Codex and Claude Code

- HN: [49173799](https://news.ycombinator.com/item?id=49173799)
- Source: [www.codewithbullet.com](https://www.codewithbullet.com)
- Score: 3
- Comments: 4
- Posted: 2026-08-04T19:33:34Z

## Translation

タイトル: Show HN: Codex や Claude Code より高速なコーディング エージェント
記事のタイトル: Bullet · 設計により高速です。
説明: Bullet は、より無駄のないオーケストレーション層を備えた非常に高速なコーディング エージェントです。
HN テキスト: こんにちは、HN、今日は皆さんとこれを共有できることに興奮しています。 TL;DR: Bullet は高速コーディング エージェントです。私たちの会社では、エージェントの実行を待機するのに何時間も費やしていました。モデルは問題ではありませんでした。エージェントのループが遅かった (不必要な計画、大きすぎるコンテキスト、繰り返しの検索、およびシリアル ツールの呼び出し)。そこで私たちは情熱を注ぐプロジェクトとして Bullet を構築しました。その目標は、品質を犠牲にすることなく、より速く反復できるようにすることです。
- 各プロンプトのモデルと推論レベルが自動的に選択されます。
- 独立した検索、読み取り、コマンドを並行して実行します。
- リポジトリ全体を埋め込むのではなく、対象を絞ったコード検索を使用します。
- Claude Code または Codex サブスクリプション、OpenAI/Anthropic/xAI API キー、またはキーのないオンデバイス モデルで動作します。 Claude Code と Codex はそれぞれ、独自のモデル エコシステムを中心としています。 OpenCode では多くのモデルが提供されますが、モデルの選択と構成はほとんどがユーザーに任されています。弾丸はモデルの上にあります。すでに支払ったものを接続し、Bullet にあらゆるタスクの最速パスを選択させます。 SWE ベンチ検証では、Bullet は 479/500 の問題 (95.8%、リーダーボードのトップ 3) を 1 回の試行で解決し、タスクあたり平均 119 秒かかりました (ミニ SWE エージェント + Fable/Sol より 35 ～ 67% 高速)。完全な結果と方法論はこちら: https://www.codewithbullet.com/blog/benchmark-results.html 製品に関する率直なフィードバックをお待ちしております。ダウンロードとセットアップには 1 分もかかりません。 Bullet の好きなところ、悪いところをぜひ聞かせてください。 Mac アプリの DMG リリースは当社の Web サイトにあります: https://www.codewithbullet.com/ 追伸: Web サイトにコードを隠しました。

フッターにある秘密ページのロックを解除します (すべて Bullet で構築されています) :) https://youtu.be/34wfiSoSw4Q

記事本文:
ブログ
/
キャリア
1.3.10をダウンロード
B-001 / はじめに
MACOS · プライベートベータ
無料アクセス
新しい SWE-ベンチで 95.8% が検証済み 結果を読む →
本当に高速なコーディングエージェント
弾丸のルート、検索、および実行は、ユーザーの状況に追いつくという 1 つの目的のために行われます。
あなたのアイデアは素早く動きます。
あなたのエージェントもそうすべきです。
エージェントの実行を待つのに何時間も費やしていました。モデルが無能だったからではありません。周囲の機械が必要以上に重かったのです。
同じモデル→ツール→結果パターン。
その周りのより緊密なループ。
簡単な作業を高速なモデルにルーティングします。ミッションで要求された場合にのみエスカレーションしてください。
ターゲットを獲得します。
ノイズは無視してください。
ターゲットを絞った検索とファイル読み取りにより、リポジトリ全体を埋め込まずに関連するコードが見つかります。
一緒に実行できるものをキューに入れないでください。
独立したツール呼び出しは並行して実行されます。重複した呼び出しとスタック ループは、さらに 1 秒も無駄になる前にインターセプトされます。
本当のプロンプト。リアルリポジトリ。
映画のようなショートカットはありません。
私たちの会社では、エージェントの実行を待機するのに何時間も費やしていました。私たちは毎日 Claude Code を使って構築し、不必要に遅い機械の中を有能なモデルが動くのを観察していました。
そこで、Bullet は情熱を注ぐプロジェクトとして始まりました。単純な作業をより速くルーティングし、重要なものだけを読み取り、独立したツールを一緒に実行し、ループがスパイラルになる前にループを停止します。
現在社内で使用しています。そのおかげで頭の痛い問題がなくなりました。私たちは、これであなたも救われるかもしれないと考えました。
コードを許可しないでください
あなたを人質に取ってください。
無料でご利用いただけます。購読はありません。
より早く発送する方法です。

## Original Extract

Bullet is a really fast coding agent with a leaner orchestration layer.

Hi HN, excited to be sharing this with you guys today. TL;DR: Bullet is a fast coding agent. At our company, we were burning hours waiting on agent runs. The models weren’t the problem. The agent loops around them were slow (unnecessary planning, oversized context, repeated searches, and serial tool calls). So we built Bullet as a passion project, the goal being to help you iterate faster, without giving up quality:
- It automatically selects the model and reasoning level for each prompt.
- It runs independent searches, reads, and commands in parallel.
- It uses targeted code search instead of embedding the entire repository.
- It works with your Claude Code or Codex subscription, OpenAI/Anthropic/xAI API keys, or an on-device model with no key. Claude Code and Codex each center on their own model ecosystem. OpenCode gives you many models, but largely leaves model selection and configuration to you. Bullet sits above the models. Connect what you already pay for, then let Bullet choose the fastest path for every task. On SWE-bench Verified, Bullet resolved 479/500 issues (95.8%, top 3 on the leaderboard) in one attempt, averaging 119 seconds per task (35–67% faster than mini-SWE-agent + Fable/Sol). Full results and methodology here: https://www.codewithbullet.com/blog/benchmark-results.html We would love honest feedback on the product! It takes less than a minute to download and set up. We’d love to hear what you like about Bullet, and what sucks. The Mac app DMG release is on our website: https://www.codewithbullet.com/ P.S: we hid a code on the website, see if you can unlock the secret page at the footer (all built with Bullet) :) https://youtu.be/34wfiSoSw4Q

BLOG
/
CAREERS
DOWNLOAD 1.3.10
B–001 / INTRODUCTION
MACOS · PRIVATE BETA
FREE ACCESS
NEW 95.8% ON SWE-BENCH VERIFIED READ THE RESULTS →
A REALLY FAST CODING AGENT
Bullet routes, searches, and executes with one purpose: keeping up with you.
YOUR IDEAS MOVE FAST.
YOUR AGENT SHOULD TOO.
We were burning hours waiting on agent runs. Not because the models were incapable. The machinery around them was simply heavier than it needed to be.
Same model → tools → results pattern.
A tighter loop around it.
Route straightforward work to fast models. Escalate only when the mission demands it.
Acquire the target.
Ignore the noise.
Targeted search and file reads find relevant code without embedding the whole repo.
Never queue what can run together.
Independent tool calls execute in parallel. Duplicate calls and stuck loops are intercepted before they waste another second.
Real prompt. Real repository.
No cinematic shortcuts.
At our company, we were burning hours waiting on agent runs. We were building with Claude Code every day, watching capable models move through needlessly slow machinery.
So Bullet began as a passion project: route simple work faster, read only what matters, run independent tools together, and stop loops before they spiral.
We use it internally now. It saved us a headache. We thought it might save you one too.
DON'T LET YOUR CODE
HOLD YOU HOSTAGE.
Free to use. No subscriptions.
Just a faster way to ship.
