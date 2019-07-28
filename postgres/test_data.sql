INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('First Blog Post', 'This is blog about something.', 'Martin', '2018-08-24 14:00:00', null, null);
INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('Second Blog Post', 'This is blog about something else...', 'Martin', '2019-02-24 13:00:00', null, null);
INSERT INTO posts (title, text, author, posted_on, next_post_id, previous_post_id) VALUES ('3rd Blog Post', 'Another dummy content', 'Martin', '2019-05-30 19:00:00', null, null);

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

INSERT INTO projects (name, thumbnail_url, url, description) VALUES ('IoT Cloud', 'https://via.placeholder.com/150', 'https://github.com/MartinHeinz/IoT-Cloud', 'Cloud framework for IoT (Internet of Things), which focuses on security and privacy of its users, their devices and data');
INSERT INTO projects (name, thumbnail_url, url, description) VALUES ('Blog & Personal Website', 'https://via.placeholder.com/150', 'https://github.com/MartinHeinz/blog-backend', 'This website. Goal of this project was to learn Go and Vue.js and as a byproduct I created personal website and blog.');

INSERT INTO tags (name, post_id, project_id) VALUES ('Python', 1, null);
INSERT INTO tags (name, post_id, project_id) VALUES ('Python', 2, null);
INSERT INTO tags (name, post_id, project_id) VALUES ('Crypto', 1, null);
INSERT INTO tags (name, post_id, project_id) VALUES ('Golang', 1, null);
INSERT INTO tags (name, post_id, project_id) VALUES ('Vue', 3, null);

INSERT INTO tags (name, post_id, project_id) VALUES ('Python', null, 1);
INSERT INTO tags (name, post_id, project_id) VALUES ('Cryptography', null, 1);
INSERT INTO tags (name, post_id, project_id) VALUES ('Privacy', null, 1);
INSERT INTO tags (name, post_id, project_id) VALUES ('IoT', null, 1);
INSERT INTO tags (name, post_id, project_id) VALUES ('Vue', null, 2);
INSERT INTO tags (name, post_id, project_id) VALUES ('Golang', null, 2);
INSERT INTO tags (name, post_id, project_id) VALUES ('Docker', null, 2);

INSERT INTO books (title, cover_url, url, alt) VALUES ('The Go Programming Language', 'https://www.gopl.io/cover.png', 'https://www.gopl.io/', 'The Go Programming Language');
INSERT INTO books (title, cover_url, url, alt) VALUES ('Clean Code', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1436202607i/3735293._SX318_.jpg', 'https://www.goodreads.com/book/show/3735293-clean-code', 'Clean Code: A Handbook of Agile Software Craftsmanship');
INSERT INTO books (title, cover_url, url, alt) VALUES ('Software Craftsmanship', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1370897661i/18054154.jpg', 'https://www.goodreads.com/book/show/18054154-software-craftsmanship', 'The Software Craftsman: Professionalism, Pragmatism, Pride');
INSERT INTO books (title, cover_url, url, alt) VALUES ('Extreme Programming Explained', 'https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1386925310i/67833.jpg', 'https://www.goodreads.com/book/show/67833.Extreme_Programming_Explained', 'Extreme Programming Explained: Embrace Change (The XP Series)');
