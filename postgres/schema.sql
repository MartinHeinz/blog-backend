SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;


CREATE TABLE public.books (
                              id integer NOT NULL,
                              created_at timestamp with time zone,
                              updated_at timestamp with time zone,
                              deleted_at timestamp with time zone,
                              title text,
                              cover_url text,
                              url text,
                              alt text
);


ALTER TABLE public.books OWNER TO postgres;


CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO postgres;


ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;



CREATE TABLE public.posts (
                              id integer NOT NULL,
                              created_at timestamp with time zone,
                              updated_at timestamp with time zone,
                              deleted_at timestamp with time zone,
                              title text,
                              text text,
                              posted_on timestamp with time zone,
                              author text,
                              post_id integer,
                              next_post_id integer,
                              previous_post_id integer
);


ALTER TABLE public.posts OWNER TO postgres;


CREATE SEQUENCE public.posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO postgres;


ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;



CREATE TABLE public.sections (
                                 post_id integer,
                                 id integer NOT NULL,
                                 created_at timestamp with time zone,
                                 updated_at timestamp with time zone,
                                 deleted_at timestamp with time zone,
                                 name text
);


ALTER TABLE public.sections OWNER TO postgres;


CREATE SEQUENCE public.sections_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sections_id_seq OWNER TO postgres;


ALTER SEQUENCE public.sections_id_seq OWNED BY public.sections.id;



CREATE TABLE public.tags (
                             post_id integer,
                             id integer NOT NULL,
                             created_at timestamp with time zone,
                             updated_at timestamp with time zone,
                             deleted_at timestamp with time zone,
                             name text
);


ALTER TABLE public.tags OWNER TO postgres;


CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO postgres;


ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;



ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);



ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);



ALTER TABLE ONLY public.sections ALTER COLUMN id SET DEFAULT nextval('public.sections_id_seq'::regclass);



ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);



ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.sections
    ADD CONSTRAINT sections_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);



CREATE INDEX idx_books_deleted_at ON public.books USING btree (deleted_at);



CREATE INDEX idx_posts_deleted_at ON public.posts USING btree (deleted_at);



CREATE INDEX idx_sections_deleted_at ON public.sections USING btree (deleted_at);



CREATE INDEX idx_tags_deleted_at ON public.tags USING btree (deleted_at);



ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_next_post_id_posts_id_foreign FOREIGN KEY (next_post_id) REFERENCES public.posts(id) ON UPDATE CASCADE ON DELETE CASCADE;



ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_previous_post_id_posts_id_foreign FOREIGN KEY (previous_post_id) REFERENCES public.posts(id) ON UPDATE CASCADE ON DELETE CASCADE;



ALTER TABLE ONLY public.sections
    ADD CONSTRAINT sections_post_id_posts_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts(id) ON UPDATE CASCADE ON DELETE CASCADE;



ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_post_id_posts_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts(id) ON UPDATE CASCADE ON DELETE CASCADE;

-- Retrieval steps:
--
-- Remove old volumes
-- docker volume prune
-- Add automigrate to Gin start-up
-- Build container (make container)
--
-- docker-compose up
-- docker exec -it <blog_db_id> /bin/bash
-- pg_dump -s blog -Upostgres | sed -e '/^--/d'
--
-- Don't forget to remove automigrate
