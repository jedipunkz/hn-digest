---
source: "https://www.lesswrong.com/posts/yrbyyvFvuaGfRAtB7/evaluating-offline-monitoring-of-internal-ai-agents-2"
hn_url: "https://news.ycombinator.com/item?id=48709299"
title: "Evaluating Offline Monitoring of Internal AI Agents"
article_title: "Evaluating Offline Monitoring of Internal AI Agents — LessWrong"
author: "joozio"
captured_at: "2026-06-28T17:25:57Z"
capture_tool: "hn-digest"
hn_id: 48709299
score: 1
comments: 0
posted_at: "2026-06-28T17:07:35Z"
tags:
  - hacker-news
  - translated
---

# Evaluating Offline Monitoring of Internal AI Agents

- HN: [48709299](https://news.ycombinator.com/item?id=48709299)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/yrbyyvFvuaGfRAtB7/evaluating-offline-monitoring-of-internal-ai-agents-2)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T17:07:35Z

## Translation

タイトル: 内部 AI エージェントのオフライン監視の評価
記事のタイトル: 内部 AI エージェントのオフライン監視の評価 — LessWrong
説明: この作業は、GovAI Winter Fellowship 2026 中に実施されました。完全なレポートは…

記事本文:
ログイン 内部 AI エージェントのオフライン監視を評価する — LessWrong AI Control AI Personal Blog 8
内部 AI エージェントのオフライン監視の評価
これは https://www.governance.ai/research-paper/evaluating-offline-monitoring-of-internal-ai-agents のリンクポストです。この作業は、GovAI Winter Fellowship 2026 中に実施されました。レポート全文
フロンティア AI 企業は、オフライン監視を使用して、社内に導入された AI エージェントからのリスクに対処します。 AI 開発者は、安全性の研究やモデルのトレーニングなどの社内作業で AI エージェントへの依存度を高めています。同時に、これらの企業は、調整されていないモデルがこのアクセスを悪用して、AI によってもたらされるリスクを理解する取り組みを妨害するなどの関連行動を起こす可能性があることを懸念しています。このようなインスタンスを特定するために、AI 企業は、AI エージェントの動作の記録をレビューし、疑わしいアクティビティにフラグを立てる「モニター」と呼ばれる別の AI モデルを備えています。人間のレビュー担当者は、モニターによって不審とフラグが付けられたアクティビティを検査し、そのアクティビティが懸念があるかどうかを判断し、適切な対応を決定します。この監視はオフラインで行われます。つまり、アクションはリアルタイムで傍受されるのではなく、実行後にレビューされます。
企業は現在、合成攻撃によるオフライン監視の有効性を評価しています。オフライン監視の有効性を評価するために、OpenAI と Anthropic は合成攻撃 (配置のずれた AI が展開中に実行する可能性のある有害なアクションを含むように構築されたトランスクリプト) を使用し、モニターが警告を発するかどうかをチェックします。
有効性の評価に関する現在の報告は不十分です。 Anthropic と OpenAI によって現在公開されている情報を考慮すると、外部関係者はオフライン監視の全体的な有効性を評価できません。 Anthropic は、m の頻度を報告します。

監視者は合成攻撃のトランスクリプトにフラグを立てますが、それらのトランスクリプトがパイプライン全体を生き残り、人間によるレビューに到達する頻度は知りません。後のフィルタでは、最初にフラグが立てられたトランスクリプトが破棄される可能性があります。 OpenAI の報告によると、同社のモニターは人間の従業員によって個別にフラグが付けられたすべての記録にフラグを立てていますが、従業員が明らかな不正行為のケースのみを発見した場合、このことはほとんど役に立ちません。私たちが知る限り、オフライン監視に関する情報を公開しているのは Anthropic と OpenAI だけです。
推奨事項。外部の関係者がオフライン監視の有効性を評価して、社内に導入されたモデルのリスクに対処できるようにするために、企業は次のことを推奨します。 さまざまな導入タイプにわたって監視されているモデル出力の割合を報告します。
監視パイプラインによって人間によるレビューに送信される合成攻撃トランスクリプトの割合を報告します。
人間のレビューに送信され、レビュー担当者によってエスカレーションされた合成攻撃トランスクリプトの割合を報告します。
攻撃とその検出の間のタイムラグを報告します。
推奨事項 (2) および (3) の推定値は、合成攻撃が実際の攻撃に似ている場合にのみ参考になるため、合成攻撃が仮想の実際の攻撃と類似していることを確認してください。
0 コメント 0 AI 制御 AI 個人ブログ 8

## Original Extract

This work was conducted during the GovAI Winter Fellowship 2026. Full report …

Login Evaluating Offline Monitoring of Internal AI Agents — LessWrong AI Control AI Personal Blog 8
Evaluating Offline Monitoring of Internal AI Agents
This is a linkpost for https://www.governance.ai/research-paper/evaluating-offline-monitoring-of-internal-ai-agents This work was conducted during the GovAI Winter Fellowship 2026. Full report
Frontier AI companies use offline monitoring to address risks from internally deployed AI agents. AI developers increasingly rely on AI agents for internal work, including for safety research and model training. At the same time, these companies are concerned that a misaligned model could exploit this access to take concerning actions, such as sabotaging efforts to understand the risks posed by AI. To identify such instances, AI companies have separate AI models called "monitors" that review transcripts of AI agents' actions and flag suspicious activity. Human reviewers examine activity flagged as suspicious by monitors, judge whether that activity is concerning, and decide on an appropriate response. This monitoring occurs offline, meaning that actions are reviewed after they have been executed rather than intercepted in real time.
Companies currently assess the effectiveness of offline monitoring via synthetic attacks. To assess the effectiveness of offline monitoring, OpenAI and Anthropic use synthetic attacks – transcripts constructed to contain the kind of harmful actions a misaligned AI might take during deployment – and then check whether monitors flag them.
Current reporting on assessments of effectiveness is insufficient. Given the information currently made public by Anthropic and OpenAI, external parties cannot assess the overall effectiveness of their offline monitoring. Anthropic reports how often monitors flag synthetic attack transcripts, but not how often those transcripts survive the full pipeline and reach human review; later filters can discard transcripts that were initially flagged. OpenAI reports that their monitor flags all transcripts independently flagged by human employees, but this tells us little if employees only catch obvious cases of misbehavior. As far as we are aware, only Anthropic and OpenAI have published information about their offline monitoring.
Recommendations. To enable external parties to assess the effectiveness of offline monitoring to address risks from internally deployed models, we recommend that companies: Report the fraction of model outputs that are monitored across different deployment types.
Report the fraction of synthetic attack transcripts that are sent to human review by the monitoring pipeline.
Report the fraction of synthetic attack transcripts sent to human review that are escalated by reviewers.
Report the time lag between an attack and its detection.
Ensure that synthetic attacks are similar to hypothetical real attacks, as the estimates in recommendations (2) and (3) are only informative if synthetic attacks resemble real ones.
0 Comments 0 AI Control AI Personal Blog 8
