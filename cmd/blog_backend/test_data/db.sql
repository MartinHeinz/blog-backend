INSERT INTO posts (title, text, posted_on) VALUES ('First Blog Post', 'This is blog about something.', '2018-08-24 14:00:00 +02:00');

INSERT INTO sections (name, post_id) VALUES ('Title', 1);
INSERT INTO sections (name, post_id) VALUES ('Subsection', 1);

INSERT INTO tags (name, post_id) VALUES ('Python', 1);
INSERT INTO tags (name, post_id) VALUES ('Crypto', 1);
