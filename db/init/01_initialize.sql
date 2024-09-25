-- users テーブルの作成
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- 自動インクリメントのユーザーID
    name TEXT NOT NULL, -- ユーザー名
    age INT, -- 年齢
    profile TEXT, -- プロフィール
    profile_image_url TEXT, -- プロフィール画像URL
    gender TEXT, -- 性別
    mbti TEXT, -- MBTI（性格診断）
    created_at TIMESTAMPTZ DEFAULT NOW(), -- 作成日時
    updated_at TIMESTAMPTZ DEFAULT NOW() -- 更新日時
);

-- updated_at フィールドを自動更新するトリガー関数
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();  -- 更新された際、updated_at に現在時刻を設定
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- users テーブルの updated_at フィールドを自動更新するトリガー
CREATE TRIGGER update_users_updated_at BEFORE
UPDATE ON users FOR EACH ROW
EXECUTE PROCEDURE update_timestamp ();

-- posts テーブルの作成
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY, -- 自動インクリメントの投稿ID
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE, -- 投稿者のユーザーID
    content TEXT NOT NULL, -- 投稿内容
    category VARCHAR(255), -- カテゴリー（恋愛、面白いなど）
    male_percentage INT DEFAULT 50, -- 男性の割合（初期値は50）
    female_percentage INT DEFAULT 50, -- 女性の割合（初期値は50）
    created_at TIMESTAMPTZ DEFAULT NOW(), -- 作成日時
    updated_at TIMESTAMPTZ DEFAULT NOW() -- 更新日時
);

-- posts テーブルの updated_at フィールドを自動更新するトリガー
CREATE TRIGGER update_posts_updated_at BEFORE
UPDATE ON posts FOR EACH ROW
EXECUTE PROCEDURE update_timestamp ();

-- replies テーブルの作成
CREATE TABLE IF NOT EXISTS replies (
    id SERIAL PRIMARY KEY, -- 自動インクリメントのリプライID
    post_id INT NOT NULL REFERENCES posts (id) ON DELETE CASCADE, -- 投稿ID（外部キー）
    sender_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE, -- 送信者のユーザーID（外部キー）
    content TEXT NOT NULL, -- リプライ内容
    created_at TIMESTAMPTZ DEFAULT NOW(), -- 作成日時
    updated_at TIMESTAMPTZ DEFAULT NOW() -- 更新日時
);

-- replies テーブルの updated_at フィールドを自動更新するトリガー
CREATE TRIGGER update_replies_updated_at BEFORE
UPDATE ON replies FOR EACH ROW
EXECUTE PROCEDURE update_timestamp ();