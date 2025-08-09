# Mahdaad Backend Problem Solving Test

Hello dear candidate,

Thank you for taking the time to participate in this evaluation.  
This test is designed to assess your backend engineering skills through a series of practical, real-world scenarios. Each task reflects challenges you may encounter in production systems, emphasizing architecture thinking, problem solving, and clarity in communication.

Before you begin, please take note of the following:

-   You are expected to solve each task in either **Node.js** or **Go**, using best practices for error handling, design, and code structure.
-   This is intentional. We want to see how you analyze the problem and choose an appropriate approach.
-   **Only a key part of the code is required**, not a full system. Focus on the core logic relevant to the challenge.
-   If you are unable or unwilling to complete any part of the test, kindly document the reason in a Reason.md file in the root of your project.
-   Organize each task in a separate branch. Once each task is complete, push it to its respective branch and share the repository in **GitHub** for review.

We wish you the best of luck!  
We look forward to seeing your ideas, code clarity, and architectural thinking in action.

## Challenge 1: Handling Slow or Unavailable External Services

**Scenario:**  
Your service depends on a third-party API for sending SMS. Sometimes, the external API becomes slow or unavailable, causing your system to lag or fail under load.

**Task:**  
Design a solution where:

-   Your system stays responsive even if the external API fails temporarily
-   You prevent overloading the API with repeated retries
-   Users receive appropriate feedback

**Deliverable:**  
Write a code snippet (in Node.js or Go) that shows how you handle the communication with the external service, including error handling and fallback logic. A complete system is not necessary. Focus on the critical logic.

**Solution:**
I have implemented a retry pattern using an external library to handle temporary failures. While a queue is another useful pattern for this scenario, I am saving its implementation for Task 4.

## Challenge 2: Multi-Step Operation Across Independent Services

**Scenario:**  
You need to implement the following flow:

1. Create an order
2. Deduct inventory
3. Process payment

These services are fully decoupled. If any step fails, all previous actions must be undone.

**Task:**  
Design a resilient, decoupled structure to perform this flow. Ensure that partial failures trigger proper rollback or compensation logic.

**Deliverable:**  
Provide a code snippet (in Node.js or Go) that shows how you control this flow, including how failure and rollback are handled. No need to implement the services, just show the control logic.

**Solution:**
I'm using the Saga pattern to solve this problem, which ensures data consistency by handling rollbacks across distributed services.

## Challenge 3: Triggering Side Effects Without Tight Coupling

**Scenario:**  
When a new course is created in your system, the following actions should happen:

-   Send a notification email to users
-   Update the admin dashboard
-   Index the course in the internal search system

These actions must happen **without tightly coupling them** to the course creation logic or blocking the main flow.

**Task:**  
Propose a structure that allows these actions to happen in a decoupled, resilient manner.

**Deliverable:**  
Write a code snippet (in Node.js or Go) that demonstrates how an event is published and consumed by another component. It can be a minimal in-memory or simulated setup; the focus is on event handling.

**Solution:**
I implemented a solution for this problem using the pub/sub pattern, where a Go channel functions as a lightweight message broker.

## Challenge 4: Keeping Distributed Data in Sync Over Time

**Scenario:**  
User data exists across multiple independent services. When a user updates their profile, the changes should eventually propagate to all relevant services, even if some are temporarily down.

**Task:**  
Design a mechanism to ensure **eventual synchronization** of data, even in the presence of temporary service failures or delivery delays.

**Deliverable:**  
Write a code snippet that shows how you queue or retry an update operation until it reaches the destination service successfully. Keep it simple and focused.
