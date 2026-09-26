---
source: "https://news.ufl.edu/2026/09/ai-powered-vehicles/"
hn_url: "https://news.ycombinator.com/item?id=49860649"
title: "Simple visual patterns can trick AI-powered vehicles and robots"
article_title: "Simple visual patterns can trick AI-powered vehicles and robots, UF research finds News | University of Florida"
image: "https://news.ufl.edu/media/newsufledu/images/2026/09/2026-09-03_Stereo-Vision-Vulnerabilities---Sarah-Rampazzi-square.jpg"
author: "geox"
captured_at: "2026-09-26T21:53:02Z"
capture_tool: "hn-digest"
hn_id: 49860649
score: 1
comments: 0
posted_at: "2026-09-26T21:20:19Z"
tags:
  - hacker-news
---

# Simple visual patterns can trick AI-powered vehicles and robots

- HN: [49860649](https://news.ycombinator.com/item?id=49860649)
- Source: [news.ufl.edu](https://news.ufl.edu/2026/09/ai-powered-vehicles/)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T21:20:19Z

## Translation

Title: Simple visual patterns can trick AI-powered vehicles and robots
Article title: Simple visual patterns can trick AI-powered vehicles and robots, UF research finds News | University of Florida
Description: A simple pattern of black-and-white stripes could cause an autonomous vehicle or robot to misjudge how far away an obstacle is, potentially triggering an unexpected maneuver or even a collision, according to new University of Florida research.

Article text:
Simple visual patterns can trick AI-powered vehicles and robots, UF research finds News | University of Florida
Skip to main content
News
School Logo Link
MENU
University of Florida
For Faculty
Faculty Resources
News Categories
Agriculture + Environment
Simple visual patterns can trick AI-powered vehicles and robots, UF research finds
Simple visual patterns can trick AI-powered vehicles and robots, UF research finds
Karen Dooley
September 22, 2026
A simple pattern of black-and-white stripes could cause an autonomous vehicle or robot to misjudge how far away an obstacle is, potentially triggering an unexpected maneuver or even a collision, according to new University of Florida research.
The finding exposes a previously overlooked vulnerability in the technology that autonomous systems use to perceive depth — a critical capability for vehicles, drones and robots navigating their surroundings.
Researchers found that repeated patterns in the environment can interfere with the algorithms and artificial intelligence models that estimate the distance between a camera and an object. By manipulating the pattern, an attacker could make an obstacle appear significantly closer or farther away than it really is.
But the problem doesn’t necessarily require an attacker, said Sara Rampazzi, Ph.D., an assistant professor in the department of computer and information science and engineering at UF.
“These things can also happen naturally, so it’s not a matter of imagining a sophisticated attacker. It becomes a safety problem,” said Rampazzi, lead of the research effort.
The research focuses on autonomous systems that use stereo cameras — two cameras positioned close together — to estimate depth. These systems can use specific algorithms or artificial intelligence to make the same calculation from the images captured by the cameras.
Although the approaches differ, the researchers found they share an underlying weakness. Repeated visual patterns, like stripes on a fence or other regularly spaced objects, can cause the systems to incorrectly calculate depth.
Rampazzi’s team, in collaboration with two universities in Japan, tested multiple sensors, algorithms and AI models and found the same fundamental vulnerability across them, suggesting the problem is not limited to a particular manufacturer or application.
In a driving scenario, for example, researchers projected a simple checkerboard-like pattern onto the back of a vehicle. The pattern caused the autonomous vehicle’s perception system to interpret part of the vehicle as being closer than it actually was and enough to trigger an automatic response such as braking.
The same type of vulnerability could affect drones and ground robots, which rely on depth perception to avoid buildings, vehicles, birds and other obstacles while following an assigned route.
“For a self-driving vehicle, an unexpected hard brake could create a hazard for vehicles following behind it,” Rampazzi said. “For a drone or robot, an incorrect depth estimate could cause a sudden maneuver or collision.”
The researchers also demonstrated that an attacker would not necessarily need to cover an entire scene with a pattern. A small strategically placed pattern can be enough to influence the depth calculation.
The study will be presented in November at the Association for Computing Machinery Conference on Computer and Communications Safety, or ACM CCS , a leading cybersecurity conference.
Rampazzi said the research team also developed defenses designed to address the underlying cause of the vulnerability rather than simply training AI models with more examples of problematic patterns.
For traditional depth-estimation algorithms, the researchers developed a software-based change that can recognize situations involving repeated patterns and prevent the system from selecting an incorrect depth calculation. For AI-based systems, they modified the models, so they can respond differently when those patterns appear.
The researchers tested their approach using simulations as well as a real vehicle and sensors at UF’s controlled autonomous-vehicle testing facility.
Rampazzi said the work illustrates a broader challenge as AI becomes increasingly responsible for interpreting the physical world.
“We try to find and resolve the root cause of the problem, instead of just adding more data on it, like having a more fancy AI model on it,” she said.
Rather than simply teaching an AI system to recognize every possible problematic pattern, she said, the researchers focus on addressing weaknesses inherited from the sensing and depth-estimation systems underneath the AI to make the technology more resilient to both deliberate attacks and unexpected conditions in the real world.
University of Florida,
Gainesville, FL 32611

## Original Extract

A simple pattern of black-and-white stripes could cause an autonomous vehicle or robot to misjudge how far away an obstacle is, potentially triggering an unexpected maneuver or even a collision, according to new University of Florida research.

Simple visual patterns can trick AI-powered vehicles and robots, UF research finds News | University of Florida
Skip to main content
News
School Logo Link
MENU
University of Florida
For Faculty
Faculty Resources
News Categories
Agriculture + Environment
Simple visual patterns can trick AI-powered vehicles and robots, UF research finds
Simple visual patterns can trick AI-powered vehicles and robots, UF research finds
Karen Dooley
September 22, 2026
A simple pattern of black-and-white stripes could cause an autonomous vehicle or robot to misjudge how far away an obstacle is, potentially triggering an unexpected maneuver or even a collision, according to new University of Florida research.
The finding exposes a previously overlooked vulnerability in the technology that autonomous systems use to perceive depth — a critical capability for vehicles, drones and robots navigating their surroundings.
Researchers found that repeated patterns in the environment can interfere with the algorithms and artificial intelligence models that estimate the distance between a camera and an object. By manipulating the pattern, an attacker could make an obstacle appear significantly closer or farther away than it really is.
But the problem doesn’t necessarily require an attacker, said Sara Rampazzi, Ph.D., an assistant professor in the department of computer and information science and engineering at UF.
“These things can also happen naturally, so it’s not a matter of imagining a sophisticated attacker. It becomes a safety problem,” said Rampazzi, lead of the research effort.
The research focuses on autonomous systems that use stereo cameras — two cameras positioned close together — to estimate depth. These systems can use specific algorithms or artificial intelligence to make the same calculation from the images captured by the cameras.
Although the approaches differ, the researchers found they share an underlying weakness. Repeated visual patterns, like stripes on a fence or other regularly spaced objects, can cause the systems to incorrectly calculate depth.
Rampazzi’s team, in collaboration with two universities in Japan, tested multiple sensors, algorithms and AI models and found the same fundamental vulnerability across them, suggesting the problem is not limited to a particular manufacturer or application.
In a driving scenario, for example, researchers projected a simple checkerboard-like pattern onto the back of a vehicle. The pattern caused the autonomous vehicle’s perception system to interpret part of the vehicle as being closer than it actually was and enough to trigger an automatic response such as braking.
The same type of vulnerability could affect drones and ground robots, which rely on depth perception to avoid buildings, vehicles, birds and other obstacles while following an assigned route.
“For a self-driving vehicle, an unexpected hard brake could create a hazard for vehicles following behind it,” Rampazzi said. “For a drone or robot, an incorrect depth estimate could cause a sudden maneuver or collision.”
The researchers also demonstrated that an attacker would not necessarily need to cover an entire scene with a pattern. A small strategically placed pattern can be enough to influence the depth calculation.
The study will be presented in November at the Association for Computing Machinery Conference on Computer and Communications Safety, or ACM CCS , a leading cybersecurity conference.
Rampazzi said the research team also developed defenses designed to address the underlying cause of the vulnerability rather than simply training AI models with more examples of problematic patterns.
For traditional depth-estimation algorithms, the researchers developed a software-based change that can recognize situations involving repeated patterns and prevent the system from selecting an incorrect depth calculation. For AI-based systems, they modified the models, so they can respond differently when those patterns appear.
The researchers tested their approach using simulations as well as a real vehicle and sensors at UF’s controlled autonomous-vehicle testing facility.
Rampazzi said the work illustrates a broader challenge as AI becomes increasingly responsible for interpreting the physical world.
“We try to find and resolve the root cause of the problem, instead of just adding more data on it, like having a more fancy AI model on it,” she said.
Rather than simply teaching an AI system to recognize every possible problematic pattern, she said, the researchers focus on addressing weaknesses inherited from the sensing and depth-estimation systems underneath the AI to make the technology more resilient to both deliberate attacks and unexpected conditions in the real world.
University of Florida,
Gainesville, FL 32611
