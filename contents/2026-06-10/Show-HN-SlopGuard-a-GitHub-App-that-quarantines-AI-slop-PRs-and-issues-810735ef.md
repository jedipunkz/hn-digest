---
source: "https://slopguard.app"
hn_url: "https://news.ycombinator.com/item?id=48479791"
title: "Show HN: SlopGuard, a GitHub App that quarantines AI slop PRs and issues"
article_title: "SlopGuard: AI slop PR/Issue guard for maintainers"
author: "blue-b"
captured_at: "2026-06-10T19:00:44Z"
capture_tool: "hn-digest"
hn_id: 48479791
score: 2
comments: 0
posted_at: "2026-06-10T17:35:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SlopGuard, a GitHub App that quarantines AI slop PRs and issues

- HN: [48479791](https://news.ycombinator.com/item?id=48479791)
- Source: [slopguard.app](https://slopguard.app)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T17:35:21Z

## Translation

タイトル: Show HN: SlopGuard、AI スロップ PR と問題を隔離する GitHub アプリ
記事タイトル: SlopGuard: メンテナ向け AI スロップ PR/問題ガード
説明: AI スロップ プル リクエストと問題をスコアリングし、出所にタグを付け、それらを隔離する GitHub アプリです。最終的な判断を下すのは常に人間です。決して自動で閉まらない。

記事本文:
SlopGuard: AI スロップ PR/メンテナー向けの問題ガード SlopGuard の仕組み 価格ドキュメント EN KO AI スロップによるリポジトリの溺死を阻止
AI が生成した PR や問題は最初の 30 秒間は現実的に見えますが、その後は時間を無駄にします。 SlopGuard はすべてのメッセージを読み取り、機械製のメッセージに理由を付けてフラグを立て、1 回のパスでキューをクリアできます。それは決して何も閉じません。あなたが電話をかけます。
GitHub アプリをインストールする 仕組みを確認する → 0 ～ 100 のスロップ スコア
自分用のソース利用可能なセルフホスト
これに見覚えがある場合は、SlopGuard が必要です
開けるには十分にもっともらしく見えますが、モデルが作ったことがわかります。それぞれ 30 秒、週に数十回です。
すべての号と PR 通知は、繰り返しになりますが、本物かいい加減かを 1 つずつ判断することを意味します。
パターンだけを終わらせると、最終的には初めての投稿者は燃えてしまいます。だから、あなたはそうではありません。
SlopGuard はすべてのファイルを読み取り、機械製のファイルに理由を付けてフラグを付けるため、本物のファイルを 1 回のパスで確認できます。決定はあなた次第です。
スライダーをドラッグしてスロップスコアラインを移動します。 PR はリアルタイムで隔離済みと合格済みを再分類します。
PR #218 依存関係を最新の 16 パスにバンプする
PR #241 認証コールバック 43 パスをリファクタリングする
#312 機能リクエスト、再現 28 パスをクリア
PR #233 79 検疫に絵文字見出しのドキュメント ページを 12 追加
PR #239 README (機械生成) 93 隔離を書き換える
しきい値以上では、SlopGuard はラベルとレビュー コメントを追加します。その下では沈黙が保たれます。境界線を設定するのはあなたです。人間は常に最後の言葉を持っています。
すべての投稿には 0 ～ 100 のスロップ スコア、その背後にある理由、および来歴の証跡が付けられます。 SlopGuard のラベルとコメント。あなたが決めてください。
チャット アシスタントの定型フレーズ (3)
流出したフレーズ「確かに！ これが更新されたコードです」
一般的な自動生成されたタイトル、空の説明
来歴: プロンプト フィンガープリント b01706d4、機械生成

編
メンテナは /slop accept 、 /slop拒否 、または /slop false-positive と応答します
AI スロップとは、人によるレビューがほとんど、またはまったくないまま、リポジトリにプッシュされる低労力の機械生成コンテンツです。LLM によって作成された PR や問題は、もっともらしいように見えますが、コンテキストを見逃し、何も追加せず、ほとんどがファームのアクティビティや貢献統計に存在します。
洗練された表現、きちんとしたタイトル、自信に満ちた要約、しかしコードベースの本当の理解はありません。
メンテナは依然としてそれを読み、質問し、それがどこにも行かないように解決する必要があります。
1 つのプロンプトは、10 個のプロジェクトにわたってほぼ同一の 10 個の PR になります。
すべてのコミットは、デフォルトのしきい値 50 で、手動でラベル付けされた 25 ケースのゴールデン セット、つまり 13 件の実際のスロップと 12 件の正当な貢献に対して SlopGuard のスコアを再スコアします。ヒューリスティックのみで、LLM キーはありません。これがまさにその方法です。
フラグを立てたすべてのことのうち、すべてが本当にひどいものでした。誤報ゼロ。
13 件の実際のスロップ ケースのうち 12 件を捕らえました。
精度と再現率が 1 つのスコアに結合されます。
25 問中 24 問正解です。この 1 回のミスは、実際にすり抜けてしまったものであり、実際の投稿者が誤ってフラグを立てたわけではありません。本物の広報担当者が処罰されたことは一度もなかった。
トリアージ支援は貢献者を尊重し、決して核攻撃にはなりません。
ワンクリックでリポジトリまたは組織にインストールします。アクション YAML、CI 構成、接続するシークレットは必要ありません。
隔離ラベルとレビュー コメントのみ。明示的なメンテナコマンドがなければ何も閉じられません。
フラグ ジェネレータのヒント、プロンプト フィンガープリント、および「AI モデルとして」などの漏洩したアシスタント フレーズ。
しきい値、ラベル、許可リスト、コメント テンプレートはリポジトリ内に存在し、他の変更と同様にレビューされます。
ヒューリスティック専用モードは API キーを使用せずに実行されますが、ゴールデン セットでは依然として 100% の精度に達します。
状態は GitHub のラベルと問題に存在します。自分で使用するためにすべてを自己ホストします。ソースはコモンズ条項に基づいて入手可能です。
ノイズをフィルターして、実際の作業を維持します。

SlopGuard はラベルとコメントのみです。閉店するかどうかは常にあなたの判断です。
GitHub メンテナー向けの AI スロップトリアージ。ラベルとコメントを付けますが、実際の投稿者を自動的にクローズすることはありません。

## Original Extract

A GitHub App that scores AI slop pull requests and issues, tags provenance, and quarantines them. A human always makes the final call. Never auto-closes.

SlopGuard: AI slop PR/Issue guard for maintainers SlopGuard How it works Pricing Docs EN KO Stop AI slop from drowning your repo
AI-generated PRs and issues look real for the first 30 seconds, then waste your time. SlopGuard reads every one, flags the machine-made ones with the reasons, and lets you clear your queue in one pass. It never closes anything; you make the call.
Install the GitHub App See how it works → 0–100 slop score
source-available self-host for yourself
You need SlopGuard if any of this sounds familiar
They look plausible enough to open, then you realize a model made it. Thirty seconds each, a few dozen a week.
Every issue and PR notification means deciding, again, real or slop, one by one.
Closing on pattern alone eventually burns a real first-time contributor. So you don't.
SlopGuard reads every one and flags the machine-made ones with the reasons, so you review the real ones in a single pass. The decision stays yours.
Drag the slider to move the slop-score line. PRs reclassify between quarantined and passed in real time.
PR #218 Bump dependencies to latest 16 pass
PR #241 Refactor the auth callback 43 pass
#312 Feature request, clear repro 28 pass
PR #233 Add 12 emoji-headed doc pages 79 quarantine
PR #239 Rewrite README (machine-generated) 93 quarantine
At or above the threshold SlopGuard adds a label and a review comment. Below it, it stays silent. You set the line; a human always has the final word.
Every contribution gets a 0 to 100 slop score, the reasons behind it, and a provenance trail. SlopGuard labels and comments. You decide.
Chat-assistant boilerplate phrases (3)
Leaked phrase “Certainly! Here is the updated code”
Generic auto-generated title, empty description
Provenance: prompt fingerprint b01706d4, machine-generated
maintainer replies /slop approve , /slop reject , or /slop false-positive
AI slop is low-effort, machine-generated content pushed to a repo with little or no human review: PRs and issues spun up by an LLM that look plausible but miss the context, add nothing, and mostly exist to farm activity or contribution stats.
Polished phrasing, a tidy title, a confident summary, and no real grasp of the codebase.
A maintainer still has to read it, ask questions, and work out that it goes nowhere.
One prompt becomes ten near-identical PRs across ten projects.
Every commit re-scores SlopGuard against a hand-labelled golden set of 25 cases, 13 real slop and 12 legit contributions, at the default threshold of 50. Heuristics only, no LLM key. Here is exactly how it did.
Of everything it flagged, all of it was real slop. Zero false alarms.
It caught 12 of the 13 real slop cases.
Precision and recall combined into one score.
24 of 25 correct. The single miss was real slop that slipped through, not a real contributor wrongly flagged. It never once punished a genuine PR.
Triage help that respects contributors and never goes nuclear.
Install on a repo or org in one click. No Action YAML, no CI config, no secrets to wire.
Quarantine label and a review comment only. Nothing is closed without an explicit maintainer command.
Flags generator hints, a prompt fingerprint, and leaked assistant phrases like “As an AI model”.
Thresholds, labels, allowlists, and comment templates live in your repo, reviewed like any other change.
Heuristics-only mode runs with zero API keys, and still hits 100% precision on the golden set.
State lives in GitHub labels and issues. Self-host the entire thing for your own use; the source is available under the Commons Clause.
Filter the noise, keep the real work.
SlopGuard only labels and comments. Whether to close is always your call.
AI-slop triage for GitHub maintainers. It labels and comments, never auto-closes a real contributor.
