-- ------------------------------------------
-- Change to the markup_scores db.
-- ------------------------------------------
USE markup_scores;

-- ------------------------------------------
-- Get the average of all scores.
-- ------------------------------------------

SELECT id, AVG(score) FROM scores GROUP BY id;
