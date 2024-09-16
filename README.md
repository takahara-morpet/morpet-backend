## ローカル環境での PostgreSQL 接続および動作確認

ローカルで PostgreSQL と繋げられるように設定済み。動作確認も完了していますが、現時点では自動でテーブルの作成は行えません。詳細は以下をご参照ください。なお、質問がある場合は中安に問い合わせてください。

### 前提条件

- PostgreSQL がローカルにインストールされていること  
  （自分は Homebrew でインストール、バージョンは`14.13`）
- Golang がローカルにインストールされていること
- この PR のブランチをローカルに pull し、コードが揃っていること

### 手順

1. **PostgreSQL の起動**

   ```bash
   brew services start postgresql@14
   ```

2. **Go サーバーの起動**

   ```bash
   go run morpet.go
   ```

3. **PostgreSQL にログイン**

   ```bash
   psql -U me -d development
   ```

   ### 例：データを挿入して、Go 側で確認する

   1. **テーブル作成**  
      ※このテーブルは一時的に Model とは関係ないものです

      ```sql
      CREATE TABLE users (
          id SERIAL PRIMARY KEY,
          name VARCHAR(50) NOT NULL,
          age INT
      );
      ```

   2. **レコード作成**
      ```sql
      INSERT INTO users (name, age) VALUES
      ('Alice', 30),
      ('Bob', 25),
      ('Charlie', 35);
      ```

4. **Go サイドで確認**  
   ブラウザで以下にアクセス：
   http://localhost:8080/users

## API 仕様

API 仕様は SwaggerUI を利用して閲覧します。

```
$ docker compose up swagger-ui
```

を実行することでローカルの Docker 上に SwaggerUI サーバが起動します。<br>
<br>
SwaggerUI サーバ起動後以下の URL から SwaggerUI へアクセスすることができます。

SwaggerUI: <http://localhost:8000/> <br>
定義ファイル: `./api-document.yaml`<br>
