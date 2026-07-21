---
source: "https://upload.sortedreceipts.com/demo"
hn_url: "https://news.ycombinator.com/item?id=48999423"
title: "Show HN: Sorted Receipts - clients dump receipts in one link, LLM sorts them"
article_title: "Try the client upload link — Sorted Receipts"
author: "jahnoikka"
captured_at: "2026-07-21T23:49:02Z"
capture_tool: "hn-digest"
hn_id: 48999423
score: 1
comments: 0
posted_at: "2026-07-21T22:51:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sorted Receipts - clients dump receipts in one link, LLM sorts them

- HN: [48999423](https://news.ycombinator.com/item?id=48999423)
- Source: [upload.sortedreceipts.com](https://upload.sortedreceipts.com/demo)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T22:51:01Z

## Translation

タイトル: HN の表示: ソートされた領収書 - クライアントは領収書を 1 つのリンクにダンプし、LLM がそれらをソートします
記事のタイトル: クライアントのアップロード リンクを試す — ソートされた領収書
HN text: ソロ創業者。 Xero は 5 月に Hubdoc を退職しました。追放されたユーザーは、ほとんどが 15 ～ 30 人の顧客を持つ個人の簿記係であり、その代わりのオプションはエンタープライズ スイート (Dext) またはドキュメントごとのクレジット (AutoEntry) です。ソートされた領収書は、彼らが求め続けているものの退屈なバージョンです。各クライアントは 1 つのアップロード リンクを取得します。アカウントもアプリも必要なく、メールをリンクに転送することもできます。ファイルは月/ベンダー/タイプごとに分類され (OpenRouter 経由の LLM)、チェックを入れると 1 つのレビュー キューに配置され、不足しているものは簿記担当者自身の名前で締切日に追跡され、月末には CSV エクスポートと一貫した名前変更が行われます。デモは、シードされたデータを含む実際のアプリであり、サインアップは必要ありません。スタック: 1 つの小規模 VPS 上の FastAPI + SQLite。既知の荒削りな点: QBO/Xero の直接投稿はありません (エクスポートのみ)、分類により手書きの領収書が破損する場合があります、レビュー UI はデスクトップファーストです。私は、レビューキューの UX と、そのようなものがレシートの最も乱雑な靴箱の中でも生き残れるかどうかについての批判を心から重視します。

記事本文:
仕分け済み ✓ 領収書
公開デモ
私は
クライアントリンクの例。これは、簿記担当者がクライアントに送信するプライベート リンクです。
この例は Acme Coffee 用であり、アカウントやパスワードは必要ありません。
アップロード.sortedreceipts.com/demo
Acme Coffee の領収書をお送りください。
携帯電話の写真または PDF をそのままドロップしてください。ソートされた領収書は山を読み取り、簿記担当者にきれいな月末を提供します。
3 つのサンプル領収書から始めます
これらは生成されたデモ ドキュメントなので、何もアップロードせずに全体を確認できます。
個々のドキュメントまたはフォルダー全体を選択します。公開デモでは、サポートされている最初の 4 つのレシートが取得され、残りは明らかに除外されます。
この公開デモでは最大 4 つのドキュメントを使用できます。余分なファイルはデバイス上に残ります。
機密記録を公開デモに含めないようにしてください。アップロードされたファイルは、この 1 つの結果のために AI プロセッサに送信されます。 Sorted Receipts は応答後にその一時コピーを削除し、アカウントに保存することはありません。 AI ルーティングでは、送信されたデータを収集する可能性のあるプロバイダーは除外されます。独自のサンプルをアップロードしたくない場合は、生成されたサンプルを使用してください。
サンプルの月末が並べ替えられました。
ベンダーまたは金額が間違っていますか?その行で「修正」を使用します。修正はルールとして保存されます。同じファイルを再度並べ替えて、それが保持されるのを確認します。
同じ山を反対側から見ると、簿記担当者が作業するダッシュボードです。各ドキュメントにチェックを入れ、月の合計を確認し、きれいな月末をエクスポートします。パイロット版では、これは簿記担当者のプライベート アクセスの背後にあり、月ごとの ZIP と不足している文書のリマインダーの下書きが含まれます。
すべての修正がルールになります。ルールは、まったく同じファイルと、同じベンダーからの今後のドキュメントに再適用され、そのメモは各種類に沿って読み取りをガイドします。ルールはこのブラウザにのみ保存されます。

## Original Extract

Solo founder. Xero retired Hubdoc in May; the displaced users are mostly solo bookkeepers with 15-30 clients whose replacement options are an enterprise suite (Dext) or per-document credits (AutoEntry). Sorted Receipts is the boring version of what they keep asking for: each client gets one upload link. No account, no app, and forwarding an email to the link works too. Files get classified by month/vendor/type (LLM via OpenRouter), land in one review queue you check off, whatever's missing gets chased on your cutoff date under the bookkeeper's own name, and month-end is a CSV export plus consistently renamed files. The demo is the real app with seeded data, no signup. Stack: FastAPI + SQLite on one small VPS. Known rough edges: no direct QBO/Xero posting (export only), classification occasionally mangles handwritten receipts, and the review UI is desktop-first. I'd genuinely value critique of the review-queue UX and whether the sort would survive your messiest shoebox of receipts.

Sorted ✓ Receipts
Public demo
i
Example client link. This is the private link a bookkeeper sends their client.
This example is for Acme Coffee and needs no account or password.
upload.sortedreceipts.com/demo
Send Acme Coffee's receipts.
Drop in phone photos or PDFs exactly as they are. Sorted Receipts reads the pile and gives your bookkeeper a clean month-end.
Start with three sample receipts
They're generated demo documents, so you can see the whole sort without uploading anything.
Choose individual documents or a whole folder. The public demo takes the first four supported receipts and clearly leaves the rest out.
Up to 4 documents in this public demo. Extra files stay on your device.
Keep sensitive records out of a public demo. Uploaded files are sent to our AI processor for this one result. Sorted Receipts deletes its temporary copy after the response and never saves it to an account. AI routing excludes providers that may collect submitted data. Use the generated samples if you'd rather not upload your own.
Your sample month-end is sorted.
Wrong vendor or amount? Use Fix on that row. The correction is saved as a rule — sort the same file again and watch it hold.
The same pile from the other side: the dashboard your bookkeeper works from. They check off each document, see month totals, and export a clean month-end. In the pilot this sits behind the bookkeeper's private access, with per-month ZIPs and reminder drafts for missing documents.
Every fix becomes a rule. A rule re-applies to the exact same file and to future documents from the same vendor, and its note rides along with each sort to guide the reading. Rules are saved in this browser only.
