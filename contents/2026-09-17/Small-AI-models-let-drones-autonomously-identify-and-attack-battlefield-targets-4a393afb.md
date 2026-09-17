---
source: "https://arstechnica.com/ai/2026/09/nato-backed-startup-adapts-ai-for-autonomous-drone-recon-and-attack-missions/"
hn_url: "https://news.ycombinator.com/item?id=49747600"
title: "Small AI models let drones autonomously identify and attack battlefield targets"
article_title: "Small AI models let drones autonomously identify and attack battlefield targets - Ars Technica"
image: "https://cdn.arstechnica.net/wp-content/uploads/2026/09/Scaleout-Systems-ALMA-demo-1152x648.png"
author: "pseudolus"
captured_at: "2026-09-17T23:24:50Z"
capture_tool: "hn-digest"
hn_id: 49747600
score: 1
comments: 0
posted_at: "2026-09-17T22:33:44Z"
tags:
  - hacker-news
---

# Small AI models let drones autonomously identify and attack battlefield targets

- HN: [49747600](https://news.ycombinator.com/item?id=49747600)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/09/nato-backed-startup-adapts-ai-for-autonomous-drone-recon-and-attack-missions/)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T22:33:44Z

## Translation

Title: Small AI models let drones autonomously identify and attack battlefield targets
Article title: Small AI models let drones autonomously identify and attack battlefield targets - Ars Technica
Description: Scaleout deploys decentralized AI-driven learning to military bases and drones.

Article text:
Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Federated learning on the battlefield
Small AI models let drones autonomously identify and attack battlefield targets
Scaleout deploys decentralized AI-driven learning to military bases and drones.
9
Demonstration of Scaleout Systems' AI-enabled target recognition and engagement capability for an autonomous attack drone under the ALMA project led by BAE Systems Bofors.
Credit:
Scaleout Systems
Demonstration of Scaleout Systems' AI-enabled target recognition and engagement capability for an autonomous attack drone under the ALMA project led by BAE Systems Bofors.
Credit:
Scaleout Systems
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
As European militaries adapt to the use of AI and drones in modern warfare, a NATO-backed startup is helping to deploy AI-driven target detection and selection that can run on small drones for surveillance and attack missions.
The company Scaleout Systems was originally founded by researchers from Uppsala University in Sweden in 2018, and initially focused on training and deploying machine learning models directly on the hardware available in commercial trucks and other vehicles. But once Russia launched its full-scale invasion of Ukraine in 2022, the company pivoted toward defense applications.
“With the war in Ukraine and a shifting world, we realized that this technology can be very important to operationalize edge data and sensor data for machine learning to make sure that NATO allies have found that strategic advantage,” Andreas Hellander, cofounder and CEO of Scaleout Systems, told Ars.
Instead of using frontier AI models from OpenAI or Anthropic , Scaleout is harnessing leaner ones, such as machine learning models that can perform computer vision tasks on the hardware of drones or computers used at forward bases. “They need to fit on forward-deployed hardware and edge hardware, which can vary quite a bit from small embedded devices to quite powerful edge workstations,” Hellander said.
Scaleout was selected to join NATO’s Defence Innovator Accelerator for the North Atlantic (DIANA) Challenge Program in 2025. There, it has worked on the Federated Aerial Intelligence for Recon project to adapt machine learning models for the edge computing hardware found in drones, drone pilot tablets, and field command posts.
Such AI models can help the drone operators with tasks such as target identification, even as the drone’s cameras and sensors collect data from the surrounding battlefield environment. They can then intermittently share selective updates with computing nodes at the local platoon or company headquarters without transmitting sensitive raw data.
Those headquarters computing nodes help to retrain the AI models on the new battlefield data aggregated from multiple sources, before pushing the updated capabilities out to the edge devices when the opportunity arises.
“Models might have been trained in a desert environment, and if we try to deploy them in an urban environment, they’re not going to perform well,” Hellander told Ars. “If we can release several new versions of this model that—during the course of a single day or certainly an operation—keep learning and keep improving from this massive amount of sensor data that is generated at a practical edge, that is the sustainable advantage.”
Testing the technology in live exercises
The technology means military drones and devices could benefit from local AI models that operate independently without relying on continuous communications with a central server hosting larger AI models in a data center. Reliance on a centralized location to run AI models looks riskier at a time when large data centers have been targeted and destroyed during the war between the US and Iran .
This approach is also incredibly useful on modern battlefields where electronic warfare and enemy jamming can frequently interfere with communication signals. A growing number of Ukrainian military drones are already incorporating onboard AI capabilities into cheap kamikaze drones .
“We built a functioning concept of how we do this for a surveillance and reconnaissance system based on a drone,” Hellander told Ars. “This is something we have also done in public demonstrations in Sweden.”
Scaleout is participating in the Affordable Loitering Modular Ammunition (ALMA) project headed by BAE Systems Bofors that aims to develop a low-cost, autonomous kamikaze drone. The ALMA concept was first publicly demonstrated during a Winter Demo 2026 event held in Sweden in January.
That demonstration showed how a drone could autonomously “detect, identify and geolocate all potential spotted threats with the use of AI,” according to a Scaleout Systems presentation about ALMA’s capabilities. “All data is handled by dedicated onboard computing, allowing the system to perform in real-time without any external processing.”
Using its onboard AI capabilities, the drone automatically prioritized the highest-value target as defined by its mission—in this case, an armored engineering vehicle—and flew to that target to drop an explosive on it. A human operator could still control and direct the drone, but the drone carried out the mission on its own without direct human commands.
In June, Scaleout also tested its technology at a Swedish Air Force base in Uppsala, where the Swedish military already has a license to use Scaleout’s main software platform. That demonstration showed how a forward-deployed computing node at the military base could still run its own AI inference and active-learning processes after losing connection with a central computing node in Scaleout Systems’ lab. Once the connection was restored, the local AI model updates were shared with the central computing node.
This federated learning strategy, which allows the decentralized network of AI models to learn from aggregated data, could eventually scale across entire geographic regions or countries, Hellander explained. “In principle, you can unlock collaboration between NATO member states,” he said.
9 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Iran strikes on Amazon data centers caused permanent loss of customer data
2.
macOS 27 Golden Gate: The Ars Technica review
3.
After being sidelined, Boeing's Starliner to get starring role in NASA's spaceflight plans
4.
Apple reportedly building server packed with M-series Ultra chips for AI
5.
This Atlantic hurricane season is about to do something that hasn't happened in 175 years
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.

## Original Extract

Scaleout deploys decentralized AI-driven learning to military bases and drones.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Federated learning on the battlefield
Small AI models let drones autonomously identify and attack battlefield targets
Scaleout deploys decentralized AI-driven learning to military bases and drones.
9
Demonstration of Scaleout Systems' AI-enabled target recognition and engagement capability for an autonomous attack drone under the ALMA project led by BAE Systems Bofors.
Credit:
Scaleout Systems
Demonstration of Scaleout Systems' AI-enabled target recognition and engagement capability for an autonomous attack drone under the ALMA project led by BAE Systems Bofors.
Credit:
Scaleout Systems
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
As European militaries adapt to the use of AI and drones in modern warfare, a NATO-backed startup is helping to deploy AI-driven target detection and selection that can run on small drones for surveillance and attack missions.
The company Scaleout Systems was originally founded by researchers from Uppsala University in Sweden in 2018, and initially focused on training and deploying machine learning models directly on the hardware available in commercial trucks and other vehicles. But once Russia launched its full-scale invasion of Ukraine in 2022, the company pivoted toward defense applications.
“With the war in Ukraine and a shifting world, we realized that this technology can be very important to operationalize edge data and sensor data for machine learning to make sure that NATO allies have found that strategic advantage,” Andreas Hellander, cofounder and CEO of Scaleout Systems, told Ars.
Instead of using frontier AI models from OpenAI or Anthropic , Scaleout is harnessing leaner ones, such as machine learning models that can perform computer vision tasks on the hardware of drones or computers used at forward bases. “They need to fit on forward-deployed hardware and edge hardware, which can vary quite a bit from small embedded devices to quite powerful edge workstations,” Hellander said.
Scaleout was selected to join NATO’s Defence Innovator Accelerator for the North Atlantic (DIANA) Challenge Program in 2025. There, it has worked on the Federated Aerial Intelligence for Recon project to adapt machine learning models for the edge computing hardware found in drones, drone pilot tablets, and field command posts.
Such AI models can help the drone operators with tasks such as target identification, even as the drone’s cameras and sensors collect data from the surrounding battlefield environment. They can then intermittently share selective updates with computing nodes at the local platoon or company headquarters without transmitting sensitive raw data.
Those headquarters computing nodes help to retrain the AI models on the new battlefield data aggregated from multiple sources, before pushing the updated capabilities out to the edge devices when the opportunity arises.
“Models might have been trained in a desert environment, and if we try to deploy them in an urban environment, they’re not going to perform well,” Hellander told Ars. “If we can release several new versions of this model that—during the course of a single day or certainly an operation—keep learning and keep improving from this massive amount of sensor data that is generated at a practical edge, that is the sustainable advantage.”
Testing the technology in live exercises
The technology means military drones and devices could benefit from local AI models that operate independently without relying on continuous communications with a central server hosting larger AI models in a data center. Reliance on a centralized location to run AI models looks riskier at a time when large data centers have been targeted and destroyed during the war between the US and Iran .
This approach is also incredibly useful on modern battlefields where electronic warfare and enemy jamming can frequently interfere with communication signals. A growing number of Ukrainian military drones are already incorporating onboard AI capabilities into cheap kamikaze drones .
“We built a functioning concept of how we do this for a surveillance and reconnaissance system based on a drone,” Hellander told Ars. “This is something we have also done in public demonstrations in Sweden.”
Scaleout is participating in the Affordable Loitering Modular Ammunition (ALMA) project headed by BAE Systems Bofors that aims to develop a low-cost, autonomous kamikaze drone. The ALMA concept was first publicly demonstrated during a Winter Demo 2026 event held in Sweden in January.
That demonstration showed how a drone could autonomously “detect, identify and geolocate all potential spotted threats with the use of AI,” according to a Scaleout Systems presentation about ALMA’s capabilities. “All data is handled by dedicated onboard computing, allowing the system to perform in real-time without any external processing.”
Using its onboard AI capabilities, the drone automatically prioritized the highest-value target as defined by its mission—in this case, an armored engineering vehicle—and flew to that target to drop an explosive on it. A human operator could still control and direct the drone, but the drone carried out the mission on its own without direct human commands.
In June, Scaleout also tested its technology at a Swedish Air Force base in Uppsala, where the Swedish military already has a license to use Scaleout’s main software platform. That demonstration showed how a forward-deployed computing node at the military base could still run its own AI inference and active-learning processes after losing connection with a central computing node in Scaleout Systems’ lab. Once the connection was restored, the local AI model updates were shared with the central computing node.
This federated learning strategy, which allows the decentralized network of AI models to learn from aggregated data, could eventually scale across entire geographic regions or countries, Hellander explained. “In principle, you can unlock collaboration between NATO member states,” he said.
9 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Iran strikes on Amazon data centers caused permanent loss of customer data
2.
macOS 27 Golden Gate: The Ars Technica review
3.
After being sidelined, Boeing's Starliner to get starring role in NASA's spaceflight plans
4.
Apple reportedly building server packed with M-series Ultra chips for AI
5.
This Atlantic hurricane season is about to do something that hasn't happened in 175 years
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
