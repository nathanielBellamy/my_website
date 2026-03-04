-- Create users safely (idempotent)
DO $$
BEGIN
  IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'marketing') THEN
    CREATE USER marketing WITH PASSWORD 'marketing';
  END IF;
END
$$;

-- Grant read-only privileges to marketing user
GRANT CONNECT ON DATABASE mw_db TO marketing;


\c mw_db;

-- Create Authors Table
CREATE TABLE authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL
);

-- Create Tags Table
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- Create Blog Posts Table
CREATE TABLE blog_posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    ordering INTEGER DEFAULT 0,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id UUID REFERENCES authors(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create a join table for the many-to-many relationship between blog_posts and tags
CREATE TABLE blog_post_tags (
    blog_post_id UUID REFERENCES blog_posts(id) ON DELETE CASCADE,
    tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (blog_post_id, tag_id)
);

-- Create Home Content Table
CREATE TABLE home_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    ordering INTEGER DEFAULT 0,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

-- Create GrooveJr Content Table
CREATE TABLE groove_jr_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    ordering INTEGER DEFAULT 0,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

-- Create About Content Table
CREATE TABLE about_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    activated_at TIMESTAMPTZ,
    deactivated_at TIMESTAMPTZ,
    ordering INTEGER DEFAULT 0,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

-- Create Images Table
CREATE TABLE images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    filename VARCHAR(255) NOT NULL UNIQUE,
    original_name VARCHAR(255) NOT NULL,
    alt_text VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Seed Data
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'seed_data_inserted') THEN
    -- Seed Authors
    INSERT INTO authors (name, activated_at) VALUES ('Nate Schieber', NOW());

    -- Seed Tags
    INSERT INTO tags (name, activated_at) VALUES ('Go', NOW()), ('PostgreSQL', NOW()), ('Angular', NOW()), ('Software Engineering', NOW());

    -- Seed Blog Posts
    WITH author_id AS (SELECT id FROM authors WHERE name = 'Nate Schieber' LIMIT 1)
    INSERT INTO blog_posts (title, content, author_id, activated_at) VALUES
    ('Getting Started with Go', 'A beginner''s guide to the Go programming language.', (SELECT id FROM author_id), NOW()),
    ('Advanced PostgreSQL', 'Exploring advanced features of PostgreSQL.', (SELECT id FROM author_id), NOW());

    -- Seed Blog Post Tags
    INSERT INTO blog_post_tags (blog_post_id, tag_id)
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Getting Started with Go' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'Go' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Advanced PostgreSQL' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'PostgreSQL' LIMIT 1);


    -- Seed Home Content
    INSERT INTO home_contents (title, content, activated_at) VALUES
    ('Welcome to my Website!', 'This is the home page.', NOW());

    -- Seed GrooveJr Content
    INSERT INTO groove_jr_contents (title, content, activated_at) VALUES
    ('GrooveJr', 'All about GrooveJr.', NOW());

    -- Seed About Content
    INSERT INTO about_contents (title, content, activated_at) VALUES
    ('About Me', 'My name is Nate and I''m a software engineer.', NOW());

    CREATE TABLE seed_data_inserted (
      id serial PRIMARY KEY,
      inserted_at timestamptz DEFAULT now()
    );
  END IF;
END $$;