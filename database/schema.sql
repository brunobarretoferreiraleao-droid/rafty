-- =========================================================
-- EXTENSIONS
-- =========================================================
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- =========================================================
-- SHARED FUNCTIONS
-- =========================================================

-- Atualização automática do campo updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- =========================================================
-- ENUM TYPES
-- =========================================================

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'post_visibility') THEN
        CREATE TYPE post_visibility AS ENUM ('public', 'friends', 'custom', 'private');
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'friendship_status') THEN
        CREATE TYPE friendship_status AS ENUM ('pending', 'accepted', 'rejected');
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'attachment_type') THEN
        CREATE TYPE attachment_type AS ENUM ('image', 'audio', 'video', 'gif', 'file');
    END IF;
END
$$;

-- =========================================================
-- USERS
-- =========================================================

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username CITEXT UNIQUE NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,

    display_name VARCHAR(100),
    bio TEXT,
    avatar_url TEXT,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (length(trim(username::text)) >= 3),
    CHECK (length(trim(email::text)) > 0),
    CHECK (display_name IS NULL OR length(trim(display_name)) > 0),
    CHECK (bio IS NULL OR length(bio) <= 500),
    CHECK (avatar_url IS NULL OR length(trim(avatar_url)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_users_username_trgm
ON users USING gin (username gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_users_display_name_trgm
ON users USING gin (display_name gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_users_email_trgm
ON users USING gin (email gin_trgm_ops);


DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================================
-- POSTS
-- =========================================================

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    title VARCHAR(255),
    content TEXT,
    visibility post_visibility NOT NULL DEFAULT 'public',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (title IS NULL OR length(trim(title)) > 0),
    CHECK (content IS NULL OR length(trim(content)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_posts_author ON posts (author_id);
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_posts_visibility ON posts (visibility);

CREATE INDEX IF NOT EXISTS idx_posts_author_created_at ON posts (author_id, created_at DESC);

DROP TRIGGER IF EXISTS update_posts_updated_at ON posts;
CREATE TRIGGER update_posts_updated_at
BEFORE UPDATE ON posts
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================================
-- POST ATTACHMENTS
-- =========================================================

CREATE TABLE IF NOT EXISTS post_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,

    file_url TEXT NOT NULL,
    file_name TEXT,
    mime_type TEXT,
    file_size BIGINT,
    attachment_type attachment_type NOT NULL,
    display_order INT NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (length(trim(file_url)) > 0),
    CHECK (file_name IS NULL OR length(trim(file_name)) > 0),
    CHECK (mime_type IS NULL OR length(trim(mime_type)) > 0),
    CHECK (file_size IS NULL OR file_size >= 0),
    CHECK (display_order >= 0)
);

CREATE INDEX IF NOT EXISTS idx_post_attachments_post_id
ON post_attachments(post_id);

-- =========================================================
-- FOLLOWS
-- =========================================================

CREATE TABLE IF NOT EXISTS follows (
    follower_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    following_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (follower_id, following_id),
    CHECK (follower_id <> following_id)
);

CREATE INDEX IF NOT EXISTS idx_follows_follower ON follows (follower_id);
CREATE INDEX IF NOT EXISTS idx_follows_following ON follows (following_id);

-- =========================================================
-- FRIENDS
-- =========================================================

CREATE TABLE IF NOT EXISTS friends (
    user_id1 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_id2 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id1, user_id2),
    CHECK (user_id1 < user_id2)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_friends_unique_pair ON friends (
    LEAST(user_id1, user_id2),
    GREATEST(user_id1, user_id2)
);

-- =========================================================
-- FRIEND REQUESTS
-- =========================================================

CREATE TABLE IF NOT EXISTS friend_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    requester_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status friendship_status NOT NULL DEFAULT 'pending',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (requester_id <> receiver_id)
);

CREATE INDEX IF NOT EXISTS idx_friend_requests_requester ON friend_requests (requester_id);
CREATE INDEX IF NOT EXISTS idx_friend_requests_receiver ON friend_requests (receiver_id);
CREATE INDEX IF NOT EXISTS idx_friend_requests_status ON friend_requests (status);

-- Garante unicidade por par (A-B == B-A)
CREATE UNIQUE INDEX IF NOT EXISTS idx_friend_requests_unique_pair ON friend_requests (
    LEAST(requester_id, receiver_id),
    GREATEST(requester_id, receiver_id)
);

DROP TRIGGER IF EXISTS update_friend_requests_updated_at ON friend_requests;
CREATE TRIGGER update_friend_requests_updated_at
BEFORE UPDATE ON friend_requests
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================================
-- FRIEND REQUEST VALIDATION FUNCTION
-- =========================================================

CREATE OR REPLACE FUNCTION prevent_duplicate_friend_requests()
RETURNS TRIGGER AS $$
BEGIN
    -- Já são amigos?
    IF EXISTS (
        SELECT 1
        FROM friends
        WHERE user_id1 = LEAST(NEW.requester_id, NEW.receiver_id)
          AND user_id2 = GREATEST(NEW.requester_id, NEW.receiver_id)
    ) THEN
        RAISE EXCEPTION 'Amizade já existe';
    END IF;

    -- Já existe solicitação pendente entre o par?
    IF EXISTS (
        SELECT 1
        FROM friend_requests
        WHERE LEAST(requester_id, receiver_id) = LEAST(NEW.requester_id, NEW.receiver_id)
          AND GREATEST(requester_id, receiver_id) = GREATEST(NEW.requester_id, NEW.receiver_id)
          AND status = 'pending'
    ) THEN
        RAISE EXCEPTION 'Solicitação já existe';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_prevent_duplicate_friend_requests ON friend_requests;
CREATE TRIGGER trg_prevent_duplicate_friend_requests
BEFORE INSERT ON friend_requests
FOR EACH ROW
EXECUTE FUNCTION prevent_duplicate_friend_requests();

-- =========================================================
-- CONVERSATIONS
-- =========================================================

-- =========================================================
-- CONVERSATIONS
-- =========================================================

CREATE TABLE IF NOT EXISTS conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user1_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user2_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_message_at TIMESTAMPTZ,

    CHECK (user1_id <> user2_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_conversations_unique_pair ON conversations (
    LEAST(user1_id, user2_id),
    GREATEST(user1_id, user2_id)
);

CREATE INDEX IF NOT EXISTS idx_conversations_user1 ON conversations(user1_id);
CREATE INDEX IF NOT EXISTS idx_conversations_user2 ON conversations(user2_id);
CREATE INDEX IF NOT EXISTS idx_conversations_ordered ON conversations(
    LEAST(user1_id, user2_id),
    GREATEST(user1_id, user2_id),
    created_at DESC
);
CREATE INDEX IF NOT EXISTS idx_conversations_last_message_at ON conversations(last_message_at DESC);
-- =========================================================
-- MESSAGES
-- =========================================================

CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    content TEXT,
    is_edited BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (content IS NULL OR length(trim(content)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_messages_conversation_created_at
ON messages(conversation_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_messages_sender_id
ON messages(sender_id);

DROP TRIGGER IF EXISTS update_messages_updated_at ON messages;
CREATE TRIGGER update_messages_updated_at
BEFORE UPDATE ON messages
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

--Validar messages.sender_id pertence à conversation_id
CREATE OR REPLACE FUNCTION validate_message_sender()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM conversations
        WHERE id = NEW.conversation_id
          AND (user1_id = NEW.sender_id OR user2_id = NEW.sender_id)
    ) THEN
        RAISE EXCEPTION 'Sender must be part of the conversation';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_validate_message_sender ON messages;
CREATE TRIGGER trg_validate_message_sender
BEFORE INSERT ON messages
FOR EACH ROW
EXECUTE FUNCTION validate_message_sender();


-- =========================================================
-- MESSAGE ATTACHMENTS
-- =========================================================

CREATE TABLE IF NOT EXISTS message_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID NOT NULL REFERENCES messages(id) ON DELETE CASCADE,

    file_url TEXT NOT NULL,
    file_name TEXT,
    mime_type TEXT,
    file_size BIGINT,
    attachment_type attachment_type NOT NULL,
    display_order INT NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (length(trim(file_url)) > 0),
    CHECK (file_name IS NULL OR length(trim(file_name)) > 0),
    CHECK (mime_type IS NULL OR length(trim(mime_type)) > 0),
    CHECK (file_size IS NULL OR file_size >= 0),
    CHECK (display_order >= 0)
);

CREATE INDEX IF NOT EXISTS idx_message_attachments_message_id
ON message_attachments(message_id);

-- Update last_message_at na tabela conversations quando uma nova mensagem é inserida
CREATE OR REPLACE FUNCTION update_conversation_last_message_at()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE conversations
    SET last_message_at = NEW.created_at
    WHERE id = NEW.conversation_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_update_conversation_last_message_at ON messages;
CREATE TRIGGER trg_update_conversation_last_message_at
AFTER INSERT ON messages
FOR EACH ROW
EXECUTE FUNCTION update_conversation_last_message_at();

-- mark message as edited
CREATE OR REPLACE FUNCTION mark_message_as_edited()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.content IS DISTINCT FROM OLD.content THEN
        NEW.is_edited = TRUE;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_mark_message_as_edited ON messages;
CREATE TRIGGER trg_mark_message_as_edited
BEFORE UPDATE ON messages
FOR EACH ROW
EXECUTE FUNCTION mark_message_as_edited();

-- Reactions enum
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reaction_type') THEN
        CREATE TYPE reaction_type AS ENUM ('like', 'deslike', 'love', 'haha', 'wow', 'sad', 'angry');
    END IF;
END
$$;

-- =========================================================
-- REACTIONS
-- =========================================================

CREATE TABLE IF NOT EXISTS post_reactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reaction reaction_type NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (post_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_post_reactions_post_id ON post_reactions (post_id);
CREATE INDEX IF NOT EXISTS idx_post_reactions_user_id ON post_reactions (user_id);
CREATE INDEX IF NOT EXISTS idx_post_reactions_reaction ON post_reactions (reaction);

DROP TRIGGER IF EXISTS update_post_reactions_updated_at ON post_reactions;
CREATE TRIGGER update_post_reactions_updated_at
BEFORE UPDATE ON post_reactions
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- =========================================================
-- Comments
-- =========================================================

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    parent_comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,

    content TEXT NOT NULL,
    is_edited BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CHECK (content IS NULL OR length(trim(content)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments (post_id);
CREATE INDEX IF NOT EXISTS idx_comments_author_id ON comments (author_id);
CREATE INDEX IF NOT EXISTS idx_comments_parents_id ON comments (parent_comment_id);
CREATE INDEX IF NOT EXISTS idx_comments_created_at ON comments (created_at DESC);

DROP TRIGGER IF EXISTS update_comments_updated_at ON comments;
CREATE TRIGGER update_comments_updated_at
BEFORE UPDATE ON comments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Notifications enum
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'notification_type') THEN
        CREATE TYPE notification_type AS ENUM (
            'friend_request',
            'friend_request_accepted',
            'follow',
            'post_from_following',
            'post_from_friend',
            'post_reaction',
            'comment',
            'comment_reply',
            'message'
        );
    END IF;
END
$$;


-- =========================================================
-- NOTIFICATIONS
-- =========================================================

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    actor_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    type notification_type NOT NULL,
    data JSONB,    

    post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
    comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    message_id UUID REFERENCES messages(id) ON DELETE CASCADE,

    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    read_at TIMESTAMPTZ
);

    CREATE INDEX if NOT EXISTS idx_notifications_user_id (user_id),
    CREATE INDEX if NOT EXISTS idx_notifications_type (type),
    CREATE INDEX if NOT EXISTS idx_notifications_created_at (created_at DESC);  

-- =========================================================
-- post_visibility_rules
-- =========================================================

CREATE TABLE IF NOT EXISTS post_visibility_rules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    can_view BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (post_id, user_id)
);
create index if not exists idx_post_visibility_rules_post_id on post_visibility_rules (post_id);
create index if not exists idx_post_visibility_rules_user_id on post_visibility_rules (user_id);

--- =========================================================

-- FULL-TEXT SEARCH
ALTER TABLE users
ADD COLUMN IF NOT EXISTS search_vector tsvector
GENERATED ALWAYS AS (
    setweight(to_tsvector('simple', coalesce(username::text, '')), 'A') ||
    setweight(to_tsvector('simple', coalesce(display_name, '')), 'B')
) STORED;

CREATE INDEX IF NOT EXISTS idx_users_search_vector
ON users USING GIN (search_vector);