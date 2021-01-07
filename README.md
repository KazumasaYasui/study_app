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
    - publish_date
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
    - memo
    - status
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
    - urgency
    - priority
    - created_at
    - updated_at
  
## ER図
![ERD](https://user-images.githubusercontent.com/26411908/103839144-132cf600-50d2-11eb-8251-e746e176e796.png)

## APIエンドポイント
- /api/v1/registration
  - POST
- /api/v1/login
  - POST
  - DELETE
- /api/v1/users
  - GET
  - POST
- /api/v1/users/{id}
  - GET
  - POST
  - DELETE
- /api/v1/books
  - GET
  - POST
- /api/v1/books/{id}
  - GET
  - PATCH
  - DELETE
- /api/v1/books/{id}/my_info
  - GET
  - POST
  - PATCH
  - DELETE
- /api/v1/studies
  - GET
  - POST
- /api/v1/studies/{id}
  - GET
  - PATCH
  - DELETE
- /api/v1/books/meta_info
  - GET

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
