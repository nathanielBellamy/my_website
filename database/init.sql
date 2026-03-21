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

-- Create Work Content Table
CREATE TABLE work_contents (
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
    INSERT INTO tags (name, activated_at) VALUES ('Go', NOW()), ('PostgreSQL', NOW()), ('Angular', NOW()), ('Software Engineering', NOW()), ('DevOps', NOW());

    -- Seed Blog Posts
    WITH author_id AS (SELECT id FROM authors WHERE name = 'Nate Schieber' LIMIT 1)
    INSERT INTO blog_posts (title, content, author_id, activated_at, ordering) VALUES
    (
      'Getting Started with Go',
      E'## Introduction to Go\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n\nPellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, commodo vitae, ornare sit amet, wisi.\n\n### Why Go?\n\nAenean fermentum risus id tortor. Integer imperdiet lectus quis justo. Nulla facilisi. Sed vel lectus. Donec odio urna, tempus molestie, porttitor ut, eperdiet a, purus. Sed non eros. Vivamus pharetra posuere sapien. Nam consectetuer. Sed aliquam, nunc eget euismod ullamcorper, lectus nunc ullamcorper orci, fermentum bibendum enim nibh eget ipsum.\n\nCras volutpat, lacus quis semper pharetra, nisi enim dignissim est, et sollicitudin quam ipsum vel mi. Sed commodo ultrices tortor. Integer vitae quam in nunc interdum laoreet. Duis vel nunc. Vivamus quis tellus vel quam varius adipiscing. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem.\n\n### Getting Your Environment Set Up\n\nFusce dui leo, imperdiet in, aliquam sit amet, feugiat eu, orci. Aenean vel massa quis mauris vehicula lacinia. Quisque tincidunt scelerisque libero. Maecenas libero. Etiam dictum tincidunt diam. Donec ipsum massa, ullamcorper in, auctor et, scelerisque sed, est. Suspendisse nisl. Sed vitae arcu. Aliquam id dolor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos.\n\nNam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante.\n\n### Writing Your First Program\n\nProin sapien ipsum, porta a, auctor quis, euismod ut, mi. Aenean viverra rhoncus pede. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut non enim eleifend felis pretium feugiat. Vivamus quis mi. Phasellus a est. Phasellus magna. In hac habitasse platea dictumst. Curabitur at lacus ac velit ornare lobortis. Cura bitur a felis in nunc fringilla tristique.',
      (SELECT id FROM author_id), NOW(), 1
    ),
    (
      'Advanced PostgreSQL Techniques',
      E'## Deep Dive into PostgreSQL\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent vestibulum molestie lacus. Aenean nonummy hendrerit mauris. Phasellus porta. Fusce suscipit varius mi. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla dui. Fusce feugiat malesuada odio. Morbi nunc odio, gravida at, cursus nec, luctus a, lorem.\n\nMaecenas tristique orci ac sem. Duis ultricies pharetra magna. Donec accumsan malesuada orci. Donec sit amet eros. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris fermentum dictum magna. Sed laoreet aliquam leo. Ut tellus dolor, dapibus eget, elementum vel, cursus eleifend, elit.\n\n### Window Functions\n\nAenean auctor wisi et urna. Aliquam erat volutpat. Duis ac turpis. Integer rutrum ante eu lacus. Vestibulum libero nisl, porta vel, scelerisque eget, malesuada at, neque. Vivamus eget nibh. Etiam cursus leo vel metus. Nulla facilisi. Aenean nec eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae.\n\nSuspendisse vel felis. Ut lorem lorem, interdum eu, tincidunt sit amet, laoreet vitae, arcu. Aenean faucibus pede eu ante. Praesent enim elit, rutrum at, molestie non, nonummy vel, nisl. Ut lectus eros, malesuada sit amet, fermentum eu, sodales cursus, magna. Donec eu purus. Quisque vehicula, urna sed ultricies auctor, pede lorem egestas dui, et convallis elit erat sed nulla.\n\n### Common Table Expressions\n\nDonec luctus. Curabitur et nunc. Aliquam dolor odio, commodo pretium, ultricies non, pharetra in, velit. Integer arcu est, nonummy in, fermentum faucibus, egestas vel, odio. Sed commodo posuere pede. Mauris ut est. Ut quis purus. Sed ac odio. Sed vehicula hendrerit sem. Duis non odio. Morbi ut dui. Sed accumsan risus eget odio.\n\nIn hac habitasse platea dictumst. Pellentesque non elit. Fusce sed justo eu urna porta tincidunt. Mauris felis odio, sollicitudin sed, volutpat a, ornare ac, erat. Morbi quis dolor. Donec pellentesque, erat ac sagittis semper, nunc dui lobortis purus, quis congue purus metus ultricies tellus. Proin et quam. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos.\n\n### Indexing Strategies\n\nNam dui ligula, fringilla a, euismod sodales, sollicitudin vel, wisi. Morbi auctor lorem non justo. Nam lacus libero, pretium at, lobortis vitae, ultricies et, tellus. Donec aliquet, tortor sed accumsan bibendum, erat ligula aliquet magna, vitae ornare odio metus a mi. Morbi ac orci et nisl hendrerit mollis. Suspendisse ut massa. Cras nec ante. Pellentesque a nulla. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aliquam tincidunt urna.',
      (SELECT id FROM author_id), NOW(), 2
    ),
    (
      'Building Modern UIs with Angular',
      E'## Angular in Practice\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus luctus urna sed urna ultricies ac tempor dui sagittis. In condimentum facilisis porta. Sed nec diam eu diam mattis viverra. Nulla fringilla, orci ac euismod semper, magna diam porttitor mauris, quis sollicitudin sapien justo in libero. Vestibulum mollis mauris enim.\n\nMorbi euismod magna ac lorem rutrum commodo. Laoreet non curabitur gravida arcu ac tortor. Ut consequat in nulla at sollicitudin. Fusce tristique felis at fermentum pharetra. Aliquam id orci ut lectus varius viverra. Nullam nunc ex, convallis sed finibus eget, sollicitudin eget ante. Phasellus eros quam, tincidunt vel vehicula ut, varius vitae metus.\n\n### Signal-Based State Management\n\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.\n\nNemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.\n\n### Component Architecture\n\nUt enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur. Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur.\n\nAt vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga.\n\n### Reactive Forms with Signals\n\nEt harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae.\n\nItaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.',
      (SELECT id FROM author_id), NOW(), 3
    ),
    (
      'DevOps Pipelines for Small Teams',
      E'## Practical DevOps\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur pretium tincidunt lacus. Nulla gravida orci a odio. Nullam varius, turpis et commodo pharetra, est eros bibendum elit, nec luctus magna felis sollicitudin mauris. Integer in mauris eu nibh euismod gravida. Duis ac tellus et risus vulputate vehicula. Donec lobortis risus a elit.\n\nEtiam vel neque nec dui dignissim bibendum. Vivamus id enim. Phasellus neque orci, porta a, aliquet quis, semper a, massa. Phasellus purus. Pellentesque tristique imperdiet tortor. Nam euismod tellus id erat. Curabitur blandit mollis lacus. Nulla sit amet magna. Pellentesque auctor neque nec urna.\n\n### CI/CD with GitHub Actions\n\nProin sapien ipsum, porta a, auctor quis, euismod ut, mi. Aenean viverra rhoncus pede. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut non enim eleifend felis pretium feugiat. Vivamus quis mi. Phasellus a est. Phasellus magna. In hac habitasse platea dictumst.\n\nCurabitur at lacus ac velit ornare lobortis. Curabitur a felis in nunc fringilla tristique. Morbi mattis ullamcorper velit. Phasellus gravida semper nisi. Nullam vel sem. Pellentesque libero tortor, tincidunt et, tincidunt eget, semper nec, quam. Sed hendrerit. Morbi ac felis.\n\n### Docker Compose in Production\n\nNunc interdum lacus sit amet orci. Vestibulum rutrum, mi nec elementum vehicula, eros quam gravida nisl, id fringilla neque ante vel mi. Morbi mollis tellus ac sapien. Phasellus volutpat, metus eget egestas mollis, lacus lacus blandit dui, id egestas quam mauris ut lacus. Fusce vel dui.\n\nSed in libero ut nibh placerat accumsan. Proin faucibus arcu quis ante. In consectetuer turpis ut velit. Nulla sit amet est. Praesent metus tellus, elementum eu, semper a, adipiscing nec, purus. Cras risus ipsum, faucibus ut, ullamcorper id, varius ac, leo. Suspendisse feugiat.\n\n### Monitoring and Alerting\n\nSuspendisse potenti. Sed tincidunt varius arcu. Mauris vitae nisi at sem facilisis semper. Proin id nisi id lacus porta fringilla. Morbi sagittis blandit sem. Cras et nunc. Pellentesque ut nulla. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non ipsum. Proin est. Ut tellus. Vestibulum sed nulla.',
      (SELECT id FROM author_id), NOW(), 4
    ),
    (
      'Deploying on NixOS: A Practical Guide',
      E'## NixOS for Server Deployments\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse varius enim in eros elementum tristique. Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut commodo diam libero vitae erat. Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet. Nunc ut sem vitae risus tristique posuere.\n\nClass aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur tempus urna at turpis condimentum lobortis. Ut commodo efficitur neque. Ut diam quam, nulla porttitor massa et, tempor lacinia erat. Vivamus in enim et enim congue condimentum. Maecenas eget condimentum nisi, vitae dignissim metus.\n\n### Declarative System Configuration\n\nPraesent commodo cursus magna, vel scelerisque nisl consectetur et. Donec sed odio dui. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla vitae elit libero, a pharetra augue. Sed posuere consectetur est at lobortis.\n\nDonec ullamcorper nulla non metus auctor fringilla. Vestibulum id ligula porta felis euismod semper. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus.\n\n### Reproducible Builds\n\nMorbi leo risus, porta ac consectetur ac, vestibulum at eros. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ut facilisis in, egestas eget quam.\n\nNullam id dolor id nibh ultricies vehicula ut id elit. Cras mattis consectetur purus sit amet fermentum. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor.\n\n### Managing Services with systemd\n\nAenean eu leo quam. Pellentesque ornare sem lacinia quam venenatis vestibulum. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Cras mattis consectetur purus sit amet fermentum. Nulla vitae elit libero, a pharetra augue.\n\nMaecenas faucibus mollis interdum. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Aenean lacinia bibendum nulla sed consectetur. Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit.',
      (SELECT id FROM author_id), NOW(), 5
    );

    -- Seed Blog Post Tags
    INSERT INTO blog_post_tags (blog_post_id, tag_id)
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Getting Started with Go' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'Go' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Getting Started with Go' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'Software Engineering' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Advanced PostgreSQL Techniques' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'PostgreSQL' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Building Modern UIs with Angular' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'Angular' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'DevOps Pipelines for Small Teams' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'DevOps' LIMIT 1)
    UNION ALL
    SELECT
        (SELECT id FROM blog_posts WHERE title = 'Deploying on NixOS: A Practical Guide' LIMIT 1),
        (SELECT id FROM tags WHERE name = 'DevOps' LIMIT 1);


    -- Seed Work Content
    INSERT INTO work_contents (title, content, activated_at, ordering) VALUES
    (
      'Welcome to My Website',
      E'## Hello and Welcome\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\n\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.\n\nNemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit.\n\nSed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur.',
      NOW(), 1
    ),
    (
      'What I Do',
      E'## Professional Focus\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper.\n\nAenean ultricies mi vitae est. Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, commodo vitae, ornare sit amet, wisi. Aenean fermentum risus id tortor. Integer imperdiet lectus quis justo.\n\nNulla facilisi. Sed vel lectus. Donec odio urna, tempus molestie, porttitor ut, eperdiet a, purus. Sed non eros. Vivamus pharetra posuere sapien. Nam consectetuer. Sed aliquam, nunc eget euismod ullamcorper, lectus nunc ullamcorper orci, fermentum bibendum enim nibh eget ipsum.\n\nCras volutpat, lacus quis semper pharetra, nisi enim dignissim est, et sollicitudin quam ipsum vel mi. Sed commodo ultrices tortor. Integer vitae quam in nunc interdum laoreet. Duis vel nunc.',
      NOW(), 2
    ),
    (
      'Latest Projects',
      E'## Current Work\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce dui leo, imperdiet in, aliquam sit amet, feugiat eu, orci. Aenean vel massa quis mauris vehicula lacinia. Quisque tincidunt scelerisque libero. Maecenas libero. Etiam dictum tincidunt diam.\n\nDonec ipsum massa, ullamcorper in, auctor et, scelerisque sed, est. Suspendisse nisl. Sed vitae arcu. Aliquam id dolor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. Nam eget dui. Etiam rhoncus.\n\nMaecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus.\n\nDonec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna.',
      NOW(), 3
    ),
    (
      'Tech Stack Overview',
      E'## Tools of the Trade\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse varius enim in eros elementum tristique. Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut commodo diam libero vitae erat. Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet.\n\nNunc ut sem vitae risus tristique posuere. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur tempus urna at turpis condimentum lobortis. Ut commodo efficitur neque.\n\nUt diam quam, nulla porttitor massa et, tempor lacinia erat. Vivamus in enim et enim congue condimentum. Maecenas eget condimentum nisi, vitae dignissim metus. Praesent commodo cursus magna, vel scelerisque nisl consectetur et.\n\nDonec sed odio dui. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla vitae elit libero, a pharetra augue.',
      NOW(), 4
    ),
    (
      'Get In Touch',
      E'## Contact Information\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ullamcorper nulla non metus auctor fringilla. Vestibulum id ligula porta felis euismod semper. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh.\n\nUt fermentum massa justo sit amet risus. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit.\n\nCras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ut facilisis in, egestas eget quam. Nullam id dolor id nibh ultricies vehicula ut id elit. Cras mattis consectetur purus sit amet fermentum.\n\nMorbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor.',
      NOW(), 5
    );

    -- Seed GrooveJr Content
    INSERT INTO groove_jr_contents (title, content, activated_at, ordering) VALUES
    (
      'What is GrooveJr?',
      E'## Introducing GrooveJr\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\n\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis.\n\nQuasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.\n\nNeque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.',
      NOW(), 1
    ),
    (
      'GrooveJr Architecture',
      E'## System Design\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper.\n\nAenean ultricies mi vitae est. Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, commodo vitae, ornare sit amet, wisi. Aenean fermentum risus id tortor.\n\nInteger imperdiet lectus quis justo. Nulla facilisi. Sed vel lectus. Donec odio urna, tempus molestie, porttitor ut, eperdiet a, purus. Sed non eros. Vivamus pharetra posuere sapien. Nam consectetuer.\n\nSed aliquam, nunc eget euismod ullamcorper, lectus nunc ullamcorper orci, fermentum bibendum enim nibh eget ipsum. Cras volutpat, lacus quis semper pharetra, nisi enim dignissim est, et sollicitudin quam ipsum vel mi.',
      NOW(), 2
    ),
    (
      'Getting Started with GrooveJr',
      E'## Quick Start Guide\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce dui leo, imperdiet in, aliquam sit amet, feugiat eu, orci. Aenean vel massa quis mauris vehicula lacinia. Quisque tincidunt scelerisque libero. Maecenas libero. Etiam dictum tincidunt diam.\n\nDonec ipsum massa, ullamcorper in, auctor et, scelerisque sed, est. Suspendisse nisl. Sed vitae arcu. Aliquam id dolor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos.\n\nNam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.\n\nMaecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh.',
      NOW(), 3
    ),
    (
      'GrooveJr Audio Engine',
      E'## The Sound Core\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse varius enim in eros elementum tristique. Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut commodo diam libero vitae erat. Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet.\n\nNunc ut sem vitae risus tristique posuere. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur tempus urna at turpis condimentum lobortis. Ut commodo efficitur neque.\n\nUt diam quam, nulla porttitor massa et, tempor lacinia erat. Vivamus in enim et enim congue condimentum. Maecenas eget condimentum nisi, vitae dignissim metus. Praesent commodo cursus magna, vel scelerisque nisl consectetur et.\n\nDonec sed odio dui. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla vitae elit libero, a pharetra augue. Sed posuere consectetur est at lobortis.',
      NOW(), 4
    ),
    (
      'GrooveJr Roadmap',
      E'## Future Plans\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ullamcorper nulla non metus auctor fringilla. Vestibulum id ligula porta felis euismod semper. Praesent commodo cursus magna, vel scelerisque nisl consectetur et.\n\nFusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.\n\nDuis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ut facilisis in, egestas eget quam.\n\nNullam id dolor id nibh ultricies vehicula ut id elit. Cras mattis consectetur purus sit amet fermentum. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et.',
      NOW(), 5
    );

    -- Seed About Content
    INSERT INTO about_contents (title, content, activated_at, ordering) VALUES
    (
      'About Me',
      E'## Who I Am\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.\n\nExcepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.\n\nNemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit.\n\nSed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam.',
      NOW(), 1
    ),
    (
      'My Background',
      E'## Education and Experience\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper.\n\nAenean ultricies mi vitae est. Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, commodo vitae, ornare sit amet, wisi. Aenean fermentum risus id tortor. Integer imperdiet lectus quis justo.\n\nNulla facilisi. Sed vel lectus. Donec odio urna, tempus molestie, porttitor ut, eperdiet a, purus. Sed non eros. Vivamus pharetra posuere sapien. Nam consectetuer. Sed aliquam, nunc eget euismod ullamcorper, lectus nunc ullamcorper orci.\n\nFermentum bibendum enim nibh eget ipsum. Cras volutpat, lacus quis semper pharetra, nisi enim dignissim est, et sollicitudin quam ipsum vel mi. Sed commodo ultrices tortor.',
      NOW(), 2
    ),
    (
      'Philosophy and Approach',
      E'## How I Think About Software\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce dui leo, imperdiet in, aliquam sit amet, feugiat eu, orci. Aenean vel massa quis mauris vehicula lacinia. Quisque tincidunt scelerisque libero. Maecenas libero. Etiam dictum tincidunt diam.\n\nDonec ipsum massa, ullamcorper in, auctor et, scelerisque sed, est. Suspendisse nisl. Sed vitae arcu. Aliquam id dolor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos.\n\nNam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem.\n\nMaecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna.',
      NOW(), 3
    ),
    (
      'Skills and Technologies',
      E'## Technical Expertise\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse varius enim in eros elementum tristique. Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut commodo diam libero vitae erat. Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet.\n\nNunc ut sem vitae risus tristique posuere. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur tempus urna at turpis condimentum lobortis. Ut commodo efficitur neque.\n\nUt diam quam, nulla porttitor massa et, tempor lacinia erat. Vivamus in enim et enim congue condimentum. Maecenas eget condimentum nisi, vitae dignissim metus. Praesent commodo cursus magna, vel scelerisque nisl consectetur et.\n\nDonec sed odio dui. Donec id elit non mi porta gravida at eget metus. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla vitae elit libero, a pharetra augue.',
      NOW(), 4
    ),
    (
      'Outside of Work',
      E'## Life Beyond Code\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ullamcorper nulla non metus auctor fringilla. Vestibulum id ligula porta felis euismod semper. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh.\n\nUt fermentum massa justo sit amet risus. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit.\n\nCras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ut facilisis in, egestas eget quam. Nullam id dolor id nibh ultricies vehicula ut id elit. Cras mattis consectetur purus sit amet fermentum.\n\nMorbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor. Aenean eu leo quam. Pellentesque ornare sem lacinia quam venenatis vestibulum.',
      NOW(), 5
    );

    CREATE TABLE seed_data_inserted (
      id serial PRIMARY KEY,
      inserted_at timestamptz DEFAULT now()
    );
  END IF;
END $$;