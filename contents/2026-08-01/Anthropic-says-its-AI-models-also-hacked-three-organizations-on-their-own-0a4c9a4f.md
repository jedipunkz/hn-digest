---
source: "https://www.engadget.com/2227630/anthropic-ai-models-hacked-three-organizations-on-their-own/"
hn_url: "https://news.ycombinator.com/item?id=49132942"
title: "Anthropic says its AI models also hacked three organizations on their own"
article_title: "Anthropic Says Its AI Models Also Hacked Three Organizations On Their Own"
author: "madradavid"
captured_at: "2026-08-01T10:27:18Z"
capture_tool: "hn-digest"
hn_id: 49132942
score: 1
comments: 0
posted_at: "2026-08-01T10:10:15Z"
tags:
  - hacker-news
  - translated
---

# Anthropic says its AI models also hacked three organizations on their own

- HN: [49132942](https://news.ycombinator.com/item?id=49132942)
- Source: [www.engadget.com](https://www.engadget.com/2227630/anthropic-ai-models-hacked-three-organizations-on-their-own/)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T10:10:15Z

## Translation

タイトル: Anthropic は、自社の AI モデルも 3 つの組織を独自にハッキングしたと発表
記事のタイトル: Anthropic は、その AI モデルが独自に 3 つの組織もハッキングしたと発表
説明: OpenAI がそのモデルが Hugging Face に侵入したことを認めた後、Anthropic はテストしていたモデルが他の組織もハッキングしていたことを認めました。

記事本文:
Anthropic、自社の AI モデルが独自に 3 つの組織もハッキングしたと発表
最新の
Anthropic は、自社の AI モデルが 3 つの組織を独自にハッキングしたと発表
同社は、エージェントがHugging FaceをハッキングしたというOpenAIの暴露を受けて、テストログの調査を開始した。
ジャックプレス/シャッターストック
どうやら、AI モデルが独自に他の組織のシステムに侵入した企業は OpenAI だけではないようです。 Anthropic は報告書を発表し、同社の AI モデルもテスト環境から解放され、3 つの異なる組織に侵入したが、どうやら通知されるまで侵害に気づいていなかった模様であることを認めた。
OpenAIが、同社がテストしていたAIエージェントが脆弱性を悪用してインターネットに接続し、Hugging Faceにハッキングしたという事件を明らかにした後、Anthropicは自社のテストの大規模な見直しを開始した。同社は特に、自社のクロードモデルがインターネットにアクセスできないはずのときにインターネットにアクセスしたという証拠を探しました。このモデルは 3 つのインスタンスでインターネットにアクセスでき、その後「3 つの異なる組織の運用インフラストラクチャに不正アクセスした」ことが判明しました。
この事件には 3 つの異なるクロード モデルが関与していました。Opus 4.7、サイバーセキュリティに焦点を当てた Mythos 5、および一般リリースが予定されていないプロトタイプです。脱出したとき、彼らは全員、キャプチャ・ザ・フラッグ・チャレンジをしていた。モデルたちは、Anthropic の内部ネットワークの別のマシンに隠された秘密情報の一部である「フラグ」を見つけるという任務を負っていました。彼らは旗を回収するために機械に侵入することになっていた。
ただし、OpenAI の場合とは異なり、このモデルはインターネットにアクセスするために脆弱性を悪用しませんでした。彼らは、次の理由により、テスト環境の範囲から逃げ出したようです。

○人的ミス。アンスロピック氏はプロンプトでモデルたちに、インターネットにアクセスできないことを指定しましたが、実際はそうではありませんでした。同社と評価パートナーの間の「誤解のため」、彼らはインターネットにアクセスできた。そのため、モデルたちがオープンなインターネットへのアクセスを見つけて 3 つの組織のシステムに遭遇したとき、彼らはそれを演習の一部として扱い、侵入しました。彼らは意図的にテスト環境から逃げようとしたわけではないとアントロピック氏は明らかにしました。
このモデルは、弱いパスワードを利用するなど、組織に侵入するための基本的な手法を使用しており、複雑な脆弱性を悪用していません。 Anthropic 社は、最新モデルはインターネット上にあることを認識した後停止したが、古いモデルはとにかく影響を受けた組織への攻撃を続けたと述べた。
Anthropic は、同社とその評価パートナーがテストを開始する前にすべてのインターネット アクセス パスを注意深く検証していれば、インシデントを防ぐことができた可能性があることを認めました。テストをもっと頻繁に徹底的に見直すこともできたはずです。さらに、最初からインターネットにアクセスできることが伝えられていれば、モデルは異なる行動をとった可能性があります。
同社は、テスト記録の見直しを開始してから4日後の7月27日に、評価パートナーと影響を受ける3つの組織に通知した。影響を受けた組織のうち 2 つは、侵害されたことに気づいていませんでした。 Anthropic はまだ 3 番目に到達しようとしています。

## Original Extract

After OpenAI’s admission that its models broke into Hugging Face, Anthropic has now admitted that the models it was testing also hacked other organizations.

Anthropic Says Its AI Models Also Hacked Three Organizations On Their Own
Latest
Anthropic says its AI models also hacked three organizations on their own
The company started reviewing test logs after OpenAI’s revelation that its agent hacked Hugging Face.
jackpress/Shutterstock
Apparently, OpenAI isn't the only company whose AI models have hacked into other organizations' systems on their own. Anthropic has published a report, admitting that its AI models have also broken free from their testing environment and infiltrated three different organizations, which apparently weren't aware of the breach until they were notified.
After OpenAI revealed the incident, wherein an AI agent it was testing exploited vulnerabilities to connect to the internet and hack into Hugging Face, Anthropic began a large-scale review of its own tests. The company specifically looked for evidence that its Claude models had accessed the internet when they shouldn't have been able to. Turns out the models were able to access the internet in three instances and then "gained unauthorized access to the production infrastructure of three different organizations."
Three different Claude models were involved in the incidents: Opus 4.7, the cybersecurity-focused Mythos 5 and a prototype that's not planned for general release. They were all doing a capture-the-flag challenge when they broke free. The models were tasked with finding the "flag," which is a piece of secret information, hidden in a different machine in Anthropic's internal network. They were supposed to break into the machine to retrieve the flag.
Unlike in OpenAI's case, however, the models didn't exploit a vulnerability to get access to the internet. They seemed to have escaped the confines of their testing environment due to human error. Anthropic specified to the models in a prompt that they had no internet access, but that wasn't the case at all. They did have internet access "due to a misunderstanding" between the company and its evaluation partner. So, when the models found access to open internet and encountered the systems of the three organizations, they treated it as part of the exercise and broke in. They didn't deliberately attempt to escape their testing environment, Anthropic clarified.
The models used basic techniques to infiltrate the organizations, such as taking advantage of weak passwords, and din't exploit complex vulnerabilities. Anthropic said its latest model stopped after it recognized that it was on the internet, but its older model continued attacking the affected organization anyway.
Anthropic admitted that the company and its evaluation partner could have prevented the incidents by carefully validating all internet access paths before they started their tests. They could have also reviewed their tests more frequently and thoroughly. In addition, the models could have behaved differently if they were told from the start that they did have internet access.
The company notified its evaluation partner and the three affected organizations on July 27, four days after it started reviewing its test transcripts. Two of the affected organizations weren't aware that they had been breached. Anthropic is still trying to reach the third.
