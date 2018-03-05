# Statistics for Hackers

Base upon the examples of the talk ["Statistics for Hackers" by Jake Vanderplas](https://www.youtube.com/watch?v=Iq9DzN6mvYA) at PyCon 2016.

## Direct simulation

Computing the Sampling Distribution is Hard.
Simulating the Sampling Distribution is Easy.

## Shuffling

Idea: Simulate the distribution by shuffling the labels repeatedely and computing the desired statistics.
Motivation: If the labels really don't matter, then switching them shouldn't change the result.

## Bootstrapping

Idea: Simulate the distribution by drawing samples with replacement.
Motivation: The data estimates its own distribution - we draw random samples from this distribution.

### Notes

- Bootstrap resampling is well studied and rests on solid theoretical grounds.
- Bootstrap often doesn't work well for rank-based statistics (e.g. maximum value).
- Works poorly with very few samples (N > 20 is a good rule of thumb).
- As always, be careful about selection biases & non-independent data!


## Cross validatoin
