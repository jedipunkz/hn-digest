---
source: "https://twitter.com/i/status/2076689215258014069"
hn_url: "https://news.ycombinator.com/item?id=48904164"
title: "Grok Build CLI accused of uploading Git repos to a Google Cloud bucket"
article_title: "International Cyber Digest on X: \"‼️ BREAKING: xAI's Grok Build CLI was uploading entire Git repositories to a Google Cloud bucket, private codebases and unredacted secrets included. The uploads quietly stopped via a hidden server-side flag, and xAI still has not said a word about scope, retention,\n[truncated]"
author: "maxloh"
captured_at: "2026-07-14T09:47:21Z"
capture_tool: "hn-digest"
hn_id: 48904164
score: 1
comments: 0
posted_at: "2026-07-14T09:28:11Z"
tags:
  - hacker-news
  - translated
---

# Grok Build CLI accused of uploading Git repos to a Google Cloud bucket

- HN: [48904164](https://news.ycombinator.com/item?id=48904164)
- Source: [twitter.com](https://twitter.com/i/status/2076689215258014069)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T09:28:11Z

## Translation

タイトル: Grok Build CLI が Git リポジトリを Google Cloud バケットにアップロードしたとして非難される
記事タイトル: International Cyber Digest on X: 「‼️ 速報: xAI の Grok Build CLI は、プライベート コードベースと編集されていないシークレットを含む Git リポジトリ全体を Google Cloud バケットにアップロードしていました。アップロードはサーバー側の非表示フラグによって静かに停止されましたが、xAI はまだ範囲、保持、
[切り捨てられた]
説明: ‼️ 速報: xAI の Grok Build CLI は、プライベート コードベースと未編集のシークレットを含む Git リポジトリ全体を Google Cloud バケットにアップロードしていました。アップロードはサーバー側の非表示フラグによって静かに停止されましたが、xAI はまだ範囲、保持、削除について何も述べていません。

記事本文:
X に関するインターナショナル サイバー ダイジェスト: 「‼️ 速報: xAI の Grok Build CLI は、プライベート コードベースと編集されていないシークレットを含む Git リポジトリ全体を Google Cloud バケットにアップロードしていました。アップロードはサーバー側の非表示フラグによって静かに停止されましたが、xAI は範囲、保持、削除についてまだ何も述べていません。 https://t.co/B2iGaPRVZq」 / X Post
@IntCyberDigest ‼️速報: xAI の Grok Build CLI は、プライベート コードベースと未編集のシークレットを含む Git リポジトリ全体を Google Cloud バケットにアップロードしていました。アップロードはサーバー側の非表示フラグによって静かに停止されましたが、xAI はまだ範囲、保持、削除について何も述べていません。
そのスケールは驚異的です。 12 GB のテスト リポジトリでは、xAI の grok-code-session-traces バケットに 5.1 GB が飛び出しましたが、実際のコーディング タスクに必要な容量はわずか 192 KB でした。このツールは、必要なファイルではなく、実行中のリポジトリを取得しました。
この修正は、研究者によるワイヤーレベルの分析の翌日に、隠しフラグ disable_codebase_upload: true として到着しました。 「モデルの改善」のオプトアウトによってアップロードが停止されることはありませんでした。
すでにアップロードされたコードが削除されるかどうかについては、まだ勧告も範囲も何もありません。 AI コーディング エージェントに独自のコードを指示する人にとっては、設定ページに記載されている内容よりも、ネットワークを通過する内容の方が重要です。 3:24 PM · 2026 年 7 月 13 日 150 万回の再生数 3 9 4 394 9 7 4 974 7 。 5K 7.5K 2． 1 K 2.1K 394 件の返信を読む X は初めてですか?
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

‼️ BREAKING: xAI's Grok Build CLI was uploading entire Git repositories to a Google Cloud bucket, private codebases and unredacted secrets included. The uploads quietly stopped via a hidden server-side flag, and xAI still has not said a word about scope, retention, or deletion.

International Cyber Digest on X: "‼️ BREAKING: xAI's Grok Build CLI was uploading entire Git repositories to a Google Cloud bucket, private codebases and unredacted secrets included. The uploads quietly stopped via a hidden server-side flag, and xAI still has not said a word about scope, retention, or deletion. https://t.co/B2iGaPRVZq" / X Post
@IntCyberDigest ‼️ BREAKING: xAI's Grok Build CLI was uploading entire Git repositories to a Google Cloud bucket, private codebases and unredacted secrets included. The uploads quietly stopped via a hidden server-side flag, and xAI still has not said a word about scope, retention, or deletion.
The scale is staggering. On a 12 GB test repo, 5.1 GB flew out the door to xAI's grok-code-session-traces bucket while the actual coding task needed just 192 KB. The tool grabbed whatever repository it ran in, not the files it needed.
The fix arrived as a hidden flag, disable_codebase_upload: true, a day after a researcher's wire-level analysis. The "Improve the model" opt-out never stopped the uploads.
Still no advisory, no scope, no word on whether already-uploaded code gets deleted. For anyone pointing AI coding agents at proprietary code, what crosses the wire matters more than what the settings page says. 3:24 PM · Jul 13, 2026 1.5M Views 3 9 4 394 9 7 4 974 7 . 5 K 7.5K 2 . 1 K 2.1K Read 394 replies New to X?
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
