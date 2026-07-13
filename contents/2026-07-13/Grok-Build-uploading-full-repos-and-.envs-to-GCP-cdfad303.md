---
source: "https://twitter.com/xbtoshi/status/2076338252051841512"
hn_url: "https://news.ycombinator.com/item?id=48899226"
title: "Grok Build uploading full repos and .envs to GCP"
article_title: "CyberSatoshi 𓆙 on X: \"the absolute state of ai dev tools. grok build is silently dumping 12gb of untouched repo data and full git commit histories to gcp just to autocomplete a script. they don't want to help you build, they are just treating local dev environments like an open buffet for training h\n[truncated]"
author: "ashleypeacock"
captured_at: "2026-07-13T21:45:08Z"
capture_tool: "hn-digest"
hn_id: 48899226
score: 2
comments: 0
posted_at: "2026-07-13T21:34:21Z"
tags:
  - hacker-news
  - translated
---

# Grok Build uploading full repos and .envs to GCP

- HN: [48899226](https://news.ycombinator.com/item?id=48899226)
- Source: [twitter.com](https://twitter.com/xbtoshi/status/2076338252051841512)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T21:34:21Z

## Translation

タイトル: Grok Build による完全なリポジトリと .env の GCP へのアップロード
記事のタイトル: CyberSatoshi 𓆙 on X: 「AI 開発ツールの絶対的な状態。grok ビルドは、スクリプトをオートコンプリートするためだけに、12 GB の未加工のリポジトリ データと完全な git コミット履歴を黙って gcp にダンプしています。彼らはビルドを手伝うつもりはなく、ローカルの開発環境をトレーニングのためのオープンビュッフェのように扱っているだけです。」
[切り捨てられた]
説明: AI 開発ツールの絶対的な状態。 grok ビルドは、スクリプトをオートコンプリートするためだけに、12 GB の未加工のリポジトリ データと完全な git コミット履歴をサイレントに gcp にダンプします。彼らは構築を手伝いたいわけではなく、ローカルの開発環境をトレーニング用のオープンビュッフェのように扱っているだけです

記事本文:
CyberSatoshi 𓆙 on X: 「AI 開発ツールの絶対的な状態。grok ビルドは、スクリプトをオートコンプリートするためだけに、12 GB の未加工のリポジトリ データと完全な git コミット履歴を黙って gcp にダンプしています。彼らはビルドを手伝いたいわけではなく、ローカルの開発環境をトレーニング用のオープンビュッフェのように扱っているだけです https://t.co/v2pGHKuNlR」 / X Post
@XBToshi AI 開発ツールの絶対的な状態。 grok ビルドは、スクリプトをオートコンプリートするためだけに、12 GB の未加工のリポジトリ データと完全な git コミット履歴をサイレントに gcp にダンプします。彼らは構築を手伝いたいのではなく、ローカルの開発環境をトレーニング データ用のオープンビュッフェのように扱っているだけです。これを機密スタックで実行すると、公開されているかどうかに関係なく、リポジトリ全体がすでに侵害されています。文字通りのスパイウェア。
ソースコードをクラウドに配布するだけでは十分にダメです。バックグラウンド同期で .env.local と .dev.vars を盲目的に取り込むのは完全な怠慢です。彼らは、オートコンプリート モデルを強化するためだけに、生の API キー、データベース認証情報、プライベート ノードを直接 gcp にバキュームしています。
開発ツールを装った大規模な認証情報侵害。これをローカルで実行した場合、秘密鍵はリモート サーバー上に存在します。焼かれたすべての秘密を考慮し、ベアメタルを完全に侵害されたものとして扱い、スタック全体を直ちにローテーションします。
@ elonmusk さん、ユーザーには真剣な説明が必要です。 蓝点网 @landiantech Jul 12 🚨🚨🚨 SpaceXAI の人工知能コード ツール #GrokBuild は、ツール自体は取得されていないコードや調整用のコンテキスト テキストと Git 完全引合履歴を含む完全な Git パブリック パッケージに公開されています。
調査では、12GB のデータが少なくとも 5GB のデータを谷歌云端末に送信し、Grok Build がこれらのデータを保存するために谷歌 GCP を使用していることも示しています。 1 241 1 。 8K 1.8K 5

3 6 536 107 件の返信を読む X は初めてですか?
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

the absolute state of ai dev tools. grok build is silently dumping 12gb of untouched repo data and full git commit histories to gcp just to autocomplete a script. they don't want to help you build, they are just treating local dev environments like an open buffet for training

CyberSatoshi 𓆙 on X: "the absolute state of ai dev tools. grok build is silently dumping 12gb of untouched repo data and full git commit histories to gcp just to autocomplete a script. they don't want to help you build, they are just treating local dev environments like an open buffet for training https://t.co/v2pGHKuNlR" / X Post
@XBToshi the absolute state of ai dev tools. grok build is silently dumping 12gb of untouched repo data and full git commit histories to gcp just to autocomplete a script. they don't want to help you build, they are just treating local dev environments like an open buffet for training data. if you run this on a sensitive stack, your entire repo is already compromised regardless if it's public or not. literal spyware.
shipping source code to the cloud is bad enough. blindly inhaling .env.local and .dev.vars in a background sync is absolute negligence. they are vacuuming up your raw api keys, database credentials, and private nodes directly to gcp just to power an autocomplete model.
a massive credential breach disguised as a dev tool. if you ran this locally, your private keys are now sitting on a remote server. consider every secret burned, treat your bare metal as completely compromised, and rotate your entire stack immediately.
@ elonmusk , your users deserve a serious explanation! 蓝点网 @landiantech Jul 12 🚨🚨🚨 SpaceXAI 的人工智能编码工具 #GrokBuild 被曝默认上传完整的 Git 仓库，包含工具本身没有读取的代码或调用的上下文以及 Git 完整提交历史。
测试还显示 12GB 的仓库数据被发送至少 5GB 的数据到谷歌云端，Grok Build 使用谷歌 GCP 来存储收集的这些数据。 Show more 4:09 PM · Jul 12, 2026 783.7K Views 1 0 7 107 2 4 1 241 1 . 8 K 1.8K 5 3 6 536 Read 107 replies New to X?
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
