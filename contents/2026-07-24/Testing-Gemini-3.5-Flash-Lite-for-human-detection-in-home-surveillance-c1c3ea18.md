---
source: "https://romanuk.org/vlm-models/"
hn_url: "https://news.ycombinator.com/item?id=49036075"
title: "Testing Gemini 3.5 Flash Lite for human detection in home surveillance"
article_title: "Testing Gemini 3.5 Flash Lite for human detection in home surveillance — Martín Romañuk"
author: "mromanuk"
captured_at: "2026-07-24T14:45:25Z"
capture_tool: "hn-digest"
hn_id: 49036075
score: 4
comments: 0
posted_at: "2026-07-24T14:19:51Z"
tags:
  - hacker-news
  - translated
---

# Testing Gemini 3.5 Flash Lite for human detection in home surveillance

- HN: [49036075](https://news.ycombinator.com/item?id=49036075)
- Source: [romanuk.org](https://romanuk.org/vlm-models/)
- Score: 4
- Comments: 0
- Posted: 2026-07-24T14:19:51Z

## Translation

タイトル: 家庭監視における人間検出のための Gemini 3.5 Flash Lite のテスト
記事のタイトル: 家庭監視における人体検出のための Gemini 3.5 Flash Lite のテスト — Martín Romañuk
説明: 家庭監視における人物検出について、Gemini-3.5-Flash-Lite VLM 機能を 3.1 と比較してベンチマークします。

記事本文:
家庭監視における人間検出のための Gemini 3.5 Flash Lite のテスト — Martín Romañuk Martín Romañuk ブログ プロジェクトについて 家庭監視における人間検出のための Gemini 3.5 Flash Lite のテスト
私はフリゲートをベースにしたカスタムの家庭用監視システムを構築しました。 Frigate は人物検出部分を実行し、より詳細な検出のために Casa Segura をフィードします。侵入検知のためのファネルのようなものです。 Frigate のアルゴリズムは TensorFlow Lite や OpenVINO などの機械学習モデルに基づいており、人物検出には MobileNet や YOLO などのモデルを使用します。システムは現在 Gemini-3.1-Flash-Lite を使用しており、うまく機能していますが、誤検知の問題があり、基本的に私の犬がシステムを人として検出してトリガーします。午前3時のアラームは聞きたくないです。それがFPを下げる最大のポイントです。
最近 Gemini-3.5-Flash-Lite が導入されましたが、価格が 50% 高騰していますが、これは奇妙に感じます。同じ階層のモデルは今後も同じ価格設定になると予想されます。おそらく彼らはモデルを大幅に改良し、それが価格高騰の正当化になったのでしょう。モデル紹介の新しいカードをすべて読むのは少し疲れます。
ゼロショットでFPを下げることができるモデルを待っていたので、テストするのが楽しみでした。以前の FP イベント (現実世界) のクリップを使用した小さなベンチマークがあります。現在、ベンチマークは kimi k3 によって管理されています (Claude Opus 4.5 および後継モデルで構築されました)。
18 個のカメラ クリップ、「手動」でフレームごとに検証（人間なしの 10 個 - 昼夜を問わず犬、横たわっている犬、遠く離れた犬、人間ありの 8 個、シーン内の人間 + 犬を含む）。
モデルが犬と人を間違えたとき。
Gemini 3.1 Pro はかなり高価です。 3.1 は 3.5 よりも 40% 安いです。
Gemini 3.1 と 3.5 は非常に似ていますが、3.1 が勝ちます。ジェミニ3

.1 Pro、本当に遅いです。
クリップごとに 3 回実行、プロダクションと同じプロンプトとトランスコード (6 秒、2 fps、768 ピクセル)、温度 0.2。
キミによれば、3.5 の単一の FP の信頼度は 0.90 でした。これはまさに、どのモデルも除去できない種類の尾部幻覚です。これに対して、より高価なモデルが勝つわけではありません。 2 番目のクエリが必要です。
悲しいことに、新しいモデルは私がすでに持っていたモデル (Gemini 3.1 Flash Lite) よりもわずかに悪かったです。誤検知を修正したのは、モデルを切り替えることではなく、同じモデルに再度チェックを依頼することでした。これはハック的で完璧ではなく、遅延が増えると思います。
コメントは管理されており、承認後に表示されます。

## Original Extract

Benchmarking Gemini-3.5-Flash-Lite VLM capabilities against 3.1 for person detection in home surveillance.

Testing Gemini 3.5 Flash Lite for human detection in home surveillance — Martín Romañuk Martín Romañuk About Blog Projects Testing Gemini 3.5 Flash Lite for human detection in home surveillance
I’ve built a custom home surveillance system based around Frigate. Frigate does the person detection part and feeds Casa Segura for finer detection. It’s like a funnel for intrusion detection. Frigate’s algorithm is based on machine learning models such as TensorFlow Lite or OpenVINO, using models like MobileNet or YOLO for person detection. The system is urrently using Gemini-3.1-Flash-Lite which works well, but there is an problem with false positives, basically my dogs triggers the system detected as persons. I don’t want to hear an alarm at 3AM. That’s the main point on lowering the FP .
Recently Gemini-3.5-Flash-Lite was introduced, there is 50% bump in pricing, which I find it odd, I would expect that the same tier of model continue to have the same pricing. Maybe they improved the model a lot, and that’s the justification of the price bump. I’m a bit tired of going through reading every new card of a model introduction.
I was excited to test it because I’m waiting for a model that can lower the FP on zero shot. I’ve a small benchmark with clips from previous FP events, which are real world, currently the benchmark is managed by kimi k3 (it was build with Claude Opus 4.5 and succesors models ).
18 camera clips, verified frame by frame by “hand” (10 without humans — dogs day/night, dog lying down, dog far away; 8 with humans, including person + dogs in scene).
When the model mistakes a dog for a person.
Gemini 3.1 Pro is pretty expensive. 3.1 is still 40% cheaper than 3.5.
Gemini 3.1 and 3.5 are very similar, but 3.1 wins. Gemini 3.1 Pro, is really slow.
3 runs per clip, same prompt and transcode as production (6s, 2 fps, 768px), temperature 0.2.
According to Kimi: The single FP of the 3.5 came with confidence 0.90 — exactly the kind of tail hallucination that no model eliminates. Against that, a more expensive model doesn’t win; a second query is needed.
Sadly the new model was marginally worse than the one I already had (Gemini 3.1 Flash Lite); what fixed the false positives wasn’t switching models, but asking the same model to check again. I feel that is hackish, not perfect and adds latency.
Comments are moderated and appear after approval.
