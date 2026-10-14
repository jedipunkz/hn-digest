---
source: "https://deepmind.google/blog/on-device-ai-scam-detection-mddi/"
hn_url: "https://news.ycombinator.com/item?id=50086596"
title: "DeepMind and Singapore fine-tune an on-device Gemma model for scam detection"
article_title: "On-Device AI Against Scams: Google DeepMind & MDDI — Google DeepMind"
image: ""
author: "acossta"
captured_at: "2026-10-14T13:21:42Z"
capture_tool: "hn-digest"
hn_id: 50086596
score: 3
comments: 1
posted_at: "2026-10-14T12:24:44Z"
tags:
  - hacker-news
---

# DeepMind and Singapore fine-tune an on-device Gemma model for scam detection

- HN: [50086596](https://news.ycombinator.com/item?id=50086596)
- Source: [deepmind.google](https://deepmind.google/blog/on-device-ai-scam-detection-mddi/)
- Score: 3
- Comments: 1
- Posted: 2026-10-14T12:24:44Z

## Translation

Title: DeepMind and Singapore fine-tune an on-device Gemma model for scam detection
Article title: On-Device AI Against Scams: Google DeepMind & MDDI — Google DeepMind

Article text:
Skip to main content Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Google DeepMind Google AI Learn about all our AI Google DeepMind Explore the frontier of AI Google Labs Try our AI experiments Google Research Explore our research Products and apps Gemini app Chat with Gemini Google AI Studio Build with our next-gen models Google Antigravity Our agentic development platform Models Research Science About Build with Gemini Try Gemini October 14, 2026 Responsibility & Safety Google DeepMind and Singapore Government collaborate to build fine-tuned AI model for scam detection
JJ Geewax and Miguel de Andrés
Share Copied Collaborating with the Government Technology Agency of Singapore (GovTech Singapore) to protect users without compromising privacy.
Scams, such as investment fraud, fake job offers, and impersonation of public officials make up over half of Singapore’s reported crimes, and account for more than S$4 billion in losses since 2019 . At the same time, scam tactics have shifted from simple phishing to complex social engineering, whereby attackers manipulate victims with convincing conversations before requesting money. Such scams now account for up to 80% of all losses.
To address this, Google DeepMind and the Government Technology Agency of Singapore (GovTech Singapore) have been working together as part of our National AI Partnership with the Singapore Government. Together, we've fine-tuned a privacy-preserving, on-device scam detection model.
On-device AI protection without compromising privacy
Generic models often struggle with local deception techniques, making scam detection harder. To address this, Google’s lightweight Gemma model was fine-tuned on local scam conversations. These ranged from impersonations of public institutions, banks, and local figures, to trending regional schemes. This ensures the fine-tuned model will be able to distinguish benign social chatter from genuine manipulation when deployed in applications installed by citizens.
Unlike traditional threat detection, which relies on server-side scanning of messages in the cloud, this model can process conversation data directly on-device. When integrated into applications, this will allow faster analysis of conversation patterns to detect manipulative intent and help citizens act before financial or credential exposure can occur.
The roll out of the model within applications will be carefully phased and developed in partnership with the Singapore Government. When deployed, the model will only be able to access on-device conversations with the user's consent, in accordance with the terms of use of that application. As the model runs entirely on-device, neither the Singapore Government nor Google, will be able to access users’ conversations through its deployment.
Fighting scams shouldn't mean sacrificing personal privacy. Working with Google DeepMind, we're bringing advanced AI protection directly onto citizens' devices. This allows us to catch highly targeted, local scam attempts in real time while keeping personal communications private.
Under the hood: local context and edge efficiency
When the model detects manipulative conversational patterns, it triggers a real-time warning directly on the user's device advising them to pause and exercise caution before sharing sensitive information or funds.
Running a complex scam classifier model on modern smartphones is no easy task, with battery efficiency being a critical area of concern. To do this, we're relying on a two-stage pipeline:
Stage 1: A lightweight, efficient classifier screens inbound messages, using minimal power, focused on catching potential threats.
Stage 2: In-context analysis with our fine-tuned Gemma model, where we analyze a larger snapshot of the conversation to confirm a threat or discard a false alarm.
Presented at Singapore International Cyber Week (SICW 2026), this collaboration provides a model for messaging platforms balancing user protection with end-to-end encryption. Strict governance was maintained throughout development:
Sovereign Sandbox: Fine-tuning took place within GovTech Singapore's managed and maintained GCP environment, keeping training data isolated.
No Citizen Data: Training strictly excluded private or personally identifiable information, utilising anonymised datasets derived from scam reports and intelligence feeds.
Stress-Testing: The model was tested against synthetic dialogues and local corpora to cover diverse scam categories and evolving impersonation tactics.
The future of on-device safety
The project demonstrates how AI can help counter criminality without compromising safety or privacy. It also shows how governments and model providers can work together to build tech for public good. With this lightweight, on-device model, we've proven the ability for local, real-time scam detection while ensuring chat interactions remain secure. We look forward to sharing more soon on our collaboration with the Singapore government, as well as exploring the potential for similar technology to be deployed in other countries.
Strengthening Singapore’s AI Future: A New National Partnership
Unlocking UK house-building with AI-accelerated planning
Announcing our partnership with the Republic of Korea
Accelerating discovery in India through AI-powered science and education
Strengthening our partnership with the UK government to support prosperity and security in the AI era
Deepening our partnership with the UK AI Security Institute
Google DeepMind supports U.S. Department of Energy on Genesis: a national mission to accelerate innovation and scientific discovery
Follow us Sign up for updates on our latest innovations I accept Google's Terms and Conditions and acknowledge that my information will be used in accordance with Google's Privacy Policy .

## Original Extract

Skip to main content Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Google DeepMind Google AI Learn about all our AI Google DeepMind Explore the frontier of AI Google Labs Try our AI experiments Google Research Explore our research Products and apps Gemini app Chat with Gemini Google AI Studio Build with our next-gen models Google Antigravity Our agentic development platform Models Research Science About Build with Gemini Try Gemini October 14, 2026 Responsibility & Safety Google DeepMind and Singapore Government collaborate to build fine-tuned AI model for scam detection
JJ Geewax and Miguel de Andrés
Share Copied Collaborating with the Government Technology Agency of Singapore (GovTech Singapore) to protect users without compromising privacy.
Scams, such as investment fraud, fake job offers, and impersonation of public officials make up over half of Singapore’s reported crimes, and account for more than S$4 billion in losses since 2019 . At the same time, scam tactics have shifted from simple phishing to complex social engineering, whereby attackers manipulate victims with convincing conversations before requesting money. Such scams now account for up to 80% of all losses.
To address this, Google DeepMind and the Government Technology Agency of Singapore (GovTech Singapore) have been working together as part of our National AI Partnership with the Singapore Government. Together, we've fine-tuned a privacy-preserving, on-device scam detection model.
On-device AI protection without compromising privacy
Generic models often struggle with local deception techniques, making scam detection harder. To address this, Google’s lightweight Gemma model was fine-tuned on local scam conversations. These ranged from impersonations of public institutions, banks, and local figures, to trending regional schemes. This ensures the fine-tuned model will be able to distinguish benign social chatter from genuine manipulation when deployed in applications installed by citizens.
Unlike traditional threat detection, which relies on server-side scanning of messages in the cloud, this model can process conversation data directly on-device. When integrated into applications, this will allow faster analysis of conversation patterns to detect manipulative intent and help citizens act before financial or credential exposure can occur.
The roll out of the model within applications will be carefully phased and developed in partnership with the Singapore Government. When deployed, the model will only be able to access on-device conversations with the user's consent, in accordance with the terms of use of that application. As the model runs entirely on-device, neither the Singapore Government nor Google, will be able to access users’ conversations through its deployment.
Fighting scams shouldn't mean sacrificing personal privacy. Working with Google DeepMind, we're bringing advanced AI protection directly onto citizens' devices. This allows us to catch highly targeted, local scam attempts in real time while keeping personal communications private.
Under the hood: local context and edge efficiency
When the model detects manipulative conversational patterns, it triggers a real-time warning directly on the user's device advising them to pause and exercise caution before sharing sensitive information or funds.
Running a complex scam classifier model on modern smartphones is no easy task, with battery efficiency being a critical area of concern. To do this, we're relying on a two-stage pipeline:
Stage 1: A lightweight, efficient classifier screens inbound messages, using minimal power, focused on catching potential threats.
Stage 2: In-context analysis with our fine-tuned Gemma model, where we analyze a larger snapshot of the conversation to confirm a threat or discard a false alarm.
Presented at Singapore International Cyber Week (SICW 2026), this collaboration provides a model for messaging platforms balancing user protection with end-to-end encryption. Strict governance was maintained throughout development:
Sovereign Sandbox: Fine-tuning took place within GovTech Singapore's managed and maintained GCP environment, keeping training data isolated.
No Citizen Data: Training strictly excluded private or personally identifiable information, utilising anonymised datasets derived from scam reports and intelligence feeds.
Stress-Testing: The model was tested against synthetic dialogues and local corpora to cover diverse scam categories and evolving impersonation tactics.
The future of on-device safety
The project demonstrates how AI can help counter criminality without compromising safety or privacy. It also shows how governments and model providers can work together to build tech for public good. With this lightweight, on-device model, we've proven the ability for local, real-time scam detection while ensuring chat interactions remain secure. We look forward to sharing more soon on our collaboration with the Singapore government, as well as exploring the potential for similar technology to be deployed in other countries.
Strengthening Singapore’s AI Future: A New National Partnership
Unlocking UK house-building with AI-accelerated planning
Announcing our partnership with the Republic of Korea
Accelerating discovery in India through AI-powered science and education
Strengthening our partnership with the UK government to support prosperity and security in the AI era
Deepening our partnership with the UK AI Security Institute
Google DeepMind supports U.S. Department of Energy on Genesis: a national mission to accelerate innovation and scientific discovery
Follow us Sign up for updates on our latest innovations I accept Google's Terms and Conditions and acknowledge that my information will be used in accordance with Google's Privacy Policy .
