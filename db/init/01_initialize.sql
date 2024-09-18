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