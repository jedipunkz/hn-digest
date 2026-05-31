---
source: "https://www.maltebuettner.eu/posts/documentai-bbox-benchmark"
hn_url: "https://news.ycombinator.com/item?id=48335498"
title: "DocumentAI Visual Benchmark - GPT 5.5, Gemini 3.5, Qwen..."
article_title: "# documentai bbox benchmark\n| 418 log"
author: "hegelsmind"
captured_at: "2026-05-31T00:30:56Z"
capture_tool: "hn-digest"
hn_id: 48335498
score: 1
comments: 0
posted_at: "2026-05-30T12:32:36Z"
tags:
  - hacker-news
  - translated
---

# DocumentAI Visual Benchmark - GPT 5.5, Gemini 3.5, Qwen...

- HN: [48335498](https://news.ycombinator.com/item?id=48335498)
- Source: [www.maltebuettner.eu](https://www.maltebuettner.eu/posts/documentai-bbox-benchmark)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T12:32:36Z

## Translation

タイトル: DocumentAI ビジュアル ベンチマーク - GPT 5.5、Gemini 3.5、Qwen...
記事タイトル：#documentai bbox ベンチマーク
| 418ログ
説明: 小規模ドキュメント抽出ベンチマークからのページレベルの例

記事本文:
# ドキュメント bbox ベンチマーク
| 418ログ
前回の投稿 では、DocumentAI 分野の最近の発展について少しお話しました。ここからは実践的な部分です。 ExtractBench データセットの tention v3 論文では、ExtractBench は抽出のみに焦点を当てましたが、モデルが返す境界ボックス参照にも興味があります。
ExtractBench には非常に限られたモデルの選択しかなく、その中にオープンウェイトモデルがなかったため、特に Qwen、Kimi、Mistral がどの程度うまく機能しているかを確認するために、OpenRouter 経由でいくつかの抽出を実行しました。そこで、FlashAttendant-3 のサンプルからページ 1 と 13 を取得し、pdfplumber (ネイティブ PDF) を参照として「参照」境界ボックスを追加しました。完璧ではありませんが、大まかな目安としては十分です。
* 一部のモデルでは、1 回の実行で抽出と bbox を生成できませんでした。これらについては、個別の抽出 + bbox プロンプトを実行しました。
注: 一部のモデルでは、数回実行した後でも OpenRouter で一貫したスコアを取得できませんでした。
bbox スコアは、カバレッジ (bbox が生成されたフィールドの数?)、和集合上の交差 (bbox が元の bbox にどの程度「適合」するかをチェックするため、Jaccard インデックスとも呼ばれます)、および重心距離 (bbox がおおよそ正しい領域にあるかどうかを確認するため) で少しオーバーエンジニアリングされています。
カバー率 × ( 0.5 × 平均 IoU + 0.5 × セントロイド スコア ) カバレッジ \times (0.5 \times \text{平均 IoU} + 0.5 \times \text{セントロイド スコア}) カバー率 × ( 0.5 × 平均 IoU + 0.5 × セントロイド スコア )
ONE_SHOT_SYSTEM_PROMPT = "指定された JSON スキーマに一致する有効な JSON のみを返します。"
one_shot_user_prompt = f"""
提供されたページ画像のみを使用してください。必ずしも連続したページである必要はありません。
オリジナルの PDF は 22 ページあります。スキーマがnumber_of_pagesを要求する場合は、22を使用します。
ページ

マッピング:
- 入力画像 1: 元の PDF ページ 1、page_index 0
- 入力画像 2: 元の PDF ページ 13、page_index 12
各スカラー抽出フィールドは、値と bbox を持つオブジェクトです。次の場合に bbox null を使用します。
値は、提供されたページ画像には表示されません。ボックスは [x1, y1, x2, y2] です。
JSON スキーマ:
{ annotated_extraction_schema_json }
「」
元の JSON スキーマを少し変更し、すべての値に bbox フィールドを追加しました。 ids フィールドの例を参照してください。
{
"id" : {
「値」: {
"タイプ" : [ "文字列" , "null" ]
} 、
"bbox" : {
"タイプ" : [ "オブジェクト" , "null" ] ,
"プロパティ" : {
"ページインデックス" : {
"型" : "整数"
} 、
「ボックス」: {
"タイプ" : "配列" ,
「アイテム」: {
"タイプ" : "数値"
} 、
"minItems" : 4 、
"最大アイテム数" : 4
}
} 、
"必須" : [ "ページインデックス" , "ボックス" ]
}
}
}
1ページ目

## Original Extract

page-level examples from a small document extraction benchmark

# documentai bbox benchmark
| 418 log
In my previous post , I talked a bit about the recent developments in the field of DocumentAI. Now comes the practical part. For the Attention v3 paper from the ExtractBench dataset, ExtractBench focused only on extraction, but I am also interested in the bounding box reference that the models return.
Because ExtractBench had only a very limited selection of models without any open-weight ones among them, I ran a few extractions via OpenRouter especially to see how well Qwen, Kimi, and Mistral are doing. So I took pages 1 and 13 from the FlashAttention-3 example from there and added "reference" bounding boxes with pdfplumber (it is a native PDF) as a reference. They are not perfect, but for a rough indication they are more than enough.
* for some models I did not manage to generate extraction and bbox in one run. For these I ran separate extraction + bbox prompts.
Note: For some models I could not really get consistent scores on OpenRouter even after several runs.
The bbox score is a bit over-engineered with coverage (for how many fields were bboxes generated?), intersection-over-union (to check how well the bbox "fits" the original one, also known as the Jaccard index ), and centroid distance (to check if the bbox is roughly in the correct area):
c o v e r a g e × ( 0.5 × mean IoU + 0.5 × centroid score ) coverage \times (0.5 \times \text{mean IoU} + 0.5 \times \text{centroid score}) co v er a g e × ( 0.5 × mean IoU + 0.5 × centroid score )
ONE_SHOT_SYSTEM_PROMPT = "Return only valid JSON matching the provided JSON Schema."
one_shot_user_prompt = f"""
Only use the provided page images. They are not necessarily consecutive pages.
The original PDF has 22 pages. If the schema asks for number_of_pages, use 22.
Page mapping:
- input image 1: original PDF page 1, page_index 0
- input image 2: original PDF page 13, page_index 12
Each scalar extraction field is an object with value and bbox. Use bbox null when
the value is not visible in the provided page images. Boxes are [x1, y1, x2, y2].
JSON Schema:
{ annotated_extraction_schema_json }
"""
I modified the original JSON schema a bit and added an additional bbox field to every value. See the example for the ids field:
{
"ids" : {
"value" : {
"type" : [ "string" , "null" ]
} ,
"bbox" : {
"type" : [ "object" , "null" ] ,
"properties" : {
"page_index" : {
"type" : "integer"
} ,
"box" : {
"type" : "array" ,
"items" : {
"type" : "number"
} ,
"minItems" : 4 ,
"maxItems" : 4
}
} ,
"required" : [ "page_index" , "box" ]
}
}
}
page 1
