# study_app

## 仕様
- 日々の学習結果を記録できるアプリ
    - 学習対象の本などを作成/参照/編集/削除
        - 本一覧ページ
        - 本作成ページ
    - 学習データを作成/編集/編集/削除
        - 学習データ一覧ページ
    - ユーザーを作成/参照/編集/削除
        - ユーザー一覧ページ
        - ユーザー作成ページ

## テーブル構造
- User
    - id
    - name
    - email
    - password_digest
    - image_url
    - created_at
    - updated_at
- Book
    - id
    - genre_id
    - publisher_id
    - author_id
    - title
    - image_url
    - published_at
    - created_at
    - updated_at
- Genre
    - id
    - name
    - created_at
    - updated_at
- Publisher
  - id
  - name
  - created_at
  - updated_at
- Author
  - id
  - name
  - created_at
  - updated_at
- Study
    - id
    - user_id
    - result
    - execution_date
    - created_at
    - updated_at
- StudyPlanItem
    - id
    - study_id
    - book_id
    - memo
    - created_at
    - updated_at
- StudyResultItem
    - id
    - study_id
    - book_id
    - memo
    - created_at
    - updated_at
- UserBookItem
    - id
    - user_id
    - book_id
    - status
    - progress_rate
    - urgency_number
    - priority_number
    - created_at
    - updated_at
  
## ER図
![ERD](https://user-images.githubusercontent.com/26411908/103175905-0d086f80-48b1-11eb-863b-f9c9479758da.png)

## APIエンドポイント
- /api/v1/books
  - GET
  - POST
- /api/v1/books/{id}
  - GET
  - PATCH
  - DELETE
- /api/v1/users
  - GET
  - POST
- /api/v1/users/{id}
  - GET
  - POST
  - DELETE
- /api/v1/studies
  - GET
  - POST
- /api/v1/studies/{id}
  - GET
  - PATCH
  - DELETE
- /api/v1/login
  - POST
  - DELETE

## 技術スタック
### フロントエンド
- HTML5 / CSS3
- JavaScript / TypeScript
- React / Next.js
- Bootstrap or Material-UI
### バックエンド
- Golang
- echo
- gorm
- OpenAPI
### DB
- PostgreSQL or MySQL
### インフラ
- AWS
    - ECS / Fargate or EKS / Fargate
    - ECR
    - RDS
    - ElastiCache
    - S3
    - VPC
    - ALB
    - Route53
    - ACM
    - SSM
    - CloudWatchLogs
### その他
- Docker / docker-compose
- Kubernetes
- github actions

## インフラ構成図
![infra](https://user-images.githubusercontent.com/26411908/103181339-65586500-48e3-11eb-8a60-c8535015d309.png)
