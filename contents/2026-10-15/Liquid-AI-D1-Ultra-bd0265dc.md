---
source: "https://www.liquid.ai/blog/d1-ultra"
hn_url: "https://news.ycombinator.com/item?id=50106907"
title: "Liquid AI D1-Ultra"
article_title: "d1-ultra: Introducing the World’s Best Decision Model | Blog | Liquid AI"
image: "https://aypchzzf9pftwuto.public.blob.vercel-storage.com/image%20(4)-mh23DVnFm2EDs4US6gJskniHIMwVY4.png"
author: "acossta"
captured_at: "2026-10-15T18:02:47Z"
capture_tool: "hn-digest"
hn_id: 50106907
score: 6
comments: 0
posted_at: "2026-10-15T17:15:43Z"
tags:
  - hacker-news
---

# Liquid AI D1-Ultra

- HN: [50106907](https://news.ycombinator.com/item?id=50106907)
- Source: [www.liquid.ai](https://www.liquid.ai/blog/d1-ultra)
- Score: 6
- Comments: 0
- Posted: 2026-10-15T17:15:43Z

## Translation

Title: Liquid AI D1-Ultra
Article title: d1-ultra: Introducing the World’s Best Decision Model | Blog | Liquid AI
Description: d1-ultra, the world's best decision model for text and images, built for hard real-world cases. Available on the Liquid AI API

Article text:
d1-ultra: Introducing the World’s Best Decision Model | Blog | Liquid AI Models
News / Models d1-ultra: Introducing the World’s Best Decision Model
Today, we release d1-ultra, our new multimodal decision model with state-of-the-art performance . Most decision models work well on easy and common tasks. We built d1-ultra for harder, messier real-world use cases, like spotting a defect on a circuit board against its spec or checking a damage claim against the customer’s photo.
d1-ultra is available today on the Liquid API at $0.042 per million input tokens.
Try on d1 Playground What d1-ultra does
d1-ultra answers the same three question types as d1:
noul : a yes/no question, answered with a probability from 0 to 1.
Choice : one label from many, with a probability per label.
Score : a position on a scale.
Compared to d1, we based d1-ultra on a larger open-weight model with a series of Liquid-specific heavy customizations. It provides significant gains when combined with our training pipeline, especially on knowledge-related tasks.
We compared it to our previous model d1 and Jev 1.14 across 112 unique evaluations.
Accuracy is not enough: decision models also output probabilities that must be properly calibrated. We realized that most decision models are overconfident in their responses. That’s why we specifically trained d1-ultra to produce more reliable probabilities.
d1-ultra outperforms Jev 1.14 across all tasks, delivering better-calibrated answers and higher accuracy at the same price. To see how it works in practice, you can find more demos running on d1-ultra in the d1 Playground .
d1-ultra also outperforms all other decision models Artificial Analysis Decision Index by at least 8 points."
d1-ultra is available on the Liquid AI API . Create your API key and start building today. d1-ultra will soon be available through Vercel and OpenRouter .
The model reads text, images, or both at the same time, like in this example:
import base64, os, requests
image = base64.b64encode(open("board.jpg", "rb").read()).decode()
response = requests.post(
"https://api.liquid.ai/decisions/v1/systemone",
headers={"Authorization": f"Bearer {os.environ['LIQUID_API_KEY']}"},
json={
"model": "d1-ultra",
"images": [f"data:image/jpeg;base64,{image}"],
"state": "Camera image of a circuit board on the production line.",
"questions":
{"defect": {
"type": "noul",
"instructions": "Does this circuit board have a defect?"
}
},
},
)
print(response.json()["answers"]["defect"]["noul"]) The full API reference is in the decision models documentation .
d1-ultra is billed on input tokens only, with no output tokens. Images are counted as input tokens at the same rate as text: 1.5 tokens per 32×32-pixel patch, so a 1024×1024 image costs 1,536 tokens. Each question is billed as its own prompt, including its text and all images.
d1-ultra is our most capable decision model so far. We’re committed to continuing our work on decision models across a range of sizes, with new features, higher decision quality, and lower latency, and we plan to provide regular updates.
For citations, please use the following reference or BibTeX:
Liquid AI, "d1-ultra: State-of-the-Art Decision Model", Liquid AI Blog, Oct 2026.
@article{liquidAI2026d1-ultra,
author = {Liquid AI},
title = {d1-ultra: State-of-the-Art Decision Model},
journal = {Liquid AI Blog},
year = {2026},
note = {www.liquid.ai/blog/d1-ultra},
}
For press inquiries, email us at press@liquid.ai
Building the fastest, most compute-efficient, and capable foundation models so intelligence can live anywhere.
Liquid Foundation Models (LFMs)
Automated Foundation Model Design

## Original Extract

d1-ultra, the world's best decision model for text and images, built for hard real-world cases. Available on the Liquid AI API

d1-ultra: Introducing the World’s Best Decision Model | Blog | Liquid AI Models
News / Models d1-ultra: Introducing the World’s Best Decision Model
Today, we release d1-ultra, our new multimodal decision model with state-of-the-art performance . Most decision models work well on easy and common tasks. We built d1-ultra for harder, messier real-world use cases, like spotting a defect on a circuit board against its spec or checking a damage claim against the customer’s photo.
d1-ultra is available today on the Liquid API at $0.042 per million input tokens.
Try on d1 Playground What d1-ultra does
d1-ultra answers the same three question types as d1:
noul : a yes/no question, answered with a probability from 0 to 1.
Choice : one label from many, with a probability per label.
Score : a position on a scale.
Compared to d1, we based d1-ultra on a larger open-weight model with a series of Liquid-specific heavy customizations. It provides significant gains when combined with our training pipeline, especially on knowledge-related tasks.
We compared it to our previous model d1 and Jev 1.14 across 112 unique evaluations.
Accuracy is not enough: decision models also output probabilities that must be properly calibrated. We realized that most decision models are overconfident in their responses. That’s why we specifically trained d1-ultra to produce more reliable probabilities.
d1-ultra outperforms Jev 1.14 across all tasks, delivering better-calibrated answers and higher accuracy at the same price. To see how it works in practice, you can find more demos running on d1-ultra in the d1 Playground .
d1-ultra also outperforms all other decision models Artificial Analysis Decision Index by at least 8 points."
d1-ultra is available on the Liquid AI API . Create your API key and start building today. d1-ultra will soon be available through Vercel and OpenRouter .
The model reads text, images, or both at the same time, like in this example:
import base64, os, requests
image = base64.b64encode(open("board.jpg", "rb").read()).decode()
response = requests.post(
"https://api.liquid.ai/decisions/v1/systemone",
headers={"Authorization": f"Bearer {os.environ['LIQUID_API_KEY']}"},
json={
"model": "d1-ultra",
"images": [f"data:image/jpeg;base64,{image}"],
"state": "Camera image of a circuit board on the production line.",
"questions":
{"defect": {
"type": "noul",
"instructions": "Does this circuit board have a defect?"
}
},
},
)
print(response.json()["answers"]["defect"]["noul"]) The full API reference is in the decision models documentation .
d1-ultra is billed on input tokens only, with no output tokens. Images are counted as input tokens at the same rate as text: 1.5 tokens per 32×32-pixel patch, so a 1024×1024 image costs 1,536 tokens. Each question is billed as its own prompt, including its text and all images.
d1-ultra is our most capable decision model so far. We’re committed to continuing our work on decision models across a range of sizes, with new features, higher decision quality, and lower latency, and we plan to provide regular updates.
For citations, please use the following reference or BibTeX:
Liquid AI, "d1-ultra: State-of-the-Art Decision Model", Liquid AI Blog, Oct 2026.
@article{liquidAI2026d1-ultra,
author = {Liquid AI},
title = {d1-ultra: State-of-the-Art Decision Model},
journal = {Liquid AI Blog},
year = {2026},
note = {www.liquid.ai/blog/d1-ultra},
}
For press inquiries, email us at press@liquid.ai
Building the fastest, most compute-efficient, and capable foundation models so intelligence can live anywhere.
Liquid Foundation Models (LFMs)
Automated Foundation Model Design
