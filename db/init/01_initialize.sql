-- users テーブルを作成
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT,
    profile TEXT,
    profile_image_url TEXT,
    gender TEXT,
    mbti TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- レコードの更新時に updatedAt を自動更新するトリガー関数を作成
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updatedAt = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- users テーブルの updatedAt フィールドを自動更新するトリガーを作成
CREATE TRIGGER update_users_updated_at BEFORE
UPDATE ON users FOR EACH ROW
EXECUTE PROCEDURE update_timestamp ();

-- posts テーブルの作成
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    category VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- レコードの更新時に updatedAt を自動更新するトリガー関数
CREATE OR REPLACE FUNCTION update_post_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updatedAt = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- posts テーブルの updatedAt フィールドを自動更新するトリガー
CREATE TRIGGER update_posts_updated_at BEFORE
UPDATE ON posts FOR EACH ROW
EXECUTE PROCEDURE update_post_timestamp ();

-- replys テーブルの作成
CREATE TABLE IF NOT EXISTS replys (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    sender_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- レコードの更新時に updatedAt を自動更新するトリガー関数
CREATE OR REPLACE FUNCTION update_reply_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updatedAt = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- replys テーブルの updatedAt フィールドを自動更新するトリガー
CREATE TRIGGER update_replys_updated_at BEFORE
UPDATE ON replys FOR EACH ROW
EXECUTE PROCEDURE update_reply_timestamp ();