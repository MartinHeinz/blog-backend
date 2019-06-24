INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('First Blog Post', 'This is blog about something.', 'Martin', '2018-08-24 14:00:00 +02:00', null, null);
INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('Second Blog Post', 'This is blog about something else...', 'Martin', '2019-02-24 13:00:00 +02:00', null, null);
INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('3rd Blog Post', 'Another dummy content', 'Martin', '2019-05-30 19:00:00 +02:00', null, null);

UPDATE posts SET next_post_id = 2, previous_post_id = null WHERE id = 1;
UPDATE posts SET next_post_id = 3, previous_post_id = 1  WHERE id = 2;
UPDATE posts SET next_post_id = null, previous_post_id = 2 WHERE id = 3;

INSERT INTO sections (name, post_id) VALUES ('Title', 1);
INSERT INTO sections (name, post_id) VALUES ('Subsection', 1);
INSERT INTO sections (name, post_id) VALUES ('Intro', 2);
INSERT INTO sections (name, post_id) VALUES ('Subsection', 2);
INSERT INTO sections (name, post_id) VALUES ('Subsection 2', 2);
INSERT INTO sections (name, post_id) VALUES ('Subsection 3', 2);
INSERT INTO sections (name, post_id) VALUES ('Intro', 3);
INSERT INTO sections (name, post_id) VALUES ('First Section', 3);

INSERT INTO tags (name, post_id) VALUES ('Python', 1);
INSERT INTO tags (name, post_id) VALUES ('Python', 2);
INSERT INTO tags (name, post_id) VALUES ('Crypto', 1);
INSERT INTO tags (name, post_id) VALUES ('Golang', 1);
INSERT INTO tags (name, post_id) VALUES ('Vue', 3);
