# Schema

This folder contains information on the database schema for the markup_scores database. It will contain 1 table with 5 columns. These columns will be:
```
unq_id      The unique id from the html file
id          The keyname value from the unq_id
id_date     The date from the unq_id
run_time    The time that this unq_id was scored
score       The score of the html content found with the unq_id
```

The second file in this folder contains a sql script to calculate the average of all scores inside the database.