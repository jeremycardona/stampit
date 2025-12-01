# Phase 2: Define - Brews N' Bites Reward System

⚠️ **ASSUMPTION-BASED DESIGN EXERCISE**  
Problem statements and design decisions in this document are based on proto-personas and have NOT been validated with actual user research. All assumptions must be tested before implementation.

---

## Core Problem Statement

**"Brews N' Bites customers struggle with a physical stamp card system that gets lost, fails to recognize their full spending value, and doesn't accommodate different visit patterns or technology preferences, resulting in unredeemed rewards and missed opportunities to build stronger customer loyalty."**

---

## Actionable Problem Statements

### Problem Statement 1: Sarah Chen (Routine Commuter)

**Sarah Chen needs a frictionless way to earn and track rewards without carrying physical items because she's always in a rush during her morning commute and frequently forgets or loses her stamp card, causing her to lose progress toward free coffees.**

- **WHO:** Sarah Chen - busy professional, high-frequency customer
- **WHAT:** Frictionless reward tracking without physical cards
- **WHY:** Time-pressed routine means physical cards get forgotten/lost, progress is wasted

---

### Problem Statement 2: Marcus Johnson (Student Regular)

**Marcus Johnson needs transparent visibility into his reward progress and recognition for his visit frequency because he's budget-conscious and frustrated when less frequent visitors seem to get rewards at the same rate despite his loyalty to Brews N' Bites.**

- **WHO:** Marcus Johnson - student, very frequent visitor, budget-conscious
- **WHAT:** Transparent progress tracking and frequency-based recognition
- **WHY:** Wants to maximize value and feel his loyalty is acknowledged over casual visitors

---

### Problem Statement 3: Jennifer Williams (Brunch Enthusiast)

**Jennifer Williams needs a reward system that recognizes her full spending across both food and beverages because she regularly spends $40-50 on brunch visits but only receives credit for her $4 coffee, making her feel undervalued as a high-spending customer.**

- **WHO:** Jennifer Williams - high-value customer, moderate frequency
- **WHAT:** Rewards that credit all purchases, not just coffee
- **WHY:** Spends significantly more per visit but receives same reward as coffee-only customers

---

### Problem Statement 4: David Park (Inconsistent Visitor)

**David Park needs a reward system that accommodates irregular visit patterns and prevents reward expiration because his freelance schedule means he visits 2-3 times monthly and never completes stamp cards before losing them, making him feel loyalty programs aren't designed for customers like him.**

- **WHO:** David Park - freelancer, low-frequency visitor, irregular schedule
- **WHAT:** Flexible rewards that don't expire and work for irregular visitors
- **WHY:** Infrequent visits + card loss means he never reaches reward threshold

---

### Problem Statement 5: Elena Rodriguez (Neighborhood Retiree)

**Elena Rodriguez needs a reward system that works without requiring a smartphone while maintaining personal recognition because she doesn't own mobile devices but values being known as a regular, and she fears technology will replace the human connections she cherishes at Brews N' Bites.**

- **WHO:** Elena Rodriguez - very high-frequency visitor, non-tech user, community-focused
- **WHAT:** Non-digital reward option that preserves personal interaction
- **WHY:** Lacks smartphone but visits most frequently; values relationship over transaction

---

## User Needs Statements (How Might We Questions)

### For Sarah (Routine Commuter)
- HMW make reward tracking effortless and automatic?
- HMW eliminate the need to carry physical cards?
- HMW speed up the transaction process?

### For Marcus (Student Regular)
- HMW reward frequency and loyalty transparently?
- HMW make progress toward rewards visible and motivating?
- HMW offer value-based incentives for budget-conscious customers?

### For Jennifer (Brunch Enthusiast)
- HMW recognize spending beyond just coffee purchases?
- HMW reward high-value transactions appropriately?
- HMW make sure customers never miss earning rewards?

### For David (Inconsistent Visitor)
- HMW keep irregular customers engaged?
- HMW ensure rewards don't expire before completion?
- HMW remind customers to return without being annoying?

### For Elena (Neighborhood Retiree)
- HMW maintain personal connection while adding technology?
- HMW accommodate customers without smartphones?
- HMW preserve the community feel of the café?

---

## Priority User Needs (Ranked)

### Critical (Must-Have)
1. **Eliminate card loss** - Affects all users
2. **Track all purchases** - Not just coffee (Jennifer's pain point, but benefits all)
3. **Provide progress visibility** - Everyone wants to know where they stand
4. **Accommodate non-smartphone users** - Can't exclude Elena and similar customers

### High Priority (Should-Have)
5. **Automatic tracking** - Remove burden from staff and customers
6. **Flexible for different visit patterns** - Work for both Sarah and David
7. **Maintain at least current reward value** - 9 purchases = 1 free

### Medium Priority (Nice-to-Have)
8. **Personalization options** - Different rewards for different preferences
9. **Engagement reminders** - For irregular visitors
10. **Special offers** - Student discounts, birthday rewards, etc.

---

## Problem Prioritization Matrix

| Problem | Impact | Severity | Priority |
|---------|--------|----------|----------|
| Card loss/forgetting | HIGH (Sarah, David, Marcus) | HIGH | **CRITICAL** |
| Only coffee tracked | MEDIUM (Jennifer, but affects all) | HIGH | **HIGH** |
| No progress visibility | HIGH (All personas) | MEDIUM | **HIGH** |
| No non-digital option | LOW (Elena, elderly) | HIGH | **HIGH** |
| Irregular visit penalization | MEDIUM (David) | MEDIUM | **MEDIUM** |

---

## Design Principles

### Principle 1: "Digital First, But Never Digital Only"
- Primary solution should leverage technology for convenience
- Must have a fallback for non-digital customers
- Staff should be empowered to help all customers

### Principle 2: "Value All Spending, Not Just Coffee"
- Reward system should recognize brunch, food, and beverage purchases
- High spenders should feel appreciated
- Points/stamps should reflect actual customer value

### Principle 3: "Transparent Progress, Zero Friction"
- Customers should always know where they stand
- Earning rewards should be automatic and effortless
- No more "did you stamp my card?" moments

### Principle 4: "Inclusive Loyalty for All Visit Patterns"
- Daily visitors should feel valued for consistency
- Irregular visitors should still be able to earn rewards
- No punitive expiration policies

### Principle 5: "Preserve the Human Touch"
- Technology should enhance, not replace, personal service
- Staff should have tools to recognize and delight regulars
- Community atmosphere remains central

---

## Success Criteria

### User Experience Metrics
- 0% reward loss due to lost/forgotten cards
- 100% of purchases automatically tracked
- Customers can check reward status in under 5 seconds
- Works for both smartphone and non-smartphone users

### Business Metrics
- Increased reward redemption rate (currently lost due to card loss)
- Higher average transaction value (food + coffee bundling)
- Improved customer retention (especially irregular visitors)
- Reduced staff time spent managing stamps

### Qualitative Goals
- Customers feel valued regardless of visit frequency
- Staff find the system easier to use
- Community atmosphere is maintained or enhanced

---

## Constraints & Considerations

### Technical Constraints
- System must work offline (in case of internet outage)
- Must integrate with existing POS system (or work independently)
- Low implementation cost for independent shop
- Solution should leverage common POS integrations available in market

### Business Constraints
- Can't be less generous than current 9-to-1 ratio
- Must be easy for staff to learn and use
- Should encourage both coffee AND food sales

### User Constraints
- Must accommodate non-tech users
- Can't require extensive customer onboarding
- Privacy concerns (some may not want tracking)

---

## Out-of-Scope Items

The following items are explicitly OUT OF SCOPE for this project to prevent scope creep and maintain focus on the core reward system problem:

### Infrastructure Changes
- **Redesigning the physical layout of the café** - Focus is on reward mechanism only
- **Replacing or upgrading the core POS system** - Solution must integrate with existing POS hardware/software
- **Major kitchen or operational workflow changes** - Reward system should fit current operations

### Business Strategy Changes
- **Developing a new Brews N' Bites branding or marketing strategy** - Brand identity remains unchanged
- **Menu redesign or pricing strategy overhaul** - Current menu and pricing structure stays
- **Loyalty program features beyond rewards** - No gamification, social features, or complex tier systems in V1

### Technology Infrastructure
- **Building custom POS software from scratch** - Must use existing or compatible third-party solutions
- **Requiring customers to create accounts with passwords** - Keep friction as low as possible
- **Complex mobile app development** - If digital solution is needed, it should be lightweight (web-based preferred)

### Scope Boundaries
- **Multi-location support** - Brews N' Bites is a single location; system doesn't need to scale beyond this
- **Third-party integrations** (delivery apps, payment processors beyond current setup) - Focus on in-store experience
- **Customer relationship management (CRM) features** - No email marketing, customer analytics dashboards, etc. in initial version

---

## Critical Design Decisions

### Decision 1: Digital-First vs. No-Smartphone Option

**The Conflict:**  
Building a system that works seamlessly for tech-savvy users (Sarah, Marcus, David) while also serving non-digital users (Elena) adds complexity and cost.

**Our Decision:**  
**Design for non-smartphone users FIRST, then enhance with digital options.**

**Rationale:**
- If we design a solution that works perfectly for Elena without a smartphone, it will work for everyone
- Digital users can benefit from the same system while optionally using enhanced digital features
- This approach is more inclusive and aligns with Design Principle 5: "Preserve the Human Touch"
- A phone number or simple identifier at POS can serve both digital and non-digital customers equally

**Implementation Approach:**
- Core system works via POS lookup (phone number, name, or simple code)
- Staff can always access and update customer rewards at register
- Digital users can optionally check status via web or text, but it's not required
- This maintains one unified system instead of managing two parallel systems

**Trade-off Accepted:**  
We sacrifice some "automatic" convenience for digital users in exchange for universal accessibility and lower implementation complexity.

---

### Decision 2: Simplicity vs. Value Recognition

**The Conflict:**  
The simplest systems (stamp-based) don't recognize spending differences. The fairest systems (dollar-based points) are more complex to implement and explain.

**Our Decision:**  
**Implement a HYBRID approach: Point-based system with simple conversion rules.**

**Rationale:**
- Most modern POS systems already track transaction amounts, making dollar-based tracking feasible
- We can design simple, easy-to-understand conversion rules (e.g., "$1 = 1 point, 100 points = free coffee")
- This recognizes Jennifer's $50 brunch while remaining transparent for Marcus to track
- Complexity is handled by POS integration, not by customers or staff

**Implementation Approach:**
- Every dollar spent = points earned (simple 1:1 ratio)
- Clear redemption thresholds (e.g., 100 points = free coffee worth $4-5)
- Maintains approximately same generosity as current system (spend $40-45, get ~$4-5 free)
- Customers see points, not complex calculations
- Staff ring up normally, system auto-applies points

**Trade-off Accepted:**  
Slightly more complex setup/configuration in exchange for fairness and recognition of full customer value. However, day-to-day operation remains simple for both staff and customers.

---

### Decision 3: Integration Approach

**The Conflict:**  
Custom-built solutions offer maximum control but require development effort. Using existing POS loyalty features is easier but may limit our ability to implement our specific design principles.

**Our Decision:**  
**Build a custom reward system for Brews N' Bites that integrates with their specific POS system.**

**Rationale:**
- Focuses scope on solving Brews N' Bites' actual needs without generalizing for multiple businesses
- Allows us to implement our hybrid point-based approach exactly as designed
- Single POS integration is manageable and keeps development straightforward
- We maintain control over user experience and reward logic
- Can iterate and improve based on real customer feedback at one location

**Implementation Approach:**
- Identify Brews N' Bites' current POS system
- Develop lightweight reward system that integrates with their specific POS via available API/webhook
- System stores customer profiles, point balances, and reward history
- Staff can access customer rewards via simple interface at the register
- Customer lookup via phone number as primary identifier

**Technical Requirements:**
- Integration with Brews N' Bites' existing POS system
- Simple database to store customer accounts and point balances
- Staff interface accessible at point-of-sale (web-based or tablet)
- Optional customer-facing interface for balance checking (web or SMS)
- Minimal infrastructure - can start with basic cloud hosting

**Trade-off Accepted:**  
Building custom solution for one location in exchange for:
- Complete control over implementation of our design principles
- Ability to iterate quickly based on real feedback
- Solution tailored exactly to Brews N' Bites' workflow
- Simpler scope allows faster time to launch

**Future Consideration:**
If the system proves successful at Brews N' Bites, it could potentially be adapted for other coffee shops later. For now, focus remains on solving one business's problem well.

---

## Problem Validation

Validating our problem definition with each persona:

✓ **Sarah Chen** - "Yes, I lose cards and want something automatic"  
✓ **Marcus Johnson** - "Yes, I want to see my progress and feel rewarded for loyalty"  
✓ **Jennifer Williams** - "Yes, I want credit for my full spending, not just coffee"  
✓ **David Park** - "Yes, I need something that doesn't punish irregular visits"  
✓ **Elena Rodriguez** - "Yes, but I don't want to lose the personal connection"

**Result:** All personas validate that this is a real problem worth solving.

---

## Key Insights Summary

### What We Learned
1. **Loss is the biggest pain point** - Cards getting lost/forgotten affects almost everyone
2. **Value recognition matters** - Customers who spend more want proportional rewards
3. **Visibility creates motivation** - Not knowing progress reduces engagement
4. **One size doesn't fit all** - Visit patterns vary widely and need accommodation
5. **Technology can't exclude** - Digital solutions must have analog alternatives

### What We're Solving
**The Problem:** Physical stamp cards create friction, get lost, don't recognize full customer value, and fail to serve diverse customer needs.

**The Opportunity:** Create a reward system that's automatic, inclusive, value-based, and preserves the community feel of Brews N' Bites.

**The Goal:** Increase customer loyalty and lifetime value while making the experience better for both customers and staff.

---

## Next Steps

Moving forward to **Phase 3: Ideate**, we will:
1. Brainstorm multiple solution concepts
2. Explore different reward mechanisms
3. Consider various technology approaches
4. Evaluate trade-offs between solutions
5. Select the most promising concept to prototype

---

## Assumptions to Validate

⚠️ **Critical Assumptions Underlying This Define Phase:**

Before proceeding to implementation, these assumptions MUST be validated with real Brews N' Bites customers:

### Assumption 1: Problem Existence
**Assumption:** Customers actually experience significant pain with the physical stamp card system  
**Validation Needed:** 
- Survey: "How satisfied are you with the current rewards program?" (1-10 scale)
- Question: "How often have you lost your stamp card in the past year?"
- Observation: Track how many customers actually use stamp cards at checkout

### Assumption 2: Card Loss Frequency
**Assumption:** Card loss is a major, widespread problem affecting multiple customer segments  
**Validation Needed:**
- Staff interview: "How often do customers say they lost their card?"
- Customer interview: "Have you ever lost your Brews N' Bites stamp card?"
- Data: What percentage of issued cards are never redeemed?

### Assumption 3: Technology Adoption
**Assumption:** Most customers would prefer or accept a digital solution  
**Validation Needed:**
- Survey: "Would you prefer a digital rewards system over physical cards?"
- Demographics: What percentage of customers actually lack smartphones?
- Observation: How many customers currently use mobile payment?

### Assumption 4: Value Recognition Matters
**Assumption:** High-spending customers feel undervalued by the per-item stamp system  
**Validation Needed:**
- Interview brunch customers: "Do you feel the rewards program values your full spending?"
- Compare: Average spend of active vs. inactive stamp card users
- Question: "Would you visit more often if rewards recognized all purchases?"

### Assumption 5: Non-Digital Users Exist
**Assumption:** There's a significant segment (like Elena) without smartphones who need accommodation  
**Validation Needed:**
- Demographic survey: "Do you own a smartphone?"
- Staff observation: How many regular customers appear to be non-tech users?
- Reality check: Is Elena representative of actual customers or an outlier?

### Assumption 6: Progress Visibility Drives Motivation
**Assumption:** Customers want to see their reward progress and this motivates return visits  
**Validation Needed:**
- Interview: "Do you currently know how many stamps you have on your card?"
- Question: "Would seeing your progress toward a reward make you visit more often?"
- A/B test: Does showing progress increase visit frequency vs. not showing?

### Assumption 7: Irregular Visitors Feel Excluded
**Assumption:** Infrequent customers feel discouraged by loyalty programs designed for daily visitors  
**Validation Needed:**
- Interview irregular customers: "Why don't you visit more often?"
- Question: "Do you feel the stamp card system works for your visit pattern?"
- Data: What's the actual distribution of visit frequencies?

### Assumption 8: Current Reward Value is Acceptable
**Assumption:** The 9-to-1 ratio (buy 9, get 1 free) is perceived as fair and generous  
**Validation Needed:**
- Survey: "Do you think the current reward (10th coffee free) is fair?"
- Competitive analysis: What do other local cafés offer?
- Question: "Would a different reward structure be more appealing?"

### Assumption 9: Staff Consistency Issues
**Assumption:** Staff sometimes forget to stamp cards, causing customer frustration  
**Validation Needed:**
- Staff interview: "How often do you forget to stamp cards?"
- Customer question: "Has staff ever forgotten to stamp your card?"
- Observation: Watch checkout process for 20+ transactions

### Assumption 10: Community Atmosphere is Critical
**Assumption:** Regular customers (like Elena) prioritize personal connection over efficiency  
**Validation Needed:**
- Interview regulars: "What brings you back to Brews N' Bites?"
- Question: "Would you be okay with a more automated system if it meant less staff interaction?"
- Staff insight: Which customers value the personal touch most?

---

## Research Questions for Validation

Before proceeding to prototyping, answer these questions with real data:

**Customer Segmentation:**
1. What are the actual customer segments at Brews N' Bites? (Don't assume our 5 personas are correct)
2. What percentage of customers fall into each segment?
3. Which segments generate the most revenue?
4. Which segments are most engaged with current rewards?

**Current System Performance:**
5. What's the redemption rate of issued stamp cards?
6. How many active stamp card users are there?
7. What's the average time to complete a card?
8. How satisfied are customers with the current program? (Quantify)

**Pain Point Validation:**
9. What are customers' actual frustrations with the current system? (Ask open-ended)
10. Which pain points affect the most customers?
11. Which pain points cause the most frustration?
12. Are there pain points we haven't considered?

**Solution Preferences:**
13. Would customers actually use a digital loyalty system?
14. What features would they want most?
15. What concerns or objections do they have?
16. How much effort are they willing to put into a new system?

**Business Context:**
17. What are the business goals for the rewards program?
18. What metrics define success from the owner's perspective?
19. What's the budget for implementation?
20. What POS system does Brews N' Bites currently use?

---

*Document Version: 1.0*  
*Date: November 19, 2025*  
*Phase: Define (2 of 5)*  
*Status: ASSUMPTION-BASED - Requires validation before implementation*
