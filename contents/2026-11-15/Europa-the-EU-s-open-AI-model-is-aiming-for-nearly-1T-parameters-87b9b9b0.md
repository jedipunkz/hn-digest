---
source: "https://maccus.fr/en/eu-open-ai-model-europa/"
hn_url: "https://news.ycombinator.com/item?id=50506796"
title: "Europa, the EU's open AI model, is aiming for nearly 1T parameters"
article_title: "EUROPA, the EU's open AI model, is aiming for nearly 1 trillion parameters - MACCUS"
image: "https://maccus.fr/wp-content/uploads/2026/11/europa-throughput-en.jpg"
author: "mstkllah"
captured_at: "2026-11-15T13:30:59Z"
capture_tool: "hn-digest"
hn_id: 50506796
score: 1
comments: 0
posted_at: "2026-11-15T13:29:48Z"
tags:
  - hacker-news
---

# Europa, the EU's open AI model, is aiming for nearly 1T parameters

- HN: [50506796](https://news.ycombinator.com/item?id=50506796)
- Source: [maccus.fr](https://maccus.fr/en/eu-open-ai-model-europa/)
- Score: 1
- Comments: 0
- Posted: 2026-11-15T13:29:48Z

## Translation

Title: Europa, the EU's open AI model, is aiming for nearly 1T parameters
Article title: EUROPA, the EU's open AI model, is aiming for nearly 1 trillion parameters - MACCUS
Description: Domyn has shared the training setup for EUROPA, the open European AI model chosen by the Commission: 951.9 billion parameters, tested on JUPITER.

Article text:
EUROPA, the EU's open AI model, is aiming for nearly 1 trillion parameters - MACCUS
Skip to content
English Français
Menu
Menu
Services
Service Overview
EUROPA, the EU’s open AI model, is aiming for nearly 1 trillion parameters
Domyn is testing the setup for EUROPA, the open AI model picked by the European Commission, on the public JUPITER supercomputer.
The Italian company Domyn has run training trials for EUROPA, the European Union’s open AI model, on a public supercomputer. The test setup has nearly 1 trillion parameters, more than double the minimum the European Commission set. Domyn, which the Commission picked in June to build EUROPA, described the trials in a technical post on November 12, 2026.
Parameters are the internal settings a model learns during training. In its test setup, EUROPA has 951.9 billion of them. The supercomputer, JUPITER, is in Jülich, Germany, and is shared with hundreds of other projects. Domyn says it made training on the machine 14 times faster than at the start of its tests.
Domyn has not yet released EUROPA’s weights, the files that hold its parameters and let anyone run it on their own hardware . It plans to publish them with details of its training methods when EUROPA comes out. The post gives no date. In June, the outlet Declic Media reported that Domyn was aiming for a release “within the year.”
A project the Commission picked in June
On February 13, 2026, the European Commission launched the Frontier AI Grand Challenge. It ran the contest with the EuroHPC Joint Undertaking, which oversees Europe’s network of public supercomputers. The goal was to pick one project to train a frontier model, a top-tier AI system, with at least 400 billion parameters. The winner gets up to 2.5% of EuroHPC’s total computing capacity for one year.
On June 19, the Commission named the winner: the EUROPA consortium. Domyn, a Milan company formerly known as iGenius, leads the group. Its partners are two German institutes from the Fraunhofer research network, IAIS and IIS. The Commission wants open models that public authorities, researchers and businesses can use. EUROPA also has to cover all 24 official EU languages.
“Europe can lead in advanced AI on its own terms,” said Henna Virkkunen, the Commission’s executive vice president for tech sovereignty, security and democracy.
How the 951.9 billion parameters are organized
Domyn says it aimed for about 1 trillion parameters, the size of today’s frontier models. The Commission’s minimum was 400 billion.
EUROPA uses a design called “mixture of experts.” The network is split into 384 specialized sub-networks, called experts. A model reads text in small units called tokens, each roughly a word or part of a word. Only 8 of the 384 experts work on any given token. As a result, only 32.8 billion of the 951.9 billion parameters are active for each token. That cuts the amount of computing needed.
During training, the model reads text in sequences of 65,536 tokens. Domyn notes that these figures come from its trials and that “some values may change before the training run starts.”
JUPITER, a shared machine built for science
According to the Jülich research center, JUPITER is Europe’s first exascale supercomputer. That means it can perform more than a quintillion (a billion billion) operations per second. The EuroHPC Joint Undertaking owns the machine and is putting €250 million into it. The German federal government and the state of North Rhine-Westphalia are each contributing €125 million.
Domyn uses JUPITER’s Booster module, a section of the machine with 5,884 nodes, the networked computers that make up the machine. Each node has four NVIDIA GH200 chips. Each chip pairs a general-purpose processor, or CPU, with a GPU, the processor that does most of the AI math.
According to Domyn, models this size have mostly been trained on dedicated clusters of NVIDIA’s newer Blackwell chips. On JUPITER, Domyn has only user-level access. That means it cannot change the operating system, the network or the nodes’ memory, all of which are set up for scientific simulation.
The network between nodes is the bottleneck. Training can run several copies of the model in parallel, and each copy spans 384 GPUs across 96 nodes. Those nodes have to exchange data. But the link between two nodes is 36 times slower than the link between a GPU and the CPU on its own chip.
From 90 to about 1,300 tokens per second per GPU
At first, each GPU processed 90 tokens per second and spent most of its time waiting for data from the network. Domyn started by making the text sequences longer, giving the GPUs more work between exchanges. The team then moved some intermediate data, the temporary results of the calculations, to the CPU’s memory to free up memory on the GPU.
Domyn also moved the heaviest part of the optimizer, an algorithm called Muon, from the CPU to the GPU. The optimizer adjusts the parameters at each training step, one round of learning. Its work on each step dropped from 41.7 seconds to 2.2 seconds.
Other changes cut the amount of data sent over the network. The experts’ weights, the values of their parameters, now travel compressed in FP8 format, at 8 bits per value. The gradients, the corrections computed at each step, travel in bf16 format, at 16 bits per value.
A fix to memory management raised throughput, the number of tokens processed per second, by 21%. In the end, each GPU processed 1,293 tokens per second, up from 90 at the start. Domyn contributed its fixes back to PyTorch and TorchTitan, the open-source software it uses to train the model.
At that rate, one copy of the model processes about 43 billion tokens a day, according to Domyn. With four copies running in parallel on 1,536 GPUs, the company puts the total at about 150 billion tokens a day. At that scale, each GPU keeps 89% of the speed it reaches when a single copy runs.
The Italian daily Quotidiano Nazionale concludes that full pretraining, the main phase of learning from raw text, would take about four months.
What the trials do not show yet
The post opens by saying the company has “recently started training” a frontier-scale model. Yet the rest of the text describes only trials. The throughput figures are median values, measured over steps 4 through 10 of each trial. The setup may also change before the training run starts.
Quotidiano Nazionale put it this way: the demonstration “does not yet amount to completing the training.”
Domyn also flags its own uncertainties. A rule that spreads tokens evenly across four groups of 96 experts, one group per GPU, speeds up the computing. But Domyn says that, based on its early runs, this rule “can indeed hurt training.”
The company says FP8 compression and bf16 gradients do not measurably change the model’s loss, the gap between its predictions and the actual text. Its comparisons, however, cover only ten training steps. The post says nothing about the training data, the share of each language, expected performance or the license.
Domyn Edge and Europe’s other projects
Domyn has already released a much smaller open model, Domyn Edge, which it began rolling out on October 20, 2026. It has 4 billion parameters, compared with 951.9 billion for EUROPA’s test setup. Domyn trained it from scratch on 13 trillion tokens. It used Leonardo, a European supercomputer hosted by the CINECA computing center in Italy.
Domyn Edge is released under the Apache 2.0 license, an open-source license that allows commercial use. The release also includes part of the training data, the code and intermediate versions of the model. French is one of its five main languages.
In June, Declic Media reported that Europe was running at least three large “sovereign” AI model projects in parallel, meaning models developed under European control. One of them, OpenEuroLLM, has a budget of €37.4 million and announced models of 7 billion to 175 billion parameters. Declic Media added that EUROPA’s target went beyond anything Europe had attempted so far. Its author also noted that Domyn’s model “hasn’t written its first line of weights.”
The November 12 post shows a project much larger than the Commission asked for. The Commission required at least 400 billion parameters, and the consortium is aiming for about 1 trillion. Domyn’s test measurements suggest that a shared public machine built for science can train a model this size at a workable speed. So far, though, this is an engineering demonstration. The full training run still has to happen, and then the model has to be evaluated and released.
In the private sector, the French company Mistral AI released the weights of Mistral Large 4 , a model with about 1 trillion parameters, on October 29. Mistral is backed by private investors and says it trained the model in its own data centers in Europe.
Running EUROPA on in-house hardware would take a lot of memory, even with the usual compression to 4 bits per parameter. At that level, its weights alone would take up at least 476 GB by our math, plus working memory. That puts it in the same range as Muse Spark Open, Meta’s open model . That model has 816 billion parameters and runs on two Apple Mac Studio desktops with 512 GB of memory each.
Several questions remain open. Neither the Commission nor Domyn has named EUROPA’s license, the legal terms for using it. The training data, the share of French in it, the model’s benchmark test scores and the release date are all still unknown.
EuroHPC’s grant of computing time lasts only one year. Three milestones are worth watching: the actual start of the full training run, the reports Domyn plans to publish “as the training run progresses,” and the technical report on Domyn Edge, expected in November.
Training a 1T-Parameter MoE on EU public infrastructure with PyTorch and TorchTitan , Domyn, November 12, 2026
Training a 1T-Parameter MoE on EU public infrastructure with PyTorch and TorchTitan , PyTorch, November 12, 2026
Commission selects EUROPA consortium as the winner of the Frontier AI Grand Challenge , European Commission, June 19, 2026
Turning strategy into action: Commission launches Frontier AI Grand Challenge , European Commission, February 13, 2026
JUPITER , Forschungszentrum Jülich, accessed November 15, 2026
IA, la sfida europea dei supercomputer: testato un modello da quasi mille miliardi di parametri , Quotidiano Nazionale, November 13, 2026
Domyn Edge Base (model card) , Domyn on Hugging Face, accessed November 15, 2026
Europe Names Its 3rd Sovereign LLM, Ships None , Declic Media, June 27, 2026
Header image: MACCUS chart, data from Domyn (technical blog post, November 12, 2026).
Google’s AI agents finished porting an OS kernel to Rust in two weeks, and it boots
68% of AI-savvy Japanese trusted themselves to spot AI-written fakes. A year on, 52% did
Microsoft apologizes for the Microsoft 365 outage and says AI wasn’t to blame
US call centers have lost 235,000 jobs since their 2016 peak. About half went before ChatGPT
1 Marketing Italy fines Meta €805 million over Facebook’s access to Android photo galleries
2 AI Google’s AI agents finished porting an OS kernel to Rust in two weeks, and it boots
3 AI 68% of AI-savvy Japanese trusted themselves to spot AI-written fakes. A year on, 52% did
Founder of MACCUS, publisher of maccus.fr
Marc Winer, born in 1997, is French American and lives in Tournai, Belgium. He founded MACCUS, an SEO agency there, and covers AI, SEO and marketing news for maccus.fr. He created marcwiner.com, a French-language Asian food site that gets more than 13 million visits a year and is translated into 23 languages. TV Tokyo filmed him for a segment on okonomiyaki. His app, HitchHive, helps people in the same city meet up for activities. It had more than 20,000 users as of September 2026. He graduated from UCLouvain and from IE Business School in Madrid.
marcwiner.com, 13 million visits a year
TV Tokyo segment (in Japanese)
Categories AI , Blog , News
Leave a Comment Cancel reply
Save my name, email, an

[truncated]

## Original Extract

Domyn has shared the training setup for EUROPA, the open European AI model chosen by the Commission: 951.9 billion parameters, tested on JUPITER.

EUROPA, the EU's open AI model, is aiming for nearly 1 trillion parameters - MACCUS
Skip to content
English Français
Menu
Menu
Services
Service Overview
EUROPA, the EU’s open AI model, is aiming for nearly 1 trillion parameters
Domyn is testing the setup for EUROPA, the open AI model picked by the European Commission, on the public JUPITER supercomputer.
The Italian company Domyn has run training trials for EUROPA, the European Union’s open AI model, on a public supercomputer. The test setup has nearly 1 trillion parameters, more than double the minimum the European Commission set. Domyn, which the Commission picked in June to build EUROPA, described the trials in a technical post on November 12, 2026.
Parameters are the internal settings a model learns during training. In its test setup, EUROPA has 951.9 billion of them. The supercomputer, JUPITER, is in Jülich, Germany, and is shared with hundreds of other projects. Domyn says it made training on the machine 14 times faster than at the start of its tests.
Domyn has not yet released EUROPA’s weights, the files that hold its parameters and let anyone run it on their own hardware . It plans to publish them with details of its training methods when EUROPA comes out. The post gives no date. In June, the outlet Declic Media reported that Domyn was aiming for a release “within the year.”
A project the Commission picked in June
On February 13, 2026, the European Commission launched the Frontier AI Grand Challenge. It ran the contest with the EuroHPC Joint Undertaking, which oversees Europe’s network of public supercomputers. The goal was to pick one project to train a frontier model, a top-tier AI system, with at least 400 billion parameters. The winner gets up to 2.5% of EuroHPC’s total computing capacity for one year.
On June 19, the Commission named the winner: the EUROPA consortium. Domyn, a Milan company formerly known as iGenius, leads the group. Its partners are two German institutes from the Fraunhofer research network, IAIS and IIS. The Commission wants open models that public authorities, researchers and businesses can use. EUROPA also has to cover all 24 official EU languages.
“Europe can lead in advanced AI on its own terms,” said Henna Virkkunen, the Commission’s executive vice president for tech sovereignty, security and democracy.
How the 951.9 billion parameters are organized
Domyn says it aimed for about 1 trillion parameters, the size of today’s frontier models. The Commission’s minimum was 400 billion.
EUROPA uses a design called “mixture of experts.” The network is split into 384 specialized sub-networks, called experts. A model reads text in small units called tokens, each roughly a word or part of a word. Only 8 of the 384 experts work on any given token. As a result, only 32.8 billion of the 951.9 billion parameters are active for each token. That cuts the amount of computing needed.
During training, the model reads text in sequences of 65,536 tokens. Domyn notes that these figures come from its trials and that “some values may change before the training run starts.”
JUPITER, a shared machine built for science
According to the Jülich research center, JUPITER is Europe’s first exascale supercomputer. That means it can perform more than a quintillion (a billion billion) operations per second. The EuroHPC Joint Undertaking owns the machine and is putting €250 million into it. The German federal government and the state of North Rhine-Westphalia are each contributing €125 million.
Domyn uses JUPITER’s Booster module, a section of the machine with 5,884 nodes, the networked computers that make up the machine. Each node has four NVIDIA GH200 chips. Each chip pairs a general-purpose processor, or CPU, with a GPU, the processor that does most of the AI math.
According to Domyn, models this size have mostly been trained on dedicated clusters of NVIDIA’s newer Blackwell chips. On JUPITER, Domyn has only user-level access. That means it cannot change the operating system, the network or the nodes’ memory, all of which are set up for scientific simulation.
The network between nodes is the bottleneck. Training can run several copies of the model in parallel, and each copy spans 384 GPUs across 96 nodes. Those nodes have to exchange data. But the link between two nodes is 36 times slower than the link between a GPU and the CPU on its own chip.
From 90 to about 1,300 tokens per second per GPU
At first, each GPU processed 90 tokens per second and spent most of its time waiting for data from the network. Domyn started by making the text sequences longer, giving the GPUs more work between exchanges. The team then moved some intermediate data, the temporary results of the calculations, to the CPU’s memory to free up memory on the GPU.
Domyn also moved the heaviest part of the optimizer, an algorithm called Muon, from the CPU to the GPU. The optimizer adjusts the parameters at each training step, one round of learning. Its work on each step dropped from 41.7 seconds to 2.2 seconds.
Other changes cut the amount of data sent over the network. The experts’ weights, the values of their parameters, now travel compressed in FP8 format, at 8 bits per value. The gradients, the corrections computed at each step, travel in bf16 format, at 16 bits per value.
A fix to memory management raised throughput, the number of tokens processed per second, by 21%. In the end, each GPU processed 1,293 tokens per second, up from 90 at the start. Domyn contributed its fixes back to PyTorch and TorchTitan, the open-source software it uses to train the model.
At that rate, one copy of the model processes about 43 billion tokens a day, according to Domyn. With four copies running in parallel on 1,536 GPUs, the company puts the total at about 150 billion tokens a day. At that scale, each GPU keeps 89% of the speed it reaches when a single copy runs.
The Italian daily Quotidiano Nazionale concludes that full pretraining, the main phase of learning from raw text, would take about four months.
What the trials do not show yet
The post opens by saying the company has “recently started training” a frontier-scale model. Yet the rest of the text describes only trials. The throughput figures are median values, measured over steps 4 through 10 of each trial. The setup may also change before the training run starts.
Quotidiano Nazionale put it this way: the demonstration “does not yet amount to completing the training.”
Domyn also flags its own uncertainties. A rule that spreads tokens evenly across four groups of 96 experts, one group per GPU, speeds up the computing. But Domyn says that, based on its early runs, this rule “can indeed hurt training.”
The company says FP8 compression and bf16 gradients do not measurably change the model’s loss, the gap between its predictions and the actual text. Its comparisons, however, cover only ten training steps. The post says nothing about the training data, the share of each language, expected performance or the license.
Domyn Edge and Europe’s other projects
Domyn has already released a much smaller open model, Domyn Edge, which it began rolling out on October 20, 2026. It has 4 billion parameters, compared with 951.9 billion for EUROPA’s test setup. Domyn trained it from scratch on 13 trillion tokens. It used Leonardo, a European supercomputer hosted by the CINECA computing center in Italy.
Domyn Edge is released under the Apache 2.0 license, an open-source license that allows commercial use. The release also includes part of the training data, the code and intermediate versions of the model. French is one of its five main languages.
In June, Declic Media reported that Europe was running at least three large “sovereign” AI model projects in parallel, meaning models developed under European control. One of them, OpenEuroLLM, has a budget of €37.4 million and announced models of 7 billion to 175 billion parameters. Declic Media added that EUROPA’s target went beyond anything Europe had attempted so far. Its author also noted that Domyn’s model “hasn’t written its first line of weights.”
The November 12 post shows a project much larger than the Commission asked for. The Commission required at least 400 billion parameters, and the consortium is aiming for about 1 trillion. Domyn’s test measurements suggest that a shared public machine built for science can train a model this size at a workable speed. So far, though, this is an engineering demonstration. The full training run still has to happen, and then the model has to be evaluated and released.
In the private sector, the French company Mistral AI released the weights of Mistral Large 4 , a model with about 1 trillion parameters, on October 29. Mistral is backed by private investors and says it trained the model in its own data centers in Europe.
Running EUROPA on in-house hardware would take a lot of memory, even with the usual compression to 4 bits per parameter. At that level, its weights alone would take up at least 476 GB by our math, plus working memory. That puts it in the same range as Muse Spark Open, Meta’s open model . That model has 816 billion parameters and runs on two Apple Mac Studio desktops with 512 GB of memory each.
Several questions remain open. Neither the Commission nor Domyn has named EUROPA’s license, the legal terms for using it. The training data, the share of French in it, the model’s benchmark test scores and the release date are all still unknown.
EuroHPC’s grant of computing time lasts only one year. Three milestones are worth watching: the actual start of the full training run, the reports Domyn plans to publish “as the training run progresses,” and the technical report on Domyn Edge, expected in November.
Training a 1T-Parameter MoE on EU public infrastructure with PyTorch and TorchTitan , Domyn, November 12, 2026
Training a 1T-Parameter MoE on EU public infrastructure with PyTorch and TorchTitan , PyTorch, November 12, 2026
Commission selects EUROPA consortium as the winner of the Frontier AI Grand Challenge , European Commission, June 19, 2026
Turning strategy into action: Commission launches Frontier AI Grand Challenge , European Commission, February 13, 2026
JUPITER , Forschungszentrum Jülich, accessed November 15, 2026
IA, la sfida europea dei supercomputer: testato un modello da quasi mille miliardi di parametri , Quotidiano Nazionale, November 13, 2026
Domyn Edge Base (model card) , Domyn on Hugging Face, accessed November 15, 2026
Europe Names Its 3rd Sovereign LLM, Ships None , Declic Media, June 27, 2026
Header image: MACCUS chart, data from Domyn (technical blog post, November 12, 2026).
Google’s AI agents finished porting an OS kernel to Rust in two weeks, and it boots
68% of AI-savvy Japanese trusted themselves to spot AI-written fakes. A year on, 52% did
Microsoft apologizes for the Microsoft 365 outage and says AI wasn’t to blame
US call centers have lost 235,000 jobs since their 2016 peak. About half went before ChatGPT
1 Marketing Italy fines Meta €805 million over Facebook’s access to Android photo galleries
2 AI Google’s AI agents finished porting an OS kernel to Rust in two weeks, and it boots
3 AI 68% of AI-savvy Japanese trusted themselves to spot AI-written fakes. A year on, 52% did
Founder of MACCUS, publisher of maccus.fr
Marc Winer, born in 1997, is French American and lives in Tournai, Belgium. He founded MACCUS, an SEO agency there, and covers AI, SEO and marketing news for maccus.fr. He created marcwiner.com, a French-language Asian food site that gets more than 13 million visits a year and is translated into 23 languages. TV Tokyo filmed him for a segment on okonomiyaki. His app, HitchHive, helps people in the same city meet up for activities. It had more than 20,000 users as of September 2026. He graduated from UCLouvain and from IE Business School in Madrid.
marcwiner.com, 13 million visits a year
TV Tokyo segment (in Japanese)
Categories AI , Blog , News
Leave a Comment Cancel reply
Save my name, email, an

[truncated]
