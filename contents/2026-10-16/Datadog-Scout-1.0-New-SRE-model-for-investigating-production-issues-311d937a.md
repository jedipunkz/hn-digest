---
source: "https://www.datadoghq.com/blog/ai/introducing-scout-sre-specialized-model-bits-ai/"
hn_url: "https://news.ycombinator.com/item?id=50111712"
title: "Datadog Scout 1.0: New SRE model for investigating production issues"
article_title: "Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost | Datadog"
image: "https://web-assets.dd-static.net/42588/1792073099-pareto_hero-v3.png"
author: "wavelander"
captured_at: "2026-10-16T00:12:43Z"
capture_tool: "hn-digest"
hn_id: 50111712
score: 1
comments: 1
posted_at: "2026-10-15T23:18:23Z"
tags:
  - hacker-news
---

# Datadog Scout 1.0: New SRE model for investigating production issues

- HN: [50111712](https://news.ycombinator.com/item?id=50111712)
- Source: [www.datadoghq.com](https://www.datadoghq.com/blog/ai/introducing-scout-sre-specialized-model-bits-ai/)
- Score: 1
- Comments: 1
- Posted: 2026-10-15T23:18:23Z

## Translation

Title: Datadog Scout 1.0: New SRE model for investigating production issues
Article title: Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost | Datadog
Description: Scout 1.0 is Datadog’s first AI model specialized for investigating production issues while achieving a better price-performance curve on our target scenarios when compared with general-purpose frontier models.

Article text:
Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost | Datadog
Datadog named a Leader in the Gartner® Magic Quadrant™ for Observability Platforms Leader in the Gartner® Magic Quadrant™ 30;
},
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'none';
}
},
checkAnnouncementBanner() {
const announcementBanner = document.querySelector('.announcement-banner') || document.querySelector('.announcement-banner--large');
if (announcementBanner) {
this.hasAnnouncementBanner = true;
} else {
this.hasAnnouncementBanner = false;
}
}
}" x-init="checkAnnouncementBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> Product {
this.openCategory = category;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`Product Category ${category} Hover`)
}
})
}, 160);
},
clearCategory() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
var observer = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = 'observability';
}
});
});
observer.observe(menu, { attributes: true });
"> The integrated platform for monitoring & security
dashboard Platform Capabilities
End-to-end, simplified visibility into your stack’s health & performance
Application Performance Monitoring
Bring Your Own Cloud Log Management
Outpace AI-powered attacks with unified security and observability
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Optimize front-end performance and enhance user experiences
Build, test, secure and ship quality code faster
Internal Developer Portal (IDP)
Integrated, streamlined workflows for faster time-to-resolution
Internal Developer Portal (IDP)
Monitor and improve model performance. Pinpoint root causes and detect anomalies
Built-in features & integrations that power the Datadog platform
Amazon Web Services Monitoring
Product host-map Infrastructure
Application Performance Monitoring
Bring Your Own Cloud Log Management
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Internal Developer Portal (IDP)
Internal Developer Portal (IDP)
dashboard Platform Capabilities
Amazon Web Services Monitoring
AI Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost
Today, we’re introducing Scout 1.0, Datadog’s first AI model specialized for investigating production issues. Scout powers Bits Investigation , our AI SRE. We trained Scout to reduce inference cost so that investigations are cheap enough to run an order of magnitude more often and earlier in an incident’s lifecycle.
Scout achieves a better price-performance curve on our target scenarios when compared with general-purpose frontier models (Figure 1). Scout 1.0 medium matches GPT-6.1 Sol high and Claude Opus 5.5 high at 3–8× lower estimated inference cost, and performs comparably to Fable medium and Astra medium at 12× and 19× lower cost, respectively. With test-time scaling, Scout 1.0 xhigh matches Fable xhigh at 4× lower cost.
We built Scout using our in-house distributed reinforcement learning infrastructure, and relied extensively on real historical incidents and telemetry data to build high-quality training simulations that don’t contain any customer data. For more technical details, see our full report .
Datadog engineering teams and a small group of customers in a closed preview are already using Scout to investigate production alerts. Sign up for the preview to try it in your environment and provide feedback before general availability.
Figure 1: Investigation quality versus estimated inference cost on AlertEval, our internal human-expert-annotated investigation benchmark. In its medium-effort mode, Scout achieves a score comparable to GPT-6.1 Sol and Claude Opus 5.5, at 3×–8× lower inference cost. Our highest-effort mode, Scout 1.0 xhigh, achieves state-of-the-art accuracy that outperforms GPT-6 Astra and matches Claude Fable 5.1 xhigh while keeping costs in the same range as Sol and Opus.* *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Figure 1: Investigation quality versus estimated inference cost on AlertEval, our internal human-expert-annotated investigation benchmark. In its medium-effort mode, Scout achieves a score comparable to GPT-6.1 Sol and Claude Opus 5.5, at 3×–8× lower inference cost. Our highest-effort mode, Scout 1.0 xhigh, achieves state-of-the-art accuracy that outperforms GPT-6 Astra and matches Claude Fable 5.1 xhigh while keeping costs in the same range as Sol and Opus.* Close dialog
Bits as a first responder for every alert
Our goal is for Bits to serve as an always-on first responder, gathering evidence and identifying likely causes before an engineer steps in. Each Bits investigation involves forming hypotheses, querying telemetry, and deciding whether the evidence supports a conclusion (Figure 2). Those steps require repeated model calls. By reducing a model’s inference cost, we can give teams more room to run investigations across their environments. Scout’s efficiency makes it practical to extend support to more alerts and investigate complex issues in greater depth, enabling greater breadth of coverage, reducing the burden of on-call work, and freeing more time for development.
Lower inference cost also makes it practical to spend additional computation on harder or higher-impact problems. For instance, Scout 1.0 xhigh uses more computation per incident to improve investigation quality, and this added investigative depth reduces engineering workload and pressure during even the highest-stakes incidents.
Figure 2: Bits Investigation presents Scout’s diagnosis and proposed recovery actions for a Postgres incident. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Figure 2: Bits Investigation presents Scout’s diagnosis and proposed recovery actions for a Postgres incident. Close dialog
Training and evaluating Scout for SRE investigations
We built Scout 1.0 on GLM-5.3 Flash and post-trained it using reinforcement learning in investigation simulations. We use historical incidents and telemetry from real alerts to build realistic simulations with known causes. Scout practices using tools modeled on Bits Investigation to gather evidence and test competing explanations. Its training rewards finding the cause and supporting its conclusions with evidence.
We also built our own training infrastructure on top of open source tools, with the simulation, evaluation, and model-training systems needed to run reinforcement learning at scale.
To test whether those skills transfer to real incidents, we built AlertEval. Datadog engineers have curated thousands of evaluation points through our incident-response workflow . AlertEval draws on that work, combining real alerts, archived telemetry, and root cause explanations written by engineers.
We are building Scout and Bits Investigation together, training the model for the tools Bits uses and improving those tools based on what we learn from its investigations.
We have also evaluated Scout on SREGym , a third-party benchmark, to test its performance outside AlertEval. We again have found that Scout 1.0 medium significantly improves upon the base model, GLM-5.3-Flash, performance, surpassing GLM-5.3 at a more than 4× reduction in cost. Through additional test time scaling, we found that Scout 1.0 high is competitive with the best models on this benchmark, while still maintaining at least 2× cost savings compared to all competitors (10× savings compared to Astra/Fable-class models).
Moving forward, we plan to simulate a wider range of failures and keep improving investigation quality while lowering inference cost. We also plan to extend this approach beyond investigation to help teams autonomously remediate issues to restore service faster. Together, these advances move Bits beyond its role as a first responder and toward becoming the only responder during an incident, resolving issues faster and reducing on-call burden.
Join the Scout 1.0 Bits Investigation preview
Preview customers are already using Bits Investigation powered by Scout on their production alerts. Sign up for the preview to try it in your environment and provide feedback before general availability.
Thanks to everyone who helped build Scout; see our technical report for the full author list and acknowledgments.
* The scores measure agreement with expert reference explanations on a 0–100 scale; they are not percentages of incidents solved. Costs are estimated inference costs per investigation, relative to Scout 1.0 medium. Lower cost appears to the right.
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-black">
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-[#FF4500]">
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-[#0A66C2]">
Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Bits as a first responder for every alert
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Training and evaluating Scout for SRE investigations
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Join the Scout 1.0 Bits Investigation preview
Did you find this article helpful?
Article thumbs down
Further reading
Datadog has been named a leader in the 2026 Gartner® Magic Quadrant™ for Observability Platforms
We're always looking for talented people to collaborate with
Start monitoring your metrics in minutes
Product Observability Infrastructure Infrastructure Monitoring
Applications Application Performance Monitoring
Bring Your Own Cloud Log Management
security-platform Security Code Security Code Security
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
rum Digital Experience Digital Experience Product Analytics
ci Software Delivery Software Delivery Continuous Testing
Internal Developer Portal (IDP)
Service Management Incident Response Incident Management
Context Graph Software Catalog
AI AI for Datadog Bits Investigation
Datadog for AI Agent Observability
dashboard Platform Capabilities Bui

[truncated]

## Original Extract

Scout 1.0 is Datadog’s first AI model specialized for investigating production issues while achieving a better price-performance curve on our target scenarios when compared with general-purpose frontier models.

Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost | Datadog
Datadog named a Leader in the Gartner® Magic Quadrant™ for Observability Platforms Leader in the Gartner® Magic Quadrant™ 30;
},
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'none';
}
},
checkAnnouncementBanner() {
const announcementBanner = document.querySelector('.announcement-banner') || document.querySelector('.announcement-banner--large');
if (announcementBanner) {
this.hasAnnouncementBanner = true;
} else {
this.hasAnnouncementBanner = false;
}
}
}" x-init="checkAnnouncementBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> Product {
this.openCategory = category;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`Product Category ${category} Hover`)
}
})
}, 160);
},
clearCategory() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
var observer = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = 'observability';
}
});
});
observer.observe(menu, { attributes: true });
"> The integrated platform for monitoring & security
dashboard Platform Capabilities
End-to-end, simplified visibility into your stack’s health & performance
Application Performance Monitoring
Bring Your Own Cloud Log Management
Outpace AI-powered attacks with unified security and observability
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Optimize front-end performance and enhance user experiences
Build, test, secure and ship quality code faster
Internal Developer Portal (IDP)
Integrated, streamlined workflows for faster time-to-resolution
Internal Developer Portal (IDP)
Monitor and improve model performance. Pinpoint root causes and detect anomalies
Built-in features & integrations that power the Datadog platform
Amazon Web Services Monitoring
Product host-map Infrastructure
Application Performance Monitoring
Bring Your Own Cloud Log Management
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Internal Developer Portal (IDP)
Internal Developer Portal (IDP)
dashboard Platform Capabilities
Amazon Web Services Monitoring
AI Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost
Today, we’re introducing Scout 1.0, Datadog’s first AI model specialized for investigating production issues. Scout powers Bits Investigation , our AI SRE. We trained Scout to reduce inference cost so that investigations are cheap enough to run an order of magnitude more often and earlier in an incident’s lifecycle.
Scout achieves a better price-performance curve on our target scenarios when compared with general-purpose frontier models (Figure 1). Scout 1.0 medium matches GPT-6.1 Sol high and Claude Opus 5.5 high at 3–8× lower estimated inference cost, and performs comparably to Fable medium and Astra medium at 12× and 19× lower cost, respectively. With test-time scaling, Scout 1.0 xhigh matches Fable xhigh at 4× lower cost.
We built Scout using our in-house distributed reinforcement learning infrastructure, and relied extensively on real historical incidents and telemetry data to build high-quality training simulations that don’t contain any customer data. For more technical details, see our full report .
Datadog engineering teams and a small group of customers in a closed preview are already using Scout to investigate production alerts. Sign up for the preview to try it in your environment and provide feedback before general availability.
Figure 1: Investigation quality versus estimated inference cost on AlertEval, our internal human-expert-annotated investigation benchmark. In its medium-effort mode, Scout achieves a score comparable to GPT-6.1 Sol and Claude Opus 5.5, at 3×–8× lower inference cost. Our highest-effort mode, Scout 1.0 xhigh, achieves state-of-the-art accuracy that outperforms GPT-6 Astra and matches Claude Fable 5.1 xhigh while keeping costs in the same range as Sol and Opus.* *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Figure 1: Investigation quality versus estimated inference cost on AlertEval, our internal human-expert-annotated investigation benchmark. In its medium-effort mode, Scout achieves a score comparable to GPT-6.1 Sol and Claude Opus 5.5, at 3×–8× lower inference cost. Our highest-effort mode, Scout 1.0 xhigh, achieves state-of-the-art accuracy that outperforms GPT-6 Astra and matches Claude Fable 5.1 xhigh while keeping costs in the same range as Sol and Opus.* Close dialog
Bits as a first responder for every alert
Our goal is for Bits to serve as an always-on first responder, gathering evidence and identifying likely causes before an engineer steps in. Each Bits investigation involves forming hypotheses, querying telemetry, and deciding whether the evidence supports a conclusion (Figure 2). Those steps require repeated model calls. By reducing a model’s inference cost, we can give teams more room to run investigations across their environments. Scout’s efficiency makes it practical to extend support to more alerts and investigate complex issues in greater depth, enabling greater breadth of coverage, reducing the burden of on-call work, and freeing more time for development.
Lower inference cost also makes it practical to spend additional computation on harder or higher-impact problems. For instance, Scout 1.0 xhigh uses more computation per incident to improve investigation quality, and this added investigative depth reduces engineering workload and pressure during even the highest-stakes incidents.
Figure 2: Bits Investigation presents Scout’s diagnosis and proposed recovery actions for a Postgres incident. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Figure 2: Bits Investigation presents Scout’s diagnosis and proposed recovery actions for a Postgres incident. Close dialog
Training and evaluating Scout for SRE investigations
We built Scout 1.0 on GLM-5.3 Flash and post-trained it using reinforcement learning in investigation simulations. We use historical incidents and telemetry from real alerts to build realistic simulations with known causes. Scout practices using tools modeled on Bits Investigation to gather evidence and test competing explanations. Its training rewards finding the cause and supporting its conclusions with evidence.
We also built our own training infrastructure on top of open source tools, with the simulation, evaluation, and model-training systems needed to run reinforcement learning at scale.
To test whether those skills transfer to real incidents, we built AlertEval. Datadog engineers have curated thousands of evaluation points through our incident-response workflow . AlertEval draws on that work, combining real alerts, archived telemetry, and root cause explanations written by engineers.
We are building Scout and Bits Investigation together, training the model for the tools Bits uses and improving those tools based on what we learn from its investigations.
We have also evaluated Scout on SREGym , a third-party benchmark, to test its performance outside AlertEval. We again have found that Scout 1.0 medium significantly improves upon the base model, GLM-5.3-Flash, performance, surpassing GLM-5.3 at a more than 4× reduction in cost. Through additional test time scaling, we found that Scout 1.0 high is competitive with the best models on this benchmark, while still maintaining at least 2× cost savings compared to all competitors (10× savings compared to Astra/Fable-class models).
Moving forward, we plan to simulate a wider range of failures and keep improving investigation quality while lowering inference cost. We also plan to extend this approach beyond investigation to help teams autonomously remediate issues to restore service faster. Together, these advances move Bits beyond its role as a first responder and toward becoming the only responder during an incident, resolving issues faster and reducing on-call burden.
Join the Scout 1.0 Bits Investigation preview
Preview customers are already using Bits Investigation powered by Scout on their production alerts. Sign up for the preview to try it in your environment and provide feedback before general availability.
Thanks to everyone who helped build Scout; see our technical report for the full author list and acknowledgments.
* The scores measure agreement with expert reference explanations on a 0–100 scale; they are not percentages of incidents solved. Costs are estimated inference costs per investigation, relative to Scout 1.0 medium. Lower cost appears to the right.
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-black">
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-[#FF4500]">
g>rect]:fill-grey-500 [&_path]:!fill-white w-7 h-7 tablet:w-8 tablet:h-8 desktop-sm:w-9 desktop-sm:h-9 group-hover:[&>g>rect]:fill-[#0A66C2]">
Introducing Scout 1.0: A new SRE model powering Bits AI with frontier performance at a fraction of the cost
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Bits as a first responder for every alert
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Training and evaluating Scout for SRE investigations
&]:border-purple-600 active:border-purple-600 [.active>&]:text-purple-600 group-has-[.active]:text-purple-600 group-has-[.active]:border-purple-600 active:text-purple-800 [.active>&]:font-bold group-has-[.active]:font-bold group-has-[.active]:hover:border-purple-700 group-has-[.active]:hover:text-purple-700 [.active]:hover:border-purple-700 [.active>&]:hover:text-purple-700 [.active]:active:border-purple-800 [.active>&]:active:text-purple-800">
Join the Scout 1.0 Bits Investigation preview
Did you find this article helpful?
Article thumbs down
Further reading
Datadog has been named a leader in the 2026 Gartner® Magic Quadrant™ for Observability Platforms
We're always looking for talented people to collaborate with
Start monitoring your metrics in minutes
Product Observability Infrastructure Infrastructure Monitoring
Applications Application Performance Monitoring
Bring Your Own Cloud Log Management
security-platform Security Code Security Code Security
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
rum Digital Experience Digital Experience Product Analytics
ci Software Delivery Software Delivery Continuous Testing
Internal Developer Portal (IDP)
Service Management Incident Response Incident Management
Context Graph Software Catalog
AI AI for Datadog Bits Investigation
Datadog for AI Agent Observability
dashboard Platform Capabilities Bui

[truncated]
