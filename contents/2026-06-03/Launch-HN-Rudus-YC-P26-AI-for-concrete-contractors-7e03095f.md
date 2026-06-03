---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48374528"
title: "Launch HN: Rudus (YC P26) – AI for concrete contractors"
article_title: ""
author: "rishipankhaniya"
captured_at: "2026-06-03T00:37:32Z"
capture_tool: "hn-digest"
hn_id: 48374528
score: 31
comments: 16
posted_at: "2026-06-02T18:51:09Z"
tags:
  - hacker-news
  - translated
---

# Launch HN: Rudus (YC P26) – AI for concrete contractors

- HN: [48374528](https://news.ycombinator.com/item?id=48374528)
- Score: 31
- Comments: 16
- Posted: 2026-06-02T18:51:09Z

## Translation

タイトル: HN: Rudus (YC P26) を開始 – コンクリート請負業者向け AI
HN テキスト: こんにちは、HN、私たちはリシとサヒルです。当社は、コンクリート下請け業者向けに構築された AI を活用した離陸および推定プラットフォームである Rudus ( https://www.rudus.ai/ ) を開発しました。テイクオフは、具体的な計画シートから材料を測定し、定量化するプロセスです。 Rudus は、すべてのコンクリート構造物 (基礎、壁、柱、スラブ) を識別し、関連する詳細情報を取得し、何時間にもわたる手動の数量計算を排除します。デモはこちらです: https://www.youtube.com/watch?v=PAMNDRWEdlI 。問題: コンクリートの下請け業者はあらゆる建物の屋台骨ですが、その見積もりワークフローは 20 年間変わっていません。現在、上級見積り担当者は PDF を開き、すべてのフーチングと勾配梁を手動でトレースし、体積、型枠、重ねスプライスと展開長を含む鉄筋サイズ別の鉄筋など、300 を超える項目を含む Excel スプレッドシートを手作業で作成しています。入札には数週間、場合によっては数か月かかる場合もあります。ほとんどの企業には見積もり担当者が数名しかいないため、利用可能な仕事のほとんどに物理的に入札することができません。この業界で使用されているソフトウェアは 2020 年以来更新されていません。さらに、市場に出ているすべての AI 数量拾いツールは GC 用に構築されており、コンクリート推定値が実際にどのように機能するかを考慮するのではなく、コンクリートを 1 つのチェックボックスとして扱います。私たちはこのトレード専用の Rudus を構築しています。これを始めたのは、サヒルが建設管理のクラスを受講し、見積もりの​​ワークフローが何十年も変わっていないことに気づいたときでした。私たちはコールドコールを始め、ドーナツを持ってオフィスに入り、現場に現れました。そして誰もが同じことを言いました。見積もりの​​遅さがビジネスを成長させる最大のボトルネックですが、彼らが試した新製品はすべて失敗しました。私たちは、これらのツールが失敗する理由は、

信頼性が高く、頻繁なエラーが後で問題を引き起こす可能性があります。見積もり業者はこれらの数字に数百万ドルから数十億ドルの入札を賭けており、ワークフローをブラックボックスと交換するつもりはないことは明らかです。私たちは別のアプローチを採用しました。つまり、当社の製品を現在の見積もりワークフローに前向きに導入することで、現在のワークフローを置き換えるのではなく、インテリジェントに加速するソフトウェアです。見積り担当者が構造 PDF を Rudus にアップロードすると、すべてのシート (基礎計画、セクション詳細、基礎集計表、フレーム立面図) が自動分類され、それぞれが適切な処理パイプラインにルーティングされます。コンピューター ビジョンは、図面セット全体で具体的な要素を検出し、シート間の相互参照に従って寸法と詳細を解決し、平面図のみのツールでは常に見逃される要素を検出します。各要素は、見積担当者が通常手作業で行うすべての計算を使用して、コンクリート、型枠、鉄筋などの完全な組立ライン項目に展開されます。一般的な基盤パッケージは、少数のアセンブリから 80 ～ 120 の価格の製品まであります。見積もり担当者はレビューし、必要に応じて上書きし、既存のワークフローに直接エクスポートします。 AI 推定の分野では、いくつかの重要な利点があります。 1 つ目は、建設業界のニッチな部分であるコンクリートに焦点を当てていることです。シートが他のサブトレードとは大幅に異なるため、コンクリートサブト用にこれを構築している人は他にいません。同じ理由で、VLM やその他の汎用ソリューションは機能しません。代わりに、大量の顧客データからのトレーニングに依存する、独自のコンピューター ビジョン モデルが必要です。当社では、お客様のテイクオフで直接トレーニングされた複数の異なるモデルを実行しており、お客様と当社のモデルとのやり取りはすべてトレーニング サンプルとなり、使用するにつれてクライアントごとの精度が向上します。当社の 2 番目の利点は当社の製品にあります

私たちはブラックボックスではなく副操縦士を構築することを選択したため、この方法論を採用しました。ほとんどの AI 離陸プラットフォームは、自律的に数量を生成することで推定器を完全に置き換えようとしていますが、現在のモデルでは出力の品質が低いため、いずれにしても離陸は手作業でやり直されます。構造コンクリート推定器のある部屋に 100 時間以上座り込み、自分たちで多数のテイクオフを完了した後、私たちは実際のワークフローを中心に構築しました。エスティメーターが離陸を開始し、ルーダスが w を延長します。

[切り捨てられた]

## Original Extract

Hi HN, we’re Rishi and Sahil. We’ve developed Rudus ( https://www.rudus.ai/ ), an AI-powered takeoff and estimation platform built for concrete subcontractors. Takeoff is the process of measuring and quantifying materials from concrete plan sheets. Rudus identifies every concrete structure (footings, walls, columns, slabs), pulls in related details, and eliminates hours of manual quantity calculation. Here’s a demo: https://www.youtube.com/watch?v=PAMNDRWEdlI . The problem: Concrete subcontractors are the backbone of every building, but their estimating workflow hasn't changed in 20 years. Right now, a senior estimator opens a PDF, manually traces every footing and grade beam, then hand-builds an Excel spreadsheet with 300+ line items- volumes, formwork, rebar by bar size with lap splices and development lengths. Bids can take weeks and even months. Most firms have just a few estimators, meaning they physically cannot bid on most of the work available to them. The software incumbent in this trade hasn’t been updated since 2020. Beyond that, every AI takeoff tool on the market was built for GCs and treats concrete as one checkbox, rather than working around how concrete estimators actually price work. We’re building Rudus for this trade and only this trade. We started this when Sahil took a construction management class and realized how the estimation workflows hadn't changed in decades. We started cold calling, walking into offices with donuts, showing up at job sites, and everyone told us the same thing: slow estimation is the biggest bottleneck in growing their business, but every new product they've tried has failed. We quickly realized that the reason those tools failed is a lack of trust and frequent errors causing later problems. Estimators stake million to billion dollar bids on these numbers, and they are clear that they won’t trade their workflow for a black box. We took a different approach: software that intelligently accelerates their current workflows rather than replacing it by forward deploying our product into their current estimation workflow. When an estimator uploads their structural PDFs to Rudus, we auto-classify every sheet (foundation plans, section details, footing schedules, frame elevations) and route each to the right processing pipeline. Computer vision detects concrete elements across the drawing set and follows cross-references across sheets to resolve dimensions and detailing, catching elements that plan-only tools always miss. Each element gets expanded into full assembly line items: concrete, formwork, and rebar with all the calculations an estimator would normally do by hand. A typical foundation package goes from a handful of assemblies to 80-120 priced line items. The estimator reviews, overrides where needed, and exports straight into their existing workflow. We have a couple key advantages in the AI estimation space. The first is our focus on concrete, a niche part of construction. No one else is building this for concrete subs because the sheets vary drastically from other subtrades. For this same reason, VLMs and other generic solutions don't work. Instead, proprietary computer vision models are required, relying on training from massive amounts of customer data. We run multiple different models trained directly on our customers' takeoffs, and every interaction from our customers with our models becomes a training example, allowing accuracy per client to sharpen with use. Our second advantage is in our product methodology, as we’ve chosen to build a copilot, not a black box. Most AI takeoff platforms try to replace the estimator completely by autonomously producing quantities, but the quality of the outputs with current models is poor, so the takeoff gets redone by hand anyway. After 100+ hours sitting in rooms with structural concrete estimators and completing numerous takeoffs ourselves, we’ve built around their actual workflow. The estimator starts the takeoff, and Rudus extends the w

[truncated]

