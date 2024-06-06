# my golang_recipe

"do one thing well". 

You won't find monolithic code in here. Most of the folders have one or two simple concepts in code. This is my golang toolbox. Use grep to scan through files and folders for that screwdriver that might help...build complex, not complicated. This repository is updated on a regular basis.

## TODO

. Blockchain project - started 04june2024
. June 2024 Project: building a mailer software for client. Send emails to office365 inbox only. This video served as reference...client claims he is not going to use it for spam https://www.youtube.com/@spamerofficial-dj9of

. The lean statup summarized https://youtu.be/RSaIOCHbuYw?si=Kn83LMSg1dWiQUp6
    learn build measure

### notes
rand.Seed deprecated? Do...
Source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
number := r.Intn(10)

...or, if you just need a random number
rand.Intn(10)

#### The Lean Startup, summarized
"The Lean Startup" by Eric Ries is a book that introduces a new approach to business development that aims to shorten product development cycles and rapidly discover if a proposed business model is viable. This methodology is designed for startups of all sizes to efficiently allocate resources and build sustainable businesses. Here’s a summary of the key concepts:

### Key Concepts

1. **Validated Learning**:
   - Startups exist not just to make stuff, make money, or serve customers. They exist to learn how to build a sustainable business.
   - This learning can be validated scientifically by running experiments that test the startup’s vision.

2. **Build-Measure-Learn Feedback Loop**:
   - This cycle is at the core of the Lean Startup model.
   - **Build**: Start with the Minimum Viable Product (MVP), a version of the product with the least features needed to start the learning process.
   - **Measure**: Use validated learning to measure how the product is performing. This involves rigorous metrics to assess the results of the experiment.
   - **Learn**: Analyze the data to learn what aspects of the product are working and what needs to change. This learning informs the next iteration of the cycle.

3. **MVP (Minimum Viable Product)**:
   - An MVP is the simplest version of a product that allows a team to quickly build and test its basic assumptions.
   - It is used to gather validated learning with the least effort, quickly iterating based on feedback.

4. **Innovation Accounting**:
   - Traditional accounting methods are not sufficient for startups.
   - Innovation accounting provides a new way of measuring progress, setting up milestones, and prioritizing work.
   - Metrics should focus on actionable, accessible, and auditable data.

5. **Pivot or Persevere**:
   - Startups must frequently decide whether to pivot (make a fundamental change to the product) or persevere (keep improving the current course).
   - Pivoting means making a significant change to the product strategy to test a new fundamental hypothesis about the product, business model, and engine of growth.

6. **Engines of Growth**:
   - Startups can achieve sustainable growth through three primary engines: the **sticky engine** (customer retention), the **viral engine** (word-of-mouth growth), and the **paid engine** (using revenue to acquire customers).

7. **Lean Thinking**:
   - Derived from lean manufacturing principles, Lean Thinking in startups focuses on eliminating waste and optimizing processes to create value efficiently.

### Implementation Strategies

- **Customer Development**: Engage with customers early and often to validate assumptions and learn what they truly need.
- **A/B Testing**: Conduct experiments comparing two versions of a product to determine which performs better.
- **Continuous Deployment**: Continuously integrate and deploy new code into the production environment to accelerate feedback and learning.
- **Actionable Metrics vs. Vanity Metrics**: Focus on metrics that can guide decision-making rather than those that just make the startup look good.

### Conclusion

"The Lean Startup" emphasizes a scientific approach to building and managing startups, focusing on rapid iteration, validated learning, and a flexible mindset to adapt and grow. This methodology aims to reduce product development cycles, measure progress accurately, and understand customer needs better, ultimately leading to building successful and sustainable businesses.
