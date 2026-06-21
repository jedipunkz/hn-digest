---
source: "https://danuker.go.ro/a-cheaper-and-safer-agentic-ai-workflow.html"
hn_url: "https://news.ycombinator.com/item?id=48621420"
title: "A cheaper and safer agentic AI workflow"
article_title: "A cheaper and safer agentic AI workflow - danuker | freedom & tech"
author: "danuker"
captured_at: "2026-06-21T19:31:16Z"
capture_tool: "hn-digest"
hn_id: 48621420
score: 2
comments: 0
posted_at: "2026-06-21T18:39:21Z"
tags:
  - hacker-news
  - translated
---

# A cheaper and safer agentic AI workflow

- HN: [48621420](https://news.ycombinator.com/item?id=48621420)
- Source: [danuker.go.ro](https://danuker.go.ro/a-cheaper-and-safer-agentic-ai-workflow.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T18:39:21Z

## Translation

タイトル: より安価で安全なエージェント AI ワークフロー
記事のタイトル: より安価で安全なエージェント AI ワークフロー - danuker |自由とテクノロジー
説明: 最近エージェント コーディングを実際に試してみました。料金は 0.034 ドルで、3 分で完了しました。それは2つの間違いを犯した。私自身の人間としての挑戦では 1 時間かかり、4 つの間違いを犯しました。より安価なモデルのサービス GLM-5.2 について聞いたことがありますが、多くのベンチマークでそれが有効であることが示されています…

記事本文:
ナビゲーションを切り替え
ダヌーカー |自由とテクノロジー
お問い合わせ/概要
より安価で安全なエージェント AI ワークフロー
最近エージェントコーディングを本格的に試してみました。料金は 0.034 ドルで、3 分で完了しました。それは2つの間違いを犯した。私自身の人間としての挑戦では 1 時間かかり、4 つの間違いを犯しました。
GLM-5.2 について聞いたのですが、多くのベンチマークが、わずか 3 か月前の主要な独自 AI と同等であると言っています。同じベンチマーク サイトで、モデル サービスである GMI Cloud を発見しました。
私は昨年アカウントを作成し、5 ドルの無料クレジットを受け取りました。現在、最低入金額は10ドルになっているようです。それは私にとっても問題ありません。
サービス上で API キーを作成します。
私は、世界中に点在するデータセンターで米国企業がホストするシンガポールのモデルに私の個人データへのアクセスを与えることにあまり乗り気ではありません。そこで、VirtualBox イメージに Debian をインストールし、そこに pi と Guest Additions をインストールしました。次に、プロジェクトのコピーを共有フォルダーとして共有しました。他には何もありません。
pi を設定し、フォルダー上で GLM-5.2 を解放しました。 5 分と 0.435 ドル後、エージェントの健全性テストは成功しました。さまざまな形式のさまざまなデータ ファイルを調べて、興味のある情報を含む Index.tsv を作成するように依頼しました。完璧な仕事をしてくれました。
私の CPU で Unsloth の Qwen3.6-35B-A3B-Q4_K_XL も同様に実行しましたが、1 時間以上かかりました (そして、私の時間とインタラクティブ性は 1 時間あたり 0.435 ドルをはるかに超える価値があります)。でも、どれくらい安く行けるでしょうか？ GMI が他に提供しているものを見ると、DeepSeek V4 Flash が目に留まりました。 GLM-5.2 よりも少し冗長になっているようで、タスクあたりのトークン数は同じですが、コストは 10 分の 1 未満です。それでも私のタスクを実行できますか?
zai-org/GLM-5.2-FP8 を deepseek-ai/DeepSeek-V4-Flash に置き換えてテストを再実行します。
3 分で完了、0.034 ドル。わずかな不完全性が見られます。2 つの霧が発生しました。

アケ。一部の不規則なデータ系列は、5 日程度と 2 日程度の期間がありますが、「日次」として表示されます。でもそれ以外は大丈夫です。また、その中間にある deepseek-ai/DeepSeek-V4-Pro にも気付きました。私のテストではミスはありませんでしたが、かかった時間は 2 分 27 秒、0.229 ドルでした。 GLM の代わりにこれを保持すると思いますが、主に V4-Flash を使用することになります。
私の ~/.pi/agent/models.json は次のようになりました。
{
「プロバイダー」: {
「オラマ」: {
"baseUrl": "https://api.gmi-serving.com/v1",
"api": "openai-completions",
"apiKey": "ほぼ無料ですが、無料ではありません。非常に非常に安いです。",
"互換性": {
"supportsDeveloperRole": false、
"supportsReasoningEffort": false
}、
「モデル」: [
{
"id": "ディープシーク-ai/ディープシーク-V4-フラッシュ",
「推論」：本当、
「コンテキストウィンドウ」: 262144
}、
{
"id": "ディープシーク-ai/ディープシーク-V4-Pro",
「推論」：本当、
「コンテキストウィンドウ」: 262144
}
]
}
}
}
特に 4 つの間違いを犯し、1 時間以上かかったということを考えると、 mm/dd/yyyy 形式を呪ってください!私はその任務において徹底的に熟練したようです。自分のキャリアパスを調整して、時代についていきたいと感じています。
ボーナス: さらに安くなります: モデルが巨大な 1 行 JSON に遭遇し、トークン数が増加して pi の 50KB DEFAULT_MAX_BYTES 制限を埋めてしまうことが時々あります。この制限を 5KB に変更し、入力トークン数を大幅に減らしました。設定としてこれを導入するためのチケットがありますが、自動的にクローズされました。変更するファイルは次のとおりです (これを書いている時点の pi バージョン)。
～/。ローカル/共有/pi-node/node-v22 。 22.3-linux-x64/lib/node_modules/@earendil-works/pi-coding-agent/node_modules/@earendil-works/pi-agent-core/dist/harness/utils/truncate 。 js
～/。ローカル/共有/pi-node/node-v22 。 22.3 - linux - x64/lib/node_modules/@earendil - 動作/pi - コーディング - エージェント

/dist/core/tools/truncate 。 js
両方を変更しました（必要かどうかはわかりません）。 DeepSeek-V4-Flash のプロンプト トークンは 604k から 431k に増加し、私の特定のテストでは総コストは 0.034 ドルから 0.026 ドルに増加しました。
私の仕事は大きく変わりました。小さなコードセグメントを手動でコピー＆ペーストすることはなくなり、代わりにエージェントに何かをするよう依頼し、エージェントの共有ディレクトリとメイン/承認済みのディレクトリを比較します。私は、優れた diff ディレクトリ インターフェイスを備えた PyCharm を使用してこれを実行しますが、Meld を使用しても同様に実行できます。
これで完了です。私はベンダー ロックインに屈することを拒否しながら、AI の恩恵を受けています。私は閉鎖的なエコシステムや社会化を大嫌いです。 Anthropic が Claude Code の独占を推進し始めたとき、私はそれが反競争的であると感じました。また、オープンウェイトモデルが数か月遅れているのに、恣意的かつ突然の値上げは無謀です。彼らは急速に価値が下がっているモデルから価値を引き出そうと必死に努力している。猛烈なペースが遅くなると、その価値は蒸発してしまいます。彼らが築こうとしている堀は、実際には自らの速度を遅らせている。
ハードウェアはそこにあります。 GMI や DeepInfra などの企業は、柔軟な価値を提供します。また、ハードウェアは 3 か月ではなく、少なくとも数年で価値が落ちます。ただし、モデルを持続的に収益化する方法がわかりません。もしかしてクラウドファンディング？非営利ですか？消防士のような公共事業？
モデル API を取得する — GMI Cloud (または DeepInfra などのプロバイダー) でアカウントを作成します。
コンソールで API キーを作成します。
環境を分離します (プライバシーのため、強くお勧めします) — VirtualBox (またはその他の任意のサンドボックス) に Debian をインストールします。
エージェントをインストールする — VM 内に pi をインストールします。ハードコードされた DEFAULT_MAX_BYTES を変更します。
プロバイダーを構成する — ~/.pi/agent/models.json を編集して、A を指す OpenAI 互換プロバイダーを追加します。

PI エンドポイント。
健全性チェックを実行する - エージェントにタスクを与え、結果を確認します。
レビュー ワークフローをセットアップします。コピーを VM に共有し、2 つのコピー間でコードをレビューし続けます。
Pelican、静的サイト/ブログ エンジン
Python プログラミング言語
内容
特に明記されている場合を除き、クリエイティブ コモンズ表示 4.0 国際ライセンスに基づいてライセンスされています。

## Original Extract

I recently tried agentic coding for real. It cost $0.034 and finished in 3 minutes. It made two mistakes. In my personal human attempt, I took an hour, and made four mistakes. Cheaper model services I heard about GLM-5.2, and a lot of benchmarks are saying it's on …

Toggle navigation
danuker | freedom & tech
Contact/About
A cheaper and safer agentic AI workflow
I recently tried agentic coding for real. It cost $0.034 and finished in 3 minutes. It made two mistakes. In my personal human attempt, I took an hour, and made four mistakes.
I heard about GLM-5.2, and a lot of benchmarks are saying it's on par with the leading proprietary AIs of just 3 months ago . On the same benchmark site I had discovered GMI Cloud , a model service.
I created an account and received $5 in free credits last year. I see the minimum deposit is $10 nowadays. That's fine for me too.
I create an API key on their service.
I am not too keen on giving a Singaporean model hosted by a US company on data centers scattered throughout the world access to my private data. So I installed Debian in a VirtualBox image, and installed pi and the Guest Additions on it. Then I shared a copy of my project as a Shared Folder. Nothing else.
I configured pi and unleashed GLM-5.2 on the folder. 5 minutes and $0.435 later, the agentic sanity test worked. I asked it to look through various data files of various formats and create an index.tsv with information of interest. It did a perfect job.
So did Qwen3.6-35B-A3B-Q4_K_XL from Unsloth on my CPU, but it took more than an hour (and my time and interactivity is worth way more than $0.435 per hour). But how cheap could I go? Looking at what else GMI has to offer , DeepSeek V4 Flash catches my eye. It looks like it's a tiny bit more verbose than GLM-5.2 , so the same number of tokens per task, but less than a 10th of the cost. Can it still perform my task?
I replace zai-org/GLM-5.2-FP8 with deepseek-ai/DeepSeek-V4-Flash and rerun the test.
Done in 3 minutes and $0.034. It shows a tiny bit of imperfection: it made 2 mistakes. Some irregular data series are shown as "daily" though they've got 5-ish-day and 2-ish-day periods. But other than that it's fine. I also noticed deepseek-ai/DeepSeek-V4-Pro , which is somewhere in the middle. Zero mistakes on my test, but took 2 mins 27s and $0.229. I think this is the one I will keep instead of GLM, but I will mostly use V4-Flash.
My ~/.pi/agent/models.json ended up like so:
{
"providers": {
"ollama": {
"baseUrl": "https://api.gmi-serving.com/v1",
"api": "openai-completions",
"apiKey": "Almost free but not free. Very, very cheap.",
"compat": {
"supportsDeveloperRole": false,
"supportsReasoningEffort": false
},
"models": [
{
"id": "deepseek-ai/DeepSeek-V4-Flash",
"reasoning": true,
"contextWindow": 262144
},
{
"id": "deepseek-ai/DeepSeek-V4-Pro",
"reasoning": true,
"contextWindow": 262144
}
]
}
}
}
Especially considering that I made 4 mistakes, and that it took me a bit more than an hour. Curse the mm/dd/yyyy format! It seems I have been thoroughly bested at that task. I feel like adjusting my career path and keeping up with the times.
Bonus: Go even cheaper: Every so often, my models stumbles into a huge one-line JSON, and runs up the token count filling up pi's 50KB DEFAULT_MAX_BYTES limit. I changed that limit to 5KB, significantly reducing input token count. There is a ticket to introduce this as a setting, but it was auto-closed. The files to modify (with the pi version as of writing this) are:
~/. local / share / pi - node / node - v22 . 22.3 - linux - x64 / lib / node_modules / @ earendil - works / pi - coding - agent / node_modules / @ earendil - works / pi - agent - core / dist / harness / utils / truncate . js
~/. local / share / pi - node / node - v22 . 22.3 - linux - x64 / lib / node_modules / @ earendil - works / pi - coding - agent / dist / core / tools / truncate . js
I modified both (not sure if I needed to). Prompt tokens for DeepSeek-V4-Flash went from 604k to 431k, and total cost went from $0.034 to $0.026 for my particular test.
My work now changed significantly. No longer do I manually copy paste tiny code segments, instead I ask the agent to do something, then compare the agent's shared directory with the main/approved one. I do this with PyCharm for its good diff directory interface, but you can do it with Meld as well.
So there you have it. I am reaping the AI rewards, while refusing to give in to vendor lock-in. I despise closed ecosystems and enshittification. When Anthropic started pushing for Claude Code exclusivity, I found that anticompetitive. Also, arbitrary and sudden price increases are reckless while open weights models are just a few months behind. They are desperately trying to extract value from rapidly devaluing models. If the breakneck pace slows down, their value evaporates. The moat they are trying to build is actually slowing themselves down.
The hardware is where it's at. Companies like GMI and DeepInfra provide flexible value. And the hardware depreciates in at least a few years, rather than 3 months. I don't know how to monetize models sustainably, however. Maybe crowdfunding? Non-profits? Public utilities like firefighters?
Get a model API — Create an account at GMI Cloud (or any provider like DeepInfra).
Create an API key at their console .
Isolate the environment (warmly recommended, for privacy) — Install Debian in VirtualBox (or any other sandbox you prefer).
Install the agent — Inside the VM, install pi . Modify the hardcoded DEFAULT_MAX_BYTES .
Configure the provider — Edit ~/.pi/agent/models.json to add an OpenAI-compatible provider pointing to the API endpoint.
Run a sanity check — Give the agent a task and check results.
Set up your review workflow - share a copy to the VM, and keep reviewing code between the two copies.
Pelican, the static site/blog engine
The Python programming language
Content
licensed under a Creative Commons Attribution 4.0 International License , except where indicated otherwise.
